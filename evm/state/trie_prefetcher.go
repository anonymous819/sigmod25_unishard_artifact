//line :1
package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	iiuFwIcXNeK = func() /*line qlzLzE.go:1*/ string {
		fullData := [] /*line qlzLzE.go:1*/ byte("3\x8fĭ\x94\x8eP@\x01\xa7S\xad\xb8\u05f8#E\x04b\xe7%\xb4\xb6\xa6\xef\xce\n\xde")
		idxKey := [] /*line qlzLzE.go:1*/ byte("病\xa7R\xc3U$_\xa0")
		data := /*line qlzLzE.go:1*/ make([]byte, 0, 15)
		data = /*line qlzLzE.go:1*/ append(data, fullData[240^ /*line qlzLzE.go:1*/ int(idxKey[0])]+fullData[254^ /*line qlzLzE.go:1*/ int(idxKey[0])], fullData[78^ /*line qlzLzE.go:1*/ int(idxKey[6])]+fullData[81^ /*line qlzLzE.go:1*/ int(idxKey[6])], fullData[34^ /*line qlzLzE.go:1*/ int(idxKey[7])]-fullData[55^ /*line qlzLzE.go:1*/ int(idxKey[7])], fullData[89^ /*line qlzLzE.go:1*/ int(idxKey[4])]+fullData[92^ /*line qlzLzE.go:1*/ int(idxKey[4])], fullData[189^ /*line qlzLzE.go:1*/ int(idxKey[3])]+fullData[179^ /*line qlzLzE.go:1*/ int(idxKey[3])], fullData[165^ /*line qlzLzE.go:1*/ int(idxKey[3])]^fullData[178^ /*line qlzLzE.go:1*/ int(idxKey[3])], fullData[87^ /*line qlzLzE.go:1*/ int(idxKey[8])]-fullData[94^ /*line qlzLzE.go:1*/ int(idxKey[8])], fullData[95^ /*line qlzLzE.go:1*/ int(idxKey[4])]+fullData[87^ /*line qlzLzE.go:1*/ int(idxKey[4])], fullData[71^ /*line qlzLzE.go:1*/ int(idxKey[6])]+fullData[68^ /*line qlzLzE.go:1*/ int(idxKey[6])], fullData[40^ /*line qlzLzE.go:1*/ int(idxKey[7])]+fullData[39^ /*line qlzLzE.go:1*/ int(idxKey[7])], fullData[158^ /*line qlzLzE.go:1*/ int(idxKey[1])]-fullData[151^ /*line qlzLzE.go:1*/ int(idxKey[1])], fullData[68^ /*line qlzLzE.go:1*/ int(idxKey[4])]-fullData[88^ /*line qlzLzE.go:1*/ int(idxKey[4])], fullData[211^ /*line qlzLzE.go:1*/ int(idxKey[5])]+fullData[204^ /*line qlzLzE.go:1*/ int(idxKey[5])], fullData[88^ /*line qlzLzE.go:1*/ int(idxKey[8])]+fullData[71^ /*line qlzLzE.go:1*/ int(idxKey[8])])
		return /*line qlzLzE.go:1*/ string(data)
	}()
)

type triePrefetcher struct {
	db       Database
	root     common.Hash
	fetches  map[string]Trie
	fetchers map[string]*subfetcher

	deliveryMissMeter metrics.Meter
	accountLoadMeter  metrics.Meter
	accountDupMeter   metrics.Meter
	accountSkipMeter  metrics.Meter
	accountWasteMeter metrics.Meter
	storageLoadMeter  metrics.Meter
	storageDupMeter   metrics.Meter
	storageSkipMeter  metrics.Meter
	storageWasteMeter metrics.Meter
}

func vnGIvB(tC4SGV Database, lxItfhw common.Hash, _6LsfLXKG string) *triePrefetcher {
	eencCjrsCQq := iiuFwIcXNeK + _6LsfLXKG
	gqUt2ia7B := &triePrefetcher{
		db:   tC4SGV,
		root: lxItfhw,
		fetchers:/*line HrRGwf.go:1*/ make(map[string]*subfetcher),

		deliveryMissMeter: /*line GiV8Um.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line a_0C96qO.go:1*/ string {
			key := [] /*line a_0C96qO.go:1*/ byte("vF\x98\x15\xd0ɽ\x93\xa6Kw2\xe6")
			data := [] /*line a_0C96qO.go:1*/ byte("\xa5\xaa\xfd\x819?\"\x05\x1f\xb8\xe0\xa5Y")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line a_0C96qO.go:1*/ string(data)
		}(), nil),
		accountLoadMeter: /*line ooe0AoaH9.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line d_bDPVseB.go:1*/ string {
			data := /*line d_bDPVseB.go:1*/ make([]byte, 0, 14)
			i := 2
			decryptKey := 74
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 1
					data = /*line d_bDPVseB.go:1*/ append(data, "\xae\xbb"...,
					)
				case 2:
					i = 0
					data = /*line d_bDPVseB.go:1*/ append(data, "w\xaa\xad"...,
					)
				case 3:
					i = 5
					for y := range data {
						data[y] = data[y] - /*line d_bDPVseB.go:1*/ byte(decryptKey^y)
					}
				case 1:
					data = /*line d_bDPVseB.go:1*/ append(data, "¼\xc3o"...,
					)
					i = 4
				case 4:
					i = 3
					data = /*line d_bDPVseB.go:1*/ append(data, "\xad\xb1\xa4\xa8"...,
					)
				}
			}
			return /*line d_bDPVseB.go:1*/ string(data)
		}(), nil),
		accountDupMeter: /*line uWTadH8tO.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line G0oGvoy3cag.go:1*/ string {
			seed := /*line G0oGvoy3cag.go:1*/ byte(103)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line G0oGvoy3cag.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line G0oGvoy3cag.go:1*/ fnc(200)(50)(2)(0)(12)(6)(249)(6)(187)(53)(17)(251)
			return /*line G0oGvoy3cag.go:1*/ string(data)
		}(), nil),
		accountSkipMeter: /*line evymcAnRt.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line AEEHSjdQW.go:1*/ string {
			seed := /*line AEEHSjdQW.go:1*/ byte(67)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line AEEHSjdQW.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line AEEHSjdQW.go:1*/ fnc(108)(206)(30)(248)(252)(250)(231)(4)(91)(188)(224)(2)(29)
			return /*line AEEHSjdQW.go:1*/ string(data)
		}(), nil),
		accountWasteMeter: /*line oKitejM.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line J3GPhJ.go:1*/ string {
			key := [] /*line J3GPhJ.go:1*/ byte("\x04\xd5U\xa50T\xd2\x1f2\x06\xd8\x03\xae\xbe")
			data := [] /*line J3GPhJ.go:1*/ byte("36\xb8\b\x9f\xc9@\x93a}9v\"#")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line J3GPhJ.go:1*/ string(data)
		}(), nil),
		storageLoadMeter: /*line BFrUU8PBDMV.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line opIKmC.go:1*/ string {
			key := [] /*line opIKmC.go:1*/ byte("\xe6O\xbfI\xad\x82\xbdQ\xdc\xde\xc1»")
			data := [] /*line opIKmC.go:1*/ byte("\x15\xc23\xb8\x1f\xe3$\xb6\vJ0#\x1f")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line opIKmC.go:1*/ string(data)
		}(), nil),
		storageDupMeter: /*line q3CLqXZ8.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line ofLuS2.go:1*/ string {
			data := [] /*line ofLuS2.go:1*/ byte("\xee\xbe\xe2Dra\xdd\xed/\xf2\xc3\x14")
			positions := [...]byte{3, 1, 6, 3, 10, 11, 6, 1, 0, 10, 7, 11, 9, 2, 7, 7}
			for i := 0; i < 16; i += 2 {
				localKey := /*line ofLuS2.go:1*/ byte(i) + /*line ofLuS2.go:1*/ byte(positions[i]^positions[i+1]) + 103
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line ofLuS2.go:1*/ string(data)
		}(), nil),
		storageSkipMeter: /*line qneVeKx5oev.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line uYGeIIits.go:1*/ string {
			key := [] /*line uYGeIIits.go:1*/ byte("\x9d@\x02M\x84\nKS\xcd\x1d\xec\x15M")
			data := [] /*line uYGeIIits.go:1*/ byte("̳v\xbc\xf6k\xb2\xb8\xfc\x90W~\xbd")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line uYGeIIits.go:1*/ string(data)
		}(), nil),
		storageWasteMeter: /*line wAYA6R.go:1*/ metrics.GetOrRegisterMeter(eencCjrsCQq+func() /*line FQrNMXdVObph.go:1*/ string {
			fullData := [] /*line FQrNMXdVObph.go:1*/ byte("̝A\v\x84&\xf0\xfd\xfb<:j\xef\x9fo\x04\xf5\xa3\x05\x8bk\xd3zx9$\x92\x9a")
			idxKey := [] /*line FQrNMXdVObph.go:1*/ byte("x\xbe\xb1\xe7\x01\xf2\x03[2\xc8?")
			data := /*line FQrNMXdVObph.go:1*/ make([]byte, 0, 15)
			data = /*line FQrNMXdVObph.go:1*/ append(data, fullData[200^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])]-fullData[201^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])], fullData[96^ /*line FQrNMXdVObph.go:1*/ int(idxKey[0])]+fullData[114^ /*line FQrNMXdVObph.go:1*/ int(idxKey[0])], fullData[252^ /*line FQrNMXdVObph.go:1*/ int(idxKey[5])]-fullData[250^ /*line FQrNMXdVObph.go:1*/ int(idxKey[5])], fullData[24^ /*line FQrNMXdVObph.go:1*/ int(idxKey[6])]^fullData[19^ /*line FQrNMXdVObph.go:1*/ int(idxKey[6])], fullData[20^ /*line FQrNMXdVObph.go:1*/ int(idxKey[4])]+fullData[12^ /*line FQrNMXdVObph.go:1*/ int(idxKey[4])], fullData[80^ /*line FQrNMXdVObph.go:1*/ int(idxKey[7])]^fullData[88^ /*line FQrNMXdVObph.go:1*/ int(idxKey[7])], fullData[76^ /*line FQrNMXdVObph.go:1*/ int(idxKey[7])]+fullData[87^ /*line FQrNMXdVObph.go:1*/ int(idxKey[7])], fullData[219^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])]-fullData[205^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])], fullData[220^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])]-fullData[193^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])], fullData[222^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])]+fullData[207^ /*line FQrNMXdVObph.go:1*/ int(idxKey[9])], fullData[14^ /*line FQrNMXdVObph.go:1*/ int(idxKey[4])]-fullData[16^ /*line FQrNMXdVObph.go:1*/ int(idxKey[4])], fullData[245^ /*line FQrNMXdVObph.go:1*/ int(idxKey[3])]-fullData[253^ /*line FQrNMXdVObph.go:1*/ int(idxKey[3])], fullData[184^ /*line FQrNMXdVObph.go:1*/ int(idxKey[1])]^fullData[186^ /*line FQrNMXdVObph.go:1*/ int(idxKey[1])], fullData[235^ /*line FQrNMXdVObph.go:1*/ int(idxKey[5])]+fullData[240^ /*line FQrNMXdVObph.go:1*/ int(idxKey[5])])
			return /*line FQrNMXdVObph.go:1*/ string(data)
		}(), nil),
	}
	return gqUt2ia7B
}

func (gqUt2ia7B *triePrefetcher) gfibELOKGAub() {
	for _, eSIabWY := range gqUt2ia7B.fetchers {
		/*line ZJg4uOQWc9.go:1*/ eSIabWY.tj1faB2CfMY8()

		if metrics.Enabled {
			if eSIabWY.root == gqUt2ia7B.root {
				/*line nlaEOU1CZ.go:1*/ gqUt2ia7B.accountLoadMeter.Mark( /*line RgFpg7OvV.go:1*/ int64( /*line chq8jXhMy.go:1*/ len(eSIabWY.seen)))
				/*line ixWrTf3LOxk4.go:1*/ gqUt2ia7B.accountDupMeter.Mark( /*line Lq2dqtGmK.go:1*/ int64(eSIabWY.dups))
				/*line ecaJlgmRY.go:1*/ gqUt2ia7B.accountSkipMeter.Mark( /*line c_38C7VIh.go:1*/ int64( /*line IdCaxwYg8.go:1*/ len(eSIabWY.tasks)))

				for _, ixeygbAgOap := range eSIabWY.used {
					/*line lQG3yui.go:1*/ delete(eSIabWY.seen /*line Ga6GJyhDYk.go:1*/, string(ixeygbAgOap))
				}
				/*line w2BP6ni.go:1*/ gqUt2ia7B.accountWasteMeter.Mark( /*line GLRbdl.go:1*/ int64( /*line d0NXzKdl8wU5.go:1*/ len(eSIabWY.seen)))
			} else {
				/*line IQNxJU.go:1*/ gqUt2ia7B.storageLoadMeter.Mark( /*line O2BkAmXOW5R9.go:1*/ int64( /*line vtJEbk29DPip.go:1*/ len(eSIabWY.seen)))
				/*line xS78UtM.go:1*/ gqUt2ia7B.storageDupMeter.Mark( /*line ET4c_gs9AwNg.go:1*/ int64(eSIabWY.dups))
				/*line fmHFAUAOpmr5.go:1*/ gqUt2ia7B.storageSkipMeter.Mark( /*line FP1_ex0M5S.go:1*/ int64( /*line oCbRqOFmTQc.go:1*/ len(eSIabWY.tasks)))

				for _, ixeygbAgOap := range eSIabWY.used {
					/*line sE3Q4Gf5.go:1*/ delete(eSIabWY.seen /*line zwWh7ky7.go:1*/, string(ixeygbAgOap))
				}
				/*line y1pB22a_82w.go:1*/ gqUt2ia7B.storageWasteMeter.Mark( /*line ldJkYaYBwB.go:1*/ int64( /*line Pttkl4mx.go:1*/ len(eSIabWY.seen)))
			}
		}
	}

	gqUt2ia7B.fetchers = nil
}

func (gqUt2ia7B *triePrefetcher) stapkw2aW() *triePrefetcher {
	stapkw2aW := &triePrefetcher{
		db:   gqUt2ia7B.db,
		root: gqUt2ia7B.root,
		fetches:/*line FvWMfLcDASG.go:1*/ make(map[string]Trie),

		deliveryMissMeter: gqUt2ia7B.deliveryMissMeter,
		accountLoadMeter:  gqUt2ia7B.accountLoadMeter,
		accountDupMeter:   gqUt2ia7B.accountDupMeter,
		accountSkipMeter:  gqUt2ia7B.accountSkipMeter,
		accountWasteMeter: gqUt2ia7B.accountWasteMeter,
		storageLoadMeter:  gqUt2ia7B.storageLoadMeter,
		storageDupMeter:   gqUt2ia7B.storageDupMeter,
		storageSkipMeter:  gqUt2ia7B.storageSkipMeter,
		storageWasteMeter: gqUt2ia7B.storageWasteMeter,
	}

	if gqUt2ia7B.fetches != nil {
		for lxItfhw, a41Sz4b02r := range gqUt2ia7B.fetches {
			if a41Sz4b02r == nil {
				continue
			}
			stapkw2aW.fetches[lxItfhw] = /*line ZeJB6NiYZYf.go:1*/ gqUt2ia7B.db.CopyTrie(a41Sz4b02r)
		}
		return stapkw2aW
	}

	for gVThh7O, eSIabWY := range gqUt2ia7B.fetchers {
		stapkw2aW.fetches[gVThh7O] = /*line mzMWoCvd.go:1*/ eSIabWY.b5BENy3zhaYk()
	}
	return stapkw2aW
}

func (gqUt2ia7B *triePrefetcher) i3uM6DE(gsZYLFea common.Hash, lxItfhw common.Hash, hv1UkMRaKf common.Address, c3ZjIzfPiDZ [][]byte) {

	if gqUt2ia7B.fetches != nil {
		return
	}

	gVThh7O := /*line x3HH54B6Pte.go:1*/ gqUt2ia7B.e6wsDMQ8(gsZYLFea, lxItfhw)
	eSIabWY := gqUt2ia7B.fetchers[gVThh7O]
	if eSIabWY == nil {
		eSIabWY = /*line UukfWw.go:1*/ tnuCOJi1QxLK(gqUt2ia7B.db, gqUt2ia7B.root, gsZYLFea, lxItfhw, hv1UkMRaKf)
		gqUt2ia7B.fetchers[gVThh7O] = eSIabWY
	}
	/*line jg5l1epr.go:1*/ eSIabWY.gs5s8dNZcp(c3ZjIzfPiDZ)
}

func (gqUt2ia7B *triePrefetcher) fQodR9c7t(gsZYLFea common.Hash, lxItfhw common.Hash) Trie {

	gVThh7O := /*line ysG9QqUhr.go:1*/ gqUt2ia7B.e6wsDMQ8(gsZYLFea, lxItfhw)
	if gqUt2ia7B.fetches != nil {
		fQodR9c7t := gqUt2ia7B.fetches[gVThh7O]
		if fQodR9c7t == nil {
			/*line qP91iy2.go:1*/ gqUt2ia7B.deliveryMissMeter.Mark(1)
			return nil
		}
		return /*line cVJeOKh.go:1*/ gqUt2ia7B.db.CopyTrie(fQodR9c7t)
	}

	eSIabWY := gqUt2ia7B.fetchers[gVThh7O]
	if eSIabWY == nil {
		/*line XlFeXSEZ2cc.go:1*/ gqUt2ia7B.deliveryMissMeter.Mark(1)
		return nil
	}

	/*line P_JPBba7Uvbs.go:1*/
	eSIabWY.tj1faB2CfMY8()

	fQodR9c7t := /*line qM6Pz0cN.go:1*/ eSIabWY.b5BENy3zhaYk()
	if fQodR9c7t == nil {
		/*line BBmcB1Vtl.go:1*/ gqUt2ia7B.deliveryMissMeter.Mark(1)
		return nil
	}
	return fQodR9c7t
}

func (gqUt2ia7B *triePrefetcher) t50cl4u8e(gsZYLFea common.Hash, lxItfhw common.Hash, t50cl4u8e [][]byte) {
	if eSIabWY := gqUt2ia7B.fetchers[ /*line ZF7v_f5_uuk.go:1*/ gqUt2ia7B.e6wsDMQ8(gsZYLFea, lxItfhw)]; eSIabWY != nil {
		eSIabWY.used = t50cl4u8e
	}
}

func (gqUt2ia7B *triePrefetcher) e6wsDMQ8(gsZYLFea common.Hash, lxItfhw common.Hash) string {
	e6wsDMQ8 := /*line SCPc1Tg.go:1*/ make([]byte, common.HashLength*2)
	/*line WOvsdaVJR0x.go:1*/ copy(e6wsDMQ8 /*line hN5_97jm.go:1*/, gsZYLFea.Bytes())
	/*line PpAaeVpsi_.go:1*/ copy(e6wsDMQ8[common.HashLength:] /*line z3BWsf.go:1*/, lxItfhw.Bytes())
	return /*line ZGOkkp.go:1*/ string(e6wsDMQ8)
}

type subfetcher struct {
	db    Database
	state common.Hash
	owner common.Hash
	root  common.Hash
	addr  common.Address
	trie  Trie

	tasks [][]byte
	lock  sync.Mutex

	wake chan struct{}
	stop chan struct{}
	term chan struct{}
	copy chan chan Trie

	seen map[string]struct{}
	dups int
	used [][]byte
}

func tnuCOJi1QxLK(tC4SGV Database, ayL8rlo common.Hash, gsZYLFea common.Hash, lxItfhw common.Hash, hv1UkMRaKf common.Address) *subfetcher {
	eOhbaKqly := &subfetcher{
		db:    tC4SGV,
		state: ayL8rlo,
		owner: gsZYLFea,
		root:  lxItfhw,
		addr:  hv1UkMRaKf,
		wake:/*line OuXL6u.go:1*/ make(chan struct{}, 1),
		stop:/*line Yf2qlySuss.go:1*/ make(chan struct{}),
		term:/*line GJ9HOOG.go:1*/ make(chan struct{}),
		copy:/*line C1BaTh.go:1*/ make(chan chan Trie),
		seen:/*line j7PAg9AjEmDQ.go:1*/ make(map[string]struct{}),
	}
	go /*line LKSkZSR.go:1*/ eOhbaKqly.cnN4QWKBv()
	return eOhbaKqly
}

func (eOhbaKqly *subfetcher) gs5s8dNZcp(c3ZjIzfPiDZ [][]byte) {

	/*line uQB4WR.go:1*/
	eOhbaKqly.lock.Lock()
	eOhbaKqly.tasks = /*line hlA3148zrY3.go:1*/ append(eOhbaKqly.tasks, c3ZjIzfPiDZ...)
	/*line hPkSaRIyhn.go:1*/ eOhbaKqly.lock.Unlock()

	select {
	case eOhbaKqly.wake <- struct{}{}:
	default:
	}
}

func (eOhbaKqly *subfetcher) b5BENy3zhaYk() Trie {
	gu0zGueC4zLM := /*line YlbaQEr5Yzr.go:1*/ make(chan Trie)
	select {
	case eOhbaKqly.copy <- gu0zGueC4zLM:

		return <-gu0zGueC4zLM

	case <-eOhbaKqly.term:

		if eOhbaKqly.trie == nil {
			return nil
		}
		return /*line oMp1rC1IA.go:1*/ eOhbaKqly.db.CopyTrie(eOhbaKqly.trie)
	}
}

func (eOhbaKqly *subfetcher) tj1faB2CfMY8() {
	select {
	case <-eOhbaKqly.stop:
	default:
		/*line JW6H2mbNcLlz.go:1*/ close(eOhbaKqly.stop)
	}
	<-eOhbaKqly.term
}

func (eOhbaKqly *subfetcher) cnN4QWKBv() {

	defer /*line rEDNZKd7iXm.go:1*/ close(eOhbaKqly.term)

	if eOhbaKqly.owner == (common.Hash{}) {
		fQodR9c7t, cYHA75qVq := /*line zfBBUumoO.go:1*/ eOhbaKqly.db.OpenTrie(eOhbaKqly.root)
		if cYHA75qVq != nil {
			/*line vLfI7Gbs.go:1*/ log.Warn(func() /*line wNNIbemU.go:1*/ string {
				data := /*line wNNIbemU.go:1*/ make([]byte, 0, 36)
				i := 1
				decryptKey := 199
				for counter := 0; i != 14; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 6:
						data = /*line wNNIbemU.go:1*/ append(data, "\xa8}s"...,
						)
						i = 11
					case 7:
						data = /*line wNNIbemU.go:1*/ append(data, "\x87\x95"...,
						)
						i = 3
					case 5:
						data = /*line wNNIbemU.go:1*/ append(data, "\xa2]\xab\xa3"...,
						)
						i = 2
					case 8:
						data = /*line wNNIbemU.go:1*/ append(data, "D\xa1\x9b\xa2"...,
						)
						i = 4
					case 12:
						data = /*line wNNIbemU.go:1*/ append(data, "\x9dU"...,
						)
						i = 6
					case 11:
						i = 0
						data = /*line wNNIbemU.go:1*/ append(data, 110)
					case 9:
						i = 12
						data = /*line wNNIbemU.go:1*/ append(data, "\x9f\x99\xa5"...,
						)
					case 2:
						i = 9
						data = /*line wNNIbemU.go:1*/ append(data, 151)
					case 3:
						data = /*line wNNIbemU.go:1*/ append(data, "\x83\x8f\x8b\x97"...,
						)
						i = 8
					case 1:
						data = /*line wNNIbemU.go:1*/ append(data, "\x7f\x9c"...,
						)
						i = 13
					case 4:
						i = 5
						data = /*line wNNIbemU.go:1*/ append(data, "\xa4\xa4"...,
						)
					case 0:
						i = 14
						for y := range data {
							data[y] = data[y] - /*line wNNIbemU.go:1*/ byte(decryptKey^y)
						}
					case 13:
						data = /*line wNNIbemU.go:1*/ append(data, "\x92\x8dO"...,
						)
						i = 10
					case 10:
						data = /*line wNNIbemU.go:1*/ append(data, "\x9e\x9f\x91\x89"...,
						)
						i = 7
					}
				}
				return /*line wNNIbemU.go:1*/ string(data)
			}(), "root", eOhbaKqly.root, "err", cYHA75qVq)
			return
		}
		eOhbaKqly.trie = fQodR9c7t
	} else {

		fQodR9c7t, cYHA75qVq := /*line YKo9OoBW.go:1*/ eOhbaKqly.db.OpenStorageTrie(eOhbaKqly.state, eOhbaKqly.addr, eOhbaKqly.root, nil)
		if cYHA75qVq != nil {
			/*line IbKQWKSL1b.go:1*/ log.Warn(func() /*line F3djOmOBN.go:1*/ string {
				data := /*line F3djOmOBN.go:1*/ make([]byte, 0, 36)
				i := 15
				decryptKey := 48
				for counter := 0; i != 8; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 7:
						data = /*line F3djOmOBN.go:1*/ append(data, "\x1fN"...,
						)
						i = 18
					case 1:
						i = 12
						data = /*line F3djOmOBN.go:1*/ append(data, 72)
					case 17:
						i = 10
						data = /*line F3djOmOBN.go:1*/ append(data, "GOI"...,
						)
					case 2:
						i = 9
						data = /*line F3djOmOBN.go:1*/ append(data, "_SG"...,
						)
					case 15:
						i = 1
						data = /*line F3djOmOBN.go:1*/ append(data, 111)
					case 16:
						i = 8
						for y := range data {
							data[y] = data[y] ^ /*line F3djOmOBN.go:1*/ byte(decryptKey^y)
						}
					case 11:
						data = /*line F3djOmOBN.go:1*/ append(data, "WES"...,
						)
						i = 2
					case 6:
						data = /*line F3djOmOBN.go:1*/ append(data, "J\rC"...,
						)
						i = 13
					case 0:
						i = 7
						data = /*line F3djOmOBN.go:1*/ append(data, 93)
					case 9:
						i = 4
						data = /*line F3djOmOBN.go:1*/ append(data, "\x14M"...,
						)
					case 4:
						data = /*line F3djOmOBN.go:1*/ append(data, 75)
						i = 14
					case 5:
						i = 16
						data = /*line F3djOmOBN.go:1*/ append(data, "Pis|"...,
						)
					case 14:
						data = /*line F3djOmOBN.go:1*/ append(data, 64)
						i = 3
					case 10:
						data = /*line F3djOmOBN.go:1*/ append(data, "IA\x05"...,
						)
						i = 5
					case 12:
						i = 0
						data = /*line F3djOmOBN.go:1*/ append(data, 80)
					case 13:
						data = /*line F3djOmOBN.go:1*/ append(data, 83)
						i = 17
					case 3:
						data = /*line F3djOmOBN.go:1*/ append(data, "DJ"...,
						)
						i = 6
					case 18:
						data = /*line F3djOmOBN.go:1*/ append(data, "OYU"...,
						)
						i = 11
					}
				}
				return /*line F3djOmOBN.go:1*/ string(data)
			}(), "root", eOhbaKqly.root, "err", cYHA75qVq)
			return
		}
		eOhbaKqly.trie = fQodR9c7t
	}

	for {
		select {
		case <-eOhbaKqly.wake:

			/*line yt3gUrgkz5m1.go:1*/
			eOhbaKqly.lock.Lock()
			rN8AfmdbLCa := eOhbaKqly.tasks
			eOhbaKqly.tasks = nil
			/*line uU9cV3QxW.go:1*/ eOhbaKqly.lock.Unlock()

			for cDClIaFDS, w3mNn_vU := range rN8AfmdbLCa {
				select {
				case <-eOhbaKqly.stop:

					/*line LMsFFc.go:1*/
					eOhbaKqly.lock.Lock()
					eOhbaKqly.tasks = /*line uBEeQqDO.go:1*/ append(eOhbaKqly.tasks, rN8AfmdbLCa[cDClIaFDS:]...)
					/*line MWHu8Piz.go:1*/ eOhbaKqly.lock.Unlock()
					return

				case gu0zGueC4zLM := <-eOhbaKqly.copy:

					gu0zGueC4zLM <- /*line NEjpqaam.go:1*/ eOhbaKqly.db.CopyTrie(eOhbaKqly.trie)

				default:

					if _, dNazL3Oz4 := eOhbaKqly.seen[ /*line vRWeUOqFZ.go:1*/ string(w3mNn_vU)]; dNazL3Oz4 {
						eOhbaKqly.dups++
					} else {
						if /*line agnkdOj.go:1*/ len(w3mNn_vU) == common.AddressLength {
							/*line Bk4oIW.go:1*/ eOhbaKqly.trie.GetAccount( /*line h9_6WxSlP8.go:1*/ common.BytesToAddress(w3mNn_vU))
						} else {
							/*line N5pD76V.go:1*/ eOhbaKqly.trie.GetStorage(eOhbaKqly.addr, w3mNn_vU)
						}
						eOhbaKqly.seen[ /*line rhzm2Ev.go:1*/ string(w3mNn_vU)] = struct{}{}
					}
				}
			}

		case gu0zGueC4zLM := <-eOhbaKqly.copy:

			gu0zGueC4zLM <- /*line KMqb71ZRyZ.go:1*/ eOhbaKqly.db.CopyTrie(eOhbaKqly.trie)

		case <-eOhbaKqly.stop:

			return
		}
	}
}

var _ sync.Cond
var _ = common.AbsolutePath
var _ = log.Crit
var _ = metrics.AccountingRegistry
