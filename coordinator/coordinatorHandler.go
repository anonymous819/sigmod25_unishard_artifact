package coordinator

import (
	"encoding/gob"
	"encoding/json"
	"sync"
	"time"

	blockchain "unishard/blockchain"
	config "unishard/config"
	message "unishard/message"
	transport "unishard/transport"
	types "unishard/types"
	utils "unishard/utils"
)

type (
	CoordinationBlockBuilder struct {
		CoordBlockBuilderReplica
		WorkerBuilderList       map[types.Shard]string
		consensusNodeTransport  map[types.NodeID]transport.Transport
		workerBuilderTransport  map[types.Shard]transport.Transport
		checkSnapshotSafetyLock sync.Mutex
	}
)

func NewInitCoordinationBlockBuilder(ip string, shard types.Shard, gatewayAddress string) *CoordinationBlockBuilder {
	cbb := CoordinationBlockBuilder{
		CoordBlockBuilderReplica: *NewCoordinationBlockBuilderReplica(ip, shard, gatewayAddress),
		WorkerBuilderList:        make(map[types.Shard]string),
		workerBuilderTransport:   make(map[types.Shard]transport.Transport),
		consensusNodeTransport:   make(map[types.NodeID]transport.Transport),
		checkSnapshotSafetyLock:  sync.Mutex{},
	}

	/* Register to gob en/decoder */
	gob.Register(message.WorkerBuilderRegister{})
	gob.Register(message.WorkerBuilderListRequest{})
	gob.Register(message.ConsensusNodeRegister{})
	gob.Register(blockchain.WorkerBlock{})
	gob.Register(blockchain.CoordinationBlock{})
	gob.Register(blockchain.CoordinationBlockWithoutHeader{})
	gob.Register(message.ClientStart{})

	/* Register message handler */
	cbb.Register(message.WorkerBuilderRegister{}, cbb.handleWorkerBuilderRegister)
	cbb.Register(message.WorkerBuilderListRequest{}, cbb.handleWorkerBuilderListRequest)
	cbb.Register(message.ConsensusNodeRegister{}, cbb.handleConsensusNodeRegister)
	cbb.Register(blockchain.WorkerBlock{}, cbb.handleWorkerBlock)
	cbb.Register(blockchain.CoordinationBlock{}, cbb.handleCoordinationBlock)

	return &cbb
}

// Main loop executed by the coordinator shard
func (cbb *CoordinationBlockBuilder) Start() {
	go func() {
		go cbb.CoordBlockBuilderReplica.Run()
		err := utils.Retry(cbb.CoordBlockBuilderReplica.gatewaynodeTransport.Dial, 100, time.Duration(50)*time.Millisecond)
		if err != nil {
			panic(err)
		}

		// Generate a block every 0.25 seconds
		interval := time.NewTicker(time.Millisecond * 250)

		// Initiate BFT consensus on the generated block
		for range interval.C {
			if cbb.CoordBlockBuilderReplica.GoFlag {
				cbb.CoordBlockBuilderReplica.ProposeCoordinatorBlock()
			}
		}
	}()

	// loop for handling received messages
	go func() {
		for {
			select {
			case msg := <-cbb.WorkerBuilderTopic:
				if v, ok := msg.(blockchain.WorkerBlock); ok {
					go func() {
						json.Marshal(v)
						cbb.checkSnapshotSafetyLock.Lock()
						// Verify the safety of the local snapshot within the received worker block, and if it is safe, add the snapshot to the SCT
						for _, ls := range v.BlockData.LocalSnapshot {
							safe := cbb.CheckSnapshotSafety(v.BlockHeader.GC, ls)
							if safe {
								cbb.snapshotControlTable.UpdateSnapshot(
									ls.Slot,
									ls.Address,
									ls.Value,
									ls.RTCS,
									v.BlockHeader.BlockHeight,
									false,
								)
							}
						}
						cbb.checkSnapshotSafetyLock.Unlock()
						// Store the cross-shard transactions and contract code contained in the received worker block
						cbb.AccumulateCrossTransaction(v.BlockData.ReceivedCrossTransaction...)
						cbb.AccumulateSmartContractBundle(v.BlockData.LocalContractBundle)
					}()
				}
			case msg := <-cbb.CoordinationBuilderTopic:
				switch v := (msg).(type) {
				case message.WorkerBuilderRegister:
					cbb.WorkerBuilderList[v.SenderShard] = v.Address
					cbb.workerBuilderTransport[v.SenderShard] = transport.NewTransport(v.Address)
					err := utils.Retry(cbb.workerBuilderTransport[v.SenderShard].Dial, 100, time.Duration(50)*time.Millisecond)
					if err != nil {
						panic(err)
					}
				case message.WorkerBuilderListRequest:
					WorkerBuilderList, _ := json.Marshal(message.WorkerBuilderListResponse{
						Builders: cbb.WorkerBuilderList,
					})
					if err := cbb.Client.PutGateway("/WorkerBuilderListResponse", WorkerBuilderList); err != nil {
						panic(err)
					}
				case blockchain.CoordinationBlock:
					for _, ct := range v.BlockData.GlobalCoordinationSequence {
						ct.LatencyDissection.CoordinationConsensusTime = time.Now().UnixMilli()
					}
					for _, t := range cbb.workerBuilderTransport {
						t.Send(v)
					}
				case blockchain.CoordinationBlockWithoutHeader:
					for _, t := range cbb.consensusNodeTransport {
						t.Send(v)
					}
				default:
				}

			case msg := <-cbb.CoordBlockBuilderReplica.ConsensusNodeTopic:
				if v, ok := (msg).(message.ConsensusNodeRegister); ok {
					t := transport.NewTransport(v.IP)
					err := utils.Retry(t.Dial, 100, time.Duration(50)*time.Millisecond)
					if err != nil {
						panic(err)
					}
					cbb.consensusNodeTransport[v.ConsensusNodeID] = t
					if len(cbb.consensusNodeTransport) == config.GetConfig().CommitteeNumber {
						cbb.GoFlag = true
						cbb.gatewaynodeTransport.Send(message.ClientStart{
							Shard: cbb.GetShard(),
						})
					}
				}
			case <-cbb.GatewayTopic:
				continue
			}
		}
	}()
}

func (r *CoordBlockBuilderReplica) handleWorkerBuilderRegister(msg message.WorkerBuilderRegister) {
	r.CoordinationBuilderTopic <- msg
}

func (r *CoordBlockBuilderReplica) handleWorkerBuilderListRequest(msg message.WorkerBuilderListRequest) {
	r.CoordinationBuilderTopic <- msg
}

func (r *CoordBlockBuilderReplica) handleConsensusNodeRegister(msg message.ConsensusNodeRegister) {
	r.ConsensusNodeTopic <- msg
}

func (r *CoordBlockBuilderReplica) handleWorkerBlock(msg blockchain.WorkerBlock) {
	// Measure for latency breakdown
	for _, ct := range msg.BlockData.ReceivedCrossTransaction {
		ct.LatencyDissection.Network1 = time.Now().UnixMilli()
	}
	r.WorkerBuilderTopic <- msg
}

func (r *CoordBlockBuilderReplica) handleCoordinationBlock(msg blockchain.CoordinationBlock) {
	r.GetBlockChain().AddCoordinationBlock(&msg)
	r.CoordinationBuilderTopic <- msg
}
