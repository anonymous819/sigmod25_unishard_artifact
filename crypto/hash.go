//line :1
package crypto

import (
	"bytes"
	"encoding/hex"
	"errors"
	"hash"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

const (
	SHA3_224 = "sha3_224"
	SHA3_256 = "sha3_256"
	SHA3_384 = "sha3_384"
	SHA3_512 = "sha3_512"
)

type KDSldlLG06As []byte

type TA6yah0 interface {
	hash.Hash
	Read([]byte) (int, error)
}

func (w126EF3Hh KDSldlLG06As) Equal(u_rNnl04a5YH KDSldlLG06As) bool {
	return /*line aDA8A9m.go:1*/ bytes.Equal(w126EF3Hh, u_rNnl04a5YH)
}

func (w126EF3Hh KDSldlLG06As) Hex() string {
	return /*line rfjCYtylR.go:1*/ hex.EncodeToString(w126EF3Hh)
}

type K3fvz3wD interface {
	Size() int

	ComputeHash([]byte) KDSldlLG06As

	io.Writer

	SumHash() KDSldlLG06As

	Reset()
}

type wmWoG2oB8SgA struct {
	iqYYUpClDS5 int
}

func GX1JDHqjdD(gMdCo6jn9 []byte) KDSldlLG06As {
	w126EF3Hh := /*line B0kQFnaGhD.go:1*/ make([]byte /*line AybCKv0.go:1*/, len(gMdCo6jn9))
	/*line OGjuZGod22A.go:1*/ copy(w126EF3Hh, gMdCo6jn9)
	return w126EF3Hh
}

func OSowWP5prKgc(mPLZ4MlJe0 []KDSldlLG06As) [][]byte {
	gMdCo6jn9 := /*line BYu7M60.go:1*/ make([][]byte /*line MeGx3xfDJ.go:1*/, len(mPLZ4MlJe0))
	for g0af0v3qZ7, w126EF3Hh := range mPLZ4MlJe0 {
		gMdCo6jn9[g0af0v3qZ7] = w126EF3Hh
	}
	return gMdCo6jn9
}

func YdkRLSPEPcjR(jBsnpUt9Vq3j string) (K3fvz3wD, error) {
	if jBsnpUt9Vq3j == func() /*line ftUPwfKR.go:1*/ string {
		data := /*line ftUPwfKR.go:1*/ make([]byte, 0, 9)
		i := 7
		decryptKey := 216
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 9
				data = /*line ftUPwfKR.go:1*/ append(data, 81)
			case 7:
				data = /*line ftUPwfKR.go:1*/ append(data, 146)
				i = 3
			case 8:
				i = 4
				data = /*line ftUPwfKR.go:1*/ append(data, 78)
			case 0:
				i = 5
				for y := range data {
					data[y] = data[y] + /*line ftUPwfKR.go:1*/ byte(decryptKey^y)
				}
			case 3:
				data = /*line ftUPwfKR.go:1*/ append(data, 136)
				i = 1
			case 6:
				i = 0
				data = /*line ftUPwfKR.go:1*/ append(data, 78)
			case 9:
				i = 8
				data = /*line ftUPwfKR.go:1*/ append(data, 122)
			case 1:
				i = 2
				data = /*line ftUPwfKR.go:1*/ append(data, 126)
			case 4:
				data = /*line ftUPwfKR.go:1*/ append(data, 75)
				i = 6
			}
		}
		return /*line ftUPwfKR.go:1*/ string(data)
	}() {
		return /*line V8PXNxX1oM7D.go:1*/ GJNstJ9Xyk(), nil
	} else if jBsnpUt9Vq3j == func() /*line e3okbN.go:1*/ string {
		data := /*line e3okbN.go:1*/ make([]byte, 0, 9)
		i := 7
		decryptKey := 181
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 8
				data = /*line e3okbN.go:1*/ append(data, 174)
			case 6:
				i = 2
				data = /*line e3okbN.go:1*/ append(data, 132)
			case 7:
				data = /*line e3okbN.go:1*/ append(data, 190)
				i = 9
			case 2:
				i = 0
				for y := range data {
					data[y] = data[y] + /*line e3okbN.go:1*/ byte(decryptKey^y)
				}
			case 3:
				i = 1
				data = /*line e3okbN.go:1*/ append(data, 125)
			case 5:
				data = /*line e3okbN.go:1*/ append(data, 130)
				i = 6
			case 8:
				i = 5
				data = /*line e3okbN.go:1*/ append(data, 130)
			case 9:
				data = /*line e3okbN.go:1*/ append(data, 180)
				i = 4
			case 4:
				data = /*line e3okbN.go:1*/ append(data, 170)
				i = 3
			}
		}
		return /*line e3okbN.go:1*/ string(data)
	}() {
		return /*line Y6jIYYzfXU.go:1*/ OamOdX(), nil
	} else if jBsnpUt9Vq3j == func() /*line dlclGnWEucL8.go:1*/ string {
		seed := /*line dlclGnWEucL8.go:1*/ byte(53)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line dlclGnWEucL8.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line dlclGnWEucL8.go:1*/ fnc(168)(69)(131)(216)(220)(140)(29)(54)
		return /*line dlclGnWEucL8.go:1*/ string(data)
	}() {
		return /*line W3MQZVnZf0.go:1*/ JiUjJXe(), nil
	} else if jBsnpUt9Vq3j == func() /*line LL2xG9z.go:1*/ string {
		data := [] /*line LL2xG9z.go:1*/ byte("sh\x03\b\xdb\xe31'")
		positions := [...]byte{5, 5, 4, 7, 4, 2, 3, 2}
		for i := 0; i < 8; i += 2 {
			localKey := /*line LL2xG9z.go:1*/ byte(i) + /*line LL2xG9z.go:1*/ byte(positions[i]^positions[i+1]) + 82
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line LL2xG9z.go:1*/ string(data)
	}() {
		return /*line SJjuLTvH0.go:1*/ XZMTba4(), nil
	} else {
		return nil /*line iaYJaj88MNb4.go:1*/, errors.New(func() /*line qD99fDHi.go:1*/ string {
			data := [] /*line qD99fDHi.go:1*/ byte("\xf2\x05vZ\xfd\xa2\r \xe5\xf7s\x86e\xfa\x88\xe5\xf2\x86\xd9!")
			positions := [...]byte{15, 1, 14, 14, 4, 6, 17, 1, 13, 4, 16, 1, 0, 17, 9, 11, 3, 8, 5, 14, 8, 18, 9, 13}
			for i := 0; i < 24; i += 2 {
				localKey := /*line qD99fDHi.go:1*/ byte(i) + /*line qD99fDHi.go:1*/ byte(positions[i]^positions[i+1]) + 97
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line qD99fDHi.go:1*/ string(data)
		}())
	}
}

func T3EsjVtT() TA6yah0 {
	return /*line bBtaeaEzpY7A.go:1*/ sha3.NewLegacyKeccak256().(TA6yah0)
}

func ElIzFWasxM(c5GIVS ...[]byte) []byte {
	gMdCo6jn9 := /*line wPElsUVSJy.go:1*/ make([]byte, 32)
	mg5sbf5aS := /*line OVd1aScwCu.go:1*/ T3EsjVtT()
	for _, gMdCo6jn9 := range c5GIVS {
		/*line HaP0A0q6C.go:1*/ mg5sbf5aS.Write(gMdCo6jn9)
	}
	_, iYBOpw := /*line vsDKSzes.go:1*/ mg5sbf5aS.Read(gMdCo6jn9)
	if iYBOpw != nil {

	}
	return gMdCo6jn9
}

func RaQjOAP(c5GIVS ...[]byte) (w126EF3Hh common.Hash) {
	mg5sbf5aS := /*line q76RRAcFNBUP.go:1*/ T3EsjVtT()
	for _, gMdCo6jn9 := range c5GIVS {
		/*line aa3k5LD6.go:1*/ mg5sbf5aS.Write(gMdCo6jn9)
	}
	_, iYBOpw := /*line JtNEacmq.go:1*/ mg5sbf5aS.Read(w126EF3Hh[:])
	if iYBOpw != nil {

	}
	return w126EF3Hh
}

var _ bytes.Buffer
var _ = hex.AppendDecode
var _ = errors.As
var _ hash.Hash
var _ io.ByteReader
var _ = common.AbsolutePath
var _ = sha3.New224
