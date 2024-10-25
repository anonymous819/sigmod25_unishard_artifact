//line :1
package trie

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

func (hkI2UXG_QBd *Trie) Prove(lhQWH7m []byte, zzPkun7 ethdb.KeyValueWriter) error {

	if hkI2UXG_QBd.committed {
		return MWbaoG
	}

	var (
		tKBrjZ     []byte
		y5wkTqRU   []node
		owwtC8asmI = hkI2UXG_QBd.root
	)
	lhQWH7m = /*line AZaZszLdiUK.go:1*/ nd12pU880Z8(lhQWH7m)
	for /*line G5KnLUDU.go:1*/ len(lhQWH7m) > 0 && owwtC8asmI != nil {
		switch gnGMaXTuiFeE := owwtC8asmI.(type) {
		case *qUKQUCfTXB:
			if /*line Mkeh0hk.go:1*/ len(lhQWH7m) < /*line SY5heBPTeE.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A) || ! /*line awQvd2.go:1*/ bytes.Equal(gnGMaXTuiFeE.ANdZYqbzJ1A, lhQWH7m[: /*line NSOosqHL.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A)]) {

				owwtC8asmI = nil
			} else {
				owwtC8asmI = gnGMaXTuiFeE.YpmEYGB
				tKBrjZ = /*line Iphrcz.go:1*/ append(tKBrjZ, gnGMaXTuiFeE.ANdZYqbzJ1A...)
				lhQWH7m = lhQWH7m[ /*line r2zIJ4HfH.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A):]
			}
			y5wkTqRU = /*line KXBqLrSnbeM.go:1*/ append(y5wkTqRU, gnGMaXTuiFeE)
		case *fullNode:
			owwtC8asmI = gnGMaXTuiFeE.Children[lhQWH7m[0]]
			tKBrjZ = /*line agtVtLrK006.go:1*/ append(tKBrjZ, lhQWH7m[0])
			lhQWH7m = lhQWH7m[1:]
			y5wkTqRU = /*line AmPtqlOd3h.go:1*/ append(y5wkTqRU, gnGMaXTuiFeE)
		case hashNode:

			aBHduJG, eZzE0KPS := /*line LTFTLXMa9.go:1*/ hkI2UXG_QBd.reader.uAjz8NxU_cL(tKBrjZ /*line sePOzXrf.go:1*/, common.BytesToHash(gnGMaXTuiFeE))
			if eZzE0KPS != nil {
				/*line J2O0uLTPjgkn.go:1*/ log.Error(func() /*line uB4GVdRKEuqv.go:1*/ string {
					data := [] /*line uB4GVdRKEuqv.go:1*/ byte("\x16\th\xae\xdb!\xd2\x05df^\x95\xb2\n\xdc\xc8\x13\xabo\xc1 \x05\x9f\xf9l\xda*\xfe\xc9P(0v\xd5")
					positions := [...]byte{9, 3, 10, 19, 24, 12, 33, 14, 31, 0, 23, 22, 17, 0, 0, 27, 23, 12, 13, 11, 6, 33, 13, 15, 19, 9, 5, 16, 7, 28, 1, 7, 10, 30, 17, 13, 24, 21, 9, 3, 23, 4, 9, 25, 33, 28, 13, 26}
					for i := 0; i < 48; i += 2 {
						localKey := /*line uB4GVdRKEuqv.go:1*/ byte(i) + /*line uB4GVdRKEuqv.go:1*/ byte(positions[i]^positions[i+1]) + 128
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
					}
					return /*line uB4GVdRKEuqv.go:1*/ string(data)
				}(), "err", eZzE0KPS)
				return eZzE0KPS
			}

			owwtC8asmI = /*line FEhxYFVLpzLz.go:1*/ hct27E(gnGMaXTuiFeE, aBHduJG)
		default:
			/*line pEX1A7jIay.go:1*/ panic( /*line unBVP0wu3zR.go:1*/ fmt.Sprintf(func() /*line TkcozGuopi97.go:1*/ string {
				data := [] /*line TkcozGuopi97.go:1*/ byte("\xb7\xd2:\xc0\x1cnv\xd4\xc5\xdc\xd7 \x04\x99\x1d{: %v")
				positions := [...]byte{9, 12, 9, 1, 13, 10, 14, 10, 1, 10, 4, 14, 3, 15, 7, 14, 8, 10, 0, 14}
				for i := 0; i < 20; i += 2 {
					localKey := /*line TkcozGuopi97.go:1*/ byte(i) + /*line TkcozGuopi97.go:1*/ byte(positions[i]^positions[i+1]) + 141
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line TkcozGuopi97.go:1*/ string(data)
			}(), owwtC8asmI, owwtC8asmI))
		}
	}
	d6imPJ := /*line qPWNN91.go:1*/ dCZJziX(false)
	defer /*line i5MdQt.go:1*/ ptIMYRK9(d6imPJ)

	for q2u2020KD6, gnGMaXTuiFeE := range y5wkTqRU {
		var ifmSiSN3 node
		gnGMaXTuiFeE, ifmSiSN3 = /*line HWkla2J71.go:1*/ d6imPJ.aCBICr2(gnGMaXTuiFeE)
		if rNuN0sMPDJ, yY_yPSd8vG := ifmSiSN3.(hashNode); yY_yPSd8vG || q2u2020KD6 == 0 {

			bh98pCjEm := /*line ANECfv9No6u.go:1*/ kG18lbsZr(gnGMaXTuiFeE)
			if !yY_yPSd8vG {
				rNuN0sMPDJ = /*line a0HYmct1R.go:1*/ d6imPJ.jfg9zxq(bh98pCjEm)
			}
			/*line vEa01SkhN5.go:1*/ zzPkun7.Put(rNuN0sMPDJ, bh98pCjEm)
		}
	}
	return nil
}

func (hkI2UXG_QBd *IetHqRYp3VR) Prove(lhQWH7m []byte, zzPkun7 ethdb.KeyValueWriter) error {
	return /*line ea6TWWV.go:1*/ hkI2UXG_QBd.dHWIsToseO.Prove(lhQWH7m, zzPkun7)
}

func UuW2zRJDUv(z8PGQUy common.Hash, lhQWH7m []byte, zzPkun7 ethdb.KeyValueReader) (ar76Sw []byte, eZzE0KPS error) {
	lhQWH7m = /*line jelcJ_KvdLMa.go:1*/ nd12pU880Z8(lhQWH7m)
	oNNWzgbv2mp := z8PGQUy
	for q2u2020KD6 := 0; ; q2u2020KD6++ {
		m2t4GBUGH, _ := /*line eNfJg8Kpe.go:1*/ zzPkun7.Get(oNNWzgbv2mp[:])
		if m2t4GBUGH == nil {
			return nil /*line P_T7FtN2bQCa.go:1*/, fmt.Errorf(func() /*line inUn1aM.go:1*/ string {
				seed := /*line inUn1aM.go:1*/ byte(94)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line inUn1aM.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line inUn1aM.go:1*/ fnc(18)(2)(253)(0)(247)(186)(78)(1)(245)(1)(187)(5)(63)(188)(8)(64)(249)(18)(245)(184)(5)(11)(6)(254)(68)(177)(247)(77)(252)(10)(0)(246)(5)(249)
				return /*line inUn1aM.go:1*/ string(data)
			}(), q2u2020KD6, oNNWzgbv2mp)
		}
		gnGMaXTuiFeE, eZzE0KPS := /*line HoBytPCbc.go:1*/ rRV2VeUAEp(oNNWzgbv2mp[:], m2t4GBUGH)
		if eZzE0KPS != nil {
			return nil /*line wKECsc.go:1*/, fmt.Errorf(func() /*line Tt32Gp.go:1*/ string {
				fullData := [] /*line Tt32Gp.go:1*/ byte("\xddє\x17\xfe\x91\xc7}z\x96\u0cdc\xb3mɑs?>\x16\v\x03e\a[\xc7H\x0f\x00\x9b\xea\x14\xde6{]\x9b\x85k\xa9\xee")
				idxKey := [] /*line Tt32Gp.go:1*/ byte("\xf8\x83\xf5\xcf\xddi\x9f\U00064865\x02\xe4m9\x06")
				data := /*line Tt32Gp.go:1*/ make([]byte, 0, 22)
				data = /*line Tt32Gp.go:1*/ append(data, fullData[239^ /*line Tt32Gp.go:1*/ int(idxKey[7])]+fullData[235^ /*line Tt32Gp.go:1*/ int(idxKey[7])], fullData[183^ /*line Tt32Gp.go:1*/ int(idxKey[10])]-fullData[132^ /*line Tt32Gp.go:1*/ int(idxKey[10])], fullData[23^ /*line Tt32Gp.go:1*/ int(idxKey[15])]-fullData[26^ /*line Tt32Gp.go:1*/ int(idxKey[15])], fullData[222^ /*line Tt32Gp.go:1*/ int(idxKey[0])]+fullData[221^ /*line Tt32Gp.go:1*/ int(idxKey[0])], fullData[200^ /*line Tt32Gp.go:1*/ int(idxKey[4])]+fullData[202^ /*line Tt32Gp.go:1*/ int(idxKey[4])], fullData[16^ /*line Tt32Gp.go:1*/ int(idxKey[15])]-fullData[3^ /*line Tt32Gp.go:1*/ int(idxKey[15])], fullData[132^ /*line Tt32Gp.go:1*/ int(idxKey[8])]+fullData[189^ /*line Tt32Gp.go:1*/ int(idxKey[8])], fullData[118^ /*line Tt32Gp.go:1*/ int(idxKey[5])]-fullData[74^ /*line Tt32Gp.go:1*/ int(idxKey[5])], fullData[215^ /*line Tt32Gp.go:1*/ int(idxKey[4])]-fullData[213^ /*line Tt32Gp.go:1*/ int(idxKey[4])], fullData[45^ /*line Tt32Gp.go:1*/ int(idxKey[14])]^fullData[27^ /*line Tt32Gp.go:1*/ int(idxKey[14])], fullData[221^ /*line Tt32Gp.go:1*/ int(idxKey[2])]^fullData[243^ /*line Tt32Gp.go:1*/ int(idxKey[2])], fullData[2^ /*line Tt32Gp.go:1*/ int(idxKey[15])]^fullData[22^ /*line Tt32Gp.go:1*/ int(idxKey[15])], fullData[164^ /*line Tt32Gp.go:1*/ int(idxKey[1])]-fullData[155^ /*line Tt32Gp.go:1*/ int(idxKey[1])], fullData[108^ /*line Tt32Gp.go:1*/ int(idxKey[13])]+fullData[111^ /*line Tt32Gp.go:1*/ int(idxKey[13])], fullData[5^ /*line Tt32Gp.go:1*/ int(idxKey[11])]-fullData[38^ /*line Tt32Gp.go:1*/ int(idxKey[11])], fullData[254^ /*line Tt32Gp.go:1*/ int(idxKey[2])]^fullData[252^ /*line Tt32Gp.go:1*/ int(idxKey[2])], fullData[36^ /*line Tt32Gp.go:1*/ int(idxKey[14])]-fullData[53^ /*line Tt32Gp.go:1*/ int(idxKey[14])], fullData[222^ /*line Tt32Gp.go:1*/ int(idxKey[4])]-fullData[221^ /*line Tt32Gp.go:1*/ int(idxKey[4])], fullData[168^ /*line Tt32Gp.go:1*/ int(idxKey[10])]+fullData[171^ /*line Tt32Gp.go:1*/ int(idxKey[10])], fullData[68^ /*line Tt32Gp.go:1*/ int(idxKey[13])]-fullData[98^ /*line Tt32Gp.go:1*/ int(idxKey[13])], fullData[220^ /*line Tt32Gp.go:1*/ int(idxKey[3])]^fullData[212^ /*line Tt32Gp.go:1*/ int(idxKey[3])])
				return /*line Tt32Gp.go:1*/ string(data)
			}(), q2u2020KD6, eZzE0KPS)
		}
		zB0wlXXn, nPjnWCe := /*line FL_7yuS1rsPS.go:1*/ cga2ibNy(gnGMaXTuiFeE, lhQWH7m, true)
		switch nPjnWCe := nPjnWCe.(type) {
		case nil:

			return nil, nil
		case hashNode:
			lhQWH7m = zB0wlXXn
			/*line K2J3os303u.go:1*/ copy(oNNWzgbv2mp[:], nPjnWCe)
		case valueNode:
			return nPjnWCe, nil
		}
	}
}

func hClFfPa(z8PGQUy common.Hash, m7S8SE node, lhQWH7m []byte, zzPkun7 ethdb.KeyValueReader, kTSWY1QX bool) (node, []byte, error) {

	oGS_E_Q := func(rNuN0sMPDJ common.Hash) (node, error) {
		m2t4GBUGH, _ := /*line duT4fZ.go:1*/ zzPkun7.Get(rNuN0sMPDJ[:])
		if m2t4GBUGH == nil {
			return nil /*line JkrNTou.go:1*/, fmt.Errorf(func() /*line DBfXYopD.go:1*/ string {
				data := [] /*line DBfXYopD.go:1*/ byte("\xda\t\xa5\xf1\xa9\xc9n\x1c|e\x12\xc2\a\x8b\xba\x9b \x90\xdf\xfa\xb6x\x04\xe6Qi\xbf\f\xa3\t\n")
				positions := [...]byte{0, 5, 17, 18, 18, 23, 7, 4, 0, 27, 28, 2, 30, 15, 19, 8, 26, 10, 20, 11, 13, 24, 14, 22, 30, 3, 24, 7, 19, 13, 28, 29, 12, 12, 2, 1}
				for i := 0; i < 36; i += 2 {
					localKey := /*line DBfXYopD.go:1*/ byte(i) + /*line DBfXYopD.go:1*/ byte(positions[i]^positions[i+1]) + 65
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line DBfXYopD.go:1*/ string(data)
			}(), rNuN0sMPDJ)
		}
		gnGMaXTuiFeE, eZzE0KPS := /*line dEjNkhq.go:1*/ rRV2VeUAEp(rNuN0sMPDJ[:], m2t4GBUGH)
		if eZzE0KPS != nil {
			return nil /*line SqlIqtAY.go:1*/, fmt.Errorf(func() /*line RosAHWg.go:1*/ string {
				fullData := [] /*line RosAHWg.go:1*/ byte("Y\x8d\xe6\xe2\xaf\xeef1xϕ92\x8c0k\xc8\t\xc3+\xa4\x0e6ꬠR\xea]\x89qZ\x815")
				idxKey := [] /*line RosAHWg.go:1*/ byte("\x9d\x03")
				data := /*line RosAHWg.go:1*/ make([]byte, 0, 18)
				data = /*line RosAHWg.go:1*/ append(data, fullData[138^ /*line RosAHWg.go:1*/ int(idxKey[0])]+fullData[149^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[134^ /*line RosAHWg.go:1*/ int(idxKey[0])]-fullData[128^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[21^ /*line RosAHWg.go:1*/ int(idxKey[1])]^fullData[25^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[7^ /*line RosAHWg.go:1*/ int(idxKey[1])]+fullData[29^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[132^ /*line RosAHWg.go:1*/ int(idxKey[0])]-fullData[147^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[137^ /*line RosAHWg.go:1*/ int(idxKey[0])]-fullData[145^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[156^ /*line RosAHWg.go:1*/ int(idxKey[0])]^fullData[158^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[18^ /*line RosAHWg.go:1*/ int(idxKey[1])]+fullData[5^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[34^ /*line RosAHWg.go:1*/ int(idxKey[1])]+fullData[4^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[14^ /*line RosAHWg.go:1*/ int(idxKey[1])]^fullData[27^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[141^ /*line RosAHWg.go:1*/ int(idxKey[0])]-fullData[130^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[6^ /*line RosAHWg.go:1*/ int(idxKey[1])]+fullData[35^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[151^ /*line RosAHWg.go:1*/ int(idxKey[0])]+fullData[148^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[146^ /*line RosAHWg.go:1*/ int(idxKey[0])]^fullData[136^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[3^ /*line RosAHWg.go:1*/ int(idxKey[1])]-fullData[8^ /*line RosAHWg.go:1*/ int(idxKey[1])], fullData[143^ /*line RosAHWg.go:1*/ int(idxKey[0])]^fullData[159^ /*line RosAHWg.go:1*/ int(idxKey[0])], fullData[129^ /*line RosAHWg.go:1*/ int(idxKey[0])]^fullData[142^ /*line RosAHWg.go:1*/ int(idxKey[0])])
				return /*line RosAHWg.go:1*/ string(data)
			}(), eZzE0KPS)
		}
		return gnGMaXTuiFeE, eZzE0KPS
	}

	if m7S8SE == nil {
		gnGMaXTuiFeE, eZzE0KPS := /*line s3S26roh6_2z.go:1*/ oGS_E_Q(z8PGQUy)
		if eZzE0KPS != nil {
			return nil, nil, eZzE0KPS
		}
		m7S8SE = gnGMaXTuiFeE
	}
	var (
		eZzE0KPS              error
		jcDLabJ7o, nmFIUUSGjl node
		zB0wlXXn              []byte
		r0QDFWcta7R           []byte
	)
	lhQWH7m, nmFIUUSGjl = /*line jP1CVDX_Y85.go:1*/ nd12pU880Z8(lhQWH7m), m7S8SE
	for {
		zB0wlXXn, jcDLabJ7o = /*line xhJg_V7W9N8.go:1*/ cga2ibNy(nmFIUUSGjl, lhQWH7m, false)
		switch nPjnWCe := jcDLabJ7o.(type) {
		case nil:

			if kTSWY1QX {
				return m7S8SE, nil, nil
			}
			return nil, nil /*line eziaWQF.go:1*/, errors.New(func() /*line ukhZRzK9Q.go:1*/ string {
				fullData := [] /*line ukhZRzK9Q.go:1*/ byte("\xe6\xe3\x13\xab\xf8\xda\x1f\xa5\x8cjš\xd8|\xff\xabK\xde\xfc\xb7\x10\x8b\xd0~\x16\x0fh\xb18\xc8+\xa2\xa2q\x8dR\xa16\xda\xf8\x83\xd6æF\xe3\x93<\xc0\xfe\xbb\n+H8\xf9\xc2UQ\xb0\xac\x1feum\xfb")
				idxKey := [] /*line ukhZRzK9Q.go:1*/ byte("QZ\x06\x98vA\xd8*\xdb\xc8uGe\xd5W~")
				data := /*line ukhZRzK9Q.go:1*/ make([]byte, 0, 34)
				data = /*line ukhZRzK9Q.go:1*/ append(data, fullData[213^ /*line ukhZRzK9Q.go:1*/ int(idxKey[8])]^fullData[206^ /*line ukhZRzK9Q.go:1*/ int(idxKey[8])], fullData[74^ /*line ukhZRzK9Q.go:1*/ int(idxKey[0])]+fullData[66^ /*line ukhZRzK9Q.go:1*/ int(idxKey[0])], fullData[105^ /*line ukhZRzK9Q.go:1*/ int(idxKey[4])]+fullData[92^ /*line ukhZRzK9Q.go:1*/ int(idxKey[4])], fullData[65^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])]+fullData[113^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])], fullData[98^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])]+fullData[91^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])], fullData[121^ /*line ukhZRzK9Q.go:1*/ int(idxKey[14])]^fullData[69^ /*line ukhZRzK9Q.go:1*/ int(idxKey[14])], fullData[12^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])]^fullData[13^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])], fullData[55^ /*line ukhZRzK9Q.go:1*/ int(idxKey[4])]+fullData[127^ /*line ukhZRzK9Q.go:1*/ int(idxKey[4])], fullData[109^ /*line ukhZRzK9Q.go:1*/ int(idxKey[10])]+fullData[70^ /*line ukhZRzK9Q.go:1*/ int(idxKey[10])], fullData[207^ /*line ukhZRzK9Q.go:1*/ int(idxKey[9])]-fullData[231^ /*line ukhZRzK9Q.go:1*/ int(idxKey[9])], fullData[125^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])]+fullData[99^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])], fullData[103^ /*line ukhZRzK9Q.go:1*/ int(idxKey[5])]+fullData[109^ /*line ukhZRzK9Q.go:1*/ int(idxKey[5])], fullData[126^ /*line ukhZRzK9Q.go:1*/ int(idxKey[14])]-fullData[77^ /*line ukhZRzK9Q.go:1*/ int(idxKey[14])], fullData[71^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])]-fullData[126^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])], fullData[233^ /*line ukhZRzK9Q.go:1*/ int(idxKey[13])]-fullData[227^ /*line ukhZRzK9Q.go:1*/ int(idxKey[13])], fullData[33^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])]^fullData[10^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])], fullData[123^ /*line ukhZRzK9Q.go:1*/ int(idxKey[4])]^fullData[112^ /*line ukhZRzK9Q.go:1*/ int(idxKey[4])], fullData[149^ /*line ukhZRzK9Q.go:1*/ int(idxKey[13])]-fullData[228^ /*line ukhZRzK9Q.go:1*/ int(idxKey[13])], fullData[201^ /*line ukhZRzK9Q.go:1*/ int(idxKey[6])]^fullData[227^ /*line ukhZRzK9Q.go:1*/ int(idxKey[6])], fullData[49^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])]^fullData[36^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])], fullData[241^ /*line ukhZRzK9Q.go:1*/ int(idxKey[13])]+fullData[229^ /*line ukhZRzK9Q.go:1*/ int(idxKey[13])], fullData[123^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])]-fullData[95^ /*line ukhZRzK9Q.go:1*/ int(idxKey[15])], fullData[16^ /*line ukhZRzK9Q.go:1*/ int(idxKey[7])]-fullData[7^ /*line ukhZRzK9Q.go:1*/ int(idxKey[7])], fullData[238^ /*line ukhZRzK9Q.go:1*/ int(idxKey[8])]-fullData[218^ /*line ukhZRzK9Q.go:1*/ int(idxKey[8])], fullData[243^ /*line ukhZRzK9Q.go:1*/ int(idxKey[8])]-fullData[230^ /*line ukhZRzK9Q.go:1*/ int(idxKey[8])], fullData[246^ /*line ukhZRzK9Q.go:1*/ int(idxKey[9])]+fullData[250^ /*line ukhZRzK9Q.go:1*/ int(idxKey[9])], fullData[68^ /*line ukhZRzK9Q.go:1*/ int(idxKey[1])]-fullData[98^ /*line ukhZRzK9Q.go:1*/ int(idxKey[1])], fullData[143^ /*line ukhZRzK9Q.go:1*/ int(idxKey[3])]^fullData[140^ /*line ukhZRzK9Q.go:1*/ int(idxKey[3])], fullData[65^ /*line ukhZRzK9Q.go:1*/ int(idxKey[0])]-fullData[101^ /*line ukhZRzK9Q.go:1*/ int(idxKey[0])], fullData[113^ /*line ukhZRzK9Q.go:1*/ int(idxKey[10])]^fullData[125^ /*line ukhZRzK9Q.go:1*/ int(idxKey[10])], fullData[38^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])]^fullData[16^ /*line ukhZRzK9Q.go:1*/ int(idxKey[2])], fullData[51^ /*line ukhZRzK9Q.go:1*/ int(idxKey[7])]-fullData[1^ /*line ukhZRzK9Q.go:1*/ int(idxKey[7])], fullData[119^ /*line ukhZRzK9Q.go:1*/ int(idxKey[10])]+fullData[86^ /*line ukhZRzK9Q.go:1*/ int(idxKey[10])])
				return /*line ukhZRzK9Q.go:1*/ string(data)
			}())
		case *qUKQUCfTXB:
			lhQWH7m, nmFIUUSGjl = zB0wlXXn, jcDLabJ7o
			continue
		case *fullNode:
			lhQWH7m, nmFIUUSGjl = zB0wlXXn, jcDLabJ7o
			continue
		case hashNode:
			jcDLabJ7o, eZzE0KPS = /*line L0t58UqHWmNv.go:1*/ oGS_E_Q( /*line SD2F5vXbvLCV.go:1*/ common.BytesToHash(nPjnWCe))
			if eZzE0KPS != nil {
				return nil, nil, eZzE0KPS
			}
		case valueNode:
			r0QDFWcta7R = nPjnWCe
		}

		switch za7g4dFR3Y := nmFIUUSGjl.(type) {
		case *qUKQUCfTXB:
			za7g4dFR3Y.YpmEYGB = jcDLabJ7o
		case *fullNode:
			za7g4dFR3Y.Children[lhQWH7m[0]] = jcDLabJ7o
		default:
			/*line m8j39kTg.go:1*/ panic( /*line Bu2c6mH9HY.go:1*/ fmt.Sprintf(func() /*line ydFlGVb.go:1*/ string {
				fullData := [] /*line ydFlGVb.go:1*/ byte("l\x90\xbb\x86\x8c\xe3i\xae\x81\xeb\xbd\xcc\x00\xdd\x149V\xecb%\x00\xefi\xd8\xfef\x8e\xa7\x1cu\xd2\x05\xdd5\x88\x86w\xd9H\v")
				idxKey := [] /*line ydFlGVb.go:1*/ byte("\x993\x06U\xc3\xcad\xbc\xfa")
				data := /*line ydFlGVb.go:1*/ make([]byte, 0, 21)
				data = /*line ydFlGVb.go:1*/ append(data, fullData[217^ /*line ydFlGVb.go:1*/ int(idxKey[5])]^fullData[222^ /*line ydFlGVb.go:1*/ int(idxKey[5])], fullData[199^ /*line ydFlGVb.go:1*/ int(idxKey[5])]+fullData[238^ /*line ydFlGVb.go:1*/ int(idxKey[5])], fullData[116^ /*line ydFlGVb.go:1*/ int(idxKey[6])]-fullData[120^ /*line ydFlGVb.go:1*/ int(idxKey[6])], fullData[117^ /*line ydFlGVb.go:1*/ int(idxKey[6])]^fullData[111^ /*line ydFlGVb.go:1*/ int(idxKey[6])], fullData[64^ /*line ydFlGVb.go:1*/ int(idxKey[3])]^fullData[118^ /*line ydFlGVb.go:1*/ int(idxKey[3])], fullData[50^ /*line ydFlGVb.go:1*/ int(idxKey[1])]^fullData[43^ /*line ydFlGVb.go:1*/ int(idxKey[1])], fullData[209^ /*line ydFlGVb.go:1*/ int(idxKey[4])]^fullData[205^ /*line ydFlGVb.go:1*/ int(idxKey[4])], fullData[188^ /*line ydFlGVb.go:1*/ int(idxKey[0])]+fullData[187^ /*line ydFlGVb.go:1*/ int(idxKey[0])], fullData[149^ /*line ydFlGVb.go:1*/ int(idxKey[0])]+fullData[153^ /*line ydFlGVb.go:1*/ int(idxKey[0])], fullData[45^ /*line ydFlGVb.go:1*/ int(idxKey[1])]-fullData[53^ /*line ydFlGVb.go:1*/ int(idxKey[1])], fullData[67^ /*line ydFlGVb.go:1*/ int(idxKey[6])]-fullData[127^ /*line ydFlGVb.go:1*/ int(idxKey[6])], fullData[86^ /*line ydFlGVb.go:1*/ int(idxKey[3])]-fullData[76^ /*line ydFlGVb.go:1*/ int(idxKey[3])], fullData[184^ /*line ydFlGVb.go:1*/ int(idxKey[0])]+fullData[150^ /*line ydFlGVb.go:1*/ int(idxKey[0])], fullData[2^ /*line ydFlGVb.go:1*/ int(idxKey[2])]^fullData[3^ /*line ydFlGVb.go:1*/ int(idxKey[2])], fullData[16^ /*line ydFlGVb.go:1*/ int(idxKey[2])]-fullData[25^ /*line ydFlGVb.go:1*/ int(idxKey[2])], fullData[66^ /*line ydFlGVb.go:1*/ int(idxKey[3])]^fullData[95^ /*line ydFlGVb.go:1*/ int(idxKey[3])], fullData[93^ /*line ydFlGVb.go:1*/ int(idxKey[3])]^fullData[87^ /*line ydFlGVb.go:1*/ int(idxKey[3])], fullData[224^ /*line ydFlGVb.go:1*/ int(idxKey[8])]^fullData[253^ /*line ydFlGVb.go:1*/ int(idxKey[8])], fullData[220^ /*line ydFlGVb.go:1*/ int(idxKey[8])]+fullData[218^ /*line ydFlGVb.go:1*/ int(idxKey[8])], fullData[195^ /*line ydFlGVb.go:1*/ int(idxKey[5])]-fullData[215^ /*line ydFlGVb.go:1*/ int(idxKey[5])])
				return /*line ydFlGVb.go:1*/ string(data)
			}(), za7g4dFR3Y, za7g4dFR3Y))
		}
		if /*line I5hNBXaypu.go:1*/ len(r0QDFWcta7R) > 0 {
			return m7S8SE, r0QDFWcta7R, nil
		}
		lhQWH7m, nmFIUUSGjl = zB0wlXXn, jcDLabJ7o
	}
}

func jPTu6P(gnGMaXTuiFeE node, a1dpIPy []byte, h3Iy1brk_ []byte) (bool, error) {
	a1dpIPy, h3Iy1brk_ = /*line bvz28kGd.go:1*/ nd12pU880Z8(a1dpIPy) /*line YQaUU4Kg.go:1*/, nd12pU880Z8(h3Iy1brk_)

	var (
		wMJhBVQ    = 0
		nmFIUUSGjl node

		fLuju8, ndcZSQ5Esf int
	)
findFork:
	for {
		switch mjh9fYwPqBOf := (gnGMaXTuiFeE).(type) {
		case *qUKQUCfTXB:
			mjh9fYwPqBOf.d6dipJ = nodeFlag{dirty: true}

			if /*line BaL4ksa0KwkP.go:1*/ len(a1dpIPy)-wMJhBVQ < /*line Ji0aa5ifS.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A) {
				fLuju8 = /*line W8IvNa.go:1*/ bytes.Compare(a1dpIPy[wMJhBVQ:], mjh9fYwPqBOf.ANdZYqbzJ1A)
			} else {
				fLuju8 = /*line gUauHc26M3.go:1*/ bytes.Compare(a1dpIPy[wMJhBVQ:wMJhBVQ+ /*line WdwonIVLkXT.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A)], mjh9fYwPqBOf.ANdZYqbzJ1A)
			}
			if /*line cEVBUg.go:1*/ len(h3Iy1brk_)-wMJhBVQ < /*line AbovOM.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A) {
				ndcZSQ5Esf = /*line fOzA63D9.go:1*/ bytes.Compare(h3Iy1brk_[wMJhBVQ:], mjh9fYwPqBOf.ANdZYqbzJ1A)
			} else {
				ndcZSQ5Esf = /*line M3Bp1cvFbf.go:1*/ bytes.Compare(h3Iy1brk_[wMJhBVQ:wMJhBVQ+ /*line q60l6p8t7.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A)], mjh9fYwPqBOf.ANdZYqbzJ1A)
			}
			if fLuju8 != 0 || ndcZSQ5Esf != 0 {
				break findFork
			}
			nmFIUUSGjl = gnGMaXTuiFeE
			gnGMaXTuiFeE, wMJhBVQ = mjh9fYwPqBOf.YpmEYGB, wMJhBVQ+ /*line VvOpJ5K.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A)
		case *fullNode:
			mjh9fYwPqBOf.flags = nodeFlag{dirty: true}

			udeVyi8, jhWfXSWw := mjh9fYwPqBOf.Children[a1dpIPy[wMJhBVQ]], mjh9fYwPqBOf.Children[h3Iy1brk_[wMJhBVQ]]
			if udeVyi8 == nil || jhWfXSWw == nil || udeVyi8 != jhWfXSWw {
				break findFork
			}
			nmFIUUSGjl = gnGMaXTuiFeE
			gnGMaXTuiFeE, wMJhBVQ = mjh9fYwPqBOf.Children[a1dpIPy[wMJhBVQ]], wMJhBVQ+1
		default:
			/*line U6MaiD_jk.go:1*/ panic( /*line kuHFO5isXuFE.go:1*/ fmt.Sprintf(func() /*line Ih2_b3C0HL.go:1*/ string {
				key := [] /*line Ih2_b3C0HL.go:1*/ byte("i\xf4/yxC\xe7\xe9.\xda\x02\xccI\xce\xde\xebg\x18-r")
				data := [] /*line Ih2_b3C0HL.go:1*/ byte("\xbc`\v\xa7\xf1+\x8fx>\x8fbT%\xa1\x86z\xd3\b\xf8\x04")
				for i, b := range key {
					data[i] = data[i] + b
				}
				return /*line Ih2_b3C0HL.go:1*/ string(data)
			}(), gnGMaXTuiFeE, gnGMaXTuiFeE))
		}
	}
	switch mjh9fYwPqBOf := gnGMaXTuiFeE.(type) {
	case *qUKQUCfTXB:

		if fLuju8 == -1 && ndcZSQ5Esf == -1 {
			return false /*line Dkni35jx18.go:1*/, errors.New(func() /*line nKXPMYU.go:1*/ string {
				data := [] /*line nKXPMYU.go:1*/ byte("05G\x11y f\x15[`\\")
				positions := [...]byte{2, 3, 10, 10, 6, 10, 10, 7, 2, 9, 8, 10, 2, 0, 6, 1}
				for i := 0; i < 16; i += 2 {
					localKey := /*line nKXPMYU.go:1*/ byte(i) + /*line nKXPMYU.go:1*/ byte(positions[i]^positions[i+1]) + 50
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line nKXPMYU.go:1*/ string(data)
			}())
		}
		if fLuju8 == 1 && ndcZSQ5Esf == 1 {
			return false /*line nv78sxEi4L3.go:1*/, errors.New(func() /*line YCXw9hNGGYH.go:1*/ string {
				seed := /*line YCXw9hNGGYH.go:1*/ byte(135)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line YCXw9hNGGYH.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line YCXw9hNGGYH.go:1*/ fnc(226)(4)(29)(254)(241)(89)(160)(19)(235)(23)(226)
				return /*line YCXw9hNGGYH.go:1*/ string(data)
			}())
		}
		if fLuju8 != 0 && ndcZSQ5Esf != 0 {

			if nmFIUUSGjl == nil {
				return true, nil
			}
			nmFIUUSGjl.(*fullNode).Children[a1dpIPy[wMJhBVQ-1]] = nil
			return false, nil
		}

		if ndcZSQ5Esf != 0 {
			if _, yY_yPSd8vG := mjh9fYwPqBOf.YpmEYGB.(valueNode); yY_yPSd8vG {

				if nmFIUUSGjl == nil {
					return true, nil
				}
				nmFIUUSGjl.(*fullNode).Children[a1dpIPy[wMJhBVQ-1]] = nil
				return false, nil
			}
			return false /*line GdVf8k_oW2.go:1*/, fpzHzrdmASxu(mjh9fYwPqBOf, mjh9fYwPqBOf.YpmEYGB, a1dpIPy[wMJhBVQ:] /*line _nuc7Z.go:1*/, len(mjh9fYwPqBOf.ANdZYqbzJ1A), false)
		}
		if fLuju8 != 0 {
			if _, yY_yPSd8vG := mjh9fYwPqBOf.YpmEYGB.(valueNode); yY_yPSd8vG {

				if nmFIUUSGjl == nil {
					return true, nil
				}
				nmFIUUSGjl.(*fullNode).Children[h3Iy1brk_[wMJhBVQ-1]] = nil
				return false, nil
			}
			return false /*line SfQVRdzakNIQ.go:1*/, fpzHzrdmASxu(mjh9fYwPqBOf, mjh9fYwPqBOf.YpmEYGB, h3Iy1brk_[wMJhBVQ:] /*line SK1cwsHe7icE.go:1*/, len(mjh9fYwPqBOf.ANdZYqbzJ1A), true)
		}
		return false, nil
	case *fullNode:

		for q2u2020KD6 := a1dpIPy[wMJhBVQ] + 1; q2u2020KD6 < h3Iy1brk_[wMJhBVQ]; q2u2020KD6++ {
			mjh9fYwPqBOf.Children[q2u2020KD6] = nil
		}
		if eZzE0KPS := /*line aiVOco1VK.go:1*/ fpzHzrdmASxu(mjh9fYwPqBOf, mjh9fYwPqBOf.Children[a1dpIPy[wMJhBVQ]], a1dpIPy[wMJhBVQ:], 1, false); eZzE0KPS != nil {
			return false, eZzE0KPS
		}
		if eZzE0KPS := /*line BgvMlngd.go:1*/ fpzHzrdmASxu(mjh9fYwPqBOf, mjh9fYwPqBOf.Children[h3Iy1brk_[wMJhBVQ]], h3Iy1brk_[wMJhBVQ:], 1, true); eZzE0KPS != nil {
			return false, eZzE0KPS
		}
		return false, nil
	default:
		/*line c0DSxfagT.go:1*/ panic( /*line UwGR14.go:1*/ fmt.Sprintf(func() /*line Are4S2p0.go:1*/ string {
			data := /*line Are4S2p0.go:1*/ make([]byte, 0, 21)
			i := 3
			decryptKey := 121
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					for y := range data {
						data[y] = data[y] ^ /*line Are4S2p0.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 0:
					i = 10
					data = /*line Are4S2p0.go:1*/ append(data, 49)
				case 8:
					i = 4
					data = /*line Are4S2p0.go:1*/ append(data, 117)
				case 2:
					i = 0
					data = /*line Are4S2p0.go:1*/ append(data, 113)
				case 4:
					data = /*line Are4S2p0.go:1*/ append(data, "0yyq"...,
					)
					i = 2
				case 7:
					data = /*line Are4S2p0.go:1*/ append(data, ",~"...,
					)
					i = 1
				case 3:
					i = 5
					data = /*line Are4S2p0.go:1*/ append(data, ">N#8"...,
					)
				case 5:
					i = 9
					data = /*line Are4S2p0.go:1*/ append(data, "vpk"...,
					)
				case 9:
					i = 8
					data = /*line Are4S2p0.go:1*/ append(data, "}\x7f{"...,
					)
				case 10:
					i = 7
					data = /*line Are4S2p0.go:1*/ append(data, 42)
				}
			}
			return /*line Are4S2p0.go:1*/ string(data)
		}(), gnGMaXTuiFeE, gnGMaXTuiFeE))
	}
}

func fpzHzrdmASxu(nmFIUUSGjl node, jcDLabJ7o node, lhQWH7m []byte, wMJhBVQ int, taaVk9y bool) error {
	switch nPjnWCe := jcDLabJ7o.(type) {
	case *fullNode:
		if taaVk9y {
			for q2u2020KD6 := 0; q2u2020KD6 < /*line B4B2QrX7.go:1*/ int(lhQWH7m[wMJhBVQ]); q2u2020KD6++ {
				nPjnWCe.Children[q2u2020KD6] = nil
			}
			nPjnWCe.flags = nodeFlag{dirty: true}
		} else {
			for q2u2020KD6 := lhQWH7m[wMJhBVQ] + 1; q2u2020KD6 < 16; q2u2020KD6++ {
				nPjnWCe.Children[q2u2020KD6] = nil
			}
			nPjnWCe.flags = nodeFlag{dirty: true}
		}
		return /*line RJy1TspxGQ.go:1*/ fpzHzrdmASxu(nPjnWCe, nPjnWCe.Children[lhQWH7m[wMJhBVQ]], lhQWH7m, wMJhBVQ+1, taaVk9y)
	case *qUKQUCfTXB:
		if /*line SBWx3Aii2G.go:1*/ len(lhQWH7m[wMJhBVQ:]) < /*line efBoEs.go:1*/ len(nPjnWCe.ANdZYqbzJ1A) || ! /*line H5MpTaQk.go:1*/ bytes.Equal(nPjnWCe.ANdZYqbzJ1A, lhQWH7m[wMJhBVQ:wMJhBVQ+ /*line AACnG5WVMxId.go:1*/ len(nPjnWCe.ANdZYqbzJ1A)]) {

			if taaVk9y {
				if /*line dWgacBf.go:1*/ bytes.Compare(nPjnWCe.ANdZYqbzJ1A, lhQWH7m[wMJhBVQ:]) < 0 {

					qC6kc1PxI8 := nmFIUUSGjl.(*fullNode)
					qC6kc1PxI8.Children[lhQWH7m[wMJhBVQ-1]] = nil
				}

			} else {
				if /*line LtBgK1E.go:1*/ bytes.Compare(nPjnWCe.ANdZYqbzJ1A, lhQWH7m[wMJhBVQ:]) > 0 {

					qC6kc1PxI8 := nmFIUUSGjl.(*fullNode)
					qC6kc1PxI8.Children[lhQWH7m[wMJhBVQ-1]] = nil
				}

			}
			return nil
		}
		if _, yY_yPSd8vG := nPjnWCe.YpmEYGB.(valueNode); yY_yPSd8vG {
			qC6kc1PxI8 := nmFIUUSGjl.(*fullNode)
			qC6kc1PxI8.Children[lhQWH7m[wMJhBVQ-1]] = nil
			return nil
		}
		nPjnWCe.d6dipJ = nodeFlag{dirty: true}
		return /*line Cp6yAxn.go:1*/ fpzHzrdmASxu(nPjnWCe, nPjnWCe.YpmEYGB, lhQWH7m, wMJhBVQ+ /*line I6WwdN_OO.go:1*/ len(nPjnWCe.ANdZYqbzJ1A), taaVk9y)
	case nil:

		return nil
	default:
		/*line Emk4Uujx3ua.go:1*/ panic(func() /*line wNZAX5Y.go:1*/ string {
			data := [] /*line wNZAX5Y.go:1*/ byte("ty ekntl\x7f\x94!\x88\x1ah\x93ppen")
			positions := [...]byte{4, 3, 8, 10, 12, 12, 3, 6, 8, 10, 0, 5, 9, 1, 6, 14, 11, 6, 1, 8, 8, 9}
			for i := 0; i < 22; i += 2 {
				localKey := /*line wNZAX5Y.go:1*/ byte(i) + /*line wNZAX5Y.go:1*/ byte(positions[i]^positions[i+1]) + 246
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line wNZAX5Y.go:1*/ string(data)
		}())
	}
}

func c5hvUl1(uAjz8NxU_cL node, lhQWH7m []byte) bool {
	wMJhBVQ, lhQWH7m := 0 /*line cITLxJdlBq.go:1*/, nd12pU880Z8(lhQWH7m)
	for uAjz8NxU_cL != nil {
		switch mjh9fYwPqBOf := uAjz8NxU_cL.(type) {
		case *fullNode:
			for q2u2020KD6 := lhQWH7m[wMJhBVQ] + 1; q2u2020KD6 < 16; q2u2020KD6++ {
				if mjh9fYwPqBOf.Children[q2u2020KD6] != nil {
					return true
				}
			}
			uAjz8NxU_cL, wMJhBVQ = mjh9fYwPqBOf.Children[lhQWH7m[wMJhBVQ]], wMJhBVQ+1
		case *qUKQUCfTXB:
			if /*line CrOCoym.go:1*/ len(lhQWH7m)-wMJhBVQ < /*line UtHx97H.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A) || ! /*line uihl8tJGa.go:1*/ bytes.Equal(mjh9fYwPqBOf.ANdZYqbzJ1A, lhQWH7m[wMJhBVQ:wMJhBVQ+ /*line IsxjLaC.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A)]) {
				return /*line lWXuMx_22.go:1*/ bytes.Compare(mjh9fYwPqBOf.ANdZYqbzJ1A, lhQWH7m[wMJhBVQ:]) > 0
			}
			uAjz8NxU_cL, wMJhBVQ = mjh9fYwPqBOf.YpmEYGB, wMJhBVQ+ /*line IrNwwgg.go:1*/ len(mjh9fYwPqBOf.ANdZYqbzJ1A)
		case valueNode:
			return false
		default:
			/*line HxWLsBaAJMyY.go:1*/ panic( /*line V3FCo7_.go:1*/ fmt.Sprintf(func() /*line km_SEk3m.go:1*/ string {
				data := [] /*line km_SEk3m.go:1*/ byte("#;\xed \x99\x87\xa4alid \xfe)d\xc6@\x9aŮ")
				positions := [...]byte{2, 17, 5, 0, 6, 2, 19, 17, 2, 12, 4, 6, 12, 13, 15, 16, 6, 1, 0, 18, 17, 0, 18, 16, 1, 1}
				for i := 0; i < 26; i += 2 {
					localKey := /*line km_SEk3m.go:1*/ byte(i) + /*line km_SEk3m.go:1*/ byte(positions[i]^positions[i+1]) + 174
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line km_SEk3m.go:1*/ string(data)
			}(), uAjz8NxU_cL, uAjz8NxU_cL))
		}
	}
	return false
}

func SPXGaGX(z8PGQUy common.Hash, tD4mxkwvh []byte, ywhO7DKTqn [][]byte, _Et7RxouL [][]byte, b95_zCscTV ethdb.KeyValueReader) (bool, error) {
	if /*line L_7LqDtjn.go:1*/ len(ywhO7DKTqn) != /*line kbsNtEK.go:1*/ len(_Et7RxouL) {
		return false /*line WEylrmq.go:1*/, fmt.Errorf(func() /*line wowOd0FfwcV6.go:1*/ string {
			key := [] /*line wowOd0FfwcV6.go:1*/ byte("\x1f\xe6J\xaah\xd5>\f䧔ԘǬw\xb6C\xe1!\x17Ĺ\x9e\xf9a\n\xc0\xe8/Т9p)\x14LF\xe3\x8f\xef\xf4\xa8\x18z")
			data := [] /*line wowOd0FfwcV6.go:1*/ byte("\x88T\xad\x19\xd6H\xa7\x7fX\f\x02H\xb87\x1e\xe6%\xa9\x01\x85x8\x1a\xca\x19\xcco9[i\xf0ǝ\x9cI\x8a\xad\xb2X\xf4b.\xc8=\xde")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line wowOd0FfwcV6.go:1*/ string(data)
		}(), /*line ABs0qvLLac.go:1*/ len(ywhO7DKTqn) /*line gsu9h8w1.go:1*/, len(_Et7RxouL))
	}

	for q2u2020KD6 := 0; q2u2020KD6 < /*line YYuev4byT4.go:1*/ len(ywhO7DKTqn)-1; q2u2020KD6++ {
		if /*line AAJ_oTO7J.go:1*/ bytes.Compare(ywhO7DKTqn[q2u2020KD6], ywhO7DKTqn[q2u2020KD6+1]) >= 0 {
			return false /*line _a9aSzmXn2B.go:1*/, errors.New(func() /*line o4By5u.go:1*/ string {
				seed := /*line o4By5u.go:1*/ byte(29)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line o4By5u.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line o4By5u.go:1*/ fnc(143)(13)(39)(71)(140)(211)(239)(232)(125)(72)(145)(39)(250)(65)(132)(7)(15)(35)(65)(129)(253)(244)(230)(215)(174)(105)(121)(59)(123)(235)(229)(189)(118)(254)(242)(233)(203)
				return /*line o4By5u.go:1*/ string(data)
			}())
		}
	}
	for _, ar76Sw := range _Et7RxouL {
		if /*line GAs7PzWuifNf.go:1*/ len(ar76Sw) == 0 {
			return false /*line Ha4kdiA9t.go:1*/, errors.New(func() /*line PSuUPUm5O.go:1*/ string {
				seed := /*line PSuUPUm5O.go:1*/ byte(198)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line PSuUPUm5O.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line PSuUPUm5O.go:1*/ fnc(180)(27)(251)(247)(226)(73)(209)(236)(1)(4)(21)(224)(7)(3)(83)(162)(13)(25)(235)(13)(239)(26)(225)
				return /*line PSuUPUm5O.go:1*/ string(data)
			}())
		}
	}

	if b95_zCscTV == nil {
		rjqmOX5MHdh := /*line CYQXbF4h.go:1*/ KYaYz2rvh7Rt(nil)
		for sCuaQRs1, lhQWH7m := range ywhO7DKTqn {
			/*line o75ll1V.go:1*/ rjqmOX5MHdh.Update(lhQWH7m, _Et7RxouL[sCuaQRs1])
		}
		if ohpfKJiElL, ddolro := /*line KvCWRBt8dvZ.go:1*/ rjqmOX5MHdh.Hash(), z8PGQUy; ohpfKJiElL != ddolro {
			return false /*line FE1GTlc3Nj.go:1*/, fmt.Errorf(func() /*line Fl2rJEn_.go:1*/ string {
				data := [] /*line Fl2rJEn_.go:1*/ byte("x\vv\xd1l\x86\x7f Wr\x89\x85Ԛ\xa0w\x81n\"r\xde\u07b8\xa41Y\xb4];`o\x9e \x99H")
				positions := [...]byte{8, 27, 1, 27, 22, 3, 14, 31, 27, 25, 26, 23, 22, 21, 16, 11, 12, 19, 34, 5, 20, 13, 22, 12, 14, 12, 5, 5, 33, 28, 3, 29, 18, 10, 0, 6, 27, 21, 28, 8, 10, 24, 12, 6, 16, 12, 0, 10, 31, 8, 28, 0}
				for i := 0; i < 52; i += 2 {
					localKey := /*line Fl2rJEn_.go:1*/ byte(i) + /*line Fl2rJEn_.go:1*/ byte(positions[i]^positions[i+1]) + 197
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line Fl2rJEn_.go:1*/ string(data)
			}(), ddolro, ohpfKJiElL)
		}
		return false, nil
	}

	if /*line CAIc9CQa.go:1*/ len(ywhO7DKTqn) == 0 {
		m7S8SE, h_pxYek4zT, eZzE0KPS := /*line i53SI_.go:1*/ hClFfPa(z8PGQUy, nil, tD4mxkwvh, b95_zCscTV, true)
		if eZzE0KPS != nil {
			return false, eZzE0KPS
		}
		if h_pxYek4zT != nil || /*line FPEDQAtlgxGV.go:1*/ c5hvUl1(m7S8SE, tD4mxkwvh) {
			return false /*line vxR3BMS30_x.go:1*/, errors.New(func() /*line DRDqW4J6.go:1*/ string {
				data := [] /*line DRDqW4J6.go:1*/ byte("Eo\xd8\x1a γ\xad\x9b\xd3e\xecfa\xd9ᣴabl\x9d")
				positions := [...]byte{0, 12, 21, 0, 3, 7, 8, 14, 5, 3, 8, 9, 21, 7, 9, 9, 15, 11, 16, 3, 2, 7, 7, 2, 21, 11, 8, 6, 15, 17}
				for i := 0; i < 30; i += 2 {
					localKey := /*line DRDqW4J6.go:1*/ byte(i) + /*line DRDqW4J6.go:1*/ byte(positions[i]^positions[i+1]) + 25
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line DRDqW4J6.go:1*/ string(data)
			}())
		}
		return false, nil
	}
	var gZix1l_Ra = ywhO7DKTqn[ /*line BCMLW9TD4c3.go:1*/ len(ywhO7DKTqn)-1]

	if /*line lJp93KOvaE.go:1*/ len(ywhO7DKTqn) == 1 && /*line gA1fgU0SBFPr.go:1*/ bytes.Equal(tD4mxkwvh, gZix1l_Ra) {
		m7S8SE, h_pxYek4zT, eZzE0KPS := /*line cAUm2mu7.go:1*/ hClFfPa(z8PGQUy, nil, tD4mxkwvh, b95_zCscTV, false)
		if eZzE0KPS != nil {
			return false, eZzE0KPS
		}
		if ! /*line WMSfPdI_01w.go:1*/ bytes.Equal(tD4mxkwvh, ywhO7DKTqn[0]) {
			return false /*line qVpY6MAium9.go:1*/, errors.New(func() /*line aAHCrGGj.go:1*/ string {
				key := [] /*line aAHCrGGj.go:1*/ byte("\xe3\x99!q\x80?\x8f\x9d\xca\vSK\x8e\a\xa2A\x18\f\xa38\xb9cU\t\x88\xad\x9a\xcb\\")
				data := [] /*line aAHCrGGj.go:1*/ byte("F\b\x93\xe3\xe5\xa2\x03\xbd:}º\xf4'\x04\xb6\x8c,\f\xa6/\xc4\xc1r\xec\xcd\x050\xd5")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return /*line aAHCrGGj.go:1*/ string(data)
			}())
		}
		if ! /*line m_SuvnY5d4Wu.go:1*/ bytes.Equal(h_pxYek4zT, _Et7RxouL[0]) {
			return false /*line oLRicG4C3.go:1*/, errors.New(func() /*line BzovPOMK.go:1*/ string {
				seed := /*line BzovPOMK.go:1*/ byte(12)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line BzovPOMK.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line BzovPOMK.go:1*/ fnc(87)(12)(3)(0)(243)(254)(17)(172)(80)(2)(253)(0)(247)(186)(66)(19)(255)(172)(73)(5)(8)(235)(11)(253)(251)(188)(68)(253)(19)(237)
				return /*line BzovPOMK.go:1*/ string(data)
			}())
		}
		return /*line DgNNlvwW.go:1*/ c5hvUl1(m7S8SE, tD4mxkwvh), nil
	}

	if /*line lKyqrEFz.go:1*/ bytes.Compare(tD4mxkwvh, gZix1l_Ra) >= 0 {
		return false /*line ItDrTAyqygS.go:1*/, errors.New(func() /*line xvO6rYe5.go:1*/ string {
			key := [] /*line xvO6rYe5.go:1*/ byte("\x8fUתI1\x00UȆN\xbeْ/F\x0e")
			data := [] /*line xvO6rYe5.go:1*/ byte("\xf8\xc3M\v\xb5\x9adu-\xea\xb5#\xf9\xfd\x94\xbf\x81")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line xvO6rYe5.go:1*/ string(data)
		}())
	}

	if /*line YlZulKMka7.go:1*/ len(tD4mxkwvh) != /*line R8ChpD7AJN.go:1*/ len(gZix1l_Ra) {
		return false /*line HjC65aHW.go:1*/, errors.New(func() /*line awVc89.go:1*/ string {
			data := [] /*line awVc89.go:1*/ byte("ih3jp\x9a\xab\x97ts\x82\x8c esg\x8av~{\x8am")
			positions := [...]byte{5, 3, 1, 16, 5, 9, 14, 1, 20, 7, 7, 4, 6, 9, 18, 3, 2, 17, 19, 21, 5, 14, 4, 10, 11, 18, 20, 9}
			for i := 0; i < 28; i += 2 {
				localKey := /*line awVc89.go:1*/ byte(i) + /*line awVc89.go:1*/ byte(positions[i]^positions[i+1]) + 240
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line awVc89.go:1*/ string(data)
		}())
	}

	m7S8SE, _, eZzE0KPS := /*line DQaDprZIA.go:1*/ hClFfPa(z8PGQUy, nil, tD4mxkwvh, b95_zCscTV, true)
	if eZzE0KPS != nil {
		return false, eZzE0KPS
	}

	m7S8SE, _, eZzE0KPS = /*line apzUQ8yaMUHt.go:1*/ hClFfPa(z8PGQUy, m7S8SE, gZix1l_Ra, b95_zCscTV, true)
	if eZzE0KPS != nil {
		return false, eZzE0KPS
	}

	b2gMag, eZzE0KPS := /*line Obum296dZ.go:1*/ jPTu6P(m7S8SE, tD4mxkwvh, gZix1l_Ra)
	if eZzE0KPS != nil {
		return false, eZzE0KPS
	}

	rjqmOX5MHdh := &Trie{root: m7S8SE, reader: /*line VExjelLpt43a.go:1*/ uAe_TJrsD(), tracer: /*line uIn8TKBaz.go:1*/ aLjRi9a0()}
	if b2gMag {
		rjqmOX5MHdh.root = nil
	}
	for sCuaQRs1, lhQWH7m := range ywhO7DKTqn {
		/*line vFfUGayb.go:1*/ rjqmOX5MHdh.Update(lhQWH7m, _Et7RxouL[sCuaQRs1])
	}
	if /*line ie8bix9.go:1*/ rjqmOX5MHdh.Hash() != z8PGQUy {
		return false /*line RS6tc3e.go:1*/, fmt.Errorf(func() /*line EnuBO5.go:1*/ string {
			fullData := [] /*line EnuBO5.go:1*/ byte("6\xba\xbe6\xc0Xd\"\\P\x80\xe8\xf5Fn\xe7\xf8\xfa~\xab\x149\xb4\xc9Lf\x15\xfa\xfc\xfeѪGju+9\x7f2\xaa;\xe3/\xbcW:\x8eJh9T\xccg.\x01ݏ`*q\x9cE0p\x8f\xab\x19\xa7\x9d\t")
			idxKey := [] /*line EnuBO5.go:1*/ byte("Y\xb3I\x02J\x8b\xf6/ST\xca!ć\xab\xcf")
			data := /*line EnuBO5.go:1*/ make([]byte, 0, 36)
			data = /*line EnuBO5.go:1*/ append(data, fullData[111^ /*line EnuBO5.go:1*/ int(idxKey[0])]+fullData[105^ /*line EnuBO5.go:1*/ int(idxKey[0])], fullData[243^ /*line EnuBO5.go:1*/ int(idxKey[6])]^fullData[246^ /*line EnuBO5.go:1*/ int(idxKey[6])], fullData[244^ /*line EnuBO5.go:1*/ int(idxKey[10])]-fullData[203^ /*line EnuBO5.go:1*/ int(idxKey[10])], fullData[236^ /*line EnuBO5.go:1*/ int(idxKey[10])]+fullData[224^ /*line EnuBO5.go:1*/ int(idxKey[10])], fullData[175^ /*line EnuBO5.go:1*/ int(idxKey[1])]+fullData[140^ /*line EnuBO5.go:1*/ int(idxKey[1])], fullData[35^ /*line EnuBO5.go:1*/ int(idxKey[11])]+fullData[50^ /*line EnuBO5.go:1*/ int(idxKey[11])], fullData[96^ /*line EnuBO5.go:1*/ int(idxKey[2])]-fullData[108^ /*line EnuBO5.go:1*/ int(idxKey[2])], fullData[109^ /*line EnuBO5.go:1*/ int(idxKey[2])]+fullData[70^ /*line EnuBO5.go:1*/ int(idxKey[2])], fullData[128^ /*line EnuBO5.go:1*/ int(idxKey[1])]-fullData[187^ /*line EnuBO5.go:1*/ int(idxKey[1])], fullData[142^ /*line EnuBO5.go:1*/ int(idxKey[13])]^fullData[128^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[222^ /*line EnuBO5.go:1*/ int(idxKey[15])]+fullData[237^ /*line EnuBO5.go:1*/ int(idxKey[15])], fullData[27^ /*line EnuBO5.go:1*/ int(idxKey[11])]^fullData[28^ /*line EnuBO5.go:1*/ int(idxKey[11])], fullData[195^ /*line EnuBO5.go:1*/ int(idxKey[15])]-fullData[143^ /*line EnuBO5.go:1*/ int(idxKey[15])], fullData[139^ /*line EnuBO5.go:1*/ int(idxKey[15])]-fullData[244^ /*line EnuBO5.go:1*/ int(idxKey[15])], fullData[251^ /*line EnuBO5.go:1*/ int(idxKey[15])]^fullData[239^ /*line EnuBO5.go:1*/ int(idxKey[15])], fullData[109^ /*line EnuBO5.go:1*/ int(idxKey[4])]^fullData[125^ /*line EnuBO5.go:1*/ int(idxKey[4])], fullData[133^ /*line EnuBO5.go:1*/ int(idxKey[12])]-fullData[235^ /*line EnuBO5.go:1*/ int(idxKey[12])], fullData[123^ /*line EnuBO5.go:1*/ int(idxKey[2])]^fullData[100^ /*line EnuBO5.go:1*/ int(idxKey[2])], fullData[7^ /*line EnuBO5.go:1*/ int(idxKey[7])]+fullData[30^ /*line EnuBO5.go:1*/ int(idxKey[7])], fullData[39^ /*line EnuBO5.go:1*/ int(idxKey[11])]+fullData[10^ /*line EnuBO5.go:1*/ int(idxKey[11])], fullData[178^ /*line EnuBO5.go:1*/ int(idxKey[5])]-fullData[155^ /*line EnuBO5.go:1*/ int(idxKey[5])], fullData[157^ /*line EnuBO5.go:1*/ int(idxKey[13])]+fullData[159^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[12^ /*line EnuBO5.go:1*/ int(idxKey[2])]+fullData[104^ /*line EnuBO5.go:1*/ int(idxKey[2])], fullData[141^ /*line EnuBO5.go:1*/ int(idxKey[13])]^fullData[140^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[169^ /*line EnuBO5.go:1*/ int(idxKey[13])]-fullData[137^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[33^ /*line EnuBO5.go:1*/ int(idxKey[3])]+fullData[25^ /*line EnuBO5.go:1*/ int(idxKey[3])], fullData[63^ /*line EnuBO5.go:1*/ int(idxKey[11])]+fullData[98^ /*line EnuBO5.go:1*/ int(idxKey[11])], fullData[215^ /*line EnuBO5.go:1*/ int(idxKey[10])]+fullData[255^ /*line EnuBO5.go:1*/ int(idxKey[10])], fullData[188^ /*line EnuBO5.go:1*/ int(idxKey[14])]+fullData[135^ /*line EnuBO5.go:1*/ int(idxKey[14])], fullData[197^ /*line EnuBO5.go:1*/ int(idxKey[13])]^fullData[149^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[132^ /*line EnuBO5.go:1*/ int(idxKey[13])]+fullData[146^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[131^ /*line EnuBO5.go:1*/ int(idxKey[13])]^fullData[145^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[190^ /*line EnuBO5.go:1*/ int(idxKey[1])]^fullData[170^ /*line EnuBO5.go:1*/ int(idxKey[1])], fullData[191^ /*line EnuBO5.go:1*/ int(idxKey[13])]^fullData[152^ /*line EnuBO5.go:1*/ int(idxKey[13])], fullData[208^ /*line EnuBO5.go:1*/ int(idxKey[12])]-fullData[248^ /*line EnuBO5.go:1*/ int(idxKey[12])])
			return /*line EnuBO5.go:1*/ string(data)
		}(), z8PGQUy /*line cixbmxeJhh.go:1*/, rjqmOX5MHdh.Hash())
	}
	return /*line jw9Mfgn.go:1*/ c5hvUl1(rjqmOX5MHdh.root, ywhO7DKTqn[ /*line nFMEWJfv.go:1*/ len(ywhO7DKTqn)-1]), nil
}

func cga2ibNy(owwtC8asmI node, lhQWH7m []byte, ngoOG_we bool) ([]byte, node) {
	for {
		switch gnGMaXTuiFeE := owwtC8asmI.(type) {
		case *qUKQUCfTXB:
			if /*line F3zNz1v.go:1*/ len(lhQWH7m) < /*line y7QxW_P7ClC.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A) || ! /*line WW2hVfoQ.go:1*/ bytes.Equal(gnGMaXTuiFeE.ANdZYqbzJ1A, lhQWH7m[: /*line Sa1hLS.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A)]) {
				return nil, nil
			}
			owwtC8asmI = gnGMaXTuiFeE.YpmEYGB
			lhQWH7m = lhQWH7m[ /*line lz4RF_.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A):]
			if !ngoOG_we {
				return lhQWH7m, owwtC8asmI
			}
		case *fullNode:
			owwtC8asmI = gnGMaXTuiFeE.Children[lhQWH7m[0]]
			lhQWH7m = lhQWH7m[1:]
			if !ngoOG_we {
				return lhQWH7m, owwtC8asmI
			}
		case hashNode:
			return lhQWH7m, gnGMaXTuiFeE
		case nil:
			return lhQWH7m, nil
		case valueNode:
			return nil, gnGMaXTuiFeE
		default:
			/*line RwAkeILAi.go:1*/ panic( /*line RnwlzRh.go:1*/ fmt.Sprintf(func() /*line mk89RmYo.go:1*/ string {
				data := [] /*line mk89RmYo.go:1*/ byte("%T\xb1,in\xd6al\xec\xc0\xddX\rf\xcd:\xa5\xfb\x82")
				positions := [...]byte{19, 18, 10, 13, 12, 12, 10, 3, 12, 6, 6, 14, 15, 6, 17, 9, 19, 2, 11, 10, 3, 2, 14, 2}
				for i := 0; i < 24; i += 2 {
					localKey := /*line mk89RmYo.go:1*/ byte(i) + /*line mk89RmYo.go:1*/ byte(positions[i]^positions[i+1]) + 166
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line mk89RmYo.go:1*/ string(data)
			}(), owwtC8asmI, owwtC8asmI))
		}
	}
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append
var _ = common.AbsolutePath
var _ ethdb.AncientReader
var _ = log.Crit
