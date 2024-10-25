//line :1
package types

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type pL0OquHHL interface {
	Bytes() []byte
}

const (
	BloomByteLength = 256

	BloomBitLength = 8 * BloomByteLength
)

type Bloom [BloomByteLength]byte

func Jyv4_JiBl3(aYao5YS []byte) Bloom {
	var zu6VnYasROCR Bloom
	/*line qxOPE14OwT.go:1*/ zu6VnYasROCR.SetBytes(aYao5YS)
	return zu6VnYasROCR
}

func (aYao5YS *Bloom) SetBytes(xUR2t9 []byte) {
	if /*line WVrio6ZAGhV7.go:1*/ len(aYao5YS) < /*line NHP4qd8.go:1*/ len(xUR2t9) {
		/*line SQH8LpcPJ.go:1*/ panic( /*line yJ1uGEY.go:1*/ fmt.Sprintf(func() /*line FbSqDavSJLw.go:1*/ string {
			data := /*line FbSqDavSJLw.go:1*/ make([]byte, 0, 26)
			i := 4
			decryptKey := 59
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					i = 1
					data = /*line FbSqDavSJLw.go:1*/ append(data, "\xa9\xaa\xa9]"...,
					)
				case 11:
					i = 5
					data = /*line FbSqDavSJLw.go:1*/ append(data, "W\x8a"...,
					)
				case 4:
					i = 7
					data = /*line FbSqDavSJLw.go:1*/ append(data, "\x9a\xa5"...,
					)
				case 0:
					data = /*line FbSqDavSJLw.go:1*/ append(data, 165)
					i = 10
				case 6:
					i = 2
					data = /*line FbSqDavSJLw.go:1*/ append(data, 132)
				case 1:
					data = /*line FbSqDavSJLw.go:1*/ append(data, "\xa0\xb8\xa4\x96"...,
					)
					i = 0
				case 5:
					data = /*line FbSqDavSJLw.go:1*/ append(data, "\x92\x91K"...,
					)
					i = 8
				case 2:
					for y := range data {
						data[y] = data[y] - /*line FbSqDavSJLw.go:1*/ byte(decryptKey^y)
					}
					i = 3
				case 10:
					data = /*line FbSqDavSJLw.go:1*/ append(data, "S\xa8"...,
					)
					i = 9
				case 9:
					i = 11
					data = /*line FbSqDavSJLw.go:1*/ append(data, "\xa4\xa5"...,
					)
				case 8:
					data = /*line FbSqDavSJLw.go:1*/ append(data, "Q\x91NT"...,
					)
					i = 6
				}
			}
			return /*line FbSqDavSJLw.go:1*/ string(data)
		}(), /*line TV5x92c.go:1*/ len(aYao5YS) /*line _E45vQTD9.go:1*/, len(xUR2t9)))
	}
	/*line vLfpKBKWp.go:1*/ copy(aYao5YS[BloomByteLength- /*line B9VfU5j.go:1*/ len(xUR2t9):], xUR2t9)
}

func (aYao5YS *Bloom) Add(xUR2t9 []byte) {
	/*line DjgndL0VZup.go:1*/ aYao5YS.yFS1Pol34fab(xUR2t9 /*line T4CqD5AK.go:1*/, make([]byte, 6))
}

func (aYao5YS *Bloom) yFS1Pol34fab(xUR2t9 []byte, djAcxNMM5s6 []byte) {
	h6YdAZ8xW2, e8AOUP, hL_F8AVD6cw, gr5o21jllI, dqQHaoy9a, gac9cdn := /*line Mc3XJO.go:1*/ jqrB6T6q0(xUR2t9, djAcxNMM5s6)
	aYao5YS[h6YdAZ8xW2] |= e8AOUP
	aYao5YS[hL_F8AVD6cw] |= gr5o21jllI
	aYao5YS[dqQHaoy9a] |= gac9cdn
}

func (aYao5YS Bloom) Big() *big.Int {
	return /*line ASpsIZPSYe.go:1*/ new(big.Int).SetBytes(aYao5YS[:])
}

func (aYao5YS Bloom) Bytes() []byte {
	return aYao5YS[:]
}

func (aYao5YS Bloom) Test(aK8lJj8mirv []byte) bool {
	h6YdAZ8xW2, e8AOUP, hL_F8AVD6cw, gr5o21jllI, dqQHaoy9a, gac9cdn := /*line skGkSia.go:1*/ jqrB6T6q0(aK8lJj8mirv /*line IL9aMy.go:1*/, make([]byte, 6))
	return e8AOUP == e8AOUP&aYao5YS[h6YdAZ8xW2] &&
		gr5o21jllI == gr5o21jllI&aYao5YS[hL_F8AVD6cw] &&
		gac9cdn == gac9cdn&aYao5YS[dqQHaoy9a]
}

func (aYao5YS Bloom) MarshalText() ([]byte, error) {
	return /*line l8kYfRrwFr.go:1*/ hexutil.Bytes(aYao5YS[:]).MarshalText()
}

func (aYao5YS *Bloom) UnmarshalText(uzD2Up []byte) error {
	return /*line HR_4CrSq4qPx.go:1*/ hexutil.UnmarshalFixedText("Bloom", uzD2Up, aYao5YS[:])
}

func T9aqYLj(zCSgFJubD2 Receipts) Bloom {
	djAcxNMM5s6 := /*line FG_69NPg8444.go:1*/ make([]byte, 6)
	var nbNfA8oE Bloom
	for _, hCHwHR1v := range zCSgFJubD2 {
		for _, iEDnq_jYGc := range hCHwHR1v.Logs {
			/*line JZR8M9v.go:1*/ nbNfA8oE.yFS1Pol34fab( /*line n33BHaUXXDUu.go:1*/ iEDnq_jYGc.Address.Bytes(), djAcxNMM5s6)
			for _, aYao5YS := range iEDnq_jYGc.Topics {
				/*line UCOeTdUcsFD.go:1*/ nbNfA8oE.yFS1Pol34fab(aYao5YS[:], djAcxNMM5s6)
			}
		}
	}
	return nbNfA8oE
}

func VhSK4u(jdmmHtmVa []*Log) []byte {
	djAcxNMM5s6 := /*line FDmElGFkuRE1.go:1*/ make([]byte, 6)
	var nbNfA8oE Bloom
	for _, iEDnq_jYGc := range jdmmHtmVa {
		/*line zqzP8Iw.go:1*/ nbNfA8oE.yFS1Pol34fab( /*line ag7hrIyDZc.go:1*/ iEDnq_jYGc.Address.Bytes(), djAcxNMM5s6)
		for _, aYao5YS := range iEDnq_jYGc.Topics {
			/*line FXQRgr6EaMta.go:1*/ nbNfA8oE.yFS1Pol34fab(aYao5YS[:], djAcxNMM5s6)
		}
	}
	return nbNfA8oE[:]
}

func TMKxUWaefJd2(bgjAklkHvJWu []byte) []byte {
	var aYao5YS Bloom
	/*line RdcKFMyrS.go:1*/ aYao5YS.SetBytes(bgjAklkHvJWu)
	return /*line ZYSFZf.go:1*/ aYao5YS.Bytes()
}

func jqrB6T6q0(bgjAklkHvJWu []byte, xaQ122eZaI []byte) (uint, byte, uint, byte, uint, byte) {
	cH4lnex0Ibo := /*line H7sF4xtMNZx.go:1*/ f3SPqrhaMlU.Get().(crypto.KeccakState)
	/*line PHCF9IwGV_.go:1*/ cH4lnex0Ibo.Reset()
	/*line KDRnWPaO2k.go:1*/ cH4lnex0Ibo.Write(bgjAklkHvJWu)
	/*line jmFqDb75.go:1*/ cH4lnex0Ibo.Read(xaQ122eZaI)
	/*line yfcC3_y04.go:1*/ f3SPqrhaMlU.Put(cH4lnex0Ibo)

	e8AOUP := /*line Cf9CU5.go:1*/ byte(1 << (xaQ122eZaI[1] & 0x7))
	gr5o21jllI := /*line piS2CmVdSta5.go:1*/ byte(1 << (xaQ122eZaI[3] & 0x7))
	gac9cdn := /*line IvVjf2An5pYn.go:1*/ byte(1 << (xaQ122eZaI[5] & 0x7))

	h6YdAZ8xW2 := BloomByteLength - /*line IxWKiZQX9K.go:1*/ uint(( /*line yXSrTm.go:1*/ binary.BigEndian.Uint16(xaQ122eZaI)&0x7ff)>>3) - 1
	hL_F8AVD6cw := BloomByteLength - /*line fEwND7a.go:1*/ uint(( /*line c_bAOaqs.go:1*/ binary.BigEndian.Uint16(xaQ122eZaI[2:])&0x7ff)>>3) - 1
	dqQHaoy9a := BloomByteLength - /*line yfK3hpM6VTC.go:1*/ uint(( /*line AM4M0voz8T.go:1*/ binary.BigEndian.Uint16(xaQ122eZaI[4:])&0x7ff)>>3) - 1

	return h6YdAZ8xW2, e8AOUP, hL_F8AVD6cw, gr5o21jllI, dqQHaoy9a, gac9cdn
}

func Uo3diyL(nbNfA8oE Bloom, aK8lJj8mirv pL0OquHHL) bool {
	return /*line tHad0WxY.go:1*/ nbNfA8oE.Test( /*line jmu7xA.go:1*/ aK8lJj8mirv.Bytes())
}

var _ = binary.Append
var _ = fmt.Append

const _ = big.Above

var _ hexutil.Big
var _ = crypto.CompressPubkey
