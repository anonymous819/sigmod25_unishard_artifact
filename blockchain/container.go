//line :1
package blockchain

import "github.com/ethereum/go-ethereum/common"

type HlSaUqroF struct {
	QaRt8EunJ O0GQK9U1
}

func (jx2s151m1 *HlSaUqroF) VertexID() common.Hash {
	return /*line KUaj1KLbM.go:1*/ jx2s151m1.QaRt8EunJ.GetBlockHash()
}

func (jx2s151m1 *HlSaUqroF) WorkerBlockLevel() uint64 {
	return /*line JnHpIdxz39w1.go:1*/ uint64( /*line LwYeryWrM5.go:1*/ jx2s151m1.QaRt8EunJ.WorkerBlockHeader().BlockHeight)
}

func (jx2s151m1 *HlSaUqroF) CoordinationBlockLevel() uint64 {
	return /*line EWZjBXSn1vay.go:1*/ uint64( /*line i7JJq64lN4NR.go:1*/ jx2s151m1.QaRt8EunJ.CoordinationBlockHeader().BlockHeight)
}

func (jx2s151m1 *HlSaUqroF) ParentWorkerBlock() (common.Hash, uint64) {
	return /*line AbsCS8zJb7Ci.go:1*/ jx2s151m1.QaRt8EunJ.WorkerBlockHeader().GvMaqmnWSRO /*line bganN3J7.go:1*/, uint64( /*line fU4O4oqrLCIS.go:1*/ jx2s151m1.QaRt8EunJ.QuorumCertificate().BlockHeight)
}
func (jx2s151m1 *HlSaUqroF) ParentCoordinationBlock() (common.Hash, uint64) {
	return /*line lIz_6e.go:1*/ jx2s151m1.QaRt8EunJ.CoordinationBlockHeader().PrevBlockHash /*line dptMbie.go:1*/, uint64( /*line ZrKAsVm3S_n.go:1*/ jx2s151m1.QaRt8EunJ.QuorumCertificate().BlockHeight)
}

func (jx2s151m1 *HlSaUqroF) GetBlock() O0GQK9U1 { return jx2s151m1.QaRt8EunJ }

var _ = common.AbsolutePath
