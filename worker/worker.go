package worker

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/holiman/uint256"
	"github.com/samber/lo"
	"go.uber.org/atomic"

	blockchain "unishard/blockchain"
	byzantine "unishard/byzantine"
	config "unishard/config"
	crypto "unishard/crypto"
	election "unishard/election"
	graph_dependency "unishard/graph_dependency"
	log "unishard/log"
	mempool "unishard/mempool"
	message "unishard/message"
	node "unishard/node"
	pacemaker "unishard/pacemaker"
	pbft "unishard/pbft"
	quorum "unishard/quorum"
	"unishard/snapshot_control_table"
	types "unishard/types"
	utils "unishard/utils"
)

// The node structure
type Replica struct {
	node.Node
	*pbft.PBFT
	*election.Static
	snapshotControlTable snapshot_control_table.SnapshotControlTable
	sct                  map[common.Address]map[common.Hash]*blockchain.SnapshotControlTable
	pd                   *mempool.Producer
	pm                   *pacemaker.Pacemaker
	start                chan bool // signal to start the node
	isStarted            atomic.Bool
	isByz                bool
	timer                *time.Timer // timeout for each view
	committedBlocks      chan *blockchain.WorkerBlock
	forkedBlocks         chan *blockchain.WorkerBlock
	reservedBlock        chan *blockchain.WorkerBlock
	eventChan            chan interface{}
	preparingBlock       *blockchain.WorkerBlock // block waiting for enough lockResponse
	reservedPreparBlock  *blockchain.WorkerBlock // reserved Preparing Block from ViewChange
	thrus                string
	lastViewTime         time.Time
	startTime            time.Time
	tmpTime              time.Time
	voteStart            time.Time
	totalDelay           time.Duration
	totalRoundTime       time.Duration
	totalVoteTime        time.Duration
	receivedNo           int
	roundNo              int
	voteNo               int
	totalCommittedTx     int
	latencyNo            int
	processedNo          int
	committedNo          int
	proposeCandidate     chan *blockchain.WorkerConsensusPayload
}

// Create a new replica instance (node)
func NewReplica(id types.NodeID, alg string, isByz bool, shard types.Shard, dbObject ethdb.Database) *Replica {
	r := new(Replica)
	r.Node = node.NewNode(id, isByz, shard)
	if isByz {
		log.Debugf("[%v] is Byzantine", r.ID())
	}

	if config.GetConfig().Master == "0" {
		r.Static = election.NewStatic(utils.NewNodeID(1), r.Shard(), config.Configuration.CommitteeNumber)
	}
	r.snapshotControlTable = snapshot_control_table.SnapshotControlTable{
		Table: map[common.Address]map[common.Hash]*snapshot_control_table.SnapshotControlEntity{},
		Mutex: sync.RWMutex{},
	}
	r.sct = make(map[common.Address]map[common.Hash]*blockchain.SnapshotControlTable)
	r.isByz = isByz
	r.pd = mempool.NewProducer()
	r.pm = pacemaker.NewPacemaker(config.GetConfig().CommitteeNumber)
	r.start = make(chan bool)
	r.eventChan = make(chan interface{})
	r.committedBlocks = make(chan *blockchain.WorkerBlock, 100)
	r.forkedBlocks = make(chan *blockchain.WorkerBlock, 100)
	r.reservedBlock = make(chan *blockchain.WorkerBlock)
	r.proposeCandidate = make(chan *blockchain.WorkerConsensusPayload, 100)
	r.preparingBlock = nil
	r.reservedPreparBlock = nil

	/* Register to gob en/decoder */
	gob.Register(blockchain.Accept{})
	gob.Register(blockchain.WorkerConsensusPayload{})
	gob.Register(blockchain.ProposedWorkerBlock{})
	gob.Register(blockchain.WorkerBlock{})
	gob.Register(message.Experiment{})
	gob.Register(blockchain.CoordinationBlock{})
	gob.Register(quorum.Collection[quorum.Vote]{})
	gob.Register(quorum.Collection[quorum.Commit]{})
	gob.Register(pacemaker.TC{})
	gob.Register(pacemaker.TMO{})
	gob.Register(byzantine.ReportByzantine{})
	gob.Register(byzantine.ReplaceByzantine{})
	gob.Register(message.ConsensusNodeRegister{})

	/* Register message handlers */
	r.Register(blockchain.Accept{}, r.HandleAccept)
	r.Register(blockchain.WorkerConsensusPayload{}, r.HandleWorkerConsensusPayload)
	r.Register(blockchain.ProposedWorkerBlock{}, r.HandleProposedWorkerBlock)
	r.Register(quorum.Collection[quorum.Vote]{}, r.HandleVote)
	r.Register(quorum.Collection[quorum.Commit]{}, r.HandleCommit)
	r.Register(pacemaker.TMO{}, r.HandleTmo)
	r.Register(pacemaker.TC{}, r.HandleTc)
	r.Register(byzantine.ReportByzantine{}, r.HandleReportByzantine)
	r.Register(byzantine.ReplaceByzantine{}, r.HandleReplaceByzantine)
	r.Register(message.Transaction{}, r.HandleTxn)
	r.Register(message.Query{}, r.handleQuery)
	r.Register(message.RequestLeader{}, r.handleRequestLeader)
	r.Register(message.ReportByzantine{}, r.handleReportByzantine)

	r.PBFT = pbft.NewPBFT(r.Node, r.pm, r.pd, r.Static, r.committedBlocks, r.forkedBlocks, r.reservedBlock, dbObject)
	return r
}

// Process the received worker consensus payload only if the node is the current leader (block proposer)
func (r *Replica) HandleWorkerConsensusPayload(bm blockchain.WorkerConsensusPayload) {
	r.receivedNo++
	r.startSignal()
	if r.pm.GetCurView() == 0 {
		r.pm.AdvanceView(0)
	}

	if !r.IsLeader(r.ID(), r.pm.GetCurView(), r.pm.GetCurEpoch()) {
		return
	}
	r.proposeCandidate <- &bm
}

// Generate a worker block using the received worker consensus payload,
// and then propose (propagate) it to all other nodes
func (r *Replica) ProposeWorkerBlock() {
	for payload := range r.proposeCandidate {
		for _, tx := range payload.Data.ReceivedLocalTransaction {
			r.HandleTxn(*tx)
		}

		// Update SCT
		r.UpdateSCT(payload.Data.GlobalCoordinationSequence)

		// Create local execution sequence
		localExecutionSequence := r.GenerateLocalExecutionSequence(payload.Data.GlobalCoordinationSequence)

		// Replicate external state variables and contracts contained in the received global snapshot
		externalAddress := r.ReplicateExternalStateVariable(payload.Data.GlobalSnapshot, payload.Data.GlobalContractBundle)

		// Execute local execution sequence
		r.PBFT.ExecuteWorkerBlock(payload, localExecutionSequence)

		// Delete the replicated external state variables and contracts
		r.DeleteExternalStateVariable(externalAddress)

		// Create local snapshot and local contract bundle
		localSnapshot, localContractBundle := r.CreateLocalSnapshot(payload.Data.ReceivedCrossTransaction, payload.Data.GC)

		// Build worker block and propagate it
		block := r.PBFT.CreateWorkerBlock(localExecutionSequence, payload.Data.ReceivedCrossTransaction, localSnapshot, localContractBundle, payload.Data.GC)
		block.Proposer = r.ID()
		leaderSignature, _ := crypto.PrivSign(crypto.IDToByte(block.BlockHash), nil)
		block.CommitteeSignature = append(block.CommitteeSignature, leaderSignature)

		workerBlock := blockchain.ProposedWorkerBlock{
			WorkerBlock:          block,
			GlobalSnapshot:       payload.Data.GlobalSnapshot,
			GlobalContractBundle: payload.Data.GlobalContractBundle,
		}
		r.proposeBlock(&workerBlock)
	}
}

// Update the snapshot control table using the read/write set from the received global coordination sequence
func (r *Replica) UpdateSCT(globalCoordinationSequence []*message.Transaction) {
	slot := []common.Hash{}
	addr := []common.Address{}

	for _, ct := range globalCoordinationSequence {
		// Verify the accessed slots and addresses based on the transaction type
		// while removing the corresponding RTCS elements from the snapshot control table
		switch ct.TXType {

		// If it is a fund transfer transaction, update SCT as follows
		case types.TRANSFER:
			slot = append(slot, utils.SlotToKey(0))
			addr = append(addr, ct.From)
			slot = append(slot, utils.SlotToKey(0))
			addr = append(addr, ct.To)
			r.snapshotControlTable.RemoveElementFromRTCS(utils.SlotToKey(0), ct.From, ct.Hash)
			r.snapshotControlTable.RemoveElementFromRTCS(utils.SlotToKey(0), ct.To, ct.Hash)

		// If it invokes a smart contract function, update SCT as follows
		case types.SMARTCONTRACT:
			for _, set := range ct.RwSet {
				for _, rv := range set.ReadSet {
					r.snapshotControlTable.RemoveElementFromRTCS(common.HexToHash(rv), set.Address, ct.Hash)
				}
				for _, wv := range set.WriteSet {
					slot = append(slot, common.HexToHash(wv))
					addr = append(addr, set.Address)
				}
			}
			slot = append(slot, utils.SlotToKey(0))
			addr = append(addr, ct.From)
			r.snapshotControlTable.RemoveElementFromRTCS(utils.SlotToKey(0), ct.From, ct.Hash)
		}
	}

	// Change all accessed snapshots to write-permitted based on the verified slots and addresses
	for i := 0; i < len(slot); i++ {
		r.snapshotControlTable.UnprotectSnapshot(slot[i], addr[i])
		if r.snapshotControlTable.Exist(slot[i], addr[i]) {
		}
	}
}

// Generate the local execution sequence using the received global coordination sequence
func (r *Replica) GenerateLocalExecutionSequence(globalCoordinationSequence []*message.Transaction) (localExecutionSequence []*message.Transaction) {
	localCommitTransaction := make([]common.Hash, 0)
	dependencyMap := make(map[common.Hash]bool)

	// Generate the write-read dependency graph for each cross-shard transaction in the global coordination sequence
	g := graph_dependency.NewGraph()
	setList := graph_dependency.GenerateSet(globalCoordinationSequence)
	for i, set := range setList {
		g.AddVertex(set.TransactionId, i, set.Write_set, set.Read_set)
	}
	g.DrawDependency()

	// Identify the local-commit cross-shard transactions for the shard
	for _, ct := range globalCoordinationSequence {
		isLocalCommit := r.CheckLocalCommit(ct)
		if ct.TXType == types.SMARTCONTRACT {
			if ct.RwSet != nil && isLocalCommit {
				localCommitTransaction = append(localCommitTransaction, ct.Hash)
			}
		} else if ct.TXType == types.TRANSFER {
			if isLocalCommit {
				localCommitTransaction = append(localCommitTransaction, ct.Hash)
			}
		}
	}

	// Identify the causal cross-shard transactions of the local-commit cross-shard transactions
	for _, ct := range globalCoordinationSequence {
		if utils.R59H6YyJ3h(localCommitTransaction, ct.Hash) {
			dependencyMap[ct.Hash] = true
		} else {
			dependencyMap[ct.Hash] = false
		}
	}
	cnt := len(localCommitTransaction)
	for cnt > 0 {
		cnt = 0
		for ct, inLocalSequence := range dependencyMap {
			if inLocalSequence {
				for _, neighbor := range g.Neighbors(ct) {
					if !dependencyMap[neighbor] {
						dependencyMap[neighbor] = true
						cnt++
					}
				}
			}
		}
	}

	// Add the local-commit cross-shard transactions and their causal cross-shard transactions to the local execution sequence
	for _, ct := range globalCoordinationSequence {
		if dependencyMap[ct.Hash] {
			localExecutionSequence = append(localExecutionSequence, ct)
		}
	}

	// Retrieve the received local transactions from the mempool
	localTransaction := r.pd.GeneratePayload()

	// Examine the write protection flags of the state variables accessed by the local transactions.
	for _, tx := range localTransaction {
		isProtected := false
		switch tx.TXType {
		case types.TRANSFER:
			fromShard := utils.CalculateShardToSend([]common.Address{tx.From}, config.GetConfig().ShardCount)[0]
			toShard := utils.CalculateShardToSend([]common.Address{tx.To}, config.GetConfig().ShardCount)[0]
			if fromShard == r.Shard() && r.snapshotControlTable.IsProtected(utils.SlotToKey(0), tx.From) {
				isProtected = true
			} else if toShard == r.Shard() && r.snapshotControlTable.IsProtected(utils.SlotToKey(0), tx.To) {
				isProtected = true
			}
		case types.SMARTCONTRACT:
			for _, rwSet := range tx.RwSet {
				for _, ws := range rwSet.WriteSet {
					snapshotShard := utils.CalculateShardToSend([]common.Address{rwSet.Address}, config.GetConfig().ShardCount)[0]
					if snapshotShard == r.Shard() && r.snapshotControlTable.IsProtected(common.HexToHash(ws), rwSet.Address) {
						isProtected = true
						goto AddOnLocalExecutionSequence
					}
				}
			}
			fromShard := utils.CalculateShardToSend([]common.Address{tx.From}, config.GetConfig().ShardCount)[0]
			if fromShard == r.Shard() && r.snapshotControlTable.IsProtected(utils.SlotToKey(0), tx.From) {
				isProtected = true
				goto AddOnLocalExecutionSequence
			}
		}

	AddOnLocalExecutionSequence:
		// If all accessed state variables are write-permitted, add the transaction to the local execution sequence;
		// otherwise, re-add it to the mempool
		if !isProtected {
			localExecutionSequence = append(localExecutionSequence, tx)
		} else {
			r.pd.AddTxn(tx)
		}
	}

	return localExecutionSequence
}

// Check if the received cross-shard transaction is a local-commit transaction in the respective shard
func (r *Replica) CheckLocalCommit(ct *message.Transaction) bool {
	switch ct.TXType {
	case types.TRANSFER:
		// Check if the 'from' or 'to' address of the fund-transfer transaction belongs to the respective shard
		fromShard := utils.CalculateShardToSend([]common.Address{ct.From}, config.GetConfig().ShardCount)[0]
		toShard := utils.CalculateShardToSend([]common.Address{ct.To}, config.GetConfig().ShardCount)[0]
		if fromShard == r.Shard() || toShard == r.Shard() {
			return true
		}
		return false
	case types.SMARTCONTRACT:
		// Check if the write set address of the smart contract transaction belongs to the respective shard
		for _, rwSet := range ct.RwSet {
			rwSetShard := utils.CalculateShardToSend([]common.Address{rwSet.Address}, config.GetConfig().ShardCount)[0]
			if rwSetShard == r.Shard() && len(rwSet.WriteSet) > 0 {
				return true
			}
		}
		fromShard := utils.CalculateShardToSend([]common.Address{ct.From}, config.GetConfig().ShardCount)[0]
		return fromShard == r.Shard()
	}

	return false
}

// Temporarily replicate the external state variables and contracts contained in the received global snapshot and global contract bundle
func (r *Replica) ReplicateExternalStateVariable(globalSnapshot []*blockchain.LocalSnapshot, globalContractBundle []*blockchain.LocalContract) (externalAddress []*common.Address) {
	stateDB := r.PBFT.GetBlockChain().GetStateDB()
	for _, code := range globalContractBundle {
		if utils.CalculateShardToSend([]common.Address{code.Address}, config.GetConfig().ShardCount)[0] != r.Shard() {
			// Temporarily create external contracts that do not belong to the shard
			stateDB.CreateTemporaryAccount(code.Address)
			stateDB.SetCode(code.Address, code.Code)
			externalAddress = append(externalAddress, &code.Address)
		}
	}
globalSnapshot:
	for _, snapshot := range globalSnapshot {
		if utils.CalculateShardToSend([]common.Address{snapshot.Address}, config.GetConfig().ShardCount)[0] != r.Shard() {
			isForeignSmartContract := func(snapshotAddress common.Address) bool {
				for _, externalAddr := range externalAddress {
					if *externalAddr == snapshotAddress {
						return true
					}
				}
				return false
			}(snapshot.Address)

			if isForeignSmartContract {
				// Update stateDB if the global snapshot contains state variables within external contracts
				for _, externalAddr := range externalAddress {
					if *externalAddr == snapshot.Address {
						stateDB.SetState(snapshot.Address, snapshot.Slot, common.HexToHash(snapshot.Value))
						continue globalSnapshot
					}
				}
			} else {
				// Create temporary account and update stateDB if the global snapshot contains state variables of external user accounts
				stateDB.CreateTemporaryAccount(snapshot.Address)
				externalAddress = append(externalAddress, &snapshot.Address)
				balance, _ := strconv.ParseInt(snapshot.Value, 0, 64)
				stateDB.SetBalance(snapshot.Address, uint256.NewInt(uint64(balance)))
				stateDB.SetNonce(snapshot.Address, 0)
			}
		}
	}

	return externalAddress
}

// Delete the replicated external state variables and contracts
func (r *Replica) DeleteExternalStateVariable(externalAddress []*common.Address) {
	for _, externalAddr := range externalAddress {
		r.PBFT.GetBlockChain().GetStateDB().DeleteAccount(*externalAddr)
	}
}

// Create local snapshot based on the received cross-shard transactions
func (r *Replica) CreateLocalSnapshot(receivedCrossTransaction []*message.Transaction, gc types.BlockHeight) (localSnapshot []*blockchain.LocalSnapshot, localContractBundle []*blockchain.LocalContract) {
	// If the RTCS in the SCT is non-empty and the protected flag is false,
	// include it in the local snapshot and update the SCT with the current state value
	stateDB := r.PBFT.GetBlockChain().GetStateDB()
	slot, addr := r.snapshotControlTable.GetNonEmptyRTCS()
	for i := 0; i < len(slot); i++ {
		if !r.snapshotControlTable.IsProtected(slot[i], addr[i]) {
			currentSnapshot := r.snapshotControlTable.FindSnapshot(slot[i], addr[i])

			// Capture the current value as a snapshot
			currentValue := ""
			if stateDB.GetCode(addr[i]) == nil {
				currentValue = stateDB.GetBalance(addr[i]).String()
			} else {
				currentValue = stateDB.GetState(addr[i], slot[i]).String()
			}

			// Update the SCT and add it to the local snapshot
			r.snapshotControlTable.UpdateSnapshot(slot[i], addr[i], currentValue, currentSnapshot.RTCS, gc, true)
			localSnapshot = append(localSnapshot, r.snapshotControlTable.FindSnapshot(slot[i], addr[i]))
		}
	}

	// Verify the read set of the received cross-shard transaction and include it in the local snapshot
	for _, ct := range receivedCrossTransaction {
		switch ct.TXType {
		case types.TRANSFER:
			// Verify if the 'from' or 'to' address of the fund-transfer transaction belongs to the respective shard,
			// update the SCT for the corresponding address, and add it to the local snapshot
			fromShard := utils.CalculateShardToSend([]common.Address{ct.From}, config.GetConfig().ShardCount)[0]
			toShard := utils.CalculateShardToSend([]common.Address{ct.To}, config.GetConfig().ShardCount)[0]
			if fromShard == r.Shard() {
				rtcs := []common.Hash{}
				if r.snapshotControlTable.Exist(utils.SlotToKey(0), ct.From) {
					currentSnapshot := r.snapshotControlTable.FindSnapshot(utils.SlotToKey(0), ct.From)
					rtcs = append(rtcs, currentSnapshot.RTCS...)
					rtcs = append(rtcs, ct.Hash)
				} else {
					rtcs = append(rtcs, ct.Hash)
				}
				currentValue := stateDB.GetBalance(ct.From).String()
				r.snapshotControlTable.UpdateSnapshot(utils.SlotToKey(0), ct.From, currentValue, rtcs, gc, true)
				localSnapshot = append(localSnapshot, r.snapshotControlTable.FindSnapshot(utils.SlotToKey(0), ct.From))
			}
			if toShard == r.Shard() {
				rtcs := []common.Hash{}
				if r.snapshotControlTable.Exist(utils.SlotToKey(0), ct.To) {
					currentSnapshot := r.snapshotControlTable.FindSnapshot(utils.SlotToKey(0), ct.To)
					rtcs = append(rtcs, currentSnapshot.RTCS...)
					rtcs = append(rtcs, ct.Hash)
				} else {
					rtcs = append(rtcs, ct.Hash)
				}
				currentValue := stateDB.GetBalance(ct.To).String()
				r.snapshotControlTable.UpdateSnapshot(utils.SlotToKey(0), ct.To, currentValue, rtcs, gc, true)
				localSnapshot = append(localSnapshot, r.snapshotControlTable.FindSnapshot(utils.SlotToKey(0), ct.To))
			}
		case types.SMARTCONTRACT:
			// Verify if the read set address of the smart contract transaction belongs to the respective shard,
			// update the SCT for the corresponding address, and add it to the local snapshot
			for _, rwSet := range ct.RwSet {
				snapshotShard := utils.CalculateShardToSend([]common.Address{rwSet.Address}, config.GetConfig().ShardCount)[0]
				if snapshotShard == r.Shard() {
					for _, rs := range rwSet.ReadSet {
						rtcs := []common.Hash{}
						if r.snapshotControlTable.Exist(common.HexToHash(rs), rwSet.Address) {
							currentSnapshot := r.snapshotControlTable.FindSnapshot(common.HexToHash(rs), rwSet.Address)
							rtcs = append(rtcs, currentSnapshot.RTCS...)
							rtcs = append(rtcs, ct.Hash)
						} else {
							rtcs = append(rtcs, ct.Hash)
						}
						currentValue := stateDB.GetState(rwSet.Address, common.HexToHash(rs)).String()
						r.snapshotControlTable.UpdateSnapshot(common.HexToHash(rs), rwSet.Address, currentValue, rtcs, gc, true)
						localSnapshot = append(localSnapshot, r.snapshotControlTable.FindSnapshot(common.HexToHash(rs), rwSet.Address))
					}

					// Add the code of the contracts belonging to the shard
					// from the received smart contract cross-shard transactions to the contract bundle
					localContract := &blockchain.LocalContract{
						Address: rwSet.Address,
						Code:    stateDB.GetCode(rwSet.Address),
					}
					localContractBundle = append(localContractBundle, localContract)

				}
			}
			// In the case of a smart contract transaction, gas is consumed from the ct.From address,
			// so update the SCT for this address and add it to the local snapshot
			fromShard := utils.CalculateShardToSend([]common.Address{ct.From}, config.GetConfig().ShardCount)[0]
			if fromShard == r.Shard() {
				rtcs := []common.Hash{}
				if r.snapshotControlTable.Exist(utils.SlotToKey(0), ct.From) {
					currentSnapshot := r.snapshotControlTable.FindSnapshot(utils.SlotToKey(0), ct.From)
					rtcs = append(rtcs, currentSnapshot.RTCS...)
					rtcs = append(rtcs, ct.Hash)
				} else {
					rtcs = append(rtcs, ct.Hash)
				}
				currentValue := stateDB.GetBalance(ct.From).String()
				r.snapshotControlTable.UpdateSnapshot(utils.SlotToKey(0), ct.From, currentValue, rtcs, gc, true)
				localSnapshot = append(localSnapshot, r.snapshotControlTable.FindSnapshot(utils.SlotToKey(0), ct.From))
			}
		}
	}

	// Remove duplicates from the local snapshot
	localSnapshot = lo.UniqBy[*blockchain.LocalSnapshot, string](localSnapshot, func(item *blockchain.LocalSnapshot) string {
		return item.Address.Hex() + item.Slot.Hex()
	})

	// Remove duplicates from the local contract bundle
	localContractBundle = lo.UniqBy[*blockchain.LocalContract, string](localContractBundle, func(item *blockchain.LocalContract) string {
		return item.Address.Hex()
	})

	return localSnapshot, localContractBundle
}

// The following functions are used for BFT consensus within a shard
// They implement PBFT algorithm

func (r *Replica) HandleProposedWorkerBlock(workerBlock blockchain.ProposedWorkerBlock) {
	r.receivedNo++
	r.startSignal()
	if r.pm.GetCurView() == 0 {
		log.Debugf("[%v] start protocol ", r.ID())
		r.pm.AdvanceView(0)
	}
	block := workerBlock.WorkerBlock
	if !r.IsCommittee(r.ID(), block.BlockHeader.Epoch) {
		return
	}
	r.eventChan <- workerBlock
}

func (r *Replica) HandleAccept(accept blockchain.Accept) {
	r.startSignal()
	if r.pm.GetCurView() == 0 {
		r.pm.AdvanceView(0)
	}

	commitedBlock := accept.CommittedBlock
	if commitedBlock.BlockHeader.BlockHeight < r.GetLastBlockHeight() {
		return
	}
	r.eventChan <- accept
}

func (r *Replica) HandleTxn(m message.Transaction) {
	r.startSignal()
	if r.pm.GetCurView() == 0 {
		r.pm.AdvanceView(0)
	}

	if m.TXType == types.ABORT {
		epoch, view := r.pm.GetCurEpochView()
		if !r.IsLeader(r.ID(), view, epoch) {
			r.Send(r.FindLeaderFor(view, epoch), m)
			return
		}
		r.pd.CollectTxn(&m)
	} else {
		r.pd.AddTxn(&m)
	}
}

func (r *Replica) HandleVote(vote quorum.Collection[quorum.Vote]) {
	r.startSignal()
	if r.State() == types.VIEWCHANGING {
		return
	}

	if !r.IsCommittee(r.ID(), vote.Epoch) {
		return
	}

	if vote.BlockHeight < r.GetHighBlockHeight() {
		return
	}

	r.eventChan <- vote
}

func (r *Replica) HandleCommit(commitMsg quorum.Collection[quorum.Commit]) {
	r.startSignal()
	if r.State() == types.VIEWCHANGING {
		return
	}

	if !r.IsCommittee(r.ID(), commitMsg.Epoch) {
		return
	}

	if commitMsg.BlockHeight < r.GetLastBlockHeight() {
		return
	}

	r.eventChan <- commitMsg
}

func (r *Replica) HandleTmo(tmo pacemaker.TMO) {
	if tmo.QmJZRap < r.pm.GetCurView() {
		return
	}

	if !r.IsCommittee(r.ID(), tmo.Epoch) {
		return
	}

	r.eventChan <- tmo
}

func (r *Replica) HandleTc(tc pacemaker.TC) {
	if !(tc.NewView > r.pm.GetCurView()) {
		return
	}
	r.eventChan <- tc
}

func (r *Replica) HandleReportByzantine(reportByzantine byzantine.ReportByzantine) {
	r.startSignal()
	r.eventChan <- reportByzantine
}

func (r *Replica) HandleReplaceByzantine(replaceByzantine byzantine.ReplaceByzantine) {
	r.startSignal()
	r.eventChan <- replaceByzantine
}

func (r *Replica) handleQuery(m message.Query) {
	latency := float64(r.totalDelay.Milliseconds()) / float64(r.latencyNo)
	r.thrus += fmt.Sprintf("Time: %v s. Throughput: %v txs/s\n", time.Since(r.startTime).Seconds(), float64(r.totalCommittedTx)/time.Since(r.tmpTime).Seconds())
	r.totalCommittedTx = 0
	r.tmpTime = time.Now()
	status := fmt.Sprintf("Latency: %v\n%s", latency, r.thrus)
	m.Reply(message.QueryReply{Info: status})
}

// handleRequestLeader replies a Leader of the node (new leader for view change)
func (r *Replica) handleRequestLeader(m message.RequestLeader) {
	r.startSignal()
	if r.pm.GetCurView() == 0 {
		r.pm.AdvanceView(0)
	}

	leader := r.FindLeaderFor(r.pm.GetCurView(), r.pm.GetCurEpoch())
	m.Reply(message.RequestLeaderReply{Leader: fmt.Sprint(leader)})
}

func (r *Replica) handleReportByzantine(m message.ReportByzantine) {
	epoch, view := r.pm.GetCurEpochView()
	committees := r.FindCommitteesFor(epoch)
	log.Debug(committees)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var publicKey types.NodeID
	for {
		publicKey = committees[rnd.Intn(utils.UB_qaFnx(committees))]
		if !r.IsLeader(publicKey, view, epoch) {
			break
		}
	}
	reportByzantine := byzantine.MakeReportByzantine(publicKey, r.Shard(), epoch, view)

	leader := r.Static.FindLeaderFor(view, epoch)
	if leader == r.ID() {
		r.HandleReportByzantine(*reportByzantine)
	} else {
		r.Send(leader, reportByzantine)
	}
}

func (r *Replica) processNewView(curEpoch types.Epoch, curView types.View) {

	if !r.IsLeader(r.ID(), curView, curEpoch) {
		return
	}

	r.SetRole(types.LEADER)
	r.eventChan <- types.EpochView{Epoch: curEpoch, View: curView}
}

func (r *Replica) proposeBlock(blockWithGlobalVariable *blockchain.ProposedWorkerBlock) {
	block := blockWithGlobalVariable.WorkerBlock
	r.BroadcastToSome(r.FindCommitteesFor(block.BlockHeader.Epoch), blockWithGlobalVariable)
	_ = r.PBFT.ProcessWorkerBlock(blockWithGlobalVariable)
	r.voteStart = time.Now()
	r.preparingBlock = nil
}

func (r *Replica) processCommittedBlock(block *blockchain.WorkerBlock) {
	r.committedNo++
	if !r.IsLeader(r.ID(), block.BlockHeader.View, block.BlockHeader.Epoch) {
		return
	}
}

func (r *Replica) processForkedBlock(block *blockchain.WorkerBlock) {}

func (r *Replica) ListenCommittedBlocks() {
	for {
		select {
		case committedBlock := <-r.committedBlocks:
			r.processCommittedBlock(committedBlock)
		case forkedBlock := <-r.forkedBlocks:
			r.processForkedBlock(forkedBlock)
		case reservedPrepareBlock := <-r.reservedBlock:
			r.reservedPreparBlock = reservedPrepareBlock
		}
	}
}

func (r *Replica) ListenLocalEvent() {
	r.lastViewTime = time.Now()
	r.timer = time.NewTimer(r.pm.GetTimerForView())
	for {
		if r.Role() == types.VALIDATOR {
			r.timer.Stop()
		} else {
			switch r.State() {
			case types.READY:
				r.timer.Reset(r.pm.GetTimerForView())
			case types.VIEWCHANGING:
				r.timer.Reset(r.pm.GetTimerForViewChange())
			}
		}
	L:
		for {
			select {
			case EpochView := <-r.pm.EnteringViewEvent():
				var (
					epoch = EpochView.Epoch
					view  = EpochView.View
				)
				if view >= 2 {
					r.totalVoteTime += time.Since(r.voteStart)
				}
				now := time.Now()
				lasts := now.Sub(r.lastViewTime)
				r.totalRoundTime += lasts
				r.roundNo++
				r.lastViewTime = now
				if r.IsCommittee(r.ID(), epoch) {
					r.SetRole(types.COMMITTEE)
				} else {
					r.SetRole(types.VALIDATOR)
				}
				r.processNewView(epoch, view)
				break L
			case <-r.pm.EnteringTmoEvent():
				r.PBFT.ProcessLocalTmo(r.pm.GetCurView())
				break L
			case <-r.timer.C:
				if r.State() == types.VIEWCHANGING {
					r.pm.UpdateNewView(r.pm.GetNewView() + 1)
				}
				r.PBFT.ProcessLocalTmo(r.pm.GetCurView())

				break L
			}
		}
	}
}

func (r *Replica) startSignal() {
	if !r.isStarted.Load() {
		r.startTime = time.Now()
		r.tmpTime = time.Now()
		r.isStarted.Store(true)
		r.start <- true
	}
}

// Start event loop
func (r *Replica) Start() {
	go r.Run()
	go r.ProposeWorkerBlock()
	r.ElectCommittees(crypto.ZsUswPaGG4Z(0), 1)
	if r.IsCommittee(r.ID(), 1) {
		r.SetRole(types.COMMITTEE)
	}
	// wait for the start signal
	<-r.start
	go r.ListenLocalEvent()
	go r.ListenCommittedBlocks()
	for {
		for r.isStarted.Load() {
			event := <-r.eventChan
			switch v := event.(type) {
			case types.EpochView:
			case blockchain.ProposedWorkerBlock:
				_ = r.PBFT.ProcessWorkerBlock(&v)
				r.voteStart = time.Now()
				r.processedNo++
			case quorum.Collection[quorum.Vote]:
				startProcessTime := time.Now()
				r.PBFT.ProcessVote(&v)
				processingDuration := time.Since(startProcessTime)
				r.totalVoteTime += processingDuration
				r.voteNo++
			case quorum.Collection[quorum.Commit]:
				r.PBFT.ProcessCommit(&v)
			case blockchain.Accept:
				r.PBFT.ProcessAccept(&v)
			case pacemaker.TMO:
				r.PBFT.ProcessRemoteTmo(&v)
			case pacemaker.TC:
				r.PBFT.ProcessTC(&v)
			case byzantine.ReportByzantine:
				r.PBFT.ProcessReportByzantine(&v)
			case byzantine.ReplaceByzantine:
				r.PBFT.ProcessReplaceByzantine(&v)
			}
		}
	}
}
