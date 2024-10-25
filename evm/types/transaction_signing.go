//line :1
package types

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

var HCy4ZyvVELw = /*line bimJ5lnEn.go:1*/ errors.New(func() /*line K5UVf8UhlPTW.go:1*/ string {
	fullData := [] /*line K5UVf8UhlPTW.go:1*/ byte("໖\x84Ѯ\u05cf\xc0ŧؼ\xeb\U000511ef\xb5\xc09\x15\r\xa5\x83\xceΤ\x17Me7W2\xb6\x87\xbch[\xd2\xe9\xb4#\xf4\x90̈́\xfe\xe6]\xc4!\xa4I")
	idxKey := [] /*line K5UVf8UhlPTW.go:1*/ byte("n\xd2\x1cy-Dʣ\x98\xe5\xc2&FΉ\x9a")
	data := /*line K5UVf8UhlPTW.go:1*/ make([]byte, 0, 28)
	data = /*line K5UVf8UhlPTW.go:1*/ append(data, fullData[197^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])]+fullData[193^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])], fullData[16^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[2])]^fullData[59^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[2])], fullData[62^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[2])]+fullData[20^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[2])], fullData[104^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[12])]-fullData[108^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[12])], fullData[143^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[8])]-fullData[140^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[8])], fullData[50^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[4])]-fullData[52^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[4])], fullData[248^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[9])]+fullData[249^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[9])], fullData[217^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])]^fullData[202^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])], fullData[46^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[4])]-fullData[30^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[4])], fullData[183^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[15])]-fullData[132^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[15])], fullData[35^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[4])]^fullData[1^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[4])], fullData[220^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])]-fullData[254^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])], fullData[223^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])]-fullData[192^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])], fullData[223^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[13])]^fullData[201^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[13])], fullData[227^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])]+fullData[216^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])], fullData[149^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[8])]-fullData[187^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[8])], fullData[228^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[10])]+fullData[203^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[10])], fullData[172^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[14])]+fullData[166^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[14])], fullData[7^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[11])]^fullData[23^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[11])], fullData[115^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[12])]-fullData[64^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[12])], fullData[20^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[11])]-fullData[61^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[11])], fullData[245^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[9])]^fullData[206^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[9])], fullData[207^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[13])]+fullData[203^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[13])], fullData[198^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[10])]+fullData[192^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[10])], fullData[226^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[10])]-fullData[234^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[10])], fullData[250^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])]^fullData[210^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[6])], fullData[98^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[12])]^fullData[92^ /*line K5UVf8UhlPTW.go:1*/ int(idxKey[12])])
	return /*line K5UVf8UhlPTW.go:1*/ string(data)
}())

type c5Q3pM6fQhA struct {
	o2bxAiUROf1O NlfWgA7
	nqGN0G8N7id_ common.Address
}

func WS4FDJv29l(yRKUDdMKWa *params.ChainConfig, zfFFIov8lRi *big.Int, eR8oHbfpQct uint64) NlfWgA7 {
	var wQYQgpeQ NlfWgA7
	switch {
	case /*line AWTjQd.go:1*/ yRKUDdMKWa.IsCancun(zfFFIov8lRi, eR8oHbfpQct):
		wQYQgpeQ = /*line mq39w1XIl.go:1*/ FIJ9kSHHl(yRKUDdMKWa.ChainID)
	case /*line D7l7bAYaa.go:1*/ yRKUDdMKWa.IsLondon(zfFFIov8lRi):
		wQYQgpeQ = /*line KMTTPk.go:1*/ X8XHgjw2AF_(yRKUDdMKWa.ChainID)
	case /*line qHK6Tf.go:1*/ yRKUDdMKWa.IsBerlin(zfFFIov8lRi):
		wQYQgpeQ = /*line Fdi8aiFQg.go:1*/ JE1FTA(yRKUDdMKWa.ChainID)
	case /*line Ka9AXR.go:1*/ yRKUDdMKWa.IsEIP155(zfFFIov8lRi):
		wQYQgpeQ = /*line pXAttL.go:1*/ GVrt8eeUJ(yRKUDdMKWa.ChainID)
	case /*line yoeEB8N.go:1*/ yRKUDdMKWa.IsHomestead(zfFFIov8lRi):
		wQYQgpeQ = H6gCgPO3{}
	default:
		wQYQgpeQ = We6m1j6t{}
	}
	return wQYQgpeQ
}

func RyhNIzr(yRKUDdMKWa *params.ChainConfig) NlfWgA7 {
	if yRKUDdMKWa.ChainID != nil {
		if yRKUDdMKWa.CancunTime != nil {
			return /*line aIBOVG3rtY.go:1*/ FIJ9kSHHl(yRKUDdMKWa.ChainID)
		}
		if yRKUDdMKWa.LondonBlock != nil {
			return /*line aBHHcuO.go:1*/ X8XHgjw2AF_(yRKUDdMKWa.ChainID)
		}
		if yRKUDdMKWa.BerlinBlock != nil {
			return /*line Xi5dTVu3Y_yz.go:1*/ JE1FTA(yRKUDdMKWa.ChainID)
		}
		if yRKUDdMKWa.EIP155Block != nil {
			return /*line AvoLug2ab.go:1*/ GVrt8eeUJ(yRKUDdMKWa.ChainID)
		}
	}
	return H6gCgPO3{}
}

func LLmYbHJ8F(aVVDgwu *big.Int) NlfWgA7 {
	if aVVDgwu == nil {
		return H6gCgPO3{}
	}
	return /*line isvqxUOzgA.go:1*/ FIJ9kSHHl(aVVDgwu)
}

func QUMkEZp81(iPfjW8i91hC *Transaction, gDp0rg NlfWgA7, auMuc5Nk_SB *ecdsa.PrivateKey) (*Transaction, error) {
	hP0dFanm4 := /*line UZRqxFZO.go:1*/ gDp0rg.Hash(iPfjW8i91hC)
	yboZCx6E, rfteMJju := /*line aXe4GK6.go:1*/ crypto.Sign(hP0dFanm4[:], auMuc5Nk_SB)
	if rfteMJju != nil {
		return nil, rfteMJju
	}
	return /*line TiIwrSmn.go:1*/ iPfjW8i91hC.WithSignature(gDp0rg, yboZCx6E)
}

func KCSux8Bg(auMuc5Nk_SB *ecdsa.PrivateKey, gDp0rg NlfWgA7, hqz9d9MXjE TxData) (*Transaction, error) {
	iPfjW8i91hC := /*line JV_qBAyI4h.go:1*/ CV6bt80q7Gn(hqz9d9MXjE)
	hP0dFanm4 := /*line CeieMH.go:1*/ gDp0rg.Hash(iPfjW8i91hC)
	yboZCx6E, rfteMJju := /*line hPhx393Gcewb.go:1*/ crypto.Sign(hP0dFanm4[:], auMuc5Nk_SB)
	if rfteMJju != nil {
		return nil, rfteMJju
	}
	return /*line dGIJS81u.go:1*/ iPfjW8i91hC.WithSignature(gDp0rg, yboZCx6E)
}

func KVcwtBNMwa(auMuc5Nk_SB *ecdsa.PrivateKey, gDp0rg NlfWgA7, hqz9d9MXjE TxData) *Transaction {
	iPfjW8i91hC, rfteMJju := /*line wtVP8n.go:1*/ KCSux8Bg(auMuc5Nk_SB, gDp0rg, hqz9d9MXjE)
	if rfteMJju != nil {
		/*line FcN49iDBR.go:1*/ panic(rfteMJju)
	}
	return iPfjW8i91hC
}

func Zsrn5l(wQYQgpeQ NlfWgA7, iPfjW8i91hC *Transaction) (common.Address, error) {
	if dR4Hfi0 := /*line FBoPayGaH.go:1*/ iPfjW8i91hC.from.Load(); dR4Hfi0 != nil {
		c5Q3pM6fQhA := dR4Hfi0.(c5Q3pM6fQhA)

		if /*line zDyGQ6fdmb.go:1*/ c5Q3pM6fQhA.o2bxAiUROf1O.Equal(wQYQgpeQ) {
			return c5Q3pM6fQhA.nqGN0G8N7id_, nil
		}
	}

	yOdETLh4, rfteMJju := /*line PeLzaNqRfc.go:1*/ wQYQgpeQ.Sender(iPfjW8i91hC)
	if rfteMJju != nil {
		return common.Address{}, rfteMJju
	}
	/*line HZQ8ji8cb.go:1*/ iPfjW8i91hC.from.Store(c5Q3pM6fQhA{o2bxAiUROf1O: wQYQgpeQ, nqGN0G8N7id_: yOdETLh4})
	return yOdETLh4, nil
}

type NlfWgA7 interface {
	Sender(iPfjW8i91hC *Transaction) (common.Address, error)

	SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (fWuxre8U, gDp0rg, pW1UMRn1s *big.Int, rfteMJju error)
	ChainID() *big.Int

	Hash(iPfjW8i91hC *Transaction) common.Hash

	Equal(NlfWgA7) bool
}

type cancunSigner struct{ londonSigner }

func FIJ9kSHHl(tkkFB3 *big.Int) NlfWgA7 {
	return cancunSigner{londonSigner{eip2930Signer{ /*line SJZHXvRL_OdQ.go:1*/ GVrt8eeUJ(tkkFB3)}}}
}

func (gDp0rg cancunSigner) Sender(iPfjW8i91hC *Transaction) (common.Address, error) {
	if /*line ySmaKPiuo.go:1*/ iPfjW8i91hC.Type() != BlobTxType {
		return /*line c5_U_jGHSA5c.go:1*/ gDp0rg.londonSigner.Sender(iPfjW8i91hC)
	}
	ICKLPq3pU, LhajwZArxs, ChPRkRQT := /*line Cuj7fL8pLgf.go:1*/ iPfjW8i91hC.RawSignatureValues()

	ICKLPq3pU = /*line acUkcBaOfNT.go:1*/ new(big.Int).Add(ICKLPq3pU /*line ISQHspdaSc.go:1*/, big.NewInt(27))
	if /*line _r8sc80NLSO.go:1*/ iPfjW8i91hC.ChainId().Cmp(gDp0rg.chainId) != 0 {
		return common.Address{} /*line qNs07qaA3.go:1*/, fmt.Errorf(func() /*line n_9tpCcJ.go:1*/ string {
			fullData := [] /*line n_9tpCcJ.go:1*/ byte("\xa2BW8\x1a&\xa3\xba\x86\x9a\xd8G\xd5\x01 F\xbc\xffL\xd9\xc3\xca\xdf\xe0똇\xb5\x1f\x1e\x133\xbc\xc0Ԭ\r\x16")
			idxKey := [] /*line n_9tpCcJ.go:1*/ byte("u\xd2")
			data := /*line n_9tpCcJ.go:1*/ make([]byte, 0, 20)
			data = /*line n_9tpCcJ.go:1*/ append(data, fullData[212^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]^fullData[218^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[119^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]+fullData[123^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[108^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]+fullData[117^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[100^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]^fullData[99^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[240^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]^fullData[242^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[214^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]+fullData[217^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[110^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]^fullData[97^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[102^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]^fullData[101^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[98^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]-fullData[84^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[205^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]^fullData[247^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[202^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]-fullData[200^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[124^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]^fullData[114^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[192^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]-fullData[222^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[105^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]+fullData[116^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[209^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])]-fullData[199^ /*line n_9tpCcJ.go:1*/ int(idxKey[1])], fullData[127^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]^fullData[86^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[81^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]+fullData[107^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[112^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]-fullData[120^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])], fullData[122^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])]+fullData[104^ /*line n_9tpCcJ.go:1*/ int(idxKey[0])])
			return /*line n_9tpCcJ.go:1*/ string(data)
		}(), HCy4ZyvVELw /*line R6iFCBN.go:1*/, iPfjW8i91hC.ChainId(), gDp0rg.chainId)
	}
	return /*line ExzHbbS1h.go:1*/ dAgH1L( /*line qSzugy.go:1*/ gDp0rg.Hash(iPfjW8i91hC), LhajwZArxs, ChPRkRQT, ICKLPq3pU, true)
}

func (gDp0rg cancunSigner) Equal(gjNSUeDUec NlfWgA7) bool {
	uO81zUpF, bCDRCu := gjNSUeDUec.(cancunSigner)
	return bCDRCu && /*line Cc2vp2wcX.go:1*/ uO81zUpF.chainId.Cmp(gDp0rg.chainId) == 0
}

func (gDp0rg cancunSigner) SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (LhajwZArxs, ChPRkRQT, ICKLPq3pU *big.Int, rfteMJju error) {
	hqz9d9MXjE, bCDRCu := iPfjW8i91hC.inner.(*BlobTx)
	if !bCDRCu {
		return /*line wCkrqa.go:1*/ gDp0rg.londonSigner.SignatureValues(iPfjW8i91hC, yboZCx6E)
	}

	if /*line N7Qun3Qseb1.go:1*/ hqz9d9MXjE.ChainID.Sign() != 0 && /*line S_vVNbyHZj.go:1*/ hqz9d9MXjE.ChainID.ToBig().Cmp(gDp0rg.chainId) != 0 {
		return nil, nil, nil /*line i9lilasLD.go:1*/, fmt.Errorf(func() /*line apDYUEev.go:1*/ string {
			data := [] /*line apDYUEev.go:1*/ byte("z\xa3e zavﻯ\xb6\x17dT\xbct\xb0\xb5\xc6")
			positions := [...]byte{9, 7, 11, 13, 4, 13, 16, 13, 10, 1, 4, 16, 4, 17, 17, 16, 18, 14, 0, 2, 12, 8, 9, 7}
			for i := 0; i < 24; i += 2 {
				localKey := /*line apDYUEev.go:1*/ byte(i) + /*line apDYUEev.go:1*/ byte(positions[i]^positions[i+1]) + 44
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line apDYUEev.go:1*/ string(data)
		}(), HCy4ZyvVELw, hqz9d9MXjE.ChainID, gDp0rg.chainId)
	}
	LhajwZArxs, ChPRkRQT, _ = /*line oauQCgnJW.go:1*/ wHPfzXWnSUX(yboZCx6E)
	ICKLPq3pU = /*line kHwioE8_fMg.go:1*/ big.NewInt( /*line ZhTUyrA2Pp.go:1*/ int64(yboZCx6E[64]))
	return LhajwZArxs, ChPRkRQT, ICKLPq3pU, nil
}

func (gDp0rg cancunSigner) Hash(iPfjW8i91hC *Transaction) common.Hash {
	if /*line U1owGl.go:1*/ iPfjW8i91hC.Type() != BlobTxType {
		return /*line Y6BBsLp6SPf.go:1*/ gDp0rg.londonSigner.Hash(iPfjW8i91hC)
	}
	return /*line tWK17vGtal.go:1*/ wrj50W(
		/*line GazV5yQdJACa.go:1*/ iPfjW8i91hC.Type(),
		[]interface{}{
			gDp0rg.chainId,
			/*line DDzeocavfw5.go:1*/ iPfjW8i91hC.Nonce(),
			/*line Mk9iJDsV.go:1*/ iPfjW8i91hC.GasTipCap(),
			/*line AyOAep0t.go:1*/ iPfjW8i91hC.GasFeeCap(),
			/*line pJAVwo.go:1*/ iPfjW8i91hC.Gas(),
			/*line s9PMt3MZ.go:1*/ iPfjW8i91hC.To(),
			/*line G9SCxmj7M2Vn.go:1*/ iPfjW8i91hC.Value(),
			/*line qnYreWzaxq.go:1*/ iPfjW8i91hC.Data(),
			/*line yT8vgh65GQ.go:1*/ iPfjW8i91hC.AccessList(),
			/*line xHbGrbfF.go:1*/ iPfjW8i91hC.BlobGasFeeCap(),
			/*line bL8NsXwSX.go:1*/ iPfjW8i91hC.BlobHashes(),
		})
}

type londonSigner struct{ eip2930Signer }

func X8XHgjw2AF_(tkkFB3 *big.Int) NlfWgA7 {
	return londonSigner{eip2930Signer{ /*line ywYPGusC3_z.go:1*/ GVrt8eeUJ(tkkFB3)}}
}

func (gDp0rg londonSigner) Sender(iPfjW8i91hC *Transaction) (common.Address, error) {
	if /*line Bz5w1b07.go:1*/ iPfjW8i91hC.Type() != DynamicFeeTxType {
		return /*line oDWV8Y7wsaDZ.go:1*/ gDp0rg.eip2930Signer.Sender(iPfjW8i91hC)
	}
	ICKLPq3pU, LhajwZArxs, ChPRkRQT := /*line FKl9WK10.go:1*/ iPfjW8i91hC.RawSignatureValues()

	ICKLPq3pU = /*line Mb5TGtlp.go:1*/ new(big.Int).Add(ICKLPq3pU /*line Wv7HETQ9.go:1*/, big.NewInt(27))
	if /*line evYGVuws.go:1*/ iPfjW8i91hC.ChainId().Cmp(gDp0rg.chainId) != 0 {
		return common.Address{} /*line BdXNooVusw.go:1*/, fmt.Errorf(func() /*line T9gzFdEkU.go:1*/ string {
			data := /*line T9gzFdEkU.go:1*/ make([]byte, 0, 20)
			i := 5
			decryptKey := 215
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					data = /*line T9gzFdEkU.go:1*/ append(data, "\xf0\xa0\xb7"...,
					)
					i = 7
				case 6:
					i = 3
					data = /*line T9gzFdEkU.go:1*/ append(data, "\xf3\xf7\xb5"...,
					)
				case 0:
					i = 2
					for y := range data {
						data[y] = data[y] ^ /*line T9gzFdEkU.go:1*/ byte(decryptKey^y)
					}
				case 1:
					i = 6
					data = /*line T9gzFdEkU.go:1*/ append(data, "\xb7\xbf\xab\xb9"...,
					)
				case 7:
					data = /*line T9gzFdEkU.go:1*/ append(data, "\xbb\xa0"...,
					)
					i = 4
				case 4:
					i = 0
					data = /*line T9gzFdEkU.go:1*/ append(data, "\xeb\xef\xad"...,
					)
				case 5:
					i = 1
					data = /*line T9gzFdEkU.go:1*/ append(data, "\xfe\xad\xe3\xf8"...,
					)
				}
			}
			return /*line T9gzFdEkU.go:1*/ string(data)
		}(), HCy4ZyvVELw /*line zQaSl1Fa.go:1*/, iPfjW8i91hC.ChainId(), gDp0rg.chainId)
	}
	return /*line eTS61r.go:1*/ dAgH1L( /*line TFlam07I4.go:1*/ gDp0rg.Hash(iPfjW8i91hC), LhajwZArxs, ChPRkRQT, ICKLPq3pU, true)
}

func (gDp0rg londonSigner) Equal(gjNSUeDUec NlfWgA7) bool {
	uO81zUpF, bCDRCu := gjNSUeDUec.(londonSigner)
	return bCDRCu && /*line Ehks0paOhWrr.go:1*/ uO81zUpF.chainId.Cmp(gDp0rg.chainId) == 0
}

func (gDp0rg londonSigner) SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (LhajwZArxs, ChPRkRQT, ICKLPq3pU *big.Int, rfteMJju error) {
	hqz9d9MXjE, bCDRCu := iPfjW8i91hC.inner.(*DynamicFeeTx)
	if !bCDRCu {
		return /*line GX6MKv.go:1*/ gDp0rg.eip2930Signer.SignatureValues(iPfjW8i91hC, yboZCx6E)
	}

	if /*line fcbgWL.go:1*/ hqz9d9MXjE.ChainID.Sign() != 0 && /*line HDseU1graCZx.go:1*/ hqz9d9MXjE.ChainID.Cmp(gDp0rg.chainId) != 0 {
		return nil, nil, nil /*line Qcer6dDD.go:1*/, fmt.Errorf(func() /*line gFB4tzLJ.go:1*/ string {
			seed := /*line gFB4tzLJ.go:1*/ byte(214)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line gFB4tzLJ.go:1*/ append(data, x+seed); seed += x; return fnc }
			/*line gFB4tzLJ.go:1*/ fnc(79)(82)(195)(230)(72)(249)(21)(239)(187)(5)(63)(188)(87)(234)(13)(6)(172)(5)(63)
			return /*line gFB4tzLJ.go:1*/ string(data)
		}(), HCy4ZyvVELw, hqz9d9MXjE.ChainID, gDp0rg.chainId)
	}
	LhajwZArxs, ChPRkRQT, _ = /*line GgERGGx3gNIi.go:1*/ wHPfzXWnSUX(yboZCx6E)
	ICKLPq3pU = /*line C3eue1.go:1*/ big.NewInt( /*line YW3RRUFBaRg.go:1*/ int64(yboZCx6E[64]))
	return LhajwZArxs, ChPRkRQT, ICKLPq3pU, nil
}

func (gDp0rg londonSigner) Hash(iPfjW8i91hC *Transaction) common.Hash {
	if /*line AC6WYp.go:1*/ iPfjW8i91hC.Type() != DynamicFeeTxType {
		return /*line kkVaGX9q_ZDk.go:1*/ gDp0rg.eip2930Signer.Hash(iPfjW8i91hC)
	}
	return /*line HSFJj_Y.go:1*/ wrj50W(
		/*line pdsGcBi.go:1*/ iPfjW8i91hC.Type(),
		[]interface{}{
			gDp0rg.chainId,
			/*line aTMqsY_sjyl.go:1*/ iPfjW8i91hC.Nonce(),
			/*line BeDr1PE.go:1*/ iPfjW8i91hC.GasTipCap(),
			/*line ty4cpMFHW.go:1*/ iPfjW8i91hC.GasFeeCap(),
			/*line tCfDcaQi.go:1*/ iPfjW8i91hC.Gas(),
			/*line qnWiuC0Q.go:1*/ iPfjW8i91hC.To(),
			/*line ou7kLMRa.go:1*/ iPfjW8i91hC.Value(),
			/*line qawbB01.go:1*/ iPfjW8i91hC.Data(),
			/*line BxDhjx7.go:1*/ iPfjW8i91hC.AccessList(),
		})
}

type eip2930Signer struct{ EIP155Signer }

func JE1FTA(tkkFB3 *big.Int) NlfWgA7 {
	return eip2930Signer{ /*line I4MGMI_zQ.go:1*/ GVrt8eeUJ(tkkFB3)}
}

func (gDp0rg eip2930Signer) ChainID() *big.Int {
	return gDp0rg.chainId
}

func (gDp0rg eip2930Signer) Equal(gjNSUeDUec NlfWgA7) bool {
	uO81zUpF, bCDRCu := gjNSUeDUec.(eip2930Signer)
	return bCDRCu && /*line nahb2rYna.go:1*/ uO81zUpF.chainId.Cmp(gDp0rg.chainId) == 0
}

func (gDp0rg eip2930Signer) Sender(iPfjW8i91hC *Transaction) (common.Address, error) {
	ICKLPq3pU, LhajwZArxs, ChPRkRQT := /*line dRl49J.go:1*/ iPfjW8i91hC.RawSignatureValues()
	switch /*line KNkZAEpIf8KL.go:1*/ iPfjW8i91hC.Type() {
	case LegacyTxType:
		return /*line AfRiMa.go:1*/ gDp0rg.EIP155Signer.Sender(iPfjW8i91hC)
	case AccessListTxType:

		ICKLPq3pU = /*line pJmvsrqJI.go:1*/ new(big.Int).Add(ICKLPq3pU /*line vuGlFRt.go:1*/, big.NewInt(27))
	default:
		return common.Address{}, B0ZTT4Gv0t
	}
	if /*line astvrTnmjqU.go:1*/ iPfjW8i91hC.ChainId().Cmp(gDp0rg.chainId) != 0 {
		return common.Address{} /*line CN0WSM5.go:1*/, fmt.Errorf(func() /*line ARMrBo6HPYh.go:1*/ string {
			fullData := [] /*line ARMrBo6HPYh.go:1*/ byte("\b\xad\xde\x1b\x04\x9b\x04\xc46\xd9L\x95*$j\x8b\\\xd9\x02\xa1\xbc.\xbf\x99\xfc>\x1e\x03\x926\xb2\xc2\xc9\\c\\A\xe0")
			idxKey := [] /*line ARMrBo6HPYh.go:1*/ byte("dC")
			data := /*line ARMrBo6HPYh.go:1*/ make([]byte, 0, 20)
			data = /*line ARMrBo6HPYh.go:1*/ append(data, fullData[115^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]^fullData[112^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[108^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]-fullData[114^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[71^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]+fullData[102^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[72^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]+fullData[76^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[88^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]-fullData[70^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[104^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]-fullData[68^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[92^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]-fullData[73^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[68^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]+fullData[80^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[122^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]-fullData[120^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[70^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]-fullData[125^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[113^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]+fullData[121^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[69^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]^fullData[78^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[103^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]+fullData[69^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[103^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]-fullData[102^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[71^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]^fullData[77^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[109^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]^fullData[101^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])], fullData[81^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]+fullData[89^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[91^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])]^fullData[82^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[1])], fullData[100^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])]+fullData[116^ /*line ARMrBo6HPYh.go:1*/ int(idxKey[0])])
			return /*line ARMrBo6HPYh.go:1*/ string(data)
		}(), HCy4ZyvVELw /*line NBf5QEsKI7U.go:1*/, iPfjW8i91hC.ChainId(), gDp0rg.chainId)
	}
	return /*line J1LICvq8ES.go:1*/ dAgH1L( /*line IeMoxTrPxO.go:1*/ gDp0rg.Hash(iPfjW8i91hC), LhajwZArxs, ChPRkRQT, ICKLPq3pU, true)
}

func (gDp0rg eip2930Signer) SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (LhajwZArxs, ChPRkRQT, ICKLPq3pU *big.Int, rfteMJju error) {
	switch hqz9d9MXjE := iPfjW8i91hC.inner.(type) {
	case *LegacyTx:
		return /*line oZSm38yTWH.go:1*/ gDp0rg.EIP155Signer.SignatureValues(iPfjW8i91hC, yboZCx6E)
	case *AccessListTx:

		if /*line V3QklcL.go:1*/ hqz9d9MXjE.ChainID.Sign() != 0 && /*line MT3d8bjn7qf.go:1*/ hqz9d9MXjE.ChainID.Cmp(gDp0rg.chainId) != 0 {
			return nil, nil, nil /*line sCsvmeDhgiJO.go:1*/, fmt.Errorf(func() /*line BgtJkVcQCb.go:1*/ string {
				key := [] /*line BgtJkVcQCb.go:1*/ byte("\x95\xa3f\xe1T\xb0%\xd4\xdb.>\x9e\x9aK\xfc\xa9#\x8e|")
				data := [] /*line BgtJkVcQCb.go:1*/ byte("\xba\x1a\xa0\x01\xbc\x11\x9b9\xfbS\xa2\xbe\x11\xacj\x1dC\xb3\xe0")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return /*line BgtJkVcQCb.go:1*/ string(data)
			}(), HCy4ZyvVELw, hqz9d9MXjE.ChainID, gDp0rg.chainId)
		}
		LhajwZArxs, ChPRkRQT, _ = /*line XkgCgzRco.go:1*/ wHPfzXWnSUX(yboZCx6E)
		ICKLPq3pU = /*line vegkD2lfj9nP.go:1*/ big.NewInt( /*line GO4TIK.go:1*/ int64(yboZCx6E[64]))
	default:
		return nil, nil, nil, B0ZTT4Gv0t
	}
	return LhajwZArxs, ChPRkRQT, ICKLPq3pU, nil
}

func (gDp0rg eip2930Signer) Hash(iPfjW8i91hC *Transaction) common.Hash {
	switch /*line gm5Laxfq.go:1*/ iPfjW8i91hC.Type() {
	case LegacyTxType:
		return /*line D6ZoVr6.go:1*/ gDp0rg.EIP155Signer.Hash(iPfjW8i91hC)
	case AccessListTxType:
		return /*line Bb2zWkEYOoq.go:1*/ wrj50W(
			/*line A7WYy6.go:1*/ iPfjW8i91hC.Type(),
			[]interface{}{
				gDp0rg.chainId,
				/*line WJK0an8eSyn.go:1*/ iPfjW8i91hC.Nonce(),
				/*line AhyRNUHHU8.go:1*/ iPfjW8i91hC.GasPrice(),
				/*line uvHACL_vf4.go:1*/ iPfjW8i91hC.Gas(),
				/*line ALgdYXgw.go:1*/ iPfjW8i91hC.To(),
				/*line T0l2b10.go:1*/ iPfjW8i91hC.Value(),
				/*line Src5UoVOw.go:1*/ iPfjW8i91hC.Data(),
				/*line SNxSvD.go:1*/ iPfjW8i91hC.AccessList(),
			})
	default:

		return common.Hash{}
	}
}

type EIP155Signer struct {
	chainId, chainIdMul *big.Int
}

func GVrt8eeUJ(tkkFB3 *big.Int) EIP155Signer {
	if tkkFB3 == nil {
		tkkFB3 = /*line cnPni_zS.go:1*/ new(big.Int)
	}
	return EIP155Signer{
		chainId: tkkFB3,
		chainIdMul:/*line Bq9bBcfcII8H.go:1*/ new(big.Int).Mul(tkkFB3 /*line pjmC2d4m.go:1*/, big.NewInt(2)),
	}
}

func (gDp0rg EIP155Signer) ChainID() *big.Int {
	return gDp0rg.chainId
}

func (gDp0rg EIP155Signer) Equal(gjNSUeDUec NlfWgA7) bool {
	y_M2k3k8N, bCDRCu := gjNSUeDUec.(EIP155Signer)
	return bCDRCu && /*line DhOVLPW.go:1*/ y_M2k3k8N.chainId.Cmp(gDp0rg.chainId) == 0
}

var v_eccb = /*line tCErmJca.go:1*/ big.NewInt(8)

func (gDp0rg EIP155Signer) Sender(iPfjW8i91hC *Transaction) (common.Address, error) {
	if /*line yXoSaPZrWb3G.go:1*/ iPfjW8i91hC.Type() != LegacyTxType {
		return common.Address{}, B0ZTT4Gv0t
	}
	if ! /*line lF9kla7.go:1*/ iPfjW8i91hC.Protected() {
		return /*line zv4gp53vlX8.go:1*/ H6gCgPO3{}.Sender(iPfjW8i91hC)
	}
	if /*line CcGKDBn.go:1*/ iPfjW8i91hC.ChainId().Cmp(gDp0rg.chainId) != 0 {
		return common.Address{} /*line HoqGSmj.go:1*/, fmt.Errorf(func() /*line EZvvTYl.go:1*/ string {
			key := [] /*line EZvvTYl.go:1*/ byte("yӟ\x9b\uedd1o/\x84\x8b\xdb\x0e6\xc9.\a\x99f")
			data := [] /*line EZvvTYl.go:1*/ byte("\x9eJٻV\x18\a\xd4O\xa9\xef\xfb\x85\x977\xa2'\xbe\xca")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line EZvvTYl.go:1*/ string(data)
		}(), HCy4ZyvVELw /*line GpV_SBFiN.go:1*/, iPfjW8i91hC.ChainId(), gDp0rg.chainId)
	}
	ICKLPq3pU, LhajwZArxs, ChPRkRQT := /*line aYQLMqC.go:1*/ iPfjW8i91hC.RawSignatureValues()
	ICKLPq3pU = /*line mViadivgyPE.go:1*/ new(big.Int).Sub(ICKLPq3pU, gDp0rg.chainIdMul)
	/*line y0cN6_wc.go:1*/ ICKLPq3pU.Sub(ICKLPq3pU, v_eccb)
	return /*line FB_hDXt.go:1*/ dAgH1L( /*line Ksi5LQq.go:1*/ gDp0rg.Hash(iPfjW8i91hC), LhajwZArxs, ChPRkRQT, ICKLPq3pU, true)
}

func (gDp0rg EIP155Signer) SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (LhajwZArxs, ChPRkRQT, ICKLPq3pU *big.Int, rfteMJju error) {
	if /*line nLyyoZnPdheF.go:1*/ iPfjW8i91hC.Type() != LegacyTxType {
		return nil, nil, nil, B0ZTT4Gv0t
	}
	LhajwZArxs, ChPRkRQT, ICKLPq3pU = /*line bio2bf.go:1*/ wHPfzXWnSUX(yboZCx6E)
	if /*line FFyhWrHbmb8c.go:1*/ gDp0rg.chainId.Sign() != 0 {
		ICKLPq3pU = /*line xa4kSNm9.go:1*/ big.NewInt( /*line UlthqSCG.go:1*/ int64(yboZCx6E[64] + 35))
		/*line r74dC3OMdM7z.go:1*/ ICKLPq3pU.Add(ICKLPq3pU, gDp0rg.chainIdMul)
	}
	return LhajwZArxs, ChPRkRQT, ICKLPq3pU, nil
}

func (gDp0rg EIP155Signer) Hash(iPfjW8i91hC *Transaction) common.Hash {
	return /*line zn_0U4BUjfXb.go:1*/ uWZWEzDAB([]interface{}{
		/*line GyLBjVSm1.go:1*/ iPfjW8i91hC.Nonce(),
		/*line KfxaHXqQ1.go:1*/ iPfjW8i91hC.GasPrice(),
		/*line xawGJaFGywDa.go:1*/ iPfjW8i91hC.Gas(),
		/*line FHN4JwL.go:1*/ iPfjW8i91hC.To(),
		/*line SUk3wFmTljNG.go:1*/ iPfjW8i91hC.Value(),
		/*line DNFbKT.go:1*/ iPfjW8i91hC.Data(),
		gDp0rg.chainId /*line NwZeLb.go:1*/, uint(0) /*line FNo3jYriE.go:1*/, uint(0),
	})
}

type H6gCgPO3 struct{ We6m1j6t }

func (gDp0rg H6gCgPO3) ChainID() *big.Int {
	return nil
}

func (gDp0rg H6gCgPO3) Equal(gjNSUeDUec NlfWgA7) bool {
	_, bCDRCu := gjNSUeDUec.(H6gCgPO3)
	return bCDRCu
}

func (zH2LoA4 H6gCgPO3) SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (fWuxre8U, gDp0rg, pW1UMRn1s *big.Int, rfteMJju error) {
	return /*line MTfe5SJkeI.go:1*/ zH2LoA4.We6m1j6t.SignatureValues(iPfjW8i91hC, yboZCx6E)
}

func (zH2LoA4 H6gCgPO3) Sender(iPfjW8i91hC *Transaction) (common.Address, error) {
	if /*line HK_bR5W6CL.go:1*/ iPfjW8i91hC.Type() != LegacyTxType {
		return common.Address{}, B0ZTT4Gv0t
	}
	pW1UMRn1s, fWuxre8U, gDp0rg := /*line n9MaJSfQkMu.go:1*/ iPfjW8i91hC.RawSignatureValues()
	return /*line oRwuvDH.go:1*/ dAgH1L( /*line CUyLlIntTCr.go:1*/ zH2LoA4.Hash(iPfjW8i91hC), fWuxre8U, gDp0rg, pW1UMRn1s, true)
}

type We6m1j6t struct{}

func (gDp0rg We6m1j6t) ChainID() *big.Int {
	return nil
}

func (gDp0rg We6m1j6t) Equal(gjNSUeDUec NlfWgA7) bool {
	_, bCDRCu := gjNSUeDUec.(We6m1j6t)
	return bCDRCu
}

func (bKdFgl We6m1j6t) Sender(iPfjW8i91hC *Transaction) (common.Address, error) {
	if /*line r3JYrJ.go:1*/ iPfjW8i91hC.Type() != LegacyTxType {
		return common.Address{}, B0ZTT4Gv0t
	}
	pW1UMRn1s, fWuxre8U, gDp0rg := /*line TX3bhh.go:1*/ iPfjW8i91hC.RawSignatureValues()
	return /*line bP64Da7H_.go:1*/ dAgH1L( /*line ANuthZ69.go:1*/ bKdFgl.Hash(iPfjW8i91hC), fWuxre8U, gDp0rg, pW1UMRn1s, false)
}

func (bKdFgl We6m1j6t) SignatureValues(iPfjW8i91hC *Transaction, yboZCx6E []byte) (fWuxre8U, gDp0rg, pW1UMRn1s *big.Int, rfteMJju error) {
	if /*line nnLQ6AZnlIH9.go:1*/ iPfjW8i91hC.Type() != LegacyTxType {
		return nil, nil, nil, B0ZTT4Gv0t
	}
	fWuxre8U, gDp0rg, pW1UMRn1s = /*line LfdEhay.go:1*/ wHPfzXWnSUX(yboZCx6E)
	return fWuxre8U, gDp0rg, pW1UMRn1s, nil
}

func (bKdFgl We6m1j6t) Hash(iPfjW8i91hC *Transaction) common.Hash {
	return /*line LEF5z6QBsLi.go:1*/ uWZWEzDAB([]interface{}{
		/*line cNNpQM.go:1*/ iPfjW8i91hC.Nonce(),
		/*line TfR6TpSfrJi.go:1*/ iPfjW8i91hC.GasPrice(),
		/*line BnJkA9Gz6LWq.go:1*/ iPfjW8i91hC.Gas(),
		/*line i8YvfE.go:1*/ iPfjW8i91hC.To(),
		/*line GFCPBNGZFx5K.go:1*/ iPfjW8i91hC.Value(),
		/*line Nt9hMty5.go:1*/ iPfjW8i91hC.Data(),
	})
}

func wHPfzXWnSUX(yboZCx6E []byte) (fWuxre8U, gDp0rg, pW1UMRn1s *big.Int) {
	if /*line glqa55s.go:1*/ len(yboZCx6E) != crypto.SignatureLength {
		/*line C_Y0xrGCvI.go:1*/ panic( /*line IC9xNa.go:1*/ fmt.Sprintf(func() /*line lwbB5Z.go:1*/ string {
			data := /*line lwbB5Z.go:1*/ make([]byte, 0, 42)
			i := 6
			decryptKey := 230
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 11:
					data = /*line lwbB5Z.go:1*/ append(data, 232)
					i = 17
				case 7:
					i = 15
					data = /*line lwbB5Z.go:1*/ append(data, 112)
				case 15:
					i = 18
					data = /*line lwbB5Z.go:1*/ append(data, "ķ\xb6"...,
					)
				case 12:
					data = /*line lwbB5Z.go:1*/ append(data, "\x9f\xf3\xde"...,
					)
					i = 11
				case 8:
					i = 14
					data = /*line lwbB5Z.go:1*/ append(data, "\xab\xb4\xb6c"...,
					)
				case 14:
					data = /*line lwbB5Z.go:1*/ append(data, "e\xa5"...,
					)
					i = 20
				case 1:
					for y := range data {
						data[y] = data[y] - /*line lwbB5Z.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 20:
					data = /*line lwbB5Z.go:1*/ append(data, 170)
					i = 12
				case 9:
					i = 2
					data = /*line lwbB5Z.go:1*/ append(data, 116)
				case 2:
					data = /*line lwbB5Z.go:1*/ append(data, "\xbb\xc1\xc5"...,
					)
					i = 7
				case 5:
					i = 4
					data = /*line lwbB5Z.go:1*/ append(data, "\xcb\xc1{\xcb"...,
					)
				case 16:
					i = 1
					data = /*line lwbB5Z.go:1*/ append(data, "\x9e\xda"...,
					)
				case 13:
					data = /*line lwbB5Z.go:1*/ append(data, "\xae\x80g"...,
					)
					i = 8
				case 18:
					i = 10
					data = /*line lwbB5Z.go:1*/ append(data, "\xba\xae"...,
					)
				case 3:
					i = 9
					data = /*line lwbB5Z.go:1*/ append(data, 188)
				case 4:
					i = 3
					data = /*line lwbB5Z.go:1*/ append(data, "\xc2\xd0"...,
					)
				case 17:
					i = 19
					data = /*line lwbB5Z.go:1*/ append(data, 239)
				case 10:
					i = 13
					data = /*line lwbB5Z.go:1*/ append(data, "\xbe\xc0\xba"...,
					)
				case 19:
					data = /*line lwbB5Z.go:1*/ append(data, 152)
					i = 16
				case 6:
					i = 5
					data = /*line lwbB5Z.go:1*/ append(data, "\xd5\xd1\xcb"...,
					)
				}
			}
			return /*line lwbB5Z.go:1*/ string(data)
		}(), /*line iUMa3tPC7pn.go:1*/ len(yboZCx6E), crypto.SignatureLength))
	}
	fWuxre8U = /*line HEnSVXNSx2O.go:1*/ new(big.Int).SetBytes(yboZCx6E[:32])
	gDp0rg = /*line JrzSvFbgAd.go:1*/ new(big.Int).SetBytes(yboZCx6E[32:64])
	pW1UMRn1s = /*line U9Blx0cE46G.go:1*/ new(big.Int).SetBytes([]byte{yboZCx6E[64] + 27})
	return fWuxre8U, gDp0rg, pW1UMRn1s
}

func dAgH1L(hl_to7t common.Hash, LhajwZArxs, ChPRkRQT, FcTIuqIt *big.Int, d9Y3yS2 bool) (common.Address, error) {
	if /*line MMXtRYO_W4a.go:1*/ FcTIuqIt.BitLen() > 8 {
		return common.Address{}, CqKl3D50bxiN
	}
	ICKLPq3pU := /*line BRPVCt6o.go:1*/ byte( /*line H5RI9vRa.go:1*/ FcTIuqIt.Uint64() - 27)
	if ! /*line QCU83eQ.go:1*/ crypto.ValidateSignatureValues(ICKLPq3pU, LhajwZArxs, ChPRkRQT, d9Y3yS2) {
		return common.Address{}, CqKl3D50bxiN
	}

	fWuxre8U, gDp0rg := /*line boM6dk.go:1*/ LhajwZArxs.Bytes() /*line IXXy75aGHG3r.go:1*/, ChPRkRQT.Bytes()
	yboZCx6E := /*line Pa4rV7N5OYE.go:1*/ make([]byte, crypto.SignatureLength)
	/*line YkwT1nPS8rC.go:1*/ copy(yboZCx6E[32- /*line wvxeWI0G8t.go:1*/ len(fWuxre8U):32], fWuxre8U)
	/*line NbbomS9BAaEX.go:1*/ copy(yboZCx6E[64- /*line MF3xV0Y.go:1*/ len(gDp0rg):64], gDp0rg)
	yboZCx6E[64] = ICKLPq3pU

	zSp3noK, rfteMJju := /*line JZsN9vUa.go:1*/ crypto.Ecrecover(hl_to7t[:], yboZCx6E)
	if rfteMJju != nil {
		return common.Address{}, rfteMJju
	}
	if /*line haGM4_.go:1*/ len(zSp3noK) == 0 || zSp3noK[0] != 4 {
		return common.Address{} /*line E1etFW.go:1*/, errors.New(func() /*line G93wVRR.go:1*/ string {
			fullData := [] /*line G93wVRR.go:1*/ byte("\xf8\xe0\xdd5\x0f0\xf0\x8c:\x88~\x8a\x1a\xf9\xe9\b\xb7\xeeP\re\x7f\x9d\xb7\xd5Ť4\xf6\x0e\x80C\xaeW^F")
			idxKey := [] /*line G93wVRR.go:1*/ byte("ϧ\x94J\xe6\x8dᏇ\x7fF")
			data := /*line G93wVRR.go:1*/ make([]byte, 0, 19)
			data = /*line G93wVRR.go:1*/ append(data, fullData[81^ /*line G93wVRR.go:1*/ int(idxKey[3])]+fullData[73^ /*line G93wVRR.go:1*/ int(idxKey[3])], fullData[102^ /*line G93wVRR.go:1*/ int(idxKey[9])]-fullData[94^ /*line G93wVRR.go:1*/ int(idxKey[9])], fullData[158^ /*line G93wVRR.go:1*/ int(idxKey[2])]-fullData[155^ /*line G93wVRR.go:1*/ int(idxKey[2])], fullData[232^ /*line G93wVRR.go:1*/ int(idxKey[4])]-fullData[239^ /*line G93wVRR.go:1*/ int(idxKey[4])], fullData[90^ /*line G93wVRR.go:1*/ int(idxKey[10])]-fullData[77^ /*line G93wVRR.go:1*/ int(idxKey[10])], fullData[220^ /*line G93wVRR.go:1*/ int(idxKey[0])]-fullData[213^ /*line G93wVRR.go:1*/ int(idxKey[0])], fullData[195^ /*line G93wVRR.go:1*/ int(idxKey[6])]^fullData[233^ /*line G93wVRR.go:1*/ int(idxKey[6])], fullData[88^ /*line G93wVRR.go:1*/ int(idxKey[3])]-fullData[79^ /*line G93wVRR.go:1*/ int(idxKey[3])], fullData[66^ /*line G93wVRR.go:1*/ int(idxKey[10])]^fullData[83^ /*line G93wVRR.go:1*/ int(idxKey[10])], fullData[153^ /*line G93wVRR.go:1*/ int(idxKey[5])]-fullData[139^ /*line G93wVRR.go:1*/ int(idxKey[5])], fullData[159^ /*line G93wVRR.go:1*/ int(idxKey[8])]^fullData[144^ /*line G93wVRR.go:1*/ int(idxKey[8])], fullData[147^ /*line G93wVRR.go:1*/ int(idxKey[2])]^fullData[149^ /*line G93wVRR.go:1*/ int(idxKey[2])], fullData[172^ /*line G93wVRR.go:1*/ int(idxKey[7])]-fullData[141^ /*line G93wVRR.go:1*/ int(idxKey[7])], fullData[237^ /*line G93wVRR.go:1*/ int(idxKey[6])]-fullData[241^ /*line G93wVRR.go:1*/ int(idxKey[6])], fullData[146^ /*line G93wVRR.go:1*/ int(idxKey[7])]-fullData[158^ /*line G93wVRR.go:1*/ int(idxKey[7])], fullData[180^ /*line G93wVRR.go:1*/ int(idxKey[2])]-fullData[139^ /*line G93wVRR.go:1*/ int(idxKey[2])], fullData[80^ /*line G93wVRR.go:1*/ int(idxKey[10])]^fullData[70^ /*line G93wVRR.go:1*/ int(idxKey[10])], fullData[248^ /*line G93wVRR.go:1*/ int(idxKey[4])]+fullData[235^ /*line G93wVRR.go:1*/ int(idxKey[4])])
			return /*line G93wVRR.go:1*/ string(data)
		}())
	}
	var yOdETLh4 common.Address
	/*line xjNC3FdqlcH3.go:1*/ copy(yOdETLh4[:] /*line fNnoPz.go:1*/, crypto.Keccak256(zSp3noK[1:])[12:])
	return yOdETLh4, nil
}

func sikVyPV4yp(pW1UMRn1s *big.Int) *big.Int {
	if /*line gmKewftniy.go:1*/ pW1UMRn1s.BitLen() <= 64 {
		pW1UMRn1s := /*line VLS3b1c3q.go:1*/ pW1UMRn1s.Uint64()
		if pW1UMRn1s == 27 || pW1UMRn1s == 28 {
			return /*line iIiO3rxHmGD.go:1*/ new(big.Int)
		}
		return /*line RjGK_TYhVYvD.go:1*/ new(big.Int).SetUint64((pW1UMRn1s - 35) / 2)
	}
	pW1UMRn1s = /*line wWfhUyf0m.go:1*/ new(big.Int).Sub(pW1UMRn1s /*line Ft5_aDjF.go:1*/, big.NewInt(35))
	return /*line LdEDzNTsgxg8.go:1*/ pW1UMRn1s.Div(pW1UMRn1s /*line opWLua.go:1*/, big.NewInt(2))
}

var _ = ecdsa.GenerateKey
var _ = errors.As
var _ = fmt.Append

const _ = big.Above

var _ = common.AbsolutePath
var _ = crypto.CompressPubkey
var _ = params.AllCliqueProtocolChanges
