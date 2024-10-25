//line :1
package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (* /*line VuWEJvA.go:1*/ n_Xw1garFo)(nil)

func (fWuxre8U Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		Type              hexutil.Uint64 `json:"type,omitempty"`
		PostState         hexutil.Bytes  `json:"root"`
		Status            hexutil.Uint64 `json:"status"`
		CumulativeGasUsed hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             Bloom          `json:"logsBloom"         gencodec:"required"`
		Logs              []*Log         `json:"logs"              gencodec:"required"`
		TxHash            common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   common.Address `json:"contractAddress"`
		GasUsed           hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		EffectiveGasPrice *hexutil.Big   `json:"effectiveGasPrice"`
		BlobGasUsed       hexutil.Uint64 `json:"blobGasUsed,omitempty"`
		BlobGasPrice      *hexutil.Big   `json:"blobGasPrice,omitempty"`
		BlockHash         common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big   `json:"blockNumber,omitempty"`
		TransactionIndex  hexutil.Uint   `json:"transactionIndex"`
	}
	var aINiWZ_Jtlzj Receipt
	aINiWZ_Jtlzj.Type = /*line PvIl_7cwj.go:1*/ hexutil.Uint64(fWuxre8U.Type)
	aINiWZ_Jtlzj.PostState = fWuxre8U.PostState
	aINiWZ_Jtlzj.Status = /*line TnflSfas5cA4.go:1*/ hexutil.Uint64(fWuxre8U.Status)
	aINiWZ_Jtlzj.CumulativeGasUsed = /*line dftRjEQCn.go:1*/ hexutil.Uint64(fWuxre8U.CumulativeGasUsed)
	aINiWZ_Jtlzj.Bloom = fWuxre8U.Bloom
	aINiWZ_Jtlzj.Logs = fWuxre8U.Logs
	aINiWZ_Jtlzj.TxHash = fWuxre8U.TxHash
	aINiWZ_Jtlzj.ContractAddress = fWuxre8U.ContractAddress
	aINiWZ_Jtlzj.GasUsed = /*line FS2kCvW9.go:1*/ hexutil.Uint64(fWuxre8U.GasUsed)
	aINiWZ_Jtlzj.EffectiveGasPrice = (* /*line pTWegHUibNqF.go:1*/ hexutil.Big)(fWuxre8U.EffectiveGasPrice)
	aINiWZ_Jtlzj.BlobGasUsed = /*line s9Zzzqe.go:1*/ hexutil.Uint64(fWuxre8U.BlobGasUsed)
	aINiWZ_Jtlzj.BlobGasPrice = (* /*line rLHaInp.go:1*/ hexutil.Big)(fWuxre8U.BlobGasPrice)
	aINiWZ_Jtlzj.BlockHash = fWuxre8U.BlockHash
	aINiWZ_Jtlzj.BlockNumber = (* /*line XLuHn3rJ07.go:1*/ hexutil.Big)(fWuxre8U.BlockNumber)
	aINiWZ_Jtlzj.TransactionIndex = /*line qAvwsNFAztfZ.go:1*/ hexutil.Uint(fWuxre8U.TransactionIndex)
	return /*line _M3q1Js.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (fWuxre8U *Receipt) UnmarshalJSON(uzD2Up []byte) error {
	type Receipt struct {
		Type              *hexutil.Uint64 `json:"type,omitempty"`
		PostState         *hexutil.Bytes  `json:"root"`
		Status            *hexutil.Uint64 `json:"status"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             *Bloom          `json:"logsBloom"         gencodec:"required"`
		Logs              []*Log          `json:"logs"              gencodec:"required"`
		TxHash            *common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   *common.Address `json:"contractAddress"`
		GasUsed           *hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		EffectiveGasPrice *hexutil.Big    `json:"effectiveGasPrice"`
		BlobGasUsed       *hexutil.Uint64 `json:"blobGasUsed,omitempty"`
		BlobGasPrice      *hexutil.Big    `json:"blobGasPrice,omitempty"`
		BlockHash         *common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big    `json:"blockNumber,omitempty"`
		TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
	}
	var otqLrvn1CD Receipt
	if rfteMJju := /*line hK_C3PxKq.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD); rfteMJju != nil {
		return rfteMJju
	}
	if otqLrvn1CD.Type != nil {
		fWuxre8U.Type = /*line QpfY3a863Yi.go:1*/ uint8(*otqLrvn1CD.Type)
	}
	if otqLrvn1CD.PostState != nil {
		fWuxre8U.PostState = *otqLrvn1CD.PostState
	}
	if otqLrvn1CD.Status != nil {
		fWuxre8U.Status = /*line lY08mjS4hqN.go:1*/ uint64(*otqLrvn1CD.Status)
	}
	if otqLrvn1CD.CumulativeGasUsed == nil {
		return /*line luj_3Lil4b.go:1*/ errors.New(func() /*line DCoaxn57.go:1*/ string {
			seed := /*line DCoaxn57.go:1*/ byte(30)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line DCoaxn57.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line DCoaxn57.go:1*/ fnc(115)(248)(250)(240)(26)(227)(23)(167)(92)(239)(8)(244)(28)(227)(17)(225)(70)(202)(31)(240)(233)(10)(88)(247)(164)(30)(228)(24)(233)(15)(9)(239)(3)(29)(210)(6)(30)(222)(26)(230)(13)(81)(231)(200)(25)(253)(172)(106)(199)(10)(22)(224)(25)(246)
			return /*line DCoaxn57.go:1*/ string(data)
		}())
	}
	fWuxre8U.CumulativeGasUsed = /*line JhMPhB2.go:1*/ uint64(*otqLrvn1CD.CumulativeGasUsed)
	if otqLrvn1CD.Bloom == nil {
		return /*line IvrdcNO_c7.go:1*/ errors.New(func() /*line bISHkjlZ.go:1*/ string {
			seed := /*line bISHkjlZ.go:1*/ byte(70)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line bISHkjlZ.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line bISHkjlZ.go:1*/ fnc(179)(98)(206)(156)(46)(97)(187)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(179)(105)(202)(160)(15)(72)(147)(38)(74)(78)(149)(112)(233)(213)(88)(226)(215)(172)(90)(184)(119)(242)
			return /*line bISHkjlZ.go:1*/ string(data)
		}())
	}
	fWuxre8U.Bloom = *otqLrvn1CD.Bloom
	if otqLrvn1CD.Logs == nil {
		return /*line Kfs4Nem.go:1*/ errors.New(func() /*line UBHYYS9EBIAT.go:1*/ string {
			seed := /*line UBHYYS9EBIAT.go:1*/ byte(36)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line UBHYYS9EBIAT.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line UBHYYS9EBIAT.go:1*/ fnc(73)(4)(2)(0)(26)(227)(23)(167)(92)(239)(8)(244)(28)(227)(17)(225)(70)(202)(31)(240)(233)(10)(88)(247)(171)(29)(232)(4)(92)(247)(168)(25)(253)(172)(106)(199)(10)(22)(224)(25)(246)
			return /*line UBHYYS9EBIAT.go:1*/ string(data)
		}())
	}
	fWuxre8U.Logs = otqLrvn1CD.Logs
	if otqLrvn1CD.TxHash == nil {
		return /*line EkrECOD5.go:1*/ errors.New(func() /*line Cg5CvS5n9.go:1*/ string {
			fullData := [] /*line Cg5CvS5n9.go:1*/ byte("\xa7uw\u009a\xfdG\xe9/\xe2\x0f\xdb2\x10-V9\x95\xef\xc0ߘ\xf0\xb60\x1a\x997v\xb2\xe8/Re\xa4\x8b\x1b\x03\x03\x90\x13}S\x06\x12\xa6\x15\xbd\xea\x83\x12İ\x88k8\x9f\xf8/̓`ya\x96b%\v\xf1\x91\xc8}\xbdu\xfa|\xca\xe2\x99Y\x92\xd9\xcd>\xab2\x85\xa0Y\xf0ȝl\xf7\x12j\t\xf8\xf4\xbb^\xfa\xfb\xea")
			idxKey := [] /*line Cg5CvS5n9.go:1*/ byte("d\xd1\xcf\xea\x13\r\x8c\xc7s\xb8\x0e!F\x19\xab\xba")
			data := /*line Cg5CvS5n9.go:1*/ make([]byte, 0, 53)
			data = /*line Cg5CvS5n9.go:1*/ append(data, fullData[8^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]^fullData[120^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[23^ /*line Cg5CvS5n9.go:1*/ int(idxKey[5])]^fullData[84^ /*line Cg5CvS5n9.go:1*/ int(idxKey[5])], fullData[164^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])]^fullData[129^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])], fullData[118^ /*line Cg5CvS5n9.go:1*/ int(idxKey[12])]-fullData[68^ /*line Cg5CvS5n9.go:1*/ int(idxKey[12])], fullData[229^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])]-fullData[215^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])], fullData[96^ /*line Cg5CvS5n9.go:1*/ int(idxKey[11])]-fullData[67^ /*line Cg5CvS5n9.go:1*/ int(idxKey[11])], fullData[59^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])]-fullData[66^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])], fullData[104^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])]-fullData[5^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])], fullData[204^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])]+fullData[194^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])], fullData[32^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]^fullData[66^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[135^ /*line Cg5CvS5n9.go:1*/ int(idxKey[9])]-fullData[174^ /*line Cg5CvS5n9.go:1*/ int(idxKey[9])], fullData[54^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])]^fullData[105^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])], fullData[218^ /*line Cg5CvS5n9.go:1*/ int(idxKey[2])]^fullData[139^ /*line Cg5CvS5n9.go:1*/ int(idxKey[2])], fullData[11^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]+fullData[37^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[176^ /*line Cg5CvS5n9.go:1*/ int(idxKey[14])]^fullData[139^ /*line Cg5CvS5n9.go:1*/ int(idxKey[14])], fullData[159^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])]+fullData[132^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])], fullData[110^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])]^fullData[123^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])], fullData[61^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])]-fullData[106^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])], fullData[19^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])]-fullData[36^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])], fullData[109^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])]+fullData[85^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])], fullData[74^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])]^fullData[90^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])], fullData[221^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])]-fullData[141^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])], fullData[223^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])]-fullData[202^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])], fullData[203^ /*line Cg5CvS5n9.go:1*/ int(idxKey[2])]^fullData[135^ /*line Cg5CvS5n9.go:1*/ int(idxKey[2])], fullData[8^ /*line Cg5CvS5n9.go:1*/ int(idxKey[12])]-fullData[4^ /*line Cg5CvS5n9.go:1*/ int(idxKey[12])], fullData[84^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])]-fullData[1^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])], fullData[44^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])]+fullData[33^ /*line Cg5CvS5n9.go:1*/ int(idxKey[10])], fullData[47^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]-fullData[28^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[43^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])]+fullData[125^ /*line Cg5CvS5n9.go:1*/ int(idxKey[0])], fullData[20^ /*line Cg5CvS5n9.go:1*/ int(idxKey[4])]-fullData[38^ /*line Cg5CvS5n9.go:1*/ int(idxKey[4])], fullData[171^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])]-fullData[130^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])], fullData[108^ /*line Cg5CvS5n9.go:1*/ int(idxKey[11])]+fullData[113^ /*line Cg5CvS5n9.go:1*/ int(idxKey[11])], fullData[68^ /*line Cg5CvS5n9.go:1*/ int(idxKey[11])]-fullData[100^ /*line Cg5CvS5n9.go:1*/ int(idxKey[11])], fullData[53^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]^fullData[94^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[82^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]^fullData[43^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[222^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])]^fullData[218^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])], fullData[135^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])]^fullData[154^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])], fullData[216^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])]-fullData[187^ /*line Cg5CvS5n9.go:1*/ int(idxKey[6])], fullData[52^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]+fullData[26^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[242^ /*line Cg5CvS5n9.go:1*/ int(idxKey[2])]-fullData[223^ /*line Cg5CvS5n9.go:1*/ int(idxKey[2])], fullData[252^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])]+fullData[237^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])], fullData[86^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])]^fullData[82^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])], fullData[90^ /*line Cg5CvS5n9.go:1*/ int(idxKey[4])]-fullData[56^ /*line Cg5CvS5n9.go:1*/ int(idxKey[4])], fullData[47^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])]-fullData[57^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])], fullData[221^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])]-fullData[143^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])], fullData[235^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])]^fullData[248^ /*line Cg5CvS5n9.go:1*/ int(idxKey[1])], fullData[219^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])]^fullData[239^ /*line Cg5CvS5n9.go:1*/ int(idxKey[7])], fullData[109^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])]^fullData[80^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])], fullData[123^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])]-fullData[63^ /*line Cg5CvS5n9.go:1*/ int(idxKey[8])], fullData[14^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]^fullData[13^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])], fullData[185^ /*line Cg5CvS5n9.go:1*/ int(idxKey[3])]+fullData[191^ /*line Cg5CvS5n9.go:1*/ int(idxKey[3])], fullData[61^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])]-fullData[25^ /*line Cg5CvS5n9.go:1*/ int(idxKey[13])])
			return /*line Cg5CvS5n9.go:1*/ string(data)
		}())
	}
	fWuxre8U.TxHash = *otqLrvn1CD.TxHash
	if otqLrvn1CD.ContractAddress != nil {
		fWuxre8U.ContractAddress = *otqLrvn1CD.ContractAddress
	}
	if otqLrvn1CD.GasUsed == nil {
		return /*line xlUa_C.go:1*/ errors.New(func() /*line QTN72sfIdwc.go:1*/ string {
			fullData := [] /*line QTN72sfIdwc.go:1*/ byte("\x1bY\x01\x17\xce7ej;3_\xebp\xb37\xfe\xbb\"\x8a\xb7\xceB\x80\x1d\xe1\xa7\uef401\x81\xacgQ\x92\x81\xf6\x93\xff\xe4\xefT\xda\x10\xc2\xe9\xcau\xd7/\x14\xf5\x8a\xad\xe2Y\x9b)J\x06h\xb7\xb1\xdaM!IS\x0eH\xbf\xf9\re\x9c\xad\xfb\x86\xa2\xcfO\x7f\xa1K\xc2\xf4\x89\xd5")
			idxKey := [] /*line QTN72sfIdwc.go:1*/ byte(")\xb0\x14\x03g")
			data := /*line QTN72sfIdwc.go:1*/ make([]byte, 0, 45)
			data = /*line QTN72sfIdwc.go:1*/ append(data, fullData[141^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]-fullData[138^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[59^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]+fullData[23^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[51^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]+fullData[89^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[128^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]+fullData[250^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[89^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])]-fullData[3^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])], fullData[40^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[77^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[247^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]-fullData[146^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[41^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])]^fullData[33^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])], fullData[16^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])]-fullData[58^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])], fullData[68^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]+fullData[64^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[165^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]+fullData[129^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[28^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[13^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[157^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]^fullData[166^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[76^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]^fullData[24^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[77^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]-fullData[72^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[6^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]^fullData[64^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[53^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]^fullData[121^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[25^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])]-fullData[84^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])], fullData[91^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]^fullData[101^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[69^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]^fullData[60^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[37^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[38^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[162^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]^fullData[170^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[82^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[9^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[139^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]^fullData[241^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[30^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])]+fullData[109^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])], fullData[49^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])]+fullData[53^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])], fullData[4^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])]-fullData[81^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])], fullData[80^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[39^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[52^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])]-fullData[65^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])], fullData[96^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]+fullData[43^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[40^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])]-fullData[26^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])], fullData[251^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]^fullData[132^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[161^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]+fullData[191^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[127^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])]^fullData[1^ /*line QTN72sfIdwc.go:1*/ int(idxKey[0])], fullData[156^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]^fullData[133^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[47^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]+fullData[97^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[145^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]-fullData[173^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[188^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]+fullData[134^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[110^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])]-fullData[99^ /*line QTN72sfIdwc.go:1*/ int(idxKey[4])], fullData[38^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])]+fullData[68^ /*line QTN72sfIdwc.go:1*/ int(idxKey[2])], fullData[45^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[74^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])], fullData[153^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]-fullData[187^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[179^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])]-fullData[169^ /*line QTN72sfIdwc.go:1*/ int(idxKey[1])], fullData[65^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])]-fullData[84^ /*line QTN72sfIdwc.go:1*/ int(idxKey[3])])
			return /*line QTN72sfIdwc.go:1*/ string(data)
		}())
	}
	fWuxre8U.GasUsed = /*line u72FMcAelV.go:1*/ uint64(*otqLrvn1CD.GasUsed)
	if otqLrvn1CD.EffectiveGasPrice != nil {
		fWuxre8U.EffectiveGasPrice = (* /*line DvVaw0PVgiD.go:1*/ big.Int)(otqLrvn1CD.EffectiveGasPrice)
	}
	if otqLrvn1CD.BlobGasUsed != nil {
		fWuxre8U.BlobGasUsed = /*line sljeQzu.go:1*/ uint64(*otqLrvn1CD.BlobGasUsed)
	}
	if otqLrvn1CD.BlobGasPrice != nil {
		fWuxre8U.BlobGasPrice = (* /*line alOoCfoay0C.go:1*/ big.Int)(otqLrvn1CD.BlobGasPrice)
	}
	if otqLrvn1CD.BlockHash != nil {
		fWuxre8U.BlockHash = *otqLrvn1CD.BlockHash
	}
	if otqLrvn1CD.BlockNumber != nil {
		fWuxre8U.BlockNumber = (* /*line ZLbSADJNym2G.go:1*/ big.Int)(otqLrvn1CD.BlockNumber)
	}
	if otqLrvn1CD.TransactionIndex != nil {
		fWuxre8U.TransactionIndex = /*line IrS9HwfGaVBi.go:1*/ uint(*otqLrvn1CD.TransactionIndex)
	}
	return nil
}

var _ = json.Compact
var _ = errors.As

const _ = big.Above

var _ = common.AbsolutePath
var _ hexutil.Big
