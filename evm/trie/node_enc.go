//line :1
package trie

import (
	"github.com/ethereum/go-ethereum/rlp"
)

func kG18lbsZr(gnGMaXTuiFeE node) []byte {
	lrCYCaYbK := /*line GlFiq1lM.go:1*/ rlp.NewEncoderBuffer(nil)
	/*line P9KxUytPBa.go:1*/ gnGMaXTuiFeE.ta85dJ1aZ(lrCYCaYbK)
	gXFlfTTwq := /*line BhLI_tnp.go:1*/ lrCYCaYbK.ToBytes()
	/*line P4jB16MVDDet.go:1*/ lrCYCaYbK.Flush()
	return gXFlfTTwq
}

func (gnGMaXTuiFeE *fullNode) ta85dJ1aZ(lrCYCaYbK rlp.EncoderBuffer) {
	i8Bz8acGb := /*line R6quimFdMcc.go:1*/ lrCYCaYbK.List()
	for _, arf3iiVz := range gnGMaXTuiFeE.Children {
		if arf3iiVz != nil {
			/*line DaihJm.go:1*/ arf3iiVz.ta85dJ1aZ(lrCYCaYbK)
		} else {
			/*line HhalT0KH3.go:1*/ lrCYCaYbK.Write(rlp.EmptyString)
		}
	}
	/*line op3S4hSNXl.go:1*/ lrCYCaYbK.ListEnd(i8Bz8acGb)
}

func (gnGMaXTuiFeE *qUKQUCfTXB) ta85dJ1aZ(lrCYCaYbK rlp.EncoderBuffer) {
	i8Bz8acGb := /*line iEvx7CDi.go:1*/ lrCYCaYbK.List()
	/*line OZuBRqMD6_.go:1*/ lrCYCaYbK.WriteBytes(gnGMaXTuiFeE.ANdZYqbzJ1A)
	if gnGMaXTuiFeE.YpmEYGB != nil {
		/*line F9xbEt67lO.go:1*/ gnGMaXTuiFeE.YpmEYGB.ta85dJ1aZ(lrCYCaYbK)
	} else {
		/*line rT_C8_b.go:1*/ lrCYCaYbK.Write(rlp.EmptyString)
	}
	/*line Gg5ascxEI7.go:1*/ lrCYCaYbK.ListEnd(i8Bz8acGb)
}

func (gnGMaXTuiFeE hashNode) ta85dJ1aZ(lrCYCaYbK rlp.EncoderBuffer) {
	/*line PL6eWJg.go:1*/ lrCYCaYbK.WriteBytes(gnGMaXTuiFeE)
}

func (gnGMaXTuiFeE valueNode) ta85dJ1aZ(lrCYCaYbK rlp.EncoderBuffer) {
	/*line YGsawfrCT.go:1*/ lrCYCaYbK.WriteBytes(gnGMaXTuiFeE)
}

func (gnGMaXTuiFeE rm74ZU64) ta85dJ1aZ(lrCYCaYbK rlp.EncoderBuffer) {
	/*line WJuJbX.go:1*/ lrCYCaYbK.Write(gnGMaXTuiFeE)
}

var _ = rlp.AppendUint64
