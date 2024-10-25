//line :1
package trie

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

var MWbaoG = /*line mDQnTv9C.go:1*/ errors.New(func() /*line BHs9zCTl4bE.go:1*/ string {
	data := /*line BHs9zCTl4bE.go:1*/ make([]byte, 0, 26)
	i := 11
	decryptKey := 228
	for counter := 0; i != 8; counter++ {
		decryptKey ^= i * counter
		switch i {
		case 2:
			data = /*line BHs9zCTl4bE.go:1*/ append(data, "3?"...,
			)
			i = 6
		case 1:
			i = 7
			data = /*line BHs9zCTl4bE.go:1*/ append(data, "67C@"...,
			)
		case 6:
			data = /*line BHs9zCTl4bE.go:1*/ append(data, 66)
			i = 9
		case 12:
			data = /*line BHs9zCTl4bE.go:1*/ append(data, ":5"...,
			)
			i = 1
		case 11:
			data = /*line BHs9zCTl4bE.go:1*/ append(data, "NMA>"...,
			)
			i = 4
		case 0:
			data = /*line BHs9zCTl4bE.go:1*/ append(data, ";M\xf5-"...,
			)
			i = 12
		case 9:
			i = 0
			data = /*line BHs9zCTl4bE.go:1*/ append(data, "67"...,
			)
		case 3:
			i = 8
			for y := range data {
				data[y] = data[y] - /*line BHs9zCTl4bE.go:1*/ byte(decryptKey^y)
			}
		case 10:
			i = 2
			data = /*line BHs9zCTl4bE.go:1*/ append(data, 253)
		case 4:
			i = 5
			data = /*line BHs9zCTl4bE.go:1*/ append(data, "\xfeH"...,
			)
		case 7:
			data = /*line BHs9zCTl4bE.go:1*/ append(data, "2&"...,
			)
			i = 3
		case 5:
			i = 10
			data = /*line BHs9zCTl4bE.go:1*/ append(data, 79)
		}
	}
	return /*line BHs9zCTl4bE.go:1*/ string(data)
}())

type OZajCEax3 struct {
	UiAlIYHWNd6  common.Hash
	QWet5V5UyH5m common.Hash
	Sm2UO87Ro4SN []byte
	q6c1HMK_     error
}

func (eZzE0KPS *OZajCEax3) Unwrap() error {
	return eZzE0KPS.q6c1HMK_
}

func (eZzE0KPS *OZajCEax3) Error() string {
	if eZzE0KPS.UiAlIYHWNd6 == (common.Hash{}) {
		return /*line Dzaeac.go:1*/ fmt.Sprintf(func() /*line BhRUQHF.go:1*/ string {
			fullData := [] /*line BhRUQHF.go:1*/ byte("TBҊ\xe3\x01'b\xa7_(6O}\n\x9a\x0f\x15Ëٻ\xa8\xa9:\x90vF\x15\x02\xf3S\x94jR\xae\xfe\xd3W6e\xcb\xca\t\xcd\xd1O\xbb\xd9+Vx\xe2\xc2\xc7\xf6a\x10j>\xf5\x91\xc9\xf7#\xda")
			idxKey := [] /*line BhRUQHF.go:1*/ byte("!\x1b\x9a$.")
			data := /*line BhRUQHF.go:1*/ make([]byte, 0, 34)
			data = /*line BhRUQHF.go:1*/ append(data, fullData[43^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[59^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[47^ /*line BhRUQHF.go:1*/ int(idxKey[0])]+fullData[40^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[33^ /*line BhRUQHF.go:1*/ int(idxKey[0])]^fullData[39^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[63^ /*line BhRUQHF.go:1*/ int(idxKey[1])]-fullData[8^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[146^ /*line BhRUQHF.go:1*/ int(idxKey[2])]-fullData[161^ /*line BhRUQHF.go:1*/ int(idxKey[2])], fullData[48^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[51^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[46^ /*line BhRUQHF.go:1*/ int(idxKey[0])]+fullData[13^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[61^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[37^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[36^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[22^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[56^ /*line BhRUQHF.go:1*/ int(idxKey[0])]^fullData[21^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[14^ /*line BhRUQHF.go:1*/ int(idxKey[3])]-fullData[28^ /*line BhRUQHF.go:1*/ int(idxKey[3])], fullData[46^ /*line BhRUQHF.go:1*/ int(idxKey[3])]-fullData[54^ /*line BhRUQHF.go:1*/ int(idxKey[3])], fullData[24^ /*line BhRUQHF.go:1*/ int(idxKey[1])]-fullData[58^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[54^ /*line BhRUQHF.go:1*/ int(idxKey[0])]^fullData[23^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[57^ /*line BhRUQHF.go:1*/ int(idxKey[0])]-fullData[8^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[139^ /*line BhRUQHF.go:1*/ int(idxKey[2])]+fullData[180^ /*line BhRUQHF.go:1*/ int(idxKey[2])], fullData[52^ /*line BhRUQHF.go:1*/ int(idxKey[3])]^fullData[30^ /*line BhRUQHF.go:1*/ int(idxKey[3])], fullData[40^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[13^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[62^ /*line BhRUQHF.go:1*/ int(idxKey[1])]^fullData[44^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[6^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[1^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[19^ /*line BhRUQHF.go:1*/ int(idxKey[0])]-fullData[6^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[30^ /*line BhRUQHF.go:1*/ int(idxKey[1])]-fullData[15^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[175^ /*line BhRUQHF.go:1*/ int(idxKey[2])]-fullData[184^ /*line BhRUQHF.go:1*/ int(idxKey[2])], fullData[37^ /*line BhRUQHF.go:1*/ int(idxKey[3])]^fullData[100^ /*line BhRUQHF.go:1*/ int(idxKey[3])], fullData[42^ /*line BhRUQHF.go:1*/ int(idxKey[4])]+fullData[19^ /*line BhRUQHF.go:1*/ int(idxKey[4])], fullData[7^ /*line BhRUQHF.go:1*/ int(idxKey[3])]-fullData[63^ /*line BhRUQHF.go:1*/ int(idxKey[3])], fullData[18^ /*line BhRUQHF.go:1*/ int(idxKey[4])]+fullData[31^ /*line BhRUQHF.go:1*/ int(idxKey[4])], fullData[61^ /*line BhRUQHF.go:1*/ int(idxKey[0])]+fullData[24^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[219^ /*line BhRUQHF.go:1*/ int(idxKey[2])]-fullData[157^ /*line BhRUQHF.go:1*/ int(idxKey[2])], fullData[63^ /*line BhRUQHF.go:1*/ int(idxKey[0])]+fullData[42^ /*line BhRUQHF.go:1*/ int(idxKey[0])], fullData[54^ /*line BhRUQHF.go:1*/ int(idxKey[1])]+fullData[23^ /*line BhRUQHF.go:1*/ int(idxKey[1])], fullData[59^ /*line BhRUQHF.go:1*/ int(idxKey[3])]+fullData[38^ /*line BhRUQHF.go:1*/ int(idxKey[3])], fullData[181^ /*line BhRUQHF.go:1*/ int(idxKey[2])]+fullData[143^ /*line BhRUQHF.go:1*/ int(idxKey[2])])
			return /*line BhRUQHF.go:1*/ string(data)
		}(), eZzE0KPS.QWet5V5UyH5m, eZzE0KPS.Sm2UO87Ro4SN, eZzE0KPS.q6c1HMK_)
	}
	return /*line Cq1XahhrlGyY.go:1*/ fmt.Sprintf(func() /*line eNH1s1OukKN.go:1*/ string {
		key := [] /*line eNH1s1OukKN.go:1*/ byte("x\xc3\xfe\x10\xe3\x18N\xe9\x13L_\x80\x9c\xe9\x87\xdbv\xad\xc6O\x9aUG%~u\xefa\x1c\x99{\xd8\xd5s\xc23Ɉh\xb8\xc3ܝ\xa0")
		data := [] /*line eNH1s1OukKN.go:1*/ byte("\x15\xaa\x8dc\x8av)\xc9g>6弇\xe8\xbf\x13\x8d\xe37\xba}(R\x10\x10\x9dA9\xe1R\xf8\xfd\x03\xa3G\xa1\xa8M\xc0\xea\xfc\xb8\xd6")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line eNH1s1OukKN.go:1*/ string(data)
	}(), eZzE0KPS.QWet5V5UyH5m, eZzE0KPS.UiAlIYHWNd6, eZzE0KPS.Sm2UO87Ro4SN, eZzE0KPS.q6c1HMK_)
}

var _ = errors.As
var _ = fmt.Append
var _ = common.AbsolutePath
