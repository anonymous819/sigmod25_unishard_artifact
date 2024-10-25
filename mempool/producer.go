//line :1
package mempool

import (
	message "unishard/message"
)

type Producer struct {
	sGOXOOj *D98fXQa
}

func NewProducer() *Producer {
	return &Producer{
		sGOXOOj: /*line RCTWE0.go:1*/ XQSPvlaDr0M8(),
	}
}

func (hzWfroqU8Nf *Producer) GeneratePayload() []*message.Transaction {
	return /*line mysRYA9sfdO.go:1*/ hzWfroqU8Nf.sGOXOOj.lrQ0VJ(1024)
}

func (hzWfroqU8Nf *Producer) AddTxn(hDdlup4 *message.Transaction) {
	/*line EV2j5LA7.go:1*/ hzWfroqU8Nf.sGOXOOj.jTeJrkZKUz9A(hDdlup4)
}

func (hzWfroqU8Nf *Producer) CollectTxn(hDdlup4 *message.Transaction) {
	/*line Xfn1ulyYTc.go:1*/ hzWfroqU8Nf.sGOXOOj.qUKUFpon0(hDdlup4)
}

func (hzWfroqU8Nf *Producer) TotalReceivedTxNo() int64 {
	return hzWfroqU8Nf.sGOXOOj.zJHJMtxDGNe
}

func (hzWfroqU8Nf *Producer) Size() int {
	return /*line TgAl9C_xd.go:1*/ hzWfroqU8Nf.sGOXOOj.eRan2aoK2()
}

var _ message.ClientStart
