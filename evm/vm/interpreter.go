//line :1
package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type Config struct {
	Tracer                  EVMLogger
	NoBaseFee               bool
	EnablePreimageRecording bool
	ExtraEips               []int
}

type GUBKEu339CA struct {
	IWhp7ux0fQ *VfNB1WFtybE
	NCfGFD     *IZHSJVjaAmxb
	YlbzoB     *WyJ004I8
}

type EVMInterpreter struct {
	evm   *EVM
	table *JumpTable

	hasher    crypto.KeccakState
	hasherBuf common.Hash

	readOnly   bool
	returnData []byte
}

func FjkJz0a0R(oqIE3EcmK *EVM) *EVMInterpreter {

	var gca1Eh *JumpTable
	switch {
	case oqIE3EcmK.chainRules.IsCancun:
		gca1Eh = &hEiTFfkGV_
	case oqIE3EcmK.chainRules.IsShanghai:
		gca1Eh = &ecLGn5YUA_
	case oqIE3EcmK.chainRules.IsMerge:
		gca1Eh = &f1EQuA
	case oqIE3EcmK.chainRules.IsLondon:
		gca1Eh = &myuFbYmfOo7
	case oqIE3EcmK.chainRules.IsBerlin:
		gca1Eh = &l7p1GOaGYiab
	case oqIE3EcmK.chainRules.IsIstanbul:
		gca1Eh = &cslDH5rwHrW2
	case oqIE3EcmK.chainRules.IsConstantinople:
		gca1Eh = &ygkeaMUJ
	case oqIE3EcmK.chainRules.IsByzantium:
		gca1Eh = &eVTfqUZ5XLKp
	case oqIE3EcmK.chainRules.IsEIP158:
		gca1Eh = &eTszvRda
	case oqIE3EcmK.chainRules.IsEIP150:
		gca1Eh = &iqRN8fb
	case oqIE3EcmK.chainRules.IsHomestead:
		gca1Eh = &slTZ7K
	default:
		gca1Eh = &zi10f_aku1a
	}
	var jSiVtfEZz []int
	if /*line NfYuJ8xw4.go:1*/ len(oqIE3EcmK.Config.ExtraEips) > 0 {

		gca1Eh = /*line KzkERE.go:1*/ fW9OT34(gca1Eh)
	}
	for _, a_wL6Q2V := range oqIE3EcmK.Config.ExtraEips {
		if v1p7xkTykQgN := /*line ZNEW5Xsl_lXp.go:1*/ RucRpvC(a_wL6Q2V, gca1Eh); v1p7xkTykQgN != nil {

			/*line a7jCmvPmEKmk.go:1*/
			log.Error(func() /*line O8X96Q4m3.go:1*/ string {
				key := [] /*line O8X96Q4m3.go:1*/ byte("\xbb\x0f]u\xbc\xe1;x\x9c\x02L\x97\x80):\x99\xa5\xfa\baf")
				data := [] /*line O8X96Q4m3.go:1*/ byte("\xfeF\rU݂O\x11\xeac8\xfe\xefG\x1a\xffēd\x04\x02")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line O8X96Q4m3.go:1*/ string(data)
			}(), "eip", a_wL6Q2V, "error", v1p7xkTykQgN)
		} else {
			jSiVtfEZz = /*line NnDCnpDyw.go:1*/ append(jSiVtfEZz, a_wL6Q2V)
		}
	}
	oqIE3EcmK.Config.ExtraEips = jSiVtfEZz
	return &EVMInterpreter{evm: oqIE3EcmK, table: gca1Eh}
}

func (kTH29c *EVMInterpreter) Run(nsYSMon3US *WyJ004I8, sIVM3RLPSKSh []byte, nIwSY3 bool) (ag8Tdqxi4 []byte, v1p7xkTykQgN error) {

	kTH29c.evm.depth++
	defer func() { /*line s795dW.go:1*/ kTH29c.evm.depth-- }()

	if nIwSY3 && !kTH29c.readOnly {
		kTH29c.readOnly = true
		defer func() { /*line wgrcDGggj.go:1*/ kTH29c.readOnly = false }()
	}

	kTH29c.returnData = nil

	if /*line AfrWEiXD5zaL.go:1*/ len(nsYSMon3US.STabmC) == 0 {
		return nil, nil
	}

	var (
		nCnJ0mc      LxosJe8
		iu5VUy1y0GT  = /*line dG0VB93.go:1*/ NRDMEb()
		eV1ooMo1qAI  = /*line dAL8ASJc.go:1*/ krqDUTXe()
		ahYrRb0Iw7v0 = &GUBKEu339CA{
			IWhp7ux0fQ: iu5VUy1y0GT,
			NCfGFD:     eV1ooMo1qAI,
			YlbzoB:     nsYSMon3US,
		}

		eAorajeVy084 = /*line QlUhn7a_mV.go:1*/ uint64(0)
		pIJBYyGRV    uint64

		eazPgic uint64
		g732GW  uint64
		aVqNrYS bool
		ekzSsj9 []byte
		hrNtUBI = kTH29c.evm.Config.Tracer != nil
	)

	defer func() {
		/*line DDHstQHY5Sm_.go:1*/ gGwn14iwA2Q_(eV1ooMo1qAI)
	}()
	nsYSMon3US.AdmClN5HlL = sIVM3RLPSKSh

	if hrNtUBI {
		defer func() {
			if /*line CjrNkr.go:1*/ v1p7xkTykQgN != nil {
				if !aVqNrYS {
					/*line rTQLufXR6u.go:1*/ kTH29c.evm.Config.Tracer.CaptureState(eazPgic, nCnJ0mc, g732GW, pIJBYyGRV, ahYrRb0Iw7v0, kTH29c.returnData, kTH29c.evm.depth, v1p7xkTykQgN)
				} else {
					/*line FfOjpj.go:1*/ kTH29c.evm.Config.Tracer.CaptureFault(eazPgic, nCnJ0mc, g732GW, pIJBYyGRV, ahYrRb0Iw7v0, kTH29c.evm.depth, v1p7xkTykQgN)
				}
			}
		}()
	}

	for {
		if hrNtUBI {

			aVqNrYS, eazPgic, g732GW = false, eAorajeVy084, nsYSMon3US.OzrTJ12T
		}

		nCnJ0mc = /*line HQoXSh.go:1*/ nsYSMon3US.GetOp(eAorajeVy084)
		mZQXEEsTFNg := kTH29c.table[nCnJ0mc]
		pIJBYyGRV = mZQXEEsTFNg.constantGas

		if bgNn4q8 := /*line vgueXjEkpMg.go:1*/ eV1ooMo1qAI.ckaB4DI(); bgNn4q8 < mZQXEEsTFNg.minStack {
			return nil, &EqA0VnW9Ok{orXbLSPr2nk: bgNn4q8, jjPUbtbD: mZQXEEsTFNg.minStack}
		} else if bgNn4q8 > mZQXEEsTFNg.maxStack {
			return nil, &FiCSfN{iJNQHm: bgNn4q8, tqEGFKAM: mZQXEEsTFNg.maxStack}
		}
		if ! /*line Dgq62ikW9jP8.go:1*/ nsYSMon3US.UseGas(pIJBYyGRV) {
			return nil, ME5lmy
		}
		if mZQXEEsTFNg.dynamicGas != nil {

			var umeiVGSX__Q uint64

			if mZQXEEsTFNg.memorySize != nil {
				n5ZF74O, r9uWm2pY := /*line NzPk1GT8C.go:1*/ mZQXEEsTFNg.memorySize(eV1ooMo1qAI)
				if r9uWm2pY {
					return nil, QiOgwB7BT
				}

				if umeiVGSX__Q, r9uWm2pY = /*line VdSdFRtww.go:1*/ math.SafeMul( /*line I50bZJ8b.go:1*/ f6N2wznhW(n5ZF74O), 32); r9uWm2pY {
					return nil, QiOgwB7BT
				}
			}

			var tZL_cc9x_vP9 uint64
			tZL_cc9x_vP9, v1p7xkTykQgN = /*line Dsmwvre0s.go:1*/ mZQXEEsTFNg.dynamicGas(kTH29c.evm, nsYSMon3US, eV1ooMo1qAI, iu5VUy1y0GT, umeiVGSX__Q)
			pIJBYyGRV += tZL_cc9x_vP9
			if v1p7xkTykQgN != nil || ! /*line RK_BHLI8F.go:1*/ nsYSMon3US.UseGas(tZL_cc9x_vP9) {
				return nil, ME5lmy
			}

			if hrNtUBI {
				/*line ZS5e8Xab.go:1*/ kTH29c.evm.Config.Tracer.CaptureState(eAorajeVy084, nCnJ0mc, g732GW, pIJBYyGRV, ahYrRb0Iw7v0, kTH29c.returnData, kTH29c.evm.depth, v1p7xkTykQgN)
				aVqNrYS = true
			}
			if umeiVGSX__Q > 0 {
				/*line C9F86VhzB.go:1*/ iu5VUy1y0GT.Resize(umeiVGSX__Q)
			}
		} else if hrNtUBI {
			/*line HsVvjHD5a.go:1*/ kTH29c.evm.Config.Tracer.CaptureState(eAorajeVy084, nCnJ0mc, g732GW, pIJBYyGRV, ahYrRb0Iw7v0, kTH29c.returnData, kTH29c.evm.depth, v1p7xkTykQgN)
			aVqNrYS = true
		}

		ekzSsj9, v1p7xkTykQgN = /*line AZPCL2uy4Vvw.go:1*/ mZQXEEsTFNg.execute(&eAorajeVy084, kTH29c, ahYrRb0Iw7v0)
		if v1p7xkTykQgN != nil {
			break
		}
		eAorajeVy084++
	}

	if v1p7xkTykQgN == deyJ2RBO4L {
		v1p7xkTykQgN = nil
	}

	return ekzSsj9, v1p7xkTykQgN
}

var _ = common.AbsolutePath
var _ = math.BigMax
var _ = crypto.CompressPubkey
var _ = log.Crit
