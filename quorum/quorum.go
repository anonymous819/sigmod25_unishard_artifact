//line :1
package quorum

import (
	"fmt"
	"sync"
	crypto "unishard/crypto"
	log "unishard/log"
	types "unishard/types"

	"github.com/ethereum/go-ethereum/common"
)

type (
	Commit byte
	Vote   byte
)

type sQXzaEG98Fq interface {
	Vote | Commit
}

type Collection[HSGwUUJCZ sQXzaEG98Fq] struct {
	types.Epoch
	types.View
	types.BlockHeight
	WpsY_aZ    types.NodeID
	NaNYri773M common.Hash
	crypto.MqlfmVE9d
}
type FZT207R[HSGwUUJCZ sQXzaEG98Fq] struct {
	oPhadLFZ6K7o int
	lDwPGMKP3u_x map[common.Hash]map[types.NodeID]*Collection[HSGwUUJCZ]
	aKxm6PV0a    sync.Mutex
}

type HPOWa1PT0xzq struct {
	AVzGLiX4RWH types.NodeID
	EmFrce      types.Epoch
	LwVL87      types.View
	types.BlockHeight
	ZWsU_2     common.Hash
	X5i3Ynndjb []types.NodeID
	i06mmUA    int
	crypto.Pp__49cd
	crypto.MqlfmVE9d
	jukqzd sync.Mutex
}

func Dq3TFZ[HSGwUUJCZ sQXzaEG98Fq](iJD83L int) *FZT207R[HSGwUUJCZ] {
	return &FZT207R[HSGwUUJCZ]{
		oPhadLFZ6K7o: iJD83L,
		lDwPGMKP3u_x:/*line jaHXT68vrG48.go:1*/ make(map[common.Hash]map[types.NodeID]*Collection[HSGwUUJCZ]),
	}
}

func G71jC_Q[HSGwUUJCZ sQXzaEG98Fq](ldX6M2dH2_c types.Epoch, bkM3WVojf1 types.View, uEugM5qRU types.BlockHeight, qpaKPuRsfb9 types.NodeID, s6k5uN0 common.Hash) *Collection[HSGwUUJCZ] {
	wZjgkbyaa71, nf3nsQ := /*line CBH9KGZeX.go:1*/ crypto.PrivSign( /*line Cxs1FF4ma5Zi.go:1*/ crypto.IDToByte(s6k5uN0), nil)
	if nf3nsQ != nil {
		/*line QESViamst4o.go:1*/ log.ZD1I5u7HLF(func() /*line YfUM_9yl4.go:1*/ string {
			seed := /*line YfUM_9yl4.go:1*/ byte(40)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line YfUM_9yl4.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line YfUM_9yl4.go:1*/ fnc(131)(208)(241)(201)(85)(242)(221)(204)(69)(203)(163)(248)(53)(119)(238)(217)(181)(24)(135)(255)(251)(255)(176)(179)(92)(182)(115)(225)(199)(135)(199)(207)(93)(16)(25)(55)(95)
			return /*line YfUM_9yl4.go:1*/ string(data)
		}(), qpaKPuRsfb9)
		return nil
	}
	return &Collection[HSGwUUJCZ]{
		Epoch:       ldX6M2dH2_c,
		View:        bkM3WVojf1,
		BlockHeight: uEugM5qRU,
		WpsY_aZ:     qpaKPuRsfb9,
		NaNYri773M:  s6k5uN0,
		MqlfmVE9d:   wZjgkbyaa71,
	}
}

func (gSe6Sbe *FZT207R[HSGwUUJCZ]) Add(wLUxkeF *Collection[HSGwUUJCZ]) (bool, *HPOWa1PT0xzq) {
	/*line ozp4n88F0.go:1*/ gSe6Sbe.aKxm6PV0a.Lock()
	defer /*line x0nlQ7O_.go:1*/ gSe6Sbe.aKxm6PV0a.Unlock()

	if /*line ZD4Th5.go:1*/ gSe6Sbe.lgWsSv(wLUxkeF.NaNYri773M) {
		return false, nil
	}
	_, eJfDOtOsF := gSe6Sbe.lDwPGMKP3u_x[wLUxkeF.NaNYri773M]
	if !eJfDOtOsF {

		gSe6Sbe.lDwPGMKP3u_x[wLUxkeF.NaNYri773M] = /*line R9D12dciNvt.go:1*/ make(map[types.NodeID]*Collection[HSGwUUJCZ])
	}
	gSe6Sbe.lDwPGMKP3u_x[wLUxkeF.NaNYri773M][wLUxkeF.WpsY_aZ] = wLUxkeF
	if /*line rRxlWgeKq5.go:1*/ gSe6Sbe.lgWsSv(wLUxkeF.NaNYri773M) {
		w1Jnuao8Vaf, aZJ5AQFhNPiv, nf3nsQ := /*line I1fdeP0k.go:1*/ gSe6Sbe.yqG1CuIwq(wLUxkeF.NaNYri773M)
		if nf3nsQ != nil {
			/*line r28JujHKJVzT.go:1*/ log.WPP4KjwN(func() /*line F2zSL6.go:1*/ string {
				seed := /*line F2zSL6.go:1*/ byte(46)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line F2zSL6.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line F2zSL6.go:1*/ fnc(6)(98)(249)(251)(239)(84)(140)(61)(242)(161)(9)(81)(226)(11)(30)(225)(27)(170)(83)(226)(7)(21)(247)(29)(237)(227)(73)(211)(165)(92)(231)(1)(7)(17)(166)(93)(234)(83)(164)(6)(31)(236)(16)(227)(11)(16)(238)(31)(226)(88)(245)(179)(88)(166)(31)(240)(242)(87)(235)(207)(168)(85)(245)(21)(236)(19)(174)(25)(35)(88)(178)(238)(31)(236)(16)(171)(95)(241)(166)(9)(77)(184)(26)(113)(179)
				return /*line F2zSL6.go:1*/ string(data)
			}(), wLUxkeF.BlockHeight, wLUxkeF.View, wLUxkeF.Epoch, wLUxkeF.NaNYri773M, nf3nsQ)
		}
		xTLMjamNN0 := &HPOWa1PT0xzq{
			EmFrce:      wLUxkeF.Epoch,
			LwVL87:      wLUxkeF.View,
			BlockHeight: wLUxkeF.BlockHeight,
			ZWsU_2:      wLUxkeF.NaNYri773M,
			Pp__49cd:    w1Jnuao8Vaf,
			X5i3Ynndjb:  aZJ5AQFhNPiv,
		}
		return true, xTLMjamNN0
	}
	return false, nil
}

func (gSe6Sbe *FZT207R[HSGwUUJCZ]) lgWsSv(s6k5uN0 common.Hash) bool {

	return /*line I2_TvcO.go:1*/ gSe6Sbe.iJnqIt(s6k5uN0) > gSe6Sbe.oPhadLFZ6K7o*2/3
}

func (gSe6Sbe *FZT207R[HSGwUUJCZ]) Majority(s6k5uN0 common.Hash) bool {

	return /*line TDCx6Iv.go:1*/ gSe6Sbe.iJnqIt(s6k5uN0) > gSe6Sbe.oPhadLFZ6K7o*1/2
}

func (gSe6Sbe *FZT207R[HSGwUUJCZ]) iJnqIt(s6k5uN0 common.Hash) int {
	return /*line a6l53NR.go:1*/ len(gSe6Sbe.lDwPGMKP3u_x[s6k5uN0])
}

func (gSe6Sbe *FZT207R[HSGwUUJCZ]) yqG1CuIwq(s6k5uN0 common.Hash) (crypto.Pp__49cd, []types.NodeID, error) {
	var y8S3i3vB2oK crypto.Pp__49cd
	var aZJ5AQFhNPiv []types.NodeID
	_, ov2W3mS3 := gSe6Sbe.lDwPGMKP3u_x[s6k5uN0]
	if !ov2W3mS3 {
		return nil, nil /*line mMIKZi9gZVK.go:1*/, fmt.Errorf(func() /*line ev2ddyuF1.go:1*/ string {
			data := /*line ev2ddyuF1.go:1*/ make([]byte, 0, 28)
			i := 4
			decryptKey := 139
			for counter := 0; i != 10; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data = /*line ev2ddyuF1.go:1*/ append(data, "ja\\i"...,
					)
					i = 9
				case 8:
					i = 13
					data = /*line ev2ddyuF1.go:1*/ append(data, "\x15e"...,
					)
				case 13:
					for y := range data {
						data[y] = data[y] + /*line ev2ddyuF1.go:1*/ byte(decryptKey^y)
					}
					i = 10
				case 6:
					data = /*line ev2ddyuF1.go:1*/ append(data, 87)
					i = 3
				case 1:
					i = 8
					data = /*line ev2ddyuF1.go:1*/ append(data, 15)
				case 12:
					i = 2
					data = /*line ev2ddyuF1.go:1*/ append(data, "\x1c^r"...,
					)
				case 5:
					i = 0
					data = /*line ev2ddyuF1.go:1*/ append(data, "[Y\x12"...,
					)
				case 7:
					i = 12
					data = /*line ev2ddyuF1.go:1*/ append(data, "mo"...,
					)
				case 0:
					data = /*line ev2ddyuF1.go:1*/ append(data, "\x03ME\x1c"...,
					)
					i = 1
				case 9:
					data = /*line ev2ddyuF1.go:1*/ append(data, "\x13X"...,
					)
					i = 11
				case 2:
					i = 5
					data = /*line ev2ddyuF1.go:1*/ append(data, 80)
				case 3:
					i = 7
					data = /*line ev2ddyuF1.go:1*/ append(data, "r k"...,
					)
				case 11:
					i = 6
					data = /*line ev2ddyuF1.go:1*/ append(data, 96)
				}
			}
			return /*line ev2ddyuF1.go:1*/ string(data)
		}(), s6k5uN0)
	}
	for _, wLUxkeF := range gSe6Sbe.lDwPGMKP3u_x[s6k5uN0] {
		y8S3i3vB2oK = /*line jyl5BsPw3Q.go:1*/ append(y8S3i3vB2oK, wLUxkeF.MqlfmVE9d)
		aZJ5AQFhNPiv = /*line DROsftElU5.go:1*/ append(aZJ5AQFhNPiv, wLUxkeF.WpsY_aZ)
	}

	return y8S3i3vB2oK, aZJ5AQFhNPiv, nil
}

func (gSe6Sbe *FZT207R[HSGwUUJCZ]) Delete(s6k5uN0 common.Hash) {
	/*line XVDOzp53.go:1*/ gSe6Sbe.aKxm6PV0a.Lock()
	defer /*line I5WOSxyFv1.go:1*/ gSe6Sbe.aKxm6PV0a.Unlock()
	_, iY3Dtu := gSe6Sbe.lDwPGMKP3u_x[s6k5uN0]
	if iY3Dtu {
		/*line jQIhOtZD0Oo.go:1*/ delete(gSe6Sbe.lDwPGMKP3u_x, s6k5uN0)
	}
}

var _ = fmt.Append
var _ sync.Cond
var _ crypto.Pp__49cd
var _ = log.CDebpj

const _ = types.ABORT

var _ = common.AbsolutePath
