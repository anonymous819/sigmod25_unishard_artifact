//line :1
package state

import "github.com/ethereum/go-ethereum/metrics"

var (
	t14aLF4Y = /*line HtWmdMaT.go:1*/ metrics.NewRegisteredMeter(func() /*line GcbgCsB.go:1*/ string {
		seed := /*line GcbgCsB.go:1*/ byte(240)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line GcbgCsB.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line GcbgCsB.go:1*/ fnc(99)(199)(123)(9)(3)(208)(230)(199)(130)(1)(21)(27)(0)(50)(102)(204)(164)(78)(149)(48)
		return /*line GcbgCsB.go:1*/ string(data)
	}(), nil)
	jWXmaMS3d = /*line AUBJiJakAD.go:1*/ metrics.NewRegisteredMeter(func() /*line TVsr1Npx.go:1*/ string {
		data := /*line TVsr1Npx.go:1*/ make([]byte, 0, 21)
		i := 6
		decryptKey := 10
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 0
				for y := range data {
					data[y] = data[y] ^ /*line TVsr1Npx.go:1*/ byte(decryptKey^y)
				}
			case 5:
				i = 7
				data = /*line TVsr1Npx.go:1*/ append(data, "cx|"...,
				)
			case 8:
				data = /*line TVsr1Npx.go:1*/ append(data, "ab"...,
				)
				i = 2
			case 4:
				data = /*line TVsr1Npx.go:1*/ append(data, 106)
				i = 9
			case 7:
				i = 4
				data = /*line TVsr1Npx.go:1*/ append(data, "jz7"...,
				)
			case 3:
				i = 5
				data = /*line TVsr1Npx.go:1*/ append(data, ">g"...,
				)
			case 1:
				data = /*line TVsr1Npx.go:1*/ append(data, "cu"...,
				)
				i = 3
			case 6:
				data = /*line TVsr1Npx.go:1*/ append(data, "gaw"...,
				)
				i = 1
			case 9:
				data = /*line TVsr1Npx.go:1*/ append(data, "ntvd"...,
				)
				i = 8
			}
		}
		return /*line TVsr1Npx.go:1*/ string(data)
	}(), nil)
	qqNAVrhZ = /*line ljXytHwg6w2A.go:1*/ metrics.NewRegisteredMeter(func() /*line Tf7O4kwEEamx.go:1*/ string {
		key := [] /*line Tf7O4kwEEamx.go:1*/ byte("\x05o\xf15\x1e:\xa5\xd2\xeft,!\x85\x7fB\x9e\x99c=\xbe")
		data := [] /*line Tf7O4kwEEamx.go:1*/ byte("v\x1b\x90A{\x15\xc1\xb7\x83\x11XD\xaa\x1e!\xfd\xf6\x16S\xca")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line Tf7O4kwEEamx.go:1*/ string(data)
	}(), nil)
	ewSUzVXKn = /*line NAvHyaYe.go:1*/ metrics.NewRegisteredMeter(func() /*line WSmITcnaaIJ.go:1*/ string {
		fullData := [] /*line WSmITcnaaIJ.go:1*/ byte("\xcb5\\S\x10\x9cX\xfb\x19p\x1c\xfc\xfd\xf0j\x8fEଜ\x17!\xeb\xbdr[\xb5\xd0\xddPJX?\x0e\x1buH\xe3y\r")
		idxKey := [] /*line WSmITcnaaIJ.go:1*/ byte("\xdcj\\t,\uebcd\\\xa3Rj\x8e\xfa\xf7\x84")
		data := /*line WSmITcnaaIJ.go:1*/ make([]byte, 0, 21)
		data = /*line WSmITcnaaIJ.go:1*/ append(data, fullData[47^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[4])]-fullData[61^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[4])], fullData[180^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[6])]-fullData[173^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[6])], fullData[169^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[9])]+fullData[179^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[9])], fullData[99^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[11])]-fullData[97^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[11])], fullData[75^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])]-fullData[90^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])], fullData[75^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[1])]+fullData[127^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[1])], fullData[152^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[12])]+fullData[168^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[12])], fullData[123^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])]+fullData[67^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])], fullData[255^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[0])]^fullData[212^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[0])], fullData[141^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[6])]+fullData[177^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[6])], fullData[88^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])]-fullData[89^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])], fullData[202^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[5])]-fullData[203^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[5])], fullData[81^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])]+fullData[124^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[8])], fullData[230^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[13])]-fullData[244^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[13])], fullData[248^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[14])]^fullData[240^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[14])], fullData[149^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[7])]+fullData[129^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[7])], fullData[227^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[14])]+fullData[238^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[14])], fullData[237^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[14])]+fullData[229^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[14])], fullData[253^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[5])]+fullData[238^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[5])], fullData[119^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[11])]^fullData[107^ /*line WSmITcnaaIJ.go:1*/ int(idxKey[11])])
		return /*line WSmITcnaaIJ.go:1*/ string(data)
	}(), nil)
	xNTiw5VKek = /*line aQtjNL.go:1*/ metrics.NewRegisteredMeter(func() /*line GI9Pgt9sxF1.go:1*/ string {
		fullData := [] /*line GI9Pgt9sxF1.go:1*/ byte("\x90?\xbfn\xaf\xab\xaa\xf4{@H\x80\x1d!G\xcb\xea!\x83\x9b\xa4_\xb8\x03_<\x0f \x9dFkc[\x96\x16N=\t,\xaf\xedԅ\xebHo\xfbm\xb2E")
		idxKey := [] /*line GI9Pgt9sxF1.go:1*/ byte("\x117\xd9z\x00\xf5\x1b\x190")
		data := /*line GI9Pgt9sxF1.go:1*/ make([]byte, 0, 26)
		data = /*line GI9Pgt9sxF1.go:1*/ append(data, fullData[215^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[2])]-fullData[240^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[2])], fullData[7^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]+fullData[60^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[48^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])]+fullData[4^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])], fullData[1^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]-fullData[50^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[40^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]+fullData[2^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[43^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[1])]-fullData[52^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[1])], fullData[61^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]^fullData[53^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[108^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[3])]-fullData[112^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[3])], fullData[24^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]^fullData[57^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[9^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])]+fullData[13^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])], fullData[26^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[1])]-fullData[25^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[1])], fullData[240^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[5])]-fullData[232^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[5])], fullData[11^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])]^fullData[39^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])], fullData[33^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])]+fullData[15^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])], fullData[120^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[3])]+fullData[110^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[3])], fullData[41^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[8])]^fullData[37^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[8])], fullData[58^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]^fullData[8^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[14^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[0])]^fullData[51^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[0])], fullData[31^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[8])]^fullData[39^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[8])], fullData[8^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])]^fullData[26^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[4])], fullData[49^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])]^fullData[11^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[7])], fullData[230^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[5])]-fullData[211^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[5])], fullData[28^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[6])]^fullData[27^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[6])], fullData[32^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[8])]-fullData[26^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[8])], fullData[249^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[5])]-fullData[243^ /*line GI9Pgt9sxF1.go:1*/ int(idxKey[5])])
		return /*line GI9Pgt9sxF1.go:1*/ string(data)
	}(), nil)
	tRW4TBhii = /*line xIdSCizun.go:1*/ metrics.NewRegisteredMeter(func() /*line RT__OOga.go:1*/ string {
		key := [] /*line RT__OOga.go:1*/ byte("(#\x94\x80\xc2\x18ot\xedS\xa3?\xf2X\xa5*\xf5AN\x8e\xda\x04k\n_")
		data := [] /*line RT__OOga.go:1*/ byte("[W\xf5\xf4\xa77\x1a\x04\x892\xd7Z\xdd+\xd1E\x87 )\xeb\xb4k\x0fo,")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line RT__OOga.go:1*/ string(data)
	}(), nil)
	db1BShRmyWbt = /*line QkZCIzaieq.go:1*/ metrics.NewRegisteredMeter(func() /*line sZvkilj.go:1*/ string {
		data := /*line sZvkilj.go:1*/ make([]byte, 0, 26)
		i := 13
		decryptKey := 234
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 0
				data = /*line sZvkilj.go:1*/ append(data, "\x83\x82\x83\x90"...,
				)
			case 4:
				data = /*line sZvkilj.go:1*/ append(data, "\x9d\x84"...,
				)
				i = 7
			case 11:
				i = 8
				data = /*line sZvkilj.go:1*/ append(data, "\x8d\x8d"...,
				)
			case 12:
				i = 1
				data = /*line sZvkilj.go:1*/ append(data, "\x9a\x8c"...,
				)
			case 1:
				i = 9
				data = /*line sZvkilj.go:1*/ append(data, 152)
			case 9:
				i = 11
				data = /*line sZvkilj.go:1*/ append(data, "\x8e\xc5"...,
				)
			case 10:
				data = /*line sZvkilj.go:1*/ append(data, "\x81\xcc"...,
				)
				i = 2
			case 6:
				i = 10
				data = /*line sZvkilj.go:1*/ append(data, "\x83\x91"...,
				)
			case 13:
				i = 12
				data = /*line sZvkilj.go:1*/ append(data, 156)
			case 3:
				data = /*line sZvkilj.go:1*/ append(data, "\x95\x95\x9d"...,
				)
				i = 4
			case 7:
				for y := range data {
					data[y] = data[y] ^ /*line sZvkilj.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 0:
				i = 3
				data = /*line sZvkilj.go:1*/ append(data, "\x8b\x93\x88"...,
				)
			case 8:
				i = 6
				data = /*line sZvkilj.go:1*/ append(data, 139)
			}
		}
		return /*line sZvkilj.go:1*/ string(data)
	}(), nil)
	qiwd7Qf6LoOn = /*line fzvYD7m1RKDz.go:1*/ metrics.NewRegisteredMeter(func() /*line KkUlFbLNa.go:1*/ string {
		fullData := [] /*line KkUlFbLNa.go:1*/ byte("Y\xcf7خ\xf7\x1b\x7f\x1f\x1e\xa5\xaa\xbc\x13E˺\xaa\xf6`;\x93a/-\xc4\x0f\x05d\x96\xd2\xde\xfdЛk(y`\xc1`\x17\xf5\x9cr͜V\xd2:")
		idxKey := [] /*line KkUlFbLNa.go:1*/ byte("m\xb3\x97\x9f\xcb\x0e\x1a\xc5")
		data := /*line KkUlFbLNa.go:1*/ make([]byte, 0, 26)
		data = /*line KkUlFbLNa.go:1*/ append(data, fullData[200^ /*line KkUlFbLNa.go:1*/ int(idxKey[7])]+fullData[214^ /*line KkUlFbLNa.go:1*/ int(idxKey[7])], fullData[45^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])]^fullData[6^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])], fullData[31^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])]-fullData[7^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])], fullData[207^ /*line KkUlFbLNa.go:1*/ int(idxKey[7])]+fullData[196^ /*line KkUlFbLNa.go:1*/ int(idxKey[7])], fullData[132^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])]^fullData[185^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])], fullData[63^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])]+fullData[36^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])], fullData[12^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])]-fullData[58^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])], fullData[39^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])]^fullData[34^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])], fullData[146^ /*line KkUlFbLNa.go:1*/ int(idxKey[1])]^fullData[191^ /*line KkUlFbLNa.go:1*/ int(idxKey[1])], fullData[141^ /*line KkUlFbLNa.go:1*/ int(idxKey[2])]+fullData[184^ /*line KkUlFbLNa.go:1*/ int(idxKey[2])], fullData[153^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])]+fullData[159^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])], fullData[222^ /*line KkUlFbLNa.go:1*/ int(idxKey[4])]^fullData[217^ /*line KkUlFbLNa.go:1*/ int(idxKey[4])], fullData[98^ /*line KkUlFbLNa.go:1*/ int(idxKey[0])]+fullData[113^ /*line KkUlFbLNa.go:1*/ int(idxKey[0])], fullData[2^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])]-fullData[10^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])], fullData[189^ /*line KkUlFbLNa.go:1*/ int(idxKey[1])]+fullData[164^ /*line KkUlFbLNa.go:1*/ int(idxKey[1])], fullData[30^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])]^fullData[61^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])], fullData[16^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])]-fullData[38^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])], fullData[9^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])]-fullData[7^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])], fullData[5^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])]^fullData[35^ /*line KkUlFbLNa.go:1*/ int(idxKey[5])], fullData[180^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])]-fullData[157^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])], fullData[52^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])]+fullData[42^ /*line KkUlFbLNa.go:1*/ int(idxKey[6])], fullData[124^ /*line KkUlFbLNa.go:1*/ int(idxKey[0])]-fullData[121^ /*line KkUlFbLNa.go:1*/ int(idxKey[0])], fullData[179^ /*line KkUlFbLNa.go:1*/ int(idxKey[2])]-fullData[142^ /*line KkUlFbLNa.go:1*/ int(idxKey[2])], fullData[128^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])]-fullData[186^ /*line KkUlFbLNa.go:1*/ int(idxKey[3])], fullData[233^ /*line KkUlFbLNa.go:1*/ int(idxKey[4])]+fullData[200^ /*line KkUlFbLNa.go:1*/ int(idxKey[4])])
		return /*line KkUlFbLNa.go:1*/ string(data)
	}(), nil)

	bKowp_ = /*line Lc684dbYuP.go:1*/ metrics.NewRegisteredGauge(func() /*line FUBEZqq.go:1*/ string {
		seed := /*line FUBEZqq.go:1*/ byte(154)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line FUBEZqq.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line FUBEZqq.go:1*/ fnc(217)(1)(237)(19)(241)(202)(53)(1)(7)(249)(15)(241)(202)(68)(1)(251)(3)(239)(6)(254)(202)(62)(244)(23)(183)(68)(249)(3)(5)
		return /*line FUBEZqq.go:1*/ string(data)
	}(), nil)
	qOuP3t = /*line Tua3ThHZE.go:1*/ metrics.NewRegisteredGauge(func() /*line mOUdp7.go:1*/ string {
		seed := /*line mOUdp7.go:1*/ byte(146)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line mOUdp7.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line mOUdp7.go:1*/ fnc(225)(1)(237)(19)(241)(202)(53)(1)(7)(249)(15)(241)(202)(68)(1)(251)(3)(239)(6)(254)(202)(62)(244)(23)(183)(68)(246)(17)(235)
		return /*line mOUdp7.go:1*/ string(data)
	}(), nil)
	fxF6osapsiIK = /*line dW8uL7s1v1lr.go:1*/ metrics.NewRegisteredResettingTimer(func() /*line wAHjlijGK.go:1*/ string {
		key := [] /*line wAHjlijGK.go:1*/ byte("X\xa7-\xf7\xc8dY\x1a2\x8a9\xb8\x81b\x89\xb7\xb7\xfb\bv_\x9c\xc53\xd2\xfd")
		data := [] /*line wAHjlijGK.go:1*/ byte("\xcb\x1b\x8ek-\x93\xbd\x7f\x9e\xef\xad\x1d\xb0\xd5\xfd&)\\oێ\x10.\xa07o")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line wAHjlijGK.go:1*/ string(data)
	}(), nil)
	jDTEIE_ = /*line S0EradaHyBa.go:1*/ metrics.NewRegisteredMeter(func() /*line omF6TOr.go:1*/ string {
		data := [] /*line omF6TOr.go:1*/ byte("Yt\xee\xee\xa6\xec\xc6el\xc3\xd8\xf9t\xef\x85A\xcag\xf6_l\x8a\xd4\xd2\xc7")
		positions := [...]byte{5, 4, 23, 0, 12, 21, 2, 3, 20, 14, 11, 15, 17, 18, 19, 19, 24, 3, 18, 23, 13, 13, 11, 14, 16, 22, 6, 21, 10, 19, 24, 9}
		for i := 0; i < 32; i += 2 {
			localKey := /*line omF6TOr.go:1*/ byte(i) + /*line omF6TOr.go:1*/ byte(positions[i]^positions[i+1]) + 136
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line omF6TOr.go:1*/ string(data)
	}(), nil)
	qt2Hcoh5RFFq = /*line Cv0OEC1uKM.go:1*/ metrics.NewRegisteredMeter(func() /*line FWI0DL.go:1*/ string {
		seed := /*line FWI0DL.go:1*/ byte(44)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line FWI0DL.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line FWI0DL.go:1*/ fnc(71)(1)(237)(19)(241)(202)(53)(1)(7)(249)(15)(241)(202)(68)(1)(251)(3)(239)(6)(254)(202)(68)(246)(17)(235)
		return /*line FWI0DL.go:1*/ string(data)
	}(), nil)
	mJkh5y = /*line Gt_KyPu1.go:1*/ metrics.NewRegisteredGauge(func() /*line earaFim.go:1*/ string {
		data := /*line earaFim.go:1*/ make([]byte, 0, 26)
		i := 8
		decryptKey := 220
		for counter := 0; i != 12; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				i = 3
				data = /*line earaFim.go:1*/ append(data, 74)
			case 14:
				data = /*line earaFim.go:1*/ append(data, 56)
				i = 9
			case 1:
				data = /*line earaFim.go:1*/ append(data, "5I"...,
				)
				i = 13
			case 6:
				data = /*line earaFim.go:1*/ append(data, 72)
				i = 14
			case 8:
				data = /*line earaFim.go:1*/ append(data, "IK"...,
				)
				i = 1
			case 10:
				data = /*line earaFim.go:1*/ append(data, "\x0246"...,
				)
				i = 5
			case 0:
				i = 12
				for y := range data {
					data[y] = data[y] - /*line earaFim.go:1*/ byte(decryptKey^y)
				}
			case 4:
				data = /*line earaFim.go:1*/ append(data, 62)
				i = 0
			case 11:
				data = /*line earaFim.go:1*/ append(data, "+*\xf1"...,
				)
				i = 7
			case 2:
				i = 6
				data = /*line earaFim.go:1*/ append(data, "\tNL"...,
				)
			case 9:
				data = /*line earaFim.go:1*/ append(data, 40)
				i = 11
			case 13:
				i = 10
				data = /*line earaFim.go:1*/ append(data, 55)
			case 3:
				data = /*line earaFim.go:1*/ append(data, "DPB"...,
				)
				i = 2
			case 7:
				i = 4
				data = /*line earaFim.go:1*/ append(data, "6+*"...,
				)
			}
		}
		return /*line earaFim.go:1*/ string(data)
	}(), nil)
)
var _ = metrics.AccountingRegistry
