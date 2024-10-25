//line :1
package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type DynamicFeeTx struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int
	GasFeeCap  *big.Int
	Gas        uint64
	To         *common.Address `rlp:"nil"`
	Value      *big.Int
	Data       []byte
	AccessList AccessList

	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

func (iPfjW8i91hC *DynamicFeeTx) acVy5yir() TxData {
	h5R0LDrm := &DynamicFeeTx{
		Nonce: iPfjW8i91hC.Nonce,
		To:/*line aGLeFyR.go:1*/ syJHwn(iPfjW8i91hC.To),
		Data:/*line Ozr07WsHK.go:1*/ common.CopyBytes(iPfjW8i91hC.Data),
		Gas: iPfjW8i91hC.Gas,

		AccessList:/*line Gu5mcEH7iv.go:1*/ make(AccessList /*line nu77NqR5.go:1*/, len(iPfjW8i91hC.AccessList)),
		Value:/*line sZzWImenkJ.go:1*/ new(big.Int),
		ChainID:/*line sa_Jy2.go:1*/ new(big.Int),
		GasTipCap:/*line BAvW50eKbF.go:1*/ new(big.Int),
		GasFeeCap:/*line vhFxg0RoDor.go:1*/ new(big.Int),
		V:/*line kQN66Ce.go:1*/ new(big.Int),
		R:/*line uWEXVyjBz.go:1*/ new(big.Int),
		S:/*line GlIUfcT.go:1*/ new(big.Int),
	}
	/*line UqySVj.go:1*/ copy(h5R0LDrm.AccessList, iPfjW8i91hC.AccessList)
	if iPfjW8i91hC.Value != nil {
		/*line DvDYYB.go:1*/ h5R0LDrm.Value.Set(iPfjW8i91hC.Value)
	}
	if iPfjW8i91hC.ChainID != nil {
		/*line tZvSGGG5amQu.go:1*/ h5R0LDrm.ChainID.Set(iPfjW8i91hC.ChainID)
	}
	if iPfjW8i91hC.GasTipCap != nil {
		/*line lbUt3IP2.go:1*/ h5R0LDrm.GasTipCap.Set(iPfjW8i91hC.GasTipCap)
	}
	if iPfjW8i91hC.GasFeeCap != nil {
		/*line kbe_F5HrEH.go:1*/ h5R0LDrm.GasFeeCap.Set(iPfjW8i91hC.GasFeeCap)
	}
	if iPfjW8i91hC.V != nil {
		/*line qgCeWwRzK8i.go:1*/ h5R0LDrm.V.Set(iPfjW8i91hC.V)
	}
	if iPfjW8i91hC.R != nil {
		/*line J2SMlZaSQj.go:1*/ h5R0LDrm.R.Set(iPfjW8i91hC.R)
	}
	if iPfjW8i91hC.S != nil {
		/*line VN9Miw2SH.go:1*/ h5R0LDrm.S.Set(iPfjW8i91hC.S)
	}
	return h5R0LDrm
}

func (iPfjW8i91hC *DynamicFeeTx) b1pCNt1O() byte             { return DynamicFeeTxType }
func (iPfjW8i91hC *DynamicFeeTx) aVVDgwu() *big.Int          { return iPfjW8i91hC.ChainID }
func (iPfjW8i91hC *DynamicFeeTx) mJq92sCJCrD() AccessList    { return iPfjW8i91hC.AccessList }
func (iPfjW8i91hC *DynamicFeeTx) bgjAklkHvJWu() []byte       { return iPfjW8i91hC.Data }
func (iPfjW8i91hC *DynamicFeeTx) go6R4MnIedw() uint64        { return iPfjW8i91hC.Gas }
func (iPfjW8i91hC *DynamicFeeTx) nwO6ag44() *big.Int         { return iPfjW8i91hC.GasFeeCap }
func (iPfjW8i91hC *DynamicFeeTx) wIXPaTO() *big.Int          { return iPfjW8i91hC.GasTipCap }
func (iPfjW8i91hC *DynamicFeeTx) dtbQrx() *big.Int           { return iPfjW8i91hC.GasFeeCap }
func (iPfjW8i91hC *DynamicFeeTx) fMqtswJ() *big.Int          { return iPfjW8i91hC.Value }
func (iPfjW8i91hC *DynamicFeeTx) xiCJb__() uint64            { return iPfjW8i91hC.Nonce }
func (iPfjW8i91hC *DynamicFeeTx) ypKkxrFoZ() *common.Address { return iPfjW8i91hC.To }

func (iPfjW8i91hC *DynamicFeeTx) iDoBrObf8(ySxt2pzafw *big.Int, oogLmTRZPjtz *big.Int) *big.Int {
	if oogLmTRZPjtz == nil {
		return /*line NgFz6CV.go:1*/ ySxt2pzafw.Set(iPfjW8i91hC.GasFeeCap)
	}
	y64eru18K5 := /*line KykeVMy.go:1*/ ySxt2pzafw.Sub(iPfjW8i91hC.GasFeeCap, oogLmTRZPjtz)
	if /*line HTbbCnoxTgQ.go:1*/ y64eru18K5.Cmp(iPfjW8i91hC.GasTipCap) > 0 {
		/*line HkJBvp76.go:1*/ y64eru18K5.Set(iPfjW8i91hC.GasTipCap)
	}
	return /*line RF_tU9q7BNT.go:1*/ y64eru18K5.Add(y64eru18K5, oogLmTRZPjtz)
}

func (iPfjW8i91hC *DynamicFeeTx) naFUnaJ8Fe() (pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	return iPfjW8i91hC.V, iPfjW8i91hC.R, iPfjW8i91hC.S
}

func (iPfjW8i91hC *DynamicFeeTx) ccFa3ozV(aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	iPfjW8i91hC.ChainID, iPfjW8i91hC.V, iPfjW8i91hC.R, iPfjW8i91hC.S = aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg
}

func (iPfjW8i91hC *DynamicFeeTx) iJPSdycNFb(aYao5YS *bytes.Buffer) error {
	return /*line WJkHNcWEq.go:1*/ rlp.Encode(aYao5YS, iPfjW8i91hC)
}

func (iPfjW8i91hC *DynamicFeeTx) h5OeLZu4u(uzD2Up []byte) error {
	return /*line DJ6REo6OrXN7.go:1*/ rlp.DecodeBytes(uzD2Up, iPfjW8i91hC)
}

var _ bytes.Buffer

const _ = big.Above

var _ = common.AbsolutePath
var _ = rlp.AppendUint64
