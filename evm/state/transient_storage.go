//line :1
package state

import (
	"github.com/ethereum/go-ethereum/common"
)

type transientStorage map[common.Address]Storage

func s9XSqaz() transientStorage {
	return /*line ccPLZPn3NWN.go:1*/ make(transientStorage)
}

func (alxknQMA7Uen transientStorage) Set(hv1UkMRaKf common.Address, ixeygbAgOap, fnngafl common.Hash) {
	if _, dNazL3Oz4 := alxknQMA7Uen[hv1UkMRaKf]; !dNazL3Oz4 {
		alxknQMA7Uen[hv1UkMRaKf] = /*line fY9XSVSko.go:1*/ make(Storage)
	}
	alxknQMA7Uen[hv1UkMRaKf][ixeygbAgOap] = fnngafl
}

func (alxknQMA7Uen transientStorage) Get(hv1UkMRaKf common.Address, ixeygbAgOap common.Hash) common.Hash {
	cGzt0N, dNazL3Oz4 := alxknQMA7Uen[hv1UkMRaKf]
	if !dNazL3Oz4 {
		return common.Hash{}
	}
	return cGzt0N[ixeygbAgOap]
}

func (alxknQMA7Uen transientStorage) Copy() transientStorage {
	vg59BBBGWI := /*line Xj9egnD10.go:1*/ make(transientStorage)
	for ixeygbAgOap, fnngafl := range alxknQMA7Uen {
		vg59BBBGWI[ixeygbAgOap] = /*line T6kzibBaEPk.go:1*/ fnngafl.Copy()
	}
	return vg59BBBGWI
}

var _ = common.AbsolutePath
