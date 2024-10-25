//line :1
package state

import (
	"bytes"
	"fmt"
	"io"
	"time"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/holiman/uint256"
)

type Code []byte

func (ts1y5Po Code) String() string {
	return /*line PsnVVAqFB.go:1*/ string(ts1y5Po)
}

type Storage map[common.Hash]common.Hash

func (peYZBpNwox Storage) String() (vUllZ82O string) {
	for ixeygbAgOap, fnngafl := range peYZBpNwox {
		vUllZ82O += /*line Dh1Nzyjege.go:1*/ fmt.Sprintf(func() /*line AvvkBOI2cd.go:1*/ string {
			seed := /*line AvvkBOI2cd.go:1*/ byte(48)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line AvvkBOI2cd.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line AvvkBOI2cd.go:1*/ fnc(245)(51)(200)(26)(230)(5)(51)(178)
			return /*line AvvkBOI2cd.go:1*/ string(data)
		}(), ixeygbAgOap, fnngafl)
	}
	return
}

func (peYZBpNwox Storage) Copy() Storage {
	mMvShC5AUZkj := /*line EgPzqVp.go:1*/ make(Storage /*line taj0o072fO.go:1*/, len(peYZBpNwox))
	for ixeygbAgOap, fnngafl := range peYZBpNwox {
		mMvShC5AUZkj[ixeygbAgOap] = fnngafl
	}
	return mMvShC5AUZkj
}

type stateObject struct {
	db       *StateDB
	address  common.Address
	addrHash common.Hash
	origin   *types.StateAccount
	data     types.StateAccount

	trie       Trie
	code       Code
	deploycode Code

	originStorage  Storage
	pendingStorage Storage
	dirtyStorage   Storage

	dirtyCode bool

	selfDestructed bool

	deleted bool

	created bool
}

func (peYZBpNwox *stateObject) usx3KByCvUvp() bool {
	return peYZBpNwox.data.Nonce == 0 && /*line AOnM7Ra1f.go:1*/ peYZBpNwox.data.Balance.IsZero() && /*line aPJXgxE6q5.go:1*/ bytes.Equal(peYZBpNwox.data.CodeHash /*line XasButAUm.go:1*/, types.JhrQnETMFm.Bytes())
}

func zkEv6T(tC4SGV *StateDB, so0TaNzLDdpc common.Address, gnXFNzo44 *types.StateAccount) *stateObject {
	var (
		deml9qsrR5FA = gnXFNzo44
		l9wReS6      = gnXFNzo44 == nil
	)
	if gnXFNzo44 == nil {
		gnXFNzo44 = /*line Etye6eak5A.go:1*/ types.D1o8Dbhahg()
	}
	return &stateObject{
		db:      tC4SGV,
		address: so0TaNzLDdpc,
		addrHash:/*line BMikiOMX.go:1*/ crypto.Keccak256Hash(so0TaNzLDdpc[:]),
		origin: deml9qsrR5FA,
		data:   *gnXFNzo44,
		originStorage:/*line V8adGuZcSf.go:1*/ make(Storage),
		pendingStorage:/*line _7fLlz.go:1*/ make(Storage),
		dirtyStorage:/*line Qz7W44a.go:1*/ make(Storage),
		created: l9wReS6,
	}
}

func (peYZBpNwox *stateObject) EncodeRLP(dPVWOjCSXmL io.Writer) error {
	return /*line OhTsp1Qv8Op.go:1*/ rlp.Encode(dPVWOjCSXmL, &peYZBpNwox.data)
}

func (peYZBpNwox *stateObject) eUznOx() {
	peYZBpNwox.selfDestructed = true
}

func (peYZBpNwox *stateObject) b2rTB7zaE9() {
	/*line yEZA2N0RtUa.go:1*/ peYZBpNwox.db.journal.uwXnrqq(fbLbWCnCjB81{
		huIthfk25W: &peYZBpNwox.address,
	})
	if peYZBpNwox.address == m8ZbLkxbBc {

		/*line myK5cukRzw_.go:1*/
		peYZBpNwox.db.journal.vdDLIhK(peYZBpNwox.address)
	}
}

func (peYZBpNwox *stateObject) z1F2PO() (Trie, error) {
	if peYZBpNwox.trie == nil {

		if peYZBpNwox.data.Root != types.NrGuaNA21 && peYZBpNwox.db.prefetcher != nil {

			peYZBpNwox.trie = /*line meDvwO6.go:1*/ peYZBpNwox.db.prefetcher.fQodR9c7t(peYZBpNwox.addrHash, peYZBpNwox.data.Root)
		}
		if peYZBpNwox.trie == nil {
			r8CzlEu, cYHA75qVq := /*line ZmMeajWhBj.go:1*/ peYZBpNwox.db.db.OpenStorageTrie(peYZBpNwox.db.originalRoot, peYZBpNwox.address, peYZBpNwox.data.Root, peYZBpNwox.db.trie)
			if cYHA75qVq != nil {
				return nil, cYHA75qVq
			}
			peYZBpNwox.trie = r8CzlEu
		}
	}
	return peYZBpNwox.trie, nil
}

func (peYZBpNwox *stateObject) GetState(ixeygbAgOap common.Hash) common.Hash {

	fnngafl, vdDLIhK := peYZBpNwox.dirtyStorage[ixeygbAgOap]
	if vdDLIhK {
		return fnngafl
	}

	return /*line pe45aTuFL.go:1*/ peYZBpNwox.GetCommittedState(ixeygbAgOap)
}

func (peYZBpNwox *stateObject) GetCommittedState(ixeygbAgOap common.Hash) common.Hash {

	if fnngafl, tKvj8dj8t := peYZBpNwox.pendingStorage[ixeygbAgOap]; tKvj8dj8t {
		return fnngafl
	}
	if fnngafl, rurYlaMG := peYZBpNwox.originStorage[ixeygbAgOap]; rurYlaMG {
		return fnngafl
	}

	if _, aSTjj8aWU9gM := peYZBpNwox.db.stateObjectsDestruct[peYZBpNwox.address]; aSTjj8aWU9gM {
		return common.Hash{}
	}

	var (
		_Ly71jodD []byte
		cYHA75qVq error
		fnngafl   common.Hash
	)
	if peYZBpNwox.db.snap != nil {
		stqEk4Brz1wF := /*line RZ2nj8z.go:1*/ time.Now()
		_Ly71jodD, cYHA75qVq = /*line eDJlsRjovHMB.go:1*/ peYZBpNwox.db.snap.Storage(peYZBpNwox.addrHash /*line jkSbgko.go:1*/, crypto.Keccak256Hash( /*line AM08p0c.go:1*/ ixeygbAgOap.Bytes()))
		if metrics.EnabledExpensive {
			peYZBpNwox.db.SnapshotStorageReads += /*line vTzD4PUPPh.go:1*/ time.Since(stqEk4Brz1wF)
		}
		if /*line Dbamph9.go:1*/ len(_Ly71jodD) > 0 {
			_, i9zJqH1SGp, _, cYHA75qVq := /*line Vbo_q5T.go:1*/ rlp.Split(_Ly71jodD)
			if cYHA75qVq != nil {
				/*line UAq37pu6hk.go:1*/ peYZBpNwox.db.eMZRw9Dc(cYHA75qVq)
			}
			/*line uuImT4YXah.go:1*/ fnngafl.SetBytes(i9zJqH1SGp)
		}
	}

	if peYZBpNwox.db.snap == nil || cYHA75qVq != nil {
		stqEk4Brz1wF := /*line CCGeIrKyV.go:1*/ time.Now()
		r8CzlEu, cYHA75qVq := /*line kxdhtTxDRo.go:1*/ peYZBpNwox.z1F2PO()
		if cYHA75qVq != nil {
			/*line ioCpQ1BbGe.go:1*/ peYZBpNwox.db.eMZRw9Dc(cYHA75qVq)
			return common.Hash{}
		}
		cGzt0N, cYHA75qVq := /*line e7DIgR2.go:1*/ r8CzlEu.GetStorage(peYZBpNwox.address /*line t0YMMpiI.go:1*/, ixeygbAgOap.Bytes())
		if metrics.EnabledExpensive {
			peYZBpNwox.db.StorageReads += /*line NzzHjpIPcX.go:1*/ time.Since(stqEk4Brz1wF)
		}
		if cYHA75qVq != nil {
			/*line u6z8siUPB.go:1*/ peYZBpNwox.db.eMZRw9Dc(cYHA75qVq)
			return common.Hash{}
		}
		/*line m6KA4f.go:1*/ fnngafl.SetBytes(cGzt0N)
	}
	peYZBpNwox.originStorage[ixeygbAgOap] = fnngafl
	return fnngafl
}

func (peYZBpNwox *stateObject) SetState(ixeygbAgOap, fnngafl common.Hash) {

	reikfgaLNK := /*line zmXfFD.go:1*/ peYZBpNwox.GetState(ixeygbAgOap)
	if reikfgaLNK == fnngafl {
		return
	}

	/*line aqhZOP.go:1*/
	peYZBpNwox.db.journal.uwXnrqq(nUsbOK{
		iLbaDHlKPp:  &peYZBpNwox.address,
		r9VPURO:     ixeygbAgOap,
		eT8enfTtJwO: reikfgaLNK,
	})
	/*line bzSuw0XdEBVf.go:1*/ peYZBpNwox.kInfXnsga(ixeygbAgOap, fnngafl)
}

func (peYZBpNwox *stateObject) kInfXnsga(ixeygbAgOap, fnngafl common.Hash) {
	peYZBpNwox.dirtyStorage[ixeygbAgOap] = fnngafl
}

func (peYZBpNwox *stateObject) tatrqu(i3uM6DE bool) {
	y3iHCyWhfb9l := /*line XnewWljgTR_7.go:1*/ make([][]byte, 0 /*line V2m6V5ahK.go:1*/, len(peYZBpNwox.dirtyStorage))
	for ixeygbAgOap, fnngafl := range peYZBpNwox.dirtyStorage {
		peYZBpNwox.pendingStorage[ixeygbAgOap] = fnngafl
		if fnngafl != peYZBpNwox.originStorage[ixeygbAgOap] {
			y3iHCyWhfb9l = /*line eaQrnqr9g_.go:1*/ append(y3iHCyWhfb9l /*line fdYoRq.go:1*/, common.CopyBytes(ixeygbAgOap[:]))
		}
	}
	if peYZBpNwox.db.prefetcher != nil && i3uM6DE && /*line mN5gJnr8.go:1*/ len(y3iHCyWhfb9l) > 0 && peYZBpNwox.data.Root != types.NrGuaNA21 {
		/*line G6QzmSa.go:1*/ peYZBpNwox.db.prefetcher.i3uM6DE(peYZBpNwox.addrHash, peYZBpNwox.data.Root, peYZBpNwox.address, y3iHCyWhfb9l)
	}
	if /*line xZqkS9Fm.go:1*/ len(peYZBpNwox.dirtyStorage) > 0 {
		peYZBpNwox.dirtyStorage = /*line ayn8OEbs.go:1*/ make(Storage)
	}
}

func (peYZBpNwox *stateObject) zvFXqsueMRod() (Trie, error) {

	/*line PGXFpp8zBF.go:1*/
	peYZBpNwox.tatrqu(false)

	if /*line vpv5tIgs8.go:1*/ len(peYZBpNwox.pendingStorage) == 0 {
		return peYZBpNwox.trie, nil
	}

	if metrics.EnabledExpensive {
		defer func( /*line yScQAFmx5aSk.go:1*/ stqEk4Brz1wF time.Time) {
			peYZBpNwox.db.StorageUpdates += /*line h84voMZk.go:1*/ time.Since(stqEk4Brz1wF)
		}( /*line PRY5qfzJo.go:1*/ time.Now())
	}

	var (
		vg59BBBGWI   map[common.Hash][]byte
		deml9qsrR5FA map[common.Hash][]byte
	)
	r8CzlEu, cYHA75qVq := /*line jKmqvA_KW.go:1*/ peYZBpNwox.z1F2PO()
	if cYHA75qVq != nil {
		/*line fAUAMWc.go:1*/ peYZBpNwox.db.eMZRw9Dc(cYHA75qVq)
		return nil, cYHA75qVq
	}

	utP5tDSa3OnI := /*line A6w3D0i6bDVF.go:1*/ make([][]byte, 0 /*line J87lgY.go:1*/, len(peYZBpNwox.pendingStorage))
	for ixeygbAgOap, fnngafl := range peYZBpNwox.pendingStorage {

		if fnngafl == peYZBpNwox.originStorage[ixeygbAgOap] {
			continue
		}
		reikfgaLNK := peYZBpNwox.originStorage[ixeygbAgOap]
		peYZBpNwox.originStorage[ixeygbAgOap] = fnngafl

		var tJL2yDNwd0 []byte
		if (fnngafl == common.Hash{}) {
			if cYHA75qVq := /*line BAg_vRg22.go:1*/ r8CzlEu.DeleteStorage(peYZBpNwox.address, ixeygbAgOap[:]); cYHA75qVq != nil {
				/*line GFbv96ftd.go:1*/ peYZBpNwox.db.eMZRw9Dc(cYHA75qVq)
				return nil, cYHA75qVq
			}
			peYZBpNwox.db.StorageDeleted += 1
		} else {

			ijNibE := /*line s6s_AKrq1.go:1*/ common.TrimLeftZeroes(fnngafl[:])
			tJL2yDNwd0, _ = /*line jZqy0iL8.go:1*/ rlp.EncodeToBytes(ijNibE)
			if cYHA75qVq := /*line U_szsPSogr.go:1*/ r8CzlEu.UpdateStorage(peYZBpNwox.address, ixeygbAgOap[:], ijNibE); cYHA75qVq != nil {
				/*line kFxTFIh9.go:1*/ peYZBpNwox.db.eMZRw9Dc(cYHA75qVq)
				return nil, cYHA75qVq
			}
			peYZBpNwox.db.StorageUpdated += 1
		}

		if vg59BBBGWI == nil {
			if vg59BBBGWI = peYZBpNwox.db.storages[peYZBpNwox.addrHash]; vg59BBBGWI == nil {
				vg59BBBGWI = /*line rrSz7o9IpV.go:1*/ make(map[common.Hash][]byte)
				peYZBpNwox.db.storages[peYZBpNwox.addrHash] = vg59BBBGWI
			}
		}
		j3syD_8Sgr := /*line t3fuSsd.go:1*/ crypto.HashData(peYZBpNwox.db.hasher, ixeygbAgOap[:])
		vg59BBBGWI[j3syD_8Sgr] = tJL2yDNwd0

		if deml9qsrR5FA == nil {
			if deml9qsrR5FA = peYZBpNwox.db.storagesOrigin[peYZBpNwox.address]; deml9qsrR5FA == nil {
				deml9qsrR5FA = /*line hvBQDw4aeyP7.go:1*/ make(map[common.Hash][]byte)
				peYZBpNwox.db.storagesOrigin[peYZBpNwox.address] = deml9qsrR5FA
			}
		}

		if _, dNazL3Oz4 := deml9qsrR5FA[j3syD_8Sgr]; !dNazL3Oz4 {
			if reikfgaLNK == (common.Hash{}) {
				deml9qsrR5FA[j3syD_8Sgr] = nil
			} else {

				ftVrYa2, _ := /*line Mlsswjt3RTlR.go:1*/ rlp.EncodeToBytes( /*line PNDuEYIy_m.go:1*/ common.TrimLeftZeroes(reikfgaLNK[:]))
				deml9qsrR5FA[j3syD_8Sgr] = ftVrYa2
			}
		}

		utP5tDSa3OnI = /*line CZYg1p56TNH.go:1*/ append(utP5tDSa3OnI /*line RU4VqVwF.go:1*/, common.CopyBytes(ixeygbAgOap[:]))
	}
	if peYZBpNwox.db.prefetcher != nil {
		/*line AFq73iLh85.go:1*/ peYZBpNwox.db.prefetcher.t50cl4u8e(peYZBpNwox.addrHash, peYZBpNwox.data.Root, utP5tDSa3OnI)
	}
	peYZBpNwox.pendingStorage = /*line _Su2XEkn.go:1*/ make(Storage)
	return r8CzlEu, nil
}

func (peYZBpNwox *stateObject) hbhGaAPh6VSz() {

	r8CzlEu, cYHA75qVq := /*line GaQSaqA.go:1*/ peYZBpNwox.zvFXqsueMRod()
	if cYHA75qVq != nil || r8CzlEu == nil {
		return
	}

	if metrics.EnabledExpensive {
		defer func( /*line u9vu2BOHc.go:1*/ stqEk4Brz1wF time.Time) {
			peYZBpNwox.db.StorageHashes += /*line J5H5KYD6iP.go:1*/ time.Since(stqEk4Brz1wF)
		}( /*line JX4TVX.go:1*/ time.Now())
	}
	peYZBpNwox.data.Root = /*line Hc8KGk.go:1*/ r8CzlEu.Hash()
}

func (peYZBpNwox *stateObject) xGnpP8a5() (*trienode.NodeSet, error) {

	if peYZBpNwox.trie == nil {
		peYZBpNwox.origin = /*line mew2tmsRJra.go:1*/ peYZBpNwox.data.Copy()
		return nil, nil
	}

	if metrics.EnabledExpensive {
		defer func( /*line GwCTLrsIa5v.go:1*/ stqEk4Brz1wF time.Time) {
			peYZBpNwox.db.StorageCommits += /*line BgZjTN04BW.go:1*/ time.Since(stqEk4Brz1wF)
		}( /*line Q_pbpJ.go:1*/ time.Now())
	}

	lxItfhw, nQXC7Td, cYHA75qVq := /*line TDLE3b.go:1*/ peYZBpNwox.trie.Commit(false)
	if cYHA75qVq != nil {
		return nil, cYHA75qVq
	}
	peYZBpNwox.data.Root = lxItfhw

	peYZBpNwox.origin = /*line h3MLD5z7.go:1*/ peYZBpNwox.data.Copy()
	return nQXC7Td, nil
}

func (peYZBpNwox *stateObject) AddBalance(zGNz6Q9KhWC *uint256.Int) {

	if /*line BCleDLvlGUc.go:1*/ zGNz6Q9KhWC.IsZero() {
		if /*line aupqrat.go:1*/ peYZBpNwox.usx3KByCvUvp() {
			/*line qSGfUW.go:1*/ peYZBpNwox.b2rTB7zaE9()
		}
		return
	}
	/*line hAnbeY5HDoR.go:1*/ peYZBpNwox.SetBalance( /*line OcCXKRVoVPm.go:1*/ new(uint256.Int).Add( /*line aa1LM2ar4.go:1*/ peYZBpNwox.Balance(), zGNz6Q9KhWC))
}

func (peYZBpNwox *stateObject) SubBalance(zGNz6Q9KhWC *uint256.Int) {
	if /*line PiV_1i.go:1*/ zGNz6Q9KhWC.IsZero() {
		return
	}
	/*line tF8Z6a.go:1*/ peYZBpNwox.SetBalance( /*line LiELzO9U3It.go:1*/ new(uint256.Int).Sub( /*line zCy3TwY8S.go:1*/ peYZBpNwox.Balance(), zGNz6Q9KhWC))
}

func (peYZBpNwox *stateObject) SetBalance(zGNz6Q9KhWC *uint256.Int) {
	/*line FObMjDpTx94x.go:1*/ peYZBpNwox.db.journal.uwXnrqq(eZiUht4Jv{
		xJtsHc1J0UJd: &peYZBpNwox.address,
		bQ3xDi:/*line JD7D0_XyOy.go:1*/ new(uint256.Int).Set(peYZBpNwox.data.Balance),
	})
	/*line t_HHTQnQ.go:1*/ peYZBpNwox.yNiGLlmM(zGNz6Q9KhWC)
}

func (peYZBpNwox *stateObject) yNiGLlmM(zGNz6Q9KhWC *uint256.Int) {
	peYZBpNwox.data.Balance = zGNz6Q9KhWC
}

func (peYZBpNwox *stateObject) eiV2C2(tC4SGV *StateDB) *stateObject {
	ws3tNR3MU := &stateObject{
		db:       tC4SGV,
		address:  peYZBpNwox.address,
		addrHash: peYZBpNwox.addrHash,
		origin:   peYZBpNwox.origin,
		data:     peYZBpNwox.data,
	}
	if peYZBpNwox.trie != nil {
		ws3tNR3MU.trie = /*line wdDXFwbQ1.go:1*/ tC4SGV.db.CopyTrie(peYZBpNwox.trie)
	}
	ws3tNR3MU.code = peYZBpNwox.code
	ws3tNR3MU.deploycode = peYZBpNwox.deploycode
	ws3tNR3MU.dirtyStorage = /*line n0wCDsYjO8.go:1*/ peYZBpNwox.dirtyStorage.Copy()
	ws3tNR3MU.originStorage = /*line L0qFy3F8.go:1*/ peYZBpNwox.originStorage.Copy()
	ws3tNR3MU.pendingStorage = /*line xpOOBw5.go:1*/ peYZBpNwox.pendingStorage.Copy()
	ws3tNR3MU.selfDestructed = peYZBpNwox.selfDestructed
	ws3tNR3MU.dirtyCode = peYZBpNwox.dirtyCode
	ws3tNR3MU.deleted = peYZBpNwox.deleted
	return ws3tNR3MU
}

func (peYZBpNwox *stateObject) Address() common.Address {
	return peYZBpNwox.address
}

func (peYZBpNwox *stateObject) Code() []byte {
	if peYZBpNwox.code != nil {
		return peYZBpNwox.code
	}
	if /*line qgQdLaX.go:1*/ bytes.Equal( /*line fSurdV5xbs.go:1*/ peYZBpNwox.CodeHash() /*line Ju7iemvK5.go:1*/, types.JhrQnETMFm.Bytes()) {
		return nil
	}
	acaO5J6, cYHA75qVq := /*line V9EOuoLak3.go:1*/ peYZBpNwox.db.db.ContractCode(peYZBpNwox.address /*line APVP3vL.go:1*/, common.BytesToHash( /*line aGb_zRG5cJr.go:1*/ peYZBpNwox.CodeHash()))
	if cYHA75qVq != nil {
		/*line Odi_UbcPOqD.go:1*/ peYZBpNwox.db.eMZRw9Dc( /*line JX7cHKXarNM.go:1*/ fmt.Errorf(func() /*line JXVNDqmt5q.go:1*/ string {
			seed := /*line JXVNDqmt5q.go:1*/ byte(22)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line JXVNDqmt5q.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line JXVNDqmt5q.go:1*/ fnc(117)(234)(27)(183)(51)(90)(184)(227)(14)(25)(182)(47)(20)(235)(31)(185)(58)(237)(10)(235)(78)(153)(45)(184)(26)(113)(179)
			return /*line JXVNDqmt5q.go:1*/ string(data)
		}(), /*line i04PCJ7Xz.go:1*/ peYZBpNwox.CodeHash(), cYHA75qVq))
	}
	peYZBpNwox.code = acaO5J6
	return acaO5J6
}

func (peYZBpNwox *stateObject) DeployCode() []byte {
	if peYZBpNwox.deploycode != nil {
		return peYZBpNwox.deploycode
	} else {
		return nil
	}
}

func (peYZBpNwox *stateObject) CodeSize() int {
	if peYZBpNwox.code != nil {
		return /*line lRYPV9.go:1*/ len(peYZBpNwox.code)
	}
	if /*line bzZXp5oKDzOP.go:1*/ bytes.Equal( /*line raIfarRJuPt.go:1*/ peYZBpNwox.CodeHash() /*line OMpQ0SBpm.go:1*/, types.JhrQnETMFm.Bytes()) {
		return 0
	}
	htUnPTZWCIPk, cYHA75qVq := /*line UCV6FOb.go:1*/ peYZBpNwox.db.db.ContractCodeSize(peYZBpNwox.address /*line CcYPBRhY6x.go:1*/, common.BytesToHash( /*line sRelNM9TVUfy.go:1*/ peYZBpNwox.CodeHash()))
	if cYHA75qVq != nil {
		/*line Cr3Nd7wBFqEP.go:1*/ peYZBpNwox.db.eMZRw9Dc( /*line I_0NOL.go:1*/ fmt.Errorf(func() /*line Le_IOTQ1_axN.go:1*/ string {
			seed := /*line Le_IOTQ1_axN.go:1*/ byte(76)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line Le_IOTQ1_axN.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line Le_IOTQ1_axN.go:1*/ fnc(47)(26)(251)(183)(51)(90)(184)(227)(14)(25)(182)(47)(20)(235)(31)(185)(33)(26)(247)(225)(69)(143)(65)(64)(154)(113)(179)
			return /*line Le_IOTQ1_axN.go:1*/ string(data)
		}(), /*line ZYeEm0cg.go:1*/ peYZBpNwox.CodeHash(), cYHA75qVq))
	}
	return htUnPTZWCIPk
}

func (peYZBpNwox *stateObject) SetCode(hCvoh5IR common.Hash, acaO5J6 []byte) {
	jj9Uv0 := /*line EEG4aXvZ.go:1*/ peYZBpNwox.Code()
	/*line LqZAm1EA.go:1*/ peYZBpNwox.db.journal.uwXnrqq(iOBxwHhpk{
		gxCiPlyej2gN: &peYZBpNwox.address,
		bxlFoAyG:/*line WS1e4iRwI.go:1*/ peYZBpNwox.CodeHash(),
		ilAkslSROBn: jj9Uv0,
	})
	/*line ykHkJBkIj.go:1*/ peYZBpNwox.rgsztNjLL9a(hCvoh5IR, acaO5J6)
}

func (peYZBpNwox *stateObject) SetDeploycode(acaO5J6 []byte) {
	peYZBpNwox.deploycode = acaO5J6
}

func (peYZBpNwox *stateObject) rgsztNjLL9a(hCvoh5IR common.Hash, acaO5J6 []byte) {
	peYZBpNwox.code = acaO5J6
	peYZBpNwox.data.CodeHash = hCvoh5IR[:]
	peYZBpNwox.dirtyCode = true
}

func (peYZBpNwox *stateObject) SetNonce(r1rvsu uint64) {
	/*line p5WAYl.go:1*/ peYZBpNwox.db.journal.uwXnrqq(dogSPq_gHV{
		pGTrAOrC:     &peYZBpNwox.address,
		tqCEH_Lia2a9: peYZBpNwox.data.Nonce,
	})
	/*line hzDrXxbHiNq.go:1*/ peYZBpNwox.ovEvA8Af(r1rvsu)
}

func (peYZBpNwox *stateObject) ovEvA8Af(r1rvsu uint64) {
	peYZBpNwox.data.Nonce = r1rvsu
}

func (peYZBpNwox *stateObject) CodeHash() []byte {
	return peYZBpNwox.data.CodeHash
}

func (peYZBpNwox *stateObject) Balance() *uint256.Int {
	return peYZBpNwox.data.Balance
}

func (peYZBpNwox *stateObject) Nonce() uint64 {
	return peYZBpNwox.data.Nonce
}

func (peYZBpNwox *stateObject) Root() common.Hash {
	return peYZBpNwox.data.Root
}

var _ bytes.Buffer
var _ = fmt.Append
var _ io.ByteReader

const _ = time.ANSIC

var _ types.AccessList
var _ = common.AbsolutePath
var _ = crypto.CompressPubkey
var _ = metrics.AccountingRegistry
var _ = rlp.AppendUint64
var _ trienode.MergedNodeSet
var _ = uint256.ErrBadBufferLength
