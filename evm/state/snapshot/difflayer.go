//line :1
package snapshot

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	bloomfilter "github.com/holiman/bloomfilter/v2"
	"golang.org/x/exp/slices"
)

var (
	axaqRIVTe_ef = /*line rshH4gMbJTU.go:1*/ uint64(4 * 1024 * 1024)

	jHWcEsAoCb = axaqRIVTe_ef / 42

	a85CXb8JknE = 0.02

	wM7IOjTIcf0b = /*line dXc1NG.go:1*/ math.Ceil( /*line HwDqr6I.go:1*/ float64(jHWcEsAoCb) * /*line tzchezFOCR92.go:1*/ math.Log(a85CXb8JknE) / /*line xLu0bURXE.go:1*/ math.Log(1/ /*line qBuc1JSOJIZe.go:1*/ math.Pow(2 /*line kvxFd31amF.go:1*/, math.Log(2))))

	azYtXORCNVk = /*line QbuLKapN7z1S.go:1*/ math.Round((wM7IOjTIcf0b / /*line G9BjJqw4.go:1*/ float64(jHWcEsAoCb)) * /*line HAi7i78uSiNy.go:1*/ math.Log(2))

	x5_wwJDKDlnN = 0
	tjIvwWThZSb  = 0
	g4XjSTnfx    = 0
)

func init() {

	x5_wwJDKDlnN = /*line IVyyGZckn8rH.go:1*/ rand.Intn(25)
	tjIvwWThZSb = /*line BMNc1DoEKQqt.go:1*/ rand.Intn(25)
	g4XjSTnfx = /*line aMo8sMJn.go:1*/ rand.Intn(25)

	for tjIvwWThZSb == x5_wwJDKDlnN {
		tjIvwWThZSb = /*line Vpg8Vt.go:1*/ rand.Intn(25)
	}
}

type diffLayer struct {
	origin *diskLayer
	parent snapshot
	memory uint64

	root  common.Hash
	stale atomic.Bool

	destructSet map[common.Hash]struct{}
	accountList []common.Hash
	accountData map[common.Hash][]byte
	storageList map[common.Hash][]common.Hash
	storageData map[common.Hash]map[common.Hash][]byte

	diffed *bloomfilter.Filter

	lock sync.RWMutex
}

func w7uXLNVT(dNQitYj3 common.Hash) uint64 {
	return /*line A2OZGpo2.go:1*/ binary.BigEndian.Uint64(dNQitYj3[x5_wwJDKDlnN : x5_wwJDKDlnN+8])
}

func g_vvx4(dNQitYj3 common.Hash) uint64 {
	return /*line H85Dcakz.go:1*/ binary.BigEndian.Uint64(dNQitYj3[tjIvwWThZSb : tjIvwWThZSb+8])
}

func iDkGlpfP_(ea7merFK, hBrNIzbDg8 common.Hash) uint64 {
	return /*line lgamLaS.go:1*/ binary.BigEndian.Uint64(ea7merFK[g4XjSTnfx:g4XjSTnfx+8]) ^
		/*line Go_iPivbFNS4.go:1*/ binary.BigEndian.Uint64(hBrNIzbDg8[g4XjSTnfx:g4XjSTnfx+8])
}

func aa7YtGB_s(f3dWTs7 snapshot, z1BBTN2UX common.Hash, zoMmGkL map[common.Hash]struct{}, bQIIyfhppAL1 map[common.Hash][]byte, agSpCMzc map[common.Hash]map[common.Hash][]byte) *diffLayer {

	hbsnnIi := &diffLayer{
		parent:      f3dWTs7,
		root:        z1BBTN2UX,
		destructSet: zoMmGkL,
		accountData: bQIIyfhppAL1,
		storageData: agSpCMzc,
		storageList:/*line V9XTgNexJ.go:1*/ make(map[common.Hash][]common.Hash),
	}
	switch f3dWTs7 := f3dWTs7.(type) {
	case *diskLayer:
		/*line DtsNqg2.go:1*/ hbsnnIi.a4OFrKzR(f3dWTs7)
	case *diffLayer:
		/*line bQ_s_EVeRF.go:1*/ hbsnnIi.a4OFrKzR(f3dWTs7.origin)
	default:
		/*line Hjp3BcECu.go:1*/ panic(func() /*line Xlavn6z4l.go:1*/ string {
			data := [] /*line Xlavn6z4l.go:1*/ byte("s\x16\x0ea\";\xdf\x17par nt \x10\rҷ")
			positions := [...]byte{0, 18, 6, 5, 3, 5, 16, 4, 3, 6, 4, 2, 11, 5, 7, 18, 3, 2, 1, 6, 3, 17, 3, 0, 11, 15}
			for i := 0; i < 26; i += 2 {
				localKey := /*line Xlavn6z4l.go:1*/ byte(i) + /*line Xlavn6z4l.go:1*/ byte(positions[i]^positions[i+1]) + 143
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line Xlavn6z4l.go:1*/ string(data)
		}())
	}

	for dcJdHV, b1jafJL := range bQIIyfhppAL1 {
		if b1jafJL == nil {
			/*line GVU_fB6X7mIj.go:1*/ panic( /*line gVI4OGL.go:1*/ fmt.Sprintf(func() /*line d7obtHjyw.go:1*/ string {
				fullData := [] /*line d7obtHjyw.go:1*/ byte("HXͥ\x03'\xfc\xdd\x1dpU5\xcbu-\xc1\x8e\x0f|\xe0\x00\x0eoe$\f9?\xf4\xe7")
				idxKey := [] /*line d7obtHjyw.go:1*/ byte("\xa2\x18\x95\xa6\x91N\xca\xfeoB\xb4\xab\xa5")
				data := /*line d7obtHjyw.go:1*/ make([]byte, 0, 16)
				data = /*line d7obtHjyw.go:1*/ append(data, fullData[247^ /*line d7obtHjyw.go:1*/ int(idxKey[7])]-fullData[239^ /*line d7obtHjyw.go:1*/ int(idxKey[7])], fullData[215^ /*line d7obtHjyw.go:1*/ int(idxKey[6])]+fullData[216^ /*line d7obtHjyw.go:1*/ int(idxKey[6])], fullData[14^ /*line d7obtHjyw.go:1*/ int(idxKey[1])]+fullData[4^ /*line d7obtHjyw.go:1*/ int(idxKey[1])], fullData[174^ /*line d7obtHjyw.go:1*/ int(idxKey[11])]+fullData[171^ /*line d7obtHjyw.go:1*/ int(idxKey[11])], fullData[168^ /*line d7obtHjyw.go:1*/ int(idxKey[12])]+fullData[177^ /*line d7obtHjyw.go:1*/ int(idxKey[12])], fullData[182^ /*line d7obtHjyw.go:1*/ int(idxKey[3])]+fullData[181^ /*line d7obtHjyw.go:1*/ int(idxKey[3])], fullData[3^ /*line d7obtHjyw.go:1*/ int(idxKey[1])]-fullData[20^ /*line d7obtHjyw.go:1*/ int(idxKey[1])], fullData[176^ /*line d7obtHjyw.go:1*/ int(idxKey[10])]+fullData[188^ /*line d7obtHjyw.go:1*/ int(idxKey[10])], fullData[26^ /*line d7obtHjyw.go:1*/ int(idxKey[1])]+fullData[25^ /*line d7obtHjyw.go:1*/ int(idxKey[1])], fullData[165^ /*line d7obtHjyw.go:1*/ int(idxKey[11])]^fullData[190^ /*line d7obtHjyw.go:1*/ int(idxKey[11])], fullData[165^ /*line d7obtHjyw.go:1*/ int(idxKey[3])]^fullData[161^ /*line d7obtHjyw.go:1*/ int(idxKey[3])], fullData[190^ /*line d7obtHjyw.go:1*/ int(idxKey[3])]+fullData[160^ /*line d7obtHjyw.go:1*/ int(idxKey[3])], fullData[228^ /*line d7obtHjyw.go:1*/ int(idxKey[7])]+fullData[245^ /*line d7obtHjyw.go:1*/ int(idxKey[7])], fullData[187^ /*line d7obtHjyw.go:1*/ int(idxKey[0])]^fullData[181^ /*line d7obtHjyw.go:1*/ int(idxKey[0])], fullData[164^ /*line d7obtHjyw.go:1*/ int(idxKey[11])]-fullData[161^ /*line d7obtHjyw.go:1*/ int(idxKey[11])])
				return /*line d7obtHjyw.go:1*/ string(data)
			}(), dcJdHV))
		}

		hbsnnIi.memory += /*line CkVyqfv6G.go:1*/ uint64(common.HashLength + /*line BU2T_Pc5u.go:1*/ len(b1jafJL))
		/*line zjvGc6CGD0.go:1*/ wranHZ1.Mark( /*line VTnh__.go:1*/ int64( /*line UeXWLWID.go:1*/ len(b1jafJL)))
	}
	for dcJdHV, etj53bd3zM := range agSpCMzc {
		if etj53bd3zM == nil {
			/*line SOHGBeW.go:1*/ panic( /*line a_tfo0m.go:1*/ fmt.Sprintf(func() /*line oVahuhcaP.go:1*/ string {
				data := [] /*line oVahuhcaP.go:1*/ byte("s\xc5Ar\x8c\xc8\xd1\xc6%#\xb7\xbd2\xd4 ")
				positions := [...]byte{2, 4, 12, 6, 14, 2, 4, 1, 7, 13, 6, 11, 7, 11, 2, 14, 10, 5}
				for i := 0; i < 18; i += 2 {
					localKey := /*line oVahuhcaP.go:1*/ byte(i) + /*line oVahuhcaP.go:1*/ byte(positions[i]^positions[i+1]) + 145
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
				}
				return /*line oVahuhcaP.go:1*/ string(data)
			}(), dcJdHV))
		}

		for _, bI1ciqVUvL4b := range etj53bd3zM {
			hbsnnIi.memory += /*line M2dI4HMqy.go:1*/ uint64(common.HashLength + /*line MD4vWlEv0.go:1*/ len(bI1ciqVUvL4b))
			/*line vOCCt7E.go:1*/ b2Hf0xz9.Mark( /*line fHbqKaaRX.go:1*/ int64( /*line cj5CKxsX7e.go:1*/ len(bI1ciqVUvL4b)))
		}
	}
	hbsnnIi.memory += /*line GkJoJNoxi.go:1*/ uint64( /*line DVNnCrfD.go:1*/ len(zoMmGkL) * common.HashLength)
	return hbsnnIi
}

func (hbsnnIi *diffLayer) a4OFrKzR(ijcATHiDv *diskLayer) {
	/*line BLv73z5B.go:1*/ hbsnnIi.lock.Lock()
	defer /*line rjWLKJziRvb.go:1*/ hbsnnIi.lock.Unlock()

	defer func( /*line JyV8jvTu4IW.go:1*/ lFRN2MXQc time.Time) {
		/*line Fbb8hqIKCmXS.go:1*/ kC2ik5HlylJ.Update( /*line VP3MaIcJa.go:1*/ time.Since(lFRN2MXQc))
	}( /*line fxyjt0.go:1*/ time.Now())

	hbsnnIi.origin = ijcATHiDv

	if f3dWTs7, jdkNTRyBJ := hbsnnIi.parent.(*diffLayer); jdkNTRyBJ {
		/*line vf6bHqY.go:1*/ f3dWTs7.lock.RLock()
		hbsnnIi.diffed, _ = /*line USwioceUY.go:1*/ f3dWTs7.diffed.Copy()
		/*line IHjvdl_C8_W5.go:1*/ f3dWTs7.lock.RUnlock()
	} else {
		hbsnnIi.diffed, _ = /*line aPbcJ4j.go:1*/ bloomfilter.New( /*line fy1mkd_KTag.go:1*/ uint64(wM7IOjTIcf0b) /*line WRZPAsUqjJO.go:1*/, uint64(azYtXORCNVk))
	}

	for xhOzkRrKDZ := range hbsnnIi.destructSet {
		/*line eBN2p87bUAN.go:1*/ hbsnnIi.diffed.AddHash( /*line gVnQwi.go:1*/ w7uXLNVT(xhOzkRrKDZ))
	}
	for xhOzkRrKDZ := range hbsnnIi.accountData {
		/*line MIYLIBGracAI.go:1*/ hbsnnIi.diffed.AddHash( /*line vVGI_mX.go:1*/ g_vvx4(xhOzkRrKDZ))
	}
	for dcJdHV, etj53bd3zM := range hbsnnIi.storageData {
		for mq_bNE9GL := range etj53bd3zM {
			/*line AAkJE4oOfQI9.go:1*/ hbsnnIi.diffed.AddHash( /*line JZKzg1yPml.go:1*/ iDkGlpfP_(dcJdHV, mq_bNE9GL))
		}
	}

	ycHKuY := /*line tSCjjG4.go:1*/ float64( /*line E68R7Kh4.go:1*/ hbsnnIi.diffed.K())
	cZL9UJradA7 := /*line Z3XN3a.go:1*/ float64( /*line qWohJ3U6R.go:1*/ hbsnnIi.diffed.N())
	dQPD9HQ := /*line TDdCXL9.go:1*/ float64( /*line N0gn9W__Ta4.go:1*/ hbsnnIi.diffed.M())
	/*line DnHoN4ina.go:1*/ sO0JKDQCa20r.Update( /*line G1alPon_A2.go:1*/ math.Pow(1.0- /*line JSKamCmvK9a.go:1*/ math.Exp((-ycHKuY)*(cZL9UJradA7+0.5)/(dQPD9HQ-1)), ycHKuY))
}

func (hbsnnIi *diffLayer) Root() common.Hash {
	return hbsnnIi.root
}

func (hbsnnIi *diffLayer) Parent() snapshot {
	/*line HxMnFS.go:1*/ hbsnnIi.lock.RLock()
	defer /*line Gs706d2o2i5i.go:1*/ hbsnnIi.lock.RUnlock()

	return hbsnnIi.parent
}

func (hbsnnIi *diffLayer) Stale() bool {
	return /*line K8151yNa9B0T.go:1*/ hbsnnIi.stale.Load()
}

func (hbsnnIi *diffLayer) Account(xhOzkRrKDZ common.Hash) (*types.SlimAccount, error) {
	bI1ciqVUvL4b, chyZ8Q := /*line DjSnzga.go:1*/ hbsnnIi.AccountRLP(xhOzkRrKDZ)
	if chyZ8Q != nil {
		return nil, chyZ8Q
	}
	if /*line NgFxrslKb9uE.go:1*/ len(bI1ciqVUvL4b) == 0 {
		return nil, nil
	}
	evzmhL1 := /*line egPeoU.go:1*/ new(types.SlimAccount)
	if chyZ8Q := /*line EaPDs__7GaC.go:1*/ rlp.DecodeBytes(bI1ciqVUvL4b, evzmhL1); chyZ8Q != nil {
		/*line jcRxAyOsoOQz.go:1*/ panic(chyZ8Q)
	}
	return evzmhL1, nil
}

func (hbsnnIi *diffLayer) AccountRLP(xhOzkRrKDZ common.Hash) ([]byte, error) {

	/*line UBf6QDeKNpt.go:1*/
	hbsnnIi.lock.RLock()
	if /*line FRPWi8d69kqs.go:1*/ hbsnnIi.Stale() {
		/*line C0tNr_uQpt5.go:1*/ hbsnnIi.lock.RUnlock()
		return nil, AiK5hkDaW
	}

	hCJXfqGTJQJ := /*line ASvpG7.go:1*/ hbsnnIi.diffed.ContainsHash( /*line QVwckvDw4.go:1*/ g_vvx4(xhOzkRrKDZ))
	if !hCJXfqGTJQJ {
		hCJXfqGTJQJ = /*line p9MCkSlnAfB.go:1*/ hbsnnIi.diffed.ContainsHash( /*line zCCLKYAM1.go:1*/ w7uXLNVT(xhOzkRrKDZ))
	}
	var ijcATHiDv *diskLayer
	if !hCJXfqGTJQJ {
		ijcATHiDv = hbsnnIi.origin
	}
	/*line MAkwwSmvJ.go:1*/ hbsnnIi.lock.RUnlock()

	if ijcATHiDv != nil {
		/*line YkiS_pFA.go:1*/ bqXxiQEUjan.Mark(1)
		return /*line VMpJJH.go:1*/ ijcATHiDv.AccountRLP(xhOzkRrKDZ)
	}

	return /*line GuhAVWr.go:1*/ hbsnnIi.yP_XjO(xhOzkRrKDZ, 0)
}

func (hbsnnIi *diffLayer) yP_XjO(xhOzkRrKDZ common.Hash, tCwDahj07TnV int) ([]byte, error) {
	/*line DsZNIy4N2W.go:1*/ hbsnnIi.lock.RLock()
	defer /*line LrVR1TM4XL8.go:1*/ hbsnnIi.lock.RUnlock()

	if /*line rDXUDDsxo.go:1*/ hbsnnIi.Stale() {
		return nil, AiK5hkDaW
	}

	if bI1ciqVUvL4b, jdkNTRyBJ := hbsnnIi.accountData[xhOzkRrKDZ]; jdkNTRyBJ {
		/*line oJhC1EWT2s.go:1*/ bF2XXz.Mark(1)
		/*line BG3kDngJbu.go:1*/ aiGEFccO.Update( /*line Ygft96.go:1*/ int64(tCwDahj07TnV))
		/*line AQ6LGKeWZXeS.go:1*/ olQwZ3qqpE0E.Mark( /*line XwRw7Ac.go:1*/ int64( /*line lkagDoe.go:1*/ len(bI1ciqVUvL4b)))
		/*line FsG95yRMRotr.go:1*/ s61yEx.Mark(1)
		return bI1ciqVUvL4b, nil
	}

	if _, jdkNTRyBJ := hbsnnIi.destructSet[xhOzkRrKDZ]; jdkNTRyBJ {
		/*line h0kxSIbB.go:1*/ bF2XXz.Mark(1)
		/*line BaGQBKt7iNha.go:1*/ aiGEFccO.Update( /*line a3KysnVk.go:1*/ int64(tCwDahj07TnV))
		/*line RlwRrgSHz2.go:1*/ hsOjrftc9.Mark(1)
		/*line aautoNjz.go:1*/ s61yEx.Mark(1)
		return nil, nil
	}

	if wMCMWH9Tfvzo, jdkNTRyBJ := hbsnnIi.parent.(*diffLayer); jdkNTRyBJ {
		return /*line GzaAynW.go:1*/ wMCMWH9Tfvzo.yP_XjO(xhOzkRrKDZ, tCwDahj07TnV+1)
	}

	/*line k_MUveB_Fwg.go:1*/
	bPVyOyY.Mark(1)
	return /*line J3_wvJF3KpZi.go:1*/ hbsnnIi.parent.AccountRLP(xhOzkRrKDZ)
}

func (hbsnnIi *diffLayer) Storage(dcJdHV, mq_bNE9GL common.Hash) ([]byte, error) {

	/*line woXZmJva6_J7.go:1*/
	hbsnnIi.lock.RLock()

	if /*line Ja3sZR.go:1*/ hbsnnIi.Stale() {
		/*line pK2ZZ6BXbRt1.go:1*/ hbsnnIi.lock.RUnlock()
		return nil, AiK5hkDaW
	}
	hCJXfqGTJQJ := /*line IPXX5Dq.go:1*/ hbsnnIi.diffed.ContainsHash( /*line kvGP6o5cu5.go:1*/ iDkGlpfP_(dcJdHV, mq_bNE9GL))
	if !hCJXfqGTJQJ {
		hCJXfqGTJQJ = /*line FihcZD.go:1*/ hbsnnIi.diffed.ContainsHash( /*line DrIAeue.go:1*/ w7uXLNVT(dcJdHV))
	}
	var ijcATHiDv *diskLayer
	if !hCJXfqGTJQJ {
		ijcATHiDv = hbsnnIi.origin
	}
	/*line OJ85kfoQ1.go:1*/ hbsnnIi.lock.RUnlock()

	if ijcATHiDv != nil {
		/*line BaHlcoRphx.go:1*/ eS34fxUGaE.Mark(1)
		return /*line Qw87awU.go:1*/ ijcATHiDv.Storage(dcJdHV, mq_bNE9GL)
	}

	return /*line GLu2wF.go:1*/ hbsnnIi.agSpCMzc(dcJdHV, mq_bNE9GL, 0)
}

func (hbsnnIi *diffLayer) agSpCMzc(dcJdHV, mq_bNE9GL common.Hash, tCwDahj07TnV int) ([]byte, error) {
	/*line kcfUIPvn7Zi.go:1*/ hbsnnIi.lock.RLock()
	defer /*line _Q3dtNy.go:1*/ hbsnnIi.lock.RUnlock()

	if /*line muAau90DbSzW.go:1*/ hbsnnIi.Stale() {
		return nil, AiK5hkDaW
	}

	if agSpCMzc, jdkNTRyBJ := hbsnnIi.storageData[dcJdHV]; jdkNTRyBJ {
		if bI1ciqVUvL4b, jdkNTRyBJ := agSpCMzc[mq_bNE9GL]; jdkNTRyBJ {
			/*line IZ_GMUpl8gwE.go:1*/ z6kgmtWyiAG.Mark(1)
			/*line zrCVPC_Q.go:1*/ w_kwXL__A.Update( /*line g8ka3s.go:1*/ int64(tCwDahj07TnV))
			if cZL9UJradA7 := /*line FPh20FYQLAJ.go:1*/ len(bI1ciqVUvL4b); cZL9UJradA7 > 0 {
				/*line H5FNdSSKeZb.go:1*/ fgr54pEYI.Mark( /*line BjXdaW.go:1*/ int64(cZL9UJradA7))
			} else {
				/*line aHSJyU.go:1*/ wn3S6_X.Mark(1)
			}
			/*line nok9svPagf.go:1*/ iSLu5zV9.Mark(1)
			return bI1ciqVUvL4b, nil
		}
	}

	if _, jdkNTRyBJ := hbsnnIi.destructSet[dcJdHV]; jdkNTRyBJ {
		/*line V_8FQ3z7W4W.go:1*/ z6kgmtWyiAG.Mark(1)
		/*line jCrH_mJHUVh.go:1*/ w_kwXL__A.Update( /*line jr0JuxgLM3.go:1*/ int64(tCwDahj07TnV))
		/*line a7x2QroH.go:1*/ wn3S6_X.Mark(1)
		/*line NRpFBt.go:1*/ iSLu5zV9.Mark(1)
		return nil, nil
	}

	if wMCMWH9Tfvzo, jdkNTRyBJ := hbsnnIi.parent.(*diffLayer); jdkNTRyBJ {
		return /*line kdaDgXUjLcZU.go:1*/ wMCMWH9Tfvzo.agSpCMzc(dcJdHV, mq_bNE9GL, tCwDahj07TnV+1)
	}

	/*line Cwg3plxZP_fw.go:1*/
	jy1zd1hC0As.Mark(1)
	return /*line rtvcBh0wg_3.go:1*/ hbsnnIi.parent.Storage(dcJdHV, mq_bNE9GL)
}

func (hbsnnIi *diffLayer) Update(wWDb6cW common.Hash, zoMmGkL map[common.Hash]struct{}, bQIIyfhppAL1 map[common.Hash][]byte, agSpCMzc map[common.Hash]map[common.Hash][]byte) *diffLayer {
	return /*line yvKB7eZcv.go:1*/ aa7YtGB_s(hbsnnIi, wWDb6cW, zoMmGkL, bQIIyfhppAL1, agSpCMzc)
}

func (hbsnnIi *diffLayer) pcFKD7() snapshot {

	f3dWTs7, jdkNTRyBJ := hbsnnIi.parent.(*diffLayer)
	if !jdkNTRyBJ {
		return hbsnnIi
	}

	f3dWTs7 = /*line pSCaI2p.go:1*/ f3dWTs7.pcFKD7().(*diffLayer)

	/*line R4xJfA0iBoLO.go:1*/
	f3dWTs7.lock.Lock()
	defer /*line uR0SiawbPpZ6.go:1*/ f3dWTs7.lock.Unlock()

	if /*line c5TLe2Ma_Wx1.go:1*/ f3dWTs7.stale.Swap(true) {
		/*line JYvpMi.go:1*/ panic(func() /*line Jk4dl0bqw8.go:1*/ string {
			seed := /*line Jk4dl0bqw8.go:1*/ byte(35)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line Jk4dl0bqw8.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line Jk4dl0bqw8.go:1*/ fnc(147)(23)(63)(113)(235)(220)(100)(12)(29)(55)(110)(150)(120)(229)(226)(176)(109)(136)(89)(188)(37)(157)(59)(99)(209)(155)
			return /*line Jk4dl0bqw8.go:1*/ string(data)
		}())
	}

	for xhOzkRrKDZ := range hbsnnIi.destructSet {
		f3dWTs7.destructSet[xhOzkRrKDZ] = struct{}{}
		/*line GAoWjYcqy.go:1*/ delete(f3dWTs7.accountData, xhOzkRrKDZ)
		/*line Pe5ClhsP.go:1*/ delete(f3dWTs7.storageData, xhOzkRrKDZ)
	}
	for xhOzkRrKDZ, bI1ciqVUvL4b := range hbsnnIi.accountData {
		f3dWTs7.accountData[xhOzkRrKDZ] = bI1ciqVUvL4b
	}

	for dcJdHV, agSpCMzc := range hbsnnIi.storageData {

		if _, jdkNTRyBJ := f3dWTs7.storageData[dcJdHV]; !jdkNTRyBJ {
			f3dWTs7.storageData[dcJdHV] = agSpCMzc
			continue
		}

		f3_D0laC := f3dWTs7.storageData[dcJdHV]
		for mq_bNE9GL, bI1ciqVUvL4b := range agSpCMzc {
			f3_D0laC[mq_bNE9GL] = bI1ciqVUvL4b
		}
	}

	return &diffLayer{
		parent:      f3dWTs7.parent,
		origin:      f3dWTs7.origin,
		root:        hbsnnIi.root,
		destructSet: f3dWTs7.destructSet,
		accountData: f3dWTs7.accountData,
		storageData: f3dWTs7.storageData,
		storageList:/*line BbIR9cdk.go:1*/ make(map[common.Hash][]common.Hash),
		diffed: hbsnnIi.diffed,
		memory: f3dWTs7.memory + hbsnnIi.memory,
	}
}

func (hbsnnIi *diffLayer) AccountList() []common.Hash {

	/*line HKyHURirLO.go:1*/
	hbsnnIi.lock.RLock()
	qn9DSiO := hbsnnIi.accountList
	/*line CYs7v_OM.go:1*/ hbsnnIi.lock.RUnlock()

	if qn9DSiO != nil {
		return qn9DSiO
	}

	/*line HaZDe_B.go:1*/
	hbsnnIi.lock.Lock()
	defer /*line Bfaw4KQ9Uw.go:1*/ hbsnnIi.lock.Unlock()

	hbsnnIi.accountList = /*line FPoMdao.go:1*/ make([]common.Hash, 0 /*line JNPoJKU.go:1*/, len(hbsnnIi.destructSet)+ /*line b5iaaC3.go:1*/ len(hbsnnIi.accountData))
	for xhOzkRrKDZ := range hbsnnIi.accountData {
		hbsnnIi.accountList = /*line Vm83QT.go:1*/ append(hbsnnIi.accountList, xhOzkRrKDZ)
	}
	for xhOzkRrKDZ := range hbsnnIi.destructSet {
		if _, jdkNTRyBJ := hbsnnIi.accountData[xhOzkRrKDZ]; !jdkNTRyBJ {
			hbsnnIi.accountList = /*line uRX8TEZrzRKr.go:1*/ append(hbsnnIi.accountList, xhOzkRrKDZ)
		}
	}
	/*line tkHy4U.go:1*/ slices.SortFunc(hbsnnIi.accountList, common.Hash.Cmp)
	hbsnnIi.memory += /*line HeYN8R.go:1*/ uint64( /*line L9gz1O2L.go:1*/ len(hbsnnIi.accountList) * common.HashLength)
	return hbsnnIi.accountList
}

func (hbsnnIi *diffLayer) StorageList(dcJdHV common.Hash) ([]common.Hash, bool) {
	/*line yi5VLBb0HVIO.go:1*/ hbsnnIi.lock.RLock()
	_, oiHhcQbh9XAf := hbsnnIi.destructSet[dcJdHV]
	if _, jdkNTRyBJ := hbsnnIi.storageData[dcJdHV]; !jdkNTRyBJ {

		/*line VbaYU5LX.go:1*/
		hbsnnIi.lock.RUnlock()
		return nil, oiHhcQbh9XAf
	}

	if qn9DSiO, h2sw6vA := hbsnnIi.storageList[dcJdHV]; h2sw6vA {
		/*line NfQ9K6np.go:1*/ hbsnnIi.lock.RUnlock()
		return qn9DSiO, oiHhcQbh9XAf
	}
	/*line O0Ita08f.go:1*/ hbsnnIi.lock.RUnlock()

	/*line uMD7sr.go:1*/
	hbsnnIi.lock.Lock()
	defer /*line OoZqJ8CGROv.go:1*/ hbsnnIi.lock.Unlock()

	dxWiS43oVr := hbsnnIi.storageData[dcJdHV]
	aR4qGPdr1lU := /*line NDl9lMbG.go:1*/ make([]common.Hash, 0 /*line XP76vQxQckn.go:1*/, len(dxWiS43oVr))
	for ycHKuY := range dxWiS43oVr {
		aR4qGPdr1lU = /*line sLlftr.go:1*/ append(aR4qGPdr1lU, ycHKuY)
	}
	/*line cQvPE9quycLi.go:1*/ slices.SortFunc(aR4qGPdr1lU, common.Hash.Cmp)
	hbsnnIi.storageList[dcJdHV] = aR4qGPdr1lU
	hbsnnIi.memory += /*line EDPO7eUyPuM.go:1*/ uint64( /*line BaFvW6HyN.go:1*/ len(hbsnnIi.storageList)*common.HashLength + common.HashLength)
	return aR4qGPdr1lU, oiHhcQbh9XAf
}

var _ = binary.Append
var _ = fmt.Append
var _ = math.Abs
var _ = rand.ExpFloat64
var _ sync.Cond
var _ = atomic.AddInt32

const _ = time.ANSIC

var _ types.AccessList
var _ = common.AbsolutePath
var _ = rlp.AppendUint64
var _ = bloomfilter.CountBitsUint64s
