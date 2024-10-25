//line :1
package vm

import (
	"sync"

	"github.com/holiman/uint256"
)

var vFgvaP = sync.Pool{
	New: func() interface{} {
		return &IZHSJVjaAmxb{sBD8Un65JT: /*line W7GVprybL95.go:1*/ make([]uint256.Int, 0, 16)}
	},
}

type IZHSJVjaAmxb struct {
	sBD8Un65JT []uint256.Int
}

func krqDUTXe() *IZHSJVjaAmxb {
	return /*line ybjWP67.go:1*/ vFgvaP.Get().(*IZHSJVjaAmxb)
}

func gGwn14iwA2Q_(ek3u6QMaV *IZHSJVjaAmxb) {
	ek3u6QMaV.sBD8Un65JT = ek3u6QMaV.sBD8Un65JT[:0]
	/*line rkde67te3b.go:1*/ vFgvaP.Put(ek3u6QMaV)
}

func (muN3frGp2cUY *IZHSJVjaAmxb) Data() []uint256.Int {
	return muN3frGp2cUY.sBD8Un65JT
}

func (muN3frGp2cUY *IZHSJVjaAmxb) p69kwi(yHAU58VWIx3 *uint256.Int) {

	muN3frGp2cUY.sBD8Un65JT = /*line JsuBHOW_Ar.go:1*/ append(muN3frGp2cUY.sBD8Un65JT, *yHAU58VWIx3)
}

func (muN3frGp2cUY *IZHSJVjaAmxb) bIh0YHTaQv2I() (ag8Tdqxi4 uint256.Int) {
	ag8Tdqxi4 = muN3frGp2cUY.sBD8Un65JT[ /*line AcucxtxYYx.go:1*/ len(muN3frGp2cUY.sBD8Un65JT)-1]
	muN3frGp2cUY.sBD8Un65JT = muN3frGp2cUY.sBD8Un65JT[: /*line AuJsCO.go:1*/ len(muN3frGp2cUY.sBD8Un65JT)-1]
	return
}

func (muN3frGp2cUY *IZHSJVjaAmxb) ckaB4DI() int {
	return /*line HFaa6aac.go:1*/ len(muN3frGp2cUY.sBD8Un65JT)
}

func (muN3frGp2cUY *IZHSJVjaAmxb) adoI6fnxF2vY(iyrnVpBe8CCl int) {
	muN3frGp2cUY.sBD8Un65JT[ /*line NAW5rluJwjaJ.go:1*/ muN3frGp2cUY.ckaB4DI()-iyrnVpBe8CCl], muN3frGp2cUY.sBD8Un65JT[ /*line aibW0ZJL.go:1*/ muN3frGp2cUY.ckaB4DI()-1] = muN3frGp2cUY.sBD8Un65JT[ /*line dJEUQC.go:1*/ muN3frGp2cUY.ckaB4DI()-1], muN3frGp2cUY.sBD8Un65JT[ /*line oIcjED57r.go:1*/ muN3frGp2cUY.ckaB4DI()-iyrnVpBe8CCl]
}

func (muN3frGp2cUY *IZHSJVjaAmxb) ngLkeTa6Y(iyrnVpBe8CCl int) {
	/*line bXatBgVSh.go:1*/ muN3frGp2cUY.p69kwi(&muN3frGp2cUY.sBD8Un65JT[ /*line Cao7a2Sbi7dt.go:1*/ muN3frGp2cUY.ckaB4DI()-iyrnVpBe8CCl])
}

func (muN3frGp2cUY *IZHSJVjaAmxb) bx3dXQEGk() *uint256.Int {
	return &muN3frGp2cUY.sBD8Un65JT[ /*line JY3MBT.go:1*/ muN3frGp2cUY.ckaB4DI()-1]
}

func (muN3frGp2cUY *IZHSJVjaAmxb) Back(iyrnVpBe8CCl int) *uint256.Int {
	return &muN3frGp2cUY.sBD8Un65JT[ /*line w49YVoyvFav.go:1*/ muN3frGp2cUY.ckaB4DI()-iyrnVpBe8CCl-1]
}

var _ sync.Cond
var _ = uint256.ErrBadBufferLength
