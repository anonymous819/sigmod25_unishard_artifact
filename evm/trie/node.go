//line :1
package trie

import (
	"fmt"
	"io"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

var gxALkQpI = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "[17]"}

type node interface {
	tGJzZYYLvK() (hashNode, bool)
	ta85dJ1aZ(lrCYCaYbK rlp.EncoderBuffer)
	tY84mvCKX(string) string
}

type (
	fullNode struct {
		Children [17]node
		flags    nodeFlag
	}
	qUKQUCfTXB struct {
		ANdZYqbzJ1A []byte
		YpmEYGB     node
		d6dipJ      nodeFlag
	}
	hashNode  []byte
	valueNode []byte
)

var e6jPsAvc = /*line ae_J1uGu.go:1*/ valueNode(nil)

func (gnGMaXTuiFeE *fullNode) EncodeRLP(lrCYCaYbK io.Writer) error {
	bqIpQy := /*line DuZw2dm.go:1*/ rlp.NewEncoderBuffer(lrCYCaYbK)
	/*line dkeOeSQ.go:1*/ gnGMaXTuiFeE.ta85dJ1aZ(bqIpQy)
	return /*line hdBPGLtD4.go:1*/ bqIpQy.Flush()
}

func (gnGMaXTuiFeE *fullNode) sBTiL7Ci() *fullNode     { sBTiL7Ci := *gnGMaXTuiFeE; return &sBTiL7Ci }
func (gnGMaXTuiFeE *qUKQUCfTXB) sBTiL7Ci() *qUKQUCfTXB { sBTiL7Ci := *gnGMaXTuiFeE; return &sBTiL7Ci }

type nodeFlag struct {
	hash  hashNode
	dirty bool
}

func (gnGMaXTuiFeE *fullNode) tGJzZYYLvK() (hashNode, bool) {
	return gnGMaXTuiFeE.flags.hash, gnGMaXTuiFeE.flags.dirty
}
func (gnGMaXTuiFeE *qUKQUCfTXB) tGJzZYYLvK() (hashNode, bool) {
	return gnGMaXTuiFeE.d6dipJ.hash, gnGMaXTuiFeE.d6dipJ.dirty
}
func (gnGMaXTuiFeE hashNode) tGJzZYYLvK() (hashNode, bool)  { return nil, true }
func (gnGMaXTuiFeE valueNode) tGJzZYYLvK() (hashNode, bool) { return nil, true }

func (gnGMaXTuiFeE *fullNode) String() string { return /*line xPjEb8j.go:1*/ gnGMaXTuiFeE.tY84mvCKX("") }
func (gnGMaXTuiFeE *qUKQUCfTXB) String() string {
	return /*line In26HsGkllw.go:1*/ gnGMaXTuiFeE.tY84mvCKX("")
}
func (gnGMaXTuiFeE hashNode) String() string { return /*line h5DU6olg.go:1*/ gnGMaXTuiFeE.tY84mvCKX("") }
func (gnGMaXTuiFeE valueNode) String() string {
	return /*line _O2g_5nGI9o.go:1*/ gnGMaXTuiFeE.tY84mvCKX("")
}

func (gnGMaXTuiFeE *fullNode) tY84mvCKX(qyT4sfWJ0WEG string) string {
	piyxg7w := /*line i5kEOD.go:1*/ fmt.Sprintf("[\n%s  ", qyT4sfWJ0WEG)
	for q2u2020KD6, uAjz8NxU_cL := range &gnGMaXTuiFeE.Children {
		if uAjz8NxU_cL == nil {
			piyxg7w += /*line BQRoPx_EORx.go:1*/ fmt.Sprintf(func() /*line qHRILa_ZLh.go:1*/ string {
				seed := /*line qHRILa_ZLh.go:1*/ byte(135)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line qHRILa_ZLh.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line qHRILa_ZLh.go:1*/ fnc(158)(78)(199)(230)(28)(50)(251)(3)(210)(226)
				return /*line qHRILa_ZLh.go:1*/ string(data)
			}(), gxALkQpI[q2u2020KD6])
		} else {
			piyxg7w += /*line m1D2bcEFkbJ.go:1*/ fmt.Sprintf("%s: %v", gxALkQpI[q2u2020KD6] /*line A_3bLvxlibM7.go:1*/, uAjz8NxU_cL.tY84mvCKX(qyT4sfWJ0WEG+"  "))
		}
	}
	return piyxg7w + /*line MrBzKPI49FiG.go:1*/ fmt.Sprintf("\n%s] ", qyT4sfWJ0WEG)
}
func (gnGMaXTuiFeE *qUKQUCfTXB) tY84mvCKX(qyT4sfWJ0WEG string) string {
	return /*line o7kR5n.go:1*/ fmt.Sprintf(func() /*line xxvdCwjLyk.go:1*/ string {
		fullData := [] /*line xxvdCwjLyk.go:1*/ byte("\xf0\xac\xde\xe0*\x9d\xea)\xf7;i\x1f1\xbe8d\x1c]")
		idxKey := [] /*line xxvdCwjLyk.go:1*/ byte("\xd7\xebq-\xf3")
		data := /*line xxvdCwjLyk.go:1*/ make([]byte, 0, 10)
		data = /*line xxvdCwjLyk.go:1*/ append(data, fullData[214^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])]-fullData[219^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])], fullData[120^ /*line xxvdCwjLyk.go:1*/ int(idxKey[2])]+fullData[119^ /*line xxvdCwjLyk.go:1*/ int(idxKey[2])], fullData[199^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])]^fullData[216^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])], fullData[239^ /*line xxvdCwjLyk.go:1*/ int(idxKey[1])]-fullData[235^ /*line xxvdCwjLyk.go:1*/ int(idxKey[1])], fullData[213^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])]-fullData[218^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])], fullData[198^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])]-fullData[217^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])], fullData[248^ /*line xxvdCwjLyk.go:1*/ int(idxKey[4])]^fullData[249^ /*line xxvdCwjLyk.go:1*/ int(idxKey[4])], fullData[210^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])]^fullData[212^ /*line xxvdCwjLyk.go:1*/ int(idxKey[0])], fullData[121^ /*line xxvdCwjLyk.go:1*/ int(idxKey[2])]+fullData[118^ /*line xxvdCwjLyk.go:1*/ int(idxKey[2])])
		return /*line xxvdCwjLyk.go:1*/ string(data)
	}(), gnGMaXTuiFeE.ANdZYqbzJ1A /*line wr691CIzG.go:1*/, gnGMaXTuiFeE.YpmEYGB.tY84mvCKX(qyT4sfWJ0WEG+"  "))
}
func (gnGMaXTuiFeE hashNode) tY84mvCKX(qyT4sfWJ0WEG string) string {
	return /*line rlETLQlJyhV0.go:1*/ fmt.Sprintf("<%x> ", [] /*line NFLusp.go:1*/ byte(gnGMaXTuiFeE))
}
func (gnGMaXTuiFeE valueNode) tY84mvCKX(qyT4sfWJ0WEG string) string {
	return /*line aOfvSItEI.go:1*/ fmt.Sprintf("%x ", [] /*line zBsELVN.go:1*/ byte(gnGMaXTuiFeE))
}

type rm74ZU64 []byte

func (gnGMaXTuiFeE rm74ZU64) tGJzZYYLvK() (hashNode, bool) {
	/*line COB4uBBI9.go:1*/ panic(func() /*line rYyafaE.go:1*/ string {
		key := [] /*line rYyafaE.go:1*/ byte("\xd3D(\x9e\x93Vb\x9c\x99?h5\n\x05\xe2\xa9\xea4`\x96\xd9d\x9cDl\x97\xf1\x8c\xa2J\xfeH\x83\x86\x99O\x9e\xa2\xbb")
		data := [] /*line rYyafaE.go:1*/ byte("\xa7,A\xed\xb3%\n\xf3\xecS\f\x15d`\x94̘\x14\x05\xf8\xbdD\xe94L\xfe\x9f\xac\xc3j\x92!\xf5\xe3\xb9;\xec\xcb\xde")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line rYyafaE.go:1*/ string(data)
	}())
}
func (gnGMaXTuiFeE rm74ZU64) tY84mvCKX(qyT4sfWJ0WEG string) string {
	/*line XvI5coSsHZ.go:1*/ panic(func() /*line C7VBo9CIo9vk.go:1*/ string {
		fullData := [] /*line C7VBo9CIo9vk.go:1*/ byte("\xf6\"j\x98\a\xdd\b?UW\b9%\x10\xb0#XI\xa3\xda5P\xee\x03\xba\xabX\xecG\x8f\xf4(pf\xaf\xd0-b\x8c\xdc~_\xdeG\xc4\xf8\f+a\f÷-J\xfe\x93\x86O\xaadÉ)\x04\xf8h\x88\x1dc\x9a\xec\x84\xe5\x9f\x18U\x04\xd9")
		idxKey := [] /*line C7VBo9CIo9vk.go:1*/ byte("\x8c\x8a\xccS@\x81\xa8WD\x9c\xfd\xb5\xe5tK\x87")
		data := /*line C7VBo9CIo9vk.go:1*/ make([]byte, 0, 40)
		data = /*line C7VBo9CIo9vk.go:1*/ append(data, fullData[0^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])]^fullData[102^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])], fullData[185^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])]-fullData[179^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])], fullData[231^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])]+fullData[205^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])], fullData[146^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[6])]^fullData[229^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[6])], fullData[193^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])]-fullData[239^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])], fullData[207^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])]+fullData[243^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])], fullData[179^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[0])]+fullData[183^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[0])], fullData[112^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[7])]+fullData[96^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[7])], fullData[241^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])]+fullData[138^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])], fullData[168^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[9])]+fullData[155^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[9])], fullData[226^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])]+fullData[220^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])], fullData[84^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[13])]-fullData[97^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[13])], fullData[75^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])]^fullData[73^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])], fullData[133^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[11])]^fullData[249^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[11])], fullData[213^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])]-fullData[251^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])], fullData[15^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[8])]+fullData[73^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[8])], fullData[8^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[14])]-fullData[82^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[14])], fullData[128^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[11])]^fullData[183^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[11])], fullData[84^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])]-fullData[99^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])], fullData[126^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])]^fullData[92^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])], fullData[144^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])]-fullData[148^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])], fullData[165^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[15])]^fullData[154^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[15])], fullData[15^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[14])]-fullData[93^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[14])], fullData[164^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[12])]-fullData[200^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[12])], fullData[214^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[9])]+fullData[150^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[9])], fullData[217^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])]-fullData[209^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[10])], fullData[152^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[15])]-fullData[159^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[15])], fullData[144^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[6])]+fullData[237^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[6])], fullData[85^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[13])]^fullData[112^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[13])], fullData[170^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[9])]^fullData[182^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[9])], fullData[81^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])]^fullData[76^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[4])], fullData[163^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])]-fullData[138^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])], fullData[75^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[8])]^fullData[76^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[8])], fullData[197^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[15])]+fullData[130^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[15])], fullData[189^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[0])]-fullData[151^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[0])], fullData[133^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])]-fullData[227^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[2])], fullData[153^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])]+fullData[137^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[1])], fullData[242^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[11])]+fullData[253^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[11])], fullData[81^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[13])]+fullData[99^ /*line C7VBo9CIo9vk.go:1*/ int(idxKey[13])])
		return /*line C7VBo9CIo9vk.go:1*/ string(data)
	}())
}

func (gnGMaXTuiFeE rm74ZU64) EncodeRLP(lrCYCaYbK io.Writer) error {
	_, eZzE0KPS := /*line C8b5CBcNI0F.go:1*/ lrCYCaYbK.Write(gnGMaXTuiFeE)
	return eZzE0KPS
}

func nyHCCZEA(rNuN0sMPDJ, m2t4GBUGH []byte) node {
	gnGMaXTuiFeE, eZzE0KPS := /*line ViM_g2.go:1*/ rRV2VeUAEp(rNuN0sMPDJ, m2t4GBUGH)
	if eZzE0KPS != nil {
		/*line h_I9_eaywB.go:1*/ panic( /*line e0AoVEW__.go:1*/ fmt.Sprintf(func() /*line OOZadW7.go:1*/ string {
			key := [] /*line OOZadW7.go:1*/ byte("B\xcc\x00\xfe\xc0\x9c\x9c\x8d.z\xa8")
			data := [] /*line OOZadW7.go:1*/ byte("\xb0;dc\xe0\xc1\x14\xc7N\x9f\x1e")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return /*line OOZadW7.go:1*/ string(data)
		}(), rNuN0sMPDJ, eZzE0KPS))
	}
	return gnGMaXTuiFeE
}

func hct27E(rNuN0sMPDJ, m2t4GBUGH []byte) node {
	gnGMaXTuiFeE, eZzE0KPS := /*line XKITzk.go:1*/ dzYqgm0ahC(rNuN0sMPDJ, m2t4GBUGH)
	if eZzE0KPS != nil {
		/*line hgB1Ofe.go:1*/ panic( /*line Co2dbu4GH.go:1*/ fmt.Sprintf(func() /*line VbHojnTEQd.go:1*/ string {
			data := /*line VbHojnTEQd.go:1*/ make([]byte, 0, 12)
			i := 4
			decryptKey := 86
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 7
					data = /*line VbHojnTEQd.go:1*/ append(data, 106)
				case 3:
					data = /*line VbHojnTEQd.go:1*/ append(data, 102)
					i = 12
				case 8:
					i = 1
					data = /*line VbHojnTEQd.go:1*/ append(data, 164)
				case 6:
					data = /*line VbHojnTEQd.go:1*/ append(data, 188)
					i = 0
				case 9:
					i = 5
					for y := range data {
						data[y] = data[y] - /*line VbHojnTEQd.go:1*/ byte(decryptKey^y)
					}
				case 0:
					i = 10
					data = /*line VbHojnTEQd.go:1*/ append(data, 127)
				case 2:
					i = 9
					data = /*line VbHojnTEQd.go:1*/ append(data, 190)
				case 7:
					i = 2
					data = /*line VbHojnTEQd.go:1*/ append(data, 112)
				case 4:
					i = 11
					data = /*line VbHojnTEQd.go:1*/ append(data, 176)
				case 12:
					i = 6
					data = /*line VbHojnTEQd.go:1*/ append(data, 108)
				case 1:
					data = /*line VbHojnTEQd.go:1*/ append(data, 166)
					i = 3
				case 11:
					i = 8
					data = /*line VbHojnTEQd.go:1*/ append(data, 178)
				}
			}
			return /*line VbHojnTEQd.go:1*/ string(data)
		}(), rNuN0sMPDJ, eZzE0KPS))
	}
	return gnGMaXTuiFeE
}

func rRV2VeUAEp(rNuN0sMPDJ, m2t4GBUGH []byte) (node, error) {
	return /*line IsEwUEQGHkEX.go:1*/ dzYqgm0ahC(rNuN0sMPDJ /*line RZmsxXVM6.go:1*/, common.CopyBytes(m2t4GBUGH))
}

func dzYqgm0ahC(rNuN0sMPDJ, m2t4GBUGH []byte) (node, error) {
	if /*line dJqoZHn0JF.go:1*/ len(m2t4GBUGH) == 0 {
		return nil, io.ErrUnexpectedEOF
	}
	hmLv2XhaDagw, _, eZzE0KPS := /*line GdfL3lLkD.go:1*/ rlp.SplitList(m2t4GBUGH)
	if eZzE0KPS != nil {
		return nil /*line E2g6sUs.go:1*/, fmt.Errorf(func() /*line EpF_NAVR.go:1*/ string {
			data := [] /*line EpF_NAVR.go:1*/ byte("\xb4\xe8\xe5\xf3d\xe9 erm;rq}\xa5`")
			positions := [...]byte{9, 2, 14, 3, 10, 3, 14, 15, 0, 12, 9, 0, 5, 2, 13, 2, 14, 10, 0, 2, 13, 3, 1, 15}
			for i := 0; i < 24; i += 2 {
				localKey := /*line EpF_NAVR.go:1*/ byte(i) + /*line EpF_NAVR.go:1*/ byte(positions[i]^positions[i+1]) + 122
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line EpF_NAVR.go:1*/ string(data)
		}(), eZzE0KPS)
	}
	switch arf3iiVz, _ := /*line g7W3MZZ.go:1*/ rlp.CountValues(hmLv2XhaDagw); arf3iiVz {
	case 2:
		gnGMaXTuiFeE, eZzE0KPS := /*line HSFgndCfy5FB.go:1*/ iqliRWY(rNuN0sMPDJ, hmLv2XhaDagw)
		return gnGMaXTuiFeE /*line Px_4nqD7w.go:1*/, dnDqoJmvhyY(eZzE0KPS, "short")
	case 17:
		gnGMaXTuiFeE, eZzE0KPS := /*line zZTxWlElA.go:1*/ z1hh7fA(rNuN0sMPDJ, hmLv2XhaDagw)
		return gnGMaXTuiFeE /*line C0GR7RDyaJ.go:1*/, dnDqoJmvhyY(eZzE0KPS, "full")
	default:
		return nil /*line GzEZ6oaBY2dv.go:1*/, fmt.Errorf(func() /*line xfVUdv0.go:1*/ string {
			data := [] /*line xfVUdv0.go:1*/ byte("i^\xd1\xedl\xf9X\xc1\x88Em\x85U\xa5\xaaKf \xf5ist\xace!em\x91q^\xb5\xe9 \x95d")
			positions := [...]byte{2, 7, 28, 9, 6, 28, 3, 22, 7, 12, 24, 1, 12, 30, 15, 6, 31, 18, 13, 5, 9, 13, 27, 34, 7, 29, 33, 15, 7, 1, 33, 18, 28, 6, 8, 34, 24, 11, 9, 14}
			for i := 0; i < 40; i += 2 {
				localKey := /*line xfVUdv0.go:1*/ byte(i) + /*line xfVUdv0.go:1*/ byte(positions[i]^positions[i+1]) + 178
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line xfVUdv0.go:1*/ string(data)
		}(), arf3iiVz)
	}
}

func iqliRWY(rNuN0sMPDJ, hmLv2XhaDagw []byte) (node, error) {
	hhojRHYKMx, ehAc5MS, eZzE0KPS := /*line vJjn206j_xb.go:1*/ rlp.SplitString(hmLv2XhaDagw)
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	h3c6blF := nodeFlag{hash: rNuN0sMPDJ}
	lhQWH7m := /*line EHKrLsA.go:1*/ o9hF7yS(hhojRHYKMx)
	if /*line C5W86DlvQK3.go:1*/ k7wsvmNwK(lhQWH7m) {

		h_pxYek4zT, _, eZzE0KPS := /*line RT0VTv.go:1*/ rlp.SplitString(ehAc5MS)
		if eZzE0KPS != nil {
			return nil /*line eMFcwwF.go:1*/, fmt.Errorf(func() /*line uuWntpI.go:1*/ string {
				key := [] /*line uuWntpI.go:1*/ byte("U\x9e^R\xc8\xf7\x82\xb1`\xc9k\uf70a\x19\x97\u0605bޢ\x0e")
				data := [] /*line uuWntpI.go:1*/ byte("\xbe\fԳ4`\xe6\xd1\xd6*\xd7d\x01\xaa\x87\x06<\xea\x9c\xfeǄ")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return /*line uuWntpI.go:1*/ string(data)
			}(), eZzE0KPS)
		}
		return &qUKQUCfTXB{lhQWH7m /*line YMIKSO8zWE.go:1*/, valueNode(h_pxYek4zT), h3c6blF}, nil
	}
	iMFHMG, _, eZzE0KPS := /*line i2MAqCGE.go:1*/ qhe4Hx0d(ehAc5MS)
	if eZzE0KPS != nil {
		return nil /*line bAQ6wK_.go:1*/, dnDqoJmvhyY(eZzE0KPS, "val")
	}
	return &qUKQUCfTXB{lhQWH7m, iMFHMG, h3c6blF}, nil
}

func z1hh7fA(rNuN0sMPDJ, hmLv2XhaDagw []byte) (*fullNode, error) {
	gnGMaXTuiFeE := &fullNode{flags: nodeFlag{hash: rNuN0sMPDJ}}
	for q2u2020KD6 := 0; q2u2020KD6 < 16; q2u2020KD6++ {
		nPjnWCe, ehAc5MS, eZzE0KPS := /*line GSPqbgAHx0v.go:1*/ qhe4Hx0d(hmLv2XhaDagw)
		if eZzE0KPS != nil {
			return gnGMaXTuiFeE /*line UftObb.go:1*/, dnDqoJmvhyY(eZzE0KPS /*line FNZM6k7J72.go:1*/, fmt.Sprintf("[%d]", q2u2020KD6))
		}
		gnGMaXTuiFeE.Children[q2u2020KD6], hmLv2XhaDagw = nPjnWCe, ehAc5MS
	}
	h_pxYek4zT, _, eZzE0KPS := /*line ZzlHatQhm.go:1*/ rlp.SplitString(hmLv2XhaDagw)
	if eZzE0KPS != nil {
		return gnGMaXTuiFeE, eZzE0KPS
	}
	if /*line HpXmQvi.go:1*/ len(h_pxYek4zT) > 0 {
		gnGMaXTuiFeE.Children[16] = /*line e5a3RKaXu4X.go:1*/ valueNode(h_pxYek4zT)
	}
	return gnGMaXTuiFeE, nil
}

const hashLen = /*line FUv8JqZU.go:1*/ len(common.Hash{})

func qhe4Hx0d(m2t4GBUGH []byte) (node, []byte, error) {
	lQi1ROE, h_pxYek4zT, ehAc5MS, eZzE0KPS := /*line zlXPhgOZ7g8.go:1*/ rlp.Split(m2t4GBUGH)
	if eZzE0KPS != nil {
		return nil, m2t4GBUGH, eZzE0KPS
	}
	switch {
	case lQi1ROE == rlp.List:

		if fa8lJCa := /*line HBXnF9zlW1cl.go:1*/ len(m2t4GBUGH) - /*line bDPiygbzaT.go:1*/ len(ehAc5MS); fa8lJCa > hashLen {
			eZzE0KPS := /*line C_9Ue4.go:1*/ fmt.Errorf(func() /*line NBZxsXyM.go:1*/ string {
				data := /*line NBZxsXyM.go:1*/ make([]byte, 0, 59)
				i := 0
				decryptKey := 171
				for counter := 0; i != 17; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 10:
						i = 1
						data = /*line NBZxsXyM.go:1*/ append(data, 13)
					case 15:
						i = 3
						data = /*line NBZxsXyM.go:1*/ append(data, 239)
					case 22:
						data = /*line NBZxsXyM.go:1*/ append(data, "][\x16"...,
						)
						i = 16
					case 12:
						i = 13
						data = /*line NBZxsXyM.go:1*/ append(data, 69)
					case 19:
						i = 22
						data = /*line NBZxsXyM.go:1*/ append(data, "cs"...,
						)
					case 18:
						data = /*line NBZxsXyM.go:1*/ append(data, "ZNN\b"...,
						)
						i = 14
					case 8:
						i = 21
						data = /*line NBZxsXyM.go:1*/ append(data, "7G1\xeb"...,
						)
					case 21:
						i = 15
						data = /*line NBZxsXyM.go:1*/ append(data, "\x06\xe9\xed+"...,
						)
					case 0:
						data = /*line NBZxsXyM.go:1*/ append(data, 110)
						i = 2
					case 14:
						data = /*line NBZxsXyM.go:1*/ append(data, 15)
						i = 5
					case 4:
						i = 6
						data = /*line NBZxsXyM.go:1*/ append(data, "\x02\xf5"...,
						)
					case 27:
						i = 26
						data = /*line NBZxsXyM.go:1*/ append(data, "^H"...,
						)
					case 9:
						data = /*line NBZxsXyM.go:1*/ append(data, 82)
						i = 10
					case 25:
						i = 24
						data = /*line NBZxsXyM.go:1*/ append(data, "S\xff"...,
						)
					case 1:
						data = /*line NBZxsXyM.go:1*/ append(data, 90)
						i = 18
					case 6:
						data = /*line NBZxsXyM.go:1*/ append(data, "K4"...,
						)
						i = 20
					case 7:
						data = /*line NBZxsXyM.go:1*/ append(data, "UTT"...,
						)
						i = 9
					case 5:
						i = 27
						data = /*line NBZxsXyM.go:1*/ append(data, "YN"...,
						)
					case 16:
						data = /*line NBZxsXyM.go:1*/ append(data, "ZaUW"...,
						)
						i = 7
					case 20:
						data = /*line NBZxsXyM.go:1*/ append(data, 64)
						i = 12
					case 26:
						data = /*line NBZxsXyM.go:1*/ append(data, "\x02J"...,
						)
						i = 25
					case 13:
						data = /*line NBZxsXyM.go:1*/ append(data, "\xf0B"...,
						)
						i = 8
					case 3:
						for y := range data {
							data[y] = data[y] - /*line NBZxsXyM.go:1*/ byte(decryptKey^y)
						}
						i = 17
					case 28:
						data = /*line NBZxsXyM.go:1*/ append(data, "=S"...,
						)
						i = 23
					case 11:
						data = /*line NBZxsXyM.go:1*/ append(data, "bnn"...,
						)
						i = 19
					case 2:
						i = 11
						data = /*line NBZxsXyM.go:1*/ append(data, 116)
					case 23:
						i = 4
						data = /*line NBZxsXyM.go:1*/ append(data, "M=J"...,
						)
					case 24:
						data = /*line NBZxsXyM.go:1*/ append(data, "\x03A\xfc"...,
						)
						i = 28
					}
				}
				return /*line NBZxsXyM.go:1*/ string(data)
			}(), fa8lJCa, hashLen)
			return nil, m2t4GBUGH, eZzE0KPS
		}
		gnGMaXTuiFeE, eZzE0KPS := /*line BR8WQc7Yh7.go:1*/ rRV2VeUAEp(nil, m2t4GBUGH)
		return gnGMaXTuiFeE, ehAc5MS, eZzE0KPS
	case lQi1ROE == rlp.String && /*line qDI6CjF.go:1*/ len(h_pxYek4zT) == 0:

		return nil, ehAc5MS, nil
	case lQi1ROE == rlp.String && /*line BxakMYI99e.go:1*/ len(h_pxYek4zT) == 32:
		return /*line S3x_CN.go:1*/ hashNode(h_pxYek4zT), ehAc5MS, nil
	default:
		return nil, nil /*line WnbEGG7G.go:1*/, fmt.Errorf(func() /*line QXDa4yzIo_t.go:1*/ string {
			data := /*line QXDa4yzIo_t.go:1*/ make([]byte, 0, 42)
			i := 2
			decryptKey := 98
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 16:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, 33)
					i = 14
				case 14:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, 101)
					i = 7
				case 7:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "7*"...,
					)
					i = 20
				case 10:
					i = 19
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "FU\x1b"...,
					)
				case 20:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "8$"...,
					)
					i = 3
				case 9:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "\x12\r"...,
					)
					i = 15
				case 15:
					i = 11
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "|(."...,
					)
				case 17:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "5p\r"...,
					)
					i = 9
				case 19:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "\x01RB"...,
					)
					i = 13
				case 1:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "+'<W"...,
					)
					i = 10
				case 3:
					i = 18
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "`j*"...,
					)
				case 8:
					i = 0
					for y := range data {
						data[y] = data[y] ^ /*line QXDa4yzIo_t.go:1*/ byte(decryptKey^y)
					}
				case 13:
					i = 5
					data = /*line QXDa4yzIo_t.go:1*/ append(data, 66)
				case 2:
					i = 4
					data = /*line QXDa4yzIo_t.go:1*/ append(data, 62)
				case 5:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, 86)
					i = 8
				case 6:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, 41)
					i = 16
				case 4:
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "8#"...,
					)
					i = 12
				case 18:
					i = 1
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "md<"...,
					)
				case 11:
					i = 6
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "+1"...,
					)
				case 12:
					i = 17
					data = /*line QXDa4yzIo_t.go:1*/ append(data, "5?;"...,
					)
				}
			}
			return /*line QXDa4yzIo_t.go:1*/ string(data)
		}(), /*line YgH1paglw.go:1*/ len(h_pxYek4zT))
	}
}

type ubKagYChj struct {
	vlECwnwN0z   error
	tZd9F7sZMvrC []string
}

func dnDqoJmvhyY(eZzE0KPS error, k8Pg6OwvP string) error {
	if eZzE0KPS == nil {
		return nil
	}
	if hOKdGMlN3MD, yY_yPSd8vG := eZzE0KPS.(*ubKagYChj); yY_yPSd8vG {
		hOKdGMlN3MD.tZd9F7sZMvrC = /*line nX6Ny7Ym.go:1*/ append(hOKdGMlN3MD.tZd9F7sZMvrC, k8Pg6OwvP)
		return hOKdGMlN3MD
	}
	return &ubKagYChj{eZzE0KPS, []string{k8Pg6OwvP}}
}

func (eZzE0KPS *ubKagYChj) Error() string {
	return /*line G_RN0afB5.go:1*/ fmt.Sprintf(func() /*line KtC7e7CIWNN.go:1*/ string {
		fullData := [] /*line KtC7e7CIWNN.go:1*/ byte("\xc5\fAb|\x9a\xacD\x19\x04k\xfa\r[Gb\x11\xef\x85\xdf}AX\x8c!\xb2\x9f\xcf$\xbf\xe9\x9c\xccBz@x#\xbc\xd8")
		idxKey := [] /*line KtC7e7CIWNN.go:1*/ byte("\rQ")
		data := /*line KtC7e7CIWNN.go:1*/ make([]byte, 0, 21)
		data = /*line KtC7e7CIWNN.go:1*/ append(data, fullData[79^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[113^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[90^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[70^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[15^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]+fullData[30^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[84^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[72^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[89^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[69^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[73^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]+fullData[86^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[87^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[74^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[91^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[88^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[42^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]^fullData[43^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[17^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]+fullData[24^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[76^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]^fullData[75^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[85^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]-fullData[80^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[40^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]^fullData[44^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[28^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]+fullData[31^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[1^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]+fullData[0^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[115^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]-fullData[114^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[117^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]-fullData[71^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[3^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]^fullData[14^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])], fullData[94^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])]+fullData[65^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[1])], fullData[13^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])]-fullData[18^ /*line KtC7e7CIWNN.go:1*/ int(idxKey[0])])
		return /*line KtC7e7CIWNN.go:1*/ string(data)
	}(), eZzE0KPS.vlECwnwN0z /*line XWB8N1bH.go:1*/, strings.Join(eZzE0KPS.tZd9F7sZMvrC, "<-"))
}

var _ = fmt.Append
var _ io.ByteReader
var _ strings.Builder
var _ = common.AbsolutePath
var _ = rlp.AppendUint64
