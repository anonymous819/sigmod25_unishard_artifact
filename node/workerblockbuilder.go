//line :1
package node

import (
	"net/http"
	"reflect"
	"sync"

	blockchain "unishard/blockchain"
	config "unishard/config"
	log "unishard/log"
	mempool "unishard/mempool"
	message "unishard/message"
	socket "unishard/socket"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BlockBuilder interface {
	socket.VN3GUe

	GetIP() string
	GetShard() types.Shard
	AddLocalTransaction(fTdm1vt5 *message.Transaction)
	GetRemainingLocalTx() int
	AddNewCrossTransaction(kcAX0S *message.Transaction)
	AddGlobalCoordinationSequence(kcAX0S *message.Transaction)
	AddGlobalSnapshot(nMYJoTttq *blockchain.LocalSnapshot)
	AddGlobalContractBundle(jZ3wbGql *blockchain.LocalContract)
	GetRemainingGlobalCoordinationSequence() []*message.Transaction
	GeneratePayload() *blockchain.Qhrpt6Ktz9
	GetBlockHeight() types.BlockHeight
	SetBlockHeight(vjxpjblE3neC types.BlockHeight)
	GetGC() types.BlockHeight
	SetGC(vjxpjblE3neC types.BlockHeight)
	Run()
	Retry(yy2jgIELlF message.Transaction)
	Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{})
}

type blockbuilder struct {
	ip    string
	shard types.Shard

	socket.VN3GUe
	server *http.Server

	handles     map[string]reflect.Value
	MessageChan chan interface{}
	TxChan      chan interface{}
	BlockChan   chan interface{}

	sync.RWMutex

	LocalMempool                 *mempool.Producer
	CrossMempool                 *mempool.Producer
	GlobalCoordinationSequence   []*message.Transaction
	GlobalSnapshot               []*blockchain.LocalSnapshot
	GlobalContractBundle         []*blockchain.LocalContract
	fdcm                         []*common.Address
	currentBlockHeight           types.BlockHeight
	latestCoordinatorBlockHeight types.BlockHeight
	txcnt                        int
}

type QeZfe7VXs struct {
	ZYK8T28dqat []*message.Transaction
	SsOFF0Vgfl  []*message.Transaction
}

func NewWorkerBlockBuilder(lq54woHa3C string, js85nByJG types.Shard) BlockBuilder {
	iai_HgOpk3 := /*line OQGJlpa.go:1*/ new(blockbuilder)
	iai_HgOpk3.ip = lq54woHa3C
	iai_HgOpk3.shard = js85nByJG
	iai_HgOpk3.VN3GUe = /*line EEeTv9CC_U.go:1*/ socket.P2LcT9Ia(lq54woHa3C, js85nByJG, config.Configuration.Addrs)
	iai_HgOpk3.handles = /*line YVanQw.go:1*/ make(map[string]reflect.Value)
	iai_HgOpk3.MessageChan = /*line mSd7r90LVsG.go:1*/ make(chan interface{}, config.Configuration.ChanBufferSize)
	iai_HgOpk3.LocalMempool = /*line NNeR8xxPWa.go:1*/ mempool.NewProducer()
	iai_HgOpk3.CrossMempool = /*line mIqi5aYoG.go:1*/ mempool.NewProducer()
	iai_HgOpk3.GlobalCoordinationSequence = /*line lyIskvF_SDv.go:1*/ make([]*message.Transaction, 0)
	iai_HgOpk3.GlobalSnapshot = /*line GzRhXnJ0GZ.go:1*/ make([]*blockchain.LocalSnapshot, 0)
	iai_HgOpk3.GlobalContractBundle = /*line oG43Fh.go:1*/ make([]*blockchain.LocalContract, 0)
	iai_HgOpk3.fdcm = nil
	iai_HgOpk3.currentBlockHeight = 0
	iai_HgOpk3.latestCoordinatorBlockHeight = 0
	iai_HgOpk3.txcnt = 0

	return iai_HgOpk3
}

func (iai_HgOpk3 *blockbuilder) GetIP() string {
	return iai_HgOpk3.ip
}

func (iai_HgOpk3 *blockbuilder) GetShard() types.Shard {
	return iai_HgOpk3.shard
}

func (iai_HgOpk3 *blockbuilder) AddLocalTransaction(fTdm1vt5 *message.Transaction) {
	/*line O4j2809ZpUmF.go:1*/ iai_HgOpk3.LocalMempool.AddTxn(fTdm1vt5)
}

func (iai_HgOpk3 *blockbuilder) GetRemainingLocalTx() int {
	return /*line MRxWxvH.go:1*/ iai_HgOpk3.LocalMempool.Size()
}

func (iai_HgOpk3 *blockbuilder) AddNewCrossTransaction(kcAX0S *message.Transaction) {
	/*line MMf6McKpOU.go:1*/ iai_HgOpk3.CrossMempool.AddTxn(kcAX0S)
}

func (iai_HgOpk3 *blockbuilder) AddGlobalCoordinationSequence(kcAX0S *message.Transaction) {
	iai_HgOpk3.GlobalCoordinationSequence = /*line l7WkhG4fb.go:1*/ append(iai_HgOpk3.GlobalCoordinationSequence, kcAX0S)
}

func (iai_HgOpk3 *blockbuilder) AddGlobalSnapshot(nMYJoTttq *blockchain.LocalSnapshot) {
	iai_HgOpk3.GlobalSnapshot = /*line BtOD5Hqe.go:1*/ append(iai_HgOpk3.GlobalSnapshot, nMYJoTttq)
}

func (iai_HgOpk3 *blockbuilder) AddGlobalContractBundle(jZ3wbGql *blockchain.LocalContract) {
	iai_HgOpk3.GlobalContractBundle = /*line rUVO0K.go:1*/ append(iai_HgOpk3.GlobalContractBundle, jZ3wbGql)
}

func (iai_HgOpk3 *blockbuilder) GetRemainingGlobalCoordinationSequence() []*message.Transaction {
	return iai_HgOpk3.GlobalCoordinationSequence
}

func (iai_HgOpk3 *blockbuilder) GetBlockHeight() types.BlockHeight {
	return iai_HgOpk3.currentBlockHeight
}

func (iai_HgOpk3 *blockbuilder) SetBlockHeight(vjxpjblE3neC types.BlockHeight) {
	iai_HgOpk3.currentBlockHeight = vjxpjblE3neC
}

func (iai_HgOpk3 *blockbuilder) GetGC() types.BlockHeight {
	return iai_HgOpk3.latestCoordinatorBlockHeight
}

func (iai_HgOpk3 *blockbuilder) SetGC(vjxpjblE3neC types.BlockHeight) {
	iai_HgOpk3.latestCoordinatorBlockHeight = vjxpjblE3neC
}

func (iai_HgOpk3 *blockbuilder) GeneratePayload() *blockchain.Qhrpt6Ktz9 {
	a8RkHaax8 := /*line bCu2Jsp.go:1*/ make([]*message.Transaction, 0)
	dgr3gjp := /*line WD0oFdydSf.go:1*/ make([]*blockchain.LocalSnapshot, 0)
	bMT754HQar := /*line hQ03zDRJ51b.go:1*/ make([]*blockchain.LocalContract, 0)

	a8RkHaax8 = /*line LnT6TO.go:1*/ append(a8RkHaax8, iai_HgOpk3.GlobalCoordinationSequence...)
	if /*line JcXfp79hO.go:1*/ len(iai_HgOpk3.GlobalCoordinationSequence) > /*line L2vudBHBa.go:1*/ len(a8RkHaax8) {
		iai_HgOpk3.GlobalCoordinationSequence = iai_HgOpk3.GlobalCoordinationSequence[ /*line eGXFGTulAEAq.go:1*/ len(a8RkHaax8):]
	} else {
		iai_HgOpk3.GlobalCoordinationSequence = []*message.Transaction{}
	}

	dgr3gjp = /*line F83Rffqe4a.go:1*/ append(dgr3gjp, iai_HgOpk3.GlobalSnapshot...)
	if /*line IgXi5AzIJm.go:1*/ len(iai_HgOpk3.GlobalSnapshot) > /*line h_O6gg2.go:1*/ len(dgr3gjp) {
		iai_HgOpk3.GlobalSnapshot = iai_HgOpk3.GlobalSnapshot[ /*line _1dTMD4O.go:1*/ len(dgr3gjp):]
	} else {
		iai_HgOpk3.GlobalSnapshot = []*blockchain.LocalSnapshot{}
	}

	bMT754HQar = /*line L_CoP6Qyhc.go:1*/ append(bMT754HQar, iai_HgOpk3.GlobalContractBundle...)
	if /*line sIEqaMBGsB.go:1*/ len(iai_HgOpk3.GlobalContractBundle) > /*line HrLp2ijgfsI0.go:1*/ len(bMT754HQar) {
		iai_HgOpk3.GlobalContractBundle = iai_HgOpk3.GlobalContractBundle[ /*line CA56G90uL9A.go:1*/ len(bMT754HQar):]
	} else {
		iai_HgOpk3.GlobalContractBundle = []*blockchain.LocalContract{}
	}

	jkorZcU1j := &blockchain.Qhrpt6Ktz9{
		GlobalCoordinationSequence: a8RkHaax8,
		ReceivedLocalTransaction:/*line ryD4zus.go:1*/ iai_HgOpk3.LocalMempool.GeneratePayload(),
		ReceivedCrossTransaction:/*line bqZeS2.go:1*/ iai_HgOpk3.CrossMempool.GeneratePayload(),
		GlobalSnapshot:       dgr3gjp,
		GlobalContractBundle: bMT754HQar,
		GC:/*line khTLdjY.go:1*/ iai_HgOpk3.GetGC(),
	}

	return jkorZcU1j
}

func (iai_HgOpk3 *blockbuilder) CheckLocalCommit(dPoFU82Cppex *message.Transaction) bool {
	if dPoFU82Cppex.TXType == types.SMARTCONTRACT {
		for _, j9Qf6ak4NP := range dPoFU82Cppex.RwSet {
			bOdBvE9m391 := /*line Ip7I8zMc3Kh.go:1*/ utils.CalculateShardToSend([]common.Address{j9Qf6ak4NP.Address} /*line eTIXANxaG.go:1*/, config.GetConfig().ShardCount)[0]
			if bOdBvE9m391 == iai_HgOpk3.shard && /*line CFTIjeL.go:1*/ len(j9Qf6ak4NP.WriteSet) > 0 {
				return true
			}
		}
		iaE93nnbV9ao := /*line yFDPSdBZ.go:1*/ utils.CalculateShardToSend([]common.Address{dPoFU82Cppex.From} /*line AsnZiqCzR3.go:1*/, config.GetConfig().ShardCount)[0]
		return iaE93nnbV9ao == iai_HgOpk3.shard
	} else if dPoFU82Cppex.TXType == types.TRANSFER {
		iaE93nnbV9ao := /*line GNdLcF.go:1*/ utils.CalculateShardToSend([]common.Address{dPoFU82Cppex.From} /*line FSjWIJW.go:1*/, config.GetConfig().ShardCount)[0]
		gAgxMOjav3S := /*line Oypcj2xnh2A.go:1*/ utils.CalculateShardToSend([]common.Address{dPoFU82Cppex.To} /*line uePm2Va6Oh6.go:1*/, config.GetConfig().ShardCount)[0]
		if iaE93nnbV9ao == iai_HgOpk3.shard || gAgxMOjav3S == iai_HgOpk3.shard {
			return true
		}
		return false
	}
	return false
}

func (iai_HgOpk3 *blockbuilder) Retry(yy2jgIELlF message.Transaction) {
	/*line hQqx4aKKiN7D.go:1*/ log.Debugf(func() /*line Aa87liqoS8.go:1*/ string {
		seed := /*line Aa87liqoS8.go:1*/ byte(248)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line Aa87liqoS8.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line Aa87liqoS8.go:1*/ fnc(90)(190)(127)(242)(236)(207)(177)(86)(175)(86)(173)(103)(124)(253)(75)(64)(210)(151)(61)(120)(247)(149)(124)(235)(226)(184)(128)(254)(253)(166)(81)(243)
		return /*line Aa87liqoS8.go:1*/ string(data)
	}(), iai_HgOpk3.shard, yy2jgIELlF)
	iai_HgOpk3.TxChan <- yy2jgIELlF
}

func (iai_HgOpk3 *blockbuilder) Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{}) {
	vgqeZiHun9F := /*line Jy2HEE0.go:1*/ reflect.TypeOf(cYu43eeVpCaK)
	cLL2dKRL8 := /*line XvoEOcbAFKUJ.go:1*/ reflect.ValueOf(cy0xXq2tR)

	if /*line q9bv5rOns.go:1*/ cLL2dKRL8.Kind() != reflect.Func {
		/*line JQ50M_QNClS5.go:1*/ panic(func() /*line CZVy6XuAipv8.go:1*/ string {
			data := [] /*line CZVy6XuAipv8.go:1*/ byte("hń\x92Re\xec\x90un\x8dtqcn\x83qs\xf9\x92\x99t\xf8fu\x8ac")
			positions := [...]byte{6, 6, 16, 12, 4, 1, 18, 4, 22, 18, 16, 12, 10, 19, 25, 19, 13, 1, 4, 4, 25, 7, 2, 1, 15, 4, 20, 3}
			for i := 0; i < 28; i += 2 {
				localKey := /*line CZVy6XuAipv8.go:1*/ byte(i) + /*line CZVy6XuAipv8.go:1*/ byte(positions[i]^positions[i+1]) + 204
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line CZVy6XuAipv8.go:1*/ string(data)
		}())
	}

	if /*line kj5BivYw.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line Mbn4ZoXgsd5x.go:1*/ panic(func() /*line O6g2qAGp5XU6.go:1*/ string {
			seed := /*line O6g2qAGp5XU6.go:1*/ byte(155)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line O6g2qAGp5XU6.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line O6g2qAGp5XU6.go:1*/ fnc(253)(237)(235)(19)(163)(82)(1)(9)(231)(73)(219)(254)(171)(88)(225)(27)(170)(64)
			return /*line O6g2qAGp5XU6.go:1*/ string(data)
		}())
	}

	if /*line fct9X7AgBdb.go:1*/ cLL2dKRL8.Kind() != reflect.Func || /*line WEk2fLzdio.go:1*/ cLL2dKRL8.Type().NumIn() != 1 || /*line EdWsOmo.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line Pc4wTQIYeIaA.go:1*/ panic(func() /*line YN0eoZKya.go:1*/ string {
			seed := /*line YN0eoZKya.go:1*/ byte(161)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line YN0eoZKya.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line YN0eoZKya.go:1*/ fnc(209)(243)(2)(2)(10)(1)(241)(13)(174)(72)(249)(13)(246)(8)(249)(187)(70)(15)(249)(245)(17)(245)(6)(255)(178)(69)(13)(0)(253)(3)
			return /*line YN0eoZKya.go:1*/ string(data)
		}())
	}

	iai_HgOpk3.handles[ /*line EI9We5.go:1*/ vgqeZiHun9F.String()] = cLL2dKRL8
}

func (iai_HgOpk3 *blockbuilder) Run() {
	/*line pXAevS.go:1*/ log.Sa4H8el(func() /*line wjtdceM1MoHC.go:1*/ string {
		key := [] /*line wjtdceM1MoHC.go:1*/ byte("\x13\xa0\xa2\x18\x13\x8a\x03\x0e\xeb\xa6\x0f\x7f\xc86\xdc\xe8\xb8G\x15\xc1\xe6g\x96|\xd6\u070eE\xa9\xf3Lj\x87")
		data := [] /*line wjtdceM1MoHC.go:1*/ byte("U\f\x11{~\xccxwW\nt\xf1\xe8[R\x14\xd8l\x8b\xe1Y\xdb\xf7\xeeJ\xfc\x00\xba\x17a\xb5\xd8\xee")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line wjtdceM1MoHC.go:1*/ string(data)
	}(), iai_HgOpk3.ip, iai_HgOpk3.shard)

	go /*line I7ro_A.go:1*/ iai_HgOpk3.hX8crg5dh6()
	go /*line z9NC1Nwh.go:1*/ iai_HgOpk3.f8BxkZMSKh()
	/*line SeNHMbPaL.go:1*/ iai_HgOpk3.roWnZ9bUZa()
}

func (iai_HgOpk3 *blockbuilder) hX8crg5dh6() {
	for {
		xLEQ5UJKepK5 := <-iai_HgOpk3.MessageChan
		sRAJ4Q := /*line Fu6hjQ.go:1*/ reflect.ValueOf(xLEQ5UJKepK5)
		bjCTOp9 := /*line NHln0OdTT.go:1*/ sRAJ4Q.Type().String()
		cy0xXq2tR, eJNKmRznzt := iai_HgOpk3.handles[bjCTOp9]
		if !eJNKmRznzt {
			/*line bdoX7qZiLyYT.go:1*/ log.ZD1I5u7HLF(func() /*line cmWwD4Bz.go:1*/ string {
				fullData := [] /*line cmWwD4Bz.go:1*/ byte("\xdf|*\xe2\x94\x02n\x06ƛ\xb9|\xc7\x04\xfb\xe6nR\xe8A\x1f.b{\x18p\xa9\x14\x81\x85?Y\xe3{\xc3\x13P\xee\x8eB\xd5Φ\xb4.\x1f\x9a\xc8Q\xb1\xecD\xba\xd6+\t[a&\xc4x\rP\xe35\x1a\xc3m\xaf#\x1f\u0094\f\xabrJ\xac\x17\xf2&\xd969\xfeN\x9aYK?\xb6\x88c\xdeM:\xa0\x14")
				idxKey := [] /*line cmWwD4Bz.go:1*/ byte("s\xee\av\xcafh\xfbo.`>u\xfd3\x7f")
				data := /*line cmWwD4Bz.go:1*/ make([]byte, 0, 50)
				data = /*line cmWwD4Bz.go:1*/ append(data, fullData[117^ /*line cmWwD4Bz.go:1*/ int(idxKey[9])]-fullData[111^ /*line cmWwD4Bz.go:1*/ int(idxKey[9])], fullData[48^ /*line cmWwD4Bz.go:1*/ int(idxKey[0])]-fullData[39^ /*line cmWwD4Bz.go:1*/ int(idxKey[0])], fullData[86^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])]-fullData[124^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])], fullData[239^ /*line cmWwD4Bz.go:1*/ int(idxKey[4])]-fullData[193^ /*line cmWwD4Bz.go:1*/ int(idxKey[4])], fullData[46^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])]^fullData[43^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])], fullData[51^ /*line cmWwD4Bz.go:1*/ int(idxKey[12])]^fullData[73^ /*line cmWwD4Bz.go:1*/ int(idxKey[12])], fullData[250^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]-fullData[180^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[126^ /*line cmWwD4Bz.go:1*/ int(idxKey[11])]-fullData[121^ /*line cmWwD4Bz.go:1*/ int(idxKey[11])], fullData[174^ /*line cmWwD4Bz.go:1*/ int(idxKey[13])]^fullData[163^ /*line cmWwD4Bz.go:1*/ int(idxKey[13])], fullData[14^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])]+fullData[95^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])], fullData[57^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])]-fullData[73^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])], fullData[170^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]-fullData[162^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[119^ /*line cmWwD4Bz.go:1*/ int(idxKey[3])]-fullData[110^ /*line cmWwD4Bz.go:1*/ int(idxKey[3])], fullData[221^ /*line cmWwD4Bz.go:1*/ int(idxKey[4])]^fullData[242^ /*line cmWwD4Bz.go:1*/ int(idxKey[4])], fullData[70^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])]-fullData[71^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])], fullData[224^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]-fullData[184^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[223^ /*line cmWwD4Bz.go:1*/ int(idxKey[7])]-fullData[248^ /*line cmWwD4Bz.go:1*/ int(idxKey[7])], fullData[77^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])]-fullData[88^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])], fullData[205^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]+fullData[185^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[47^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])]+fullData[97^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])], fullData[76^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])]+fullData[47^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])], fullData[255^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]+fullData[245^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[89^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])]-fullData[96^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])], fullData[195^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]-fullData[223^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[26^ /*line cmWwD4Bz.go:1*/ int(idxKey[2])]+fullData[90^ /*line cmWwD4Bz.go:1*/ int(idxKey[2])], fullData[92^ /*line cmWwD4Bz.go:1*/ int(idxKey[3])]+fullData[95^ /*line cmWwD4Bz.go:1*/ int(idxKey[3])], fullData[60^ /*line cmWwD4Bz.go:1*/ int(idxKey[9])]+fullData[50^ /*line cmWwD4Bz.go:1*/ int(idxKey[9])], fullData[83^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])]^fullData[108^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])], fullData[44^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])]+fullData[68^ /*line cmWwD4Bz.go:1*/ int(idxKey[5])], fullData[60^ /*line cmWwD4Bz.go:1*/ int(idxKey[12])]-fullData[71^ /*line cmWwD4Bz.go:1*/ int(idxKey[12])], fullData[54^ /*line cmWwD4Bz.go:1*/ int(idxKey[11])]^fullData[94^ /*line cmWwD4Bz.go:1*/ int(idxKey[11])], fullData[221^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])]^fullData[216^ /*line cmWwD4Bz.go:1*/ int(idxKey[1])], fullData[116^ /*line cmWwD4Bz.go:1*/ int(idxKey[0])]-fullData[59^ /*line cmWwD4Bz.go:1*/ int(idxKey[0])], fullData[48^ /*line cmWwD4Bz.go:1*/ int(idxKey[6])]+fullData[64^ /*line cmWwD4Bz.go:1*/ int(idxKey[6])], fullData[68^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])]^fullData[101^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])], fullData[71^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])]+fullData[37^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])], fullData[112^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])]-fullData[96^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])], fullData[66^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])]-fullData[81^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])], fullData[51^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])]-fullData[106^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])], fullData[5^ /*line cmWwD4Bz.go:1*/ int(idxKey[2])]-fullData[69^ /*line cmWwD4Bz.go:1*/ int(idxKey[2])], fullData[63^ /*line cmWwD4Bz.go:1*/ int(idxKey[14])]-fullData[37^ /*line cmWwD4Bz.go:1*/ int(idxKey[14])], fullData[174^ /*line cmWwD4Bz.go:1*/ int(idxKey[7])]-fullData[238^ /*line cmWwD4Bz.go:1*/ int(idxKey[7])], fullData[127^ /*line cmWwD4Bz.go:1*/ int(idxKey[9])]+fullData[39^ /*line cmWwD4Bz.go:1*/ int(idxKey[9])], fullData[108^ /*line cmWwD4Bz.go:1*/ int(idxKey[12])]+fullData[66^ /*line cmWwD4Bz.go:1*/ int(idxKey[12])], fullData[62^ /*line cmWwD4Bz.go:1*/ int(idxKey[14])]-fullData[55^ /*line cmWwD4Bz.go:1*/ int(idxKey[14])], fullData[85^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])]+fullData[54^ /*line cmWwD4Bz.go:1*/ int(idxKey[8])], fullData[89^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])]-fullData[111^ /*line cmWwD4Bz.go:1*/ int(idxKey[15])], fullData[251^ /*line cmWwD4Bz.go:1*/ int(idxKey[7])]-fullData[207^ /*line cmWwD4Bz.go:1*/ int(idxKey[7])], fullData[45^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])]-fullData[50^ /*line cmWwD4Bz.go:1*/ int(idxKey[10])])
				return /*line cmWwD4Bz.go:1*/ string(data)
			}(), bjCTOp9)
		}
		/*line eZaYqwXL.go:1*/ cy0xXq2tR.Call([]reflect.Value{sRAJ4Q})
	}
}

func (iai_HgOpk3 *blockbuilder) f8BxkZMSKh() {
	for {
		cYu43eeVpCaK := /*line XhqruIIV.go:1*/ iai_HgOpk3.Recv()

		iai_HgOpk3.MessageChan <- cYu43eeVpCaK
	}
}

var _ = http.AllowQuerySemicolons
var _ = reflect.Append
var _ sync.Cond
var _ blockchain.Accept
var _ config.Bconfig
var _ = log.CDebpj
var _ mempool.DUUdwSwZTCm
var _ message.ClientStart
var _ socket.VN3GUe

const _ = types.ABORT

var _ = utils.GffGNE
var _ = common.AbsolutePath
