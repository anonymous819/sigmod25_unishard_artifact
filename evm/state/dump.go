//line :1
package state

import (
	"encoding/json"
	"fmt"
	"time"

	types "unishard/evm/types"

	trie "unishard/evm/trie"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

type FREnjDRFF struct {
	WLCvmi0w2L  bool
	SLFA_jI     bool
	D_4N8FCzH   bool
	R8sL1CkH    []byte
	FaNH1vN2d2M uint64
}

type TSzQva interface {
	OnRoot(common.Hash)

	OnAccount(*common.Address, DumpAccount)
}

type DumpAccount struct {
	Balance     string                 `json:"balance"`
	Nonce       uint64                 `json:"nonce"`
	Root        hexutil.Bytes          `json:"root"`
	CodeHash    hexutil.Bytes          `json:"codeHash"`
	Code        hexutil.Bytes          `json:"code,omitempty"`
	Storage     map[common.Hash]string `json:"storage,omitempty"`
	Address     *common.Address        `json:"address,omitempty"`
	AddressHash hexutil.Bytes          `json:"key,omitempty"`
}

type Cag7Uq struct {
	JOoaQOQA1    string                 `json:"root"`
	JPjIz1vACvvV map[string]DumpAccount `json:"accounts"`

	PK5ng8Tn []byte `json:"next,omitempty"`
}

func (pajPTPPInV *Cag7Uq) OnRoot(lxItfhw common.Hash) {
	pajPTPPInV.JOoaQOQA1 = /*line WXiyO8.go:1*/ fmt.Sprintf("%x", lxItfhw)
}

func (pajPTPPInV *Cag7Uq) OnAccount(hv1UkMRaKf *common.Address, hJAE3TNQR_k2 DumpAccount) {
	if hv1UkMRaKf == nil {
		pajPTPPInV.JPjIz1vACvvV[ /*line JcY1HC.go:1*/ fmt.Sprintf("pre(%s)", hJAE3TNQR_k2.AddressHash)] = hJAE3TNQR_k2
	}
	if hv1UkMRaKf != nil {
		pajPTPPInV.JPjIz1vACvvV[(* /*line OLVgb2SGT.go:1*/ hv1UkMRaKf).String()] = hJAE3TNQR_k2
	}
}

type dHA8ny struct {
	*json.Encoder
}

func (pajPTPPInV dHA8ny) OnAccount(hv1UkMRaKf *common.Address, hJAE3TNQR_k2 DumpAccount) {
	jvaR40CUi := &DumpAccount{
		Balance:     hJAE3TNQR_k2.Balance,
		Nonce:       hJAE3TNQR_k2.Nonce,
		Root:        hJAE3TNQR_k2.Root,
		CodeHash:    hJAE3TNQR_k2.CodeHash,
		Code:        hJAE3TNQR_k2.Code,
		Storage:     hJAE3TNQR_k2.Storage,
		AddressHash: hJAE3TNQR_k2.AddressHash,
		Address:     hv1UkMRaKf,
	}
	/*line E7uju2as.go:1*/ pajPTPPInV.Encode(jvaR40CUi)
}

func (pajPTPPInV dHA8ny) OnRoot(lxItfhw common.Hash) {
	/*line EDv8eD_jtG.go:1*/ pajPTPPInV.Encode(struct {
		Root common.Hash `json:"root"`
	}{lxItfhw})
}

func (peYZBpNwox *StateDB) DumpToCollector(ts1y5Po TSzQva, pXDmAV88ut1 *FREnjDRFF) (nI13nxu []byte) {

	if pXDmAV88ut1 == nil {
		pXDmAV88ut1 = /*line rTqZcMcL.go:1*/ new(FREnjDRFF)
	}
	var (
		qlzJL5Jva    int
		cWM9wg4djowd uint64
		stqEk4Brz1wF = /*line LyLnpByvrZHD.go:1*/ time.Now()
		uOGokjlZcG7  = /*line Fj0yB0JZDSYx.go:1*/ time.Now()
	)
	/*line gWjocD.go:1*/ log.Info(func() /*line QMMa7i.go:1*/ string {
		fullData := [] /*line QMMa7i.go:1*/ byte("\xeah\xc5\x10d\xa5\x88\xba\x18\xa8\x12#F\x8c\x8b=\u05f52\xc8\xcb\xd9\xe6T0\x7f\x8f(J\xf8'\x8b\x98\x82\x9c\xf6%`\xafL")
		idxKey := [] /*line QMMa7i.go:1*/ byte("%\xa1\xa7\x82A\x8f\xe4\xca\xf2\xd9R\xc6\x05!Ί")
		data := /*line QMMa7i.go:1*/ make([]byte, 0, 21)
		data = /*line QMMa7i.go:1*/ append(data, fullData[12^ /*line QMMa7i.go:1*/ int(idxKey[12])]-fullData[18^ /*line QMMa7i.go:1*/ int(idxKey[12])], fullData[75^ /*line QMMa7i.go:1*/ int(idxKey[4])]+fullData[100^ /*line QMMa7i.go:1*/ int(idxKey[4])], fullData[42^ /*line QMMa7i.go:1*/ int(idxKey[13])]^fullData[61^ /*line QMMa7i.go:1*/ int(idxKey[13])], fullData[30^ /*line QMMa7i.go:1*/ int(idxKey[12])]+fullData[10^ /*line QMMa7i.go:1*/ int(idxKey[12])], fullData[200^ /*line QMMa7i.go:1*/ int(idxKey[14])]-fullData[207^ /*line QMMa7i.go:1*/ int(idxKey[14])], fullData[225^ /*line QMMa7i.go:1*/ int(idxKey[8])]-fullData[246^ /*line QMMa7i.go:1*/ int(idxKey[8])], fullData[160^ /*line QMMa7i.go:1*/ int(idxKey[3])]-fullData[156^ /*line QMMa7i.go:1*/ int(idxKey[3])], fullData[198^ /*line QMMa7i.go:1*/ int(idxKey[9])]^fullData[207^ /*line QMMa7i.go:1*/ int(idxKey[9])], fullData[182^ /*line QMMa7i.go:1*/ int(idxKey[2])]^fullData[165^ /*line QMMa7i.go:1*/ int(idxKey[2])], fullData[211^ /*line QMMa7i.go:1*/ int(idxKey[14])]-fullData[212^ /*line QMMa7i.go:1*/ int(idxKey[14])], fullData[38^ /*line QMMa7i.go:1*/ int(idxKey[13])]-fullData[6^ /*line QMMa7i.go:1*/ int(idxKey[13])], fullData[192^ /*line QMMa7i.go:1*/ int(idxKey[9])]^fullData[209^ /*line QMMa7i.go:1*/ int(idxKey[9])], fullData[61^ /*line QMMa7i.go:1*/ int(idxKey[0])]^fullData[38^ /*line QMMa7i.go:1*/ int(idxKey[0])], fullData[162^ /*line QMMa7i.go:1*/ int(idxKey[3])]-fullData[166^ /*line QMMa7i.go:1*/ int(idxKey[3])], fullData[134^ /*line QMMa7i.go:1*/ int(idxKey[2])]^fullData[132^ /*line QMMa7i.go:1*/ int(idxKey[2])], fullData[242^ /*line QMMa7i.go:1*/ int(idxKey[8])]^fullData[252^ /*line QMMa7i.go:1*/ int(idxKey[8])], fullData[68^ /*line QMMa7i.go:1*/ int(idxKey[4])]^fullData[81^ /*line QMMa7i.go:1*/ int(idxKey[4])], fullData[144^ /*line QMMa7i.go:1*/ int(idxKey[3])]^fullData[142^ /*line QMMa7i.go:1*/ int(idxKey[3])], fullData[180^ /*line QMMa7i.go:1*/ int(idxKey[1])]+fullData[172^ /*line QMMa7i.go:1*/ int(idxKey[1])], fullData[169^ /*line QMMa7i.go:1*/ int(idxKey[5])]^fullData[155^ /*line QMMa7i.go:1*/ int(idxKey[5])])
		return /*line QMMa7i.go:1*/ string(data)
	}(), "root" /*line I9lCkla.go:1*/, peYZBpNwox.trie.Hash())
	/*line IjXabqJL.go:1*/ ts1y5Po.OnRoot( /*line SHSMR3D6J.go:1*/ peYZBpNwox.trie.Hash())

	yVQYehp9F, cYHA75qVq := /*line MOaHUVLeBJ.go:1*/ peYZBpNwox.trie.NodeIterator(pXDmAV88ut1.R8sL1CkH)
	if cYHA75qVq != nil {
		/*line QM1e8jGOf.go:1*/ log.Error(func() /*line a9SGN5n.go:1*/ string {
			data := [] /*line a9SGN5n.go:1*/ byte("\xbd\xaci\xb8\"d\xb2g頇\xa6\xba\xb9r\xb5\xb3r")
			positions := [...]byte{15, 6, 13, 16, 12, 1, 4, 8, 7, 11, 11, 3, 12, 11, 9, 11, 12, 8, 0, 10}
			for i := 0; i < 20; i += 2 {
				localKey := /*line a9SGN5n.go:1*/ byte(i) + /*line a9SGN5n.go:1*/ byte(positions[i]^positions[i+1]) + 183
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line a9SGN5n.go:1*/ string(data)
		}(), "err", cYHA75qVq)
		return nil
	}
	s8hOm6P := /*line PzMvNpceA.go:1*/ trie.YH_VLHxOhT4(yVQYehp9F)
	for /*line Jm0aJcqt_a.go:1*/ s8hOm6P.Next() {
		var tfksAzzqIpKa types.StateAccount
		if cYHA75qVq := /*line CDkGR1MLnA.go:1*/ rlp.DecodeBytes(s8hOm6P.H884Yc, &tfksAzzqIpKa); cYHA75qVq != nil {
			/*line wwmILLk.go:1*/ panic(cYHA75qVq)
		}
		var (
			hJAE3TNQR_k2 = DumpAccount{
				Balance:/*line HytPYG0.go:1*/ tfksAzzqIpKa.Balance.String(),
				Nonce:       tfksAzzqIpKa.Nonce,
				Root:        tfksAzzqIpKa.Root[:],
				CodeHash:    tfksAzzqIpKa.CodeHash,
				AddressHash: s8hOm6P.JyqRfQ2XM6,
			}
			so0TaNzLDdpc *common.Address
			hv1UkMRaKf   common.Address
			x_CxRf_e5DB9 = /*line mFLUuV.go:1*/ peYZBpNwox.trie.GetKey(s8hOm6P.JyqRfQ2XM6)
		)
		if x_CxRf_e5DB9 == nil {
			qlzJL5Jva++
			if pXDmAV88ut1.D_4N8FCzH {
				continue
			}
		} else {
			hv1UkMRaKf = /*line Xy3An5pg6.go:1*/ common.BytesToAddress(x_CxRf_e5DB9)
			so0TaNzLDdpc = &hv1UkMRaKf
			hJAE3TNQR_k2.Address = so0TaNzLDdpc
		}
		ws3tNR3MU := /*line j3_PVG.go:1*/ zkEv6T(peYZBpNwox, hv1UkMRaKf, &tfksAzzqIpKa)
		if !pXDmAV88ut1.WLCvmi0w2L {
			hJAE3TNQR_k2.Code = /*line TKD51kG.go:1*/ ws3tNR3MU.Code()
		}
		if !pXDmAV88ut1.SLFA_jI {
			hJAE3TNQR_k2.Storage = /*line adGZ98oYyVcK.go:1*/ make(map[common.Hash]string)
			r8CzlEu, cYHA75qVq := /*line lSy6g_.go:1*/ ws3tNR3MU.z1F2PO()
			if cYHA75qVq != nil {
				/*line uKuBd0IeZXsb.go:1*/ log.Error(func() /*line CLDQQ1hGtU.go:1*/ string {
					data := /*line CLDQQ1hGtU.go:1*/ make([]byte, 0, 28)
					i := 8
					decryptKey := 19
					for counter := 0; i != 3; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 1:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xfe\xfc"...,
							)
							i = 6
						case 12:
							i = 9
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xae\xfb"...,
							)
						case 2:
							i = 12
							data = /*line CLDQQ1hGtU.go:1*/ append(data, 232)
						case 11:
							i = 2
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xe6\xf8\xea\xeb"...,
							)
						case 0:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xbe\xeb"...,
							)
							i = 7
						case 9:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xf2\xe8"...,
							)
							i = 10
						case 6:
							i = 4
							data = /*line CLDQQ1hGtU.go:1*/ append(data, 245)
						case 13:
							for y := range data {
								data[y] = data[y] ^ /*line CLDQQ1hGtU.go:1*/ byte(decryptKey^y)
							}
							i = 3
						case 7:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xff\xb1"...,
							)
							i = 1
						case 4:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xf1\xb6\xe4\xfc"...,
							)
							i = 11
						case 10:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, 231)
							i = 13
						case 5:
							i = 0
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xf9\xf9"...,
							)
						case 8:
							data = /*line CLDQQ1hGtU.go:1*/ append(data, "\xde\xf8\xf3\xf7"...,
							)
							i = 5
						}
					}
					return /*line CLDQQ1hGtU.go:1*/ string(data)
				}(), "err", cYHA75qVq)
				continue
			}
			yVQYehp9F, cYHA75qVq := /*line Aj5PCWXZ.go:1*/ r8CzlEu.NodeIterator(nil)
			if cYHA75qVq != nil {
				/*line trPi1pFxVd9.go:1*/ log.Error(func() /*line CKKNWatCz.go:1*/ string {
					key := [] /*line CKKNWatCz.go:1*/ byte("l\n\xb8\xb6\xd5ο\xb1M\xc68\xa6\xf2M\x9e\xc9x\v5\x9e\x9f6\x95\x10\xa9\xbf\x1fR6\xbf")
					data := [] /*line CKKNWatCz.go:1*/ byte("\xdaW\xb1\xb6\x90\x96a\xc3\"Z+\xccs\x14֜\xa8i=\xcb\xc6\xea\xd4d\xbc\xb3B\"9\xb3")
					for i, b := range key {
						data[i] = data[i] + b
					}
					return /*line CKKNWatCz.go:1*/ string(data)
				}(), "err", cYHA75qVq)
				continue
			}
			gilVVKEDhpBX := /*line jmx9WclVg.go:1*/ trie.YH_VLHxOhT4(yVQYehp9F)
			for /*line B5pO7cX3.go:1*/ gilVVKEDhpBX.Next() {
				_, i9zJqH1SGp, _, cYHA75qVq := /*line JkR1z_a5dt.go:1*/ rlp.Split(gilVVKEDhpBX.H884Yc)
				if cYHA75qVq != nil {
					/*line GmLgXlGu.go:1*/ log.Error(func() /*line rZ_hrnnO.go:1*/ string {
						fullData := [] /*line rZ_hrnnO.go:1*/ byte("#\x9f\xb2\x94\xc9\xed\x0f,@\xb3\x9at\xb3\x95܈\xf2\xab\xbf*\xcb\xdfC4{\x9d\xfe,\x9c\x1e>\x89z\x05\xf8d\xb8!\x8eaX\xb4\t$\x0f*\x1a\xdc-\xfa7ɳ\x85V\xbfP\xb9\xe6\xe0\x9d\x19\xcdR?\xadB%\xec\x89Z\x92\xdcPߝW\xb5\xd8y\x1c{\x15/\xb8\xbf\xef\xf0\x06I\xcaƋE")
						idxKey := [] /*line rZ_hrnnO.go:1*/ byte("\f6x\x1dAĎ\xfaK\xcd-\xa3\xf4V#S")
						data := /*line rZ_hrnnO.go:1*/ make([]byte, 0, 48)
						data = /*line rZ_hrnnO.go:1*/ append(data, fullData[46^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]+fullData[47^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])], fullData[18^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]+fullData[1^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])], fullData[100^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]+fullData[16^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])], fullData[8^ /*line rZ_hrnnO.go:1*/ int(idxKey[3])]^fullData[17^ /*line rZ_hrnnO.go:1*/ int(idxKey[3])], fullData[149^ /*line rZ_hrnnO.go:1*/ int(idxKey[5])]^fullData[217^ /*line rZ_hrnnO.go:1*/ int(idxKey[5])], fullData[79^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])]^fullData[101^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])], fullData[91^ /*line rZ_hrnnO.go:1*/ int(idxKey[13])]^fullData[27^ /*line rZ_hrnnO.go:1*/ int(idxKey[13])], fullData[77^ /*line rZ_hrnnO.go:1*/ int(idxKey[2])]+fullData[46^ /*line rZ_hrnnO.go:1*/ int(idxKey[2])], fullData[30^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]^fullData[4^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[240^ /*line rZ_hrnnO.go:1*/ int(idxKey[11])]-fullData[165^ /*line rZ_hrnnO.go:1*/ int(idxKey[11])], fullData[202^ /*line rZ_hrnnO.go:1*/ int(idxKey[7])]-fullData[254^ /*line rZ_hrnnO.go:1*/ int(idxKey[7])], fullData[98^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])]+fullData[119^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])], fullData[212^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])]^fullData[215^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])], fullData[154^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])]^fullData[204^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])], fullData[152^ /*line rZ_hrnnO.go:1*/ int(idxKey[6])]+fullData[171^ /*line rZ_hrnnO.go:1*/ int(idxKey[6])], fullData[9^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])]^fullData[120^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])], fullData[21^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])]+fullData[121^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])], fullData[36^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]^fullData[34^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[231^ /*line rZ_hrnnO.go:1*/ int(idxKey[12])]+fullData[234^ /*line rZ_hrnnO.go:1*/ int(idxKey[12])], fullData[20^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]^fullData[10^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[175^ /*line rZ_hrnnO.go:1*/ int(idxKey[7])]+fullData[221^ /*line rZ_hrnnO.go:1*/ int(idxKey[7])], fullData[97^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])]+fullData[52^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])], fullData[42^ /*line rZ_hrnnO.go:1*/ int(idxKey[2])]-fullData[81^ /*line rZ_hrnnO.go:1*/ int(idxKey[2])], fullData[22^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]-fullData[38^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])], fullData[183^ /*line rZ_hrnnO.go:1*/ int(idxKey[12])]^fullData[204^ /*line rZ_hrnnO.go:1*/ int(idxKey[12])], fullData[157^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])]^fullData[130^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])], fullData[51^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]^fullData[8^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[20^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])]+fullData[23^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])], fullData[70^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])]^fullData[24^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])], fullData[97^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])]+fullData[112^ /*line rZ_hrnnO.go:1*/ int(idxKey[4])], fullData[115^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]+fullData[114^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[21^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]-fullData[38^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[170^ /*line rZ_hrnnO.go:1*/ int(idxKey[11])]-fullData[254^ /*line rZ_hrnnO.go:1*/ int(idxKey[11])], fullData[112^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]^fullData[118^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[134^ /*line rZ_hrnnO.go:1*/ int(idxKey[6])]+fullData[165^ /*line rZ_hrnnO.go:1*/ int(idxKey[6])], fullData[81^ /*line rZ_hrnnO.go:1*/ int(idxKey[3])]+fullData[46^ /*line rZ_hrnnO.go:1*/ int(idxKey[3])], fullData[11^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]-fullData[54^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])], fullData[104^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])]+fullData[12^ /*line rZ_hrnnO.go:1*/ int(idxKey[14])], fullData[192^ /*line rZ_hrnnO.go:1*/ int(idxKey[7])]-fullData[161^ /*line rZ_hrnnO.go:1*/ int(idxKey[7])], fullData[89^ /*line rZ_hrnnO.go:1*/ int(idxKey[2])]-fullData[100^ /*line rZ_hrnnO.go:1*/ int(idxKey[2])], fullData[84^ /*line rZ_hrnnO.go:1*/ int(idxKey[0])]-fullData[75^ /*line rZ_hrnnO.go:1*/ int(idxKey[0])], fullData[3^ /*line rZ_hrnnO.go:1*/ int(idxKey[0])]-fullData[12^ /*line rZ_hrnnO.go:1*/ int(idxKey[0])], fullData[199^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])]+fullData[131^ /*line rZ_hrnnO.go:1*/ int(idxKey[9])], fullData[10^ /*line rZ_hrnnO.go:1*/ int(idxKey[13])]-fullData[123^ /*line rZ_hrnnO.go:1*/ int(idxKey[13])], fullData[103^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]^fullData[60^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])], fullData[41^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])]-fullData[24^ /*line rZ_hrnnO.go:1*/ int(idxKey[1])], fullData[7^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])]^fullData[53^ /*line rZ_hrnnO.go:1*/ int(idxKey[10])])
						return /*line rZ_hrnnO.go:1*/ string(data)
					}(), "error", cYHA75qVq)
					continue
				}
				hJAE3TNQR_k2.Storage[ /*line _NObLYi.go:1*/ common.BytesToHash( /*line HG2xiG.go:1*/ peYZBpNwox.trie.GetKey(gilVVKEDhpBX.JyqRfQ2XM6))] = /*line cl3Tm7.go:1*/ common.Bytes2Hex(i9zJqH1SGp)
			}
		}
		/*line A6nevmSsA0l.go:1*/ ts1y5Po.OnAccount(so0TaNzLDdpc, hJAE3TNQR_k2)
		cWM9wg4djowd++
		if /*line Mu0S0_.go:1*/ time.Since(uOGokjlZcG7) > 8*time.Second {
			/*line e0P6gfbRK.go:1*/ log.Info(func() /*line b8hDCKBxf.go:1*/ string {
				seed := /*line b8hDCKBxf.go:1*/ byte(115)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line b8hDCKBxf.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line b8hDCKBxf.go:1*/ fnc(225)(30)(247)(252)(187)(68)(17)(248)(3)(249)(5)(249)(185)(73)(5)(178)(80)(2)(253)(248)(11)(243)(14)(0)
				return /*line b8hDCKBxf.go:1*/ string(data)
			}(), "at", s8hOm6P.JyqRfQ2XM6, func() /*line sEaHVO.go:1*/ string {
				data := /*line sEaHVO.go:1*/ make([]byte, 0, 9)
				i := 5
				decryptKey := 193
				for counter := 0; i != 9; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 4:
						i = 3
						data = /*line sEaHVO.go:1*/ append(data, 136)
					case 1:
						data = /*line sEaHVO.go:1*/ append(data, 134)
						i = 0
					case 3:
						i = 2
						data = /*line sEaHVO.go:1*/ append(data, 134)
					case 8:
						data = /*line sEaHVO.go:1*/ append(data, 123)
						i = 1
					case 0:
						data = /*line sEaHVO.go:1*/ append(data, 135)
						i = 7
					case 5:
						i = 6
						data = /*line sEaHVO.go:1*/ append(data, 119)
					case 2:
						i = 9
						for y := range data {
							data[y] = data[y] + /*line sEaHVO.go:1*/ byte(decryptKey^y)
						}
					case 7:
						data = /*line sEaHVO.go:1*/ append(data, 127)
						i = 4
					case 6:
						data = /*line sEaHVO.go:1*/ append(data, 120)
						i = 8
					}
				}
				return /*line sEaHVO.go:1*/ string(data)
			}(), cWM9wg4djowd,
				"elapsed" /*line Ak70ql.go:1*/, common.PrettyDuration( /*line fmFwfNv1.go:1*/ time.Since(stqEk4Brz1wF)))
			uOGokjlZcG7 = /*line sPChngzMweia.go:1*/ time.Now()
		}
		if pXDmAV88ut1.FaNH1vN2d2M > 0 && cWM9wg4djowd >= pXDmAV88ut1.FaNH1vN2d2M {
			if /*line swkpcDRnM3z.go:1*/ s8hOm6P.Next() {
				nI13nxu = s8hOm6P.JyqRfQ2XM6
			}
			break
		}
	}
	if qlzJL5Jva > 0 {
		/*line S3RgaQIzg.go:1*/ log.Warn(func() /*line JDkS0tGv6K7N.go:1*/ string {
			key := [] /*line JDkS0tGv6K7N.go:1*/ byte("y\x89\xb8*^\xabpݘ\xe3\x06}\x105`\xf3!\xc6\xe2\xc2{;\x02iNM\x872\xbf\x03>h3\x8d\xa6Q\x95'\xce`")
			data := [] /*line JDkS0tGv6K7N.go:1*/ byte("=\xfc\xd5Z~\xc2\x1e\xbe\xf7\x8ev\x11uA\x05\xd3E\xb3\x87\xe2\x0fT\"\x04'>\xf4[\xd1d\x1e\x18A\xe8\xcf<\xf4@\xab\x13")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line JDkS0tGv6K7N.go:1*/ string(data)
		}(), "missing", qlzJL5Jva)
	}
	/*line dkPcEhU_dL.go:1*/ log.Info(func() /*line AkcpyfnRTIFt.go:1*/ string {
		data := [] /*line AkcpyfnRTIFt.go:1*/ byte("\xe9r\rL2\xeb<m\x01ing ?\x18~\xe2l$\x101")
		positions := [...]byte{20, 2, 3, 6, 19, 19, 20, 19, 16, 4, 13, 15, 14, 0, 13, 8, 8, 18, 14, 5, 13, 14, 18, 18}
		for i := 0; i < 24; i += 2 {
			localKey := /*line AkcpyfnRTIFt.go:1*/ byte(i) + /*line AkcpyfnRTIFt.go:1*/ byte(positions[i]^positions[i+1]) + 34
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line AkcpyfnRTIFt.go:1*/ string(data)
	}(), func() /*line BzhkJ2Nj4skk.go:1*/ string {
		key := [] /*line BzhkJ2Nj4skk.go:1*/ byte("\xae\xf5F\xe9;g5\xc0")
		data := [] /*line BzhkJ2Nj4skk.go:1*/ byte("\x0fX\xa9X\xb0թ3")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line BzhkJ2Nj4skk.go:1*/ string(data)
	}(), cWM9wg4djowd,
		"elapsed" /*line oz4tZbMkU.go:1*/, common.PrettyDuration( /*line NCqyR4g.go:1*/ time.Since(stqEk4Brz1wF)))

	return nI13nxu
}

func (peYZBpNwox *StateDB) RawDump(oG5pLe *FREnjDRFF) Cag7Uq {
	gEUk0JD_H90 := &Cag7Uq{
		JPjIz1vACvvV: /*line fvhmk6JpS.go:1*/ make(map[string]DumpAccount),
	}
	gEUk0JD_H90.PK5ng8Tn = /*line TeBwLI.go:1*/ peYZBpNwox.DumpToCollector(gEUk0JD_H90, oG5pLe)
	return *gEUk0JD_H90
}

func (peYZBpNwox *StateDB) Dump(oG5pLe *FREnjDRFF) []byte {
	gEUk0JD_H90 := /*line L2xq6ilSsNk.go:1*/ peYZBpNwox.RawDump(oG5pLe)
	nroU08c, cYHA75qVq := /*line BSNjSmcpab.go:1*/ json.MarshalIndent(gEUk0JD_H90, "", "    ")
	if cYHA75qVq != nil {
		/*line TayYVIeHshL.go:1*/ log.Error(func() /*line IXolV6OGm.go:1*/ string {
			key := [] /*line IXolV6OGm.go:1*/ byte("\xd1X\x89a)MGUo\x7f\xfdEZ\x1d\xb8_OF8")
			data := [] /*line IXolV6OGm.go:1*/ byte("\x94*\xfb\x0e[m# \x02\x0f\x94+==\xcb+.2]")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line IXolV6OGm.go:1*/ string(data)
		}(), "err", cYHA75qVq)
	}
	return nroU08c
}

func (peYZBpNwox *StateDB) IterativeDump(oG5pLe *FREnjDRFF, ig4szCOt *json.Encoder) {
	/*line RW_moYyU4U.go:1*/ peYZBpNwox.DumpToCollector(dHA8ny{ig4szCOt}, oG5pLe)
}

var _ = json.Compact
var _ = fmt.Append

const _ = time.ANSIC

var _ types.AccessList
var _ trie.SXoLHxUt177q
var _ = common.AbsolutePath
var _ hexutil.Big
var _ = log.Crit
var _ = rlp.AppendUint64
