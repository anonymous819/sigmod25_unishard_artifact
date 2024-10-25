//line :1
package pacemaker

import (
	"sync"

	blockchain "unishard/blockchain"
	types "unishard/types"
)

type IWNgjZJHY struct {
	n2VXca       int
	qow6rdzJrc_n map[types.View]map[types.NodeID]*TMO
	z_2ET95aBu   sync.Mutex
}

func Cfc0vlx(iDvh6le2Hy7Q int) *IWNgjZJHY {
	ffhG_AUcI := /*line Zc3a_PrMWjk.go:1*/ new(IWNgjZJHY)
	ffhG_AUcI.n2VXca = iDvh6le2Hy7Q
	ffhG_AUcI.qow6rdzJrc_n = /*line d0jUDq3e.go:1*/ make(map[types.View]map[types.NodeID]*TMO)

	return ffhG_AUcI
}

func (ffhG_AUcI *IWNgjZJHY) AddTmo(eO94A1 *TMO) (bool, *TC, *blockchain.WorkerBlock) {
	/*line EeqhqqBdGS.go:1*/ ffhG_AUcI.z_2ET95aBu.Lock()
	defer /*line FKg_2yTR.go:1*/ ffhG_AUcI.z_2ET95aBu.Unlock()
	if /*line fmq13Aa6.go:1*/ ffhG_AUcI.rtp1gNaU_fz(eO94A1.Ig3RQ9F) {
		return false, nil, nil
	}
	_, dy9ax2z0dP := ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F]
	if !dy9ax2z0dP {

		ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F] = /*line nB4DY0sljjm7.go:1*/ make(map[types.NodeID]*TMO)
	}
	ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F][eO94A1.Ln9AEaadIIY] = eO94A1
	if /*line h48TL3frbDW.go:1*/ ffhG_AUcI.rtp1gNaU_fz(eO94A1.Ig3RQ9F) {
		tkGFesoo, y5d2lf := /*line JMQX5bIoE7.go:1*/ CDzFORl9BQA(eO94A1.Ig3RQ9F, ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F])
		return true, tkGFesoo, y5d2lf
	}

	return false, nil, nil
}

func (ffhG_AUcI *IWNgjZJHY) AddMjorityTmo(eO94A1 *TMO) (bool, *TC, *blockchain.WorkerBlock) {
	/*line WMVEWRr.go:1*/ ffhG_AUcI.z_2ET95aBu.Lock()
	defer /*line fF_d1kCw.go:1*/ ffhG_AUcI.z_2ET95aBu.Unlock()
	if /*line ZKaKLa.go:1*/ ffhG_AUcI.Majority(eO94A1.Ig3RQ9F) {
		return false, nil, nil
	}
	_, dy9ax2z0dP := ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F]
	if !dy9ax2z0dP {

		ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F] = /*line s4mrfd.go:1*/ make(map[types.NodeID]*TMO)
	}
	ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F][eO94A1.Ln9AEaadIIY] = eO94A1
	if /*line a6PnbUIE.go:1*/ ffhG_AUcI.Majority(eO94A1.Ig3RQ9F) {
		tkGFesoo, y5d2lf := /*line lJS4J4MrPwWG.go:1*/ CDzFORl9BQA(eO94A1.Ig3RQ9F, ffhG_AUcI.qow6rdzJrc_n[eO94A1.Ig3RQ9F])
		return true, tkGFesoo, y5d2lf
	}

	return false, nil, nil
}

func (ffhG_AUcI *IWNgjZJHY) rtp1gNaU_fz(hwXqUFsD types.View) bool {
	return /*line vF110KE629Bq.go:1*/ ffhG_AUcI.fWUDWuAEZYiJ(hwXqUFsD) > ffhG_AUcI.n2VXca*2/3
}

func (ffhG_AUcI *IWNgjZJHY) Majority(hwXqUFsD types.View) bool {
	return /*line oYLzGuKkW.go:1*/ ffhG_AUcI.fWUDWuAEZYiJ(hwXqUFsD) > ffhG_AUcI.n2VXca*1/2
}

func (ffhG_AUcI *IWNgjZJHY) fWUDWuAEZYiJ(hwXqUFsD types.View) int {
	return /*line JJT8GDj.go:1*/ len(ffhG_AUcI.qow6rdzJrc_n[hwXqUFsD])
}

var _ sync.Cond
var _ blockchain.Accept

const _ = types.ABORT
