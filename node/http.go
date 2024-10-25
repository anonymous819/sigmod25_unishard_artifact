//line :1
package node

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	log "unishard/log"
	message "unishard/message"
	utils "unishard/utils"
)

const (
	HTTPClientID  = "Id"
	HTTPCommandID = "Cid"
)

func (eL4EyO3pzuT *node) roWnZ9bUZa() {
	bzee5Hc06 := /*line gxOwWOhu.go:1*/ http.NewServeMux()
	/*line Jqa7c5Zq9jU.go:1*/ bzee5Hc06.HandleFunc("/", eL4EyO3pzuT.hZ4qb61uGeW)
	/*line CD2nWdTP_p.go:1*/ bzee5Hc06.HandleFunc("/query", eL4EyO3pzuT.dyeyfsXei)
	/*line ax5fpoC6g.go:1*/ bzee5Hc06.HandleFunc(func() /*line xE3vjzIW.go:1*/ string {
		fullData := [] /*line xE3vjzIW.go:1*/ byte("\xfa@urf\x04\xf0]\xd5}\xf4\xe8\x0e\fo\x11:j\x90l\xc5\xf5\t\\\x91-\x17\xa5")
		idxKey := [] /*line xE3vjzIW.go:1*/ byte("%(KA\x88q>")
		data := /*line xE3vjzIW.go:1*/ make([]byte, 0, 15)
		data = /*line xE3vjzIW.go:1*/ append(data, fullData[52^ /*line xE3vjzIW.go:1*/ int(idxKey[6])]-fullData[42^ /*line xE3vjzIW.go:1*/ int(idxKey[6])], fullData[48^ /*line xE3vjzIW.go:1*/ int(idxKey[0])]+fullData[44^ /*line xE3vjzIW.go:1*/ int(idxKey[0])], fullData[62^ /*line xE3vjzIW.go:1*/ int(idxKey[0])]-fullData[36^ /*line xE3vjzIW.go:1*/ int(idxKey[0])], fullData[102^ /*line xE3vjzIW.go:1*/ int(idxKey[5])]^fullData[104^ /*line xE3vjzIW.go:1*/ int(idxKey[5])], fullData[48^ /*line xE3vjzIW.go:1*/ int(idxKey[6])]-fullData[62^ /*line xE3vjzIW.go:1*/ int(idxKey[6])], fullData[93^ /*line xE3vjzIW.go:1*/ int(idxKey[2])]^fullData[88^ /*line xE3vjzIW.go:1*/ int(idxKey[2])], fullData[74^ /*line xE3vjzIW.go:1*/ int(idxKey[3])]-fullData[67^ /*line xE3vjzIW.go:1*/ int(idxKey[3])], fullData[45^ /*line xE3vjzIW.go:1*/ int(idxKey[1])]-fullData[58^ /*line xE3vjzIW.go:1*/ int(idxKey[1])], fullData[34^ /*line xE3vjzIW.go:1*/ int(idxKey[0])]-fullData[42^ /*line xE3vjzIW.go:1*/ int(idxKey[0])], fullData[152^ /*line xE3vjzIW.go:1*/ int(idxKey[4])]-fullData[128^ /*line xE3vjzIW.go:1*/ int(idxKey[4])], fullData[77^ /*line xE3vjzIW.go:1*/ int(idxKey[2])]^fullData[83^ /*line xE3vjzIW.go:1*/ int(idxKey[2])], fullData[71^ /*line xE3vjzIW.go:1*/ int(idxKey[2])]^fullData[90^ /*line xE3vjzIW.go:1*/ int(idxKey[2])], fullData[139^ /*line xE3vjzIW.go:1*/ int(idxKey[4])]^fullData[146^ /*line xE3vjzIW.go:1*/ int(idxKey[4])], fullData[76^ /*line xE3vjzIW.go:1*/ int(idxKey[3])]+fullData[69^ /*line xE3vjzIW.go:1*/ int(idxKey[3])])
		return /*line xE3vjzIW.go:1*/ string(data)
	}(), eL4EyO3pzuT.sHpWltzZK)
	/*line vXLavKc.go:1*/ bzee5Hc06.HandleFunc(func() /*line aCtyVI.go:1*/ string {
		data := /*line aCtyVI.go:1*/ make([]byte, 0, 17)
		i := 5
		decryptKey := 123
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				data = /*line aCtyVI.go:1*/ append(data, 20)
				i = 2
			case 2:
				data = /*line aCtyVI.go:1*/ append(data, "\b\f\x06"...,
				)
				i = 7
			case 4:
				data = /*line aCtyVI.go:1*/ append(data, 31)
				i = 8
			case 7:
				i = 6
				for y := range data {
					data[y] = data[y] ^ /*line aCtyVI.go:1*/ byte(decryptKey^y)
				}
			case 0:
				data = /*line aCtyVI.go:1*/ append(data, "\x1e)\x1d"...,
				)
				i = 4
			case 8:
				data = /*line aCtyVI.go:1*/ append(data, "\a\t"...,
				)
				i = 1
			case 3:
				i = 0
				data = /*line aCtyVI.go:1*/ append(data, "\a\x1b"...,
				)
			case 5:
				data = /*line aCtyVI.go:1*/ append(data, "C\x1f\v\x1f"...,
				)
				i = 3
			}
		}
		return /*line aCtyVI.go:1*/ string(data)
	}(), eL4EyO3pzuT.mODHJaCQQIF)

	lq54woHa3C, c2qNJpNX8d := /*line OMsLJxSHJ.go:1*/ url.Parse(func() /*line SqbEKlgp.go:1*/ string {
		seed := /*line SqbEKlgp.go:1*/ byte(211)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line SqbEKlgp.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line SqbEKlgp.go:1*/ fnc(59)(130)(4)(4)(210)(153)(50)(102)(205)(159)(53)(108)(214)(174)(90)(183)(119)
		return /*line SqbEKlgp.go:1*/ string(data)
	}() + /*line TyVzhQ8mz.go:1*/ strconv.Itoa(9000+ /*line vuUyd6PuMgv7.go:1*/ utils.BD1ZTohx( /*line oZh7ITLLloMN.go:1*/ eL4EyO3pzuT.ID())+ /*line IiN9IWjrjy.go:1*/ int(eL4EyO3pzuT.shard)*100))
	if c2qNJpNX8d != nil {
		/*line H1_q9qaQNDS0.go:1*/ log.FB5S6xdVix(func() /*line Sos17sN.go:1*/ string {
			fullData := [] /*line Sos17sN.go:1*/ byte("aA,\xe0Ox&8\xc7\xfc\xaaŹ\x89\x93\xd1\xfa\x80\x15\x10r\xfd\xddiU\x02\x1e\xfbH\xe3v\xf04c\xfc\xb0\x0e\x12\x92\xb2\xc3qa\x88")
			idxKey := [] /*line Sos17sN.go:1*/ byte("\a\xfc\xbeN\xf5\xbd\xcev\xf8\xa0")
			data := /*line Sos17sN.go:1*/ make([]byte, 0, 23)
			data = /*line Sos17sN.go:1*/ append(data, fullData[173^ /*line Sos17sN.go:1*/ int(idxKey[2])]^fullData[187^ /*line Sos17sN.go:1*/ int(idxKey[2])], fullData[162^ /*line Sos17sN.go:1*/ int(idxKey[9])]+fullData[188^ /*line Sos17sN.go:1*/ int(idxKey[9])], fullData[160^ /*line Sos17sN.go:1*/ int(idxKey[9])]^fullData[178^ /*line Sos17sN.go:1*/ int(idxKey[9])], fullData[111^ /*line Sos17sN.go:1*/ int(idxKey[7])]^fullData[98^ /*line Sos17sN.go:1*/ int(idxKey[7])], fullData[177^ /*line Sos17sN.go:1*/ int(idxKey[2])]+fullData[186^ /*line Sos17sN.go:1*/ int(idxKey[2])], fullData[111^ /*line Sos17sN.go:1*/ int(idxKey[3])]+fullData[107^ /*line Sos17sN.go:1*/ int(idxKey[3])], fullData[172^ /*line Sos17sN.go:1*/ int(idxKey[5])]-fullData[153^ /*line Sos17sN.go:1*/ int(idxKey[5])], fullData[238^ /*line Sos17sN.go:1*/ int(idxKey[6])]+fullData[201^ /*line Sos17sN.go:1*/ int(idxKey[6])], fullData[243^ /*line Sos17sN.go:1*/ int(idxKey[4])]+fullData[229^ /*line Sos17sN.go:1*/ int(idxKey[4])], fullData[9^ /*line Sos17sN.go:1*/ int(idxKey[0])]^fullData[26^ /*line Sos17sN.go:1*/ int(idxKey[0])], fullData[209^ /*line Sos17sN.go:1*/ int(idxKey[6])]+fullData[231^ /*line Sos17sN.go:1*/ int(idxKey[6])], fullData[182^ /*line Sos17sN.go:1*/ int(idxKey[2])]-fullData[166^ /*line Sos17sN.go:1*/ int(idxKey[2])], fullData[109^ /*line Sos17sN.go:1*/ int(idxKey[3])]^fullData[102^ /*line Sos17sN.go:1*/ int(idxKey[3])], fullData[183^ /*line Sos17sN.go:1*/ int(idxKey[9])]+fullData[169^ /*line Sos17sN.go:1*/ int(idxKey[9])], fullData[181^ /*line Sos17sN.go:1*/ int(idxKey[9])]-fullData[182^ /*line Sos17sN.go:1*/ int(idxKey[9])], fullData[226^ /*line Sos17sN.go:1*/ int(idxKey[8])]-fullData[244^ /*line Sos17sN.go:1*/ int(idxKey[8])], fullData[156^ /*line Sos17sN.go:1*/ int(idxKey[2])]+fullData[160^ /*line Sos17sN.go:1*/ int(idxKey[2])], fullData[67^ /*line Sos17sN.go:1*/ int(idxKey[3])]^fullData[85^ /*line Sos17sN.go:1*/ int(idxKey[3])], fullData[170^ /*line Sos17sN.go:1*/ int(idxKey[9])]^fullData[171^ /*line Sos17sN.go:1*/ int(idxKey[9])], fullData[77^ /*line Sos17sN.go:1*/ int(idxKey[3])]+fullData[104^ /*line Sos17sN.go:1*/ int(idxKey[3])], fullData[211^ /*line Sos17sN.go:1*/ int(idxKey[8])]+fullData[223^ /*line Sos17sN.go:1*/ int(idxKey[8])], fullData[249^ /*line Sos17sN.go:1*/ int(idxKey[8])]^fullData[210^ /*line Sos17sN.go:1*/ int(idxKey[8])])
			return /*line Sos17sN.go:1*/ string(data)
		}(), c2qNJpNX8d)
	}
	g4CtYNh := ":" + /*line FBn0LriM.go:1*/ lq54woHa3C.Port()
	eL4EyO3pzuT.server = &http.Server{
		Addr:    g4CtYNh,
		Handler: bzee5Hc06,
	}
	/*line Sl9yWtRp.go:1*/ log.ZWwzboe(func() /*line QN34mpava10L.go:1*/ string {
		seed := /*line QN34mpava10L.go:1*/ byte(253)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line QN34mpava10L.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line QN34mpava10L.go:1*/ fnc(107)(12)(0)(252)(176)(83)(242)(13)(4)(239)(13)(174)(83)(1)(237)(17)(2)(245)(5)(249)(185)(79)(255)(178)
		return /*line QN34mpava10L.go:1*/ string(data)
	}(), g4CtYNh)
	/*line G5Z92oD4Jw.go:1*/ log.FB5S6xdVix( /*line X6UZLJTw.go:1*/ eL4EyO3pzuT.server.ListenAndServe())
}

func (eL4EyO3pzuT *node) dyeyfsXei(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	defer /*line eIgoetIk.go:1*/ yy2jgIELlF.Body.Close()
	var _toZNNZD3f message.Query
	_toZNNZD3f.VDVxPqDv = /*line n7AqZKMsMDI.go:1*/ make(chan message.QueryReply)
	eL4EyO3pzuT.TxChan <- _toZNNZD3f
	_H3YUcFXKO := <-_toZNNZD3f.VDVxPqDv
	_, c2qNJpNX8d := /*line dWQ6wg.go:1*/ io.WriteString(xmroCbCgMk4O, _H3YUcFXKO.Info)
	if c2qNJpNX8d != nil {
		/*line hRt7Qf3S6m.go:1*/ log.CIfHzNc72c(c2qNJpNX8d)
	}
}

func (eL4EyO3pzuT *node) sHpWltzZK(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	defer /*line FZlt5S.go:1*/ yy2jgIELlF.Body.Close()
	var fB9Blo4or5_I message.RequestLeader
	fB9Blo4or5_I.H1yn44Ukl = /*line z1BwGk.go:1*/ make(chan message.RequestLeaderReply)
	eL4EyO3pzuT.TxChan <- fB9Blo4or5_I
	_H3YUcFXKO := <-fB9Blo4or5_I.H1yn44Ukl
	_, c2qNJpNX8d := /*line SarilOsIa_o.go:1*/ io.WriteString(xmroCbCgMk4O, _H3YUcFXKO.Leader)
	if c2qNJpNX8d != nil {
		/*line iBi9lYe.go:1*/ log.CIfHzNc72c(c2qNJpNX8d)
	}
}

func (eL4EyO3pzuT *node) hZ4qb61uGeW(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {
	var qOlEl0 message.Transaction
	defer /*line D4eHAT.go:1*/ yy2jgIELlF.Body.Close()

	sRAJ4Q, _ := /*line BKkIcrt.go:1*/ io.ReadAll(yy2jgIELlF.Body)
	/*line AuAW0BaV7.go:1*/ log.Debugf(func() /*line HL4FaD1yTuMd.go:1*/ string {
		data := [] /*line HL4FaD1yTuMd.go:1*/ byte("\xb4\xb5v\xc1 \x1fg\xac\xdeo(\x8f is \xcex")
		positions := [...]byte{7, 6, 0, 10, 3, 5, 1, 6, 16, 11, 7, 8, 1, 7, 7, 0, 3, 8}
		for i := 0; i < 18; i += 2 {
			localKey := /*line HL4FaD1yTuMd.go:1*/ byte(i) + /*line HL4FaD1yTuMd.go:1*/ byte(positions[i]^positions[i+1]) + 71
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line HL4FaD1yTuMd.go:1*/ string(data)
	}(), eL4EyO3pzuT.id, sRAJ4Q)
	eL4EyO3pzuT.TxChan <- qOlEl0
}

func (eL4EyO3pzuT *node) mODHJaCQQIF(xmroCbCgMk4O http.ResponseWriter, yy2jgIELlF *http.Request) {

	var qOlEl0 message.ReportByzantine

	eL4EyO3pzuT.TxChan <- qOlEl0
}

var _ io.ByteReader
var _ = http.AllowQuerySemicolons
var _ url.Error
var _ = strconv.AppendBool
var _ = log.CDebpj
var _ message.ClientStart
var _ = utils.GffGNE
