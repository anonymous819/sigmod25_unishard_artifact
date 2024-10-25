//line :1
package runtime

import (
	"math"
	"math/big"

	state "unishard/evm/state"
	types "unishard/evm/types"
	vm "unishard/evm/vm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type Config struct {
	ChainConfig *params.ChainConfig
	Difficulty  *big.Int
	Origin      common.Address
	Coinbase    common.Address
	BlockNumber *big.Int
	Time        uint64
	GasLimit    uint64
	GasPrice    *big.Int
	Value       *big.Int
	Debug       bool
	EVMConfig   vm.Config
	BaseFee     *big.Int
	BlobBaseFee *big.Int
	BlobHashes  []common.Hash
	BlobFeeCap  *big.Int
	Random      *common.Hash

	State     *state.StateDB
	GetHashFn func(oi6lELrTw5uX uint64) common.Hash
}

func kvwS4ZaTs(v6t6po8y *Config) {
	if v6t6po8y.ChainConfig == nil {
		v6t6po8y.ChainConfig = &params.ChainConfig{
			ChainID:/*line fsAkQoTqt.go:1*/ big.NewInt(1),
			HomesteadBlock:/*line JL_H9y.go:1*/ new(big.Int),
			DAOForkBlock:/*line Sd7I63Z6xz.go:1*/ new(big.Int),
			DAOForkSupport: false,
			EIP150Block:/*line cwauE9etBG3.go:1*/ new(big.Int),
			EIP155Block:/*line CpvK05.go:1*/ new(big.Int),
			EIP158Block:/*line am6bNVO7yy.go:1*/ new(big.Int),
			ByzantiumBlock:/*line gIGh5yOiUc.go:1*/ new(big.Int),
			ConstantinopleBlock:/*line JgoVx8EOLjIr.go:1*/ new(big.Int),
			PetersburgBlock:/*line Q1aTgDi.go:1*/ new(big.Int),
			IstanbulBlock:/*line LeYWByUDShC.go:1*/ new(big.Int),
			MuirGlacierBlock:/*line iEtlp880W7Uv.go:1*/ new(big.Int),
			BerlinBlock:/*line S2rxMkBQ5f.go:1*/ new(big.Int),
			LondonBlock:/*line _ISwFdzh0qT.go:1*/ new(big.Int),
		}
	}

	if v6t6po8y.Difficulty == nil {
		v6t6po8y.Difficulty = /*line rrWSFadND.go:1*/ new(big.Int)
	}
	if v6t6po8y.GasLimit == 0 {
		v6t6po8y.GasLimit = math.MaxUint64
	}
	if v6t6po8y.GasPrice == nil {
		v6t6po8y.GasPrice = /*line zdSYDB3UVYAj.go:1*/ new(big.Int)
	}
	if v6t6po8y.Value == nil {
		v6t6po8y.Value = /*line VLixuL.go:1*/ new(big.Int)
	}
	if v6t6po8y.BlockNumber == nil {
		v6t6po8y.BlockNumber = /*line Ky2ebu.go:1*/ new(big.Int)
	}
	if v6t6po8y.GetHashFn == nil {
		v6t6po8y.GetHashFn = func(oi6lELrTw5uX uint64) common.Hash {
			return /*line kqLilpSA.go:1*/ common.BytesToHash( /*line IGuEys.go:1*/ crypto.Keccak256([] /*line yl2ure0_.go:1*/ byte( /*line eGPl0z.go:1*/ new(big.Int).SetUint64(oi6lELrTw5uX).String())))
		}
	}
	if v6t6po8y.BaseFee == nil {
		v6t6po8y.BaseFee = /*line lerD6RJY7B.go:1*/ big.NewInt(params.InitialBaseFee)
	}
	if v6t6po8y.BlobBaseFee == nil {
		v6t6po8y.BlobBaseFee = /*line D69Mcy.go:1*/ big.NewInt(params.BlobTxMinBlobGasprice)
	}
}

func Fu0mUiBTpA(ka7x627sFovc, hlmJfP []byte, v6t6po8y *Config) ([]byte, *state.StateDB, error) {
	if v6t6po8y == nil {
		v6t6po8y = /*line NBREM8xOupp.go:1*/ new(Config)
	}
	/*line fbJGiqiT.go:1*/ kvwS4ZaTs(v6t6po8y)

	if v6t6po8y.State == nil {
		v6t6po8y.State, _ = /*line CaYzyw3.go:1*/ state.Fwoiwa1(types.NrGuaNA21 /*line dxjM3OP.go:1*/, state.LPPxdRUd( /*line rY6Vvw.go:1*/ rawdb.NewMemoryDatabase()), nil)
	}
	var (
		bB3TE71 = /*line jDcZnV0U.go:1*/ common.BytesToAddress([] /*line Ku1ESbaMHU.go:1*/ byte(func() /*line u6XScH.go:1*/ string {
			key := [] /*line u6XScH.go:1*/ byte("\x1f\xa4/\xdf5[j\xdc")
			data := [] /*line u6XScH.go:1*/ byte("D\xcb?\x95=\x06\xf9\x98")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line u6XScH.go:1*/ string(data)
		}()))
		fUK1k5       = /*line UMMZqF0UGL3.go:1*/ FzlY0sma(v6t6po8y, v6t6po8y.State)
		zKYlvkIkp5   = /*line CcvGrUa4ov54.go:1*/ vm.AccountRef(v6t6po8y.Origin)
		eT3htgRUwVuo = /*line RiPsHh2jHoT.go:1*/ v6t6po8y.ChainConfig.Rules(fUK1k5.Context.BlockNumber, fUK1k5.Context.Random != nil, fUK1k5.Context.Time)
	)

	/*line pRk5ARgW.go:1*/
	v6t6po8y.State.Prepare(eT3htgRUwVuo, v6t6po8y.Origin, v6t6po8y.Coinbase, &bB3TE71 /*line RyBsdzf.go:1*/, vm.WNX2wLgB2IK(eT3htgRUwVuo), nil)
	/*line zS9dvU3jzti_.go:1*/ v6t6po8y.State.CreateAccount(bB3TE71)

	/*line CsG6TWvg8n.go:1*/
	v6t6po8y.State.SetCode(bB3TE71, ka7x627sFovc)

	d5phwHP, _, eUeKdu := /*line vSPYJ66dsDT.go:1*/ fUK1k5.Call(
		zKYlvkIkp5,
		/*line rx0aD2j_.go:1*/ common.BytesToAddress([] /*line GZVp60hzJj5.go:1*/ byte(func() /*line UR5SHRIOLmu.go:1*/ string {
			data := [] /*line UR5SHRIOLmu.go:1*/ byte("\x02\x11\x05pn\x03mt")
			positions := [...]byte{4, 1, 0, 3, 2, 6, 0, 3, 0, 2, 5, 1}
			for i := 0; i < 12; i += 2 {
				localKey := /*line UR5SHRIOLmu.go:1*/ byte(i) + /*line UR5SHRIOLmu.go:1*/ byte(positions[i]^positions[i+1]) + 94
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line UR5SHRIOLmu.go:1*/ string(data)
		}())),
		hlmJfP,
		v6t6po8y.GasLimit,
		/*line Ia6hja_5HEcG.go:1*/ uint256.MustFromBig(v6t6po8y.Value),
	)
	return d5phwHP, v6t6po8y.State, eUeKdu
}

func IGLFaF(hlmJfP []byte, v6t6po8y *Config) ([]byte, common.Address, uint64, error) {
	if v6t6po8y == nil {
		v6t6po8y = /*line fCxIfEta.go:1*/ new(Config)
	}
	/*line qR7xAm2.go:1*/ kvwS4ZaTs(v6t6po8y)

	if v6t6po8y.State == nil {
		v6t6po8y.State, _ = /*line n9qzeT3GUNRA.go:1*/ state.Fwoiwa1(types.NrGuaNA21 /*line Mv4QnQJI.go:1*/, state.LPPxdRUd( /*line z5ur8yfUt1r.go:1*/ rawdb.NewMemoryDatabase()), nil)
	}
	var (
		fUK1k5       = /*line PwX9h8aDV6FD.go:1*/ FzlY0sma(v6t6po8y, v6t6po8y.State)
		zKYlvkIkp5   = /*line HFNf02RM.go:1*/ vm.AccountRef(v6t6po8y.Origin)
		eT3htgRUwVuo = /*line ySVWa5A.go:1*/ v6t6po8y.ChainConfig.Rules(fUK1k5.Context.BlockNumber, fUK1k5.Context.Random != nil, fUK1k5.Context.Time)
	)

	/*line GFMsybaIMzF.go:1*/
	v6t6po8y.State.Prepare(eT3htgRUwVuo, v6t6po8y.Origin, v6t6po8y.Coinbase, nil /*line BhpDleTRx.go:1*/, vm.WNX2wLgB2IK(eT3htgRUwVuo), nil)

	ka7x627sFovc, bB3TE71, c4T1rt7k5i, eUeKdu := /*line H5tlGN8s0.go:1*/ fUK1k5.Create(
		zKYlvkIkp5,
		hlmJfP,
		v6t6po8y.GasLimit,
		/*line dHnadTS.go:1*/ uint256.MustFromBig(v6t6po8y.Value),
	)
	return ka7x627sFovc, bB3TE71, c4T1rt7k5i, eUeKdu
}

func FpzWvzX3(bB3TE71 common.Address, hlmJfP []byte, v6t6po8y *Config) ([]byte, uint64, error) {
	/*line SXz_jju.go:1*/ kvwS4ZaTs(v6t6po8y)

	var (
		fUK1k5       = /*line Vte619w7.go:1*/ FzlY0sma(v6t6po8y, v6t6po8y.State)
		zKYlvkIkp5   = /*line Wv6j3K1S.go:1*/ vm.AccountRef(v6t6po8y.Origin)
		aSoSIE_      = v6t6po8y.State
		eT3htgRUwVuo = /*line Evkwda.go:1*/ v6t6po8y.ChainConfig.Rules(fUK1k5.Context.BlockNumber, fUK1k5.Context.Random != nil, fUK1k5.Context.Time)
	)

	/*line QazOnno.go:1*/
	aSoSIE_.Prepare(eT3htgRUwVuo, v6t6po8y.Origin, v6t6po8y.Coinbase, &bB3TE71 /*line SFRBZwuatVE.go:1*/, vm.WNX2wLgB2IK(eT3htgRUwVuo), nil)

	d5phwHP, c4T1rt7k5i, eUeKdu := /*line I0Zybd.go:1*/ fUK1k5.Call(
		zKYlvkIkp5,
		bB3TE71,
		hlmJfP,
		v6t6po8y.GasLimit,
		/*line v9t4zSw.go:1*/ uint256.MustFromBig(v6t6po8y.Value),
	)
	return d5phwHP, c4T1rt7k5i, eUeKdu
}

var _ = math.Abs

const _ = big.Above

var _ state.Code
var _ types.AccessList

const _ = vm.ADD

var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ = crypto.CompressPubkey
var _ = params.AllCliqueProtocolChanges
var _ = uint256.ErrBadBufferLength
