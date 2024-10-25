//line :1
package types

import (
	"bytes"
	"fmt"
	"math"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

var f3SPqrhaMlU = sync.Pool{
	New: func() interface{} { return /*line FCSmIRNxr.go:1*/ sha3.NewLegacyKeccak256() },
}

var fgxbiq = sync.Pool{
	New: func() interface{} { return /*line wnsOQlJj.go:1*/ new(bytes.Buffer) },
}

func cGon36m7lTS(rMe9pIag uint64) ([]byte, *bytes.Buffer, error) {
	if rMe9pIag > math.MaxInt {
		return nil, nil /*line Ngb5me3mTo2o.go:1*/, fmt.Errorf(func() /*line Wh6AhkT8Fsp.go:1*/ string {
			seed := /*line Wh6AhkT8Fsp.go:1*/ byte(71)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line Wh6AhkT8Fsp.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line Wh6AhkT8Fsp.go:1*/ fnc(36)(10)(27)(183)(51)(90)(179)(226)(29)(166)(78)(15)(239)(30)(243)(251)(164)(71)(9)(88)(163)(26)(247)(225)(69)(143)(93)
			return /*line Wh6AhkT8Fsp.go:1*/ string(data)
		}(), rMe9pIag)
	}
	djAcxNMM5s6 := /*line eRop482De.go:1*/ fgxbiq.Get().(*bytes.Buffer)
	/*line CGZToWcl9A.go:1*/ djAcxNMM5s6.Reset()
	/*line za6Kf6Aw25.go:1*/ djAcxNMM5s6.Grow( /*line JzaGpWnBfm.go:1*/ int(rMe9pIag))
	aYao5YS := /*line t27o93J.go:1*/ djAcxNMM5s6.Bytes()[: /*line uMAaIz1C5.go:1*/ int(rMe9pIag)]
	return aYao5YS, djAcxNMM5s6, nil
}

func uWZWEzDAB(uO81zUpF interface{}) (hP0dFanm4 common.Hash) {
	cH4lnex0Ibo := /*line WB6aR3zE.go:1*/ f3SPqrhaMlU.Get().(crypto.KeccakState)
	defer /*line gss0ol4i.go:1*/ f3SPqrhaMlU.Put(cH4lnex0Ibo)
	/*line iD8zIjByqNMI.go:1*/ cH4lnex0Ibo.Reset()
	/*line FNoOzrS2xM.go:1*/ rlp.Encode(cH4lnex0Ibo, uO81zUpF)
	/*line OIkgSy0Qd4xR.go:1*/ cH4lnex0Ibo.Read(hP0dFanm4[:])
	return hP0dFanm4
}

func wrj50W(gX0mWyd1Ly byte, uO81zUpF interface{}) (hP0dFanm4 common.Hash) {
	cH4lnex0Ibo := /*line F3Zs7z.go:1*/ f3SPqrhaMlU.Get().(crypto.KeccakState)
	defer /*line tB0NBj.go:1*/ f3SPqrhaMlU.Put(cH4lnex0Ibo)
	/*line ReGfnIFfie.go:1*/ cH4lnex0Ibo.Reset()
	/*line mYmcnkqnbsm.go:1*/ cH4lnex0Ibo.Write([]byte{gX0mWyd1Ly})
	/*line UyYyRIMfi.go:1*/ rlp.Encode(cH4lnex0Ibo, uO81zUpF)
	/*line y49pewdYfzyk.go:1*/ cH4lnex0Ibo.Read(hP0dFanm4[:])
	return hP0dFanm4
}

type ZGJpq7HtE interface {
	Reset()
	Update([]byte, []byte) error
	Hash() common.Hash
}

type ItPhZaPh interface {
	Len() int
	EncodeIndex(int, *bytes.Buffer)
}

func jbAmeY2C(uqEGc5a ItPhZaPh, ibcOzM6f int, djAcxNMM5s6 *bytes.Buffer) []byte {
	/*line FPg2Tr.go:1*/ djAcxNMM5s6.Reset()
	/*line LDOwR4Oea.go:1*/ uqEGc5a.EncodeIndex(ibcOzM6f, djAcxNMM5s6)

	return /*line IyMO7ffcz.go:1*/ common.CopyBytes( /*line AA0Obi1Z.go:1*/ djAcxNMM5s6.Bytes())
}

func HXf5z8W_(uqEGc5a ItPhZaPh, qs4y7ialyr ZGJpq7HtE) common.Hash {
	/*line dHIztyr.go:1*/ qs4y7ialyr.Reset()

	btsGSgwXXhK := /*line G8wxa9S.go:1*/ fgxbiq.Get().(*bytes.Buffer)
	defer /*line WLHSvCdU.go:1*/ fgxbiq.Put(btsGSgwXXhK)

	var iYSu7v []byte
	for ibcOzM6f := 1; ibcOzM6f < /*line Xw0aGGHF1.go:1*/ uqEGc5a.Len() && ibcOzM6f <= 0x7f; ibcOzM6f++ {
		iYSu7v = /*line mRSvAABCH4aY.go:1*/ rlp.AppendUint64(iYSu7v[:0] /*line oPaNLsJAGCs.go:1*/, uint64(ibcOzM6f))
		fMqtswJ := /*line D90wJu.go:1*/ jbAmeY2C(uqEGc5a, ibcOzM6f, btsGSgwXXhK)
		/*line LgxECowe8Rlp.go:1*/ qs4y7ialyr.Update(iYSu7v, fMqtswJ)
	}
	if /*line pK6BMSl_.go:1*/ uqEGc5a.Len() > 0 {
		iYSu7v = /*line azJenMHhkI.go:1*/ rlp.AppendUint64(iYSu7v[:0], 0)
		fMqtswJ := /*line ysPqJTVqPo.go:1*/ jbAmeY2C(uqEGc5a, 0, btsGSgwXXhK)
		/*line GRMGCLQ.go:1*/ qs4y7ialyr.Update(iYSu7v, fMqtswJ)
	}
	for ibcOzM6f := 0x80; ibcOzM6f < /*line PHRgAhF7Oz.go:1*/ uqEGc5a.Len(); ibcOzM6f++ {
		iYSu7v = /*line se6jgIlK.go:1*/ rlp.AppendUint64(iYSu7v[:0] /*line kt5Wwq.go:1*/, uint64(ibcOzM6f))
		fMqtswJ := /*line BKVGoUn7.go:1*/ jbAmeY2C(uqEGc5a, ibcOzM6f, btsGSgwXXhK)
		/*line dV2R8enf.go:1*/ qs4y7ialyr.Update(iYSu7v, fMqtswJ)
	}
	return /*line CATDWsR.go:1*/ qs4y7ialyr.Hash()
}

var _ bytes.Buffer
var _ = fmt.Append
var _ = math.Abs
var _ sync.Cond
var _ = common.AbsolutePath
var _ = crypto.CompressPubkey
var _ = rlp.AppendUint64
var _ = sha3.New224
