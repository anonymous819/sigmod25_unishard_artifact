//line :1
package snapshot

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

type FcPREnURc interface {
	Next() bool

	Error() error

	Hash() common.Hash

	Release()
}

type Nq4YsH_ interface {
	FcPREnURc

	Account() []byte
}

type PE_7UyJghxy interface {
	FcPREnURc

	Slot() []byte
}

type diffAccountIterator struct {
	curHash common.Hash

	layer *diffLayer
	keys  []common.Hash
	fail  error
}

func (hbsnnIi *diffLayer) AccountIterator(gnGBaeoLX common.Hash) Nq4YsH_ {

	vvERYE8 := /*line zh53jKEvQYA8.go:1*/ hbsnnIi.AccountList()
	as3oS5CqGh := /*line Y7gL5m6RhOmn.go:1*/ sort.Search( /*line yeiWliZX.go:1*/ len(vvERYE8), func(fSpMhzHR int) bool {
		return /*line rVGda6diJ.go:1*/ bytes.Compare(gnGBaeoLX[:], vvERYE8[fSpMhzHR][:]) <= 0
	})

	return &diffAccountIterator{
		layer: hbsnnIi,
		keys:  vvERYE8[as3oS5CqGh:],
	}
}

func (cu8ZKpmK *diffAccountIterator) Next() bool {

	if cu8ZKpmK.fail != nil {
		/*line kTna01.go:1*/ panic( /*line XZTOLEVyM.go:1*/ fmt.Sprintf(func() /*line RK3cksyIL.go:1*/ string {
			data := [] /*line RK3cksyIL.go:1*/ byte("call\x17\xea &\xd8\xd2\xdd\x1e\xe4T\xb6faiWe\xc7r:tera\xbeo(\xa1\xe6\xaav")
			positions := [...]byte{11, 30, 18, 12, 20, 22, 13, 27, 4, 20, 8, 10, 29, 21, 4, 31, 9, 13, 11, 11, 14, 30, 32, 7, 11, 7, 5, 12, 11, 14, 32, 27, 8, 9, 9, 29, 31, 20, 21, 11, 32, 27}
			for i := 0; i < 42; i += 2 {
				localKey := /*line RK3cksyIL.go:1*/ byte(i) + /*line RK3cksyIL.go:1*/ byte(positions[i]^positions[i+1]) + 88
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line RK3cksyIL.go:1*/ string(data)
		}(), cu8ZKpmK.fail))
	}

	if /*line OpysVona7.go:1*/ len(cu8ZKpmK.keys) == 0 {
		return false
	}
	if /*line lhCrQQjrG_GA.go:1*/ cu8ZKpmK.layer.Stale() {
		cu8ZKpmK.fail, cu8ZKpmK.keys = AiK5hkDaW, nil
		return false
	}

	cu8ZKpmK.curHash = cu8ZKpmK.keys[0]

	cu8ZKpmK.keys = cu8ZKpmK.keys[1:]
	return true
}

func (cu8ZKpmK *diffAccountIterator) Error() error {
	return cu8ZKpmK.fail
}

func (cu8ZKpmK *diffAccountIterator) Hash() common.Hash {
	return cu8ZKpmK.curHash
}

func (cu8ZKpmK *diffAccountIterator) Account() []byte {
	/*line sIApsjw8goXZ.go:1*/ cu8ZKpmK.layer.lock.RLock()
	b1jafJL, jdkNTRyBJ := cu8ZKpmK.layer.accountData[cu8ZKpmK.curHash]
	if !jdkNTRyBJ {
		if _, jdkNTRyBJ := cu8ZKpmK.layer.destructSet[cu8ZKpmK.curHash]; jdkNTRyBJ {
			/*line JCOf5O_l9l.go:1*/ cu8ZKpmK.layer.lock.RUnlock()
			return nil
		}
		/*line a5KLswOXBfN.go:1*/ panic( /*line Efz68s94.go:1*/ fmt.Sprintf(func() /*line OhcrmWAij.go:1*/ string {
			data := /*line OhcrmWAij.go:1*/ make([]byte, 0, 45)
			i := 0
			decryptKey := 71
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					i = 19
					data = /*line OhcrmWAij.go:1*/ append(data, "s\xb0"...,
					)
				case 13:
					i = 2
					data = /*line OhcrmWAij.go:1*/ append(data, "\xc2̦\xa9"...,
					)
				case 1:
					data = /*line OhcrmWAij.go:1*/ append(data, "굜\x9e"...,
					)
					i = 18
				case 3:
					data = /*line OhcrmWAij.go:1*/ append(data, 200)
					i = 5
				case 11:
					for y := range data {
						data[y] = data[y] + /*line OhcrmWAij.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 15:
					data = /*line OhcrmWAij.go:1*/ append(data, 196)
					i = 3
				case 14:
					data = /*line OhcrmWAij.go:1*/ append(data, "\xd5\xe6\xed\xe3"...,
					)
					i = 1
				case 2:
					i = 12
					data = /*line OhcrmWAij.go:1*/ append(data, 165)
				case 16:
					data = /*line OhcrmWAij.go:1*/ append(data, "\xb8\xcc"...,
					)
					i = 15
				case 10:
					i = 8
					data = /*line OhcrmWAij.go:1*/ append(data, "õ"...,
					)
				case 18:
					i = 11
					data = /*line OhcrmWAij.go:1*/ append(data, 242)
				case 19:
					i = 10
					data = /*line OhcrmWAij.go:1*/ append(data, "Ĳ\xbd"...,
					)
				case 8:
					data = /*line OhcrmWAij.go:1*/ append(data, "\xbb\u0093\xd5"...,
					)
					i = 9
				case 7:
					data = /*line OhcrmWAij.go:1*/ append(data, "\xc4\xd2"...,
					)
					i = 13
				case 9:
					data = /*line OhcrmWAij.go:1*/ append(data, 212)
					i = 14
				case 17:
					data = /*line OhcrmWAij.go:1*/ append(data, 192)
					i = 7
				case 12:
					i = 4
					data = /*line OhcrmWAij.go:1*/ append(data, "b\xb5\xb7\xb3"...,
					)
				case 0:
					i = 16
					data = /*line OhcrmWAij.go:1*/ append(data, "\xbcȶ\xc4"...,
					)
				case 5:
					data = /*line OhcrmWAij.go:1*/ append(data, "{ξ"...,
					)
					i = 17
				}
			}
			return /*line OhcrmWAij.go:1*/ string(data)
		}(), cu8ZKpmK.curHash))
	}
	/*line as7JNxjRwO9.go:1*/ cu8ZKpmK.layer.lock.RUnlock()
	if /*line gw6kNULE_.go:1*/ cu8ZKpmK.layer.Stale() {
		cu8ZKpmK.fail, cu8ZKpmK.keys = AiK5hkDaW, nil
	}
	return b1jafJL
}

func (cu8ZKpmK *diffAccountIterator) Release() {}

type uwdOD3Fd struct {
	cJV_ZD4D   *diskLayer
	lQA3hTamBK ethdb.Iterator
}

func (hbsnnIi *diskLayer) AccountIterator(gnGBaeoLX common.Hash) Nq4YsH_ {
	aBqfVt := /*line v_9Vj1.go:1*/ common.TrimRightZeroes(gnGBaeoLX[:])
	return &uwdOD3Fd{
		cJV_ZD4D: hbsnnIi,
		lQA3hTamBK:/*line qTkIbRaqR.go:1*/ hbsnnIi.diskdb.NewIterator(rawdb.SnapshotAccountPrefix, aBqfVt),
	}
}

func (cu8ZKpmK *uwdOD3Fd) Next() bool {

	if cu8ZKpmK.lQA3hTamBK == nil {
		return false
	}

	for {
		if ! /*line aXg_sROaQz5.go:1*/ cu8ZKpmK.lQA3hTamBK.Next() {
			/*line FkWwknD.go:1*/ cu8ZKpmK.lQA3hTamBK.Release()
			cu8ZKpmK.lQA3hTamBK = nil
			return false
		}
		if /*line CChFWP.go:1*/ len( /*line tL4uay.go:1*/ cu8ZKpmK.lQA3hTamBK.Key()) == /*line r24JSkcrL.go:1*/ len(rawdb.SnapshotAccountPrefix)+common.HashLength {
			break
		}
	}
	return true
}

func (cu8ZKpmK *uwdOD3Fd) Error() error {
	if cu8ZKpmK.lQA3hTamBK == nil {
		return nil
	}
	return /*line BONiIbKCJ3.go:1*/ cu8ZKpmK.lQA3hTamBK.Error()
}

func (cu8ZKpmK *uwdOD3Fd) Hash() common.Hash {
	return /*line hvIttN5aLow.go:1*/ common.BytesToHash( /*line UUiS9qu9eT.go:1*/ cu8ZKpmK.lQA3hTamBK.Key())
}

func (cu8ZKpmK *uwdOD3Fd) Account() []byte {
	return /*line A3DESkaqO.go:1*/ cu8ZKpmK.lQA3hTamBK.Value()
}

func (cu8ZKpmK *uwdOD3Fd) Release() {

	if cu8ZKpmK.lQA3hTamBK != nil {
		/*line sG4F_QtQ.go:1*/ cu8ZKpmK.lQA3hTamBK.Release()
		cu8ZKpmK.lQA3hTamBK = nil
	}
}

type diffStorageIterator struct {
	curHash common.Hash
	account common.Hash

	layer *diffLayer
	keys  []common.Hash
	fail  error
}

func (hbsnnIi *diffLayer) StorageIterator(evzmhL1 common.Hash, gnGBaeoLX common.Hash) (PE_7UyJghxy, bool) {

	vvERYE8, oiHhcQbh9XAf := /*line I0PwoCWV.go:1*/ hbsnnIi.StorageList(evzmhL1)
	as3oS5CqGh := /*line AsIqEo.go:1*/ sort.Search( /*line MYhc1WDYM.go:1*/ len(vvERYE8), func(fSpMhzHR int) bool {
		return /*line IYKRfn3I59.go:1*/ bytes.Compare(gnGBaeoLX[:], vvERYE8[fSpMhzHR][:]) <= 0
	})

	return &diffStorageIterator{
		layer:   hbsnnIi,
		account: evzmhL1,
		keys:    vvERYE8[as3oS5CqGh:],
	}, oiHhcQbh9XAf
}

func (cu8ZKpmK *diffStorageIterator) Next() bool {

	if cu8ZKpmK.fail != nil {
		/*line BAUrxa3.go:1*/ panic( /*line QzKa2cLc1w.go:1*/ fmt.Sprintf(func() /*line skA1qu2L.go:1*/ string {
			data := [] /*line skA1qu2L.go:1*/ byte("\xa3\x8el ed \x9c\x9fxۖof܌auOe\x9d \x8a\x92e\x8batD\bg\xb1g\xd7")
			positions := [...]byte{32, 17, 31, 18, 30, 33, 32, 10, 30, 14, 28, 29, 13, 23, 23, 15, 31, 23, 29, 1, 25, 20, 7, 22, 0, 7, 3, 8, 17, 0, 12, 22, 28, 31, 11, 8}
			for i := 0; i < 36; i += 2 {
				localKey := /*line skA1qu2L.go:1*/ byte(i) + /*line skA1qu2L.go:1*/ byte(positions[i]^positions[i+1]) + 206
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line skA1qu2L.go:1*/ string(data)
		}(), cu8ZKpmK.fail))
	}

	if /*line mzYPEmtTN1.go:1*/ len(cu8ZKpmK.keys) == 0 {
		return false
	}
	if /*line PwpFDQ9_2awK.go:1*/ cu8ZKpmK.layer.Stale() {
		cu8ZKpmK.fail, cu8ZKpmK.keys = AiK5hkDaW, nil
		return false
	}

	cu8ZKpmK.curHash = cu8ZKpmK.keys[0]

	cu8ZKpmK.keys = cu8ZKpmK.keys[1:]
	return true
}

func (cu8ZKpmK *diffStorageIterator) Error() error {
	return cu8ZKpmK.fail
}

func (cu8ZKpmK *diffStorageIterator) Hash() common.Hash {
	return cu8ZKpmK.curHash
}

func (cu8ZKpmK *diffStorageIterator) Slot() []byte {
	/*line YoUzjXNk.go:1*/ cu8ZKpmK.layer.lock.RLock()
	agSpCMzc, jdkNTRyBJ := cu8ZKpmK.layer.storageData[cu8ZKpmK.account]
	if !jdkNTRyBJ {
		/*line BhfufkbSA.go:1*/ panic( /*line Cw9aXs.go:1*/ fmt.Sprintf(func() /*line Y2bhtMacrPzy.go:1*/ string {
			key := [] /*line Y2bhtMacrPzy.go:1*/ byte("\r\x0e\x9c\xfe\x008\xaeGFI\x94ߨ=\xc7f$N\x1e\x98T\xab\xfc\xc7\xee6z\xfb;\xa5\xb6q\xb7\xf9Ԋ:\xf1\x99ʗ\x13\fZ\x04\xcd\xc48@\xa9\xbc\x15")
			data := [] /*line Y2bhtMacrPzy.go:1*/ byte("dz\xf9\x8caL\xc15f;\xf1\xb9\xcdO\xa2\bG+z\xb8:Ē\xea\x8bN\x13\x88O\xc0\xd8\x05\x97\x98\xb7\xe9U\x84\xf7\xbe\xb7`x5v\xac\xa3]z\x89\x99m")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line Y2bhtMacrPzy.go:1*/ string(data)
		}(), cu8ZKpmK.account))
	}

	b1jafJL, jdkNTRyBJ := agSpCMzc[cu8ZKpmK.curHash]
	if !jdkNTRyBJ {
		/*line R4EuwhPW2.go:1*/ panic( /*line RtBvfCH6m7.go:1*/ fmt.Sprintf(func() /*line EACMSF.go:1*/ string {
			seed := /*line EACMSF.go:1*/ byte(160)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line EACMSF.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line EACMSF.go:1*/ fnc(201)(11)(241)(13)(239)(19)(251)(3)(174)(82)(243)(1)(255)(13)(243)(9)(245)(2)(255)(188)(78)(1)(255)(191)(56)(19)(241)(10)(1)(241)(9)(6)(172)(83)(1)(251)(3)(239)(6)(254)(187)(83)(249)(3)(5)(198)(230)(5)(83)
			return /*line EACMSF.go:1*/ string(data)
		}(), cu8ZKpmK.curHash))
	}
	/*line NW6Bm1.go:1*/ cu8ZKpmK.layer.lock.RUnlock()
	if /*line JJoSd9.go:1*/ cu8ZKpmK.layer.Stale() {
		cu8ZKpmK.fail, cu8ZKpmK.keys = AiK5hkDaW, nil
	}
	return b1jafJL
}

func (cu8ZKpmK *diffStorageIterator) Release() {}

type swgK1tJor struct {
	gYwAfTJ62yAe *diskLayer
	l4uSdvMRLp1  common.Hash
	gfHvlTv      ethdb.Iterator
}

func (hbsnnIi *diskLayer) StorageIterator(evzmhL1 common.Hash, gnGBaeoLX common.Hash) (PE_7UyJghxy, bool) {
	aBqfVt := /*line QsL4xTV37H.go:1*/ common.TrimRightZeroes(gnGBaeoLX[:])
	return &swgK1tJor{
		gYwAfTJ62yAe: hbsnnIi,
		l4uSdvMRLp1:  evzmhL1,
		gfHvlTv:/*line vY99ZQ.go:1*/ hbsnnIi.diskdb.NewIterator( /*line DOL6t9.go:1*/ append(rawdb.SnapshotStoragePrefix /*line HxvxCh1lJe0D.go:1*/, evzmhL1.Bytes()...), aBqfVt),
	}, false
}

func (cu8ZKpmK *swgK1tJor) Next() bool {

	if cu8ZKpmK.gfHvlTv == nil {
		return false
	}

	for {
		if ! /*line gUo2eX36EQW.go:1*/ cu8ZKpmK.gfHvlTv.Next() {
			/*line FrBclvmXL.go:1*/ cu8ZKpmK.gfHvlTv.Release()
			cu8ZKpmK.gfHvlTv = nil
			return false
		}
		if /*line yAalyrbR2x.go:1*/ len( /*line Sn6osBa5i753.go:1*/ cu8ZKpmK.gfHvlTv.Key()) == /*line fkCPva0.go:1*/ len(rawdb.SnapshotStoragePrefix)+common.HashLength+common.HashLength {
			break
		}
	}
	return true
}

func (cu8ZKpmK *swgK1tJor) Error() error {
	if cu8ZKpmK.gfHvlTv == nil {
		return nil
	}
	return /*line AjV1KAfxF.go:1*/ cu8ZKpmK.gfHvlTv.Error()
}

func (cu8ZKpmK *swgK1tJor) Hash() common.Hash {
	return /*line lOeSTAZ5R.go:1*/ common.BytesToHash( /*line tv2EtXWnJ4.go:1*/ cu8ZKpmK.gfHvlTv.Key())
}

func (cu8ZKpmK *swgK1tJor) Slot() []byte {
	return /*line KVcUSxD.go:1*/ cu8ZKpmK.gfHvlTv.Value()
}

func (cu8ZKpmK *swgK1tJor) Release() {

	if cu8ZKpmK.gfHvlTv != nil {
		/*line f4Aao3ozX.go:1*/ cu8ZKpmK.gfHvlTv.Release()
		cu8ZKpmK.gfHvlTv = nil
	}
}

var _ bytes.Buffer
var _ = fmt.Append
var _ = sort.Find
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
