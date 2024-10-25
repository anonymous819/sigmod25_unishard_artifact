//line :1
package trie

import (
	"errors"
	"fmt"
	"sync"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/prque"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

var BENf3izjF = /*line uQjNzGZ_k1.go:1*/ errors.New(func() /*line dXamD3m4.go:1*/ string {
	seed := /*line dXamD3m4.go:1*/ byte(34)
	var data []byte
	type decFunc func(byte) decFunc
	var fnc decFunc
	fnc = func(x byte) decFunc { data = /*line dXamD3m4.go:1*/ append(data, x-seed); seed += x; return fnc }
	/*line dXamD3m4.go:1*/ fnc(144)(33)(71)(58)(198)(127)(10)(24)(32)(78)(157)(43)(85)
	return /*line dXamD3m4.go:1*/ string(data)
}())

var A7S3wVrk = /*line JXVvZawOu.go:1*/ errors.New(func() /*line FU5iRdtirG.go:1*/ string {
	seed := /*line FU5iRdtirG.go:1*/ byte(137)
	var data []byte
	type decFunc func(byte) decFunc
	var fnc decFunc
	fnc = func(x byte) decFunc { data = /*line FU5iRdtirG.go:1*/ append(data, x-seed); seed += x; return fnc }
	/*line FU5iRdtirG.go:1*/ fnc(234)(223)(196)(123)(242)(231)(227)(109)(42)(86)(169)(70)(142)(42)(84)(154)(51)
	return /*line FU5iRdtirG.go:1*/ string(data)
}())

const maxFetchesPerDepth = 16384

var (
	_zqqN05gnNB = /*line PBb1O0t4.go:1*/ metrics.NewRegisteredGauge(func() /*line aRKByrjpsxIY.go:1*/ string {
		data := /*line aRKByrjpsxIY.go:1*/ make([]byte, 0, 17)
		i := 7
		decryptKey := 162
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				data = /*line aRKByrjpsxIY.go:1*/ append(data, "\x00\f\xfe"...,
				)
				i = 8
			case 7:
				i = 1
				data = /*line aRKByrjpsxIY.go:1*/ append(data, 10)
			case 3:
				i = 0
				data = /*line aRKByrjpsxIY.go:1*/ append(data, "\xce\x00\x02\x06"...,
				)
			case 1:
				data = /*line aRKByrjpsxIY.go:1*/ append(data, 9)
				i = 2
			case 5:
				data = /*line aRKByrjpsxIY.go:1*/ append(data, "\xff\x01"...,
				)
				i = 3
			case 2:
				i = 6
				data = /*line aRKByrjpsxIY.go:1*/ append(data, "\xfd\xfa\xc1"...,
				)
			case 8:
				for y := range data {
					data[y] = data[y] - /*line aRKByrjpsxIY.go:1*/ byte(decryptKey^y)
				}
				i = 4
			case 6:
				data = /*line aRKByrjpsxIY.go:1*/ append(data, "\x06\t"...,
				)
				i = 5
			}
		}
		return /*line aRKByrjpsxIY.go:1*/ string(data)
	}(), nil)

	lnsopxDpNL = /*line CDnCBLS8.go:1*/ metrics.NewRegisteredGauge(func() /*line ACsnKsK2PU.go:1*/ string {
		key := [] /*line ACsnKsK2PU.go:1*/ byte("oH\xaf\xbd\xe5\aj\xdb\xe4̥\xfe\v\x88\xf3\xa6")
		data := [] /*line ACsnKsK2PU.go:1*/ byte("\x05*\xba\xa8Jl\x0f\x93\x7fc\xc7qd\xe3\x82\xca")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line ACsnKsK2PU.go:1*/ string(data)
	}(), nil)

	fvkJlHhIoeT2 = /*line AOoVEUf.go:1*/ metrics.NewRegisteredGauge(func() /*line JEOY9shu.go:1*/ string {
		seed := /*line JEOY9shu.go:1*/ byte(0)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line JEOY9shu.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line JEOY9shu.go:1*/ fnc(116)(6)(19)(232)(90)(188)(242)(19)(243)(172)(65)(31)(235)(31)(234)(172)(78)(30)(248)(252)(250)(231)(4)
		return /*line JEOY9shu.go:1*/ string(data)
	}(), nil)

	g7uspIlC4v = /*line kujspSgjgN2.go:1*/ metrics.NewRegisteredGauge(func() /*line l5fqBo86fV.go:1*/ string {
		data := /*line l5fqBo86fV.go:1*/ make([]byte, 0, 24)
		i := 3
		decryptKey := 118
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 8:
				for y := range data {
					data[y] = data[y] - /*line l5fqBo86fV.go:1*/ byte(decryptKey^y)
				}
				i = 2
			case 0:
				data = /*line l5fqBo86fV.go:1*/ append(data, "je6y"...,
				)
				i = 7
			case 3:
				i = 6
				data = /*line l5fqBo86fV.go:1*/ append(data, 119)
			case 5:
				i = 9
				data = /*line l5fqBo86fV.go:1*/ append(data, "9ww"...,
				)
			case 6:
				i = 0
				data = /*line l5fqBo86fV.go:1*/ append(data, 116)
			case 9:
				i = 4
				data = /*line l5fqBo86fV.go:1*/ append(data, "ss\x80"...,
				)
			case 1:
				data = /*line l5fqBo86fV.go:1*/ append(data, "}z"...,
				)
				i = 8
			case 4:
				data = /*line l5fqBo86fV.go:1*/ append(data, ";\x86\x86"...,
				)
				i = 10
			case 11:
				data = /*line l5fqBo86fV.go:1*/ append(data, "\x82x"...,
				)
				i = 1
			case 10:
				i = 11
				data = /*line l5fqBo86fV.go:1*/ append(data, 128)
			case 7:
				i = 5
				data = /*line l5fqBo86fV.go:1*/ append(data, "~rn"...,
				)
			}
		}
		return /*line l5fqBo86fV.go:1*/ string(data)
	}(), nil)

	s125TYfsu0W = /*line O69u9yRX8QX.go:1*/ metrics.NewRegisteredGauge(func() /*line jqu0oxBrhszj.go:1*/ string {
		seed := /*line jqu0oxBrhszj.go:1*/ byte(40)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line jqu0oxBrhszj.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line jqu0oxBrhszj.go:1*/ fnc(92)(246)(19)(232)(90)(188)(242)(19)(243)(172)(76)(20)(235)(31)(234)
		return /*line jqu0oxBrhszj.go:1*/ string(data)
	}(), nil)
)

type UiEAYQ [][]byte

func JwkofVV(qiRG6lpeaFl []byte) UiEAYQ {

	if /*line GoUABRRQa.go:1*/ len(qiRG6lpeaFl) < 64 {
		return UiEAYQ{ /*line IrCzwephKON.go:1*/ pSumbut7Z(qiRG6lpeaFl)}
	}
	return UiEAYQ{ /*line VpwGTi1mPtwe.go:1*/ iZl0DAYMB(qiRG6lpeaFl[:64]) /*line sEYq5GFIPaW.go:1*/, pSumbut7Z(qiRG6lpeaFl[64:])}
}

type C2UpTEtk3IR func(ywhO7DKTqn [][]byte, qiRG6lpeaFl []byte, xePr0Cqmao []byte, nmFIUUSGjl common.Hash, w78Jh0If6Su0 []byte) error

type ppmclce_a struct {
	gkgnXT8aL common.Hash
	dE1ARYil  []byte
	mtTpVM    []byte

	yx1kuQ0QEs0u *ppmclce_a
	hHtND2       int
	baDbjQlsCI   C2UpTEtk3IR
}

type nLazi_HEl3 struct {
	pt0XJRWq     common.Hash
	fNxs_ct2     []byte
	gxZxr1e4F0Ms []byte
	aK_sok       []*ppmclce_a
}

type XDAbLg0qTTF struct {
	I0rxZfpi string
	OAQsZC7I []byte
}

type F0vWRiZdlN0v struct {
	Ou81btbhI common.Hash
	H_jvwIxZq []byte
}

type fwA_J0erfLQ struct {
	gAvK9j      common.Hash
	qZSlzWMGYFA []byte
	jms_XktV    []byte
	pBaEHW1     common.Hash
}

func (dMBBt1tBjpd *fwA_J0erfLQ) g7xlc_Y1Bm() bool {
	return /*line JUzFGD.go:1*/ len(dMBBt1tBjpd.jms_XktV) == 0
}

type fP5DKl7 struct {
	u0OT5c     string
	dble5Cg3th map[common.Hash][]byte
	kTnGvbt2g  []fwA_J0erfLQ
	tRaHa9t    uint64
}

func nAsYnT1(d07z8Tvz string) *fP5DKl7 {
	return &fP5DKl7{
		u0OT5c: d07z8Tvz,
		dble5Cg3th:/*line Hjo_RBWWpp.go:1*/ make(map[common.Hash][]byte),
	}
}

func (xzcSEuMx *fP5DKl7) hhiYbrnJ(rNuN0sMPDJ common.Hash) bool {
	_, yY_yPSd8vG := xzcSEuMx.dble5Cg3th[rNuN0sMPDJ]
	return yY_yPSd8vG
}

func (xzcSEuMx *fP5DKl7) eWegJqd6F6(rNuN0sMPDJ common.Hash, abS5SZPC []byte) {
	xzcSEuMx.dble5Cg3th[rNuN0sMPDJ] = abS5SZPC
	xzcSEuMx.tRaHa9t += common.HashLength + /*line EL3C4ps.go:1*/ uint64( /*line Bq_eiE.go:1*/ len(abS5SZPC))
}

func (xzcSEuMx *fP5DKl7) hgsdDYY(zalM3_NR4XT common.Hash, qiRG6lpeaFl []byte, aBHduJG []byte, rNuN0sMPDJ common.Hash) {
	if xzcSEuMx.u0OT5c == rawdb.PathScheme {
		if zalM3_NR4XT == (common.Hash{}) {
			xzcSEuMx.tRaHa9t += /*line _CassQf.go:1*/ uint64( /*line mFEK0u4M.go:1*/ len(qiRG6lpeaFl) + /*line cQyjfiLN4.go:1*/ len(aBHduJG))
		} else {
			xzcSEuMx.tRaHa9t += common.HashLength + /*line NJtB41q_7b.go:1*/ uint64( /*line d9NjISWN.go:1*/ len(qiRG6lpeaFl)+ /*line iM8oqzaJO5rT.go:1*/ len(aBHduJG))
		}
	} else {
		xzcSEuMx.tRaHa9t += common.HashLength + /*line UeCjo9uzkD.go:1*/ uint64( /*line bxG2zl7.go:1*/ len(aBHduJG))
	}
	xzcSEuMx.kTnGvbt2g = /*line G6yAd6.go:1*/ append(xzcSEuMx.kTnGvbt2g, fwA_J0erfLQ{
		gAvK9j:      zalM3_NR4XT,
		qZSlzWMGYFA: qiRG6lpeaFl,
		jms_XktV:    aBHduJG,
		pBaEHW1:     rNuN0sMPDJ,
	})
}

func (xzcSEuMx *fP5DKl7) yyrftSiJ4F5N(zalM3_NR4XT common.Hash, qiRG6lpeaFl []byte) {
	if xzcSEuMx.u0OT5c != rawdb.PathScheme {
		/*line Anht39.go:1*/ log.Error(func() /*line s5yZ_BDyzB.go:1*/ string {
			key := [] /*line s5yZ_BDyzB.go:1*/ byte("Y\xdc\xce\x1d\x019\xb5\xf0\xb7\xbb\x1d\x12\x0e\xee\xff\xe0\x1a\x94l]\xe2\x8eN\xfd")
			data := [] /*line s5yZ_BDyzB.go:1*/ byte("\xaeJ3\x95q\x9e\x18d\x1c\x1f=\x80}Rd\x00~\xf9\xd8\xc2V\xf7\xbdk")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line s5yZ_BDyzB.go:1*/ string(data)
		}(), "owner", zalM3_NR4XT, "path", qiRG6lpeaFl, "scheme", xzcSEuMx.u0OT5c)
		return
	}
	if zalM3_NR4XT == (common.Hash{}) {
		xzcSEuMx.tRaHa9t += /*line HYgP4G8DDo.go:1*/ uint64( /*line q2o9ys.go:1*/ len(qiRG6lpeaFl))
	} else {
		xzcSEuMx.tRaHa9t += common.HashLength + /*line PYNMVrU9_Nf.go:1*/ uint64( /*line hIvvy1Hj.go:1*/ len(qiRG6lpeaFl))
	}
	xzcSEuMx.kTnGvbt2g = /*line Gwjb5vym.go:1*/ append(xzcSEuMx.kTnGvbt2g, fwA_J0erfLQ{
		gAvK9j:      zalM3_NR4XT,
		qZSlzWMGYFA: qiRG6lpeaFl,
	})
}

type HZAeEZr_ struct {
	yHer_HF4aV   string
	gBuukqS      ethdb.KeyValueReader
	hcmulRbxXrsz *fP5DKl7
	u1AhfCK      map[string]*ppmclce_a
	gVAbVpvBn    map[common.Hash]*nLazi_HEl3
	pafmayfkH    *prque.Prque[int64, any]
	bLY5FqFWk    map[int]int
}

func DMuKeU6T5BQ5(m7S8SE common.Hash, g5YbxqvpQ ethdb.KeyValueReader, irP7xy57BA35 C2UpTEtk3IR, d07z8Tvz string) *HZAeEZr_ {
	nnRE_k := &HZAeEZr_{
		yHer_HF4aV: d07z8Tvz,
		gBuukqS:    g5YbxqvpQ,
		hcmulRbxXrsz:/*line TDpYcYQ.go:1*/ nAsYnT1(d07z8Tvz),
		u1AhfCK:/*line BvGWnDw.go:1*/ make(map[string]*ppmclce_a),
		gVAbVpvBn:/*line F360ZK4.go:1*/ make(map[common.Hash]*nLazi_HEl3),
		pafmayfkH:/*line B2QWcDT.go:1*/ prque.New[int64, any](nil),
		bLY5FqFWk:/*line xHiU1ni5S.go:1*/ make(map[int]int),
	}
	/*line Jy80h9S.go:1*/ nnRE_k.AddSubTrie(m7S8SE, nil, common.Hash{}, nil, irP7xy57BA35)
	return nnRE_k
}

func (mvwKyZFJ6 *HZAeEZr_) AddSubTrie(m7S8SE common.Hash, qiRG6lpeaFl []byte, nmFIUUSGjl common.Hash, w78Jh0If6Su0 []byte, irP7xy57BA35 C2UpTEtk3IR) {
	if m7S8SE == types.NrGuaNA21 {
		return
	}
	zalM3_NR4XT, z9OkvU := /*line KlBBRPXf.go:1*/ RVW0hv81imRn(qiRG6lpeaFl)
	aqgXaP1JtoO, yOI3M1hHTg := /*line aH83u_OXZ9.go:1*/ mvwKyZFJ6.j7FeOp5b(zalM3_NR4XT, z9OkvU, m7S8SE)
	if aqgXaP1JtoO {

		return
	} else if yOI3M1hHTg {

		/*line yL4tGW.go:1*/
		mvwKyZFJ6.hcmulRbxXrsz.yyrftSiJ4F5N(zalM3_NR4XT, z9OkvU)
	}

	raCuLyGdWdp1 := &ppmclce_a{
		gkgnXT8aL:  m7S8SE,
		dE1ARYil:   qiRG6lpeaFl,
		baDbjQlsCI: irP7xy57BA35,
	}

	if nmFIUUSGjl != (common.Hash{}) {
		bvp1dT2K5 := mvwKyZFJ6.u1AhfCK[ /*line kE7OGTqT.go:1*/ string(w78Jh0If6Su0)]
		if bvp1dT2K5 == nil {
			/*line Uo2nN1.go:1*/ panic( /*line I3IcPbw.go:1*/ fmt.Sprintf(func() /*line RZQR_t5ZVKBy.go:1*/ string {
				fullData := [] /*line RZQR_t5ZVKBy.go:1*/ byte("\xd7u\xb7\x05\xc28\x9d\x8d\x9dj\xa5\x1a\xee\xf6\xf0\x03\v\xbc\xbd\xb3\xdb\xf7ܷ\xa1\x060H\xbcz\x94\xfcN)o\x14\x91,V]\aq2J\xe3\xff\xa9j:\x96\xe8W\x80%U1\xad\x10_\x882\x87")
				idxKey := [] /*line RZQR_t5ZVKBy.go:1*/ byte("\xf9lǋ\x12)\xd2\xfcK\x84 \x18f\xdf\f\x94")
				data := /*line RZQR_t5ZVKBy.go:1*/ make([]byte, 0, 32)
				data = /*line RZQR_t5ZVKBy.go:1*/ append(data, fullData[62^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[11])]-fullData[52^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[11])], fullData[47^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])]+fullData[30^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])], fullData[193^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[13])]-fullData[227^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[13])], fullData[234^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[6])]^fullData[230^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[6])], fullData[45^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[5])]-fullData[9^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[5])], fullData[105^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])]^fullData[71^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])], fullData[223^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[7])]+fullData[202^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[7])], fullData[222^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[2])]-fullData[223^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[2])], fullData[25^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])]^fullData[34^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])], fullData[65^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[12])]-fullData[121^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[12])], fullData[39^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])]-fullData[16^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])], fullData[253^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[6])]-fullData[250^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[6])], fullData[109^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])]-fullData[85^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])], fullData[235^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[7])]+fullData[237^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[7])], fullData[102^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])]-fullData[91^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])], fullData[111^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])]-fullData[93^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])], fullData[6^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])]^fullData[60^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[4])], fullData[4^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[14])]^fullData[30^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[14])], fullData[94^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])]-fullData[113^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])], fullData[200^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[6])]^fullData[232^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[6])], fullData[59^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[10])]+fullData[5^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[10])], fullData[176^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[15])]-fullData[189^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[15])], fullData[23^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[11])]-fullData[30^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[11])], fullData[58^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[5])]^fullData[63^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[5])], fullData[41^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[10])]+fullData[48^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[10])], fullData[166^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[3])]+fullData[169^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[3])], fullData[254^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[0])]+fullData[249^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[0])], fullData[45^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[10])]-fullData[60^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[10])], fullData[106^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[8])]+fullData[94^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[8])], fullData[85^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[12])]-fullData[76^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[12])], fullData[98^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])]^fullData[87^ /*line RZQR_t5ZVKBy.go:1*/ int(idxKey[1])])
				return /*line RZQR_t5ZVKBy.go:1*/ string(data)
			}(), nmFIUUSGjl))
		}
		bvp1dT2K5.hHtND2++
		raCuLyGdWdp1.yx1kuQ0QEs0u = bvp1dT2K5
	}
	/*line tWFbM9c5h.go:1*/ mvwKyZFJ6.bfuqiir1d2(raCuLyGdWdp1)
}

func (mvwKyZFJ6 *HZAeEZr_) AddCodeEntry(rNuN0sMPDJ common.Hash, qiRG6lpeaFl []byte, nmFIUUSGjl common.Hash, w78Jh0If6Su0 []byte) {

	if rNuN0sMPDJ == types.JhrQnETMFm {
		return
	}
	if /*line gtvqZhz5X.go:1*/ mvwKyZFJ6.hcmulRbxXrsz.hhiYbrnJ(rNuN0sMPDJ) {
		return
	}

	if /*line gnV4vQF.go:1*/ rawdb.HasCodeWithPrefix(mvwKyZFJ6.gBuukqS, rNuN0sMPDJ) {
		return
	}

	raCuLyGdWdp1 := &nLazi_HEl3{
		fNxs_ct2: qiRG6lpeaFl,
		pt0XJRWq: rNuN0sMPDJ,
	}

	if nmFIUUSGjl != (common.Hash{}) {
		bvp1dT2K5 := mvwKyZFJ6.u1AhfCK[ /*line Mdp7z89mg.go:1*/ string(w78Jh0If6Su0)]
		if bvp1dT2K5 == nil {
			/*line QXbZZM.go:1*/ panic( /*line PgGL10ldUrsI.go:1*/ fmt.Sprintf(func() /*line F87KDOXcCuS.go:1*/ string {
				fullData := [] /*line F87KDOXcCuS.go:1*/ byte("\\x\x1ag\xc1>_.\x153\x93\xed\x92\xd4@rHe/V\xfd\b15\xecx$#\t\xb9\xe8\x87\af\xce<\x00\x06`I\x98vw\x96\xf0\x04W\xee\xcc\x1c\xe2x[\xa8\x8a\b^\xf0\x83:\x92\xef\xe4\xa7")
				idxKey := [] /*line F87KDOXcCuS.go:1*/ byte("\xb5\xfeo\x86")
				data := /*line F87KDOXcCuS.go:1*/ make([]byte, 0, 33)
				data = /*line F87KDOXcCuS.go:1*/ append(data, fullData[122^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])]-fullData[68^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])], fullData[177^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])]-fullData[147^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])], fullData[153^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]+fullData[170^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[91^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])]^fullData[70^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])], fullData[141^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]+fullData[181^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[133^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]^fullData[154^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[83^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])]+fullData[93^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])], fullData[200^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]+fullData[224^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[248^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]+fullData[252^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[233^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]-fullData[246^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[164^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]+fullData[140^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[201^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]^fullData[223^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[207^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]-fullData[227^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[162^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]+fullData[151^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[193^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]^fullData[243^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[136^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]-fullData[182^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[176^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])]+fullData[163^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])], fullData[149^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]^fullData[156^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[157^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])]-fullData[180^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])], fullData[134^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]-fullData[169^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[158^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]^fullData[188^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])], fullData[221^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]^fullData[238^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[211^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]-fullData[192^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[236^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]^fullData[217^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[198^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]-fullData[195^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[96^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])]-fullData[123^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])], fullData[128^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])]-fullData[142^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])], fullData[242^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])]-fullData[249^ /*line F87KDOXcCuS.go:1*/ int(idxKey[1])], fullData[79^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])]+fullData[102^ /*line F87KDOXcCuS.go:1*/ int(idxKey[2])], fullData[159^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])]-fullData[155^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])], fullData[144^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])]^fullData[174^ /*line F87KDOXcCuS.go:1*/ int(idxKey[0])], fullData[191^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])]-fullData[159^ /*line F87KDOXcCuS.go:1*/ int(idxKey[3])])
				return /*line F87KDOXcCuS.go:1*/ string(data)
			}(), nmFIUUSGjl))
		}
		bvp1dT2K5.hHtND2++
		raCuLyGdWdp1.aK_sok = /*line h135atwIZ7.go:1*/ append(raCuLyGdWdp1.aK_sok, bvp1dT2K5)
	}
	/*line fHsX2HnG.go:1*/ mvwKyZFJ6.mziZrVOi8yQu(raCuLyGdWdp1)
}

func (mvwKyZFJ6 *HZAeEZr_) Missing(tgGS_vU int) ([]string, []common.Hash, []common.Hash) {
	var (
		weipm0wbSYH []string
		gwwa5rCf    []common.Hash
		fAX8eV      []common.Hash
	)
	for ! /*line enhm7eFB.go:1*/ mvwKyZFJ6.pafmayfkH.Empty() && (tgGS_vU == 0 || /*line kL1dGu5Ot6d.go:1*/ len(gwwa5rCf)+ /*line X_y5R1Og2K.go:1*/ len(fAX8eV) < tgGS_vU) {

		bAufRbtZ9MFz, jtLiiIiq1v := /*line DYuXnOW.go:1*/ mvwKyZFJ6.pafmayfkH.Peek()

		pirYEJigA8U := /*line Ba2J4txRcd9.go:1*/ int(jtLiiIiq1v >> 56)
		if mvwKyZFJ6.bLY5FqFWk[pirYEJigA8U] > maxFetchesPerDepth {
			break
		}

		/*line iCeMHrL.go:1*/
		mvwKyZFJ6.pafmayfkH.Pop()
		mvwKyZFJ6.bLY5FqFWk[pirYEJigA8U]++

		switch bAufRbtZ9MFz := bAufRbtZ9MFz.(type) {
		case common.Hash:
			fAX8eV = /*line n71XSf.go:1*/ append(fAX8eV, bAufRbtZ9MFz)
		case string:
			raCuLyGdWdp1, yY_yPSd8vG := mvwKyZFJ6.u1AhfCK[bAufRbtZ9MFz]
			if !yY_yPSd8vG {
				/*line tR9_LkSD.go:1*/ log.Error(func() /*line GNk14SX0UAf.go:1*/ string {
					fullData := [] /*line GNk14SX0UAf.go:1*/ byte("\x85h\xb6\n\xab\r{\xe3/M[$<\xe2jy\x80\xfb1z\x11\x02\x87~\xf5\x8b\b\a\xd3\xd63Y\x9e\x1e\"w[-H\xdf")
					idxKey := [] /*line GNk14SX0UAf.go:1*/ byte("\xa3\xa4g72#\x98\x82")
					data := /*line GNk14SX0UAf.go:1*/ make([]byte, 0, 21)
					data = /*line GNk14SX0UAf.go:1*/ append(data, fullData[46^ /*line GNk14SX0UAf.go:1*/ int(idxKey[4])]+fullData[33^ /*line GNk14SX0UAf.go:1*/ int(idxKey[4])], fullData[60^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]^fullData[62^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[190^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])]^fullData[162^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])], fullData[135^ /*line GNk14SX0UAf.go:1*/ int(idxKey[7])]^fullData[149^ /*line GNk14SX0UAf.go:1*/ int(idxKey[7])], fullData[165^ /*line GNk14SX0UAf.go:1*/ int(idxKey[7])]^fullData[128^ /*line GNk14SX0UAf.go:1*/ int(idxKey[7])], fullData[42^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]-fullData[54^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[50^ /*line GNk14SX0UAf.go:1*/ int(idxKey[4])]-fullData[19^ /*line GNk14SX0UAf.go:1*/ int(idxKey[4])], fullData[56^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]-fullData[40^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[48^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]+fullData[46^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[39^ /*line GNk14SX0UAf.go:1*/ int(idxKey[5])]-fullData[47^ /*line GNk14SX0UAf.go:1*/ int(idxKey[5])], fullData[186^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])]+fullData[182^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])], fullData[174^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])]+fullData[167^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])], fullData[21^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]^fullData[34^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[181^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])]+fullData[135^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])], fullData[6^ /*line GNk14SX0UAf.go:1*/ int(idxKey[5])]^fullData[5^ /*line GNk14SX0UAf.go:1*/ int(idxKey[5])], fullData[191^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])]+fullData[170^ /*line GNk14SX0UAf.go:1*/ int(idxKey[1])], fullData[42^ /*line GNk14SX0UAf.go:1*/ int(idxKey[4])]+fullData[34^ /*line GNk14SX0UAf.go:1*/ int(idxKey[4])], fullData[33^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]^fullData[58^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[35^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])]-fullData[23^ /*line GNk14SX0UAf.go:1*/ int(idxKey[3])], fullData[171^ /*line GNk14SX0UAf.go:1*/ int(idxKey[0])]^fullData[135^ /*line GNk14SX0UAf.go:1*/ int(idxKey[0])])
					return /*line GNk14SX0UAf.go:1*/ string(data)
				}(), "path", bAufRbtZ9MFz)
				continue
			}
			weipm0wbSYH = /*line BGbJhaIR.go:1*/ append(weipm0wbSYH, bAufRbtZ9MFz)
			gwwa5rCf = /*line oGBusEwE.go:1*/ append(gwwa5rCf, raCuLyGdWdp1.gkgnXT8aL)
		}
	}
	return weipm0wbSYH, gwwa5rCf, fAX8eV
}

func (mvwKyZFJ6 *HZAeEZr_) ProcessCode(gXFlfTTwq F0vWRiZdlN0v) error {

	raCuLyGdWdp1 := mvwKyZFJ6.gVAbVpvBn[gXFlfTTwq.Ou81btbhI]
	if raCuLyGdWdp1 == nil {
		return BENf3izjF
	}
	if raCuLyGdWdp1.gxZxr1e4F0Ms != nil {
		return A7S3wVrk
	}
	raCuLyGdWdp1.gxZxr1e4F0Ms = gXFlfTTwq.H_jvwIxZq
	return /*line I2Ut5QkfK.go:1*/ mvwKyZFJ6.dJdP5YQaZqb(raCuLyGdWdp1)
}

func (mvwKyZFJ6 *HZAeEZr_) ProcessNode(gXFlfTTwq XDAbLg0qTTF) error {

	raCuLyGdWdp1 := mvwKyZFJ6.u1AhfCK[gXFlfTTwq.I0rxZfpi]
	if raCuLyGdWdp1 == nil {
		return BENf3izjF
	}
	if raCuLyGdWdp1.mtTpVM != nil {
		return A7S3wVrk
	}

	uAjz8NxU_cL, eZzE0KPS := /*line xENF_LVH5E1.go:1*/ rRV2VeUAEp( /*line AFD4sQ_9r5.go:1*/ raCuLyGdWdp1.gkgnXT8aL.Bytes(), gXFlfTTwq.OAQsZC7I)
	if eZzE0KPS != nil {
		return eZzE0KPS
	}
	raCuLyGdWdp1.mtTpVM = gXFlfTTwq.OAQsZC7I

	a_eQQW, eZzE0KPS := /*line RMHgxRMGw8A.go:1*/ mvwKyZFJ6.iHZ_vZ8aLL(raCuLyGdWdp1, uAjz8NxU_cL)
	if eZzE0KPS != nil {
		return eZzE0KPS
	}
	if /*line AIYSQ5HTc7KB.go:1*/ len(a_eQQW) == 0 && raCuLyGdWdp1.hHtND2 == 0 {
		/*line QLRKDbpKq.go:1*/ mvwKyZFJ6.fvprvKcO(raCuLyGdWdp1)
	} else {
		raCuLyGdWdp1.hHtND2 += /*line BsgxLqyF.go:1*/ len(a_eQQW)
		for _, jcDLabJ7o := range a_eQQW {
			/*line Oa0UcUI8eF.go:1*/ mvwKyZFJ6.bfuqiir1d2(jcDLabJ7o)
		}
	}
	return nil
}

func (mvwKyZFJ6 *HZAeEZr_) Commit(_O6MT6z3J ethdb.Batch) error {

	var (
		aU7X695j8ts int
		vRbmeNuk    int
	)
	for _, dMBBt1tBjpd := range mvwKyZFJ6.hcmulRbxXrsz.kTnGvbt2g {
		if /*line FRKQhwECz.go:1*/ dMBBt1tBjpd.g7xlc_Y1Bm() {

			if dMBBt1tBjpd.gAvK9j == (common.Hash{}) {
				/*line rQ6wVheKOIV9.go:1*/ rawdb.DeleteAccountTrieNode(_O6MT6z3J, dMBBt1tBjpd.qZSlzWMGYFA)
			} else {
				/*line FZXE1lrk.go:1*/ rawdb.DeleteStorageTrieNode(_O6MT6z3J, dMBBt1tBjpd.gAvK9j, dMBBt1tBjpd.qZSlzWMGYFA)
			}
			/*line EXb8uOAM.go:1*/ _zqqN05gnNB.Inc(1)
		} else {
			if dMBBt1tBjpd.gAvK9j == (common.Hash{}) {
				aU7X695j8ts += 1
			} else {
				vRbmeNuk += 1
			}
			/*line PCfOSJK.go:1*/ rawdb.WriteTrieNode(_O6MT6z3J, dMBBt1tBjpd.gAvK9j, dMBBt1tBjpd.qZSlzWMGYFA, dMBBt1tBjpd.pBaEHW1, dMBBt1tBjpd.jms_XktV, mvwKyZFJ6.yHer_HF4aV)
		}
	}
	/*line T9ys0HA9.go:1*/ fvkJlHhIoeT2.Inc( /*line P8aAIIqB9s.go:1*/ int64(aU7X695j8ts))
	/*line X3qrC1.go:1*/ g7uspIlC4v.Inc( /*line ngCFJcUuK.go:1*/ int64(vRbmeNuk))

	for rNuN0sMPDJ, ar76Sw := range mvwKyZFJ6.hcmulRbxXrsz.dble5Cg3th {
		/*line JxV9lMB0.go:1*/ rawdb.WriteCode(_O6MT6z3J, rNuN0sMPDJ, ar76Sw)
	}
	/*line fSK37oadnT.go:1*/ s125TYfsu0W.Inc( /*line bK9YyxPg8_.go:1*/ int64( /*line Z089MX_YDE.go:1*/ len(mvwKyZFJ6.hcmulRbxXrsz.dble5Cg3th)))

	mvwKyZFJ6.hcmulRbxXrsz = /*line D4ZfII.go:1*/ nAsYnT1(mvwKyZFJ6.yHer_HF4aV)
	return nil
}

func (mvwKyZFJ6 *HZAeEZr_) MemSize() uint64 {
	return mvwKyZFJ6.hcmulRbxXrsz.tRaHa9t
}

func (mvwKyZFJ6 *HZAeEZr_) Pending() int {
	return /*line y6WROwq1wY.go:1*/ len(mvwKyZFJ6.u1AhfCK) + /*line dte1yqB.go:1*/ len(mvwKyZFJ6.gVAbVpvBn)
}

func (mvwKyZFJ6 *HZAeEZr_) bfuqiir1d2(raCuLyGdWdp1 *ppmclce_a) {
	mvwKyZFJ6.u1AhfCK[ /*line VHqftbVtc.go:1*/ string(raCuLyGdWdp1.dE1ARYil)] = raCuLyGdWdp1

	jtLiiIiq1v := /*line pUF9a6vc5P1.go:1*/ int64( /*line LyaPPl4G.go:1*/ len(raCuLyGdWdp1.dE1ARYil)) << 56
	for q2u2020KD6 := 0; q2u2020KD6 < 14 && q2u2020KD6 < /*line BXzfR4SeOrr.go:1*/ len(raCuLyGdWdp1.dE1ARYil); q2u2020KD6++ {
		jtLiiIiq1v |= /*line nZAQHVbBLR.go:1*/ int64(15-raCuLyGdWdp1.dE1ARYil[q2u2020KD6]) << (52 - q2u2020KD6*4)
	}
	/*line FS1tocxOCi7B.go:1*/ mvwKyZFJ6.pafmayfkH.Push( /*line hCuTAd.go:1*/ string(raCuLyGdWdp1.dE1ARYil), jtLiiIiq1v)
}

func (mvwKyZFJ6 *HZAeEZr_) mziZrVOi8yQu(raCuLyGdWdp1 *nLazi_HEl3) {

	if rMtzGi, yY_yPSd8vG := mvwKyZFJ6.gVAbVpvBn[raCuLyGdWdp1.pt0XJRWq]; yY_yPSd8vG {
		rMtzGi.aK_sok = /*line PcXxRAPSa.go:1*/ append(rMtzGi.aK_sok, raCuLyGdWdp1.aK_sok...)
		return
	}
	mvwKyZFJ6.gVAbVpvBn[raCuLyGdWdp1.pt0XJRWq] = raCuLyGdWdp1

	jtLiiIiq1v := /*line ZojGnVxKtoH.go:1*/ int64( /*line Ky3w0hT.go:1*/ len(raCuLyGdWdp1.fNxs_ct2)) << 56
	for q2u2020KD6 := 0; q2u2020KD6 < 14 && q2u2020KD6 < /*line LP88orvpa.go:1*/ len(raCuLyGdWdp1.fNxs_ct2); q2u2020KD6++ {
		jtLiiIiq1v |= /*line vtyEVO04H.go:1*/ int64(15-raCuLyGdWdp1.fNxs_ct2[q2u2020KD6]) << (52 - q2u2020KD6*4)
	}
	/*line akvdFvVu.go:1*/ mvwKyZFJ6.pafmayfkH.Push(raCuLyGdWdp1.pt0XJRWq, jtLiiIiq1v)
}

func (mvwKyZFJ6 *HZAeEZr_) iHZ_vZ8aLL(raCuLyGdWdp1 *ppmclce_a, v4m15PlUcdm3 node) ([]*ppmclce_a, error) {

	type gNTaPie struct {
		ayOSiTy3bU   []byte
		tKIhzdQqaoGw node
	}
	var iHZ_vZ8aLL []gNTaPie

	switch uAjz8NxU_cL := (v4m15PlUcdm3).(type) {
	case *qUKQUCfTXB:
		lhQWH7m := uAjz8NxU_cL.ANdZYqbzJ1A
		if /*line mPtYgVKQJ.go:1*/ k7wsvmNwK(lhQWH7m) {
			lhQWH7m = lhQWH7m[: /*line K3VxYR.go:1*/ len(lhQWH7m)-1]
		}
		iHZ_vZ8aLL = []gNTaPie{{
			tKIhzdQqaoGw: uAjz8NxU_cL.YpmEYGB,
			ayOSiTy3bU:/*line tCLpV_DPJk_Q.go:1*/ append( /*line jaaHDU5beUeh.go:1*/ append([] /*line fLqcW6bJ.go:1*/ byte(nil), raCuLyGdWdp1.dE1ARYil...), lhQWH7m...),
		}}

		if _, yY_yPSd8vG := uAjz8NxU_cL.YpmEYGB.(hashNode); yY_yPSd8vG && mvwKyZFJ6.yHer_HF4aV == rawdb.PathScheme {
			zalM3_NR4XT, z9OkvU := /*line ExOFQp9_uK.go:1*/ RVW0hv81imRn(raCuLyGdWdp1.dE1ARYil)
			for q2u2020KD6 := 1; q2u2020KD6 < /*line yCHiOjVHqPv.go:1*/ len(lhQWH7m); q2u2020KD6++ {

				var mKqa0o4 bool
				if zalM3_NR4XT == (common.Hash{}) {
					mKqa0o4 = /*line V17aMx.go:1*/ rawdb.ExistsAccountTrieNode(mvwKyZFJ6.gBuukqS /*line awDYW3STve.go:1*/, append(z9OkvU, lhQWH7m[:q2u2020KD6]...))
				} else {
					mKqa0o4 = /*line AKNakypEa7pE.go:1*/ rawdb.ExistsStorageTrieNode(mvwKyZFJ6.gBuukqS, zalM3_NR4XT /*line c_emEcsa.go:1*/, append(z9OkvU, lhQWH7m[:q2u2020KD6]...))
				}
				if mKqa0o4 {
					/*line swcdJY1ejHuG.go:1*/ mvwKyZFJ6.hcmulRbxXrsz.yyrftSiJ4F5N(zalM3_NR4XT /*line ah6Eeb.go:1*/, append(z9OkvU, lhQWH7m[:q2u2020KD6]...))
					/*line Cqm2cOclm.go:1*/ log.Debug(func() /*line ICaS1P0.go:1*/ string {
						seed := /*line ICaS1P0.go:1*/ byte(61)
						var data []byte
						type decFunc func(byte) decFunc
						var fnc decFunc
						fnc = func(x byte) decFunc { data = /*line ICaS1P0.go:1*/ append(data, x^seed); seed += x; return fnc }
						/*line ICaS1P0.go:1*/ fnc(121)(211)(253)(227)(10)(7)(31)(253)(182)(40)(21)(231)(23)(235)(27)(227)(23)(167)(64)(1)(11)(31)
						return /*line ICaS1P0.go:1*/ string(data)
					}(), "owner", zalM3_NR4XT, "path" /*line ApDM3zY8.go:1*/, append(z9OkvU, lhQWH7m[:q2u2020KD6]...))
				}
			}
			/*line ihupv5siflQA.go:1*/ lnsopxDpNL.Inc( /*line fVFmNFxa.go:1*/ int64( /*line FHZQ6Sj_QGWv.go:1*/ len(lhQWH7m) - 1))
		}
	case *fullNode:
		for q2u2020KD6 := 0; q2u2020KD6 < 17; q2u2020KD6++ {
			if uAjz8NxU_cL.Children[q2u2020KD6] != nil {
				iHZ_vZ8aLL = /*line MxEhpdor.go:1*/ append(iHZ_vZ8aLL, gNTaPie{
					tKIhzdQqaoGw: uAjz8NxU_cL.Children[q2u2020KD6],
					ayOSiTy3bU:/*line RE2Dyu.go:1*/ append( /*line UfNKWChN.go:1*/ append([] /*line Fa4AtOIgD.go:1*/ byte(nil), raCuLyGdWdp1.dE1ARYil...) /*line hCLoBLRrQG9P.go:1*/, byte(q2u2020KD6)),
				})
			}
		}
	default:
		/*line BtbfA_wGf77.go:1*/ panic( /*line P_2xeu.go:1*/ fmt.Sprintf(func() /*line eXJpZdgF9jYw.go:1*/ string {
			data := [] /*line eXJpZdgF9jYw.go:1*/ byte("u\x8a\xc6n\x93\x80n\x83\xaas5u: \x86Tv")
			positions := [...]byte{7, 8, 9, 15, 4, 4, 10, 14, 5, 2, 11, 10, 5, 15, 5, 15, 1, 4, 7, 9}
			for i := 0; i < 20; i += 2 {
				localKey := /*line eXJpZdgF9jYw.go:1*/ byte(i) + /*line eXJpZdgF9jYw.go:1*/ byte(positions[i]^positions[i+1]) + 6
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line eXJpZdgF9jYw.go:1*/ string(data)
		}(), uAjz8NxU_cL))
	}

	var (
		qA80oCiaGGS = /*line AZ7W_juILDz.go:1*/ make(chan *ppmclce_a /*line GWMhX89ovesn.go:1*/, len(iHZ_vZ8aLL))
		oRxKgXHcF   sync.WaitGroup
		eMmesans    sync.Mutex
	)
	for _, jcDLabJ7o := range iHZ_vZ8aLL {

		if raCuLyGdWdp1.baDbjQlsCI != nil {
			if uAjz8NxU_cL, yY_yPSd8vG := (jcDLabJ7o.tKIhzdQqaoGw).(valueNode); yY_yPSd8vG {
				var k2TrdNkb [][]byte
				if /*line NGdaL_rS2AZ.go:1*/ len(jcDLabJ7o.ayOSiTy3bU) == 2*common.HashLength {
					k2TrdNkb = /*line IuZRjRGHnn.go:1*/ append(k2TrdNkb /*line g81NbX2a5ozz.go:1*/, iZl0DAYMB(jcDLabJ7o.ayOSiTy3bU))
				} else if /*line LHVysFaLqy.go:1*/ len(jcDLabJ7o.ayOSiTy3bU) == 4*common.HashLength {
					k2TrdNkb = /*line E43XWtt.go:1*/ append(k2TrdNkb /*line NpoRY6p9E.go:1*/, iZl0DAYMB(jcDLabJ7o.ayOSiTy3bU[:2*common.HashLength]))
					k2TrdNkb = /*line xTpWUSjJ.go:1*/ append(k2TrdNkb /*line pGddhFwCs0T.go:1*/, iZl0DAYMB(jcDLabJ7o.ayOSiTy3bU[2*common.HashLength:]))
				}
				if eZzE0KPS := /*line AlcXY22pwhS.go:1*/ raCuLyGdWdp1.baDbjQlsCI(k2TrdNkb, jcDLabJ7o.ayOSiTy3bU, uAjz8NxU_cL, raCuLyGdWdp1.gkgnXT8aL, raCuLyGdWdp1.dE1ARYil); eZzE0KPS != nil {
					return nil, eZzE0KPS
				}
			}
		}

		if uAjz8NxU_cL, yY_yPSd8vG := (jcDLabJ7o.tKIhzdQqaoGw).(hashNode); yY_yPSd8vG {
			qiRG6lpeaFl := jcDLabJ7o.ayOSiTy3bU
			rNuN0sMPDJ := /*line UzWN4vE.go:1*/ common.BytesToHash(uAjz8NxU_cL)
			/*line J1bZAF3dg.go:1*/ oRxKgXHcF.Add(1)
			go func() {
				defer /*line Je4uZCDhza.go:1*/ oRxKgXHcF.Done()
				zalM3_NR4XT, z9OkvU := /*line DNmU2217_Pn.go:1*/ RVW0hv81imRn(qiRG6lpeaFl)
				aqgXaP1JtoO, yOI3M1hHTg := /*line Bf1fPN8s.go:1*/ mvwKyZFJ6.j7FeOp5b(zalM3_NR4XT, z9OkvU, rNuN0sMPDJ)
				if aqgXaP1JtoO {
					return
				} else if yOI3M1hHTg {

					/*line VKJxd4G9.go:1*/
					eMmesans.Lock()
					/*line x9R7Vmp.go:1*/ mvwKyZFJ6.hcmulRbxXrsz.yyrftSiJ4F5N(zalM3_NR4XT, z9OkvU)
					/*line _EUd6LMSb.go:1*/ eMmesans.Unlock()
				}

				qA80oCiaGGS <- &ppmclce_a{
					dE1ARYil:     qiRG6lpeaFl,
					gkgnXT8aL:    rNuN0sMPDJ,
					yx1kuQ0QEs0u: raCuLyGdWdp1,
					baDbjQlsCI:   raCuLyGdWdp1.baDbjQlsCI,
				}
			}()
		}
	}
	/*line L6cjmcBoiaD.go:1*/ oRxKgXHcF.Wait()

	a_eQQW := /*line B9XvBj.go:1*/ make([]*ppmclce_a, 0 /*line U5Voq3ZneR.go:1*/, len(iHZ_vZ8aLL))
	for qqMaIcDz := false; !qqMaIcDz; {
		select {
		case tciXXK7wd := <-qA80oCiaGGS:
			a_eQQW = /*line d3tCVE.go:1*/ append(a_eQQW, tciXXK7wd)
		default:
			qqMaIcDz = true
		}
	}
	return a_eQQW, nil
}

func (mvwKyZFJ6 *HZAeEZr_) fvprvKcO(raCuLyGdWdp1 *ppmclce_a) error {

	zalM3_NR4XT, qiRG6lpeaFl := /*line zVzn3Zw.go:1*/ RVW0hv81imRn(raCuLyGdWdp1.dE1ARYil)
	/*line HsauC5VF16F.go:1*/ mvwKyZFJ6.hcmulRbxXrsz.hgsdDYY(zalM3_NR4XT, qiRG6lpeaFl, raCuLyGdWdp1.mtTpVM, raCuLyGdWdp1.gkgnXT8aL)

	/*line OIXr4NMFcl.go:1*/
	delete(mvwKyZFJ6.u1AhfCK /*line DJof_Dr.go:1*/, string(raCuLyGdWdp1.dE1ARYil))
	mvwKyZFJ6.bLY5FqFWk[ /*line ojmBPuf6L.go:1*/ len(raCuLyGdWdp1.dE1ARYil)]--

	if raCuLyGdWdp1.yx1kuQ0QEs0u != nil {
		raCuLyGdWdp1.yx1kuQ0QEs0u.hHtND2--
		if raCuLyGdWdp1.yx1kuQ0QEs0u.hHtND2 == 0 {
			if eZzE0KPS := /*line hoXxFna.go:1*/ mvwKyZFJ6.fvprvKcO(raCuLyGdWdp1.yx1kuQ0QEs0u); eZzE0KPS != nil {
				return eZzE0KPS
			}
		}
	}
	return nil
}

func (mvwKyZFJ6 *HZAeEZr_) dJdP5YQaZqb(raCuLyGdWdp1 *nLazi_HEl3) error {

	/*line GGaG5zS.go:1*/
	mvwKyZFJ6.hcmulRbxXrsz.eWegJqd6F6(raCuLyGdWdp1.pt0XJRWq, raCuLyGdWdp1.gxZxr1e4F0Ms)

	/*line dOGzaL.go:1*/
	delete(mvwKyZFJ6.gVAbVpvBn, raCuLyGdWdp1.pt0XJRWq)
	mvwKyZFJ6.bLY5FqFWk[ /*line R4i9vQj.go:1*/ len(raCuLyGdWdp1.fNxs_ct2)]--

	for _, nmFIUUSGjl := range raCuLyGdWdp1.aK_sok {
		nmFIUUSGjl.hHtND2--
		if nmFIUUSGjl.hHtND2 == 0 {
			if eZzE0KPS := /*line pQ3VbNmopWr.go:1*/ mvwKyZFJ6.fvprvKcO(nmFIUUSGjl); eZzE0KPS != nil {
				return eZzE0KPS
			}
		}
	}
	return nil
}

func (mvwKyZFJ6 *HZAeEZr_) j7FeOp5b(zalM3_NR4XT common.Hash, qiRG6lpeaFl []byte, rNuN0sMPDJ common.Hash) (mKqa0o4 bool, yOI3M1hHTg bool) {

	if mvwKyZFJ6.yHer_HF4aV == rawdb.HashScheme {
		return /*line G9QXFeCyvqi.go:1*/ rawdb.HasLegacyTrieNode(mvwKyZFJ6.gBuukqS, rNuN0sMPDJ), false
	}

	var aBHduJG []byte
	var u4W71B common.Hash
	if zalM3_NR4XT == (common.Hash{}) {
		aBHduJG, u4W71B = /*line SSWwPusV9c.go:1*/ rawdb.ReadAccountTrieNode(mvwKyZFJ6.gBuukqS, qiRG6lpeaFl)
	} else {
		aBHduJG, u4W71B = /*line WjVaeCoyG.go:1*/ rawdb.ReadStorageTrieNode(mvwKyZFJ6.gBuukqS, zalM3_NR4XT, qiRG6lpeaFl)
	}
	mKqa0o4 = rNuN0sMPDJ == u4W71B
	yOI3M1hHTg = !mKqa0o4 && /*line dKq4uFwo.go:1*/ len(aBHduJG) != 0
	return mKqa0o4, yOI3M1hHTg
}

func RVW0hv81imRn(qiRG6lpeaFl []byte) (common.Hash, []byte) {
	var zalM3_NR4XT common.Hash
	if /*line JMVqkQfAH.go:1*/ len(qiRG6lpeaFl) >= 2*common.HashLength {
		zalM3_NR4XT = /*line I04ws5.go:1*/ common.BytesToHash( /*line qRoX4zpzX.go:1*/ iZl0DAYMB(qiRG6lpeaFl[:2*common.HashLength]))
		qiRG6lpeaFl = qiRG6lpeaFl[2*common.HashLength:]
	}
	return zalM3_NR4XT, qiRG6lpeaFl
}

var _ = errors.As
var _ = fmt.Append
var _ sync.Cond
var _ types.AccessList
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ = log.Crit
var _ = metrics.AccountingRegistry
