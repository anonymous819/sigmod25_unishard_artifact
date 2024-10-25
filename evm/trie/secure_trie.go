//line :1
package trie

import (
	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/triedb/database"
)

type YBvUpdiJAOa = IetHqRYp3VR

func L9jleDwJXS(iiQYh2c common.Hash, zalM3_NR4XT common.Hash, m7S8SE common.Hash, kRC_1aK database.Database) (*YBvUpdiJAOa, error) {
	cIFJQQh := &ID{
		StateRoot: iiQYh2c,
		Owner:     zalM3_NR4XT,
		Root:      m7S8SE,
	}
	return /*line EVBPlk3yi.go:1*/ MMAH7dV2(cIFJQQh, kRC_1aK)
}

type IetHqRYp3VR struct {
	dHWIsToseO  Trie
	ouUFp32t    database.Database
	gOF8I4F9ry0 [common.HashLength]byte
	jIKLs9WGejG map[string][]byte
	o0bAtMty    *IetHqRYp3VR
}

func MMAH7dV2(cIFJQQh *ID, kRC_1aK database.Database) (*IetHqRYp3VR, error) {
	if kRC_1aK == nil {
		/*line EpP4RT.go:1*/ panic(func() /*line hJFLlr.go:1*/ string {
			seed := /*line hJFLlr.go:1*/ byte(74)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line hJFLlr.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line hJFLlr.go:1*/ fnc(42)(254)(247)(252)(201)(32)(23)(18)(220)(33)(237)(19)(241)(239)(30)(247)(252)(187)(67)(254)(11)(0)(249)(255)(188)(87)(242)(11)(244)(7)(6)(255)(172)(65)(191)(68)(253)(19)(237)(1)(255)(18)(242)
			return /*line hJFLlr.go:1*/ string(data)
		}())
	}
	gExqOlCxa53E, eZzE0KPS := /*line FkTSQ2nBm5b0.go:1*/ RJaQ3esB(cIFJQQh, kRC_1aK)
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	return &IetHqRYp3VR{dHWIsToseO: *gExqOlCxa53E, ouUFp32t: kRC_1aK}, nil
}

func (hkI2UXG_QBd *IetHqRYp3VR) MustGet(lhQWH7m []byte) []byte {
	return /*line AZonwMAl.go:1*/ hkI2UXG_QBd.dHWIsToseO.MustGet( /*line VC8GpEf1U.go:1*/ hkI2UXG_QBd.n0Fz11Tp(lhQWH7m))
}

func (hkI2UXG_QBd *IetHqRYp3VR) GetStorage(_ common.Address, lhQWH7m []byte) ([]byte, error) {
	bh98pCjEm, eZzE0KPS := /*line Pn1AKqiIQEf.go:1*/ hkI2UXG_QBd.dHWIsToseO.Get( /*line Iv8qCAFB8H.go:1*/ hkI2UXG_QBd.n0Fz11Tp(lhQWH7m))
	if eZzE0KPS != nil || /*line ldaoS40adswi.go:1*/ len(bh98pCjEm) == 0 {
		return nil, eZzE0KPS
	}
	_, eac03YO, _, eZzE0KPS := /*line zTQPJaxt6oN.go:1*/ rlp.Split(bh98pCjEm)
	return eac03YO, eZzE0KPS
}

func (hkI2UXG_QBd *IetHqRYp3VR) GetAccount(l98Od6T1zF common.Address) (*types.StateAccount, error) {
	b6Gjyt0P8, eZzE0KPS := /*line C3n7N4H.go:1*/ hkI2UXG_QBd.dHWIsToseO.Get( /*line ZH9pfMIXH.go:1*/ hkI2UXG_QBd.n0Fz11Tp( /*line lfaptpcFJWhX.go:1*/ l98Od6T1zF.Bytes()))
	if b6Gjyt0P8 == nil || eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	jnXtTaTmH := /*line mvnN7zi.go:1*/ new(types.StateAccount)
	eZzE0KPS = /*line tAto515r.go:1*/ rlp.DecodeBytes(b6Gjyt0P8, jnXtTaTmH)
	return jnXtTaTmH, eZzE0KPS
}

func (hkI2UXG_QBd *IetHqRYp3VR) GetAccountByHash(kzkViH5 common.Hash) (*types.StateAccount, error) {
	b6Gjyt0P8, eZzE0KPS := /*line cs_w97.go:1*/ hkI2UXG_QBd.dHWIsToseO.Get( /*line HRtpba.go:1*/ kzkViH5.Bytes())
	if b6Gjyt0P8 == nil || eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	jnXtTaTmH := /*line SEMzcqErSC9k.go:1*/ new(types.StateAccount)
	eZzE0KPS = /*line UZhgrE5.go:1*/ rlp.DecodeBytes(b6Gjyt0P8, jnXtTaTmH)
	return jnXtTaTmH, eZzE0KPS
}

func (hkI2UXG_QBd *IetHqRYp3VR) GetNode(qiRG6lpeaFl []byte) ([]byte, int, error) {
	return /*line psOsMddQ6.go:1*/ hkI2UXG_QBd.dHWIsToseO.GetNode(qiRG6lpeaFl)
}

func (hkI2UXG_QBd *IetHqRYp3VR) MustUpdate(lhQWH7m, ar76Sw []byte) {
	gyjPR3tec := /*line qFt4Kj4q.go:1*/ hkI2UXG_QBd.n0Fz11Tp(lhQWH7m)
	/*line GtBJmuCTg.go:1*/ hkI2UXG_QBd.dHWIsToseO.MustUpdate(gyjPR3tec, ar76Sw)
	/*line Rg7gGPh.go:1*/ hkI2UXG_QBd.q1pomc6s()[ /*line Yae3PFa1Ds_b.go:1*/ string(gyjPR3tec)] = /*line WQ6pR_f8rR.go:1*/ common.CopyBytes(lhQWH7m)
}

func (hkI2UXG_QBd *IetHqRYp3VR) UpdateStorage(_ common.Address, lhQWH7m, ar76Sw []byte) error {
	gyjPR3tec := /*line m3UM_l.go:1*/ hkI2UXG_QBd.n0Fz11Tp(lhQWH7m)
	eBWWX0caQi, _ := /*line myDI6nS9V1.go:1*/ rlp.EncodeToBytes(ar76Sw)
	eZzE0KPS := /*line T9Q4qYpIQ.go:1*/ hkI2UXG_QBd.dHWIsToseO.Update(gyjPR3tec, eBWWX0caQi)
	if eZzE0KPS != nil {
		return eZzE0KPS
	}
	/*line NX2a3Sg.go:1*/ hkI2UXG_QBd.q1pomc6s()[ /*line F41DMOiry.go:1*/ string(gyjPR3tec)] = /*line UOxsaEBZni.go:1*/ common.CopyBytes(lhQWH7m)
	return nil
}

func (hkI2UXG_QBd *IetHqRYp3VR) UpdateAccount(l98Od6T1zF common.Address, gU1J_yS *types.StateAccount) error {
	gyjPR3tec := /*line g871EYRqx45.go:1*/ hkI2UXG_QBd.n0Fz11Tp( /*line EmbGhECV.go:1*/ l98Od6T1zF.Bytes())
	qNY8zG, eZzE0KPS := /*line DVM2Ij6.go:1*/ rlp.EncodeToBytes(gU1J_yS)
	if eZzE0KPS != nil {
		return eZzE0KPS
	}
	if eZzE0KPS := /*line J7LBSaMCU.go:1*/ hkI2UXG_QBd.dHWIsToseO.Update(gyjPR3tec, qNY8zG); eZzE0KPS != nil {
		return eZzE0KPS
	}
	/*line NsOuuJaS.go:1*/ hkI2UXG_QBd.q1pomc6s()[ /*line CC0gC0.go:1*/ string(gyjPR3tec)] = /*line DSGZLj_QEs84.go:1*/ l98Od6T1zF.Bytes()
	return nil
}

func (hkI2UXG_QBd *IetHqRYp3VR) UpdateContractCode(_ common.Address, _ common.Hash, _ []byte) error {
	return nil
}

func (hkI2UXG_QBd *IetHqRYp3VR) MustDelete(lhQWH7m []byte) {
	gyjPR3tec := /*line TDnts6GDNdxJ.go:1*/ hkI2UXG_QBd.n0Fz11Tp(lhQWH7m)
	/*line y5WbZk7i5uvD.go:1*/ delete( /*line q71xK3Q.go:1*/ hkI2UXG_QBd.q1pomc6s() /*line he5N8MJvlk6.go:1*/, string(gyjPR3tec))
	/*line PnUXUGaWUep.go:1*/ hkI2UXG_QBd.dHWIsToseO.MustDelete(gyjPR3tec)
}

func (hkI2UXG_QBd *IetHqRYp3VR) DeleteStorage(_ common.Address, lhQWH7m []byte) error {
	gyjPR3tec := /*line Rx9r4Q.go:1*/ hkI2UXG_QBd.n0Fz11Tp(lhQWH7m)
	/*line uYEcyRoHEZ.go:1*/ delete( /*line X_5_hWhs8ye.go:1*/ hkI2UXG_QBd.q1pomc6s() /*line oLJw00M.go:1*/, string(gyjPR3tec))
	return /*line yvLUJG.go:1*/ hkI2UXG_QBd.dHWIsToseO.Delete(gyjPR3tec)
}

func (hkI2UXG_QBd *IetHqRYp3VR) DeleteAccount(l98Od6T1zF common.Address) error {
	gyjPR3tec := /*line TecICavK1xuW.go:1*/ hkI2UXG_QBd.n0Fz11Tp( /*line oVQxbCK0q.go:1*/ l98Od6T1zF.Bytes())
	/*line lflp9JIIl4F.go:1*/ delete( /*line nvj5_Z.go:1*/ hkI2UXG_QBd.q1pomc6s() /*line sTd67JX.go:1*/, string(gyjPR3tec))
	return /*line sMwY5S2YYT.go:1*/ hkI2UXG_QBd.dHWIsToseO.Delete(gyjPR3tec)
}

func (hkI2UXG_QBd *IetHqRYp3VR) GetKey(yc5bYOK65eC []byte) []byte {
	if lhQWH7m, yY_yPSd8vG := /*line JqexhT5uy5q.go:1*/ hkI2UXG_QBd.q1pomc6s()[ /*line OstnOuVab.go:1*/ string(yc5bYOK65eC)]; yY_yPSd8vG {
		return lhQWH7m
	}
	return /*line uE60XjzE.go:1*/ hkI2UXG_QBd.ouUFp32t.Preimage( /*line lUMvyH4ni6H.go:1*/ common.BytesToHash(yc5bYOK65eC))
}

func (hkI2UXG_QBd *IetHqRYp3VR) Commit(xfPJhFZM25R bool) (common.Hash, *trienode.NodeSet, error) {

	if /*line CtcEK8H1GvnX.go:1*/ len( /*line mYJbMLJP4.go:1*/ hkI2UXG_QBd.q1pomc6s()) > 0 {
		vmMt3OgPQcVH := /*line AJPlTu5On.go:1*/ make(map[common.Hash][]byte)
		for gyjPR3tec, lhQWH7m := range hkI2UXG_QBd.jIKLs9WGejG {
			vmMt3OgPQcVH[ /*line _EovCZM.go:1*/ common.BytesToHash([] /*line LjKvBkg.go:1*/ byte(gyjPR3tec))] = lhQWH7m
		}
		/*line xFudfRg5T9F.go:1*/ hkI2UXG_QBd.ouUFp32t.InsertPreimage(vmMt3OgPQcVH)
		hkI2UXG_QBd.jIKLs9WGejG = /*line fCL4cSoTqQSE.go:1*/ make(map[string][]byte)
	}

	return /*line Ucfvrdk.go:1*/ hkI2UXG_QBd.dHWIsToseO.Commit(xfPJhFZM25R)
}

func (hkI2UXG_QBd *IetHqRYp3VR) Hash() common.Hash {
	return /*line IESaAa.go:1*/ hkI2UXG_QBd.dHWIsToseO.Hash()
}

func (hkI2UXG_QBd *IetHqRYp3VR) Copy() *IetHqRYp3VR {
	return &IetHqRYp3VR{
		dHWIsToseO:  * /*line aavLbTzx228.go:1*/ hkI2UXG_QBd.dHWIsToseO.Copy(),
		ouUFp32t:    hkI2UXG_QBd.ouUFp32t,
		jIKLs9WGejG: hkI2UXG_QBd.jIKLs9WGejG,
	}
}

func (hkI2UXG_QBd *IetHqRYp3VR) NodeIterator(iS1397t []byte) (FBy5J0gbkg, error) {
	return /*line fdqyDjkdQ1.go:1*/ hkI2UXG_QBd.dHWIsToseO.NodeIterator(iS1397t)
}

func (hkI2UXG_QBd *IetHqRYp3VR) MustNodeIterator(iS1397t []byte) FBy5J0gbkg {
	return /*line BJWXobPe.go:1*/ hkI2UXG_QBd.dHWIsToseO.MustNodeIterator(iS1397t)
}

func (hkI2UXG_QBd *IetHqRYp3VR) n0Fz11Tp(lhQWH7m []byte) []byte {
	klE3zy := /*line WsWXbe.go:1*/ dCZJziX(false)
	/*line vYVGaWMlhMaH.go:1*/ klE3zy.vbM2IWKj6j.Reset()
	/*line CLeLlPkjPqG.go:1*/ klE3zy.vbM2IWKj6j.Write(lhQWH7m)
	/*line tei9_ZJQTLFV.go:1*/ klE3zy.vbM2IWKj6j.Read(hkI2UXG_QBd.gOF8I4F9ry0[:])
	/*line DGkg64.go:1*/ ptIMYRK9(klE3zy)
	return hkI2UXG_QBd.gOF8I4F9ry0[:]
}

func (hkI2UXG_QBd *IetHqRYp3VR) q1pomc6s() map[string][]byte {
	if hkI2UXG_QBd != hkI2UXG_QBd.o0bAtMty {
		hkI2UXG_QBd.o0bAtMty = hkI2UXG_QBd
		hkI2UXG_QBd.jIKLs9WGejG = /*line oq4Zkrs.go:1*/ make(map[string][]byte)
	}
	return hkI2UXG_QBd.jIKLs9WGejG
}

var _ types.AccessList
var _ = common.AbsolutePath
var _ = rlp.AppendUint64
var _ trienode.MergedNodeSet
var _ database.Database
