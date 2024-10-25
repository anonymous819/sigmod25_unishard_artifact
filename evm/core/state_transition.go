//line :1
package core

import (
	"fmt"
	"math"
	"math/big"

	types "unishard/evm/types"
	vm "unishard/evm/vm"

	"github.com/ethereum/go-ethereum/common"
	cmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type YHD4JD4i7 struct {
	A7_zDWIb   uint64
	ZY9Ly6apaW uint64
	BahScRGbl  error
	A3aOS4Y4   []byte
}

func (kzCAWNp7 *YHD4JD4i7) Unwrap() error {
	return kzCAWNp7.BahScRGbl
}

func (kzCAWNp7 *YHD4JD4i7) Failed() bool { return kzCAWNp7.BahScRGbl != nil }

func (kzCAWNp7 *YHD4JD4i7) Return() []byte {
	if kzCAWNp7.BahScRGbl != nil {
		return nil
	}
	return /*line SqZF9Y_3BH.go:1*/ common.CopyBytes(kzCAWNp7.A3aOS4Y4)
}

func (kzCAWNp7 *YHD4JD4i7) Revert() []byte {
	if kzCAWNp7.BahScRGbl != vm.T_LmvOGib {
		return nil
	}
	return /*line QxgjZJYiyr.go:1*/ common.CopyBytes(kzCAWNp7.A3aOS4Y4)
}

func QCDeiF8gmvU(rVnDLTx5 []byte, gec_2ITSwoS types.AccessList, pwLQ5O_S4Ti bool, wDICbt4XMZ, rBK8D4mVo, a4RnqwA62PI bool) (uint64, error) {

	var mtuuPqNr_M uint64
	if pwLQ5O_S4Ti && wDICbt4XMZ {
		mtuuPqNr_M = params.TxGasContractCreation
	} else {
		mtuuPqNr_M = params.TxGas
	}
	gK7TpxC4 := /*line PYn3TnvsD.go:1*/ uint64( /*line ZHV4gW_cI.go:1*/ len(rVnDLTx5))

	if gK7TpxC4 > 0 {

		var g8ZpToHH uint64
		for _, o3UZwsbu41P := range rVnDLTx5 {
			if o3UZwsbu41P != 0 {
				g8ZpToHH++
			}
		}

		bF7BuagElLC := params.TxDataNonZeroGasFrontier
		if rBK8D4mVo {
			bF7BuagElLC = params.TxDataNonZeroGasEIP2028
		}
		if (math.MaxUint64-mtuuPqNr_M)/bF7BuagElLC < g8ZpToHH {
			return 0, ZuKrAa8z_b8
		}
		mtuuPqNr_M += g8ZpToHH * bF7BuagElLC

		p6_xWv := gK7TpxC4 - g8ZpToHH
		if (math.MaxUint64-mtuuPqNr_M)/params.TxDataZeroGas < p6_xWv {
			return 0, ZuKrAa8z_b8
		}
		mtuuPqNr_M += p6_xWv * params.TxDataZeroGas

		if pwLQ5O_S4Ti && a4RnqwA62PI {
			dTHjGgqrd := /*line l100ETbvS.go:1*/ lEao6DA(gK7TpxC4)
			if (math.MaxUint64-mtuuPqNr_M)/params.InitCodeWordGas < dTHjGgqrd {
				return 0, ZuKrAa8z_b8
			}
			mtuuPqNr_M += dTHjGgqrd * params.InitCodeWordGas
		}
	}
	if gec_2ITSwoS != nil {
		mtuuPqNr_M += /*line WyjI2kMHL.go:1*/ uint64( /*line AmzcL3H.go:1*/ len(gec_2ITSwoS)) * params.TxAccessListAddressGas
		mtuuPqNr_M += /*line u7aR0Fh.go:1*/ uint64( /*line tzEji6.go:1*/ gec_2ITSwoS.StorageKeys()) * params.TxAccessListStorageKeyGas
	}
	return mtuuPqNr_M, nil
}

func lEao6DA(niFNN0AcIr uint64) uint64 {
	if niFNN0AcIr > math.MaxUint64-31 {
		return math.MaxUint64/32 + 1
	}

	return (niFNN0AcIr + 31) / 32
}

type Message struct {
	To            *common.Address
	From          common.Address
	Nonce         uint64
	Value         *big.Int
	GasLimit      uint64
	GasPrice      *big.Int
	GasFeeCap     *big.Int
	GasTipCap     *big.Int
	Data          []byte
	AccessList    types.AccessList
	BlobGasFeeCap *big.Int
	BlobHashes    []common.Hash

	SkipAccountChecks bool
}

func RDwaGH8G(nWrNkD6jsTa *types.Transaction, aG9YbQZeaqCa types.NlfWgA7, q1VdXTanDXB *big.Int) (*Message, error) {
	rrEywqc_se := &Message{
		Nonce:/*line Ya89d9Xx_E.go:1*/ nWrNkD6jsTa.Nonce(),
		GasLimit:/*line RlXmgF5.go:1*/ nWrNkD6jsTa.Gas(),
		GasPrice:/*line dRxbkX.go:1*/ new(big.Int).Set( /*line CsTAkRhgRn2K.go:1*/ nWrNkD6jsTa.GasPrice()),
		GasFeeCap:/*line KTBWVvXihIm.go:1*/ new(big.Int).Set( /*line jiu4PM79ox.go:1*/ nWrNkD6jsTa.GasFeeCap()),
		GasTipCap:/*line OX1kqJrDwZyS.go:1*/ new(big.Int).Set( /*line NncNa3ZKp.go:1*/ nWrNkD6jsTa.GasTipCap()),
		To:/*line x2xtaz.go:1*/ nWrNkD6jsTa.To(),
		Value:/*line JvJZdFFbz.go:1*/ nWrNkD6jsTa.Value(),
		Data:/*line geplZOSqa12B.go:1*/ nWrNkD6jsTa.Data(),
		AccessList:/*line JKDFTau.go:1*/ nWrNkD6jsTa.AccessList(),
		SkipAccountChecks: false,
		BlobHashes:/*line F1FkdNFwzkdg.go:1*/ nWrNkD6jsTa.BlobHashes(),
		BlobGasFeeCap:/*line RJThhB3HZs.go:1*/ nWrNkD6jsTa.BlobGasFeeCap(),
	}

	if q1VdXTanDXB != nil {
		rrEywqc_se.GasPrice = /*line NAjiYLpXV7zh.go:1*/ cmath.BigMin( /*line OTVfBKFP1_.go:1*/ rrEywqc_se.GasPrice.Add(rrEywqc_se.GasTipCap, q1VdXTanDXB), rrEywqc_se.GasFeeCap)
	}
	var beOvhSsCC3c error
	rrEywqc_se.From, beOvhSsCC3c = /*line FQCmzWKa.go:1*/ types.Zsrn5l(aG9YbQZeaqCa, nWrNkD6jsTa)
	return rrEywqc_se, beOvhSsCC3c
}

func UaA5Qe(eVip3B2X *vm.EVM, rrEywqc_se *Message, mxECAEZXrJtk *GasPool) (*YHD4JD4i7, error) {
	return /*line JjrZR4dP.go:1*/ DWkacNkh(eVip3B2X, rrEywqc_se, mxECAEZXrJtk).TransitionDb()
}

type StateTransition struct {
	gp           *GasPool
	msg          *Message
	gasRemaining uint64
	initialGas   uint64
	state        vm.StateDB
	evm          *vm.EVM
}

func DWkacNkh(eVip3B2X *vm.EVM, rrEywqc_se *Message, mxECAEZXrJtk *GasPool) *StateTransition {
	return &StateTransition{
		gp:    mxECAEZXrJtk,
		evm:   eVip3B2X,
		msg:   rrEywqc_se,
		state: eVip3B2X.StateDB,
	}
}

func (azD5aZP7c *StateTransition) rJ6LH9n() common.Address {
	if azD5aZP7c.msg == nil || azD5aZP7c.msg.To == nil {
		return common.Address{}
	}
	return *azD5aZP7c.msg.To
}

func (azD5aZP7c *StateTransition) fvZooT7() error {
	fRvA7i := /*line F9qm6X.go:1*/ new(big.Int).SetUint64(azD5aZP7c.msg.GasLimit)
	fRvA7i = /*line N3vb0J1baQvN.go:1*/ fRvA7i.Mul(fRvA7i, azD5aZP7c.msg.GasPrice)
	aaIqwx := /*line B5fn_Lyci.go:1*/ new(big.Int).Set(fRvA7i)
	if azD5aZP7c.msg.GasFeeCap != nil {
		/*line Wl982rZsC9.go:1*/ aaIqwx.SetUint64(azD5aZP7c.msg.GasLimit)
		aaIqwx = /*line nBjezTOTNC.go:1*/ aaIqwx.Mul(aaIqwx, azD5aZP7c.msg.GasFeeCap)
		/*line o96djB.go:1*/ aaIqwx.Add(aaIqwx, azD5aZP7c.msg.Value)
	}
	if /*line gKWwe1.go:1*/ azD5aZP7c.evm.ChainConfig().IsCancun(azD5aZP7c.evm.Context.BlockNumber, azD5aZP7c.evm.Context.Time) {
		if o2Yia7q := /*line Vy0y6Wbvpy.go:1*/ azD5aZP7c.tNkAbGXqUFBd(); o2Yia7q > 0 {

			jBTzGqphMc := /*line a95_GN.go:1*/ new(big.Int).SetUint64(o2Yia7q)
			/*line Isic9Ppo77.go:1*/ jBTzGqphMc.Mul(jBTzGqphMc, azD5aZP7c.msg.BlobGasFeeCap)
			/*line iPlgvd.go:1*/ aaIqwx.Add(aaIqwx, jBTzGqphMc)

			ycTOoDK7gyI := /*line BELtw1gl.go:1*/ new(big.Int).SetUint64(o2Yia7q)
			/*line BALRP1.go:1*/ ycTOoDK7gyI.Mul(ycTOoDK7gyI, azD5aZP7c.evm.Context.BlobBaseFee)
			/*line WwybGu.go:1*/ fRvA7i.Add(fRvA7i, ycTOoDK7gyI)
		}
	}
	kvYjmBqi, reTKtI3DRV7M := /*line cqFsG_L.go:1*/ uint256.FromBig(aaIqwx)
	if reTKtI3DRV7M {
		return /*line ukp_bH.go:1*/ fmt.Errorf(func() /*line q7qqyX9.go:1*/ string {
			seed := /*line q7qqyX9.go:1*/ byte(196)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line q7qqyX9.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line q7qqyX9.go:1*/ fnc(233)(36)(11)(252)(57)(117)(234)(226)(183)(124)(248)(157)(63)(207)(72)(226)(183)(122)(248)(228)(209)(149)(41)(14)(94)(187)(129)(247)(251)(235)(216)(107)(27)(73)(125)(252)(248)(239)(237)(135)(32)(67)(135)(248)(50)(107)(225)(193)
			return /*line q7qqyX9.go:1*/ string(data)
		}(), IS8iCTU /*line SwlIMkhN3T0f.go:1*/, azD5aZP7c.msg.From.Hex())
	}
	if kIY609j, ifu6pB1YTVF := /*line IdD1snLBfp5L.go:1*/ azD5aZP7c.state.GetBalance(azD5aZP7c.msg.From), kvYjmBqi; /*line AnXzSBEAVa6.go:1*/ kIY609j.Cmp(ifu6pB1YTVF) < 0 {
		return /*line diUXDAln.go:1*/ fmt.Errorf(func() /*line dATEqsd5S5p.go:1*/ string {
			seed := /*line dATEqsd5S5p.go:1*/ byte(107)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line dATEqsd5S5p.go:1*/ append(data, x^seed); seed += x; return fnc }
			/*line dATEqsd5S5p.go:1*/ fnc(78)(206)(189)(100)(201)(21)(226)(26)(231)(26)(240)(83)(227)(223)(168)(88)(233)(7)(29)(181)(111)(207)(168)(71)(22)(227)(4)(84)(237)(195)
			return /*line dATEqsd5S5p.go:1*/ string(data)
		}(), IS8iCTU /*line hGrLdg.go:1*/, azD5aZP7c.msg.From.Hex(), kIY609j, ifu6pB1YTVF)
	}
	if beOvhSsCC3c := /*line BTptrS4kP.go:1*/ azD5aZP7c.gp.SubGas(azD5aZP7c.msg.GasLimit); beOvhSsCC3c != nil {
		return beOvhSsCC3c
	}
	azD5aZP7c.gasRemaining += azD5aZP7c.msg.GasLimit

	azD5aZP7c.initialGas = azD5aZP7c.msg.GasLimit
	cj9oL3msd, _ := /*line GwJK73Vq.go:1*/ uint256.FromBig(fRvA7i)
	/*line FcbQegcRSo.go:1*/ azD5aZP7c.state.SubBalance(azD5aZP7c.msg.From, cj9oL3msd)
	return nil
}

func (azD5aZP7c *StateTransition) frxPpELhXK() error {

	rrEywqc_se := azD5aZP7c.msg
	if !rrEywqc_se.SkipAccountChecks {

		cK_zLw9Fn := /*line vECJm3wVl_s.go:1*/ azD5aZP7c.state.GetNonce(rrEywqc_se.From)
		if m0UsArn2wT := rrEywqc_se.Nonce; cK_zLw9Fn < m0UsArn2wT {
			return /*line eAcuWnk4z.go:1*/ fmt.Errorf(func() /*line HoZ2pQz.go:1*/ string {
				data := [] /*line HoZ2pQz.go:1*/ byte("zY:\xf7\x9d*\x18re\xb3\xf3\xa0\xe5\xd8\xfdl\x8e\xe3\xaai\xf7\xffϢt\xb8\xe2\x8d:\xa9?\xc4")
				positions := [...]byte{20, 1, 15, 19, 16, 29, 31, 16, 30, 11, 0, 22, 30, 0, 23, 14, 6, 4, 15, 16, 29, 26, 0, 18, 13, 5, 19, 31, 21, 17, 29, 11, 13, 11, 25, 22, 10, 6, 3, 5, 30, 1, 9, 12, 25, 3, 27, 3}
				for i := 0; i < 48; i += 2 {
					localKey := /*line HoZ2pQz.go:1*/ byte(i) + /*line HoZ2pQz.go:1*/ byte(positions[i]^positions[i+1]) + 103
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line HoZ2pQz.go:1*/ string(data)
			}(), BTDSak,
				/*line aGAfIfS.go:1*/ rrEywqc_se.From.Hex(), m0UsArn2wT, cK_zLw9Fn)
		} else if cK_zLw9Fn > m0UsArn2wT {
			return /*line JWxoQdpBMK.go:1*/ fmt.Errorf(func() /*line mYzROEH59Oel.go:1*/ string {
				seed := /*line mYzROEH59Oel.go:1*/ byte(232)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line mYzROEH59Oel.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line mYzROEH59Oel.go:1*/ fnc(13)(108)(155)(28)(121)(245)(234)(226)(183)(124)(248)(157)(63)(207)(84)(156)(140)(28)(250)(218)(185)(177)(30)(143)(31)(43)(105)(195)(91)(156)(61)(185)
				return /*line mYzROEH59Oel.go:1*/ string(data)
			}(), PUInji920l,
				/*line DtVhZO2.go:1*/ rrEywqc_se.From.Hex(), m0UsArn2wT, cK_zLw9Fn)
		} else if cK_zLw9Fn+1 < cK_zLw9Fn {
			return /*line AyHg_Ba.go:1*/ fmt.Errorf(func() /*line pVlxAJ.go:1*/ string {
				seed := /*line pVlxAJ.go:1*/ byte(37)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line pVlxAJ.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line pVlxAJ.go:1*/ fnc(74)(230)(143)(4)(73)(149)(42)(98)(183)(124)(248)(157)(63)(207)(84)(156)(134)(13)(25)(39)(80)(117)(208)(165)(137)
				return /*line pVlxAJ.go:1*/ string(data)
			}(), DRCaKsEo,
				/*line tmoPKuna.go:1*/ rrEywqc_se.From.Hex(), cK_zLw9Fn)
		}

		f6vTM4zZ := /*line Ykchj3kQ.go:1*/ azD5aZP7c.state.GetCodeHash(rrEywqc_se.From)
		if f6vTM4zZ != (common.Hash{}) && f6vTM4zZ != types.JhrQnETMFm {
			return /*line VtqETeDOpY.go:1*/ fmt.Errorf(func() /*line LPar8tA_Y.go:1*/ string {
				data := [] /*line LPar8tA_Y.go:1*/ byte("%\x1d}e\x14\x83dtessn\aq,1/\x1cR\x81\\\x89sh:\x85ks")
				positions := [...]byte{4, 15, 12, 1, 1, 3, 21, 11, 17, 26, 25, 11, 20, 16, 11, 20, 20, 4, 19, 7, 25, 26, 1, 18, 26, 20, 1, 2, 5, 13}
				for i := 0; i < 30; i += 2 {
					localKey := /*line LPar8tA_Y.go:1*/ byte(i) + /*line LPar8tA_Y.go:1*/ byte(positions[i]^positions[i+1]) + 233
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line LPar8tA_Y.go:1*/ string(data)
			}(), QmP7XWS,
				/*line DvaCBfEz7soH.go:1*/ rrEywqc_se.From.Hex(), f6vTM4zZ)
		}
	}

	if /*line t7VWRXQWmK39.go:1*/ azD5aZP7c.evm.ChainConfig().IsLondon(azD5aZP7c.evm.Context.BlockNumber) {

		daKHNXRl := azD5aZP7c.evm.Config.NoBaseFee && /*line T0TDh7.go:1*/ rrEywqc_se.GasFeeCap.BitLen() == 0 && /*line T4sWWOZXgKg.go:1*/ rrEywqc_se.GasTipCap.BitLen() == 0
		if !daKHNXRl {
			if iFBW2xFkLp := /*line ao_6pnNH.go:1*/ rrEywqc_se.GasFeeCap.BitLen(); iFBW2xFkLp > 256 {
				return /*line TkRv8K.go:1*/ fmt.Errorf(func() /*line igW2DqG.go:1*/ string {
					key := [] /*line igW2DqG.go:1*/ byte("\xd4#qiyF,'\xb7U\xf1?!A\xb8\x03\xa1\xcez\x1bf\fPd\x9d\x16-\x89\x92N~\xf374W\xb7M\xfc\xa5\x03D\x8c\xfd")
					data := [] /*line igW2DqG.go:1*/ byte("\xf9\x9a\xab\x89ڪ\x90\x99\x1c\xc8d_F\xb7\xe4#\x0e/\xf2a\xcbq\xa0\xc9\x0f]\x8e\xfc\xb2\xb0\xe7gW\xa0\xbc%\xb4p\r=d\xb1a")
					for i, b := range key {
						data[i] = data[i] - b
					}
					return /*line igW2DqG.go:1*/ string(data)
				}(), NZGyk3sP,
					/*line fl7aaTPyr.go:1*/ rrEywqc_se.From.Hex(), iFBW2xFkLp)
			}
			if iFBW2xFkLp := /*line bO9somJf.go:1*/ rrEywqc_se.GasTipCap.BitLen(); iFBW2xFkLp > 256 {
				return /*line uPwrXbA.go:1*/ fmt.Errorf(func() /*line TKRSLGLb.go:1*/ string {
					data := /*line TKRSLGLb.go:1*/ make([]byte, 0, 52)
					i := 2
					decryptKey := 41
					for counter := 0; i != 12; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 11:
							data = /*line TKRSLGLb.go:1*/ append(data, 20)
							i = 18
						case 17:
							data = /*line TKRSLGLb.go:1*/ append(data, "\x02\f\x1e"...,
							)
							i = 8
						case 1:
							i = 11
							data = /*line TKRSLGLb.go:1*/ append(data, "C\x00\b"...,
							)
						case 9:
							for y := range data {
								data[y] = data[y] ^ /*line TKRSLGLb.go:1*/ byte(decryptKey^y)
							}
							i = 12
						case 5:
							data = /*line TKRSLGLb.go:1*/ append(data, ":7"...,
							)
							i = 7
						case 3:
							i = 19
							data = /*line TKRSLGLb.go:1*/ append(data, "%2*="...,
							)
						case 13:
							i = 3
							data = /*line TKRSLGLb.go:1*/ append(data, "d\"&"...,
							)
						case 2:
							i = 10
							data = /*line TKRSLGLb.go:1*/ append(data, "b1"...,
							)
						case 15:
							i = 9
							data = /*line TKRSLGLb.go:1*/ append(data, "S\x11"...,
							)
						case 18:
							i = 17
							data = /*line TKRSLGLb.go:1*/ append(data, "O\x02\b"...,
							)
						case 16:
							i = 14
							data = /*line TKRSLGLb.go:1*/ append(data, "?\t="...,
							)
						case 19:
							data = /*line TKRSLGLb.go:1*/ append(data, ">ln<"...,
							)
							i = 0
						case 10:
							data = /*line TKRSLGLb.go:1*/ append(data, 127)
							i = 13
						case 7:
							i = 4
							data = /*line TKRSLGLb.go:1*/ append(data, "-\x04!;"...,
							)
						case 20:
							i = 6
							data = /*line TKRSLGLb.go:1*/ append(data, "\"6"...,
							)
						case 0:
							data = /*line TKRSLGLb.go:1*/ append(data, "eh"...,
							)
							i = 5
						case 6:
							i = 16
							data = /*line TKRSLGLb.go:1*/ append(data, "*$\x1a>"...,
							)
						case 14:
							data = /*line TKRSLGLb.go:1*/ append(data, "\x15!\x04\x17"...,
							)
							i = 1
						case 8:
							i = 15
							data = /*line TKRSLGLb.go:1*/ append(data, "\x01RW"...,
							)
						case 4:
							i = 20
							data = /*line TKRSLGLb.go:1*/ append(data, 62)
						}
					}
					return /*line TKRSLGLb.go:1*/ string(data)
				}(), Ld_lTHQa,
					/*line Q29hvlO17K.go:1*/ rrEywqc_se.From.Hex(), iFBW2xFkLp)
			}
			if /*line euUpsauw.go:1*/ rrEywqc_se.GasFeeCap.Cmp(rrEywqc_se.GasTipCap) < 0 {
				return /*line QHjSOVEE8Giq.go:1*/ fmt.Errorf(func() /*line J4YiAqK.go:1*/ string {
					seed := /*line J4YiAqK.go:1*/ byte(64)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data = /*line J4YiAqK.go:1*/ append(data, x+seed); seed += x; return fnc }
					/*line J4YiAqK.go:1*/ fnc(229)(82)(195)(230)(65)(3)(0)(14)(243)(14)(0)(173)(5)(81)(182)(244)(77)(244)(23)(216)(34)(247)(6)(3)(247)(11)(5)(205)(31)(0)(235)(21)(13)(213)(26)(18)(199)(230)(5)(78)(185)(244)(77)(244)(23)(206)(31)(0)(235)(21)(13)(213)(26)(18)(199)(230)(5)(78)
					return /*line J4YiAqK.go:1*/ string(data)
				}(), ASHz0aLCD1b,
					/*line ElqRrK.go:1*/ rrEywqc_se.From.Hex(), rrEywqc_se.GasTipCap, rrEywqc_se.GasFeeCap)
			}

			if /*line qQS1Ucu2Jf.go:1*/ rrEywqc_se.GasFeeCap.Cmp(azD5aZP7c.evm.Context.BaseFee) < 0 {
				return /*line nxdCRUw68n.go:1*/ fmt.Errorf(func() /*line d3uOiF.go:1*/ string {
					data := /*line d3uOiF.go:1*/ make([]byte, 0, 46)
					i := 1
					decryptKey := 236
					for counter := 0; i != 4; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 13:
							i = 15
							data = /*line d3uOiF.go:1*/ append(data, "\xf4\xdb\x1d!"...,
							)
						case 3:
							i = 0
							data = /*line d3uOiF.go:1*/ append(data, "\x16\xde\xc5\xcb"...,
							)
						case 6:
							for y := range data {
								data[y] = data[y] - /*line d3uOiF.go:1*/ byte(decryptKey^y)
							}
							i = 4
						case 11:
							i = 10
							data = /*line d3uOiF.go:1*/ append(data, "\xe2\xd7\x15\n"...,
							)
						case 10:
							data = /*line d3uOiF.go:1*/ append(data, "\"\xf1\x11\x12"...,
							)
							i = 16
						case 7:
							data = /*line d3uOiF.go:1*/ append(data, "\xd9+"...,
							)
							i = 11
						case 1:
							i = 13
							data = /*line d3uOiF.go:1*/ append(data, "\xdd0"...,
							)
						case 16:
							data = /*line d3uOiF.go:1*/ append(data, "\xfe\x14"...,
							)
							i = 12
						case 12:
							data = /*line d3uOiF.go:1*/ append(data, "\x12\xe8\x03"...,
							)
							i = 3
						case 15:
							i = 17
							data = /*line d3uOiF.go:1*/ append(data, "\"1\x15"...,
							)
						case 0:
							data = /*line d3uOiF.go:1*/ append(data, "\x1a\xc4"...,
							)
							i = 2
						case 5:
							i = 8
							data = /*line d3uOiF.go:1*/ append(data, "\xe4\x04\xf5\xcb"...,
							)
						case 8:
							data = /*line d3uOiF.go:1*/ append(data, 178)
							i = 9
						case 9:
							data = /*line d3uOiF.go:1*/ append(data, "\xb8\a"...,
							)
							i = 6
						case 17:
							data = /*line d3uOiF.go:1*/ append(data, "$%\xd3"...,
							)
							i = 7
						case 2:
							i = 14
							data = /*line d3uOiF.go:1*/ append(data, "\xb9\xfc\xfc"...,
							)
						case 14:
							data = /*line d3uOiF.go:1*/ append(data, "\x0f\x02"...,
							)
							i = 5
						}
					}
					return /*line d3uOiF.go:1*/ string(data)
				}(), KMdaRv6na,
					/*line FBeQSp.go:1*/ rrEywqc_se.From.Hex(), rrEywqc_se.GasFeeCap, azD5aZP7c.evm.Context.BaseFee)
			}
		}
	}

	if rrEywqc_se.BlobHashes != nil {

		if rrEywqc_se.To == nil {
			return ZBGarFk1UOg
		}
		if /*line KQpcwwsu9dVz.go:1*/ len(rrEywqc_se.BlobHashes) == 0 {
			return CQgsoNkePe5
		}
		for x4YIdUciVB, n2z3WgWjM := range rrEywqc_se.BlobHashes {
			if ! /*line aCSHxUwtP.go:1*/ kzg4844.IsValidVersionedHash(n2z3WgWjM[:]) {
				return /*line F1mBx6pJH.go:1*/ fmt.Errorf(func() /*line DX70aO.go:1*/ string {
					data := /*line DX70aO.go:1*/ make([]byte, 0, 33)
					i := 11
					decryptKey := 191
					for counter := 0; i != 8; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 12:
							data = /*line DX70aO.go:1*/ append(data, "\xed\xfc\xaa"...,
							)
							i = 3
						case 10:
							for y := range data {
								data[y] = data[y] + /*line DX70aO.go:1*/ byte(decryptKey^y)
							}
							i = 8
						case 1:
							data = /*line DX70aO.go:1*/ append(data, 253)
							i = 18
						case 0:
							data = /*line DX70aO.go:1*/ append(data, "\xfe\xbb"...,
							)
							i = 7
						case 11:
							i = 6
							data = /*line DX70aO.go:1*/ append(data, 229)
						case 2:
							i = 13
							data = /*line DX70aO.go:1*/ append(data, "\t\f"...,
							)
						case 13:
							data = /*line DX70aO.go:1*/ append(data, 12)
							i = 10
						case 7:
							data = /*line DX70aO.go:1*/ append(data, "\x12\xfe\f\x12"...,
							)
							i = 2
						case 15:
							data = /*line DX70aO.go:1*/ append(data, "\xfe\x03"...,
							)
							i = 4
						case 17:
							i = 14
							data = /*line DX70aO.go:1*/ append(data, "\xb2\xff\xf9"...,
							)
						case 5:
							data = /*line DX70aO.go:1*/ append(data, 233)
							i = 9
						case 4:
							data = /*line DX70aO.go:1*/ append(data, "\xef\xff"...,
							)
							i = 1
						case 9:
							data = /*line DX70aO.go:1*/ append(data, "\xa6\xf3"...,
							)
							i = 12
						case 16:
							i = 5
							data = /*line DX70aO.go:1*/ append(data, "\xf0䧭"...,
							)
						case 6:
							data = /*line DX70aO.go:1*/ append(data, 240)
							i = 16
						case 18:
							i = 17
							data = /*line DX70aO.go:1*/ append(data, 245)
						case 3:
							data = /*line DX70aO.go:1*/ append(data, 248)
							i = 15
						case 14:
							i = 0
							data = /*line DX70aO.go:1*/ append(data, 8)
						}
					}
					return /*line DX70aO.go:1*/ string(data)
				}(), x4YIdUciVB)
			}
		}
	}

	if /*line YTIMjktuuSI.go:1*/ azD5aZP7c.evm.ChainConfig().IsCancun(azD5aZP7c.evm.Context.BlockNumber, azD5aZP7c.evm.Context.Time) {
		if /*line HlmpeKitEZOh.go:1*/ azD5aZP7c.tNkAbGXqUFBd() > 0 {

			daKHNXRl := azD5aZP7c.evm.Config.NoBaseFee && /*line opNRO3VK.go:1*/ rrEywqc_se.BlobGasFeeCap.BitLen() == 0
			if !daKHNXRl {

				if /*line xt26EvbaM2v.go:1*/ rrEywqc_se.BlobGasFeeCap.Cmp(azD5aZP7c.evm.Context.BlobBaseFee) < 0 {
					return /*line UgMiG8hLI.go:1*/ fmt.Errorf(func() /*line ARDpHjc44rQ.go:1*/ string {
						seed := /*line ARDpHjc44rQ.go:1*/ byte(146)
						var data []byte
						type decFunc func(byte) decFunc
						var fnc decFunc
						fnc = func(x byte) decFunc { data = /*line ARDpHjc44rQ.go:1*/ append(data, x+seed); seed += x; return fnc }
						/*line ARDpHjc44rQ.go:1*/ fnc(147)(82)(195)(230)(65)(3)(0)(14)(243)(14)(0)(173)(5)(81)(170)(66)(10)(3)(243)(229)(26)(18)(211)(31)(0)(222)(30)(15)(202)(230)(5)(81)(182)(244)(66)(10)(3)(243)(224)(31)(18)(242)(225)(31)(0)(213)(230)(5)(81)
						return /*line ARDpHjc44rQ.go:1*/ string(data)
					}(), DKYkxv6_M,
						/*line Cew3_m4kmwr.go:1*/ rrEywqc_se.From.Hex(), rrEywqc_se.BlobGasFeeCap, azD5aZP7c.evm.Context.BlobBaseFee)
				}
			}
		}
	}
	return /*line C4zXIm47.go:1*/ azD5aZP7c.fvZooT7()
}

func (azD5aZP7c *StateTransition) TransitionDb() (*YHD4JD4i7, error) {

	if beOvhSsCC3c := /*line nnLkmka0o.go:1*/ azD5aZP7c.frxPpELhXK(); beOvhSsCC3c != nil {
		return nil, beOvhSsCC3c
	}

	if cCIxZ1iUM := azD5aZP7c.evm.Config.Tracer; cCIxZ1iUM != nil {
		/*line oogj2VnWg8bn.go:1*/ cCIxZ1iUM.CaptureTxStart(azD5aZP7c.initialGas)
		defer func() {
			/*line yTIsUy3_av.go:1*/ cCIxZ1iUM.CaptureTxEnd(azD5aZP7c.gasRemaining)
		}()
	}

	var (
		rrEywqc_se = azD5aZP7c.msg
		iqQD9nu    = /*line HV0Et3_zWMpP.go:1*/ vm.AccountRef(rrEywqc_se.From)
		mdeiFQ     = /*line q3lcU7rXY8.go:1*/ azD5aZP7c.evm.ChainConfig().Rules(azD5aZP7c.evm.Context.BlockNumber, azD5aZP7c.evm.Context.Random != nil, azD5aZP7c.evm.Context.Time)
		gX5PsmvWX  = rrEywqc_se.To == nil
	)

	mtuuPqNr_M, beOvhSsCC3c := /*line JasYL6.go:1*/ QCDeiF8gmvU(rrEywqc_se.Data, rrEywqc_se.AccessList, gX5PsmvWX, mdeiFQ.IsHomestead, mdeiFQ.IsIstanbul, mdeiFQ.IsShanghai)
	if beOvhSsCC3c != nil {
		return nil, beOvhSsCC3c
	}
	if azD5aZP7c.gasRemaining < mtuuPqNr_M {
		return nil /*line YSP71CQ9.go:1*/, fmt.Errorf(func() /*line esYNq1.go:1*/ string {
			key := [] /*line esYNq1.go:1*/ byte("'\xb5-\xae4J\x9c{6\x98\x97:\xce\x03\x949b\xcd\t\xcc")
			data := [] /*line esYNq1.go:1*/ byte("\xfe\xc2\rr4\x17\xda\xea\xea\x8d\xcd\xf2Rt\xcd5\x12S\x1c\x98")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return /*line esYNq1.go:1*/ string(data)
		}(), HdRvVhQk, azD5aZP7c.gasRemaining, mtuuPqNr_M)
	}
	azD5aZP7c.gasRemaining -= mtuuPqNr_M

	ufc7SdmXHV7, reTKtI3DRV7M := /*line BOFLgO.go:1*/ uint256.FromBig(rrEywqc_se.Value)
	if reTKtI3DRV7M {
		return nil /*line TaQgqBmWgc.go:1*/, fmt.Errorf(func() /*line Mdgawr6GS.go:1*/ string {
			data := [] /*line Mdgawr6GS.go:1*/ byte("k9\xf3 \xe5<\xaf\xf5m\xa0\xf9\xf3\xeev")
			positions := [...]byte{12, 9, 1, 7, 0, 0, 2, 7, 5, 8, 5, 4, 11, 6, 8, 0, 10, 10}
			for i := 0; i < 18; i += 2 {
				localKey := /*line Mdgawr6GS.go:1*/ byte(i) + /*line Mdgawr6GS.go:1*/ byte(positions[i]^positions[i+1]) + 118
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return /*line Mdgawr6GS.go:1*/ string(data)
		}(), WrwQcX4l /*line oZqhIv92p4S.go:1*/, rrEywqc_se.From.Hex())
	}
	if ! /*line C9xo3sX0G3.go:1*/ ufc7SdmXHV7.IsZero() && ! /*line x3dOJU.go:1*/ azD5aZP7c.evm.Context.CanTransfer(azD5aZP7c.state, rrEywqc_se.From, ufc7SdmXHV7) {
		return nil /*line U4hzIaQamj7z.go:1*/, fmt.Errorf(func() /*line vxoH4X8.go:1*/ string {
			data := /*line vxoH4X8.go:1*/ make([]byte, 0, 15)
			i := 1
			decryptKey := 94
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 2
					data = /*line vxoH4X8.go:1*/ append(data, "\xa7\xa4"...,
					)
				case 2:
					data = /*line vxoH4X8.go:1*/ append(data, "\xb3\xb3¿"...,
					)
					i = 3
				case 0:
					i = 7
					for y := range data {
						data[y] = data[y] - /*line vxoH4X8.go:1*/ byte(decryptKey^y)
					}
				case 4:
					data = /*line vxoH4X8.go:1*/ append(data, "e\xa3"...,
					)
					i = 6
				case 1:
					i = 5
					data = /*line vxoH4X8.go:1*/ append(data, 107)
				case 3:
					i = 0
					data = /*line vxoH4X8.go:1*/ append(data, "mo\xc1"...,
					)
				case 5:
					data = /*line vxoH4X8.go:1*/ append(data, "\xbe~"...,
					)
					i = 4
				}
			}
			return /*line vxoH4X8.go:1*/ string(data)
		}(), WrwQcX4l /*line GoqicS0lahw.go:1*/, rrEywqc_se.From.Hex())
	}

	if mdeiFQ.IsShanghai && gX5PsmvWX && /*line X_02Ru.go:1*/ len(rrEywqc_se.Data) > params.MaxInitCodeSize {
		return nil /*line hKftNIfK.go:1*/, fmt.Errorf(func() /*line JvmPVyYzMUP9.go:1*/ string {
			fullData := [] /*line JvmPVyYzMUP9.go:1*/ byte("\xad\xe1\xf4\x05ѡ\x0f\x9c\x9d\b\x92\x82\xcd\"\x7f_\xf49Tw\xae\x03\xba\xe0\xa5c-\xd0\xe3\x1fɥ%\xf0\xbd!\xfa\x7f\xc5t\x19mpi\x19W>D\xfbD")
			idxKey := [] /*line JvmPVyYzMUP9.go:1*/ byte("m\x00d-\xf0\x97\xbap\x9e'd\x19P\x83m\xff")
			data := /*line JvmPVyYzMUP9.go:1*/ make([]byte, 0, 26)
			data = /*line JvmPVyYzMUP9.go:1*/ append(data, fullData[122^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[0])]^fullData[75^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[0])], fullData[37^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[1])]-fullData[9^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[1])], fullData[170^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[6])]-fullData[172^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[6])], fullData[182^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[6])]-fullData[186^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[6])], fullData[228^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[15])]-fullData[214^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[15])], fullData[74^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[0])]+fullData[93^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[0])], fullData[160^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[13])]-fullData[161^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[13])], fullData[87^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])]+fullData[78^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])], fullData[94^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])]+fullData[85^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])], fullData[81^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])]+fullData[90^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])], fullData[45^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[1])]^fullData[46^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[1])], fullData[77^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])]-fullData[79^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[12])], fullData[152^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[5])]-fullData[179^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[5])], fullData[208^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[4])]-fullData[243^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[4])], fullData[12^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[9])]-fullData[22^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[9])], fullData[140^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[8])]+fullData[147^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[8])], fullData[12^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[11])]-fullData[5^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[11])], fullData[116^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[14])]^fullData[107^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[14])], fullData[216^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[4])]^fullData[218^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[4])], fullData[76^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[14])]^fullData[101^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[14])], fullData[106^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[7])]^fullData[95^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[7])], fullData[135^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[13])]^fullData[155^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[13])], fullData[175^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[13])]^fullData[146^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[13])], fullData[227^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[4])]+fullData[228^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[4])], fullData[37^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[9])]^fullData[44^ /*line JvmPVyYzMUP9.go:1*/ int(idxKey[9])])
			return /*line JvmPVyYzMUP9.go:1*/ string(data)
		}(), FQA1MFfYa1Ac /*line tXBlABaqTcy.go:1*/, len(rrEywqc_se.Data), params.MaxInitCodeSize)
	}

	/*line n6uSqI.go:1*/
	azD5aZP7c.state.Prepare(mdeiFQ, rrEywqc_se.From, azD5aZP7c.evm.Context.Coinbase, rrEywqc_se.To /*line oUoQVsaN.go:1*/, vm.WNX2wLgB2IK(mdeiFQ), rrEywqc_se.AccessList)

	var (
		wwOxMoKfe1 []byte
		sRxmbmVnY  error
	)
	if gX5PsmvWX {
		wwOxMoKfe1, _, azD5aZP7c.gasRemaining, sRxmbmVnY = /*line yNAZW5xElHQa.go:1*/ azD5aZP7c.evm.Create(iqQD9nu, rrEywqc_se.Data, azD5aZP7c.gasRemaining, ufc7SdmXHV7)
	} else {

		/*line eE2jzGlN.go:1*/
		azD5aZP7c.state.SetNonce(rrEywqc_se.From /*line OtveilYIo8Ly.go:1*/, azD5aZP7c.state.GetNonce( /*line SozsMwrNbhJ8.go:1*/ iqQD9nu.Address())+1)
		wwOxMoKfe1, azD5aZP7c.gasRemaining, sRxmbmVnY = /*line Gtta2ccxiaaI.go:1*/ azD5aZP7c.evm.Call(iqQD9nu /*line DXXCqI_2A.go:1*/, azD5aZP7c.rJ6LH9n(), rrEywqc_se.Data, azD5aZP7c.gasRemaining, ufc7SdmXHV7)
	}

	var kqpHiQ uint64
	if !mdeiFQ.IsLondon {

		kqpHiQ = /*line UkqrI7.go:1*/ azD5aZP7c.bX_8tQukI(params.RefundQuotient)
	} else {

		kqpHiQ = /*line qhdaNNV.go:1*/ azD5aZP7c.bX_8tQukI(params.RefundQuotientEIP3529)
	}
	axW3q7 := rrEywqc_se.GasPrice
	if mdeiFQ.IsLondon {
		axW3q7 = /*line VdsZ7i1.go:1*/ cmath.BigMin(rrEywqc_se.GasTipCap /*line aX0ybXk8mObA.go:1*/, new(big.Int).Sub(rrEywqc_se.GasFeeCap, azD5aZP7c.evm.Context.BaseFee))
	}
	jh7jCW, _ := /*line osxvCq.go:1*/ uint256.FromBig(axW3q7)

	if azD5aZP7c.evm.Config.NoBaseFee && /*line BjNuCi8IXDX.go:1*/ rrEywqc_se.GasFeeCap.Sign() == 0 && /*line da0772Is.go:1*/ rrEywqc_se.GasTipCap.Sign() == 0 {

	} else {
		uq07oir := /*line GtXr_ulBL.go:1*/ new(uint256.Int).SetUint64( /*line H9TKMlAki.go:1*/ azD5aZP7c.yA6Rl84cKd3())
		/*line s1Mnv_.go:1*/ uq07oir.Mul(uq07oir, jh7jCW)
		/*line WcNhOhkhPE4.go:1*/ azD5aZP7c.state.AddBalance(azD5aZP7c.evm.Context.Coinbase, uq07oir)
	}

	return &YHD4JD4i7{
		A7_zDWIb:/*line u77pEKdJl.go:1*/ azD5aZP7c.yA6Rl84cKd3(),
		ZY9Ly6apaW: kqpHiQ,
		BahScRGbl:  sRxmbmVnY,
		A3aOS4Y4:   wwOxMoKfe1,
	}, nil
}

func (azD5aZP7c *StateTransition) bX_8tQukI(mW36BhfB uint64) uint64 {

	gorOlsNidK := /*line EoNLRwbMe.go:1*/ azD5aZP7c.yA6Rl84cKd3() / mW36BhfB
	if gorOlsNidK > /*line lApNDFq7cGM.go:1*/ azD5aZP7c.state.GetRefund() {
		gorOlsNidK = /*line BdthZiAJpHP.go:1*/ azD5aZP7c.state.GetRefund()
	}
	azD5aZP7c.gasRemaining += gorOlsNidK

	cItcF3_b := /*line pQOzx5JQtSqe.go:1*/ uint256.NewInt(azD5aZP7c.gasRemaining)
	cItcF3_b = /*line rk0k5LDoLwZ.go:1*/ cItcF3_b.Mul(cItcF3_b /*line KOrOQwMHSGEF.go:1*/, uint256.MustFromBig(azD5aZP7c.msg.GasPrice))
	/*line CBBTV_J.go:1*/ azD5aZP7c.state.AddBalance(azD5aZP7c.msg.From, cItcF3_b)

	/*line _3ytHx.go:1*/
	azD5aZP7c.gp.AddGas(azD5aZP7c.gasRemaining)

	return gorOlsNidK
}

func (azD5aZP7c *StateTransition) yA6Rl84cKd3() uint64 {
	return azD5aZP7c.initialGas - azD5aZP7c.gasRemaining
}

func (azD5aZP7c *StateTransition) tNkAbGXqUFBd() uint64 {
	return /*line __3ineu.go:1*/ uint64( /*line pDFdWLSNWcg.go:1*/ len(azD5aZP7c.msg.BlobHashes) * params.BlobTxBlobGasPerBlob)
}

var _ = fmt.Append
var _ = math.Abs

const _ = big.Above

var _ types.AccessList

const _ = vm.ADD

var _ = common.AbsolutePath
var _ = cmath.BigMax
var _ kzg4844.Blob
var _ = params.AllCliqueProtocolChanges
var _ = uint256.ErrBadBufferLength
