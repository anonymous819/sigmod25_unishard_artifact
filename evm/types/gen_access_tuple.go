//line :1
package types

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

func (wp2wJGQyqGR AccessTuple) MarshalJSON() ([]byte, error) {
	type AccessTuple struct {
		Address     common.Address `json:"address"     gencodec:"required"`
		StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
	}
	var aINiWZ_Jtlzj AccessTuple
	aINiWZ_Jtlzj.Address = wp2wJGQyqGR.Address
	aINiWZ_Jtlzj.StorageKeys = wp2wJGQyqGR.StorageKeys
	return /*line NnpoAnVYQmT.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (wp2wJGQyqGR *AccessTuple) UnmarshalJSON(uzD2Up []byte) error {
	type AccessTuple struct {
		Address     *common.Address `json:"address"     gencodec:"required"`
		StorageKeys []common.Hash   `json:"storageKeys" gencodec:"required"`
	}
	var otqLrvn1CD AccessTuple
	if rfteMJju := /*line KaCvBz.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD); rfteMJju != nil {
		return rfteMJju
	}
	if otqLrvn1CD.Address == nil {
		return /*line ishzR9z.go:1*/ errors.New(func() /*line sCU4rkT_9.go:1*/ string {
			seed := /*line sCU4rkT_9.go:1*/ byte(226)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line sCU4rkT_9.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line sCU4rkT_9.go:1*/ fnc(79)(154)(62)(124)(238)(225)(187)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(168)(83)(166)(90)(167)(92)(184)(36)(65)(200)(153)(53)(24)(81)(196)(136)(18)(50)(100)(169)(115)(225)(190)(117)
			return /*line sCU4rkT_9.go:1*/ string(data)
		}())
	}
	wp2wJGQyqGR.Address = *otqLrvn1CD.Address
	if otqLrvn1CD.StorageKeys == nil {
		return /*line r55NEzUT.go:1*/ errors.New(func() /*line Kgbn4Pxc7Ga.go:1*/ string {
			fullData := [] /*line Kgbn4Pxc7Ga.go:1*/ byte("\xb0<\xfe~\xdb4$F\xad#\xeb¶\xcfaV;\xc7 \xab^kg\xb3\xbf\xd0mU،\xb1\xe2\xadϘ:M\xd8sF'\xdd\xef\x9dӰ\x17U\n:\xad\x94\xc2\xcd,\x06\x000f\x92\xd0Pe>R\xb1\xc3vأi|z\xcfx\x0ej\x8c_/\x9b\xdc\xc6\xd7\x1f\xae\xa6\xa9\xcaK6\"¹\x19[\xbc\xf2\t\x05\x94\xbaH\f")
			idxKey := [] /*line Kgbn4Pxc7Ga.go:1*/ byte("+\xfb\x84\x95JQ`\xdc0\xe4<\n\xd4'\xf0'")
			data := /*line Kgbn4Pxc7Ga.go:1*/ make([]byte, 0, 53)
			data = /*line Kgbn4Pxc7Ga.go:1*/ append(data, fullData[98^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])]+fullData[109^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])], fullData[129^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[7])]-fullData[225^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[7])], fullData[213^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])]^fullData[227^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])], fullData[73^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])]+fullData[103^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])], fullData[57^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])]^fullData[99^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])], fullData[112^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[8])]^fullData[49^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[8])], fullData[122^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])]^fullData[80^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])], fullData[31^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])]+fullData[53^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])], fullData[39^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])]^fullData[19^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])], fullData[32^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[11])]+fullData[73^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[11])], fullData[226^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])]-fullData[243^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])], fullData[18^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[11])]+fullData[6^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[11])], fullData[165^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])]-fullData[190^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])], fullData[204^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[1])]-fullData[200^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[1])], fullData[104^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[10])]+fullData[59^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[10])], fullData[105^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])]-fullData[101^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])], fullData[239^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])]+fullData[240^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])], fullData[4^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])]+fullData[17^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])], fullData[159^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])]+fullData[139^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])], fullData[212^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])]^fullData[134^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])], fullData[61^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])]+fullData[72^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])], fullData[213^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[1])]^fullData[221^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[1])], fullData[133^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])]-fullData[180^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])], fullData[162^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])]+fullData[254^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])], fullData[221^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])]-fullData[152^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])], fullData[89^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[5])]+fullData[64^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[5])], fullData[52^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[8])]+fullData[84^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[8])], fullData[243^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])]^fullData[151^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])], fullData[46^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])]^fullData[4^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[0])], fullData[158^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[1])]^fullData[210^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[1])], fullData[155^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[7])]^fullData[130^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[7])], fullData[12^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])]+fullData[85^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])], fullData[104^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])]+fullData[127^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])], fullData[209^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[7])]-fullData[211^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[7])], fullData[196^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])]^fullData[178^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])], fullData[125^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[8])]-fullData[14^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[8])], fullData[168^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])]+fullData[160^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])], fullData[170^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])]-fullData[204^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[14])], fullData[191^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])]-fullData[141^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[2])], fullData[103^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])]-fullData[117^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[4])], fullData[116^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[10])]+fullData[106^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[10])], fullData[70^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[15])]-fullData[102^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[15])], fullData[201^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])]+fullData[135^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])], fullData[42^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[6])]+fullData[106^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[6])], fullData[218^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[3])]-fullData[205^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[3])], fullData[214^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])]-fullData[213^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])], fullData[161^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])]-fullData[221^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[9])], fullData[55^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[6])]-fullData[123^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[6])], fullData[124^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])]-fullData[7^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[13])], fullData[136^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])]+fullData[129^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])], fullData[190^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[3])]+fullData[220^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[3])], fullData[205^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])]-fullData[193^ /*line Kgbn4Pxc7Ga.go:1*/ int(idxKey[12])])
			return /*line Kgbn4Pxc7Ga.go:1*/ string(data)
		}())
	}
	wp2wJGQyqGR.StorageKeys = otqLrvn1CD.StorageKeys
	return nil
}

var _ = json.Compact
var _ = errors.As
var _ = common.AbsolutePath
