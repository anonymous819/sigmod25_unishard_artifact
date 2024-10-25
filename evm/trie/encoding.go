//line :1
package trie

func pSumbut7Z(erLmPgwaEDk []byte) []byte {
	a9hfr8Gsl9C := /*line g0Wf3YB.go:1*/ byte(0)
	if /*line vq3duwD.go:1*/ k7wsvmNwK(erLmPgwaEDk) {
		a9hfr8Gsl9C = 1
		erLmPgwaEDk = erLmPgwaEDk[: /*line VQtTqzu.go:1*/ len(erLmPgwaEDk)-1]
	}
	m2t4GBUGH := /*line LAaB0pRE.go:1*/ make([]byte /*line aiKWdVWtNl.go:1*/, len(erLmPgwaEDk)/2+1)
	m2t4GBUGH[0] = a9hfr8Gsl9C << 5
	if /*line gI9pLl.go:1*/ len(erLmPgwaEDk)&1 == 1 {
		m2t4GBUGH[0] |= 1 << 4
		m2t4GBUGH[0] |= erLmPgwaEDk[0]
		erLmPgwaEDk = erLmPgwaEDk[1:]
	}
	/*line GnPXPCTJirkU.go:1*/ oIX_bO0aP5L(erLmPgwaEDk, m2t4GBUGH[1:])
	return m2t4GBUGH
}

func rs0SaPUuBe(erLmPgwaEDk []byte) []byte {
	var (
		tXOrlLdbmrh = /*line ZFTwMNf.go:1*/ len(erLmPgwaEDk)
		jXd1eY1a    = /*line HdLPdqfjb.go:1*/ byte(0)
	)

	if tXOrlLdbmrh > 0 && erLmPgwaEDk[tXOrlLdbmrh-1] == 16 {
		jXd1eY1a = 1 << 5
		tXOrlLdbmrh--
	}
	var (
		bFYlZEsw0 = tXOrlLdbmrh/2 + 1
		qF3Tq4X8x = 0
		hKIvF2    = 1
	)
	if tXOrlLdbmrh&1 == 1 {
		jXd1eY1a |= 1 << 4
		jXd1eY1a |= erLmPgwaEDk[0]
		qF3Tq4X8x++
	}
	for ; qF3Tq4X8x < tXOrlLdbmrh; hKIvF2, qF3Tq4X8x = hKIvF2+1, qF3Tq4X8x+2 {
		erLmPgwaEDk[hKIvF2] = erLmPgwaEDk[qF3Tq4X8x]<<4 | erLmPgwaEDk[qF3Tq4X8x+1]
	}
	erLmPgwaEDk[0] = jXd1eY1a
	return erLmPgwaEDk[:bFYlZEsw0]
}

func o9hF7yS(wu0pNs []byte) []byte {
	if /*line n4RNVSUM.go:1*/ len(wu0pNs) == 0 {
		return wu0pNs
	}
	gPBv_0Ab := /*line D80cO8bNLve2.go:1*/ nd12pU880Z8(wu0pNs)

	if gPBv_0Ab[0] < 2 {
		gPBv_0Ab = gPBv_0Ab[: /*line eePTTHH.go:1*/ len(gPBv_0Ab)-1]
	}

	yHwtlOAJ := 2 - gPBv_0Ab[0]&1
	return gPBv_0Ab[yHwtlOAJ:]
}

func nd12pU880Z8(vOV6g1 []byte) []byte {
	aHj7_4 := /*line wcb2J4R6g.go:1*/ len(vOV6g1)*2 + 1
	var w5bfDSBs_Z = /*line t4gZX8YPvS.go:1*/ make([]byte, aHj7_4)
	for q2u2020KD6, mZsNbw0 := range vOV6g1 {
		w5bfDSBs_Z[q2u2020KD6*2] = mZsNbw0 / 16
		w5bfDSBs_Z[q2u2020KD6*2+1] = mZsNbw0 % 16
	}
	w5bfDSBs_Z[aHj7_4-1] = 16
	return w5bfDSBs_Z
}

func iZl0DAYMB(erLmPgwaEDk []byte) []byte {
	if /*line cocta3Yt8K.go:1*/ k7wsvmNwK(erLmPgwaEDk) {
		erLmPgwaEDk = erLmPgwaEDk[: /*line FCNQ4P2.go:1*/ len(erLmPgwaEDk)-1]
	}
	if /*line I3ti08SjhOp.go:1*/ len(erLmPgwaEDk)&1 != 0 {
		/*line _unXEyCQ.go:1*/ panic(func() /*line ruFkdu.go:1*/ string {
			data := [] /*line ruFkdu.go:1*/ byte("ZSn't\x16kK`vo-\x17a\x7f\x03x.keX 1f\x1cqid[Zo22Wa")
			positions := [...]byte{10, 15, 28, 22, 22, 0, 6, 20, 10, 7, 25, 22, 30, 14, 31, 34, 20, 1, 26, 0, 24, 5, 5, 34, 29, 12, 30, 0, 29, 32, 33, 33, 14, 6, 17, 8, 7, 13, 10, 25, 7, 1, 31, 11}
			for i := 0; i < 44; i += 2 {
				localKey := /*line ruFkdu.go:1*/ byte(i) + /*line ruFkdu.go:1*/ byte(positions[i]^positions[i+1]) + 5
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return /*line ruFkdu.go:1*/ string(data)
		}())
	}
	lhQWH7m := /*line puePifBuQWaK.go:1*/ make([]byte /*line EoqHnVpKTdo.go:1*/, len(erLmPgwaEDk)/2)
	/*line FiqM08o4JUNI.go:1*/ oIX_bO0aP5L(erLmPgwaEDk, lhQWH7m)
	return lhQWH7m
}

func oIX_bO0aP5L(w5bfDSBs_Z []byte, sNrBzyUXd6av []byte) {
	for hKIvF2, qF3Tq4X8x := 0, 0; qF3Tq4X8x < /*line EjuAHE7MfO.go:1*/ len(w5bfDSBs_Z); hKIvF2, qF3Tq4X8x = hKIvF2+1, qF3Tq4X8x+2 {
		sNrBzyUXd6av[hKIvF2] = w5bfDSBs_Z[qF3Tq4X8x]<<4 | w5bfDSBs_Z[qF3Tq4X8x+1]
	}
}

func e0augA(vXUk1wf, mZsNbw0 []byte) int {
	var q2u2020KD6, cewciTKe5 = 0 /*line EASqkUPb6cK7.go:1*/, len(vXUk1wf)
	if /*line LC2GG1v1cUYh.go:1*/ len(mZsNbw0) < cewciTKe5 {
		cewciTKe5 = /*line rPX46_.go:1*/ len(mZsNbw0)
	}
	for ; q2u2020KD6 < cewciTKe5; q2u2020KD6++ {
		if vXUk1wf[q2u2020KD6] != mZsNbw0[q2u2020KD6] {
			break
		}
	}
	return q2u2020KD6
}

func k7wsvmNwK(mvwKyZFJ6 []byte) bool {
	return /*line rexrth.go:1*/ len(mvwKyZFJ6) > 0 && mvwKyZFJ6[ /*line sMh2oH7Z.go:1*/ len(mvwKyZFJ6)-1] == 16
}
