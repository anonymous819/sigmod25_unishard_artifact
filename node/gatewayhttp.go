//line :1
package node

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	log "unishard/log"
	message "unishard/message"
	types "unishard/types"
)

func (mM6kT_d *gatewaynode) roWnZ9bUZa() {
	bzee5Hc06 := /*line i6YnJMdCeaSS.go:1*/ http.NewServeMux()
	/*line ua9G6P.go:1*/ bzee5Hc06.HandleFunc("/", mM6kT_d.hZ4qb61uGeW)
	/*line toGspJfod.go:1*/ bzee5Hc06.HandleFunc(func() /*line zW6Q4QIr_Q.go:1*/ string {
		seed := /*line zW6Q4QIr_Q.go:1*/ byte(185)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line zW6Q4QIr_Q.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line zW6Q4QIr_Q.go:1*/ fnc(118)(40)(24)(3)(249)(250)(13)(225)(21)(249)(17)(243)(255)(242)(11)(17)(247)(248)(1)(10)(249)(237)(19)(2)(2)(10)(1)(241)(13)(224)(19)(14)(253)(255)(255)(5)(242)
		return /*line zW6Q4QIr_Q.go:1*/ string(data)
	}(), mM6kT_d.e1WUIv)
	/*line uqDmlfe.go:1*/ bzee5Hc06.HandleFunc(func() /*line Ch5UyjcP.go:1*/ string {
		fullData := [] /*line Ch5UyjcP.go:1*/ byte("\xa7\xd6b\x8a\xae\xac.\xd4NV\xf0:JJ\xdf:\xf2c?\xb9\xfd@\x1d\xd7;\xcdʭ\x80\xeb\xc2M\xdf\x11L9\xb3i=\xa9^s%\xa6P\x11[2\xb7}\xbb\x18")
		idxKey := [] /*line Ch5UyjcP.go:1*/ byte("%\xd1")
		data := /*line Ch5UyjcP.go:1*/ make([]byte, 0, 27)
		data = /*line Ch5UyjcP.go:1*/ append(data, fullData[224^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]-fullData[217^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[199^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]^fullData[220^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[194^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]^fullData[208^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[212^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]-fullData[218^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[34^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]-fullData[0^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[9^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]-fullData[56^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[226^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]-fullData[250^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[52^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]+fullData[43^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[41^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]^fullData[55^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[15^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]^fullData[7^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[5^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]-fullData[12^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[198^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]^fullData[245^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[209^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]^fullData[207^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[53^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]-fullData[57^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[240^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]+fullData[201^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[254^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]^fullData[255^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[203^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]+fullData[246^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[58^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]^fullData[6^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[39^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]+fullData[47^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[23^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]-fullData[44^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[202^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]-fullData[222^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[247^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]-fullData[200^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[252^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])]+fullData[249^ /*line Ch5UyjcP.go:1*/ int(idxKey[1])], fullData[35^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]^fullData[48^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[49^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]-fullData[38^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])], fullData[21^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])]+fullData[33^ /*line Ch5UyjcP.go:1*/ int(idxKey[0])])
		return /*line Ch5UyjcP.go:1*/ string(data)
	}(), mM6kT_d.l7EaUZFpb)

	lq54woHa3C, c2qNJpNX8d := /*line E64kZqJxud.go:1*/ url.Parse(func() /*line z56fXJ.go:1*/ string {
		seed := /*line z56fXJ.go:1*/ byte(255)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line z56fXJ.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line z56fXJ.go:1*/ fnc(105)(12)(0)(252)(202)(245)(0)(2)(1)(5)(247)(2)(254)(2)(254)(3)(9)(254)(248)(0)(0)
		return /*line z56fXJ.go:1*/ string(data)
	}())
	if c2qNJpNX8d != nil {
		/*line uVSV7TQrEOmi.go:1*/ log.FB5S6xdVix(func() /*line jfN2T6ob.go:1*/ string {
			key := [] /*line jfN2T6ob.go:1*/ byte("Ձ5\xac\f<\xaa2\xbc\x0e\x18r\x1fʊ'\x0f\xb8\xadi\x87\x80")
			data := [] /*line jfN2T6ob.go:1*/ byte("\x93\xf3?\xc4\x149\xc8:dbI\x00T\x9b\x96>c\xba\xc2\t\xb3\xa0")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line jfN2T6ob.go:1*/ string(data)
		}(), c2qNJpNX8d)
	}
	g4CtYNh := ":" + /*line ejAUtGQ0NV.go:1*/ lq54woHa3C.Port()
	mM6kT_d.server = &http.Server{
		Addr:    g4CtYNh,
		Handler: bzee5Hc06,
	}
	/*line eTCwHkQdmda.go:1*/ log.ZWwzboe(func() /*line YGUHw_GPNbz8.go:1*/ string {
		data := /*line YGUHw_GPNbz8.go:1*/ make([]byte, 0, 33)
		i := 8
		decryptKey := 64
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 11
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "ziw\x90"...,
				)
			case 8:
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "Ql|n"...,
				)
				i = 12
			case 3:
				i = 6
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "9\x91"...,
				)
			case 9:
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "\x86|~"...,
				)
				i = 7
			case 2:
				i = 1
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "wtq&"...,
				)
			case 11:
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "\x80\x8a"...,
				)
				i = 3
			case 4:
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "\x86\x825"...,
				)
				i = 0
			case 0:
				for y := range data {
					data[y] = data[y] - /*line YGUHw_GPNbz8.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 7:
				i = 4
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "x6"...,
				)
			case 12:
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "\x85p\x85"...,
				)
				i = 10
			case 6:
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "\x93}\x8f"...,
				)
				i = 9
			case 10:
				i = 2
				data = /*line YGUHw_GPNbz8.go:1*/ append(data, "-j"...,
				)
			}
		}
		return /*line YGUHw_GPNbz8.go:1*/ string(data)
	}(), g4CtYNh)
	/*line qjleS_lgmNL_.go:1*/ log.FB5S6xdVix( /*line Ghovrp77DH5.go:1*/ mM6kT_d.server.ListenAndServe())
}

func (mM6kT_d *gatewaynode) hZ4qb61uGeW(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.TransactionForm
	defer /*line DKkxua.go:1*/ yy2jgIELlF.Body.Close()
	sRAJ4Q, _ := /*line ZwiQa_Yi5ba.go:1*/ io.ReadAll(yy2jgIELlF.Body)
	if c2qNJpNX8d := /*line AHagE5beBap.go:1*/ json.Unmarshal(sRAJ4Q, &qOlEl0); c2qNJpNX8d != nil {
		/*line XdQwBMG7.go:1*/ log.Jp980o6YjM(func() /*line DS4CTlO.go:1*/ string {
			key := [] /*line DS4CTlO.go:1*/ byte("pD$u\x18\xfb\xd4t?\xe9/\xe4\xadI\xffD\x0f&nq\xd3\x18\x1a-+\x87x\xe2\xe8\xf8M2\xf3\xfc\x94VX^3\x9a\xda")
			data := [] /*line DS4CTlO.go:1*/ byte("ն\x96\xe4\x8a\x1bCע^\xa1V\x12\xad\x1f\xbbw\x8bܑE}}\x92\x94\xfd\xe1PO\x18\xc1\xa4Tj\a\xb7\xbbҜ\tH")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line DS4CTlO.go:1*/ string(data)
		}())
	}

	mM6kT_d.MessageChan <- qOlEl0
}

func (mM6kT_d *gatewaynode) e1WUIv(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.WorkerSharedVariableRegisterResponse
	defer /*line Do1rPKT_0CE.go:1*/ yy2jgIELlF.Body.Close()
	cUffk93, _ := /*line IgGzcYS.go:1*/ strconv.Atoi( /*line u0rEPVAnaOU.go:1*/ yy2jgIELlF.URL.Query().Get(func() /*line DubWbPamDf5t.go:1*/ string {
		key := [] /*line DubWbPamDf5t.go:1*/ byte("V\x03\x90\x9bu\xf8\x9fF\xf7W\xb0")
		data := [] /*line DubWbPamDf5t.go:1*/ byte("\x05f\xfe\xff\x10\x8a\xcc.\x96%\xd4")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line DubWbPamDf5t.go:1*/ string(data)
	}()))
	uLlOFc := /*line ozBxA8jGeQuB.go:1*/ yy2jgIELlF.URL.Query().Get(func() /*line Txjxrqbaz9.go:1*/ string {
		seed := /*line Txjxrqbaz9.go:1*/ byte(81)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line Txjxrqbaz9.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line Txjxrqbaz9.go:1*/ fnc(167)(89)(195)(125)(242)(229)(212)(161)(80)
		return /*line Txjxrqbaz9.go:1*/ string(data)
	}())
	qOlEl0.PI77Zg94 = /*line Jint_F.go:1*/ types.Shard(cUffk93)
	qOlEl0.QxT5nG3X5Gn4 = uLlOFc
	mM6kT_d.MessageChan <- qOlEl0
}
func (mM6kT_d *gatewaynode) l7EaUZFpb(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.WorkerBuilderListResponse
	defer /*line HWSi8ncIkNz.go:1*/ yy2jgIELlF.Body.Close()
	sRAJ4Q, _ := /*line hOzuXB2vb3.go:1*/ io.ReadAll(yy2jgIELlF.Body)
	c2qNJpNX8d := /*line hNSVaDz.go:1*/ json.Unmarshal(sRAJ4Q, &qOlEl0)
	if c2qNJpNX8d != nil {
		/*line zNtVSTR.go:1*/ log.ViJSfja(func() /*line xLIewfWFP6HQ.go:1*/ string {
			key := [] /*line xLIewfWFP6HQ.go:1*/ byte("p\xa5\xae\x8e\xff9Y\xdeiC\u0557\x12\x89LXl7w\xaf;hx\x99\\\xcd\xd5F\xe3\xe6^c\xc6%P\xf1t{=00\xc2>\xb7\x87r{Σ\xba\x16)f\xc7T")
			data := [] /*line xLIewfWFP6HQ.go:1*/ byte("\xf5\xcd\xc4\xe1s\xe7\x16\x85\xfa2\x9d\xdbS\xdb\xd4\x1f\xfc.\xf7q7\xfd\xeb\xcc\r\xa9\x94(\x84:\xf9\f\xacF\x15\x81\xce\xfa,<4\xa34\x95\xe2\x01\xf9\x84¹ZF\b\xac\x11")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line xLIewfWFP6HQ.go:1*/ string(data)
		}())
	}
	mM6kT_d.MessageChan <- qOlEl0
}

var _ = json.Compact
var _ io.ByteReader
var _ = http.AllowQuerySemicolons
var _ url.Error
var _ = strconv.AppendBool
var _ = log.CDebpj
var _ message.ClientStart

const _ = types.ABORT
