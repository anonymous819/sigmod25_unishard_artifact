//line :1
package blockchain

import (
	"fmt"
	"os"
	"strconv"

	config "unishard/config"
	crypto "unishard/crypto"
	evm "unishard/evm"
	state "unishard/evm/state"
	log "unishard/log"
	message "unishard/message"
	quorum "unishard/quorum"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

type Y3t7C8s0m struct {
	gamo3nN *VE_4Hib
	eyUr46t *state.StateDB
	dN52FYW *state.StateDB
	i9CVg80 state.Database

	uraHsZ_PG map[common.Address]map[common.Hash]*SnapshotControlTable

	zZzr7abV_N  int
	njakPUdNatA int
	snDFqw      int
	doByMIZ     int
}

type SnapshotControlTable struct {
	NQB0qlzATMvq string
	BlC2JXSmAuvQ []common.Hash
	UPa0r5o      bool
}

func O8OnUJH0nc2(fddpyk_YAhM types.Shard, mubE4Ycelim int, krk3qIRD3K3 ethdb.Database) *Y3t7C8s0m {
	j4ZllBpzGP9 := /*line ic91WHuHC0.go:1*/ new(Y3t7C8s0m)
	j4ZllBpzGP9.gamo3nN = /*line mdK5hVv.go:1*/ Ok0v5OJZ3H()
	j4ZllBpzGP9.uraHsZ_PG = /*line sVEOGnK.go:1*/ make(map[common.Address]map[common.Hash]*SnapshotControlTable)

	_SVuw80vd_pA := []string{}

	for vYVJPf9SlbQs := 1; vYVJPf9SlbQs <= /*line ZKprPcp32wC.go:1*/ config.GetConfig().ShardCount; vYVJPf9SlbQs++ {
		knIJqGXv5X, _ := /*line MxbBmFYKm.go:1*/ os.ReadFile( /*line SWyiCg6.go:1*/ fmt.Sprintf(func() /*line GbpP72w6.go:1*/ string {
			key := [] /*line GbpP72w6.go:1*/ byte("\xc0\xd5\x0eۘ\x14\xb0\x9c\xd4r\xbd\xa8\x8e\x8a\xc3/ (")
			data := [] /*line GbpP72w6.go:1*/ byte("\xee\xfb!\xb8\xf7y\xdd\xf3\xba]\xce\xdc\xef\xfe\xa6KB\a")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line GbpP72w6.go:1*/ string(data)
		}()+ /*line QKFg3O.go:1*/ strconv.Itoa( /*line nae5nb.go:1*/ config.GetConfig().ShardCount)+func() /*line CRittkj.go:1*/ string {
			fullData := [] /*line CRittkj.go:1*/ byte("\xd0)\xd0\xc2\xd7D\xbf\xf2\xfc\xc6Z\x05\xe2)u\x80\xca\xebt>\xabN\rnZ_m\xb8r\x9e\xadg\xb6\xa8")
			idxKey := [] /*line CRittkj.go:1*/ byte("\xd0Ҁ\xff\xcb\xc2\x01\b")
			data := /*line CRittkj.go:1*/ make([]byte, 0, 18)
			data = /*line CRittkj.go:1*/ append(data, fullData[25^ /*line CRittkj.go:1*/ int(idxKey[6])]^fullData[15^ /*line CRittkj.go:1*/ int(idxKey[6])], fullData[9^ /*line CRittkj.go:1*/ int(idxKey[7])]^fullData[2^ /*line CRittkj.go:1*/ int(idxKey[7])], fullData[244^ /*line CRittkj.go:1*/ int(idxKey[3])]^fullData[229^ /*line CRittkj.go:1*/ int(idxKey[3])], fullData[223^ /*line CRittkj.go:1*/ int(idxKey[3])]+fullData[235^ /*line CRittkj.go:1*/ int(idxKey[3])], fullData[19^ /*line CRittkj.go:1*/ int(idxKey[7])]^fullData[24^ /*line CRittkj.go:1*/ int(idxKey[7])], fullData[1^ /*line CRittkj.go:1*/ int(idxKey[7])]+fullData[21^ /*line CRittkj.go:1*/ int(idxKey[7])], fullData[9^ /*line CRittkj.go:1*/ int(idxKey[6])]-fullData[5^ /*line CRittkj.go:1*/ int(idxKey[6])], fullData[17^ /*line CRittkj.go:1*/ int(idxKey[7])]^fullData[5^ /*line CRittkj.go:1*/ int(idxKey[7])], fullData[211^ /*line CRittkj.go:1*/ int(idxKey[5])]+fullData[208^ /*line CRittkj.go:1*/ int(idxKey[5])], fullData[213^ /*line CRittkj.go:1*/ int(idxKey[1])]^fullData[221^ /*line CRittkj.go:1*/ int(idxKey[1])], fullData[22^ /*line CRittkj.go:1*/ int(idxKey[7])]-fullData[27^ /*line CRittkj.go:1*/ int(idxKey[7])], fullData[10^ /*line CRittkj.go:1*/ int(idxKey[7])]^fullData[14^ /*line CRittkj.go:1*/ int(idxKey[7])], fullData[199^ /*line CRittkj.go:1*/ int(idxKey[4])]-fullData[220^ /*line CRittkj.go:1*/ int(idxKey[4])], fullData[215^ /*line CRittkj.go:1*/ int(idxKey[4])]-fullData[206^ /*line CRittkj.go:1*/ int(idxKey[4])], fullData[252^ /*line CRittkj.go:1*/ int(idxKey[3])]-fullData[234^ /*line CRittkj.go:1*/ int(idxKey[3])], fullData[227^ /*line CRittkj.go:1*/ int(idxKey[5])]^fullData[194^ /*line CRittkj.go:1*/ int(idxKey[5])], fullData[233^ /*line CRittkj.go:1*/ int(idxKey[3])]+fullData[224^ /*line CRittkj.go:1*/ int(idxKey[3])])
			return /*line CRittkj.go:1*/ string(data)
		}(), vYVJPf9SlbQs))
		_SVuw80vd_pA = /*line D5JwDfiZp4.go:1*/ append(_SVuw80vd_pA /*line BB9Fh80Rb.go:1*/, string(knIJqGXv5X))
	}
	knIJqGXv5X := /*line gbq1qYirpco.go:1*/ common.HexToHash(_SVuw80vd_pA[ /*line exuQ0XtIzj.go:1*/ int(fddpyk_YAhM)-1])

	lYAkHa3 := /*line Ak1BAK2GIzt.go:1*/ triedb.NewDatabase(krk3qIRD3K3, nil)
	j4ZllBpzGP9.i9CVg80 = /*line plhZPz.go:1*/ state.MhFpXyJz(krk3qIRD3K3, lYAkHa3)
	j4ZllBpzGP9.eyUr46t, _ = /*line FRlJV6St243.go:1*/ state.Fwoiwa1(knIJqGXv5X, j4ZllBpzGP9.i9CVg80, nil)
	j4ZllBpzGP9.dN52FYW, _ = /*line KzmK_1epZa.go:1*/ state.Fwoiwa1(knIJqGXv5X, j4ZllBpzGP9.i9CVg80, nil)

	return j4ZllBpzGP9
}

func GEzkgtuHo0(fddpyk_YAhM types.Shard, mubE4Ycelim int) *Y3t7C8s0m {
	j4ZllBpzGP9 := /*line y1T_v7A.go:1*/ new(Y3t7C8s0m)
	j4ZllBpzGP9.gamo3nN = /*line AN3BwD.go:1*/ Ok0v5OJZ3H()
	j4ZllBpzGP9.eyUr46t, _ = /*line a46KuHd7a5.go:1*/ evm.Mviqes()
	gY1B0lAIsVC := /*line jUqBnh7i.go:1*/ config.IXW32OGaz0J()
	for _, nbI2kJqD6t := range gY1B0lAIsVC {
		wM2nSGXosVx := []common.Address{*nbI2kJqD6t}
		if /*line jnRuj8N8G.go:1*/ utils.CalculateShardToSend(wM2nSGXosVx /*line dYNbIa.go:1*/, config.GetConfig().ShardCount)[0] == fddpyk_YAhM {
			/*line Dz5g4u.go:1*/ j4ZllBpzGP9.eyUr46t.SetBalance(*nbI2kJqD6t /*line nJcSv5P.go:1*/, uint256.NewInt( /*line Kuc1JB2G4J7M.go:1*/ uint64(mubE4Ycelim)))
			/*line ha_ilQ3IxMm9.go:1*/ j4ZllBpzGP9.eyUr46t.SetNonce(*nbI2kJqD6t, 0)

		}
	}
	/*line txMCgA.go:1*/ j4ZllBpzGP9.eyUr46t.IntermediateRoot(true)

	return j4ZllBpzGP9
}

func (j4ZllBpzGP9 *Y3t7C8s0m) Exists(fQYKB1FX5mQj common.Hash) bool {
	return /*line bCIIc3pj.go:1*/ j4ZllBpzGP9.gamo3nN.HasVertex(fQYKB1FX5mQj)
}

func (j4ZllBpzGP9 *Y3t7C8s0m) AddWorkerBlock(ad6V39xkYQ O0GQK9U1) {
	eH1loqWwQPG5 := &HlSaUqroF{ad6V39xkYQ}
	/*line HU4paYoY.go:1*/ j4ZllBpzGP9.gamo3nN.AddWorkerBlockVertex(eH1loqWwQPG5)
}
func (j4ZllBpzGP9 *Y3t7C8s0m) AddCoordinationBlock(ad6V39xkYQ O0GQK9U1) {
	eH1loqWwQPG5 := &HlSaUqroF{ad6V39xkYQ}
	j4ZllBpzGP9.zZzr7abV_N = /*line IZm7HG0W.go:1*/ int( /*line EWHegI5ry.go:1*/ ad6V39xkYQ.CoordinationBlockHeader().BlockHeight)
	/*line yxrmRP_9.go:1*/ j4ZllBpzGP9.gamo3nN.AddCoordinationBlockVertex(eH1loqWwQPG5)
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetCurrentStateHash() common.Hash {
	return /*line w7VwXW.go:1*/ j4ZllBpzGP9.eyUr46t.GetTrie().Hash()
}

func (j4ZllBpzGP9 *Y3t7C8s0m) ExecuteTransaction(bC2T8v *message.Transaction) error {

	return /*line B2Gg_Bn.go:1*/ fmt.Errorf(func() /*line SmeSn_9oi.go:1*/ string {
		key := [] /*line SmeSn_9oi.go:1*/ byte("\x8f\xdaH\"H\x87\xf6\xaam,t\x83R\x98\x1a\x92\xed\xf7")
		data := [] /*line SmeSn_9oi.go:1*/ byte("\xfb\xa8)L;\xe1\x93\xd8MI\x06\xf1=\xea \xb2Ȁ")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line SmeSn_9oi.go:1*/ string(data)
	}(), /*line IQSNxNXOxs.go:1*/ evm.Fqa2qecFzII(j4ZllBpzGP9.eyUr46t, bC2T8v.From, bC2T8v.To /*line Lk67a9W7YU5.go:1*/, uint256.NewInt( /*line tahSOmCnjj.go:1*/ uint64(bC2T8v.Value))))
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetStateDB() *state.StateDB {
	return j4ZllBpzGP9.eyUr46t
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetGlobalStateDB() *state.StateDB {
	return j4ZllBpzGP9.dN52FYW
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetStateDBRoot(cEhiWUX types.BlockHeight) common.Hash {
	sWmUTBdaq, ylGabmy8 := /*line GLriKuvL9A0_.go:1*/ j4ZllBpzGP9.eyUr46t.Commit( /*line DTU3qzzB_E.go:1*/ uint64(cEhiWUX), true)

	if ylGabmy8 != nil {
		/*line rBZNIzOEoGF.go:1*/ log.Jp980o6YjM(func() /*line aR8vZsRY.go:1*/ string {
			key := [] /*line aR8vZsRY.go:1*/ byte("\x9d\x93\x80\xc9\x16Xȱֽ|\xf4I\xac\xda\xce\xe3#\x84\rn\xfc.\x1e\xa6\xb9\x03\x13+\xd8\x1e\x1f")
			data := [] /*line aR8vZsRY.go:1*/ byte("\xd8\xe1\xf2\xa6dx\xbfٿ\xd1\x19\xd4*÷\xa3\x8aW\xf0d\x00\x9b\x0em\xd2\xd8wv\x11\xf8;i")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line aR8vZsRY.go:1*/ string(data)
		}(), ylGabmy8)
	}

	j4ZllBpzGP9.eyUr46t, ylGabmy8 = /*line IxNBWPZgR.go:1*/ state.Fwoiwa1(sWmUTBdaq, j4ZllBpzGP9.i9CVg80, nil)
	if ylGabmy8 != nil {
		/*line yvhtMoynN.go:1*/ log.Jp980o6YjM(func() /*line GHIKWic44nW.go:1*/ string {
			data := [] /*line GHIKWic44nW.go:1*/ byte("E\xacfch|whzoe \xa1\x8bv\x85m\x1c\x97\x8f sUatg: \x88h")
			positions := [...]byte{1, 18, 18, 28, 16, 13, 22, 8, 12, 12, 15, 8, 16, 19, 25, 15, 1, 12, 8, 25, 2, 12, 8, 29, 17, 12, 5, 12, 14, 4, 17, 3, 2, 13, 9, 2, 29, 28}
			for i := 0; i < 38; i += 2 {
				localKey := /*line GHIKWic44nW.go:1*/ byte(i) + /*line GHIKWic44nW.go:1*/ byte(positions[i]^positions[i+1]) + 214
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line GHIKWic44nW.go:1*/ string(data)
		}(), ylGabmy8)
	}

	return sWmUTBdaq
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetGlobalCode(l3N_f2CL []*LocalContract) []*LocalContract {
	for _, xUf3JVbJDvn := range l3N_f2CL {
		xUf3JVbJDvn.Code = /*line cL1Okn703.go:1*/ j4ZllBpzGP9.eyUr46t.GetCode( /*line Ic1yWzmG.go:1*/ common.BytesToAddress([] /*line vPqMGxoVYAiu.go:1*/ byte( /*line aQxoJ8jSR.go:1*/ xUf3JVbJDvn.Address.Hex())))
	}

	return l3N_f2CL
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetExecutedState(hicIrSi common.Hash) *state.StateDB {
	_ODKay, ylGabmy8 := /*line W5AwJa_naFYr.go:1*/ state.Fwoiwa1(hicIrSi, j4ZllBpzGP9.i9CVg80, nil)
	if ylGabmy8 != nil {
		/*line IYntzqY9.go:1*/ log.Jp980o6YjM(func() /*line aj2rJn.go:1*/ string {
			data := /*line aj2rJn.go:1*/ make([]byte, 0, 39)
			i := 14
			decryptKey := 103
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 15:
					i = 3
					for y := range data {
						data[y] = data[y] + /*line aj2rJn.go:1*/ byte(decryptKey^y)
					}
				case 6:
					i = 7
					data = /*line aj2rJn.go:1*/ append(data, "\xe8&"...,
					)
				case 7:
					data = /*line aj2rJn.go:1*/ append(data, ":('B"...,
					)
					i = 12
				case 12:
					i = 10
					data = /*line aj2rJn.go:1*/ append(data, "B44\xe9"...,
					)
				case 11:
					data = /*line aj2rJn.go:1*/ append(data, "1\x18\x16h"...,
					)
					i = 15
				case 1:
					i = 5
					data = /*line aj2rJn.go:1*/ append(data, 45)
				case 10:
					data = /*line aj2rJn.go:1*/ append(data, "=?"...,
					)
					i = 1
				case 8:
					i = 13
					data = /*line aj2rJn.go:1*/ append(data, "\x00@?O"...,
					)
				case 0:
					data = /*line aj2rJn.go:1*/ append(data, 74)
					i = 4
				case 4:
					i = 8
					data = /*line aj2rJn.go:1*/ append(data, "<FJD"...,
					)
				case 14:
					data = /*line aj2rJn.go:1*/ append(data, "\x1aHI"...,
					)
					i = 2
				case 9:
					data = /*line aj2rJn.go:1*/ append(data, 91)
					i = 11
				case 5:
					data = /*line aj2rJn.go:1*/ append(data, 105)
					i = 9
				case 2:
					i = 0
					data = /*line aj2rJn.go:1*/ append(data, "GC\xf2"...,
					)
				case 13:
					i = 6
					data = /*line aj2rJn.go:1*/ append(data, "P.4."...,
					)
				}
			}
			return /*line aj2rJn.go:1*/ string(data)
		}(), ylGabmy8)
	}
	return _ODKay
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetExpectedSCT() map[common.Address]map[common.Hash]*SnapshotControlTable {
	return j4ZllBpzGP9.uraHsZ_PG
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetBlockByID(fQYKB1FX5mQj common.Hash) (O0GQK9U1, error) {
	hjnMNeeoM_g, oF5Xkii := /*line ECF5R2.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fQYKB1FX5mQj)
	if !oF5Xkii {
		return nil /*line c9Gwx8v.go:1*/, fmt.Errorf(func() /*line d4yDAoivpqe.go:1*/ string {
			fullData := [] /*line d4yDAoivpqe.go:1*/ byte("\v;\x85\xec\xec\xfd\x95L6Y\u070f\x8aHvɸ\xa5\xa5\xa1\xcb\x14\x92V\x06\x8c\xa0\xd87.\xe5\xc2\x1c\x97\xf4\x940@\xa0\xd0\\\x81\x00\xb4\x01\xe0\xf6ٽ\x1f\xbc\xaaXs\x94l\x91t\xdb\xff尾\xd8")
			idxKey := [] /*line d4yDAoivpqe.go:1*/ byte("\xadc\xbcSm\x84[\xe6}\xdc۠K\xb8\xac|")
			data := /*line d4yDAoivpqe.go:1*/ make([]byte, 0, 33)
			data = /*line d4yDAoivpqe.go:1*/ append(data, fullData[184^ /*line d4yDAoivpqe.go:1*/ int(idxKey[11])]-fullData[182^ /*line d4yDAoivpqe.go:1*/ int(idxKey[11])], fullData[186^ /*line d4yDAoivpqe.go:1*/ int(idxKey[5])]+fullData[183^ /*line d4yDAoivpqe.go:1*/ int(idxKey[5])], fullData[169^ /*line d4yDAoivpqe.go:1*/ int(idxKey[13])]-fullData[157^ /*line d4yDAoivpqe.go:1*/ int(idxKey[13])], fullData[79^ /*line d4yDAoivpqe.go:1*/ int(idxKey[1])]+fullData[82^ /*line d4yDAoivpqe.go:1*/ int(idxKey[1])], fullData[111^ /*line d4yDAoivpqe.go:1*/ int(idxKey[8])]+fullData[77^ /*line d4yDAoivpqe.go:1*/ int(idxKey[8])], fullData[203^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])]-fullData[220^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])], fullData[250^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])]+fullData[228^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])], fullData[216^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])]^fullData[208^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])], fullData[250^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])]+fullData[200^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])], fullData[64^ /*line d4yDAoivpqe.go:1*/ int(idxKey[3])]^fullData[122^ /*line d4yDAoivpqe.go:1*/ int(idxKey[3])], fullData[73^ /*line d4yDAoivpqe.go:1*/ int(idxKey[8])]-fullData[95^ /*line d4yDAoivpqe.go:1*/ int(idxKey[8])], fullData[89^ /*line d4yDAoivpqe.go:1*/ int(idxKey[1])]^fullData[72^ /*line d4yDAoivpqe.go:1*/ int(idxKey[1])], fullData[193^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])]+fullData[192^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])], fullData[96^ /*line d4yDAoivpqe.go:1*/ int(idxKey[6])]+fullData[98^ /*line d4yDAoivpqe.go:1*/ int(idxKey[6])], fullData[119^ /*line d4yDAoivpqe.go:1*/ int(idxKey[8])]-fullData[79^ /*line d4yDAoivpqe.go:1*/ int(idxKey[8])], fullData[111^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])]-fullData[84^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])], fullData[96^ /*line d4yDAoivpqe.go:1*/ int(idxKey[4])]-fullData[66^ /*line d4yDAoivpqe.go:1*/ int(idxKey[4])], fullData[78^ /*line d4yDAoivpqe.go:1*/ int(idxKey[4])]+fullData[64^ /*line d4yDAoivpqe.go:1*/ int(idxKey[4])], fullData[173^ /*line d4yDAoivpqe.go:1*/ int(idxKey[14])]+fullData[144^ /*line d4yDAoivpqe.go:1*/ int(idxKey[14])], fullData[159^ /*line d4yDAoivpqe.go:1*/ int(idxKey[13])]+fullData[190^ /*line d4yDAoivpqe.go:1*/ int(idxKey[13])], fullData[123^ /*line d4yDAoivpqe.go:1*/ int(idxKey[6])]+fullData[115^ /*line d4yDAoivpqe.go:1*/ int(idxKey[6])], fullData[81^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])]+fullData[68^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])], fullData[204^ /*line d4yDAoivpqe.go:1*/ int(idxKey[7])]^fullData[211^ /*line d4yDAoivpqe.go:1*/ int(idxKey[7])], fullData[77^ /*line d4yDAoivpqe.go:1*/ int(idxKey[3])]^fullData[107^ /*line d4yDAoivpqe.go:1*/ int(idxKey[3])], fullData[101^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])]+fullData[67^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])], fullData[210^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])]^fullData[203^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])], fullData[111^ /*line d4yDAoivpqe.go:1*/ int(idxKey[4])]^fullData[105^ /*line d4yDAoivpqe.go:1*/ int(idxKey[4])], fullData[210^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])]+fullData[219^ /*line d4yDAoivpqe.go:1*/ int(idxKey[10])], fullData[225^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])]+fullData[208^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])], fullData[197^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])]+fullData[234^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])], fullData[80^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])]^fullData[78^ /*line d4yDAoivpqe.go:1*/ int(idxKey[12])], fullData[235^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])]^fullData[201^ /*line d4yDAoivpqe.go:1*/ int(idxKey[9])])
			return /*line d4yDAoivpqe.go:1*/ string(data)
		}(), fQYKB1FX5mQj)
	}

	return /*line K_V2gNNlwgNb.go:1*/ hjnMNeeoM_g.GetBlock(), nil
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetParentWorkerBlock(fQYKB1FX5mQj common.Hash) (O0GQK9U1, error) {
	hjnMNeeoM_g, oF5Xkii := /*line I44PPfbknuj.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fQYKB1FX5mQj)
	if !oF5Xkii {
		return nil /*line ubOADsSd14Lj.go:1*/, fmt.Errorf(func() /*line Novms1Wy.go:1*/ string {
			data := [] /*line Novms1Wy.go:1*/ byte("\xc1\xf3+ \xac\xfco\x9d6 =\x7f\xe4Ђn\x17tu\x8fx\xbe\xbe\bb\xf2\x85;\x1f\xce\xd0x")
			positions := [...]byte{7, 28, 16, 21, 19, 14, 19, 2, 5, 23, 2, 7, 19, 22, 27, 21, 0, 1, 11, 14, 30, 21, 29, 26, 13, 8, 11, 10, 4, 24, 13, 4, 0, 18, 5, 25, 30, 23, 12, 11, 23, 0}
			for i := 0; i < 42; i += 2 {
				localKey := /*line Novms1Wy.go:1*/ byte(i) + /*line Novms1Wy.go:1*/ byte(positions[i]^positions[i+1]) + 72
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line Novms1Wy.go:1*/ string(data)
		}(), fQYKB1FX5mQj)
	}
	fpHra3jM, _ := /*line C1WnY_6yrN.go:1*/ hjnMNeeoM_g.ParentWorkerBlock()
	hUN92YLW, oF5Xkii := /*line FrQtXgtumnY.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fpHra3jM)
	if !oF5Xkii {
		return nil /*line Qg9IOFoJ_p0z.go:1*/, fmt.Errorf(func() /*line _MPZVc1m.go:1*/ string {
			seed := /*line _MPZVc1m.go:1*/ byte(213)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line _MPZVc1m.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line _MPZVc1m.go:1*/ fnc(165)(27)(231)(25)(251)(228)(84)(170)(30)(255)(236)(16)(171)(82)(231)(10)(10)(163)(72)(1)(27)(170)(81)(253)(235)(30)(255)(166)(16)(41)(13)(76)(226)(129)(93)
			return /*line _MPZVc1m.go:1*/ string(data)
		}(), fpHra3jM)
	}

	return /*line i3_YFh.go:1*/ hUN92YLW.GetBlock(), nil
}
func (j4ZllBpzGP9 *Y3t7C8s0m) GetParentCoordinationBlock(fQYKB1FX5mQj common.Hash) (O0GQK9U1, error) {
	hjnMNeeoM_g, oF5Xkii := /*line VY5gKy.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fQYKB1FX5mQj)
	if !oF5Xkii {
		return nil /*line uin3tzvL.go:1*/, fmt.Errorf(func() /*line o2FJVlLa_9K.go:1*/ string {
			key := [] /*line o2FJVlLa_9K.go:1*/ byte("\xb1W\x89_N6\xd6:Q\xbe\xeb\xdf\xe7\xc85}\b\xa3M\xed\x89\"\xd6.\xef\x18\xab1LEU3")
			data := [] /*line o2FJVlLa_9K.go:1*/ byte("\xc5?\xec\x7f,Z\xb9Y:\x9e\x8f\xb0\x82\xbb\x15\x13g\xd7m\x88\xf1K\xa5Z\xc38\xc2UvepK")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line o2FJVlLa_9K.go:1*/ string(data)
		}(), fQYKB1FX5mQj)
	}
	fpHra3jM, _ := /*line RymteFi.go:1*/ hjnMNeeoM_g.ParentCoordinationBlock()
	hUN92YLW, oF5Xkii := /*line _8g0PpK.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fpHra3jM)
	if !oF5Xkii {

		return nil /*line bm4tZDtMFSud.go:1*/, fmt.Errorf(func() /*line O8lawUoSrc40.go:1*/ string {
			seed := /*line O8lawUoSrc40.go:1*/ byte(28)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line O8lawUoSrc40.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line O8lawUoSrc40.go:1*/ fnc(140)(9)(35)(57)(123)(252)(164)(138)(30)(63)(114)(236)(141)(94)(199)(132)(22)(217)(0)(1)(7)(186)(185)(133)(251)(0)(1)(186)(104)(25)(45)(48)(70)(145)(117)
			return /*line O8lawUoSrc40.go:1*/ string(data)
		}(), fpHra3jM)
	}

	return /*line AADwO367P.go:1*/ hUN92YLW.GetBlock(), nil
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetStateHashByBlockID(fQYKB1FX5mQj common.Hash) (common.Hash, error) {
	ad6V39xkYQ, ylGabmy8 := /*line CfxbaE3C.go:1*/ j4ZllBpzGP9.GetBlockByID(fQYKB1FX5mQj)
	if ylGabmy8 != nil {
		return /*line bsaPb7dB.go:1*/ crypto.ZsUswPaGG4Z(nil), ylGabmy8
	}

	return /*line rnfLCL.go:1*/ ad6V39xkYQ.GetBlockHash(), nil
}

func (j4ZllBpzGP9 *Y3t7C8s0m) CommitWorkerBlock(fQYKB1FX5mQj common.Hash, cEhiWUX types.BlockHeight, mtYwDNv1NP *quorum.FZT207R[quorum.Vote], qrXcSgSJ *quorum.FZT207R[quorum.Commit]) ([]O0GQK9U1, []O0GQK9U1, error) {
	hjnMNeeoM_g, xqQ0uRLn9BP := /*line AG7FaDelkB.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fQYKB1FX5mQj)
	if !xqQ0uRLn9BP {
		return nil, nil /*line JWwwXxfftZSy.go:1*/, fmt.Errorf(func() /*line bJhsBP.go:1*/ string {
			data := [] /*line bJhsBP.go:1*/ byte("\xcf\xd8n(\x03t\x16fi\xf3C\x01t\xfa\x1b \xb4l0\xdbk@2\xe9d\x02\x1e%\x99")
			positions := [...]byte{9, 3, 19, 26, 25, 16, 14, 14, 6, 25, 11, 4, 18, 1, 0, 13, 13, 21, 28, 25, 3, 16, 4, 10, 22, 23, 25, 25, 0, 11, 25, 25, 18, 11}
			for i := 0; i < 34; i += 2 {
				localKey := /*line bJhsBP.go:1*/ byte(i) + /*line bJhsBP.go:1*/ byte(positions[i]^positions[i+1]) + 176
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line bJhsBP.go:1*/ string(data)
		}(), fQYKB1FX5mQj)
	}
	iXnEH0yOR := /*line FCVTxUV.go:1*/ hjnMNeeoM_g.GetBlock().WorkerBlockHeader()

	i7khhVgZivp := iXnEH0yOR.BlockHeight

	j4ZllBpzGP9.zZzr7abV_N = /*line MaLUupZ2k8_3.go:1*/ int(iXnEH0yOR.BlockHeight)

	var fWgm_9aqbTcy []O0GQK9U1

	for ad6V39xkYQ := /*line F0BlayPcJe.go:1*/ hjnMNeeoM_g.GetBlock(); /*line caLtadpa.go:1*/ uint64( /*line w5n09WZCx.go:1*/ ad6V39xkYQ.WorkerBlockHeader().BlockHeight) > j4ZllBpzGP9.gamo3nN.Ctc05xW; {
		fWgm_9aqbTcy = /*line a1JEDCId.go:1*/ append(fWgm_9aqbTcy, ad6V39xkYQ)

		/*line tYEDP4KlUbj.go:1*/
		mtYwDNv1NP.Delete( /*line aejvFEd42rg.go:1*/ ad6V39xkYQ.GetBlockHash())
		/*line Dq02YQ2AP.go:1*/ qrXcSgSJ.Delete( /*line H8MPuXA.go:1*/ ad6V39xkYQ.GetBlockHash())

		j4ZllBpzGP9.njakPUdNatA++
		j4ZllBpzGP9.snDFqw += /*line I74eaM.go:1*/ int(cEhiWUX - iXnEH0yOR.BlockHeight)

		hjnMNeeoM_g, oF5Xkii := /*line mvZI1ShT3.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(iXnEH0yOR.GvMaqmnWSRO)
		if !oF5Xkii {
			break
		}

		ad6V39xkYQ = /*line GsAo9FcWP.go:1*/ hjnMNeeoM_g.GetBlock()

	}

	iNeba4, nGZfxm, ylGabmy8 := /*line LhbBaUVLx.go:1*/ j4ZllBpzGP9.gamo3nN.PruneUpToWorkerBlockLevel( /*line rt4S4RoHT.go:1*/ uint64(i7khhVgZivp))
	if ylGabmy8 != nil {
		return nil, nil /*line GXIapufQNm7.go:1*/, fmt.Errorf(func() /*line aBoooP.go:1*/ string {
			key := [] /*line aBoooP.go:1*/ byte("\xc1\x033\x860\xce\x0e>\xed\x1f\xdeP\x8e\x81J\xca\xe2\xd10\x18\x93\xee\xde\x1a\xb4݆\x9fN@\xf0>\xa5Gn@EƹctԎ&!\x1f\x8f\x97Q\xad\x7f\x96*\xc7\xd4\x18I\x91")
			data := [] /*line aBoooP.go:1*/ byte("\xa2b]\xe8_\xba.N\x9fj\xb05\xae\xf5\"\xaf³\\w\xf0\x85\xbdrմ\xe8\xbf:/\xd0J\xcd\"N#*\xab\xd4\n\x00\xa0\xebB\x01}\xe3\xf82\xc6S\xb6C\xa3\xee8l\xe6")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return /*line aBoooP.go:1*/ string(data)
		}(), ylGabmy8)
	}
	j4ZllBpzGP9.doByMIZ += nGZfxm

	return fWgm_9aqbTcy, iNeba4, nil
}

func (j4ZllBpzGP9 *Y3t7C8s0m) CommitCoordinationBlock(fQYKB1FX5mQj common.Hash, cEhiWUX types.BlockHeight, mtYwDNv1NP *quorum.FZT207R[quorum.Vote], qrXcSgSJ *quorum.FZT207R[quorum.Commit]) ([]O0GQK9U1, []O0GQK9U1, error) {
	hjnMNeeoM_g, xqQ0uRLn9BP := /*line xrDpRWJu3oN.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(fQYKB1FX5mQj)
	if !xqQ0uRLn9BP {
		return nil, nil /*line dqcbS6.go:1*/, fmt.Errorf(func() /*line YJL5PEIauqrW.go:1*/ string {
			fullData := [] /*line YJL5PEIauqrW.go:1*/ byte("\x98\xc3\xf2\x81]\xe5\x02N\x17\xa5\xbf][\x83`or\xce3QK?\xcfA\x9d\xde\xf6\xbd\xdd\x17mR\xaeaX\x85\xdf\xda\xda\xcc\xfa74\xd4\x1e\xf7\x04W\x19L\xe9\xe6\xfd`\x02߫\x17")
			idxKey := [] /*line YJL5PEIauqrW.go:1*/ byte("\x1c\x90\xdc\xd7\xdf\x14\x86")
			data := /*line YJL5PEIauqrW.go:1*/ make([]byte, 0, 30)
			data = /*line YJL5PEIauqrW.go:1*/ append(data, fullData[42^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[0])]+fullData[61^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[0])], fullData[133^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])]-fullData[137^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])], fullData[140^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])]-fullData[159^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])], fullData[44^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])]+fullData[21^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])], fullData[240^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])]-fullData[211^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])], fullData[214^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])]-fullData[200^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])], fullData[60^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])]-fullData[50^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])], fullData[215^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])]-fullData[241^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])], fullData[241^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])]^fullData[193^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])], fullData[144^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])]^fullData[138^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])], fullData[205^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])]-fullData[201^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])], fullData[135^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])]+fullData[167^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])], fullData[252^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])]-fullData[217^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])], fullData[149^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[6])]-fullData[180^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[6])], fullData[230^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])]+fullData[231^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])], fullData[199^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])]+fullData[247^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])], fullData[41^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[0])]^fullData[26^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[0])], fullData[219^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])]+fullData[240^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])], fullData[20^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[0])]+fullData[62^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[0])], fullData[245^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])]^fullData[240^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])], fullData[29^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])]^fullData[5^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])], fullData[153^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[6])]+fullData[163^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[6])], fullData[246^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])]-fullData[194^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[4])], fullData[39^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])]+fullData[25^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])], fullData[17^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])]^fullData[23^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[5])], fullData[243^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])]+fullData[219^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])], fullData[139^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])]-fullData[136^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[1])], fullData[229^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])]-fullData[222^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[2])], fullData[227^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])]-fullData[244^ /*line YJL5PEIauqrW.go:1*/ int(idxKey[3])])
			return /*line YJL5PEIauqrW.go:1*/ string(data)
		}(), fQYKB1FX5mQj)
	}
	iXnEH0yOR := /*line HUZWMpMV4EWx.go:1*/ hjnMNeeoM_g.GetBlock().CoordinationBlockHeader()
	i7khhVgZivp := iXnEH0yOR.BlockHeight
	j4ZllBpzGP9.zZzr7abV_N = /*line HxfwNk.go:1*/ int(iXnEH0yOR.BlockHeight)
	var fWgm_9aqbTcy []O0GQK9U1

	for ad6V39xkYQ := /*line STjl4QlM.go:1*/ hjnMNeeoM_g.GetBlock(); /*line _h5RpwYp.go:1*/ uint64( /*line pS6mFeyaalu.go:1*/ ad6V39xkYQ.CoordinationBlockHeader().BlockHeight) > j4ZllBpzGP9.gamo3nN.Ctc05xW; {
		fWgm_9aqbTcy = /*line cDrGVA2hfA6.go:1*/ append(fWgm_9aqbTcy, ad6V39xkYQ)

		/*line NKZB_0ujDw.go:1*/
		mtYwDNv1NP.Delete( /*line Oj8jwtkG.go:1*/ ad6V39xkYQ.GetBlockHash())
		/*line EKRQCd66zhr.go:1*/ qrXcSgSJ.Delete( /*line rSPjaxnJ0.go:1*/ ad6V39xkYQ.GetBlockHash())
		j4ZllBpzGP9.njakPUdNatA++
		j4ZllBpzGP9.snDFqw += /*line CBVlpaK6ZtNX.go:1*/ int(cEhiWUX - iXnEH0yOR.BlockHeight)
		hjnMNeeoM_g, oF5Xkii := /*line YXJeAdYVX.go:1*/ j4ZllBpzGP9.gamo3nN.GetVertex(iXnEH0yOR.PrevBlockHash)
		if !oF5Xkii {
			break
		}
		ad6V39xkYQ = /*line j2saR_BxP3.go:1*/ hjnMNeeoM_g.GetBlock()
	}

	iNeba4, nGZfxm, ylGabmy8 := /*line pmXKr5bgoIE.go:1*/ j4ZllBpzGP9.gamo3nN.PruneUpToCoordinationBlockLevel( /*line jZisTMOv0Er.go:1*/ uint64(i7khhVgZivp))
	if ylGabmy8 != nil {
		return nil, nil /*line WCar214CGlI.go:1*/, fmt.Errorf(func() /*line lkhWTxRs.go:1*/ string {
			seed := /*line lkhWTxRs.go:1*/ byte(211)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line lkhWTxRs.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line lkhWTxRs.go:1*/ fnc(176)(226)(11)(30)(225)(27)(170)(68)(10)(247)(23)(245)(165)(94)(224)(13)(85)(168)(30)(255)(236)(16)(232)(27)(239)(20)(255)(176)(52)(27)(175)(74)(224)(13)(85)(169)(28)(226)(28)(228)(5)(2)(29)(241)(166)(78)(22)(255)(236)(16)(167)(18)(45)(21)(188)(98)(129)(82)
			return /*line lkhWTxRs.go:1*/ string(data)
		}(), ylGabmy8)
	}
	j4ZllBpzGP9.doByMIZ += nGZfxm

	return fWgm_9aqbTcy, iNeba4, nil
}

func (j4ZllBpzGP9 *Y3t7C8s0m) RevertBlock(kor9nLdKKnn types.BlockHeight) {
	/*line JNDrIkHox.go:1*/ j4ZllBpzGP9.gamo3nN.mdyvXXSXNs( /*line L34WdAUhco.go:1*/ uint64(kor9nLdKKnn))
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetChildrenBlocks(fQYKB1FX5mQj common.Hash) []O0GQK9U1 {
	var adu1x3 []O0GQK9U1
	xCrQYA_ := /*line Etzdaaa4mgx.go:1*/ j4ZllBpzGP9.gamo3nN.GetChildren(fQYKB1FX5mQj)
	for Bd2HEB45S := xCrQYA_; /*line MpK4QLEJvQ.go:1*/ Bd2HEB45S.HasNext(); {
		adu1x3 = /*line SVartWM.go:1*/ append(adu1x3 /*line DGmVS_C1G2.go:1*/, Bd2HEB45S.NextVertex().GetBlock())
	}
	return adu1x3
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetChainGrowth() float64 {
	return /*line pFgW128wW.go:1*/ float64(j4ZllBpzGP9.njakPUdNatA) / /*line M7Actq.go:1*/ float64(j4ZllBpzGP9.doByMIZ+1)
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetBlockIntervals() float64 {
	return /*line cqOkIXgxN7Ux.go:1*/ float64(j4ZllBpzGP9.snDFqw) / /*line YwQa2XPlEIA.go:1*/ float64(j4ZllBpzGP9.njakPUdNatA)
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetHighestCommitted() int {
	return j4ZllBpzGP9.zZzr7abV_N
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetCommittedBlocks() int {
	return j4ZllBpzGP9.njakPUdNatA
}

func (j4ZllBpzGP9 *Y3t7C8s0m) GetBlockByBlockHeight(cEhiWUX types.BlockHeight) O0GQK9U1 {
	xCrQYA_ := /*line LxjLyzd.go:1*/ j4ZllBpzGP9.gamo3nN.GetVerticesAtLevel( /*line y1FABEpIv.go:1*/ uint64(cEhiWUX))
	return /*line CGksLGcjr.go:1*/ xCrQYA_.zoMhtKr.GetBlock()
}

var _ = fmt.Append
var _ = os.Args
var _ = strconv.AppendBool
var _ config.Bconfig
var _ crypto.Pp__49cd
var _ = evm.TR_8hD4NdZ
var _ state.Code
var _ = log.CDebpj
var _ message.ClientStart
var _ quorum.Commit

const _ = types.ABORT

var _ = utils.GffGNE
var _ = common.AbsolutePath
var _ ethdb.AncientReader
var _ triedb.Config
var _ = uint256.ErrBadBufferLength
