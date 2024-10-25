//line :1
package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LegacyTx struct {
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       *common.Address `rlp:"nil"`
	Value    *big.Int
	Data     []byte
	V, R, S  *big.Int
}

func ZcgLfhKfFo1(xiCJb__ uint64, ypKkxrFoZ common.Address, ak5MYZiydpyX *big.Int, aKaQ4RCfo uint64, dtbQrx *big.Int, bgjAklkHvJWu []byte) *Transaction {
	return /*line GHKsIn3ku.go:1*/ CV6bt80q7Gn(&LegacyTx{
		Nonce:    xiCJb__,
		To:       &ypKkxrFoZ,
		Value:    ak5MYZiydpyX,
		Gas:      aKaQ4RCfo,
		GasPrice: dtbQrx,
		Data:     bgjAklkHvJWu,
	})
}

func UIWEQVl1(xiCJb__ uint64, ak5MYZiydpyX *big.Int, aKaQ4RCfo uint64, dtbQrx *big.Int, bgjAklkHvJWu []byte) *Transaction {
	return /*line JgA1Lk.go:1*/ CV6bt80q7Gn(&LegacyTx{
		Nonce:    xiCJb__,
		Value:    ak5MYZiydpyX,
		Gas:      aKaQ4RCfo,
		GasPrice: dtbQrx,
		Data:     bgjAklkHvJWu,
	})
}

func (iPfjW8i91hC *LegacyTx) acVy5yir() TxData {
	h5R0LDrm := &LegacyTx{
		Nonce: iPfjW8i91hC.Nonce,
		To:/*line Bc6i8tczO6.go:1*/ syJHwn(iPfjW8i91hC.To),
		Data:/*line W4PfTmkREmt.go:1*/ common.CopyBytes(iPfjW8i91hC.Data),
		Gas: iPfjW8i91hC.Gas,

		Value:/*line lCFUvGa3Fccx.go:1*/ new(big.Int),
		GasPrice:/*line fY5R0sQ.go:1*/ new(big.Int),
		V:/*line y4zQoWSh.go:1*/ new(big.Int),
		R:/*line t6MgroOMQ.go:1*/ new(big.Int),
		S:/*line JHVgSpDKry.go:1*/ new(big.Int),
	}
	if iPfjW8i91hC.Value != nil {
		/*line E8C6P13p.go:1*/ h5R0LDrm.Value.Set(iPfjW8i91hC.Value)
	}
	if iPfjW8i91hC.GasPrice != nil {
		/*line t6XBPT6WcFaR.go:1*/ h5R0LDrm.GasPrice.Set(iPfjW8i91hC.GasPrice)
	}
	if iPfjW8i91hC.V != nil {
		/*line oZtPX6orX9.go:1*/ h5R0LDrm.V.Set(iPfjW8i91hC.V)
	}
	if iPfjW8i91hC.R != nil {
		/*line J1t66VhYUVGG.go:1*/ h5R0LDrm.R.Set(iPfjW8i91hC.R)
	}
	if iPfjW8i91hC.S != nil {
		/*line RFbEvcd8KLG.go:1*/ h5R0LDrm.S.Set(iPfjW8i91hC.S)
	}
	return h5R0LDrm
}

func (iPfjW8i91hC *LegacyTx) b1pCNt1O() byte { return LegacyTxType }
func (iPfjW8i91hC *LegacyTx) aVVDgwu() *big.Int {
	return /*line MzW5y5GZi4h.go:1*/ sikVyPV4yp(iPfjW8i91hC.V)
}
func (iPfjW8i91hC *LegacyTx) mJq92sCJCrD() AccessList    { return nil }
func (iPfjW8i91hC *LegacyTx) bgjAklkHvJWu() []byte       { return iPfjW8i91hC.Data }
func (iPfjW8i91hC *LegacyTx) go6R4MnIedw() uint64        { return iPfjW8i91hC.Gas }
func (iPfjW8i91hC *LegacyTx) dtbQrx() *big.Int           { return iPfjW8i91hC.GasPrice }
func (iPfjW8i91hC *LegacyTx) wIXPaTO() *big.Int          { return iPfjW8i91hC.GasPrice }
func (iPfjW8i91hC *LegacyTx) nwO6ag44() *big.Int         { return iPfjW8i91hC.GasPrice }
func (iPfjW8i91hC *LegacyTx) fMqtswJ() *big.Int          { return iPfjW8i91hC.Value }
func (iPfjW8i91hC *LegacyTx) xiCJb__() uint64            { return iPfjW8i91hC.Nonce }
func (iPfjW8i91hC *LegacyTx) ypKkxrFoZ() *common.Address { return iPfjW8i91hC.To }

func (iPfjW8i91hC *LegacyTx) iDoBrObf8(ySxt2pzafw *big.Int, oogLmTRZPjtz *big.Int) *big.Int {
	return /*line AcFoxN.go:1*/ ySxt2pzafw.Set(iPfjW8i91hC.GasPrice)
}

func (iPfjW8i91hC *LegacyTx) naFUnaJ8Fe() (pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	return iPfjW8i91hC.V, iPfjW8i91hC.R, iPfjW8i91hC.S
}

func (iPfjW8i91hC *LegacyTx) ccFa3ozV(aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	iPfjW8i91hC.V, iPfjW8i91hC.R, iPfjW8i91hC.S = pW1UMRn1s, fWuxre8U, gDp0rg
}

func (iPfjW8i91hC *LegacyTx) iJPSdycNFb(*bytes.Buffer) error {
	/*line WOLdawQI.go:1*/ panic(func() /*line hRb2OyT_V.go:1*/ string {
		key := [] /*line hRb2OyT_V.go:1*/ byte("#\xaa\x14\x9a뵺\xc2\xc8[\xbc\xa5\xabk\x93\xc6!\xdc`Ύ\xbe\xb7\xff\x7f")
		data := [] /*line hRb2OyT_V.go:1*/ byte("\x88\x18w\tO\x1a\xda%)\xc7(\n\x0f\x8b\x024A(\xc55\xef!0S\xf7")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line hRb2OyT_V.go:1*/ string(data)
	}())
}

func (iPfjW8i91hC *LegacyTx) h5OeLZu4u([]byte) error {
	/*line rQrCCdFbUrg0.go:1*/ panic(func() /*line _JAucRBCJ.go:1*/ string {
		key := [] /*line _JAucRBCJ.go:1*/ byte("r\xf7B\xb5ǀ\xe2:\x88\x7f\xe9\xca\xc2\xd4fʧ\xb5)qm\x15\xd9]\xf1\x02")
		data := [] /*line _JAucRBCJ.go:1*/ byte("\xf2n!\xba\x9d\xe5>)\xd9탛\xa2L\t\xa4y\x97<\xf6\xf4N\xa0\xf7\x87'")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line _JAucRBCJ.go:1*/ string(data)
	}())
}

var _ bytes.Buffer

const _ = big.Above

var _ = common.AbsolutePath
