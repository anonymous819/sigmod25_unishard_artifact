//line :1
package trie

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie/trienode"
)

type nO6upQo struct {
	bRTjxOq    *trienode.NodeSet
	b2r12tecD  *tracer
	iOE_Oy20yk bool
}

func egS1cdL6TQG(mff5TAZylrk *trienode.NodeSet, oM5uOnaI7X *tracer, xfPJhFZM25R bool) *nO6upQo {
	return &nO6upQo{
		bRTjxOq:    mff5TAZylrk,
		b2r12tecD:  oM5uOnaI7X,
		iOE_Oy20yk: xfPJhFZM25R,
	}
}

func (arf3iiVz *nO6upQo) Commit(gnGMaXTuiFeE node) hashNode {
	return /*line oc0bQaqQ7.go:1*/ arf3iiVz.vpjQVV(nil, gnGMaXTuiFeE).(hashNode)
}

func (arf3iiVz *nO6upQo) vpjQVV(qiRG6lpeaFl []byte, gnGMaXTuiFeE node) node {

	rNuN0sMPDJ, c0YPpA0vA5f := /*line aaDDgrpyhm.go:1*/ gnGMaXTuiFeE.tGJzZYYLvK()
	if rNuN0sMPDJ != nil && !c0YPpA0vA5f {
		return rNuN0sMPDJ
	}

	switch bXl46lLQR := gnGMaXTuiFeE.(type) {
	case *qUKQUCfTXB:

		iRpa_yqzzsFI := /*line yTkXE4Ls2VJ.go:1*/ bXl46lLQR.sBTiL7Ci()

		if _, yY_yPSd8vG := bXl46lLQR.YpmEYGB.(*fullNode); yY_yPSd8vG {
			iRpa_yqzzsFI.YpmEYGB = /*line awzhpCzm.go:1*/ arf3iiVz.vpjQVV( /*line jl1BA0J_WoL.go:1*/ append(qiRG6lpeaFl, bXl46lLQR.ANdZYqbzJ1A...), bXl46lLQR.YpmEYGB)
		}

		iRpa_yqzzsFI.ANdZYqbzJ1A = /*line NbL_ZsB.go:1*/ pSumbut7Z(bXl46lLQR.ANdZYqbzJ1A)
		klVnl8zZd := /*line bfB1f1JG.go:1*/ arf3iiVz.wAO5xfUp97X(qiRG6lpeaFl, iRpa_yqzzsFI)
		if ifmSiSN3, yY_yPSd8vG := klVnl8zZd.(hashNode); yY_yPSd8vG {
			return ifmSiSN3
		}
		return iRpa_yqzzsFI
	case *fullNode:
		q0WWgK1KXhVQ := /*line L09CxNi3.go:1*/ arf3iiVz.rFK_Un119XB(qiRG6lpeaFl, bXl46lLQR)
		iRpa_yqzzsFI := /*line BNnLBZJNgl1H.go:1*/ bXl46lLQR.sBTiL7Ci()
		iRpa_yqzzsFI.Children = q0WWgK1KXhVQ

		klVnl8zZd := /*line OU1R_IooNX46.go:1*/ arf3iiVz.wAO5xfUp97X(qiRG6lpeaFl, iRpa_yqzzsFI)
		if ifmSiSN3, yY_yPSd8vG := klVnl8zZd.(hashNode); yY_yPSd8vG {
			return ifmSiSN3
		}
		return iRpa_yqzzsFI
	case hashNode:
		return bXl46lLQR
	default:

		/*line A9akXvpkh0.go:1*/
		panic( /*line AoDIti2pS.go:1*/ fmt.Sprintf(func() /*line cfGg5h.go:1*/ string {
			data := /*line cfGg5h.go:1*/ make([]byte, 0, 21)
			i := 1
			decryptKey := 92
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 9
					data = /*line cfGg5h.go:1*/ append(data, "\faz4"...,
					)
				case 0:
					i = 2
					data = /*line cfGg5h.go:1*/ append(data, "?=9"...,
					)
				case 8:
					data = /*line cfGg5h.go:1*/ append(data, "33s"...,
					)
					i = 3
				case 5:
					i = 7
					for y := range data {
						data[y] = data[y] ^ /*line cfGg5h.go:1*/ byte(decryptKey^y)
					}
				case 2:
					i = 4
					data = /*line cfGg5h.go:1*/ append(data, "7r"...,
					)
				case 1:
					data = /*line cfGg5h.go:1*/ append(data, 124)
					i = 6
				case 9:
					data = /*line cfGg5h.go:1*/ append(data, "2)"...,
					)
					i = 0
				case 4:
					i = 8
					data = /*line cfGg5h.go:1*/ append(data, ";;"...,
					)
				case 3:
					i = 5
					data = /*line cfGg5h.go:1*/ append(data, "hn<"...,
					)
				}
			}
			return /*line cfGg5h.go:1*/ string(data)
		}(), gnGMaXTuiFeE, gnGMaXTuiFeE))
	}
}

func (arf3iiVz *nO6upQo) rFK_Un119XB(qiRG6lpeaFl []byte, gnGMaXTuiFeE *fullNode) [17]node {
	var iHZ_vZ8aLL [17]node
	for q2u2020KD6 := 0; q2u2020KD6 < 16; q2u2020KD6++ {
		jcDLabJ7o := gnGMaXTuiFeE.Children[q2u2020KD6]
		if jcDLabJ7o == nil {
			continue
		}

		if ifmSiSN3, yY_yPSd8vG := jcDLabJ7o.(hashNode); yY_yPSd8vG {
			iHZ_vZ8aLL[q2u2020KD6] = ifmSiSN3
			continue
		}

		iHZ_vZ8aLL[q2u2020KD6] = /*line O8MjkmOocH.go:1*/ arf3iiVz.vpjQVV( /*line XHPtEd5aIodq.go:1*/ append(qiRG6lpeaFl /*line raaObJVDRmcf.go:1*/, byte(q2u2020KD6)), jcDLabJ7o)
	}

	if gnGMaXTuiFeE.Children[16] != nil {
		iHZ_vZ8aLL[16] = gnGMaXTuiFeE.Children[16]
	}
	return iHZ_vZ8aLL
}

func (arf3iiVz *nO6upQo) wAO5xfUp97X(qiRG6lpeaFl []byte, gnGMaXTuiFeE node) node {

	var rNuN0sMPDJ, _ = /*line zIrGXDSgDHDc.go:1*/ gnGMaXTuiFeE.tGJzZYYLvK()

	if rNuN0sMPDJ == nil {

		_, yY_yPSd8vG := arf3iiVz.b2r12tecD.accessList[ /*line qUk7aoH.go:1*/ string(qiRG6lpeaFl)]
		if yY_yPSd8vG {
			/*line VvUOa7j6dCnK.go:1*/ arf3iiVz.bRTjxOq.AddNode(qiRG6lpeaFl /*line auaHOm.go:1*/, trienode.NewDeleted())
		}
		return gnGMaXTuiFeE
	}

	e6aJjpmHDa8 := /*line UymyIDE1gsC.go:1*/ common.BytesToHash(rNuN0sMPDJ)
	/*line YM9RJQVS2HQf.go:1*/ arf3iiVz.bRTjxOq.AddNode(qiRG6lpeaFl /*line Gyy8Wt.go:1*/, trienode.New(e6aJjpmHDa8 /*line Fif0zc.go:1*/, kG18lbsZr(gnGMaXTuiFeE)))

	if arf3iiVz.iOE_Oy20yk {
		if pOsORZ8, yY_yPSd8vG := gnGMaXTuiFeE.(*qUKQUCfTXB); yY_yPSd8vG {
			if h_pxYek4zT, yY_yPSd8vG := pOsORZ8.YpmEYGB.(valueNode); yY_yPSd8vG {
				/*line BKlrbD.go:1*/ arf3iiVz.bRTjxOq.AddLeaf(e6aJjpmHDa8, h_pxYek4zT)
			}
		}
	}
	return rNuN0sMPDJ
}

type SDUpMtPV struct{}

func (aRra_C_9g SDUpMtPV) ForEach(uAjz8NxU_cL []byte, bk2DZUvPms func(common.Hash)) {
	/*line vAuUtA_MSp.go:1*/ bLpKhe4kLfy( /*line U5zOsU.go:1*/ hct27E(nil, uAjz8NxU_cL), bk2DZUvPms)
}

func bLpKhe4kLfy(gnGMaXTuiFeE node, bk2DZUvPms func(rNuN0sMPDJ common.Hash)) {
	switch gnGMaXTuiFeE := gnGMaXTuiFeE.(type) {
	case *qUKQUCfTXB:
		/*line Fb1pjXiVKGJ.go:1*/ bLpKhe4kLfy(gnGMaXTuiFeE.YpmEYGB, bk2DZUvPms)
	case *fullNode:
		for q2u2020KD6 := 0; q2u2020KD6 < 16; q2u2020KD6++ {
			/*line Dkwg7ZeuFWM1.go:1*/ bLpKhe4kLfy(gnGMaXTuiFeE.Children[q2u2020KD6], bk2DZUvPms)
		}
	case hashNode:
		/*line B7GKlGFRnul.go:1*/ bk2DZUvPms( /*line rlWA4O.go:1*/ common.BytesToHash(gnGMaXTuiFeE))
	case valueNode, nil:
	default:
		/*line QAixqxVxT8u.go:1*/ panic( /*line wrJ8pHTlv.go:1*/ fmt.Sprintf(func() /*line h3VNaFG.go:1*/ string {
			data := [] /*line h3VNaFG.go:1*/ byte("P\x12L?\xeb\x1a\xf9En\rde;\xf5y)K\xfbF%T")
			positions := [...]byte{12, 13, 5, 15, 16, 9, 12, 7, 17, 7, 5, 12, 16, 3, 18, 1, 15, 12, 13, 2, 0, 3, 12, 6, 18, 4}
			for i := 0; i < 26; i += 2 {
				localKey := /*line h3VNaFG.go:1*/ byte(i) + /*line h3VNaFG.go:1*/ byte(positions[i]^positions[i+1]) + 7
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line h3VNaFG.go:1*/ string(data)
		}(), gnGMaXTuiFeE))
	}
}

var _ = fmt.Append
var _ = common.AbsolutePath
var _ trienode.MergedNodeSet
