//line :1
package node

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	log "unishard/log"
	message "unishard/message"
)

func (iai_HgOpk3 *coordblockbuilder) roWnZ9bUZa() {
	bzee5Hc06 := /*line lOaoeJS.go:1*/ http.NewServeMux()
	/*line vGljkQ6hVMjo.go:1*/ bzee5Hc06.HandleFunc("/", iai_HgOpk3.hZ4qb61uGeW)
	/*line vmhdL5.go:1*/ bzee5Hc06.HandleFunc("/query", iai_HgOpk3.dyeyfsXei)
	/*line Ix2c7d6v7w4Q.go:1*/ bzee5Hc06.HandleFunc(func() /*line H5r2YGRmomY.go:1*/ string {
		seed := /*line H5r2YGRmomY.go:1*/ byte(241)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line H5r2YGRmomY.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line H5r2YGRmomY.go:1*/ fnc(62)(67)(243)(12)(4)(240)(14)(1)(216)(25)(252)(3)(1)(13)
		return /*line H5r2YGRmomY.go:1*/ string(data)
	}(), iai_HgOpk3.sHpWltzZK)
	/*line KB8KOpD.go:1*/ bzee5Hc06.HandleFunc(func() /*line CaPteoWJQZ.go:1*/ string {
		seed := /*line CaPteoWJQZ.go:1*/ byte(167)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line CaPteoWJQZ.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line CaPteoWJQZ.go:1*/ fnc(214)(239)(209)(173)(89)(181)(108)(166)(131)(7)(245)(247)(244)(221)(191)(117)
		return /*line CaPteoWJQZ.go:1*/ string(data)
	}(), iai_HgOpk3.mODHJaCQQIF)
	/*line icKuI1hMKsDN.go:1*/ bzee5Hc06.HandleFunc(func() /*line TsoURSU.go:1*/ string {
		data := /*line TsoURSU.go:1*/ make([]byte, 0, 23)
		i := 7
		decryptKey := 126
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				i = 5
				data = /*line TsoURSU.go:1*/ append(data, 139)
			case 3:
				data = /*line TsoURSU.go:1*/ append(data, "þ\xcc"...,
				)
				i = 6
			case 10:
				i = 9
				data = /*line TsoURSU.go:1*/ append(data, "\xb6\xc1í"...,
				)
			case 4:
				i = 10
				data = /*line TsoURSU.go:1*/ append(data, "ä\xb8\xb3"...,
				)
			case 8:
				data = /*line TsoURSU.go:1*/ append(data, 187)
				i = 1
			case 9:
				data = /*line TsoURSU.go:1*/ append(data, 187)
				i = 0
			case 5:
				i = 3
				data = /*line TsoURSU.go:1*/ append(data, "\xb4\xcd\xd1"...,
				)
			case 1:
				i = 4
				data = /*line TsoURSU.go:1*/ append(data, 181)
			case 6:
				i = 8
				data = /*line TsoURSU.go:1*/ append(data, "\x9dɾ\xc2"...,
				)
			case 0:
				i = 2
				for y := range data {
					data[y] = data[y] - /*line TsoURSU.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return /*line TsoURSU.go:1*/ string(data)
	}(), iai_HgOpk3.qj9Jzl3)
	/*line Sa0lvH.go:1*/ bzee5Hc06.HandleFunc(func() /*line g84eLOsouE.go:1*/ string {
		data := [] /*line g84eLOsouE.go:1*/ byte("/W\x83rkerbu\x8dl\x86\x85{\x8eeM\x9d\x96eeUe\xb0w")
		positions := [...]byte{13, 7, 18, 14, 24, 20, 16, 24, 12, 2, 13, 15, 12, 17, 7, 24, 16, 11, 18, 20, 21, 18, 12, 21, 2, 15, 20, 9, 14, 23, 15, 13, 14, 20}
		for i := 0; i < 34; i += 2 {
			localKey := /*line g84eLOsouE.go:1*/ byte(i) + /*line g84eLOsouE.go:1*/ byte(positions[i]^positions[i+1]) + 232
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line g84eLOsouE.go:1*/ string(data)
	}(), iai_HgOpk3.aEknSj)

	lq54woHa3C, c2qNJpNX8d := /*line OHuRn9.go:1*/ url.Parse(func() /*line LZKogS.go:1*/ string {
		data := /*line LZKogS.go:1*/ make([]byte, 0, 22)
		i := 0
		decryptKey := 38
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 1
				data = /*line LZKogS.go:1*/ append(data, "\x04\b\x04\t"...,
				)
			case 8:
				data = /*line LZKogS.go:1*/ append(data, "\x01\x02"...,
				)
				i = 6
			case 0:
				data = /*line LZKogS.go:1*/ append(data, "FQT"...,
				)
				i = 3
			case 1:
				i = 4
				data = /*line LZKogS.go:1*/ append(data, "\b\x03"...,
				)
			case 7:
				data = /*line LZKogS.go:1*/ append(data, "\v\f\b\f"...,
				)
				i = 2
			case 4:
				for y := range data {
					data[y] = data[y] + /*line LZKogS.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 2:
				i = 9
				data = /*line LZKogS.go:1*/ append(data, "\x06\a"...,
				)
			case 3:
				data = /*line LZKogS.go:1*/ append(data, "O\x14\b"...,
				)
				i = 7
			case 9:
				i = 8
				data = /*line LZKogS.go:1*/ append(data, 0)
			}
		}
		return /*line LZKogS.go:1*/ string(data)
	}())
	if c2qNJpNX8d != nil {
		/*line IOMxNpw.go:1*/ log.FB5S6xdVix(func() /*line F5QDnJC.go:1*/ string {
			fullData := [] /*line F5QDnJC.go:1*/ byte("\xead\xef&JX\xb0\xa4\xf2\x8e\xbeg\xb3a\xd0\x0e\xf0G܀jkb\x9a(#\x9e\xb9f\x1f5+\xc3\xcf\xdc}\x0e\x18\xb9\n\xb0\x828\xaf")
			idxKey := [] /*line F5QDnJC.go:1*/ byte("\xee\xce\xc1<G.H~\x98\xf3\f<\xcb\xdb\xe0")
			data := /*line F5QDnJC.go:1*/ make([]byte, 0, 23)
			data = /*line F5QDnJC.go:1*/ append(data, fullData[194^ /*line F5QDnJC.go:1*/ int(idxKey[12])]-fullData[200^ /*line F5QDnJC.go:1*/ int(idxKey[12])], fullData[151^ /*line F5QDnJC.go:1*/ int(idxKey[8])]+fullData[132^ /*line F5QDnJC.go:1*/ int(idxKey[8])], fullData[60^ /*line F5QDnJC.go:1*/ int(idxKey[11])]^fullData[38^ /*line F5QDnJC.go:1*/ int(idxKey[11])], fullData[33^ /*line F5QDnJC.go:1*/ int(idxKey[3])]-fullData[23^ /*line F5QDnJC.go:1*/ int(idxKey[3])], fullData[193^ /*line F5QDnJC.go:1*/ int(idxKey[14])]^fullData[226^ /*line F5QDnJC.go:1*/ int(idxKey[14])], fullData[156^ /*line F5QDnJC.go:1*/ int(idxKey[8])]+fullData[135^ /*line F5QDnJC.go:1*/ int(idxKey[8])], fullData[228^ /*line F5QDnJC.go:1*/ int(idxKey[9])]-fullData[235^ /*line F5QDnJC.go:1*/ int(idxKey[9])], fullData[230^ /*line F5QDnJC.go:1*/ int(idxKey[1])]^fullData[236^ /*line F5QDnJC.go:1*/ int(idxKey[1])], fullData[203^ /*line F5QDnJC.go:1*/ int(idxKey[1])]-fullData[228^ /*line F5QDnJC.go:1*/ int(idxKey[1])], fullData[104^ /*line F5QDnJC.go:1*/ int(idxKey[6])]^fullData[68^ /*line F5QDnJC.go:1*/ int(idxKey[6])], fullData[25^ /*line F5QDnJC.go:1*/ int(idxKey[10])]^fullData[43^ /*line F5QDnJC.go:1*/ int(idxKey[10])], fullData[90^ /*line F5QDnJC.go:1*/ int(idxKey[7])]+fullData[127^ /*line F5QDnJC.go:1*/ int(idxKey[7])], fullData[55^ /*line F5QDnJC.go:1*/ int(idxKey[5])]-fullData[40^ /*line F5QDnJC.go:1*/ int(idxKey[5])], fullData[226^ /*line F5QDnJC.go:1*/ int(idxKey[2])]^fullData[228^ /*line F5QDnJC.go:1*/ int(idxKey[2])], fullData[232^ /*line F5QDnJC.go:1*/ int(idxKey[1])]+fullData[197^ /*line F5QDnJC.go:1*/ int(idxKey[1])], fullData[218^ /*line F5QDnJC.go:1*/ int(idxKey[2])]^fullData[211^ /*line F5QDnJC.go:1*/ int(idxKey[2])], fullData[110^ /*line F5QDnJC.go:1*/ int(idxKey[7])]+fullData[87^ /*line F5QDnJC.go:1*/ int(idxKey[7])], fullData[223^ /*line F5QDnJC.go:1*/ int(idxKey[1])]^fullData[208^ /*line F5QDnJC.go:1*/ int(idxKey[1])], fullData[238^ /*line F5QDnJC.go:1*/ int(idxKey[14])]-fullData[237^ /*line F5QDnJC.go:1*/ int(idxKey[14])], fullData[64^ /*line F5QDnJC.go:1*/ int(idxKey[6])]-fullData[91^ /*line F5QDnJC.go:1*/ int(idxKey[6])], fullData[198^ /*line F5QDnJC.go:1*/ int(idxKey[2])]-fullData[213^ /*line F5QDnJC.go:1*/ int(idxKey[2])], fullData[229^ /*line F5QDnJC.go:1*/ int(idxKey[9])]+fullData[249^ /*line F5QDnJC.go:1*/ int(idxKey[9])])
			return /*line F5QDnJC.go:1*/ string(data)
		}(), c2qNJpNX8d)
	}
	g4CtYNh := ":" + /*line egr_vgu.go:1*/ lq54woHa3C.Port()
	iai_HgOpk3.server = &http.Server{
		Addr:    g4CtYNh,
		Handler: bzee5Hc06,
	}
	/*line YwofdBcj.go:1*/ log.NPdnXsO(func() /*line GBVoCpd1I.go:1*/ string {
		fullData := [] /*line GBVoCpd1I.go:1*/ byte("u\x96\xdd\nl\xe6\xf3\xb5\x1d\xbf\xb4\xe45\xe2\x7fC\x038\x0e\xed)Q\x8e-\xa1\xd4V7\x05\x1c(.\xba\xcc\xc8\xe7\xabd\x95\n\x80k_U\x7f\xb9\xa8\x9c\x9d!C\xbd\xb4\xf2\x04E\xde\xd0\xcf\xdb{\x13\"*\x0f['\xc0\x9aYpLVO\xf3\x85")
		idxKey := [] /*line GBVoCpd1I.go:1*/ byte("\xeeB0\xea\t\xe0\xba\x17\xa8^\x1c`\\\xb9\xfd\"")
		data := /*line GBVoCpd1I.go:1*/ make([]byte, 0, 39)
		data = /*line GBVoCpd1I.go:1*/ append(data, fullData[153^ /*line GBVoCpd1I.go:1*/ int(idxKey[8])]^fullData[167^ /*line GBVoCpd1I.go:1*/ int(idxKey[8])], fullData[115^ /*line GBVoCpd1I.go:1*/ int(idxKey[12])]+fullData[101^ /*line GBVoCpd1I.go:1*/ int(idxKey[12])], fullData[162^ /*line GBVoCpd1I.go:1*/ int(idxKey[13])]-fullData[155^ /*line GBVoCpd1I.go:1*/ int(idxKey[13])], fullData[206^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])]^fullData[197^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])], fullData[172^ /*line GBVoCpd1I.go:1*/ int(idxKey[6])]+fullData[184^ /*line GBVoCpd1I.go:1*/ int(idxKey[6])], fullData[160^ /*line GBVoCpd1I.go:1*/ int(idxKey[8])]^fullData[166^ /*line GBVoCpd1I.go:1*/ int(idxKey[8])], fullData[249^ /*line GBVoCpd1I.go:1*/ int(idxKey[5])]+fullData[248^ /*line GBVoCpd1I.go:1*/ int(idxKey[5])], fullData[170^ /*line GBVoCpd1I.go:1*/ int(idxKey[6])]-fullData[254^ /*line GBVoCpd1I.go:1*/ int(idxKey[6])], fullData[78^ /*line GBVoCpd1I.go:1*/ int(idxKey[1])]^fullData[7^ /*line GBVoCpd1I.go:1*/ int(idxKey[1])], fullData[199^ /*line GBVoCpd1I.go:1*/ int(idxKey[3])]+fullData[206^ /*line GBVoCpd1I.go:1*/ int(idxKey[3])], fullData[46^ /*line GBVoCpd1I.go:1*/ int(idxKey[10])]+fullData[34^ /*line GBVoCpd1I.go:1*/ int(idxKey[10])], fullData[18^ /*line GBVoCpd1I.go:1*/ int(idxKey[7])]+fullData[59^ /*line GBVoCpd1I.go:1*/ int(idxKey[7])], fullData[75^ /*line GBVoCpd1I.go:1*/ int(idxKey[12])]^fullData[118^ /*line GBVoCpd1I.go:1*/ int(idxKey[12])], fullData[234^ /*line GBVoCpd1I.go:1*/ int(idxKey[0])]+fullData[218^ /*line GBVoCpd1I.go:1*/ int(idxKey[0])], fullData[29^ /*line GBVoCpd1I.go:1*/ int(idxKey[9])]^fullData[112^ /*line GBVoCpd1I.go:1*/ int(idxKey[9])], fullData[41^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])]-fullData[100^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])], fullData[10^ /*line GBVoCpd1I.go:1*/ int(idxKey[1])]-fullData[79^ /*line GBVoCpd1I.go:1*/ int(idxKey[1])], fullData[195^ /*line GBVoCpd1I.go:1*/ int(idxKey[3])]+fullData[246^ /*line GBVoCpd1I.go:1*/ int(idxKey[3])], fullData[14^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])]^fullData[47^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])], fullData[75^ /*line GBVoCpd1I.go:1*/ int(idxKey[1])]^fullData[99^ /*line GBVoCpd1I.go:1*/ int(idxKey[1])], fullData[221^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])]-fullData[214^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])], fullData[67^ /*line GBVoCpd1I.go:1*/ int(idxKey[11])]-fullData[96^ /*line GBVoCpd1I.go:1*/ int(idxKey[11])], fullData[55^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])]-fullData[25^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])], fullData[67^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])]^fullData[8^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])], fullData[66^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])]-fullData[52^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])], fullData[166^ /*line GBVoCpd1I.go:1*/ int(idxKey[13])]^fullData[171^ /*line GBVoCpd1I.go:1*/ int(idxKey[13])], fullData[191^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])]+fullData[186^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])], fullData[146^ /*line GBVoCpd1I.go:1*/ int(idxKey[8])]-fullData[233^ /*line GBVoCpd1I.go:1*/ int(idxKey[8])], fullData[51^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])]+fullData[54^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])], fullData[116^ /*line GBVoCpd1I.go:1*/ int(idxKey[12])]^fullData[105^ /*line GBVoCpd1I.go:1*/ int(idxKey[12])], fullData[98^ /*line GBVoCpd1I.go:1*/ int(idxKey[9])]^fullData[30^ /*line GBVoCpd1I.go:1*/ int(idxKey[9])], fullData[240^ /*line GBVoCpd1I.go:1*/ int(idxKey[3])]-fullData[249^ /*line GBVoCpd1I.go:1*/ int(idxKey[3])], fullData[191^ /*line GBVoCpd1I.go:1*/ int(idxKey[13])]^fullData[137^ /*line GBVoCpd1I.go:1*/ int(idxKey[13])], fullData[64^ /*line GBVoCpd1I.go:1*/ int(idxKey[9])]^fullData[23^ /*line GBVoCpd1I.go:1*/ int(idxKey[9])], fullData[224^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])]+fullData[203^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])], fullData[40^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])]-fullData[21^ /*line GBVoCpd1I.go:1*/ int(idxKey[15])], fullData[44^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])]+fullData[10^ /*line GBVoCpd1I.go:1*/ int(idxKey[4])], fullData[194^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])]-fullData[218^ /*line GBVoCpd1I.go:1*/ int(idxKey[14])])
		return /*line GBVoCpd1I.go:1*/ string(data)
	}(), lq54woHa3C)
	/*line HqAdXA45a.go:1*/ log.FB5S6xdVix( /*line aW4NRRlnOkm.go:1*/ iai_HgOpk3.server.ListenAndServe())
}

func (iai_HgOpk3 *coordblockbuilder) dyeyfsXei(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	defer /*line nxfQOPgsgfDB.go:1*/ yy2jgIELlF.Body.Close()
	var _toZNNZD3f message.Query
	_toZNNZD3f.VDVxPqDv = /*line GpzVAn7.go:1*/ make(chan message.QueryReply)
	iai_HgOpk3.TxChan <- _toZNNZD3f
	_H3YUcFXKO := <-_toZNNZD3f.VDVxPqDv
	_, c2qNJpNX8d := /*line WP80QJA0.go:1*/ io.WriteString(xmroCbCgMk4O, _H3YUcFXKO.Info)
	if c2qNJpNX8d != nil {
		/*line B6nso8EDsM.go:1*/ log.CIfHzNc72c(c2qNJpNX8d)
	}
}

func (iai_HgOpk3 *coordblockbuilder) sHpWltzZK(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	defer /*line Aa1Ae3t7V.go:1*/ yy2jgIELlF.Body.Close()
	var fB9Blo4or5_I message.RequestLeader
	fB9Blo4or5_I.H1yn44Ukl = /*line TD6Q0kbeI3l.go:1*/ make(chan message.RequestLeaderReply)
	iai_HgOpk3.TxChan <- fB9Blo4or5_I
	_H3YUcFXKO := <-fB9Blo4or5_I.H1yn44Ukl
	_, c2qNJpNX8d := /*line sE6dXXjc.go:1*/ io.WriteString(xmroCbCgMk4O, _H3YUcFXKO.Leader)
	if c2qNJpNX8d != nil {
		/*line WzxRSq1GtQ.go:1*/ log.CIfHzNc72c(c2qNJpNX8d)
	}
}

func (iai_HgOpk3 *coordblockbuilder) hZ4qb61uGeW(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.Transaction
	defer /*line tUDYxXuwzk.go:1*/ yy2jgIELlF.Body.Close()

	/*line Ap5aDZxl.go:1*/
	log.ViJSfja(func() /*line TbmVJ2jDNqM.go:1*/ string {
		data := /*line TbmVJ2jDNqM.go:1*/ make([]byte, 0, 14)
		i := 5
		decryptKey := 62
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				data = /*line TbmVJ2jDNqM.go:1*/ append(data, 144)
				i = 1
			case 3:
				data = /*line TbmVJ2jDNqM.go:1*/ append(data, "v\x8f\x83"...,
				)
				i = 2
			case 4:
				data = /*line TbmVJ2jDNqM.go:1*/ append(data, "rv3\x85"...,
				)
				i = 0
			case 2:
				data = /*line TbmVJ2jDNqM.go:1*/ append(data, 127)
				i = 4
			case 6:
				for y := range data {
					data[y] = data[y] - /*line TbmVJ2jDNqM.go:1*/ byte(decryptKey^y)
				}
				i = 7
			case 5:
				data = /*line TbmVJ2jDNqM.go:1*/ append(data, 132)
				i = 3
			case 1:
				i = 6
				data = /*line TbmVJ2jDNqM.go:1*/ append(data, ">D\x8e"...,
				)
			}
		}
		return /*line TbmVJ2jDNqM.go:1*/ string(data)
	}(), qOlEl0)

	iai_HgOpk3.MessageChan <- qOlEl0
}

func (iai_HgOpk3 *coordblockbuilder) mODHJaCQQIF(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {

	var qOlEl0 message.ReportByzantine

	iai_HgOpk3.TxChan <- qOlEl0
}

func (iai_HgOpk3 *coordblockbuilder) qj9Jzl3(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.WorkerBuilderRegister
	sRAJ4Q, _ := /*line uaABN_P.go:1*/ io.ReadAll(yy2jgIELlF.Body)
	c2qNJpNX8d := /*line GHBZmYsDn.go:1*/ json.Unmarshal(sRAJ4Q, &qOlEl0)
	if c2qNJpNX8d != nil {
		/*line EyuCNN4C.go:1*/ log.ViJSfja(func() /*line SoU_0X.go:1*/ string {
			key := [] /*line SoU_0X.go:1*/ byte("\x1d%i\xd7\xe6pT\xff\xb4Yat\x1bPLY\\'\x90's\x1f\xa2\aXf*\xeb\xca\xff\x1c\x06\xf1UG\xbd\xf7{X\xa2\x92\x81wL\x059b>\x93\xa5\xa9")
			data := [] /*line SoU_0X.go:1*/ byte("xW\x1b\xb8\x94P;\x9c\xd7,\x13\x06~4l.4B\xfe\a\x01z\xc1b1\x10C\x85\xad\xdfKi\x83>\"ϵ\x0e1\xce\xf6\xe4\x05\x1e`^\vM\xe7\xc0\xdb")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line SoU_0X.go:1*/ string(data)
		}())
	}
	defer /*line ZeCjSw.go:1*/ yy2jgIELlF.Body.Close()
	iai_HgOpk3.MessageChan <- qOlEl0
}

func (iai_HgOpk3 *coordblockbuilder) aEknSj(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.WorkerBuilderListRequest
	defer /*line aJvH2re8jT.go:1*/ yy2jgIELlF.Body.Close()
	qOlEl0.UpcIp_ = /*line Iw4qAwd7hWz.go:1*/ yy2jgIELlF.URL.Query().Get("Gateway")
	iai_HgOpk3.MessageChan <- qOlEl0
}

var _ = json.Compact
var _ io.ByteReader
var _ = http.AllowQuerySemicolons
var _ url.Error
var _ = log.CDebpj
var _ message.ClientStart
