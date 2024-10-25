//line :1

package state

import (
	"fmt"
	"sort"
	"time"

	snapshot "unishard/evm/state/snapshot"
	trie "unishard/evm/trie"
	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/trie/triestate"
	"github.com/holiman/uint256"
)

const (
	storageDeleteLimit = 512 * 1024 * 1024
)

type revision struct {
	id           int
	journalIndex int
}

type StateDB struct {
	db         Database
	prefetcher *triePrefetcher
	trie       Trie
	hasher     crypto.KeccakState
	snaps      *snapshot.CE2m1DUB6wW
	snap       snapshot.ExbZ15xa

	originalRoot common.Hash

	accounts       map[common.Hash][]byte
	storages       map[common.Hash]map[common.Hash][]byte
	accountsOrigin map[common.Address][]byte
	storagesOrigin map[common.Address]map[common.Hash][]byte

	stateObjects         map[common.Address]*stateObject
	stateObjectsPending  map[common.Address]struct{}
	stateObjectsDirty    map[common.Address]struct{}
	stateObjectsDestruct map[common.Address]*types.StateAccount

	dbErr error

	refund uint64

	thash   common.Hash
	txIndex int
	logs    map[common.Hash][]*types.Log
	logSize uint

	preimages map[common.Hash][]byte

	accessList *accessList

	transientStorage transientStorage

	journal        *journal
	validRevisions []revision
	nextRevisionId int

	AccountReads         time.Duration
	AccountHashes        time.Duration
	AccountUpdates       time.Duration
	AccountCommits       time.Duration
	StorageReads         time.Duration
	StorageHashes        time.Duration
	StorageUpdates       time.Duration
	StorageCommits       time.Duration
	SnapshotAccountReads time.Duration
	SnapshotStorageReads time.Duration
	SnapshotCommits      time.Duration
	TrieDBCommits        time.Duration

	AccountUpdated int
	StorageUpdated int
	AccountDeleted int
	StorageDeleted int

	onCommit func(d3JBDDpmp *triestate.Set)
}

func Fwoiwa1(lxItfhw common.Hash, tC4SGV Database, d5ez3p *snapshot.CE2m1DUB6wW) (*StateDB, error) {
	r8CzlEu, cYHA75qVq := /*line Jw9Cq51yqZ.go:1*/ tC4SGV.OpenTrie(lxItfhw)
	if cYHA75qVq != nil {
		return nil, cYHA75qVq
	}
	b8nxFWz4g8 := &StateDB{
		db:           tC4SGV,
		trie:         r8CzlEu,
		originalRoot: lxItfhw,
		snaps:        d5ez3p,
		accounts:/*line DpWz8631K.go:1*/ make(map[common.Hash][]byte),
		storages:/*line EVZ3l5861ZUQ.go:1*/ make(map[common.Hash]map[common.Hash][]byte),
		accountsOrigin:/*line C2j0mjp.go:1*/ make(map[common.Address][]byte),
		storagesOrigin:/*line lctfzwgdZyv.go:1*/ make(map[common.Address]map[common.Hash][]byte),
		stateObjects:/*line PYc9Mk.go:1*/ make(map[common.Address]*stateObject),
		stateObjectsPending:/*line lD6DVPOa.go:1*/ make(map[common.Address]struct{}),
		stateObjectsDirty:/*line ReDcT2gb.go:1*/ make(map[common.Address]struct{}),
		stateObjectsDestruct:/*line m8jJIDLEb.go:1*/ make(map[common.Address]*types.StateAccount),
		logs:/*line f4lVBz.go:1*/ make(map[common.Hash][]*types.Log),
		preimages:/*line KuPhgjdYSB7R.go:1*/ make(map[common.Hash][]byte),
		journal:/*line NBj54aPnWa3.go:1*/ vuoy40vfttW(),
		accessList:/*line mnlZlUCHk2eb.go:1*/ viqfAs6EZc(),
		transientStorage:/*line l80uLXSD.go:1*/ s9XSqaz(),
		hasher:/*line NL9tgz.go:1*/ crypto.NewKeccakState(),
	}
	if b8nxFWz4g8.snaps != nil {
		b8nxFWz4g8.snap = /*line rOscKSx.go:1*/ b8nxFWz4g8.snaps.Snapshot(lxItfhw)
	}
	return b8nxFWz4g8, nil
}

func (peYZBpNwox *StateDB) GetTrie() Trie {
	return peYZBpNwox.trie
}

func (peYZBpNwox *StateDB) StartPrefetcher(_6LsfLXKG string) {
	if peYZBpNwox.prefetcher != nil {
		/*line iaagwX5tb.go:1*/ peYZBpNwox.prefetcher.gfibELOKGAub()
		peYZBpNwox.prefetcher = nil
	}
	if peYZBpNwox.snap != nil {
		peYZBpNwox.prefetcher = /*line _Itssp.go:1*/ vnGIvB(peYZBpNwox.db, peYZBpNwox.originalRoot, _6LsfLXKG)
	}
}

func (peYZBpNwox *StateDB) StopPrefetcher() {
	if peYZBpNwox.prefetcher != nil {
		/*line oQLhda.go:1*/ peYZBpNwox.prefetcher.gfibELOKGAub()
		peYZBpNwox.prefetcher = nil
	}
}

func (peYZBpNwox *StateDB) eMZRw9Dc(cYHA75qVq error) {
	if peYZBpNwox.dbErr == nil {
		peYZBpNwox.dbErr = cYHA75qVq
	}
}

func (peYZBpNwox *StateDB) Error() error {
	return peYZBpNwox.dbErr
}

func (peYZBpNwox *StateDB) AddLog(n4_a0RnFhM *types.Log) {
	/*line b175G6.go:1*/ peYZBpNwox.journal.uwXnrqq(uRRh4T8d{inLGNw: peYZBpNwox.thash})

	n4_a0RnFhM.TxHash = peYZBpNwox.thash
	n4_a0RnFhM.TxIndex = /*line ufCZgXH.go:1*/ uint(peYZBpNwox.txIndex)
	n4_a0RnFhM.Index = peYZBpNwox.logSize
	peYZBpNwox.logs[peYZBpNwox.thash] = /*line I9XnM4.go:1*/ append(peYZBpNwox.logs[peYZBpNwox.thash], n4_a0RnFhM)
	peYZBpNwox.logSize++
}

func (peYZBpNwox *StateDB) GetLogs(cSqPoIL2n5 common.Hash, zSJdvuTU3FR uint64, fWo1fW common.Hash) []*types.Log {
	jrZPAbGCstr := peYZBpNwox.logs[cSqPoIL2n5]
	for _, fuNTsaDfte := range jrZPAbGCstr {
		fuNTsaDfte.BlockNumber = zSJdvuTU3FR
		fuNTsaDfte.BlockHash = fWo1fW
	}
	return jrZPAbGCstr
}

func (peYZBpNwox *StateDB) Logs() []*types.Log {
	var jrZPAbGCstr []*types.Log
	for _, dzuknsn := range peYZBpNwox.logs {
		jrZPAbGCstr = /*line wv2g7TGsQXo.go:1*/ append(jrZPAbGCstr, dzuknsn...)
	}
	return jrZPAbGCstr
}

func (peYZBpNwox *StateDB) AddPreimage(cSqPoIL2n5 common.Hash, ikSYP8 []byte) {
	if _, dNazL3Oz4 := peYZBpNwox.preimages[cSqPoIL2n5]; !dNazL3Oz4 {
		/*line KlGb4hvfJLEM.go:1*/ peYZBpNwox.journal.uwXnrqq(qhj_H1rUkvcr{gefaQDBFugQa: cSqPoIL2n5})
		iw6356Qvw0MZ := /*line CZBVCj69aLzm.go:1*/ make([]byte /*line U8bDnSWYZv.go:1*/, len(ikSYP8))
		/*line c2CpQ98ya.go:1*/ copy(iw6356Qvw0MZ, ikSYP8)
		peYZBpNwox.preimages[cSqPoIL2n5] = iw6356Qvw0MZ
	}
}

func (peYZBpNwox *StateDB) Preimages() map[common.Hash][]byte {
	return peYZBpNwox.preimages
}

func (peYZBpNwox *StateDB) AddRefund(mHQ9DMmu uint64) {
	/*line LfQ86Whhe.go:1*/ peYZBpNwox.journal.uwXnrqq(foFrwR8S4{eT55nqw_VBa: peYZBpNwox.refund})
	peYZBpNwox.refund += mHQ9DMmu
}

func (peYZBpNwox *StateDB) SubRefund(mHQ9DMmu uint64) {
	/*line Mso8ve3m.go:1*/ peYZBpNwox.journal.uwXnrqq(foFrwR8S4{eT55nqw_VBa: peYZBpNwox.refund})
	if mHQ9DMmu > peYZBpNwox.refund {
		/*line jrpbjHau62.go:1*/ panic( /*line xTKcv5IaDZqd.go:1*/ fmt.Sprintf(func() /*line J6QSG2OMH.go:1*/ string {
			fullData := [] /*line J6QSG2OMH.go:1*/ byte("\xd6\xec\xe4t\xa2\xee\xcf@A\xf9\x94c\x05p;\xb6\x15K\xd1}\x019^!u\xb3\x8bт\xbf\x91\x06a\x83Rk\xcf\xe35:]\xaa\xa7\xe2\x17i\xd17\xe7\x9d\\\xd3S\xab\xe3\x8aR\xfaC\xa6<\xb8\x8dq\xcb\xc3n\x87\xcd9\xca\x13\x8bJ\xef\x93\xffI2\x8d\x9f\xaf>y\x13\x0e\x17\x82β\x00\xc5Ư\x86\n")
			idxKey := [] /*line J6QSG2OMH.go:1*/ byte("\xee\x15")
			data := /*line J6QSG2OMH.go:1*/ make([]byte, 0, 49)
			data = /*line J6QSG2OMH.go:1*/ append(data, fullData[86^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[85^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[40^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[33^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[163^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[203^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[164^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[176^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[253^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]^fullData[186^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[248^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[215^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[58^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[67^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[8^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[39^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[18^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[14^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[221^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[234^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[34^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[23^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[168^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[199^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[76^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[12^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[42^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[89^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[92^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[21^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[31^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[91^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[216^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[185^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[19^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[30^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[242^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[169^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[195^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[187^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[15^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[54^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[80^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[29^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[5^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[24^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[251^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]^fullData[255^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[51^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]-fullData[73^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[84^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[61^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[231^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]^fullData[252^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[53^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[10^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[90^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[20^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[11^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[62^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[201^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[180^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[202^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[191^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[210^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[194^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[78^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[69^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[165^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[208^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[212^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[226^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[77^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[55^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[249^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[179^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[93^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[16^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[81^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[32^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[237^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[250^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[223^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[192^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[225^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[214^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[246^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]-fullData[224^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[196^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[189^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])], fullData[71^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]+fullData[37^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[87^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])]^fullData[74^ /*line J6QSG2OMH.go:1*/ int(idxKey[1])], fullData[207^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])]+fullData[213^ /*line J6QSG2OMH.go:1*/ int(idxKey[0])])
			return /*line J6QSG2OMH.go:1*/ string(data)
		}(), mHQ9DMmu, peYZBpNwox.refund))
	}
	peYZBpNwox.refund -= mHQ9DMmu
}

func (peYZBpNwox *StateDB) Exist(hv1UkMRaKf common.Address) bool {
	return /*line H29ayt877.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf) != nil
}

func (peYZBpNwox *StateDB) Empty(hv1UkMRaKf common.Address) bool {
	o8hlB0H := /*line kLXGg0_klHUa.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	return o8hlB0H == nil || /*line GMpNdNXY.go:1*/ o8hlB0H.usx3KByCvUvp()
}

func (peYZBpNwox *StateDB) GetBalance(hv1UkMRaKf common.Address) *uint256.Int {
	p82q2vrTl7HW := /*line I9LmJitIDiX.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line gKtRCuSQ_S0g.go:1*/ p82q2vrTl7HW.Balance()
	}
	return common.U2560
}

func (peYZBpNwox *StateDB) GetNonce(hv1UkMRaKf common.Address) uint64 {
	p82q2vrTl7HW := /*line sMYEUmQ2me.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line EAbaB0C_.go:1*/ p82q2vrTl7HW.Nonce()
	}

	return 0
}

func (peYZBpNwox *StateDB) GetStorageRoot(hv1UkMRaKf common.Address) common.Hash {
	p82q2vrTl7HW := /*line GCODzz3.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line Ki3Et40BVI5.go:1*/ p82q2vrTl7HW.Root()
	}
	return common.Hash{}
}

func (peYZBpNwox *StateDB) TxIndex() int {
	return peYZBpNwox.txIndex
}

func (peYZBpNwox *StateDB) GetCode(hv1UkMRaKf common.Address) []byte {
	p82q2vrTl7HW := /*line aRyoNNRP5.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line em7ptW_8y6Bz.go:1*/ p82q2vrTl7HW.Code()
	}
	return nil
}

func (peYZBpNwox *StateDB) GetDeploycode(hv1UkMRaKf common.Address) []byte {
	p82q2vrTl7HW := /*line lkwhn9.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line IMkL7N.go:1*/ p82q2vrTl7HW.DeployCode()
	}
	return nil
}

func (peYZBpNwox *StateDB) GetCodeSize(hv1UkMRaKf common.Address) int {
	p82q2vrTl7HW := /*line enmBYs.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line HGsEaWl3MRmx.go:1*/ p82q2vrTl7HW.CodeSize()
	}
	return 0
}

func (peYZBpNwox *StateDB) GetCodeHash(hv1UkMRaKf common.Address) common.Hash {
	p82q2vrTl7HW := /*line MUWNWe7e5.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line oAWdpCa.go:1*/ common.BytesToHash( /*line VwSDV8.go:1*/ p82q2vrTl7HW.CodeHash())
	}
	return common.Hash{}
}

func (peYZBpNwox *StateDB) GetState(hv1UkMRaKf common.Address, cSqPoIL2n5 common.Hash) common.Hash {
	p82q2vrTl7HW := /*line Ewc9TJo9.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line BKWh8_3MLi.go:1*/ p82q2vrTl7HW.GetState(cSqPoIL2n5)
	}
	return common.Hash{}
}

func (peYZBpNwox *StateDB) GetCommittedState(hv1UkMRaKf common.Address, cSqPoIL2n5 common.Hash) common.Hash {
	p82q2vrTl7HW := /*line kdGy69.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return /*line uaORo34mGYS1.go:1*/ p82q2vrTl7HW.GetCommittedState(cSqPoIL2n5)
	}
	return common.Hash{}
}

func (peYZBpNwox *StateDB) Database() Database {
	return peYZBpNwox.db
}

func (peYZBpNwox *StateDB) HasSelfDestructed(hv1UkMRaKf common.Address) bool {
	p82q2vrTl7HW := /*line bnPmoD.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		return p82q2vrTl7HW.selfDestructed
	}
	return false
}

func (peYZBpNwox *StateDB) AddBalance(hv1UkMRaKf common.Address, zGNz6Q9KhWC *uint256.Int) {
	p82q2vrTl7HW := /*line eouGlmo.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line sWaJikqo.go:1*/ p82q2vrTl7HW.AddBalance(zGNz6Q9KhWC)
	}
}

func (peYZBpNwox *StateDB) SubBalance(hv1UkMRaKf common.Address, zGNz6Q9KhWC *uint256.Int) {
	p82q2vrTl7HW := /*line J2aCAfqla2.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line PfIMLfv1c.go:1*/ p82q2vrTl7HW.SubBalance(zGNz6Q9KhWC)
	}
}

func (peYZBpNwox *StateDB) SetBalance(hv1UkMRaKf common.Address, zGNz6Q9KhWC *uint256.Int) {
	p82q2vrTl7HW := /*line fM8lW78.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line CDW1AhH043F.go:1*/ p82q2vrTl7HW.SetBalance(zGNz6Q9KhWC)
	}
}

func (peYZBpNwox *StateDB) SetNonce(hv1UkMRaKf common.Address, r1rvsu uint64) {
	p82q2vrTl7HW := /*line aqoPFA6XYdC.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line h_TWWGF.go:1*/ p82q2vrTl7HW.SetNonce(r1rvsu)
	}
}

func (peYZBpNwox *StateDB) SetCode(hv1UkMRaKf common.Address, acaO5J6 []byte) {
	p82q2vrTl7HW := /*line s8dQgvg.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line C7ai8gz.go:1*/ p82q2vrTl7HW.SetCode( /*line QXFGOfHvr.go:1*/ crypto.Keccak256Hash(acaO5J6), acaO5J6)
	}
}

func (peYZBpNwox *StateDB) SetDeployCode(hv1UkMRaKf common.Address, acaO5J6 []byte) {
	p82q2vrTl7HW := /*line a6MkEf8nm.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line qFv_9af.go:1*/ p82q2vrTl7HW.SetDeploycode(acaO5J6)
	}
}

func (peYZBpNwox *StateDB) SetState(hv1UkMRaKf common.Address, ixeygbAgOap, fnngafl common.Hash) {
	p82q2vrTl7HW := /*line Ui_9dUUZ7es.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	if p82q2vrTl7HW != nil {
		/*line j90HgqEr99W.go:1*/ p82q2vrTl7HW.SetState(ixeygbAgOap, fnngafl)
	}
}

func (peYZBpNwox *StateDB) SetStorage(hv1UkMRaKf common.Address, vg59BBBGWI map[common.Hash]common.Hash) {

	if _, dNazL3Oz4 := peYZBpNwox.stateObjectsDestruct[hv1UkMRaKf]; !dNazL3Oz4 {
		peYZBpNwox.stateObjectsDestruct[hv1UkMRaKf] = nil
	}
	p82q2vrTl7HW := /*line Hc8s0VRrp.go:1*/ peYZBpNwox.b47R9p3B(hv1UkMRaKf)
	for _3JfOPUaHxI, gz5sHd := range vg59BBBGWI {
		/*line GUiT88.go:1*/ p82q2vrTl7HW.SetState(_3JfOPUaHxI, gz5sHd)
	}
}

func (peYZBpNwox *StateDB) SelfDestruct(hv1UkMRaKf common.Address) {
	p82q2vrTl7HW := /*line EFgyX5B3.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW == nil {
		return
	}
	/*line JyKhZ384P.go:1*/ peYZBpNwox.journal.uwXnrqq(kzbdktS{
		sQRjYTvw:  &hv1UkMRaKf,
		yqkgNNBcP: p82q2vrTl7HW.selfDestructed,
		ba3Vrv:/*line DNhbHLs.go:1*/ new(uint256.Int).Set( /*line Im8rBxaDy.go:1*/ p82q2vrTl7HW.Balance()),
	})
	/*line H5f2Q_4UKs_7.go:1*/ p82q2vrTl7HW.eUznOx()
	p82q2vrTl7HW.data.Balance = /*line Ndketb2xDA.go:1*/ new(uint256.Int)
}

func (peYZBpNwox *StateDB) DeleteAccount(hv1UkMRaKf common.Address) {
	p82q2vrTl7HW := /*line lD5y_4XYf.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW == nil {
		return
	}

	/*line BM8oCM3zX.go:1*/
	delete(peYZBpNwox.accounts, p82q2vrTl7HW.addrHash)
	/*line FeFUIoTYneAa.go:1*/ delete(peYZBpNwox.storages, p82q2vrTl7HW.addrHash)
	/*line H1CnST2x.go:1*/ delete(peYZBpNwox.accountsOrigin, p82q2vrTl7HW.address)
	/*line DBkpafAi5e8.go:1*/ delete(peYZBpNwox.storagesOrigin, p82q2vrTl7HW.address)
	p82q2vrTl7HW.data.Balance = /*line i_cYVzWbuqe8.go:1*/ new(uint256.Int)
	p82q2vrTl7HW.data.Nonce = 0
	p82q2vrTl7HW.data.Root = types.NrGuaNA21
	p82q2vrTl7HW.data.CodeHash = /*line gs_hnglK.go:1*/ types.JhrQnETMFm.Bytes()
}

func (peYZBpNwox *StateDB) Selfdestruct6780(hv1UkMRaKf common.Address) {
	p82q2vrTl7HW := /*line UEch6rAZHOAc.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW == nil {
		return
	}

	if p82q2vrTl7HW.created {
		/*line wWsp67oFaW1O.go:1*/ peYZBpNwox.SelfDestruct(hv1UkMRaKf)
	}
}

func (peYZBpNwox *StateDB) SetTransientState(hv1UkMRaKf common.Address, ixeygbAgOap, fnngafl common.Hash) {
	reikfgaLNK := /*line iavMZg.go:1*/ peYZBpNwox.GetTransientState(hv1UkMRaKf, ixeygbAgOap)
	if reikfgaLNK == fnngafl {
		return
	}
	/*line VYMPTZy4cM7.go:1*/ peYZBpNwox.journal.uwXnrqq(yzAiyxK{
		iLbaDHlKPp:  &hv1UkMRaKf,
		r9VPURO:     ixeygbAgOap,
		eT8enfTtJwO: reikfgaLNK,
	})
	/*line tVAhWa.go:1*/ peYZBpNwox.vfrWO0H(hv1UkMRaKf, ixeygbAgOap, fnngafl)
}

func (peYZBpNwox *StateDB) vfrWO0H(hv1UkMRaKf common.Address, ixeygbAgOap, fnngafl common.Hash) {
	/*line VgZy54R.go:1*/ peYZBpNwox.transientStorage.Set(hv1UkMRaKf, ixeygbAgOap, fnngafl)
}

func (peYZBpNwox *StateDB) GetTransientState(hv1UkMRaKf common.Address, ixeygbAgOap common.Hash) common.Hash {
	return /*line gPTL676q.go:1*/ peYZBpNwox.transientStorage.Get(hv1UkMRaKf, ixeygbAgOap)
}

func (peYZBpNwox *StateDB) yXyxeb11X(ws3tNR3MU *stateObject) {

	if metrics.EnabledExpensive {
		defer func( /*line Fm33Dcv.go:1*/ stqEk4Brz1wF time.Time) {
			peYZBpNwox.AccountUpdates += /*line NL020sI.go:1*/ time.Since(stqEk4Brz1wF)
		}( /*line bJXcBmJuIM__.go:1*/ time.Now())
	}

	hv1UkMRaKf := /*line HNaahz.go:1*/ ws3tNR3MU.Address()
	if cYHA75qVq := /*line JNoOLi_4.go:1*/ peYZBpNwox.trie.UpdateAccount(hv1UkMRaKf, &ws3tNR3MU.data); cYHA75qVq != nil {
		/*line DlWBx1UUO.go:1*/ peYZBpNwox.eMZRw9Dc( /*line SnK8viQa.go:1*/ fmt.Errorf(func() /*line hyiGI7_hkNl.go:1*/ string {
			key := [] /*line hyiGI7_hkNl.go:1*/ byte("ܛT\n\x82\tހ\xb3&\to\x8d#%.\xb4u\xc9\xf4\x7fת\xcaM\xa39\xce\xcb=\t\x88")
			data := [] /*line hyiGI7_hkNl.go:1*/ byte("\x99\xd5\x10W\xf2\\u\xf4\xaeN\\\xe0\xd5G@5\xc0\xab_1\xf9Rv\x9b%\xcf6\xa4o\xe3\x1c\xee")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line hyiGI7_hkNl.go:1*/ string(data)
		}(), hv1UkMRaKf[:], cYHA75qVq))
	}
	if ws3tNR3MU.dirtyCode {
		/*line CR59rG.go:1*/ peYZBpNwox.trie.UpdateContractCode( /*line FUoU9w5.go:1*/ ws3tNR3MU.Address() /*line _d2WCs1Zsce.go:1*/, common.BytesToHash( /*line r5D5LITq.go:1*/ ws3tNR3MU.CodeHash()), ws3tNR3MU.code)
	}

	peYZBpNwox.accounts[ws3tNR3MU.addrHash] = /*line KEaqk_kdnEW.go:1*/ types.K2bJTLXfmaB(ws3tNR3MU.data)

	if _, dNazL3Oz4 := peYZBpNwox.accountsOrigin[ws3tNR3MU.address]; !dNazL3Oz4 {
		if ws3tNR3MU.origin == nil {
			peYZBpNwox.accountsOrigin[ws3tNR3MU.address] = nil
		} else {
			peYZBpNwox.accountsOrigin[ws3tNR3MU.address] = /*line PqoQ3My6I7.go:1*/ types.K2bJTLXfmaB(*ws3tNR3MU.origin)
		}
	}
}

func (peYZBpNwox *StateDB) weye1mu(ws3tNR3MU *stateObject) {

	if metrics.EnabledExpensive {
		defer func( /*line ZviQNzvRJ.go:1*/ stqEk4Brz1wF time.Time) {
			peYZBpNwox.AccountUpdates += /*line gOihoVnlYb.go:1*/ time.Since(stqEk4Brz1wF)
		}( /*line Mc0K6l.go:1*/ time.Now())
	}

	hv1UkMRaKf := /*line SZpGwAb.go:1*/ ws3tNR3MU.Address()
	if cYHA75qVq := /*line I7A1Bg.go:1*/ peYZBpNwox.trie.DeleteAccount(hv1UkMRaKf); cYHA75qVq != nil {
		/*line Dlmp3Sbb2.go:1*/ peYZBpNwox.eMZRw9Dc( /*line StvVYkm.go:1*/ fmt.Errorf(func() /*line bAVPTBxC.go:1*/ string {
			key := [] /*line bAVPTBxC.go:1*/ byte("\xf5!\xfd\xd6M\xb0N\xb4\xa9\xfc#HW\xe6\x8cO\xf0\xe0\xb7\xc6\xcc\vx\xd3{d\x89Q\xdfu\x95h")
			data := [] /*line bAVPTBxC.go:1*/ byte("oDo\x8f'\xb5\x05\xc0\xb8xB\a\v\x84\xd9\x14\x84@q_\xac\x1e\xa8\x92\xf7\x0e\xe6![\xab\x90\x0e")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line bAVPTBxC.go:1*/ string(data)
		}(), hv1UkMRaKf[:], cYHA75qVq))
	}
}

func (peYZBpNwox *StateDB) nxJzqNXi(hv1UkMRaKf common.Address) *stateObject {
	if ws3tNR3MU := /*line GX3JHYiLJGF.go:1*/ peYZBpNwox.aOEYXRzaHwFc(hv1UkMRaKf); ws3tNR3MU != nil && !ws3tNR3MU.deleted {
		return ws3tNR3MU
	}
	return nil
}

func (peYZBpNwox *StateDB) aOEYXRzaHwFc(hv1UkMRaKf common.Address) *stateObject {

	if ws3tNR3MU := peYZBpNwox.stateObjects[hv1UkMRaKf]; ws3tNR3MU != nil {
		return ws3tNR3MU
	}

	var tfksAzzqIpKa *types.StateAccount
	if peYZBpNwox.snap != nil {
		stqEk4Brz1wF := /*line E6_0k9f5O.go:1*/ time.Now()
		b0gZOIQWE, cYHA75qVq := /*line lCLrW1.go:1*/ peYZBpNwox.snap.Account( /*line _Vxy4aRk.go:1*/ crypto.HashData(peYZBpNwox.hasher /*line Y36WWJ07PkY.go:1*/, hv1UkMRaKf.Bytes()))
		if metrics.EnabledExpensive {
			peYZBpNwox.SnapshotAccountReads += /*line X4zVmVUrxJbJ.go:1*/ time.Since(stqEk4Brz1wF)
		}
		if cYHA75qVq == nil {
			if b0gZOIQWE == nil {
				return nil
			}
			tfksAzzqIpKa = &types.StateAccount{
				Nonce:    b0gZOIQWE.Nonce,
				Balance:  b0gZOIQWE.Balance,
				CodeHash: b0gZOIQWE.CodeHash,
				Root:/*line pp6EODWg_j.go:1*/ common.BytesToHash(b0gZOIQWE.Root),
			}
			if /*line HmdoSCRg17w.go:1*/ len(tfksAzzqIpKa.CodeHash) == 0 {
				tfksAzzqIpKa.CodeHash = /*line MCbJ_5m0Ns_E.go:1*/ types.JhrQnETMFm.Bytes()
			}
			if tfksAzzqIpKa.Root == (common.Hash{}) {
				tfksAzzqIpKa.Root = types.NrGuaNA21
			}
		}
	}

	if tfksAzzqIpKa == nil {
		stqEk4Brz1wF := /*line xA9zOyO8.go:1*/ time.Now()
		var cYHA75qVq error
		tfksAzzqIpKa, cYHA75qVq = /*line PAWBGS6VS9L.go:1*/ peYZBpNwox.trie.GetAccount(hv1UkMRaKf)
		if metrics.EnabledExpensive {
			peYZBpNwox.AccountReads += /*line vFh6pgebZ.go:1*/ time.Since(stqEk4Brz1wF)
		}
		if cYHA75qVq != nil {
			/*line O6FxVmBVFSl.go:1*/ peYZBpNwox.eMZRw9Dc( /*line _a7QDRk.go:1*/ fmt.Errorf(func() /*line vInPvugcPhpO.go:1*/ string {
				data := /*line vInPvugcPhpO.go:1*/ make([]byte, 0, 36)
				i := 1
				decryptKey := 249
				for counter := 0; i != 0; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 5:
						i = 0
						for y := range data {
							data[y] = data[y] - /*line vInPvugcPhpO.go:1*/ byte(decryptKey^y)
						}
					case 7:
						i = 3
						data = /*line vInPvugcPhpO.go:1*/ append(data, "\xf9M"...,
						)
					case 6:
						i = 7
						data = /*line vInPvugcPhpO.go:1*/ append(data, "\xf6\xff"...,
						)
					case 8:
						i = 6
						data = /*line vInPvugcPhpO.go:1*/ append(data, "3E"...,
						)
					case 3:
						data = /*line vInPvugcPhpO.go:1*/ append(data, "\x03\xfb"...,
						)
						i = 2
					case 1:
						i = 4
						data = /*line vInPvugcPhpO.go:1*/ append(data, ")(4"...,
						)
					case 14:
						data = /*line vInPvugcPhpO.go:1*/ append(data, ")9/"...,
						)
						i = 11
					case 15:
						i = 5
						data = /*line vInPvugcPhpO.go:1*/ append(data, "\x17\x02\bW"...,
						)
					case 11:
						i = 9
						data = /*line vInPvugcPhpO.go:1*/ append(data, "\x1e<"...,
						)
					case 2:
						data = /*line vInPvugcPhpO.go:1*/ append(data, "=KPN"...,
						)
						i = 13
					case 10:
						data = /*line vInPvugcPhpO.go:1*/ append(data, "<8"...,
						)
						i = 8
					case 9:
						data = /*line vInPvugcPhpO.go:1*/ append(data, 42)
						i = 12
					case 12:
						i = 10
						data = /*line vInPvugcPhpO.go:1*/ append(data, "B4\x1b/"...,
						)
					case 4:
						data = /*line vInPvugcPhpO.go:1*/ append(data, "\x05+3"...,
						)
						i = 14
					case 13:
						data = /*line vInPvugcPhpO.go:1*/ append(data, 78)
						i = 15
					}
				}
				return /*line vInPvugcPhpO.go:1*/ string(data)
			}(), /*line Qdy4Al.go:1*/ hv1UkMRaKf.Bytes(), cYHA75qVq))
			return nil
		}
		if tfksAzzqIpKa == nil {
			return nil
		}
	}

	ws3tNR3MU := /*line C3pvNOT24rUy.go:1*/ zkEv6T(peYZBpNwox, hv1UkMRaKf, tfksAzzqIpKa)
	/*line aQamt_.go:1*/ peYZBpNwox.wTQEdk(ws3tNR3MU)
	return ws3tNR3MU
}

func (peYZBpNwox *StateDB) wTQEdk(gds4Bcqt *stateObject) {
	peYZBpNwox.stateObjects[ /*line sTyLTk.go:1*/ gds4Bcqt.Address()] = gds4Bcqt
}

func (peYZBpNwox *StateDB) b47R9p3B(hv1UkMRaKf common.Address) *stateObject {
	p82q2vrTl7HW := /*line x1aueF.go:1*/ peYZBpNwox.nxJzqNXi(hv1UkMRaKf)
	if p82q2vrTl7HW == nil {
		p82q2vrTl7HW, _ = /*line uRcr47.go:1*/ peYZBpNwox.mXXLmsylBEwb(hv1UkMRaKf)
	}
	return p82q2vrTl7HW
}

func (peYZBpNwox *StateDB) mXXLmsylBEwb(hv1UkMRaKf common.Address) (ikfPxhlba, reikfgaLNK *stateObject) {
	reikfgaLNK = /*line KWUG8z1eNj.go:1*/ peYZBpNwox.aOEYXRzaHwFc(hv1UkMRaKf)
	ikfPxhlba = /*line C3VK4nMq.go:1*/ zkEv6T(peYZBpNwox, hv1UkMRaKf, nil)
	if reikfgaLNK == nil {
		/*line ERkX3Xwgi.go:1*/ peYZBpNwox.journal.uwXnrqq(kLOasI4ckgD{huIthfk25W: &hv1UkMRaKf})
	} else {

		_, ssgiIj7R5c := peYZBpNwox.stateObjectsDestruct[reikfgaLNK.address]
		if !ssgiIj7R5c {
			peYZBpNwox.stateObjectsDestruct[reikfgaLNK.address] = reikfgaLNK.origin
		}

		tSTepB, dNazL3Oz4 := peYZBpNwox.accountsOrigin[reikfgaLNK.address]
		/*line FaMM2fcA.go:1*/ peYZBpNwox.journal.uwXnrqq(vF3b_q{
			yqSjsAK:      &hv1UkMRaKf,
			lpZWUm0WTZ5T: reikfgaLNK,
			aeIv_xPXxXq4: ssgiIj7R5c,
			_JFcGsAN:     peYZBpNwox.accounts[reikfgaLNK.addrHash],
			dTp3Yn8laadd: peYZBpNwox.storages[reikfgaLNK.addrHash],
			b0aCZklQW:    dNazL3Oz4,
			p3F8IHE:      tSTepB,
			aItHQjJHk:    peYZBpNwox.storagesOrigin[reikfgaLNK.address],
		})
		/*line HgrdF2oi.go:1*/ delete(peYZBpNwox.accounts, reikfgaLNK.addrHash)
		/*line CEn194pJ9i.go:1*/ delete(peYZBpNwox.storages, reikfgaLNK.addrHash)
		/*line bwyIkVSgl7.go:1*/ delete(peYZBpNwox.accountsOrigin, reikfgaLNK.address)
		/*line nP5CJGVyg_.go:1*/ delete(peYZBpNwox.storagesOrigin, reikfgaLNK.address)
	}
	/*line JC1EVmqLKs00.go:1*/ peYZBpNwox.wTQEdk(ikfPxhlba)
	if reikfgaLNK != nil && !reikfgaLNK.deleted {
		return ikfPxhlba, reikfgaLNK
	}
	return ikfPxhlba, nil
}

func (peYZBpNwox *StateDB) CreateAccount(hv1UkMRaKf common.Address) {
	llyAGqvcW, reikfgaLNK := /*line aMaZa3.go:1*/ peYZBpNwox.mXXLmsylBEwb(hv1UkMRaKf)
	if reikfgaLNK != nil {
		/*line ujhBNPABDnmj.go:1*/ llyAGqvcW.yNiGLlmM(reikfgaLNK.data.Balance)
	}
}

func (peYZBpNwox *StateDB) CreateTemporaryAccount(hv1UkMRaKf common.Address) *stateObject {
	llyAGqvcW := /*line z25l7eSh0R7l.go:1*/ zkEv6T(peYZBpNwox, hv1UkMRaKf, nil)
	/*line bsnF9kjIYTR.go:1*/ peYZBpNwox.journal.uwXnrqq(kLOasI4ckgD{huIthfk25W: &hv1UkMRaKf})
	/*line NUOLNK.go:1*/ peYZBpNwox.wTQEdk(llyAGqvcW)

	return llyAGqvcW
}

func (peYZBpNwox *StateDB) Copy() *StateDB {

	ayL8rlo := &StateDB{
		db: peYZBpNwox.db,
		trie:/*line t0ztyTgH32on.go:1*/ peYZBpNwox.db.CopyTrie(peYZBpNwox.trie),
		originalRoot: peYZBpNwox.originalRoot,
		accounts:/*line fhIdSqvr.go:1*/ make(map[common.Hash][]byte),
		storages:/*line QIJ28g8YFW.go:1*/ make(map[common.Hash]map[common.Hash][]byte),
		accountsOrigin:/*line G6S7JpFZZ.go:1*/ make(map[common.Address][]byte),
		storagesOrigin:/*line WLvEcza.go:1*/ make(map[common.Address]map[common.Hash][]byte),
		stateObjects:/*line C3b2zkkBdJx.go:1*/ make(map[common.Address]*stateObject /*line ueSWhru53.go:1*/, len(peYZBpNwox.journal.dirties)),
		stateObjectsPending:/*line So3VYMU14plI.go:1*/ make(map[common.Address]struct{} /*line xrUl96.go:1*/, len(peYZBpNwox.stateObjectsPending)),
		stateObjectsDirty:/*line a8JT1wY.go:1*/ make(map[common.Address]struct{} /*line axIzvNaQi.go:1*/, len(peYZBpNwox.journal.dirties)),
		stateObjectsDestruct:/*line G1ncO5.go:1*/ make(map[common.Address]*types.StateAccount /*line znGL60.go:1*/, len(peYZBpNwox.stateObjectsDestruct)),
		refund: peYZBpNwox.refund,
		logs:/*line ffdw4W6cjqt.go:1*/ make(map[common.Hash][]*types.Log /*line RQSeO5o.go:1*/, len(peYZBpNwox.logs)),
		logSize: peYZBpNwox.logSize,
		preimages:/*line or8gJi.go:1*/ make(map[common.Hash][]byte /*line lWIdUeeK0.go:1*/, len(peYZBpNwox.preimages)),
		journal:/*line FjWuDKbQi.go:1*/ vuoy40vfttW(),
		hasher:/*line ArCFSgI.go:1*/ crypto.NewKeccakState(),

		snaps: peYZBpNwox.snaps,
		snap:  peYZBpNwox.snap,
	}

	for hv1UkMRaKf := range peYZBpNwox.journal.dirties {

		if gds4Bcqt, _bRYnhxrC := peYZBpNwox.stateObjects[hv1UkMRaKf]; _bRYnhxrC {

			ayL8rlo.stateObjects[hv1UkMRaKf] = /*line DHrAYS15.go:1*/ gds4Bcqt.eiV2C2(ayL8rlo)

			ayL8rlo.stateObjectsDirty[hv1UkMRaKf] = struct{}{}
			ayL8rlo.stateObjectsPending[hv1UkMRaKf] = struct{}{}
		}
	}

	for hv1UkMRaKf := range peYZBpNwox.stateObjectsPending {
		if _, _bRYnhxrC := ayL8rlo.stateObjects[hv1UkMRaKf]; !_bRYnhxrC {
			ayL8rlo.stateObjects[hv1UkMRaKf] = /*line dGNepuAWNTb.go:1*/ peYZBpNwox.stateObjects[hv1UkMRaKf].eiV2C2(ayL8rlo)
		}
		ayL8rlo.stateObjectsPending[hv1UkMRaKf] = struct{}{}
	}
	for hv1UkMRaKf := range peYZBpNwox.stateObjectsDirty {
		if _, _bRYnhxrC := ayL8rlo.stateObjects[hv1UkMRaKf]; !_bRYnhxrC {
			ayL8rlo.stateObjects[hv1UkMRaKf] = /*line iLhi7DeuYn.go:1*/ peYZBpNwox.stateObjects[hv1UkMRaKf].eiV2C2(ayL8rlo)
		}
		ayL8rlo.stateObjectsDirty[hv1UkMRaKf] = struct{}{}
	}

	for hv1UkMRaKf, fnngafl := range peYZBpNwox.stateObjectsDestruct {
		ayL8rlo.stateObjectsDestruct[hv1UkMRaKf] = fnngafl
	}

	ayL8rlo.accounts = /*line iEX8IAXhh.go:1*/ eh1lN9hBaUMv(peYZBpNwox.accounts)
	ayL8rlo.storages = /*line X9rjpTLY.go:1*/ ngMfLQ4pubvm(peYZBpNwox.storages)
	ayL8rlo.accountsOrigin = /*line KS5Zr8V.go:1*/ eh1lN9hBaUMv(ayL8rlo.accountsOrigin)
	ayL8rlo.storagesOrigin = /*line QOHRCjjvzu.go:1*/ ngMfLQ4pubvm(ayL8rlo.storagesOrigin)

	for cSqPoIL2n5, jrZPAbGCstr := range peYZBpNwox.logs {
		mMvShC5AUZkj := /*line xo1xTTNua.go:1*/ make([]*types.Log /*line F4ivCXwJVJN.go:1*/, len(jrZPAbGCstr))
		for cDClIaFDS, fuNTsaDfte := range jrZPAbGCstr {
			mMvShC5AUZkj[cDClIaFDS] = /*line XySMNeF.go:1*/ new(types.Log)
			*mMvShC5AUZkj[cDClIaFDS] = *fuNTsaDfte
		}
		ayL8rlo.logs[cSqPoIL2n5] = mMvShC5AUZkj
	}

	for cSqPoIL2n5, ikSYP8 := range peYZBpNwox.preimages {
		ayL8rlo.preimages[cSqPoIL2n5] = ikSYP8
	}

	ayL8rlo.accessList = /*line pJ0xFd.go:1*/ peYZBpNwox.accessList.Copy()
	ayL8rlo.transientStorage = /*line WJMI4lPpR1KN.go:1*/ peYZBpNwox.transientStorage.Copy()

	if peYZBpNwox.prefetcher != nil {
		ayL8rlo.prefetcher = /*line w6ekbDfzo4.go:1*/ peYZBpNwox.prefetcher.stapkw2aW()
	}
	return ayL8rlo
}

func (peYZBpNwox *StateDB) Snapshot() int {
	gVThh7O := peYZBpNwox.nextRevisionId
	peYZBpNwox.nextRevisionId++
	peYZBpNwox.validRevisions = /*line ZWQ1M6VVR84.go:1*/ append(peYZBpNwox.validRevisions, revision{gVThh7O /*line AeLbsSk.go:1*/, peYZBpNwox.journal.fgH6Ck43r()})
	return gVThh7O
}

func (peYZBpNwox *StateDB) RevertToSnapshot(vSFoChk int) {

	iJQtkRishD3g := /*line f8SYhPs.go:1*/ sort.Search( /*line lFnKoxX.go:1*/ len(peYZBpNwox.validRevisions), func(cDClIaFDS int) bool {
		return peYZBpNwox.validRevisions[cDClIaFDS].id >= vSFoChk
	})
	if iJQtkRishD3g == /*line kFo34TMX.go:1*/ len(peYZBpNwox.validRevisions) || peYZBpNwox.validRevisions[iJQtkRishD3g].id != vSFoChk {
		/*line HbGRDAQhR34E.go:1*/ panic( /*line JvsAvlKWc.go:1*/ fmt.Errorf(func() /*line A5vmezsMVHK9.go:1*/ string {
			seed := /*line A5vmezsMVHK9.go:1*/ byte(250)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line A5vmezsMVHK9.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line A5vmezsMVHK9.go:1*/ fnc(136)(231)(31)(225)(26)(234)(2)(1)(80)(169)(13)(86)(233)(195)(88)(179)(226)(11)(30)(225)(27)(170)(86)(239)(89)(160)(23)(255)(237)(7)(8)(225)(1)
			return /*line A5vmezsMVHK9.go:1*/ string(data)
		}(), vSFoChk))
	}
	anRdfDemah := peYZBpNwox.validRevisions[iJQtkRishD3g].journalIndex

	/*line G6v9Cs9.go:1*/
	peYZBpNwox.journal.oOOlwUrQ0Ck(peYZBpNwox, anRdfDemah)
	peYZBpNwox.validRevisions = peYZBpNwox.validRevisions[:iJQtkRishD3g]
}

func (peYZBpNwox *StateDB) GetRefund() uint64 {
	return peYZBpNwox.refund
}

func (peYZBpNwox *StateDB) Finalise(qdnMhVci bool) {
	rKWa5M_ZQh := /*line TipN_c8Wej.go:1*/ make([][]byte, 0 /*line X_NbnIaK.go:1*/, len(peYZBpNwox.journal.dirties))
	for hv1UkMRaKf := range peYZBpNwox.journal.dirties {
		ws3tNR3MU, _bRYnhxrC := peYZBpNwox.stateObjects[hv1UkMRaKf]
		if !_bRYnhxrC {

			continue
		}
		if ws3tNR3MU.selfDestructed || (qdnMhVci && /*line gqYnXFpC.go:1*/ ws3tNR3MU.usx3KByCvUvp()) {
			ws3tNR3MU.deleted = true

			if _, dNazL3Oz4 := peYZBpNwox.stateObjectsDestruct[ws3tNR3MU.address]; !dNazL3Oz4 {
				peYZBpNwox.stateObjectsDestruct[ws3tNR3MU.address] = ws3tNR3MU.origin
			}

			/*line gdkDtSSAklf.go:1*/
			delete(peYZBpNwox.accounts, ws3tNR3MU.addrHash)
			/*line QCyXbyaFx.go:1*/ delete(peYZBpNwox.storages, ws3tNR3MU.addrHash)
			/*line d429zk.go:1*/ delete(peYZBpNwox.accountsOrigin, ws3tNR3MU.address)
			/*line FQ6ZnLQX.go:1*/ delete(peYZBpNwox.storagesOrigin, ws3tNR3MU.address)
		} else {
			/*line iazfEdOyEPLs.go:1*/ ws3tNR3MU.tatrqu(true)
		}
		ws3tNR3MU.created = false
		peYZBpNwox.stateObjectsPending[hv1UkMRaKf] = struct{}{}
		peYZBpNwox.stateObjectsDirty[hv1UkMRaKf] = struct{}{}

		rKWa5M_ZQh = /*line ZX8oDbbqU.go:1*/ append(rKWa5M_ZQh /*line Mtvb51fX.go:1*/, common.CopyBytes(hv1UkMRaKf[:]))
	}
	if peYZBpNwox.prefetcher != nil && /*line inrAUu.go:1*/ len(rKWa5M_ZQh) > 0 {
		/*line w0i67Lcd2.go:1*/ peYZBpNwox.prefetcher.i3uM6DE(common.Hash{}, peYZBpNwox.originalRoot, common.Address{}, rKWa5M_ZQh)
	}

	/*line duz3Kbe42.go:1*/
	peYZBpNwox.vbJ77c8A()
}

func (peYZBpNwox *StateDB) IntermediateRoot(qdnMhVci bool) common.Hash {

	/*line W008OdWG.go:1*/
	peYZBpNwox.Finalise(qdnMhVci)

	lS6Suv4 := peYZBpNwox.prefetcher
	if peYZBpNwox.prefetcher != nil {
		defer func() {
			/*line fdQQori.go:1*/ peYZBpNwox.prefetcher.gfibELOKGAub()
			peYZBpNwox.prefetcher = nil
		}()
	}

	for hv1UkMRaKf := range peYZBpNwox.stateObjectsPending {
		if ws3tNR3MU := peYZBpNwox.stateObjects[hv1UkMRaKf]; !ws3tNR3MU.deleted {
			/*line _p3Fb4zUaF.go:1*/ ws3tNR3MU.hbhGaAPh6VSz()
		}
	}

	if lS6Suv4 != nil {
		if fQodR9c7t := /*line vGOY5OSgtcZ.go:1*/ lS6Suv4.fQodR9c7t(common.Hash{}, peYZBpNwox.originalRoot); fQodR9c7t != nil {
			peYZBpNwox.trie = fQodR9c7t
		}
	}
	eQfnePg := /*line JcFTuhZoN0QD.go:1*/ make([][]byte, 0 /*line bFvmquuken.go:1*/, len(peYZBpNwox.stateObjectsPending))
	for hv1UkMRaKf := range peYZBpNwox.stateObjectsPending {
		if ws3tNR3MU := peYZBpNwox.stateObjects[hv1UkMRaKf]; ws3tNR3MU.deleted {
			/*line E8vBaSm2Sn9c.go:1*/ peYZBpNwox.weye1mu(ws3tNR3MU)
			peYZBpNwox.AccountDeleted += 1
		} else {
			/*line amAEkVEsAFH.go:1*/ peYZBpNwox.yXyxeb11X(ws3tNR3MU)
			peYZBpNwox.AccountUpdated += 1
		}
		eQfnePg = /*line NLF6PIV_QP.go:1*/ append(eQfnePg /*line Zto5WJ9s7.go:1*/, common.CopyBytes(hv1UkMRaKf[:]))
	}
	if lS6Suv4 != nil {
		/*line TRfqnMXS.go:1*/ lS6Suv4.t50cl4u8e(common.Hash{}, peYZBpNwox.originalRoot, eQfnePg)
	}
	if /*line Ykk3NTwMPFER.go:1*/ len(peYZBpNwox.stateObjectsPending) > 0 {
		peYZBpNwox.stateObjectsPending = /*line EfbWLXhrtOo.go:1*/ make(map[common.Address]struct{})
	}

	if metrics.EnabledExpensive {
		defer func( /*line aw2waTPq.go:1*/ stqEk4Brz1wF time.Time) {
			peYZBpNwox.AccountHashes += /*line FzAcuraAkD4.go:1*/ time.Since(stqEk4Brz1wF)
		}( /*line C0os7NU.go:1*/ time.Now())
	}

	return /*line ywzzX2w9.go:1*/ peYZBpNwox.trie.Hash()
}

func (peYZBpNwox *StateDB) SetTxContext(vj_kb4Zaw27c common.Hash, whYnPzBO2 int) {
	peYZBpNwox.thash = vj_kb4Zaw27c
	peYZBpNwox.txIndex = whYnPzBO2
}

func (peYZBpNwox *StateDB) vbJ77c8A() {
	if /*line rGKLgIME7.go:1*/ len(peYZBpNwox.journal.entries) > 0 {
		peYZBpNwox.journal = /*line EdM_Da.go:1*/ vuoy40vfttW()
		peYZBpNwox.refund = 0
	}
	peYZBpNwox.validRevisions = peYZBpNwox.validRevisions[:0]
}

func (peYZBpNwox *StateDB) ppE03tfQO(bRSewLor2S common.Hash, lxItfhw common.Hash) (bool, common.StorageSize, map[common.Hash][]byte, *trienode.NodeSet, error) {
	x60KDSaRsxSR, cYHA75qVq := /*line urBRL_TkQs2l.go:1*/ peYZBpNwox.snaps.StorageIterator(peYZBpNwox.originalRoot, bRSewLor2S, common.Hash{})
	if cYHA75qVq != nil {
		return false, 0, nil, nil, cYHA75qVq
	}
	defer /*line ptz8vy5JepC.go:1*/ x60KDSaRsxSR.Release()

	var (
		htUnPTZWCIPk common.StorageSize
		nQXC7Td      = /*line IP2WFYgraTzn.go:1*/ trienode.NewNodeSet(bRSewLor2S)
		cZ1Z_E5qkNa  = /*line tOD5a2iHp9.go:1*/ make(map[common.Hash][]byte)
	)
	d1U7tZu := /*line DazyYKR.go:1*/ trie.HN5J4sr7gxe()
	d1U7tZu = /*line XX8rrbbZjSC.go:1*/ d1U7tZu.WithWriter(func(dbphQ9 []byte, cSqPoIL2n5 common.Hash, m7cu_z []byte) {
		/*line B572Pcjop5.go:1*/ nQXC7Td.AddNode(dbphQ9 /*line bPZu6ina2m.go:1*/, trienode.NewDeleted())
		htUnPTZWCIPk += /*line JJSB1JY.go:1*/ common.StorageSize( /*line kT0wLGRdf.go:1*/ len(dbphQ9))
	})
	jwPBVnjlhG := /*line IQiZd5.go:1*/ trie.KYaYz2rvh7Rt(d1U7tZu)
	for /*line Hn3gDlMd.go:1*/ x60KDSaRsxSR.Next() {
		if htUnPTZWCIPk > storageDeleteLimit {
			return true, htUnPTZWCIPk, nil, nil, nil
		}
		ibmfxwhUDRC := /*line Rb_QwXT6.go:1*/ common.CopyBytes( /*line vg5bSJMgKU.go:1*/ x60KDSaRsxSR.Slot())
		if cYHA75qVq := /*line COp2Emo4Sfq.go:1*/ x60KDSaRsxSR.Error(); cYHA75qVq != nil {
			return false, 0, nil, nil, cYHA75qVq
		}
		htUnPTZWCIPk += /*line QuUGrbIjck.go:1*/ common.StorageSize(common.HashLength + /*line HK2SgHaB.go:1*/ len(ibmfxwhUDRC))
		cZ1Z_E5qkNa[ /*line pjr6pMHYAV5.go:1*/ x60KDSaRsxSR.Hash()] = ibmfxwhUDRC

		if cYHA75qVq := /*line kv4xWGvYsaoc.go:1*/ jwPBVnjlhG.Update( /*line CNETai7ab.go:1*/ x60KDSaRsxSR.Hash().Bytes(), ibmfxwhUDRC); cYHA75qVq != nil {
			return false, 0, nil, nil, cYHA75qVq
		}
	}
	if cYHA75qVq := /*line iYhfXQo.go:1*/ x60KDSaRsxSR.Error(); cYHA75qVq != nil {
		return false, 0, nil, nil, cYHA75qVq
	}
	if /*line Xb6KdGEEH.go:1*/ jwPBVnjlhG.Hash() != lxItfhw {
		return false, 0, nil, nil /*line WeaDQk5A.go:1*/, fmt.Errorf(func() /*line u2QG48H.go:1*/ string {
			key := [] /*line u2QG48H.go:1*/ byte("\U0008672c\x89m\xa5\xc9\b(\xbeB\xa3\xc5\a\x84\xc3U\x16\xdd\xe0@\xef@^<dJ\xfdӤ\x88\xe6\xfdC\x89N\x00\xf6")
			data := [] /*line u2QG48H.go:1*/ byte("\x81\xe8\xc5\xc4\xea\xfbʫ\x18A\xb5\xde˪m\x9c\xaa\f^\x86\x88%u\xec\xc2)\x14&#RԤ:j,\xeb\xd2%\x82")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line u2QG48H.go:1*/ string(data)
		}(), lxItfhw /*line PZvii9P4VkA.go:1*/, jwPBVnjlhG.Hash())
	}
	return false, htUnPTZWCIPk, cZ1Z_E5qkNa, nQXC7Td, nil
}

func (peYZBpNwox *StateDB) i2TnpHp1vCxW(hv1UkMRaKf common.Address, bRSewLor2S common.Hash, lxItfhw common.Hash) (bool, common.StorageSize, map[common.Hash][]byte, *trienode.NodeSet, error) {
	r8CzlEu, cYHA75qVq := /*line zaKMX0Tx4pq.go:1*/ peYZBpNwox.db.OpenStorageTrie(peYZBpNwox.originalRoot, hv1UkMRaKf, lxItfhw, peYZBpNwox.trie)
	if cYHA75qVq != nil {
		return false, 0, nil, nil /*line CyNqaw4k.go:1*/, fmt.Errorf(func() /*line FYV2rb6C.go:1*/ string {
			fullData := [] /*line FYV2rb6C.go:1*/ byte("\x14n֨Fm\x00\xd2}|\x85\x8f\x1ac:\x9b\t<\\9\x06\xb5}s\x95\x83 \xdb\xd7o\xec{\xe6\xf4\xb5\xcc\v\xa4\x184\xee9s\xfc\xf4\x1fdPUR37\x92f\xa1\x88s\x02\a\xa6\x01c\xe7\x1c\x03\xc6֠\xa2\xad\xaf`")
			idxKey := [] /*line FYV2rb6C.go:1*/ byte("\xf4\xc7g{0=u\t\xa0;AB\x89H^\xfe")
			data := /*line FYV2rb6C.go:1*/ make([]byte, 0, 37)
			data = /*line FYV2rb6C.go:1*/ append(data, fullData[112^ /*line FYV2rb6C.go:1*/ int(idxKey[11])]^fullData[114^ /*line FYV2rb6C.go:1*/ int(idxKey[11])], fullData[208^ /*line FYV2rb6C.go:1*/ int(idxKey[1])]+fullData[239^ /*line FYV2rb6C.go:1*/ int(idxKey[1])], fullData[29^ /*line FYV2rb6C.go:1*/ int(idxKey[5])]-fullData[53^ /*line FYV2rb6C.go:1*/ int(idxKey[5])], fullData[22^ /*line FYV2rb6C.go:1*/ int(idxKey[9])]^fullData[3^ /*line FYV2rb6C.go:1*/ int(idxKey[9])], fullData[76^ /*line FYV2rb6C.go:1*/ int(idxKey[6])]+fullData[72^ /*line FYV2rb6C.go:1*/ int(idxKey[6])], fullData[228^ /*line FYV2rb6C.go:1*/ int(idxKey[8])]^fullData[225^ /*line FYV2rb6C.go:1*/ int(idxKey[8])], fullData[98^ /*line FYV2rb6C.go:1*/ int(idxKey[3])]-fullData[118^ /*line FYV2rb6C.go:1*/ int(idxKey[3])], fullData[241^ /*line FYV2rb6C.go:1*/ int(idxKey[0])]+fullData[206^ /*line FYV2rb6C.go:1*/ int(idxKey[0])], fullData[71^ /*line FYV2rb6C.go:1*/ int(idxKey[10])]+fullData[92^ /*line FYV2rb6C.go:1*/ int(idxKey[10])], fullData[116^ /*line FYV2rb6C.go:1*/ int(idxKey[14])]+fullData[27^ /*line FYV2rb6C.go:1*/ int(idxKey[14])], fullData[151^ /*line FYV2rb6C.go:1*/ int(idxKey[12])]-fullData[159^ /*line FYV2rb6C.go:1*/ int(idxKey[12])], fullData[93^ /*line FYV2rb6C.go:1*/ int(idxKey[11])]-fullData[102^ /*line FYV2rb6C.go:1*/ int(idxKey[11])], fullData[69^ /*line FYV2rb6C.go:1*/ int(idxKey[2])]-fullData[72^ /*line FYV2rb6C.go:1*/ int(idxKey[2])], fullData[189^ /*line FYV2rb6C.go:1*/ int(idxKey[12])]^fullData[162^ /*line FYV2rb6C.go:1*/ int(idxKey[12])], fullData[83^ /*line FYV2rb6C.go:1*/ int(idxKey[11])]-fullData[125^ /*line FYV2rb6C.go:1*/ int(idxKey[11])], fullData[135^ /*line FYV2rb6C.go:1*/ int(idxKey[12])]+fullData[160^ /*line FYV2rb6C.go:1*/ int(idxKey[12])], fullData[85^ /*line FYV2rb6C.go:1*/ int(idxKey[10])]+fullData[64^ /*line FYV2rb6C.go:1*/ int(idxKey[10])], fullData[10^ /*line FYV2rb6C.go:1*/ int(idxKey[5])]+fullData[3^ /*line FYV2rb6C.go:1*/ int(idxKey[5])], fullData[111^ /*line FYV2rb6C.go:1*/ int(idxKey[6])]+fullData[68^ /*line FYV2rb6C.go:1*/ int(idxKey[6])], fullData[40^ /*line FYV2rb6C.go:1*/ int(idxKey[4])]^fullData[17^ /*line FYV2rb6C.go:1*/ int(idxKey[4])], fullData[114^ /*line FYV2rb6C.go:1*/ int(idxKey[2])]^fullData[96^ /*line FYV2rb6C.go:1*/ int(idxKey[2])], fullData[108^ /*line FYV2rb6C.go:1*/ int(idxKey[2])]+fullData[37^ /*line FYV2rb6C.go:1*/ int(idxKey[2])], fullData[72^ /*line FYV2rb6C.go:1*/ int(idxKey[13])]-fullData[100^ /*line FYV2rb6C.go:1*/ int(idxKey[13])], fullData[230^ /*line FYV2rb6C.go:1*/ int(idxKey[8])]^fullData[187^ /*line FYV2rb6C.go:1*/ int(idxKey[8])], fullData[207^ /*line FYV2rb6C.go:1*/ int(idxKey[0])]+fullData[215^ /*line FYV2rb6C.go:1*/ int(idxKey[0])], fullData[54^ /*line FYV2rb6C.go:1*/ int(idxKey[6])]-fullData[70^ /*line FYV2rb6C.go:1*/ int(idxKey[6])], fullData[215^ /*line FYV2rb6C.go:1*/ int(idxKey[1])]+fullData[213^ /*line FYV2rb6C.go:1*/ int(idxKey[1])], fullData[102^ /*line FYV2rb6C.go:1*/ int(idxKey[10])]^fullData[103^ /*line FYV2rb6C.go:1*/ int(idxKey[10])], fullData[250^ /*line FYV2rb6C.go:1*/ int(idxKey[15])]^fullData[203^ /*line FYV2rb6C.go:1*/ int(idxKey[15])], fullData[126^ /*line FYV2rb6C.go:1*/ int(idxKey[11])]^fullData[108^ /*line FYV2rb6C.go:1*/ int(idxKey[11])], fullData[203^ /*line FYV2rb6C.go:1*/ int(idxKey[1])]-fullData[196^ /*line FYV2rb6C.go:1*/ int(idxKey[1])], fullData[84^ /*line FYV2rb6C.go:1*/ int(idxKey[13])]+fullData[71^ /*line FYV2rb6C.go:1*/ int(idxKey[13])], fullData[59^ /*line FYV2rb6C.go:1*/ int(idxKey[3])]^fullData[104^ /*line FYV2rb6C.go:1*/ int(idxKey[3])], fullData[209^ /*line FYV2rb6C.go:1*/ int(idxKey[0])]+fullData[253^ /*line FYV2rb6C.go:1*/ int(idxKey[0])], fullData[49^ /*line FYV2rb6C.go:1*/ int(idxKey[9])]-fullData[124^ /*line FYV2rb6C.go:1*/ int(idxKey[9])], fullData[63^ /*line FYV2rb6C.go:1*/ int(idxKey[7])]+fullData[11^ /*line FYV2rb6C.go:1*/ int(idxKey[7])])
			return /*line FYV2rb6C.go:1*/ string(data)
		}(), cYHA75qVq)
	}
	s8hOm6P, cYHA75qVq := /*line SFVmjr6.go:1*/ r8CzlEu.NodeIterator(nil)
	if cYHA75qVq != nil {
		return false, 0, nil, nil /*line qEWl2dT.go:1*/, fmt.Errorf(func() /*line aJCBu9b.go:1*/ string {
			seed := /*line aJCBu9b.go:1*/ byte(190)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line aJCBu9b.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line aJCBu9b.go:1*/ fnc(36)(67)(142)(31)(55)(109)(150)(128)(251)(167)(157)(59)(107)(223)(112)(51)(103)(201)(149)(25)(56)(110)(151)(119)(249)(227)(211)(149)(61)(117)(237)(148)(28)(125)(7)(14)(228)(174)(97)(20)
			return /*line aJCBu9b.go:1*/ string(data)
		}(), cYHA75qVq)
	}
	var (
		htUnPTZWCIPk common.StorageSize
		nQXC7Td      = /*line NSXUwn.go:1*/ trienode.NewNodeSet(bRSewLor2S)
		cZ1Z_E5qkNa  = /*line jwPYic427qNT.go:1*/ make(map[common.Hash][]byte)
	)
	for /*line zV6mqcHMs.go:1*/ s8hOm6P.Next(true) {
		if htUnPTZWCIPk > storageDeleteLimit {
			return true, htUnPTZWCIPk, nil, nil, nil
		}
		if /*line U4mf24O.go:1*/ s8hOm6P.Leaf() {
			cZ1Z_E5qkNa[ /*line adKHjThArigY.go:1*/ common.BytesToHash( /*line U4hDeJ.go:1*/ s8hOm6P.LeafKey())] = /*line DO4v1KyMTG.go:1*/ common.CopyBytes( /*line g3nhe8o.go:1*/ s8hOm6P.LeafBlob())
			htUnPTZWCIPk += /*line XgJp7GqMyHXY.go:1*/ common.StorageSize(common.HashLength + /*line EaLTgAxd.go:1*/ len( /*line slzFfgScKOq.go:1*/ s8hOm6P.LeafBlob()))
			continue
		}
		if /*line hCMeqO0qXas.go:1*/ s8hOm6P.Hash() == (common.Hash{}) {
			continue
		}
		htUnPTZWCIPk += /*line KVKlJ2WilPu.go:1*/ common.StorageSize( /*line B3ALax.go:1*/ len( /*line W1EpsREm.go:1*/ s8hOm6P.Path()))
		/*line AFanko.go:1*/ nQXC7Td.AddNode( /*line uWAW3uJfl3.go:1*/ s8hOm6P.Path() /*line EzqxTbtN.go:1*/, trienode.NewDeleted())
	}
	if cYHA75qVq := /*line dmurBg.go:1*/ s8hOm6P.Error(); cYHA75qVq != nil {
		return false, 0, nil, nil, cYHA75qVq
	}
	return false, htUnPTZWCIPk, cZ1Z_E5qkNa, nQXC7Td, nil
}

func (peYZBpNwox *StateDB) j4_Yr4ciW(hv1UkMRaKf common.Address, bRSewLor2S common.Hash, lxItfhw common.Hash) (bool, map[common.Hash][]byte, *trienode.NodeSet, error) {
	var (
		stqEk4Brz1wF = /*line VnwtQ2.go:1*/ time.Now()
		cYHA75qVq    error
		tzHJLW_      bool
		htUnPTZWCIPk common.StorageSize
		cZ1Z_E5qkNa  map[common.Hash][]byte
		nQXC7Td      *trienode.NodeSet
	)

	if peYZBpNwox.snap != nil {
		tzHJLW_, htUnPTZWCIPk, cZ1Z_E5qkNa, nQXC7Td, cYHA75qVq = /*line X7404jAQaPE.go:1*/ peYZBpNwox.ppE03tfQO(bRSewLor2S, lxItfhw)
	}
	if peYZBpNwox.snap == nil || cYHA75qVq != nil {
		tzHJLW_, htUnPTZWCIPk, cZ1Z_E5qkNa, nQXC7Td, cYHA75qVq = /*line OQaC4eUV_wX0.go:1*/ peYZBpNwox.i2TnpHp1vCxW(hv1UkMRaKf, bRSewLor2S, lxItfhw)
	}
	if cYHA75qVq != nil {
		return false, nil, nil, cYHA75qVq
	}
	if metrics.EnabledExpensive {
		if tzHJLW_ {
			/*line L3lEq6shFGop.go:1*/ mJkh5y.Inc(1)
		}
		t6rBh2SfX := /*line q8efTo.go:1*/ int64( /*line nqhr47i4rMb.go:1*/ len(cZ1Z_E5qkNa))

		/*line CVibqU2ccu.go:1*/
		bKowp_.UpdateIfGt( /*line wpoEec.go:1*/ int64( /*line F4D4esytg7.go:1*/ len(cZ1Z_E5qkNa)))
		/*line CMTTZ9QMzMwl.go:1*/ qOuP3t.UpdateIfGt( /*line k8rX7RvMk.go:1*/ int64(htUnPTZWCIPk))

		/*line Q68Rn50gRE.go:1*/
		fxF6osapsiIK.UpdateSince(stqEk4Brz1wF)
		/*line j_OJxQ.go:1*/ jDTEIE_.Mark(t6rBh2SfX)
		/*line MmbW7Y.go:1*/ qt2Hcoh5RFFq.Mark( /*line F_wg7G.go:1*/ int64(htUnPTZWCIPk))
	}
	return tzHJLW_, cZ1Z_E5qkNa, nQXC7Td, nil
}

func (peYZBpNwox *StateDB) mC9AgIi(nQXC7Td *trienode.MergedNodeSet) (map[common.Address]struct{}, error) {

	avY0t4RP1I := /*line rJsSETZN.go:1*/ make(map[common.Address]struct{})
	if /*line sB9a2v.go:1*/ peYZBpNwox.db.TrieDB().Scheme() == rawdb.HashScheme {
		return avY0t4RP1I, nil
	}
	for hv1UkMRaKf, reikfgaLNK := range peYZBpNwox.stateObjectsDestruct {

		bRSewLor2S := /*line B5Wu4oyZe.go:1*/ crypto.Keccak256Hash(hv1UkMRaKf[:])
		if reikfgaLNK == nil {
			if _, dNazL3Oz4 := peYZBpNwox.accounts[bRSewLor2S]; dNazL3Oz4 {
				peYZBpNwox.accountsOrigin[hv1UkMRaKf] = nil
			}
			continue
		}

		peYZBpNwox.accountsOrigin[hv1UkMRaKf] = /*line DR9R6Ru_mZ.go:1*/ types.K2bJTLXfmaB(*reikfgaLNK)

		if reikfgaLNK.Root == types.NrGuaNA21 {
			continue
		}

		tzHJLW_, cZ1Z_E5qkNa, ifeYDU90, cYHA75qVq := /*line DJgfvZ4BHmHv.go:1*/ peYZBpNwox.j4_Yr4ciW(hv1UkMRaKf, bRSewLor2S, reikfgaLNK.Root)
		if cYHA75qVq != nil {
			return nil /*line HScIR_St.go:1*/, fmt.Errorf(func() /*line tdTxaAS66RSO.go:1*/ string {
				data := [] /*line tdTxaAS66RSO.go:1*/ byte("f\xa7c\x1aed\xe9\xcdͳ\xe0e\xc8Ptx st\x81FǠx\xb4 ñrP \xdco")
				positions := [...]byte{10, 26, 3, 8, 2, 6, 9, 8, 20, 9, 27, 12, 21, 12, 31, 27, 26, 1, 24, 13, 7, 2, 27, 3, 10, 26, 22, 12, 1, 32, 29, 31, 23, 15, 19, 2, 19, 3, 13, 12, 22, 22}
				for i := 0; i < 42; i += 2 {
					localKey := /*line tdTxaAS66RSO.go:1*/ byte(i) + /*line tdTxaAS66RSO.go:1*/ byte(positions[i]^positions[i+1]) + 181
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line tdTxaAS66RSO.go:1*/ string(data)
			}(), cYHA75qVq)
		}

		if tzHJLW_ {
			avY0t4RP1I[hv1UkMRaKf] = struct{}{}
			/*line P2mM2pzYbunC.go:1*/ delete(peYZBpNwox.storagesOrigin, hv1UkMRaKf)
			continue
		}
		if peYZBpNwox.storagesOrigin[hv1UkMRaKf] == nil {
			peYZBpNwox.storagesOrigin[hv1UkMRaKf] = cZ1Z_E5qkNa
		} else {

			for ixeygbAgOap, cGzt0N := range cZ1Z_E5qkNa {
				peYZBpNwox.storagesOrigin[hv1UkMRaKf][ixeygbAgOap] = cGzt0N
			}
		}
		if cYHA75qVq := /*line WN_YimDINSZ.go:1*/ nQXC7Td.Merge(ifeYDU90); cYHA75qVq != nil {
			return nil, cYHA75qVq
		}
	}
	return avY0t4RP1I, nil
}

func (peYZBpNwox *StateDB) Commit(cCPWpPefu uint64, qdnMhVci bool) (common.Hash, error) {

	if peYZBpNwox.dbErr != nil {
		return common.Hash{} /*line tswbS9pR.go:1*/, fmt.Errorf(func() /*line COzvAQGCQfr.go:1*/ string {
			data := /*line COzvAQGCQfr.go:1*/ make([]byte, 0, 40)
			i := 0
			decryptKey := 7
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 13
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xee\xf8\xc7\xdb"...,
					)
				case 1:
					i = 15
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xe6\xe3\xff\xe0"...,
					)
				case 0:
					i = 2
					data = /*line COzvAQGCQfr.go:1*/ append(data, 246)
				case 16:
					i = 1
					data = /*line COzvAQGCQfr.go:1*/ append(data, 160)
				case 7:
					i = 6
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xe4\xb3\xf3"...,
					)
				case 13:
					i = 4
					data = /*line COzvAQGCQfr.go:1*/ append(data, 197)
				case 8:
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xfc\xfc\xbb\xfe"...,
					)
					i = 14
				case 9:
					data = /*line COzvAQGCQfr.go:1*/ append(data, 234)
					i = 8
				case 11:
					data = /*line COzvAQGCQfr.go:1*/ append(data, 197)
					i = 12
				case 5:
					i = 16
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xe1\xa7\xf2\xee"...,
					)
				case 14:
					data = /*line COzvAQGCQfr.go:1*/ append(data, 240)
					i = 5
				case 4:
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\x8c\x91\x95"...,
					)
					i = 11
				case 12:
					for y := range data {
						data[y] = data[y] ^ /*line COzvAQGCQfr.go:1*/ byte(decryptKey^y)
					}
					i = 3
				case 6:
					i = 9
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xff\xf3\xed"...,
					)
				case 2:
					i = 7
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xfb\xfa\xfb\xf8"...,
					)
				case 15:
					i = 10
					data = /*line COzvAQGCQfr.go:1*/ append(data, "\xe6\xeb\xfb\xa8"...,
					)
				}
			}
			return /*line COzvAQGCQfr.go:1*/ string(data)
		}(), peYZBpNwox.dbErr)
	}

	/*line wa51WapL.go:1*/
	peYZBpNwox.IntermediateRoot(qdnMhVci)

	var (
		ayLK4KQXyY   int
		fGenfN868u5  int
		xT719aaps    int
		cjQJXhTj     int
		nQXC7Td      = /*line dEv3HrfpZn8_.go:1*/ trienode.NewMergedNodeSet()
		j7iae69otGcZ = /*line kxNM09BYy.go:1*/ peYZBpNwox.db.DiskDB().NewBatch()
	)

	avY0t4RP1I, cYHA75qVq := /*line F_IH60X.go:1*/ peYZBpNwox.mC9AgIi(nQXC7Td)
	if cYHA75qVq != nil {
		return common.Hash{}, cYHA75qVq
	}

	for hv1UkMRaKf := range peYZBpNwox.stateObjectsDirty {
		ws3tNR3MU := peYZBpNwox.stateObjects[hv1UkMRaKf]
		if ws3tNR3MU.deleted {
			continue
		}

		if ws3tNR3MU.code != nil && ws3tNR3MU.dirtyCode {
			/*line BaOF8VazJY.go:1*/ rawdb.WriteCode(j7iae69otGcZ /*line SJI2BV.go:1*/, common.BytesToHash( /*line eBYrAC_P_.go:1*/ ws3tNR3MU.CodeHash()), ws3tNR3MU.code)
			ws3tNR3MU.dirtyCode = false
		}

		ifeYDU90, cYHA75qVq := /*line Zc0pHPEj2sg.go:1*/ ws3tNR3MU.xGnpP8a5()
		if cYHA75qVq != nil {
			return common.Hash{}, cYHA75qVq
		}

		if ifeYDU90 != nil {
			if cYHA75qVq := /*line HOTCDV9iNKv.go:1*/ nQXC7Td.Merge(ifeYDU90); cYHA75qVq != nil {
				return common.Hash{}, cYHA75qVq
			}
			sD9HmJFewE7, gnLNi6M := /*line sqhg2ZNlv5.go:1*/ ifeYDU90.Size()
			xT719aaps += sD9HmJFewE7
			cjQJXhTj += gnLNi6M
		}
	}
	if /*line SdoODb.go:1*/ j7iae69otGcZ.ValueSize() > 0 {
		if cYHA75qVq := /*line _UzTKiJ.go:1*/ j7iae69otGcZ.Write(); cYHA75qVq != nil {
			/*line b2GyW9C.go:1*/ log.Crit(func() /*line OpTkDPbc4LGC.go:1*/ string {
				data := /*line OpTkDPbc4LGC.go:1*/ make([]byte, 0, 29)
				i := 5
				decryptKey := 118
				for counter := 0; i != 6; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 8:
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xb0o"...,
						)
						i = 10
					case 1:
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xbc\xc4\xcc\xd1"...,
						)
						i = 4
					case 9:
						i = 6
						for y := range data {
							data[y] = data[y] - /*line OpTkDPbc4LGC.go:1*/ byte(decryptKey^y)
						}
					case 5:
						i = 7
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, 143)
					case 2:
						i = 8
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xb6\xb2"...,
						)
					case 7:
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xa9\xb4"...,
						)
						i = 2
					case 3:
						i = 1
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xbay"...,
						)
					case 4:
						i = 11
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xd5\x7f\xc1\xc0"...,
						)
					case 0:
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xb1\xb2\xb1\xb0"...,
						)
						i = 3
					case 10:
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "°`\xa6"...,
						)
						i = 0
					case 11:
						data = /*line OpTkDPbc4LGC.go:1*/ append(data, "\xb4\xb8\xc5"...,
						)
						i = 9
					}
				}
				return /*line OpTkDPbc4LGC.go:1*/ string(data)
			}(), "error", cYHA75qVq)
		}
	}

	var stqEk4Brz1wF time.Time
	if metrics.EnabledExpensive {
		stqEk4Brz1wF = /*line AXqhP8zga4o.go:1*/ time.Now()
	}
	lxItfhw, ifeYDU90, cYHA75qVq := /*line araUGMR.go:1*/ peYZBpNwox.trie.Commit(true)
	if cYHA75qVq != nil {
		return common.Hash{}, cYHA75qVq
	}

	if ifeYDU90 != nil {
		if cYHA75qVq := /*line O08hdhQT9nkv.go:1*/ nQXC7Td.Merge(ifeYDU90); cYHA75qVq != nil {
			return common.Hash{}, cYHA75qVq
		}
		ayLK4KQXyY, fGenfN868u5 = /*line Bb0DrRQO5ejK.go:1*/ ifeYDU90.Size()
	}
	if metrics.EnabledExpensive {
		peYZBpNwox.AccountCommits += /*line mO4uTqaqMP.go:1*/ time.Since(stqEk4Brz1wF)

		/*line u_iM9ag.go:1*/
		t14aLF4Y.Mark( /*line zj2SyBPvV0.go:1*/ int64(peYZBpNwox.AccountUpdated))
		/*line uSqZV5x.go:1*/ jWXmaMS3d.Mark( /*line uEOvOYzg2TRx.go:1*/ int64(peYZBpNwox.StorageUpdated))
		/*line UbENBGE.go:1*/ qqNAVrhZ.Mark( /*line mr96dpfp.go:1*/ int64(peYZBpNwox.AccountDeleted))
		/*line yDtS0o6i6f.go:1*/ ewSUzVXKn.Mark( /*line WJJVKbD.go:1*/ int64(peYZBpNwox.StorageDeleted))
		/*line pR04eM17FVi.go:1*/ xNTiw5VKek.Mark( /*line QNwEdQ_tU3T.go:1*/ int64(ayLK4KQXyY))
		/*line DeVG7tJAQqr.go:1*/ db1BShRmyWbt.Mark( /*line CDFIGdj96.go:1*/ int64(fGenfN868u5))
		/*line efnCWId8I6tf.go:1*/ tRW4TBhii.Mark( /*line Tt3tv1yda1lk.go:1*/ int64(xT719aaps))
		/*line KsEOGaKy.go:1*/ qiwd7Qf6LoOn.Mark( /*line JL8sjskTDl6o.go:1*/ int64(cjQJXhTj))
		peYZBpNwox.AccountUpdated, peYZBpNwox.AccountDeleted = 0, 0
		peYZBpNwox.StorageUpdated, peYZBpNwox.StorageDeleted = 0, 0
	}

	if peYZBpNwox.snap != nil {
		stqEk4Brz1wF := /*line GfSMrN_DRWXf.go:1*/ time.Now()

		if orvSpr3dK := /*line lbT_1ByaK.go:1*/ peYZBpNwox.snap.Root(); orvSpr3dK != lxItfhw {
			if cYHA75qVq := /*line b_sZ2uX.go:1*/ peYZBpNwox.snaps.Update(lxItfhw, orvSpr3dK /*line CPBCGjan.go:1*/, peYZBpNwox.tQsAgQ71(peYZBpNwox.stateObjectsDestruct), peYZBpNwox.accounts, peYZBpNwox.storages); cYHA75qVq != nil {
				/*line wSuNE1.go:1*/ log.Warn(func() /*line N8H9fKwDOZ3l.go:1*/ string {
					seed := /*line N8H9fKwDOZ3l.go:1*/ byte(24)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line N8H9fKwDOZ3l.go:1*/ append(data, x+seed); seed += x; return fnc }
					/*line N8H9fKwDOZ3l.go:1*/ fnc(46)(27)(8)(3)(249)(255)(188)(84)(251)(177)(85)(251)(244)(253)(19)(241)(187)(83)(251)(243)(15)(3)(245)(7)(5)(172)(84)(254)(243)(0)
					return /*line N8H9fKwDOZ3l.go:1*/ string(data)
				}(), "from", orvSpr3dK, "to", lxItfhw, "err", cYHA75qVq)
			}

			if cYHA75qVq := /*line LCSmU8.go:1*/ peYZBpNwox.snaps.Cap(lxItfhw, 128); cYHA75qVq != nil {
				/*line SfauL_rvO6H.go:1*/ log.Warn(func() /*line vDvPJCw.go:1*/ string {
					key := [] /*line vDvPJCw.go:1*/ byte("\xd1\x19$\xb1\xe8\xd3\xc9{\x87\xcfS\xfb\x95\x0f\"\xb6\xb4\x93\x03\x06\xf3\x9aR\xb1E'2")
					data := [] /*line vDvPJCw.go:1*/ byte("\x97xMݍ\xb7\xe9\x0f\xe8\xef0\x9a\xe5/Q\xd8\xd5\xe3pn\x9c\xeer\xc57BW")
					for i, b := range key {
						data[i] = data[i] ^ b
					}
					return /*line vDvPJCw.go:1*/ string(data)
				}(), "root", lxItfhw, "layers", 128, "err", cYHA75qVq)
			}
		}
		if metrics.EnabledExpensive {
			peYZBpNwox.SnapshotCommits += /*line CPkuFqtcy.go:1*/ time.Since(stqEk4Brz1wF)
		}
		peYZBpNwox.snap = nil
	}
	if lxItfhw == (common.Hash{}) {
		lxItfhw = types.NrGuaNA21
	}
	deml9qsrR5FA := peYZBpNwox.originalRoot
	if deml9qsrR5FA == (common.Hash{}) {
		deml9qsrR5FA = types.NrGuaNA21
	}
	if lxItfhw != deml9qsrR5FA {
		stqEk4Brz1wF := /*line zT0oDspuuuB.go:1*/ time.Now()
		ifeYDU90 := /*line JXjNE9RtxX.go:1*/ triestate.New(peYZBpNwox.accountsOrigin, peYZBpNwox.storagesOrigin, avY0t4RP1I)
		if cYHA75qVq := /*line cPa9BcW6O.go:1*/ peYZBpNwox.db.TrieDB().Update(lxItfhw, deml9qsrR5FA, cCPWpPefu, nQXC7Td, ifeYDU90); cYHA75qVq != nil {
			return common.Hash{}, cYHA75qVq
		}
		peYZBpNwox.originalRoot = lxItfhw
		if metrics.EnabledExpensive {
			peYZBpNwox.TrieDBCommits += /*line EqjwuKuhT.go:1*/ time.Since(stqEk4Brz1wF)
		}
		if peYZBpNwox.onCommit != nil {
			/*line whjaJdwy0fSC.go:1*/ peYZBpNwox.onCommit(ifeYDU90)
		}
	}

	peYZBpNwox.accounts = /*line a09G9QW.go:1*/ make(map[common.Hash][]byte)
	peYZBpNwox.storages = /*line ESnJz0A1PgdK.go:1*/ make(map[common.Hash]map[common.Hash][]byte)
	peYZBpNwox.accountsOrigin = /*line qKpawIMftKwY.go:1*/ make(map[common.Address][]byte)
	peYZBpNwox.storagesOrigin = /*line ONAHFykUt.go:1*/ make(map[common.Address]map[common.Hash][]byte)
	peYZBpNwox.stateObjectsDirty = /*line qtwZO2Wv.go:1*/ make(map[common.Address]struct{})
	peYZBpNwox.stateObjectsDestruct = /*line V_uYpTBXGJCq.go:1*/ make(map[common.Address]*types.StateAccount)
	return lxItfhw, nil
}

func (peYZBpNwox *StateDB) Prepare(t6NRfk6xn0 params.Rules, f7dNjq8iK6I, aV8ZhY6R8H4 common.Address, bVZjFaT37U7 *common.Address, vyaO1EcASWyu []common.Address, f7dIdwRLoax types.AccessList) {
	if t6NRfk6xn0.IsBerlin {

		kFF49KO := /*line fWZbEs.go:1*/ viqfAs6EZc()
		peYZBpNwox.accessList = kFF49KO

		/*line ccAptdksEqo.go:1*/
		kFF49KO.AddAddress(f7dNjq8iK6I)
		if bVZjFaT37U7 != nil {
			/*line bE_Lwa.go:1*/ kFF49KO.AddAddress(*bVZjFaT37U7)

		}
		for _, hv1UkMRaKf := range vyaO1EcASWyu {
			/*line yjSL86xUJ.go:1*/ kFF49KO.AddAddress(hv1UkMRaKf)
		}
		for _, acM7_8MbU4 := range f7dIdwRLoax {
			/*line EuawQDkIqZa.go:1*/ kFF49KO.AddAddress(acM7_8MbU4.Address)
			for _, ixeygbAgOap := range acM7_8MbU4.StorageKeys {
				/*line j6dO0w6vJ3.go:1*/ kFF49KO.AddSlot(acM7_8MbU4.Address, ixeygbAgOap)
			}
		}
		if t6NRfk6xn0.IsShanghai {
			/*line JIL4BFFaoEB.go:1*/ kFF49KO.AddAddress(aV8ZhY6R8H4)
		}
	}

	peYZBpNwox.transientStorage = /*line PPYm2GiGh.go:1*/ s9XSqaz()
}

func (peYZBpNwox *StateDB) AddAddressToAccessList(hv1UkMRaKf common.Address) {
	if /*line bqUbMx_.go:1*/ peYZBpNwox.accessList.AddAddress(hv1UkMRaKf) {
		/*line vgdgoOmnHe.go:1*/ peYZBpNwox.journal.uwXnrqq(vPXduwMcrq{&hv1UkMRaKf})
	}
}

func (peYZBpNwox *StateDB) AddSlotToAccessList(hv1UkMRaKf common.Address, ibmfxwhUDRC common.Hash) {
	ajaHPOARCE, bBdaEE := /*line aRIwS0aTH.go:1*/ peYZBpNwox.accessList.AddSlot(hv1UkMRaKf, ibmfxwhUDRC)
	if ajaHPOARCE {

		/*line JE6ZwqtZ.go:1*/
		peYZBpNwox.journal.uwXnrqq(vPXduwMcrq{&hv1UkMRaKf})
	}
	if bBdaEE {
		/*line a96waZ5WKf.go:1*/ peYZBpNwox.journal.uwXnrqq(eboE8NSCCQU{
			p7jEjRwYLtL: &hv1UkMRaKf,
			ythV5glxZGp: &ibmfxwhUDRC,
		})
	}
}

func (peYZBpNwox *StateDB) AddressInAccessList(hv1UkMRaKf common.Address) bool {
	return /*line wSHIZ4gjDT.go:1*/ peYZBpNwox.accessList.ContainsAddress(hv1UkMRaKf)
}

func (peYZBpNwox *StateDB) SlotInAccessList(hv1UkMRaKf common.Address, ibmfxwhUDRC common.Hash) (fsAQ1g bool, nZAaaiHQ9 bool) {
	return /*line af2S0YRcL.go:1*/ peYZBpNwox.accessList.Contains(hv1UkMRaKf, ibmfxwhUDRC)
}

func (peYZBpNwox *StateDB) tQsAgQ71(ifeYDU90 map[common.Address]*types.StateAccount) map[common.Hash]struct{} {
	vImTncmiqQaT := /*line iCMZafx2Cf9I.go:1*/ make(map[common.Hash]struct{} /*line yKhmuFx.go:1*/, len(ifeYDU90))
	for hv1UkMRaKf := range ifeYDU90 {
		ws3tNR3MU, _bRYnhxrC := peYZBpNwox.stateObjects[hv1UkMRaKf]
		if !_bRYnhxrC {
			vImTncmiqQaT[ /*line EAiVw0SZ1f0.go:1*/ crypto.Keccak256Hash(hv1UkMRaKf[:])] = struct{}{}
		} else {
			vImTncmiqQaT[ws3tNR3MU.addrHash] = struct{}{}
		}
	}
	return vImTncmiqQaT
}

func eh1lN9hBaUMv[_3JfOPUaHxI comparable](ifeYDU90 map[_3JfOPUaHxI][]byte) map[_3JfOPUaHxI][]byte {
	_079lbTfgS := /*line FKZTdVH.go:1*/ make(map[_3JfOPUaHxI][]byte /*line IIOIEFySn6.go:1*/, len(ifeYDU90))
	for ixeygbAgOap, cGzt0N := range ifeYDU90 {
		_079lbTfgS[ixeygbAgOap] = /*line HQt1ucm3DJj.go:1*/ common.CopyBytes(cGzt0N)
	}
	return _079lbTfgS
}

func ngMfLQ4pubvm[_3JfOPUaHxI comparable](ifeYDU90 map[_3JfOPUaHxI]map[common.Hash][]byte) map[_3JfOPUaHxI]map[common.Hash][]byte {
	_079lbTfgS := /*line L7Ipfset1.go:1*/ make(map[_3JfOPUaHxI]map[common.Hash][]byte /*line CTMt2I.go:1*/, len(ifeYDU90))
	for hv1UkMRaKf, jIhdoI := range ifeYDU90 {
		_079lbTfgS[hv1UkMRaKf] = /*line BpRzQwzdUan.go:1*/ make(map[common.Hash][]byte /*line erLej9RQns.go:1*/, len(jIhdoI))
		for ixeygbAgOap, cGzt0N := range jIhdoI {
			_079lbTfgS[hv1UkMRaKf][ixeygbAgOap] = /*line SZcEl46Na.go:1*/ common.CopyBytes(cGzt0N)
		}
	}
	return _079lbTfgS
}

var _ = fmt.Append
var _ = sort.Find

const _ = time.ANSIC

var _ snapshot.Nq4YsH_
var _ trie.SXoLHxUt177q
var _ types.AccessList
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ = crypto.CompressPubkey
var _ = log.Crit
var _ = metrics.AccountingRegistry
var _ = params.AllCliqueProtocolChanges
var _ trienode.MergedNodeSet
var _ = triestate.Apply
var _ = uint256.ErrBadBufferLength
