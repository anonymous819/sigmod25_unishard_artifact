//line :1
package core

import (
	"math/big"

	types "unishard/evm/types"
	vm "unishard/evm/vm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/holiman/uint256"
)

type Pbl4im interface {
	GetHeader(common.Hash, uint64) *types.Header
}

func BuvnqH(df_4ahmAJA *types.Header, gNmiziZ Pbl4im, tC55ScwGQ7B *common.Address) vm.BlockContext {
	var (
		uq8GzQYLJA   common.Address
		q1VdXTanDXB  *big.Int
		ckd79_J2IWeM *big.Int
		o5WBISVgeTEo *common.Hash
	)

	uq8GzQYLJA = *tC55ScwGQ7B
	if df_4ahmAJA.BaseFee != nil {
		q1VdXTanDXB = /*line w3XvpX.go:1*/ new(big.Int).Set(df_4ahmAJA.BaseFee)
	}
	if df_4ahmAJA.ExcessBlobGas != nil {
		ckd79_J2IWeM = /*line aF3RKRB2ls.go:1*/ eip4844.CalcBlobFee(*df_4ahmAJA.ExcessBlobGas)
	}
	if /*line xmnbEdILxq.go:1*/ df_4ahmAJA.Difficulty.Cmp(common.Big0) == 0 {
		o5WBISVgeTEo = &df_4ahmAJA.MixDigest
	}
	return vm.BlockContext{
		CanTransfer: S3EqaQJ,
		Transfer:    BvvkC2thz,
		GetHash:/*line fQkPrFq4p.go:1*/ UpC7Sh(df_4ahmAJA, gNmiziZ),
		Coinbase: uq8GzQYLJA,
		BlockNumber:/*line DH3q1hd.go:1*/ new(big.Int).Set(df_4ahmAJA.Number),
		Time: df_4ahmAJA.Time,
		Difficulty:/*line v4Rruz.go:1*/ new(big.Int).Set(df_4ahmAJA.Difficulty),
		BaseFee:     q1VdXTanDXB,
		BlobBaseFee: ckd79_J2IWeM,
		GasLimit:    df_4ahmAJA.GasLimit,
		Random:      o5WBISVgeTEo,
	}
}

func N2dzdjR(rrEywqc_se *Message) vm.TxContext {
	xeXRxQuT95Di := vm.TxContext{
		Origin: rrEywqc_se.From,
		GasPrice:/*line xIQtNkP2r8Am.go:1*/ new(big.Int).Set(rrEywqc_se.GasPrice),
		BlobHashes: rrEywqc_se.BlobHashes,
	}
	if rrEywqc_se.BlobGasFeeCap != nil {
		xeXRxQuT95Di.BlobFeeCap = /*line dxua2Go_iKC.go:1*/ new(big.Int).Set(rrEywqc_se.BlobGasFeeCap)
	}
	return xeXRxQuT95Di
}

func UpC7Sh(nGHqiV *types.Header, gNmiziZ Pbl4im) func(plVIdO2O7A uint64) common.Hash {

	var hWBpybGBso []common.Hash

	return func(plVIdO2O7A uint64) common.Hash {
		if /*line UELKJ3.go:1*/ nGHqiV.Number.Uint64() <= plVIdO2O7A {

			return common.Hash{}
		}

		if /*line YK2kCPTzl2S7.go:1*/ len(hWBpybGBso) == 0 {
			hWBpybGBso = /*line EbaRjwDH2.go:1*/ append(hWBpybGBso, nGHqiV.ParentHash)
		}
		if q69c3HSGS2XF := /*line nC6Pl_eK.go:1*/ nGHqiV.Number.Uint64() - plVIdO2O7A - 1; q69c3HSGS2XF < /*line fjkQJsa3uW.go:1*/ uint64( /*line PLv_mqYuXm.go:1*/ len(hWBpybGBso)) {
			return hWBpybGBso[q69c3HSGS2XF]
		}

		j6kUIZmzb := hWBpybGBso[ /*line D8MXz5ad.go:1*/ len(hWBpybGBso)-1]
		minxPhNNXU := /*line ISb_Jie5d.go:1*/ nGHqiV.Number.Uint64() - /*line pgbjMDgLDeMK.go:1*/ uint64( /*line qVq79aZkv.go:1*/ len(hWBpybGBso))

		for {
			df_4ahmAJA := /*line bFBDi_wT.go:1*/ gNmiziZ.GetHeader(j6kUIZmzb, minxPhNNXU)
			if df_4ahmAJA == nil {
				break
			}
			hWBpybGBso = /*line eU_hkI.go:1*/ append(hWBpybGBso, df_4ahmAJA.ParentHash)
			j6kUIZmzb = df_4ahmAJA.ParentHash
			minxPhNNXU = /*line A1g6RS.go:1*/ df_4ahmAJA.Number.Uint64() - 1
			if plVIdO2O7A == minxPhNNXU {
				return j6kUIZmzb
			}
		}
		return common.Hash{}
	}
}

func S3EqaQJ(lB6T8T4u4o vm.StateDB, auymTOeVmG common.Address, yMzfLaazE *uint256.Int) bool {
	return /*line H2lFlb.go:1*/ lB6T8T4u4o.GetBalance(auymTOeVmG).Cmp(yMzfLaazE) >= 0
}

func BvvkC2thz(lB6T8T4u4o vm.StateDB, iqQD9nu, reeHa6 common.Address, yMzfLaazE *uint256.Int) {
	/*line y3GmaCDdT.go:1*/ lB6T8T4u4o.SubBalance(iqQD9nu, yMzfLaazE)
	/*line ffHlqJaH.go:1*/ lB6T8T4u4o.AddBalance(reeHa6, yMzfLaazE)
}

const _ = big.Above

var _ types.AccessList

const _ = vm.ADD

var _ = common.AbsolutePath
var _ = eip4844.CalcBlobFee
var _ = uint256.ErrBadBufferLength
