//line :1
package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type journalEntry interface {
	oOOlwUrQ0Ck(*StateDB)

	lMB3DJJd() *common.Address
}

type journal struct {
	entries []journalEntry
	dirties map[common.Address]int
}

func vuoy40vfttW() *journal {
	return &journal{
		dirties: /*line BgRBNOk.go:1*/ make(map[common.Address]int),
	}
}

func (xAAU8HoeJU *journal) uwXnrqq(eKg26VhLs journalEntry) {
	xAAU8HoeJU.entries = /*line UpgaWOam4eu.go:1*/ append(xAAU8HoeJU.entries, eKg26VhLs)
	if hv1UkMRaKf := /*line l7OsKFjT9k.go:1*/ eKg26VhLs.lMB3DJJd(); hv1UkMRaKf != nil {
		xAAU8HoeJU.dirties[*hv1UkMRaKf]++
	}
}

func (xAAU8HoeJU *journal) oOOlwUrQ0Ck(bPZabpmiZd *StateDB, anRdfDemah int) {
	for cDClIaFDS := /*line kpS84Tov.go:1*/ len(xAAU8HoeJU.entries) - 1; cDClIaFDS >= anRdfDemah; cDClIaFDS-- {

		/*line FI3xi3Te.go:1*/
		xAAU8HoeJU.entries[cDClIaFDS].oOOlwUrQ0Ck(bPZabpmiZd)

		if hv1UkMRaKf := /*line OvqeLk0.go:1*/ xAAU8HoeJU.entries[cDClIaFDS].lMB3DJJd(); hv1UkMRaKf != nil {
			if xAAU8HoeJU.dirties[*hv1UkMRaKf]--; xAAU8HoeJU.dirties[*hv1UkMRaKf] == 0 {
				/*line Y7kOdQne.go:1*/ delete(xAAU8HoeJU.dirties, *hv1UkMRaKf)
			}
		}
	}
	xAAU8HoeJU.entries = xAAU8HoeJU.entries[:anRdfDemah]
}

func (xAAU8HoeJU *journal) vdDLIhK(hv1UkMRaKf common.Address) {
	xAAU8HoeJU.dirties[hv1UkMRaKf]++
}

func (xAAU8HoeJU *journal) fgH6Ck43r() int {
	return /*line CLid99OaCD9t.go:1*/ len(xAAU8HoeJU.entries)
}

type (
	kLOasI4ckgD struct {
		huIthfk25W *common.Address
	}
	vF3b_q struct {
		yqSjsAK      *common.Address
		lpZWUm0WTZ5T *stateObject
		aeIv_xPXxXq4 bool
		_JFcGsAN     []byte
		dTp3Yn8laadd map[common.Hash][]byte

		b0aCZklQW bool
		p3F8IHE   []byte
		aItHQjJHk map[common.Hash][]byte
	}
	kzbdktS struct {
		sQRjYTvw  *common.Address
		yqkgNNBcP bool
		ba3Vrv    *uint256.Int
	}

	eZiUht4Jv struct {
		xJtsHc1J0UJd *common.Address
		bQ3xDi       *uint256.Int
	}
	dogSPq_gHV struct {
		pGTrAOrC     *common.Address
		tqCEH_Lia2a9 uint64
	}
	auS1vO struct {
		pGTrAOrC     *common.Address
		tqCEH_Lia2a9 uint64
	}
	nUsbOK struct {
		iLbaDHlKPp           *common.Address
		r9VPURO, eT8enfTtJwO common.Hash
	}
	iOBxwHhpk struct {
		gxCiPlyej2gN          *common.Address
		ilAkslSROBn, bxlFoAyG []byte
	}

	foFrwR8S4 struct {
		eT55nqw_VBa uint64
	}
	uRRh4T8d struct {
		inLGNw common.Hash
	}
	qhj_H1rUkvcr struct {
		gefaQDBFugQa common.Hash
	}
	fbLbWCnCjB81 struct {
		huIthfk25W *common.Address
	}

	vPXduwMcrq struct {
		rZPG8MA8NQb *common.Address
	}
	eboE8NSCCQU struct {
		p7jEjRwYLtL *common.Address
		ythV5glxZGp *common.Hash
	}

	yzAiyxK struct {
		iLbaDHlKPp           *common.Address
		r9VPURO, eT8enfTtJwO common.Hash
	}
)

func (gu0zGueC4zLM kLOasI4ckgD) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line OO4Hba5.go:1*/ delete(peYZBpNwox.stateObjects, *gu0zGueC4zLM.huIthfk25W)
	/*line Jnbbk8JhWmxU.go:1*/ delete(peYZBpNwox.stateObjectsDirty, *gu0zGueC4zLM.huIthfk25W)
}

func (gu0zGueC4zLM kLOasI4ckgD) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.huIthfk25W
}

func (gu0zGueC4zLM vF3b_q) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line UOTDcRT4D.go:1*/ peYZBpNwox.wTQEdk(gu0zGueC4zLM.lpZWUm0WTZ5T)
	if !gu0zGueC4zLM.aeIv_xPXxXq4 {
		/*line zGqVNRlng.go:1*/ delete(peYZBpNwox.stateObjectsDestruct, gu0zGueC4zLM.lpZWUm0WTZ5T.address)
	}
	if gu0zGueC4zLM._JFcGsAN != nil {
		peYZBpNwox.accounts[gu0zGueC4zLM.lpZWUm0WTZ5T.addrHash] = gu0zGueC4zLM._JFcGsAN
	}
	if gu0zGueC4zLM.dTp3Yn8laadd != nil {
		peYZBpNwox.storages[gu0zGueC4zLM.lpZWUm0WTZ5T.addrHash] = gu0zGueC4zLM.dTp3Yn8laadd
	}
	if gu0zGueC4zLM.b0aCZklQW {
		peYZBpNwox.accountsOrigin[gu0zGueC4zLM.lpZWUm0WTZ5T.address] = gu0zGueC4zLM.p3F8IHE
	}
	if gu0zGueC4zLM.aItHQjJHk != nil {
		peYZBpNwox.storagesOrigin[gu0zGueC4zLM.lpZWUm0WTZ5T.address] = gu0zGueC4zLM.aItHQjJHk
	}
}

func (gu0zGueC4zLM vF3b_q) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.yqSjsAK
}

func (gu0zGueC4zLM kzbdktS) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	ws3tNR3MU := /*line Nu8ziDtS.go:1*/ peYZBpNwox.nxJzqNXi(*gu0zGueC4zLM.sQRjYTvw)
	if ws3tNR3MU != nil {
		ws3tNR3MU.selfDestructed = gu0zGueC4zLM.yqkgNNBcP
		/*line sCbzgu3iZnk.go:1*/ ws3tNR3MU.yNiGLlmM(gu0zGueC4zLM.ba3Vrv)
	}
}

func (gu0zGueC4zLM kzbdktS) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.sQRjYTvw
}

var m8ZbLkxbBc = /*line md3JXml.go:1*/ common.HexToAddress(func() /*line NLa0KOJewX.go:1*/ string {
	data := [] /*line NLa0KOJewX.go:1*/ byte("\xc7\xeb\xd8\x0f\xfb0\xf4\xf2\x05\xf90\xff\xf400ƻ\x9600\xdf0\xe0\xa200?0\"\xfb\xf80/\xfd400?03")
	positions := [...]byte{16, 9, 29, 4, 9, 7, 11, 17, 33, 15, 6, 12, 28, 33, 30, 20, 3, 28, 1, 11, 3, 23, 2, 0, 1, 3, 1, 22, 1, 2, 32, 17, 37, 26, 11, 22, 0, 8, 30, 15, 33, 0, 34, 8}
	for i := 0; i < 44; i += 2 {
		localKey := /*line NLa0KOJewX.go:1*/ byte(i) + /*line NLa0KOJewX.go:1*/ byte(positions[i]^positions[i+1]) + 176
		data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
	}
	return /*line NLa0KOJewX.go:1*/ string(data)
}())

func (gu0zGueC4zLM fbLbWCnCjB81) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
}

func (gu0zGueC4zLM fbLbWCnCjB81) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.huIthfk25W
}

func (gu0zGueC4zLM eZiUht4Jv) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line keXPgl.go:1*/ peYZBpNwox.nxJzqNXi(*gu0zGueC4zLM.xJtsHc1J0UJd).yNiGLlmM(gu0zGueC4zLM.bQ3xDi)
}

func (gu0zGueC4zLM eZiUht4Jv) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.xJtsHc1J0UJd
}

func (gu0zGueC4zLM dogSPq_gHV) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line VfUMc7.go:1*/ peYZBpNwox.nxJzqNXi(*gu0zGueC4zLM.pGTrAOrC).ovEvA8Af(gu0zGueC4zLM.tqCEH_Lia2a9)
}

func (gu0zGueC4zLM dogSPq_gHV) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.pGTrAOrC
}

func (gu0zGueC4zLM iOBxwHhpk) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line qGyimQ8c.go:1*/ peYZBpNwox.nxJzqNXi(*gu0zGueC4zLM.gxCiPlyej2gN).rgsztNjLL9a( /*line sPQmHLeI_.go:1*/ common.BytesToHash(gu0zGueC4zLM.bxlFoAyG), gu0zGueC4zLM.ilAkslSROBn)
}

func (gu0zGueC4zLM iOBxwHhpk) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.gxCiPlyej2gN
}

func (gu0zGueC4zLM nUsbOK) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line bpTjVmn1.go:1*/ peYZBpNwox.nxJzqNXi(*gu0zGueC4zLM.iLbaDHlKPp).kInfXnsga(gu0zGueC4zLM.r9VPURO, gu0zGueC4zLM.eT8enfTtJwO)
}

func (gu0zGueC4zLM nUsbOK) lMB3DJJd() *common.Address {
	return gu0zGueC4zLM.iLbaDHlKPp
}

func (gu0zGueC4zLM yzAiyxK) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line uAj9EQDAKa.go:1*/ peYZBpNwox.vfrWO0H(*gu0zGueC4zLM.iLbaDHlKPp, gu0zGueC4zLM.r9VPURO, gu0zGueC4zLM.eT8enfTtJwO)
}

func (gu0zGueC4zLM yzAiyxK) lMB3DJJd() *common.Address {
	return nil
}

func (gu0zGueC4zLM foFrwR8S4) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	peYZBpNwox.refund = gu0zGueC4zLM.eT55nqw_VBa
}

func (gu0zGueC4zLM foFrwR8S4) lMB3DJJd() *common.Address {
	return nil
}

func (gu0zGueC4zLM uRRh4T8d) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	jrZPAbGCstr := peYZBpNwox.logs[gu0zGueC4zLM.inLGNw]
	if /*line RlCvSpvY6gU.go:1*/ len(jrZPAbGCstr) == 1 {
		/*line Pa9HLZ.go:1*/ delete(peYZBpNwox.logs, gu0zGueC4zLM.inLGNw)
	} else {
		peYZBpNwox.logs[gu0zGueC4zLM.inLGNw] = jrZPAbGCstr[: /*line i8uKrmQm.go:1*/ len(jrZPAbGCstr)-1]
	}
	peYZBpNwox.logSize--
}

func (gu0zGueC4zLM uRRh4T8d) lMB3DJJd() *common.Address {
	return nil
}

func (gu0zGueC4zLM qhj_H1rUkvcr) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line ASBpyrodkkI.go:1*/ delete(peYZBpNwox.preimages, gu0zGueC4zLM.gefaQDBFugQa)
}

func (gu0zGueC4zLM qhj_H1rUkvcr) lMB3DJJd() *common.Address {
	return nil
}

func (gu0zGueC4zLM vPXduwMcrq) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {

	/*line EMO7kna9tQE.go:1*/
	peYZBpNwox.accessList.DeleteAddress(*gu0zGueC4zLM.rZPG8MA8NQb)
}

func (gu0zGueC4zLM vPXduwMcrq) lMB3DJJd() *common.Address {
	return nil
}

func (gu0zGueC4zLM eboE8NSCCQU) oOOlwUrQ0Ck(peYZBpNwox *StateDB) {
	/*line Cl7w0Kx0gM.go:1*/ peYZBpNwox.accessList.DeleteSlot(*gu0zGueC4zLM.p7jEjRwYLtL, *gu0zGueC4zLM.ythV5glxZGp)
}

func (gu0zGueC4zLM eboE8NSCCQU) lMB3DJJd() *common.Address {
	return nil
}

var _ = common.AbsolutePath
var _ = uint256.ErrBadBufferLength
