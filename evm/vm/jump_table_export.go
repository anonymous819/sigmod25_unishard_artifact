//line :1
package vm

import (
	"errors"

	"github.com/ethereum/go-ethereum/params"
)

func BtWI62(avYeWr0Mmle params.Rules) (JumpTable, error) {
	switch {
	case avYeWr0Mmle.IsVerkle:
		return /*line GJATTEsyT.go:1*/ s8uFHpJ() /*line Yzwv_5.go:1*/, errors.New(func() /*line HawzsQi67l3.go:1*/ string {
			data := [] /*line HawzsQi67l3.go:1*/ byte("ve\xb7\xec\x13\xd3\xe8f\xa9r\x01 \x9fo\xab d\xae=\xd1n\xb4\\ y\x99\xa1")
			positions := [...]byte{18, 4, 19, 6, 5, 8, 21, 21, 2, 3, 4, 4, 14, 4, 17, 19, 10, 18, 26, 18, 26, 12, 6, 22, 10, 2, 8, 25}
			for i := 0; i < 28; i += 2 {
				localKey := /*line HawzsQi67l3.go:1*/ byte(i) + /*line HawzsQi67l3.go:1*/ byte(positions[i]^positions[i+1]) + 171
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line HawzsQi67l3.go:1*/ string(data)
		}())
	case avYeWr0Mmle.IsPrague:
		return /*line xDkmaoGF.go:1*/ s8uFHpJ() /*line sh6vH7Qtlrrx.go:1*/, errors.New(func() /*line AtHD33h7X.go:1*/ string {
			data := [] /*line AtHD33h7X.go:1*/ byte("p-\xe0\xce>eן\x95rޢ\xa1\xda\x1a ڻ\x91iV\xda\x04 \xdae]")
			positions := [...]byte{7, 12, 16, 14, 22, 16, 18, 16, 1, 3, 3, 26, 20, 10, 1, 11, 16, 17, 24, 12, 8, 6, 13, 1, 13, 4, 3, 24, 7, 21, 21, 2, 17, 10, 13, 13}
			for i := 0; i < 36; i += 2 {
				localKey := /*line AtHD33h7X.go:1*/ byte(i) + /*line AtHD33h7X.go:1*/ byte(positions[i]^positions[i+1]) + 70
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line AtHD33h7X.go:1*/ string(data)
		}())
	case avYeWr0Mmle.IsCancun:
		return /*line axq8PXrpYm.go:1*/ s8uFHpJ(), nil
	case avYeWr0Mmle.IsShanghai:
		return /*line ECwSVXsh.go:1*/ ddIsmUpbc(), nil
	case avYeWr0Mmle.IsMerge:
		return /*line DPQf6ejcAM.go:1*/ in9K_X62jm(), nil
	case avYeWr0Mmle.IsLondon:
		return /*line h65ab9SJ.go:1*/ fuEI3sHVt9hH(), nil
	case avYeWr0Mmle.IsBerlin:
		return /*line HsZVcypP.go:1*/ f8rBfy_rdC(), nil
	case avYeWr0Mmle.IsIstanbul:
		return /*line H8naov.go:1*/ qsNj3ExQbuEI(), nil
	case avYeWr0Mmle.IsConstantinople:
		return /*line BU91Jq.go:1*/ aaQtRKQWQEX(), nil
	case avYeWr0Mmle.IsByzantium:
		return /*line mB0V5m.go:1*/ akQ0LksfZ(), nil
	case avYeWr0Mmle.IsEIP158:
		return /*line aKLl7pSd6.go:1*/ r55_aGYuJs(), nil
	case avYeWr0Mmle.IsEIP150:
		return /*line DpIIa79NJGx.go:1*/ mVuYzPpq(), nil
	case avYeWr0Mmle.IsHomestead:
		return /*line GYr8pRpR.go:1*/ j1_EIof(), nil
	}
	return /*line C0M35erlK.go:1*/ bATEwUY(), nil
}

func (nCnJ0mc *operation) Stack() (int, int) {
	return nCnJ0mc.minStack, nCnJ0mc.maxStack
}

func (nCnJ0mc *operation) HasCost() bool {

	return nCnJ0mc.dynamicGas != nil || nCnJ0mc.constantGas != 0
}

var _ = errors.As
var _ = params.AllCliqueProtocolChanges
