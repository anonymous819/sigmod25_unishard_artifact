//line :1
package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (* /*line BIVwUdA.go:1*/ dCCagad0IRZF)(nil)

func (hP0dFanm4 Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash       common.Hash     `json:"parentHash"       gencodec:"required"`
		UncleHash        common.Hash     `json:"sha3Uncles"       gencodec:"required"`
		Coinbase         common.Address  `json:"miner"`
		Root             common.Hash     `json:"stateRoot"        gencodec:"required"`
		TxHash           common.Hash     `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash      common.Hash     `json:"receiptsRoot"     gencodec:"required"`
		Bloom            Bloom           `json:"logsBloom"        gencodec:"required"`
		Difficulty       *hexutil.Big    `json:"difficulty"       gencodec:"required"`
		Number           *hexutil.Big    `json:"number"           gencodec:"required"`
		GasLimit         hexutil.Uint64  `json:"gasLimit"         gencodec:"required"`
		GasUsed          hexutil.Uint64  `json:"gasUsed"          gencodec:"required"`
		Time             hexutil.Uint64  `json:"timestamp"        gencodec:"required"`
		Extra            hexutil.Bytes   `json:"extraData"        gencodec:"required"`
		MixDigest        common.Hash     `json:"mixHash"`
		Nonce            BlockNonce      `json:"nonce"`
		BaseFee          *hexutil.Big    `json:"baseFeePerGas" rlp:"optional"`
		WithdrawalsHash  *common.Hash    `json:"withdrawalsRoot" rlp:"optional"`
		BlobGasUsed      *hexutil.Uint64 `json:"blobGasUsed" rlp:"optional"`
		ExcessBlobGas    *hexutil.Uint64 `json:"excessBlobGas" rlp:"optional"`
		ParentBeaconRoot *common.Hash    `json:"parentBeaconBlockRoot" rlp:"optional"`
		Hash             common.Hash     `json:"hash"`
	}
	var aINiWZ_Jtlzj Header
	aINiWZ_Jtlzj.ParentHash = hP0dFanm4.ParentHash
	aINiWZ_Jtlzj.UncleHash = hP0dFanm4.UncleHash
	aINiWZ_Jtlzj.Coinbase = hP0dFanm4.Coinbase
	aINiWZ_Jtlzj.Root = hP0dFanm4.Root
	aINiWZ_Jtlzj.TxHash = hP0dFanm4.TxHash
	aINiWZ_Jtlzj.ReceiptHash = hP0dFanm4.ReceiptHash
	aINiWZ_Jtlzj.Bloom = hP0dFanm4.Bloom
	aINiWZ_Jtlzj.Difficulty = (* /*line P4axKlDxgy.go:1*/ hexutil.Big)(hP0dFanm4.Difficulty)
	aINiWZ_Jtlzj.Number = (* /*line H6AWpSV.go:1*/ hexutil.Big)(hP0dFanm4.Number)
	aINiWZ_Jtlzj.GasLimit = /*line WJXzHfuRi6ME.go:1*/ hexutil.Uint64(hP0dFanm4.GasLimit)
	aINiWZ_Jtlzj.GasUsed = /*line WtPFPPFKO.go:1*/ hexutil.Uint64(hP0dFanm4.GasUsed)
	aINiWZ_Jtlzj.Time = /*line DBRUbK_9.go:1*/ hexutil.Uint64(hP0dFanm4.Time)
	aINiWZ_Jtlzj.Extra = hP0dFanm4.Extra
	aINiWZ_Jtlzj.MixDigest = hP0dFanm4.MixDigest
	aINiWZ_Jtlzj.Nonce = hP0dFanm4.Nonce
	aINiWZ_Jtlzj.BaseFee = (* /*line LE2AJvuVM.go:1*/ hexutil.Big)(hP0dFanm4.BaseFee)
	aINiWZ_Jtlzj.WithdrawalsHash = hP0dFanm4.WithdrawalsHash
	aINiWZ_Jtlzj.BlobGasUsed = (* /*line sD1uy_D.go:1*/ hexutil.Uint64)(hP0dFanm4.BlobGasUsed)
	aINiWZ_Jtlzj.ExcessBlobGas = (* /*line CcUJzjsfZdCp.go:1*/ hexutil.Uint64)(hP0dFanm4.ExcessBlobGas)
	aINiWZ_Jtlzj.ParentBeaconRoot = hP0dFanm4.ParentBeaconRoot
	aINiWZ_Jtlzj.Hash = /*line evg1qS.go:1*/ hP0dFanm4.Hash()
	return /*line J7EF1u.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (hP0dFanm4 *Header) UnmarshalJSON(uzD2Up []byte) error {
	type Header struct {
		ParentHash       *common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash        *common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase         *common.Address `json:"miner"`
		Root             *common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash           *common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash      *common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom            *Bloom          `json:"logsBloom"        gencodec:"required"`
		Difficulty       *hexutil.Big    `json:"difficulty"       gencodec:"required"`
		Number           *hexutil.Big    `json:"number"           gencodec:"required"`
		GasLimit         *hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed          *hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		Time             *hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
		Extra            *hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		MixDigest        *common.Hash    `json:"mixHash"`
		Nonce            *BlockNonce     `json:"nonce"`
		BaseFee          *hexutil.Big    `json:"baseFeePerGas" rlp:"optional"`
		WithdrawalsHash  *common.Hash    `json:"withdrawalsRoot" rlp:"optional"`
		BlobGasUsed      *hexutil.Uint64 `json:"blobGasUsed" rlp:"optional"`
		ExcessBlobGas    *hexutil.Uint64 `json:"excessBlobGas" rlp:"optional"`
		ParentBeaconRoot *common.Hash    `json:"parentBeaconBlockRoot" rlp:"optional"`
	}
	var otqLrvn1CD Header
	if rfteMJju := /*line P6mLxoZKa1S5.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD); rfteMJju != nil {
		return rfteMJju
	}
	if otqLrvn1CD.ParentHash == nil {
		return /*line bonKBmAJL.go:1*/ errors.New(func() /*line jh6eO5KFxNvF.go:1*/ string {
			fullData := [] /*line jh6eO5KFxNvF.go:1*/ byte("\xae\x1e\xb1V츈\xe5\xf5T\x9cч\x05\xc4\xe6\xafx\x98\xd9r!\x00w\x1f\xf4^\xb8\x8dx\xf8\x1b\xfe\xb0~\xa1\xfd\xd7B\xb3\xc7#Y\x9c\x9a\x112\xea@\xa0\x00J\x18\x1d6ր\x93\x03j3\xc18\xddz\x8f\x82\xdfN\xcbo\xe1_n\xf4+H\xa0\xba\x8eO\x16\xb5*\xbd)\x95Dmy\xf4\x80")
			idxKey := [] /*line jh6eO5KFxNvF.go:1*/ byte(";\xd4\t\xf4\x80;\xdf~\xb4\xdf\x0e\x92\x98\x10e\xbf")
			data := /*line jh6eO5KFxNvF.go:1*/ make([]byte, 0, 47)
			data = /*line jh6eO5KFxNvF.go:1*/ append(data, fullData[159^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[12])]+fullData[158^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[12])], fullData[171^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])]^fullData[160^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])], fullData[186^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])]-fullData[155^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])], fullData[222^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])]^fullData[167^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])], fullData[29^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[5])]-fullData[40^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[5])], fullData[32^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[0])]^fullData[12^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[0])], fullData[151^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])]+fullData[130^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])], fullData[146^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[6])]-fullData[231^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[6])], fullData[192^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])]+fullData[198^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])], fullData[163^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[4])]^fullData[142^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[4])], fullData[74^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])]-fullData[64^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])], fullData[139^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[8])]+fullData[166^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[8])], fullData[92^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])]^fullData[57^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])], fullData[246^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[9])]+fullData[143^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[9])], fullData[232^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])]^fullData[170^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])], fullData[210^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[6])]+fullData[151^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[6])], fullData[169^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])]-fullData[161^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])], fullData[248^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])]^fullData[179^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])], fullData[154^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])]^fullData[152^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[11])], fullData[241^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[1])]+fullData[155^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[1])], fullData[238^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])]^fullData[218^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])], fullData[198^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])]-fullData[223^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])], fullData[229^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[1])]-fullData[143^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[1])], fullData[22^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[10])]^fullData[48^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[10])], fullData[37^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[5])]+fullData[38^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[5])], fullData[98^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[5])]^fullData[15^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[5])], fullData[251^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[6])]^fullData[158^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[6])], fullData[53^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[7])]^fullData[58^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[7])], fullData[180^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])]+fullData[190^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])], fullData[30^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])]^fullData[51^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])], fullData[55^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[13])]+fullData[70^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[13])], fullData[22^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[0])]-fullData[26^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[0])], fullData[163^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])]^fullData[159^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])], fullData[60^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[7])]^fullData[81^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[7])], fullData[240^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])]^fullData[177^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])], fullData[226^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[1])]^fullData[133^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[1])], fullData[15^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[10])]^fullData[31^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[10])], fullData[249^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])]+fullData[169^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[15])], fullData[246^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])]+fullData[201^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[3])], fullData[71^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])]-fullData[37^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])], fullData[187^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[8])]^fullData[180^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[8])], fullData[69^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])]+fullData[60^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[2])], fullData[38^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[7])]+fullData[103^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[7])], fullData[110^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[14])]+fullData[92^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[14])], fullData[131^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[4])]^fullData[188^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[4])], fullData[194^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[12])]+fullData[186^ /*line jh6eO5KFxNvF.go:1*/ int(idxKey[12])])
			return /*line jh6eO5KFxNvF.go:1*/ string(data)
		}())
	}
	hP0dFanm4.ParentHash = *otqLrvn1CD.ParentHash
	if otqLrvn1CD.UncleHash == nil {
		return /*line A3fzai0Sc.go:1*/ errors.New(func() /*line KARQr2YJv8F.go:1*/ string {
			key := [] /*line KARQr2YJv8F.go:1*/ byte("\xc4>\x93\x06\x99\xe1\x90n\xe28o\xfd\xcc0\x8a\x02\xa4n\xa08\xb0\xe7\x0e.A4\xc5\xd8eo~\xa2v\r\x12\xfa\x13r\xae\aQq\xae\xe7]\x9e")
			data := [] /*line KARQr2YJv8F.go:1*/ byte("\xa9+\xe0mЍײ\x90-\x02x\x9dB\xdbb|\xf8\xc9-\xbc}\x12\xf924\x9c[\xf0\xff\xe5\xca\xeff\x15&S\xfd\xc4\x19\xf7\xf4\xb3}\b\xd4")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line KARQr2YJv8F.go:1*/ string(data)
		}())
	}
	hP0dFanm4.UncleHash = *otqLrvn1CD.UncleHash
	if otqLrvn1CD.Coinbase != nil {
		hP0dFanm4.Coinbase = *otqLrvn1CD.Coinbase
	}
	if otqLrvn1CD.Root == nil {
		return /*line zaYvlCTfn1C.go:1*/ errors.New(func() /*line HK6Ud5eG.go:1*/ string {
			fullData := [] /*line HK6Ud5eG.go:1*/ byte("\x87Ҷ6\xee#Toz\x1d䥮\u07b3\xb6\xd8+\xd9RGFR\xbfYZ\xad#\xb0G\"\xb5Og\xfc\xb9Qꂽ<\xfe\xdf \x02\x80\x1c+m\xf6<\x1e\xe2\xfe\xb5\x15\x03?\xc20\xe1\x05\xc0\xf0HD\x7f\xbd\x98L\xf4*\xbb$P}\x10Tu\xb6#\xab\x8e\xf2wo\x9c\x00#\x1f")
			idxKey := [] /*line HK6Ud5eG.go:1*/ byte("\x92\xeb\x18\xeeH\xbfm\xa6(֥")
			data := /*line HK6Ud5eG.go:1*/ make([]byte, 0, 46)
			data = /*line HK6Ud5eG.go:1*/ append(data, fullData[83^ /*line HK6Ud5eG.go:1*/ int(idxKey[2])]+fullData[39^ /*line HK6Ud5eG.go:1*/ int(idxKey[2])], fullData[241^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])]^fullData[150^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])], fullData[106^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])]+fullData[110^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])], fullData[117^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])]^fullData[42^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])], fullData[247^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])]+fullData[179^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])], fullData[151^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])]^fullData[169^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])], fullData[119^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])]-fullData[120^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])], fullData[167^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])]^fullData[208^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])], fullData[244^ /*line HK6Ud5eG.go:1*/ int(idxKey[9])]^fullData[132^ /*line HK6Ud5eG.go:1*/ int(idxKey[9])], fullData[129^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])]^fullData[182^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])], fullData[185^ /*line HK6Ud5eG.go:1*/ int(idxKey[0])]+fullData[182^ /*line HK6Ud5eG.go:1*/ int(idxKey[0])], fullData[252^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])]+fullData[184^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])], fullData[24^ /*line HK6Ud5eG.go:1*/ int(idxKey[2])]^fullData[28^ /*line HK6Ud5eG.go:1*/ int(idxKey[2])], fullData[195^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])]+fullData[189^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])], fullData[102^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])]+fullData[83^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])], fullData[56^ /*line HK6Ud5eG.go:1*/ int(idxKey[2])]+fullData[47^ /*line HK6Ud5eG.go:1*/ int(idxKey[2])], fullData[245^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])]-fullData[214^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])], fullData[61^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])]-fullData[46^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])], fullData[247^ /*line HK6Ud5eG.go:1*/ int(idxKey[9])]+fullData[250^ /*line HK6Ud5eG.go:1*/ int(idxKey[9])], fullData[183^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])]^fullData[230^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])], fullData[2^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])]^fullData[38^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])], fullData[176^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])]^fullData[190^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])], fullData[95^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])]-fullData[67^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])], fullData[152^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])]+fullData[187^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])], fullData[227^ /*line HK6Ud5eG.go:1*/ int(idxKey[9])]+fullData[152^ /*line HK6Ud5eG.go:1*/ int(idxKey[9])], fullData[154^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])]-fullData[150^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])], fullData[244^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])]-fullData[166^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])], fullData[245^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])]^fullData[246^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])], fullData[161^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])]+fullData[151^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])], fullData[160^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])]+fullData[143^ /*line HK6Ud5eG.go:1*/ int(idxKey[7])], fullData[200^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])]+fullData[233^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])], fullData[174^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])]^fullData[179^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])], fullData[248^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])]-fullData[230^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])], fullData[214^ /*line HK6Ud5eG.go:1*/ int(idxKey[0])]^fullData[133^ /*line HK6Ud5eG.go:1*/ int(idxKey[0])], fullData[75^ /*line HK6Ud5eG.go:1*/ int(idxKey[4])]+fullData[109^ /*line HK6Ud5eG.go:1*/ int(idxKey[4])], fullData[225^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])]+fullData[205^ /*line HK6Ud5eG.go:1*/ int(idxKey[1])], fullData[7^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])]^fullData[105^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])], fullData[172^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])]^fullData[240^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])], fullData[145^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])]^fullData[159^ /*line HK6Ud5eG.go:1*/ int(idxKey[10])], fullData[45^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])]+fullData[58^ /*line HK6Ud5eG.go:1*/ int(idxKey[6])], fullData[84^ /*line HK6Ud5eG.go:1*/ int(idxKey[4])]+fullData[126^ /*line HK6Ud5eG.go:1*/ int(idxKey[4])], fullData[161^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])]+fullData[191^ /*line HK6Ud5eG.go:1*/ int(idxKey[3])], fullData[53^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])]^fullData[45^ /*line HK6Ud5eG.go:1*/ int(idxKey[8])], fullData[113^ /*line HK6Ud5eG.go:1*/ int(idxKey[4])]^fullData[81^ /*line HK6Ud5eG.go:1*/ int(idxKey[4])], fullData[174^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])]+fullData[171^ /*line HK6Ud5eG.go:1*/ int(idxKey[5])])
			return /*line HK6Ud5eG.go:1*/ string(data)
		}())
	}
	hP0dFanm4.Root = *otqLrvn1CD.Root
	if otqLrvn1CD.TxHash == nil {
		return /*line E3fj3RDEFAg.go:1*/ errors.New(func() /*line aoMUaY0dt5.go:1*/ string {
			fullData := [] /*line aoMUaY0dt5.go:1*/ byte("\xd0\xf49i*6^PS\xad\xe5\xc0\x95\xf6\xe6愰`J\xd3\t\vӡU\x89\xa1\x82\xbc\xe5\x9a͆\\U\xd4d\xb7\xde\xfb\xe2\xc2\xf3\xc5\x1a$\"=G\x9c\x93Z{L'u2A\x91P\xf4Պ1\xf4\x86\xe5C\xe5\xb2;\x80\x81N\x15\x92\a\xeaZtQ4riYS\xfd\x1b\xd9\x1bK\xca!\x9b\x90찹ޯ\xe9\xfe\xd2")
			idxKey := [] /*line aoMUaY0dt5.go:1*/ byte("\x84F?\x82\xcb\xe9\xdb\x1et\xe7\bx\x1f>\x7f\xbd")
			data := /*line aoMUaY0dt5.go:1*/ make([]byte, 0, 53)
			data = /*line aoMUaY0dt5.go:1*/ append(data, fullData[84^ /*line aoMUaY0dt5.go:1*/ int(idxKey[1])]-fullData[109^ /*line aoMUaY0dt5.go:1*/ int(idxKey[1])], fullData[43^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])]-fullData[109^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])], fullData[176^ /*line aoMUaY0dt5.go:1*/ int(idxKey[5])]+fullData[246^ /*line aoMUaY0dt5.go:1*/ int(idxKey[5])], fullData[111^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])]-fullData[93^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])], fullData[21^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])]-fullData[69^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])], fullData[73^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])]^fullData[80^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])], fullData[95^ /*line aoMUaY0dt5.go:1*/ int(idxKey[12])]-fullData[67^ /*line aoMUaY0dt5.go:1*/ int(idxKey[12])], fullData[120^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])]+fullData[127^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])], fullData[48^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])]-fullData[111^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])], fullData[113^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])]^fullData[124^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])], fullData[166^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])]+fullData[207^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])], fullData[195^ /*line aoMUaY0dt5.go:1*/ int(idxKey[5])]^fullData[207^ /*line aoMUaY0dt5.go:1*/ int(idxKey[5])], fullData[79^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])]-fullData[111^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])], fullData[173^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])]^fullData[219^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])], fullData[119^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])]+fullData[33^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])], fullData[112^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])]-fullData[31^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])], fullData[127^ /*line aoMUaY0dt5.go:1*/ int(idxKey[12])]+fullData[77^ /*line aoMUaY0dt5.go:1*/ int(idxKey[12])], fullData[91^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])]+fullData[51^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])], fullData[124^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])]+fullData[47^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])], fullData[244^ /*line aoMUaY0dt5.go:1*/ int(idxKey[9])]+fullData[189^ /*line aoMUaY0dt5.go:1*/ int(idxKey[9])], fullData[109^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])]^fullData[118^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])], fullData[42^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])]-fullData[19^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])], fullData[203^ /*line aoMUaY0dt5.go:1*/ int(idxKey[3])]^fullData[153^ /*line aoMUaY0dt5.go:1*/ int(idxKey[3])], fullData[173^ /*line aoMUaY0dt5.go:1*/ int(idxKey[9])]^fullData[179^ /*line aoMUaY0dt5.go:1*/ int(idxKey[9])], fullData[37^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])]+fullData[71^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])], fullData[117^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])]-fullData[104^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])], fullData[211^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])]^fullData[182^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])], fullData[43^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])]+fullData[97^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])], fullData[6^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])]^fullData[5^ /*line aoMUaY0dt5.go:1*/ int(idxKey[2])], fullData[115^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])]^fullData[96^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])], fullData[55^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])]-fullData[63^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])], fullData[136^ /*line aoMUaY0dt5.go:1*/ int(idxKey[15])]-fullData[240^ /*line aoMUaY0dt5.go:1*/ int(idxKey[15])], fullData[45^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])]-fullData[26^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])], fullData[194^ /*line aoMUaY0dt5.go:1*/ int(idxKey[9])]^fullData[241^ /*line aoMUaY0dt5.go:1*/ int(idxKey[9])], fullData[79^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])]^fullData[41^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])], fullData[177^ /*line aoMUaY0dt5.go:1*/ int(idxKey[15])]^fullData[179^ /*line aoMUaY0dt5.go:1*/ int(idxKey[15])], fullData[62^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])]+fullData[121^ /*line aoMUaY0dt5.go:1*/ int(idxKey[14])], fullData[214^ /*line aoMUaY0dt5.go:1*/ int(idxKey[4])]^fullData[220^ /*line aoMUaY0dt5.go:1*/ int(idxKey[4])], fullData[35^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])]^fullData[86^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])], fullData[93^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])]-fullData[2^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])], fullData[44^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])]-fullData[73^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])], fullData[64^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])]-fullData[91^ /*line aoMUaY0dt5.go:1*/ int(idxKey[11])], fullData[23^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])]+fullData[124^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])], fullData[134^ /*line aoMUaY0dt5.go:1*/ int(idxKey[15])]^fullData[219^ /*line aoMUaY0dt5.go:1*/ int(idxKey[15])], fullData[109^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])]+fullData[18^ /*line aoMUaY0dt5.go:1*/ int(idxKey[10])], fullData[193^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])]-fullData[168^ /*line aoMUaY0dt5.go:1*/ int(idxKey[0])], fullData[122^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])]-fullData[22^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])], fullData[147^ /*line aoMUaY0dt5.go:1*/ int(idxKey[3])]^fullData[188^ /*line aoMUaY0dt5.go:1*/ int(idxKey[3])], fullData[90^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])]+fullData[120^ /*line aoMUaY0dt5.go:1*/ int(idxKey[13])], fullData[84^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])]-fullData[119^ /*line aoMUaY0dt5.go:1*/ int(idxKey[8])], fullData[92^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])]-fullData[67^ /*line aoMUaY0dt5.go:1*/ int(idxKey[7])], fullData[122^ /*line aoMUaY0dt5.go:1*/ int(idxKey[1])]-fullData[97^ /*line aoMUaY0dt5.go:1*/ int(idxKey[1])])
			return /*line aoMUaY0dt5.go:1*/ string(data)
		}())
	}
	hP0dFanm4.TxHash = *otqLrvn1CD.TxHash
	if otqLrvn1CD.ReceiptHash == nil {
		return /*line oikF4h1Iz.go:1*/ errors.New(func() /*line z2y27ph.go:1*/ string {
			key := [] /*line z2y27ph.go:1*/ byte("\xa2\x85{ #zf\x02\xac\x9fp\x9e\x99\U00055aa6\xb4\xa9\x93\xcd\xe9SU\x03\v\"\xe9u\xd6\xf8\xed\x99x\x7fF\nI\xff\x94\xc6\xd4\xd5\xe9\x92@\x80q")
			data := [] /*line z2y27ph.go:1*/ byte("\xcf\xec\bSJ\x14\x01\"\xde\xfa\x01\xeb\xf0\x83\xf0Ά\xd2\xc0\xf6\xa1\x8dsrqnA\x8c\x1c\xa6\x8c\x9e\xcb\x17\x102-i\x99\xfb\xb4\xf4\x9d\x8c\xf3$\xe5\x03")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line z2y27ph.go:1*/ string(data)
		}())
	}
	hP0dFanm4.ReceiptHash = *otqLrvn1CD.ReceiptHash
	if otqLrvn1CD.Bloom == nil {
		return /*line D_fuMNyxXp.go:1*/ errors.New(func() /*line YxE2nrlbfbI.go:1*/ string {
			fullData := [] /*line YxE2nrlbfbI.go:1*/ byte("~\xf3\xdd\xc0\x99\xb3\xab\xc3]\x10\xec0% \xcb\xfb\x89`\xec\xbcf\f$J\x97\xf2\x99\xae\x02\xfb\xe6\xbe\xe8ߌ\xac\x91\xb3\xe7\x97\a\xce4\x80_P\x98\x83\xeb\xf3׳)\x10\xc7\xc3\x15\xe0\xfe5\x83\x12\xfe5\xe2\xa4c^%\a\x9f;tU\xd9_\xf1\x1f\xfdu\x10\x83{\xdap\xa3\x91\x7f̣")
			idxKey := [] /*line YxE2nrlbfbI.go:1*/ byte("\xd6\xd8=Vxq")
			data := /*line YxE2nrlbfbI.go:1*/ make([]byte, 0, 46)
			data = /*line YxE2nrlbfbI.go:1*/ append(data, fullData[88^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]-fullData[21^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[5^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]^fullData[83^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[107^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]-fullData[16^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[49^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]^fullData[39^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[74^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]-fullData[82^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[102^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]^fullData[103^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[76^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]^fullData[108^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[141^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]-fullData[228^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[81^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])]^fullData[107^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])], fullData[16^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])]+fullData[5^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])], fullData[152^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]^fullData[244^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[213^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]+fullData[145^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[34^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])]+fullData[59^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])], fullData[101^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])]-fullData[104^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])], fullData[209^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]^fullData[151^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[54^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])]-fullData[101^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])], fullData[71^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]^fullData[87^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[40^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]+fullData[70^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[106^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])]-fullData[41^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])], fullData[248^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]-fullData[247^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[125^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]+fullData[92^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[123^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])]^fullData[57^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])], fullData[40^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])]+fullData[77^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])], fullData[138^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]+fullData[251^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[225^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]-fullData[144^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[67^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]^fullData[20^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[254^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]^fullData[199^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[59^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]-fullData[101^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[121^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]^fullData[60^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[250^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]-fullData[215^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[232^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]^fullData[242^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[148^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]+fullData[216^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[63^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])]-fullData[105^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])], fullData[252^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]+fullData[231^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[125^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]+fullData[126^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[147^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]+fullData[157^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[120^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]+fullData[100^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[231^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])]-fullData[223^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[1])], fullData[102^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]+fullData[109^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[80^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])]-fullData[86^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[5])], fullData[79^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])]^fullData[78^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[3])], fullData[243^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]+fullData[205^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[76^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])]+fullData[63^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[4])], fullData[200^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])]+fullData[129^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[0])], fullData[121^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])]-fullData[14^ /*line YxE2nrlbfbI.go:1*/ int(idxKey[2])])
			return /*line YxE2nrlbfbI.go:1*/ string(data)
		}())
	}
	hP0dFanm4.Bloom = *otqLrvn1CD.Bloom
	if otqLrvn1CD.Difficulty == nil {
		return /*line CIeUx6.go:1*/ errors.New(func() /*line IPZwIiC9ibTr.go:1*/ string {
			data := /*line IPZwIiC9ibTr.go:1*/ make([]byte, 0, 47)
			i := 5
			decryptKey := 100
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					i = 18
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 156)
				case 11:
					i = 16
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xa8\xf3"...,
					)
				case 15:
					i = 19
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xb3\xbf"...,
					)
				case 14:
					i = 13
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\x81\x89\x97\xc4"...,
					)
				case 9:
					i = 20
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 182)
				case 19:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 190)
					i = 9
				case 20:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xbd\xa8"...,
					)
					i = 3
				case 17:
					i = 11
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xbc\xa8"...,
					)
				case 10:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\x9b\xc6\xc0"...,
					)
					i = 14
				case 21:
					i = 15
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xb2\xf5\xf3\xbf"...,
					)
				case 2:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 179)
					i = 1
				case 16:
					i = 22
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xb4\xb8\xb5"...,
					)
				case 3:
					i = 10
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xb0\x97"...,
					)
				case 13:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 163)
					i = 6
				case 6:
					i = 0
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\x8f\x88\x8c"...,
					)
				case 18:
					for y := range data {
						data[y] = data[y] ^ /*line IPZwIiC9ibTr.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 7:
					i = 12
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xb9\xaf\xb8"...,
					)
				case 5:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xae\xab\xb2"...,
					)
					i = 2
				case 12:
					i = 17
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xbd\xa6"...,
					)
				case 0:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 138)
					i = 4
				case 1:
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, "\xae\xa8\xa2\xe4"...,
					)
					i = 7
				case 22:
					i = 21
					data = /*line IPZwIiC9ibTr.go:1*/ append(data, 187)
				}
			}
			return /*line IPZwIiC9ibTr.go:1*/ string(data)
		}())
	}
	hP0dFanm4.Difficulty = (* /*line fzj1U9A.go:1*/ big.Int)(otqLrvn1CD.Difficulty)
	if otqLrvn1CD.Number == nil {
		return /*line wv6oszyDEQ9.go:1*/ errors.New(func() /*line sjMPKL6z.go:1*/ string {
			key := [] /*line sjMPKL6z.go:1*/ byte("\xd3w\xc4h`=\xd0\xf3K\x8c\x16\x10v\xa4iBY\xbf\xe5\x93\xdcy\xad\x0eH\xe3]+\xbaCqK\x16\xc3(n\a\x19\x96\x00\xf0Y")
			data := [] /*line sjMPKL6z.go:1*/ byte("\x9a\xf2\xaf\v\t1\x97-'\xd9[e\xf3\xce\xfc\"ǧ\x84Ґ\xebs\x19&\x92\x107\xab/\xb6\xd5P\xacJ\xb2AL\xcbdu\x19")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line sjMPKL6z.go:1*/ string(data)
		}())
	}
	hP0dFanm4.Number = (* /*line bycjPPajhc.go:1*/ big.Int)(otqLrvn1CD.Number)
	if otqLrvn1CD.GasLimit == nil {
		return /*line QCKYzqKTP591.go:1*/ errors.New(func() /*line Ky8qdSPsf.go:1*/ string {
			data := [] /*line Ky8qdSPsf.go:1*/ byte("Tessicg2_equ]r\xa7 eky?l\x92\xa1'}\x86\xc6LIJS:w fVrIHlSAs_")
			positions := [...]byte{30, 39, 26, 22, 12, 14, 17, 29, 0, 24, 42, 14, 22, 3, 21, 25, 22, 29, 0, 32, 12, 17, 7, 42, 15, 3, 39, 28, 24, 16, 25, 5, 24, 40, 32, 12, 35, 1, 19, 41, 37, 29, 31, 18, 8, 43, 37, 3, 41, 39, 15, 1, 18, 42}
			for i := 0; i < 54; i += 2 {
				localKey := /*line Ky8qdSPsf.go:1*/ byte(i) + /*line Ky8qdSPsf.go:1*/ byte(positions[i]^positions[i+1]) + 196
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line Ky8qdSPsf.go:1*/ string(data)
		}())
	}
	hP0dFanm4.GasLimit = /*line CxlNv_W.go:1*/ uint64(*otqLrvn1CD.GasLimit)
	if otqLrvn1CD.GasUsed == nil {
		return /*line ZDrgIgF8UIb.go:1*/ errors.New(func() /*line A3r3_l.go:1*/ string {
			data := /*line A3r3_l.go:1*/ make([]byte, 0, 44)
			i := 14
			decryptKey := 75
			for counter := 0; i != 16; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 5:
					i = 3
					data = /*line A3r3_l.go:1*/ append(data, "CG"...,
					)
				case 3:
					i = 12
					data = /*line A3r3_l.go:1*/ append(data, "C\xfb"...,
					)
				case 18:
					data = /*line A3r3_l.go:1*/ append(data, "&;\x1c"...,
					)
					i = 4
				case 7:
					i = 9
					data = /*line A3r3_l.go:1*/ append(data, "46-\xec"...,
					)
				case 8:
					i = 7
					data = /*line A3r3_l.go:1*/ append(data, "7\xee39"...,
					)
				case 4:
					data = /*line A3r3_l.go:1*/ append(data, "5&"...,
					)
					i = 17
				case 11:
					data = /*line A3r3_l.go:1*/ append(data, "SR"...,
					)
					i = 5
				case 12:
					i = 2
					data = /*line A3r3_l.go:1*/ append(data, "H:I"...,
					)
				case 2:
					i = 15
					data = /*line A3r3_l.go:1*/ append(data, "L;C"...,
					)
				case 9:
					data = /*line A3r3_l.go:1*/ append(data, "\xf2-"...,
					)
					i = 18
				case 10:
					data = /*line A3r3_l.go:1*/ append(data, 234)
					i = 6
				case 6:
					i = 13
					data = /*line A3r3_l.go:1*/ append(data, "\x1ecoq"...,
					)
				case 17:
					data = /*line A3r3_l.go:1*/ append(data, 40)
					i = 10
				case 15:
					i = 8
					data = /*line A3r3_l.go:1*/ append(data, 57)
				case 13:
					data = /*line A3r3_l.go:1*/ append(data, "\x1aAa\\"...,
					)
					i = 0
				case 0:
					i = 1
					data = /*line A3r3_l.go:1*/ append(data, "ZZj"...,
					)
				case 14:
					data = /*line A3r3_l.go:1*/ append(data, "KF"...,
					)
					i = 11
				case 1:
					i = 16
					for y := range data {
						data[y] = data[y] + /*line A3r3_l.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return /*line A3r3_l.go:1*/ string(data)
		}())
	}
	hP0dFanm4.GasUsed = /*line HTZYe6s9.go:1*/ uint64(*otqLrvn1CD.GasUsed)
	if otqLrvn1CD.Time == nil {
		return /*line p8Go2kSwLpeo.go:1*/ errors.New(func() /*line w4bEKim1s8d.go:1*/ string {
			data := /*line w4bEKim1s8d.go:1*/ make([]byte, 0, 46)
			i := 0
			decryptKey := 175
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 17:
					i = 8
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xc6\xef\r\n"...,
					)
				case 11:
					i = 18
					data = /*line w4bEKim1s8d.go:1*/ append(data, 194)
				case 10:
					data = /*line w4bEKim1s8d.go:1*/ append(data, 7)
					i = 20
				case 7:
					i = 21
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xee\xfb\x00"...,
					)
				case 8:
					data = /*line w4bEKim1s8d.go:1*/ append(data, 14)
					i = 13
				case 5:
					for y := range data {
						data[y] = data[y] - /*line w4bEKim1s8d.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 13:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\x10\x1e"...,
					)
					i = 5
				case 14:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xf3\xed\xa7\xfa"...,
					)
					i = 7
				case 18:
					i = 17
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\t\x13\x17"...,
					)
				case 16:
					i = 14
					data = /*line w4bEKim1s8d.go:1*/ append(data, 237)
				case 21:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xf5\xff"...,
					)
					i = 4
				case 12:
					i = 10
					data = /*line w4bEKim1s8d.go:1*/ append(data, 2)
				case 1:
					i = 11
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\x10\xc8"...,
					)
				case 4:
					i = 2
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xf3\xf3\xb0"...,
					)
				case 2:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xf7\xfb\xf8"...,
					)
					i = 19
				case 20:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\x00\x0f"...,
					)
					i = 15
				case 19:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\x00\xf9\xb6\xbe"...,
					)
					i = 9
				case 9:
					i = 12
					data = /*line w4bEKim1s8d.go:1*/ append(data, 12)
				case 15:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\x11\xff\f"...,
					)
					i = 1
				case 0:
					i = 3
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xed\xea"...,
					)
				case 3:
					data = /*line w4bEKim1s8d.go:1*/ append(data, "\xf5\xf6"...,
					)
					i = 16
				}
			}
			return /*line w4bEKim1s8d.go:1*/ string(data)
		}())
	}
	hP0dFanm4.Time = /*line HIDS524TwNO.go:1*/ uint64(*otqLrvn1CD.Time)
	if otqLrvn1CD.Extra == nil {
		return /*line r8Tcxf.go:1*/ errors.New(func() /*line H4pfVexEQD.go:1*/ string {
			seed := /*line H4pfVexEQD.go:1*/ byte(81)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line H4pfVexEQD.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line H4pfVexEQD.go:1*/ fnc(28)(252)(10)(0)(246)(5)(249)(185)(82)(243)(12)(4)(244)(9)(243)(255)(188)(70)(3)(252)(7)(248)(188)(7)(62)(19)(252)(254)(239)(227)(29)(19)(237)(198)(249)(70)(9)(3)(174)(40)(29)(252)(3)(1)(13)
			return /*line H4pfVexEQD.go:1*/ string(data)
		}())
	}
	hP0dFanm4.Extra = *otqLrvn1CD.Extra
	if otqLrvn1CD.MixDigest != nil {
		hP0dFanm4.MixDigest = *otqLrvn1CD.MixDigest
	}
	if otqLrvn1CD.Nonce != nil {
		hP0dFanm4.Nonce = *otqLrvn1CD.Nonce
	}
	if otqLrvn1CD.BaseFee != nil {
		hP0dFanm4.BaseFee = (* /*line j8pb5sXHul.go:1*/ big.Int)(otqLrvn1CD.BaseFee)
	}
	if otqLrvn1CD.WithdrawalsHash != nil {
		hP0dFanm4.WithdrawalsHash = otqLrvn1CD.WithdrawalsHash
	}
	if otqLrvn1CD.BlobGasUsed != nil {
		hP0dFanm4.BlobGasUsed = (* /*line PqO8qr6ogA.go:1*/ uint64)(otqLrvn1CD.BlobGasUsed)
	}
	if otqLrvn1CD.ExcessBlobGas != nil {
		hP0dFanm4.ExcessBlobGas = (* /*line mSxcw0V7R.go:1*/ uint64)(otqLrvn1CD.ExcessBlobGas)
	}
	if otqLrvn1CD.ParentBeaconRoot != nil {
		hP0dFanm4.ParentBeaconRoot = otqLrvn1CD.ParentBeaconRoot
	}
	return nil
}

var _ = json.Compact
var _ = errors.As

const _ = big.Above

var _ = common.AbsolutePath
var _ hexutil.Big
