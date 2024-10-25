//line :1
package types

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate go run github.com/fjl/gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go

var (
	fYLiFP      = []byte{}
	fRso7AaJq9a = []byte{0x01}
)

var lMbK8q0 = /*line ABWgvaG.go:1*/ errors.New(func() /*line jqqEWdmG.go:1*/ string {
	fullData := [] /*line jqqEWdmG.go:1*/ byte("\x17\x91y\xbb\x17\x99\x03\xedg\x9b\xafGh\xa0\x04\x970hD'\xd0\xcf\xef$\x1a\xb6\x85\xbc\xafS7\x19\xa7\xfe\xe3\f@\xf0\x965e\xb4s\xf6K\xff")
	idxKey := [] /*line jqqEWdmG.go:1*/ byte("\xcbAe\xaa\x12C砬?r'\xc7fV")
	data := /*line jqqEWdmG.go:1*/ make([]byte, 0, 24)
	data = /*line jqqEWdmG.go:1*/ append(data, fullData[98^ /*line jqqEWdmG.go:1*/ int(idxKey[2])]^fullData[96^ /*line jqqEWdmG.go:1*/ int(idxKey[2])], fullData[68^ /*line jqqEWdmG.go:1*/ int(idxKey[2])]-fullData[127^ /*line jqqEWdmG.go:1*/ int(idxKey[2])], fullData[26^ /*line jqqEWdmG.go:1*/ int(idxKey[4])]^fullData[22^ /*line jqqEWdmG.go:1*/ int(idxKey[4])], fullData[126^ /*line jqqEWdmG.go:1*/ int(idxKey[13])]+fullData[74^ /*line jqqEWdmG.go:1*/ int(idxKey[13])], fullData[246^ /*line jqqEWdmG.go:1*/ int(idxKey[6])]-fullData[233^ /*line jqqEWdmG.go:1*/ int(idxKey[6])], fullData[170^ /*line jqqEWdmG.go:1*/ int(idxKey[3])]^fullData[180^ /*line jqqEWdmG.go:1*/ int(idxKey[3])], fullData[103^ /*line jqqEWdmG.go:1*/ int(idxKey[13])]^fullData[68^ /*line jqqEWdmG.go:1*/ int(idxKey[13])], fullData[94^ /*line jqqEWdmG.go:1*/ int(idxKey[1])]-fullData[104^ /*line jqqEWdmG.go:1*/ int(idxKey[1])], fullData[128^ /*line jqqEWdmG.go:1*/ int(idxKey[7])]+fullData[187^ /*line jqqEWdmG.go:1*/ int(idxKey[7])], fullData[219^ /*line jqqEWdmG.go:1*/ int(idxKey[12])]+fullData[222^ /*line jqqEWdmG.go:1*/ int(idxKey[12])], fullData[134^ /*line jqqEWdmG.go:1*/ int(idxKey[7])]^fullData[141^ /*line jqqEWdmG.go:1*/ int(idxKey[7])], fullData[20^ /*line jqqEWdmG.go:1*/ int(idxKey[4])]^fullData[56^ /*line jqqEWdmG.go:1*/ int(idxKey[4])], fullData[34^ /*line jqqEWdmG.go:1*/ int(idxKey[9])]^fullData[44^ /*line jqqEWdmG.go:1*/ int(idxKey[9])], fullData[72^ /*line jqqEWdmG.go:1*/ int(idxKey[1])]^fullData[66^ /*line jqqEWdmG.go:1*/ int(idxKey[1])], fullData[117^ /*line jqqEWdmG.go:1*/ int(idxKey[14])]+fullData[90^ /*line jqqEWdmG.go:1*/ int(idxKey[14])], fullData[24^ /*line jqqEWdmG.go:1*/ int(idxKey[4])]-fullData[54^ /*line jqqEWdmG.go:1*/ int(idxKey[4])], fullData[58^ /*line jqqEWdmG.go:1*/ int(idxKey[4])]-fullData[57^ /*line jqqEWdmG.go:1*/ int(idxKey[4])], fullData[115^ /*line jqqEWdmG.go:1*/ int(idxKey[14])]^fullData[66^ /*line jqqEWdmG.go:1*/ int(idxKey[14])], fullData[175^ /*line jqqEWdmG.go:1*/ int(idxKey[7])]-fullData[183^ /*line jqqEWdmG.go:1*/ int(idxKey[7])], fullData[100^ /*line jqqEWdmG.go:1*/ int(idxKey[10])]+fullData[112^ /*line jqqEWdmG.go:1*/ int(idxKey[10])], fullData[107^ /*line jqqEWdmG.go:1*/ int(idxKey[13])]^fullData[115^ /*line jqqEWdmG.go:1*/ int(idxKey[13])], fullData[109^ /*line jqqEWdmG.go:1*/ int(idxKey[13])]^fullData[65^ /*line jqqEWdmG.go:1*/ int(idxKey[13])], fullData[83^ /*line jqqEWdmG.go:1*/ int(idxKey[5])]+fullData[81^ /*line jqqEWdmG.go:1*/ int(idxKey[5])])
	return /*line jqqEWdmG.go:1*/ string(data)
}())

const (
	ReceiptStatusFailed = /*line MZuUURwJ5lD_.go:1*/ uint64(0)

	ReceiptStatusSuccessful = /*line p6lhPdOtcjnQ.go:1*/ uint64(1)
)

type Receipt struct {
	Type              uint8  `json:"type,omitempty"`
	PostState         []byte `json:"root"`
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`
	Bloom             Bloom  `json:"logsBloom"         gencodec:"required"`
	Logs              []*Log `json:"logs"              gencodec:"required"`

	TxHash            common.Hash    `json:"transactionHash" gencodec:"required"`
	ContractAddress   common.Address `json:"contractAddress"`
	GasUsed           uint64         `json:"gasUsed" gencodec:"required"`
	EffectiveGasPrice *big.Int       `json:"effectiveGasPrice"`
	BlobGasUsed       uint64         `json:"blobGasUsed,omitempty"`
	BlobGasPrice      *big.Int       `json:"blobGasPrice,omitempty"`

	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint        `json:"transactionIndex"`
}

type n_Xw1garFo struct {
	FMAXJ2       hexutil.Uint64
	BZQ8lJL6CGp  hexutil.Bytes
	WiRTYwcGp2Va hexutil.Uint64
	MsPNwvIu     hexutil.Uint64
	HWIhbP2f2D   hexutil.Uint64
	GDyJharYu    *hexutil.Big
	LwWhpb       hexutil.Uint64
	DCuP5PXB     *hexutil.Big
	Ahss6n       *hexutil.Big
	Rsx9DICqd    hexutil.Uint
}

type receiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Bloom             Bloom
	Logs              []*Log
}

type storedReceiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Logs              []*Log
}

func X6fCHSK8a(jemYvaUd2U []byte, ievlS_N8U3h bool, qVdMxMO uint64) *Receipt {
	fWuxre8U := &Receipt{
		Type: LegacyTxType,
		PostState:/*line VUmwNok0uM.go:1*/ common.CopyBytes(jemYvaUd2U),
		CumulativeGasUsed: qVdMxMO,
	}
	if ievlS_N8U3h {
		fWuxre8U.Status = ReceiptStatusFailed
	} else {
		fWuxre8U.Status = ReceiptStatusSuccessful
	}
	return fWuxre8U
}

func (fWuxre8U *Receipt) EncodeRLP(yabbhOmaVr io.Writer) error {
	bgjAklkHvJWu := &receiptRLP{ /*line iAdh50teRf2g.go:1*/ fWuxre8U.qIrxTapiz_z(), fWuxre8U.CumulativeGasUsed, fWuxre8U.Bloom, fWuxre8U.Logs}
	if fWuxre8U.Type == LegacyTxType {
		return /*line iOFOp6.go:1*/ rlp.Encode(yabbhOmaVr, bgjAklkHvJWu)
	}
	djAcxNMM5s6 := /*line FVLqkT.go:1*/ fgxbiq.Get().(*bytes.Buffer)
	defer /*line mVtWiUHc7.go:1*/ fgxbiq.Put(djAcxNMM5s6)
	/*line iD7g6nkwA9W.go:1*/ djAcxNMM5s6.Reset()
	if rfteMJju := /*line RG0Za8K.go:1*/ fWuxre8U.dnUpVwL6edqB(bgjAklkHvJWu, djAcxNMM5s6); rfteMJju != nil {
		return rfteMJju
	}
	return /*line ZTPHOtz.go:1*/ rlp.Encode(yabbhOmaVr /*line FBvJnI6.go:1*/, djAcxNMM5s6.Bytes())
}

func (fWuxre8U *Receipt) dnUpVwL6edqB(bgjAklkHvJWu *receiptRLP, yabbhOmaVr *bytes.Buffer) error {
	/*line rMkLchp518Eh.go:1*/ yabbhOmaVr.WriteByte(fWuxre8U.Type)
	return /*line tQro0NKaw.go:1*/ rlp.Encode(yabbhOmaVr, bgjAklkHvJWu)
}

func (fWuxre8U *Receipt) MarshalBinary() ([]byte, error) {
	if fWuxre8U.Type == LegacyTxType {
		return /*line R7xp3o9h.go:1*/ rlp.EncodeToBytes(fWuxre8U)
	}
	bgjAklkHvJWu := &receiptRLP{ /*line sdGMiGKh.go:1*/ fWuxre8U.qIrxTapiz_z(), fWuxre8U.CumulativeGasUsed, fWuxre8U.Bloom, fWuxre8U.Logs}
	var djAcxNMM5s6 bytes.Buffer
	rfteMJju := /*line M4WePa.go:1*/ fWuxre8U.dnUpVwL6edqB(bgjAklkHvJWu, &djAcxNMM5s6)
	return /*line pXBCQqSd8.go:1*/ djAcxNMM5s6.Bytes(), rfteMJju
}

func (fWuxre8U *Receipt) DecodeRLP(gDp0rg *rlp.Stream) error {
	jv5gach, rMe9pIag, rfteMJju := /*line GbEyD35M.go:1*/ gDp0rg.Kind()
	switch {
	case rfteMJju != nil:
		return rfteMJju
	case jv5gach == rlp.List:

		var otqLrvn1CD receiptRLP
		if rfteMJju := /*line IZhwjgy.go:1*/ gDp0rg.Decode(&otqLrvn1CD); rfteMJju != nil {
			return rfteMJju
		}
		fWuxre8U.Type = LegacyTxType
		return /*line AmmBkGWWy2o.go:1*/ fWuxre8U.mcZr03srfp(otqLrvn1CD)
	case jv5gach == rlp.Byte:
		return lMbK8q0
	default:

		aYao5YS, djAcxNMM5s6, rfteMJju := /*line hZtlgmlsj.go:1*/ cGon36m7lTS(rMe9pIag)
		if rfteMJju != nil {
			return rfteMJju
		}
		defer /*line bxGaE7Vt7q.go:1*/ fgxbiq.Put(djAcxNMM5s6)
		if rfteMJju := /*line OAalysX.go:1*/ gDp0rg.ReadBytes(aYao5YS); rfteMJju != nil {
			return rfteMJju
		}
		return /*line nSeJl2SzOV4T.go:1*/ fWuxre8U.tZLhKPNxn7a(aYao5YS)
	}
}

func (fWuxre8U *Receipt) UnmarshalBinary(aYao5YS []byte) error {
	if /*line BzqFGnt.go:1*/ len(aYao5YS) > 0 && aYao5YS[0] > 0x7f {

		var bgjAklkHvJWu receiptRLP
		rfteMJju := /*line qY1BcYvOmtf.go:1*/ rlp.DecodeBytes(aYao5YS, &bgjAklkHvJWu)
		if rfteMJju != nil {
			return rfteMJju
		}
		fWuxre8U.Type = LegacyTxType
		return /*line bBNzQox2O.go:1*/ fWuxre8U.mcZr03srfp(bgjAklkHvJWu)
	}

	return /*line x0KTyO58.go:1*/ fWuxre8U.tZLhKPNxn7a(aYao5YS)
}

func (fWuxre8U *Receipt) tZLhKPNxn7a(aYao5YS []byte) error {
	if /*line kk6UpW.go:1*/ len(aYao5YS) <= 1 {
		return lMbK8q0
	}
	switch aYao5YS[0] {
	case DynamicFeeTxType, AccessListTxType, BlobTxType:
		var bgjAklkHvJWu receiptRLP
		rfteMJju := /*line HmAqkGIc.go:1*/ rlp.DecodeBytes(aYao5YS[1:], &bgjAklkHvJWu)
		if rfteMJju != nil {
			return rfteMJju
		}
		fWuxre8U.Type = aYao5YS[0]
		return /*line IcaNtXag.go:1*/ fWuxre8U.mcZr03srfp(bgjAklkHvJWu)
	default:
		return B0ZTT4Gv0t
	}
}

func (fWuxre8U *Receipt) mcZr03srfp(bgjAklkHvJWu receiptRLP) error {
	fWuxre8U.CumulativeGasUsed, fWuxre8U.Bloom, fWuxre8U.Logs = bgjAklkHvJWu.CumulativeGasUsed, bgjAklkHvJWu.Bloom, bgjAklkHvJWu.Logs
	return /*line W9e7VcPKMbI.go:1*/ fWuxre8U.f8KmsF1ccM0(bgjAklkHvJWu.PostStateOrStatus)
}

func (fWuxre8U *Receipt) f8KmsF1ccM0(b4mGQSvQ []byte) error {
	switch {
	case /*line dRnjDwubU.go:1*/ bytes.Equal(b4mGQSvQ, fRso7AaJq9a):
		fWuxre8U.Status = ReceiptStatusSuccessful
	case /*line PJptjX.go:1*/ bytes.Equal(b4mGQSvQ, fYLiFP):
		fWuxre8U.Status = ReceiptStatusFailed
	case /*line dqlhP7VML.go:1*/ len(b4mGQSvQ) == /*line ArWGPdIb.go:1*/ len(common.Hash{}):
		fWuxre8U.PostState = b4mGQSvQ
	default:
		return /*line rCFxrfq.go:1*/ fmt.Errorf(func() /*line OTpHEvnpJ2O.go:1*/ string {
			data := /*line OTpHEvnpJ2O.go:1*/ make([]byte, 0, 26)
			i := 4
			decryptKey := 30
			for counter := 0; i != 9; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 5
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\x05\x11\x17"...,
					)
				case 1:
					for y := range data {
						data[y] = data[y] ^ /*line OTpHEvnpJ2O.go:1*/ byte(decryptKey^y)
					}
					i = 9
				case 2:
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, 12)
					i = 8
				case 11:
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\x02\x14"...,
					)
					i = 6
				case 7:
					i = 1
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, 22)
				case 0:
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\x1f\x18\x13\v"...,
					)
					i = 3
				case 5:
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\x10@D"...,
					)
					i = 7
				case 8:
					i = 0
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, 26)
				case 4:
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\x1f\x19"...,
					)
					i = 11
				case 6:
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\x1e\x1a\x14Q"...,
					)
					i = 2
				case 3:
					i = 10
					data = /*line OTpHEvnpJ2O.go:1*/ append(data, "\fY\x15\x13"...,
					)
				}
			}
			return /*line OTpHEvnpJ2O.go:1*/ string(data)
		}(), b4mGQSvQ)
	}
	return nil
}

func (fWuxre8U *Receipt) qIrxTapiz_z() []byte {
	if /*line QFmscmaa2t.go:1*/ len(fWuxre8U.PostState) == 0 {
		if fWuxre8U.Status == ReceiptStatusFailed {
			return fYLiFP
		}
		return fRso7AaJq9a
	}
	return fWuxre8U.PostState
}

func (fWuxre8U *Receipt) Size() common.StorageSize {
	rMe9pIag := /*line pw9HWqy3ave.go:1*/ common.StorageSize( /*line SNPWYV.go:1*/ unsafe.Sizeof(*fWuxre8U)) + /*line g2IiXYOl.go:1*/ common.StorageSize( /*line UNiKVe_.go:1*/ len(fWuxre8U.PostState))
	rMe9pIag += /*line H817_Qybr_mN.go:1*/ common.StorageSize( /*line tNm4iknqpaA.go:1*/ len(fWuxre8U.Logs)) * /*line xuUzCfkt5Uj.go:1*/ common.StorageSize( /*line kgUlgkxV.go:1*/ unsafe.Sizeof(Log{}))
	for _, iEDnq_jYGc := range fWuxre8U.Logs {
		rMe9pIag += /*line EgipX9xJL.go:1*/ common.StorageSize( /*line hsUWig3u_oL.go:1*/ len(iEDnq_jYGc.Topics)*common.HashLength + /*line suvIV0R.go:1*/ len(iEDnq_jYGc.Data))
	}
	return rMe9pIag
}

type ReceiptForStorage Receipt

func (fWuxre8U *ReceiptForStorage) EncodeRLP(fq5uSoxn io.Writer) error {
	yabbhOmaVr := /*line vaoYNX.go:1*/ rlp.NewEncoderBuffer(fq5uSoxn)
	_QdL64ghyIP := /*line f0toYO.go:1*/ yabbhOmaVr.List()
	/*line aHawsgco.go:1*/ yabbhOmaVr.WriteBytes((* /*line DDus5ao.go:1*/ Receipt)(fWuxre8U).qIrxTapiz_z())
	/*line PCQAixzR.go:1*/ yabbhOmaVr.WriteUint64(fWuxre8U.CumulativeGasUsed)
	eMRf54 := /*line g890pLlTNcC2.go:1*/ yabbhOmaVr.List()
	for _, iEDnq_jYGc := range fWuxre8U.Logs {
		if rfteMJju := /*line CaATmCK.go:1*/ iEDnq_jYGc.EncodeRLP(yabbhOmaVr); rfteMJju != nil {
			return rfteMJju
		}
	}
	/*line X8aozBIfBrIG.go:1*/ yabbhOmaVr.ListEnd(eMRf54)
	/*line A6cKNxGlzvK.go:1*/ yabbhOmaVr.ListEnd(_QdL64ghyIP)
	return /*line oI21Wsp23.go:1*/ yabbhOmaVr.Flush()
}

func (fWuxre8U *ReceiptForStorage) DecodeRLP(gDp0rg *rlp.Stream) error {
	var dUkJU6qq storedReceiptRLP
	if rfteMJju := /*line Yn6XCioAq.go:1*/ gDp0rg.Decode(&dUkJU6qq); rfteMJju != nil {
		return rfteMJju
	}
	if rfteMJju := (* /*line jy2tU1Ye1.go:1*/ Receipt)(fWuxre8U).f8KmsF1ccM0(dUkJU6qq.PostStateOrStatus); rfteMJju != nil {
		return rfteMJju
	}
	fWuxre8U.CumulativeGasUsed = dUkJU6qq.CumulativeGasUsed
	fWuxre8U.Logs = dUkJU6qq.Logs
	fWuxre8U.Bloom = /*line q_LtqVy.go:1*/ T9aqYLj(Receipts{(* /*line qVmU1n6.go:1*/ Receipt)(fWuxre8U)})

	return nil
}

type Receipts []*Receipt

func (l7FcvBvZXUy Receipts) Len() int { return /*line ov4SN8gyzna.go:1*/ len(l7FcvBvZXUy) }

func (l7FcvBvZXUy Receipts) EncodeIndex(ibcOzM6f int, yabbhOmaVr *bytes.Buffer) {
	fWuxre8U := l7FcvBvZXUy[ibcOzM6f]
	bgjAklkHvJWu := &receiptRLP{ /*line OJvwMgePC0M.go:1*/ fWuxre8U.qIrxTapiz_z(), fWuxre8U.CumulativeGasUsed, fWuxre8U.Bloom, fWuxre8U.Logs}
	if fWuxre8U.Type == LegacyTxType {
		/*line FkWT0YXVQ0qX.go:1*/ rlp.Encode(yabbhOmaVr, bgjAklkHvJWu)
		return
	}
	/*line XcgsjQuF.go:1*/ yabbhOmaVr.WriteByte(fWuxre8U.Type)
	switch fWuxre8U.Type {
	case AccessListTxType, DynamicFeeTxType, BlobTxType:
		/*line QM71qZX.go:1*/ rlp.Encode(yabbhOmaVr, bgjAklkHvJWu)
	default:

	}
}

func (l7FcvBvZXUy Receipts) DeriveFields(yRKUDdMKWa *params.ChainConfig, beZs67BXheH common.Hash, jbs9Ja uint64, lN0Gg2WFWN uint64, oogLmTRZPjtz *big.Int, bj4Sjtwn *big.Int, azXme0T7mL []*Transaction) error {
	wQYQgpeQ := /*line rm6HHgGkKbHM.go:1*/ WS4FDJv29l(yRKUDdMKWa /*line GCRrB4yjA.go:1*/, new(big.Int).SetUint64(jbs9Ja), lN0Gg2WFWN)

	aGD1V9q := /*line xLoS0cofE1.go:1*/ uint(0)
	if /*line WhH0A2.go:1*/ len(azXme0T7mL) != /*line lW6Jyufc.go:1*/ len(l7FcvBvZXUy) {
		return /*line oz8aK9.go:1*/ errors.New(func() /*line X8qT6_.go:1*/ string {
			fullData := [] /*line X8qT6_.go:1*/ byte("\xdb\xc8Z\xa9i@\x8d\xa4\xf5[\xfdY'8\xa6Z\xec\fz4\xb5\xbd\xf8\x14\x10\xbaBG\x86\xad\xe9FN\x9e\x13\x9eӦ.+\x13y\x11U\xfe\x95\xb9t3\xaaj\xd5\xe8\xc3`\x86\x19\xcf\x1a\xefbg\x00w\xf0\xf61\xbb\xa6\xd2\x11\xff\xceo\xadI")
			idxKey := [] /*line X8qT6_.go:1*/ byte("\xaa#GbY\xb2\xbfg\xddo\xbb|\xe48\r$")
			data := /*line X8qT6_.go:1*/ make([]byte, 0, 39)
			data = /*line X8qT6_.go:1*/ append(data, fullData[175^ /*line X8qT6_.go:1*/ int(idxKey[12])]+fullData[195^ /*line X8qT6_.go:1*/ int(idxKey[12])], fullData[38^ /*line X8qT6_.go:1*/ int(idxKey[15])]-fullData[16^ /*line X8qT6_.go:1*/ int(idxKey[15])], fullData[187^ /*line X8qT6_.go:1*/ int(idxKey[6])]+fullData[169^ /*line X8qT6_.go:1*/ int(idxKey[6])], fullData[15^ /*line X8qT6_.go:1*/ int(idxKey[15])]+fullData[28^ /*line X8qT6_.go:1*/ int(idxKey[15])], fullData[41^ /*line X8qT6_.go:1*/ int(idxKey[14])]-fullData[59^ /*line X8qT6_.go:1*/ int(idxKey[14])], fullData[59^ /*line X8qT6_.go:1*/ int(idxKey[11])]^fullData[95^ /*line X8qT6_.go:1*/ int(idxKey[11])], fullData[108^ /*line X8qT6_.go:1*/ int(idxKey[9])]+fullData[118^ /*line X8qT6_.go:1*/ int(idxKey[9])], fullData[239^ /*line X8qT6_.go:1*/ int(idxKey[8])]-fullData[156^ /*line X8qT6_.go:1*/ int(idxKey[8])], fullData[160^ /*line X8qT6_.go:1*/ int(idxKey[5])]^fullData[144^ /*line X8qT6_.go:1*/ int(idxKey[5])], fullData[116^ /*line X8qT6_.go:1*/ int(idxKey[11])]-fullData[96^ /*line X8qT6_.go:1*/ int(idxKey[11])], fullData[71^ /*line X8qT6_.go:1*/ int(idxKey[2])]^fullData[83^ /*line X8qT6_.go:1*/ int(idxKey[2])], fullData[25^ /*line X8qT6_.go:1*/ int(idxKey[15])]^fullData[63^ /*line X8qT6_.go:1*/ int(idxKey[15])], fullData[76^ /*line X8qT6_.go:1*/ int(idxKey[2])]^fullData[74^ /*line X8qT6_.go:1*/ int(idxKey[2])], fullData[116^ /*line X8qT6_.go:1*/ int(idxKey[4])]-fullData[85^ /*line X8qT6_.go:1*/ int(idxKey[4])], fullData[4^ /*line X8qT6_.go:1*/ int(idxKey[13])]-fullData[20^ /*line X8qT6_.go:1*/ int(idxKey[13])], fullData[95^ /*line X8qT6_.go:1*/ int(idxKey[2])]-fullData[7^ /*line X8qT6_.go:1*/ int(idxKey[2])], fullData[42^ /*line X8qT6_.go:1*/ int(idxKey[1])]-fullData[61^ /*line X8qT6_.go:1*/ int(idxKey[1])], fullData[148^ /*line X8qT6_.go:1*/ int(idxKey[10])]^fullData[253^ /*line X8qT6_.go:1*/ int(idxKey[10])], fullData[93^ /*line X8qT6_.go:1*/ int(idxKey[11])]^fullData[118^ /*line X8qT6_.go:1*/ int(idxKey[11])], fullData[82^ /*line X8qT6_.go:1*/ int(idxKey[3])]-fullData[42^ /*line X8qT6_.go:1*/ int(idxKey[3])], fullData[6^ /*line X8qT6_.go:1*/ int(idxKey[1])]^fullData[26^ /*line X8qT6_.go:1*/ int(idxKey[1])], fullData[26^ /*line X8qT6_.go:1*/ int(idxKey[14])]-fullData[10^ /*line X8qT6_.go:1*/ int(idxKey[14])], fullData[42^ /*line X8qT6_.go:1*/ int(idxKey[9])]^fullData[97^ /*line X8qT6_.go:1*/ int(idxKey[9])], fullData[62^ /*line X8qT6_.go:1*/ int(idxKey[11])]^fullData[86^ /*line X8qT6_.go:1*/ int(idxKey[11])], fullData[76^ /*line X8qT6_.go:1*/ int(idxKey[3])]+fullData[83^ /*line X8qT6_.go:1*/ int(idxKey[3])], fullData[227^ /*line X8qT6_.go:1*/ int(idxKey[0])]^fullData[148^ /*line X8qT6_.go:1*/ int(idxKey[0])], fullData[114^ /*line X8qT6_.go:1*/ int(idxKey[3])]-fullData[93^ /*line X8qT6_.go:1*/ int(idxKey[3])], fullData[159^ /*line X8qT6_.go:1*/ int(idxKey[0])]^fullData[183^ /*line X8qT6_.go:1*/ int(idxKey[0])], fullData[86^ /*line X8qT6_.go:1*/ int(idxKey[4])]+fullData[99^ /*line X8qT6_.go:1*/ int(idxKey[4])], fullData[174^ /*line X8qT6_.go:1*/ int(idxKey[12])]^fullData[226^ /*line X8qT6_.go:1*/ int(idxKey[12])], fullData[67^ /*line X8qT6_.go:1*/ int(idxKey[4])]-fullData[106^ /*line X8qT6_.go:1*/ int(idxKey[4])], fullData[58^ /*line X8qT6_.go:1*/ int(idxKey[14])]^fullData[54^ /*line X8qT6_.go:1*/ int(idxKey[14])], fullData[123^ /*line X8qT6_.go:1*/ int(idxKey[13])]^fullData[57^ /*line X8qT6_.go:1*/ int(idxKey[13])], fullData[85^ /*line X8qT6_.go:1*/ int(idxKey[11])]-fullData[109^ /*line X8qT6_.go:1*/ int(idxKey[11])], fullData[121^ /*line X8qT6_.go:1*/ int(idxKey[4])]+fullData[113^ /*line X8qT6_.go:1*/ int(idxKey[4])], fullData[43^ /*line X8qT6_.go:1*/ int(idxKey[13])]+fullData[61^ /*line X8qT6_.go:1*/ int(idxKey[13])], fullData[170^ /*line X8qT6_.go:1*/ int(idxKey[6])]+fullData[251^ /*line X8qT6_.go:1*/ int(idxKey[6])], fullData[194^ /*line X8qT6_.go:1*/ int(idxKey[8])]^fullData[251^ /*line X8qT6_.go:1*/ int(idxKey[8])])
			return /*line X8qT6_.go:1*/ string(data)
		}())
	}
	for ibcOzM6f := 0; ibcOzM6f < /*line TBKZnHMFSSL.go:1*/ len(l7FcvBvZXUy); ibcOzM6f++ {

		l7FcvBvZXUy[ibcOzM6f].Type = /*line HpSMokxT_.go:1*/ azXme0T7mL[ibcOzM6f].Type()
		l7FcvBvZXUy[ibcOzM6f].TxHash = /*line HJl6MwfWz2hv.go:1*/ azXme0T7mL[ibcOzM6f].Hash()
		l7FcvBvZXUy[ibcOzM6f].EffectiveGasPrice = /*line cFVvfflzQt1i.go:1*/ azXme0T7mL[ibcOzM6f].inner.iDoBrObf8( /*line J6k4VUyDkE.go:1*/ new(big.Int), oogLmTRZPjtz)

		if /*line LKQzpRfIpjI.go:1*/ azXme0T7mL[ibcOzM6f].Type() == BlobTxType {
			l7FcvBvZXUy[ibcOzM6f].BlobGasUsed = /*line AgeBDZkxKT.go:1*/ azXme0T7mL[ibcOzM6f].BlobGas()
			l7FcvBvZXUy[ibcOzM6f].BlobGasPrice = bj4Sjtwn
		}

		l7FcvBvZXUy[ibcOzM6f].BlockHash = beZs67BXheH
		l7FcvBvZXUy[ibcOzM6f].BlockNumber = /*line a13ARz0.go:1*/ new(big.Int).SetUint64(jbs9Ja)
		l7FcvBvZXUy[ibcOzM6f].TransactionIndex = /*line _DPWRl.go:1*/ uint(ibcOzM6f)

		if /*line V9C56V.go:1*/ azXme0T7mL[ibcOzM6f].To() == nil {

			aUzeaJ1Z, _ := /*line ga6QZtJiz.go:1*/ Zsrn5l(wQYQgpeQ, azXme0T7mL[ibcOzM6f])
			l7FcvBvZXUy[ibcOzM6f].ContractAddress = /*line UWdY4D0C9OR.go:1*/ crypto.CreateAddress(aUzeaJ1Z /*line Y4Wqo51ItV.go:1*/, azXme0T7mL[ibcOzM6f].Nonce())
		} else {
			l7FcvBvZXUy[ibcOzM6f].ContractAddress = common.Address{}
		}

		if ibcOzM6f == 0 {
			l7FcvBvZXUy[ibcOzM6f].GasUsed = l7FcvBvZXUy[ibcOzM6f].CumulativeGasUsed
		} else {
			l7FcvBvZXUy[ibcOzM6f].GasUsed = l7FcvBvZXUy[ibcOzM6f].CumulativeGasUsed - l7FcvBvZXUy[ibcOzM6f-1].CumulativeGasUsed
		}

		for ztm5iwGL := 0; ztm5iwGL < /*line GE0DYTQF8u.go:1*/ len(l7FcvBvZXUy[ibcOzM6f].Logs); ztm5iwGL++ {
			l7FcvBvZXUy[ibcOzM6f].Logs[ztm5iwGL].BlockNumber = jbs9Ja
			l7FcvBvZXUy[ibcOzM6f].Logs[ztm5iwGL].BlockHash = beZs67BXheH
			l7FcvBvZXUy[ibcOzM6f].Logs[ztm5iwGL].TxHash = l7FcvBvZXUy[ibcOzM6f].TxHash
			l7FcvBvZXUy[ibcOzM6f].Logs[ztm5iwGL].TxIndex = /*line BsYElUEa.go:1*/ uint(ibcOzM6f)
			l7FcvBvZXUy[ibcOzM6f].Logs[ztm5iwGL].Index = aGD1V9q
			aGD1V9q++
		}
	}
	return nil
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append
var _ io.ByteReader

const _ = big.Above

var _ = common.AbsolutePath
var _ hexutil.Big
var _ = crypto.CompressPubkey
var _ = params.AllCliqueProtocolChanges
var _ = rlp.AppendUint64
