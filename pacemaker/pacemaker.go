//line :1
package pacemaker

import (
	"sync"
	"time"

	blockchain "unishard/blockchain"
	config "unishard/config"

	types "unishard/types"
)

const ROTATING_COMMITTEE string = "Rotating_Committee"
const ROTATING_LEADER string = "Rotating_Leader"

type Pacemaker struct {
	as1T2l3UoGCM types.View
	oUiCc2oAv    types.View
	yRkDfB       types.View
	fs8vqPD2     int
	fL0kya3z     types.Epoch
	orlsDhQ      int
	l3G8aKyXh    chan types.EpochView
	ica_LoPo0j   chan types.EpochView
	mjgQzFnUN    *IWNgjZJHY
	yVo0TtId     sync.Mutex
}

func NewPacemaker(iDvh6le2Hy7Q int) *Pacemaker {
	vusGHk := /*line zd6uCYp.go:1*/ new(Pacemaker)
	vusGHk.l3G8aKyXh = /*line ml47tp.go:1*/ make(chan types.EpochView, 100)
	vusGHk.ica_LoPo0j = /*line PLAEqI53y1Go.go:1*/ make(chan types.EpochView, 100)
	vusGHk.mjgQzFnUN = /*line FUM8rO6mU.go:1*/ Cfc0vlx(iDvh6le2Hy7Q)
	vusGHk.orlsDhQ = /*line yfNVasuqIC1.go:1*/ config.GetConfig().ViewChangePeriod
	return vusGHk
}

func (vLHPFogOKZ *Pacemaker) ProcessRemoteTmo(eO94A1 *TMO) (bool, *TC, *blockchain.WorkerBlock) {
	if eO94A1.QmJZRap < vLHPFogOKZ.as1T2l3UoGCM {
		return false, nil, nil
	}
	return /*line m4JsYW3TY2I.go:1*/ vLHPFogOKZ.mjgQzFnUN.AddTmo(eO94A1)
}

func (vLHPFogOKZ *Pacemaker) ProcessRemoteMjorityTmo(eO94A1 *TMO) (bool, *TC, *blockchain.WorkerBlock) {
	if eO94A1.QmJZRap < vLHPFogOKZ.as1T2l3UoGCM {
		return false, nil, nil
	}
	return /*line Hw0MZx.go:1*/ vLHPFogOKZ.mjgQzFnUN.AddMjorityTmo(eO94A1)
}

func (vLHPFogOKZ *Pacemaker) AdvanceView(hwXqUFsD types.View) {
	/*line aERDGYXuhbUu.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line ELMYyvMnBHag.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	if hwXqUFsD < vLHPFogOKZ.as1T2l3UoGCM {
		return
	}

	if /*line EfqVQz4.go:1*/ config.GetConfig().RotatingElection == func() /*line StEnLGB1UirZ.go:1*/ string {
		fullData := [] /*line StEnLGB1UirZ.go:1*/ byte("\xb9\xb4a\x82tGS\x02\xc1\xb4\xbb'\x13\x8b\x16\xf5\x15^\xa5k\xaa\x9a\xffI\xbe\x86\x06B'\xee\xeb\xeam{\xd6\xd3")
		idxKey := [] /*line StEnLGB1UirZ.go:1*/ byte("\xf4\x10\xb0\xa6\x93\x031Q\xe7g")
		data := /*line StEnLGB1UirZ.go:1*/ make([]byte, 0, 19)
		data = /*line StEnLGB1UirZ.go:1*/ append(data, fullData[245^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[8])]-fullData[225^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[8])], fullData[50^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[1])]^fullData[16^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[1])], fullData[83^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[7])]+fullData[93^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[7])], fullData[247^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[8])]-fullData[238^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[8])], fullData[113^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[9])]^fullData[106^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[9])], fullData[66^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[7])]-fullData[86^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[7])], fullData[123^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[9])]^fullData[112^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[9])], fullData[179^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])]-fullData[137^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])], fullData[174^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[2])]+fullData[180^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[2])], fullData[57^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[6])]+fullData[50^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[6])], fullData[5^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[1])]^fullData[31^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[1])], fullData[80^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[7])]-fullData[84^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[7])], fullData[8^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[1])]^fullData[51^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[1])], fullData[42^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[6])]+fullData[58^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[6])], fullData[168^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[3])]+fullData[183^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[3])], fullData[142^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])]+fullData[138^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])], fullData[135^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])]+fullData[153^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])], fullData[140^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])]+fullData[178^ /*line StEnLGB1UirZ.go:1*/ int(idxKey[4])])
		return /*line StEnLGB1UirZ.go:1*/ string(data)
	}() {

		vLHPFogOKZ.fL0kya3z += 1
	} else if /*line iZKrhzUdzA7j.go:1*/ config.GetConfig().RotatingElection == func() /*line TI3saM.go:1*/ string {
		fullData := [] /*line TI3saM.go:1*/ byte("\t\xb3\xb4\x83\xa6p\xb2<\xfc\xefö\xe7\x9a#\xc1\xcbb\xb5sX\xc2lv\xac0\xa3\xaf\xd2\x11")
		idxKey := [] /*line TI3saM.go:1*/ byte("\xe6\x03\xb8\xef3\xa8m\xf6bڅ\aW\xe7")
		data := /*line TI3saM.go:1*/ make([]byte, 0, 16)
		data = /*line TI3saM.go:1*/ append(data, fullData[246^ /*line TI3saM.go:1*/ int(idxKey[13])]^fullData[254^ /*line TI3saM.go:1*/ int(idxKey[13])], fullData[113^ /*line TI3saM.go:1*/ int(idxKey[8])]+fullData[106^ /*line TI3saM.go:1*/ int(idxKey[8])], fullData[164^ /*line TI3saM.go:1*/ int(idxKey[2])]^fullData[188^ /*line TI3saM.go:1*/ int(idxKey[2])], fullData[173^ /*line TI3saM.go:1*/ int(idxKey[5])]^fullData[181^ /*line TI3saM.go:1*/ int(idxKey[5])], fullData[102^ /*line TI3saM.go:1*/ int(idxKey[6])]^fullData[120^ /*line TI3saM.go:1*/ int(idxKey[6])], fullData[88^ /*line TI3saM.go:1*/ int(idxKey[12])]-fullData[67^ /*line TI3saM.go:1*/ int(idxKey[12])], fullData[120^ /*line TI3saM.go:1*/ int(idxKey[8])]+fullData[114^ /*line TI3saM.go:1*/ int(idxKey[8])], fullData[132^ /*line TI3saM.go:1*/ int(idxKey[10])]+fullData[135^ /*line TI3saM.go:1*/ int(idxKey[10])], fullData[225^ /*line TI3saM.go:1*/ int(idxKey[3])]+fullData[232^ /*line TI3saM.go:1*/ int(idxKey[3])], fullData[190^ /*line TI3saM.go:1*/ int(idxKey[2])]+fullData[181^ /*line TI3saM.go:1*/ int(idxKey[2])], fullData[230^ /*line TI3saM.go:1*/ int(idxKey[0])]^fullData[240^ /*line TI3saM.go:1*/ int(idxKey[0])], fullData[31^ /*line TI3saM.go:1*/ int(idxKey[11])]+fullData[21^ /*line TI3saM.go:1*/ int(idxKey[11])], fullData[137^ /*line TI3saM.go:1*/ int(idxKey[10])]^fullData[134^ /*line TI3saM.go:1*/ int(idxKey[10])], fullData[238^ /*line TI3saM.go:1*/ int(idxKey[13])]+fullData[240^ /*line TI3saM.go:1*/ int(idxKey[13])], fullData[193^ /*line TI3saM.go:1*/ int(idxKey[9])]+fullData[208^ /*line TI3saM.go:1*/ int(idxKey[9])])
		return /*line TI3saM.go:1*/ string(data)
	}() {

		if vLHPFogOKZ.fs8vqPD2%(vLHPFogOKZ.orlsDhQ* /*line BG3Tn9v.go:1*/ config.GetConfig().CommitteeNumber) == 0 {
			vLHPFogOKZ.fL0kya3z += 1
		}
	}

	if vLHPFogOKZ.fs8vqPD2%vLHPFogOKZ.orlsDhQ == 0 && vLHPFogOKZ.fs8vqPD2/vLHPFogOKZ.orlsDhQ+1 > /*line pAvYkFhTr.go:1*/ int(hwXqUFsD) {
		vLHPFogOKZ.fs8vqPD2 += 1
		/*line ARhca2.go:1*/ vLHPFogOKZ.UpdateView(hwXqUFsD)
		return
	}
	vLHPFogOKZ.fs8vqPD2 += 1
	vLHPFogOKZ.l3G8aKyXh <- types.EpochView{Epoch: vLHPFogOKZ.fL0kya3z, View: vLHPFogOKZ.as1T2l3UoGCM}
}

func (vLHPFogOKZ *Pacemaker) UpdateView(hwXqUFsD types.View) {
	if /*line O_j0ksRSiYe.go:1*/ vLHPFogOKZ.yVo0TtId.TryLock() {
		defer /*line WiIHfjpqPzO.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	}

	if hwXqUFsD < vLHPFogOKZ.as1T2l3UoGCM {
		return
	}

	vLHPFogOKZ.as1T2l3UoGCM = hwXqUFsD + 1
	vLHPFogOKZ.oUiCc2oAv = (hwXqUFsD + 1) + 1
	vLHPFogOKZ.yRkDfB = hwXqUFsD + 1

	vLHPFogOKZ.l3G8aKyXh <- types.EpochView{Epoch: vLHPFogOKZ.fL0kya3z, View: vLHPFogOKZ.as1T2l3UoGCM}
}

func (vLHPFogOKZ *Pacemaker) AdvanceViewForFillHoleNode(hwXqUFsD types.View) {
	/*line Ufg247Nir.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line EgHdy47UGF.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	if hwXqUFsD < vLHPFogOKZ.as1T2l3UoGCM {
		return
	}

	if /*line dhsT8aS_b.go:1*/ config.GetConfig().RotatingElection == func() /*line YbzsHr.go:1*/ string {
		data := /*line YbzsHr.go:1*/ make([]byte, 0, 19)
		i := 3
		decryptKey := 71
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data = /*line YbzsHr.go:1*/ append(data, "\xe1\xef\xee"...,
				)
				i = 4
			case 8:
				i = 7
				data = /*line YbzsHr.go:1*/ append(data, "\xbf\xee\xeb\xe6"...,
				)
			case 1:
				i = 6
				data = /*line YbzsHr.go:1*/ append(data, "\xeb\xd7\xe5"...,
				)
			case 4:
				i = 2
				data = /*line YbzsHr.go:1*/ append(data, 202)
			case 5:
				for y := range data {
					data[y] = data[y] - /*line YbzsHr.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 3:
				i = 1
				data = /*line YbzsHr.go:1*/ append(data, "\xc7\xe3"...,
				)
			case 2:
				data = /*line YbzsHr.go:1*/ append(data, 201)
				i = 5
			case 6:
				i = 8
				data = /*line YbzsHr.go:1*/ append(data, "\xd9\xe1\xd9\xdc"...,
				)
			}
		}
		return /*line YbzsHr.go:1*/ string(data)
	}() {

		vLHPFogOKZ.fL0kya3z += 1
	} else if /*line Gsd0POh2QTXN.go:1*/ config.GetConfig().RotatingElection == func() /*line ziAaovlpWvqp.go:1*/ string {
		key := [] /*line ziAaovlpWvqp.go:1*/ byte("Z\x7f\xb6!l0\xf2\x14)\\eK7\x8f\xb9")
		data := [] /*line ziAaovlpWvqp.go:1*/ byte("\b\x10\xc2@\x18Y\x9csv\x10\x00*S\xea\xcb")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line ziAaovlpWvqp.go:1*/ string(data)
	}() {

		if vLHPFogOKZ.fs8vqPD2%(vLHPFogOKZ.orlsDhQ* /*line F3CJZTbYFdW.go:1*/ config.GetConfig().CommitteeNumber) == 0 {
			vLHPFogOKZ.fL0kya3z += 1
		}
	}

	if vLHPFogOKZ.fs8vqPD2%vLHPFogOKZ.orlsDhQ == 0 && vLHPFogOKZ.fs8vqPD2/vLHPFogOKZ.orlsDhQ+1 > /*line dhMyvnDpsq.go:1*/ int(hwXqUFsD) {
		vLHPFogOKZ.fs8vqPD2 += 1
		/*line O_E3GoyL_.go:1*/ vLHPFogOKZ.UpdateViewForFillholeNode(hwXqUFsD)
		return
	}
	vLHPFogOKZ.fs8vqPD2 += 1
}

func (vLHPFogOKZ *Pacemaker) UpdateViewForFillholeNode(hwXqUFsD types.View) {
	if /*line uHbdThedNAv0.go:1*/ vLHPFogOKZ.yVo0TtId.TryLock() {
		defer /*line EbUqZfjlSaf.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	}

	if hwXqUFsD < vLHPFogOKZ.as1T2l3UoGCM {
		return
	}

	vLHPFogOKZ.as1T2l3UoGCM = hwXqUFsD + 1
	vLHPFogOKZ.oUiCc2oAv = (hwXqUFsD + 1) + 1
	vLHPFogOKZ.yRkDfB = hwXqUFsD + 1
}

func (vLHPFogOKZ *Pacemaker) UpdateEpoch(stjIrJW types.Epoch) {
	/*line nTlwOjBQYt.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line FPYMHX6.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()

	if stjIrJW < vLHPFogOKZ.fL0kya3z {
		return
	}

	vLHPFogOKZ.fL0kya3z = stjIrJW
}

func (vLHPFogOKZ *Pacemaker) ExecuteViewChange(f8bkQu_wY types.View) {
	vLHPFogOKZ.ica_LoPo0j <- types.EpochView{Epoch: vLHPFogOKZ.fL0kya3z, View: f8bkQu_wY}
}

func (vLHPFogOKZ *Pacemaker) EnteringViewEvent() chan types.EpochView {
	return vLHPFogOKZ.l3G8aKyXh
}

func (vLHPFogOKZ *Pacemaker) EnteringTmoEvent() chan types.EpochView {
	return vLHPFogOKZ.ica_LoPo0j
}

func (vLHPFogOKZ *Pacemaker) UpdateAnchorView(h83Fjm types.View) {
	/*line eLM4EJKDpWOU.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line uwquoknw.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	vLHPFogOKZ.yRkDfB = h83Fjm
}

func (vLHPFogOKZ *Pacemaker) UpdateNewView(f8bkQu_wY types.View) {
	/*line E62SH9JldPUD.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line E9g_TYJQVv1B.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	vLHPFogOKZ.oUiCc2oAv = f8bkQu_wY
}

func (vLHPFogOKZ *Pacemaker) GetCurView() types.View {
	/*line wiM6kDoar9.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line hNGf2XF.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return vLHPFogOKZ.as1T2l3UoGCM
}
func (vLHPFogOKZ *Pacemaker) GetNewView() types.View {
	/*line ufzOpb.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line IVJlIf.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return vLHPFogOKZ.oUiCc2oAv
}
func (vLHPFogOKZ *Pacemaker) GetAnchorView() types.View {
	/*line DarXy2LGC7Y.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line AOV3SHpo8.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return vLHPFogOKZ.yRkDfB
}
func (vLHPFogOKZ *Pacemaker) GetCurRound() int {
	/*line BR_RuA8Sf.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line JstJI_RHltgJ.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return vLHPFogOKZ.fs8vqPD2
}

func (vLHPFogOKZ *Pacemaker) GetCurEpoch() types.Epoch {
	/*line qzMsBYjkrr.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line RKqJ9dMCDDO.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return vLHPFogOKZ.fL0kya3z
}

func (vLHPFogOKZ *Pacemaker) GetCurEpochView() (types.Epoch, types.View) {
	/*line yS9GTn.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line KetxM1.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return vLHPFogOKZ.fL0kya3z, vLHPFogOKZ.as1T2l3UoGCM
}

func (vLHPFogOKZ *Pacemaker) GetTimerForView() time.Duration {
	return /*line RhAUNS.go:1*/ time.Duration( /*line alTUQbeJG.go:1*/ config.GetConfig().Timeout) * time.Millisecond
}

func (vLHPFogOKZ *Pacemaker) GetTimerForViewChange() time.Duration {
	/*line PjNJde6JxP1.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line B40KNUtI.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()
	return /*line D98mKiu6usE.go:1*/ time.Duration( /*line MCYXVO3oB.go:1*/ int(vLHPFogOKZ.oUiCc2oAv-vLHPFogOKZ.yRkDfB)* /*line A88p6arr_.go:1*/ config.GetConfig().ViewChangeTimeout) * time.Millisecond
}

func (vLHPFogOKZ *Pacemaker) IsTimeToElect() bool {
	/*line XpdL4w6.go:1*/ vLHPFogOKZ.yVo0TtId.Lock()
	defer /*line fDOsGNCWcXo.go:1*/ vLHPFogOKZ.yVo0TtId.Unlock()

	if /*line JYpiv6No.go:1*/ config.GetConfig().RotatingElection == func() /*line oaJHiizKm8.go:1*/ string {
		key := [] /*line oaJHiizKm8.go:1*/ byte("\\\xc5\a\xe5+-\xa6g\xb7%\xef\v\xec)u\xd3oH")
		data := [] /*line oaJHiizKm8.go:1*/ byte("\xf6\xaam|I<\xc8\x00\xa8\x1e\x80b\x81@\xff\xa1\xf6\x1d")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line oaJHiizKm8.go:1*/ string(data)
	}() {
		return true
	}

	return vLHPFogOKZ.fs8vqPD2%(vLHPFogOKZ.orlsDhQ* /*line xxIwe0.go:1*/ config.GetConfig().CommitteeNumber) == 1
}

var _ sync.Cond

const _ = time.ANSIC

var _ blockchain.Accept
var _ config.Bconfig

const _ = types.ABORT
