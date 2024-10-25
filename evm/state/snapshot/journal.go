//line :1
package snapshot

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/triedb"
)

const journalVersion uint64 = 0

type journalGenerator struct {
	Wiping bool

	Done     bool
	Marker   []byte
	Accounts uint64
	Slots    uint64
	Storage  uint64
}

type journalDestruct struct {
	Hash common.Hash
}

type journalAccount struct {
	Hash common.Hash
	Blob []byte
}

type journalStorage struct {
	Hash common.Hash
	Keys []common.Hash
	Vals [][]byte
}

func DHCaqaAWGa7(gJTjXnF []byte) string {
	if /*line uDvRXOf.go:1*/ len(gJTjXnF) == 0 {
		return ""
	}
	var iatAGX journalGenerator
	if chyZ8Q := /*line Ij4MNlHb.go:1*/ rlp.DecodeBytes(gJTjXnF, &iatAGX); chyZ8Q != nil {
		/*line bg2oBxJAPr.go:1*/ log.Warn(func() /*line XqJfYo4S.go:1*/ string {
			seed := /*line XqJfYo4S.go:1*/ byte(122)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line XqJfYo4S.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line XqJfYo4S.go:1*/ fnc(224)(187)(126)(255)(247)(237)(150)(128)(251)(167)(146)(37)(72)(156)(45)(91)(113)(53)(101)(189)(137)(21)(31)(69)(143)(202)(219)(180)(113)(217)(191)(109)(237)(213)(173)
			return /*line XqJfYo4S.go:1*/ string(data)
		}(), "err", chyZ8Q)
		return ""
	}

	var dQPD9HQ string
	switch al5iRGOo := iatAGX.Marker; /*line Ixv1H2.go:1*/ len(al5iRGOo) {
	case common.HashLength:
		dQPD9HQ = /*line xUngMe643ia9.go:1*/ fmt.Sprintf("at %#x", al5iRGOo)
	case 2 * common.HashLength:
		dQPD9HQ = /*line DC5A4lvd.go:1*/ fmt.Sprintf(func() /*line _JoPqOXn4.go:1*/ string {
			seed := /*line _JoPqOXn4.go:1*/ byte(59)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line _JoPqOXn4.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line _JoPqOXn4.go:1*/ fnc(164)(77)(76)(157)(56)(197)(50)(165)(93)(102)(209)(160)(149)
			return /*line _JoPqOXn4.go:1*/ string(data)
		}(), al5iRGOo[:common.HashLength], al5iRGOo[common.HashLength:])
	default:
		dQPD9HQ = /*line ALi7gw.go:1*/ fmt.Sprintf("%#x", al5iRGOo)
	}
	return /*line Jt9MjRL.go:1*/ fmt.Sprintf(func() /*line Bv284p.go:1*/ string {
		fullData := [] /*line Bv284p.go:1*/ byte("\xee\xd4\xf8W!\x83\x8a\xae\xab\xdd\xf3_\x9cG\x80\xb8\xf1\xf3@h氄\xd7w\xceM\x17\x86\x17\xe01\x8a\x8f\xffm\x1e\xee\xdb\x1c\xefs\bI\xab\x80\xc1h\xb7\xbb\xfb\x06L\x95\xe3\xf4\xf9\x1cfywwM\xb4O\x7f_\x9fF&Iv\x15J\xefS\xae{7\xc0\x06\x9e\xe9\xd7\x02\xbc\x8c0緌o\xf7\f\x8cu\x87,\xecHZ\x88\xf0\x02J\x99z\xd7j\x9b\xdf\x04,\x96\xd9\xfe")
		idxKey := [] /*line Bv284p.go:1*/ byte("\xf3i\xdc\xe3\a\xbb \x10=\xbe\x9af\x8c\x9a\xa3\xd5")
		data := /*line Bv284p.go:1*/ make([]byte, 0, 59)
		data = /*line Bv284p.go:1*/ append(data, fullData[234^ /*line Bv284p.go:1*/ int(idxKey[10])]^fullData[181^ /*line Bv284p.go:1*/ int(idxKey[10])], fullData[133^ /*line Bv284p.go:1*/ int(idxKey[2])]+fullData[211^ /*line Bv284p.go:1*/ int(idxKey[2])], fullData[103^ /*line Bv284p.go:1*/ int(idxKey[6])]-fullData[10^ /*line Bv284p.go:1*/ int(idxKey[6])], fullData[70^ /*line Bv284p.go:1*/ int(idxKey[11])]^fullData[44^ /*line Bv284p.go:1*/ int(idxKey[11])], fullData[89^ /*line Bv284p.go:1*/ int(idxKey[7])]+fullData[118^ /*line Bv284p.go:1*/ int(idxKey[7])], fullData[155^ /*line Bv284p.go:1*/ int(idxKey[14])]^fullData[209^ /*line Bv284p.go:1*/ int(idxKey[14])], fullData[221^ /*line Bv284p.go:1*/ int(idxKey[12])]-fullData[183^ /*line Bv284p.go:1*/ int(idxKey[12])], fullData[13^ /*line Bv284p.go:1*/ int(idxKey[1])]+fullData[80^ /*line Bv284p.go:1*/ int(idxKey[1])], fullData[186^ /*line Bv284p.go:1*/ int(idxKey[12])]+fullData[167^ /*line Bv284p.go:1*/ int(idxKey[12])], fullData[144^ /*line Bv284p.go:1*/ int(idxKey[15])]^fullData[230^ /*line Bv284p.go:1*/ int(idxKey[15])], fullData[13^ /*line Bv284p.go:1*/ int(idxKey[11])]-fullData[23^ /*line Bv284p.go:1*/ int(idxKey[11])], fullData[228^ /*line Bv284p.go:1*/ int(idxKey[9])]^fullData[150^ /*line Bv284p.go:1*/ int(idxKey[9])], fullData[67^ /*line Bv284p.go:1*/ int(idxKey[7])]+fullData[78^ /*line Bv284p.go:1*/ int(idxKey[7])], fullData[71^ /*line Bv284p.go:1*/ int(idxKey[1])]^fullData[37^ /*line Bv284p.go:1*/ int(idxKey[1])], fullData[237^ /*line Bv284p.go:1*/ int(idxKey[5])]-fullData[160^ /*line Bv284p.go:1*/ int(idxKey[5])], fullData[252^ /*line Bv284p.go:1*/ int(idxKey[9])]-fullData[174^ /*line Bv284p.go:1*/ int(idxKey[9])], fullData[37^ /*line Bv284p.go:1*/ int(idxKey[4])]+fullData[88^ /*line Bv284p.go:1*/ int(idxKey[4])], fullData[166^ /*line Bv284p.go:1*/ int(idxKey[0])]+fullData[195^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[37^ /*line Bv284p.go:1*/ int(idxKey[8])]^fullData[39^ /*line Bv284p.go:1*/ int(idxKey[8])], fullData[191^ /*line Bv284p.go:1*/ int(idxKey[13])]-fullData[131^ /*line Bv284p.go:1*/ int(idxKey[13])], fullData[223^ /*line Bv284p.go:1*/ int(idxKey[0])]-fullData[239^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[249^ /*line Bv284p.go:1*/ int(idxKey[13])]^fullData[251^ /*line Bv284p.go:1*/ int(idxKey[13])], fullData[228^ /*line Bv284p.go:1*/ int(idxKey[0])]^fullData[193^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[158^ /*line Bv284p.go:1*/ int(idxKey[0])]-fullData[190^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[48^ /*line Bv284p.go:1*/ int(idxKey[8])]+fullData[96^ /*line Bv284p.go:1*/ int(idxKey[8])], fullData[222^ /*line Bv284p.go:1*/ int(idxKey[0])]+fullData[145^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[19^ /*line Bv284p.go:1*/ int(idxKey[4])]-fullData[59^ /*line Bv284p.go:1*/ int(idxKey[4])], fullData[161^ /*line Bv284p.go:1*/ int(idxKey[14])]-fullData[181^ /*line Bv284p.go:1*/ int(idxKey[14])], fullData[173^ /*line Bv284p.go:1*/ int(idxKey[14])]+fullData[178^ /*line Bv284p.go:1*/ int(idxKey[14])], fullData[252^ /*line Bv284p.go:1*/ int(idxKey[3])]-fullData[191^ /*line Bv284p.go:1*/ int(idxKey[3])], fullData[241^ /*line Bv284p.go:1*/ int(idxKey[3])]+fullData[253^ /*line Bv284p.go:1*/ int(idxKey[3])], fullData[222^ /*line Bv284p.go:1*/ int(idxKey[13])]+fullData[244^ /*line Bv284p.go:1*/ int(idxKey[13])], fullData[139^ /*line Bv284p.go:1*/ int(idxKey[12])]-fullData[228^ /*line Bv284p.go:1*/ int(idxKey[12])], fullData[230^ /*line Bv284p.go:1*/ int(idxKey[9])]-fullData[143^ /*line Bv284p.go:1*/ int(idxKey[9])], fullData[121^ /*line Bv284p.go:1*/ int(idxKey[7])]+fullData[112^ /*line Bv284p.go:1*/ int(idxKey[7])], fullData[213^ /*line Bv284p.go:1*/ int(idxKey[13])]-fullData[185^ /*line Bv284p.go:1*/ int(idxKey[13])], fullData[199^ /*line Bv284p.go:1*/ int(idxKey[12])]+fullData[136^ /*line Bv284p.go:1*/ int(idxKey[12])], fullData[31^ /*line Bv284p.go:1*/ int(idxKey[6])]^fullData[6^ /*line Bv284p.go:1*/ int(idxKey[6])], fullData[97^ /*line Bv284p.go:1*/ int(idxKey[6])]+fullData[42^ /*line Bv284p.go:1*/ int(idxKey[6])], fullData[59^ /*line Bv284p.go:1*/ int(idxKey[1])]^fullData[12^ /*line Bv284p.go:1*/ int(idxKey[1])], fullData[224^ /*line Bv284p.go:1*/ int(idxKey[3])]^fullData[180^ /*line Bv284p.go:1*/ int(idxKey[3])], fullData[201^ /*line Bv284p.go:1*/ int(idxKey[14])]-fullData[235^ /*line Bv284p.go:1*/ int(idxKey[14])], fullData[182^ /*line Bv284p.go:1*/ int(idxKey[9])]+fullData[159^ /*line Bv284p.go:1*/ int(idxKey[9])], fullData[242^ /*line Bv284p.go:1*/ int(idxKey[0])]+fullData[199^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[102^ /*line Bv284p.go:1*/ int(idxKey[11])]+fullData[40^ /*line Bv284p.go:1*/ int(idxKey[11])], fullData[200^ /*line Bv284p.go:1*/ int(idxKey[15])]^fullData[252^ /*line Bv284p.go:1*/ int(idxKey[15])], fullData[208^ /*line Bv284p.go:1*/ int(idxKey[2])]^fullData[201^ /*line Bv284p.go:1*/ int(idxKey[2])], fullData[156^ /*line Bv284p.go:1*/ int(idxKey[13])]-fullData[246^ /*line Bv284p.go:1*/ int(idxKey[13])], fullData[179^ /*line Bv284p.go:1*/ int(idxKey[0])]-fullData[148^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[99^ /*line Bv284p.go:1*/ int(idxKey[7])]^fullData[83^ /*line Bv284p.go:1*/ int(idxKey[7])], fullData[7^ /*line Bv284p.go:1*/ int(idxKey[8])]-fullData[10^ /*line Bv284p.go:1*/ int(idxKey[8])], fullData[7^ /*line Bv284p.go:1*/ int(idxKey[6])]^fullData[29^ /*line Bv284p.go:1*/ int(idxKey[6])], fullData[202^ /*line Bv284p.go:1*/ int(idxKey[10])]+fullData[145^ /*line Bv284p.go:1*/ int(idxKey[10])], fullData[21^ /*line Bv284p.go:1*/ int(idxKey[6])]+fullData[41^ /*line Bv284p.go:1*/ int(idxKey[6])], fullData[187^ /*line Bv284p.go:1*/ int(idxKey[9])]-fullData[248^ /*line Bv284p.go:1*/ int(idxKey[9])], fullData[190^ /*line Bv284p.go:1*/ int(idxKey[10])]+fullData[206^ /*line Bv284p.go:1*/ int(idxKey[10])], fullData[205^ /*line Bv284p.go:1*/ int(idxKey[0])]^fullData[224^ /*line Bv284p.go:1*/ int(idxKey[0])], fullData[6^ /*line Bv284p.go:1*/ int(idxKey[1])]+fullData[50^ /*line Bv284p.go:1*/ int(idxKey[1])])
		return /*line Bv284p.go:1*/ string(data)
	}(),
		iatAGX.Done, iatAGX.Accounts, iatAGX.Slots, iatAGX.Storage, dQPD9HQ)
}

func aQkgwpB129k(lDBEbd ethdb.KeyValueStore, v4faEBKJgI2 *diskLayer) (snapshot, journalGenerator, error) {

	gJTjXnF := /*line MkuysxR.go:1*/ rawdb.ReadSnapshotGenerator(lDBEbd)
	if /*line t73SUPCTBk.go:1*/ len(gJTjXnF) == 0 {
		return nil, journalGenerator{} /*line _HkoTVjIKFb.go:1*/, errors.New(func() /*line EGDfSmG_xy.go:1*/ string {
			key := [] /*line EGDfSmG_xy.go:1*/ byte("z\xa8\xba\xa8s\xc8\xec^V\xdb\xc3)!\x04\v\xe0L\xe8\xd9&\xff\xddh\xf67\n")
			data := [] /*line EGDfSmG_xy.go:1*/ byte("\xe7\x11-\x1b\xdc6S~\xc9I$\x99\x94lzTlO>\x94dO\xc9j\xa6|")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line EGDfSmG_xy.go:1*/ string(data)
		}())
	}
	var iatAGX journalGenerator
	if chyZ8Q := /*line KETtYRSAaZi.go:1*/ rlp.DecodeBytes(gJTjXnF, &iatAGX); chyZ8Q != nil {
		return nil, journalGenerator{} /*line CO7OVhULDwT.go:1*/, fmt.Errorf(func() /*line DOwn8XDNE.go:1*/ string {
			data := [] /*line DOwn8XDNE.go:1*/ byte("\xdc\xe4\xd1le\xb1\x14\xf2X\xcc;ec\xd4d\x01\xe0\xe7\x82a\xf0\xfd\u0099tU\xaeC\xaeer\x8f\xd2\xeb\x16\xb8 !c")
			positions := [...]byte{35, 37, 18, 31, 6, 25, 10, 9, 13, 2, 34, 15, 25, 1, 8, 16, 22, 8, 21, 23, 20, 20, 7, 27, 7, 13, 32, 23, 38, 15, 17, 26, 8, 2, 1, 34, 28, 15, 34, 35, 35, 25, 2, 28, 5, 10, 28, 33, 20, 13, 25, 37, 26, 27, 0, 34}
			for i := 0; i < 56; i += 2 {
				localKey := /*line DOwn8XDNE.go:1*/ byte(i) + /*line DOwn8XDNE.go:1*/ byte(positions[i]^positions[i+1]) + 18
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line DOwn8XDNE.go:1*/ string(data)
		}(), chyZ8Q)
	}

	var y84IW11J snapshot = v4faEBKJgI2
	chyZ8Q := /*line WDXb1fa_9b5.go:1*/ twkYxFLo(lDBEbd, func(f3dWTs7 common.Hash, z1BBTN2UX common.Hash, vgaPgkN map[common.Hash]struct{}, oVnS8XskWEE map[common.Hash][]byte, bJNw07c3Namn map[common.Hash]map[common.Hash][]byte) error {
		y84IW11J = /*line jMjcEX.go:1*/ aa7YtGB_s(y84IW11J, z1BBTN2UX, vgaPgkN, oVnS8XskWEE, bJNw07c3Namn)
		return nil
	})
	if chyZ8Q != nil {
		return v4faEBKJgI2, iatAGX, nil
	}
	return y84IW11J, iatAGX, nil
}

func nsDET4Qxptc(iYSzoWnzsqd ethdb.KeyValueStore, rhls5wQnNg *triedb.Database, z1BBTN2UX common.Hash, duGwhWx7d5J int, c9futN1PG6d bool, dVteLD bool) (snapshot, bool, error) {

	if /*line jSRjwXJy.go:1*/ rawdb.ReadSnapshotDisabled(iYSzoWnzsqd) {
		return nil, true, nil
	}

	koxaxufVjnP := /*line we_IbLlW.go:1*/ rawdb.ReadSnapshotRoot(iYSzoWnzsqd)
	if koxaxufVjnP == (common.Hash{}) {
		return nil, false /*line FEFgkKN_bB.go:1*/, errors.New(func() /*line OuQvtZtrO.go:1*/ string {
			data := [] /*line OuQvtZtrO.go:1*/ byte(":i\xdf\xc099\x17 A,2\xcboD\x00\x19\x14\xcae@ 1\xfbapI6o\xc6")
			positions := [...]byte{22, 4, 3, 11, 8, 3, 9, 26, 13, 14, 21, 19, 0, 5, 13, 25, 4, 8, 11, 13, 11, 17, 2, 10, 28, 13, 17, 15, 11, 6, 13, 16}
			for i := 0; i < 32; i += 2 {
				localKey := /*line OuQvtZtrO.go:1*/ byte(i) + /*line OuQvtZtrO.go:1*/ byte(positions[i]^positions[i+1]) + 35
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line OuQvtZtrO.go:1*/ string(data)
		}())
	}
	v4faEBKJgI2 := &diskLayer{
		diskdb: iYSzoWnzsqd,
		triedb: rhls5wQnNg,
		cache:/*line J2zeXabCw.go:1*/ fastcache.New(duGwhWx7d5J * 1024 * 1024),
		root: koxaxufVjnP,
	}
	yk1GostL, iatAGX, chyZ8Q := /*line MehKNHT6.go:1*/ aQkgwpB129k(iYSzoWnzsqd, v4faEBKJgI2)
	if chyZ8Q != nil {
		/*line T1K4EgIXxamN.go:1*/ log.Warn(func() /*line ExwjNVYTEO97.go:1*/ string {
			data := [] /*line ExwjNVYTEO97.go:1*/ byte("cB-Z[d\x1fY\x16pl5aB %ouUlaK")
			positions := [...]byte{2, 21, 3, 21, 11, 1, 2, 18, 7, 11, 13, 3, 19, 8, 19, 6, 15, 15, 4, 3, 0, 15, 9, 19, 19, 3, 1, 1}
			for i := 0; i < 28; i += 2 {
				localKey := /*line ExwjNVYTEO97.go:1*/ byte(i) + /*line ExwjNVYTEO97.go:1*/ byte(positions[i]^positions[i+1]) + 214
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line ExwjNVYTEO97.go:1*/ string(data)
		}(), "error", chyZ8Q)
		return nil, false, chyZ8Q
	}

	if j9xp0xEW9BzB := /*line UkGqkrnbv.go:1*/ yk1GostL.Root(); j9xp0xEW9BzB != z1BBTN2UX {

		if !c9futN1PG6d {
			return nil, false /*line dH2Jr5.go:1*/, fmt.Errorf(func() /*line PLg46Pm.go:1*/ string {
				data := [] /*line PLg46Pm.go:1*/ byte("\xfce\xf5d\x82\xca\x06\xbe\xb3,'YSm\xc1\xc7Oh\xb6\xf6napdVs\x9bV ӏ\xf4\x04h\xd2xx\x92\x90wEn\xa7\x7fLFx")
				positions := [...]byte{23, 38, 45, 30, 31, 26, 38, 29, 26, 18, 30, 8, 29, 37, 7, 6, 15, 45, 19, 27, 37, 5, 44, 31, 43, 15, 11, 44, 6, 2, 12, 0, 29, 35, 25, 16, 31, 24, 42, 44, 19, 16, 2, 18, 31, 25, 6, 33, 15, 24, 7, 4, 44, 11, 35, 14, 9, 43, 0, 15, 16, 40, 35, 34, 18, 32}
				for i := 0; i < 66; i += 2 {
					localKey := /*line PLg46Pm.go:1*/ byte(i) + /*line PLg46Pm.go:1*/ byte(positions[i]^positions[i+1]) + 178
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line PLg46Pm.go:1*/ string(data)
			}(), j9xp0xEW9BzB, z1BBTN2UX)
		}

		/*line AWCL2X.go:1*/
		log.Warn(func() /*line BG4nqm.go:1*/ string {
			data := [] /*line BG4nqm.go:1*/ byte("c\xb5\x82\xb5s\x8coi i\x85\x88I\x9c\x8b\xbaconti\x9dU܀\x81x\x02\u05ca\xbb\x88s\x82\x85i6")
			positions := [...]byte{30, 0, 12, 3, 30, 2, 28, 14, 22, 27, 1, 33, 14, 34, 26, 36, 13, 21, 19, 1, 5, 1, 7, 3, 27, 3, 23, 15, 29, 11, 7, 25, 14, 7, 19, 19, 11, 34, 25, 10, 24, 25, 31, 22, 32, 5, 33, 23}
			for i := 0; i < 48; i += 2 {
				localKey := /*line BG4nqm.go:1*/ byte(i) + /*line BG4nqm.go:1*/ byte(positions[i]^positions[i+1]) + 202
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line BG4nqm.go:1*/ string(data)
		}(), func() /*line IHgvdde9JCJ.go:1*/ string {
			seed := /*line IHgvdde9JCJ.go:1*/ byte(183)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line IHgvdde9JCJ.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line IHgvdde9JCJ.go:1*/ fnc(42)(79)(145)(49)(100)(197)(138)(25)
			return /*line IHgvdde9JCJ.go:1*/ string(data)
		}(), j9xp0xEW9BzB, func() /*line BCiaCl.go:1*/ string {
			data := [] /*line BCiaCl.go:1*/ byte("\x17\xc6\xc9\x1dn\f\vo\xc8")
			positions := [...]byte{8, 8, 8, 3, 5, 2, 1, 0, 6, 0, 5, 5}
			for i := 0; i < 12; i += 2 {
				localKey := /*line BCiaCl.go:1*/ byte(i) + /*line BCiaCl.go:1*/ byte(positions[i]^positions[i+1]) + 74
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line BCiaCl.go:1*/ string(data)
		}(), z1BBTN2UX)
	}

	if !iatAGX.Done {
		v4faEBKJgI2.genMarker = iatAGX.Marker
		if v4faEBKJgI2.genMarker == nil {
			v4faEBKJgI2.genMarker = []byte{}
		}
	}

	if !iatAGX.Done && !dVteLD {
		v4faEBKJgI2.genPending = /*line rO_T2glBTKz.go:1*/ make(chan struct{})
		v4faEBKJgI2.genAbort = /*line E3YafmG.go:1*/ make(chan chan *generatorStats)

		var ijcATHiDv uint64
		if /*line BVj_kTQt0.go:1*/ len(iatAGX.Marker) >= 8 {
			ijcATHiDv = /*line iVazTDlP7HJR.go:1*/ binary.BigEndian.Uint64(iatAGX.Marker)
		}
		go /*line LAteZMmqODrm.go:1*/ v4faEBKJgI2.tliqpCic(&generatorStats{
			origin: ijcATHiDv,
			start:/*line IVprOh_A.go:1*/ time.Now(),
			accounts: iatAGX.Accounts,
			slots:    iatAGX.Slots,
			storage:/*line QQ735Bv9_YK.go:1*/ common.StorageSize(iatAGX.Storage),
		})
	}
	return yk1GostL, false, nil
}

func (hbsnnIi *diskLayer) Journal(z1J97J9 *bytes.Buffer) (common.Hash, error) {

	var drjkE0 *generatorStats
	if hbsnnIi.genAbort != nil {
		stpiKC6okfS8 := /*line jws2EaHH_0Am.go:1*/ make(chan *generatorStats)
		hbsnnIi.genAbort <- stpiKC6okfS8

		if drjkE0 = <-stpiKC6okfS8; drjkE0 != nil {
			/*line czloGBe.go:1*/ drjkE0.Log(func() /*line z5Xs99rp.go:1*/ string {
				data := [] /*line z5Xs99rp.go:1*/ byte("NKurnKAl\x1dn\xefXQn\xf5pq>\xd3&wC^\xf4\x0fnn_\xefiot")
				positions := [...]byte{14, 21, 16, 20, 28, 14, 23, 10, 21, 11, 18, 10, 27, 26, 19, 8, 20, 22, 10, 10, 29, 28, 5, 24, 14, 23, 0, 10, 12, 1, 8, 19, 18, 0, 17, 6, 5, 18}
				for i := 0; i < 38; i += 2 {
					localKey := /*line z5Xs99rp.go:1*/ byte(i) + /*line z5Xs99rp.go:1*/ byte(positions[i]^positions[i+1]) + 245
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line z5Xs99rp.go:1*/ string(data)
			}(), hbsnnIi.root, hbsnnIi.genMarker)
		}
	}

	/*line ItDCwFlfrUp.go:1*/
	hbsnnIi.lock.RLock()
	defer /*line uSsT2zEz3cB.go:1*/ hbsnnIi.lock.RUnlock()

	if hbsnnIi.stale {
		return common.Hash{}, AiK5hkDaW
	}

	/*line F8GuJ78PU.go:1*/
	bQ5V3QyAnw(hbsnnIi.diskdb, hbsnnIi.genMarker, drjkE0)

	/*line cQC8a2SgC.go:1*/
	log.Debug(func() /*line _VICfTBOxEw.go:1*/ string {
		data := [] /*line _VICfTBOxEw.go:1*/ byte("*\xcbu\xc0\x8a\x03\a\xd2ed d\x0e\xb2\xa3 \xba\x8c\xc1\xcdr")
		positions := [...]byte{17, 5, 0, 13, 17, 4, 6, 17, 5, 4, 12, 16, 19, 3, 5, 13, 7, 6, 1, 18, 16, 4, 14, 17}
		for i := 0; i < 24; i += 2 {
			localKey := /*line _VICfTBOxEw.go:1*/ byte(i) + /*line _VICfTBOxEw.go:1*/ byte(positions[i]^positions[i+1]) + 137
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line _VICfTBOxEw.go:1*/ string(data)
	}(), "root", hbsnnIi.root)
	return hbsnnIi.root, nil
}

func (hbsnnIi *diffLayer) Journal(z1J97J9 *bytes.Buffer) (common.Hash, error) {

	v4faEBKJgI2, chyZ8Q := /*line fafoZyY.go:1*/ hbsnnIi.parent.Journal(z1J97J9)
	if chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}

	/*line X4tIp3K8pk.go:1*/
	hbsnnIi.lock.RLock()
	defer /*line O1WsWGS.go:1*/ hbsnnIi.lock.RUnlock()

	if /*line c4LaB7_9.go:1*/ hbsnnIi.Stale() {
		return common.Hash{}, AiK5hkDaW
	}

	if chyZ8Q := /*line aM8rlglu.go:1*/ rlp.Encode(z1J97J9, hbsnnIi.root); chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}
	zoMmGkL := /*line urbhWWLsiy.go:1*/ make([]journalDestruct, 0 /*line vQJcR4YS.go:1*/, len(hbsnnIi.destructSet))
	for xhOzkRrKDZ := range hbsnnIi.destructSet {
		zoMmGkL = /*line fxaPVXvyDVmX.go:1*/ append(zoMmGkL, journalDestruct{Hash: xhOzkRrKDZ})
	}
	if chyZ8Q := /*line PuZ0DkfiwwAI.go:1*/ rlp.Encode(z1J97J9, zoMmGkL); chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}
	bQIIyfhppAL1 := /*line CnF4febQPX4f.go:1*/ make([]journalAccount, 0 /*line AlSi8spRumY.go:1*/, len(hbsnnIi.accountData))
	for xhOzkRrKDZ, b1jafJL := range hbsnnIi.accountData {
		bQIIyfhppAL1 = /*line mHSsGBJ68Oq_.go:1*/ append(bQIIyfhppAL1, journalAccount{Hash: xhOzkRrKDZ, Blob: b1jafJL})
	}
	if chyZ8Q := /*line OHBSkha7_Wq.go:1*/ rlp.Encode(z1J97J9, bQIIyfhppAL1); chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}
	agSpCMzc := /*line eNlV7VCP.go:1*/ make([]journalStorage, 0 /*line HFcyhRkUaO2N.go:1*/, len(hbsnnIi.storageData))
	for xhOzkRrKDZ, etj53bd3zM := range hbsnnIi.storageData {
		glD6VIl86 := /*line p5gdIh51.go:1*/ make([]common.Hash, 0 /*line tAArtDnq.go:1*/, len(etj53bd3zM))
		xELkOUKoeEz := /*line UO3uQ6HhS4qk.go:1*/ make([][]byte, 0 /*line Y0KYbfc.go:1*/, len(etj53bd3zM))
		for nVUwQz8Zi, jIbih8achE := range etj53bd3zM {
			glD6VIl86 = /*line r30FyYIs.go:1*/ append(glD6VIl86, nVUwQz8Zi)
			xELkOUKoeEz = /*line IcO6kMZAyXs.go:1*/ append(xELkOUKoeEz, jIbih8achE)
		}
		agSpCMzc = /*line pKnWemU.go:1*/ append(agSpCMzc, journalStorage{Hash: xhOzkRrKDZ, Keys: glD6VIl86, Vals: xELkOUKoeEz})
	}
	if chyZ8Q := /*line pqhciR.go:1*/ rlp.Encode(z1J97J9, agSpCMzc); chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}
	/*line d62WiXQ.go:1*/ log.Debug(func() /*line cYBtCa.go:1*/ string {
		data := [] /*line cYBtCa.go:1*/ byte("\x0e\\?KCB.le] b5ff eeGer")
		positions := [...]byte{6, 0, 3, 3, 2, 12, 2, 9, 6, 0, 5, 4, 4, 16, 0, 1, 17, 11, 1, 16, 0, 18, 12, 12}
		for i := 0; i < 24; i += 2 {
			localKey := /*line cYBtCa.go:1*/ byte(i) + /*line cYBtCa.go:1*/ byte(positions[i]^positions[i+1]) + 215
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line cYBtCa.go:1*/ string(data)
	}(), "root", hbsnnIi.root, "parent" /*line wFpusiKtP2co.go:1*/, hbsnnIi.parent.Root())
	return v4faEBKJgI2, nil
}

type gTVIrBGl = func(f3dWTs7 common.Hash, z1BBTN2UX common.Hash, zoMmGkL map[common.Hash]struct{}, bQIIyfhppAL1 map[common.Hash][]byte, agSpCMzc map[common.Hash]map[common.Hash][]byte) error

func twkYxFLo(lDBEbd ethdb.KeyValueReader, kGNYh2Kjh gTVIrBGl) error {
	g0iutUWAEY := /*line DGVkVJUi7.go:1*/ rawdb.ReadSnapshotJournal(lDBEbd)
	if /*line AikS0u.go:1*/ len(g0iutUWAEY) == 0 {
		/*line FNOVTFaKB.go:1*/ log.Warn(func() /*line KuK2bSL.go:1*/ string {
			fullData := [] /*line KuK2bSL.go:1*/ byte("\x17U$\xb2w\x80RS\x92\\\x81\xbeQ\xbex\"\xf3\xed\xc2\xe4m\xdc\xeaj\x1b\\\x02\x10q\x16\x8a\x8c\xa6=Ĭ#V\x13<\x05\x99\x0e\xa1~O")
			idxKey := [] /*line KuK2bSL.go:1*/ byte("#\x19of\x10\xed\fn\x89\xa9\x97H\xe4%")
			data := /*line KuK2bSL.go:1*/ make([]byte, 0, 24)
			data = /*line KuK2bSL.go:1*/ append(data, fullData[12^ /*line KuK2bSL.go:1*/ int(idxKey[4])]^fullData[49^ /*line KuK2bSL.go:1*/ int(idxKey[4])], fullData[16^ /*line KuK2bSL.go:1*/ int(idxKey[1])]+fullData[63^ /*line KuK2bSL.go:1*/ int(idxKey[1])], fullData[19^ /*line KuK2bSL.go:1*/ int(idxKey[6])]^fullData[29^ /*line KuK2bSL.go:1*/ int(idxKey[6])], fullData[57^ /*line KuK2bSL.go:1*/ int(idxKey[1])]+fullData[18^ /*line KuK2bSL.go:1*/ int(idxKey[1])], fullData[34^ /*line KuK2bSL.go:1*/ int(idxKey[0])]+fullData[56^ /*line KuK2bSL.go:1*/ int(idxKey[0])], fullData[106^ /*line KuK2bSL.go:1*/ int(idxKey[2])]^fullData[124^ /*line KuK2bSL.go:1*/ int(idxKey[2])], fullData[14^ /*line KuK2bSL.go:1*/ int(idxKey[13])]-fullData[47^ /*line KuK2bSL.go:1*/ int(idxKey[13])], fullData[8^ /*line KuK2bSL.go:1*/ int(idxKey[13])]^fullData[2^ /*line KuK2bSL.go:1*/ int(idxKey[13])], fullData[187^ /*line KuK2bSL.go:1*/ int(idxKey[9])]+fullData[138^ /*line KuK2bSL.go:1*/ int(idxKey[9])], fullData[227^ /*line KuK2bSL.go:1*/ int(idxKey[5])]-fullData[237^ /*line KuK2bSL.go:1*/ int(idxKey[5])], fullData[138^ /*line KuK2bSL.go:1*/ int(idxKey[8])]+fullData[132^ /*line KuK2bSL.go:1*/ int(idxKey[8])], fullData[133^ /*line KuK2bSL.go:1*/ int(idxKey[8])]^fullData[134^ /*line KuK2bSL.go:1*/ int(idxKey[8])], fullData[200^ /*line KuK2bSL.go:1*/ int(idxKey[12])]-fullData[249^ /*line KuK2bSL.go:1*/ int(idxKey[12])], fullData[13^ /*line KuK2bSL.go:1*/ int(idxKey[1])]^fullData[3^ /*line KuK2bSL.go:1*/ int(idxKey[1])], fullData[120^ /*line KuK2bSL.go:1*/ int(idxKey[7])]+fullData[112^ /*line KuK2bSL.go:1*/ int(idxKey[7])], fullData[60^ /*line KuK2bSL.go:1*/ int(idxKey[13])]+fullData[7^ /*line KuK2bSL.go:1*/ int(idxKey[13])], fullData[135^ /*line KuK2bSL.go:1*/ int(idxKey[10])]^fullData[190^ /*line KuK2bSL.go:1*/ int(idxKey[10])], fullData[158^ /*line KuK2bSL.go:1*/ int(idxKey[8])]+fullData[161^ /*line KuK2bSL.go:1*/ int(idxKey[8])], fullData[143^ /*line KuK2bSL.go:1*/ int(idxKey[8])]+fullData[173^ /*line KuK2bSL.go:1*/ int(idxKey[8])], fullData[109^ /*line KuK2bSL.go:1*/ int(idxKey[11])]^fullData[74^ /*line KuK2bSL.go:1*/ int(idxKey[11])], fullData[24^ /*line KuK2bSL.go:1*/ int(idxKey[4])]+fullData[5^ /*line KuK2bSL.go:1*/ int(idxKey[4])], fullData[199^ /*line KuK2bSL.go:1*/ int(idxKey[5])]+fullData[234^ /*line KuK2bSL.go:1*/ int(idxKey[5])], fullData[33^ /*line KuK2bSL.go:1*/ int(idxKey[13])]^fullData[61^ /*line KuK2bSL.go:1*/ int(idxKey[13])])
			return /*line KuK2bSL.go:1*/ string(data)
		}(), "diffs", "missing")
		return nil
	}
	vaD4aIy := /*line eGom7kwVj.go:1*/ rlp.NewStream( /*line F8O7vD7z8F.go:1*/ bytes.NewReader(g0iutUWAEY), 0)

	ihZoVs, chyZ8Q := /*line Fyq6PFLSrXmr.go:1*/ vaD4aIy.Uint64()
	if chyZ8Q != nil {
		/*line Jcor5LA.go:1*/ log.Warn(func() /*line x7tVtyaROoUB.go:1*/ string {
			fullData := [] /*line x7tVtyaROoUB.go:1*/ byte("\xab\xe9\xfa\xc6\x15\x80n\xa15\x88\x02C3\x93\x80D\xd2\xd7\xe1\x1ej\xee\x047\xbe,z\x8b?\x1c\xa01\x94\xe7<\xab)\xbch\xf4\xce.R\xfa\x13\xe0\xb3\xc1:?7JW\x06\x8b\x1a\x02J\xf8\xcfL\xa46NE\xb2\\r\xe5\xc7^\x1d\x1f\x1a")
			idxKey := [] /*line x7tVtyaROoUB.go:1*/ byte("\x9e\x8f\xech\x16\xc2\xc1S\x10%")
			data := /*line x7tVtyaROoUB.go:1*/ make([]byte, 0, 38)
			data = /*line x7tVtyaROoUB.go:1*/ append(data, fullData[128^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[1])]^fullData[183^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[1])], fullData[8^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])]-fullData[39^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])], fullData[196^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]+fullData[192^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[149^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])]+fullData[186^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])], fullData[165^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])]-fullData[138^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])], fullData[132^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[5])]+fullData[247^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[5])], fullData[169^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])]-fullData[156^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])], fullData[16^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])]+fullData[89^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])], fullData[100^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])]^fullData[42^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])], fullData[241^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]-fullData[136^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[110^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])]+fullData[126^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])], fullData[248^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]-fullData[133^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[179^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])]+fullData[147^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])], fullData[13^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])]-fullData[11^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])], fullData[210^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]-fullData[128^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[113^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])]^fullData[96^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])], fullData[158^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])]^fullData[182^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])], fullData[254^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]-fullData[232^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[56^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])]+fullData[57^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])], fullData[144^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[1])]+fullData[152^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[1])], fullData[170^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[1])]-fullData[187^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[1])], fullData[57^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[9])]^fullData[109^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[9])], fullData[163^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])]+fullData[157^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[0])], fullData[46^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[8])]-fullData[85^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[8])], fullData[81^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])]^fullData[48^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])], fullData[172^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[2])]^fullData[222^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[2])], fullData[120^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])]^fullData[115^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])], fullData[25^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[9])]+fullData[33^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[9])], fullData[211^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]+fullData[247^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[219^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[5])]+fullData[229^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[5])], fullData[226^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])]-fullData[201^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[6])], fullData[68^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])]+fullData[66^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])], fullData[82^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])]+fullData[114^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[3])], fullData[67^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])]+fullData[84^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[7])], fullData[218^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[5])]^fullData[211^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[5])], fullData[31^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])]+fullData[55^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[4])], fullData[5^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[8])]+fullData[30^ /*line x7tVtyaROoUB.go:1*/ int(idxKey[8])])
			return /*line x7tVtyaROoUB.go:1*/ string(data)
		}(), "error", chyZ8Q)
		return /*line BFmSuoR.go:1*/ errors.New(func() /*line jhckJNXn0s.go:1*/ string {
			seed := /*line jhckJNXn0s.go:1*/ byte(81)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line jhckJNXn0s.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line jhckJNXn0s.go:1*/ fnc(183)(105)(218)(183)(103)(205)(86)(0)(251)(167)(160)(51)(116)(228)(197)(148)(23)(233)(28)(61)(128)(253)(246)(223)(201)(70)(226)(179)(115)(231)(196)(142)(27)
			return /*line jhckJNXn0s.go:1*/ string(data)
		}())
	}
	if ihZoVs != journalVersion {
		/*line BoaVl2kJbk.go:1*/ log.Warn(func() /*line d0amKdqyj.go:1*/ string {
			data := /*line d0amKdqyj.go:1*/ make([]byte, 0, 50)
			i := 11
			decryptKey := 243
			for counter := 0; i != 10; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 14:
					i = 15
					data = /*line d0amKdqyj.go:1*/ append(data, 63)
				case 7:
					i = 0
					data = /*line d0amKdqyj.go:1*/ append(data, "72"...,
					)
				case 2:
					i = 10
					for y := range data {
						data[y] = data[y] ^ /*line d0amKdqyj.go:1*/ byte(decryptKey^y)
					}
				case 3:
					i = 18
					data = /*line d0amKdqyj.go:1*/ append(data, 4)
				case 11:
					data = /*line d0amKdqyj.go:1*/ append(data, "\x13?&"...,
					)
					i = 7
				case 5:
					i = 3
					data = /*line d0amKdqyj.go:1*/ append(data, "\x1dT"...,
					)
				case 6:
					data = /*line d0amKdqyj.go:1*/ append(data, "^\v\x19"...,
					)
					i = 19
				case 12:
					i = 14
					data = /*line d0amKdqyj.go:1*/ append(data, 59)
				case 19:
					data = /*line d0amKdqyj.go:1*/ append(data, "\t\t\x10\x17"...,
					)
					i = 4
				case 0:
					i = 13
					data = /*line d0amKdqyj.go:1*/ append(data, 32)
				case 16:
					data = /*line d0amKdqyj.go:1*/ append(data, ")4"...,
					)
					i = 9
				case 4:
					data = /*line d0amKdqyj.go:1*/ append(data, 9)
					i = 2
				case 15:
					i = 17
					data = /*line d0amKdqyj.go:1*/ append(data, "\"*&"...,
					)
				case 20:
					i = 8
					data = /*line d0amKdqyj.go:1*/ append(data, "6&66"...,
					)
				case 9:
					data = /*line d0amKdqyj.go:1*/ append(data, ">z*"...,
					)
					i = 20
				case 17:
					data = /*line d0amKdqyj.go:1*/ append(data, "i?\x1e\x02"...,
					)
					i = 5
				case 8:
					i = 1
					data = /*line d0amKdqyj.go:1*/ append(data, ",,"...,
					)
				case 1:
					data = /*line d0amKdqyj.go:1*/ append(data, "6a* "...,
					)
					i = 12
				case 13:
					data = /*line d0amKdqyj.go:1*/ append(data, "55;~"...,
					)
					i = 16
				case 18:
					i = 6
					data = /*line d0amKdqyj.go:1*/ append(data, "\x00\x1e\x1e\x18"...,
					)
				}
			}
			return /*line d0amKdqyj.go:1*/ string(data)
		}(), func() /*line c5wvGBYpBqzr.go:1*/ string {
			data := /*line c5wvGBYpBqzr.go:1*/ make([]byte, 0, 9)
			i := 7
			decryptKey := 221
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 9:
					i = 1
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 156)
				case 6:
					for y := range data {
						data[y] = data[y] ^ /*line c5wvGBYpBqzr.go:1*/ byte(decryptKey^y)
					}
					i = 3
				case 8:
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 153)
					i = 9
				case 2:
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 142)
					i = 8
				case 7:
					i = 2
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 152)
				case 4:
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 157)
					i = 5
				case 5:
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 137)
					i = 0
				case 1:
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 135)
					i = 4
				case 0:
					i = 6
					data = /*line c5wvGBYpBqzr.go:1*/ append(data, 137)
				}
			}
			return /*line c5wvGBYpBqzr.go:1*/ string(data)
		}(), journalVersion, "got", ihZoVs)
		return /*line xreOdKrwVcAK.go:1*/ errors.New(func() /*line R1mvDope.go:1*/ string {
			fullData := [] /*line R1mvDope.go:1*/ byte("\x11ϯY\x81@\x81\xa4x\x11l\xa1\xba\xe2\xe0\xa1װ\xf5\xea\x02\xc1^\xab\xd5c\xd2\xd6\xc2f\x13\xd4\xe7\xbcz\x9c!̇9\x14\x9a")
			idxKey := [] /*line R1mvDope.go:1*/ byte("Mv\xe2ͅ\x93\xd97\x80E\x17t\x1a\x06")
			data := /*line R1mvDope.go:1*/ make([]byte, 0, 22)
			data = /*line R1mvDope.go:1*/ append(data, fullData[9^ /*line R1mvDope.go:1*/ int(idxKey[13])]^fullData[29^ /*line R1mvDope.go:1*/ int(idxKey[13])], fullData[9^ /*line R1mvDope.go:1*/ int(idxKey[12])]-fullData[18^ /*line R1mvDope.go:1*/ int(idxKey[12])], fullData[28^ /*line R1mvDope.go:1*/ int(idxKey[13])]-fullData[31^ /*line R1mvDope.go:1*/ int(idxKey[13])], fullData[246^ /*line R1mvDope.go:1*/ int(idxKey[2])]+fullData[232^ /*line R1mvDope.go:1*/ int(idxKey[2])], fullData[32^ /*line R1mvDope.go:1*/ int(idxKey[7])]^fullData[18^ /*line R1mvDope.go:1*/ int(idxKey[7])], fullData[141^ /*line R1mvDope.go:1*/ int(idxKey[8])]^fullData[156^ /*line R1mvDope.go:1*/ int(idxKey[8])], fullData[106^ /*line R1mvDope.go:1*/ int(idxKey[0])]-fullData[76^ /*line R1mvDope.go:1*/ int(idxKey[0])], fullData[12^ /*line R1mvDope.go:1*/ int(idxKey[12])]+fullData[26^ /*line R1mvDope.go:1*/ int(idxKey[12])], fullData[60^ /*line R1mvDope.go:1*/ int(idxKey[7])]^fullData[40^ /*line R1mvDope.go:1*/ int(idxKey[7])], fullData[183^ /*line R1mvDope.go:1*/ int(idxKey[5])]-fullData[145^ /*line R1mvDope.go:1*/ int(idxKey[5])], fullData[0^ /*line R1mvDope.go:1*/ int(idxKey[13])]-fullData[24^ /*line R1mvDope.go:1*/ int(idxKey[13])], fullData[22^ /*line R1mvDope.go:1*/ int(idxKey[12])]-fullData[25^ /*line R1mvDope.go:1*/ int(idxKey[12])], fullData[145^ /*line R1mvDope.go:1*/ int(idxKey[8])]+fullData[161^ /*line R1mvDope.go:1*/ int(idxKey[8])], fullData[20^ /*line R1mvDope.go:1*/ int(idxKey[12])]+fullData[31^ /*line R1mvDope.go:1*/ int(idxKey[12])], fullData[151^ /*line R1mvDope.go:1*/ int(idxKey[5])]+fullData[129^ /*line R1mvDope.go:1*/ int(idxKey[5])], fullData[66^ /*line R1mvDope.go:1*/ int(idxKey[9])]^fullData[80^ /*line R1mvDope.go:1*/ int(idxKey[9])], fullData[105^ /*line R1mvDope.go:1*/ int(idxKey[11])]^fullData[92^ /*line R1mvDope.go:1*/ int(idxKey[11])], fullData[242^ /*line R1mvDope.go:1*/ int(idxKey[2])]+fullData[193^ /*line R1mvDope.go:1*/ int(idxKey[2])], fullData[103^ /*line R1mvDope.go:1*/ int(idxKey[9])]-fullData[76^ /*line R1mvDope.go:1*/ int(idxKey[9])], fullData[108^ /*line R1mvDope.go:1*/ int(idxKey[11])]+fullData[93^ /*line R1mvDope.go:1*/ int(idxKey[11])], fullData[196^ /*line R1mvDope.go:1*/ int(idxKey[2])]+fullData[194^ /*line R1mvDope.go:1*/ int(idxKey[2])])
			return /*line R1mvDope.go:1*/ string(data)
		}())
	}

	var f3dWTs7 common.Hash
	if chyZ8Q := /*line FYuaNdmyQ.go:1*/ vaD4aIy.Decode(&f3dWTs7); chyZ8Q != nil {
		return /*line OGDVIhfq.go:1*/ errors.New(func() /*line b3rkBsE9l.go:1*/ string {
			data := /*line b3rkBsE9l.go:1*/ make([]byte, 0, 24)
			i := 2
			decryptKey := 169
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					i = 9
					data = /*line b3rkBsE9l.go:1*/ append(data, 128)
				case 6:
					i = 0
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x97\x8f"...,
					)
				case 4:
					i = 12
					data = /*line b3rkBsE9l.go:1*/ append(data, 128)
				case 11:
					data = /*line b3rkBsE9l.go:1*/ append(data, 155)
					i = 3
				case 1:
					i = 11
					data = /*line b3rkBsE9l.go:1*/ append(data, 130)
				case 12:
					i = 10
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x86\x8c\xca"...,
					)
				case 3:
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x98\x8e\xdf"...,
					)
					i = 7
				case 9:
					i = 4
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x85\x9c\x9d"...,
					)
				case 0:
					i = 8
					for y := range data {
						data[y] = data[y] ^ /*line b3rkBsE9l.go:1*/ byte(decryptKey^y)
					}
				case 7:
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x8c\x96"...,
					)
					i = 6
				case 10:
					i = 5
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x81\x8d\x94"...,
					)
				case 5:
					i = 1
					data = /*line b3rkBsE9l.go:1*/ append(data, "\x8d\xc1\x8c"...,
					)
				}
			}
			return /*line b3rkBsE9l.go:1*/ string(data)
		}())
	}
	if koxaxufVjnP := /*line ERZumpIPNF.go:1*/ rawdb.ReadSnapshotRoot(lDBEbd); koxaxufVjnP != f3dWTs7 {
		/*line DknYblRxxl6.go:1*/ log.Warn(func() /*line NU4Jyz1Uf.go:1*/ string {
			fullData := [] /*line NU4Jyz1Uf.go:1*/ byte("?4A\x98\x02\x8ej\xffy\xae\xcde\x05\x13\x85\xa8\x83\xe7\xfe\xb2;(j\x06\x02\x1fq.>r[&\xf5~\x0ex^%b\xacG\x02\xb5\xf7\xe7\xcf")
			idxKey := [] /*line NU4Jyz1Uf.go:1*/ byte("\x8e\x17\x8a<\xb8|")
			data := /*line NU4Jyz1Uf.go:1*/ make([]byte, 0, 24)
			data = /*line NU4Jyz1Uf.go:1*/ append(data, fullData[156^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])]^fullData[157^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])], fullData[38^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])]-fullData[36^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])], fullData[86^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[5])]+fullData[91^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[5])], fullData[157^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])]^fullData[186^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])], fullData[183^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])]^fullData[178^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])], fullData[43^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])]^fullData[26^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])], fullData[63^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])]-fullData[31^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])], fullData[151^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])]-fullData[141^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])], fullData[55^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[1])]+fullData[31^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[1])], fullData[167^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])]^fullData[131^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])], fullData[12^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[1])]^fullData[51^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[1])], fullData[172^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])]+fullData[133^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])], fullData[120^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[5])]^fullData[122^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[5])], fullData[48^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])]+fullData[42^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])], fullData[163^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])]-fullData[143^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])], fullData[184^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])]-fullData[161^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])], fullData[154^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])]+fullData[155^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[2])], fullData[41^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])]+fullData[20^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[3])], fullData[153^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])]+fullData[147^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[4])], fullData[143^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])]+fullData[146^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])], fullData[131^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])]+fullData[144^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[0])], fullData[99^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[5])]+fullData[104^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[5])], fullData[59^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[1])]+fullData[25^ /*line NU4Jyz1Uf.go:1*/ int(idxKey[1])])
			return /*line NU4Jyz1Uf.go:1*/ string(data)
		}(), func() /*line sqeKdVMnVs.go:1*/ string {
			fullData := [] /*line sqeKdVMnVs.go:1*/ byte("ܰ\xae\xf5\x86\xb7\xd2\xc4\xde>Z2-\ná")
			idxKey := [] /*line sqeKdVMnVs.go:1*/ byte("\xb10\x9bJ\x18")
			data := /*line sqeKdVMnVs.go:1*/ make([]byte, 0, 9)
			data = /*line sqeKdVMnVs.go:1*/ append(data, fullData[21^ /*line sqeKdVMnVs.go:1*/ int(idxKey[4])]+fullData[18^ /*line sqeKdVMnVs.go:1*/ int(idxKey[4])], fullData[147^ /*line sqeKdVMnVs.go:1*/ int(idxKey[2])]^fullData[158^ /*line sqeKdVMnVs.go:1*/ int(idxKey[2])], fullData[69^ /*line sqeKdVMnVs.go:1*/ int(idxKey[3])]+fullData[76^ /*line sqeKdVMnVs.go:1*/ int(idxKey[3])], fullData[60^ /*line sqeKdVMnVs.go:1*/ int(idxKey[1])]+fullData[57^ /*line sqeKdVMnVs.go:1*/ int(idxKey[1])], fullData[24^ /*line sqeKdVMnVs.go:1*/ int(idxKey[4])]^fullData[26^ /*line sqeKdVMnVs.go:1*/ int(idxKey[4])], fullData[51^ /*line sqeKdVMnVs.go:1*/ int(idxKey[1])]-fullData[52^ /*line sqeKdVMnVs.go:1*/ int(idxKey[1])], fullData[186^ /*line sqeKdVMnVs.go:1*/ int(idxKey[0])]-fullData[191^ /*line sqeKdVMnVs.go:1*/ int(idxKey[0])], fullData[176^ /*line sqeKdVMnVs.go:1*/ int(idxKey[0])]+fullData[182^ /*line sqeKdVMnVs.go:1*/ int(idxKey[0])])
			return /*line sqeKdVMnVs.go:1*/ string(data)
		}(), koxaxufVjnP, "diffs", func() /*line RLKeqrvrz.go:1*/ string {
			key := [] /*line RLKeqrvrz.go:1*/ byte("Fʭ\x16\xb1\"\xe1Z.")
			data := [] /*line RLKeqrvrz.go:1*/ byte("\xbb8\x1aw%\x85I\xbf\x92")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line RLKeqrvrz.go:1*/ string(data)
		}())
		return /*line UPDfjyu.go:1*/ errors.New(func() /*line WAdUCZe40vi.go:1*/ string {
			data := /*line WAdUCZe40vi.go:1*/ make([]byte, 0, 32)
			i := 7
			decryptKey := 15
			for counter := 0; i != 14; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 12:
					i = 1
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\x00\x14\x00\x06"...,
					)
				case 6:
					i = 0
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xec\xfa\xed"...,
					)
				case 5:
					data = /*line WAdUCZe40vi.go:1*/ append(data, 182)
					i = 6
				case 9:
					i = 10
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xb1\xf6\x00"...,
					)
				case 13:
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xf9\xf3"...,
					)
					i = 11
				case 1:
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xf8\xf8"...,
					)
					i = 9
				case 0:
					i = 13
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xaa\xf3"...,
					)
				case 3:
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xe2\xfb\xec\xfa"...,
					)
					i = 8
				case 7:
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\b\x05\f\a"...,
					)
					i = 12
				case 10:
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\v\x00"...,
					)
					i = 5
				case 8:
					data = /*line WAdUCZe40vi.go:1*/ append(data, 248)
					i = 4
				case 11:
					i = 2
					data = /*line WAdUCZe40vi.go:1*/ append(data, 244)
				case 2:
					i = 3
					data = /*line WAdUCZe40vi.go:1*/ append(data, "\xa3\xf0"...,
					)
				case 4:
					i = 14
					for y := range data {
						data[y] = data[y] + /*line WAdUCZe40vi.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return /*line WAdUCZe40vi.go:1*/ string(data)
		}())
	}
	for {
		var (
			z1BBTN2UX    common.Hash
			zoMmGkL      []journalDestruct
			bQIIyfhppAL1 []journalAccount
			agSpCMzc     []journalStorage
			vgaPgkN      = /*line c2WI26Ld.go:1*/ make(map[common.Hash]struct{})
			oVnS8XskWEE  = /*line SAdYTsOK45.go:1*/ make(map[common.Hash][]byte)
			bJNw07c3Namn = /*line XBpA4djwnm6j.go:1*/ make(map[common.Hash]map[common.Hash][]byte)
		)

		if chyZ8Q := /*line _FB1W_.go:1*/ vaD4aIy.Decode(&z1BBTN2UX); chyZ8Q != nil {

			if /*line MyRPCavOs.go:1*/ errors.Is(chyZ8Q, io.EOF) {
				return nil
			}
			return /*line JueaaZnKYa.go:1*/ fmt.Errorf(func() /*line GKEoii2P.go:1*/ string {
				seed := /*line GKEoii2P.go:1*/ byte(0)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line GKEoii2P.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line GKEoii2P.go:1*/ fnc(108)(219)(168)(83)(98)(8)(21)(39)(78)(86)(254)(249)(242)(233)(152)(22)(49)(179)
				return /*line GKEoii2P.go:1*/ string(data)
			}(), chyZ8Q)
		}
		if chyZ8Q := /*line RQKbjA.go:1*/ vaD4aIy.Decode(&zoMmGkL); chyZ8Q != nil {
			return /*line E9lw9DDHQB.go:1*/ fmt.Errorf(func() /*line DYNsMM6el4Go.go:1*/ string {
				data := /*line DYNsMM6el4Go.go:1*/ make([]byte, 0, 24)
				i := 0
				decryptKey := 0
				for counter := 0; i != 2; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 7:
						data = /*line DYNsMM6el4Go.go:1*/ append(data, "09/"...,
						)
						i = 4
					case 8:
						data = /*line DYNsMM6el4Go.go:1*/ append(data, "$c$$"...,
						)
						i = 3
					case 5:
						data = /*line DYNsMM6el4Go.go:1*/ append(data, "-n+%"...,
						)
						i = 1
					case 6:
						data = /*line DYNsMM6el4Go.go:1*/ append(data, 42)
						i = 9
					case 4:
						data = /*line DYNsMM6el4Go.go:1*/ append(data, "+c~z"...,
						)
						i = 6
					case 9:
						for y := range data {
							data[y] = data[y] ^ /*line DYNsMM6el4Go.go:1*/ byte(decryptKey^y)
						}
						i = 2
					case 0:
						i = 5
						data = /*line DYNsMM6el4Go.go:1*/ append(data, "&$)"...,
						)
					case 1:
						data = /*line DYNsMM6el4Go.go:1*/ append(data, 43)
						i = 8
					case 3:
						i = 7
						data = /*line DYNsMM6el4Go.go:1*/ append(data, "536"...,
						)
					}
				}
				return /*line DYNsMM6el4Go.go:1*/ string(data)
			}(), chyZ8Q)
		}
		if chyZ8Q := /*line W2OuQv.go:1*/ vaD4aIy.Decode(&bQIIyfhppAL1); chyZ8Q != nil {
			return /*line eDa2BWScWZ.go:1*/ fmt.Errorf(func() /*line uR5qPEfQt_.go:1*/ string {
				key := [] /*line uR5qPEfQt_.go:1*/ byte("#3\x8du\xadqbZU)\x1e\xe5\x98v6(\xdcZ\x03%v\x19")
				data := [] /*line uR5qPEfQt_.go:1*/ byte("I<\xd4\xefs\xf3\a\f\x11\xf7C~\xcb\xf9?F\x98\x197\xfb\xaf]")
				for i, b := range key {
					data[i] = data[i] + b
				}
				return /*line uR5qPEfQt_.go:1*/ string(data)
			}(), chyZ8Q)
		}
		if chyZ8Q := /*line qcuB8tJ0ZvA.go:1*/ vaD4aIy.Decode(&agSpCMzc); chyZ8Q != nil {
			return /*line mxUnLUCdbTt.go:1*/ fmt.Errorf(func() /*line zFNaVOFBlmaC.go:1*/ string {
				seed := /*line zFNaVOFBlmaC.go:1*/ byte(88)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line zFNaVOFBlmaC.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line zFNaVOFBlmaC.go:1*/ fnc(20)(3)(242)(3)(188)(68)(5)(253)(0)(186)(83)(1)(251)(3)(239)(6)(254)(213)(230)(5)(81)
				return /*line zFNaVOFBlmaC.go:1*/ string(data)
			}(), chyZ8Q)
		}
		for _, gUZorXrZXPhE := range zoMmGkL {
			vgaPgkN[gUZorXrZXPhE.Hash] = struct{}{}
		}
		for _, gUZorXrZXPhE := range bQIIyfhppAL1 {
			if /*line IKjIad_3FBuF.go:1*/ len(gUZorXrZXPhE.Blob) > 0 {
				oVnS8XskWEE[gUZorXrZXPhE.Hash] = gUZorXrZXPhE.Blob
			} else {
				oVnS8XskWEE[gUZorXrZXPhE.Hash] = nil
			}
		}
		for _, gUZorXrZXPhE := range agSpCMzc {
			etj53bd3zM := /*line apyFkdhfQZu.go:1*/ make(map[common.Hash][]byte)
			for fSpMhzHR, nVUwQz8Zi := range gUZorXrZXPhE.Keys {
				if /*line jVVCenK.go:1*/ len(gUZorXrZXPhE.Vals[fSpMhzHR]) > 0 {
					etj53bd3zM[nVUwQz8Zi] = gUZorXrZXPhE.Vals[fSpMhzHR]
				} else {
					etj53bd3zM[nVUwQz8Zi] = nil
				}
			}
			bJNw07c3Namn[gUZorXrZXPhE.Hash] = etj53bd3zM
		}
		if chyZ8Q := /*line Ku4jbYTmh.go:1*/ kGNYh2Kjh(f3dWTs7, z1BBTN2UX, vgaPgkN, oVnS8XskWEE, bJNw07c3Namn); chyZ8Q != nil {
			return chyZ8Q
		}
		f3dWTs7 = z1BBTN2UX
	}
}

var _ bytes.Buffer
var _ = binary.Append
var _ = errors.As
var _ = fmt.Append
var _ io.ByteReader

const _ = time.ANSIC

var _ fastcache.BigStats
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ = log.Crit
var _ = rlp.AppendUint64
var _ triedb.Config
