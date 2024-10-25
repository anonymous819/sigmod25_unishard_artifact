//line :1
package pbft

import (
	"errors"
	"fmt"
	"sync"
	"time"

	blockchain "unishard/blockchain"
	byzantine "unishard/byzantine"
	config "unishard/config"
	crypto "unishard/crypto"
	election "unishard/election"
	log "unishard/log"
	node "unishard/node"
	pacemaker "unishard/pacemaker"
	quorum "unishard/quorum"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
)

type L4qxkgqr2rNN struct {
	node.Node
	*election.Static
	mGpWFGC2    *pacemaker.UDpSaa3
	_BUwIvrsOe  types.BlockHeight
	sN6W66jL    types.BlockHeight
	bQsSnXJ     *quorum.HPOWa1PT0xzq
	cj2GSMP     *quorum.HPOWa1PT0xzq
	iSgBUFf9    *blockchain.Y3t7C8s0m
	ffpSbXbMEX0 *quorum.FZT207R[quorum.Vote]
	aK7JnwJ     *quorum.FZT207R[quorum.Commit]
	xwa44pOLa1  chan *blockchain.CoordinationBlock
	dZp5Ovm     chan *blockchain.CoordinationBlock
	jtokCZw_Dtx chan *blockchain.CoordinationBlock
	vHhOhURml   map[common.Hash]*quorum.HPOWa1PT0xzq
	d6XkweCEz   map[common.Hash]*quorum.HPOWa1PT0xzq
	acQiWvFYDH  chan *quorum.HPOWa1PT0xzq
	ta2Y2hhO    *quorum.HPOWa1PT0xzq
	wdUdVrSEU   map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.CoordinationBlock
	sa21VPHg    map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.Qi_7sYpWjR8
	cA3KSEC5m   map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.CoordinationBlock
	tRK_bDOT    sync.Mutex

	nbOZ17j map[types.NodeID]*pacemaker.H8NP1AOKTF
}

func WtH1Mn_f(
	sh_ipJP6 node.Node, yapKe2jjkRuB *pacemaker.UDpSaa3,
	keasWXics5qx *election.Static, l2pHR1ZhFG chan *blockchain.CoordinationBlock,
	gDcB23PkLL chan *blockchain.CoordinationBlock,
	eb1fVADHfFP chan *blockchain.CoordinationBlock) *L4qxkgqr2rNN {
	cu1RSg2YR := /*line VaCz3VSV4.go:1*/ new(L4qxkgqr2rNN)
	cu1RSg2YR.Node = sh_ipJP6
	cu1RSg2YR.Static = keasWXics5qx
	cu1RSg2YR.mGpWFGC2 = yapKe2jjkRuB
	cu1RSg2YR.iSgBUFf9 = /*line GNkFLU21El.go:1*/ blockchain.GEzkgtuHo0( /*line rE7bBPheo61Y.go:1*/ cu1RSg2YR.Shard() /*line Bn9N8FuJz8c.go:1*/, config.GetConfig().DefaultBalance)
	cu1RSg2YR.ffpSbXbMEX0 = /*line h1gfp3w.go:1*/ quorum.Dq3TFZ[quorum.Vote]( /*line z0LOhQQC.go:1*/ config.GetConfig().CommitteeNumber)
	cu1RSg2YR.aK7JnwJ = /*line B74lHPdo1Fr.go:1*/ quorum.Dq3TFZ[quorum.Commit]( /*line D4Cmi2ei0f5G.go:1*/ config.GetConfig().CommitteeNumber)
	cu1RSg2YR.wdUdVrSEU = /*line r_ENCoinOI.go:1*/ make(map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.CoordinationBlock)
	cu1RSg2YR.sa21VPHg = /*line WAtnaX2Y.go:1*/ make(map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.Qi_7sYpWjR8)
	cu1RSg2YR.cA3KSEC5m = /*line vR6h07P5RlOY.go:1*/ make(map[types.Epoch]map[types.View]map[types.BlockHeight]*blockchain.CoordinationBlock)
	cu1RSg2YR.vHhOhURml = /*line P5DihU.go:1*/ make(map[common.Hash]*quorum.HPOWa1PT0xzq)
	cu1RSg2YR.bQsSnXJ = &quorum.HPOWa1PT0xzq{LwVL87: 0}
	cu1RSg2YR.d6XkweCEz = /*line ECUpaO.go:1*/ make(map[common.Hash]*quorum.HPOWa1PT0xzq)
	cu1RSg2YR.ta2Y2hhO = &quorum.HPOWa1PT0xzq{LwVL87: 0}
	cu1RSg2YR.acQiWvFYDH = /*line hDvMAFU8n.go:1*/ make(chan *quorum.HPOWa1PT0xzq, 1)
	cu1RSg2YR.cj2GSMP = &quorum.HPOWa1PT0xzq{LwVL87: 0}
	cu1RSg2YR.xwa44pOLa1 = l2pHR1ZhFG
	cu1RSg2YR.dZp5Ovm = gDcB23PkLL
	cu1RSg2YR.jtokCZw_Dtx = eb1fVADHfFP

	cu1RSg2YR.nbOZ17j = /*line olNkex.go:1*/ make(map[types.NodeID]*pacemaker.H8NP1AOKTF)
	return cu1RSg2YR
}

func (cu1RSg2YR *L4qxkgqr2rNN) GetBlockChain() *blockchain.Y3t7C8s0m {
	return cu1RSg2YR.iSgBUFf9
}

func (cu1RSg2YR *L4qxkgqr2rNN) CreateCoordinationBlock(ndnWqNS *blockchain.CoordinationBlockWithoutHeader) *blockchain.CoordinationBlock {

	var cWdSpmiBLOz *quorum.HPOWa1PT0xzq
	if cu1RSg2YR.bQsSnXJ.LwVL87 == 0 {
		cWdSpmiBLOz = cu1RSg2YR.bQsSnXJ
	} else if cu1RSg2YR.bQsSnXJ.BlockHeight == cu1RSg2YR.ta2Y2hhO.BlockHeight {
		cWdSpmiBLOz = <-cu1RSg2YR.acQiWvFYDH
	} else {
		cWdSpmiBLOz = cu1RSg2YR.bQsSnXJ
		<-cu1RSg2YR.acQiWvFYDH
	}

	ouJpb2aF3mlv := /*line dqeCgmtgQvkX.go:1*/ blockchain.AiKaUM7(
		ndnWqNS,
		/*line yRbzaFa5.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch(),
		/*line w4wJOQP.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurView(),
		/*line s1X1dHT.go:1*/ cu1RSg2YR.GetLastBlockHeight(),
		/*line HD2RLAUL.go:1*/ cu1RSg2YR.iSgBUFf9.GetCurrentStateHash(),
		cWdSpmiBLOz.ZWsU_2,
		/*line FFvW8VTy.go:1*/ cu1RSg2YR.GetHighBlockHeight(),
		cWdSpmiBLOz,
	)
	cu1RSg2YR.ta2Y2hhO = ouJpb2aF3mlv.(*blockchain.CoordinationBlock).QC

	return ouJpb2aF3mlv.(*blockchain.CoordinationBlock)
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessCoordinationBlock(ouJpb2aF3mlv *blockchain.CoordinationBlock) error {
	/*line H03xXNG.go:1*/ log.ViJSfja(func() /*line UKHai5ByPMaL.go:1*/ string {
		seed := /*line UKHai5ByPMaL.go:1*/ byte(113)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line UKHai5ByPMaL.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line UKHai5ByPMaL.go:1*/ fnc(204)(98)(21)(17)(229)(210)(204)(186)(113)(214)(174)(106)(212)(120)(28)(56)(115)(216)(181)(111)(209)(181)(95)(196)(135)(226)(238)(223)(178)(108)(150)(35)(150)(46)(89)(166)(78)(170)(84)(158)(65)(123)(175)(160)(74)(151)(34)(76)(77)(237)(219)(163)(87)(176)
		return /*line UKHai5ByPMaL.go:1*/ string(data)
	}(), /*line Eydsjjq.go:1*/ cu1RSg2YR.ID())

	hYhbNJBbO := /*line tOVBPfqIh3Yk.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch()
	r0VHvbK := /*line GXnAUPgMyhu.go:1*/ cu1RSg2YR.GetHighBlockHeight()

	/*line ak9YwAgo95f.go:1*/
	log.ViJSfja(func() /*line aQWIoY9U.go:1*/ string {
		data := /*line aQWIoY9U.go:1*/ make([]byte, 0, 102)
		i := 11
		decryptKey := 43
		for counter := 0; i != 12; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 33:
				i = 40
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8f\xd3\xd1"...,
				)
			case 20:
				data = /*line aQWIoY9U.go:1*/ append(data, "PF\x89\x8a"...,
				)
				i = 8
			case 28:
				data = /*line aQWIoY9U.go:1*/ append(data, 157)
				i = 31
			case 25:
				i = 21
				data = /*line aQWIoY9U.go:1*/ append(data, "\x82\x8f\x92\x87"...,
				)
			case 3:
				data = /*line aQWIoY9U.go:1*/ append(data, 152)
				i = 0
			case 40:
				i = 19
				data = /*line aQWIoY9U.go:1*/ append(data, 207)
			case 23:
				i = 6
				data = /*line aQWIoY9U.go:1*/ append(data, 213)
			case 16:
				i = 23
				data = /*line aQWIoY9U.go:1*/ append(data, "\xa2\x9c{\x7f"...,
				)
			case 6:
				for y := range data {
					data[y] = data[y] - /*line aQWIoY9U.go:1*/ byte(decryptKey^y)
				}
				i = 12
			case 29:
				i = 24
				data = /*line aQWIoY9U.go:1*/ append(data, "]d"...,
				)
			case 36:
				data = /*line aQWIoY9U.go:1*/ append(data, "n}"...,
				)
				i = 10
			case 4:
				i = 3
				data = /*line aQWIoY9U.go:1*/ append(data, "\x9b\x8c\x94"...,
				)
			case 18:
				data = /*line aQWIoY9U.go:1*/ append(data, "z,{"...,
				)
				i = 15
			case 1:
				i = 16
				data = /*line aQWIoY9U.go:1*/ append(data, "݆"...,
				)
			case 14:
				i = 9
				data = /*line aQWIoY9U.go:1*/ append(data, "\xc4\xe4\xe7"...,
				)
			case 10:
				i = 18
				data = /*line aQWIoY9U.go:1*/ append(data, 121)
			case 17:
				data = /*line aQWIoY9U.go:1*/ append(data, "ht%"...,
				)
				i = 39
			case 15:
				i = 41
				data = /*line aQWIoY9U.go:1*/ append(data, 115)
			case 8:
				i = 25
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8a}"...,
				)
			case 5:
				i = 30
				data = /*line aQWIoY9U.go:1*/ append(data, 129)
			case 19:
				i = 1
				data = /*line aQWIoY9U.go:1*/ append(data, "\xc6ʅ\x89"...,
				)
			case 32:
				i = 35
				data = /*line aQWIoY9U.go:1*/ append(data, "\xa0\x93"...,
				)
			case 35:
				i = 34
				data = /*line aQWIoY9U.go:1*/ append(data, "\x98\xa5\xa8w"...,
				)
			case 31:
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8fb"...,
				)
				i = 37
			case 26:
				data = /*line aQWIoY9U.go:1*/ append(data, "\x81\x83z"...,
				)
				i = 5
			case 27:
				i = 13
				data = /*line aQWIoY9U.go:1*/ append(data, "\x9a\xea\x97\xec"...,
				)
			case 39:
				data = /*line aQWIoY9U.go:1*/ append(data, 41)
				i = 22
			case 0:
				i = 28
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8e\xa0\x98"...,
				)
			case 24:
				i = 32
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8f\xb0"...,
				)
			case 30:
				data = /*line aQWIoY9U.go:1*/ append(data, 41)
				i = 36
			case 9:
				i = 27
				data = /*line aQWIoY9U.go:1*/ append(data, "\xd8\xd8\xe7\x92"...,
				)
			case 7:
				i = 14
				data = /*line aQWIoY9U.go:1*/ append(data, "\xe4\xea\xdd\xe8"...,
				)
			case 34:
				i = 4
				data = /*line aQWIoY9U.go:1*/ append(data, "\xa6\xa5"...,
				)
			case 22:
				data = /*line aQWIoY9U.go:1*/ append(data, "}&\xdb"...,
				)
				i = 7
			case 21:
				i = 26
				data = /*line aQWIoY9U.go:1*/ append(data, "\x7fw3t"...,
				)
			case 37:
				i = 20
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8f\x91\x88\x8f"...,
				)
			case 38:
				data = /*line aQWIoY9U.go:1*/ append(data, 226)
				i = 33
			case 2:
				data = /*line aQWIoY9U.go:1*/ append(data, "\x8a\x92"...,
				)
				i = 38
			case 13:
				i = 2
				data = /*line aQWIoY9U.go:1*/ append(data, "\xd2\xcd\xe2"...,
				)
			case 11:
				data = /*line aQWIoY9U.go:1*/ append(data, "\x94]\xb1\x97"...,
				)
				i = 29
			case 41:
				i = 17
				data = /*line aQWIoY9U.go:1*/ append(data, "bd"...,
				)
			}
		}
		return /*line aQWIoY9U.go:1*/ string(data)
	}(), /*line aHQw9gc4aBwp.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash)

	if ouJpb2aF3mlv.BlockHeader.Epoch > hYhbNJBbO {
		/*line ivLgZvuLxEO.go:1*/ utils.OUWLEOgmMfR(cu1RSg2YR.wdUdVrSEU, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, ouJpb2aF3mlv)

		return nil
	}

	k1qj7vimKIy7 := /*line DjHUqA_Fi5.go:1*/ cu1RSg2YR.zV6vwpsi9W1(ouJpb2aF3mlv)
	if k1qj7vimKIy7 != nil {
		/*line clpsfLKSs.go:1*/ log.Jp980o6YjM(func() /*line lOqtpc.go:1*/ string {
			fullData := [] /*line lOqtpc.go:1*/ byte("7\xc7+\xda6\uf194>\xd8~\xac\x90\xf1K\xe3\xc8m\x19W@\xb1\x00\xee#$\xce\x7fa\x94\t4,\xdaHu\x9b\\\xeb\x80\xc6D\xff7X\xf57\xb9\xec\xa2@c\x9e!\x88Eʏ\xf7\x9d\x879\x17Z\x19b]%\x11\xf9Z\t\x88c\x15\xc9)\x99\xea\x97,\xf5\x1d\xc1H\xd0*@\xdb\xd0F\x16ϭ\x1c\xa7[\xbaߓsX\xa8\xb9\xc7\xe6å\xf0\xec9\x01\x9cX̨\x80\xd6\xee\xfa")
			idxKey := [] /*line lOqtpc.go:1*/ byte("Ɂ\xcd\x1f\xf9܅E\x06\x1dF\x8a\x81\xb5\x9e\x1d")
			data := /*line lOqtpc.go:1*/ make([]byte, 0, 61)
			data = /*line lOqtpc.go:1*/ append(data, fullData[143^ /*line lOqtpc.go:1*/ int(idxKey[11])]-fullData[151^ /*line lOqtpc.go:1*/ int(idxKey[11])], fullData[199^ /*line lOqtpc.go:1*/ int(idxKey[5])]-fullData[154^ /*line lOqtpc.go:1*/ int(idxKey[5])], fullData[157^ /*line lOqtpc.go:1*/ int(idxKey[6])]-fullData[216^ /*line lOqtpc.go:1*/ int(idxKey[6])], fullData[20^ /*line lOqtpc.go:1*/ int(idxKey[10])]+fullData[17^ /*line lOqtpc.go:1*/ int(idxKey[10])], fullData[113^ /*line lOqtpc.go:1*/ int(idxKey[9])]-fullData[72^ /*line lOqtpc.go:1*/ int(idxKey[9])], fullData[169^ /*line lOqtpc.go:1*/ int(idxKey[1])]-fullData[181^ /*line lOqtpc.go:1*/ int(idxKey[1])], fullData[136^ /*line lOqtpc.go:1*/ int(idxKey[1])]^fullData[183^ /*line lOqtpc.go:1*/ int(idxKey[1])], fullData[181^ /*line lOqtpc.go:1*/ int(idxKey[4])]^fullData[153^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[198^ /*line lOqtpc.go:1*/ int(idxKey[4])]+fullData[179^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[194^ /*line lOqtpc.go:1*/ int(idxKey[2])]-fullData[185^ /*line lOqtpc.go:1*/ int(idxKey[2])], fullData[199^ /*line lOqtpc.go:1*/ int(idxKey[11])]-fullData[149^ /*line lOqtpc.go:1*/ int(idxKey[11])], fullData[226^ /*line lOqtpc.go:1*/ int(idxKey[2])]+fullData[172^ /*line lOqtpc.go:1*/ int(idxKey[2])], fullData[89^ /*line lOqtpc.go:1*/ int(idxKey[8])]+fullData[116^ /*line lOqtpc.go:1*/ int(idxKey[8])], fullData[30^ /*line lOqtpc.go:1*/ int(idxKey[15])]-fullData[82^ /*line lOqtpc.go:1*/ int(idxKey[15])], fullData[104^ /*line lOqtpc.go:1*/ int(idxKey[10])]^fullData[106^ /*line lOqtpc.go:1*/ int(idxKey[10])], fullData[143^ /*line lOqtpc.go:1*/ int(idxKey[6])]+fullData[136^ /*line lOqtpc.go:1*/ int(idxKey[6])], fullData[187^ /*line lOqtpc.go:1*/ int(idxKey[14])]+fullData[197^ /*line lOqtpc.go:1*/ int(idxKey[14])], fullData[238^ /*line lOqtpc.go:1*/ int(idxKey[2])]-fullData[137^ /*line lOqtpc.go:1*/ int(idxKey[2])], fullData[202^ /*line lOqtpc.go:1*/ int(idxKey[14])]^fullData[171^ /*line lOqtpc.go:1*/ int(idxKey[14])], fullData[14^ /*line lOqtpc.go:1*/ int(idxKey[3])]-fullData[53^ /*line lOqtpc.go:1*/ int(idxKey[3])], fullData[78^ /*line lOqtpc.go:1*/ int(idxKey[3])]^fullData[24^ /*line lOqtpc.go:1*/ int(idxKey[3])], fullData[249^ /*line lOqtpc.go:1*/ int(idxKey[4])]-fullData[147^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[179^ /*line lOqtpc.go:1*/ int(idxKey[14])]^fullData[238^ /*line lOqtpc.go:1*/ int(idxKey[14])], fullData[170^ /*line lOqtpc.go:1*/ int(idxKey[0])]-fullData[208^ /*line lOqtpc.go:1*/ int(idxKey[0])], fullData[65^ /*line lOqtpc.go:1*/ int(idxKey[9])]-fullData[1^ /*line lOqtpc.go:1*/ int(idxKey[9])], fullData[177^ /*line lOqtpc.go:1*/ int(idxKey[11])]^fullData[232^ /*line lOqtpc.go:1*/ int(idxKey[11])], fullData[167^ /*line lOqtpc.go:1*/ int(idxKey[1])]^fullData[189^ /*line lOqtpc.go:1*/ int(idxKey[1])], fullData[159^ /*line lOqtpc.go:1*/ int(idxKey[4])]^fullData[248^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[227^ /*line lOqtpc.go:1*/ int(idxKey[13])]+fullData[219^ /*line lOqtpc.go:1*/ int(idxKey[13])], fullData[199^ /*line lOqtpc.go:1*/ int(idxKey[6])]^fullData[129^ /*line lOqtpc.go:1*/ int(idxKey[6])], fullData[235^ /*line lOqtpc.go:1*/ int(idxKey[5])]-fullData[130^ /*line lOqtpc.go:1*/ int(idxKey[5])], fullData[22^ /*line lOqtpc.go:1*/ int(idxKey[8])]+fullData[119^ /*line lOqtpc.go:1*/ int(idxKey[8])], fullData[37^ /*line lOqtpc.go:1*/ int(idxKey[9])]^fullData[22^ /*line lOqtpc.go:1*/ int(idxKey[9])], fullData[222^ /*line lOqtpc.go:1*/ int(idxKey[0])]+fullData[173^ /*line lOqtpc.go:1*/ int(idxKey[0])], fullData[21^ /*line lOqtpc.go:1*/ int(idxKey[15])]^fullData[14^ /*line lOqtpc.go:1*/ int(idxKey[15])], fullData[128^ /*line lOqtpc.go:1*/ int(idxKey[0])]+fullData[215^ /*line lOqtpc.go:1*/ int(idxKey[0])], fullData[83^ /*line lOqtpc.go:1*/ int(idxKey[7])]-fullData[97^ /*line lOqtpc.go:1*/ int(idxKey[7])], fullData[155^ /*line lOqtpc.go:1*/ int(idxKey[5])]-fullData[183^ /*line lOqtpc.go:1*/ int(idxKey[5])], fullData[35^ /*line lOqtpc.go:1*/ int(idxKey[15])]-fullData[39^ /*line lOqtpc.go:1*/ int(idxKey[15])], fullData[97^ /*line lOqtpc.go:1*/ int(idxKey[8])]^fullData[39^ /*line lOqtpc.go:1*/ int(idxKey[8])], fullData[187^ /*line lOqtpc.go:1*/ int(idxKey[13])]-fullData[220^ /*line lOqtpc.go:1*/ int(idxKey[13])], fullData[253^ /*line lOqtpc.go:1*/ int(idxKey[11])]^fullData[194^ /*line lOqtpc.go:1*/ int(idxKey[11])], fullData[63^ /*line lOqtpc.go:1*/ int(idxKey[3])]+fullData[61^ /*line lOqtpc.go:1*/ int(idxKey[3])], fullData[52^ /*line lOqtpc.go:1*/ int(idxKey[15])]-fullData[69^ /*line lOqtpc.go:1*/ int(idxKey[15])], fullData[208^ /*line lOqtpc.go:1*/ int(idxKey[5])]+fullData[169^ /*line lOqtpc.go:1*/ int(idxKey[5])], fullData[226^ /*line lOqtpc.go:1*/ int(idxKey[0])]-fullData[211^ /*line lOqtpc.go:1*/ int(idxKey[0])], fullData[43^ /*line lOqtpc.go:1*/ int(idxKey[10])]^fullData[127^ /*line lOqtpc.go:1*/ int(idxKey[10])], fullData[80^ /*line lOqtpc.go:1*/ int(idxKey[7])]^fullData[28^ /*line lOqtpc.go:1*/ int(idxKey[7])], fullData[229^ /*line lOqtpc.go:1*/ int(idxKey[13])]^fullData[208^ /*line lOqtpc.go:1*/ int(idxKey[13])], fullData[172^ /*line lOqtpc.go:1*/ int(idxKey[14])]+fullData[221^ /*line lOqtpc.go:1*/ int(idxKey[14])], fullData[196^ /*line lOqtpc.go:1*/ int(idxKey[4])]-fullData[185^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[179^ /*line lOqtpc.go:1*/ int(idxKey[5])]^fullData[239^ /*line lOqtpc.go:1*/ int(idxKey[5])], fullData[72^ /*line lOqtpc.go:1*/ int(idxKey[8])]^fullData[0^ /*line lOqtpc.go:1*/ int(idxKey[8])], fullData[145^ /*line lOqtpc.go:1*/ int(idxKey[4])]+fullData[138^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[200^ /*line lOqtpc.go:1*/ int(idxKey[4])]+fullData[170^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[138^ /*line lOqtpc.go:1*/ int(idxKey[14])]+fullData[156^ /*line lOqtpc.go:1*/ int(idxKey[14])], fullData[222^ /*line lOqtpc.go:1*/ int(idxKey[4])]-fullData[163^ /*line lOqtpc.go:1*/ int(idxKey[4])], fullData[13^ /*line lOqtpc.go:1*/ int(idxKey[3])]-fullData[90^ /*line lOqtpc.go:1*/ int(idxKey[3])], fullData[112^ /*line lOqtpc.go:1*/ int(idxKey[8])]-fullData[77^ /*line lOqtpc.go:1*/ int(idxKey[8])], fullData[203^ /*line lOqtpc.go:1*/ int(idxKey[11])]-fullData[186^ /*line lOqtpc.go:1*/ int(idxKey[11])])
			return /*line lOqtpc.go:1*/ string(data)
		}(), /*line bnOwIT.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)

		return k1qj7vimKIy7
	}

	if ouJpb2aF3mlv.BlockHeader.BlockHeight > r0VHvbK+1 {

		/*line ei80aH.go:1*/
		utils.OUWLEOgmMfR(cu1RSg2YR.wdUdVrSEU, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, ouJpb2aF3mlv)

		return nil
	} else if /*line jGIGjv4.go:1*/ cu1RSg2YR.State() == types.VIEWCHANGING {
		/*line ddZvhwtQI7KS.go:1*/ utils.OUWLEOgmMfR(cu1RSg2YR.wdUdVrSEU, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv)

		return nil
	}

	if ouJpb2aF3mlv.BlockHeader.StateRoot != /*line HaIb4qXEwZ6.go:1*/ cu1RSg2YR.iSgBUFf9.GetCurrentStateHash() {

		/*line EySoiOE05iYa.go:1*/
		log.WPP4KjwN(func() /*line v6K9gaC.go:1*/ string {
			key := [] /*line v6K9gaC.go:1*/ byte("\x1e\xc90>\xbdc\b^\x99\x04\bϐB\xcfj\xfa\xe6ҙzB\u0080\xe3VD\x02\xc1\xdf\xe1\xc2G\xcc\xe6\x15\x1f\x9e\x11\xbfa\xfd&\"\x9dt\x1c\xbf\x16O\xa8\x14\xda\xfb\xbcR")
			data := [] /*line v6K9gaC.go:1*/ byte("y\ue99b\u074bX\xd0\bgmB\x03\x99>\xdceKD\xdb\xe6\xb1%\xeb\fv\xb7v\"SF\n\xa8?N5\x8d\r\x85\xdf\xd7^\x92\x8b\x01\x94~+\x85\xb2\x13]\x1e\x1b\xe1\xca")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line v6K9gaC.go:1*/ string(data)
		}(), /*line IsLXcfxSjIKL.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)

		return nil
	}

	/*line JGPJYHX.go:1*/
	cu1RSg2YR.SetState(types.PREPARED)

	/*line b9YAAcqbF.go:1*/
	utils.OUWLEOgmMfR(cu1RSg2YR.cA3KSEC5m, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv)

	cWdSpmiBLOz, i3KmQgjb := cu1RSg2YR.vHhOhURml[ouJpb2aF3mlv.BlockHash]
	if i3KmQgjb {
		/*line o7B5TJ.go:1*/ cu1RSg2YR.aVd2nRenj(cWdSpmiBLOz)
		/*line E5ZxGletn.go:1*/ delete(cu1RSg2YR.vHhOhURml, ouJpb2aF3mlv.BlockHash)
		qd9YlUhDZv7a := /*line G5Fmg36y.go:1*/ quorum.G71jC_Q[quorum.Vote](
			ouJpb2aF3mlv.BlockHeader.Epoch,
			ouJpb2aF3mlv.BlockHeader.View,
			ouJpb2aF3mlv.BlockHeader.BlockHeight,
			/*line XktO9Nx65.go:1*/ cu1RSg2YR.ID(),
			ouJpb2aF3mlv.BlockHash,
		)

		/*line Grokton.go:1*/
		cu1RSg2YR.BroadcastToSome( /*line BazHNWNy.go:1*/ cu1RSg2YR.FindCommitteesFor(ouJpb2aF3mlv.BlockHeader.Epoch), qd9YlUhDZv7a)

		return nil
	}

	ae4RXEhV, k1qj7vimKIy7 := /*line ljyKAe.go:1*/ cu1RSg2YR.uS1zadl60yhO(ouJpb2aF3mlv)
	if k1qj7vimKIy7 != nil {

		/*line KdqzK_2R.go:1*/
		log.Jp980o6YjM(func() /*line H6QK5jT.go:1*/ string {
			fullData := [] /*line H6QK5jT.go:1*/ byte("4_\xa8\x1cQ\xf9w\xb5M\x0f\x18&\xc5v\xecj\xd5\xf5\x80\xfc\xdeDZk\xa0x\x9e\n1\x8d\xcf5\xdcMm\xc8\xe5\x0f4b(\xda\xfdX\xe3\xed\\ɨ\xc2B\xefv\x801\xe8\xbbO\bi\xa5\xf6\xf0\x82\xf0\xb0\xfd\x9e\xeem$c\xba\x85ڒj\x12x\xf2\xfe\xea\xd5\ti\xbd~\xfa\xf5\x8c0#\xf7\xe0\xc1R\x8c\x84\xfa\xb2\xf6z:\xfd9\x9b\x1cIqIh\xce~\xe8d\xf5\x12\xe6\x8dudn#\x9c\x8f\xa8\xa7:\xaf\x11T\xfd\xa0e\x82\xb1\x944")
			idxKey := [] /*line H6QK5jT.go:1*/ byte("}^\x19N\x9e@\xcd\xf0jI\xa1\xb4\xd7D\xa0\xf0")
			data := /*line H6QK5jT.go:1*/ make([]byte, 0, 70)
			data = /*line H6QK5jT.go:1*/ append(data, fullData[207^ /*line H6QK5jT.go:1*/ int(idxKey[4])]^fullData[25^ /*line H6QK5jT.go:1*/ int(idxKey[4])], fullData[161^ /*line H6QK5jT.go:1*/ int(idxKey[14])]-fullData[223^ /*line H6QK5jT.go:1*/ int(idxKey[14])], fullData[182^ /*line H6QK5jT.go:1*/ int(idxKey[15])]^fullData[175^ /*line H6QK5jT.go:1*/ int(idxKey[15])], fullData[193^ /*line H6QK5jT.go:1*/ int(idxKey[9])]^fullData[102^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[75^ /*line H6QK5jT.go:1*/ int(idxKey[9])]+fullData[7^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[28^ /*line H6QK5jT.go:1*/ int(idxKey[1])]-fullData[12^ /*line H6QK5jT.go:1*/ int(idxKey[1])], fullData[173^ /*line H6QK5jT.go:1*/ int(idxKey[6])]^fullData[237^ /*line H6QK5jT.go:1*/ int(idxKey[6])], fullData[32^ /*line H6QK5jT.go:1*/ int(idxKey[3])]-fullData[115^ /*line H6QK5jT.go:1*/ int(idxKey[3])], fullData[166^ /*line H6QK5jT.go:1*/ int(idxKey[11])]-fullData[53^ /*line H6QK5jT.go:1*/ int(idxKey[11])], fullData[53^ /*line H6QK5jT.go:1*/ int(idxKey[2])]+fullData[44^ /*line H6QK5jT.go:1*/ int(idxKey[2])], fullData[157^ /*line H6QK5jT.go:1*/ int(idxKey[4])]+fullData[243^ /*line H6QK5jT.go:1*/ int(idxKey[4])], fullData[79^ /*line H6QK5jT.go:1*/ int(idxKey[9])]+fullData[90^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[16^ /*line H6QK5jT.go:1*/ int(idxKey[3])]^fullData[45^ /*line H6QK5jT.go:1*/ int(idxKey[3])], fullData[236^ /*line H6QK5jT.go:1*/ int(idxKey[12])]+fullData[147^ /*line H6QK5jT.go:1*/ int(idxKey[12])], fullData[103^ /*line H6QK5jT.go:1*/ int(idxKey[2])]^fullData[58^ /*line H6QK5jT.go:1*/ int(idxKey[2])], fullData[183^ /*line H6QK5jT.go:1*/ int(idxKey[15])]+fullData[249^ /*line H6QK5jT.go:1*/ int(idxKey[15])], fullData[197^ /*line H6QK5jT.go:1*/ int(idxKey[10])]+fullData[214^ /*line H6QK5jT.go:1*/ int(idxKey[10])], fullData[113^ /*line H6QK5jT.go:1*/ int(idxKey[0])]+fullData[101^ /*line H6QK5jT.go:1*/ int(idxKey[0])], fullData[129^ /*line H6QK5jT.go:1*/ int(idxKey[7])]-fullData[253^ /*line H6QK5jT.go:1*/ int(idxKey[7])], fullData[125^ /*line H6QK5jT.go:1*/ int(idxKey[9])]-fullData[192^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[246^ /*line H6QK5jT.go:1*/ int(idxKey[14])]-fullData[212^ /*line H6QK5jT.go:1*/ int(idxKey[14])], fullData[179^ /*line H6QK5jT.go:1*/ int(idxKey[11])]+fullData[252^ /*line H6QK5jT.go:1*/ int(idxKey[11])], fullData[175^ /*line H6QK5jT.go:1*/ int(idxKey[14])]+fullData[165^ /*line H6QK5jT.go:1*/ int(idxKey[14])], fullData[189^ /*line H6QK5jT.go:1*/ int(idxKey[12])]+fullData[238^ /*line H6QK5jT.go:1*/ int(idxKey[12])], fullData[152^ /*line H6QK5jT.go:1*/ int(idxKey[14])]-fullData[235^ /*line H6QK5jT.go:1*/ int(idxKey[14])], fullData[4^ /*line H6QK5jT.go:1*/ int(idxKey[1])]+fullData[96^ /*line H6QK5jT.go:1*/ int(idxKey[1])], fullData[188^ /*line H6QK5jT.go:1*/ int(idxKey[15])]^fullData[163^ /*line H6QK5jT.go:1*/ int(idxKey[15])], fullData[228^ /*line H6QK5jT.go:1*/ int(idxKey[4])]-fullData[175^ /*line H6QK5jT.go:1*/ int(idxKey[4])], fullData[148^ /*line H6QK5jT.go:1*/ int(idxKey[12])]^fullData[151^ /*line H6QK5jT.go:1*/ int(idxKey[12])], fullData[30^ /*line H6QK5jT.go:1*/ int(idxKey[9])]-fullData[16^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[49^ /*line H6QK5jT.go:1*/ int(idxKey[9])]-fullData[17^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[250^ /*line H6QK5jT.go:1*/ int(idxKey[7])]+fullData[222^ /*line H6QK5jT.go:1*/ int(idxKey[7])], fullData[59^ /*line H6QK5jT.go:1*/ int(idxKey[9])]^fullData[92^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[123^ /*line H6QK5jT.go:1*/ int(idxKey[9])]-fullData[93^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[121^ /*line H6QK5jT.go:1*/ int(idxKey[9])]+fullData[28^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[4^ /*line H6QK5jT.go:1*/ int(idxKey[0])]+fullData[14^ /*line H6QK5jT.go:1*/ int(idxKey[0])], fullData[155^ /*line H6QK5jT.go:1*/ int(idxKey[7])]-fullData[173^ /*line H6QK5jT.go:1*/ int(idxKey[7])], fullData[127^ /*line H6QK5jT.go:1*/ int(idxKey[5])]^fullData[53^ /*line H6QK5jT.go:1*/ int(idxKey[5])], fullData[29^ /*line H6QK5jT.go:1*/ int(idxKey[2])]^fullData[63^ /*line H6QK5jT.go:1*/ int(idxKey[2])], fullData[112^ /*line H6QK5jT.go:1*/ int(idxKey[15])]-fullData[140^ /*line H6QK5jT.go:1*/ int(idxKey[15])], fullData[245^ /*line H6QK5jT.go:1*/ int(idxKey[11])]-fullData[220^ /*line H6QK5jT.go:1*/ int(idxKey[11])], fullData[212^ /*line H6QK5jT.go:1*/ int(idxKey[15])]^fullData[134^ /*line H6QK5jT.go:1*/ int(idxKey[15])], fullData[207^ /*line H6QK5jT.go:1*/ int(idxKey[14])]-fullData[244^ /*line H6QK5jT.go:1*/ int(idxKey[14])], fullData[237^ /*line H6QK5jT.go:1*/ int(idxKey[14])]+fullData[135^ /*line H6QK5jT.go:1*/ int(idxKey[14])], fullData[37^ /*line H6QK5jT.go:1*/ int(idxKey[13])]^fullData[74^ /*line H6QK5jT.go:1*/ int(idxKey[13])], fullData[86^ /*line H6QK5jT.go:1*/ int(idxKey[2])]-fullData[4^ /*line H6QK5jT.go:1*/ int(idxKey[2])], fullData[84^ /*line H6QK5jT.go:1*/ int(idxKey[0])]^fullData[0^ /*line H6QK5jT.go:1*/ int(idxKey[0])], fullData[108^ /*line H6QK5jT.go:1*/ int(idxKey[0])]^fullData[109^ /*line H6QK5jT.go:1*/ int(idxKey[0])], fullData[252^ /*line H6QK5jT.go:1*/ int(idxKey[4])]+fullData[251^ /*line H6QK5jT.go:1*/ int(idxKey[4])], fullData[238^ /*line H6QK5jT.go:1*/ int(idxKey[4])]-fullData[187^ /*line H6QK5jT.go:1*/ int(idxKey[4])], fullData[202^ /*line H6QK5jT.go:1*/ int(idxKey[15])]^fullData[216^ /*line H6QK5jT.go:1*/ int(idxKey[15])], fullData[224^ /*line H6QK5jT.go:1*/ int(idxKey[12])]^fullData[205^ /*line H6QK5jT.go:1*/ int(idxKey[12])], fullData[24^ /*line H6QK5jT.go:1*/ int(idxKey[4])]+fullData[179^ /*line H6QK5jT.go:1*/ int(idxKey[4])], fullData[228^ /*line H6QK5jT.go:1*/ int(idxKey[12])]^fullData[190^ /*line H6QK5jT.go:1*/ int(idxKey[12])], fullData[198^ /*line H6QK5jT.go:1*/ int(idxKey[13])]^fullData[114^ /*line H6QK5jT.go:1*/ int(idxKey[13])], fullData[107^ /*line H6QK5jT.go:1*/ int(idxKey[9])]-fullData[104^ /*line H6QK5jT.go:1*/ int(idxKey[9])], fullData[34^ /*line H6QK5jT.go:1*/ int(idxKey[3])]-fullData[100^ /*line H6QK5jT.go:1*/ int(idxKey[3])], fullData[35^ /*line H6QK5jT.go:1*/ int(idxKey[13])]+fullData[83^ /*line H6QK5jT.go:1*/ int(idxKey[13])], fullData[214^ /*line H6QK5jT.go:1*/ int(idxKey[6])]-fullData[241^ /*line H6QK5jT.go:1*/ int(idxKey[6])], fullData[150^ /*line H6QK5jT.go:1*/ int(idxKey[6])]+fullData[78^ /*line H6QK5jT.go:1*/ int(idxKey[6])], fullData[186^ /*line H6QK5jT.go:1*/ int(idxKey[7])]-fullData[233^ /*line H6QK5jT.go:1*/ int(idxKey[7])], fullData[219^ /*line H6QK5jT.go:1*/ int(idxKey[7])]^fullData[240^ /*line H6QK5jT.go:1*/ int(idxKey[7])], fullData[183^ /*line H6QK5jT.go:1*/ int(idxKey[10])]^fullData[190^ /*line H6QK5jT.go:1*/ int(idxKey[10])], fullData[197^ /*line H6QK5jT.go:1*/ int(idxKey[5])]+fullData[16^ /*line H6QK5jT.go:1*/ int(idxKey[5])], fullData[177^ /*line H6QK5jT.go:1*/ int(idxKey[12])]-fullData[201^ /*line H6QK5jT.go:1*/ int(idxKey[12])], fullData[97^ /*line H6QK5jT.go:1*/ int(idxKey[0])]-fullData[33^ /*line H6QK5jT.go:1*/ int(idxKey[0])], fullData[197^ /*line H6QK5jT.go:1*/ int(idxKey[6])]^fullData[136^ /*line H6QK5jT.go:1*/ int(idxKey[6])], fullData[13^ /*line H6QK5jT.go:1*/ int(idxKey[13])]^fullData[192^ /*line H6QK5jT.go:1*/ int(idxKey[13])], fullData[17^ /*line H6QK5jT.go:1*/ int(idxKey[8])]-fullData[97^ /*line H6QK5jT.go:1*/ int(idxKey[8])])
			return /*line H6QK5jT.go:1*/ string(data)
		}(), /*line emfFgiJguRGv.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)

		return k1qj7vimKIy7
	}
	if !ae4RXEhV {

		/*line kOAhoEYKeTrW.go:1*/
		log.WPP4KjwN(func() /*line iMK2Ib7Ei3.go:1*/ string {
			data := [] /*line iMK2Ib7Ei3.go:1*/ byte("[\x1dt]\"\x1cF$#ceTseo(_\xff\f\x19l\x14C\xa1\x00y^\x1b\xd2nZ\x9e\x12goi&\xeb \xec\xb1\v\xf9\xf65e\x7ffPv bE\x11\x1ad 5D2\xf5j")
			positions := [...]byte{22, 22, 16, 57, 41, 46, 40, 31, 6, 11, 19, 17, 1, 36, 27, 6, 13, 49, 8, 27, 16, 43, 15, 30, 7, 16, 27, 26, 53, 18, 48, 44, 24, 52, 49, 21, 59, 61, 19, 40, 36, 15, 28, 42, 7, 60, 4, 2, 42, 54, 1, 1, 28, 25, 26, 5, 17, 23, 27, 37, 41, 60, 43, 44, 31, 39, 32, 55, 17, 37, 13, 15, 54, 57, 40, 32}
			for i := 0; i < 76; i += 2 {
				localKey := /*line iMK2Ib7Ei3.go:1*/ byte(i) + /*line iMK2Ib7Ei3.go:1*/ byte(positions[i]^positions[i+1]) + 32
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line iMK2Ib7Ei3.go:1*/ string(data)
		}(), /*line A6ZQnu.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)

		return nil
	}

	qd9YlUhDZv7a := /*line MlnRYn.go:1*/ quorum.G71jC_Q[quorum.Vote](
		ouJpb2aF3mlv.BlockHeader.Epoch,
		ouJpb2aF3mlv.BlockHeader.View,
		ouJpb2aF3mlv.BlockHeader.BlockHeight,
		/*line cqFaIX4.go:1*/ cu1RSg2YR.ID(),
		ouJpb2aF3mlv.BlockHash,
	)

	/*line dfdqmIG.go:1*/
	cu1RSg2YR.BroadcastToSome( /*line eeKaWKS.go:1*/ cu1RSg2YR.FindCommitteesFor(ouJpb2aF3mlv.BlockHeader.Epoch), qd9YlUhDZv7a)
	/*line NLG2JAt9PE.go:1*/ cu1RSg2YR.ProcessVote(qd9YlUhDZv7a)

	return nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessVote(qd9YlUhDZv7a *quorum.Collection[quorum.Vote]) {

	/*line _llNoM.go:1*/
	log.ViJSfja(func() /*line qVFyOcadnky.go:1*/ string {
		data := [] /*line qVFyOcadnky.go:1*/ byte("\v;vj\x18\xc8broa\xfdbsVo\xa9\x99: ;ro\xd8\xf3s\xa1\x12\x9d\xe2\xd0\xc6>5{)Gor!b\x9do\xef\xe2\re\t\xd8htg%\xb6\x1c\xdeyKw %v\t\x10~o\x87\xd0 \x06A Y\xd4\xf2\t\x0f\a\a@\xd0h % ")
		positions := [...]byte{72, 72, 71, 38, 55, 52, 73, 73, 62, 69, 27, 74, 3, 15, 66, 4, 16, 9, 42, 15, 26, 54, 25, 23, 0, 46, 75, 50, 33, 40, 61, 76, 27, 69, 79, 26, 44, 68, 77, 17, 34, 28, 72, 52, 74, 26, 15, 83, 33, 10, 69, 63, 5, 29, 19, 32, 38, 76, 53, 22, 26, 9, 80, 11, 35, 31, 47, 53, 71, 62, 23, 30, 62, 15, 72, 3, 6, 78, 23, 3, 61, 43, 44, 40, 16, 33, 65, 32, 0, 77, 1, 33, 76, 55, 73, 56, 27, 43}
		for i := 0; i < 98; i += 2 {
			localKey := /*line qVFyOcadnky.go:1*/ byte(i) + /*line qVFyOcadnky.go:1*/ byte(positions[i]^positions[i+1]) + 92
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line qVFyOcadnky.go:1*/ string(data)
	}(), /*line wmRKVTHNNR_.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.BlockHeight, qd9YlUhDZv7a.View, qd9YlUhDZv7a.Epoch, qd9YlUhDZv7a.NaNYri773M, qd9YlUhDZv7a.WpsY_aZ)

	cAIldNjRExFl, k1qj7vimKIy7 := /*line Gcz_cCOn_.go:1*/ crypto.SkrCuscT(qd9YlUhDZv7a.MqlfmVE9d /*line safkpXmh.go:1*/, crypto.IDToByte(qd9YlUhDZv7a.NaNYri773M))
	if k1qj7vimKIy7 != nil {
		/*line G3lzAQoE.go:1*/ log.Jp980o6YjM(func() /*line E6Yi7tqrZnZ.go:1*/ string {
			data := /*line E6Yi7tqrZnZ.go:1*/ make([]byte, 0, 66)
			i := 28
			decryptKey := 157
			for counter := 0; i != 17; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 159)
					i = 23
				case 21:
					i = 12
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 125)
				case 8:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x98\x91"...,
					)
					i = 27
				case 20:
					i = 26
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "}_{\x7f"...,
					)
				case 19:
					i = 22
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x89A\x98"...,
					)
				case 25:
					i = 7
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x9dW\x9b"...,
					)
				case 11:
					i = 13
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 145)
				case 27:
					i = 2
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x97\x8d"...,
					)
				case 28:
					i = 5
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 97)
				case 26:
					i = 10
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "{>"...,
					)
				case 31:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x8e\x84\x8c\x9e"...,
					)
					i = 11
				case 18:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "d\")"...,
					)
					i = 6
				case 4:
					i = 21
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\xa5_\x83"...,
					)
				case 9:
					i = 20
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "pu\x82"...,
					)
				case 6:
					i = 14
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 84)
				case 16:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\xa9\xad\xb1"...,
					)
					i = 4
				case 24:
					i = 3
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x83\x85>"...,
					)
				case 22:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x8b\x93"...,
					)
					i = 15
				case 15:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "M\xa3"...,
					)
					i = 8
				case 14:
					i = 9
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "u}"...,
					)
				case 0:
					i = 29
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "`\xbe"...,
					)
				case 10:
					i = 24
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "8|\x84\x83"...,
					)
				case 2:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x9f\xab\xa7"...,
					)
					i = 25
				case 13:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 149)
					i = 19
				case 23:
					i = 16
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 84)
				case 3:
					i = 1
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x86\x8e"...,
					)
				case 5:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "*~"...,
					)
					i = 18
				case 12:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 92)
					i = 0
				case 1:
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, 63)
					i = 30
				case 30:
					i = 31
					data = /*line E6Yi7tqrZnZ.go:1*/ append(data, "\x90~"...,
					)
				case 29:
					i = 17
					for y := range data {
						data[y] = data[y] + /*line E6Yi7tqrZnZ.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return /*line E6Yi7tqrZnZ.go:1*/ string(data)
		}(), /*line HTniaWT2Q.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.NaNYri773M)

		return
	}
	if !cAIldNjRExFl {
		/*line jSpyuG2QzC_J.go:1*/ log.WPP4KjwN(func() /*line HhDLXczSs6.go:1*/ string {
			data := [] /*line HhDLXczSs6.go:1*/ byte("[\xdbv\x80 \x86\xc9$\xcfc/?\x81\xdeot\xe6ǭ\x00M\x1b\xdf\xd0v(dK\x01\" oJ\xed\x04\x0e\x1d\xb7h\v\x8dn\xad\xf2ii\x01\xf8s\xbegD\xa4t\x1e;-\xf0Io\x90k\x13")
			positions := [...]byte{7, 49, 42, 46, 3, 23, 5, 17, 22, 7, 13, 34, 62, 18, 60, 40, 37, 39, 12, 33, 52, 21, 8, 30, 1, 6, 57, 17, 12, 43, 56, 23, 30, 10, 35, 33, 27, 54, 62, 27, 28, 32, 46, 36, 3, 12, 47, 60, 60, 40, 21, 29, 16, 29, 20, 11, 32, 51, 25, 19, 7, 59, 51, 6, 23, 44, 56, 18, 51, 55, 19, 61, 13, 25, 3, 44, 61, 35}
			for i := 0; i < 78; i += 2 {
				localKey := /*line HhDLXczSs6.go:1*/ byte(i) + /*line HhDLXczSs6.go:1*/ byte(positions[i]^positions[i+1]) + 133
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line HhDLXczSs6.go:1*/ string(data)
		}(), /*line WrYMr5MDS0.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.NaNYri773M)

		return
	}

	qNDcZn, cWdSpmiBLOz := /*line C1wzwwmpv3in.go:1*/ cu1RSg2YR.ffpSbXbMEX0.Add(qd9YlUhDZv7a)
	if !qNDcZn {
		/*line G8U1Wz.go:1*/ log.ViJSfja(func() /*line DYNk5R.go:1*/ string {
			key := [] /*line DYNk5R.go:1*/ byte("\x85ƳH*\x83h\xd63[\xf4+\x05\x00\xe6\xf1\xb6\xe8M\x069f\x84\xb9\x94\x83\xfa\ak\xa7\x11K\xfbEp\x84c\x96\xba\x81鉺ų\xbcG\xa2\xcc\t\xb9\xa5V\xe1\xd9O\ap`\xa9\xda\xf8\xbf")
			data := [] /*line DYNk5R.go:1*/ byte("\xe0\xeb)\xa5J\xab\xb8H\xa2\xbeY\x9exVUe\x1b\x11m|\xa8\xda\xe9,\xb4\xe4ll\x8b\x15\x80\xbf\x1b\xb8\xe5\xea\xc9\xff\x1d\xeaN\xf7.\xe5'+g\x04Ar%\tvB\xf9\xc0j\x90\xa9\xed\xfa\x1d7")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line DYNk5R.go:1*/ string(data)
		}(), /*line jLQyCjb0.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.NaNYri773M)

		return
	}

	cWdSpmiBLOz.AVzGLiX4RWH = /*line arBqDVD_oyH.go:1*/ cu1RSg2YR.ID()

	_, kbgkwuEjk := cu1RSg2YR.cA3KSEC5m[cWdSpmiBLOz.EmFrce][cWdSpmiBLOz.LwVL87][cWdSpmiBLOz.BlockHeight]
	if !kbgkwuEjk {
		cu1RSg2YR.vHhOhURml[cWdSpmiBLOz.ZWsU_2] = cWdSpmiBLOz
		/*line BAbgqy.go:1*/ log.ViJSfja(func() /*line pBtX95a.go:1*/ string {
			key := [] /*line pBtX95a.go:1*/ byte("\xa9'\xad'\xd1\xce\xf69\x17_p\x7f2\x03\x92s\xa7uf\x067\xbc\xecb\xc6$n(\xf6#\xc0_\xff\x0e\x9d\xac\xdaJ}3 \xc2͔\xc9,\xf2\xa9w+7\xc0\x8c\xf5d\xe83ٵf\xb5\xf3\x86ܗ")
			data := [] /*line pBtX95a.go:1*/ byte("\xf2\x02\xdbz\xf1\xe6\xa6Kx<\x15\fAU\xfd\a\xc2\\Fr_\xd9\xcc\x13\xa5\x04\a[\xd6A\xb59\x99k\xefɾj\x1b\\R\xe2\xa3\xfb\xbd\f\x93\xdb\x05BA\xa5\xe8\xd5\x06\x84\\\xba\xdeF\xfc\xb7\xa6\xf9\xef")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line pBtX95a.go:1*/ string(data)
		}(), /*line im9zM88fq6I.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2)

		return
	}

	/*line urVjXqvvR4xn.go:1*/
	cu1RSg2YR.aVd2nRenj(cWdSpmiBLOz)

	/*line KVq7J6ub.go:1*/
	log.ViJSfja(func() /*line B1f2MhAHu.go:1*/ string {
		seed := /*line B1f2MhAHu.go:1*/ byte(37)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line B1f2MhAHu.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line B1f2MhAHu.go:1*/ fnc(128)(202)(229)(177)(37)(82)(204)(186)(113)(214)(174)(106)(212)(139)(47)(99)(183)(50)(91)(252)(251)(251)(241)(236)(205)(151)(45)(22)(124)(250)(241)(214)(174)(106)(212)(158)(65)(123)(175)(180)(97)(199)(127)(185)(184)(121)(245)(152)(114)(238)(223)(178)(108)(181)(135)(18)(34)(69)(150)(216)(181)(187)(32)(150)(31)(58)(134)(181)(111)(47)(8)(85)(181)(105)(198)(145)(218)(185)(195)(48)(137)(13)(246)(241)(53)
		return /*line B1f2MhAHu.go:1*/ string(data)
	}(), /*line xsjMW067_r.go:1*/ cu1RSg2YR.ID(), qd9YlUhDZv7a.BlockHeight, qd9YlUhDZv7a.View, qd9YlUhDZv7a.Epoch, qd9YlUhDZv7a.NaNYri773M)
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessCommit(gFIlzELaH *quorum.Collection[quorum.Commit]) {
	/*line gG24Q_itrCc.go:1*/ log.ViJSfja(func() /*line fYbaUB.go:1*/ string {
		data := [] /*line fYbaUB.go:1*/ byte("\x8f%v]\xf3(P\x12o\x95.&9C\x88\xccT\xa6t\xc6؞\xdaoc\r\x01\x19\xa7\xa6\x9eYg\x11$m\xd0t fo\x1b\xb2\x1c \x03\xf4\"\xea\xfai\xd7h#\xd7(v\x1d\x8c\x0fe\xc6Ӱq\x88/\xd2@\x8b\xa7\xa9\x9e\x1eRI\x1a \x01x")
		positions := [...]byte{59, 63, 34, 9, 65, 67, 71, 31, 36, 20, 44, 27, 48, 62, 68, 74, 21, 58, 78, 78, 21, 55, 74, 48, 76, 76, 64, 51, 45, 25, 20, 42, 72, 65, 7, 36, 46, 17, 22, 42, 10, 29, 73, 66, 33, 4, 26, 0, 46, 43, 31, 15, 11, 12, 29, 49, 47, 41, 34, 42, 30, 42, 65, 11, 19, 68, 48, 49, 57, 61, 22, 71, 32, 58, 30, 4, 51, 53, 54, 14, 70, 34, 71, 33, 51, 69, 30, 58, 67, 28, 10, 26, 62, 7, 59, 7, 74, 16}
		for i := 0; i < 98; i += 2 {
			localKey := /*line fYbaUB.go:1*/ byte(i) + /*line fYbaUB.go:1*/ byte(positions[i]^positions[i+1]) + 18
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line fYbaUB.go:1*/ string(data)
	}(), /*line aqgnwKJa.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.BlockHeight, gFIlzELaH.View, gFIlzELaH.Epoch, gFIlzELaH.NaNYri773M)
	rDMAJ7T, k1qj7vimKIy7 := /*line BhYLGjS6.go:1*/ crypto.SkrCuscT(gFIlzELaH.MqlfmVE9d /*line Oelrh9xi_.go:1*/, crypto.IDToByte(gFIlzELaH.NaNYri773M))
	if k1qj7vimKIy7 != nil {
		/*line mJhcrKqR.go:1*/ log.Jp980o6YjM(func() /*line G7F16R.go:1*/ string {
			data := /*line G7F16R.go:1*/ make([]byte, 0, 70)
			i := 23
			decryptKey := 57
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					data = /*line G7F16R.go:1*/ append(data, "\a\x1b"...,
					)
					i = 16
				case 15:
					i = 3
					data = /*line G7F16R.go:1*/ append(data, "=;1"...,
					)
				case 7:
					i = 10
					data = /*line G7F16R.go:1*/ append(data, ":oy"...,
					)
				case 30:
					for y := range data {
						data[y] = data[y] ^ /*line G7F16R.go:1*/ byte(decryptKey^y)
					}
					i = 4
				case 18:
					data = /*line G7F16R.go:1*/ append(data, "\x0f>"...,
					)
					i = 22
				case 27:
					data = /*line G7F16R.go:1*/ append(data, 63)
					i = 29
				case 28:
					data = /*line G7F16R.go:1*/ append(data, "\x05\x17\x1a\x1f"...,
					)
					i = 5
				case 3:
					data = /*line G7F16R.go:1*/ append(data, "w,1"...,
					)
					i = 27
				case 0:
					i = 20
					data = /*line G7F16R.go:1*/ append(data, 24)
				case 25:
					i = 6
					data = /*line G7F16R.go:1*/ append(data, "!e/)"...,
					)
				case 12:
					data = /*line G7F16R.go:1*/ append(data, 52)
					i = 21
				case 6:
					data = /*line G7F16R.go:1*/ append(data, 104)
					i = 26
				case 16:
					data = /*line G7F16R.go:1*/ append(data, "J\x02\x02"...,
					)
					i = 17
				case 26:
					i = 14
					data = /*line G7F16R.go:1*/ append(data, "*%&!"...,
					)
				case 24:
					i = 2
					data = /*line G7F16R.go:1*/ append(data, "\x04.T"...,
					)
				case 29:
					data = /*line G7F16R.go:1*/ append(data, "{/4"...,
					)
					i = 11
				case 22:
					data = /*line G7F16R.go:1*/ append(data, "\x11\x12\r"...,
					)
					i = 1
				case 9:
					data = /*line G7F16R.go:1*/ append(data, 74)
					i = 19
				case 13:
					data = /*line G7F16R.go:1*/ append(data, "71"...,
					)
					i = 25
				case 2:
					data = /*line G7F16R.go:1*/ append(data, "]&"...,
					)
					i = 28
				case 19:
					data = /*line G7F16R.go:1*/ append(data, "D\x00\x14\x15"...,
					)
					i = 8
				case 21:
					data = /*line G7F16R.go:1*/ append(data, 42)
					i = 15
				case 1:
					data = /*line G7F16R.go:1*/ append(data, "\b\x16"...,
					)
					i = 9
				case 23:
					i = 24
					data = /*line G7F16R.go:1*/ append(data, "+T"...,
					)
				case 5:
					i = 18
					data = /*line G7F16R.go:1*/ append(data, 8)
				case 20:
					i = 12
					data = /*line G7F16R.go:1*/ append(data, "\n\"8"...,
					)
				case 14:
					i = 7
					data = /*line G7F16R.go:1*/ append(data, 36)
				case 11:
					data = /*line G7F16R.go:1*/ append(data, "91!5"...,
					)
					i = 13
				case 17:
					data = /*line G7F16R.go:1*/ append(data, 77)
					i = 0
				case 10:
					data = /*line G7F16R.go:1*/ append(data, "u\x12\x16L"...,
					)
					i = 30
				}
			}
			return /*line G7F16R.go:1*/ string(data)
		}(), /*line _BSpDC6.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.NaNYri773M)

		return
	}
	if !rDMAJ7T {
		/*line jnA488Eal.go:1*/ log.WPP4KjwN(func() /*line uUD7zXEF4v.go:1*/ string {
			fullData := [] /*line uUD7zXEF4v.go:1*/ byte("\x8c\x933\x89\xc0\x8a*\x12\xd90͜\x15\xc5\xf4\x9e\xb2\xf7s\xf6꿸\x9d\xa9\x96\xa9\n\x17?\x0f\xd4\xf8\f\x19\x92\x7f/\xa5\xc9(\xe4d\xe1l\x01\x01\x7f\x00\xbc\x92v\x9dI-\x96=$\f\xe8\x12\xb9\xbc[b\xc6M';^\xe3\x8cS(M\xd7\xd9$\x14{\xf69[\xf7\xd1\xe2FI`a\xc7vn{\x12\xcft\x16\xe8\xdaY\xceHAN\xff\x9d)\x9ew\x97S\x8b^\xcd\xd5\f\xdaRsͷ\xa1\uf4ecg\xa1V\xb9\xcf\x14\x1d/")
			idxKey := [] /*line uUD7zXEF4v.go:1*/ byte("\xa6\x8b\xa0\xea\xb0\xdf\f74")
			data := /*line uUD7zXEF4v.go:1*/ make([]byte, 0, 68)
			data = /*line uUD7zXEF4v.go:1*/ append(data, fullData[104^ /*line uUD7zXEF4v.go:1*/ int(idxKey[7])]+fullData[112^ /*line uUD7zXEF4v.go:1*/ int(idxKey[7])], fullData[69^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]+fullData[110^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[47^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]+fullData[24^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[119^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])]-fullData[47^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])], fullData[197^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]+fullData[149^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[12^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]^fullData[56^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[148^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]^fullData[216^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[29^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]^fullData[45^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[109^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])]+fullData[104^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])], fullData[1^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])]+fullData[96^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])], fullData[203^ /*line uUD7zXEF4v.go:1*/ int(idxKey[5])]+fullData[144^ /*line uUD7zXEF4v.go:1*/ int(idxKey[5])], fullData[181^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]-fullData[190^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[93^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]-fullData[52^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[62^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]-fullData[49^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[156^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]+fullData[110^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[80^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])]-fullData[34^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])], fullData[130^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]-fullData[154^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[100^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]+fullData[38^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[212^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]^fullData[39^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[205^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]+fullData[221^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[134^ /*line uUD7zXEF4v.go:1*/ int(idxKey[5])]+fullData[202^ /*line uUD7zXEF4v.go:1*/ int(idxKey[5])], fullData[120^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]-fullData[74^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[198^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]^fullData[150^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[230^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]^fullData[8^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[164^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]-fullData[195^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[133^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]-fullData[225^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[184^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]-fullData[187^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[133^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]+fullData[237^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[137^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]-fullData[177^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[67^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]^fullData[124^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[128^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]-fullData[206^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[145^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]-fullData[136^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[184^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]-fullData[219^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[184^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])]+fullData[135^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])], fullData[155^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]^fullData[185^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[207^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]^fullData[11^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[183^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]^fullData[174^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[155^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]-fullData[192^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[25^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]-fullData[31^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[178^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]^fullData[246^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[216^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])]^fullData[243^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])], fullData[236^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]^fullData[175^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[220^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]+fullData[211^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[105^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]-fullData[11^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[168^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]^fullData[211^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[170^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]+fullData[203^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[238^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])]-fullData[187^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])], fullData[46^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])]-fullData[26^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])], fullData[134^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])]^fullData[135^ /*line uUD7zXEF4v.go:1*/ int(idxKey[2])], fullData[59^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]^fullData[37^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[253^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])]-fullData[180^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])], fullData[127^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]^fullData[103^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[214^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]+fullData[196^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[101^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]+fullData[61^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[14^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]+fullData[102^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[141^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]^fullData[111^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[99^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]^fullData[28^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[128^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]-fullData[129^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[222^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])]+fullData[138^ /*line uUD7zXEF4v.go:1*/ int(idxKey[1])], fullData[228^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])]+fullData[202^ /*line uUD7zXEF4v.go:1*/ int(idxKey[4])], fullData[41^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])]-fullData[65^ /*line uUD7zXEF4v.go:1*/ int(idxKey[8])], fullData[192^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]+fullData[219^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[62^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])]+fullData[117^ /*line uUD7zXEF4v.go:1*/ int(idxKey[6])], fullData[204^ /*line uUD7zXEF4v.go:1*/ int(idxKey[5])]-fullData[207^ /*line uUD7zXEF4v.go:1*/ int(idxKey[5])], fullData[164^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])]+fullData[158^ /*line uUD7zXEF4v.go:1*/ int(idxKey[3])], fullData[236^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])]-fullData[239^ /*line uUD7zXEF4v.go:1*/ int(idxKey[0])], fullData[47^ /*line uUD7zXEF4v.go:1*/ int(idxKey[7])]+fullData[181^ /*line uUD7zXEF4v.go:1*/ int(idxKey[7])])
			return /*line uUD7zXEF4v.go:1*/ string(data)
		}(), /*line wCl82ad1.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.NaNYri773M)

		return
	}

	qNDcZn, dWmn_Gtk1d := /*line EuJYEU.go:1*/ cu1RSg2YR.aK7JnwJ.Add(gFIlzELaH)
	if !qNDcZn {
		/*line URjrdSJ.go:1*/ log.ViJSfja(func() /*line IPmQj9.go:1*/ string {
			data := [] /*line IPmQj9.go:1*/ byte("[%;g\xb3\xca\xf6\x11\xabԹs\x7f\xb0o\xf5!\x8f\xbc\xf6\xa4commi\xe9\xfb\xe5ar\x8e\xb1no\xa8 \xbc\x15ffici\xbe\xc21\xc9\xfd\xab[bu\x9f[^\x9d\xec\x1c\x8b\x19\xa0 ʽ\xf7%x")
			positions := [...]byte{38, 56, 38, 65, 5, 17, 26, 56, 49, 8, 2, 64, 60, 3, 26, 44, 19, 59, 28, 50, 45, 18, 15, 44, 26, 6, 56, 7, 38, 53, 46, 61, 55, 20, 58, 5, 17, 37, 12, 63, 3, 60, 27, 16, 56, 27, 48, 26, 3, 35, 56, 54, 19, 19, 61, 32, 47, 12, 10, 26, 16, 64, 13, 31, 4, 15, 5, 63, 32, 37, 55, 50, 57, 9, 5, 9, 59, 44}
			for i := 0; i < 78; i += 2 {
				localKey := /*line IPmQj9.go:1*/ byte(i) + /*line IPmQj9.go:1*/ byte(positions[i]^positions[i+1]) + 251
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line IPmQj9.go:1*/ string(data)
		}(), /*line EmQaZN7aw.go:1*/ cu1RSg2YR.ID(), gFIlzELaH.NaNYri773M)

		return
	}

	dWmn_Gtk1d.AVzGLiX4RWH = /*line Uw2SO2ovuaDY.go:1*/ cu1RSg2YR.ID()

	/*line PjFSrXu.go:1*/
	cu1RSg2YR.emCoevqAq(dWmn_Gtk1d)

	/*line HpKa_a7G4.go:1*/
	log.ViJSfja(func() /*line Yt4rnkR.go:1*/ string {
		data := /*line Yt4rnkR.go:1*/ make([]byte, 0, 90)
		i := 6
		decryptKey := 15
		for counter := 0; i != 23; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				data = /*line Yt4rnkR.go:1*/ append(data, "}x"...,
				)
				i = 9
			case 25:
				i = 26
				data = /*line Yt4rnkR.go:1*/ append(data, "HELB"...,
				)
			case 21:
				i = 7
				data = /*line Yt4rnkR.go:1*/ append(data, 58)
			case 13:
				i = 1
				data = /*line Yt4rnkR.go:1*/ append(data, "ae7"...,
				)
			case 15:
				i = 10
				data = /*line Yt4rnkR.go:1*/ append(data, "RQV"...,
				)
			case 22:
				data = /*line Yt4rnkR.go:1*/ append(data, "\":7&"...,
				)
				i = 19
			case 17:
				data = /*line Yt4rnkR.go:1*/ append(data, "0gc3"...,
				)
				i = 0
			case 6:
				data = /*line Yt4rnkR.go:1*/ append(data, "L3cI"...,
				)
				i = 20
			case 1:
				for y := range data {
					data[y] = data[y] ^ /*line Yt4rnkR.go:1*/ byte(decryptKey^y)
				}
				i = 23
			case 27:
				i = 32
				data = /*line Yt4rnkR.go:1*/ append(data, "#dhn"...,
				)
			case 8:
				data = /*line Yt4rnkR.go:1*/ append(data, "9+5:"...,
				)
				i = 17
			case 29:
				i = 18
				data = /*line Yt4rnkR.go:1*/ append(data, "PG@["...,
				)
			case 12:
				data = /*line Yt4rnkR.go:1*/ append(data, "CDf"...,
				)
				i = 25
			case 16:
				data = /*line Yt4rnkR.go:1*/ append(data, 117)
				i = 22
			case 18:
				data = /*line Yt4rnkR.go:1*/ append(data, "_W\x1f]"...,
				)
				i = 15
			case 19:
				data = /*line Yt4rnkR.go:1*/ append(data, "pz(}"...,
				)
				i = 8
			case 4:
				i = 3
				data = /*line Yt4rnkR.go:1*/ append(data, 87)
			case 2:
				data = /*line Yt4rnkR.go:1*/ append(data, "yzXU"...,
				)
				i = 29
			case 20:
				i = 21
				data = /*line Yt4rnkR.go:1*/ append(data, 51)
			case 0:
				i = 13
				data = /*line Yt4rnkR.go:1*/ append(data, "d\n\x06"...,
				)
			case 3:
				data = /*line Yt4rnkR.go:1*/ append(data, "\x04AN"...,
				)
				i = 28
			case 14:
				i = 5
				data = /*line Yt4rnkR.go:1*/ append(data, "bp"...,
				)
			case 31:
				data = /*line Yt4rnkR.go:1*/ append(data, 117)
				i = 11
			case 24:
				data = /*line Yt4rnkR.go:1*/ append(data, "M\x18AI"...,
				)
				i = 4
			case 32:
				i = 30
				data = /*line Yt4rnkR.go:1*/ append(data, "f}ei"...,
				)
			case 28:
				i = 12
				data = /*line Yt4rnkR.go:1*/ append(data, 78)
			case 11:
				data = /*line Yt4rnkR.go:1*/ append(data, "joq-"...,
				)
				i = 27
			case 9:
				i = 31
				data = /*line Yt4rnkR.go:1*/ append(data, "ohYv"...,
				)
			case 10:
				i = 24
				data = /*line Yt4rnkR.go:1*/ append(data, 83)
			case 7:
				i = 14
				data = /*line Yt4rnkR.go:1*/ append(data, 65)
			case 30:
				data = /*line Yt4rnkR.go:1*/ append(data, "o*"...,
				)
				i = 2
			case 26:
				data = /*line Yt4rnkR.go:1*/ append(data, "]\br "...,
				)
				i = 16
			}
		}
		return /*line Yt4rnkR.go:1*/ string(data)
	}(), /*line nML6TLCgH.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessAccept(bZwrdESj *blockchain.Qi_7sYpWjR8) {
	if /*line kITOALPJo.go:1*/ cu1RSg2YR.IsCommittee( /*line q5D3OCd2iD.go:1*/ cu1RSg2YR.ID(), bZwrdESj.XjPd77f0.BlockHeader.Epoch) {
		return
	}

	ouJpb2aF3mlv := bZwrdESj.XjPd77f0

	hYhbNJBbO := /*line ACaobqnIzVMT.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch()
	r0VHvbK := /*line GgF0rPz.go:1*/ cu1RSg2YR.GetHighBlockHeight()

	if ouJpb2aF3mlv.BlockHeader.Epoch > hYhbNJBbO {
		/*line Q7x_Z5vH9n.go:1*/ utils.OUWLEOgmMfR(cu1RSg2YR.sa21VPHg, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, bZwrdESj)

		return
	}

	k1qj7vimKIy7 := /*line g_VgGrOMrl.go:1*/ cu1RSg2YR.zV6vwpsi9W1(ouJpb2aF3mlv)
	if k1qj7vimKIy7 != nil {
		/*line b_dOgqhoP.go:1*/ log.Jp980o6YjM(func() /*line JTWtZDocPatS.go:1*/ string {
			data := [] /*line JTWtZDocPatS.go:1*/ byte("\n[\xf6o \f\x8a\xeeoc\xda=]A먱p\x9e!\x06x۳\x11\xcae\xd4ce\xa8N\xeb\xf9i\xf4a\xdaөbr\xback:\xdc%4")
			positions := [...]byte{39, 21, 1, 42, 20, 15, 31, 23, 32, 21, 23, 0, 12, 32, 38, 39, 24, 12, 5, 41, 19, 18, 3, 46, 26, 2, 22, 32, 33, 2, 25, 38, 14, 37, 31, 11, 5, 16, 31, 20, 27, 41, 46, 48, 10, 7, 25, 42, 6, 46, 35, 33, 0, 7, 30, 19, 35, 0}
			for i := 0; i < 58; i += 2 {
				localKey := /*line JTWtZDocPatS.go:1*/ byte(i) + /*line JTWtZDocPatS.go:1*/ byte(positions[i]^positions[i+1]) + 62
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line JTWtZDocPatS.go:1*/ string(data)
		}(), /*line uTtX4Z.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)

		return
	}

	if ouJpb2aF3mlv.BlockHeader.BlockHeight > r0VHvbK+1 {

		/*line QzOOx4U.go:1*/
		utils.OUWLEOgmMfR(cu1RSg2YR.sa21VPHg, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight-1, bZwrdESj)

		/*line lJlqaRzBqV.go:1*/
		log.ViJSfja(func() /*line Jo0yI2jMOfO.go:1*/ string {
			fullData := [] /*line Jo0yI2jMOfO.go:1*/ byte("\x93\xfb\xa8\xeaӴX\xb9\xb3\xa6\x8fx\xb8\xeb\uedf5\xf8\x14X\xf5\f\x91\x8d㪄\xb9\xb3\x1e6ޙd\x7f\xc1\x12\x0e<\x7f\xf5\xe6Tb\x90~\xf9\xc9v\xf0ʐ\xe4Y:D\x8e\x9f7\xe8@\xe9A{\xb4|)\xd4'\xcc\xdd3T\xd6{v\xec\xe7\x1e\xecs\xb3\x82\xd73(\x9e\x96o\xb9\a\x9b-\xeaI\xce\xd0\x1f\x8fV4\x17\xb1`\f\x89\xb7\xeb\x1aU\xf5\xbaGS\faJ8\xa9\xc3\t\xf3\xa7\xf1\x820\xae\x80v\x88\x14\xa4\x96W\xa4\xe0\xbb\x1a\xd4]\xfc\x86")
			idxKey := [] /*line Jo0yI2jMOfO.go:1*/ byte("\xf86\x9dHN][m\x8dA\xc3\xe8ve\xfd(")
			data := /*line Jo0yI2jMOfO.go:1*/ make([]byte, 0, 72)
			data = /*line Jo0yI2jMOfO.go:1*/ append(data, fullData[243^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]-fullData[252^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])], fullData[148^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[11])]-fullData[99^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[11])], fullData[121^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])]-fullData[35^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])], fullData[49^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]-fullData[44^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[74^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])]+fullData[241^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])], fullData[205^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])]-fullData[2^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])], fullData[115^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[3])]^fullData[68^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[3])], fullData[33^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]^fullData[162^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[40^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]-fullData[110^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[171^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]+fullData[201^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])], fullData[41^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]-fullData[172^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[220^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[11])]+fullData[226^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[11])], fullData[11^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])]+fullData[80^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])], fullData[62^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])]-fullData[58^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])], fullData[253^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[2])]^fullData[129^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[2])], fullData[100^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[3])]^fullData[49^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[3])], fullData[95^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])]+fullData[81^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])], fullData[166^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]+fullData[168^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])], fullData[218^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])]+fullData[23^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])], fullData[28^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])]-fullData[123^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])], fullData[109^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])]^fullData[62^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])], fullData[232^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])]-fullData[198^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])], fullData[117^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])]-fullData[74^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])], fullData[69^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])]^fullData[98^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])], fullData[40^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])]^fullData[55^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])], fullData[228^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])]+fullData[242^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])], fullData[61^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]^fullData[13^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[140^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])]^fullData[186^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])], fullData[27^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])]-fullData[57^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])], fullData[210^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])]^fullData[164^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])], fullData[39^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])]-fullData[26^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])], fullData[250^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])]-fullData[221^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])], fullData[55^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]^fullData[66^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[217^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])]^fullData[142^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])], fullData[205^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[10])]^fullData[156^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[10])], fullData[10^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]+fullData[103^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[250^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[10])]^fullData[192^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[10])], fullData[44^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])]^fullData[228^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])], fullData[85^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])]+fullData[12^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])], fullData[14^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]+fullData[174^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])], fullData[107^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])]+fullData[15^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])], fullData[83^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])]-fullData[85^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])], fullData[178^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]-fullData[232^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])], fullData[15^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]+fullData[152^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])], fullData[173^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]-fullData[83^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[115^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])]+fullData[221^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[5])], fullData[4^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])]^fullData[91^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])], fullData[36^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])]-fullData[60^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])], fullData[142^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[11])]+fullData[96^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[11])], fullData[6^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])]+fullData[26^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[12])], fullData[90^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])]^fullData[115^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])], fullData[203^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[2])]-fullData[167^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[2])], fullData[59^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])]+fullData[18^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[4])], fullData[117^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]^fullData[26^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[245^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[10])]^fullData[208^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[10])], fullData[48^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]^fullData[74^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[224^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])]+fullData[80^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[7])], fullData[162^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])]^fullData[217^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])], fullData[248^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])]-fullData[163^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])], fullData[99^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]-fullData[84^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[144^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[2])]^fullData[165^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[2])], fullData[9^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])]-fullData[42^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[9])], fullData[62^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]-fullData[1^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[120^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]+fullData[70^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[248^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])]-fullData[153^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[0])], fullData[30^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])]^fullData[20^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[6])], fullData[99^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])]+fullData[44^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[15])], fullData[36^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[1])]+fullData[75^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[1])], fullData[227^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])]-fullData[127^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[13])], fullData[189^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])]^fullData[235^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[14])], fullData[224^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])]-fullData[203^ /*line Jo0yI2jMOfO.go:1*/ int(idxKey[8])])
			return /*line Jo0yI2jMOfO.go:1*/ string(data)
		}(), /*line ykNW9grCcyx.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash)
		return
	}

	if ouJpb2aF3mlv.Proposer != /*line aNeDmyFOTua.go:1*/ cu1RSg2YR.ID() {

		if ouJpb2aF3mlv.BlockHeader.StateRoot != /*line SwbZXYwY5gA.go:1*/ cu1RSg2YR.iSgBUFf9.GetCurrentStateHash() {

			/*line jmF1aQzKzn9.go:1*/
			log.WPP4KjwN(func() /*line fUSQJNAZ3B.go:1*/ string {
				seed := /*line fUSQJNAZ3B.go:1*/ byte(179)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line fUSQJNAZ3B.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line fUSQJNAZ3B.go:1*/ fnc(14)(230)(29)(33)(5)(18)(76)(186)(113)(214)(174)(106)(212)(118)(14)(28)(58)(127)(2)(185)(105)(37)(75)(131)(25)(35)(41)(107)(232)(197)(66)(210)(165)(79)(74)(234)(191)(137)(15)(25)(238)(5)(5)(230)(209)(245)(146)(102)(214)(175)(82)(172)(53)(135)(18)(34)(69)(150)(216)(181)(187)
				return /*line fUSQJNAZ3B.go:1*/ string(data)
			}(), /*line dsLbO3TY3xH8.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHash, ouJpb2aF3mlv.BlockHeader.BlockHeight)
			return
		}
	}

	/*line VUayuAa2q3q.go:1*/
	cu1RSg2YR.iSgBUFf9.AddCoordinationBlock(ouJpb2aF3mlv)

	_ = /*line BPZnTc.go:1*/ cu1RSg2YR.wMIZeBk(ouJpb2aF3mlv.CQC)
	_ = /*line H4RK165niGz.go:1*/ cu1RSg2YR.xAnHtUT9ATI(ouJpb2aF3mlv.BlockHeader.BlockHeight)
	/*line Zk8n4XgB.go:1*/ cu1RSg2YR.aDTslULytPz_(ouJpb2aF3mlv.CQC)

	if /*line dCDmwmAvq.go:1*/ cu1RSg2YR.mGpWFGC2.IsTimeToElect() {

		hYhbNJBbO := /*line dYwN_wbQZc51.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch()
		ey93G03TMe := /*line RTmwMoyzdp.go:1*/ cu1RSg2YR.ElectCommittees(ouJpb2aF3mlv.BlockHash, hYhbNJBbO+1)
		if ey93G03TMe != nil {
			/*line F2cOXB.go:1*/ log.ViJSfja(func() /*line H8fmWmLyFOW.go:1*/ string {
				data := [] /*line H8fmWmLyFOW.go:1*/ byte("A\xb7v]\xcc( \xd7oceTs\xb6Ɯe0j)\xbde<h\x8dP|\xad*DC\x9c\x92mg\xdeutej,OBv\x7ffi\x19,\r:w epocs\x9e\x7f\x82")
				positions := [...]byte{29, 34, 6, 60, 46, 29, 58, 32, 23, 13, 36, 28, 42, 17, 4, 20, 15, 27, 29, 57, 26, 48, 32, 15, 13, 24, 36, 18, 13, 1, 7, 1, 14, 31, 41, 6, 23, 32, 35, 41, 48, 22, 31, 59, 39, 50, 58, 35, 30, 49, 37, 37, 25, 34, 58, 60, 4, 0, 44, 40, 11, 17, 18, 39, 4, 47}
				for i := 0; i < 66; i += 2 {
					localKey := /*line H8fmWmLyFOW.go:1*/ byte(i) + /*line H8fmWmLyFOW.go:1*/ byte(positions[i]^positions[i+1]) + 206
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line H8fmWmLyFOW.go:1*/ string(data)
			}(), /*line WvaD4Uc_S.go:1*/ cu1RSg2YR.ID(), ey93G03TMe, hYhbNJBbO+1)
		}
	}

	/*line KmZhIcX.go:1*/
	cu1RSg2YR.mGpWFGC2.AdvanceView(ouJpb2aF3mlv.BlockHeader.View)

	/*line ayqLPxiOP.go:1*/
	cu1RSg2YR.z6VaPTrjpK(ouJpb2aF3mlv.CQC)

	k1qj7vimKIy7 = /*line sj3JraqH7khz.go:1*/ cu1RSg2YR.fhcoYWd(ouJpb2aF3mlv.CQC)

	if k1qj7vimKIy7 != nil {
		/*line BBBE9jorgPxS.go:1*/ log.Jp980o6YjM(func() /*line H9VvlcyyA.go:1*/ string {
			key := [] /*line H9VvlcyyA.go:1*/ byte("S\"I[\n{\x80h^\xc57N\x8d\xf9\xbdR\f'\xcf\xca\x14\xe2\xadH\xaf\xa1\xf0\xdbO\xfe\x7f\x18\xa5\xf1\xf1\x13*K\xe0<1q\xe9\xa3")
			data := [] /*line H9VvlcyyA.go:1*/ byte("\b\x03-\x02\x16\xad\xd0\n\x11\x9e.%\xe6H\xa6\x11YI\xa5_\f\x84\xb4!\xbd\xc4tE\x14q\xeeUă//B$\x83/\t\xaf<\xd3")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line H9VvlcyyA.go:1*/ string(data)
		}(), /*line obE_23.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)
		return
	}
	hYhbNJBbO, uxNBF8Z7 := /*line C_SVw7o.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpochView()

	if bZwrdESj, i3KmQgjb := cu1RSg2YR.sa21VPHg[hYhbNJBbO][uxNBF8Z7][ouJpb2aF3mlv.BlockHeader.BlockHeight]; i3KmQgjb {
		/*line gtBpZa0Sy.go:1*/ cu1RSg2YR.ProcessAccept(bZwrdESj)
		/*line CnAkB2b.go:1*/ delete(cu1RSg2YR.sa21VPHg[hYhbNJBbO][uxNBF8Z7], ouJpb2aF3mlv.BlockHeader.BlockHeight)
	}

	if ! /*line A11dw40G.go:1*/ cu1RSg2YR.IsCommittee( /*line YJzntCBekuW.go:1*/ cu1RSg2YR.ID(), hYhbNJBbO) {
		return
	}

	if ouJpb2aF3mlv, kbgkwuEjk := cu1RSg2YR.wdUdVrSEU[hYhbNJBbO][uxNBF8Z7][ouJpb2aF3mlv.BlockHeader.BlockHeight]; kbgkwuEjk {
		_ = /*line xTfE7N.go:1*/ cu1RSg2YR.ProcessCoordinationBlock(ouJpb2aF3mlv)
		/*line DgPm_iLj8AP.go:1*/ delete(cu1RSg2YR.wdUdVrSEU[hYhbNJBbO][uxNBF8Z7], ouJpb2aF3mlv.BlockHeader.BlockHeight)
	}
}

func (cu1RSg2YR *L4qxkgqr2rNN) zV6vwpsi9W1(ouJpb2aF3mlv *blockchain.CoordinationBlock) error {

	cWdSpmiBLOz := ouJpb2aF3mlv.QC
	if cWdSpmiBLOz.BlockHeight < /*line RBUg3V.go:1*/ cu1RSg2YR.GetHighBlockHeight() {
		return /*line gGagjOn.go:1*/ fmt.Errorf(func() /*line VfjEOyL.go:1*/ string {
			data := /*line VfjEOyL.go:1*/ make([]byte, 0, 32)
			i := 8
			decryptKey := 112
			for counter := 0; i != 15; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 5:
					data = /*line VfjEOyL.go:1*/ append(data, "<559"...,
					)
					i = 11
				case 6:
					i = 14
					data = /*line VfjEOyL.go:1*/ append(data, "b7!+"...,
					)
				case 4:
					i = 10
					data = /*line VfjEOyL.go:1*/ append(data, "!<"...,
					)
				case 3:
					i = 6
					data = /*line VfjEOyL.go:1*/ append(data, "#7"...,
					)
				case 10:
					i = 3
					data = /*line VfjEOyL.go:1*/ append(data, "n#"...,
					)
				case 11:
					data = /*line VfjEOyL.go:1*/ append(data, ":59"...,
					)
					i = 9
				case 8:
					data = /*line VfjEOyL.go:1*/ append(data, ":)"...,
					)
					i = 0
				case 9:
					i = 2
					data = /*line VfjEOyL.go:1*/ append(data, 36)
				case 12:
					data = /*line VfjEOyL.go:1*/ append(data, "{85"...,
					)
					i = 13
				case 14:
					data = /*line VfjEOyL.go:1*/ append(data, "/!h"...,
					)
					i = 12
				case 2:
					data = /*line VfjEOyL.go:1*/ append(data, "ws"...,
					)
					i = 1
				case 0:
					data = /*line VfjEOyL.go:1*/ append(data, 105)
					i = 4
				case 13:
					data = /*line VfjEOyL.go:1*/ append(data, 55)
					i = 5
				case 1:
					i = 7
					data = /*line VfjEOyL.go:1*/ append(data, 35)
				case 7:
					i = 15
					for y := range data {
						data[y] = data[y] ^ /*line VfjEOyL.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return /*line VfjEOyL.go:1*/ string(data)
		}(), cWdSpmiBLOz.BlockHeight)
	}

	if ouJpb2aF3mlv.QC.AVzGLiX4RWH != /*line UR5KrNZGa.go:1*/ cu1RSg2YR.ID() {
		lCSMkWR3Qe, _ := /*line Of_PDrV.go:1*/ crypto.BFo6_c(cWdSpmiBLOz.Pp__49cd, cWdSpmiBLOz.ZWsU_2, cWdSpmiBLOz.X5i3Ynndjb)
		if !lCSMkWR3Qe {
			return /*line mR6lAB.go:1*/ errors.New(func() /*line An7JQ36Ya.go:1*/ string {
				data := [] /*line An7JQ36Ya.go:1*/ byte("\xf2ece\xb3vFd \xa3\xaa\xb1\xb1arum9eE\xe1JHi}v\xaf\x90Rd \xbdrggrt\\`Hs")
				positions := [...]byte{22, 20, 13, 37, 19, 6, 10, 28, 19, 39, 14, 27, 20, 37, 11, 27, 12, 0, 9, 14, 13, 4, 19, 21, 26, 31, 17, 9, 12, 32, 32, 17, 19, 6, 24, 35, 38, 18, 10, 9, 0, 28, 34, 14}
				for i := 0; i < 44; i += 2 {
					localKey := /*line An7JQ36Ya.go:1*/ byte(i) + /*line An7JQ36Ya.go:1*/ byte(positions[i]^positions[i+1]) + 191
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line An7JQ36Ya.go:1*/ string(data)
			}())
		}
	}

	hYhbNJBbO, uxNBF8Z7 := /*line X_C1T0.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpochView()
	r0VHvbK := /*line y_gFYa9MjjC.go:1*/ cu1RSg2YR.GetHighBlockHeight()

	if ! /*line SYrdKD.go:1*/ cu1RSg2YR.Static.IsLeader(ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch) {
		return /*line AXtwcf7I.go:1*/ fmt.Errorf(func() /*line AU5l8ztBXl.go:1*/ string {
			data := [] /*line AU5l8ztBXl.go:1*/ byte("\xc0\x93K\x92i3\xa6\xa6 a\xfe\xb8\xd3'p\x18sa\xff View\r%v\xd5I\xaa\x7f\xc6\xe59Rv\xef \x95\xe5c a\xb9 :\x99C\xc2\xf2\x1d\x87,lea\x14\x00V\xfd$\v")
			positions := [...]byte{27, 43, 5, 18, 5, 2, 7, 57, 20, 33, 0, 34, 47, 48, 40, 15, 3, 59, 37, 50, 11, 57, 13, 20, 46, 60, 15, 33, 43, 58, 6, 45, 45, 56, 47, 57, 28, 32, 36, 30, 6, 60, 24, 43, 61, 38, 48, 36, 29, 59, 12, 51, 15, 12, 27, 31, 1, 52, 34, 32, 1, 10, 36, 0, 13, 39, 49, 43, 61, 49}
			for i := 0; i < 70; i += 2 {
				localKey := /*line AU5l8ztBXl.go:1*/ byte(i) + /*line AU5l8ztBXl.go:1*/ byte(positions[i]^positions[i+1]) + 32
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line AU5l8ztBXl.go:1*/ string(data)
		}(), ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.Proposer)
	}

	if ouJpb2aF3mlv.BlockHeader.Epoch < hYhbNJBbO || ouJpb2aF3mlv.BlockHeader.View < uxNBF8Z7 || ouJpb2aF3mlv.BlockHeader.BlockHeight < r0VHvbK {
		return /*line MC5k0J_.go:1*/ fmt.Errorf(func() /*line kPvLnOfMZe.go:1*/ string {
			seed := /*line kPvLnOfMZe.go:1*/ byte(221)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line kPvLnOfMZe.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line kPvLnOfMZe.go:1*/ fnc(149)(243)(254)(2)(4)(13)(239)(255)(188)(65)(191)(83)(1)(237)(11)(249)(187)(80)(2)(253)(1)(255)(4)(238)(11)(180)(70)(12)(253)(254)(179)(5)(81)(170)(86)(243)(252)(18)(169)(5)(81)(170)(66)(10)(3)(244)(8)(253)(253)(4)(254)(1)(12)(172)(5)(81)
			return /*line kPvLnOfMZe.go:1*/ string(data)
		}(), ouJpb2aF3mlv.Proposer, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.BlockHeight)
	}

	if ouJpb2aF3mlv.Proposer != /*line Pb4uiu5.go:1*/ cu1RSg2YR.ID() {
		iG6LeZAvtH1B, _ := /*line HySRCLcIN.go:1*/ crypto.SkrCuscT(ouJpb2aF3mlv.CommitteeSignature[1] /*line l7SZnU.go:1*/, crypto.IDToByte(ouJpb2aF3mlv.BlockHash))
		if !iG6LeZAvtH1B {
			return /*line stWyQZS3.go:1*/ errors.New(func() /*line EcavWBO0.go:1*/ string {
				key := [] /*line EcavWBO0.go:1*/ byte("\xd3k\tW1\xd6\xffnm\xb5O\xf8\xa9\xbb[\xdf|\\\x8b\xbe}v\x983\x92\x8bz\xcd\"n\x13\x1a\x12\x85\x88\xe4|\x0eXy/\x9d")
				data := [] /*line EcavWBO0.go:1*/ byte("\xa1\x0ej2X\xa0\x9a\nM\xd4o\x9a\xc5\xd48\xb4\\+\xe2\xca\x15V\xf9]\xb2\xe2\x14\xbbC\x02z~2\xf6\xe1\x83\x12o,\f]\xf8")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line EcavWBO0.go:1*/ string(data)
			}())
		}
	}

	if ouJpb2aF3mlv.QC == nil {
		return /*line Y7kbr5.go:1*/ errors.New(func() /*line MdNIsBl.go:1*/ string {
			fullData := [] /*line MdNIsBl.go:1*/ byte("\xd9\xff\\ś\xb7\xc9\xf3|\x1a\xffS\xef\xf0\x9c\xdd\\\xfdT\x89`\x833\xe6\x94x{\xe4\x13[w`[;\xd3,\x9dL\xf2\u05f7\x01\xe0\xf3\xe3؛\xe1\x1f\x04\xc1:\bw:\x8e\x8e\x99")
			idxKey := [] /*line MdNIsBl.go:1*/ byte("j\xb60 \x86\x17\xf9\x15\x1c<\x01\xca\x16x\f.")
			data := /*line MdNIsBl.go:1*/ make([]byte, 0, 30)
			data = /*line MdNIsBl.go:1*/ append(data, fullData[116^ /*line MdNIsBl.go:1*/ int(idxKey[13])]^fullData[124^ /*line MdNIsBl.go:1*/ int(idxKey[13])], fullData[60^ /*line MdNIsBl.go:1*/ int(idxKey[14])]^fullData[18^ /*line MdNIsBl.go:1*/ int(idxKey[14])], fullData[27^ /*line MdNIsBl.go:1*/ int(idxKey[2])]-fullData[8^ /*line MdNIsBl.go:1*/ int(idxKey[2])], fullData[101^ /*line MdNIsBl.go:1*/ int(idxKey[13])]-fullData[89^ /*line MdNIsBl.go:1*/ int(idxKey[13])], fullData[29^ /*line MdNIsBl.go:1*/ int(idxKey[8])]^fullData[56^ /*line MdNIsBl.go:1*/ int(idxKey[8])], fullData[230^ /*line MdNIsBl.go:1*/ int(idxKey[11])]-fullData[255^ /*line MdNIsBl.go:1*/ int(idxKey[11])], fullData[144^ /*line MdNIsBl.go:1*/ int(idxKey[1])]-fullData[163^ /*line MdNIsBl.go:1*/ int(idxKey[1])], fullData[31^ /*line MdNIsBl.go:1*/ int(idxKey[12])]-fullData[19^ /*line MdNIsBl.go:1*/ int(idxKey[12])], fullData[185^ /*line MdNIsBl.go:1*/ int(idxKey[1])]+fullData[129^ /*line MdNIsBl.go:1*/ int(idxKey[1])], fullData[147^ /*line MdNIsBl.go:1*/ int(idxKey[1])]-fullData[149^ /*line MdNIsBl.go:1*/ int(idxKey[1])], fullData[43^ /*line MdNIsBl.go:1*/ int(idxKey[14])]+fullData[2^ /*line MdNIsBl.go:1*/ int(idxKey[14])], fullData[227^ /*line MdNIsBl.go:1*/ int(idxKey[6])]^fullData[229^ /*line MdNIsBl.go:1*/ int(idxKey[6])], fullData[20^ /*line MdNIsBl.go:1*/ int(idxKey[12])]^fullData[0^ /*line MdNIsBl.go:1*/ int(idxKey[12])], fullData[16^ /*line MdNIsBl.go:1*/ int(idxKey[10])]+fullData[24^ /*line MdNIsBl.go:1*/ int(idxKey[10])], fullData[191^ /*line MdNIsBl.go:1*/ int(idxKey[4])]+fullData[164^ /*line MdNIsBl.go:1*/ int(idxKey[4])], fullData[40^ /*line MdNIsBl.go:1*/ int(idxKey[15])]+fullData[0^ /*line MdNIsBl.go:1*/ int(idxKey[15])], fullData[183^ /*line MdNIsBl.go:1*/ int(idxKey[4])]-fullData[157^ /*line MdNIsBl.go:1*/ int(idxKey[4])], fullData[74^ /*line MdNIsBl.go:1*/ int(idxKey[0])]+fullData[94^ /*line MdNIsBl.go:1*/ int(idxKey[0])], fullData[59^ /*line MdNIsBl.go:1*/ int(idxKey[9])]+fullData[52^ /*line MdNIsBl.go:1*/ int(idxKey[9])], fullData[63^ /*line MdNIsBl.go:1*/ int(idxKey[5])]^fullData[23^ /*line MdNIsBl.go:1*/ int(idxKey[5])], fullData[60^ /*line MdNIsBl.go:1*/ int(idxKey[12])]^fullData[14^ /*line MdNIsBl.go:1*/ int(idxKey[12])], fullData[35^ /*line MdNIsBl.go:1*/ int(idxKey[9])]+fullData[21^ /*line MdNIsBl.go:1*/ int(idxKey[9])], fullData[2^ /*line MdNIsBl.go:1*/ int(idxKey[10])]-fullData[17^ /*line MdNIsBl.go:1*/ int(idxKey[10])], fullData[14^ /*line MdNIsBl.go:1*/ int(idxKey[8])]^fullData[47^ /*line MdNIsBl.go:1*/ int(idxKey[8])], fullData[36^ /*line MdNIsBl.go:1*/ int(idxKey[12])]^fullData[57^ /*line MdNIsBl.go:1*/ int(idxKey[12])], fullData[52^ /*line MdNIsBl.go:1*/ int(idxKey[3])]-fullData[42^ /*line MdNIsBl.go:1*/ int(idxKey[3])], fullData[27^ /*line MdNIsBl.go:1*/ int(idxKey[14])]+fullData[58^ /*line MdNIsBl.go:1*/ int(idxKey[14])], fullData[59^ /*line MdNIsBl.go:1*/ int(idxKey[12])]^fullData[5^ /*line MdNIsBl.go:1*/ int(idxKey[12])], fullData[103^ /*line MdNIsBl.go:1*/ int(idxKey[0])]+fullData[97^ /*line MdNIsBl.go:1*/ int(idxKey[0])])
			return /*line MdNIsBl.go:1*/ string(data)
		}())
	}

	if /*line d5_YSm.go:1*/ cu1RSg2YR.IsByz() && /*line jshwgUm.go:1*/ config.GetConfig().Strategy == FORK && /*line dCjrzS.go:1*/ cu1RSg2YR.IsLeader( /*line JWpS6A7BX6H.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.LwVL87+1 /*line l0lPxs9KA.go:1*/, cu1RSg2YR.mGpWFGC2.GetCurEpoch()) {
		/*line E2LEddd.go:1*/ cu1RSg2YR.mGpWFGC2.AdvanceView(cWdSpmiBLOz.LwVL87)
		return nil
	}

	/*line XopmN1O028.go:1*/
	log.ViJSfja(func() /*line IM1gpKbCnP.go:1*/ string {
		seed := /*line IM1gpKbCnP.go:1*/ byte(243)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line IM1gpKbCnP.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line IM1gpKbCnP.go:1*/ fnc(168)(190)(47)(213)(125)(242)(188)(250)(237)(12)(30)(234)(240)(48)(198)(27)(240)(29)(247)(225)(10)(18)(241)(19)(203)(56)(227)(12)(16)(162)(13)(92)(255)(251)(249)(250)(235)(11)(29)(182)(60)(250)(237)(12)(30)(234)(240)(26)(227)(23)(167)(77)(30)(235)(240)(29)(247)(225)(10)(18)(241)(19)(169)(80)(238)(31)(236)(16)(171)(84)(230)(31)(236)(16)(195)(43)(16)(238)(31)(226)(66)(154)(113)(179)(88)(166)(31)(240)(242)(77)(228)(141)(67)(88)(181)(245)(21)(236)(19)(180)(98)(129)(83)(88)(153)(45)(172)(98)(129)(93)
		return /*line IM1gpKbCnP.go:1*/ string(data)
	}(), /*line SXVwJo_k.go:1*/ cu1RSg2YR.ID(), ouJpb2aF3mlv.BlockHeader.BlockHeight, ouJpb2aF3mlv.BlockHeader.View, ouJpb2aF3mlv.BlockHeader.Epoch, ouJpb2aF3mlv.BlockHash)

	return nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) aVd2nRenj(cWdSpmiBLOz *quorum.HPOWa1PT0xzq) {
	/*line qGswgms.go:1*/ log.ViJSfja(func() /*line sbEo5u6KWx.go:1*/ string {
		key := [] /*line sbEo5u6KWx.go:1*/ byte("\xe9\x12\x10kuRh\xf9\xfdG2\x99\x01Q\xfdq\x0f\xe5\xfe@~\x95\xf1\xf8\x0fg\xab\x99{ͮ\xb120\x87&\xa1\xbaY\xd1F\x95\xe7_\b\x11vH\xc0\xb9\xa7\x9dS$t:\xb29\x8a3\x0ep\xf4\x03dҫ\xfc\xe0\xc5\x16\xc6\xfa\x9e\xdc@RA\xf4\xbe\xec\xaa@:\x82I\xd3\xf2N\xcb\xd7\xca^2\x16\xf9\x87ö")
		data := [] /*line sbEo5u6KWx.go:1*/ byte("\xb27f6Uz\x18\x8b\x92$W\xear\x12\x98\x03{\x8c\x98)\x1d\xf4\x85\x9dY\b\xdf\xfcR\xed\xde\xc3]S\xe2U\xd2\xd37\xb6f\xe4\x920zd\x1bh\xa3\xdc\xd5\xe9:B\x1dY\xd3M\xe3\\`P\x96o\v\xb1\xc0\xb4\x85\xacq\xae\x8e\xbe\xf96r7\x9dۛ\x8aeL\xa2,\xa3\x9d-\xa3\xf7\xef(\x12_\xbd\xa7\xe6\xce")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line sbEo5u6KWx.go:1*/ string(data)
	}(), /*line cIg0B0bx6Q.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.BlockHeight, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.ZWsU_2)

	if cWdSpmiBLOz.BlockHeight < /*line Xo0OttbWK5.go:1*/ cu1RSg2YR.GetHighBlockHeight() {
		return
	}

	ouJpb2aF3mlv, kbgkwuEjk := cu1RSg2YR.cA3KSEC5m[cWdSpmiBLOz.EmFrce][cWdSpmiBLOz.LwVL87][cWdSpmiBLOz.BlockHeight]
	if !kbgkwuEjk {
		return
	}

	/*line ZtJN0Am.go:1*/
	cu1RSg2YR.iSgBUFf9.AddCoordinationBlock(ouJpb2aF3mlv)
	/*line etC_BDs.go:1*/ delete(cu1RSg2YR.cA3KSEC5m[cWdSpmiBLOz.EmFrce][cWdSpmiBLOz.LwVL87], cWdSpmiBLOz.BlockHeight)
	/*line EfgyIk5.go:1*/ log.ViJSfja(func() /*line FKTt8y2oB6H.go:1*/ string {
		key := [] /*line FKTt8y2oB6H.go:1*/ byte("$\xe3&$\xa3\xa2\xf3\x19\xb5\xa5\x05\xd0A\x1d\x92e\xe5Z\xe5\v\xe2+]\xc270\x88\x1a\xb6\x00\x8c\x14\x93v\x96b»\n\x93\xadrQ;\xf5\x92=\x84c\xba\xfb\x03#\xa3\x86")
		data := [] /*line FKTt8y2oB6H.go:1*/ byte("\x7f\xc6Py\x83\x8a\x83k\xda\xc6`\xa32^\xf7\x17\x913\x83b\x81J)\xa7a_\xfc\x7f\x9f \xfca\xe7V\xf4\x0e\xad\xd8a\xb3\xc4\x1c%T\xd5\xf0Q\xeb\x00јkB\xca\xe8")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line FKTt8y2oB6H.go:1*/ string(data)
	}(), /*line _Fe6IuX.go:1*/ cu1RSg2YR.ID())

	if /*line Bxm6929GZ.go:1*/ cu1RSg2YR.IsByz() && /*line wXhmqY3QUTVP.go:1*/ config.GetConfig().Strategy == FORK && /*line HgEztQOWVvO.go:1*/ cu1RSg2YR.IsLeader( /*line ZJ9eyaC.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.LwVL87+1 /*line RJrQ4P.go:1*/, cu1RSg2YR.mGpWFGC2.GetCurEpoch()) {
		/*line CBoW3Av3D3lG.go:1*/ cu1RSg2YR.mGpWFGC2.AdvanceView(cWdSpmiBLOz.LwVL87)
		return
	}

	k1qj7vimKIy7 := /*line QHhGyjKtFT.go:1*/ cu1RSg2YR.wMIZeBk(cWdSpmiBLOz)
	if k1qj7vimKIy7 != nil {
		cu1RSg2YR.vHhOhURml[cWdSpmiBLOz.ZWsU_2] = cWdSpmiBLOz
		/*line rt6cuAZxi.go:1*/ log.ViJSfja(func() /*line U7d1McnK.go:1*/ string {
			fullData := [] /*line U7d1McnK.go:1*/ byte("\x8aV@\xa6\x04\xa5''\x80\xf3\x14\xac\xa0\x10z\xc0Ϥ+ۃ\xad\x0f\xcf>\xb2L\xeeD\xcaÂ\x18%g\x9d\x14;\xfd \xeb\xfd\x02\x18\xdb8\x02CVn\xb2T\xa9p\xbd\x8eҎ\x90d\x12t\xf4j\xdf\xc0\xcfD\xc4J&m\xe4\x94&3-\xd6\x0e$\xc4\x0eIa\xf6\xf16IK\xb8̋\xed\xa7\xebK\x12\x9f\xed&\"\x00\xb7\xadot\x1b\xe4]\xe7\f\x17V\xa6տ\xa0\x88\xb5\xae\x8e܆\x90,%\xae\x03\fd/\xae\x13\xa7\xfc\xbbH\xfb\x9cv\xf7\x88Ӆ3\xc2\bkY\x99\xe6\xc5S|?\xc5W\x88\xa6\x92\xc0\x01\x93\xe3\xbf\xd9\r\x06aT")
			idxKey := [] /*line U7d1McnK.go:1*/ byte("\xd7Я\xe4i\f\xdf\xfd\x0e\t\x04\x03\xaf\xd8\xf7y")
			data := /*line U7d1McnK.go:1*/ make([]byte, 0, 86)
			data = /*line U7d1McnK.go:1*/ append(data, fullData[79^ /*line U7d1McnK.go:1*/ int(idxKey[13])]-fullData[231^ /*line U7d1McnK.go:1*/ int(idxKey[13])], fullData[158^ /*line U7d1McnK.go:1*/ int(idxKey[5])]^fullData[64^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[124^ /*line U7d1McnK.go:1*/ int(idxKey[5])]^fullData[43^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[160^ /*line U7d1McnK.go:1*/ int(idxKey[3])]-fullData[198^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[116^ /*line U7d1McnK.go:1*/ int(idxKey[7])]-fullData[238^ /*line U7d1McnK.go:1*/ int(idxKey[7])], fullData[170^ /*line U7d1McnK.go:1*/ int(idxKey[13])]-fullData[205^ /*line U7d1McnK.go:1*/ int(idxKey[13])], fullData[175^ /*line U7d1McnK.go:1*/ int(idxKey[14])]^fullData[210^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[7^ /*line U7d1McnK.go:1*/ int(idxKey[11])]-fullData[156^ /*line U7d1McnK.go:1*/ int(idxKey[11])], fullData[20^ /*line U7d1McnK.go:1*/ int(idxKey[15])]^fullData[228^ /*line U7d1McnK.go:1*/ int(idxKey[15])], fullData[4^ /*line U7d1McnK.go:1*/ int(idxKey[10])]-fullData[3^ /*line U7d1McnK.go:1*/ int(idxKey[10])], fullData[182^ /*line U7d1McnK.go:1*/ int(idxKey[1])]+fullData[83^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[223^ /*line U7d1McnK.go:1*/ int(idxKey[14])]+fullData[130^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[170^ /*line U7d1McnK.go:1*/ int(idxKey[1])]+fullData[178^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[129^ /*line U7d1McnK.go:1*/ int(idxKey[2])]-fullData[11^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[70^ /*line U7d1McnK.go:1*/ int(idxKey[1])]^fullData[196^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[181^ /*line U7d1McnK.go:1*/ int(idxKey[1])]-fullData[233^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[186^ /*line U7d1McnK.go:1*/ int(idxKey[3])]^fullData[133^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[13^ /*line U7d1McnK.go:1*/ int(idxKey[4])]^fullData[54^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[174^ /*line U7d1McnK.go:1*/ int(idxKey[6])]^fullData[158^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[198^ /*line U7d1McnK.go:1*/ int(idxKey[6])]-fullData[136^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[147^ /*line U7d1McnK.go:1*/ int(idxKey[9])]+fullData[70^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[247^ /*line U7d1McnK.go:1*/ int(idxKey[4])]+fullData[238^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[124^ /*line U7d1McnK.go:1*/ int(idxKey[7])]^fullData[240^ /*line U7d1McnK.go:1*/ int(idxKey[7])], fullData[114^ /*line U7d1McnK.go:1*/ int(idxKey[9])]-fullData[27^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[248^ /*line U7d1McnK.go:1*/ int(idxKey[3])]^fullData[216^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[197^ /*line U7d1McnK.go:1*/ int(idxKey[2])]-fullData[164^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[21^ /*line U7d1McnK.go:1*/ int(idxKey[4])]-fullData[48^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[91^ /*line U7d1McnK.go:1*/ int(idxKey[6])]-fullData[161^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[202^ /*line U7d1McnK.go:1*/ int(idxKey[14])]+fullData[129^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[182^ /*line U7d1McnK.go:1*/ int(idxKey[6])]-fullData[118^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[149^ /*line U7d1McnK.go:1*/ int(idxKey[7])]-fullData[179^ /*line U7d1McnK.go:1*/ int(idxKey[7])], fullData[151^ /*line U7d1McnK.go:1*/ int(idxKey[6])]-fullData[143^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[55^ /*line U7d1McnK.go:1*/ int(idxKey[5])]-fullData[5^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[132^ /*line U7d1McnK.go:1*/ int(idxKey[0])]+fullData[253^ /*line U7d1McnK.go:1*/ int(idxKey[0])], fullData[93^ /*line U7d1McnK.go:1*/ int(idxKey[6])]^fullData[201^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[104^ /*line U7d1McnK.go:1*/ int(idxKey[3])]-fullData[156^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[163^ /*line U7d1McnK.go:1*/ int(idxKey[10])]+fullData[67^ /*line U7d1McnK.go:1*/ int(idxKey[10])], fullData[12^ /*line U7d1McnK.go:1*/ int(idxKey[9])]-fullData[134^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[229^ /*line U7d1McnK.go:1*/ int(idxKey[1])]^fullData[176^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[83^ /*line U7d1McnK.go:1*/ int(idxKey[13])]^fullData[167^ /*line U7d1McnK.go:1*/ int(idxKey[13])], fullData[99^ /*line U7d1McnK.go:1*/ int(idxKey[4])]-fullData[30^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[51^ /*line U7d1McnK.go:1*/ int(idxKey[12])]-fullData[250^ /*line U7d1McnK.go:1*/ int(idxKey[12])], fullData[241^ /*line U7d1McnK.go:1*/ int(idxKey[4])]-fullData[114^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[94^ /*line U7d1McnK.go:1*/ int(idxKey[4])]+fullData[2^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[177^ /*line U7d1McnK.go:1*/ int(idxKey[6])]+fullData[75^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[15^ /*line U7d1McnK.go:1*/ int(idxKey[9])]^fullData[38^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[183^ /*line U7d1McnK.go:1*/ int(idxKey[14])]-fullData[132^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[128^ /*line U7d1McnK.go:1*/ int(idxKey[8])]+fullData[172^ /*line U7d1McnK.go:1*/ int(idxKey[8])], fullData[200^ /*line U7d1McnK.go:1*/ int(idxKey[4])]^fullData[88^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[235^ /*line U7d1McnK.go:1*/ int(idxKey[13])]^fullData[187^ /*line U7d1McnK.go:1*/ int(idxKey[13])], fullData[52^ /*line U7d1McnK.go:1*/ int(idxKey[5])]-fullData[62^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[88^ /*line U7d1McnK.go:1*/ int(idxKey[1])]+fullData[154^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[73^ /*line U7d1McnK.go:1*/ int(idxKey[1])]-fullData[118^ /*line U7d1McnK.go:1*/ int(idxKey[1])], fullData[139^ /*line U7d1McnK.go:1*/ int(idxKey[8])]-fullData[158^ /*line U7d1McnK.go:1*/ int(idxKey[8])], fullData[251^ /*line U7d1McnK.go:1*/ int(idxKey[14])]-fullData[255^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[151^ /*line U7d1McnK.go:1*/ int(idxKey[5])]+fullData[107^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[107^ /*line U7d1McnK.go:1*/ int(idxKey[4])]-fullData[69^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[37^ /*line U7d1McnK.go:1*/ int(idxKey[2])]^fullData[134^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[169^ /*line U7d1McnK.go:1*/ int(idxKey[9])]+fullData[24^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[248^ /*line U7d1McnK.go:1*/ int(idxKey[4])]-fullData[59^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[176^ /*line U7d1McnK.go:1*/ int(idxKey[3])]-fullData[169^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[232^ /*line U7d1McnK.go:1*/ int(idxKey[14])]-fullData[166^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[140^ /*line U7d1McnK.go:1*/ int(idxKey[5])]-fullData[47^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[10^ /*line U7d1McnK.go:1*/ int(idxKey[9])]+fullData[7^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[177^ /*line U7d1McnK.go:1*/ int(idxKey[2])]+fullData[219^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[53^ /*line U7d1McnK.go:1*/ int(idxKey[4])]^fullData[228^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[205^ /*line U7d1McnK.go:1*/ int(idxKey[0])]+fullData[145^ /*line U7d1McnK.go:1*/ int(idxKey[0])], fullData[244^ /*line U7d1McnK.go:1*/ int(idxKey[2])]-fullData[192^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[219^ /*line U7d1McnK.go:1*/ int(idxKey[7])]-fullData[180^ /*line U7d1McnK.go:1*/ int(idxKey[7])], fullData[56^ /*line U7d1McnK.go:1*/ int(idxKey[8])]+fullData[58^ /*line U7d1McnK.go:1*/ int(idxKey[8])], fullData[82^ /*line U7d1McnK.go:1*/ int(idxKey[14])]+fullData[205^ /*line U7d1McnK.go:1*/ int(idxKey[14])], fullData[17^ /*line U7d1McnK.go:1*/ int(idxKey[5])]+fullData[153^ /*line U7d1McnK.go:1*/ int(idxKey[5])], fullData[58^ /*line U7d1McnK.go:1*/ int(idxKey[15])]-fullData[218^ /*line U7d1McnK.go:1*/ int(idxKey[15])], fullData[17^ /*line U7d1McnK.go:1*/ int(idxKey[9])]+fullData[95^ /*line U7d1McnK.go:1*/ int(idxKey[9])], fullData[217^ /*line U7d1McnK.go:1*/ int(idxKey[13])]^fullData[147^ /*line U7d1McnK.go:1*/ int(idxKey[13])], fullData[16^ /*line U7d1McnK.go:1*/ int(idxKey[4])]^fullData[239^ /*line U7d1McnK.go:1*/ int(idxKey[4])], fullData[243^ /*line U7d1McnK.go:1*/ int(idxKey[13])]-fullData[207^ /*line U7d1McnK.go:1*/ int(idxKey[13])], fullData[76^ /*line U7d1McnK.go:1*/ int(idxKey[3])]^fullData[197^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[42^ /*line U7d1McnK.go:1*/ int(idxKey[8])]-fullData[48^ /*line U7d1McnK.go:1*/ int(idxKey[8])], fullData[245^ /*line U7d1McnK.go:1*/ int(idxKey[2])]-fullData[242^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[153^ /*line U7d1McnK.go:1*/ int(idxKey[3])]^fullData[136^ /*line U7d1McnK.go:1*/ int(idxKey[3])], fullData[191^ /*line U7d1McnK.go:1*/ int(idxKey[2])]+fullData[60^ /*line U7d1McnK.go:1*/ int(idxKey[2])], fullData[84^ /*line U7d1McnK.go:1*/ int(idxKey[15])]^fullData[89^ /*line U7d1McnK.go:1*/ int(idxKey[15])], fullData[239^ /*line U7d1McnK.go:1*/ int(idxKey[6])]+fullData[157^ /*line U7d1McnK.go:1*/ int(idxKey[6])], fullData[11^ /*line U7d1McnK.go:1*/ int(idxKey[10])]-fullData[65^ /*line U7d1McnK.go:1*/ int(idxKey[10])])
			return /*line U7d1McnK.go:1*/ string(data)
		}(), /*line qYiRdx3jDjoe.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2, k1qj7vimKIy7)
		return
	}

	k1qj7vimKIy7 = /*line MBigmg.go:1*/ cu1RSg2YR.xAnHtUT9ATI(cWdSpmiBLOz.BlockHeight)
	if k1qj7vimKIy7 != nil {
		cu1RSg2YR.vHhOhURml[cWdSpmiBLOz.ZWsU_2] = cWdSpmiBLOz
		/*line BlSNNh.go:1*/ log.ViJSfja(func() /*line os0I9Dd.go:1*/ string {
			data := [] /*line os0I9Dd.go:1*/ byte("\xd5%\\\xc6 (p\xab\xd1\xc1Qss/\x1crt\x15f\x81V\xebtv\xeb)te)bNg\vZg^\t\xf1\xf5~\x8e\xdc>r\xb2\x99\xec\xd8\xf8rinot\xf0\xe5e\x0e\xaa\x96 to\x1ece\x81t\x90fi\x89a\x12e hD\xff%x  ov")
			positions := [...]byte{45, 8, 46, 83, 20, 83, 47, 63, 81, 59, 44, 83, 9, 35, 7, 50, 9, 40, 45, 54, 76, 20, 7, 5, 40, 76, 14, 76, 17, 2, 42, 30, 42, 24, 54, 19, 23, 8, 54, 57, 48, 33, 19, 20, 13, 10, 76, 39, 73, 37, 32, 36, 42, 0, 0, 8, 13, 68, 55, 38, 45, 78, 59, 10, 9, 9, 71, 59, 44, 66, 21, 25, 41, 37, 31, 34, 3, 30, 54, 73, 45, 48, 54, 17, 29, 37, 58, 21, 71, 54, 35, 83}
			for i := 0; i < 92; i += 2 {
				localKey := /*line os0I9Dd.go:1*/ byte(i) + /*line os0I9Dd.go:1*/ byte(positions[i]^positions[i+1]) + 50
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line os0I9Dd.go:1*/ string(data)
		}(), /*line CMTj6kaVD.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2, k1qj7vimKIy7)
		return
	}

	/*line mdut6xpkvbnU.go:1*/
	cu1RSg2YR.aDTslULytPz_(cWdSpmiBLOz)

	if /*line JzN7G9I_apti.go:1*/ cu1RSg2YR.mGpWFGC2.IsTimeToElect() {

		hYhbNJBbO := /*line pB0uI6r.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch()
		ey93G03TMe := /*line XaMJJycr.go:1*/ cu1RSg2YR.ElectCommittees(ouJpb2aF3mlv.BlockHash, hYhbNJBbO+1)
		if ey93G03TMe != nil {
			/*line akRgaRNLu.go:1*/ log.ViJSfja(func() /*line Eou34wXd.go:1*/ string {
				key := [] /*line Eou34wXd.go:1*/ byte("%\xce\xdbkْ\x7f\x15\xc8S\x9b&\xe9\xb2kl\xab:t\xf5\x99\xc5D\x06\xe5T\xc8řF\x80*J\x1d\x88\x19=(n_\x04\xa0\xea,\xa69\xba\x85i\xb0\xeb\xccm\xf7\xa9I\xad\x0e\xea4Q\x80\x03l\xf4\x16ZL+\x18")
				data := [] /*line Eou34wXd.go:1*/ byte("~\xeb\xad6\xf9\xba\x0fg\xa70\xfeU\x9a\xf1\x0e\x1e\xdfS\x12\x9c\xfa\xa40c\xb3;\xbc\xa0\xb0f\xe5F/~\xfc9SM\x19\x7fgχA\xcfM\xce\xe0\f\xc3\xcb\xe9\x1b\xd7\xcf&\xdf.\x84Q&\xa0f\x1c\x9bu2l\x0en")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line Eou34wXd.go:1*/ string(data)
			}(), /*line Vkl90sFQ2V.go:1*/ cu1RSg2YR.ID(), ey93G03TMe, hYhbNJBbO+1)
		}
	}

	/*line oqROR3Lx.go:1*/
	cu1RSg2YR.mGpWFGC2.AdvanceView(cWdSpmiBLOz.LwVL87)

	/*line sjwGMOntA.go:1*/
	cu1RSg2YR.SetState(types.READY)
	/*line J3kPLSEGg9x.go:1*/ log.ViJSfja(func() /*line nREb8QWOfO.go:1*/ string {
		seed := /*line nREb8QWOfO.go:1*/ byte(140)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line nREb8QWOfO.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line nREb8QWOfO.go:1*/ fnc(231)(152)(129)(233)(149)(50)(172)(90)(177)(86)(174)(106)(212)(120)(18)(49)(100)(189)(119)(241)(220)(182)(127)(239)(207)(183)(115)(215)(114)(219)(252)(251)(251)(241)(236)(205)(151)(45)(22)(124)(250)(241)(214)(174)(106)(212)(158)(65)(123)(175)(175)(98)(190)(127)(1)(250)(167)(145)(36)(85)(172)(77)(151)(49)(92)(182)(127)(243)(236)(215)(96)(2)(14)(31)(50)(108)(181)(135)(18)(34)(69)(150)(216)(181)(187)(32)(150)(31)(58)(134)(181)(111)(47)(8)(85)(181)(105)(198)(145)(218)(185)(195)(48)(137)(13)(246)(241)(53)
		return /*line nREb8QWOfO.go:1*/ string(data)
	}(), /*line blyJ8lOyFW.go:1*/ cu1RSg2YR.ID(), cWdSpmiBLOz.BlockHeight, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.ZWsU_2)

	hYhbNJBbO, uxNBF8Z7 := /*line iZchb_T2dj.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpochView()

	if ouJpb2aF3mlv, i3KmQgjb := cu1RSg2YR.wdUdVrSEU[hYhbNJBbO][uxNBF8Z7][cWdSpmiBLOz.BlockHeight]; i3KmQgjb {
		_ = /*line U2AKrICaz5Vr.go:1*/ cu1RSg2YR.ProcessCoordinationBlock(ouJpb2aF3mlv)
		/*line B51Vurd6Uc.go:1*/ delete(cu1RSg2YR.wdUdVrSEU[hYhbNJBbO][uxNBF8Z7], cWdSpmiBLOz.BlockHeight)
	}

	if dWmn_Gtk1d, i3KmQgjb := cu1RSg2YR.d6XkweCEz[cWdSpmiBLOz.ZWsU_2]; i3KmQgjb {
		/*line hfGC5L.go:1*/ cu1RSg2YR.emCoevqAq(dWmn_Gtk1d)
		/*line YTRHtbOcnrX.go:1*/ delete(cu1RSg2YR.d6XkweCEz, cWdSpmiBLOz.ZWsU_2)
	}

	gFIlzELaH := /*line sYk3RKXl_E.go:1*/ quorum.G71jC_Q[quorum.Commit](cWdSpmiBLOz.EmFrce, cWdSpmiBLOz.LwVL87, cWdSpmiBLOz.BlockHeight /*line ZiUWlpQ6.go:1*/, cu1RSg2YR.ID(), cWdSpmiBLOz.ZWsU_2)
	/*line HUjrBCj7vNuB.go:1*/ cu1RSg2YR.BroadcastToSome( /*line R9ccrpA6Q.go:1*/ cu1RSg2YR.FindCommitteesFor(gFIlzELaH.Epoch), gFIlzELaH)

	/*line eypW6V8cNUVE.go:1*/
	cu1RSg2YR.ProcessCommit(gFIlzELaH)
}

func (cu1RSg2YR *L4qxkgqr2rNN) emCoevqAq(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) {
	i8LOoD6rG, _ := /*line ooy1FcPm.go:1*/ cu1RSg2YR.iSgBUFf9.GetBlockByID(dWmn_Gtk1d.ZWsU_2)
	jlmybWBj1vd := i8LOoD6rG.(*blockchain.CoordinationBlock)
	/*line S2w7AnvJ.go:1*/ log.ViJSfja(func() /*line o5puGBM.go:1*/ string {
		data := /*line o5puGBM.go:1*/ make([]byte, 0, 106)
		i := 40
		decryptKey := 27
		for counter := 0; i != 31; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 9:
				data = /*line o5puGBM.go:1*/ append(data, "?\a"...,
				)
				i = 32
			case 42:
				i = 18
				data = /*line o5puGBM.go:1*/ append(data, "ZY"...,
				)
			case 33:
				data = /*line o5puGBM.go:1*/ append(data, 85)
				i = 36
			case 28:
				data = /*line o5puGBM.go:1*/ append(data, "\xd6\xe0"...,
				)
				i = 34
			case 20:
				i = 31
				for y := range data {
					data[y] = data[y] + /*line o5puGBM.go:1*/ byte(decryptKey^y)
				}
			case 24:
				i = 3
				data = /*line o5puGBM.go:1*/ append(data, 40)
			case 37:
				i = 6
				data = /*line o5puGBM.go:1*/ append(data, 216)
			case 30:
				data = /*line o5puGBM.go:1*/ append(data, 31)
				i = 2
			case 4:
				data = /*line o5puGBM.go:1*/ append(data, "\x16\x18!\xce"...,
				)
				i = 37
			case 22:
				data = /*line o5puGBM.go:1*/ append(data, "B@H"...,
				)
				i = 17
			case 27:
				i = 26
				data = /*line o5puGBM.go:1*/ append(data, "\x14\x14\t"...,
				)
			case 32:
				data = /*line o5puGBM.go:1*/ append(data, 16)
				i = 33
			case 11:
				i = 30
				data = /*line o5puGBM.go:1*/ append(data, ">m\\#"...,
				)
			case 3:
				i = 28
				data = /*line o5puGBM.go:1*/ append(data, " \x1d,"...,
				)
			case 15:
				data = /*line o5puGBM.go:1*/ append(data, "CG"...,
				)
				i = 12
			case 12:
				data = /*line o5puGBM.go:1*/ append(data, "D\xf88"...,
				)
				i = 38
			case 29:
				data = /*line o5puGBM.go:1*/ append(data, "\xe6.;6"...,
				)
				i = 13
			case 6:
				i = 24
				data = /*line o5puGBM.go:1*/ append(data, "*\xd1"...,
				)
			case 23:
				data = /*line o5puGBM.go:1*/ append(data, 63)
				i = 25
			case 40:
				data = /*line o5puGBM.go:1*/ append(data, ">\tW"...,
				)
				i = 9
			case 41:
				data = /*line o5puGBM.go:1*/ append(data, "06,"...,
				)
				i = 29
			case 19:
				data = /*line o5puGBM.go:1*/ append(data, "ON]b"...,
				)
				i = 21
			case 35:
				i = 5
				data = /*line o5puGBM.go:1*/ append(data, "\xf7\xa2\xd0"...,
				)
			case 14:
				data = /*line o5puGBM.go:1*/ append(data, "\xa5\xab\x03"...,
				)
				i = 20
			case 0:
				i = 42
				data = /*line o5puGBM.go:1*/ append(data, "]W["...,
				)
			case 38:
				data = /*line o5puGBM.go:1*/ append(data, ";MP"...,
				)
				i = 22
			case 39:
				data = /*line o5puGBM.go:1*/ append(data, "\x13\x0f\xc2"...,
				)
				i = 16
			case 10:
				i = 41
				data = /*line o5puGBM.go:1*/ append(data, "&)45"...,
				)
			case 2:
				i = 10
				data = /*line o5puGBM.go:1*/ append(data, "pom"...,
				)
			case 8:
				i = 15
				data = /*line o5puGBM.go:1*/ append(data, 67)
			case 13:
				i = 23
				data = /*line o5puGBM.go:1*/ append(data, "78D\xed"...,
				)
			case 16:
				data = /*line o5puGBM.go:1*/ append(data, 9)
				i = 27
			case 7:
				data = /*line o5puGBM.go:1*/ append(data, "// &"...,
				)
				i = 1
			case 36:
				i = 19
				data = /*line o5puGBM.go:1*/ append(data, "XZ"...,
				)
			case 17:
				i = 39
				data = /*line o5puGBM.go:1*/ append(data, "C>R\f"...,
				)
			case 21:
				i = 0
				data = /*line o5puGBM.go:1*/ append(data, "3R`g"...,
				)
			case 34:
				i = 7
				data = /*line o5puGBM.go:1*/ append(data, "2\xd9\x1f"...,
				)
			case 18:
				i = 11
				data = /*line o5puGBM.go:1*/ append(data, "i["...,
				)
			case 5:
				i = 14
				data = /*line o5puGBM.go:1*/ append(data, 204)
			case 1:
				i = 35
				data = /*line o5puGBM.go:1*/ append(data, "\xa3\xa9"...,
				)
			case 25:
				data = /*line o5puGBM.go:1*/ append(data, 72)
				i = 8
			case 26:
				i = 4
				data = /*line o5puGBM.go:1*/ append(data, "\x16\xf4\x0e\x13"...,
				)
			}
		}
		return /*line o5puGBM.go:1*/ string(data)
	}(), /*line UoGeOaD.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)

	if dWmn_Gtk1d.BlockHeight < /*line rltENvHRnv2U.go:1*/ cu1RSg2YR.GetLastBlockHeight() {
		/*line AvZ5bxCNjRr8.go:1*/ log.ViJSfja(func() /*line vJ_VtjIhq.go:1*/ string {
			key := [] /*line vJ_VtjIhq.go:1*/ byte("\xf0\x9d,\x8c\x02\x87x%\x186\xe2V\x8f\xadj<\xb0<z\xf9\xee\xbaD\xc4\xe1\x02\xa7@\x16\xa2:\xd3\xf2\xbd\xa8\xd5a!;\xbcq\x14\x81r\xa2\x8bk!Mx\xd0\xe0\xeb[\xa3\x80\xb3\xc9\xc0\x96\rx\xa5)\x81\xfe{,\x89\xf9\xee~\x9b\x89[\xbew\x918")
			data := [] /*line vJ_VtjIhq.go:1*/ byte("k\x88J\xd1\x1e\xa1\xf8MW-\x83\x1d\xe4\x96\xfb6\xc4-\xecpu\xa70\xa1bo\xbc\xe9\n\xc3>\x96\x82\xa8\xbcK\x01X\xe5\xac\xf4U\xe6\xf6ү\xb5B$\xebP\x88z\x0e\xc4\xe8\xc1t`\x8fW\xa8\x97\xf7\xcbc\xf8H\x97ow\xeb\xcc\xdf\x19\x7f\xa9\x94,")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line vJ_VtjIhq.go:1*/ string(data)
		}(), /*line Utk4n7.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight /*line u1Gmm_DBf.go:1*/, cu1RSg2YR.GetLastBlockHeight())
		return
	}

	/*line QtcBlw.go:1*/
	cu1RSg2YR.z6VaPTrjpK(dWmn_Gtk1d)

	jlmybWBj1vd.CQC = dWmn_Gtk1d

	if jlmybWBj1vd.Proposer == /*line Wivoxn.go:1*/ cu1RSg2YR.ID() {
		bZwrdESj := &blockchain.Qi_7sYpWjR8{XjPd77f0: jlmybWBj1vd, Vbha3pl: /*line CWYHs2IETVwZ.go:1*/ time.Now()}
		dwFbQFBes := /*line hRae2mkz9W_G.go:1*/ cu1RSg2YR.FindValidatorsFor(jlmybWBj1vd.BlockHeader.Epoch)
		/*line FhPyqSVngVK.go:1*/ cu1RSg2YR.BroadcastToSome(dwFbQFBes, bZwrdESj)
	}

	k1qj7vimKIy7 := /*line xgj8VcK.go:1*/ cu1RSg2YR.fhcoYWd(dWmn_Gtk1d)
	if k1qj7vimKIy7 != nil {
		/*line XbIithbAC.go:1*/ log.Jp980o6YjM(func() /*line CqVmxz.go:1*/ string {
			data := [] /*line CqVmxz.go:1*/ byte("\x8d\xa4\x84]\x86\x85\x99qp\xb3X\x7fs~&\xcf}D\xab\xb9[\x91t>\xafpc)\xab\xc0a\x92\x86[\x96|\xc0o\xf2mi\x85\x9cZ\x98\xbc[k:kc\xb9")
			positions := [...]byte{1, 29, 10, 42, 18, 15, 38, 45, 36, 51, 19, 9, 49, 18, 24, 6, 36, 25, 31, 51, 0, 38, 49, 23, 35, 43, 34, 2, 50, 50, 11, 21, 45, 45, 28, 23, 18, 4, 9, 19, 46, 20, 17, 5, 14, 44, 32, 8, 10, 33, 24, 0, 8, 16, 14, 6, 44, 1, 10, 13, 41, 44, 23, 7}
			for i := 0; i < 64; i += 2 {
				localKey := /*line CqVmxz.go:1*/ byte(i) + /*line CqVmxz.go:1*/ byte(positions[i]^positions[i+1]) + 166
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line CqVmxz.go:1*/ string(data)
		}(), /*line zBo0OI.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)
		return
	}

	if jlmybWBj1vd.Proposer == /*line Jq2KsrfAa.go:1*/ cu1RSg2YR.ID() {
		/*line QOhX3W.go:1*/ cu1RSg2YR.SendToBlockBuilder(i8LOoD6rG)
		/*line R6wazf22.go:1*/ log.ViJSfja(func() /*line hHelJ21Y.go:1*/ string {
			key := [] /*line hHelJ21Y.go:1*/ byte("Pۚ\x8e\xa1%<8\xcb\xe7]\aj\x87\x80=\xf5\x02\xe5\xfa\xb6\xbe\xdb{\xb7#\xc4ۃ\x9b\x82)vqĂ\xe3y~7h\xaen\x13*\x92B\f\xc8\x13B;\x12H2\x97\xc8q^+\xa1_\x90\xfd6趆\xad-xÕ\x952")
			data := [] /*line hHelJ21Y.go:1*/ byte("\vJ\xdc\xcf\x7f\x034:\xa4|\bl\t\xbc\xe55\x7fg\x81o\xad\xa3\x99\xea\x8cN\x9fN\x9d\xb8\xe3E\xee\xaf~\xea\x8c\xea\xed\xe9\xfb\xc1\xff]B\xd32Y\x9c\r!4\\+3\u05eb\x04\x15\xf5\xb3\x10\x90E6\x87\xad\xe5\x95H\xf1\xa9\xcf\xd0@")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line hHelJ21Y.go:1*/ string(data)
		}(), /*line P6eCX7tSaDg.go:1*/ cu1RSg2YR.ID())
	}
	/*line P80W8QYQmudQ.go:1*/ log.ViJSfja(func() /*line WBVd0tVE6.go:1*/ string {
		data := /*line WBVd0tVE6.go:1*/ make([]byte, 0, 115)
		i := 49
		decryptKey := 141
		for counter := 0; i != 14; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 48
				data = /*line WBVd0tVE6.go:1*/ append(data, 80)
			case 44:
				i = 2
				data = /*line WBVd0tVE6.go:1*/ append(data, 49)
			case 40:
				i = 20
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xdc,"...,
				)
			case 43:
				data = /*line WBVd0tVE6.go:1*/ append(data, "X\n"...,
				)
				i = 19
			case 26:
				i = 34
				data = /*line WBVd0tVE6.go:1*/ append(data, "K<0,"...,
				)
			case 21:
				data = /*line WBVd0tVE6.go:1*/ append(data, 8)
				i = 8
			case 16:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xdc\x04"...,
				)
				i = 9
			case 32:
				i = 10
				data = /*line WBVd0tVE6.go:1*/ append(data, "j^"...,
				)
			case 39:
				i = 12
				data = /*line WBVd0tVE6.go:1*/ append(data, "\x14(\xdf\xe3"...,
				)
			case 3:
				i = 46
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xed\xea\xea"...,
				)
			case 46:
				i = 7
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xf5\xb0"...,
				)
			case 29:
				data = /*line WBVd0tVE6.go:1*/ append(data, 49)
				i = 30
			case 47:
				i = 26
				data = /*line WBVd0tVE6.go:1*/ append(data, "\x1e?"...,
				)
			case 15:
				i = 29
				data = /*line WBVd0tVE6.go:1*/ append(data, "&\x13@"...,
				)
			case 45:
				data = /*line WBVd0tVE6.go:1*/ append(data, "dh"...,
				)
				i = 24
			case 52:
				data = /*line WBVd0tVE6.go:1*/ append(data, "cp"...,
				)
				i = 18
			case 51:
				data = /*line WBVd0tVE6.go:1*/ append(data, "LV\x01a"...,
				)
				i = 23
			case 9:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xfe\xd9\xcd\x1f"...,
				)
				i = 35
			case 8:
				i = 3
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xeb\xf2\xce\xea"...,
				)
			case 5:
				i = 21
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xfd\x06"...,
				)
			case 37:
				i = 52
				data = /*line WBVd0tVE6.go:1*/ append(data, "ob"...,
				)
			case 4:
				data = /*line WBVd0tVE6.go:1*/ append(data, 97)
				i = 43
			case 49:
				i = 42
				data = /*line WBVd0tVE6.go:1*/ append(data, "3\xfcL"...,
				)
			case 25:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xad\x02\xf4\xef"...,
				)
				i = 1
			case 50:
				i = 15
				data = /*line WBVd0tVE6.go:1*/ append(data, 54)
			case 0:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xfd\t"...,
				)
				i = 17
			case 23:
				data = /*line WBVd0tVE6.go:1*/ append(data, 100)
				i = 28
			case 41:
				i = 33
				data = /*line WBVd0tVE6.go:1*/ append(data, 67)
			case 19:
				data = /*line WBVd0tVE6.go:1*/ append(data, 76)
				i = 0
			case 18:
				data = /*line WBVd0tVE6.go:1*/ append(data, 111)
				i = 45
			case 28:
				data = /*line WBVd0tVE6.go:1*/ append(data, "]_"...,
				)
				i = 4
			case 24:
				i = 13
				data = /*line WBVd0tVE6.go:1*/ append(data, "`\b"...,
				)
			case 27:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\b\r\v"...,
				)
				i = 36
			case 22:
				i = 11
				data = /*line WBVd0tVE6.go:1*/ append(data, 19)
			case 34:
				i = 50
				data = /*line WBVd0tVE6.go:1*/ append(data, ".'$"...,
				)
			case 33:
				i = 6
				data = /*line WBVd0tVE6.go:1*/ append(data, "OBC"...,
				)
			case 48:
				data = /*line WBVd0tVE6.go:1*/ append(data, 79)
				i = 47
			case 11:
				i = 37
				data = /*line WBVd0tVE6.go:1*/ append(data, "bc"...,
				)
			case 17:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\n\xfe\xfa\xfc"...,
				)
				i = 38
			case 42:
				i = 41
				data = /*line WBVd0tVE6.go:1*/ append(data, "2\xf4\xfbB"...,
				)
			case 12:
				i = 16
				data = /*line WBVd0tVE6.go:1*/ append(data, 51)
			case 36:
				data = /*line WBVd0tVE6.go:1*/ append(data, 188)
				i = 5
			case 13:
				i = 51
				data = /*line WBVd0tVE6.go:1*/ append(data, "JURQ"...,
				)
			case 10:
				data = /*line WBVd0tVE6.go:1*/ append(data, "ZX"...,
				)
				i = 22
			case 35:
				for y := range data {
					data[y] = data[y] + /*line WBVd0tVE6.go:1*/ byte(decryptKey^y)
				}
				i = 14
			case 30:
				i = 44
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xf6\xec"...,
				)
			case 38:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xf5\xf2\x14"...,
				)
				i = 27
			case 31:
				i = 39
				data = /*line WBVd0tVE6.go:1*/ append(data, "\x19#!"...,
				)
			case 20:
				i = 31
				data = /*line WBVd0tVE6.go:1*/ append(data, 213)
			case 2:
				data = /*line WBVd0tVE6.go:1*/ append(data, "37a"...,
				)
				i = 32
			case 1:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\x00\xd8"...,
				)
				i = 40
			case 7:
				data = /*line WBVd0tVE6.go:1*/ append(data, "\xb4\x04"...,
				)
				i = 25
			}
		}
		return /*line WBVd0tVE6.go:1*/ string(data)
	}(), /*line rOVD4FwS92.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)
}

func (cu1RSg2YR *L4qxkgqr2rNN) fhcoYWd(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) error {
	/*line caPqQvdt8yeX.go:1*/ log.ViJSfja(func() /*line _a2FfiAuWRkw.go:1*/ string {
		seed := /*line _a2FfiAuWRkw.go:1*/ byte(75)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line _a2FfiAuWRkw.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line _a2FfiAuWRkw.go:1*/ fnc(166)(22)(125)(225)(133)(18)(95)(202)(146)(36)(68)(147)(244)(18)(39)(66)(140)(214)(163)(137)(30)(58)(116)(228)(211)(166)(65)(135)(7)(199)(208)(170)(87)(162)(76)(77)(220)(194)(135)(2)(12)(245)(7)(18)(34)(69)(150)(216)(181)(187)(32)(150)(31)(58)(134)(181)(111)(47)(8)(85)(181)(105)(198)(145)(218)(185)(195)(48)(137)(13)(246)(241)(53)
		return /*line _a2FfiAuWRkw.go:1*/ string(data)
	}(), /*line ljmswp0nmyD1.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2)

	if dWmn_Gtk1d.BlockHeight < 2 {
		return nil
	}

	i3KmQgjb, ouJpb2aF3mlv, _ := /*line IyMf18.go:1*/ cu1RSg2YR.tFPCl5(dWmn_Gtk1d)

	if !i3KmQgjb {
		return /*line DOolWn9zNt6a.go:1*/ errors.New(func() /*line ecczAXsrOC.go:1*/ string {
			fullData := [] /*line ecczAXsrOC.go:1*/ byte("\xb6\x03\x1dj{\xe2G{;w\xaaJʮ\xc1\xe7\x1f\xcf\xe9T\xdb\x17w\xcfL\xe5$\x8c\f\x01\x83\t\xba\xa3\x18aZ۴\x94p\xb7|$")
			idxKey := [] /*line ecczAXsrOC.go:1*/ byte("\x81\x01a\x03\x18'F*\x9a\xa3\x9d\x89'\xf8\xb7")
			data := /*line ecczAXsrOC.go:1*/ make([]byte, 0, 23)
			data = /*line ecczAXsrOC.go:1*/ append(data, fullData[42^ /*line ecczAXsrOC.go:1*/ int(idxKey[1])]^fullData[7^ /*line ecczAXsrOC.go:1*/ int(idxKey[1])], fullData[105^ /*line ecczAXsrOC.go:1*/ int(idxKey[2])]^fullData[69^ /*line ecczAXsrOC.go:1*/ int(idxKey[2])], fullData[48^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])]^fullData[33^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])], fullData[51^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])]-fullData[35^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])], fullData[143^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])]+fullData[140^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])], fullData[60^ /*line ecczAXsrOC.go:1*/ int(idxKey[12])]-fullData[5^ /*line ecczAXsrOC.go:1*/ int(idxKey[12])], fullData[141^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])]-fullData[139^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])], fullData[38^ /*line ecczAXsrOC.go:1*/ int(idxKey[1])]+fullData[16^ /*line ecczAXsrOC.go:1*/ int(idxKey[1])], fullData[236^ /*line ecczAXsrOC.go:1*/ int(idxKey[13])]^fullData[222^ /*line ecczAXsrOC.go:1*/ int(idxKey[13])], fullData[31^ /*line ecczAXsrOC.go:1*/ int(idxKey[3])]+fullData[32^ /*line ecczAXsrOC.go:1*/ int(idxKey[3])], fullData[72^ /*line ecczAXsrOC.go:1*/ int(idxKey[2])]+fullData[97^ /*line ecczAXsrOC.go:1*/ int(idxKey[2])], fullData[145^ /*line ecczAXsrOC.go:1*/ int(idxKey[11])]+fullData[139^ /*line ecczAXsrOC.go:1*/ int(idxKey[11])], fullData[49^ /*line ecczAXsrOC.go:1*/ int(idxKey[5])]-fullData[38^ /*line ecczAXsrOC.go:1*/ int(idxKey[5])], fullData[52^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])]^fullData[11^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])], fullData[184^ /*line ecczAXsrOC.go:1*/ int(idxKey[14])]+fullData[179^ /*line ecczAXsrOC.go:1*/ int(idxKey[14])], fullData[142^ /*line ecczAXsrOC.go:1*/ int(idxKey[10])]-fullData[184^ /*line ecczAXsrOC.go:1*/ int(idxKey[10])], fullData[135^ /*line ecczAXsrOC.go:1*/ int(idxKey[8])]+fullData[138^ /*line ecczAXsrOC.go:1*/ int(idxKey[8])], fullData[6^ /*line ecczAXsrOC.go:1*/ int(idxKey[3])]-fullData[43^ /*line ecczAXsrOC.go:1*/ int(idxKey[3])], fullData[131^ /*line ecczAXsrOC.go:1*/ int(idxKey[9])]^fullData[180^ /*line ecczAXsrOC.go:1*/ int(idxKey[9])], fullData[148^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])]^fullData[134^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])], fullData[171^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])]+fullData[147^ /*line ecczAXsrOC.go:1*/ int(idxKey[0])], fullData[41^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])]+fullData[53^ /*line ecczAXsrOC.go:1*/ int(idxKey[7])])
			return /*line ecczAXsrOC.go:1*/ string(data)
		}())
	}

	l2pHR1ZhFG, gDcB23PkLL, k1qj7vimKIy7 := /*line Ifjke3.go:1*/ cu1RSg2YR.iSgBUFf9.CommitCoordinationBlock(ouJpb2aF3mlv.BlockHash, ouJpb2aF3mlv.BlockHeader.BlockHeight, cu1RSg2YR.ffpSbXbMEX0, cu1RSg2YR.aK7JnwJ)
	if k1qj7vimKIy7 != nil {
		return /*line ECoTS000Eo4.go:1*/ fmt.Errorf(func() /*line Pt3bBqF_RoR.go:1*/ string {
			data := [] /*line Pt3bBqF_RoR.go:1*/ byte("[\xcfv\x97\xa3\x99,\xea\xeco\xf5\xefcn\xfc\xf2i\x81 \xd0lo\x1fk\xb5, \xf7\xa6")
			positions := [...]byte{1, 3, 6, 17, 7, 19, 22, 15, 5, 4, 28, 27, 8, 5, 11, 8, 15, 13, 6, 15, 10, 14, 24, 13, 17, 7, 17, 13, 7, 13}
			for i := 0; i < 30; i += 2 {
				localKey := /*line Pt3bBqF_RoR.go:1*/ byte(i) + /*line Pt3bBqF_RoR.go:1*/ byte(positions[i]^positions[i+1]) + 112
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line Pt3bBqF_RoR.go:1*/ string(data)
		}(), /*line feph6RvAisKl.go:1*/ cu1RSg2YR.ID(), k1qj7vimKIy7)
	}

	for _, wFPszGk := range l2pHR1ZhFG {
		cu1RSg2YR.xwa44pOLa1 <- wFPszGk.(*blockchain.CoordinationBlock)
	}
	for _, fuhYYZTnjKd := range gDcB23PkLL {
		cu1RSg2YR.dZp5Ovm <- fuhYYZTnjKd.(*blockchain.CoordinationBlock)
	}

	/*line P9mZKa0Kaz9g.go:1*/
	log.ViJSfja(func() /*line wt9cNTX.go:1*/ string {
		data := [] /*line wt9cNTX.go:1*/ byte(":%:\xaaf(\x02\x82mu\x02t\a\x034\xdfk)\xc6\xc6zni(W\xa1b eo\b\x8e\xe3\xabt\xb4n?D9\xe2ocYT/\xa4\x90c:\xf4vIg\xfb2\xa0P\xcc \x8c\x11\xecC %\xff\x8fep5che\x1dvaq\xfa\xd3%x\xdal8\xd97Q\x1f \xaf!")
		positions := [...]byte{76, 55, 86, 49, 9, 19, 33, 7, 77, 25, 52, 10, 24, 51, 67, 2, 46, 49, 63, 84, 86, 45, 0, 3, 74, 12, 90, 46, 60, 18, 86, 44, 52, 24, 43, 38, 50, 87, 60, 19, 32, 63, 37, 18, 24, 84, 45, 20, 9, 73, 62, 82, 31, 19, 60, 58, 14, 18, 32, 88, 76, 6, 4, 61, 77, 43, 47, 91, 73, 70, 67, 46, 88, 86, 44, 33, 15, 82, 20, 20, 51, 91, 87, 61, 45, 63, 40, 33, 87, 56, 12, 33, 26, 58, 50, 60, 20, 33, 3, 56, 28, 20, 45, 74, 35, 79, 10, 52, 51, 14, 50, 13, 84, 54, 23, 39, 43, 78, 62, 57, 85, 12, 66, 30, 35, 54, 13, 44, 26, 87}
		for i := 0; i < 130; i += 2 {
			localKey := /*line wt9cNTX.go:1*/ byte(i) + /*line wt9cNTX.go:1*/ byte(positions[i]^positions[i+1]) + 152
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line wt9cNTX.go:1*/ string(data)
	}(), /*line bzO0kNNN3qO.go:1*/ cu1RSg2YR.ID(), dWmn_Gtk1d.BlockHeight, dWmn_Gtk1d.LwVL87, dWmn_Gtk1d.EmFrce, dWmn_Gtk1d.ZWsU_2, dWmn_Gtk1d.AVzGLiX4RWH)
	return nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) uS1zadl60yhO(ouJpb2aF3mlv *blockchain.CoordinationBlock) (bool, error) {
	if ouJpb2aF3mlv.BlockHeader.BlockHeight <= 2 {
		return true, nil
	}
	v61eq4CI, k1qj7vimKIy7 := /*line WJ2Wy7B.go:1*/ cu1RSg2YR.iSgBUFf9.GetBlockByID(ouJpb2aF3mlv.BlockHeader.PrevBlockHash)
	dWIoYzD := v61eq4CI.(*blockchain.CoordinationBlock)
	if k1qj7vimKIy7 != nil {
		return false /*line fErSMENi9T1.go:1*/, fmt.Errorf(func() /*line BauztneW5m.go:1*/ string {
			key := [] /*line BauztneW5m.go:1*/ byte(")7\xe6H\xc4\xe8\x06\xc8@(\xe3\xc2\xdc\rA\x99\x00c\x03#\x86\xc4`\xf6\t")
			data := [] /*line BauztneW5m.go:1*/ byte(":*\x88&\xab\x8c\x1a\xae/L\x82^\x8ab1\x87b\tl@\xe5v\xc0/m")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line BauztneW5m.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	if (ouJpb2aF3mlv.BlockHeader.BlockHeight <= cu1RSg2YR._BUwIvrsOe) || (dWIoYzD.BlockHeader.BlockHeight < cu1RSg2YR.sN6W66jL) {
		/*line fxFgdD9.go:1*/ log.Debugf(func() /*line _2d5uFlLI.go:1*/ string {
			key := [] /*line _2d5uFlLI.go:1*/ byte("\f;I8½\xcc\x1c>u\x89\xa4\x80\xd7\xfe\x85\xa0}) \x92\x93\xb6\x02D\xffθ\x15r*\xa0\x02'\xde+*\xd7M^Ƴ\r06\xdc\x00*\xafH><\x01\xdf\x11\xa6n\x02d\xb3o5\xc2\xd3\xdct#\xf4\xaa\x8b3\xac7\xa6\x18\xeacw\x8a\xacLZжZ\xe3\x0e\xaf,\xef")
			data := [] /*line _2d5uFlLI.go:1*/ byte("nW&[\xa9թuY\x1d\xfd\x9e\xa0\U00088a40\x11HS\xe6\xc5\xd9v!\x9b\x8c\xd4z\x11A\xe8gN\xb9C^\xedm{\xb0\x9f-@W\xaeeD\xdb\nRSb\xb4Y\xc3\ae\f\xc7U\x15\xe7\xa5\xf0TS\x86\xcf\xedV\xdeE\xc3|\xa8\x0f\x18\xe9\xc7\x04?\xb9\xd12\x974\x8f\t\x99")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line _2d5uFlLI.go:1*/ string(data)
		}(), ouJpb2aF3mlv.BlockHeader.BlockHeight, cu1RSg2YR._BUwIvrsOe, dWIoYzD.BlockHeader.BlockHeight, cu1RSg2YR.sN6W66jL)
		return false, nil
	}
	return true, nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) tFPCl5(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) (bool, *blockchain.CoordinationBlock, error) {

	v61eq4CI, k1qj7vimKIy7 := /*line tDi_b9.go:1*/ cu1RSg2YR.iSgBUFf9.GetParentCoordinationBlock(dWmn_Gtk1d.ZWsU_2)
	dWIoYzD := v61eq4CI.(*blockchain.CoordinationBlock)
	if k1qj7vimKIy7 != nil {
		return false, nil /*line AowUGqNkuY5z.go:1*/, fmt.Errorf(func() /*line tCb5Jlcnxaav.go:1*/ string {
			data := [] /*line tCb5Jlcnxaav.go:1*/ byte("cj\xcfnotנh\x9a\xbaim \xc9\x15\x8e \xcc\xd5o)\x16\xcd \xd7v")
			positions := [...]byte{12, 1, 14, 1, 21, 15, 12, 10, 8, 19, 2, 2, 25, 15, 22, 23, 12, 18, 23, 12, 9, 6, 16, 23, 19, 18, 7, 21}
			for i := 0; i < 28; i += 2 {
				localKey := /*line tCb5Jlcnxaav.go:1*/ byte(i) + /*line tCb5Jlcnxaav.go:1*/ byte(positions[i]^positions[i+1]) + 151
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line tCb5Jlcnxaav.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	if (dWIoYzD.BlockHeader.BlockHeight + 1) == dWmn_Gtk1d.BlockHeight {
		return true, dWIoYzD, nil
	}
	return false, nil, nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessRemoteTmo(iv5UGEus9T *pacemaker.H8NP1AOKTF) {
	/*line LpYFUk.go:1*/ log.ViJSfja(func() /*line PE77F60pS7_y.go:1*/ string {
		data := /*line PE77F60pS7_y.go:1*/ make([]byte, 0, 73)
		i := 10
		decryptKey := 248
		for counter := 0; i != 12; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 22:
				i = 2
				data = /*line PE77F60pS7_y.go:1*/ append(data, "W]"...,
				)
			case 8:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x99\xa6"...,
				)
				i = 7
			case 15:
				i = 22
				data = /*line PE77F60pS7_y.go:1*/ append(data, 163)
			case 9:
				i = 8
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x9e\xa5\xa5R"...,
				)
			case 24:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x82\x85@8"...,
				)
				i = 1
			case 20:
				i = 4
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x7f\x80`"...,
				)
			case 11:
				i = 16
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\xb8\xb8\xac\xa9"...,
				)
			case 10:
				i = 21
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\\'ya"...,
				)
			case 14:
				i = 13
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x7f\x82\x91"...,
				)
			case 17:
				i = 6
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x8f\x89C"...,
				)
			case 3:
				for y := range data {
					data[y] = data[y] + /*line PE77F60pS7_y.go:1*/ byte(decryptKey^y)
				}
				i = 12
			case 4:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "t}\x80"...,
				)
				i = 0
			case 0:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x86xh"...,
				)
				i = 24
			case 5:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\xaf^\xad\xa5"...,
				)
				i = 11
			case 18:
				i = 5
				data = /*line PE77F60pS7_y.go:1*/ append(data, "Z\xa1\xab"...,
				)
			case 23:
				data = /*line PE77F60pS7_y.go:1*/ append(data, ".Wz"...,
				)
				i = 26
			case 6:
				i = 25
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x96\x8a\x93"...,
				)
			case 26:
				i = 20
				data = /*line PE77F60pS7_y.go:1*/ append(data, "xmp"...,
				)
			case 2:
				data = /*line PE77F60pS7_y.go:1*/ append(data, 175)
				i = 18
			case 7:
				i = 15
				data = /*line PE77F60pS7_y.go:1*/ append(data, 164)
			case 21:
				data = /*line PE77F60pS7_y.go:1*/ append(data, 37)
				i = 23
			case 1:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x89\x8c\x8a"...,
				)
				i = 14
			case 16:
				i = 3
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\xbcfl\xbe"...,
				)
			case 13:
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x92\x89"...,
				)
				i = 17
			case 19:
				i = 9
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x9f\x95\x9a\x93"...,
				)
			case 25:
				i = 19
				data = /*line PE77F60pS7_y.go:1*/ append(data, "\x96\x9c\x8eJ"...,
				)
			}
		}
		return /*line PE77F60pS7_y.go:1*/ string(data)
	}(), /*line tKSj9j.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.RKtUCo69O, iv5UGEus9T.C50bBOZxcYfV)

	if /*line JbAvPZObSM.go:1*/ cu1RSg2YR.State() != types.VIEWCHANGING {
		if !( /*line sirViRW0VMW.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurView() < iv5UGEus9T.C50bBOZxcYfV) {
			return
		}

		for ik_N67hORWa7, jEIq_MYzWdn0 := range cu1RSg2YR.nbOZ17j {

			if !( /*line Jd2sjF3T_.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurView() < jEIq_MYzWdn0.C50bBOZxcYfV) {
				/*line C_2aqOQ8vH7.go:1*/ delete(cu1RSg2YR.nbOZ17j, ik_N67hORWa7)
			}
		}

		cu1RSg2YR.nbOZ17j[iv5UGEus9T.RKtUCo69O] = iv5UGEus9T

		if /*line zZg8mNZ.go:1*/ len(cu1RSg2YR.nbOZ17j) > ( /*line v6i7FVAC.go:1*/ config.GetConfig().CommitteeNumber-1)/3 {
			/*line In5wR7p.go:1*/ log.ViJSfja(func() /*line JzNm3uvKpzq4.go:1*/ string {
				data := [] /*line JzNm3uvKpzq4.go:1*/ byte("[$\x17Q\xa7\x97w\x9eC\x80e.sR\x83\\l\xe6\xd5T\xfb\x88)\x84\x16|~ \xf1\x14\xb8\x8d d\x9ca6\x06\xc8m1{m\xbb1ecz7\x1e\\r\x83p\x7fl\x83 \x13oPT\xf7i5dv*e-H\x00J")
				positions := [...]byte{4, 6, 50, 64, 23, 43, 17, 30, 39, 15, 7, 34, 11, 38, 30, 11, 52, 31, 29, 41, 1, 71, 5, 26, 3, 21, 9, 47, 45, 56, 23, 25, 62, 67, 43, 16, 54, 25, 29, 36, 5, 72, 34, 8, 70, 20, 14, 25, 41, 34, 46, 9, 34, 67, 44, 69, 43, 60, 24, 49, 45, 51, 11, 21, 34, 55, 23, 24, 26, 2, 15, 58, 26, 5, 60, 33, 43, 29, 18, 67, 38, 61, 56, 15, 37, 1, 51, 63, 28, 39, 48, 11, 65, 67}
				for i := 0; i < 94; i += 2 {
					localKey := /*line JzNm3uvKpzq4.go:1*/ byte(i) + /*line JzNm3uvKpzq4.go:1*/ byte(positions[i]^positions[i+1]) + 167
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line JzNm3uvKpzq4.go:1*/ string(data)
			}(), /*line FyZiqY.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.C50bBOZxcYfV)

			var cbz6DUSD types.View

			for _, iv5UGEus9T := range cu1RSg2YR.nbOZ17j {
				/*line N06zcm.go:1*/ cu1RSg2YR.mGpWFGC2.ProcessRemoteTmo(iv5UGEus9T)

				if cbz6DUSD < iv5UGEus9T.RR2FRcA9 {
					cbz6DUSD = iv5UGEus9T.RR2FRcA9
				}
			}

			if /*line aqZbo0Fi.go:1*/ cu1RSg2YR.mGpWFGC2.GetAnchorView() < cbz6DUSD {
				/*line BwzvwC.go:1*/ cu1RSg2YR.mGpWFGC2.UpdateAnchorView(cbz6DUSD)
			}

			/*line QUiutMSxLkp.go:1*/
			cu1RSg2YR.mGpWFGC2.ExecuteViewChange( /*line lcgvOAk2.go:1*/ cu1RSg2YR.mGpWFGC2.GetNewView())

			cu1RSg2YR.nbOZ17j = /*line JvBlUGYDdZ.go:1*/ make(map[types.NodeID]*pacemaker.H8NP1AOKTF)
		}
		return
	}

	if /*line zUdHcMa_c.go:1*/ len(cu1RSg2YR.nbOZ17j) > 0 {
		cu1RSg2YR.nbOZ17j = /*line CaropQ.go:1*/ make(map[types.NodeID]*pacemaker.H8NP1AOKTF)
	}

	if iv5UGEus9T.RR2FRcA9 > /*line DYblk1kk.go:1*/ cu1RSg2YR.mGpWFGC2.GetAnchorView() {
		/*line ITYp77ffI.go:1*/ cu1RSg2YR.mGpWFGC2.UpdateAnchorView(iv5UGEus9T.RR2FRcA9)
		/*line edYEzcNvfgo.go:1*/ cu1RSg2YR.mGpWFGC2.ExecuteViewChange( /*line hXfNxwU.go:1*/ cu1RSg2YR.mGpWFGC2.GetNewView())
		return
	}

	qNDcZn, aiM0S8js, hByvRzR := /*line E8m3MFhvtQmB.go:1*/ cu1RSg2YR.mGpWFGC2.ProcessRemoteTmo(iv5UGEus9T)
	if !qNDcZn {
		return
	}

	aiM0S8js.Epoch = /*line sULQmaug.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch()

	if /*line tk80kTTbkx5c.go:1*/ cu1RSg2YR.IsLeader( /*line pdQGGaIX.go:1*/ cu1RSg2YR.ID(), aiM0S8js.APadJA, aiM0S8js.Epoch) {
		return
	}

	if hByvRzR != nil {
		cu1RSg2YR.jtokCZw_Dtx <- hByvRzR
	}

	/*line h0zOCSv.go:1*/
	log.ViJSfja(func() /*line nOiRvj.go:1*/ string {
		data := /*line nOiRvj.go:1*/ make([]byte, 0, 54)
		i := 4
		decryptKey := 18
		for counter := 0; i != 14; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				data = /*line nOiRvj.go:1*/ append(data, 245)
				i = 18
			case 8:
				data = /*line nOiRvj.go:1*/ append(data, "\xf5\xe3\xa7\xef"...,
				)
				i = 16
			case 4:
				i = 1
				data = /*line nOiRvj.go:1*/ append(data, "\xf6\xbf\x0f"...,
				)
			case 11:
				i = 17
				data = /*line nOiRvj.go:1*/ append(data, "\xdc\xfc\xfd\xb6"...,
				)
			case 17:
				data = /*line nOiRvj.go:1*/ append(data, "\xac\xe4\xa2"...,
				)
				i = 8
			case 15:
				i = 5
				data = /*line nOiRvj.go:1*/ append(data, "$3\xde#"...,
				)
			case 7:
				i = 11
				data = /*line nOiRvj.go:1*/ append(data, "\x01\xfa\xfe\xee"...,
				)
			case 6:
				data = /*line nOiRvj.go:1*/ append(data, "\xa4\x1d/\""...,
				)
				i = 15
			case 5:
				data = /*line nOiRvj.go:1*/ append(data, "+%\xd2\xff"...,
				)
				i = 3
			case 13:
				data = /*line nOiRvj.go:1*/ append(data, "\x02\xf5\xf6\x03"...,
				)
				i = 2
			case 10:
				i = 9
				data = /*line nOiRvj.go:1*/ append(data, 11)
			case 2:
				i = 7
				data = /*line nOiRvj.go:1*/ append(data, "\n\xe8\xfa"...,
				)
			case 3:
				data = /*line nOiRvj.go:1*/ append(data, "\x15.\xd6"...,
				)
				i = 10
			case 9:
				data = /*line nOiRvj.go:1*/ append(data, "\x1d\x10!\xc9"...,
				)
				i = 0
			case 0:
				data = /*line nOiRvj.go:1*/ append(data, "\xcd%"...,
				)
				i = 12
			case 12:
				i = 14
				for y := range data {
					data[y] = data[y] - /*line nOiRvj.go:1*/ byte(decryptKey^y)
				}
			case 16:
				data = /*line nOiRvj.go:1*/ append(data, 248)
				i = 6
			case 18:
				data = /*line nOiRvj.go:1*/ append(data, "\xbf\xc6\xed\x0e"...,
				)
				i = 13
			}
		}
		return /*line nOiRvj.go:1*/ string(data)
	}(), /*line W9YQIfZ4EMwN.go:1*/ cu1RSg2YR.ID(), aiM0S8js.APadJA)

	aiM0S8js.H5NbPTQ = /*line GqoJxcQ.go:1*/ cu1RSg2YR.ID()

	/*line QiJaZDAXVvME.go:1*/
	cu1RSg2YR.SetRole(types.LEADER)

	/*line BqpFvHc.go:1*/
	log.ViJSfja(func() /*line uzk3HA0.go:1*/ string {
		key := [] /*line uzk3HA0.go:1*/ byte("\x1c\xb9㹠n\n\x06K\x12̵?_\\\x89\xb7\xe4\r\x99\xcfb\xe6\x8bq\xc6\xf4\x9f\xa8h\xaf\x10C\xa9,\xd52m\x1a3\x05\vR\x15~\xd7\xf7VuT4'\t\x9e\xe8\x82}\x11+Y\x83\x1c\x919y,P\x9d\xbfȁ`5^\xfb\x89\xb7\x12\x16\xd8\xc8")
		data := [] /*line uzk3HA0.go:1*/ byte("?l\x93\xa4\x80\xbaFl$Q\x99\xbe4\xf3\t丐X\xbb\x9e\rC\x95\xf5\xa3z\xca\xcb\x00\xb6T\xdd\xc7F\x9a1\xf8Y@dc\x15\v\xf4\x8ev\x19\xff\x11\xecM`\xcf}\xed\xf8c\xf5\r\xefS\xdc\xe7\xacJ\xd0ɰ\xaa\x9f\x0e0\x19{\xe0\xaee\nM\xae")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line uzk3HA0.go:1*/ string(data)
	}(), /*line Et7Io52G.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.RKtUCo69O, iv5UGEus9T.C50bBOZxcYfV)

	/*line tySPe5u.go:1*/
	cu1RSg2YR.Broadcast(aiM0S8js)
	/*line NP2yYgmvLxlb.go:1*/ cu1RSg2YR.ProcessTC(aiM0S8js)
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessLocalTmo(a3qJWd types.View) {
	/*line BCsUrrVRE.go:1*/ log.ViJSfja(func() /*line A3WpfWE.go:1*/ string {
		key := [] /*line A3WpfWE.go:1*/ byte("׳Q\xac\x0f\xa5\x1em\x8dEb\xe5S\x9e\xe2\x90rA#\xb2\xe2\xc1\xb9\xa2\x8a\x89\xa9=\xa3\x88\xd2\xd8<S\xe5<\xb5J̺\xbcZ\x86\xda\xd1\x10!/\xa2N\xa4\xb7\x19u\xa8$K\x8d\xd3")
		data := [] /*line A3WpfWE.go:1*/ byte("\x8c\x96'\xf1/\x8dN\x1f\xe2&\a\x96 ҍ\xf3\x13-wߍ\xe8\x99\xd2\xf8\xe6\xcaX\xd0\xfb\xbb\xb6[s\x89S\xd6+\xa0\x9a\xc83뿾eU\x0f\xc4!֗o\x1c\xcdSk\xa8\xa5")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line A3WpfWE.go:1*/ string(data)
	}(), /*line LXEshT0hI1JG.go:1*/ cu1RSg2YR.ID(), a3qJWd)

	/*line JSTuWH23aF.go:1*/
	cu1RSg2YR.SetState(types.VIEWCHANGING)

	iv5UGEus9T := &pacemaker.H8NP1AOKTF{
		Shard:/*line IcdYLZOpMeW.go:1*/ cu1RSg2YR.Shard(),
		Epoch:/*line xsDQr2mKVo.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpoch(),
		W72t5Nhk: a3qJWd,
		C50bBOZxcYfV:/*line x_fxYdjO6t.go:1*/ cu1RSg2YR.mGpWFGC2.GetNewView(),
		RR2FRcA9:/*line PkbwmtZEE7uX.go:1*/ cu1RSg2YR.mGpWFGC2.GetAnchorView(),
		RKtUCo69O:/*line Mcb79Fv0UaEa.go:1*/ cu1RSg2YR.ID(),
		AWApIIs:/*line GkVToz.go:1*/ cu1RSg2YR.GetHighBlockHeight(),
		A2Q6Z0L:/*line A_ULi7G.go:1*/ cu1RSg2YR.GetHighBlockHeight(),
		QSw1K7B:/*line _qRE6TF4UOEq.go:1*/ cu1RSg2YR.iSgBUFf9.GetBlockByBlockHeight( /*line EP_nDbAS.go:1*/ cu1RSg2YR.GetHighBlockHeight()).(*blockchain.CoordinationBlock),
		Jth9oL0xv: nil,
		DSNq8K7x:/*line aacNYojbb.go:1*/ cu1RSg2YR.GetHighQC(),
	}

	if ouJpb2aF3mlv, kbgkwuEjk := cu1RSg2YR.cA3KSEC5m[iv5UGEus9T.Epoch][iv5UGEus9T.W72t5Nhk][iv5UGEus9T.AWApIIs+1]; kbgkwuEjk {
		iv5UGEus9T.A2Q6Z0L += 1
		iv5UGEus9T.Jth9oL0xv = ouJpb2aF3mlv
	}

	/*line P4PsLX.go:1*/
	log.ViJSfja(func() /*line jKdXJ1.go:1*/ string {
		fullData := [] /*line jKdXJ1.go:1*/ byte("\xb2BQ\x842t\xcet-Ā\xd5st'\xa2\xa7\x17\x0f\x04\xb5G,\xfb8.B&A\x1eՁ\xf4S\x9fvA\xed\bkg\xb0\xbc\x92\x1f\xd9\x03#\xb2<\x1f\xe3\xdc~s\xb8\r\xa6w\xb8s㟍\x0e\xf5\xb7X\xe3Y\xe3Ul\x9cy\xa5\x1f2\t\x98=\xbdl(Z|\xac\x02K㮦/?\xfd\x93S\x1e\xec\xd5\x1dXB\xa4h\x92a\x0f\x1b\bz\t#3\xf0\xf5{\x15\xc6y\xea\xe8\xac\xdc\xf2\xaaҬ\xf6ߊY7C\xe0h㔋\xd7\xff<d\x1a\xd7\x11\xcb>\x86\xcb#\xc4\xc5TP\x17VZ\xde\xe2\xa1Fa'\x9d\x8b#\x03t\xb5\x9aߍ\x1e8t\xa4?\x1bx\xfa里\a\x7f\x91{>\xbar\xaa7n\xfc\xd8\x7f\\\xdf\xf6\x9a\f\xafN\xd7Bz\x88\x01\xc1ӷg\x99\x98\xe7$\xb4-\xac\x15.t\xf3\xdb\x1c\x89\x9eI=e\xc7\xdet\xe4\xc6")
		idxKey := [] /*line jKdXJ1.go:1*/ byte("\x14\xf1&\xf7\xa1X\x0e \x9b\x14\xc9Ɵ")
		data := /*line jKdXJ1.go:1*/ make([]byte, 0, 119)
		data = /*line jKdXJ1.go:1*/ append(data, fullData[249^ /*line jKdXJ1.go:1*/ int(idxKey[2])]-fullData[240^ /*line jKdXJ1.go:1*/ int(idxKey[2])], fullData[229^ /*line jKdXJ1.go:1*/ int(idxKey[11])]+fullData[12^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[106^ /*line jKdXJ1.go:1*/ int(idxKey[0])]^fullData[164^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[35^ /*line jKdXJ1.go:1*/ int(idxKey[3])]+fullData[48^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[192^ /*line jKdXJ1.go:1*/ int(idxKey[10])]-fullData[174^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[191^ /*line jKdXJ1.go:1*/ int(idxKey[5])]^fullData[237^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[98^ /*line jKdXJ1.go:1*/ int(idxKey[4])]+fullData[18^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[47^ /*line jKdXJ1.go:1*/ int(idxKey[10])]^fullData[82^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[71^ /*line jKdXJ1.go:1*/ int(idxKey[12])]-fullData[54^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[133^ /*line jKdXJ1.go:1*/ int(idxKey[3])]^fullData[168^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[204^ /*line jKdXJ1.go:1*/ int(idxKey[12])]+fullData[207^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[77^ /*line jKdXJ1.go:1*/ int(idxKey[7])]+fullData[7^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[239^ /*line jKdXJ1.go:1*/ int(idxKey[1])]^fullData[170^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[208^ /*line jKdXJ1.go:1*/ int(idxKey[8])]+fullData[45^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[55^ /*line jKdXJ1.go:1*/ int(idxKey[7])]+fullData[45^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[8^ /*line jKdXJ1.go:1*/ int(idxKey[7])]+fullData[226^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[161^ /*line jKdXJ1.go:1*/ int(idxKey[12])]-fullData[12^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[130^ /*line jKdXJ1.go:1*/ int(idxKey[3])]^fullData[189^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[178^ /*line jKdXJ1.go:1*/ int(idxKey[0])]^fullData[46^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[126^ /*line jKdXJ1.go:1*/ int(idxKey[8])]^fullData[1^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[173^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[178^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[142^ /*line jKdXJ1.go:1*/ int(idxKey[10])]^fullData[156^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[199^ /*line jKdXJ1.go:1*/ int(idxKey[4])]+fullData[73^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[26^ /*line jKdXJ1.go:1*/ int(idxKey[11])]^fullData[172^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[68^ /*line jKdXJ1.go:1*/ int(idxKey[11])]+fullData[128^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[44^ /*line jKdXJ1.go:1*/ int(idxKey[1])]^fullData[213^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[40^ /*line jKdXJ1.go:1*/ int(idxKey[10])]^fullData[68^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[214^ /*line jKdXJ1.go:1*/ int(idxKey[12])]^fullData[222^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[161^ /*line jKdXJ1.go:1*/ int(idxKey[7])]^fullData[90^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[211^ /*line jKdXJ1.go:1*/ int(idxKey[1])]-fullData[53^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[0^ /*line jKdXJ1.go:1*/ int(idxKey[9])]^fullData[132^ /*line jKdXJ1.go:1*/ int(idxKey[9])], fullData[26^ /*line jKdXJ1.go:1*/ int(idxKey[5])]-fullData[89^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[196^ /*line jKdXJ1.go:1*/ int(idxKey[4])]+fullData[48^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[208^ /*line jKdXJ1.go:1*/ int(idxKey[12])]-fullData[137^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[125^ /*line jKdXJ1.go:1*/ int(idxKey[10])]+fullData[167^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[28^ /*line jKdXJ1.go:1*/ int(idxKey[11])]-fullData[254^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[62^ /*line jKdXJ1.go:1*/ int(idxKey[3])]-fullData[206^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[43^ /*line jKdXJ1.go:1*/ int(idxKey[0])]-fullData[9^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[13^ /*line jKdXJ1.go:1*/ int(idxKey[7])]+fullData[245^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[52^ /*line jKdXJ1.go:1*/ int(idxKey[1])]+fullData[102^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[181^ /*line jKdXJ1.go:1*/ int(idxKey[10])]^fullData[202^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[112^ /*line jKdXJ1.go:1*/ int(idxKey[0])]^fullData[187^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[251^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[69^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[42^ /*line jKdXJ1.go:1*/ int(idxKey[1])]^fullData[17^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[123^ /*line jKdXJ1.go:1*/ int(idxKey[0])]+fullData[5^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[245^ /*line jKdXJ1.go:1*/ int(idxKey[2])]-fullData[79^ /*line jKdXJ1.go:1*/ int(idxKey[2])], fullData[176^ /*line jKdXJ1.go:1*/ int(idxKey[10])]-fullData[119^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[89^ /*line jKdXJ1.go:1*/ int(idxKey[2])]+fullData[142^ /*line jKdXJ1.go:1*/ int(idxKey[2])], fullData[67^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[19^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[181^ /*line jKdXJ1.go:1*/ int(idxKey[9])]^fullData[100^ /*line jKdXJ1.go:1*/ int(idxKey[9])], fullData[134^ /*line jKdXJ1.go:1*/ int(idxKey[0])]^fullData[62^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[180^ /*line jKdXJ1.go:1*/ int(idxKey[3])]-fullData[239^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[22^ /*line jKdXJ1.go:1*/ int(idxKey[9])]^fullData[183^ /*line jKdXJ1.go:1*/ int(idxKey[9])], fullData[229^ /*line jKdXJ1.go:1*/ int(idxKey[3])]+fullData[106^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[15^ /*line jKdXJ1.go:1*/ int(idxKey[5])]-fullData[252^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[248^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[243^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[184^ /*line jKdXJ1.go:1*/ int(idxKey[2])]+fullData[60^ /*line jKdXJ1.go:1*/ int(idxKey[2])], fullData[246^ /*line jKdXJ1.go:1*/ int(idxKey[11])]+fullData[250^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[151^ /*line jKdXJ1.go:1*/ int(idxKey[11])]^fullData[83^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[163^ /*line jKdXJ1.go:1*/ int(idxKey[6])]^fullData[178^ /*line jKdXJ1.go:1*/ int(idxKey[6])], fullData[61^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[210^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[20^ /*line jKdXJ1.go:1*/ int(idxKey[8])]^fullData[114^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[159^ /*line jKdXJ1.go:1*/ int(idxKey[0])]-fullData[202^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[194^ /*line jKdXJ1.go:1*/ int(idxKey[11])]^fullData[146^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[205^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[247^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[70^ /*line jKdXJ1.go:1*/ int(idxKey[1])]-fullData[117^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[80^ /*line jKdXJ1.go:1*/ int(idxKey[3])]^fullData[216^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[198^ /*line jKdXJ1.go:1*/ int(idxKey[6])]^fullData[108^ /*line jKdXJ1.go:1*/ int(idxKey[6])], fullData[184^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[57^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[84^ /*line jKdXJ1.go:1*/ int(idxKey[7])]^fullData[65^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[33^ /*line jKdXJ1.go:1*/ int(idxKey[2])]-fullData[120^ /*line jKdXJ1.go:1*/ int(idxKey[2])], fullData[163^ /*line jKdXJ1.go:1*/ int(idxKey[7])]^fullData[87^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[91^ /*line jKdXJ1.go:1*/ int(idxKey[1])]+fullData[123^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[137^ /*line jKdXJ1.go:1*/ int(idxKey[1])]-fullData[244^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[120^ /*line jKdXJ1.go:1*/ int(idxKey[10])]-fullData[251^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[40^ /*line jKdXJ1.go:1*/ int(idxKey[1])]+fullData[198^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[76^ /*line jKdXJ1.go:1*/ int(idxKey[1])]-fullData[180^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[170^ /*line jKdXJ1.go:1*/ int(idxKey[4])]+fullData[66^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[215^ /*line jKdXJ1.go:1*/ int(idxKey[3])]+fullData[253^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[212^ /*line jKdXJ1.go:1*/ int(idxKey[5])]^fullData[243^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[234^ /*line jKdXJ1.go:1*/ int(idxKey[8])]+fullData[199^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[114^ /*line jKdXJ1.go:1*/ int(idxKey[3])]-fullData[59^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[93^ /*line jKdXJ1.go:1*/ int(idxKey[8])]^fullData[178^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[9^ /*line jKdXJ1.go:1*/ int(idxKey[10])]^fullData[80^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[74^ /*line jKdXJ1.go:1*/ int(idxKey[6])]^fullData[193^ /*line jKdXJ1.go:1*/ int(idxKey[6])], fullData[132^ /*line jKdXJ1.go:1*/ int(idxKey[4])]+fullData[26^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[222^ /*line jKdXJ1.go:1*/ int(idxKey[6])]+fullData[128^ /*line jKdXJ1.go:1*/ int(idxKey[6])], fullData[35^ /*line jKdXJ1.go:1*/ int(idxKey[5])]-fullData[110^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[102^ /*line jKdXJ1.go:1*/ int(idxKey[11])]+fullData[176^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[38^ /*line jKdXJ1.go:1*/ int(idxKey[1])]-fullData[72^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[255^ /*line jKdXJ1.go:1*/ int(idxKey[1])]^fullData[208^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[198^ /*line jKdXJ1.go:1*/ int(idxKey[8])]^fullData[183^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[191^ /*line jKdXJ1.go:1*/ int(idxKey[1])]-fullData[27^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[35^ /*line jKdXJ1.go:1*/ int(idxKey[8])]-fullData[33^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[97^ /*line jKdXJ1.go:1*/ int(idxKey[3])]^fullData[217^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[155^ /*line jKdXJ1.go:1*/ int(idxKey[8])]-fullData[86^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[147^ /*line jKdXJ1.go:1*/ int(idxKey[5])]^fullData[105^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[34^ /*line jKdXJ1.go:1*/ int(idxKey[10])]-fullData[107^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[184^ /*line jKdXJ1.go:1*/ int(idxKey[0])]+fullData[39^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[77^ /*line jKdXJ1.go:1*/ int(idxKey[5])]^fullData[67^ /*line jKdXJ1.go:1*/ int(idxKey[5])], fullData[166^ /*line jKdXJ1.go:1*/ int(idxKey[7])]+fullData[11^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[88^ /*line jKdXJ1.go:1*/ int(idxKey[0])]^fullData[218^ /*line jKdXJ1.go:1*/ int(idxKey[0])], fullData[8^ /*line jKdXJ1.go:1*/ int(idxKey[6])]^fullData[177^ /*line jKdXJ1.go:1*/ int(idxKey[6])], fullData[154^ /*line jKdXJ1.go:1*/ int(idxKey[4])]+fullData[201^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[40^ /*line jKdXJ1.go:1*/ int(idxKey[4])]-fullData[236^ /*line jKdXJ1.go:1*/ int(idxKey[4])], fullData[237^ /*line jKdXJ1.go:1*/ int(idxKey[1])]^fullData[249^ /*line jKdXJ1.go:1*/ int(idxKey[1])], fullData[158^ /*line jKdXJ1.go:1*/ int(idxKey[11])]-fullData[242^ /*line jKdXJ1.go:1*/ int(idxKey[11])], fullData[11^ /*line jKdXJ1.go:1*/ int(idxKey[9])]+fullData[139^ /*line jKdXJ1.go:1*/ int(idxKey[9])], fullData[169^ /*line jKdXJ1.go:1*/ int(idxKey[10])]^fullData[103^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[217^ /*line jKdXJ1.go:1*/ int(idxKey[10])]+fullData[24^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[162^ /*line jKdXJ1.go:1*/ int(idxKey[10])]-fullData[180^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[168^ /*line jKdXJ1.go:1*/ int(idxKey[7])]+fullData[180^ /*line jKdXJ1.go:1*/ int(idxKey[7])], fullData[170^ /*line jKdXJ1.go:1*/ int(idxKey[10])]-fullData[8^ /*line jKdXJ1.go:1*/ int(idxKey[10])], fullData[202^ /*line jKdXJ1.go:1*/ int(idxKey[3])]^fullData[82^ /*line jKdXJ1.go:1*/ int(idxKey[3])], fullData[185^ /*line jKdXJ1.go:1*/ int(idxKey[12])]+fullData[215^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[170^ /*line jKdXJ1.go:1*/ int(idxKey[12])]+fullData[144^ /*line jKdXJ1.go:1*/ int(idxKey[12])], fullData[73^ /*line jKdXJ1.go:1*/ int(idxKey[8])]^fullData[27^ /*line jKdXJ1.go:1*/ int(idxKey[8])], fullData[134^ /*line jKdXJ1.go:1*/ int(idxKey[11])]+fullData[65^ /*line jKdXJ1.go:1*/ int(idxKey[11])])
		return /*line jKdXJ1.go:1*/ string(data)
	}(), /*line bBU29ej.go:1*/ cu1RSg2YR.ID(), iv5UGEus9T.W72t5Nhk, iv5UGEus9T.C50bBOZxcYfV, iv5UGEus9T.RR2FRcA9, iv5UGEus9T.AWApIIs, iv5UGEus9T.A2Q6Z0L)

	/*line iTa0zohN0O.go:1*/
	cu1RSg2YR.BroadcastToSome( /*line GOawekYO.go:1*/ cu1RSg2YR.FindCommitteesFor(iv5UGEus9T.Epoch), iv5UGEus9T)
	/*line XZhfYTF5lC.go:1*/ cu1RSg2YR.ProcessRemoteTmo(iv5UGEus9T)

	/*line J8l9MY.go:1*/
	log.ViJSfja(func() /*line M5vimUa2.go:1*/ string {
		fullData := [] /*line M5vimUa2.go:1*/ byte("^\x863\xa5\xba\x82\x88Ⱦ\xcf\x05f\x13D>\xd2\xd9$4\x83|\x9ebM\x06\xb7\v\xc3P\xe7h\xd4I\xc3ī\x8b\xe0\a@#\xd9\xecż\x9bi\x1b\x8d\x17OUv\xf9+z֕\xe9\xb5\xdb|\x1f,S\x9c%\xb0\a\xa5\xfd\xaf\xe2a\xb9\x99\xf6\xb67g\xe8\xf3\xb0̸\xe7\xcc\xd1gK\xfeK\x8c0\xd8\x1b\xa3\xa1\xdcp(\x9f\xa9\t\xc4u繄v\x99z\x82\vZ\xeax5\xb0\xa8\xeb\x9e\x10\xbd&\xc1\xa5\x8c8\xd4:@K\xc3\xea\xd7")
		idxKey := [] /*line M5vimUa2.go:1*/ byte("\x16\x8dV")
		data := /*line M5vimUa2.go:1*/ make([]byte, 0, 69)
		data = /*line M5vimUa2.go:1*/ append(data, fullData[29^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[12^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[160^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[224^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[43^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[79^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[47^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[211^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[45^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[47^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[14^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[100^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[57^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[79^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[129^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[196^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[189^ /*line M5vimUa2.go:1*/ int(idxKey[1])]+fullData[197^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[87^ /*line M5vimUa2.go:1*/ int(idxKey[2])]-fullData[126^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[203^ /*line M5vimUa2.go:1*/ int(idxKey[1])]+fullData[147^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[119^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[4^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[105^ /*line M5vimUa2.go:1*/ int(idxKey[2])]-fullData[28^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[22^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[104^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[187^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[161^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[18^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[6^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[229^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[200^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[135^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[163^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[19^ /*line M5vimUa2.go:1*/ int(idxKey[0])]+fullData[25^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[225^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[188^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[11^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[16^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[142^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[153^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[117^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[10^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[168^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[186^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[251^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[164^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[158^ /*line M5vimUa2.go:1*/ int(idxKey[1])]+fullData[245^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[215^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[45^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[34^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[39^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[53^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[13^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[177^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[185^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[81^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[77^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[120^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[125^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[234^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[198^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[127^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[48^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[212^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[101^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[50^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[70^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[207^ /*line M5vimUa2.go:1*/ int(idxKey[1])]+fullData[170^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[146^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[72^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[231^ /*line M5vimUa2.go:1*/ int(idxKey[1])]+fullData[242^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[64^ /*line M5vimUa2.go:1*/ int(idxKey[2])]-fullData[99^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[138^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[255^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[88^ /*line M5vimUa2.go:1*/ int(idxKey[2])]-fullData[209^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[167^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[219^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[38^ /*line M5vimUa2.go:1*/ int(idxKey[2])]+fullData[37^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[88^ /*line M5vimUa2.go:1*/ int(idxKey[0])]+fullData[150^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[73^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[66^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[94^ /*line M5vimUa2.go:1*/ int(idxKey[2])]+fullData[54^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[21^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[52^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[90^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[46^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[69^ /*line M5vimUa2.go:1*/ int(idxKey[0])]+fullData[97^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[156^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[154^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[7^ /*line M5vimUa2.go:1*/ int(idxKey[2])]+fullData[57^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[194^ /*line M5vimUa2.go:1*/ int(idxKey[1])]+fullData[215^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[115^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[75^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[118^ /*line M5vimUa2.go:1*/ int(idxKey[2])]-fullData[73^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[68^ /*line M5vimUa2.go:1*/ int(idxKey[2])]^fullData[213^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[201^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[216^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[31^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[112^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[247^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[236^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[175^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[192^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[106^ /*line M5vimUa2.go:1*/ int(idxKey[0])]-fullData[14^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[141^ /*line M5vimUa2.go:1*/ int(idxKey[1])]^fullData[233^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[152^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[248^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[44^ /*line M5vimUa2.go:1*/ int(idxKey[0])]^fullData[74^ /*line M5vimUa2.go:1*/ int(idxKey[0])], fullData[84^ /*line M5vimUa2.go:1*/ int(idxKey[2])]+fullData[91^ /*line M5vimUa2.go:1*/ int(idxKey[2])], fullData[204^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[176^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[11^ /*line M5vimUa2.go:1*/ int(idxKey[1])]-fullData[166^ /*line M5vimUa2.go:1*/ int(idxKey[1])], fullData[65^ /*line M5vimUa2.go:1*/ int(idxKey[0])]+fullData[104^ /*line M5vimUa2.go:1*/ int(idxKey[0])])
		return /*line M5vimUa2.go:1*/ string(data)
	}(), /*line IVbkB_Cd.go:1*/ cu1RSg2YR.ID(), a3qJWd)
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessTC(aiM0S8js *pacemaker.RZ65hTAdXj) {
	/*line MNX_hzj.go:1*/ log.ViJSfja(func() /*line Ce1CAcNNZ0x.go:1*/ string {
		fullData := [] /*line Ce1CAcNNZ0x.go:1*/ byte("\xa7\v\x02Dݑ\x1e\xbe\xe4\xfe\xe1z\xbd\xeea\x14}Sp\xfa\xf7\xb0\xd5k\xf2\"{\x17c=L\x15\x952\t\xb1y\xb8\xb4\x11i\xfc\xedl\xc9\xff|D\xa8\x11\xcexZx\xa1\x94\xefw\xa1\xdb\xd9\xee6\x90\xc7W5\x91[\x15\xc5;\xc7\xdb\xc2j\xf7K\xb5\xbf\xbe`\x1d\xb12@N\xd2G\xb4\xd2\xc1\x88\xbc\xccs\xaaQ\x9de\x93-&\x88a\xe3\x9bC\b\xb5W>\x9d\x0fuz\xa5\xbd\x8a\xed#\xccr\t!\xa4\xeb\xec\xf9\xb9")
		idxKey := [] /*line Ce1CAcNNZ0x.go:1*/ byte("\xb7tK\xfaJ7}\xd9f\x91\ft\x9e`Md")
		data := /*line Ce1CAcNNZ0x.go:1*/ make([]byte, 0, 66)
		data = /*line Ce1CAcNNZ0x.go:1*/ append(data, fullData[80^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])]^fullData[117^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])], fullData[119^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]-fullData[18^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[68^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[4])]^fullData[81^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[4])], fullData[164^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]^fullData[89^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[60^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])]+fullData[57^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])], fullData[191^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[0])]+fullData[180^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[0])], fullData[86^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]-fullData[97^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[10^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]-fullData[97^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[79^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[13])]-fullData[118^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[13])], fullData[131^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[3])]-fullData[210^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[3])], fullData[51^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]+fullData[62^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[115^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])]-fullData[40^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])], fullData[87^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])]-fullData[27^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])], fullData[41^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]^fullData[5^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[75^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])]-fullData[59^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])], fullData[19^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])]+fullData[78^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])], fullData[216^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]+fullData[156^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[18^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])]-fullData[36^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])], fullData[51^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]^fullData[42^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[115^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])]-fullData[77^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])], fullData[205^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]+fullData[242^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[168^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]-fullData[185^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[96^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])]-fullData[47^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])], fullData[6^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])]+fullData[125^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])], fullData[247^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]^fullData[198^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[253^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[12])]+fullData[188^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[12])], fullData[58^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]+fullData[70^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[34^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])]+fullData[46^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])], fullData[248^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]-fullData[222^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[20^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])]+fullData[58^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])], fullData[107^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]-fullData[21^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[18^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]+fullData[85^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[148^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[9])]-fullData[136^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[9])], fullData[121^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])]+fullData[41^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])], fullData[120^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])]+fullData[67^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])], fullData[121^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])]+fullData[32^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])], fullData[33^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[6])]+fullData[52^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[6])], fullData[194^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[3])]-fullData[140^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[3])], fullData[127^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[6])]+fullData[111^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[6])], fullData[84^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])]-fullData[82^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])], fullData[74^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[4])]+fullData[0^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[4])], fullData[63^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])]^fullData[10^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])], fullData[247^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[9])]+fullData[250^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[9])], fullData[8^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])]+fullData[32^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])], fullData[9^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[5])]^fullData[118^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[5])], fullData[119^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])]^fullData[28^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])], fullData[207^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[12])]-fullData[210^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[12])], fullData[123^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])]^fullData[110^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[1])], fullData[40^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[6])]-fullData[42^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[6])], fullData[4^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])]^fullData[106^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[8])], fullData[46^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]^fullData[6^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[25^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]^fullData[49^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[13^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]-fullData[90^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[194^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[9])]-fullData[210^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[9])], fullData[182^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[5])]+fullData[90^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[5])], fullData[95^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])]^fullData[52^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[15])], fullData[36^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])]^fullData[122^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[14])], fullData[35^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])]+fullData[4^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[2])], fullData[54^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[4])]^fullData[36^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[4])], fullData[151^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])]-fullData[199^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[7])], fullData[138^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[0])]+fullData[142^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[0])], fullData[97^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])]^fullData[52^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])], fullData[126^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])]-fullData[47^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[11])], fullData[51^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])]+fullData[44^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[10])], fullData[134^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[0])]-fullData[221^ /*line Ce1CAcNNZ0x.go:1*/ int(idxKey[0])])
		return /*line Ce1CAcNNZ0x.go:1*/ string(data)
	}(), /*line iRWhNUus.go:1*/ cu1RSg2YR.ID(), aiM0S8js.APadJA)

	/*line zc3TvO.go:1*/
	cu1RSg2YR.SetState(types.READY)

	hYhbNJBbO, uxNBF8Z7 := /*line EPHu_9.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpochView()

	/*line ZDAfRtR.go:1*/
	delete(cu1RSg2YR.cA3KSEC5m[hYhbNJBbO][uxNBF8Z7], aiM0S8js.BecWOQ4Lltj)

	/*line l2I0gSTZvY0.go:1*/
	cu1RSg2YR.mGpWFGC2.UpdateView(aiM0S8js.APadJA - 1)
	/*line yrT9of1ayB6Y.go:1*/ cu1RSg2YR.mGpWFGC2.UpdateAnchorView(aiM0S8js.C8IzV0maUK)

	/*line rsns6pn.go:1*/
	log.ViJSfja(func() /*line __ZjE6.go:1*/ string {
		key := [] /*line __ZjE6.go:1*/ byte("#S\xa5\rK>w'\xc6\x1av\x16iNqZ\x19\x97\xf4\xcdƒx\\D\x7f F>\xbd\x17\xbc\x82\u058bf\x90\x8c\x83\f\x95\xa7-z\xe0\xfc\xad \x97\ve\xec\x16Z\xec\x83 !\x8a\x1f\xba,\x1c\r6\xa8\x05\xa9\x8fM}B?\xe9")
		data := [] /*line __ZjE6.go:1*/ byte("8\xd2\xd1P\xd5\xea\xd9K\xa9I\xef]\n\x06\xd2\xcf\a\xcfu\xa1\xa3\xe1\xf0\t \xa1P,1\xa6N\xb7\xf1\x93\xe3\x01\x90\xe8\xe6a\xd0\xc8H\xfa@g\xb8R\xdd^\x01}M\a\x88\xe6OM\x96G\xb5F\x04a/\xcf\x1b\xcd\xda\x18\xfa\xde\xe6\x8d")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line __ZjE6.go:1*/ string(data)
	}(), /*line D1zShTNJMXr.go:1*/ cu1RSg2YR.ID(), aiM0S8js.APadJA)

	hYhbNJBbO, uxNBF8Z7 = /*line Dqlo5xab.go:1*/ cu1RSg2YR.mGpWFGC2.GetCurEpochView()

	if ouJpb2aF3mlv, kbgkwuEjk := cu1RSg2YR.wdUdVrSEU[hYhbNJBbO][uxNBF8Z7][aiM0S8js.BecWOQ4Lltj]; kbgkwuEjk {
		_ = /*line TsmsW8wcRqRj.go:1*/ cu1RSg2YR.ProcessCoordinationBlock(ouJpb2aF3mlv)
		/*line szzDMESV6L.go:1*/ delete(cu1RSg2YR.wdUdVrSEU[hYhbNJBbO][uxNBF8Z7], aiM0S8js.BecWOQ4Lltj)
	}

}

func (cu1RSg2YR *L4qxkgqr2rNN) GetHighQC() *quorum.HPOWa1PT0xzq {
	/*line s5PxyiGd.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line GVlcvEHXVk.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	return cu1RSg2YR.bQsSnXJ
}

func (cu1RSg2YR *L4qxkgqr2rNN) GetHighCQC() *quorum.HPOWa1PT0xzq {
	/*line jPC0uiiq.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line B561ab.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	return cu1RSg2YR.cj2GSMP
}

func (cu1RSg2YR *L4qxkgqr2rNN) GetChainStatus() string {
	vO8QgiLWR_a := /*line hMmxNiwb2a.go:1*/ cu1RSg2YR.iSgBUFf9.GetChainGrowth()
	luidf2 := /*line Jw1twLssU.go:1*/ cu1RSg2YR.iSgBUFf9.GetBlockIntervals()
	return /*line JJlBp8.go:1*/ fmt.Sprintf(func() /*line ZfhCgQup.go:1*/ string {
		data := [] /*line ZfhCgQup.go:1*/ byte("\x05\xf4vZ\x94\x1fb,\x97*h\xab\x9c\x82\x19Dpqk\x93\\ i8:\xb5\x98h+\x88c\xa1diT \xb6r\xb0\xae\x87\\\xf0ɩte_\xf6s\xadi\xf3\xb8s\x7f\xc7\xec7\x03\xc0\xd9\xe3c\xa5 Kχb\xb0\xdfa\xab\xde\xe9\x9c]lo\xf4")
		positions := [...]byte{16, 36, 54, 14, 57, 52, 28, 78, 40, 75, 58, 52, 7, 6, 59, 25, 50, 41, 13, 68, 38, 27, 51, 5, 73, 61, 15, 50, 31, 70, 39, 5, 73, 14, 17, 10, 34, 58, 66, 16, 59, 79, 0, 8, 57, 67, 77, 67, 53, 69, 48, 38, 43, 71, 9, 28, 14, 17, 36, 47, 6, 18, 7, 3, 26, 53, 77, 47, 32, 52, 41, 55, 12, 62, 13, 10, 5, 76, 52, 40, 13, 31, 23, 4, 66, 0, 11, 19, 74, 43, 70, 1, 38, 80, 29, 79, 23, 19, 47, 40, 52, 20, 55, 13, 42, 64, 44, 3, 26, 70, 56, 12, 27, 60}
		for i := 0; i < 114; i += 2 {
			localKey := /*line ZfhCgQup.go:1*/ byte(i) + /*line ZfhCgQup.go:1*/ byte(positions[i]^positions[i+1]) + 179
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line ZfhCgQup.go:1*/ string(data)
	}(), /*line T6Wr2sH.go:1*/ cu1RSg2YR.ID() /*line sqSztpj.go:1*/, cu1RSg2YR.mGpWFGC2.GetCurView(), vO8QgiLWR_a, luidf2)
}

func (cu1RSg2YR *L4qxkgqr2rNN) GetHighBlockHeight() types.BlockHeight {
	/*line Nvw2NVYspBy.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line H0LsFuawki5.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	return cu1RSg2YR.bQsSnXJ.BlockHeight
}

func (cu1RSg2YR *L4qxkgqr2rNN) GetLastBlockHeight() types.BlockHeight {
	/*line G5SfVXCtzYYw.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line p6aIycT2PHwZ.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	return cu1RSg2YR.cj2GSMP.BlockHeight
}

func (cu1RSg2YR *L4qxkgqr2rNN) GetLastVotedBlockHeight() types.BlockHeight {
	/*line ANNAPdCPW.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line FufXusTa.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	return cu1RSg2YR._BUwIvrsOe
}

func (cu1RSg2YR *L4qxkgqr2rNN) aDTslULytPz_(cWdSpmiBLOz *quorum.HPOWa1PT0xzq) {
	/*line jghJyXIpsql.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line DE3tOrNRA.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	if cWdSpmiBLOz.BlockHeight > cu1RSg2YR.bQsSnXJ.BlockHeight {
		cu1RSg2YR.bQsSnXJ = cWdSpmiBLOz
		if /*line VNJVgP.go:1*/ cu1RSg2YR.IsLeader( /*line F182fAP.go:1*/ cu1RSg2YR.ID() /*line m5_mdto7.go:1*/, cu1RSg2YR.mGpWFGC2.GetCurView() /*line dcyvFZ.go:1*/, cu1RSg2YR.mGpWFGC2.GetCurEpoch()) {
			cu1RSg2YR.acQiWvFYDH <- cWdSpmiBLOz
		}
	}
}

func (cu1RSg2YR *L4qxkgqr2rNN) z6VaPTrjpK(dWmn_Gtk1d *quorum.HPOWa1PT0xzq) {
	/*line NKOP7gnUy.go:1*/ cu1RSg2YR.tRK_bDOT.Lock()
	defer /*line txOEs2uHvgBs.go:1*/ cu1RSg2YR.tRK_bDOT.Unlock()
	if dWmn_Gtk1d.BlockHeight > cu1RSg2YR.cj2GSMP.BlockHeight {
		cu1RSg2YR.cj2GSMP = dWmn_Gtk1d
	}
}

func (cu1RSg2YR *L4qxkgqr2rNN) xAnHtUT9ATI(bQ_GL8MJ_E types.BlockHeight) error {
	if bQ_GL8MJ_E < cu1RSg2YR._BUwIvrsOe {
		return /*line Br4qLv.go:1*/ errors.New(func() /*line owaGd7.go:1*/ string {
			key := [] /*line owaGd7.go:1*/ byte("\xfa1\x05s\xcck\xf3\x93\v\x84\xbc_\xcdC\xd3\xf8\x9fBDA-\xfa\xa7ЀA\x8c\x89\xad1\xa3\x1b\b\xb1\xc2ѐ\xc0&\xfc\x0f\x7f\x05\x816\xaa\xd1;\x01}fO\xb7H{`\x155\x0f")
			data := [] /*line owaGd7.go:1*/ byte("n\x92w\xda1\xdf\x13\xf5w\xf3\x1f\xca\x15\xa8<_\a\xb6d\xaa\xa0\x1a\x13?\xf7\xa6\xfe\xa9!\x99\x04\x89(%*6\xb0,\x87o\x83\x9f{\xf0\xaa\x0f5[c\xe9ղ\"\x90\xe0\xc9|\x9d\x83")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line owaGd7.go:1*/ string(data)
		}())
	}
	cu1RSg2YR._BUwIvrsOe = bQ_GL8MJ_E
	return nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) wMIZeBk(cWdSpmiBLOz *quorum.HPOWa1PT0xzq) error {
	if cWdSpmiBLOz.BlockHeight <= 1 {
		return nil
	}
	_, k1qj7vimKIy7 := /*line ulmMXUzjYDK.go:1*/ cu1RSg2YR.iSgBUFf9.GetBlockByID(cWdSpmiBLOz.ZWsU_2)
	if k1qj7vimKIy7 != nil {
		return /*line Nuj2sxsKa.go:1*/ fmt.Errorf(func() /*line klOsffQqj.go:1*/ string {
			data := /*line klOsffQqj.go:1*/ make([]byte, 0, 40)
			i := 5
			decryptKey := 16
			for counter := 0; i != 15; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data = /*line klOsffQqj.go:1*/ append(data, "\xdc߰\xb0"...,
					)
					i = 11
				case 1:
					i = 10
					data = /*line klOsffQqj.go:1*/ append(data, "\xd8\xdc\xe0\x8f"...,
					)
				case 7:
					i = 13
					data = /*line klOsffQqj.go:1*/ append(data, "\xd7\xd8"...,
					)
				case 9:
					data = /*line klOsffQqj.go:1*/ append(data, "\xd6ʄ"...,
					)
					i = 7
				case 13:
					i = 6
					data = /*line klOsffQqj.go:1*/ append(data, "\xde\xde\xe0"...,
					)
				case 14:
					i = 15
					for y := range data {
						data[y] = data[y] - /*line klOsffQqj.go:1*/ byte(decryptKey^y)
					}
				case 5:
					data = /*line klOsffQqj.go:1*/ append(data, "\xcc\xc9\xd9"...,
					)
					i = 1
				case 4:
					data = /*line klOsffQqj.go:1*/ append(data, "\xe2\xd5\xe0"...,
					)
					i = 8
				case 6:
					i = 16
					data = /*line klOsffQqj.go:1*/ append(data, "\xec\xef\xe1"...,
					)
				case 0:
					data = /*line klOsffQqj.go:1*/ append(data, "\x84mq"...,
					)
					i = 3
				case 8:
					data = /*line klOsffQqj.go:1*/ append(data, 188)
					i = 2
				case 10:
					i = 9
					data = /*line klOsffQqj.go:1*/ append(data, "\xe3\xd1\xc4\xc4"...,
					)
				case 12:
					i = 4
					data = /*line klOsffQqj.go:1*/ append(data, "\xd3\xdc"...,
					)
				case 16:
					i = 12
					data = /*line klOsffQqj.go:1*/ append(data, "\xe3\x9e"...,
					)
				case 3:
					i = 14
					data = /*line klOsffQqj.go:1*/ append(data, 197)
				case 11:
					i = 0
					data = /*line klOsffQqj.go:1*/ append(data, 191)
				}
			}
			return /*line klOsffQqj.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	v61eq4CI, k1qj7vimKIy7 := /*line g8pmHL5.go:1*/ cu1RSg2YR.iSgBUFf9.GetParentCoordinationBlock(cWdSpmiBLOz.ZWsU_2)
	dWIoYzD := v61eq4CI.(*blockchain.CoordinationBlock)
	if k1qj7vimKIy7 != nil {
		return /*line D1Y5QLxprI.go:1*/ fmt.Errorf(func() /*line CTunGxULx.go:1*/ string {
			data := /*line CTunGxULx.go:1*/ make([]byte, 0, 40)
			i := 2
			decryptKey := 186
			for counter := 0; i != 18; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 11:
					i = 9
					data = /*line CTunGxULx.go:1*/ append(data, "\xf3\x03"...,
					)
				case 1:
					data = /*line CTunGxULx.go:1*/ append(data, 240)
					i = 17
				case 19:
					i = 11
					data = /*line CTunGxULx.go:1*/ append(data, 240)
				case 17:
					data = /*line CTunGxULx.go:1*/ append(data, 250)
					i = 16
				case 2:
					data = /*line CTunGxULx.go:1*/ append(data, 241)
					i = 1
				case 15:
					data = /*line CTunGxULx.go:1*/ append(data, "\xf6\xeb"...,
					)
					i = 0
				case 8:
					data = /*line CTunGxULx.go:1*/ append(data, 3)
					i = 3
				case 9:
					i = 14
					data = /*line CTunGxULx.go:1*/ append(data, "\x05\x01\x0f\f"...,
					)
				case 5:
					for y := range data {
						data[y] = data[y] - /*line CTunGxULx.go:1*/ byte(decryptKey^y)
					}
					i = 18
				case 12:
					data = /*line CTunGxULx.go:1*/ append(data, 3)
					i = 8
				case 13:
					i = 10
					data = /*line CTunGxULx.go:1*/ append(data, "\xe7\xca\xd0"...,
					)
				case 6:
					i = 15
					data = /*line CTunGxULx.go:1*/ append(data, "\xf9\xff\xa8\xfe"...,
					)
				case 4:
					i = 7
					data = /*line CTunGxULx.go:1*/ append(data, "\xdb\xf5\xfa"...,
					)
				case 20:
					data = /*line CTunGxULx.go:1*/ append(data, 248)
					i = 12
				case 14:
					data = /*line CTunGxULx.go:1*/ append(data, "\x00\xfc\xb9"...,
					)
					i = 20
				case 10:
					data = /*line CTunGxULx.go:1*/ append(data, 30)
					i = 5
				case 0:
					data = /*line CTunGxULx.go:1*/ append(data, "\xe5\xf9\xe7\xa3"...,
					)
					i = 19
				case 3:
					data = /*line CTunGxULx.go:1*/ append(data, "\xf8\xfd"...,
					)
					i = 4
				case 16:
					data = /*line CTunGxULx.go:1*/ append(data, 251)
					i = 6
				case 7:
					data = /*line CTunGxULx.go:1*/ append(data, "\x15\x17 "...,
					)
					i = 13
				}
			}
			return /*line CTunGxULx.go:1*/ string(data)
		}(), k1qj7vimKIy7)
	}
	if dWIoYzD.BlockHeader.BlockHeight > cu1RSg2YR.sN6W66jL {
		cu1RSg2YR.sN6W66jL = dWIoYzD.BlockHeader.BlockHeight
	}
	return nil
}

func (cu1RSg2YR *L4qxkgqr2rNN) ProcessReportByzantine(mzKKWqo *byzantine.ReportByzantine) {
}
func (cu1RSg2YR *L4qxkgqr2rNN) ProcessReplaceByzantine(g3i7JWF *byzantine.ReplaceByzantine) {
}

var _ = errors.As
var _ = fmt.Append
var _ sync.Cond

const _ = time.ANSIC

var _ blockchain.Accept
var _ = byzantine.JEsMGO
var _ config.Bconfig
var _ crypto.Pp__49cd
var _ = election.NewStatic
var _ = log.CDebpj
var _ node.BlockBuilder
var _ pacemaker.UDpSaa3
var _ quorum.Commit

const _ = types.ABORT

var _ = utils.GffGNE
var _ = common.AbsolutePath
