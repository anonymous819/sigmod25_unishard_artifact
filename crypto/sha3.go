//line :1
package crypto

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

const (
	HashLenSha3_224 = 32
	HashLenSha3_256 = 32
	HashLenSha3_384 = 32
	HashLenSha3_512 = 32
)

type zISqBRgUDEs struct {
	*wmWoG2oB8SgA
	hash.Hash
}

type uGXVarOR74 struct {
	*wmWoG2oB8SgA
	hash.Hash
}

type ciNltK struct {
	*wmWoG2oB8SgA
	hash.Hash
}

type rsEVqUD struct {
	*wmWoG2oB8SgA
	hash.Hash
}

func GJNstJ9Xyk() K3fvz3wD {
	return &zISqBRgUDEs{
		wmWoG2oB8SgA: &wmWoG2oB8SgA{
			iqYYUpClDS5: HashLenSha3_224},
		Hash:/*line vDsa8k0EaSGD.go:1*/ sha3.New224()}
}

func OamOdX() K3fvz3wD {
	return &uGXVarOR74{
		wmWoG2oB8SgA: &wmWoG2oB8SgA{
			iqYYUpClDS5: HashLenSha3_256},
		Hash:/*line brMXRZR.go:1*/ sha3.New256()}
}

func JiUjJXe() K3fvz3wD {
	return &ciNltK{
		wmWoG2oB8SgA: &wmWoG2oB8SgA{
			iqYYUpClDS5: HashLenSha3_384},
		Hash:/*line o6_kC9s.go:1*/ sha3.New384()}
}

func XZMTba4() K3fvz3wD {
	return &rsEVqUD{
		wmWoG2oB8SgA: &wmWoG2oB8SgA{
			iqYYUpClDS5: HashLenSha3_512},
		Hash:/*line EUidTFeH.go:1*/ sha3.New512()}
}

func (eJtyIP *zISqBRgUDEs) ComputeHash(c5GIVS []byte) KDSldlLG06As {
	/*line rVbixRafror.go:1*/ eJtyIP.Reset()
	_, _ = /*line E_KmltqrA.go:1*/ eJtyIP.Write(c5GIVS)
	wUbR1hG := /*line dlM2epZrwk.go:1*/ make(KDSldlLG06As, 0, HashLenSha3_224)
	return /*line _aJetq.go:1*/ eJtyIP.Sum(wUbR1hG)
}

func (eJtyIP *uGXVarOR74) ComputeHash(c5GIVS []byte) KDSldlLG06As {
	/*line YAV2mCpPu.go:1*/ eJtyIP.Reset()
	_, _ = /*line cOi4zmNV85na.go:1*/ eJtyIP.Write(c5GIVS)
	wUbR1hG := /*line Xm11ECzz.go:1*/ make(KDSldlLG06As, 0, HashLenSha3_256)
	return /*line Ru5hY3LOdhth.go:1*/ eJtyIP.Sum(wUbR1hG)
}

func (eJtyIP *ciNltK) ComputeHash(c5GIVS []byte) KDSldlLG06As {
	/*line ZfTIW0.go:1*/ eJtyIP.Reset()
	_, _ = /*line RNDOrkCuJb5.go:1*/ eJtyIP.Write(c5GIVS)
	wUbR1hG := /*line saNLLJ0W7ZF.go:1*/ make(KDSldlLG06As, 0, HashLenSha3_384)
	return /*line PIDnWaFXJh.go:1*/ eJtyIP.Sum(wUbR1hG)
}

func (eJtyIP *rsEVqUD) ComputeHash(c5GIVS []byte) KDSldlLG06As {
	/*line ZXnD7w7R.go:1*/ eJtyIP.Reset()
	_, _ = /*line tznMs3q8.go:1*/ eJtyIP.Write(c5GIVS)
	wUbR1hG := /*line AvJ3sOl.go:1*/ make(KDSldlLG06As, 0, HashLenSha3_512)
	return /*line lgibUShrwW.go:1*/ eJtyIP.Sum(wUbR1hG)
}

func (eJtyIP *zISqBRgUDEs) SumHash() KDSldlLG06As {
	wUbR1hG := /*line nlk61tE.go:1*/ make(KDSldlLG06As, HashLenSha3_224)
	/*line w6gK6V.go:1*/ eJtyIP.Sum(wUbR1hG[:0])
	/*line mdP09NmcZI.go:1*/ eJtyIP.Reset()
	return wUbR1hG
}

func (eJtyIP *uGXVarOR74) SumHash() KDSldlLG06As {
	wUbR1hG := /*line b20dzW4X9_.go:1*/ make(KDSldlLG06As, HashLenSha3_256)
	/*line DxmJTw3Cwjh.go:1*/ eJtyIP.Sum(wUbR1hG[:0])
	/*line Chv3gHD.go:1*/ eJtyIP.Reset()
	return wUbR1hG
}

func (eJtyIP *ciNltK) SumHash() KDSldlLG06As {
	wUbR1hG := /*line azPvF8o.go:1*/ make(KDSldlLG06As, HashLenSha3_384)
	/*line b6p6hiszv.go:1*/ eJtyIP.Sum(wUbR1hG[:0])
	/*line harfjtqJgdF8.go:1*/ eJtyIP.Reset()
	return wUbR1hG
}

func (eJtyIP *rsEVqUD) SumHash() KDSldlLG06As {
	wUbR1hG := /*line uqwcPGFi.go:1*/ make(KDSldlLG06As, HashLenSha3_512)
	/*line TqVWmNG4e.go:1*/ eJtyIP.Sum(wUbR1hG[:0])
	/*line ty7MQfq.go:1*/ eJtyIP.Reset()
	return wUbR1hG
}

var _ hash.Hash
var _ = sha3.New224
