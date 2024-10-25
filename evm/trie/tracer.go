//line :1
package trie

import (
	"github.com/ethereum/go-ethereum/common"
)

type tracer struct {
	inserts    map[string]struct{}
	deletes    map[string]struct{}
	accessList map[string][]byte
}

func aLjRi9a0() *tracer {
	return &tracer{
		inserts:/*line wwbW2suf.go:1*/ make(map[string]struct{}),
		deletes:/*line SzYb120nUlT.go:1*/ make(map[string]struct{}),
		accessList:/*line dgTavusJ35Lb.go:1*/ make(map[string][]byte),
	}
}

func (hkI2UXG_QBd *tracer) d5Xj5n(qiRG6lpeaFl []byte, h_pxYek4zT []byte) {
	hkI2UXG_QBd.accessList[ /*line NOKRyy1Hazed.go:1*/ string(qiRG6lpeaFl)] = h_pxYek4zT
}

func (hkI2UXG_QBd *tracer) har2k8rr3s(qiRG6lpeaFl []byte) {
	if _, pSKuSIB := hkI2UXG_QBd.deletes[ /*line V6eearSA.go:1*/ string(qiRG6lpeaFl)]; pSKuSIB {
		/*line k0v5jkmwSRj.go:1*/ delete(hkI2UXG_QBd.deletes /*line DfMa6nInH7.go:1*/, string(qiRG6lpeaFl))
		return
	}
	hkI2UXG_QBd.inserts[ /*line MEQIj7U.go:1*/ string(qiRG6lpeaFl)] = struct{}{}
}

func (hkI2UXG_QBd *tracer) xwjztifet3(qiRG6lpeaFl []byte) {
	if _, pSKuSIB := hkI2UXG_QBd.inserts[ /*line COau_V9l6.go:1*/ string(qiRG6lpeaFl)]; pSKuSIB {
		/*line QmXcVmA.go:1*/ delete(hkI2UXG_QBd.inserts /*line kbkdKdr.go:1*/, string(qiRG6lpeaFl))
		return
	}
	hkI2UXG_QBd.deletes[ /*line yHRImeHdeI9n.go:1*/ string(qiRG6lpeaFl)] = struct{}{}
}

func (hkI2UXG_QBd *tracer) x9PQbEXy() {
	hkI2UXG_QBd.inserts = /*line rDz3hRHR0JP.go:1*/ make(map[string]struct{})
	hkI2UXG_QBd.deletes = /*line rwSNduvYZZc.go:1*/ make(map[string]struct{})
	hkI2UXG_QBd.accessList = /*line WQrjSrH.go:1*/ make(map[string][]byte)
}

func (hkI2UXG_QBd *tracer) sBTiL7Ci() *tracer {
	var (
		g0m8ME     = /*line gwRfqyYfUNx.go:1*/ make(map[string]struct{})
		tfaheLZOu_ = /*line gWRDFrz.go:1*/ make(map[string]struct{})
		yxVEXPKP   = /*line pah9mqHp8K.go:1*/ make(map[string][]byte)
	)
	for qiRG6lpeaFl := range hkI2UXG_QBd.inserts {
		g0m8ME[qiRG6lpeaFl] = struct{}{}
	}
	for qiRG6lpeaFl := range hkI2UXG_QBd.deletes {
		tfaheLZOu_[qiRG6lpeaFl] = struct{}{}
	}
	for qiRG6lpeaFl, aBHduJG := range hkI2UXG_QBd.accessList {
		yxVEXPKP[qiRG6lpeaFl] = /*line sEO5Ar.go:1*/ common.CopyBytes(aBHduJG)
	}
	return &tracer{
		inserts:    g0m8ME,
		deletes:    tfaheLZOu_,
		accessList: yxVEXPKP,
	}
}

func (hkI2UXG_QBd *tracer) tnNWw3vS() []string {
	var k2TrdNkb []string
	for qiRG6lpeaFl := range hkI2UXG_QBd.deletes {

		_, yY_yPSd8vG := hkI2UXG_QBd.accessList[qiRG6lpeaFl]
		if !yY_yPSd8vG {
			continue
		}
		k2TrdNkb = /*line VzdAYbOI1.go:1*/ append(k2TrdNkb, qiRG6lpeaFl)
	}
	return k2TrdNkb
}

var _ = common.AbsolutePath
