//line :1
package snapshot

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
)

const (
	snapAccount = "account"
	snapStorage = "storage"
)

type generatorStats struct {
	origin   uint64
	start    time.Time
	accounts uint64
	slots    uint64
	dangling uint64
	storage  common.StorageSize
}

func (pY22ntC *generatorStats) Log(wnl9hzf60nbg string, z1BBTN2UX common.Hash, al5iRGOo []byte) {
	var dnam11P []interface{}
	if z1BBTN2UX != (common.Hash{}) {
		dnam11P = /*line DXSJO4cFo.go:1*/ append(dnam11P, []interface{}{"root", z1BBTN2UX}...)
	}

	switch /*line MaZW7sd8d.go:1*/ len(al5iRGOo) {
	case common.HashLength:
		dnam11P = /*line IxIvuE2G.go:1*/ append(dnam11P, []interface{}{"at" /*line RNQwFr.go:1*/, common.BytesToHash(al5iRGOo)}...)
	case 2 * common.HashLength:
		dnam11P = /*line UbQeMd9vzA.go:1*/ append(dnam11P, []interface{}{
			"in" /*line vzBal1xYvwH.go:1*/, common.BytesToHash(al5iRGOo[:common.HashLength]),
			"at" /*line qd4ymo4lGV.go:1*/, common.BytesToHash(al5iRGOo[common.HashLength:]),
		}...)
	}

	dnam11P = /*line zgiBako.go:1*/ append(dnam11P, []interface{}{
		func() /*line uWX43zQ.go:1*/ string {
			fullData := [] /*line uWX43zQ.go:1*/ byte("\x92I\xa4\x01\x8aރ*\x91t\a)\xf9\xed\xf3K")
			idxKey := [] /*line uWX43zQ.go:1*/ byte("\x84\xc1")
			data := /*line uWX43zQ.go:1*/ make([]byte, 0, 9)
			data = /*line uWX43zQ.go:1*/ append(data, fullData[207^ /*line uWX43zQ.go:1*/ int(idxKey[1])]^fullData[193^ /*line uWX43zQ.go:1*/ int(idxKey[1])], fullData[203^ /*line uWX43zQ.go:1*/ int(idxKey[1])]-fullData[195^ /*line uWX43zQ.go:1*/ int(idxKey[1])], fullData[192^ /*line uWX43zQ.go:1*/ int(idxKey[1])]^fullData[198^ /*line uWX43zQ.go:1*/ int(idxKey[1])], fullData[140^ /*line uWX43zQ.go:1*/ int(idxKey[0])]+fullData[129^ /*line uWX43zQ.go:1*/ int(idxKey[0])], fullData[200^ /*line uWX43zQ.go:1*/ int(idxKey[1])]+fullData[194^ /*line uWX43zQ.go:1*/ int(idxKey[1])], fullData[199^ /*line uWX43zQ.go:1*/ int(idxKey[1])]^fullData[204^ /*line uWX43zQ.go:1*/ int(idxKey[1])], fullData[143^ /*line uWX43zQ.go:1*/ int(idxKey[0])]+fullData[139^ /*line uWX43zQ.go:1*/ int(idxKey[0])], fullData[136^ /*line uWX43zQ.go:1*/ int(idxKey[0])]^fullData[128^ /*line uWX43zQ.go:1*/ int(idxKey[0])])
			return /*line uWX43zQ.go:1*/ string(data)
		}(), pY22ntC.accounts,
		"slots", pY22ntC.slots,
		"storage", pY22ntC.storage,
		func() /*line zaP4QVday.go:1*/ string {
			seed := /*line zaP4QVday.go:1*/ byte(127)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line zaP4QVday.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line zaP4QVday.go:1*/ fnc(227)(195)(147)(31)(67)(131)(11)(15)
			return /*line zaP4QVday.go:1*/ string(data)
		}(), pY22ntC.dangling,
		"elapsed" /*line ADLLU3a.go:1*/, common.PrettyDuration( /*line ZGbKtgUAUbE.go:1*/ time.Since(pY22ntC.start)),
	}...)

	if /*line eHHjYPa2ZN.go:1*/ len(al5iRGOo) > 0 {
		if k5lF2W9o := /*line u1Ii7Q7.go:1*/ binary.BigEndian.Uint64(al5iRGOo[:8]) - pY22ntC.origin; k5lF2W9o > 0 {
			lPjY_TaIBHq := math.MaxUint64 - /*line SAxL3fVjaG.go:1*/ binary.BigEndian.Uint64(al5iRGOo[:8])

			jB5nFJLw := k5lF2W9o/ /*line unZtaF.go:1*/ uint64( /*line X8kYCu3i4_z.go:1*/ time.Since(pY22ntC.start)/time.Millisecond+1) + 1
			dnam11P = /*line GDtAGSrzC0XB.go:1*/ append(dnam11P, []interface{}{
				"eta" /*line tYCutla_Svh.go:1*/, common.PrettyDuration( /*line ABfJKJe.go:1*/ time.Duration(lPjY_TaIBHq/jB5nFJLw) * time.Millisecond),
			}...)
		}
	}
	/*line xfx_Sw.go:1*/ log.Info(wnl9hzf60nbg, dnam11P...)
}

type qo43sXy struct {
	g3wtYhVIxH4  *generatorStats
	eVwCwm       ethdb.KeyValueStore
	yeG2yi_8     *r5WWYA_
	cAUfxRkqDgk4 *r5WWYA_
	fxipo_3eD    ethdb.Batch
	aRunvr       time.Time
}

func wNiHmW_Ox8k(drjkE0 *generatorStats, lDBEbd ethdb.KeyValueStore, higPZ5 []byte, qeaXSK []byte) *qo43sXy {
	dnam11P := &qo43sXy{
		g3wtYhVIxH4: drjkE0,
		eVwCwm:      lDBEbd,
		fxipo_3eD:/*line BIwTxOUmFL0q.go:1*/ lDBEbd.NewBatch(),
		aRunvr:/*line HgcJg2.go:1*/ time.Now(),
	}
	/*line DgQRpwT.go:1*/ dnam11P.uA98lbQW(snapAccount, higPZ5)
	/*line zv10fXDv.go:1*/ dnam11P.uA98lbQW(snapStorage, qeaXSK)
	return dnam11P
}

func (dnam11P *qo43sXy) uA98lbQW(eWW3WCQF string, lFRN2MXQc []byte) {
	if eWW3WCQF == snapAccount {
		lEtpKFfaKmLF := /*line CtFE5aH.go:1*/ dnam11P.eVwCwm.NewIterator(rawdb.SnapshotAccountPrefix, lFRN2MXQc)
		dnam11P.yeG2yi_8 = /*line m6MP_2.go:1*/ szfg10dIoV( /*line xKVjUlt.go:1*/ rawdb.NewKeyLengthIterator(lEtpKFfaKmLF, 1+common.HashLength))
		return
	}
	lEtpKFfaKmLF := /*line DcgypSc.go:1*/ dnam11P.eVwCwm.NewIterator(rawdb.SnapshotStoragePrefix, lFRN2MXQc)
	dnam11P.cAUfxRkqDgk4 = /*line Pkv9NSrUJ.go:1*/ szfg10dIoV( /*line xrNVUpGvh9M.go:1*/ rawdb.NewKeyLengthIterator(lEtpKFfaKmLF, 1+2*common.HashLength))
}

func (dnam11P *qo43sXy) wOAalaayDr4C(eWW3WCQF string) {

	var lEtpKFfaKmLF = dnam11P.yeG2yi_8
	if eWW3WCQF == snapStorage {
		lEtpKFfaKmLF = dnam11P.cAUfxRkqDgk4
	}
	gHv1IbVh04no := /*line TsYT1RXWw.go:1*/ lEtpKFfaKmLF.Next()
	if !gHv1IbVh04no {

		/*line lAeTZdu.go:1*/
		lEtpKFfaKmLF.Release()
		if eWW3WCQF == snapAccount {
			dnam11P.yeG2yi_8 = /*line VOcHksPk.go:1*/ szfg10dIoV( /*line e5hufzykoVNF.go:1*/ memorydb.New().NewIterator(nil, nil))
			return
		}
		dnam11P.cAUfxRkqDgk4 = /*line zRqV0aoK.go:1*/ szfg10dIoV( /*line aGppL39p.go:1*/ memorydb.New().NewIterator(nil, nil))
		return
	}
	zYawlDUNOaHH := /*line l0KyNY.go:1*/ lEtpKFfaKmLF.Key()
	/*line RdtVFxPsEO0t.go:1*/ lEtpKFfaKmLF.Release()
	/*line FD_n5G.go:1*/ dnam11P.uA98lbQW(eWW3WCQF, zYawlDUNOaHH[1:])
}

func (dnam11P *qo43sXy) f_Fw0rNs4() {
	/*line icWjjqu.go:1*/ dnam11P.yeG2yi_8.Release()
	/*line E9hyma.go:1*/ dnam11P.cAUfxRkqDgk4.Release()
}

func (dnam11P *qo43sXy) ahOoDUdo(eWW3WCQF string) *r5WWYA_ {
	if eWW3WCQF == snapAccount {
		return dnam11P.yeG2yi_8
	}
	return dnam11P.cAUfxRkqDgk4
}

func (dnam11P *qo43sXy) q8QodbI(evzmhL1 common.Hash) {
	var (
		nKfEsguBw    uint64
		lFRN2MXQc    = /*line F3aKjTwPcPJ.go:1*/ time.Now()
		lEtpKFfaKmLF = dnam11P.cAUfxRkqDgk4
	)
	for /*line ufVDvF.go:1*/ lEtpKFfaKmLF.Next() {
		nVUwQz8Zi := /*line oB5DOqw3ENN.go:1*/ lEtpKFfaKmLF.Key()
		if /*line UFzNaAWWmu.go:1*/ bytes.Compare(nVUwQz8Zi[1:1+common.HashLength] /*line uw9Oo_4FA35.go:1*/, evzmhL1.Bytes()) >= 0 {
			/*line h1B9q2W7Fkf.go:1*/ lEtpKFfaKmLF.Hold()
			break
		}
		nKfEsguBw++
		/*line jMHpltroD.go:1*/ dnam11P.fxipo_3eD.Delete(nVUwQz8Zi)
		if /*line HpIKC1.go:1*/ dnam11P.fxipo_3eD.ValueSize() > ethdb.IdealBatchSize {
			/*line Uif5qqL2Jmnp.go:1*/ dnam11P.fxipo_3eD.Write()
			/*line FoetsPC5.go:1*/ dnam11P.fxipo_3eD.Reset()
		}
	}
	dnam11P.g3wtYhVIxH4.dangling += nKfEsguBw
	/*line m2pmZw6iL2J.go:1*/ pjvp9N8ibUBI.Inc( /*line AcaXo3jItJKn.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
}

func (dnam11P *qo43sXy) scD2LUFBkjzR(evzmhL1 common.Hash) error {
	var (
		nKfEsguBw    int64
		lFRN2MXQc    = /*line ozoKCd.go:1*/ time.Now()
		lEtpKFfaKmLF = dnam11P.cAUfxRkqDgk4
	)
	for /*line CVQ5vwYNH6.go:1*/ lEtpKFfaKmLF.Next() {
		nVUwQz8Zi := /*line BN1Ir3An.go:1*/ lEtpKFfaKmLF.Key()
		mGzSMuzsfdY3 := /*line bPoEUf.go:1*/ bytes.Compare(nVUwQz8Zi[1:1+common.HashLength] /*line uxOCwi.go:1*/, evzmhL1.Bytes())
		if mGzSMuzsfdY3 < 0 {
			return /*line caTL3XdRlH.go:1*/ errors.New(func() /*line Gk30U4ky.go:1*/ string {
				seed := /*line Gk30U4ky.go:1*/ byte(53)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line Gk30U4ky.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line Gk30U4ky.go:1*/ fnc(52)(5)(8)(235)(11)(253)(251)(188)(73)(11)(241)(13)(239)(19)(251)(3)(174)(80)(255)(4)(246)(11)(245)(6)(255)
				return /*line Gk30U4ky.go:1*/ string(data)
			}())
		}
		if mGzSMuzsfdY3 > 0 {
			/*line cZxlaoMW7NQ.go:1*/ lEtpKFfaKmLF.Hold()
			break
		}
		nKfEsguBw++
		/*line SVYNGi7Spp.go:1*/ dnam11P.fxipo_3eD.Delete(nVUwQz8Zi)
		if /*line AXN_KSDaBuN.go:1*/ dnam11P.fxipo_3eD.ValueSize() > ethdb.IdealBatchSize {
			/*line NAVMuef2.go:1*/ dnam11P.fxipo_3eD.Write()
			/*line Jo8WrC.go:1*/ dnam11P.fxipo_3eD.Reset()
		}
	}
	/*line Pm37aN5te1ba.go:1*/ aUn5w2V.Mark(nKfEsguBw)
	/*line L4mGzlwdXaRj.go:1*/ pjvp9N8ibUBI.Inc( /*line D__3o703VOIh.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
	return nil
}

func (dnam11P *qo43sXy) fhK1vcHFHY() {
	var (
		nKfEsguBw    uint64
		lFRN2MXQc    = /*line VaAYu4Q.go:1*/ time.Now()
		lEtpKFfaKmLF = dnam11P.cAUfxRkqDgk4
	)
	for /*line v0nwgN6Zp7.go:1*/ lEtpKFfaKmLF.Next() {
		nKfEsguBw++
		/*line FxQG_sMM8HK8.go:1*/ dnam11P.fxipo_3eD.Delete( /*line WAOGrg8yIs.go:1*/ lEtpKFfaKmLF.Key())
		if /*line O_SxJ5Ixq_.go:1*/ dnam11P.fxipo_3eD.ValueSize() > ethdb.IdealBatchSize {
			/*line FByUeRp.go:1*/ dnam11P.fxipo_3eD.Write()
			/*line wJNp7S.go:1*/ dnam11P.fxipo_3eD.Reset()
		}
	}
	dnam11P.g3wtYhVIxH4.dangling += nKfEsguBw
	/*line zsacZT1.go:1*/ krsDyKF9M1a.Mark( /*line INRFpPacN.go:1*/ int64(nKfEsguBw))
	/*line bwY1Ee.go:1*/ pjvp9N8ibUBI.Inc( /*line mldAMIfTZw.go:1*/ time.Since(lFRN2MXQc).Nanoseconds())
}

var _ bytes.Buffer
var _ = binary.Append
var _ = errors.As

const _ = time.ANSIC

var _ = common.AbsolutePath
var _ = math.BigMax
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ memorydb.Database
var _ = log.Crit
