//line :1
package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Ck0iTH interface {
	Address() common.Address
}

type AccountRef common.Address

func (ofmoCVtQ0KZ AccountRef) Address() common.Address {
	return ( /*line TDsY8QZAwVu.go:1*/ common.Address)(ofmoCVtQ0KZ)
}

type WyJ004I8 struct {
	HDwb4iNVhYJ1 common.Address
	nEhOfBstgMa  Ck0iTH
	aYunfbcy     Ck0iTH

	s8_ItraG     map[common.Hash]gYd8QoJh3
	fuFjqE7z1wS2 gYd8QoJh3

	STabmC       []byte
	YfrpaFi2iY3  common.Hash
	DkqqBAVFZj9n *common.Address
	AdmClN5HlL   []byte

	OzrTJ12T  uint64
	mddAyugJc *uint256.Int
}

func Jd5jjadjJuy(dcewTNVj Ck0iTH, agaO1kq6 Ck0iTH, gVUwww *uint256.Int, a0SqxOtD uint64) *WyJ004I8 {
	oPtEQcrX := &WyJ004I8{HDwb4iNVhYJ1: /*line ESf0PVh.go:1*/ dcewTNVj.Address(), nEhOfBstgMa: dcewTNVj, aYunfbcy: agaO1kq6}

	if fzrwnz96pkv, m5pPAO := dcewTNVj.(*WyJ004I8); m5pPAO {

		oPtEQcrX.s8_ItraG = fzrwnz96pkv.s8_ItraG
	} else {
		oPtEQcrX.s8_ItraG = /*line Qpkw8rEsh.go:1*/ make(map[common.Hash]gYd8QoJh3)
	}

	oPtEQcrX.OzrTJ12T = a0SqxOtD

	oPtEQcrX.mddAyugJc = gVUwww

	return oPtEQcrX
}

func (oPtEQcrX *WyJ004I8) k_hb9giFj(diQ3lfnTEsj *uint256.Int) bool {
	epzuYYusHWT, r9uWm2pY := /*line rWeJTCWLtGOW.go:1*/ diQ3lfnTEsj.Uint64WithOverflow()

	if r9uWm2pY || epzuYYusHWT >= /*line S1Rhn6Qt.go:1*/ uint64( /*line ZYV0j8Z.go:1*/ len(oPtEQcrX.STabmC)) {
		return false
	}

	if /*line OkOi4HjnQA8.go:1*/ LxosJe8(oPtEQcrX.STabmC[epzuYYusHWT]) != JUMPDEST {
		return false
	}
	return /*line KenwCt.go:1*/ oPtEQcrX.dP9Ui9Qa(epzuYYusHWT)
}

func (oPtEQcrX *WyJ004I8) dP9Ui9Qa(epzuYYusHWT uint64) bool {

	if oPtEQcrX.fuFjqE7z1wS2 != nil {
		return /*line n3axdUGy.go:1*/ oPtEQcrX.fuFjqE7z1wS2.uR1Fotw(epzuYYusHWT)
	}

	if oPtEQcrX.YfrpaFi2iY3 != (common.Hash{}) {

		cPu2Qp4RUGO, fqiXg9 := oPtEQcrX.s8_ItraG[oPtEQcrX.YfrpaFi2iY3]
		if !fqiXg9 {

			cPu2Qp4RUGO = /*line d813QRgE3.go:1*/ t27hPJQn(oPtEQcrX.STabmC)
			oPtEQcrX.s8_ItraG[oPtEQcrX.YfrpaFi2iY3] = cPu2Qp4RUGO
		}

		oPtEQcrX.fuFjqE7z1wS2 = cPu2Qp4RUGO
		return /*line EynG4556.go:1*/ cPu2Qp4RUGO.uR1Fotw(epzuYYusHWT)
	}

	if oPtEQcrX.fuFjqE7z1wS2 == nil {
		oPtEQcrX.fuFjqE7z1wS2 = /*line XKTQL81.go:1*/ t27hPJQn(oPtEQcrX.STabmC)
	}
	return /*line ms3x9zqGvB.go:1*/ oPtEQcrX.fuFjqE7z1wS2.uR1Fotw(epzuYYusHWT)
}

func (oPtEQcrX *WyJ004I8) AsDelegate() *WyJ004I8 {

	fzrwnz96pkv := oPtEQcrX.nEhOfBstgMa.(*WyJ004I8)
	oPtEQcrX.HDwb4iNVhYJ1 = fzrwnz96pkv.HDwb4iNVhYJ1
	oPtEQcrX.mddAyugJc = fzrwnz96pkv.mddAyugJc

	return oPtEQcrX
}

func (oPtEQcrX *WyJ004I8) GetOp(iyrnVpBe8CCl uint64) LxosJe8 {
	if iyrnVpBe8CCl < /*line T_6ALU.go:1*/ uint64( /*line VIvfnNEfqchw.go:1*/ len(oPtEQcrX.STabmC)) {
		return /*line sJaWVl17bt.go:1*/ LxosJe8(oPtEQcrX.STabmC[iyrnVpBe8CCl])
	}

	return STOP
}

func (oPtEQcrX *WyJ004I8) Caller() common.Address {
	return oPtEQcrX.HDwb4iNVhYJ1
}

func (oPtEQcrX *WyJ004I8) UseGas(a0SqxOtD uint64) (m5pPAO bool) {
	if oPtEQcrX.OzrTJ12T < a0SqxOtD {
		return false
	}
	oPtEQcrX.OzrTJ12T -= a0SqxOtD
	return true
}

func (oPtEQcrX *WyJ004I8) Address() common.Address {
	return /*line iAAj0amWaT3.go:1*/ oPtEQcrX.aYunfbcy.Address()
}

func (oPtEQcrX *WyJ004I8) Value() *uint256.Int {
	return oPtEQcrX.mddAyugJc
}

func (oPtEQcrX *WyJ004I8) SetCallCode(ewLbwkms *common.Address, vvL5SsAl common.Hash, euva86f0e []byte) {
	oPtEQcrX.STabmC = euva86f0e
	oPtEQcrX.YfrpaFi2iY3 = vvL5SsAl
	oPtEQcrX.DkqqBAVFZj9n = ewLbwkms
}

func (oPtEQcrX *WyJ004I8) SetCodeOptionalHash(ewLbwkms *common.Address, hAawFa8 *codeAndHash) {
	oPtEQcrX.STabmC = hAawFa8.code
	oPtEQcrX.YfrpaFi2iY3 = hAawFa8.hash
	oPtEQcrX.DkqqBAVFZj9n = ewLbwkms
}

var _ = common.AbsolutePath
var _ = uint256.ErrBadBufferLength
