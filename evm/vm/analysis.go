//line :1
package vm

const (
	set2BitsMask = /*line V8zCFmy3uo8.go:1*/ uint16(0b11)
	set3BitsMask = /*line a9DM7ZD.go:1*/ uint16(0b111)
	set4BitsMask = /*line vzWKNWABLZ.go:1*/ uint16(0b1111)
	set5BitsMask = /*line CmrDbhPI.go:1*/ uint16(0b1_1111)
	set6BitsMask = /*line ojPVUEhcf6uP.go:1*/ uint16(0b11_1111)
	set7BitsMask = /*line oyVAK_.go:1*/ uint16(0b111_1111)
)

type gYd8QoJh3 []byte

func (v3P7bJLmQ_ gYd8QoJh3) dFQMCOSkt(tzxj_ga uint64) {
	v3P7bJLmQ_[tzxj_ga/8] |= 1 << (tzxj_ga % 8)
}

func (v3P7bJLmQ_ gYd8QoJh3) mj6sG6QRVBkX(jQsgbIsa uint16, tzxj_ga uint64) {
	dbalZotRLEH := jQsgbIsa << (tzxj_ga % 8)
	v3P7bJLmQ_[tzxj_ga/8] |= /*line FYYIp3tfap0.go:1*/ byte(dbalZotRLEH)
	if l1KtNnRwliPB := /*line H4Z5GCavELNV.go:1*/ byte(dbalZotRLEH >> 8); l1KtNnRwliPB != 0 {
		v3P7bJLmQ_[tzxj_ga/8+1] = l1KtNnRwliPB
	}
}

func (v3P7bJLmQ_ gYd8QoJh3) tajEg0(tzxj_ga uint64) {
	dbalZotRLEH := /*line CsxNgVzdM.go:1*/ byte(0xFF << (tzxj_ga % 8))
	v3P7bJLmQ_[tzxj_ga/8] |= dbalZotRLEH
	v3P7bJLmQ_[tzxj_ga/8+1] = ^dbalZotRLEH
}

func (v3P7bJLmQ_ gYd8QoJh3) d4YwcX(tzxj_ga uint64) {
	dbalZotRLEH := /*line FXoHaNtA2a.go:1*/ byte(0xFF << (tzxj_ga % 8))
	v3P7bJLmQ_[tzxj_ga/8] |= dbalZotRLEH
	v3P7bJLmQ_[tzxj_ga/8+1] = 0xFF
	v3P7bJLmQ_[tzxj_ga/8+2] = ^dbalZotRLEH
}

func (v3P7bJLmQ_ *gYd8QoJh3) uR1Fotw(tzxj_ga uint64) bool {
	return (((*v3P7bJLmQ_)[tzxj_ga/8] >> (tzxj_ga % 8)) & 1) == 0
}

func t27hPJQn(euva86f0e []byte) gYd8QoJh3 {

	v3P7bJLmQ_ := /*line AL1EnIpb6.go:1*/ make(gYd8QoJh3 /*line Wbsu2a.go:1*/, len(euva86f0e)/8+1+4)
	return /*line ZDXV_8S7.go:1*/ ih0B5koZE8F(euva86f0e, v3P7bJLmQ_)
}

func ih0B5koZE8F(euva86f0e, v3P7bJLmQ_ gYd8QoJh3) gYd8QoJh3 {
	for eAorajeVy084 := /*line HZlgGZ8RlF31.go:1*/ uint64(0); eAorajeVy084 < /*line BKoqW6JW3.go:1*/ uint64( /*line QegUNLLneM7.go:1*/ len(euva86f0e)); {
		nCnJ0mc := /*line tjIDNkMCMGc.go:1*/ LxosJe8(euva86f0e[eAorajeVy084])
		eAorajeVy084++
		if /*line zlMuAS.go:1*/ int8(nCnJ0mc) < /*line rBBpTBHA.go:1*/ int8(PUSH1) {
			continue
		}
		uWwFMRCpQ := nCnJ0mc - PUSH1 + 1
		if uWwFMRCpQ >= 8 {
			for ; uWwFMRCpQ >= 16; uWwFMRCpQ -= 16 {
				/*line EMvO8F1m8W.go:1*/ v3P7bJLmQ_.d4YwcX(eAorajeVy084)
				eAorajeVy084 += 16
			}
			for ; uWwFMRCpQ >= 8; uWwFMRCpQ -= 8 {
				/*line B5OeLnGkLS.go:1*/ v3P7bJLmQ_.tajEg0(eAorajeVy084)
				eAorajeVy084 += 8
			}
		}
		switch uWwFMRCpQ {
		case 1:
			/*line IadMHE7m.go:1*/ v3P7bJLmQ_.dFQMCOSkt(eAorajeVy084)
			eAorajeVy084 += 1
		case 2:
			/*line dqbqDai.go:1*/ v3P7bJLmQ_.mj6sG6QRVBkX(set2BitsMask, eAorajeVy084)
			eAorajeVy084 += 2
		case 3:
			/*line vMVD_a.go:1*/ v3P7bJLmQ_.mj6sG6QRVBkX(set3BitsMask, eAorajeVy084)
			eAorajeVy084 += 3
		case 4:
			/*line HsmvGxPK.go:1*/ v3P7bJLmQ_.mj6sG6QRVBkX(set4BitsMask, eAorajeVy084)
			eAorajeVy084 += 4
		case 5:
			/*line EZwdUXz7y.go:1*/ v3P7bJLmQ_.mj6sG6QRVBkX(set5BitsMask, eAorajeVy084)
			eAorajeVy084 += 5
		case 6:
			/*line AmJv0oUNiD4Y.go:1*/ v3P7bJLmQ_.mj6sG6QRVBkX(set6BitsMask, eAorajeVy084)
			eAorajeVy084 += 6
		case 7:
			/*line BQ6iAI_Or.go:1*/ v3P7bJLmQ_.mj6sG6QRVBkX(set7BitsMask, eAorajeVy084)
			eAorajeVy084 += 7
		}
	}
	return v3P7bJLmQ_
}
