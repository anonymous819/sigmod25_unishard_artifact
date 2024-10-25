//line :1
package transport

import (
	"bytes"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"
	"sync"

	log "unishard/log"
)

var BCNGV6 = /*line t7zCzo.go:1*/ flag.String(func() /*line voIV1XyT09.go:1*/ string {
	fullData := [] /*line voIV1XyT09.go:1*/ byte("<\xbb\r\x98/>#\xf7|p5\xb0\x02\xae\xf6\xc2c\x1e")
	idxKey := [] /*line voIV1XyT09.go:1*/ byte("h\x8f\x98\xca")
	data := /*line voIV1XyT09.go:1*/ make([]byte, 0, 10)
	data = /*line voIV1XyT09.go:1*/ append(data, fullData[132^ /*line voIV1XyT09.go:1*/ int(idxKey[1])]-fullData[143^ /*line voIV1XyT09.go:1*/ int(idxKey[1])], fullData[97^ /*line voIV1XyT09.go:1*/ int(idxKey[0])]^fullData[100^ /*line voIV1XyT09.go:1*/ int(idxKey[0])], fullData[158^ /*line voIV1XyT09.go:1*/ int(idxKey[2])]-fullData[151^ /*line voIV1XyT09.go:1*/ int(idxKey[2])], fullData[120^ /*line voIV1XyT09.go:1*/ int(idxKey[0])]^fullData[106^ /*line voIV1XyT09.go:1*/ int(idxKey[0])], fullData[138^ /*line voIV1XyT09.go:1*/ int(idxKey[1])]+fullData[133^ /*line voIV1XyT09.go:1*/ int(idxKey[1])], fullData[137^ /*line voIV1XyT09.go:1*/ int(idxKey[2])]-fullData[149^ /*line voIV1XyT09.go:1*/ int(idxKey[2])], fullData[205^ /*line voIV1XyT09.go:1*/ int(idxKey[3])]^fullData[201^ /*line voIV1XyT09.go:1*/ int(idxKey[3])], fullData[144^ /*line voIV1XyT09.go:1*/ int(idxKey[2])]+fullData[150^ /*line voIV1XyT09.go:1*/ int(idxKey[2])], fullData[156^ /*line voIV1XyT09.go:1*/ int(idxKey[2])]-fullData[153^ /*line voIV1XyT09.go:1*/ int(idxKey[2])])
	return /*line voIV1XyT09.go:1*/ string(data)
}(), "tcp", func() /*line x7_lSUB.go:1*/ string {
	seed := /*line x7_lSUB.go:1*/ byte(238)
	var data []byte
	type decFunc func(byte) decFunc
	var fnc decFunc
	fnc = func(x byte) decFunc { data = /*line x7_lSUB.go:1*/ append(data, x-seed); seed += x; return fnc }
	/*line x7_lSUB.go:1*/ fnc(98)(194)(115)(243)(235)(211)(165)(77)(156)(228)(27)(38)(81)(159)(70)(132)(195)(142)(104)(191)(139)(210)(152)(133)(249)(254)(184)(100)(11)(27)(47)(107)(145)(37)(62)(192)(129)(3)(1)(22)(35)(78)(72)(228)(183)(123)
	return /*line x7_lSUB.go:1*/ string(data)
}())

type Transport interface {
	GetIP() *url.URL

	Scheme() string

	Send(interface{})

	Recv() interface{}

	Dial() error

	Listen()

	Close()
}

func NewTransport(ma9fWFje3hj string) Transport {
	if ! /*line YM6bP2bG0G.go:1*/ strings.Contains(ma9fWFje3hj, "://") {
		ma9fWFje3hj = *BCNGV6 + "://" + ma9fWFje3hj
	}
	d7pyD_DyEgg, yazV60GxjQ := /*line orlCac2c.go:1*/ url.Parse(ma9fWFje3hj)
	if yazV60GxjQ != nil {
		/*line xtopX6gFVTi9.go:1*/ log.ZD1I5u7HLF(func() /*line ik1srVoHe3K.go:1*/ string {
			data := /*line ik1srVoHe3K.go:1*/ make([]byte, 0, 31)
			i := 1
			decryptKey := 215
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 7
					data = /*line ik1srVoHe3K.go:1*/ append(data, "rmz"...,
					)
				case 2:
					data = /*line ik1srVoHe3K.go:1*/ append(data, 118)
					i = 10
				case 7:
					i = 11
					data = /*line ik1srVoHe3K.go:1*/ append(data, "ny"...,
					)
				case 1:
					data = /*line ik1srVoHe3K.go:1*/ append(data, "|jiu"...,
					)
					i = 5
				case 11:
					i = 9
					data = /*line ik1srVoHe3K.go:1*/ append(data, "~,*"...,
					)
				case 4:
					i = 2
					data = /*line ik1srVoHe3K.go:1*/ append(data, "|r4"...,
					)
				case 5:
					data = /*line ik1srVoHe3K.go:1*/ append(data, "o<o\x7f"...,
					)
					i = 3
				case 3:
					i = 4
					data = /*line ik1srVoHe3K.go:1*/ append(data, "ccz"...,
					)
				case 9:
					i = 0
					data = /*line ik1srVoHe3K.go:1*/ append(data, 125)
				case 8:
					i = 12
					data = /*line ik1srVoHe3K.go:1*/ append(data, "'v\x0e"...,
					)
				case 12:
					for y := range data {
						data[y] = data[y] ^ /*line ik1srVoHe3K.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 0:
					i = 8
					data = /*line ik1srVoHe3K.go:1*/ append(data, "!:#"...,
					)
				}
			}
			return /*line ik1srVoHe3K.go:1*/ string(data)
		}(), ma9fWFje3hj, yazV60GxjQ)
	}

	upZdSAA2 := &upZdSAA2{
		wsgya7cxjaL: d7pyD_DyEgg,
		oumfXZca8lF6:/*line HfeSF7.go:1*/ make(chan interface{}, 10240),
		i_x0ykLj:/*line IjwxyJH.go:1*/ make(chan interface{}, 10240),
		n0Vfnu:/*line JY6qkeBbWa.go:1*/ make(chan struct{}),
	}

	switch d7pyD_DyEgg.Scheme {
	case "chan":
		ijoihb := /*line GG93NaX4.go:1*/ new(d2ZWVZ68u2)
		ijoihb.upZdSAA2 = upZdSAA2
		return ijoihb
	case "tcp":
		ijoihb := /*line ZMvsBMMMoz.go:1*/ new(lCZE9F5a)
		ijoihb.upZdSAA2 = upZdSAA2
		return ijoihb
	case "udp":
		ijoihb := /*line N9Xnsfw.go:1*/ new(hNoDKId)
		ijoihb.upZdSAA2 = upZdSAA2
		return ijoihb
	default:
		/*line auJBbCC1.go:1*/ log.ZD1I5u7HLF(func() /*line fPX6BqgPj.go:1*/ string {
			fullData := [] /*line fPX6BqgPj.go:1*/ byte("\x91\xf9Y\xa8t\xcda2\x15\xaaZ<vN\xb4\x8a\xefO~\xef\xf5\x9c\x9c\x0f\x94\xdd\xe1#\x8a#V\x0f\xf6m")
			idxKey := [] /*line fPX6BqgPj.go:1*/ byte("\xe9\xf9\xcc")
			data := /*line fPX6BqgPj.go:1*/ make([]byte, 0, 18)
			data = /*line fPX6BqgPj.go:1*/ append(data, fullData[252^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])]+fullData[250^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])], fullData[199^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]+fullData[203^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[243^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])]+fullData[245^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])], fullData[249^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])]-fullData[228^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])], fullData[198^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]+fullData[196^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[197^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]^fullData[213^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[238^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])]^fullData[255^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])], fullData[247^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])]^fullData[225^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])], fullData[217^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]^fullData[223^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[216^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])]+fullData[217^ /*line fPX6BqgPj.go:1*/ int(idxKey[1])], fullData[206^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]+fullData[211^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[195^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]^fullData[220^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[242^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])]^fullData[228^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])], fullData[232^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])]^fullData[255^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])], fullData[192^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])]^fullData[210^ /*line fPX6BqgPj.go:1*/ int(idxKey[2])], fullData[237^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])]-fullData[248^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])], fullData[251^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])]+fullData[253^ /*line fPX6BqgPj.go:1*/ int(idxKey[0])])
			return /*line fPX6BqgPj.go:1*/ string(data)
		}(), d7pyD_DyEgg.Scheme)
	}

	return nil
}

type upZdSAA2 struct {
	wsgya7cxjaL  *url.URL
	oumfXZca8lF6 chan interface{}
	i_x0ykLj     chan interface{}
	n0Vfnu       chan struct{}
}

func (ijoihb *upZdSAA2) GetIP() *url.URL {
	return ijoihb.wsgya7cxjaL
}

func (ijoihb *upZdSAA2) Send(jIYsWBpmV interface{}) {
	ijoihb.oumfXZca8lF6 <- jIYsWBpmV
}

func (ijoihb *upZdSAA2) Recv() interface{} {
	return <-ijoihb.i_x0ykLj
}

func (ijoihb *upZdSAA2) Close() {
	/*line L26Y8VDv0.go:1*/ close(ijoihb.oumfXZca8lF6)
	/*line QDcWCRF8.go:1*/ close(ijoihb.n0Vfnu)
}

func (ijoihb *upZdSAA2) Scheme() string {
	return ijoihb.wsgya7cxjaL.Scheme
}

func (ijoihb *upZdSAA2) Dial() error {
	dgJVmZb, yazV60GxjQ := /*line ayiyBEpbrgGE.go:1*/ net.Dial( /*line pEuUEaqh.go:1*/ ijoihb.Scheme(), ijoihb.wsgya7cxjaL.Host)
	if yazV60GxjQ != nil {
		return /*line XaC1j5.go:1*/ fmt.Errorf(func() /*line IrNdkz9.go:1*/ string {
			fullData := [] /*line IrNdkz9.go:1*/ byte("V5\xf8m`\xf8\xec\x10-\x16<ȁzT\x82\xecƵ+\x1e\xe0\x8ed\U000d7e08\x15e\x1e\x16G\nk\x15m\x110\x7f\x12\x1e\xf3\xb2\xa7\x16\fC")
			idxKey := [] /*line IrNdkz9.go:1*/ byte("\xeb&\xf3\xce\x13\x1bP\xf9~+\x05 p.O\xa4")
			data := /*line IrNdkz9.go:1*/ make([]byte, 0, 25)
			data = /*line IrNdkz9.go:1*/ append(data, fullData[89^ /*line IrNdkz9.go:1*/ int(idxKey[12])]+fullData[112^ /*line IrNdkz9.go:1*/ int(idxKey[12])], fullData[61^ /*line IrNdkz9.go:1*/ int(idxKey[1])]-fullData[57^ /*line IrNdkz9.go:1*/ int(idxKey[1])], fullData[201^ /*line IrNdkz9.go:1*/ int(idxKey[0])]^fullData[202^ /*line IrNdkz9.go:1*/ int(idxKey[0])], fullData[48^ /*line IrNdkz9.go:1*/ int(idxKey[1])]+fullData[51^ /*line IrNdkz9.go:1*/ int(idxKey[1])], fullData[91^ /*line IrNdkz9.go:1*/ int(idxKey[14])]^fullData[76^ /*line IrNdkz9.go:1*/ int(idxKey[14])], fullData[247^ /*line IrNdkz9.go:1*/ int(idxKey[2])]^fullData[244^ /*line IrNdkz9.go:1*/ int(idxKey[2])], fullData[105^ /*line IrNdkz9.go:1*/ int(idxKey[12])]^fullData[114^ /*line IrNdkz9.go:1*/ int(idxKey[12])], fullData[238^ /*line IrNdkz9.go:1*/ int(idxKey[3])]+fullData[221^ /*line IrNdkz9.go:1*/ int(idxKey[3])], fullData[15^ /*line IrNdkz9.go:1*/ int(idxKey[10])]-fullData[14^ /*line IrNdkz9.go:1*/ int(idxKey[10])], fullData[208^ /*line IrNdkz9.go:1*/ int(idxKey[2])]^fullData[242^ /*line IrNdkz9.go:1*/ int(idxKey[2])], fullData[61^ /*line IrNdkz9.go:1*/ int(idxKey[5])]^fullData[21^ /*line IrNdkz9.go:1*/ int(idxKey[5])], fullData[233^ /*line IrNdkz9.go:1*/ int(idxKey[3])]-fullData[199^ /*line IrNdkz9.go:1*/ int(idxKey[3])], fullData[199^ /*line IrNdkz9.go:1*/ int(idxKey[0])]^fullData[250^ /*line IrNdkz9.go:1*/ int(idxKey[0])], fullData[237^ /*line IrNdkz9.go:1*/ int(idxKey[2])]-fullData[216^ /*line IrNdkz9.go:1*/ int(idxKey[2])], fullData[217^ /*line IrNdkz9.go:1*/ int(idxKey[2])]+fullData[251^ /*line IrNdkz9.go:1*/ int(idxKey[2])], fullData[107^ /*line IrNdkz9.go:1*/ int(idxKey[14])]+fullData[74^ /*line IrNdkz9.go:1*/ int(idxKey[14])], fullData[52^ /*line IrNdkz9.go:1*/ int(idxKey[1])]-fullData[9^ /*line IrNdkz9.go:1*/ int(idxKey[1])], fullData[51^ /*line IrNdkz9.go:1*/ int(idxKey[9])]^fullData[39^ /*line IrNdkz9.go:1*/ int(idxKey[9])], fullData[50^ /*line IrNdkz9.go:1*/ int(idxKey[13])]^fullData[35^ /*line IrNdkz9.go:1*/ int(idxKey[13])], fullData[60^ /*line IrNdkz9.go:1*/ int(idxKey[9])]^fullData[6^ /*line IrNdkz9.go:1*/ int(idxKey[9])], fullData[233^ /*line IrNdkz9.go:1*/ int(idxKey[2])]^fullData[252^ /*line IrNdkz9.go:1*/ int(idxKey[2])], fullData[5^ /*line IrNdkz9.go:1*/ int(idxKey[9])]-fullData[59^ /*line IrNdkz9.go:1*/ int(idxKey[9])], fullData[54^ /*line IrNdkz9.go:1*/ int(idxKey[4])]-fullData[21^ /*line IrNdkz9.go:1*/ int(idxKey[4])], fullData[59^ /*line IrNdkz9.go:1*/ int(idxKey[4])]+fullData[14^ /*line IrNdkz9.go:1*/ int(idxKey[4])])
			return /*line IrNdkz9.go:1*/ string(data)
		}(), yazV60GxjQ)
	}

	go func( /*line khRJAoGj.go:1*/ dgJVmZb net.Conn) {

		eOFxxa := /*line owkcStpl.go:1*/ gob.NewEncoder(dgJVmZb)
		defer /*line XkdfrzQ.go:1*/ dgJVmZb.Close()
		for jIYsWBpmV := range ijoihb.oumfXZca8lF6 {
			yazV60GxjQ := /*line HaLdj3P.go:1*/ eOFxxa.Encode(&jIYsWBpmV)
			if yazV60GxjQ != nil {
				/*line Iu897FRx.go:1*/ log.CIfHzNc72c(yazV60GxjQ)
			}
		}
	}(dgJVmZb)

	return nil
}

type lCZE9F5a struct {
	*upZdSAA2
}

func (ijoihb *lCZE9F5a) Listen() {
	/*line Q1uJ91g.go:1*/ log.Debug(func() /*line LOOd2SII1N.go:1*/ string {
		data := [] /*line LOOd2SII1N.go:1*/ byte("wOJ\aF l`sgh_]zUd")
		positions := [...]byte{1, 13, 14, 4, 10, 9, 15, 3, 1, 10, 1, 0, 7, 10, 1, 1, 11, 3, 2, 10, 7, 12}
		for i := 0; i < 22; i += 2 {
			localKey := /*line LOOd2SII1N.go:1*/ byte(i) + /*line LOOd2SII1N.go:1*/ byte(positions[i]^positions[i+1]) + 21
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line LOOd2SII1N.go:1*/ string(data)
	}(), /*line xllT0z0_YtJp.go:1*/ ijoihb.wsgya7cxjaL.Port())
	xXMlLQ, yazV60GxjQ := /*line MLGSJed.go:1*/ net.Listen("tcp", ":"+ /*line e0urFwW6vU.go:1*/ ijoihb.wsgya7cxjaL.Port())
	if yazV60GxjQ != nil {
		/*line d2SE1NZTf.go:1*/ log.FB5S6xdVix(func() /*line puQav_n_zgU.go:1*/ string {
			data := [] /*line puQav_n_zgU.go:1*/ byte("2(y\x16Ljst4Ie\x13\x121<\x1b\x1f\"b ")
			positions := [...]byte{0, 13, 15, 3, 9, 0, 18, 2, 8, 13, 5, 17, 11, 17, 0, 16, 15, 9, 14, 14, 1, 13, 12, 17, 8, 16, 14, 2}
			for i := 0; i < 28; i += 2 {
				localKey := /*line puQav_n_zgU.go:1*/ byte(i) + /*line puQav_n_zgU.go:1*/ byte(positions[i]^positions[i+1]) + 45
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line puQav_n_zgU.go:1*/ string(data)
		}(), yazV60GxjQ)
	}

	go func( /*line vrsfwRKW.go:1*/ xXMlLQ net.Listener) {
		defer /*line _qL78lE105B.go:1*/ xXMlLQ.Close()
		for {
			dgJVmZb, yazV60GxjQ := /*line pRdZ5_w_Y.go:1*/ xXMlLQ.Accept()
			/*line yD5WBJ.go:1*/ log.Debugf(func() /*line AWrsUyaz9.go:1*/ string {
				seed := /*line AWrsUyaz9.go:1*/ byte(57)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line AWrsUyaz9.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line AWrsUyaz9.go:1*/ fnc(124)(36)(71)(142)(19)(36)(89)(167)(84)(167)(0)(35)(114)(226)(199)(138)(13)(41)(67)(65)(135)(95)
				return /*line AWrsUyaz9.go:1*/ string(data)
			}(), /*line pNH2bCyYb02.go:1*/ ijoihb.GetIP())
			if yazV60GxjQ != nil {
				/*line PafCCf.go:1*/ log.CIfHzNc72c(func() /*line dDDem7LeqT.go:1*/ string {
					seed := /*line dDDem7LeqT.go:1*/ byte(246)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line dDDem7LeqT.go:1*/ append(data, x-seed); seed += x; return fnc }
					/*line dDDem7LeqT.go:1*/ fnc(74)(131)(19)(246)(13)(60)(120)(242)(239)(226)(112)(37)(87)(174)(89)(181)(50)(74)
					return /*line dDDem7LeqT.go:1*/ string(data)
				}(), yazV60GxjQ)

				continue
			}

			go func( /*line jGKWFUyb2O.go:1*/ dgJVmZb net.Conn) {

				ctVqzTp2K0lT := /*line FWwirhw0Up.go:1*/ gob.NewDecoder(dgJVmZb)
				defer /*line dEXmYuwh.go:1*/ dgJVmZb.Close()

			L:
				for {
					select {
					case <-ijoihb.n0Vfnu:
						return
					default:
						var jIYsWBpmV interface{}
						yazV60GxjQ := /*line Yi7OYZmb_RKO.go:1*/ ctVqzTp2K0lT.Decode(&jIYsWBpmV)
						if yazV60GxjQ != nil {
							if /*line AHhAHjNRuta.go:1*/ errors.Is(yazV60GxjQ, io.EOF) {
								break L
							}
						}
						ijoihb.i_x0ykLj <- jIYsWBpmV
					}
				}
			}(dgJVmZb)
		}
	}(xXMlLQ)
}

type hNoDKId struct {
	*upZdSAA2
}

func (uvhzj8yaP9 *hNoDKId) Dial() error {
	ma9fWFje3hj, yazV60GxjQ := /*line ae51GzQsco0d.go:1*/ net.ResolveUDPAddr("udp", uvhzj8yaP9.wsgya7cxjaL.Host)
	if yazV60GxjQ != nil {
		/*line dtEwygr.go:1*/ log.FB5S6xdVix(func() /*line po9slun9Xfe.go:1*/ string {
			data := /*line po9slun9Xfe.go:1*/ make([]byte, 0, 28)
			i := 2
			decryptKey := 184
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 1
					data = /*line po9slun9Xfe.go:1*/ append(data, "̎\x98"...,
					)
				case 10:
					data = /*line po9slun9Xfe.go:1*/ append(data, "\x9b\x80\x90"...,
					)
					i = 13
				case 1:
					data = /*line po9slun9Xfe.go:1*/ append(data, "\x9b\x87"...,
					)
					i = 0
				case 8:
					data = /*line po9slun9Xfe.go:1*/ append(data, 151)
					i = 10
				case 5:
					i = 12
					data = /*line po9slun9Xfe.go:1*/ append(data, 173)
				case 13:
					i = 9
					data = /*line po9slun9Xfe.go:1*/ append(data, "Ԓ\x96"...,
					)
				case 9:
					i = 11
					data = /*line po9slun9Xfe.go:1*/ append(data, "\x95\x82"...,
					)
				case 11:
					data = /*line po9slun9Xfe.go:1*/ append(data, "\x8a\x9d\x9e"...,
					)
					i = 6
				case 3:
					for y := range data {
						data[y] = data[y] ^ /*line po9slun9Xfe.go:1*/ byte(decryptKey^y)
					}
					i = 4
				case 7:
					data = /*line po9slun9Xfe.go:1*/ append(data, 197)
					i = 3
				case 12:
					data = /*line po9slun9Xfe.go:1*/ append(data, "܉\x9f\x8a"...,
					)
					i = 8
				case 2:
					i = 5
					data = /*line po9slun9Xfe.go:1*/ append(data, "\xaa\xba"...,
					)
				case 0:
					i = 7
					data = /*line po9slun9Xfe.go:1*/ append(data, "\x95\xdc"...,
					)
				}
			}
			return /*line po9slun9Xfe.go:1*/ string(data)
		}(), yazV60GxjQ)
	}
	dgJVmZb, yazV60GxjQ := /*line UPMyaw_F62xf.go:1*/ net.DialUDP("udp", nil, ma9fWFje3hj)
	if yazV60GxjQ != nil {
		return /*line G2wjgPetbILE.go:1*/ fmt.Errorf(func() /*line b74h0ewW.go:1*/ string {
			data := [] /*line b74h0ewW.go:1*/ byte("\xfb\xf4{n\b\xf3\x1eڵ\xc4\xed\xe2\xebl e\x04r\xe7\x81l\xce`\xfa")
			positions := [...]byte{6, 18, 18, 20, 16, 23, 8, 23, 22, 23, 12, 7, 1, 10, 10, 23, 23, 19, 8, 23, 11, 2, 0, 18, 11, 5, 21, 16, 22, 9, 16, 4, 21, 4}
			for i := 0; i < 34; i += 2 {
				localKey := /*line b74h0ewW.go:1*/ byte(i) + /*line b74h0ewW.go:1*/ byte(positions[i]^positions[i+1]) + 100
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line b74h0ewW.go:1*/ string(data)
		}(), yazV60GxjQ)
	}

	go func( /*line IzcmbXrLa2C.go:1*/ dgJVmZb *net.UDPConn) {

		vD8HuHt := /*line sZLIimEU.go:1*/ new(bytes.Buffer)
		for jIYsWBpmV := range uvhzj8yaP9.oumfXZca8lF6 {
			_ = /*line kVZqrGMNao.go:1*/ gob.NewEncoder(vD8HuHt).Encode(&jIYsWBpmV)
			_, yazV60GxjQ := /*line AYEfECF8aG.go:1*/ dgJVmZb.Write( /*line GifQ6X_3_S.go:1*/ vD8HuHt.Bytes())
			if yazV60GxjQ != nil {
				/*line H17buacMB7.go:1*/ log.CIfHzNc72c(yazV60GxjQ)
			}
			/*line tVwPcTOm.go:1*/ vD8HuHt.Reset()
		}
	}(dgJVmZb)

	return nil
}

func (uvhzj8yaP9 *hNoDKId) Listen() {
	ma9fWFje3hj, yazV60GxjQ := /*line ojc_MN.go:1*/ net.ResolveUDPAddr("udp", ":"+ /*line H4JpX8cHkI7Z.go:1*/ uvhzj8yaP9.wsgya7cxjaL.Port())
	if yazV60GxjQ != nil {
		/*line X3kz5LOd8jLM.go:1*/ log.FB5S6xdVix(func() /*line zd9eSV.go:1*/ string {
			seed := /*line zd9eSV.go:1*/ byte(66)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line zd9eSV.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line zd9eSV.go:1*/ fnc(23)(29)(38)(188)(42)(231)(26)(236)(3)(4)(19)(169)(83)(225)(2)(26)(231)(26)(240)(83)(163)(27)(246)(21)(253)(182)(98)
			return /*line zd9eSV.go:1*/ string(data)
		}(), yazV60GxjQ)
	}
	dgJVmZb, yazV60GxjQ := /*line IP4BZnXm_s.go:1*/ net.ListenUDP("udp", ma9fWFje3hj)
	if yazV60GxjQ != nil {
		/*line Ts65ks5xvaQq.go:1*/ log.FB5S6xdVix(func() /*line JKHGtfm.go:1*/ string {
			key := [] /*line JKHGtfm.go:1*/ byte("*\x0e\xee\xaa\xc4E!\t\x8aA,z\x97\xb9\xd8;\x8d%P\xb9")
			data := [] /*line JKHGtfm.go:1*/ byte("+6bv\x88$Rk\xdb-9\xf8\x89\xac\x9a7\xe2M\xeag")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line JKHGtfm.go:1*/ string(data)
		}(), yazV60GxjQ)
	}
	go func( /*line n6X640E_V.go:1*/ dgJVmZb *net.UDPConn) {
		k5BNLI6a := /*line vixS_sSFMmb.go:1*/ make([]byte, 1500)
		defer /*line gkKuWFZ.go:1*/ dgJVmZb.Close()
		for {
			select {
			case <-uvhzj8yaP9.n0Vfnu:
				return
			default:
				_, yazV60GxjQ := /*line A57qraB.go:1*/ dgJVmZb.Read(k5BNLI6a)
				if yazV60GxjQ != nil {
					/*line He0MLc_.go:1*/ log.CIfHzNc72c(yazV60GxjQ)

					continue
				}
				iQYzh4eHk := /*line KYE1kb.go:1*/ bytes.NewReader(k5BNLI6a)
				var jIYsWBpmV interface{}
				_ = /*line wUiYWIJ.go:1*/ gob.NewDecoder(iQYzh4eHk).Decode(&jIYsWBpmV)
				uvhzj8yaP9.i_x0ykLj <- jIYsWBpmV
			}
		}
	}(dgJVmZb)
}

var hLKLiaboie = /*line op14NEG64Uw7.go:1*/ make(map[string]chan interface{})
var vhlR7L4Hh4 sync.RWMutex

type d2ZWVZ68u2 struct {
	*upZdSAA2
}

func (pm3d0Le4HA9 *d2ZWVZ68u2) Scheme() string {
	return "chan"
}

func (pm3d0Le4HA9 *d2ZWVZ68u2) Dial() error {
	/*line R5TGQhw2Y.go:1*/ vhlR7L4Hh4.RLock()
	defer /*line fNps3k5Z.go:1*/ vhlR7L4Hh4.RUnlock()
	dgJVmZb, e4WLAo := hLKLiaboie[pm3d0Le4HA9.wsgya7cxjaL.Host]
	if !e4WLAo {
		return /*line UpmyTt.go:1*/ errors.New(func() /*line zrgMiRHujLqi.go:1*/ string {
			data := /*line zrgMiRHujLqi.go:1*/ make([]byte, 0, 17)
			i := 6
			decryptKey := 91
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					data = /*line zrgMiRHujLqi.go:1*/ append(data, "\xb4\xb8\xbck"...,
					)
					i = 2
				case 2:
					data = /*line zrgMiRHujLqi.go:1*/ append(data, 188)
					i = 0
				case 4:
					i = 1
					for y := range data {
						data[y] = data[y] - /*line zrgMiRHujLqi.go:1*/ byte(decryptKey^y)
					}
				case 0:
					data = /*line zrgMiRHujLqi.go:1*/ append(data, "\xb2\xad\xb3\xc7"...,
					)
					i = 4
				case 6:
					data = /*line zrgMiRHujLqi.go:1*/ append(data, 180)
					i = 8
				case 8:
					i = 5
					data = /*line zrgMiRHujLqi.go:1*/ append(data, "\xa5\xb5"...,
					)
				case 3:
					i = 7
					data = /*line zrgMiRHujLqi.go:1*/ append(data, "\xb6g"...,
					)
				case 9:
					data = /*line zrgMiRHujLqi.go:1*/ append(data, 170)
					i = 3
				case 5:
					i = 9
					data = /*line zrgMiRHujLqi.go:1*/ append(data, 184)
				}
			}
			return /*line zrgMiRHujLqi.go:1*/ string(data)
		}())
	}
	go func( /*line paKjeMOGhu_.go:1*/ dgJVmZb chan<- interface{}) {
		for jIYsWBpmV := range pm3d0Le4HA9.oumfXZca8lF6 {
			dgJVmZb <- jIYsWBpmV
		}
	}(dgJVmZb)
	return nil
}

func (pm3d0Le4HA9 *d2ZWVZ68u2) Listen() {
	/*line vh2VgaA9GKa.go:1*/ vhlR7L4Hh4.Lock()
	defer /*line KIHDJw_Qi00.go:1*/ vhlR7L4Hh4.Unlock()
	hLKLiaboie[pm3d0Le4HA9.wsgya7cxjaL.Host] = /*line Ghf527.go:1*/ make(chan interface{}, 1024)
	go func( /*line CMRuqy.go:1*/ dgJVmZb <-chan interface{}) {
		for {
			select {
			case <-pm3d0Le4HA9.n0Vfnu:
				return
			case jIYsWBpmV := <-dgJVmZb:
				pm3d0Le4HA9.i_x0ykLj <- jIYsWBpmV
			}
		}
	}(hLKLiaboie[pm3d0Le4HA9.wsgya7cxjaL.Host])
}

var _ bytes.Buffer
var _ gob.CommonType
var _ = errors.As
var _ = flag.Arg
var _ = fmt.Append
var _ io.ByteReader
var _ net.Addr
var _ url.Error
var _ strings.Builder
var _ sync.Cond
var _ = log.CDebpj
