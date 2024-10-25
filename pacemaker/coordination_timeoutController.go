//line :1
package pacemaker

import (
	"sync"

	blockchain "unishard/blockchain"
	types "unishard/types"
)

type FjZVwJmDPT struct {
	jMiIWe      int
	lgXJXgaVPjr map[types.View]map[types.NodeID]*H8NP1AOKTF
	r4gRf7aAj0  sync.Mutex
}

func Kg1AusRa8_(iDvh6le2Hy7Q int) *FjZVwJmDPT {
	ffhG_AUcI := /*line GrK49z4.go:1*/ new(FjZVwJmDPT)
	ffhG_AUcI.jMiIWe = iDvh6le2Hy7Q
	ffhG_AUcI.lgXJXgaVPjr = /*line AAUoSE.go:1*/ make(map[types.View]map[types.NodeID]*H8NP1AOKTF)

	return ffhG_AUcI
}

func (ffhG_AUcI *FjZVwJmDPT) AddTmo(eO94A1 *H8NP1AOKTF) (bool, *RZ65hTAdXj, *blockchain.CoordinationBlock) {
	/*line tliCOYt9JSk.go:1*/ ffhG_AUcI.r4gRf7aAj0.Lock()
	defer /*line fZc1_dT_N.go:1*/ ffhG_AUcI.r4gRf7aAj0.Unlock()
	if /*line wni5dxN.go:1*/ ffhG_AUcI.rtp1gNaU_fz(eO94A1.C50bBOZxcYfV) {
		return false, nil, nil
	}
	_, dy9ax2z0dP := ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV]
	if !dy9ax2z0dP {

		ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV] = /*line ArfNFcafNa.go:1*/ make(map[types.NodeID]*H8NP1AOKTF)
	}
	ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV][eO94A1.RKtUCo69O] = eO94A1
	if /*line CA8h1zq8.go:1*/ ffhG_AUcI.rtp1gNaU_fz(eO94A1.C50bBOZxcYfV) {
		tkGFesoo, y5d2lf := /*line kMTayguC.go:1*/ JLIUrx3(eO94A1.C50bBOZxcYfV, ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV])
		return true, tkGFesoo, y5d2lf
	}

	return false, nil, nil
}

func (ffhG_AUcI *FjZVwJmDPT) AddMjorityTmo(eO94A1 *H8NP1AOKTF) (bool, *RZ65hTAdXj, *blockchain.CoordinationBlock) {
	/*line qVLju0I1D7.go:1*/ ffhG_AUcI.r4gRf7aAj0.Lock()
	defer /*line CvHYRtV5ItUX.go:1*/ ffhG_AUcI.r4gRf7aAj0.Unlock()
	if /*line vNbZm3femge.go:1*/ ffhG_AUcI.Majority(eO94A1.C50bBOZxcYfV) {
		return false, nil, nil
	}
	_, dy9ax2z0dP := ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV]
	if !dy9ax2z0dP {

		ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV] = /*line nz09d1G9fW.go:1*/ make(map[types.NodeID]*H8NP1AOKTF)
	}
	ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV][eO94A1.RKtUCo69O] = eO94A1
	if /*line CtI4o5.go:1*/ ffhG_AUcI.Majority(eO94A1.C50bBOZxcYfV) {
		tkGFesoo, y5d2lf := /*line rSloO2H.go:1*/ JLIUrx3(eO94A1.C50bBOZxcYfV, ffhG_AUcI.lgXJXgaVPjr[eO94A1.C50bBOZxcYfV])
		return true, tkGFesoo, y5d2lf
	}

	return false, nil, nil
}

func (ffhG_AUcI *FjZVwJmDPT) rtp1gNaU_fz(hwXqUFsD types.View) bool {
	return /*line Fj1NTMtjUzDN.go:1*/ ffhG_AUcI.fWUDWuAEZYiJ(hwXqUFsD) > ffhG_AUcI.jMiIWe*2/3
}

func (ffhG_AUcI *FjZVwJmDPT) Majority(hwXqUFsD types.View) bool {
	return /*line K_IQPyC9WjP.go:1*/ ffhG_AUcI.fWUDWuAEZYiJ(hwXqUFsD) > ffhG_AUcI.jMiIWe*1/2
}

func (ffhG_AUcI *FjZVwJmDPT) fWUDWuAEZYiJ(hwXqUFsD types.View) int {
	return /*line _gI8hxJT.go:1*/ len(ffhG_AUcI.lgXJXgaVPjr[hwXqUFsD])
}

var _ sync.Cond
var _ blockchain.Accept

const _ = types.ABORT
