//line :1
package node

import (
	"net/http"
	"reflect"
	"sync"

	blockchain "unishard/blockchain"
	config "unishard/config"
	crypto "unishard/crypto"
	log "unishard/log"
	mempool "unishard/mempool"
	message "unishard/message"
	socket "unishard/socket"
	types "unishard/types"
)

type CoordBlockBuilder interface {
	socket.VN3GUe

	GetIP() string
	GetShard() types.Shard
	GetBlockChain() *blockchain.Y3t7C8s0m
	AddLocalTransaction(fTdm1vt5 *message.Transaction)
	AddNewCrossTransaction(kcAX0S *message.Transaction)
	AddSendedCrossTransaction(kcAX0S *message.Transaction)
	CreateCoordinationBlockWithoutHeader(jNjVFG0Hrn *blockchain.CoordinationBlockData) *blockchain.CoordinationBlockWithoutHeader
	Run()
	Retry(yy2jgIELlF message.Transaction)
	Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{})
}

type coordblockbuilder struct {
	ip    string
	shard types.Shard

	socket.VN3GUe
	server *http.Server

	handles     map[string]reflect.Value
	MessageChan chan interface{}
	TxChan      chan interface{}
	BlockChan   chan interface{}

	sync.RWMutex

	bc    *blockchain.Y3t7C8s0m
	ltmp  *mempool.Producer
	nctmp *mempool.Producer
	sctmp *mempool.Producer
	srtmp *mempool.Producer
	txcnt int
}

func NewCoordinationBlockBuilder(lq54woHa3C string, js85nByJG types.Shard) CoordBlockBuilder {
	iai_HgOpk3 := /*line H83HcCHOPQ4.go:1*/ new(coordblockbuilder)
	iai_HgOpk3.ip = lq54woHa3C
	iai_HgOpk3.shard = js85nByJG
	iai_HgOpk3.VN3GUe = /*line iyax6yvH7J.go:1*/ socket.P2LcT9Ia(lq54woHa3C, js85nByJG, config.Configuration.Addrs)
	iai_HgOpk3.handles = /*line Ad_LU5rPdX.go:1*/ make(map[string]reflect.Value)
	iai_HgOpk3.MessageChan = /*line VzDa9sDw_u.go:1*/ make(chan interface{}, config.Configuration.ChanBufferSize)
	iai_HgOpk3.bc = /*line c5xfxPQWw.go:1*/ blockchain.GEzkgtuHo0(iai_HgOpk3.shard /*line LQ1tadHef.go:1*/, config.GetConfig().DefaultBalance)
	iai_HgOpk3.ltmp = /*line aU30caV4X.go:1*/ mempool.NewProducer()
	iai_HgOpk3.nctmp = /*line GJSOVHRH.go:1*/ mempool.NewProducer()
	iai_HgOpk3.sctmp = /*line vJDBGlH.go:1*/ mempool.NewProducer()
	iai_HgOpk3.srtmp = /*line NtqlffLb9.go:1*/ mempool.NewProducer()
	iai_HgOpk3.txcnt = 0

	return iai_HgOpk3
}

func (iai_HgOpk3 *coordblockbuilder) GetIP() string {
	return iai_HgOpk3.ip
}

func (iai_HgOpk3 *coordblockbuilder) GetShard() types.Shard {
	return iai_HgOpk3.shard
}

func (iai_HgOpk3 *coordblockbuilder) GetBlockChain() *blockchain.Y3t7C8s0m {
	return iai_HgOpk3.bc
}

func (iai_HgOpk3 *coordblockbuilder) AddLocalTransaction(fTdm1vt5 *message.Transaction) {
	/*line xvWGCwiLzAe.go:1*/ iai_HgOpk3.ltmp.AddTxn(fTdm1vt5)
}

func (iai_HgOpk3 *coordblockbuilder) AddNewCrossTransaction(kcAX0S *message.Transaction) {
	/*line GshQLEM.go:1*/ iai_HgOpk3.nctmp.AddTxn(kcAX0S)
}

func (iai_HgOpk3 *coordblockbuilder) AddSendedCrossTransaction(kcAX0S *message.Transaction) {
	/*line YQsfUDBwU.go:1*/ iai_HgOpk3.sctmp.AddTxn(kcAX0S)
}

func (iai_HgOpk3 *coordblockbuilder) AddSequenceReadyCrossTransaction(pd906M_h *message.Transaction) {
	/*line sfhuS1sNyuN.go:1*/ iai_HgOpk3.srtmp.AddTxn(pd906M_h)
}

func (iai_HgOpk3 *coordblockbuilder) CreateCoordinationBlockWithoutHeader(jNjVFG0Hrn *blockchain.CoordinationBlockData) *blockchain.CoordinationBlockWithoutHeader {
	s6GAIqC6Mda7 := /*line w36mQdO3.go:1*/ new(blockchain.CoordinationBlockWithoutHeader)
	s6GAIqC6Mda7.CoordinationBlockData = jNjVFG0Hrn
	omuFrxdTot, _ := /*line U5QFOgae8_96.go:1*/ crypto.PrivSign( /*line O2fr5MyaNygP.go:1*/ crypto.IDToByte( /*line Hf6WhpCVTfc.go:1*/ s6GAIqC6Mda7.MakeHash(jNjVFG0Hrn)), nil)
	s6GAIqC6Mda7.CommitteeSignature = /*line oSEVR_t.go:1*/ append(s6GAIqC6Mda7.CommitteeSignature, omuFrxdTot)

	return s6GAIqC6Mda7
}

func (iai_HgOpk3 *coordblockbuilder) Retry(yy2jgIELlF message.Transaction) {
	/*line kZ0Xw0AXxoiA.go:1*/ log.ViJSfja(func() /*line anMqTgp.go:1*/ string {
		seed := /*line anMqTgp.go:1*/ byte(241)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line anMqTgp.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line anMqTgp.go:1*/ fnc(113)(10)(3)(244)(8)(247)(19)(244)(3)(248)(1)(13)(174)(5)(81)(170)(82)(243)(15)(254)(7)(167)(82)(243)(12)(244)(16)(254)(1)(172)(5)(81)
		return /*line anMqTgp.go:1*/ string(data)
	}(), iai_HgOpk3.shard, yy2jgIELlF)
	iai_HgOpk3.TxChan <- yy2jgIELlF
}

func (iai_HgOpk3 *coordblockbuilder) Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{}) {
	vgqeZiHun9F := /*line JOAIAA.go:1*/ reflect.TypeOf(cYu43eeVpCaK)

	cLL2dKRL8 := /*line cml5mRq9.go:1*/ reflect.ValueOf(cy0xXq2tR)

	if /*line TpyQ6LS5Id.go:1*/ cLL2dKRL8.Kind() != reflect.Func {
		/*line hy7DGftw.go:1*/ panic(func() /*line kvUT1i.go:1*/ string {
			seed := /*line kvUT1i.go:1*/ byte(52)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line kvUT1i.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line kvUT1i.go:1*/ fnc(52)(249)(13)(246)(8)(249)(187)(70)(15)(249)(245)(17)(245)(6)(255)(178)(73)(10)(173)(78)(1)(5)(172)(70)(15)(249)(245)
			return /*line kvUT1i.go:1*/ string(data)
		}())
	}

	if /*line ipES4jher.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line B77nqukwe.go:1*/ log.Jp980o6YjM("%v %v" /*line k8uEURK.go:1*/, cLL2dKRL8.Type().In(0), vgqeZiHun9F)
		/*line idnU4f.go:1*/ panic(func() /*line MfaotR68ZU.go:1*/ string {
			data := [] /*line MfaotR68ZU.go:1*/ byte("fu\xa6݉\xc00\xa4\xb1 \x02synoI\xdd\x04")
			positions := [...]byte{4, 6, 8, 12, 3, 7, 5, 3, 16, 17, 10, 16, 2, 10, 7, 17, 10, 7, 3, 17, 16, 7, 2, 7, 15, 12}
			for i := 0; i < 26; i += 2 {
				localKey := /*line MfaotR68ZU.go:1*/ byte(i) + /*line MfaotR68ZU.go:1*/ byte(positions[i]^positions[i+1]) + 14
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line MfaotR68ZU.go:1*/ string(data)
		}())
	}

	if /*line vaW5wnQg.go:1*/ cLL2dKRL8.Kind() != reflect.Func || /*line tRzLj2q7nuR.go:1*/ cLL2dKRL8.Type().NumIn() != 1 || /*line EMXYJRGn.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line hlkYGJ5VAKq.go:1*/ panic(func() /*line NzBS9c.go:1*/ string {
			key := [] /*line NzBS9c.go:1*/ byte("/\xfdX\xaey\n\xfdy7ކي\xd5Kz6\xaa\xa2\xfc\xc8 V-\x91\xa8!\xabs\x98")
			data := [] /*line NzBS9c.go:1*/ byte("\xa1b\xbf\x17\xec~b\xebWF\xe7G\xeeA\xb0\x9a\x9c\x1f\x10_<\x89ś\xb1\r\x93\x1d\xe2\n")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line NzBS9c.go:1*/ string(data)
		}())
	}

	iai_HgOpk3.handles[ /*line HAT6Zo.go:1*/ vgqeZiHun9F.String()] = cLL2dKRL8
}

func (iai_HgOpk3 *coordblockbuilder) Run() {
	/*line iS_9ZmxZ6.go:1*/ log.XWmhQ9(func() /*line TAuUL1.go:1*/ string {
		fullData := [] /*line TAuUL1.go:1*/ byte("\x81\xaf\xe7\xd61\xffM\xb08h\xc6\x7f\tK\x06\x03\x94iB\x94)\x11,\xeck\x14Zz\x8bg\x80\xce[\f\xf1\x00\x00r\x183W\xea\xbd}\xf3i\xcd2\x8a\x0f\xa1\f4\x8d\xb9AX%z>~\x0e\xf8.J\xb8")
		idxKey := [] /*line TAuUL1.go:1*/ byte("\n\xf6\xfe\xcbg\x0e/\xfbo\xc9\xc7\xe6\xb5g\xe6\x87")
		data := /*line TAuUL1.go:1*/ make([]byte, 0, 34)
		data = /*line TAuUL1.go:1*/ append(data, fullData[71^ /*line TAuUL1.go:1*/ int(idxKey[4])]+fullData[101^ /*line TAuUL1.go:1*/ int(idxKey[4])], fullData[142^ /*line TAuUL1.go:1*/ int(idxKey[12])]+fullData[138^ /*line TAuUL1.go:1*/ int(idxKey[12])], fullData[255^ /*line TAuUL1.go:1*/ int(idxKey[9])]-fullData[137^ /*line TAuUL1.go:1*/ int(idxKey[9])], fullData[66^ /*line TAuUL1.go:1*/ int(idxKey[13])]-fullData[86^ /*line TAuUL1.go:1*/ int(idxKey[13])], fullData[38^ /*line TAuUL1.go:1*/ int(idxKey[5])]+fullData[23^ /*line TAuUL1.go:1*/ int(idxKey[5])], fullData[204^ /*line TAuUL1.go:1*/ int(idxKey[11])]^fullData[227^ /*line TAuUL1.go:1*/ int(idxKey[11])], fullData[205^ /*line TAuUL1.go:1*/ int(idxKey[3])]^fullData[195^ /*line TAuUL1.go:1*/ int(idxKey[3])], fullData[197^ /*line TAuUL1.go:1*/ int(idxKey[11])]+fullData[247^ /*line TAuUL1.go:1*/ int(idxKey[11])], fullData[46^ /*line TAuUL1.go:1*/ int(idxKey[0])]-fullData[26^ /*line TAuUL1.go:1*/ int(idxKey[0])], fullData[0^ /*line TAuUL1.go:1*/ int(idxKey[6])]-fullData[48^ /*line TAuUL1.go:1*/ int(idxKey[6])], fullData[43^ /*line TAuUL1.go:1*/ int(idxKey[0])]^fullData[39^ /*line TAuUL1.go:1*/ int(idxKey[0])], fullData[103^ /*line TAuUL1.go:1*/ int(idxKey[4])]+fullData[69^ /*line TAuUL1.go:1*/ int(idxKey[4])], fullData[17^ /*line TAuUL1.go:1*/ int(idxKey[0])]^fullData[16^ /*line TAuUL1.go:1*/ int(idxKey[0])], fullData[95^ /*line TAuUL1.go:1*/ int(idxKey[4])]^fullData[76^ /*line TAuUL1.go:1*/ int(idxKey[4])], fullData[227^ /*line TAuUL1.go:1*/ int(idxKey[2])]^fullData[235^ /*line TAuUL1.go:1*/ int(idxKey[2])], fullData[59^ /*line TAuUL1.go:1*/ int(idxKey[6])]+fullData[32^ /*line TAuUL1.go:1*/ int(idxKey[6])], fullData[215^ /*line TAuUL1.go:1*/ int(idxKey[3])]-fullData[211^ /*line TAuUL1.go:1*/ int(idxKey[3])], fullData[13^ /*line TAuUL1.go:1*/ int(idxKey[5])]^fullData[34^ /*line TAuUL1.go:1*/ int(idxKey[5])], fullData[225^ /*line TAuUL1.go:1*/ int(idxKey[14])]^fullData[236^ /*line TAuUL1.go:1*/ int(idxKey[14])], fullData[233^ /*line TAuUL1.go:1*/ int(idxKey[2])]+fullData[202^ /*line TAuUL1.go:1*/ int(idxKey[2])], fullData[236^ /*line TAuUL1.go:1*/ int(idxKey[2])]^fullData[250^ /*line TAuUL1.go:1*/ int(idxKey[2])], fullData[252^ /*line TAuUL1.go:1*/ int(idxKey[3])]-fullData[229^ /*line TAuUL1.go:1*/ int(idxKey[3])], fullData[245^ /*line TAuUL1.go:1*/ int(idxKey[11])]-fullData[193^ /*line TAuUL1.go:1*/ int(idxKey[11])], fullData[89^ /*line TAuUL1.go:1*/ int(idxKey[4])]^fullData[87^ /*line TAuUL1.go:1*/ int(idxKey[4])], fullData[221^ /*line TAuUL1.go:1*/ int(idxKey[3])]-fullData[138^ /*line TAuUL1.go:1*/ int(idxKey[3])], fullData[60^ /*line TAuUL1.go:1*/ int(idxKey[5])]+fullData[5^ /*line TAuUL1.go:1*/ int(idxKey[5])], fullData[224^ /*line TAuUL1.go:1*/ int(idxKey[2])]-fullData[195^ /*line TAuUL1.go:1*/ int(idxKey[2])], fullData[54^ /*line TAuUL1.go:1*/ int(idxKey[0])]-fullData[6^ /*line TAuUL1.go:1*/ int(idxKey[0])], fullData[38^ /*line TAuUL1.go:1*/ int(idxKey[6])]^fullData[33^ /*line TAuUL1.go:1*/ int(idxKey[6])], fullData[93^ /*line TAuUL1.go:1*/ int(idxKey[4])]-fullData[84^ /*line TAuUL1.go:1*/ int(idxKey[4])], fullData[73^ /*line TAuUL1.go:1*/ int(idxKey[8])]-fullData[110^ /*line TAuUL1.go:1*/ int(idxKey[8])], fullData[34^ /*line TAuUL1.go:1*/ int(idxKey[6])]^fullData[22^ /*line TAuUL1.go:1*/ int(idxKey[6])], fullData[39^ /*line TAuUL1.go:1*/ int(idxKey[5])]^fullData[59^ /*line TAuUL1.go:1*/ int(idxKey[5])])
		return /*line TAuUL1.go:1*/ string(data)
	}(), iai_HgOpk3.ip, iai_HgOpk3.shard)

	go /*line Qxb0JE.go:1*/ iai_HgOpk3.hX8crg5dh6()
	go /*line GY1mOpi6UAem.go:1*/ iai_HgOpk3.f8BxkZMSKh()
	/*line _38uGN.go:1*/ iai_HgOpk3.roWnZ9bUZa()
}

func (iai_HgOpk3 *coordblockbuilder) hX8crg5dh6() {
	for {
		xLEQ5UJKepK5 := <-iai_HgOpk3.MessageChan
		sRAJ4Q := /*line C1JEquKjlZd.go:1*/ reflect.ValueOf(xLEQ5UJKepK5)
		bjCTOp9 := /*line M6ni3M3.go:1*/ sRAJ4Q.Type().String()
		cy0xXq2tR, eJNKmRznzt := iai_HgOpk3.handles[bjCTOp9]
		if !eJNKmRznzt {
			/*line FESpwPkJY.go:1*/ log.ZD1I5u7HLF(func() /*line lPxh_Mn.go:1*/ string {
				fullData := [] /*line lPxh_Mn.go:1*/ byte("3\x9c\xac\xdd\x1f/V\xa60=wy\x84-\x0f\bQ5D\x80\"ً2\xf9i\x12f\xc7h\xe1\xf6\x11:54b\xe8\xcc\x02\xb8\xee\xa7\xfc\xf7\xf7\x96\x12\xa92\xa33?\x12\x95\xa9\x12ASɬ\xae\x93\xa4\xaf#)\xea\xaf\xe8!\x99\xf3U\xe3\xff^pDgN\xc0\x7f\xbbԼ\xf6\x1f\av\x9e\x15\x16\x1d\xe4\x86x\xd5")
				idxKey := [] /*line lPxh_Mn.go:1*/ byte("\x83\xb6\xea\xf7\xbdW\x8c\x96K\x98!\x11")
				data := /*line lPxh_Mn.go:1*/ make([]byte, 0, 50)
				data = /*line lPxh_Mn.go:1*/ append(data, fullData[174^ /*line lPxh_Mn.go:1*/ int(idxKey[4])]-fullData[167^ /*line lPxh_Mn.go:1*/ int(idxKey[4])], fullData[206^ /*line lPxh_Mn.go:1*/ int(idxKey[9])]+fullData[147^ /*line lPxh_Mn.go:1*/ int(idxKey[9])], fullData[118^ /*line lPxh_Mn.go:1*/ int(idxKey[10])]-fullData[106^ /*line lPxh_Mn.go:1*/ int(idxKey[10])], fullData[81^ /*line lPxh_Mn.go:1*/ int(idxKey[11])]^fullData[18^ /*line lPxh_Mn.go:1*/ int(idxKey[11])], fullData[47^ /*line lPxh_Mn.go:1*/ int(idxKey[11])]^fullData[14^ /*line lPxh_Mn.go:1*/ int(idxKey[11])], fullData[87^ /*line lPxh_Mn.go:1*/ int(idxKey[5])]+fullData[116^ /*line lPxh_Mn.go:1*/ int(idxKey[5])], fullData[8^ /*line lPxh_Mn.go:1*/ int(idxKey[8])]+fullData[25^ /*line lPxh_Mn.go:1*/ int(idxKey[8])], fullData[225^ /*line lPxh_Mn.go:1*/ int(idxKey[3])]+fullData[178^ /*line lPxh_Mn.go:1*/ int(idxKey[3])], fullData[130^ /*line lPxh_Mn.go:1*/ int(idxKey[0])]^fullData[166^ /*line lPxh_Mn.go:1*/ int(idxKey[0])], fullData[199^ /*line lPxh_Mn.go:1*/ int(idxKey[3])]-fullData[185^ /*line lPxh_Mn.go:1*/ int(idxKey[3])], fullData[197^ /*line lPxh_Mn.go:1*/ int(idxKey[6])]-fullData[198^ /*line lPxh_Mn.go:1*/ int(idxKey[6])], fullData[208^ /*line lPxh_Mn.go:1*/ int(idxKey[9])]^fullData[182^ /*line lPxh_Mn.go:1*/ int(idxKey[9])], fullData[152^ /*line lPxh_Mn.go:1*/ int(idxKey[0])]^fullData[164^ /*line lPxh_Mn.go:1*/ int(idxKey[0])], fullData[161^ /*line lPxh_Mn.go:1*/ int(idxKey[9])]-fullData[222^ /*line lPxh_Mn.go:1*/ int(idxKey[9])], fullData[174^ /*line lPxh_Mn.go:1*/ int(idxKey[6])]+fullData[191^ /*line lPxh_Mn.go:1*/ int(idxKey[6])], fullData[160^ /*line lPxh_Mn.go:1*/ int(idxKey[4])]-fullData[229^ /*line lPxh_Mn.go:1*/ int(idxKey[4])], fullData[251^ /*line lPxh_Mn.go:1*/ int(idxKey[3])]-fullData[171^ /*line lPxh_Mn.go:1*/ int(idxKey[3])], fullData[18^ /*line lPxh_Mn.go:1*/ int(idxKey[8])]-fullData[115^ /*line lPxh_Mn.go:1*/ int(idxKey[8])], fullData[237^ /*line lPxh_Mn.go:1*/ int(idxKey[1])]-fullData[129^ /*line lPxh_Mn.go:1*/ int(idxKey[1])], fullData[238^ /*line lPxh_Mn.go:1*/ int(idxKey[4])]-fullData[187^ /*line lPxh_Mn.go:1*/ int(idxKey[4])], fullData[172^ /*line lPxh_Mn.go:1*/ int(idxKey[9])]-fullData[156^ /*line lPxh_Mn.go:1*/ int(idxKey[9])], fullData[99^ /*line lPxh_Mn.go:1*/ int(idxKey[8])]+fullData[118^ /*line lPxh_Mn.go:1*/ int(idxKey[8])], fullData[46^ /*line lPxh_Mn.go:1*/ int(idxKey[11])]-fullData[20^ /*line lPxh_Mn.go:1*/ int(idxKey[11])], fullData[16^ /*line lPxh_Mn.go:1*/ int(idxKey[5])]+fullData[54^ /*line lPxh_Mn.go:1*/ int(idxKey[5])], fullData[151^ /*line lPxh_Mn.go:1*/ int(idxKey[1])]+fullData[244^ /*line lPxh_Mn.go:1*/ int(idxKey[1])], fullData[85^ /*line lPxh_Mn.go:1*/ int(idxKey[8])]^fullData[125^ /*line lPxh_Mn.go:1*/ int(idxKey[8])], fullData[187^ /*line lPxh_Mn.go:1*/ int(idxKey[7])]^fullData[204^ /*line lPxh_Mn.go:1*/ int(idxKey[7])], fullData[92^ /*line lPxh_Mn.go:1*/ int(idxKey[8])]+fullData[66^ /*line lPxh_Mn.go:1*/ int(idxKey[8])], fullData[222^ /*line lPxh_Mn.go:1*/ int(idxKey[0])]-fullData[199^ /*line lPxh_Mn.go:1*/ int(idxKey[0])], fullData[78^ /*line lPxh_Mn.go:1*/ int(idxKey[11])]^fullData[22^ /*line lPxh_Mn.go:1*/ int(idxKey[11])], fullData[55^ /*line lPxh_Mn.go:1*/ int(idxKey[5])]-fullData[98^ /*line lPxh_Mn.go:1*/ int(idxKey[5])], fullData[185^ /*line lPxh_Mn.go:1*/ int(idxKey[7])]-fullData[164^ /*line lPxh_Mn.go:1*/ int(idxKey[7])], fullData[221^ /*line lPxh_Mn.go:1*/ int(idxKey[6])]-fullData[220^ /*line lPxh_Mn.go:1*/ int(idxKey[6])], fullData[128^ /*line lPxh_Mn.go:1*/ int(idxKey[9])]^fullData[141^ /*line lPxh_Mn.go:1*/ int(idxKey[9])], fullData[227^ /*line lPxh_Mn.go:1*/ int(idxKey[4])]-fullData[183^ /*line lPxh_Mn.go:1*/ int(idxKey[4])], fullData[220^ /*line lPxh_Mn.go:1*/ int(idxKey[3])]+fullData[238^ /*line lPxh_Mn.go:1*/ int(idxKey[3])], fullData[192^ /*line lPxh_Mn.go:1*/ int(idxKey[2])]^fullData[190^ /*line lPxh_Mn.go:1*/ int(idxKey[2])], fullData[173^ /*line lPxh_Mn.go:1*/ int(idxKey[4])]^fullData[169^ /*line lPxh_Mn.go:1*/ int(idxKey[4])], fullData[219^ /*line lPxh_Mn.go:1*/ int(idxKey[7])]^fullData[182^ /*line lPxh_Mn.go:1*/ int(idxKey[7])], fullData[192^ /*line lPxh_Mn.go:1*/ int(idxKey[6])]-fullData[160^ /*line lPxh_Mn.go:1*/ int(idxKey[6])], fullData[138^ /*line lPxh_Mn.go:1*/ int(idxKey[7])]-fullData[178^ /*line lPxh_Mn.go:1*/ int(idxKey[7])], fullData[135^ /*line lPxh_Mn.go:1*/ int(idxKey[1])]+fullData[159^ /*line lPxh_Mn.go:1*/ int(idxKey[1])], fullData[229^ /*line lPxh_Mn.go:1*/ int(idxKey[3])]+fullData[255^ /*line lPxh_Mn.go:1*/ int(idxKey[3])], fullData[70^ /*line lPxh_Mn.go:1*/ int(idxKey[5])]-fullData[2^ /*line lPxh_Mn.go:1*/ int(idxKey[5])], fullData[135^ /*line lPxh_Mn.go:1*/ int(idxKey[4])]^fullData[252^ /*line lPxh_Mn.go:1*/ int(idxKey[4])], fullData[163^ /*line lPxh_Mn.go:1*/ int(idxKey[9])]^fullData[154^ /*line lPxh_Mn.go:1*/ int(idxKey[9])], fullData[113^ /*line lPxh_Mn.go:1*/ int(idxKey[5])]-fullData[107^ /*line lPxh_Mn.go:1*/ int(idxKey[5])], fullData[153^ /*line lPxh_Mn.go:1*/ int(idxKey[7])]^fullData[155^ /*line lPxh_Mn.go:1*/ int(idxKey[7])], fullData[4^ /*line lPxh_Mn.go:1*/ int(idxKey[8])]+fullData[69^ /*line lPxh_Mn.go:1*/ int(idxKey[8])])
				return /*line lPxh_Mn.go:1*/ string(data)
			}(), bjCTOp9)
		}
		/*line hEcRjlGka.go:1*/ cy0xXq2tR.Call([]reflect.Value{sRAJ4Q})
	}
}

func (iai_HgOpk3 *coordblockbuilder) f8BxkZMSKh() {
	for {
		cYu43eeVpCaK := /*line ji_cJFgy.go:1*/ iai_HgOpk3.Recv()
		iai_HgOpk3.MessageChan <- cYu43eeVpCaK
	}
}

var _ = http.AllowQuerySemicolons
var _ = reflect.Append
var _ sync.Cond
var _ blockchain.Accept
var _ config.Bconfig
var _ crypto.Pp__49cd
var _ = log.CDebpj
var _ mempool.DUUdwSwZTCm
var _ message.ClientStart
var _ socket.VN3GUe

const _ = types.ABORT
