//line :1
package types

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate go run github.com/fjl/gencodec -type Withdrawal -field-override withdrawalMarshaling -out gen_withdrawal_json.go
//go:generate go run ../../rlp/rlpgen -type Withdrawal -out gen_withdrawal_rlp.go

type Withdrawal struct {
	Index     uint64         `json:"index"`
	Validator uint64         `json:"validatorIndex"`
	Address   common.Address `json:"address"`
	Amount    uint64         `json:"amount"`
}

type e78qeVSzvAI struct {
	P8nEiskvNgKQ hexutil.Uint64
	Po_PG0pix    hexutil.Uint64
	Fa4xzQUq1    hexutil.Uint64
}

type Withdrawals []*Withdrawal

func (gDp0rg Withdrawals) Len() int { return /*line GQ3bRCwBj.go:1*/ len(gDp0rg) }

func (gDp0rg Withdrawals) EncodeIndex(ibcOzM6f int, yabbhOmaVr *bytes.Buffer) {
	/*line DpRxn42.go:1*/ rlp.Encode(yabbhOmaVr, gDp0rg[ibcOzM6f])
}

var _ bytes.Buffer
var _ = common.AbsolutePath
var _ hexutil.Big
var _ = rlp.AppendUint64
