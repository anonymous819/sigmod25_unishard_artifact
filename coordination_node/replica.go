//line :1
package coordination_node

import (
	"encoding/gob"
	"time"

	"go.uber.org/atomic"

	blockchain "unishard/blockchain"
	byzantine "unishard/byzantine"
	config "unishard/config"
	crypto "unishard/crypto"
	election "unishard/election"
	log "unishard/log"
	mempool "unishard/mempool"
	message "unishard/message"
	node "unishard/node"
	pacemaker "unishard/pacemaker"
	pbft "unishard/pbft"
	quorum "unishard/quorum"
	types "unishard/types"
	utils "unishard/utils"
)

type (
	NkMVKP1upck struct {
		node.Node
		*pbft.L4qxkgqr2rNN
		*election.Static
		mT78Bso      *mempool.Producer
		eR3ITxPztL   *pacemaker.UDpSaa3
		hRxv_rDfIto  chan bool
		vSTadWcPct   atomic.Bool
		qE0frqWyh5   bool
		awyhpy       *time.Timer
		voFDpf2C3j   chan *blockchain.CoordinationBlock
		fSekax0HKJhe chan *blockchain.CoordinationBlock
		tjCx_L       chan *blockchain.CoordinationBlock
		j2dWReenf    chan interface{}
		bD6rp11      blockchain.O0GQK9U1
		jiEQJK_Z     blockchain.O0GQK9U1

		xeW2u54H     string
		b85Um3o      time.Time
		nra9E5SGz    time.Time
		kAULyZH20mY  time.Time
		noKmvQPr0iyr time.Time
		lya6BrALg4   time.Duration
		hQ06kw2n2k9z time.Duration
		pjfCawH97    time.Duration
		eUSdlh8pJ    time.Duration
		qSwbEXSzclch int
		vJew7rVYC    int
		rnwgyGvAt    int
		mpwDsO_Z3hdH int
		pQP7KRLBm    int
		cQDC845n     int
		n3KLOC       int
		ilcpjBW      int

		bSIArm0 chan *blockchain.CoordinationBlockWithoutHeader
	}
)

func I7kUh8uWo(foraXRqlF5 types.NodeID, tifmnW7 string, _yBYReQVMl bool, mk3Bfj types.Shard) *NkMVKP1upck {
	sNycjx := /*line aIaYJOGfd.go:1*/ new(NkMVKP1upck)
	sNycjx.Node = /*line qak4An.go:1*/ node.NewNode(foraXRqlF5, _yBYReQVMl, mk3Bfj)
	if _yBYReQVMl {
		/*line UFk_0axSl.go:1*/ log.ViJSfja(func() /*line aVABQ7QNa.go:1*/ string {
			fullData := [] /*line aVABQ7QNa.go:1*/ byte(" E\xb9\xac\xac\xcdN\x13K.Ql`m\x9f\xfct\xa4\uee72\xeeԃ\x03գ\x10\xab\x00/\xc3\xed\xc4")
			idxKey := [] /*line aVABQ7QNa.go:1*/ byte("\xa1\xf6QY6\xb3;\x99")
			data := /*line aVABQ7QNa.go:1*/ make([]byte, 0, 18)
			data = /*line aVABQ7QNa.go:1*/ append(data, fullData[169^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])]^fullData[186^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])], fullData[252^ /*line aVABQ7QNa.go:1*/ int(idxKey[1])]+fullData[224^ /*line aVABQ7QNa.go:1*/ int(idxKey[1])], fullData[128^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])]+fullData[181^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])], fullData[144^ /*line aVABQ7QNa.go:1*/ int(idxKey[7])]+fullData[135^ /*line aVABQ7QNa.go:1*/ int(idxKey[7])], fullData[214^ /*line aVABQ7QNa.go:1*/ int(idxKey[1])]-fullData[243^ /*line aVABQ7QNa.go:1*/ int(idxKey[1])], fullData[190^ /*line aVABQ7QNa.go:1*/ int(idxKey[5])]+fullData[188^ /*line aVABQ7QNa.go:1*/ int(idxKey[5])], fullData[158^ /*line aVABQ7QNa.go:1*/ int(idxKey[7])]+fullData[149^ /*line aVABQ7QNa.go:1*/ int(idxKey[7])], fullData[187^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])]-fullData[182^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])], fullData[50^ /*line aVABQ7QNa.go:1*/ int(idxKey[4])]^fullData[35^ /*line aVABQ7QNa.go:1*/ int(idxKey[4])], fullData[136^ /*line aVABQ7QNa.go:1*/ int(idxKey[7])]+fullData[128^ /*line aVABQ7QNa.go:1*/ int(idxKey[7])], fullData[67^ /*line aVABQ7QNa.go:1*/ int(idxKey[2])]-fullData[65^ /*line aVABQ7QNa.go:1*/ int(idxKey[2])], fullData[188^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])]-fullData[175^ /*line aVABQ7QNa.go:1*/ int(idxKey[0])], fullData[69^ /*line aVABQ7QNa.go:1*/ int(idxKey[3])]+fullData[70^ /*line aVABQ7QNa.go:1*/ int(idxKey[3])], fullData[244^ /*line aVABQ7QNa.go:1*/ int(idxKey[1])]-fullData[247^ /*line aVABQ7QNa.go:1*/ int(idxKey[1])], fullData[48^ /*line aVABQ7QNa.go:1*/ int(idxKey[6])]-fullData[35^ /*line aVABQ7QNa.go:1*/ int(idxKey[6])], fullData[181^ /*line aVABQ7QNa.go:1*/ int(idxKey[5])]^fullData[179^ /*line aVABQ7QNa.go:1*/ int(idxKey[5])], fullData[66^ /*line aVABQ7QNa.go:1*/ int(idxKey[2])]+fullData[82^ /*line aVABQ7QNa.go:1*/ int(idxKey[2])])
			return /*line aVABQ7QNa.go:1*/ string(data)
		}(), /*line Crfco8K.go:1*/ sNycjx.ID())
	}

	if /*line XaAnV4NwM.go:1*/ config.GetConfig().Master == "0" {
		sNycjx.Static = /*line na45PhaRH.go:1*/ election.NewStatic( /*line g_xsUva.go:1*/ utils.NewNodeID(1) /*line XzebIQZYgU9a.go:1*/, sNycjx.Shard(), config.Configuration.CommitteeNumber)
	}
	sNycjx.qE0frqWyh5 = _yBYReQVMl
	sNycjx.mT78Bso = /*line VVON2C3aai.go:1*/ mempool.NewProducer()
	sNycjx.eR3ITxPztL = /*line gTaTpaD.go:1*/ pacemaker.IM6FKYh( /*line Csk6Mg.go:1*/ config.GetConfig().CommitteeNumber)
	sNycjx.hRxv_rDfIto = /*line iUYQLDA.go:1*/ make(chan bool)
	sNycjx.j2dWReenf = /*line _XvuPuSSml.go:1*/ make(chan interface{})
	sNycjx.voFDpf2C3j = /*line alOqxVL.go:1*/ make(chan *blockchain.CoordinationBlock, 100)
	sNycjx.fSekax0HKJhe = /*line WoDbXcJ3L1O7.go:1*/ make(chan *blockchain.CoordinationBlock, 100)
	sNycjx.tjCx_L = /*line awblyoFtxq.go:1*/ make(chan *blockchain.CoordinationBlock)
	sNycjx.bD6rp11 = nil
	sNycjx.jiEQJK_Z = nil
	sNycjx.bSIArm0 = /*line Qltf_z.go:1*/ make(chan *blockchain.CoordinationBlockWithoutHeader, 100)

	/*line ClOWn_KN.go:1*/
	sNycjx.Register(blockchain.Qi_7sYpWjR8{}, sNycjx.HandleAccept)
	/*line fyAND4.go:1*/ sNycjx.Register(blockchain.CoordinationBlockWithoutHeader{}, sNycjx.HandleCoordinationBlockWithoutHeader)
	/*line dQNG1blzxU.go:1*/ sNycjx.Register(blockchain.CoordinationBlock{}, sNycjx.HandleCoordinationBlock)
	/*line lRal9AtBlS3o.go:1*/ sNycjx.Register(quorum.Collection[quorum.Vote]{}, sNycjx.HandleVote)
	/*line Nll_M3qr.go:1*/ sNycjx.Register(quorum.Collection[quorum.Commit]{}, sNycjx.HandleCommit)
	/*line vIalzwYzNg.go:1*/ sNycjx.Register(pacemaker.H8NP1AOKTF{}, sNycjx.HandleTmo)
	/*line G8KEN6gFJFNC.go:1*/ sNycjx.Register(pacemaker.RZ65hTAdXj{}, sNycjx.HandleTc)
	/*line AsCCWrpa.go:1*/ sNycjx.Register(byzantine.ReportByzantine{}, sNycjx.HandleReportByzantine)
	/*line CB9cakT.go:1*/ sNycjx.Register(byzantine.ReplaceByzantine{}, sNycjx.HandleReplaceByzantine)
	/*line srD_m_DQ.go:1*/ sNycjx.Register(message.Transaction{}, sNycjx.HandleTxn)
	/*line Vma5aj.go:1*/ sNycjx.Register(message.Query{}, sNycjx.tacS0M)
	/*line gRQRAw3jf.go:1*/ sNycjx.Register(message.RequestLeader{}, sNycjx.i5Y3A9uT)
	/*line EQPYA46.go:1*/ sNycjx.Register(message.ReportByzantine{}, sNycjx.ncNZka4T5b)

	/*line Mhk1F1u.go:1*/
	gob.Register(blockchain.Qi_7sYpWjR8{})
	/*line SLc4VcGK.go:1*/ gob.Register(blockchain.CoordinationBlock{})
	/*line DP6TIYVm.go:1*/ gob.Register(blockchain.CoordinationBlockWithoutHeader{})
	/*line HJ9wG4WxmUP.go:1*/ gob.Register(quorum.Collection[quorum.Vote]{})
	/*line iiawmmFVGa.go:1*/ gob.Register(quorum.Collection[quorum.Commit]{})
	/*line snUMzPq76Rc.go:1*/ gob.Register(pacemaker.RZ65hTAdXj{})
	/*line wqaFoml.go:1*/ gob.Register(pacemaker.H8NP1AOKTF{})
	/*line OmRnWg9nn.go:1*/ gob.Register(byzantine.ReportByzantine{})
	/*line JXGlvaTR.go:1*/ gob.Register(byzantine.ReplaceByzantine{})
	/*line aomrypd9t.go:1*/ gob.Register(message.ConsensusNodeRegister{})

	sNycjx.L4qxkgqr2rNN = /*line KtoeMXt3ujX.go:1*/ pbft.WtH1Mn_f(
		sNycjx.Node,
		sNycjx.eR3ITxPztL,
		sNycjx.Static,
		sNycjx.voFDpf2C3j,
		sNycjx.fSekax0HKJhe,
		sNycjx.tjCx_L,
	)
	return sNycjx
}

func (sNycjx *NkMVKP1upck) aJf0PZ(e0u7vzTo types.Epoch, qfFEdrtkLgdU types.View) {

	/*line KZFwyjOmKk1m.go:1*/
	log.ViJSfja(func() /*line M0sG9DD0o3dQ.go:1*/ string {
		key := [] /*line M0sG9DD0o3dQ.go:1*/ byte("kD.\xee2Tc\xb1\xd5\xcdO\xd5x\xc81O\x89\xaeN@\xa64\xb6@ \x87\x81}^^\\Pձ\x9b\x91eE?~\xae2\x0f?\xabH{[\xca\xd1\xd4\xd1\xd5\xd2\x1b]\xc9P\xfa_\xdbC^")
		data := [] /*line M0sG9DD0o3dQ.go:1*/ byte("0aX\xb3\x12$\x11\u07b6\xa8<\xa6X\xa6T8\xa9\xdc!5\xc8P\x96&O\xf5\xa1+7;+j\xf5\x94\xed\xb1 5P\x1d\xc6\b/\x1a\xddh\x90\xfdf:YE\xf59\x9e\xe5\"\xc3f\x04\xfe5\x03")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line M0sG9DD0o3dQ.go:1*/ string(data)
	}(), /*line Hs06i6we3c.go:1*/ sNycjx.ID(), qfFEdrtkLgdU, e0u7vzTo /*line CcI1UYd3.go:1*/, sNycjx.FindLeaderFor(qfFEdrtkLgdU, e0u7vzTo))
	if ! /*line R_K8ipYBP8.go:1*/ sNycjx.IsLeader( /*line J8wAUTu.go:1*/ sNycjx.ID(), qfFEdrtkLgdU, e0u7vzTo) {

		return
	}

	/*line jmvnDssG0.go:1*/
	sNycjx.SetRole(types.LEADER)
	sNycjx.j2dWReenf <- types.EpochView{Epoch: e0u7vzTo, View: qfFEdrtkLgdU}
}

func (sNycjx *NkMVKP1upck) a2JwLa(oW6IBJJc8PaV blockchain.O0GQK9U1) {
}

func (sNycjx *NkMVKP1upck) cDpcrL(oW6IBJJc8PaV blockchain.O0GQK9U1) {
}

func (sNycjx *NkMVKP1upck) x7D59YJrnEj(_06ZtlXy types.Epoch, grAmVu types.View) {
	kj0Ckj8mRolX := /*line J9LtMdRX.go:1*/ time.Now()
	oW6IBJJc8PaV := /*line qME9wi8rOwnb.go:1*/ new(blockchain.CoordinationBlock)
	oW6IBJJc8PaV.BlockHeader = /*line IhEthLz5vgI9.go:1*/ new(blockchain.CoordinationBlockHeader)
	oW6IBJJc8PaV.BlockData = /*line phUiwQ.go:1*/ new(blockchain.CoordinationBlockData)
	oW6IBJJc8PaV.QC = /*line OvnsAzh.go:1*/ new(quorum.HPOWa1PT0xzq)
	oW6IBJJc8PaV.CQC = /*line aePEyFXI.go:1*/ new(quorum.HPOWa1PT0xzq)

	sNycjx.cQDC845n++
	vvHMEj0_WXo := /*line ShqqZRmSMr.go:1*/ time.Now()
	ikPTNs := /*line X2rYjygH.go:1*/ vvHMEj0_WXo.Sub(kj0Ckj8mRolX)
	oW6IBJJc8PaV.BlockHeader.Timestamp = /*line g7q__rH.go:1*/ time.Now()
	sNycjx.lya6BrALg4 += ikPTNs

	/*line V1BMa8.go:1*/
	sNycjx.smzUvNu(oW6IBJJc8PaV)
}

func (sNycjx *NkMVKP1upck) smzUvNu(oW6IBJJc8PaV *blockchain.CoordinationBlock) {
	/*line JJjney.go:1*/ sNycjx.BroadcastToSome( /*line tCcOu9b.go:1*/ sNycjx.FindCommitteesFor(oW6IBJJc8PaV.BlockHeader.Epoch), oW6IBJJc8PaV)
	_ = /*line MzCba0np.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessCoordinationBlock(oW6IBJJc8PaV)
	sNycjx.noKmvQPr0iyr = /*line znhmxIFjU.go:1*/ time.Now()
	sNycjx.bD6rp11 = nil
}

func (sNycjx *NkMVKP1upck) ListenCommittedBlocks() {
	for {
		select {
		case zK8TOi := <-sNycjx.voFDpf2C3j:
			/*line KFaoDfh4oq.go:1*/ sNycjx.a2JwLa(zK8TOi)
		case huV7WEU := <-sNycjx.fSekax0HKJhe:
			/*line TIzwvaVV.go:1*/ sNycjx.cDpcrL(huV7WEU)
		case zmBgw0nld0 := <-sNycjx.tjCx_L:
			sNycjx.jiEQJK_Z = zmBgw0nld0
		}
	}
}

func (sNycjx *NkMVKP1upck) ListenLocalEvent() {
	sNycjx.b85Um3o = /*line e0d0PDyw5GWd.go:1*/ time.Now()
	sNycjx.awyhpy = /*line vcgsfWaUrOJa.go:1*/ time.NewTimer( /*line Yrpt77Y1.go:1*/ sNycjx.eR3ITxPztL.GetTimerForView())
	for {

		if /*line eEcUDyVVWKki.go:1*/ sNycjx.Role() == types.VALIDATOR {
			/*line ydhbA9Yw7z.go:1*/ sNycjx.awyhpy.Stop()
		} else {
			switch /*line dENxVk.go:1*/ sNycjx.State() {
			case types.READY:
				/*line aehgORTeZxZ.go:1*/ sNycjx.awyhpy.Reset( /*line aImHoiwsB.go:1*/ sNycjx.eR3ITxPztL.GetTimerForView())
			case types.VIEWCHANGING:
				/*line hjO20cBbd8.go:1*/ sNycjx.awyhpy.Reset( /*line mcwU95I.go:1*/ sNycjx.eR3ITxPztL.GetTimerForViewChange())
			}
		}
	L:
		for {
			select {
			case Y_eLX7_pU := <- /*line GrtCyULP.go:1*/ sNycjx.eR3ITxPztL.EnteringViewEvent():
				var (
					_06ZtlXy = Y_eLX7_pU.Epoch
					grAmVu   = Y_eLX7_pU.View
				)
				if grAmVu >= 2 {
					sNycjx.eUSdlh8pJ += /*line w5yS8mN4.go:1*/ time.Since(sNycjx.noKmvQPr0iyr)
				}

				wyXftBXbye := /*line HPQSYa.go:1*/ time.Now()
				viYjbvGJGar := /*line BnELaXpiujkj.go:1*/ wyXftBXbye.Sub(sNycjx.b85Um3o)
				sNycjx.pjfCawH97 += viYjbvGJGar
				sNycjx.vJew7rVYC++
				sNycjx.b85Um3o = wyXftBXbye

				if /*line DJHwiqf5.go:1*/ sNycjx.IsCommittee( /*line XBMA5O.go:1*/ sNycjx.ID(), _06ZtlXy) {
					/*line En9CrAazonc7.go:1*/ sNycjx.SetRole(types.COMMITTEE)
				} else {
					/*line eJi2lYM8e2.go:1*/ sNycjx.SetRole(types.VALIDATOR)
				}
				/*line JpEkqL3O5op.go:1*/ log.ViJSfja(func() /*line Q0Kx3l0RW88h.go:1*/ string {
					seed := /*line Q0Kx3l0RW88h.go:1*/ byte(133)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line Q0Kx3l0RW88h.go:1*/ append(data, x+seed); seed += x; return fnc }
					/*line Q0Kx3l0RW88h.go:1*/ fnc(214)(202)(81)(231)(195)(84)(244)(253)(187)(76)(245)(18)(1)(172)(86)(243)(252)(18)(169)(5)(81)(170)(76)(245)(18)(1)(255)(173)(5)(81)(170)(77)(252)(3)(0)(253)(10)(242)(254)(12)(255)(246)(15)
					return /*line Q0Kx3l0RW88h.go:1*/ string(data)
				}(), /*line ASI7Usp.go:1*/ sNycjx.ID(), grAmVu-1 /*line SCelIoil.go:1*/, viYjbvGJGar.Milliseconds())

				/*line kzhabIdaR.go:1*/
				sNycjx.aJf0PZ(_06ZtlXy, grAmVu)
				break L
			case vIdkuGW := <- /*line et34h70B.go:1*/ sNycjx.eR3ITxPztL.EnteringTmoEvent():

				/*line dAephfkP.go:1*/
				log.ViJSfja(func() /*line OZwymA.go:1*/ string {
					data := /*line OZwymA.go:1*/ make([]byte, 0, 41)
					i := 3
					decryptKey := 103
					for counter := 0; i != 15; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 3:
							data = /*line OZwymA.go:1*/ append(data, "\nu"...,
							)
							i = 4
						case 17:
							data = /*line OZwymA.go:1*/ append(data, "u\">"...,
							)
							i = 6
						case 7:
							i = 15
							for y := range data {
								data[y] = data[y] ^ /*line OZwymA.go:1*/ byte(decryptKey^y)
							}
						case 1:
							i = 13
							data = /*line OZwymA.go:1*/ append(data, "l!"...,
							)
						case 8:
							i = 5
							data = /*line OZwymA.go:1*/ append(data, "u8"...,
							)
						case 12:
							i = 9
							data = /*line OZwymA.go:1*/ append(data, ";a37"...,
							)
						case 13:
							i = 11
							data = /*line OZwymA.go:1*/ append(data, "+\x06\x06\x1a"...,
							)
						case 5:
							i = 12
							data = /*line OZwymA.go:1*/ append(data, "2<28"...,
							)
						case 11:
							i = 0
							data = /*line OZwymA.go:1*/ append(data, "\x17\x02"...,
							)
						case 0:
							i = 14
							data = /*line OZwymA.go:1*/ append(data, "TR"...,
							)
						case 2:
							data = /*line OZwymA.go:1*/ append(data, 37)
							i = 10
						case 6:
							i = 8
							data = /*line OZwymA.go:1*/ append(data, "3."...,
							)
						case 14:
							i = 7
							data = /*line OZwymA.go:1*/ append(data, 0)
						case 16:
							i = 2
							data = /*line OZwymA.go:1*/ append(data, ")'h-"...,
							)
						case 9:
							data = /*line OZwymA.go:1*/ append(data, "#70g"...,
							)
							i = 16
						case 10:
							i = 1
							data = /*line OZwymA.go:1*/ append(data, 63)
						case 4:
							data = /*line OZwymA.go:1*/ append(data, "%\x0f"...,
							)
							i = 17
						}
					}
					return /*line OZwymA.go:1*/ string(data)
				}(), /*line kbqELF2LTq5.go:1*/ sNycjx.ID(), vIdkuGW)
				/*line vk0dzBalo.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessLocalTmo( /*line nX304Iz0D.go:1*/ sNycjx.eR3ITxPztL.GetCurView())
				break L
			case <-sNycjx.awyhpy.C:

				/*line Z9NPyv1.go:1*/
				log.ViJSfja(func() /*line mtXa1cPeIoqE.go:1*/ string {
					fullData := [] /*line mtXa1cPeIoqE.go:1*/ byte("sk\xc8ś\xb5&\x8a\x0fA\x8e9?\x84=_\\\xfa\x8d\x1f%E\xf3,a\xb1\x8e\xf5\x94\x00\xd5x\xa9\xf5)$诸\x8c\xd3v \x16\xbe\xe5<\xcfʬ5+\xa7\xe4\xdb,\x8e\xf2S\xd6C]؉\xfc\xaeov\x9e\x8e\xec+\xacp@O")
					idxKey := [] /*line mtXa1cPeIoqE.go:1*/ byte("\xb8\xcc")
					data := /*line mtXa1cPeIoqE.go:1*/ make([]byte, 0, 39)
					data = /*line mtXa1cPeIoqE.go:1*/ append(data, fullData[234^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]-fullData[241^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[134^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[225^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[232^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]^fullData[136^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[167^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[172^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[151^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[157^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[223^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]^fullData[205^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[136^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[160^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[222^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]-fullData[230^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[231^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]^fullData[204^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[248^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]^fullData[206^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[161^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[150^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[176^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[188^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[173^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[142^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[132^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[143^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[141^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[201^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[180^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[168^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[202^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[135^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[251^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[248^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[165^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[128^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[166^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[241^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[250^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[163^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[139^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[237^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[169^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[135^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[134^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[137^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[219^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[254^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[141^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[178^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[183^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[139^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[235^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[208^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[254^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[191^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[240^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[187^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[177^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[155^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[174^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]^fullData[181^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[179^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[131^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[229^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]+fullData[245^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[144^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[162^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[246^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])]^fullData[194^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[1])], fullData[152^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]+fullData[148^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])], fullData[253^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])]-fullData[154^ /*line mtXa1cPeIoqE.go:1*/ int(idxKey[0])])
					return /*line mtXa1cPeIoqE.go:1*/ string(data)
				}(), /*line aBMMBUt8Uor.go:1*/ sNycjx.ID())

				if /*line QR3AMvU6Xxw.go:1*/ sNycjx.State() == types.VIEWCHANGING {
					/*line HAL6uo.go:1*/ sNycjx.eR3ITxPztL.UpdateNewView( /*line zwGF9A.go:1*/ sNycjx.eR3ITxPztL.GetNewView() + 1)
				}
				/*line Na3EyU9.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessLocalTmo( /*line mhqZ_TfUle7g.go:1*/ sNycjx.eR3ITxPztL.GetCurView())

				break L
			}
		}
	}
}

func (sNycjx *NkMVKP1upck) mH3iaIUnLeFn() {
	if ! /*line uK8_UyJgcC77.go:1*/ sNycjx.vSTadWcPct.Load() {
		sNycjx.nra9E5SGz = /*line bFEJlIOTkM.go:1*/ time.Now()
		sNycjx.kAULyZH20mY = /*line ti1Vw2Dg.go:1*/ time.Now()
		/*line YohV6kiM0jUG.go:1*/ log.ViJSfja(func() /*line On_vKgNwmCI.go:1*/ string {
			fullData := [] /*line On_vKgNwmCI.go:1*/ byte("\x82\xd5s\"\xbd\xaf.\xb4\xe7}\x8f\xfe\xf6\xfc\f'\xa5\x0f3\x0eo\x8a\xfd\x1c\xe0\xd6\v\xe6\xdd.\xb1@%\x02\n\xea|\xe3:1")
			idxKey := [] /*line On_vKgNwmCI.go:1*/ byte("3\xa7\xf07\xfb.\x9eCt#\\4\x8c\x80")
			data := /*line On_vKgNwmCI.go:1*/ make([]byte, 0, 21)
			data = /*line On_vKgNwmCI.go:1*/ append(data, fullData[129^ /*line On_vKgNwmCI.go:1*/ int(idxKey[13])]-fullData[163^ /*line On_vKgNwmCI.go:1*/ int(idxKey[13])], fullData[16^ /*line On_vKgNwmCI.go:1*/ int(idxKey[3])]^fullData[48^ /*line On_vKgNwmCI.go:1*/ int(idxKey[3])], fullData[144^ /*line On_vKgNwmCI.go:1*/ int(idxKey[12])]-fullData[172^ /*line On_vKgNwmCI.go:1*/ int(idxKey[12])], fullData[55^ /*line On_vKgNwmCI.go:1*/ int(idxKey[0])]+fullData[46^ /*line On_vKgNwmCI.go:1*/ int(idxKey[0])], fullData[95^ /*line On_vKgNwmCI.go:1*/ int(idxKey[10])]^fullData[66^ /*line On_vKgNwmCI.go:1*/ int(idxKey[10])], fullData[251^ /*line On_vKgNwmCI.go:1*/ int(idxKey[4])]-fullData[224^ /*line On_vKgNwmCI.go:1*/ int(idxKey[4])], fullData[164^ /*line On_vKgNwmCI.go:1*/ int(idxKey[13])]^fullData[143^ /*line On_vKgNwmCI.go:1*/ int(idxKey[13])], fullData[49^ /*line On_vKgNwmCI.go:1*/ int(idxKey[9])]-fullData[48^ /*line On_vKgNwmCI.go:1*/ int(idxKey[9])], fullData[156^ /*line On_vKgNwmCI.go:1*/ int(idxKey[6])]-fullData[136^ /*line On_vKgNwmCI.go:1*/ int(idxKey[6])], fullData[130^ /*line On_vKgNwmCI.go:1*/ int(idxKey[12])]-fullData[137^ /*line On_vKgNwmCI.go:1*/ int(idxKey[12])], fullData[231^ /*line On_vKgNwmCI.go:1*/ int(idxKey[2])]-fullData[253^ /*line On_vKgNwmCI.go:1*/ int(idxKey[2])], fullData[47^ /*line On_vKgNwmCI.go:1*/ int(idxKey[3])]+fullData[45^ /*line On_vKgNwmCI.go:1*/ int(idxKey[3])], fullData[34^ /*line On_vKgNwmCI.go:1*/ int(idxKey[5])]-fullData[49^ /*line On_vKgNwmCI.go:1*/ int(idxKey[5])], fullData[157^ /*line On_vKgNwmCI.go:1*/ int(idxKey[12])]^fullData[134^ /*line On_vKgNwmCI.go:1*/ int(idxKey[12])], fullData[12^ /*line On_vKgNwmCI.go:1*/ int(idxKey[5])]+fullData[11^ /*line On_vKgNwmCI.go:1*/ int(idxKey[5])], fullData[38^ /*line On_vKgNwmCI.go:1*/ int(idxKey[5])]+fullData[62^ /*line On_vKgNwmCI.go:1*/ int(idxKey[5])], fullData[22^ /*line On_vKgNwmCI.go:1*/ int(idxKey[3])]-fullData[62^ /*line On_vKgNwmCI.go:1*/ int(idxKey[3])], fullData[166^ /*line On_vKgNwmCI.go:1*/ int(idxKey[13])]^fullData[153^ /*line On_vKgNwmCI.go:1*/ int(idxKey[13])], fullData[253^ /*line On_vKgNwmCI.go:1*/ int(idxKey[4])]-fullData[238^ /*line On_vKgNwmCI.go:1*/ int(idxKey[4])], fullData[56^ /*line On_vKgNwmCI.go:1*/ int(idxKey[0])]^fullData[39^ /*line On_vKgNwmCI.go:1*/ int(idxKey[0])])
			return /*line On_vKgNwmCI.go:1*/ string(data)
		}(), /*line urtuYqGv3jOi.go:1*/ sNycjx.ID())
		/*line ItgiIoqRO3PK.go:1*/ sNycjx.vSTadWcPct.Store(true)
		sNycjx.hRxv_rDfIto <- true
	}
}

func (sNycjx *NkMVKP1upck) s3bZV5() {
	for {
		qhYcu_hfD := <-sNycjx.bSIArm0
		oW6IBJJc8PaV := /*line I2R5c5xGvVn.go:1*/ sNycjx.L4qxkgqr2rNN.CreateCoordinationBlock(qhYcu_hfD)
		oW6IBJJc8PaV.Proposer = /*line dwMIXFERoAi1.go:1*/ sNycjx.ID()
		iWZXicPh, _ := /*line Fwdq4QSaJik.go:1*/ crypto.PrivSign( /*line fwRqZfR.go:1*/ crypto.IDToByte(oW6IBJJc8PaV.BlockHash), nil)
		oW6IBJJc8PaV.CommitteeSignature = /*line YuvHPN.go:1*/ append(oW6IBJJc8PaV.CommitteeSignature, iWZXicPh)

		/*line aZ8FtMYW.go:1*/
		sNycjx.smzUvNu(oW6IBJJc8PaV)
		/*line mBfaWka.go:1*/ log.ViJSfja(func() /*line xDqTKAead5E.go:1*/ string {
			fullData := [] /*line xDqTKAead5E.go:1*/ byte("]\xbay\x9e\xcfq\xe7l\xf3\xbe\x91C\"Lc&\xce\xc4\xfb\xe2\xef\xf1!\xfb\xac\xa4\"\x03qޓ\xfa\u008ef\x88\x01 |\xfe\xb5\xb0T/h%\x99рQޖ\xbd\xa3\xa9\x0e\xd1Ï\xc1\f\x01\xfb\x0fC\x7f(\xebrX\xb1\xb16qF\xa1V:B6\xb4X\x1f\xc3!\xe6U\x0etǾ\xa7e\xc59\xad\xb6\x00\xaf\xfe\xe6\xf9\xc4\xc8ץ\x02\xb5\x8c;\xa2\xe1V1\xba[\x1ej\xd3d\x89\x14\xe5\x00\x85\x87\xd7;\x00\xe4̬Rt\x82c\xb3N\x99\x18\x1a\x1ej\x95\xf9\xe8")
			idxKey := [] /*line xDqTKAead5E.go:1*/ byte("\x92G\xa7\xf0\xc7f\xc4]\xdco\x85")
			data := /*line xDqTKAead5E.go:1*/ make([]byte, 0, 74)
			data = /*line xDqTKAead5E.go:1*/ append(data, fullData[243^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]+fullData[234^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[132^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]^fullData[215^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[74^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]+fullData[130^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[221^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]-fullData[151^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[253^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]+fullData[196^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[71^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])]-fullData[68^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])], fullData[221^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]^fullData[78^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[192^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])]-fullData[242^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])], fullData[190^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]^fullData[219^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[146^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]^fullData[253^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[138^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])]^fullData[202^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])], fullData[206^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])]^fullData[18^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])], fullData[74^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]-fullData[214^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[234^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])]^fullData[195^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])], fullData[251^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]+fullData[67^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[250^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])]-fullData[190^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])], fullData[229^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])]+fullData[17^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])], fullData[184^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])]^fullData[237^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])], fullData[115^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])]^fullData[224^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])], fullData[199^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]-fullData[198^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[133^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])]^fullData[234^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])], fullData[238^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]+fullData[197^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[23^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]+fullData[64^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[205^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]^fullData[254^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[28^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]-fullData[42^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[20^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]^fullData[95^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[7^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])]-fullData[46^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])], fullData[229^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]+fullData[231^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[131^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])]-fullData[220^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])], fullData[184^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])]-fullData[231^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])], fullData[87^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])]+fullData[1^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])], fullData[226^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]+fullData[188^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[245^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]-fullData[212^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[114^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])]+fullData[22^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])], fullData[214^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]+fullData[156^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[67^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]+fullData[179^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])], fullData[154^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]^fullData[180^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[54^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])]+fullData[84^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])], fullData[153^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]-fullData[148^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[145^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]+fullData[135^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[100^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]^fullData[94^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[191^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])]-fullData[238^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])], fullData[129^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])]+fullData[162^ /*line xDqTKAead5E.go:1*/ int(idxKey[0])], fullData[196^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]+fullData[146^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[153^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]-fullData[136^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[133^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]+fullData[149^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[123^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])]^fullData[78^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])], fullData[83^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])]-fullData[9^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])], fullData[123^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]-fullData[120^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[57^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])]^fullData[119^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])], fullData[21^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])]-fullData[89^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])], fullData[201^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]^fullData[45^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[246^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]-fullData[38^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[181^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]+fullData[208^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[53^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]-fullData[200^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[67^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])]^fullData[39^ /*line xDqTKAead5E.go:1*/ int(idxKey[7])], fullData[137^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]+fullData[218^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[212^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])]^fullData[160^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])], fullData[221^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])]^fullData[170^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])], fullData[9^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])]^fullData[196^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])], fullData[76^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])]-fullData[197^ /*line xDqTKAead5E.go:1*/ int(idxKey[1])], fullData[186^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])]-fullData[255^ /*line xDqTKAead5E.go:1*/ int(idxKey[3])], fullData[33^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])]+fullData[227^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])], fullData[163^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])]+fullData[154^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])], fullData[216^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])]-fullData[245^ /*line xDqTKAead5E.go:1*/ int(idxKey[6])], fullData[252^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])]+fullData[233^ /*line xDqTKAead5E.go:1*/ int(idxKey[8])], fullData[43^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])]+fullData[72^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])], fullData[15^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])]-fullData[41^ /*line xDqTKAead5E.go:1*/ int(idxKey[5])], fullData[204^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])]-fullData[210^ /*line xDqTKAead5E.go:1*/ int(idxKey[10])], fullData[255^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])]-fullData[101^ /*line xDqTKAead5E.go:1*/ int(idxKey[9])], fullData[203^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]^fullData[36^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[171^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])]+fullData[188^ /*line xDqTKAead5E.go:1*/ int(idxKey[2])], fullData[190^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])]-fullData[196^ /*line xDqTKAead5E.go:1*/ int(idxKey[4])])
			return /*line xDqTKAead5E.go:1*/ string(data)
		}(), /*line GG9UgjYr2.go:1*/ sNycjx.ID(), oW6IBJJc8PaV.BlockHash /*line Fn9rfOoJ.go:1*/, sNycjx.L4qxkgqr2rNN.GetLastBlockHeight() /*line Uy1XLdg9AD.go:1*/, sNycjx.eR3ITxPztL.GetCurView() /*line C8Jn2woLEE.go:1*/, sNycjx.eR3ITxPztL.GetCurEpoch())
	}
}

func (sNycjx *NkMVKP1upck) Start() {
	go /*line UVHPD_qLlv2g.go:1*/ sNycjx.Run()
	go /*line psAca2.go:1*/ sNycjx.s3bZV5()
	qgMD59zRcRm := /*line FHbhJXSC.go:1*/ sNycjx.ElectCommittees( /*line UxaoB1PQH.go:1*/ crypto.ZsUswPaGG4Z(0), 1)
	/*line wu14n_jm7Wzt.go:1*/ log.ViJSfja(func() /*line Ii8lATQ95oS.go:1*/ string {
		seed := /*line Ii8lATQ95oS.go:1*/ byte(218)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line Ii8lATQ95oS.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line Ii8lATQ95oS.go:1*/ fnc(129)(202)(81)(231)(195)(69)(7)(249)(254)(17)(172)(78)(247)(18)(169)(67)(12)(254)(0)(252)(11)(0)(241)(0)(14)(173)(5)(81)(170)(70)(9)(3)(174)(78)(247)(18)(169)(69)(11)(255)(244)(5)(184)(17)
		return /*line Ii8lATQ95oS.go:1*/ string(data)
	}(), /*line TBP28EJ9.go:1*/ sNycjx.ID(), qgMD59zRcRm)
	if /*line vi_VfhDJ2.go:1*/ sNycjx.IsCommittee( /*line IxWzuXyVZY.go:1*/ sNycjx.ID(), 1) {
		/*line c6AgxJ6.go:1*/ sNycjx.SetRole(types.COMMITTEE)
	}

	<-sNycjx.hRxv_rDfIto
	go /*line dhrEPWCF.go:1*/ sNycjx.ListenLocalEvent()
	go /*line Vf9KXSL0Yr.go:1*/ sNycjx.ListenCommittedBlocks()
	for {
		for /*line Fb_EQ5aPt.go:1*/ sNycjx.vSTadWcPct.Load() {
			g7MVoY84op := <-sNycjx.j2dWReenf
			switch yTCcCswT_ := g7MVoY84op.(type) {
			case types.EpochView:

			case blockchain.CoordinationBlock:
				/*line ByGuGW3yg.go:1*/ log.ViJSfja(func() /*line AOIqjZUcly.go:1*/ string {
					seed := /*line AOIqjZUcly.go:1*/ byte(214)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line AOIqjZUcly.go:1*/ append(data, x^seed); seed += x; return fnc }
					/*line AOIqjZUcly.go:1*/ fnc(165)(15)(235)(7)(8)(164)(120)(210)(29)(236)(30)(234)(240)(83)(133)(36)(0)(29)(232)(29)(255)(241)(245)(31)(250)(225)(50)(206)(31)(236)(16)
					return /*line AOIqjZUcly.go:1*/ string(data)
				}())
				_ = /*line G0GCYgYFH.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessCoordinationBlock(&yTCcCswT_)
				sNycjx.noKmvQPr0iyr = /*line _ydzMJPtWra.go:1*/ time.Now()
				sNycjx.n3KLOC++

			case quorum.Collection[quorum.Vote]:
				dsXeGBVkajz := /*line udnjMbT.go:1*/ time.Now()
				/*line APQljZnzUqiG.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessVote(&yTCcCswT_)
				ejHNkCq__r := /*line uWkMZz75S.go:1*/ time.Since(dsXeGBVkajz)
				sNycjx.eUSdlh8pJ += ejHNkCq__r
				sNycjx.rnwgyGvAt++
			case quorum.Collection[quorum.Commit]:

				/*line VUO7AFb.go:1*/
				sNycjx.L4qxkgqr2rNN.ProcessCommit(&yTCcCswT_)

			case blockchain.Qi_7sYpWjR8:
				/*line zvgA6JNNo.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessAccept(&yTCcCswT_)
			case pacemaker.H8NP1AOKTF:
				/*line PSfJ29a.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessRemoteTmo(&yTCcCswT_)
			case pacemaker.RZ65hTAdXj:
				/*line FLX5EMgBXd.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessTC(&yTCcCswT_)
			case byzantine.ReportByzantine:
				/*line JjTZstQ5.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessReportByzantine(&yTCcCswT_)
			case byzantine.ReplaceByzantine:
				/*line jQjI92QjY39J.go:1*/ sNycjx.L4qxkgqr2rNN.ProcessReplaceByzantine(&yTCcCswT_)
			}
		}
	}
}

var _ gob.CommonType

const _ = time.ANSIC

var _ atomic.Bool
var _ blockchain.Accept
var _ = byzantine.JEsMGO
var _ config.Bconfig
var _ crypto.Pp__49cd
var _ = election.NewStatic
var _ = log.CDebpj
var _ mempool.DUUdwSwZTCm
var _ message.ClientStart
var _ node.BlockBuilder
var _ pacemaker.UDpSaa3
var _ pbft.L4qxkgqr2rNN
var _ quorum.Commit

const _ = types.ABORT

var _ = utils.GffGNE
