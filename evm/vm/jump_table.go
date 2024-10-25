//line :1
package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/params"
)

type (
	executionFunc func(eAorajeVy084 *uint64, i6HsYcs_9RSM *EVMInterpreter, ahYrRb0Iw7v0 *GUBKEu339CA) ([]byte, error)
	gasFunc       func(*EVM, *WyJ004I8, *IZHSJVjaAmxb, *VfNB1WFtybE, uint64) (uint64, error)

	memorySizeFunc func(*IZHSJVjaAmxb) (jLa0ywd uint64, r9uWm2pY bool)
)

type operation struct {
	execute     executionFunc
	constantGas uint64
	dynamicGas  gasFunc

	minStack int

	maxStack int

	memorySize memorySizeFunc
}

var (
	zi10f_aku1a  = /*line BRtOHm.go:1*/ bATEwUY()
	slTZ7K       = /*line OM0ekunPZuXW.go:1*/ j1_EIof()
	iqRN8fb      = /*line jiOPSNagGPW.go:1*/ mVuYzPpq()
	eTszvRda     = /*line lPrzwg.go:1*/ r55_aGYuJs()
	eVTfqUZ5XLKp = /*line SEwpCzhFFI.go:1*/ akQ0LksfZ()
	ygkeaMUJ     = /*line ETth_Oa.go:1*/ aaQtRKQWQEX()
	cslDH5rwHrW2 = /*line SSo5E4SN.go:1*/ qsNj3ExQbuEI()
	l7p1GOaGYiab = /*line JoWTaCm2Et1.go:1*/ f8rBfy_rdC()
	myuFbYmfOo7  = /*line l0xL78hm5.go:1*/ fuEI3sHVt9hH()
	f1EQuA       = /*line pR7QM6B.go:1*/ in9K_X62jm()
	ecLGn5YUA_   = /*line BnNVyvNRE49F.go:1*/ ddIsmUpbc()
	hEiTFfkGV_   = /*line y8kY9qb.go:1*/ s8uFHpJ()
)

type JumpTable [256]*operation

func tmwbqHY(kOAWieB3PO8 JumpTable) JumpTable {
	for hv0ZaFQN, nCnJ0mc := range kOAWieB3PO8 {
		if nCnJ0mc == nil {
			/*line _BD6444T44R.go:1*/ panic( /*line qLrs43e.go:1*/ fmt.Sprintf(func() /*line ZQM9wsSj.go:1*/ string {
				fullData := [] /*line ZQM9wsSj.go:1*/ byte("%\xa1\x8b\xb2\x10\xed\xd7:\xbf\xe40\xb2x\xbdĄ\x18Q8\xc6;[\xb6\x98\xab~I\x11\x9b\x9f\xca\xc9C\v")
				idxKey := [] /*line ZQM9wsSj.go:1*/ byte("\xea&")
				data := /*line ZQM9wsSj.go:1*/ make([]byte, 0, 18)
				data = /*line ZQM9wsSj.go:1*/ append(data, fullData[253^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]+fullData[236^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[7^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]-fullData[58^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[247^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]^fullData[226^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[39^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]+fullData[41^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[54^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]^fullData[50^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[45^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]^fullData[56^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[44^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]^fullData[34^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[248^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]^fullData[251^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[234^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]-fullData[233^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[242^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]-fullData[232^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[245^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]-fullData[255^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[61^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]^fullData[63^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[231^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]-fullData[240^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[227^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]^fullData[228^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])], fullData[48^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]-fullData[6^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[42^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])]+fullData[35^ /*line ZQM9wsSj.go:1*/ int(idxKey[1])], fullData[237^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])]-fullData[249^ /*line ZQM9wsSj.go:1*/ int(idxKey[0])])
				return /*line ZQM9wsSj.go:1*/ string(data)
			}(), hv0ZaFQN))
		}

		if nCnJ0mc.memorySize != nil && nCnJ0mc.dynamicGas == nil {
			/*line Toz0QXNm.go:1*/ panic( /*line YPYJVxk.go:1*/ fmt.Sprintf(func() /*line wwYPmU6UFD.go:1*/ string {
				seed := /*line wwYPmU6UFD.go:1*/ byte(116)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line wwYPmU6UFD.go:1*/ append(data, x-seed); seed += x; return fnc }
				/*line wwYPmU6UFD.go:1*/ fnc(227)(199)(62)(129)(83)(80)(232)(201)(164)(245)(46)(113)(215)(161)(78)(152)(42)(17)(111)(214)(180)(106)(215)(181)(17)(100)(219)(181)(22)(122)(245)(239)(138)(88)(197)(127)(241)(238)(216)(170)(17)(105)(204)(170)
				return /*line wwYPmU6UFD.go:1*/ string(data)
			}(), /*line eCYGi4cFL6vr.go:1*/ LxosJe8(hv0ZaFQN).String()))
		}
	}
	return kOAWieB3PO8
}

func s8uFHpJ() JumpTable {
	ve85g6Vu := /*line o2KKDpZ.go:1*/ ddIsmUpbc()
	/*line lmPIA7w.go:1*/ z8cuFbRWqBJd(&ve85g6Vu)
	/*line QgMn4JjQp.go:1*/ uE3Hi9(&ve85g6Vu)
	/*line Hq9G_7OXBb.go:1*/ agVtvxDSUs(&ve85g6Vu)
	/*line UvkdGF.go:1*/ ghNjp4tedh(&ve85g6Vu)
	/*line GfFckvoPa.go:1*/ eC3GlL4(&ve85g6Vu)
	return /*line diDd7JWeTgr.go:1*/ tmwbqHY(ve85g6Vu)
}

func ddIsmUpbc() JumpTable {
	ve85g6Vu := /*line rb8ZnoS.go:1*/ in9K_X62jm()
	/*line YIo4sDI.go:1*/ wKpQXZpUwK(&ve85g6Vu)
	/*line CZNd9Qb.go:1*/ wuw9VQEHA6(&ve85g6Vu)

	return /*line k9XCXf.go:1*/ tmwbqHY(ve85g6Vu)
}

func in9K_X62jm() JumpTable {
	ve85g6Vu := /*line oEtWa8N.go:1*/ fuEI3sHVt9hH()
	ve85g6Vu[PREVRANDAO] = &operation{
		execute:     vZaNI8ewO9y,
		constantGas: GasQuickStep,
		minStack:/*line EQsJaa8.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line lOaFsdkQJ.go:1*/ mETN8x(0, 1),
	}
	return /*line K4SvvEMNqsM.go:1*/ tmwbqHY(ve85g6Vu)
}

func fuEI3sHVt9hH() JumpTable {
	ve85g6Vu := /*line AcH0Os_Ft5KM.go:1*/ f8rBfy_rdC()
	/*line AEroaZQ.go:1*/ dHGM24Ut0F(&ve85g6Vu)
	/*line Hv857i.go:1*/ hgndeOut9ePa(&ve85g6Vu)
	return /*line DAGtfw07LiZB.go:1*/ tmwbqHY(ve85g6Vu)
}

func f8rBfy_rdC() JumpTable {
	ve85g6Vu := /*line lClScun.go:1*/ qsNj3ExQbuEI()
	/*line Idz6mDi1_.go:1*/ wIyo_cTzT4(&ve85g6Vu)
	return /*line YH2meagmf.go:1*/ tmwbqHY(ve85g6Vu)
}

func qsNj3ExQbuEI() JumpTable {
	ve85g6Vu := /*line b8LkHXd1HxqN.go:1*/ aaQtRKQWQEX()

	/*line x5WvNnJQSEA.go:1*/
	ntMBqQpN5sa(&ve85g6Vu)
	/*line rkY25qdXD.go:1*/ bIhOZSNloD(&ve85g6Vu)
	/*line DE63nlDkB.go:1*/ e4gyUrMOn3L(&ve85g6Vu)

	return /*line MgJSojkZrt_4.go:1*/ tmwbqHY(ve85g6Vu)
}

func aaQtRKQWQEX() JumpTable {
	ve85g6Vu := /*line gqSWlxVfaUV.go:1*/ akQ0LksfZ()
	ve85g6Vu[SHL] = &operation{
		execute:     pI3Ubpr,
		constantGas: GasFastestStep,
		minStack:/*line QR5bTDGz2.go:1*/ d0cam_EE09tL(2, 1),
		maxStack:/*line lWlajqTltdlg.go:1*/ mETN8x(2, 1),
	}
	ve85g6Vu[SHR] = &operation{
		execute:     hCmHXJSR0,
		constantGas: GasFastestStep,
		minStack:/*line FCtvVt.go:1*/ d0cam_EE09tL(2, 1),
		maxStack:/*line bXhd9fhW.go:1*/ mETN8x(2, 1),
	}
	ve85g6Vu[SAR] = &operation{
		execute:     xkQje0o0,
		constantGas: GasFastestStep,
		minStack:/*line rcLp51W.go:1*/ d0cam_EE09tL(2, 1),
		maxStack:/*line JAqoYazX.go:1*/ mETN8x(2, 1),
	}
	ve85g6Vu[EXTCODEHASH] = &operation{
		execute:     hq9dOMVJu,
		constantGas: params.ExtcodeHashGasConstantinople,
		minStack:/*line ofkYSX.go:1*/ d0cam_EE09tL(1, 1),
		maxStack:/*line xDKybNnLlP0.go:1*/ mETN8x(1, 1),
	}
	ve85g6Vu[CREATE2] = &operation{
		execute:     ytM3XsXC,
		constantGas: params.Create2Gas,
		dynamicGas:  zBgFdndQ,
		minStack:/*line p2JH6j.go:1*/ d0cam_EE09tL(4, 1),
		maxStack:/*line YKxUBGUkCS.go:1*/ mETN8x(4, 1),
		memorySize: g9l43wp,
	}
	return /*line QYL_PcLdR.go:1*/ tmwbqHY(ve85g6Vu)
}

func akQ0LksfZ() JumpTable {
	ve85g6Vu := /*line ZHTeTE5pmP.go:1*/ r55_aGYuJs()
	ve85g6Vu[STATICCALL] = &operation{
		execute:     r2rEDFn,
		constantGas: params.CallGasEIP150,
		dynamicGas:  svRiUsX,
		minStack:/*line AyFBPrM.go:1*/ d0cam_EE09tL(6, 1),
		maxStack:/*line vSIIRGv.go:1*/ mETN8x(6, 1),
		memorySize: ayRQao7EU,
	}
	ve85g6Vu[RETURNDATASIZE] = &operation{
		execute:     bwz5npcZy_x,
		constantGas: GasQuickStep,
		minStack:/*line MkaBsHq37.go:1*/ d0cam_EE09tL(0, 1),
		maxStack:/*line N8FqNtMkzuJh.go:1*/ mETN8x(0, 1),
	}
	ve85g6Vu[RETURNDATACOPY] = &operation{
		execute:     pGrLEaIDs,
		constantGas: GasFastestStep,
		dynamicGas:  g8KCX7,
		minStack:/*line y_qrXvUVnO9.go:1*/ d0cam_EE09tL(3, 0),
		maxStack:/*line Rhf1qe3M_vr.go:1*/ mETN8x(3, 0),
		memorySize: nMYGGlURs,
	}
	ve85g6Vu[REVERT] = &operation{
		execute:    aKDpaP,
		dynamicGas: gt2cFYPNW,
		minStack:/*line m8g8jw4UOd6.go:1*/ d0cam_EE09tL(2, 0),
		maxStack:/*line MeMd17uD.go:1*/ mETN8x(2, 0),
		memorySize: nfLk0ccx3SSZ,
	}
	return /*line JMDDVrk7.go:1*/ tmwbqHY(ve85g6Vu)
}

func r55_aGYuJs() JumpTable {
	ve85g6Vu := /*line aWtAEnTeT.go:1*/ mVuYzPpq()
	ve85g6Vu[EXP].dynamicGas = cpcGfh
	return /*line DxoG4duH.go:1*/ tmwbqHY(ve85g6Vu)
}

func mVuYzPpq() JumpTable {
	ve85g6Vu := /*line dOVDTLRP.go:1*/ j1_EIof()
	ve85g6Vu[BALANCE].constantGas = params.BalanceGasEIP150
	ve85g6Vu[EXTCODESIZE].constantGas = params.ExtcodeSizeGasEIP150
	ve85g6Vu[SLOAD].constantGas = params.SloadGasEIP150
	ve85g6Vu[EXTCODECOPY].constantGas = params.ExtcodeCopyBaseEIP150
	ve85g6Vu[CALL].constantGas = params.CallGasEIP150
	ve85g6Vu[CALLCODE].constantGas = params.CallGasEIP150
	ve85g6Vu[DELEGATECALL].constantGas = params.CallGasEIP150
	return /*line PHpCzIFrmg.go:1*/ tmwbqHY(ve85g6Vu)
}

func j1_EIof() JumpTable {
	ve85g6Vu := /*line QH97WSre7sCP.go:1*/ bATEwUY()
	ve85g6Vu[DELEGATECALL] = &operation{
		execute:     hDef9GRB,
		dynamicGas:  kQ1IHor,
		constantGas: params.CallGasFrontier,
		minStack:/*line CcGTqr_d4J.go:1*/ d0cam_EE09tL(6, 1),
		maxStack:/*line mR9FxGm.go:1*/ mETN8x(6, 1),
		memorySize: aeayOiZ,
	}
	return /*line kxUzdj5o1.go:1*/ tmwbqHY(ve85g6Vu)
}

func bATEwUY() JumpTable {
	gEGKOX := JumpTable{
		STOP: {
			execute:     bstKGE2y,
			constantGas: 0,
			minStack:/*line Qfq3fukBm.go:1*/ d0cam_EE09tL(0, 0),
			maxStack:/*line ICjD88.go:1*/ mETN8x(0, 0),
		},
		ADD: {
			execute:     oX7kObHW,
			constantGas: GasFastestStep,
			minStack:/*line kWuTU6IL.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line y3nZCwDj3GSA.go:1*/ mETN8x(2, 1),
		},
		MUL: {
			execute:     jyWLifLT6,
			constantGas: GasFastStep,
			minStack:/*line CCWVQKtzkHI.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line XHKmMgaL6hxW.go:1*/ mETN8x(2, 1),
		},
		SUB: {
			execute:     qSFoLqsn,
			constantGas: GasFastestStep,
			minStack:/*line EH7ywqqCfa.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line A4qNg0Q816Zd.go:1*/ mETN8x(2, 1),
		},
		DIV: {
			execute:     dpEU4dBZ_2,
			constantGas: GasFastStep,
			minStack:/*line JvFsgRih.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line gD1wwi.go:1*/ mETN8x(2, 1),
		},
		SDIV: {
			execute:     aUETnuDpdMcV,
			constantGas: GasFastStep,
			minStack:/*line W0NcGc4x.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line jaW50R.go:1*/ mETN8x(2, 1),
		},
		MOD: {
			execute:     jUisG0sV,
			constantGas: GasFastStep,
			minStack:/*line hg84unE.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line MRnknEVasS.go:1*/ mETN8x(2, 1),
		},
		SMOD: {
			execute:     iZ3IEay,
			constantGas: GasFastStep,
			minStack:/*line bPzlF61.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line wZMCpnBq.go:1*/ mETN8x(2, 1),
		},
		ADDMOD: {
			execute:     zD9oyOC009E,
			constantGas: GasMidStep,
			minStack:/*line q0cUjvE0.go:1*/ d0cam_EE09tL(3, 1),
			maxStack:/*line V5Fqg2Wx3.go:1*/ mETN8x(3, 1),
		},
		MULMOD: {
			execute:     diTvlz3,
			constantGas: GasMidStep,
			minStack:/*line Kqc6_EcE.go:1*/ d0cam_EE09tL(3, 1),
			maxStack:/*line Ox4mBWocH.go:1*/ mETN8x(3, 1),
		},
		EXP: {
			execute:    hKqv_CFbs,
			dynamicGas: pJOGLJAn,
			minStack:/*line HHh7Wm0BZ.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line SStuHsCW8.go:1*/ mETN8x(2, 1),
		},
		SIGNEXTEND: {
			execute:     zuO1DQ7,
			constantGas: GasFastStep,
			minStack:/*line B63RhaFJG.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line xC2oJWFf.go:1*/ mETN8x(2, 1),
		},
		LT: {
			execute:     tdJw_kOUY,
			constantGas: GasFastestStep,
			minStack:/*line bTDdpO.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line FHgPaEx.go:1*/ mETN8x(2, 1),
		},
		GT: {
			execute:     mwHPKgTNAb2,
			constantGas: GasFastestStep,
			minStack:/*line FQBTj3nE.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line taieo3rKWK2.go:1*/ mETN8x(2, 1),
		},
		SLT: {
			execute:     aN4rhyAjO,
			constantGas: GasFastestStep,
			minStack:/*line Bnif__q4maG.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line Uofsf8g.go:1*/ mETN8x(2, 1),
		},
		SGT: {
			execute:     gaqtaLjFVoO,
			constantGas: GasFastestStep,
			minStack:/*line Dq8Dhx.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line cXXtyhioRJ.go:1*/ mETN8x(2, 1),
		},
		EQ: {
			execute:     aoaLUkw1,
			constantGas: GasFastestStep,
			minStack:/*line MBMAfI.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line Myp3RuWjM2.go:1*/ mETN8x(2, 1),
		},
		ISZERO: {
			execute:     ovsI44pN,
			constantGas: GasFastestStep,
			minStack:/*line eYspIPstqHa.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line yEyhZq32.go:1*/ mETN8x(1, 1),
		},
		AND: {
			execute:     xrOGoCDobg,
			constantGas: GasFastestStep,
			minStack:/*line ZX3nWL9.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line NE8ffo3wjz.go:1*/ mETN8x(2, 1),
		},
		XOR: {
			execute:     cJPL60zZ,
			constantGas: GasFastestStep,
			minStack:/*line jAmsP5B7o.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line rbAtRcPSadiN.go:1*/ mETN8x(2, 1),
		},
		OR: {
			execute:     kXM7xm3,
			constantGas: GasFastestStep,
			minStack:/*line GoSrbZKj.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line BeDqFNAronf.go:1*/ mETN8x(2, 1),
		},
		NOT: {
			execute:     laabqEP,
			constantGas: GasFastestStep,
			minStack:/*line BagBgBr.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line I32vcNu9hqT.go:1*/ mETN8x(1, 1),
		},
		BYTE: {
			execute:     iLNxwWV1g,
			constantGas: GasFastestStep,
			minStack:/*line rBdWpmHJ.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line xiAwS6J.go:1*/ mETN8x(2, 1),
		},
		KECCAK256: {
			execute:     b2t6gaCJyUCQ,
			constantGas: params.Keccak256Gas,
			dynamicGas:  jNJPVlmt,
			minStack:/*line JA4pewk.go:1*/ d0cam_EE09tL(2, 1),
			maxStack:/*line cpx5SOhqy1Y.go:1*/ mETN8x(2, 1),
			memorySize: xsTtg2_dZhV,
		},
		ADDRESS: {
			execute:     aYNCt_kHR1aG,
			constantGas: GasQuickStep,
			minStack:/*line l2E8PdFp_LbV.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line eqeFY8SSX.go:1*/ mETN8x(0, 1),
		},
		BALANCE: {
			execute:     wfFEz6PeFzF2,
			constantGas: params.BalanceGasFrontier,
			minStack:/*line G5CfwpHVe.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line Mpzhk95.go:1*/ mETN8x(1, 1),
		},
		ORIGIN: {
			execute:     q4uqMXP,
			constantGas: GasQuickStep,
			minStack:/*line fPd7Ier03C.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line DEoTCx1.go:1*/ mETN8x(0, 1),
		},
		CALLER: {
			execute:     kZIPn77Wz,
			constantGas: GasQuickStep,
			minStack:/*line toOvl9ehQLbt.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line rlmi57w_6Yx.go:1*/ mETN8x(0, 1),
		},
		CALLVALUE: {
			execute:     uOad9IVxe,
			constantGas: GasQuickStep,
			minStack:/*line RL5HSmbZck.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line aHQbu8a.go:1*/ mETN8x(0, 1),
		},
		CALLDATALOAD: {
			execute:     fvPF_mdlH,
			constantGas: GasFastestStep,
			minStack:/*line mW9nw3S.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line BcQQX7NulaKY.go:1*/ mETN8x(1, 1),
		},
		CALLDATASIZE: {
			execute:     dLFWUbq2o,
			constantGas: GasQuickStep,
			minStack:/*line vJz4pUX.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line HY01zLD.go:1*/ mETN8x(0, 1),
		},
		CALLDATACOPY: {
			execute:     vJImE5oGbH,
			constantGas: GasFastestStep,
			dynamicGas:  uFCmbjA,
			minStack:/*line X664yI_vDups.go:1*/ d0cam_EE09tL(3, 0),
			maxStack:/*line Vnmo2a_6R0.go:1*/ mETN8x(3, 0),
			memorySize: gkWdHhXy3,
		},
		CODESIZE: {
			execute:     ywlX7e,
			constantGas: GasQuickStep,
			minStack:/*line D_FmAYr.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line RwbMtju.go:1*/ mETN8x(0, 1),
		},
		CODECOPY: {
			execute:     jj_sMGNoEwv,
			constantGas: GasFastestStep,
			dynamicGas:  eaHmd75RS,
			minStack:/*line qj5Iba.go:1*/ d0cam_EE09tL(3, 0),
			maxStack:/*line HoqQzGEK2CFR.go:1*/ mETN8x(3, 0),
			memorySize: tatHA4q4SN,
		},
		GASPRICE: {
			execute:     eYiQu4Nk,
			constantGas: GasQuickStep,
			minStack:/*line bINaamazAANX.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line a7_Ipc3.go:1*/ mETN8x(0, 1),
		},
		EXTCODESIZE: {
			execute:     bkxY032LmmI,
			constantGas: params.ExtcodeSizeGasFrontier,
			minStack:/*line AKLe2uq.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line GUaO6He_.go:1*/ mETN8x(1, 1),
		},
		EXTCODECOPY: {
			execute:     gwwrLw_N8,
			constantGas: params.ExtcodeCopyBaseFrontier,
			dynamicGas:  jgfKmr,
			minStack:/*line tbpQoUcCg.go:1*/ d0cam_EE09tL(4, 0),
			maxStack:/*line CVJp1YAcbbg.go:1*/ mETN8x(4, 0),
			memorySize: cC9LQO2b1,
		},
		BLOCKHASH: {
			execute:     e1Psf94lt,
			constantGas: GasExtStep,
			minStack:/*line W9vpuXt3nd.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line dhZggkrf4dxi.go:1*/ mETN8x(1, 1),
		},
		COINBASE: {
			execute:     iTOQ8c6yAs,
			constantGas: GasQuickStep,
			minStack:/*line lVIVuLNP.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line IjCslKMweC.go:1*/ mETN8x(0, 1),
		},
		TIMESTAMP: {
			execute:     pNlxWyL,
			constantGas: GasQuickStep,
			minStack:/*line TL2D2ljfuq9.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line jel8iM.go:1*/ mETN8x(0, 1),
		},
		NUMBER: {
			execute:     na_6wKu,
			constantGas: GasQuickStep,
			minStack:/*line JbTlP1Gz.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line xGHOMAhr.go:1*/ mETN8x(0, 1),
		},
		DIFFICULTY: {
			execute:     v3cd0zrgYBMK,
			constantGas: GasQuickStep,
			minStack:/*line IQffXoEw2.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line K8sn2FmFZ.go:1*/ mETN8x(0, 1),
		},
		GASLIMIT: {
			execute:     j4nGWV_j,
			constantGas: GasQuickStep,
			minStack:/*line rTMsSHq.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line HCEnH3n83Bjp.go:1*/ mETN8x(0, 1),
		},
		POP: {
			execute:     adf4Ga,
			constantGas: GasQuickStep,
			minStack:/*line CIGKH4YQs8u.go:1*/ d0cam_EE09tL(1, 0),
			maxStack:/*line WY1JpqTWu.go:1*/ mETN8x(1, 0),
		},
		MLOAD: {
			execute:     gWp7v4Mkx,
			constantGas: GasFastestStep,
			dynamicGas:  mqEO_xnUyuwM,
			minStack:/*line rTbaqILf.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line MQEOWdlyNUB.go:1*/ mETN8x(1, 1),
			memorySize: ayJ1VM5kO,
		},
		MSTORE: {
			execute:     rYq4Xq3Kw8,
			constantGas: GasFastestStep,
			dynamicGas:  yfEk_0hk,
			minStack:/*line Gi4khSlLz1.go:1*/ d0cam_EE09tL(2, 0),
			maxStack:/*line ALMCEmsigYfx.go:1*/ mETN8x(2, 0),
			memorySize: xaWFBW,
		},
		MSTORE8: {
			execute:     fNrDMOD,
			constantGas: GasFastestStep,
			dynamicGas:  jUCdrJVETKmE,
			memorySize:  aRHOlcXhZa,
			minStack:/*line NtP817GZ5.go:1*/ d0cam_EE09tL(2, 0),
			maxStack:/*line y5W0ZmZZOWwt.go:1*/ mETN8x(2, 0),
		},
		SLOAD: {
			execute:     eqyviYiwS7,
			constantGas: params.SloadGasFrontier,
			minStack:/*line R0_CUAyCBdAI.go:1*/ d0cam_EE09tL(1, 1),
			maxStack:/*line pEpSm1_xL.go:1*/ mETN8x(1, 1),
		},
		SSTORE: {
			execute:    x0FlrlvFaav,
			dynamicGas: i7tr1rh,
			minStack:/*line sCRwmHfRk.go:1*/ d0cam_EE09tL(2, 0),
			maxStack:/*line Mki0mQwHar4r.go:1*/ mETN8x(2, 0),
		},
		JUMP: {
			execute:     ccGG7bFdd,
			constantGas: GasMidStep,
			minStack:/*line Cyj_Hh.go:1*/ d0cam_EE09tL(1, 0),
			maxStack:/*line EJmSgZy1TN9.go:1*/ mETN8x(1, 0),
		},
		JUMPI: {
			execute:     aDsWoEpue2,
			constantGas: GasSlowStep,
			minStack:/*line ptkadn.go:1*/ d0cam_EE09tL(2, 0),
			maxStack:/*line aFdqnW1u95L.go:1*/ mETN8x(2, 0),
		},
		PC: {
			execute:     s_89rC8B,
			constantGas: GasQuickStep,
			minStack:/*line MKqJuWqq9.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line BH02REu9.go:1*/ mETN8x(0, 1),
		},
		MSIZE: {
			execute:     f0WYg67Ke,
			constantGas: GasQuickStep,
			minStack:/*line Z9OLdh0F2.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line zY8Wgn.go:1*/ mETN8x(0, 1),
		},
		GAS: {
			execute:     iYgVTmLTIQ,
			constantGas: GasQuickStep,
			minStack:/*line PKp8GDXR.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line pZu0IW.go:1*/ mETN8x(0, 1),
		},
		JUMPDEST: {
			execute:     ixRWzLWLcn,
			constantGas: params.JumpdestGas,
			minStack:/*line qBqdXjl4ijn.go:1*/ d0cam_EE09tL(0, 0),
			maxStack:/*line JEaqdqOnmA.go:1*/ mETN8x(0, 0),
		},
		PUSH1: {
			execute:     e2uCnqwu,
			constantGas: GasFastestStep,
			minStack:/*line FniO5owKTz.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line GidKHpm.go:1*/ mETN8x(0, 1),
		},
		PUSH2: {
			execute:/*line b9MaMCaEYlkR.go:1*/ k5OkhRfA(2, 2),
			constantGas: GasFastestStep,
			minStack:/*line at33ILF5l_.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line rm9Rvt3riC.go:1*/ mETN8x(0, 1),
		},
		PUSH3: {
			execute:/*line lvL0a93gj.go:1*/ k5OkhRfA(3, 3),
			constantGas: GasFastestStep,
			minStack:/*line Qjgk_Y150Fux.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line kTKFWv.go:1*/ mETN8x(0, 1),
		},
		PUSH4: {
			execute:/*line Sti4IRzYSZBg.go:1*/ k5OkhRfA(4, 4),
			constantGas: GasFastestStep,
			minStack:/*line A1g8ylWk.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line CSJHuB0Zz_.go:1*/ mETN8x(0, 1),
		},
		PUSH5: {
			execute:/*line JSm2OM.go:1*/ k5OkhRfA(5, 5),
			constantGas: GasFastestStep,
			minStack:/*line F8c2EII.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line h9bv1CBoL7v.go:1*/ mETN8x(0, 1),
		},
		PUSH6: {
			execute:/*line PxYrG0.go:1*/ k5OkhRfA(6, 6),
			constantGas: GasFastestStep,
			minStack:/*line wKyltlrY_i.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line wirlVQkA.go:1*/ mETN8x(0, 1),
		},
		PUSH7: {
			execute:/*line aafQIliX9a5.go:1*/ k5OkhRfA(7, 7),
			constantGas: GasFastestStep,
			minStack:/*line BGFawEMSv.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line Fa2PH3_W.go:1*/ mETN8x(0, 1),
		},
		PUSH8: {
			execute:/*line YjRiTJDxo.go:1*/ k5OkhRfA(8, 8),
			constantGas: GasFastestStep,
			minStack:/*line hf9JRSF7O.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line SRiBNq.go:1*/ mETN8x(0, 1),
		},
		PUSH9: {
			execute:/*line VaE6fnz.go:1*/ k5OkhRfA(9, 9),
			constantGas: GasFastestStep,
			minStack:/*line Zq6TCpk7W.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line XaUGfkVQi.go:1*/ mETN8x(0, 1),
		},
		PUSH10: {
			execute:/*line DFT6aTWw8M.go:1*/ k5OkhRfA(10, 10),
			constantGas: GasFastestStep,
			minStack:/*line yWw2vsTCca.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line T2hOV_B2J.go:1*/ mETN8x(0, 1),
		},
		PUSH11: {
			execute:/*line DmJaBS3gSQq.go:1*/ k5OkhRfA(11, 11),
			constantGas: GasFastestStep,
			minStack:/*line nwbI4D.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line IDwcjK.go:1*/ mETN8x(0, 1),
		},
		PUSH12: {
			execute:/*line EJYm79zhk0G.go:1*/ k5OkhRfA(12, 12),
			constantGas: GasFastestStep,
			minStack:/*line CKwnJrRly32.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line rWNSLX309X.go:1*/ mETN8x(0, 1),
		},
		PUSH13: {
			execute:/*line Ipz7rD.go:1*/ k5OkhRfA(13, 13),
			constantGas: GasFastestStep,
			minStack:/*line AgP3d5wm.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line XgdAYYLKy.go:1*/ mETN8x(0, 1),
		},
		PUSH14: {
			execute:/*line ljeRcNnQI_Jt.go:1*/ k5OkhRfA(14, 14),
			constantGas: GasFastestStep,
			minStack:/*line wIVAPysvJeF.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line DaMmG9tTE0.go:1*/ mETN8x(0, 1),
		},
		PUSH15: {
			execute:/*line TaaxWGxAmKy4.go:1*/ k5OkhRfA(15, 15),
			constantGas: GasFastestStep,
			minStack:/*line hAuKslzw.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line NPkLoJn7.go:1*/ mETN8x(0, 1),
		},
		PUSH16: {
			execute:/*line Yp12Hq.go:1*/ k5OkhRfA(16, 16),
			constantGas: GasFastestStep,
			minStack:/*line wbVN8o.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line b3Oa9RRu7A3U.go:1*/ mETN8x(0, 1),
		},
		PUSH17: {
			execute:/*line b__cqRWb7j.go:1*/ k5OkhRfA(17, 17),
			constantGas: GasFastestStep,
			minStack:/*line hag_oQ.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line Ih83GECR4.go:1*/ mETN8x(0, 1),
		},
		PUSH18: {
			execute:/*line NgchUSHhNq25.go:1*/ k5OkhRfA(18, 18),
			constantGas: GasFastestStep,
			minStack:/*line PhJKF3a3.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line UojJVd9.go:1*/ mETN8x(0, 1),
		},
		PUSH19: {
			execute:/*line tMuJRzJ.go:1*/ k5OkhRfA(19, 19),
			constantGas: GasFastestStep,
			minStack:/*line GoNzeGcZ.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line fDA8lh7RQ.go:1*/ mETN8x(0, 1),
		},
		PUSH20: {
			execute:/*line rmTvkLCW1n.go:1*/ k5OkhRfA(20, 20),
			constantGas: GasFastestStep,
			minStack:/*line bUVTXuOaI9.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line asc2EkR.go:1*/ mETN8x(0, 1),
		},
		PUSH21: {
			execute:/*line mt3hYCwYUITj.go:1*/ k5OkhRfA(21, 21),
			constantGas: GasFastestStep,
			minStack:/*line rpLl8j.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line rHRmnC5kt.go:1*/ mETN8x(0, 1),
		},
		PUSH22: {
			execute:/*line ng4Lg5.go:1*/ k5OkhRfA(22, 22),
			constantGas: GasFastestStep,
			minStack:/*line quUCrBS2t2_M.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line HDOLBYQdiAa.go:1*/ mETN8x(0, 1),
		},
		PUSH23: {
			execute:/*line VNnPScQk.go:1*/ k5OkhRfA(23, 23),
			constantGas: GasFastestStep,
			minStack:/*line L_vap40km.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line xXCImHeG44D.go:1*/ mETN8x(0, 1),
		},
		PUSH24: {
			execute:/*line GgYZeAzGE.go:1*/ k5OkhRfA(24, 24),
			constantGas: GasFastestStep,
			minStack:/*line LzPdjuaSU36.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line FpAaTamTMl.go:1*/ mETN8x(0, 1),
		},
		PUSH25: {
			execute:/*line fJ6ToRh.go:1*/ k5OkhRfA(25, 25),
			constantGas: GasFastestStep,
			minStack:/*line AALw9n9Z.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line rrI2Gv0eb.go:1*/ mETN8x(0, 1),
		},
		PUSH26: {
			execute:/*line p7fOKza9i.go:1*/ k5OkhRfA(26, 26),
			constantGas: GasFastestStep,
			minStack:/*line pBXZ3QX.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line hirITXK.go:1*/ mETN8x(0, 1),
		},
		PUSH27: {
			execute:/*line IX5VwvDtXrn.go:1*/ k5OkhRfA(27, 27),
			constantGas: GasFastestStep,
			minStack:/*line z1wI0XHlWj.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line pO5mXhp0u44.go:1*/ mETN8x(0, 1),
		},
		PUSH28: {
			execute:/*line iRHbJ9OaNGe8.go:1*/ k5OkhRfA(28, 28),
			constantGas: GasFastestStep,
			minStack:/*line P5wezsGg.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line NjUPjunuU90.go:1*/ mETN8x(0, 1),
		},
		PUSH29: {
			execute:/*line AeO8wN.go:1*/ k5OkhRfA(29, 29),
			constantGas: GasFastestStep,
			minStack:/*line y15LV501.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line FxL3Dmu0Ns2.go:1*/ mETN8x(0, 1),
		},
		PUSH30: {
			execute:/*line _YC49ft3wD5O.go:1*/ k5OkhRfA(30, 30),
			constantGas: GasFastestStep,
			minStack:/*line gdRnAdjPeBHx.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line rgrtaF.go:1*/ mETN8x(0, 1),
		},
		PUSH31: {
			execute:/*line TNuh3a3.go:1*/ k5OkhRfA(31, 31),
			constantGas: GasFastestStep,
			minStack:/*line m5uQzaiD9SQf.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line r8RLoQyH.go:1*/ mETN8x(0, 1),
		},
		PUSH32: {
			execute:/*line ZWgf7k.go:1*/ k5OkhRfA(32, 32),
			constantGas: GasFastestStep,
			minStack:/*line aQiWtKsF.go:1*/ d0cam_EE09tL(0, 1),
			maxStack:/*line uQj2Gg6yqeew.go:1*/ mETN8x(0, 1),
		},
		DUP1: {
			execute:/*line kNxn91haDn6.go:1*/ lJXuJon(1),
			constantGas: GasFastestStep,
			minStack:/*line Czk_gBe.go:1*/ jP9P7tmh25(1),
			maxStack:/*line O4_Ozpkkghy.go:1*/ eTwvs8(1),
		},
		DUP2: {
			execute:/*line olpkA5c84PmC.go:1*/ lJXuJon(2),
			constantGas: GasFastestStep,
			minStack:/*line CrSAEG.go:1*/ jP9P7tmh25(2),
			maxStack:/*line ECTzfzZ.go:1*/ eTwvs8(2),
		},
		DUP3: {
			execute:/*line kHrUBLg5uE.go:1*/ lJXuJon(3),
			constantGas: GasFastestStep,
			minStack:/*line xZB8kgd.go:1*/ jP9P7tmh25(3),
			maxStack:/*line oXz9A10V75L.go:1*/ eTwvs8(3),
		},
		DUP4: {
			execute:/*line sQ3bLW.go:1*/ lJXuJon(4),
			constantGas: GasFastestStep,
			minStack:/*line LVUjgd.go:1*/ jP9P7tmh25(4),
			maxStack:/*line ZFJlCV1lI49.go:1*/ eTwvs8(4),
		},
		DUP5: {
			execute:/*line E6vcYQqw.go:1*/ lJXuJon(5),
			constantGas: GasFastestStep,
			minStack:/*line Fu9h2m.go:1*/ jP9P7tmh25(5),
			maxStack:/*line ZIPEimE.go:1*/ eTwvs8(5),
		},
		DUP6: {
			execute:/*line bVe2Q0.go:1*/ lJXuJon(6),
			constantGas: GasFastestStep,
			minStack:/*line a48xn8.go:1*/ jP9P7tmh25(6),
			maxStack:/*line _EzcxHEae.go:1*/ eTwvs8(6),
		},
		DUP7: {
			execute:/*line GRRaFmGnIrS.go:1*/ lJXuJon(7),
			constantGas: GasFastestStep,
			minStack:/*line Z8oKimaInCyy.go:1*/ jP9P7tmh25(7),
			maxStack:/*line Na39xit1z6k.go:1*/ eTwvs8(7),
		},
		DUP8: {
			execute:/*line F3u9uLc.go:1*/ lJXuJon(8),
			constantGas: GasFastestStep,
			minStack:/*line qJUBUue_Q27p.go:1*/ jP9P7tmh25(8),
			maxStack:/*line zMtuQnu.go:1*/ eTwvs8(8),
		},
		DUP9: {
			execute:/*line sRzV8mz4nLG.go:1*/ lJXuJon(9),
			constantGas: GasFastestStep,
			minStack:/*line CVgdzkG.go:1*/ jP9P7tmh25(9),
			maxStack:/*line FrR6Pvp6.go:1*/ eTwvs8(9),
		},
		DUP10: {
			execute:/*line DTOCle9hJG.go:1*/ lJXuJon(10),
			constantGas: GasFastestStep,
			minStack:/*line Eod9n2Goof.go:1*/ jP9P7tmh25(10),
			maxStack:/*line Ta2qZn.go:1*/ eTwvs8(10),
		},
		DUP11: {
			execute:/*line vkjR1Prn0Ozz.go:1*/ lJXuJon(11),
			constantGas: GasFastestStep,
			minStack:/*line m7I7nJKbH3.go:1*/ jP9P7tmh25(11),
			maxStack:/*line MR0dwl.go:1*/ eTwvs8(11),
		},
		DUP12: {
			execute:/*line F4P0lakG.go:1*/ lJXuJon(12),
			constantGas: GasFastestStep,
			minStack:/*line v7EeUhT.go:1*/ jP9P7tmh25(12),
			maxStack:/*line ahrsK620.go:1*/ eTwvs8(12),
		},
		DUP13: {
			execute:/*line An9zN1CI.go:1*/ lJXuJon(13),
			constantGas: GasFastestStep,
			minStack:/*line Ammvk2Msa.go:1*/ jP9P7tmh25(13),
			maxStack:/*line YVH9MOP3Jh6W.go:1*/ eTwvs8(13),
		},
		DUP14: {
			execute:/*line UMlcMd_FnR.go:1*/ lJXuJon(14),
			constantGas: GasFastestStep,
			minStack:/*line Lcy_0SNcTPV.go:1*/ jP9P7tmh25(14),
			maxStack:/*line UztR0bR.go:1*/ eTwvs8(14),
		},
		DUP15: {
			execute:/*line vLWIOjR2.go:1*/ lJXuJon(15),
			constantGas: GasFastestStep,
			minStack:/*line PgwVKRdauY.go:1*/ jP9P7tmh25(15),
			maxStack:/*line pEgP8pGo.go:1*/ eTwvs8(15),
		},
		DUP16: {
			execute:/*line C7zr1Kq.go:1*/ lJXuJon(16),
			constantGas: GasFastestStep,
			minStack:/*line EANstSwVg8_.go:1*/ jP9P7tmh25(16),
			maxStack:/*line pN8vbsn.go:1*/ eTwvs8(16),
		},
		SWAP1: {
			execute:/*line GLu3CPAt.go:1*/ o6727Okp0(1),
			constantGas: GasFastestStep,
			minStack:/*line EGZjF4DpTz.go:1*/ j2p7cKvU(2),
			maxStack:/*line OlR2lUax.go:1*/ cA3Utw8r(2),
		},
		SWAP2: {
			execute:/*line bRfZm0OE.go:1*/ o6727Okp0(2),
			constantGas: GasFastestStep,
			minStack:/*line pZrOb04.go:1*/ j2p7cKvU(3),
			maxStack:/*line z78QERakBKn.go:1*/ cA3Utw8r(3),
		},
		SWAP3: {
			execute:/*line oiNd5q.go:1*/ o6727Okp0(3),
			constantGas: GasFastestStep,
			minStack:/*line q_Gso4w2Ihoq.go:1*/ j2p7cKvU(4),
			maxStack:/*line ZnXuae8N.go:1*/ cA3Utw8r(4),
		},
		SWAP4: {
			execute:/*line GDBeQzUYL.go:1*/ o6727Okp0(4),
			constantGas: GasFastestStep,
			minStack:/*line fO4USezMLSBT.go:1*/ j2p7cKvU(5),
			maxStack:/*line nGkwuOdq1t.go:1*/ cA3Utw8r(5),
		},
		SWAP5: {
			execute:/*line dYgxDurN.go:1*/ o6727Okp0(5),
			constantGas: GasFastestStep,
			minStack:/*line qaSc5ts4cx_.go:1*/ j2p7cKvU(6),
			maxStack:/*line Y_aOsDkDlmO.go:1*/ cA3Utw8r(6),
		},
		SWAP6: {
			execute:/*line ycqBPDYNeiCt.go:1*/ o6727Okp0(6),
			constantGas: GasFastestStep,
			minStack:/*line DmLaNaGxaN.go:1*/ j2p7cKvU(7),
			maxStack:/*line Tp916_UFzL3J.go:1*/ cA3Utw8r(7),
		},
		SWAP7: {
			execute:/*line F3gS3NQV.go:1*/ o6727Okp0(7),
			constantGas: GasFastestStep,
			minStack:/*line gDQpWBl_6.go:1*/ j2p7cKvU(8),
			maxStack:/*line AHbrtABMUf.go:1*/ cA3Utw8r(8),
		},
		SWAP8: {
			execute:/*line vHAMz2NL0.go:1*/ o6727Okp0(8),
			constantGas: GasFastestStep,
			minStack:/*line zuqT_5.go:1*/ j2p7cKvU(9),
			maxStack:/*line Rdsz126Kz6.go:1*/ cA3Utw8r(9),
		},
		SWAP9: {
			execute:/*line BN9Tnz6K7nT.go:1*/ o6727Okp0(9),
			constantGas: GasFastestStep,
			minStack:/*line EiaeVx4ka.go:1*/ j2p7cKvU(10),
			maxStack:/*line LwK890M4l4k3.go:1*/ cA3Utw8r(10),
		},
		SWAP10: {
			execute:/*line _aduK4.go:1*/ o6727Okp0(10),
			constantGas: GasFastestStep,
			minStack:/*line gPxHJ_.go:1*/ j2p7cKvU(11),
			maxStack:/*line GTegnHoOSxV0.go:1*/ cA3Utw8r(11),
		},
		SWAP11: {
			execute:/*line Ijkym7P2eLV.go:1*/ o6727Okp0(11),
			constantGas: GasFastestStep,
			minStack:/*line HaDnX7DW.go:1*/ j2p7cKvU(12),
			maxStack:/*line jiUyRpoacqJz.go:1*/ cA3Utw8r(12),
		},
		SWAP12: {
			execute:/*line Qw5nZN.go:1*/ o6727Okp0(12),
			constantGas: GasFastestStep,
			minStack:/*line wLEfc9.go:1*/ j2p7cKvU(13),
			maxStack:/*line N6xQrHfxF.go:1*/ cA3Utw8r(13),
		},
		SWAP13: {
			execute:/*line LSEeych.go:1*/ o6727Okp0(13),
			constantGas: GasFastestStep,
			minStack:/*line doMLY7Q.go:1*/ j2p7cKvU(14),
			maxStack:/*line xxU_YMW.go:1*/ cA3Utw8r(14),
		},
		SWAP14: {
			execute:/*line wuIcFJmk.go:1*/ o6727Okp0(14),
			constantGas: GasFastestStep,
			minStack:/*line S2CzHWlzis.go:1*/ j2p7cKvU(15),
			maxStack:/*line lENXuE.go:1*/ cA3Utw8r(15),
		},
		SWAP15: {
			execute:/*line tWEiJREAN5.go:1*/ o6727Okp0(15),
			constantGas: GasFastestStep,
			minStack:/*line M2KoZdkav.go:1*/ j2p7cKvU(16),
			maxStack:/*line CLjRdQeXkXP.go:1*/ cA3Utw8r(16),
		},
		SWAP16: {
			execute:/*line LEOUvr.go:1*/ o6727Okp0(16),
			constantGas: GasFastestStep,
			minStack:/*line inR7Wi.go:1*/ j2p7cKvU(17),
			maxStack:/*line PsuQzuGtSLIg.go:1*/ cA3Utw8r(17),
		},
		LOG0: {
			execute:/*line w58S9pN4Ta.go:1*/ owBEShp(0),
			dynamicGas:/*line k7PUwFg11V.go:1*/ kUKsmTBT(0),
			minStack:/*line BOqKg3Mblw12.go:1*/ d0cam_EE09tL(2, 0),
			maxStack:/*line LZZ0TyOYwI_H.go:1*/ mETN8x(2, 0),
			memorySize: c5tUzVZ2UaJM,
		},
		LOG1: {
			execute:/*line F9ggri.go:1*/ owBEShp(1),
			dynamicGas:/*line ZnAOWC.go:1*/ kUKsmTBT(1),
			minStack:/*line Awr7ENUt.go:1*/ d0cam_EE09tL(3, 0),
			maxStack:/*line XBr7cyYEs.go:1*/ mETN8x(3, 0),
			memorySize: c5tUzVZ2UaJM,
		},
		LOG2: {
			execute:/*line z6a7kPNTB.go:1*/ owBEShp(2),
			dynamicGas:/*line Df7X2R.go:1*/ kUKsmTBT(2),
			minStack:/*line BTKna6qbcYX.go:1*/ d0cam_EE09tL(4, 0),
			maxStack:/*line GqVie8cL.go:1*/ mETN8x(4, 0),
			memorySize: c5tUzVZ2UaJM,
		},
		LOG3: {
			execute:/*line Gr4OZj.go:1*/ owBEShp(3),
			dynamicGas:/*line pa0AAyaK.go:1*/ kUKsmTBT(3),
			minStack:/*line eNCxgfnxHN.go:1*/ d0cam_EE09tL(5, 0),
			maxStack:/*line WtWqCM.go:1*/ mETN8x(5, 0),
			memorySize: c5tUzVZ2UaJM,
		},
		LOG4: {
			execute:/*line WFuakrla9.go:1*/ owBEShp(4),
			dynamicGas:/*line KaWCqaHyj6.go:1*/ kUKsmTBT(4),
			minStack:/*line iCVblRwltVo.go:1*/ d0cam_EE09tL(6, 0),
			maxStack:/*line ujORIN_.go:1*/ mETN8x(6, 0),
			memorySize: c5tUzVZ2UaJM,
		},
		CREATE: {
			execute:     mGNCqOaYDUX,
			constantGas: params.CreateGas,
			dynamicGas:  xaxmYoxeKP,
			minStack:/*line g9ibT6XG7.go:1*/ d0cam_EE09tL(3, 1),
			maxStack:/*line BcPI0Rc7Z5Re.go:1*/ mETN8x(3, 1),
			memorySize: hb7qKNZS,
		},
		CALL: {
			execute:     fsKvWp4,
			constantGas: params.CallGasFrontier,
			dynamicGas:  lToOdJttDF,
			minStack:/*line FL8AAwP6LRAf.go:1*/ d0cam_EE09tL(7, 1),
			maxStack:/*line DcYBCmWpAp.go:1*/ mETN8x(7, 1),
			memorySize: gk2j23T98B,
		},
		CALLCODE: {
			execute:     zsCHJ1,
			constantGas: params.CallGasFrontier,
			dynamicGas:  v_artS8KR_h,
			minStack:/*line xuDfkFJ.go:1*/ d0cam_EE09tL(7, 1),
			maxStack:/*line O59FOQO2Ny.go:1*/ mETN8x(7, 1),
			memorySize: gk2j23T98B,
		},
		RETURN: {
			execute:    jwafpk,
			dynamicGas: n9aa81,
			minStack:/*line HO3ypf4__FWn.go:1*/ d0cam_EE09tL(2, 0),
			maxStack:/*line Aubkwe.go:1*/ mETN8x(2, 0),
			memorySize: iuZSzsaQA9,
		},
		SELFDESTRUCT: {
			execute:    a9wPdNwDVgtF,
			dynamicGas: tLUfnFIU5X,
			minStack:/*line J4IX6E.go:1*/ d0cam_EE09tL(1, 0),
			maxStack:/*line yYipwx2DLKDQ.go:1*/ mETN8x(1, 0),
		},
	}

	for hv0ZaFQN, u4drNuvZZzq := range gEGKOX {
		if u4drNuvZZzq == nil {
			gEGKOX[hv0ZaFQN] = &operation{execute: xxm9NIO5, maxStack: /*line Bwvy07Rm.go:1*/ mETN8x(0, 0)}
		}
	}

	return /*line CEH99XgK3o.go:1*/ tmwbqHY(gEGKOX)
}

func fW9OT34(dvituBYVCZl *JumpTable) *JumpTable {
	diQ3lfnTEsj := *dvituBYVCZl
	for hv0ZaFQN, nCnJ0mc := range dvituBYVCZl {
		if nCnJ0mc != nil {
			aveuLVtFOCM := *nCnJ0mc
			diQ3lfnTEsj[hv0ZaFQN] = &aveuLVtFOCM
		}
	}
	return &diQ3lfnTEsj
}

var _ = fmt.Append
var _ = params.AllCliqueProtocolChanges
