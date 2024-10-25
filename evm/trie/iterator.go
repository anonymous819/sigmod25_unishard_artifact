//line :1
package trie

import (
	"bytes"
	"container/heap"
	"errors"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
)

type NodeResolver func(zalM3_NR4XT common.Hash, qiRG6lpeaFl []byte, rNuN0sMPDJ common.Hash) []byte

type ZojeMiTW struct {
	miALa3amW FBy5J0gbkg

	JyqRfQ2XM6 []byte
	H884Yc     []byte
	WitCUHMMV  error
}

func YH_VLHxOhT4(qGFIOZJlpK8 FBy5J0gbkg) *ZojeMiTW {
	return &ZojeMiTW{
		miALa3amW: qGFIOZJlpK8,
	}
}

func (qGFIOZJlpK8 *ZojeMiTW) Next() bool {
	for /*line DE8wD8ikt.go:1*/ qGFIOZJlpK8.miALa3amW.Next(true) {
		if /*line eym3BML.go:1*/ qGFIOZJlpK8.miALa3amW.Leaf() {
			qGFIOZJlpK8.JyqRfQ2XM6 = /*line AF8yy6l.go:1*/ qGFIOZJlpK8.miALa3amW.LeafKey()
			qGFIOZJlpK8.H884Yc = /*line WWCgjNOo5Ga.go:1*/ qGFIOZJlpK8.miALa3amW.LeafBlob()
			return true
		}
	}
	qGFIOZJlpK8.JyqRfQ2XM6 = nil
	qGFIOZJlpK8.H884Yc = nil
	qGFIOZJlpK8.WitCUHMMV = /*line H6dKdw.go:1*/ qGFIOZJlpK8.miALa3amW.Error()
	return false
}

func (qGFIOZJlpK8 *ZojeMiTW) Prove() [][]byte {
	return /*line AwoYox1VWQDK.go:1*/ qGFIOZJlpK8.miALa3amW.LeafProof()
}

type FBy5J0gbkg interface {
	Next(bool) bool

	Error() error

	Hash() common.Hash

	Parent() common.Hash

	Path() []byte

	NodeBlob() []byte

	Leaf() bool

	LeafKey() []byte

	LeafBlob() []byte

	LeafProof() [][]byte

	AddResolver(NodeResolver)
}

type nodeIteratorState struct {
	hash    common.Hash
	node    node
	parent  common.Hash
	index   int
	pathlen int
}

type nodeIterator struct {
	trie  *Trie
	stack []*nodeIteratorState
	path  []byte
	err   error

	resolver NodeResolver
	pool     []*nodeIteratorState
}

var i3mk7ZCo9qL = /*line jaKx_SAYp.go:1*/ errors.New(func() /*line im04_i.go:1*/ string {
	data := [] /*line im04_i.go:1*/ byte("\"!)/ $ d;era+io4")
	positions := [...]byte{4, 12, 0, 13, 15, 5, 7, 3, 3, 5, 12, 5, 2, 5, 12, 0, 1, 8, 3, 0}
	for i := 0; i < 20; i += 2 {
		localKey := /*line im04_i.go:1*/ byte(i) + /*line im04_i.go:1*/ byte(positions[i]^positions[i+1]) + 60
		data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
	}
	return /*line im04_i.go:1*/ string(data)
}())

type daOLQi struct {
	cEQRuz     []byte
	zR6VsQS3rD error
}

func (svhEA52x0Bm daOLQi) Error() string {
	return func() /*line Da051m6.go:1*/ string {
		fullData := [] /*line Da051m6.go:1*/ byte("Z,\x18G\x9d\xeb\xc85\x0eu\x19\xecR\xa8Z\"\xbf\x83\xcd\x02\v\xb0\xf4\xd4")
		idxKey := [] /*line Da051m6.go:1*/ byte("`\xf7")
		data := /*line Da051m6.go:1*/ make([]byte, 0, 13)
		data = /*line Da051m6.go:1*/ append(data, fullData[254^ /*line Da051m6.go:1*/ int(idxKey[1])]-fullData[228^ /*line Da051m6.go:1*/ int(idxKey[1])], fullData[229^ /*line Da051m6.go:1*/ int(idxKey[1])]^fullData[250^ /*line Da051m6.go:1*/ int(idxKey[1])], fullData[231^ /*line Da051m6.go:1*/ int(idxKey[1])]-fullData[249^ /*line Da051m6.go:1*/ int(idxKey[1])], fullData[108^ /*line Da051m6.go:1*/ int(idxKey[0])]+fullData[106^ /*line Da051m6.go:1*/ int(idxKey[0])], fullData[227^ /*line Da051m6.go:1*/ int(idxKey[1])]-fullData[242^ /*line Da051m6.go:1*/ int(idxKey[1])], fullData[100^ /*line Da051m6.go:1*/ int(idxKey[0])]+fullData[102^ /*line Da051m6.go:1*/ int(idxKey[0])], fullData[248^ /*line Da051m6.go:1*/ int(idxKey[1])]-fullData[226^ /*line Da051m6.go:1*/ int(idxKey[1])], fullData[96^ /*line Da051m6.go:1*/ int(idxKey[0])]+fullData[98^ /*line Da051m6.go:1*/ int(idxKey[0])], fullData[107^ /*line Da051m6.go:1*/ int(idxKey[0])]+fullData[113^ /*line Da051m6.go:1*/ int(idxKey[0])], fullData[99^ /*line Da051m6.go:1*/ int(idxKey[0])]^fullData[103^ /*line Da051m6.go:1*/ int(idxKey[0])], fullData[246^ /*line Da051m6.go:1*/ int(idxKey[1])]+fullData[255^ /*line Da051m6.go:1*/ int(idxKey[1])], fullData[118^ /*line Da051m6.go:1*/ int(idxKey[0])]^fullData[119^ /*line Da051m6.go:1*/ int(idxKey[0])])
		return /*line Da051m6.go:1*/ string(data)
	}() + /*line a4hrEDK131xN.go:1*/ svhEA52x0Bm.zR6VsQS3rD.Error()
}

func u98WTduaeU2Y(gExqOlCxa53E *Trie, iS1397t []byte) FBy5J0gbkg {
	if /*line vuK3oG7mIy.go:1*/ gExqOlCxa53E.Hash() == types.NrGuaNA21 {
		return &nodeIterator{
			trie: gExqOlCxa53E,
			err:  i3mk7ZCo9qL,
		}
	}
	qGFIOZJlpK8 := &nodeIterator{trie: gExqOlCxa53E}
	qGFIOZJlpK8.err = /*line aHzV990.go:1*/ qGFIOZJlpK8.nIzh_B(iS1397t)
	return qGFIOZJlpK8
}

func (qGFIOZJlpK8 *nodeIterator) qzNv8_3fx6C(bAufRbtZ9MFz *nodeIteratorState) {
	if /*line EaQw1qgT4i5.go:1*/ len(qGFIOZJlpK8.pool) < 40 {
		bAufRbtZ9MFz.node = nil
		qGFIOZJlpK8.pool = /*line wRsqcgxcJPc4.go:1*/ append(qGFIOZJlpK8.pool, bAufRbtZ9MFz)
	}
}

func (qGFIOZJlpK8 *nodeIterator) qknotTQoU() *nodeIteratorState {
	hOtJPXr := /*line Lb9_oB7ggwY.go:1*/ len(qGFIOZJlpK8.pool) - 1
	if hOtJPXr < 0 {
		return /*line VYrtulu.go:1*/ new(nodeIteratorState)
	}
	dz9jA8sz4 := qGFIOZJlpK8.pool[hOtJPXr]
	qGFIOZJlpK8.pool[hOtJPXr] = nil
	qGFIOZJlpK8.pool = qGFIOZJlpK8.pool[:hOtJPXr]
	return dz9jA8sz4
}

func (qGFIOZJlpK8 *nodeIterator) AddResolver(aRra_C_9g NodeResolver) {
	qGFIOZJlpK8.resolver = aRra_C_9g
}

func (qGFIOZJlpK8 *nodeIterator) Hash() common.Hash {
	if /*line SNdooybZPF.go:1*/ len(qGFIOZJlpK8.stack) == 0 {
		return common.Hash{}
	}
	return qGFIOZJlpK8.stack[ /*line QW2NeUuOc.go:1*/ len(qGFIOZJlpK8.stack)-1].hash
}

func (qGFIOZJlpK8 *nodeIterator) Parent() common.Hash {
	if /*line grUhlyggu.go:1*/ len(qGFIOZJlpK8.stack) == 0 {
		return common.Hash{}
	}
	return qGFIOZJlpK8.stack[ /*line c3X17Q.go:1*/ len(qGFIOZJlpK8.stack)-1].parent
}

func (qGFIOZJlpK8 *nodeIterator) Leaf() bool {
	return /*line BartNpQpD.go:1*/ k7wsvmNwK(qGFIOZJlpK8.path)
}

func (qGFIOZJlpK8 *nodeIterator) LeafKey() []byte {
	if /*line hWaIeK.go:1*/ len(qGFIOZJlpK8.stack) > 0 {
		if _, yY_yPSd8vG := qGFIOZJlpK8.stack[ /*line fBPwraqd3bj.go:1*/ len(qGFIOZJlpK8.stack)-1].node.(valueNode); yY_yPSd8vG {
			return /*line DPO2EqvM7VVj.go:1*/ iZl0DAYMB(qGFIOZJlpK8.path)
		}
	}
	/*line BIfZFop.go:1*/ panic(func() /*line KOaOaBaT.go:1*/ string {
		fullData := [] /*line KOaOaBaT.go:1*/ byte("]\xa9\xe7%\xf5Y\x90\xa9\x19\xcaV\xb1\xc3\xe8ݸM\xb6\x87$\x865")
		idxKey := [] /*line KOaOaBaT.go:1*/ byte("\xdd7")
		data := /*line KOaOaBaT.go:1*/ make([]byte, 0, 12)
		data = /*line KOaOaBaT.go:1*/ append(data, fullData[36^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]-fullData[38^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[208^ /*line KOaOaBaT.go:1*/ int(idxKey[0])]^fullData[207^ /*line KOaOaBaT.go:1*/ int(idxKey[0])], fullData[52^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]-fullData[60^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[59^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]+fullData[55^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[53^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]^fullData[35^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[48^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]^fullData[57^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[62^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]+fullData[61^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[200^ /*line KOaOaBaT.go:1*/ int(idxKey[0])]^fullData[216^ /*line KOaOaBaT.go:1*/ int(idxKey[0])], fullData[51^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]^fullData[49^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[54^ /*line KOaOaBaT.go:1*/ int(idxKey[1])]+fullData[56^ /*line KOaOaBaT.go:1*/ int(idxKey[1])], fullData[213^ /*line KOaOaBaT.go:1*/ int(idxKey[0])]+fullData[205^ /*line KOaOaBaT.go:1*/ int(idxKey[0])])
		return /*line KOaOaBaT.go:1*/ string(data)
	}())
}

func (qGFIOZJlpK8 *nodeIterator) LeafBlob() []byte {
	if /*line fX8iUB5JUav.go:1*/ len(qGFIOZJlpK8.stack) > 0 {
		if uAjz8NxU_cL, yY_yPSd8vG := qGFIOZJlpK8.stack[ /*line byFeNrL.go:1*/ len(qGFIOZJlpK8.stack)-1].node.(valueNode); yY_yPSd8vG {
			return uAjz8NxU_cL
		}
	}
	/*line JsRqLp.go:1*/ panic(func() /*line oTvTq6p6cs.go:1*/ string {
		fullData := [] /*line oTvTq6p6cs.go:1*/ byte("I\xd2\x1dڶ|\xd2<\x18Ɲ\xf2\xc8\x1b\xe1Ī\xa0\t\xb6\x95)")
		idxKey := [] /*line oTvTq6p6cs.go:1*/ byte("\xec\xc62\x15z\xd2^i\x1d")
		data := /*line oTvTq6p6cs.go:1*/ make([]byte, 0, 12)
		data = /*line oTvTq6p6cs.go:1*/ append(data, fullData[106^ /*line oTvTq6p6cs.go:1*/ int(idxKey[4])]+fullData[117^ /*line oTvTq6p6cs.go:1*/ int(idxKey[4])], fullData[23^ /*line oTvTq6p6cs.go:1*/ int(idxKey[8])]+fullData[28^ /*line oTvTq6p6cs.go:1*/ int(idxKey[8])], fullData[18^ /*line oTvTq6p6cs.go:1*/ int(idxKey[3])]-fullData[25^ /*line oTvTq6p6cs.go:1*/ int(idxKey[3])], fullData[123^ /*line oTvTq6p6cs.go:1*/ int(idxKey[7])]^fullData[124^ /*line oTvTq6p6cs.go:1*/ int(idxKey[7])], fullData[107^ /*line oTvTq6p6cs.go:1*/ int(idxKey[7])]^fullData[108^ /*line oTvTq6p6cs.go:1*/ int(idxKey[7])], fullData[38^ /*line oTvTq6p6cs.go:1*/ int(idxKey[2])]^fullData[60^ /*line oTvTq6p6cs.go:1*/ int(idxKey[2])], fullData[57^ /*line oTvTq6p6cs.go:1*/ int(idxKey[2])]-fullData[52^ /*line oTvTq6p6cs.go:1*/ int(idxKey[2])], fullData[105^ /*line oTvTq6p6cs.go:1*/ int(idxKey[4])]^fullData[121^ /*line oTvTq6p6cs.go:1*/ int(idxKey[4])], fullData[16^ /*line oTvTq6p6cs.go:1*/ int(idxKey[8])]-fullData[25^ /*line oTvTq6p6cs.go:1*/ int(idxKey[8])], fullData[21^ /*line oTvTq6p6cs.go:1*/ int(idxKey[3])]+fullData[29^ /*line oTvTq6p6cs.go:1*/ int(idxKey[3])], fullData[4^ /*line oTvTq6p6cs.go:1*/ int(idxKey[3])]^fullData[28^ /*line oTvTq6p6cs.go:1*/ int(idxKey[3])])
		return /*line oTvTq6p6cs.go:1*/ string(data)
	}())
}

func (qGFIOZJlpK8 *nodeIterator) LeafProof() [][]byte {
	if /*line wabtCZOy_iZC.go:1*/ len(qGFIOZJlpK8.stack) > 0 {
		if _, yY_yPSd8vG := qGFIOZJlpK8.stack[ /*line _voLDrJHIBqC.go:1*/ len(qGFIOZJlpK8.stack)-1].node.(valueNode); yY_yPSd8vG {
			d6imPJ := /*line LIrdIFsUS.go:1*/ dCZJziX(false)
			defer /*line UKQZe_rHY.go:1*/ ptIMYRK9(d6imPJ)
			a8EIiGJB := /*line M9WwV4t2q.go:1*/ make([][]byte, 0 /*line YaGNzMe.go:1*/, len(qGFIOZJlpK8.stack))

			for q2u2020KD6, bAufRbtZ9MFz := range qGFIOZJlpK8.stack[: /*line Ww_P78vf.go:1*/ len(qGFIOZJlpK8.stack)-1] {

				uAjz8NxU_cL, iVcAtTyHaK := /*line E4V8mQ.go:1*/ d6imPJ.aCBICr2(bAufRbtZ9MFz.node)
				if _, yY_yPSd8vG := iVcAtTyHaK.(hashNode); yY_yPSd8vG || q2u2020KD6 == 0 {
					a8EIiGJB = /*line DN396sxv0Enq.go:1*/ append(a8EIiGJB /*line Z8EfpMP.go:1*/, kG18lbsZr(uAjz8NxU_cL))
				}
			}
			return a8EIiGJB
		}
	}
	/*line jWx8ar4lO75v.go:1*/ panic(func() /*line QfZHDNMJ0jVO.go:1*/ string {
		data := [] /*line QfZHDNMJ0jVO.go:1*/ byte("nom\xc9a\xd4 \x92.0\"")
		positions := [...]byte{3, 10, 5, 7, 9, 3, 8, 2, 3, 5, 8, 3, 3, 8, 2, 2}
		for i := 0; i < 16; i += 2 {
			localKey := /*line QfZHDNMJ0jVO.go:1*/ byte(i) + /*line QfZHDNMJ0jVO.go:1*/ byte(positions[i]^positions[i+1]) + 148
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line QfZHDNMJ0jVO.go:1*/ string(data)
	}())
}

func (qGFIOZJlpK8 *nodeIterator) Path() []byte {
	return qGFIOZJlpK8.path
}

func (qGFIOZJlpK8 *nodeIterator) NodeBlob() []byte {
	if /*line Eg2tY_9CkDTK.go:1*/ qGFIOZJlpK8.Hash() == (common.Hash{}) {
		return nil
	}
	aBHduJG, eZzE0KPS := /*line jaqLoEq.go:1*/ qGFIOZJlpK8.gm5ThOsCwp( /*line n_37co4QKX.go:1*/ qGFIOZJlpK8.Hash().Bytes() /*line gbNflb.go:1*/, qGFIOZJlpK8.Path())
	if eZzE0KPS != nil {
		qGFIOZJlpK8.err = eZzE0KPS
		return nil
	}
	return aBHduJG
}

func (qGFIOZJlpK8 *nodeIterator) Error() error {
	if qGFIOZJlpK8.err == i3mk7ZCo9qL {
		return nil
	}
	if nIzh_B, yY_yPSd8vG := qGFIOZJlpK8.err.(daOLQi); yY_yPSd8vG {
		return nIzh_B.zR6VsQS3rD
	}
	return qGFIOZJlpK8.err
}

func (qGFIOZJlpK8 *nodeIterator) Next(h6UJjHlET bool) bool {
	if qGFIOZJlpK8.err == i3mk7ZCo9qL {
		return false
	}
	if nIzh_B, yY_yPSd8vG := qGFIOZJlpK8.err.(daOLQi); yY_yPSd8vG {
		if qGFIOZJlpK8.err = /*line kAAvy0KFs5.go:1*/ qGFIOZJlpK8.nIzh_B(nIzh_B.cEQRuz); qGFIOZJlpK8.err != nil {
			return false
		}
	}

	b9P6UtNKOlhb, dnUkxKJn, qiRG6lpeaFl, eZzE0KPS := /*line bYa42e7amn.go:1*/ qGFIOZJlpK8.mYiaufDE98NH(h6UJjHlET)
	qGFIOZJlpK8.err = eZzE0KPS
	if qGFIOZJlpK8.err != nil {
		return false
	}
	/*line AqvcgaiM5I.go:1*/ qGFIOZJlpK8.uWm6GHNVGM(b9P6UtNKOlhb, dnUkxKJn, qiRG6lpeaFl)
	return true
}

func (qGFIOZJlpK8 *nodeIterator) nIzh_B(tKBrjZ []byte) error {

	lhQWH7m := /*line PipffRb.go:1*/ nd12pU880Z8(tKBrjZ)
	lhQWH7m = lhQWH7m[: /*line UdebbGhdJNG.go:1*/ len(lhQWH7m)-1]

	for {
		b9P6UtNKOlhb, dnUkxKJn, qiRG6lpeaFl, eZzE0KPS := /*line LIhWw_00.go:1*/ qGFIOZJlpK8.rqOxMpJmm5DJ(lhQWH7m)
		if eZzE0KPS == i3mk7ZCo9qL {
			return i3mk7ZCo9qL
		} else if eZzE0KPS != nil {
			return daOLQi{tKBrjZ, eZzE0KPS}
		} else if /*line Qj12az0dUp.go:1*/ bytes.Compare(qiRG6lpeaFl, lhQWH7m) >= 0 {
			return nil
		}
		/*line CQ0yNQp6r.go:1*/ qGFIOZJlpK8.uWm6GHNVGM(b9P6UtNKOlhb, dnUkxKJn, qiRG6lpeaFl)
	}
}

func (qGFIOZJlpK8 *nodeIterator) init() (*nodeIteratorState, error) {
	m7S8SE := /*line H754EDz.go:1*/ qGFIOZJlpK8.trie.Hash()
	b9P6UtNKOlhb := &nodeIteratorState{node: qGFIOZJlpK8.trie.root, index: -1}
	if m7S8SE != types.NrGuaNA21 {
		b9P6UtNKOlhb.hash = m7S8SE
	}
	return b9P6UtNKOlhb /*line MDjDDCL.go:1*/, b9P6UtNKOlhb.v2ULJu(qGFIOZJlpK8, nil)
}

func (qGFIOZJlpK8 *nodeIterator) mYiaufDE98NH(h6UJjHlET bool) (*nodeIteratorState, *int, []byte, error) {

	if /*line pcjnndRyv5IG.go:1*/ len(qGFIOZJlpK8.stack) == 0 {
		b9P6UtNKOlhb, eZzE0KPS := /*line h3ZETsx6C.go:1*/ qGFIOZJlpK8.init()
		return b9P6UtNKOlhb, nil, nil, eZzE0KPS
	}
	if !h6UJjHlET {

		/*line klgNOdwlZEW.go:1*/
		qGFIOZJlpK8.cT7GLkyQ6_B()
	}

	for /*line Xlf59MBxx.go:1*/ len(qGFIOZJlpK8.stack) > 0 {
		nmFIUUSGjl := qGFIOZJlpK8.stack[ /*line NGJiKYMkDK.go:1*/ len(qGFIOZJlpK8.stack)-1]
		bvp1dT2K5 := nmFIUUSGjl.hash
		if (bvp1dT2K5 == common.Hash{}) {
			bvp1dT2K5 = nmFIUUSGjl.parent
		}
		b9P6UtNKOlhb, qiRG6lpeaFl, yY_yPSd8vG := /*line HcMgUFHP.go:1*/ qGFIOZJlpK8.oYT6z7M3SR(nmFIUUSGjl, bvp1dT2K5)
		if yY_yPSd8vG {
			if eZzE0KPS := /*line wVpZJaZ.go:1*/ b9P6UtNKOlhb.v2ULJu(qGFIOZJlpK8, qiRG6lpeaFl); eZzE0KPS != nil {
				return nmFIUUSGjl, &nmFIUUSGjl.index, qiRG6lpeaFl, eZzE0KPS
			}
			return b9P6UtNKOlhb, &nmFIUUSGjl.index, qiRG6lpeaFl, nil
		}

		/*line iR9Qe_F.go:1*/
		qGFIOZJlpK8.cT7GLkyQ6_B()
	}
	return nil, nil, nil, i3mk7ZCo9qL
}

func (qGFIOZJlpK8 *nodeIterator) rqOxMpJmm5DJ(jlz6KI_1 []byte) (*nodeIteratorState, *int, []byte, error) {

	if /*line Fmv3lk_FfQu.go:1*/ len(qGFIOZJlpK8.stack) == 0 {
		b9P6UtNKOlhb, eZzE0KPS := /*line Jbwj1Yk2OJe.go:1*/ qGFIOZJlpK8.init()
		return b9P6UtNKOlhb, nil, nil, eZzE0KPS
	}
	if ! /*line OzTkojHJK.go:1*/ bytes.HasPrefix(jlz6KI_1, qGFIOZJlpK8.path) {

		/*line _1x2pIB9LOY.go:1*/
		qGFIOZJlpK8.cT7GLkyQ6_B()
	}

	for /*line h2aDm4GWgZ.go:1*/ len(qGFIOZJlpK8.stack) > 0 {
		nmFIUUSGjl := qGFIOZJlpK8.stack[ /*line LNO6jigkrxoO.go:1*/ len(qGFIOZJlpK8.stack)-1]
		bvp1dT2K5 := nmFIUUSGjl.hash
		if (bvp1dT2K5 == common.Hash{}) {
			bvp1dT2K5 = nmFIUUSGjl.parent
		}
		b9P6UtNKOlhb, qiRG6lpeaFl, yY_yPSd8vG := /*line YQpswmy.go:1*/ qGFIOZJlpK8.nqoonNhJ5(nmFIUUSGjl, bvp1dT2K5, jlz6KI_1)
		if yY_yPSd8vG {
			if eZzE0KPS := /*line QYJI9Do7oaPq.go:1*/ b9P6UtNKOlhb.v2ULJu(qGFIOZJlpK8, qiRG6lpeaFl); eZzE0KPS != nil {
				return nmFIUUSGjl, &nmFIUUSGjl.index, qiRG6lpeaFl, eZzE0KPS
			}
			return b9P6UtNKOlhb, &nmFIUUSGjl.index, qiRG6lpeaFl, nil
		}

		/*line gybm8xvXg.go:1*/
		qGFIOZJlpK8.cT7GLkyQ6_B()
	}
	return nil, nil, nil, i3mk7ZCo9qL
}

func (qGFIOZJlpK8 *nodeIterator) zz0UEW(rNuN0sMPDJ hashNode, qiRG6lpeaFl []byte) (node, error) {
	if qGFIOZJlpK8.resolver != nil {
		if aBHduJG := /*line em5nQa0LEM.go:1*/ qGFIOZJlpK8.resolver(qGFIOZJlpK8.trie.owner, qiRG6lpeaFl /*line Zrfg7HdF.go:1*/, common.BytesToHash(rNuN0sMPDJ)); /*line UWyQgasV.go:1*/ len(aBHduJG) > 0 {
			if lR8wDJYky, eZzE0KPS := /*line GSJW2PuzjA.go:1*/ rRV2VeUAEp(rNuN0sMPDJ, aBHduJG); eZzE0KPS == nil {
				return lR8wDJYky, nil
			}
		}
	}

	aBHduJG, eZzE0KPS := /*line OASPYiOQzf.go:1*/ qGFIOZJlpK8.trie.reader.uAjz8NxU_cL(qiRG6lpeaFl /*line RMdaL2rd.go:1*/, common.BytesToHash(rNuN0sMPDJ))
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}

	return /*line bCfxM8RB_q.go:1*/ hct27E(rNuN0sMPDJ, aBHduJG), nil
}

func (qGFIOZJlpK8 *nodeIterator) gm5ThOsCwp(rNuN0sMPDJ hashNode, qiRG6lpeaFl []byte) ([]byte, error) {
	if qGFIOZJlpK8.resolver != nil {
		if aBHduJG := /*line cufApJL.go:1*/ qGFIOZJlpK8.resolver(qGFIOZJlpK8.trie.owner, qiRG6lpeaFl /*line opzNJRD.go:1*/, common.BytesToHash(rNuN0sMPDJ)); /*line WdzxCG.go:1*/ len(aBHduJG) > 0 {
			return aBHduJG, nil
		}
	}

	return /*line AacJrcvXT.go:1*/ qGFIOZJlpK8.trie.reader.uAjz8NxU_cL(qiRG6lpeaFl /*line lo58yaVFH.go:1*/, common.BytesToHash(rNuN0sMPDJ))
}

func (jIgKH4h *nodeIteratorState) v2ULJu(qGFIOZJlpK8 *nodeIterator, qiRG6lpeaFl []byte) error {
	if rNuN0sMPDJ, yY_yPSd8vG := jIgKH4h.node.(hashNode); yY_yPSd8vG {
		lR8wDJYky, eZzE0KPS := /*line Knn_nIlkq9.go:1*/ qGFIOZJlpK8.zz0UEW(rNuN0sMPDJ, qiRG6lpeaFl)
		if eZzE0KPS != nil {
			return eZzE0KPS
		}
		jIgKH4h.node = lR8wDJYky
		jIgKH4h.hash = /*line TsSPjrkkZG4d.go:1*/ common.BytesToHash(rNuN0sMPDJ)
	}
	return nil
}

func (qGFIOZJlpK8 *nodeIterator) gN9gDh9oc(gnGMaXTuiFeE *fullNode, sCuaQRs1 int, bvp1dT2K5 common.Hash) (node, *nodeIteratorState, []byte, int) {
	var (
		qiRG6lpeaFl  = qGFIOZJlpK8.path
		jcDLabJ7o    node
		b9P6UtNKOlhb *nodeIteratorState
		sSJ1NN       []byte
	)
	for ; sCuaQRs1 < /*line Hd4oTFH676QE.go:1*/ len(gnGMaXTuiFeE.Children); sCuaQRs1++ {
		if gnGMaXTuiFeE.Children[sCuaQRs1] != nil {
			jcDLabJ7o = gnGMaXTuiFeE.Children[sCuaQRs1]
			rNuN0sMPDJ, _ := /*line I0zoO2pC.go:1*/ jcDLabJ7o.tGJzZYYLvK()
			b9P6UtNKOlhb = /*line kqmya6.go:1*/ qGFIOZJlpK8.qknotTQoU()
			b9P6UtNKOlhb.hash = /*line meC40Gi62zyU.go:1*/ common.BytesToHash(rNuN0sMPDJ)
			b9P6UtNKOlhb.node = jcDLabJ7o
			b9P6UtNKOlhb.parent = bvp1dT2K5
			b9P6UtNKOlhb.index = -1
			b9P6UtNKOlhb.pathlen = /*line ak_eBtm.go:1*/ len(qiRG6lpeaFl)
			sSJ1NN = /*line ob7R5v.go:1*/ append(sSJ1NN, qiRG6lpeaFl...)
			sSJ1NN = /*line NViKAbt.go:1*/ append(sSJ1NN /*line axJoqQMq.go:1*/, byte(sCuaQRs1))
			return jcDLabJ7o, b9P6UtNKOlhb, sSJ1NN, sCuaQRs1
		}
	}
	return nil, nil, nil, 0
}

func (qGFIOZJlpK8 *nodeIterator) oYT6z7M3SR(nmFIUUSGjl *nodeIteratorState, bvp1dT2K5 common.Hash) (*nodeIteratorState, []byte, bool) {
	switch uAjz8NxU_cL := nmFIUUSGjl.node.(type) {
	case *fullNode:

		if jcDLabJ7o, b9P6UtNKOlhb, qiRG6lpeaFl, sCuaQRs1 := /*line m9iVksc.go:1*/ qGFIOZJlpK8.gN9gDh9oc(uAjz8NxU_cL, nmFIUUSGjl.index+1, bvp1dT2K5); jcDLabJ7o != nil {
			nmFIUUSGjl.index = sCuaQRs1 - 1
			return b9P6UtNKOlhb, qiRG6lpeaFl, true
		}
	case *qUKQUCfTXB:

		if nmFIUUSGjl.index < 0 {
			rNuN0sMPDJ, _ := /*line eoANG3oBHyq.go:1*/ uAjz8NxU_cL.YpmEYGB.tGJzZYYLvK()
			b9P6UtNKOlhb := /*line LTPUUs5q.go:1*/ qGFIOZJlpK8.qknotTQoU()
			b9P6UtNKOlhb.hash = /*line ZG6gaPaam.go:1*/ common.BytesToHash(rNuN0sMPDJ)
			b9P6UtNKOlhb.node = uAjz8NxU_cL.YpmEYGB
			b9P6UtNKOlhb.parent = bvp1dT2K5
			b9P6UtNKOlhb.index = -1
			b9P6UtNKOlhb.pathlen = /*line cRDhf5X.go:1*/ len(qGFIOZJlpK8.path)
			qiRG6lpeaFl := /*line VSKsEe.go:1*/ append(qGFIOZJlpK8.path, uAjz8NxU_cL.ANdZYqbzJ1A...)
			return b9P6UtNKOlhb, qiRG6lpeaFl, true
		}
	}
	return nmFIUUSGjl, qGFIOZJlpK8.path, false
}

func (qGFIOZJlpK8 *nodeIterator) nqoonNhJ5(nmFIUUSGjl *nodeIteratorState, bvp1dT2K5 common.Hash, lhQWH7m []byte) (*nodeIteratorState, []byte, bool) {
	switch gnGMaXTuiFeE := nmFIUUSGjl.node.(type) {
	case *fullNode:

		jcDLabJ7o, b9P6UtNKOlhb, qiRG6lpeaFl, sCuaQRs1 := /*line pl5VnkOv.go:1*/ qGFIOZJlpK8.gN9gDh9oc(gnGMaXTuiFeE, nmFIUUSGjl.index+1, bvp1dT2K5)
		if jcDLabJ7o == nil {

			return nmFIUUSGjl, qGFIOZJlpK8.path, false
		}

		if /*line noOaRY6.go:1*/ bytes.Compare(qiRG6lpeaFl, lhQWH7m) >= 0 {
			nmFIUUSGjl.index = sCuaQRs1 - 1
			return b9P6UtNKOlhb, qiRG6lpeaFl, true
		}

		for {
			oYT6z7M3SR, kyflesS, hk2Cq6rkt, oPnci_ := /*line F9AdrAKg.go:1*/ qGFIOZJlpK8.gN9gDh9oc(gnGMaXTuiFeE, sCuaQRs1+1, bvp1dT2K5)

			if oYT6z7M3SR == nil || /*line vnROBYcmRvSg.go:1*/ bytes.Compare(hk2Cq6rkt, lhQWH7m) >= 0 {
				nmFIUUSGjl.index = sCuaQRs1 - 1
				return b9P6UtNKOlhb, qiRG6lpeaFl, true
			}

			b9P6UtNKOlhb, qiRG6lpeaFl, sCuaQRs1 = kyflesS, hk2Cq6rkt, oPnci_
		}
	case *qUKQUCfTXB:

		if nmFIUUSGjl.index < 0 {
			rNuN0sMPDJ, _ := /*line bAOVbt.go:1*/ gnGMaXTuiFeE.YpmEYGB.tGJzZYYLvK()
			b9P6UtNKOlhb := /*line o6lPSO.go:1*/ qGFIOZJlpK8.qknotTQoU()
			b9P6UtNKOlhb.hash = /*line XOje86RV.go:1*/ common.BytesToHash(rNuN0sMPDJ)
			b9P6UtNKOlhb.node = gnGMaXTuiFeE.YpmEYGB
			b9P6UtNKOlhb.parent = bvp1dT2K5
			b9P6UtNKOlhb.index = -1
			b9P6UtNKOlhb.pathlen = /*line BagGLCOXcnN.go:1*/ len(qGFIOZJlpK8.path)
			qiRG6lpeaFl := /*line hA9__PC.go:1*/ append(qGFIOZJlpK8.path, gnGMaXTuiFeE.ANdZYqbzJ1A...)
			return b9P6UtNKOlhb, qiRG6lpeaFl, true
		}
	}
	return nmFIUUSGjl, qGFIOZJlpK8.path, false
}

func (qGFIOZJlpK8 *nodeIterator) uWm6GHNVGM(b9P6UtNKOlhb *nodeIteratorState, dnUkxKJn *int, qiRG6lpeaFl []byte) {
	qGFIOZJlpK8.path = qiRG6lpeaFl
	qGFIOZJlpK8.stack = /*line nD3070eoa.go:1*/ append(qGFIOZJlpK8.stack, b9P6UtNKOlhb)
	if dnUkxKJn != nil {
		*dnUkxKJn++
	}
}

func (qGFIOZJlpK8 *nodeIterator) cT7GLkyQ6_B() {
	k0F4spwdWm := qGFIOZJlpK8.stack[ /*line PUzqg9T17.go:1*/ len(qGFIOZJlpK8.stack)-1]
	qGFIOZJlpK8.path = qGFIOZJlpK8.path[:k0F4spwdWm.pathlen]
	qGFIOZJlpK8.stack[ /*line Oona1vMDwV.go:1*/ len(qGFIOZJlpK8.stack)-1] = nil
	qGFIOZJlpK8.stack = qGFIOZJlpK8.stack[: /*line ENxW4G5ke.go:1*/ len(qGFIOZJlpK8.stack)-1]

	/*line wKqmcN7SYKpN.go:1*/
	qGFIOZJlpK8.qzNv8_3fx6C(k0F4spwdWm)
}

func cSqBHdRTg3(vXUk1wf, mZsNbw0 FBy5J0gbkg) int {
	if h_9XtYGhV5 := /*line U8dkA2ZQ_c9.go:1*/ bytes.Compare( /*line xIGxUjd.go:1*/ vXUk1wf.Path() /*line a19503zR.go:1*/, mZsNbw0.Path()); h_9XtYGhV5 != 0 {
		return h_9XtYGhV5
	}
	if /*line JBewL3.go:1*/ vXUk1wf.Leaf() && ! /*line KhMbi9E069C.go:1*/ mZsNbw0.Leaf() {
		return -1
	} else if /*line _tvkB4Zfp4.go:1*/ mZsNbw0.Leaf() && ! /*line txFtInoXLt.go:1*/ vXUk1wf.Leaf() {
		return 1
	}
	if h_9XtYGhV5 := /*line nBC0C_.go:1*/ bytes.Compare( /*line LZNEHmCa35x0.go:1*/ vXUk1wf.Hash().Bytes() /*line oZp1A6.go:1*/, mZsNbw0.Hash().Bytes()); h_9XtYGhV5 != 0 {
		return h_9XtYGhV5
	}
	if /*line dZ2PoK.go:1*/ vXUk1wf.Leaf() && /*line B0e4urOzQU.go:1*/ mZsNbw0.Leaf() {
		return /*line x4OVF9eVnh.go:1*/ bytes.Compare( /*line tWIfMaHr.go:1*/ vXUk1wf.LeafBlob() /*line THjHBs8CFBCz.go:1*/, mZsNbw0.LeafBlob())
	}
	return 0
}

type dYYPvRalkDg struct {
	aOkHsWzYTr, cOPQYOO FBy5J0gbkg
	zz35wCI             bool
	q5wAp0XpMA          int
}

func Rs861HQqYsmv(vXUk1wf, mZsNbw0 FBy5J0gbkg) (FBy5J0gbkg, *int) {
	/*line Emj23gM0b.go:1*/ vXUk1wf.Next(true)
	qGFIOZJlpK8 := &dYYPvRalkDg{
		aOkHsWzYTr: vXUk1wf,
		cOPQYOO:    mZsNbw0,
	}
	return qGFIOZJlpK8, &qGFIOZJlpK8.q5wAp0XpMA
}

func (qGFIOZJlpK8 *dYYPvRalkDg) Hash() common.Hash {
	return /*line sujMGrA5x.go:1*/ qGFIOZJlpK8.cOPQYOO.Hash()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) Parent() common.Hash {
	return /*line ee62yXY.go:1*/ qGFIOZJlpK8.cOPQYOO.Parent()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) Leaf() bool {
	return /*line parirr.go:1*/ qGFIOZJlpK8.cOPQYOO.Leaf()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) LeafKey() []byte {
	return /*line dT9V82eJKM.go:1*/ qGFIOZJlpK8.cOPQYOO.LeafKey()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) LeafBlob() []byte {
	return /*line eXqEJqIKP.go:1*/ qGFIOZJlpK8.cOPQYOO.LeafBlob()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) LeafProof() [][]byte {
	return /*line bmB8ZXlnrV.go:1*/ qGFIOZJlpK8.cOPQYOO.LeafProof()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) Path() []byte {
	return /*line ky2tCTOtci.go:1*/ qGFIOZJlpK8.cOPQYOO.Path()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) NodeBlob() []byte {
	return /*line NVxJWXb.go:1*/ qGFIOZJlpK8.cOPQYOO.NodeBlob()
}

func (qGFIOZJlpK8 *dYYPvRalkDg) AddResolver(aRra_C_9g NodeResolver) {
	/*line DzfmyN7imu.go:1*/ panic(func() /*line caCyUX.go:1*/ string {
		data := /*line caCyUX.go:1*/ make([]byte, 0, 16)
		i := 1
		decryptKey := 5
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				data = /*line caCyUX.go:1*/ append(data, "\x9a\x9d\xa3"...,
				)
				i = 4
			case 1:
				data = /*line caCyUX.go:1*/ append(data, "\xa3\xa3\xabV"...,
				)
				i = 5
			case 6:
				data = /*line caCyUX.go:1*/ append(data, 169)
				i = 3
			case 0:
				i = 2
				for y := range data {
					data[y] = data[y] - /*line caCyUX.go:1*/ byte(decryptKey^y)
				}
			case 7:
				i = 0
				data = /*line caCyUX.go:1*/ append(data, "\x9d\x9f"...,
				)
			case 3:
				i = 7
				data = /*line caCyUX.go:1*/ append(data, "\xa4\xac\xad"...,
				)
			case 4:
				i = 6
				data = /*line caCyUX.go:1*/ append(data, "\x9e\xa2"...,
				)
			}
		}
		return /*line caCyUX.go:1*/ string(data)
	}())
}

func (qGFIOZJlpK8 *dYYPvRalkDg) Next(bool) bool {

	if ! /*line DTxFtwa6I9.go:1*/ qGFIOZJlpK8.cOPQYOO.Next(true) {
		return false
	}
	qGFIOZJlpK8.q5wAp0XpMA++

	if qGFIOZJlpK8.zz35wCI {

		return true
	}

	for {
		switch /*line WdsZMz.go:1*/ cSqBHdRTg3(qGFIOZJlpK8.aOkHsWzYTr, qGFIOZJlpK8.cOPQYOO) {
		case -1:

			if ! /*line va1aJx2yDM.go:1*/ qGFIOZJlpK8.aOkHsWzYTr.Next(true) {
				qGFIOZJlpK8.zz35wCI = true
				return true
			}
			qGFIOZJlpK8.q5wAp0XpMA++
		case 1:

			return true
		case 0:

			j8fdCH02mY := /*line BA_DgFCRLt.go:1*/ qGFIOZJlpK8.aOkHsWzYTr.Hash() == common.Hash{}
			if ! /*line B30ehBNyX.go:1*/ qGFIOZJlpK8.cOPQYOO.Next(j8fdCH02mY) {
				return false
			}
			qGFIOZJlpK8.q5wAp0XpMA++
			if ! /*line kA86JWy6.go:1*/ qGFIOZJlpK8.aOkHsWzYTr.Next(j8fdCH02mY) {
				qGFIOZJlpK8.zz35wCI = true
				return true
			}
			qGFIOZJlpK8.q5wAp0XpMA++
		}
	}
}

func (qGFIOZJlpK8 *dYYPvRalkDg) Error() error {
	if eZzE0KPS := /*line oa_yR8a.go:1*/ qGFIOZJlpK8.aOkHsWzYTr.Error(); eZzE0KPS != nil {
		return eZzE0KPS
	}
	return /*line KCsjDFIbRiu.go:1*/ qGFIOZJlpK8.cOPQYOO.Error()
}

type ae4aXu7 []FBy5J0gbkg

func (klE3zy ae4aXu7) Len() int { return /*line paF6KiLgcR.go:1*/ len(klE3zy) }
func (klE3zy ae4aXu7) Less(q2u2020KD6, dhOtbRx_ int) bool {
	return /*line JljhTJfgiif.go:1*/ cSqBHdRTg3(klE3zy[q2u2020KD6], klE3zy[dhOtbRx_]) < 0
}
func (klE3zy ae4aXu7) Swap(q2u2020KD6, dhOtbRx_ int) {
	klE3zy[q2u2020KD6], klE3zy[dhOtbRx_] = klE3zy[dhOtbRx_], klE3zy[q2u2020KD6]
}
func (klE3zy *ae4aXu7) Push(kkMqX5HLl0n interface{}) {
	*klE3zy = /*line RojNeg2.go:1*/ append(*klE3zy, kkMqX5HLl0n.(FBy5J0gbkg))
}
func (klE3zy *ae4aXu7) Pop() interface{} {
	gnGMaXTuiFeE := /*line btwh17JrmLCI.go:1*/ len(*klE3zy)
	kkMqX5HLl0n := (*klE3zy)[gnGMaXTuiFeE-1]
	*klE3zy = (*klE3zy)[0 : gnGMaXTuiFeE-1]
	return kkMqX5HLl0n
}

type odcxjPIq struct {
	k4xbZk0XYN63 *ae4aXu7
	h9a1udIyKRBJ int
}

func Bofu18o(t3RjyuErN []FBy5J0gbkg) (FBy5J0gbkg, *int) {
	klE3zy := /*line tzJPU5IiwIl.go:1*/ make(ae4aXu7 /*line HWdMlw.go:1*/, len(t3RjyuErN))
	/*line vskb8T.go:1*/ copy(klE3zy, t3RjyuErN)
	/*line Ol2uXbtB.go:1*/ heap.Init(&klE3zy)

	vxSJBqqT := &odcxjPIq{k4xbZk0XYN63: &klE3zy}
	return vxSJBqqT, &vxSJBqqT.h9a1udIyKRBJ
}

func (qGFIOZJlpK8 *odcxjPIq) Hash() common.Hash {
	return (* /*line DyRVCUTNiKu.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].Hash()
}

func (qGFIOZJlpK8 *odcxjPIq) Parent() common.Hash {
	return (* /*line jkai_dPdU.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].Parent()
}

func (qGFIOZJlpK8 *odcxjPIq) Leaf() bool {
	return (* /*line BlNaaW6Yaj5.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].Leaf()
}

func (qGFIOZJlpK8 *odcxjPIq) LeafKey() []byte {
	return (* /*line URr_PW.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].LeafKey()
}

func (qGFIOZJlpK8 *odcxjPIq) LeafBlob() []byte {
	return (* /*line FCssQXHZLA.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].LeafBlob()
}

func (qGFIOZJlpK8 *odcxjPIq) LeafProof() [][]byte {
	return (* /*line mLtpOd2ANtAk.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].LeafProof()
}

func (qGFIOZJlpK8 *odcxjPIq) Path() []byte {
	return (* /*line Ol7GFwj.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].Path()
}

func (qGFIOZJlpK8 *odcxjPIq) NodeBlob() []byte {
	return (* /*line C4zMqx.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].NodeBlob()
}

func (qGFIOZJlpK8 *odcxjPIq) AddResolver(aRra_C_9g NodeResolver) {
	/*line TiuVg_Hi3I.go:1*/ panic(func() /*line hD8pD2cilAk.go:1*/ string {
		data := [] /*line hD8pD2cilAk.go:1*/ byte("\x1a\xaf\x83\x87\xbe\xc7\x14\xb0e÷\xd3\xcd\xc7[")
		positions := [...]byte{12, 11, 2, 3, 13, 6, 2, 0, 0, 5, 13, 14, 9, 2, 7, 10, 13, 1, 4, 1}
		for i := 0; i < 20; i += 2 {
			localKey := /*line hD8pD2cilAk.go:1*/ byte(i) + /*line hD8pD2cilAk.go:1*/ byte(positions[i]^positions[i+1]) + 154
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line hD8pD2cilAk.go:1*/ string(data)
	}())
}

func (qGFIOZJlpK8 *odcxjPIq) Next(h6UJjHlET bool) bool {
	if /*line eTIDGcZVz.go:1*/ len(*qGFIOZJlpK8.k4xbZk0XYN63) == 0 {
		return false
	}

	dS536f21 := /*line GuIvw8hnaZf.go:1*/ heap.Pop(qGFIOZJlpK8.k4xbZk0XYN63).(FBy5J0gbkg)

	for /*line vTLhf53EM7Ub.go:1*/ len(*qGFIOZJlpK8.k4xbZk0XYN63) > 0 && ((!h6UJjHlET && /*line dlZzzD.go:1*/ bytes.HasPrefix((* /*line LjgiuT.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[0].Path() /*line Hhlw0acx.go:1*/, dS536f21.Path())) || /*line vBlUh2kTrPW.go:1*/ cSqBHdRTg3(dS536f21, (*qGFIOZJlpK8.k4xbZk0XYN63)[0]) == 0) {
		abjmscii := /*line VYgps8iOnLIx.go:1*/ heap.Pop(qGFIOZJlpK8.k4xbZk0XYN63).(FBy5J0gbkg)

		if /*line WBW6Fgcq0Zh.go:1*/ abjmscii.Next( /*line HrPZa6.go:1*/ abjmscii.Hash() == common.Hash{}) {
			qGFIOZJlpK8.h9a1udIyKRBJ++

			/*line F5cJbhu0.go:1*/
			heap.Push(qGFIOZJlpK8.k4xbZk0XYN63, abjmscii)
		}
	}
	if /*line W9HX4CPLz2.go:1*/ dS536f21.Next(h6UJjHlET) {
		qGFIOZJlpK8.h9a1udIyKRBJ++
		/*line vGu1YN.go:1*/ heap.Push(qGFIOZJlpK8.k4xbZk0XYN63, dS536f21)
	}
	return /*line _cc2UWXw1OI.go:1*/ len(*qGFIOZJlpK8.k4xbZk0XYN63) > 0
}

func (qGFIOZJlpK8 *odcxjPIq) Error() error {
	for q2u2020KD6 := 0; q2u2020KD6 < /*line yWHeufqQ3.go:1*/ len(*qGFIOZJlpK8.k4xbZk0XYN63); q2u2020KD6++ {
		if eZzE0KPS := (* /*line MEcDF22EmEew.go:1*/ qGFIOZJlpK8.k4xbZk0XYN63)[q2u2020KD6].Error(); eZzE0KPS != nil {
			return eZzE0KPS
		}
	}
	return nil
}

var _ bytes.Buffer
var _ = heap.Fix
var _ = errors.As
var _ types.AccessList
var _ = common.AbsolutePath
