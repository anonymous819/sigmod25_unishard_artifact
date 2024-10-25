//line :1
package message

import (
	"fmt"

	crypto "unishard/crypto"
	types "unishard/types"

	"github.com/ethereum/go-ethereum/common"
)

type ClientStart struct {
	Shard types.Shard
}
type WorkerBuilderRegister struct {
	SenderShard types.Shard
	Address     string
}

type WorkerBuilderListRequest struct {
	UpcIp_ string
}

type WorkerBuilderListResponse struct {
	Builders map[types.Shard]string
}

type WorkerSharedVariableRegisterRequest struct {
	UpcIp_ string
}

type WorkerSharedVariableRegisterResponse struct {
	PI77Zg94     types.Shard
	QxT5nG3X5Gn4 string
}

type ConsensusNodeRegister struct {
	ConsensusNodeID types.NodeID
	IP              string
}

type X9s0iU4 struct {
	IeZCwdnRUn string
}

type TN51p2hVBir struct {
	PKZzTa     common.Hash
	Dfyxn2pdX  string
	F2x1L86P   string
	SojSbIsHsW []byte
}

type Bncgq6Ys struct {
	W4puDs      common.Hash
	DESc6onmvvK int64
}

type SBXZ5rXiXHpc struct {
	JR5HajI  common.Hash
	UCK4axEV []int64
}
type Experiment struct {
	types.Shard
	Da8uC0Sys2u  R8SUVOg
	Bkq028NruW   []*Bncgq6Ys
	RoQbttkP     float64
	JUgObcOr     []*SBXZ5rXiXHpc
	W2g99STnJRA1 []*SBXZ5rXiXHpc
	E2v01O_Ii    bool
}

type R8SUVOg struct {
	RRTdrUMYO int
	Zjb9_abe  int
	GV9IpLO9e int
	XRLlpce9L float64
}

type CrossShardTransactionLatencyDissection struct {
	WBBWaitingTime            int64
	WorkerConsensusTime       int64
	Network1                  int64
	CBBWaitingTIme            int64
	CoordinationConsensusTime int64
	Network2                  int64

	WBBWaitingTime2      int64
	WorkerExecuteStart   int64
	WorkerExecuteEnd     int64
	WorkerConsensusTime2 int64
}

type RwVariable struct {
	Address common.Address
	Name    string
	types.J_zoDMw
}

type TransactionForm struct {
	From                common.Address
	To                  common.Address
	Value               int
	Data                []byte
	ExternalAddressList []common.Address
	MappingIdx          []byte
	Timestamp           int64
}

type Transaction struct {
	NodeID types.NodeID
	Hash   common.Hash

	TXType types.MaZWZHI

	From                common.Address
	To                  common.Address
	Value               int
	Data                []byte
	ExternalAddressList []common.Address
	RwSet               []RwVariable
	IsCrossShardTx      bool
	Timestamp           int64
	LatencyDissection   CrossShardTransactionLatencyDissection
}

func UjchOXDlIcT(qlfQnm TransactionForm, aM5X1f9 bool, zxQSw84qK types.NodeID) *Transaction {
	uQalDtAfP := /*line rqaAPL0fFF9S.go:1*/ new(Transaction)

	uQalDtAfP.NodeID = zxQSw84qK

	uQalDtAfP.From = qlfQnm.From
	uQalDtAfP.To = qlfQnm.To
	uQalDtAfP.Value = qlfQnm.Value
	uQalDtAfP.IsCrossShardTx = aM5X1f9
	uQalDtAfP.Timestamp = qlfQnm.Timestamp
	uQalDtAfP.TXType = types.TRANSFER
	/*line MmOErQXC.go:1*/ uQalDtAfP.caObMR()
	return uQalDtAfP
}

func I9j76PaEYB(qlfQnm TransactionForm, aM5X1f9 bool, zxQSw84qK types.NodeID) *Transaction {
	uQalDtAfP := /*line a2J5ec.go:1*/ new(Transaction)
	uQalDtAfP.NodeID = zxQSw84qK
	uQalDtAfP.From = qlfQnm.From
	uQalDtAfP.To = qlfQnm.To
	uQalDtAfP.Value = qlfQnm.Value
	uQalDtAfP.Data = qlfQnm.Data
	uQalDtAfP.IsCrossShardTx = aM5X1f9
	uQalDtAfP.Timestamp = qlfQnm.Timestamp
	uQalDtAfP.TXType = types.DEPLOY
	/*line K6A2TLdY.go:1*/ uQalDtAfP.caObMR()
	return uQalDtAfP
}

func AtmQ9HD(qlfQnm TransactionForm, udzyukn []RwVariable, aM5X1f9 bool, zxQSw84qK types.NodeID) *Transaction {
	uQalDtAfP := /*line BEbejU5S.go:1*/ new(Transaction)
	uQalDtAfP.NodeID = zxQSw84qK
	uQalDtAfP.From = qlfQnm.From
	uQalDtAfP.To = qlfQnm.To
	uQalDtAfP.Value = qlfQnm.Value
	uQalDtAfP.Data = qlfQnm.Data
	uQalDtAfP.ExternalAddressList = qlfQnm.ExternalAddressList
	uQalDtAfP.RwSet = udzyukn
	uQalDtAfP.IsCrossShardTx = aM5X1f9
	uQalDtAfP.Timestamp = qlfQnm.Timestamp
	uQalDtAfP.TXType = types.SMARTCONTRACT
	/*line dmaKVBLGN.go:1*/ uQalDtAfP.caObMR()
	return uQalDtAfP
}

type rawTransaction struct {
	From      common.Address
	To        common.Address
	Value     int
	Data      []byte
	TXType    types.MaZWZHI
	Timestamp int64
}

func (uQalDtAfP *Transaction) caObMR() {
	fl4suoJ := &rawTransaction{
		From:      uQalDtAfP.From,
		To:        uQalDtAfP.To,
		Value:     uQalDtAfP.Value,
		Data:      uQalDtAfP.Data,
		TXType:    uQalDtAfP.TXType,
		Timestamp: uQalDtAfP.Timestamp,
	}

	uQalDtAfP.Hash = /*line MuQFNs.go:1*/ crypto.ZsUswPaGG4Z(fl4suoJ)
}

func (uQalDtAfP *Transaction) String() string {

	var arB5fS5b string
	switch uQalDtAfP.TXType {
	case types.TRANSFER:
		arB5fS5b = /*line Ba3OodaYh.go:1*/ fmt.Sprintf(func() /*line fuzjoC.go:1*/ string {
			data := /*line fuzjoC.go:1*/ make([]byte, 0, 88)
			i := 34
			decryptKey := 103
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 30:
					i = 21
					data = /*line fuzjoC.go:1*/ append(data, "\xb6\xdb\x06\x1a"...,
					)
				case 34:
					i = 13
					data = /*line fuzjoC.go:1*/ append(data, 223)
				case 14:
					data = /*line fuzjoC.go:1*/ append(data, "\xa3\xd6\xf3"...,
					)
					i = 0
				case 20:
					data = /*line fuzjoC.go:1*/ append(data, 73)
					i = 25
				case 24:
					data = /*line fuzjoC.go:1*/ append(data, "\xfe\x0f\xfb\xcc"...,
					)
					i = 19
				case 25:
					i = 17
					data = /*line fuzjoC.go:1*/ append(data, "\x12\xff"...,
					)
				case 31:
					i = 6
					data = /*line fuzjoC.go:1*/ append(data, "\xee\xf2B"...,
					)
				case 10:
					i = 29
					data = /*line fuzjoC.go:1*/ append(data, "\xe7\x0f\x03\b"...,
					)
				case 8:
					i = 18
					data = /*line fuzjoC.go:1*/ append(data, "\x15)4"...,
					)
				case 12:
					i = 22
					data = /*line fuzjoC.go:1*/ append(data, 245)
				case 9:
					data = /*line fuzjoC.go:1*/ append(data, "\xd4$\xd9\xcc"...,
					)
					i = 33
				case 26:
					data = /*line fuzjoC.go:1*/ append(data, 9)
					i = 31
				case 33:
					i = 5
					data = /*line fuzjoC.go:1*/ append(data, "\xf7\x11\xdb"...,
					)
				case 18:
					data = /*line fuzjoC.go:1*/ append(data, "+88<"...,
					)
					i = 16
				case 23:
					i = 7
					data = /*line fuzjoC.go:1*/ append(data, "\xd7'"...,
					)
				case 27:
					i = 15
					data = /*line fuzjoC.go:1*/ append(data, "\x19-!\""...,
					)
				case 19:
					i = 30
					data = /*line fuzjoC.go:1*/ append(data, "\xb1\xb5\x0f"...,
					)
				case 3:
					i = 4
					data = /*line fuzjoC.go:1*/ append(data, "\xff(\xf7"...,
					)
				case 21:
					i = 9
					data = /*line fuzjoC.go:1*/ append(data, "\x17\xe3\xc8"...,
					)
				case 16:
					i = 20
					data = /*line fuzjoC.go:1*/ append(data, 71)
				case 4:
					data = /*line fuzjoC.go:1*/ append(data, "=9<;"...,
					)
					i = 26
				case 7:
					data = /*line fuzjoC.go:1*/ append(data, "\xdc\xd7"...,
					)
					i = 3
				case 0:
					data = /*line fuzjoC.go:1*/ append(data, 225)
					i = 12
				case 6:
					i = 8
					data = /*line fuzjoC.go:1*/ append(data, "\xef\xe2"...,
					)
				case 5:
					data = /*line fuzjoC.go:1*/ append(data, "\xc0\xcc\x1c\xd1"...,
					)
					i = 32
				case 17:
					i = 28
					data = /*line fuzjoC.go:1*/ append(data, "\x03S"...,
					)
				case 13:
					data = /*line fuzjoC.go:1*/ append(data, "\xfc\xea\xf6\x02"...,
					)
					i = 2
				case 15:
					data = /*line fuzjoC.go:1*/ append(data, "\xf6\xd3"...,
					)
					i = 23
				case 11:
					data = /*line fuzjoC.go:1*/ append(data, 37)
					i = 27
				case 32:
					data = /*line fuzjoC.go:1*/ append(data, "\xc4\xfd\x1b"...,
					)
					i = 11
				case 28:
					for y := range data {
						data[y] = data[y] - /*line fuzjoC.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 22:
					data = /*line fuzjoC.go:1*/ append(data, "\xf9\xe6"...,
					)
					i = 10
				case 29:
					i = 24
					data = /*line fuzjoC.go:1*/ append(data, "\x06\xbf\xe6"...,
					)
				case 2:
					data = /*line fuzjoC.go:1*/ append(data, "\xf4\xf2\xfe"...,
					)
					i = 14
				}
			}
			return /*line fuzjoC.go:1*/ string(data)
		}(), uQalDtAfP.Hash, uQalDtAfP.From, uQalDtAfP.To, uQalDtAfP.Value, uQalDtAfP.IsCrossShardTx, uQalDtAfP.Timestamp)

	case types.SMARTCONTRACT:
		arB5fS5b = /*line k1Qppefvk9Wp.go:1*/ fmt.Sprintf(func() /*line PNHIkz.go:1*/ string {
			key := [] /*line PNHIkz.go:1*/ byte("m\xa2\x99\xa7\xa2a\xea\xab\xc8\xc2J\f[6\x83s\x8b\x03\xf3\xe1\x97\xf5V1H\xa5)\xde%\x81\x90\xd3\xecjӨ\u1cc8\x1fYq\x93\x11\x1a\xbc\xe3\xbc\x13|Ѯ\xe3F3zT\x1bQ\xcf\x11\xbd.\xf1\xbc)W\x90\x99\x95f \xdc\x02\xd5A\fI\xc6\x1c\xbfR")
			data := [] /*line PNHIkz.go:1*/ byte("\xee\x83ݶ~\U00083daa\xb2\xf9c\x13>\xef\xee\xd8q-s\xdbl\x18B\x19\xbeK\x8bJ\xed\x90uu\t\x95\x92?r\xf0\x01\xfb\xfe\xa7\x0f\v\xbaId1壳W\xda\xf2\xfc\xd8\x05\x01\xa8B\xa8FId\xfc\x1f\x9c\x87\xbf\x03M\x89q\x9f a't\x04f$")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line PNHIkz.go:1*/ string(data)
		}(), uQalDtAfP.NodeID, uQalDtAfP.Hash, uQalDtAfP.To, uQalDtAfP.Data, uQalDtAfP.RwSet, uQalDtAfP.Timestamp)

	default:
		arB5fS5b = func() /*line aSa1fEWaY_ro.go:1*/ string {
			seed := /*line aSa1fEWaY_ro.go:1*/ byte(30)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line aSa1fEWaY_ro.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line aSa1fEWaY_ro.go:1*/ fnc(48)(33)(5)(172)(51)(34)(251)(0)(255)(3)(2)(241)(255)(188)(52)(30)(239)(13)(5)(238)(2)(17)(245)(6)(255)
			return /*line aSa1fEWaY_ro.go:1*/ string(data)
		}()
	}
	return arB5fS5b
}

type Query struct {
	VDVxPqDv chan QueryReply
}

func (hwP2PpHahGr *Query) Reply(d8Y9UnsZ8 QueryReply) {
	hwP2PpHahGr.VDVxPqDv <- d8Y9UnsZ8
}

type QueryReply struct {
	Info string
}

type RequestLeader struct {
	H1yn44Ukl chan RequestLeaderReply
}

func (hwP2PpHahGr *RequestLeader) Reply(d8Y9UnsZ8 RequestLeaderReply) {
	hwP2PpHahGr.H1yn44Ukl <- d8Y9UnsZ8
}

type RequestLeaderReply struct {
	Leader string
}

type ReportByzantine struct {
	IEmX_C chan LZ6CwlA
}

func (hwP2PpHahGr *ReportByzantine) Reply(d8Y9UnsZ8 LZ6CwlA) {
	hwP2PpHahGr.IEmX_C <- d8Y9UnsZ8
}

type LZ6CwlA struct {
	Leader string
}

var _ = fmt.Append
var _ crypto.Pp__49cd

const _ = types.ABORT

var _ = common.AbsolutePath
