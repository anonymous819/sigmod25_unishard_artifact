//line :1
package vm

import (
	"github.com/holiman/uint256"
)

const (
	GasQuickStep   uint64 = 2
	GasFastestStep uint64 = 3
	GasFastStep    uint64 = 5
	GasMidStep     uint64 = 8
	GasSlowStep    uint64 = 10
	GasExtStep     uint64 = 20
)

func iHblnaC3VG(yR814Up bool, tRhtsiHO, aY3v02xiEWPQ uint64, x9soFqW *uint256.Int) (uint64, error) {
	if yR814Up {
		tRhtsiHO = tRhtsiHO - aY3v02xiEWPQ
		a0SqxOtD := tRhtsiHO - tRhtsiHO/64

		if ! /*line xpGjeCo.go:1*/ x9soFqW.IsUint64() || a0SqxOtD < /*line GkE2ayZs0L.go:1*/ x9soFqW.Uint64() {
			return a0SqxOtD, nil
		}
	}
	if ! /*line os2IAr2n.go:1*/ x9soFqW.IsUint64() {
		return 0, QiOgwB7BT
	}

	return /*line wLPsRRDm6UDi.go:1*/ x9soFqW.Uint64(), nil
}

var _ = uint256.ErrBadBufferLength
