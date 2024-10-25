//line :1
package state

import (
	"errors"
	"fmt"

	types "unishard/evm/types"

	trie "unishard/evm/trie"

	"github.com/crate-crypto/go-ipa/banderwagon"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/trie/utils"
	"github.com/ethereum/go-ethereum/triedb"
)

const (
	codeSizeCacheSize = 100000

	codeCacheSize = 64 * 1024 * 1024

	commitmentSize = banderwagon.UncompressedSize

	commitmentCacheItems = 64 * 1024 * 1024 / (commitmentSize + common.AddressLength)
)

type Database interface {
	OpenTrie(lxItfhw common.Hash) (Trie, error)

	OpenStorageTrie(d1eqKPB3 common.Hash, so0TaNzLDdpc common.Address, lxItfhw common.Hash, fQodR9c7t Trie) (Trie, error)

	CopyTrie(Trie) Trie

	ContractCode(hv1UkMRaKf common.Address, hCvoh5IR common.Hash) ([]byte, error)

	ContractCodeSize(hv1UkMRaKf common.Address, hCvoh5IR common.Hash) (int, error)

	DiskDB() ethdb.KeyValueStore

	TrieDB() *triedb.Database
}

type Trie interface {
	GetKey([]byte) []byte

	GetAccount(so0TaNzLDdpc common.Address) (*types.StateAccount, error)

	GetStorage(hv1UkMRaKf common.Address, ixeygbAgOap []byte) ([]byte, error)

	UpdateAccount(so0TaNzLDdpc common.Address, hJAE3TNQR_k2 *types.StateAccount) error

	UpdateStorage(hv1UkMRaKf common.Address, ixeygbAgOap, fnngafl []byte) error

	DeleteAccount(so0TaNzLDdpc common.Address) error

	DeleteStorage(hv1UkMRaKf common.Address, ixeygbAgOap []byte) error

	UpdateContractCode(so0TaNzLDdpc common.Address, hCvoh5IR common.Hash, acaO5J6 []byte) error

	Hash() common.Hash

	Commit(mIdefX_D bool) (common.Hash, *trienode.NodeSet, error)

	NodeIterator(qH4BWRy []byte) (trie.FBy5J0gbkg, error)

	Prove(ixeygbAgOap []byte, tqi8IH9 ethdb.KeyValueWriter) error
}

func LPPxdRUd(tC4SGV ethdb.Database) Database {
	return /*line U1XdxDKaon.go:1*/ TKhm11(tC4SGV, nil)
}

func TKhm11(tC4SGV ethdb.Database, xk0ZHJ41b *triedb.Config) Database {
	return &cjBz8E{
		smG3p5Z6y87J: tC4SGV,
		bHyLEEfII1Q:/*line TrJNMZy94I.go:1*/ lru.NewCache[common.Hash, int](codeSizeCacheSize),
		w7wN69lPjSe:/*line gne0vad.go:1*/ lru.NewSizeConstrainedCache[common.Hash, []byte](codeCacheSize),
		gmoAKzY8:/*line I8pIuvMY.go:1*/ triedb.NewDatabase(tC4SGV, xk0ZHJ41b),
	}
}

func MhFpXyJz(tC4SGV ethdb.Database, aY5CHa *triedb.Database) Database {
	return &cjBz8E{
		smG3p5Z6y87J: tC4SGV,
		bHyLEEfII1Q:/*line hiZa58z9h0X.go:1*/ lru.NewCache[common.Hash, int](codeSizeCacheSize),
		w7wN69lPjSe:/*line SezWzR2Z.go:1*/ lru.NewSizeConstrainedCache[common.Hash, []byte](codeCacheSize),
		gmoAKzY8: aY5CHa,
	}
}

type cjBz8E struct {
	smG3p5Z6y87J ethdb.KeyValueStore
	bHyLEEfII1Q  *lru.Cache[common.Hash, int]
	w7wN69lPjSe  *lru.SizeConstrainedCache[common.Hash, []byte]
	gmoAKzY8     *triedb.Database
}

func (tC4SGV *cjBz8E) OpenTrie(lxItfhw common.Hash) (Trie, error) {
	if /*line eD86Fi.go:1*/ tC4SGV.gmoAKzY8.IsVerkle() {
		return /*line naWS_2s.go:1*/ trie.DDmplr(lxItfhw, tC4SGV.gmoAKzY8 /*line BbsKTJzP.go:1*/, utils.NewPointCache(commitmentCacheItems))
	}
	r8CzlEu, cYHA75qVq := /*line Zu_GoTJ.go:1*/ trie.MMAH7dV2( /*line CAzLCjHR.go:1*/ trie.P5xt0mP5FivA(lxItfhw), tC4SGV.gmoAKzY8)
	if cYHA75qVq != nil {
		return nil, cYHA75qVq
	}
	return r8CzlEu, nil
}

func (tC4SGV *cjBz8E) OpenStorageTrie(d1eqKPB3 common.Hash, so0TaNzLDdpc common.Address, lxItfhw common.Hash, aC3ahpdz Trie) (Trie, error) {

	if /*line b8dCa7NCA2.go:1*/ tC4SGV.gmoAKzY8.IsVerkle() {
		return aC3ahpdz, nil
	}
	r8CzlEu, cYHA75qVq := /*line mJao7x.go:1*/ trie.MMAH7dV2( /*line H8wUYwRqc1.go:1*/ trie.NYkaq0T(d1eqKPB3 /*line cL7eY16afhe.go:1*/, crypto.Keccak256Hash( /*line C7ncX07Mz.go:1*/ so0TaNzLDdpc.Bytes()), lxItfhw), tC4SGV.gmoAKzY8)
	if cYHA75qVq != nil {
		return nil, cYHA75qVq
	}
	return r8CzlEu, nil
}

func (tC4SGV *cjBz8E) CopyTrie(alxknQMA7Uen Trie) Trie {
	switch alxknQMA7Uen := alxknQMA7Uen.(type) {
	case *trie.IetHqRYp3VR:
		return /*line uQ3fc8.go:1*/ alxknQMA7Uen.Copy()
	default:
		/*line F0ptk7k8_kr.go:1*/ panic( /*line nXl3uXL.go:1*/ fmt.Errorf(func() /*line c8LplpH.go:1*/ string {
			fullData := [] /*line c8LplpH.go:1*/ byte("\xb0+\x1b O\xaa\\\x9d\xc3j\b\x8d\x04O\xb7rQޘ+q\xfd\x16\xaf)}\xf7\xaa\xdd\x7f\xbdd\xc1\x0e\xd8\xc0\xc0\xdf\x14T")
			idxKey := [] /*line c8LplpH.go:1*/ byte("\x1b,v\x19\x9a\xfa\x03\x83\x0f\xceU\xae\xc6")
			data := /*line c8LplpH.go:1*/ make([]byte, 0, 21)
			data = /*line c8LplpH.go:1*/ append(data, fullData[20^ /*line c8LplpH.go:1*/ int(idxKey[8])]^fullData[42^ /*line c8LplpH.go:1*/ int(idxKey[8])], fullData[22^ /*line c8LplpH.go:1*/ int(idxKey[3])]-fullData[21^ /*line c8LplpH.go:1*/ int(idxKey[3])], fullData[155^ /*line c8LplpH.go:1*/ int(idxKey[4])]-fullData[185^ /*line c8LplpH.go:1*/ int(idxKey[4])], fullData[142^ /*line c8LplpH.go:1*/ int(idxKey[11])]^fullData[185^ /*line c8LplpH.go:1*/ int(idxKey[11])], fullData[47^ /*line c8LplpH.go:1*/ int(idxKey[1])]+fullData[40^ /*line c8LplpH.go:1*/ int(idxKey[1])], fullData[9^ /*line c8LplpH.go:1*/ int(idxKey[8])]+fullData[13^ /*line c8LplpH.go:1*/ int(idxKey[8])], fullData[161^ /*line c8LplpH.go:1*/ int(idxKey[7])]-fullData[138^ /*line c8LplpH.go:1*/ int(idxKey[7])], fullData[210^ /*line c8LplpH.go:1*/ int(idxKey[9])]^fullData[219^ /*line c8LplpH.go:1*/ int(idxKey[9])], fullData[28^ /*line c8LplpH.go:1*/ int(idxKey[3])]^fullData[8^ /*line c8LplpH.go:1*/ int(idxKey[3])], fullData[74^ /*line c8LplpH.go:1*/ int(idxKey[10])]+fullData[116^ /*line c8LplpH.go:1*/ int(idxKey[10])], fullData[176^ /*line c8LplpH.go:1*/ int(idxKey[11])]-fullData[137^ /*line c8LplpH.go:1*/ int(idxKey[11])], fullData[63^ /*line c8LplpH.go:1*/ int(idxKey[3])]+fullData[9^ /*line c8LplpH.go:1*/ int(idxKey[3])], fullData[1^ /*line c8LplpH.go:1*/ int(idxKey[3])]+fullData[3^ /*line c8LplpH.go:1*/ int(idxKey[3])], fullData[93^ /*line c8LplpH.go:1*/ int(idxKey[10])]^fullData[91^ /*line c8LplpH.go:1*/ int(idxKey[10])], fullData[38^ /*line c8LplpH.go:1*/ int(idxKey[1])]^fullData[56^ /*line c8LplpH.go:1*/ int(idxKey[1])], fullData[167^ /*line c8LplpH.go:1*/ int(idxKey[7])]^fullData[131^ /*line c8LplpH.go:1*/ int(idxKey[7])], fullData[184^ /*line c8LplpH.go:1*/ int(idxKey[11])]+fullData[163^ /*line c8LplpH.go:1*/ int(idxKey[11])], fullData[169^ /*line c8LplpH.go:1*/ int(idxKey[11])]-fullData[183^ /*line c8LplpH.go:1*/ int(idxKey[11])], fullData[145^ /*line c8LplpH.go:1*/ int(idxKey[7])]+fullData[136^ /*line c8LplpH.go:1*/ int(idxKey[7])], fullData[63^ /*line c8LplpH.go:1*/ int(idxKey[1])]^fullData[49^ /*line c8LplpH.go:1*/ int(idxKey[1])])
			return /*line c8LplpH.go:1*/ string(data)
		}(), alxknQMA7Uen))
	}
}

func (tC4SGV *cjBz8E) ContractCode(so0TaNzLDdpc common.Address, hCvoh5IR common.Hash) ([]byte, error) {
	acaO5J6, _ := /*line yipPXBr8H.go:1*/ tC4SGV.w7wN69lPjSe.Get(hCvoh5IR)
	if /*line XKATiUFh.go:1*/ len(acaO5J6) > 0 {
		return acaO5J6, nil
	}
	acaO5J6 = /*line uYrqPg.go:1*/ rawdb.ReadCode(tC4SGV.smG3p5Z6y87J, hCvoh5IR)
	if /*line aBj05hBdG.go:1*/ len(acaO5J6) > 0 {
		/*line _p2C7YTP9.go:1*/ tC4SGV.w7wN69lPjSe.Add(hCvoh5IR, acaO5J6)
		/*line yxD_zyErkyAv.go:1*/ tC4SGV.bHyLEEfII1Q.Add(hCvoh5IR /*line xvhiqkIysQ.go:1*/, len(acaO5J6))
		return acaO5J6, nil
	}
	return nil /*line Aavrb7PW0g8.go:1*/, errors.New(func() /*line FwLHL5.go:1*/ string {
		seed := /*line FwLHL5.go:1*/ byte(196)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line FwLHL5.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line FwLHL5.go:1*/ fnc(170)(1)(27)(170)(82)(233)(26)(231)(20)
		return /*line FwLHL5.go:1*/ string(data)
	}())
}

func (tC4SGV *cjBz8E) ContractCodeWithPrefix(so0TaNzLDdpc common.Address, hCvoh5IR common.Hash) ([]byte, error) {
	acaO5J6, _ := /*line Lk2o3wYyT.go:1*/ tC4SGV.w7wN69lPjSe.Get(hCvoh5IR)
	if /*line NgDat3xx7G0.go:1*/ len(acaO5J6) > 0 {
		return acaO5J6, nil
	}
	acaO5J6 = /*line W09i9w.go:1*/ rawdb.ReadCodeWithPrefix(tC4SGV.smG3p5Z6y87J, hCvoh5IR)
	if /*line HFzfhxVA.go:1*/ len(acaO5J6) > 0 {
		/*line LMIXVbz8746.go:1*/ tC4SGV.w7wN69lPjSe.Add(hCvoh5IR, acaO5J6)
		/*line XL9qsjaansR.go:1*/ tC4SGV.bHyLEEfII1Q.Add(hCvoh5IR /*line cbHgBvZ.go:1*/, len(acaO5J6))
		return acaO5J6, nil
	}
	return nil /*line f6ggaW.go:1*/, errors.New(func() /*line CtaxmVrKY.go:1*/ string {
		data := [] /*line CtaxmVrKY.go:1*/ byte("\xbf\xb4j iou\xa2\xb3")
		positions := [...]byte{1, 2, 7, 2, 8, 1, 4, 0, 0, 7}
		for i := 0; i < 10; i += 2 {
			localKey := /*line CtaxmVrKY.go:1*/ byte(i) + /*line CtaxmVrKY.go:1*/ byte(positions[i]^positions[i+1]) + 207
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line CtaxmVrKY.go:1*/ string(data)
	}())
}

func (tC4SGV *cjBz8E) ContractCodeSize(hv1UkMRaKf common.Address, hCvoh5IR common.Hash) (int, error) {
	if rurYlaMG, dNazL3Oz4 := /*line DDjefbr.go:1*/ tC4SGV.bHyLEEfII1Q.Get(hCvoh5IR); dNazL3Oz4 {
		return rurYlaMG, nil
	}
	acaO5J6, cYHA75qVq := /*line II1hYkaecDq.go:1*/ tC4SGV.ContractCode(hv1UkMRaKf, hCvoh5IR)
	return /*line UoefLlH1I.go:1*/ len(acaO5J6), cYHA75qVq
}

func (tC4SGV *cjBz8E) DiskDB() ethdb.KeyValueStore {
	return tC4SGV.smG3p5Z6y87J
}

func (tC4SGV *cjBz8E) TrieDB() *triedb.Database {
	return tC4SGV.gmoAKzY8
}

var _ = errors.As
var _ = fmt.Append
var _ types.AccessList
var _ trie.SXoLHxUt177q
var _ = banderwagon.BatchMapToScalarField
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ = crypto.CompressPubkey
var _ ethdb.AncientReader
var _ trienode.MergedNodeSet
var _ = utils.BalanceKey
var _ triedb.Config
