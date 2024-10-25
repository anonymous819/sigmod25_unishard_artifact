//line :1
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//go:generate go run ../../rlp/rlpgen -type Log -out gen_log_rlp.go
//go:generate go run github.com/fjl/gencodec -type Log -field-override logMarshaling -out gen_log_json.go

type Log struct {
	Address common.Address `json:"address" gencodec:"required"`

	Topics []common.Hash `json:"topics" gencodec:"required"`

	Data []byte `json:"data" gencodec:"required"`

	BlockNumber uint64 `json:"blockNumber" rlp:"-"`

	TxHash common.Hash `json:"transactionHash" gencodec:"required" rlp:"-"`

	TxIndex uint `json:"transactionIndex" rlp:"-"`

	BlockHash common.Hash `json:"blockHash" rlp:"-"`

	Index uint `json:"logIndex" rlp:"-"`

	Removed bool `json:"removed" rlp:"-"`
}

type aBHOK_3Irh struct {
	WozQCMKZq4   hexutil.Bytes
	CSRoXZlGoz   hexutil.Uint64
	DzjeOJaeS    hexutil.Uint
	HVF5RGf3SL9f hexutil.Uint
}

var _ = common.AbsolutePath
var _ hexutil.Big
