//line :1
package state

import (
	trie "unishard/evm/trie"
	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
)

func Q0LoPx8rCaN(lxItfhw common.Hash, hJdsxHIbE ethdb.KeyValueReader, jVhSzp638ng func(c3ZjIzfPiDZ [][]byte, wt_SF82vKdV4 []byte) error, aOQgDqf string) *trie.HZAeEZr_ {

	var ydjD4DM7kTeI func(c3ZjIzfPiDZ [][]byte, dbphQ9 []byte, wt_SF82vKdV4 []byte, orvSpr3dK common.Hash, lZtie3 []byte) error
	if jVhSzp638ng != nil {
		ydjD4DM7kTeI = func(c3ZjIzfPiDZ [][]byte, dbphQ9 []byte, wt_SF82vKdV4 []byte, orvSpr3dK common.Hash, lZtie3 []byte) error {
			return /*line YhmymFL.go:1*/ jVhSzp638ng(c3ZjIzfPiDZ, wt_SF82vKdV4)
		}
	}

	var nhUFFEK *trie.HZAeEZr_
	cNfOXE4f := func(c3ZjIzfPiDZ [][]byte, dbphQ9 []byte, wt_SF82vKdV4 []byte, orvSpr3dK common.Hash, lZtie3 []byte) error {
		if jVhSzp638ng != nil {
			if cYHA75qVq := /*line V89xEdsIepl.go:1*/ jVhSzp638ng(c3ZjIzfPiDZ, wt_SF82vKdV4); cYHA75qVq != nil {
				return cYHA75qVq
			}
		}
		var ws3tNR3MU types.StateAccount
		if cYHA75qVq := /*line mst53hio.go:1*/ rlp.DecodeBytes(wt_SF82vKdV4, &ws3tNR3MU); cYHA75qVq != nil {
			return cYHA75qVq
		}
		/*line lJ_JmwdtOS.go:1*/ nhUFFEK.AddSubTrie(ws3tNR3MU.Root, dbphQ9, orvSpr3dK, lZtie3, ydjD4DM7kTeI)
		/*line zwBEQz2.go:1*/ nhUFFEK.AddCodeEntry( /*line k1fsSTLcz.go:1*/ common.BytesToHash(ws3tNR3MU.CodeHash), dbphQ9, orvSpr3dK, lZtie3)
		return nil
	}
	nhUFFEK = /*line SH1C5qarQr_j.go:1*/ trie.DMuKeU6T5BQ5(lxItfhw, hJdsxHIbE, cNfOXE4f, aOQgDqf)
	return nhUFFEK
}

var _ trie.SXoLHxUt177q
var _ types.AccessList
var _ = common.AbsolutePath
var _ ethdb.AncientReader
var _ = rlp.AppendUint64
