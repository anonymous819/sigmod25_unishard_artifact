//line :1
package snapshot

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

type r5WWYA_ struct {
	cx6PS_f1JyN  ethdb.Iterator
	kuI1imakf    []byte
	m1F2JPPlH1E2 []byte
	kalBe7jI77r4 bool
}

func szfg10dIoV(cu8ZKpmK ethdb.Iterator) *r5WWYA_ {
	return &r5WWYA_{cx6PS_f1JyN: cu8ZKpmK}
}

func (cu8ZKpmK *r5WWYA_) Hold() {
	if /*line BGLj635J.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Key() == nil {
		return
	}
	cu8ZKpmK.kuI1imakf = /*line VKY0MXJKYIR.go:1*/ common.CopyBytes( /*line UHcMSe4GKnO.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Key())
	cu8ZKpmK.m1F2JPPlH1E2 = /*line IqfWFUvp.go:1*/ common.CopyBytes( /*line mwsHlj.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Value())
	cu8ZKpmK.kalBe7jI77r4 = false
}

func (cu8ZKpmK *r5WWYA_) Next() bool {
	if !cu8ZKpmK.kalBe7jI77r4 && cu8ZKpmK.kuI1imakf != nil {
		cu8ZKpmK.kalBe7jI77r4 = true
	} else if cu8ZKpmK.kalBe7jI77r4 {
		cu8ZKpmK.kalBe7jI77r4 = false
		cu8ZKpmK.kuI1imakf = nil
		cu8ZKpmK.m1F2JPPlH1E2 = nil
	}
	if cu8ZKpmK.kuI1imakf != nil {
		return true
	}
	return /*line FAJV3rawfh7a.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Next()
}

func (cu8ZKpmK *r5WWYA_) Error() error { return /*line zahmM2.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Error() }

func (cu8ZKpmK *r5WWYA_) Release() {
	cu8ZKpmK.kalBe7jI77r4 = false
	cu8ZKpmK.kuI1imakf = nil
	cu8ZKpmK.m1F2JPPlH1E2 = nil
	/*line DAWLfevl_d.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Release()
}

func (cu8ZKpmK *r5WWYA_) Key() []byte {
	if cu8ZKpmK.kuI1imakf != nil {
		return cu8ZKpmK.kuI1imakf
	}
	return /*line PAMUGqDl.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Key()
}

func (cu8ZKpmK *r5WWYA_) Value() []byte {
	if cu8ZKpmK.m1F2JPPlH1E2 != nil {
		return cu8ZKpmK.m1F2JPPlH1E2
	}
	return /*line L3iWQYzybd3h.go:1*/ cu8ZKpmK.cx6PS_f1JyN.Value()
}

var _ = common.AbsolutePath
var _ ethdb.AncientReader
