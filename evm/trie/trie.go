//line :1

package trie

import (
	"bytes"
	"errors"
	"fmt"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/triedb/database"
)

type Trie struct {
	root  node
	owner common.Hash

	committed bool

	unhashed int

	reader *trieReader

	tracer *tracer
}

func (hkI2UXG_QBd *Trie) gp9ZtkuvtStq() nodeFlag {
	return nodeFlag{dirty: true}
}

func (hkI2UXG_QBd *Trie) Copy() *Trie {
	return &Trie{
		root:      hkI2UXG_QBd.root,
		owner:     hkI2UXG_QBd.owner,
		committed: hkI2UXG_QBd.committed,
		unhashed:  hkI2UXG_QBd.unhashed,
		reader:    hkI2UXG_QBd.reader,
		tracer:/*line sxz2515UbH5r.go:1*/ hkI2UXG_QBd.tracer.sBTiL7Ci(),
	}
}

func RJaQ3esB(cIFJQQh *ID, kRC_1aK database.Database) (*Trie, error) {
	dmVOxRYtQ, eZzE0KPS := /*line h2a35CEU8Qy.go:1*/ fHP68lzj9vD1(cIFJQQh.StateRoot, cIFJQQh.Owner, kRC_1aK)
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	gExqOlCxa53E := &Trie{
		owner:  cIFJQQh.Owner,
		reader: dmVOxRYtQ,
		tracer:/*line aWSCPUFf90YR.go:1*/ aLjRi9a0(),
	}
	if cIFJQQh.Root != (common.Hash{}) && cIFJQQh.Root != types.NrGuaNA21 {
		eHO3qzYODd, eZzE0KPS := /*line DNxNLA.go:1*/ gExqOlCxa53E.xIqJE3(cIFJQQh.Root[:], nil)
		if eZzE0KPS != nil {
			return nil, eZzE0KPS
		}
		gExqOlCxa53E.root = eHO3qzYODd
	}
	return gExqOlCxa53E, nil
}

func GysSMGsv(kRC_1aK database.Database) *Trie {
	rjqmOX5MHdh, _ := /*line vZaIiTpyUQ.go:1*/ RJaQ3esB( /*line U3SElpcd5C.go:1*/ K9sqwbQ(types.NrGuaNA21), kRC_1aK)
	return rjqmOX5MHdh
}

func (hkI2UXG_QBd *Trie) MustNodeIterator(iS1397t []byte) FBy5J0gbkg {
	qGFIOZJlpK8, eZzE0KPS := /*line wHmpx6YU65t.go:1*/ hkI2UXG_QBd.NodeIterator(iS1397t)
	if eZzE0KPS != nil {
		/*line yKFYTvoq5o.go:1*/ log.Error(func() /*line hzl14k.go:1*/ string {
			key := [] /*line hzl14k.go:1*/ byte("\xc3T\x90\x84\xed\xa4\x14\xd6!N\xb8T]Z\xea_//\x9f\b\xc8\f\x03\x92\x1f\xee\x1b\x0ep\xf2K\xab=W\au\xe7\x91&\x98\xfa")
			data := [] /*line hzl14k.go:1*/ byte("\x96:\xf8\xe5\x83\xc0x\xb3En\xcc&4?\xca:]]\xf0z\xe8em\xb2K\x9crk^\xbc$\xcfX\x1es\x10\x95\xf0R\xf7\x88")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line hzl14k.go:1*/ string(data)
		}(), "err", eZzE0KPS)
	}
	return qGFIOZJlpK8
}

func (hkI2UXG_QBd *Trie) NodeIterator(iS1397t []byte) (FBy5J0gbkg, error) {

	if hkI2UXG_QBd.committed {
		return nil, MWbaoG
	}
	return /*line anGzwG.go:1*/ u98WTduaeU2Y(hkI2UXG_QBd, iS1397t), nil
}

func (hkI2UXG_QBd *Trie) MustGet(lhQWH7m []byte) []byte {
	b6Gjyt0P8, eZzE0KPS := /*line Es73h0p.go:1*/ hkI2UXG_QBd.Get(lhQWH7m)
	if eZzE0KPS != nil {
		/*line bii6gpmo.go:1*/ log.Error(func() /*line NAKn76cw.go:1*/ string {
			data := /*line NAKn76cw.go:1*/ make([]byte, 0, 33)
			i := 0
			decryptKey := 153
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					data = /*line NAKn76cw.go:1*/ append(data, 39)
					i = 4
				case 11:
					data = /*line NAKn76cw.go:1*/ append(data, "\x16\x10\x16"...,
					)
					i = 2
				case 4:
					i = 8
					for y := range data {
						data[y] = data[y] - /*line NAKn76cw.go:1*/ byte(decryptKey^y)
					}
				case 1:
					data = /*line NAKn76cw.go:1*/ append(data, 27)
					i = 11
				case 0:
					i = 1
					data = /*line NAKn76cw.go:1*/ append(data, 1)
				case 7:
					data = /*line NAKn76cw.go:1*/ append(data, "-1\xd8\""...,
					)
					i = 12
				case 9:
					i = 13
					data = /*line NAKn76cw.go:1*/ append(data, "\x19\t\x06\xc2"...,
					)
				case 10:
					data = /*line NAKn76cw.go:1*/ append(data, "\x1f\x1c"...,
					)
					i = 5
				case 2:
					data = /*line NAKn76cw.go:1*/ append(data, "\r\x16"...,
					)
					i = 3
				case 5:
					data = /*line NAKn76cw.go:1*/ append(data, "\xde\xf8\x17"...,
					)
					i = 6
				case 3:
					i = 9
					data = /*line NAKn76cw.go:1*/ append(data, "\x10\b\xc5\x1a"...,
					)
				case 13:
					i = 7
					data = /*line NAKn76cw.go:1*/ append(data, "\b./"...,
					)
				case 12:
					data = /*line NAKn76cw.go:1*/ append(data, "(\xdb\b'"...,
					)
					i = 10
				}
			}
			return /*line NAKn76cw.go:1*/ string(data)
		}(), "err", eZzE0KPS)
	}
	return b6Gjyt0P8
}

func (hkI2UXG_QBd *Trie) Get(lhQWH7m []byte) ([]byte, error) {

	if hkI2UXG_QBd.committed {
		return nil, MWbaoG
	}
	ar76Sw, bkYJpPY, kdibOGeF, eZzE0KPS := /*line YAaCHXY.go:1*/ hkI2UXG_QBd.cga2ibNy(hkI2UXG_QBd.root /*line KVMT5ouaid4.go:1*/, nd12pU880Z8(lhQWH7m), 0)
	if eZzE0KPS == nil && kdibOGeF {
		hkI2UXG_QBd.root = bkYJpPY
	}
	return ar76Sw, eZzE0KPS
}

func (hkI2UXG_QBd *Trie) cga2ibNy(uJUzjlW8 node, lhQWH7m []byte, wMJhBVQ int) (ar76Sw []byte, eT6HmXgwbfB node, kdibOGeF bool, eZzE0KPS error) {
	switch gnGMaXTuiFeE := (uJUzjlW8).(type) {
	case nil:
		return nil, nil, false, nil
	case valueNode:
		return gnGMaXTuiFeE, gnGMaXTuiFeE, false, nil
	case *qUKQUCfTXB:
		if /*line QuOBEi.go:1*/ len(lhQWH7m)-wMJhBVQ < /*line JDfA23Y.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A) || ! /*line tXzlmB.go:1*/ bytes.Equal(gnGMaXTuiFeE.ANdZYqbzJ1A, lhQWH7m[wMJhBVQ:wMJhBVQ+ /*line HYTP7wi.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A)]) {

			return nil, gnGMaXTuiFeE, false, nil
		}
		ar76Sw, eT6HmXgwbfB, kdibOGeF, eZzE0KPS = /*line zAe3MGq5dHTs.go:1*/ hkI2UXG_QBd.cga2ibNy(gnGMaXTuiFeE.YpmEYGB, lhQWH7m, wMJhBVQ+ /*line m9wAcTy3t6fs.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A))
		if eZzE0KPS == nil && kdibOGeF {
			gnGMaXTuiFeE = /*line iZUZWIXAzDpB.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
			gnGMaXTuiFeE.YpmEYGB = eT6HmXgwbfB
		}
		return ar76Sw, gnGMaXTuiFeE, kdibOGeF, eZzE0KPS
	case *fullNode:
		ar76Sw, eT6HmXgwbfB, kdibOGeF, eZzE0KPS = /*line ysW1a8u_7z3.go:1*/ hkI2UXG_QBd.cga2ibNy(gnGMaXTuiFeE.Children[lhQWH7m[wMJhBVQ]], lhQWH7m, wMJhBVQ+1)
		if eZzE0KPS == nil && kdibOGeF {
			gnGMaXTuiFeE = /*line AwUbnXnKDj.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
			gnGMaXTuiFeE.Children[lhQWH7m[wMJhBVQ]] = eT6HmXgwbfB
		}
		return ar76Sw, gnGMaXTuiFeE, kdibOGeF, eZzE0KPS
	case hashNode:
		jcDLabJ7o, eZzE0KPS := /*line YzYDQbtRioY.go:1*/ hkI2UXG_QBd.xIqJE3(gnGMaXTuiFeE, lhQWH7m[:wMJhBVQ])
		if eZzE0KPS != nil {
			return nil, gnGMaXTuiFeE, true, eZzE0KPS
		}
		ar76Sw, eT6HmXgwbfB, _, eZzE0KPS := /*line aYDXPb.go:1*/ hkI2UXG_QBd.cga2ibNy(jcDLabJ7o, lhQWH7m, wMJhBVQ)
		return ar76Sw, eT6HmXgwbfB, true, eZzE0KPS
	default:
		/*line ayuLC_MB.go:1*/ panic( /*line BZBe0Bl.go:1*/ fmt.Sprintf(func() /*line eoAnfr.go:1*/ string {
			fullData := [] /*line eoAnfr.go:1*/ byte("A\xd5\x7fp\"\xbds.J\x7f=\x1e\x16\x9c\f\xd4H\x05\x8f\xe59\f\xe4\xf73L9\x94|\xab-Q_\xa1\xa6\nz\xef\xca\xfd")
			idxKey := [] /*line eoAnfr.go:1*/ byte(";m")
			data := /*line eoAnfr.go:1*/ make([]byte, 0, 21)
			data = /*line eoAnfr.go:1*/ append(data, fullData[45^ /*line eoAnfr.go:1*/ int(idxKey[0])]+fullData[59^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[99^ /*line eoAnfr.go:1*/ int(idxKey[1])]+fullData[125^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[58^ /*line eoAnfr.go:1*/ int(idxKey[0])]^fullData[30^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[78^ /*line eoAnfr.go:1*/ int(idxKey[1])]+fullData[97^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[28^ /*line eoAnfr.go:1*/ int(idxKey[0])]^fullData[32^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[61^ /*line eoAnfr.go:1*/ int(idxKey[0])]-fullData[42^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[100^ /*line eoAnfr.go:1*/ int(idxKey[1])]+fullData[122^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[102^ /*line eoAnfr.go:1*/ int(idxKey[1])]-fullData[104^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[36^ /*line eoAnfr.go:1*/ int(idxKey[0])]-fullData[40^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[35^ /*line eoAnfr.go:1*/ int(idxKey[0])]-fullData[29^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[51^ /*line eoAnfr.go:1*/ int(idxKey[0])]^fullData[60^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[73^ /*line eoAnfr.go:1*/ int(idxKey[1])]+fullData[79^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[116^ /*line eoAnfr.go:1*/ int(idxKey[1])]+fullData[105^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[96^ /*line eoAnfr.go:1*/ int(idxKey[1])]-fullData[115^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[56^ /*line eoAnfr.go:1*/ int(idxKey[0])]-fullData[46^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[119^ /*line eoAnfr.go:1*/ int(idxKey[1])]-fullData[98^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[127^ /*line eoAnfr.go:1*/ int(idxKey[1])]+fullData[112^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[57^ /*line eoAnfr.go:1*/ int(idxKey[0])]-fullData[27^ /*line eoAnfr.go:1*/ int(idxKey[0])], fullData[76^ /*line eoAnfr.go:1*/ int(idxKey[1])]-fullData[113^ /*line eoAnfr.go:1*/ int(idxKey[1])], fullData[49^ /*line eoAnfr.go:1*/ int(idxKey[0])]+fullData[47^ /*line eoAnfr.go:1*/ int(idxKey[0])])
			return /*line eoAnfr.go:1*/ string(data)
		}(), uJUzjlW8, uJUzjlW8))
	}
}

func (hkI2UXG_QBd *Trie) MustGetNode(qiRG6lpeaFl []byte) ([]byte, int) {
	bAufRbtZ9MFz, lR8wDJYky, eZzE0KPS := /*line L_PkKVE.go:1*/ hkI2UXG_QBd.GetNode(qiRG6lpeaFl)
	if eZzE0KPS != nil {
		/*line QVzNDaSt.go:1*/ log.Error(func() /*line gNQDo8Ngsr0R.go:1*/ string {
			data := [] /*line gNQDo8Ngsr0R.go:1*/ byte("\x97n]\x19b\x8bD}T\x8agg\x8ae\x1c}\xc1U{S i\x1a]l}i\x85܀y\x8c\xa6\x92de")
			positions := [...]byte{8, 11, 6, 3, 19, 23, 2, 4, 3, 24, 30, 15, 17, 0, 7, 31, 33, 9, 14, 9, 10, 2, 22, 23, 12, 17, 25, 5, 30, 16, 3, 29, 32, 30, 15, 11, 6, 14, 32, 28, 32, 8, 27, 18, 7, 12}
			for i := 0; i < 46; i += 2 {
				localKey := /*line gNQDo8Ngsr0R.go:1*/ byte(i) + /*line gNQDo8Ngsr0R.go:1*/ byte(positions[i]^positions[i+1]) + 227
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line gNQDo8Ngsr0R.go:1*/ string(data)
		}(), "err", eZzE0KPS)
	}
	return bAufRbtZ9MFz, lR8wDJYky
}

func (hkI2UXG_QBd *Trie) GetNode(qiRG6lpeaFl []byte) ([]byte, int, error) {

	if hkI2UXG_QBd.committed {
		return nil, 0, MWbaoG
	}
	bAufRbtZ9MFz, bkYJpPY, lR8wDJYky, eZzE0KPS := /*line x_dGabpSYVD.go:1*/ hkI2UXG_QBd.cdTba5HX(hkI2UXG_QBd.root /*line tnZ3Lg.go:1*/, o9hF7yS(qiRG6lpeaFl), 0)
	if eZzE0KPS != nil {
		return nil, lR8wDJYky, eZzE0KPS
	}
	if lR8wDJYky > 0 {
		hkI2UXG_QBd.root = bkYJpPY
	}
	if bAufRbtZ9MFz == nil {
		return nil, lR8wDJYky, nil
	}
	return bAufRbtZ9MFz, lR8wDJYky, nil
}

func (hkI2UXG_QBd *Trie) cdTba5HX(uJUzjlW8 node, qiRG6lpeaFl []byte, wMJhBVQ int) (bAufRbtZ9MFz []byte, eT6HmXgwbfB node, lR8wDJYky int, eZzE0KPS error) {

	if uJUzjlW8 == nil {
		return nil, nil, 0, nil
	}

	if wMJhBVQ >= /*line C5uTNH.go:1*/ len(qiRG6lpeaFl) {

		var rNuN0sMPDJ hashNode
		if uAjz8NxU_cL, yY_yPSd8vG := uJUzjlW8.(hashNode); yY_yPSd8vG {
			rNuN0sMPDJ = uAjz8NxU_cL
		} else {
			rNuN0sMPDJ, _ = /*line wn7Ali.go:1*/ uJUzjlW8.tGJzZYYLvK()
		}
		if rNuN0sMPDJ == nil {
			return nil, uJUzjlW8, 0 /*line DCJ6_U.go:1*/, errors.New(func() /*line hoUSSYv3K.go:1*/ string {
				data := [] /*line hoUSSYv3K.go:1*/ byte("\xf1\xc5\xc5\x00\xc6o\xcbsj\xcbH3s\xa4Woٞ")
				positions := [...]byte{11, 10, 8, 1, 14, 16, 4, 9, 1, 6, 0, 0, 3, 17, 16, 13, 17, 10, 11, 2}
				for i := 0; i < 20; i += 2 {
					localKey := /*line hoUSSYv3K.go:1*/ byte(i) + /*line hoUSSYv3K.go:1*/ byte(positions[i]^positions[i+1]) + 149
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line hoUSSYv3K.go:1*/ string(data)
			}())
		}
		aBHduJG, eZzE0KPS := /*line WYQFRkyex.go:1*/ hkI2UXG_QBd.reader.uAjz8NxU_cL(qiRG6lpeaFl /*line Rs0awzQg0.go:1*/, common.BytesToHash(rNuN0sMPDJ))
		return aBHduJG, uJUzjlW8, 1, eZzE0KPS
	}

	switch gnGMaXTuiFeE := (uJUzjlW8).(type) {
	case valueNode:

		return nil, nil, 0, nil

	case *qUKQUCfTXB:
		if /*line k8PWPN.go:1*/ len(qiRG6lpeaFl)-wMJhBVQ < /*line i278zbzWdmf.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A) || ! /*line Woz4_MP74Q6.go:1*/ bytes.Equal(gnGMaXTuiFeE.ANdZYqbzJ1A, qiRG6lpeaFl[wMJhBVQ:wMJhBVQ+ /*line RBEHm9lsy.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A)]) {

			return nil, gnGMaXTuiFeE, 0, nil
		}
		bAufRbtZ9MFz, eT6HmXgwbfB, lR8wDJYky, eZzE0KPS = /*line Jxb1PK8.go:1*/ hkI2UXG_QBd.cdTba5HX(gnGMaXTuiFeE.YpmEYGB, qiRG6lpeaFl, wMJhBVQ+ /*line EEjRUfG.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A))
		if eZzE0KPS == nil && lR8wDJYky > 0 {
			gnGMaXTuiFeE = /*line GEuWV2Iuu.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
			gnGMaXTuiFeE.YpmEYGB = eT6HmXgwbfB
		}
		return bAufRbtZ9MFz, gnGMaXTuiFeE, lR8wDJYky, eZzE0KPS

	case *fullNode:
		bAufRbtZ9MFz, eT6HmXgwbfB, lR8wDJYky, eZzE0KPS = /*line AywCMaMliVr.go:1*/ hkI2UXG_QBd.cdTba5HX(gnGMaXTuiFeE.Children[qiRG6lpeaFl[wMJhBVQ]], qiRG6lpeaFl, wMJhBVQ+1)
		if eZzE0KPS == nil && lR8wDJYky > 0 {
			gnGMaXTuiFeE = /*line gA5EjBUusF.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
			gnGMaXTuiFeE.Children[qiRG6lpeaFl[wMJhBVQ]] = eT6HmXgwbfB
		}
		return bAufRbtZ9MFz, gnGMaXTuiFeE, lR8wDJYky, eZzE0KPS

	case hashNode:
		jcDLabJ7o, eZzE0KPS := /*line NebCxxoAy.go:1*/ hkI2UXG_QBd.xIqJE3(gnGMaXTuiFeE, qiRG6lpeaFl[:wMJhBVQ])
		if eZzE0KPS != nil {
			return nil, gnGMaXTuiFeE, 1, eZzE0KPS
		}
		bAufRbtZ9MFz, eT6HmXgwbfB, lR8wDJYky, eZzE0KPS := /*line ESNDie.go:1*/ hkI2UXG_QBd.cdTba5HX(jcDLabJ7o, qiRG6lpeaFl, wMJhBVQ)
		return bAufRbtZ9MFz, eT6HmXgwbfB, lR8wDJYky + 1, eZzE0KPS

	default:
		/*line H55Oca6LFU4.go:1*/ panic( /*line tpzWQ3A.go:1*/ fmt.Sprintf(func() /*line Va6wW2ZK_2S.go:1*/ string {
			fullData := [] /*line Va6wW2ZK_2S.go:1*/ byte("\x11\x8d\x1e\x05\x86\r#p@-\x0eA'%\x18\xd2>\xafd\xbcp\r\x9c\x88\xa4n\xd1\xef24\x02;Sms\xa1\x84\\\xb2'")
			idxKey := [] /*line Va6wW2ZK_2S.go:1*/ byte("\x02\x88\xd4\xd8\xe9}\x9d\x17$2=\x9dX")
			data := /*line Va6wW2ZK_2S.go:1*/ make([]byte, 0, 21)
			data = /*line Va6wW2ZK_2S.go:1*/ append(data, fullData[214^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[3])]+fullData[221^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[3])], fullData[84^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])]^fullData[122^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])], fullData[210^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[3])]^fullData[197^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[3])], fullData[142^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[6])]-fullData[139^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[6])], fullData[54^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[8])]^fullData[49^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[8])], fullData[68^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])]^fullData[125^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])], fullData[34^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[8])]+fullData[4^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[8])], fullData[50^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[9])]^fullData[38^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[9])], fullData[148^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[11])]^fullData[150^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[11])], fullData[147^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[1])]^fullData[140^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[1])], fullData[91^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])]-fullData[123^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])], fullData[232^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[4])]-fullData[200^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[4])], fullData[58^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[10])]-fullData[35^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[10])], fullData[149^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[6])]-fullData[135^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[6])], fullData[13^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[0])]-fullData[27^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[0])], fullData[152^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[1])]+fullData[175^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[1])], fullData[187^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[11])]+fullData[138^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[11])], fullData[124^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])]^fullData[64^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[12])], fullData[235^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[4])]^fullData[246^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[4])], fullData[217^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[2])]-fullData[197^ /*line Va6wW2ZK_2S.go:1*/ int(idxKey[2])])
			return /*line Va6wW2ZK_2S.go:1*/ string(data)
		}(), uJUzjlW8, uJUzjlW8))
	}
}

func (hkI2UXG_QBd *Trie) MustUpdate(lhQWH7m, ar76Sw []byte) {
	if eZzE0KPS := /*line kGjk0ISuIZ.go:1*/ hkI2UXG_QBd.Update(lhQWH7m, ar76Sw); eZzE0KPS != nil {
		/*line Fg59QuR.go:1*/ log.Error(func() /*line qY99Gk8e5W.go:1*/ string {
			data := [] /*line qY99Gk8e5W.go:1*/ byte("\x06n\xee\xf5n\x7f\xae\x89=\x01\xa6ri\x94 \xdb\xe7r\x98r\xf8\xa2ѐ\xf6\x96i\x88@\xc3\xf0da[$")
			positions := [...]byte{22, 18, 25, 18, 16, 34, 21, 0, 10, 5, 27, 9, 20, 29, 28, 23, 18, 0, 23, 16, 16, 8, 7, 5, 33, 9, 15, 24, 15, 30, 6, 28, 16, 23, 6, 8, 6, 13, 5, 8, 6, 30, 2, 3, 22, 10, 30, 9}
			for i := 0; i < 48; i += 2 {
				localKey := /*line qY99Gk8e5W.go:1*/ byte(i) + /*line qY99Gk8e5W.go:1*/ byte(positions[i]^positions[i+1]) + 72
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line qY99Gk8e5W.go:1*/ string(data)
		}(), "err", eZzE0KPS)
	}
}

func (hkI2UXG_QBd *Trie) Update(lhQWH7m, ar76Sw []byte) error {

	if hkI2UXG_QBd.committed {
		return MWbaoG
	}
	return /*line khmylRi3ua.go:1*/ hkI2UXG_QBd.lnptOb8HyLD(lhQWH7m, ar76Sw)
}

func (hkI2UXG_QBd *Trie) lnptOb8HyLD(lhQWH7m, ar76Sw []byte) error {
	hkI2UXG_QBd.unhashed++
	ec_diEU := /*line CJKoCasp5Rf.go:1*/ nd12pU880Z8(lhQWH7m)
	if /*line CahcaihPWoq.go:1*/ len(ar76Sw) != 0 {
		_, gnGMaXTuiFeE, eZzE0KPS := /*line Q_jp3zV.go:1*/ hkI2UXG_QBd.bfEQU4T(hkI2UXG_QBd.root, nil, ec_diEU /*line nA0Rf4AB.go:1*/, valueNode(ar76Sw))
		if eZzE0KPS != nil {
			return eZzE0KPS
		}
		hkI2UXG_QBd.root = gnGMaXTuiFeE
	} else {
		_, gnGMaXTuiFeE, eZzE0KPS := /*line aZeIAY.go:1*/ hkI2UXG_QBd.mWPnovPMg4(hkI2UXG_QBd.root, nil, ec_diEU)
		if eZzE0KPS != nil {
			return eZzE0KPS
		}
		hkI2UXG_QBd.root = gnGMaXTuiFeE
	}
	return nil
}

func (hkI2UXG_QBd *Trie) bfEQU4T(gnGMaXTuiFeE node, tKBrjZ, lhQWH7m []byte, ar76Sw node) (bool, node, error) {
	if /*line yc5Er0JKAH.go:1*/ len(lhQWH7m) == 0 {
		if eBWWX0caQi, yY_yPSd8vG := gnGMaXTuiFeE.(valueNode); yY_yPSd8vG {
			return ! /*line hyr52n.go:1*/ bytes.Equal(eBWWX0caQi, ar76Sw.(valueNode)), ar76Sw, nil
		}
		return true, ar76Sw, nil
	}
	switch gnGMaXTuiFeE := gnGMaXTuiFeE.(type) {
	case *qUKQUCfTXB:
		imXgnpQ5W := /*line a30F1ZKae.go:1*/ e0augA(lhQWH7m, gnGMaXTuiFeE.ANdZYqbzJ1A)

		if imXgnpQ5W == /*line CsvxNI.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A) {
			c0YPpA0vA5f, nqnEea, eZzE0KPS := /*line c7z6DorDHN2.go:1*/ hkI2UXG_QBd.bfEQU4T(gnGMaXTuiFeE.YpmEYGB /*line XiNpMEBF1M.go:1*/, append(tKBrjZ, lhQWH7m[:imXgnpQ5W]...), lhQWH7m[imXgnpQ5W:], ar76Sw)
			if !c0YPpA0vA5f || eZzE0KPS != nil {
				return false, gnGMaXTuiFeE, eZzE0KPS
			}
			return true, &qUKQUCfTXB{gnGMaXTuiFeE.ANdZYqbzJ1A, nqnEea /*line A5bqXg.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil
		}

		cT6QHM64Leu := &fullNode{flags: /*line BYuBqVfH1.go:1*/ hkI2UXG_QBd.gp9ZtkuvtStq()}
		var eZzE0KPS error
		_, cT6QHM64Leu.Children[gnGMaXTuiFeE.ANdZYqbzJ1A[imXgnpQ5W]], eZzE0KPS = /*line DUk36_E.go:1*/ hkI2UXG_QBd.bfEQU4T(nil /*line xBnqFv_kRSJ0.go:1*/, append(tKBrjZ, gnGMaXTuiFeE.ANdZYqbzJ1A[:imXgnpQ5W+1]...), gnGMaXTuiFeE.ANdZYqbzJ1A[imXgnpQ5W+1:], gnGMaXTuiFeE.YpmEYGB)
		if eZzE0KPS != nil {
			return false, nil, eZzE0KPS
		}
		_, cT6QHM64Leu.Children[lhQWH7m[imXgnpQ5W]], eZzE0KPS = /*line N47rGvU1MJN.go:1*/ hkI2UXG_QBd.bfEQU4T(nil /*line B1Hala1T.go:1*/, append(tKBrjZ, lhQWH7m[:imXgnpQ5W+1]...), lhQWH7m[imXgnpQ5W+1:], ar76Sw)
		if eZzE0KPS != nil {
			return false, nil, eZzE0KPS
		}

		if imXgnpQ5W == 0 {
			return true, cT6QHM64Leu, nil
		}

		/*line YTQwPF3TSwW.go:1*/
		hkI2UXG_QBd.tracer.har2k8rr3s( /*line CIy830aEo.go:1*/ append(tKBrjZ, lhQWH7m[:imXgnpQ5W]...))

		return true, &qUKQUCfTXB{lhQWH7m[:imXgnpQ5W], cT6QHM64Leu /*line dhJPcn.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil

	case *fullNode:
		c0YPpA0vA5f, nqnEea, eZzE0KPS := /*line zLRVqpx.go:1*/ hkI2UXG_QBd.bfEQU4T(gnGMaXTuiFeE.Children[lhQWH7m[0]] /*line KHkLTfZ.go:1*/, append(tKBrjZ, lhQWH7m[0]), lhQWH7m[1:], ar76Sw)
		if !c0YPpA0vA5f || eZzE0KPS != nil {
			return false, gnGMaXTuiFeE, eZzE0KPS
		}
		gnGMaXTuiFeE = /*line aCaDC5N.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
		gnGMaXTuiFeE.flags = /*line A8YTsPD1.go:1*/ hkI2UXG_QBd.gp9ZtkuvtStq()
		gnGMaXTuiFeE.Children[lhQWH7m[0]] = nqnEea
		return true, gnGMaXTuiFeE, nil

	case nil:

		/*line URv9qYa4.go:1*/
		hkI2UXG_QBd.tracer.har2k8rr3s(tKBrjZ)

		return true, &qUKQUCfTXB{lhQWH7m, ar76Sw /*line z5kq0aa.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil

	case hashNode:

		mjh9fYwPqBOf, eZzE0KPS := /*line KknXkbh33P.go:1*/ hkI2UXG_QBd.xIqJE3(gnGMaXTuiFeE, tKBrjZ)
		if eZzE0KPS != nil {
			return false, nil, eZzE0KPS
		}
		c0YPpA0vA5f, nqnEea, eZzE0KPS := /*line XARkwB6c8.go:1*/ hkI2UXG_QBd.bfEQU4T(mjh9fYwPqBOf, tKBrjZ, lhQWH7m, ar76Sw)
		if !c0YPpA0vA5f || eZzE0KPS != nil {
			return false, mjh9fYwPqBOf, eZzE0KPS
		}
		return true, nqnEea, nil

	default:
		/*line Q578HR2.go:1*/ panic( /*line pUmdkvt4.go:1*/ fmt.Sprintf(func() /*line ODWTPu85GF.go:1*/ string {
			data := [] /*line ODWTPu85GF.go:1*/ byte("0\f\xf0\x9e\x01\xcevali/\xb4n1d\xe2&\r\xe7v")
			positions := [...]byte{3, 11, 17, 15, 18, 13, 1, 16, 0, 4, 18, 13, 0, 5, 3, 11, 10, 10, 15, 15, 3, 2, 17, 0}
			for i := 0; i < 24; i += 2 {
				localKey := /*line ODWTPu85GF.go:1*/ byte(i) + /*line ODWTPu85GF.go:1*/ byte(positions[i]^positions[i+1]) + 187
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line ODWTPu85GF.go:1*/ string(data)
		}(), gnGMaXTuiFeE, gnGMaXTuiFeE))
	}
}

func (hkI2UXG_QBd *Trie) MustDelete(lhQWH7m []byte) {
	if eZzE0KPS := /*line Ft9Wh5Nqre8.go:1*/ hkI2UXG_QBd.Delete(lhQWH7m); eZzE0KPS != nil {
		/*line Rtz6wrHuGNp.go:1*/ log.Error(func() /*line PslzrnR.go:1*/ string {
			fullData := [] /*line PslzrnR.go:1*/ byte(")A\xb9\x90O\r3˰@\x17d\x89\u07bb\xa2\xc6w\x85M\xb5\x1c\x9b\xefY\x8a\xdaW\xd4\xf3\x93\xb7\xd1l\x82\x16\xdbA\x18\xe9\tN)a%fl̔\xda$M\x98\x03\xbf\x92\xf4\bL\xd5\x1d\xd6\xd1a\xb4\xf6\x9a\xc9\xccT")
			idxKey := [] /*line PslzrnR.go:1*/ byte("'\xbf")
			data := /*line PslzrnR.go:1*/ make([]byte, 0, 36)
			data = /*line PslzrnR.go:1*/ append(data, fullData[99^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[54^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[13^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[41^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[29^ /*line PslzrnR.go:1*/ int(idxKey[0])]+fullData[50^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[23^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[33^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[24^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[58^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[161^ /*line PslzrnR.go:1*/ int(idxKey[1])]+fullData[159^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[56^ /*line PslzrnR.go:1*/ int(idxKey[0])]+fullData[51^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[188^ /*line PslzrnR.go:1*/ int(idxKey[1])]+fullData[132^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[62^ /*line PslzrnR.go:1*/ int(idxKey[0])]+fullData[61^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[31^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[59^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[182^ /*line PslzrnR.go:1*/ int(idxKey[1])]-fullData[144^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[172^ /*line PslzrnR.go:1*/ int(idxKey[1])]+fullData[147^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[32^ /*line PslzrnR.go:1*/ int(idxKey[0])]^fullData[40^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[17^ /*line PslzrnR.go:1*/ int(idxKey[0])]^fullData[22^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[12^ /*line PslzrnR.go:1*/ int(idxKey[0])]^fullData[38^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[10^ /*line PslzrnR.go:1*/ int(idxKey[0])]^fullData[18^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[100^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[60^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[43^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[45^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[4^ /*line PslzrnR.go:1*/ int(idxKey[0])]+fullData[63^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[140^ /*line PslzrnR.go:1*/ int(idxKey[1])]-fullData[155^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[173^ /*line PslzrnR.go:1*/ int(idxKey[1])]+fullData[169^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[34^ /*line PslzrnR.go:1*/ int(idxKey[0])]^fullData[44^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[134^ /*line PslzrnR.go:1*/ int(idxKey[1])]-fullData[253^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[39^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[15^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[145^ /*line PslzrnR.go:1*/ int(idxKey[1])]-fullData[153^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[42^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[6^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[189^ /*line PslzrnR.go:1*/ int(idxKey[1])]+fullData[183^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[14^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[0^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[55^ /*line PslzrnR.go:1*/ int(idxKey[0])]-fullData[19^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[130^ /*line PslzrnR.go:1*/ int(idxKey[1])]^fullData[136^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[255^ /*line PslzrnR.go:1*/ int(idxKey[1])]^fullData[129^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[35^ /*line PslzrnR.go:1*/ int(idxKey[0])]+fullData[27^ /*line PslzrnR.go:1*/ int(idxKey[0])], fullData[154^ /*line PslzrnR.go:1*/ int(idxKey[1])]+fullData[141^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[254^ /*line PslzrnR.go:1*/ int(idxKey[1])]-fullData[157^ /*line PslzrnR.go:1*/ int(idxKey[1])], fullData[250^ /*line PslzrnR.go:1*/ int(idxKey[1])]-fullData[168^ /*line PslzrnR.go:1*/ int(idxKey[1])])
			return /*line PslzrnR.go:1*/ string(data)
		}(), "err", eZzE0KPS)
	}
}

func (hkI2UXG_QBd *Trie) Delete(lhQWH7m []byte) error {

	if hkI2UXG_QBd.committed {
		return MWbaoG
	}
	hkI2UXG_QBd.unhashed++
	ec_diEU := /*line dnaLGhk_Qvl.go:1*/ nd12pU880Z8(lhQWH7m)
	_, gnGMaXTuiFeE, eZzE0KPS := /*line v7ZsVChvjwO.go:1*/ hkI2UXG_QBd.mWPnovPMg4(hkI2UXG_QBd.root, nil, ec_diEU)
	if eZzE0KPS != nil {
		return eZzE0KPS
	}
	hkI2UXG_QBd.root = gnGMaXTuiFeE
	return nil
}

func (hkI2UXG_QBd *Trie) mWPnovPMg4(gnGMaXTuiFeE node, tKBrjZ, lhQWH7m []byte) (bool, node, error) {
	switch gnGMaXTuiFeE := gnGMaXTuiFeE.(type) {
	case *qUKQUCfTXB:
		imXgnpQ5W := /*line u9S8xitiU.go:1*/ e0augA(lhQWH7m, gnGMaXTuiFeE.ANdZYqbzJ1A)
		if imXgnpQ5W < /*line hG9Z5_pN.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A) {
			return false, gnGMaXTuiFeE, nil
		}
		if imXgnpQ5W == /*line wa7M5SmBes.go:1*/ len(lhQWH7m) {

			/*line Of38sUY.go:1*/
			hkI2UXG_QBd.tracer.xwjztifet3(tKBrjZ)

			return true, nil, nil
		}

		c0YPpA0vA5f, jcDLabJ7o, eZzE0KPS := /*line hazE4KXN8y7p.go:1*/ hkI2UXG_QBd.mWPnovPMg4(gnGMaXTuiFeE.YpmEYGB /*line AaNlrBJ.go:1*/, append(tKBrjZ, lhQWH7m[: /*line SbZ3Ot5Uk.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A)]...), lhQWH7m[ /*line SlkVoVE6hk.go:1*/ len(gnGMaXTuiFeE.ANdZYqbzJ1A):])
		if !c0YPpA0vA5f || eZzE0KPS != nil {
			return false, gnGMaXTuiFeE, eZzE0KPS
		}
		switch jcDLabJ7o := jcDLabJ7o.(type) {
		case *qUKQUCfTXB:

			/*line HnGiNFs21C00.go:1*/
			hkI2UXG_QBd.tracer.xwjztifet3( /*line fhgejXy.go:1*/ append(tKBrjZ, gnGMaXTuiFeE.ANdZYqbzJ1A...))

			return true, &qUKQUCfTXB{ /*line lRVa25j4Laf.go:1*/ x_TVg_Nwq3L(gnGMaXTuiFeE.ANdZYqbzJ1A, jcDLabJ7o.ANdZYqbzJ1A...), jcDLabJ7o.YpmEYGB /*line Ye5TOCda56e.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil
		default:
			return true, &qUKQUCfTXB{gnGMaXTuiFeE.ANdZYqbzJ1A, jcDLabJ7o /*line vl8BWKA_rorp.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil
		}

	case *fullNode:
		c0YPpA0vA5f, nqnEea, eZzE0KPS := /*line QCY8l118dC.go:1*/ hkI2UXG_QBd.mWPnovPMg4(gnGMaXTuiFeE.Children[lhQWH7m[0]] /*line rUn96pWW8e5.go:1*/, append(tKBrjZ, lhQWH7m[0]), lhQWH7m[1:])
		if !c0YPpA0vA5f || eZzE0KPS != nil {
			return false, gnGMaXTuiFeE, eZzE0KPS
		}
		gnGMaXTuiFeE = /*line JaLZMdUYZdi.go:1*/ gnGMaXTuiFeE.sBTiL7Ci()
		gnGMaXTuiFeE.flags = /*line GHz59RaYd.go:1*/ hkI2UXG_QBd.gp9ZtkuvtStq()
		gnGMaXTuiFeE.Children[lhQWH7m[0]] = nqnEea

		if nqnEea != nil {
			return true, gnGMaXTuiFeE, nil
		}

		wMJhBVQ := -1
		for q2u2020KD6, nPjnWCe := range &gnGMaXTuiFeE.Children {
			if nPjnWCe != nil {
				if wMJhBVQ == -1 {
					wMJhBVQ = q2u2020KD6
				} else {
					wMJhBVQ = -2
					break
				}
			}
		}
		if wMJhBVQ >= 0 {
			if wMJhBVQ != 16 {

				baPHVYk, eZzE0KPS := /*line aWSjFSjeNJC.go:1*/ hkI2UXG_QBd.v2ULJu(gnGMaXTuiFeE.Children[wMJhBVQ] /*line mhpyrA.go:1*/, append(tKBrjZ /*line XkmcZ_EED.go:1*/, byte(wMJhBVQ)))
				if eZzE0KPS != nil {
					return false, nil, eZzE0KPS
				}
				if baPHVYk, yY_yPSd8vG := baPHVYk.(*qUKQUCfTXB); yY_yPSd8vG {

					/*line MXnYxhfOL.go:1*/
					hkI2UXG_QBd.tracer.xwjztifet3( /*line my00rl.go:1*/ append(tKBrjZ /*line iGugFU7bDd.go:1*/, byte(wMJhBVQ)))

					ec_diEU := /*line X5laAZBg4oGk.go:1*/ append([]byte{ /*line Z4aB8J_QTeK.go:1*/ byte(wMJhBVQ)}, baPHVYk.ANdZYqbzJ1A...)
					return true, &qUKQUCfTXB{ec_diEU, baPHVYk.YpmEYGB /*line JAer2Fj.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil
				}
			}

			return true, &qUKQUCfTXB{[]byte{ /*line N7MaQ68n6LOg.go:1*/ byte(wMJhBVQ)}, gnGMaXTuiFeE.Children[wMJhBVQ] /*line TWuOypl6JaG.go:1*/, hkI2UXG_QBd.gp9ZtkuvtStq()}, nil
		}

		return true, gnGMaXTuiFeE, nil

	case valueNode:
		return true, nil, nil

	case nil:
		return false, nil, nil

	case hashNode:

		mjh9fYwPqBOf, eZzE0KPS := /*line _Tz7QT32iJ.go:1*/ hkI2UXG_QBd.xIqJE3(gnGMaXTuiFeE, tKBrjZ)
		if eZzE0KPS != nil {
			return false, nil, eZzE0KPS
		}
		c0YPpA0vA5f, nqnEea, eZzE0KPS := /*line z9hyoNLIyKO.go:1*/ hkI2UXG_QBd.mWPnovPMg4(mjh9fYwPqBOf, tKBrjZ, lhQWH7m)
		if !c0YPpA0vA5f || eZzE0KPS != nil {
			return false, mjh9fYwPqBOf, eZzE0KPS
		}
		return true, nqnEea, nil

	default:
		/*line b7Raso.go:1*/ panic( /*line i4MVoPa.go:1*/ fmt.Sprintf(func() /*line SSMkrbGtha.go:1*/ string {
			data := /*line SSMkrbGtha.go:1*/ make([]byte, 0, 26)
			i := 4
			decryptKey := 32
			for counter := 0; i != 11; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xb3\xf6"...,
					)
					i = 10
				case 6:
					i = 2
					data = /*line SSMkrbGtha.go:1*/ append(data, 182)
				case 3:
					i = 6
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xfe\xb0"...,
					)
				case 1:
					i = 5
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xe0\xee"...,
					)
				case 10:
					data = /*line SSMkrbGtha.go:1*/ append(data, 191)
					i = 14
				case 5:
					data = /*line SSMkrbGtha.go:1*/ append(data, 188)
					i = 9
				case 2:
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xad\xbb\xb9\xbd"...,
					)
					i = 7
				case 12:
					i = 1
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xb8\xe9"...,
					)
				case 0:
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xb7\xf7\xec\xea"...,
					)
					i = 12
				case 13:
					i = 11
					for y := range data {
						data[y] = data[y] ^ /*line SSMkrbGtha.go:1*/ byte(decryptKey^y)
					}
				case 8:
					i = 3
					data = /*line SSMkrbGtha.go:1*/ append(data, 229)
				case 9:
					i = 13
					data = /*line SSMkrbGtha.go:1*/ append(data, 236)
				case 4:
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xf8\x88"...,
					)
					i = 8
				case 14:
					i = 0
					data = /*line SSMkrbGtha.go:1*/ append(data, "\xbf\xb7"...,
					)
				}
			}
			return /*line SSMkrbGtha.go:1*/ string(data)
		}(), gnGMaXTuiFeE, gnGMaXTuiFeE, lhQWH7m))
	}
}

func x_TVg_Nwq3L(lBuz0jb []byte, aQ1jonyD ...byte) []byte {
	iMFHMG := /*line DVp4Tm_Ch.go:1*/ make([]byte /*line xs6Y_rwW1_t.go:1*/, len(lBuz0jb)+ /*line AY7cHr2.go:1*/ len(aQ1jonyD))
	/*line mbZXb8o.go:1*/ copy(iMFHMG, lBuz0jb)
	/*line FoBTqQsQR.go:1*/ copy(iMFHMG[ /*line N3IR4b.go:1*/ len(lBuz0jb):], aQ1jonyD)
	return iMFHMG
}

func (hkI2UXG_QBd *Trie) v2ULJu(gnGMaXTuiFeE node, tKBrjZ []byte) (node, error) {
	if gnGMaXTuiFeE, yY_yPSd8vG := gnGMaXTuiFeE.(hashNode); yY_yPSd8vG {
		return /*line jd1YVqwTDX0n.go:1*/ hkI2UXG_QBd.xIqJE3(gnGMaXTuiFeE, tKBrjZ)
	}
	return gnGMaXTuiFeE, nil
}

func (hkI2UXG_QBd *Trie) xIqJE3(gnGMaXTuiFeE hashNode, tKBrjZ []byte) (node, error) {
	aBHduJG, eZzE0KPS := /*line DbEigIte9a.go:1*/ hkI2UXG_QBd.reader.uAjz8NxU_cL(tKBrjZ /*line sQ3pVy.go:1*/, common.BytesToHash(gnGMaXTuiFeE))
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	/*line aDIOr17.go:1*/ hkI2UXG_QBd.tracer.d5Xj5n(tKBrjZ, aBHduJG)
	return /*line qUeYe6a8.go:1*/ nyHCCZEA(gnGMaXTuiFeE, aBHduJG), nil
}

func (hkI2UXG_QBd *Trie) Hash() common.Hash {
	rNuN0sMPDJ, kdSzj0 := /*line G_LQlz.go:1*/ hkI2UXG_QBd.gLYHF6c__H()
	hkI2UXG_QBd.root = kdSzj0
	return /*line ZC028etI.go:1*/ common.BytesToHash(rNuN0sMPDJ.(hashNode))
}

func (hkI2UXG_QBd *Trie) Commit(xfPJhFZM25R bool) (common.Hash, *trienode.NodeSet, error) {
	defer /*line ZB5Dq3kuHCT.go:1*/ hkI2UXG_QBd.tracer.x9PQbEXy()
	defer func() {
		/*line J8TaqdZGBMm.go:1*/ hkI2UXG_QBd.committed = true
	}()

	if hkI2UXG_QBd.root == nil {
		k2TrdNkb := /*line ZWf6oZai4aQ.go:1*/ hkI2UXG_QBd.tracer.tnNWw3vS()
		if /*line D64ouHia.go:1*/ len(k2TrdNkb) == 0 {
			return types.NrGuaNA21, nil, nil
		}
		y5wkTqRU := /*line XRAnOT7.go:1*/ trienode.NewNodeSet(hkI2UXG_QBd.owner)
		for _, qiRG6lpeaFl := range k2TrdNkb {
			/*line KuxcF2QDEWe.go:1*/ y5wkTqRU.AddNode([] /*line TxFZfcnOuW.go:1*/ byte(qiRG6lpeaFl) /*line DS1196Ef0kY.go:1*/, trienode.NewDeleted())
		}
		return types.NrGuaNA21, y5wkTqRU, nil
	}

	z8PGQUy := /*line JLKFk4EpoRN.go:1*/ hkI2UXG_QBd.Hash()

	if klVnl8zZd, c0YPpA0vA5f := /*line EUw7N83yDISH.go:1*/ hkI2UXG_QBd.root.tGJzZYYLvK(); !c0YPpA0vA5f {

		hkI2UXG_QBd.root = klVnl8zZd
		return z8PGQUy, nil, nil
	}
	y5wkTqRU := /*line qNiEjtUIJT.go:1*/ trienode.NewNodeSet(hkI2UXG_QBd.owner)
	for _, qiRG6lpeaFl := range /*line CNpcCW23.go:1*/ hkI2UXG_QBd.tracer.tnNWw3vS() {
		/*line IvA05OCeWUV.go:1*/ y5wkTqRU.AddNode([] /*line JhfHzXa4Y1WD.go:1*/ byte(qiRG6lpeaFl) /*line QjZjxm.go:1*/, trienode.NewDeleted())
	}
	hkI2UXG_QBd.root = /*line wFbzjUPmOQ8.go:1*/ egS1cdL6TQG(y5wkTqRU, hkI2UXG_QBd.tracer, xfPJhFZM25R).Commit(hkI2UXG_QBd.root)
	return z8PGQUy, y5wkTqRU, nil
}

func (hkI2UXG_QBd *Trie) gLYHF6c__H() (node, node) {
	if hkI2UXG_QBd.root == nil {
		return /*line I6wHw8uA.go:1*/ hashNode( /*line f4aroXtS7Gs.go:1*/ types.NrGuaNA21.Bytes()), nil
	}

	klE3zy := /*line NaNw1VQrDaL.go:1*/ dCZJziX(hkI2UXG_QBd.unhashed >= 100)
	defer func() {
		/*line s52W_D4v6oC.go:1*/ ptIMYRK9(klE3zy)
		hkI2UXG_QBd.unhashed = 0
	}()
	iVcAtTyHaK, kdSzj0 := /*line t3KDocXGT.go:1*/ klE3zy.rNuN0sMPDJ(hkI2UXG_QBd.root, true)
	return iVcAtTyHaK, kdSzj0
}

func (hkI2UXG_QBd *Trie) Reset() {
	hkI2UXG_QBd.root = nil
	hkI2UXG_QBd.owner = common.Hash{}
	hkI2UXG_QBd.unhashed = 0
	/*line YoG0m77C2q.go:1*/ hkI2UXG_QBd.tracer.x9PQbEXy()
	hkI2UXG_QBd.committed = false
}

var _ bytes.Buffer
var _ = errors.As
var _ = fmt.Append
var _ types.AccessList
var _ = common.AbsolutePath
var _ = log.Crit
var _ trienode.MergedNodeSet
var _ database.Database
