//line :1
package vm

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/params"
)

func sssg6yQAr(iu5VUy1y0GT *VfNB1WFtybE, hfuPa0Wld2 uint64) (uint64, error) {
	if hfuPa0Wld2 == 0 {
		return 0, nil
	}

	if hfuPa0Wld2 > 0x1FFFFFFFE0 {
		return 0, QiOgwB7BT
	}
	woNLZzHpa7 := /*line GbMNBQeQAUV.go:1*/ f6N2wznhW(hfuPa0Wld2)
	hfuPa0Wld2 = woNLZzHpa7 * 32

	if hfuPa0Wld2 > /*line jkMOpZGOvzpA.go:1*/ uint64( /*line CTboqn.go:1*/ iu5VUy1y0GT.Len()) {
		jVoNrJ := woNLZzHpa7 * woNLZzHpa7
		yQaDA3L1 := woNLZzHpa7 * params.MemoryGas
		ba0sUoWL := jVoNrJ / params.QuadCoeffDiv
		rZ1hn2aV := yQaDA3L1 + ba0sUoWL

		blKMeD2 := rZ1hn2aV - iu5VUy1y0GT.fKgy76T6uU
		iu5VUy1y0GT.fKgy76T6uU = rZ1hn2aV

		return blKMeD2, nil
	}
	return 0, nil
}

func mR9SfM1gvky(dk_HEZvJo72Q int) gasFunc {
	return func(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {

		a0SqxOtD, v1p7xkTykQgN := /*line EeD5aeW1crB.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
		if v1p7xkTykQgN != nil {
			return 0, v1p7xkTykQgN
		}

		c6Qppz, r9uWm2pY := /*line IBrEbBx.go:1*/ eV1ooMo1qAI.Back(dk_HEZvJo72Q).Uint64WithOverflow()
		if r9uWm2pY {
			return 0, QiOgwB7BT
		}

		if c6Qppz, r9uWm2pY = /*line dxq4IJSQ.go:1*/ math.SafeMul( /*line l_s17n5Ura5K.go:1*/ f6N2wznhW(c6Qppz), params.CopyGas); r9uWm2pY {
			return 0, QiOgwB7BT
		}

		if a0SqxOtD, r9uWm2pY = /*line Al_yN5usz.go:1*/ math.SafeAdd(a0SqxOtD, c6Qppz); r9uWm2pY {
			return 0, QiOgwB7BT
		}
		return a0SqxOtD, nil
	}
}

var (
	uFCmbjA   = /*line kcd08ibw8aYP.go:1*/ mR9SfM1gvky(2)
	eaHmd75RS = /*line RQFrnACGWq.go:1*/ mR9SfM1gvky(2)
	dNvDvP    = /*line G20N4i8ND.go:1*/ mR9SfM1gvky(2)
	jgfKmr    = /*line PhLLiu84.go:1*/ mR9SfM1gvky(3)
	g8KCX7    = /*line Qd9HLCqcNOUJ.go:1*/ mR9SfM1gvky(2)
)

func i7tr1rh(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	var (
		eK4i7p94C79, hpgMKBm = /*line cgYlkxmfw0vD.go:1*/ eV1ooMo1qAI.Back(1) /*line Pya5Q6hPrx9.go:1*/, eV1ooMo1qAI.Back(0)
		ia6V464EksUP         = /*line WQ35c98h.go:1*/ oqIE3EcmK.StateDB.GetState( /*line wEO1ojDE.go:1*/ nsYSMon3US.Address() /*line HijZXCPnxFiH.go:1*/, hpgMKBm.Bytes32())
	)

	if oqIE3EcmK.chainRules.IsPetersburg || !oqIE3EcmK.chainRules.IsConstantinople {

		switch {
		case ia6V464EksUP == (common.Hash{}) && /*line x8Z_U2Y2Ze.go:1*/ eK4i7p94C79.Sign() != 0:
			return params.SstoreSetGas, nil
		case ia6V464EksUP != (common.Hash{}) && /*line mxanE0daQ.go:1*/ eK4i7p94C79.Sign() == 0:
			/*line eDSFnODQD.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.SstoreRefundGas)
			return params.SstoreClearGas, nil
		default:
			return params.SstoreResetGas, nil
		}
	}

	gVUwww := /*line uq94qIt.go:1*/ common.Hash( /*line P27BgyOO.go:1*/ eK4i7p94C79.Bytes32())
	if ia6V464EksUP == gVUwww {
		return params.NetSstoreNoopGas, nil
	}
	qltkB7G8Emox := /*line YdLnHVrh_K.go:1*/ oqIE3EcmK.StateDB.GetCommittedState( /*line kkQYopayu.go:1*/ nsYSMon3US.Address() /*line KsSET8.go:1*/, hpgMKBm.Bytes32())
	if qltkB7G8Emox == ia6V464EksUP {
		if qltkB7G8Emox == (common.Hash{}) {
			return params.NetSstoreInitGas, nil
		}
		if gVUwww == (common.Hash{}) {
			/*line D5gIJNk.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.NetSstoreClearRefund)
		}
		return params.NetSstoreCleanGas, nil
	}
	if qltkB7G8Emox != (common.Hash{}) {
		if ia6V464EksUP == (common.Hash{}) {
			/*line ndazFuvWGCR.go:1*/ oqIE3EcmK.StateDB.SubRefund(params.NetSstoreClearRefund)
		} else if gVUwww == (common.Hash{}) {
			/*line DZHHZJN6YMu.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.NetSstoreClearRefund)
		}
	}
	if qltkB7G8Emox == gVUwww {
		if qltkB7G8Emox == (common.Hash{}) {
			/*line JnbRahC.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.NetSstoreResetClearRefund)
		} else {
			/*line CoMiWe.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.NetSstoreResetRefund)
		}
	}
	return params.NetSstoreDirtyGas, nil
}

func r8YSTJBT6L7(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {

	if nsYSMon3US.OzrTJ12T <= params.SstoreSentryGasEIP2200 {
		return 0 /*line gHkpS9.go:1*/, errors.New(func() /*line gAzJTphroqVL.go:1*/ string {
			key := [] /*line gAzJTphroqVL.go:1*/ byte("\x12\x1a\x82\x1b8\xd6\x18&=qF!\xc6T\xbc<\xe4\xc2\v\x8a<a/\xf2\xa1\xf9\x8c\xa0\x0fH}R(\xa1\xe5\x92")
			data := [] /*line gAzJTphroqVL.go:1*/ byte("\\U\xf2\x05-\x98WO*\xf7\xdaF\x9b\x1fd*\x8b\xb0\x15\xe8)\x04?\x82\xd1h\xe2\xc3j\xd8\xf6\x13FӍ\xe7")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line gAzJTphroqVL.go:1*/ string(data)
		}())
	}

	var (
		eK4i7p94C79, hpgMKBm = /*line rDKcwK7jr.go:1*/ eV1ooMo1qAI.Back(1) /*line XdAPnJP.go:1*/, eV1ooMo1qAI.Back(0)
		ia6V464EksUP         = /*line hLaKTDs4h.go:1*/ oqIE3EcmK.StateDB.GetState( /*line VnNqwWWmK6.go:1*/ nsYSMon3US.Address() /*line Ir2IOKt.go:1*/, hpgMKBm.Bytes32())
	)
	gVUwww := /*line swkvTB.go:1*/ common.Hash( /*line iIkdp26.go:1*/ eK4i7p94C79.Bytes32())

	if ia6V464EksUP == gVUwww {
		return params.SloadGasEIP2200, nil
	}
	qltkB7G8Emox := /*line nClBOHcmPn.go:1*/ oqIE3EcmK.StateDB.GetCommittedState( /*line WxaCjZFOCD.go:1*/ nsYSMon3US.Address() /*line OeB947jC.go:1*/, hpgMKBm.Bytes32())
	if qltkB7G8Emox == ia6V464EksUP {
		if qltkB7G8Emox == (common.Hash{}) {
			return params.SstoreSetGasEIP2200, nil
		}
		if gVUwww == (common.Hash{}) {
			/*line uXKJn_E8w5Er.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.SstoreClearsScheduleRefundEIP2200)
		}
		return params.SstoreResetGasEIP2200, nil
	}
	if qltkB7G8Emox != (common.Hash{}) {
		if ia6V464EksUP == (common.Hash{}) {
			/*line n8vnQRYuQgPr.go:1*/ oqIE3EcmK.StateDB.SubRefund(params.SstoreClearsScheduleRefundEIP2200)
		} else if gVUwww == (common.Hash{}) {
			/*line T5JXZRKVL.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.SstoreClearsScheduleRefundEIP2200)
		}
	}
	if qltkB7G8Emox == gVUwww {
		if qltkB7G8Emox == (common.Hash{}) {
			/*line GhuZO4.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.SstoreSetGasEIP2200 - params.SloadGasEIP2200)
		} else {
			/*line aXDWvwKFzXBJ.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.SstoreResetGasEIP2200 - params.SloadGasEIP2200)
		}
	}
	return params.SloadGasEIP2200, nil
}

func kUKsmTBT(iyrnVpBe8CCl uint64) gasFunc {
	return func(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
		v1tagXR, r9uWm2pY := /*line VKCOjvztZK.go:1*/ eV1ooMo1qAI.Back(1).Uint64WithOverflow()
		if r9uWm2pY {
			return 0, QiOgwB7BT
		}

		a0SqxOtD, v1p7xkTykQgN := /*line sTLVYdaI.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
		if v1p7xkTykQgN != nil {
			return 0, v1p7xkTykQgN
		}

		if a0SqxOtD, r9uWm2pY = /*line yL9r52uFNl.go:1*/ math.SafeAdd(a0SqxOtD, params.LogGas); r9uWm2pY {
			return 0, QiOgwB7BT
		}
		if a0SqxOtD, r9uWm2pY = /*line JydLD9Ff_kD.go:1*/ math.SafeAdd(a0SqxOtD, iyrnVpBe8CCl*params.LogTopicGas); r9uWm2pY {
			return 0, QiOgwB7BT
		}

		var tUyZrnJX uint64
		if tUyZrnJX, r9uWm2pY = /*line AmW50gWJGBvy.go:1*/ math.SafeMul(v1tagXR, params.LogDataGas); r9uWm2pY {
			return 0, QiOgwB7BT
		}
		if a0SqxOtD, r9uWm2pY = /*line BTa2XVJ2.go:1*/ math.SafeAdd(a0SqxOtD, tUyZrnJX); r9uWm2pY {
			return 0, QiOgwB7BT
		}
		return a0SqxOtD, nil
	}
}

func jNJPVlmt(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a0SqxOtD, v1p7xkTykQgN := /*line kRjaIyQ0xPM.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	zw2dfNv, r9uWm2pY := /*line MzU5J0p5.go:1*/ eV1ooMo1qAI.Back(1).Uint64WithOverflow()
	if r9uWm2pY {
		return 0, QiOgwB7BT
	}
	if zw2dfNv, r9uWm2pY = /*line xh4aL1fpkyT.go:1*/ math.SafeMul( /*line ftUMRJPhO4j.go:1*/ f6N2wznhW(zw2dfNv), params.Keccak256WordGas); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	if a0SqxOtD, r9uWm2pY = /*line FQUfMQq6q.go:1*/ math.SafeAdd(a0SqxOtD, zw2dfNv); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func vdC38M(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	return /*line aIVUUvQoE.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
}

var (
	n9aa81       = vdC38M
	gt2cFYPNW    = vdC38M
	mqEO_xnUyuwM = vdC38M
	jUCdrJVETKmE = vdC38M
	yfEk_0hk     = vdC38M
	xaxmYoxeKP   = vdC38M
)

func zBgFdndQ(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a0SqxOtD, v1p7xkTykQgN := /*line tZTzJYxoNSoL.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	zw2dfNv, r9uWm2pY := /*line HAGzbSD7O.go:1*/ eV1ooMo1qAI.Back(2).Uint64WithOverflow()
	if r9uWm2pY {
		return 0, QiOgwB7BT
	}
	if zw2dfNv, r9uWm2pY = /*line a6b7rSxep.go:1*/ math.SafeMul( /*line V1008tYjWIUn.go:1*/ f6N2wznhW(zw2dfNv), params.Keccak256WordGas); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	if a0SqxOtD, r9uWm2pY = /*line rf8CGXDTA.go:1*/ math.SafeAdd(a0SqxOtD, zw2dfNv); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func hDQ5MD(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a0SqxOtD, v1p7xkTykQgN := /*line ccxOxxm77rG.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	jLa0ywd, r9uWm2pY := /*line KyMu4x2.go:1*/ eV1ooMo1qAI.Back(2).Uint64WithOverflow()
	if r9uWm2pY || jLa0ywd > params.MaxInitCodeSize {
		return 0, QiOgwB7BT
	}

	eX82VCSk := params.InitCodeWordGas * ((jLa0ywd + 31) / 32)
	if a0SqxOtD, r9uWm2pY = /*line UMB01x.go:1*/ math.SafeAdd(a0SqxOtD, eX82VCSk); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}
func dUKWTgDCL(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a0SqxOtD, v1p7xkTykQgN := /*line QLzjf8.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	jLa0ywd, r9uWm2pY := /*line B2nrrO5nD.go:1*/ eV1ooMo1qAI.Back(2).Uint64WithOverflow()
	if r9uWm2pY || jLa0ywd > params.MaxInitCodeSize {
		return 0, QiOgwB7BT
	}

	eX82VCSk := (params.InitCodeWordGas + params.Keccak256WordGas) * ((jLa0ywd + 31) / 32)
	if a0SqxOtD, r9uWm2pY = /*line OQ4Wl2XECbVI.go:1*/ math.SafeAdd(a0SqxOtD, eX82VCSk); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func pJOGLJAn(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	udP9Bu := /*line YuW1FgvO.go:1*/ uint64(( /*line wY2rwP.go:1*/ eV1ooMo1qAI.sBD8Un65JT[ /*line nlqIYTyb0Ir.go:1*/ eV1ooMo1qAI.ckaB4DI()-2].BitLen() + 7) / 8)

	var (
		a0SqxOtD = udP9Bu * params.ExpByteFrontier
		r9uWm2pY bool
	)
	if a0SqxOtD, r9uWm2pY = /*line xaG_w2.go:1*/ math.SafeAdd(a0SqxOtD, params.ExpGas); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func cpcGfh(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	udP9Bu := /*line F4PBxb.go:1*/ uint64(( /*line nKDRh_N.go:1*/ eV1ooMo1qAI.sBD8Un65JT[ /*line MTanNY.go:1*/ eV1ooMo1qAI.ckaB4DI()-2].BitLen() + 7) / 8)

	var (
		a0SqxOtD = udP9Bu * params.ExpByteEIP158
		r9uWm2pY bool
	)
	if a0SqxOtD, r9uWm2pY = /*line pKzemUSjN.go:1*/ math.SafeAdd(a0SqxOtD, params.ExpGas); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func lToOdJttDF(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	var (
		a0SqxOtD uint64
		ylCZTV8y = ! /*line tb23z89kB_L.go:1*/ eV1ooMo1qAI.Back(2).IsZero()
		qtbv6eO  = /*line Et8a4MN_E.go:1*/ common.Address( /*line PkBma9S.go:1*/ eV1ooMo1qAI.Back(1).Bytes20())
	)
	if oqIE3EcmK.chainRules.IsEIP158 {
		if ylCZTV8y && /*line im7WeJ6H.go:1*/ oqIE3EcmK.StateDB.Empty(qtbv6eO) {
			a0SqxOtD += params.CallNewAccountGas
		}
	} else if ! /*line jP5iF3Ka2a4.go:1*/ oqIE3EcmK.StateDB.Exist(qtbv6eO) {
		a0SqxOtD += params.CallNewAccountGas
	}
	if ylCZTV8y {
		a0SqxOtD += params.CallValueTransferGas
	}
	a4uYoMHWG, v1p7xkTykQgN := /*line AEeOOmc.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	var r9uWm2pY bool
	if a0SqxOtD, r9uWm2pY = /*line GIutisollz.go:1*/ math.SafeAdd(a0SqxOtD, a4uYoMHWG); r9uWm2pY {
		return 0, QiOgwB7BT
	}

	oqIE3EcmK.callGasTemp, v1p7xkTykQgN = /*line OkT7D6.go:1*/ iHblnaC3VG(oqIE3EcmK.chainRules.IsEIP150, nsYSMon3US.OzrTJ12T, a0SqxOtD /*line kiBi67.go:1*/, eV1ooMo1qAI.Back(0))
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	if a0SqxOtD, r9uWm2pY = /*line GffYcylyS.go:1*/ math.SafeAdd(a0SqxOtD, oqIE3EcmK.callGasTemp); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func v_artS8KR_h(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a4uYoMHWG, v1p7xkTykQgN := /*line XdQz_JoqZnR.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	var (
		a0SqxOtD uint64
		r9uWm2pY bool
	)
	if /*line q59IGnI.go:1*/ eV1ooMo1qAI.Back(2).Sign() != 0 {
		a0SqxOtD += params.CallValueTransferGas
	}
	if a0SqxOtD, r9uWm2pY = /*line s639qZvrAHTo.go:1*/ math.SafeAdd(a0SqxOtD, a4uYoMHWG); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	oqIE3EcmK.callGasTemp, v1p7xkTykQgN = /*line p9AHMbritBB.go:1*/ iHblnaC3VG(oqIE3EcmK.chainRules.IsEIP150, nsYSMon3US.OzrTJ12T, a0SqxOtD /*line k4YwtL1Z6rP.go:1*/, eV1ooMo1qAI.Back(0))
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	if a0SqxOtD, r9uWm2pY = /*line nrS5XoM4JX0o.go:1*/ math.SafeAdd(a0SqxOtD, oqIE3EcmK.callGasTemp); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func kQ1IHor(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a0SqxOtD, v1p7xkTykQgN := /*line L5AfYh.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	oqIE3EcmK.callGasTemp, v1p7xkTykQgN = /*line IX06LeJEco.go:1*/ iHblnaC3VG(oqIE3EcmK.chainRules.IsEIP150, nsYSMon3US.OzrTJ12T, a0SqxOtD /*line NGMX8r2qB.go:1*/, eV1ooMo1qAI.Back(0))
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	var r9uWm2pY bool
	if a0SqxOtD, r9uWm2pY = /*line RqrJkMA.go:1*/ math.SafeAdd(a0SqxOtD, oqIE3EcmK.callGasTemp); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func svRiUsX(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	a0SqxOtD, v1p7xkTykQgN := /*line fi3oz9m.go:1*/ sssg6yQAr(iu5VUy1y0GT, umeiVGSX__Q)
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	oqIE3EcmK.callGasTemp, v1p7xkTykQgN = /*line FESvJrfIFi.go:1*/ iHblnaC3VG(oqIE3EcmK.chainRules.IsEIP150, nsYSMon3US.OzrTJ12T, a0SqxOtD /*line dnfuaUvzrB.go:1*/, eV1ooMo1qAI.Back(0))
	if v1p7xkTykQgN != nil {
		return 0, v1p7xkTykQgN
	}
	var r9uWm2pY bool
	if a0SqxOtD, r9uWm2pY = /*line LCoDjPfZ.go:1*/ math.SafeAdd(a0SqxOtD, oqIE3EcmK.callGasTemp); r9uWm2pY {
		return 0, QiOgwB7BT
	}
	return a0SqxOtD, nil
}

func tLUfnFIU5X(oqIE3EcmK *EVM, nsYSMon3US *WyJ004I8, eV1ooMo1qAI *IZHSJVjaAmxb, iu5VUy1y0GT *VfNB1WFtybE, umeiVGSX__Q uint64) (uint64, error) {
	var a0SqxOtD uint64

	if oqIE3EcmK.chainRules.IsEIP150 {
		a0SqxOtD = params.SelfdestructGasEIP150
		var qtbv6eO = /*line KMMRA_z.go:1*/ common.Address( /*line Z0hdI90gmM4.go:1*/ eV1ooMo1qAI.Back(0).Bytes20())

		if oqIE3EcmK.chainRules.IsEIP158 {

			if /*line XFEojxdK.go:1*/ oqIE3EcmK.StateDB.Empty(qtbv6eO) && /*line eY9Fdd.go:1*/ oqIE3EcmK.StateDB.GetBalance( /*line HqhMrBnQ10.go:1*/ nsYSMon3US.Address()).Sign() != 0 {
				a0SqxOtD += params.CreateBySelfdestructGas
			}
		} else if ! /*line _FHbrU1.go:1*/ oqIE3EcmK.StateDB.Exist(qtbv6eO) {
			a0SqxOtD += params.CreateBySelfdestructGas
		}
	}

	if ! /*line ESqafr8q.go:1*/ oqIE3EcmK.StateDB.HasSelfDestructed( /*line IWV2PnaVw.go:1*/ nsYSMon3US.Address()) {
		/*line FqgjfHS5Q.go:1*/ oqIE3EcmK.StateDB.AddRefund(params.SelfdestructRefundGas)
	}
	return a0SqxOtD, nil
}

var _ = errors.As
var _ = common.AbsolutePath
var _ = math.BigMax
var _ = params.AllCliqueProtocolChanges
