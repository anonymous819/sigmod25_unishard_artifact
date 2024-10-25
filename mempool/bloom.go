//line :1
package mempool

import (
	crypto "unishard/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/willf/bitset"
)

const DEFAULT_SIZE = 2 << 24

var bs8sOGLAY7x = []uint{7, 11, 13, 31, 37, 61}

type FhaZQKI struct {
	QbdaoaCnZu  uint
	Ig0O0y4MkoO uint
}

type EV4KClspBSq struct {
	HMWR9gy    *bitset.BitSet
	NI0sp9IGUa [6]FhaZQKI
}

func Rq6zjAUwOnR6() *EV4KClspBSq {
	jqCJgx := /*line yfHX8uOcXNVr.go:1*/ new(EV4KClspBSq)
	for iG7hNFS := 0; iG7hNFS < /*line E5JNxI.go:1*/ len(jqCJgx.NI0sp9IGUa); iG7hNFS++ {
		jqCJgx.NI0sp9IGUa[iG7hNFS] = FhaZQKI{DEFAULT_SIZE, bs8sOGLAY7x[iG7hNFS]}
	}
	jqCJgx.HMWR9gy = /*line IXkLEcTNglaq.go:1*/ bitset.New(DEFAULT_SIZE)
	return jqCJgx
}

func (jqCJgx EV4KClspBSq) Add(kSTvTSmia3 common.Hash) {
	for _, f6_gnD2_nA6y := range jqCJgx.NI0sp9IGUa {
		/*line AtOcERy.go:1*/ jqCJgx.HMWR9gy.Set( /*line CE26mGA.go:1*/ f6_gnD2_nA6y.niAPaP(kSTvTSmia3))
	}
}

func (jqCJgx EV4KClspBSq) Contains(kSTvTSmia3 common.Hash) bool {
	if kSTvTSmia3 == /*line RC2wNLdZSC.go:1*/ crypto.ZsUswPaGG4Z(nil) {
		return false
	}
	aiFV8nw := true
	for _, f6_gnD2_nA6y := range jqCJgx.NI0sp9IGUa {
		aiFV8nw = aiFV8nw && /*line oN0bKFmT.go:1*/ jqCJgx.HMWR9gy.Test( /*line rOIXWp.go:1*/ f6_gnD2_nA6y.niAPaP(kSTvTSmia3))
	}
	return aiFV8nw
}

func (fDqwNOG FhaZQKI) niAPaP(kSTvTSmia3 common.Hash) uint {
	var tYibDeO uint = 0
	for iG7hNFS := 0; iG7hNFS < /*line AEe7bPPK7MQm.go:1*/ len(kSTvTSmia3); iG7hNFS++ {
		tYibDeO = tYibDeO*fDqwNOG.Ig0O0y4MkoO + /*line JvRz1ms6w.go:1*/ uint(kSTvTSmia3[iG7hNFS])
	}
	return (fDqwNOG.QbdaoaCnZu - 1) & tYibDeO
}

var _ = common.AbsolutePath
var _ = bitset.Base64StdEncoding
var _ crypto.Pp__49cd
