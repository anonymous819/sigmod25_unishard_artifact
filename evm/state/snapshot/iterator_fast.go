//line :1
package snapshot

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

type xlNMI97NK_ struct {
	mZR9T1H3Tp65 FcPREnURc
	eyJQwba      int
}

func (cu8ZKpmK *xlNMI97NK_) Cmp(lzcwPxiAHf *xlNMI97NK_) int {

	a9dhFIMqJL1O := /*line dESgjV6P.go:1*/ cu8ZKpmK.mZR9T1H3Tp65.Hash()
	fWLSPhu54z := /*line ZV53kOk.go:1*/ lzcwPxiAHf.mZR9T1H3Tp65.Hash()

	switch /*line IHoD2qX0ih.go:1*/ bytes.Compare(a9dhFIMqJL1O[:], fWLSPhu54z[:]) {
	case -1:
		return -1
	case 1:
		return 1
	}

	if cu8ZKpmK.eyJQwba < lzcwPxiAHf.eyJQwba {
		return -1
	}
	if cu8ZKpmK.eyJQwba > lzcwPxiAHf.eyJQwba {
		return 1
	}
	return 0
}

type bJzmW7YOgIxt struct {
	cIT8woywCq1 *CE2m1DUB6wW
	xXYPhv2w    common.Hash

	nrTwyBJ      []byte
	yxDFcc_lELc5 []byte

	qR2G2CL      []*xlNMI97NK_
	r4DzlIJJPs   bool
	_EgzAfKfzWaQ bool
	zJG9at0m     error
}

func l8VM92(abbWTExt *CE2m1DUB6wW, z1BBTN2UX common.Hash, evzmhL1 common.Hash, gnGBaeoLX common.Hash, tw1Vjc bool) (*bJzmW7YOgIxt, error) {
	pc2Fq1jYVnOO := /*line m9rJRA7GRy0A.go:1*/ abbWTExt.Snapshot(z1BBTN2UX)
	if pc2Fq1jYVnOO == nil {
		return nil /*line Tn41guQOFzf.go:1*/, fmt.Errorf(func() /*line NwkEGtisVVXh.go:1*/ string {
			key := [] /*line NwkEGtisVVXh.go:1*/ byte("\t\xac_\xc6|\xc8\xdc*;\xcc|]\xad)>%N\xb3\x80\xb5")
			data := [] /*line NwkEGtisVVXh.go:1*/ byte("|\xc24\xa8\x13\xbf\xb2\nH\xa2\x1d-\xdeAQQt\x93\xa5\xcd")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line NwkEGtisVVXh.go:1*/ string(data)
		}(), z1BBTN2UX)
	}
	lffjFkFU5 := &bJzmW7YOgIxt{
		cIT8woywCq1:  abbWTExt,
		xXYPhv2w:     z1BBTN2UX,
		_EgzAfKfzWaQ: tw1Vjc,
	}
	y84IW11J := pc2Fq1jYVnOO.(snapshot)
	for tCwDahj07TnV := 0; y84IW11J != nil; tCwDahj07TnV++ {
		if tw1Vjc {
			lffjFkFU5.qR2G2CL = /*line vlE_IY3WbgME.go:1*/ append(lffjFkFU5.qR2G2CL, &xlNMI97NK_{
				mZR9T1H3Tp65:/*line dASsJS.go:1*/ y84IW11J.AccountIterator(gnGBaeoLX),
				eyJQwba: tCwDahj07TnV,
			})
		} else {

			cu8ZKpmK, oiHhcQbh9XAf := /*line N5LQPU1.go:1*/ y84IW11J.StorageIterator(evzmhL1, gnGBaeoLX)
			lffjFkFU5.qR2G2CL = /*line hlM8T4qzWJw.go:1*/ append(lffjFkFU5.qR2G2CL, &xlNMI97NK_{
				mZR9T1H3Tp65: cu8ZKpmK,
				eyJQwba:      tCwDahj07TnV,
			})
			if oiHhcQbh9XAf {
				break
			}
		}
		y84IW11J = /*line S2YbfY2T.go:1*/ y84IW11J.Parent()
	}
	/*line P59NXW9kTEn.go:1*/ lffjFkFU5.init()
	return lffjFkFU5, nil
}

func (lffjFkFU5 *bJzmW7YOgIxt) init() {

	var m8Y7Dc = /*line Cbz7GVa.go:1*/ make(map[common.Hash]int)

	for fSpMhzHR := 0; fSpMhzHR < /*line pImzZbKg.go:1*/ len(lffjFkFU5.qR2G2CL); fSpMhzHR++ {

		cu8ZKpmK := lffjFkFU5.qR2G2CL[fSpMhzHR]
		for {

			if ! /*line YhbY_S.go:1*/ cu8ZKpmK.mZR9T1H3Tp65.Next() {
				/*line ua24Kt0.go:1*/ cu8ZKpmK.mZR9T1H3Tp65.Release()
				h6VFlt6g := /*line vNr77zB9s8.go:1*/ len(lffjFkFU5.qR2G2CL) - 1

				lffjFkFU5.qR2G2CL[fSpMhzHR] = lffjFkFU5.qR2G2CL[h6VFlt6g]
				lffjFkFU5.qR2G2CL[h6VFlt6g] = nil
				lffjFkFU5.qR2G2CL = lffjFkFU5.qR2G2CL[:h6VFlt6g]

				fSpMhzHR--
				break
			}

			xhOzkRrKDZ := /*line FhSMwIU.go:1*/ cu8ZKpmK.mZR9T1H3Tp65.Hash()
			if lzcwPxiAHf, h2sw6vA := m8Y7Dc[xhOzkRrKDZ]; !h2sw6vA {
				m8Y7Dc[xhOzkRrKDZ] = fSpMhzHR
				break
			} else {

				if lffjFkFU5.qR2G2CL[lzcwPxiAHf].eyJQwba < cu8ZKpmK.eyJQwba {

					continue
				} else {

					cu8ZKpmK = lffjFkFU5.qR2G2CL[lzcwPxiAHf]
					lffjFkFU5.qR2G2CL[lzcwPxiAHf], lffjFkFU5.qR2G2CL[fSpMhzHR] = lffjFkFU5.qR2G2CL[fSpMhzHR], lffjFkFU5.qR2G2CL[lzcwPxiAHf]
					continue
				}
			}
		}
	}

	/*line UK0SgdYo3.go:1*/
	slices.SortFunc(lffjFkFU5.qR2G2CL, func(rmHPuz8jT, zfco8dsA *xlNMI97NK_) int { return /*line HISsitPbxz.go:1*/ rmHPuz8jT.Cmp(zfco8dsA) })
	lffjFkFU5.r4DzlIJJPs = false
}

func (lffjFkFU5 *bJzmW7YOgIxt) Next() bool {
	if /*line PIWPy2acGMJ.go:1*/ len(lffjFkFU5.qR2G2CL) == 0 {
		return false
	}
	if !lffjFkFU5.r4DzlIJJPs {

		lffjFkFU5.r4DzlIJJPs = true
		if lffjFkFU5._EgzAfKfzWaQ {
			lffjFkFU5.nrTwyBJ = /*line oTdta9iD.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.(Nq4YsH_).Account()
		} else {
			lffjFkFU5.yxDFcc_lELc5 = /*line J46Oasml.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.(PE_7UyJghxy).Slot()
		}
		if vATHS0Ybywa_ := /*line QFaQmHt2W4.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.Error(); vATHS0Ybywa_ != nil {
			lffjFkFU5.zJG9at0m = vATHS0Ybywa_
			return false
		}
		if lffjFkFU5.nrTwyBJ != nil || lffjFkFU5.yxDFcc_lELc5 != nil {
			return true
		}

	}

	for {
		if ! /*line I1qm8nJh.go:1*/ lffjFkFU5.zYawlDUNOaHH(0) {
			return false
		}
		if lffjFkFU5._EgzAfKfzWaQ {
			lffjFkFU5.nrTwyBJ = /*line GLCPivL.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.(Nq4YsH_).Account()
		} else {
			lffjFkFU5.yxDFcc_lELc5 = /*line ycPsaM.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.(PE_7UyJghxy).Slot()
		}
		if vATHS0Ybywa_ := /*line LqyPgW.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.Error(); vATHS0Ybywa_ != nil {
			lffjFkFU5.zJG9at0m = vATHS0Ybywa_
			return false
		}
		if lffjFkFU5.nrTwyBJ != nil || lffjFkFU5.yxDFcc_lELc5 != nil {
			break
		}
	}
	return true
}

func (lffjFkFU5 *bJzmW7YOgIxt) zYawlDUNOaHH(uMgqUt int) bool {

	if cu8ZKpmK := lffjFkFU5.qR2G2CL[uMgqUt].mZR9T1H3Tp65; ! /*line JPguDD.go:1*/ cu8ZKpmK.Next() {
		/*line CzraThA.go:1*/ cu8ZKpmK.Release()

		lffjFkFU5.qR2G2CL = /*line o_CN5M.go:1*/ append(lffjFkFU5.qR2G2CL[:uMgqUt], lffjFkFU5.qR2G2CL[uMgqUt+1:]...)
		return /*line B3gYiyVuYua.go:1*/ len(lffjFkFU5.qR2G2CL) > 0
	}

	if uMgqUt == /*line yLVHoa.go:1*/ len(lffjFkFU5.qR2G2CL)-1 {
		return true
	}

	var (
		ciCjABysa, zYawlDUNOaHH    = lffjFkFU5.qR2G2CL[uMgqUt], lffjFkFU5.qR2G2CL[uMgqUt+1]
		uQW3MXk_v9Od, rUAf7rXWlcQn = /*line Pn6QsPkFay8.go:1*/ ciCjABysa.mZR9T1H3Tp65.Hash() /*line DfjwUfL0J7iL.go:1*/, zYawlDUNOaHH.mZR9T1H3Tp65.Hash()
	)
	if wMCMWH9Tfvzo := /*line RIEP2_PLkp3.go:1*/ bytes.Compare(uQW3MXk_v9Od[:], rUAf7rXWlcQn[:]); wMCMWH9Tfvzo < 0 {

		return true
	} else if wMCMWH9Tfvzo == 0 && ciCjABysa.eyJQwba < zYawlDUNOaHH.eyJQwba {

		/*line BQmLOLTbaTed.go:1*/
		lffjFkFU5.zYawlDUNOaHH(uMgqUt + 1)
		return true
	}

	kSMxyXqcZ1U := -1
	as3oS5CqGh := /*line shp73BQEKf.go:1*/ sort.Search( /*line hpltZD.go:1*/ len(lffjFkFU5.qR2G2CL), func(cZL9UJradA7 int) bool {

		if cZL9UJradA7 < uMgqUt {
			return false
		}
		if cZL9UJradA7 == /*line F7hHoDxXFko.go:1*/ len(lffjFkFU5.qR2G2CL)-1 {

			return true
		}
		rUAf7rXWlcQn := /*line CMzEfe.go:1*/ lffjFkFU5.qR2G2CL[cZL9UJradA7+1].mZR9T1H3Tp65.Hash()
		if wMCMWH9Tfvzo := /*line CqYdglmIv.go:1*/ bytes.Compare(uQW3MXk_v9Od[:], rUAf7rXWlcQn[:]); wMCMWH9Tfvzo < 0 {
			return true
		} else if wMCMWH9Tfvzo > 0 {
			return false
		}

		kSMxyXqcZ1U = cZL9UJradA7 + 1

		return ciCjABysa.eyJQwba < lffjFkFU5.qR2G2CL[cZL9UJradA7+1].eyJQwba
	})
	/*line lgUMeXyKwH.go:1*/ lffjFkFU5.glWqeJ(uMgqUt, as3oS5CqGh)
	if kSMxyXqcZ1U != -1 {
		/*line Fjq9J4H.go:1*/ lffjFkFU5.zYawlDUNOaHH(kSMxyXqcZ1U)
	}
	return true
}

func (lffjFkFU5 *bJzmW7YOgIxt) glWqeJ(as3oS5CqGh, omsSScD int) {
	jZ1tXHzelGns := lffjFkFU5.qR2G2CL[as3oS5CqGh]
	/*line pUyNVaUc.go:1*/ copy(lffjFkFU5.qR2G2CL[as3oS5CqGh:], lffjFkFU5.qR2G2CL[as3oS5CqGh+1:omsSScD+1])
	lffjFkFU5.qR2G2CL[omsSScD] = jZ1tXHzelGns
}

func (lffjFkFU5 *bJzmW7YOgIxt) Error() error {
	return lffjFkFU5.zJG9at0m
}

func (lffjFkFU5 *bJzmW7YOgIxt) Hash() common.Hash {
	return /*line KRRlIU.go:1*/ lffjFkFU5.qR2G2CL[0].mZR9T1H3Tp65.Hash()
}

func (lffjFkFU5 *bJzmW7YOgIxt) Account() []byte {
	return lffjFkFU5.nrTwyBJ
}

func (lffjFkFU5 *bJzmW7YOgIxt) Slot() []byte {
	return lffjFkFU5.yxDFcc_lELc5
}

func (lffjFkFU5 *bJzmW7YOgIxt) Release() {
	for _, cu8ZKpmK := range lffjFkFU5.qR2G2CL {
		/*line ISgUvwgn.go:1*/ cu8ZKpmK.mZR9T1H3Tp65.Release()
	}
	lffjFkFU5.qR2G2CL = nil
}

func (lffjFkFU5 *bJzmW7YOgIxt) Debug() {
	for _, cu8ZKpmK := range lffjFkFU5.qR2G2CL {
		/*line FkujouQ.go:1*/ fmt.Printf(func() /*line HapQfvnO.go:1*/ string {
			data := /*line HapQfvnO.go:1*/ make([]byte, 0, 13)
			i := 3
			decryptKey := 215
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					i = 5
					data = /*line HapQfvnO.go:1*/ append(data, "\xfd\xa8\xe2\xf5"...,
					)
				case 4:
					data = /*line HapQfvnO.go:1*/ append(data, "\xa9\xe7\xfe\xaa"...,
					)
					i = 2
				case 3:
					i = 4
					data = /*line HapQfvnO.go:1*/ append(data, 131)
				case 0:
					i = 1
					for y := range data {
						data[y] = data[y] ^ /*line HapQfvnO.go:1*/ byte(decryptKey^y)
					}
				case 5:
					data = /*line HapQfvnO.go:1*/ append(data, "\xa7\x8f\xf3"...,
					)
					i = 0
				}
			}
			return /*line HapQfvnO.go:1*/ string(data)
		}(), cu8ZKpmK.eyJQwba /*line kXg8llcWc7P.go:1*/, cu8ZKpmK.mZR9T1H3Tp65.Hash()[0])
	}
	/*line pQZsSVnY.go:1*/ fmt.Println()
}

func qEhLP2hRX(abbWTExt *CE2m1DUB6wW, z1BBTN2UX common.Hash, gnGBaeoLX common.Hash) (Nq4YsH_, error) {
	return /*line Y08_aEmYa8j9.go:1*/ l8VM92(abbWTExt, z1BBTN2UX, common.Hash{}, gnGBaeoLX, true)
}

func h4wrLvWBbqCm(abbWTExt *CE2m1DUB6wW, z1BBTN2UX common.Hash, evzmhL1 common.Hash, gnGBaeoLX common.Hash) (PE_7UyJghxy, error) {
	return /*line n_dVST6oPxg.go:1*/ l8VM92(abbWTExt, z1BBTN2UX, evzmhL1, gnGBaeoLX, false)
}

var _ bytes.Buffer
var _ = fmt.Append
var _ = sort.Find
var _ = common.AbsolutePath
