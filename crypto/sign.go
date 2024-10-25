//line :1
package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	types "unishard/types"

	"github.com/ethereum/go-ethereum/common"
)

const (
	BLS_BLS12381    = "BLS_BLS12381"
	ECDSA_P256      = "ECDSA_P256"
	ECDSA_SECp256k1 = "ECDSA_SECp256k1"
)

var jYDXng *ecdsa.PrivateKey
var nRHqoC0RfsK_ *ecdsa.PublicKey

var lYn7kpz []HtOBJwcl
var sid5XmISl []H9kkE4IQM

var cU1mItV6O1 []HtOBJwcl
var i7nKmshfEUp []H9kkE4IQM

type HtOBJwcl interface {
	Algorithm() string

	Sign([]byte, K3fvz3wD) (MqlfmVE9d, error)

	PrivateKey() *ecdsa.PrivateKey
}

type H9kkE4IQM interface {
	Algorithm() string

	Verify(MqlfmVE9d, KDSldlLG06As) (bool, error)

	Encode() ([]byte, error)
	PublicKey() *ecdsa.PublicKey
}

func vd9I9J7lF(fSlpSB0To *ecdsa.PrivateKey, t0Vj5K *ecdsa.PublicKey) (string, string) {
	epcT49, _ := /*line IH3Y8D1MO.go:1*/ x509.MarshalECPrivateKey(fSlpSB0To)
	lC8uBJo := /*line BgeHF7OJ4a.go:1*/ pem.EncodeToMemory(&pem.Block{Type: func() /*line _vWuVe6Hw.go:1*/ string {
		data := /*line _vWuVe6Hw.go:1*/ make([]byte, 0, 12)
		i := 3
		decryptKey := 221
		for counter := 0; i != 10; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 12:
				i = 7
				data = /*line _vWuVe6Hw.go:1*/ append(data, 210)
			case 8:
				data = /*line _vWuVe6Hw.go:1*/ append(data, 212)
				i = 0
			case 7:
				data = /*line _vWuVe6Hw.go:1*/ append(data, 194)
				i = 9
			case 0:
				i = 10
				for y := range data {
					data[y] = data[y] ^ /*line _vWuVe6Hw.go:1*/ byte(decryptKey^y)
				}
			case 5:
				i = 4
				data = /*line _vWuVe6Hw.go:1*/ append(data, 196)
			case 9:
				i = 5
				data = /*line _vWuVe6Hw.go:1*/ append(data, 214)
			case 3:
				i = 11
				data = /*line _vWuVe6Hw.go:1*/ append(data, 215)
			case 11:
				i = 6
				data = /*line _vWuVe6Hw.go:1*/ append(data, 212)
			case 2:
				i = 8
				data = /*line _vWuVe6Hw.go:1*/ append(data, 203)
			case 4:
				i = 1
				data = /*line _vWuVe6Hw.go:1*/ append(data, 160)
			case 1:
				data = /*line _vWuVe6Hw.go:1*/ append(data, 196)
				i = 2
			case 6:
				i = 12
				data = /*line _vWuVe6Hw.go:1*/ append(data, 204)
			}
		}
		return /*line _vWuVe6Hw.go:1*/ string(data)
	}(), Bytes: epcT49})

	bTBGd3V08IlM, _ := /*line FvtvsR.go:1*/ x509.MarshalPKIXPublicKey(t0Vj5K)
	vGuNtzn := /*line UNRy6j.go:1*/ pem.EncodeToMemory(&pem.Block{Type: func() /*line HVrTX1.go:1*/ string {
		fullData := [] /*line HVrTX1.go:1*/ byte("/\xe8\x7f\x8c\xcdoe\xdf\xc9h\xe6\x00q \x94q\xb7+:'")
		idxKey := [] /*line HVrTX1.go:1*/ byte("\xe6\x12\xa4\xbf\xba")
		data := /*line HVrTX1.go:1*/ make([]byte, 0, 11)
		data = /*line HVrTX1.go:1*/ append(data, fullData[184^ /*line HVrTX1.go:1*/ int(idxKey[3])]+fullData[179^ /*line HVrTX1.go:1*/ int(idxKey[3])], fullData[244^ /*line HVrTX1.go:1*/ int(idxKey[0])]^fullData[227^ /*line HVrTX1.go:1*/ int(idxKey[0])], fullData[224^ /*line HVrTX1.go:1*/ int(idxKey[0])]^fullData[245^ /*line HVrTX1.go:1*/ int(idxKey[0])], fullData[187^ /*line HVrTX1.go:1*/ int(idxKey[3])]+fullData[189^ /*line HVrTX1.go:1*/ int(idxKey[3])], fullData[230^ /*line HVrTX1.go:1*/ int(idxKey[0])]-fullData[236^ /*line HVrTX1.go:1*/ int(idxKey[0])], fullData[181^ /*line HVrTX1.go:1*/ int(idxKey[2])]^fullData[173^ /*line HVrTX1.go:1*/ int(idxKey[2])], fullData[180^ /*line HVrTX1.go:1*/ int(idxKey[3])]^fullData[178^ /*line HVrTX1.go:1*/ int(idxKey[3])], fullData[170^ /*line HVrTX1.go:1*/ int(idxKey[4])]+fullData[180^ /*line HVrTX1.go:1*/ int(idxKey[4])], fullData[26^ /*line HVrTX1.go:1*/ int(idxKey[1])]^fullData[17^ /*line HVrTX1.go:1*/ int(idxKey[1])], fullData[187^ /*line HVrTX1.go:1*/ int(idxKey[4])]+fullData[181^ /*line HVrTX1.go:1*/ int(idxKey[4])])
		return /*line HVrTX1.go:1*/ string(data)
	}(), Bytes: bTBGd3V08IlM})

	return /*line dpaYMt2YcxP.go:1*/ string(lC8uBJo) /*line _rAyrPFn.go:1*/, string(vGuNtzn)
}

func kglRi1viUo(lC8uBJo string, vGuNtzn string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	cbp8vZeGjcFA, _ := /*line q85PqVxxxsb.go:1*/ pem.Decode([] /*line skBY_DSYi.go:1*/ byte(lC8uBJo))
	epcT49 := cbp8vZeGjcFA.Bytes
	fSlpSB0To, _ := /*line tyWNvMa.go:1*/ x509.ParseECPrivateKey(epcT49)

	exkD9dSI, _ := /*line BQNWPz.go:1*/ pem.Decode([] /*line Fn3ncz.go:1*/ byte(vGuNtzn))
	bTBGd3V08IlM := exkD9dSI.Bytes
	c88w65UGV, _ := /*line VbRlD2lMBTg.go:1*/ x509.ParsePKIXPublicKey(bTBGd3V08IlM)
	t0Vj5K := c88w65UGV.(*ecdsa.PublicKey)

	return fSlpSB0To, t0Vj5K
}

func HIZ_hi() error {
	v6UI6RSaE := func() /*line zkBDZI0xZIl.go:1*/ string {
		seed := /*line zkBDZI0xZIl.go:1*/ byte(102)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line zkBDZI0xZIl.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line zkBDZI0xZIl.go:1*/ fnc(75)(156)(96)(128)(0)(111)(217)(50)(238)(219)(80)(144)(2)(27)(59)(233)(197)(19)(73)(249)(238)(192)(116)(224)(128)(0)(0)(39)(25)(37)(241)(192)(2)(20)(28)(48)(236)(218)(8)(42)(254)(254)(236)(174)(10)(100)(245)(226)(34)(253)(163)(72)(180)(66)(76)(158)(54)(206)(20)(78)(138)(0)(63)(231)(54)(237)(170)(48)(66)(248)(205)(79)(172)(33)(211)(14)(49)(222)(239)(79)(169)(46)(242)(200)(20)(40)(226)(50)(244)(214)(69)(143)(79)(213)(30)(194)(1)(37)(58)(248)(229)(215)(58)(214)(40)(214)(64)(201)(12)(203)(46)(79)(187)(176)(113)(143)(80)(186)(122)(242)(197)(26)(210)(25)(74)(231)(139)(85)(183)(126)(200)(177)(118)(221)(207)(31)(11)(27)(215)(61)(238)(170)(3)(60)(232)(3)(54)(210)(20)(164)(122)(216)(199)(3)(44)(5)(54)(184)(26)(171)(99)(199)(63)(233)(199)(125)(170)(21)(220)(17)(66)(203)(77)(167)(30)(255)(253)(254)(249)(239)(68)(206)(237)(2)(28)(255)(239)(46)(204)(67)(142)(47)(74)(252)(183)(89)(224)(128)(0)(0)(104)(219)(52)(132)(120)(242)(219)(59)(233)(197)(19)(73)(249)(238)(192)(116)(224)(128)(0)(0)
		return /*line zkBDZI0xZIl.go:1*/ string(data)
	}()
	fWthKymgNO := func() /*line Qp_SGVU.go:1*/ string {
		seed := /*line Qp_SGVU.go:1*/ byte(230)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line Qp_SGVU.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line Qp_SGVU.go:1*/ fnc(19)(38)(76)(152)(48)(117)(237)(220)(186)(121)(196)(184)(117)(215)(184)(109)(212)(133)(53)(100)(220)(140)(24)(48)(96)(192)(93)(253)(243)(11)(34)(18)(86)(142)(11)(25)(86)(151)(29)(107)(198)(82)(183)(108)(232)(216)(160)(66)(168)(59)(101)(251)(230)(146)(56)(109)(234)(230)(173)(103)(228)(162)(72)(142)(4)(63)(148)(251)(16)(241)(31)(254)(11)(255)(54)(66)(146)(34)(107)(222)(162)(45)(80)(193)(87)(249)(171)(104)(252)(183)(130)(202)(255)(211)(167)(123)(227)(175)(93)(176)(78)(177)(144)(24)(51)(80)(190)(104)(156)(94)(216)(133)(15)(58)(127)(214)(210)(99)(227)(207)(175)(105)(167)(61)(180)(115)(196)(161)(4)(77)(85)(227)(198)(143)(33)(66)(133)(3)(208)(226)(181)(116)(228)(201)(131)(248)(4)(213)(170)(135)(228)(200)(93)(221)(186)(116)(232)(208)(184)(121)(232)(172)(136)(21)(23)(56)(109)(212)(133)(53)(100)(220)(140)(24)(48)(96)(192)
		return /*line Qp_SGVU.go:1*/ string(data)
	}()
	ixv_COBPVvD, rntVhF1_K := /*line Mtl5nZDcI.go:1*/ kglRi1viUo(v6UI6RSaE, fWthKymgNO)
	jYDXng = ixv_COBPVvD
	nRHqoC0RfsK_ = rntVhF1_K
	return nil
}

func BFo6_c(nM3HxVbGk Pp__49cd, _X9xGSTi common.Hash, sUTM2ieid8Na []types.NodeID) (bool, error) {
	var nXejsF9of0ys bool
	var dQA_vo0K7 error
	for g0af0v3qZ7, _ := range sUTM2ieid8Na {
		nXejsF9of0ys, dQA_vo0K7 = /*line wx8JR5ag.go:1*/ SkrCuscT(nM3HxVbGk[g0af0v3qZ7] /*line uVXrfN02KG.go:1*/, IDToByte(_X9xGSTi))
		if dQA_vo0K7 != nil {
			return false, dQA_vo0K7
		}
		if !nXejsF9of0ys {
			return false, nil
		}
	}
	return true, nil
}

var _ = ecdsa.GenerateKey

const _ = x509.CANotAuthorizedForExtKeyUsage

var _ pem.Block

const _ = types.ABORT

var _ = common.AbsolutePath
