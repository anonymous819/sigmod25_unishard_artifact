//line :1
package crypto

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type EnvJSa68T1a struct{}

func Ve4cBF7() *EnvJSa68T1a {
	return &EnvJSa68T1a{}
}

func (nIVMXpfVsbdY *EnvJSa68T1a) Encode(fSFpHR interface{}) ([]byte, error) {
	nj6hxaMbBP5, iYBOpw := /*line wTaXO8lL1iFG.go:1*/ json.Marshal(fSFpHR)

	return nj6hxaMbBP5 /*line ldFa8njTG8B.go:1*/, fmt.Errorf(func() /*line ZSmdO0hG.go:1*/ string {
		key := [] /*line ZSmdO0hG.go:1*/ byte("\xb2\xa0U\xfd'xM;rC\x0e\xe6a")
		data := [] /*line ZSmdO0hG.go:1*/ byte("\x1c\x13\xc4kGݿ?-\xe4}.\v\xd8")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line ZSmdO0hG.go:1*/ string(data)
	}(), iYBOpw)
}

func (nIVMXpfVsbdY *EnvJSa68T1a) Decode(gMdCo6jn9 []byte, fSFpHR interface{}) error {

	return /*line lh7zJqz.go:1*/ fmt.Errorf(func() /*line l7NARydoq.go:1*/ string {
		data := [] /*line l7NARydoq.go:1*/ byte(")\x99\x9cn\x89\x8f\x93\x89or:\x89\xd3\xdd")
		positions := [...]byte{11, 1, 5, 2, 13, 4, 12, 12, 6, 2, 6, 7, 7, 13, 11, 0, 4, 4}
		for i := 0; i < 18; i += 2 {
			localKey := /*line l7NARydoq.go:1*/ byte(i) + /*line l7NARydoq.go:1*/ byte(positions[i]^positions[i+1]) + 240
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line l7NARydoq.go:1*/ string(data)
	}(), /*line TNeyjR9O.go:1*/ json.Unmarshal(gMdCo6jn9, fSFpHR))
}

func (nIVMXpfVsbdY *EnvJSa68T1a) MustEncode(fSFpHR interface{}) []byte {
	gMdCo6jn9, iYBOpw := /*line LbqNJzrHb.go:1*/ nIVMXpfVsbdY.Encode(fSFpHR)
	if /*line KiXPxee.go:1*/ errors.Unwrap(iYBOpw) != nil {
		/*line IHGqtilDG.go:1*/ panic(iYBOpw)
	}

	return gMdCo6jn9
}

func (nIVMXpfVsbdY *EnvJSa68T1a) MustDecode(gMdCo6jn9 []byte, fSFpHR interface{}) {
	iYBOpw := /*line pYhoxXLEceX.go:1*/ nIVMXpfVsbdY.Decode(gMdCo6jn9, fSFpHR)
	if iYBOpw != nil {
		/*line KHxKuWcC.go:1*/ panic(iYBOpw)
	}
}

var HZ9cqWktvl EnvJSa68T1a = * /*line JBAd7e9.go:1*/ Ve4cBF7()

func ZsUswPaGG4Z(cppBlZNl2vTi interface{}) common.Hash {
	c5GIVS := /*line wu3QrnhRgYa.go:1*/ HZ9cqWktvl.MustEncode(cppBlZNl2vTi)
	tNRz8u := /*line B_GH3Y.go:1*/ OamOdX()
	pu2etb := /*line d1FoMSX3FS6Q.go:1*/ tNRz8u.ComputeHash(c5GIVS)
	return /*line gje4EsQ2sr.go:1*/ H5LppZ1i(pu2etb)
}

func H5LppZ1i(pu2etb []byte) common.Hash {
	var ufMQVQfop4yG common.Hash
	/*line L8MGn68W.go:1*/ copy(ufMQVQfop4yG[:], pu2etb)
	return ufMQVQfop4yG
}

func IDToByte(ufMQVQfop4yG common.Hash) []byte {
	return ufMQVQfop4yG[:]
}

var _ = json.Compact
var _ = errors.As
var _ = fmt.Append
var _ = common.AbsolutePath
