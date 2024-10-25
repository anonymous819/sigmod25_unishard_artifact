//line :1
package coordination_node

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	blockchain "unishard/blockchain"
	byzantine "unishard/byzantine"
	log "unishard/log"
	message "unishard/message"
	pacemaker "unishard/pacemaker"
	quorum "unishard/quorum"
	types "unishard/types"
	utils "unishard/utils"
)

func (sNycjx *NkMVKP1upck) HandleCoordinationBlockWithoutHeader(vEBx6Yr0E4Ro blockchain.CoordinationBlockWithoutHeader) {
	/*line VWBFJ4X8Q0s.go:1*/ log.ViJSfja(func() /*line BNRo5maO.go:1*/ string {
		data := [] /*line BNRo5maO.go:1*/ byte("(\x12\x84\x16N\x9e\xac\xd1\"\x15rSin7{\x92<\t\x1aSR8\x1dW\xf5\"`Wu\xef\x88\xe52dI\xa8\x9f>\xf7~=r?O\x18\xcb\x10$Eid\xaa\xde\x02v\x1aL>p\x8d[\x8ds \u0604\x85G\x12\xf5}\xcfa")
		positions := [...]byte{4, 59, 16, 68, 45, 71, 61, 20, 32, 47, 15, 11, 47, 65, 21, 35, 31, 63, 37, 11, 43, 22, 27, 48, 25, 53, 6, 69, 7, 69, 32, 47, 1, 25, 17, 33, 26, 14, 58, 67, 72, 63, 40, 58, 70, 9, 25, 28, 5, 54, 45, 30, 47, 9, 18, 59, 4, 15, 41, 58, 47, 4, 45, 49, 16, 56, 38, 57, 70, 1, 72, 53, 35, 3, 20, 20, 54, 11, 52, 68, 44, 56, 53, 54, 7, 46, 1, 8, 66, 60, 67, 62, 23, 6, 54, 69, 21, 23, 2, 54, 19, 20, 2, 39, 49, 36}
		for i := 0; i < 106; i += 2 {
			localKey := /*line BNRo5maO.go:1*/ byte(i) + /*line BNRo5maO.go:1*/ byte(positions[i]^positions[i+1]) + 11
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line BNRo5maO.go:1*/ string(data)
	}(), /*line bD4wFg.go:1*/ sNycjx.Shard() /*line A4tbLmZDDmXy.go:1*/, sNycjx.ID())

	sNycjx.qSwbEXSzclch++
	/*line S3ZkmzCTAo.go:1*/ sNycjx.mH3iaIUnLeFn()

	if /*line tCvNa4iS9BD.go:1*/ sNycjx.eR3ITxPztL.GetCurView() == 0 {
		/*line G_zX53uaDlwi.go:1*/ log.ViJSfja(func() /*line giQuI4CM.go:1*/ string {
			key := [] /*line giQuI4CM.go:1*/ byte("\t\x8bGNY\xbd\xf8\xcd\x1b\xf4\x8c\xa5\xfa>\xc02\xbcl &")
			data := [] /*line giQuI4CM.go:1*/ byte("R\x9a/\x0fǶ|\x94W\x80\x94\xcbx1\xb4=\xa7\x03L\xfa")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line giQuI4CM.go:1*/ string(data)
		}(), /*line nvuRO5CeD.go:1*/ sNycjx.ID())
		/*line qq8ZaLg9f.go:1*/ sNycjx.eR3ITxPztL.AdvanceView(0)
	}

	if ! /*line G_TMQoRYhoat.go:1*/ sNycjx.IsLeader( /*line BycTBI.go:1*/ sNycjx.ID() /*line qIt_KDaOmZ.go:1*/, sNycjx.eR3ITxPztL.GetCurView() /*line AaByk_QxHy.go:1*/, sNycjx.eR3ITxPztL.GetCurEpoch()) {
		return
	}

	/*line hp5Me27.go:1*/
	json.Marshal(vEBx6Yr0E4Ro)

	sNycjx.bSIArm0 <- &vEBx6Yr0E4Ro
}

func (sNycjx *NkMVKP1upck) HandleCoordinationBlock(oW6IBJJc8PaV blockchain.CoordinationBlock) {
	/*line h_KamJG.go:1*/ log.ViJSfja(func() /*line WdIKAKfy.go:1*/ string {
		seed := /*line WdIKAKfy.go:1*/ byte(0)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line WdIKAKfy.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line WdIKAKfy.go:1*/ fnc(72)(169)(95)(180)(112)(217)(144)(76)(152)(51)(88)(181)(111)(209)(181)(95)(196)(135)(226)(238)(223)(178)(108)(167)(52)(109)(216)(3)
		return /*line WdIKAKfy.go:1*/ string(data)
	}(), oW6IBJJc8PaV)
	sNycjx.qSwbEXSzclch++
	/*line Jiayf6irmhc.go:1*/ sNycjx.mH3iaIUnLeFn()

	if /*line vnukwm5.go:1*/ sNycjx.eR3ITxPztL.GetCurView() == 0 {
		/*line POkd6cx.go:1*/ log.ViJSfja(func() /*line vYg0bSi.go:1*/ string {
			data := /*line vYg0bSi.go:1*/ make([]byte, 0, 21)
			i := 6
			decryptKey := 93
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					data = /*line vYg0bSi.go:1*/ append(data, "![\t\r"...,
					)
					i = 3
				case 6:
					i = 0
					data = /*line vYg0bSi.go:1*/ append(data, "$[\v"...,
					)
				case 4:
					data = /*line vYg0bSi.go:1*/ append(data, "\x01\x01L"...,
					)
					i = 1
				case 1:
					for y := range data {
						data[y] = data[y] ^ /*line vYg0bSi.go:1*/ byte(decryptKey^y)
					}
					i = 2
				case 8:
					data = /*line vYg0bSi.go:1*/ append(data, "\x1f\f"...,
					)
					i = 4
				case 3:
					i = 5
					data = /*line vYg0bSi.go:1*/ append(data, "\x19\x05"...,
					)
				case 7:
					i = 8
					data = /*line vYg0bSi.go:1*/ append(data, "\x04\x01\x1d\x05"...,
					)
				case 5:
					data = /*line vYg0bSi.go:1*/ append(data, "\x02U"...,
					)
					i = 7
				}
			}
			return /*line vYg0bSi.go:1*/ string(data)
		}(), /*line pOLucpV1.go:1*/ sNycjx.ID())
		/*line LZTZaDdumLFy.go:1*/ sNycjx.eR3ITxPztL.AdvanceView(0)
	}

	if ! /*line CkI4NSI1Ejb.go:1*/ sNycjx.IsCommittee( /*line Czn5QVXpZL2u.go:1*/ sNycjx.ID() /*line XyUxMXQ2tl.go:1*/, oW6IBJJc8PaV.CoordinationBlockHeader().Epoch) {
		return
	}

	if oW6IBJJc8PaV.BlockData.GlobalCoordinationSequence == nil {
		oW6IBJJc8PaV.BlockData.GlobalCoordinationSequence = blockchain.Sequence{}
	}
	if oW6IBJJc8PaV.BlockData.GlobalSnapshot == nil {
		oW6IBJJc8PaV.BlockData.GlobalSnapshot = blockchain.LocalSnapshots{}
	}
	if oW6IBJJc8PaV.BlockData.GlobalContractBundle == nil {
		oW6IBJJc8PaV.BlockData.GlobalContractBundle = []*blockchain.LocalContract{}
	}

	sNycjx.j2dWReenf <- oW6IBJJc8PaV
}

func (sNycjx *NkMVKP1upck) HandleTxn(aPUydzsz message.Transaction) {

	/*line O7WjSo.go:1*/
	sNycjx.mH3iaIUnLeFn()

	if /*line btOV5hz6faN.go:1*/ sNycjx.eR3ITxPztL.GetCurView() == 0 {
		/*line JmUAfnOB06.go:1*/ log.ViJSfja(func() /*line ZG7ydT6.go:1*/ string {
			fullData := [] /*line ZG7ydT6.go:1*/ byte("\x93\xf7u\ab\xf1\xa6\x1e\xd6V\x9crm\ay\x99m\xf4\x9f\xf4QN\xa5\x01\x9b\xc1\x02\"V\xa5\xdfx\x86\xff\x16\xd5\n\xca4\xbf")
			idxKey := [] /*line ZG7ydT6.go:1*/ byte("\x18\x99\xedE\xb8|\xbf\xe9m\xb5\xae\x05\xf0")
			data := /*line ZG7ydT6.go:1*/ make([]byte, 0, 21)
			data = /*line ZG7ydT6.go:1*/ append(data, fullData[109^ /*line ZG7ydT6.go:1*/ int(idxKey[5])]-fullData[115^ /*line ZG7ydT6.go:1*/ int(idxKey[5])], fullData[29^ /*line ZG7ydT6.go:1*/ int(idxKey[0])]+fullData[62^ /*line ZG7ydT6.go:1*/ int(idxKey[0])], fullData[229^ /*line ZG7ydT6.go:1*/ int(idxKey[7])]-fullData[232^ /*line ZG7ydT6.go:1*/ int(idxKey[7])], fullData[118^ /*line ZG7ydT6.go:1*/ int(idxKey[5])]^fullData[101^ /*line ZG7ydT6.go:1*/ int(idxKey[5])], fullData[190^ /*line ZG7ydT6.go:1*/ int(idxKey[4])]-fullData[152^ /*line ZG7ydT6.go:1*/ int(idxKey[4])], fullData[13^ /*line ZG7ydT6.go:1*/ int(idxKey[11])]^fullData[19^ /*line ZG7ydT6.go:1*/ int(idxKey[11])], fullData[249^ /*line ZG7ydT6.go:1*/ int(idxKey[7])]+fullData[234^ /*line ZG7ydT6.go:1*/ int(idxKey[7])], fullData[177^ /*line ZG7ydT6.go:1*/ int(idxKey[9])]-fullData[162^ /*line ZG7ydT6.go:1*/ int(idxKey[9])], fullData[254^ /*line ZG7ydT6.go:1*/ int(idxKey[12])]-fullData[253^ /*line ZG7ydT6.go:1*/ int(idxKey[12])], fullData[96^ /*line ZG7ydT6.go:1*/ int(idxKey[5])]^fullData[103^ /*line ZG7ydT6.go:1*/ int(idxKey[5])], fullData[61^ /*line ZG7ydT6.go:1*/ int(idxKey[0])]+fullData[17^ /*line ZG7ydT6.go:1*/ int(idxKey[0])], fullData[165^ /*line ZG7ydT6.go:1*/ int(idxKey[4])]^fullData[155^ /*line ZG7ydT6.go:1*/ int(idxKey[4])], fullData[88^ /*line ZG7ydT6.go:1*/ int(idxKey[5])]^fullData[99^ /*line ZG7ydT6.go:1*/ int(idxKey[5])], fullData[189^ /*line ZG7ydT6.go:1*/ int(idxKey[10])]^fullData[182^ /*line ZG7ydT6.go:1*/ int(idxKey[10])], fullData[243^ /*line ZG7ydT6.go:1*/ int(idxKey[7])]+fullData[226^ /*line ZG7ydT6.go:1*/ int(idxKey[7])], fullData[121^ /*line ZG7ydT6.go:1*/ int(idxKey[8])]+fullData[106^ /*line ZG7ydT6.go:1*/ int(idxKey[8])], fullData[189^ /*line ZG7ydT6.go:1*/ int(idxKey[6])]^fullData[157^ /*line ZG7ydT6.go:1*/ int(idxKey[6])], fullData[229^ /*line ZG7ydT6.go:1*/ int(idxKey[12])]-fullData[238^ /*line ZG7ydT6.go:1*/ int(idxKey[12])], fullData[36^ /*line ZG7ydT6.go:1*/ int(idxKey[11])]-fullData[5^ /*line ZG7ydT6.go:1*/ int(idxKey[11])], fullData[167^ /*line ZG7ydT6.go:1*/ int(idxKey[9])]^fullData[146^ /*line ZG7ydT6.go:1*/ int(idxKey[9])])
			return /*line ZG7ydT6.go:1*/ string(data)
		}(), /*line v_oI2Xe10Lk.go:1*/ sNycjx.ID())
		/*line uY2rfG.go:1*/ sNycjx.eR3ITxPztL.AdvanceView(0)
	}

	if aPUydzsz.TXType == types.ABORT {
		_06ZtlXy, grAmVu := /*line EYXQsgQaESB.go:1*/ sNycjx.eR3ITxPztL.GetCurEpochView()
		if ! /*line xt1gBsLx.go:1*/ sNycjx.IsLeader( /*line frcCpG1df.go:1*/ sNycjx.ID(), grAmVu, _06ZtlXy) {
			/*line c7NhrvSaKMTB.go:1*/ sNycjx.Send( /*line kHI5rE.go:1*/ sNycjx.FindLeaderFor(grAmVu, _06ZtlXy), aPUydzsz)
			return
		}

		/*line AqUgqlrm862n.go:1*/
		sNycjx.mT78Bso.CollectTxn(&aPUydzsz)
	} else {
		/*line SNF1vSzVoJu.go:1*/ sNycjx.mT78Bso.AddTxn(&aPUydzsz)
	}
}

func (sNycjx *NkMVKP1upck) HandleVote(nM_xlemn2jC quorum.Collection[quorum.Vote]) {
	/*line GouWo2pYF.go:1*/ sNycjx.mH3iaIUnLeFn()
	if /*line ae0ApQqXDld.go:1*/ sNycjx.State() == types.VIEWCHANGING {
		return
	}

	if ! /*line obmoFFx9RP.go:1*/ sNycjx.IsCommittee( /*line rs2bvKu1JHFP.go:1*/ sNycjx.ID(), nM_xlemn2jC.Epoch) {
		return
	}

	if nM_xlemn2jC.BlockHeight < /*line aAkoIXa.go:1*/ sNycjx.GetHighBlockHeight() {
		return
	}

	/*line K3PHb0AsZl8.go:1*/
	log.ViJSfja(func() /*line qTIePDc.go:1*/ string {
		seed := /*line qTIePDc.go:1*/ byte(102)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line qTIePDc.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line qTIePDc.go:1*/ fnc(245)(202)(81)(231)(195)(8)(32)(25)(13)(246)(8)(249)(241)(25)(5)(241)(196)(247)(82)(243)(254)(2)(4)(13)(239)(255)(188)(86)(249)(5)(241)(187)(70)(12)(253)(254)(179)(59)(202)(81)(231)(207)(244)(66)(10)(3)(244)(8)(221)(29)(4)(254)(1)(12)(198)(230)(5)(81)(170)(86)(243)(252)(18)(195)(230)(5)(81)(170)(69)(11)(255)(244)(5)(210)(230)(5)(81)(170)(41)(251)(246)(230)(5)(83)
		return /*line qTIePDc.go:1*/ string(data)
	}(), /*line LMrfblSn.go:1*/ sNycjx.ID(), nM_xlemn2jC.WpsY_aZ, nM_xlemn2jC.BlockHeight, nM_xlemn2jC.View, nM_xlemn2jC.Epoch, nM_xlemn2jC.NaNYri773M)
	sNycjx.j2dWReenf <- nM_xlemn2jC
}

func (sNycjx *NkMVKP1upck) HandleCommit(_1clKS_ quorum.Collection[quorum.Commit]) {
	/*line LCsfDcH.go:1*/ sNycjx.mH3iaIUnLeFn()
	if /*line GcjxgWHX.go:1*/ sNycjx.State() == types.VIEWCHANGING {
		return
	}

	if ! /*line lCy13R3N5.go:1*/ sNycjx.IsCommittee( /*line KhaN8rC5.go:1*/ sNycjx.ID(), _1clKS_.Epoch) {
		return
	}

	if _1clKS_.BlockHeight < /*line s0Z8PED4Zg0.go:1*/ sNycjx.GetLastBlockHeight() {
		return
	}
	/*line cjXlpbkV.go:1*/ log.ViJSfja(func() /*line clBV_Ma6b0C0.go:1*/ string {
		key := [] /*line clBV_Ma6b0C0.go:1*/ byte("\x045\xed\"?\xed\xb4\xf9\x16\x1cw^Z=Mj\x1c\xa4 \x14\xcc\ts3\xed\xbd\xf9\x8a\xb6'$1\xc7z\xb8w\x80.<\xca\xedo\x0e\x1f\xc5\xf7\xdc\xd9u\xf1\x02\x88\\\xb60\x15~S3嶥Z\x981\xa9\x81_(і\xfa\xc2ӹ\xadR~\x8cm\x06\xc3\xc4r\x93\x95\xe61")
		data := [] /*line clBV_Ma6b0C0.go:1*/ byte("_\x10\x9b\x7f\x1f\xc5\xfc\x98xx\x1b;\x19R \au\xd0\t4\xbel\x10V\x84˜\xee\x96DK\\\xaa\x13\xccW\xe6\\S\xa7\xcd4+i\x98\xdb\xfc\xbb\x19\x9ea\xe34\xd3Yr\x16'\tœ\xd3z\xeeX\xcc\xf6e\b\xf4\xe0ڧ\xa3\xd6\xce:D\xacHp\xe3\x8d6\xa9\xb5\xc3I")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line clBV_Ma6b0C0.go:1*/ string(data)
	}(), /*line mhpWXS.go:1*/ sNycjx.ID(), _1clKS_.WpsY_aZ, _1clKS_.BlockHeight, _1clKS_.View, _1clKS_.Epoch, _1clKS_.NaNYri773M)
	sNycjx.j2dWReenf <- _1clKS_
}

func (sNycjx *NkMVKP1upck) HandleAccept(eIrGsbbA2B2 blockchain.Qi_7sYpWjR8) {
	/*line gKbG9Ot.go:1*/ sNycjx.mH3iaIUnLeFn()

	if /*line I5lTpc4HqsF.go:1*/ sNycjx.eR3ITxPztL.GetCurView() == 0 {
		/*line _S40P5.go:1*/ log.ViJSfja(func() /*line J7GMvBawjaE.go:1*/ string {
			data := [] /*line J7GMvBawjaE.go:1*/ byte("\x12u\xc0\xb9zs\xa2irthǭ\x14\xc0\xbaro\xcc ")
			positions := [...]byte{7, 10, 18, 13, 6, 3, 16, 0, 2, 15, 12, 0, 7, 16, 12, 2, 14, 12, 11, 14, 4, 1, 16, 18}
			for i := 0; i < 24; i += 2 {
				localKey := /*line J7GMvBawjaE.go:1*/ byte(i) + /*line J7GMvBawjaE.go:1*/ byte(positions[i]^positions[i+1]) + 60
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line J7GMvBawjaE.go:1*/ string(data)
		}(), /*line oAebqBjz3heK.go:1*/ sNycjx.ID())
		/*line qU1RUXgd3ET.go:1*/ sNycjx.eR3ITxPztL.AdvanceView(0)
	}

	yfSSaR := eIrGsbbA2B2.XjPd77f0
	if /*line TbSaa4r.go:1*/ yfSSaR.CoordinationBlockHeader().BlockHeight < /*line EiIj3I.go:1*/ sNycjx.GetLastBlockHeight() {
		return
	}
	sNycjx.j2dWReenf <- eIrGsbbA2B2
}

func (sNycjx *NkMVKP1upck) HandleTmo(pNSBGys pacemaker.H8NP1AOKTF) {
	if pNSBGys.W72t5Nhk < /*line HUUN1DXgLF.go:1*/ sNycjx.eR3ITxPztL.GetCurView() {
		return
	}

	if ! /*line p60CBD4aP6N.go:1*/ sNycjx.IsCommittee( /*line zLRSGVb.go:1*/ sNycjx.ID(), pNSBGys.Epoch) {
		return
	}

	/*line g9Uy42s.go:1*/
	log.ViJSfja(func() /*line a6T1EwTyFre.go:1*/ string {
		key := [] /*line a6T1EwTyFre.go:1*/ byte("\xfa4\x02\xb8\xb6v\xe6\x1f\xc1k\xa1~*\xf6\x8c\xdc[`\xfa_gY\xf9夒zP\t\xea\xd6\xd0%\x87\xfd\x8b\xdd2j\xc0\x8f\bt\xdb~\xcfX\x93\x10\xe8\x00\xcd\xc1g!")
		data := [] /*line a6T1EwTyFre.go:1*/ byte("\xa1\x11t\xe5\x96^\xae~\xaf\x0f\xcd\x1b~\x9b\xe3\xf5{\x12\x9f<\x020\x8f\x80\xc0\xb2\x1bp}\x83\xbb\xb5J\U00089afb@\x05\xad\xaf-\x02\xfb\x18\xa0*\xb3f\x81e\xba\xe1BW")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line a6T1EwTyFre.go:1*/ string(data)
	}(), /*line ql_Oy4A.go:1*/ sNycjx.ID(), pNSBGys.RKtUCo69O, pNSBGys.W72t5Nhk)
	sNycjx.j2dWReenf <- pNSBGys
}

func (sNycjx *NkMVKP1upck) HandleTc(mlMN1Gj8 pacemaker.RZ65hTAdXj) {
	if !(mlMN1Gj8.APadJA > /*line OOjbVk5B3GtQ.go:1*/ sNycjx.eR3ITxPztL.GetCurView()) {
		return
	}
	/*line Isv7aLBzbz.go:1*/ log.ViJSfja(func() /*line cClImgTFVoYU.go:1*/ string {
		data := /*line cClImgTFVoYU.go:1*/ make([]byte, 0, 50)
		i := 19
		decryptKey := 68
		for counter := 0; i != 9; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 0
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xdf$,"...,
				)
			case 5:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xda\xde"...,
				)
				i = 3
			case 11:
				data = /*line cClImgTFVoYU.go:1*/ append(data, 236)
				i = 16
			case 10:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xfe\xf0\xed"...,
				)
				i = 15
			case 20:
				i = 18
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\x1d."...,
				)
			case 12:
				i = 13
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xc0\xc7\xe6\xfe"...,
				)
			case 14:
				data = /*line cClImgTFVoYU.go:1*/ append(data, 164)
				i = 2
			case 0:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "&\xd3(\x1a"...,
				)
				i = 20
			case 7:
				i = 11
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xa2\xf5\xeb\xa7"...,
				)
			case 8:
				for y := range data {
					data[y] = data[y] + /*line cClImgTFVoYU.go:1*/ byte(decryptKey^y)
				}
				i = 9
			case 15:
				data = /*line cClImgTFVoYU.go:1*/ append(data, 238)
				i = 4
			case 6:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xbf\xb5"...,
				)
				i = 10
			case 2:
				data = /*line cClImgTFVoYU.go:1*/ append(data, 228)
				i = 7
			case 18:
				i = 8
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xd6\xda\""...,
				)
			case 4:
				i = 14
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xf9\x05\xf3\xf1"...,
				)
			case 19:
				i = 12
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xf7\xc0\x10\xf6"...,
				)
			case 16:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xf7+"...,
				)
				i = 21
			case 21:
				i = 5
				data = /*line cClImgTFVoYU.go:1*/ append(data, 40)
			case 17:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\xfe\xf6\xec\xfa"...,
				)
				i = 6
			case 3:
				i = 1
				data = /*line cClImgTFVoYU.go:1*/ append(data, 54)
			case 13:
				data = /*line cClImgTFVoYU.go:1*/ append(data, "\x02\xf7"...,
				)
				i = 17
			}
		}
		return /*line cClImgTFVoYU.go:1*/ string(data)
	}(), /*line VmXC9Gqyzt.go:1*/ sNycjx.ID(), mlMN1Gj8.H5NbPTQ, mlMN1Gj8.APadJA)
	sNycjx.j2dWReenf <- mlMN1Gj8
}

func (sNycjx *NkMVKP1upck) HandleReportByzantine(gE5mOZ byzantine.ReportByzantine) {
	/*line N1Ts63.go:1*/ log.ViJSfja(func() /*line M3WNMOy84.go:1*/ string {
		data := [] /*line M3WNMOy84.go:1*/ byte("[\xd1v\xa2 r\x91\xd4\xd9\x18\xbe ;\xeczqT{0ne\x0er\xf24R,\x95h9\xc2oV\xdcpMb\xb9\x7f\xbf\xe7\f\fy %\xec")
		positions := [...]byte{7, 25, 46, 17, 33, 27, 25, 37, 28, 40, 27, 29, 18, 38, 37, 32, 16, 15, 9, 10, 37, 32, 15, 35, 10, 7, 3, 12, 8, 25, 38, 39, 46, 16, 30, 26, 8, 24, 15, 30, 40, 10, 3, 8, 27, 13, 1, 6, 21, 32, 23, 41, 18, 42}
		for i := 0; i < 54; i += 2 {
			localKey := /*line M3WNMOy84.go:1*/ byte(i) + /*line M3WNMOy84.go:1*/ byte(positions[i]^positions[i+1]) + 55
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line M3WNMOy84.go:1*/ string(data)
	}(), /*line ijRMK6.go:1*/ sNycjx.ID(), gE5mOZ.DDj8mepzC)

	/*line EfTl2OuoQ885.go:1*/
	sNycjx.mH3iaIUnLeFn()

	sNycjx.j2dWReenf <- gE5mOZ
}

func (sNycjx *NkMVKP1upck) HandleReplaceByzantine(h2cC0Il8uqlq byzantine.ReplaceByzantine) {
	/*line EQdDQx.go:1*/ log.ViJSfja(func() /*line QRdDCkvv.go:1*/ string {
		data := /*line QRdDCkvv.go:1*/ make([]byte, 0, 74)
		i := 5
		decryptKey := 238
		for counter := 0; i != 18; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 12:
				data = /*line QRdDCkvv.go:1*/ append(data, 126)
				i = 22
			case 27:
				i = 8
				data = /*line QRdDCkvv.go:1*/ append(data, "\x7fi"...,
				)
			case 5:
				data = /*line QRdDCkvv.go:1*/ append(data, "y\x06V|"...,
				)
				i = 23
			case 2:
				i = 20
				data = /*line QRdDCkvv.go:1*/ append(data, "5t"...,
				)
			case 20:
				i = 7
				data = /*line QRdDCkvv.go:1*/ append(data, "~o"...,
				)
			case 9:
				i = 0
				data = /*line QRdDCkvv.go:1*/ append(data, 70)
			case 11:
				data = /*line QRdDCkvv.go:1*/ append(data, " nj"...,
				)
				i = 3
			case 16:
				data = /*line QRdDCkvv.go:1*/ append(data, "KL\x0e"...,
				)
				i = 21
			case 10:
				data = /*line QRdDCkvv.go:1*/ append(data, "\x1c[mq"...,
				)
				i = 11
			case 6:
				i = 13
				data = /*line QRdDCkvv.go:1*/ append(data, "k35g"...,
				)
			case 15:
				data = /*line QRdDCkvv.go:1*/ append(data, "M\\MK"...,
				)
				i = 10
			case 21:
				data = /*line QRdDCkvv.go:1*/ append(data, "mUWS"...,
				)
				i = 26
			case 13:
				i = 2
				data = /*line QRdDCkvv.go:1*/ append(data, "6c{"...,
				)
			case 22:
				i = 4
				data = /*line QRdDCkvv.go:1*/ append(data, "q\v\x00@"...,
				)
			case 23:
				i = 9
				data = /*line QRdDCkvv.go:1*/ append(data, "\x06UAU"...,
				)
			case 26:
				data = /*line QRdDCkvv.go:1*/ append(data, "]D"...,
				)
				i = 24
			case 8:
				data = /*line QRdDCkvv.go:1*/ append(data, "d`m"...,
				)
				i = 25
			case 0:
				i = 16
				data = /*line QRdDCkvv.go:1*/ append(data, 74)
			case 4:
				i = 17
				data = /*line QRdDCkvv.go:1*/ append(data, 10)
			case 24:
				data = /*line QRdDCkvv.go:1*/ append(data, "XXR"...,
				)
				i = 14
			case 7:
				data = /*line QRdDCkvv.go:1*/ append(data, "9nj"...,
				)
				i = 12
			case 17:
				i = 1
				data = /*line QRdDCkvv.go:1*/ append(data, 3)
			case 25:
				i = 6
				data = /*line QRdDCkvv.go:1*/ append(data, "/gh"...,
				)
			case 1:
				i = 19
				data = /*line QRdDCkvv.go:1*/ append(data, "\x1eD@\x1c"...,
				)
			case 19:
				for y := range data {
					data[y] = data[y] ^ /*line QRdDCkvv.go:1*/ byte(decryptKey^y)
				}
				i = 18
			case 3:
				data = /*line QRdDCkvv.go:1*/ append(data, "c$u"...,
				)
				i = 27
			case 14:
				data = /*line QRdDCkvv.go:1*/ append(data, "\x14G_J"...,
				)
				i = 15
			}
		}
		return /*line QRdDCkvv.go:1*/ string(data)
	}(), /*line yoylsOaM.go:1*/ sNycjx.ID(), h2cC0Il8uqlq.NDKcLy, h2cC0Il8uqlq.JEKGSzf)
	/*line Iqsi8_.go:1*/ sNycjx.mH3iaIUnLeFn()
	sNycjx.j2dWReenf <- h2cC0Il8uqlq
}

func (sNycjx *NkMVKP1upck) tacS0M(aPUydzsz message.Query) {

	yzoMqvBskBJi := /*line auYywONG.go:1*/ float64( /*line yFkwKrwkt.go:1*/ sNycjx.hQ06kw2n2k9z.Milliseconds()) / /*line GVmEESf.go:1*/ float64(sNycjx.pQP7KRLBm)
	sNycjx.xeW2u54H += /*line dCoe5hwfT.go:1*/ fmt.Sprintf(func() /*line KCWEE1F.go:1*/ string {
		seed := /*line KCWEE1F.go:1*/ byte(136)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line KCWEE1F.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line KCWEE1F.go:1*/ fnc(220)(13)(28)(232)(79)(228)(141)(67)(88)(163)(93)(240)(148)(60)(226)(29)(250)(238)(31)(230)(9)(241)(76)(226)(129)(83)(88)(164)(12)(243)(92)(188)(129)
		return /*line KCWEE1F.go:1*/ string(data)
	}(), /*line wy3y1I2tlBZT.go:1*/ time.Since(sNycjx.nra9E5SGz).Seconds() /*line GCckvriL.go:1*/, float64(sNycjx.mpwDsO_Z3hdH)/ /*line kspN7i.go:1*/ time.Since(sNycjx.kAULyZH20mY).Seconds())
	sNycjx.mpwDsO_Z3hdH = 0
	sNycjx.kAULyZH20mY = /*line T2EnLfjS6xoK.go:1*/ time.Now()
	bRsf_K := /*line BqUljhn7.go:1*/ fmt.Sprintf(func() /*line Tlqdoe.go:1*/ string {
		fullData := [] /*line Tlqdoe.go:1*/ byte("\xfe\xe9\xfd\x12\x16\xda\"\x1b|\x83\xb9\xad\xd2\xc97\xe7\xe6u\x0f\xb3\xb4}\xccq\xf2\x05v\x9c")
		idxKey := [] /*line Tlqdoe.go:1*/ byte("\x19~")
		data := /*line Tlqdoe.go:1*/ make([]byte, 0, 15)
		data = /*line Tlqdoe.go:1*/ append(data, fullData[119^ /*line Tlqdoe.go:1*/ int(idxKey[1])]+fullData[115^ /*line Tlqdoe.go:1*/ int(idxKey[1])], fullData[17^ /*line Tlqdoe.go:1*/ int(idxKey[0])]-fullData[30^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[105^ /*line Tlqdoe.go:1*/ int(idxKey[1])]^fullData[103^ /*line Tlqdoe.go:1*/ int(idxKey[1])], fullData[23^ /*line Tlqdoe.go:1*/ int(idxKey[0])]-fullData[21^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[31^ /*line Tlqdoe.go:1*/ int(idxKey[0])]-fullData[13^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[12^ /*line Tlqdoe.go:1*/ int(idxKey[0])]+fullData[9^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[100^ /*line Tlqdoe.go:1*/ int(idxKey[1])]-fullData[124^ /*line Tlqdoe.go:1*/ int(idxKey[1])], fullData[22^ /*line Tlqdoe.go:1*/ int(idxKey[0])]-fullData[18^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[125^ /*line Tlqdoe.go:1*/ int(idxKey[1])]-fullData[102^ /*line Tlqdoe.go:1*/ int(idxKey[1])], fullData[11^ /*line Tlqdoe.go:1*/ int(idxKey[0])]+fullData[29^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[101^ /*line Tlqdoe.go:1*/ int(idxKey[1])]+fullData[123^ /*line Tlqdoe.go:1*/ int(idxKey[1])], fullData[10^ /*line Tlqdoe.go:1*/ int(idxKey[0])]^fullData[19^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[15^ /*line Tlqdoe.go:1*/ int(idxKey[0])]^fullData[24^ /*line Tlqdoe.go:1*/ int(idxKey[0])], fullData[8^ /*line Tlqdoe.go:1*/ int(idxKey[0])]+fullData[25^ /*line Tlqdoe.go:1*/ int(idxKey[0])])
		return /*line Tlqdoe.go:1*/ string(data)
	}(), yzoMqvBskBJi, sNycjx.xeW2u54H)

	/*line KX3da2JFm.go:1*/
	aPUydzsz.Reply(message.QueryReply{Info: bRsf_K})
}

func (sNycjx *NkMVKP1upck) i5Y3A9uT(aPUydzsz message.RequestLeader) {

	/*line nI6pPuv.go:1*/
	sNycjx.mH3iaIUnLeFn()

	if /*line jVHkc2.go:1*/ sNycjx.eR3ITxPztL.GetCurView() == 0 {
		/*line BzbjH0n.go:1*/ log.ViJSfja(func() /*line lVLiqE3fKj7b.go:1*/ string {
			seed := /*line lVLiqE3fKj7b.go:1*/ byte(132)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line lVLiqE3fKj7b.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line lVLiqE3fKj7b.go:1*/ fnc(223)(136)(97)(169)(21)(125)(251)(227)(215)(176)(12)(104)(210)(161)(71)(137)(6)(24)(45)(14)
			return /*line lVLiqE3fKj7b.go:1*/ string(data)
		}(), /*line JW0Gmr1w1.go:1*/ sNycjx.ID())
		/*line C8P_wDot.go:1*/ sNycjx.eR3ITxPztL.AdvanceView(0)
	}

	afOagFB5JyhY := /*line Zdg2C53AgfN6.go:1*/ sNycjx.FindLeaderFor( /*line HorKh9WXxR.go:1*/ sNycjx.eR3ITxPztL.GetCurView() /*line XVzLiEQ16v.go:1*/, sNycjx.eR3ITxPztL.GetCurEpoch())
	/*line eTI4LNl6Etg2.go:1*/ aPUydzsz.Reply(message.RequestLeaderReply{Leader: /*line m4raHaO.go:1*/ fmt.Sprint(afOagFB5JyhY)})
}

func (sNycjx *NkMVKP1upck) ncNZka4T5b(aPUydzsz message.ReportByzantine) {

	_06ZtlXy, grAmVu := /*line VFQGlzJzMtnd.go:1*/ sNycjx.eR3ITxPztL.GetCurEpochView()

	qHLJZ14 := /*line YdqL2tCFoUth.go:1*/ sNycjx.FindCommitteesFor(_06ZtlXy)
	/*line FDW6Sv0CI.go:1*/ log.CDebpj(qHLJZ14)
	aQSLv2i5cV := /*line oVZl5DQ.go:1*/ rand.New( /*line Y_1OFD.go:1*/ rand.NewSource( /*line Rrib4H.go:1*/ time.Now().UnixNano()))
	var lKC_UE83W types.NodeID
	for {
		lKC_UE83W = qHLJZ14[ /*line LEaWmjeq.go:1*/ aQSLv2i5cV.Intn( /*line rah73X4.go:1*/ utils.UB_qaFnx(qHLJZ14))]
		if ! /*line Q0XM68_U44x.go:1*/ sNycjx.IsLeader(lKC_UE83W, grAmVu, _06ZtlXy) {
			/*line VljELQT.go:1*/ log.ViJSfja(func() /*line UmGIotAaN2.go:1*/ string {
				key := [] /*line UmGIotAaN2.go:1*/ byte("dv\x13R\x7fIo\xf4I\xa5\xed\xf2xQ\xa6\x95!\xcd\n\xb8\x00\x7f\x938\xbeߝ")
				data := [] /*line UmGIotAaN2.go:1*/ byte("?Se\x0f_\x06\x03\x90i\xf5\x98\x9034ߵo\xa2n\xdd 6\xf7\x02\x9e\xfa\xeb")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line UmGIotAaN2.go:1*/ string(data)
			}(), /*line fxBEq4FsTEdc.go:1*/ sNycjx.ID(), lKC_UE83W)
			break
		}
	}
	gE5mOZ := /*line i1_bXShhwaB.go:1*/ byzantine.MakeReportByzantine(lKC_UE83W /*line MmtmwBGi.go:1*/, sNycjx.Shard(), _06ZtlXy, grAmVu)

	afOagFB5JyhY := /*line TV1zfof.go:1*/ sNycjx.Static.FindLeaderFor(grAmVu, _06ZtlXy)
	if afOagFB5JyhY == /*line wwSd1Xz03.go:1*/ sNycjx.ID() {
		/*line ZAnKBmbffj.go:1*/ sNycjx.HandleReportByzantine(*gE5mOZ)
	} else {
		/*line x1i56D.go:1*/ sNycjx.Send(afOagFB5JyhY, gE5mOZ)
	}
}

var _ = json.Compact
var _ = fmt.Append
var _ = rand.ExpFloat64

const _ = time.ANSIC

var _ blockchain.Accept
var _ = byzantine.JEsMGO
var _ = log.CDebpj
var _ message.ClientStart
var _ pacemaker.UDpSaa3
var _ quorum.Commit

const _ = types.ABORT

var _ = utils.GffGNE
