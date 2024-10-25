//line :1
package pacemaker

import (
	"sync"
	"time"

	blockchain "unishard/blockchain"
	config "unishard/config"
	types "unishard/types"
)

type UDpSaa3 struct {
	b2PPNyvSz9a  types.View
	luUebPvwl    types.View
	jyjxwMSaZpI  types.View
	thDXeOFC     int
	iKJu2AzWYZO  types.Epoch
	mCfmfrnUGBo  int
	c4awd7Ya0T   chan types.EpochView
	bjf05wRHXA   chan types.EpochView
	e8HbUxtMNVpD *FjZVwJmDPT
	wcxgEdqaXcC  sync.Mutex
}

func IM6FKYh(iDvh6le2Hy7Q int) *UDpSaa3 {
	vusGHk := /*line DJTAhiRL9lTc.go:1*/ new(UDpSaa3)
	vusGHk.c4awd7Ya0T = /*line c3Na_5zjl.go:1*/ make(chan types.EpochView, 100)
	vusGHk.bjf05wRHXA = /*line HbwZYkK6.go:1*/ make(chan types.EpochView, 100)
	vusGHk.e8HbUxtMNVpD = /*line s4aitVM57YRk.go:1*/ Kg1AusRa8_(iDvh6le2Hy7Q)
	vusGHk.mCfmfrnUGBo = /*line Fx2fK4k8f.go:1*/ config.GetConfig().ViewChangePeriod
	return vusGHk
}

func (vLHPFogOKZ *UDpSaa3) ProcessRemoteTmo(eO94A1 *H8NP1AOKTF) (bool, *RZ65hTAdXj, *blockchain.CoordinationBlock) {
	if eO94A1.W72t5Nhk < vLHPFogOKZ.b2PPNyvSz9a {
		return false, nil, nil
	}
	return /*line LD4QZf.go:1*/ vLHPFogOKZ.e8HbUxtMNVpD.AddTmo(eO94A1)
}

func (vLHPFogOKZ *UDpSaa3) ProcessRemoteMjorityTmo(eO94A1 *H8NP1AOKTF) (bool, *RZ65hTAdXj, *blockchain.CoordinationBlock) {
	if eO94A1.W72t5Nhk < vLHPFogOKZ.b2PPNyvSz9a {
		return false, nil, nil
	}
	return /*line AUK2udqa237p.go:1*/ vLHPFogOKZ.e8HbUxtMNVpD.AddMjorityTmo(eO94A1)
}

func (vLHPFogOKZ *UDpSaa3) AdvanceView(hwXqUFsD types.View) {
	/*line beW2fnuIf.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line ydUC2GbS84q.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	if hwXqUFsD < vLHPFogOKZ.b2PPNyvSz9a {
		return
	}

	if /*line GY6f5GCMt21q.go:1*/ config.GetConfig().RotatingElection == func() /*line bGoW177x.go:1*/ string {
		key := [] /*line bGoW177x.go:1*/ byte("BZKG)\xb8'\xe6\xe6\x00\x88\x86\x92\x8c\x9b\xe1\xc1<")
		data := [] /*line bGoW177x.go:1*/ byte("\x94ɿ\xa8\x9d!\x95MEC\xf7\xf3\xff\xf5\x0fU&\xa1")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line bGoW177x.go:1*/ string(data)
	}() {

		vLHPFogOKZ.iKJu2AzWYZO += 1
	} else if /*line bn0GPfa13P.go:1*/ config.GetConfig().RotatingElection == func() /*line NmcvI8u2hO.go:1*/ string {
		key := [] /*line NmcvI8u2hO.go:1*/ byte("\xac\xe8{wz\x87\xa9\xadnJ\xbeC\xa0e\n")
		data := [] /*line NmcvI8u2hO.go:1*/ byte("\xfe\x87\x0f\x16\x0e\xee\xc7\xca1\x06\xdb\"\xc4\x00x")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line NmcvI8u2hO.go:1*/ string(data)
	}() {

		if vLHPFogOKZ.thDXeOFC%(vLHPFogOKZ.mCfmfrnUGBo* /*line MXEfQl.go:1*/ config.GetConfig().CommitteeNumber) == 0 {
			vLHPFogOKZ.iKJu2AzWYZO += 1
		}
	}

	if vLHPFogOKZ.thDXeOFC%vLHPFogOKZ.mCfmfrnUGBo == 0 && vLHPFogOKZ.thDXeOFC/vLHPFogOKZ.mCfmfrnUGBo+1 > /*line pKDcZxgM.go:1*/ int(hwXqUFsD) {
		vLHPFogOKZ.thDXeOFC += 1
		/*line XfbKLC33r16A.go:1*/ vLHPFogOKZ.UpdateView(hwXqUFsD)
		return
	}
	vLHPFogOKZ.thDXeOFC += 1
	vLHPFogOKZ.c4awd7Ya0T <- types.EpochView{Epoch: vLHPFogOKZ.iKJu2AzWYZO, View: vLHPFogOKZ.b2PPNyvSz9a}
}

func (vLHPFogOKZ *UDpSaa3) UpdateView(hwXqUFsD types.View) {
	if /*line h6SMysya.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.TryLock() {
		defer /*line LCIakNQbawRS.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	}

	if hwXqUFsD < vLHPFogOKZ.b2PPNyvSz9a {
		return
	}

	vLHPFogOKZ.b2PPNyvSz9a = hwXqUFsD + 1
	vLHPFogOKZ.luUebPvwl = (hwXqUFsD + 1) + 1
	vLHPFogOKZ.jyjxwMSaZpI = hwXqUFsD + 1

	vLHPFogOKZ.c4awd7Ya0T <- types.EpochView{Epoch: vLHPFogOKZ.iKJu2AzWYZO, View: vLHPFogOKZ.b2PPNyvSz9a}
}

func (vLHPFogOKZ *UDpSaa3) AdvanceViewForFillHoleNode(hwXqUFsD types.View) {
	/*line MtXgsSVbm.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line T0IiUWZFny.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	if hwXqUFsD < vLHPFogOKZ.b2PPNyvSz9a {
		return
	}

	if /*line anS_AEu.go:1*/ config.GetConfig().RotatingElection == func() /*line JlGxVL5.go:1*/ string {
		fullData := [] /*line JlGxVL5.go:1*/ byte("\xb3h\x1b\xff:9\xe8\x86EAs\xfb\xf9M\x1f\x8cR\xbbj\xb7\xb0\x90!\xe1\x8czԮ\xd3{u[\xf3(&\xb6")
		idxKey := [] /*line JlGxVL5.go:1*/ byte("\xf7\xd4")
		data := /*line JlGxVL5.go:1*/ make([]byte, 0, 19)
		data = /*line JlGxVL5.go:1*/ append(data, fullData[253^ /*line JlGxVL5.go:1*/ int(idxKey[0])]^fullData[225^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[237^ /*line JlGxVL5.go:1*/ int(idxKey[0])]^fullData[230^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[251^ /*line JlGxVL5.go:1*/ int(idxKey[0])]+fullData[234^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[203^ /*line JlGxVL5.go:1*/ int(idxKey[1])]^fullData[208^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[196^ /*line JlGxVL5.go:1*/ int(idxKey[1])]^fullData[246^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[246^ /*line JlGxVL5.go:1*/ int(idxKey[0])]-fullData[244^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[202^ /*line JlGxVL5.go:1*/ int(idxKey[1])]^fullData[214^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[195^ /*line JlGxVL5.go:1*/ int(idxKey[1])]-fullData[205^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[200^ /*line JlGxVL5.go:1*/ int(idxKey[1])]^fullData[204^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[247^ /*line JlGxVL5.go:1*/ int(idxKey[0])]+fullData[226^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[192^ /*line JlGxVL5.go:1*/ int(idxKey[1])]-fullData[221^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[245^ /*line JlGxVL5.go:1*/ int(idxKey[1])]+fullData[220^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[244^ /*line JlGxVL5.go:1*/ int(idxKey[1])]-fullData[211^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[218^ /*line JlGxVL5.go:1*/ int(idxKey[1])]-fullData[247^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[242^ /*line JlGxVL5.go:1*/ int(idxKey[0])]^fullData[250^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[248^ /*line JlGxVL5.go:1*/ int(idxKey[0])]+fullData[241^ /*line JlGxVL5.go:1*/ int(idxKey[0])], fullData[199^ /*line JlGxVL5.go:1*/ int(idxKey[1])]+fullData[207^ /*line JlGxVL5.go:1*/ int(idxKey[1])], fullData[252^ /*line JlGxVL5.go:1*/ int(idxKey[0])]+fullData[229^ /*line JlGxVL5.go:1*/ int(idxKey[0])])
		return /*line JlGxVL5.go:1*/ string(data)
	}() {

		vLHPFogOKZ.iKJu2AzWYZO += 1
	} else if /*line crcjrQKWr.go:1*/ config.GetConfig().RotatingElection == func() /*line Eg1Eitmcag.go:1*/ string {
		fullData := [] /*line Eg1Eitmcag.go:1*/ byte("f\xfd\x95;\xe2\xdfCO\b\x1aJ \xbc\x9f8\xe2\x9b\xf1ڢ\xbb~`\xe8\n\x96\xd4\\\t\x01")
		idxKey := [] /*line Eg1Eitmcag.go:1*/ byte("\xe6\xf7\x97\x90\xfa ")
		data := /*line Eg1Eitmcag.go:1*/ make([]byte, 0, 16)
		data = /*line Eg1Eitmcag.go:1*/ append(data, fullData[140^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])]-fullData[143^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])], fullData[58^ /*line Eg1Eitmcag.go:1*/ int(idxKey[5])]+fullData[48^ /*line Eg1Eitmcag.go:1*/ int(idxKey[5])], fullData[142^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])]^fullData[147^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])], fullData[252^ /*line Eg1Eitmcag.go:1*/ int(idxKey[4])]-fullData[245^ /*line Eg1Eitmcag.go:1*/ int(idxKey[4])], fullData[227^ /*line Eg1Eitmcag.go:1*/ int(idxKey[0])]+fullData[228^ /*line Eg1Eitmcag.go:1*/ int(idxKey[0])], fullData[140^ /*line Eg1Eitmcag.go:1*/ int(idxKey[3])]+fullData[134^ /*line Eg1Eitmcag.go:1*/ int(idxKey[3])], fullData[235^ /*line Eg1Eitmcag.go:1*/ int(idxKey[4])]^fullData[247^ /*line Eg1Eitmcag.go:1*/ int(idxKey[4])], fullData[151^ /*line Eg1Eitmcag.go:1*/ int(idxKey[3])]-fullData[135^ /*line Eg1Eitmcag.go:1*/ int(idxKey[3])], fullData[251^ /*line Eg1Eitmcag.go:1*/ int(idxKey[4])]^fullData[233^ /*line Eg1Eitmcag.go:1*/ int(idxKey[4])], fullData[159^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])]-fullData[155^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])], fullData[43^ /*line Eg1Eitmcag.go:1*/ int(idxKey[5])]-fullData[52^ /*line Eg1Eitmcag.go:1*/ int(idxKey[5])], fullData[148^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])]-fullData[133^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])], fullData[158^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])]^fullData[130^ /*line Eg1Eitmcag.go:1*/ int(idxKey[2])], fullData[230^ /*line Eg1Eitmcag.go:1*/ int(idxKey[0])]-fullData[251^ /*line Eg1Eitmcag.go:1*/ int(idxKey[0])], fullData[236^ /*line Eg1Eitmcag.go:1*/ int(idxKey[0])]^fullData[232^ /*line Eg1Eitmcag.go:1*/ int(idxKey[0])])
		return /*line Eg1Eitmcag.go:1*/ string(data)
	}() {

		if vLHPFogOKZ.thDXeOFC%(vLHPFogOKZ.mCfmfrnUGBo* /*line OAav7f.go:1*/ config.GetConfig().CommitteeNumber) == 0 {
			vLHPFogOKZ.iKJu2AzWYZO += 1
		}
	}

	if vLHPFogOKZ.thDXeOFC%vLHPFogOKZ.mCfmfrnUGBo == 0 && vLHPFogOKZ.thDXeOFC/vLHPFogOKZ.mCfmfrnUGBo+1 > /*line JZGqhiYDmL_g.go:1*/ int(hwXqUFsD) {
		vLHPFogOKZ.thDXeOFC += 1
		/*line izJK0NujjMC.go:1*/ vLHPFogOKZ.UpdateViewForFillholeNode(hwXqUFsD)
		return
	}
	vLHPFogOKZ.thDXeOFC += 1
}

func (vLHPFogOKZ *UDpSaa3) UpdateViewForFillholeNode(hwXqUFsD types.View) {
	if /*line IYJltw0nLJCp.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.TryLock() {
		defer /*line v9xR2X7.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	}

	if hwXqUFsD < vLHPFogOKZ.b2PPNyvSz9a {
		return
	}

	vLHPFogOKZ.b2PPNyvSz9a = hwXqUFsD + 1
	vLHPFogOKZ.luUebPvwl = (hwXqUFsD + 1) + 1
	vLHPFogOKZ.jyjxwMSaZpI = hwXqUFsD + 1
}

func (vLHPFogOKZ *UDpSaa3) UpdateEpoch(stjIrJW types.Epoch) {
	/*line TvUhzdE2.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line mpS7MW5.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()

	if stjIrJW < vLHPFogOKZ.iKJu2AzWYZO {
		return
	}

	vLHPFogOKZ.iKJu2AzWYZO = stjIrJW
}

func (vLHPFogOKZ *UDpSaa3) ExecuteViewChange(f8bkQu_wY types.View) {
	vLHPFogOKZ.bjf05wRHXA <- types.EpochView{Epoch: vLHPFogOKZ.iKJu2AzWYZO, View: f8bkQu_wY}
}

func (vLHPFogOKZ *UDpSaa3) EnteringViewEvent() chan types.EpochView {
	return vLHPFogOKZ.c4awd7Ya0T
}

func (vLHPFogOKZ *UDpSaa3) EnteringTmoEvent() chan types.EpochView {
	return vLHPFogOKZ.bjf05wRHXA
}

func (vLHPFogOKZ *UDpSaa3) UpdateAnchorView(h83Fjm types.View) {
	/*line EH3wPoG5rkk.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line DjHP7yuqS.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	vLHPFogOKZ.jyjxwMSaZpI = h83Fjm
}

func (vLHPFogOKZ *UDpSaa3) UpdateNewView(f8bkQu_wY types.View) {
	/*line _fV7pe.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line tZUoPWmy.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	vLHPFogOKZ.luUebPvwl = f8bkQu_wY
}

func (vLHPFogOKZ *UDpSaa3) GetCurView() types.View {
	/*line PI4onLJft8.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line _YnPaJHFb.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return vLHPFogOKZ.b2PPNyvSz9a
}
func (vLHPFogOKZ *UDpSaa3) GetNewView() types.View {
	/*line xW_Rk_kgguM.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line vmWTiOxh.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return vLHPFogOKZ.luUebPvwl
}
func (vLHPFogOKZ *UDpSaa3) GetAnchorView() types.View {
	/*line idyfqX.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line INmQA39E2Dv0.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return vLHPFogOKZ.jyjxwMSaZpI
}
func (vLHPFogOKZ *UDpSaa3) GetCurRound() int {
	/*line msypU5EWGKpz.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line Owo5Koqys.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return vLHPFogOKZ.thDXeOFC
}

func (vLHPFogOKZ *UDpSaa3) GetCurEpoch() types.Epoch {
	/*line _vr5o8d5sPzc.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line lYaKSBkqlqM.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return vLHPFogOKZ.iKJu2AzWYZO
}

func (vLHPFogOKZ *UDpSaa3) GetCurEpochView() (types.Epoch, types.View) {
	/*line SlhxHUXDK3lH.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line xjMca7z.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return vLHPFogOKZ.iKJu2AzWYZO, vLHPFogOKZ.b2PPNyvSz9a
}

func (vLHPFogOKZ *UDpSaa3) GetTimerForView() time.Duration {
	return /*line wG1Zxvm.go:1*/ time.Duration( /*line aDY_LR2Palgi.go:1*/ config.GetConfig().Timeout) * time.Millisecond
}

func (vLHPFogOKZ *UDpSaa3) GetTimerForViewChange() time.Duration {
	/*line O1TFK36lmfW.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line SEgCah4bY8d.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()
	return /*line AakSVMTe3.go:1*/ time.Duration( /*line JNl0buey0.go:1*/ int(vLHPFogOKZ.luUebPvwl-vLHPFogOKZ.jyjxwMSaZpI)* /*line QxyTcdF73fkD.go:1*/ config.GetConfig().ViewChangeTimeout) * time.Millisecond
}

func (vLHPFogOKZ *UDpSaa3) IsTimeToElect() bool {
	/*line kufevhq.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Lock()
	defer /*line Jp6XDqwZ8s.go:1*/ vLHPFogOKZ.wcxgEdqaXcC.Unlock()

	if /*line xSYvLjZI8_I2.go:1*/ config.GetConfig().RotatingElection == func() /*line EcUoqg.go:1*/ string {
		fullData := [] /*line EcUoqg.go:1*/ byte("y\az\xe1q\xfc\x97#^\xbf\xf7?\x946h#1\f\x95S\xedKa\xe6\x92*\x87B\xe2\xc5\xd7\xc2\xf1\xb6\xa4V")
		idxKey := [] /*line EcUoqg.go:1*/ byte("\xdd2e\xc3\aH\xfb+\xaf\xd3\xd2@!\xd6\x7f")
		data := /*line EcUoqg.go:1*/ make([]byte, 0, 19)
		data = /*line EcUoqg.go:1*/ append(data, fullData[213^ /*line EcUoqg.go:1*/ int(idxKey[3])]+fullData[227^ /*line EcUoqg.go:1*/ int(idxKey[3])], fullData[216^ /*line EcUoqg.go:1*/ int(idxKey[13])]+fullData[215^ /*line EcUoqg.go:1*/ int(idxKey[13])], fullData[192^ /*line EcUoqg.go:1*/ int(idxKey[10])]^fullData[209^ /*line EcUoqg.go:1*/ int(idxKey[10])], fullData[38^ /*line EcUoqg.go:1*/ int(idxKey[12])]-fullData[62^ /*line EcUoqg.go:1*/ int(idxKey[12])], fullData[19^ /*line EcUoqg.go:1*/ int(idxKey[1])]-fullData[41^ /*line EcUoqg.go:1*/ int(idxKey[1])], fullData[202^ /*line EcUoqg.go:1*/ int(idxKey[3])]-fullData[224^ /*line EcUoqg.go:1*/ int(idxKey[3])], fullData[225^ /*line EcUoqg.go:1*/ int(idxKey[3])]-fullData[206^ /*line EcUoqg.go:1*/ int(idxKey[3])], fullData[249^ /*line EcUoqg.go:1*/ int(idxKey[6])]+fullData[239^ /*line EcUoqg.go:1*/ int(idxKey[6])], fullData[48^ /*line EcUoqg.go:1*/ int(idxKey[12])]+fullData[50^ /*line EcUoqg.go:1*/ int(idxKey[12])], fullData[209^ /*line EcUoqg.go:1*/ int(idxKey[0])]^fullData[195^ /*line EcUoqg.go:1*/ int(idxKey[0])], fullData[15^ /*line EcUoqg.go:1*/ int(idxKey[4])]^fullData[23^ /*line EcUoqg.go:1*/ int(idxKey[4])], fullData[215^ /*line EcUoqg.go:1*/ int(idxKey[9])]+fullData[214^ /*line EcUoqg.go:1*/ int(idxKey[9])], fullData[114^ /*line EcUoqg.go:1*/ int(idxKey[2])]-fullData[101^ /*line EcUoqg.go:1*/ int(idxKey[2])], fullData[101^ /*line EcUoqg.go:1*/ int(idxKey[14])]+fullData[99^ /*line EcUoqg.go:1*/ int(idxKey[14])], fullData[219^ /*line EcUoqg.go:1*/ int(idxKey[0])]-fullData[210^ /*line EcUoqg.go:1*/ int(idxKey[0])], fullData[116^ /*line EcUoqg.go:1*/ int(idxKey[14])]^fullData[106^ /*line EcUoqg.go:1*/ int(idxKey[14])], fullData[117^ /*line EcUoqg.go:1*/ int(idxKey[14])]^fullData[103^ /*line EcUoqg.go:1*/ int(idxKey[14])], fullData[30^ /*line EcUoqg.go:1*/ int(idxKey[4])]-fullData[26^ /*line EcUoqg.go:1*/ int(idxKey[4])])
		return /*line EcUoqg.go:1*/ string(data)
	}() {
		return true
	}

	return vLHPFogOKZ.thDXeOFC%(vLHPFogOKZ.mCfmfrnUGBo* /*line tL1TLFapwxv.go:1*/ config.GetConfig().CommitteeNumber) == 1
}

var _ sync.Cond

const _ = time.ANSIC

var _ blockchain.Accept
var _ config.Bconfig

const _ = types.ABORT
