//line :1
package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/holiman/uint256"
)

func nIR3b9(waclpXm8dPXx, eKJjZpH *uint256.Int) (uint64, bool) {
	if ! /*line VRYxtAIb0M.go:1*/ eKJjZpH.IsUint64() {
		return 0, true
	}
	return /*line Mlikzu096.go:1*/ xn5Tz_ue(waclpXm8dPXx /*line LcRmydvbrrvz.go:1*/, eKJjZpH.Uint64())
}

func xn5Tz_ue(waclpXm8dPXx *uint256.Int, hPOvCmO4 uint64) (uint64, bool) {

	if hPOvCmO4 == 0 {
		return 0, false
	}

	_sqWZsdE, r9uWm2pY := /*line aSfbqWL4M.go:1*/ waclpXm8dPXx.Uint64WithOverflow()
	if r9uWm2pY {
		return 0, true
	}
	rVFSadsv := _sqWZsdE + hPOvCmO4

	return rVFSadsv, rVFSadsv < _sqWZsdE
}

func m8O5lZzG(eiZFOFAwLXjh []byte, baL8MX8dX6 uint64, jLa0ywd uint64) []byte {
	kpLBYg := /*line hmzPTVNI.go:1*/ uint64( /*line tp1pa2t.go:1*/ len(eiZFOFAwLXjh))
	if baL8MX8dX6 > kpLBYg {
		baL8MX8dX6 = kpLBYg
	}
	kd9DfDjo2zph := baL8MX8dX6 + jLa0ywd
	if kd9DfDjo2zph > kpLBYg {
		kd9DfDjo2zph = kpLBYg
	}
	return /*line OFwiB66f6al.go:1*/ common.RightPadBytes(eiZFOFAwLXjh[baL8MX8dX6:kd9DfDjo2zph] /*line eS60x9y3R8.go:1*/, int(jLa0ywd))
}

func f6N2wznhW(jLa0ywd uint64) uint64 {
	if jLa0ywd > math.MaxUint64-31 {
		return math.MaxUint64/32 + 1
	}

	return (jLa0ywd + 31) / 32
}

func aFNbM1_(l1KtNnRwliPB []byte) bool {
	for _, pW_50s4bg := range l1KtNnRwliPB {
		if pW_50s4bg != 0 {
			return false
		}
	}
	return true
}

var _ = common.AbsolutePath
var _ = math.BigMax
var _ = uint256.ErrBadBufferLength
