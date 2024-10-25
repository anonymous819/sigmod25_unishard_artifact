//line :1
package trie

import (
	"encoding/binary"
	"errors"
	"fmt"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/trie/utils"
	"github.com/ethereum/go-ethereum/triedb/database"
	"github.com/gballet/go-verkle"
	"github.com/holiman/uint256"
)

var (
	uROBXD2cL [32]byte
	cKf61egNB = /*line pWaJqMZAaRJo.go:1*/ errors.New(func() /*line FrP8ujeWqQXL.go:1*/ string {
		data := [] /*line FrP8ujeWqQXL.go:1*/ byte("\x04\n\x02\tv\x84\fN'o1\x11 \x98y1\x1e f\x14\x1a\xfd\x1f\a\t\x95")
		positions := [...]byte{23, 6, 1, 10, 21, 3, 15, 8, 3, 13, 7, 21, 2, 11, 0, 22, 20, 5, 4, 7, 24, 24, 10, 7, 3, 25, 15, 20, 16, 19, 3, 3, 20, 13}
		for i := 0; i < 34; i += 2 {
			localKey := /*line FrP8ujeWqQXL.go:1*/ byte(i) + /*line FrP8ujeWqQXL.go:1*/ byte(positions[i]^positions[i+1]) + 82
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line FrP8ujeWqQXL.go:1*/ string(data)
	}())
)

type VerkleTrie struct {
	root   verkle.VerkleNode
	cache  *utils.PointCache
	reader *trieReader
}

func DDmplr(m7S8SE common.Hash, kRC_1aK database.Database, tGJzZYYLvK *utils.PointCache) (*VerkleTrie, error) {
	dmVOxRYtQ, eZzE0KPS := /*line QdqwvlGNF0.go:1*/ fHP68lzj9vD1(m7S8SE, common.Hash{}, kRC_1aK)
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}

	uAjz8NxU_cL := /*line JqrHdU0RTR.go:1*/ verkle.New()
	if m7S8SE != types.L748VNVQ && m7S8SE != types.NrGuaNA21 {
		aBHduJG, eZzE0KPS := /*line G6l2Mxxnboi.go:1*/ dmVOxRYtQ.uAjz8NxU_cL(nil, common.Hash{})
		if eZzE0KPS != nil {
			return nil, eZzE0KPS
		}
		uAjz8NxU_cL, eZzE0KPS = /*line DFt1GJTZ.go:1*/ verkle.ParseNode(aBHduJG, 0)
		if eZzE0KPS != nil {
			return nil, eZzE0KPS
		}
	}
	return &VerkleTrie{
		root:   uAjz8NxU_cL,
		cache:  tGJzZYYLvK,
		reader: dmVOxRYtQ,
	}, nil
}

func (hkI2UXG_QBd *VerkleTrie) GetKey(lhQWH7m []byte) []byte {
	return lhQWH7m
}

func (hkI2UXG_QBd *VerkleTrie) GetAccount(sJ3vb1eSGg common.Address) (*types.StateAccount, error) {
	var (
		gU1J_yS   = &types.StateAccount{}
		_Et7RxouL [][]byte
		eZzE0KPS  error
	)
	switch gnGMaXTuiFeE := hkI2UXG_QBd.root.(type) {
	case *verkle.InternalNode:
		_Et7RxouL, eZzE0KPS = /*line RDhUHtPzfM.go:1*/ gnGMaXTuiFeE.GetValuesAtStem( /*line ZkfkogQBvXta.go:1*/ hkI2UXG_QBd.cache.GetStem(sJ3vb1eSGg[:]), hkI2UXG_QBd.hQUG60b46)
		if eZzE0KPS != nil {
			return nil /*line eDdIyI.go:1*/, fmt.Errorf(func() /*line JPDtDTOAd8.go:1*/ string {
				fullData := [] /*line JPDtDTOAd8.go:1*/ byte("-\xf8\xb5x\xcc\a\xa4\xfdɰy\xec\xff.[\xde\x06\xcfB\xcd\xe62\x89\x15z\xa9\xae\xe7z\xa3\x94P\xe9\xbf\xd5P\x00u\x8cv\xf5\x0e\xf0\x85Ƀ&X[H")
				idxKey := [] /*line JPDtDTOAd8.go:1*/ byte("\xeeXp\xd3\x0f\x83\xaa\x8aG\xc2m\x8f\xe0f:\r")
				data := /*line JPDtDTOAd8.go:1*/ make([]byte, 0, 26)
				data = /*line JPDtDTOAd8.go:1*/ append(data, fullData[59^ /*line JPDtDTOAd8.go:1*/ int(idxKey[14])]^fullData[27^ /*line JPDtDTOAd8.go:1*/ int(idxKey[14])], fullData[217^ /*line JPDtDTOAd8.go:1*/ int(idxKey[3])]+fullData[216^ /*line JPDtDTOAd8.go:1*/ int(idxKey[3])], fullData[174^ /*line JPDtDTOAd8.go:1*/ int(idxKey[7])]-fullData[172^ /*line JPDtDTOAd8.go:1*/ int(idxKey[7])], fullData[193^ /*line JPDtDTOAd8.go:1*/ int(idxKey[3])]+fullData[223^ /*line JPDtDTOAd8.go:1*/ int(idxKey[3])], fullData[226^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])]+fullData[250^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])], fullData[42^ /*line JPDtDTOAd8.go:1*/ int(idxKey[15])]^fullData[26^ /*line JPDtDTOAd8.go:1*/ int(idxKey[15])], fullData[78^ /*line JPDtDTOAd8.go:1*/ int(idxKey[1])]+fullData[76^ /*line JPDtDTOAd8.go:1*/ int(idxKey[1])], fullData[238^ /*line JPDtDTOAd8.go:1*/ int(idxKey[0])]^fullData[193^ /*line JPDtDTOAd8.go:1*/ int(idxKey[0])], fullData[206^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])]^fullData[209^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])], fullData[255^ /*line JPDtDTOAd8.go:1*/ int(idxKey[0])]-fullData[222^ /*line JPDtDTOAd8.go:1*/ int(idxKey[0])], fullData[146^ /*line JPDtDTOAd8.go:1*/ int(idxKey[11])]^fullData[162^ /*line JPDtDTOAd8.go:1*/ int(idxKey[11])], fullData[110^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])]+fullData[100^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])], fullData[224^ /*line JPDtDTOAd8.go:1*/ int(idxKey[9])]+fullData[225^ /*line JPDtDTOAd8.go:1*/ int(idxKey[9])], fullData[106^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])]-fullData[70^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])], fullData[99^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])]-fullData[120^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])], fullData[74^ /*line JPDtDTOAd8.go:1*/ int(idxKey[8])]^fullData[110^ /*line JPDtDTOAd8.go:1*/ int(idxKey[8])], fullData[219^ /*line JPDtDTOAd8.go:1*/ int(idxKey[9])]^fullData[198^ /*line JPDtDTOAd8.go:1*/ int(idxKey[9])], fullData[125^ /*line JPDtDTOAd8.go:1*/ int(idxKey[13])]-fullData[67^ /*line JPDtDTOAd8.go:1*/ int(idxKey[13])], fullData[29^ /*line JPDtDTOAd8.go:1*/ int(idxKey[15])]-fullData[19^ /*line JPDtDTOAd8.go:1*/ int(idxKey[15])], fullData[248^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])]+fullData[200^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])], fullData[114^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])]-fullData[98^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])], fullData[229^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])]-fullData[243^ /*line JPDtDTOAd8.go:1*/ int(idxKey[12])], fullData[47^ /*line JPDtDTOAd8.go:1*/ int(idxKey[4])]^fullData[35^ /*line JPDtDTOAd8.go:1*/ int(idxKey[4])], fullData[130^ /*line JPDtDTOAd8.go:1*/ int(idxKey[7])]-fullData[140^ /*line JPDtDTOAd8.go:1*/ int(idxKey[7])], fullData[71^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])]-fullData[113^ /*line JPDtDTOAd8.go:1*/ int(idxKey[10])])
				return /*line JPDtDTOAd8.go:1*/ string(data)
			}(), sJ3vb1eSGg, eZzE0KPS)
		}
	default:
		return nil, cKf61egNB
	}
	if _Et7RxouL == nil {
		return nil, nil
	}

	if /*line NyJT9muHWjB.go:1*/ len(_Et7RxouL[utils.NonceLeafKey]) > 0 {
		gU1J_yS.Nonce = /*line OjjhHI.go:1*/ binary.LittleEndian.Uint64(_Et7RxouL[utils.NonceLeafKey])
	}

	var ixqfJZh [32]byte
	/*line Et06dYgNPu6G.go:1*/ copy(ixqfJZh[:], _Et7RxouL[utils.BalanceLeafKey])
	for q2u2020KD6 := 0; q2u2020KD6 < /*line l5Mldqt.go:1*/ len(ixqfJZh)/2; q2u2020KD6++ {
		ixqfJZh[ /*line q5qU9KJ817A.go:1*/ len(ixqfJZh)-q2u2020KD6-1], ixqfJZh[q2u2020KD6] = ixqfJZh[q2u2020KD6], ixqfJZh[ /*line OLdJQgZfCSW.go:1*/ len(ixqfJZh)-q2u2020KD6-1]
	}
	gU1J_yS.Balance = /*line CcpeuO.go:1*/ new(uint256.Int).SetBytes32(ixqfJZh[:])

	gU1J_yS.CodeHash = _Et7RxouL[utils.CodeKeccakLeafKey]

	return gU1J_yS, nil
}

func (hkI2UXG_QBd *VerkleTrie) GetStorage(sJ3vb1eSGg common.Address, lhQWH7m []byte) ([]byte, error) {
	ec_diEU := /*line GeQVIHqzX0.go:1*/ utils.StorageSlotKeyWithEvaluatedAddress( /*line HdoHHlaJ.go:1*/ hkI2UXG_QBd.cache.Get( /*line VOPTZ8MtQG.go:1*/ sJ3vb1eSGg.Bytes()), lhQWH7m)
	h_pxYek4zT, eZzE0KPS := /*line jMhVHcpl28.go:1*/ hkI2UXG_QBd.root.Get(ec_diEU, hkI2UXG_QBd.hQUG60b46)
	if eZzE0KPS != nil {
		return nil, eZzE0KPS
	}
	return /*line FgjxNECY.go:1*/ common.TrimLeftZeroes(h_pxYek4zT), nil
}

func (hkI2UXG_QBd *VerkleTrie) UpdateAccount(sJ3vb1eSGg common.Address, gU1J_yS *types.StateAccount) error {
	var (
		eZzE0KPS          error
		e0ZQ53Ob, ixqfJZh [32]byte
		_Et7RxouL         = /*line gkMbECr9LA.go:1*/ make([][]byte, verkle.NodeWidth)
	)
	_Et7RxouL[utils.VersionLeafKey] = uROBXD2cL[:]
	_Et7RxouL[utils.CodeKeccakLeafKey] = gU1J_yS.CodeHash[:]

	/*line syCUTO1S.go:1*/
	binary.LittleEndian.PutUint64(e0ZQ53Ob[:], gU1J_yS.Nonce)
	_Et7RxouL[utils.NonceLeafKey] = e0ZQ53Ob[:]

	sNrBzyUXd6av := /*line XzNpEkYi.go:1*/ gU1J_yS.Balance.Bytes()
	if /*line GAtkXATa4eUZ.go:1*/ len(sNrBzyUXd6av) > 0 {
		for q2u2020KD6, mZsNbw0 := range sNrBzyUXd6av {
			ixqfJZh[ /*line JhgmbsEYp.go:1*/ len(sNrBzyUXd6av)-q2u2020KD6-1] = mZsNbw0
		}
	}
	_Et7RxouL[utils.BalanceLeafKey] = ixqfJZh[:]

	switch gnGMaXTuiFeE := hkI2UXG_QBd.root.(type) {
	case *verkle.InternalNode:
		eZzE0KPS = /*line s2pcGs_jVwkr.go:1*/ gnGMaXTuiFeE.InsertValuesAtStem( /*line VHj0Vl.go:1*/ hkI2UXG_QBd.cache.GetStem(sJ3vb1eSGg[:]), _Et7RxouL, hkI2UXG_QBd.hQUG60b46)
		if eZzE0KPS != nil {
			return /*line IPQxGhJRweZU.go:1*/ fmt.Errorf(func() /*line NATMuRH.go:1*/ string {
				data := /*line NATMuRH.go:1*/ make([]byte, 0, 29)
				i := 14
				decryptKey := 37
				for counter := 0; i != 7; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 6:
						data = /*line NATMuRH.go:1*/ append(data, "\x9d\x97"...,
						)
						i = 5
					case 5:
						data = /*line NATMuRH.go:1*/ append(data, "\xdb\xe3"...,
						)
						i = 8
					case 3:
						i = 10
						data = /*line NATMuRH.go:1*/ append(data, 220)
					case 13:
						data = /*line NATMuRH.go:1*/ append(data, "\x8f\xed"...,
						)
						i = 6
					case 2:
						data = /*line NATMuRH.go:1*/ append(data, "Ť\xc5\xd0"...,
						)
						i = 12
					case 8:
						i = 11
						data = /*line NATMuRH.go:1*/ append(data, 226)
					case 1:
						i = 7
						for y := range data {
							data[y] = data[y] - /*line NATMuRH.go:1*/ byte(decryptKey^y)
						}
					case 9:
						data = /*line NATMuRH.go:1*/ append(data, "\xc7\xd5"...,
						)
						i = 2
					case 10:
						i = 13
						data = /*line NATMuRH.go:1*/ append(data, "݈\x93"...,
						)
					case 12:
						i = 3
						data = /*line NATMuRH.go:1*/ append(data, "\xdb\xe4"...,
						)
					case 4:
						i = 1
						data = /*line NATMuRH.go:1*/ append(data, 244)
					case 14:
						data = /*line NATMuRH.go:1*/ append(data, "\xba\xd4\xcb"...,
						)
						i = 9
					case 0:
						data = /*line NATMuRH.go:1*/ append(data, "䷜\xa4"...,
						)
						i = 4
					case 11:
						i = 0
						data = /*line NATMuRH.go:1*/ append(data, 226)
					}
				}
				return /*line NATMuRH.go:1*/ string(data)
			}(), sJ3vb1eSGg, eZzE0KPS)
		}
	default:
		return cKf61egNB
	}

	return nil
}

func (hkI2UXG_QBd *VerkleTrie) UpdateStorage(l98Od6T1zF common.Address, lhQWH7m, ar76Sw []byte) error {

	var eBWWX0caQi [32]byte
	if /*line _PZBdMVB.go:1*/ len(ar76Sw) >= 32 {
		/*line mpvb4kv.go:1*/ copy(eBWWX0caQi[:], ar76Sw[:32])
	} else {
		/*line u0QcLcwqdj.go:1*/ copy(eBWWX0caQi[32- /*line zgQkaxHdN.go:1*/ len(ar76Sw):], ar76Sw[:])
	}
	ec_diEU := /*line BzM8lQa.go:1*/ utils.StorageSlotKeyWithEvaluatedAddress( /*line rXLUwzqKrurh.go:1*/ hkI2UXG_QBd.cache.Get( /*line cVCJoGSY8bWm.go:1*/ l98Od6T1zF.Bytes()), lhQWH7m)
	return /*line PzQAFjTRbP5.go:1*/ hkI2UXG_QBd.root.Insert(ec_diEU, eBWWX0caQi[:], hkI2UXG_QBd.hQUG60b46)
}

func (hkI2UXG_QBd *VerkleTrie) DeleteAccount(sJ3vb1eSGg common.Address) error {
	var (
		eZzE0KPS  error
		_Et7RxouL = /*line MqHxut.go:1*/ make([][]byte, verkle.NodeWidth)
	)
	for q2u2020KD6 := 0; q2u2020KD6 < verkle.NodeWidth; q2u2020KD6++ {
		_Et7RxouL[q2u2020KD6] = uROBXD2cL[:]
	}
	switch gnGMaXTuiFeE := hkI2UXG_QBd.root.(type) {
	case *verkle.InternalNode:
		eZzE0KPS = /*line K30wag.go:1*/ gnGMaXTuiFeE.InsertValuesAtStem( /*line GBq4RT.go:1*/ hkI2UXG_QBd.cache.GetStem( /*line I6gVJPbSl.go:1*/ sJ3vb1eSGg.Bytes()), _Et7RxouL, hkI2UXG_QBd.hQUG60b46)
		if eZzE0KPS != nil {
			return /*line kLt1zLM7eN.go:1*/ fmt.Errorf(func() /*line HALR1Ty9ssI.go:1*/ string {
				data := /*line HALR1Ty9ssI.go:1*/ make([]byte, 0, 29)
				i := 2
				decryptKey := 84
				for counter := 0; i != 3; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 0:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, 135)
						i = 5
					case 6:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "zlEh"...,
						)
						i = 1
					case 7:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "v\x88\x89\x83"...,
						)
						i = 0
					case 5:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "T;="...,
						)
						i = 4
					case 2:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "Fhlf"...,
						)
						i = 6
					case 10:
						i = 9
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "w\x82/4"...,
						)
					case 9:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "2\x8a<0"...,
						)
						i = 7
					case 1:
						data = /*line HALR1Ty9ssI.go:1*/ append(data, "mz}"...,
						)
						i = 10
					case 8:
						for y := range data {
							data[y] = data[y] - /*line HALR1Ty9ssI.go:1*/ byte(decryptKey^y)
						}
						i = 3
					case 4:
						i = 8
						data = /*line HALR1Ty9ssI.go:1*/ append(data, 143)
					}
				}
				return /*line HALR1Ty9ssI.go:1*/ string(data)
			}(), sJ3vb1eSGg, eZzE0KPS)
		}
	default:
		return cKf61egNB
	}
	return nil
}

func (hkI2UXG_QBd *VerkleTrie) DeleteStorage(sJ3vb1eSGg common.Address, lhQWH7m []byte) error {
	var uROBXD2cL [32]byte
	ec_diEU := /*line RX8nyYhEI.go:1*/ utils.StorageSlotKeyWithEvaluatedAddress( /*line QHwFJjM.go:1*/ hkI2UXG_QBd.cache.Get( /*line dRLoAK.go:1*/ sJ3vb1eSGg.Bytes()), lhQWH7m)
	return /*line TMRFlsAE4O.go:1*/ hkI2UXG_QBd.root.Insert(ec_diEU, uROBXD2cL[:], hkI2UXG_QBd.hQUG60b46)
}

func (hkI2UXG_QBd *VerkleTrie) Hash() common.Hash {
	return /*line JLxzEmhmHP.go:1*/ hkI2UXG_QBd.root.Commit().Bytes()
}

func (hkI2UXG_QBd *VerkleTrie) Commit(_ bool) (common.Hash, *trienode.NodeSet, error) {
	m7S8SE, yY_yPSd8vG := hkI2UXG_QBd.root.(*verkle.InternalNode)
	if !yY_yPSd8vG {
		return common.Hash{}, nil /*line xICTtNsS.go:1*/, errors.New(func() /*line pF63N10.go:1*/ string {
			data := [] /*line pF63N10.go:1*/ byte("ua$\x0f4eKtadU`NV-Qnon\xf3\xdety\xf4e")
			positions := [...]byte{20, 14, 20, 23, 3, 2, 13, 19, 10, 4, 13, 10, 3, 19, 2, 14, 15, 2, 6, 12, 11, 18, 12, 1, 14, 8}
			for i := 0; i < 26; i += 2 {
				localKey := /*line pF63N10.go:1*/ byte(i) + /*line pF63N10.go:1*/ byte(positions[i]^positions[i+1]) + 207
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line pF63N10.go:1*/ string(data)
		}())
	}
	y5wkTqRU, eZzE0KPS := /*line aUAAPwb.go:1*/ m7S8SE.BatchSerialize()
	if eZzE0KPS != nil {
		return common.Hash{}, nil /*line F02VolRm4.go:1*/, fmt.Errorf(func() /*line CZ8aHI.go:1*/ string {
			data := [] /*line CZ8aHI.go:1*/ byte("sH6\xbb-liID9g t\xf69e nod\xcd\x0e\xe51\xfb\xa6")
			positions := [...]byte{23, 3, 24, 23, 1, 8, 13, 20, 7, 3, 25, 7, 21, 24, 13, 24, 7, 24, 21, 2, 4, 20, 25, 9, 9, 22, 9, 14}
			for i := 0; i < 28; i += 2 {
				localKey := /*line CZ8aHI.go:1*/ byte(i) + /*line CZ8aHI.go:1*/ byte(positions[i]^positions[i+1]) + 20
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line CZ8aHI.go:1*/ string(data)
		}(), eZzE0KPS)
	}
	mff5TAZylrk := /*line BMeOh9E9R.go:1*/ trienode.NewNodeSet(common.Hash{})
	for _, uAjz8NxU_cL := range y5wkTqRU {

		/*line lznqHVLUkDc.go:1*/
		mff5TAZylrk.AddNode(uAjz8NxU_cL.Path /*line HSME8xp.go:1*/, trienode.New(common.Hash{}, uAjz8NxU_cL.SerializedBytes))
	}

	return /*line liLxJa.go:1*/ hkI2UXG_QBd.Hash(), mff5TAZylrk, nil
}

func (hkI2UXG_QBd *VerkleTrie) NodeIterator(hphffFll9Zv []byte) (FBy5J0gbkg, error) {
	/*line DFl47mk9Qr.go:1*/ panic(func() /*line KlWODiqg.go:1*/ string {
		key := [] /*line KlWODiqg.go:1*/ byte("{\xe0SKd\xe7}\b\xaa\xfaANA0\xad")
		data := [] /*line KlWODiqg.go:1*/ byte("\xf3\x8f!\xd5\x05\x86\xf3d\xbbs$ 35\xb7")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line KlWODiqg.go:1*/ string(data)
	}())
}

func (hkI2UXG_QBd *VerkleTrie) Prove(lhQWH7m []byte, zzPkun7 ethdb.KeyValueWriter) error {
	/*line at1RzJCoj.go:1*/ panic(func() /*line ij1Vanmo8_xi.go:1*/ string {
		key := [] /*line ij1Vanmo8_xi.go:1*/ byte("\x1dܻE{\xe0\x9e\x84\"\x9c\x91Wdi\x9d")
		data := [] /*line ij1Vanmo8_xi.go:1*/ byte("s\xb3\xcfe\x12\x8d\xee\xe8G\xf1\xf49\x10\f\xf9")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line ij1Vanmo8_xi.go:1*/ string(data)
	}())
}

func (hkI2UXG_QBd *VerkleTrie) Copy() *VerkleTrie {
	return &VerkleTrie{
		root:/*line cnAaIrvP.go:1*/ hkI2UXG_QBd.root.Copy(),
		cache:  hkI2UXG_QBd.cache,
		reader: hkI2UXG_QBd.reader,
	}
}

func (hkI2UXG_QBd *VerkleTrie) IsVerkle() bool {
	return true
}

type SXoLHxUt177q []byte

const (
	PUSH1  = /*line N6YCIPgX.go:1*/ byte(0x60)
	PUSH32 = /*line AsFHAw_F.go:1*/ byte(0x7f)
)

func TPNBHlItln7n(abS5SZPC []byte) SXoLHxUt177q {
	var (
		u_GBUE       = 0
		f9LPLtKA1Yg0 = /*line Ngv7S3nCe.go:1*/ len(abS5SZPC) / 31
		xp5a1CLm1    = 0
	)
	if /*line zk4elh4.go:1*/ len(abS5SZPC)%31 != 0 {
		f9LPLtKA1Yg0++
	}
	eOPsnAqnato := /*line aSa11GjA.go:1*/ make([]byte, f9LPLtKA1Yg0*32)
	for q2u2020KD6 := 0; q2u2020KD6 < f9LPLtKA1Yg0; q2u2020KD6++ {

		cP6eH_ := 31 * (q2u2020KD6 + 1)
		if /*line JFbpi0.go:1*/ len(abS5SZPC) < cP6eH_ {
			cP6eH_ = /*line ToPP11LpR0G.go:1*/ len(abS5SZPC)
		}
		/*line qMToKV5g.go:1*/ copy(eOPsnAqnato[q2u2020KD6*32+1:], abS5SZPC[31*q2u2020KD6:cP6eH_])

		if u_GBUE > 31 {

			eOPsnAqnato[q2u2020KD6*32] = 31
			u_GBUE = 1
			continue
		}
		eOPsnAqnato[32*q2u2020KD6] = /*line tU1pSPj.go:1*/ byte(u_GBUE)
		u_GBUE = 0

		for ; xp5a1CLm1 < cP6eH_; xp5a1CLm1++ {
			if abS5SZPC[xp5a1CLm1] >= PUSH1 && abS5SZPC[xp5a1CLm1] <= PUSH32 {
				xp5a1CLm1 += /*line RyvZqCTXg.go:1*/ int(abS5SZPC[xp5a1CLm1] - PUSH1 + 1)
				if xp5a1CLm1+1 >= 31*(q2u2020KD6+1) {
					xp5a1CLm1++
					u_GBUE = xp5a1CLm1 - 31*(q2u2020KD6+1)
					break
				}
			}
		}
	}
	return eOPsnAqnato
}

func (hkI2UXG_QBd *VerkleTrie) UpdateContractCode(sJ3vb1eSGg common.Address, axeVrX common.Hash, abS5SZPC []byte) error {
	var (
		eOPsnAqnato = /*line xTdfV9J.go:1*/ TPNBHlItln7n(abS5SZPC)
		_Et7RxouL   [][]byte
		lhQWH7m     []byte
		eZzE0KPS    error
	)
	for q2u2020KD6, tqlGm3JVO4rM := 0 /*line W9we4NX9d0X.go:1*/, uint64(0); q2u2020KD6 < /*line IT0aNaT.go:1*/ len(eOPsnAqnato); q2u2020KD6, tqlGm3JVO4rM = q2u2020KD6+32, tqlGm3JVO4rM+1 {
		sLGV_KjaCeT := (tqlGm3JVO4rM + 128) % 256
		if sLGV_KjaCeT == 0 || tqlGm3JVO4rM == 0 {
			_Et7RxouL = /*line nf8_8GQpI3.go:1*/ make([][]byte, verkle.NodeWidth)
			lhQWH7m = /*line ADLis_j.go:1*/ utils.CodeChunkKeyWithEvaluatedAddress( /*line BPfaPCZ.go:1*/ hkI2UXG_QBd.cache.Get( /*line tV1VQE6Dk7.go:1*/ sJ3vb1eSGg.Bytes()) /*line KK81FDvCYPK.go:1*/, uint256.NewInt(tqlGm3JVO4rM))
		}
		_Et7RxouL[sLGV_KjaCeT] = eOPsnAqnato[q2u2020KD6 : q2u2020KD6+32]

		if q2u2020KD6 == 0 {
			hiRBsibTQK := /*line Za4U0ITK.go:1*/ make([]byte, 32)
			/*line E9XJuRUsi9V.go:1*/ binary.LittleEndian.PutUint64(hiRBsibTQK /*line g9_Nmdh1Yf0.go:1*/, uint64( /*line w8yQNcc8qqw.go:1*/ len(abS5SZPC)))
			_Et7RxouL[utils.CodeSizeLeafKey] = hiRBsibTQK
		}
		if sLGV_KjaCeT == 255 || /*line tNAdWyIXbpc.go:1*/ len(eOPsnAqnato)-q2u2020KD6 <= 32 {
			switch m7S8SE := hkI2UXG_QBd.root.(type) {
			case *verkle.InternalNode:
				eZzE0KPS = /*line EJhEdOo_S.go:1*/ m7S8SE.InsertValuesAtStem(lhQWH7m[:31], _Et7RxouL, hkI2UXG_QBd.hQUG60b46)
				if eZzE0KPS != nil {
					return /*line fyGfqA.go:1*/ fmt.Errorf(func() /*line ZzH8APRu80r.go:1*/ string {
						fullData := [] /*line ZzH8APRu80r.go:1*/ byte("\xba\xad\xb9\t\xaf\"LB\xe6\x89\xe7\x02j\x91U\xe0D\xad\xdfVJ\x0e\xea>\xe1P\"/\x98\xaaPۛ\xf8Ŷ\xb0\xf3j*Lm\x90\x9a\xbe\xaab\xd1\x18N0Y\x16>i=JVj\xc1\xc8\x7f\"ӟWf\x97\xffEc\x85\xf0\xc0\xae\x0f")
						idxKey := [] /*line ZzH8APRu80r.go:1*/ byte("\x00a\fIH\x8eRh\xabFEOl\xe9\xeb\xc3")
						data := /*line ZzH8APRu80r.go:1*/ make([]byte, 0, 39)
						data = /*line ZzH8APRu80r.go:1*/ append(data, fullData[202^ /*line ZzH8APRu80r.go:1*/ int(idxKey[13])]+fullData[169^ /*line ZzH8APRu80r.go:1*/ int(idxKey[13])], fullData[124^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])]^fullData[10^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])], fullData[101^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])]+fullData[72^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])], fullData[106^ /*line ZzH8APRu80r.go:1*/ int(idxKey[10])]+fullData[111^ /*line ZzH8APRu80r.go:1*/ int(idxKey[10])], fullData[213^ /*line ZzH8APRu80r.go:1*/ int(idxKey[14])]-fullData[161^ /*line ZzH8APRu80r.go:1*/ int(idxKey[14])], fullData[113^ /*line ZzH8APRu80r.go:1*/ int(idxKey[3])]^fullData[82^ /*line ZzH8APRu80r.go:1*/ int(idxKey[3])], fullData[99^ /*line ZzH8APRu80r.go:1*/ int(idxKey[9])]+fullData[88^ /*line ZzH8APRu80r.go:1*/ int(idxKey[9])], fullData[19^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])]+fullData[98^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])], fullData[88^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])]^fullData[86^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])], fullData[162^ /*line ZzH8APRu80r.go:1*/ int(idxKey[14])]-fullData[195^ /*line ZzH8APRu80r.go:1*/ int(idxKey[14])], fullData[185^ /*line ZzH8APRu80r.go:1*/ int(idxKey[8])]^fullData[186^ /*line ZzH8APRu80r.go:1*/ int(idxKey[8])], fullData[123^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])]-fullData[90^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])], fullData[116^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])]+fullData[111^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])], fullData[109^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])]+fullData[75^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])], fullData[234^ /*line ZzH8APRu80r.go:1*/ int(idxKey[14])]-fullData[209^ /*line ZzH8APRu80r.go:1*/ int(idxKey[14])], fullData[90^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])]+fullData[91^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])], fullData[70^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])]^fullData[80^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])], fullData[4^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])]^fullData[67^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])], fullData[15^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])]+fullData[104^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])], fullData[2^ /*line ZzH8APRu80r.go:1*/ int(idxKey[0])]-fullData[13^ /*line ZzH8APRu80r.go:1*/ int(idxKey[0])], fullData[152^ /*line ZzH8APRu80r.go:1*/ int(idxKey[8])]-fullData[138^ /*line ZzH8APRu80r.go:1*/ int(idxKey[8])], fullData[143^ /*line ZzH8APRu80r.go:1*/ int(idxKey[8])]-fullData[173^ /*line ZzH8APRu80r.go:1*/ int(idxKey[8])], fullData[41^ /*line ZzH8APRu80r.go:1*/ int(idxKey[0])]-fullData[3^ /*line ZzH8APRu80r.go:1*/ int(idxKey[0])], fullData[119^ /*line ZzH8APRu80r.go:1*/ int(idxKey[10])]-fullData[105^ /*line ZzH8APRu80r.go:1*/ int(idxKey[10])], fullData[107^ /*line ZzH8APRu80r.go:1*/ int(idxKey[12])]^fullData[81^ /*line ZzH8APRu80r.go:1*/ int(idxKey[12])], fullData[90^ /*line ZzH8APRu80r.go:1*/ int(idxKey[12])]-fullData[124^ /*line ZzH8APRu80r.go:1*/ int(idxKey[12])], fullData[95^ /*line ZzH8APRu80r.go:1*/ int(idxKey[7])]^fullData[45^ /*line ZzH8APRu80r.go:1*/ int(idxKey[7])], fullData[92^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])]+fullData[112^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])], fullData[120^ /*line ZzH8APRu80r.go:1*/ int(idxKey[12])]^fullData[74^ /*line ZzH8APRu80r.go:1*/ int(idxKey[12])], fullData[109^ /*line ZzH8APRu80r.go:1*/ int(idxKey[9])]^fullData[2^ /*line ZzH8APRu80r.go:1*/ int(idxKey[9])], fullData[97^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])]-fullData[7^ /*line ZzH8APRu80r.go:1*/ int(idxKey[11])], fullData[122^ /*line ZzH8APRu80r.go:1*/ int(idxKey[9])]+fullData[91^ /*line ZzH8APRu80r.go:1*/ int(idxKey[9])], fullData[183^ /*line ZzH8APRu80r.go:1*/ int(idxKey[5])]-fullData[132^ /*line ZzH8APRu80r.go:1*/ int(idxKey[5])], fullData[94^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])]^fullData[84^ /*line ZzH8APRu80r.go:1*/ int(idxKey[4])], fullData[126^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])]^fullData[121^ /*line ZzH8APRu80r.go:1*/ int(idxKey[1])], fullData[87^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])]-fullData[89^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])], fullData[3^ /*line ZzH8APRu80r.go:1*/ int(idxKey[10])]-fullData[112^ /*line ZzH8APRu80r.go:1*/ int(idxKey[10])], fullData[93^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])]+fullData[17^ /*line ZzH8APRu80r.go:1*/ int(idxKey[6])])
						return /*line ZzH8APRu80r.go:1*/ string(data)
					}(), sJ3vb1eSGg[:], eZzE0KPS)
				}
			default:
				return cKf61egNB
			}
		}
	}
	return nil
}

func (hkI2UXG_QBd *VerkleTrie) ToDot() string {
	return /*line pgiawkq7.go:1*/ verkle.ToDot(hkI2UXG_QBd.root)
}

func (hkI2UXG_QBd *VerkleTrie) hQUG60b46(qiRG6lpeaFl []byte) ([]byte, error) {
	return /*line q45HQuJwe.go:1*/ hkI2UXG_QBd.reader.uAjz8NxU_cL(qiRG6lpeaFl, common.Hash{})
}

var _ = binary.Append
var _ = errors.As
var _ = fmt.Append
var _ types.AccessList
var _ = common.AbsolutePath
var _ ethdb.AncientReader
var _ trienode.MergedNodeSet
var _ = utils.BalanceKey
var _ database.Database
var _ = verkle.BatchNewLeafNode
var _ = uint256.ErrBadBufferLength
