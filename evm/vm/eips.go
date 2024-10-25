//line :1
package vm

import (
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

var rEgwHroU = map[int]func(*JumpTable){
	5656: ghNjp4tedh,
	6780: eC3GlL4,
	3855: wKpQXZpUwK,
	3860: wuw9VQEHA6,
	3529: dHGM24Ut0F,
	3198: hgndeOut9ePa,
	2929: wIyo_cTzT4,
	2200: e4gyUrMOn3L,
	1884: bIhOZSNloD,
	1344: ntMBqQpN5sa,
	1153: agVtvxDSUs,
}

func RucRpvC(gLddIqU int, kOAWieB3PO8 *JumpTable) error {
	fc1qAAhlLac8, m5pPAO := rEgwHroU[gLddIqU]
	if !m5pPAO {
		return /*line gssUAc2joaHH.go:1*/ fmt.Errorf(func() /*line a0YRT7U.go:1*/ string {
			data := /*line a0YRT7U.go:1*/ make([]byte, 0, 17)
			i := 5
			decryptKey := 123
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data = /*line a0YRT7U.go:1*/ append(data, 62)
					i = 3
				case 2:
					data = /*line a0YRT7U.go:1*/ append(data, " qw7"...,
					)
					i = 6
				case 3:
					i = 0
					data = /*line a0YRT7U.go:1*/ append(data, "0u"...,
					)
				case 6:
					for y := range data {
						data[y] = data[y] ^ /*line a0YRT7U.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 9:
					data = /*line a0YRT7U.go:1*/ append(data, 48)
					i = 7
				case 1:
					i = 9
					data = /*line a0YRT7U.go:1*/ append(data, "3::>"...,
					)
				case 5:
					i = 1
					data = /*line a0YRT7U.go:1*/ append(data, 41)
				case 7:
					data = /*line a0YRT7U.go:1*/ append(data, 52)
					i = 4
				case 0:
					i = 2
					data = /*line a0YRT7U.go:1*/ append(data, "3>"...,
					)
				}
			}
			return /*line a0YRT7U.go:1*/ string(data)
		}(), gLddIqU)
	}
	/*line PfOr_v8i7.go:1*/ fc1qAAhlLac8(kOAWieB3PO8)
	return nil
}

func Ko1f8yFYO(gLddIqU int) bool {
	_, m5pPAO := rEgwHroU[gLddIqU]
	return m5pPAO
}
func IASUT9nDKtFI() []string {
	var chn1DWE []string
	for zsVIRN6k := range rEgwHroU {
		chn1DWE = /*line fXeIMJ.go:1*/ append(chn1DWE /*line fg5oEtN7.go:1*/, fmt.Sprintf("%d", zsVIRN6k))
	}
	/*line mKHZS8VFP.go:1*/ sort.Strings(chn1DWE)
	return chn1DWE
}

func bIhOZSNloD(kOAWieB3PO8 *JumpTable) {

	kOAWieB3PO8[SLOAD].constantGas = params.SloadGasEIP1884
	kOAWieB3PO8[BALANCE].constantGas = params.BalanceGasEIP1884
	kOAWieB3PO8[EXTCODEHASH].constantGas = params.ExtcodeHashGasEIP1884

	kOAWieB3PO8[SELFBALANCE] = &operation{
		execute:     zcHineC21,
		constantGas: GasFastStep,
		minStack:/*line q2NikY.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line jaH_BlvQr7D.go:1*/ mETN8x(0, 1),
	}
}

func zcHineC21(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	coqYjCIx := /*line D9zSNIdZA.go:1*/ i6HsYcs_9RSM.evm.StateDB.GetBalance( /*line qBuBia.go:1*/ oXFAbeW4qW.YlbzoB.Address())
	/*line GK_jSC9F6K.go:1*/ oXFAbeW4qW.NCfGFD.p69kwi(coqYjCIx)
	return nil, nil
}

func ntMBqQpN5sa(kOAWieB3PO8 *JumpTable) {

	kOAWieB3PO8[CHAINID] = &operation{
		execute:     bYiTeEy,
		constantGas: GasQuickStep,
		minStack:/*line U3U0Sf.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line mEg81eAaqZR.go:1*/ mETN8x(0, 1),
	}
}

func bYiTeEy(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	hbk9nZvPouh5, _ := /*line zXZtAG7.go:1*/ uint256.FromBig(i6HsYcs_9RSM.evm.chainConfig.ChainID)
	/*line HjsjVqmC5taY.go:1*/ oXFAbeW4qW.NCfGFD.p69kwi(hbk9nZvPouh5)
	return nil, nil
}

func e4gyUrMOn3L(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[SLOAD].constantGas = params.SloadGasEIP2200
	kOAWieB3PO8[SSTORE].dynamicGas = r8YSTJBT6L7
}

func wIyo_cTzT4(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[SSTORE].dynamicGas = eyogaBFMJnN

	kOAWieB3PO8[SLOAD].constantGas = 0
	kOAWieB3PO8[SLOAD].dynamicGas = xo7aSal

	kOAWieB3PO8[EXTCODECOPY].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[EXTCODECOPY].dynamicGas = dmS0QyHL3jW

	kOAWieB3PO8[EXTCODESIZE].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[EXTCODESIZE].dynamicGas = wgZqJIOJy_

	kOAWieB3PO8[EXTCODEHASH].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[EXTCODEHASH].dynamicGas = wgZqJIOJy_

	kOAWieB3PO8[BALANCE].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[BALANCE].dynamicGas = wgZqJIOJy_

	kOAWieB3PO8[CALL].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[CALL].dynamicGas = o6YZbUtLF

	kOAWieB3PO8[CALLCODE].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[CALLCODE].dynamicGas = ebMHFr2anx3

	kOAWieB3PO8[STATICCALL].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[STATICCALL].dynamicGas = y8J5Tzv

	kOAWieB3PO8[DELEGATECALL].constantGas = params.WarmStorageReadCostEIP2929
	kOAWieB3PO8[DELEGATECALL].dynamicGas = wm3iLpdO

	kOAWieB3PO8[SELFDESTRUCT].constantGas = params.SelfdestructGasEIP150
	kOAWieB3PO8[SELFDESTRUCT].dynamicGas = iIZATkPa
}

func dHGM24Ut0F(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[SSTORE].dynamicGas = m3aQyqq
	kOAWieB3PO8[SELFDESTRUCT].dynamicGas = ec6eoV
}

func hgndeOut9ePa(kOAWieB3PO8 *JumpTable) {

	kOAWieB3PO8[BASEFEE] = &operation{
		execute:     gM5gy9o,
		constantGas: GasQuickStep,
		minStack:/*line Cg8UfMi.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line H6RS4Qni3kr.go:1*/ mETN8x(0, 1),
	}
}

func agVtvxDSUs(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[TLOAD] = &operation{
		execute:     aAIlWeqq,
		constantGas: params.WarmStorageReadCostEIP2929,
		minStack:/*line eLwFFEEf.go:1*/ d0cam_EE09tL(1, 1),
		maxStack:/*line vFmOcTeKRiE.go:1*/ mETN8x(1, 1),
	}

	kOAWieB3PO8[TSTORE] = &operation{
		execute:     fNSI_8,
		constantGas: params.WarmStorageReadCostEIP2929,
		minStack:/*line ASW5hF6zR.go:1*/ d0cam_EE09tL(2, 0),
		maxStack:/*line D64ZcfpT8.go:1*/ mETN8x(2, 0),
	}
}

func aAIlWeqq(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	ia4Vzxc4yCMp := /*line mtRUDGsV.go:1*/ oXFAbeW4qW.NCfGFD.bx3dXQEGk()
	vvL5SsAl := /*line Na1rlctA.go:1*/ common.Hash( /*line El36soIl.go:1*/ ia4Vzxc4yCMp.Bytes32())
	rVFSadsv := /*line kieWZZfda1.go:1*/ i6HsYcs_9RSM.evm.StateDB.GetTransientState( /*line DEC9nfCWs.go:1*/ oXFAbeW4qW.YlbzoB.Address(), vvL5SsAl)
	/*line _L1P3lTM.go:1*/ ia4Vzxc4yCMp.SetBytes( /*line JfwUmIuJqAC7.go:1*/ rVFSadsv.Bytes())
	return nil, nil
}

func fNSI_8(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	if i6HsYcs_9RSM.readOnly {
		return nil, FgcTjgE_LXA
	}
	ia4Vzxc4yCMp := /*line EPUNHx.go:1*/ oXFAbeW4qW.NCfGFD.bIh0YHTaQv2I()
	rVFSadsv := /*line p4p9GyOc.go:1*/ oXFAbeW4qW.NCfGFD.bIh0YHTaQv2I()
	/*line AyGUkJ3CwW.go:1*/ i6HsYcs_9RSM.evm.StateDB.SetTransientState( /*line qAqtjT0D.go:1*/ oXFAbeW4qW.YlbzoB.Address() /*line AcaOMRvw.go:1*/, ia4Vzxc4yCMp.Bytes32() /*line VLCCgDfVCPX7.go:1*/, rVFSadsv.Bytes32())
	return nil, nil
}

func gM5gy9o(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	yXruNkhqX, _ := /*line CT1FRV.go:1*/ uint256.FromBig(i6HsYcs_9RSM.evm.Context.BaseFee)
	/*line K7QQKtINg.go:1*/ oXFAbeW4qW.NCfGFD.p69kwi(yXruNkhqX)
	return nil, nil
}

func wKpQXZpUwK(kOAWieB3PO8 *JumpTable) {

	kOAWieB3PO8[PUSH0] = &operation{
		execute:     nSKpPyAd,
		constantGas: GasQuickStep,
		minStack:/*line yPxhpKcnArja.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line XiaTSwBpRn_v.go:1*/ mETN8x(0, 1),
	}
}

func nSKpPyAd(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	/*line FycADuzY.go:1*/ oXFAbeW4qW.NCfGFD.p69kwi( /*line GFARhYF.go:1*/ new(uint256.Int))
	return nil, nil
}

func wuw9VQEHA6(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[CREATE].dynamicGas = hDQ5MD
	kOAWieB3PO8[CREATE2].dynamicGas = dUKWTgDCL
}

func ghNjp4tedh(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[MCOPY] = &operation{
		execute:     tGHXMtATx,
		constantGas: GasFastestStep,
		dynamicGas:  dNvDvP,
		minStack:/*line yLkC67G1.go:1*/ d0cam_EE09tL(3, 0),
		maxStack:/*line CcJzgr3jKGji.go:1*/ mETN8x(3, 0),
		memorySize: d16ZZ0VJUl,
	}
}

func tGHXMtATx(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	var (
		rafqvqt     = /*line vFJbCNbZrh7.go:1*/ oXFAbeW4qW.NCfGFD.bIh0YHTaQv2I()
		nwiuVz9sYqY = /*line nba1pn3Wh.go:1*/ oXFAbeW4qW.NCfGFD.bIh0YHTaQv2I()
		kpLBYg      = /*line sZwNIl.go:1*/ oXFAbeW4qW.NCfGFD.bIh0YHTaQv2I()
	)

	/*line vka6NBoYpTVs.go:1*/
	oXFAbeW4qW.IWhp7ux0fQ.Copy( /*line B4euyoue70yd.go:1*/ rafqvqt.Uint64() /*line y_jxz6.go:1*/, nwiuVz9sYqY.Uint64() /*line oHntck46yb.go:1*/, kpLBYg.Uint64())
	return nil, nil
}

func x5XENHYIC(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	qPF5C99xK7I := /*line CRAf6W_Cc3U.go:1*/ oXFAbeW4qW.NCfGFD.bx3dXQEGk()
	if /*line _Ckbwc.go:1*/ qPF5C99xK7I.LtUint64( /*line rQJAZSaxe0q.go:1*/ uint64( /*line Hf_MKJ3V.go:1*/ len(i6HsYcs_9RSM.evm.TxContext.BlobHashes))) {
		sLxAOvR := i6HsYcs_9RSM.evm.TxContext.BlobHashes[ /*line EO_4_ozng.go:1*/ qPF5C99xK7I.Uint64()]
		/*line GCelWO7.go:1*/ qPF5C99xK7I.SetBytes32(sLxAOvR[:])
	} else {
		/*line wbatns.go:1*/ qPF5C99xK7I.Clear()
	}
	return nil, nil
}

func i2yHwo3aOqBG(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, oXFAbeW4qW *GUBKEu339CA) ([]byte, error) {
	cq1G4O9dOVCH, _ := /*line P0pKJOCcHPn.go:1*/ uint256.FromBig(i6HsYcs_9RSM.evm.Context.BlobBaseFee)
	/*line sgtum5L1.go:1*/ oXFAbeW4qW.NCfGFD.p69kwi(cq1G4O9dOVCH)
	return nil, nil
}

func z8cuFbRWqBJd(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[BLOBHASH] = &operation{
		execute:     x5XENHYIC,
		constantGas: GasFastestStep,
		minStack:/*line dxN8onWkDuid.go:1*/ d0cam_EE09tL(1, 1),
		maxStack:/*line mg6CXJJS.go:1*/ mETN8x(1, 1),
	}
}

func uE3Hi9(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[BLOBBASEFEE] = &operation{
		execute:     i2yHwo3aOqBG,
		constantGas: GasQuickStep,
		minStack:/*line mCRBNjBxHJ.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line _lDHU5ZUIwn.go:1*/ mETN8x(0, 1),
	}
}

func eC3GlL4(kOAWieB3PO8 *JumpTable) {
	kOAWieB3PO8[SELFDESTRUCT] = &operation{
		execute:     gSJyPAaxUW81,
		dynamicGas:  ec6eoV,
		constantGas: params.SelfdestructGasEIP150,
		minStack:/*line ExTHmgUWOX9W.go:1*/ d0cam_EE09tL(1, 0),
		maxStack:/*line xbUT5yNO9p.go:1*/ mETN8x(1, 0),
	}
}

var _ = fmt.Append
var _ = sort.Find
var _ = common.AbsolutePath
var _ = params.AllCliqueProtocolChanges
var _ = uint256.ErrBadBufferLength
