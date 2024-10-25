//line :1
package runtime

import (
	core "unishard/evm/core"
	state "unishard/evm/state"
	vm "unishard/evm/vm"
)

func FzlY0sma(v6t6po8y *Config, aSoSIE_ *state.StateDB) *vm.EVM {
	pSw6MDQP := vm.TxContext{
		Origin:     v6t6po8y.Origin,
		GasPrice:   v6t6po8y.GasPrice,
		BlobHashes: v6t6po8y.BlobHashes,
		BlobFeeCap: v6t6po8y.BlobFeeCap,
	}
	lB3DdQXLCVL := vm.BlockContext{
		CanTransfer: core.S3EqaQJ,
		Transfer:    core.BvvkC2thz,
		GetHash:     v6t6po8y.GetHashFn,
		Coinbase:    v6t6po8y.Coinbase,
		BlockNumber: v6t6po8y.BlockNumber,
		Time:        v6t6po8y.Time,
		Difficulty:  v6t6po8y.Difficulty,
		GasLimit:    v6t6po8y.GasLimit,
		BaseFee:     v6t6po8y.BaseFee,
		BlobBaseFee: v6t6po8y.BlobBaseFee,
		Random:      v6t6po8y.Random,
	}

	return /*line kEFtC4yF.go:1*/ vm.C4K6fJpFpRBB(lB3DdQXLCVL, pSw6MDQP, aSoSIE_, v6t6po8y.ChainConfig, v6t6po8y.EVMConfig)
}

var _ = core.UaA5Qe
var _ state.Code

const _ = vm.ADD
