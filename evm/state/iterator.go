//line :1
package state

import (
	"bytes"
	"errors"
	"fmt"

	types "unishard/evm/types"

	trie "unishard/evm/trie"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type nodeIterator struct {
	state *StateDB

	stateIt trie.FBy5J0gbkg
	dataIt  trie.FBy5J0gbkg

	accountHash common.Hash
	codeHash    common.Hash
	code        []byte

	Hash   common.Hash
	Parent common.Hash

	Error error
}

func cgDmmeS5Mw(ayL8rlo *StateDB) *nodeIterator {
	return &nodeIterator{
		state: ayL8rlo,
	}
}

func (s8hOm6P *nodeIterator) Next() bool {

	if s8hOm6P.Error != nil {
		return false
	}

	if cYHA75qVq := /*line hux3kHDcuamm.go:1*/ s8hOm6P.aTQMSxZ_nXx4(); cYHA75qVq != nil {
		s8hOm6P.Error = cYHA75qVq
		return false
	}
	return /*line gE8rIu.go:1*/ s8hOm6P.oxvOEE()
}

func (s8hOm6P *nodeIterator) aTQMSxZ_nXx4() error {

	if s8hOm6P.state == nil {
		return nil
	}

	var cYHA75qVq error
	if s8hOm6P.stateIt == nil {
		s8hOm6P.stateIt, cYHA75qVq = /*line waqVaIze0c.go:1*/ s8hOm6P.state.trie.NodeIterator(nil)
		if cYHA75qVq != nil {
			return cYHA75qVq
		}
	}

	if s8hOm6P.dataIt != nil {
		if pXikPJkp3aW := /*line l7Nlahg.go:1*/ s8hOm6P.dataIt.Next(true); !pXikPJkp3aW {
			if /*line yJkKSXD.go:1*/ s8hOm6P.dataIt.Error() != nil {
				return /*line q0Rg1dI2CNG.go:1*/ s8hOm6P.dataIt.Error()
			}
			s8hOm6P.dataIt = nil
		}
		return nil
	}

	if s8hOm6P.code != nil {
		s8hOm6P.code = nil
		return nil
	}

	if pXikPJkp3aW := /*line NtqX5pTS6fD.go:1*/ s8hOm6P.stateIt.Next(true); !pXikPJkp3aW {
		if /*line HMsWwz.go:1*/ s8hOm6P.stateIt.Error() != nil {
			return /*line Xskdr6FW1GhX.go:1*/ s8hOm6P.stateIt.Error()
		}
		s8hOm6P.state, s8hOm6P.stateIt = nil, nil
		return nil
	}

	if ! /*line i66YBb.go:1*/ s8hOm6P.stateIt.Leaf() {
		return nil
	}

	var hJAE3TNQR_k2 types.StateAccount
	if cYHA75qVq := /*line fW_SRmviuwVF.go:1*/ rlp.DecodeBytes( /*line DYPx1p.go:1*/ s8hOm6P.stateIt.LeafBlob(), &hJAE3TNQR_k2); cYHA75qVq != nil {
		return cYHA75qVq
	}

	ikSYP8 := /*line vg9MyRY.go:1*/ s8hOm6P.state.trie.GetKey( /*line WCQZzT3c.go:1*/ s8hOm6P.stateIt.LeafKey())
	if ikSYP8 == nil {
		return /*line H_wVhuW.go:1*/ errors.New(func() /*line CnBdp2u2T.go:1*/ string {
			key := [] /*line CnBdp2u2T.go:1*/ byte("*\xc4Rx\xcdݗ\x8a\xb6U>\xa33[\xc0\xcds9\xec\x15*A\xb28CҨ!p\vb\xa5")
			data := [] /*line CnBdp2u2T.go:1*/ byte("7\x9f\x11\xf7\xa8\x91ݖ\xab\x0f&\xcf2\x18\xb3S\xf6:4YE3n)3\x8f\xc1K\xf1W\n\xc0")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line CnBdp2u2T.go:1*/ string(data)
		}())
	}
	so0TaNzLDdpc := /*line sMh_j0c9RAfr.go:1*/ common.BytesToAddress(ikSYP8)

	wf4EYFKE3, cYHA75qVq := /*line UTpJUAg.go:1*/ s8hOm6P.state.db.OpenStorageTrie(s8hOm6P.state.originalRoot, so0TaNzLDdpc, hJAE3TNQR_k2.Root, s8hOm6P.state.trie)
	if cYHA75qVq != nil {
		return cYHA75qVq
	}
	s8hOm6P.dataIt, cYHA75qVq = /*line E26LS8N.go:1*/ wf4EYFKE3.NodeIterator(nil)
	if cYHA75qVq != nil {
		return cYHA75qVq
	}
	if ! /*line L4jh5T.go:1*/ s8hOm6P.dataIt.Next(true) {
		s8hOm6P.dataIt = nil
	}
	if ! /*line FNGm_AZe6u.go:1*/ bytes.Equal(hJAE3TNQR_k2.CodeHash /*line Aahjaf7.go:1*/, types.JhrQnETMFm.Bytes()) {
		s8hOm6P.codeHash = /*line IgxBRRUpCOe.go:1*/ common.BytesToHash(hJAE3TNQR_k2.CodeHash)
		s8hOm6P.code, cYHA75qVq = /*line wLDBZLgE2P.go:1*/ s8hOm6P.state.db.ContractCode(so0TaNzLDdpc /*line fQyUG4moUJ.go:1*/, common.BytesToHash(hJAE3TNQR_k2.CodeHash))
		if cYHA75qVq != nil {
			return /*line BphhBboHqTi.go:1*/ fmt.Errorf(func() /*line o1PkIOSeEmot.go:1*/ string {
				data := [] /*line o1PkIOSeEmot.go:1*/ byte("coRQ()t6\x14\x13v")
				positions := [...]byte{5, 4, 7, 6, 3, 8, 7, 6, 4, 4, 2, 9, 5, 5}
				for i := 0; i < 14; i += 2 {
					localKey := /*line o1PkIOSeEmot.go:1*/ byte(i) + /*line o1PkIOSeEmot.go:1*/ byte(positions[i]^positions[i+1]) + 98
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line o1PkIOSeEmot.go:1*/ string(data)
			}(), hJAE3TNQR_k2.CodeHash, cYHA75qVq)
		}
	}
	s8hOm6P.accountHash = /*line G2atDGVF.go:1*/ s8hOm6P.stateIt.Parent()
	return nil
}

func (s8hOm6P *nodeIterator) oxvOEE() bool {

	s8hOm6P.Hash = common.Hash{}

	if s8hOm6P.state == nil {
		return false
	}

	switch {
	case s8hOm6P.dataIt != nil:
		s8hOm6P.Hash, s8hOm6P.Parent = /*line GtGKUAB.go:1*/ s8hOm6P.dataIt.Hash() /*line NmM3Poop0qq.go:1*/, s8hOm6P.dataIt.Parent()
		if s8hOm6P.Parent == (common.Hash{}) {
			s8hOm6P.Parent = s8hOm6P.accountHash
		}
	case s8hOm6P.code != nil:
		s8hOm6P.Hash, s8hOm6P.Parent = s8hOm6P.codeHash, s8hOm6P.accountHash
	case s8hOm6P.stateIt != nil:
		s8hOm6P.Hash, s8hOm6P.Parent = /*line I8NWZ2pYC.go:1*/ s8hOm6P.stateIt.Hash() /*line FMZns2QcoY.go:1*/, s8hOm6P.stateIt.Parent()
	}
	return true
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append
var _ types.AccessList
var _ trie.SXoLHxUt177q
var _ = common.AbsolutePath
var _ = rlp.AppendUint64
