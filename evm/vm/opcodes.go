//line :1
package vm

import (
	"fmt"
)

type LxosJe8 byte

func (nCnJ0mc LxosJe8) IsPush() bool {
	return PUSH0 <= nCnJ0mc && nCnJ0mc <= PUSH32
}

const (
	STOP       LxosJe8 = 0x0
	ADD        LxosJe8 = 0x1
	MUL        LxosJe8 = 0x2
	SUB        LxosJe8 = 0x3
	DIV        LxosJe8 = 0x4
	SDIV       LxosJe8 = 0x5
	MOD        LxosJe8 = 0x6
	SMOD       LxosJe8 = 0x7
	ADDMOD     LxosJe8 = 0x8
	MULMOD     LxosJe8 = 0x9
	EXP        LxosJe8 = 0xa
	SIGNEXTEND LxosJe8 = 0xb
)

const (
	LT     LxosJe8 = 0x10
	GT     LxosJe8 = 0x11
	SLT    LxosJe8 = 0x12
	SGT    LxosJe8 = 0x13
	EQ     LxosJe8 = 0x14
	ISZERO LxosJe8 = 0x15
	AND    LxosJe8 = 0x16
	OR     LxosJe8 = 0x17
	XOR    LxosJe8 = 0x18
	NOT    LxosJe8 = 0x19
	BYTE   LxosJe8 = 0x1a
	SHL    LxosJe8 = 0x1b
	SHR    LxosJe8 = 0x1c
	SAR    LxosJe8 = 0x1d
)

const (
	KECCAK256 LxosJe8 = 0x20
)

const (
	ADDRESS        LxosJe8 = 0x30
	BALANCE        LxosJe8 = 0x31
	ORIGIN         LxosJe8 = 0x32
	CALLER         LxosJe8 = 0x33
	CALLVALUE      LxosJe8 = 0x34
	CALLDATALOAD   LxosJe8 = 0x35
	CALLDATASIZE   LxosJe8 = 0x36
	CALLDATACOPY   LxosJe8 = 0x37
	CODESIZE       LxosJe8 = 0x38
	CODECOPY       LxosJe8 = 0x39
	GASPRICE       LxosJe8 = 0x3a
	EXTCODESIZE    LxosJe8 = 0x3b
	EXTCODECOPY    LxosJe8 = 0x3c
	RETURNDATASIZE LxosJe8 = 0x3d
	RETURNDATACOPY LxosJe8 = 0x3e
	EXTCODEHASH    LxosJe8 = 0x3f
)

const (
	BLOCKHASH   LxosJe8 = 0x40
	COINBASE    LxosJe8 = 0x41
	TIMESTAMP   LxosJe8 = 0x42
	NUMBER      LxosJe8 = 0x43
	DIFFICULTY  LxosJe8 = 0x44
	RANDOM      LxosJe8 = 0x44
	PREVRANDAO  LxosJe8 = 0x44
	GASLIMIT    LxosJe8 = 0x45
	CHAINID     LxosJe8 = 0x46
	SELFBALANCE LxosJe8 = 0x47
	BASEFEE     LxosJe8 = 0x48
	BLOBHASH    LxosJe8 = 0x49
	BLOBBASEFEE LxosJe8 = 0x4a
)

const (
	POP      LxosJe8 = 0x50
	MLOAD    LxosJe8 = 0x51
	MSTORE   LxosJe8 = 0x52
	MSTORE8  LxosJe8 = 0x53
	SLOAD    LxosJe8 = 0x54
	SSTORE   LxosJe8 = 0x55
	JUMP     LxosJe8 = 0x56
	JUMPI    LxosJe8 = 0x57
	PC       LxosJe8 = 0x58
	MSIZE    LxosJe8 = 0x59
	GAS      LxosJe8 = 0x5a
	JUMPDEST LxosJe8 = 0x5b
	TLOAD    LxosJe8 = 0x5c
	TSTORE   LxosJe8 = 0x5d
	MCOPY    LxosJe8 = 0x5e
	PUSH0    LxosJe8 = 0x5f
)

const (
	PUSH1 LxosJe8 = 0x60 + iota
	PUSH2
	PUSH3
	PUSH4
	PUSH5
	PUSH6
	PUSH7
	PUSH8
	PUSH9
	PUSH10
	PUSH11
	PUSH12
	PUSH13
	PUSH14
	PUSH15
	PUSH16
	PUSH17
	PUSH18
	PUSH19
	PUSH20
	PUSH21
	PUSH22
	PUSH23
	PUSH24
	PUSH25
	PUSH26
	PUSH27
	PUSH28
	PUSH29
	PUSH30
	PUSH31
	PUSH32
)

const (
	DUP1 = 0x80 + iota
	DUP2
	DUP3
	DUP4
	DUP5
	DUP6
	DUP7
	DUP8
	DUP9
	DUP10
	DUP11
	DUP12
	DUP13
	DUP14
	DUP15
	DUP16
)

const (
	SWAP1 = 0x90 + iota
	SWAP2
	SWAP3
	SWAP4
	SWAP5
	SWAP6
	SWAP7
	SWAP8
	SWAP9
	SWAP10
	SWAP11
	SWAP12
	SWAP13
	SWAP14
	SWAP15
	SWAP16
)

const (
	LOG0 LxosJe8 = 0xa0 + iota
	LOG1
	LOG2
	LOG3
	LOG4
)

const (
	CREATE       LxosJe8 = 0xf0
	CALL         LxosJe8 = 0xf1
	CALLCODE     LxosJe8 = 0xf2
	RETURN       LxosJe8 = 0xf3
	DELEGATECALL LxosJe8 = 0xf4
	CREATE2      LxosJe8 = 0xf5

	STATICCALL   LxosJe8 = 0xfa
	REVERT       LxosJe8 = 0xfd
	INVALID      LxosJe8 = 0xfe
	SELFDESTRUCT LxosJe8 = 0xff
)

var pQzNDaP5fC = [256]string{

	STOP:   "STOP",
	ADD:    "ADD",
	MUL:    "MUL",
	SUB:    "SUB",
	DIV:    "DIV",
	SDIV:   "SDIV",
	MOD:    "MOD",
	SMOD:   "SMOD",
	EXP:    "EXP",
	NOT:    "NOT",
	LT:     "LT",
	GT:     "GT",
	SLT:    "SLT",
	SGT:    "SGT",
	EQ:     "EQ",
	ISZERO: "ISZERO",
	SIGNEXTEND: func() /*line h8spl7pKy0E.go:1*/ string {
		key := [] /*line h8spl7pKy0E.go:1*/ byte("\xe40\x84\"ɚf\xbcK\x93")
		data := [] /*line h8spl7pKy0E.go:1*/ byte("\xb7y\xc3l\x8c\xc22\xf9\x05\xd7")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line h8spl7pKy0E.go:1*/ string(data)
	}(),

	AND:    "AND",
	OR:     "OR",
	XOR:    "XOR",
	BYTE:   "BYTE",
	SHL:    "SHL",
	SHR:    "SHR",
	SAR:    "SAR",
	ADDMOD: "ADDMOD",
	MULMOD: "MULMOD",

	KECCAK256: func() /*line bX9YIKziTZ_.go:1*/ string {
		seed := /*line bX9YIKziTZ_.go:1*/ byte(213)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line bX9YIKziTZ_.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line bX9YIKziTZ_.go:1*/ fnc(118)(250)(254)(0)(254)(10)(231)(3)(1)
		return /*line bX9YIKziTZ_.go:1*/ string(data)
	}(),

	ADDRESS: "ADDRESS",
	BALANCE: "BALANCE",
	ORIGIN:  "ORIGIN",
	CALLER:  "CALLER",
	CALLVALUE: func() /*line EG1Xm7ZPH7s3.go:1*/ string {
		key := [] /*line EG1Xm7ZPH7s3.go:1*/ byte("HM\xd0\x00\x1c\xca\x1fO\t")
		data := [] /*line EG1Xm7ZPH7s3.go:1*/ byte("\v\f\x9cLJ\x8bS\x1aL")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line EG1Xm7ZPH7s3.go:1*/ string(data)
	}(),
	CALLDATALOAD: func() /*line P8EraLd2pFYd.go:1*/ string {
		seed := /*line P8EraLd2pFYd.go:1*/ byte(78)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line P8EraLd2pFYd.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line P8EraLd2pFYd.go:1*/ fnc(13)(26)(57)(226)(212)(37)(221)(39)(193)(1)(14)(25)
		return /*line P8EraLd2pFYd.go:1*/ string(data)
	}(),
	CALLDATASIZE: func() /*line uot4w7i.go:1*/ string {
		data := /*line uot4w7i.go:1*/ make([]byte, 0, 13)
		i := 1
		decryptKey := 187
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				data = /*line uot4w7i.go:1*/ append(data, "\xf7\xf4\xfa\xfb"...,
				)
				i = 5
			case 6:
				i = 4
				data = /*line uot4w7i.go:1*/ append(data, 250)
			case 5:
				i = 3
				data = /*line uot4w7i.go:1*/ append(data, "\xf4\xf0"...,
				)
			case 0:
				i = 6
				data = /*line uot4w7i.go:1*/ append(data, "\xf4\xe4"...,
				)
			case 4:
				i = 2
				for y := range data {
					data[y] = data[y] ^ /*line uot4w7i.go:1*/ byte(decryptKey^y)
				}
			case 3:
				data = /*line uot4w7i.go:1*/ append(data, "\xe6\xf2\xef"...,
				)
				i = 0
			}
		}
		return /*line uot4w7i.go:1*/ string(data)
	}(),
	CALLDATACOPY: func() /*line hvpAasInCm.go:1*/ string {
		key := [] /*line hvpAasInCm.go:1*/ byte("H\xc5\xe2H;\xa7\x80\b0Ő")
		data := [] /*line hvpAasInCm.go:1*/ byte("\v\x84\xae\x04\x7f\xe6\xd4Is\x00\x9c\xd2")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line hvpAasInCm.go:1*/ string(data)
	}(),
	CODESIZE: func() /*line OxeBv80.go:1*/ string {
		data := /*line OxeBv80.go:1*/ make([]byte, 0, 9)
		i := 0
		decryptKey := 220
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				data = /*line OxeBv80.go:1*/ append(data, 10)
				i = 3
			case 8:
				data = /*line OxeBv80.go:1*/ append(data, 12)
				i = 9
			case 0:
				i = 7
				data = /*line OxeBv80.go:1*/ append(data, 7)
			case 4:
				i = 6
				data = /*line OxeBv80.go:1*/ append(data, 8)
			case 2:
				data = /*line OxeBv80.go:1*/ append(data, 10)
				i = 8
			case 7:
				i = 2
				data = /*line OxeBv80.go:1*/ append(data, 20)
			case 6:
				for y := range data {
					data[y] = data[y] - /*line OxeBv80.go:1*/ byte(decryptKey^y)
				}
				i = 1
			case 3:
				data = /*line OxeBv80.go:1*/ append(data, 28)
				i = 4
			case 9:
				data = /*line OxeBv80.go:1*/ append(data, 19)
				i = 5
			}
		}
		return /*line OxeBv80.go:1*/ string(data)
	}(),
	CODECOPY: func() /*line pURwjUH_VGj.go:1*/ string {
		fullData := [] /*line pURwjUH_VGj.go:1*/ byte("t\xe5\xf1}o\xd6\"Ƅ\xaa\xe2^nfݷ")
		idxKey := [] /*line pURwjUH_VGj.go:1*/ byte("\xd9\x0e\xbe\xcc")
		data := /*line pURwjUH_VGj.go:1*/ make([]byte, 0, 9)
		data = /*line pURwjUH_VGj.go:1*/ append(data, fullData[207^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])]+fullData[203^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])], fullData[183^ /*line pURwjUH_VGj.go:1*/ int(idxKey[2])]^fullData[191^ /*line pURwjUH_VGj.go:1*/ int(idxKey[2])], fullData[202^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])]^fullData[193^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])], fullData[200^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])]+fullData[201^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])], fullData[195^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])]-fullData[204^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])], fullData[206^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])]+fullData[199^ /*line pURwjUH_VGj.go:1*/ int(idxKey[3])], fullData[2^ /*line pURwjUH_VGj.go:1*/ int(idxKey[1])]+fullData[4^ /*line pURwjUH_VGj.go:1*/ int(idxKey[1])], fullData[6^ /*line pURwjUH_VGj.go:1*/ int(idxKey[1])]^fullData[0^ /*line pURwjUH_VGj.go:1*/ int(idxKey[1])])
		return /*line pURwjUH_VGj.go:1*/ string(data)
	}(),
	GASPRICE: func() /*line g9iVnV37.go:1*/ string {
		key := [] /*line g9iVnV37.go:1*/ byte("\xdb\xce\xfeL\xee\xfbN~")
		data := [] /*line g9iVnV37.go:1*/ byte("lsU\x04dN\xf5\xc7")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line g9iVnV37.go:1*/ string(data)
	}(),
	EXTCODESIZE: func() /*line KPGy1IteBtO1.go:1*/ string {
		data := [] /*line KPGy1IteBtO1.go:1*/ byte("hnwQgSTScI^")
		positions := [...]byte{6, 0, 9, 10, 4, 5, 2, 8, 1, 8, 3, 0, 6, 2, 8, 5}
		for i := 0; i < 16; i += 2 {
			localKey := /*line KPGy1IteBtO1.go:1*/ byte(i) + /*line KPGy1IteBtO1.go:1*/ byte(positions[i]^positions[i+1]) + 255
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line KPGy1IteBtO1.go:1*/ string(data)
	}(),
	EXTCODECOPY: func() /*line rOV9OZTtm.go:1*/ string {
		key := [] /*line rOV9OZTtm.go:1*/ byte("\xcd\xcd@\xe1\f\xd9.\x06\xfba]")
		data := [] /*line rOV9OZTtm.go:1*/ byte("x\x8b\x14bCk\x17=T\xef\xfc")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line rOV9OZTtm.go:1*/ string(data)
	}(),
	RETURNDATASIZE: func() /*line YZJargiByrHT.go:1*/ string {
		data := /*line YZJargiByrHT.go:1*/ make([]byte, 0, 15)
		i := 3
		decryptKey := 107
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data = /*line YZJargiByrHT.go:1*/ append(data, 45)
				i = 6
			case 6:
				data = /*line YZJargiByrHT.go:1*/ append(data, ";)))"...,
				)
				i = 0
			case 0:
				i = 4
				data = /*line YZJargiByrHT.go:1*/ append(data, "4=9#"...,
				)
			case 5:
				i = 2
				data = /*line YZJargiByrHT.go:1*/ append(data, 55)
			case 4:
				data = /*line YZJargiByrHT.go:1*/ append(data, "7&=)"...,
				)
				i = 5
			case 2:
				i = 1
				for y := range data {
					data[y] = data[y] ^ /*line YZJargiByrHT.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return /*line YZJargiByrHT.go:1*/ string(data)
	}(),
	RETURNDATACOPY: func() /*line AyYNb30klaUP.go:1*/ string {
		key := [] /*line AyYNb30klaUP.go:1*/ byte("X\xec\xe1\x11B\xb1\x95t\xaey\x03\xe4\xce,")
		data := [] /*line AyYNb30klaUP.go:1*/ byte("\n\xa9\xb5D\x10\xff\xd15\xfa8@\xab\x9eu")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line AyYNb30klaUP.go:1*/ string(data)
	}(),
	EXTCODEHASH: func() /*line k9X54WoFrt.go:1*/ string {
		seed := /*line k9X54WoFrt.go:1*/ byte(157)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line k9X54WoFrt.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line k9X54WoFrt.go:1*/ fnc(216)(45)(246)(219)(60)(235)(223)(49)(235)(198)(19)
		return /*line k9X54WoFrt.go:1*/ string(data)
	}(),

	BLOCKHASH: func() /*line CrvpUaZ.go:1*/ string {
		fullData := [] /*line CrvpUaZ.go:1*/ byte("\x17\bB\xf2\x15\xd3\xfbX\x91\xb6ͱ\x06\xd4\xf9C~\xce")
		idxKey := [] /*line CrvpUaZ.go:1*/ byte("\xa3\xa8")
		data := /*line CrvpUaZ.go:1*/ make([]byte, 0, 10)
		data = /*line CrvpUaZ.go:1*/ append(data, fullData[173^ /*line CrvpUaZ.go:1*/ int(idxKey[1])]-fullData[160^ /*line CrvpUaZ.go:1*/ int(idxKey[1])], fullData[179^ /*line CrvpUaZ.go:1*/ int(idxKey[0])]+fullData[178^ /*line CrvpUaZ.go:1*/ int(idxKey[0])], fullData[170^ /*line CrvpUaZ.go:1*/ int(idxKey[0])]^fullData[173^ /*line CrvpUaZ.go:1*/ int(idxKey[0])], fullData[163^ /*line CrvpUaZ.go:1*/ int(idxKey[0])]-fullData[174^ /*line CrvpUaZ.go:1*/ int(idxKey[0])], fullData[169^ /*line CrvpUaZ.go:1*/ int(idxKey[1])]^fullData[167^ /*line CrvpUaZ.go:1*/ int(idxKey[1])], fullData[172^ /*line CrvpUaZ.go:1*/ int(idxKey[1])]-fullData[162^ /*line CrvpUaZ.go:1*/ int(idxKey[1])], fullData[160^ /*line CrvpUaZ.go:1*/ int(idxKey[0])]-fullData[168^ /*line CrvpUaZ.go:1*/ int(idxKey[0])], fullData[165^ /*line CrvpUaZ.go:1*/ int(idxKey[0])]+fullData[164^ /*line CrvpUaZ.go:1*/ int(idxKey[0])], fullData[164^ /*line CrvpUaZ.go:1*/ int(idxKey[1])]+fullData[170^ /*line CrvpUaZ.go:1*/ int(idxKey[1])])
		return /*line CrvpUaZ.go:1*/ string(data)
	}(),
	COINBASE: func() /*line FDT173VjV.go:1*/ string {
		key := [] /*line FDT173VjV.go:1*/ byte("\x95:0\xfc)fչ")
		data := [] /*line FDT173VjV.go:1*/ byte("\xae\x15\x19R\x19\xdb~\x8c")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line FDT173VjV.go:1*/ string(data)
	}(),
	TIMESTAMP: func() /*line I0xOEBV.go:1*/ string {
		seed := /*line I0xOEBV.go:1*/ byte(161)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line I0xOEBV.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line I0xOEBV.go:1*/ fnc(245)(223)(56)(232)(198)(15)(43)(216)(61)
		return /*line I0xOEBV.go:1*/ string(data)
	}(),
	NUMBER: "NUMBER",
	DIFFICULTY: func() /*line WQvPNGX.go:1*/ string {
		data := /*line WQvPNGX.go:1*/ make([]byte, 0, 11)
		i := 1
		decryptKey := 249
		for counter := 0; i != 11; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data = /*line WQvPNGX.go:1*/ append(data, 136)
				i = 6
			case 6:
				data = /*line WQvPNGX.go:1*/ append(data, 132)
				i = 5
			case 2:
				i = 0
				data = /*line WQvPNGX.go:1*/ append(data, 147)
			case 10:
				i = 11
				for y := range data {
					data[y] = data[y] ^ /*line WQvPNGX.go:1*/ byte(decryptKey^y)
				}
			case 7:
				data = /*line WQvPNGX.go:1*/ append(data, 156)
				i = 9
			case 1:
				i = 3
				data = /*line WQvPNGX.go:1*/ append(data, 132)
			case 8:
				i = 2
				data = /*line WQvPNGX.go:1*/ append(data, 134)
			case 0:
				data = /*line WQvPNGX.go:1*/ append(data, 139)
				i = 7
			case 4:
				data = /*line WQvPNGX.go:1*/ append(data, 141)
				i = 8
			case 5:
				i = 4
				data = /*line WQvPNGX.go:1*/ append(data, 133)
			case 9:
				data = /*line WQvPNGX.go:1*/ append(data, 144)
				i = 10
			}
		}
		return /*line WQvPNGX.go:1*/ string(data)
	}(),
	GASLIMIT: func() /*line uJtCr1Rj.go:1*/ string {
		data := [] /*line uJtCr1Rj.go:1*/ byte("[pbLWMOR")
		positions := [...]byte{6, 7, 0, 0, 4, 4, 2, 1}
		for i := 0; i < 8; i += 2 {
			localKey := /*line uJtCr1Rj.go:1*/ byte(i) + /*line uJtCr1Rj.go:1*/ byte(positions[i]^positions[i+1]) + 26
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line uJtCr1Rj.go:1*/ string(data)
	}(),
	CHAINID: "CHAINID",
	SELFBALANCE: func() /*line CAP6bF7zSmi.go:1*/ string {
		data := [] /*line CAP6bF7zSmi.go:1*/ byte("\xf1\xf85\xfe\xf7\x80LAY\xfeE")
		positions := [...]byte{8, 4, 4, 0, 2, 1, 1, 9, 3, 4, 5, 9, 5, 3}
		for i := 0; i < 14; i += 2 {
			localKey := /*line CAP6bF7zSmi.go:1*/ byte(i) + /*line CAP6bF7zSmi.go:1*/ byte(positions[i]^positions[i+1]) + 173
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line CAP6bF7zSmi.go:1*/ string(data)
	}(),
	BASEFEE: "BASEFEE",
	BLOBHASH: func() /*line yzcIJRRNcQ.go:1*/ string {
		data := [] /*line yzcIJRRNcQ.go:1*/ byte("5EFB6A\x17;")
		positions := [...]byte{1, 6, 2, 2, 1, 7, 0, 0, 4, 7}
		for i := 0; i < 10; i += 2 {
			localKey := /*line yzcIJRRNcQ.go:1*/ byte(i) + /*line yzcIJRRNcQ.go:1*/ byte(positions[i]^positions[i+1]) + 7
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line yzcIJRRNcQ.go:1*/ string(data)
	}(),
	BLOBBASEFEE: func() /*line otIYL7P0ymH.go:1*/ string {
		data := [] /*line otIYL7P0ymH.go:1*/ byte("\xbfu\x84\xa9~\xb1S\x9dy|\x87")
		positions := [...]byte{2, 0, 3, 8, 2, 10, 4, 5, 8, 7, 4, 8, 9, 1}
		for i := 0; i < 14; i += 2 {
			localKey := /*line otIYL7P0ymH.go:1*/ byte(i) + /*line otIYL7P0ymH.go:1*/ byte(positions[i]^positions[i+1]) + 188
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line otIYL7P0ymH.go:1*/ string(data)
	}(),

	POP:     "POP",
	MLOAD:   "MLOAD",
	MSTORE:  "MSTORE",
	MSTORE8: "MSTORE8",
	SLOAD:   "SLOAD",
	SSTORE:  "SSTORE",
	JUMP:    "JUMP",
	JUMPI:   "JUMPI",
	PC:      "PC",
	MSIZE:   "MSIZE",
	GAS:     "GAS",
	JUMPDEST: func() /*line IY3ZXuvq.go:1*/ string {
		seed := /*line IY3ZXuvq.go:1*/ byte(54)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line IY3ZXuvq.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line IY3ZXuvq.go:1*/ fnc(124)(231)(212)(61)(238)(221)(38)(207)
		return /*line IY3ZXuvq.go:1*/ string(data)
	}(),
	TLOAD:  "TLOAD",
	TSTORE: "TSTORE",
	MCOPY:  "MCOPY",
	PUSH0:  "PUSH0",

	PUSH1:  "PUSH1",
	PUSH2:  "PUSH2",
	PUSH3:  "PUSH3",
	PUSH4:  "PUSH4",
	PUSH5:  "PUSH5",
	PUSH6:  "PUSH6",
	PUSH7:  "PUSH7",
	PUSH8:  "PUSH8",
	PUSH9:  "PUSH9",
	PUSH10: "PUSH10",
	PUSH11: "PUSH11",
	PUSH12: "PUSH12",
	PUSH13: "PUSH13",
	PUSH14: "PUSH14",
	PUSH15: "PUSH15",
	PUSH16: "PUSH16",
	PUSH17: "PUSH17",
	PUSH18: "PUSH18",
	PUSH19: "PUSH19",
	PUSH20: "PUSH20",
	PUSH21: "PUSH21",
	PUSH22: "PUSH22",
	PUSH23: "PUSH23",
	PUSH24: "PUSH24",
	PUSH25: "PUSH25",
	PUSH26: "PUSH26",
	PUSH27: "PUSH27",
	PUSH28: "PUSH28",
	PUSH29: "PUSH29",
	PUSH30: "PUSH30",
	PUSH31: "PUSH31",
	PUSH32: "PUSH32",

	DUP1:  "DUP1",
	DUP2:  "DUP2",
	DUP3:  "DUP3",
	DUP4:  "DUP4",
	DUP5:  "DUP5",
	DUP6:  "DUP6",
	DUP7:  "DUP7",
	DUP8:  "DUP8",
	DUP9:  "DUP9",
	DUP10: "DUP10",
	DUP11: "DUP11",
	DUP12: "DUP12",
	DUP13: "DUP13",
	DUP14: "DUP14",
	DUP15: "DUP15",
	DUP16: "DUP16",

	SWAP1:  "SWAP1",
	SWAP2:  "SWAP2",
	SWAP3:  "SWAP3",
	SWAP4:  "SWAP4",
	SWAP5:  "SWAP5",
	SWAP6:  "SWAP6",
	SWAP7:  "SWAP7",
	SWAP8:  "SWAP8",
	SWAP9:  "SWAP9",
	SWAP10: "SWAP10",
	SWAP11: "SWAP11",
	SWAP12: "SWAP12",
	SWAP13: "SWAP13",
	SWAP14: "SWAP14",
	SWAP15: "SWAP15",
	SWAP16: "SWAP16",

	LOG0: "LOG0",
	LOG1: "LOG1",
	LOG2: "LOG2",
	LOG3: "LOG3",
	LOG4: "LOG4",

	CREATE: "CREATE",
	CALL:   "CALL",
	RETURN: "RETURN",
	CALLCODE: func() /*line Zx_Q9adRaIT5.go:1*/ string {
		data := /*line Zx_Q9adRaIT5.go:1*/ make([]byte, 0, 9)
		i := 5
		decryptKey := 69
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 3
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 117)
			case 5:
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 120)
				i = 6
			case 3:
				i = 4
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 116)
			case 0:
				i = 2
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 121)
			case 6:
				i = 1
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 123)
			case 4:
				i = 9
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 124)
			case 2:
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 121)
				i = 8
			case 9:
				i = 0
				data = /*line Zx_Q9adRaIT5.go:1*/ append(data, 113)
			case 8:
				for y := range data {
					data[y] = data[y] ^ /*line Zx_Q9adRaIT5.go:1*/ byte(decryptKey^y)
				}
				i = 7
			}
		}
		return /*line Zx_Q9adRaIT5.go:1*/ string(data)
	}(),
	DELEGATECALL: func() /*line TwYO3l.go:1*/ string {
		fullData := [] /*line TwYO3l.go:1*/ byte("}E%\xe4\x89r\xed\xc93(\x85ݹ\xd3\x031\x11\x98P\x127{\xe1@")
		idxKey := [] /*line TwYO3l.go:1*/ byte("\x80\xe2")
		data := /*line TwYO3l.go:1*/ make([]byte, 0, 13)
		data = /*line TwYO3l.go:1*/ append(data, fullData[132^ /*line TwYO3l.go:1*/ int(idxKey[0])]-fullData[129^ /*line TwYO3l.go:1*/ int(idxKey[0])], fullData[239^ /*line TwYO3l.go:1*/ int(idxKey[1])]+fullData[231^ /*line TwYO3l.go:1*/ int(idxKey[1])], fullData[237^ /*line TwYO3l.go:1*/ int(idxKey[1])]^fullData[226^ /*line TwYO3l.go:1*/ int(idxKey[1])], fullData[243^ /*line TwYO3l.go:1*/ int(idxKey[1])]^fullData[233^ /*line TwYO3l.go:1*/ int(idxKey[1])], fullData[137^ /*line TwYO3l.go:1*/ int(idxKey[0])]-fullData[150^ /*line TwYO3l.go:1*/ int(idxKey[0])], fullData[240^ /*line TwYO3l.go:1*/ int(idxKey[1])]^fullData[242^ /*line TwYO3l.go:1*/ int(idxKey[1])], fullData[228^ /*line TwYO3l.go:1*/ int(idxKey[1])]^fullData[238^ /*line TwYO3l.go:1*/ int(idxKey[1])], fullData[147^ /*line TwYO3l.go:1*/ int(idxKey[0])]+fullData[136^ /*line TwYO3l.go:1*/ int(idxKey[0])], fullData[236^ /*line TwYO3l.go:1*/ int(idxKey[1])]+fullData[245^ /*line TwYO3l.go:1*/ int(idxKey[1])], fullData[130^ /*line TwYO3l.go:1*/ int(idxKey[0])]-fullData[131^ /*line TwYO3l.go:1*/ int(idxKey[0])], fullData[135^ /*line TwYO3l.go:1*/ int(idxKey[0])]^fullData[138^ /*line TwYO3l.go:1*/ int(idxKey[0])], fullData[148^ /*line TwYO3l.go:1*/ int(idxKey[0])]^fullData[149^ /*line TwYO3l.go:1*/ int(idxKey[0])])
		return /*line TwYO3l.go:1*/ string(data)
	}(),
	CREATE2: "CREATE2",
	STATICCALL: func() /*line XuCDqP1.go:1*/ string {
		seed := /*line XuCDqP1.go:1*/ byte(41)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line XuCDqP1.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line XuCDqP1.go:1*/ fnc(42)(1)(237)(19)(245)(250)(0)(254)(11)(0)
		return /*line XuCDqP1.go:1*/ string(data)
	}(),
	REVERT:  "REVERT",
	INVALID: "INVALID",
	SELFDESTRUCT: func() /*line G8wsDWYC_.go:1*/ string {
		seed := /*line G8wsDWYC_.go:1*/ byte(241)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line G8wsDWYC_.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line G8wsDWYC_.go:1*/ fnc(98)(242)(7)(250)(254)(1)(14)(1)(254)(3)(238)(17)
		return /*line G8wsDWYC_.go:1*/ string(data)
	}(),
}

func (nCnJ0mc LxosJe8) String() string {
	if ek3u6QMaV := pQzNDaP5fC[nCnJ0mc]; ek3u6QMaV != "" {
		return ek3u6QMaV
	}
	return /*line rf0W7ttveMuD.go:1*/ fmt.Sprintf(func() /*line NSWR7QyOH.go:1*/ string {
		fullData := [] /*line NSWR7QyOH.go:1*/ byte("$\x11\x02\xa6\x1c\x82\x8c\xed\x04p\xb5\xc0w\xabo\xb9\x01<_\xe1<TM\xa9m3W@(|?8\xcaթ\xce2\\\xa1\x18ǖ\xffW")
		idxKey := [] /*line NSWR7QyOH.go:1*/ byte("\x15\xa6\x95\x1c,\xa5")
		data := /*line NSWR7QyOH.go:1*/ make([]byte, 0, 23)
		data = /*line NSWR7QyOH.go:1*/ append(data, fullData[190^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])]^fullData[164^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])], fullData[12^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]+fullData[18^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[133^ /*line NSWR7QyOH.go:1*/ int(idxKey[5])]^fullData[178^ /*line NSWR7QyOH.go:1*/ int(idxKey[5])], fullData[191^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])]^fullData[131^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])], fullData[16^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]+fullData[27^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[182^ /*line NSWR7QyOH.go:1*/ int(idxKey[5])]-fullData[184^ /*line NSWR7QyOH.go:1*/ int(idxKey[5])], fullData[162^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])]^fullData[183^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])], fullData[63^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]+fullData[6^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[51^ /*line NSWR7QyOH.go:1*/ int(idxKey[0])]+fullData[16^ /*line NSWR7QyOH.go:1*/ int(idxKey[0])], fullData[58^ /*line NSWR7QyOH.go:1*/ int(idxKey[4])]-fullData[13^ /*line NSWR7QyOH.go:1*/ int(idxKey[4])], fullData[183^ /*line NSWR7QyOH.go:1*/ int(idxKey[5])]-fullData[187^ /*line NSWR7QyOH.go:1*/ int(idxKey[5])], fullData[169^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])]+fullData[172^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])], fullData[17^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]-fullData[8^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[157^ /*line NSWR7QyOH.go:1*/ int(idxKey[2])]+fullData[156^ /*line NSWR7QyOH.go:1*/ int(idxKey[2])], fullData[51^ /*line NSWR7QyOH.go:1*/ int(idxKey[4])]^fullData[11^ /*line NSWR7QyOH.go:1*/ int(idxKey[4])], fullData[7^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]^fullData[28^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[20^ /*line NSWR7QyOH.go:1*/ int(idxKey[0])]+fullData[0^ /*line NSWR7QyOH.go:1*/ int(idxKey[0])], fullData[23^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]+fullData[31^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[140^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])]^fullData[143^ /*line NSWR7QyOH.go:1*/ int(idxKey[1])], fullData[62^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]^fullData[52^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[55^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])]^fullData[56^ /*line NSWR7QyOH.go:1*/ int(idxKey[3])], fullData[42^ /*line NSWR7QyOH.go:1*/ int(idxKey[4])]-fullData[48^ /*line NSWR7QyOH.go:1*/ int(idxKey[4])])
		return /*line NSWR7QyOH.go:1*/ string(data)
	}(), /*line Cf4KNNZ.go:1*/ int(nCnJ0mc))
}

var coCamfmb = map[string]LxosJe8{
	"STOP":   STOP,
	"ADD":    ADD,
	"MUL":    MUL,
	"SUB":    SUB,
	"DIV":    DIV,
	"SDIV":   SDIV,
	"MOD":    MOD,
	"SMOD":   SMOD,
	"EXP":    EXP,
	"NOT":    NOT,
	"LT":     LT,
	"GT":     GT,
	"SLT":    SLT,
	"SGT":    SGT,
	"EQ":     EQ,
	"ISZERO": ISZERO,
	func() /*line TX8Wqto61G.go:1*/ string {
		data := /*line TX8Wqto61G.go:1*/ make([]byte, 0, 11)
		i := 5
		decryptKey := 31
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				data = /*line TX8Wqto61G.go:1*/ append(data, 92)
				i = 6
			case 10:
				data = /*line TX8Wqto61G.go:1*/ append(data, 87)
				i = 0
			case 5:
				data = /*line TX8Wqto61G.go:1*/ append(data, 66)
				i = 2
			case 1:
				i = 11
				data = /*line TX8Wqto61G.go:1*/ append(data, 67)
			case 6:
				for y := range data {
					data[y] = data[y] ^ /*line TX8Wqto61G.go:1*/ byte(decryptKey^y)
				}
				i = 3
			case 7:
				i = 8
				data = /*line TX8Wqto61G.go:1*/ append(data, 84)
			case 8:
				data = /*line TX8Wqto61G.go:1*/ append(data, 92)
				i = 4
			case 9:
				i = 1
				data = /*line TX8Wqto61G.go:1*/ append(data, 76)
			case 2:
				data = /*line TX8Wqto61G.go:1*/ append(data, 89)
				i = 7
			case 11:
				data = /*line TX8Wqto61G.go:1*/ append(data, 83)
				i = 10
			case 4:
				data = /*line TX8Wqto61G.go:1*/ append(data, 80)
				i = 9
			}
		}
		return /*line TX8Wqto61G.go:1*/ string(data)
	}(): SIGNEXTEND,
	"AND":    AND,
	"OR":     OR,
	"XOR":    XOR,
	"BYTE":   BYTE,
	"SHL":    SHL,
	"SHR":    SHR,
	"SAR":    SAR,
	"ADDMOD": ADDMOD,
	"MULMOD": MULMOD,
	func() /*line j_WuYklCG.go:1*/ string {
		data := [] /*line j_WuYklCG.go:1*/ byte("\xb3E\xb2\xba\xb1\xbb256")
		positions := [...]byte{4, 2, 0, 4, 3, 3, 5, 3, 5, 0}
		for i := 0; i < 10; i += 2 {
			localKey := /*line j_WuYklCG.go:1*/ byte(i) + /*line j_WuYklCG.go:1*/ byte(positions[i]^positions[i+1]) + 236
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line j_WuYklCG.go:1*/ string(data)
	}(): KECCAK256,
	"ADDRESS": ADDRESS,
	"BALANCE": BALANCE,
	"ORIGIN":  ORIGIN,
	"CALLER":  CALLER,
	func() /*line tQSyFgc.go:1*/ string {
		data := /*line tQSyFgc.go:1*/ make([]byte, 0, 10)
		i := 6
		decryptKey := 147
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 4
				data = /*line tQSyFgc.go:1*/ append(data, 174)
			case 1:
				data = /*line tQSyFgc.go:1*/ append(data, 181)
				i = 3
			case 6:
				i = 7
				data = /*line tQSyFgc.go:1*/ append(data, 171)
			case 7:
				i = 5
				data = /*line tQSyFgc.go:1*/ append(data, 168)
			case 3:
				i = 2
				for y := range data {
					data[y] = data[y] + /*line tQSyFgc.go:1*/ byte(decryptKey^y)
				}
			case 9:
				data = /*line tQSyFgc.go:1*/ append(data, 177)
				i = 8
			case 10:
				i = 0
				data = /*line tQSyFgc.go:1*/ append(data, 164)
			case 8:
				data = /*line tQSyFgc.go:1*/ append(data, 186)
				i = 10
			case 5:
				data = /*line tQSyFgc.go:1*/ append(data, 178)
				i = 9
			case 4:
				i = 1
				data = /*line tQSyFgc.go:1*/ append(data, 182)
			}
		}
		return /*line tQSyFgc.go:1*/ string(data)
	}(): CALLVALUE,
	func() /*line ScXrAyt5Dg5.go:1*/ string {
		data := /*line ScXrAyt5Dg5.go:1*/ make([]byte, 0, 13)
		i := 1
		decryptKey := 58
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 2
				data = /*line ScXrAyt5Dg5.go:1*/ append(data, "ae"...,
				)
			case 3:
				data = /*line ScXrAyt5Dg5.go:1*/ append(data, "ejn"...,
				)
				i = 0
			case 4:
				i = 3
				data = /*line ScXrAyt5Dg5.go:1*/ append(data, 100)
			case 2:
				for y := range data {
					data[y] = data[y] ^ /*line ScXrAyt5Dg5.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 0:
				data = /*line ScXrAyt5Dg5.go:1*/ append(data, "xlnl"...,
				)
				i = 6
			case 1:
				data = /*line ScXrAyt5Dg5.go:1*/ append(data, "ij"...,
				)
				i = 4
			}
		}
		return /*line ScXrAyt5Dg5.go:1*/ string(data)
	}(): CALLDATALOAD,
	func() /*line sYqcNAYY.go:1*/ string {
		data := [] /*line sYqcNAYY.go:1*/ byte("J\x8bLL\x8eQ~\xb2\x93IB]")
		positions := [...]byte{0, 8, 1, 4, 10, 11, 8, 5, 6, 8, 8, 7, 11, 10}
		for i := 0; i < 14; i += 2 {
			localKey := /*line sYqcNAYY.go:1*/ byte(i) + /*line sYqcNAYY.go:1*/ byte(positions[i]^positions[i+1]) + 200
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line sYqcNAYY.go:1*/ string(data)
	}(): CALLDATASIZE,
	func() /*line ArB1SITdL.go:1*/ string {
		seed := /*line ArB1SITdL.go:1*/ byte(194)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line ArB1SITdL.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line ArB1SITdL.go:1*/ fnc(5)(8)(27)(54)(100)(197)(157)(39)(80)(172)(89)(187)
		return /*line ArB1SITdL.go:1*/ string(data)
	}(): CALLDATACOPY,
	"CHAINID": CHAINID,
	"BASEFEE": BASEFEE,
	func() /*line E79u9dDmDqoh.go:1*/ string {
		key := [] /*line E79u9dDmDqoh.go:1*/ byte("b!x\x11\x86\xbf\xff|")
		data := [] /*line E79u9dDmDqoh.go:1*/ byte("\xa4m\xc7S\xce\x00R\xc4")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line E79u9dDmDqoh.go:1*/ string(data)
	}(): BLOBHASH,
	func() /*line QPcneTTfT_.go:1*/ string {
		data := [] /*line QPcneTTfT_.go:1*/ byte("J;CI)N$.F55")
		positions := [...]byte{2, 0, 1, 6, 2, 5, 3, 2, 4, 5, 7, 4, 2, 0, 9, 10}
		for i := 0; i < 16; i += 2 {
			localKey := /*line QPcneTTfT_.go:1*/ byte(i) + /*line QPcneTTfT_.go:1*/ byte(positions[i]^positions[i+1]) + 95
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line QPcneTTfT_.go:1*/ string(data)
	}(): BLOBBASEFEE,
	func() /*line OrdIq5nF.go:1*/ string {
		seed := /*line OrdIq5nF.go:1*/ byte(161)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line OrdIq5nF.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line OrdIq5nF.go:1*/ fnc(229)(195)(5)(11)(30)(54)(249)(227)(202)(18)(41)(194)
		return /*line OrdIq5nF.go:1*/ string(data)
	}(): DELEGATECALL,
	func() /*line gBuohfr5.go:1*/ string {
		data := [] /*line gBuohfr5.go:1*/ byte("%)\x1bTI>\xfe<&6")
		positions := [...]byte{2, 8, 6, 0, 1, 7, 9, 6, 0, 5, 8, 2, 5, 5}
		for i := 0; i < 14; i += 2 {
			localKey := /*line gBuohfr5.go:1*/ byte(i) + /*line gBuohfr5.go:1*/ byte(positions[i]^positions[i+1]) + 222
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line gBuohfr5.go:1*/ string(data)
	}(): STATICCALL,
	func() /*line HPtEXB94e.go:1*/ string {
		key := [] /*line HPtEXB94e.go:1*/ byte("\xba`\xe5{\xe5\xc9d~")
		data := [] /*line HPtEXB94e.go:1*/ byte("\xfd\xaf)\xc08\x12\xbe\xc3")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line HPtEXB94e.go:1*/ string(data)
	}(): CODESIZE,
	func() /*line aaxSfhShruFu.go:1*/ string {
		data := [] /*line aaxSfhShruFu.go:1*/ byte("C\xa5D\x8d\x80@r\x95")
		positions := [...]byte{7, 4, 6, 3, 3, 5, 3, 4, 3, 1}
		for i := 0; i < 10; i += 2 {
			localKey := /*line aaxSfhShruFu.go:1*/ byte(i) + /*line aaxSfhShruFu.go:1*/ byte(positions[i]^positions[i+1]) + 214
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line aaxSfhShruFu.go:1*/ string(data)
	}(): CODECOPY,
	func() /*line O2qUfSr.go:1*/ string {
		data := /*line O2qUfSr.go:1*/ make([]byte, 0, 9)
		i := 4
		decryptKey := 104
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 9:
				data = /*line O2qUfSr.go:1*/ append(data, 222)
				i = 2
			case 5:
				i = 1
				data = /*line O2qUfSr.go:1*/ append(data, 210)
			case 4:
				data = /*line O2qUfSr.go:1*/ append(data, 217)
				i = 5
			case 0:
				for y := range data {
					data[y] = data[y] + /*line O2qUfSr.go:1*/ byte(decryptKey^y)
				}
				i = 8
			case 1:
				i = 3
				data = /*line O2qUfSr.go:1*/ append(data, 231)
			case 6:
				data = /*line O2qUfSr.go:1*/ append(data, 232)
				i = 9
			case 2:
				i = 7
				data = /*line O2qUfSr.go:1*/ append(data, 219)
			case 3:
				i = 6
				data = /*line O2qUfSr.go:1*/ append(data, 227)
			case 7:
				i = 0
				data = /*line O2qUfSr.go:1*/ append(data, 220)
			}
		}
		return /*line O2qUfSr.go:1*/ string(data)
	}(): GASPRICE,
	func() /*line oaNQBWNtr.go:1*/ string {
		data := [] /*line oaNQBWNtr.go:1*/ byte("E\xa3Tr\x83\x14E\xf5\xdbZ\xa3")
		positions := [...]byte{8, 5, 3, 10, 7, 8, 5, 7, 5, 1, 3, 5, 8, 4, 7, 4}
		for i := 0; i < 16; i += 2 {
			localKey := /*line oaNQBWNtr.go:1*/ byte(i) + /*line oaNQBWNtr.go:1*/ byte(positions[i]^positions[i+1]) + 34
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return /*line oaNQBWNtr.go:1*/ string(data)
	}(): EXTCODESIZE,
	func() /*line At_nYeHf0.go:1*/ string {
		key := [] /*line At_nYeHf0.go:1*/ byte("\x83\x13W\xcda\x15rZ\xe5\x99q")
		data := [] /*line At_nYeHf0.go:1*/ byte("\xc6K\x03\x8e.Q7\x19\xaa\xc9(")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line At_nYeHf0.go:1*/ string(data)
	}(): EXTCODECOPY,
	func() /*line DAATO9z5cVM7.go:1*/ string {
		data := /*line DAATO9z5cVM7.go:1*/ make([]byte, 0, 15)
		i := 0
		decryptKey := 155
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				data = /*line DAATO9z5cVM7.go:1*/ append(data, 174)
				i = 4
			case 7:
				i = 1
				data = /*line DAATO9z5cVM7.go:1*/ append(data, 196)
			case 0:
				i = 5
				data = /*line DAATO9z5cVM7.go:1*/ append(data, "\xb8\xaa\xbc\xbc"...,
				)
			case 5:
				i = 6
				data = /*line DAATO9z5cVM7.go:1*/ append(data, "\xb4\xaf"...,
				)
			case 6:
				i = 3
				data = /*line DAATO9z5cVM7.go:1*/ append(data, "\xa8\xa4"...,
				)
			case 4:
				i = 2
				for y := range data {
					data[y] = data[y] + /*line DAATO9z5cVM7.go:1*/ byte(decryptKey^y)
				}
			case 3:
				data = /*line DAATO9z5cVM7.go:1*/ append(data, "®ø"...,
				)
				i = 7
			}
		}
		return /*line DAATO9z5cVM7.go:1*/ string(data)
	}(): RETURNDATASIZE,
	func() /*line BoVC9qs1.go:1*/ string {
		key := [] /*line BoVC9qs1.go:1*/ byte("qx\x06%B\x9cP.\x8c\xbbG\xc7\xf1\xc2")
		data := [] /*line BoVC9qs1.go:1*/ byte("ýZz\x94\xea\x94o\xe0\xfc\x8a\x16A\x1b")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line BoVC9qs1.go:1*/ string(data)
	}(): RETURNDATACOPY,
	func() /*line F7PMuaqtr0.go:1*/ string {
		seed := /*line F7PMuaqtr0.go:1*/ byte(139)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line F7PMuaqtr0.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line F7PMuaqtr0.go:1*/ fnc(186)(19)(252)(239)(12)(245)(1)(3)(249)(18)(245)
		return /*line F7PMuaqtr0.go:1*/ string(data)
	}(): EXTCODEHASH,
	func() /*line lJLwhM2_q.go:1*/ string {
		key := [] /*line lJLwhM2_q.go:1*/ byte("\xb47\xea\x82d\x9a;\xeb\x01")
		data := [] /*line lJLwhM2_q.go:1*/ byte("\xf6{\xa5\xc1/\xd2z\xb8I")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line lJLwhM2_q.go:1*/ string(data)
	}(): BLOCKHASH,
	func() /*line lXq5rd.go:1*/ string {
		data := /*line lXq5rd.go:1*/ make([]byte, 0, 9)
		i := 6
		decryptKey := 96
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				data = /*line lXq5rd.go:1*/ append(data, 229)
				i = 0
			case 4:
				i = 7
				data = /*line lXq5rd.go:1*/ append(data, 224)
			case 8:
				i = 2
				data = /*line lXq5rd.go:1*/ append(data, 243)
			case 7:
				data = /*line lXq5rd.go:1*/ append(data, 222)
				i = 8
			case 6:
				data = /*line lXq5rd.go:1*/ append(data, 221)
				i = 9
			case 2:
				i = 1
				data = /*line lXq5rd.go:1*/ append(data, 228)
			case 0:
				data = /*line lXq5rd.go:1*/ append(data, 233)
				i = 4
			case 9:
				i = 5
				data = /*line lXq5rd.go:1*/ append(data, 232)
			case 1:
				i = 3
				for y := range data {
					data[y] = data[y] + /*line lXq5rd.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return /*line lXq5rd.go:1*/ string(data)
	}(): COINBASE,
	func() /*line MLqOPqwf.go:1*/ string {
		data := [] /*line MLqOPqwf.go:1*/ byte("rxukSTytP")
		positions := [...]byte{2, 3, 0, 6, 6, 3, 3, 6, 3, 7, 1, 7}
		for i := 0; i < 12; i += 2 {
			localKey := /*line MLqOPqwf.go:1*/ byte(i) + /*line MLqOPqwf.go:1*/ byte(positions[i]^positions[i+1]) + 37
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line MLqOPqwf.go:1*/ string(data)
	}(): TIMESTAMP,
	"NUMBER": NUMBER,
	func() /*line MS61wzyyWLDx.go:1*/ string {
		fullData := [] /*line MS61wzyyWLDx.go:1*/ byte("\xe8=Ry\xcbNS[\x0f\xf9\x1f\x9f\xba\xbf\xac\x1bB\xf8i\xa9")
		idxKey := [] /*line MS61wzyyWLDx.go:1*/ byte("X\xbd")
		data := /*line MS61wzyyWLDx.go:1*/ make([]byte, 0, 11)
		data = /*line MS61wzyyWLDx.go:1*/ append(data, fullData[181^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])]-fullData[185^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])], fullData[72^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])]-fullData[81^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])], fullData[172^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])]+fullData[184^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])], fullData[85^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])]-fullData[91^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])], fullData[87^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])]^fullData[90^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])], fullData[189^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])]+fullData[186^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])], fullData[174^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])]+fullData[179^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])], fullData[94^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])]^fullData[82^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[0])], fullData[188^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])]^fullData[175^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])], fullData[182^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])]+fullData[177^ /*line MS61wzyyWLDx.go:1*/ int(idxKey[1])])
		return /*line MS61wzyyWLDx.go:1*/ string(data)
	}(): DIFFICULTY,
	func() /*line r1UDWmv6.go:1*/ string {
		key := [] /*line r1UDWmv6.go:1*/ byte("\x9ej\u008f\"\x86G7")
		data := [] /*line r1UDWmv6.go:1*/ byte("\xe5\xab\x15\xdbkӐ\x8b")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line r1UDWmv6.go:1*/ string(data)
	}(): GASLIMIT,
	func() /*line Ni_b7cFVhBYi.go:1*/ string {
		data := [] /*line Ni_b7cFVhBYi.go:1*/ byte("=:L\xa5\x92ALAទ")
		positions := [...]byte{1, 0, 1, 8, 10, 1, 0, 3, 3, 4, 9, 8}
		for i := 0; i < 12; i += 2 {
			localKey := /*line Ni_b7cFVhBYi.go:1*/ byte(i) + /*line Ni_b7cFVhBYi.go:1*/ byte(positions[i]^positions[i+1]) + 165
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line Ni_b7cFVhBYi.go:1*/ string(data)
	}(): SELFBALANCE,
	"POP":     POP,
	"MLOAD":   MLOAD,
	"MSTORE":  MSTORE,
	"MSTORE8": MSTORE8,
	"SLOAD":   SLOAD,
	"SSTORE":  SSTORE,
	"JUMP":    JUMP,
	"JUMPI":   JUMPI,
	"PC":      PC,
	"MSIZE":   MSIZE,
	"GAS":     GAS,
	func() /*line Ahyy3j.go:1*/ string {
		data := [] /*line Ahyy3j.go:1*/ byte("\xbe\x8d\xc7P\xcā\x88")
		positions := [...]byte{1, 0, 6, 5, 2, 7, 4, 4, 1, 4, 6, 7}
		for i := 0; i < 12; i += 2 {
			localKey := /*line Ahyy3j.go:1*/ byte(i) + /*line Ahyy3j.go:1*/ byte(positions[i]^positions[i+1]) + 188
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line Ahyy3j.go:1*/ string(data)
	}(): JUMPDEST,
	"TLOAD":   TLOAD,
	"TSTORE":  TSTORE,
	"MCOPY":   MCOPY,
	"PUSH0":   PUSH0,
	"PUSH1":   PUSH1,
	"PUSH2":   PUSH2,
	"PUSH3":   PUSH3,
	"PUSH4":   PUSH4,
	"PUSH5":   PUSH5,
	"PUSH6":   PUSH6,
	"PUSH7":   PUSH7,
	"PUSH8":   PUSH8,
	"PUSH9":   PUSH9,
	"PUSH10":  PUSH10,
	"PUSH11":  PUSH11,
	"PUSH12":  PUSH12,
	"PUSH13":  PUSH13,
	"PUSH14":  PUSH14,
	"PUSH15":  PUSH15,
	"PUSH16":  PUSH16,
	"PUSH17":  PUSH17,
	"PUSH18":  PUSH18,
	"PUSH19":  PUSH19,
	"PUSH20":  PUSH20,
	"PUSH21":  PUSH21,
	"PUSH22":  PUSH22,
	"PUSH23":  PUSH23,
	"PUSH24":  PUSH24,
	"PUSH25":  PUSH25,
	"PUSH26":  PUSH26,
	"PUSH27":  PUSH27,
	"PUSH28":  PUSH28,
	"PUSH29":  PUSH29,
	"PUSH30":  PUSH30,
	"PUSH31":  PUSH31,
	"PUSH32":  PUSH32,
	"DUP1":    DUP1,
	"DUP2":    DUP2,
	"DUP3":    DUP3,
	"DUP4":    DUP4,
	"DUP5":    DUP5,
	"DUP6":    DUP6,
	"DUP7":    DUP7,
	"DUP8":    DUP8,
	"DUP9":    DUP9,
	"DUP10":   DUP10,
	"DUP11":   DUP11,
	"DUP12":   DUP12,
	"DUP13":   DUP13,
	"DUP14":   DUP14,
	"DUP15":   DUP15,
	"DUP16":   DUP16,
	"SWAP1":   SWAP1,
	"SWAP2":   SWAP2,
	"SWAP3":   SWAP3,
	"SWAP4":   SWAP4,
	"SWAP5":   SWAP5,
	"SWAP6":   SWAP6,
	"SWAP7":   SWAP7,
	"SWAP8":   SWAP8,
	"SWAP9":   SWAP9,
	"SWAP10":  SWAP10,
	"SWAP11":  SWAP11,
	"SWAP12":  SWAP12,
	"SWAP13":  SWAP13,
	"SWAP14":  SWAP14,
	"SWAP15":  SWAP15,
	"SWAP16":  SWAP16,
	"LOG0":    LOG0,
	"LOG1":    LOG1,
	"LOG2":    LOG2,
	"LOG3":    LOG3,
	"LOG4":    LOG4,
	"CREATE":  CREATE,
	"CREATE2": CREATE2,
	"CALL":    CALL,
	"RETURN":  RETURN,
	func() /*line GMUlJcU5nT.go:1*/ string {
		fullData := [] /*line GMUlJcU5nT.go:1*/ byte(";G\xeeG\b\xa2:\xbdsu\xd0\xf1\x03\xf8\xe3\xd0")
		idxKey := [] /*line GMUlJcU5nT.go:1*/ byte("\xd4Q{\xeeD\x00")
		data := /*line GMUlJcU5nT.go:1*/ make([]byte, 0, 9)
		data = /*line GMUlJcU5nT.go:1*/ append(data, fullData[238^ /*line GMUlJcU5nT.go:1*/ int(idxKey[3])]-fullData[227^ /*line GMUlJcU5nT.go:1*/ int(idxKey[3])], fullData[74^ /*line GMUlJcU5nT.go:1*/ int(idxKey[4])]-fullData[65^ /*line GMUlJcU5nT.go:1*/ int(idxKey[4])], fullData[87^ /*line GMUlJcU5nT.go:1*/ int(idxKey[1])]-fullData[83^ /*line GMUlJcU5nT.go:1*/ int(idxKey[1])], fullData[79^ /*line GMUlJcU5nT.go:1*/ int(idxKey[4])]^fullData[67^ /*line GMUlJcU5nT.go:1*/ int(idxKey[4])], fullData[10^ /*line GMUlJcU5nT.go:1*/ int(idxKey[5])]+fullData[8^ /*line GMUlJcU5nT.go:1*/ int(idxKey[5])], fullData[82^ /*line GMUlJcU5nT.go:1*/ int(idxKey[1])]+fullData[85^ /*line GMUlJcU5nT.go:1*/ int(idxKey[1])], fullData[239^ /*line GMUlJcU5nT.go:1*/ int(idxKey[3])]-fullData[226^ /*line GMUlJcU5nT.go:1*/ int(idxKey[3])], fullData[9^ /*line GMUlJcU5nT.go:1*/ int(idxKey[5])]+fullData[15^ /*line GMUlJcU5nT.go:1*/ int(idxKey[5])])
		return /*line GMUlJcU5nT.go:1*/ string(data)
	}(): CALLCODE,
	"REVERT":  REVERT,
	"INVALID": INVALID,
	func() /*line JJHOrsL13.go:1*/ string {
		seed := /*line JJHOrsL13.go:1*/ byte(9)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line JJHOrsL13.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line JJHOrsL13.go:1*/ fnc(92)(170)(91)(176)(94)(189)(136)(17)(32)(67)(116)(249)
		return /*line JJHOrsL13.go:1*/ string(data)
	}(): SELFDESTRUCT,
}

func PCaY86YPpL(xaFC3hH string) LxosJe8 {
	return coCamfmb[xaFC3hH]
}

var _ = fmt.Append
