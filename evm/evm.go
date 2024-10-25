//line :1
package evm

import (
	"fmt"
	"math"
	"math/big"
	"time"

	core "unishard/evm/core"
	state "unishard/evm/state"
	types "unishard/evm/types"
	vm "unishard/evm/vm"
	runtime "unishard/evm/vm/runtime"
	unitype "unishard/types"

	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const SHANGHAI_BLOCK_COUNT unitype.BlockHeight = 16830000

func KiHNvp1iNg() (*state.StateDB, error) {
	return /*line cNUpVjU6yJf.go:1*/ state.Fwoiwa1(types.NrGuaNA21 /*line qEzZHs90WF7O.go:1*/, state.LPPxdRUd( /*line IKBNKCC.go:1*/ rawdb.NewMemoryDatabase()), nil)
}

func Mviqes() (*state.StateDB, error) {
	return /*line ovFWp0w6.go:1*/ state.Fwoiwa1(types.NrGuaNA21 /*line RJ3nWsZ_2Q6.go:1*/, state.LPPxdRUd( /*line PzED5zDW.go:1*/ rawdb.NewMemoryDatabase()), nil)
}

func Fqa2qecFzII(wROuk2t *state.StateDB, se71p0 common.Address, g0gfe9SN9o common.Address, o2ZTIeD7C *uint256.Int) error {
	if /*line UacgOL3UdZ.go:1*/ core.S3EqaQJ(wROuk2t, se71p0, o2ZTIeD7C) {
		/*line yN1B313vi.go:1*/ core.BvvkC2thz(wROuk2t, se71p0, g0gfe9SN9o, o2ZTIeD7C)
		/*line AA54JJ.go:1*/ wROuk2t.SetNonce(se71p0 /*line huo8lGSydEDS.go:1*/, wROuk2t.GetNonce(se71p0)+1)
		return nil
	} else {
		return /*line ThymO3.go:1*/ fmt.Errorf(func() /*line DqIoaELp.go:1*/ string {
			fullData := [] /*line DqIoaELp.go:1*/ byte("1\xeb\x87\f3\xcca\xe1\xd3\xcb颲\x15O\xc4~\xa1\xe3Ɇ:\xc5 \xbc\x8d}@\x87\xad}\x14\xe3,Ud\xf7B2\x18`\x01\x1f\xabd\xc0")
			idxKey := [] /*line DqIoaELp.go:1*/ byte("\x10\xfc\x01KI\xa2Wc\xecd\x96m\xa5E\x9d")
			data := /*line DqIoaELp.go:1*/ make([]byte, 0, 24)
			data = /*line DqIoaELp.go:1*/ append(data, fullData[71^ /*line DqIoaELp.go:1*/ int(idxKey[11])]+fullData[99^ /*line DqIoaELp.go:1*/ int(idxKey[11])], fullData[215^ /*line DqIoaELp.go:1*/ int(idxKey[1])]^fullData[243^ /*line DqIoaELp.go:1*/ int(idxKey[1])], fullData[68^ /*line DqIoaELp.go:1*/ int(idxKey[6])]-fullData[117^ /*line DqIoaELp.go:1*/ int(idxKey[6])], fullData[239^ /*line DqIoaELp.go:1*/ int(idxKey[8])]^fullData[205^ /*line DqIoaELp.go:1*/ int(idxKey[8])], fullData[191^ /*line DqIoaELp.go:1*/ int(idxKey[10])]^fullData[186^ /*line DqIoaELp.go:1*/ int(idxKey[10])], fullData[127^ /*line DqIoaELp.go:1*/ int(idxKey[11])]^fullData[116^ /*line DqIoaELp.go:1*/ int(idxKey[11])], fullData[183^ /*line DqIoaELp.go:1*/ int(idxKey[5])]-fullData[171^ /*line DqIoaELp.go:1*/ int(idxKey[5])], fullData[15^ /*line DqIoaELp.go:1*/ int(idxKey[0])]^fullData[22^ /*line DqIoaELp.go:1*/ int(idxKey[0])], fullData[10^ /*line DqIoaELp.go:1*/ int(idxKey[2])]+fullData[23^ /*line DqIoaELp.go:1*/ int(idxKey[2])], fullData[184^ /*line DqIoaELp.go:1*/ int(idxKey[5])]-fullData[175^ /*line DqIoaELp.go:1*/ int(idxKey[5])], fullData[209^ /*line DqIoaELp.go:1*/ int(idxKey[1])]+fullData[212^ /*line DqIoaELp.go:1*/ int(idxKey[1])], fullData[64^ /*line DqIoaELp.go:1*/ int(idxKey[9])]+fullData[116^ /*line DqIoaELp.go:1*/ int(idxKey[9])], fullData[20^ /*line DqIoaELp.go:1*/ int(idxKey[0])]+fullData[11^ /*line DqIoaELp.go:1*/ int(idxKey[0])], fullData[133^ /*line DqIoaELp.go:1*/ int(idxKey[5])]^fullData[188^ /*line DqIoaELp.go:1*/ int(idxKey[5])], fullData[173^ /*line DqIoaELp.go:1*/ int(idxKey[12])]^fullData[180^ /*line DqIoaELp.go:1*/ int(idxKey[12])], fullData[190^ /*line DqIoaELp.go:1*/ int(idxKey[14])]+fullData[133^ /*line DqIoaELp.go:1*/ int(idxKey[14])], fullData[110^ /*line DqIoaELp.go:1*/ int(idxKey[9])]-fullData[120^ /*line DqIoaELp.go:1*/ int(idxKey[9])], fullData[72^ /*line DqIoaELp.go:1*/ int(idxKey[11])]-fullData[106^ /*line DqIoaELp.go:1*/ int(idxKey[11])], fullData[0^ /*line DqIoaELp.go:1*/ int(idxKey[2])]^fullData[3^ /*line DqIoaELp.go:1*/ int(idxKey[2])], fullData[233^ /*line DqIoaELp.go:1*/ int(idxKey[8])]^fullData[241^ /*line DqIoaELp.go:1*/ int(idxKey[8])], fullData[22^ /*line DqIoaELp.go:1*/ int(idxKey[2])]-fullData[13^ /*line DqIoaELp.go:1*/ int(idxKey[2])], fullData[75^ /*line DqIoaELp.go:1*/ int(idxKey[3])]+fullData[109^ /*line DqIoaELp.go:1*/ int(idxKey[3])], fullData[119^ /*line DqIoaELp.go:1*/ int(idxKey[6])]^fullData[67^ /*line DqIoaELp.go:1*/ int(idxKey[6])])
			return /*line DqIoaELp.go:1*/ string(data)
		}())
	}
}

func WTPTcYkXjM(j33SLxQ unitype.BlockHeight, h_YHcC string) *runtime.Config {
	ssMMsa := /*line S26csgu.go:1*/ new(runtime.Config)
	if ssMMsa.Difficulty == nil {
		ssMMsa.Difficulty = /*line ASAga16dJxD.go:1*/ big.NewInt(0)
	}
	if ssMMsa.GasLimit == 0 {
		ssMMsa.GasLimit = math.MaxUint64
	}
	if ssMMsa.GasPrice == nil {
		ssMMsa.GasPrice = /*line trYxels03Lw6.go:1*/ new(big.Int)
	}
	if ssMMsa.Value == nil {
		ssMMsa.Value = /*line RVJ8hCrbHB8J.go:1*/ new(big.Int)
	}
	if ssMMsa.GetHashFn == nil {
		ssMMsa.GetHashFn = func(sUhUhqBNIbsu uint64) common.Hash {
			return /*line aCSzJU.go:1*/ common.BytesToHash( /*line omBkta.go:1*/ crypto.Keccak256([] /*line uY2J3EIxko.go:1*/ byte( /*line eFqv7Z5.go:1*/ new(big.Int).SetUint64(sUhUhqBNIbsu).String())))
		}
	}
	if ssMMsa.BaseFee == nil {
		ssMMsa.BaseFee = /*line LIPfLjdlB.go:1*/ big.NewInt(params.InitialBaseFee)
	}
	if ssMMsa.BlobBaseFee == nil {
		ssMMsa.BlobBaseFee = /*line llKooacFBN.go:1*/ big.NewInt(params.BlobTxMinBlobGasprice)
	}

	ssMMsa.ChainConfig = params.SepoliaChainConfig

	ssMMsa.BlockNumber = /*line tchhltI.go:1*/ big.NewInt( /*line pcOXjpvu.go:1*/ int64(SHANGHAI_BLOCK_COUNT) + /*line wVrAt5F7F8V.go:1*/ int64(j33SLxQ))

	ssMMsa.Time = /*line waPGZes9d.go:1*/ uint64( /*line wCxxBq1FQTR.go:1*/ time.Now().Unix())
	zeSImsa := /*line uL7ZkoH.go:1*/ common.BytesToHash([] /*line B_8NqCl__a.go:1*/ byte(h_YHcC))
	ssMMsa.Random = &zeSImsa

	return ssMMsa
}

func TR_8hD4NdZ(d9FAR2c6EjZ *vm.EVM, wROuk2t *state.StateDB, mLQs1JTkvJv []byte, f8vS_VypKauT common.Address) (uint64, common.Address, error) {
	h1dxeG_Iv := /*line pXxHaOM497Y.go:1*/ vm.AccountRef(f8vS_VypKauT)
	_, dw3w0z, hvQz5nVym4GL, cMNj_dfUsT := /*line BowJWtZFBD.go:1*/ d9FAR2c6EjZ.Create(
		h1dxeG_Iv,
		mLQs1JTkvJv,
		/*line F5yZqcaVp.go:1*/ wROuk2t.GetBalance(f8vS_VypKauT).Uint64()/1e10,
		/*line X90_TCocNWI.go:1*/ uint256.NewInt(0),
	)
	if cMNj_dfUsT != nil {
		return 0, common.Address{0}, cMNj_dfUsT
	}

	return hvQz5nVym4GL, dw3w0z, nil
}

func DFPUN4b8lsOr(d9FAR2c6EjZ *vm.EVM, wROuk2t *state.StateDB, qo_tMHL7zqB []byte, fxblJyU common.Address, z0GfpJ common.Address) (uint64, []byte, error) {
	se71p0 := /*line Y1Y4nJc5Yta.go:1*/ vm.AccountRef(fxblJyU)

	kZa8SM8V, hvQz5nVym4GL, cMNj_dfUsT := /*line ElBPze.go:1*/ d9FAR2c6EjZ.Call(
		se71p0,
		z0GfpJ,
		qo_tMHL7zqB,
		/*line N6muLF.go:1*/ wROuk2t.GetBalance(fxblJyU).Uint64(),
		/*line jZnfrmx4E.go:1*/ uint256.NewInt(0),
	)
	if cMNj_dfUsT != nil {
		return 0, []byte{}, cMNj_dfUsT
	}

	return hvQz5nVym4GL, kZa8SM8V, nil
}

func ZMProjGaK_b(d9FAR2c6EjZ *vm.EVM, wROuk2t *state.StateDB, qo_tMHL7zqB []byte, fxblJyU common.Address, z0GfpJ common.Address) (uint64, []byte, error) {
	se71p0 := /*line HnFEljGf4.go:1*/ vm.AccountRef(fxblJyU)
	kZa8SM8V, hvQz5nVym4GL, cMNj_dfUsT := /*line D0ZvpC.go:1*/ d9FAR2c6EjZ.Call(
		se71p0,
		z0GfpJ,
		qo_tMHL7zqB,
		/*line D1MfD5eSHO.go:1*/ wROuk2t.GetBalance(fxblJyU).Uint64(),
		/*line pdAJvzqSwsZF.go:1*/ uint256.NewInt(20000000),
	)
	if cMNj_dfUsT != nil {
		return 0, []byte{}, cMNj_dfUsT
	}

	return hvQz5nVym4GL, kZa8SM8V, nil
}

var _ = fmt.Append
var _ = math.Abs

const _ = big.Above
const _ = time.ANSIC

var _ = core.UaA5Qe
var _ state.Code
var _ types.AccessList

const _ = vm.ADD

var _ = runtime.FpzWvzX3

const _ = unitype.ABORT

var _ = uint256.ErrBadBufferLength
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ = crypto.CompressPubkey
var _ = params.AllCliqueProtocolChanges
