//line :1
package snapshot

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	trie "unishard/evm/trie"
	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

type nTKI2Efv struct {
	vENhG6_Sjr common.Hash
	lw3vbv     []byte
}

type (
	h6VYD59LL func(lDBEbd ethdb.KeyValueWriter, jtZSrgtJ3bB string, gUkUcM510UeB common.Hash, rCUYPH11 chan (nTKI2Efv), w_jHRmYmG chan (common.Hash))

	oGgEtanTv func(lDBEbd ethdb.KeyValueWriter, dcJdHV, j8UX1s common.Hash, lXazFVVqJfD_ *bS1vBqz) (common.Hash, error)
)

func ZMz3a4JjAC(cu8ZKpmK Nq4YsH_) (common.Hash, error) {
	return /*line veMyjPsrb.go:1*/ qQF5aSpK2E9(nil, "", cu8ZKpmK, common.Hash{}, d8k2UL, nil /*line o5oKW4KXli0J.go:1*/, kT7u5MT_Lv(), true)
}

func ZNvyOsNDgSYx(evzmhL1 common.Hash, cu8ZKpmK PE_7UyJghxy) (common.Hash, error) {
	return /*line wcx_DZdqL5C.go:1*/ qQF5aSpK2E9(nil, "", cu8ZKpmK, evzmhL1, d8k2UL, nil /*line jrgaFAh.go:1*/, kT7u5MT_Lv(), true)
}

func MzPaxfxF(vT262n *CE2m1DUB6wW, z1BBTN2UX common.Hash, eswHK65 ethdb.Database, uaANK1cwczz ethdb.KeyValueWriter) error {

	avYdDZ2tZiG, chyZ8Q := /*line IXsLdcUoXZE.go:1*/ vT262n.AccountIterator(z1BBTN2UX, common.Hash{})
	if chyZ8Q != nil {
		return chyZ8Q
	}
	defer /*line jXW9gTTJ.go:1*/ avYdDZ2tZiG.Release()

	jtZSrgtJ3bB := /*line VuJOrYIDJ.go:1*/ vT262n.sKxDZofwkrd.Scheme()
	pfeznfxya, chyZ8Q := /*line nLma7a.go:1*/ qQF5aSpK2E9(uaANK1cwczz, jtZSrgtJ3bB, avYdDZ2tZiG, common.Hash{}, d8k2UL, func(uaANK1cwczz ethdb.KeyValueWriter, dcJdHV, j8UX1s common.Hash, lXazFVVqJfD_ *bS1vBqz) (common.Hash, error) {

		if j8UX1s != types.JhrQnETMFm {
			jTcZY5bQSaNL := /*line Haad753_SS.go:1*/ rawdb.ReadCode(eswHK65, j8UX1s)
			if /*line bCwIao2BHRJ7.go:1*/ len(jTcZY5bQSaNL) == 0 {
				return common.Hash{} /*line Vcn72kJ.go:1*/, errors.New(func() /*line U8F9_oW9IJN.go:1*/ string {
					key := [] /*line U8F9_oW9IJN.go:1*/ byte("<jI\xcc)}\xe2-\x9d\r\x12f,\xe0\xc6\xefNe\fZ\x03\xbd\x9c\xf2#\x96\x16\x01")
					data := [] /*line U8F9_oW9IJN.go:1*/ byte("*\xf7 \xa0<\xe7>G\xd2\x13`\xff5\x84Zt!\th\x18^\xa6\xd8.@\xd9Nd")
					for i, b := range key {
						data[i] = data[i] + b
					}
					return /*line U8F9_oW9IJN.go:1*/ string(data)
				}())
			}
			/*line Y5DLmt2YHD4.go:1*/ rawdb.WriteCode(uaANK1cwczz, j8UX1s, jTcZY5bQSaNL)
		}

		sK5vOn6BL2OZ, chyZ8Q := /*line J76FxcR.go:1*/ vT262n.StorageIterator(z1BBTN2UX, dcJdHV, common.Hash{})
		if chyZ8Q != nil {
			return common.Hash{}, chyZ8Q
		}
		defer /*line Uw2r01cetCO.go:1*/ sK5vOn6BL2OZ.Release()

		xhOzkRrKDZ, chyZ8Q := /*line DXdihn1erHr.go:1*/ qQF5aSpK2E9(uaANK1cwczz, jtZSrgtJ3bB, sK5vOn6BL2OZ, dcJdHV, d8k2UL, nil, lXazFVVqJfD_, false)
		if chyZ8Q != nil {
			return common.Hash{}, chyZ8Q
		}
		return xhOzkRrKDZ, nil
	}, /*line hY4cJr.go:1*/ kT7u5MT_Lv(), true)

	if chyZ8Q != nil {
		return chyZ8Q
	}
	if pfeznfxya != z1BBTN2UX {
		return /*line _2C9cZH7Hn.go:1*/ fmt.Errorf(func() /*line nnl6I90x7.go:1*/ string {
			key := [] /*line nnl6I90x7.go:1*/ byte("ɤm\x8d߂`F\x89\x93\xae{'\x9e\x16\xe2\x87\xef\xf3\x88J\xaa\xca\x17Y\xf9\xd2m\xab\xe3\xa4\xd4~C\x00\x90Q\x01P\\\xb9")
			data := [] /*line nnl6I90x7.go:1*/ byte("\xba\xd0\f\xf9\xba\xa2\x12)\xe6\xe7\x8e\x13F\xed~\xc2ꆀ\xe5+ީ\x7fcٵ\x02\xdfÁ\xacRcw\xf1?upy\xc1")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line nnl6I90x7.go:1*/ string(data)
		}(), pfeznfxya, z1BBTN2UX)
	}
	return nil
}

type bS1vBqz struct {
	m6uZdar common.Hash
	_nnPqAe time.Time

	kywPZCe55Otv uint64
	nW44olrH     uint64

	oItojHaaoQw map[common.Hash]time.Time
	jMgEhGkh    map[common.Hash]common.Hash

	aej95zWMdd sync.RWMutex
}

func kT7u5MT_Lv() *bS1vBqz {
	return &bS1vBqz{
		oItojHaaoQw:/*line Wq1XBI.go:1*/ make(map[common.Hash]time.Time),
		jMgEhGkh:/*line w4AetXV.go:1*/ make(map[common.Hash]common.Hash),
		_nnPqAe:/*line wbXSF2GU.go:1*/ time.Now(),
	}
}

func (lXazFVVqJfD_ *bS1vBqz) erG4tHs(evzmhL1 common.Hash, k5lF2W9o uint64) {
	/*line LPQtQ4QO.go:1*/ lXazFVVqJfD_.aej95zWMdd.Lock()
	defer /*line VUOrCk.go:1*/ lXazFVVqJfD_.aej95zWMdd.Unlock()

	lXazFVVqJfD_.kywPZCe55Otv += k5lF2W9o
	lXazFVVqJfD_.m6uZdar = evzmhL1
}

func (lXazFVVqJfD_ *bS1vBqz) xhRltz12_Jds(k5lF2W9o uint64) {
	/*line hXCHucU75Th.go:1*/ lXazFVVqJfD_.aej95zWMdd.Lock()
	defer /*line is50354a13ag.go:1*/ lXazFVVqJfD_.aej95zWMdd.Unlock()

	lXazFVVqJfD_.kywPZCe55Otv += k5lF2W9o
}

func (lXazFVVqJfD_ *bS1vBqz) fwC0zLiduCLo(evzmhL1 common.Hash, j3V0_MwKAiG common.Hash, k5lF2W9o uint64) {
	/*line HgNLOuRFxPU.go:1*/ lXazFVVqJfD_.aej95zWMdd.Lock()
	defer /*line rhad4YemSieY.go:1*/ lXazFVVqJfD_.aej95zWMdd.Unlock()

	lXazFVVqJfD_.nW44olrH += k5lF2W9o
	lXazFVVqJfD_.jMgEhGkh[evzmhL1] = j3V0_MwKAiG
	if _, jdkNTRyBJ := lXazFVVqJfD_.oItojHaaoQw[evzmhL1]; !jdkNTRyBJ {
		lXazFVVqJfD_.oItojHaaoQw[evzmhL1] = /*line LUPGToBonfTa.go:1*/ time.Now()
	}
}

func (lXazFVVqJfD_ *bS1vBqz) epwx18(evzmhL1 common.Hash, k5lF2W9o uint64) {
	/*line DyaRUaWSdKD.go:1*/ lXazFVVqJfD_.aej95zWMdd.Lock()
	defer /*line O7L52s.go:1*/ lXazFVVqJfD_.aej95zWMdd.Unlock()

	lXazFVVqJfD_.nW44olrH += k5lF2W9o
	/*line VnUvOt8md.go:1*/ delete(lXazFVVqJfD_.jMgEhGkh, evzmhL1)
	/*line t0CoUbilcy.go:1*/ delete(lXazFVVqJfD_.oItojHaaoQw, evzmhL1)
}

func (lXazFVVqJfD_ *bS1vBqz) s2mS29OcT1w() {
	/*line PUcEES4H.go:1*/ lXazFVVqJfD_.aej95zWMdd.RLock()
	defer /*line Nyzvb5KOnX.go:1*/ lXazFVVqJfD_.aej95zWMdd.RUnlock()

	dnam11P := []interface{}{
		func() /*line mbMMht.go:1*/ string {
			key := [] /*line mbMMht.go:1*/ byte("\xb9\xdaWt\x1a\x9f\xbcl")
			data := [] /*line mbMMht.go:1*/ byte("\xa8\x89\f\xfb[ϸ\a")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line mbMMht.go:1*/ string(data)
		}(), lXazFVVqJfD_.kywPZCe55Otv,
		"slots", lXazFVVqJfD_.nW44olrH,
		"elapsed" /*line EO4WsPiPj.go:1*/, common.PrettyDuration( /*line DENXRA.go:1*/ time.Since(lXazFVVqJfD_._nnPqAe)),
	}
	if lXazFVVqJfD_.kywPZCe55Otv > 0 {

		if k5lF2W9o := /*line Ie525eLMrb7O.go:1*/ binary.BigEndian.Uint64(lXazFVVqJfD_.m6uZdar[:8]) / lXazFVVqJfD_.kywPZCe55Otv; k5lF2W9o > 0 {
			var (
				lPjY_TaIBHq = (math.MaxUint64 - /*line Rj1CgCdx.go:1*/ binary.BigEndian.Uint64(lXazFVVqJfD_.m6uZdar[:8])) / lXazFVVqJfD_.kywPZCe55Otv
				jB5nFJLw    = k5lF2W9o/ /*line qFxMfihS1R.go:1*/ uint64( /*line MsW7xt0w9zTq.go:1*/ time.Since(lXazFVVqJfD_._nnPqAe)/time.Millisecond+1) + 1
				eykU3i      = /*line G1tNYf4.go:1*/ time.Duration(lPjY_TaIBHq/jB5nFJLw) * time.Millisecond
			)

			for h1PKeV3DPS, j9xp0xEW9BzB := range lXazFVVqJfD_.jMgEhGkh {
				lFRN2MXQc := lXazFVVqJfD_.oItojHaaoQw[h1PKeV3DPS]
				if k5lF2W9o := /*line fuH8Av7kSAay.go:1*/ binary.BigEndian.Uint64(j9xp0xEW9BzB[:8]); k5lF2W9o > 0 {
					var (
						lPjY_TaIBHq = math.MaxUint64 - /*line kn3TGOSbY.go:1*/ binary.BigEndian.Uint64(j9xp0xEW9BzB[:8])
						jB5nFJLw    = k5lF2W9o/ /*line xHh6u8_As.go:1*/ uint64( /*line XaRtzu1b.go:1*/ time.Since(lFRN2MXQc)/time.Millisecond+1) + 1
					)

					if x1LBZUayY4 := /*line SSUiUmSpaL0.go:1*/ time.Duration(lPjY_TaIBHq/jB5nFJLw) * time.Millisecond; eykU3i < x1LBZUayY4 {
						eykU3i = x1LBZUayY4
					}
				}
			}
			dnam11P = /*line oiO6yX.go:1*/ append(dnam11P, []interface{}{
				"eta" /*line A7UYZY.go:1*/, common.PrettyDuration(eykU3i),
			}...)
		}
	}
	/*line zAzKGKUgd4.go:1*/ log.Info(func() /*line TLkOP70.go:1*/ string {
		seed := /*line TLkOP70.go:1*/ byte(166)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line TLkOP70.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line TLkOP70.go:1*/ fnc(163)(43)(241)(13)(239)(19)(245)(5)(249)(185)(83)(1)(237)(19)(241)(187)(83)(251)(243)(15)(3)(245)(7)(5)
		return /*line TLkOP70.go:1*/ string(data)
	}(), dnam11P...)
}

func (lXazFVVqJfD_ *bS1vBqz) jh1PsTWf6w() {
	/*line IuBKV0A_t.go:1*/ lXazFVVqJfD_.aej95zWMdd.RLock()
	defer /*line HQFMRgp0xq.go:1*/ lXazFVVqJfD_.aej95zWMdd.RUnlock()

	var dnam11P []interface{}
	dnam11P = /*line KmR1JzzGQAHa.go:1*/ append(dnam11P, []interface{}{func() /*line uOAmOkHtGI.go:1*/ string {
		fullData := [] /*line uOAmOkHtGI.go:1*/ byte("Oh\xe8\xden\x86\xc0iq\x05:\x1e'M\xb2\x06")
		idxKey := [] /*line uOAmOkHtGI.go:1*/ byte(">j")
		data := /*line uOAmOkHtGI.go:1*/ make([]byte, 0, 9)
		data = /*line uOAmOkHtGI.go:1*/ append(data, fullData[50^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]+fullData[52^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])], fullData[107^ /*line uOAmOkHtGI.go:1*/ int(idxKey[1])]-fullData[99^ /*line uOAmOkHtGI.go:1*/ int(idxKey[1])], fullData[48^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]-fullData[62^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])], fullData[53^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]^fullData[54^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])], fullData[61^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]-fullData[57^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])], fullData[60^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]+fullData[59^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])], fullData[49^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]+fullData[58^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])], fullData[56^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])]-fullData[51^ /*line uOAmOkHtGI.go:1*/ int(idxKey[0])])
		return /*line uOAmOkHtGI.go:1*/ string(data)
	}(), lXazFVVqJfD_.kywPZCe55Otv}...)
	if lXazFVVqJfD_.nW44olrH != 0 {
		dnam11P = /*line pme6Z_.go:1*/ append(dnam11P, []interface{}{"slots", lXazFVVqJfD_.nW44olrH}...)
	}
	dnam11P = /*line qOGwP04iKfi.go:1*/ append(dnam11P, []interface{}{"elapsed" /*line mWIb9S1ziLFF.go:1*/, common.PrettyDuration( /*line jee7aH6.go:1*/ time.Since(lXazFVVqJfD_._nnPqAe))}...)
	/*line sSmHRYEMkm.go:1*/ log.Info(func() /*line prX3Mqa4p50i.go:1*/ string {
		key := [] /*line prX3Mqa4p50i.go:1*/ byte("\xbb\xd1\xf5\x1b\xa1i4\x16\x86ۿ\\N!l\xd2\xf5")
		data := [] /*line prX3Mqa4p50i.go:1*/ byte("\x04EZ\x8d\x02ݙz\xa6N-\xbd\xbe\x94\xd4Ai")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line prX3Mqa4p50i.go:1*/ string(data)
	}(), dnam11P...)
}

func ndaq_GfvBr(drjkE0 *bS1vBqz, dI4AXBMU chan bool) {
	sa0gIG_6S := /*line ifpqh0UPljgj.go:1*/ time.NewTimer(0)
	defer /*line psMqBgHeN.go:1*/ sa0gIG_6S.Stop()

	for {
		select {
		case <-sa0gIG_6S.C:
			/*line Tivb4dbfbHV.go:1*/ drjkE0.s2mS29OcT1w()
			/*line QjgpDgekHR.go:1*/ sa0gIG_6S.Reset(time.Second * 8)
		case qI7KXIVdU := <-dI4AXBMU:
			if qI7KXIVdU {
				/*line HVCCyO_TK.go:1*/ drjkE0.jh1PsTWf6w()
			}
			return
		}
	}
}

func qQF5aSpK2E9(lDBEbd ethdb.KeyValueWriter, jtZSrgtJ3bB string, cu8ZKpmK FcPREnURc, evzmhL1 common.Hash, g8wn3jnO h6VYD59LL, gOm37Ll8RA2 oGgEtanTv, drjkE0 *bS1vBqz, s2mS29OcT1w bool) (common.Hash, error) {
	var (
		rCUYPH11   = /*line mZWcPRjg.go:1*/ make(chan nTKI2Efv)
		w_jHRmYmG  = /*line AARuKombpf.go:1*/ make(chan common.Hash, 1)
		yzn54DGo6O = /*line MyzURoK.go:1*/ make(chan bool, 1)
		eT8ykpf9   sync.WaitGroup
	)

	/*line JaRw9Z.go:1*/
	eT8ykpf9.Add(1)
	go func() {
		defer /*line FmHG8PKa6pTF.go:1*/ eT8ykpf9.Done()
		/*line CzIWnIU1F9X9.go:1*/ g8wn3jnO(lDBEbd, jtZSrgtJ3bB, evzmhL1, rCUYPH11, w_jHRmYmG)
	}()

	if s2mS29OcT1w && drjkE0 != nil {
		/*line xmADCsUuXK.go:1*/ eT8ykpf9.Add(1)
		go func() {
			defer /*line FvFdkyXM.go:1*/ eT8ykpf9.Done()
			/*line aNn4l49.go:1*/ ndaq_GfvBr(drjkE0, yzn54DGo6O)
		}()
	}

	gapnXx := /*line raxSwgqBcjf.go:1*/ runtime.NumCPU()
	pU9TrEjNm := /*line vs24KnHRa8sD.go:1*/ make(chan error, gapnXx)
	for fSpMhzHR := 0; fSpMhzHR < gapnXx; fSpMhzHR++ {
		pU9TrEjNm <- nil
	}

	dI4AXBMU := func(fW1k6t error) (common.Hash, error) {
		/*line rbK80tR.go:1*/ close(rCUYPH11)
		hftNwnRF := <-w_jHRmYmG
		for fSpMhzHR := 0; fSpMhzHR < gapnXx; fSpMhzHR++ {
			if chyZ8Q := <-pU9TrEjNm; chyZ8Q != nil && fW1k6t == nil {
				fW1k6t = chyZ8Q
			}
		}
		yzn54DGo6O <- fW1k6t == nil

		/*line tmMPdg.go:1*/
		eT8ykpf9.Wait()
		return hftNwnRF, fW1k6t
	}
	var (
		jy_QWtkwazhi = /*line lXNkYAa.go:1*/ time.Now()
		vdgsVi_D     = /*line I2hg3o.go:1*/ uint64(0)
		eKg15ZOwwj   nTKI2Efv
	)

	for /*line jv_MxMedmkx3.go:1*/ cu8ZKpmK.Next() {
		if evzmhL1 == (common.Hash{}) {
			var (
				chyZ8Q  error
				fV7iQ3I []byte
			)
			if gOm37Ll8RA2 == nil {
				fV7iQ3I, chyZ8Q = /*line VUir_Q.go:1*/ types.R3iwnzGSe( /*line BNT3tS_.go:1*/ cu8ZKpmK.(Nq4YsH_).Account())
				if chyZ8Q != nil {
					return /*line RGbDzpuBJas.go:1*/ dI4AXBMU(chyZ8Q)
				}
			} else {

				if chyZ8Q := <-pU9TrEjNm; chyZ8Q != nil {
					pU9TrEjNm <- nil
					return /*line hsy_jqi9GQhu.go:1*/ dI4AXBMU(chyZ8Q)
				}

				evzmhL1, chyZ8Q := /*line NQ3svo9i.go:1*/ types.VGE1getZq8( /*line ozJbsBs.go:1*/ cu8ZKpmK.(Nq4YsH_).Account())
				if chyZ8Q != nil {
					return /*line A9rc2Qx19.go:1*/ dI4AXBMU(chyZ8Q)
				}
				go func( /*line Yea3ynFMZMTV.go:1*/ xhOzkRrKDZ common.Hash) {
					cz4eo5R, chyZ8Q := /*line SX5WDcU.go:1*/ gOm37Ll8RA2(lDBEbd, xhOzkRrKDZ /*line fwCENNr.go:1*/, common.BytesToHash(evzmhL1.CodeHash), drjkE0)
					if chyZ8Q != nil {
						pU9TrEjNm <- chyZ8Q
						return
					}
					if evzmhL1.Root != cz4eo5R {
						pU9TrEjNm <- /*line B3oDzgyscqM2.go:1*/ fmt.Errorf(func() /*line FkNDh245A.go:1*/ string {
							key := [] /*line FkNDh245A.go:1*/ byte("\xedx\x13\xb7{`\xd1\xf8p\xf5Iнġ\xf8\xa5\x9d0iDV{GK<\xe5FF\xb7[D\x9a&\xb7\xa31\x13,xd\xbe")
							data := [] /*line FkNDh245A.go:1*/ byte("|\xf6c\xaa\xf1\t\x93(\x03\x80\x19\xa2\xb2\xab\xd30\xcb\xc4D\xff\xdc\xcf\xfd\xe2\xe1\xe4\x92\x1b(\xbd\xc5\xe1\xde\x06i\xc50c9\xa8\xc1\xba")
							for i, b := range key {
								data[i] = data[i] + b
							}
							return /*line FkNDh245A.go:1*/ string(data)
						}(), xhOzkRrKDZ, evzmhL1.Root, cz4eo5R)
						return
					}
					pU9TrEjNm <- nil
				}( /*line cMdYOMz1N.go:1*/ cu8ZKpmK.Hash())
				fV7iQ3I, chyZ8Q = /*line BIEFEA02KlA.go:1*/ rlp.EncodeToBytes(evzmhL1)
				if chyZ8Q != nil {
					return /*line A0RSDmH.go:1*/ dI4AXBMU(chyZ8Q)
				}
			}
			eKg15ZOwwj = nTKI2Efv{ /*line Oswih22Zw.go:1*/ cu8ZKpmK.Hash(), fV7iQ3I}
		} else {
			eKg15ZOwwj = nTKI2Efv{ /*line g0OOdG.go:1*/ cu8ZKpmK.Hash() /*line dUSLzET.go:1*/, common.CopyBytes( /*line TbjnVqMD.go:1*/ cu8ZKpmK.(PE_7UyJghxy).Slot())}
		}
		rCUYPH11 <- eKg15ZOwwj

		vdgsVi_D++
		if /*line IbFAYD6.go:1*/ time.Since(jy_QWtkwazhi) > 3*time.Second && drjkE0 != nil {
			if evzmhL1 == (common.Hash{}) {
				/*line WqLEwXl2XMX.go:1*/ drjkE0.erG4tHs( /*line LhAIl2DnhX3k.go:1*/ cu8ZKpmK.Hash(), vdgsVi_D)
			} else {
				/*line X2kkVVd5ZOW.go:1*/ drjkE0.fwC0zLiduCLo(evzmhL1 /*line fwDtQG9kqtOe.go:1*/, cu8ZKpmK.Hash(), vdgsVi_D)
			}
			jy_QWtkwazhi, vdgsVi_D = /*line FGXCG1gn.go:1*/ time.Now(), 0
		}
	}

	if vdgsVi_D > 0 && drjkE0 != nil {
		if evzmhL1 == (common.Hash{}) {
			/*line E9D1OZp.go:1*/ drjkE0.xhRltz12_Jds(vdgsVi_D)
		} else {
			/*line xSzzJ0M.go:1*/ drjkE0.epwx18(evzmhL1, vdgsVi_D)
		}
	}
	return /*line Ri87s8at.go:1*/ dI4AXBMU(nil)
}

func d8k2UL(lDBEbd ethdb.KeyValueWriter, jtZSrgtJ3bB string, gUkUcM510UeB common.Hash, rCUYPH11 chan nTKI2Efv, w_jHRmYmG chan common.Hash) {
	jMwa6C := /*line m88z07.go:1*/ trie.HN5J4sr7gxe()
	if lDBEbd != nil {
		jMwa6C = /*line OiE9u3eAO9.go:1*/ jMwa6C.WithWriter(func(al0tnXGUu []byte, xhOzkRrKDZ common.Hash, b1jafJL []byte) {
			/*line GHvAOf7XWjf.go:1*/ rawdb.WriteTrieNode(lDBEbd, gUkUcM510UeB, al0tnXGUu, xhOzkRrKDZ, b1jafJL, jtZSrgtJ3bB)
		})
	}
	rNemY7HB := /*line n2awMD_nGQN.go:1*/ trie.KYaYz2rvh7Rt(jMwa6C)
	for eKg15ZOwwj := range rCUYPH11 {
		/*line ToIYiDK8.go:1*/ rNemY7HB.Update(eKg15ZOwwj.vENhG6_Sjr[:], eKg15ZOwwj.lw3vbv)
	}
	w_jHRmYmG <- /*line GnVYToQA8aSI.go:1*/ rNemY7HB.Commit()
}

var _ = binary.Append
var _ = errors.As
var _ = fmt.Append
var _ = math.Abs
var _ = runtime.BlockProfile
var _ sync.Cond

const _ = time.ANSIC

var _ trie.SXoLHxUt177q
var _ types.AccessList
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ = log.Crit
var _ = rlp.AppendUint64
