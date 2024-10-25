//line :1
package vm

import (
	"math/big"
	"sync/atomic"

	config "unishard/config"
	types "unishard/evm/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type (
	CanTransferFunc func(StateDB, common.Address, *uint256.Int) bool

	TransferFunc func(StateDB, common.Address, common.Address, *uint256.Int)

	GetHashFunc func(uint64) common.Hash
)

func (oqIE3EcmK *EVM) x9L1UOmWAUf(ewLbwkms common.Address) (VeZhcup, bool) {
	var m6aBsCMavS map[common.Address]VeZhcup
	switch {
	case oqIE3EcmK.chainRules.IsCancun:
		m6aBsCMavS = H5auzN
	case oqIE3EcmK.chainRules.IsBerlin:
		m6aBsCMavS = Agr7Uo
	case oqIE3EcmK.chainRules.IsIstanbul:
		m6aBsCMavS = VN3ry_pWNww
	case oqIE3EcmK.chainRules.IsByzantium:
		m6aBsCMavS = J0k7Wc9OJi
	default:
		m6aBsCMavS = AwN8pSb35
	}
	gTaLKvaWJ, m5pPAO := m6aBsCMavS[ewLbwkms]
	return gTaLKvaWJ, m5pPAO
}

type BlockContext struct {
	CanTransfer CanTransferFunc

	Transfer TransferFunc

	GetHash GetHashFunc

	Coinbase    common.Address
	GasLimit    uint64
	BlockNumber *big.Int
	Time        uint64
	Difficulty  *big.Int
	BaseFee     *big.Int
	BlobBaseFee *big.Int
	Random      *common.Hash
}

type TxContext struct {
	Origin     common.Address
	GasPrice   *big.Int
	BlobHashes []common.Hash
	BlobFeeCap *big.Int
}

type EVM struct {
	Context BlockContext
	TxContext

	StateDB StateDB

	depth int

	chainConfig *params.ChainConfig

	chainRules params.Rules

	Config Config

	interpreter *EVMInterpreter

	abort atomic.Bool

	callGasTemp uint64
}

func C4K6fJpFpRBB(cxffQq BlockContext, cCN8zVgo TxContext, r99Fja StateDB, ayCFTRM *params.ChainConfig, ydlJdL Config) *EVM {

	if ydlJdL.NoBaseFee {
		if /*line Imoi9kYB4Qx.go:1*/ cCN8zVgo.GasPrice.BitLen() == 0 {
			cxffQq.BaseFee = /*line n1afywEKj.go:1*/ new(big.Int)
		}
		if cCN8zVgo.BlobFeeCap != nil && /*line UGlc34njPGhC.go:1*/ cCN8zVgo.BlobFeeCap.BitLen() == 0 {
			cxffQq.BlobBaseFee = /*line C6JcZt.go:1*/ new(big.Int)
		}
	}
	oqIE3EcmK := &EVM{
		Context:     cxffQq,
		TxContext:   cCN8zVgo,
		StateDB:     r99Fja,
		Config:      ydlJdL,
		chainConfig: ayCFTRM,
		chainRules:/*line fM7NANe.go:1*/ ayCFTRM.Rules(cxffQq.BlockNumber, cxffQq.Random != nil, cxffQq.Time),
	}
	oqIE3EcmK.interpreter = /*line GlojJex.go:1*/ FjkJz0a0R(oqIE3EcmK)
	return oqIE3EcmK
}

func (oqIE3EcmK *EVM) Reset(cCN8zVgo TxContext, r99Fja StateDB) {
	oqIE3EcmK.TxContext = cCN8zVgo
	oqIE3EcmK.StateDB = r99Fja
}

func (oqIE3EcmK *EVM) Cancel() {
	/*line m2SQFmSGeB.go:1*/ oqIE3EcmK.abort.Store(true)
}

func (oqIE3EcmK *EVM) Cancelled() bool {
	return /*line HZeqcXfr.go:1*/ oqIE3EcmK.abort.Load()
}

func (oqIE3EcmK *EVM) Interpreter() *EVMInterpreter {
	return oqIE3EcmK.interpreter
}

func (oqIE3EcmK *EVM) Call(dcewTNVj Ck0iTH, ewLbwkms common.Address, sIVM3RLPSKSh []byte, a0SqxOtD uint64, gVUwww *uint256.Int) (ag8Tdqxi4 []byte, lWknatVReio uint64, v1p7xkTykQgN error) {

	if oqIE3EcmK.depth > /*line JxMYcr3Lp.go:1*/ int(params.CallCreateDepth) {
		return nil, a0SqxOtD, IaLTy2I1
	}

	if ! /*line CAmJIM.go:1*/ gVUwww.IsZero() && ! /*line OQxWiwAyq30Z.go:1*/ oqIE3EcmK.Context.CanTransfer(oqIE3EcmK.StateDB /*line FhURzEjO.go:1*/, dcewTNVj.Address(), gVUwww) {
		return nil, a0SqxOtD, ShnqjVTzAmq
	}
	mNx7CEGW_l := /*line ozwjQGW.go:1*/ oqIE3EcmK.StateDB.Snapshot()
	gTaLKvaWJ, djHONazGQG4 := /*line DgfBtNGrG.go:1*/ oqIE3EcmK.x9L1UOmWAUf(ewLbwkms)
	hrNtUBI := oqIE3EcmK.Config.Tracer != nil

	if ! /*line _NJCdMvL.go:1*/ oqIE3EcmK.StateDB.Exist(ewLbwkms) {
		if !djHONazGQG4 && oqIE3EcmK.chainRules.IsEIP158 && /*line Krd6AV.go:1*/ gVUwww.IsZero() {

			if hrNtUBI {
				if oqIE3EcmK.depth == 0 {
					/*line AethLo_.go:1*/ oqIE3EcmK.Config.Tracer.CaptureStart(oqIE3EcmK /*line Aa8fRr.go:1*/, dcewTNVj.Address(), ewLbwkms, false, sIVM3RLPSKSh, a0SqxOtD /*line qGeHGlYBa.go:1*/, gVUwww.ToBig())
					/*line I3CW9BIP.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnd(ag8Tdqxi4, 0, nil)
				} else {
					/*line NgItXo.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnter(CALL /*line fjAmZe.go:1*/, dcewTNVj.Address(), ewLbwkms, sIVM3RLPSKSh, a0SqxOtD /*line uZ2x5XuZO.go:1*/, gVUwww.ToBig())
					/*line Qi2aNcasQt.go:1*/ oqIE3EcmK.Config.Tracer.CaptureExit(ag8Tdqxi4, 0, nil)
				}
			}
			return nil, a0SqxOtD, nil
		}
		/*line BMk8LD8Pj8d.go:1*/ oqIE3EcmK.StateDB.CreateAccount(ewLbwkms)
	}
	/*line _hLBqJT2KeB.go:1*/ oqIE3EcmK.Context.Transfer(oqIE3EcmK.StateDB /*line aIzI0aaR3wc.go:1*/, dcewTNVj.Address(), ewLbwkms, gVUwww)

	if hrNtUBI {
		if oqIE3EcmK.depth == 0 {
			/*line JrbpxF.go:1*/ oqIE3EcmK.Config.Tracer.CaptureStart(oqIE3EcmK /*line Jo2COggt.go:1*/, dcewTNVj.Address(), ewLbwkms, false, sIVM3RLPSKSh, a0SqxOtD /*line fqLNy12mJl.go:1*/, gVUwww.ToBig())
			defer func( /*line ptVLe9A.go:1*/ tp9Ih7YU3 uint64) {
				/*line cCNOW3ZIVq2.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnd(ag8Tdqxi4, tp9Ih7YU3-a0SqxOtD, v1p7xkTykQgN)
			}(a0SqxOtD)
		} else {

			/*line GiY7sWa1Al4.go:1*/
			oqIE3EcmK.Config.Tracer.CaptureEnter(CALL /*line DC1Fq7SN.go:1*/, dcewTNVj.Address(), ewLbwkms, sIVM3RLPSKSh, a0SqxOtD /*line wIg91RAtUe.go:1*/, gVUwww.ToBig())
			defer func( /*line ccEW25eUq.go:1*/ tp9Ih7YU3 uint64) {
				/*line GQ0sSW1OAU.go:1*/ oqIE3EcmK.Config.Tracer.CaptureExit(ag8Tdqxi4, tp9Ih7YU3-a0SqxOtD, v1p7xkTykQgN)
			}(a0SqxOtD)
		}
	}

	if djHONazGQG4 {
		ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN = /*line oikKgU.go:1*/ WCpcUHmYy(gTaLKvaWJ, sIVM3RLPSKSh, a0SqxOtD)
	} else {

		euva86f0e := /*line XDamCI.go:1*/ oqIE3EcmK.StateDB.GetCode(ewLbwkms)
		if /*line jjl3cg.go:1*/ len(euva86f0e) == 0 {
			ag8Tdqxi4, v1p7xkTykQgN = nil, nil
		} else {
			gQerigH := ewLbwkms

			nsYSMon3US := /*line aEnTh0.go:1*/ Jd5jjadjJuy(dcewTNVj /*line AadIH8O1dn.go:1*/, AccountRef(gQerigH), gVUwww, a0SqxOtD)
			/*line HTgdIYacvG.go:1*/ nsYSMon3US.SetCallCode(&gQerigH /*line YAvjE5a.go:1*/, oqIE3EcmK.StateDB.GetCodeHash(gQerigH), euva86f0e)
			ag8Tdqxi4, v1p7xkTykQgN = /*line oFXc_dqq.go:1*/ oqIE3EcmK.interpreter.Run(nsYSMon3US, sIVM3RLPSKSh, false)
			a0SqxOtD = nsYSMon3US.OzrTJ12T
		}
	}

	if v1p7xkTykQgN != nil {
		/*line rt9HS2Fpu.go:1*/ oqIE3EcmK.StateDB.RevertToSnapshot(mNx7CEGW_l)
		if v1p7xkTykQgN != T_LmvOGib {
			a0SqxOtD = 0
		}

	}
	return ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN
}

func (oqIE3EcmK *EVM) CallCode(dcewTNVj Ck0iTH, ewLbwkms common.Address, sIVM3RLPSKSh []byte, a0SqxOtD uint64, gVUwww *uint256.Int) (ag8Tdqxi4 []byte, lWknatVReio uint64, v1p7xkTykQgN error) {

	if oqIE3EcmK.depth > /*line JivSWDa.go:1*/ int(params.CallCreateDepth) {
		return nil, a0SqxOtD, IaLTy2I1
	}

	if ! /*line BvVQPwRgS.go:1*/ oqIE3EcmK.Context.CanTransfer(oqIE3EcmK.StateDB /*line bfODtv.go:1*/, dcewTNVj.Address(), gVUwww) {
		return nil, a0SqxOtD, ShnqjVTzAmq
	}
	var mNx7CEGW_l = /*line NPN3CDz0jF7.go:1*/ oqIE3EcmK.StateDB.Snapshot()

	if oqIE3EcmK.Config.Tracer != nil {
		/*line twcAW_9BkrEK.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnter(CALLCODE /*line zwE0PJoOnmX.go:1*/, dcewTNVj.Address(), ewLbwkms, sIVM3RLPSKSh, a0SqxOtD /*line vPxgbQN8.go:1*/, gVUwww.ToBig())
		defer func( /*line E5sNYCF7gU.go:1*/ tp9Ih7YU3 uint64) {
			/*line _vuH9OE7ZGo6.go:1*/ oqIE3EcmK.Config.Tracer.CaptureExit(ag8Tdqxi4, tp9Ih7YU3-a0SqxOtD, v1p7xkTykQgN)
		}(a0SqxOtD)
	}

	if gTaLKvaWJ, djHONazGQG4 := /*line dDAn6txo.go:1*/ oqIE3EcmK.x9L1UOmWAUf(ewLbwkms); djHONazGQG4 {
		ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN = /*line F3wCGAd2ZsfR.go:1*/ WCpcUHmYy(gTaLKvaWJ, sIVM3RLPSKSh, a0SqxOtD)
	} else {
		gQerigH := ewLbwkms

		nsYSMon3US := /*line Jb8pPo7MT.go:1*/ Jd5jjadjJuy(dcewTNVj /*line rhEqDO49E4.go:1*/, AccountRef( /*line LKIODve.go:1*/ dcewTNVj.Address()), gVUwww, a0SqxOtD)
		/*line yHFdn6v5.go:1*/ nsYSMon3US.SetCallCode(&gQerigH /*line Ha_FCun.go:1*/, oqIE3EcmK.StateDB.GetCodeHash(gQerigH) /*line e3NYxqhU.go:1*/, oqIE3EcmK.StateDB.GetCode(gQerigH))
		ag8Tdqxi4, v1p7xkTykQgN = /*line hOOR56.go:1*/ oqIE3EcmK.interpreter.Run(nsYSMon3US, sIVM3RLPSKSh, false)
		a0SqxOtD = nsYSMon3US.OzrTJ12T
	}
	if v1p7xkTykQgN != nil {
		/*line PJNyGIYH.go:1*/ oqIE3EcmK.StateDB.RevertToSnapshot(mNx7CEGW_l)
		if v1p7xkTykQgN != T_LmvOGib {
			a0SqxOtD = 0
		}
	}
	return ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN
}

func (oqIE3EcmK *EVM) DelegateCall(dcewTNVj Ck0iTH, ewLbwkms common.Address, sIVM3RLPSKSh []byte, a0SqxOtD uint64) (ag8Tdqxi4 []byte, lWknatVReio uint64, v1p7xkTykQgN error) {

	if oqIE3EcmK.depth > /*line DQz1mgKLz_0c.go:1*/ int(params.CallCreateDepth) {
		return nil, a0SqxOtD, IaLTy2I1
	}
	var mNx7CEGW_l = /*line E7K_lFvewfr_.go:1*/ oqIE3EcmK.StateDB.Snapshot()

	if oqIE3EcmK.Config.Tracer != nil {

		fzrwnz96pkv := dcewTNVj.(*WyJ004I8)

		/*line lauOiJIcBK.go:1*/
		oqIE3EcmK.Config.Tracer.CaptureEnter(DELEGATECALL /*line jsGVC47oeY.go:1*/, dcewTNVj.Address(), ewLbwkms, sIVM3RLPSKSh, a0SqxOtD /*line mGQNlXV7Bst.go:1*/, fzrwnz96pkv.mddAyugJc.ToBig())
		defer func( /*line fUZTeywtX.go:1*/ tp9Ih7YU3 uint64) {
			/*line SHo5DgV.go:1*/ oqIE3EcmK.Config.Tracer.CaptureExit(ag8Tdqxi4, tp9Ih7YU3-a0SqxOtD, v1p7xkTykQgN)
		}(a0SqxOtD)
	}

	if gTaLKvaWJ, djHONazGQG4 := /*line sHmn1_Ej3Xf.go:1*/ oqIE3EcmK.x9L1UOmWAUf(ewLbwkms); djHONazGQG4 {
		ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN = /*line BDc2Az.go:1*/ WCpcUHmYy(gTaLKvaWJ, sIVM3RLPSKSh, a0SqxOtD)
	} else {
		gQerigH := ewLbwkms

		nsYSMon3US := /*line asvDEK4ka6rl.go:1*/ Jd5jjadjJuy(dcewTNVj /*line ncFNeOCs.go:1*/, AccountRef( /*line gjNMaX.go:1*/ dcewTNVj.Address()), nil, a0SqxOtD).AsDelegate()
		/*line AvQiZNNmFp.go:1*/ nsYSMon3US.SetCallCode(&gQerigH /*line eLYQgJQx0uY.go:1*/, oqIE3EcmK.StateDB.GetCodeHash(gQerigH) /*line YZrx10b9J.go:1*/, oqIE3EcmK.StateDB.GetCode(gQerigH))
		ag8Tdqxi4, v1p7xkTykQgN = /*line IMPwaQgRKE.go:1*/ oqIE3EcmK.interpreter.Run(nsYSMon3US, sIVM3RLPSKSh, false)
		a0SqxOtD = nsYSMon3US.OzrTJ12T
	}
	if v1p7xkTykQgN != nil {
		/*line XPt8gp3d_.go:1*/ oqIE3EcmK.StateDB.RevertToSnapshot(mNx7CEGW_l)
		if v1p7xkTykQgN != T_LmvOGib {
			a0SqxOtD = 0
		}
	}
	return ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN
}

func (oqIE3EcmK *EVM) StaticCall(dcewTNVj Ck0iTH, ewLbwkms common.Address, sIVM3RLPSKSh []byte, a0SqxOtD uint64) (ag8Tdqxi4 []byte, lWknatVReio uint64, v1p7xkTykQgN error) {

	if oqIE3EcmK.depth > /*line vNowB_.go:1*/ int(params.CallCreateDepth) {
		return nil, a0SqxOtD, IaLTy2I1
	}

	var mNx7CEGW_l = /*line muTTyf6f.go:1*/ oqIE3EcmK.StateDB.Snapshot()

	/*line or1vYTYph7A.go:1*/
	oqIE3EcmK.StateDB.AddBalance(ewLbwkms /*line ylUQPcWkI.go:1*/, new(uint256.Int))

	if oqIE3EcmK.Config.Tracer != nil {
		/*line Bez3omPxVhk.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnter(STATICCALL /*line plmTJaWhlruA.go:1*/, dcewTNVj.Address(), ewLbwkms, sIVM3RLPSKSh, a0SqxOtD, nil)
		defer func( /*line fJFIPV.go:1*/ tp9Ih7YU3 uint64) {
			/*line IthrBzAA.go:1*/ oqIE3EcmK.Config.Tracer.CaptureExit(ag8Tdqxi4, tp9Ih7YU3-a0SqxOtD, v1p7xkTykQgN)
		}(a0SqxOtD)
	}

	if gTaLKvaWJ, djHONazGQG4 := /*line AeNTCSim2nJ.go:1*/ oqIE3EcmK.x9L1UOmWAUf(ewLbwkms); djHONazGQG4 {
		ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN = /*line AGjP0d.go:1*/ WCpcUHmYy(gTaLKvaWJ, sIVM3RLPSKSh, a0SqxOtD)
	} else {

		gQerigH := ewLbwkms

		nsYSMon3US := /*line IP7FStQ1TR.go:1*/ Jd5jjadjJuy(dcewTNVj /*line H5SF0lvF09s.go:1*/, AccountRef(gQerigH) /*line C6D_bZxTQ.go:1*/, new(uint256.Int), a0SqxOtD)
		/*line DtcxBFd.go:1*/ nsYSMon3US.SetCallCode(&gQerigH /*line pS3LiOyvA.go:1*/, oqIE3EcmK.StateDB.GetCodeHash(gQerigH) /*line ZvQGiXj2WW7G.go:1*/, oqIE3EcmK.StateDB.GetCode(gQerigH))

		ag8Tdqxi4, v1p7xkTykQgN = /*line hQk4EgQ.go:1*/ oqIE3EcmK.interpreter.Run(nsYSMon3US, sIVM3RLPSKSh, true)
		a0SqxOtD = nsYSMon3US.OzrTJ12T
	}
	if v1p7xkTykQgN != nil {
		/*line Ec5Bh3ald2P1.go:1*/ oqIE3EcmK.StateDB.RevertToSnapshot(mNx7CEGW_l)
		if v1p7xkTykQgN != T_LmvOGib {
			a0SqxOtD = 0
		}
	}
	return ag8Tdqxi4, a0SqxOtD, v1p7xkTykQgN
}

type codeAndHash struct {
	code []byte
	hash common.Hash
}

func (oPtEQcrX *codeAndHash) Hash() common.Hash {
	if oPtEQcrX.hash == (common.Hash{}) {
		oPtEQcrX.hash = /*line haJmmpa2pbb.go:1*/ crypto.Keccak256Hash(oPtEQcrX.code)
	}
	return oPtEQcrX.hash
}

func (oqIE3EcmK *EVM) wGdwdlv(dcewTNVj Ck0iTH, hAawFa8 *codeAndHash, euva86f0e []byte, a0SqxOtD uint64, gVUwww *uint256.Int, qtbv6eO common.Address, oyspArz0 LxosJe8) ([]byte, common.Address, uint64, error) {

	if oqIE3EcmK.depth > /*line bW8tADn.go:1*/ int(params.CallCreateDepth) {
		return nil, common.Address{}, a0SqxOtD, IaLTy2I1
	}
	if ! /*line F6TPm3S4.go:1*/ oqIE3EcmK.Context.CanTransfer(oqIE3EcmK.StateDB /*line _bJ5J_sc.go:1*/, dcewTNVj.Address(), gVUwww) {
		return nil, common.Address{}, a0SqxOtD, ShnqjVTzAmq
	}
	cK7czJ5f := /*line y9di4o.go:1*/ oqIE3EcmK.StateDB.GetNonce( /*line WGp_RaT.go:1*/ dcewTNVj.Address())
	if cK7czJ5f+1 < cK7czJ5f {
		return nil, common.Address{}, a0SqxOtD, TtKxX91
	}
	/*line ovupPt.go:1*/ oqIE3EcmK.StateDB.SetNonce( /*line Clma7Hm.go:1*/ dcewTNVj.Address(), cK7czJ5f+1)

	if oqIE3EcmK.chainRules.IsBerlin {
		/*line FhpQjNix.go:1*/ oqIE3EcmK.StateDB.AddAddressToAccessList(qtbv6eO)
	}

	yOOMGRlw := /*line BEzKdEcQPyy.go:1*/ oqIE3EcmK.StateDB.GetCodeHash(qtbv6eO)
	if /*line TE3Fqpg9.go:1*/ oqIE3EcmK.StateDB.GetNonce(qtbv6eO) != 0 || (yOOMGRlw != (common.Hash{}) && yOOMGRlw != types.JhrQnETMFm) {
		return nil, common.Address{}, 0, ADSu6fWP
	}

	mNx7CEGW_l := /*line SdsaL3M6zAIC.go:1*/ oqIE3EcmK.StateDB.Snapshot()
	/*line Jk0qKP4.go:1*/ oqIE3EcmK.StateDB.CreateAccount(qtbv6eO)
	if oqIE3EcmK.chainRules.IsEIP158 {
		/*line zYI0Jzx.go:1*/ oqIE3EcmK.StateDB.SetNonce(qtbv6eO, 1)
	}
	/*line HvzYJIuLa.go:1*/ oqIE3EcmK.Context.Transfer(oqIE3EcmK.StateDB /*line TwElZn2SYM.go:1*/, dcewTNVj.Address(), qtbv6eO, gVUwww)

	nsYSMon3US := /*line __039Bkx.go:1*/ Jd5jjadjJuy(dcewTNVj /*line iY78P9c.go:1*/, AccountRef(qtbv6eO), gVUwww, a0SqxOtD)
	/*line RCIrNz6.go:1*/ nsYSMon3US.SetCodeOptionalHash(&qtbv6eO, hAawFa8)

	if oqIE3EcmK.Config.Tracer != nil {
		if oqIE3EcmK.depth == 0 {
			/*line RitMkqHev.go:1*/ oqIE3EcmK.Config.Tracer.CaptureStart(oqIE3EcmK /*line Up0UBe.go:1*/, dcewTNVj.Address(), qtbv6eO, true, hAawFa8.code, a0SqxOtD /*line gjqztGZ.go:1*/, gVUwww.ToBig())
		} else {
			/*line wJXEof.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnter(oyspArz0 /*line WaXUzGiIdVN4.go:1*/, dcewTNVj.Address(), qtbv6eO, hAawFa8.code, a0SqxOtD /*line H_SWS344JW.go:1*/, gVUwww.ToBig())
		}
	}

	ag8Tdqxi4, v1p7xkTykQgN := /*line a144xuA3nT3f.go:1*/ oqIE3EcmK.interpreter.Run(nsYSMon3US, nil, false)

	if v1p7xkTykQgN == nil && oqIE3EcmK.chainRules.IsEIP158 && /*line Ahc_K5.go:1*/ len(ag8Tdqxi4) > params.MaxCodeSize {
		v1p7xkTykQgN = HRghUPhaXRh
	}

	if v1p7xkTykQgN == nil && /*line nmSW_NDU.go:1*/ len(ag8Tdqxi4) >= 1 && ag8Tdqxi4[0] == 0xEF && oqIE3EcmK.chainRules.IsLondon {
		v1p7xkTykQgN = GB05_B8mV9sm
	}

	if v1p7xkTykQgN == nil {
		mWc4WLa := /*line dDmorU.go:1*/ uint64( /*line gVLWu5iBY.go:1*/ len(ag8Tdqxi4)) * params.CreateDataGas
		if /*line xqgYDIpERUk.go:1*/ nsYSMon3US.UseGas(mWc4WLa) {
			/*line AoRVmAKD.go:1*/ oqIE3EcmK.StateDB.SetCode(qtbv6eO, ag8Tdqxi4)
			/*line I41atxx.go:1*/ oqIE3EcmK.StateDB.SetDeployCode(qtbv6eO, euva86f0e)
		} else {
			v1p7xkTykQgN = Wrtcx8Yb8
		}
	}

	if v1p7xkTykQgN != nil && (oqIE3EcmK.chainRules.IsHomestead || v1p7xkTykQgN != Wrtcx8Yb8) {
		/*line ArPAdqJs.go:1*/ oqIE3EcmK.StateDB.RevertToSnapshot(mNx7CEGW_l)
		if v1p7xkTykQgN != T_LmvOGib {
			/*line aZfsmp.go:1*/ nsYSMon3US.UseGas(nsYSMon3US.OzrTJ12T)
		}
	}

	if oqIE3EcmK.Config.Tracer != nil {
		if oqIE3EcmK.depth == 0 {
			/*line azbQT1lA6n5.go:1*/ oqIE3EcmK.Config.Tracer.CaptureEnd(ag8Tdqxi4, a0SqxOtD-nsYSMon3US.OzrTJ12T, v1p7xkTykQgN)
		} else {
			/*line GYFU0KaRlR.go:1*/ oqIE3EcmK.Config.Tracer.CaptureExit(ag8Tdqxi4, a0SqxOtD-nsYSMon3US.OzrTJ12T, v1p7xkTykQgN)
		}
	}
	return ag8Tdqxi4, qtbv6eO, nsYSMon3US.OzrTJ12T, v1p7xkTykQgN
}

func (oqIE3EcmK *EVM) Create(dcewTNVj Ck0iTH, euva86f0e []byte, a0SqxOtD uint64, gVUwww *uint256.Int) (ag8Tdqxi4 []byte, qfxiJ_1HIhg6 common.Address, lWknatVReio uint64, v1p7xkTykQgN error) {
	qfxiJ_1HIhg6 = /*line zItArd0ucT.go:1*/ crypto.CreateAddress( /*line tDbMed1.go:1*/ dcewTNVj.Address() /*line W1cQskYrTCT.go:1*/, oqIE3EcmK.StateDB.GetNonce( /*line K2ccVM0Q0ap1.go:1*/ dcewTNVj.Address()))
	wHMwHa6E3 := /*line WTXvIfYsN_g2.go:1*/ utils.CalculateShardToSend([]common.Address{qfxiJ_1HIhg6} /*line BQAyetNxS.go:1*/, config.GetConfig().ShardCount)[0]
	kOdSwR := /*line EtYtqdtgP.go:1*/ utils.CalculateShardToSend([]common.Address{ /*line D2jP9WRIz.go:1*/ dcewTNVj.Address()} /*line k6iNAPunxaG.go:1*/, config.GetConfig().ShardCount)[0]
	if wHMwHa6E3 != kOdSwR {
		qfxiJ_1HIhg6 = /*line GiOiX9H5x.go:1*/ common.BigToAddress( /*line UZ3PUVo_rqqa.go:1*/ qfxiJ_1HIhg6.Big().Add( /*line PJTSq06.go:1*/ qfxiJ_1HIhg6.Big() /*line exV2Bik.go:1*/, big.NewInt( /*line RlrWmY.go:1*/ int64(kOdSwR)- /*line usz3xu.go:1*/ int64(wHMwHa6E3))))
	}
	return /*line eCX9ITTk4qJx.go:1*/ oqIE3EcmK.wGdwdlv(dcewTNVj, &codeAndHash{code: euva86f0e}, euva86f0e, a0SqxOtD, gVUwww, qfxiJ_1HIhg6, CREATE)
}

func (oqIE3EcmK *EVM) CreateFDCM(vvq4w6UE49Q common.Address, dcewTNVj Ck0iTH, euva86f0e []byte, a0SqxOtD uint64, gVUwww *uint256.Int) (ag8Tdqxi4 []byte, qfxiJ_1HIhg6 common.Address, lWknatVReio uint64, v1p7xkTykQgN error) {
	return /*line fnA4ma5y7sa.go:1*/ oqIE3EcmK.wGdwdlv(dcewTNVj, &codeAndHash{code: euva86f0e}, euva86f0e, a0SqxOtD, gVUwww, vvq4w6UE49Q, CREATE)
}

func (oqIE3EcmK *EVM) Create2(dcewTNVj Ck0iTH, euva86f0e []byte, a0SqxOtD uint64, eom5CarDaJ *uint256.Int, m_rrjN *uint256.Int) (ag8Tdqxi4 []byte, qfxiJ_1HIhg6 common.Address, lWknatVReio uint64, v1p7xkTykQgN error) {
	hAawFa8 := &codeAndHash{code: euva86f0e}
	qfxiJ_1HIhg6 = /*line dzJ2Gxryr.go:1*/ crypto.CreateAddress2( /*line iTQjJ1.go:1*/ dcewTNVj.Address() /*line JOoTGJbJ.go:1*/, m_rrjN.Bytes32() /*line kqs5sEZwPc.go:1*/, hAawFa8.Hash().Bytes())
	return /*line sY1Q7X.go:1*/ oqIE3EcmK.wGdwdlv(dcewTNVj, hAawFa8, euva86f0e, a0SqxOtD, eom5CarDaJ, qfxiJ_1HIhg6, CREATE2)
}

func (oqIE3EcmK *EVM) ChainConfig() *params.ChainConfig { return oqIE3EcmK.chainConfig }

const _ = big.Above

var _ = atomic.AddInt32
var _ config.Bconfig
var _ types.AccessList
var _ = utils.GffGNE
var _ = common.AbsolutePath
var _ = crypto.CompressPubkey
var _ = params.AllCliqueProtocolChanges
var _ = uint256.ErrBadBufferLength
