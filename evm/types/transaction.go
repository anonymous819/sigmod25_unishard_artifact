//line :1
package types

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	CqKl3D50bxiN = /*line L6VzPd460_lw.go:1*/ errors.New(func() /*line EY0QYAjvsK.go:1*/ string {
		data := /*line EY0QYAjvsK.go:1*/ make([]byte, 0, 35)
		i := 16
		decryptKey := 87
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				data = /*line EY0QYAjvsK.go:1*/ append(data, 239)
				i = 10
			case 2:
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xb4\xf1\xaa\xad"...,
				)
				i = 8
			case 10:
				i = 7
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xe0\xb3\xe2\xef"...,
				)
			case 7:
				i = 0
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xbf\xed\xbc\xaa"...,
				)
			case 8:
				data = /*line EY0QYAjvsK.go:1*/ append(data, 189)
				i = 4
			case 1:
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xbe\xba"...,
				)
				i = 2
			case 13:
				i = 1
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xb9\xa2\xb4"...,
				)
			case 9:
				for y := range data {
					data[y] = data[y] ^ /*line EY0QYAjvsK.go:1*/ byte(decryptKey^y)
				}
				i = 6
			case 0:
				i = 15
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xa4\xbc"...,
				)
			case 16:
				data = /*line EY0QYAjvsK.go:1*/ append(data, 191)
				i = 13
			case 11:
				i = 17
				data = /*line EY0QYAjvsK.go:1*/ append(data, 168)
			case 14:
				i = 12
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xbb\xad"...,
				)
			case 3:
				data = /*line EY0QYAjvsK.go:1*/ append(data, 186)
				i = 14
			case 4:
				i = 3
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xb3\xa9"...,
				)
			case 17:
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\xaa\xe5\xb4"...,
				)
				i = 5
			case 15:
				data = /*line EY0QYAjvsK.go:1*/ append(data, "\x93\x84"...,
				)
				i = 9
			case 12:
				data = /*line EY0QYAjvsK.go:1*/ append(data, 175)
				i = 11
			}
		}
		return /*line EY0QYAjvsK.go:1*/ string(data)
	}())
	A4ZTcBQ = /*line IvShi5.go:1*/ errors.New(func() /*line IcvF9u5IA.go:1*/ string {
		seed := /*line IcvF9u5IA.go:1*/ byte(64)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line IcvF9u5IA.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line IcvF9u5IA.go:1*/ fnc(180)(102)(187)(131)(11)(4)(10)(37)(63)(132)(7)(192)(212)(173)(81)(151)(233)(22)(55)(100)(214)(89)(0)(1)(7)(186)(199)(144)(27)(54)(107)(217)(180)(89)(177)(30)(97)(198)(147)(3)(10)(24)(48)(75)(230)(206)(153)(55)(95)(188)(137)(3)(5)(198)(223)(180)(102)(211)(153)(69)(139)(19)(25)(64)
		return /*line IcvF9u5IA.go:1*/ string(data)
	}())
	ASX0AwZr = /*line AjStQ6mUD.go:1*/ errors.New(func() /*line HZHqgfR.go:1*/ string {
		key := [] /*line HZHqgfR.go:1*/ byte("ޤh\xb7\x9dI\x00\xcb\xefMȐw\xf1\xf7r\xbb\x8dfn\x13\x7f\xaa\xaa\a\x15T<(\x8d\xf7\xa3\x8b\xdc\xfe\xe0'\xd0\xc4\xec\xeb\xa3")
		data := [] /*line HZHqgfR.go:1*/ byte("\xaa\xd6\t\xd9\xee(c\xbf\x86\"\xa6\xb0\x03\x88\x87\x17\x9b\xe3\t\x1a3\t\xcb\xc6nqtUF\xad\x83\xcb\xe2\xafރH\xbe\xb0\x89\x93\xd7")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line HZHqgfR.go:1*/ string(data)
	}())
	B0ZTT4Gv0t = /*line OhKGcTvmBm.go:1*/ errors.New(func() /*line Drc1Yb47Eyal.go:1*/ string {
		key := [] /*line Drc1Yb47Eyal.go:1*/ byte("6,\x9d\x8eP5\x89\x8f\xd1\xe3@\xb7*\xf1r\xd6\xfdF\x0f\xcd\xd4b\x9c\x9d\xc9\x12\xafB\x15\xe7")
		data := [] /*line Drc1Yb47Eyal.go:1*/ byte("B^\xfc\xe0#T\xea\xfb\xb8\x8c.\x97^\x88\x02\xb3\xdd(`\xb9\xf4\x11\xe9\xed\xb9}\xdd6p\x83")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line Drc1Yb47Eyal.go:1*/ string(data)
	}())
	ZLykSo = /*line hY_p8KZN.go:1*/ errors.New(func() /*line s0Z8VTX8M.go:1*/ string {
		data := [] /*line s0Z8VTX8M.go:1*/ byte("fe\xdd Za4-v\xd0;s t\xec\xc8V\xc6ba\xd9\xd2$\xe9e\xc3")
		positions := [...]byte{14, 14, 23, 16, 4, 2, 10, 21, 22, 6, 8, 22, 21, 23, 9, 2, 9, 22, 6, 25, 15, 20, 10, 17, 15, 7, 15, 9}
		for i := 0; i < 28; i += 2 {
			localKey := /*line s0Z8VTX8M.go:1*/ byte(i) + /*line s0Z8VTX8M.go:1*/ byte(positions[i]^positions[i+1]) + 124
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line s0Z8VTX8M.go:1*/ string(data)
	}())
	xVgRUgbb9tPH = /*line _8cG7aBzP.go:1*/ errors.New(func() /*line pLJg0x2.go:1*/ string {
		data := [] /*line pLJg0x2.go:1*/ byte("q\xd8xe\xc2\xf0\x95\x00az\xa0\x13\x04tio\xb1\xf0\x18e\xdax\xbe\xc4or\xbc")
		positions := [...]byte{19, 10, 1, 10, 21, 6, 6, 12, 2, 11, 23, 20, 9, 7, 16, 7, 26, 0, 6, 19, 9, 23, 6, 4, 11, 11, 18, 16, 9, 18, 5, 17, 22, 26, 16, 2}
		for i := 0; i < 36; i += 2 {
			localKey := /*line pLJg0x2.go:1*/ byte(i) + /*line pLJg0x2.go:1*/ byte(positions[i]^positions[i+1]) + 158
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line pLJg0x2.go:1*/ string(data)
	}())
	vnbtviuT = /*line pF8aDWeb.go:1*/ errors.New(func() /*line FQAYiITp7Ogm.go:1*/ string {
		data := [] /*line FQAYiITp7Ogm.go:1*/ byte("wyP\xda\xdfkm\x9e'x~s\x8flߍ.u\x9ea b`\xc70 Ægk")
		positions := [...]byte{7, 18, 14, 23, 26, 27, 10, 22, 3, 12, 29, 7, 3, 19, 18, 0, 4, 15, 29, 14, 9, 27, 12, 0, 22, 6, 11, 7, 18, 18, 16, 28, 14, 5, 16, 22}
		for i := 0; i < 36; i += 2 {
			localKey := /*line FQAYiITp7Ogm.go:1*/ byte(i) + /*line FQAYiITp7Ogm.go:1*/ byte(positions[i]^positions[i+1]) + 228
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line FQAYiITp7Ogm.go:1*/ string(data)
	}())
	sn89ktn_dI = /*line wJGAHGv.go:1*/ errors.New(func() /*line IiiA7DbizKRm.go:1*/ string {
		key := [] /*line IiiA7DbizKRm.go:1*/ byte("\x1b'x\xdbHH\xec\xe5\xfe\x19*J\x9c\x1c\xden\xa9\xf2\xf9ff\x8c\x80\xc3Ġd0A&&\xfa<\x14\xc0\xc4i")
		data := [] /*line IiiA7DbizKRm.go:1*/ byte("B\x9d\x9f\xfb\xa9\xb6P\x05%\x92z\xab\x0e\x85R\xe7\xd0\x12_\xcf\xcb\xf8\xe46\xe4\x04\xd3P\xaf\x95\x9a\x1a\xa9u4'\xd1")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line IiiA7DbizKRm.go:1*/ string(data)
	}())
	tH2RXxQ = /*line jVhlFHSMX.go:1*/ errors.New(func() /*line DA1aIt.go:1*/ string {
		fullData := [] /*line DA1aIt.go:1*/ byte("\x13\xf1\xaf\x93%(\x14\x8c\bR\x15\xc4ќ\xc6\xe2$\xb9NF| \x95\x15q\x88\x9b\xe5\x9eg\xabR\x05\x1c\x8aM˙\xb8\xcacB`S\x03\xa2Y\xea\x0e\xed\x83\xd0c\x96y\f\xa5A\\\x06va\xfc\tQS#\x12`\xa4\x96\x1c.\xc2%\xa2EYأ<\xf7\xb7\xdcr[\xb6,\xdfH")
		idxKey := [] /*line DA1aIt.go:1*/ byte("\x056\xcf=]Kz{\xda{dr\x0f\xc5^\x93")
		data := /*line DA1aIt.go:1*/ make([]byte, 0, 46)
		data = /*line DA1aIt.go:1*/ append(data, fullData[44^ /*line DA1aIt.go:1*/ int(idxKey[7])]^fullData[66^ /*line DA1aIt.go:1*/ int(idxKey[7])], fullData[85^ /*line DA1aIt.go:1*/ int(idxKey[5])]-fullData[98^ /*line DA1aIt.go:1*/ int(idxKey[5])], fullData[99^ /*line DA1aIt.go:1*/ int(idxKey[11])]^fullData[85^ /*line DA1aIt.go:1*/ int(idxKey[11])], fullData[107^ /*line DA1aIt.go:1*/ int(idxKey[4])]-fullData[102^ /*line DA1aIt.go:1*/ int(idxKey[4])], fullData[116^ /*line DA1aIt.go:1*/ int(idxKey[14])]+fullData[97^ /*line DA1aIt.go:1*/ int(idxKey[14])], fullData[207^ /*line DA1aIt.go:1*/ int(idxKey[13])]+fullData[235^ /*line DA1aIt.go:1*/ int(idxKey[13])], fullData[145^ /*line DA1aIt.go:1*/ int(idxKey[15])]-fullData[202^ /*line DA1aIt.go:1*/ int(idxKey[15])], fullData[56^ /*line DA1aIt.go:1*/ int(idxKey[12])]+fullData[9^ /*line DA1aIt.go:1*/ int(idxKey[12])], fullData[88^ /*line DA1aIt.go:1*/ int(idxKey[6])]-fullData[78^ /*line DA1aIt.go:1*/ int(idxKey[6])], fullData[67^ /*line DA1aIt.go:1*/ int(idxKey[14])]+fullData[29^ /*line DA1aIt.go:1*/ int(idxKey[14])], fullData[109^ /*line DA1aIt.go:1*/ int(idxKey[7])]-fullData[55^ /*line DA1aIt.go:1*/ int(idxKey[7])], fullData[122^ /*line DA1aIt.go:1*/ int(idxKey[11])]+fullData[63^ /*line DA1aIt.go:1*/ int(idxKey[11])], fullData[16^ /*line DA1aIt.go:1*/ int(idxKey[0])]^fullData[26^ /*line DA1aIt.go:1*/ int(idxKey[0])], fullData[194^ /*line DA1aIt.go:1*/ int(idxKey[13])]^fullData[222^ /*line DA1aIt.go:1*/ int(idxKey[13])], fullData[156^ /*line DA1aIt.go:1*/ int(idxKey[15])]^fullData[166^ /*line DA1aIt.go:1*/ int(idxKey[15])], fullData[65^ /*line DA1aIt.go:1*/ int(idxKey[4])]-fullData[23^ /*line DA1aIt.go:1*/ int(idxKey[4])], fullData[83^ /*line DA1aIt.go:1*/ int(idxKey[7])]-fullData[43^ /*line DA1aIt.go:1*/ int(idxKey[7])], fullData[52^ /*line DA1aIt.go:1*/ int(idxKey[11])]-fullData[78^ /*line DA1aIt.go:1*/ int(idxKey[11])], fullData[57^ /*line DA1aIt.go:1*/ int(idxKey[3])]-fullData[107^ /*line DA1aIt.go:1*/ int(idxKey[3])], fullData[65^ /*line DA1aIt.go:1*/ int(idxKey[11])]+fullData[95^ /*line DA1aIt.go:1*/ int(idxKey[11])], fullData[219^ /*line DA1aIt.go:1*/ int(idxKey[8])]-fullData[214^ /*line DA1aIt.go:1*/ int(idxKey[8])], fullData[28^ /*line DA1aIt.go:1*/ int(idxKey[0])]-fullData[56^ /*line DA1aIt.go:1*/ int(idxKey[0])], fullData[66^ /*line DA1aIt.go:1*/ int(idxKey[5])]^fullData[91^ /*line DA1aIt.go:1*/ int(idxKey[5])], fullData[60^ /*line DA1aIt.go:1*/ int(idxKey[10])]-fullData[66^ /*line DA1aIt.go:1*/ int(idxKey[10])], fullData[28^ /*line DA1aIt.go:1*/ int(idxKey[14])]-fullData[114^ /*line DA1aIt.go:1*/ int(idxKey[14])], fullData[64^ /*line DA1aIt.go:1*/ int(idxKey[0])]^fullData[76^ /*line DA1aIt.go:1*/ int(idxKey[0])], fullData[209^ /*line DA1aIt.go:1*/ int(idxKey[8])]-fullData[143^ /*line DA1aIt.go:1*/ int(idxKey[8])], fullData[46^ /*line DA1aIt.go:1*/ int(idxKey[12])]-fullData[93^ /*line DA1aIt.go:1*/ int(idxKey[12])], fullData[53^ /*line DA1aIt.go:1*/ int(idxKey[0])]-fullData[78^ /*line DA1aIt.go:1*/ int(idxKey[0])], fullData[12^ /*line DA1aIt.go:1*/ int(idxKey[12])]^fullData[94^ /*line DA1aIt.go:1*/ int(idxKey[12])], fullData[83^ /*line DA1aIt.go:1*/ int(idxKey[5])]^fullData[11^ /*line DA1aIt.go:1*/ int(idxKey[5])], fullData[92^ /*line DA1aIt.go:1*/ int(idxKey[5])]^fullData[95^ /*line DA1aIt.go:1*/ int(idxKey[5])], fullData[97^ /*line DA1aIt.go:1*/ int(idxKey[11])]+fullData[119^ /*line DA1aIt.go:1*/ int(idxKey[11])], fullData[114^ /*line DA1aIt.go:1*/ int(idxKey[3])]^fullData[15^ /*line DA1aIt.go:1*/ int(idxKey[3])], fullData[120^ /*line DA1aIt.go:1*/ int(idxKey[1])]+fullData[59^ /*line DA1aIt.go:1*/ int(idxKey[1])], fullData[23^ /*line DA1aIt.go:1*/ int(idxKey[0])]-fullData[86^ /*line DA1aIt.go:1*/ int(idxKey[0])], fullData[255^ /*line DA1aIt.go:1*/ int(idxKey[13])]+fullData[229^ /*line DA1aIt.go:1*/ int(idxKey[13])], fullData[111^ /*line DA1aIt.go:1*/ int(idxKey[5])]^fullData[115^ /*line DA1aIt.go:1*/ int(idxKey[5])], fullData[5^ /*line DA1aIt.go:1*/ int(idxKey[0])]^fullData[65^ /*line DA1aIt.go:1*/ int(idxKey[0])], fullData[81^ /*line DA1aIt.go:1*/ int(idxKey[5])]+fullData[69^ /*line DA1aIt.go:1*/ int(idxKey[5])], fullData[71^ /*line DA1aIt.go:1*/ int(idxKey[10])]^fullData[44^ /*line DA1aIt.go:1*/ int(idxKey[10])], fullData[244^ /*line DA1aIt.go:1*/ int(idxKey[13])]^fullData[224^ /*line DA1aIt.go:1*/ int(idxKey[13])], fullData[241^ /*line DA1aIt.go:1*/ int(idxKey[8])]-fullData[245^ /*line DA1aIt.go:1*/ int(idxKey[8])], fullData[136^ /*line DA1aIt.go:1*/ int(idxKey[2])]+fullData[142^ /*line DA1aIt.go:1*/ int(idxKey[2])], fullData[142^ /*line DA1aIt.go:1*/ int(idxKey[8])]+fullData[228^ /*line DA1aIt.go:1*/ int(idxKey[8])])
		return /*line DA1aIt.go:1*/ string(data)
	}())
)

const (
	LegacyTxType     = 0x00
	AccessListTxType = 0x01
	DynamicFeeTxType = 0x02
	BlobTxType       = 0x03
)

type Transaction struct {
	inner TxData
	time  time.Time

	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

func CV6bt80q7Gn(vv4w9F TxData) *Transaction {
	iPfjW8i91hC := /*line af3wAe7QQG.go:1*/ new(Transaction)
	/*line vh4sMRSDr.go:1*/ iPfjW8i91hC.secbe_rMaO( /*line dtzPbJ.go:1*/ vv4w9F.acVy5yir(), 0)
	return iPfjW8i91hC
}

type TxData interface {
	b1pCNt1O() byte
	acVy5yir() TxData

	aVVDgwu() *big.Int
	mJq92sCJCrD() AccessList
	bgjAklkHvJWu() []byte
	go6R4MnIedw() uint64
	dtbQrx() *big.Int
	wIXPaTO() *big.Int
	nwO6ag44() *big.Int
	fMqtswJ() *big.Int
	xiCJb__() uint64
	ypKkxrFoZ() *common.Address

	naFUnaJ8Fe() (pW1UMRn1s, fWuxre8U, gDp0rg *big.Int)
	ccFa3ozV(aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg *big.Int)

	iDoBrObf8(ySxt2pzafw *big.Int, oogLmTRZPjtz *big.Int) *big.Int

	iJPSdycNFb(*bytes.Buffer) error
	h5OeLZu4u([]byte) error
}

func (iPfjW8i91hC *Transaction) EncodeRLP(yabbhOmaVr io.Writer) error {
	if /*line HLGUivD_.go:1*/ iPfjW8i91hC.Type() == LegacyTxType {
		return /*line AyaxrE.go:1*/ rlp.Encode(yabbhOmaVr, iPfjW8i91hC.inner)
	}

	djAcxNMM5s6 := /*line UxpRvRby50W.go:1*/ fgxbiq.Get().(*bytes.Buffer)
	defer /*line ggybUGreqFw.go:1*/ fgxbiq.Put(djAcxNMM5s6)
	/*line Yt4fwnwfYS.go:1*/ djAcxNMM5s6.Reset()
	if rfteMJju := /*line TvHsn6xS.go:1*/ iPfjW8i91hC.dnUpVwL6edqB(djAcxNMM5s6); rfteMJju != nil {
		return rfteMJju
	}
	return /*line wqvZ653Q.go:1*/ rlp.Encode(yabbhOmaVr /*line DivqZaew.go:1*/, djAcxNMM5s6.Bytes())
}

func (iPfjW8i91hC *Transaction) dnUpVwL6edqB(yabbhOmaVr *bytes.Buffer) error {
	/*line K8nwE0qBM.go:1*/ yabbhOmaVr.WriteByte( /*line HtXpQrsn.go:1*/ iPfjW8i91hC.Type())
	return /*line x7PpxZB.go:1*/ iPfjW8i91hC.inner.iJPSdycNFb(yabbhOmaVr)
}

func (iPfjW8i91hC *Transaction) MarshalBinary() ([]byte, error) {
	if /*line A7_qyexFMa.go:1*/ iPfjW8i91hC.Type() == LegacyTxType {
		return /*line DoZXdL.go:1*/ rlp.EncodeToBytes(iPfjW8i91hC.inner)
	}
	var djAcxNMM5s6 bytes.Buffer
	rfteMJju := /*line yza05vKz.go:1*/ iPfjW8i91hC.dnUpVwL6edqB(&djAcxNMM5s6)
	return /*line D2dxDK.go:1*/ djAcxNMM5s6.Bytes(), rfteMJju
}

func (iPfjW8i91hC *Transaction) DecodeRLP(gDp0rg *rlp.Stream) error {
	jv5gach, rMe9pIag, rfteMJju := /*line EILqN0.go:1*/ gDp0rg.Kind()
	switch {
	case rfteMJju != nil:
		return rfteMJju
	case jv5gach == rlp.List:

		var vv4w9F LegacyTx
		rfteMJju := /*line hoaxp8e.go:1*/ gDp0rg.Decode(&vv4w9F)
		if rfteMJju == nil {
			/*line _vauyd6XZ_.go:1*/ iPfjW8i91hC.secbe_rMaO(&vv4w9F /*line S0AYI7.go:1*/, rlp.ListSize(rMe9pIag))
		}
		return rfteMJju
	case jv5gach == rlp.Byte:
		return xVgRUgbb9tPH
	default:

		aYao5YS, djAcxNMM5s6, rfteMJju := /*line vA3a2OrQM.go:1*/ cGon36m7lTS(rMe9pIag)
		if rfteMJju != nil {
			return rfteMJju
		}
		defer /*line RZXs6M.go:1*/ fgxbiq.Put(djAcxNMM5s6)
		if rfteMJju := /*line UwdgZeOlv4QO.go:1*/ gDp0rg.ReadBytes(aYao5YS); rfteMJju != nil {
			return rfteMJju
		}

		vv4w9F, rfteMJju := /*line sPN_31wZ.go:1*/ iPfjW8i91hC.tZLhKPNxn7a(aYao5YS)
		if rfteMJju == nil {
			/*line wh3auq70dMH1.go:1*/ iPfjW8i91hC.secbe_rMaO(vv4w9F, rMe9pIag)
		}
		return rfteMJju
	}
}

func (iPfjW8i91hC *Transaction) UnmarshalBinary(aYao5YS []byte) error {
	if /*line GFfXkEaY0VM.go:1*/ len(aYao5YS) > 0 && aYao5YS[0] > 0x7f {

		var bgjAklkHvJWu LegacyTx
		rfteMJju := /*line be1p_VnSPN.go:1*/ rlp.DecodeBytes(aYao5YS, &bgjAklkHvJWu)
		if rfteMJju != nil {
			return rfteMJju
		}
		/*line bR_u_zw61.go:1*/ iPfjW8i91hC.secbe_rMaO(&bgjAklkHvJWu /*line mWntw08lm64.go:1*/, uint64( /*line EhpAeK.go:1*/ len(aYao5YS)))
		return nil
	}

	vv4w9F, rfteMJju := /*line gSghnnj.go:1*/ iPfjW8i91hC.tZLhKPNxn7a(aYao5YS)
	if rfteMJju != nil {
		return rfteMJju
	}
	/*line IONaHl8jw.go:1*/ iPfjW8i91hC.secbe_rMaO(vv4w9F /*line KTai8ihcmF.go:1*/, uint64( /*line mtFsOn.go:1*/ len(aYao5YS)))
	return nil
}

func (iPfjW8i91hC *Transaction) tZLhKPNxn7a(aYao5YS []byte) (TxData, error) {
	if /*line H0fhJQV2.go:1*/ len(aYao5YS) <= 1 {
		return nil, xVgRUgbb9tPH
	}
	var vv4w9F TxData
	switch aYao5YS[0] {
	case AccessListTxType:
		vv4w9F = /*line OlMQI3Dg1PeO.go:1*/ new(AccessListTx)
	case DynamicFeeTxType:
		vv4w9F = /*line wPd_6QlwL.go:1*/ new(DynamicFeeTx)
	case BlobTxType:
		vv4w9F = /*line q5sM4EEk.go:1*/ new(BlobTx)
	default:
		return nil, B0ZTT4Gv0t
	}
	rfteMJju := /*line koAl9jBt.go:1*/ vv4w9F.h5OeLZu4u(aYao5YS[1:])
	return vv4w9F, rfteMJju
}

func (iPfjW8i91hC *Transaction) secbe_rMaO(vv4w9F TxData, rMe9pIag uint64) {
	iPfjW8i91hC.inner = vv4w9F
	iPfjW8i91hC.time = /*line sMm4__FPQErW.go:1*/ time.Now()
	if rMe9pIag > 0 {
		/*line BpXTDGH2.go:1*/ iPfjW8i91hC.size.Store(rMe9pIag)
	}
}

func vwXAfyu(pW1UMRn1s *big.Int, fWuxre8U *big.Int, gDp0rg *big.Int, ubWxiPM64Wj bool) error {
	if /*line CvfzIbT.go:1*/ eHSrmIkThFCt(pW1UMRn1s) && !ubWxiPM64Wj {
		return A4ZTcBQ
	}

	var xbLDnkMEucK byte
	if /*line D8lOGAlQFIl.go:1*/ eHSrmIkThFCt(pW1UMRn1s) {
		aVVDgwu := /*line tk_smnsp.go:1*/ sikVyPV4yp(pW1UMRn1s).Uint64()
		xbLDnkMEucK = /*line UdKAJ2tXE.go:1*/ byte( /*line rGZOXfdZs.go:1*/ pW1UMRn1s.Uint64() - 35 - 2*aVVDgwu)
	} else if ubWxiPM64Wj {

		xbLDnkMEucK = /*line BIfZaH9s.go:1*/ byte( /*line jEAubKsDghR.go:1*/ pW1UMRn1s.Uint64() - 27)
	} else {

		xbLDnkMEucK = /*line KJCy6dDjzD.go:1*/ byte( /*line lKAYXWYnEK.go:1*/ pW1UMRn1s.Uint64())
	}
	if ! /*line h83_IH.go:1*/ crypto.ValidateSignatureValues(xbLDnkMEucK, fWuxre8U, gDp0rg, false) {
		return CqKl3D50bxiN
	}

	return nil
}

func eHSrmIkThFCt(ICKLPq3pU *big.Int) bool {
	if /*line Z60w6mGIaG4O.go:1*/ ICKLPq3pU.BitLen() <= 8 {
		pW1UMRn1s := /*line mo3dpXHu6Vbw.go:1*/ ICKLPq3pU.Uint64()
		return pW1UMRn1s != 27 && pW1UMRn1s != 28 && pW1UMRn1s != 1 && pW1UMRn1s != 0
	}

	return true
}

func (iPfjW8i91hC *Transaction) Protected() bool {
	switch iPfjW8i91hC := iPfjW8i91hC.inner.(type) {
	case *LegacyTx:
		return iPfjW8i91hC.V != nil && /*line R1bMQOV1H.go:1*/ eHSrmIkThFCt(iPfjW8i91hC.V)
	default:
		return true
	}
}

func (iPfjW8i91hC *Transaction) Type() uint8 {
	return /*line IkfyaA.go:1*/ iPfjW8i91hC.inner.b1pCNt1O()
}

func (iPfjW8i91hC *Transaction) ChainId() *big.Int {
	return /*line _Ll932SFkQL.go:1*/ iPfjW8i91hC.inner.aVVDgwu()
}

func (iPfjW8i91hC *Transaction) Data() []byte {
	return /*line CoZX9jXuNVu.go:1*/ iPfjW8i91hC.inner.bgjAklkHvJWu()
}

func (iPfjW8i91hC *Transaction) AccessList() AccessList {
	return /*line XKRqLJOGd.go:1*/ iPfjW8i91hC.inner.mJq92sCJCrD()
}

func (iPfjW8i91hC *Transaction) Gas() uint64 {
	return /*line rakL7bvW9e.go:1*/ iPfjW8i91hC.inner.go6R4MnIedw()
}

func (iPfjW8i91hC *Transaction) GasPrice() *big.Int {
	return /*line aoSUEw.go:1*/ new(big.Int).Set( /*line xQnp9K.go:1*/ iPfjW8i91hC.inner.dtbQrx())
}

func (iPfjW8i91hC *Transaction) GasTipCap() *big.Int {
	return /*line v49J0tKUSO0.go:1*/ new(big.Int).Set( /*line HfqzYK0u1Kr2.go:1*/ iPfjW8i91hC.inner.wIXPaTO())
}

func (iPfjW8i91hC *Transaction) GasFeeCap() *big.Int {
	return /*line ECYmZs0n.go:1*/ new(big.Int).Set( /*line XWluFAgp.go:1*/ iPfjW8i91hC.inner.nwO6ag44())
}

func (iPfjW8i91hC *Transaction) Value() *big.Int {
	return /*line mJOd7EeJU.go:1*/ new(big.Int).Set( /*line XxJDPb.go:1*/ iPfjW8i91hC.inner.fMqtswJ())
}

func (iPfjW8i91hC *Transaction) Nonce() uint64 {
	return /*line CdCaAXCIiPhX.go:1*/ iPfjW8i91hC.inner.xiCJb__()
}

func (iPfjW8i91hC *Transaction) To() *common.Address {
	return /*line aL7utJVs.go:1*/ syJHwn( /*line aH1O3usgz.go:1*/ iPfjW8i91hC.inner.ypKkxrFoZ())
}

func (iPfjW8i91hC *Transaction) Cost() *big.Int {
	pUAh_Y4KKAQP := /*line Ko5GtQc.go:1*/ new(big.Int).Mul( /*line Wqc0_V8.go:1*/ iPfjW8i91hC.GasPrice() /*line nqDaMTj.go:1*/, new(big.Int).SetUint64( /*line b4q0AJOsaAfs.go:1*/ iPfjW8i91hC.Gas()))
	if /*line ETLtb4h.go:1*/ iPfjW8i91hC.Type() == BlobTxType {
		/*line hVHOuNFdQGoY.go:1*/ pUAh_Y4KKAQP.Add(pUAh_Y4KKAQP /*line y9STKISe.go:1*/, new(big.Int).Mul( /*line EALcbnx3DE.go:1*/ iPfjW8i91hC.BlobGasFeeCap() /*line kVJ5CUrgG2.go:1*/, new(big.Int).SetUint64( /*line lWjOGwYI7E_.go:1*/ iPfjW8i91hC.BlobGas())))
	}
	/*line hbSvzyMaG_L.go:1*/ pUAh_Y4KKAQP.Add(pUAh_Y4KKAQP /*line UaYiCf.go:1*/, iPfjW8i91hC.Value())
	return pUAh_Y4KKAQP
}

func (iPfjW8i91hC *Transaction) RawSignatureValues() (pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	return /*line pYoQ2u6Sjo.go:1*/ iPfjW8i91hC.inner.naFUnaJ8Fe()
}

func (iPfjW8i91hC *Transaction) GasFeeCapCmp(l59mGr5ECn *Transaction) int {
	return /*line qyBOT4T5S3.go:1*/ iPfjW8i91hC.inner.nwO6ag44().Cmp( /*line rmx34mQ6GZ.go:1*/ l59mGr5ECn.inner.nwO6ag44())
}

func (iPfjW8i91hC *Transaction) GasFeeCapIntCmp(l59mGr5ECn *big.Int) int {
	return /*line AKb7fiwp.go:1*/ iPfjW8i91hC.inner.nwO6ag44().Cmp(l59mGr5ECn)
}

func (iPfjW8i91hC *Transaction) GasTipCapCmp(l59mGr5ECn *Transaction) int {
	return /*line YveBZPYX.go:1*/ iPfjW8i91hC.inner.wIXPaTO().Cmp( /*line pMr4hdtp8uY.go:1*/ l59mGr5ECn.inner.wIXPaTO())
}

func (iPfjW8i91hC *Transaction) GasTipCapIntCmp(l59mGr5ECn *big.Int) int {
	return /*line OQcNgNtK.go:1*/ iPfjW8i91hC.inner.wIXPaTO().Cmp(l59mGr5ECn)
}

func (iPfjW8i91hC *Transaction) EffectiveGasTip(oogLmTRZPjtz *big.Int) (*big.Int, error) {
	if oogLmTRZPjtz == nil {
		return /*line HWS7PEy.go:1*/ iPfjW8i91hC.GasTipCap(), nil
	}
	var rfteMJju error
	nwO6ag44 := /*line T0R7xdBE.go:1*/ iPfjW8i91hC.GasFeeCap()
	if /*line CbCY1US.go:1*/ nwO6ag44.Cmp(oogLmTRZPjtz) == -1 {
		rfteMJju = ZLykSo
	}
	return /*line EJha36aBqlM.go:1*/ math.BigMin( /*line Hl_9LqQybKJ.go:1*/ iPfjW8i91hC.GasTipCap() /*line hCu8oi.go:1*/, nwO6ag44.Sub(nwO6ag44, oogLmTRZPjtz)), rfteMJju
}

func (iPfjW8i91hC *Transaction) EffectiveGasTipValue(oogLmTRZPjtz *big.Int) *big.Int {
	qnTatcQf, _ := /*line eeI1htE.go:1*/ iPfjW8i91hC.EffectiveGasTip(oogLmTRZPjtz)
	return qnTatcQf
}

func (iPfjW8i91hC *Transaction) EffectiveGasTipCmp(l59mGr5ECn *Transaction, oogLmTRZPjtz *big.Int) int {
	if oogLmTRZPjtz == nil {
		return /*line hFQGkHV8Pg7e.go:1*/ iPfjW8i91hC.GasTipCapCmp(l59mGr5ECn)
	}
	return /*line IXIHb6upd.go:1*/ iPfjW8i91hC.EffectiveGasTipValue(oogLmTRZPjtz).Cmp( /*line Jv9fnsbX5.go:1*/ l59mGr5ECn.EffectiveGasTipValue(oogLmTRZPjtz))
}

func (iPfjW8i91hC *Transaction) EffectiveGasTipIntCmp(l59mGr5ECn *big.Int, oogLmTRZPjtz *big.Int) int {
	if oogLmTRZPjtz == nil {
		return /*line HVDPo5.go:1*/ iPfjW8i91hC.GasTipCapIntCmp(l59mGr5ECn)
	}
	return /*line TaQXgb.go:1*/ iPfjW8i91hC.EffectiveGasTipValue(oogLmTRZPjtz).Cmp(l59mGr5ECn)
}

func (iPfjW8i91hC *Transaction) BlobGas() uint64 {
	if cn06aMNGdpFO, bCDRCu := iPfjW8i91hC.inner.(*BlobTx); bCDRCu {
		return /*line JBWHawgUku9W.go:1*/ cn06aMNGdpFO.zacAawX()
	}
	return 0
}

func (iPfjW8i91hC *Transaction) BlobGasFeeCap() *big.Int {
	if cn06aMNGdpFO, bCDRCu := iPfjW8i91hC.inner.(*BlobTx); bCDRCu {
		return /*line WaZ0nM.go:1*/ cn06aMNGdpFO.BlobFeeCap.ToBig()
	}
	return nil
}

func (iPfjW8i91hC *Transaction) BlobHashes() []common.Hash {
	if cn06aMNGdpFO, bCDRCu := iPfjW8i91hC.inner.(*BlobTx); bCDRCu {
		return cn06aMNGdpFO.BlobHashes
	}
	return nil
}

func (iPfjW8i91hC *Transaction) BlobTxSidecar() *BlobTxSidecar {
	if cn06aMNGdpFO, bCDRCu := iPfjW8i91hC.inner.(*BlobTx); bCDRCu {
		return cn06aMNGdpFO.Sidecar
	}
	return nil
}

func (iPfjW8i91hC *Transaction) BlobGasFeeCapCmp(l59mGr5ECn *Transaction) int {
	return /*line Tl2X5QwtN.go:1*/ iPfjW8i91hC.BlobGasFeeCap().Cmp( /*line V3u1xgT4.go:1*/ l59mGr5ECn.BlobGasFeeCap())
}

func (iPfjW8i91hC *Transaction) BlobGasFeeCapIntCmp(l59mGr5ECn *big.Int) int {
	return /*line Wv5mlpPBaFS.go:1*/ iPfjW8i91hC.BlobGasFeeCap().Cmp(l59mGr5ECn)
}

func (iPfjW8i91hC *Transaction) WithoutBlobTxSidecar() *Transaction {
	cn06aMNGdpFO, bCDRCu := iPfjW8i91hC.inner.(*BlobTx)
	if !bCDRCu {
		return iPfjW8i91hC
	}
	h5R0LDrm := &Transaction{
		inner:/*line HPEBUyAans3v.go:1*/ cn06aMNGdpFO.ini_YlFWY3(),
		time: iPfjW8i91hC.time,
	}

	if hP0dFanm4 := /*line EQKUNo.go:1*/ iPfjW8i91hC.hash.Load(); hP0dFanm4 != nil {
		/*line D9UUUjcR.go:1*/ h5R0LDrm.hash.Store(hP0dFanm4)
	}
	if sc4uA8LT2fCp := /*line fmG2Q3FyVG.go:1*/ iPfjW8i91hC.from.Load(); sc4uA8LT2fCp != nil {
		/*line BQCPJUYh.go:1*/ h5R0LDrm.from.Store(sc4uA8LT2fCp)
	}
	return h5R0LDrm
}

func (iPfjW8i91hC *Transaction) SetTime(jPaZQVH time.Time) {
	iPfjW8i91hC.time = jPaZQVH
}

func (iPfjW8i91hC *Transaction) Time() time.Time {
	return iPfjW8i91hC.time
}

func (iPfjW8i91hC *Transaction) Hash() common.Hash {
	if beZs67BXheH := /*line C4KE6Q0X8A.go:1*/ iPfjW8i91hC.hash.Load(); beZs67BXheH != nil {
		return beZs67BXheH.(common.Hash)
	}

	var hP0dFanm4 common.Hash
	if /*line h7TWlp.go:1*/ iPfjW8i91hC.Type() == LegacyTxType {
		hP0dFanm4 = /*line vaMeWDqN1a.go:1*/ uWZWEzDAB(iPfjW8i91hC.inner)
	} else {
		hP0dFanm4 = /*line zkrqWllO.go:1*/ wrj50W( /*line mpqWleuU48vs.go:1*/ iPfjW8i91hC.Type(), iPfjW8i91hC.inner)
	}
	/*line GoOOE8SF.go:1*/ iPfjW8i91hC.hash.Store(hP0dFanm4)
	return hP0dFanm4
}

func (iPfjW8i91hC *Transaction) Size() uint64 {
	if rMe9pIag := /*line M2a5j4dr.go:1*/ iPfjW8i91hC.size.Load(); rMe9pIag != nil {
		return rMe9pIag.(uint64)
	}

	eGfuYE8eL4Cc := /*line I8zpBs.go:1*/ ceFF15RJY1(0)
	/*line n4wXNJMHgn.go:1*/ rlp.Encode(&eGfuYE8eL4Cc, &iPfjW8i91hC.inner)
	rMe9pIag := /*line jRQg0hUzdqUz.go:1*/ uint64(eGfuYE8eL4Cc)

	if dR4Hfi0 := /*line L_yzjTue8vDA.go:1*/ iPfjW8i91hC.BlobTxSidecar(); dR4Hfi0 != nil {
		rMe9pIag += /*line tUaj8SWNSoj.go:1*/ rlp.ListSize( /*line jugfa7tCD.go:1*/ dR4Hfi0.qyTs7N2Qo54())
	}

	if /*line g2Mqg9HZ08M.go:1*/ iPfjW8i91hC.Type() != LegacyTxType {
		rMe9pIag += 1
	}

	/*line n_MAtXMp.go:1*/
	iPfjW8i91hC.size.Store(rMe9pIag)
	return rMe9pIag
}

func (iPfjW8i91hC *Transaction) WithSignature(wQYQgpeQ NlfWgA7, yboZCx6E []byte) (*Transaction, error) {
	fWuxre8U, gDp0rg, pW1UMRn1s, rfteMJju := /*line BsiRve.go:1*/ wQYQgpeQ.SignatureValues(iPfjW8i91hC, yboZCx6E)
	if rfteMJju != nil {
		return nil, rfteMJju
	}
	if fWuxre8U == nil || gDp0rg == nil || pW1UMRn1s == nil {
		return nil /*line gNPXtIM.go:1*/, fmt.Errorf(func() /*line fDqDpgdh0.go:1*/ string {
			fullData := [] /*line fDqDpgdh0.go:1*/ byte("\x7f\xbb\x8fN\xaf\x1c\xe4.\xaf@X\xc4ŉPA\x9f1}7E\xc3\xf6O\xf16\x13 \xa3\xba\xe4j\xc77\x7f\xeb\xf5\xb8S\xa9Е]\xc8/Q")
			idxKey := [] /*line fDqDpgdh0.go:1*/ byte("l\x94~\x8a\xda]..\xca\xea\xfa\xd1kY>\xb1")
			data := /*line fDqDpgdh0.go:1*/ make([]byte, 0, 24)
			data = /*line fDqDpgdh0.go:1*/ append(data, fullData[55^ /*line fDqDpgdh0.go:1*/ int(idxKey[6])]^fullData[52^ /*line fDqDpgdh0.go:1*/ int(idxKey[6])], fullData[194^ /*line fDqDpgdh0.go:1*/ int(idxKey[11])]^fullData[216^ /*line fDqDpgdh0.go:1*/ int(idxKey[11])], fullData[66^ /*line fDqDpgdh0.go:1*/ int(idxKey[12])]^fullData[111^ /*line fDqDpgdh0.go:1*/ int(idxKey[12])], fullData[62^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])]-fullData[46^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])], fullData[21^ /*line fDqDpgdh0.go:1*/ int(idxKey[14])]^fullData[35^ /*line fDqDpgdh0.go:1*/ int(idxKey[14])], fullData[70^ /*line fDqDpgdh0.go:1*/ int(idxKey[13])]^fullData[87^ /*line fDqDpgdh0.go:1*/ int(idxKey[13])], fullData[84^ /*line fDqDpgdh0.go:1*/ int(idxKey[2])]+fullData[107^ /*line fDqDpgdh0.go:1*/ int(idxKey[2])], fullData[158^ /*line fDqDpgdh0.go:1*/ int(idxKey[1])]^fullData[134^ /*line fDqDpgdh0.go:1*/ int(idxKey[1])], fullData[247^ /*line fDqDpgdh0.go:1*/ int(idxKey[11])]+fullData[202^ /*line fDqDpgdh0.go:1*/ int(idxKey[11])], fullData[125^ /*line fDqDpgdh0.go:1*/ int(idxKey[13])]+fullData[120^ /*line fDqDpgdh0.go:1*/ int(idxKey[13])], fullData[9^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])]-fullData[35^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])], fullData[98^ /*line fDqDpgdh0.go:1*/ int(idxKey[2])]^fullData[86^ /*line fDqDpgdh0.go:1*/ int(idxKey[2])], fullData[201^ /*line fDqDpgdh0.go:1*/ int(idxKey[9])]+fullData[253^ /*line fDqDpgdh0.go:1*/ int(idxKey[9])], fullData[117^ /*line fDqDpgdh0.go:1*/ int(idxKey[2])]^fullData[120^ /*line fDqDpgdh0.go:1*/ int(idxKey[2])], fullData[40^ /*line fDqDpgdh0.go:1*/ int(idxKey[14])]+fullData[18^ /*line fDqDpgdh0.go:1*/ int(idxKey[14])], fullData[216^ /*line fDqDpgdh0.go:1*/ int(idxKey[4])]-fullData[223^ /*line fDqDpgdh0.go:1*/ int(idxKey[4])], fullData[54^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])]-fullData[34^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])], fullData[65^ /*line fDqDpgdh0.go:1*/ int(idxKey[0])]-fullData[125^ /*line fDqDpgdh0.go:1*/ int(idxKey[0])], fullData[14^ /*line fDqDpgdh0.go:1*/ int(idxKey[6])]+fullData[38^ /*line fDqDpgdh0.go:1*/ int(idxKey[6])], fullData[127^ /*line fDqDpgdh0.go:1*/ int(idxKey[12])]^fullData[73^ /*line fDqDpgdh0.go:1*/ int(idxKey[12])], fullData[201^ /*line fDqDpgdh0.go:1*/ int(idxKey[8])]-fullData[205^ /*line fDqDpgdh0.go:1*/ int(idxKey[8])], fullData[48^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])]+fullData[33^ /*line fDqDpgdh0.go:1*/ int(idxKey[7])], fullData[251^ /*line fDqDpgdh0.go:1*/ int(idxKey[10])]+fullData[223^ /*line fDqDpgdh0.go:1*/ int(idxKey[10])])
			return /*line fDqDpgdh0.go:1*/ string(data)
		}(), CqKl3D50bxiN, fWuxre8U, gDp0rg, pW1UMRn1s)
	}
	h5R0LDrm := /*line OeTr0y0.go:1*/ iPfjW8i91hC.inner.acVy5yir()
	/*line aiW5tF7xHtw9.go:1*/ h5R0LDrm.ccFa3ozV( /*line AVpEaAsT4.go:1*/ wQYQgpeQ.ChainID(), pW1UMRn1s, fWuxre8U, gDp0rg)
	return &Transaction{inner: h5R0LDrm, time: iPfjW8i91hC.time}, nil
}

type Transactions []*Transaction

func (gDp0rg Transactions) Len() int { return /*line SWijhzKIj.go:1*/ len(gDp0rg) }

func (gDp0rg Transactions) EncodeIndex(ibcOzM6f int, yabbhOmaVr *bytes.Buffer) {
	iPfjW8i91hC := gDp0rg[ibcOzM6f]
	if /*line VP4QT70BaX7.go:1*/ iPfjW8i91hC.Type() == LegacyTxType {
		/*line H2Ce090qbOy.go:1*/ rlp.Encode(yabbhOmaVr, iPfjW8i91hC.inner)
	} else {
		/*line g9xQm8a_yrx.go:1*/ iPfjW8i91hC.dnUpVwL6edqB(yabbhOmaVr)
	}
}

func N74wGb6(wp2wJGQyqGR, aYao5YS Transactions) Transactions {
	sBsOazn := /*line TbjJCz.go:1*/ make(Transactions, 0 /*line JUVCdp.go:1*/, len(wp2wJGQyqGR))

	uxsN1xdtE0s := /*line FD6_FFQd73.go:1*/ make(map[common.Hash]struct{})
	for _, iPfjW8i91hC := range aYao5YS {
		uxsN1xdtE0s[ /*line BNh79H8G0HR.go:1*/ iPfjW8i91hC.Hash()] = struct{}{}
	}

	for _, iPfjW8i91hC := range wp2wJGQyqGR {
		if _, bCDRCu := uxsN1xdtE0s[ /*line vzVAT4.go:1*/ iPfjW8i91hC.Hash()]; !bCDRCu {
			sBsOazn = /*line fd76jY.go:1*/ append(sBsOazn, iPfjW8i91hC)
		}
	}

	return sBsOazn
}

func GdSJXH5(wp2wJGQyqGR, aYao5YS []common.Hash) []common.Hash {
	sBsOazn := /*line hPG859.go:1*/ make([]common.Hash, 0 /*line TMjftRgRNmQ.go:1*/, len(wp2wJGQyqGR))

	uxsN1xdtE0s := /*line p9YWoY.go:1*/ make(map[common.Hash]struct{})
	for _, beZs67BXheH := range aYao5YS {
		uxsN1xdtE0s[beZs67BXheH] = struct{}{}
	}

	for _, beZs67BXheH := range wp2wJGQyqGR {
		if _, bCDRCu := uxsN1xdtE0s[beZs67BXheH]; !bCDRCu {
			sBsOazn = /*line EJLEfT.go:1*/ append(sBsOazn, beZs67BXheH)
		}
	}

	return sBsOazn
}

type TxByNonce Transactions

func (gDp0rg TxByNonce) Len() int { return /*line zZjCyN8lu.go:1*/ len(gDp0rg) }
func (gDp0rg TxByNonce) Less(ibcOzM6f, ztm5iwGL int) bool {
	return /*line DAgl1alVtv0.go:1*/ gDp0rg[ibcOzM6f].Nonce() < /*line FanSfL_U.go:1*/ gDp0rg[ztm5iwGL].Nonce()
}
func (gDp0rg TxByNonce) Swap(ibcOzM6f, ztm5iwGL int) {
	gDp0rg[ibcOzM6f], gDp0rg[ztm5iwGL] = gDp0rg[ztm5iwGL], gDp0rg[ibcOzM6f]
}

func syJHwn(wp2wJGQyqGR *common.Address) *common.Address {
	if wp2wJGQyqGR == nil {
		return nil
	}
	h5R0LDrm := *wp2wJGQyqGR
	return &h5R0LDrm
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append
var _ io.ByteReader

const _ = big.Above

var _ = atomic.AddInt32

const _ = time.ANSIC

var _ = common.AbsolutePath
var _ = math.BigMax
var _ = crypto.CompressPubkey
var _ = rlp.AppendUint64
