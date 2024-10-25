//line :1
package state

import (
	"github.com/ethereum/go-ethereum/common"
)

type accessList struct {
	addresses map[common.Address]int
	slots     []map[common.Hash]struct{}
}

func (kFF49KO *accessList) ContainsAddress(so0TaNzLDdpc common.Address) bool {
	_, dNazL3Oz4 := kFF49KO.addresses[so0TaNzLDdpc]
	return dNazL3Oz4
}

func (kFF49KO *accessList) Contains(so0TaNzLDdpc common.Address, ibmfxwhUDRC common.Hash) (fsAQ1g bool, nZAaaiHQ9 bool) {
	iJQtkRishD3g, dNazL3Oz4 := kFF49KO.addresses[so0TaNzLDdpc]
	if !dNazL3Oz4 {

		return false, false
	}
	if iJQtkRishD3g == -1 {

		return true, false
	}
	_, nZAaaiHQ9 = kFF49KO.slots[iJQtkRishD3g][ibmfxwhUDRC]
	return true, nZAaaiHQ9
}

func viqfAs6EZc() *accessList {
	return &accessList{
		addresses: /*line rgFFab.go:1*/ make(map[common.Address]int),
	}
}

func (g5XBfGEf *accessList) Copy() *accessList {
	naayhAG6aM := /*line JarTYaA62ws.go:1*/ viqfAs6EZc()
	for _3JfOPUaHxI, gz5sHd := range g5XBfGEf.addresses {
		naayhAG6aM.addresses[_3JfOPUaHxI] = gz5sHd
	}
	naayhAG6aM.slots = /*line dOT2gdyKre86.go:1*/ make([]map[common.Hash]struct{} /*line HCu8KMqd.go:1*/, len(g5XBfGEf.slots))
	for cDClIaFDS, fn1gtmQa6 := range g5XBfGEf.slots {
		hS2HCc0r := /*line Ny2t4EkX.go:1*/ make(map[common.Hash]struct{} /*line jG412DYknP.go:1*/, len(fn1gtmQa6))
		for _3JfOPUaHxI := range fn1gtmQa6 {
			hS2HCc0r[_3JfOPUaHxI] = struct{}{}
		}
		naayhAG6aM.slots[cDClIaFDS] = hS2HCc0r
	}
	return naayhAG6aM
}

func (kFF49KO *accessList) AddAddress(so0TaNzLDdpc common.Address) bool {
	if _, vMmdZkjx960 := kFF49KO.addresses[so0TaNzLDdpc]; vMmdZkjx960 {
		return false
	}
	kFF49KO.addresses[so0TaNzLDdpc] = -1
	return true
}

func (kFF49KO *accessList) AddSlot(so0TaNzLDdpc common.Address, ibmfxwhUDRC common.Hash) (aswYwYJ bool, d9kkLAua bool) {
	iJQtkRishD3g, gfMgjM3q := kFF49KO.addresses[so0TaNzLDdpc]
	if !gfMgjM3q || iJQtkRishD3g == -1 {

		kFF49KO.addresses[so0TaNzLDdpc] = /*line Jzv3d2xim7T.go:1*/ len(kFF49KO.slots)
		iKg5R72b := map[common.Hash]struct{}{ibmfxwhUDRC: {}}
		kFF49KO.slots = /*line p4DnHs8aP.go:1*/ append(kFF49KO.slots, iKg5R72b)
		return !gfMgjM3q, true
	}

	iKg5R72b := kFF49KO.slots[iJQtkRishD3g]
	if _, dNazL3Oz4 := iKg5R72b[ibmfxwhUDRC]; !dNazL3Oz4 {
		iKg5R72b[ibmfxwhUDRC] = struct{}{}

		return false, true
	}

	return false, false
}

func (kFF49KO *accessList) DeleteSlot(so0TaNzLDdpc common.Address, ibmfxwhUDRC common.Hash) {
	iJQtkRishD3g, hgsbH7iL8n := kFF49KO.addresses[so0TaNzLDdpc]

	if !hgsbH7iL8n {
		/*line GxrAzzVVhQj.go:1*/ panic(func() /*line IS6I2pXWI4Z.go:1*/ string {
			fullData := [] /*line IS6I2pXWI4Z.go:1*/ byte("4\xadzy\r\x1f\x9dh\xd6U9w\xf9\x01!J\xec\xc80C\xbbo\xd3\xf7Q9a\xaa\xa6\xd3\xe0\xce7\x9e\x1b\x85\x99\x85C\xdb#\xb9\x19AC\a\xdeQ\xfa-O_\xf7*R\x92\x8d\xb5:\x90\x00+K^u-\x93\xc8\xcc-\x9d?&\xc8\x03?\x84?\xe8v2ބ([\x86\xa41\xab&\x9dHړ#\xc7\x13o\xd74")
			idxKey := [] /*line IS6I2pXWI4Z.go:1*/ byte("U\xf9\xb9\xe3\x14a=\x1f\xfdlË\xc6z\x96\xc2")
			data := /*line IS6I2pXWI4Z.go:1*/ make([]byte, 0, 51)
			data = /*line IS6I2pXWI4Z.go:1*/ append(data, fullData[33^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[7])]^fullData[21^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[7])], fullData[114^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])]-fullData[48^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])], fullData[211^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])]^fullData[236^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])], fullData[222^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])]+fullData[209^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])], fullData[222^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])]-fullData[217^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])], fullData[94^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[13])]+fullData[93^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[13])], fullData[207^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])]+fullData[186^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])], fullData[241^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])]+fullData[198^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])], fullData[201^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[1])]^fullData[255^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[1])], fullData[197^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[8])]+fullData[160^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[8])], fullData[34^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[13])]+fullData[57^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[13])], fullData[78^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])]^fullData[103^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])], fullData[197^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[15])]^fullData[239^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[15])], fullData[28^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]+fullData[53^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[172^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])]+fullData[248^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])], fullData[104^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[6])]-fullData[21^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[6])], fullData[200^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])]-fullData[239^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])], fullData[20^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]-fullData[2^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[97^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])]^fullData[13^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])], fullData[129^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])]+fullData[216^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])], fullData[249^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])]+fullData[254^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])], fullData[166^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])]^fullData[139^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])], fullData[16^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]^fullData[81^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[161^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])]+fullData[252^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])], fullData[10^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]^fullData[88^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[191^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[1])]-fullData[224^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[1])], fullData[88^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])]^fullData[79^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])], fullData[226^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[2])]^fullData[136^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[2])], fullData[22^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]+fullData[24^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[41^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[13])]^fullData[46^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[13])], fullData[50^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])]-fullData[38^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[9])], fullData[14^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[6])]^fullData[106^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[6])], fullData[250^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])]^fullData[159^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[10])], fullData[184^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[2])]+fullData[230^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[2])], fullData[176^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[2])]^fullData[249^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[2])], fullData[23^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]+fullData[3^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[130^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])]+fullData[218^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])], fullData[237^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[15])]^fullData[161^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[15])], fullData[49^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])]^fullData[74^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])], fullData[156^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])]+fullData[215^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[12])], fullData[169^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])]^fullData[132^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])], fullData[65^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])]^fullData[71^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])], fullData[51^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])]^fullData[55^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])], fullData[233^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[8])]-fullData[203^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[8])], fullData[40^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]-fullData[35^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])], fullData[216^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])]-fullData[223^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[14])], fullData[168^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])]+fullData[162^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[3])], fullData[113^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])]^fullData[68^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[5])], fullData[229^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[8])]-fullData[211^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[8])], fullData[33^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])]+fullData[27^ /*line IS6I2pXWI4Z.go:1*/ int(idxKey[4])])
			return /*line IS6I2pXWI4Z.go:1*/ string(data)
		}())
	}
	iKg5R72b := kFF49KO.slots[iJQtkRishD3g]
	/*line FCQJkAHFY0s.go:1*/ delete(iKg5R72b, ibmfxwhUDRC)

	if /*line OaAcaKa2wUiC.go:1*/ len(iKg5R72b) == 0 {
		kFF49KO.slots = kFF49KO.slots[:iJQtkRishD3g]
		kFF49KO.addresses[so0TaNzLDdpc] = -1
	}
}

func (kFF49KO *accessList) DeleteAddress(so0TaNzLDdpc common.Address) {
	/*line bSdsNfM0B4gR.go:1*/ delete(kFF49KO.addresses, so0TaNzLDdpc)
}

var _ = common.AbsolutePath
