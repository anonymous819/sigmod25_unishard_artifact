//line :1
package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate go run github.com/fjl/gencodec -type AccessTuple -out gen_access_tuple.go

type AccessList []AccessTuple

type AccessTuple struct {
	Address     common.Address `json:"address"     gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
}

func (a6iaOe1nsP AccessList) StorageKeys() int {
	as9fqKk := 0
	for _, wa4flad := range a6iaOe1nsP {
		as9fqKk += /*line GMbPLsYA.go:1*/ len(wa4flad.StorageKeys)
	}
	return as9fqKk
}

type AccessListTx struct {
	ChainID    *big.Int
	Nonce      uint64
	GasPrice   *big.Int
	Gas        uint64
	To         *common.Address `rlp:"nil"`
	Value      *big.Int
	Data       []byte
	AccessList AccessList
	V, R, S    *big.Int
}

func (iPfjW8i91hC *AccessListTx) acVy5yir() TxData {
	h5R0LDrm := &AccessListTx{
		Nonce: iPfjW8i91hC.Nonce,
		To:/*line _Pe2_mN.go:1*/ syJHwn(iPfjW8i91hC.To),
		Data:/*line JSWPQhvJ1z.go:1*/ common.CopyBytes(iPfjW8i91hC.Data),
		Gas: iPfjW8i91hC.Gas,

		AccessList:/*line WXgKErXF.go:1*/ make(AccessList /*line G3kF2aoyf.go:1*/, len(iPfjW8i91hC.AccessList)),
		Value:/*line pd_kattjj.go:1*/ new(big.Int),
		ChainID:/*line f9SyicxxcLUs.go:1*/ new(big.Int),
		GasPrice:/*line QsHqag.go:1*/ new(big.Int),
		V:/*line B4YcSrMt.go:1*/ new(big.Int),
		R:/*line XgS1AxvaD.go:1*/ new(big.Int),
		S:/*line bhKprnB.go:1*/ new(big.Int),
	}
	/*line QRmsp72n.go:1*/ copy(h5R0LDrm.AccessList, iPfjW8i91hC.AccessList)
	if iPfjW8i91hC.Value != nil {
		/*line rZEFyT.go:1*/ h5R0LDrm.Value.Set(iPfjW8i91hC.Value)
	}
	if iPfjW8i91hC.ChainID != nil {
		/*line aIs3qYJH.go:1*/ h5R0LDrm.ChainID.Set(iPfjW8i91hC.ChainID)
	}
	if iPfjW8i91hC.GasPrice != nil {
		/*line BwkOKWMfHc8R.go:1*/ h5R0LDrm.GasPrice.Set(iPfjW8i91hC.GasPrice)
	}
	if iPfjW8i91hC.V != nil {
		/*line czLw_9q7p.go:1*/ h5R0LDrm.V.Set(iPfjW8i91hC.V)
	}
	if iPfjW8i91hC.R != nil {
		/*line FJE_RL.go:1*/ h5R0LDrm.R.Set(iPfjW8i91hC.R)
	}
	if iPfjW8i91hC.S != nil {
		/*line lT1JdUAVT_lP.go:1*/ h5R0LDrm.S.Set(iPfjW8i91hC.S)
	}
	return h5R0LDrm
}

func (iPfjW8i91hC *AccessListTx) b1pCNt1O() byte             { return AccessListTxType }
func (iPfjW8i91hC *AccessListTx) aVVDgwu() *big.Int          { return iPfjW8i91hC.ChainID }
func (iPfjW8i91hC *AccessListTx) mJq92sCJCrD() AccessList    { return iPfjW8i91hC.AccessList }
func (iPfjW8i91hC *AccessListTx) bgjAklkHvJWu() []byte       { return iPfjW8i91hC.Data }
func (iPfjW8i91hC *AccessListTx) go6R4MnIedw() uint64        { return iPfjW8i91hC.Gas }
func (iPfjW8i91hC *AccessListTx) dtbQrx() *big.Int           { return iPfjW8i91hC.GasPrice }
func (iPfjW8i91hC *AccessListTx) wIXPaTO() *big.Int          { return iPfjW8i91hC.GasPrice }
func (iPfjW8i91hC *AccessListTx) nwO6ag44() *big.Int         { return iPfjW8i91hC.GasPrice }
func (iPfjW8i91hC *AccessListTx) fMqtswJ() *big.Int          { return iPfjW8i91hC.Value }
func (iPfjW8i91hC *AccessListTx) xiCJb__() uint64            { return iPfjW8i91hC.Nonce }
func (iPfjW8i91hC *AccessListTx) ypKkxrFoZ() *common.Address { return iPfjW8i91hC.To }

func (iPfjW8i91hC *AccessListTx) iDoBrObf8(ySxt2pzafw *big.Int, oogLmTRZPjtz *big.Int) *big.Int {
	return /*line CkwSyMmU21b.go:1*/ ySxt2pzafw.Set(iPfjW8i91hC.GasPrice)
}

func (iPfjW8i91hC *AccessListTx) naFUnaJ8Fe() (pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	return iPfjW8i91hC.V, iPfjW8i91hC.R, iPfjW8i91hC.S
}

func (iPfjW8i91hC *AccessListTx) ccFa3ozV(aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	iPfjW8i91hC.ChainID, iPfjW8i91hC.V, iPfjW8i91hC.R, iPfjW8i91hC.S = aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg
}

func (iPfjW8i91hC *AccessListTx) iJPSdycNFb(aYao5YS *bytes.Buffer) error {
	return /*line NeRMhOyIhy.go:1*/ rlp.Encode(aYao5YS, iPfjW8i91hC)
}

func (iPfjW8i91hC *AccessListTx) h5OeLZu4u(uzD2Up []byte) error {
	return /*line UogKQ6Uhx.go:1*/ rlp.DecodeBytes(uzD2Up, iPfjW8i91hC)
}

var _ bytes.Buffer

const _ = big.Above

var _ = common.AbsolutePath
var _ = rlp.AppendUint64
