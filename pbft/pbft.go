//line :1
package pbft

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	blockchain "unishard/blockchain"
	byzantine "unishard/byzantine"
	config "unishard/config"
	crypto "unishard/crypto"
	election "unishard/election"
	evm "unishard/evm"
	state "unishard/evm/state"
	runtime "unishard/evm/vm/runtime"
	log "unishard/log"
	"unishard/measure"
	mempool "unishard/mempool"
	message "unishard/message"
	node "unishard/node"
	pacemaker "unishard/pacemaker"
	quorum "unishard/quorum"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/holiman/uint256"
	"github.com/samber/lo"
)

const FORK = "fork"

type PBFT struct {
	node.Node
	*election.Static
	pm                   *pacemaker.Pacemaker
	pd                   *mempool.Producer
	lastVotedBlockHeight types.BlockHeight
	preferredBlockHeight types.BlockHeight
	highQC               *quorum.HPOWa1PT0xzq
	highCQC              *quorum.HPOWa1PT0xzq
	bc                   *blockchain.Y3t7C8s0m
	voteQuorum           *quorum.FZT207R[quorum.Vote]
	commitQuorum         *quorum.FZT207R[quorum.Commit]
	committedBlocks      chan *blockchain.WorkerBlock
	forkedBlocks         chan *blockchain.WorkerBlock
	reservedBlock        chan *blockchain.WorkerBlock
	updatedQC            chan *quorum.HPOWa1PT0xzq
	lastCreatedBlockQC   *quorum.HPOWa1PT0xzq
	bufferedQCs          map[common.Hash]*quorum.HPOWa1PT0xzq
	bufferedCQCs         map[common.Hash]*quorum.HPOWa1PT0xzq
	bufferedBlocks       map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.ProposedWorkerBlock
	bufferedAccepts      map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.Accept
	agreeingBlocks       map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.ProposedWorkerBlock
	executedState        map[types.BlockHeight]*state.StateDB
	mu                   sync.Mutex
	gSnapshot            []*blockchain.LocalSnapshot
	gCode                []*blockchain.LocalContract
	fdcm                 []*common.Address
	pbftMeasure          *measure.PBFTMeasure

	detectedTmos map[types.NodeID]*pacemaker.TMO
}

func NewPBFT(
	sh_ipJP6 node.Node, yapKe2jjkRuB *pacemaker.Pacemaker, ftImr220 *mempool.Producer,
	keasWXics5qx *election.Static, l2pHR1ZhFG chan *blockchain.WorkerBlock,
	gDcB23PkLL chan *blockchain.WorkerBlock,
	eb1fVADHfFP chan *blockchain.WorkerBlock,
	iXgpvsBPoj ethdb.Database) *PBFT {
	cu1RSg2YR := /*line Lt237b_M73KV.go:1*/ new(PBFT)
	cu1RSg2YR.Node = sh_ipJP6
	cu1RSg2YR.Static = keasWXics5qx
	cu1RSg2YR.pm = yapKe2jjkRuB
	cu1RSg2YR.pd = ftImr220
	cu1RSg2YR.bc = /*line qIr_3fR.go:1*/ blockchain.O8OnUJH0nc2( /*line EfKcgp4b.go:1*/ cu1RSg2YR.Shard() /*line KxSWuq3.go:1*/, config.GetConfig().DefaultBalance, iXgpvsBPoj)
	cu1RSg2YR.voteQuorum = /*line IiVb2umlJPL.go:1*/ quorum.Dq3TFZ[quorum.Vote]( /*line NrRPEVI.go:1*/ config.GetConfig().CommitteeNumber)
	cu1RSg2YR.commitQuorum = /*line AO5MCs.go:1*/ quorum.Dq3TFZ[quorum.Commit]( /*line JBA0r_.go:1*/ config.GetConfig().CommitteeNumber)
	cu1RSg2YR.bufferedBlocks = /*line Vqa0sO_RLwRJ.go:1*/ make(map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.ProposedWorkerBlock)
	cu1RSg2YR.bufferedAccepts = /*line JPqxAjMq4XaG.go:1*/ make(map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.Accept)
	cu1RSg2YR.agreeingBlocks = /*line eNkdlnid.go:1*/ make(map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.ProposedWorkerBlock)
	cu1RSg2YR.executedState = /*line O3sGGZaTQM_M.go:1*/ make(map[types.BlockHeight]*state.StateDB)
	cu1RSg2YR.bufferedQCs = /*line pyp1Ha4.go:1*/ make(map[common.Hash]*quorum.HPOWa1PT0xzq)
	cu1RSg2YR.highQC = &quorum.HPOWa1PT0xzq{LwVL87: 0}
	cu1RSg2YR.bufferedCQCs = /*line XSO5lzYG6.go:1*/ make(map[common.Hash]*quorum.HPOWa1PT0xzq)
	cu1RSg2YR.highCQC = &quorum.HPOWa1PT0xzq{LwVL87: 0}
	cu1RSg2YR.committedBlocks = l2pHR1ZhFG
	cu1RSg2YR.forkedBlocks = gDcB23PkLL
	cu1RSg2YR.reservedBlock = eb1fVADHfFP
	cu1RSg2YR.lastCreatedBlockQC = &quorum.HPOWa1PT0xzq{LwVL87: 0}
	cu1RSg2YR.updatedQC = /*line BPMR_6_oEv.go:1*/ make(chan *quorum.HPOWa1PT0xzq, 1)

	cu1RSg2YR.gSnapshot = /*line t7c5cGtCPg.go:1*/ make([]*blockchain.LocalSnapshot, 0)
	cu1RSg2YR.gCode = /*line RL0xVRukPTY.go:1*/ make([]*blockchain.LocalContract, 0)
	cu1RSg2YR.fdcm = nil
	cu1RSg2YR.detectedTmos = /*line i8M07fiy.go:1*/ make(map[types.NodeID]*pacemaker.TMO)
	cu1RSg2YR.pbftMeasure = /*line HvDa26TYVmXa.go:1*/ measure.NewPBFTMeasure()

	return cu1RSg2YR
}

func (cu1RSg2YR *PBFT) GetBlockChain() *blockchain.Y3t7C8s0m {
	return cu1RSg2YR.bc
}

func (cu1RSg2YR *PBFT) ExecuteWorkerBlock(vf8hbwv *blockchain.WorkerConsensusPayload, duD12ISu []*message.Transaction) {
	cu1RSg2YR.gSnapshot = vf8hbwv.Data.GlobalSnapshot
	cu1RSg2YR.gCode = vf8hbwv.Data.GlobalContractBundle

	for _, bBrNWS3ghg := range vf8hbwv.Data.ReceivedCrossTransaction {
		bBrNWS3ghg.LatencyDissection.WBBWaitingTime = /*line tsUPKFGV9_.go:1*/ time.Now().UnixMilli()
	}

	for _, bQZ9a4xL := range duD12ISu {
		if bQZ9a4xL.IsCrossShardTx {
			bQZ9a4xL.LatencyDissection.WBBWaitingTime2 = /*line Q4P0P1Y.go:1*/ time.Now().UnixMilli()
		} else {
			bQZ9a4xL.LatencyDissection.WBBWaitingTime = /*line GlKEjEf7KVl.go:1*/ time.Now().UnixMilli()
		}
	}

	if /*line padWOuc.go:1*/ cu1RSg2YR.pbftMeasure.StartTime.IsZero() && /*line aSW3fV6D.go:1*/ len(duD12ISu) > 0 {
		cu1RSg2YR.pbftMeasure.StartTime = /*line u7gCu0e.go:1*/ time.Now()
	}

	/*line w7P0aU.go:1*/
	cu1RSg2YR.vhgIe3tbXH(duD12ISu, true)
}

func (cu1RSg2YR *PBFT) CreateWorkerBlock(duD12ISu, gapFR6 []*message.Transaction, thxhyC_T9ga []*blockchain.LocalSnapshot, ia9eCe8vlOON []*blockchain.LocalContract, Gmwh551Kb types.BlockHeight) *blockchain.WorkerBlock {

	hHuUreZez := /*line U7KRaHto_.go:1*/ blockchain.BC1Hpz9(duD12ISu, gapFR6, thxhyC_T9ga, ia9eCe8vlOON)
	s9tPOLmew := /*line GNtHdXcISrPl.go:1*/ new(blockchain.Y7M2DYD)
	s9tPOLmew.Uy2puNVLTPk = hHuUreZez

	var cWdSpmiBLOz *quorum.HPOWa1PT0xzq
	if cu1RSg2YR.highQC.LwVL87 == 0 {
		cWdSpmiBLOz = cu1RSg2YR.highQC
	} else if cu1RSg2YR.highQC.BlockHeight == cu1RSg2YR.lastCreatedBlockQC.BlockHeight {
		cWdSpmiBLOz = <-cu1RSg2YR.updatedQC
	} else {
		cWdSpmiBLOz = cu1RSg2YR.highQC
		<-cu1RSg2YR.updatedQC
	}

	ouJpb2aF3mlv := /*line jehqgA5oIUWD.go:1*/ blockchain.BRx2UrRXo(s9tPOLmew /*line e_coZ3.go:1*/, cu1RSg2YR.pm.GetCurEpoch() /*line J7_veroP4qw.go:1*/, cu1RSg2YR.pm.GetCurView(), Gmwh551Kb /*line VNxidxNt1.go:1*/, cu1RSg2YR.bc.GetCurrentStateHash(), cWdSpmiBLOz.ZWsU_2 /*line GANWxr.go:1*/, cu1RSg2YR.GetHighBlockHeight(), cWdSpmiBLOz /*line rnmLSgmm.go:1*/, cu1RSg2YR.Shard())
	ouVUaI5j8OTr := ouJpb2aF3mlv.(*blockchain.WorkerBlock)
	cu1RSg2YR.lastCreatedBlockQC = ouVUaI5j8OTr.O7aFnBRNh
	/*line BJu1w3nfO6r.go:1*/ log.Debugf(func() /*line Vx5U4uCM_wM.go:1*/ string {
		key := [] /*line Vx5U4uCM_wM.go:1*/ byte("\xb0\xff\x165\xd7\xc3竿\xa4._\x16E\xbe\xac\x86\x138\xea\xf6kK\xfa/\xafh_za\x88`߶\x1a?9\xdc:\xe0\x7f즇\xaf\xe5\xd4DN\xe0\x06\x97@;\x03r\xd3;\xa7V\x18\x8e\x18\xb8\xbf\xa2\xceu\x88\xdbb\xe7ПAM:\x10\xff\x97L\xa8\xa0G\x9d\xa1\xa3]\xb3\x84\v\x1b\xf6\x99\xf1\x1cVI\x14\xdcȅi3:(\x05\xd5O\x11\x846~IMgkd\x11\xca\xd6\xf1\x15\x1aK\xb3\x97\x97")
		data := [] /*line Vx5U4uCM_wM.go:1*/ byte("\v$\x8c\x92\xf7\xeb*\x1d$\x05\xa2\xc4m\xb40\x17\xeb\x85zVeζ#O\x15\xd1\xcd\xe3\xd4\xf0\xc5Cև\xa0\xa4E\xa8G\x9fM\xc6\xf7!TD\xb3\xc1Ar\xb7\xa6\xaau\x925\xa7\x16\xb9\x83\xf6}!&\nB\x95\xadQ\x82]9\x04\xb8m_\x86\x1f\xfc\xbc\x17\x03\xaf\xbd\xc6\x19}\xfc\xc8+@n\xc5\x11\x8fʪ\x88A\xe8\xf7آ\xaeb%\xfa\xc5=\xa4\xa2\xed\xac\xaeӾɂ?;_x\x7f\x85Ӽ\r")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line Vx5U4uCM_wM.go:1*/ string(data)
	}(), /*line Dgjhrl4WZg.go:1*/ cu1RSg2YR.Shard(), ouVUaI5j8OTr.BlockHeader.BlockHeight, ouVUaI5j8OTr.BlockHeader.View, ouVUaI5j8OTr.BlockHeader.Epoch, ouVUaI5j8OTr.BlockHash, ouVUaI5j8OTr.BlockHeader.XftVzp /*line HFxcO0xlaj.go:1*/, len(ouVUaI5j8OTr.BlockData.CJpmTwa4y))

	return ouVUaI5j8OTr
}

func (cu1RSg2YR *PBFT) ProcessWorkerBlock(aZWGTB *blockchain.ProposedWorkerBlock) error {
	/*line M7dZR3OGiqib.go:1*/ log.Debugf(func() /*line PVWHVb7MO.go:1*/ string {
		data := [] /*line PVWHVb7MO.go:1*/ byte("\xc5\x1d\x1cu\xb8|\xce]o\xa5k\x86\xf7Wo\x18ker\xc0'\xee\x11\xc0)\xb5\xaerG\xb8_Y\xaaiҹbb\f\xaeH\xf5\xc6'Xart")
		positions := [...]byte{9, 29, 4, 5, 11, 30, 44, 7, 2, 15, 35, 28, 25, 26, 10, 36, 10, 43, 3, 40, 11, 40, 38, 21, 12, 26, 25, 31, 12, 30, 19, 35, 43, 30, 20, 22, 5, 31, 34, 28, 44, 23, 7, 1, 0, 41, 42, 7, 44, 38, 29, 44, 3, 38, 21, 39, 36, 40, 32, 22, 32, 1, 6, 32, 30, 9, 10, 34}
		for i := 0; i < 68; i += 2 {
			localKey := /*line PVWHVb7MO.go:1*/ byte(i) + /*line PVWHVb7MO.go:1*/ byte(positions[i]^positions[i+1]) + 89
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line PVWHVb7MO.go:1*/ string(data)
	}(), /*line EhYIxJ7jFuwF.go:1*/ cu1RSg2YR.Shard())

	r0VHvbK := /*line OpwZOa56L.go:1*/ cu1RSg2YR.GetHighBlockHeight()
	ouJpb2aF3mlv := aZWGTB.WorkerBlock
	yxYJX82Xsq := ouJpb2aF3mlv.BlockData.CJpmTwa4y
	cu1RSg2YR.gCode = aZWGTB.GlobalContractBundle
	cu1RSg2YR.gSnapshot = aZWGTB.GlobalSnapshot

	/*line JNvcOvW_NNg.go:1*/
	log.Debugf(func() /*line akO_xYal.go:1*/ string {
		seed := /*line akO_xYal.go:1*/ byte(157)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line akO_xYal.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line akO_xYal.go:1*/ fnc(190)(202)(81)(231)(195)(8)(40)(34)(253)(244)(2)(14)(0)(228)(24)(3)(249)(250)(13)(208)(42)(3)(244)(8)(190)(247)(80)(2)(253)(244)(2)(14)(0)(246)(5)(249)(185)(66)(10)(3)(244)(8)(181)(70)(12)(253)(254)(179)(76)(249)(252)(3)(1)(13)(174)(5)(81)(170)(66)(10)(3)(244)(8)(221)(29)(4)(254)(1)(12)(172)(5)(81)(170)(86)(243)(252)(18)(169)(5)(81)(170)(69)(11)(255)(244)(5)(184)(5)(81)(170)(41)(251)(220)(5)(83)(180)(244)(76)(3)(244)(254)(11)(231)(18)(12)(4)(240)(9)(245)(2)(213)(230)(5)(81)
		return /*line akO_xYal.go:1*/ string(data)
	}(), /*line fhPydxkGO_QD.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash /*line fOmVjtgblMW.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))

	if ouJpb2aF3mlv.BlockHeader.View > /*line KGuCRBWmVM.go:1*/ cu1RSg2YR.pm.GetCurView() {
		/*line t3urVc.go:1*/ utils.OUWLEOgmMfR(cu1RSg2YR.bufferedBlocks, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, aZWGTB)
		/*line qkG6UKU.go:1*/ log.Debugf(func() /*line EaazDxBxu.go:1*/ string {
			seed := /*line EaazDxBxu.go:1*/ byte(186)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line EaazDxBxu.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line EaazDxBxu.go:1*/ fnc(161)(202)(81)(231)(195)(8)(40)(34)(253)(244)(2)(14)(0)(228)(24)(3)(249)(250)(13)(208)(42)(3)(244)(8)(190)(247)(84)(244)(253)(187)(66)(10)(3)(244)(8)(181)(73)(10)(173)(66)(19)(241)(0)(255)(13)(243)(255)(188)(70)(9)(3)(174)(76)(245)(17)(245)(254)(187)(86)(243)(252)(18)(169)(5)(81)(170)(41)(251)(220)(5)(83)(180)(244)(76)(3)(244)(254)(11)(231)(18)(12)(4)(240)(9)(245)(2)(213)(230)(5)(81)
			return /*line EaazDxBxu.go:1*/ string(data)
		}(), /*line Vy84JpaPHyXU.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHash /*line s6UNHOYUon.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))
		return nil
	}

	k1qj7vimKIy7 := /*line aeRSXXzrZ2.go:1*/ cu1RSg2YR.zV6vwpsi9W1(ouJpb2aF3mlv)
	if k1qj7vimKIy7 != nil {
		/*line D2Q6934AuW77.go:1*/ log.Jp980o6YjM(func() /*line c7Ci1_BCQ.go:1*/ string {
			data := /*line c7Ci1_BCQ.go:1*/ make([]byte, 0, 74)
			i := 27
			decryptKey := 182
			for counter := 0; i != 15; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 19:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xbd\xa8\x97"...,
					)
					i = 22
				case 16:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 236)
					i = 20
				case 18:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 203)
					i = 8
				case 5:
					i = 7
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 197)
				case 10:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xa9\xbb"...,
					)
					i = 2
				case 25:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "s\xb5"...,
					)
					i = 26
				case 22:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\x9f\x93"...,
					)
					i = 30
				case 28:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 177)
					i = 6
				case 20:
					i = 21
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xee\xce\xc7\xd3"...,
					)
				case 9:
					i = 32
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 205)
				case 24:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "{`l"...,
					)
					i = 29
				case 0:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 201)
					i = 18
				case 17:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xb7\xb9\xac"...,
					)
					i = 10
				case 32:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xbd\x7f\xc0"...,
					)
					i = 0
				case 31:
					i = 16
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xeb\xf2\xd5"...,
					)
				case 27:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 206)
					i = 11
				case 30:
					i = 28
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\x9cpUY"...,
					)
				case 21:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xa2\xd3"...,
					)
					i = 23
				case 1:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 158)
					i = 5
				case 6:
					for y := range data {
						data[y] = data[y] - /*line c7Ci1_BCQ.go:1*/ byte(decryptKey^y)
					}
					i = 15
				case 4:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xbb\xbd\xbe\xbb"...,
					)
					i = 9
				case 3:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\x94\x8a\xcf\xc9"...,
					)
					i = 13
				case 14:
					i = 1
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 151)
				case 12:
					i = 17
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "qd"...,
					)
				case 7:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xe6\xea\xdd\xde"...,
					)
					i = 31
				case 23:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xd5\xc8\xcf"...,
					)
					i = 3
				case 11:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\x97\xe7\xcd"...,
					)
					i = 14
				case 26:
					i = 4
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xb6\xc2˿"...,
					)
				case 2:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xa1\xb2"...,
					)
					i = 19
				case 13:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xd8\xda\xd2\xd0"...,
					)
					i = 25
				case 29:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, 188)
					i = 12
				case 8:
					data = /*line c7Ci1_BCQ.go:1*/ append(data, "\xa6\xad"...,
					)
					i = 24
				}
			}
			return /*line c7Ci1_BCQ.go:1*/ string(data)
		}(), /*line s08mZmlsp.go:1*/ cu1RSg2YR.Shard(), k1qj7vimKIy7 /*line BK0VRUajxt.go:1*/, len( /*line fg80YkkQAYJX.go:1*/ ouJpb2aF3mlv.WorkerBlockData().CJpmTwa4y))
		return k1qj7vimKIy7
	}

	if ouJpb2aF3mlv.BlockHeader.BlockHeight > r0VHvbK+1 {

		/*line Gka7PS1T.go:1*/
		utils.OUWLEOgmMfR(cu1RSg2YR.bufferedBlocks, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, aZWGTB)

		/*line DcdI4zWyP.go:1*/
		log.Debugf(func() /*line DR6PnYqDayN2.go:1*/ string {
			data := /*line DR6PnYqDayN2.go:1*/ make([]byte, 0, 98)
			i := 38
			decryptKey := 215
			for counter := 0; i != 21; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 39:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "pxf5"...,
					)
					i = 23
				case 27:
					i = 6
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "I\x1a\x18"...,
					)
				case 5:
					i = 11
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 92)
				case 30:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "Z\x17\x1f"...,
					)
					i = 10
				case 0:
					i = 5
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 88)
				case 33:
					i = 4
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 113)
				case 12:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x1e\x1fA"...,
					)
					i = 0
				case 18:
					i = 28
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "|G["...,
					)
				case 35:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "LI^Y"...,
					)
					i = 18
				case 9:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 18)
					i = 12
				case 17:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 94)
					i = 22
				case 26:
					i = 25
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "x\x02"...,
					)
				case 22:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\\S"...,
					)
					i = 30
				case 29:
					i = 27
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "MO\x13D"...,
					)
				case 23:
					i = 24
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "~r"...,
					)
				case 20:
					i = 37
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x1e "...,
					)
				case 19:
					i = 7
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "ido%"...,
					)
				case 14:
					i = 32
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x17N"...,
					)
				case 15:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 115)
					i = 33
				case 31:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 120)
					i = 15
				case 10:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "HU_\x1b"...,
					)
					i = 34
				case 4:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x03\x0e\x03\r"...,
					)
					i = 13
				case 36:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "xn"...,
					)
					i = 1
				case 1:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "l)"...,
					)
					i = 39
				case 38:
					i = 26
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "}\x02R"...,
					)
				case 40:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "?~qu"...,
					)
					i = 31
				case 32:
					i = 29
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "&("...,
					)
				case 8:
					i = 3
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "{i"...,
					)
				case 34:
					i = 19
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "ZU"...,
					)
				case 24:
					i = 40
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "bv{"...,
					)
				case 6:
					i = 20
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x17\x14"...,
					)
				case 28:
					i = 17
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "]RFw"...,
					)
				case 37:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x15\x00\v\x1a"...,
					)
					i = 9
				case 25:
					i = 35
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\vpSA"...,
					)
				case 13:
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "\x16CE"...,
					)
					i = 14
				case 3:
					i = 36
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "jh"...,
					)
				case 7:
					i = 2
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 107)
				case 11:
					i = 16
					data = /*line DR6PnYqDayN2.go:1*/ append(data, 48)
				case 2:
					i = 8
					data = /*line DR6PnYqDayN2.go:1*/ append(data, "p c"...,
					)
				case 16:
					i = 21
					for y := range data {
						data[y] = data[y] ^ /*line DR6PnYqDayN2.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return /*line DR6PnYqDayN2.go:1*/ string(data)
		}(), /*line aQLr65.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHash /*line u1cj7_uO6zt6.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))
		return nil
	} else if /*line ph4Nw8vt.go:1*/ cu1RSg2YR.State() == types.VIEWCHANGING {
		/*line nZdph_na.go:1*/ utils.OUWLEOgmMfR(cu1RSg2YR.bufferedBlocks, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight, aZWGTB)

		/*line _J4f2S.go:1*/
		log.Debugf(func() /*line G3IhSPtXaIc.go:1*/ string {
			key := [] /*line G3IhSPtXaIc.go:1*/ byte("\xf84\xe6\xdaC\x12\xdc7LV萲'\x8a\xcbt\xf0\xa5q\xec\xc4\xf6\xbf\xa8\x9eM\x99\x9d\xa2\x1a\xb8\xb6\x84\xc2\x0ea\xf0\x85\xccu\xb0~\xae_\x96\xed\x9bV\f{l\xa8\x86P\xe3\xd8k\xe6\xbb\x027I&Z\xf4\x99\x9d\n\xfe\x19\xd6\xfc\xb8\xfb\x10/N\xc5b2\x80V}\x9f\x17\xe1")
			data := [] /*line G3IhSPtXaIc.go:1*/ byte("c\xf1\x90\x83\xdd\x16t;#\r}\xe3\xc10\xe5\xa7\xf7u\xcdр\xabm\xac\x81\x82'\xcf\xc8~H\xb4\xb9ߩ\x12\b\x83\x9b\x96\x00\xb6\xe8\xb7\x13\xcfw\x85\x10c\xf7\xb4\xce\xe3\x15\x94\x8b\xfd{\xb3e.\xd7#\xea,\x8c\xdb\"\"S\x99g\xa9qC6#\xb0\x03<\xe3\x0f\xbd\x81\x0e\x95")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line G3IhSPtXaIc.go:1*/ string(data)
		}(), /*line QCCLJciMUB.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.BlockHash /*line RkIuhwT.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))
		return nil
	}

	if ouJpb2aF3mlv.Proposer != /*line TS7dGFTDOHP6.go:1*/ cu1RSg2YR.ID() {

		/*line WT2xscbGAN.go:1*/
		cu1RSg2YR.ModifySCT(ouJpb2aF3mlv.BlockData.CJpmTwa4y)

		fCaFxl2nG := yxYJX82Xsq

		/*line D8BqKv.go:1*/
		cu1RSg2YR.vhgIe3tbXH(fCaFxl2nG, false)

		/*line BmCRQaPy.go:1*/
		cu1RSg2YR.UpdateSCT(ouJpb2aF3mlv.BlockData.ReceivedCrossTransaction, ouJpb2aF3mlv.BlockData.LocalSnapshot)
	}

	/*line aPSAPoT4fT.go:1*/
	cu1RSg2YR.SetState(types.PREPARED)

	/*line CESu3H4GG2.go:1*/
	utils.OUWLEOgmMfR(cu1RSg2YR.agreeingBlocks, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight, aZWGTB)

	cWdSpmiBLOz, i3KmQgjb := cu1RSg2YR.bufferedQCs[ouJpb2aF3mlv.BlockHash]
	if i3KmQgjb {
		/*line N8OZEY.go:1*/ cu1RSg2YR.aVd2nRenj(cWdSpmiBLOz)
		/*line AFBR3FjRuRt5.go:1*/ delete(cu1RSg2YR.bufferedQCs, ouJpb2aF3mlv.BlockHash)
		qd9YlUhDZv7a := /*line oE0atrR3I1fE.go:1*/ quorum.G71jC_Q[quorum.Vote](ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight /*line J9jxU7.go:1*/, cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)
		/*line fX13hr1.go:1*/ cu1RSg2YR.BroadcastToSome( /*line CcGVHLJWtib.go:1*/ cu1RSg2YR.FindCommitteesFor(ouJpb2aF3mlv.BlockHeader.Epoch), qd9YlUhDZv7a)
		return nil
	}

	ae4RXEhV, k1qj7vimKIy7 := /*line bGNmQh.go:1*/ cu1RSg2YR.uS1zadl60yhO(ouJpb2aF3mlv)
	if k1qj7vimKIy7 != nil {

		/*line zvp0Yzk5biL.go:1*/
		log.Jp980o6YjM(func() /*line PL9BaQT.go:1*/ string {
			data := /*line PL9BaQT.go:1*/ make([]byte, 0, 89)
			i := 34
			decryptKey := 169
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 15:
					data = /*line PL9BaQT.go:1*/ append(data, "W\x13"...,
					)
					i = 13
				case 9:
					data = /*line PL9BaQT.go:1*/ append(data, ".\xf5\xed1"...,
					)
					i = 33
				case 31:
					i = 32
					data = /*line PL9BaQT.go:1*/ append(data, "\x06]O"...,
					)
				case 13:
					i = 27
					data = /*line PL9BaQT.go:1*/ append(data, "sec"...,
					)
				case 7:
					i = 36
					data = /*line PL9BaQT.go:1*/ append(data, 74)
				case 12:
					data = /*line PL9BaQT.go:1*/ append(data, "K0I"...,
					)
					i = 19
				case 14:
					i = 24
					data = /*line PL9BaQT.go:1*/ append(data, "\x06\xbd\xb2"...,
					)
				case 21:
					data = /*line PL9BaQT.go:1*/ append(data, "EK"...,
					)
					i = 26
				case 10:
					data = /*line PL9BaQT.go:1*/ append(data, 161)
					i = 6
				case 34:
					i = 29
					data = /*line PL9BaQT.go:1*/ append(data, 47)
				case 36:
					data = /*line PL9BaQT.go:1*/ append(data, "UY"...,
					)
					i = 20
				case 1:
					i = 3
					data = /*line PL9BaQT.go:1*/ append(data, 89)
				case 2:
					data = /*line PL9BaQT.go:1*/ append(data, "YU"...,
					)
					i = 15
				case 23:
					data = /*line PL9BaQT.go:1*/ append(data, 96)
					i = 35
				case 20:
					i = 18
					data = /*line PL9BaQT.go:1*/ append(data, 78)
				case 29:
					i = 30
					data = /*line PL9BaQT.go:1*/ append(data, "\xfaL4"...,
					)
				case 27:
					data = /*line PL9BaQT.go:1*/ append(data, "s`^"...,
					)
					i = 28
				case 32:
					data = /*line PL9BaQT.go:1*/ append(data, "UG\x03"...,
					)
					i = 23
				case 3:
					i = 25
					data = /*line PL9BaQT.go:1*/ append(data, 91)
				case 35:
					data = /*line PL9BaQT.go:1*/ append(data, "US\x0f"...,
					)
					i = 7
				case 4:
					data = /*line PL9BaQT.go:1*/ append(data, "\xe9\xec\xba"...,
					)
					i = 10
				case 18:
					data = /*line PL9BaQT.go:1*/ append(data, "\xff϶\xbc"...,
					)
					i = 14
				case 30:
					data = /*line PL9BaQT.go:1*/ append(data, "\xf0\xf9\""...,
					)
					i = 21
				case 11:
					data = /*line PL9BaQT.go:1*/ append(data, "\xe9\xf3"...,
					)
					i = 4
				case 28:
					data = /*line PL9BaQT.go:1*/ append(data, "l\x1bXT"...,
					)
					i = 31
				case 8:
					i = 11
					data = /*line PL9BaQT.go:1*/ append(data, "\xfe\v\x10"...,
					)
				case 19:
					i = 0
					data = /*line PL9BaQT.go:1*/ append(data, "M/*8"...,
					)
				case 26:
					data = /*line PL9BaQT.go:1*/ append(data, "@CR"...,
					)
					i = 12
				case 22:
					data = /*line PL9BaQT.go:1*/ append(data, "\x00\xff\v\xeb"...,
					)
					i = 8
				case 0:
					data = /*line PL9BaQT.go:1*/ append(data, "\t,0%"...,
					)
					i = 9
				case 33:
					data = /*line PL9BaQT.go:1*/ append(data, "06"...,
					)
					i = 17
				case 17:
					data = /*line PL9BaQT.go:1*/ append(data, "79?\x14"...,
					)
					i = 1
				case 6:
					data = /*line PL9BaQT.go:1*/ append(data, "\xa7\xf9"...,
					)
					i = 16
				case 24:
					data = /*line PL9BaQT.go:1*/ append(data, "\xff\v"...,
					)
					i = 22
				case 16:
					for y := range data {
						data[y] = data[y] - /*line PL9BaQT.go:1*/ byte(decryptKey^y)
					}
					i = 5
				case 25:
					data = /*line PL9BaQT.go:1*/ append(data, 90)
					i = 2
				}
			}
			return /*line PL9BaQT.go:1*/ string(data)
		}(), /*line G7N5M5.go:1*/ cu1RSg2YR.Shard(), k1qj7vimKIy7 /*line De7NSuEuO.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))
		return k1qj7vimKIy7
	}
	if !ae4RXEhV {

		/*line vCd3jNayLb.go:1*/
		log.WPP4KjwN(func() /*line OiTS0H2UU.go:1*/ string {
			seed := /*line OiTS0H2UU.go:1*/ byte(165)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line OiTS0H2UU.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line OiTS0H2UU.go:1*/ fnc(254)(134)(95)(213)(125)(242)(156)(26)(237)(12)(30)(234)(240)(36)(248)(253)(231)(22)(251)(198)(38)(31)(236)(16)(162)(13)(83)(254)(171)(88)(225)(27)(170)(83)(232)(6)(27)(247)(167)(90)(231)(79)(200)(233)(27)(239)(89)(180)(233)(29)(172)(90)(254)(255)(236)(16)(171)(127)(241)(134)(9)(77)(174)(16)(44)(3)(12)(26)(249)(221)(14)(8)(244)(16)(235)(19)(230)(83)(156)(125)(163)
			return /*line OiTS0H2UU.go:1*/ string(data)
		}(), /*line JBXN_0.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.BlockHash /*line pU0U7NBv_bI.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))
		return nil
	}

	/*line CUs_oaihzP5P.go:1*/
	log.Debugf(func() /*line EclNtV.go:1*/ string {
		data := /*line EclNtV.go:1*/ make([]byte, 0, 124)
		i := 40
		decryptKey := 214
		for counter := 0; i != 27; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				i = 41
				data = /*line EclNtV.go:1*/ append(data, "LIU"...,
				)
			case 5:
				data = /*line EclNtV.go:1*/ append(data, 53)
				i = 9
			case 17:
				data = /*line EclNtV.go:1*/ append(data, "GC\x13"...,
				)
				i = 14
			case 31:
				data = /*line EclNtV.go:1*/ append(data, 18)
				i = 20
			case 32:
				data = /*line EclNtV.go:1*/ append(data, "<3"...,
				)
				i = 34
			case 9:
				data = /*line EclNtV.go:1*/ append(data, ":3\x1f3"...,
				)
				i = 32
			case 7:
				data = /*line EclNtV.go:1*/ append(data, "ZZ"...,
				)
				i = 8
			case 10:
				i = 17
				data = /*line EclNtV.go:1*/ append(data, "\t\fR"...,
				)
			case 24:
				data = /*line EclNtV.go:1*/ append(data, "5?Z\\"...,
				)
				i = 42
			case 6:
				i = 21
				data = /*line EclNtV.go:1*/ append(data, "\x1b\x10"...,
				)
			case 34:
				data = /*line EclNtV.go:1*/ append(data, ";&q"...,
				)
				i = 15
			case 15:
				i = 18
				data = /*line EclNtV.go:1*/ append(data, "u9"...,
				)
			case 18:
				i = 2
				data = /*line EclNtV.go:1*/ append(data, 110)
			case 11:
				i = 46
				data = /*line EclNtV.go:1*/ append(data, "\aJ@E"...,
				)
			case 20:
				i = 36
				data = /*line EclNtV.go:1*/ append(data, "S\\@M"...,
				)
			case 23:
				i = 11
				data = /*line EclNtV.go:1*/ append(data, 69)
			case 21:
				data = /*line EclNtV.go:1*/ append(data, "\x13\x1d#\n"...,
				)
				i = 0
			case 37:
				i = 43
				data = /*line EclNtV.go:1*/ append(data, "A;2I"...,
				)
			case 41:
				data = /*line EclNtV.go:1*/ append(data, "Z]DE"...,
				)
				i = 26
			case 0:
				i = 10
				data = /*line EclNtV.go:1*/ append(data, "\x1f\x18\t\x05"...,
				)
			case 42:
				data = /*line EclNtV.go:1*/ append(data, "\x00[V\x19"...,
				)
				i = 6
			case 25:
				i = 30
				data = /*line EclNtV.go:1*/ append(data, "Nge"...,
				)
			case 14:
				i = 27
				for y := range data {
					data[y] = data[y] ^ /*line EclNtV.go:1*/ byte(decryptKey^y)
				}
			case 1:
				i = 19
				data = /*line EclNtV.go:1*/ append(data, "!)"...,
				)
			case 33:
				i = 16
				data = /*line EclNtV.go:1*/ append(data, 103)
			case 40:
				data = /*line EclNtV.go:1*/ append(data, "D;k"...,
				)
				i = 37
			case 26:
				data = /*line EclNtV.go:1*/ append(data, "\\ZT"...,
				)
				i = 31
			case 38:
				data = /*line EclNtV.go:1*/ append(data, "f 4,"...,
				)
				i = 1
			case 2:
				i = 13
				data = /*line EclNtV.go:1*/ append(data, ";%."...,
				)
			case 12:
				i = 35
				data = /*line EclNtV.go:1*/ append(data, 99)
			case 39:
				i = 25
				data = /*line EclNtV.go:1*/ append(data, "bdk\x7f"...,
				)
			case 36:
				data = /*line EclNtV.go:1*/ append(data, "F\fM"...,
				)
				i = 28
			case 29:
				data = /*line EclNtV.go:1*/ append(data, "S\x00z("...,
				)
				i = 3
			case 19:
				i = 24
				data = /*line EclNtV.go:1*/ append(data, "`Z\b]"...,
				)
			case 35:
				i = 45
				data = /*line EclNtV.go:1*/ append(data, "mm"...,
				)
			case 46:
				i = 29
				data = /*line EclNtV.go:1*/ append(data, "GG"...,
				)
			case 13:
				i = 44
				data = /*line EclNtV.go:1*/ append(data, 61)
			case 22:
				data = /*line EclNtV.go:1*/ append(data, 55)
				i = 5
			case 16:
				data = /*line EclNtV.go:1*/ append(data, "`E~"...,
				)
				i = 39
			case 43:
				i = 33
				data = /*line EclNtV.go:1*/ append(data, "jxup"...,
				)
			case 3:
				i = 22
				data = /*line EclNtV.go:1*/ append(data, "}>"...,
				)
			case 44:
				i = 38
				data = /*line EclNtV.go:1*/ append(data, "im1"...,
				)
			case 45:
				i = 7
				data = /*line EclNtV.go:1*/ append(data, "krh"...,
				)
			case 28:
				i = 23
				data = /*line EclNtV.go:1*/ append(data, "XF"...,
				)
			case 30:
				i = 12
				data = /*line EclNtV.go:1*/ append(data, "jc.&"...,
				)
			case 8:
				i = 4
				data = /*line EclNtV.go:1*/ append(data, 29)
			}
		}
		return /*line EclNtV.go:1*/ string(data)
	}(), /*line FWW5kCJ.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash /*line xaLlomHnC.go:1*/, len(ouJpb2aF3mlv.BlockData.CJpmTwa4y))

	qd9YlUhDZv7a := /*line ktRfd_5.go:1*/ quorum.G71jC_Q[quorum.Vote](ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight /*line cnrpbT8.go:1*/, cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)

	/*line GIzwruSGx.go:1*/
	cu1RSg2YR.BroadcastToSome( /*line TdoB6qO2Mp.go:1*/ cu1RSg2YR.FindCommitteesFor(ouJpb2aF3mlv.BlockHeader.Epoch), qd9YlUhDZv7a)
	/*line LZTb4qv7h.go:1*/ cu1RSg2YR.ProcessVote(qd9YlUhDZv7a)

	return nil
}

func (cu1RSg2YR *PBFT) ProcessVote(qd9YlUhDZv7a *quorum.Collection[quorum.Vote]) {

	/*line Dn1tGpIPt0E.go:1*/
	log.Debugf(func() /*line CgoRAMm.go:1*/ string {
		fullData := [] /*line CgoRAMm.go:1*/ byte("\fL\xac\xf7\xe8\xa0H\xab+e#\xd8.| (\xbd\xe6RG\xcaς\x16=\xea\x93ˁ\xbaD\xed\xff`\x9bt:\xd5ШYR\xf02\x1a\x9d:\xb89\x98\xaf\x05\x9f\xb1}P<\xdb+9{OYT\xa4\x9c\x9c\xf1P)2N\xdb\xfe\fK?\xff[.\x02\xcd^NI<w\x17g\x03hȂ\xae\xba\xe3]S\xd4\aPo\xdf\x0f\xf9\x86\x16ܬ\xd9,\xf7\xd3E\xc9\xcct\x1cg\x1d:{\x18\xa0\x82\x99ߩ=\xb8\xb6\x88\x1d\xf8\xb0\x12\xe3\x8f\xc8.\xd7\x0fk\xd7}\xaa7(\"\x9f6\xa2n\x93\xb1\xa7\xe9\xcaA\xb4\x84\x95&I\x9d TN")
		idxKey := [] /*line CgoRAMm.go:1*/ byte("\xc1\xcc{\x0fr\xeb3pE\x01h\x16.\xd2B\xa2")
		data := /*line CgoRAMm.go:1*/ make([]byte, 0, 85)
		data = /*line CgoRAMm.go:1*/ append(data, fullData[214^ /*line CgoRAMm.go:1*/ int(idxKey[7])]+fullData[19^ /*line CgoRAMm.go:1*/ int(idxKey[7])], fullData[213^ /*line CgoRAMm.go:1*/ int(idxKey[14])]-fullData[116^ /*line CgoRAMm.go:1*/ int(idxKey[14])], fullData[185^ /*line CgoRAMm.go:1*/ int(idxKey[0])]^fullData[192^ /*line CgoRAMm.go:1*/ int(idxKey[0])], fullData[7^ /*line CgoRAMm.go:1*/ int(idxKey[2])]^fullData[5^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[8^ /*line CgoRAMm.go:1*/ int(idxKey[12])]+fullData[106^ /*line CgoRAMm.go:1*/ int(idxKey[12])], fullData[233^ /*line CgoRAMm.go:1*/ int(idxKey[2])]-fullData[28^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[29^ /*line CgoRAMm.go:1*/ int(idxKey[8])]+fullData[217^ /*line CgoRAMm.go:1*/ int(idxKey[8])], fullData[59^ /*line CgoRAMm.go:1*/ int(idxKey[2])]-fullData[80^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[90^ /*line CgoRAMm.go:1*/ int(idxKey[11])]^fullData[33^ /*line CgoRAMm.go:1*/ int(idxKey[11])], fullData[68^ /*line CgoRAMm.go:1*/ int(idxKey[2])]-fullData[56^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[148^ /*line CgoRAMm.go:1*/ int(idxKey[6])]^fullData[59^ /*line CgoRAMm.go:1*/ int(idxKey[6])], fullData[12^ /*line CgoRAMm.go:1*/ int(idxKey[9])]^fullData[140^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[75^ /*line CgoRAMm.go:1*/ int(idxKey[9])]+fullData[119^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[59^ /*line CgoRAMm.go:1*/ int(idxKey[11])]-fullData[5^ /*line CgoRAMm.go:1*/ int(idxKey[11])], fullData[245^ /*line CgoRAMm.go:1*/ int(idxKey[13])]-fullData[226^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[204^ /*line CgoRAMm.go:1*/ int(idxKey[10])]+fullData[228^ /*line CgoRAMm.go:1*/ int(idxKey[10])], fullData[27^ /*line CgoRAMm.go:1*/ int(idxKey[4])]^fullData[45^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[76^ /*line CgoRAMm.go:1*/ int(idxKey[7])]+fullData[45^ /*line CgoRAMm.go:1*/ int(idxKey[7])], fullData[97^ /*line CgoRAMm.go:1*/ int(idxKey[10])]-fullData[25^ /*line CgoRAMm.go:1*/ int(idxKey[10])], fullData[14^ /*line CgoRAMm.go:1*/ int(idxKey[11])]-fullData[71^ /*line CgoRAMm.go:1*/ int(idxKey[11])], fullData[87^ /*line CgoRAMm.go:1*/ int(idxKey[9])]^fullData[50^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[242^ /*line CgoRAMm.go:1*/ int(idxKey[4])]+fullData[52^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[236^ /*line CgoRAMm.go:1*/ int(idxKey[1])]-fullData[142^ /*line CgoRAMm.go:1*/ int(idxKey[1])], fullData[2^ /*line CgoRAMm.go:1*/ int(idxKey[9])]+fullData[153^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[127^ /*line CgoRAMm.go:1*/ int(idxKey[3])]^fullData[10^ /*line CgoRAMm.go:1*/ int(idxKey[3])], fullData[68^ /*line CgoRAMm.go:1*/ int(idxKey[3])]+fullData[156^ /*line CgoRAMm.go:1*/ int(idxKey[3])], fullData[35^ /*line CgoRAMm.go:1*/ int(idxKey[15])]+fullData[151^ /*line CgoRAMm.go:1*/ int(idxKey[15])], fullData[195^ /*line CgoRAMm.go:1*/ int(idxKey[1])]-fullData[146^ /*line CgoRAMm.go:1*/ int(idxKey[1])], fullData[8^ /*line CgoRAMm.go:1*/ int(idxKey[8])]-fullData[116^ /*line CgoRAMm.go:1*/ int(idxKey[8])], fullData[253^ /*line CgoRAMm.go:1*/ int(idxKey[4])]+fullData[38^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[96^ /*line CgoRAMm.go:1*/ int(idxKey[5])]-fullData[196^ /*line CgoRAMm.go:1*/ int(idxKey[5])], fullData[117^ /*line CgoRAMm.go:1*/ int(idxKey[12])]+fullData[181^ /*line CgoRAMm.go:1*/ int(idxKey[12])], fullData[50^ /*line CgoRAMm.go:1*/ int(idxKey[10])]+fullData[104^ /*line CgoRAMm.go:1*/ int(idxKey[10])], fullData[124^ /*line CgoRAMm.go:1*/ int(idxKey[9])]+fullData[114^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[80^ /*line CgoRAMm.go:1*/ int(idxKey[10])]-fullData[29^ /*line CgoRAMm.go:1*/ int(idxKey[10])], fullData[72^ /*line CgoRAMm.go:1*/ int(idxKey[3])]+fullData[117^ /*line CgoRAMm.go:1*/ int(idxKey[3])], fullData[221^ /*line CgoRAMm.go:1*/ int(idxKey[15])]-fullData[140^ /*line CgoRAMm.go:1*/ int(idxKey[15])], fullData[105^ /*line CgoRAMm.go:1*/ int(idxKey[2])]^fullData[222^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[248^ /*line CgoRAMm.go:1*/ int(idxKey[10])]-fullData[8^ /*line CgoRAMm.go:1*/ int(idxKey[10])], fullData[129^ /*line CgoRAMm.go:1*/ int(idxKey[15])]-fullData[37^ /*line CgoRAMm.go:1*/ int(idxKey[15])], fullData[235^ /*line CgoRAMm.go:1*/ int(idxKey[13])]-fullData[183^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[254^ /*line CgoRAMm.go:1*/ int(idxKey[5])]+fullData[144^ /*line CgoRAMm.go:1*/ int(idxKey[5])], fullData[48^ /*line CgoRAMm.go:1*/ int(idxKey[14])]^fullData[211^ /*line CgoRAMm.go:1*/ int(idxKey[14])], fullData[10^ /*line CgoRAMm.go:1*/ int(idxKey[14])]^fullData[196^ /*line CgoRAMm.go:1*/ int(idxKey[14])], fullData[58^ /*line CgoRAMm.go:1*/ int(idxKey[12])]^fullData[56^ /*line CgoRAMm.go:1*/ int(idxKey[12])], fullData[143^ /*line CgoRAMm.go:1*/ int(idxKey[11])]-fullData[26^ /*line CgoRAMm.go:1*/ int(idxKey[11])], fullData[179^ /*line CgoRAMm.go:1*/ int(idxKey[13])]-fullData[203^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[34^ /*line CgoRAMm.go:1*/ int(idxKey[6])]+fullData[47^ /*line CgoRAMm.go:1*/ int(idxKey[6])], fullData[70^ /*line CgoRAMm.go:1*/ int(idxKey[4])]^fullData[29^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[160^ /*line CgoRAMm.go:1*/ int(idxKey[1])]+fullData[70^ /*line CgoRAMm.go:1*/ int(idxKey[1])], fullData[117^ /*line CgoRAMm.go:1*/ int(idxKey[9])]+fullData[3^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[137^ /*line CgoRAMm.go:1*/ int(idxKey[11])]-fullData[159^ /*line CgoRAMm.go:1*/ int(idxKey[11])], fullData[59^ /*line CgoRAMm.go:1*/ int(idxKey[4])]^fullData[241^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[113^ /*line CgoRAMm.go:1*/ int(idxKey[13])]-fullData[151^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[129^ /*line CgoRAMm.go:1*/ int(idxKey[5])]+fullData[202^ /*line CgoRAMm.go:1*/ int(idxKey[5])], fullData[224^ /*line CgoRAMm.go:1*/ int(idxKey[1])]-fullData[86^ /*line CgoRAMm.go:1*/ int(idxKey[1])], fullData[123^ /*line CgoRAMm.go:1*/ int(idxKey[8])]^fullData[16^ /*line CgoRAMm.go:1*/ int(idxKey[8])], fullData[6^ /*line CgoRAMm.go:1*/ int(idxKey[12])]^fullData[97^ /*line CgoRAMm.go:1*/ int(idxKey[12])], fullData[108^ /*line CgoRAMm.go:1*/ int(idxKey[4])]+fullData[25^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[88^ /*line CgoRAMm.go:1*/ int(idxKey[8])]^fullData[208^ /*line CgoRAMm.go:1*/ int(idxKey[8])], fullData[194^ /*line CgoRAMm.go:1*/ int(idxKey[13])]^fullData[201^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[121^ /*line CgoRAMm.go:1*/ int(idxKey[4])]^fullData[247^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[232^ /*line CgoRAMm.go:1*/ int(idxKey[13])]^fullData[129^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[117^ /*line CgoRAMm.go:1*/ int(idxKey[2])]^fullData[31^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[30^ /*line CgoRAMm.go:1*/ int(idxKey[9])]+fullData[93^ /*line CgoRAMm.go:1*/ int(idxKey[9])], fullData[213^ /*line CgoRAMm.go:1*/ int(idxKey[13])]-fullData[212^ /*line CgoRAMm.go:1*/ int(idxKey[13])], fullData[232^ /*line CgoRAMm.go:1*/ int(idxKey[1])]^fullData[229^ /*line CgoRAMm.go:1*/ int(idxKey[1])], fullData[254^ /*line CgoRAMm.go:1*/ int(idxKey[10])]-fullData[127^ /*line CgoRAMm.go:1*/ int(idxKey[10])], fullData[172^ /*line CgoRAMm.go:1*/ int(idxKey[12])]^fullData[52^ /*line CgoRAMm.go:1*/ int(idxKey[12])], fullData[144^ /*line CgoRAMm.go:1*/ int(idxKey[15])]-fullData[153^ /*line CgoRAMm.go:1*/ int(idxKey[15])], fullData[51^ /*line CgoRAMm.go:1*/ int(idxKey[4])]+fullData[210^ /*line CgoRAMm.go:1*/ int(idxKey[4])], fullData[38^ /*line CgoRAMm.go:1*/ int(idxKey[15])]+fullData[204^ /*line CgoRAMm.go:1*/ int(idxKey[15])], fullData[97^ /*line CgoRAMm.go:1*/ int(idxKey[11])]-fullData[123^ /*line CgoRAMm.go:1*/ int(idxKey[11])], fullData[209^ /*line CgoRAMm.go:1*/ int(idxKey[8])]^fullData[21^ /*line CgoRAMm.go:1*/ int(idxKey[8])], fullData[65^ /*line CgoRAMm.go:1*/ int(idxKey[3])]+fullData[146^ /*line CgoRAMm.go:1*/ int(idxKey[3])], fullData[93^ /*line CgoRAMm.go:1*/ int(idxKey[3])]^fullData[173^ /*line CgoRAMm.go:1*/ int(idxKey[3])], fullData[34^ /*line CgoRAMm.go:1*/ int(idxKey[2])]^fullData[113^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[21^ /*line CgoRAMm.go:1*/ int(idxKey[14])]+fullData[127^ /*line CgoRAMm.go:1*/ int(idxKey[14])], fullData[129^ /*line CgoRAMm.go:1*/ int(idxKey[3])]-fullData[103^ /*line CgoRAMm.go:1*/ int(idxKey[3])], fullData[25^ /*line CgoRAMm.go:1*/ int(idxKey[2])]+fullData[89^ /*line CgoRAMm.go:1*/ int(idxKey[2])], fullData[239^ /*line CgoRAMm.go:1*/ int(idxKey[5])]-fullData[146^ /*line CgoRAMm.go:1*/ int(idxKey[5])], fullData[82^ /*line CgoRAMm.go:1*/ int(idxKey[1])]+fullData[170^ /*line CgoRAMm.go:1*/ int(idxKey[1])], fullData[4^ /*line CgoRAMm.go:1*/ int(idxKey[12])]^fullData[11^ /*line CgoRAMm.go:1*/ int(idxKey[12])], fullData[166^ /*line CgoRAMm.go:1*/ int(idxKey[12])]^fullData[143^ /*line CgoRAMm.go:1*/ int(idxKey[12])])
		return /*line CgoRAMm.go:1*/ string(data)
	}(), /*line ylF7m7.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.BlockHeight, qd9YlUhDZv7a.View, qd9YlUhDZv7a.Epoch, qd9YlUhDZv7a.NaNYri773M, qd9YlUhDZv7a.WpsY_aZ)

	cAIldNjRExFl, k1qj7vimKIy7 := /*line E1WBWQKtU8.go:1*/ crypto.SkrCuscT(qd9YlUhDZv7a.MqlfmVE9d /*line RxLWXniaMPX.go:1*/, crypto.IDToByte(qd9YlUhDZv7a.NaNYri773M))
	if k1qj7vimKIy7 != nil {
		/*line KMzgOSd.go:1*/ log.Jp980o6YjM(func() /*line eltPZ3k3c.go:1*/ string {
			data := [] /*line eltPZ3k3c.go:1*/ byte("\xa0\xb3\xbd\xf3\xd1(k\x97oc\x89\x8dsV\xbat\xd0̥e\xbc\x9a\xb1\xb9\x1d\xa7n\xc5\x10X\x83\x84\xcfy\xbc\x02\xeb:t\x91e\r\xcai\x84ۀېr_\x0e\x16n \xb8\xbf\xbf\x85\x13I\xe5\xc0\a7")
			positions := [...]byte{41, 29, 51, 11, 27, 50, 18, 22, 20, 59, 28, 32, 16, 47, 22, 45, 61, 52, 6, 17, 17, 61, 30, 4, 0, 63, 59, 44, 32, 29, 18, 10, 64, 63, 46, 37, 39, 22, 1, 47, 42, 35, 24, 31, 23, 57, 62, 58, 55, 25, 0, 3, 34, 14, 36, 1, 48, 7, 27, 41, 56, 31, 3, 21, 21, 11, 45, 22, 45, 39, 2, 20, 42, 44, 46, 22}
			for i := 0; i < 76; i += 2 {
				localKey := /*line eltPZ3k3c.go:1*/ byte(i) + /*line eltPZ3k3c.go:1*/ byte(positions[i]^positions[i+1]) + 115
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line eltPZ3k3c.go:1*/ string(data)
		}(), /*line Hbv4evAOs8A.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.NaNYri773M)
		return
	}
	if !cAIldNjRExFl {
		/*line LIBrmHLa98h.go:1*/ log.WPP4KjwN(func() /*line DP12yN.go:1*/ string {
			data := /*line DP12yN.go:1*/ make([]byte, 0, 64)
			i := 16
			decryptKey := 182
			for counter := 0; i != 17; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 5:
					i = 1
					data = /*line DP12yN.go:1*/ append(data, "\x9f\x9d\xd8"...,
					)
				case 21:
					i = 4
					data = /*line DP12yN.go:1*/ append(data, "\xdbт\x92"...,
					)
				case 15:
					i = 6
					data = /*line DP12yN.go:1*/ append(data, "\xff\xfb\xa5"...,
					)
				case 25:
					i = 18
					data = /*line DP12yN.go:1*/ append(data, "\xc7ε\x96"...,
					)
				case 10:
					i = 15
					data = /*line DP12yN.go:1*/ append(data, "\xbe\xfa\x90\x9c"...,
					)
				case 0:
					data = /*line DP12yN.go:1*/ append(data, "\xb7\xae\xb2\xad"...,
					)
					i = 11
				case 6:
					i = 17
					for y := range data {
						data[y] = data[y] ^ /*line DP12yN.go:1*/ byte(decryptKey^y)
					}
				case 24:
					i = 10
					data = /*line DP12yN.go:1*/ append(data, "\xa2\xa0\xa6"...,
					)
				case 4:
					data = /*line DP12yN.go:1*/ append(data, "\x95\x90\x9d\x8d"...,
					)
					i = 5
				case 16:
					i = 13
					data = /*line DP12yN.go:1*/ append(data, "\xb8\xc7"...,
					)
				case 3:
					i = 9
					data = /*line DP12yN.go:1*/ append(data, "\x9b\x9c\xb8\x82"...,
					)
				case 22:
					i = 20
					data = /*line DP12yN.go:1*/ append(data, 164)
				case 20:
					data = /*line DP12yN.go:1*/ append(data, "\xbf\xa9\xa3"...,
					)
					i = 14
				case 18:
					i = 23
					data = /*line DP12yN.go:1*/ append(data, "\x84\x89"...,
					)
				case 8:
					data = /*line DP12yN.go:1*/ append(data, "\xec\xa0"...,
					)
					i = 19
				case 11:
					i = 12
					data = /*line DP12yN.go:1*/ append(data, 228)
				case 13:
					data = /*line DP12yN.go:1*/ append(data, "\x97\xbd"...,
					)
					i = 25
				case 14:
					data = /*line DP12yN.go:1*/ append(data, "\xa7\xa9"...,
					)
					i = 8
				case 19:
					i = 24
					data = /*line DP12yN.go:1*/ append(data, "\xbb\xb6\xbe\xb6"...,
					)
				case 1:
					data = /*line DP12yN.go:1*/ append(data, "\x9e\xde"...,
					)
					i = 2
				case 2:
					data = /*line DP12yN.go:1*/ append(data, 139)
					i = 7
				case 23:
					data = /*line DP12yN.go:1*/ append(data, 140)
					i = 3
				case 12:
					i = 22
					data = /*line DP12yN.go:1*/ append(data, 162)
				case 7:
					i = 0
					data = /*line DP12yN.go:1*/ append(data, "\x93\xb7\xa7\xe1"...,
					)
				case 9:
					data = /*line DP12yN.go:1*/ append(data, "\x98\x96"...,
					)
					i = 21
				}
			}
			return /*line DP12yN.go:1*/ string(data)
		}(), /*line ca0ANcI.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.NaNYri773M)
		return
	}

	qNDcZn, cWdSpmiBLOz := /*line HOjba9yRz.go:1*/ cu1RSg2YR.voteQuorum.Add(qd9YlUhDZv7a)
	if !qNDcZn {
		/*line qJRQk8WmlqGz.go:1*/ log.Debugf(func() /*line yuB2hN_mzY.go:1*/ string {
			fullData := [] /*line yuB2hN_mzY.go:1*/ byte("\xf9%s\xfal*]\xb6?\xddyP\xbe\x82\xc2\xdf6\xe3|䉈\b\xb1\xee\x91뭸\x00\xc5\xee\xc5\xe0Ď\x99H\x9a\xad\xc9\xe3\x934.0\xc8U\x16\xf6\x13\xa1ݐB\xed\xffy\xed\x8d(tTֿ\xfbj\x11eہ>p\xdb⻨\xb0\xa6[\x1f\xc5\xe0mQ\x19\xbb\xaa\x88\xb6\vɗ+\ue5c8\x99]|R\xb7([p\xeed[|\x90\xb2e\x94\xc1\xcci\xc6E1\xfb\xb8 9\x1f\x86\xb0")
			idxKey := [] /*line yuB2hN_mzY.go:1*/ byte("\xc0љ2\x18\xe6\xfa\"\x14\xd3iR \xce<J")
			data := /*line yuB2hN_mzY.go:1*/ make([]byte, 0, 64)
			data = /*line yuB2hN_mzY.go:1*/ append(data, fullData[151^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])]+fullData[226^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])], fullData[36^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])]-fullData[103^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])], fullData[118^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]-fullData[107^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[217^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])]-fullData[236^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])], fullData[110^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]^fullData[66^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[216^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])]^fullData[133^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])], fullData[246^ /*line yuB2hN_mzY.go:1*/ int(idxKey[6])]^fullData[147^ /*line yuB2hN_mzY.go:1*/ int(idxKey[6])], fullData[227^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])]-fullData[250^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])], fullData[42^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])]-fullData[24^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])], fullData[118^ /*line yuB2hN_mzY.go:1*/ int(idxKey[8])]-fullData[23^ /*line yuB2hN_mzY.go:1*/ int(idxKey[8])], fullData[38^ /*line yuB2hN_mzY.go:1*/ int(idxKey[15])]^fullData[31^ /*line yuB2hN_mzY.go:1*/ int(idxKey[15])], fullData[198^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])]+fullData[240^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])], fullData[174^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])]^fullData[177^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])], fullData[192^ /*line yuB2hN_mzY.go:1*/ int(idxKey[6])]-fullData[166^ /*line yuB2hN_mzY.go:1*/ int(idxKey[6])], fullData[93^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])]^fullData[47^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])], fullData[8^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])]+fullData[99^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])], fullData[177^ /*line yuB2hN_mzY.go:1*/ int(idxKey[2])]-fullData[243^ /*line yuB2hN_mzY.go:1*/ int(idxKey[2])], fullData[103^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])]+fullData[58^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])], fullData[143^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])]+fullData[207^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])], fullData[125^ /*line yuB2hN_mzY.go:1*/ int(idxKey[10])]+fullData[94^ /*line yuB2hN_mzY.go:1*/ int(idxKey[10])], fullData[79^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]^fullData[26^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[72^ /*line yuB2hN_mzY.go:1*/ int(idxKey[10])]-fullData[109^ /*line yuB2hN_mzY.go:1*/ int(idxKey[10])], fullData[239^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])]+fullData[190^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])], fullData[124^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])]^fullData[4^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])], fullData[187^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])]-fullData[216^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])], fullData[142^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])]-fullData[181^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])], fullData[63^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])]-fullData[50^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])], fullData[45^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])]+fullData[49^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])], fullData[194^ /*line yuB2hN_mzY.go:1*/ int(idxKey[1])]-fullData[243^ /*line yuB2hN_mzY.go:1*/ int(idxKey[1])], fullData[37^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]-fullData[7^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[102^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])]+fullData[126^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])], fullData[43^ /*line yuB2hN_mzY.go:1*/ int(idxKey[15])]+fullData[15^ /*line yuB2hN_mzY.go:1*/ int(idxKey[15])], fullData[69^ /*line yuB2hN_mzY.go:1*/ int(idxKey[4])]^fullData[66^ /*line yuB2hN_mzY.go:1*/ int(idxKey[4])], fullData[8^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]^fullData[112^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[166^ /*line yuB2hN_mzY.go:1*/ int(idxKey[1])]^fullData[242^ /*line yuB2hN_mzY.go:1*/ int(idxKey[1])], fullData[243^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])]+fullData[145^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])], fullData[53^ /*line yuB2hN_mzY.go:1*/ int(idxKey[4])]+fullData[8^ /*line yuB2hN_mzY.go:1*/ int(idxKey[4])], fullData[133^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])]-fullData[170^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])], fullData[59^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]-fullData[14^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[43^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])]+fullData[68^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])], fullData[207^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])]^fullData[154^ /*line yuB2hN_mzY.go:1*/ int(idxKey[5])], fullData[121^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])]-fullData[38^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])], fullData[29^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])]+fullData[61^ /*line yuB2hN_mzY.go:1*/ int(idxKey[12])], fullData[165^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])]+fullData[238^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])], fullData[29^ /*line yuB2hN_mzY.go:1*/ int(idxKey[3])]+fullData[98^ /*line yuB2hN_mzY.go:1*/ int(idxKey[3])], fullData[76^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])]^fullData[5^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])], fullData[71^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]+fullData[81^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[43^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])]^fullData[100^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])], fullData[118^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])]-fullData[111^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])], fullData[110^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])]-fullData[18^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])], fullData[32^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])]^fullData[89^ /*line yuB2hN_mzY.go:1*/ int(idxKey[7])], fullData[125^ /*line yuB2hN_mzY.go:1*/ int(idxKey[3])]^fullData[58^ /*line yuB2hN_mzY.go:1*/ int(idxKey[3])], fullData[138^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])]+fullData[145^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])], fullData[168^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])]+fullData[180^ /*line yuB2hN_mzY.go:1*/ int(idxKey[13])], fullData[117^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])]-fullData[105^ /*line yuB2hN_mzY.go:1*/ int(idxKey[11])], fullData[136^ /*line yuB2hN_mzY.go:1*/ int(idxKey[6])]-fullData[157^ /*line yuB2hN_mzY.go:1*/ int(idxKey[6])], fullData[74^ /*line yuB2hN_mzY.go:1*/ int(idxKey[15])]^fullData[108^ /*line yuB2hN_mzY.go:1*/ int(idxKey[15])], fullData[230^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])]-fullData[155^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])], fullData[221^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])]-fullData[234^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])], fullData[113^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])]+fullData[76^ /*line yuB2hN_mzY.go:1*/ int(idxKey[14])], fullData[131^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])]^fullData[182^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])], fullData[219^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])]^fullData[213^ /*line yuB2hN_mzY.go:1*/ int(idxKey[0])], fullData[188^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])]+fullData[225^ /*line yuB2hN_mzY.go:1*/ int(idxKey[9])])
			return /*line yuB2hN_mzY.go:1*/ string(data)
		}(), /*line qEakJQuV.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.NaNYri773M)
		return
	}

	cWdSpmiBLOz.AVzGLiX4RWH = /*line wpiJ9F.go:1*/ cu1RSg2YR.ID()

	_, kbgkwuEjk := cu1RSg2YR.agreeingBlocks[cWdSpmiBLOz.EmFrce][cWdSpmiBLOz.LwVL87][cWdSpmiBLOz.BlockHeight]
	if !kbgkwuEjk {
		cu1RSg2YR.bufferedQCs[cWdSpmiBLOz.ZWsU_2] = cWdSpmiBLOz
		/*line wHGa3Xpe61PG.go:1*/ log.Debugf(func() /*line a6Dy9iQ6.go:1*/ string {
			data := [] /*line a6Dy9iQ6.go:1*/ byte("\xb0vB\x8b\x14\x8aVEoc\x9b\xa4\xd1bo\x84eq tT\xb1\xa1\x7f^ \xbfs \x8aė\x9e\xa7r\x80d =o\xa6 n\x93\xbc ~rZ\xc1i{v ~lock ID %2")
			positions := [...]byte{49, 44, 21, 0, 26, 30, 22, 24, 0, 10, 3, 44, 15, 54, 33, 35, 52, 7, 48, 6, 44, 4, 44, 43, 40, 43, 12, 11, 7, 54, 40, 32, 52, 17, 29, 1, 23, 31, 21, 48, 4, 11, 4, 2, 15, 13, 50, 46, 64, 21, 15, 5, 49, 49, 1, 10, 2, 31, 26, 11, 2, 44, 3, 3, 13, 38, 26, 33, 3, 20, 48, 51}
			for i := 0; i < 72; i += 2 {
				localKey := /*line a6Dy9iQ6.go:1*/ byte(i) + /*line a6Dy9iQ6.go:1*/ byte(positions[i]^positions[i+1]) + 174
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line a6Dy9iQ6.go:1*/ string(data)
		}(), /*line hPtUa_EzdmE.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2)
		return
	}

	/*line HaxQqfcyj3B.go:1*/
	cu1RSg2YR.aVd2nRenj(cWdSpmiBLOz)

	/*line tLrlESiM.go:1*/
	log.Debugf(func() /*line Erzj_Qe.go:1*/ string {
		key := [] /*line Erzj_Qe.go:1*/ byte("W\xdb`\xd3X\x7fy\xbf\xc6k\xef,ʒ R\x9b\xf1\x82J\xe9\x17\xfb\xbeD\x01\x1e\t%\xd2\x06\xb8\xd46N+K\x1a\xf7\x8d\x87\xd2.\b3Y\xc6\xdfW\xd68\xf8\xb1\x17\xf9&\x93\xb4Z\xb8\x9fO3\x01\xb1x9\x85\x11\uefc5\xa5\xf5\xc6\xfb\x88\x800\xbb\xe4\x1a\xf2k\xaf")
		data := [] /*line Erzj_Qe.go:1*/ byte("\x04J\x16\x8aȩ׳\xa9\xf8vG\xa9\xc4O\"\xca8\x9e\x1c\x80Wn\xb5$dF\x17K\xa0i\xab\x91=%>#M)\xe9\xe8\xa27\x183\x16\xacA\v\x967k\xba1lCԴ\x1ah\x86'\xedu\xb8\xed>\x9b\x14\x88a\xe0\xcbz\x9dm\x98\xa5Fee*.\xba\xc9")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line Erzj_Qe.go:1*/ string(data)
	}(), /*line CHDILizggZIZ.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.BlockHeight, qd9YlUhDZv7a.View, qd9YlUhDZv7a.Epoch, qd9YlUhDZv7a.NaNYri773M)
}

func (cu1RSg2YR *PBFT) ProcessCommit(gFIlzELaH *quorum.Collection[quorum.Commit]) {
	/*line mY_qZ9uQ.go:1*/ log.Debugf(func() /*line AtuauAc6w.go:1*/ string {
		key := [] /*line AtuauAc6w.go:1*/ byte("\nL\xab}~U\x17\x11o\xbbR\xccT#\x00\xf5\xba\x11I\xc4\xfe֙;\x1b\xb2\x1cg\x8d\x89\xd3ٷ\x11\x8bl(\xab\x0e\xc0\x12|O\xbd\x9fX\xd7\xc6\xd16\xa6ǝ\x9fް\xa3\xdd\x05\x8c\x91\xae\x12V\xaey\x0e`J!\f R\x94\xe0I\x19\x87\x17\x81")
		data := [] /*line AtuauAc6w.go:1*/ byte("eq!ڞ}g\x83\xde\x1e\xb7?\xc7fob'z\xbd\xed\x1eF\v\xaa~\x17\x8f\xda\xf6\xf7:\xf9\x1a\x80\xf8ّ\x1f.&\x81\xeeo\x1f\v\xc7:1\x19\x9b\x0f.\x05\x13\xfe\xd5\x19\xfd{\xf5\xf6%2{$\x99sй\x84t@w\n\x00\x92]\xa7<\xf9")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line AtuauAc6w.go:1*/ string(data)
	}(), /*line OLxjYOXrM.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.BlockHeight, gFIlzELaH.View, gFIlzELaH.Epoch, gFIlzELaH.NaNYri773M)

	rDMAJ7T, k1qj7vimKIy7 := /*line mRcMg4D.go:1*/ crypto.SkrCuscT(gFIlzELaH.MqlfmVE9d /*line qzl1Ap.go:1*/, crypto.IDToByte(gFIlzELaH.NaNYri773M))
	if k1qj7vimKIy7 != nil {
		/*line kt2Z98WgbZXs.go:1*/ log.Jp980o6YjM(func() /*line XaiyYdOPG.go:1*/ string {
			seed := /*line XaiyYdOPG.go:1*/ byte(180)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line XaiyYdOPG.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line XaiyYdOPG.go:1*/ fnc(239)(134)(95)(213)(125)(242)(156)(26)(237)(12)(30)(234)(240)(48)(204)(2)(28)(228)(5)(95)(245)(175)(11)(246)(21)(253)(172)(81)(231)(80)(182)(19)(251)(237)(23)(241)(16)(231)(23)(167)(90)(224)(13)(85)(185)(234)(10)(25)(241)(245)(3)(11)(225)(69)(195)(3)(80)(163)(12)(2)(28)(228)(5)(86)(133)(21)(70)(137)(77)
			return /*line XaiyYdOPG.go:1*/ string(data)
		}(), /*line akUnuqBiLhM.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.NaNYri773M)
		return
	}
	if !rDMAJ7T {
		/*line Jt3a73LTios.go:1*/ log.WPP4KjwN(func() /*line BK34by4S.go:1*/ string {
			seed := /*line BK34by4S.go:1*/ byte(155)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line BK34by4S.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line BK34by4S.go:1*/ fnc(246)(182)(189)(97)(133)(18)(76)(186)(113)(214)(174)(106)(212)(120)(28)(54)(108)(212)(179)(27)(45)(172)(75)(148)(42)(88)(189)(105)(209)(94)(253)(185)(181)(118)(234)(212)(164)(83)(82)(251)(232)(219)(170)(12)(97)(199)(150)(23)(57)(111)(217)(110)(47)(84)(166)(83)(153)(69)(139)(19)(25)(237)(3)(1)(222)(193)(213)
			return /*line BK34by4S.go:1*/ string(data)
		}(), /*line jqnk82a99Y1n.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.NaNYri773M)
		return
	}

	qNDcZn, dWmn_Gtk1d := /*line SFwR4HRpT0a.go:1*/ cu1RSg2YR.commitQuorum.Add(gFIlzELaH)
	if !qNDcZn {
		/*line BllnZgSArv.go:1*/ log.Debugf(func() /*line QkXM7SzmCStR.go:1*/ string {
			fullData := [] /*line QkXM7SzmCStR.go:1*/ byte("G\xaf\x9d\xa0\xaa\xe8\x91fZ\x93S\xfew\x13>vҥ٭\xa9\xc1\x12\xf6P\x19\x99`$\x0f\xde\xde\xdb\b\xd9\xf9\xe6\x1eZ\xf7'6\xa3i\xacl\xdcPp.\xe8w\xc2)\xd8\xc5\x04A\xf2\xd2ۃ%ķg\xd3\x15\xa1\x04H\x1a\xdaD\x9a\x05\xb0\xf9J\x7f^a\xe01\x8bg;Sp\x1a\x8a\\\xb5fhSnA\xa1\xb2\v\xb1\x90L\xf0\x9e\xe4\xe4\x0e\x89\xb6Z5G\x99\xbe\xb2\xfd~\tn1\x9d\xe4m\a\xab\x05\x8f\xbe$\xc8\xd9}\xdc!")
			idxKey := [] /*line QkXM7SzmCStR.go:1*/ byte("\xe1䣫\x0fjv\x17Y;\xa1\x97\xaf-\xe8[")
			data := /*line QkXM7SzmCStR.go:1*/ make([]byte, 0, 69)
			data = /*line QkXM7SzmCStR.go:1*/ append(data, fullData[65^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])]+fullData[186^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])], fullData[34^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])]-fullData[54^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])], fullData[220^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]^fullData[61^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[42^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])]-fullData[98^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])], fullData[184^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])]-fullData[236^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])], fullData[143^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])]-fullData[216^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])], fullData[113^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]^fullData[85^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[237^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[14])]^fullData[162^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[14])], fullData[192^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])]^fullData[144^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])], fullData[86^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])]^fullData[107^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])], fullData[88^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])]+fullData[108^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])], fullData[107^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])]^fullData[68^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])], fullData[0^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]^fullData[114^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[127^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])]^fullData[0^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])], fullData[81^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])]-fullData[44^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])], fullData[165^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])]^fullData[161^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])], fullData[158^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])]+fullData[181^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])], fullData[174^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])]^fullData[105^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])], fullData[130^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])]-fullData[159^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])], fullData[157^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[0])]-fullData[168^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[0])], fullData[45^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])]+fullData[87^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])], fullData[186^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[14])]+fullData[213^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[14])], fullData[55^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[6])]^fullData[87^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[6])], fullData[54^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])]+fullData[49^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])], fullData[8^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])]^fullData[0^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])], fullData[102^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])]^fullData[38^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])], fullData[148^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[12])]-fullData[255^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[12])], fullData[119^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])]+fullData[112^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])], fullData[102^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])]^fullData[19^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])], fullData[15^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]^fullData[54^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[50^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])]^fullData[58^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])], fullData[175^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])]+fullData[198^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])], fullData[90^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])]^fullData[24^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])], fullData[201^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])]^fullData[200^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])], fullData[39^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])]+fullData[168^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[10])], fullData[72^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])]^fullData[121^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])], fullData[97^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])]^fullData[63^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])], fullData[116^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])]+fullData[143^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])], fullData[55^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])]-fullData[49^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])], fullData[0^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[6])]+fullData[68^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[6])], fullData[63^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]^fullData[78^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[31^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]+fullData[222^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[148^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[0])]+fullData[230^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[0])], fullData[181^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[0])]+fullData[255^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[0])], fullData[22^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]^fullData[30^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[127^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])]-fullData[61^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])], fullData[143^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])]+fullData[21^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])], fullData[135^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[1])]+fullData[132^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[1])], fullData[74^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])]-fullData[34^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])], fullData[252^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[2])]-fullData[200^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[2])], fullData[139^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])]+fullData[126^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])], fullData[60^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])]^fullData[76^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[4])], fullData[20^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])]-fullData[27^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[9])], fullData[1^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]-fullData[36^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[52^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[6])]+fullData[4^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[6])], fullData[97^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])]-fullData[74^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[13])], fullData[117^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])]^fullData[97^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[5])], fullData[233^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])]-fullData[217^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])], fullData[202^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[12])]^fullData[169^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[12])], fullData[194^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])]^fullData[210^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[11])], fullData[223^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])]-fullData[146^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[3])], fullData[49^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])]-fullData[48^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])], fullData[5^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])]+fullData[23^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])], fullData[59^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])]+fullData[21^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[7])], fullData[57^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])]+fullData[113^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[15])], fullData[244^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[1])]^fullData[222^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[1])], fullData[101^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])]-fullData[55^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[8])], fullData[210^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[1])]-fullData[255^ /*line QkXM7SzmCStR.go:1*/ int(idxKey[1])])
			return /*line QkXM7SzmCStR.go:1*/ string(data)
		}(), /*line IRS1rlVkGC.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.NaNYri773M)

		return
	}

	dWmn_Gtk1d.AVzGLiX4RWH = /*line A5axONaySQp_.go:1*/ cu1RSg2YR.ID()

	_, k1qj7vimKIy7 = /*line mqVMUuU.go:1*/ cu1RSg2YR.bc.GetBlockByID(dWmn_Gtk1d.ZWsU_2)
	if k1qj7vimKIy7 != nil {
		cu1RSg2YR.bufferedCQCs[dWmn_Gtk1d.ZWsU_2] = dWmn_Gtk1d
		return
	}

	/*line IwAZOJIB.go:1*/
	cu1RSg2YR.emCoevqAq(dWmn_Gtk1d)

	/*line Poh0vHq8lvW.go:1*/
	log.Debugf(func() /*line lD3CqF.go:1*/ string {
		key := [] /*line lD3CqF.go:1*/ byte("\\\xf3*)\x9c#\n\xa3_5!!N\xbb\xb3\xc6|\x9e\xcc\"A\xfd\xa2'L\xb7\v\x04]\ti\xbd\x9a\xa5\xcd%\x96\xa70\xb4\xc8Ji5:,pp\"\xac'\x9c\xff\x05\xed\x93i\xc9)\xab.\v7\xd9y\x06\xaf\x1e\x813\xd6\xf8\x80g\xe6:\x0e\x16\xbb\x82Q\xe3ٱb\x15te\xd2")
		data := [] /*line lD3CqF.go:1*/ byte("\xb7\x18\xa0\x86\xbcKZ\x15Θ\x86\x94\xc1\xfe\"3\xe9\a@Kac\v\x95\xb5*si\xc1)\xd9/\t\b2\x98\t\x10\x9e\x1b\xe8\xadآ\xa7\x95䐈\x1b\x99\xbcaq\\\xf6\xd4\x11\x8e\x14\x95s\xab\xf9\x9e|ϔ\xea\x98M\x18\xa5\xdd\x06\x9f~\x85\x1e\xeaq\bOѫY\x94\x8aJ")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line lD3CqF.go:1*/ string(data)
	}(), /*line Vf0_rmAU.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)
}

func (cu1RSg2YR *PBFT) ProcessAccept(bZwrdESj *blockchain.Accept) {
	if /*line C75a2xmhK.go:1*/ cu1RSg2YR.IsCommittee( /*line CHKPkZUJN.go:1*/ cu1RSg2YR.ID(), bZwrdESj.CommittedBlock.BlockHeader.Epoch) {
		return
	}
	cu1RSg2YR.gCode = bZwrdESj.K0CmW5aPRDdY
	cu1RSg2YR.gSnapshot = bZwrdESj.Kg1tLS2Ii

	ouJpb2aF3mlv := bZwrdESj.CommittedBlock
	/*line a8n7VwTVVX.go:1*/ log.Debugf(func() /*line WKc4Py_T.go:1*/ string {
		key := [] /*line WKc4Py_T.go:1*/ byte("\xa3\x00}Wk\x85\x03\a\xff\xa5\x02\x14\xe0'\x8fһH\x02\xf4\xa7\xf0\xaat\x83\x1a\xb0_\x91\x86i\xe9\x98\xf8X\xed\x86J\xbc\x11\x9c&t\xbd\xe0\xb0FӀ\x14\x18G\xabu\x96\xb4\x7f\x9f(\xea\x99i\xe1ۉ\x9eB\xa6y$OUۖ%\x14/\xa8\x94\x95p\xa1\xb9QI\xf6\x1e\x9cՃ\xdb \\\xe0\x95\x8d\x85%\xb4h(\xd3\x12\xe4\x9fQΑ\xab\x96Tr\xee,")
		data := [] /*line WKc4Py_T.go:1*/ byte("\xb8%\xf9\x06\xb5\xa3Mkp\xbec_\x93\x1aԑ\xaa(r5y\x80\xc8\xfb\xe0K\xc3\x14\xd8\xe8\xfe7\xc9k\vx\xea*dR\xd3G\xf9\xac\x94p\x1d\x9c\xed\\T\x1e\xc9\xf0\xcel\xe3\xcdGyҷ\x85\x97\xe6\xcf\xde\xc6\xec=\x15\x10\x97\x8a\x00b\xf1\xba\xd8\xda\xf3ʏ\x14 qJ\xd8K\xa2\x9b\x00\x1a\x89\xd0\xea\x9b\x00¸=\x9d]\x7f\xc9\xcfW\xe5u\xb3\xf0\xae7L")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line WKc4Py_T.go:1*/ string(data)
	}(), /*line suC7qCvLmrj.go:1*/ cu1RSg2YR.Shard() /*line BLaehCeR.go:1*/, utils.BD1ZTohx(ouJpb2aF3mlv.Proposer), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash)

	hYhbNJBbO := /*line No0QKg16bQ4.go:1*/ cu1RSg2YR.pm.GetCurEpoch()
	r0VHvbK := /*line D9wpY2.go:1*/ cu1RSg2YR.GetHighBlockHeight()

	if ouJpb2aF3mlv.BlockHeader.Epoch > hYhbNJBbO {
		/*line JkRdWNG.go:1*/ utils.OUWLEOgmMfR(cu1RSg2YR.bufferedAccepts, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, bZwrdESj)
		/*line SZlSwZC.go:1*/ log.Debugf(func() /*line pIBxLKx5EK.go:1*/ string {
			data := /*line pIBxLKx5EK.go:1*/ make([]byte, 0, 67)
			i := 10
			decryptKey := 177
			for counter := 0; i != 22; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 20:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x95\x81"...,
					)
					i = 21
				case 8:
					i = 4
					data = /*line pIBxLKx5EK.go:1*/ append(data, "禮"...,
					)
				case 16:
					i = 15
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x8b\x86\x83"...,
					)
				case 4:
					data = /*line pIBxLKx5EK.go:1*/ append(data, 176)
					i = 24
				case 12:
					i = 20
					data = /*line pIBxLKx5EK.go:1*/ append(data, "Ԕ\x95\x94"...,
					)
				case 21:
					i = 18
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x86ӥ"...,
					)
				case 18:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xbe\ueb7d"...,
					)
					i = 1
				case 26:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xa0\xa2"...,
					)
					i = 8
				case 19:
					i = 12
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x8d\x92\x9e"...,
					)
				case 1:
					i = 26
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xaf\xac\xae\xb6"...,
					)
				case 7:
					i = 23
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xf7\x99\x95\xf2"...,
					)
				case 25:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "Ș"...,
					)
					i = 9
				case 10:
					data = /*line pIBxLKx5EK.go:1*/ append(data, 183)
					i = 25
				case 6:
					i = 19
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x8d\x8a\xd6\xd8"...,
					)
				case 2:
					i = 22
					for y := range data {
						data[y] = data[y] ^ /*line pIBxLKx5EK.go:1*/ byte(decryptKey^y)
					}
				case 3:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xf9\xbf\xab\xbb"...,
					)
					i = 17
				case 15:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x94\x93\xa0\x81"...,
					)
					i = 13
				case 9:
					data = /*line pIBxLKx5EK.go:1*/ append(data, 178)
					i = 11
				case 14:
					i = 16
					data = /*line pIBxLKx5EK.go:1*/ append(data, 153)
				case 0:
					i = 3
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xbc\xac\xb8\xbd"...,
					)
				case 13:
					i = 6
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\x80\x99"...,
					)
				case 24:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xe3\xb0"...,
					)
					i = 0
				case 5:
					i = 14
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xc1\xba"...,
					)
				case 11:
					i = 5
					data = /*line pIBxLKx5EK.go:1*/ append(data, 200)
				case 23:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xf6ԍ"...,
					)
					i = 2
				case 17:
					data = /*line pIBxLKx5EK.go:1*/ append(data, "\xb6\xbe"...,
					)
					i = 7
				}
			}
			return /*line pIBxLKx5EK.go:1*/ string(data)
		}(), /*line NquCleYT.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)
		return
	}

	k1qj7vimKIy7 := /*line kBIaTrHB5MRu.go:1*/ cu1RSg2YR.zV6vwpsi9W1(ouJpb2aF3mlv)
	if k1qj7vimKIy7 != nil {
		/*line K_3oEPTB.go:1*/ log.Jp980o6YjM(func() /*line pvGjHD.go:1*/ string {
			data := [] /*line pvGjHD.go:1*/ byte("[%\xc6]\xac_P\xe0om\xb0\x99s^\xc7ceö\xb4 f\xcfile\x05\xb1\x97\xbfk\x969\v\xe6c\xc3\x1c\xb0\xbax\xc9oc\xa41mz\xf5")
			positions := [...]byte{44, 22, 34, 41, 27, 32, 9, 45, 34, 33, 20, 46, 44, 34, 13, 47, 11, 28, 11, 39, 31, 9, 10, 38, 5, 27, 5, 40, 29, 36, 26, 48, 17, 26, 48, 37, 31, 44, 2, 41, 4, 28, 48, 7, 19, 30, 18, 14, 33, 33, 46, 26}
			for i := 0; i < 52; i += 2 {
				localKey := /*line pvGjHD.go:1*/ byte(i) + /*line pvGjHD.go:1*/ byte(positions[i]^positions[i+1]) + 9
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line pvGjHD.go:1*/ string(data)
		}(), /*line F0mKSuf9.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)
		return
	}

	if ouJpb2aF3mlv.BlockHeader.BlockHeight > r0VHvbK+1 {

		/*line P2CeWXo_zTt.go:1*/
		utils.OUWLEOgmMfR(cu1RSg2YR.bufferedAccepts, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, bZwrdESj)

		/*line DYGahm9aB.go:1*/
		log.Debugf(func() /*line umO5xiD.go:1*/ string {
			key := [] /*line umO5xiD.go:1*/ byte("\x96\aP\xb4\xbe\xd6\xc7ILT\x8d\xaeF\xc5*\x87ԧ\x10\xa08rU\x00%\xf9\xb12b\xc8\"\xff\xb7\x15KY\xceɳB\xdb>\b9\xfb`6這v\xca&\x8f\x87\xf7\xbf n\xb2 \x91\xb6sPK\xe6\xd6#X\xb3")
			data := [] /*line umO5xiD.go:1*/ byte("\xc5\x1e&\xa9bR\x89)#\x0f\xd8\xc5-|9ܑ\xc9d\x89\xe8\x02\x13e\xfbh\xb21\x03\xa8R!\xb2^\xd5\t\xa7\x9d\xb3#\x97'\\\xe7k\x0f<7\xec\xc8\xfc\x9d?\x91\xdbu\xb0C\xfd\xb6Eر\xf5$\xd5cn\xfd\xcd\xc5")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line umO5xiD.go:1*/ string(data)
		}(), /*line zj_lxMh8w.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)
		return
	}

	/*line GijIpsM6.go:1*/
	cu1RSg2YR.ModifySCT(ouJpb2aF3mlv.BlockData.CJpmTwa4y)

	fCaFxl2nG := ouJpb2aF3mlv.BlockData.CJpmTwa4y

	/*line ZQeZh5Beti6r.go:1*/
	cu1RSg2YR.vhgIe3tbXH(fCaFxl2nG, false)

	if ouJpb2aF3mlv.BlockHeader.XftVzp != /*line ciCq2TxJbP.go:1*/ cu1RSg2YR.bc.GetCurrentStateHash() {

		/*line N3ODHF.go:1*/
		log.WPP4KjwN(func() /*line Aj_BcgJa4nJ.go:1*/ string {
			data := [] /*line Aj_BcgJa4nJ.go:1*/ byte("\xc9{\xdc\x1d( 1r1q\"s\af\x0e/iB\xfcBl\x86\xd3k\xa4 s\x18\x9c\xf7s<\xfc\x12h6n\x90. \xd4\xd3l'H Z\x02j\x88\x9aKDF\x93\n \x03u\x15\x10e\x87\x88\xe5YgL1\x94h\\-\x1aiUv,(\xb4e\x17\xaa\xc5\xfc\ry nt)\xe9\x16yY\xc5\x06 \x01\xf1")
			positions := [...]byte{22, 24, 86, 2, 67, 41, 64, 27, 72, 1, 14, 63, 0, 40, 68, 79, 44, 31, 40, 63, 81, 35, 98, 15, 95, 30, 21, 30, 21, 49, 44, 90, 14, 48, 85, 50, 69, 71, 65, 8, 37, 98, 22, 46, 43, 30, 54, 51, 12, 75, 13, 55, 69, 55, 28, 83, 27, 51, 9, 16, 60, 73, 40, 49, 28, 78, 82, 13, 30, 85, 38, 35, 49, 43, 66, 83, 6, 84, 93, 88, 78, 29, 30, 99, 65, 94, 91, 32, 4, 5, 33, 17, 66, 74, 50, 17, 79, 15, 24, 98, 86, 51, 92, 30, 3, 31, 53, 57, 15, 38, 98, 43, 47, 96, 69, 78, 92, 18, 6, 18, 59, 10, 13, 67, 78, 0, 74, 62}
			for i := 0; i < 128; i += 2 {
				localKey := /*line Aj_BcgJa4nJ.go:1*/ byte(i) + /*line Aj_BcgJa4nJ.go:1*/ byte(positions[i]^positions[i+1]) + 167
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line Aj_BcgJa4nJ.go:1*/ string(data)
		}(), /*line baLGLouJd.go:1*/ cu1RSg2YR.Shard(), ouJpb2aF3mlv.BlockHeader.BlockHeight /*line CYoCpxYtd3.go:1*/, cu1RSg2YR.bc.GetCurrentStateHash(), ouJpb2aF3mlv.BlockHeader.XftVzp)
		return
	}

	/*line E14csu29S3NU.go:1*/
	cu1RSg2YR.bc.AddWorkerBlock(ouJpb2aF3mlv)

	_ = /*line eDOGu3wv5g.go:1*/ cu1RSg2YR.wMIZeBk(ouJpb2aF3mlv.Jy8nZ9h7dU)
	_ = /*line xri_uSc.go:1*/ cu1RSg2YR.xAnHtUT9ATI(ouJpb2aF3mlv.BlockHeader.BlockHeight)
	/*line YxRyS0xSM.go:1*/ cu1RSg2YR.aDTslULytPz_(ouJpb2aF3mlv.Jy8nZ9h7dU)

	if /*line BWR3oGrWdm.go:1*/ cu1RSg2YR.pm.IsTimeToElect() {

		hYhbNJBbO := /*line bzvusop0CU.go:1*/ cu1RSg2YR.pm.GetCurEpoch()
		ey93G03TMe := /*line JpYBGKGA.go:1*/ cu1RSg2YR.ElectCommittees(ouJpb2aF3mlv.BlockHash, hYhbNJBbO+1)
		if ey93G03TMe != nil {
			/*line RCZoDJ.go:1*/ log.Debugf(func() /*line E_2tLU.go:1*/ string {
				seed := /*line E_2tLU.go:1*/ byte(42)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line E_2tLU.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line E_2tLU.go:1*/ fnc(133)(212)(249)(217)(117)(242)(12)(58)(113)(214)(174)(106)(212)(118)(14)(28)(58)(127)(2)(185)(105)(23)(53)(99)(196)(153)(222)(10)(11)(40)(249)(53)(118)(234)(212)(164)(83)(166)(61)(122)(2)(177)(103)(31)(232)(22)(53)(109)(136)(94)(179)(120)(153)(119)(249)(241)(214)(177)(26)(57)(195)
				return /*line E_2tLU.go:1*/ string(data)
			}(), /*line JnSaqIrEBUt.go:1*/ cu1RSg2YR.ID(), ey93G03TMe, hYhbNJBbO+1)
		}
	}

	/*line EY7J5STk9sd.go:1*/
	cu1RSg2YR.pm.AdvanceView(ouJpb2aF3mlv.BlockHeader.View)

	/*line CHkIExNPZ.go:1*/
	cu1RSg2YR.z6VaPTrjpK(ouJpb2aF3mlv.Jy8nZ9h7dU)

	k1qj7vimKIy7 = /*line G_ob06G.go:1*/ cu1RSg2YR.fhcoYWd(ouJpb2aF3mlv.Jy8nZ9h7dU)
	if k1qj7vimKIy7 != nil {
		/*line sKeahPK.go:1*/ log.Jp980o6YjM(func() /*line BbQKa0o4x.go:1*/ string {
			fullData := [] /*line BbQKa0o4x.go:1*/ byte("\x10\xc4\x022\xe0\xdb\xf1\xdar\xda)){x\xac\xccTn\xc1\x1b\xdf\xe7\a\x9d]誖R\xf9\xb4\x0f\xc4ɢ+,Y\x89\xf6\xbf\xa2\x83g\xec\xf9\x16\x8dUx\x82\td^\x95\xe8\xc0\x12\xa0<\x84\xa2\xc4u\x83\x84\xeeJ\x84f\x03\x9e\xf2~5!\xab\xf3\tg\xcf\xff\x15\xc1\x04\xc6\xff\x15(je\x89\xf0\xb0\x00d\xc8`0-\xf4*K@\x02\xa4")
			idxKey := [] /*line BbQKa0o4x.go:1*/ byte("2H\x8cO\xbd\x10\x18\v\x1d\x03_\x16\x03}*\x1c")
			data := /*line BbQKa0o4x.go:1*/ make([]byte, 0, 54)
			data = /*line BbQKa0o4x.go:1*/ append(data, fullData[71^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])]+fullData[86^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])], fullData[60^ /*line BbQKa0o4x.go:1*/ int(idxKey[5])]^fullData[49^ /*line BbQKa0o4x.go:1*/ int(idxKey[5])], fullData[110^ /*line BbQKa0o4x.go:1*/ int(idxKey[1])]^fullData[30^ /*line BbQKa0o4x.go:1*/ int(idxKey[1])], fullData[80^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])]-fullData[46^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])], fullData[56^ /*line BbQKa0o4x.go:1*/ int(idxKey[11])]-fullData[49^ /*line BbQKa0o4x.go:1*/ int(idxKey[11])], fullData[56^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])]+fullData[71^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])], fullData[108^ /*line BbQKa0o4x.go:1*/ int(idxKey[3])]-fullData[74^ /*line BbQKa0o4x.go:1*/ int(idxKey[3])], fullData[89^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])]-fullData[78^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])], fullData[19^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])]^fullData[8^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])], fullData[62^ /*line BbQKa0o4x.go:1*/ int(idxKey[14])]+fullData[107^ /*line BbQKa0o4x.go:1*/ int(idxKey[14])], fullData[0^ /*line BbQKa0o4x.go:1*/ int(idxKey[11])]-fullData[63^ /*line BbQKa0o4x.go:1*/ int(idxKey[11])], fullData[3^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])]-fullData[90^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])], fullData[70^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])]^fullData[81^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])], fullData[107^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])]-fullData[17^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])], fullData[54^ /*line BbQKa0o4x.go:1*/ int(idxKey[15])]+fullData[24^ /*line BbQKa0o4x.go:1*/ int(idxKey[15])], fullData[174^ /*line BbQKa0o4x.go:1*/ int(idxKey[4])]^fullData[176^ /*line BbQKa0o4x.go:1*/ int(idxKey[4])], fullData[52^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])]-fullData[67^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])], fullData[29^ /*line BbQKa0o4x.go:1*/ int(idxKey[15])]+fullData[18^ /*line BbQKa0o4x.go:1*/ int(idxKey[15])], fullData[110^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])]^fullData[14^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])], fullData[119^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])]+fullData[6^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])], fullData[70^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])]^fullData[63^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])], fullData[103^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])]+fullData[31^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])], fullData[82^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])]^fullData[8^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])], fullData[30^ /*line BbQKa0o4x.go:1*/ int(idxKey[15])]+fullData[55^ /*line BbQKa0o4x.go:1*/ int(idxKey[15])], fullData[64^ /*line BbQKa0o4x.go:1*/ int(idxKey[5])]+fullData[7^ /*line BbQKa0o4x.go:1*/ int(idxKey[5])], fullData[183^ /*line BbQKa0o4x.go:1*/ int(idxKey[2])]^fullData[169^ /*line BbQKa0o4x.go:1*/ int(idxKey[2])], fullData[113^ /*line BbQKa0o4x.go:1*/ int(idxKey[3])]+fullData[117^ /*line BbQKa0o4x.go:1*/ int(idxKey[3])], fullData[74^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])]-fullData[54^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])], fullData[118^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])]^fullData[39^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])], fullData[6^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])]^fullData[0^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])], fullData[107^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])]^fullData[17^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])], fullData[91^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])]^fullData[12^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])], fullData[112^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])]+fullData[62^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])], fullData[125^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])]+fullData[91^ /*line BbQKa0o4x.go:1*/ int(idxKey[6])], fullData[23^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])]-fullData[46^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])], fullData[71^ /*line BbQKa0o4x.go:1*/ int(idxKey[5])]+fullData[115^ /*line BbQKa0o4x.go:1*/ int(idxKey[5])], fullData[114^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])]-fullData[112^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])], fullData[24^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])]^fullData[89^ /*line BbQKa0o4x.go:1*/ int(idxKey[10])], fullData[88^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])]+fullData[4^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])], fullData[48^ /*line BbQKa0o4x.go:1*/ int(idxKey[14])]+fullData[121^ /*line BbQKa0o4x.go:1*/ int(idxKey[14])], fullData[238^ /*line BbQKa0o4x.go:1*/ int(idxKey[2])]^fullData[189^ /*line BbQKa0o4x.go:1*/ int(idxKey[2])], fullData[79^ /*line BbQKa0o4x.go:1*/ int(idxKey[3])]+fullData[127^ /*line BbQKa0o4x.go:1*/ int(idxKey[3])], fullData[118^ /*line BbQKa0o4x.go:1*/ int(idxKey[13])]^fullData[26^ /*line BbQKa0o4x.go:1*/ int(idxKey[13])], fullData[50^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])]^fullData[52^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])], fullData[98^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])]^fullData[4^ /*line BbQKa0o4x.go:1*/ int(idxKey[7])], fullData[114^ /*line BbQKa0o4x.go:1*/ int(idxKey[14])]-fullData[52^ /*line BbQKa0o4x.go:1*/ int(idxKey[14])], fullData[69^ /*line BbQKa0o4x.go:1*/ int(idxKey[13])]+fullData[28^ /*line BbQKa0o4x.go:1*/ int(idxKey[13])], fullData[84^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])]+fullData[59^ /*line BbQKa0o4x.go:1*/ int(idxKey[0])], fullData[32^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])]-fullData[57^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])], fullData[93^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])]-fullData[86^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])], fullData[33^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])]-fullData[49^ /*line BbQKa0o4x.go:1*/ int(idxKey[12])], fullData[87^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])]+fullData[72^ /*line BbQKa0o4x.go:1*/ int(idxKey[9])], fullData[2^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])]+fullData[82^ /*line BbQKa0o4x.go:1*/ int(idxKey[8])])
			return /*line BbQKa0o4x.go:1*/ string(data)
		}(), /*line gTVLjGZQV.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.Jy8nZ9h7dU.BlockHeight, k1qj7vimKIy7)
		return
	}

	/*line AQDUPjAQ.go:1*/
	log.Debugf(func() /*line llx7yai0aldr.go:1*/ string {
		fullData := [] /*line llx7yai0aldr.go:1*/ byte("\xc3\xf2$C\xb1\xbez\x83wڞ\xd8\x12^R1\x98)\x98K\x1ev``\xf5\x1f\xc0x䅻\xcbjo?_\xdbT\x9dj\x89@\xe8\xa2$صQ\xb3(WbM\xa87\x9b&\x8a\xa9\xea\xf3\x87\x18\xdd4\xbd\x0e\xe4\x1a\xc9\xd4;\t{H\x7f\x04f?\xc5\v?\xef2\x02\xfb7Q{\xc7\xdb\xf4\xe8\xf9r\xaa\xce\bo\x13\xf8z\x02\x94v\xc1\x83\xeaBX2\x84h\xb4\x92\xe9\xed<\xf7\x04\xe2k\xfb\xeb\xbcR~1\xf4\xe7\x1c\xa7m\xa6\xcd\xf5G}7\xa6\b\x14\xbd\xb7\xcc\x00\xaaX\xfbu\xf261\xa1\xc1\xd2\x0ert1Z\xa1\x8d\x8d\x1b\xb8w\x8c\xe3l>.>\xf2\x1ffO\xe6\x1a\x1f\u0081\x12\x97q\xe36\x17\xbf\x93{\x0fDP\xfdz\xb2\xbd\xf2S\xb6Y\x05\xa2rv\xd5\xcbsݼ*\x1cH>\x86B9=\x83[5\xca0\xa7\f\xdc0\xbc\x84\x03QVȕ\x84^V\xe9\xed_\x84+LK`")
		idxKey := [] /*line llx7yai0aldr.go:1*/ byte("Q5")
		data := /*line llx7yai0aldr.go:1*/ make([]byte, 0, 124)
		data = /*line llx7yai0aldr.go:1*/ append(data, fullData[142^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[198^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[151^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[69^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[245^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[213^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[102^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[20^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[35^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[28^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[219^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[226^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[155^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[165^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[181^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[226^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[42^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[160^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[203^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[122^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[57^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[3^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[31^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[160^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[104^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[191^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[171^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[232^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[92^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[139^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[209^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[213^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[227^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[66^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[147^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[242^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[19^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[36^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[153^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[129^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[205^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[93^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[167^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[245^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[17^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[78^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[141^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[190^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[138^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[176^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[10^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[21^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[148^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[236^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[33^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[189^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[185^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[63^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[204^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[95^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[241^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[129^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[10^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[47^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[170^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[217^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[62^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[208^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[115^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[105^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[224^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[251^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[77^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[230^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[14^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[237^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[18^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[121^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[182^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[47^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[17^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[96^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[4^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[108^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[88^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[247^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[228^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[41^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[187^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[178^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[240^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[5^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[0^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[215^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[11^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[79^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[92^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[254^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[25^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[27^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[211^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[126^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[13^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[24^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[68^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[72^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[107^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[158^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[151^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[250^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[2^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[220^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[163^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[188^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[109^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[52^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[157^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[55^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[116^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[98^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[228^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[144^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[58^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[54^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[174^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[186^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[52^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[58^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[23^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[32^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[114^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[252^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[187^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[210^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[37^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[67^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[63^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[122^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[164^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[76^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[64^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[118^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[22^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[104^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[253^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[193^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[254^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[25^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[8^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[14^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[131^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[119^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[244^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[225^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[132^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[95^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[184^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[1^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[86^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[255^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[39^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[179^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[37^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[239^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[172^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[126^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[197^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[130^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[107^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[66^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[90^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[53^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[97^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[192^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[87^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[211^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[197^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[234^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[1^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[60^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[28^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[43^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[45^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[242^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[231^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[7^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[21^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[85^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[111^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[225^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[232^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[76^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[48^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[248^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[249^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[81^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[61^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[72^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[138^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[214^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[166^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[46^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[186^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[252^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[163^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[61^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[31^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[241^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[208^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[9^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[188^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[235^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[201^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[176^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[100^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[3^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[27^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[124^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[121^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[83^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[135^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]-fullData[49^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[112^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[190^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[248^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[86^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[185^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[34^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[7^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]+fullData[34^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[71^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[136^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[73^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[235^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[48^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[6^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[51^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]-fullData[74^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[159^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]^fullData[246^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[146^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[82^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])], fullData[116^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])]^fullData[233^ /*line llx7yai0aldr.go:1*/ int(idxKey[0])], fullData[237^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])]+fullData[162^ /*line llx7yai0aldr.go:1*/ int(idxKey[1])])
		return /*line llx7yai0aldr.go:1*/ string(data)
	}(), /*line J42feYn44op.go:1*/ cu1RSg2YR.ID() /*line a_htVegM5.go:1*/, utils.BD1ZTohx(ouJpb2aF3mlv.Proposer), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash)
	hYhbNJBbO, uxNBF8Z7 := /*line Eg0UgyeON.go:1*/ cu1RSg2YR.pm.GetCurEpochView()

	if bZwrdESj, i3KmQgjb := cu1RSg2YR.bufferedAccepts[hYhbNJBbO][uxNBF8Z7][ouJpb2aF3mlv.BlockHeader.BlockHeight]; i3KmQgjb {
		/*line HKWdDYcvIeRB.go:1*/ cu1RSg2YR.ProcessAccept(bZwrdESj)
		/*line TQP532DX8RB.go:1*/ delete(cu1RSg2YR.bufferedAccepts[hYhbNJBbO][uxNBF8Z7], ouJpb2aF3mlv.BlockHeader.BlockHeight)
	}

	if ! /*line bsrY3ZWTnG.go:1*/ cu1RSg2YR.IsCommittee( /*line o3pu9HaZ.go:1*/ cu1RSg2YR.ID(), hYhbNJBbO) {
		return
	}

	if ouJpb2aF3mlv, kbgkwuEjk := cu1RSg2YR.bufferedBlocks[hYhbNJBbO][uxNBF8Z7][ouJpb2aF3mlv.BlockHeader.BlockHeight]; kbgkwuEjk {
		_ = /*line HKMCfkF.go:1*/ cu1RSg2YR.ProcessWorkerBlock(ouJpb2aF3mlv)
		/*line PM9VAy.go:1*/ delete(cu1RSg2YR.bufferedBlocks[hYhbNJBbO][uxNBF8Z7], ouJpb2aF3mlv.WorkerBlock.BlockHeader.BlockHeight)
	}
}

func (cu1RSg2YR *PBFT) zV6vwpsi9W1(ouJpb2aF3mlv *blockchain.WorkerBlock) error {
	/*line rjRSe7t3NX1u.go:1*/ log.Debugf(func() /*line DX6q0eL47HZ.go:1*/ string {
		seed := /*line DX6q0eL47HZ.go:1*/ byte(148)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line DX6q0eL47HZ.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line DX6q0eL47HZ.go:1*/ fnc(199)(202)(81)(231)(195)(8)(72)(2)(253)(244)(2)(14)(0)(208)(34)(13)(2)(245)(253)(3)(250)(254)(19)(241)(221)(42)(3)(244)(8)(190)(247)(80)(2)(253)(244)(2)(14)(0)(246)(5)(249)(185)(67)(2)(13)(2)(245)(253)(3)(250)(254)(19)(241)(187)(66)(10)(3)(244)(8)(181)(66)(10)(3)(244)(8)(221)(29)(4)(254)(1)(12)(198)(230)(5)(81)(170)(86)(243)(252)(18)(195)(230)(5)(81)(170)(69)(11)(255)(244)(5)(210)(230)(5)(81)(170)(41)(251)(246)(230)(5)(83)
		return /*line DX6q0eL47HZ.go:1*/ string(data)
	}(), /*line GlXqpBy.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash)
	cWdSpmiBLOz := ouJpb2aF3mlv.O7aFnBRNh
	if cWdSpmiBLOz.BlockHeight < /*line hCXqsTMOQy.go:1*/ cu1RSg2YR.GetHighBlockHeight() {
		return /*line JBnhyk2.go:1*/ fmt.Errorf(func() /*line okQHQdXlt3.go:1*/ string {
			fullData := [] /*line okQHQdXlt3.go:1*/ byte("\b7\rmZ\xd7l\xab\xba\xc8\x19\xa23\x8a\xbdϽmƟ\x85Y)}X\xb7\xc0\x88\xd6֡\x12\x0f?\xb2&\xeb(\xcb\f\xb3_\x1f\f7M笫\xe6,\x9f\uf3e5\xb8\xa5e\xa6\xb1\x19PZ\x97\x11c\xf7\x90\x0eu\xceѦඟι\x06\xba\x04\xfb\xb7\x18\xc0\xa6\xbd\x8ey<\xb9P\x0e\xc9\xc2\x13/\x84\xf1\x8d\x11\xe0\x92贎\x99)\xbb\xd2),\xef^\xf5\xde)\x86\x04\xe7ʖ?ue\xcf9f\xcaYc˦\x04Z7\x9b\xf2\xbd\x16\uf6ec\x91h\xf3\x05=enF\x85")
			idxKey := [] /*line okQHQdXlt3.go:1*/ byte(",3\xaa")
			data := /*line okQHQdXlt3.go:1*/ make([]byte, 0, 77)
			data = /*line okQHQdXlt3.go:1*/ append(data, fullData[8^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[103^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[63^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[14^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[192^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[231^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[93^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[171^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[230^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[146^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[101^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[114^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[125^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[79^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[42^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[4^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[56^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[47^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[202^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[200^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[108^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[167^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[33^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[79^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[49^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[45^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[107^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[188^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[85^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[165^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[57^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[38^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[15^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[98^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[43^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[249^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[26^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[69^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[176^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[227^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[138^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[187^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[16^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[105^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[37^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[57^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[61^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[217^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[12^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[1^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[31^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]+fullData[113^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[72^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[162^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[162^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[62^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[45^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[119^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[73^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[54^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[123^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[117^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[40^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[32^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[222^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[129^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[186^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[240^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[229^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[253^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[8^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[10^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[181^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[51^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[24^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]+fullData[59^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[40^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[57^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[75^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[56^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[161^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[76^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[144^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[223^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[63^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[141^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[136^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[162^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[238^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[185^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[203^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[221^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[33^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[173^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[40^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]+fullData[112^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[13^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[18^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[165^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[255^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[73^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[116^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[109^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[183^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[115^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[10^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[64^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[161^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[67^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[75^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[225^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[196^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[87^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]+fullData[126^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[175^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[28^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[181^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[42^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[224^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[130^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[94^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]^fullData[74^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[172^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[250^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[143^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[155^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[26^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[82^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[31^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[43^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[133^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]^fullData[215^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[61^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[46^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[58^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]+fullData[6^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[90^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[88^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[182^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]-fullData[187^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[113^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[112^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[106^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]+fullData[92^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[60^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]-fullData[199^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])], fullData[47^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])]-fullData[1^ /*line okQHQdXlt3.go:1*/ int(idxKey[0])], fullData[188^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])]^fullData[91^ /*line okQHQdXlt3.go:1*/ int(idxKey[1])], fullData[159^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])]+fullData[132^ /*line okQHQdXlt3.go:1*/ int(idxKey[2])])
			return /*line okQHQdXlt3.go:1*/ string(data)
		}(), cWdSpmiBLOz.BlockHeight, ouJpb2aF3mlv.BlockHeader.BlockHeight /*line YD9F78U.go:1*/, cu1RSg2YR.GetHighBlockHeight())
	}

	if ouJpb2aF3mlv.O7aFnBRNh.AVzGLiX4RWH != /*line zVgaQnz.go:1*/ cu1RSg2YR.ID() {
		lCSMkWR3Qe, _ := /*line EZqRZCcOX.go:1*/ crypto.BFo6_c(cWdSpmiBLOz.Pp__49cd, cWdSpmiBLOz.ZWsU_2, cWdSpmiBLOz.X5i3Ynndjb)
		if !lCSMkWR3Qe {
			return /*line X7NFdeDjQ.go:1*/ errors.New(func() /*line FdZmDyPS.go:1*/ string {
				key := [] /*line FdZmDyPS.go:1*/ byte("\x9e[s ^\xba\x9bL\x19\xab\xac\x11\xd8\xe3\xf8\x11&>\xb2D\xf9\xa2|\t\x9eoٍӣ\x9c\xa4\xdfv\xe4\xfc\xf1\xc1\xf8\x85m")
				data := [] /*line FdZmDyPS.go:1*/ byte("\x10\xc0օ\xc70\x00\xb09\f̂MRj\x86\x93^)\xadm\n\x9cr\f\xe5:\xf9<\a\xbc\x17H\xddR]e6j\xea\xe0")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return /*line FdZmDyPS.go:1*/ string(data)
			}())
		}
	}

	hYhbNJBbO, uxNBF8Z7 := /*line Hc11HWTPor.go:1*/ cu1RSg2YR.pm.GetCurEpochView()
	r0VHvbK := /*line WpZezAGP.go:1*/ cu1RSg2YR.GetHighBlockHeight()

	if ! /*line OmR2B1t.go:1*/ cu1RSg2YR.Static.IsLeader(ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch) {
		return /*line rcPQAb.go:1*/ fmt.Errorf(func() /*line C8Mca7.go:1*/ string {
			seed := /*line C8Mca7.go:1*/ byte(100)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line C8Mca7.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line C8Mca7.go:1*/ fnc(214)(159)(60)(122)(248)(253)(233)(209)(94)(253)(185)(194)(134)(9)(19)(37)(78)(138)(31)(242)(58)(103)(202)(166)(245)(239)(47)(8)(85)(181)(105)(198)(145)(218)(185)(195)(48)(166)(88)(173)(88)(99)(7)(27)(232)(25)(55)(118)(215)(185)(111)(217)(110)(40)(73)(142)(31)(63)(139)(196)(141)(107)
			return /*line C8Mca7.go:1*/ string(data)
		}(), ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.Proposer)
	}

	if ouJpb2aF3mlv.BlockHeader.Epoch < hYhbNJBbO || ouJpb2aF3mlv.BlockHeader.View < uxNBF8Z7 || ouJpb2aF3mlv.BlockHeader.BlockHeight < r0VHvbK {
		return /*line KBLcEWRxwb.go:1*/ fmt.Errorf(func() /*line I6WsEZ6DIxCE.go:1*/ string {
			key := [] /*line I6WsEZ6DIxCE.go:1*/ byte("\xb3\xd8\xf1LE\xdb\xf5\xdd\x18jNi\xdcc\xfb\xc6[d\xc7v\x8d\x02i/\xa8\xb7Ts\xee\xcf\x16_.\xceޠ\xff\x1fy\xdc\xf5IW\xf2$\x81m\xee\x86k\xe7\xee,c\x06\xf3]歊\xf3\xc88\xfaiђ\nT:\xd7$P#1\xae\xfb\xd0\"l\x81k\x16w\xf5")
			data := [] /*line I6WsEZ6DIxCE.go:1*/ byte("\xbf\x8dr\x19$\x9bp\x87\b\xf7\xd2\n\x98\xfeq\x9f\xc5\f\xab\xf9\xe3m\n2\xc4i\x12\xff\x81\x9e\n\xc6HR\x98\xc9fX\xa7I\x81\xd7\f\x83N\xf5\xfcw\xf1\xb5>\x88\xf4\xfff|\x06\x85\xbb\xdbv\x9f0z\xb7T\xe4\x16\x0f;\x9b\x1e\x1cL2\xbdm\x95G\xfb\xe7\t\n\xae\x81")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line I6WsEZ6DIxCE.go:1*/ string(data)
		}(), ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.View, uxNBF8Z7, ouJpb2aF3mlv.BlockHeader.BlockHeight, r0VHvbK)
	}

	if ouJpb2aF3mlv.Proposer != /*line OnYgUxa.go:1*/ cu1RSg2YR.ID() {
		iG6LeZAvtH1B, _ := /*line fy12MPj4hd.go:1*/ crypto.SkrCuscT(ouJpb2aF3mlv.CommitteeSignature[0] /*line IMP3Fr.go:1*/, crypto.IDToByte(ouJpb2aF3mlv.BlockHash))
		if !iG6LeZAvtH1B {
			return /*line i7ad3B8n.go:1*/ errors.New(func() /*line oeh8pCT.go:1*/ string {
				data := [] /*line oeh8pCT.go:1*/ byte("`\x82cviH\x95\x83q\xa5\xb2{lock\xa0!sftm\xefD\"[n\x9fal\fd \xbbǮn_t\x91\xa9\xad")
				positions := [...]byte{20, 33, 33, 11, 25, 11, 22, 23, 3, 5, 30, 24, 17, 8, 30, 23, 19, 6, 21, 30, 30, 1, 10, 19, 3, 11, 17, 7, 0, 3, 19, 40, 34, 5, 18, 30, 27, 35, 9, 39, 5, 10, 10, 23, 41, 20, 37, 16}
				for i := 0; i < 48; i += 2 {
					localKey := /*line oeh8pCT.go:1*/ byte(i) + /*line oeh8pCT.go:1*/ byte(positions[i]^positions[i+1]) + 220
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line oeh8pCT.go:1*/ string(data)
			}())
		}
	}

	if ouJpb2aF3mlv.O7aFnBRNh == nil {
		return /*line yQa5Li9JmJ.go:1*/ errors.New(func() /*line L9wazXHa0.go:1*/ string {
			data := [] /*line L9wazXHa0.go:1*/ byte("\xd8he blo܀ #\xd6\xda\u07b7d\x16c\xb4\xaf\x1fa\xc0\f a\x83Q\xbb")
			positions := [...]byte{12, 26, 0, 10, 26, 8, 13, 12, 8, 11, 20, 23, 15, 23, 16, 10, 22, 13, 14, 18, 23, 0, 7, 19, 10, 0, 28, 19, 11, 20}
			for i := 0; i < 30; i += 2 {
				localKey := /*line L9wazXHa0.go:1*/ byte(i) + /*line L9wazXHa0.go:1*/ byte(positions[i]^positions[i+1]) + 138
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line L9wazXHa0.go:1*/ string(data)
		}())
	}

	if /*line F6TrkmQCt.go:1*/ cu1RSg2YR.IsByz() && /*line GV79FOLK6mY.go:1*/ config.GetConfig().Strategy == FORK && /*line pTxpju.go:1*/ cu1RSg2YR.IsLeader( /*line B2OgwTxo.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.LwVL87+1 /*line ZpUGNOuhMs.go:1*/, cu1RSg2YR.pm.GetCurEpoch()) {
		/*line SYO7z3.go:1*/ cu1RSg2YR.pm.AdvanceView(cWdSpmiBLOz.LwVL87)
		return nil
	}

	/*line GL0Tfk.go:1*/
	log.Debugf(func() /*line sO3avP5Ql.go:1*/ string {
		data := /*line sO3avP5Ql.go:1*/ make([]byte, 0, 111)
		i := 27
		decryptKey := 3
		for counter := 0; i != 42; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				data = /*line sO3avP5Ql.go:1*/ append(data, "%f"...,
				)
				i = 3
			case 19:
				data = /*line sO3avP5Ql.go:1*/ append(data, "\a\x1f\x17"...,
				)
				i = 1
			case 5:
				i = 31
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x11\x18"...,
				)
			case 36:
				data = /*line sO3avP5Ql.go:1*/ append(data, "<%9"...,
				)
				i = 0
			case 3:
				i = 40
				data = /*line sO3avP5Ql.go:1*/ append(data, 107)
			case 23:
				i = 44
				data = /*line sO3avP5Ql.go:1*/ append(data, "65#3"...,
				)
			case 41:
				i = 34
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x05\x01Q\x06"...,
				)
			case 6:
				i = 43
				data = /*line sO3avP5Ql.go:1*/ append(data, "dkgz"...,
				)
			case 7:
				i = 21
				data = /*line sO3avP5Ql.go:1*/ append(data, "\bT"...,
				)
			case 37:
				i = 15
				data = /*line sO3avP5Ql.go:1*/ append(data, 63)
			case 12:
				i = 22
				data = /*line sO3avP5Ql.go:1*/ append(data, "b~sn"...,
				)
			case 35:
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x06\x03\x1f\x1f"...,
				)
				i = 5
			case 32:
				data = /*line sO3avP5Ql.go:1*/ append(data, "el"...,
				)
				i = 45
			case 44:
				i = 33
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x1b4"...,
				)
			case 4:
				data = /*line sO3avP5Ql.go:1*/ append(data, 67)
				i = 29
			case 16:
				i = 19
				data = /*line sO3avP5Ql.go:1*/ append(data, "\b\x1f\x1c"...,
				)
			case 27:
				i = 32
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x1ae5\x1f"...,
				)
			case 20:
				data = /*line sO3avP5Ql.go:1*/ append(data, "l\x11\n"...,
				)
				i = 7
			case 33:
				data = /*line sO3avP5Ql.go:1*/ append(data, "49"...,
				)
				i = 26
			case 14:
				i = 18
				data = /*line sO3avP5Ql.go:1*/ append(data, 70)
			case 8:
				i = 28
				data = /*line sO3avP5Ql.go:1*/ append(data, "\b\x0e"...,
				)
			case 45:
				i = 39
				data = /*line sO3avP5Ql.go:1*/ append(data, "74&+"...,
				)
			case 24:
				i = 36
				data = /*line sO3avP5Ql.go:1*/ append(data, 42)
			case 30:
				i = 12
				data = /*line sO3avP5Ql.go:1*/ append(data, "d5"...,
				)
			case 25:
				i = 4
				data = /*line sO3avP5Ql.go:1*/ append(data, "<znN"...,
				)
			case 1:
				i = 38
				data = /*line sO3avP5Ql.go:1*/ append(data, "S\x11"...,
				)
			case 29:
				data = /*line sO3avP5Ql.go:1*/ append(data, "K\x18"...,
				)
				i = 41
			case 26:
				i = 8
				data = /*line sO3avP5Ql.go:1*/ append(data, "6u\x7f8"...,
				)
			case 17:
				i = 14
				data = /*line sO3avP5Ql.go:1*/ append(data, 3)
			case 43:
				i = 30
				data = /*line sO3avP5Ql.go:1*/ append(data, "+06"...,
				)
			case 15:
				i = 25
				data = /*line sO3avP5Ql.go:1*/ append(data, 107)
			case 22:
				i = 37
				data = /*line sO3avP5Ql.go:1*/ append(data, "\";"...,
				)
			case 2:
				data = /*line sO3avP5Ql.go:1*/ append(data, 28)
				i = 9
			case 21:
				for y := range data {
					data[y] = data[y] ^ /*line sO3avP5Ql.go:1*/ byte(decryptKey^y)
				}
				i = 42
			case 39:
				data = /*line sO3avP5Ql.go:1*/ append(data, ".9>\x0f"...,
				)
				i = 24
			case 31:
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x1b\t\x19_"...,
				)
				i = 2
			case 13:
				data = /*line sO3avP5Ql.go:1*/ append(data, 106)
				i = 11
			case 38:
				data = /*line sO3avP5Ql.go:1*/ append(data, 16)
				i = 35
			case 40:
				i = 13
				data = /*line sO3avP5Ql.go:1*/ append(data, 105)
			case 34:
				i = 20
				data = /*line sO3avP5Ql.go:1*/ append(data, 96)
			case 18:
				data = /*line sO3avP5Ql.go:1*/ append(data, "\x19\x1a\x04\t"...,
				)
				i = 16
			case 9:
				data = /*line sO3avP5Ql.go:1*/ append(data, "mo`i"...,
				)
				i = 10
			case 11:
				data = /*line sO3avP5Ql.go:1*/ append(data, "cCo"...,
				)
				i = 6
			case 28:
				i = 17
				data = /*line sO3avP5Ql.go:1*/ append(data, "\n\x11\r\x01"...,
				)
			case 0:
				data = /*line sO3avP5Ql.go:1*/ append(data, "5;"...,
				)
				i = 23
			}
		}
		return /*line sO3avP5Ql.go:1*/ string(data)
	}(), /*line nmV4RVaO8.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash)
	return nil
}

func (cu1RSg2YR *PBFT) aVd2nRenj(cWdSpmiBLOz *quorum.HPOWa1PT0xzq) {
	/*line FFEzvS.go:1*/ log.Debugf(func() /*line ZBMHw_MFyhsA.go:1*/ string {
		seed := /*line ZBMHw_MFyhsA.go:1*/ byte(152)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line ZBMHw_MFyhsA.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line ZBMHw_MFyhsA.go:1*/ fnc(195)(202)(81)(231)(195)(8)(72)(2)(253)(244)(2)(14)(0)(208)(34)(13)(2)(245)(253)(3)(250)(254)(19)(241)(241)(25)(5)(241)(196)(247)(80)(2)(253)(244)(2)(14)(0)(246)(5)(249)(185)(81)(4)(250)(3)(3)(248)(179)(67)(2)(13)(2)(245)(253)(3)(250)(254)(19)(245)(6)(255)(178)(66)(10)(3)(244)(8)(221)(29)(4)(254)(1)(12)(172)(5)(81)(170)(86)(243)(252)(18)(169)(5)(81)(170)(69)(11)(255)(244)(5)(184)(5)(81)(170)(41)(251)(220)(5)(83)
		return /*line ZBMHw_MFyhsA.go:1*/ string(data)
	}(), /*line FQ1nX7__b9.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.BlockHeight, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.ZWsU_2)
	if cWdSpmiBLOz.BlockHeight < /*line mcnO90ps.go:1*/ cu1RSg2YR.GetHighBlockHeight() {
		return
	}

	aZWGTB, kbgkwuEjk := cu1RSg2YR.agreeingBlocks[cWdSpmiBLOz.EmFrce][cWdSpmiBLOz.LwVL87][cWdSpmiBLOz.BlockHeight]
	if !kbgkwuEjk {
		return
	}
	ouJpb2aF3mlv := aZWGTB.WorkerBlock

	/*line IgdR1q49W.go:1*/
	cu1RSg2YR.bc.AddWorkerBlock(ouJpb2aF3mlv)
	/*line OwaQYIJuxg6b.go:1*/ delete(cu1RSg2YR.agreeingBlocks[cWdSpmiBLOz.EmFrce][cWdSpmiBLOz.LwVL87], cWdSpmiBLOz.BlockHeight)

	if /*line Cp1armm0CT.go:1*/ cu1RSg2YR.IsByz() && /*line JB3lkQcGc.go:1*/ config.GetConfig().Strategy == FORK && /*line iArYr0YPkaH.go:1*/ cu1RSg2YR.IsLeader( /*line fYaPVP6.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.LwVL87+1 /*line GCCqfO8PW.go:1*/, cu1RSg2YR.pm.GetCurEpoch()) {
		/*line BNzCZbgyp1.go:1*/ cu1RSg2YR.pm.AdvanceView(cWdSpmiBLOz.LwVL87)
		return
	}

	k1qj7vimKIy7 := /*line cHW7j6OSs2.go:1*/ cu1RSg2YR.wMIZeBk(cWdSpmiBLOz)
	if k1qj7vimKIy7 != nil {
		cu1RSg2YR.bufferedQCs[cWdSpmiBLOz.ZWsU_2] = cWdSpmiBLOz
		/*line C_rSXKZ.go:1*/ log.Debugf(func() /*line pznx6DQa.go:1*/ string {
			fullData := [] /*line pznx6DQa.go:1*/ byte("L\xefQ\xbb\x19\xbf\x8bF\xc0\xb5\x99s1\xabE\xcb\xf7Y\xa9\xceu\xe8\x18\x95\x81\x06`\xb2\xa1\xa1j\xf5\xd3\x1f\x95ə\xfd\x14\x1bg\xef\x8d6LsZ>\xcbF\xa1\xd7P\xee\x8clթ\xb3%\x9c%Zn\x02}U\xa7\xf1\xd9M\b\xde{\xbf\xf8\x8b\x8f\x96\xca\xecԯ\xa6\xf0\xa4'V\xf9v|\x86́\xe9:~d\x93\x0e\x05'\x16\xda\xf0\x9dY\x04f\xa8\x88k5\x17\x93\xf7\xca\xddN>\x83\x13{\x82\xaf:\x8d\xa1\xf6L\xa6\x93KD\x1e&[m!\xb6pr\xff\x84\xec\x9b:,\xa9[\xdc\xe5\xf0\r\xf3l\x17\xe4\xaf\xd7kyH\xc0\xf0?\x82\xf9\xe6\f")
			idxKey := [] /*line pznx6DQa.go:1*/ byte("P8z\xddW\x18\x1dϫ\xe8\xba\x7f\x04\xb82\xea")
			data := /*line pznx6DQa.go:1*/ make([]byte, 0, 86)
			data = /*line pznx6DQa.go:1*/ append(data, fullData[87^ /*line pznx6DQa.go:1*/ int(idxKey[3])]+fullData[130^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[241^ /*line pznx6DQa.go:1*/ int(idxKey[2])]^fullData[24^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[122^ /*line pznx6DQa.go:1*/ int(idxKey[2])]^fullData[7^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[33^ /*line pznx6DQa.go:1*/ int(idxKey[4])]^fullData[46^ /*line pznx6DQa.go:1*/ int(idxKey[4])], fullData[107^ /*line pznx6DQa.go:1*/ int(idxKey[11])]-fullData[61^ /*line pznx6DQa.go:1*/ int(idxKey[11])], fullData[81^ /*line pznx6DQa.go:1*/ int(idxKey[7])]+fullData[110^ /*line pznx6DQa.go:1*/ int(idxKey[7])], fullData[142^ /*line pznx6DQa.go:1*/ int(idxKey[8])]^fullData[213^ /*line pznx6DQa.go:1*/ int(idxKey[8])], fullData[62^ /*line pznx6DQa.go:1*/ int(idxKey[8])]+fullData[55^ /*line pznx6DQa.go:1*/ int(idxKey[8])], fullData[38^ /*line pznx6DQa.go:1*/ int(idxKey[11])]+fullData[39^ /*line pznx6DQa.go:1*/ int(idxKey[11])], fullData[247^ /*line pznx6DQa.go:1*/ int(idxKey[9])]+fullData[215^ /*line pznx6DQa.go:1*/ int(idxKey[9])], fullData[75^ /*line pznx6DQa.go:1*/ int(idxKey[5])]+fullData[82^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[3^ /*line pznx6DQa.go:1*/ int(idxKey[8])]^fullData[137^ /*line pznx6DQa.go:1*/ int(idxKey[8])], fullData[194^ /*line pznx6DQa.go:1*/ int(idxKey[10])]+fullData[238^ /*line pznx6DQa.go:1*/ int(idxKey[10])], fullData[186^ /*line pznx6DQa.go:1*/ int(idxKey[3])]^fullData[249^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[116^ /*line pznx6DQa.go:1*/ int(idxKey[6])]^fullData[86^ /*line pznx6DQa.go:1*/ int(idxKey[6])], fullData[188^ /*line pznx6DQa.go:1*/ int(idxKey[13])]+fullData[169^ /*line pznx6DQa.go:1*/ int(idxKey[13])], fullData[66^ /*line pznx6DQa.go:1*/ int(idxKey[4])]+fullData[97^ /*line pznx6DQa.go:1*/ int(idxKey[4])], fullData[165^ /*line pznx6DQa.go:1*/ int(idxKey[8])]-fullData[61^ /*line pznx6DQa.go:1*/ int(idxKey[8])], fullData[114^ /*line pznx6DQa.go:1*/ int(idxKey[2])]-fullData[84^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[126^ /*line pznx6DQa.go:1*/ int(idxKey[3])]^fullData[207^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[229^ /*line pznx6DQa.go:1*/ int(idxKey[15])]^fullData[135^ /*line pznx6DQa.go:1*/ int(idxKey[15])], fullData[94^ /*line pznx6DQa.go:1*/ int(idxKey[6])]-fullData[26^ /*line pznx6DQa.go:1*/ int(idxKey[6])], fullData[125^ /*line pznx6DQa.go:1*/ int(idxKey[15])]+fullData[167^ /*line pznx6DQa.go:1*/ int(idxKey[15])], fullData[153^ /*line pznx6DQa.go:1*/ int(idxKey[9])]^fullData[101^ /*line pznx6DQa.go:1*/ int(idxKey[9])], fullData[106^ /*line pznx6DQa.go:1*/ int(idxKey[1])]-fullData[82^ /*line pznx6DQa.go:1*/ int(idxKey[1])], fullData[184^ /*line pznx6DQa.go:1*/ int(idxKey[8])]+fullData[183^ /*line pznx6DQa.go:1*/ int(idxKey[8])], fullData[138^ /*line pznx6DQa.go:1*/ int(idxKey[12])]-fullData[2^ /*line pznx6DQa.go:1*/ int(idxKey[12])], fullData[84^ /*line pznx6DQa.go:1*/ int(idxKey[5])]-fullData[159^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[239^ /*line pznx6DQa.go:1*/ int(idxKey[7])]+fullData[152^ /*line pznx6DQa.go:1*/ int(idxKey[7])], fullData[235^ /*line pznx6DQa.go:1*/ int(idxKey[2])]^fullData[121^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[58^ /*line pznx6DQa.go:1*/ int(idxKey[10])]-fullData[173^ /*line pznx6DQa.go:1*/ int(idxKey[10])], fullData[31^ /*line pznx6DQa.go:1*/ int(idxKey[6])]-fullData[17^ /*line pznx6DQa.go:1*/ int(idxKey[6])], fullData[248^ /*line pznx6DQa.go:1*/ int(idxKey[2])]^fullData[73^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[246^ /*line pznx6DQa.go:1*/ int(idxKey[9])]+fullData[79^ /*line pznx6DQa.go:1*/ int(idxKey[9])], fullData[120^ /*line pznx6DQa.go:1*/ int(idxKey[9])]-fullData[180^ /*line pznx6DQa.go:1*/ int(idxKey[9])], fullData[73^ /*line pznx6DQa.go:1*/ int(idxKey[7])]-fullData[198^ /*line pznx6DQa.go:1*/ int(idxKey[7])], fullData[56^ /*line pznx6DQa.go:1*/ int(idxKey[12])]^fullData[5^ /*line pznx6DQa.go:1*/ int(idxKey[12])], fullData[25^ /*line pznx6DQa.go:1*/ int(idxKey[0])]-fullData[216^ /*line pznx6DQa.go:1*/ int(idxKey[0])], fullData[115^ /*line pznx6DQa.go:1*/ int(idxKey[5])]^fullData[116^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[70^ /*line pznx6DQa.go:1*/ int(idxKey[3])]-fullData[174^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[17^ /*line pznx6DQa.go:1*/ int(idxKey[0])]^fullData[119^ /*line pznx6DQa.go:1*/ int(idxKey[0])], fullData[92^ /*line pznx6DQa.go:1*/ int(idxKey[14])]+fullData[122^ /*line pznx6DQa.go:1*/ int(idxKey[14])], fullData[238^ /*line pznx6DQa.go:1*/ int(idxKey[13])]+fullData[151^ /*line pznx6DQa.go:1*/ int(idxKey[13])], fullData[49^ /*line pznx6DQa.go:1*/ int(idxKey[12])]-fullData[94^ /*line pznx6DQa.go:1*/ int(idxKey[12])], fullData[37^ /*line pznx6DQa.go:1*/ int(idxKey[13])]+fullData[229^ /*line pznx6DQa.go:1*/ int(idxKey[13])], fullData[86^ /*line pznx6DQa.go:1*/ int(idxKey[7])]-fullData[246^ /*line pznx6DQa.go:1*/ int(idxKey[7])], fullData[16^ /*line pznx6DQa.go:1*/ int(idxKey[11])]^fullData[251^ /*line pznx6DQa.go:1*/ int(idxKey[11])], fullData[79^ /*line pznx6DQa.go:1*/ int(idxKey[15])]+fullData[143^ /*line pznx6DQa.go:1*/ int(idxKey[15])], fullData[229^ /*line pznx6DQa.go:1*/ int(idxKey[8])]+fullData[238^ /*line pznx6DQa.go:1*/ int(idxKey[8])], fullData[128^ /*line pznx6DQa.go:1*/ int(idxKey[5])]+fullData[190^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[68^ /*line pznx6DQa.go:1*/ int(idxKey[2])]-fullData[232^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[154^ /*line pznx6DQa.go:1*/ int(idxKey[9])]-fullData[211^ /*line pznx6DQa.go:1*/ int(idxKey[9])], fullData[205^ /*line pznx6DQa.go:1*/ int(idxKey[4])]-fullData[216^ /*line pznx6DQa.go:1*/ int(idxKey[4])], fullData[79^ /*line pznx6DQa.go:1*/ int(idxKey[1])]+fullData[19^ /*line pznx6DQa.go:1*/ int(idxKey[1])], fullData[109^ /*line pznx6DQa.go:1*/ int(idxKey[4])]^fullData[212^ /*line pznx6DQa.go:1*/ int(idxKey[4])], fullData[27^ /*line pznx6DQa.go:1*/ int(idxKey[2])]^fullData[28^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[102^ /*line pznx6DQa.go:1*/ int(idxKey[4])]+fullData[118^ /*line pznx6DQa.go:1*/ int(idxKey[4])], fullData[242^ /*line pznx6DQa.go:1*/ int(idxKey[7])]^fullData[74^ /*line pznx6DQa.go:1*/ int(idxKey[7])], fullData[116^ /*line pznx6DQa.go:1*/ int(idxKey[11])]+fullData[59^ /*line pznx6DQa.go:1*/ int(idxKey[11])], fullData[37^ /*line pznx6DQa.go:1*/ int(idxKey[0])]^fullData[5^ /*line pznx6DQa.go:1*/ int(idxKey[0])], fullData[248^ /*line pznx6DQa.go:1*/ int(idxKey[7])]-fullData[227^ /*line pznx6DQa.go:1*/ int(idxKey[7])], fullData[40^ /*line pznx6DQa.go:1*/ int(idxKey[14])]-fullData[98^ /*line pznx6DQa.go:1*/ int(idxKey[14])], fullData[103^ /*line pznx6DQa.go:1*/ int(idxKey[6])]-fullData[180^ /*line pznx6DQa.go:1*/ int(idxKey[6])], fullData[133^ /*line pznx6DQa.go:1*/ int(idxKey[12])]+fullData[85^ /*line pznx6DQa.go:1*/ int(idxKey[12])], fullData[82^ /*line pznx6DQa.go:1*/ int(idxKey[6])]^fullData[137^ /*line pznx6DQa.go:1*/ int(idxKey[6])], fullData[120^ /*line pznx6DQa.go:1*/ int(idxKey[12])]^fullData[112^ /*line pznx6DQa.go:1*/ int(idxKey[12])], fullData[166^ /*line pznx6DQa.go:1*/ int(idxKey[3])]+fullData[181^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[29^ /*line pznx6DQa.go:1*/ int(idxKey[5])]^fullData[40^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[33^ /*line pznx6DQa.go:1*/ int(idxKey[2])]^fullData[83^ /*line pznx6DQa.go:1*/ int(idxKey[2])], fullData[203^ /*line pznx6DQa.go:1*/ int(idxKey[3])]^fullData[189^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[248^ /*line pznx6DQa.go:1*/ int(idxKey[13])]^fullData[24^ /*line pznx6DQa.go:1*/ int(idxKey[13])], fullData[28^ /*line pznx6DQa.go:1*/ int(idxKey[11])]-fullData[114^ /*line pznx6DQa.go:1*/ int(idxKey[11])], fullData[194^ /*line pznx6DQa.go:1*/ int(idxKey[15])]^fullData[243^ /*line pznx6DQa.go:1*/ int(idxKey[15])], fullData[139^ /*line pznx6DQa.go:1*/ int(idxKey[5])]+fullData[186^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[23^ /*line pznx6DQa.go:1*/ int(idxKey[0])]^fullData[217^ /*line pznx6DQa.go:1*/ int(idxKey[0])], fullData[148^ /*line pznx6DQa.go:1*/ int(idxKey[5])]-fullData[44^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[118^ /*line pznx6DQa.go:1*/ int(idxKey[0])]+fullData[32^ /*line pznx6DQa.go:1*/ int(idxKey[0])], fullData[205^ /*line pznx6DQa.go:1*/ int(idxKey[3])]+fullData[155^ /*line pznx6DQa.go:1*/ int(idxKey[3])], fullData[67^ /*line pznx6DQa.go:1*/ int(idxKey[6])]-fullData[62^ /*line pznx6DQa.go:1*/ int(idxKey[6])], fullData[161^ /*line pznx6DQa.go:1*/ int(idxKey[10])]-fullData[144^ /*line pznx6DQa.go:1*/ int(idxKey[10])], fullData[53^ /*line pznx6DQa.go:1*/ int(idxKey[5])]+fullData[124^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[123^ /*line pznx6DQa.go:1*/ int(idxKey[12])]+fullData[14^ /*line pznx6DQa.go:1*/ int(idxKey[12])], fullData[42^ /*line pznx6DQa.go:1*/ int(idxKey[14])]^fullData[0^ /*line pznx6DQa.go:1*/ int(idxKey[14])], fullData[188^ /*line pznx6DQa.go:1*/ int(idxKey[5])]^fullData[32^ /*line pznx6DQa.go:1*/ int(idxKey[5])], fullData[165^ /*line pznx6DQa.go:1*/ int(idxKey[13])]^fullData[39^ /*line pznx6DQa.go:1*/ int(idxKey[13])])
			return /*line pznx6DQa.go:1*/ string(data)
		}(), /*line pgqCVqmcGnDy.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2, k1qj7vimKIy7)
		return
	}

	k1qj7vimKIy7 = /*line Wh0jLqLdkm.go:1*/ cu1RSg2YR.xAnHtUT9ATI(cWdSpmiBLOz.BlockHeight)
	if k1qj7vimKIy7 != nil {
		cu1RSg2YR.bufferedQCs[cWdSpmiBLOz.ZWsU_2] = cWdSpmiBLOz
		/*line ehg6nsAwVu.go:1*/ log.Debugf(func() /*line yZDELIVb3.go:1*/ string {
			data := [] /*line yZDELIVb3.go:1*/ byte("\v%@]\xd8\x1a\x8br\x17avE\x1e\xf8\x00\x00w\x7ff\x18p`\x00\\Vv-\x1c;y\x1d\x8fhu\xf2isclZOe\vH\x19d0U3r gotGrCadyT\no \x88e\x9fti\xcbp\xb1\xf1keLOm zx@\xc1%\xdf")
			positions := [...]byte{19, 70, 60, 22, 66, 64, 10, 6, 9, 16, 37, 75, 19, 47, 77, 84, 25, 75, 26, 29, 14, 43, 13, 33, 13, 19, 56, 20, 43, 25, 73, 73, 79, 72, 21, 81, 43, 54, 17, 46, 70, 0, 15, 68, 76, 64, 73, 8, 32, 4, 28, 12, 47, 23, 2, 11, 34, 71, 20, 79, 41, 79, 15, 51, 39, 47, 9, 28, 27, 31, 40, 46, 38, 54, 32, 82, 39, 8, 27, 79, 69, 33, 30, 44, 61, 12, 42, 48, 0, 10, 15, 21, 42, 5, 77, 0, 46, 56}
			for i := 0; i < 98; i += 2 {
				localKey := /*line yZDELIVb3.go:1*/ byte(i) + /*line yZDELIVb3.go:1*/ byte(positions[i]^positions[i+1]) + 244
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line yZDELIVb3.go:1*/ string(data)
		}(), /*line t24poxn21Myv.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2, k1qj7vimKIy7)
		return
	}

	/*line BiG8gk.go:1*/
	cu1RSg2YR.pm.AdvanceView(cWdSpmiBLOz.LwVL87)
	/*line I6BAZmlOD.go:1*/ cu1RSg2YR.aDTslULytPz_(cWdSpmiBLOz)

	if /*line FxN4Ps.go:1*/ cu1RSg2YR.pm.IsTimeToElect() {

		hYhbNJBbO := /*line fqqwbR1Mm5PK.go:1*/ cu1RSg2YR.pm.GetCurEpoch()
		ey93G03TMe := /*line jRILCdf.go:1*/ cu1RSg2YR.ElectCommittees(ouJpb2aF3mlv.BlockHash, hYhbNJBbO+1)
		if ey93G03TMe != nil {
			/*line Pvho5V5mhBV.go:1*/ log.Debugf(func() /*line OPMQm_.go:1*/ string {
				fullData := [] /*line OPMQm_.go:1*/ byte("\xc1\x13\t\xf7O8\xb1\x86y\xf1\xec\b\x1d\xab\xcf3\x87X\xafT\xd9\xd8.\xa1\x9c\xeb\x84\xde\x1b\x1a> 1\x05\xd8}\x85\xab\xd9!y\n\x1e\xa5\xdeF\x00T\xb7.p\bC\xb2t\x0f\xf5\xd0猢\xe1&\xd2D\\\xd5C\x9a\xc3\\\xac\xe4\x8cX\vd\xb2\xbd\x11\xfa;\xa8\x90Ŷ\xb2\xaa_q+\x93\x19(⹂xgN\xa4\x8d\x13\xb29\x95\xfe\x1b\xed\xcaX\x91e\x88u\x1e\xdbj\ue7e5\x19\x06E\x00\xc1\u05cb\xf8\x88Ge:NՕ\xbd\a\a\xb9")
				idxKey := [] /*line OPMQm_.go:1*/ byte("\xe6\xff\xb6f\x10\x97,\xb9\xb7\xcb\x0fDX\xc7\xc0%")
				data := /*line OPMQm_.go:1*/ make([]byte, 0, 71)
				data = /*line OPMQm_.go:1*/ append(data, fullData[74^ /*line OPMQm_.go:1*/ int(idxKey[11])]+fullData[13^ /*line OPMQm_.go:1*/ int(idxKey[11])], fullData[51^ /*line OPMQm_.go:1*/ int(idxKey[12])]+fullData[113^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[149^ /*line OPMQm_.go:1*/ int(idxKey[2])]-fullData[63^ /*line OPMQm_.go:1*/ int(idxKey[2])], fullData[162^ /*line OPMQm_.go:1*/ int(idxKey[5])]+fullData[178^ /*line OPMQm_.go:1*/ int(idxKey[5])], fullData[146^ /*line OPMQm_.go:1*/ int(idxKey[7])]^fullData[157^ /*line OPMQm_.go:1*/ int(idxKey[7])], fullData[122^ /*line OPMQm_.go:1*/ int(idxKey[15])]-fullData[74^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[59^ /*line OPMQm_.go:1*/ int(idxKey[3])]^fullData[119^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[74^ /*line OPMQm_.go:1*/ int(idxKey[3])]^fullData[33^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[147^ /*line OPMQm_.go:1*/ int(idxKey[13])]^fullData[144^ /*line OPMQm_.go:1*/ int(idxKey[13])], fullData[89^ /*line OPMQm_.go:1*/ int(idxKey[12])]^fullData[106^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[164^ /*line OPMQm_.go:1*/ int(idxKey[8])]+fullData[248^ /*line OPMQm_.go:1*/ int(idxKey[8])], fullData[140^ /*line OPMQm_.go:1*/ int(idxKey[14])]+fullData[247^ /*line OPMQm_.go:1*/ int(idxKey[14])], fullData[156^ /*line OPMQm_.go:1*/ int(idxKey[0])]^fullData[148^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[120^ /*line OPMQm_.go:1*/ int(idxKey[12])]-fullData[46^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[11^ /*line OPMQm_.go:1*/ int(idxKey[4])]+fullData[0^ /*line OPMQm_.go:1*/ int(idxKey[4])], fullData[214^ /*line OPMQm_.go:1*/ int(idxKey[0])]-fullData[157^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[224^ /*line OPMQm_.go:1*/ int(idxKey[5])]^fullData[142^ /*line OPMQm_.go:1*/ int(idxKey[5])], fullData[174^ /*line OPMQm_.go:1*/ int(idxKey[6])]^fullData[29^ /*line OPMQm_.go:1*/ int(idxKey[6])], fullData[12^ /*line OPMQm_.go:1*/ int(idxKey[11])]+fullData[36^ /*line OPMQm_.go:1*/ int(idxKey[11])], fullData[155^ /*line OPMQm_.go:1*/ int(idxKey[0])]+fullData[180^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[220^ /*line OPMQm_.go:1*/ int(idxKey[9])]-fullData[213^ /*line OPMQm_.go:1*/ int(idxKey[9])], fullData[12^ /*line OPMQm_.go:1*/ int(idxKey[4])]+fullData[61^ /*line OPMQm_.go:1*/ int(idxKey[4])], fullData[244^ /*line OPMQm_.go:1*/ int(idxKey[0])]-fullData[183^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[125^ /*line OPMQm_.go:1*/ int(idxKey[15])]-fullData[117^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[126^ /*line OPMQm_.go:1*/ int(idxKey[3])]^fullData[11^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[242^ /*line OPMQm_.go:1*/ int(idxKey[8])]-fullData[152^ /*line OPMQm_.go:1*/ int(idxKey[8])], fullData[10^ /*line OPMQm_.go:1*/ int(idxKey[3])]-fullData[78^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[249^ /*line OPMQm_.go:1*/ int(idxKey[14])]+fullData[71^ /*line OPMQm_.go:1*/ int(idxKey[14])], fullData[41^ /*line OPMQm_.go:1*/ int(idxKey[6])]+fullData[37^ /*line OPMQm_.go:1*/ int(idxKey[6])], fullData[255^ /*line OPMQm_.go:1*/ int(idxKey[1])]^fullData[194^ /*line OPMQm_.go:1*/ int(idxKey[1])], fullData[214^ /*line OPMQm_.go:1*/ int(idxKey[5])]-fullData[148^ /*line OPMQm_.go:1*/ int(idxKey[5])], fullData[208^ /*line OPMQm_.go:1*/ int(idxKey[0])]-fullData[213^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[95^ /*line OPMQm_.go:1*/ int(idxKey[6])]-fullData[167^ /*line OPMQm_.go:1*/ int(idxKey[6])], fullData[140^ /*line OPMQm_.go:1*/ int(idxKey[0])]+fullData[150^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[160^ /*line OPMQm_.go:1*/ int(idxKey[15])]+fullData[27^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[144^ /*line OPMQm_.go:1*/ int(idxKey[9])]+fullData[174^ /*line OPMQm_.go:1*/ int(idxKey[9])], fullData[36^ /*line OPMQm_.go:1*/ int(idxKey[3])]-fullData[4^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[221^ /*line OPMQm_.go:1*/ int(idxKey[0])]+fullData[192^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[100^ /*line OPMQm_.go:1*/ int(idxKey[12])]+fullData[222^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[66^ /*line OPMQm_.go:1*/ int(idxKey[6])]^fullData[77^ /*line OPMQm_.go:1*/ int(idxKey[6])], fullData[89^ /*line OPMQm_.go:1*/ int(idxKey[10])]-fullData[11^ /*line OPMQm_.go:1*/ int(idxKey[10])], fullData[173^ /*line OPMQm_.go:1*/ int(idxKey[6])]-fullData[85^ /*line OPMQm_.go:1*/ int(idxKey[6])], fullData[17^ /*line OPMQm_.go:1*/ int(idxKey[15])]^fullData[51^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[206^ /*line OPMQm_.go:1*/ int(idxKey[11])]-fullData[0^ /*line OPMQm_.go:1*/ int(idxKey[11])], fullData[24^ /*line OPMQm_.go:1*/ int(idxKey[12])]-fullData[44^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[163^ /*line OPMQm_.go:1*/ int(idxKey[1])]-fullData[135^ /*line OPMQm_.go:1*/ int(idxKey[1])], fullData[162^ /*line OPMQm_.go:1*/ int(idxKey[9])]-fullData[236^ /*line OPMQm_.go:1*/ int(idxKey[9])], fullData[236^ /*line OPMQm_.go:1*/ int(idxKey[0])]+fullData[238^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[162^ /*line OPMQm_.go:1*/ int(idxKey[8])]^fullData[63^ /*line OPMQm_.go:1*/ int(idxKey[8])], fullData[67^ /*line OPMQm_.go:1*/ int(idxKey[11])]-fullData[34^ /*line OPMQm_.go:1*/ int(idxKey[11])], fullData[102^ /*line OPMQm_.go:1*/ int(idxKey[11])]^fullData[196^ /*line OPMQm_.go:1*/ int(idxKey[11])], fullData[199^ /*line OPMQm_.go:1*/ int(idxKey[7])]-fullData[222^ /*line OPMQm_.go:1*/ int(idxKey[7])], fullData[60^ /*line OPMQm_.go:1*/ int(idxKey[12])]^fullData[103^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[235^ /*line OPMQm_.go:1*/ int(idxKey[0])]-fullData[153^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[21^ /*line OPMQm_.go:1*/ int(idxKey[10])]-fullData[37^ /*line OPMQm_.go:1*/ int(idxKey[10])], fullData[41^ /*line OPMQm_.go:1*/ int(idxKey[12])]+fullData[98^ /*line OPMQm_.go:1*/ int(idxKey[12])], fullData[78^ /*line OPMQm_.go:1*/ int(idxKey[4])]^fullData[67^ /*line OPMQm_.go:1*/ int(idxKey[4])], fullData[85^ /*line OPMQm_.go:1*/ int(idxKey[10])]+fullData[55^ /*line OPMQm_.go:1*/ int(idxKey[10])], fullData[168^ /*line OPMQm_.go:1*/ int(idxKey[0])]+fullData[224^ /*line OPMQm_.go:1*/ int(idxKey[0])], fullData[100^ /*line OPMQm_.go:1*/ int(idxKey[3])]+fullData[32^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[156^ /*line OPMQm_.go:1*/ int(idxKey[1])]^fullData[151^ /*line OPMQm_.go:1*/ int(idxKey[1])], fullData[33^ /*line OPMQm_.go:1*/ int(idxKey[10])]^fullData[16^ /*line OPMQm_.go:1*/ int(idxKey[10])], fullData[115^ /*line OPMQm_.go:1*/ int(idxKey[10])]+fullData[140^ /*line OPMQm_.go:1*/ int(idxKey[10])], fullData[188^ /*line OPMQm_.go:1*/ int(idxKey[1])]^fullData[240^ /*line OPMQm_.go:1*/ int(idxKey[1])], fullData[235^ /*line OPMQm_.go:1*/ int(idxKey[1])]-fullData[138^ /*line OPMQm_.go:1*/ int(idxKey[1])], fullData[45^ /*line OPMQm_.go:1*/ int(idxKey[3])]+fullData[44^ /*line OPMQm_.go:1*/ int(idxKey[3])], fullData[104^ /*line OPMQm_.go:1*/ int(idxKey[15])]+fullData[112^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[161^ /*line OPMQm_.go:1*/ int(idxKey[15])]-fullData[56^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[46^ /*line OPMQm_.go:1*/ int(idxKey[15])]+fullData[41^ /*line OPMQm_.go:1*/ int(idxKey[15])], fullData[63^ /*line OPMQm_.go:1*/ int(idxKey[3])]+fullData[71^ /*line OPMQm_.go:1*/ int(idxKey[3])])
				return /*line OPMQm_.go:1*/ string(data)
			}(), /*line GOj2XKnxFM.go:1*/ cu1RSg2YR.ID(), ey93G03TMe, hYhbNJBbO+1)
		}
	}

	/*line AavxawQgU.go:1*/
	cu1RSg2YR.SetState(types.READY)
	/*line vaurj4VmjyQA.go:1*/ log.Debugf(func() /*line nezqZvd.go:1*/ string {
		fullData := [] /*line nezqZvd.go:1*/ byte("\x86g\x02\x04ۘ\xc1ܩ\x9c4\xe3A\xf13\xdf\xf7@5\xf2\xc4\xc7S\xbe\\\xa1'\xe1cѱ\x81\xc8g\x04\xd9% I\xa9)wk\x1a;\x15\x15_\x12\xd3s\xceߞ:?\x90\x90\xfb\xae\xf4\xea\x1b⧅\b$\xa79*\xb5\xe4g\x9e\x15ۦ\xf12\xb2\x10\xe3\xc3\xe3\x93T\xe7\x10\x8e\x04\xe9H\xb0(\xc6\u0084Z\xee\x1d<\xce\xfa\x16\x05|s\xb5/\xe1f\xbc\xa7\x9fΒ\xbc\n\xa9\xff\x03\xa6]\x99\x03\u05ec\xfb\xa1\xa4.\xd7˜\xb5Ċ\x82\xa3\xa2\xca틯]\xc1\x89(\xc5\x1d0wp\x1a@\x9ej2\x92\xaa\xa0\xbb%\xb4\x1b\xb7\x157\xe3ux͆\x85H\xc6d\xed\xc1K\xa6ע\x0e0\xd1%q\xc51\xea\xf0\xee@\xf6\x8b!qV\xf7\xfb\xb4ܶ\xe1B\xcfs\xcfC\x86U\x82\xbb\x19")
		idxKey := [] /*line nezqZvd.go:1*/ byte("U\x02\xc0\xbdȭP\xbb\xb3\xe1\xea\bw\x84\xf5\xd6")
		data := /*line nezqZvd.go:1*/ make([]byte, 0, 109)
		data = /*line nezqZvd.go:1*/ append(data, fullData[212^ /*line nezqZvd.go:1*/ int(idxKey[13])]+fullData[140^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[41^ /*line nezqZvd.go:1*/ int(idxKey[0])]^fullData[32^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[16^ /*line nezqZvd.go:1*/ int(idxKey[6])]-fullData[238^ /*line nezqZvd.go:1*/ int(idxKey[6])], fullData[54^ /*line nezqZvd.go:1*/ int(idxKey[5])]^fullData[59^ /*line nezqZvd.go:1*/ int(idxKey[5])], fullData[151^ /*line nezqZvd.go:1*/ int(idxKey[3])]+fullData[209^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[123^ /*line nezqZvd.go:1*/ int(idxKey[1])]-fullData[78^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[200^ /*line nezqZvd.go:1*/ int(idxKey[12])]+fullData[164^ /*line nezqZvd.go:1*/ int(idxKey[12])], fullData[51^ /*line nezqZvd.go:1*/ int(idxKey[12])]-fullData[101^ /*line nezqZvd.go:1*/ int(idxKey[12])], fullData[140^ /*line nezqZvd.go:1*/ int(idxKey[10])]^fullData[243^ /*line nezqZvd.go:1*/ int(idxKey[10])], fullData[47^ /*line nezqZvd.go:1*/ int(idxKey[12])]+fullData[97^ /*line nezqZvd.go:1*/ int(idxKey[12])], fullData[7^ /*line nezqZvd.go:1*/ int(idxKey[0])]^fullData[248^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[228^ /*line nezqZvd.go:1*/ int(idxKey[0])]-fullData[88^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[31^ /*line nezqZvd.go:1*/ int(idxKey[8])]^fullData[164^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[170^ /*line nezqZvd.go:1*/ int(idxKey[9])]^fullData[38^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[134^ /*line nezqZvd.go:1*/ int(idxKey[7])]-fullData[21^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[216^ /*line nezqZvd.go:1*/ int(idxKey[7])]^fullData[178^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[148^ /*line nezqZvd.go:1*/ int(idxKey[14])]-fullData[164^ /*line nezqZvd.go:1*/ int(idxKey[14])], fullData[161^ /*line nezqZvd.go:1*/ int(idxKey[4])]-fullData[78^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[8^ /*line nezqZvd.go:1*/ int(idxKey[0])]+fullData[153^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[149^ /*line nezqZvd.go:1*/ int(idxKey[13])]-fullData[50^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[192^ /*line nezqZvd.go:1*/ int(idxKey[7])]-fullData[220^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[118^ /*line nezqZvd.go:1*/ int(idxKey[8])]^fullData[113^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[176^ /*line nezqZvd.go:1*/ int(idxKey[4])]+fullData[98^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[1^ /*line nezqZvd.go:1*/ int(idxKey[2])]-fullData[83^ /*line nezqZvd.go:1*/ int(idxKey[2])], fullData[63^ /*line nezqZvd.go:1*/ int(idxKey[3])]^fullData[174^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[169^ /*line nezqZvd.go:1*/ int(idxKey[9])]^fullData[37^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[86^ /*line nezqZvd.go:1*/ int(idxKey[13])]-fullData[85^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[54^ /*line nezqZvd.go:1*/ int(idxKey[1])]+fullData[2^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[50^ /*line nezqZvd.go:1*/ int(idxKey[1])]^fullData[46^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[235^ /*line nezqZvd.go:1*/ int(idxKey[6])]+fullData[153^ /*line nezqZvd.go:1*/ int(idxKey[6])], fullData[184^ /*line nezqZvd.go:1*/ int(idxKey[15])]+fullData[151^ /*line nezqZvd.go:1*/ int(idxKey[15])], fullData[211^ /*line nezqZvd.go:1*/ int(idxKey[12])]+fullData[240^ /*line nezqZvd.go:1*/ int(idxKey[12])], fullData[244^ /*line nezqZvd.go:1*/ int(idxKey[15])]+fullData[75^ /*line nezqZvd.go:1*/ int(idxKey[15])], fullData[96^ /*line nezqZvd.go:1*/ int(idxKey[4])]-fullData[187^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[212^ /*line nezqZvd.go:1*/ int(idxKey[9])]^fullData[111^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[198^ /*line nezqZvd.go:1*/ int(idxKey[5])]^fullData[8^ /*line nezqZvd.go:1*/ int(idxKey[5])], fullData[8^ /*line nezqZvd.go:1*/ int(idxKey[1])]-fullData[205^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[150^ /*line nezqZvd.go:1*/ int(idxKey[6])]^fullData[126^ /*line nezqZvd.go:1*/ int(idxKey[6])], fullData[131^ /*line nezqZvd.go:1*/ int(idxKey[13])]-fullData[244^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[57^ /*line nezqZvd.go:1*/ int(idxKey[11])]-fullData[20^ /*line nezqZvd.go:1*/ int(idxKey[11])], fullData[132^ /*line nezqZvd.go:1*/ int(idxKey[7])]^fullData[131^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[76^ /*line nezqZvd.go:1*/ int(idxKey[10])]-fullData[182^ /*line nezqZvd.go:1*/ int(idxKey[10])], fullData[20^ /*line nezqZvd.go:1*/ int(idxKey[13])]+fullData[78^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[16^ /*line nezqZvd.go:1*/ int(idxKey[11])]^fullData[77^ /*line nezqZvd.go:1*/ int(idxKey[11])], fullData[135^ /*line nezqZvd.go:1*/ int(idxKey[6])]-fullData[29^ /*line nezqZvd.go:1*/ int(idxKey[6])], fullData[132^ /*line nezqZvd.go:1*/ int(idxKey[3])]+fullData[20^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[3^ /*line nezqZvd.go:1*/ int(idxKey[0])]+fullData[120^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[155^ /*line nezqZvd.go:1*/ int(idxKey[1])]-fullData[0^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[93^ /*line nezqZvd.go:1*/ int(idxKey[11])]^fullData[52^ /*line nezqZvd.go:1*/ int(idxKey[11])], fullData[90^ /*line nezqZvd.go:1*/ int(idxKey[4])]^fullData[5^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[225^ /*line nezqZvd.go:1*/ int(idxKey[2])]-fullData[3^ /*line nezqZvd.go:1*/ int(idxKey[2])], fullData[79^ /*line nezqZvd.go:1*/ int(idxKey[2])]-fullData[168^ /*line nezqZvd.go:1*/ int(idxKey[2])], fullData[37^ /*line nezqZvd.go:1*/ int(idxKey[7])]^fullData[42^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[218^ /*line nezqZvd.go:1*/ int(idxKey[6])]+fullData[144^ /*line nezqZvd.go:1*/ int(idxKey[6])], fullData[99^ /*line nezqZvd.go:1*/ int(idxKey[12])]^fullData[105^ /*line nezqZvd.go:1*/ int(idxKey[12])], fullData[154^ /*line nezqZvd.go:1*/ int(idxKey[1])]^fullData[41^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[200^ /*line nezqZvd.go:1*/ int(idxKey[9])]+fullData[198^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[194^ /*line nezqZvd.go:1*/ int(idxKey[10])]+fullData[220^ /*line nezqZvd.go:1*/ int(idxKey[10])], fullData[179^ /*line nezqZvd.go:1*/ int(idxKey[3])]-fullData[142^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[224^ /*line nezqZvd.go:1*/ int(idxKey[9])]^fullData[70^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[150^ /*line nezqZvd.go:1*/ int(idxKey[5])]+fullData[242^ /*line nezqZvd.go:1*/ int(idxKey[5])], fullData[249^ /*line nezqZvd.go:1*/ int(idxKey[13])]+fullData[235^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[154^ /*line nezqZvd.go:1*/ int(idxKey[2])]-fullData[92^ /*line nezqZvd.go:1*/ int(idxKey[2])], fullData[193^ /*line nezqZvd.go:1*/ int(idxKey[8])]+fullData[62^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[200^ /*line nezqZvd.go:1*/ int(idxKey[5])]-fullData[142^ /*line nezqZvd.go:1*/ int(idxKey[5])], fullData[205^ /*line nezqZvd.go:1*/ int(idxKey[7])]-fullData[204^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[75^ /*line nezqZvd.go:1*/ int(idxKey[4])]^fullData[170^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[68^ /*line nezqZvd.go:1*/ int(idxKey[1])]-fullData[4^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[96^ /*line nezqZvd.go:1*/ int(idxKey[9])]-fullData[174^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[160^ /*line nezqZvd.go:1*/ int(idxKey[10])]-fullData[83^ /*line nezqZvd.go:1*/ int(idxKey[10])], fullData[81^ /*line nezqZvd.go:1*/ int(idxKey[0])]-fullData[247^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[139^ /*line nezqZvd.go:1*/ int(idxKey[1])]-fullData[92^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[240^ /*line nezqZvd.go:1*/ int(idxKey[6])]+fullData[48^ /*line nezqZvd.go:1*/ int(idxKey[6])], fullData[205^ /*line nezqZvd.go:1*/ int(idxKey[8])]+fullData[182^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[57^ /*line nezqZvd.go:1*/ int(idxKey[13])]^fullData[49^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[194^ /*line nezqZvd.go:1*/ int(idxKey[3])]-fullData[177^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[112^ /*line nezqZvd.go:1*/ int(idxKey[0])]+fullData[193^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[26^ /*line nezqZvd.go:1*/ int(idxKey[7])]+fullData[46^ /*line nezqZvd.go:1*/ int(idxKey[7])], fullData[168^ /*line nezqZvd.go:1*/ int(idxKey[9])]^fullData[89^ /*line nezqZvd.go:1*/ int(idxKey[9])], fullData[115^ /*line nezqZvd.go:1*/ int(idxKey[3])]^fullData[153^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[229^ /*line nezqZvd.go:1*/ int(idxKey[0])]+fullData[226^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[72^ /*line nezqZvd.go:1*/ int(idxKey[0])]+fullData[222^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[117^ /*line nezqZvd.go:1*/ int(idxKey[3])]^fullData[57^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[88^ /*line nezqZvd.go:1*/ int(idxKey[10])]^fullData[202^ /*line nezqZvd.go:1*/ int(idxKey[10])], fullData[245^ /*line nezqZvd.go:1*/ int(idxKey[13])]^fullData[62^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[134^ /*line nezqZvd.go:1*/ int(idxKey[4])]+fullData[165^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[255^ /*line nezqZvd.go:1*/ int(idxKey[3])]-fullData[34^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[190^ /*line nezqZvd.go:1*/ int(idxKey[13])]-fullData[240^ /*line nezqZvd.go:1*/ int(idxKey[13])], fullData[22^ /*line nezqZvd.go:1*/ int(idxKey[3])]^fullData[217^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[143^ /*line nezqZvd.go:1*/ int(idxKey[15])]+fullData[141^ /*line nezqZvd.go:1*/ int(idxKey[15])], fullData[159^ /*line nezqZvd.go:1*/ int(idxKey[4])]-fullData[221^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[41^ /*line nezqZvd.go:1*/ int(idxKey[8])]^fullData[132^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[103^ /*line nezqZvd.go:1*/ int(idxKey[8])]-fullData[188^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[107^ /*line nezqZvd.go:1*/ int(idxKey[4])]+fullData[72^ /*line nezqZvd.go:1*/ int(idxKey[4])], fullData[241^ /*line nezqZvd.go:1*/ int(idxKey[10])]-fullData[128^ /*line nezqZvd.go:1*/ int(idxKey[10])], fullData[231^ /*line nezqZvd.go:1*/ int(idxKey[8])]-fullData[99^ /*line nezqZvd.go:1*/ int(idxKey[8])], fullData[65^ /*line nezqZvd.go:1*/ int(idxKey[14])]-fullData[62^ /*line nezqZvd.go:1*/ int(idxKey[14])], fullData[239^ /*line nezqZvd.go:1*/ int(idxKey[14])]-fullData[125^ /*line nezqZvd.go:1*/ int(idxKey[14])], fullData[143^ /*line nezqZvd.go:1*/ int(idxKey[3])]^fullData[131^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[70^ /*line nezqZvd.go:1*/ int(idxKey[14])]+fullData[218^ /*line nezqZvd.go:1*/ int(idxKey[14])], fullData[46^ /*line nezqZvd.go:1*/ int(idxKey[11])]-fullData[75^ /*line nezqZvd.go:1*/ int(idxKey[11])], fullData[208^ /*line nezqZvd.go:1*/ int(idxKey[2])]-fullData[223^ /*line nezqZvd.go:1*/ int(idxKey[2])], fullData[94^ /*line nezqZvd.go:1*/ int(idxKey[0])]^fullData[6^ /*line nezqZvd.go:1*/ int(idxKey[0])], fullData[190^ /*line nezqZvd.go:1*/ int(idxKey[3])]-fullData[107^ /*line nezqZvd.go:1*/ int(idxKey[3])], fullData[79^ /*line nezqZvd.go:1*/ int(idxKey[11])]-fullData[180^ /*line nezqZvd.go:1*/ int(idxKey[11])], fullData[142^ /*line nezqZvd.go:1*/ int(idxKey[1])]^fullData[215^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[135^ /*line nezqZvd.go:1*/ int(idxKey[1])]-fullData[120^ /*line nezqZvd.go:1*/ int(idxKey[1])], fullData[98^ /*line nezqZvd.go:1*/ int(idxKey[14])]+fullData[90^ /*line nezqZvd.go:1*/ int(idxKey[14])])
		return /*line nezqZvd.go:1*/ string(data)
	}(), /*line TXYTTEeiFgkO.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.BlockHeight, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.ZWsU_2)

	hYhbNJBbO, uxNBF8Z7 := /*line czXowiML9rJ.go:1*/ cu1RSg2YR.pm.GetCurEpochView()

	for pHERGPksgf30, d1zM2qw := range cu1RSg2YR.bufferedBlocks {
		for a3qJWd, u4oFYmVjv := range d1zM2qw {
			for gNdxFNoW, ouJpb2aF3mlv := range u4oFYmVjv {
				/*line sSk7gOkdXY.go:1*/ log.Debugf(func() /*line tmqykJ8.go:1*/ string {
					seed := /*line tmqykJ8.go:1*/ byte(169)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line tmqykJ8.go:1*/ append(data, x-seed); seed += x; return fnc }
					/*line tmqykJ8.go:1*/ fnc(4)(210)(245)(209)(101)(210)(236)(218)(177)(86)(174)(106)(212)(120)(18)(49)(100)(189)(119)(241)(220)(182)(127)(239)(207)(183)(115)(215)(114)(219)(248)(3)(247)(238)(219)(195)(121)(241)(158)(126)(6)(15)(18)(44)(13)(95)(201)(145)(22)(49)(26)(57)(195)(48)(182)(95)(186)(134)(181)(111)(47)(8)(88)(173)(94)(186)(117)(246)(152)(53)(187)(32)(105)(205)(118)(241)(53)
					return /*line tmqykJ8.go:1*/ string(data)
				}(), /*line CRTKSTon.go:1*/ cu1RSg2YR.ID(), pHERGPksgf30, a3qJWd, gNdxFNoW, ouJpb2aF3mlv.WorkerBlock.BlockHash)
			}
		}
	}

	if ouJpb2aF3mlv, i3KmQgjb := cu1RSg2YR.bufferedBlocks[hYhbNJBbO][uxNBF8Z7][cWdSpmiBLOz.BlockHeight]; i3KmQgjb {
		/*line f90aFFua6B.go:1*/ log.Debugf(func() /*line SenZ7nUQbZ.go:1*/ string {
			data := [] /*line SenZ7nUQbZ.go:1*/ byte("\xbb\xc3z\f \"\xa5\xbao\xa5©\x88\xdd\x06r|\rfʹat\x8aXoGz)\x90s\xada\xb9ƺ\x8bons\xfbns\xee9F\xc6\x03&\x9f\r¶\xf0|-~sq\xac\x80oSk }m\x12\x8a\xb5\x8d\x12\x9do\xb4\x9d\xb7ei\xf3h\xb4 \x93\xa4wwi\xf0wn%v9][\x842h\x97\xb7\xd4 ğ'r\x10")
			positions := [...]byte{13, 33, 99, 55, 99, 58, 3, 56, 50, 47, 1, 60, 70, 24, 14, 103, 31, 90, 59, 35, 79, 72, 100, 93, 88, 58, 81, 107, 0, 74, 17, 106, 52, 11, 2, 76, 6, 56, 75, 9, 40, 104, 96, 63, 34, 10, 44, 99, 86, 54, 24, 51, 18, 58, 57, 94, 48, 85, 57, 74, 6, 49, 68, 103, 46, 65, 104, 88, 94, 24, 90, 43, 55, 67, 88, 107, 85, 51, 46, 16, 51, 95, 69, 84, 29, 44, 47, 14, 0, 68, 3, 83, 0, 23, 53, 93, 45, 49, 50, 18, 100, 29, 97, 31, 35, 59, 105, 19, 36, 53, 60, 27, 93, 99, 62, 107, 10, 47, 59, 71, 20, 88, 55, 62, 57, 57, 62, 79, 88, 66, 2, 12, 81, 45, 95, 26, 96, 79, 33, 60, 3, 27, 29, 65, 29, 66, 27, 106, 90, 106, 5, 101, 86, 65, 7, 105}
			for i := 0; i < 156; i += 2 {
				localKey := /*line SenZ7nUQbZ.go:1*/ byte(i) + /*line SenZ7nUQbZ.go:1*/ byte(positions[i]^positions[i+1]) + 94
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line SenZ7nUQbZ.go:1*/ string(data)
		}(), /*line GoppWs.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.BlockHeight, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.ZWsU_2)
		_ = /*line IVYdw9HAi.go:1*/ cu1RSg2YR.ProcessWorkerBlock(ouJpb2aF3mlv)
		/*line QrEP0YN1S.go:1*/ delete(cu1RSg2YR.bufferedBlocks[hYhbNJBbO][uxNBF8Z7], cWdSpmiBLOz.BlockHeight)
	}

	if dWmn_Gtk1d, i3KmQgjb := cu1RSg2YR.bufferedCQCs[cWdSpmiBLOz.ZWsU_2]; i3KmQgjb {
		/*line xtn3XNnjcw.go:1*/ cu1RSg2YR.emCoevqAq(dWmn_Gtk1d)
		/*line zPSw209g.go:1*/ delete(cu1RSg2YR.bufferedCQCs, cWdSpmiBLOz.ZWsU_2)
	}

	gFIlzELaH := /*line Q2UErDVX3_Y.go:1*/ quorum.G71jC_Q[quorum.Commit](cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.BlockHeight /*line K5Jo0tDEb.go:1*/, cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2)
	/*line QIoNH1lN7P.go:1*/ cu1RSg2YR.BroadcastToSome( /*line MG6MXZWa.go:1*/ cu1RSg2YR.FindCommitteesFor(gFIlzELaH.Epoch), gFIlzELaH)

	/*line QMQ9LRLwub.go:1*/
	cu1RSg2YR.ProcessCommit(gFIlzELaH)
}

func (cu1RSg2YR *PBFT) emCoevqAq(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) {
	/*line KFkWco.go:1*/ log.Debugf(func() /*line tm8wx8.go:1*/ string {
		data := /*line tm8wx8.go:1*/ make([]byte, 0, 106)
		i := 18
		decryptKey := 115
		for counter := 0; i != 14; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 42:
				data = /*line tm8wx8.go:1*/ append(data, "\x87\x8b\xec\xe8"...,
				)
				i = 43
			case 37:
				i = 7
				data = /*line tm8wx8.go:1*/ append(data, "\xdf\xcd\xc0"...,
				)
			case 3:
				i = 21
				data = /*line tm8wx8.go:1*/ append(data, "\xe1\xec"...,
				)
			case 38:
				i = 3
				data = /*line tm8wx8.go:1*/ append(data, 173)
			case 9:
				i = 30
				data = /*line tm8wx8.go:1*/ append(data, 239)
			case 13:
				data = /*line tm8wx8.go:1*/ append(data, 143)
				i = 23
			case 8:
				data = /*line tm8wx8.go:1*/ append(data, "\x85\xd0"...,
				)
				i = 33
			case 46:
				data = /*line tm8wx8.go:1*/ append(data, 210)
				i = 0
			case 34:
				i = 35
				data = /*line tm8wx8.go:1*/ append(data, "\x85\x88\x81\x8f"...,
				)
			case 12:
				data = /*line tm8wx8.go:1*/ append(data, "\xda\xe9\xee\xfb"...,
				)
				i = 36
			case 6:
				data = /*line tm8wx8.go:1*/ append(data, "\x8d\xd8"...,
				)
				i = 13
			case 29:
				data = /*line tm8wx8.go:1*/ append(data, 142)
				i = 28
			case 40:
				data = /*line tm8wx8.go:1*/ append(data, "\xf9\xf7\xff"...,
				)
				i = 20
			case 19:
				i = 15
				data = /*line tm8wx8.go:1*/ append(data, "\x86Ɍ\x83"...,
				)
			case 1:
				i = 8
				data = /*line tm8wx8.go:1*/ append(data, "\x9a\x8b\xdd\xd7"...,
				)
			case 24:
				data = /*line tm8wx8.go:1*/ append(data, "\x83\x84"...,
				)
				i = 19
			case 2:
				data = /*line tm8wx8.go:1*/ append(data, "\x9d\xea\xee"...,
				)
				i = 22
			case 44:
				data = /*line tm8wx8.go:1*/ append(data, "\xd2\xde\xd0\xdd"...,
				)
				i = 32
			case 23:
				data = /*line tm8wx8.go:1*/ append(data, 151)
				i = 1
			case 11:
				i = 25
				data = /*line tm8wx8.go:1*/ append(data, 228)
			case 36:
				data = /*line tm8wx8.go:1*/ append(data, "\xfa\xe7\xe1\xeb"...,
				)
				i = 38
			case 0:
				i = 11
				data = /*line tm8wx8.go:1*/ append(data, 213)
			case 43:
				data = /*line tm8wx8.go:1*/ append(data, 186)
				i = 26
			case 35:
				i = 10
				data = /*line tm8wx8.go:1*/ append(data, "\x90\xc5"...,
				)
			case 4:
				data = /*line tm8wx8.go:1*/ append(data, 206)
				i = 44
			case 31:
				data = /*line tm8wx8.go:1*/ append(data, "\xc8\xd8\xf1\xc2"...,
				)
				i = 27
			case 15:
				i = 41
				data = /*line tm8wx8.go:1*/ append(data, "\x83\x8e"...,
				)
			case 10:
				i = 6
				data = /*line tm8wx8.go:1*/ append(data, 223)
			case 27:
				data = /*line tm8wx8.go:1*/ append(data, "Ә"...,
				)
				i = 39
			case 20:
				i = 24
				data = /*line tm8wx8.go:1*/ append(data, "\xf4\xf5\xe1"...,
				)
			case 33:
				data = /*line tm8wx8.go:1*/ append(data, "\x94\x86\x98\x97"...,
				)
				i = 2
			case 26:
				i = 14
				for y := range data {
					data[y] = data[y] ^ /*line tm8wx8.go:1*/ byte(decryptKey^y)
				}
			case 5:
				i = 40
				data = /*line tm8wx8.go:1*/ append(data, "\xf8\xe0\xe7"...,
				)
			case 32:
				data = /*line tm8wx8.go:1*/ append(data, 222)
				i = 31
			case 22:
				i = 42
				data = /*line tm8wx8.go:1*/ append(data, "\xbe\xe9"...,
				)
			case 30:
				data = /*line tm8wx8.go:1*/ append(data, "\xf3\xa4\xf4\xef"...,
				)
				i = 17
			case 45:
				i = 12
				data = /*line tm8wx8.go:1*/ append(data, 198)
			case 25:
				i = 4
				data = /*line tm8wx8.go:1*/ append(data, "\xc1\xd7"...,
				)
			case 39:
				data = /*line tm8wx8.go:1*/ append(data, "\x96\xc7"...,
				)
				i = 45
			case 41:
				i = 34
				data = /*line tm8wx8.go:1*/ append(data, "\x89\xab"...,
				)
			case 7:
				i = 46
				data = /*line tm8wx8.go:1*/ append(data, 197)
			case 28:
				i = 37
				data = /*line tm8wx8.go:1*/ append(data, "\x87\xdc"...,
				)
			case 21:
				i = 9
				data = /*line tm8wx8.go:1*/ append(data, "\xed\xec"...,
				)
			case 18:
				data = /*line tm8wx8.go:1*/ append(data, "\xf1\x8e\xde\xf4"...,
				)
				i = 29
			case 17:
				data = /*line tm8wx8.go:1*/ append(data, "\xf4\xea\xec"...,
				)
				i = 16
			case 16:
				i = 5
				data = /*line tm8wx8.go:1*/ append(data, "\xf3\xbf\xff"...,
				)
			}
		}
		return /*line tm8wx8.go:1*/ string(data)
	}(), /*line oVGhvhg7.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)

	if dWmn_Gtk1d.BlockHeight < /*line PvniHQ39Fm0g.go:1*/ cu1RSg2YR.GetLastBlockHeight() {
		return
	}

	/*line I5MA5t.go:1*/
	cu1RSg2YR.z6VaPTrjpK(dWmn_Gtk1d)

	i8LOoD6rG, _ := /*line AzfmykJEtt.go:1*/ cu1RSg2YR.bc.GetBlockByID(dWmn_Gtk1d.ZWsU_2)
	aVkBDTlb_E := i8LOoD6rG.(*blockchain.WorkerBlock)

	aVkBDTlb_E.Jy8nZ9h7dU = dWmn_Gtk1d

	if aVkBDTlb_E.Proposer == /*line R_mJbni.go:1*/ cu1RSg2YR.ID() {
		bZwrdESj := &blockchain.Accept{CommittedBlock: aVkBDTlb_E, Kg1tLS2Ii: cu1RSg2YR.gSnapshot, K0CmW5aPRDdY: cu1RSg2YR.gCode, JTK381N2H: /*line fgrnhP9p.go:1*/ time.Now()}
		dwFbQFBes := /*line LRBy48nWME.go:1*/ cu1RSg2YR.FindValidatorsFor(aVkBDTlb_E.BlockHeader.Epoch)
		/*line EBstv3.go:1*/ cu1RSg2YR.BroadcastToSome(dwFbQFBes, bZwrdESj)
	}

	k1qj7vimKIy7 := /*line f5Y_Jk.go:1*/ cu1RSg2YR.fhcoYWd(dWmn_Gtk1d)
	if k1qj7vimKIy7 != nil {
		/*line plgC3zcP1.go:1*/ log.Jp980o6YjM(func() /*line Ft3_q1zgMB2i.go:1*/ string {
			fullData := [] /*line Ft3_q1zgMB2i.go:1*/ byte("\xee\xf7S\x87nH\xb1\x03\xce&\x88\nŵ\xb5\f\x90\x96j\xecP\f\x16D:h\x96\x84-N\x92\x95T\x96\xc4w\xaf\xe1ԥ\xe9\xed\xceK\xd0\xd5\xfc\xba\xbd\xa3\xd1N\x8cI\x12.\x1bm~\xf1\xf9\x1d\xa2Wj>\x13\xbe\x8a*\x99\xda\xdadF\xb6\x8a:\xc7\x0e\x8by\xe1L\x82\xf7\xe7\x17\xd4;Fٲ\b\xea;\xae֏<\x85\xbf\x16=\xe3\xd11.\xf2\x93\xee\xda\xc6\xe1Z=\x14\xd0\xd101\x99")
			idxKey := [] /*line Ft3_q1zgMB2i.go:1*/ byte("\xcaU\x93\x93?d\x1f\xc2\x06\x8f\x9e")
			data := /*line Ft3_q1zgMB2i.go:1*/ make([]byte, 0, 62)
			data = /*line Ft3_q1zgMB2i.go:1*/ append(data, fullData[230^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]^fullData[195^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[11^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]-fullData[89^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[107^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]^fullData[115^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[171^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])]+fullData[201^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])], fullData[184^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]-fullData[192^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[1^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])]^fullData[48^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])], fullData[102^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])]^fullData[20^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])], fullData[212^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])]^fullData[139^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])], fullData[200^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])]-fullData[211^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])], fullData[106^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]-fullData[100^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[178^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])]-fullData[249^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])], fullData[211^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[10])]-fullData[208^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[10])], fullData[199^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])]-fullData[239^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])], fullData[111^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]^fullData[50^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[50^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]^fullData[46^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[114^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])]^fullData[77^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])], fullData[121^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])]-fullData[11^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])], fullData[119^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]+fullData[12^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[110^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]^fullData[98^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[179^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])]^fullData[224^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])], fullData[114^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]-fullData[0^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[62^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])]-fullData[37^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])], fullData[176^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])]^fullData[148^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])], fullData[228^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])]^fullData[140^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])], fullData[102^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]^fullData[47^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[215^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]^fullData[168^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[14^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]+fullData[25^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[19^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])]+fullData[59^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[8])], fullData[157^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]+fullData[196^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[36^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])]-fullData[79^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])], fullData[220^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[10])]-fullData[194^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[10])], fullData[222^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])]-fullData[156^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])], fullData[225^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]^fullData[219^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[88^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]^fullData[121^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[18^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])]-fullData[93^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])], fullData[196^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])]-fullData[198^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])], fullData[210^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])]-fullData[222^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])], fullData[14^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])]^fullData[35^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])], fullData[92^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])]+fullData[71^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])], fullData[37^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])]-fullData[86^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])], fullData[255^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]^fullData[171^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[210^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])]^fullData[249^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])], fullData[85^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])]-fullData[22^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])], fullData[6^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]-fullData[94^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[122^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])]-fullData[124^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])], fullData[212^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])]^fullData[157^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])], fullData[4^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]-fullData[51^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[134^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])]^fullData[231^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])], fullData[225^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])]+fullData[147^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[3])], fullData[185^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]+fullData[141^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[17^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])]+fullData[29^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[5])], fullData[175^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[10])]^fullData[188^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[10])], fullData[70^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])]^fullData[78^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[1])], fullData[40^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])]^fullData[72^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])], fullData[171^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])]^fullData[204^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])], fullData[224^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])]+fullData[245^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[0])], fullData[167^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])]-fullData[247^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[7])], fullData[248^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])]+fullData[156^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[2])], fullData[119^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])]+fullData[101^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[4])], fullData[150^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])]+fullData[191^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[9])], fullData[66^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])]+fullData[27^ /*line Ft3_q1zgMB2i.go:1*/ int(idxKey[6])])
			return /*line Ft3_q1zgMB2i.go:1*/ string(data)
		}(), /*line CXd8VM.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, k1qj7vimKIy7)
		return
	}

	if aVkBDTlb_E.Proposer == /*line QCgWXhWE.go:1*/ cu1RSg2YR.ID() {
		for _, bQZ9a4xL := range aVkBDTlb_E.BlockData.CJpmTwa4y {
			if bQZ9a4xL.IsCrossShardTx {
				bQZ9a4xL.LatencyDissection.WorkerConsensusTime2 = /*line WrpD0nu.go:1*/ time.Now().UnixMilli()
			} else {
				bQZ9a4xL.LatencyDissection.WorkerConsensusTime = /*line HglFai98.go:1*/ time.Now().UnixMilli()
			}
		}
		hl5hslyPV := /*line d6ocvZrf.go:1*/ cu1RSg2YR.pbftMeasure.CalculateMeasurements(aVkBDTlb_E /*line QgnxcXtsqg.go:1*/, cu1RSg2YR.Shard())
		/*line eB8APAV.go:1*/ cu1RSg2YR.SendToBlockBuilder(&hl5hslyPV)
	}

	if aVkBDTlb_E.Proposer == /*line FNswXgB.go:1*/ cu1RSg2YR.ID() {
		/*line Xm0WKCR3D.go:1*/ cu1RSg2YR.SendToBlockBuilder(aVkBDTlb_E)
		/*line ad41hmkvhKS.go:1*/ log.Debugf(func() /*line _N9Bp2l.go:1*/ string {
			seed := /*line _N9Bp2l.go:1*/ byte(0)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line _N9Bp2l.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line _N9Bp2l.go:1*/ fnc(91)(202)(81)(231)(195)(8)(72)(2)(253)(244)(2)(14)(0)(208)(34)(13)(2)(245)(253)(3)(250)(254)(19)(241)(222)(46)(242)(198)(247)(51)(18)(9)(246)(188)(34)(42)(3)(244)(8)(181)(67)(12)(254)(3)(252)(249)(15)(241)(255)(188)(67)(12)(255)(5)(242)(9)(5)(2)(254)(173)(52)(27)(177)(34)(42)(3)(244)(8)(215)(51)(244)(3)(248)(1)(13)
			return /*line _N9Bp2l.go:1*/ string(data)
		}(), /*line oib7WYsBZr.go:1*/ cu1RSg2YR.ID())
	}
	/*line FlAXWSIbYixu.go:1*/ log.Debugf(func() /*line DMFUcA3x.go:1*/ string {
		data := [] /*line DMFUcA3x.go:1*/ byte("\n6-]C\x8f*r$\x8fe}I\xe9er\x17\xef\xdfKc\x1e/_4+\a$\xc6f\x93nizyTd\xfd\x11%\x84We\fbi\xe4\xef c/\xacmcm q\x05 r/\x05 `ebTi\xad`;2t'5\x06 \xbf1\x053k\x1903gr.QAv\x1cv|\x1f~ %v\x154c)g\x1f %\x90 I\x1e %\x9e")
		positions := [...]byte{88, 53, 88, 6, 30, 9, 66, 40, 68, 66, 58, 73, 23, 65, 17, 18, 104, 69, 75, 103, 89, 26, 51, 70, 86, 50, 79, 33, 113, 103, 99, 22, 63, 22, 21, 44, 69, 27, 47, 54, 73, 70, 61, 84, 95, 74, 43, 21, 19, 99, 102, 34, 23, 80, 93, 79, 70, 5, 12, 41, 11, 12, 102, 39, 100, 0, 54, 73, 95, 101, 25, 87, 88, 57, 83, 18, 23, 35, 110, 60, 63, 91, 74, 74, 17, 80, 1, 33, 16, 4, 24, 9, 2, 35, 1, 84, 27, 19, 50, 21, 88, 82, 94, 23, 8, 44, 95, 78, 37, 77, 28, 107, 33, 5, 113, 46, 70, 13, 93, 71, 38, 9}
		for i := 0; i < 122; i += 2 {
			localKey := /*line DMFUcA3x.go:1*/ byte(i) + /*line DMFUcA3x.go:1*/ byte(positions[i]^positions[i+1]) + 203
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line DMFUcA3x.go:1*/ string(data)
	}(), /*line sUjukIDxbfD.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)
}

func (cu1RSg2YR *PBFT) fhcoYWd(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) error {
	/*line iyTT3g.go:1*/ log.Debugf(func() /*line DhzIAv.go:1*/ string {
		data := /*line DhzIAv.go:1*/ make([]byte, 0, 74)
		i := 19
		decryptKey := 196
		for counter := 0; i != 13; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 20:
				i = 13
				for y := range data {
					data[y] = data[y] ^ /*line DhzIAv.go:1*/ byte(decryptKey^y)
				}
			case 23:
				i = 2
				data = /*line DhzIAv.go:1*/ append(data, "݈"...,
				)
			case 16:
				data = /*line DhzIAv.go:1*/ append(data, "\xdc\xd3\xd3"...,
				)
				i = 1
			case 8:
				data = /*line DhzIAv.go:1*/ append(data, "\xb7\xff\xf2"...,
				)
				i = 21
			case 11:
				data = /*line DhzIAv.go:1*/ append(data, "\xe1\xe2\xe1"...,
				)
				i = 25
			case 22:
				data = /*line DhzIAv.go:1*/ append(data, 190)
				i = 8
			case 21:
				i = 3
				data = /*line DhzIAv.go:1*/ append(data, "\xff\xfe"...,
				)
			case 17:
				data = /*line DhzIAv.go:1*/ append(data, "\xae\xf9\x97\x9b"...,
				)
				i = 6
			case 2:
				i = 7
				data = /*line DhzIAv.go:1*/ append(data, 223)
			case 25:
				i = 27
				data = /*line DhzIAv.go:1*/ append(data, "\xe4\xf6\xf7"...,
				)
			case 9:
				data = /*line DhzIAv.go:1*/ append(data, "\x8d\x87\xd5"...,
				)
				i = 5
			case 14:
				i = 23
				data = /*line DhzIAv.go:1*/ append(data, "\xc0\x95\x8f"...,
				)
			case 19:
				data = /*line DhzIAv.go:1*/ append(data, "\xc1\xbe\xee\xc4"...,
				)
				i = 22
			case 6:
				i = 28
				data = /*line DhzIAv.go:1*/ append(data, "\xfc\xf8"...,
				)
			case 15:
				data = /*line DhzIAv.go:1*/ append(data, "\xcd\xfa\xfe"...,
				)
				i = 17
			case 1:
				data = /*line DhzIAv.go:1*/ append(data, "\xde\xd9\xfb\xd5"...,
				)
				i = 12
			case 10:
				i = 16
				data = /*line DhzIAv.go:1*/ append(data, "ә"...,
				)
			case 12:
				data = /*line DhzIAv.go:1*/ append(data, "\xd8\xd1"...,
				)
				i = 0
			case 26:
				i = 15
				data = /*line DhzIAv.go:1*/ append(data, "\xc8\xc7"...,
				)
			case 27:
				data = /*line DhzIAv.go:1*/ append(data, "\xe9\xef\xe1\xa7"...,
				)
				i = 24
			case 5:
				i = 26
				data = /*line DhzIAv.go:1*/ append(data, "\x80\xc4\xd6"...,
				)
			case 29:
				i = 9
				data = /*line DhzIAv.go:1*/ append(data, 219)
			case 3:
				data = /*line DhzIAv.go:1*/ append(data, "\xf9\xe5\xd4"...,
				)
				i = 4
			case 18:
				i = 11
				data = /*line DhzIAv.go:1*/ append(data, "\xa2\xa8\xea"...,
				)
			case 4:
				data = /*line DhzIAv.go:1*/ append(data, "\xfb\xfb\xf6\xe1"...,
				)
				i = 18
			case 28:
				i = 20
				data = /*line DhzIAv.go:1*/ append(data, 170)
			case 24:
				i = 10
				data = /*line DhzIAv.go:1*/ append(data, "\xe6\xe9\xd5\xd8"...,
				)
			case 7:
				i = 29
				data = /*line DhzIAv.go:1*/ append(data, "\xc7\xca"...,
				)
			case 0:
				i = 14
				data = /*line DhzIAv.go:1*/ append(data, 223)
			}
		}
		return /*line DhzIAv.go:1*/ string(data)
	}(), /*line Do3VniNhwvfA.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)

	if dWmn_Gtk1d.BlockHeight < 2 {
		return nil
	}

	i3KmQgjb, ouJpb2aF3mlv, _ := /*line gfcfLR9_.go:1*/ cu1RSg2YR.tFPCl5(dWmn_Gtk1d)

	if !i3KmQgjb {
		return /*line N8cmTB.go:1*/ errors.New(func() /*line viWlZVFSlRrx.go:1*/ string {
			seed := /*line viWlZVFSlRrx.go:1*/ byte(29)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line viWlZVFSlRrx.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line viWlZVFSlRrx.go:1*/ fnc(126)(250)(251)(254)(225)(27)(170)(87)(228)(2)(28)(228)(5)(86)(174)(3)(93)(168)(247)(21)(235)(10)
			return /*line viWlZVFSlRrx.go:1*/ string(data)
		}())
	}

	l2pHR1ZhFG, gDcB23PkLL, k1qj7vimKIy7 := /*line FLIabN8Rl.go:1*/ cu1RSg2YR.bc.CommitWorkerBlock(ouJpb2aF3mlv.BlockHash, ouJpb2aF3mlv.BlockHeader.BlockHeight, cu1RSg2YR.voteQuorum, cu1RSg2YR.commitQuorum)

	if k1qj7vimKIy7 != nil {
		return /*line dOet0aJn.go:1*/ fmt.Errorf(func() /*line MHIgYs20GR1.go:1*/ string {
			key := [] /*line MHIgYs20GR1.go:1*/ byte("\xa6\xa4f}P\xbbg\xda6ʛ|o\x14A\x11Xe\u008cW\x04\v\xd2\xf5\x1bh3\x85")
			data := [] /*line MHIgYs20GR1.go:1*/ byte("\x01\xc9\xdc\xdap\x1e\xc8H\xa49\x0f\x9c҃\xae~\xc1\xd9\xe2\xee\xc3sn=hG\x88X\xfb")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line MHIgYs20GR1.go:1*/ string(data)
		}(), /*line haJxXapO.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)
	}

	for _, wFPszGk := range l2pHR1ZhFG {
		cu1RSg2YR.committedBlocks <- wFPszGk.(*blockchain.WorkerBlock)
	}

	for _, fuhYYZTnjKd := range gDcB23PkLL {
		cu1RSg2YR.forkedBlocks <- fuhYYZTnjKd.(*blockchain.WorkerBlock)
	}

	/*line weCYOTag.go:1*/
	log.Debugf(func() /*line SjmeRGJNI.go:1*/ string {
		data := [] /*line SjmeRGJNI.go:1*/ byte("[[9]K(#*mci\x85Bl1c\xe3\x91+\xfb\x0fn\x1as\x1c'\x82\fc\xd1\xcd\xf0ttt\xbdn\xec\xfa\xb9\xf6\x03lk@bdyJ\x80\xfc\xca.>\xd21+\xdbv d5\xbdw Tv X\xbf\n\xe6\xd2 \xfd\xefG\xfeD<%x\x0f1\xf2a\xa2\xac8 jv")
		positions := [...]byte{52, 87, 2, 84, 1, 38, 76, 1, 65, 54, 61, 47, 30, 79, 9, 46, 1, 22, 79, 48, 20, 7, 83, 26, 18, 27, 54, 75, 68, 47, 64, 29, 74, 62, 60, 41, 16, 90, 25, 17, 70, 25, 86, 20, 75, 39, 86, 83, 19, 4, 86, 86, 38, 79, 72, 1, 50, 88, 50, 50, 16, 77, 20, 56, 52, 75, 19, 82, 68, 37, 6, 6, 39, 71, 44, 37, 4, 54, 49, 79, 47, 2, 31, 70, 42, 57, 6, 14, 42, 7, 71, 11, 69, 51, 35, 86, 55, 53, 24, 60, 40, 50, 82, 32, 57, 4, 2, 40, 77, 71, 18, 52}
		for i := 0; i < 112; i += 2 {
			localKey := /*line SjmeRGJNI.go:1*/ byte(i) + /*line SjmeRGJNI.go:1*/ byte(positions[i]^positions[i+1]) + 212
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line SjmeRGJNI.go:1*/ string(data)
	}(), /*line VXJDjrZ5An_.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2, dWmn_Gtk1d.AVzGLiX4RWH)
	return nil
}

func (cu1RSg2YR *PBFT) uS1zadl60yhO(ouJpb2aF3mlv *blockchain.WorkerBlock) (bool, error) {
	if ouJpb2aF3mlv.BlockHeader.BlockHeight <= 2 {
		return true, nil
	}
	v61eq4CI, k1qj7vimKIy7 := /*line HYGhQWrSsmb8.go:1*/ cu1RSg2YR.bc.GetBlockByID(ouJpb2aF3mlv.BlockHeader.GvMaqmnWSRO)
	l_CyibPF := v61eq4CI.(*blockchain.WorkerBlock)
	if k1qj7vimKIy7 != nil {
		return false /*line xPwef0W2.go:1*/, fmt.Errorf(func() /*line LDoNsIKLgak.go:1*/ string {
			fullData := [] /*line LDoNsIKLgak.go:1*/ byte("\x1fT\xb7\xb8T\xf6A\x1ftڷ\xfd\xb0\x0f\x18y8*Ty\xd0\x11u\x9b\xe6T\x8eՅ\xbc'\xd3\xd0vB\x89!\xec\xf07p\x99\x96\xd3\xf9\x80\xaf-\xf8\x8c")
			idxKey := [] /*line LDoNsIKLgak.go:1*/ byte("]\xec9\xad\t\xdb")
			data := /*line LDoNsIKLgak.go:1*/ make([]byte, 0, 26)
			data = /*line LDoNsIKLgak.go:1*/ append(data, fullData[203^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]-fullData[192^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[246^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]-fullData[219^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[40^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])]+fullData[57^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])], fullData[30^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])]-fullData[38^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])], fullData[35^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])]-fullData[23^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])], fullData[142^ /*line LDoNsIKLgak.go:1*/ int(idxKey[3])]^fullData[166^ /*line LDoNsIKLgak.go:1*/ int(idxKey[3])], fullData[56^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])]^fullData[49^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])], fullData[193^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]-fullData[213^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[239^ /*line LDoNsIKLgak.go:1*/ int(idxKey[1])]+fullData[238^ /*line LDoNsIKLgak.go:1*/ int(idxKey[1])], fullData[247^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]-fullData[199^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[42^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])]+fullData[28^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])], fullData[214^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]+fullData[206^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[84^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])]^fullData[64^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])], fullData[127^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])]-fullData[118^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])], fullData[200^ /*line LDoNsIKLgak.go:1*/ int(idxKey[1])]-fullData[194^ /*line LDoNsIKLgak.go:1*/ int(idxKey[1])], fullData[29^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])]-fullData[5^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])], fullData[108^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])]-fullData[76^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])], fullData[243^ /*line LDoNsIKLgak.go:1*/ int(idxKey[1])]+fullData[197^ /*line LDoNsIKLgak.go:1*/ int(idxKey[1])], fullData[33^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])]^fullData[14^ /*line LDoNsIKLgak.go:1*/ int(idxKey[4])], fullData[51^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])]-fullData[32^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])], fullData[187^ /*line LDoNsIKLgak.go:1*/ int(idxKey[3])]+fullData[168^ /*line LDoNsIKLgak.go:1*/ int(idxKey[3])], fullData[195^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]+fullData[201^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[125^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])]^fullData[123^ /*line LDoNsIKLgak.go:1*/ int(idxKey[0])], fullData[212^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])]-fullData[223^ /*line LDoNsIKLgak.go:1*/ int(idxKey[5])], fullData[63^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])]^fullData[30^ /*line LDoNsIKLgak.go:1*/ int(idxKey[2])])
			return /*line LDoNsIKLgak.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	if (ouJpb2aF3mlv.BlockHeader.BlockHeight <= cu1RSg2YR.lastVotedBlockHeight) || (l_CyibPF.BlockHeader.BlockHeight < cu1RSg2YR.preferredBlockHeight) {
		return false, nil
	}
	return true, nil
}

func (cu1RSg2YR *PBFT) tFPCl5(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) (bool, *blockchain.WorkerBlock, error) {
	v61eq4CI, k1qj7vimKIy7 := /*line H1fuJdzJhwm.go:1*/ cu1RSg2YR.bc.GetParentWorkerBlock(dWmn_Gtk1d.ZWsU_2)
	l_CyibPF := v61eq4CI.(*blockchain.WorkerBlock)
	if k1qj7vimKIy7 != nil {
		return false, nil /*line K7Tpt5.go:1*/, fmt.Errorf(func() /*line MlBtgd.go:1*/ string {
			fullData := [] /*line MlBtgd.go:1*/ byte("\x83\xfe\x88k+\xb5\xb1\xeb\xdc\xed\xf8\x87߅\xab\xe8\x9c\xfa\x97\xa5K\xbf\x00\xb0Ⱥ\x81\x8b\x84\xe0Bt\x9a\xb18\xdb-n\xa4\xb4\xbe\xb9\x00\xb6\xd5\x1a\xb7CtC\xc3E\x91\xaa")
			idxKey := [] /*line MlBtgd.go:1*/ byte("^\x88\xa8\x9fb\xcc*\rsԶ\xff\x01\r\xc0")
			data := /*line MlBtgd.go:1*/ make([]byte, 0, 28)
			data = /*line MlBtgd.go:1*/ append(data, fullData[25^ /*line MlBtgd.go:1*/ int(idxKey[12])]^fullData[15^ /*line MlBtgd.go:1*/ int(idxKey[12])], fullData[67^ /*line MlBtgd.go:1*/ int(idxKey[0])]+fullData[68^ /*line MlBtgd.go:1*/ int(idxKey[0])], fullData[71^ /*line MlBtgd.go:1*/ int(idxKey[4])]-fullData[116^ /*line MlBtgd.go:1*/ int(idxKey[4])], fullData[228^ /*line MlBtgd.go:1*/ int(idxKey[14])]^fullData[241^ /*line MlBtgd.go:1*/ int(idxKey[14])], fullData[130^ /*line MlBtgd.go:1*/ int(idxKey[2])]-fullData[156^ /*line MlBtgd.go:1*/ int(idxKey[2])], fullData[91^ /*line MlBtgd.go:1*/ int(idxKey[0])]+fullData[75^ /*line MlBtgd.go:1*/ int(idxKey[0])], fullData[195^ /*line MlBtgd.go:1*/ int(idxKey[14])]^fullData[212^ /*line MlBtgd.go:1*/ int(idxKey[14])], fullData[180^ /*line MlBtgd.go:1*/ int(idxKey[3])]^fullData[179^ /*line MlBtgd.go:1*/ int(idxKey[3])], fullData[100^ /*line MlBtgd.go:1*/ int(idxKey[4])]-fullData[124^ /*line MlBtgd.go:1*/ int(idxKey[4])], fullData[249^ /*line MlBtgd.go:1*/ int(idxKey[5])]+fullData[254^ /*line MlBtgd.go:1*/ int(idxKey[5])], fullData[119^ /*line MlBtgd.go:1*/ int(idxKey[8])]-fullData[91^ /*line MlBtgd.go:1*/ int(idxKey[8])], fullData[192^ /*line MlBtgd.go:1*/ int(idxKey[14])]-fullData[237^ /*line MlBtgd.go:1*/ int(idxKey[14])], fullData[37^ /*line MlBtgd.go:1*/ int(idxKey[6])]^fullData[58^ /*line MlBtgd.go:1*/ int(idxKey[6])], fullData[12^ /*line MlBtgd.go:1*/ int(idxKey[6])]-fullData[54^ /*line MlBtgd.go:1*/ int(idxKey[6])], fullData[9^ /*line MlBtgd.go:1*/ int(idxKey[6])]^fullData[51^ /*line MlBtgd.go:1*/ int(idxKey[6])], fullData[221^ /*line MlBtgd.go:1*/ int(idxKey[5])]+fullData[211^ /*line MlBtgd.go:1*/ int(idxKey[5])], fullData[6^ /*line MlBtgd.go:1*/ int(idxKey[13])]^fullData[12^ /*line MlBtgd.go:1*/ int(idxKey[13])], fullData[205^ /*line MlBtgd.go:1*/ int(idxKey[14])]^fullData[211^ /*line MlBtgd.go:1*/ int(idxKey[14])], fullData[7^ /*line MlBtgd.go:1*/ int(idxKey[13])]^fullData[45^ /*line MlBtgd.go:1*/ int(idxKey[13])], fullData[151^ /*line MlBtgd.go:1*/ int(idxKey[3])]^fullData[136^ /*line MlBtgd.go:1*/ int(idxKey[3])], fullData[42^ /*line MlBtgd.go:1*/ int(idxKey[13])]-fullData[62^ /*line MlBtgd.go:1*/ int(idxKey[13])], fullData[157^ /*line MlBtgd.go:1*/ int(idxKey[3])]^fullData[152^ /*line MlBtgd.go:1*/ int(idxKey[3])], fullData[164^ /*line MlBtgd.go:1*/ int(idxKey[2])]-fullData[152^ /*line MlBtgd.go:1*/ int(idxKey[2])], fullData[147^ /*line MlBtgd.go:1*/ int(idxKey[1])]^fullData[169^ /*line MlBtgd.go:1*/ int(idxKey[1])], fullData[141^ /*line MlBtgd.go:1*/ int(idxKey[3])]^fullData[177^ /*line MlBtgd.go:1*/ int(idxKey[3])], fullData[197^ /*line MlBtgd.go:1*/ int(idxKey[5])]+fullData[238^ /*line MlBtgd.go:1*/ int(idxKey[5])], fullData[40^ /*line MlBtgd.go:1*/ int(idxKey[12])]-fullData[46^ /*line MlBtgd.go:1*/ int(idxKey[12])])
			return /*line MlBtgd.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	if (l_CyibPF.BlockHeader.BlockHeight + 1) == dWmn_Gtk1d.BlockHeight {
		return true, l_CyibPF, nil
	}
	return false, nil, nil
}

func (cu1RSg2YR *PBFT) GenerateFDCM() {
	k2gdNf := /*line MagsFK.go:1*/ make([]*blockchain.LocalSnapshot, 0)
	k2gdNf = /*line V2n4iO.go:1*/ append(k2gdNf, cu1RSg2YR.gSnapshot...)
	if /*line ZshECW.go:1*/ len(cu1RSg2YR.gSnapshot) > /*line yfPt4K8Zj.go:1*/ len(k2gdNf) {
		cu1RSg2YR.gSnapshot = cu1RSg2YR.gSnapshot[ /*line IHA2DA0K.go:1*/ len(k2gdNf):]
	} else {
		cu1RSg2YR.gSnapshot = []*blockchain.LocalSnapshot{}
	}
	JEvdtw := /*line slq_N3PjNKep.go:1*/ make([]*blockchain.LocalContract, 0)
	JEvdtw = /*line YzmALL.go:1*/ append(JEvdtw, cu1RSg2YR.gCode...)
	if /*line zVoYxG.go:1*/ len(cu1RSg2YR.gCode) > /*line yard0CSTPq4.go:1*/ len(JEvdtw) {
		cu1RSg2YR.gCode = cu1RSg2YR.gCode[ /*line Rs_mby1Wr4CY.go:1*/ len(JEvdtw):]
	} else {
		cu1RSg2YR.gCode = []*blockchain.LocalContract{}
	}
	for _, lWx2EeY := range JEvdtw {
		if /*line POOfU_qzhm.go:1*/ utils.CalculateShardToSend([]common.Address{lWx2EeY.Address} /*line IeDG_Tj7.go:1*/, config.GetConfig().ShardCount)[0] != /*line S74YNjupm.go:1*/ cu1RSg2YR.Shard() {
			var vByID1 []byte

			if /*line DN0Lao1.go:1*/ len(lWx2EeY.Code) > 0 {
				vByID1 = lWx2EeY.Code
			} else {
				/*line Hj4MK3qTx.go:1*/ log.Jp980o6YjM(func() /*line GpXs7RavkHf3.go:1*/ string {
					data := [] /*line GpXs7RavkHf3.go:1*/ byte("\xe7gE\x7fe\xbaaP,I\x04C\x18^NE+\xd7t\x18 ck1!rajt\x18od\x116\x04rl\xd6r")
					positions := [...]byte{33, 34, 33, 13, 0, 3, 10, 27, 24, 12, 16, 22, 3, 32, 2, 14, 1, 2, 14, 9, 16, 5, 8, 10, 23, 29, 23, 32, 2, 19, 32, 22, 17, 36, 37, 17, 7, 32, 32, 16}
					for i := 0; i < 40; i += 2 {
						localKey := /*line GpXs7RavkHf3.go:1*/ byte(i) + /*line GpXs7RavkHf3.go:1*/ byte(positions[i]^positions[i+1]) + 80
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
					}
					return /*line GpXs7RavkHf3.go:1*/ string(data)
				}())
			}
			/*line J0N_Sb.go:1*/ cu1RSg2YR.bc.GetStateDB().CreateTemporaryAccount(lWx2EeY.Address)
			/*line jQ_bNuUw.go:1*/ cu1RSg2YR.bc.GetStateDB().SetCode(lWx2EeY.Address, vByID1)

			cu1RSg2YR.fdcm = /*line kRNPdaKXjEo7.go:1*/ append(cu1RSg2YR.fdcm, &lWx2EeY.Address)
		}
	}
globalSnapshot:
	for _, aVU16qtmHkO := range k2gdNf {
		if /*line IwWfXmD.go:1*/ utils.CalculateShardToSend([]common.Address{aVU16qtmHkO.Address} /*line nvOhgP.go:1*/, config.GetConfig().ShardCount)[0] != /*line DuTCHPKpz12.go:1*/ cu1RSg2YR.Shard() {
			iRgBar := func( /*line K538tS.go:1*/ b8kFRD common.Address) bool {
				for _, omhAOEXAcby := range cu1RSg2YR.fdcm {
					if *omhAOEXAcby == b8kFRD {
						return true
					}
				}
				return false
			}(aVU16qtmHkO.Address)

			if iRgBar {

				for _, omhAOEXAcby := range cu1RSg2YR.fdcm {
					if *omhAOEXAcby == aVU16qtmHkO.Address {
						/*line vmjOYJk.go:1*/ cu1RSg2YR.bc.GetStateDB().SetState(aVU16qtmHkO.Address, aVU16qtmHkO.Slot /*line Ec5V9m97.go:1*/, common.HexToHash(aVU16qtmHkO.Value))
						continue globalSnapshot
					}
				}
			} else {

				/*line ykmTpI.go:1*/
				cu1RSg2YR.bc.GetStateDB().CreateTemporaryAccount(aVU16qtmHkO.Address)
				cu1RSg2YR.fdcm = /*line FoVySTPz4.go:1*/ append(cu1RSg2YR.fdcm, &aVU16qtmHkO.Address)
				ybWVGGO4SaK, k1qj7vimKIy7 := /*line R51jxq9eOs8.go:1*/ strconv.ParseInt(aVU16qtmHkO.Value, 0, 64)
				if k1qj7vimKIy7 != nil {
					/*line dQsimv.go:1*/ log.Jp980o6YjM(func() /*line ndQKYD.go:1*/ string {
						seed := /*line ndQKYD.go:1*/ byte(206)
						var data []byte
						type decFunc func(byte) decFunc
						var fnc decFunc
						fnc = func(x byte) decFunc { data = /*line ndQKYD.go:1*/ append(data, x+seed); seed += x; return fnc }
						/*line ndQKYD.go:1*/ fnc(141)(202)(81)(231)(195)(8)(31)(30)(9)(247)(13)(239)(19)(241)(225)(254)(255)(10)(220)(247)(37)(45)(0)(253)(3)(174)(5)(81)
						return /*line ndQKYD.go:1*/ string(data)
					}(), /*line E2idAM5MnnW4.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)
					/*line H62LDw7.go:1*/ cu1RSg2YR.CloseFDCM()
					return
				}
				/*line hm6m66dMsk.go:1*/ cu1RSg2YR.bc.GetStateDB().SetBalance(aVU16qtmHkO.Address /*line AMwNha.go:1*/, uint256.NewInt( /*line U_F9_M0Ag.go:1*/ uint64(ybWVGGO4SaK)))
				/*line GfoxPI4LvlPn.go:1*/ cu1RSg2YR.bc.GetStateDB().SetNonce(aVU16qtmHkO.Address, 0)
			}
		}
	}
}

func (cu1RSg2YR *PBFT) vhgIe3tbXH(yxYJX82Xsq []*message.Transaction, kt6kudfae bool) {
	prF0uu2Fu2 := /*line QhpEVqSD.go:1*/ evm.WTPTcYkXjM( /*line qfBLu7EsDzy.go:1*/ cu1RSg2YR.GetHighBlockHeight() /*line cx00S4ir2q_.go:1*/, string( /*line LpyeWJpyb_.go:1*/ rune( /*line TzUJdnScH.go:1*/ time.Now().Unix())))
	cXyBcMXjM := /*line AAdKyG.go:1*/ runtime.FzlY0sma(prF0uu2Fu2 /*line Jwi7fH.go:1*/, cu1RSg2YR.bc.GetStateDB())
	if !kt6kudfae {
		/*line IayTRxHOS.go:1*/ cu1RSg2YR.GenerateFDCM()
	}
	for _, bQZ9a4xL := range yxYJX82Xsq {
		if bQZ9a4xL.IsCrossShardTx {

			if bQZ9a4xL.TXType == types.TRANSFER {
				k1qj7vimKIy7 := /*line Bro1Fckk2aSB.go:1*/ evm.Fqa2qecFzII( /*line ace3FcQn.go:1*/ cu1RSg2YR.bc.GetStateDB(), bQZ9a4xL.From, bQZ9a4xL.To /*line vfuaWjG0mw.go:1*/, uint256.NewInt( /*line CMlp9bb.go:1*/ uint64(bQZ9a4xL.Value)))
				if k1qj7vimKIy7 != nil {
					bQZ9a4xL.TXType = types.ABORT
					/*line J13S3ZCK.go:1*/ log.Jp980o6YjM(func() /*line hdTaY4K_Hwta.go:1*/ string {
						fullData := [] /*line hdTaY4K_Hwta.go:1*/ byte("\xd6]>\x99Q]X\x1d\x14r\xb9\xfc\x17r\xa6\xdb\x04\xe2\xbdֻ\x16\xbc\xfc;Fx75-Oۧ\n\xe1\xf7\xe8\xfc\xdc\xf15\x8d;\xe8\xa4σO\xcf\xc5z\xf5\xff\x93\xeaAbYu\xd1\x7f\xdf^\x92&\xcbk;G\x17\x19\xb4\x97\xe8p\xfb\xdc5\xb0\xed\xdf\x15\xa8\vJ\xdd[8Gl\xa1\xc8RP\x84\x7f\xdc\xe3\x96[N\x1a\xc8'\x12\x9f\xa9b\xb6:J\x163S\x84K\xc2\xd3\xe9[\xbb_zD\x12\xca;\xec%\xc3\x1e\x8d[(\x80n\x9c\x1f\x152\xceo\xcf\x0e\v\x11]\x10\xd7ly\x98\x13\x96Mb")
						idxKey := [] /*line hdTaY4K_Hwta.go:1*/ byte("2lxŉ\xf95$a\xe4NR\x82[\xbb)")
						data := /*line hdTaY4K_Hwta.go:1*/ make([]byte, 0, 79)
						data = /*line hdTaY4K_Hwta.go:1*/ append(data, fullData[111^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]^fullData[162^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[217^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])]+fullData[131^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])], fullData[123^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]+fullData[36^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[225^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])]^fullData[151^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])], fullData[110^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]-fullData[185^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[41^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]^fullData[107^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[7^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]-fullData[181^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[33^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]+fullData[17^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[160^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]^fullData[129^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[31^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])]^fullData[3^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])], fullData[179^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])]-fullData[19^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])], fullData[117^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])]-fullData[106^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])], fullData[70^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]+fullData[94^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[172^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])]-fullData[168^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])], fullData[252^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[5])]^fullData[251^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[5])], fullData[121^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])]-fullData[206^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])], fullData[74^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])]-fullData[27^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])], fullData[36^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])]^fullData[108^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])], fullData[79^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]+fullData[53^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[63^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[6])]-fullData[109^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[6])], fullData[125^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]-fullData[169^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[244^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[1])]+fullData[232^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[1])], fullData[204^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])]+fullData[84^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])], fullData[33^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])]-fullData[29^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])], fullData[213^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]-fullData[127^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[213^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])]^fullData[143^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])], fullData[113^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[5])]^fullData[202^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[5])], fullData[126^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])]^fullData[38^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])], fullData[31^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]+fullData[119^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[117^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]-fullData[207^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[23^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[1])]+fullData[74^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[1])], fullData[203^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])]+fullData[13^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])], fullData[195^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])]+fullData[162^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])], fullData[194^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])]+fullData[20^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])], fullData[85^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])]-fullData[43^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])], fullData[49^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]-fullData[118^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[40^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])]+fullData[79^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])], fullData[26^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]^fullData[51^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[84^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]+fullData[123^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[54^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]-fullData[29^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[125^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]-fullData[73^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[57^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]-fullData[119^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[11^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])]+fullData[75^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])], fullData[5^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]^fullData[58^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[82^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]+fullData[60^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[61^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])]-fullData[218^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])], fullData[73^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])]^fullData[213^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[13])], fullData[192^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])]-fullData[89^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[11])], fullData[15^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])]^fullData[100^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])], fullData[17^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]-fullData[64^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[105^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])]^fullData[86^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])], fullData[180^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])]+fullData[86^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])], fullData[118^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]^fullData[122^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[114^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])]^fullData[91^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[10])], fullData[166^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]-fullData[243^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[162^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])]-fullData[88^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])], fullData[79^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]^fullData[96^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[80^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]+fullData[2^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[201^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])]^fullData[72^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])], fullData[133^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[14])]^fullData[197^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[14])], fullData[232^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])]-fullData[155^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])], fullData[115^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])]-fullData[165^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])], fullData[20^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])]+fullData[201^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[12])], fullData[100^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]-fullData[242^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[59^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])]^fullData[53^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[0])], fullData[51^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]-fullData[8^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[147^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])]-fullData[221^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[3])], fullData[171^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])]+fullData[34^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[7])], fullData[76^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])]+fullData[94^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])], fullData[226^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])]+fullData[45^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])], fullData[252^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])]+fullData[138^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])], fullData[76^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])]-fullData[69^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[2])], fullData[31^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[6])]^fullData[183^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[6])], fullData[221^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])]+fullData[109^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[9])], fullData[201^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])]-fullData[246^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[4])], fullData[163^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])]^fullData[1^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[15])], fullData[137^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[14])]^fullData[194^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[14])], fullData[117^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])]+fullData[25^ /*line hdTaY4K_Hwta.go:1*/ int(idxKey[8])])
						return /*line hdTaY4K_Hwta.go:1*/ string(data)
					}(), /*line Ioc79ktS.go:1*/ cu1RSg2YR.Shard() /*line bZMjvPapa.go:1*/, cu1RSg2YR.ID(), bQZ9a4xL.Hash, k1qj7vimKIy7)
					continue
				}

			} else if bQZ9a4xL.TXType == types.SMARTCONTRACT {
				_, _, k1qj7vimKIy7 := /*line J0PK3SU.go:1*/ evm.DFPUN4b8lsOr(cXyBcMXjM /*line IP0ZjgzG.go:1*/, cu1RSg2YR.bc.GetStateDB(), bQZ9a4xL.Data, bQZ9a4xL.From, bQZ9a4xL.To)
				if k1qj7vimKIy7 != nil {
					bQZ9a4xL.TXType = types.ABORT
					/*line NfXUvfgQl.go:1*/ log.Jp980o6YjM(func() /*line RM0YWSe7.go:1*/ string {
						key := [] /*line RM0YWSe7.go:1*/ byte("\x9c\xa7\x04\x00R\xcel\xc6y\x81\xce#QC\xfa\x9b\x83P[V\x02\x01\xabj\xfb;{\x89\xad\\\xb1V\n.\x9d\x0f\x9b\xa7\xeaD}Z}\x11E\x824\x97\xd9\x0f\xad\xb6Z\xb39\xf9T,\xbf\xaa+\x14\xb5\xe0\x1c\x05W\x82\xc3<thE\xf0\xd5\U000df6c0\xe3\x1b\x9d\x9e\xacS\x91\x03\x9a\x15\xd1%;\x88\xcdͲ\xf0ܮ\x88E\xc3\xf5\x80\x8fҏ@y*\x88\x06\x1a\xf4\x03\\\x8f#\xb0㰱\xf2\x95\n\xd9Ih\xfa\xfb\x04\xa4\xd1\xc77\xb2\xfe")
						data := [] /*line RM0YWSe7.go:1*/ byte("\xf7\xccz]r\xf6\xd1>\xde\xe4C\x97\xb6\x97l\xfc\xf1ü\xb9vj\x1a\xd8$[\xde\xfb\x1c\xcf$v}\x9b\xfe\x81\x0f\nY\xb2\xf1\xcc\xdet\xb9\xa2\xa8\t:} \x17\xbd'\xa2h\xc2L$\"\x90w*T\x85tŢ)\x9d\xdd\xd4e\x15M-\xbf\xc0\xf6\x0f;\xff\xff\x18\xb4\xfff\xffO\xf1J\xb1\xb4\xedA$QE\x1c\xa8\xb8/d\xf4\xc9\U000b4da5J\xf0u\x8eYo|\x02\x8f\x1fW\xea\xd1\x17\v6\xf9\xad\xcdjjw\rE\x01W\xd7t")
						for i, b := range key {
							data[i] = data[i] - b
						}
						return /*line RM0YWSe7.go:1*/ string(data)
					}(), /*line W92JGz0ih.go:1*/ cu1RSg2YR.Shard(), bQZ9a4xL.Hash, k1qj7vimKIy7 /*line DwAm2vz3.go:1*/, cu1RSg2YR.bc.GetStateDB().GetBalance(bQZ9a4xL.From) /*line R_wJM_uWG_.go:1*/, cu1RSg2YR.bc.GetStateDB().GetState(bQZ9a4xL.ExternalAddressList[0] /*line iZBWlGoEu.go:1*/, common.HexToHash("0")) /*line EAUb4MiiXaa.go:1*/, cu1RSg2YR.bc.GetStateDB().GetState(bQZ9a4xL.ExternalAddressList[2] /*line WAS_k0CMDE.go:1*/, common.HexToHash("0")) /*line sARsz34nyj.go:1*/, cu1RSg2YR.bc.GetStateDB().GetState(bQZ9a4xL.ExternalAddressList[1] /*line j00cm9ZHKrt.go:1*/, common.BytesToHash( /*line Z9uqxK3yUT.go:1*/ utils.GffGNE(bQZ9a4xL.From, 0))))
					continue
				}

			}
		} else {

			if bQZ9a4xL.TXType == types.TRANSFER {
				k1qj7vimKIy7 := /*line yLmr9xUvuuHX.go:1*/ evm.Fqa2qecFzII( /*line IupUD8zPPqs.go:1*/ cu1RSg2YR.bc.GetStateDB(), bQZ9a4xL.From, bQZ9a4xL.To /*line l6Mpwf.go:1*/, uint256.NewInt( /*line YI2TxSmMGS.go:1*/ uint64(bQZ9a4xL.Value)))
				if k1qj7vimKIy7 != nil {
					bQZ9a4xL.TXType = types.ABORT
					/*line bJGzlfDC49.go:1*/ log.Jp980o6YjM(func() /*line KCl9HCIkJr.go:1*/ string {
						data := [] /*line KCl9HCIkJr.go:1*/ byte("g%v\xa0 (\x9bI\xabĥt\xb6\xc0\x10\xbf\x8dsa\xa0\xbd\xa9o\xbc) \xf4oc\x9alG\x8f\xa9aZYwr)\n{\x8bnsac\x88io\x90we\x8ce\xb5ut\x83[}\xa6\x1d|i\xb9\xe4\xd7J\x9e\x1a\xc0\xdf")
						positions := [...]byte{59, 35, 51, 8, 40, 72, 0, 69, 69, 70, 59, 68, 58, 20, 55, 3, 72, 14, 0, 7, 21, 65, 62, 13, 68, 70, 37, 6, 62, 7, 37, 36, 53, 66, 67, 47, 32, 19, 51, 31, 23, 50, 41, 58, 15, 31, 59, 72, 71, 21, 50, 13, 39, 39, 42, 55, 41, 14, 10, 40, 36, 3, 26, 61, 14, 61, 33, 12, 60, 63, 50, 9, 60, 31, 16, 29}
						for i := 0; i < 76; i += 2 {
							localKey := /*line KCl9HCIkJr.go:1*/ byte(i) + /*line KCl9HCIkJr.go:1*/ byte(positions[i]^positions[i+1]) + 213
							data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
						}
						return /*line KCl9HCIkJr.go:1*/ string(data)
					}(), /*line T5WuX1yU.go:1*/ cu1RSg2YR.ID(), bQZ9a4xL.Hash, k1qj7vimKIy7)
					continue
				}

			} else if bQZ9a4xL.TXType == types.SMARTCONTRACT {

				_, _, k1qj7vimKIy7 := /*line OakyKSdLhu.go:1*/ evm.DFPUN4b8lsOr(cXyBcMXjM /*line AsW7Eb.go:1*/, cu1RSg2YR.bc.GetStateDB(), bQZ9a4xL.Data, bQZ9a4xL.From, bQZ9a4xL.To)
				if k1qj7vimKIy7 != nil {
					bQZ9a4xL.TXType = types.ABORT
					/*line TfbkDas.go:1*/ log.Jp980o6YjM(func() /*line DEOEBQi3JG.go:1*/ string {
						seed := /*line DEOEBQi3JG.go:1*/ byte(73)
						var data []byte
						type decFunc func(byte) decFunc
						var fnc decFunc
						fnc = func(x byte) decFunc { data = /*line DEOEBQi3JG.go:1*/ append(data, x-seed); seed += x; return fnc }
						/*line DEOEBQi3JG.go:1*/ fnc(164)(18)(117)(209)(101)(210)(225)(213)(151)(44)(106)(211)(151)(29)(88)(159)(75)(155)(36)(74)(165)(63)(132)(7)(201)(137)(94)(191)(114)(226)(207)(82)(247)(232)(196)(153)(52)(87)(186)(115)(236)(214)(155)(56)(129)(174)(176)(94)(171)(99)(203)(132)(10)(37)(63)(132)(7)(192)(197)(157)(39)(76)(170)(83)(155)(60)(119)(160)(134)(7)(22)(47)(18)(41)(165)(12)(254)(1)(83)(92)(172)(154)(51)(113)(215)(187)(107)(216)(133)(240)(229)(27)(236)(204)(236)(214)(155)(62)(129)(180)(187)(111)(225)(199)(84)(142)(33)(147)(220)(172)(160)(71)(147)(23)(53)(30)(143)(23)(49)(103)(148)(14)(33)(147)(220)(172)(156)(57)(125)(249)(246)(226)(207)(100)(174)(97)(19)
						return /*line DEOEBQi3JG.go:1*/ string(data)
					}(), /*line QsIBRyqI4FL.go:1*/ cu1RSg2YR.Shard(), bQZ9a4xL.Hash, k1qj7vimKIy7 /*line MB2cCn7zA4at.go:1*/, cu1RSg2YR.bc.GetStateDB().GetBalance(bQZ9a4xL.From) /*line gvrTmH.go:1*/, cu1RSg2YR.bc.GetStateDB().GetState(bQZ9a4xL.ExternalAddressList[0] /*line J6t1D5hlvuKJ.go:1*/, common.HexToHash("0")) /*line fb_LEPJvaa.go:1*/, cu1RSg2YR.bc.GetStateDB().GetState(bQZ9a4xL.ExternalAddressList[2] /*line uTZQLSNafC.go:1*/, common.HexToHash("0")) /*line N8VfL7xLX.go:1*/, cu1RSg2YR.bc.GetStateDB().GetState(bQZ9a4xL.ExternalAddressList[1] /*line Ghf9Ix.go:1*/, common.BytesToHash( /*line sKvFabw.go:1*/ utils.GffGNE(bQZ9a4xL.From, 0))))
					continue
				}

			} else if bQZ9a4xL.TXType == types.DEPLOY {
				_, _, k1qj7vimKIy7 := /*line pU3CFqZ.go:1*/ evm.TR_8hD4NdZ(cXyBcMXjM /*line tv_eiM.go:1*/, cu1RSg2YR.bc.GetStateDB(), bQZ9a4xL.Data, bQZ9a4xL.From)
				if k1qj7vimKIy7 != nil {
					bQZ9a4xL.TXType = types.ABORT
					/*line rm4Mn3X.go:1*/ log.Jp980o6YjM(func() /*line xa4vIfjXNXMF.go:1*/ string {
						data := [] /*line xa4vIfjXNXMF.go:1*/ byte("z\x82v\x87\xb0eq\xac\x1e\xfduV\xa8Tra\xc3}\x85\x11@_%ƙݠoc@r$\xf4\xab\x17]AQR\xae\u05fd\ts\xe8ct\nop Xxee\xd5+\xa9o\x93ݔ\x9d\xd9;s1\x17T\x06pv")
						positions := [...]byte{32, 4, 65, 20, 68, 59, 62, 26, 16, 9, 33, 42, 22, 22, 1, 63, 5, 33, 18, 12, 41, 57, 70, 47, 39, 55, 40, 3, 11, 7, 49, 24, 7, 22, 9, 59, 65, 34, 64, 64, 60, 25, 56, 56, 56, 33, 30, 55, 24, 23, 68, 64, 44, 61, 19, 15, 33, 17, 23, 54, 69, 62, 41, 18, 31, 61, 41, 63, 61, 31, 18, 37, 36, 8, 39, 64, 47, 21, 29, 0, 35, 68, 51, 69, 18, 30, 61, 3, 59, 26, 38, 15, 36, 41, 39, 26, 6, 66, 17, 67, 16, 41}
						for i := 0; i < 102; i += 2 {
							localKey := /*line xa4vIfjXNXMF.go:1*/ byte(i) + /*line xa4vIfjXNXMF.go:1*/ byte(positions[i]^positions[i+1]) + 176
							data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
						}
						return /*line xa4vIfjXNXMF.go:1*/ string(data)
					}(), /*line Acw8YABm6Pe.go:1*/ cu1RSg2YR.ID(), bQZ9a4xL.Hash, k1qj7vimKIy7)
					continue
				}

			}
		}
	}
	if !kt6kudfae {
		/*line ApiKfJs.go:1*/ cu1RSg2YR.CloseFDCM()
	}

	iSve7rE_k1 := /*line Xm8wsHc8uv.go:1*/ cu1RSg2YR.bc.GetStateDBRoot( /*line in644kO9Txf.go:1*/ cu1RSg2YR.GetHighBlockHeight() + 1)
	jBIbIsh := /*line _fudTR6.go:1*/ cu1RSg2YR.bc.GetExecutedState(iSve7rE_k1)
	cu1RSg2YR.executedState[ /*line V3epHG2juF.go:1*/ cu1RSg2YR.GetHighBlockHeight()+1] = jBIbIsh
}

func (cu1RSg2YR *PBFT) CloseFDCM() {
	xBKFpDBB9kW := /*line ckhFzW2Xb0D.go:1*/ make([]*common.Address, 0)
	xBKFpDBB9kW = /*line a972ZDQjcku.go:1*/ append(xBKFpDBB9kW, cu1RSg2YR.fdcm...)
	if /*line DFfWMCHIj.go:1*/ len(cu1RSg2YR.fdcm) > /*line TfPlhE9boT.go:1*/ len(xBKFpDBB9kW) {
		cu1RSg2YR.fdcm = cu1RSg2YR.fdcm[ /*line KPy10W8Suy.go:1*/ len(xBKFpDBB9kW):]
	} else {
		cu1RSg2YR.fdcm = []*common.Address{}
	}
	for _, omhAOEXAcby := range xBKFpDBB9kW {
		/*line abo7Eb8NuY6n.go:1*/ cu1RSg2YR.bc.GetStateDB().DeleteAccount(*omhAOEXAcby)
	}
	cu1RSg2YR.fdcm = []*common.Address{}
}

func (cu1RSg2YR *PBFT) ProcessRemoteTmo(iv5UGEus9T *pacemaker.TMO) {
	/*line Kme8Hrao8VCw.go:1*/ log.Debugf(func() /*line wGrcjemLn.go:1*/ string {
		data := [] /*line wGrcjemLn.go:1*/ byte("\xfb\xd3\x1e]8(P6o4\xf6ssR\xf8\x91\xe8t\x97TiT\xd7 p\xd5oce\xb7\x1cicg%\xfbemooW\xf9\xfeW\xea\xc4oX\x95\x91\xf7\xbdom.\x9e\xe5 f\xdf\xe7\"\x04ewv\xd3P7\x7f)v")
		positions := [...]byte{45, 21, 69, 20, 67, 60, 66, 51, 44, 15, 50, 14, 54, 59, 34, 40, 2, 4, 70, 18, 21, 7, 43, 48, 34, 32, 22, 67, 48, 1, 16, 43, 39, 68, 21, 48, 21, 55, 56, 10, 29, 67, 39, 0, 30, 29, 48, 20, 35, 60, 4, 49, 42, 62, 69, 47, 32, 25, 47, 61, 66, 34, 44, 15, 41, 45, 18, 68, 49, 47, 18, 9, 44, 21, 44, 18}
		for i := 0; i < 76; i += 2 {
			localKey := /*line wGrcjemLn.go:1*/ byte(i) + /*line wGrcjemLn.go:1*/ byte(positions[i]^positions[i+1]) + 40
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line wGrcjemLn.go:1*/ string(data)
	}(), /*line Hn9OamdM.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.Ln9AEaadIIY, iv5UGEus9T.Ig3RQ9F)

	if /*line QlavXdX.go:1*/ cu1RSg2YR.State() != types.VIEWCHANGING {
		if !( /*line DZ3mBG.go:1*/ cu1RSg2YR.pm.GetCurView() < iv5UGEus9T.Ig3RQ9F) {
			return
		}

		for ik_N67hORWa7, jEIq_MYzWdn0 := range cu1RSg2YR.detectedTmos {

			if !( /*line InfbhYNKH.go:1*/ cu1RSg2YR.pm.GetCurView() < jEIq_MYzWdn0.Ig3RQ9F) {
				/*line obx13h6jHe.go:1*/ delete(cu1RSg2YR.detectedTmos, ik_N67hORWa7)
			}
		}

		cu1RSg2YR.detectedTmos[iv5UGEus9T.Ln9AEaadIIY] = iv5UGEus9T

		if /*line YX8BaacTBNg.go:1*/ len(cu1RSg2YR.detectedTmos) > ( /*line asuj_3V.go:1*/ config.GetConfig().CommitteeNumber-1)/3 {
			/*line OD0QCaLgbY.go:1*/ log.Debugf(func() /*line DVpUieavcx.go:1*/ string {
				data := /*line DVpUieavcx.go:1*/ make([]byte, 0, 74)
				i := 10
				decryptKey := 59
				for counter := 0; i != 21; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 33:
						i = 29
						data = /*line DVpUieavcx.go:1*/ append(data, "\xaa\xa3\x83"...,
						)
					case 16:
						i = 2
						data = /*line DVpUieavcx.go:1*/ append(data, "\x98S"...,
						)
					case 7:
						i = 32
						data = /*line DVpUieavcx.go:1*/ append(data, "\x869\x80F"...,
						)
					case 15:
						i = 26
						data = /*line DVpUieavcx.go:1*/ append(data, 94)
					case 13:
						i = 14
						data = /*line DVpUieavcx.go:1*/ append(data, "\x9b\xa1"...,
						)
					case 30:
						data = /*line DVpUieavcx.go:1*/ append(data, "\x94\x9a"...,
						)
						i = 22
					case 31:
						data = /*line DVpUieavcx.go:1*/ append(data, 128)
						i = 7
					case 3:
						data = /*line DVpUieavcx.go:1*/ append(data, 134)
						i = 31
					case 4:
						i = 28
						data = /*line DVpUieavcx.go:1*/ append(data, "b\xb4"...,
						)
					case 20:
						i = 23
						data = /*line DVpUieavcx.go:1*/ append(data, "\x83|\x83\x84"...,
						)
					case 8:
						data = /*line DVpUieavcx.go:1*/ append(data, "a\x8a\xad\xa3"...,
						)
						i = 27
					case 19:
						i = 11
						data = /*line DVpUieavcx.go:1*/ append(data, "i%lv"...,
						)
					case 1:
						i = 33
						data = /*line DVpUieavcx.go:1*/ append(data, 155)
					case 23:
						i = 15
						data = /*line DVpUieavcx.go:1*/ append(data, "szq-"...,
						)
					case 26:
						i = 25
						data = /*line DVpUieavcx.go:1*/ append(data, 129)
					case 27:
						data = /*line DVpUieavcx.go:1*/ append(data, 152)
						i = 1
					case 32:
						i = 20
						data = /*line DVpUieavcx.go:1*/ append(data, "E5"...,
						)
					case 18:
						i = 3
						data = /*line DVpUieavcx.go:1*/ append(data, "\x94\x88<\x91"...,
						)
					case 6:
						data = /*line DVpUieavcx.go:1*/ append(data, "\xe8\xdd\xf0"...,
						)
						i = 17
					case 12:
						data = /*line DVpUieavcx.go:1*/ append(data, 126)
						i = 19
					case 10:
						i = 4
						data = /*line DVpUieavcx.go:1*/ append(data, 151)
					case 11:
						data = /*line DVpUieavcx.go:1*/ append(data, "r!p"...,
						)
						i = 9
					case 25:
						i = 0
						data = /*line DVpUieavcx.go:1*/ append(data, "wy"...,
						)
					case 14:
						data = /*line DVpUieavcx.go:1*/ append(data, "\x93\x83\x95"...,
						)
						i = 16
					case 22:
						data = /*line DVpUieavcx.go:1*/ append(data, "G\x8d\x90"...,
						)
						i = 18
					case 28:
						data = /*line DVpUieavcx.go:1*/ append(data, "\x9cX"...,
						)
						i = 8
					case 5:
						for y := range data {
							data[y] = data[y] - /*line DVpUieavcx.go:1*/ byte(decryptKey^y)
						}
						i = 21
					case 0:
						data = /*line DVpUieavcx.go:1*/ append(data, 121)
						i = 12
					case 29:
						i = 13
						data = /*line DVpUieavcx.go:1*/ append(data, "\x97\xa0"...,
						)
					case 17:
						i = 24
						data = /*line DVpUieavcx.go:1*/ append(data, "\x9a\xa0"...,
						)
					case 2:
						i = 30
						data = /*line DVpUieavcx.go:1*/ append(data, "Kk"...,
						)
					case 9:
						data = /*line DVpUieavcx.go:1*/ append(data, "h\xf3\x9d\xf4"...,
						)
						i = 6
					case 24:
						i = 5
						data = /*line DVpUieavcx.go:1*/ append(data, 234)
					}
				}
				return /*line DVpUieavcx.go:1*/ string(data)
			}(), /*line ELfRNmRJSid.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.Ig3RQ9F)

			var cbz6DUSD types.View

			for _, iv5UGEus9T := range cu1RSg2YR.detectedTmos {
				/*line TazuPg3zy.go:1*/ cu1RSg2YR.pm.ProcessRemoteTmo(iv5UGEus9T)

				if cbz6DUSD < iv5UGEus9T.BcEyAIp_ {
					cbz6DUSD = iv5UGEus9T.BcEyAIp_
				}
			}

			if /*line RznaQTvh.go:1*/ cu1RSg2YR.pm.GetAnchorView() < cbz6DUSD {
				/*line ilE96sKE1.go:1*/ cu1RSg2YR.pm.UpdateAnchorView(cbz6DUSD)
			}

			/*line Hx57XbE06Vsa.go:1*/
			cu1RSg2YR.pm.ExecuteViewChange( /*line FQ_fA50CNrAK.go:1*/ cu1RSg2YR.pm.GetNewView())

			cu1RSg2YR.detectedTmos = /*line tA7xK3zS15D.go:1*/ make(map[types.NodeID]*pacemaker.TMO)
		}
		return
	}

	if /*line De35UZRR.go:1*/ len(cu1RSg2YR.detectedTmos) > 0 {
		cu1RSg2YR.detectedTmos = /*line PCNdsMzm6G.go:1*/ make(map[types.NodeID]*pacemaker.TMO)
	}

	if iv5UGEus9T.BcEyAIp_ > /*line UCX4FiHDuIgE.go:1*/ cu1RSg2YR.pm.GetAnchorView() {
		/*line IDfUiqavj4pS.go:1*/ cu1RSg2YR.pm.UpdateAnchorView(iv5UGEus9T.BcEyAIp_)
		/*line P2bFSWm844_P.go:1*/ cu1RSg2YR.pm.ExecuteViewChange( /*line IYrmbTRSJ.go:1*/ cu1RSg2YR.pm.GetNewView())
		return
	}

	qNDcZn, aiM0S8js, hByvRzR := /*line VXOH_VlexE.go:1*/ cu1RSg2YR.pm.ProcessRemoteTmo(iv5UGEus9T)
	if !qNDcZn {
		return
	}

	aiM0S8js.Epoch = /*line MPkrep60.go:1*/ cu1RSg2YR.pm.GetCurEpoch()

	if /*line f0G5ATu.go:1*/ cu1RSg2YR.IsLeader( /*line IUA8peLmpfl.go:1*/ cu1RSg2YR.ID(), aiM0S8js.NewView, aiM0S8js.Epoch) {
		return
	}

	if hByvRzR != nil {
		cu1RSg2YR.reservedBlock <- hByvRzR
	}

	/*line JIn8u0Td0Gm.go:1*/
	log.Debugf(func() /*line HNYDejhhHY_a.go:1*/ string {
		seed := /*line HNYDejhhHY_a.go:1*/ byte(86)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line HNYDejhhHY_a.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line HNYDejhhHY_a.go:1*/ fnc(13)(70)(223)(213)(125)(242)(156)(26)(237)(12)(30)(234)(240)(33)(241)(232)(2)(27)(239)(45)(203)(30)(166)(21)(43)(85)(190)(235)(83)(175)(6)(91)(180)(255)(224)(5)(26)(168)(86)(233)(29)(172)(118)(203)(14)(167)(120)(207)(16)(242)(87)(235)(207)
		return /*line HNYDejhhHY_a.go:1*/ string(data)
	}(), /*line B2QmluUanr.go:1*/ cu1RSg2YR.ID(), aiM0S8js.NewView)

	aiM0S8js.Uc9aMd = /*line ANTDdalSn.go:1*/ cu1RSg2YR.ID()

	/*line WFa5PN6r.go:1*/
	cu1RSg2YR.SetRole(types.LEADER)

	/*line FcCLgAzO.go:1*/
	log.Debugf(func() /*line bgehVv.go:1*/ string {
		seed := /*line bgehVv.go:1*/ byte(50)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line bgehVv.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line bgehVv.go:1*/ fnc(41)(202)(81)(231)(195)(8)(40)(34)(253)(244)(2)(14)(0)(223)(19)(8)(2)(5)(241)(239)(25)(2)(186)(247)(70)(3)(5)(251)(10)(245)(253)(255)(188)(80)(2)(253)(244)(2)(14)(0)(246)(5)(249)(185)(82)(243)(8)(2)(5)(241)(187)(84)(245)(4)(248)(10)(6)(255)(172)(70)(12)(253)(254)(179)(5)(81)(170)(70)(9)(3)(174)(78)(247)(18)(255)(243)(252)(18)(169)(5)(81)
		return /*line bgehVv.go:1*/ string(data)
	}(), /*line Iq856zVUA.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.Ln9AEaadIIY, iv5UGEus9T.Ig3RQ9F)

	/*line Alkp_Pa.go:1*/
	cu1RSg2YR.Broadcast(aiM0S8js)
	/*line S5OBOX.go:1*/ cu1RSg2YR.ProcessTC(aiM0S8js)
}

func (cu1RSg2YR *PBFT) ProcessLocalTmo(a3qJWd types.View) {
	/*line OuDYjFath.go:1*/ log.Debugf(func() /*line I7xsNzTaGz.go:1*/ string {
		seed := /*line I7xsNzTaGz.go:1*/ byte(84)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line I7xsNzTaGz.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line I7xsNzTaGz.go:1*/ fnc(7)(202)(81)(231)(195)(8)(40)(34)(253)(244)(2)(14)(0)(217)(35)(244)(254)(11)(232)(25)(2)(186)(247)(80)(2)(253)(244)(2)(14)(0)(246)(5)(249)(185)(76)(3)(244)(254)(11)(180)(84)(245)(4)(248)(10)(6)(255)(172)(70)(9)(3)(174)(86)(243)(252)(18)(169)(5)(81)
		return /*line I7xsNzTaGz.go:1*/ string(data)
	}(), /*line U_WL5m.go:1*/ cu1RSg2YR.ID(), a3qJWd)

	/*line EqYhkr8nM.go:1*/
	cu1RSg2YR.SetState(types.VIEWCHANGING)

	iv5UGEus9T := &pacemaker.TMO{
		Shard:/*line Hi_FamMz9M.go:1*/ cu1RSg2YR.Shard(),
		Epoch:/*line d0uej6Lz.go:1*/ cu1RSg2YR.pm.GetCurEpoch(),
		QmJZRap: a3qJWd,
		Ig3RQ9F:/*line wu1Fn5jDkilz.go:1*/ cu1RSg2YR.pm.GetNewView(),
		BcEyAIp_:/*line At1jm6gAoG.go:1*/ cu1RSg2YR.pm.GetAnchorView(),
		Ln9AEaadIIY:/*line t0dyfF.go:1*/ cu1RSg2YR.ID(),
		UFs51prD6H1:/*line fej1Paqqn5.go:1*/ cu1RSg2YR.GetHighBlockHeight(),
		K13ceE04nX:/*line YwXeY6fUmsd.go:1*/ cu1RSg2YR.GetHighBlockHeight(),
		N9CdZ7bx2WlQ:/*line EsEmPvRQNQlv.go:1*/ cu1RSg2YR.bc.GetBlockByBlockHeight( /*line Euuhyq.go:1*/ cu1RSg2YR.GetHighBlockHeight()).(*blockchain.WorkerBlock),
		GtrfyVMK9: nil,
		W52craOOoraa:/*line v2QKO7a_I5bY.go:1*/ cu1RSg2YR.GetHighQC(),
	}

	if aZWGTB, kbgkwuEjk := cu1RSg2YR.agreeingBlocks[iv5UGEus9T.Epoch][iv5UGEus9T.QmJZRap][iv5UGEus9T.UFs51prD6H1+1]; kbgkwuEjk {
		iv5UGEus9T.K13ceE04nX += 1
		iv5UGEus9T.GtrfyVMK9 = aZWGTB.WorkerBlock
	}

	/*line bwl8bY.go:1*/
	log.Debugf(func() /*line DClTcsF8.go:1*/ string {
		seed := /*line DClTcsF8.go:1*/ byte(192)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line DClTcsF8.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line DClTcsF8.go:1*/ fnc(155)(202)(81)(231)(195)(8)(40)(34)(253)(244)(2)(14)(0)(217)(35)(244)(254)(11)(232)(25)(2)(186)(247)(84)(249)(2)(177)(73)(10)(173)(66)(19)(244)(3)(8)(172)(70)(9)(3)(174)(86)(243)(252)(18)(169)(5)(81)(170)(78)(247)(18)(169)(86)(243)(252)(18)(169)(5)(81)(170)(65)(13)(245)(5)(7)(3)(174)(86)(243)(252)(18)(169)(5)(81)(170)(76)(245)(18)(1)(172)(66)(10)(3)(244)(8)(253)(253)(4)(254)(1)(12)(172)(5)(81)(170)(80)(2)(243)(11)(241)(17)(243)(255)(188)(66)(10)(3)(244)(8)(253)(253)(4)(254)(1)(12)(172)(5)(81)
		return /*line DClTcsF8.go:1*/ string(data)
	}(), /*line SKnwGxU.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.QmJZRap, iv5UGEus9T.Ig3RQ9F, iv5UGEus9T.BcEyAIp_, iv5UGEus9T.UFs51prD6H1, iv5UGEus9T.K13ceE04nX)

	/*line QuxL_oInx.go:1*/
	cu1RSg2YR.BroadcastToSome( /*line FSoa90Vi7B.go:1*/ cu1RSg2YR.FindCommitteesFor(iv5UGEus9T.Epoch), iv5UGEus9T)
	/*line zrhkJN.go:1*/ cu1RSg2YR.ProcessRemoteTmo(iv5UGEus9T)

	/*line JQuVMPXSbKF.go:1*/
	log.Debugf(func() /*line drXa6hD.go:1*/ string {
		key := [] /*line drXa6hD.go:1*/ byte("m%2\xbf\xfc<\xd5\xeeM\xc5\xe3\xa3$'\xf8\x87(\xefe(\x94OA\xd8P\x8eb\x01\x18\xf9\ud7adA\x1eu\xa1\xdd\x06\x89m\x95\x05O\xab\x90\x03\xf8\x9a\x86)O\xa7\r\v\xeev}\xc8S$\xdaB\xb7\xfc\x16\xbb\xc0")
		data := [] /*line drXa6hD.go:1*/ byte("\xee\x00D\x9e$\xec{\x84\"\x9e\x82\xd0O%w\xdc9}\xefE\xdb\xdaߎ\x19\xe0\arPlw\x82\xc31Q\xeeĖm\xe0\x01\xd2\x1b\x1d\xc4\xd3^t\x86\xee@\x1e\xbebj\x86\xaa\xe9\xa7\x1f\xfc\x9c'\xae{\nj\xb6")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line drXa6hD.go:1*/ string(data)
	}(), /*line ya_CLRwQi8.go:1*/ cu1RSg2YR.ID(), a3qJWd)
}

func (cu1RSg2YR *PBFT) ProcessTC(aiM0S8js *pacemaker.TC) {
	/*line FWpH8zkY43.go:1*/ log.Debugf(func() /*line GIMYbvQna.go:1*/ string {
		fullData := [] /*line GIMYbvQna.go:1*/ byte("\x99@\x1e\x01\xffAS\v\t\xfd+)6\xc1f\xb0Vfq\x90M\xcdOE\xee\xc5ȣ\x02UZ\x8b\x02\x8f\xc1Ф\x95=P\x10\xa1xEUi;\x1f@\x95\x91qX\xd6)\x1d,ӈ\x98\xab\x13l\xca\xdb-\xba#\xaa\xe3\x16\xd3:}<PmM)̧\x16\xef#\x12?4CHT!\xc3ҙ3\xce\x04\x01\xa3\xaf\xc3R+\x1dw\xff\xbf3\x8b\x1e\x93\xce\xc9\xf0\xc2Du\xccg\xbc \xe8\xea\xe6$s\xacb\xd1r")
		idxKey := [] /*line GIMYbvQna.go:1*/ byte("\xa3\x1f 2\x8al\x808`\x97\x83\x89\nN\xb3\x15")
		data := /*line GIMYbvQna.go:1*/ make([]byte, 0, 66)
		data = /*line GIMYbvQna.go:1*/ append(data, fullData[126^ /*line GIMYbvQna.go:1*/ int(idxKey[8])]+fullData[99^ /*line GIMYbvQna.go:1*/ int(idxKey[8])], fullData[124^ /*line GIMYbvQna.go:1*/ int(idxKey[13])]-fullData[112^ /*line GIMYbvQna.go:1*/ int(idxKey[13])], fullData[73^ /*line GIMYbvQna.go:1*/ int(idxKey[5])]-fullData[67^ /*line GIMYbvQna.go:1*/ int(idxKey[5])], fullData[33^ /*line GIMYbvQna.go:1*/ int(idxKey[2])]+fullData[23^ /*line GIMYbvQna.go:1*/ int(idxKey[2])], fullData[145^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]+fullData[130^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[194^ /*line GIMYbvQna.go:1*/ int(idxKey[0])]^fullData[237^ /*line GIMYbvQna.go:1*/ int(idxKey[0])], fullData[105^ /*line GIMYbvQna.go:1*/ int(idxKey[12])]^fullData[14^ /*line GIMYbvQna.go:1*/ int(idxKey[12])], fullData[50^ /*line GIMYbvQna.go:1*/ int(idxKey[8])]-fullData[41^ /*line GIMYbvQna.go:1*/ int(idxKey[8])], fullData[27^ /*line GIMYbvQna.go:1*/ int(idxKey[5])]-fullData[120^ /*line GIMYbvQna.go:1*/ int(idxKey[5])], fullData[173^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]+fullData[215^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[78^ /*line GIMYbvQna.go:1*/ int(idxKey[1])]^fullData[98^ /*line GIMYbvQna.go:1*/ int(idxKey[1])], fullData[186^ /*line GIMYbvQna.go:1*/ int(idxKey[4])]^fullData[212^ /*line GIMYbvQna.go:1*/ int(idxKey[4])], fullData[18^ /*line GIMYbvQna.go:1*/ int(idxKey[7])]^fullData[63^ /*line GIMYbvQna.go:1*/ int(idxKey[7])], fullData[136^ /*line GIMYbvQna.go:1*/ int(idxKey[10])]+fullData[137^ /*line GIMYbvQna.go:1*/ int(idxKey[10])], fullData[203^ /*line GIMYbvQna.go:1*/ int(idxKey[14])]+fullData[240^ /*line GIMYbvQna.go:1*/ int(idxKey[14])], fullData[134^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]-fullData[177^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[240^ /*line GIMYbvQna.go:1*/ int(idxKey[10])]-fullData[255^ /*line GIMYbvQna.go:1*/ int(idxKey[10])], fullData[233^ /*line GIMYbvQna.go:1*/ int(idxKey[0])]-fullData[214^ /*line GIMYbvQna.go:1*/ int(idxKey[0])], fullData[114^ /*line GIMYbvQna.go:1*/ int(idxKey[15])]+fullData[57^ /*line GIMYbvQna.go:1*/ int(idxKey[15])], fullData[59^ /*line GIMYbvQna.go:1*/ int(idxKey[8])]-fullData[57^ /*line GIMYbvQna.go:1*/ int(idxKey[8])], fullData[25^ /*line GIMYbvQna.go:1*/ int(idxKey[2])]^fullData[47^ /*line GIMYbvQna.go:1*/ int(idxKey[2])], fullData[32^ /*line GIMYbvQna.go:1*/ int(idxKey[2])]+fullData[111^ /*line GIMYbvQna.go:1*/ int(idxKey[2])], fullData[226^ /*line GIMYbvQna.go:1*/ int(idxKey[6])]^fullData[163^ /*line GIMYbvQna.go:1*/ int(idxKey[6])], fullData[128^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]-fullData[203^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[35^ /*line GIMYbvQna.go:1*/ int(idxKey[7])]-fullData[112^ /*line GIMYbvQna.go:1*/ int(idxKey[7])], fullData[37^ /*line GIMYbvQna.go:1*/ int(idxKey[13])]-fullData[87^ /*line GIMYbvQna.go:1*/ int(idxKey[13])], fullData[25^ /*line GIMYbvQna.go:1*/ int(idxKey[7])]^fullData[65^ /*line GIMYbvQna.go:1*/ int(idxKey[7])], fullData[244^ /*line GIMYbvQna.go:1*/ int(idxKey[6])]+fullData[188^ /*line GIMYbvQna.go:1*/ int(idxKey[6])], fullData[9^ /*line GIMYbvQna.go:1*/ int(idxKey[13])]+fullData[103^ /*line GIMYbvQna.go:1*/ int(idxKey[13])], fullData[154^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]-fullData[163^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[24^ /*line GIMYbvQna.go:1*/ int(idxKey[2])]+fullData[37^ /*line GIMYbvQna.go:1*/ int(idxKey[2])], fullData[4^ /*line GIMYbvQna.go:1*/ int(idxKey[2])]+fullData[2^ /*line GIMYbvQna.go:1*/ int(idxKey[2])], fullData[23^ /*line GIMYbvQna.go:1*/ int(idxKey[12])]-fullData[113^ /*line GIMYbvQna.go:1*/ int(idxKey[12])], fullData[83^ /*line GIMYbvQna.go:1*/ int(idxKey[5])]^fullData[6^ /*line GIMYbvQna.go:1*/ int(idxKey[5])], fullData[161^ /*line GIMYbvQna.go:1*/ int(idxKey[14])]-fullData[186^ /*line GIMYbvQna.go:1*/ int(idxKey[14])], fullData[187^ /*line GIMYbvQna.go:1*/ int(idxKey[4])]+fullData[230^ /*line GIMYbvQna.go:1*/ int(idxKey[4])], fullData[40^ /*line GIMYbvQna.go:1*/ int(idxKey[13])]^fullData[22^ /*line GIMYbvQna.go:1*/ int(idxKey[13])], fullData[96^ /*line GIMYbvQna.go:1*/ int(idxKey[13])]-fullData[123^ /*line GIMYbvQna.go:1*/ int(idxKey[13])], fullData[206^ /*line GIMYbvQna.go:1*/ int(idxKey[4])]+fullData[144^ /*line GIMYbvQna.go:1*/ int(idxKey[4])], fullData[220^ /*line GIMYbvQna.go:1*/ int(idxKey[0])]^fullData[229^ /*line GIMYbvQna.go:1*/ int(idxKey[0])], fullData[226^ /*line GIMYbvQna.go:1*/ int(idxKey[4])]^fullData[136^ /*line GIMYbvQna.go:1*/ int(idxKey[4])], fullData[36^ /*line GIMYbvQna.go:1*/ int(idxKey[3])]^fullData[4^ /*line GIMYbvQna.go:1*/ int(idxKey[3])], fullData[144^ /*line GIMYbvQna.go:1*/ int(idxKey[6])]^fullData[213^ /*line GIMYbvQna.go:1*/ int(idxKey[6])], fullData[102^ /*line GIMYbvQna.go:1*/ int(idxKey[3])]^fullData[1^ /*line GIMYbvQna.go:1*/ int(idxKey[3])], fullData[81^ /*line GIMYbvQna.go:1*/ int(idxKey[13])]^fullData[52^ /*line GIMYbvQna.go:1*/ int(idxKey[13])], fullData[15^ /*line GIMYbvQna.go:1*/ int(idxKey[3])]^fullData[68^ /*line GIMYbvQna.go:1*/ int(idxKey[3])], fullData[178^ /*line GIMYbvQna.go:1*/ int(idxKey[11])]+fullData[9^ /*line GIMYbvQna.go:1*/ int(idxKey[11])], fullData[22^ /*line GIMYbvQna.go:1*/ int(idxKey[12])]+fullData[70^ /*line GIMYbvQna.go:1*/ int(idxKey[12])], fullData[206^ /*line GIMYbvQna.go:1*/ int(idxKey[0])]+fullData[232^ /*line GIMYbvQna.go:1*/ int(idxKey[0])], fullData[101^ /*line GIMYbvQna.go:1*/ int(idxKey[2])]^fullData[68^ /*line GIMYbvQna.go:1*/ int(idxKey[2])], fullData[175^ /*line GIMYbvQna.go:1*/ int(idxKey[0])]^fullData[132^ /*line GIMYbvQna.go:1*/ int(idxKey[0])], fullData[5^ /*line GIMYbvQna.go:1*/ int(idxKey[5])]-fullData[127^ /*line GIMYbvQna.go:1*/ int(idxKey[5])], fullData[162^ /*line GIMYbvQna.go:1*/ int(idxKey[11])]+fullData[200^ /*line GIMYbvQna.go:1*/ int(idxKey[11])], fullData[191^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]-fullData[230^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[82^ /*line GIMYbvQna.go:1*/ int(idxKey[1])]+fullData[69^ /*line GIMYbvQna.go:1*/ int(idxKey[1])], fullData[51^ /*line GIMYbvQna.go:1*/ int(idxKey[5])]-fullData[65^ /*line GIMYbvQna.go:1*/ int(idxKey[5])], fullData[215^ /*line GIMYbvQna.go:1*/ int(idxKey[4])]^fullData[146^ /*line GIMYbvQna.go:1*/ int(idxKey[4])], fullData[111^ /*line GIMYbvQna.go:1*/ int(idxKey[12])]+fullData[101^ /*line GIMYbvQna.go:1*/ int(idxKey[12])], fullData[104^ /*line GIMYbvQna.go:1*/ int(idxKey[8])]-fullData[14^ /*line GIMYbvQna.go:1*/ int(idxKey[8])], fullData[208^ /*line GIMYbvQna.go:1*/ int(idxKey[6])]+fullData[242^ /*line GIMYbvQna.go:1*/ int(idxKey[6])], fullData[28^ /*line GIMYbvQna.go:1*/ int(idxKey[5])]^fullData[18^ /*line GIMYbvQna.go:1*/ int(idxKey[5])], fullData[193^ /*line GIMYbvQna.go:1*/ int(idxKey[9])]^fullData[192^ /*line GIMYbvQna.go:1*/ int(idxKey[9])], fullData[132^ /*line GIMYbvQna.go:1*/ int(idxKey[4])]+fullData[200^ /*line GIMYbvQna.go:1*/ int(idxKey[4])], fullData[42^ /*line GIMYbvQna.go:1*/ int(idxKey[12])]+fullData[89^ /*line GIMYbvQna.go:1*/ int(idxKey[12])], fullData[1^ /*line GIMYbvQna.go:1*/ int(idxKey[6])]^fullData[224^ /*line GIMYbvQna.go:1*/ int(idxKey[6])])
		return /*line GIMYbvQna.go:1*/ string(data)
	}(), /*line dmOOniCDE9.go:1*/ cu1RSg2YR.ID(), aiM0S8js.NewView)

	/*line KIJLFyHB.go:1*/
	cu1RSg2YR.SetState(types.READY)

	hYhbNJBbO, uxNBF8Z7 := /*line WrdiA0Pobg.go:1*/ cu1RSg2YR.pm.GetCurEpochView()

	/*line YBbXLX7.go:1*/
	delete(cu1RSg2YR.agreeingBlocks[hYhbNJBbO][uxNBF8Z7], aiM0S8js.U6G7XuY_)

	/*line XayACf6BG.go:1*/
	cu1RSg2YR.pm.UpdateView(aiM0S8js.NewView - 1)
	/*line BVum__ERGX.go:1*/ cu1RSg2YR.pm.UpdateAnchorView(aiM0S8js.Z7uwaNa)

	/*line EFjFTOa.go:1*/
	log.Debugf(func() /*line d7fOqRzk0.go:1*/ string {
		fullData := [] /*line d7fOqRzk0.go:1*/ byte("\v\x0fe\xd7G\x02t\x95\xea5i\a\xbb\x96\xe4\xf2\xac\xc1\x1ck=\b\x96\xff\xa09\x03B\xbb\x1a\x9d\x94Ј(\x9a\xc9b\x92P\xa9\x9e\xf6Z1\xe5\xb4%\x14\x93D\\W\x86l~\xecs>\xfc^\xa2L\x1d\xe5\x0f\b\x7f\xec\x0e)\xf7\x944ɱJi\x82'\xe2\xb3}/\xf1p\xc1)0\xa3\xea\xb1*c\xc9\x16\xe3\x13\xb7\xc1\xc0\xc4 \xb2\x1fM\xb7\xe3s\xd3=&\xfd\xf8\xb4\xd4M\x8ba\x97\x93\xf5~\x98\x91\xb9\xd7`D\xbbHDH)v{v\x8d(_\xe8{\xe8r\t\xdbLJ")
		idxKey := [] /*line d7fOqRzk0.go:1*/ byte("\xa2\x15MG\x94\xdd\x1c\x9aS\xf0\xc37Um\x19t")
		data := /*line d7fOqRzk0.go:1*/ make([]byte, 0, 75)
		data = /*line d7fOqRzk0.go:1*/ append(data, fullData[24^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]^fullData[91^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[119^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]+fullData[96^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[99^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]-fullData[119^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[58^ /*line d7fOqRzk0.go:1*/ int(idxKey[6])]-fullData[21^ /*line d7fOqRzk0.go:1*/ int(idxKey[6])], fullData[201^ /*line d7fOqRzk0.go:1*/ int(idxKey[5])]+fullData[189^ /*line d7fOqRzk0.go:1*/ int(idxKey[5])], fullData[77^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]-fullData[38^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[239^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])]-fullData[112^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])], fullData[129^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]^fullData[44^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[110^ /*line d7fOqRzk0.go:1*/ int(idxKey[12])]-fullData[220^ /*line d7fOqRzk0.go:1*/ int(idxKey[12])], fullData[211^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])]^fullData[220^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])], fullData[5^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])]-fullData[12^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])], fullData[111^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])]-fullData[98^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])], fullData[55^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]^fullData[2^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[147^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]+fullData[179^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[58^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])]^fullData[23^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])], fullData[208^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]^fullData[188^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[113^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])]-fullData[14^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])], fullData[28^ /*line d7fOqRzk0.go:1*/ int(idxKey[3])]^fullData[57^ /*line d7fOqRzk0.go:1*/ int(idxKey[3])], fullData[85^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]+fullData[7^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[101^ /*line d7fOqRzk0.go:1*/ int(idxKey[11])]-fullData[118^ /*line d7fOqRzk0.go:1*/ int(idxKey[11])], fullData[167^ /*line d7fOqRzk0.go:1*/ int(idxKey[11])]+fullData[72^ /*line d7fOqRzk0.go:1*/ int(idxKey[11])], fullData[239^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]^fullData[191^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[116^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])]+fullData[150^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])], fullData[72^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])]^fullData[102^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])], fullData[80^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]-fullData[63^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[131^ /*line d7fOqRzk0.go:1*/ int(idxKey[5])]-fullData[245^ /*line d7fOqRzk0.go:1*/ int(idxKey[5])], fullData[161^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])]^fullData[190^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])], fullData[41^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])]^fullData[68^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])], fullData[65^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])]+fullData[140^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])], fullData[144^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])]+fullData[134^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])], fullData[206^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])]+fullData[25^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])], fullData[122^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])]^fullData[3^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])], fullData[159^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]-fullData[241^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[180^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])]^fullData[233^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])], fullData[20^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])]-fullData[59^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])], fullData[17^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])]^fullData[45^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])], fullData[180^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]-fullData[36^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[45^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])]+fullData[38^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])], fullData[144^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]-fullData[51^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[92^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])]^fullData[68^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])], fullData[40^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]^fullData[35^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[168^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])]-fullData[166^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])], fullData[119^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])]^fullData[229^ /*line d7fOqRzk0.go:1*/ int(idxKey[13])], fullData[45^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]+fullData[167^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[211^ /*line d7fOqRzk0.go:1*/ int(idxKey[7])]-fullData[170^ /*line d7fOqRzk0.go:1*/ int(idxKey[7])], fullData[143^ /*line d7fOqRzk0.go:1*/ int(idxKey[6])]^fullData[153^ /*line d7fOqRzk0.go:1*/ int(idxKey[6])], fullData[37^ /*line d7fOqRzk0.go:1*/ int(idxKey[6])]^fullData[67^ /*line d7fOqRzk0.go:1*/ int(idxKey[6])], fullData[174^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]+fullData[192^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[112^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])]^fullData[0^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])], fullData[179^ /*line d7fOqRzk0.go:1*/ int(idxKey[5])]-fullData[174^ /*line d7fOqRzk0.go:1*/ int(idxKey[5])], fullData[7^ /*line d7fOqRzk0.go:1*/ int(idxKey[3])]-fullData[4^ /*line d7fOqRzk0.go:1*/ int(idxKey[3])], fullData[194^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])]+fullData[232^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])], fullData[238^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]^fullData[245^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[168^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]-fullData[183^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[59^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]+fullData[44^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[200^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]+fullData[197^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])], fullData[17^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]-fullData[204^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[45^ /*line d7fOqRzk0.go:1*/ int(idxKey[1])]-fullData[34^ /*line d7fOqRzk0.go:1*/ int(idxKey[1])], fullData[73^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])]^fullData[129^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])], fullData[193^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]+fullData[55^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[86^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])]-fullData[32^ /*line d7fOqRzk0.go:1*/ int(idxKey[2])], fullData[181^ /*line d7fOqRzk0.go:1*/ int(idxKey[7])]+fullData[238^ /*line d7fOqRzk0.go:1*/ int(idxKey[7])], fullData[198^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])]^fullData[206^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])], fullData[216^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]^fullData[127^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[21^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])]^fullData[193^ /*line d7fOqRzk0.go:1*/ int(idxKey[8])], fullData[104^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])]-fullData[247^ /*line d7fOqRzk0.go:1*/ int(idxKey[15])], fullData[69^ /*line d7fOqRzk0.go:1*/ int(idxKey[12])]+fullData[83^ /*line d7fOqRzk0.go:1*/ int(idxKey[12])], fullData[41^ /*line d7fOqRzk0.go:1*/ int(idxKey[12])]+fullData[120^ /*line d7fOqRzk0.go:1*/ int(idxKey[12])], fullData[192^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])]-fullData[181^ /*line d7fOqRzk0.go:1*/ int(idxKey[4])], fullData[133^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])]-fullData[159^ /*line d7fOqRzk0.go:1*/ int(idxKey[9])], fullData[103^ /*line d7fOqRzk0.go:1*/ int(idxKey[11])]+fullData[48^ /*line d7fOqRzk0.go:1*/ int(idxKey[11])], fullData[62^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])]^fullData[76^ /*line d7fOqRzk0.go:1*/ int(idxKey[14])], fullData[255^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])]^fullData[68^ /*line d7fOqRzk0.go:1*/ int(idxKey[10])], fullData[177^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])]-fullData[219^ /*line d7fOqRzk0.go:1*/ int(idxKey[0])])
		return /*line d7fOqRzk0.go:1*/ string(data)
	}(), /*line Pva3oJOIN1m.go:1*/ cu1RSg2YR.ID(), aiM0S8js.NewView)

	hYhbNJBbO, uxNBF8Z7 = /*line AroUWt.go:1*/ cu1RSg2YR.pm.GetCurEpochView()

	if ouJpb2aF3mlv, kbgkwuEjk := cu1RSg2YR.bufferedBlocks[hYhbNJBbO][uxNBF8Z7][aiM0S8js.U6G7XuY_]; kbgkwuEjk {
		_ = /*line a2AvdJyTF.go:1*/ cu1RSg2YR.ProcessWorkerBlock(ouJpb2aF3mlv)
		/*line uep_wmcN.go:1*/ delete(cu1RSg2YR.bufferedBlocks[hYhbNJBbO][uxNBF8Z7], aiM0S8js.U6G7XuY_)
	}
}

func (cu1RSg2YR *PBFT) UpdateSCT(kDYQNbksZq []*message.Transaction, thxhyC_T9ga []*blockchain.LocalSnapshot) {

	for _, mLro608nVKLl := range kDYQNbksZq {
		pqKOi5u0pS := 0
		if mLro608nVKLl.TXType == types.TRANSFER {
			if /*line zkjLYNapzZ.go:1*/ utils.CalculateShardToSend([]common.Address{mLro608nVKLl.To} /*line uEn20jvI.go:1*/, config.GetConfig().ShardCount)[0] != /*line AsIAyrYpak.go:1*/ cu1RSg2YR.Shard() {
				for _, aVU16qtmHkO := range thxhyC_T9ga {
					if aVU16qtmHkO.Address == mLro608nVKLl.From {
						b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line I8FzTaIoq.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address /*line KeYM20z6.go:1*/, utils.SlotToKey(0), aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
						/*line HaBQ5r.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						pqKOi5u0pS++
						break
					}
				}
				if pqKOi5u0pS == 0 {
					scDQLKCT := mLro608nVKLl.From
					fsNx3bxqUIt := /*line RKS_FvaujVx.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(mLro608nVKLl.From).String()
					b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line SazLtx2OZ.go:1*/ cu1RSg2YR.CreateSCTObj(scDQLKCT /*line iiL6b_6B2y.go:1*/, utils.SlotToKey(0), fsNx3bxqUIt, []common.Hash{mLro608nVKLl.Hash})
					/*line JPBkaIQ.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
				}
			} else if /*line HXSzuaLaZl.go:1*/ utils.CalculateShardToSend([]common.Address{mLro608nVKLl.From} /*line aPOwxjHYCBo.go:1*/, config.GetConfig().ShardCount)[0] != /*line DNY6xK5ypyi.go:1*/ cu1RSg2YR.Shard() {
				for _, aVU16qtmHkO := range thxhyC_T9ga {
					if aVU16qtmHkO.Address == mLro608nVKLl.To {
						b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line X8F0qoqLDUF3.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address /*line y4R1Q4wdht.go:1*/, utils.SlotToKey(0), aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
						/*line FXb6hQG_9.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						pqKOi5u0pS++
						break
					}
				}
				if pqKOi5u0pS == 0 {
					scDQLKCT := mLro608nVKLl.To
					fsNx3bxqUIt := /*line tcJIiQM0.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(mLro608nVKLl.To).String()
					b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line ZMid8x.go:1*/ cu1RSg2YR.CreateSCTObj(scDQLKCT /*line lb2zK2zxgwi.go:1*/, utils.SlotToKey(0), fsNx3bxqUIt, []common.Hash{mLro608nVKLl.Hash})
					/*line UuHq4apaMx.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
				}
			}
		} else if mLro608nVKLl.TXType == types.SMARTCONTRACT {
			if /*line YIc7cpKN11q.go:1*/ utils.CalculateShardToSend([]common.Address{mLro608nVKLl.From} /*line GLvo4t_i.go:1*/, config.GetConfig().ShardCount)[0] == /*line YtHe56v.go:1*/ cu1RSg2YR.Shard() {
				for _, aVU16qtmHkO := range thxhyC_T9ga {
					if aVU16qtmHkO.Address == mLro608nVKLl.From {
						b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line ENQVWxyePTH.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address /*line KPGVy4cvOtQU.go:1*/, utils.SlotToKey(0), aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
						/*line XYZ9aAoy3s.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						pqKOi5u0pS++
						break
					}
				}
				if pqKOi5u0pS == 0 {
					scDQLKCT := mLro608nVKLl.From
					fsNx3bxqUIt := /*line Ad2n6CqLHm.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(mLro608nVKLl.From).String()
					b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line tDW1hJdGdz51.go:1*/ cu1RSg2YR.CreateSCTObj(scDQLKCT /*line WPVMjhUFvAx4.go:1*/, utils.SlotToKey(0), fsNx3bxqUIt, []common.Hash{mLro608nVKLl.Hash})
					/*line VGMRA5_A.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
				}
			}
			for _, ds58OBJNw := range mLro608nVKLl.RwSet {
				if /*line wmyzGJuaa2.go:1*/ utils.CalculateShardToSend([]common.Address{ds58OBJNw.Address} /*line wCHDQi8.go:1*/, config.GetConfig().ShardCount)[0] == /*line inb1kCrk5e.go:1*/ cu1RSg2YR.Shard() {
					for _, ceBkC2 := range ds58OBJNw.ReadSet {
						pqKOi5u0pS = 0
						for _, aVU16qtmHkO := range thxhyC_T9ga {
							if aVU16qtmHkO.Address == ds58OBJNw.Address && aVU16qtmHkO.Slot == /*line I0ygkXSdL.go:1*/ common.HexToHash(ceBkC2) {
								for _, OTC_a4FQrin1 := range aVU16qtmHkO.RTCS {
									if OTC_a4FQrin1 == mLro608nVKLl.Hash {
										pqKOi5u0pS++
										break
									}
								}
								if pqKOi5u0pS == 0 {
									b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line BAmWkIS29U6o.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address, aVU16qtmHkO.Slot, aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
									/*line ltZFLFW3fHj.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
									pqKOi5u0pS++
								}
								break
							}
						}
						if pqKOi5u0pS == 0 {
							scDQLKCT := ds58OBJNw.Address
							fsNx3bxqUIt := /*line LiXoiQQnBV.go:1*/ cu1RSg2YR.bc.GetStateDB().GetState(ds58OBJNw.Address /*line bXzNpjoiiN.go:1*/, common.HexToHash(ceBkC2)).String()
							b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line ei3Rl0X3OJK.go:1*/ cu1RSg2YR.CreateSCTObj(scDQLKCT /*line CHMjVtAu6Aa.go:1*/, common.HexToHash(ceBkC2), fsNx3bxqUIt, []common.Hash{mLro608nVKLl.Hash})
							/*line Gata44LpnDl5.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						}
					}
				}
			}
		}
	}
}

func (cu1RSg2YR *PBFT) ModifySCT(o9gw2jO []*message.Transaction) {
	for _, bQZ9a4xL := range o9gw2jO {
		if bQZ9a4xL.IsCrossShardTx {
			if bQZ9a4xL.TXType == types.TRANSFER {
				if fsNx3bxqUIt, i3KmQgjb := /*line NhdpIQ.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From]; i3KmQgjb {
					for mTZDwMx, sTaCTHWa4L := range fsNx3bxqUIt[ /*line _D8jpsJo.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ {
						if sTaCTHWa4L == bQZ9a4xL.Hash {
							fsNx3bxqUIt[ /*line Z9J3Ok.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ = /*line v1k__F.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx, fsNx3bxqUIt[ /*line W_33ym8b5X.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ)
							break
						}
					}
				} else if fsNx3bxqUIt, i3KmQgjb := /*line LAN9tyW.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.To]; i3KmQgjb {
					for mTZDwMx, sTaCTHWa4L := range fsNx3bxqUIt[ /*line FzDnIX2Zrn.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ {
						if sTaCTHWa4L == bQZ9a4xL.Hash {
							fsNx3bxqUIt[ /*line iGc5uDF4Jkv4.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ = /*line _K2zOH.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx, fsNx3bxqUIt[ /*line Ylv5d838L1X.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ)
							break
						}
					}
				}
			} else if bQZ9a4xL.TXType == types.SMARTCONTRACT {
				for _, mvaFahgWX := range bQZ9a4xL.RwSet {
					if fsNx3bxqUIt, i3KmQgjb := /*line Idxknyx_.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[mvaFahgWX.Address]; i3KmQgjb {
						for _, ceBkC2 := range mvaFahgWX.ReadSet {
							if ceBkC2 == func() /*line Gf_26U.go:1*/ string {
								data := [] /*line Gf_26U.go:1*/ byte("\x93\x8a\x8a.s\xcc\xcad\xce\xc8")
								positions := [...]byte{8, 1, 8, 0, 2, 9, 0, 9, 5, 8, 6, 6, 5, 0}
								for i := 0; i < 14; i += 2 {
									localKey := /*line Gf_26U.go:1*/ byte(i) + /*line Gf_26U.go:1*/ byte(positions[i]^positions[i+1]) + 82
									data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
								}
								return /*line Gf_26U.go:1*/ string(data)
							}() {
								for mTZDwMx, sTaCTHWa4L := range /*line YV6QKOS.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line Cda6zl.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ {
									if sTaCTHWa4L == bQZ9a4xL.Hash {
										/*line zrhsr2Nc.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line ItQOkHCv.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ = /*line A2dkxaxAyVa.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx /*line q1vdn6TXZ3CO.go:1*/, cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line Ib2joFUWdvSW.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ)
										for _, kmDgMw_b := range mvaFahgWX.WriteSet {
											if kmDgMw_b == func() /*line DP_A6gNSC.go:1*/ string {
												data := [] /*line DP_A6gNSC.go:1*/ byte("\x8fnm.s{\x9e\x94ef")
												positions := [...]byte{7, 6, 0, 1, 5, 9, 0, 0, 5, 2}
												for i := 0; i < 10; i += 2 {
													localKey := /*line DP_A6gNSC.go:1*/ byte(i) + /*line DP_A6gNSC.go:1*/ byte(positions[i]^positions[i+1]) + 249
													data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
												}
												return /*line DP_A6gNSC.go:1*/ string(data)
											}() {
												/*line EXiIo6jc0MWk.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line TMr2CGsca.go:1*/ utils.SlotToKey(0)].UPa0r5o = false
											}
										}
										break
									}
								}
							} else {
								if _, kbgkwuEjk := fsNx3bxqUIt[ /*line uGY_i4ouB77.go:1*/ common.HexToHash(ceBkC2)]; !kbgkwuEjk {
									fsNx3bxqUIt[ /*line mkBK9NsN.go:1*/ common.HexToHash(ceBkC2)] = &blockchain.SnapshotControlTable{}
								}
								for mTZDwMx, sTaCTHWa4L := range fsNx3bxqUIt[ /*line zOQAlEvY.go:1*/ common.HexToHash(ceBkC2)].BlC2JXSmAuvQ {
									if sTaCTHWa4L == bQZ9a4xL.Hash {
										fsNx3bxqUIt[ /*line hiUqxWh.go:1*/ common.HexToHash(ceBkC2)].BlC2JXSmAuvQ = /*line WOntRVuOi1.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx, fsNx3bxqUIt[ /*line GROwdxCkKam.go:1*/ common.HexToHash(ceBkC2)].BlC2JXSmAuvQ)
										break
									}
								}
							}
						}
					}
				}

				if /*line wYhig8yF6MW.go:1*/ utils.CalculateShardToSend([]common.Address{bQZ9a4xL.From} /*line NZj9ckUL6.go:1*/, config.GetConfig().ShardCount)[0] == /*line m6KbX1l9T.go:1*/ cu1RSg2YR.Shard() {
					if _, kbgkwuEjk := /*line KyVfXJA4FJ.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line jPZz6Ph.go:1*/ common.HexToHash("0")]; !kbgkwuEjk {
						/*line FAtYnwq_y.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From] = /*line S1r3F7y6.go:1*/ make(map[common.Hash]*blockchain.SnapshotControlTable)
						/*line IvjOvxMPdlb.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line kUnIcSMcp.go:1*/ common.HexToHash("0")] = &blockchain.SnapshotControlTable{}
					}
					for mTZDwMx, sTaCTHWa4L := range /*line CtawXNQmmRw.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line l7Onjtcsc.go:1*/ common.HexToHash("0")].BlC2JXSmAuvQ {
						if sTaCTHWa4L == bQZ9a4xL.Hash {
							/*line Av9TGyYml.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line F8tT1rhayxqp.go:1*/ common.HexToHash("0")].BlC2JXSmAuvQ = /*line ho1PFYoQRnCK.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx /*line Iq3qnkU.go:1*/, cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line VAcOfu512VHc.go:1*/ common.HexToHash("0")].BlC2JXSmAuvQ)
							break
						}
					}
				}
			}
		}
		for _, _4arcQEvNi := range bQZ9a4xL.RwSet {
			for _, kmDgMw_b := range _4arcQEvNi.WriteSet {
				if fsNx3bxqUIt, i3KmQgjb := /*line FEZ8tGnHK.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[_4arcQEvNi.Address][ /*line OoMDYaknv.go:1*/ common.HexToHash(kmDgMw_b)]; i3KmQgjb {
					fsNx3bxqUIt.UPa0r5o = false
				}
			}
			for _, ceBkC2 := range _4arcQEvNi.ReadSet {
				if fsNx3bxqUIt, i3KmQgjb := /*line waKi1oHwFKI.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[_4arcQEvNi.Address][ /*line SuXSWh46Z8s.go:1*/ common.HexToHash(ceBkC2)]; i3KmQgjb {
					fsNx3bxqUIt.UPa0r5o = false
				}
			}
		}

		if fsNx3bxqUIt, i3KmQgjb := /*line Goz9dh.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.From][ /*line uXgIGH2.go:1*/ common.HexToHash("0")]; i3KmQgjb {
			fsNx3bxqUIt.UPa0r5o = false
		}
		if bQZ9a4xL.TXType == types.TRANSFER {
			if fsNx3bxqUIt, i3KmQgjb := /*line IBqjgBt.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bQZ9a4xL.To][ /*line XF1HnJ5Eex.go:1*/ common.HexToHash("0")]; i3KmQgjb {
				fsNx3bxqUIt.UPa0r5o = false
			}
		}
	}
}

func (cu1RSg2YR *PBFT) CreateLocalSequence(Yw2ZXLA []*message.Transaction) []*message.Transaction {
	yxYJX82Xsq := /*line wuxvAOT.go:1*/ make([]*message.Transaction, 0)
	yxYJX82Xsq = /*line aGJsuViJ1a.go:1*/ append(yxYJX82Xsq, Yw2ZXLA...)

	for _, bBrNWS3ghg := range Yw2ZXLA {

		if bBrNWS3ghg.TXType == types.TRANSFER {
			if fsNx3bxqUIt, i3KmQgjb := /*line uQU0wraXsrtJ.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From]; i3KmQgjb {
				for mTZDwMx, sTaCTHWa4L := range fsNx3bxqUIt[ /*line EQep1rGt.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ {
					if sTaCTHWa4L == bBrNWS3ghg.Hash {
						fsNx3bxqUIt[ /*line IbgCPH.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ = /*line YaFxRp.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx, fsNx3bxqUIt[ /*line ClOteK_bk.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ)
						break
					}
				}
			} else if fsNx3bxqUIt, i3KmQgjb := /*line Di7_0jUqfbv.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.To]; i3KmQgjb {
				for mTZDwMx, sTaCTHWa4L := range fsNx3bxqUIt[ /*line FYH7mlHQC.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ {
					if sTaCTHWa4L == bBrNWS3ghg.Hash {
						fsNx3bxqUIt[ /*line AsyFasHZxs.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ = /*line HLv73jS.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx, fsNx3bxqUIt[ /*line BUc76xqI.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ)
						break
					}
				}
			}
		} else if bBrNWS3ghg.TXType == types.SMARTCONTRACT {
			for _, mvaFahgWX := range bBrNWS3ghg.RwSet {
				if fsNx3bxqUIt, i3KmQgjb := /*line UFUdrk.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[mvaFahgWX.Address]; i3KmQgjb {
					for _, ceBkC2 := range mvaFahgWX.ReadSet {
						if ceBkC2 == func() /*line YMHVwe.go:1*/ string {
							data := /*line YMHVwe.go:1*/ make([]byte, 0, 11)
							i := 6
							decryptKey := 7
							for counter := 0; i != 9; counter++ {
								decryptKey ^= i * counter
								switch i {
								case 1:
									data = /*line YMHVwe.go:1*/ append(data, 201)
									i = 8
								case 8:
									i = 5
									data = /*line YMHVwe.go:1*/ append(data, 213)
								case 6:
									i = 11
									data = /*line YMHVwe.go:1*/ append(data, 206)
								case 3:
									data = /*line YMHVwe.go:1*/ append(data, 202)
									i = 2
								case 5:
									i = 4
									data = /*line YMHVwe.go:1*/ append(data, 202)
								case 7:
									i = 10
									data = /*line YMHVwe.go:1*/ append(data, 218)
								case 2:
									data = /*line YMHVwe.go:1*/ append(data, 144)
									i = 0
								case 0:
									data = /*line YMHVwe.go:1*/ append(data, 216)
									i = 1
								case 4:
									data = /*line YMHVwe.go:1*/ append(data, 206)
									i = 7
								case 11:
									data = /*line YMHVwe.go:1*/ append(data, 211)
									i = 3
								case 10:
									i = 9
									for y := range data {
										data[y] = data[y] - /*line YMHVwe.go:1*/ byte(decryptKey^y)
									}
								}
							}
							return /*line YMHVwe.go:1*/ string(data)
						}() {
							for mTZDwMx, sTaCTHWa4L := range /*line C5YBXrr5ziUx.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line lbWaKl.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ {
								if sTaCTHWa4L == bBrNWS3ghg.Hash {
									/*line VSXEvyQ.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line kl_gIwL45.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ = /*line jYKWQoj9tS.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx /*line B5tqy2dlLw.go:1*/, cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line csJjc_wUS0.go:1*/ utils.SlotToKey(0)].BlC2JXSmAuvQ)
									for _, kmDgMw_b := range mvaFahgWX.WriteSet {
										if kmDgMw_b == func() /*line SML7hAfQE.go:1*/ string {
											key := [] /*line SML7hAfQE.go:1*/ byte("\xcb\xc3\x12\x92|*\xeaj\xa1R")
											data := [] /*line SML7hAfQE.go:1*/ byte("86y\xc0\xef\x8fX\xce\x06\xc4")
											for i, b := range key {
												data[i] = data[i] - b
											}
											return /*line SML7hAfQE.go:1*/ string(data)
										}() {
											/*line EWmWHE.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line YJmQ9zbN.go:1*/ utils.SlotToKey(0)].UPa0r5o = false
										}
									}
									break
								}
							}
						} else {
							if _, kbgkwuEjk := fsNx3bxqUIt[ /*line h5dwtpzw.go:1*/ common.HexToHash(ceBkC2)]; !kbgkwuEjk {
								fsNx3bxqUIt[ /*line Vh7Yl71oS.go:1*/ common.HexToHash(ceBkC2)] = &blockchain.SnapshotControlTable{}
							}
							for mTZDwMx, sTaCTHWa4L := range fsNx3bxqUIt[ /*line In36kX7FYp.go:1*/ common.HexToHash(ceBkC2)].BlC2JXSmAuvQ {
								if sTaCTHWa4L == bBrNWS3ghg.Hash {
									fsNx3bxqUIt[ /*line kvpCxheIpc.go:1*/ common.HexToHash(ceBkC2)].BlC2JXSmAuvQ = /*line Yn4gOJI4X.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx, fsNx3bxqUIt[ /*line HAD7td.go:1*/ common.HexToHash(ceBkC2)].BlC2JXSmAuvQ)
									break
								}
							}
						}
					}
				}
			}

			if /*line GSXow3z.go:1*/ utils.CalculateShardToSend([]common.Address{bBrNWS3ghg.From} /*line m31eWxDb.go:1*/, config.GetConfig().ShardCount)[0] == /*line BiVWRF1C.go:1*/ cu1RSg2YR.Shard() {
				if _, kbgkwuEjk := /*line _dlSSxuGzVU.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line KkHB8lHZn.go:1*/ common.HexToHash("0")]; !kbgkwuEjk {
					/*line JvrEFO.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From] = /*line NPlhFACbBYy2.go:1*/ make(map[common.Hash]*blockchain.SnapshotControlTable)
					/*line GlXKtU.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line WarzNvqB59v.go:1*/ common.HexToHash("0")] = &blockchain.SnapshotControlTable{}
				}
				for mTZDwMx, sTaCTHWa4L := range /*line rQYpaBAGZ.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line YPdh17m.go:1*/ common.HexToHash("0")].BlC2JXSmAuvQ {
					if sTaCTHWa4L == bBrNWS3ghg.Hash {
						/*line BT5ha1h.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line kmtcyfGamKF.go:1*/ common.HexToHash("0")].BlC2JXSmAuvQ = /*line qg6aTtjMp.go:1*/ utils.CmTWaXnCjAw[common.Hash](mTZDwMx /*line Mj7mkh.go:1*/, cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line AMEo3G7Ga.go:1*/ common.HexToHash("0")].BlC2JXSmAuvQ)
						break
					}
				}
			}
		}

		for _, _4arcQEvNi := range bBrNWS3ghg.RwSet {
			for _, kmDgMw_b := range _4arcQEvNi.WriteSet {
				if fsNx3bxqUIt, i3KmQgjb := /*line FTBN1kCBNat.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[_4arcQEvNi.Address][ /*line G2CJaGWO.go:1*/ common.HexToHash(kmDgMw_b)]; i3KmQgjb {
					fsNx3bxqUIt.UPa0r5o = false
				}
			}
			for _, ceBkC2 := range _4arcQEvNi.ReadSet {
				if fsNx3bxqUIt, i3KmQgjb := /*line uF67MNj9QnKY.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[_4arcQEvNi.Address][ /*line wzmUHDu3.go:1*/ common.HexToHash(ceBkC2)]; i3KmQgjb {
					fsNx3bxqUIt.UPa0r5o = false
				}
			}
		}

		if fsNx3bxqUIt, i3KmQgjb := /*line xFpsSU.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.From][ /*line Lm88eu8j.go:1*/ common.HexToHash("0")]; i3KmQgjb {
			fsNx3bxqUIt.UPa0r5o = false
		}
		if bBrNWS3ghg.TXType == types.TRANSFER {
			if fsNx3bxqUIt, i3KmQgjb := /*line Zp5S0Xych.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[bBrNWS3ghg.To][ /*line QfK9_h8nuKIC.go:1*/ common.HexToHash("0")]; i3KmQgjb {
				fsNx3bxqUIt.UPa0r5o = false
			}
		}
	}

	is2bz5spWcWF := /*line A1TNLecyZH.go:1*/ cu1RSg2YR.pd.GeneratePayload()
	for _, k9KT1tUY := range is2bz5spWcWF {
		kt6kudfae := false
		if k9KT1tUY.TXType == types.TRANSFER {
			if fsNx3bxqUIt, i3KmQgjb := /*line Cf77_Q64g.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[k9KT1tUY.From]; i3KmQgjb {
				if fsNx3bxqUIt[ /*line SHyxTn3xa.go:1*/ utils.SlotToKey(0)].UPa0r5o {
					kt6kudfae = true
				}
			} else if fsNx3bxqUIt, i3KmQgjb := /*line GDWm49.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[k9KT1tUY.To]; i3KmQgjb {
				if fsNx3bxqUIt[ /*line BksOAf7y.go:1*/ utils.SlotToKey(0)].UPa0r5o {
					kt6kudfae = true
				}
			}
		} else if k9KT1tUY.TXType == types.SMARTCONTRACT {
		Rwset:
			for _, ds58OBJNw := range k9KT1tUY.RwSet {
				if fsNx3bxqUIt, i3KmQgjb := /*line oMvynO.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[ds58OBJNw.Address]; i3KmQgjb {
					for _, kmDgMw_b := range ds58OBJNw.WriteSet {
						if kmDgMw_b == func() /*line S_8OFf.go:1*/ string {
							data := [] /*line S_8OFf.go:1*/ byte("{LNYZR/dQr")
							positions := [...]byte{4, 5, 8, 4, 6, 3, 0, 2, 3, 3, 6, 5, 1, 0}
							for i := 0; i < 14; i += 2 {
								localKey := /*line S_8OFf.go:1*/ byte(i) + /*line S_8OFf.go:1*/ byte(positions[i]^positions[i+1]) + 20
								data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
							}
							return /*line S_8OFf.go:1*/ string(data)
						}() {
							if ceGD01WXm4W, i3KmQgjb := /*line XeJXQv.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[k9KT1tUY.From]; i3KmQgjb {
								if ceGD01WXm4W[ /*line Hfp0RegMya.go:1*/ utils.SlotToKey(0)].UPa0r5o {
									kt6kudfae = true
									break Rwset
								}
							}
						} else {
							if z28l_XYKz, i3KmQgjb := fsNx3bxqUIt[ /*line gcY0aCv_.go:1*/ common.HexToHash(kmDgMw_b)]; i3KmQgjb {
								if z28l_XYKz.UPa0r5o {
									kt6kudfae = true
									break Rwset
								}
							}
						}
					}
				}
			}
		}
		if !kt6kudfae {
			yxYJX82Xsq = /*line EJ_Lt14g.go:1*/ append(yxYJX82Xsq, k9KT1tUY)
		} else {

			if _, i3KmQgjb := cu1RSg2YR.pbftMeasure.ConflictTransaction[k9KT1tUY.Hash]; !i3KmQgjb {
				cu1RSg2YR.pbftMeasure.ConflictTransaction[k9KT1tUY.Hash] = 1
			} else {
				cu1RSg2YR.pbftMeasure.ConflictTransaction[k9KT1tUY.Hash]++
			}
			/*line Pel6rIaJ7UX.go:1*/ cu1RSg2YR.pd.AddTxn(k9KT1tUY)
		}
	}

	return yxYJX82Xsq
}

func (cu1RSg2YR *PBFT) GetSCT() map[common.Address]map[common.Hash]*blockchain.SnapshotControlTable {
	return /*line SYQjVE.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()
}

func (cu1RSg2YR *PBFT) CreateSnapshot(AmaNjdAdS []*message.Transaction, a3qJWd types.View, pHERGPksgf30 types.Epoch) ([]*blockchain.LocalSnapshot, []*blockchain.LocalContract) {
	thxhyC_T9ga := /*line JbAQOqT2vEqw.go:1*/ make([]*blockchain.LocalSnapshot, 0)
	ia9eCe8vlOON := /*line IAqX19nx.go:1*/ make([]*blockchain.LocalContract, 0)
	lEJEixo2I := /*line S7arrCZli0.go:1*/ make(map[common.Hash]bool, 0)

	for b8kFRD, uaTLqEF := range /*line aH7gyp0.go:1*/ cu1RSg2YR.bc.GetExpectedSCT() {
		for mTZDwMx, z28l_XYKz := range uaTLqEF {
			if /*line fCDy3eVnu.go:1*/ len(z28l_XYKz.BlC2JXSmAuvQ) > 0 && !z28l_XYKz.UPa0r5o {
				xy5c2O := &blockchain.LocalSnapshot{
					Address: b8kFRD,
					Slot:    mTZDwMx,
				}
				for _, bQZ9a4xL := range z28l_XYKz.BlC2JXSmAuvQ {
					xy5c2O.RTCS = /*line gUxDVuPo.go:1*/ append(xy5c2O.RTCS, bQZ9a4xL)
					lEJEixo2I[bQZ9a4xL] = true
				}
				if /*line cdSNnraR.go:1*/ cu1RSg2YR.bc.GetStateDB().GetCode(b8kFRD) == nil {
					xy5c2O.Value = /*line Zgjo42.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(b8kFRD).String()
					z28l_XYKz.NQB0qlzATMvq = /*line qaMhnnkEw.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(b8kFRD).String()
				} else {
					xy5c2O.Value = /*line mSuKTUAtaDsq.go:1*/ cu1RSg2YR.bc.GetStateDB().GetState(b8kFRD, mTZDwMx).String()
					z28l_XYKz.NQB0qlzATMvq = /*line VTwMSbq6nM.go:1*/ cu1RSg2YR.bc.GetStateDB().GetState(b8kFRD, mTZDwMx).String()
				}
				z28l_XYKz.UPa0r5o = true
				thxhyC_T9ga = /*line P9_LVE.go:1*/ append(thxhyC_T9ga, xy5c2O)
			}
		}
	}

	for _, mLro608nVKLl := range AmaNjdAdS {
		if mLro608nVKLl.TXType == types.SMARTCONTRACT {
			for _, ds58OBJNw := range mLro608nVKLl.RwSet {
				ezBex70Hy := &blockchain.LocalContract{
					Address: ds58OBJNw.Address,
					Code:/*line FyihJ0Hls8S.go:1*/ cu1RSg2YR.bc.GetStateDB().GetCode(ds58OBJNw.Address),
				}
				ia9eCe8vlOON = /*line pzb6Ts3ANia.go:1*/ append(ia9eCe8vlOON, ezBex70Hy)
			}
		}
	}

	ia9eCe8vlOON = /*line d5jOH8Sp.go:1*/ lo.UniqBy[*blockchain.LocalContract, string](ia9eCe8vlOON, func(txGhxt1Xx *blockchain.LocalContract) string {
		return /*line oDOiJsq.go:1*/ txGhxt1Xx.Address.Hex()
	})

	for _, mLro608nVKLl := range AmaNjdAdS {
		pqKOi5u0pS := 0
		if mLro608nVKLl.TXType == types.TRANSFER {
			if /*line kF7dpy.go:1*/ utils.CalculateShardToSend([]common.Address{mLro608nVKLl.To} /*line lo4hXO_2c.go:1*/, config.GetConfig().ShardCount)[0] != /*line CSeAKKddZajV.go:1*/ cu1RSg2YR.Shard() {
				for _, aVU16qtmHkO := range thxhyC_T9ga {
					if aVU16qtmHkO.Address == mLro608nVKLl.From {
						aVU16qtmHkO.RTCS = /*line MqJXlsFkZf.go:1*/ append(aVU16qtmHkO.RTCS, mLro608nVKLl.Hash)
						b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line sDW9KB.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address /*line eI1jOSi4.go:1*/, utils.SlotToKey(0), aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
						/*line GTdeNACvzA6W.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						pqKOi5u0pS++
						break
					}
				}
				if pqKOi5u0pS == 0 {
					xy5c2O := &blockchain.LocalSnapshot{
						Address: mLro608nVKLl.From,
						Slot:/*line Ff2jjTLFFQ.go:1*/ utils.SlotToKey(0),
						Value:/*line qviPlc3q3ajR.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(mLro608nVKLl.From).String(),
					}
					xy5c2O.RTCS = /*line vjo4NWh0dtY9.go:1*/ append(xy5c2O.RTCS, mLro608nVKLl.Hash)
					thxhyC_T9ga = /*line i3v_VdkWa.go:1*/ append(thxhyC_T9ga, xy5c2O)
					b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line fj92Ta.go:1*/ cu1RSg2YR.CreateSCTObj(xy5c2O.Address /*line cnccusQoKyy.go:1*/, utils.SlotToKey(0), xy5c2O.Value, []common.Hash{mLro608nVKLl.Hash})
					/*line F_RpLWMF.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
				}
			} else if /*line DcC0z5rADHUB.go:1*/ utils.CalculateShardToSend([]common.Address{mLro608nVKLl.From} /*line JltZcDzWG.go:1*/, config.GetConfig().ShardCount)[0] != /*line rFmk4Ra6qhX.go:1*/ cu1RSg2YR.Shard() {
				for _, aVU16qtmHkO := range thxhyC_T9ga {
					if aVU16qtmHkO.Address == mLro608nVKLl.To {
						aVU16qtmHkO.RTCS = /*line pqjaawd_dj.go:1*/ append(aVU16qtmHkO.RTCS, mLro608nVKLl.Hash)
						b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line BAvEh2XaKa9.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address /*line RuCeZjOLxh.go:1*/, utils.SlotToKey(0), aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
						/*line ZlKEZ0lO.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						pqKOi5u0pS++
						break
					}
				}
				if pqKOi5u0pS == 0 {
					xy5c2O := &blockchain.LocalSnapshot{
						Address: mLro608nVKLl.To,
						Slot:/*line uMCSLXoAV.go:1*/ utils.SlotToKey(0),
						Value:/*line ISvwOkQQ.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(mLro608nVKLl.To).String(),
					}
					xy5c2O.RTCS = /*line GdavpHT.go:1*/ append(xy5c2O.RTCS, mLro608nVKLl.Hash)
					thxhyC_T9ga = /*line CMuaZM8a.go:1*/ append(thxhyC_T9ga, xy5c2O)
					b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line rhHFTDYUcc.go:1*/ cu1RSg2YR.CreateSCTObj(xy5c2O.Address /*line k6lvxMpBSLu.go:1*/, utils.SlotToKey(0), xy5c2O.Value, []common.Hash{mLro608nVKLl.Hash})
					/*line DaG3ouR9Yw.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
				}
			}
		} else if mLro608nVKLl.TXType == types.SMARTCONTRACT {
			if /*line aK2MPaWONgL.go:1*/ utils.CalculateShardToSend([]common.Address{mLro608nVKLl.From} /*line jH2q19_kRpMZ.go:1*/, config.GetConfig().ShardCount)[0] == /*line L7hpAH.go:1*/ cu1RSg2YR.Shard() {
				for _, aVU16qtmHkO := range thxhyC_T9ga {
					if aVU16qtmHkO.Address == mLro608nVKLl.From {
						aVU16qtmHkO.RTCS = /*line kMCWppP.go:1*/ append(aVU16qtmHkO.RTCS, mLro608nVKLl.Hash)
						b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line U9A6Kf.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address /*line jgo3oQY.go:1*/, utils.SlotToKey(0), aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
						/*line bFWAfE.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						pqKOi5u0pS++
						break
					}
				}
				if pqKOi5u0pS == 0 {
					xy5c2O := &blockchain.LocalSnapshot{
						Address: mLro608nVKLl.From,
						Slot:/*line G2VciF.go:1*/ utils.SlotToKey(0),
						Value:/*line Jv84F_Bd1.go:1*/ cu1RSg2YR.bc.GetStateDB().GetBalance(mLro608nVKLl.From).String(),
					}
					xy5c2O.RTCS = /*line SZkn3Y2T3.go:1*/ append(xy5c2O.RTCS, mLro608nVKLl.Hash)
					thxhyC_T9ga = /*line J6bafoKgq.go:1*/ append(thxhyC_T9ga, xy5c2O)
					b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line rUUhksda4s.go:1*/ cu1RSg2YR.CreateSCTObj(xy5c2O.Address /*line DEKyOc.go:1*/, utils.SlotToKey(0), xy5c2O.Value, []common.Hash{mLro608nVKLl.Hash})
					/*line bBmltCqq.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
				}
			}
			for _, ds58OBJNw := range mLro608nVKLl.RwSet {
				if /*line V2RaD6hCI.go:1*/ utils.CalculateShardToSend([]common.Address{ds58OBJNw.Address} /*line JSE54z_u.go:1*/, config.GetConfig().ShardCount)[0] == /*line VP9R75mN.go:1*/ cu1RSg2YR.Shard() {
					for _, ceBkC2 := range ds58OBJNw.ReadSet {
						pqKOi5u0pS = 0
						for _, aVU16qtmHkO := range thxhyC_T9ga {
							if aVU16qtmHkO.Address == ds58OBJNw.Address && aVU16qtmHkO.Slot == /*line mtjGU7.go:1*/ common.HexToHash(ceBkC2) {
								for _, OTC_a4FQrin1 := range aVU16qtmHkO.RTCS {
									if OTC_a4FQrin1 == mLro608nVKLl.Hash {
										pqKOi5u0pS++
										break
									}
								}
								if pqKOi5u0pS == 0 {
									aVU16qtmHkO.RTCS = /*line BQLpeVnR.go:1*/ append(aVU16qtmHkO.RTCS, mLro608nVKLl.Hash)
									b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line magB0c8.go:1*/ cu1RSg2YR.CreateSCTObj(aVU16qtmHkO.Address, aVU16qtmHkO.Slot, aVU16qtmHkO.Value, []common.Hash{mLro608nVKLl.Hash})
									/*line XJpEYjdofgvC.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
									pqKOi5u0pS++
								}
								break
							}
						}
						if pqKOi5u0pS == 0 {
							xy5c2O := &blockchain.LocalSnapshot{
								Address: ds58OBJNw.Address,
								Slot:/*line CktKUaGO.go:1*/ common.HexToHash(ceBkC2),
								Value:/*line xDbcMAsi.go:1*/ cu1RSg2YR.bc.GetStateDB().GetState(ds58OBJNw.Address /*line jhF7KH0aO.go:1*/, common.HexToHash(ceBkC2)).String(),
							}
							xy5c2O.RTCS = /*line cXItRECP.go:1*/ append(xy5c2O.RTCS, mLro608nVKLl.Hash)
							thxhyC_T9ga = /*line Nb8tv2Jx.go:1*/ append(thxhyC_T9ga, xy5c2O)
							b8kFRD, mfg9OCAuCN6, uaTLqEF := /*line qbDAFowg05tn.go:1*/ cu1RSg2YR.CreateSCTObj(xy5c2O.Address, xy5c2O.Slot, xy5c2O.Value, []common.Hash{mLro608nVKLl.Hash})
							/*line RbT8BZZk0Us.go:1*/ cu1RSg2YR.AddAndUpdateSCT(b8kFRD, mfg9OCAuCN6, uaTLqEF)
						}
					}
				}
			}
		}
	}

	return thxhyC_T9ga, ia9eCe8vlOON
}

func (cu1RSg2YR *PBFT) CreateSCTObj(b8kFRD common.Address, iRXysfs98Do common.Hash, pmRleS7K string, t2OXXM3pqr []common.Hash) (common.Address, common.Hash, *blockchain.SnapshotControlTable) {
	return b8kFRD, iRXysfs98Do, &blockchain.SnapshotControlTable{
		NQB0qlzATMvq: pmRleS7K,
		BlC2JXSmAuvQ: t2OXXM3pqr,
		UPa0r5o:      true,
	}
}

func (cu1RSg2YR *PBFT) AddAndUpdateSCT(rmKrOzZtQLVb common.Address, mfg9OCAuCN6 common.Hash, aVU16qtmHkO *blockchain.SnapshotControlTable) {
	if uaTLqEF, i3KmQgjb := /*line yOuPSvLGrMp.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[rmKrOzZtQLVb][mfg9OCAuCN6]; i3KmQgjb {
		if /*line uRrLAHi.go:1*/ len(aVU16qtmHkO.BlC2JXSmAuvQ) > 0 {
			uaTLqEF.BlC2JXSmAuvQ = /*line EdP6sNB5B2.go:1*/ append(uaTLqEF.BlC2JXSmAuvQ, aVU16qtmHkO.BlC2JXSmAuvQ...)
			uaTLqEF.UPa0r5o = true
		} else if /*line dayjE3r5G.go:1*/ len(aVU16qtmHkO.BlC2JXSmAuvQ) == 0 && /*line gn_kDIwaq.go:1*/ len(uaTLqEF.BlC2JXSmAuvQ) == 0 {
			uaTLqEF.UPa0r5o = false
		}
	} else {
		if /*line ma_avuP4Au.go:1*/ len(aVU16qtmHkO.BlC2JXSmAuvQ) > 0 {
			aVU16qtmHkO.UPa0r5o = true
		}

		if _, kbgkwuEjk := /*line lCFjXFV.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[rmKrOzZtQLVb]; !kbgkwuEjk {
			/*line wh8w7BX4.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[rmKrOzZtQLVb] = /*line FlYCZfpn.go:1*/ make(map[common.Hash]*blockchain.SnapshotControlTable)
		}
		/*line WP630Vta.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[rmKrOzZtQLVb][mfg9OCAuCN6] = /*line FI996v.go:1*/ new(blockchain.SnapshotControlTable)
		/*line x8bKXLpg2Q.go:1*/ cu1RSg2YR.bc.GetExpectedSCT()[rmKrOzZtQLVb][mfg9OCAuCN6] = aVU16qtmHkO
	}
}

func (cu1RSg2YR *PBFT) GetHighQC() *quorum.HPOWa1PT0xzq {
	/*line LXMFJfoUgOO.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line o7GzpfamsPAP.go:1*/ cu1RSg2YR.mu.Unlock()
	return cu1RSg2YR.highQC
}

func (cu1RSg2YR *PBFT) GetHighCQC() *quorum.HPOWa1PT0xzq {
	/*line oNTVnEf6uv4.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line JdXkvKXZPfpP.go:1*/ cu1RSg2YR.mu.Unlock()
	return cu1RSg2YR.highCQC
}

func (cu1RSg2YR *PBFT) GetChainStatus() string {
	vO8QgiLWR_a := /*line O647Mu9k.go:1*/ cu1RSg2YR.bc.GetChainGrowth()
	luidf2 := /*line fnoq7dn50T2L.go:1*/ cu1RSg2YR.bc.GetBlockIntervals()
	return /*line H9xp2rAv.go:1*/ fmt.Sprintf(func() /*line aAaKNn.go:1*/ string {
		fullData := [] /*line aAaKNn.go:1*/ byte("\xeej\x99\xe4\xba=\x82\xc4B\xf3\xbb\xb5۴\x94\xe1\xdd\xca\xcd.=\x1d\x9a\xb6\xa8\xa0-)\x03\x13\x89\xc1\xb2\uf0a3 \x88\xff\xa1\x0ff\x11R\\\x9c\xa1\xea9\x02\x1d\xdb`\x12\xac\xb6\x15dr3:\xd5\uf523D \xec\x84a`\x84V\xafBFqCes\x8f\x19+]>\b\xfdI%\xdb+\xa2\xe24^˹\x8c\xe8\xc0\xe7\x04A^\xd6>\xaa\xfc\x10O\x8a\xd7j(\xc2-\xdf6h\xfc#\xef\vtn9\xfe\xab\xcb\x11\x16v\x14\xf0\xaf\xfb\xc5\x16T\xbe\xd2\x18JhD\xc1\xa9\xd3:\xee\xae\xd1g\xc4z\x8dc\xa5tX&\x1e")
		idxKey := [] /*line aAaKNn.go:1*/ byte("(c\xd0w\t*]\xdcN\x9f\xf2v.\x1e\x01S")
		data := /*line aAaKNn.go:1*/ make([]byte, 0, 82)
		data = /*line aAaKNn.go:1*/ append(data, fullData[184^ /*line aAaKNn.go:1*/ int(idxKey[2])]^fullData[75^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[218^ /*line aAaKNn.go:1*/ int(idxKey[2])]+fullData[209^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[158^ /*line aAaKNn.go:1*/ int(idxKey[7])]-fullData[182^ /*line aAaKNn.go:1*/ int(idxKey[7])], fullData[71^ /*line aAaKNn.go:1*/ int(idxKey[5])]^fullData[31^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[90^ /*line aAaKNn.go:1*/ int(idxKey[11])]+fullData[239^ /*line aAaKNn.go:1*/ int(idxKey[11])], fullData[37^ /*line aAaKNn.go:1*/ int(idxKey[5])]^fullData[33^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[176^ /*line aAaKNn.go:1*/ int(idxKey[9])]-fullData[189^ /*line aAaKNn.go:1*/ int(idxKey[9])], fullData[34^ /*line aAaKNn.go:1*/ int(idxKey[12])]-fullData[173^ /*line aAaKNn.go:1*/ int(idxKey[12])], fullData[0^ /*line aAaKNn.go:1*/ int(idxKey[4])]+fullData[122^ /*line aAaKNn.go:1*/ int(idxKey[4])], fullData[66^ /*line aAaKNn.go:1*/ int(idxKey[6])]-fullData[3^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[211^ /*line aAaKNn.go:1*/ int(idxKey[9])]+fullData[250^ /*line aAaKNn.go:1*/ int(idxKey[9])], fullData[66^ /*line aAaKNn.go:1*/ int(idxKey[2])]^fullData[227^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[46^ /*line aAaKNn.go:1*/ int(idxKey[0])]+fullData[173^ /*line aAaKNn.go:1*/ int(idxKey[0])], fullData[11^ /*line aAaKNn.go:1*/ int(idxKey[5])]-fullData[68^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[174^ /*line aAaKNn.go:1*/ int(idxKey[12])]+fullData[13^ /*line aAaKNn.go:1*/ int(idxKey[12])], fullData[113^ /*line aAaKNn.go:1*/ int(idxKey[5])]-fullData[57^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[215^ /*line aAaKNn.go:1*/ int(idxKey[2])]^fullData[211^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[28^ /*line aAaKNn.go:1*/ int(idxKey[15])]+fullData[79^ /*line aAaKNn.go:1*/ int(idxKey[15])], fullData[37^ /*line aAaKNn.go:1*/ int(idxKey[3])]+fullData[30^ /*line aAaKNn.go:1*/ int(idxKey[3])], fullData[0^ /*line aAaKNn.go:1*/ int(idxKey[13])]^fullData[93^ /*line aAaKNn.go:1*/ int(idxKey[13])], fullData[47^ /*line aAaKNn.go:1*/ int(idxKey[5])]^fullData[164^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[120^ /*line aAaKNn.go:1*/ int(idxKey[10])]-fullData[175^ /*line aAaKNn.go:1*/ int(idxKey[10])], fullData[77^ /*line aAaKNn.go:1*/ int(idxKey[6])]^fullData[80^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[136^ /*line aAaKNn.go:1*/ int(idxKey[13])]+fullData[150^ /*line aAaKNn.go:1*/ int(idxKey[13])], fullData[78^ /*line aAaKNn.go:1*/ int(idxKey[2])]-fullData[68^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[33^ /*line aAaKNn.go:1*/ int(idxKey[15])]^fullData[15^ /*line aAaKNn.go:1*/ int(idxKey[15])], fullData[145^ /*line aAaKNn.go:1*/ int(idxKey[4])]^fullData[1^ /*line aAaKNn.go:1*/ int(idxKey[4])], fullData[100^ /*line aAaKNn.go:1*/ int(idxKey[5])]+fullData[171^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[89^ /*line aAaKNn.go:1*/ int(idxKey[0])]-fullData[67^ /*line aAaKNn.go:1*/ int(idxKey[0])], fullData[41^ /*line aAaKNn.go:1*/ int(idxKey[15])]+fullData[107^ /*line aAaKNn.go:1*/ int(idxKey[15])], fullData[109^ /*line aAaKNn.go:1*/ int(idxKey[0])]^fullData[25^ /*line aAaKNn.go:1*/ int(idxKey[0])], fullData[64^ /*line aAaKNn.go:1*/ int(idxKey[3])]+fullData[87^ /*line aAaKNn.go:1*/ int(idxKey[3])], fullData[107^ /*line aAaKNn.go:1*/ int(idxKey[13])]-fullData[35^ /*line aAaKNn.go:1*/ int(idxKey[13])], fullData[189^ /*line aAaKNn.go:1*/ int(idxKey[5])]-fullData[92^ /*line aAaKNn.go:1*/ int(idxKey[5])], fullData[239^ /*line aAaKNn.go:1*/ int(idxKey[9])]-fullData[232^ /*line aAaKNn.go:1*/ int(idxKey[9])], fullData[123^ /*line aAaKNn.go:1*/ int(idxKey[6])]^fullData[41^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[201^ /*line aAaKNn.go:1*/ int(idxKey[2])]-fullData[224^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[131^ /*line aAaKNn.go:1*/ int(idxKey[14])]^fullData[56^ /*line aAaKNn.go:1*/ int(idxKey[14])], fullData[63^ /*line aAaKNn.go:1*/ int(idxKey[0])]+fullData[72^ /*line aAaKNn.go:1*/ int(idxKey[0])], fullData[54^ /*line aAaKNn.go:1*/ int(idxKey[4])]-fullData[28^ /*line aAaKNn.go:1*/ int(idxKey[4])], fullData[138^ /*line aAaKNn.go:1*/ int(idxKey[10])]-fullData[187^ /*line aAaKNn.go:1*/ int(idxKey[10])], fullData[148^ /*line aAaKNn.go:1*/ int(idxKey[7])]^fullData[136^ /*line aAaKNn.go:1*/ int(idxKey[7])], fullData[55^ /*line aAaKNn.go:1*/ int(idxKey[13])]^fullData[85^ /*line aAaKNn.go:1*/ int(idxKey[13])], fullData[78^ /*line aAaKNn.go:1*/ int(idxKey[0])]^fullData[19^ /*line aAaKNn.go:1*/ int(idxKey[0])], fullData[101^ /*line aAaKNn.go:1*/ int(idxKey[3])]+fullData[121^ /*line aAaKNn.go:1*/ int(idxKey[3])], fullData[193^ /*line aAaKNn.go:1*/ int(idxKey[6])]+fullData[119^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[54^ /*line aAaKNn.go:1*/ int(idxKey[1])]+fullData[48^ /*line aAaKNn.go:1*/ int(idxKey[1])], fullData[143^ /*line aAaKNn.go:1*/ int(idxKey[2])]-fullData[175^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[44^ /*line aAaKNn.go:1*/ int(idxKey[11])]^fullData[60^ /*line aAaKNn.go:1*/ int(idxKey[11])], fullData[206^ /*line aAaKNn.go:1*/ int(idxKey[6])]-fullData[27^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[177^ /*line aAaKNn.go:1*/ int(idxKey[12])]-fullData[143^ /*line aAaKNn.go:1*/ int(idxKey[12])], fullData[198^ /*line aAaKNn.go:1*/ int(idxKey[2])]-fullData[74^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[45^ /*line aAaKNn.go:1*/ int(idxKey[15])]^fullData[10^ /*line aAaKNn.go:1*/ int(idxKey[15])], fullData[216^ /*line aAaKNn.go:1*/ int(idxKey[7])]-fullData[76^ /*line aAaKNn.go:1*/ int(idxKey[7])], fullData[64^ /*line aAaKNn.go:1*/ int(idxKey[14])]^fullData[142^ /*line aAaKNn.go:1*/ int(idxKey[14])], fullData[135^ /*line aAaKNn.go:1*/ int(idxKey[2])]-fullData[203^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[87^ /*line aAaKNn.go:1*/ int(idxKey[12])]+fullData[20^ /*line aAaKNn.go:1*/ int(idxKey[12])], fullData[125^ /*line aAaKNn.go:1*/ int(idxKey[14])]^fullData[140^ /*line aAaKNn.go:1*/ int(idxKey[14])], fullData[51^ /*line aAaKNn.go:1*/ int(idxKey[12])]+fullData[5^ /*line aAaKNn.go:1*/ int(idxKey[12])], fullData[83^ /*line aAaKNn.go:1*/ int(idxKey[11])]^fullData[110^ /*line aAaKNn.go:1*/ int(idxKey[11])], fullData[73^ /*line aAaKNn.go:1*/ int(idxKey[6])]+fullData[5^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[0^ /*line aAaKNn.go:1*/ int(idxKey[1])]+fullData[85^ /*line aAaKNn.go:1*/ int(idxKey[1])], fullData[91^ /*line aAaKNn.go:1*/ int(idxKey[7])]+fullData[167^ /*line aAaKNn.go:1*/ int(idxKey[7])], fullData[30^ /*line aAaKNn.go:1*/ int(idxKey[15])]+fullData[119^ /*line aAaKNn.go:1*/ int(idxKey[15])], fullData[122^ /*line aAaKNn.go:1*/ int(idxKey[13])]+fullData[89^ /*line aAaKNn.go:1*/ int(idxKey[13])], fullData[2^ /*line aAaKNn.go:1*/ int(idxKey[15])]^fullData[46^ /*line aAaKNn.go:1*/ int(idxKey[15])], fullData[110^ /*line aAaKNn.go:1*/ int(idxKey[14])]^fullData[138^ /*line aAaKNn.go:1*/ int(idxKey[14])], fullData[183^ /*line aAaKNn.go:1*/ int(idxKey[2])]+fullData[188^ /*line aAaKNn.go:1*/ int(idxKey[2])], fullData[96^ /*line aAaKNn.go:1*/ int(idxKey[14])]+fullData[99^ /*line aAaKNn.go:1*/ int(idxKey[14])], fullData[217^ /*line aAaKNn.go:1*/ int(idxKey[6])]-fullData[219^ /*line aAaKNn.go:1*/ int(idxKey[6])], fullData[219^ /*line aAaKNn.go:1*/ int(idxKey[9])]+fullData[10^ /*line aAaKNn.go:1*/ int(idxKey[9])], fullData[178^ /*line aAaKNn.go:1*/ int(idxKey[9])]-fullData[63^ /*line aAaKNn.go:1*/ int(idxKey[9])], fullData[89^ /*line aAaKNn.go:1*/ int(idxKey[4])]^fullData[9^ /*line aAaKNn.go:1*/ int(idxKey[4])], fullData[84^ /*line aAaKNn.go:1*/ int(idxKey[8])]-fullData[223^ /*line aAaKNn.go:1*/ int(idxKey[8])], fullData[183^ /*line aAaKNn.go:1*/ int(idxKey[9])]-fullData[161^ /*line aAaKNn.go:1*/ int(idxKey[9])], fullData[65^ /*line aAaKNn.go:1*/ int(idxKey[14])]-fullData[61^ /*line aAaKNn.go:1*/ int(idxKey[14])], fullData[81^ /*line aAaKNn.go:1*/ int(idxKey[11])]^fullData[250^ /*line aAaKNn.go:1*/ int(idxKey[11])], fullData[76^ /*line aAaKNn.go:1*/ int(idxKey[8])]+fullData[96^ /*line aAaKNn.go:1*/ int(idxKey[8])], fullData[28^ /*line aAaKNn.go:1*/ int(idxKey[12])]-fullData[120^ /*line aAaKNn.go:1*/ int(idxKey[12])], fullData[114^ /*line aAaKNn.go:1*/ int(idxKey[1])]-fullData[254^ /*line aAaKNn.go:1*/ int(idxKey[1])], fullData[128^ /*line aAaKNn.go:1*/ int(idxKey[4])]+fullData[61^ /*line aAaKNn.go:1*/ int(idxKey[4])])
		return /*line aAaKNn.go:1*/ string(data)
	}(), /*line aO26YF.go:1*/ cu1RSg2YR.ID() /*line L6vEjaGbJF.go:1*/, cu1RSg2YR.pm.GetCurView(), vO8QgiLWR_a, luidf2)
}

func (cu1RSg2YR *PBFT) GetHighBlockHeight() types.BlockHeight {
	/*line F4z74VtEw8cp.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line rB0IIvJ8a7.go:1*/ cu1RSg2YR.mu.Unlock()
	return cu1RSg2YR.highQC.BlockHeight
}

func (cu1RSg2YR *PBFT) GetLastBlockHeight() types.BlockHeight {
	/*line uSJ0db.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line pxdbTEd_z.go:1*/ cu1RSg2YR.mu.Unlock()
	return cu1RSg2YR.highCQC.BlockHeight
}

func (cu1RSg2YR *PBFT) GetLastVotedBlockHeight() types.BlockHeight {
	/*line hxQ9GxMs.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line yLhWYp.go:1*/ cu1RSg2YR.mu.Unlock()
	return cu1RSg2YR.lastVotedBlockHeight
}

func (cu1RSg2YR *PBFT) aDTslULytPz_(cWdSpmiBLOz *quorum.HPOWa1PT0xzq) {
	/*line GcyE13H0.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line zJjrBBjWI9Z4.go:1*/ cu1RSg2YR.mu.Unlock()
	if cWdSpmiBLOz.BlockHeight > cu1RSg2YR.highQC.BlockHeight {
		cu1RSg2YR.highQC = cWdSpmiBLOz
	}
}

func (cu1RSg2YR *PBFT) z6VaPTrjpK(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) {
	/*line e4agMG.go:1*/ cu1RSg2YR.mu.Lock()
	defer /*line j0PhrZaBeqN.go:1*/ cu1RSg2YR.mu.Unlock()
	if dWmn_Gtk1d.BlockHeight > cu1RSg2YR.highCQC.BlockHeight {
		cu1RSg2YR.highCQC = dWmn_Gtk1d
		if /*line tAWcXFt.go:1*/ cu1RSg2YR.IsLeader( /*line FPtAhvmJ27.go:1*/ cu1RSg2YR.ID() /*line WXVNQL_FQx3S.go:1*/, cu1RSg2YR.pm.GetCurView() /*line cs5P4CYakKl.go:1*/, cu1RSg2YR.pm.GetCurEpoch()) {
			cu1RSg2YR.updatedQC <- dWmn_Gtk1d
		}
	}
}

func (cu1RSg2YR *PBFT) xAnHtUT9ATI(bQ_GL8MJ_E types.BlockHeight) error {
	if bQ_GL8MJ_E < cu1RSg2YR.lastVotedBlockHeight {
		return /*line yh5GCITa0.go:1*/ errors.New(func() /*line lZz6mrr.go:1*/ string {
			key := [] /*line lZz6mrr.go:1*/ byte("S]\xc3\xff\xb0\x83m\xa7\xee\xfd{\x87\x9cM>\xfd:\xaeG\xda`\xa6\a\x84\xaaX\xdaa\x98b\x15-g\xb9\xed\xe6\x85\xc7[\xb6\xb1s\x88\x8b\x9d\xcb\xd9\x12\x82_1N\x94\x93\xd3\x19z\xc92")
			data := [] /*line lZz6mrr.go:1*/ byte("Ǿ5f\x15\xf7\x8d\tZl\xde\xf2䲧d\xa2\"gC\xd3\xc6s\xf3!\xbdL\x81\f\xcav\x9b\x87-UK\xa53\xbc)%\x93\xfe\xfa\x110=2\xe4ˠ\xb1\xff\xdb8\x82\xe11\xa6")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line lZz6mrr.go:1*/ string(data)
		}())
	}
	cu1RSg2YR.lastVotedBlockHeight = bQ_GL8MJ_E
	return nil
}

func (cu1RSg2YR *PBFT) wMIZeBk(cWdSpmiBLOz *quorum.HPOWa1PT0xzq) error {
	if cWdSpmiBLOz.BlockHeight <= 1 {
		return nil
	}
	_, k1qj7vimKIy7 := /*line WHf7SkSoo.go:1*/ cu1RSg2YR.bc.GetBlockByID(cWdSpmiBLOz.ZWsU_2)
	if k1qj7vimKIy7 != nil {
		return /*line mVO0meMI.go:1*/ fmt.Errorf(func() /*line SD28Nb.go:1*/ string {
			fullData := [] /*line SD28Nb.go:1*/ byte("\xd9\xcaCk\xf2<G\xc3\f!N\x1a$\xd0\xd07;CBC\xd3\x17\x89\"\xfbq\x9d\xb1Af.\xe7\xfe\x16\xa6\xa2\x92FA\xb5\xa8~\x97\x03\xa5\x06\xb3N\x81k\xd70\x03n\x88#\xccT\xf3\x1a\xaf\x0e\x98\x922\x8e\xb9\x8a \xf9\xddr!\xe8\xcfۊx")
			idxKey := [] /*line SD28Nb.go:1*/ byte("`\x1d\xbf")
			data := /*line SD28Nb.go:1*/ make([]byte, 0, 40)
			data = /*line SD28Nb.go:1*/ append(data, fullData[29^ /*line SD28Nb.go:1*/ int(idxKey[1])]+fullData[94^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[119^ /*line SD28Nb.go:1*/ int(idxKey[0])]^fullData[113^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[134^ /*line SD28Nb.go:1*/ int(idxKey[2])]+fullData[180^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[4^ /*line SD28Nb.go:1*/ int(idxKey[1])]-fullData[41^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[137^ /*line SD28Nb.go:1*/ int(idxKey[2])]^fullData[160^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[183^ /*line SD28Nb.go:1*/ int(idxKey[2])]^fullData[242^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[25^ /*line SD28Nb.go:1*/ int(idxKey[1])]+fullData[3^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[177^ /*line SD28Nb.go:1*/ int(idxKey[2])]^fullData[147^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[27^ /*line SD28Nb.go:1*/ int(idxKey[1])]-fullData[47^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[85^ /*line SD28Nb.go:1*/ int(idxKey[1])]+fullData[31^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[11^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[84^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[33^ /*line SD28Nb.go:1*/ int(idxKey[1])]-fullData[13^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[133^ /*line SD28Nb.go:1*/ int(idxKey[2])]-fullData[254^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[50^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[40^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[32^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[52^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[149^ /*line SD28Nb.go:1*/ int(idxKey[2])]+fullData[244^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[163^ /*line SD28Nb.go:1*/ int(idxKey[2])]^fullData[179^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[148^ /*line SD28Nb.go:1*/ int(idxKey[2])]-fullData[165^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[106^ /*line SD28Nb.go:1*/ int(idxKey[0])]+fullData[117^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[70^ /*line SD28Nb.go:1*/ int(idxKey[0])]-fullData[42^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[62^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[16^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[151^ /*line SD28Nb.go:1*/ int(idxKey[2])]-fullData[172^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[91^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[95^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[99^ /*line SD28Nb.go:1*/ int(idxKey[0])]+fullData[71^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[140^ /*line SD28Nb.go:1*/ int(idxKey[2])]+fullData[255^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[157^ /*line SD28Nb.go:1*/ int(idxKey[2])]^fullData[190^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[36^ /*line SD28Nb.go:1*/ int(idxKey[0])]-fullData[123^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[5^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[35^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[45^ /*line SD28Nb.go:1*/ int(idxKey[1])]-fullData[60^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[145^ /*line SD28Nb.go:1*/ int(idxKey[2])]-fullData[142^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[116^ /*line SD28Nb.go:1*/ int(idxKey[0])]+fullData[68^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[243^ /*line SD28Nb.go:1*/ int(idxKey[2])]-fullData[182^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[250^ /*line SD28Nb.go:1*/ int(idxKey[2])]-fullData[128^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[38^ /*line SD28Nb.go:1*/ int(idxKey[1])]^fullData[90^ /*line SD28Nb.go:1*/ int(idxKey[1])], fullData[111^ /*line SD28Nb.go:1*/ int(idxKey[0])]-fullData[103^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[186^ /*line SD28Nb.go:1*/ int(idxKey[2])]+fullData[159^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[125^ /*line SD28Nb.go:1*/ int(idxKey[0])]^fullData[69^ /*line SD28Nb.go:1*/ int(idxKey[0])], fullData[146^ /*line SD28Nb.go:1*/ int(idxKey[2])]^fullData[136^ /*line SD28Nb.go:1*/ int(idxKey[2])], fullData[114^ /*line SD28Nb.go:1*/ int(idxKey[0])]-fullData[88^ /*line SD28Nb.go:1*/ int(idxKey[0])])
			return /*line SD28Nb.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	v61eq4CI, k1qj7vimKIy7 := /*line qA6hgAs5.go:1*/ cu1RSg2YR.bc.GetParentWorkerBlock(cWdSpmiBLOz.ZWsU_2)
	l_CyibPF := v61eq4CI.(*blockchain.WorkerBlock)
	if k1qj7vimKIy7 != nil {
		return /*line HFQDndtm.go:1*/ fmt.Errorf(func() /*line _eB1b8.go:1*/ string {
			key := [] /*line _eB1b8.go:1*/ byte("U81\xa0\xff̌ՙ\a!\xbcb\r\xc2~\xdd\x1a(\xf9; \x99!\xe9\xddd\xd1\x18\x93(^\x16*\xd2\xe1\xe2\x03\xe5")
			data := [] /*line _eB1b8.go:1*/ byte("6Y_ΐ\xb8\xac\xa0\xe9c@\xc8\a-\xb2\f\xb8|M\x8bIE\xfd\x01\x8b\xb1\v\xb2s\xdbM7qB\xa6\xdb\xc2&\x93")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line _eB1b8.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	if l_CyibPF.BlockHeader.BlockHeight > cu1RSg2YR.preferredBlockHeight {
		cu1RSg2YR.preferredBlockHeight = l_CyibPF.BlockHeader.BlockHeight
	}

	return nil
}

func (cu1RSg2YR *PBFT) ProcessReportByzantine(mzKKWqo *byzantine.ReportByzantine) {
}
func (cu1RSg2YR *PBFT) ProcessReplaceByzantine(g3i7JWF *byzantine.ReplaceByzantine) {
}

var _ = errors.As
var _ = fmt.Append
var _ = strconv.AppendBool
var _ sync.Cond

const _ = time.ANSIC

var _ blockchain.Accept
var _ = byzantine.JEsMGO
var _ config.Bconfig
var _ crypto.Pp__49cd
var _ = election.NewStatic
var _ = evm.TR_8hD4NdZ
var _ state.Code
var _ = runtime.FpzWvzX3
var _ = log.CDebpj
var _ = measure.CalculateAverageTimeDifference
var _ mempool.DUUdwSwZTCm
var _ message.ClientStart
var _ node.BlockBuilder
var _ pacemaker.UDpSaa3
var _ quorum.Commit

const _ = types.ABORT

var _ = utils.GffGNE
var _ = common.AbsolutePath
var _ ethdb.AncientReader
var _ = uint256.ErrBadBufferLength
var _ = lo.AllCharset
