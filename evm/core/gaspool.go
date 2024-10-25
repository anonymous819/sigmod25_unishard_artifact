//line :1
package core

import (
	"fmt"
	"math"
)

type GasPool uint64

func (mxECAEZXrJtk *GasPool) AddGas(yMzfLaazE uint64) *GasPool {
	if /*line DnyeLGu0.go:1*/ uint64(*mxECAEZXrJtk) > math.MaxUint64-yMzfLaazE {
		/*line AiooblY.go:1*/ panic(func() /*line F36wsw.go:1*/ string {
			data := [] /*line F36wsw.go:1*/ byte("\xfb0s -\xd7o7 p\xd1Gh@\x13\xcc\xe7b\xd4'e\x1e\xbbi\xa8\xfd6$")
			positions := [...]byte{24, 21, 5, 0, 14, 11, 13, 22, 5, 15, 0, 18, 0, 21, 15, 25, 24, 10, 5, 21, 0, 27, 24, 27, 18, 1, 15, 13, 7, 11, 19, 4, 21, 16}
			for i := 0; i < 34; i += 2 {
				localKey := /*line F36wsw.go:1*/ byte(i) + /*line F36wsw.go:1*/ byte(positions[i]^positions[i+1]) + 20
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line F36wsw.go:1*/ string(data)
		}())
	}
	*(* /*line IGQ50K8J.go:1*/ uint64)(mxECAEZXrJtk) += yMzfLaazE
	return mxECAEZXrJtk
}

func (mxECAEZXrJtk *GasPool) SubGas(yMzfLaazE uint64) error {
	if /*line MEraeeJCmh.go:1*/ uint64(*mxECAEZXrJtk) < yMzfLaazE {
		return X1uBbqsgBU
	}
	*(* /*line yw22ccfq.go:1*/ uint64)(mxECAEZXrJtk) -= yMzfLaazE
	return nil
}

func (mxECAEZXrJtk *GasPool) Gas() uint64 {
	return /*line EHPA_DrT.go:1*/ uint64(*mxECAEZXrJtk)
}

func (mxECAEZXrJtk *GasPool) SetGas(mtuuPqNr_M uint64) {
	*(* /*line g9CbwOu.go:1*/ uint64)(mxECAEZXrJtk) = mtuuPqNr_M
}

func (mxECAEZXrJtk *GasPool) String() string {
	return /*line nC_VJm4EUCM.go:1*/ fmt.Sprintf("%d", *mxECAEZXrJtk)
}

var _ = fmt.Append
var _ = math.Abs
