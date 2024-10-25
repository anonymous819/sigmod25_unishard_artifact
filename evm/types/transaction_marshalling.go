//line :1
package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

type txJSON struct {
	Type hexutil.Uint64 `json:"type"`

	ChainID              *hexutil.Big    `json:"chainId,omitempty"`
	Nonce                *hexutil.Uint64 `json:"nonce"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxFeePerBlobGas     *hexutil.Big    `json:"maxFeePerBlobGas,omitempty"`
	Value                *hexutil.Big    `json:"value"`
	Input                *hexutil.Bytes  `json:"input"`
	AccessList           *AccessList     `json:"accessList,omitempty"`
	BlobVersionedHashes  []common.Hash   `json:"blobVersionedHashes,omitempty"`
	V                    *hexutil.Big    `json:"v"`
	R                    *hexutil.Big    `json:"r"`
	S                    *hexutil.Big    `json:"s"`
	YParity              *hexutil.Uint64 `json:"yParity,omitempty"`

	Blobs       []kzg4844.Blob       `json:"blobs,omitempty"`
	Commitments []kzg4844.Commitment `json:"commitments,omitempty"`
	Proofs      []kzg4844.Proof      `json:"proofs,omitempty"`

	Hash common.Hash `json:"hash"`
}

func (iPfjW8i91hC *txJSON) jDvTUb() (*big.Int, error) {
	if iPfjW8i91hC.YParity != nil {
		wgXRO0hhah := /*line oIuxlvIvzKBa.go:1*/ uint64(*iPfjW8i91hC.YParity)
		if wgXRO0hhah != 0 && wgXRO0hhah != 1 {
			return nil, vnbtviuT
		}
		qcB_ur5ib053 := /*line DwlirF41qp.go:1*/ new(big.Int).SetUint64(wgXRO0hhah)
		if iPfjW8i91hC.V != nil && /*line FCuCGhl__.go:1*/ iPfjW8i91hC.V.ToInt().Cmp(qcB_ur5ib053) != 0 {
			return nil, sn89ktn_dI
		}
		return qcB_ur5ib053, nil
	}
	if iPfjW8i91hC.V != nil {
		return /*line dzXk308K5Y.go:1*/ iPfjW8i91hC.V.ToInt(), nil
	}
	return nil, tH2RXxQ
}

func (iPfjW8i91hC *Transaction) MarshalJSON() ([]byte, error) {
	var aINiWZ_Jtlzj txJSON

	aINiWZ_Jtlzj.Hash = /*line EQZvwfzLC.go:1*/ iPfjW8i91hC.Hash()
	aINiWZ_Jtlzj.Type = /*line NDvAZrg.go:1*/ hexutil.Uint64( /*line _Y95Eogh.go:1*/ iPfjW8i91hC.Type())

	switch r5Uxg5_sE2 := iPfjW8i91hC.inner.(type) {
	case *LegacyTx:
		aINiWZ_Jtlzj.Nonce = (* /*line G69WPYaNI9p2.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Nonce)
		aINiWZ_Jtlzj.To = /*line Na54WyvW.go:1*/ iPfjW8i91hC.To()
		aINiWZ_Jtlzj.Gas = (* /*line uFO5Por.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Gas)
		aINiWZ_Jtlzj.GasPrice = (* /*line W56dBC.go:1*/ hexutil.Big)(r5Uxg5_sE2.GasPrice)
		aINiWZ_Jtlzj.Value = (* /*line EK1fcz3joBV7.go:1*/ hexutil.Big)(r5Uxg5_sE2.Value)
		aINiWZ_Jtlzj.Input = (* /*line DmdQkl.go:1*/ hexutil.Bytes)(&r5Uxg5_sE2.Data)
		aINiWZ_Jtlzj.V = (* /*line qRGHsLLV.go:1*/ hexutil.Big)(r5Uxg5_sE2.V)
		aINiWZ_Jtlzj.R = (* /*line MVd_ZDz.go:1*/ hexutil.Big)(r5Uxg5_sE2.R)
		aINiWZ_Jtlzj.S = (* /*line iKNH9_LT.go:1*/ hexutil.Big)(r5Uxg5_sE2.S)
		if /*line bL51atNEl.go:1*/ iPfjW8i91hC.Protected() {
			aINiWZ_Jtlzj.ChainID = (* /*line O6ot6HWC.go:1*/ hexutil.Big)( /*line JcLAtE58y.go:1*/ iPfjW8i91hC.ChainId())
		}

	case *AccessListTx:
		aINiWZ_Jtlzj.ChainID = (* /*line qrumiAR.go:1*/ hexutil.Big)(r5Uxg5_sE2.ChainID)
		aINiWZ_Jtlzj.Nonce = (* /*line aCdlDYCoYuf.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Nonce)
		aINiWZ_Jtlzj.To = /*line H0BNqkm.go:1*/ iPfjW8i91hC.To()
		aINiWZ_Jtlzj.Gas = (* /*line Xnm_BaXOxu.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Gas)
		aINiWZ_Jtlzj.GasPrice = (* /*line vYIB4kBsi_dj.go:1*/ hexutil.Big)(r5Uxg5_sE2.GasPrice)
		aINiWZ_Jtlzj.Value = (* /*line KkN249H4mrNh.go:1*/ hexutil.Big)(r5Uxg5_sE2.Value)
		aINiWZ_Jtlzj.Input = (* /*line iul1HIAfAd.go:1*/ hexutil.Bytes)(&r5Uxg5_sE2.Data)
		aINiWZ_Jtlzj.AccessList = &r5Uxg5_sE2.AccessList
		aINiWZ_Jtlzj.V = (* /*line OLA4bh7YJ.go:1*/ hexutil.Big)(r5Uxg5_sE2.V)
		aINiWZ_Jtlzj.R = (* /*line sn5Lzb0H.go:1*/ hexutil.Big)(r5Uxg5_sE2.R)
		aINiWZ_Jtlzj.S = (* /*line xBPe6FH.go:1*/ hexutil.Big)(r5Uxg5_sE2.S)
		bYmLTIZW := /*line CsCRCa0Blt.go:1*/ r5Uxg5_sE2.V.Uint64()
		aINiWZ_Jtlzj.YParity = (* /*line e0WIUrvx.go:1*/ hexutil.Uint64)(&bYmLTIZW)

	case *DynamicFeeTx:
		aINiWZ_Jtlzj.ChainID = (* /*line HFP0MO1qm0.go:1*/ hexutil.Big)(r5Uxg5_sE2.ChainID)
		aINiWZ_Jtlzj.Nonce = (* /*line GSgDbBaBTcMV.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Nonce)
		aINiWZ_Jtlzj.To = /*line l6di5A6Bz5M.go:1*/ iPfjW8i91hC.To()
		aINiWZ_Jtlzj.Gas = (* /*line UGns8p.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Gas)
		aINiWZ_Jtlzj.MaxFeePerGas = (* /*line zoH7DQdHz.go:1*/ hexutil.Big)(r5Uxg5_sE2.GasFeeCap)
		aINiWZ_Jtlzj.MaxPriorityFeePerGas = (* /*line yoTgwg_z.go:1*/ hexutil.Big)(r5Uxg5_sE2.GasTipCap)
		aINiWZ_Jtlzj.Value = (* /*line JXxZAtSfZau.go:1*/ hexutil.Big)(r5Uxg5_sE2.Value)
		aINiWZ_Jtlzj.Input = (* /*line khULanHnXE.go:1*/ hexutil.Bytes)(&r5Uxg5_sE2.Data)
		aINiWZ_Jtlzj.AccessList = &r5Uxg5_sE2.AccessList
		aINiWZ_Jtlzj.V = (* /*line urS9dlNA9q.go:1*/ hexutil.Big)(r5Uxg5_sE2.V)
		aINiWZ_Jtlzj.R = (* /*line K4aKuf7tFEzd.go:1*/ hexutil.Big)(r5Uxg5_sE2.R)
		aINiWZ_Jtlzj.S = (* /*line rXBl_jQkZ.go:1*/ hexutil.Big)(r5Uxg5_sE2.S)
		bYmLTIZW := /*line GJiwmXxqa.go:1*/ r5Uxg5_sE2.V.Uint64()
		aINiWZ_Jtlzj.YParity = (* /*line sHDAp64.go:1*/ hexutil.Uint64)(&bYmLTIZW)

	case *BlobTx:
		aINiWZ_Jtlzj.ChainID = (* /*line u2TsdG.go:1*/ hexutil.Big)( /*line H8OFXx.go:1*/ r5Uxg5_sE2.ChainID.ToBig())
		aINiWZ_Jtlzj.Nonce = (* /*line CqrNFSj43fMa.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Nonce)
		aINiWZ_Jtlzj.Gas = (* /*line zEu5wM01xCE5.go:1*/ hexutil.Uint64)(&r5Uxg5_sE2.Gas)
		aINiWZ_Jtlzj.MaxFeePerGas = (* /*line k4zS9UDL8.go:1*/ hexutil.Big)( /*line uoMyIbwFa.go:1*/ r5Uxg5_sE2.GasFeeCap.ToBig())
		aINiWZ_Jtlzj.MaxPriorityFeePerGas = (* /*line SD5OGyL9Wy1.go:1*/ hexutil.Big)( /*line LyPp4GY.go:1*/ r5Uxg5_sE2.GasTipCap.ToBig())
		aINiWZ_Jtlzj.MaxFeePerBlobGas = (* /*line CynRzvaLY.go:1*/ hexutil.Big)( /*line hD7A7Z0OQ6c.go:1*/ r5Uxg5_sE2.BlobFeeCap.ToBig())
		aINiWZ_Jtlzj.Value = (* /*line Ep13ls2kw2.go:1*/ hexutil.Big)( /*line VAIk7lDxkYED.go:1*/ r5Uxg5_sE2.Value.ToBig())
		aINiWZ_Jtlzj.Input = (* /*line LCwyMwh.go:1*/ hexutil.Bytes)(&r5Uxg5_sE2.Data)
		aINiWZ_Jtlzj.AccessList = &r5Uxg5_sE2.AccessList
		aINiWZ_Jtlzj.BlobVersionedHashes = r5Uxg5_sE2.BlobHashes
		aINiWZ_Jtlzj.To = /*line kauh7mhWIEqc.go:1*/ iPfjW8i91hC.To()
		aINiWZ_Jtlzj.V = (* /*line GaTXu8NR.go:1*/ hexutil.Big)( /*line qkQAA_G9p703.go:1*/ r5Uxg5_sE2.V.ToBig())
		aINiWZ_Jtlzj.R = (* /*line uLIPqC0UXge.go:1*/ hexutil.Big)( /*line j0S55aa.go:1*/ r5Uxg5_sE2.R.ToBig())
		aINiWZ_Jtlzj.S = (* /*line lH0ZguGSzd.go:1*/ hexutil.Big)( /*line _kAk5gA5rS9Z.go:1*/ r5Uxg5_sE2.S.ToBig())
		bYmLTIZW := /*line E8y5rF.go:1*/ r5Uxg5_sE2.V.Uint64()
		aINiWZ_Jtlzj.YParity = (* /*line iaxN200d.go:1*/ hexutil.Uint64)(&bYmLTIZW)
		if vymDa9n := r5Uxg5_sE2.Sidecar; vymDa9n != nil {
			aINiWZ_Jtlzj.Blobs = r5Uxg5_sE2.Sidecar.Blobs
			aINiWZ_Jtlzj.Commitments = r5Uxg5_sE2.Sidecar.Commitments
			aINiWZ_Jtlzj.Proofs = r5Uxg5_sE2.Sidecar.Proofs
		}
	}
	return /*line BBBnXY5vPa.go:1*/ json.Marshal(&aINiWZ_Jtlzj)
}

func (iPfjW8i91hC *Transaction) UnmarshalJSON(uzD2Up []byte) error {
	var otqLrvn1CD txJSON
	rfteMJju := /*line xxpYrPTBEx.go:1*/ json.Unmarshal(uzD2Up, &otqLrvn1CD)
	if rfteMJju != nil {
		return rfteMJju
	}

	var vv4w9F TxData
	switch otqLrvn1CD.Type {
	case LegacyTxType:
		var r5Uxg5_sE2 LegacyTx
		vv4w9F = &r5Uxg5_sE2
		if otqLrvn1CD.Nonce == nil {
			return /*line fMtrz0gJzitu.go:1*/ errors.New(func() /*line otXd6lY6J.go:1*/ string {
				data := /*line otXd6lY6J.go:1*/ make([]byte, 0, 46)
				i := 4
				decryptKey := 225
				for counter := 0; i != 17; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 11:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xb4\xb2"...,
						)
						i = 7
					case 15:
						i = 21
						data = /*line otXd6lY6J.go:1*/ append(data, 185)
					case 5:
						data = /*line otXd6lY6J.go:1*/ append(data, 218)
						i = 16
					case 4:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xdd\xd8\xe1"...,
						)
						i = 2
					case 2:
						i = 9
						data = /*line otXd6lY6J.go:1*/ append(data, 224)
					case 7:
						for y := range data {
							data[y] = data[y] + /*line otXd6lY6J.go:1*/ byte(decryptKey^y)
						}
						i = 17
					case 10:
						i = 8
						data = /*line otXd6lY6J.go:1*/ append(data, "\xe7\xe2\xe8"...,
						)
					case 22:
						i = 14
						data = /*line otXd6lY6J.go:1*/ append(data, "\xd7\xda\xcd"...,
						)
					case 0:
						data = /*line otXd6lY6J.go:1*/ append(data, "ٚ\x92"...,
						)
						i = 5
					case 12:
						i = 10
						data = /*line otXd6lY6J.go:1*/ append(data, "\xa0\xe5"...,
						)
					case 16:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xbeo\xc2"...,
						)
						i = 3
					case 21:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xbd\xaa"...,
						)
						i = 6
					case 13:
						i = 12
						data = /*line otXd6lY6J.go:1*/ append(data, 197)
					case 1:
						data = /*line otXd6lY6J.go:1*/ append(data, 173)
						i = 15
					case 19:
						data = /*line otXd6lY6J.go:1*/ append(data, 216)
						i = 0
					case 8:
						i = 18
						data = /*line otXd6lY6J.go:1*/ append(data, "ߚ"...,
						)
					case 3:
						data = /*line otXd6lY6J.go:1*/ append(data, 191)
						i = 1
					case 20:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xda\xcc"...,
						)
						i = 22
					case 6:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xab\xbb\xaf"...,
						)
						i = 11
					case 18:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xa0\xe6\xe6\xe4"...,
						)
						i = 19
					case 9:
						data = /*line otXd6lY6J.go:1*/ append(data, "\xd5\xd9щ"...,
						)
						i = 20
					case 14:
						i = 13
						data = /*line otXd6lY6J.go:1*/ append(data, "\xd5\xc7"...,
						)
					}
				}
				return /*line otXd6lY6J.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Nonce = /*line dWZ3PCc.go:1*/ uint64(*otqLrvn1CD.Nonce)
		if otqLrvn1CD.To != nil {
			r5Uxg5_sE2.To = otqLrvn1CD.To
		}
		if otqLrvn1CD.Gas == nil {
			return /*line GiZuXS2oE.go:1*/ errors.New(func() /*line X9eoZWpS32My.go:1*/ string {
				seed := /*line X9eoZWpS32My.go:1*/ byte(12)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line X9eoZWpS32My.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line X9eoZWpS32My.go:1*/ fnc(121)(238)(230)(204)(142)(33)(59)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(174)(86)(190)(48)(89)(251)(251)(168)(164)(70)(123)(3)(11)(4)(10)(37)(63)(132)(7)
				return /*line X9eoZWpS32My.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Gas = /*line M8XgZqluvYiF.go:1*/ uint64(*otqLrvn1CD.Gas)
		if otqLrvn1CD.GasPrice == nil {
			return /*line Dao4jZL7.go:1*/ errors.New(func() /*line GTchaq0Brqjp.go:1*/ string {
				data := /*line GTchaq0Brqjp.go:1*/ make([]byte, 0, 49)
				i := 22
				decryptKey := 184
				for counter := 0; i != 6; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 8:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 255)
						i = 14
					case 15:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "\x02\xfaBF"...,
						)
						i = 8
					case 18:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "X\v"...,
						)
						i = 17
					case 12:
						i = 7
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 104)
					case 14:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "RO="...,
						)
						i = 13
					case 3:
						i = 9
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "\x1ce"...,
						)
					case 13:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "AE23"...,
						)
						i = 0
					case 2:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 89)
						i = 19
					case 5:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 74)
						i = 21
					case 22:
						i = 16
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "hcl"...,
						)
					case 0:
						i = 4
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 75)
					case 23:
						i = 15
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 73)
					case 19:
						i = 23
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "OH"...,
						)
					case 4:
						i = 11
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "?D"...,
						)
					case 1:
						for y := range data {
							data[y] = data[y] - /*line GTchaq0Brqjp.go:1*/ byte(decryptKey^y)
						}
						i = 6
					case 9:
						i = 12
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "Wbe`"...,
						)
					case 20:
						i = 10
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 82)
					case 11:
						i = 1
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 66)
					case 17:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "PRM["...,
						)
						i = 20
					case 16:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "khld"...,
						)
						i = 3
					case 7:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, 90)
						i = 18
					case 21:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "CT0"...,
						)
						i = 2
					case 10:
						data = /*line GTchaq0Brqjp.go:1*/ append(data, "\r\x13"...,
						)
						i = 5
					}
				}
				return /*line GTchaq0Brqjp.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.GasPrice = (* /*line GW2l2GQh0G.go:1*/ big.Int)(otqLrvn1CD.GasPrice)
		if otqLrvn1CD.Value == nil {
			return /*line spZN2hTevMA.go:1*/ errors.New(func() /*line GWbMd9yFjH.go:1*/ string {
				key := [] /*line GWbMd9yFjH.go:1*/ byte("\xeao*\x1d\x13\x93mҏ\xb1\xcd\x17\x8c \xfbu\xbb\x8d?\xc0Z\x04!\xac\xd2ӡ\x9c\x9a\xa1\xb4V\x99\xbeW/\x13:\xd7\xc3-\x19\xc1\xaau")
				data := [] /*line GWbMd9yFjH.go:1*/ byte("\x87\x06Ynz\xfd\n\xf2\xfdԼb\xe5R\x9e\x11\x9b\xebV\xa56`\x01\x8b\xa4\xb2\xcd\xe9\xff\x86\x94?\xf7\x9e#]rT\xa4\xa2Nm\xa8\xc5\x1b")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line GWbMd9yFjH.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Value = (* /*line CBT9MFceyqk.go:1*/ big.Int)(otqLrvn1CD.Value)
		if otqLrvn1CD.Input == nil {
			return /*line occfOynudEh.go:1*/ errors.New(func() /*line y5XV1wjDi0.go:1*/ string {
				key := [] /*line y5XV1wjDi0.go:1*/ byte("\x94B\x84\xbdǝ\"0\xe1\xf5\x16O`\xcc\xc4T\xac\xfe\xf1\xf3\xe6$\x00hk@'O\xe4\x93\xd1`k(\nra\x8bs\xbeI\x04\xc3\xf6\x8c")
				data := [] /*line y5XV1wjDi0.go:1*/ byte("\x01\xab\xf700\v\x89PSZ\x87\xc4\xc9>)\xb8\xccdZXR\x88 \x8fԮ\x97\xc4X\xba\xf1\xc9\xd9H~\xe4\xc2\xf9\xe6\x1f\xacx,e\xfa")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return /*line y5XV1wjDi0.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Data = *otqLrvn1CD.Input

		if otqLrvn1CD.R == nil {
			return /*line h6yArW34.go:1*/ errors.New(func() /*line d3V6I0KuZ8.go:1*/ string {
				seed := /*line d3V6I0KuZ8.go:1*/ byte(209)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line d3V6I0KuZ8.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line d3V6I0KuZ8.go:1*/ fnc(188)(228)(2)(0)(26)(227)(23)(167)(92)(239)(8)(244)(28)(227)(17)(225)(70)(202)(31)(240)(233)(10)(88)(247)(181)(91)(247)(167)(27)(176)(52)(6)(27)(251)(227)(18)(230)(31)(227)(2)(1)
				return /*line d3V6I0KuZ8.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.R = (* /*line FPgtXua6.go:1*/ big.Int)(otqLrvn1CD.R)

		if otqLrvn1CD.S == nil {
			return /*line tcgUGEe6Zqc.go:1*/ errors.New(func() /*line EXNtFAQ2A53.go:1*/ string {
				data := [] /*line EXNtFAQ2A53.go:1*/ byte("m\tQs\x05`\x165\x90^4u|:ed f1bG\fX\x98s7xz]'5\xedanpa]yi!7")
				positions := [...]byte{4, 13, 23, 31, 13, 22, 36, 28, 39, 26, 9, 40, 31, 30, 22, 29, 18, 9, 34, 27, 25, 6, 1, 22, 9, 6, 8, 39, 10, 2, 21, 7, 20, 21, 2, 19, 28, 13, 28, 39, 29, 37, 29, 30, 4, 31, 12, 40, 5, 21}
				for i := 0; i < 50; i += 2 {
					localKey := /*line EXNtFAQ2A53.go:1*/ byte(i) + /*line EXNtFAQ2A53.go:1*/ byte(positions[i]^positions[i+1]) + 188
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line EXNtFAQ2A53.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.S = (* /*line DI6jfz.go:1*/ big.Int)(otqLrvn1CD.S)

		if otqLrvn1CD.V == nil {
			return /*line FxM14URppZQ_.go:1*/ errors.New(func() /*line Q5X5SNH1XS.go:1*/ string {
				data := /*line Q5X5SNH1XS.go:1*/ make([]byte, 0, 42)
				i := 10
				decryptKey := 9
				for counter := 0; i != 19; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 1:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x82\x7f\x95\xa1"...,
						)
						i = 8
					case 18:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, 42)
						i = 12
					case 17:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x8e\x80\x8b"...,
						)
						i = 3
					case 3:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, 142)
						i = 16
					case 13:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "~/"...,
						)
						i = 1
					case 4:
						i = 19
						for y := range data {
							data[y] = data[y] + /*line Q5X5SNH1XS.go:1*/ byte(decryptKey^y)
						}
					case 8:
						i = 5
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\xa5\x92\x9b\xab"...,
						)
					case 12:
						i = 13
						data = /*line Q5X5SNH1XS.go:1*/ append(data, 114)
					case 7:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "k&,\x82"...,
						)
						i = 14
					case 10:
						i = 15
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x81|\x85"...,
						)
					case 9:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "kft"...,
						)
						i = 7
					case 14:
						i = 18
						data = /*line Q5X5SNH1XS.go:1*/ append(data, 50)
					case 11:
						i = 0
						data = /*line Q5X5SNH1XS.go:1*/ append(data, 133)
					case 5:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x9f\xa4\xaa"...,
						)
						i = 4
					case 16:
						i = 6
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x89\x91"...,
						)
					case 0:
						i = 17
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "}5"...,
						)
					case 6:
						i = 2
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x83\x81"...,
						)
					case 2:
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "$i"...,
						)
						i = 9
					case 15:
						i = 11
						data = /*line Q5X5SNH1XS.go:1*/ append(data, "\x84\x81"...,
						)
					}
				}
				return /*line Q5X5SNH1XS.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.V = (* /*line v7rKTgjIZ2.go:1*/ big.Int)(otqLrvn1CD.V)
		if /*line v4yXEVBXSCV.go:1*/ r5Uxg5_sE2.V.Sign() != 0 || /*line DxsZrKAD.go:1*/ r5Uxg5_sE2.R.Sign() != 0 || /*line TL_T4Q.go:1*/ r5Uxg5_sE2.S.Sign() != 0 {
			if rfteMJju := /*line HpnRLmdFQz2y.go:1*/ vwXAfyu(r5Uxg5_sE2.V, r5Uxg5_sE2.R, r5Uxg5_sE2.S, true); rfteMJju != nil {
				return rfteMJju
			}
		}

	case AccessListTxType:
		var r5Uxg5_sE2 AccessListTx
		vv4w9F = &r5Uxg5_sE2
		if otqLrvn1CD.ChainID == nil {
			return /*line XazxFy8.go:1*/ errors.New(func() /*line jaJgCQe.go:1*/ string {
				key := [] /*line jaJgCQe.go:1*/ byte("\xb9il\xbc:\b>\x88\x14Tډ\xd68\xdd\x12\x16\xa5\x9fI(\xd9\x04\xba'\xa2\xf3\xa5i\\\x80c\x8d\x1ap\xf4\xb8\xf7L\u0083f\x96fG\x8d\xe1")
				data := [] /*line jaJgCQe.go:1*/ byte("\xd4\x00\x1f\xcfSfY\xa8f1\xab\xfc\xbfJ\xb8v6\xc3\xf6,D\xbd$\x9dDʒ\xcc\a\x15\xe4D\xads\x1e\xd4̅-\xac\xf0\a\xf5\x12.\xe2\x8f")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line jaJgCQe.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.ChainID = (* /*line NWsw_n9OJ2g.go:1*/ big.Int)(otqLrvn1CD.ChainID)
		if otqLrvn1CD.Nonce == nil {
			return /*line eZJX95.go:1*/ errors.New(func() /*line DysVdyxv6V.go:1*/ string {
				data := /*line DysVdyxv6V.go:1*/ make([]byte, 0, 46)
				i := 3
				decryptKey := 86
				for counter := 0; i != 10; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 4:
						i = 15
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xe3\xea\xea"...,
						)
					case 19:
						data = /*line DysVdyxv6V.go:1*/ append(data, "fn\xb6"...,
						)
						i = 17
					case 8:
						i = 6
						data = /*line DysVdyxv6V.go:1*/ append(data, "ƽý"...,
						)
					case 11:
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xe3\xe9"...,
						)
						i = 9
					case 3:
						i = 13
						data = /*line DysVdyxv6V.go:1*/ append(data, 189)
					case 1:
						i = 2
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xc3\xc3`"...,
						)
					case 7:
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xd0\xc5\xcf"...,
						)
						i = 1
					case 9:
						i = 16
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xd8\xdb"...,
						)
					case 13:
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xba\xc5"...,
						)
						i = 8
					case 18:
						i = 19
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xab\xa8\xb0\xa9"...,
						)
					case 0:
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xe5\xd5"...,
						)
						i = 11
					case 2:
						data = /*line DysVdyxv6V.go:1*/ append(data, 167)
						i = 18
					case 15:
						for y := range data {
							data[y] = data[y] - /*line DysVdyxv6V.go:1*/ byte(decryptKey^y)
						}
						i = 10
					case 6:
						i = 7
						data = /*line DysVdyxv6V.go:1*/ append(data, "wʾ\xcb"...,
						)
					case 17:
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xb8\xb8"...,
						)
						i = 14
					case 16:
						data = /*line DysVdyxv6V.go:1*/ append(data, 237)
						i = 4
					case 12:
						data = /*line DysVdyxv6V.go:1*/ append(data, 184)
						i = 5
					case 14:
						data = /*line DysVdyxv6V.go:1*/ append(data, "\xae\xb1tn"...,
						)
						i = 12
					case 5:
						i = 0
						data = /*line DysVdyxv6V.go:1*/ append(data, "ޑ\xe6"...,
						)
					}
				}
				return /*line DysVdyxv6V.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Nonce = /*line IgbBs9NE0.go:1*/ uint64(*otqLrvn1CD.Nonce)
		if otqLrvn1CD.To != nil {
			r5Uxg5_sE2.To = otqLrvn1CD.To
		}
		if otqLrvn1CD.Gas == nil {
			return /*line IB5h2Zz.go:1*/ errors.New(func() /*line p1Prg87eMHTv.go:1*/ string {
				data := [] /*line p1Prg87eMHTv.go:1*/ byte("\x1a(\x7fsiSg\xf4'(?\x03,\xae\xef&m<,1<d\x85Sg\x12s(\xe6in\xa0k@e\xe2sLc=i+n")
				positions := [...]byte{35, 10, 35, 19, 13, 2, 27, 32, 25, 20, 8, 25, 22, 19, 13, 25, 41, 15, 13, 14, 18, 12, 0, 14, 19, 33, 23, 33, 1, 8, 31, 27, 7, 27, 25, 9, 11, 1, 10, 39, 17, 39, 10, 32, 39, 5, 37, 1, 41, 2, 28, 31, 16, 41, 17, 13, 34, 25}
				for i := 0; i < 58; i += 2 {
					localKey := /*line p1Prg87eMHTv.go:1*/ byte(i) + /*line p1Prg87eMHTv.go:1*/ byte(positions[i]^positions[i+1]) + 145
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line p1Prg87eMHTv.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Gas = /*line dJYBJ5VZ_fi.go:1*/ uint64(*otqLrvn1CD.Gas)
		if otqLrvn1CD.GasPrice == nil {
			return /*line k3EYlMESB.go:1*/ errors.New(func() /*line JtDYoEjiME.go:1*/ string {
				fullData := [] /*line JtDYoEjiME.go:1*/ byte("oipo\x14X˻\x89\xee\x16\x9dJ\xe9\x90\xc6\x12^X\xa3\xa5X*\r\xe5\x117\xbbS\xa3y\xc8d\xc5\xc4;\xe2\xa3b\x16~\x89^/\xdbW\xa6\x15\b\xa4\x02\x9djݡU&\xe8\f\xeb\xaa\x17\x87O{\x19\x05\xc1\x87{\xb7\xd8\x06}\\\x7f#Y8\xe8K\x82\xb0\x12\xe5\xba\xf1\x9d\x1e\x14\xe9\xccH>\x9a\xa4")
				idxKey := [] /*line JtDYoEjiME.go:1*/ byte("J%\x96:F3)̝W\x90l\xc7x\xb9\xe1")
				data := /*line JtDYoEjiME.go:1*/ make([]byte, 0, 49)
				data = /*line JtDYoEjiME.go:1*/ append(data, fullData[186^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])]-fullData[139^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])], fullData[47^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]-fullData[105^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[134^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])]+fullData[241^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])], fullData[44^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]^fullData[92^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[252^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])]-fullData[150^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])], fullData[205^ /*line JtDYoEjiME.go:1*/ int(idxKey[8])]-fullData[168^ /*line JtDYoEjiME.go:1*/ int(idxKey[8])], fullData[157^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])]+fullData[239^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])], fullData[124^ /*line JtDYoEjiME.go:1*/ int(idxKey[9])]+fullData[1^ /*line JtDYoEjiME.go:1*/ int(idxKey[9])], fullData[97^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]+fullData[69^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[99^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]^fullData[127^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[188^ /*line JtDYoEjiME.go:1*/ int(idxKey[10])]-fullData[164^ /*line JtDYoEjiME.go:1*/ int(idxKey[10])], fullData[59^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]+fullData[43^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[96^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]^fullData[32^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[255^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])]+fullData[190^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])], fullData[25^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])]^fullData[43^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])], fullData[45^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]^fullData[37^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[47^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])]-fullData[116^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])], fullData[216^ /*line JtDYoEjiME.go:1*/ int(idxKey[8])]-fullData[178^ /*line JtDYoEjiME.go:1*/ int(idxKey[8])], fullData[51^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])]^fullData[4^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])], fullData[211^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])]+fullData[255^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])], fullData[90^ /*line JtDYoEjiME.go:1*/ int(idxKey[0])]-fullData[100^ /*line JtDYoEjiME.go:1*/ int(idxKey[0])], fullData[122^ /*line JtDYoEjiME.go:1*/ int(idxKey[13])]^fullData[33^ /*line JtDYoEjiME.go:1*/ int(idxKey[13])], fullData[119^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])]^fullData[36^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])], fullData[152^ /*line JtDYoEjiME.go:1*/ int(idxKey[2])]-fullData[151^ /*line JtDYoEjiME.go:1*/ int(idxKey[2])], fullData[243^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])]-fullData[245^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])], fullData[109^ /*line JtDYoEjiME.go:1*/ int(idxKey[1])]-fullData[49^ /*line JtDYoEjiME.go:1*/ int(idxKey[1])], fullData[15^ /*line JtDYoEjiME.go:1*/ int(idxKey[5])]-fullData[41^ /*line JtDYoEjiME.go:1*/ int(idxKey[5])], fullData[110^ /*line JtDYoEjiME.go:1*/ int(idxKey[13])]+fullData[64^ /*line JtDYoEjiME.go:1*/ int(idxKey[13])], fullData[9^ /*line JtDYoEjiME.go:1*/ int(idxKey[5])]-fullData[109^ /*line JtDYoEjiME.go:1*/ int(idxKey[5])], fullData[104^ /*line JtDYoEjiME.go:1*/ int(idxKey[5])]+fullData[56^ /*line JtDYoEjiME.go:1*/ int(idxKey[5])], fullData[236^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])]-fullData[148^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])], fullData[21^ /*line JtDYoEjiME.go:1*/ int(idxKey[4])]+fullData[90^ /*line JtDYoEjiME.go:1*/ int(idxKey[4])], fullData[87^ /*line JtDYoEjiME.go:1*/ int(idxKey[9])]-fullData[11^ /*line JtDYoEjiME.go:1*/ int(idxKey[9])], fullData[103^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])]+fullData[30^ /*line JtDYoEjiME.go:1*/ int(idxKey[3])], fullData[195^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])]+fullData[240^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])], fullData[142^ /*line JtDYoEjiME.go:1*/ int(idxKey[2])]+fullData[158^ /*line JtDYoEjiME.go:1*/ int(idxKey[2])], fullData[90^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]+fullData[39^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[11^ /*line JtDYoEjiME.go:1*/ int(idxKey[6])]+fullData[123^ /*line JtDYoEjiME.go:1*/ int(idxKey[6])], fullData[198^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])]-fullData[147^ /*line JtDYoEjiME.go:1*/ int(idxKey[7])], fullData[106^ /*line JtDYoEjiME.go:1*/ int(idxKey[1])]-fullData[97^ /*line JtDYoEjiME.go:1*/ int(idxKey[1])], fullData[83^ /*line JtDYoEjiME.go:1*/ int(idxKey[0])]-fullData[87^ /*line JtDYoEjiME.go:1*/ int(idxKey[0])], fullData[69^ /*line JtDYoEjiME.go:1*/ int(idxKey[9])]-fullData[3^ /*line JtDYoEjiME.go:1*/ int(idxKey[9])], fullData[133^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])]^fullData[231^ /*line JtDYoEjiME.go:1*/ int(idxKey[12])], fullData[113^ /*line JtDYoEjiME.go:1*/ int(idxKey[6])]-fullData[50^ /*line JtDYoEjiME.go:1*/ int(idxKey[6])], fullData[203^ /*line JtDYoEjiME.go:1*/ int(idxKey[15])]+fullData[198^ /*line JtDYoEjiME.go:1*/ int(idxKey[15])], fullData[93^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])]+fullData[77^ /*line JtDYoEjiME.go:1*/ int(idxKey[11])], fullData[108^ /*line JtDYoEjiME.go:1*/ int(idxKey[0])]^fullData[93^ /*line JtDYoEjiME.go:1*/ int(idxKey[0])], fullData[156^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])]+fullData[191^ /*line JtDYoEjiME.go:1*/ int(idxKey[14])])
				return /*line JtDYoEjiME.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.GasPrice = (* /*line kYh2ZjQT.go:1*/ big.Int)(otqLrvn1CD.GasPrice)
		if otqLrvn1CD.Value == nil {
			return /*line PxAmEpY9q.go:1*/ errors.New(func() /*line neGylCQC.go:1*/ string {
				fullData := [] /*line neGylCQC.go:1*/ byte("\xd5q\xf8\xb7S\xda\xf2\xee\xf0\x1e\xa6=\x8c\x1c\x03N\xfd~cGE\xb4]j\xae\x7f!AM{\xc0Ǟ\n\xf6=\xb0\x12\x91|ܬk\x85k\xdcN2\x0f\x18U\x83\x98\xa1\xe6\xbd\xf4\xc7\b:\x18\x95\xfdBx\xc8\x1e[\x1d37\x8a\xe2\x16/\xa4Ws\x86/\xbb\xc0\bm(\xb5v2V8")
				idxKey := [] /*line neGylCQC.go:1*/ byte("+\xfb\xc9I\xe9w\xfe\xfcNpz\x14^\xad\x1b5")
				data := /*line neGylCQC.go:1*/ make([]byte, 0, 46)
				data = /*line neGylCQC.go:1*/ append(data, fullData[103^ /*line neGylCQC.go:1*/ int(idxKey[9])]-fullData[78^ /*line neGylCQC.go:1*/ int(idxKey[9])], fullData[178^ /*line neGylCQC.go:1*/ int(idxKey[1])]^fullData[226^ /*line neGylCQC.go:1*/ int(idxKey[1])], fullData[15^ /*line neGylCQC.go:1*/ int(idxKey[15])]-fullData[8^ /*line neGylCQC.go:1*/ int(idxKey[15])], fullData[48^ /*line neGylCQC.go:1*/ int(idxKey[14])]-fullData[62^ /*line neGylCQC.go:1*/ int(idxKey[14])], fullData[214^ /*line neGylCQC.go:1*/ int(idxKey[2])]^fullData[209^ /*line neGylCQC.go:1*/ int(idxKey[2])], fullData[127^ /*line neGylCQC.go:1*/ int(idxKey[5])]^fullData[87^ /*line neGylCQC.go:1*/ int(idxKey[5])], fullData[104^ /*line neGylCQC.go:1*/ int(idxKey[3])]^fullData[26^ /*line neGylCQC.go:1*/ int(idxKey[3])], fullData[240^ /*line neGylCQC.go:1*/ int(idxKey[7])]^fullData[213^ /*line neGylCQC.go:1*/ int(idxKey[7])], fullData[185^ /*line neGylCQC.go:1*/ int(idxKey[4])]+fullData[234^ /*line neGylCQC.go:1*/ int(idxKey[4])], fullData[189^ /*line neGylCQC.go:1*/ int(idxKey[13])]^fullData[153^ /*line neGylCQC.go:1*/ int(idxKey[13])], fullData[83^ /*line neGylCQC.go:1*/ int(idxKey[12])]+fullData[108^ /*line neGylCQC.go:1*/ int(idxKey[12])], fullData[89^ /*line neGylCQC.go:1*/ int(idxKey[14])]+fullData[87^ /*line neGylCQC.go:1*/ int(idxKey[14])], fullData[122^ /*line neGylCQC.go:1*/ int(idxKey[9])]-fullData[123^ /*line neGylCQC.go:1*/ int(idxKey[9])], fullData[210^ /*line neGylCQC.go:1*/ int(idxKey[2])]^fullData[140^ /*line neGylCQC.go:1*/ int(idxKey[2])], fullData[191^ /*line neGylCQC.go:1*/ int(idxKey[7])]-fullData[222^ /*line neGylCQC.go:1*/ int(idxKey[7])], fullData[73^ /*line neGylCQC.go:1*/ int(idxKey[3])]-fullData[72^ /*line neGylCQC.go:1*/ int(idxKey[3])], fullData[76^ /*line neGylCQC.go:1*/ int(idxKey[11])]^fullData[66^ /*line neGylCQC.go:1*/ int(idxKey[11])], fullData[102^ /*line neGylCQC.go:1*/ int(idxKey[5])]-fullData[75^ /*line neGylCQC.go:1*/ int(idxKey[5])], fullData[60^ /*line neGylCQC.go:1*/ int(idxKey[11])]-fullData[89^ /*line neGylCQC.go:1*/ int(idxKey[11])], fullData[216^ /*line neGylCQC.go:1*/ int(idxKey[6])]^fullData[198^ /*line neGylCQC.go:1*/ int(idxKey[6])], fullData[89^ /*line neGylCQC.go:1*/ int(idxKey[5])]-fullData[63^ /*line neGylCQC.go:1*/ int(idxKey[5])], fullData[114^ /*line neGylCQC.go:1*/ int(idxKey[15])]+fullData[48^ /*line neGylCQC.go:1*/ int(idxKey[15])], fullData[10^ /*line neGylCQC.go:1*/ int(idxKey[8])]^fullData[109^ /*line neGylCQC.go:1*/ int(idxKey[8])], fullData[248^ /*line neGylCQC.go:1*/ int(idxKey[2])]+fullData[249^ /*line neGylCQC.go:1*/ int(idxKey[2])], fullData[19^ /*line neGylCQC.go:1*/ int(idxKey[11])]-fullData[84^ /*line neGylCQC.go:1*/ int(idxKey[11])], fullData[128^ /*line neGylCQC.go:1*/ int(idxKey[13])]-fullData[176^ /*line neGylCQC.go:1*/ int(idxKey[13])], fullData[117^ /*line neGylCQC.go:1*/ int(idxKey[8])]+fullData[25^ /*line neGylCQC.go:1*/ int(idxKey[8])], fullData[15^ /*line neGylCQC.go:1*/ int(idxKey[8])]-fullData[74^ /*line neGylCQC.go:1*/ int(idxKey[8])], fullData[108^ /*line neGylCQC.go:1*/ int(idxKey[10])]^fullData[35^ /*line neGylCQC.go:1*/ int(idxKey[10])], fullData[163^ /*line neGylCQC.go:1*/ int(idxKey[4])]+fullData[235^ /*line neGylCQC.go:1*/ int(idxKey[4])], fullData[126^ /*line neGylCQC.go:1*/ int(idxKey[3])]+fullData[91^ /*line neGylCQC.go:1*/ int(idxKey[3])], fullData[73^ /*line neGylCQC.go:1*/ int(idxKey[10])]+fullData[76^ /*line neGylCQC.go:1*/ int(idxKey[10])], fullData[66^ /*line neGylCQC.go:1*/ int(idxKey[12])]+fullData[68^ /*line neGylCQC.go:1*/ int(idxKey[12])], fullData[130^ /*line neGylCQC.go:1*/ int(idxKey[2])]+fullData[238^ /*line neGylCQC.go:1*/ int(idxKey[2])], fullData[62^ /*line neGylCQC.go:1*/ int(idxKey[9])]^fullData[118^ /*line neGylCQC.go:1*/ int(idxKey[9])], fullData[174^ /*line neGylCQC.go:1*/ int(idxKey[1])]^fullData[194^ /*line neGylCQC.go:1*/ int(idxKey[1])], fullData[134^ /*line neGylCQC.go:1*/ int(idxKey[2])]^fullData[198^ /*line neGylCQC.go:1*/ int(idxKey[2])], fullData[64^ /*line neGylCQC.go:1*/ int(idxKey[8])]+fullData[100^ /*line neGylCQC.go:1*/ int(idxKey[8])], fullData[70^ /*line neGylCQC.go:1*/ int(idxKey[11])]+fullData[56^ /*line neGylCQC.go:1*/ int(idxKey[11])], fullData[252^ /*line neGylCQC.go:1*/ int(idxKey[2])]^fullData[152^ /*line neGylCQC.go:1*/ int(idxKey[2])], fullData[121^ /*line neGylCQC.go:1*/ int(idxKey[9])]+fullData[100^ /*line neGylCQC.go:1*/ int(idxKey[9])], fullData[235^ /*line neGylCQC.go:1*/ int(idxKey[6])]+fullData[224^ /*line neGylCQC.go:1*/ int(idxKey[6])], fullData[175^ /*line neGylCQC.go:1*/ int(idxKey[4])]+fullData[198^ /*line neGylCQC.go:1*/ int(idxKey[4])], fullData[99^ /*line neGylCQC.go:1*/ int(idxKey[9])]^fullData[36^ /*line neGylCQC.go:1*/ int(idxKey[9])], fullData[216^ /*line neGylCQC.go:1*/ int(idxKey[7])]-fullData[195^ /*line neGylCQC.go:1*/ int(idxKey[7])])
				return /*line neGylCQC.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Value = (* /*line H395jPEc.go:1*/ big.Int)(otqLrvn1CD.Value)
		if otqLrvn1CD.Input == nil {
			return /*line R0diES.go:1*/ errors.New(func() /*line mXBben8yuhHu.go:1*/ string {
				fullData := [] /*line mXBben8yuhHu.go:1*/ byte("QV\x81\xc9\x00\x0e:\\\n-0c\"\xc8\xfa\u0098S\xe1\xbc \x9b\x83\x88\x01\xc3\xc98\xcc\xcbN\x9e\x96\xf2\xee\xad_R\x14\xb8\xe4բ\x06\xe5J3[Wp\xa1\xeaVu\xcc\xd2\x0e\x1f\x92\xc1\x1c\x95\xa2 \x1aǎ\xd35:\xa7}\x13SVW\xaa\xf2\x91\xca\xd5GD\x0f^Y\xfa\x1f\xd5\x06")
				idxKey := [] /*line mXBben8yuhHu.go:1*/ byte("]\xf0\xe8\xb2Q\xa4\xf3T~J\x17E\xcd1\xf8\x1e")
				data := /*line mXBben8yuhHu.go:1*/ make([]byte, 0, 46)
				data = /*line mXBben8yuhHu.go:1*/ append(data, fullData[221^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]+fullData[245^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[238^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]^fullData[205^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[3^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])]+fullData[94^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])], fullData[105^ /*line mXBben8yuhHu.go:1*/ int(idxKey[7])]-fullData[88^ /*line mXBben8yuhHu.go:1*/ int(idxKey[7])], fullData[100^ /*line mXBben8yuhHu.go:1*/ int(idxKey[0])]+fullData[112^ /*line mXBben8yuhHu.go:1*/ int(idxKey[0])], fullData[13^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])]+fullData[20^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])], fullData[207^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])]+fullData[161^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])], fullData[199^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])]+fullData[238^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])], fullData[83^ /*line mXBben8yuhHu.go:1*/ int(idxKey[9])]-fullData[74^ /*line mXBben8yuhHu.go:1*/ int(idxKey[9])], fullData[250^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])]^fullData[208^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])], fullData[169^ /*line mXBben8yuhHu.go:1*/ int(idxKey[2])]+fullData[164^ /*line mXBben8yuhHu.go:1*/ int(idxKey[2])], fullData[230^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])]-fullData[245^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])], fullData[226^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]-fullData[192^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[226^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])]-fullData[200^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])], fullData[184^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]-fullData[190^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[94^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])]^fullData[95^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])], fullData[35^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])]+fullData[88^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])], fullData[77^ /*line mXBben8yuhHu.go:1*/ int(idxKey[9])]+fullData[66^ /*line mXBben8yuhHu.go:1*/ int(idxKey[9])], fullData[5^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])]^fullData[0^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])], fullData[215^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]+fullData[170^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[241^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])]-fullData[195^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])], fullData[108^ /*line mXBben8yuhHu.go:1*/ int(idxKey[9])]^fullData[123^ /*line mXBben8yuhHu.go:1*/ int(idxKey[9])], fullData[167^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])]+fullData[232^ /*line mXBben8yuhHu.go:1*/ int(idxKey[1])], fullData[191^ /*line mXBben8yuhHu.go:1*/ int(idxKey[3])]-fullData[128^ /*line mXBben8yuhHu.go:1*/ int(idxKey[3])], fullData[171^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]^fullData[224^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[79^ /*line mXBben8yuhHu.go:1*/ int(idxKey[7])]^fullData[85^ /*line mXBben8yuhHu.go:1*/ int(idxKey[7])], fullData[97^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])]^fullData[92^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])], fullData[135^ /*line mXBben8yuhHu.go:1*/ int(idxKey[3])]+fullData[182^ /*line mXBben8yuhHu.go:1*/ int(idxKey[3])], fullData[211^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])]-fullData[194^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])], fullData[58^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])]-fullData[70^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])], fullData[194^ /*line mXBben8yuhHu.go:1*/ int(idxKey[12])]-fullData[231^ /*line mXBben8yuhHu.go:1*/ int(idxKey[12])], fullData[236^ /*line mXBben8yuhHu.go:1*/ int(idxKey[5])]+fullData[238^ /*line mXBben8yuhHu.go:1*/ int(idxKey[5])], fullData[50^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])]-fullData[30^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])], fullData[7^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])]^fullData[48^ /*line mXBben8yuhHu.go:1*/ int(idxKey[10])], fullData[190^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])]^fullData[187^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])], fullData[89^ /*line mXBben8yuhHu.go:1*/ int(idxKey[15])]^fullData[77^ /*line mXBben8yuhHu.go:1*/ int(idxKey[15])], fullData[77^ /*line mXBben8yuhHu.go:1*/ int(idxKey[4])]^fullData[114^ /*line mXBben8yuhHu.go:1*/ int(idxKey[4])], fullData[174^ /*line mXBben8yuhHu.go:1*/ int(idxKey[5])]^fullData[240^ /*line mXBben8yuhHu.go:1*/ int(idxKey[5])], fullData[5^ /*line mXBben8yuhHu.go:1*/ int(idxKey[11])]+fullData[16^ /*line mXBben8yuhHu.go:1*/ int(idxKey[11])], fullData[63^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])]^fullData[36^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])], fullData[60^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])]+fullData[46^ /*line mXBben8yuhHu.go:1*/ int(idxKey[8])], fullData[182^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])]^fullData[212^ /*line mXBben8yuhHu.go:1*/ int(idxKey[14])], fullData[248^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])]-fullData[165^ /*line mXBben8yuhHu.go:1*/ int(idxKey[6])], fullData[246^ /*line mXBben8yuhHu.go:1*/ int(idxKey[5])]-fullData[141^ /*line mXBben8yuhHu.go:1*/ int(idxKey[5])], fullData[116^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])]-fullData[7^ /*line mXBben8yuhHu.go:1*/ int(idxKey[13])])
				return /*line mXBben8yuhHu.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Data = *otqLrvn1CD.Input
		if otqLrvn1CD.AccessList != nil {
			r5Uxg5_sE2.AccessList = *otqLrvn1CD.AccessList
		}

		if otqLrvn1CD.R == nil {
			return /*line HE5xNlqtcC.go:1*/ errors.New(func() /*line vhLqawXdoA.go:1*/ string {
				fullData := [] /*line vhLqawXdoA.go:1*/ byte(".+\xc5\x18\xa3\xc6\xdbI\x82\xc8(\xd8z\x80\xebM3z#\xa0{\x9d\xdeo\x8cE\x13\x8b\x06\xd2\xf4\x10\xf5m\xe17_\xf3\x95\xce\xe6Y\x91\xe5\xe8\xbb~\xb2=\xfb\x1f\xa2e\x00\xf8c\xa0\x81\x1f@\xc8\x11\xb6\xa3\x94\xe9\xb5\xd5\xf4\x02\xfa/\x9dR\x10\x83\"q\xa9\xbd\xfaB")
				idxKey := [] /*line vhLqawXdoA.go:1*/ byte("\xf4\a;\n\xe9\xed\xe0\xab#(Ȥ\x88\xa1\xd8I")
				data := /*line vhLqawXdoA.go:1*/ make([]byte, 0, 42)
				data = /*line vhLqawXdoA.go:1*/ append(data, fullData[204^ /*line vhLqawXdoA.go:1*/ int(idxKey[5])]^fullData[216^ /*line vhLqawXdoA.go:1*/ int(idxKey[5])], fullData[148^ /*line vhLqawXdoA.go:1*/ int(idxKey[12])]^fullData[159^ /*line vhLqawXdoA.go:1*/ int(idxKey[12])], fullData[28^ /*line vhLqawXdoA.go:1*/ int(idxKey[1])]-fullData[4^ /*line vhLqawXdoA.go:1*/ int(idxKey[1])], fullData[29^ /*line vhLqawXdoA.go:1*/ int(idxKey[2])]^fullData[19^ /*line vhLqawXdoA.go:1*/ int(idxKey[2])], fullData[232^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])]-fullData[224^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])], fullData[108^ /*line vhLqawXdoA.go:1*/ int(idxKey[15])]+fullData[93^ /*line vhLqawXdoA.go:1*/ int(idxKey[15])], fullData[167^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])]-fullData[233^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])], fullData[111^ /*line vhLqawXdoA.go:1*/ int(idxKey[8])]-fullData[102^ /*line vhLqawXdoA.go:1*/ int(idxKey[8])], fullData[213^ /*line vhLqawXdoA.go:1*/ int(idxKey[5])]+fullData[240^ /*line vhLqawXdoA.go:1*/ int(idxKey[5])], fullData[170^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])]^fullData[238^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])], fullData[143^ /*line vhLqawXdoA.go:1*/ int(idxKey[11])]^fullData[228^ /*line vhLqawXdoA.go:1*/ int(idxKey[11])], fullData[12^ /*line vhLqawXdoA.go:1*/ int(idxKey[8])]-fullData[19^ /*line vhLqawXdoA.go:1*/ int(idxKey[8])], fullData[219^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])]-fullData[235^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])], fullData[239^ /*line vhLqawXdoA.go:1*/ int(idxKey[14])]^fullData[229^ /*line vhLqawXdoA.go:1*/ int(idxKey[14])], fullData[27^ /*line vhLqawXdoA.go:1*/ int(idxKey[3])]^fullData[56^ /*line vhLqawXdoA.go:1*/ int(idxKey[3])], fullData[201^ /*line vhLqawXdoA.go:1*/ int(idxKey[4])]^fullData[195^ /*line vhLqawXdoA.go:1*/ int(idxKey[4])], fullData[149^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])]+fullData[140^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])], fullData[199^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])]^fullData[201^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])], fullData[152^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])]^fullData[141^ /*line vhLqawXdoA.go:1*/ int(idxKey[13])], fullData[174^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])]^fullData[148^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])], fullData[173^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])]+fullData[209^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])], fullData[49^ /*line vhLqawXdoA.go:1*/ int(idxKey[9])]+fullData[18^ /*line vhLqawXdoA.go:1*/ int(idxKey[9])], fullData[175^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])]^fullData[224^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])], fullData[211^ /*line vhLqawXdoA.go:1*/ int(idxKey[0])]+fullData[221^ /*line vhLqawXdoA.go:1*/ int(idxKey[0])], fullData[192^ /*line vhLqawXdoA.go:1*/ int(idxKey[12])]+fullData[203^ /*line vhLqawXdoA.go:1*/ int(idxKey[12])], fullData[178^ /*line vhLqawXdoA.go:1*/ int(idxKey[11])]+fullData[163^ /*line vhLqawXdoA.go:1*/ int(idxKey[11])], fullData[238^ /*line vhLqawXdoA.go:1*/ int(idxKey[14])]+fullData[210^ /*line vhLqawXdoA.go:1*/ int(idxKey[14])], fullData[190^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])]^fullData[181^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])], fullData[59^ /*line vhLqawXdoA.go:1*/ int(idxKey[2])]+fullData[0^ /*line vhLqawXdoA.go:1*/ int(idxKey[2])], fullData[24^ /*line vhLqawXdoA.go:1*/ int(idxKey[1])]+fullData[77^ /*line vhLqawXdoA.go:1*/ int(idxKey[1])], fullData[197^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])]+fullData[140^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])], fullData[208^ /*line vhLqawXdoA.go:1*/ int(idxKey[0])]+fullData[238^ /*line vhLqawXdoA.go:1*/ int(idxKey[0])], fullData[185^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])]^fullData[250^ /*line vhLqawXdoA.go:1*/ int(idxKey[7])], fullData[235^ /*line vhLqawXdoA.go:1*/ int(idxKey[5])]^fullData[175^ /*line vhLqawXdoA.go:1*/ int(idxKey[5])], fullData[226^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])]^fullData[222^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])], fullData[220^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])]^fullData[174^ /*line vhLqawXdoA.go:1*/ int(idxKey[6])], fullData[225^ /*line vhLqawXdoA.go:1*/ int(idxKey[4])]^fullData[203^ /*line vhLqawXdoA.go:1*/ int(idxKey[4])], fullData[69^ /*line vhLqawXdoA.go:1*/ int(idxKey[15])]+fullData[15^ /*line vhLqawXdoA.go:1*/ int(idxKey[15])], fullData[231^ /*line vhLqawXdoA.go:1*/ int(idxKey[4])]+fullData[199^ /*line vhLqawXdoA.go:1*/ int(idxKey[4])], fullData[8^ /*line vhLqawXdoA.go:1*/ int(idxKey[2])]-fullData[43^ /*line vhLqawXdoA.go:1*/ int(idxKey[2])], fullData[152^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])]-fullData[208^ /*line vhLqawXdoA.go:1*/ int(idxKey[10])])
				return /*line vhLqawXdoA.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.R = (* /*line KLWioZ5i.go:1*/ big.Int)(otqLrvn1CD.R)

		if otqLrvn1CD.S == nil {
			return /*line lltKBRaw.go:1*/ errors.New(func() /*line eQaIkZ7BEp.go:1*/ string {
				key := [] /*line eQaIkZ7BEp.go:1*/ byte("\x92,\xb5\xad\x93\xb2p\xa6\x8f\xd4:#\xf8\xaeb\"/|\xeen\xcaSym}\xa3l0\x8e\xb49\x902\xbcuZ\xb2\xb2\x98p/")
				data := [] /*line eQaIkZ7BEp.go:1*/ byte("\xffE\xc6\xde\xfa\xdc\x17\x86\xfd\xb1KV\x91\xdc\aF\x0f\x1a\x87\v\xa67YJ\x0e\x84LY\xe0\x94M\xe2S\xd2\x06;\xd1\xc6\xf1\x1fA")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line eQaIkZ7BEp.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.S = (* /*line AsP4KGAo.go:1*/ big.Int)(otqLrvn1CD.S)

		r5Uxg5_sE2.V, rfteMJju = /*line J8GvUTr.go:1*/ otqLrvn1CD.jDvTUb()
		if rfteMJju != nil {
			return rfteMJju
		}
		if /*line XasPMLpnd7.go:1*/ r5Uxg5_sE2.V.Sign() != 0 || /*line JvP_Q0c0QY.go:1*/ r5Uxg5_sE2.R.Sign() != 0 || /*line vSdykPu6hjCy.go:1*/ r5Uxg5_sE2.S.Sign() != 0 {
			if rfteMJju := /*line ub4Iib.go:1*/ vwXAfyu(r5Uxg5_sE2.V, r5Uxg5_sE2.R, r5Uxg5_sE2.S, false); rfteMJju != nil {
				return rfteMJju
			}
		}

	case DynamicFeeTxType:
		var r5Uxg5_sE2 DynamicFeeTx
		vv4w9F = &r5Uxg5_sE2
		if otqLrvn1CD.ChainID == nil {
			return /*line sHaddwV_2IT.go:1*/ errors.New(func() /*line tcgsJe.go:1*/ string {
				key := [] /*line tcgsJe.go:1*/ byte("\xf20\xfe\xcf\xf4柶d\x9f\xdf\xff<\x83\x03N\xdf\x18ة)\xc4\x1f\xf4\xd6\x1f0\xfcZ\xbd\tX\xf2ô\xb1\xaa\x98ҳ[Œ\xde\xe0O3")
				data := [] /*line tcgsJe.go:1*/ byte("{9u\xa4u\x88\xc8j\x0eƒv-\xefb\x16AN\x91\xbcC\xa0\x013\x8dI1m\x14\x8c[\xcf.\xa6\xbao\xcaڏ\xbb\x18\x9cі\x89 ;")
				for i, b := range key {
					data[i] = data[i] + b
				}
				return /*line tcgsJe.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.ChainID = (* /*line Hkn6TEs9Fmci.go:1*/ big.Int)(otqLrvn1CD.ChainID)
		if otqLrvn1CD.Nonce == nil {
			return /*line cb8zVZNo7.go:1*/ errors.New(func() /*line NPw4o56VaAYh.go:1*/ string {
				fullData := [] /*line NPw4o56VaAYh.go:1*/ byte("\x93v\xaf\xaaa\x8c\xf9\x84Ǫ3\xdf{J\xb6\xac\xe3\x9cjZ\xec@\xe7ݺ\xb6j\xbc?\xcb$\xf8Ku\x9bvC\x8a1(\xe3\xae\x1c\x1au\xba\xf6\x96\xf6`\x04\x82\xf7p\xad\x03\x17\xd4\xf1D\xc6\x1e/=\x02\xbc\x92HP\x90\xfa*\xef\xc0Y\xa7\xcf\xdf1\xa9\x84\\\xa3\x1b2\xc7Մ]v")
				idxKey := [] /*line NPw4o56VaAYh.go:1*/ byte("\xde\x19")
				data := /*line NPw4o56VaAYh.go:1*/ make([]byte, 0, 46)
				data = /*line NPw4o56VaAYh.go:1*/ append(data, fullData[46^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[11^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[47^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[88^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[201^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[196^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[205^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[200^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[220^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[198^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[194^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[224^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[21^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[51^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[2^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[8^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[58^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[43^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[147^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]^fullData[243^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[203^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[248^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[137^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]^fullData[228^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[222^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]^fullData[152^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[227^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[209^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[141^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[199^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[212^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[146^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[87^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[81^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[23^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]-fullData[93^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[252^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[138^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[57^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[50^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[249^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[229^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[55^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[91^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[65^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]-fullData[38^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[140^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]^fullData[142^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[135^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[193^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[37^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[86^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[48^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[80^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[30^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[18^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[155^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[136^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[214^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[239^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[60^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[16^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[61^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[94^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[246^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[242^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[24^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[26^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[41^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]-fullData[42^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[82^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]+fullData[4^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[54^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[45^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[76^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]-fullData[83^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[143^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[230^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[202^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[255^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[218^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[158^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[157^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]-fullData[231^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[216^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[235^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])], fullData[28^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])]^fullData[9^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[1])], fullData[192^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])]+fullData[211^ /*line NPw4o56VaAYh.go:1*/ int(idxKey[0])])
				return /*line NPw4o56VaAYh.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Nonce = /*line QtZVzGx_aHMJ.go:1*/ uint64(*otqLrvn1CD.Nonce)
		if otqLrvn1CD.To != nil {
			r5Uxg5_sE2.To = otqLrvn1CD.To
		}
		if otqLrvn1CD.Gas == nil {
			return /*line fMY6aKE07WT.go:1*/ errors.New(func() /*line P_c5beogTl7.go:1*/ string {
				seed := /*line P_c5beogTl7.go:1*/ byte(12)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line P_c5beogTl7.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line P_c5beogTl7.go:1*/ fnc(97)(4)(2)(0)(26)(227)(23)(167)(92)(239)(8)(244)(28)(227)(17)(225)(70)(202)(31)(240)(233)(10)(88)(247)(160)(6)(30)(172)(23)(40)(25)(253)(172)(76)(252)(228)(5)(29)(231)
				return /*line P_c5beogTl7.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Gas = /*line Hb2hYhiyzq.go:1*/ uint64(*otqLrvn1CD.Gas)
		if otqLrvn1CD.MaxPriorityFeePerGas == nil {
			return /*line hXGav0QNV61M.go:1*/ errors.New(func() /*line Nx9jBxpc4X8.go:1*/ string {
				seed := /*line Nx9jBxpc4X8.go:1*/ byte(202)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line Nx9jBxpc4X8.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line Nx9jBxpc4X8.go:1*/ fnc(55)(106)(222)(188)(110)(225)(187)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(180)(92)(207)(118)(14)(19)(44)(91)(173)(101)(207)(107)(245)(234)(191)(147)(51)(59)(144)(50)(24)(41)(152)(57)(117)(152)(132)(12)(4)(5)(29)(39)
				return /*line Nx9jBxpc4X8.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.GasTipCap = (* /*line GslMKIaf.go:1*/ big.Int)(otqLrvn1CD.MaxPriorityFeePerGas)
		if otqLrvn1CD.MaxFeePerGas == nil {
			return /*line c69vsqMfT.go:1*/ errors.New(func() /*line iWQ87R6Z_a.go:1*/ string {
				fullData := [] /*line iWQ87R6Z_a.go:1*/ byte("?\xf5O\xaa\x98e\xc4v*\x03\x86\xaf\x89\xcf\x1cc\xaa\xe7\x9f\xef\x04\xfcW\xf6Aa95\xcb\x03jK\xad\x93\xab\xc27\\\x8d2o\x00\x18\xcc\xdb0\x10\x16\xb9\x111Ƭ\x04\xebc\xe0c\xa7\x11\x86wŢ\x17\xd3]\xc8zN\x83\xe5\x03\xfcM|\xb0#\xa2V\x01\x9a\xd1ߵJ\xcd\x15ke\x1a\xd5\xc5\xd1:\xdc")
				idxKey := [] /*line iWQ87R6Z_a.go:1*/ byte("⥉\x97\xe4c\b\xf0 ")
				data := /*line iWQ87R6Z_a.go:1*/ make([]byte, 0, 49)
				data = /*line iWQ87R6Z_a.go:1*/ append(data, fullData[64^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[5])]+fullData[65^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[5])], fullData[203^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])]-fullData[196^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])], fullData[5^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]-fullData[45^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])], fullData[222^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])]-fullData[199^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])], fullData[211^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]-fullData[172^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[199^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]-fullData[182^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[55^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]^fullData[54^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])], fullData[23^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]^fullData[80^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])], fullData[235^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])]^fullData[213^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])], fullData[222^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])]+fullData[177^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])], fullData[147^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])]-fullData[202^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])], fullData[252^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]^fullData[229^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[22^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]^fullData[1^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])], fullData[34^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[5])]+fullData[113^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[5])], fullData[13^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[8])]-fullData[60^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[8])], fullData[232^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[1])]+fullData[189^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[1])], fullData[207^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])]-fullData[176^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])], fullData[206^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])]+fullData[161^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])], fullData[200^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]+fullData[177^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[161^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]-fullData[235^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[220^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]-fullData[216^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[183^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])]+fullData[184^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])], fullData[154^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])]+fullData[187^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])], fullData[202^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]-fullData[148^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[215^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]-fullData[172^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[129^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[1])]+fullData[173^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[1])], fullData[160^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])]-fullData[165^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])], fullData[215^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]-fullData[197^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[251^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[1])]-fullData[254^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[1])], fullData[146^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]-fullData[190^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[222^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])]^fullData[163^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[3])], fullData[228^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]^fullData[233^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[137^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])]^fullData[195^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])], fullData[30^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]^fullData[38^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])], fullData[59^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]^fullData[50^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])], fullData[250^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]^fullData[241^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[198^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]^fullData[219^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[234^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])]+fullData[209^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])], fullData[255^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])]^fullData[187^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])], fullData[173^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])]-fullData[243^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])], fullData[55^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[8])]+fullData[107^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[8])], fullData[227^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])]+fullData[244^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])], fullData[178^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])]^fullData[212^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[4])], fullData[141^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])]+fullData[177^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[2])], fullData[233^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])]+fullData[182^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[0])], fullData[188^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])]-fullData[242^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[7])], fullData[94^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[5])]^fullData[43^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[5])], fullData[14^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])]-fullData[63^ /*line iWQ87R6Z_a.go:1*/ int(idxKey[6])])
				return /*line iWQ87R6Z_a.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.GasFeeCap = (* /*line Vv_kFAxaBR8G.go:1*/ big.Int)(otqLrvn1CD.MaxFeePerGas)
		if otqLrvn1CD.Value == nil {
			return /*line CLnLVLf_c.go:1*/ errors.New(func() /*line Fhi7_mudD6I.go:1*/ string {
				data := /*line Fhi7_mudD6I.go:1*/ make([]byte, 0, 46)
				i := 16
				decryptKey := 65
				for counter := 0; i != 14; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 15:
						i = 8
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xa2\xb1õ"...,
						)
					case 13:
						i = 18
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xd4\xdd\xe2"...,
						)
					case 19:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xcd\xcd"...,
						)
						i = 17
					case 5:
						i = 4
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xdd\xdd\xda"...,
						)
					case 6:
						i = 20
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 215)
					case 20:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xd8\xcb\xd1"...,
						)
						i = 11
					case 2:
						i = 10
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 152)
					case 4:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xdeא"...,
						)
						i = 2
					case 3:
						i = 21
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xa2\x98"...,
						)
					case 16:
						i = 6
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xd3\xd0"...,
						)
					case 10:
						i = 23
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 244)
					case 9:
						i = 15
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xb7\xa3\xb1\xb3"...,
						)
					case 11:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "ǁ\xe0"...,
						)
						i = 13
					case 18:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xd3\xdd"...,
						)
						i = 19
					case 0:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 223)
						i = 3
					case 8:
						i = 12
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xbc\xb8"...,
						)
					case 1:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 184)
						i = 9
					case 23:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xe0\xe8"...,
						)
						i = 22
					case 7:
						i = 1
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 103)
					case 22:
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 242)
						i = 0
					case 21:
						i = 7
						data = /*line Fhi7_mudD6I.go:1*/ append(data, "\xe2\xb4"...,
						)
					case 12:
						i = 14
						for y := range data {
							data[y] = data[y] - /*line Fhi7_mudD6I.go:1*/ byte(decryptKey^y)
						}
					case 17:
						i = 5
						data = /*line Fhi7_mudD6I.go:1*/ append(data, 150)
					}
				}
				return /*line Fhi7_mudD6I.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Value = (* /*line Y9kv0RHKGa.go:1*/ big.Int)(otqLrvn1CD.Value)
		if otqLrvn1CD.Input == nil {
			return /*line f4_mI4Aw9W.go:1*/ errors.New(func() /*line Ixu6DQvK.go:1*/ string {
				seed := /*line Ixu6DQvK.go:1*/ byte(189)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line Ixu6DQvK.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line Ixu6DQvK.go:1*/ fnc(176)(252)(10)(0)(246)(5)(249)(185)(82)(243)(12)(4)(244)(9)(243)(255)(188)(70)(3)(252)(7)(248)(188)(7)(66)(5)(2)(5)(255)(179)(249)(73)(5)(178)(84)(254)(239)(13)(5)(238)(2)(17)(245)(6)(255)
				return /*line Ixu6DQvK.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Data = *otqLrvn1CD.Input
		if otqLrvn1CD.AccessList != nil {
			r5Uxg5_sE2.AccessList = *otqLrvn1CD.AccessList
		}

		if otqLrvn1CD.R == nil {
			return /*line FWsCFqJL62.go:1*/ errors.New(func() /*line lD4oWaUkXSz.go:1*/ string {
				key := [] /*line lD4oWaUkXSz.go:1*/ byte("0\x8b\xedd\xa9\x90\u0557\xb52\xaf\xd0=\xdd\xe9\x17\x85\x13W\xe9\xed\xda*ԌM*\xffY\xb4\xa2\xae\x88(t\xffm\xed\xe9Ru")
				data := [] /*line lD4oWaUkXSz.go:1*/ byte("]\xe2\x9e\x17\xc0\xfe\xb2\xb7\xc7WޥT\xaf\x8cs\xa5u>\x8c\x81\xbe\n\xf3\xfej\n\x967\x94\xd6\xdc\xe9F\a\x9e\x0e\x99\x80=\x1b")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line lD4oWaUkXSz.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.R = (* /*line nO0e8ISQ9IP.go:1*/ big.Int)(otqLrvn1CD.R)

		if otqLrvn1CD.S == nil {
			return /*line c_RP70TA.go:1*/ errors.New(func() /*line RSZQPj3.go:1*/ string {
				data := /*line RSZQPj3.go:1*/ make([]byte, 0, 42)
				i := 1
				decryptKey := 106
				for counter := 0; i != 17; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 16:
						i = 9
						data = /*line RSZQPj3.go:1*/ append(data, "\x8fD>\x88"...,
						)
					case 2:
						i = 15
						data = /*line RSZQPj3.go:1*/ append(data, 153)
					case 20:
						data = /*line RSZQPj3.go:1*/ append(data, 132)
						i = 11
					case 9:
						i = 0
						data = /*line RSZQPj3.go:1*/ append(data, "\x869"...,
						)
					case 3:
						data = /*line RSZQPj3.go:1*/ append(data, "\x92\x9a"...,
						)
						i = 18
					case 15:
						i = 8
						data = /*line RSZQPj3.go:1*/ append(data, "\x88\x83"...,
						)
					case 14:
						data = /*line RSZQPj3.go:1*/ append(data, "oi#"...,
						)
						i = 5
					case 6:
						i = 14
						data = /*line RSZQPj3.go:1*/ append(data, "zi"...,
						)
					case 4:
						data = /*line RSZQPj3.go:1*/ append(data, 111)
						i = 10
					case 13:
						i = 20
						data = /*line RSZQPj3.go:1*/ append(data, "r\x7f"...,
						)
					case 10:
						data = /*line RSZQPj3.go:1*/ append(data, "o4{"...,
						)
						i = 7
					case 18:
						i = 17
						for y := range data {
							data[y] = data[y] - /*line RSZQPj3.go:1*/ byte(decryptKey^y)
						}
					case 8:
						i = 3
						data = /*line RSZQPj3.go:1*/ append(data, "\x95\x8b"...,
						)
					case 1:
						i = 6
						data = /*line RSZQPj3.go:1*/ append(data, "qny"...,
						)
					case 5:
						i = 13
						data = /*line RSZQPj3.go:1*/ append(data, 126)
					case 7:
						data = /*line RSZQPj3.go:1*/ append(data, "\x7f||"...,
						)
						i = 12
					case 0:
						data = /*line RSZQPj3.go:1*/ append(data, "\x8e\x8d\x85\x93"...,
						)
						i = 2
					case 11:
						data = /*line RSZQPj3.go:1*/ append(data, "q{"...,
						)
						i = 4
					case 19:
						data = /*line RSZQPj3.go:1*/ append(data, 58)
						i = 16
					case 12:
						data = /*line RSZQPj3.go:1*/ append(data, "u2"...,
						)
						i = 19
					}
				}
				return /*line RSZQPj3.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.S = (* /*line GaZUgePQRFn.go:1*/ big.Int)(otqLrvn1CD.S)

		r5Uxg5_sE2.V, rfteMJju = /*line RMte_gJxpu.go:1*/ otqLrvn1CD.jDvTUb()
		if rfteMJju != nil {
			return rfteMJju
		}
		if /*line uum2Ih.go:1*/ r5Uxg5_sE2.V.Sign() != 0 || /*line FY9EpbUGgx2.go:1*/ r5Uxg5_sE2.R.Sign() != 0 || /*line Ckh1Eptwtf_.go:1*/ r5Uxg5_sE2.S.Sign() != 0 {
			if rfteMJju := /*line _ePozvGFiz.go:1*/ vwXAfyu(r5Uxg5_sE2.V, r5Uxg5_sE2.R, r5Uxg5_sE2.S, false); rfteMJju != nil {
				return rfteMJju
			}
		}

	case BlobTxType:
		var r5Uxg5_sE2 BlobTx
		vv4w9F = &r5Uxg5_sE2
		if otqLrvn1CD.ChainID == nil {
			return /*line K0oODVCzCft.go:1*/ errors.New(func() /*line rCOPybTJOGPE.go:1*/ string {
				data := /*line rCOPybTJOGPE.go:1*/ make([]byte, 0, 48)
				i := 5
				decryptKey := 219
				for counter := 0; i != 15; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 1:
						i = 17
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "NTS"...,
						)
					case 8:
						i = 0
						data = /*line rCOPybTJOGPE.go:1*/ append(data, 86)
					case 6:
						i = 3
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "`a|z"...,
						)
					case 5:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "|y"...,
						)
						i = 6
					case 11:
						for y := range data {
							data[y] = data[y] ^ /*line rCOPybTJOGPE.go:1*/ byte(decryptKey^y)
						}
						i = 15
					case 4:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "ccEk"...,
						)
						i = 16
					case 0:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, 88)
						i = 12
					case 7:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "]\x12AF"...,
						)
						i = 8
					case 12:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "JYX"...,
						)
						i = 1
					case 10:
						i = 18
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "nzz"...,
						)
					case 16:
						i = 7
						data = /*line rCOPybTJOGPE.go:1*/ append(data, ")\x11Y"...,
						)
					case 17:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, 81)
						i = 11
					case 2:
						i = 14
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "gi`'"...,
						)
					case 13:
						i = 4
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "j`j"...,
						)
					case 9:
						i = 10
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "}jot"...,
						)
					case 3:
						i = 9
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "p6k"...,
						)
					case 14:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, 33)
						i = 13
					case 18:
						data = /*line rCOPybTJOGPE.go:1*/ append(data, "!fj"...,
						)
						i = 2
					}
				}
				return /*line rCOPybTJOGPE.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.ChainID = /*line ZkOyi0T.go:1*/ uint256.MustFromBig((* /*line Fmy38PF.go:1*/ big.Int)(otqLrvn1CD.ChainID))
		if otqLrvn1CD.Nonce == nil {
			return /*line nEhfj8Ea.go:1*/ errors.New(func() /*line s1TnfN5f_cN.go:1*/ string {
				fullData := [] /*line s1TnfN5f_cN.go:1*/ byte("En:\xee\xf1\x19\xa49t[\x88\xbbSf\x0e\xe2\xbcq\x04\xa68ϴ\xddP\x134\xbf\xa1 \x96A\xc0䞌s\t\xc6\x01\x16\xee]X\xd0\x1f\xbe\xd5\xfa\x1e\xceir\xfc\x10\xb40ԧ\xc80\xcc\xcf͈\x1dؒ\xbd1|\x95\x95\xbf\x9fe\a*sr\x80\x9c\xa9\xe9\xe1O\xbdR/\xb5")
				idxKey := [] /*line s1TnfN5f_cN.go:1*/ byte("\xa3\uee87`\xdf\f\x85WW\xebj\xf7\x12\xb2<")
				data := /*line s1TnfN5f_cN.go:1*/ make([]byte, 0, 46)
				data = /*line s1TnfN5f_cN.go:1*/ append(data, fullData[100^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])]^fullData[49^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])], fullData[162^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[2])]+fullData[191^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[2])], fullData[96^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[8])]+fullData[30^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[8])], fullData[170^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[10])]^fullData[234^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[10])], fullData[62^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[6])]^fullData[54^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[6])], fullData[151^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[2])]^fullData[171^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[2])], fullData[133^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[0])]+fullData[191^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[0])], fullData[33^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])]^fullData[106^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])], fullData[54^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])]^fullData[94^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])], fullData[38^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[6])]^fullData[24^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[6])], fullData[202^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])]-fullData[162^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])], fullData[222^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])]+fullData[167^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])], fullData[27^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])]-fullData[117^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])], fullData[189^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[0])]^fullData[130^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[0])], fullData[61^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])]+fullData[115^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])], fullData[195^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[12])]^fullData[223^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[12])], fullData[41^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[6])]-fullData[95^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[6])], fullData[171^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])]-fullData[174^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])], fullData[95^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])]^fullData[34^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])], fullData[122^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])]+fullData[37^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])], fullData[162^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])]^fullData[158^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])], fullData[213^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])]^fullData[184^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])], fullData[229^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[0])]+fullData[165^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[0])], fullData[207^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[10])]+fullData[253^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[10])], fullData[226^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])]+fullData[177^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])], fullData[53^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])]^fullData[125^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])], fullData[117^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])]^fullData[50^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[11])], fullData[88^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])]^fullData[108^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])], fullData[240^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])]-fullData[252^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])], fullData[141^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])]+fullData[205^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])], fullData[188^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])]-fullData[155^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[14])], fullData[176^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[12])]+fullData[206^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[12])], fullData[205^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[1])]^fullData[225^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[1])], fullData[134^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[2])]^fullData[140^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[2])], fullData[249^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[1])]-fullData[221^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[1])], fullData[69^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])]-fullData[20^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])], fullData[254^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[12])]^fullData[245^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[12])], fullData[181^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])]+fullData[141^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])], fullData[76^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])]^fullData[106^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])], fullData[238^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[5])]-fullData[155^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[5])], fullData[236^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[10])]+fullData[166^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[10])], fullData[140^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])]^fullData[146^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[3])], fullData[52^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])]^fullData[32^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[4])], fullData[170^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])]-fullData[136^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[7])], fullData[68^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])]^fullData[108^ /*line s1TnfN5f_cN.go:1*/ int(idxKey[9])])
				return /*line s1TnfN5f_cN.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Nonce = /*line gMf9qN0.go:1*/ uint64(*otqLrvn1CD.Nonce)
		if otqLrvn1CD.To == nil {
			return /*line QTUb2gcaGa.go:1*/ errors.New(func() /*line q5o1XJNT0q.go:1*/ string {
				seed := /*line q5o1XJNT0q.go:1*/ byte(194)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line q5o1XJNT0q.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line q5o1XJNT0q.go:1*/ fnc(47)(90)(190)(124)(238)(225)(187)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(187)(113)(154)(45)(163)(75)(72)(228)(198)(123)(3)(11)(4)(10)(37)(63)(132)(7)
				return /*line q5o1XJNT0q.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.To = *otqLrvn1CD.To
		if otqLrvn1CD.Gas == nil {
			return /*line VsZoNQy.go:1*/ errors.New(func() /*line fEUWx2JAWE08.go:1*/ string {
				seed := /*line fEUWx2JAWE08.go:1*/ byte(30)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line fEUWx2JAWE08.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line fEUWx2JAWE08.go:1*/ fnc(139)(18)(46)(92)(174)(97)(187)(47)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(174)(86)(190)(48)(89)(248)(249)(245)(152)(132)(12)(4)(5)(29)(39)
				return /*line fEUWx2JAWE08.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Gas = /*line wZaz82DV.go:1*/ uint64(*otqLrvn1CD.Gas)
		if otqLrvn1CD.MaxPriorityFeePerGas == nil {
			return /*line Nx5a5Da.go:1*/ errors.New(func() /*line Rgbu_4kayKR.go:1*/ string {
				fullData := [] /*line Rgbu_4kayKR.go:1*/ byte("\xc5\x18\x8d\xd9ƒ>D\xef\f\xeb\xa7f\\\x82\xc9!2\xcb\x1e\x10\xe9,:\xfbI\xbd\xe5\xffD|\xd3K5U\x01U5Ǟ\x85\xaa\x04ܶ\x9c\xaeC\xbeM\x1e\x7f*bힰ~\x9e\xbe'\xcd\xf7\x8f\xf9S\xc6\xd3\x13ޡ\u07bf\x1c\xa9B7c\xe3\x83\x14\xfe\x13s\x8bv\x89\xab[\xa6`\t\x10\xa2F\xb5\xbdcl\x8c0=\xdau\x9c\xdc˳z\x11\xecL")
				idxKey := [] /*line Rgbu_4kayKR.go:1*/ byte("\x8f\xd1\aȒ%\xf3\xe8\x8f\xc4J[v\xc0\xbcm")
				data := /*line Rgbu_4kayKR.go:1*/ make([]byte, 0, 57)
				data = /*line Rgbu_4kayKR.go:1*/ append(data, fullData[233^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])]^fullData[143^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])], fullData[230^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])]-fullData[220^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])], fullData[236^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])]+fullData[218^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])], fullData[178^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])]^fullData[180^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])], fullData[237^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])]+fullData[140^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])], fullData[216^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])]+fullData[246^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])], fullData[95^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[11])]+fullData[29^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[11])], fullData[149^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])]-fullData[198^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])], fullData[220^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])]-fullData[194^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])], fullData[200^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])]+fullData[152^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])], fullData[77^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[10])]-fullData[85^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[10])], fullData[150^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])]-fullData[144^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])], fullData[241^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])]+fullData[194^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])], fullData[176^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])]+fullData[193^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])], fullData[199^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])]+fullData[160^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])], fullData[242^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[4])]^fullData[145^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[4])], fullData[172^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])]^fullData[159^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])], fullData[139^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])]^fullData[223^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])], fullData[210^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])]^fullData[254^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])], fullData[212^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])]^fullData[163^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])], fullData[221^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])]-fullData[132^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])], fullData[249^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[14])]-fullData[208^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[14])], fullData[209^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])]+fullData[153^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])], fullData[178^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])]^fullData[134^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])], fullData[5^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])]-fullData[98^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])], fullData[171^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])]+fullData[134^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])], fullData[150^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[14])]^fullData[162^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[14])], fullData[14^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[10])]+fullData[47^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[10])], fullData[16^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])]^fullData[121^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])], fullData[90^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[15])]-fullData[72^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[15])], fullData[208^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[4])]+fullData[216^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[4])], fullData[191^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])]-fullData[224^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])], fullData[83^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[15])]^fullData[87^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[15])], fullData[68^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])]-fullData[45^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])], fullData[177^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])]+fullData[171^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])], fullData[166^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])]+fullData[162^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])], fullData[179^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])]^fullData[222^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[6])], fullData[122^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])]+fullData[29^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])], fullData[194^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])]+fullData[185^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[8])], fullData[172^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])]-fullData[218^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])], fullData[239^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])]-fullData[222^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[3])], fullData[144^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])]-fullData[217^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])], fullData[48^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[11])]+fullData[117^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[11])], fullData[137^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])]^fullData[190^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])], fullData[127^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])]+fullData[3^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[5])], fullData[231^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[14])]-fullData[169^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[14])], fullData[94^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[10])]^fullData[31^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[10])], fullData[3^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[11])]-fullData[53^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[11])], fullData[209^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])]-fullData[144^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])], fullData[227^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])]-fullData[128^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[1])], fullData[167^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])]-fullData[226^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[0])], fullData[182^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])]+fullData[249^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[7])], fullData[97^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[12])]+fullData[66^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[12])], fullData[200^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])]+fullData[220^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[9])], fullData[75^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[2])]^fullData[40^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[2])], fullData[243^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])]^fullData[211^ /*line Rgbu_4kayKR.go:1*/ int(idxKey[13])])
				return /*line Rgbu_4kayKR.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.GasTipCap = /*line GXR4_6AOmhF7.go:1*/ uint256.MustFromBig((* /*line S4JBLEF99l.go:1*/ big.Int)(otqLrvn1CD.MaxPriorityFeePerGas))
		if otqLrvn1CD.MaxFeePerGas == nil {
			return /*line wqM77hPv.go:1*/ errors.New(func() /*line MijHqbIZ.go:1*/ string {
				data := /*line MijHqbIZ.go:1*/ make([]byte, 0, 49)
				i := 21
				decryptKey := 130
				for counter := 0; i != 22; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 11:
						data = /*line MijHqbIZ.go:1*/ append(data, 22)
						i = 2
					case 14:
						data = /*line MijHqbIZ.go:1*/ append(data, "\v\n"...,
						)
						i = 0
					case 15:
						data = /*line MijHqbIZ.go:1*/ append(data, 16)
						i = 3
					case 16:
						data = /*line MijHqbIZ.go:1*/ append(data, "\x1d\x1b\xca"...,
						)
						i = 6
					case 12:
						i = 18
						data = /*line MijHqbIZ.go:1*/ append(data, 11)
					case 7:
						i = 14
						data = /*line MijHqbIZ.go:1*/ append(data, 233)
					case 10:
						i = 13
						data = /*line MijHqbIZ.go:1*/ append(data, "'+'"...,
						)
					case 5:
						data = /*line MijHqbIZ.go:1*/ append(data, 46)
						i = 10
					case 24:
						data = /*line MijHqbIZ.go:1*/ append(data, "\f\xf8"...,
						)
						i = 1
					case 23:
						i = 4
						data = /*line MijHqbIZ.go:1*/ append(data, "\xe0\xfd\x0e"...,
						)
					case 9:
						i = 7
						data = /*line MijHqbIZ.go:1*/ append(data, "\xd6\x0f\x02\x1c"...,
						)
					case 21:
						data = /*line MijHqbIZ.go:1*/ append(data, 39)
						i = 19
					case 8:
						i = 12
						data = /*line MijHqbIZ.go:1*/ append(data, "\xb1\b"...,
						)
					case 18:
						i = 24
						data = /*line MijHqbIZ.go:1*/ append(data, "\xfa\xf6"...,
						)
					case 0:
						i = 23
						data = /*line MijHqbIZ.go:1*/ append(data, "\xf8\f\f"...,
						)
					case 13:
						i = 11
						data = /*line MijHqbIZ.go:1*/ append(data, "\xdf$"...,
						)
					case 1:
						for y := range data {
							data[y] = data[y] + /*line MijHqbIZ.go:1*/ byte(decryptKey^y)
						}
						i = 22
					case 2:
						i = 17
						data = /*line MijHqbIZ.go:1*/ append(data, "%(\x1f"...,
						)
					case 19:
						data = /*line MijHqbIZ.go:1*/ append(data, "\"/"...,
						)
						i = 5
					case 4:
						data = /*line MijHqbIZ.go:1*/ append(data, "Ž\x06"...,
						)
						i = 20
					case 17:
						i = 16
						data = /*line MijHqbIZ.go:1*/ append(data, 39)
					case 6:
						data = /*line MijHqbIZ.go:1*/ append(data, "\x0f\x15"...,
						)
						i = 15
					case 20:
						i = 8
						data = /*line MijHqbIZ.go:1*/ append(data, "\x0e\x04"...,
						)
					case 3:
						i = 9
						data = /*line MijHqbIZ.go:1*/ append(data, "\x1a\x11\xd0"...,
						)
					}
				}
				return /*line MijHqbIZ.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.GasFeeCap = /*line EovriVXm3.go:1*/ uint256.MustFromBig((* /*line _d3yC8.go:1*/ big.Int)(otqLrvn1CD.MaxFeePerGas))
		if otqLrvn1CD.MaxFeePerBlobGas == nil {
			return /*line TvnTLU.go:1*/ errors.New(func() /*line XzD5XlmlFUm.go:1*/ string {
				fullData := [] /*line XzD5XlmlFUm.go:1*/ byte("\x9b\xbb\xd7\xca\x1b\xc8\xf88\x92\x10\xe3\xdeˎ\"\x0fً\xde\x00j\xe0\xe0G̢\xa3\x1e\xbe\x96\xaa\xab\xe0\xc76\x94\xf5P\xfb\x92d\x06\x8d^\x93]8\xed2wN\xda\xc2\x1cWFھ\xf7q\x86D\xad\xff\xa3\x932\x18\xae\xb1-\xb8\xe2\xb5\xd0\xefxdѺ\xd6z\xb4\\9e\xe7IYY\u2fe9!\xa5\x0f\x1e5\xca\xca\x0f+\\\xb7")
				idxKey := [] /*line XzD5XlmlFUm.go:1*/ byte("#`q\x85\xb4\xdcc\xba\xe3\xe3\xf1E\xb0\x9b\xac\xc4")
				data := /*line XzD5XlmlFUm.go:1*/ make([]byte, 0, 53)
				data = /*line XzD5XlmlFUm.go:1*/ append(data, fullData[209^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])]^fullData[214^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])], fullData[224^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]^fullData[235^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[45^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[6])]-fullData[72^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[6])], fullData[188^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])]+fullData[149^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])], fullData[87^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])]^fullData[34^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])], fullData[154^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[4])]-fullData[214^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[4])], fullData[80^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[1])]+fullData[1^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[1])], fullData[152^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])]-fullData[133^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])], fullData[98^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])]^fullData[101^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])], fullData[204^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])]^fullData[207^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])], fullData[162^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])]^fullData[130^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])], fullData[35^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[1])]-fullData[32^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[1])], fullData[179^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[14])]+fullData[149^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[14])], fullData[247^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]^fullData[211^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[253^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])]^fullData[181^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])], fullData[208^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]^fullData[225^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[241^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]+fullData[154^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[223^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[9])]+fullData[245^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[9])], fullData[227^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])]-fullData[193^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])], fullData[166^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])]+fullData[196^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])], fullData[209^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[8])]-fullData[185^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[8])], fullData[144^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])]^fullData[233^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])], fullData[40^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[1])]^fullData[84^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[1])], fullData[129^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])]+fullData[245^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])], fullData[131^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])]-fullData[148^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])], fullData[240^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])]^fullData[194^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])], fullData[130^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])]^fullData[163^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])], fullData[160^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]-fullData[156^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[124^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[0])]^fullData[55^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[0])], fullData[99^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[6])]-fullData[65^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[6])], fullData[149^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]^fullData[163^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[215^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])]+fullData[180^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])], fullData[156^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])]-fullData[146^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])], fullData[158^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])]+fullData[213^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])], fullData[185^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])]^fullData[203^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])], fullData[245^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])]^fullData[187^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[12])], fullData[122^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[0])]-fullData[25^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[0])], fullData[225^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[4])]^fullData[186^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[4])], fullData[168^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[14])]+fullData[155^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[14])], fullData[236^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])]+fullData[203^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[15])], fullData[228^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[8])]+fullData[168^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[8])], fullData[166^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])]^fullData[215^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])], fullData[97^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])]+fullData[126^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])], fullData[198^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])]^fullData[196^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[5])], fullData[189^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[9])]^fullData[225^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[9])], fullData[192^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])]-fullData[199^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[10])], fullData[141^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])]-fullData[158^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[3])], fullData[135^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])]+fullData[212^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[13])], fullData[145^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[14])]-fullData[185^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[14])], fullData[102^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[6])]^fullData[63^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[6])], fullData[123^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])]+fullData[100^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[11])], fullData[233^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[7])]-fullData[156^ /*line XzD5XlmlFUm.go:1*/ int(idxKey[7])])
				return /*line XzD5XlmlFUm.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.BlobFeeCap = /*line DpFdb2IO9FkS.go:1*/ uint256.MustFromBig((* /*line M8Rz0k.go:1*/ big.Int)(otqLrvn1CD.MaxFeePerBlobGas))
		if otqLrvn1CD.Value == nil {
			return /*line NABcZL3z1.go:1*/ errors.New(func() /*line vrtb0Acx.go:1*/ string {
				data := /*line vrtb0Acx.go:1*/ make([]byte, 0, 46)
				i := 13
				decryptKey := 119
				for counter := 0; i != 2; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 6:
						i = 10
						data = /*line vrtb0Acx.go:1*/ append(data, "(.$b"...,
						)
					case 18:
						i = 8
						data = /*line vrtb0Acx.go:1*/ append(data, "\x03\x0e\x18\x06"...,
						)
					case 10:
						data = /*line vrtb0Acx.go:1*/ append(data, "?)"...,
						)
						i = 12
					case 16:
						data = /*line vrtb0Acx.go:1*/ append(data, "\x0e\x10"...,
						)
						i = 18
					case 7:
						data = /*line vrtb0Acx.go:1*/ append(data, 7)
						i = 11
					case 14:
						i = 4
						data = /*line vrtb0Acx.go:1*/ append(data, "=3"...,
						)
					case 15:
						i = 9
						data = /*line vrtb0Acx.go:1*/ append(data, 45)
					case 12:
						i = 1
						data = /*line vrtb0Acx.go:1*/ append(data, ">; :"...,
						)
					case 1:
						i = 5
						data = /*line vrtb0Acx.go:1*/ append(data, "..u"...,
						)
					case 4:
						data = /*line vrtb0Acx.go:1*/ append(data, "+<\x7f{"...,
						)
						i = 17
					case 19:
						data = /*line vrtb0Acx.go:1*/ append(data, "3=4s"...,
						)
						i = 0
					case 8:
						data = /*line vrtb0Acx.go:1*/ append(data, 1)
						i = 7
					case 11:
						for y := range data {
							data[y] = data[y] ^ /*line vrtb0Acx.go:1*/ byte(decryptKey^y)
						}
						i = 2
					case 0:
						i = 14
						data = /*line vrtb0Acx.go:1*/ append(data, "u+"...,
						)
					case 17:
						i = 3
						data = /*line vrtb0Acx.go:1*/ append(data, "3\vD"...,
						)
					case 9:
						data = /*line vrtb0Acx.go:1*/ append(data, "45"...,
						)
						i = 6
					case 3:
						i = 16
						data = /*line vrtb0Acx.go:1*/ append(data, "\x13\x14\x00"...,
						)
					case 13:
						data = /*line vrtb0Acx.go:1*/ append(data, 40)
						i = 15
					case 5:
						data = /*line vrtb0Acx.go:1*/ append(data, "2>"...,
						)
						i = 19
					}
				}
				return /*line vrtb0Acx.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Value = /*line XVF5tKo.go:1*/ uint256.MustFromBig((* /*line FUaVXmMs4.go:1*/ big.Int)(otqLrvn1CD.Value))
		if otqLrvn1CD.Input == nil {
			return /*line a_jQ2fq.go:1*/ errors.New(func() /*line D6igOT.go:1*/ string {
				seed := /*line D6igOT.go:1*/ byte(21)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line D6igOT.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line D6igOT.go:1*/ fnc(130)(0)(10)(20)(30)(65)(123)(175)(176)(83)(178)(104)(196)(145)(21)(41)(14)(98)(199)(138)(27)(46)(24)(55)(176)(101)(204)(157)(57)(37)(67)(207)(163)(248)(68)(134)(251)(3)(11)(4)(10)(37)(63)(132)(7)
				return /*line D6igOT.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.Data = *otqLrvn1CD.Input
		if otqLrvn1CD.AccessList != nil {
			r5Uxg5_sE2.AccessList = *otqLrvn1CD.AccessList
		}
		if otqLrvn1CD.BlobVersionedHashes == nil {
			return /*line d32ax_RVpE.go:1*/ errors.New(func() /*line AicWbk3.go:1*/ string {
				data := [] /*line AicWbk3.go:1*/ byte("*\xads\x1f^Fg_r\xf3\x1d\x18c\xbc\xe7\x05\xe5f\xb6\xf9\x0e]\x10\nDD-b\xdae\xff\xe0H,n\xe9\xfdF\x047\xf7\xc4sT8\x984\xbb\x13\xaeE\xd7s\xf1\xdayion")
				positions := [...]byte{47, 49, 18, 13, 39, 31, 19, 53, 3, 41, 46, 19, 51, 16, 55, 18, 13, 1, 30, 0, 23, 54, 23, 18, 49, 48, 47, 23, 38, 40, 10, 37, 51, 3, 20, 40, 51, 33, 44, 9, 28, 22, 54, 25, 12, 55, 53, 24, 11, 15, 45, 40, 43, 13, 40, 11, 39, 11, 35, 21, 15, 32, 11, 14, 21, 23, 11, 11, 39, 7, 3, 26, 50, 38, 36, 47, 5, 0, 3, 15, 18, 33, 11, 4}
				for i := 0; i < 84; i += 2 {
					localKey := /*line AicWbk3.go:1*/ byte(i) + /*line AicWbk3.go:1*/ byte(positions[i]^positions[i+1]) + 136
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
				}
				return /*line AicWbk3.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.BlobHashes = otqLrvn1CD.BlobVersionedHashes

		var hoaGFM bool
		if otqLrvn1CD.R == nil {
			return /*line WISUUUW30z.go:1*/ errors.New(func() /*line QQxfx1z.go:1*/ string {
				data := [] /*line QQxfx1z.go:1*/ byte("\xe3\xb5\xf7s\x1fmg\x91r\xf2\xedu\x1c\xa9u\xf3\r\xc0\x9e\xa5Qd\xfb\xcar'\xe0\xfe\xf9\x8et\xa0\xc7P\f\xf0\xe1\x19\xcbon")
				positions := [...]byte{18, 16, 7, 31, 4, 2, 29, 12, 4, 33, 10, 29, 22, 13, 17, 32, 35, 37, 34, 28, 31, 20, 0, 14, 33, 29, 26, 19, 10, 27, 38, 31, 16, 1, 23, 34, 7, 37, 29, 36, 23, 5, 4, 33, 12, 0, 9, 15, 16, 28, 28, 23}
				for i := 0; i < 52; i += 2 {
					localKey := /*line QQxfx1z.go:1*/ byte(i) + /*line QQxfx1z.go:1*/ byte(positions[i]^positions[i+1]) + 98
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return /*line QQxfx1z.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.R, hoaGFM = /*line oEzDReaK9N.go:1*/ uint256.FromBig((* /*line dcj97ClHr.go:1*/ big.Int)(otqLrvn1CD.R))
		if hoaGFM {
			return /*line P5C2UKHy.go:1*/ errors.New(func() /*line AnC5z3V.go:1*/ string {
				seed := /*line AnC5z3V.go:1*/ byte(69)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line AnC5z3V.go:1*/ append(data, x+seed); seed += x; return fnc }
				/*line AnC5z3V.go:1*/ fnc(226)(75)(181)(249)(86)(235)(11)(9)(240)(187)(79)(7)(239)(13)(244)(6)(3)(8)(252)(173)(85)(244)(5)(6)(190)(3)(1)
				return /*line AnC5z3V.go:1*/ string(data)
			}())
		}

		if otqLrvn1CD.S == nil {
			return /*line NE4hsS8a.go:1*/ errors.New(func() /*line SsypIFaHtu8.go:1*/ string {
				data := /*line SsypIFaHtu8.go:1*/ make([]byte, 0, 42)
				i := 3
				decryptKey := 27
				for counter := 0; i != 10; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 9:
						i = 0
						data = /*line SsypIFaHtu8.go:1*/ append(data, "}uzt"...,
						)
					case 2:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\xa9\\X"...,
						)
						i = 12
					case 4:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\xa9\xa4\xa6"...,
						)
						i = 5
					case 16:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\x8b\x93"...,
						)
						i = 15
					case 14:
						i = 4
						data = /*line SsypIFaHtu8.go:1*/ append(data, "^\xa3"...,
						)
					case 1:
						data = /*line SsypIFaHtu8.go:1*/ append(data, 123)
						i = 6
					case 0:
						i = 10
						for y := range data {
							data[y] = data[y] + /*line SsypIFaHtu8.go:1*/ byte(decryptKey^y)
						}
					case 7:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "K\x98"...,
						)
						i = 8
					case 5:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\x9d\\b"...,
						)
						i = 2
					case 8:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\x8a\x99\x9c"...,
						)
						i = 16
					case 17:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\xa8\xa5o"...,
						)
						i = 1
					case 11:
						i = 7
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\xa2\x93\x97\x93"...,
						)
					case 15:
						i = 14
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\x89\x87"...,
						)
					case 13:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\x96\xa3"...,
						)
						i = 11
					case 12:
						i = 17
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\xa0\xa0Q"...,
						)
					case 6:
						data = /*line SsypIFaHtu8.go:1*/ append(data, "\x83pm"...,
						)
						i = 9
					case 3:
						i = 13
						data = /*line SsypIFaHtu8.go:1*/ append(data, 155)
					}
				}
				return /*line SsypIFaHtu8.go:1*/ string(data)
			}())
		}
		r5Uxg5_sE2.S, hoaGFM = /*line pL22wYV6F.go:1*/ uint256.FromBig((* /*line BVRg1Zm.go:1*/ big.Int)(otqLrvn1CD.S))
		if hoaGFM {
			return /*line EWmWDDOUoJ.go:1*/ errors.New(func() /*line KN3sbaQaL.go:1*/ string {
				key := [] /*line KN3sbaQaL.go:1*/ byte("\xc1Ƶ\teM\xa8*\xfb\x83\xb5\x7f8CsjE\xef\x7f\xd2\xff\xf9S\x94\xdc\xd0\x15")
				data := [] /*line KN3sbaQaL.go:1*/ byte("f\xadr\x17\x11\x14\xc4Kj\x9d\xba\xf7-/\xf3\x02*\x88\xf4Nvp\x1b\xe0Ve!")
				for i, b := range key {
					data[i] = data[i] + b
				}
				return /*line KN3sbaQaL.go:1*/ string(data)
			}())
		}

		c1obyP5J3_L, rfteMJju := /*line WpA92fwDp.go:1*/ otqLrvn1CD.jDvTUb()
		if rfteMJju != nil {
			return rfteMJju
		}
		r5Uxg5_sE2.V, hoaGFM = /*line S9rlEZFK.go:1*/ uint256.FromBig(c1obyP5J3_L)
		if hoaGFM {
			return /*line R71LK2qO1.go:1*/ errors.New(func() /*line HgD_Ct.go:1*/ string {
				fullData := [] /*line HgD_Ct.go:1*/ byte("P\x14\xfb\x96\x82=6\"\xdeIѵ\x1c\xcakK\xfaD\x9b\xb5UC\xd9\xd4?ٙI\x93\xd5\ai\xec;C})k\x11\x14\xf3j\x01\f<5ZN\xec7ePd&")
				idxKey := [] /*line HgD_Ct.go:1*/ byte("]\xfe?\xfc\xe2\x0fƻ\x14\x92i>\x8a\x9ev~")
				data := /*line HgD_Ct.go:1*/ make([]byte, 0, 28)
				data = /*line HgD_Ct.go:1*/ append(data, fullData[244^ /*line HgD_Ct.go:1*/ int(idxKey[4])]+fullData[205^ /*line HgD_Ct.go:1*/ int(idxKey[4])], fullData[39^ /*line HgD_Ct.go:1*/ int(idxKey[2])]+fullData[14^ /*line HgD_Ct.go:1*/ int(idxKey[2])], fullData[10^ /*line HgD_Ct.go:1*/ int(idxKey[2])]+fullData[21^ /*line HgD_Ct.go:1*/ int(idxKey[2])], fullData[213^ /*line HgD_Ct.go:1*/ int(idxKey[1])]+fullData[255^ /*line HgD_Ct.go:1*/ int(idxKey[1])], fullData[159^ /*line HgD_Ct.go:1*/ int(idxKey[12])]^fullData[167^ /*line HgD_Ct.go:1*/ int(idxKey[12])], fullData[33^ /*line HgD_Ct.go:1*/ int(idxKey[2])]+fullData[17^ /*line HgD_Ct.go:1*/ int(idxKey[2])], fullData[146^ /*line HgD_Ct.go:1*/ int(idxKey[13])]+fullData[173^ /*line HgD_Ct.go:1*/ int(idxKey[13])], fullData[100^ /*line HgD_Ct.go:1*/ int(idxKey[15])]^fullData[94^ /*line HgD_Ct.go:1*/ int(idxKey[15])], fullData[136^ /*line HgD_Ct.go:1*/ int(idxKey[12])]+fullData[163^ /*line HgD_Ct.go:1*/ int(idxKey[12])], fullData[156^ /*line HgD_Ct.go:1*/ int(idxKey[9])]+fullData[153^ /*line HgD_Ct.go:1*/ int(idxKey[9])], fullData[235^ /*line HgD_Ct.go:1*/ int(idxKey[3])]-fullData[206^ /*line HgD_Ct.go:1*/ int(idxKey[3])], fullData[149^ /*line HgD_Ct.go:1*/ int(idxKey[12])]-fullData[162^ /*line HgD_Ct.go:1*/ int(idxKey[12])], fullData[50^ /*line HgD_Ct.go:1*/ int(idxKey[2])]+fullData[45^ /*line HgD_Ct.go:1*/ int(idxKey[2])], fullData[221^ /*line HgD_Ct.go:1*/ int(idxKey[3])]^fullData[245^ /*line HgD_Ct.go:1*/ int(idxKey[3])], fullData[0^ /*line HgD_Ct.go:1*/ int(idxKey[8])]+fullData[50^ /*line HgD_Ct.go:1*/ int(idxKey[8])], fullData[168^ /*line HgD_Ct.go:1*/ int(idxKey[12])]+fullData[174^ /*line HgD_Ct.go:1*/ int(idxKey[12])], fullData[170^ /*line HgD_Ct.go:1*/ int(idxKey[7])]-fullData[166^ /*line HgD_Ct.go:1*/ int(idxKey[7])], fullData[226^ /*line HgD_Ct.go:1*/ int(idxKey[4])]-fullData[251^ /*line HgD_Ct.go:1*/ int(idxKey[4])], fullData[151^ /*line HgD_Ct.go:1*/ int(idxKey[9])]+fullData[148^ /*line HgD_Ct.go:1*/ int(idxKey[9])], fullData[26^ /*line HgD_Ct.go:1*/ int(idxKey[2])]-fullData[48^ /*line HgD_Ct.go:1*/ int(idxKey[2])], fullData[37^ /*line HgD_Ct.go:1*/ int(idxKey[11])]^fullData[18^ /*line HgD_Ct.go:1*/ int(idxKey[11])], fullData[47^ /*line HgD_Ct.go:1*/ int(idxKey[2])]^fullData[35^ /*line HgD_Ct.go:1*/ int(idxKey[2])], fullData[58^ /*line HgD_Ct.go:1*/ int(idxKey[11])]+fullData[14^ /*line HgD_Ct.go:1*/ int(idxKey[11])], fullData[97^ /*line HgD_Ct.go:1*/ int(idxKey[10])]+fullData[106^ /*line HgD_Ct.go:1*/ int(idxKey[10])], fullData[237^ /*line HgD_Ct.go:1*/ int(idxKey[1])]+fullData[221^ /*line HgD_Ct.go:1*/ int(idxKey[1])], fullData[152^ /*line HgD_Ct.go:1*/ int(idxKey[9])]+fullData[166^ /*line HgD_Ct.go:1*/ int(idxKey[9])], fullData[193^ /*line HgD_Ct.go:1*/ int(idxKey[6])]+fullData[225^ /*line HgD_Ct.go:1*/ int(idxKey[6])])
				return /*line HgD_Ct.go:1*/ string(data)
			}())
		}
		if /*line __d4yhY.go:1*/ r5Uxg5_sE2.V.Sign() != 0 || /*line W1avss7w.go:1*/ r5Uxg5_sE2.R.Sign() != 0 || /*line rX5JTT0_4I.go:1*/ r5Uxg5_sE2.S.Sign() != 0 {
			if rfteMJju := /*line akGmSQWR.go:1*/ vwXAfyu(c1obyP5J3_L /*line y8L2IE8Y.go:1*/, r5Uxg5_sE2.R.ToBig() /*line jv_AWdPaWyW.go:1*/, r5Uxg5_sE2.S.ToBig(), false); rfteMJju != nil {
				return rfteMJju
			}
		}

	default:
		return B0ZTT4Gv0t
	}

	/*line zavimbIYITh.go:1*/
	iPfjW8i91hC.secbe_rMaO(vv4w9F, 0)

	return nil
}

var _ = json.Compact
var _ = errors.As

const _ = big.Above

var _ = common.AbsolutePath
var _ hexutil.Big
var _ kzg4844.Blob
var _ = uint256.ErrBadBufferLength
