//line :1
package blockchain

import (
	"fmt"
	log "unishard/log"

	"github.com/ethereum/go-ethereum/common"
)

const prunedBlockLimit uint64 = 128

type VE_4Hib struct {
	nfMni6SCvI   GRCKSNhw
	g86EQDJVDIFv map[uint64]BJtoC0zGWx05
	Ctc05xW      uint64
	Q8x5uzuCgKKv uint64
}

type BJtoC0zGWx05 []*y90TfLDGWMoI
type GRCKSNhw map[common.Hash]*y90TfLDGWMoI

type y90TfLDGWMoI struct {
	i48zUQEQZh2 common.Hash
	nF6BWYO     uint64
	hzUDw01     BJtoC0zGWx05

	bv0LsjY5 S9fI7Wj
}

func Ok0v5OJZ3H() *VE_4Hib {
	return &VE_4Hib{
		nfMni6SCvI:/*line haFSWocLC.go:1*/ make(GRCKSNhw),
		g86EQDJVDIFv:/*line IjYTBXS.go:1*/ make(map[uint64]BJtoC0zGWx05),
	}
}

func (xsiB_r *VE_4Hib) PruneUpToWorkerBlockLevel(zziPJBj1a uint64) ([]O0GQK9U1, int, error) {

	var xz1QbM65 int
	iNeba4 := /*line LPthxMyVS71d.go:1*/ make([]O0GQK9U1, 0)
	p2fQIZsi2uG := /*line zkkldrax0aWq.go:1*/ make(map[uint64]bool)
	if zziPJBj1a < xsiB_r.Ctc05xW {
		return nil, xz1QbM65 /*line Vt0UMS9Z.go:1*/, fmt.Errorf(func() /*line zGNHH966a3o5.go:1*/ string {
			key := [] /*line zGNHH966a3o5.go:1*/ byte("-\rV\x16\xebo\x14\x11g\xb3V\x05\x16\x99\xc4\xefB\x7fRfy\x83\x81\xb27m`\xb6\xfdQ\x0eWV\xa9\xe3\x9a\xde\xf4\x94\x8bh͚\x97i\xe0l#\xa9'\xd8ѭ\xb0ѲƟ\x97\xbd\xa7*m\x0f\xac\a \xf9\xd2\xce]<q\x13")
			data := [] /*line zGNHH966a3o5.go:1*/ byte("AX!\n\x81\x00cT\f\xc1\xcagOݡ}ަ\x12\xba\xea\xde\xed\xbc8\a\xc0\xach\xcfe\x16\vÉ˔,\xe0\xdd\xf9\xa1\x86\xd9\t\x85\nF\xc6N\x9bO\xbf\xb1\xa2\xc2Z\xd3η\xba?\x01V\xb8\x19Ll\xa4\x97\x0f\xe4\xb4Q")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line zGNHH966a3o5.go:1*/ string(data)
		}(), zziPJBj1a, xsiB_r.Ctc05xW)
	}

	for rgkb3PXc := zziPJBj1a; rgkb3PXc >= xsiB_r.Ctc05xW && rgkb3PXc > 1; {

		hjnMNeeoM_g := xsiB_r.g86EQDJVDIFv[rgkb3PXc][0].bv0LsjY5
		fpHra3jM, _ := /*line HiFiauAkh4H.go:1*/ hjnMNeeoM_g.ParentWorkerBlock()
		hUN92YLW, xqQ0uRLn9BP := /*line HzSoibK.go:1*/ xsiB_r.GetVertex(fpHra3jM)
		if !xqQ0uRLn9BP || /*line tfiTnX.go:1*/ hUN92YLW.WorkerBlockLevel() < xsiB_r.Ctc05xW {
			break
		}
		p2fQIZsi2uG[ /*line BxZXem3a.go:1*/ hUN92YLW.WorkerBlockLevel()] = true
		rgkb3PXc = /*line gOGBhDg.go:1*/ hUN92YLW.WorkerBlockLevel()
	}

	for rgkb3PXc := xsiB_r.Ctc05xW; rgkb3PXc < zziPJBj1a; rgkb3PXc++ {

		for _, h_5_caanZ := range xsiB_r.g86EQDJVDIFv[rgkb3PXc] {
			if !p2fQIZsi2uG[rgkb3PXc] && rgkb3PXc > 1 {
				if h_5_caanZ.bv0LsjY5 != nil {
					/*line ruL7vR.go:1*/ log.Debugf(func() /*line A7DwCiia.go:1*/ string {
						key := [] /*line A7DwCiia.go:1*/ byte("\x89¤-\x14\x01\xf2\xc3\xf1\xa4B\xdd\x1f=5\x1e#\v\xa1Z,#\x81s\xc1H\x91^\x9b\x930\x06\xff\x04g|A_")
						data := [] /*line A7DwCiia.go:1*/ byte("\xef\xad\xd1Cp!\x93\xe3\x97\xcb0\xb6zY\x15|Od\xc21\x00\x03\xf7\x1a\xa4?\xab~\xbe\xe5\x1c&\x96`]\\d'")
						for i, b := range key {
							data[i] = data[i] ^ b
						}
						return /*line A7DwCiia.go:1*/ string(data)
					}(), /*line kz9qVnx.go:1*/ h_5_caanZ.bv0LsjY5.WorkerBlockLevel() /*line BXOy6W.go:1*/, h_5_caanZ.bv0LsjY5.VertexID())
					iNeba4 = /*line Ll94RYdgx.go:1*/ append(iNeba4 /*line EaBcaE.go:1*/, h_5_caanZ.bv0LsjY5.GetBlock())
				}
			}
		}
	}

	xsiB_r.Ctc05xW = zziPJBj1a

	if zziPJBj1a < prunedBlockLimit {
		return iNeba4, xz1QbM65, nil
	}

	for rgkb3PXc := xsiB_r.Q8x5uzuCgKKv; rgkb3PXc <= zziPJBj1a-prunedBlockLimit; rgkb3PXc++ {

		for _, h_5_caanZ := range xsiB_r.g86EQDJVDIFv[rgkb3PXc] {
			xz1QbM65++
			/*line faRZ_hBE.go:1*/ delete(xsiB_r.nfMni6SCvI, h_5_caanZ.i48zUQEQZh2)
		}
		/*line MdFCB_VC.go:1*/ delete(xsiB_r.g86EQDJVDIFv, rgkb3PXc)
	}

	xsiB_r.Q8x5uzuCgKKv = zziPJBj1a - prunedBlockLimit

	return iNeba4, xz1QbM65, nil
}

func (xsiB_r *VE_4Hib) PruneUpToCoordinationBlockLevel(zziPJBj1a uint64) ([]O0GQK9U1, int, error) {

	var xz1QbM65 int
	iNeba4 := /*line Neb0Tfjm.go:1*/ make([]O0GQK9U1, 0)
	p2fQIZsi2uG := /*line C5XpxExB.go:1*/ make(map[uint64]bool)
	if zziPJBj1a < xsiB_r.Ctc05xW {
		return nil, xz1QbM65 /*line q5z_pA.go:1*/, fmt.Errorf(func() /*line aeO7uJ.go:1*/ string {
			data := [] /*line aeO7uJ.go:1*/ byte("\xb8\x17w lo\xa7e\x8d\xe0\xb0\x01=v\xcd\x05\x9e\x86˺\x98\xe5\x99noB\xfabe 1\xb1alleȜ\\uar \xber\x9fvw\x04ud [ݶ\xf7\x9f\xb1etŃ{\xceX\x12\xf0e\xc2\xdb\xed %\x1e")
			positions := [...]byte{62, 38, 1, 19, 65, 56, 25, 22, 16, 10, 8, 14, 43, 64, 8, 55, 69, 30, 10, 6, 22, 21, 50, 38, 70, 21, 57, 69, 73, 10, 16, 38, 47, 41, 54, 61, 12, 26, 62, 69, 63, 15, 20, 11, 20, 68, 17, 52, 62, 62, 55, 31, 57, 10, 57, 45, 0, 60, 9, 25, 52, 37, 30, 21, 66, 39, 43, 70, 25, 48, 8, 36, 66, 22, 15, 14, 69, 12, 65, 6, 6, 53, 62, 61, 25, 1, 63, 18, 56, 48}
			for i := 0; i < 90; i += 2 {
				localKey := /*line aeO7uJ.go:1*/ byte(i) + /*line aeO7uJ.go:1*/ byte(positions[i]^positions[i+1]) + 227
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line aeO7uJ.go:1*/ string(data)
		}(), zziPJBj1a, xsiB_r.Ctc05xW)
	}

	for rgkb3PXc := zziPJBj1a; rgkb3PXc >= xsiB_r.Ctc05xW && rgkb3PXc > 1; {

		hjnMNeeoM_g := xsiB_r.g86EQDJVDIFv[rgkb3PXc][0].bv0LsjY5
		fpHra3jM, _ := /*line BdIaWuB8RP.go:1*/ hjnMNeeoM_g.ParentCoordinationBlock()
		hUN92YLW, xqQ0uRLn9BP := /*line cavIxC.go:1*/ xsiB_r.GetVertex(fpHra3jM)
		if !xqQ0uRLn9BP || /*line mf2L2GY.go:1*/ hUN92YLW.CoordinationBlockLevel() < xsiB_r.Ctc05xW {
			break
		}
		p2fQIZsi2uG[ /*line bOSNtf9V.go:1*/ hUN92YLW.CoordinationBlockLevel()] = true
		rgkb3PXc = /*line wLtSnDc5Q.go:1*/ hUN92YLW.CoordinationBlockLevel()
	}

	for rgkb3PXc := xsiB_r.Ctc05xW; rgkb3PXc < zziPJBj1a; rgkb3PXc++ {

		for _, h_5_caanZ := range xsiB_r.g86EQDJVDIFv[rgkb3PXc] {
			if !p2fQIZsi2uG[rgkb3PXc] && rgkb3PXc > 1 {
				if h_5_caanZ.bv0LsjY5 != nil {
					/*line hRJYIj8SoVr.go:1*/ log.Debugf(func() /*line Bm4SXtfsay.go:1*/ string {
						data := [] /*line Bm4SXtfsay.go:1*/ byte("f}j5d\x1d\x96\x89fotkz|\xa4$xjtk, \x12\rdW~w%\x83\x91 id:\xae\x91\x9c")
						positions := [...]byte{22, 7, 5, 23, 5, 17, 30, 15, 36, 3, 29, 1, 27, 16, 23, 14, 13, 18, 24, 27, 17, 27, 12, 2, 26, 26, 37, 6, 25, 35, 1, 16, 24, 3, 26, 22, 5, 15, 10, 6, 18, 18, 6, 26, 23, 10, 5, 23}
						for i := 0; i < 48; i += 2 {
							localKey := /*line Bm4SXtfsay.go:1*/ byte(i) + /*line Bm4SXtfsay.go:1*/ byte(positions[i]^positions[i+1]) + 225
							data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
						}
						return /*line Bm4SXtfsay.go:1*/ string(data)
					}(), /*line GCzPQm.go:1*/ h_5_caanZ.bv0LsjY5.CoordinationBlockLevel() /*line wz1ODI.go:1*/, h_5_caanZ.bv0LsjY5.VertexID())
					iNeba4 = /*line jnrfMzg.go:1*/ append(iNeba4 /*line jOViSLYdW9H5.go:1*/, h_5_caanZ.bv0LsjY5.GetBlock())
				}
			}
		}
	}

	xsiB_r.Ctc05xW = zziPJBj1a

	if zziPJBj1a < prunedBlockLimit {
		return iNeba4, xz1QbM65, nil
	}

	for rgkb3PXc := xsiB_r.Q8x5uzuCgKKv; rgkb3PXc <= zziPJBj1a-prunedBlockLimit; rgkb3PXc++ {

		for _, h_5_caanZ := range xsiB_r.g86EQDJVDIFv[rgkb3PXc] {
			xz1QbM65++
			/*line qlO3QP.go:1*/ delete(xsiB_r.nfMni6SCvI, h_5_caanZ.i48zUQEQZh2)
		}
		/*line cWteUdnZjg.go:1*/ delete(xsiB_r.g86EQDJVDIFv, rgkb3PXc)
	}

	xsiB_r.Q8x5uzuCgKKv = zziPJBj1a - prunedBlockLimit

	return iNeba4, xz1QbM65, nil
}

func (xsiB_r *VE_4Hib) HasVertex(fQYKB1FX5mQj common.Hash) bool {
	aSzFIBKnn6, oF5Xkii := xsiB_r.nfMni6SCvI[fQYKB1FX5mQj]
	return oF5Xkii && ! /*line vEayZAiBCy.go:1*/ xsiB_r.d_TYsnDG301f(aSzFIBKnn6)
}

func (xsiB_r *VE_4Hib) d_TYsnDG301f(y90TfLDGWMoI *y90TfLDGWMoI) bool {
	return y90TfLDGWMoI.bv0LsjY5 == nil
}

func (xsiB_r *VE_4Hib) GetVertex(fQYKB1FX5mQj common.Hash) (S9fI7Wj, bool) {
	aSzFIBKnn6, oF5Xkii := xsiB_r.nfMni6SCvI[fQYKB1FX5mQj]
	if !oF5Xkii || /*line RLoY2jrRLxMm.go:1*/ xsiB_r.d_TYsnDG301f(aSzFIBKnn6) {
		return nil, false
	}
	return aSzFIBKnn6.bv0LsjY5, true
}

func (xsiB_r *VE_4Hib) GetChildren(fQYKB1FX5mQj common.Hash) ZOPXV6ybU {
	aSzFIBKnn6 := xsiB_r.nfMni6SCvI[fQYKB1FX5mQj]

	return /*line lGGdNPj.go:1*/ qP0V4i3Ogp_(aSzFIBKnn6.hzUDw01)
}

func (xsiB_r *VE_4Hib) GetNumberOfChildren(fQYKB1FX5mQj common.Hash) int {
	aSzFIBKnn6 := xsiB_r.nfMni6SCvI[fQYKB1FX5mQj]
	vuEjfGzPKDi9 := 0
	for _, jcve0JEXylM := range aSzFIBKnn6.hzUDw01 {
		if jcve0JEXylM.bv0LsjY5 != nil {
			vuEjfGzPKDi9++
		}
	}
	return vuEjfGzPKDi9
}

func (xsiB_r *VE_4Hib) GetVerticesAtLevel(zziPJBj1a uint64) ZOPXV6ybU {
	return /*line TS_P_L0HF4s.go:1*/ qP0V4i3Ogp_(xsiB_r.g86EQDJVDIFv[zziPJBj1a])
}

func (xsiB_r *VE_4Hib) AddWorkerBlockVertex(hjnMNeeoM_g S9fI7Wj) {
	if /*line UwlXZvyNWD9.go:1*/ hjnMNeeoM_g.WorkerBlockLevel() < xsiB_r.Ctc05xW {
		return
	}
	aSzFIBKnn6 := /*line zFJYkldaf.go:1*/ xsiB_r.iwZGLIzT( /*line pzTWlaV.go:1*/ hjnMNeeoM_g.VertexID() /*line Bb48N5Wzgdc.go:1*/, hjnMNeeoM_g.WorkerBlockLevel())
	if ! /*line _9IizlxvV2.go:1*/ xsiB_r.d_TYsnDG301f(aSzFIBKnn6) {
		return
	}

	aSzFIBKnn6.bv0LsjY5 = hjnMNeeoM_g
	/*line iN68kr3Q5Hz.go:1*/ xsiB_r.hCsTRM(aSzFIBKnn6)
}
func (xsiB_r *VE_4Hib) AddCoordinationBlockVertex(hjnMNeeoM_g S9fI7Wj) {
	if /*line ELWhBYY_kKm_.go:1*/ hjnMNeeoM_g.CoordinationBlockLevel() < xsiB_r.Ctc05xW {
		return
	}
	aSzFIBKnn6 := /*line qlJA0yzu.go:1*/ xsiB_r.iwZGLIzT( /*line vtYYaII9H.go:1*/ hjnMNeeoM_g.VertexID() /*line EMr0Fgac.go:1*/, hjnMNeeoM_g.CoordinationBlockLevel())
	if ! /*line mMBcyH1Gz5EZ.go:1*/ xsiB_r.d_TYsnDG301f(aSzFIBKnn6) {
		return
	}

	aSzFIBKnn6.bv0LsjY5 = hjnMNeeoM_g
	/*line HUL6dIeqT.go:1*/ xsiB_r.wslGd_xIuT(aSzFIBKnn6)
}

func (xsiB_r *VE_4Hib) hCsTRM(y90TfLDGWMoI *y90TfLDGWMoI) {

	if y90TfLDGWMoI.nF6BWYO <= xsiB_r.Ctc05xW {
		return
	}

	_, irHKT_H8vt2d := /*line QL79Lv_yvX1A.go:1*/ y90TfLDGWMoI.bv0LsjY5.ParentWorkerBlock()
	if irHKT_H8vt2d < xsiB_r.Ctc05xW {
		return
	}
	eAc6JnCF := /*line lGILGpq.go:1*/ xsiB_r.iwZGLIzT( /*line mgRas7M4.go:1*/ y90TfLDGWMoI.bv0LsjY5.ParentWorkerBlock())
	eAc6JnCF.hzUDw01 = /*line adGC0rk1U.go:1*/ append(eAc6JnCF.hzUDw01, y90TfLDGWMoI)
}

func (xsiB_r *VE_4Hib) wslGd_xIuT(y90TfLDGWMoI *y90TfLDGWMoI) {

	if y90TfLDGWMoI.nF6BWYO <= xsiB_r.Ctc05xW {
		return
	}

	_, irHKT_H8vt2d := /*line tFfCoWTYbh.go:1*/ y90TfLDGWMoI.bv0LsjY5.ParentCoordinationBlock()
	if irHKT_H8vt2d < xsiB_r.Ctc05xW {
		return
	}
	eAc6JnCF := /*line B13ArkYO0WO.go:1*/ xsiB_r.iwZGLIzT( /*line eBjpWRr.go:1*/ y90TfLDGWMoI.bv0LsjY5.ParentCoordinationBlock())
	eAc6JnCF.hzUDw01 = /*line agdtVYKiKaP.go:1*/ append(eAc6JnCF.hzUDw01, y90TfLDGWMoI)
}

func (xsiB_r *VE_4Hib) iwZGLIzT(fQYKB1FX5mQj common.Hash, zziPJBj1a uint64) *y90TfLDGWMoI {
	aSzFIBKnn6, oF5Xkii := xsiB_r.nfMni6SCvI[fQYKB1FX5mQj]
	if !oF5Xkii {
		aSzFIBKnn6 = &y90TfLDGWMoI{
			i48zUQEQZh2: fQYKB1FX5mQj,
			nF6BWYO:     zziPJBj1a,
		}
		xsiB_r.nfMni6SCvI[aSzFIBKnn6.i48zUQEQZh2] = aSzFIBKnn6
		xj7E3dvlfZM := xsiB_r.g86EQDJVDIFv[aSzFIBKnn6.nF6BWYO]
		xsiB_r.g86EQDJVDIFv[aSzFIBKnn6.nF6BWYO] = /*line UzUvR4e48.go:1*/ append(xj7E3dvlfZM, aSzFIBKnn6)
	}
	return aSzFIBKnn6
}

func (xsiB_r *VE_4Hib) mdyvXXSXNs(zziPJBj1a uint64) {
	for _, h_5_caanZ := range xsiB_r.g86EQDJVDIFv[zziPJBj1a] {
		/*line ozayBaCW8.go:1*/ delete(xsiB_r.nfMni6SCvI, h_5_caanZ.i48zUQEQZh2)
	}
	/*line wmaawE.go:1*/ delete(xsiB_r.g86EQDJVDIFv, zziPJBj1a)
}

var _ = fmt.Append
var _ = log.CDebpj
var _ = common.AbsolutePath
