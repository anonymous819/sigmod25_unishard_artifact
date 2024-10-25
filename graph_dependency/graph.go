//line :1
package graph_dependency

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

type EnUOfMJb2Rs struct {
	Zjo4ai   map[common.Hash]*ZTebsfJN
	CVgrY1RA map[common.Hash]int
	Up00xvsd map[common.Hash]int

	RvR6XAfo map[common.Hash][]ZTebsfJN
}

type ZTebsfJN struct {
	Rnf3h0N      common.Hash
	LU2OwRAg_    int
	Dp1dzRVI__Id map[common.Hash]*Yg0tyIM
	EKlrom_1     []string
	GjiISH       []string
}

type Yg0tyIM struct {
	DjAYL6McRG *ZTebsfJN
}
type ZaaIbDTS struct {
	UgLbBl  int
	XEGPBy3 int
}
type L3xCZgM []ZaaIbDTS

func (lIIutWz8Tb L3xCZgM) Len() int { return /*line Ogdp3Vg.go:1*/ len(lIIutWz8Tb) }
func (lIIutWz8Tb L3xCZgM) Less(mOeLHltB, lAl17O3L int) bool {
	return lIIutWz8Tb[mOeLHltB].XEGPBy3 < lIIutWz8Tb[lAl17O3L].XEGPBy3
}
func (lIIutWz8Tb L3xCZgM) Swap(mOeLHltB, lAl17O3L int) {
	lIIutWz8Tb[mOeLHltB], lIIutWz8Tb[lAl17O3L] = lIIutWz8Tb[lAl17O3L], lIIutWz8Tb[mOeLHltB]
}

func NewGraph() *EnUOfMJb2Rs {
	hQ5E7Uxw := &EnUOfMJb2Rs{Zjo4ai: map[common.Hash]*ZTebsfJN{}, CVgrY1RA: map[common.Hash]int{}, Up00xvsd: map[common.Hash]int{}}

	return hQ5E7Uxw
}

func (hQ5E7Uxw *EnUOfMJb2Rs) AddVertex(gwUbrwbq common.Hash, peUvaDCt5 int, aqEVaDQsSMdE []string, dxoIw86_NN []string) {
	hQ5E7Uxw.Zjo4ai[gwUbrwbq] = &ZTebsfJN{Rnf3h0N: gwUbrwbq, LU2OwRAg_: peUvaDCt5, Dp1dzRVI__Id: map[common.Hash]*Yg0tyIM{}, GjiISH: aqEVaDQsSMdE, EKlrom_1: dxoIw86_NN}
}

func (hQ5E7Uxw *EnUOfMJb2Rs) DrawDependency() {
	for _, lLYhSAaTx3xF := range hQ5E7Uxw.Zjo4ai {
		for _, xwdQGJWqXBlw := range hQ5E7Uxw.Zjo4ai {
			if lLYhSAaTx3xF.LU2OwRAg_ < xwdQGJWqXBlw.LU2OwRAg_ {
				if /*line u1hpui.go:1*/ fCqQNeT(lLYhSAaTx3xF.GjiISH, xwdQGJWqXBlw.EKlrom_1) {
					/*line sAQFAKEQi1lO.go:1*/ hQ5E7Uxw.AddEdge(xwdQGJWqXBlw.Rnf3h0N, lLYhSAaTx3xF.Rnf3h0N)
				}
			}
		}
	}
}

func (hQ5E7Uxw *EnUOfMJb2Rs) AddEdge(h3jXv2 common.Hash, vEXq1J7ekuYN common.Hash) {
	if _, w7NucW5 := hQ5E7Uxw.Zjo4ai[h3jXv2]; !w7NucW5 {
		return
	}
	if _, w7NucW5 := hQ5E7Uxw.Zjo4ai[vEXq1J7ekuYN]; !w7NucW5 {
		return
	}

	hQ5E7Uxw.Zjo4ai[h3jXv2].Dp1dzRVI__Id[vEXq1J7ekuYN] = &Yg0tyIM{DjAYL6McRG: hQ5E7Uxw.Zjo4ai[vEXq1J7ekuYN]}
	hQ5E7Uxw.CVgrY1RA[h3jXv2] += 1
	hQ5E7Uxw.Up00xvsd[vEXq1J7ekuYN] += 1
}

func (hQ5E7Uxw *EnUOfMJb2Rs) Neighbors(h3jXv2 common.Hash) []common.Hash {
	vXwlj7sly := []common.Hash{}

	if _, w7NucW5 := hQ5E7Uxw.Zjo4ai[h3jXv2]; !w7NucW5 {
		return vXwlj7sly
	}

	for _, zdFhR6 := range hQ5E7Uxw.Zjo4ai[h3jXv2].Dp1dzRVI__Id {
		vXwlj7sly = /*line eLENaE__.go:1*/ append(vXwlj7sly, zdFhR6.DjAYL6McRG.Rnf3h0N)
	}

	return vXwlj7sly
}

func (hQ5E7Uxw *EnUOfMJb2Rs) GetEnteringEdges(gwUbrwbq common.Hash) int {
	return hQ5E7Uxw.Up00xvsd[gwUbrwbq]
}

func (hQ5E7Uxw *EnUOfMJb2Rs) GetLeavingEdges(gwUbrwbq common.Hash) int {
	return hQ5E7Uxw.CVgrY1RA[gwUbrwbq]
}

func (hQ5E7Uxw *EnUOfMJb2Rs) GetVertices() []common.Hash {
	vXwlj7sly := []common.Hash{}
	for _, jrMBt6Qvyef := range hQ5E7Uxw.Zjo4ai {
		if ! /*line lDLr9G.go:1*/ slices.Contains(vXwlj7sly, jrMBt6Qvyef.Rnf3h0N) {
			vXwlj7sly = /*line LeP9HJn.go:1*/ append(vXwlj7sly, jrMBt6Qvyef.Rnf3h0N)
		}
	}

	return vXwlj7sly
}

func fCqQNeT(rk8WiS9uE []string, qkwkMHHcS []string) bool {
	for _, y7sfM1V_MgVE := range rk8WiS9uE {
		for _, vNxUYuhJxEPG := range qkwkMHHcS {
			if y7sfM1V_MgVE == vNxUYuhJxEPG {
				return true
			}
		}
	}
	return false
}

var _ = common.AbsolutePath
