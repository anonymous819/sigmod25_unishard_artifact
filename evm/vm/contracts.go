//line :1
package vm

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/blake2b"
	"github.com/ethereum/go-ethereum/crypto/bls12381"
	"github.com/ethereum/go-ethereum/crypto/bn256"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/ripemd160"
)

type VeZhcup interface {
	RequiredGas(sIVM3RLPSKSh []byte) uint64
	Run(sIVM3RLPSKSh []byte) ([]byte, error)
}

var AwN8pSb35 = map[common.Address]VeZhcup{
	/*line Cm6RK2.go:1*/ common.BytesToAddress([]byte{1}): &dNNZJ_u{},
	/*line nlR75cA23NA.go:1*/ common.BytesToAddress([]byte{2}): &a5WVQJPdJaAC{},
	/*line sVvSUvY.go:1*/ common.BytesToAddress([]byte{3}): &le3W0S0CCS{},
	/*line ZGma6u.go:1*/ common.BytesToAddress([]byte{4}): &o31NjJItcX{},
}

var J0k7Wc9OJi = map[common.Address]VeZhcup{
	/*line ButF0fWsMkD.go:1*/ common.BytesToAddress([]byte{1}): &dNNZJ_u{},
	/*line EhERSxYbd.go:1*/ common.BytesToAddress([]byte{2}): &a5WVQJPdJaAC{},
	/*line bCg0aDG4Z.go:1*/ common.BytesToAddress([]byte{3}): &le3W0S0CCS{},
	/*line HfiypFv.go:1*/ common.BytesToAddress([]byte{4}): &o31NjJItcX{},
	/*line rw92AGO.go:1*/ common.BytesToAddress([]byte{5}): &tVYGZaqW3a{xI9wE2: false},
	/*line ywCQshkcN3q.go:1*/ common.BytesToAddress([]byte{6}): &u0cWnoUxkq1D{},
	/*line NlsnAm.go:1*/ common.BytesToAddress([]byte{7}): &rPzaYtsJ{},
	/*line rD9l6qo.go:1*/ common.BytesToAddress([]byte{8}): &e8otDl{},
}

var VN3ry_pWNww = map[common.Address]VeZhcup{
	/*line uohX01TwO.go:1*/ common.BytesToAddress([]byte{1}): &dNNZJ_u{},
	/*line F34KjITi7n6.go:1*/ common.BytesToAddress([]byte{2}): &a5WVQJPdJaAC{},
	/*line KkFSOC01X.go:1*/ common.BytesToAddress([]byte{3}): &le3W0S0CCS{},
	/*line xFparlqR9nan.go:1*/ common.BytesToAddress([]byte{4}): &o31NjJItcX{},
	/*line exW33ej.go:1*/ common.BytesToAddress([]byte{5}): &tVYGZaqW3a{xI9wE2: false},
	/*line jq09jHIfR.go:1*/ common.BytesToAddress([]byte{6}): &lEsxPQVLShM{},
	/*line ZMnWGa.go:1*/ common.BytesToAddress([]byte{7}): &uAwS_e1Y{},
	/*line PwulcK.go:1*/ common.BytesToAddress([]byte{8}): &lzLDwHcHQDpl{},
	/*line M7sItAY.go:1*/ common.BytesToAddress([]byte{9}): &zAsr27{},
}

var Agr7Uo = map[common.Address]VeZhcup{
	/*line qRzMLyQ3Ut.go:1*/ common.BytesToAddress([]byte{1}): &dNNZJ_u{},
	/*line mAQebw.go:1*/ common.BytesToAddress([]byte{2}): &a5WVQJPdJaAC{},
	/*line M4X977T_fOa.go:1*/ common.BytesToAddress([]byte{3}): &le3W0S0CCS{},
	/*line JIQmXpZe.go:1*/ common.BytesToAddress([]byte{4}): &o31NjJItcX{},
	/*line F0slConY1md3.go:1*/ common.BytesToAddress([]byte{5}): &tVYGZaqW3a{xI9wE2: true},
	/*line TUtipjllyZW.go:1*/ common.BytesToAddress([]byte{6}): &lEsxPQVLShM{},
	/*line ioRxZvWtuoiC.go:1*/ common.BytesToAddress([]byte{7}): &uAwS_e1Y{},
	/*line K_2ToUp4xS.go:1*/ common.BytesToAddress([]byte{8}): &lzLDwHcHQDpl{},
	/*line H3JQiDqT.go:1*/ common.BytesToAddress([]byte{9}): &zAsr27{},
}

var H5auzN = map[common.Address]VeZhcup{
	/*line s44Sai7BrWav.go:1*/ common.BytesToAddress([]byte{1}): &dNNZJ_u{},
	/*line GVNsvU.go:1*/ common.BytesToAddress([]byte{2}): &a5WVQJPdJaAC{},
	/*line qwJCBWdnjU.go:1*/ common.BytesToAddress([]byte{3}): &le3W0S0CCS{},
	/*line uzVCjYUB.go:1*/ common.BytesToAddress([]byte{4}): &o31NjJItcX{},
	/*line y0m1bHkTA.go:1*/ common.BytesToAddress([]byte{5}): &tVYGZaqW3a{xI9wE2: true},
	/*line ZDNrBWHmW1R.go:1*/ common.BytesToAddress([]byte{6}): &lEsxPQVLShM{},
	/*line EjYDdHsjZg.go:1*/ common.BytesToAddress([]byte{7}): &uAwS_e1Y{},
	/*line agps7i0y6x85.go:1*/ common.BytesToAddress([]byte{8}): &lzLDwHcHQDpl{},
	/*line CnnH1rIE.go:1*/ common.BytesToAddress([]byte{9}): &zAsr27{},
	/*line VRjk62yjFO6.go:1*/ common.BytesToAddress([]byte{0x0a}): &b6Inm46Q{},
}

var W2sr7_V3 = map[common.Address]VeZhcup{
	/*line dNXoZfgB.go:1*/ common.BytesToAddress([]byte{10}): &rvFNzWYNNj{},
	/*line X9h1bYM.go:1*/ common.BytesToAddress([]byte{11}): &hF6SEzF{},
	/*line M3YkAbr7B.go:1*/ common.BytesToAddress([]byte{12}): &vfwOEaC{},
	/*line WW1EPeSC9Mb.go:1*/ common.BytesToAddress([]byte{13}): &yUf2fKtK{},
	/*line hIRUMEob.go:1*/ common.BytesToAddress([]byte{14}): &vJqscRtml_{},
	/*line IDpanoIridt9.go:1*/ common.BytesToAddress([]byte{15}): &k8odgmq8WI{},
	/*line sRaBKqBvpX.go:1*/ common.BytesToAddress([]byte{16}): &fQxEr40b{},
	/*line YAcWUZ1V.go:1*/ common.BytesToAddress([]byte{17}): &hLGLU0Dn{},
	/*line T1r3_uBcs.go:1*/ common.BytesToAddress([]byte{18}): &gRXncStcih6U{},
}

var (
	O9bce3DtYWdM []common.Address
	Nbre163a5ojA []common.Address
	IS7Lk3si     []common.Address
	AlLAWXk      []common.Address
	PVW5rksGv0Yg []common.Address
)

func init() {
	for zsVIRN6k := range AwN8pSb35 {
		PVW5rksGv0Yg = /*line wyKguXT.go:1*/ append(PVW5rksGv0Yg, zsVIRN6k)
	}
	for zsVIRN6k := range J0k7Wc9OJi {
		AlLAWXk = /*line We52NlNXnlBD.go:1*/ append(AlLAWXk, zsVIRN6k)
	}
	for zsVIRN6k := range VN3ry_pWNww {
		IS7Lk3si = /*line gPNB2VwG6L.go:1*/ append(IS7Lk3si, zsVIRN6k)
	}
	for zsVIRN6k := range Agr7Uo {
		Nbre163a5ojA = /*line ENpint4up.go:1*/ append(Nbre163a5ojA, zsVIRN6k)
	}
	for zsVIRN6k := range H5auzN {
		O9bce3DtYWdM = /*line Ebe7nf4SHK.go:1*/ append(O9bce3DtYWdM, zsVIRN6k)
	}
}

func WNX2wLgB2IK(avYeWr0Mmle params.Rules) []common.Address {
	switch {
	case avYeWr0Mmle.IsCancun:
		return O9bce3DtYWdM
	case avYeWr0Mmle.IsBerlin:
		return Nbre163a5ojA
	case avYeWr0Mmle.IsIstanbul:
		return IS7Lk3si
	case avYeWr0Mmle.IsByzantium:
		return AlLAWXk
	default:
		return PVW5rksGv0Yg
	}
}

func WCpcUHmYy(gTaLKvaWJ VeZhcup, sIVM3RLPSKSh []byte, dB48qy uint64) (ag8Tdqxi4 []byte, c3jVruxp7Ow uint64, v1p7xkTykQgN error) {
	ma9RgGF0t4 := /*line FWDoPXWuJc.go:1*/ gTaLKvaWJ.RequiredGas(sIVM3RLPSKSh)
	if dB48qy < ma9RgGF0t4 {
		return nil, 0, ME5lmy
	}
	dB48qy -= ma9RgGF0t4
	dhOXtmt2Uz9, v1p7xkTykQgN := /*line JanlwK_aFt7U.go:1*/ gTaLKvaWJ.Run(sIVM3RLPSKSh)
	return dhOXtmt2Uz9, dB48qy, v1p7xkTykQgN
}

type dNNZJ_u struct{}

func (oPtEQcrX *dNNZJ_u) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.EcrecoverGas
}

func (oPtEQcrX *dNNZJ_u) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	const ecRecoverInputLength = 128

	sIVM3RLPSKSh = /*line W7FQa5.go:1*/ common.RightPadBytes(sIVM3RLPSKSh, ecRecoverInputLength)

	gGFrnFgT := /*line pITFIRd6.go:1*/ new(big.Int).SetBytes(sIVM3RLPSKSh[64:96])
	ek3u6QMaV := /*line tLXLnaght.go:1*/ new(big.Int).SetBytes(sIVM3RLPSKSh[96:128])
	gbvC5akim := sIVM3RLPSKSh[63] - 27

	if ! /*line a5o_tpbD.go:1*/ aFNbM1_(sIVM3RLPSKSh[32:63]) || ! /*line ZI7JtL.go:1*/ crypto.ValidateSignatureValues(gbvC5akim, gGFrnFgT, ek3u6QMaV, false) {
		return nil, nil
	}

	ecRXoaaNiz := /*line QmkG0mgX.go:1*/ make([]byte, 65)
	/*line Uvk5Pin.go:1*/ copy(ecRXoaaNiz, sIVM3RLPSKSh[64:128])
	ecRXoaaNiz[64] = gbvC5akim

	w_xLwnF, v1p7xkTykQgN := /*line EgkAbclX.go:1*/ crypto.Ecrecover(sIVM3RLPSKSh[:32], ecRXoaaNiz)

	if v1p7xkTykQgN != nil {
		return nil, nil
	}

	return /*line TBZFOPEZH.go:1*/ common.LeftPadBytes( /*line uqpc0d.go:1*/ crypto.Keccak256(w_xLwnF[1:])[12:], 32), nil
}

type a5WVQJPdJaAC struct{}

func (oPtEQcrX *a5WVQJPdJaAC) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return /*line CCiP8nXycP.go:1*/ uint64( /*line CxnWuitcH.go:1*/ len(sIVM3RLPSKSh)+31)/32*params.Sha256PerWordGas + params.Sha256BaseGas
}
func (oPtEQcrX *a5WVQJPdJaAC) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	eNKBLUKPi := /*line Qu7tG_W.go:1*/ sha256.Sum256(sIVM3RLPSKSh)
	return eNKBLUKPi[:], nil
}

type le3W0S0CCS struct{}

func (oPtEQcrX *le3W0S0CCS) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return /*line IlFqaf.go:1*/ uint64( /*line C_XLP113.go:1*/ len(sIVM3RLPSKSh)+31)/32*params.Ripemd160PerWordGas + params.Ripemd160BaseGas
}
func (oPtEQcrX *le3W0S0CCS) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	vW3fwSV9JWl := /*line JObp9damlr.go:1*/ ripemd160.New()
	/*line HxaAevNr35a.go:1*/ vW3fwSV9JWl.Write(sIVM3RLPSKSh)
	return /*line C0HRIK.go:1*/ common.LeftPadBytes( /*line M04aye0bwV.go:1*/ vW3fwSV9JWl.Sum(nil), 32), nil
}

type o31NjJItcX struct{}

func (oPtEQcrX *o31NjJItcX) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return /*line Drwkjn.go:1*/ uint64( /*line JnyCAzkNT.go:1*/ len(sIVM3RLPSKSh)+31)/32*params.IdentityPerWordGas + params.IdentityBaseGas
}
func (oPtEQcrX *o31NjJItcX) Run(kTH29c []byte) ([]byte, error) {
	return /*line l5B9tcc0a7.go:1*/ common.CopyBytes(kTH29c), nil
}

type tVYGZaqW3a struct {
	xI9wE2 bool
}

var (
	y_M1FClUrWEw = /*line TrwryNJdEoN.go:1*/ big.NewInt(1)
	avGeszl      = /*line obL7_Sdw5.go:1*/ big.NewInt(3)
	iG4W2eY      = /*line VJqYgloV9U.go:1*/ big.NewInt(4)
	wPQ0J5z      = /*line tSWxuexwqj1.go:1*/ big.NewInt(7)
	tDlfSa10     = /*line cWJNB5vYT.go:1*/ big.NewInt(8)
	aReziea      = /*line C8JeSljh.go:1*/ big.NewInt(16)
	rB1__KHUhrF  = /*line TheJh4czLN.go:1*/ big.NewInt(20)
	hAVadq5EaG   = /*line pikKJCmSsUt9.go:1*/ big.NewInt(32)
	aagylA       = /*line CPmzmp54rrZx.go:1*/ big.NewInt(64)
	lNn5LKfra    = /*line u1_79yZ10n2.go:1*/ big.NewInt(96)
	sJhVpCDF     = /*line RfXy5hw5GfC.go:1*/ big.NewInt(480)
	aK9qBRBtoaGe = /*line nokyq3U.go:1*/ big.NewInt(1024)
	mbFOkOo      = /*line cIy4w0x.go:1*/ big.NewInt(3072)
	cRRFBG       = /*line LiLUqDlwGg.go:1*/ big.NewInt(199680)
)

func wsQl4Xqf(hpgMKBm *big.Int) *big.Int {
	switch {
	case /*line fbwQKCjhQ.go:1*/ hpgMKBm.Cmp(aagylA) <= 0:
		/*line I4in3qF88.go:1*/ hpgMKBm.Mul(hpgMKBm, hpgMKBm)
	case /*line Cl4hOsgBO.go:1*/ hpgMKBm.Cmp(aK9qBRBtoaGe) <= 0:

		hpgMKBm = /*line DHYdfbJED.go:1*/ new(big.Int).Add(
			/*line EYBP4YER.go:1*/ new(big.Int).Div( /*line gDc3zo9DQb6N.go:1*/ new(big.Int).Mul(hpgMKBm, hpgMKBm), iG4W2eY),
			/*line ZKYTmaOqN4t.go:1*/ new(big.Int).Sub( /*line ZVFy0KBhsOSY.go:1*/ new(big.Int).Mul(lNn5LKfra, hpgMKBm), mbFOkOo),
		)
	default:

		hpgMKBm = /*line qrCnE7WfE.go:1*/ new(big.Int).Add(
			/*line FuuK8a9SJ.go:1*/ new(big.Int).Div( /*line C7JuRp6.go:1*/ new(big.Int).Mul(hpgMKBm, hpgMKBm), aReziea),
			/*line DqaRwp.go:1*/ new(big.Int).Sub( /*line VsuXv8Ce7.go:1*/ new(big.Int).Mul(sJhVpCDF, hpgMKBm), cRRFBG),
		)
	}
	return hpgMKBm
}

func (oPtEQcrX *tVYGZaqW3a) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	var (
		nw6_fH = /*line BJ3k2a3i0GkW.go:1*/ new(big.Int).SetBytes( /*line es2W1nw8x49b.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 0, 32))
		cuKNXu = /*line nyaWw3UgMf.go:1*/ new(big.Int).SetBytes( /*line ozynoCuDK5v.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 32, 32))
		bSdZdr = /*line de3stIDw3.go:1*/ new(big.Int).SetBytes( /*line ACZN4HFz.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 64, 32))
	)
	if /*line k7tYbw.go:1*/ len(sIVM3RLPSKSh) > 96 {
		sIVM3RLPSKSh = sIVM3RLPSKSh[96:]
	} else {
		sIVM3RLPSKSh = sIVM3RLPSKSh[:0]
	}

	var lC6ndo7wga *big.Int
	if /*line H7DA6wIM1X.go:1*/ big.NewInt( /*line aHa0UW.go:1*/ int64( /*line SKOa5UX.go:1*/ len(sIVM3RLPSKSh))).Cmp(nw6_fH) <= 0 {
		lC6ndo7wga = /*line JBO_GXP6P.go:1*/ new(big.Int)
	} else {
		if /*line AN9Rw_eilJc.go:1*/ cuKNXu.Cmp(hAVadq5EaG) > 0 {
			lC6ndo7wga = /*line QMEgt8QYcH.go:1*/ new(big.Int).SetBytes( /*line nlyBwUJ.go:1*/ m8O5lZzG(sIVM3RLPSKSh /*line rcP_sKiCPMJ.go:1*/, nw6_fH.Uint64(), 32))
		} else {
			lC6ndo7wga = /*line rJ75oSoV_P.go:1*/ new(big.Int).SetBytes( /*line ZaRA5c2Hqv.go:1*/ m8O5lZzG(sIVM3RLPSKSh /*line CMFmrjxB09.go:1*/, nw6_fH.Uint64() /*line oI3m012B.go:1*/, cuKNXu.Uint64()))
		}
	}

	var fqa2jV4Ny int
	if b5Fd9QE7Qf7v := /*line jNt5B6Bjz.go:1*/ lC6ndo7wga.BitLen(); b5Fd9QE7Qf7v > 0 {
		fqa2jV4Ny = b5Fd9QE7Qf7v - 1
	}
	d25OXcYBofJ5 := /*line vQoU2v.go:1*/ new(big.Int)
	if /*line GPSX8vg.go:1*/ cuKNXu.Cmp(hAVadq5EaG) > 0 {
		/*line wV8T26cPwwoe.go:1*/ d25OXcYBofJ5.Sub(cuKNXu, hAVadq5EaG)
		/*line KzjgBaQ.go:1*/ d25OXcYBofJ5.Mul(tDlfSa10, d25OXcYBofJ5)
	}
	/*line J_D3YMv9nk.go:1*/ d25OXcYBofJ5.Add(d25OXcYBofJ5 /*line JglMl37UtwGN.go:1*/, big.NewInt( /*line VthekLKoWvs.go:1*/ int64(fqa2jV4Ny)))

	a0SqxOtD := /*line nDZ_1ajZrFX.go:1*/ new(big.Int).Set( /*line um8BlSsWF0rz.go:1*/ math.BigMax(bSdZdr, nw6_fH))
	if oPtEQcrX.xI9wE2 {

		a0SqxOtD = /*line GgV4Z60I.go:1*/ a0SqxOtD.Add(a0SqxOtD, wPQ0J5z)
		a0SqxOtD = /*line bayj9dDVYqb.go:1*/ a0SqxOtD.Div(a0SqxOtD, tDlfSa10)
		/*line AuZjeBPWF.go:1*/ a0SqxOtD.Mul(a0SqxOtD, a0SqxOtD)

		/*line Z_iPYYf.go:1*/
		a0SqxOtD.Mul(a0SqxOtD /*line KPGMlE.go:1*/, math.BigMax(d25OXcYBofJ5, y_M1FClUrWEw))

		/*line QkSoriGAj.go:1*/
		a0SqxOtD.Div(a0SqxOtD, avGeszl)
		if /*line oeeS6NQ94.go:1*/ a0SqxOtD.BitLen() > 64 {
			return math.MaxUint64
		}

		if /*line auM92_.go:1*/ a0SqxOtD.Uint64() < 200 {
			return 200
		}
		return /*line MHHsYMXIlz.go:1*/ a0SqxOtD.Uint64()
	}
	a0SqxOtD = /*line Ocy4_IRtiOA1.go:1*/ wsQl4Xqf(a0SqxOtD)
	/*line K1BaGuL.go:1*/ a0SqxOtD.Mul(a0SqxOtD /*line dId34Y_T.go:1*/, math.BigMax(d25OXcYBofJ5, y_M1FClUrWEw))
	/*line A05b9uEowK.go:1*/ a0SqxOtD.Div(a0SqxOtD, rB1__KHUhrF)

	if /*line HlZo9m.go:1*/ a0SqxOtD.BitLen() > 64 {
		return math.MaxUint64
	}
	return /*line WY3dhZxoA.go:1*/ a0SqxOtD.Uint64()
}

func (oPtEQcrX *tVYGZaqW3a) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	var (
		nw6_fH = /*line qyrnp4iczS8.go:1*/ new(big.Int).SetBytes( /*line NKHBHE_.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 0, 32)).Uint64()
		cuKNXu = /*line CRm4GadHjX.go:1*/ new(big.Int).SetBytes( /*line fDsSg11POlmK.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 32, 32)).Uint64()
		bSdZdr = /*line ExwL1Ll.go:1*/ new(big.Int).SetBytes( /*line Pk1qYez.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 64, 32)).Uint64()
	)
	if /*line Gg4tj9Xf.go:1*/ len(sIVM3RLPSKSh) > 96 {
		sIVM3RLPSKSh = sIVM3RLPSKSh[96:]
	} else {
		sIVM3RLPSKSh = sIVM3RLPSKSh[:0]
	}

	if nw6_fH == 0 && bSdZdr == 0 {
		return []byte{}, nil
	}

	var (
		aY3v02xiEWPQ = /*line va2Rm_DsQ5F5.go:1*/ new(big.Int).SetBytes( /*line mHJe3UPq.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 0, nw6_fH))
		d9VP52t8g5n  = /*line I03PF1oiQa.go:1*/ new(big.Int).SetBytes( /*line tXTpA4.go:1*/ m8O5lZzG(sIVM3RLPSKSh, nw6_fH, cuKNXu))
		fPqbysU      = /*line AYuBWeNJ2z.go:1*/ new(big.Int).SetBytes( /*line lpajuwNaM.go:1*/ m8O5lZzG(sIVM3RLPSKSh, nw6_fH+cuKNXu, bSdZdr))
		gbvC5akim    []byte
	)
	switch {
	case /*line zWfbviAnYx.go:1*/ fPqbysU.BitLen() == 0:

		return /*line hfNzW7aOrHrA.go:1*/ common.LeftPadBytes([]byte{} /*line kgWfJUMop.go:1*/, int(bSdZdr)), nil
	case /*line ufaAaaM.go:1*/ aY3v02xiEWPQ.BitLen() == 1:

		gbvC5akim = /*line zKGIDBxv.go:1*/ aY3v02xiEWPQ.Mod(aY3v02xiEWPQ, fPqbysU).Bytes()
	default:
		gbvC5akim = /*line u1I4yxEEM9c.go:1*/ aY3v02xiEWPQ.Exp(aY3v02xiEWPQ, d9VP52t8g5n, fPqbysU).Bytes()
	}
	return /*line xhFJTVAl2k8.go:1*/ common.LeftPadBytes(gbvC5akim /*line ljVmSHJdaf9.go:1*/, int(bSdZdr)), nil
}

func pHNesSW(f45cPAJ []byte) (*bn256.G1, error) {
	gTaLKvaWJ := /*line RPco_TIy.go:1*/ new(bn256.G1)
	if _, v1p7xkTykQgN := /*line IFyvknYd9FTb.go:1*/ gTaLKvaWJ.Unmarshal(f45cPAJ); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	return gTaLKvaWJ, nil
}

func s_v2h5ibfNY(f45cPAJ []byte) (*bn256.G2, error) {
	gTaLKvaWJ := /*line y0g_7b1e.go:1*/ new(bn256.G2)
	if _, v1p7xkTykQgN := /*line JZdMTcY.go:1*/ gTaLKvaWJ.Unmarshal(f45cPAJ); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	return gTaLKvaWJ, nil
}

func wDzVC4x(sIVM3RLPSKSh []byte) ([]byte, error) {
	hpgMKBm, v1p7xkTykQgN := /*line It45e6r.go:1*/ pHNesSW( /*line ZkVGHil.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 0, 64))
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	eK4i7p94C79, v1p7xkTykQgN := /*line B8HVbtSCRNT.go:1*/ pHNesSW( /*line V2QYaPna9BY7.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 64, 64))
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	ekzSsj9 := /*line UjW8L81.go:1*/ new(bn256.G1)
	/*line GQFmD7.go:1*/ ekzSsj9.Add(hpgMKBm, eK4i7p94C79)
	return /*line cLoWakjDM2.go:1*/ ekzSsj9.Marshal(), nil
}

type lEsxPQVLShM struct{}

func (oPtEQcrX *lEsxPQVLShM) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bn256AddGasIstanbul
}

func (oPtEQcrX *lEsxPQVLShM) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	return /*line DemyS4hXCu.go:1*/ wDzVC4x(sIVM3RLPSKSh)
}

type u0cWnoUxkq1D struct{}

func (oPtEQcrX *u0cWnoUxkq1D) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bn256AddGasByzantium
}

func (oPtEQcrX *u0cWnoUxkq1D) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	return /*line AamVaZ.go:1*/ wDzVC4x(sIVM3RLPSKSh)
}

func i0DFph3i93U6(sIVM3RLPSKSh []byte) ([]byte, error) {
	gTaLKvaWJ, v1p7xkTykQgN := /*line I4Zrd7c.go:1*/ pHNesSW( /*line px0Lkjq3s8mr.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 0, 64))
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	ekzSsj9 := /*line IiDnPEScfSrx.go:1*/ new(bn256.G1)
	/*line pxqWKs6true.go:1*/ ekzSsj9.ScalarMult(gTaLKvaWJ /*line JKwB0FUP.go:1*/, new(big.Int).SetBytes( /*line fMPt2I4Ha.go:1*/ m8O5lZzG(sIVM3RLPSKSh, 64, 32)))
	return /*line Kp8ozYIfT6.go:1*/ ekzSsj9.Marshal(), nil
}

type uAwS_e1Y struct{}

func (oPtEQcrX *uAwS_e1Y) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bn256ScalarMulGasIstanbul
}

func (oPtEQcrX *uAwS_e1Y) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	return /*line k9gsICw22.go:1*/ i0DFph3i93U6(sIVM3RLPSKSh)
}

type rPzaYtsJ struct{}

func (oPtEQcrX *rPzaYtsJ) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bn256ScalarMulGasByzantium
}

func (oPtEQcrX *rPzaYtsJ) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	return /*line Bau9RsKJmQ.go:1*/ i0DFph3i93U6(sIVM3RLPSKSh)
}

var (
	tB0x6FoeGT = func() [] /*line GuPSKXE2.go:1*/ byte {
		data := [] /*line GuPSKXE2.go:1*/ byte("\x1a~\x00-\x11:\xc2\x00$:\x1e/\x00:\x00\x00\x00\x00\x00\x1f\x00\x18U\xfa\x00W&jVl\x041")
		positions := [...]byte{22, 5, 3, 29, 23, 25, 27, 25, 13, 6, 13, 27, 27, 1, 3, 25, 13, 26, 23, 8, 9, 19, 28, 27, 1, 11, 0, 25, 5, 19, 27, 10, 26, 28, 31, 4, 21, 4, 11, 21, 10, 23, 1, 30}
		for i := 0; i < 44; i += 2 {
			localKey := /*line GuPSKXE2.go:1*/ byte(i) + /*line GuPSKXE2.go:1*/ byte(positions[i]^positions[i+1]) + 179
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return data
	}()

	cyt1LXav = /*line xPaRca.go:1*/ make([]byte, 32)

	fQLehcag = /*line zBudHqU.go:1*/ errors.New(func() /*line IMIVIXfrw.go:1*/ string {
		fullData := [] /*line IMIVIXfrw.go:1*/ byte("'\xcf\xcb-ZOMD\xbc\n\x89ؓ\xdfi\x03<\x9d\v\xd1\xc8\x1b\x1c\x14\r\xf0\v\xf8\xfdO\x05Q\x9b\x1dt\xe3z\x90\xaa\x12y\xac\x94\xd1!\xb4\xa4\xe0\xdd\x1a\x8bN\x1c\x15d\x98(\x00u\xed+<")
		idxKey := [] /*line IMIVIXfrw.go:1*/ byte("ƑF\xdbԉ,\x04\xc02\xa2x2\xfd\xf4\x18")
		data := /*line IMIVIXfrw.go:1*/ make([]byte, 0, 32)
		data = /*line IMIVIXfrw.go:1*/ append(data, fullData[199^ /*line IMIVIXfrw.go:1*/ int(idxKey[0])]+fullData[202^ /*line IMIVIXfrw.go:1*/ int(idxKey[0])], fullData[6^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])]^fullData[34^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])], fullData[37^ /*line IMIVIXfrw.go:1*/ int(idxKey[15])]-fullData[19^ /*line IMIVIXfrw.go:1*/ int(idxKey[15])], fullData[33^ /*line IMIVIXfrw.go:1*/ int(idxKey[9])]+fullData[47^ /*line IMIVIXfrw.go:1*/ int(idxKey[9])], fullData[41^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])]^fullData[47^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])], fullData[82^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])]-fullData[64^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])], fullData[2^ /*line IMIVIXfrw.go:1*/ int(idxKey[6])]^fullData[56^ /*line IMIVIXfrw.go:1*/ int(idxKey[6])], fullData[44^ /*line IMIVIXfrw.go:1*/ int(idxKey[6])]^fullData[31^ /*line IMIVIXfrw.go:1*/ int(idxKey[6])], fullData[30^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])]-fullData[36^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])], fullData[205^ /*line IMIVIXfrw.go:1*/ int(idxKey[14])]+fullData[214^ /*line IMIVIXfrw.go:1*/ int(idxKey[14])], fullData[217^ /*line IMIVIXfrw.go:1*/ int(idxKey[8])]+fullData[232^ /*line IMIVIXfrw.go:1*/ int(idxKey[8])], fullData[135^ /*line IMIVIXfrw.go:1*/ int(idxKey[5])]^fullData[128^ /*line IMIVIXfrw.go:1*/ int(idxKey[5])], fullData[23^ /*line IMIVIXfrw.go:1*/ int(idxKey[15])]+fullData[57^ /*line IMIVIXfrw.go:1*/ int(idxKey[15])], fullData[66^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])]-fullData[95^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])], fullData[236^ /*line IMIVIXfrw.go:1*/ int(idxKey[3])]^fullData[224^ /*line IMIVIXfrw.go:1*/ int(idxKey[3])], fullData[232^ /*line IMIVIXfrw.go:1*/ int(idxKey[14])]-fullData[198^ /*line IMIVIXfrw.go:1*/ int(idxKey[14])], fullData[110^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])]+fullData[124^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])], fullData[99^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])]^fullData[105^ /*line IMIVIXfrw.go:1*/ int(idxKey[11])], fullData[84^ /*line IMIVIXfrw.go:1*/ int(idxKey[2])]+fullData[115^ /*line IMIVIXfrw.go:1*/ int(idxKey[2])], fullData[6^ /*line IMIVIXfrw.go:1*/ int(idxKey[12])]-fullData[27^ /*line IMIVIXfrw.go:1*/ int(idxKey[12])], fullData[52^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])]^fullData[12^ /*line IMIVIXfrw.go:1*/ int(idxKey[7])], fullData[204^ /*line IMIVIXfrw.go:1*/ int(idxKey[4])]^fullData[226^ /*line IMIVIXfrw.go:1*/ int(idxKey[4])], fullData[150^ /*line IMIVIXfrw.go:1*/ int(idxKey[5])]-fullData[132^ /*line IMIVIXfrw.go:1*/ int(idxKey[5])], fullData[55^ /*line IMIVIXfrw.go:1*/ int(idxKey[9])]+fullData[3^ /*line IMIVIXfrw.go:1*/ int(idxKey[9])], fullData[158^ /*line IMIVIXfrw.go:1*/ int(idxKey[5])]^fullData[173^ /*line IMIVIXfrw.go:1*/ int(idxKey[5])], fullData[86^ /*line IMIVIXfrw.go:1*/ int(idxKey[2])]+fullData[122^ /*line IMIVIXfrw.go:1*/ int(idxKey[2])], fullData[227^ /*line IMIVIXfrw.go:1*/ int(idxKey[13])]+fullData[232^ /*line IMIVIXfrw.go:1*/ int(idxKey[13])], fullData[227^ /*line IMIVIXfrw.go:1*/ int(idxKey[0])]+fullData[229^ /*line IMIVIXfrw.go:1*/ int(idxKey[0])], fullData[155^ /*line IMIVIXfrw.go:1*/ int(idxKey[1])]+fullData[190^ /*line IMIVIXfrw.go:1*/ int(idxKey[1])], fullData[42^ /*line IMIVIXfrw.go:1*/ int(idxKey[6])]+fullData[47^ /*line IMIVIXfrw.go:1*/ int(idxKey[6])], fullData[52^ /*line IMIVIXfrw.go:1*/ int(idxKey[15])]+fullData[31^ /*line IMIVIXfrw.go:1*/ int(idxKey[15])])
		return /*line IMIVIXfrw.go:1*/ string(data)
	}())
)

func dyDgxjCa(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line x3lYylGx.go:1*/ len(sIVM3RLPSKSh)%192 > 0 {
		return nil, fQLehcag
	}

	var (
		eeIy0sWZqmg []*bn256.G1
		ahx9Mrld75  []*bn256.G2
	)
	for hv0ZaFQN := 0; hv0ZaFQN < /*line n3adrRsdqpfF.go:1*/ len(sIVM3RLPSKSh); hv0ZaFQN += 192 {
		oPtEQcrX, v1p7xkTykQgN := /*line N_AOTY5f.go:1*/ pHNesSW(sIVM3RLPSKSh[hv0ZaFQN : hv0ZaFQN+64])
		if v1p7xkTykQgN != nil {
			return nil, v1p7xkTykQgN
		}
		k_aCKc, v1p7xkTykQgN := /*line EWTcBm.go:1*/ s_v2h5ibfNY(sIVM3RLPSKSh[hv0ZaFQN+64 : hv0ZaFQN+192])
		if v1p7xkTykQgN != nil {
			return nil, v1p7xkTykQgN
		}
		eeIy0sWZqmg = /*line dxkWY8i.go:1*/ append(eeIy0sWZqmg, oPtEQcrX)
		ahx9Mrld75 = /*line wg4NYNL.go:1*/ append(ahx9Mrld75, k_aCKc)
	}

	if /*line wrG9K5NpLIK.go:1*/ bn256.PairingCheck(eeIy0sWZqmg, ahx9Mrld75) {
		return tB0x6FoeGT, nil
	}
	return cyt1LXav, nil
}

type lzLDwHcHQDpl struct{}

func (oPtEQcrX *lzLDwHcHQDpl) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bn256PairingBaseGasIstanbul + /*line c7VACa.go:1*/ uint64( /*line Pemfir21Y6dz.go:1*/ len(sIVM3RLPSKSh)/192)*params.Bn256PairingPerPointGasIstanbul
}

func (oPtEQcrX *lzLDwHcHQDpl) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	return /*line Uncfh0.go:1*/ dyDgxjCa(sIVM3RLPSKSh)
}

type e8otDl struct{}

func (oPtEQcrX *e8otDl) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bn256PairingBaseGasByzantium + /*line lb_3fZehIZ.go:1*/ uint64( /*line cZoTiPUook57.go:1*/ len(sIVM3RLPSKSh)/192)*params.Bn256PairingPerPointGasByzantium
}

func (oPtEQcrX *e8otDl) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	return /*line BvJ1tzWWa.go:1*/ dyDgxjCa(sIVM3RLPSKSh)
}

type zAsr27 struct{}

func (oPtEQcrX *zAsr27) RequiredGas(sIVM3RLPSKSh []byte) uint64 {

	if /*line ilSRwHOTnB.go:1*/ len(sIVM3RLPSKSh) != blake2FInputLength {
		return 0
	}
	return /*line k9w92a.go:1*/ uint64( /*line zi9BOEa.go:1*/ binary.BigEndian.Uint32(sIVM3RLPSKSh[0:4]))
}

const (
	blake2FInputLength        = 213
	blake2FFinalBlockBytes    = /*line RF87A9yd9n.go:1*/ byte(1)
	blake2FNonFinalBlockBytes = /*line WyIuodpr1.go:1*/ byte(0)
)

var (
	kXYV8wOxLoeU = /*line aaF_0Kp.go:1*/ errors.New(func() /*line HQXI580z.go:1*/ string {
		data := /*line HQXI580z.go:1*/ make([]byte, 0, 21)
		i := 11
		decryptKey := 35
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data = /*line HQXI580z.go:1*/ append(data, "\xea\xe4\xf2"...,
				)
				i = 8
			case 9:
				data = /*line HQXI580z.go:1*/ append(data, 200)
				i = 7
			case 1:
				data = /*line HQXI580z.go:1*/ append(data, 220)
				i = 6
			case 10:
				i = 4
				data = /*line HQXI580z.go:1*/ append(data, 210)
			case 3:
				i = 10
				data = /*line HQXI580z.go:1*/ append(data, "\xe4\xd0\xd4"...,
				)
			case 2:
				i = 1
				data = /*line HQXI580z.go:1*/ append(data, "\xd3\xd6"...,
				)
			case 8:
				i = 5
				data = /*line HQXI580z.go:1*/ append(data, 231)
			case 4:
				i = 2
				data = /*line HQXI580z.go:1*/ append(data, "\u038b\xcd"...,
				)
			case 6:
				data = /*line HQXI580z.go:1*/ append(data, "ԁ\xce"...,
				)
				i = 9
			case 11:
				i = 3
				data = /*line HQXI580z.go:1*/ append(data, "\xd5\xdb"...,
				)
			case 5:
				for y := range data {
					data[y] = data[y] - /*line HQXI580z.go:1*/ byte(decryptKey^y)
				}
				i = 0
			}
		}
		return /*line HQXI580z.go:1*/ string(data)
	}())
	d21J4LSI3 = /*line FXjSNkdt9xD.go:1*/ errors.New(func() /*line KAjqT4R.go:1*/ string {
		data := [] /*line KAjqT4R.go:1*/ byte("\x8d\x9f|ali\x95\x947z\xdaql\x84\x8b\x9fgg")
		positions := [...]byte{2, 6, 9, 15, 8, 0, 15, 1, 11, 14, 9, 13, 16, 11, 10, 7, 13, 0, 14, 6}
		for i := 0; i < 20; i += 2 {
			localKey := /*line KAjqT4R.go:1*/ byte(i) + /*line KAjqT4R.go:1*/ byte(positions[i]^positions[i+1]) + 223
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line KAjqT4R.go:1*/ string(data)
	}())
)

func (oPtEQcrX *zAsr27) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line WiAsbagO.go:1*/ len(sIVM3RLPSKSh) != blake2FInputLength {
		return nil, kXYV8wOxLoeU
	}
	if sIVM3RLPSKSh[212] != blake2FNonFinalBlockBytes && sIVM3RLPSKSh[212] != blake2FFinalBlockBytes {
		return nil, d21J4LSI3
	}

	var (
		xQho8iK   = /*line WmX13NDa.go:1*/ binary.BigEndian.Uint32(sIVM3RLPSKSh[0:4])
		pbPswdCmR = sIVM3RLPSKSh[212] == blake2FFinalBlockBytes

		eNKBLUKPi [8]uint64
		bHALbG    [16]uint64
		k_aCKc    [2]uint64
	)
	for hv0ZaFQN := 0; hv0ZaFQN < 8; hv0ZaFQN++ {
		gmjJkRQEj := 4 + hv0ZaFQN*8
		eNKBLUKPi[hv0ZaFQN] = /*line bRabfFq5RM.go:1*/ binary.LittleEndian.Uint64(sIVM3RLPSKSh[gmjJkRQEj : gmjJkRQEj+8])
	}
	for hv0ZaFQN := 0; hv0ZaFQN < 16; hv0ZaFQN++ {
		gmjJkRQEj := 68 + hv0ZaFQN*8
		bHALbG[hv0ZaFQN] = /*line GjjgR4o.go:1*/ binary.LittleEndian.Uint64(sIVM3RLPSKSh[gmjJkRQEj : gmjJkRQEj+8])
	}
	k_aCKc[0] = /*line iuvcne.go:1*/ binary.LittleEndian.Uint64(sIVM3RLPSKSh[196:204])
	k_aCKc[1] = /*line PeV1YOG1amaC.go:1*/ binary.LittleEndian.Uint64(sIVM3RLPSKSh[204:212])

	/*line CEVkHiP0aa2Q.go:1*/
	blake2b.F(&eNKBLUKPi, bHALbG, k_aCKc, pbPswdCmR, xQho8iK)

	dhOXtmt2Uz9 := /*line mz9kbo5xW02v.go:1*/ make([]byte, 64)
	for hv0ZaFQN := 0; hv0ZaFQN < 8; hv0ZaFQN++ {
		gmjJkRQEj := hv0ZaFQN * 8
		/*line CBR9pZ6GHR.go:1*/ binary.LittleEndian.PutUint64(dhOXtmt2Uz9[gmjJkRQEj:gmjJkRQEj+8], eNKBLUKPi[hv0ZaFQN])
	}
	return dhOXtmt2Uz9, nil
}

var (
	l2R6GU = /*line qKhNvCfTKj.go:1*/ errors.New(func() /*line fkF3T89t.go:1*/ string {
		data := /*line fkF3T89t.go:1*/ make([]byte, 0, 21)
		i := 7
		decryptKey := 157
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				for y := range data {
					data[y] = data[y] - /*line fkF3T89t.go:1*/ byte(decryptKey^y)
				}
				i = 6
			case 3:
				data = /*line fkF3T89t.go:1*/ append(data, "\x1f\xda\x1e\""...,
				)
				i = 0
			case 9:
				i = 1
				data = /*line fkF3T89t.go:1*/ append(data, 43)
			case 1:
				data = /*line fkF3T89t.go:1*/ append(data, "%\xd0\x1f"...,
				)
				i = 2
			case 2:
				i = 5
				data = /*line fkF3T89t.go:1*/ append(data, "\x17\x1b"...,
				)
			case 0:
				data = /*line fkF3T89t.go:1*/ append(data, 39)
				i = 9
			case 7:
				i = 8
				data = /*line fkF3T89t.go:1*/ append(data, "&*5\x1f"...,
				)
			case 8:
				i = 3
				data = /*line fkF3T89t.go:1*/ append(data, "%!"...,
				)
			case 5:
				data = /*line fkF3T89t.go:1*/ append(data, "\x13#\x16"...,
				)
				i = 4
			}
		}
		return /*line fkF3T89t.go:1*/ string(data)
	}())
	h5R02Uy = /*line anOaYRXZ.go:1*/ errors.New(func() /*line I7U0LENi.go:1*/ string {
		data := /*line I7U0LENi.go:1*/ make([]byte, 0, 32)
		i := 3
		decryptKey := 149
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 2
				for y := range data {
					data[y] = data[y] + /*line I7U0LENi.go:1*/ byte(decryptKey^y)
				}
			case 7:
				i = 11
				data = /*line I7U0LENi.go:1*/ append(data, "\x02\f\x0f"...,
				)
			case 13:
				i = 6
				data = /*line I7U0LENi.go:1*/ append(data, 4)
			case 11:
				data = /*line I7U0LENi.go:1*/ append(data, "\xbc\r"...,
				)
				i = 10
			case 12:
				i = 0
				data = /*line I7U0LENi.go:1*/ append(data, "\x0f\a\xf9\x04"...,
				)
			case 8:
				i = 13
				data = /*line I7U0LENi.go:1*/ append(data, "\xa4\xe6\xee"...,
				)
			case 1:
				i = 4
				data = /*line I7U0LENi.go:1*/ append(data, 237)
			case 4:
				i = 5
				data = /*line I7U0LENi.go:1*/ append(data, "\xaa\xed\xf1\xea"...,
				)
			case 9:
				i = 1
				data = /*line I7U0LENi.go:1*/ append(data, "\xf7\xf5"...,
				)
			case 10:
				i = 12
				data = /*line I7U0LENi.go:1*/ append(data, "\t\a\xb8\xf7"...,
				)
			case 3:
				data = /*line I7U0LENi.go:1*/ append(data, "\xf8\xfe\x03\xef"...,
				)
				i = 9
			case 5:
				data = /*line I7U0LENi.go:1*/ append(data, "\xf2\xe7"...,
				)
				i = 8
			case 6:
				i = 7
				data = /*line I7U0LENi.go:1*/ append(data, 13)
			}
		}
		return /*line I7U0LENi.go:1*/ string(data)
	}())
	tuYLU7Vb9 = /*line H1PjPn7ssNfY.go:1*/ errors.New(func() /*line aDbhimmECi7.go:1*/ string {
		data := [] /*line aDbhimmECi7.go:1*/ byte("\x971 p\xae\nn\x1e i\x15\xeb\xf6o\xf0\xa6\xe1\xa7wco\xb4\r_cU,\xa8\xc7\xec\x92r\xc5\xdb\xd7")
		positions := [...]byte{32, 28, 21, 27, 28, 5, 22, 7, 14, 22, 10, 4, 25, 0, 30, 25, 26, 23, 5, 11, 29, 16, 23, 12, 25, 14, 34, 10, 32, 18, 32, 21, 11, 15, 17, 18, 0, 27, 12, 33}
		for i := 0; i < 40; i += 2 {
			localKey := /*line aDbhimmECi7.go:1*/ byte(i) + /*line aDbhimmECi7.go:1*/ byte(positions[i]^positions[i+1]) + 98
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line aDbhimmECi7.go:1*/ string(data)
	}())
	drqKDxu6OB = /*line qVXkHum1.go:1*/ errors.New(func() /*line W0GtzmEE.go:1*/ string {
		key := [] /*line W0GtzmEE.go:1*/ byte("\f\xea\x96)s\xaf\v]\xf2\x12Μ\xaf\xaf\xa0ȇ\xce;\x89\xa9\xcd\x06\xd1`P£u\x82\xd4o\x88@\x97")
		data := [] /*line W0GtzmEE.go:1*/ byte("[H\x8aG\xfc\xbac\x17.W\xa5\x84\xbf\xc0\xd4X\xe8\xa0\xe5\xdaƥl\x94\x03$^\xd0\x00\xe0\x93\x03\xe75\xd9")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line W0GtzmEE.go:1*/ string(data)
	}())
)

type rvFNzWYNNj struct{}

func (oPtEQcrX *rvFNzWYNNj) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381G1AddGas
}

func (oPtEQcrX *rvFNzWYNNj) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line AQwa9hwn.go:1*/ len(sIVM3RLPSKSh) != 256 {
		return nil, l2R6GU
	}
	var v1p7xkTykQgN error
	var fgX6hYEn9, vzqai0UFHHxu *bls12381.PointG1

	cJWe1uw3Q := /*line UwUU0VIyVg.go:1*/ bls12381.NewG1()

	if fgX6hYEn9, v1p7xkTykQgN = /*line CFKNNGyAwSA.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[:128]); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	if vzqai0UFHHxu, v1p7xkTykQgN = /*line st5EhZ.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[128:]); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	gGFrnFgT := /*line TAFNhIkD.go:1*/ cJWe1uw3Q.New()
	/*line cIfP_nSshc2.go:1*/ cJWe1uw3Q.Add(gGFrnFgT, fgX6hYEn9, vzqai0UFHHxu)

	return /*line JoxNw1sxk5Y7.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type hF6SEzF struct{}

func (oPtEQcrX *hF6SEzF) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381G1MulGas
}

func (oPtEQcrX *hF6SEzF) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line jLWe_DlN5ng.go:1*/ len(sIVM3RLPSKSh) != 160 {
		return nil, l2R6GU
	}
	var v1p7xkTykQgN error
	var fgX6hYEn9 *bls12381.PointG1

	cJWe1uw3Q := /*line qg65_b.go:1*/ bls12381.NewG1()

	if fgX6hYEn9, v1p7xkTykQgN = /*line FSipc2.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[:128]); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	mfy2T9 := /*line DWhJnsDEyY.go:1*/ new(big.Int).SetBytes(sIVM3RLPSKSh[128:])

	gGFrnFgT := /*line z0EQkmMuKeGp.go:1*/ cJWe1uw3Q.New()
	/*line jd85iB2El.go:1*/ cJWe1uw3Q.MulScalar(gGFrnFgT, fgX6hYEn9, mfy2T9)

	return /*line xTtqCbeb_.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type vfwOEaC struct{}

func (oPtEQcrX *vfwOEaC) RequiredGas(sIVM3RLPSKSh []byte) uint64 {

	zsVIRN6k := /*line vPQAYNrLV.go:1*/ len(sIVM3RLPSKSh) / 160
	if zsVIRN6k == 0 {

		return 0
	}

	var fgDU47SPo uint64
	if dAZ75o8e := /*line a48OgzL.go:1*/ len(params.Bls12381MultiExpDiscountTable); zsVIRN6k < dAZ75o8e {
		fgDU47SPo = params.Bls12381MultiExpDiscountTable[zsVIRN6k-1]
	} else {
		fgDU47SPo = params.Bls12381MultiExpDiscountTable[dAZ75o8e-1]
	}

	return ( /*line jGxF7GEUPJ.go:1*/ uint64(zsVIRN6k) * params.Bls12381G1MulGas * fgDU47SPo) / 1000
}

func (oPtEQcrX *vfwOEaC) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	zsVIRN6k := /*line m03mV_PFdO.go:1*/ len(sIVM3RLPSKSh) / 160
	if /*line F9x2iBI.go:1*/ len(sIVM3RLPSKSh) == 0 || /*line lR8Z4wm.go:1*/ len(sIVM3RLPSKSh)%160 != 0 {
		return nil, l2R6GU
	}
	var v1p7xkTykQgN error
	e1evqtmO := /*line LKsU9T.go:1*/ make([]*bls12381.PointG1, zsVIRN6k)
	eeMvYXikG := /*line mORs3Os.go:1*/ make([]*big.Int, zsVIRN6k)

	cJWe1uw3Q := /*line xe6pG1zR.go:1*/ bls12381.NewG1()

	for hv0ZaFQN := 0; hv0ZaFQN < zsVIRN6k; hv0ZaFQN++ {
		waclpXm8dPXx := 160 * hv0ZaFQN
		vQFJJk7haJxo, kWKhkYo8d, tIyAZNGq := waclpXm8dPXx, waclpXm8dPXx+128, waclpXm8dPXx+160

		if e1evqtmO[hv0ZaFQN], v1p7xkTykQgN = /*line qk8Bxf.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[vQFJJk7haJxo:kWKhkYo8d]); v1p7xkTykQgN != nil {
			return nil, v1p7xkTykQgN
		}

		eeMvYXikG[hv0ZaFQN] = /*line CapStB2avS9_.go:1*/ new(big.Int).SetBytes(sIVM3RLPSKSh[kWKhkYo8d:tIyAZNGq])
	}

	gGFrnFgT := /*line qQY3W2rRjx.go:1*/ cJWe1uw3Q.New()
	/*line kog3FalL.go:1*/ cJWe1uw3Q.MultiExp(gGFrnFgT, e1evqtmO, eeMvYXikG)

	return /*line D7ValoY68h.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type yUf2fKtK struct{}

func (oPtEQcrX *yUf2fKtK) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381G2AddGas
}

func (oPtEQcrX *yUf2fKtK) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line Z8RgAmR.go:1*/ len(sIVM3RLPSKSh) != 512 {
		return nil, l2R6GU
	}
	var v1p7xkTykQgN error
	var fgX6hYEn9, vzqai0UFHHxu *bls12381.PointG2

	cJWe1uw3Q := /*line baou11pv.go:1*/ bls12381.NewG2()
	gGFrnFgT := /*line zj2tplvO6gu.go:1*/ cJWe1uw3Q.New()

	if fgX6hYEn9, v1p7xkTykQgN = /*line ERXv5sP.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[:256]); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	if vzqai0UFHHxu, v1p7xkTykQgN = /*line RCC1ttZUPXzB.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[256:]); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	/*line qQI4Wx.go:1*/
	cJWe1uw3Q.Add(gGFrnFgT, fgX6hYEn9, vzqai0UFHHxu)

	return /*line TLSqfc3C8.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type vJqscRtml_ struct{}

func (oPtEQcrX *vJqscRtml_) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381G2MulGas
}

func (oPtEQcrX *vJqscRtml_) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line vuZ4DWw.go:1*/ len(sIVM3RLPSKSh) != 288 {
		return nil, l2R6GU
	}
	var v1p7xkTykQgN error
	var fgX6hYEn9 *bls12381.PointG2

	cJWe1uw3Q := /*line IjYm61UzKFYa.go:1*/ bls12381.NewG2()

	if fgX6hYEn9, v1p7xkTykQgN = /*line a3XFM6jaA.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[:256]); v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	mfy2T9 := /*line GlYFAfN.go:1*/ new(big.Int).SetBytes(sIVM3RLPSKSh[256:])

	gGFrnFgT := /*line R2uOyfSS.go:1*/ cJWe1uw3Q.New()
	/*line USbehYd.go:1*/ cJWe1uw3Q.MulScalar(gGFrnFgT, fgX6hYEn9, mfy2T9)

	return /*line IuCob9r.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type k8odgmq8WI struct{}

func (oPtEQcrX *k8odgmq8WI) RequiredGas(sIVM3RLPSKSh []byte) uint64 {

	zsVIRN6k := /*line uaqsCrCh.go:1*/ len(sIVM3RLPSKSh) / 288
	if zsVIRN6k == 0 {

		return 0
	}

	var fgDU47SPo uint64
	if dAZ75o8e := /*line acYC1V3.go:1*/ len(params.Bls12381MultiExpDiscountTable); zsVIRN6k < dAZ75o8e {
		fgDU47SPo = params.Bls12381MultiExpDiscountTable[zsVIRN6k-1]
	} else {
		fgDU47SPo = params.Bls12381MultiExpDiscountTable[dAZ75o8e-1]
	}

	return ( /*line GayicYzPD.go:1*/ uint64(zsVIRN6k) * params.Bls12381G2MulGas * fgDU47SPo) / 1000
}

func (oPtEQcrX *k8odgmq8WI) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	zsVIRN6k := /*line HrmZmhEYuLkq.go:1*/ len(sIVM3RLPSKSh) / 288
	if /*line a77y478B5k6.go:1*/ len(sIVM3RLPSKSh) == 0 || /*line Ho9IV9sHDT4M.go:1*/ len(sIVM3RLPSKSh)%288 != 0 {
		return nil, l2R6GU
	}
	var v1p7xkTykQgN error
	e1evqtmO := /*line zlkJm8ePXr.go:1*/ make([]*bls12381.PointG2, zsVIRN6k)
	eeMvYXikG := /*line G0GQ3c2.go:1*/ make([]*big.Int, zsVIRN6k)

	cJWe1uw3Q := /*line nDHfTB.go:1*/ bls12381.NewG2()

	for hv0ZaFQN := 0; hv0ZaFQN < zsVIRN6k; hv0ZaFQN++ {
		waclpXm8dPXx := 288 * hv0ZaFQN
		vQFJJk7haJxo, kWKhkYo8d, tIyAZNGq := waclpXm8dPXx, waclpXm8dPXx+256, waclpXm8dPXx+288

		if e1evqtmO[hv0ZaFQN], v1p7xkTykQgN = /*line OPyEsV5eJtFd.go:1*/ cJWe1uw3Q.DecodePoint(sIVM3RLPSKSh[vQFJJk7haJxo:kWKhkYo8d]); v1p7xkTykQgN != nil {
			return nil, v1p7xkTykQgN
		}

		eeMvYXikG[hv0ZaFQN] = /*line VAO53IoMx3jy.go:1*/ new(big.Int).SetBytes(sIVM3RLPSKSh[kWKhkYo8d:tIyAZNGq])
	}

	gGFrnFgT := /*line HrG7XVp.go:1*/ cJWe1uw3Q.New()
	/*line NHqfvI.go:1*/ cJWe1uw3Q.MultiExp(gGFrnFgT, e1evqtmO, eeMvYXikG)

	return /*line HFWdsPpD.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type fQxEr40b struct{}

func (oPtEQcrX *fQxEr40b) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381PairingBaseGas + /*line EzviIwKKgg8.go:1*/ uint64( /*line IT1kBTgnp4KV.go:1*/ len(sIVM3RLPSKSh)/384)*params.Bls12381PairingPerPairGas
}

func (oPtEQcrX *fQxEr40b) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	zsVIRN6k := /*line ZsdP7E.go:1*/ len(sIVM3RLPSKSh) / 384
	if /*line BKJ4ITRE.go:1*/ len(sIVM3RLPSKSh) == 0 || /*line BfqfZd.go:1*/ len(sIVM3RLPSKSh)%384 != 0 {
		return nil, l2R6GU
	}

	mfy2T9 := /*line hCEsa621Ubw.go:1*/ bls12381.NewPairingEngine()
	uAe7gysrq4E, bioOK0nofzsu := mfy2T9.G1, mfy2T9.G2

	for hv0ZaFQN := 0; hv0ZaFQN < zsVIRN6k; hv0ZaFQN++ {
		waclpXm8dPXx := 384 * hv0ZaFQN
		vQFJJk7haJxo, kWKhkYo8d, tIyAZNGq := waclpXm8dPXx, waclpXm8dPXx+128, waclpXm8dPXx+384

		vzqai0UFHHxu, v1p7xkTykQgN := /*line ppY52e.go:1*/ uAe7gysrq4E.DecodePoint(sIVM3RLPSKSh[vQFJJk7haJxo:kWKhkYo8d])
		if v1p7xkTykQgN != nil {
			return nil, v1p7xkTykQgN
		}

		nFi04GDa, v1p7xkTykQgN := /*line Cu4JgI2icL_.go:1*/ bioOK0nofzsu.DecodePoint(sIVM3RLPSKSh[kWKhkYo8d:tIyAZNGq])
		if v1p7xkTykQgN != nil {
			return nil, v1p7xkTykQgN
		}

		if ! /*line bgu7pjayvj.go:1*/ uAe7gysrq4E.InCorrectSubgroup(vzqai0UFHHxu) {
			return nil, tuYLU7Vb9
		}
		if ! /*line VkVUu1W.go:1*/ bioOK0nofzsu.InCorrectSubgroup(nFi04GDa) {
			return nil, drqKDxu6OB
		}

		/*line dvbiGG.go:1*/
		mfy2T9.AddPair(vzqai0UFHHxu, nFi04GDa)
	}

	yFSnEdFC := /*line BgVVrFbDNOos.go:1*/ make([]byte, 32)

	if /*line BPlM12.go:1*/ mfy2T9.Check() {
		yFSnEdFC[31] = 1
	}
	return yFSnEdFC, nil
}

func bia_Y_BhZii(kTH29c []byte) ([]byte, error) {
	if /*line dz34FBtob.go:1*/ len(kTH29c) != 64 {
		return nil /*line b9Vdx0l_MsIG.go:1*/, errors.New(func() /*line zQnrdt.go:1*/ string {
			data := [] /*line zQnrdt.go:1*/ byte("y\x91v\xd8l\x19\x95\xbb\xdc\xcdeg\xd0 \xda\xd5e\xd2\xc5\xce\xd6 \xc3\xd1\xd4g\x97h")
			positions := [...]byte{1, 6, 1, 22, 15, 23, 26, 7, 11, 6, 6, 26, 17, 12, 8, 24, 12, 9, 9, 0, 3, 15, 18, 19, 14, 9, 5, 20, 6, 20, 22, 11, 3, 3}
			for i := 0; i < 34; i += 2 {
				localKey := /*line zQnrdt.go:1*/ byte(i) + /*line zQnrdt.go:1*/ byte(positions[i]^positions[i+1]) + 148
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line zQnrdt.go:1*/ string(data)
		}())
	}

	for hv0ZaFQN := 0; hv0ZaFQN < 16; hv0ZaFQN++ {
		if kTH29c[hv0ZaFQN] != /*line b2lHTF.go:1*/ byte(0x00) {
			return nil, h5R02Uy
		}
	}
	yFSnEdFC := /*line P2Aal_.go:1*/ make([]byte, 48)
	/*line df2z4Djz.go:1*/ copy(yFSnEdFC[:], kTH29c[16:])
	return yFSnEdFC, nil
}

type hLGLU0Dn struct{}

func (oPtEQcrX *hLGLU0Dn) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381MapG1Gas
}

func (oPtEQcrX *hLGLU0Dn) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line gdwPRSF.go:1*/ len(sIVM3RLPSKSh) != 64 {
		return nil, l2R6GU
	}

	bN6vvP, v1p7xkTykQgN := /*line pAA7cuDQm.go:1*/ bia_Y_BhZii(sIVM3RLPSKSh)
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	cJWe1uw3Q := /*line aEvQ7mq.go:1*/ bls12381.NewG1()

	gGFrnFgT, v1p7xkTykQgN := /*line IF4bmJ.go:1*/ cJWe1uw3Q.MapToCurve(bN6vvP)
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	return /*line rBfxq0.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type gRXncStcih6U struct{}

func (oPtEQcrX *gRXncStcih6U) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.Bls12381MapG2Gas
}

func (oPtEQcrX *gRXncStcih6U) Run(sIVM3RLPSKSh []byte) ([]byte, error) {

	if /*line S65dj_.go:1*/ len(sIVM3RLPSKSh) != 128 {
		return nil, l2R6GU
	}

	bN6vvP := /*line syWa_0Sb.go:1*/ make([]byte, 96)
	hz0e6kliyGeb, v1p7xkTykQgN := /*line DkdhIJU.go:1*/ bia_Y_BhZii(sIVM3RLPSKSh[:64])
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	/*line ePxKlcx.go:1*/ copy(bN6vvP[48:], hz0e6kliyGeb)
	jNVXZm, v1p7xkTykQgN := /*line OPoihzFZUKz.go:1*/ bia_Y_BhZii(sIVM3RLPSKSh[64:])
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}
	/*line ruZ7yIqxQlY7.go:1*/ copy(bN6vvP[:48], jNVXZm)

	cJWe1uw3Q := /*line _zamsP.go:1*/ bls12381.NewG2()

	gGFrnFgT, v1p7xkTykQgN := /*line WSJuBhPBkukF.go:1*/ cJWe1uw3Q.MapToCurve(bN6vvP)
	if v1p7xkTykQgN != nil {
		return nil, v1p7xkTykQgN
	}

	return /*line gWZf6LO.go:1*/ cJWe1uw3Q.EncodePoint(gGFrnFgT), nil
}

type b6Inm46Q struct{}

func (l1KtNnRwliPB *b6Inm46Q) RequiredGas(sIVM3RLPSKSh []byte) uint64 {
	return params.BlobTxPointEvaluationPrecompileGas
}

const (
	blobVerifyInputLength           = 192
	blobCommitmentVersionKZG  uint8 = 0x01
	blobPrecompileReturnValue       = "000000000000000000000000000000000000000000000000000000000000100073eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001"
)

var (
	vWPlbW = /*line DLj1Awb8AbK.go:1*/ errors.New(func() /*line smsyYBnG.go:1*/ string {
		seed := /*line smsyYBnG.go:1*/ byte(220)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line smsyYBnG.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line smsyYBnG.go:1*/ fnc(69)(143)(38)(55)(121)(239)(217)(110)(37)(79)(160)(69)(137)(190)(200)(137)(27)(47)(107)(202)
		return /*line smsyYBnG.go:1*/ string(data)
	}())
	h2K2bm8 = /*line _qB7AqiX.go:1*/ errors.New(func() /*line FOiWX4j.go:1*/ string {
		fullData := [] /*line FOiWX4j.go:1*/ byte("+Bc\xbf{\x06\x1e0\xb6\x98t\xed\xe4j\x9b\x8a\xceiz\xedO\x1d\xfej/I\xf0\x19\xe4b\x03\xff\x80{\xd4)\x12<\x8b\xaacK\xec\x1c~<e[O\xdb")
		idxKey := [] /*line FOiWX4j.go:1*/ byte("B\xe0~D\xdb\x12'5\xe1\xe6J_")
		data := /*line FOiWX4j.go:1*/ make([]byte, 0, 26)
		data = /*line FOiWX4j.go:1*/ append(data, fullData[52^ /*line FOiWX4j.go:1*/ int(idxKey[6])]-fullData[7^ /*line FOiWX4j.go:1*/ int(idxKey[6])], fullData[199^ /*line FOiWX4j.go:1*/ int(idxKey[4])]-fullData[250^ /*line FOiWX4j.go:1*/ int(idxKey[4])], fullData[241^ /*line FOiWX4j.go:1*/ int(idxKey[9])]^fullData[253^ /*line FOiWX4j.go:1*/ int(idxKey[9])], fullData[231^ /*line FOiWX4j.go:1*/ int(idxKey[8])]+fullData[209^ /*line FOiWX4j.go:1*/ int(idxKey[8])], fullData[104^ /*line FOiWX4j.go:1*/ int(idxKey[3])]-fullData[81^ /*line FOiWX4j.go:1*/ int(idxKey[3])], fullData[67^ /*line FOiWX4j.go:1*/ int(idxKey[10])]^fullData[96^ /*line FOiWX4j.go:1*/ int(idxKey[10])], fullData[89^ /*line FOiWX4j.go:1*/ int(idxKey[3])]-fullData[91^ /*line FOiWX4j.go:1*/ int(idxKey[3])], fullData[108^ /*line FOiWX4j.go:1*/ int(idxKey[0])]+fullData[92^ /*line FOiWX4j.go:1*/ int(idxKey[0])], fullData[213^ /*line FOiWX4j.go:1*/ int(idxKey[4])]^fullData[205^ /*line FOiWX4j.go:1*/ int(idxKey[4])], fullData[102^ /*line FOiWX4j.go:1*/ int(idxKey[2])]^fullData[87^ /*line FOiWX4j.go:1*/ int(idxKey[2])], fullData[3^ /*line FOiWX4j.go:1*/ int(idxKey[5])]^fullData[11^ /*line FOiWX4j.go:1*/ int(idxKey[5])], fullData[57^ /*line FOiWX4j.go:1*/ int(idxKey[5])]^fullData[31^ /*line FOiWX4j.go:1*/ int(idxKey[5])], fullData[226^ /*line FOiWX4j.go:1*/ int(idxKey[1])]^fullData[229^ /*line FOiWX4j.go:1*/ int(idxKey[1])], fullData[121^ /*line FOiWX4j.go:1*/ int(idxKey[2])]+fullData[127^ /*line FOiWX4j.go:1*/ int(idxKey[2])], fullData[58^ /*line FOiWX4j.go:1*/ int(idxKey[5])]-fullData[8^ /*line FOiWX4j.go:1*/ int(idxKey[5])], fullData[99^ /*line FOiWX4j.go:1*/ int(idxKey[3])]+fullData[71^ /*line FOiWX4j.go:1*/ int(idxKey[3])], fullData[98^ /*line FOiWX4j.go:1*/ int(idxKey[3])]+fullData[72^ /*line FOiWX4j.go:1*/ int(idxKey[3])], fullData[122^ /*line FOiWX4j.go:1*/ int(idxKey[11])]-fullData[79^ /*line FOiWX4j.go:1*/ int(idxKey[11])], fullData[79^ /*line FOiWX4j.go:1*/ int(idxKey[2])]+fullData[113^ /*line FOiWX4j.go:1*/ int(idxKey[2])], fullData[80^ /*line FOiWX4j.go:1*/ int(idxKey[3])]^fullData[68^ /*line FOiWX4j.go:1*/ int(idxKey[3])], fullData[109^ /*line FOiWX4j.go:1*/ int(idxKey[0])]^fullData[70^ /*line FOiWX4j.go:1*/ int(idxKey[0])], fullData[17^ /*line FOiWX4j.go:1*/ int(idxKey[7])]^fullData[39^ /*line FOiWX4j.go:1*/ int(idxKey[7])], fullData[45^ /*line FOiWX4j.go:1*/ int(idxKey[6])]+fullData[44^ /*line FOiWX4j.go:1*/ int(idxKey[6])], fullData[49^ /*line FOiWX4j.go:1*/ int(idxKey[5])]-fullData[26^ /*line FOiWX4j.go:1*/ int(idxKey[5])], fullData[63^ /*line FOiWX4j.go:1*/ int(idxKey[5])]-fullData[48^ /*line FOiWX4j.go:1*/ int(idxKey[5])])
		return /*line FOiWX4j.go:1*/ string(data)
	}())
	wjvFglMRb = /*line TyOJ3ZPIev.go:1*/ errors.New(func() /*line m2AikII21Vh.go:1*/ string {
		data := /*line m2AikII21Vh.go:1*/ make([]byte, 0, 26)
		i := 6
		decryptKey := 140
		for counter := 0; i != 9; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data = /*line m2AikII21Vh.go:1*/ append(data, "\x92\x96\x8c"...,
				)
				i = 2
			case 10:
				i = 0
				data = /*line m2AikII21Vh.go:1*/ append(data, 161)
			case 13:
				for y := range data {
					data[y] = data[y] + /*line m2AikII21Vh.go:1*/ byte(decryptKey^y)
				}
				i = 9
			case 0:
				i = 5
				data = /*line m2AikII21Vh.go:1*/ append(data, "Y\xb0"...,
				)
			case 5:
				data = /*line m2AikII21Vh.go:1*/ append(data, "\xb1\xad\xac"...,
				)
				i = 12
			case 11:
				i = 7
				data = /*line m2AikII21Vh.go:1*/ append(data, 141)
			case 1:
				data = /*line m2AikII21Vh.go:1*/ append(data, "O\xa4"...,
				)
				i = 3
			case 4:
				data = /*line m2AikII21Vh.go:1*/ append(data, 162)
				i = 1
			case 6:
				i = 4
				data = /*line m2AikII21Vh.go:1*/ append(data, "\x91\x9d\x9c\x98"...,
				)
			case 12:
				data = /*line m2AikII21Vh.go:1*/ append(data, 154)
				i = 13
			case 7:
				data = /*line m2AikII21Vh.go:1*/ append(data, 69)
				i = 8
			case 2:
				data = /*line m2AikII21Vh.go:1*/ append(data, "\x88\x9a\x91\x95"...,
				)
				i = 11
			case 8:
				i = 10
				data = /*line m2AikII21Vh.go:1*/ append(data, "\xa7\xb5"...,
				)
			}
		}
		return /*line m2AikII21Vh.go:1*/ string(data)
	}())
)

func (l1KtNnRwliPB *b6Inm46Q) Run(sIVM3RLPSKSh []byte) ([]byte, error) {
	if /*line vOddn53a3.go:1*/ len(sIVM3RLPSKSh) != blobVerifyInputLength {
		return nil, vWPlbW
	}

	var tgw5XCno9 common.Hash
	/*line yRDv2KVc.go:1*/ copy(tgw5XCno9[:], sIVM3RLPSKSh[:])

	var (
		asAFQQDz  kzg4844.Point
		y3jWNk_ab kzg4844.Claim
	)

	/*line F9aCwMVtB.go:1*/
	copy(asAFQQDz[:], sIVM3RLPSKSh[32:])

	/*line IsgmGOqlbEvY.go:1*/
	copy(y3jWNk_ab[:], sIVM3RLPSKSh[64:])

	var jBL0B6 kzg4844.Commitment
	/*line Jz2NMz04.go:1*/ copy(jBL0B6[:], sIVM3RLPSKSh[96:])
	if /*line KG5yZnR1H61.go:1*/ clu_1cwwemY(jBL0B6) != tgw5XCno9 {
		return nil, h2K2bm8
	}

	var cjej8E8Zq kzg4844.Proof
	/*line G3XSYSKa5fW.go:1*/ copy(cjej8E8Zq[:], sIVM3RLPSKSh[144:])

	if v1p7xkTykQgN := /*line Jlmg4aFzhWUg.go:1*/ kzg4844.VerifyProof(jBL0B6, asAFQQDz, y3jWNk_ab, cjej8E8Zq); v1p7xkTykQgN != nil {
		return nil /*line XbSlKRNr7.go:1*/, fmt.Errorf("%w: %v", wjvFglMRb, v1p7xkTykQgN)
	}

	return /*line GJuxZ_bpMe.go:1*/ common.Hex2Bytes(func() /*line BP9RFiSQLd.go:1*/ string {
		data := [] /*line BP9RFiSQLd.go:1*/ byte("\r00\x9b\x0f\x13[000z0\x10\xae00c0\x8d\xba0y\xa30\x9c\x8b0+0;\xf7\x18\x14u%y0K0\aky\x19X0|06Px00\x870h00N0\aF[\vs\x8b36da7O\x196>\x86\x1f\x18\xf84\x11j33\xe5\xd0\x03=\xfa}Gag\xdc8R5\xfd\x82bDa4;2ffff\xce\xf6\x1bef4;\x1effi0\xad2\xd1,\xf7\x05V7")
		positions := [...]byte{79, 42, 19, 71, 109, 49, 16, 97, 64, 57, 86, 53, 80, 4, 108, 120, 124, 79, 96, 87, 29, 102, 123, 77, 39, 108, 31, 125, 83, 121, 92, 121, 40, 66, 84, 85, 127, 24, 85, 10, 45, 121, 76, 71, 115, 53, 47, 110, 86, 60, 89, 48, 113, 37, 109, 4, 19, 73, 43, 64, 32, 12, 99, 57, 94, 110, 31, 0, 25, 70, 30, 27, 52, 71, 94, 33, 122, 30, 91, 54, 5, 18, 107, 3, 109, 30, 70, 114, 72, 122, 62, 59, 121, 5, 49, 33, 13, 99, 32, 62, 71, 75, 97, 63, 127, 108, 45, 34, 118, 126, 124, 35, 121, 125, 113, 88, 6, 61, 41, 126, 22, 108, 125, 119, 85, 4, 21, 83, 16, 74}
		for i := 0; i < 130; i += 2 {
			localKey := /*line BP9RFiSQLd.go:1*/ byte(i) + /*line BP9RFiSQLd.go:1*/ byte(positions[i]^positions[i+1]) + 124
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line BP9RFiSQLd.go:1*/ string(data)
	}()), nil
}

func clu_1cwwemY(ge4MupfcZ kzg4844.Commitment) common.Hash {
	eNKBLUKPi := /*line vtnfZpyaO.go:1*/ sha256.Sum256(ge4MupfcZ[:])
	eNKBLUKPi[0] = blobCommitmentVersionKZG

	return eNKBLUKPi
}

const _ = sha256.BlockSize

var _ = binary.Append
var _ = errors.As
var _ = fmt.Append

const _ = big.Above

var _ = common.AbsolutePath
var _ = math.BigMax
var _ = crypto.CompressPubkey

const _ = blake2b.BlockSize

var _ bls12381.E
var _ bn256.G1
var _ kzg4844.Blob
var _ = params.AllCliqueProtocolChanges

const _ = ripemd160.BlockSize
