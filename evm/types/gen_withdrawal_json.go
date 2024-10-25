//line :1
package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (* /*line Ajbl18q.go:1*/ e78qeVSzvAI)(nil)

func (yabbhOmaVr Withdrawal) MarshalJSON() ([]byte, error) {
	type Withdrawal struct {
		Index     hexutil.Uint64 `json:"index"`
		Validator hexutil.Uint64 `json:"validatorIndex"`
		Address   common.Address `json:"address"`
		Amount    hexutil.Uint64 `json:"amount"`
	}
	var aINiWZ_Jtlzj Withdrawal
	aINiWZ_Jtlzj.Index = /*line bjpzBhA.go:1*/ hexutil.Uint64(yabbhOmaVr.Index)
	aINiWZ_Jtlzj.Validator = /*line uUofvSEgOlYm.go:1*/ hexutil.Uint64(yabbhOmaVr.Validator)
	aINiWZ_Jtlzj.Address = yabbhOmaVr.Address
	aINiWZ_Jtlzj.Amount = /*line k16FhXUdE.go:1*/ hexutil.Uint64(yabbhOmaVr.Amount)
	return /*line FjCj5jALJz.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (yabbhOmaVr *Withdrawal) UnmarshalJSON(uzD2Up []byte) error {
	type Withdrawal struct {
		Index     *hexutil.Uint64 `json:"index"`
		Validator *hexutil.Uint64 `json:"validatorIndex"`
		Address   *common.Address `json:"address"`
		Amount    *hexutil.Uint64 `json:"amount"`
	}
	var otqLrvn1CD Withdrawal
	if rfteMJju := /*line uX5sXEaAV.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD); rfteMJju != nil {
		return rfteMJju
	}
	if otqLrvn1CD.Index != nil {
		yabbhOmaVr.Index = /*line uVau2lcz.go:1*/ uint64(*otqLrvn1CD.Index)
	}
	if otqLrvn1CD.Validator != nil {
		yabbhOmaVr.Validator = /*line J6z1da4j5A2.go:1*/ uint64(*otqLrvn1CD.Validator)
	}
	if otqLrvn1CD.Address != nil {
		yabbhOmaVr.Address = *otqLrvn1CD.Address
	}
	if otqLrvn1CD.Amount != nil {
		yabbhOmaVr.Amount = /*line DXPKSnRm69c.go:1*/ uint64(*otqLrvn1CD.Amount)
	}
	return nil
}

var _ = json.Compact
var _ = common.AbsolutePath
var _ hexutil.Big
