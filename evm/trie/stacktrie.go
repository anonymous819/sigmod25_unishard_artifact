//line :1
package trie

import (
	"bytes"
	"errors"
	"sync"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	fbnoIefSAHah = sync.Pool{New: func() any { return /*line F_QqChz4s_.go:1*/ new(waMJLSBe) }}
	_            = /*line DFNHtRwSd.go:1*/ types.ZGJpq7HtE((* /*line C940TcyJ_l.go:1*/ JKAJrncEF)(nil))
)

type Kr3mtK4 struct {
	KILMj82    func(qiRG6lpeaFl []byte, rNuN0sMPDJ common.Hash, aBHduJG []byte)
	O1sMUJt_OI func(qiRG6lpeaFl []byte)

	D1MUlGq     bool
	AbM2xa0De   bool
	pLNVqr5xZzH metrics.Gauge
}

func HN5J4sr7gxe() *Kr3mtK4 { return &Kr3mtK4{} }

func (tTpi95sF750 *Kr3mtK4) WithWriter(vtVGAZ7 func(qiRG6lpeaFl []byte, rNuN0sMPDJ common.Hash, aBHduJG []byte)) *Kr3mtK4 {
	tTpi95sF750.KILMj82 = vtVGAZ7
	return tTpi95sF750
}

func (tTpi95sF750 *Kr3mtK4) WithCleaner(iyDLnQLEfhc func(qiRG6lpeaFl []byte)) *Kr3mtK4 {
	tTpi95sF750.O1sMUJt_OI = iyDLnQLEfhc
	return tTpi95sF750
}

func (tTpi95sF750 *Kr3mtK4) WithSkipBoundary(zTbiLa34, gNHFqKvzHWYp bool, f4DprUQ0ACzA metrics.Gauge) *Kr3mtK4 {
	tTpi95sF750.D1MUlGq = zTbiLa34
	tTpi95sF750.AbM2xa0De = gNHFqKvzHWYp
	tTpi95sF750.pLNVqr5xZzH = f4DprUQ0ACzA
	return tTpi95sF750
}

type JKAJrncEF struct {
	kmfQF3JHh1b *Kr3mtK4
	dbm3YIa     *waMJLSBe
	hEg7x8gVZOy *d6imPJ

	enzdgudgyf4 []byte
	aCCP1ML     []byte
}

func KYaYz2rvh7Rt(wgt9q9wGfaIH *Kr3mtK4) *JKAJrncEF {
	if wgt9q9wGfaIH == nil {
		wgt9q9wGfaIH = /*line seAKJWbMJ7aN.go:1*/ HN5J4sr7gxe()
	}
	return &JKAJrncEF{
		kmfQF3JHh1b: wgt9q9wGfaIH,
		dbm3YIa:/*line ivwNNbJCB.go:1*/ fbnoIefSAHah.Get().(*waMJLSBe),
		hEg7x8gVZOy:/*line lzgWsZMFkQAB.go:1*/ dCZJziX(false),
	}
}

func (hkI2UXG_QBd *JKAJrncEF) Update(lhQWH7m, ar76Sw []byte) error {
	if /*line c1QbYNEe.go:1*/ len(ar76Sw) == 0 {
		return /*line He2xPWEbu.go:1*/ errors.New(func() /*line JbabXsDdv.go:1*/ string {
			data := /*line JbabXsDdv.go:1*/ make([]byte, 0, 34)
			i := 10
			decryptKey := 154
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 14:
					i = 3
					data = /*line JbabXsDdv.go:1*/ append(data, "6*)%"...,
					)
				case 2:
					i = 12
					data = /*line JbabXsDdv.go:1*/ append(data, ",6!5"...,
					)
				case 6:
					i = 0
					data = /*line JbabXsDdv.go:1*/ append(data, 61)
				case 0:
					i = 8
					data = /*line JbabXsDdv.go:1*/ append(data, "88@"...,
					)
				case 13:
					data = /*line JbabXsDdv.go:1*/ append(data, "?7!"...,
					)
					i = 6
				case 8:
					i = 5
					for y := range data {
						data[y] = data[y] ^ /*line JbabXsDdv.go:1*/ byte(decryptKey^y)
					}
				case 7:
					data = /*line JbabXsDdv.go:1*/ append(data, 53)
					i = 4
				case 12:
					i = 14
					data = /*line JbabXsDdv.go:1*/ append(data, "2y="...,
					)
				case 9:
					i = 2
					data = /*line JbabXsDdv.go:1*/ append(data, ".`*"...,
					)
				case 10:
					i = 11
					data = /*line JbabXsDdv.go:1*/ append(data, "=:"...,
					)
				case 4:
					i = 13
					data = /*line JbabXsDdv.go:1*/ append(data, 53)
				case 1:
					data = /*line JbabXsDdv.go:1*/ append(data, "o:"...,
					)
					i = 9
				case 3:
					data = /*line JbabXsDdv.go:1*/ append(data, "\x7fv"...,
					)
					i = 7
				case 11:
					data = /*line JbabXsDdv.go:1*/ append(data, "2##+"...,
					)
					i = 1
				}
			}
			return /*line JbabXsDdv.go:1*/ string(data)
		}())
	}
	ec_diEU := /*line hzwCldJpO.go:1*/ nd12pU880Z8(lhQWH7m)
	ec_diEU = ec_diEU[: /*line rI4NA1L1d.go:1*/ len(ec_diEU)-1]
	if /*line VSvvn1Zpr.go:1*/ bytes.Compare(hkI2UXG_QBd.aCCP1ML, ec_diEU) >= 0 {
		return /*line Bl2l_fN4l4Y.go:1*/ errors.New(func() /*line IbPnxg.go:1*/ string {
			seed := /*line IbPnxg.go:1*/ byte(224)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line IbPnxg.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line IbPnxg.go:1*/ fnc(142)(1)(255)(191)(52)(18)(240)(2)(9)(246)(5)(5)(249)(185)(75)(250)(20)(167)(79)(3)(242)(1)(13)
			return /*line IbPnxg.go:1*/ string(data)
		}())
	}

	if hkI2UXG_QBd.enzdgudgyf4 == nil {
		hkI2UXG_QBd.enzdgudgyf4 = /*line NWWDECg.go:1*/ append([]byte{}, ec_diEU...)
	}
	if hkI2UXG_QBd.aCCP1ML == nil {
		hkI2UXG_QBd.aCCP1ML = /*line xxWl49mDMAd.go:1*/ append([]byte{}, ec_diEU...)
	} else {
		hkI2UXG_QBd.aCCP1ML = /*line bDwpdcGm.go:1*/ append(hkI2UXG_QBd.aCCP1ML[:0], ec_diEU...)
	}
	/*line KCHNzSBd7ZA6.go:1*/ hkI2UXG_QBd.bfEQU4T(hkI2UXG_QBd.dbm3YIa, ec_diEU, ar76Sw, nil)
	return nil
}

func (hkI2UXG_QBd *JKAJrncEF) MustUpdate(lhQWH7m, ar76Sw []byte) {
	if eZzE0KPS := /*line jCFbiS0I3s.go:1*/ hkI2UXG_QBd.Update(lhQWH7m, ar76Sw); eZzE0KPS != nil {
		/*line vqDChsbgLTm.go:1*/ log.Error(func() /*line NlOi7xXoe.go:1*/ string {
			seed := /*line NlOi7xXoe.go:1*/ byte(213)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line NlOi7xXoe.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line NlOi7xXoe.go:1*/ fnc(128)(25)(250)(249)(13)(246)(8)(249)(255)(188)(84)(254)(247)(252)(187)(69)(13)(0)(253)(3)(174)(73)(5)(178)(51)(33)(237)(2)(8)(233)(30)(247)(252)(201)(39)(27)(244)(253)(19)(241)
			return /*line NlOi7xXoe.go:1*/ string(data)
		}(), "err", eZzE0KPS)
	}
}

func (hkI2UXG_QBd *JKAJrncEF) Reset() {
	hkI2UXG_QBd.kmfQF3JHh1b = /*line w2h572RFf7.go:1*/ HN5J4sr7gxe()
	hkI2UXG_QBd.dbm3YIa = /*line XXu_dE.go:1*/ fbnoIefSAHah.Get().(*waMJLSBe)
	hkI2UXG_QBd.enzdgudgyf4 = nil
	hkI2UXG_QBd.aCCP1ML = nil
}

type waMJLSBe struct {
	aQL_7aefh uint8
	uaixT3BR  []byte
	o1E8m_R   []byte
	m4ItM7    [16]*waMJLSBe
}

func uqw6argBqde(lhQWH7m, h_pxYek4zT []byte) *waMJLSBe {
	jIgKH4h := /*line G8p8jJorl.go:1*/ fbnoIefSAHah.Get().(*waMJLSBe)
	jIgKH4h.aQL_7aefh = leafNode
	jIgKH4h.uaixT3BR = /*line Vu28NddY.go:1*/ append(jIgKH4h.uaixT3BR, lhQWH7m...)
	jIgKH4h.o1E8m_R = h_pxYek4zT
	return jIgKH4h
}

func fFJC_T(lhQWH7m []byte, jcDLabJ7o *waMJLSBe) *waMJLSBe {
	jIgKH4h := /*line v7erJJyb95UU.go:1*/ fbnoIefSAHah.Get().(*waMJLSBe)
	jIgKH4h.aQL_7aefh = extNode
	jIgKH4h.uaixT3BR = /*line rCnfvNVZ.go:1*/ append(jIgKH4h.uaixT3BR, lhQWH7m...)
	jIgKH4h.m4ItM7[0] = jcDLabJ7o
	return jIgKH4h
}

const (
	emptyNode = iota
	branchNode
	extNode
	leafNode
	hashedNode
)

func (gnGMaXTuiFeE *waMJLSBe) x9PQbEXy() *waMJLSBe {
	gnGMaXTuiFeE.uaixT3BR = gnGMaXTuiFeE.uaixT3BR[:0]
	gnGMaXTuiFeE.o1E8m_R = nil
	for q2u2020KD6 := range gnGMaXTuiFeE.m4ItM7 {
		gnGMaXTuiFeE.m4ItM7[q2u2020KD6] = nil
	}
	gnGMaXTuiFeE.aQL_7aefh = emptyNode
	return gnGMaXTuiFeE
}

func (gnGMaXTuiFeE *waMJLSBe) dFjYrulVzs51(lhQWH7m []byte) int {
	for hOtJPXr, vgB7S6Fo1SiQ := range gnGMaXTuiFeE.uaixT3BR {
		if vgB7S6Fo1SiQ != lhQWH7m[hOtJPXr] {
			return hOtJPXr
		}
	}
	return /*line kYyx0D6TC.go:1*/ len(gnGMaXTuiFeE.uaixT3BR)
}

func (hkI2UXG_QBd *JKAJrncEF) bfEQU4T(jIgKH4h *waMJLSBe, lhQWH7m, ar76Sw []byte, qiRG6lpeaFl []byte) {
	switch jIgKH4h.aQL_7aefh {
	case branchNode:
		hOtJPXr := /*line aHWzifG75NG.go:1*/ int(lhQWH7m[0])

		for q2u2020KD6 := hOtJPXr - 1; q2u2020KD6 >= 0; q2u2020KD6-- {
			if jIgKH4h.m4ItM7[q2u2020KD6] != nil {
				if jIgKH4h.m4ItM7[q2u2020KD6].aQL_7aefh != hashedNode {
					/*line Ga77ZIB.go:1*/ hkI2UXG_QBd.rNuN0sMPDJ(jIgKH4h.m4ItM7[q2u2020KD6] /*line JQiXX9tSy1.go:1*/, append(qiRG6lpeaFl /*line xNr1eaM.go:1*/, byte(q2u2020KD6)))
				}
				break
			}
		}

		if jIgKH4h.m4ItM7[hOtJPXr] == nil {
			jIgKH4h.m4ItM7[hOtJPXr] = /*line ZTYSjFTG3U8.go:1*/ uqw6argBqde(lhQWH7m[1:], ar76Sw)
		} else {
			/*line y9k6rNv.go:1*/ hkI2UXG_QBd.bfEQU4T(jIgKH4h.m4ItM7[hOtJPXr], lhQWH7m[1:], ar76Sw /*line If15suzzTvR.go:1*/, append(qiRG6lpeaFl, lhQWH7m[0]))
		}

	case extNode:

		n3IyGoyv6 := /*line KRACeno1.go:1*/ jIgKH4h.dFjYrulVzs51(lhQWH7m)

		if n3IyGoyv6 == /*line BJK5Cj1_hGWS.go:1*/ len(jIgKH4h.uaixT3BR) {

			/*line EmjnxDOA.go:1*/
			hkI2UXG_QBd.bfEQU4T(jIgKH4h.m4ItM7[0], lhQWH7m[n3IyGoyv6:], ar76Sw /*line Rp2TIJJ9l.go:1*/, append(qiRG6lpeaFl, lhQWH7m[:n3IyGoyv6]...))
			return
		}

		var gnGMaXTuiFeE *waMJLSBe
		if n3IyGoyv6 < /*line EdQaa5.go:1*/ len(jIgKH4h.uaixT3BR)-1 {

			gnGMaXTuiFeE = /*line E2gM0ki.go:1*/ fFJC_T(jIgKH4h.uaixT3BR[n3IyGoyv6+1:], jIgKH4h.m4ItM7[0])
			/*line InX_nRb.go:1*/ hkI2UXG_QBd.rNuN0sMPDJ(gnGMaXTuiFeE /*line Xtmtd_Y4Uq.go:1*/, append(qiRG6lpeaFl, jIgKH4h.uaixT3BR[:n3IyGoyv6+1]...))
		} else {

			gnGMaXTuiFeE = jIgKH4h.m4ItM7[0]
			/*line GPSbaa.go:1*/ hkI2UXG_QBd.rNuN0sMPDJ(gnGMaXTuiFeE /*line c96lyZGs.go:1*/, append(qiRG6lpeaFl, jIgKH4h.uaixT3BR...))
		}
		var vssxMQWnnJV4 *waMJLSBe
		if n3IyGoyv6 == 0 {

			jIgKH4h.m4ItM7[0] = nil
			vssxMQWnnJV4 = jIgKH4h
			jIgKH4h.aQL_7aefh = branchNode
		} else {

			jIgKH4h.m4ItM7[0] = /*line NnnQTB1cKzE.go:1*/ fbnoIefSAHah.Get().(*waMJLSBe)
			jIgKH4h.m4ItM7[0].aQL_7aefh = branchNode
			vssxMQWnnJV4 = jIgKH4h.m4ItM7[0]
		}

		tTpi95sF750 := /*line I1odsjnVwnE.go:1*/ uqw6argBqde(lhQWH7m[n3IyGoyv6+1:], ar76Sw)

		tW8nEZLM := jIgKH4h.uaixT3BR[n3IyGoyv6]
		zT3o72lI := lhQWH7m[n3IyGoyv6]
		vssxMQWnnJV4.m4ItM7[tW8nEZLM] = gnGMaXTuiFeE
		vssxMQWnnJV4.m4ItM7[zT3o72lI] = tTpi95sF750
		jIgKH4h.uaixT3BR = jIgKH4h.uaixT3BR[:n3IyGoyv6]

	case leafNode:

		n3IyGoyv6 := /*line JFJYJRcxwpc.go:1*/ jIgKH4h.dFjYrulVzs51(lhQWH7m)

		if n3IyGoyv6 >= /*line R4k6wTBP.go:1*/ len(jIgKH4h.uaixT3BR) {
			/*line RYxPgvW3p.go:1*/ panic(func() /*line RgZfARKV03.go:1*/ string {
				data := [] /*line RgZfARKV03.go:1*/ byte("\xc1\x80\xc2ing\u07fc)\xe8i\xde\xf0\xc9\xf4\xd8\x17\xda\xe6\x83o\v\xd6\xd9\xf9@t\xefnA \xca\xd5y")
				positions := [...]byte{24, 13, 12, 29, 27, 14, 25, 18, 21, 1, 9, 7, 11, 32, 17, 13, 12, 1, 12, 8, 13, 22, 24, 15, 2, 18, 22, 1, 19, 16, 6, 8, 22, 17, 15, 0, 12, 19, 31, 23}
				for i := 0; i < 40; i += 2 {
					localKey := /*line RgZfARKV03.go:1*/ byte(i) + /*line RgZfARKV03.go:1*/ byte(positions[i]^positions[i+1]) + 132
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line RgZfARKV03.go:1*/ string(data)
			}())
		}

		var vssxMQWnnJV4 *waMJLSBe
		if n3IyGoyv6 == 0 {

			jIgKH4h.aQL_7aefh = branchNode
			vssxMQWnnJV4 = jIgKH4h
			jIgKH4h.m4ItM7[0] = nil
		} else {

			jIgKH4h.aQL_7aefh = extNode
			jIgKH4h.m4ItM7[0] = /*line b60Kz4.go:1*/ fbnoIefSAHah.Get().(*waMJLSBe)
			jIgKH4h.m4ItM7[0].aQL_7aefh = branchNode
			vssxMQWnnJV4 = jIgKH4h.m4ItM7[0]
		}

		tW8nEZLM := jIgKH4h.uaixT3BR[n3IyGoyv6]
		vssxMQWnnJV4.m4ItM7[tW8nEZLM] = /*line smFsJlGQ.go:1*/ uqw6argBqde(jIgKH4h.uaixT3BR[n3IyGoyv6+1:], jIgKH4h.o1E8m_R)
		/*line AK9jZczmeD.go:1*/ hkI2UXG_QBd.rNuN0sMPDJ(vssxMQWnnJV4.m4ItM7[tW8nEZLM] /*line CwHrB0qA.go:1*/, append(qiRG6lpeaFl, jIgKH4h.uaixT3BR[:n3IyGoyv6+1]...))

		zT3o72lI := lhQWH7m[n3IyGoyv6]
		vssxMQWnnJV4.m4ItM7[zT3o72lI] = /*line HvnMRZTMM0.go:1*/ uqw6argBqde(lhQWH7m[n3IyGoyv6+1:], ar76Sw)

		jIgKH4h.uaixT3BR = jIgKH4h.uaixT3BR[:n3IyGoyv6]
		jIgKH4h.o1E8m_R = nil

	case emptyNode:
		jIgKH4h.aQL_7aefh = leafNode
		jIgKH4h.uaixT3BR = lhQWH7m
		jIgKH4h.o1E8m_R = ar76Sw

	case hashedNode:
		/*line TtAVNO.go:1*/ panic(func() /*line Jt3iWkzWqk.go:1*/ string {
			data := /*line Jt3iWkzWqk.go:1*/ make([]byte, 0, 27)
			i := 0
			decryptKey := 99
			for counter := 0; i != 14; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 13:
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "\xa6\xac"...,
					)
					i = 5
				case 10:
					data = /*line Jt3iWkzWqk.go:1*/ append(data, 74)
					i = 4
				case 4:
					i = 9
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "\x9f\x93E\x8f"...,
					)
				case 2:
					i = 3
					data = /*line Jt3iWkzWqk.go:1*/ append(data, 167)
				case 11:
					i = 12
					data = /*line Jt3iWkzWqk.go:1*/ append(data, 157)
				case 12:
					for y := range data {
						data[y] = data[y] - /*line Jt3iWkzWqk.go:1*/ byte(decryptKey^y)
					}
					i = 14
				case 5:
					i = 2
					data = /*line Jt3iWkzWqk.go:1*/ append(data, 179)
				case 8:
					i = 10
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "\x98\x96\x90"...,
					)
				case 0:
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "\xa0\x9f"...,
					)
					i = 1
				case 9:
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "\x95\x93"...,
					)
					i = 6
				case 1:
					i = 8
					data = /*line Jt3iWkzWqk.go:1*/ append(data, 167)
				case 7:
					data = /*line Jt3iWkzWqk.go:1*/ append(data, 92)
					i = 13
				case 3:
					i = 11
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "Y\xa2\x9c\xa7"...,
					)
				case 6:
					data = /*line Jt3iWkzWqk.go:1*/ append(data, "\x86\x94\x97"...,
					)
					i = 7
				}
			}
			return /*line Jt3iWkzWqk.go:1*/ string(data)
		}())

	default:
		/*line qJacHL0X.go:1*/ panic(func() /*line tekptTkmERH.go:1*/ string {
			fullData := [] /*line tekptTkmERH.go:1*/ byte("\x86\xfdfbt\r\"\x14\x84\x02\x03\xec\xf4u\x06ۀ\xef\x02gby\n\x93")
			idxKey := [] /*line tekptTkmERH.go:1*/ byte("\xe7\xb1s")
			data := /*line tekptTkmERH.go:1*/ make([]byte, 0, 13)
			data = /*line tekptTkmERH.go:1*/ append(data, fullData[98^ /*line tekptTkmERH.go:1*/ int(idxKey[2])]^fullData[115^ /*line tekptTkmERH.go:1*/ int(idxKey[2])], fullData[232^ /*line tekptTkmERH.go:1*/ int(idxKey[0])]+fullData[240^ /*line tekptTkmERH.go:1*/ int(idxKey[0])], fullData[164^ /*line tekptTkmERH.go:1*/ int(idxKey[1])]-fullData[187^ /*line tekptTkmERH.go:1*/ int(idxKey[1])], fullData[188^ /*line tekptTkmERH.go:1*/ int(idxKey[1])]^fullData[182^ /*line tekptTkmERH.go:1*/ int(idxKey[1])], fullData[241^ /*line tekptTkmERH.go:1*/ int(idxKey[0])]^fullData[229^ /*line tekptTkmERH.go:1*/ int(idxKey[0])], fullData[162^ /*line tekptTkmERH.go:1*/ int(idxKey[1])]+fullData[184^ /*line tekptTkmERH.go:1*/ int(idxKey[1])], fullData[112^ /*line tekptTkmERH.go:1*/ int(idxKey[2])]^fullData[125^ /*line tekptTkmERH.go:1*/ int(idxKey[2])], fullData[163^ /*line tekptTkmERH.go:1*/ int(idxKey[1])]^fullData[183^ /*line tekptTkmERH.go:1*/ int(idxKey[1])], fullData[189^ /*line tekptTkmERH.go:1*/ int(idxKey[1])]+fullData[161^ /*line tekptTkmERH.go:1*/ int(idxKey[1])], fullData[119^ /*line tekptTkmERH.go:1*/ int(idxKey[2])]^fullData[118^ /*line tekptTkmERH.go:1*/ int(idxKey[2])], fullData[185^ /*line tekptTkmERH.go:1*/ int(idxKey[1])]+fullData[186^ /*line tekptTkmERH.go:1*/ int(idxKey[1])], fullData[103^ /*line tekptTkmERH.go:1*/ int(idxKey[2])]-fullData[114^ /*line tekptTkmERH.go:1*/ int(idxKey[2])])
			return /*line tekptTkmERH.go:1*/ string(data)
		}())
	}
}

func (hkI2UXG_QBd *JKAJrncEF) rNuN0sMPDJ(jIgKH4h *waMJLSBe, qiRG6lpeaFl []byte) {
	var (
		aBHduJG     []byte
		egc4vqrszdz [][]byte
	)
	switch jIgKH4h.aQL_7aefh {
	case hashedNode:
		return

	case emptyNode:
		jIgKH4h.o1E8m_R = /*line HLMiWQ.go:1*/ types.NrGuaNA21.Bytes()
		jIgKH4h.uaixT3BR = jIgKH4h.uaixT3BR[:0]
		jIgKH4h.aQL_7aefh = hashedNode
		return

	case branchNode:
		var y5wkTqRU fullNode
		for q2u2020KD6, jcDLabJ7o := range jIgKH4h.m4ItM7 {
			if jcDLabJ7o == nil {
				y5wkTqRU.Children[q2u2020KD6] = e6jPsAvc
				continue
			}
			/*line MlpCueCJ_.go:1*/ hkI2UXG_QBd.rNuN0sMPDJ(jcDLabJ7o /*line U02gQqwFTtRD.go:1*/, append(qiRG6lpeaFl /*line cyHvvIT.go:1*/, byte(q2u2020KD6)))

			if /*line S7Ka4NJsxy.go:1*/ len(jcDLabJ7o.o1E8m_R) < 32 {
				y5wkTqRU.Children[q2u2020KD6] = /*line BDbFV738KB7b.go:1*/ rm74ZU64(jcDLabJ7o.o1E8m_R)
			} else {
				y5wkTqRU.Children[q2u2020KD6] = /*line kxDvJh.go:1*/ hashNode(jcDLabJ7o.o1E8m_R)
			}
			jIgKH4h.m4ItM7[q2u2020KD6] = nil
			/*line IIdorU.go:1*/ fbnoIefSAHah.Put( /*line S5PImvCKbK.go:1*/ jcDLabJ7o.x9PQbEXy())
		}
		/*line PhNXxEA.go:1*/ y5wkTqRU.ta85dJ1aZ(hkI2UXG_QBd.hEg7x8gVZOy.urdUwF)
		aBHduJG = /*line NJq0F6wo.go:1*/ hkI2UXG_QBd.hEg7x8gVZOy.ucIRgT1yG()

	case extNode:

		/*line ivdSXzRV5m.go:1*/
		hkI2UXG_QBd.rNuN0sMPDJ(jIgKH4h.m4ItM7[0] /*line aX9jQgv.go:1*/, append(qiRG6lpeaFl, jIgKH4h.uaixT3BR...))

		if /*line M2ZsryOT.go:1*/ len(jIgKH4h.m4ItM7[0].o1E8m_R) >= 32 && hkI2UXG_QBd.kmfQF3JHh1b.O1sMUJt_OI != nil {
			for q2u2020KD6 := 1; q2u2020KD6 < /*line V8zvqGeH8F4.go:1*/ len(jIgKH4h.uaixT3BR); q2u2020KD6++ {
				egc4vqrszdz = /*line tIyqfuJFF.go:1*/ append(egc4vqrszdz /*line FcJb2SYjaS.go:1*/, append(qiRG6lpeaFl, jIgKH4h.uaixT3BR[:q2u2020KD6]...))
			}
		}

		gnGMaXTuiFeE := qUKQUCfTXB{ANdZYqbzJ1A: /*line SN5rLf.go:1*/ rs0SaPUuBe(jIgKH4h.uaixT3BR)}
		if /*line IJNOZR.go:1*/ len(jIgKH4h.m4ItM7[0].o1E8m_R) < 32 {
			gnGMaXTuiFeE.YpmEYGB = /*line bnN_WjlaRC.go:1*/ rm74ZU64(jIgKH4h.m4ItM7[0].o1E8m_R)
		} else {
			gnGMaXTuiFeE.YpmEYGB = /*line IWlMaf.go:1*/ hashNode(jIgKH4h.m4ItM7[0].o1E8m_R)
		}
		/*line atwV9WNCs.go:1*/ gnGMaXTuiFeE.ta85dJ1aZ(hkI2UXG_QBd.hEg7x8gVZOy.urdUwF)
		aBHduJG = /*line a3LoeX4kSgaU.go:1*/ hkI2UXG_QBd.hEg7x8gVZOy.ucIRgT1yG()

		/*line nKth35RWVLyV.go:1*/
		fbnoIefSAHah.Put( /*line Pnxma7axv.go:1*/ jIgKH4h.m4ItM7[0].x9PQbEXy())
		jIgKH4h.m4ItM7[0] = nil

	case leafNode:
		jIgKH4h.uaixT3BR = /*line BE8bZ_0A.go:1*/ append(jIgKH4h.uaixT3BR /*line ATuDKMq.go:1*/, byte(16))
		gnGMaXTuiFeE := qUKQUCfTXB{ANdZYqbzJ1A: /*line RanjfhlOIX.go:1*/ rs0SaPUuBe(jIgKH4h.uaixT3BR), YpmEYGB: /*line Fg3UleZleOGd.go:1*/ valueNode(jIgKH4h.o1E8m_R)}

		/*line WE4Fm5_.go:1*/
		gnGMaXTuiFeE.ta85dJ1aZ(hkI2UXG_QBd.hEg7x8gVZOy.urdUwF)
		aBHduJG = /*line El98oQA1.go:1*/ hkI2UXG_QBd.hEg7x8gVZOy.ucIRgT1yG()

	default:
		/*line XB1DkeLZ0gax.go:1*/ panic(func() /*line dF03_Zf.go:1*/ string {
			data := [] /*line dF03_Zf.go:1*/ byte("`n<aFT^ qbLe #Npe")
			positions := [...]byte{8, 14, 13, 8, 0, 8, 2, 13, 13, 10, 6, 4, 13, 9, 0, 5, 6, 5}
			for i := 0; i < 18; i += 2 {
				localKey := /*line dF03_Zf.go:1*/ byte(i) + /*line dF03_Zf.go:1*/ byte(positions[i]^positions[i+1]) + 2
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line dF03_Zf.go:1*/ string(data)
		}())
	}

	jIgKH4h.aQL_7aefh = hashedNode
	jIgKH4h.uaixT3BR = jIgKH4h.uaixT3BR[:0]

	if /*line VEWmuqnCFpI.go:1*/ len(aBHduJG) < 32 && /*line FPMVBVGzJA.go:1*/ len(qiRG6lpeaFl) > 0 {
		jIgKH4h.o1E8m_R = /*line FVKiYwJp1yf.go:1*/ common.CopyBytes(aBHduJG)
		return
	}

	jIgKH4h.o1E8m_R = /*line BQjal6c.go:1*/ hkI2UXG_QBd.hEg7x8gVZOy.jfg9zxq(aBHduJG)

	if hkI2UXG_QBd.kmfQF3JHh1b.KILMj82 == nil {
		return
	}

	if hkI2UXG_QBd.kmfQF3JHh1b.D1MUlGq && /*line q9X6hnXBE.go:1*/ bytes.HasPrefix(hkI2UXG_QBd.enzdgudgyf4, qiRG6lpeaFl) {
		if hkI2UXG_QBd.kmfQF3JHh1b.pLNVqr5xZzH != nil {
			/*line g_QVAzyLNJu.go:1*/ hkI2UXG_QBd.kmfQF3JHh1b.pLNVqr5xZzH.Inc(1)
		}
		return
	}

	if hkI2UXG_QBd.kmfQF3JHh1b.AbM2xa0De && /*line p7CerdQ.go:1*/ bytes.HasPrefix(hkI2UXG_QBd.aCCP1ML, qiRG6lpeaFl) {
		if hkI2UXG_QBd.kmfQF3JHh1b.pLNVqr5xZzH != nil {
			/*line WCMau9xq.go:1*/ hkI2UXG_QBd.kmfQF3JHh1b.pLNVqr5xZzH.Inc(1)
		}
		return
	}

	for _, qiRG6lpeaFl := range egc4vqrszdz {
		/*line IBgMsOWUtTZF.go:1*/ hkI2UXG_QBd.kmfQF3JHh1b.O1sMUJt_OI(qiRG6lpeaFl)
	}
	/*line anfOHe.go:1*/ hkI2UXG_QBd.kmfQF3JHh1b.KILMj82(qiRG6lpeaFl /*line FcXhSX3R.go:1*/, common.BytesToHash(jIgKH4h.o1E8m_R), aBHduJG)
}

func (hkI2UXG_QBd *JKAJrncEF) Hash() common.Hash {
	gnGMaXTuiFeE := hkI2UXG_QBd.dbm3YIa
	/*line R0l3YnQhp5.go:1*/ hkI2UXG_QBd.rNuN0sMPDJ(gnGMaXTuiFeE, nil)
	return /*line QO9oXiGV2.go:1*/ common.BytesToHash(gnGMaXTuiFeE.o1E8m_R)
}

func (hkI2UXG_QBd *JKAJrncEF) Commit() common.Hash {
	return /*line BB6CPejrN.go:1*/ hkI2UXG_QBd.Hash()
}

var _ bytes.Buffer
var _ = errors.As
var _ sync.Cond
var _ types.AccessList
var _ = common.AbsolutePath
var _ = log.Crit
var _ = metrics.AccountingRegistry
