//line :1
package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

var _ = (* /*line sWs6sEQJSy.go:1*/ lfUHe_P)(nil)

func (wp2wJGQyqGR Account) MarshalJSON() ([]byte, error) {
	type Account struct {
		Code       hexutil.Bytes               `json:"code,omitempty"`
		Storage    map[storageJSON]storageJSON `json:"storage,omitempty"`
		Balance    *math.HexOrDecimal256       `json:"balance" gencodec:"required"`
		Nonce      math.HexOrDecimal64         `json:"nonce,omitempty"`
		PrivateKey hexutil.Bytes               `json:"secretKey,omitempty"`
	}
	var aINiWZ_Jtlzj Account
	aINiWZ_Jtlzj.Code = wp2wJGQyqGR.Code
	if wp2wJGQyqGR.Storage != nil {
		aINiWZ_Jtlzj.Storage = /*line NMIIM8jZtE.go:1*/ make(map[storageJSON]storageJSON /*line kcwLUt.go:1*/, len(wp2wJGQyqGR.Storage))
		for eyxLwPi2, pW1UMRn1s := range wp2wJGQyqGR.Storage {
			aINiWZ_Jtlzj.Storage[ /*line BeJKX88zXhOB.go:1*/ storageJSON(eyxLwPi2)] = /*line qcMWgMA.go:1*/ storageJSON(pW1UMRn1s)
		}
	}
	aINiWZ_Jtlzj.Balance = (* /*line tPuhnE.go:1*/ math.HexOrDecimal256)(wp2wJGQyqGR.Balance)
	aINiWZ_Jtlzj.Nonce = /*line J__vqVFG8EYU.go:1*/ math.HexOrDecimal64(wp2wJGQyqGR.Nonce)
	aINiWZ_Jtlzj.PrivateKey = wp2wJGQyqGR.PrivateKey
	return /*line PyRqcJH.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (wp2wJGQyqGR *Account) UnmarshalJSON(uzD2Up []byte) error {
	type Account struct {
		Code       *hexutil.Bytes              `json:"code,omitempty"`
		Storage    map[storageJSON]storageJSON `json:"storage,omitempty"`
		Balance    *math.HexOrDecimal256       `json:"balance" gencodec:"required"`
		Nonce      *math.HexOrDecimal64        `json:"nonce,omitempty"`
		PrivateKey *hexutil.Bytes              `json:"secretKey,omitempty"`
	}
	var otqLrvn1CD Account
	if rfteMJju := /*line vOoW5ZS.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD); rfteMJju != nil {
		return rfteMJju
	}
	if otqLrvn1CD.Code != nil {
		wp2wJGQyqGR.Code = *otqLrvn1CD.Code
	}
	if otqLrvn1CD.Storage != nil {
		wp2wJGQyqGR.Storage = /*line H1ZPS2Vf3yrb.go:1*/ make(map[common.Hash]common.Hash /*line mNjcONmU4wOG.go:1*/, len(otqLrvn1CD.Storage))
		for eyxLwPi2, pW1UMRn1s := range otqLrvn1CD.Storage {
			wp2wJGQyqGR.Storage[ /*line ctvxm2zx.go:1*/ common.Hash(eyxLwPi2)] = /*line rCjXQw.go:1*/ common.Hash(pW1UMRn1s)
		}
	}
	if otqLrvn1CD.Balance == nil {
		return /*line L0VRFxNum.go:1*/ errors.New(func() /*line a3MtuysajUFt.go:1*/ string {
			data := [] /*line a3MtuysajUFt.go:1*/ byte("m%\x8er\x93\x0f\x8c\x92\xbb\x04\xef\xf4ᎇ\xf1 \xf9\x02zl\xf1\xf8'ba\xd9an*u'vPoo\x18AՐ\xfeunt")
			positions := [...]byte{33, 15, 38, 36, 1, 2, 32, 39, 7, 6, 18, 4, 7, 15, 15, 5, 35, 32, 21, 15, 5, 14, 29, 32, 18, 30, 5, 17, 19, 3, 12, 19, 22, 17, 21, 14, 21, 32, 35, 22, 11, 9, 40, 40, 18, 6, 1, 26, 22, 13, 39, 8, 10, 3}
			for i := 0; i < 54; i += 2 {
				localKey := /*line a3MtuysajUFt.go:1*/ byte(i) + /*line a3MtuysajUFt.go:1*/ byte(positions[i]^positions[i+1]) + 71
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line a3MtuysajUFt.go:1*/ string(data)
		}())
	}
	wp2wJGQyqGR.Balance = (* /*line oEN7GWN.go:1*/ big.Int)(otqLrvn1CD.Balance)
	if otqLrvn1CD.Nonce != nil {
		wp2wJGQyqGR.Nonce = /*line MaxWK0q.go:1*/ uint64(*otqLrvn1CD.Nonce)
	}
	if otqLrvn1CD.PrivateKey != nil {
		wp2wJGQyqGR.PrivateKey = *otqLrvn1CD.PrivateKey
	}
	return nil
}

var _ = json.Compact
var _ = errors.As

const _ = big.Above

var _ = common.AbsolutePath
var _ hexutil.Big
var _ = math.BigMax
