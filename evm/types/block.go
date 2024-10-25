//line :1
package types

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

type BlockNonce [8]byte

func DXY1_SFDpeE(ibcOzM6f uint64) BlockNonce {
	var hqUDDp5MZm6a BlockNonce
	/*line kV8_sp.go:1*/ binary.BigEndian.PutUint64(hqUDDp5MZm6a[:], ibcOzM6f)
	return hqUDDp5MZm6a
}

func (hqUDDp5MZm6a BlockNonce) Uint64() uint64 {
	return /*line Bunkx2a.go:1*/ binary.BigEndian.Uint64(hqUDDp5MZm6a[:])
}

func (hqUDDp5MZm6a BlockNonce) MarshalText() ([]byte, error) {
	return /*line _62hbyb.go:1*/ hexutil.Bytes(hqUDDp5MZm6a[:]).MarshalText()
}

func (hqUDDp5MZm6a *BlockNonce) UnmarshalText(uzD2Up []byte) error {
	return /*line zNuLCTiockrp.go:1*/ hexutil.UnmarshalFixedText(func() /*line QYzyssVux.go:1*/ string {
		key := [] /*line QYzyssVux.go:1*/ byte("\x1dv\x91=\xf5[LC\xd4V")
		data := [] /*line QYzyssVux.go:1*/ byte("_\xe2\x00\xa0`\xa9\xbb\xb17\xbb")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line QYzyssVux.go:1*/ string(data)
	}(), uzD2Up, hqUDDp5MZm6a[:])
}

//go:generate go run github.com/fjl/gencodec -type Header -field-override headerMarshaling -out gen_header_json.go
//go:generate go run ../../rlp/rlpgen -type Header -out gen_header_rlp.go

type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`

	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`
}

type dCCagad0IRZF struct {
	IObokkp      *hexutil.Big
	XsxJTK9PQ    *hexutil.Big
	He9kWVPVz7   hexutil.Uint64
	XRcxeYOK     hexutil.Uint64
	WAlHbp0      hexutil.Uint64
	RIkWsDOET5k  hexutil.Bytes
	JBht7167XLt  *hexutil.Big
	AwEC86JVc2Fm common.Hash `json:"hash"`
	Tz9tjHuAdd   *hexutil.Uint64
	RFTHgc36fvip *hexutil.Uint64
}

func (hP0dFanm4 *Header) Hash() common.Hash {
	return /*line RzUWe6gc.go:1*/ uWZWEzDAB(hP0dFanm4)
}

var f6NhdF37pYe = /*line UEmC1R.go:1*/ common.StorageSize( /*line xa4Kga2J8g.go:1*/ reflect.TypeOf(Header{}).Size())

func (hP0dFanm4 *Header) Size() common.StorageSize {
	var ro9Fb3 int
	if hP0dFanm4.BaseFee != nil {
		ro9Fb3 = /*line taBcpOzaE.go:1*/ hP0dFanm4.BaseFee.BitLen()
	}
	return f6NhdF37pYe + /*line dwQppK3n0PR.go:1*/ common.StorageSize( /*line Yq_rG07W3lLR.go:1*/ len(hP0dFanm4.Extra)+( /*line HZq5AXEygjG.go:1*/ hP0dFanm4.Difficulty.BitLen()+ /*line Ra8g74aBuzdf.go:1*/ hP0dFanm4.Number.BitLen()+ro9Fb3)/8)
}

func (hP0dFanm4 *Header) SanityCheck() error {
	if hP0dFanm4.Number != nil && ! /*line hlUQvWon7.go:1*/ hP0dFanm4.Number.IsUint64() {
		return /*line b7Q50X_.go:1*/ fmt.Errorf(func() /*line XiFtwl.go:1*/ string {
			data := /*line XiFtwl.go:1*/ make([]byte, 0, 34)
			i := 11
			decryptKey := 208
			for counter := 0; i != 13; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					data = /*line XiFtwl.go:1*/ append(data, "\xba\xa1\xaf\xaf"...,
					)
					i = 17
				case 16:
					i = 7
					data = /*line XiFtwl.go:1*/ append(data, "\xad\xe0"...,
					)
				case 17:
					i = 1
					data = /*line XiFtwl.go:1*/ append(data, "\xb9\xf2\xe9"...,
					)
				case 6:
					i = 14
					data = /*line XiFtwl.go:1*/ append(data, "\xb0\xb3\xfd"...,
					)
				case 10:
					i = 15
					data = /*line XiFtwl.go:1*/ append(data, 182)
				case 12:
					i = 16
					data = /*line XiFtwl.go:1*/ append(data, 167)
				case 3:
					data = /*line XiFtwl.go:1*/ append(data, "\xf1\xa0"...,
					)
					i = 8
				case 15:
					i = 3
					data = /*line XiFtwl.go:1*/ append(data, "\xb9\xbd\xb0\xbb"...,
					)
				case 11:
					i = 6
					data = /*line XiFtwl.go:1*/ append(data, 170)
				case 7:
					data = /*line XiFtwl.go:1*/ append(data, "\xe4\x9a"...,
					)
					i = 9
				case 9:
					for y := range data {
						data[y] = data[y] ^ /*line XiFtwl.go:1*/ byte(decryptKey^y)
					}
					i = 13
				case 1:
					i = 5
					data = /*line XiFtwl.go:1*/ append(data, 164)
				case 5:
					i = 2
					data = /*line XiFtwl.go:1*/ append(data, "\xae\xb0"...,
					)
				case 0:
					i = 4
					data = /*line XiFtwl.go:1*/ append(data, 170)
				case 14:
					data = /*line XiFtwl.go:1*/ append(data, "\xb6\xba"...,
					)
					i = 0
				case 2:
					data = /*line XiFtwl.go:1*/ append(data, 169)
					i = 12
				case 4:
					data = /*line XiFtwl.go:1*/ append(data, "\xbe\xb3\xf7"...,
					)
					i = 10
				}
			}
			return /*line XiFtwl.go:1*/ string(data)
		}(), /*line ipkQzbpyC4Q.go:1*/ hP0dFanm4.Number.BitLen())
	}
	if hP0dFanm4.Difficulty != nil {
		if jA0bwDc3B := /*line ldBKRz1m.go:1*/ hP0dFanm4.Difficulty.BitLen(); jA0bwDc3B > 80 {
			return /*line gWDkXr7.go:1*/ fmt.Errorf(func() /*line HaOCYM.go:1*/ string {
				data := [] /*line HaOCYM.go:1*/ byte("t\xaf\xfe\xf0\x8d*\xfb\x06\x01p\x9bl\xad3\xe1Ѭ\xd8f\"i\xc3\xe1l\x062l b\xad4\x06\xec2\xfc\xd4\xd7")
				positions := [...]byte{12, 34, 24, 33, 24, 13, 1, 5, 25, 19, 32, 30, 10, 24, 1, 13, 30, 19, 17, 26, 33, 8, 5, 30, 8, 2, 29, 9, 3, 16, 36, 13, 10, 22, 24, 31, 14, 17, 7, 6, 21, 32, 33, 15, 35, 4, 19, 14, 19, 29}
				for i := 0; i < 50; i += 2 {
					localKey := /*line HaOCYM.go:1*/ byte(i) + /*line HaOCYM.go:1*/ byte(positions[i]^positions[i+1]) + 69
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line HaOCYM.go:1*/ string(data)
			}(), jA0bwDc3B)
		}
	}
	if oBVxZOkL := /*line xOaH81.go:1*/ len(hP0dFanm4.Extra); oBVxZOkL > 100*1024 {
		return /*line Fnz0oBT.go:1*/ fmt.Errorf(func() /*line XOQFbSla.go:1*/ string {
			fullData := [] /*line XOQFbSla.go:1*/ byte("\xdc-\x98t\x83\xd8h\xf32\x94\x8b\xd0ܼܣ#HI\xc2^\xeeI\xcfJ\xe3\xaaL\xbf\x9d\x94\x12\xb2\x9dl\xbd\xd7\xf7jjF2lh)\x14\xa3'\xed\xe3\xc6\xe9\xa7\xef\x9c\x05\x98L\xa0\xc6\x02M\x06=\xc1\x8f\xb1\xea")
			idxKey := [] /*line XOQFbSla.go:1*/ byte("y\xc5\xfe5\xf91\xa9\x92\xee\x19\x9d\x92l")
			data := /*line XOQFbSla.go:1*/ make([]byte, 0, 35)
			data = /*line XOQFbSla.go:1*/ append(data, fullData[240^ /*line XOQFbSla.go:1*/ int(idxKey[2])]-fullData[248^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[229^ /*line XOQFbSla.go:1*/ int(idxKey[2])]^fullData[238^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[55^ /*line XOQFbSla.go:1*/ int(idxKey[3])]^fullData[16^ /*line XOQFbSla.go:1*/ int(idxKey[3])], fullData[238^ /*line XOQFbSla.go:1*/ int(idxKey[1])]^fullData[212^ /*line XOQFbSla.go:1*/ int(idxKey[1])], fullData[37^ /*line XOQFbSla.go:1*/ int(idxKey[9])]+fullData[62^ /*line XOQFbSla.go:1*/ int(idxKey[9])], fullData[175^ /*line XOQFbSla.go:1*/ int(idxKey[7])]+fullData[191^ /*line XOQFbSla.go:1*/ int(idxKey[7])], fullData[167^ /*line XOQFbSla.go:1*/ int(idxKey[11])]^fullData[143^ /*line XOQFbSla.go:1*/ int(idxKey[11])], fullData[31^ /*line XOQFbSla.go:1*/ int(idxKey[3])]-fullData[2^ /*line XOQFbSla.go:1*/ int(idxKey[3])], fullData[136^ /*line XOQFbSla.go:1*/ int(idxKey[10])]^fullData[151^ /*line XOQFbSla.go:1*/ int(idxKey[10])], fullData[58^ /*line XOQFbSla.go:1*/ int(idxKey[9])]^fullData[56^ /*line XOQFbSla.go:1*/ int(idxKey[9])], fullData[118^ /*line XOQFbSla.go:1*/ int(idxKey[0])]+fullData[101^ /*line XOQFbSla.go:1*/ int(idxKey[0])], fullData[105^ /*line XOQFbSla.go:1*/ int(idxKey[12])]+fullData[101^ /*line XOQFbSla.go:1*/ int(idxKey[12])], fullData[164^ /*line XOQFbSla.go:1*/ int(idxKey[11])]-fullData[147^ /*line XOQFbSla.go:1*/ int(idxKey[11])], fullData[25^ /*line XOQFbSla.go:1*/ int(idxKey[5])]-fullData[40^ /*line XOQFbSla.go:1*/ int(idxKey[5])], fullData[72^ /*line XOQFbSla.go:1*/ int(idxKey[12])]^fullData[97^ /*line XOQFbSla.go:1*/ int(idxKey[12])], fullData[70^ /*line XOQFbSla.go:1*/ int(idxKey[0])]+fullData[72^ /*line XOQFbSla.go:1*/ int(idxKey[0])], fullData[188^ /*line XOQFbSla.go:1*/ int(idxKey[2])]-fullData[199^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[175^ /*line XOQFbSla.go:1*/ int(idxKey[8])]+fullData[221^ /*line XOQFbSla.go:1*/ int(idxKey[8])], fullData[145^ /*line XOQFbSla.go:1*/ int(idxKey[10])]+fullData[165^ /*line XOQFbSla.go:1*/ int(idxKey[10])], fullData[49^ /*line XOQFbSla.go:1*/ int(idxKey[5])]-fullData[23^ /*line XOQFbSla.go:1*/ int(idxKey[5])], fullData[113^ /*line XOQFbSla.go:1*/ int(idxKey[5])]^fullData[11^ /*line XOQFbSla.go:1*/ int(idxKey[5])], fullData[229^ /*line XOQFbSla.go:1*/ int(idxKey[8])]-fullData[204^ /*line XOQFbSla.go:1*/ int(idxKey[8])], fullData[227^ /*line XOQFbSla.go:1*/ int(idxKey[4])]-fullData[239^ /*line XOQFbSla.go:1*/ int(idxKey[4])], fullData[127^ /*line XOQFbSla.go:1*/ int(idxKey[12])]+fullData[76^ /*line XOQFbSla.go:1*/ int(idxKey[12])], fullData[218^ /*line XOQFbSla.go:1*/ int(idxKey[8])]^fullData[213^ /*line XOQFbSla.go:1*/ int(idxKey[8])], fullData[209^ /*line XOQFbSla.go:1*/ int(idxKey[2])]-fullData[206^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[224^ /*line XOQFbSla.go:1*/ int(idxKey[2])]-fullData[253^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[221^ /*line XOQFbSla.go:1*/ int(idxKey[1])]+fullData[233^ /*line XOQFbSla.go:1*/ int(idxKey[1])], fullData[150^ /*line XOQFbSla.go:1*/ int(idxKey[7])]^fullData[209^ /*line XOQFbSla.go:1*/ int(idxKey[7])], fullData[143^ /*line XOQFbSla.go:1*/ int(idxKey[10])]-fullData[138^ /*line XOQFbSla.go:1*/ int(idxKey[10])], fullData[204^ /*line XOQFbSla.go:1*/ int(idxKey[2])]^fullData[208^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[208^ /*line XOQFbSla.go:1*/ int(idxKey[4])]^fullData[230^ /*line XOQFbSla.go:1*/ int(idxKey[4])], fullData[249^ /*line XOQFbSla.go:1*/ int(idxKey[2])]+fullData[246^ /*line XOQFbSla.go:1*/ int(idxKey[2])], fullData[134^ /*line XOQFbSla.go:1*/ int(idxKey[7])]+fullData[172^ /*line XOQFbSla.go:1*/ int(idxKey[7])])
			return /*line XOQFbSla.go:1*/ string(data)
		}(), oBVxZOkL)
	}
	if hP0dFanm4.BaseFee != nil {
		if qYOIEJ := /*line _jw0H2aL1EEW.go:1*/ hP0dFanm4.BaseFee.BitLen(); qYOIEJ > 256 {
			return /*line RtwPBaBpYLge.go:1*/ fmt.Errorf(func() /*line nqIsaC0P.go:1*/ string {
				data := /*line nqIsaC0P.go:1*/ make([]byte, 0, 30)
				i := 6
				decryptKey := 67
				for counter := 0; i != 2; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 3:
						data = /*line nqIsaC0P.go:1*/ append(data, 51)
						i = 0
					case 8:
						data = /*line nqIsaC0P.go:1*/ append(data, 48)
						i = 11
					case 6:
						data = /*line nqIsaC0P.go:1*/ append(data, 42)
						i = 8
					case 7:
						data = /*line nqIsaC0P.go:1*/ append(data, "p7"...,
						)
						i = 12
					case 11:
						i = 1
						data = /*line nqIsaC0P.go:1*/ append(data, "3}6:"...,
						)
					case 12:
						i = 13
						data = /*line nqIsaC0P.go:1*/ append(data, "+*vm"...,
						)
					case 5:
						data = /*line nqIsaC0P.go:1*/ append(data, "#)d"...,
						)
						i = 10
					case 9:
						for y := range data {
							data[y] = data[y] ^ /*line nqIsaC0P.go:1*/ byte(decryptKey^y)
						}
						i = 2
					case 0:
						data = /*line nqIsaC0P.go:1*/ append(data, 119)
						i = 4
					case 10:
						i = 9
						data = /*line nqIsaC0P.go:1*/ append(data, "`&"...,
						)
					case 4:
						i = 7
						data = /*line nqIsaC0P.go:1*/ append(data, "64!6"...,
						)
					case 13:
						data = /*line nqIsaC0P.go:1*/ append(data, "(\"<%"...,
						)
						i = 5
					case 1:
						data = /*line nqIsaC0P.go:1*/ append(data, "*>"...,
						)
						i = 3
					}
				}
				return /*line nqIsaC0P.go:1*/ string(data)
			}(), qYOIEJ)
		}
	}
	return nil
}

func (hP0dFanm4 *Header) EmptyBody() bool {
	if hP0dFanm4.WithdrawalsHash != nil {
		return hP0dFanm4.TxHash == DZtBYFgNK && *hP0dFanm4.WithdrawalsHash == VuQdlzM
	}
	return hP0dFanm4.TxHash == DZtBYFgNK && hP0dFanm4.UncleHash == G23a9r1r
}

func (hP0dFanm4 *Header) EmptyReceipts() bool {
	return hP0dFanm4.ReceiptHash == JKp5Sg4r1
}

type AKziBPwSf struct {
	EBPTugX   []*Transaction
	UdMUXwaMt []*Header
	WzYFfWv   []*Withdrawal `rlp:"optional"`
}

type Block struct {
	header       *Header
	uncles       []*Header
	transactions Transactions
	withdrawals  Withdrawals

	hash atomic.Value
	size atomic.Value

	ReceivedAt   time.Time
	ReceivedFrom interface{}
}

type extblock struct {
	Header      *Header
	Txs         []*Transaction
	Uncles      []*Header
	Withdrawals []*Withdrawal `rlp:"optional"`
}

func JtnNKnZ(a2qBKtL *Header, azXme0T7mL []*Transaction, e44zjta []*Header, zCSgFJubD2 []*Receipt, qs4y7ialyr ZGJpq7HtE) *Block {
	aYao5YS := &Block{header: /*line QYi_B1K.go:1*/ WVZZUslxWdK(a2qBKtL)}

	if /*line DbDQRz.go:1*/ len(azXme0T7mL) == 0 {
		aYao5YS.header.TxHash = DZtBYFgNK
	} else {
		aYao5YS.header.TxHash = /*line Po2cbigzAb.go:1*/ HXf5z8W_( /*line B4UcGcMaU1.go:1*/ Transactions(azXme0T7mL), qs4y7ialyr)
		aYao5YS.transactions = /*line BVp0VyZ.go:1*/ make(Transactions /*line x9A1XU.go:1*/, len(azXme0T7mL))
		/*line M14x6cwIzTK.go:1*/ copy(aYao5YS.transactions, azXme0T7mL)
	}

	if /*line IR0pkiYa50ag.go:1*/ len(zCSgFJubD2) == 0 {
		aYao5YS.header.ReceiptHash = JKp5Sg4r1
	} else {
		aYao5YS.header.ReceiptHash = /*line Uim6Xkn7tmW.go:1*/ HXf5z8W_( /*line Cdosbw.go:1*/ Receipts(zCSgFJubD2), qs4y7ialyr)
		aYao5YS.header.Bloom = /*line D8iCeU.go:1*/ T9aqYLj(zCSgFJubD2)
	}

	if /*line LKTeogpZnd.go:1*/ len(e44zjta) == 0 {
		aYao5YS.header.UncleHash = G23a9r1r
	} else {
		aYao5YS.header.UncleHash = /*line t8mQ3k69dMG0.go:1*/ YpHxy5(e44zjta)
		aYao5YS.uncles = /*line KAtGUHzseQWR.go:1*/ make([]*Header /*line WLGBJkAnwjZI.go:1*/, len(e44zjta))
		for ibcOzM6f := range e44zjta {
			aYao5YS.uncles[ibcOzM6f] = /*line grDxj6H.go:1*/ WVZZUslxWdK(e44zjta[ibcOzM6f])
		}
	}

	return aYao5YS
}

func COlwYyG(a2qBKtL *Header, azXme0T7mL []*Transaction, e44zjta []*Header, zCSgFJubD2 []*Receipt, yhduqYTk_Lzk []*Withdrawal, qs4y7ialyr ZGJpq7HtE) *Block {
	aYao5YS := /*line _s7KOLtN.go:1*/ JtnNKnZ(a2qBKtL, azXme0T7mL, e44zjta, zCSgFJubD2, qs4y7ialyr)

	if yhduqYTk_Lzk == nil {
		aYao5YS.header.WithdrawalsHash = nil
	} else if /*line QIzT0n.go:1*/ len(yhduqYTk_Lzk) == 0 {
		aYao5YS.header.WithdrawalsHash = &VuQdlzM
	} else {
		hP0dFanm4 := /*line CtBytV.go:1*/ HXf5z8W_( /*line d1uMFh.go:1*/ Withdrawals(yhduqYTk_Lzk), qs4y7ialyr)
		aYao5YS.header.WithdrawalsHash = &hP0dFanm4
	}

	return /*line oZ5yiM.go:1*/ aYao5YS.WithWithdrawals(yhduqYTk_Lzk)
}

func WVZZUslxWdK(hP0dFanm4 *Header) *Header {
	h5R0LDrm := *hP0dFanm4
	if h5R0LDrm.Difficulty = /*line t6P1QxJE28.go:1*/ new(big.Int); hP0dFanm4.Difficulty != nil {
		/*line vInV5ArCJ.go:1*/ h5R0LDrm.Difficulty.Set(hP0dFanm4.Difficulty)
	}
	if h5R0LDrm.Number = /*line awfoG52r.go:1*/ new(big.Int); hP0dFanm4.Number != nil {
		/*line SxNjHcy0.go:1*/ h5R0LDrm.Number.Set(hP0dFanm4.Number)
	}
	if hP0dFanm4.BaseFee != nil {
		h5R0LDrm.BaseFee = /*line gPN_4qH.go:1*/ new(big.Int).Set(hP0dFanm4.BaseFee)
	}
	if /*line OmaNsDmel.go:1*/ len(hP0dFanm4.Extra) > 0 {
		h5R0LDrm.Extra = /*line VGO7NuPMl.go:1*/ make([]byte /*line SGDWCWo.go:1*/, len(hP0dFanm4.Extra))
		/*line HPbutDyhe.go:1*/ copy(h5R0LDrm.Extra, hP0dFanm4.Extra)
	}
	if hP0dFanm4.WithdrawalsHash != nil {
		h5R0LDrm.WithdrawalsHash = /*line CqQ_5naBc.go:1*/ new(common.Hash)
		*h5R0LDrm.WithdrawalsHash = *hP0dFanm4.WithdrawalsHash
	}
	if hP0dFanm4.ExcessBlobGas != nil {
		h5R0LDrm.ExcessBlobGas = /*line GcrfTEOBXtr.go:1*/ new(uint64)
		*h5R0LDrm.ExcessBlobGas = *hP0dFanm4.ExcessBlobGas
	}
	if hP0dFanm4.BlobGasUsed != nil {
		h5R0LDrm.BlobGasUsed = /*line GJPV1MU.go:1*/ new(uint64)
		*h5R0LDrm.BlobGasUsed = *hP0dFanm4.BlobGasUsed
	}
	if hP0dFanm4.ParentBeaconRoot != nil {
		h5R0LDrm.ParentBeaconRoot = /*line HmBmDaATW.go:1*/ new(common.Hash)
		*h5R0LDrm.ParentBeaconRoot = *hP0dFanm4.ParentBeaconRoot
	}
	return &h5R0LDrm
}

func (aYao5YS *Block) DecodeRLP(gDp0rg *rlp.Stream) error {
	var gKX11Q extblock
	_, rMe9pIag, _ := /*line hn8dse_nu.go:1*/ gDp0rg.Kind()
	if rfteMJju := /*line WbAjftAeMF.go:1*/ gDp0rg.Decode(&gKX11Q); rfteMJju != nil {
		return rfteMJju
	}
	aYao5YS.header, aYao5YS.uncles, aYao5YS.transactions, aYao5YS.withdrawals = gKX11Q.Header, gKX11Q.Uncles, gKX11Q.Txs, gKX11Q.Withdrawals
	/*line Sxtit5Di5nq.go:1*/ aYao5YS.size.Store( /*line CujJAPd.go:1*/ rlp.ListSize(rMe9pIag))
	return nil
}

func (aYao5YS *Block) EncodeRLP(yabbhOmaVr io.Writer) error {
	return /*line Hb6xRBbuA82d.go:1*/ rlp.Encode(yabbhOmaVr, &extblock{
		Header:      aYao5YS.header,
		Txs:         aYao5YS.transactions,
		Uncles:      aYao5YS.uncles,
		Withdrawals: aYao5YS.withdrawals,
	})
}

func (aYao5YS *Block) Body() *AKziBPwSf {
	return &AKziBPwSf{aYao5YS.transactions, aYao5YS.uncles, aYao5YS.withdrawals}
}

func (aYao5YS *Block) Uncles() []*Header          { return aYao5YS.uncles }
func (aYao5YS *Block) Transactions() Transactions { return aYao5YS.transactions }
func (aYao5YS *Block) Withdrawals() Withdrawals   { return aYao5YS.withdrawals }

func (aYao5YS *Block) Transaction(beZs67BXheH common.Hash) *Transaction {
	for _, gvjsfo := range aYao5YS.transactions {
		if /*line tXPCItr4q.go:1*/ gvjsfo.Hash() == beZs67BXheH {
			return gvjsfo
		}
	}
	return nil
}

func (aYao5YS *Block) Header() *Header {
	return /*line lbn35VsBMRr.go:1*/ WVZZUslxWdK(aYao5YS.header)
}

func (aYao5YS *Block) Number() *big.Int {
	return /*line qE2PpEic8UMZ.go:1*/ new(big.Int).Set(aYao5YS.header.Number)
}
func (aYao5YS *Block) GasLimit() uint64 { return aYao5YS.header.GasLimit }
func (aYao5YS *Block) GasUsed() uint64  { return aYao5YS.header.GasUsed }
func (aYao5YS *Block) Difficulty() *big.Int {
	return /*line JRgatY.go:1*/ new(big.Int).Set(aYao5YS.header.Difficulty)
}
func (aYao5YS *Block) Time() uint64 { return aYao5YS.header.Time }

func (aYao5YS *Block) NumberU64() uint64 {
	return /*line _yup3ozTKIsL.go:1*/ aYao5YS.header.Number.Uint64()
}
func (aYao5YS *Block) MixDigest() common.Hash { return aYao5YS.header.MixDigest }
func (aYao5YS *Block) Nonce() uint64 {
	return /*line W0qPvaZK_W.go:1*/ binary.BigEndian.Uint64(aYao5YS.header.Nonce[:])
}
func (aYao5YS *Block) Bloom() Bloom             { return aYao5YS.header.Bloom }
func (aYao5YS *Block) Coinbase() common.Address { return aYao5YS.header.Coinbase }
func (aYao5YS *Block) Root() common.Hash        { return aYao5YS.header.Root }
func (aYao5YS *Block) ParentHash() common.Hash  { return aYao5YS.header.ParentHash }
func (aYao5YS *Block) TxHash() common.Hash      { return aYao5YS.header.TxHash }
func (aYao5YS *Block) ReceiptHash() common.Hash { return aYao5YS.header.ReceiptHash }
func (aYao5YS *Block) UncleHash() common.Hash   { return aYao5YS.header.UncleHash }
func (aYao5YS *Block) Extra() []byte {
	return /*line JOwZw6n.go:1*/ common.CopyBytes(aYao5YS.header.Extra)
}

func (aYao5YS *Block) BaseFee() *big.Int {
	if aYao5YS.header.BaseFee == nil {
		return nil
	}
	return /*line bPzX512.go:1*/ new(big.Int).Set(aYao5YS.header.BaseFee)
}

func (aYao5YS *Block) BeaconRoot() *common.Hash { return aYao5YS.header.ParentBeaconRoot }

func (aYao5YS *Block) ExcessBlobGas() *uint64 {
	var y_6Jb0Dumwf7 *uint64
	if aYao5YS.header.ExcessBlobGas != nil {
		y_6Jb0Dumwf7 = /*line NvEE9sYqA.go:1*/ new(uint64)
		*y_6Jb0Dumwf7 = *aYao5YS.header.ExcessBlobGas
	}
	return y_6Jb0Dumwf7
}

func (aYao5YS *Block) BlobGasUsed() *uint64 {
	var he60UtaRoQBu *uint64
	if aYao5YS.header.BlobGasUsed != nil {
		he60UtaRoQBu = /*line Cy2Gwa.go:1*/ new(uint64)
		*he60UtaRoQBu = *aYao5YS.header.BlobGasUsed
	}
	return he60UtaRoQBu
}

func (aYao5YS *Block) Size() uint64 {
	if rMe9pIag := /*line TH9GarHQDc.go:1*/ aYao5YS.size.Load(); rMe9pIag != nil {
		return rMe9pIag.(uint64)
	}
	eGfuYE8eL4Cc := /*line ID0g_PO.go:1*/ ceFF15RJY1(0)
	/*line PkXGfhD7r1c.go:1*/ rlp.Encode(&eGfuYE8eL4Cc, aYao5YS)
	/*line SKaPu9a.go:1*/ aYao5YS.size.Store( /*line CfwGd6KLrk.go:1*/ uint64(eGfuYE8eL4Cc))
	return /*line QimdyG.go:1*/ uint64(eGfuYE8eL4Cc)
}

func (aYao5YS *Block) SanityCheck() error {
	return /*line IwwEIUyTpL2W.go:1*/ aYao5YS.header.SanityCheck()
}

type ceFF15RJY1 uint64

func (eGfuYE8eL4Cc *ceFF15RJY1) Write(aYao5YS []byte) (int, error) {
	*eGfuYE8eL4Cc += /*line Fx4LGhER0y6l.go:1*/ ceFF15RJY1( /*line JaOl1L.go:1*/ len(aYao5YS))
	return /*line LtpXpCT.go:1*/ len(aYao5YS), nil
}

func YpHxy5(e44zjta []*Header) common.Hash {
	if /*line iCfUgv.go:1*/ len(e44zjta) == 0 {
		return G23a9r1r
	}
	return /*line RL40JH3K.go:1*/ uWZWEzDAB(e44zjta)
}

func KgyKRR_LlB(a2qBKtL *Header) *Block {
	return &Block{header: /*line tNcSYv4.go:1*/ WVZZUslxWdK(a2qBKtL)}
}

func (aYao5YS *Block) WithSeal(a2qBKtL *Header) *Block {
	return &Block{
		header:/*line xwSfN4gQ.go:1*/ WVZZUslxWdK(a2qBKtL),
		transactions: aYao5YS.transactions,
		uncles:       aYao5YS.uncles,
		withdrawals:  aYao5YS.withdrawals,
	}
}

func (aYao5YS *Block) WithBody(aBDOglTT []*Transaction, e44zjta []*Header) *Block {
	cDLgvpe2H := &Block{
		header: aYao5YS.header,
		transactions:/*line Ya4sML8Dy2.go:1*/ make([]*Transaction /*line mcIqeoNpm.go:1*/, len(aBDOglTT)),
		uncles:/*line Hzog2qG5.go:1*/ make([]*Header /*line jl48WzrQ.go:1*/, len(e44zjta)),
		withdrawals: aYao5YS.withdrawals,
	}
	/*line BhQAjNsYUA_.go:1*/ copy(cDLgvpe2H.transactions, aBDOglTT)
	for ibcOzM6f := range e44zjta {
		cDLgvpe2H.uncles[ibcOzM6f] = /*line ZL3uMJB4alA8.go:1*/ WVZZUslxWdK(e44zjta[ibcOzM6f])
	}
	return cDLgvpe2H
}

func (aYao5YS *Block) WithWithdrawals(yhduqYTk_Lzk []*Withdrawal) *Block {
	cDLgvpe2H := &Block{
		header:       aYao5YS.header,
		transactions: aYao5YS.transactions,
		uncles:       aYao5YS.uncles,
	}
	if yhduqYTk_Lzk != nil {
		cDLgvpe2H.withdrawals = /*line lv2fUSki.go:1*/ make([]*Withdrawal /*line Q9oxsqrZvlsS.go:1*/, len(yhduqYTk_Lzk))
		/*line esl6nOg69D.go:1*/ copy(cDLgvpe2H.withdrawals, yhduqYTk_Lzk)
	}
	return cDLgvpe2H
}

func (aYao5YS *Block) Hash() common.Hash {
	if beZs67BXheH := /*line nNURLKaU.go:1*/ aYao5YS.hash.Load(); beZs67BXheH != nil {
		return beZs67BXheH.(common.Hash)
	}
	pW1UMRn1s := /*line BmC8rp4HGgc.go:1*/ aYao5YS.header.Hash()
	/*line w1dFlg9sV9.go:1*/ aYao5YS.hash.Store(pW1UMRn1s)
	return pW1UMRn1s
}

type Atnp8i6A []*Block

func UoWI0WX(a2qBKtL []byte) common.Hash {

	ud_t2GI4, _, rfteMJju := /*line D0aMIwY6.go:1*/ rlp.SplitList(a2qBKtL)
	if rfteMJju != nil {
		return common.Hash{}
	}
	oBTwIB, _, rfteMJju := /*line VpbrYW4aco06.go:1*/ rlp.SplitString(ud_t2GI4)
	if rfteMJju != nil {
		return common.Hash{}
	}
	if /*line aE_6ogaJywo3.go:1*/ len(oBTwIB) != 32 {
		return common.Hash{}
	}
	return /*line jScvBVT1M.go:1*/ common.BytesToHash(oBTwIB)
}

var _ = binary.Append
var _ = fmt.Append
var _ io.ByteReader

const _ = big.Above

var _ = reflect.Append
var _ = atomic.AddInt32

const _ = time.ANSIC

var _ = common.AbsolutePath
var _ hexutil.Big
var _ = rlp.AppendUint64
