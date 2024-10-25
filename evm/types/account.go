//line :1
package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

//go:generate go run github.com/fjl/gencodec -type Account -field-override accountMarshaling -out gen_account.go

type Account struct {
	Code    []byte                      `json:"code,omitempty"`
	Storage map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance *big.Int                    `json:"balance" gencodec:"required"`
	Nonce   uint64                      `json:"nonce,omitempty"`

	PrivateKey []byte `json:"secretKey,omitempty"`
}

type lfUHe_P struct {
	RNVKDwVfSwx  hexutil.Bytes
	HU_4xK       *math.HexOrDecimal256
	Mbak_aIb     math.HexOrDecimal64
	GJEa6Z6CMJxy map[storageJSON]storageJSON
	MouBg77isE   hexutil.Bytes
}

type storageJSON common.Hash

func (hP0dFanm4 *storageJSON) UnmarshalText(oyayybCohvzt []byte) error {
	oyayybCohvzt = /*line CC9xmfL.go:1*/ bytes.TrimPrefix(oyayybCohvzt, [] /*line CjqGndjwv.go:1*/ byte("0x"))
	if /*line _KVrvEVg.go:1*/ len(oyayybCohvzt) > 64 {
		return /*line IBbv4ngl.go:1*/ fmt.Errorf(func() /*line yvTeRwn.go:1*/ string {
			seed := /*line yvTeRwn.go:1*/ byte(31)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line yvTeRwn.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line yvTeRwn.go:1*/ fnc(107)(229)(0)(79)(211)(240)(239)(9)(89)(186)(233)(13)(162)(71)(3)(15)(15)(237)(26)(231)(31)(235)(247)(91)(191)(251)(176)(51)(7)(21)(253)(237)(30)(242)(169)(89)(238)(0)(86)(185)(233)(29)(251)(236)(85)(239)(200)
			return /*line yvTeRwn.go:1*/ string(data)
		}(), oyayybCohvzt)
	}
	_b7Govm8wsWN := /*line pOl0KfTo9.go:1*/ len(hP0dFanm4) - /*line ic3lavO.go:1*/ len(oyayybCohvzt)/2
	if _, rfteMJju := /*line EJ44buD66H.go:1*/ hex.Decode(hP0dFanm4[_b7Govm8wsWN:], oyayybCohvzt); rfteMJju != nil {
		return /*line wlBjwa2JhU.go:1*/ fmt.Errorf(func() /*line V0DD2f1ozz.go:1*/ string {
			data := [] /*line V0DD2f1ozz.go:1*/ byte("\x93\xe1\x82a\xa1w\xa2\x9a\x95\xa9\xc2ps|gr\x1d\xa2?xk\x8by/v\x91\x9c\x99\xc2 \xb6\x7f")
			positions := [...]byte{10, 30, 28, 11, 9, 11, 6, 4, 17, 13, 27, 8, 7, 1, 26, 26, 2, 17, 7, 2, 0, 25, 28, 30, 0, 27, 21, 31, 19, 13, 18, 19, 2, 10, 5, 8, 11, 6, 11, 16, 28, 25, 14, 28}
			for i := 0; i < 44; i += 2 {
				localKey := /*line V0DD2f1ozz.go:1*/ byte(i) + /*line V0DD2f1ozz.go:1*/ byte(positions[i]^positions[i+1]) + 194
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line V0DD2f1ozz.go:1*/ string(data)
		}(), oyayybCohvzt)
	}
	return nil
}

func (hP0dFanm4 storageJSON) MarshalText() ([]byte, error) {
	return /*line fcsTz8uY.go:1*/ hexutil.Bytes(hP0dFanm4[:]).MarshalText()
}

type Wu07iXgy map[common.Address]Account

func (cg5_GVcDy *Wu07iXgy) UnmarshalJSON(bgjAklkHvJWu []byte) error {
	s38sBl1pk1u := /*line Bf_TfchVdYah.go:1*/ make(map[common.UnprefixedAddress]Account)
	if rfteMJju := /*line HXIxd7Y.go:1*/ json.Unmarshal(bgjAklkHvJWu, &s38sBl1pk1u); rfteMJju != nil {
		return rfteMJju
	}
	*cg5_GVcDy = /*line hvMy52Y7.go:1*/ make(Wu07iXgy)
	for yOdETLh4, wp2wJGQyqGR := range s38sBl1pk1u {
		(*cg5_GVcDy)[ /*line M1jOtz1TM6xa.go:1*/ common.Address(yOdETLh4)] = wp2wJGQyqGR
	}
	return nil
}

var _ bytes.Buffer
var _ = hex.AppendDecode
var _ = json.Compact
var _ = fmt.Append

const _ = big.Above

var _ = common.AbsolutePath
var _ hexutil.Big
var _ = math.BigMax
