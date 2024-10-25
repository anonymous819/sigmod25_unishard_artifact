//line :1
package types

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (* /*line Ma1nvllpLt.go:1*/ aBHOK_3Irh)(nil)

func (mwH7XS Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address     common.Address `json:"address" gencodec:"required"`
		Topics      []common.Hash  `json:"topics" gencodec:"required"`
		Data        hexutil.Bytes  `json:"data" gencodec:"required"`
		BlockNumber hexutil.Uint64 `json:"blockNumber" rlp:"-"`
		TxHash      common.Hash    `json:"transactionHash" gencodec:"required" rlp:"-"`
		TxIndex     hexutil.Uint   `json:"transactionIndex" rlp:"-"`
		BlockHash   common.Hash    `json:"blockHash" rlp:"-"`
		Index       hexutil.Uint   `json:"logIndex" rlp:"-"`
		Removed     bool           `json:"removed" rlp:"-"`
	}
	var aINiWZ_Jtlzj Log
	aINiWZ_Jtlzj.Address = mwH7XS.Address
	aINiWZ_Jtlzj.Topics = mwH7XS.Topics
	aINiWZ_Jtlzj.Data = mwH7XS.Data
	aINiWZ_Jtlzj.BlockNumber = /*line Uj7T_Q.go:1*/ hexutil.Uint64(mwH7XS.BlockNumber)
	aINiWZ_Jtlzj.TxHash = mwH7XS.TxHash
	aINiWZ_Jtlzj.TxIndex = /*line Cznj2vn.go:1*/ hexutil.Uint(mwH7XS.TxIndex)
	aINiWZ_Jtlzj.BlockHash = mwH7XS.BlockHash
	aINiWZ_Jtlzj.Index = /*line RdD6pd1n3Idw.go:1*/ hexutil.Uint(mwH7XS.Index)
	aINiWZ_Jtlzj.Removed = mwH7XS.Removed
	return /*line SUZivu3K.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (mwH7XS *Log) UnmarshalJSON(uzD2Up []byte) error {
	type Log struct {
		Address     *common.Address `json:"address" gencodec:"required"`
		Topics      []common.Hash   `json:"topics" gencodec:"required"`
		Data        *hexutil.Bytes  `json:"data" gencodec:"required"`
		BlockNumber *hexutil.Uint64 `json:"blockNumber" rlp:"-"`
		TxHash      *common.Hash    `json:"transactionHash" gencodec:"required" rlp:"-"`
		TxIndex     *hexutil.Uint   `json:"transactionIndex" rlp:"-"`
		BlockHash   *common.Hash    `json:"blockHash" rlp:"-"`
		Index       *hexutil.Uint   `json:"logIndex" rlp:"-"`
		Removed     *bool           `json:"removed" rlp:"-"`
	}
	var otqLrvn1CD Log
	if rfteMJju := /*line GEaEzQk5YjjH.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD); rfteMJju != nil {
		return rfteMJju
	}
	if otqLrvn1CD.Address == nil {
		return /*line rF7eGPXZJ9.go:1*/ errors.New(func() /*line IajollypaRN.go:1*/ string {
			data := /*line IajollypaRN.go:1*/ make([]byte, 0, 41)
			i := 3
			decryptKey := 98
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					data = /*line IajollypaRN.go:1*/ append(data, "\x8e\x99\x9a\x89"...,
					)
					i = 1
				case 5:
					data = /*line IajollypaRN.go:1*/ append(data, 156)
					i = 8
				case 14:
					data = /*line IajollypaRN.go:1*/ append(data, "\x9e\x92"...,
					)
					i = 11
				case 16:
					i = 2
					for y := range data {
						data[y] = data[y] - /*line IajollypaRN.go:1*/ byte(decryptKey^y)
					}
				case 15:
					i = 17
					data = /*line IajollypaRN.go:1*/ append(data, "ku"...,
					)
				case 3:
					i = 10
					data = /*line IajollypaRN.go:1*/ append(data, 145)
				case 1:
					data = /*line IajollypaRN.go:1*/ append(data, 143)
					i = 9
				case 8:
					data = /*line IajollypaRN.go:1*/ append(data, "\x9c\x95R"...,
					)
					i = 12
				case 9:
					i = 0
					data = /*line IajollypaRN.go:1*/ append(data, 137)
				case 6:
					data = /*line IajollypaRN.go:1*/ append(data, 159)
					i = 5
				case 12:
					i = 7
					data = /*line IajollypaRN.go:1*/ append(data, "Z\x9d\xa1\xa2"...,
					)
				case 13:
					data = /*line IajollypaRN.go:1*/ append(data, "\x8f\x8fT\x9b"...,
					)
					i = 6
				case 4:
					i = 16
					data = /*line IajollypaRN.go:1*/ append(data, "Mqj"...,
					)
				case 18:
					data = /*line IajollypaRN.go:1*/ append(data, "\xac\xadb$"...,
					)
					i = 15
				case 7:
					data = /*line IajollypaRN.go:1*/ append(data, "\xb1\x9d"...,
					)
					i = 18
				case 11:
					data = /*line IajollypaRN.go:1*/ append(data, "\x9f\xa4\x91\x9b"...,
					)
					i = 13
				case 17:
					data = /*line IajollypaRN.go:1*/ append(data, "y "...,
					)
					i = 4
				case 0:
					i = 14
					data = /*line IajollypaRN.go:1*/ append(data, 67)
				}
			}
			return /*line IajollypaRN.go:1*/ string(data)
		}())
	}
	mwH7XS.Address = *otqLrvn1CD.Address
	if otqLrvn1CD.Topics == nil {
		return /*line mBhD6HboI.go:1*/ errors.New(func() /*line HAkd9s1Mq2Z.go:1*/ string {
			seed := /*line HAkd9s1Mq2Z.go:1*/ byte(194)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line HAkd9s1Mq2Z.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line HAkd9s1Mq2Z.go:1*/ fnc(47)(90)(190)(124)(238)(225)(187)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(187)(113)(227)(191)(120)(0)(180)(97)(8)(25)(53)(24)(92)(219)(174)
			return /*line HAkd9s1Mq2Z.go:1*/ string(data)
		}())
	}
	mwH7XS.Topics = otqLrvn1CD.Topics
	if otqLrvn1CD.Data == nil {
		return /*line akoe3tO.go:1*/ errors.New(func() /*line ayyF_4q4Q.go:1*/ string {
			data := /*line ayyF_4q4Q.go:1*/ make([]byte, 0, 38)
			i := 4
			decryptKey := 159
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					data = /*line ayyF_4q4Q.go:1*/ append(data, "1\xeb"...,
					)
					i = 11
				case 4:
					i = 10
					data = /*line ayyF_4q4Q.go:1*/ append(data, "96AB"...,
					)
				case 11:
					data = /*line ayyF_4q4Q.go:1*/ append(data, "6*"...,
					)
					i = 7
				case 7:
					i = 20
					data = /*line ayyF_4q4Q.go:1*/ append(data, 55)
				case 17:
					i = 15
					data = /*line ayyF_4q4Q.go:1*/ append(data, 94)
				case 3:
					i = 8
					data = /*line ayyF_4q4Q.go:1*/ append(data, "'\xfcC"...,
					)
				case 14:
					for y := range data {
						data[y] = data[y] - /*line ayyF_4q4Q.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 5:
					data = /*line ayyF_4q4Q.go:1*/ append(data, 79)
					i = 14
				case 19:
					i = 9
					data = /*line ayyF_4q4Q.go:1*/ append(data, "\xfa\x02"...,
					)
				case 1:
					data = /*line ayyF_4q4Q.go:1*/ append(data, ":^"...,
					)
					i = 5
				case 9:
					i = 12
					data = /*line ayyF_4q4Q.go:1*/ append(data, 56)
				case 20:
					i = 13
					data = /*line ayyF_4q4Q.go:1*/ append(data, "<)3"...,
					)
				case 16:
					i = 17
					data = /*line ayyF_4q4Q.go:1*/ append(data, "8B"...,
					)
				case 12:
					data = /*line ayyF_4q4Q.go:1*/ append(data, "6J"...,
					)
					i = 18
				case 10:
					i = 0
					data = /*line ayyF_4q4Q.go:1*/ append(data, "17"...,
					)
				case 2:
					data = /*line ayyF_4q4Q.go:1*/ append(data, 61)
					i = 19
				case 18:
					data = /*line ayyF_4q4Q.go:1*/ append(data, "8\xf7\xf1"...,
					)
					i = 16
				case 13:
					i = 3
					data = /*line ayyF_4q4Q.go:1*/ append(data, 39)
				case 8:
					data = /*line ayyF_4q4Q.go:1*/ append(data, "GDD"...,
					)
					i = 2
				case 15:
					data = /*line ayyF_4q4Q.go:1*/ append(data, 13)
					i = 1
				}
			}
			return /*line ayyF_4q4Q.go:1*/ string(data)
		}())
	}
	mwH7XS.Data = *otqLrvn1CD.Data
	if otqLrvn1CD.BlockNumber != nil {
		mwH7XS.BlockNumber = /*line PNvV8gPOB.go:1*/ uint64(*otqLrvn1CD.BlockNumber)
	}
	if otqLrvn1CD.TxHash == nil {
		return /*line XDaUZQacP.go:1*/ errors.New(func() /*line oCh3bJfGR8.go:1*/ string {
			data := /*line oCh3bJfGR8.go:1*/ make([]byte, 0, 49)
			i := 13
			decryptKey := 240
			for counter := 0; i != 20; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 11:
					data = /*line oCh3bJfGR8.go:1*/ append(data, "+-$\xe3"...,
					)
					i = 4
				case 9:
					i = 20
					for y := range data {
						data[y] = data[y] - /*line oCh3bJfGR8.go:1*/ byte(decryptKey^y)
					}
				case 7:
					i = 5
					data = /*line oCh3bJfGR8.go:1*/ append(data, ">R"...,
					)
				case 4:
					data = /*line oCh3bJfGR8.go:1*/ append(data, "\xe9A"...,
					)
					i = 17
				case 2:
					i = 7
					data = /*line oCh3bJfGR8.go:1*/ append(data, "^ce"...,
					)
				case 17:
					data = /*line oCh3bJfGR8.go:1*/ append(data, ">0"...,
					)
					i = 1
				case 14:
					data = /*line oCh3bJfGR8.go:1*/ append(data, "SBJ"...,
					)
					i = 6
				case 6:
					i = 3
					data = /*line oCh3bJfGR8.go:1*/ append(data, "@>"...,
					)
				case 3:
					i = 11
					data = /*line oCh3bJfGR8.go:1*/ append(data, "\xe5*0"...,
					)
				case 1:
					i = 18
					data = /*line oCh3bJfGR8.go:1*/ append(data, "<<"...,
					)
				case 0:
					data = /*line oCh3bJfGR8.go:1*/ append(data, 68)
					i = 12
				case 18:
					data = /*line oCh3bJfGR8.go:1*/ append(data, ").>"...,
					)
					i = 2
				case 15:
					i = 10
					data = /*line oCh3bJfGR8.go:1*/ append(data, 98)
				case 19:
					data = /*line oCh3bJfGR8.go:1*/ append(data, ":\xf2O"...,
					)
					i = 16
				case 13:
					i = 8
					data = /*line oCh3bJfGR8.go:1*/ append(data, "B="...,
					)
				case 16:
					data = /*line oCh3bJfGR8.go:1*/ append(data, "AP"...,
					)
					i = 14
				case 5:
					data = /*line oCh3bJfGR8.go:1*/ append(data, "c[\x19\x1d"...,
					)
					i = 15
				case 10:
					i = 0
					data = /*line oCh3bJfGR8.go:1*/ append(data, "np\x19"...,
					)
				case 12:
					i = 9
					data = /*line oCh3bJfGR8.go:1*/ append(data, "ja"...,
					)
				case 8:
					data = /*line oCh3bJfGR8.go:1*/ append(data, "JI:>"...,
					)
					i = 19
				}
			}
			return /*line oCh3bJfGR8.go:1*/ string(data)
		}())
	}
	mwH7XS.TxHash = *otqLrvn1CD.TxHash
	if otqLrvn1CD.TxIndex != nil {
		mwH7XS.TxIndex = /*line mgr6lGWxG9q.go:1*/ uint(*otqLrvn1CD.TxIndex)
	}
	if otqLrvn1CD.BlockHash != nil {
		mwH7XS.BlockHash = *otqLrvn1CD.BlockHash
	}
	if otqLrvn1CD.Index != nil {
		mwH7XS.Index = /*line bT0ATNj6NHpz.go:1*/ uint(*otqLrvn1CD.Index)
	}
	if otqLrvn1CD.Removed != nil {
		mwH7XS.Removed = *otqLrvn1CD.Removed
	}
	return nil
}

var _ = json.Compact
var _ = errors.As
var _ = common.AbsolutePath
var _ hexutil.Big
