//line :1

package snapshot

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/triedb"
)

var (
	vA2HcXHs0F = /*line gKv88Nf.go:1*/ metrics.NewRegisteredMeter(func() /*line G2HxWyXl.go:1*/ string {
		seed := /*line G2HxWyXl.go:1*/ byte(163)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line G2HxWyXl.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line G2HxWyXl.go:1*/ fnc(208)(1)(237)(19)(241)(202)(68)(251)(243)(15)(3)(245)(7)(5)(187)(52)(9)(249)(252)(13)(193)(50)(2)(0)(12)(6)(249)(6)(187)(57)(1)(11)
		return /*line G2HxWyXl.go:1*/ string(data)
	}(), nil)
	xL3GIW = /*line uvhzcV.go:1*/ metrics.NewRegisteredMeter(func() /*line mkAZlYd.go:1*/ string {
		seed := /*line mkAZlYd.go:1*/ byte(84)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line mkAZlYd.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line mkAZlYd.go:1*/ fnc(31)(1)(237)(19)(241)(202)(68)(251)(243)(15)(3)(245)(7)(5)(187)(52)(9)(249)(252)(13)(193)(50)(2)(0)(12)(6)(249)(6)(187)(62)(252)(10)(0)
		return /*line mkAZlYd.go:1*/ string(data)
	}(), nil)
	msaMf9 = /*line Tv7FrMJETkWz.go:1*/ metrics.NewRegisteredMeter(func() /*line Mc03KyjOA.go:1*/ string {
		seed := /*line Mc03KyjOA.go:1*/ byte(162)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line Mc03KyjOA.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line Mc03KyjOA.go:1*/ fnc(209)(7)(27)(225)(19)(166)(92)(229)(17)(241)(1)(27)(225)(27)(165)(76)(23)(247)(232)(31)(191)(46)(30)(248)(252)(250)(231)(4)(91)(166)(27)(245)(253)
		return /*line Mc03KyjOA.go:1*/ string(data)
	}(), nil)
	nQLCSIOQS = /*line G0dCnM0_ih.go:1*/ metrics.NewRegisteredMeter(func() /*line I3A_AnC13Wk.go:1*/ string {
		data := /*line I3A_AnC13Wk.go:1*/ make([]byte, 0, 34)
		i := 6
		decryptKey := 149
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 9
				data = /*line I3A_AnC13Wk.go:1*/ append(data, 79)
			case 0:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, 51)
				i = 12
			case 16:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "2/"...,
				)
				i = 5
			case 14:
				i = 0
				data = /*line I3A_AnC13Wk.go:1*/ append(data, ":4;\xef"...,
				)
			case 12:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "'$"...,
				)
				i = 13
			case 5:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, 61)
				i = 15
			case 15:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "\xf7*"...,
				)
				i = 1
			case 8:
				i = 4
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "?E\x01"...,
				)
			case 13:
				i = 2
				data = /*line I3A_AnC13Wk.go:1*/ append(data, 96)
			case 7:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, ".3"...,
				)
				i = 14
			case 9:
				i = 10
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "Q?S"...,
				)
			case 2:
				for y := range data {
					data[y] = data[y] - /*line I3A_AnC13Wk.go:1*/ byte(decryptKey^y)
				}
				i = 3
			case 1:
				i = 7
				data = /*line I3A_AnC13Wk.go:1*/ append(data, 45)
			case 11:
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "5EI?"...,
				)
				i = 8
			case 17:
				i = 11
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "MI"...,
				)
			case 4:
				i = 16
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "68"...,
				)
			case 10:
				i = 17
				data = /*line I3A_AnC13Wk.go:1*/ append(data, "=\b"...,
				)
			}
		}
		return /*line I3A_AnC13Wk.go:1*/ string(data)
	}(), nil)
	t9NCla4Ghrl = /*line p_ubeey2.go:1*/ metrics.NewRegisteredMeter(func() /*line yWjXk0QM.go:1*/ string {
		data := /*line yWjXk0QM.go:1*/ make([]byte, 0, 35)
		i := 7
		decryptKey := 19
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 12
				data = /*line yWjXk0QM.go:1*/ append(data, ".6-q"...,
				)
			case 4:
				data = /*line yWjXk0QM.go:1*/ append(data, ";;!!"...,
				)
				i = 6
			case 9:
				i = 2
				data = /*line yWjXk0QM.go:1*/ append(data, "6765"...,
				)
			case 3:
				for y := range data {
					data[y] = data[y] ^ /*line yWjXk0QM.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 11:
				i = 4
				data = /*line yWjXk0QM.go:1*/ append(data, "++"...,
				)
			case 8:
				data = /*line yWjXk0QM.go:1*/ append(data, "5#h7"...,
				)
				i = 11
			case 1:
				data = /*line yWjXk0QM.go:1*/ append(data, "61?y"...,
				)
				i = 9
			case 0:
				i = 3
				data = /*line yWjXk0QM.go:1*/ append(data, "4\x16\x06"...,
				)
			case 12:
				i = 0
				data = /*line yWjXk0QM.go:1*/ append(data, "(."...,
				)
			case 10:
				i = 8
				data = /*line yWjXk0QM.go:1*/ append(data, 33)
			case 7:
				i = 10
				data = /*line yWjXk0QM.go:1*/ append(data, "17"...,
				)
			case 6:
				i = 1
				data = /*line yWjXk0QM.go:1*/ append(data, ";c.>"...,
				)
			}
		}
		return /*line yWjXk0QM.go:1*/ string(data)
	}(), nil)

	vG6wMf = /*line wvhcnd.go:1*/ metrics.NewRegisteredMeter(func() /*line Ik1Cwa.go:1*/ string {
		data := [] /*line Ik1Cwa.go:1*/ byte("s:af\xadp\x89na\xb1\x9dhotoBlekt)6\x8a'j\x8fjc\x9e\x84sT")
		positions := [...]byte{19, 14, 1, 24, 9, 6, 22, 25, 26, 5, 25, 10, 24, 4, 3, 14, 23, 5, 25, 24, 28, 29, 24, 18, 14, 21, 24, 1, 28, 4, 20, 27, 27, 20, 19, 3, 25, 6, 10, 15, 29, 21, 30, 26, 10, 31}
		for i := 0; i < 46; i += 2 {
			localKey := /*line Ik1Cwa.go:1*/ byte(i) + /*line Ik1Cwa.go:1*/ byte(positions[i]^positions[i+1]) + 230
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line Ik1Cwa.go:1*/ string(data)
	}(), nil)
	jzc5tPpFnz = /*line vrbJdJ.go:1*/ metrics.NewRegisteredMeter(func() /*line MABdRf.go:1*/ string {
		key := [] /*line MABdRf.go:1*/ byte("t\xd1\xc8O\xfa~\t\u070e[\n\xbe\x8c \xec\xc4\rם\xd7\a2Ϊڸ\xab\x8a\xc1V\xe6Q\x02")
		data := [] /*line MABdRf.go:1*/ byte("\a\xa5\xa9;\x9fQz\xb2\xef+y\xd6\xe3Tça\xb2\xfc\xb9(A\xbaŨ\xd9\xcc\xef\xee;\x8f\"q")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line MABdRf.go:1*/ string(data)
	}(), nil)
	hqqacN = /*line j70FcY.go:1*/ metrics.NewRegisteredMeter(func() /*line qkyiZnPzLYH.go:1*/ string {
		key := [] /*line qkyiZnPzLYH.go:1*/ byte("f\xea\xa4Y/:\x90\x8foN\awW\xac+6|\x88\xaa-GZ]\xd5\x06\x03\x8a\f\x94\xcas\xd2\xd5")
		data := [] /*line qkyiZnPzLYH.go:1*/ byte("\xd9^\x05͔i\x03\xfdоz\xdf\xc6 Z\x99\xe8\xed\v\x9bv\xcd\xd1Dxd\xf1q\xc33\xe17M")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line qkyiZnPzLYH.go:1*/ string(data)
	}(), nil)
	jJdpjkAbM4 = /*line _DyKJJA.go:1*/ metrics.NewRegisteredMeter(func() /*line T5jeMv7jbK.go:1*/ string {
		fullData := [] /*line T5jeMv7jbK.go:1*/ byte("\x939\x15\x98\xe4\r\xe1\xb0}I\xbf\x87\xb16\x8c\x18x\x06\xf0d\xec\x8aթ\t\a\xcbNd\xaaqK&\x88\x107\xe4P\x84\xf0h(_\xe2\b\x92\f\x04\xcf'ׄ\xc6ߚ\x9e3\xa3\x17̰f\xa0\xadt:")
		idxKey := [] /*line T5jeMv7jbK.go:1*/ byte("þ=\x02\xc9=\xbe\x9en\xe0-S\xec\x90\xec\x1f")
		data := /*line T5jeMv7jbK.go:1*/ make([]byte, 0, 34)
		data = /*line T5jeMv7jbK.go:1*/ append(data, fullData[78^ /*line T5jeMv7jbK.go:1*/ int(idxKey[11])]-fullData[112^ /*line T5jeMv7jbK.go:1*/ int(idxKey[11])], fullData[0^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])]-fullData[45^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])], fullData[110^ /*line T5jeMv7jbK.go:1*/ int(idxKey[11])]^fullData[74^ /*line T5jeMv7jbK.go:1*/ int(idxKey[11])], fullData[12^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])]-fullData[56^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])], fullData[28^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])]-fullData[39^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])], fullData[37^ /*line T5jeMv7jbK.go:1*/ int(idxKey[10])]-fullData[54^ /*line T5jeMv7jbK.go:1*/ int(idxKey[10])], fullData[222^ /*line T5jeMv7jbK.go:1*/ int(idxKey[4])]-fullData[196^ /*line T5jeMv7jbK.go:1*/ int(idxKey[4])], fullData[249^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])]^fullData[232^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])], fullData[76^ /*line T5jeMv7jbK.go:1*/ int(idxKey[8])]^fullData[112^ /*line T5jeMv7jbK.go:1*/ int(idxKey[8])], fullData[28^ /*line T5jeMv7jbK.go:1*/ int(idxKey[2])]-fullData[50^ /*line T5jeMv7jbK.go:1*/ int(idxKey[2])], fullData[216^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])]+fullData[211^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])], fullData[250^ /*line T5jeMv7jbK.go:1*/ int(idxKey[4])]^fullData[221^ /*line T5jeMv7jbK.go:1*/ int(idxKey[4])], fullData[8^ /*line T5jeMv7jbK.go:1*/ int(idxKey[3])]-fullData[39^ /*line T5jeMv7jbK.go:1*/ int(idxKey[3])], fullData[229^ /*line T5jeMv7jbK.go:1*/ int(idxKey[0])]^fullData[209^ /*line T5jeMv7jbK.go:1*/ int(idxKey[0])], fullData[233^ /*line T5jeMv7jbK.go:1*/ int(idxKey[4])]+fullData[209^ /*line T5jeMv7jbK.go:1*/ int(idxKey[4])], fullData[131^ /*line T5jeMv7jbK.go:1*/ int(idxKey[0])]^fullData[249^ /*line T5jeMv7jbK.go:1*/ int(idxKey[0])], fullData[36^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])]^fullData[33^ /*line T5jeMv7jbK.go:1*/ int(idxKey[15])], fullData[235^ /*line T5jeMv7jbK.go:1*/ int(idxKey[9])]^fullData[203^ /*line T5jeMv7jbK.go:1*/ int(idxKey[9])], fullData[142^ /*line T5jeMv7jbK.go:1*/ int(idxKey[1])]+fullData[147^ /*line T5jeMv7jbK.go:1*/ int(idxKey[1])], fullData[253^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])]+fullData[196^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])], fullData[246^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])]+fullData[240^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])], fullData[86^ /*line T5jeMv7jbK.go:1*/ int(idxKey[11])]-fullData[101^ /*line T5jeMv7jbK.go:1*/ int(idxKey[11])], fullData[252^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])]^fullData[194^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])], fullData[217^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])]^fullData[235^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])], fullData[146^ /*line T5jeMv7jbK.go:1*/ int(idxKey[13])]-fullData[169^ /*line T5jeMv7jbK.go:1*/ int(idxKey[13])], fullData[82^ /*line T5jeMv7jbK.go:1*/ int(idxKey[8])]+fullData[98^ /*line T5jeMv7jbK.go:1*/ int(idxKey[8])], fullData[198^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])]+fullData[192^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])], fullData[226^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])]-fullData[221^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])], fullData[45^ /*line T5jeMv7jbK.go:1*/ int(idxKey[3])]-fullData[20^ /*line T5jeMv7jbK.go:1*/ int(idxKey[3])], fullData[236^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])]^fullData[234^ /*line T5jeMv7jbK.go:1*/ int(idxKey[12])], fullData[151^ /*line T5jeMv7jbK.go:1*/ int(idxKey[7])]-fullData[186^ /*line T5jeMv7jbK.go:1*/ int(idxKey[7])], fullData[237^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])]+fullData[197^ /*line T5jeMv7jbK.go:1*/ int(idxKey[14])], fullData[53^ /*line T5jeMv7jbK.go:1*/ int(idxKey[3])]-fullData[67^ /*line T5jeMv7jbK.go:1*/ int(idxKey[3])])
		return /*line T5jeMv7jbK.go:1*/ string(data)
	}(), nil)
	wZpZiLVo4M1 = /*line ZaUgUibANM.go:1*/ metrics.NewRegisteredMeter(func() /*line dCv7jJ0Ly.go:1*/ string {
		fullData := [] /*line dCv7jJ0Ly.go:1*/ byte("\x93\x9c\xe1\x8et\x1e\xa1\x00\x88Y\x19\x9b\x93\xba҉R?ݺÂ~组C0;ˠ-\x9c-\xb3\xe6\x1f\xad\xe6?\vi\xbdL\x83V9ȯ\xc37\xa6>d7\xecX\x80NF\a\xd7\xd4\x0f\xcam\x97\xf4")
		idxKey := [] /*line dCv7jJ0Ly.go:1*/ byte("\xd6d\x0e\t\xa6\x85\x80\x88\x1d\xe7\xafb\x96*")
		data := /*line dCv7jJ0Ly.go:1*/ make([]byte, 0, 35)
		data = /*line dCv7jJ0Ly.go:1*/ append(data, fullData[185^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[4])]-fullData[171^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[4])], fullData[11^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[13])]^fullData[35^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[13])], fullData[43^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])]-fullData[25^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])], fullData[37^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])]-fullData[54^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])], fullData[237^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])]+fullData[242^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])], fullData[175^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])]-fullData[154^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])], fullData[37^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[2])]^fullData[41^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[2])], fullData[188^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[6])]^fullData[169^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[6])], fullData[176^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[7])]^fullData[166^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[7])], fullData[50^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[13])]^fullData[55^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[13])], fullData[108^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[11])]^fullData[100^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[11])], fullData[66^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])]^fullData[103^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])], fullData[146^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[7])]-fullData[182^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[7])], fullData[76^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])]-fullData[38^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])], fullData[150^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])]+fullData[159^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])], fullData[173^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[4])]+fullData[137^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[4])], fullData[149^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])]+fullData[170^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])], fullData[156^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[5])]+fullData[135^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[5])], fullData[251^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])]^fullData[228^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])], fullData[6^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[8])]+fullData[41^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[8])], fullData[145^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[4])]^fullData[151^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[4])], fullData[52^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])]+fullData[8^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])], fullData[14^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])]+fullData[13^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[3])], fullData[140^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])]+fullData[160^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[10])], fullData[202^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])]+fullData[224^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])], fullData[16^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[2])]-fullData[31^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[2])], fullData[243^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])]+fullData[197^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[0])], fullData[108^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])]+fullData[118^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])], fullData[68^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])]-fullData[37^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[1])], fullData[205^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[9])]^fullData[167^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[9])], fullData[105^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[13])]-fullData[63^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[13])], fullData[243^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[9])]+fullData[212^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[9])], fullData[235^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[9])]^fullData[240^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[9])], fullData[116^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[11])]-fullData[104^ /*line dCv7jJ0Ly.go:1*/ int(idxKey[11])])
		return /*line dCv7jJ0Ly.go:1*/ string(data)
	}(), nil)

	bF2XXz = /*line VYfmBoargn.go:1*/ metrics.NewRegisteredMeter(func() /*line tERcLDbOU.go:1*/ string {
		data := /*line tERcLDbOU.go:1*/ make([]byte, 0, 33)
		i := 14
		decryptKey := 198
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 13
				data = /*line tERcLDbOU.go:1*/ append(data, "\x15\x0eR\x14"...,
				)
			case 4:
				data = /*line tERcLDbOU.go:1*/ append(data, "\a\vZ\x15"...,
				)
				i = 0
			case 1:
				for y := range data {
					data[y] = data[y] ^ /*line tERcLDbOU.go:1*/ byte(decryptKey^y)
				}
				i = 2
			case 16:
				i = 11
				data = /*line tERcLDbOU.go:1*/ append(data, "\x18\x02"...,
				)
			case 14:
				i = 17
				data = /*line tERcLDbOU.go:1*/ append(data, "\x12\x14"...,
				)
			case 13:
				data = /*line tERcLDbOU.go:1*/ append(data, 22)
				i = 9
			case 15:
				data = /*line tERcLDbOU.go:1*/ append(data, 2)
				i = 4
			case 7:
				data = /*line tERcLDbOU.go:1*/ append(data, 8)
				i = 3
			case 9:
				data = /*line tERcLDbOU.go:1*/ append(data, 10)
				i = 1
			case 8:
				i = 12
				data = /*line tERcLDbOU.go:1*/ append(data, 0)
			case 0:
				i = 6
				data = /*line tERcLDbOU.go:1*/ append(data, "\x14\x15\x16\r"...,
				)
			case 17:
				data = /*line tERcLDbOU.go:1*/ append(data, "\x02\x16"...,
				)
				i = 8
			case 10:
				i = 15
				data = /*line tERcLDbOU.go:1*/ append(data, "@\n\x18"...,
				)
			case 12:
				data = /*line tERcLDbOU.go:1*/ append(data, "K\x14"...,
				)
				i = 5
			case 5:
				data = /*line tERcLDbOU.go:1*/ append(data, 8)
				i = 7
			case 11:
				i = 10
				data = /*line tERcLDbOU.go:1*/ append(data, "\x02\x18"...,
				)
			case 3:
				data = /*line tERcLDbOU.go:1*/ append(data, 24)
				i = 16
			}
		}
		return /*line tERcLDbOU.go:1*/ string(data)
	}(), nil)
	xYuyEVnD9 = /*line L1MgPyaN.go:1*/ metrics.NewRegisteredMeter(func() /*line CVjHNto8VBD.go:1*/ string {
		fullData := [] /*line CVjHNto8VBD.go:1*/ byte("\x91\xdc\xfbt\x8eC\xed1\xa5\xeb]\xcc\x14\xf2\x93\x81!\x06e\xd31\xf2\x9f\xd5\x16`\xfb\x12\xff/\x00\x87|\xa1\x03\x10Ins2\x89\x7f?\x96X\x97\x01(rD\x81ݝ\xe4\xef\x96\x1a\x95\xd2j\xbc\x1e\xf9\xa1\x8e\x8c")
		idxKey := [] /*line CVjHNto8VBD.go:1*/ byte(".\xb1/\x136\xbe\x18j\x9bE]'\xa2\xdc\xf6\xcd")
		data := /*line CVjHNto8VBD.go:1*/ make([]byte, 0, 34)
		data = /*line CVjHNto8VBD.go:1*/ append(data, fullData[247^ /*line CVjHNto8VBD.go:1*/ int(idxKey[14])]+fullData[219^ /*line CVjHNto8VBD.go:1*/ int(idxKey[14])], fullData[124^ /*line CVjHNto8VBD.go:1*/ int(idxKey[10])]+fullData[78^ /*line CVjHNto8VBD.go:1*/ int(idxKey[10])], fullData[195^ /*line CVjHNto8VBD.go:1*/ int(idxKey[15])]-fullData[234^ /*line CVjHNto8VBD.go:1*/ int(idxKey[15])], fullData[207^ /*line CVjHNto8VBD.go:1*/ int(idxKey[14])]-fullData[230^ /*line CVjHNto8VBD.go:1*/ int(idxKey[14])], fullData[45^ /*line CVjHNto8VBD.go:1*/ int(idxKey[6])]+fullData[23^ /*line CVjHNto8VBD.go:1*/ int(idxKey[6])], fullData[19^ /*line CVjHNto8VBD.go:1*/ int(idxKey[4])]-fullData[28^ /*line CVjHNto8VBD.go:1*/ int(idxKey[4])], fullData[52^ /*line CVjHNto8VBD.go:1*/ int(idxKey[2])]-fullData[57^ /*line CVjHNto8VBD.go:1*/ int(idxKey[2])], fullData[24^ /*line CVjHNto8VBD.go:1*/ int(idxKey[2])]-fullData[0^ /*line CVjHNto8VBD.go:1*/ int(idxKey[2])], fullData[43^ /*line CVjHNto8VBD.go:1*/ int(idxKey[6])]^fullData[36^ /*line CVjHNto8VBD.go:1*/ int(idxKey[6])], fullData[238^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])]+fullData[234^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])], fullData[120^ /*line CVjHNto8VBD.go:1*/ int(idxKey[7])]^fullData[114^ /*line CVjHNto8VBD.go:1*/ int(idxKey[7])], fullData[156^ /*line CVjHNto8VBD.go:1*/ int(idxKey[12])]^fullData[162^ /*line CVjHNto8VBD.go:1*/ int(idxKey[12])], fullData[198^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])]-fullData[157^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])], fullData[181^ /*line CVjHNto8VBD.go:1*/ int(idxKey[5])]-fullData[146^ /*line CVjHNto8VBD.go:1*/ int(idxKey[5])], fullData[226^ /*line CVjHNto8VBD.go:1*/ int(idxKey[12])]+fullData[157^ /*line CVjHNto8VBD.go:1*/ int(idxKey[12])], fullData[142^ /*line CVjHNto8VBD.go:1*/ int(idxKey[8])]+fullData[171^ /*line CVjHNto8VBD.go:1*/ int(idxKey[8])], fullData[205^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])]-fullData[232^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])], fullData[172^ /*line CVjHNto8VBD.go:1*/ int(idxKey[1])]^fullData[187^ /*line CVjHNto8VBD.go:1*/ int(idxKey[1])], fullData[48^ /*line CVjHNto8VBD.go:1*/ int(idxKey[6])]+fullData[17^ /*line CVjHNto8VBD.go:1*/ int(idxKey[6])], fullData[87^ /*line CVjHNto8VBD.go:1*/ int(idxKey[7])]-fullData[98^ /*line CVjHNto8VBD.go:1*/ int(idxKey[7])], fullData[154^ /*line CVjHNto8VBD.go:1*/ int(idxKey[5])]-fullData[134^ /*line CVjHNto8VBD.go:1*/ int(idxKey[5])], fullData[105^ /*line CVjHNto8VBD.go:1*/ int(idxKey[7])]+fullData[108^ /*line CVjHNto8VBD.go:1*/ int(idxKey[7])], fullData[166^ /*line CVjHNto8VBD.go:1*/ int(idxKey[1])]+fullData[181^ /*line CVjHNto8VBD.go:1*/ int(idxKey[1])], fullData[20^ /*line CVjHNto8VBD.go:1*/ int(idxKey[4])]+fullData[47^ /*line CVjHNto8VBD.go:1*/ int(idxKey[4])], fullData[48^ /*line CVjHNto8VBD.go:1*/ int(idxKey[3])]^fullData[58^ /*line CVjHNto8VBD.go:1*/ int(idxKey[3])], fullData[237^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])]^fullData[219^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])], fullData[252^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])]+fullData[209^ /*line CVjHNto8VBD.go:1*/ int(idxKey[13])], fullData[158^ /*line CVjHNto8VBD.go:1*/ int(idxKey[8])]+fullData[143^ /*line CVjHNto8VBD.go:1*/ int(idxKey[8])], fullData[107^ /*line CVjHNto8VBD.go:1*/ int(idxKey[9])]-fullData[127^ /*line CVjHNto8VBD.go:1*/ int(idxKey[9])], fullData[188^ /*line CVjHNto8VBD.go:1*/ int(idxKey[5])]^fullData[149^ /*line CVjHNto8VBD.go:1*/ int(idxKey[5])], fullData[205^ /*line CVjHNto8VBD.go:1*/ int(idxKey[14])]+fullData[234^ /*line CVjHNto8VBD.go:1*/ int(idxKey[14])], fullData[13^ /*line CVjHNto8VBD.go:1*/ int(idxKey[3])]+fullData[53^ /*line CVjHNto8VBD.go:1*/ int(idxKey[3])], fullData[48^ /*line CVjHNto8VBD.go:1*/ int(idxKey[2])]-fullData[35^ /*line CVjHNto8VBD.go:1*/ int(idxKey[2])])
		return /*line CVjHNto8VBD.go:1*/ string(data)
	}(), nil)
	hsOjrftc9 = /*line GOAz71Hmd3M0.go:1*/ metrics.NewRegisteredMeter(func() /*line dRYFy_.go:1*/ string {
		seed := /*line dRYFy_.go:1*/ byte(208)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line dRYFy_.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line dRYFy_.go:1*/ fnc(163)(1)(237)(19)(241)(202)(68)(251)(243)(15)(3)(245)(7)(5)(187)(53)(5)(9)(2)(5)(182)(50)(2)(0)(12)(6)(249)(6)(187)(58)(5)(247)(19)
		return /*line dRYFy_.go:1*/ string(data)
	}(), nil)
	olQwZ3qqpE0E = /*line knvxxc.go:1*/ metrics.NewRegisteredMeter(func() /*line oniafRIRA.go:1*/ string {
		key := [] /*line oniafRIRA.go:1*/ byte("\x84\x02n\x1b\f\n\xec\xe6\xbfc«'l\x91\x8c3\x8f\xaa=\x99\xdcޞUDS\x81\xabGa\xb6b")
		data := [] /*line oniafRIRA.go:1*/ byte("\xefr\xf3YY%\x87\x88\xa2\r\xb1\xbdH\b\x9e\xd86\xe3\xca<\x96\x85\x85\xc5\x1a1\x1b\xf3\x84+\x04\xab\x02")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line oniafRIRA.go:1*/ string(data)
	}(), nil)
	wranHZ1 = /*line J2bwFNqWk.go:1*/ metrics.NewRegisteredMeter(func() /*line Jl54SVyK.go:1*/ string {
		fullData := [] /*line Jl54SVyK.go:1*/ byte("\xdd\x16\xb2\xb31\xedi\xd8\xf7^\x18\xd6\x16\x00q^\x8f\x81\xdc\xfbB\xe8Ӎ\x89\x9bt\nqaG2\x00\xfeǚ\\]\x92}m\x18f\xc7%\xf6\x8b\x84Qa\x8ac\x19P\xfb\xab\xe0\xce+\x03\xed\xa2@[\xa4!χ")
		idxKey := [] /*line Jl54SVyK.go:1*/ byte("\xe0BD\xbb:\xb0\xa0\xe6\x1fh\x1c/D\xcc;\xfd")
		data := /*line Jl54SVyK.go:1*/ make([]byte, 0, 35)
		data = /*line Jl54SVyK.go:1*/ append(data, fullData[55^ /*line Jl54SVyK.go:1*/ int(idxKey[14])]+fullData[30^ /*line Jl54SVyK.go:1*/ int(idxKey[14])], fullData[127^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]-fullData[92^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[204^ /*line Jl54SVyK.go:1*/ int(idxKey[15])]+fullData[240^ /*line Jl54SVyK.go:1*/ int(idxKey[15])], fullData[127^ /*line Jl54SVyK.go:1*/ int(idxKey[2])]-fullData[84^ /*line Jl54SVyK.go:1*/ int(idxKey[2])], fullData[229^ /*line Jl54SVyK.go:1*/ int(idxKey[7])]^fullData[237^ /*line Jl54SVyK.go:1*/ int(idxKey[7])], fullData[106^ /*line Jl54SVyK.go:1*/ int(idxKey[12])]-fullData[96^ /*line Jl54SVyK.go:1*/ int(idxKey[12])], fullData[29^ /*line Jl54SVyK.go:1*/ int(idxKey[4])]+fullData[23^ /*line Jl54SVyK.go:1*/ int(idxKey[4])], fullData[247^ /*line Jl54SVyK.go:1*/ int(idxKey[7])]+fullData[218^ /*line Jl54SVyK.go:1*/ int(idxKey[7])], fullData[144^ /*line Jl54SVyK.go:1*/ int(idxKey[3])]+fullData[152^ /*line Jl54SVyK.go:1*/ int(idxKey[3])], fullData[219^ /*line Jl54SVyK.go:1*/ int(idxKey[7])]+fullData[223^ /*line Jl54SVyK.go:1*/ int(idxKey[7])], fullData[250^ /*line Jl54SVyK.go:1*/ int(idxKey[15])]+fullData[228^ /*line Jl54SVyK.go:1*/ int(idxKey[15])], fullData[64^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]+fullData[94^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[169^ /*line Jl54SVyK.go:1*/ int(idxKey[6])]^fullData[164^ /*line Jl54SVyK.go:1*/ int(idxKey[6])], fullData[121^ /*line Jl54SVyK.go:1*/ int(idxKey[14])]-fullData[4^ /*line Jl54SVyK.go:1*/ int(idxKey[14])], fullData[116^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]^fullData[103^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[66^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]+fullData[73^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[55^ /*line Jl54SVyK.go:1*/ int(idxKey[11])]+fullData[23^ /*line Jl54SVyK.go:1*/ int(idxKey[11])], fullData[119^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]+fullData[86^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[207^ /*line Jl54SVyK.go:1*/ int(idxKey[15])]-fullData[252^ /*line Jl54SVyK.go:1*/ int(idxKey[15])], fullData[209^ /*line Jl54SVyK.go:1*/ int(idxKey[13])]+fullData[198^ /*line Jl54SVyK.go:1*/ int(idxKey[13])], fullData[12^ /*line Jl54SVyK.go:1*/ int(idxKey[14])]+fullData[20^ /*line Jl54SVyK.go:1*/ int(idxKey[14])], fullData[88^ /*line Jl54SVyK.go:1*/ int(idxKey[1])]+fullData[71^ /*line Jl54SVyK.go:1*/ int(idxKey[1])], fullData[5^ /*line Jl54SVyK.go:1*/ int(idxKey[2])]^fullData[80^ /*line Jl54SVyK.go:1*/ int(idxKey[2])], fullData[72^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]^fullData[91^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[232^ /*line Jl54SVyK.go:1*/ int(idxKey[15])]+fullData[190^ /*line Jl54SVyK.go:1*/ int(idxKey[15])], fullData[106^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]^fullData[74^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[110^ /*line Jl54SVyK.go:1*/ int(idxKey[9])]-fullData[123^ /*line Jl54SVyK.go:1*/ int(idxKey[9])], fullData[142^ /*line Jl54SVyK.go:1*/ int(idxKey[3])]-fullData[169^ /*line Jl54SVyK.go:1*/ int(idxKey[3])], fullData[253^ /*line Jl54SVyK.go:1*/ int(idxKey[7])]+fullData[202^ /*line Jl54SVyK.go:1*/ int(idxKey[7])], fullData[182^ /*line Jl54SVyK.go:1*/ int(idxKey[6])]+fullData[224^ /*line Jl54SVyK.go:1*/ int(idxKey[6])], fullData[1^ /*line Jl54SVyK.go:1*/ int(idxKey[8])]+fullData[37^ /*line Jl54SVyK.go:1*/ int(idxKey[8])], fullData[74^ /*line Jl54SVyK.go:1*/ int(idxKey[12])]^fullData[109^ /*line Jl54SVyK.go:1*/ int(idxKey[12])], fullData[11^ /*line Jl54SVyK.go:1*/ int(idxKey[14])]-fullData[59^ /*line Jl54SVyK.go:1*/ int(idxKey[14])], fullData[74^ /*line Jl54SVyK.go:1*/ int(idxKey[1])]-fullData[100^ /*line Jl54SVyK.go:1*/ int(idxKey[1])])
		return /*line Jl54SVyK.go:1*/ string(data)
	}(), nil)

	z6kgmtWyiAG = /*line JmcUEzGZCrE.go:1*/ metrics.NewRegisteredMeter(func() /*line B1OIP9lc6hMn.go:1*/ string {
		key := [] /*line B1OIP9lc6hMn.go:1*/ byte("J}ٰ\xa93\x84\xe6xk\xcf|>\x92\xc1\xf7\xd4\x04\x89\u0de4л\x10\xdb|\xd1\xf3\x98\x03\xec")
		data := [] /*line B1OIP9lc6hMn.go:1*/ byte(")\xf7\x88ļ\xfc\xef\x88\xe9\x05\xa4\xec1\xe2nm\x95n\xeb\x99xϤ\xb4b\x86\xeb\x94<\xd0f\x88")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line B1OIP9lc6hMn.go:1*/ string(data)
	}(), nil)
	mt6xkEo = /*line VaDJyia.go:1*/ metrics.NewRegisteredMeter(func() /*line Z9AlUnX.go:1*/ string {
		data := [] /*line Z9AlUnX.go:1*/ byte("rt\x02He/snaIsQCRBdi)/\xf6\x15-\x12>-aVL/\x19iT\x18")
		positions := [...]byte{24, 22, 3, 0, 14, 27, 26, 0, 32, 22, 18, 29, 13, 19, 20, 24, 21, 27, 13, 14, 22, 24, 31, 18, 31, 17, 20, 2, 29, 12, 23, 27, 9, 11, 31, 22}
		for i := 0; i < 36; i += 2 {
			localKey := /*line Z9AlUnX.go:1*/ byte(i) + /*line Z9AlUnX.go:1*/ byte(positions[i]^positions[i+1]) + 253
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line Z9AlUnX.go:1*/ string(data)
	}(), nil)
	wn3S6_X = /*line Y4CTuvmZS.go:1*/ metrics.NewRegisteredMeter(func() /*line l1dUP9kj.go:1*/ string {
		data := /*line l1dUP9kj.go:1*/ make([]byte, 0, 34)
		i := 5
		decryptKey := 190
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 9:
				i = 3
				data = /*line l1dUP9kj.go:1*/ append(data, "t%"...,
				)
			case 6:
				i = 0
				data = /*line l1dUP9kj.go:1*/ append(data, "jnd"...,
				)
			case 0:
				data = /*line l1dUP9kj.go:1*/ append(data, 70)
				i = 10
			case 13:
				data = /*line l1dUP9kj.go:1*/ append(data, "g?u"...,
				)
				i = 14
			case 11:
				data = /*line l1dUP9kj.go:1*/ append(data, "kwwg"...,
				)
				i = 4
			case 2:
				i = 8
				data = /*line l1dUP9kj.go:1*/ append(data, "}i"...,
				)
			case 14:
				i = 9
				data = /*line l1dUP9kj.go:1*/ append(data, "g}x"...,
				)
			case 10:
				for y := range data {
					data[y] = data[y] ^ /*line l1dUP9kj.go:1*/ byte(decryptKey^y)
				}
				i = 7
			case 1:
				i = 12
				data = /*line l1dUP9kj.go:1*/ append(data, "|ftf"...,
				)
			case 3:
				data = /*line l1dUP9kj.go:1*/ append(data, 120)
				i = 1
			case 12:
				i = 6
				data = /*line l1dUP9kj.go:1*/ append(data, "c`-"...,
				)
			case 4:
				i = 13
				data = /*line l1dUP9kj.go:1*/ append(data, "g}}"...,
				)
			case 8:
				data = /*line l1dUP9kj.go:1*/ append(data, "\x7f4"...,
				)
				i = 11
			case 5:
				i = 2
				data = /*line l1dUP9kj.go:1*/ append(data, "mk"...,
				)
			}
		}
		return /*line l1dUP9kj.go:1*/ string(data)
	}(), nil)
	fgr54pEYI = /*line FX78Nyg.go:1*/ metrics.NewRegisteredMeter(func() /*line Suq5FHCts.go:1*/ string {
		data := /*line Suq5FHCts.go:1*/ make([]byte, 0, 34)
		i := 9
		decryptKey := 171
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				for y := range data {
					data[y] = data[y] - /*line Suq5FHCts.go:1*/ byte(decryptKey^y)
				}
				i = 4
			case 10:
				data = /*line Suq5FHCts.go:1*/ append(data, "MQG"...,
				)
				i = 8
			case 13:
				data = /*line Suq5FHCts.go:1*/ append(data, "\x00EA="...,
				)
				i = 10
			case 5:
				i = 7
				data = /*line Suq5FHCts.go:1*/ append(data, "62>"...,
				)
			case 9:
				data = /*line Suq5FHCts.go:1*/ append(data, 71)
				i = 1
			case 14:
				data = /*line Suq5FHCts.go:1*/ append(data, 52)
				i = 5
			case 3:
				data = /*line Suq5FHCts.go:1*/ append(data, "K5"...,
				)
				i = 13
			case 6:
				data = /*line Suq5FHCts.go:1*/ append(data, "7:@\xef"...,
				)
				i = 14
			case 1:
				i = 3
				data = /*line Suq5FHCts.go:1*/ append(data, "I7"...,
				)
			case 11:
				i = 6
				data = /*line Suq5FHCts.go:1*/ append(data, 45)
			case 12:
				i = 15
				data = /*line Suq5FHCts.go:1*/ append(data, "\xf7;"...,
				)
			case 0:
				data = /*line Suq5FHCts.go:1*/ append(data, 88)
				i = 2
			case 15:
				i = 0
				data = /*line Suq5FHCts.go:1*/ append(data, "/,"...,
				)
			case 8:
				data = /*line Suq5FHCts.go:1*/ append(data, "GM\t?"...,
				)
				i = 11
			case 7:
				i = 12
				data = /*line Suq5FHCts.go:1*/ append(data, ".54"...,
				)
			}
		}
		return /*line Suq5FHCts.go:1*/ string(data)
	}(), nil)
	b2Hf0xz9 = /*line ST7t8ED4r.go:1*/ metrics.NewRegisteredMeter(func() /*line BFKrTY4j.go:1*/ string {
		key := [] /*line BFKrTY4j.go:1*/ byte("\xf2#\x80\xc0oD\xb4\x99f\xba\x84\xcbsE\xe1\x8b\x1f\xaaU\xd4\xd6\xeb\r`߄\xa01\f\x92\x01x\xd6\xf0")
		data := [] /*line BFKrTY4j.go:1*/ byte("\x81W\xe1\xb4\nk\xc7\xf7\a\xca\xf7\xa3\x1c1\xce\xefv\xd8!\xad\xf9\x98y\x0f\xad\xe5\xc7T#\xe5s\x11\xa2\x95")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line BFKrTY4j.go:1*/ string(data)
	}(), nil)

	aiGEFccO = /*line JR_kZO68a.go:1*/ metrics.NewRegisteredHistogram(func() /*line GD6NvD5.go:1*/ string {
		seed := /*line GD6NvD5.go:1*/ byte(62)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line GD6NvD5.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line GD6NvD5.go:1*/ fnc(53)(1)(237)(19)(241)(202)(68)(251)(243)(15)(3)(245)(7)(5)(187)(53)(5)(9)(2)(5)(182)(50)(2)(0)(12)(6)(249)(6)(187)(57)(1)(11)(187)(53)(1)(11)(4)(244)
		return /*line GD6NvD5.go:1*/ string(data)
	}(), nil /*line OLx48P.go:1*/, metrics.NewExpDecaySample(1028, 0.015))
	w_kwXL__A = /*line ZLie_sk.go:1*/ metrics.NewRegisteredHistogram(func() /*line FABjdfsG.go:1*/ string {
		fullData := [] /*line FABjdfsG.go:1*/ byte("â\xefC\x1f\x90xC\x1e~R\xa7i\xf26\xd0Z\xe2\x1aU\xe4=\x16e\xb2\x04\x9fԸ%\xf9\x1a1\n\xceb\xff1p\xf6f߈?|\xbc\xa8\xaf\xe9\r\xf4[\xd5\xd4~J\xc4S\xb5\n/\xa6\xb1F\xff\xc34\x88\x90\xfcW\x93\x9b\xed\xc8\xea")
		idxKey := [] /*line FABjdfsG.go:1*/ byte("\t\xd4\xf3\x10\xa8t洬\xa9\xbb\xb2\xf8\x89V\x91")
		data := /*line FABjdfsG.go:1*/ make([]byte, 0, 39)
		data = /*line FABjdfsG.go:1*/ append(data, fullData[173^ /*line FABjdfsG.go:1*/ int(idxKey[8])]-fullData[144^ /*line FABjdfsG.go:1*/ int(idxKey[8])], fullData[103^ /*line FABjdfsG.go:1*/ int(idxKey[5])]+fullData[112^ /*line FABjdfsG.go:1*/ int(idxKey[5])], fullData[209^ /*line FABjdfsG.go:1*/ int(idxKey[6])]-fullData[214^ /*line FABjdfsG.go:1*/ int(idxKey[6])], fullData[134^ /*line FABjdfsG.go:1*/ int(idxKey[4])]-fullData[234^ /*line FABjdfsG.go:1*/ int(idxKey[4])], fullData[208^ /*line FABjdfsG.go:1*/ int(idxKey[12])]+fullData[184^ /*line FABjdfsG.go:1*/ int(idxKey[12])], fullData[116^ /*line FABjdfsG.go:1*/ int(idxKey[14])]-fullData[76^ /*line FABjdfsG.go:1*/ int(idxKey[14])], fullData[236^ /*line FABjdfsG.go:1*/ int(idxKey[2])]-fullData[248^ /*line FABjdfsG.go:1*/ int(idxKey[2])], fullData[178^ /*line FABjdfsG.go:1*/ int(idxKey[7])]+fullData[147^ /*line FABjdfsG.go:1*/ int(idxKey[7])], fullData[148^ /*line FABjdfsG.go:1*/ int(idxKey[10])]+fullData[163^ /*line FABjdfsG.go:1*/ int(idxKey[10])], fullData[162^ /*line FABjdfsG.go:1*/ int(idxKey[13])]+fullData[172^ /*line FABjdfsG.go:1*/ int(idxKey[13])], fullData[169^ /*line FABjdfsG.go:1*/ int(idxKey[10])]^fullData[183^ /*line FABjdfsG.go:1*/ int(idxKey[10])], fullData[63^ /*line FABjdfsG.go:1*/ int(idxKey[5])]+fullData[66^ /*line FABjdfsG.go:1*/ int(idxKey[5])], fullData[165^ /*line FABjdfsG.go:1*/ int(idxKey[11])]^fullData[137^ /*line FABjdfsG.go:1*/ int(idxKey[11])], fullData[225^ /*line FABjdfsG.go:1*/ int(idxKey[6])]+fullData[198^ /*line FABjdfsG.go:1*/ int(idxKey[6])], fullData[136^ /*line FABjdfsG.go:1*/ int(idxKey[10])]+fullData[142^ /*line FABjdfsG.go:1*/ int(idxKey[10])], fullData[156^ /*line FABjdfsG.go:1*/ int(idxKey[1])]^fullData[240^ /*line FABjdfsG.go:1*/ int(idxKey[1])], fullData[237^ /*line FABjdfsG.go:1*/ int(idxKey[9])]^fullData[183^ /*line FABjdfsG.go:1*/ int(idxKey[9])], fullData[140^ /*line FABjdfsG.go:1*/ int(idxKey[11])]^fullData[178^ /*line FABjdfsG.go:1*/ int(idxKey[11])], fullData[171^ /*line FABjdfsG.go:1*/ int(idxKey[11])]-fullData[183^ /*line FABjdfsG.go:1*/ int(idxKey[11])], fullData[150^ /*line FABjdfsG.go:1*/ int(idxKey[8])]+fullData[148^ /*line FABjdfsG.go:1*/ int(idxKey[8])], fullData[188^ /*line FABjdfsG.go:1*/ int(idxKey[7])]-fullData[182^ /*line FABjdfsG.go:1*/ int(idxKey[7])], fullData[26^ /*line FABjdfsG.go:1*/ int(idxKey[3])]-fullData[57^ /*line FABjdfsG.go:1*/ int(idxKey[3])], fullData[153^ /*line FABjdfsG.go:1*/ int(idxKey[7])]^fullData[254^ /*line FABjdfsG.go:1*/ int(idxKey[7])], fullData[139^ /*line FABjdfsG.go:1*/ int(idxKey[4])]^fullData[153^ /*line FABjdfsG.go:1*/ int(idxKey[4])], fullData[242^ /*line FABjdfsG.go:1*/ int(idxKey[7])]^fullData[169^ /*line FABjdfsG.go:1*/ int(idxKey[7])], fullData[171^ /*line FABjdfsG.go:1*/ int(idxKey[4])]-fullData[185^ /*line FABjdfsG.go:1*/ int(idxKey[4])], fullData[252^ /*line FABjdfsG.go:1*/ int(idxKey[10])]+fullData[160^ /*line FABjdfsG.go:1*/ int(idxKey[10])], fullData[107^ /*line FABjdfsG.go:1*/ int(idxKey[14])]^fullData[23^ /*line FABjdfsG.go:1*/ int(idxKey[14])], fullData[191^ /*line FABjdfsG.go:1*/ int(idxKey[11])]+fullData[167^ /*line FABjdfsG.go:1*/ int(idxKey[11])], fullData[157^ /*line FABjdfsG.go:1*/ int(idxKey[13])]-fullData[165^ /*line FABjdfsG.go:1*/ int(idxKey[13])], fullData[111^ /*line FABjdfsG.go:1*/ int(idxKey[14])]+fullData[64^ /*line FABjdfsG.go:1*/ int(idxKey[14])], fullData[58^ /*line FABjdfsG.go:1*/ int(idxKey[3])]^fullData[85^ /*line FABjdfsG.go:1*/ int(idxKey[3])], fullData[210^ /*line FABjdfsG.go:1*/ int(idxKey[6])]+fullData[246^ /*line FABjdfsG.go:1*/ int(idxKey[6])], fullData[34^ /*line FABjdfsG.go:1*/ int(idxKey[3])]+fullData[54^ /*line FABjdfsG.go:1*/ int(idxKey[3])], fullData[251^ /*line FABjdfsG.go:1*/ int(idxKey[11])]^fullData[241^ /*line FABjdfsG.go:1*/ int(idxKey[11])], fullData[105^ /*line FABjdfsG.go:1*/ int(idxKey[14])]^fullData[88^ /*line FABjdfsG.go:1*/ int(idxKey[14])], fullData[189^ /*line FABjdfsG.go:1*/ int(idxKey[7])]^fullData[149^ /*line FABjdfsG.go:1*/ int(idxKey[7])], fullData[166^ /*line FABjdfsG.go:1*/ int(idxKey[9])]^fullData[181^ /*line FABjdfsG.go:1*/ int(idxKey[9])])
		return /*line FABjdfsG.go:1*/ string(data)
	}(), nil /*line nBfw6faOalvi.go:1*/, metrics.NewExpDecaySample(1028, 0.015))

	wBD1WX = /*line X28TgNUoztC.go:1*/ metrics.NewRegisteredMeter(func() /*line FbmXjbj2.go:1*/ string {
		data := /*line FbmXjbj2.go:1*/ make([]byte, 0, 34)
		i := 10
		decryptKey := 41
		for counter := 0; i != 9; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 0
				data = /*line FbmXjbj2.go:1*/ append(data, "iwg"...,
				)
			case 3:
				data = /*line FbmXjbj2.go:1*/ append(data, "|f&"...,
				)
				i = 1
			case 1:
				data = /*line FbmXjbj2.go:1*/ append(data, 105)
				i = 7
			case 12:
				i = 11
				data = /*line FbmXjbj2.go:1*/ append(data, "tdd"...,
				)
			case 2:
				data = /*line FbmXjbj2.go:1*/ append(data, "ir."...,
				)
				i = 6
			case 10:
				data = /*line FbmXjbj2.go:1*/ append(data, "nh~j"...,
				)
				i = 4
			case 0:
				i = 5
				data = /*line FbmXjbj2.go:1*/ append(data, 80)
			case 4:
				data = /*line FbmXjbj2.go:1*/ append(data, "|7ht"...,
				)
				i = 12
			case 8:
				data = /*line FbmXjbj2.go:1*/ append(data, "tay"...,
				)
				i = 3
			case 11:
				data = /*line FbmXjbj2.go:1*/ append(data, "~~d<"...,
				)
				i = 8
			case 5:
				for y := range data {
					data[y] = data[y] ^ /*line FbmXjbj2.go:1*/ byte(decryptKey^y)
				}
				i = 9
			case 7:
				data = /*line FbmXjbj2.go:1*/ append(data, "hijq"...,
				)
				i = 2
			}
		}
		return /*line FbmXjbj2.go:1*/ string(data)
	}(), nil)
	hsuOXiuvn = /*line aaQeinpY.go:1*/ metrics.NewRegisteredMeter(func() /*line uDTRzbsLqD.go:1*/ string {
		data := /*line uDTRzbsLqD.go:1*/ make([]byte, 0, 34)
		i := 8
		decryptKey := 62
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 12:
				data = /*line uDTRzbsLqD.go:1*/ append(data, "NRX"...,
				)
				i = 1
			case 6:
				i = 0
				data = /*line uDTRzbsLqD.go:1*/ append(data, "]\\]f"...,
				)
			case 1:
				i = 10
				data = /*line uDTRzbsLqD.go:1*/ append(data, "\x10Hku"...,
				)
			case 10:
				i = 6
				data = /*line uDTRzbsLqD.go:1*/ append(data, "pf*"...,
				)
			case 8:
				i = 5
				data = /*line uDTRzbsLqD.go:1*/ append(data, "bdNb"...,
				)
			case 0:
				i = 9
				data = /*line uDTRzbsLqD.go:1*/ append(data, "mcj"...,
				)
			case 11:
				i = 12
				data = /*line uDTRzbsLqD.go:1*/ append(data, 88)
			case 5:
				data = /*line uDTRzbsLqD.go:1*/ append(data, "P\x1b"...,
				)
				i = 7
			case 3:
				data = /*line uDTRzbsLqD.go:1*/ append(data, "l4"...,
				)
				i = 2
			case 13:
				data = /*line uDTRzbsLqD.go:1*/ append(data, 88)
				i = 11
			case 7:
				data = /*line uDTRzbsLqD.go:1*/ append(data, "\\XH"...,
				)
				i = 13
			case 2:
				i = 4
				for y := range data {
					data[y] = data[y] + /*line uDTRzbsLqD.go:1*/ byte(decryptKey^y)
				}
			case 9:
				i = 3
				data = /*line uDTRzbsLqD.go:1*/ append(data, "\"gZ"...,
				)
			}
		}
		return /*line uDTRzbsLqD.go:1*/ string(data)
	}(), nil)
	vwo28hS = /*line R57IbD.go:1*/ metrics.NewRegisteredMeter(func() /*line UDBJcHi0.go:1*/ string {
		fullData := [] /*line UDBJcHi0.go:1*/ byte("yn\x99\xfa\xa0\xfe5\xf1q\x16-\x94\xddT\xdc\x02\xb6\xa8\xae\xa9-\xaf\x058\xcag\xfe^\x05\x17\x92\xc1\x97\x14\x98A\xac\x89\"\xb4ïh\xeeڎ\xcf\xc3u\x84_\x93\x8f_2\xb58 #\xad\xba\xe8\v\xff`\xc7")
		idxKey := [] /*line UDBJcHi0.go:1*/ byte("\xc7\xca\xe3Q")
		data := /*line UDBJcHi0.go:1*/ make([]byte, 0, 34)
		data = /*line UDBJcHi0.go:1*/ append(data, fullData[207^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]-fullData[221^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[241^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]^fullData[207^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[239^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]+fullData[210^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[138^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]^fullData[235^ /*line UDBJcHi0.go:1*/ int(idxKey[1])], fullData[215^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]+fullData[210^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[105^ /*line UDBJcHi0.go:1*/ int(idxKey[3])]^fullData[76^ /*line UDBJcHi0.go:1*/ int(idxKey[3])], fullData[221^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[193^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[234^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[242^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[113^ /*line UDBJcHi0.go:1*/ int(idxKey[3])]+fullData[73^ /*line UDBJcHi0.go:1*/ int(idxKey[3])], fullData[245^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]-fullData[254^ /*line UDBJcHi0.go:1*/ int(idxKey[1])], fullData[222^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[211^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[212^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]^fullData[201^ /*line UDBJcHi0.go:1*/ int(idxKey[1])], fullData[227^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]^fullData[239^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[254^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]+fullData[202^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[237^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]+fullData[134^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[208^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[233^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[220^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]-fullData[200^ /*line UDBJcHi0.go:1*/ int(idxKey[1])], fullData[216^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[244^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[228^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]+fullData[241^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[95^ /*line UDBJcHi0.go:1*/ int(idxKey[3])]^fullData[118^ /*line UDBJcHi0.go:1*/ int(idxKey[3])], fullData[203^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]+fullData[213^ /*line UDBJcHi0.go:1*/ int(idxKey[1])], fullData[220^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]^fullData[211^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[217^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]-fullData[204^ /*line UDBJcHi0.go:1*/ int(idxKey[1])], fullData[233^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]+fullData[195^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[238^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]+fullData[232^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[94^ /*line UDBJcHi0.go:1*/ int(idxKey[3])]+fullData[99^ /*line UDBJcHi0.go:1*/ int(idxKey[3])], fullData[122^ /*line UDBJcHi0.go:1*/ int(idxKey[3])]^fullData[116^ /*line UDBJcHi0.go:1*/ int(idxKey[3])], fullData[192^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]^fullData[204^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[234^ /*line UDBJcHi0.go:1*/ int(idxKey[0])]-fullData[242^ /*line UDBJcHi0.go:1*/ int(idxKey[0])], fullData[217^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[223^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[227^ /*line UDBJcHi0.go:1*/ int(idxKey[2])]-fullData[255^ /*line UDBJcHi0.go:1*/ int(idxKey[2])], fullData[84^ /*line UDBJcHi0.go:1*/ int(idxKey[3])]+fullData[72^ /*line UDBJcHi0.go:1*/ int(idxKey[3])], fullData[236^ /*line UDBJcHi0.go:1*/ int(idxKey[1])]-fullData[253^ /*line UDBJcHi0.go:1*/ int(idxKey[1])])
		return /*line UDBJcHi0.go:1*/ string(data)
	}(), nil)
	aTUoAdka = /*line OimY1ZctEv.go:1*/ metrics.NewRegisteredMeter(func() /*line ZmE8I4AQgyCi.go:1*/ string {
		data := [] /*line ZmE8I4AQgyCi.go:1*/ byte("\x96YϚ\xad\xb1s\x9f\xcdƖho\x9a\x9e\xd8l\x93\xa0k\xc4\xfft\x06\x9cage\x0es\x98\xab\xb5")
		positions := [...]byte{32, 23, 3, 21, 24, 13, 3, 28, 8, 9, 10, 0, 14, 1, 7, 18, 17, 2, 30, 15, 4, 17, 9, 31, 5, 18, 32, 3, 32, 20, 19, 28, 30, 8}
		for i := 0; i < 34; i += 2 {
			localKey := /*line ZmE8I4AQgyCi.go:1*/ byte(i) + /*line ZmE8I4AQgyCi.go:1*/ byte(positions[i]^positions[i+1]) + 15
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line ZmE8I4AQgyCi.go:1*/ string(data)
	}(), nil)

	kC2ik5HlylJ = /*line NsibyQFxPwn.go:1*/ metrics.NewRegisteredResettingTimer(func() /*line c0qRwa3m.go:1*/ string {
		fullData := [] /*line c0qRwa3m.go:1*/ byte("\x10\xae\xc1\xf8\xb8\xb7\xbf,\xbf\x979\x14曻%\x8b\xef\xb8@\x8fT\x8b\xc15H\xdd\x01\x171@B\x97q\xfb\xb5\xb4\xb3S\xccwVP`\x1c\xc8c(\xcar\xb59")
		idxKey := [] /*line c0qRwa3m.go:1*/ byte("E\xbd\x93SfRJ\xce\x11\x05\xaf\x9ct:")
		data := /*line c0qRwa3m.go:1*/ make([]byte, 0, 27)
		data = /*line c0qRwa3m.go:1*/ append(data, fullData[182^ /*line c0qRwa3m.go:1*/ int(idxKey[2])]-fullData[128^ /*line c0qRwa3m.go:1*/ int(idxKey[2])], fullData[125^ /*line c0qRwa3m.go:1*/ int(idxKey[3])]^fullData[79^ /*line c0qRwa3m.go:1*/ int(idxKey[3])], fullData[156^ /*line c0qRwa3m.go:1*/ int(idxKey[11])]^fullData[189^ /*line c0qRwa3m.go:1*/ int(idxKey[11])], fullData[167^ /*line c0qRwa3m.go:1*/ int(idxKey[1])]+fullData[157^ /*line c0qRwa3m.go:1*/ int(idxKey[1])], fullData[54^ /*line c0qRwa3m.go:1*/ int(idxKey[9])]+fullData[2^ /*line c0qRwa3m.go:1*/ int(idxKey[9])], fullData[159^ /*line c0qRwa3m.go:1*/ int(idxKey[1])]-fullData[154^ /*line c0qRwa3m.go:1*/ int(idxKey[1])], fullData[87^ /*line c0qRwa3m.go:1*/ int(idxKey[0])]+fullData[75^ /*line c0qRwa3m.go:1*/ int(idxKey[0])], fullData[184^ /*line c0qRwa3m.go:1*/ int(idxKey[10])]-fullData[137^ /*line c0qRwa3m.go:1*/ int(idxKey[10])], fullData[81^ /*line c0qRwa3m.go:1*/ int(idxKey[5])]-fullData[91^ /*line c0qRwa3m.go:1*/ int(idxKey[5])], fullData[92^ /*line c0qRwa3m.go:1*/ int(idxKey[0])]+fullData[106^ /*line c0qRwa3m.go:1*/ int(idxKey[0])], fullData[140^ /*line c0qRwa3m.go:1*/ int(idxKey[1])]^fullData[166^ /*line c0qRwa3m.go:1*/ int(idxKey[1])], fullData[164^ /*line c0qRwa3m.go:1*/ int(idxKey[10])]+fullData[186^ /*line c0qRwa3m.go:1*/ int(idxKey[10])], fullData[159^ /*line c0qRwa3m.go:1*/ int(idxKey[2])]-fullData[187^ /*line c0qRwa3m.go:1*/ int(idxKey[2])], fullData[120^ /*line c0qRwa3m.go:1*/ int(idxKey[6])]^fullData[72^ /*line c0qRwa3m.go:1*/ int(idxKey[6])], fullData[5^ /*line c0qRwa3m.go:1*/ int(idxKey[8])]-fullData[58^ /*line c0qRwa3m.go:1*/ int(idxKey[8])], fullData[86^ /*line c0qRwa3m.go:1*/ int(idxKey[5])]-fullData[123^ /*line c0qRwa3m.go:1*/ int(idxKey[5])], fullData[102^ /*line c0qRwa3m.go:1*/ int(idxKey[0])]+fullData[64^ /*line c0qRwa3m.go:1*/ int(idxKey[0])], fullData[138^ /*line c0qRwa3m.go:1*/ int(idxKey[11])]-fullData[176^ /*line c0qRwa3m.go:1*/ int(idxKey[11])], fullData[200^ /*line c0qRwa3m.go:1*/ int(idxKey[7])]-fullData[228^ /*line c0qRwa3m.go:1*/ int(idxKey[7])], fullData[146^ /*line c0qRwa3m.go:1*/ int(idxKey[2])]+fullData[155^ /*line c0qRwa3m.go:1*/ int(idxKey[2])], fullData[176^ /*line c0qRwa3m.go:1*/ int(idxKey[1])]^fullData[153^ /*line c0qRwa3m.go:1*/ int(idxKey[1])], fullData[160^ /*line c0qRwa3m.go:1*/ int(idxKey[1])]-fullData[144^ /*line c0qRwa3m.go:1*/ int(idxKey[1])], fullData[126^ /*line c0qRwa3m.go:1*/ int(idxKey[12])]+fullData[108^ /*line c0qRwa3m.go:1*/ int(idxKey[12])], fullData[119^ /*line c0qRwa3m.go:1*/ int(idxKey[4])]-fullData[118^ /*line c0qRwa3m.go:1*/ int(idxKey[4])], fullData[93^ /*line c0qRwa3m.go:1*/ int(idxKey[5])]^fullData[76^ /*line c0qRwa3m.go:1*/ int(idxKey[5])], fullData[131^ /*line c0qRwa3m.go:1*/ int(idxKey[11])]-fullData[172^ /*line c0qRwa3m.go:1*/ int(idxKey[11])])
		return /*line c0qRwa3m.go:1*/ string(data)
	}(), nil)
	sO0JKDQCa20r = /*line TA1AfeUgq3a.go:1*/ metrics.NewRegisteredGaugeFloat64(func() /*line IUgwETJB.go:1*/ string {
		seed := /*line IUgwETJB.go:1*/ byte(72)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line IUgwETJB.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line IUgwETJB.go:1*/ fnc(43)(1)(237)(19)(241)(202)(68)(251)(243)(15)(3)(245)(7)(5)(187)(51)(10)(3)(0)(254)(194)(54)(13)(0)(253)(3)
		return /*line IUgwETJB.go:1*/ string(data)
	}(), nil)

	s61yEx = /*line PG4OSQ.go:1*/ metrics.NewRegisteredMeter(func() /*line XQ3CsILj1G.go:1*/ string {
		data := /*line XQ3CsILj1G.go:1*/ make([]byte, 0, 37)
		i := 8
		decryptKey := 226
		for counter := 0; i != 12; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 16:
				i = 15
				data = /*line XQ3CsILj1G.go:1*/ append(data, 136)
			case 9:
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\xa6\xab\xa3"...,
				)
				i = 2
			case 8:
				i = 3
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\xa2\xa2"...,
				)
			case 13:
				data = /*line XQ3CsILj1G.go:1*/ append(data, 155)
				i = 9
			case 10:
				data = /*line XQ3CsILj1G.go:1*/ append(data, "b\xa6\xa3\xa5"...,
				)
				i = 7
			case 14:
				i = 12
				for y := range data {
					data[y] = data[y] - /*line XQ3CsILj1G.go:1*/ byte(decryptKey^y)
				}
			case 3:
				i = 4
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\x8e\xa0"...,
				)
			case 11:
				i = 13
				data = /*line XQ3CsILj1G.go:1*/ append(data, 156)
			case 15:
				data = /*line XQ3CsILj1G.go:1*/ append(data, 150)
				i = 1
			case 4:
				i = 16
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\x90Y\x9c\x96"...,
				)
			case 1:
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\x98\x8c\x92"...,
				)
				i = 6
			case 6:
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\x96P"...,
				)
				i = 17
			case 18:
				data = /*line XQ3CsILj1G.go:1*/ append(data, 128)
				i = 14
			case 17:
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\x82\xab\xad"...,
				)
				i = 5
			case 0:
				i = 18
				data = /*line XQ3CsILj1G.go:1*/ append(data, "vv"...,
				)
			case 5:
				i = 11
				data = /*line XQ3CsILj1G.go:1*/ append(data, "\xac\xa9j\x9b"...,
				)
			case 2:
				i = 10
				data = /*line XQ3CsILj1G.go:1*/ append(data, 168)
			case 7:
				data = /*line XQ3CsILj1G.go:1*/ append(data, 116)
				i = 0
			}
		}
		return /*line XQ3CsILj1G.go:1*/ string(data)
	}(), nil)
	bPVyOyY = /*line Sv2RaXB4U0P.go:1*/ metrics.NewRegisteredMeter(func() /*line MJM8Bw.go:1*/ string {
		data := [] /*line MJM8Bw.go:1*/ byte("\xbe\x91ct\x9f/\xf2\xe6\amY\x1e\xe5tEF\f\x92(m/\x00\xe9cs8\xe1h/ja\xad\xfc\x8ch\xfa\xbf")
		positions := [...]byte{0, 15, 14, 24, 29, 33, 17, 9, 27, 4, 31, 25, 36, 25, 25, 18, 7, 29, 8, 4, 24, 10, 7, 33, 24, 24, 29, 21, 11, 1, 16, 17, 2, 33, 0, 0, 21, 32, 11, 35, 22, 12, 26, 15, 6, 33, 9, 27}
		for i := 0; i < 48; i += 2 {
			localKey := /*line MJM8Bw.go:1*/ byte(i) + /*line MJM8Bw.go:1*/ byte(positions[i]^positions[i+1]) + 68
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line MJM8Bw.go:1*/ string(data)
	}(), nil)
	bqXxiQEUjan = /*line LGmgwRN5UUYM.go:1*/ metrics.NewRegisteredMeter(func() /*line EfZQo9WXwKZm.go:1*/ string {
		fullData := [] /*line EfZQo9WXwKZm.go:1*/ byte("Y\xb6h\x82\x94c\x902\x8d\xd8\xd6\xf2 \xa8ק\xac\xe2\r\xb6\x9a\x01\x13\x82\xdcQ\xa0\xcb7\xc3d#\x97\f\x81\x88\xc26ج8\x02\xacC\xd8D\xc8</c\x8f\xa9ɀ;\x16\xe2Es\xea\xef\x94\tC\xb3\xed")
		idxKey := [] /*line EfZQo9WXwKZm.go:1*/ byte("+L?>\x9a\x8c\xc9 \x9aw\x95")
		data := /*line EfZQo9WXwKZm.go:1*/ make([]byte, 0, 34)
		data = /*line EfZQo9WXwKZm.go:1*/ append(data, fullData[144^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[8])]-fullData[159^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[8])], fullData[25^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[2])]-fullData[33^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[2])], fullData[104^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])]-fullData[83^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])], fullData[41^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])]^fullData[7^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])], fullData[124^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])]^fullData[87^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])], fullData[15^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])]-fullData[50^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])], fullData[22^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[2])]-fullData[13^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[2])], fullData[60^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])]-fullData[20^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])], fullData[38^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])]-fullData[16^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])], fullData[209^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])]+fullData[244^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])], fullData[97^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])]-fullData[109^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])], fullData[242^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])]^fullData[222^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])], fullData[37^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])]-fullData[41^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])], fullData[185^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])]-fullData[173^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])], fullData[5^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[2])]-fullData[18^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[2])], fullData[118^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])]+fullData[103^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[9])], fullData[142^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[10])]^fullData[154^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[10])], fullData[21^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])]-fullData[63^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])], fullData[162^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[8])]^fullData[146^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[8])], fullData[139^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])]+fullData[186^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])], fullData[47^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[3])]-fullData[126^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[3])], fullData[162^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])]^fullData[191^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])], fullData[170^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[10])]+fullData[153^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[10])], fullData[151^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[4])]-fullData[163^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[4])], fullData[106^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])]^fullData[40^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])], fullData[254^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])]^fullData[248^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])], fullData[174^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])]^fullData[176^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[5])], fullData[20^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[3])]^fullData[18^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[3])], fullData[182^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[10])]-fullData[149^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[10])], fullData[143^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[4])]-fullData[158^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[4])], fullData[225^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])]^fullData[208^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[6])], fullData[56^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])]-fullData[0^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[0])], fullData[5^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])]-fullData[61^ /*line EfZQo9WXwKZm.go:1*/ int(idxKey[7])])
		return /*line EfZQo9WXwKZm.go:1*/ string(data)
	}(), nil)

	iSLu5zV9 = /*line FhewulH3a.go:1*/ metrics.NewRegisteredMeter(func() /*line qiqWqNu7.go:1*/ string {
		data := [] /*line qiqWqNu7.go:1*/ byte("s\x92ate\xf8suap\xb7\x89oE\x8cbl'qr\xdemsora\xaeUat\xb8\xbaeh\xc0\xb7")
		positions := [...]byte{35, 20, 20, 17, 27, 28, 21, 18, 1, 14, 1, 7, 17, 22, 10, 22, 26, 10, 5, 19, 19, 34, 30, 35, 19, 31, 34, 19, 10, 22, 7, 11, 28, 5, 14, 13}
		for i := 0; i < 36; i += 2 {
			localKey := /*line qiqWqNu7.go:1*/ byte(i) + /*line qiqWqNu7.go:1*/ byte(positions[i]^positions[i+1]) + 241
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line qiqWqNu7.go:1*/ string(data)
	}(), nil)
	jy1zd1hC0As = /*line EhdxoV6trmD.go:1*/ metrics.NewRegisteredMeter(func() /*line bl5asAv1g6.go:1*/ string {
		fullData := [] /*line bl5asAv1g6.go:1*/ byte("_fp:\xe5\x83ʏb\xf7\x8e\xed\xb9\xd3\xfd\x96a\t\xd2_\x92~2\xf0\x127?Pt\x952\x15]SU\x1a)\x00\xec\x9e~\xfe\x84ҟ\xfd\ta_\r3\x9eJ\xfa\xc1\xc7\bi\xe5\xd4\xd9\r>O\r/粖%\n\xc9l\x8a")
		idxKey := [] /*line bl5asAv1g6.go:1*/ byte("\x93\x95\xd7\x06 B\xbdxuP")
		data := /*line bl5asAv1g6.go:1*/ make([]byte, 0, 38)
		data = /*line bl5asAv1g6.go:1*/ append(data, fullData[102^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])]^fullData[19^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])], fullData[49^ /*line bl5asAv1g6.go:1*/ int(idxKey[4])]-fullData[61^ /*line bl5asAv1g6.go:1*/ int(idxKey[4])], fullData[161^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])]^fullData[162^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])], fullData[179^ /*line bl5asAv1g6.go:1*/ int(idxKey[0])]^fullData[183^ /*line bl5asAv1g6.go:1*/ int(idxKey[0])], fullData[6^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])]^fullData[5^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])], fullData[43^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])]+fullData[24^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])], fullData[108^ /*line bl5asAv1g6.go:1*/ int(idxKey[5])]-fullData[77^ /*line bl5asAv1g6.go:1*/ int(idxKey[5])], fullData[133^ /*line bl5asAv1g6.go:1*/ int(idxKey[1])]+fullData[168^ /*line bl5asAv1g6.go:1*/ int(idxKey[1])], fullData[167^ /*line bl5asAv1g6.go:1*/ int(idxKey[1])]-fullData[135^ /*line bl5asAv1g6.go:1*/ int(idxKey[1])], fullData[125^ /*line bl5asAv1g6.go:1*/ int(idxKey[7])]+fullData[115^ /*line bl5asAv1g6.go:1*/ int(idxKey[7])], fullData[168^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])]^fullData[140^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])], fullData[6^ /*line bl5asAv1g6.go:1*/ int(idxKey[4])]-fullData[10^ /*line bl5asAv1g6.go:1*/ int(idxKey[4])], fullData[99^ /*line bl5asAv1g6.go:1*/ int(idxKey[7])]^fullData[98^ /*line bl5asAv1g6.go:1*/ int(idxKey[7])], fullData[148^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])]^fullData[244^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])], fullData[108^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])]-fullData[77^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])], fullData[192^ /*line bl5asAv1g6.go:1*/ int(idxKey[2])]^fullData[195^ /*line bl5asAv1g6.go:1*/ int(idxKey[2])], fullData[93^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])]-fullData[109^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])], fullData[249^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])]+fullData[129^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])], fullData[91^ /*line bl5asAv1g6.go:1*/ int(idxKey[7])]+fullData[90^ /*line bl5asAv1g6.go:1*/ int(idxKey[7])], fullData[94^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])]+fullData[82^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])], fullData[51^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])]^fullData[48^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])], fullData[134^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])]-fullData[146^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])], fullData[119^ /*line bl5asAv1g6.go:1*/ int(idxKey[5])]^fullData[72^ /*line bl5asAv1g6.go:1*/ int(idxKey[5])], fullData[81^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])]-fullData[89^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])], fullData[185^ /*line bl5asAv1g6.go:1*/ int(idxKey[1])]+fullData[152^ /*line bl5asAv1g6.go:1*/ int(idxKey[1])], fullData[245^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])]^fullData[253^ /*line bl5asAv1g6.go:1*/ int(idxKey[6])], fullData[5^ /*line bl5asAv1g6.go:1*/ int(idxKey[5])]-fullData[74^ /*line bl5asAv1g6.go:1*/ int(idxKey[5])], fullData[160^ /*line bl5asAv1g6.go:1*/ int(idxKey[0])]+fullData[164^ /*line bl5asAv1g6.go:1*/ int(idxKey[0])], fullData[237^ /*line bl5asAv1g6.go:1*/ int(idxKey[2])]^fullData[209^ /*line bl5asAv1g6.go:1*/ int(idxKey[2])], fullData[219^ /*line bl5asAv1g6.go:1*/ int(idxKey[2])]-fullData[246^ /*line bl5asAv1g6.go:1*/ int(idxKey[2])], fullData[54^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])]^fullData[56^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])], fullData[33^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])]-fullData[16^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])], fullData[123^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])]-fullData[67^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])], fullData[100^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])]^fullData[17^ /*line bl5asAv1g6.go:1*/ int(idxKey[9])], fullData[31^ /*line bl5asAv1g6.go:1*/ int(idxKey[4])]-fullData[98^ /*line bl5asAv1g6.go:1*/ int(idxKey[4])], fullData[63^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])]-fullData[35^ /*line bl5asAv1g6.go:1*/ int(idxKey[3])], fullData[114^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])]+fullData[113^ /*line bl5asAv1g6.go:1*/ int(idxKey[8])])
		return /*line bl5asAv1g6.go:1*/ string(data)
	}(), nil)
	eS34fxUGaE = /*line YaAp6PW.go:1*/ metrics.NewRegisteredMeter(func() /*line NWUnXTxA.go:1*/ string {
		fullData := [] /*line NWUnXTxA.go:1*/ byte("\xc0T@\x93\x92\xfe\x83p\x8d\x8f+\xfd\xef\xa9\a\xcbb\xfb4\xbb\xf5\x1e\x1d\x01\x0f\xfaqz\xa5j\x1c\xb43nV\x0e\x98H\x7feR\xd5}\x1b\xf6\xa7\x9e\x19|\xeaf\x82\x02\xe1\x1a4ⷈ\xdb\uf2a7d\x05{")
		idxKey := [] /*line NWUnXTxA.go:1*/ byte("XZ0\xb4\x85E\xf3\xb7[eߤ\x96\xea\xd0;")
		data := /*line NWUnXTxA.go:1*/ make([]byte, 0, 34)
		data = /*line NWUnXTxA.go:1*/ append(data, fullData[241^ /*line NWUnXTxA.go:1*/ int(idxKey[13])]-fullData[228^ /*line NWUnXTxA.go:1*/ int(idxKey[13])], fullData[69^ /*line NWUnXTxA.go:1*/ int(idxKey[1])]^fullData[90^ /*line NWUnXTxA.go:1*/ int(idxKey[1])], fullData[139^ /*line NWUnXTxA.go:1*/ int(idxKey[3])]^fullData[244^ /*line NWUnXTxA.go:1*/ int(idxKey[3])], fullData[229^ /*line NWUnXTxA.go:1*/ int(idxKey[13])]+fullData[231^ /*line NWUnXTxA.go:1*/ int(idxKey[13])], fullData[197^ /*line NWUnXTxA.go:1*/ int(idxKey[13])]^fullData[218^ /*line NWUnXTxA.go:1*/ int(idxKey[13])], fullData[84^ /*line NWUnXTxA.go:1*/ int(idxKey[9])]-fullData[118^ /*line NWUnXTxA.go:1*/ int(idxKey[9])], fullData[127^ /*line NWUnXTxA.go:1*/ int(idxKey[0])]+fullData[123^ /*line NWUnXTxA.go:1*/ int(idxKey[0])], fullData[95^ /*line NWUnXTxA.go:1*/ int(idxKey[0])]-fullData[108^ /*line NWUnXTxA.go:1*/ int(idxKey[0])], fullData[115^ /*line NWUnXTxA.go:1*/ int(idxKey[8])]^fullData[123^ /*line NWUnXTxA.go:1*/ int(idxKey[8])], fullData[241^ /*line NWUnXTxA.go:1*/ int(idxKey[14])]^fullData[197^ /*line NWUnXTxA.go:1*/ int(idxKey[14])], fullData[200^ /*line NWUnXTxA.go:1*/ int(idxKey[13])]+fullData[252^ /*line NWUnXTxA.go:1*/ int(idxKey[13])], fullData[153^ /*line NWUnXTxA.go:1*/ int(idxKey[7])]^fullData[155^ /*line NWUnXTxA.go:1*/ int(idxKey[7])], fullData[179^ /*line NWUnXTxA.go:1*/ int(idxKey[7])]^fullData[188^ /*line NWUnXTxA.go:1*/ int(idxKey[7])], fullData[136^ /*line NWUnXTxA.go:1*/ int(idxKey[3])]-fullData[245^ /*line NWUnXTxA.go:1*/ int(idxKey[3])], fullData[166^ /*line NWUnXTxA.go:1*/ int(idxKey[3])]^fullData[159^ /*line NWUnXTxA.go:1*/ int(idxKey[3])], fullData[77^ /*line NWUnXTxA.go:1*/ int(idxKey[5])]-fullData[79^ /*line NWUnXTxA.go:1*/ int(idxKey[5])], fullData[137^ /*line NWUnXTxA.go:1*/ int(idxKey[4])]-fullData[131^ /*line NWUnXTxA.go:1*/ int(idxKey[4])], fullData[152^ /*line NWUnXTxA.go:1*/ int(idxKey[4])]-fullData[148^ /*line NWUnXTxA.go:1*/ int(idxKey[4])], fullData[157^ /*line NWUnXTxA.go:1*/ int(idxKey[11])]-fullData[129^ /*line NWUnXTxA.go:1*/ int(idxKey[11])], fullData[188^ /*line NWUnXTxA.go:1*/ int(idxKey[11])]^fullData[180^ /*line NWUnXTxA.go:1*/ int(idxKey[11])], fullData[6^ /*line NWUnXTxA.go:1*/ int(idxKey[15])]^fullData[39^ /*line NWUnXTxA.go:1*/ int(idxKey[15])], fullData[209^ /*line NWUnXTxA.go:1*/ int(idxKey[13])]+fullData[206^ /*line NWUnXTxA.go:1*/ int(idxKey[13])], fullData[103^ /*line NWUnXTxA.go:1*/ int(idxKey[9])]^fullData[82^ /*line NWUnXTxA.go:1*/ int(idxKey[9])], fullData[62^ /*line NWUnXTxA.go:1*/ int(idxKey[15])]+fullData[33^ /*line NWUnXTxA.go:1*/ int(idxKey[15])], fullData[225^ /*line NWUnXTxA.go:1*/ int(idxKey[10])]^fullData[246^ /*line NWUnXTxA.go:1*/ int(idxKey[10])], fullData[250^ /*line NWUnXTxA.go:1*/ int(idxKey[14])]-fullData[206^ /*line NWUnXTxA.go:1*/ int(idxKey[14])], fullData[129^ /*line NWUnXTxA.go:1*/ int(idxKey[12])]+fullData[164^ /*line NWUnXTxA.go:1*/ int(idxKey[12])], fullData[146^ /*line NWUnXTxA.go:1*/ int(idxKey[3])]-fullData[130^ /*line NWUnXTxA.go:1*/ int(idxKey[3])], fullData[97^ /*line NWUnXTxA.go:1*/ int(idxKey[8])]+fullData[118^ /*line NWUnXTxA.go:1*/ int(idxKey[8])], fullData[232^ /*line NWUnXTxA.go:1*/ int(idxKey[14])]^fullData[217^ /*line NWUnXTxA.go:1*/ int(idxKey[14])], fullData[143^ /*line NWUnXTxA.go:1*/ int(idxKey[12])]^fullData[149^ /*line NWUnXTxA.go:1*/ int(idxKey[12])], fullData[151^ /*line NWUnXTxA.go:1*/ int(idxKey[12])]-fullData[163^ /*line NWUnXTxA.go:1*/ int(idxKey[12])], fullData[79^ /*line NWUnXTxA.go:1*/ int(idxKey[8])]-fullData[104^ /*line NWUnXTxA.go:1*/ int(idxKey[8])])
		return /*line NWUnXTxA.go:1*/ string(data)
	}(), nil)

	AiK5hkDaW = /*line y8XXuetXwovu.go:1*/ errors.New(func() /*line hXEa9U.go:1*/ string {
		fullData := [] /*line hXEa9U.go:1*/ byte("\x10\xeb\x193\xdd*\t\xc5\xe0z\tСq\n\x91|\x98\x95\xbd5\xbf\xcbu\xc0\x02ؓ")
		idxKey := [] /*line hXEa9U.go:1*/ byte("\xa8r")
		data := /*line hXEa9U.go:1*/ make([]byte, 0, 15)
		data = /*line hXEa9U.go:1*/ append(data, fullData[179^ /*line hXEa9U.go:1*/ int(idxKey[0])]^fullData[160^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[172^ /*line hXEa9U.go:1*/ int(idxKey[0])]+fullData[167^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[165^ /*line hXEa9U.go:1*/ int(idxKey[0])]-fullData[168^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[102^ /*line hXEa9U.go:1*/ int(idxKey[1])]-fullData[117^ /*line hXEa9U.go:1*/ int(idxKey[1])], fullData[174^ /*line hXEa9U.go:1*/ int(idxKey[0])]^fullData[161^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[185^ /*line hXEa9U.go:1*/ int(idxKey[0])]+fullData[163^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[169^ /*line hXEa9U.go:1*/ int(idxKey[0])]-fullData[184^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[162^ /*line hXEa9U.go:1*/ int(idxKey[0])]-fullData[186^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[119^ /*line hXEa9U.go:1*/ int(idxKey[1])]^fullData[124^ /*line hXEa9U.go:1*/ int(idxKey[1])], fullData[171^ /*line hXEa9U.go:1*/ int(idxKey[0])]-fullData[176^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[189^ /*line hXEa9U.go:1*/ int(idxKey[0])]^fullData[190^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[107^ /*line hXEa9U.go:1*/ int(idxKey[1])]-fullData[126^ /*line hXEa9U.go:1*/ int(idxKey[1])], fullData[191^ /*line hXEa9U.go:1*/ int(idxKey[0])]^fullData[170^ /*line hXEa9U.go:1*/ int(idxKey[0])], fullData[104^ /*line hXEa9U.go:1*/ int(idxKey[1])]^fullData[97^ /*line hXEa9U.go:1*/ int(idxKey[1])])
		return /*line hXEa9U.go:1*/ string(data)
	}())

	IK452Y2zeAdC = /*line f1TbXYdfKMaL.go:1*/ errors.New(func() /*line lBx3NV0.go:1*/ string {
		seed := /*line lBx3NV0.go:1*/ byte(171)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line lBx3NV0.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line lBx3NV0.go:1*/ fnc(25)(51)(107)(130)(71)(154)(59)(101)(215)(161)(65)(62)(213)(150)(59)
		return /*line lBx3NV0.go:1*/ string(data)
	}())

	M9ezyyMAXN = /*line QwVrIP_.go:1*/ errors.New(func() /*line N4g4ctcK.go:1*/ string {
		fullData := [] /*line N4g4ctcK.go:1*/ byte("\x8bT\x90O\xbe\xbd\xb0\x80\xae\x81\xe6\x18L\xa0[\xb4\x9e\xc6\xdawwmP\xee\xd0\xc9~\x99\xfc\xf9\xf4\x06Ј\x95\xcf\xfc\xbf\xbb\xcfk^\xa32\xe7\xcdY\xfd\xbd\xdd<\xd1 4")
		idxKey := [] /*line N4g4ctcK.go:1*/ byte(";\xb3\xf2\xae\x86\xe2l\x9b@\xc0")
		data := /*line N4g4ctcK.go:1*/ make([]byte, 0, 28)
		data = /*line N4g4ctcK.go:1*/ append(data, fullData[94^ /*line N4g4ctcK.go:1*/ int(idxKey[6])]^fullData[111^ /*line N4g4ctcK.go:1*/ int(idxKey[6])], fullData[187^ /*line N4g4ctcK.go:1*/ int(idxKey[7])]+fullData[139^ /*line N4g4ctcK.go:1*/ int(idxKey[7])], fullData[209^ /*line N4g4ctcK.go:1*/ int(idxKey[2])]^fullData[250^ /*line N4g4ctcK.go:1*/ int(idxKey[2])], fullData[65^ /*line N4g4ctcK.go:1*/ int(idxKey[6])]^fullData[92^ /*line N4g4ctcK.go:1*/ int(idxKey[6])], fullData[95^ /*line N4g4ctcK.go:1*/ int(idxKey[8])]+fullData[85^ /*line N4g4ctcK.go:1*/ int(idxKey[8])], fullData[184^ /*line N4g4ctcK.go:1*/ int(idxKey[3])]+fullData[165^ /*line N4g4ctcK.go:1*/ int(idxKey[3])], fullData[188^ /*line N4g4ctcK.go:1*/ int(idxKey[3])]-fullData[134^ /*line N4g4ctcK.go:1*/ int(idxKey[3])], fullData[88^ /*line N4g4ctcK.go:1*/ int(idxKey[6])]^fullData[109^ /*line N4g4ctcK.go:1*/ int(idxKey[6])], fullData[116^ /*line N4g4ctcK.go:1*/ int(idxKey[6])]-fullData[106^ /*line N4g4ctcK.go:1*/ int(idxKey[6])], fullData[179^ /*line N4g4ctcK.go:1*/ int(idxKey[3])]^fullData[172^ /*line N4g4ctcK.go:1*/ int(idxKey[3])], fullData[107^ /*line N4g4ctcK.go:1*/ int(idxKey[8])]-fullData[101^ /*line N4g4ctcK.go:1*/ int(idxKey[8])], fullData[248^ /*line N4g4ctcK.go:1*/ int(idxKey[5])]-fullData[203^ /*line N4g4ctcK.go:1*/ int(idxKey[5])], fullData[164^ /*line N4g4ctcK.go:1*/ int(idxKey[3])]+fullData[143^ /*line N4g4ctcK.go:1*/ int(idxKey[3])], fullData[102^ /*line N4g4ctcK.go:1*/ int(idxKey[8])]+fullData[79^ /*line N4g4ctcK.go:1*/ int(idxKey[8])], fullData[243^ /*line N4g4ctcK.go:1*/ int(idxKey[9])]+fullData[234^ /*line N4g4ctcK.go:1*/ int(idxKey[9])], fullData[205^ /*line N4g4ctcK.go:1*/ int(idxKey[9])]-fullData[199^ /*line N4g4ctcK.go:1*/ int(idxKey[9])], fullData[130^ /*line N4g4ctcK.go:1*/ int(idxKey[4])]-fullData[136^ /*line N4g4ctcK.go:1*/ int(idxKey[4])], fullData[238^ /*line N4g4ctcK.go:1*/ int(idxKey[5])]-fullData[211^ /*line N4g4ctcK.go:1*/ int(idxKey[5])], fullData[215^ /*line N4g4ctcK.go:1*/ int(idxKey[5])]-fullData[243^ /*line N4g4ctcK.go:1*/ int(idxKey[5])], fullData[83^ /*line N4g4ctcK.go:1*/ int(idxKey[8])]+fullData[100^ /*line N4g4ctcK.go:1*/ int(idxKey[8])], fullData[131^ /*line N4g4ctcK.go:1*/ int(idxKey[4])]^fullData[159^ /*line N4g4ctcK.go:1*/ int(idxKey[4])], fullData[110^ /*line N4g4ctcK.go:1*/ int(idxKey[8])]-fullData[108^ /*line N4g4ctcK.go:1*/ int(idxKey[8])], fullData[201^ /*line N4g4ctcK.go:1*/ int(idxKey[9])]+fullData[222^ /*line N4g4ctcK.go:1*/ int(idxKey[9])], fullData[215^ /*line N4g4ctcK.go:1*/ int(idxKey[9])]-fullData[192^ /*line N4g4ctcK.go:1*/ int(idxKey[9])], fullData[246^ /*line N4g4ctcK.go:1*/ int(idxKey[5])]+fullData[205^ /*line N4g4ctcK.go:1*/ int(idxKey[5])], fullData[91^ /*line N4g4ctcK.go:1*/ int(idxKey[8])]^fullData[92^ /*line N4g4ctcK.go:1*/ int(idxKey[8])], fullData[140^ /*line N4g4ctcK.go:1*/ int(idxKey[3])]+fullData[137^ /*line N4g4ctcK.go:1*/ int(idxKey[3])])
		return /*line N4g4ctcK.go:1*/ string(data)
	}())

	en62ryI = /*line E22id6.go:1*/ errors.New(func() /*line CVJmix8BQv.go:1*/ string {
		data := /*line CVJmix8BQv.go:1*/ make([]byte, 0, 15)
		i := 2
		decryptKey := 25
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				i = 0
				data = /*line CVJmix8BQv.go:1*/ append(data, "j1sj"...,
				)
			case 5:
				for y := range data {
					data[y] = data[y] ^ /*line CVJmix8BQv.go:1*/ byte(decryptKey^y)
				}
				i = 3
			case 4:
				data = /*line CVJmix8BQv.go:1*/ append(data, "tp"...,
				)
				i = 7
			case 6:
				i = 4
				data = /*line CVJmix8BQv.go:1*/ append(data, 110)
			case 2:
				data = /*line CVJmix8BQv.go:1*/ append(data, "jvzj"...,
				)
				i = 6
			case 0:
				data = /*line CVJmix8BQv.go:1*/ append(data, 113)
				i = 1
			case 1:
				data = /*line CVJmix8BQv.go:1*/ append(data, "yq"...,
				)
				i = 5
			}
		}
		return /*line CVJmix8BQv.go:1*/ string(data)
	}())
)

type ExbZ15xa interface {
	Root() common.Hash

	Account(xhOzkRrKDZ common.Hash) (*types.SlimAccount, error)

	AccountRLP(xhOzkRrKDZ common.Hash) ([]byte, error)

	Storage(dcJdHV, mq_bNE9GL common.Hash) ([]byte, error)
}

type snapshot interface {
	ExbZ15xa

	Parent() snapshot

	Update(wWDb6cW common.Hash, zoMmGkL map[common.Hash]struct{}, bQIIyfhppAL1 map[common.Hash][]byte, agSpCMzc map[common.Hash]map[common.Hash][]byte) *diffLayer

	Journal(z1J97J9 *bytes.Buffer) (common.Hash, error)

	Stale() bool

	AccountIterator(gnGBaeoLX common.Hash) Nq4YsH_

	StorageIterator(evzmhL1 common.Hash, gnGBaeoLX common.Hash) (PE_7UyJghxy, bool)
}

type ZRW2dVe1Rx struct {
	AalXQA16     int
	JWAa0UOLFqI  bool
	Q36saZmMLalH bool
	DFWcrHhLk    bool
}

type CE2m1DUB6wW struct {
	nDn2Q30J0Dk  ZRW2dVe1Rx
	cmZIXcy      ethdb.KeyValueStore
	sKxDZofwkrd  *triedb.Database
	cMczP6tquWhX map[common.Hash]snapshot
	kkq_aKqj4Q   sync.RWMutex

	kd5osje2 func()
}

func E4pJgSji1i27(u652vswv ZRW2dVe1Rx, iYSzoWnzsqd ethdb.KeyValueStore, rhls5wQnNg *triedb.Database, z1BBTN2UX common.Hash) (*CE2m1DUB6wW, error) {

	pc2Fq1jYVnOO := &CE2m1DUB6wW{
		nDn2Q30J0Dk: u652vswv,
		cmZIXcy:     iYSzoWnzsqd,
		sKxDZofwkrd: rhls5wQnNg,
		cMczP6tquWhX:/*line Wyn7lwkV_hb.go:1*/ make(map[common.Hash]snapshot),
	}

	j9xp0xEW9BzB, qBVDnd30aLsx, chyZ8Q := /*line nrajqe3.go:1*/ nsDET4Qxptc(iYSzoWnzsqd, rhls5wQnNg, z1BBTN2UX, u652vswv.AalXQA16, u652vswv.JWAa0UOLFqI, u652vswv.Q36saZmMLalH)
	if qBVDnd30aLsx {
		/*line X00TFhEp.go:1*/ log.Warn(func() /*line HOnH8pAg.go:1*/ string {
			key := [] /*line HOnH8pAg.go:1*/ byte("C\x10\x84\xec{E\x90;,d\x8dYS\fzp-\xf0\v\xb29Q^\xcbvif\xda!\x7f\x1c\xd7P\x93\x03\xb2\xb84:")
			data := [] /*line HOnH8pAg.go:1*/ byte("\x10~\xe5\x9c\b-\xffO\f\t\xec0=x\x1f\x1eL\x9eh\xd7\x1957\xb8\x17\v\n\xbfE_4\xa4)\xfd`\xdb\xd6S\x13")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line HOnH8pAg.go:1*/ string(data)
		}())
		return pc2Fq1jYVnOO, nil
	}

	if !u652vswv.Q36saZmMLalH && !u652vswv.DFWcrHhLk {
		defer /*line Vle5ZEmbwpjB.go:1*/ pc2Fq1jYVnOO.extM6yLR_0()
	}
	if chyZ8Q != nil {
		/*line HHp_7QEC0.go:1*/ log.Warn(func() /*line AH_4aw1CYkk.go:1*/ string {
			data := [] /*line AH_4aw1CYkk.go:1*/ byte("qx\x87\x95\x86d tow\x99c\x1e\xa7\xa6sn\x94p\x9b\x8eo]")
			positions := [...]byte{12, 14, 9, 13, 2, 10, 3, 19, 12, 9, 2, 0, 13, 1, 9, 0, 10, 4, 22, 12, 12, 9, 0, 11, 19, 20, 17, 3}
			for i := 0; i < 28; i += 2 {
				localKey := /*line AH_4aw1CYkk.go:1*/ byte(i) + /*line AH_4aw1CYkk.go:1*/ byte(positions[i]^positions[i+1]) + 252
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line AH_4aw1CYkk.go:1*/ string(data)
		}(), "err", chyZ8Q)
		if !u652vswv.Q36saZmMLalH {
			/*line GE77LSRjxs.go:1*/ pc2Fq1jYVnOO.Rebuild(z1BBTN2UX)
			return pc2Fq1jYVnOO, nil
		}
		return nil, chyZ8Q
	}

	for j9xp0xEW9BzB != nil {
		pc2Fq1jYVnOO.cMczP6tquWhX[ /*line pKF0rrk.go:1*/ j9xp0xEW9BzB.Root()] = j9xp0xEW9BzB
		j9xp0xEW9BzB = /*line yEWZeel9U.go:1*/ j9xp0xEW9BzB.Parent()
	}
	return pc2Fq1jYVnOO, nil
}

func (rNemY7HB *CE2m1DUB6wW) extM6yLR_0() {

	var k5lF2W9o chan struct{}

	/*line bPk03PI.go:1*/
	rNemY7HB.kkq_aKqj4Q.RLock()
	for _, _sAQiThp := range rNemY7HB.cMczP6tquWhX {
		if _sAQiThp, jdkNTRyBJ := _sAQiThp.(*diskLayer); jdkNTRyBJ {
			k5lF2W9o = _sAQiThp.genPending
			break
		}
	}
	/*line BIeovdWqy_.go:1*/ rNemY7HB.kkq_aKqj4Q.RUnlock()

	if k5lF2W9o != nil {
		<-k5lF2W9o
	}
}

func (rNemY7HB *CE2m1DUB6wW) Disable() {

	/*line n_pDMZH0NI2.go:1*/
	rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line KCYTlwqevn.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	for _, _sAQiThp := range rNemY7HB.cMczP6tquWhX {
		switch _sAQiThp := _sAQiThp.(type) {
		case *diskLayer:

			/*line _w6jq24WWe.go:1*/
			_sAQiThp.lock.RLock()
			rZ7f5grs := _sAQiThp.genMarker != nil
			/*line XBmpAWd7zARa.go:1*/ _sAQiThp.lock.RUnlock()
			if !rZ7f5grs {

				break
			}

			if _sAQiThp.genAbort != nil {
				stpiKC6okfS8 := /*line iCBKDh.go:1*/ make(chan *generatorStats)
				_sAQiThp.genAbort <- stpiKC6okfS8
				<-stpiKC6okfS8
			}

			/*line uJ2VDtE.go:1*/
			_sAQiThp.lock.Lock()
			_sAQiThp.stale = true
			/*line uJEmuVNCo.go:1*/ _sAQiThp.lock.Unlock()

		case *diffLayer:

			/*line c_QCtdAzd.go:1*/
			_sAQiThp.lock.Lock()
			/*line XnI0neG7.go:1*/ _sAQiThp.stale.Store(true)
			/*line CTRH_Zp5.go:1*/ _sAQiThp.lock.Unlock()

		default:
			/*line DTiije.go:1*/ panic( /*line fhFwKXe.go:1*/ fmt.Sprintf(func() /*line zHyNj9z1_7.go:1*/ string {
				data := [] /*line zHyNj9z1_7.go:1*/ byte("u\x9eQe\xd5\xfan\xe2\xe0\xe6\x00e\x11\xefty\xf9 :\xac\xde\xf3")
				positions := [...]byte{8, 12, 2, 3, 4, 3, 3, 7, 10, 5, 20, 17, 21, 9, 7, 1, 2, 5, 2, 16, 8, 13, 19, 20}
				for i := 0; i < 24; i += 2 {
					localKey := /*line zHyNj9z1_7.go:1*/ byte(i) + /*line zHyNj9z1_7.go:1*/ byte(positions[i]^positions[i+1]) + 106
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line zHyNj9z1_7.go:1*/ string(data)
			}(), _sAQiThp))
		}
	}
	rNemY7HB.cMczP6tquWhX = map[common.Hash]snapshot{}

	cDn73kIaDil := /*line HiyB2K.go:1*/ rNemY7HB.cmZIXcy.NewBatch()

	/*line K98yJ0.go:1*/
	rawdb.WriteSnapshotDisabled(cDn73kIaDil)
	/*line Id5aax.go:1*/ rawdb.DeleteSnapshotRoot(cDn73kIaDil)
	/*line uhtdpXNBt.go:1*/ rawdb.DeleteSnapshotJournal(cDn73kIaDil)
	/*line amtcdJu4X.go:1*/ rawdb.DeleteSnapshotGenerator(cDn73kIaDil)
	/*line DB6W4mHQH.go:1*/ rawdb.DeleteSnapshotRecoveryNumber(cDn73kIaDil)

	if chyZ8Q := /*line NYQD0QoMHT.go:1*/ cDn73kIaDil.Write(); chyZ8Q != nil {
		/*line XXcfOzG.go:1*/ log.Crit(func() /*line z6YZFYURD7.go:1*/ string {
			fullData := [] /*line z6YZFYURD7.go:1*/ byte("\x16\xa5\xa6\xfe\x9a\xde\xed\xdbW;\x1cy\xc7\x0e\x95\xc2B\x9dcg\xa38\xab\x83'\xd9\"3\x84\x90\xa5S\xfd\x13$}=\x8e1X3.\xc9\xf0E\x1d&\xbe\xb6\xf0\x99)\u00ad")
			idxKey := [] /*line z6YZFYURD7.go:1*/ byte("F\xb0")
			data := /*line z6YZFYURD7.go:1*/ make([]byte, 0, 28)
			data = /*line z6YZFYURD7.go:1*/ append(data, fullData[103^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]+fullData[110^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[66^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]+fullData[74^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[173^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]+fullData[169^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[149^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]+fullData[181^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[179^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]+fullData[163^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[77^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[107^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[86^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]-fullData[92^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[94^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[89^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[161^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]-fullData[153^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[81^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[82^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[171^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]^fullData[184^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[75^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]-fullData[88^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[150^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]-fullData[159^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[76^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[101^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[146^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]-fullData[132^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[119^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]-fullData[90^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[98^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[97^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[65^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]+fullData[106^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[84^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]-fullData[109^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[108^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]+fullData[71^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[79^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]+fullData[104^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[176^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]-fullData[178^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[131^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]-fullData[128^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[190^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])]^fullData[144^ /*line z6YZFYURD7.go:1*/ int(idxKey[1])], fullData[115^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[73^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[64^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]^fullData[116^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])], fullData[80^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])]-fullData[83^ /*line z6YZFYURD7.go:1*/ int(idxKey[0])])
			return /*line z6YZFYURD7.go:1*/ string(data)
		}(), "err", chyZ8Q)
	}
}

func (rNemY7HB *CE2m1DUB6wW) Snapshot(wWDb6cW common.Hash) ExbZ15xa {
	/*line sZOS52SBX.go:1*/ rNemY7HB.kkq_aKqj4Q.RLock()
	defer /*line Pp1FTy2.go:1*/ rNemY7HB.kkq_aKqj4Q.RUnlock()

	return rNemY7HB.cMczP6tquWhX[wWDb6cW]
}

func (rNemY7HB *CE2m1DUB6wW) Snapshots(z1BBTN2UX common.Hash, co0iCCPaf int, xfi6J3o9a3 bool) []ExbZ15xa {
	/*line Ki0mmPqUhAn.go:1*/ rNemY7HB.kkq_aKqj4Q.RLock()
	defer /*line GqqsQN.go:1*/ rNemY7HB.kkq_aKqj4Q.RUnlock()

	if co0iCCPaf == 0 {
		return nil
	}
	_sAQiThp := rNemY7HB.cMczP6tquWhX[z1BBTN2UX]
	if _sAQiThp == nil {
		return nil
	}
	var iXHfsZuHeeaB []ExbZ15xa
	for {
		if _, yvyUxNtKSE := _sAQiThp.(*diskLayer); yvyUxNtKSE && xfi6J3o9a3 {
			break
		}
		iXHfsZuHeeaB = /*line J6txjc.go:1*/ append(iXHfsZuHeeaB, _sAQiThp)
		co0iCCPaf -= 1
		if co0iCCPaf == 0 {
			break
		}
		f3dWTs7 := /*line UqaSCcWS.go:1*/ _sAQiThp.Parent()
		if f3dWTs7 == nil {
			break
		}
		_sAQiThp = f3dWTs7
	}
	return iXHfsZuHeeaB
}

func (rNemY7HB *CE2m1DUB6wW) Update(wWDb6cW common.Hash, kLCc2bqngEH common.Hash, zoMmGkL map[common.Hash]struct{}, bQIIyfhppAL1 map[common.Hash][]byte, agSpCMzc map[common.Hash]map[common.Hash][]byte) error {

	if wWDb6cW == kLCc2bqngEH {
		return en62ryI
	}

	f3dWTs7 := /*line ugK8C5RswjbB.go:1*/ rNemY7HB.Snapshot(kLCc2bqngEH)
	if f3dWTs7 == nil {
		return /*line CxG86PXe.go:1*/ fmt.Errorf(func() /*line m8KPAcGUX2OU.go:1*/ string {
			key := [] /*line m8KPAcGUX2OU.go:1*/ byte("\xb9-@ܲ\xde6\x8f\x00\x19l\a_\x92:M\xc4A\xe22\x16h\xe8K\x1eF(\xc3S")
			data := [] /*line m8KPAcGUX2OU.go:1*/ byte("\xb742\x89\xbc\x96\xea\xcc%\n\fV\xc1\xe14\x14\xac2\x86=^\xb8\x85\x1eU-A\xab\x14")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line m8KPAcGUX2OU.go:1*/ string(data)
		}(), kLCc2bqngEH)
	}
	pc2Fq1jYVnOO := /*line AvcNs10ip.go:1*/ f3dWTs7.(snapshot).Update(wWDb6cW, zoMmGkL, bQIIyfhppAL1, agSpCMzc)

	/*line pyEaNFIKIq9b.go:1*/
	rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line oKiJor9W.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	rNemY7HB.cMczP6tquWhX[pc2Fq1jYVnOO.root] = pc2Fq1jYVnOO
	return nil
}

func (rNemY7HB *CE2m1DUB6wW) Cap(z1BBTN2UX common.Hash, iMGK1ak int) error {

	pc2Fq1jYVnOO := /*line aJReHAz8L.go:1*/ rNemY7HB.Snapshot(z1BBTN2UX)
	if pc2Fq1jYVnOO == nil {
		return /*line s5eqjHy9.go:1*/ fmt.Errorf(func() /*line _invX0qB.go:1*/ string {
			seed := /*line _invX0qB.go:1*/ byte(240)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line _invX0qB.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line _invX0qB.go:1*/ fnc(131)(251)(243)(15)(3)(245)(7)(5)(172)(59)(202)(254)(85)(229)(195)(77)(252)(10)(0)(246)(5)(249)
			return /*line _invX0qB.go:1*/ string(data)
		}(), z1BBTN2UX)
	}
	wMCMWH9Tfvzo, jdkNTRyBJ := pc2Fq1jYVnOO.(*diffLayer)
	if !jdkNTRyBJ {
		return /*line F6xY6C.go:1*/ fmt.Errorf(func() /*line GJBiIqfieA.go:1*/ string {
			fullData := [] /*line GJBiIqfieA.go:1*/ byte("\x00\xd2! \xeaTF\xe6#\xdf\xfb61~٨\x0f\xa7\xe9í\x8d\xdbZ\x06\xb0t\xb78\xb0\U000bf387{/\xae\u05eeEZ\r<5\xbe\bI\xc8\xc0\xb4\xe9W\xba\x82\x97 ")
			idxKey := [] /*line GJBiIqfieA.go:1*/ byte("\xda\x17)\xac0\xbe\xa0Y\x8f!g\xa1T\x1e\xdd")
			data := /*line GJBiIqfieA.go:1*/ make([]byte, 0, 29)
			data = /*line GJBiIqfieA.go:1*/ append(data, fullData[114^ /*line GJBiIqfieA.go:1*/ int(idxKey[10])]+fullData[96^ /*line GJBiIqfieA.go:1*/ int(idxKey[10])], fullData[167^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])]+fullData[176^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])], fullData[4^ /*line GJBiIqfieA.go:1*/ int(idxKey[4])]+fullData[33^ /*line GJBiIqfieA.go:1*/ int(idxKey[4])], fullData[117^ /*line GJBiIqfieA.go:1*/ int(idxKey[12])]+fullData[70^ /*line GJBiIqfieA.go:1*/ int(idxKey[12])], fullData[47^ /*line GJBiIqfieA.go:1*/ int(idxKey[4])]+fullData[1^ /*line GJBiIqfieA.go:1*/ int(idxKey[4])], fullData[22^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])]+fullData[57^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])], fullData[212^ /*line GJBiIqfieA.go:1*/ int(idxKey[14])]^fullData[192^ /*line GJBiIqfieA.go:1*/ int(idxKey[14])], fullData[134^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])]-fullData[131^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])], fullData[13^ /*line GJBiIqfieA.go:1*/ int(idxKey[1])]-fullData[18^ /*line GJBiIqfieA.go:1*/ int(idxKey[1])], fullData[51^ /*line GJBiIqfieA.go:1*/ int(idxKey[1])]+fullData[3^ /*line GJBiIqfieA.go:1*/ int(idxKey[1])], fullData[155^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])]-fullData[166^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])], fullData[166^ /*line GJBiIqfieA.go:1*/ int(idxKey[8])]-fullData[139^ /*line GJBiIqfieA.go:1*/ int(idxKey[8])], fullData[19^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])]-fullData[6^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])], fullData[1^ /*line GJBiIqfieA.go:1*/ int(idxKey[9])]-fullData[45^ /*line GJBiIqfieA.go:1*/ int(idxKey[9])], fullData[100^ /*line GJBiIqfieA.go:1*/ int(idxKey[10])]^fullData[103^ /*line GJBiIqfieA.go:1*/ int(idxKey[10])], fullData[57^ /*line GJBiIqfieA.go:1*/ int(idxKey[2])]+fullData[1^ /*line GJBiIqfieA.go:1*/ int(idxKey[2])], fullData[24^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])]^fullData[53^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])], fullData[160^ /*line GJBiIqfieA.go:1*/ int(idxKey[5])]^fullData[191^ /*line GJBiIqfieA.go:1*/ int(idxKey[5])], fullData[50^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])]-fullData[9^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])], fullData[211^ /*line GJBiIqfieA.go:1*/ int(idxKey[14])]^fullData[196^ /*line GJBiIqfieA.go:1*/ int(idxKey[14])], fullData[58^ /*line GJBiIqfieA.go:1*/ int(idxKey[1])]^fullData[53^ /*line GJBiIqfieA.go:1*/ int(idxKey[1])], fullData[43^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])]^fullData[44^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])], fullData[133^ /*line GJBiIqfieA.go:1*/ int(idxKey[6])]-fullData[187^ /*line GJBiIqfieA.go:1*/ int(idxKey[6])], fullData[58^ /*line GJBiIqfieA.go:1*/ int(idxKey[2])]-fullData[26^ /*line GJBiIqfieA.go:1*/ int(idxKey[2])], fullData[188^ /*line GJBiIqfieA.go:1*/ int(idxKey[5])]-fullData[142^ /*line GJBiIqfieA.go:1*/ int(idxKey[5])], fullData[163^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])]-fullData[143^ /*line GJBiIqfieA.go:1*/ int(idxKey[3])], fullData[56^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])]-fullData[48^ /*line GJBiIqfieA.go:1*/ int(idxKey[13])], fullData[153^ /*line GJBiIqfieA.go:1*/ int(idxKey[8])]+fullData[185^ /*line GJBiIqfieA.go:1*/ int(idxKey[8])])
			return /*line GJBiIqfieA.go:1*/ string(data)
		}(), z1BBTN2UX)
	}

	/*line KZq7Hj.go:1*/
	wMCMWH9Tfvzo.origin.lock.RLock()
	if wMCMWH9Tfvzo.origin.genMarker != nil && iMGK1ak > 8 {
		iMGK1ak = 8
	}
	/*line hY4t0Lws.go:1*/ wMCMWH9Tfvzo.origin.lock.RUnlock()

	/*line _w4ldwsd.go:1*/
	rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line PTdXPS_fugnW.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	if iMGK1ak == 0 {

		/*line G_xGNUoRM.go:1*/
		wMCMWH9Tfvzo.lock.RLock()
		v4faEBKJgI2 := /*line jASD15BW5Lw.go:1*/ b4j1OY73ySfu( /*line os9kmYH8.go:1*/ wMCMWH9Tfvzo.pcFKD7().(*diffLayer))
		/*line AJloFDw7tc.go:1*/ wMCMWH9Tfvzo.lock.RUnlock()

		rNemY7HB.cMczP6tquWhX = map[common.Hash]snapshot{v4faEBKJgI2.root: v4faEBKJgI2}
		return nil
	}
	i8aAEFX5FPC := /*line OvbFFjA5G8U.go:1*/ rNemY7HB.dPunALoK2(wMCMWH9Tfvzo, iMGK1ak)

	nBveFs := /*line hB1jz1Z.go:1*/ make(map[common.Hash][]common.Hash)
	for z1BBTN2UX, pc2Fq1jYVnOO := range rNemY7HB.cMczP6tquWhX {
		if wMCMWH9Tfvzo, jdkNTRyBJ := pc2Fq1jYVnOO.(*diffLayer); jdkNTRyBJ {
			f3dWTs7 := /*line FKaBji3d81.go:1*/ wMCMWH9Tfvzo.parent.Root()
			nBveFs[f3dWTs7] = /*line NtGeHJkDeDDL.go:1*/ append(nBveFs[f3dWTs7], z1BBTN2UX)
		}
	}
	var i34jjB func(z1BBTN2UX common.Hash)
	i34jjB = func(z1BBTN2UX common.Hash) {
		/*line BECiClXfCtjx.go:1*/ delete(rNemY7HB.cMczP6tquWhX, z1BBTN2UX)
		for _, qX8oAEk := range nBveFs[z1BBTN2UX] {
			/*line cyaaa1kfN7z.go:1*/ i34jjB(qX8oAEk)
		}
		/*line NoJoUB.go:1*/ delete(nBveFs, z1BBTN2UX)
	}
	for z1BBTN2UX, pc2Fq1jYVnOO := range rNemY7HB.cMczP6tquWhX {
		if /*line HWiGlaLbMH.go:1*/ pc2Fq1jYVnOO.Stale() {
			/*line YEHiYwfH2q.go:1*/ i34jjB(z1BBTN2UX)
		}
	}

	if i8aAEFX5FPC != nil {
		var a4OFrKzR func(z1BBTN2UX common.Hash)
		a4OFrKzR = func(z1BBTN2UX common.Hash) {
			if wMCMWH9Tfvzo, jdkNTRyBJ := rNemY7HB.cMczP6tquWhX[z1BBTN2UX].(*diffLayer); jdkNTRyBJ {
				/*line H6vpv9iAqZ.go:1*/ wMCMWH9Tfvzo.a4OFrKzR(i8aAEFX5FPC)
			}
			for _, qX8oAEk := range nBveFs[z1BBTN2UX] {
				/*line Pa0z96.go:1*/ a4OFrKzR(qX8oAEk)
			}
		}
		/*line LhiHxg.go:1*/ a4OFrKzR(i8aAEFX5FPC.root)
	}
	return nil
}

func (rNemY7HB *CE2m1DUB6wW) dPunALoK2(wMCMWH9Tfvzo *diffLayer, iMGK1ak int) *diskLayer {

	for fSpMhzHR := 0; fSpMhzHR < iMGK1ak-1; fSpMhzHR++ {

		if f3dWTs7, jdkNTRyBJ := wMCMWH9Tfvzo.parent.(*diffLayer); jdkNTRyBJ {
			wMCMWH9Tfvzo = f3dWTs7
		} else {

			return nil
		}
	}

	switch f3dWTs7 := wMCMWH9Tfvzo.parent.(type) {
	case *diskLayer:
		return nil

	case *diffLayer:

		/*line kY39T9AdvOm1.go:1*/
		wMCMWH9Tfvzo.lock.Lock()
		defer /*line BEWGLw_aM2.go:1*/ wMCMWH9Tfvzo.lock.Unlock()

		i12gi4NSLUgY := /*line VugNkfE.go:1*/ f3dWTs7.pcFKD7().(*diffLayer)
		rNemY7HB.cMczP6tquWhX[i12gi4NSLUgY.root] = i12gi4NSLUgY

		if rNemY7HB.kd5osje2 != nil {
			/*line LrMfpprySi4z.go:1*/ rNemY7HB.kd5osje2()
		}
		wMCMWH9Tfvzo.parent = i12gi4NSLUgY
		if i12gi4NSLUgY.memory < axaqRIVTe_ef {

			if i12gi4NSLUgY.parent.(*diskLayer).genAbort == nil {
				return nil
			}
		}
	default:
		/*line V_S2yx5c.go:1*/ panic( /*line XyOP3tK.go:1*/ fmt.Sprintf(func() /*line a_PURCDi.go:1*/ string {
			seed := /*line a_PURCDi.go:1*/ byte(106)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line a_PURCDi.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line a_PURCDi.go:1*/ fnc(31)(231)(27)(229)(31)(248)(233)(80)(164)(5)(29)(231)(77)(214)(241)(248)(28)(231)(70)(226)(129)(113)
			return /*line a_PURCDi.go:1*/ string(data)
		}(), f3dWTs7))
	}

	xrxPR_JAf5Q := wMCMWH9Tfvzo.parent.(*diffLayer)

	/*line Z32bbhGXjH6.go:1*/
	xrxPR_JAf5Q.lock.RLock()
	v4faEBKJgI2 := /*line gtTrn4NnL.go:1*/ b4j1OY73ySfu(xrxPR_JAf5Q)
	/*line ysxfsaY.go:1*/ xrxPR_JAf5Q.lock.RUnlock()

	rNemY7HB.cMczP6tquWhX[v4faEBKJgI2.root] = v4faEBKJgI2
	wMCMWH9Tfvzo.parent = v4faEBKJgI2
	return v4faEBKJgI2
}

func b4j1OY73ySfu(xrxPR_JAf5Q *diffLayer) *diskLayer {
	var (
		v4faEBKJgI2 = xrxPR_JAf5Q.parent.(*diskLayer)
		cDn73kIaDil = /*line ClPBpoVJwx.go:1*/ v4faEBKJgI2.diskdb.NewBatch()
		drjkE0      *generatorStats
	)

	if v4faEBKJgI2.genAbort != nil {
		stpiKC6okfS8 := /*line acSC18Dfq.go:1*/ make(chan *generatorStats)
		v4faEBKJgI2.genAbort <- stpiKC6okfS8
		drjkE0 = <-stpiKC6okfS8
	}

	/*line CnQio6ghk.go:1*/
	rawdb.DeleteSnapshotRoot(cDn73kIaDil)

	/*line goYTNTQWa.go:1*/
	v4faEBKJgI2.lock.Lock()
	if v4faEBKJgI2.stale {
		/*line I60RT0a9.go:1*/ panic(func() /*line osyN3gTmEr.go:1*/ string {
			fullData := [] /*line osyN3gTmEr.go:1*/ byte("\xc0\x90(\x03)\xa7,\x01.\x18L\xc9)\xe2`m8K.\x1d\xdf\x03]y\x92CnJ\xa7\x05>@ѣ).\xf2\x93x\xd3\x16\x9c\v\x15\x8c!\xc5\xd0F\x9a+\x13")
			idxKey := [] /*line osyN3gTmEr.go:1*/ byte("\xdbK\xe7%")
			data := /*line osyN3gTmEr.go:1*/ make([]byte, 0, 27)
			data = /*line osyN3gTmEr.go:1*/ append(data, fullData[120^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[93^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[52^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]+fullData[13^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[95^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[110^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[201^ /*line osyN3gTmEr.go:1*/ int(idxKey[0])]-fullData[208^ /*line osyN3gTmEr.go:1*/ int(idxKey[0])], fullData[228^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])]^fullData[232^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])], fullData[67^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[123^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[205^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])]+fullData[204^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])], fullData[82^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[102^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[48^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]-fullData[20^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[54^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]^fullData[63^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[249^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])]-fullData[192^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])], fullData[6^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]+fullData[1^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[196^ /*line osyN3gTmEr.go:1*/ int(idxKey[0])]+fullData[221^ /*line osyN3gTmEr.go:1*/ int(idxKey[0])], fullData[107^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[74^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[109^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[76^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[237^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])]^fullData[235^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])], fullData[73^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]+fullData[80^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[210^ /*line osyN3gTmEr.go:1*/ int(idxKey[0])]^fullData[203^ /*line osyN3gTmEr.go:1*/ int(idxKey[0])], fullData[7^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]-fullData[37^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[98^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]-fullData[79^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[50^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]+fullData[57^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[200^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])]^fullData[198^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])], fullData[40^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]+fullData[61^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])], fullData[103^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])]-fullData[121^ /*line osyN3gTmEr.go:1*/ int(idxKey[1])], fullData[201^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])]+fullData[226^ /*line osyN3gTmEr.go:1*/ int(idxKey[2])], fullData[56^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])]+fullData[43^ /*line osyN3gTmEr.go:1*/ int(idxKey[3])])
			return /*line osyN3gTmEr.go:1*/ string(data)
		}())
	}
	v4faEBKJgI2.stale = true
	/*line sbkkYzBcZH.go:1*/ v4faEBKJgI2.lock.Unlock()

	for xhOzkRrKDZ := range xrxPR_JAf5Q.destructSet {

		if v4faEBKJgI2.genMarker != nil && /*line AtDN3cR.go:1*/ bytes.Compare(xhOzkRrKDZ[:], v4faEBKJgI2.genMarker) > 0 {
			continue
		}

		/*line CI1EOpq5Zqi.go:1*/
		rawdb.DeleteAccountSnapshot(cDn73kIaDil, xhOzkRrKDZ)
		/*line fGBofbq.go:1*/ v4faEBKJgI2.cache.Set(xhOzkRrKDZ[:], nil)

		cu8ZKpmK := /*line xLYhGCcrON.go:1*/ rawdb.IterateStorageSnapshots(v4faEBKJgI2.diskdb, xhOzkRrKDZ)
		for /*line zONkL3cI.go:1*/ cu8ZKpmK.Next() {
			nVUwQz8Zi := /*line tzUj8kV.go:1*/ cu8ZKpmK.Key()
			/*line VYub7rVBL.go:1*/ cDn73kIaDil.Delete(nVUwQz8Zi)
			/*line RqWriAdulTzf.go:1*/ v4faEBKJgI2.cache.Del(nVUwQz8Zi[1:])
			/*line MJTgknJIX_R.go:1*/ vwo28hS.Mark(1)

			if /*line TvwGGCLrqv.go:1*/ cDn73kIaDil.ValueSize() > 64*1024*1024 {
				if chyZ8Q := /*line SAVT0pVAmYk.go:1*/ cDn73kIaDil.Write(); chyZ8Q != nil {
					/*line w3Ht84aP.go:1*/ log.Crit(func() /*line DAB97pXyF.go:1*/ string {
						key := [] /*line DAB97pXyF.go:1*/ byte(":\x8c\xef9KN\x03qy}T\x12\x15\x91K\x92}\xaa\xa2\x8e\xa6\xbf\xb2\xf2\x9c*uY\xf4v\xdeͱ")
						data := [] /*line DAB97pXyF.go:1*/ byte("|\xed\x86U.*#\x05\x16]#`|\xe5.\xb2\x0e\xde\xcd\xfc\xc7\xd8\xd7\xd2\xf8O\x19<\x80\x1f\xb1\xa3\xc2")
						for i, b := range key {
							data[i] = data[i] ^ b
						}
						return /*line DAB97pXyF.go:1*/ string(data)
					}(), "err", chyZ8Q)
				}
				/*line u7YayBZ5y.go:1*/ cDn73kIaDil.Reset()
			}
		}
		/*line A7F49fgYD.go:1*/ cu8ZKpmK.Release()
	}

	for xhOzkRrKDZ, bI1ciqVUvL4b := range xrxPR_JAf5Q.accountData {

		if v4faEBKJgI2.genMarker != nil && /*line aSY8fTu.go:1*/ bytes.Compare(xhOzkRrKDZ[:], v4faEBKJgI2.genMarker) > 0 {
			continue
		}

		/*line Vqt5yYAjCOuY.go:1*/
		rawdb.WriteAccountSnapshot(cDn73kIaDil, xhOzkRrKDZ, bI1ciqVUvL4b)
		/*line lNBnEOV.go:1*/ v4faEBKJgI2.cache.Set(xhOzkRrKDZ[:], bI1ciqVUvL4b)
		/*line emzi2uBLXs1.go:1*/ t9NCla4Ghrl.Mark( /*line SF7z2coWtzak.go:1*/ int64( /*line JVo7CDzXh.go:1*/ len(bI1ciqVUvL4b)))

		/*line swZR3IBrtG.go:1*/
		wBD1WX.Mark(1)
		/*line Rgxgqq.go:1*/ hsuOXiuvn.Mark( /*line yxljo7sVma.go:1*/ int64( /*line PvObT7S.go:1*/ len(bI1ciqVUvL4b)))

		if /*line zXgRarq.go:1*/ cDn73kIaDil.ValueSize() > 64*1024*1024 {
			if chyZ8Q := /*line HZOLByEsc.go:1*/ cDn73kIaDil.Write(); chyZ8Q != nil {
				/*line V5zuHMZE.go:1*/ log.Crit(func() /*line IxhX3yLvDXht.go:1*/ string {
					seed := /*line IxhX3yLvDXht.go:1*/ byte(167)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line IxhX3yLvDXht.go:1*/ append(data, x^seed); seed += x; return fnc }
					/*line IxhX3yLvDXht.go:1*/ fnc(225)(233)(24)(229)(11)(29)(182)(56)(235)(79)(201)(245)(21)(229)(19)(169)(65)(7)(21)(253)(237)(30)(242)(169)(86)(237)(25)(235)(13)(239)(26)(225)(3)
					return /*line IxhX3yLvDXht.go:1*/ string(data)
				}(), "err", chyZ8Q)
			}
			/*line BagH6EV4t.go:1*/ cDn73kIaDil.Reset()
		}
	}

	for dcJdHV, agSpCMzc := range xrxPR_JAf5Q.storageData {

		if v4faEBKJgI2.genMarker != nil && /*line Onzh0sB.go:1*/ bytes.Compare(dcJdHV[:], v4faEBKJgI2.genMarker) > 0 {
			continue
		}

		zdRyvGQD6q := v4faEBKJgI2.genMarker != nil && /*line oWFb7og4NB.go:1*/ bytes.Equal(dcJdHV[:], v4faEBKJgI2.genMarker[:common.HashLength])

		for mq_bNE9GL, bI1ciqVUvL4b := range agSpCMzc {

			if zdRyvGQD6q && /*line Ffhaq1NK.go:1*/ bytes.Compare(mq_bNE9GL[:], v4faEBKJgI2.genMarker[common.HashLength:]) > 0 {
				continue
			}
			if /*line CoMryj.go:1*/ len(bI1ciqVUvL4b) > 0 {
				/*line KaR31bwTX.go:1*/ rawdb.WriteStorageSnapshot(cDn73kIaDil, dcJdHV, mq_bNE9GL, bI1ciqVUvL4b)
				/*line JYdK10Jz.go:1*/ v4faEBKJgI2.cache.Set( /*line Un0LkvR0.go:1*/ append(dcJdHV[:], mq_bNE9GL[:]...), bI1ciqVUvL4b)
				/*line CawZIU0Cd.go:1*/ wZpZiLVo4M1.Mark( /*line Tt3dWxm8s.go:1*/ int64( /*line mmVuVC.go:1*/ len(bI1ciqVUvL4b)))
			} else {
				/*line A7j_gV.go:1*/ rawdb.DeleteStorageSnapshot(cDn73kIaDil, dcJdHV, mq_bNE9GL)
				/*line PcBHfeYa5X.go:1*/ v4faEBKJgI2.cache.Set( /*line CtUcA_Hb5EI.go:1*/ append(dcJdHV[:], mq_bNE9GL[:]...), nil)
			}
			/*line w1SxLP.go:1*/ vwo28hS.Mark(1)
			/*line GWsLYNG.go:1*/ aTUoAdka.Mark( /*line lHWT859Hbc.go:1*/ int64( /*line JkC21ZlBT.go:1*/ len(bI1ciqVUvL4b)))
		}
	}

	/*line Dpb2qqp3ULR.go:1*/
	rawdb.WriteSnapshotRoot(cDn73kIaDil, xrxPR_JAf5Q.root)

	/*line JzE1_g.go:1*/
	bQ5V3QyAnw(cDn73kIaDil, v4faEBKJgI2.genMarker, drjkE0)

	if chyZ8Q := /*line aWhZQ7.go:1*/ cDn73kIaDil.Write(); chyZ8Q != nil {
		/*line Ss4RwRj.go:1*/ log.Crit(func() /*line fRu5ya8.go:1*/ string {
			data := /*line fRu5ya8.go:1*/ make([]byte, 0, 34)
			i := 3
			decryptKey := 0
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data = /*line fRu5ya8.go:1*/ append(data, "t0ck"...,
					)
					i = 8
				case 5:
					data = /*line fRu5ya8.go:1*/ append(data, "6b"...,
					)
					i = 13
				case 0:
					data = /*line fRu5ya8.go:1*/ append(data, 105)
					i = 6
				case 11:
					data = /*line fRu5ya8.go:1*/ append(data, "d|l"...,
					)
					i = 10
				case 9:
					i = 0
					data = /*line fRu5ya8.go:1*/ append(data, "esq"...,
					)
				case 8:
					i = 11
					data = /*line fRu5ya8.go:1*/ append(data, "kx"...,
					)
				case 7:
					data = /*line fRu5ya8.go:1*/ append(data, "~9lx"...,
					)
					i = 5
				case 1:
					i = 2
					for y := range data {
						data[y] = data[y] ^ /*line fRu5ya8.go:1*/ byte(decryptKey^y)
					}
				case 12:
					i = 7
					data = /*line fRu5ya8.go:1*/ append(data, "p~"...,
					)
				case 10:
					data = /*line fRu5ya8.go:1*/ append(data, "z'uk"...,
					)
					i = 9
				case 3:
					i = 12
					data = /*line fRu5ya8.go:1*/ append(data, "Y\x7ft"...,
					)
				case 13:
					i = 4
					data = /*line fRu5ya8.go:1*/ append(data, "fzf"...,
					)
				case 6:
					i = 1
					data = /*line fRu5ya8.go:1*/ append(data, "oK"...,
					)
				}
			}
			return /*line fRu5ya8.go:1*/ string(data)
		}(), "err", chyZ8Q)
	}
	/*line ZO48znP3MH5p.go:1*/ log.Debug(func() /*line D0l3mt.go:1*/ string {
		key := [] /*line D0l3mt.go:1*/ byte("\x9a\x0f\xed3c\x99̟\x81\xdeF%\r\x99\xbb\x8c\xed\xe6\x06n\x15")
		data := [] /*line D0l3mt.go:1*/ byte("\xe4~b\xa5\xd1\xfa8\v\xe6Bf\x89v\f&\xacYG\x7fӇ")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line D0l3mt.go:1*/ string(data)
	}(), "root", xrxPR_JAf5Q.root, func() /*line OFxH_UU6acaw.go:1*/ string {
		key := [] /*line OFxH_UU6acaw.go:1*/ byte("\xbf\xd6}\x8a\xd3\xcb\xe7\xd1")
		data := [] /*line OFxH_UU6acaw.go:1*/ byte("ܹ\x10\xfa\xbf\xae\x93\xb4")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line OFxH_UU6acaw.go:1*/ string(data)
	}(), v4faEBKJgI2.genMarker == nil)
	jToehJAGNky := &diskLayer{
		root:       xrxPR_JAf5Q.root,
		cache:      v4faEBKJgI2.cache,
		diskdb:     v4faEBKJgI2.diskdb,
		triedb:     v4faEBKJgI2.triedb,
		genMarker:  v4faEBKJgI2.genMarker,
		genPending: v4faEBKJgI2.genPending,
	}

	if v4faEBKJgI2.genMarker != nil && v4faEBKJgI2.genAbort != nil {
		jToehJAGNky.genMarker = v4faEBKJgI2.genMarker
		jToehJAGNky.genAbort = /*line Mx7sRca.go:1*/ make(chan chan *generatorStats)
		go /*line ql8nokhUf6.go:1*/ jToehJAGNky.tliqpCic(drjkE0)
	}
	return jToehJAGNky
}

func (rNemY7HB *CE2m1DUB6wW) Release() {
	if hbsnnIi := /*line IKsQQxo.go:1*/ rNemY7HB.fPkcGryZI(); hbsnnIi != nil {
		/*line AU5cE1MSj7.go:1*/ hbsnnIi.Release()
	}
}

func (rNemY7HB *CE2m1DUB6wW) Journal(z1BBTN2UX common.Hash) (common.Hash, error) {

	pc2Fq1jYVnOO := /*line pBNBTuLoMuyG.go:1*/ rNemY7HB.Snapshot(z1BBTN2UX)
	if pc2Fq1jYVnOO == nil {
		return common.Hash{} /*line DT4Yfc46h8.go:1*/, fmt.Errorf(func() /*line LbQsOYahqE.go:1*/ string {
			key := [] /*line LbQsOYahqE.go:1*/ byte("\xd52\xcbVI\xf6y\xba\x0e\xf4r\t\xdf|\xac\xf2\xe0\x95\xf2\xa7\xd7\x06")
			data := [] /*line LbQsOYahqE.go:1*/ byte("\x9e<\x96\x1a*r\xf6\xba\x12g\xb3\x1a\x99\xe1t{\x89ށ\u0097a")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line LbQsOYahqE.go:1*/ string(data)
		}(), z1BBTN2UX)
	}

	/*line Idh2TBoCA4IN.go:1*/
	rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line fugerWKKTD.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	g0iutUWAEY := /*line VnEaxXG00jT.go:1*/ new(bytes.Buffer)
	if chyZ8Q := /*line LnGuZ3.go:1*/ rlp.Encode(g0iutUWAEY, journalVersion); chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}
	w05cVcQ := /*line GpjX5q.go:1*/ rNemY7HB.dHqydahw0x9t()
	if w05cVcQ == (common.Hash{}) {
		return common.Hash{} /*line In2f1RW.go:1*/, errors.New(func() /*line PuyjqTWe6De.go:1*/ string {
			fullData := [] /*line PuyjqTWe6De.go:1*/ byte("\xa9ۃ\xbc\xf5\x97\x13>`K?7MRϠI Ŀ\x89\"\xce\xc1\xa8\xaf6\xcd\xc5a\xaa\xd5>\xa2")
			idxKey := [] /*line PuyjqTWe6De.go:1*/ byte("\x1f\x19D\x9a\x8a\xf1죩")
			data := /*line PuyjqTWe6De.go:1*/ make([]byte, 0, 18)
			data = /*line PuyjqTWe6De.go:1*/ append(data, fullData[155^ /*line PuyjqTWe6De.go:1*/ int(idxKey[4])]^fullData[154^ /*line PuyjqTWe6De.go:1*/ int(idxKey[4])], fullData[22^ /*line PuyjqTWe6De.go:1*/ int(idxKey[1])]+fullData[15^ /*line PuyjqTWe6De.go:1*/ int(idxKey[1])], fullData[79^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])]-fullData[83^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])], fullData[243^ /*line PuyjqTWe6De.go:1*/ int(idxKey[5])]-fullData[228^ /*line PuyjqTWe6De.go:1*/ int(idxKey[5])], fullData[63^ /*line PuyjqTWe6De.go:1*/ int(idxKey[0])]^fullData[18^ /*line PuyjqTWe6De.go:1*/ int(idxKey[0])], fullData[144^ /*line PuyjqTWe6De.go:1*/ int(idxKey[4])]-fullData[145^ /*line PuyjqTWe6De.go:1*/ int(idxKey[4])], fullData[62^ /*line PuyjqTWe6De.go:1*/ int(idxKey[0])]-fullData[24^ /*line PuyjqTWe6De.go:1*/ int(idxKey[0])], fullData[64^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])]-fullData[91^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])], fullData[150^ /*line PuyjqTWe6De.go:1*/ int(idxKey[4])]-fullData[151^ /*line PuyjqTWe6De.go:1*/ int(idxKey[4])], fullData[137^ /*line PuyjqTWe6De.go:1*/ int(idxKey[3])]+fullData[132^ /*line PuyjqTWe6De.go:1*/ int(idxKey[3])], fullData[136^ /*line PuyjqTWe6De.go:1*/ int(idxKey[3])]+fullData[131^ /*line PuyjqTWe6De.go:1*/ int(idxKey[3])], fullData[66^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])]-fullData[92^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])], fullData[68^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])]-fullData[80^ /*line PuyjqTWe6De.go:1*/ int(idxKey[2])], fullData[30^ /*line PuyjqTWe6De.go:1*/ int(idxKey[0])]+fullData[26^ /*line PuyjqTWe6De.go:1*/ int(idxKey[0])], fullData[26^ /*line PuyjqTWe6De.go:1*/ int(idxKey[1])]-fullData[21^ /*line PuyjqTWe6De.go:1*/ int(idxKey[1])], fullData[23^ /*line PuyjqTWe6De.go:1*/ int(idxKey[1])]-fullData[17^ /*line PuyjqTWe6De.go:1*/ int(idxKey[1])], fullData[248^ /*line PuyjqTWe6De.go:1*/ int(idxKey[5])]^fullData[251^ /*line PuyjqTWe6De.go:1*/ int(idxKey[5])])
			return /*line PuyjqTWe6De.go:1*/ string(data)
		}())
	}

	if chyZ8Q := /*line JWP9kafdrHf.go:1*/ rlp.Encode(g0iutUWAEY, w05cVcQ); chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}

	v4faEBKJgI2, chyZ8Q := /*line irjK83WZha.go:1*/ pc2Fq1jYVnOO.(snapshot).Journal(g0iutUWAEY)
	if chyZ8Q != nil {
		return common.Hash{}, chyZ8Q
	}

	/*line Mw7XhJA1mWg.go:1*/
	rawdb.WriteSnapshotJournal(rNemY7HB.cmZIXcy /*line zOd1M83YQyw.go:1*/, g0iutUWAEY.Bytes())
	return v4faEBKJgI2, nil
}

func (rNemY7HB *CE2m1DUB6wW) Rebuild(z1BBTN2UX common.Hash) {
	/*line Ci0gBmsm6r.go:1*/ rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line A7fMUhd9.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	/*line xTU4guT.go:1*/
	rawdb.DeleteSnapshotRecoveryNumber(rNemY7HB.cmZIXcy)
	/*line oCPzkI.go:1*/ rawdb.DeleteSnapshotDisabled(rNemY7HB.cmZIXcy)

	for _, _sAQiThp := range rNemY7HB.cMczP6tquWhX {
		switch _sAQiThp := _sAQiThp.(type) {
		case *diskLayer:

			if _sAQiThp.genAbort != nil {
				stpiKC6okfS8 := /*line rEfyPec0xz.go:1*/ make(chan *generatorStats)
				_sAQiThp.genAbort <- stpiKC6okfS8
				<-stpiKC6okfS8
			}

			/*line Y0Oc6H_Qk.go:1*/
			_sAQiThp.lock.Lock()
			_sAQiThp.stale = true
			/*line R8m4ZhxmC55C.go:1*/ _sAQiThp.lock.Unlock()

		case *diffLayer:

			/*line DkAf1SE.go:1*/
			_sAQiThp.lock.Lock()
			/*line GfJoZmS.go:1*/ _sAQiThp.stale.Store(true)
			/*line R8L_GTNc.go:1*/ _sAQiThp.lock.Unlock()

		default:
			/*line yk7nOSu.go:1*/ panic( /*line pPBcs8fcUE.go:1*/ fmt.Sprintf(func() /*line ydDwiZ9K.go:1*/ string {
				seed := /*line ydDwiZ9K.go:1*/ byte(130)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line ydDwiZ9K.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line ydDwiZ9K.go:1*/ fnc(243)(249)(253)(3)(1)(8)(247)(178)(76)(245)(24)(236)(13)(174)(84)(5)(247)(245)(213)(230)(5)(47)
				return /*line ydDwiZ9K.go:1*/ string(data)
			}(), _sAQiThp))
		}
	}

	/*line EiB3qCw2.go:1*/
	log.Info(func() /*line m5al8fP.go:1*/ string {
		data := /*line m5al8fP.go:1*/ make([]byte, 0, 26)
		i := 7
		decryptKey := 192
		for counter := 0; i != 13; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				i = 13
				for y := range data {
					data[y] = data[y] - /*line m5al8fP.go:1*/ byte(decryptKey^y)
				}
			case 3:
				data = /*line m5al8fP.go:1*/ append(data, "S\xa5\xa9"...,
				)
				i = 9
			case 6:
				data = /*line m5al8fP.go:1*/ append(data, 153)
				i = 11
			case 9:
				data = /*line m5al8fP.go:1*/ append(data, 149)
				i = 10
			case 1:
				i = 5
				data = /*line m5al8fP.go:1*/ append(data, 159)
			case 0:
				data = /*line m5al8fP.go:1*/ append(data, "\xaf\xa6\xa8"...,
				)
				i = 8
			case 12:
				data = /*line m5al8fP.go:1*/ append(data, 149)
				i = 4
			case 10:
				i = 6
				data = /*line m5al8fP.go:1*/ append(data, "\xab\x9bI\x9b"...,
				)
			case 8:
				data = /*line m5al8fP.go:1*/ append(data, "\xa3\xa7\x9f\x97"...,
				)
				i = 3
			case 2:
				i = 1
				data = /*line m5al8fP.go:1*/ append(data, 157)
			case 7:
				i = 0
				data = /*line m5al8fP.go:1*/ append(data, "\x8b\x9d\x9d"...,
				)
			case 5:
				i = 12
				data = /*line m5al8fP.go:1*/ append(data, "\x97\x9d"...,
				)
			case 11:
				data = /*line m5al8fP.go:1*/ append(data, 139)
				i = 2
			}
		}
		return /*line m5al8fP.go:1*/ string(data)
	}())
	rNemY7HB.cMczP6tquWhX = map[common.Hash]snapshot{
		z1BBTN2UX: /*line xzHpSNP5YUm.go:1*/ gazN04(rNemY7HB.cmZIXcy, rNemY7HB.sKxDZofwkrd, rNemY7HB.nDn2Q30J0Dk.AalXQA16, z1BBTN2UX),
	}
}

func (rNemY7HB *CE2m1DUB6wW) AccountIterator(z1BBTN2UX common.Hash, gnGBaeoLX common.Hash) (Nq4YsH_, error) {
	jdkNTRyBJ, chyZ8Q := /*line aik8NzfAcl.go:1*/ rNemY7HB.rZ7f5grs()
	if chyZ8Q != nil {
		return nil, chyZ8Q
	}
	if jdkNTRyBJ {
		return nil, M9ezyyMAXN
	}
	return /*line fLMkXjaLx.go:1*/ qEhLP2hRX(rNemY7HB, z1BBTN2UX, gnGBaeoLX)
}

func (rNemY7HB *CE2m1DUB6wW) StorageIterator(z1BBTN2UX common.Hash, evzmhL1 common.Hash, gnGBaeoLX common.Hash) (PE_7UyJghxy, error) {
	jdkNTRyBJ, chyZ8Q := /*line ZY3iOQFzo.go:1*/ rNemY7HB.rZ7f5grs()
	if chyZ8Q != nil {
		return nil, chyZ8Q
	}
	if jdkNTRyBJ {
		return nil, M9ezyyMAXN
	}
	return /*line u_DLk84iXLF.go:1*/ h4wrLvWBbqCm(rNemY7HB, z1BBTN2UX, evzmhL1, gnGBaeoLX)
}

func (rNemY7HB *CE2m1DUB6wW) Verify(z1BBTN2UX common.Hash) error {
	avYdDZ2tZiG, chyZ8Q := /*line DVEAoEjg.go:1*/ rNemY7HB.AccountIterator(z1BBTN2UX, common.Hash{})
	if chyZ8Q != nil {
		return chyZ8Q
	}
	defer /*line V91SX6.go:1*/ avYdDZ2tZiG.Release()

	pfeznfxya, chyZ8Q := /*line nSGdshTaO.go:1*/ qQF5aSpK2E9(nil, "", avYdDZ2tZiG, common.Hash{}, d8k2UL, func(lDBEbd ethdb.KeyValueWriter, dcJdHV, j8UX1s common.Hash, lXazFVVqJfD_ *bS1vBqz) (common.Hash, error) {
		sK5vOn6BL2OZ, chyZ8Q := /*line ySXjL0.go:1*/ rNemY7HB.StorageIterator(z1BBTN2UX, dcJdHV, common.Hash{})
		if chyZ8Q != nil {
			return common.Hash{}, chyZ8Q
		}
		defer /*line HQ0RD01ZO.go:1*/ sK5vOn6BL2OZ.Release()

		xhOzkRrKDZ, chyZ8Q := /*line P1AxHNpNcYSB.go:1*/ qQF5aSpK2E9(nil, "", sK5vOn6BL2OZ, dcJdHV, d8k2UL, nil, lXazFVVqJfD_, false)
		if chyZ8Q != nil {
			return common.Hash{}, chyZ8Q
		}
		return xhOzkRrKDZ, nil
	}, /*line fqPPvBDWbWHP.go:1*/ kT7u5MT_Lv(), true)

	if chyZ8Q != nil {
		return chyZ8Q
	}
	if pfeznfxya != z1BBTN2UX {
		return /*line JyIOEaa.go:1*/ fmt.Errorf(func() /*line GWR8R1L.go:1*/ string {
			seed := /*line GWR8R1L.go:1*/ byte(122)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line GWR8R1L.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line GWR8R1L.go:1*/ fnc(249)(1)(237)(19)(241)(187)(82)(253)(0)(5)(172)(72)(249)(18)(245)(184)(77)(252)(10)(250)(244)(19)(239)(5)(210)(230)(71)(8)(5)(172)(5)(83)(180)(244)(87)(234)(13)(6)(172)(5)(83)
			return /*line GWR8R1L.go:1*/ string(data)
		}(), pfeznfxya, z1BBTN2UX)
	}
	return nil
}

func (rNemY7HB *CE2m1DUB6wW) fPkcGryZI() *diskLayer {
	var pc2Fq1jYVnOO snapshot
	for _, ui48PUZCY := range rNemY7HB.cMczP6tquWhX {
		pc2Fq1jYVnOO = ui48PUZCY
		break
	}
	if pc2Fq1jYVnOO == nil {
		return nil
	}
	switch _sAQiThp := pc2Fq1jYVnOO.(type) {
	case *diskLayer:
		return _sAQiThp
	case *diffLayer:
		return _sAQiThp.origin
	default:
		/*line fNXEuCp.go:1*/ panic( /*line Kr_htWTYlH.go:1*/ fmt.Sprintf(func() /*line S9LZRz1.go:1*/ string {
			fullData := [] /*line S9LZRz1.go:1*/ byte("~°\x89h.\x06L\t\xbc\xab-~\xa0\x96\x97\xa2$\x93\xbfQ\xe6\x18\x0fi\xb2\xab)\xc5\xcc\xd1\x1bj\xca\xff\x91\x8e\xd7")
			idxKey := [] /*line S9LZRz1.go:1*/ byte("\xb8\r\xcf\xf6\xf3vM")
			data := /*line S9LZRz1.go:1*/ make([]byte, 0, 20)
			data = /*line S9LZRz1.go:1*/ append(data, fullData[177^ /*line S9LZRz1.go:1*/ int(idxKey[0])]-fullData[183^ /*line S9LZRz1.go:1*/ int(idxKey[0])], fullData[84^ /*line S9LZRz1.go:1*/ int(idxKey[6])]^fullData[88^ /*line S9LZRz1.go:1*/ int(idxKey[6])], fullData[110^ /*line S9LZRz1.go:1*/ int(idxKey[6])]^fullData[71^ /*line S9LZRz1.go:1*/ int(idxKey[6])], fullData[117^ /*line S9LZRz1.go:1*/ int(idxKey[5])]-fullData[110^ /*line S9LZRz1.go:1*/ int(idxKey[5])], fullData[103^ /*line S9LZRz1.go:1*/ int(idxKey[5])]+fullData[98^ /*line S9LZRz1.go:1*/ int(idxKey[5])], fullData[17^ /*line S9LZRz1.go:1*/ int(idxKey[1])]^fullData[23^ /*line S9LZRz1.go:1*/ int(idxKey[1])], fullData[100^ /*line S9LZRz1.go:1*/ int(idxKey[5])]+fullData[104^ /*line S9LZRz1.go:1*/ int(idxKey[5])], fullData[214^ /*line S9LZRz1.go:1*/ int(idxKey[3])]^fullData[225^ /*line S9LZRz1.go:1*/ int(idxKey[3])], fullData[65^ /*line S9LZRz1.go:1*/ int(idxKey[6])]^fullData[91^ /*line S9LZRz1.go:1*/ int(idxKey[6])], fullData[212^ /*line S9LZRz1.go:1*/ int(idxKey[3])]^fullData[248^ /*line S9LZRz1.go:1*/ int(idxKey[3])], fullData[242^ /*line S9LZRz1.go:1*/ int(idxKey[3])]^fullData[240^ /*line S9LZRz1.go:1*/ int(idxKey[3])], fullData[105^ /*line S9LZRz1.go:1*/ int(idxKey[5])]^fullData[118^ /*line S9LZRz1.go:1*/ int(idxKey[5])], fullData[8^ /*line S9LZRz1.go:1*/ int(idxKey[1])]-fullData[44^ /*line S9LZRz1.go:1*/ int(idxKey[1])], fullData[86^ /*line S9LZRz1.go:1*/ int(idxKey[6])]^fullData[69^ /*line S9LZRz1.go:1*/ int(idxKey[6])], fullData[235^ /*line S9LZRz1.go:1*/ int(idxKey[3])]^fullData[251^ /*line S9LZRz1.go:1*/ int(idxKey[3])], fullData[30^ /*line S9LZRz1.go:1*/ int(idxKey[1])]+fullData[29^ /*line S9LZRz1.go:1*/ int(idxKey[1])], fullData[253^ /*line S9LZRz1.go:1*/ int(idxKey[3])]+fullData[241^ /*line S9LZRz1.go:1*/ int(idxKey[3])], fullData[157^ /*line S9LZRz1.go:1*/ int(idxKey[0])]+fullData[156^ /*line S9LZRz1.go:1*/ int(idxKey[0])], fullData[206^ /*line S9LZRz1.go:1*/ int(idxKey[2])]+fullData[205^ /*line S9LZRz1.go:1*/ int(idxKey[2])])
			return /*line S9LZRz1.go:1*/ string(data)
		}(), pc2Fq1jYVnOO))
	}
}

func (rNemY7HB *CE2m1DUB6wW) dHqydahw0x9t() common.Hash {
	fPkcGryZI := /*line Z9Lxk3fMf.go:1*/ rNemY7HB.fPkcGryZI()
	if fPkcGryZI == nil {
		return common.Hash{}
	}
	return /*line eD16CS9.go:1*/ fPkcGryZI.Root()
}

func (rNemY7HB *CE2m1DUB6wW) rZ7f5grs() (bool, error) {
	/*line BnVhhDNwj.go:1*/ rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line DUsC4zzvs.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	_sAQiThp := /*line w_EuBJA.go:1*/ rNemY7HB.fPkcGryZI()
	if _sAQiThp == nil {
		return false /*line KlVI6ra.go:1*/, errors.New(func() /*line MJNskADE3ayW.go:1*/ string {
			seed := /*line MJNskADE3ayW.go:1*/ byte(140)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line MJNskADE3ayW.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line MJNskADE3ayW.go:1*/ fnc(240)(229)(212)(160)(245)(54)(97)(218)(160)(77)(72)(217)(188)(37)(151)(42)(94)(188)(110)(225)(187)
			return /*line MJNskADE3ayW.go:1*/ string(data)
		}())
	}
	/*line aJJQFnCa.go:1*/ _sAQiThp.lock.RLock()
	defer /*line zk1PMX5uLf_.go:1*/ _sAQiThp.lock.RUnlock()
	return _sAQiThp.genMarker != nil, nil
}

func (rNemY7HB *CE2m1DUB6wW) DiskRoot() common.Hash {
	/*line HKCNuTR8.go:1*/ rNemY7HB.kkq_aKqj4Q.Lock()
	defer /*line SikCV7U07kLK.go:1*/ rNemY7HB.kkq_aKqj4Q.Unlock()

	return /*line XhIPhMybQI.go:1*/ rNemY7HB.dHqydahw0x9t()
}

func (rNemY7HB *CE2m1DUB6wW) Size() (fvKNpIjqNl common.StorageSize, a1mADLUs common.StorageSize) {
	/*line Vofhq_l8.go:1*/ rNemY7HB.kkq_aKqj4Q.RLock()
	defer /*line wiV0Ka.go:1*/ rNemY7HB.kkq_aKqj4Q.RUnlock()

	var f4_oKO4jp common.StorageSize
	for _, _sAQiThp := range rNemY7HB.cMczP6tquWhX {
		if _sAQiThp, jdkNTRyBJ := _sAQiThp.(*diffLayer); jdkNTRyBJ {
			f4_oKO4jp += /*line boOhC9tLn5.go:1*/ common.StorageSize(_sAQiThp.memory)
		}
	}
	return f4_oKO4jp, 0
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append
var _ sync.Cond
var _ types.AccessList
var _ = common.AbsolutePath
var _ = rawdb.BestUpdateKey
var _ ethdb.AncientReader
var _ = log.Crit
var _ = metrics.AccountingRegistry
var _ = rlp.AppendUint64
var _ triedb.Config
