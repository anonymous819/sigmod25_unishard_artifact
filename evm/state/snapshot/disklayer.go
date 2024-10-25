//line :1
package snapshot

import (
	"bytes"
	"sync"

	types "unishard/evm/types"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/triedb"
)

type diskLayer struct {
	diskdb ethdb.KeyValueStore
	triedb *triedb.Database
	cache  *fastcache.Cache

	root  common.Hash
	stale bool

	genMarker  []byte
	genPending chan struct{}
	genAbort   chan chan *generatorStats

	lock sync.RWMutex
}

func (hbsnnIi *diskLayer) Release() error {
	if hbsnnIi.cache != nil {
		/*line jFYYXdUl43rs.go:1*/ hbsnnIi.cache.Reset()
	}
	return nil
}

func (hbsnnIi *diskLayer) Root() common.Hash {
	return hbsnnIi.root
}

func (hbsnnIi *diskLayer) Parent() snapshot {
	return nil
}

func (hbsnnIi *diskLayer) Stale() bool {
	/*line _Pj4xYmQa.go:1*/ hbsnnIi.lock.RLock()
	defer /*line sbL55CpD_2q.go:1*/ hbsnnIi.lock.RUnlock()

	return hbsnnIi.stale
}

func (hbsnnIi *diskLayer) Account(xhOzkRrKDZ common.Hash) (*types.SlimAccount, error) {
	bI1ciqVUvL4b, chyZ8Q := /*line A7gB0aW3e.go:1*/ hbsnnIi.AccountRLP(xhOzkRrKDZ)
	if chyZ8Q != nil {
		return nil, chyZ8Q
	}
	if /*line vzaEGPqBD5.go:1*/ len(bI1ciqVUvL4b) == 0 {
		return nil, nil
	}
	evzmhL1 := /*line IXX9Hg7wiqFp.go:1*/ new(types.SlimAccount)
	if chyZ8Q := /*line DcJfLAk8p.go:1*/ rlp.DecodeBytes(bI1ciqVUvL4b, evzmhL1); chyZ8Q != nil {
		/*line ncNQcyAA75j5.go:1*/ panic(chyZ8Q)
	}
	return evzmhL1, nil
}

func (hbsnnIi *diskLayer) AccountRLP(xhOzkRrKDZ common.Hash) ([]byte, error) {
	/*line IdpyaR.go:1*/ hbsnnIi.lock.RLock()
	defer /*line Cw5bmXY.go:1*/ hbsnnIi.lock.RUnlock()

	if hbsnnIi.stale {
		return nil, AiK5hkDaW
	}

	if hbsnnIi.genMarker != nil && /*line IYT_RRYzr.go:1*/ bytes.Compare(xhOzkRrKDZ[:], hbsnnIi.genMarker) > 0 {
		return nil, IK452Y2zeAdC
	}

	/*line Afo7zZ.go:1*/
	xYuyEVnD9.Mark(1)

	if b1jafJL, qpHiF3kG3W := /*line cTRhfzd2.go:1*/ hbsnnIi.cache.HasGet(nil, xhOzkRrKDZ[:]); qpHiF3kG3W {
		/*line KVW1JRMMp.go:1*/ vA2HcXHs0F.Mark(1)
		/*line yS78trT258T.go:1*/ nQLCSIOQS.Mark( /*line fIhkDSwC_E.go:1*/ int64( /*line haMLcgtl8C.go:1*/ len(b1jafJL)))
		return b1jafJL, nil
	}

	b1jafJL := /*line BJL0JQPCzgq.go:1*/ rawdb.ReadAccountSnapshot(hbsnnIi.diskdb, xhOzkRrKDZ)
	/*line GD9_2tz1.go:1*/ hbsnnIi.cache.Set(xhOzkRrKDZ[:], b1jafJL)

	/*line Ky_xd1RqHz2.go:1*/
	xL3GIW.Mark(1)
	if cZL9UJradA7 := /*line sRAsAe.go:1*/ len(b1jafJL); cZL9UJradA7 > 0 {
		/*line HuSYUqxJO6.go:1*/ t9NCla4Ghrl.Mark( /*line Ytwgrs.go:1*/ int64(cZL9UJradA7))
	} else {
		/*line hcnTea.go:1*/ msaMf9.Mark(1)
	}
	return b1jafJL, nil
}

func (hbsnnIi *diskLayer) Storage(dcJdHV, mq_bNE9GL common.Hash) ([]byte, error) {
	/*line BwRO8BKz.go:1*/ hbsnnIi.lock.RLock()
	defer /*line cap92SDvo.go:1*/ hbsnnIi.lock.RUnlock()

	if hbsnnIi.stale {
		return nil, AiK5hkDaW
	}
	nVUwQz8Zi := /*line gpk7QknKlQ.go:1*/ append(dcJdHV[:], mq_bNE9GL[:]...)

	if hbsnnIi.genMarker != nil && /*line fLIASMTTa.go:1*/ bytes.Compare(nVUwQz8Zi, hbsnnIi.genMarker) > 0 {
		return nil, IK452Y2zeAdC
	}

	/*line BF4VQTFxB.go:1*/
	mt6xkEo.Mark(1)

	if b1jafJL, qpHiF3kG3W := /*line DpQLQj6Bal.go:1*/ hbsnnIi.cache.HasGet(nil, nVUwQz8Zi); qpHiF3kG3W {
		/*line BbQF1zTWs.go:1*/ vG6wMf.Mark(1)
		/*line eC6XJIS9FgM.go:1*/ jJdpjkAbM4.Mark( /*line aJ88oo.go:1*/ int64( /*line gREQfhVWJ6jB.go:1*/ len(b1jafJL)))
		return b1jafJL, nil
	}

	b1jafJL := /*line nHEFUq6dzw.go:1*/ rawdb.ReadStorageSnapshot(hbsnnIi.diskdb, dcJdHV, mq_bNE9GL)
	/*line R1rJf68d.go:1*/ hbsnnIi.cache.Set(nVUwQz8Zi, b1jafJL)

	/*line X3GyGQd32.go:1*/
	jzc5tPpFnz.Mark(1)
	if cZL9UJradA7 := /*line vZb6gfM.go:1*/ len(b1jafJL); cZL9UJradA7 > 0 {
		/*line l6UvgZL_1mP.go:1*/ wZpZiLVo4M1.Mark( /*line OmiBscNHcRk.go:1*/ int64(cZL9UJradA7))
	} else {
		/*line E8UcxNuG4ceA.go:1*/ hqqacN.Mark(1)
	}
	return b1jafJL, nil
}

func (hbsnnIi *diskLayer) Update(sBF5LXy5I8g common.Hash, zoMmGkL map[common.Hash]struct{}, bQIIyfhppAL1 map[common.Hash][]byte, agSpCMzc map[common.Hash]map[common.Hash][]byte) *diffLayer {
	return /*line BQrOkR.go:1*/ aa7YtGB_s(hbsnnIi, sBF5LXy5I8g, zoMmGkL, bQIIyfhppAL1, agSpCMzc)
}

var _ bytes.Buffer
var _ sync.Cond
var _ types.AccessList
var _ fastcache.BigStats
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ = rlp.AppendUint64
var _ triedb.Config
