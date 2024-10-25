//line :1
package blockchain

import "github.com/ethereum/go-ethereum/common"

type S9fI7Wj interface {
	VertexID() common.Hash

	WorkerBlockLevel() uint64
	CoordinationBlockLevel() uint64

	ParentWorkerBlock() (common.Hash, uint64)
	ParentCoordinationBlock() (common.Hash, uint64)

	GetBlock() O0GQK9U1
}

type ZOPXV6ybU struct {
	wizeHT       BJtoC0zGWx05
	f6UGFTwqF1st int
	zoMhtKr      S9fI7Wj
}

func (x4tCF3OaKG *ZOPXV6ybU) pCPuwkk() {
	for x4tCF3OaKG.f6UGFTwqF1st < /*line DHU3JDq.go:1*/ len(x4tCF3OaKG.wizeHT) {
		h_5_caanZ := x4tCF3OaKG.wizeHT[x4tCF3OaKG.f6UGFTwqF1st].bv0LsjY5
		x4tCF3OaKG.f6UGFTwqF1st++
		if h_5_caanZ != nil {
			x4tCF3OaKG.zoMhtKr = h_5_caanZ
			return
		}
	}
	x4tCF3OaKG.zoMhtKr = nil
}

func (x4tCF3OaKG *ZOPXV6ybU) NextVertex() S9fI7Wj {
	joHO51TkdLTO := x4tCF3OaKG.zoMhtKr
	/*line UHBWHVin9xi0.go:1*/ x4tCF3OaKG.pCPuwkk()
	return joHO51TkdLTO
}

func (x4tCF3OaKG *ZOPXV6ybU) HasNext() bool {
	return x4tCF3OaKG.zoMhtKr != nil
}

func qP0V4i3Ogp_(jW9hHXhqFYY8 BJtoC0zGWx05) ZOPXV6ybU {
	x4tCF3OaKG := ZOPXV6ybU{
		wizeHT: jW9hHXhqFYY8,
	}
	/*line QxknCL0.go:1*/ x4tCF3OaKG.pCPuwkk()
	return x4tCF3OaKG
}

var _ = common.AbsolutePath
