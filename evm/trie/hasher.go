//line :1
package trie

import (
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type d6imPJ struct {
	vbM2IWKj6j   crypto.KeccakState
	cNadhB3aI    []byte
	urdUwF       rlp.EncoderBuffer
	xqxEulHqlWrc bool
}

var ataGrYxs2 = sync.Pool{
	New: func() interface{} {
		return &d6imPJ{
			cNadhB3aI:/*line HQuvYWt6A0.go:1*/ make([]byte, 0, 550),
			vbM2IWKj6j:/*line hta9yK.go:1*/ sha3.NewLegacyKeccak256().(crypto.KeccakState),
			urdUwF:/*line apn_gi4.go:1*/ rlp.NewEncoderBuffer(nil),
		}
	},
}

func dCZJziX(jrJP19nCXDet bool) *d6imPJ {
	klE3zy := /*line wC2gcbu.go:1*/ ataGrYxs2.Get().(*d6imPJ)
	klE3zy.xqxEulHqlWrc = jrJP19nCXDet
	return klE3zy
}

func ptIMYRK9(klE3zy *d6imPJ) {
	/*line _8acsdzAepJ.go:1*/ ataGrYxs2.Put(klE3zy)
}

func (klE3zy *d6imPJ) rNuN0sMPDJ(gnGMaXTuiFeE node, mAlQwO bool) (iVcAtTyHaK node, kdSzj0 node) {

	if rNuN0sMPDJ, _ := /*line z0GQIyN2aWfX.go:1*/ gnGMaXTuiFeE.tGJzZYYLvK(); rNuN0sMPDJ != nil {
		return rNuN0sMPDJ, gnGMaXTuiFeE
	}

	switch gnGMaXTuiFeE := gnGMaXTuiFeE.(type) {
	case *qUKQUCfTXB:
		iRpa_yqzzsFI, kdSzj0 := /*line Pf5qwPbaLqG.go:1*/ klE3zy.hOn0b1lf1(gnGMaXTuiFeE)
		iVcAtTyHaK := /*line p9d7msD3.go:1*/ klE3zy.c8RixUz(iRpa_yqzzsFI, mAlQwO)

		if ifmSiSN3, yY_yPSd8vG := iVcAtTyHaK.(hashNode); yY_yPSd8vG {
			kdSzj0.d6dipJ.hash = ifmSiSN3
		} else {
			kdSzj0.d6dipJ.hash = nil
		}
		return iVcAtTyHaK, kdSzj0
	case *fullNode:
		iRpa_yqzzsFI, kdSzj0 := /*line RFpeI3.go:1*/ klE3zy.k4i1rgXhDgLW(gnGMaXTuiFeE)
		iVcAtTyHaK = /*line B8UvGdHPhz.go:1*/ klE3zy.bHDsn6fZ(iRpa_yqzzsFI, mAlQwO)
		if ifmSiSN3, yY_yPSd8vG := iVcAtTyHaK.(hashNode); yY_yPSd8vG {
			kdSzj0.flags.hash = ifmSiSN3
		} else {
			kdSzj0.flags.hash = nil
		}
		return iVcAtTyHaK, kdSzj0
	default:

		return gnGMaXTuiFeE, gnGMaXTuiFeE
	}
}

func (klE3zy *d6imPJ) hOn0b1lf1(gnGMaXTuiFeE *qUKQUCfTXB) (iRpa_yqzzsFI, kdSzj0 *qUKQUCfTXB) {

	iRpa_yqzzsFI, kdSzj0 = /*line JGlEyaUcEL.go:1*/ gnGMaXTuiFeE.sBTiL7Ci() /*line J7E1X38wG4A9.go:1*/, gnGMaXTuiFeE.sBTiL7Ci()

	iRpa_yqzzsFI.ANdZYqbzJ1A = /*line puL75o.go:1*/ pSumbut7Z(gnGMaXTuiFeE.ANdZYqbzJ1A)

	switch gnGMaXTuiFeE.YpmEYGB.(type) {
	case *fullNode, *qUKQUCfTXB:
		iRpa_yqzzsFI.YpmEYGB, kdSzj0.YpmEYGB = /*line znW00l.go:1*/ klE3zy.rNuN0sMPDJ(gnGMaXTuiFeE.YpmEYGB, false)
	}
	return iRpa_yqzzsFI, kdSzj0
}

func (klE3zy *d6imPJ) k4i1rgXhDgLW(gnGMaXTuiFeE *fullNode) (iRpa_yqzzsFI *fullNode, kdSzj0 *fullNode) {

	kdSzj0 = /*line RxZaLLBg.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
	iRpa_yqzzsFI = /*line JHvalvDtQz7.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
	if klE3zy.xqxEulHqlWrc {
		var vgNRFx2c1h2e sync.WaitGroup
		/*line F_eK1sniq64R.go:1*/ vgNRFx2c1h2e.Add(16)
		for q2u2020KD6 := 0; q2u2020KD6 < 16; q2u2020KD6++ {
			go func( /*line AHA7yoiGUgk.go:1*/ q2u2020KD6 int) {
				d6imPJ := /*line MayftDR5eSI.go:1*/ dCZJziX(false)
				if jcDLabJ7o := gnGMaXTuiFeE.Children[q2u2020KD6]; jcDLabJ7o != nil {
					iRpa_yqzzsFI.Children[q2u2020KD6], kdSzj0.Children[q2u2020KD6] = /*line yayvWS.go:1*/ d6imPJ.rNuN0sMPDJ(jcDLabJ7o, false)
				} else {
					iRpa_yqzzsFI.Children[q2u2020KD6] = e6jPsAvc
				}
				/*line He1Y4igHll.go:1*/ ptIMYRK9(d6imPJ)
				/*line NBbW24IG.go:1*/ vgNRFx2c1h2e.Done()
			}(q2u2020KD6)
		}
		/*line AblPEw50Yl.go:1*/ vgNRFx2c1h2e.Wait()
	} else {
		for q2u2020KD6 := 0; q2u2020KD6 < 16; q2u2020KD6++ {
			if jcDLabJ7o := gnGMaXTuiFeE.Children[q2u2020KD6]; jcDLabJ7o != nil {
				iRpa_yqzzsFI.Children[q2u2020KD6], kdSzj0.Children[q2u2020KD6] = /*line avGGfpt.go:1*/ klE3zy.rNuN0sMPDJ(jcDLabJ7o, false)
			} else {
				iRpa_yqzzsFI.Children[q2u2020KD6] = e6jPsAvc
			}
		}
	}
	return iRpa_yqzzsFI, kdSzj0
}

func (klE3zy *d6imPJ) c8RixUz(gnGMaXTuiFeE *qUKQUCfTXB, mAlQwO bool) node {
	/*line GyJy93yK.go:1*/ gnGMaXTuiFeE.ta85dJ1aZ(klE3zy.urdUwF)
	bh98pCjEm := /*line LLZ3c8h8qgH.go:1*/ klE3zy.ucIRgT1yG()

	if /*line v3_RYS79HQ.go:1*/ len(bh98pCjEm) < 32 && !mAlQwO {
		return gnGMaXTuiFeE
	}
	return /*line Bv7syP.go:1*/ klE3zy.jfg9zxq(bh98pCjEm)
}

func (klE3zy *d6imPJ) bHDsn6fZ(gnGMaXTuiFeE *fullNode, mAlQwO bool) node {
	/*line Cs3wgac2xX.go:1*/ gnGMaXTuiFeE.ta85dJ1aZ(klE3zy.urdUwF)
	bh98pCjEm := /*line ImhbeZ.go:1*/ klE3zy.ucIRgT1yG()

	if /*line ZlAlZJ.go:1*/ len(bh98pCjEm) < 32 && !mAlQwO {
		return gnGMaXTuiFeE
	}
	return /*line F2i6aANSj_2L.go:1*/ klE3zy.jfg9zxq(bh98pCjEm)
}

func (klE3zy *d6imPJ) ucIRgT1yG() []byte {
	klE3zy.cNadhB3aI = /*line IvRtpa.go:1*/ klE3zy.urdUwF.AppendToBytes(klE3zy.cNadhB3aI[:0])
	/*line vIduv467o.go:1*/ klE3zy.urdUwF.Reset(nil)
	return klE3zy.cNadhB3aI
}

func (klE3zy *d6imPJ) jfg9zxq(qNY8zG []byte) hashNode {
	gnGMaXTuiFeE := /*line Iqnrs3QAv.go:1*/ make(hashNode, 32)
	/*line GPhzUOu85Ju.go:1*/ klE3zy.vbM2IWKj6j.Reset()
	/*line Jx1tsWM.go:1*/ klE3zy.vbM2IWKj6j.Write(qNY8zG)
	/*line TwSrH4.go:1*/ klE3zy.vbM2IWKj6j.Read(gnGMaXTuiFeE)
	return gnGMaXTuiFeE
}

func (klE3zy *d6imPJ) aCBICr2(v4UNa4AppV node) (iRpa_yqzzsFI, iVcAtTyHaK node) {
	switch gnGMaXTuiFeE := v4UNa4AppV.(type) {
	case *qUKQUCfTXB:
		pOsORZ8, _ := /*line I9x7nT41E.go:1*/ klE3zy.hOn0b1lf1(gnGMaXTuiFeE)
		return pOsORZ8 /*line FSxHll8xm.go:1*/, klE3zy.c8RixUz(pOsORZ8, false)
	case *fullNode:
		qC6kc1PxI8, _ := /*line xaGFc4ncqv.go:1*/ klE3zy.k4i1rgXhDgLW(gnGMaXTuiFeE)
		return qC6kc1PxI8 /*line g6qhnBzMv.go:1*/, klE3zy.bHDsn6fZ(qC6kc1PxI8, false)
	default:

		return gnGMaXTuiFeE, gnGMaXTuiFeE
	}
}

var _ sync.Cond
var _ = crypto.CompressPubkey
var _ = rlp.AppendUint64
var _ = sha3.New224
