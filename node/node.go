//line :1
package node

import (
	"net/http"
	"reflect"
	"sync"

	config "unishard/config"
	log "unishard/log"
	message "unishard/message"
	socket "unishard/socket"
	types "unishard/types"
)

type Node interface {
	socket.EbznnrSS9z

	ID() types.NodeID
	Shard() types.Shard
	State() types.T30k6y_tK
	Role() types.FP9GjX
	SetState(z_hgAx types.T30k6y_tK)
	SetRole(wZNBvH types.FP9GjX)
	Run()
	Retry(yy2jgIELlF message.Transaction)
	Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{})
	IsByz() bool
}

type node struct {
	id    types.NodeID
	shard types.Shard
	state types.T30k6y_tK
	role  types.FP9GjX

	socket.EbznnrSS9z

	MessageChan           chan interface{}
	CrossChainMessageChan chan interface{}
	TxChan                chan interface{}
	handles               map[string]reflect.Value
	server                *http.Server
	isByz                 bool
	totalTxn              int

	sync.RWMutex
	forwards map[string]*message.Transaction
}

func NewNode(wCjuxLIbrSy types.NodeID, gDODt11q3pkZ bool, js85nByJG types.Shard) Node {
	return &node{
		id:    wCjuxLIbrSy,
		shard: js85nByJG,
		state: types.READY,
		role:  types.VALIDATOR,
		isByz: gDODt11q3pkZ,
		EbznnrSS9z:/*line Eoa_Ovcte.go:1*/ socket.UxTMLURlR(wCjuxLIbrSy, config.Configuration.Addrs, js85nByJG),

		MessageChan:/*line h9b1CISHV.go:1*/ make(chan interface{}, config.Configuration.ChanBufferSize),
		CrossChainMessageChan:/*line VafKVKqwiDC7.go:1*/ make(chan interface{}, config.Configuration.ChanBufferSize),
		TxChan:/*line gtkt6jqadpue.go:1*/ make(chan interface{}, config.Configuration.ChanBufferSize),
		handles:/*line hmIVQK.go:1*/ make(map[string]reflect.Value),
		forwards:/*line xqqwd_dd.go:1*/ make(map[string]*message.Transaction),
	}
}

func (eL4EyO3pzuT *node) ID() types.NodeID {
	return eL4EyO3pzuT.id
}

func (eL4EyO3pzuT *node) State() types.T30k6y_tK {
	/*line ysmh6h.go:1*/ eL4EyO3pzuT.RWMutex.Lock()
	defer /*line CkjmMGZv1wqX.go:1*/ eL4EyO3pzuT.RWMutex.Unlock()
	return eL4EyO3pzuT.state
}

func (eL4EyO3pzuT *node) Role() types.FP9GjX {
	/*line GOJ4dX.go:1*/ eL4EyO3pzuT.RWMutex.Lock()
	defer /*line m7nnS7.go:1*/ eL4EyO3pzuT.RWMutex.Unlock()
	return eL4EyO3pzuT.role
}

func (eL4EyO3pzuT *node) Shard() types.Shard {
	return eL4EyO3pzuT.shard
}

func (eL4EyO3pzuT *node) IsByz() bool {
	return eL4EyO3pzuT.isByz
}

func (eL4EyO3pzuT *node) SetState(z_hgAx types.T30k6y_tK) {
	/*line yaflZMK.go:1*/ eL4EyO3pzuT.RWMutex.Lock()
	defer /*line oao59xz3RXa.go:1*/ eL4EyO3pzuT.RWMutex.Unlock()
	eL4EyO3pzuT.state = z_hgAx
}

func (eL4EyO3pzuT *node) SetRole(wZNBvH types.FP9GjX) {
	/*line dfSoTgaI.go:1*/ eL4EyO3pzuT.RWMutex.Lock()
	defer /*line iEpuBjbfGN.go:1*/ eL4EyO3pzuT.RWMutex.Unlock()
	eL4EyO3pzuT.role = wZNBvH
}

func (eL4EyO3pzuT *node) Retry(yy2jgIELlF message.Transaction) {
	/*line Fm0UwPy82j.go:1*/ log.Debugf(func() /*line zya2MzvaRs_q.go:1*/ string {
		fullData := [] /*line zya2MzvaRs_q.go:1*/ byte("\x8e\xb4\x9c\xa0\xb9\xd1(\xbcJn\x01\xbd\xf7ӲΔd\xd0ɀ\xacOҦ\x13\xe5\xea\xf1\x19Z.\x88\xb7\xdb\xe5\xcbe\x12\b\xe3zm\xb9\x94\xc3_\xbd")
		idxKey := [] /*line zya2MzvaRs_q.go:1*/ byte("fi9\b\xbd\x06\x00K\xb5U\xe0\xce\\")
		data := /*line zya2MzvaRs_q.go:1*/ make([]byte, 0, 25)
		data = /*line zya2MzvaRs_q.go:1*/ append(data, fullData[144^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[8])]-fullData[185^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[8])], fullData[73^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[12])]+fullData[113^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[12])], fullData[180^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[8])]^fullData[167^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[8])], fullData[196^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[11])]-fullData[204^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[11])], fullData[33^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[5])]^fullData[0^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[5])], fullData[29^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[2])]+fullData[39^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[2])], fullData[164^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[8])]+fullData[147^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[8])], fullData[82^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[12])]+fullData[85^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[12])], fullData[110^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[1])]^fullData[102^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[1])], fullData[65^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[9])]^fullData[118^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[9])], fullData[29^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[6])]^fullData[42^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[6])], fullData[125^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[0])]+fullData[70^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[0])], fullData[69^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[1])]+fullData[115^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[1])], fullData[203^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[11])]+fullData[216^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[11])], fullData[38^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[3])]+fullData[17^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[3])], fullData[87^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[7])]^fullData[91^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[7])], fullData[6^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[5])]+fullData[46^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[5])], fullData[17^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[5])]^fullData[39^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[5])], fullData[255^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[10])]-fullData[228^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[10])], fullData[205^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[11])]+fullData[195^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[11])], fullData[98^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[1])]^fullData[122^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[1])], fullData[68^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[12])]+fullData[117^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[12])], fullData[159^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[4])]+fullData[181^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[4])], fullData[22^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[2])]+fullData[18^ /*line zya2MzvaRs_q.go:1*/ int(idxKey[2])])
		return /*line zya2MzvaRs_q.go:1*/ string(data)
	}(), eL4EyO3pzuT.id, yy2jgIELlF)
	eL4EyO3pzuT.MessageChan <- yy2jgIELlF
}

func (eL4EyO3pzuT *node) Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{}) {
	vgqeZiHun9F := /*line kuLUunSkrmbA.go:1*/ reflect.TypeOf(cYu43eeVpCaK)
	cLL2dKRL8 := /*line X7yXNBEFd.go:1*/ reflect.ValueOf(cy0xXq2tR)

	if /*line natQfpf0YRSF.go:1*/ cLL2dKRL8.Kind() != reflect.Func {
		/*line R35bjT6bngX2.go:1*/ panic(func() /*line W13U_gDZYWcp.go:1*/ string {
			fullData := [] /*line W13U_gDZYWcp.go:1*/ byte("^\xc0\xffg\"\xa2\x81\x9c\x14\f\x9b\x94\x89},\x9cU~\xc3\x1b\x02\xfb\xc1\xe0ڄP潩p`^>\xcd\xd9b#\xc7_\xfevU\xa1\xf2s\x1aTU~\xfb\xf5\xffz")
			idxKey := [] /*line W13U_gDZYWcp.go:1*/ byte("\xaf\x99VY\x06\x03\x06\xc6")
			data := /*line W13U_gDZYWcp.go:1*/ make([]byte, 0, 28)
			data = /*line W13U_gDZYWcp.go:1*/ append(data, fullData[164^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[0])]-fullData[161^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[0])], fullData[79^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])]-fullData[115^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])], fullData[123^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])]+fullData[100^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])], fullData[55^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])]+fullData[29^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])], fullData[40^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[5])]^fullData[33^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[5])], fullData[194^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[7])]-fullData[218^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[7])], fullData[25^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])]+fullData[7^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])], fullData[154^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])]+fullData[155^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])], fullData[94^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])]+fullData[122^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])], fullData[226^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[7])]+fullData[207^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[7])], fullData[98^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])]^fullData[89^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])], fullData[41^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])]-fullData[17^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])], fullData[115^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])]+fullData[81^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])], fullData[101^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])]+fullData[99^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])], fullData[131^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])]^fullData[184^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])], fullData[79^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])]+fullData[126^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])], fullData[112^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])]-fullData[118^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])], fullData[117^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])]+fullData[95^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])], fullData[15^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[5])]^fullData[30^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[5])], fullData[72^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])]+fullData[126^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[2])], fullData[22^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])]+fullData[40^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[4])], fullData[141^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])]^fullData[176^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])], fullData[18^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[5])]+fullData[6^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[5])], fullData[138^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])]^fullData[148^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])], fullData[65^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])]+fullData[83^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[3])], fullData[139^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])]-fullData[169^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[1])], fullData[6^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[6])]-fullData[19^ /*line W13U_gDZYWcp.go:1*/ int(idxKey[6])])
			return /*line W13U_gDZYWcp.go:1*/ string(data)
		}())
	}

	if /*line MD87ZHOi.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line CDgPYYyTrP.go:1*/ panic(func() /*line WPNFY0WPJ5V.go:1*/ string {
			data := /*line WPNFY0WPJ5V.go:1*/ make([]byte, 0, 19)
			i := 2
			decryptKey := 128
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, "\xbd\xcdù"...,
					)
					i = 4
				case 0:
					i = 1
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, "\xc8\xca\xc2\xc4"...,
					)
				case 9:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, 200)
					i = 6
				case 3:
					for y := range data {
						data[y] = data[y] + /*line WPNFY0WPJ5V.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 6:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, 206)
					i = 5
				case 7:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, "\xc6\xd1{\xca"...,
					)
					i = 9
				case 5:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, "g\xbc"...,
					)
					i = 3
				case 4:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, 115)
					i = 0
				case 1:
					data = /*line WPNFY0WPJ5V.go:1*/ append(data, 128)
					i = 7
				}
			}
			return /*line WPNFY0WPJ5V.go:1*/ string(data)
		}())
	}

	if /*line r7UN0OgIzXhG.go:1*/ cLL2dKRL8.Kind() != reflect.Func || /*line QNzU0Gy2a.go:1*/ cLL2dKRL8.Type().NumIn() != 1 || /*line JwNSFxj.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line mZ4GtrBs.go:1*/ panic(func() /*line vawvdi.go:1*/ string {
			seed := /*line vawvdi.go:1*/ byte(86)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line vawvdi.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line vawvdi.go:1*/ fnc(36)(31)(254)(254)(230)(15)(239)(11)(164)(64)(9)(31)(244)(232)(9)(85)(172)(3)(23)(243)(247)(19)(226)(1)(80)(165)(23)(14)(229)(29)
			return /*line vawvdi.go:1*/ string(data)
		}())
	}
	eL4EyO3pzuT.handles[ /*line RyaZfszY.go:1*/ vgqeZiHun9F.String()] = cLL2dKRL8
}

func (eL4EyO3pzuT *node) Run() {
	/*line CMUetZT.go:1*/ log.Sa4H8el(func() /*line gLLv755.go:1*/ string {
		key := [] /*line gLLv755.go:1*/ byte("8\xa3\x19\xf8H\x8b{\xbe\f3$P긩\xd0\t\xf1V\x8bu")
		data := [] /*line gLLv755.go:1*/ byte("\xa6\x12}]h\xb0\xf1\xde\x7f\xa7\x85\xc2^\xd8\x1bEw_\xbf\xf9\xdc")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line gLLv755.go:1*/ string(data)
	}(), eL4EyO3pzuT.id)
	if /*line BWy8_In.go:1*/ len(eL4EyO3pzuT.handles) > 0 {
		go /*line SIogkWeMNQ.go:1*/ eL4EyO3pzuT.hX8crg5dh6()
		go /*line Ix4gOTFYzDCu.go:1*/ eL4EyO3pzuT.o4PwoT23()
		go /*line aVb1Rwfu.go:1*/ eL4EyO3pzuT.f8BxkZMSKh()
		go /*line pFUZ0IaGDJ1.go:1*/ eL4EyO3pzuT.deKdeLzuSxIn()
	}

	r8e8mXbo := message.ConsensusNodeRegister{
		ConsensusNodeID: eL4EyO3pzuT.id,
		IP:/*line EHSs3ek.go:1*/ eL4EyO3pzuT.EbznnrSS9z.GetAddresses()[eL4EyO3pzuT.shard][eL4EyO3pzuT.id],
	}
	/*line wkiWmw_.go:1*/ eL4EyO3pzuT.SendToBlockBuilder(r8e8mXbo)
	/*line HiZ0sS6.go:1*/ log.Debugf(func() /*line Pa7lRlE.go:1*/ string {
		data := /*line Pa7lRlE.go:1*/ make([]byte, 0, 64)
		i := 19
		decryptKey := 203
		for counter := 0; i != 17; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data = /*line Pa7lRlE.go:1*/ append(data, "\xa0\xab\xae\xa5"...,
				)
				i = 13
			case 0:
				for y := range data {
					data[y] = data[y] + /*line Pa7lRlE.go:1*/ byte(decryptKey^y)
				}
				i = 17
			case 19:
				data = /*line Pa7lRlE.go:1*/ append(data, "\x8f\x83"...,
				)
				i = 27
			case 27:
				i = 24
				data = /*line Pa7lRlE.go:1*/ append(data, "{\x8b\x84"...,
				)
			case 13:
				data = /*line Pa7lRlE.go:1*/ append(data, "\xb2\xb2]"...,
				)
				i = 15
			case 18:
				i = 21
				data = /*line Pa7lRlE.go:1*/ append(data, "kt|l"...,
				)
			case 10:
				data = /*line Pa7lRlE.go:1*/ append(data, 155)
				i = 4
			case 20:
				data = /*line Pa7lRlE.go:1*/ append(data, 158)
				i = 10
			case 4:
				i = 12
				data = /*line Pa7lRlE.go:1*/ append(data, "U\xa0\x9a"...,
				)
			case 21:
				i = 23
				data = /*line Pa7lRlE.go:1*/ append(data, 120)
			case 12:
				data = /*line Pa7lRlE.go:1*/ append(data, "J\x8b\x9c\x9e"...,
				)
				i = 14
			case 25:
				i = 1
				data = /*line Pa7lRlE.go:1*/ append(data, "6:\x82+"...,
				)
			case 1:
				data = /*line Pa7lRlE.go:1*/ append(data, 125)
				i = 8
			case 2:
				data = /*line Pa7lRlE.go:1*/ append(data, "\x8b\x8d\x8c\x8c"...,
				)
				i = 22
			case 23:
				data = /*line Pa7lRlE.go:1*/ append(data, 37)
				i = 11
			case 8:
				i = 26
				data = /*line Pa7lRlE.go:1*/ append(data, "n~s."...,
				)
			case 5:
				i = 20
				data = /*line Pa7lRlE.go:1*/ append(data, "\x98\xa5\xa4\x99"...,
				)
			case 24:
				i = 6
				data = /*line Pa7lRlE.go:1*/ append(data, "Y>B"...,
				)
			case 3:
				i = 16
				data = /*line Pa7lRlE.go:1*/ append(data, "z|"...,
				)
			case 15:
				data = /*line Pa7lRlE.go:1*/ append(data, 161)
				i = 5
			case 6:
				data = /*line Pa7lRlE.go:1*/ append(data, "\x8a?2"...,
				)
				i = 3
			case 14:
				i = 2
				data = /*line Pa7lRlE.go:1*/ append(data, "\x91\x98\x86\x98"...,
				)
			case 9:
				data = /*line Pa7lRlE.go:1*/ append(data, 106)
				i = 18
			case 16:
				i = 25
				data = /*line Pa7lRlE.go:1*/ append(data, 81)
			case 26:
				i = 9
				data = /*line Pa7lRlE.go:1*/ append(data, "\x7fi"...,
				)
			case 11:
				i = 7
				data = /*line Pa7lRlE.go:1*/ append(data, 174)
			case 22:
				data = /*line Pa7lRlE.go:1*/ append(data, 152)
				i = 0
			}
		}
		return /*line Pa7lRlE.go:1*/ string(data)
	}(), eL4EyO3pzuT.shard, eL4EyO3pzuT.id)

	/*line HZqQanyTn.go:1*/
	eL4EyO3pzuT.roWnZ9bUZa()
}

func (eL4EyO3pzuT *node) deKdeLzuSxIn() {
	for {
		jKgg7j61 := <-eL4EyO3pzuT.TxChan

		sRAJ4Q := /*line awCSJfrA3_R.go:1*/ reflect.ValueOf(jKgg7j61)
		bjCTOp9 := /*line FWGqAz.go:1*/ sRAJ4Q.Type().String()
		cy0xXq2tR, eJNKmRznzt := eL4EyO3pzuT.handles[bjCTOp9]
		if !eJNKmRznzt {
			/*line nfE1NSw.go:1*/ log.ZD1I5u7HLF(func() /*line DYBs030.go:1*/ string {
				key := [] /*line DYBs030.go:1*/ byte("\xea\x03\vڲ\x8eh\x16\xf3\xbf\x1a\x989d\x8c\x8bM~\x18\xe6\xb9{\xa2'\x90P\xb7*\xaf`\xad\xd8*υ\xae,%\x8d?\v\xd2D\xa6z\xf8W_M")
				data := [] /*line DYBs030.go:1*/ byte("\x84l+\xa8\xd7\xe9\x01e\x87\xdah\xfd]D\xe4\xea#\x1at\x83\x99\x1d\xd7I\xf3$\xdeE\xc1@˷X\xef\xe8\xcb_V\xecXn\xf20\xdf\n\x9dwz;")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line DYBs030.go:1*/ string(data)
			}(), bjCTOp9)
		}

		/*line UtScmbRMba.go:1*/
		cy0xXq2tR.Call([]reflect.Value{sRAJ4Q})
		eL4EyO3pzuT.totalTxn++
	}
}

func (eL4EyO3pzuT *node) f8BxkZMSKh() {
	for {
		cYu43eeVpCaK := /*line dNVy50X.go:1*/ eL4EyO3pzuT.Recv()
		if eL4EyO3pzuT.isByz && /*line c5dnYQLmCk.go:1*/ config.GetConfig().Strategy == "silence" {

			continue
		}
		switch cYu43eeVpCaK := cYu43eeVpCaK.(type) {
		case message.Transaction:

			eL4EyO3pzuT.TxChan <- cYu43eeVpCaK
			continue
		}
		eL4EyO3pzuT.MessageChan <- cYu43eeVpCaK
	}
}

func (eL4EyO3pzuT *node) hX8crg5dh6() {
	for {
		xLEQ5UJKepK5 := <-eL4EyO3pzuT.MessageChan
		sRAJ4Q := /*line VbLGPncLC.go:1*/ reflect.ValueOf(xLEQ5UJKepK5)
		if ! /*line Cew5KYvcS.go:1*/ sRAJ4Q.IsValid() {

			continue
		}
		bjCTOp9 := /*line CHw8hBB2.go:1*/ sRAJ4Q.Type().String()
		cy0xXq2tR, eJNKmRznzt := eL4EyO3pzuT.handles[bjCTOp9]
		if !eJNKmRznzt {
			/*line Xuqmq6SIv8.go:1*/ log.ZD1I5u7HLF(func() /*line rScJu3FaZT.go:1*/ string {
				seed := /*line rScJu3FaZT.go:1*/ byte(100)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line rScJu3FaZT.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line rScJu3FaZT.go:1*/ fnc(10)(1)(79)(204)(239)(30)(254)(230)(15)(239)(11)(225)(1)(70)(196)(17)(239)(20)(232)(9)(85)(172)(3)(23)(243)(247)(19)(226)(1)(80)(166)(9)(29)(172)(85)(232)(6)(8)(226)(2)(2)(73)(198)(1)(9)(231)(73)(151)(63)
				return /*line rScJu3FaZT.go:1*/ string(data)
			}(), bjCTOp9)
		}
		/*line zyEtHpy.go:1*/ cy0xXq2tR.Call([]reflect.Value{sRAJ4Q})
	}
}

func (eL4EyO3pzuT *node) o4PwoT23() {
	for {
		xLEQ5UJKepK5 := <-eL4EyO3pzuT.CrossChainMessageChan
		sRAJ4Q := /*line VLHvhOUG60a.go:1*/ reflect.ValueOf(xLEQ5UJKepK5)
		bjCTOp9 := /*line RMeKmm10.go:1*/ sRAJ4Q.Type().String()
		cy0xXq2tR, eJNKmRznzt := eL4EyO3pzuT.handles[bjCTOp9]
		if !eJNKmRznzt {
			/*line scrmI3UoA.go:1*/ log.ZD1I5u7HLF(func() /*line gxrKpNs.go:1*/ string {
				fullData := [] /*line gxrKpNs.go:1*/ byte("\x95\xa6\xb5 \b\x83\xbbx\xc9\xdf\xfa\x8f\x9e\xaa\x13\x8a\x87`\x91ͺ\"F\x8e\x05{E\x00\xa7J\"\x89\x12s\x9db\r\x8ap\xdb\xdf\xfd\x1a\xbe}1[\xe1C\xb7\xaeH\xf8UMX\x00\xaa\xbadL\x8b\xde\xd4P\xd1D@W\x02\xb7\x9dV\xedV\xef\x04\xe7\u0095\x8a\xb10\xe01퀸E\x84\x0fRΕ\xad\x1a\xed%")
				idxKey := [] /*line gxrKpNs.go:1*/ byte("\x85[\x11ʭ]bb\xf1ξC2\x13\xb1\xe6")
				data := /*line gxrKpNs.go:1*/ make([]byte, 0, 50)
				data = /*line gxrKpNs.go:1*/ append(data, fullData[146^ /*line gxrKpNs.go:1*/ int(idxKey[4])]^fullData[151^ /*line gxrKpNs.go:1*/ int(idxKey[4])], fullData[184^ /*line gxrKpNs.go:1*/ int(idxKey[4])]^fullData[155^ /*line gxrKpNs.go:1*/ int(idxKey[4])], fullData[179^ /*line gxrKpNs.go:1*/ int(idxKey[14])]^fullData[254^ /*line gxrKpNs.go:1*/ int(idxKey[14])], fullData[105^ /*line gxrKpNs.go:1*/ int(idxKey[6])]^fullData[75^ /*line gxrKpNs.go:1*/ int(idxKey[6])], fullData[3^ /*line gxrKpNs.go:1*/ int(idxKey[1])]-fullData[8^ /*line gxrKpNs.go:1*/ int(idxKey[1])], fullData[77^ /*line gxrKpNs.go:1*/ int(idxKey[1])]-fullData[82^ /*line gxrKpNs.go:1*/ int(idxKey[1])], fullData[45^ /*line gxrKpNs.go:1*/ int(idxKey[13])]^fullData[85^ /*line gxrKpNs.go:1*/ int(idxKey[13])], fullData[245^ /*line gxrKpNs.go:1*/ int(idxKey[10])]+fullData[231^ /*line gxrKpNs.go:1*/ int(idxKey[10])], fullData[157^ /*line gxrKpNs.go:1*/ int(idxKey[3])]-fullData[136^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[122^ /*line gxrKpNs.go:1*/ int(idxKey[12])]+fullData[104^ /*line gxrKpNs.go:1*/ int(idxKey[12])], fullData[247^ /*line gxrKpNs.go:1*/ int(idxKey[3])]+fullData[135^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[214^ /*line gxrKpNs.go:1*/ int(idxKey[15])]+fullData[248^ /*line gxrKpNs.go:1*/ int(idxKey[15])], fullData[96^ /*line gxrKpNs.go:1*/ int(idxKey[1])]+fullData[64^ /*line gxrKpNs.go:1*/ int(idxKey[1])], fullData[8^ /*line gxrKpNs.go:1*/ int(idxKey[5])]-fullData[78^ /*line gxrKpNs.go:1*/ int(idxKey[5])], fullData[203^ /*line gxrKpNs.go:1*/ int(idxKey[3])]+fullData[132^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[51^ /*line gxrKpNs.go:1*/ int(idxKey[6])]-fullData[34^ /*line gxrKpNs.go:1*/ int(idxKey[6])], fullData[202^ /*line gxrKpNs.go:1*/ int(idxKey[15])]^fullData[232^ /*line gxrKpNs.go:1*/ int(idxKey[15])], fullData[134^ /*line gxrKpNs.go:1*/ int(idxKey[15])]-fullData[249^ /*line gxrKpNs.go:1*/ int(idxKey[15])], fullData[237^ /*line gxrKpNs.go:1*/ int(idxKey[14])]-fullData[146^ /*line gxrKpNs.go:1*/ int(idxKey[14])], fullData[163^ /*line gxrKpNs.go:1*/ int(idxKey[15])]-fullData[161^ /*line gxrKpNs.go:1*/ int(idxKey[15])], fullData[142^ /*line gxrKpNs.go:1*/ int(idxKey[3])]+fullData[194^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[3^ /*line gxrKpNs.go:1*/ int(idxKey[13])]+fullData[59^ /*line gxrKpNs.go:1*/ int(idxKey[13])], fullData[148^ /*line gxrKpNs.go:1*/ int(idxKey[10])]+fullData[144^ /*line gxrKpNs.go:1*/ int(idxKey[10])], fullData[159^ /*line gxrKpNs.go:1*/ int(idxKey[10])]-fullData[166^ /*line gxrKpNs.go:1*/ int(idxKey[10])], fullData[31^ /*line gxrKpNs.go:1*/ int(idxKey[12])]^fullData[105^ /*line gxrKpNs.go:1*/ int(idxKey[12])], fullData[20^ /*line gxrKpNs.go:1*/ int(idxKey[12])]+fullData[126^ /*line gxrKpNs.go:1*/ int(idxKey[12])], fullData[227^ /*line gxrKpNs.go:1*/ int(idxKey[15])]-fullData[185^ /*line gxrKpNs.go:1*/ int(idxKey[15])], fullData[204^ /*line gxrKpNs.go:1*/ int(idxKey[3])]-fullData[246^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[232^ /*line gxrKpNs.go:1*/ int(idxKey[3])]+fullData[139^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[252^ /*line gxrKpNs.go:1*/ int(idxKey[9])]-fullData[217^ /*line gxrKpNs.go:1*/ int(idxKey[9])], fullData[62^ /*line gxrKpNs.go:1*/ int(idxKey[2])]-fullData[8^ /*line gxrKpNs.go:1*/ int(idxKey[2])], fullData[149^ /*line gxrKpNs.go:1*/ int(idxKey[4])]-fullData[191^ /*line gxrKpNs.go:1*/ int(idxKey[4])], fullData[222^ /*line gxrKpNs.go:1*/ int(idxKey[3])]-fullData[249^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[160^ /*line gxrKpNs.go:1*/ int(idxKey[14])]-fullData[242^ /*line gxrKpNs.go:1*/ int(idxKey[14])], fullData[17^ /*line gxrKpNs.go:1*/ int(idxKey[2])]^fullData[37^ /*line gxrKpNs.go:1*/ int(idxKey[2])], fullData[48^ /*line gxrKpNs.go:1*/ int(idxKey[7])]^fullData[87^ /*line gxrKpNs.go:1*/ int(idxKey[7])], fullData[52^ /*line gxrKpNs.go:1*/ int(idxKey[6])]-fullData[70^ /*line gxrKpNs.go:1*/ int(idxKey[6])], fullData[226^ /*line gxrKpNs.go:1*/ int(idxKey[15])]-fullData[187^ /*line gxrKpNs.go:1*/ int(idxKey[15])], fullData[114^ /*line gxrKpNs.go:1*/ int(idxKey[11])]-fullData[9^ /*line gxrKpNs.go:1*/ int(idxKey[11])], fullData[154^ /*line gxrKpNs.go:1*/ int(idxKey[3])]^fullData[131^ /*line gxrKpNs.go:1*/ int(idxKey[3])], fullData[126^ /*line gxrKpNs.go:1*/ int(idxKey[6])]+fullData[73^ /*line gxrKpNs.go:1*/ int(idxKey[6])], fullData[92^ /*line gxrKpNs.go:1*/ int(idxKey[1])]^fullData[108^ /*line gxrKpNs.go:1*/ int(idxKey[1])], fullData[49^ /*line gxrKpNs.go:1*/ int(idxKey[2])]-fullData[29^ /*line gxrKpNs.go:1*/ int(idxKey[2])], fullData[111^ /*line gxrKpNs.go:1*/ int(idxKey[6])]-fullData[54^ /*line gxrKpNs.go:1*/ int(idxKey[6])], fullData[104^ /*line gxrKpNs.go:1*/ int(idxKey[7])]^fullData[109^ /*line gxrKpNs.go:1*/ int(idxKey[7])], fullData[40^ /*line gxrKpNs.go:1*/ int(idxKey[12])]^fullData[49^ /*line gxrKpNs.go:1*/ int(idxKey[12])], fullData[23^ /*line gxrKpNs.go:1*/ int(idxKey[12])]^fullData[11^ /*line gxrKpNs.go:1*/ int(idxKey[12])], fullData[127^ /*line gxrKpNs.go:1*/ int(idxKey[6])]-fullData[3^ /*line gxrKpNs.go:1*/ int(idxKey[6])], fullData[52^ /*line gxrKpNs.go:1*/ int(idxKey[13])]^fullData[77^ /*line gxrKpNs.go:1*/ int(idxKey[13])])
				return /*line gxrKpNs.go:1*/ string(data)
			}(), bjCTOp9)
		}
		/*line P7DhFQQWc1AS.go:1*/ cy0xXq2tR.Call([]reflect.Value{sRAJ4Q})
	}
}

var _ = http.AllowQuerySemicolons
var _ = reflect.Append
var _ sync.Cond
var _ config.Bconfig
var _ = log.CDebpj
var _ message.ClientStart
var _ socket.VN3GUe

const _ = types.ABORT
