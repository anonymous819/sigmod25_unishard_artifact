//line :1
package snapshot

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	types "unishard/evm/types"

	trie "unishard/evm/trie"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/triedb"
)

var (
	inviux = 128

	lZKEBSD = 1024

	j0_A4Kp3 = /*line qXArWCYCTiM.go:1*/ errors.New(func() /*line VVzIBwz.go:1*/ string {
		fullData := [] /*line VVzIBwz.go:1*/ byte("\v.\xa2\x01B;\x86\xa6<~\xf5\xc1\xe1Τ]b\fxb\xd0r-\xd0")
		idxKey := [] /*line VVzIBwz.go:1*/ byte("\x86'")
		data := /*line VVzIBwz.go:1*/ make([]byte, 0, 13)
		data = /*line VVzIBwz.go:1*/ append(data, fullData[53^ /*line VVzIBwz.go:1*/ int(idxKey[1])]-fullData[39^ /*line VVzIBwz.go:1*/ int(idxKey[1])], fullData[34^ /*line VVzIBwz.go:1*/ int(idxKey[1])]+fullData[38^ /*line VVzIBwz.go:1*/ int(idxKey[1])], fullData[45^ /*line VVzIBwz.go:1*/ int(idxKey[1])]+fullData[46^ /*line VVzIBwz.go:1*/ int(idxKey[1])], fullData[147^ /*line VVzIBwz.go:1*/ int(idxKey[0])]^fullData[133^ /*line VVzIBwz.go:1*/ int(idxKey[0])], fullData[137^ /*line VVzIBwz.go:1*/ int(idxKey[0])]+fullData[151^ /*line VVzIBwz.go:1*/ int(idxKey[0])], fullData[145^ /*line VVzIBwz.go:1*/ int(idxKey[0])]-fullData[149^ /*line VVzIBwz.go:1*/ int(idxKey[0])], fullData[43^ /*line VVzIBwz.go:1*/ int(idxKey[1])]^fullData[33^ /*line VVzIBwz.go:1*/ int(idxKey[1])], fullData[130^ /*line VVzIBwz.go:1*/ int(idxKey[0])]^fullData[150^ /*line VVzIBwz.go:1*/ int(idxKey[0])], fullData[32^ /*line VVzIBwz.go:1*/ int(idxKey[1])]+fullData[42^ /*line VVzIBwz.go:1*/ int(idxKey[1])], fullData[132^ /*line VVzIBwz.go:1*/ int(idxKey[0])]+fullData[146^ /*line VVzIBwz.go:1*/ int(idxKey[0])], fullData[144^ /*line VVzIBwz.go:1*/ int(idxKey[0])]+fullData[142^ /*line VVzIBwz.go:1*/ int(idxKey[0])], fullData[41^ /*line VVzIBwz.go:1*/ int(idxKey[1])]+fullData[44^ /*line VVzIBwz.go:1*/ int(idxKey[1])])
		return /*line VVzIBwz.go:1*/ string(data)
	}())
)

func gazN04(iYSzoWnzsqd ethdb.KeyValueStore, rhls5wQnNg *triedb.Database, duGwhWx7d5J int, z1BBTN2UX common.Hash) *diskLayer {

	var (
		drjkE0      = &generatorStats{start: /*line m_G5zPYG.go:1*/ time.Now()}
		cDn73kIaDil = /*line DojV7_H5_h.go:1*/ iYSzoWnzsqd.NewBatch()
		nrnBdVLnvxF = []byte{}
	)
	/*line GOQ7Eji.go:1*/ rawdb.WriteSnapshotRoot(cDn73kIaDil, z1BBTN2UX)
	/*line McklJYAacRO.go:1*/ bQ5V3QyAnw(cDn73kIaDil, nrnBdVLnvxF, drjkE0)
	if chyZ8Q := /*line jYEamXy8mtEq.go:1*/ cDn73kIaDil.Write(); chyZ8Q != nil {
		/*line J5zc7GF.go:1*/ log.Crit(func() /*line BRk3XNXr.go:1*/ string {
			seed := /*line BRk3XNXr.go:1*/ byte(224)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line BRk3XNXr.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line BRk3XNXr.go:1*/ fnc(166)(231)(4)(29)(235)(29)(182)(56)(235)(79)(201)(245)(21)(229)(19)(169)(91)(227)(25)(253)(239)(20)(229)(7)(15)(225)(1)(70)(223)(255)(235)(1)(19)(169)(95)(240)(243)(31)(246)(251)
			return /*line BRk3XNXr.go:1*/ string(data)
		}(), "err", chyZ8Q)
	}
	v4faEBKJgI2 := &diskLayer{
		diskdb: iYSzoWnzsqd,
		triedb: rhls5wQnNg,
		root:   z1BBTN2UX,
		cache:/*line iYoVbOY2.go:1*/ fastcache.New(duGwhWx7d5J * 1024 * 1024),
		genMarker: nrnBdVLnvxF,
		genPending:/*line clxUL2E.go:1*/ make(chan struct{}),
		genAbort:/*line I_559RliZ.go:1*/ make(chan chan *generatorStats),
	}
	go /*line _cbsadoXH.go:1*/ v4faEBKJgI2.tliqpCic(drjkE0)
	/*line WJeHNzTRC9.go:1*/ log.Debug(func() /*line Rwonqu4h9.go:1*/ string {
		seed := /*line Rwonqu4h9.go:1*/ byte(92)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line Rwonqu4h9.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line Rwonqu4h9.go:1*/ fnc(15)(31)(235)(7)(8)(164)(91)(237)(17)(241)(1)(27)(225)(27)(170)(83)(226)(7)(21)(247)(29)(237)(239)(26)(225)
		return /*line Rwonqu4h9.go:1*/ string(data)
	}(), "root", z1BBTN2UX)
	return v4faEBKJgI2
}

func bQ5V3QyAnw(lDBEbd ethdb.KeyValueWriter, al5iRGOo []byte, drjkE0 *generatorStats) {

	gUZorXrZXPhE := journalGenerator{
		Done:   al5iRGOo == nil,
		Marker: al5iRGOo,
	}
	if drjkE0 != nil {
		gUZorXrZXPhE.Accounts = drjkE0.accounts
		gUZorXrZXPhE.Slots = drjkE0.slots
		gUZorXrZXPhE.Storage = /*line H7RUq_n.go:1*/ uint64(drjkE0.storage)
	}
	b1jafJL, chyZ8Q := /*line IurOafq7ia_.go:1*/ rlp.EncodeToBytes(gUZorXrZXPhE)
	if chyZ8Q != nil {
		/*line vXDI4FEyNnFg.go:1*/ panic(chyZ8Q)
	}
	var uZhnATB string
	switch {
	case al5iRGOo == nil:
		uZhnATB = "done"
	case /*line IGlWIKwBfqP.go:1*/ bytes.Equal(al5iRGOo, []byte{}):
		uZhnATB = "empty"
	case /*line sTrVSdxzsri.go:1*/ len(al5iRGOo) == common.HashLength:
		uZhnATB = /*line eueyQaKh.go:1*/ fmt.Sprintf("%#x", al5iRGOo)
	default:
		uZhnATB = /*line JxBMca3.go:1*/ fmt.Sprintf("%#x:%#x", al5iRGOo[:common.HashLength], al5iRGOo[common.HashLength:])
	}
	/*line BPMdACeP6aq.go:1*/ log.Debug(func() /*line z39dWm.go:1*/ string {
		fullData := [] /*line z39dWm.go:1*/ byte("\xfc\xb9\xea\x1e\x85\xe3?\xf7\x97\x9c%3c\xf3n\x85\x11\xbb-3\x9d\x9c\xe1\xa1_\x91/\x82\x1cɛvt\x99\xde\x1d\x15\x94o\xbb\xf1\x96\xeaveۆ\x94\xe6Z1\x84>W\x02\xf0l\x81")
		idxKey := [] /*line z39dWm.go:1*/ byte("\x18ak\aƗn-\xb2UbR\x01N\xb0\xd0")
		data := /*line z39dWm.go:1*/ make([]byte, 0, 30)
		data = /*line z39dWm.go:1*/ append(data, fullData[25^ /*line z39dWm.go:1*/ int(idxKey[12])]^fullData[37^ /*line z39dWm.go:1*/ int(idxKey[12])], fullData[184^ /*line z39dWm.go:1*/ int(idxKey[5])]+fullData[186^ /*line z39dWm.go:1*/ int(idxKey[5])], fullData[144^ /*line z39dWm.go:1*/ int(idxKey[8])]+fullData[186^ /*line z39dWm.go:1*/ int(idxKey[8])], fullData[169^ /*line z39dWm.go:1*/ int(idxKey[14])]^fullData[181^ /*line z39dWm.go:1*/ int(idxKey[14])], fullData[128^ /*line z39dWm.go:1*/ int(idxKey[5])]-fullData[132^ /*line z39dWm.go:1*/ int(idxKey[5])], fullData[158^ /*line z39dWm.go:1*/ int(idxKey[8])]+fullData[178^ /*line z39dWm.go:1*/ int(idxKey[8])], fullData[160^ /*line z39dWm.go:1*/ int(idxKey[5])]^fullData[158^ /*line z39dWm.go:1*/ int(idxKey[5])], fullData[182^ /*line z39dWm.go:1*/ int(idxKey[5])]-fullData[133^ /*line z39dWm.go:1*/ int(idxKey[5])], fullData[73^ /*line z39dWm.go:1*/ int(idxKey[1])]+fullData[65^ /*line z39dWm.go:1*/ int(idxKey[1])], fullData[48^ /*line z39dWm.go:1*/ int(idxKey[7])]+fullData[51^ /*line z39dWm.go:1*/ int(idxKey[7])], fullData[72^ /*line z39dWm.go:1*/ int(idxKey[1])]-fullData[126^ /*line z39dWm.go:1*/ int(idxKey[1])], fullData[214^ /*line z39dWm.go:1*/ int(idxKey[4])]^fullData[237^ /*line z39dWm.go:1*/ int(idxKey[4])], fullData[122^ /*line z39dWm.go:1*/ int(idxKey[1])]-fullData[66^ /*line z39dWm.go:1*/ int(idxKey[1])], fullData[12^ /*line z39dWm.go:1*/ int(idxKey[12])]^fullData[21^ /*line z39dWm.go:1*/ int(idxKey[12])], fullData[11^ /*line z39dWm.go:1*/ int(idxKey[3])]+fullData[49^ /*line z39dWm.go:1*/ int(idxKey[3])], fullData[194^ /*line z39dWm.go:1*/ int(idxKey[4])]^fullData[193^ /*line z39dWm.go:1*/ int(idxKey[4])], fullData[114^ /*line z39dWm.go:1*/ int(idxKey[9])]-fullData[100^ /*line z39dWm.go:1*/ int(idxKey[9])], fullData[177^ /*line z39dWm.go:1*/ int(idxKey[14])]+fullData[161^ /*line z39dWm.go:1*/ int(idxKey[14])], fullData[7^ /*line z39dWm.go:1*/ int(idxKey[7])]+fullData[34^ /*line z39dWm.go:1*/ int(idxKey[7])], fullData[109^ /*line z39dWm.go:1*/ int(idxKey[9])]^fullData[86^ /*line z39dWm.go:1*/ int(idxKey[9])], fullData[43^ /*line z39dWm.go:1*/ int(idxKey[0])]+fullData[13^ /*line z39dWm.go:1*/ int(idxKey[0])], fullData[64^ /*line z39dWm.go:1*/ int(idxKey[6])]+fullData[108^ /*line z39dWm.go:1*/ int(idxKey[6])], fullData[133^ /*line z39dWm.go:1*/ int(idxKey[14])]^fullData[186^ /*line z39dWm.go:1*/ int(idxKey[14])], fullData[244^ /*line z39dWm.go:1*/ int(idxKey[4])]+fullData[242^ /*line z39dWm.go:1*/ int(idxKey[4])], fullData[167^ /*line z39dWm.go:1*/ int(idxKey[5])]^fullData[174^ /*line z39dWm.go:1*/ int(idxKey[5])], fullData[84^ /*line z39dWm.go:1*/ int(idxKey[11])]+fullData[89^ /*line z39dWm.go:1*/ int(idxKey[11])], fullData[61^ /*line z39dWm.go:1*/ int(idxKey[0])]-fullData[2^ /*line z39dWm.go:1*/ int(idxKey[0])], fullData[119^ /*line z39dWm.go:1*/ int(idxKey[2])]^fullData[77^ /*line z39dWm.go:1*/ int(idxKey[2])], fullData[208^ /*line z39dWm.go:1*/ int(idxKey[4])]-fullData[200^ /*line z39dWm.go:1*/ int(idxKey[4])])
		return /*line z39dWm.go:1*/ string(data)
	}(), func() /*line HlmwqZtkwB8.go:1*/ string {
		key := [] /*line HlmwqZtkwB8.go:1*/ byte("-\t\xbd\xc6,D%\x82")
		data := [] /*line HlmwqZtkwB8.go:1*/ byte("\x9d{,-\x9e\xa9\x98\xf5")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line HlmwqZtkwB8.go:1*/ string(data)
	}(), uZhnATB)
	/*line ueahJ84NM.go:1*/ rawdb.WriteSnapshotGenerator(lDBEbd, b1jafJL)
}

type sU5j5p struct {
	rACLehqaKtw  [][]byte
	fJRwxa3TtWCJ [][]byte
	zvUxtDgF     bool
	cMH_V88pSut  bool
	wOXUSabjPWCu error
	ahjfa5       *trie.Trie
}

func (hftNwnRF *sU5j5p) aIRqDs() bool {
	return hftNwnRF.wOXUSabjPWCu == nil
}

func (hftNwnRF *sU5j5p) h6VFlt6g() []byte {
	var h6VFlt6g []byte
	if /*line HJK1bBld85.go:1*/ len(hftNwnRF.rACLehqaKtw) > 0 {
		h6VFlt6g = hftNwnRF.rACLehqaKtw[ /*line J64lsX.go:1*/ len(hftNwnRF.rACLehqaKtw)-1]
	}
	return h6VFlt6g
}

func (hftNwnRF *sU5j5p) q9r2hg5(kGNYh2Kjh func(nVUwQz8Zi []byte, jIbih8achE []byte) error) error {
	for fSpMhzHR := 0; fSpMhzHR < /*line za3OHIJBSa.go:1*/ len(hftNwnRF.rACLehqaKtw); fSpMhzHR++ {
		nVUwQz8Zi, jIbih8achE := hftNwnRF.rACLehqaKtw[fSpMhzHR], hftNwnRF.fJRwxa3TtWCJ[fSpMhzHR]
		if chyZ8Q := /*line brG5pdNjWeY.go:1*/ kGNYh2Kjh(nVUwQz8Zi, jIbih8achE); chyZ8Q != nil {
			return chyZ8Q
		}
	}
	return nil
}

func (hbsnnIi *diskLayer) dRlH06yalc(dnam11P *qo43sXy, maaFGYOvd *trie.ID, sjBQgeKoXl []byte, eWW3WCQF string, ijcATHiDv []byte, gkQugurnMa int, hld1IO func([]byte) ([]byte, error)) (*sU5j5p, error) {
	var (
		glD6VIl86    [][]byte
		xELkOUKoeEz  [][]byte
		dc2CEa7W     = /*line DZClF1kH.go:1*/ rawdb.NewMemoryDatabase()
		maGwUs       = false
		lEtpKFfaKmLF = /*line oGGy11O.go:1*/ dnam11P.ahOoDUdo(eWW3WCQF)
		lFRN2MXQc    = /*line SnmN4Evaq.go:1*/ time.Now()
		mj7OGjax2GH  = /*line GsQ7ylr12a.go:1*/ append(sjBQgeKoXl, ijcATHiDv...)
	)
	for /*line DhznooPfoT8.go:1*/ lEtpKFfaKmLF.Next() {

		nVUwQz8Zi := /*line HbbhetA.go:1*/ lEtpKFfaKmLF.Key()
		if /*line YL8gD2s9ml.go:1*/ bytes.Compare(nVUwQz8Zi, mj7OGjax2GH) < 0 {
			return nil /*line ErPacNWl.go:1*/, errors.New(func() /*line ltIBpY0yW.go:1*/ string {
				fullData := [] /*line ltIBpY0yW.go:1*/ byte("\xb8%\xb9\xd9\xfe\xafYV/eq \x86:\xc5\xc2\xe4qg\xac\xd8\xdfs@\xf3v\xe110B\x88}\x8f\x02\x99\xe8\xfe\xef.\x13\x14\xe3MA\x86N\xe4\n6EPZ")
				idxKey := [] /*line ltIBpY0yW.go:1*/ byte("u\xed\x18\xc2\xd5")
				data := /*line ltIBpY0yW.go:1*/ make([]byte, 0, 27)
				data = /*line ltIBpY0yW.go:1*/ append(data, fullData[240^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])]-fullData[249^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])], fullData[234^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]-fullData[206^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[81^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]-fullData[107^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[110^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]+fullData[105^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[71^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]-fullData[101^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[196^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]+fullData[225^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[94^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]^fullData[116^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[249^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]-fullData[237^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[192^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])]-fullData[204^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])], fullData[199^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])]^fullData[242^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])], fullData[93^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]-fullData[112^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[238^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]+fullData[207^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[85^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]-fullData[83^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[239^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]-fullData[220^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[254^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]^fullData[227^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[237^ /*line ltIBpY0yW.go:1*/ int(idxKey[3])]^fullData[203^ /*line ltIBpY0yW.go:1*/ int(idxKey[3])], fullData[111^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]-fullData[99^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[230^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])]-fullData[216^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])], fullData[221^ /*line ltIBpY0yW.go:1*/ int(idxKey[3])]+fullData[218^ /*line ltIBpY0yW.go:1*/ int(idxKey[3])], fullData[125^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]+fullData[98^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[127^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]+fullData[84^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[226^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])]-fullData[235^ /*line ltIBpY0yW.go:1*/ int(idxKey[1])], fullData[229^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])]^fullData[200^ /*line ltIBpY0yW.go:1*/ int(idxKey[4])], fullData[95^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]-fullData[91^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[113^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])]+fullData[100^ /*line ltIBpY0yW.go:1*/ int(idxKey[0])], fullData[19^ /*line ltIBpY0yW.go:1*/ int(idxKey[2])]+fullData[53^ /*line ltIBpY0yW.go:1*/ int(idxKey[2])])
				return /*line ltIBpY0yW.go:1*/ string(data)
			}())
		}

		if ! /*line vPxXeXTMVd0.go:1*/ bytes.Equal(nVUwQz8Zi[: /*line lNA3F_tjc9I_.go:1*/ len(sjBQgeKoXl)], sjBQgeKoXl) {
			/*line H1UD1H.go:1*/ lEtpKFfaKmLF.Hold()
			break
		}

		if /*line YZPO7y.go:1*/ len(glD6VIl86) == gkQugurnMa {
			/*line GTtuY9X9.go:1*/ lEtpKFfaKmLF.Hold()
			maGwUs = true
			break
		}
		glD6VIl86 = /*line JJB7aq6i.go:1*/ append(glD6VIl86 /*line Gqc9hp09F.go:1*/, common.CopyBytes(nVUwQz8Zi[ /*line VC2GmBhs.go:1*/ len(sjBQgeKoXl):]))

		if hld1IO == nil {
			xELkOUKoeEz = /*line WjhotyjQUQV.go:1*/ append(xELkOUKoeEz /*line irTz786H5.go:1*/, common.CopyBytes( /*line hvQt89nAFH.go:1*/ lEtpKFfaKmLF.Value()))
		} else {
			jIbih8achE, chyZ8Q := /*line K3QQFpHN.go:1*/ hld1IO( /*line _ZljYuv7t6y.go:1*/ lEtpKFfaKmLF.Value())
			if chyZ8Q != nil {

				xELkOUKoeEz = /*line veW6KK0zeR.go:1*/ append(xELkOUKoeEz /*line SjQJtpI.go:1*/, common.CopyBytes( /*line pRqMYj.go:1*/ lEtpKFfaKmLF.Value()))
				/*line HZII_9y.go:1*/ log.Error(func() /*line _KySEE.go:1*/ string {
					data := /*line _KySEE.go:1*/ make([]byte, 0, 37)
					i := 6
					decryptKey := 10
					for counter := 0; i != 16; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 1:
							data = /*line _KySEE.go:1*/ append(data, 89)
							i = 3
						case 15:
							i = 14
							data = /*line _KySEE.go:1*/ append(data, 241)
						case 2:
							data = /*line _KySEE.go:1*/ append(data, 64)
							i = 13
						case 6:
							i = 11
							data = /*line _KySEE.go:1*/ append(data, "\x1d9>"...,
							)
						case 0:
							i = 8
							data = /*line _KySEE.go:1*/ append(data, "&36"...,
							)
						case 9:
							i = 12
							data = /*line _KySEE.go:1*/ append(data, ";\xe8"...,
							)
						case 13:
							i = 4
							data = /*line _KySEE.go:1*/ append(data, "MI"...,
							)
						case 12:
							i = 7
							data = /*line _KySEE.go:1*/ append(data, 38)
						case 18:
							data = /*line _KySEE.go:1*/ append(data, "N\x00"...,
							)
							i = 2
						case 8:
							i = 17
							data = /*line _KySEE.go:1*/ append(data, "0C\xf0"...,
							)
						case 4:
							i = 9
							data = /*line _KySEE.go:1*/ append(data, "R>L"...,
							)
						case 10:
							i = 16
							for y := range data {
								data[y] = data[y] + /*line _KySEE.go:1*/ byte(decryptKey^y)
							}
						case 5:
							data = /*line _KySEE.go:1*/ append(data, ".\xea["...,
							)
							i = 1
						case 7:
							i = 0
							data = /*line _KySEE.go:1*/ append(data, 41)
						case 3:
							i = 10
							data = /*line _KySEE.go:1*/ append(data, "iW"...,
							)
						case 17:
							data = /*line _KySEE.go:1*/ append(data, "@B,@"...,
							)
							i = 5
						case 14:
							i = 18
							data = /*line _KySEE.go:1*/ append(data, 70)
						case 11:
							data = /*line _KySEE.go:1*/ append(data, "B88"...,
							)
							i = 15
						}
					}
					return /*line _KySEE.go:1*/ string(data)
				}(), "err", chyZ8Q)
			} else {
				xELkOUKoeEz = /*line Oi3YtJ0ZMW.go:1*/ append(xELkOUKoeEz, jIbih8achE)
			}
		}
	}

	if eWW3WCQF == snapStorage {
		/*line DcjLJa.go:1*/ kCr_5N_kcGi.Inc( /*line jIYgwFZmGS.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
	} else {
		/*line lVJjb4Yw.go:1*/ rfeXPR.Inc( /*line GBMaGv904u.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
	}
	defer func( /*line Uk7ll18vTL.go:1*/ lFRN2MXQc time.Time) {
		if eWW3WCQF == snapStorage {
			/*line ynog0aECbw.go:1*/ g3EaYq.Inc( /*line Z5vCov62.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
		} else {
			/*line JUtPWvg.go:1*/ cauaAYaAEFtV.Inc( /*line zJLiS6bdKM6.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
		}
	}( /*line GCiUdjP.go:1*/ time.Now())

	z1BBTN2UX := maaFGYOvd.Root
	if ijcATHiDv == nil && !maGwUs {
		t0OEuB09q1 := /*line GQtJaVjJ.go:1*/ trie.KYaYz2rvh7Rt(nil)
		for fSpMhzHR, nVUwQz8Zi := range glD6VIl86 {
			if chyZ8Q := /*line rBgptL.go:1*/ t0OEuB09q1.Update(nVUwQz8Zi, xELkOUKoeEz[fSpMhzHR]); chyZ8Q != nil {
				return nil, chyZ8Q
			}
		}
		if iL1CHTcbRt := /*line pgKcO5lMQv.go:1*/ t0OEuB09q1.Hash(); iL1CHTcbRt != z1BBTN2UX {
			return &sU5j5p{
				rACLehqaKtw:  glD6VIl86,
				fJRwxa3TtWCJ: xELkOUKoeEz,
				wOXUSabjPWCu: /*line bLk_PcXZ1.go:1*/ fmt.Errorf(func() /*line KwZCK23.go:1*/ string {
					fullData := [] /*line KwZCK23.go:1*/ byte("\x12\xe6\xff\xbb\xccu5\x10\x16\x92\b\x89Zg\xf6\xff\x1e\xe8\xc6<\xf0{\xe36qT\xaf\xcc\xcaMz\x8dע]\x87\xe5|D\x15\xeb[\xd6>,\xcb\xff\xfd\x8a\x11\x90:_\xa45\xd3\xf4B")
					idxKey := [] /*line KwZCK23.go:1*/ byte("j\xe6")
					data := /*line KwZCK23.go:1*/ make([]byte, 0, 30)
					data = /*line KwZCK23.go:1*/ append(data, fullData[73^ /*line KwZCK23.go:1*/ int(idxKey[0])]+fullData[126^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[200^ /*line KwZCK23.go:1*/ int(idxKey[1])]^fullData[249^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[235^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[236^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[199^ /*line KwZCK23.go:1*/ int(idxKey[1])]^fullData[253^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[67^ /*line KwZCK23.go:1*/ int(idxKey[0])]-fullData[82^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[125^ /*line KwZCK23.go:1*/ int(idxKey[0])]^fullData[98^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[193^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[196^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[201^ /*line KwZCK23.go:1*/ int(idxKey[1])]^fullData[239^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[214^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[194^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[111^ /*line KwZCK23.go:1*/ int(idxKey[0])]+fullData[104^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[232^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[192^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[119^ /*line KwZCK23.go:1*/ int(idxKey[0])]+fullData[93^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[245^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[202^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[115^ /*line KwZCK23.go:1*/ int(idxKey[0])]^fullData[92^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[101^ /*line KwZCK23.go:1*/ int(idxKey[0])]^fullData[97^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[64^ /*line KwZCK23.go:1*/ int(idxKey[0])]-fullData[114^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[248^ /*line KwZCK23.go:1*/ int(idxKey[1])]^fullData[234^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[124^ /*line KwZCK23.go:1*/ int(idxKey[0])]+fullData[83^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[79^ /*line KwZCK23.go:1*/ int(idxKey[0])]^fullData[94^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[112^ /*line KwZCK23.go:1*/ int(idxKey[0])]^fullData[74^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[206^ /*line KwZCK23.go:1*/ int(idxKey[1])]-fullData[203^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[229^ /*line KwZCK23.go:1*/ int(idxKey[1])]^fullData[226^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[107^ /*line KwZCK23.go:1*/ int(idxKey[0])]+fullData[127^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[211^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[250^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[89^ /*line KwZCK23.go:1*/ int(idxKey[0])]-fullData[120^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[205^ /*line KwZCK23.go:1*/ int(idxKey[1])]-fullData[246^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[108^ /*line KwZCK23.go:1*/ int(idxKey[0])]^fullData[109^ /*line KwZCK23.go:1*/ int(idxKey[0])], fullData[230^ /*line KwZCK23.go:1*/ int(idxKey[1])]+fullData[215^ /*line KwZCK23.go:1*/ int(idxKey[1])], fullData[247^ /*line KwZCK23.go:1*/ int(idxKey[1])]^fullData[212^ /*line KwZCK23.go:1*/ int(idxKey[1])])
					return /*line KwZCK23.go:1*/ string(data)
				}(), iL1CHTcbRt, z1BBTN2UX),
			}, nil
		}
		return &sU5j5p{rACLehqaKtw: glD6VIl86, fJRwxa3TtWCJ: xELkOUKoeEz}, nil
	}

	bFlDuoT, chyZ8Q := /*line F6jomQdx.go:1*/ trie.RJaQ3esB(maaFGYOvd, hbsnnIi.triedb)
	if chyZ8Q != nil {
		/*line jrMbFxPbvOp.go:1*/ dnam11P.g3wtYhVIxH4.Log(func() /*line i1phq12usJip.go:1*/ string {
			data := [] /*line i1phq12usJip.go:1*/ byte("T1i\xc1\x10m\x7f!\xcf\xe3),\xeb s\xf5\x17\xc6e\xeesnapsh\x99\x18\xe5\xfdnx po\x17s\x1d\xa1")
			positions := [...]byte{6, 15, 31, 34, 12, 19, 3, 38, 1, 26, 29, 15, 11, 31, 37, 6, 34, 4, 4, 29, 9, 15, 35, 15, 38, 11, 28, 10, 1, 19, 3, 8, 16, 3, 7, 17, 27, 7, 10, 8}
			for i := 0; i < 40; i += 2 {
				localKey := /*line i1phq12usJip.go:1*/ byte(i) + /*line i1phq12usJip.go:1*/ byte(positions[i]^positions[i+1]) + 27
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line i1phq12usJip.go:1*/ string(data)
		}(), hbsnnIi.root, hbsnnIi.genMarker)
		return nil, j0_A4Kp3
	}

	if ijcATHiDv == nil {
		ijcATHiDv = /*line B8QUvZS.go:1*/ common.Hash{}.Bytes()
	}
	if chyZ8Q := /*line GcrwRMyw.go:1*/ bFlDuoT.Prove(ijcATHiDv, dc2CEa7W); chyZ8Q != nil {
		/*line YqgMjlIpVe4o.go:1*/ log.Debug(func() /*line _trK21AGIR.go:1*/ string {
			fullData := [] /*line _trK21AGIR.go:1*/ byte("(\xd5X\xc2H\x87\xf8g\x18D\xa2\xb8\xc8\x1e3)\x8e1\xae\xa1`HQ:V.\x10\x92\x17\xb1\xd5\xdfF*\xc5\xe7\xbe\xdey\x99L-")
			idxKey := [] /*line _trK21AGIR.go:1*/ byte("\x89\x902\x80\xe1\xa9G\xee\x02\x01Q\x1b\xb7\xa4\x02M")
			data := /*line _trK21AGIR.go:1*/ make([]byte, 0, 22)
			data = /*line _trK21AGIR.go:1*/ append(data, fullData[18^ /*line _trK21AGIR.go:1*/ int(idxKey[8])]^fullData[14^ /*line _trK21AGIR.go:1*/ int(idxKey[8])], fullData[23^ /*line _trK21AGIR.go:1*/ int(idxKey[9])]+fullData[27^ /*line _trK21AGIR.go:1*/ int(idxKey[9])], fullData[19^ /*line _trK21AGIR.go:1*/ int(idxKey[14])]^fullData[0^ /*line _trK21AGIR.go:1*/ int(idxKey[14])], fullData[141^ /*line _trK21AGIR.go:1*/ int(idxKey[5])]+fullData[187^ /*line _trK21AGIR.go:1*/ int(idxKey[5])], fullData[249^ /*line _trK21AGIR.go:1*/ int(idxKey[4])]^fullData[239^ /*line _trK21AGIR.go:1*/ int(idxKey[4])], fullData[3^ /*line _trK21AGIR.go:1*/ int(idxKey[14])]^fullData[31^ /*line _trK21AGIR.go:1*/ int(idxKey[14])], fullData[148^ /*line _trK21AGIR.go:1*/ int(idxKey[1])]-fullData[144^ /*line _trK21AGIR.go:1*/ int(idxKey[1])], fullData[86^ /*line _trK21AGIR.go:1*/ int(idxKey[15])]-fullData[64^ /*line _trK21AGIR.go:1*/ int(idxKey[15])], fullData[167^ /*line _trK21AGIR.go:1*/ int(idxKey[3])]-fullData[161^ /*line _trK21AGIR.go:1*/ int(idxKey[3])], fullData[172^ /*line _trK21AGIR.go:1*/ int(idxKey[13])]-fullData[162^ /*line _trK21AGIR.go:1*/ int(idxKey[13])], fullData[7^ /*line _trK21AGIR.go:1*/ int(idxKey[11])]^fullData[28^ /*line _trK21AGIR.go:1*/ int(idxKey[11])], fullData[84^ /*line _trK21AGIR.go:1*/ int(idxKey[15])]+fullData[68^ /*line _trK21AGIR.go:1*/ int(idxKey[15])], fullData[66^ /*line _trK21AGIR.go:1*/ int(idxKey[15])]^fullData[109^ /*line _trK21AGIR.go:1*/ int(idxKey[15])], fullData[142^ /*line _trK21AGIR.go:1*/ int(idxKey[1])]+fullData[131^ /*line _trK21AGIR.go:1*/ int(idxKey[1])], fullData[178^ /*line _trK21AGIR.go:1*/ int(idxKey[1])]-fullData[132^ /*line _trK21AGIR.go:1*/ int(idxKey[1])], fullData[24^ /*line _trK21AGIR.go:1*/ int(idxKey[11])]-fullData[17^ /*line _trK21AGIR.go:1*/ int(idxKey[11])], fullData[21^ /*line _trK21AGIR.go:1*/ int(idxKey[8])]^fullData[23^ /*line _trK21AGIR.go:1*/ int(idxKey[8])], fullData[141^ /*line _trK21AGIR.go:1*/ int(idxKey[13])]^fullData[140^ /*line _trK21AGIR.go:1*/ int(idxKey[13])], fullData[161^ /*line _trK21AGIR.go:1*/ int(idxKey[13])]+fullData[135^ /*line _trK21AGIR.go:1*/ int(idxKey[13])], fullData[29^ /*line _trK21AGIR.go:1*/ int(idxKey[8])]^fullData[9^ /*line _trK21AGIR.go:1*/ int(idxKey[8])], fullData[39^ /*line _trK21AGIR.go:1*/ int(idxKey[8])]-fullData[36^ /*line _trK21AGIR.go:1*/ int(idxKey[8])])
			return /*line _trK21AGIR.go:1*/ string(data)
		}(), "kind", eWW3WCQF, "origin", ijcATHiDv, "err", chyZ8Q)
		return &sU5j5p{
			rACLehqaKtw:  glD6VIl86,
			fJRwxa3TtWCJ: xELkOUKoeEz,
			zvUxtDgF:     maGwUs,
			wOXUSabjPWCu: chyZ8Q,
			ahjfa5:       bFlDuoT,
		}, nil
	}
	if /*line Naf3VzAO8tU.go:1*/ len(glD6VIl86) > 0 {
		if chyZ8Q := /*line OQ7Y9Htd1eF.go:1*/ bFlDuoT.Prove(glD6VIl86[ /*line qPahaAEtye1Y.go:1*/ len(glD6VIl86)-1], dc2CEa7W); chyZ8Q != nil {
			/*line b6aaaNCBU.go:1*/ log.Debug(func() /*line Pimv8z.go:1*/ string {
				seed := /*line Pimv8z.go:1*/ byte(13)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line Pimv8z.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line Pimv8z.go:1*/ fnc(75)(57)(248)(229)(11)(29)(182)(56)(235)(79)(206)(254)(229)(25)(237)(85)(184)(227)(11)(23)(226)
				return /*line Pimv8z.go:1*/ string(data)
			}(), "kind", eWW3WCQF, "last", glD6VIl86[ /*line qd9qvt8.go:1*/ len(glD6VIl86)-1], "err", chyZ8Q)
			return &sU5j5p{
				rACLehqaKtw:  glD6VIl86,
				fJRwxa3TtWCJ: xELkOUKoeEz,
				zvUxtDgF:     maGwUs,
				wOXUSabjPWCu: chyZ8Q,
				ahjfa5:       bFlDuoT,
			}, nil
		}
	}

	j5wt3kguOQ3, chyZ8Q := /*line KlguqLomRFa.go:1*/ trie.SPXGaGX(z1BBTN2UX, ijcATHiDv, glD6VIl86, xELkOUKoeEz, dc2CEa7W)
	return &sU5j5p{
			rACLehqaKtw:  glD6VIl86,
			fJRwxa3TtWCJ: xELkOUKoeEz,
			zvUxtDgF:     maGwUs,
			cMH_V88pSut:  j5wt3kguOQ3,
			wOXUSabjPWCu: chyZ8Q,
			ahjfa5:       bFlDuoT},
		nil
}

type jraijxKc func(nVUwQz8Zi []byte, jIbih8achE []byte, n8I7342 bool, gi7Xm3 bool) error

func (hbsnnIi *diskLayer) dw0mjyXBWN(dnam11P *qo43sXy, maaFGYOvd *trie.ID, sjBQgeKoXl []byte, eWW3WCQF string, ijcATHiDv []byte, gkQugurnMa int, vaTRmEh jraijxKc, hld1IO func([]byte) ([]byte, error)) (bool, []byte, error) {

	hftNwnRF, chyZ8Q := /*line u419UAd4sg4.go:1*/ hbsnnIi.dRlH06yalc(dnam11P, maaFGYOvd, sjBQgeKoXl, eWW3WCQF, ijcATHiDv, gkQugurnMa, hld1IO)
	if chyZ8Q != nil {
		return false, nil, chyZ8Q
	}
	h6VFlt6g := /*line SY7E15.go:1*/ hftNwnRF.h6VFlt6g()

	naUAO7HaG := []interface{}{"kind", eWW3WCQF, "prefix" /*line kag6WeSy.go:1*/, hexutil.Encode(sjBQgeKoXl)}
	if /*line JmJavOfYv.go:1*/ len(ijcATHiDv) > 0 {
		naUAO7HaG = /*line a7mzb0IbBl2G.go:1*/ append(naUAO7HaG, "origin" /*line ydyAspl.go:1*/, hexutil.Encode(ijcATHiDv))
	}
	hhTZiFI_0 := /*line BKLyXCw9RqF.go:1*/ log.New(naUAO7HaG...)

	if /*line PJD9uNEu.go:1*/ hftNwnRF.aIRqDs() {
		/*line rChQ95qh.go:1*/ m1L8ptKM24al.Mark(1)
		/*line TTn6qO.go:1*/ hhTZiFI_0.Trace(func() /*line u9ke0pJ_P.go:1*/ string {
			data := /*line u9ke0pJ_P.go:1*/ make([]byte, 0, 19)
			i := 0
			decryptKey := 68
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					data = /*line u9ke0pJ_P.go:1*/ append(data, "\r."...,
					)
					i = 1
				case 1:
					data = /*line u9ke0pJ_P.go:1*/ append(data, "0(<"...,
					)
					i = 6
				case 2:
					data = /*line u9ke0pJ_P.go:1*/ append(data, "*)"...,
					)
					i = 8
				case 5:
					i = 3
					data = /*line u9ke0pJ_P.go:1*/ append(data, ")!5#"...,
					)
				case 8:
					i = 4
					for y := range data {
						data[y] = data[y] ^ /*line u9ke0pJ_P.go:1*/ byte(decryptKey^y)
					}
				case 3:
					i = 7
					data = /*line u9ke0pJ_P.go:1*/ append(data, "3q"...,
					)
				case 7:
					i = 2
					data = /*line u9ke0pJ_P.go:1*/ append(data, "\"2<"...,
					)
				case 6:
					i = 5
					data = /*line u9ke0pJ_P.go:1*/ append(data, "<{"...,
					)
				}
			}
			return /*line u9ke0pJ_P.go:1*/ string(data)
		}(), "last" /*line FrhD3Q.go:1*/, hexutil.Encode(h6VFlt6g))

		if chyZ8Q := /*line lprhu_Ram.go:1*/ hftNwnRF.q9r2hg5(func(nVUwQz8Zi []byte, jIbih8achE []byte) error {
			return /*line QoKXA0klAVw.go:1*/ vaTRmEh(nVUwQz8Zi, jIbih8achE, false, false)
		}); chyZ8Q != nil {
			return false, nil, chyZ8Q
		}

		return !hftNwnRF.zvUxtDgF && !hftNwnRF.cMH_V88pSut, h6VFlt6g, nil
	}
	/*line g33BeTS.go:1*/ hhTZiFI_0.Trace(func() /*line J2RUuS.go:1*/ string {
		data := /*line J2RUuS.go:1*/ make([]byte, 0, 30)
		i := 6
		decryptKey := 156
		for counter := 0; i != 12; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 11:
				data = /*line J2RUuS.go:1*/ append(data, 125)
				i = 8
			case 1:
				i = 0
				data = /*line J2RUuS.go:1*/ append(data, "rwu"...,
				)
			case 5:
				i = 11
				data = /*line J2RUuS.go:1*/ append(data, "\x86t\x80x"...,
				)
			case 4:
				data = /*line J2RUuS.go:1*/ append(data, "~n"...,
				)
				i = 3
			case 8:
				for y := range data {
					data[y] = data[y] + /*line J2RUuS.go:1*/ byte(decryptKey^y)
				}
				i = 12
			case 3:
				data = /*line J2RUuS.go:1*/ append(data, 115)
				i = 7
			case 2:
				i = 5
				data = /*line J2RUuS.go:1*/ append(data, 61)
			case 9:
				data = /*line J2RUuS.go:1*/ append(data, "j\x80;\x8d"...,
				)
				i = 10
			case 10:
				data = /*line J2RUuS.go:1*/ append(data, "\x8d\x81\x93\x83"...,
				)
				i = 2
			case 6:
				data = /*line J2RUuS.go:1*/ append(data, "Pp"...,
				)
				i = 4
			case 7:
				data = /*line J2RUuS.go:1*/ append(data, "\x83sq$"...,
				)
				i = 1
			case 0:
				i = 9
				data = /*line J2RUuS.go:1*/ append(data, "lhz"...,
				)
			}
		}
		return /*line J2RUuS.go:1*/ string(data)
	}(), "last" /*line FTlf0h.go:1*/, hexutil.Encode(h6VFlt6g), "err", hftNwnRF.wOXUSabjPWCu)
	/*line OJaWaPZuJ.go:1*/ xsxTpq.Mark(1)

	if ijcATHiDv == nil && h6VFlt6g == nil {
		vuHmniJGKa7 := nNgqpt
		if eWW3WCQF == snapStorage {
			vuHmniJGKa7 = aKb5Mgikf
		}
		/*line Jx77bp.go:1*/ vuHmniJGKa7.Mark(1)
	}

	var o3BrVUwKh trie.NodeResolver
	if /*line BEaC7sT3wlhk.go:1*/ len(hftNwnRF.rACLehqaKtw) > 0 {
		h72TRHP0sl2 := /*line ISHEaFH82aNk.go:1*/ rawdb.NewMemoryDatabase()
		qxGcdD := /*line IeuYxGcsMB.go:1*/ triedb.NewDatabase(h72TRHP0sl2, triedb.HashDefaults)
		defer /*line qshD2QC.go:1*/ qxGcdD.Close()
		_FZ1mMvU := /*line ZjPat3wYm.go:1*/ trie.GysSMGsv(qxGcdD)
		for fSpMhzHR, nVUwQz8Zi := range hftNwnRF.rACLehqaKtw {
			/*line FNQeGZP.go:1*/ _FZ1mMvU.Update(nVUwQz8Zi, hftNwnRF.fJRwxa3TtWCJ[fSpMhzHR])
		}
		z1BBTN2UX, ogk6H1uqA, chyZ8Q := /*line DYKiPsaEQy.go:1*/ _FZ1mMvU.Commit(false)
		if chyZ8Q != nil {
			return false, nil, chyZ8Q
		}
		if ogk6H1uqA != nil {
			/*line MfwC3tHgA6.go:1*/ qxGcdD.Update(z1BBTN2UX, types.NrGuaNA21, 0 /*line bLVMTiNwkii.go:1*/, trienode.NewWithNodeSet(ogk6H1uqA), nil)
			/*line p3MDPoG0UY6.go:1*/ qxGcdD.Commit(z1BBTN2UX, false)
		}
		o3BrVUwKh = func(gUkUcM510UeB common.Hash, al0tnXGUu []byte, xhOzkRrKDZ common.Hash) []byte {
			return /*line WtAJvFLNe_r.go:1*/ rawdb.ReadTrieNode(h72TRHP0sl2, gUkUcM510UeB, al0tnXGUu, xhOzkRrKDZ /*line ByAVTK9P.go:1*/, qxGcdD.Scheme())
		}
	}

	bFlDuoT := hftNwnRF.ahjfa5
	if bFlDuoT == nil {
		bFlDuoT, chyZ8Q = /*line iBnj0a8Oc.go:1*/ trie.RJaQ3esB(maaFGYOvd, hbsnnIi.triedb)
		if chyZ8Q != nil {
			/*line pLeaFRund.go:1*/ dnam11P.g3wtYhVIxH4.Log(func() /*line RRQaCF9b.go:1*/ string {
				data := /*line RRQaCF9b.go:1*/ make([]byte, 0, 40)
				i := 17
				decryptKey := 193
				for counter := 0; i != 12; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 4:
						i = 10
						data = /*line RRQaCF9b.go:1*/ append(data, "\xc2\xd8\xdc\xc6"...,
						)
					case 9:
						data = /*line RRQaCF9b.go:1*/ append(data, "\xd2\xd8\xcc\xde"...,
						)
						i = 6
					case 13:
						i = 16
						data = /*line RRQaCF9b.go:1*/ append(data, "\xe9\xeb"...,
						)
					case 14:
						i = 13
						data = /*line RRQaCF9b.go:1*/ append(data, 254)
					case 16:
						i = 12
						for y := range data {
							data[y] = data[y] ^ /*line RRQaCF9b.go:1*/ byte(decryptKey^y)
						}
					case 6:
						data = /*line RRQaCF9b.go:1*/ append(data, "\x9a\xce\xd2"...,
						)
						i = 8
					case 17:
						data = /*line RRQaCF9b.go:1*/ append(data, "\xfd\xda"...,
						)
						i = 1
					case 1:
						data = /*line RRQaCF9b.go:1*/ append(data, 194)
						i = 3
					case 15:
						data = /*line RRQaCF9b.go:1*/ append(data, "\xdd\xd9\xd1"...,
						)
						i = 11
					case 2:
						data = /*line RRQaCF9b.go:1*/ append(data, 132)
						i = 18
					case 7:
						data = /*line RRQaCF9b.go:1*/ append(data, "ŉ"...,
						)
						i = 2
					case 0:
						i = 7
						data = /*line RRQaCF9b.go:1*/ append(data, "\xdd\xd2\xc9\xcd"...,
						)
					case 3:
						i = 5
						data = /*line RRQaCF9b.go:1*/ append(data, "ύ\xc1"...,
						)
					case 11:
						data = /*line RRQaCF9b.go:1*/ append(data, "\xa9\xf8\xea\xff"...,
						)
						i = 14
					case 18:
						i = 9
						data = /*line RRQaCF9b.go:1*/ append(data, 212)
					case 10:
						data = /*line RRQaCF9b.go:1*/ append(data, 193)
						i = 15
					case 5:
						data = /*line RRQaCF9b.go:1*/ append(data, 198)
						i = 0
					case 8:
						data = /*line RRQaCF9b.go:1*/ append(data, "\xde\xce"...,
						)
						i = 4
					}
				}
				return /*line RRQaCF9b.go:1*/ string(data)
			}(), hbsnnIi.root, hbsnnIi.genMarker)
			return false, nil, j0_A4Kp3
		}
	}
	var (
		l4NmGKfMkQss         bool
		ujlUWsOM13, wIxIkV6h = hftNwnRF.rACLehqaKtw, hftNwnRF.fJRwxa3TtWCJ

		nKfEsguBw = 0
		aBY2AlaT  = 0
		pBGkUnVA  = 0
		x4TqPFl4  = 0
		rmAaDQ    = 0

		lFRN2MXQc = /*line R8yjG0TP.go:1*/ time.Now()
		bBn_s0Tcb time.Duration
	)
	f1AgbL, chyZ8Q := /*line B5FLVwV.go:1*/ bFlDuoT.NodeIterator(ijcATHiDv)
	if chyZ8Q != nil {
		return false, nil, chyZ8Q
	}
	/*line SM1mrVkH.go:1*/ f1AgbL.AddResolver(o3BrVUwKh)
	lEtpKFfaKmLF := /*line noyGjEaA.go:1*/ trie.YH_VLHxOhT4(f1AgbL)

	for /*line Jf93C4BG4.go:1*/ lEtpKFfaKmLF.Next() {
		if h6VFlt6g != nil && /*line NEI68aRGG1o.go:1*/ bytes.Compare(lEtpKFfaKmLF.JyqRfQ2XM6, h6VFlt6g) > 0 {
			l4NmGKfMkQss = true
			break
		}
		nKfEsguBw++
		n8I7342 := true
		aBY2AlaT++
		for /*line YQ_h6HaX.go:1*/ len(ujlUWsOM13) > 0 {
			if mGzSMuzsfdY3 := /*line UnQQaz.go:1*/ bytes.Compare(ujlUWsOM13[0], lEtpKFfaKmLF.JyqRfQ2XM6); mGzSMuzsfdY3 < 0 {

				aRJ2USeHC := /*line nQn06dAg.go:1*/ time.Now()
				if chyZ8Q := /*line J0OlgD0rjSee.go:1*/ vaTRmEh(ujlUWsOM13[0], nil, false, true); chyZ8Q != nil {
					return false, nil, chyZ8Q
				}
				ujlUWsOM13 = ujlUWsOM13[1:]
				wIxIkV6h = wIxIkV6h[1:]
				x4TqPFl4++
				bBn_s0Tcb += /*line a1JOq1y.go:1*/ time.Since(aRJ2USeHC)
				continue
			} else if mGzSMuzsfdY3 == 0 {

				aBY2AlaT--
				if n8I7342 = ! /*line cZO_3bM3Nk.go:1*/ bytes.Equal(wIxIkV6h[0], lEtpKFfaKmLF.H884Yc); n8I7342 {
					pBGkUnVA++
				} else {
					rmAaDQ++
				}
				ujlUWsOM13 = ujlUWsOM13[1:]
				wIxIkV6h = wIxIkV6h[1:]
			}
			break
		}
		aRJ2USeHC := /*line G0GawPh29.go:1*/ time.Now()
		if chyZ8Q := /*line TwUK6Q.go:1*/ vaTRmEh(lEtpKFfaKmLF.JyqRfQ2XM6, lEtpKFfaKmLF.H884Yc, n8I7342, false); chyZ8Q != nil {
			return false, nil, chyZ8Q
		}
		bBn_s0Tcb += /*line xxiQCa.go:1*/ time.Since(aRJ2USeHC)
	}
	if lEtpKFfaKmLF.WitCUHMMV != nil {

		/*line idtaHgc.go:1*/
		log.Error(func() /*line j8gIg59nywU.go:1*/ string {
			data := [] /*line j8gIg59nywU.go:1*/ byte("S\xdea\xaee\xa0\xd1\xe2\xe5\xe2բo\xb3\xb0\xde\xe4\xdcW\x87Ol\xfb\xbf-To\xd1\xc6t\xbdr\xb5\xd4e ,\x9ciX")
			positions := [...]byte{20, 23, 10, 33, 28, 36, 28, 15, 32, 22, 37, 23, 25, 8, 18, 16, 8, 6, 16, 23, 16, 7, 20, 17, 1, 9, 14, 39, 19, 24, 19, 13, 11, 32, 17, 14, 27, 16, 17, 7, 3, 30, 32, 5}
			for i := 0; i < 44; i += 2 {
				localKey := /*line j8gIg59nywU.go:1*/ byte(i) + /*line j8gIg59nywU.go:1*/ byte(positions[i]^positions[i+1]) + 114
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line j8gIg59nywU.go:1*/ string(data)
		}(), "err", lEtpKFfaKmLF.WitCUHMMV)
		return false, nil, lEtpKFfaKmLF.WitCUHMMV
	}

	aRJ2USeHC := /*line ltXRQ5kmm.go:1*/ time.Now()
	for _, nVUwQz8Zi := range ujlUWsOM13 {
		if chyZ8Q := /*line JcRQI46Makqv.go:1*/ vaTRmEh(nVUwQz8Zi, nil, false, true); chyZ8Q != nil {
			return false, nil, chyZ8Q
		}
		x4TqPFl4 += 1
	}
	bBn_s0Tcb += /*line npPS_M.go:1*/ time.Since(aRJ2USeHC)

	if eWW3WCQF == snapStorage {
		/*line absTrj.go:1*/ hGMYBIAVt7.Inc(( /*line uQEbWKAL.go:1*/ time.Since(lFRN2MXQc) - bBn_s0Tcb).Nanoseconds())
	} else {
		/*line AvG3Rid.go:1*/ ewyA1FlVy.Inc(( /*line obNluVXnk7.go:1*/ time.Since(lFRN2MXQc) - bBn_s0Tcb).Nanoseconds())
	}
	/*line WNVZ98dO.go:1*/ hhTZiFI_0.Debug(func() /*line fBLKOUge0xM.go:1*/ string {
		data := [] /*line fBLKOUge0xM.go:1*/ byte("Reg\x94+\x8d\x93ate\xa1\xa4\xa8\xb7\x9a\x9dw\xb7Q\x8bnt\xd4")
		positions := [...]byte{11, 13, 18, 17, 3, 18, 14, 22, 21, 14, 5, 21, 19, 19, 15, 6, 10, 6, 5, 12, 11, 5, 6, 22, 18, 12, 5, 21, 4, 11, 22, 16}
		for i := 0; i < 32; i += 2 {
			localKey := /*line fBLKOUge0xM.go:1*/ byte(i) + /*line fBLKOUge0xM.go:1*/ byte(positions[i]^positions[i+1]) + 202
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line fBLKOUge0xM.go:1*/ string(data)
	}(), "root", maaFGYOvd.Root, "last" /*line IRfBcNb.go:1*/, hexutil.Encode(h6VFlt6g),
		"count", nKfEsguBw, "created", aBY2AlaT, "updated", pBGkUnVA, func() /*line KUmVnpBU.go:1*/ string {
			fullData := [] /*line KUmVnpBU.go:1*/ byte("\xdd\xe2\xe1ui\x10\xfd\rn\xf2hrx\x8e\a|\x00P")
			idxKey := [] /*line KUmVnpBU.go:1*/ byte("9Q\xab\x0e-B")
			data := /*line KUmVnpBU.go:1*/ make([]byte, 0, 10)
			data = /*line KUmVnpBU.go:1*/ append(data, fullData[30^ /*line KUmVnpBU.go:1*/ int(idxKey[3])]^fullData[13^ /*line KUmVnpBU.go:1*/ int(idxKey[3])], fullData[64^ /*line KUmVnpBU.go:1*/ int(idxKey[1])]-fullData[80^ /*line KUmVnpBU.go:1*/ int(idxKey[1])], fullData[66^ /*line KUmVnpBU.go:1*/ int(idxKey[5])]-fullData[70^ /*line KUmVnpBU.go:1*/ int(idxKey[5])], fullData[92^ /*line KUmVnpBU.go:1*/ int(idxKey[1])]+fullData[83^ /*line KUmVnpBU.go:1*/ int(idxKey[1])], fullData[1^ /*line KUmVnpBU.go:1*/ int(idxKey[3])]-fullData[0^ /*line KUmVnpBU.go:1*/ int(idxKey[3])], fullData[62^ /*line KUmVnpBU.go:1*/ int(idxKey[0])]^fullData[49^ /*line KUmVnpBU.go:1*/ int(idxKey[0])], fullData[2^ /*line KUmVnpBU.go:1*/ int(idxKey[3])]^fullData[11^ /*line KUmVnpBU.go:1*/ int(idxKey[3])], fullData[8^ /*line KUmVnpBU.go:1*/ int(idxKey[3])]+fullData[4^ /*line KUmVnpBU.go:1*/ int(idxKey[3])], fullData[75^ /*line KUmVnpBU.go:1*/ int(idxKey[5])]+fullData[73^ /*line KUmVnpBU.go:1*/ int(idxKey[5])])
			return /*line KUmVnpBU.go:1*/ string(data)
		}(), rmAaDQ, "deleted", x4TqPFl4)

	return !l4NmGKfMkQss && !hftNwnRF.zvUxtDgF, h6VFlt6g, nil
}

func (hbsnnIi *diskLayer) vGfNmysg3W(dnam11P *qo43sXy, y84IW11J []byte) error {
	var stpiKC6okfS8 chan *generatorStats
	select {
	case stpiKC6okfS8 = <-hbsnnIi.genAbort:
	default:
	}
	if /*line d2CBZXYC.go:1*/ dnam11P.fxipo_3eD.ValueSize() > ethdb.IdealBatchSize || stpiKC6okfS8 != nil {
		if /*line Kwk3S48.go:1*/ bytes.Compare(y84IW11J, hbsnnIi.genMarker) < 0 {
			/*line G8IHGY1_u8.go:1*/ log.Error(func() /*line BQHj8ec.go:1*/ string {
				seed := /*line BQHj8ec.go:1*/ byte(171)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line BQHj8ec.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line BQHj8ec.go:1*/ fnc(254)(23)(33)(81)(165)(63)(133)(15)(202)(219)(180)(113)(217)(191)(109)(237)(213)(173)(8)(103)(188)(129)(8)(188)(186)(115)(232)(216)(188)(98)(213)(156)(71)
				return /*line BQHj8ec.go:1*/ string(data)
			}(), "current" /*line TECJrBcIb.go:1*/, fmt.Sprintf("%x", y84IW11J), func() /*line n1INca89HCi.go:1*/ string {
				data := [] /*line n1INca89HCi.go:1*/ byte("\x7fw|nxtkWt")
				positions := [...]byte{3, 4, 5, 0, 3, 8, 3, 7, 2, 1, 1, 8}
				for i := 0; i < 12; i += 2 {
					localKey := /*line n1INca89HCi.go:1*/ byte(i) + /*line n1INca89HCi.go:1*/ byte(positions[i]^positions[i+1]) + 236
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line n1INca89HCi.go:1*/ string(data)
			}(), /*line DIzobjjDR.go:1*/ fmt.Sprintf("%x", hbsnnIi.genMarker))
		}

		/*line z_8h4DM.go:1*/
		bQ5V3QyAnw(dnam11P.fxipo_3eD, y84IW11J, dnam11P.g3wtYhVIxH4)

		if chyZ8Q := /*line APTldu.go:1*/ dnam11P.fxipo_3eD.Write(); chyZ8Q != nil {
			return chyZ8Q
		}
		/*line GNFQFNCs.go:1*/ dnam11P.fxipo_3eD.Reset()

		/*line Yy1zafZP.go:1*/
		hbsnnIi.lock.Lock()
		hbsnnIi.genMarker = y84IW11J
		/*line AW8rY7.go:1*/ hbsnnIi.lock.Unlock()

		if stpiKC6okfS8 != nil {
			/*line LXSKPfcc.go:1*/ dnam11P.g3wtYhVIxH4.Log(func() /*line trJkgP.go:1*/ string {
				data := [] /*line trJkgP.go:1*/ byte("\bb\a\x1e\xd0i\x06] [\x01\x0et%\xcfK\xc4\xe9p3&(&\xed\xd4eFer;8\x1d:\x16")
				positions := [...]byte{11, 24, 23, 14, 2, 2, 0, 7, 9, 0, 0, 16, 31, 18, 4, 32, 0, 17, 20, 11, 33, 33, 13, 31, 3, 3, 16, 9, 19, 13, 6, 29, 21, 14, 10, 0, 30, 19, 29, 22, 14, 32, 16, 15, 9, 26}
				for i := 0; i < 46; i += 2 {
					localKey := /*line trJkgP.go:1*/ byte(i) + /*line trJkgP.go:1*/ byte(positions[i]^positions[i+1]) + 148
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line trJkgP.go:1*/ string(data)
			}(), hbsnnIi.root, y84IW11J)
			return /*line ao56_d4IiV.go:1*/ qnWW8_(stpiKC6okfS8)
		}

		/*line pDj3fLRClpFS.go:1*/
		dnam11P.wOAalaayDr4C(snapAccount)
		/*line b2ekIea_eP.go:1*/ dnam11P.wOAalaayDr4C(snapStorage)
	}
	if /*line f0hg1UPJR.go:1*/ time.Since(dnam11P.aRunvr) > 8*time.Second {
		/*line hmjoFIlRp.go:1*/ dnam11P.g3wtYhVIxH4.Log(func() /*line T2p3kYw.go:1*/ string {
			key := [] /*line T2p3kYw.go:1*/ byte("n\x8d\n\x1a\xceOWj\x84\xefn\x16!(z\xdfk\xbf\xfc.c)\x9b(\xca")
			data := [] /*line T2p3kYw.go:1*/ byte("\xd9\xd8dK\xa4\x12\x1d\xff\xeax\xb2]S9\xfa\x86\xb5\xb4r3\rJ\xcdG\xaa")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line T2p3kYw.go:1*/ string(data)
		}(), hbsnnIi.root, y84IW11J)
		dnam11P.aRunvr = /*line ZSaE2REizWVu.go:1*/ time.Now()
	}
	return nil
}

func pv9KAud4D(dnam11P *qo43sXy, hbsnnIi *diskLayer, fGWXYaz common.Hash, evzmhL1 common.Hash, a3_RBWxx08hW common.Hash, tNjZuUpnjlXC []byte) error {
	lrGP0LSN := func(nVUwQz8Zi []byte, jIbih8achE []byte, n8I7342 bool, gi7Xm3 bool) error {
		defer func( /*line X40ZQT.go:1*/ lFRN2MXQc time.Time) {
			/*line CXHD2ep0qi.go:1*/ sTEsaYjvIa.Inc( /*line FDCjxCdeXuM.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
		}( /*line DnWvQ7DQQ9Y.go:1*/ time.Now())

		if gi7Xm3 {
			/*line t7cK9xA.go:1*/ rawdb.DeleteStorageSnapshot(dnam11P.fxipo_3eD, evzmhL1 /*line abAtZw.go:1*/, common.BytesToHash(nVUwQz8Zi))
			/*line kQ41Q8WSo.go:1*/ aUn5w2V.Mark(1)
			return nil
		}
		if n8I7342 {
			/*line ifjGZQT.go:1*/ rawdb.WriteStorageSnapshot(dnam11P.fxipo_3eD, evzmhL1 /*line px9Jh1MVwk_o.go:1*/, common.BytesToHash(nVUwQz8Zi), jIbih8achE)
			/*line B1gxKz.go:1*/ i50G1Ge.Mark(1)
		} else {
			/*line EqlaxKQ4Uar.go:1*/ orgtTMsZp.Mark(1)
		}
		dnam11P.g3wtYhVIxH4.storage += /*line NaTnaD0_.go:1*/ common.StorageSize(1 + 2*common.HashLength + /*line zBT8xH.go:1*/ len(jIbih8achE))
		dnam11P.g3wtYhVIxH4.slots++

		if chyZ8Q := /*line xfttVgA.go:1*/ hbsnnIi.vGfNmysg3W(dnam11P /*line hrwdtiyw.go:1*/, append(evzmhL1[:], nVUwQz8Zi...)); chyZ8Q != nil {
			return chyZ8Q
		}
		return nil
	}

	var ijcATHiDv = /*line BGyfba12RWG.go:1*/ common.CopyBytes(tNjZuUpnjlXC)
	for {
		lMBced9 := /*line JkWQodbf3h.go:1*/ trie.NYkaq0T(fGWXYaz, evzmhL1, a3_RBWxx08hW)
		rsGxVsqZR5F9, h6VFlt6g, chyZ8Q := /*line qAbpTjNb.go:1*/ hbsnnIi.dw0mjyXBWN(dnam11P, lMBced9 /*line pYvzp6PEie.go:1*/, append(rawdb.SnapshotStoragePrefix /*line Pa_xMtHCR5.go:1*/, evzmhL1.Bytes()...), snapStorage, ijcATHiDv, lZKEBSD, lrGP0LSN, nil)
		if chyZ8Q != nil {
			return chyZ8Q
		}

		if rsGxVsqZR5F9 {
			break
		}
		if ijcATHiDv = /*line BpBxwR.go:1*/ uZWCppliS(h6VFlt6g); ijcATHiDv == nil {
			break
		}
	}
	return nil
}

func zPhbCBxk(dnam11P *qo43sXy, hbsnnIi *diskLayer, higPZ5 []byte) error {
	zesPF2d := func(nVUwQz8Zi []byte, jIbih8achE []byte, n8I7342 bool, gi7Xm3 bool) error {

		evzmhL1 := /*line J0WSid.go:1*/ common.BytesToHash(nVUwQz8Zi)
		/*line Hyjou4KVmShV.go:1*/ dnam11P.q8QodbI(evzmhL1)

		lFRN2MXQc := /*line fvU24vhoyx.go:1*/ time.Now()
		if gi7Xm3 {
			/*line syN0cr9f8.go:1*/ rawdb.DeleteAccountSnapshot(dnam11P.fxipo_3eD, evzmhL1)
			/*line NpaMdPBn.go:1*/ fBYs5Maa.Mark(1)
			/*line PDFQ4xw1.go:1*/ f1iu2ct2.Inc( /*line YFmuNK.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())

			/*line IevmJU5_06.go:1*/
			dnam11P.scD2LUFBkjzR(evzmhL1)
			return nil
		}

		var h1PKeV3DPS types.StateAccount
		if chyZ8Q := /*line m19Z63P2nS.go:1*/ rlp.DecodeBytes(jIbih8achE, &h1PKeV3DPS); chyZ8Q != nil {
			/*line n4ta8yG2hUp.go:1*/ log.Crit(func() /*line i_cIWzKDcex5.go:1*/ string {
				fullData := [] /*line i_cIWzKDcex5.go:1*/ byte("\xb3Q\n\xeb)\xa0\xf1_\xb5*Md\xaaH\xf3\x05\xe7\x17?\xb6L+\x9c<U\x91\x90\xb2\xb7\x12\xae\x1e\xfd\xf4\xab\xa0\x14\x0f\x9a&\x94Z5\xe7\xd2\xf6\x0e\xba\x88\x90~\xeb\xd8S;č!}}\x93\xf7\xc0%0~;jW\a!\xffO\x84\xf9=\x9d\xfc\x92\x8eH@\x8a\xf7\x83qڼP\xbb\xfd\xaf\xbb\xfe8ݞ\xe4iP\xf6\x11\xbb ")
				idxKey := [] /*line i_cIWzKDcex5.go:1*/ byte("\u07b7\x13\x9a[a<\x846\xb0\xfb\x1a\aa\x99\x87")
				data := /*line i_cIWzKDcex5.go:1*/ make([]byte, 0, 53)
				data = /*line i_cIWzKDcex5.go:1*/ append(data, fullData[166^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[15])]-fullData[165^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[15])], fullData[138^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])]^fullData[200^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])], fullData[174^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])]^fullData[130^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])], fullData[88^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])]+fullData[61^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])], fullData[163^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])]^fullData[208^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])], fullData[183^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])]+fullData[182^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])], fullData[46^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[13])]-fullData[104^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[13])], fullData[171^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])]-fullData[162^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])], fullData[248^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])]-fullData[169^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])], fullData[78^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])]-fullData[62^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])], fullData[97^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[13])]-fullData[2^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[13])], fullData[197^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])]^fullData[225^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])], fullData[104^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])]^fullData[17^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])], fullData[125^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])]^fullData[3^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])], fullData[210^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])]^fullData[154^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])], fullData[201^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])]+fullData[173^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])], fullData[211^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])]-fullData[177^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])], fullData[153^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])]-fullData[148^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])], fullData[12^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[2])]-fullData[117^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[2])], fullData[101^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])]+fullData[0^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])], fullData[28^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])]^fullData[12^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])], fullData[36^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])]^fullData[10^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])], fullData[82^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])]-fullData[93^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])], fullData[42^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])]^fullData[26^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])], fullData[210^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[9])]-fullData[227^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[9])], fullData[2^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])]+fullData[87^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])], fullData[3^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])]+fullData[46^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])], fullData[184^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])]+fullData[188^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[7])], fullData[147^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])]+fullData[239^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])], fullData[66^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[2])]^fullData[57^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[2])], fullData[82^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])]-fullData[69^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])], fullData[189^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])]^fullData[246^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])], fullData[86^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])]-fullData[118^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])], fullData[140^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])]+fullData[142^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[14])], fullData[255^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])]+fullData[198^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])], fullData[13^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])]+fullData[82^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])], fullData[252^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])]+fullData[222^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[10])], fullData[93^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])]^fullData[106^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])], fullData[114^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])]-fullData[29^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[8])], fullData[187^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[9])]^fullData[161^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[9])], fullData[18^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])]-fullData[16^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])], fullData[174^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])]^fullData[234^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[1])], fullData[191^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[0])]+fullData[196^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[0])], fullData[66^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])]-fullData[23^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])], fullData[12^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])]^fullData[93^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[11])], fullData[137^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[3])]+fullData[205^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[3])], fullData[51^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])]-fullData[57^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])], fullData[236^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[9])]-fullData[153^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[9])], fullData[127^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])]+fullData[62^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[6])], fullData[6^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])]^fullData[89^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])], fullData[61^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])]^fullData[26^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[12])], fullData[22^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])]^fullData[21^ /*line i_cIWzKDcex5.go:1*/ int(idxKey[4])])
				return /*line i_cIWzKDcex5.go:1*/ string(data)
			}(), "err", chyZ8Q)
		}

		if higPZ5 == nil || ! /*line r0TQYdCbixL_.go:1*/ bytes.Equal(evzmhL1[:], higPZ5) {
			rrGntsl := /*line eTG8sX1hEIfk.go:1*/ len(jIbih8achE)
			if !n8I7342 {
				if /*line YoORu2EDSaqX.go:1*/ bytes.Equal(h1PKeV3DPS.CodeHash, types.JhrQnETMFm[:]) {
					rrGntsl -= 32
				}
				if h1PKeV3DPS.Root == types.NrGuaNA21 {
					rrGntsl -= 32
				}
				/*line doVkba.go:1*/ lW_o9U5IloX.Mark(1)
			} else {
				bI1ciqVUvL4b := /*line MukPK3Rfa.go:1*/ types.K2bJTLXfmaB(h1PKeV3DPS)
				rrGntsl = /*line ysQaXO7.go:1*/ len(bI1ciqVUvL4b)
				/*line EM5zl5S.go:1*/ rawdb.WriteAccountSnapshot(dnam11P.fxipo_3eD, evzmhL1, bI1ciqVUvL4b)
				/*line Fh2c4L.go:1*/ ew7oLoPaR.Mark(1)
			}
			dnam11P.g3wtYhVIxH4.storage += /*line BABrrMR3q.go:1*/ common.StorageSize(1 + common.HashLength + rrGntsl)
			dnam11P.g3wtYhVIxH4.accounts++
		}

		al5iRGOo := evzmhL1[:]
		if higPZ5 != nil && /*line RdXLxkP.go:1*/ bytes.Equal(al5iRGOo, higPZ5) && /*line _dA4Y8.go:1*/ len(hbsnnIi.genMarker) > common.HashLength {
			al5iRGOo = hbsnnIi.genMarker[:]
		}

		if chyZ8Q := /*line EPdBZsdYd.go:1*/ hbsnnIi.vGfNmysg3W(dnam11P, al5iRGOo); chyZ8Q != nil {
			return chyZ8Q
		}
		/*line Lys91q.go:1*/ f1iu2ct2.Inc( /*line hdJtqigZJjdO.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())

		if h1PKeV3DPS.Root == types.NrGuaNA21 {
			/*line dKaBgxH.go:1*/ dnam11P.scD2LUFBkjzR(evzmhL1)
		} else {
			var tNjZuUpnjlXC []byte
			if higPZ5 != nil && /*line DKFeXW0KeUe.go:1*/ bytes.Equal(evzmhL1[:], higPZ5) && /*line XpaP_EeWJcU.go:1*/ len(hbsnnIi.genMarker) > common.HashLength {
				tNjZuUpnjlXC = hbsnnIi.genMarker[common.HashLength:]
			}
			if chyZ8Q := /*line z_m6pDU9r1iO.go:1*/ pv9KAud4D(dnam11P, hbsnnIi, hbsnnIi.root, evzmhL1, h1PKeV3DPS.Root, tNjZuUpnjlXC); chyZ8Q != nil {
				return chyZ8Q
			}
		}

		higPZ5 = nil
		return nil
	}

	var bISCIEW9w63 = inviux
	if /*line aiyNTcyCb.go:1*/ len(higPZ5) > 0 {
		bISCIEW9w63 = 1
	}
	ijcATHiDv := /*line AWY28D7_B.go:1*/ common.CopyBytes(higPZ5)
	for {
		lMBced9 := /*line a8Q1gA.go:1*/ trie.P5xt0mP5FivA(hbsnnIi.root)
		rsGxVsqZR5F9, h6VFlt6g, chyZ8Q := /*line p2FjelC.go:1*/ hbsnnIi.dw0mjyXBWN(dnam11P, lMBced9, rawdb.SnapshotAccountPrefix, snapAccount, ijcATHiDv, bISCIEW9w63, zesPF2d, types.R3iwnzGSe)
		if chyZ8Q != nil {
			return chyZ8Q
		}
		ijcATHiDv = /*line bmhjT_ANN.go:1*/ uZWCppliS(h6VFlt6g)

		if ijcATHiDv == nil || rsGxVsqZR5F9 {
			/*line MhAaj9ee_Z7v.go:1*/ dnam11P.fhK1vcHFHY()
			break
		}
		bISCIEW9w63 = inviux
	}
	return nil
}

func (hbsnnIi *diskLayer) tliqpCic(drjkE0 *generatorStats) {
	var (
		higPZ5       []byte
		stpiKC6okfS8 chan *generatorStats
	)
	if /*line KLHO1W.go:1*/ len(hbsnnIi.genMarker) > 0 {
		higPZ5 = hbsnnIi.genMarker[:common.HashLength]
	}
	/*line iCJmEUuXA.go:1*/ drjkE0.Log(func() /*line iiT_YcIzpU.go:1*/ string {
		data := [] /*line iiT_YcIzpU.go:1*/ byte("\xe2\x17s\xf7min\xdd\x19g`atw\xfe\x10T\xedpQ\xdc\xe3\xd4axen\xb3\x91a\xecoo\x17")
		positions := [...]byte{7, 9, 17, 19, 31, 24, 31, 22, 10, 0, 28, 23, 22, 8, 15, 28, 14, 13, 3, 22, 21, 16, 17, 28, 30, 13, 0, 20, 14, 19, 21, 17, 21, 7, 27, 14, 3, 33, 24, 7, 30, 1}
		for i := 0; i < 42; i += 2 {
			localKey := /*line iiT_YcIzpU.go:1*/ byte(i) + /*line iiT_YcIzpU.go:1*/ byte(positions[i]^positions[i+1]) + 92
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line iiT_YcIzpU.go:1*/ string(data)
	}(), hbsnnIi.root, hbsnnIi.genMarker)

	dnam11P := /*line rbfNkpin1I6.go:1*/ wNiHmW_Ox8k(drjkE0, hbsnnIi.diskdb, higPZ5, hbsnnIi.genMarker)
	defer /*line fMOWxZEDiar.go:1*/ dnam11P.f_Fw0rNs4()

	if chyZ8Q := /*line lCGJV8P.go:1*/ zPhbCBxk(dnam11P, hbsnnIi, higPZ5); chyZ8Q != nil {

		if rq5ZXwEn, jdkNTRyBJ := chyZ8Q.(*c5NDRd); jdkNTRyBJ {
			stpiKC6okfS8 = rq5ZXwEn.v9cWaTe
		}

		if stpiKC6okfS8 == nil {
			stpiKC6okfS8 = <-hbsnnIi.genAbort
		}
		stpiKC6okfS8 <- drjkE0
		return
	}

	/*line mUVo4AbIbP0.go:1*/
	bQ5V3QyAnw(dnam11P.fxipo_3eD, nil, drjkE0)
	if chyZ8Q := /*line hE72E15A.go:1*/ dnam11P.fxipo_3eD.Write(); chyZ8Q != nil {
		/*line GHXQ2Il.go:1*/ log.Error(func() /*line ateoiJjOVGLa.go:1*/ string {
			seed := /*line ateoiJjOVGLa.go:1*/ byte(225)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line ateoiJjOVGLa.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line ateoiJjOVGLa.go:1*/ fnc(167)(233)(24)(229)(11)(29)(182)(56)(235)(79)(216)(250)(229)(6)(19)(174)(94)(251)(225)(21)(227)
			return /*line ateoiJjOVGLa.go:1*/ string(data)
		}(), "err", chyZ8Q)

		stpiKC6okfS8 = <-hbsnnIi.genAbort
		stpiKC6okfS8 <- drjkE0
		return
	}
	/*line X5NaAHlj.go:1*/ dnam11P.fxipo_3eD.Reset()

	/*line Tg6AIBKL.go:1*/
	log.Info(func() /*line cGAlWOB_aJ.go:1*/ string {
		seed := /*line cGAlWOB_aJ.go:1*/ byte(89)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line cGAlWOB_aJ.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line cGAlWOB_aJ.go:1*/ fnc(30)(18)(231)(21)(247)(29)(237)(227)(13)(86)(191)(255)(235)(1)(19)(169)(65)(29)(241)(241)(1)(27)(225)(27)
		return /*line cGAlWOB_aJ.go:1*/ string(data)
	}(), func() /*line wqYIebByj.go:1*/ string {
		data := [] /*line wqYIebByj.go:1*/ byte("[\xe2T\xd8\xdb\xd7\xe2s")
		positions := [...]byte{2, 0, 6, 6, 1, 5, 4, 2, 0, 3}
		for i := 0; i < 10; i += 2 {
			localKey := /*line wqYIebByj.go:1*/ byte(i) + /*line wqYIebByj.go:1*/ byte(positions[i]^positions[i+1]) + 108
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line wqYIebByj.go:1*/ string(data)
	}(), drjkE0.accounts, "slots", drjkE0.slots,
		"storage", drjkE0.storage, func() /*line YipmdE.go:1*/ string {
			data := /*line YipmdE.go:1*/ make([]byte, 0, 9)
			i := 5
			decryptKey := 228
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data = /*line YipmdE.go:1*/ append(data, 137)
					i = 0
				case 6:
					i = 4
					data = /*line YipmdE.go:1*/ append(data, 133)
				case 0:
					i = 1
					data = /*line YipmdE.go:1*/ append(data, 131)
				case 8:
					data = /*line YipmdE.go:1*/ append(data, 127)
					i = 9
				case 9:
					data = /*line YipmdE.go:1*/ append(data, 141)
					i = 7
				case 7:
					i = 6
					data = /*line YipmdE.go:1*/ append(data, 135)
				case 5:
					i = 8
					data = /*line YipmdE.go:1*/ append(data, 129)
				case 1:
					i = 3
					for y := range data {
						data[y] = data[y] + /*line YipmdE.go:1*/ byte(decryptKey^y)
					}
				case 4:
					i = 2
					data = /*line YipmdE.go:1*/ append(data, 131)
				}
			}
			return /*line YipmdE.go:1*/ string(data)
		}(), drjkE0.dangling, "elapsed" /*line c1eCVw5kBqZ.go:1*/, common.PrettyDuration( /*line XYBqmjwXHtcW.go:1*/ time.Since(drjkE0.start)))

	/*line NxrDt5xoDVf.go:1*/
	hbsnnIi.lock.Lock()
	hbsnnIi.genMarker = nil
	/*line uisB9aeLC.go:1*/ close(hbsnnIi.genPending)
	/*line IdEiyy.go:1*/ hbsnnIi.lock.Unlock()

	stpiKC6okfS8 = <-hbsnnIi.genAbort
	stpiKC6okfS8 <- nil
}

func uZWCppliS(nVUwQz8Zi []byte) []byte {
	for fSpMhzHR := /*line E7nAIP1PB0w.go:1*/ len(nVUwQz8Zi) - 1; fSpMhzHR >= 0; fSpMhzHR-- {
		nVUwQz8Zi[fSpMhzHR]++
		if nVUwQz8Zi[fSpMhzHR] != 0x0 {
			return nVUwQz8Zi
		}
	}
	return nil
}

type c5NDRd struct {
	v9cWaTe chan *generatorStats
}

func qnWW8_(stpiKC6okfS8 chan *generatorStats) error {
	return &c5NDRd{v9cWaTe: stpiKC6okfS8}
}

func (chyZ8Q *c5NDRd) Error() string {
	return "aborted"
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append

const _ = time.ANSIC

var _ types.AccessList
var _ trie.SXoLHxUt177q
var _ fastcache.BigStats
var _ = common.AbsolutePath
var _ hexutil.Big
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ = log.Crit
var _ = rlp.AppendUint64
var _ trienode.MergedNodeSet
var _ triedb.Config
