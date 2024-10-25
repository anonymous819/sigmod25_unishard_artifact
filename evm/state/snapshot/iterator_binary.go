//line :1
package snapshot

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

type aIOlc0qZvA struct {
	pgz9u076x_Q  FcPREnURc
	aFrRKX       FcPREnURc
	ccteAaL6V    bool
	h9AkahsaXXlY bool
	gwoHDzL7sp   bool
	mZ_uN01xOx   common.Hash
	dE4DeenWH    common.Hash
	ieH_2wJvG    error
}

func (hbsnnIi *diffLayer) smV5bw0OO() FcPREnURc {
	f3dWTs7, jdkNTRyBJ := hbsnnIi.parent.(*diffLayer)
	if !jdkNTRyBJ {
		faCFaIZW := &aIOlc0qZvA{
			pgz9u076x_Q:/*line crX27R0.go:1*/ hbsnnIi.AccountIterator(common.Hash{}),
			aFrRKX:/*line HXaVBVJd.go:1*/ hbsnnIi.Parent().AccountIterator(common.Hash{}),
			gwoHDzL7sp: true,
		}
		faCFaIZW.ccteAaL6V = ! /*line JzxJ29kQ2ez.go:1*/ faCFaIZW.pgz9u076x_Q.Next()
		faCFaIZW.h9AkahsaXXlY = ! /*line _PaWfu.go:1*/ faCFaIZW.aFrRKX.Next()
		return faCFaIZW
	}
	faCFaIZW := &aIOlc0qZvA{
		pgz9u076x_Q:/*line GPLRtyt.go:1*/ hbsnnIi.AccountIterator(common.Hash{}),
		aFrRKX:/*line CVNmX2Oya.go:1*/ f3dWTs7.smV5bw0OO(),
		gwoHDzL7sp: true,
	}
	faCFaIZW.ccteAaL6V = ! /*line IcqilyQPwsT_.go:1*/ faCFaIZW.pgz9u076x_Q.Next()
	faCFaIZW.h9AkahsaXXlY = ! /*line aFodF95yj.go:1*/ faCFaIZW.aFrRKX.Next()
	return faCFaIZW
}

func (hbsnnIi *diffLayer) z4rQNTf6c(evzmhL1 common.Hash) FcPREnURc {
	f3dWTs7, jdkNTRyBJ := hbsnnIi.parent.(*diffLayer)
	if !jdkNTRyBJ {

		rmHPuz8jT, oiHhcQbh9XAf := /*line B8sdbo.go:1*/ hbsnnIi.StorageIterator(evzmhL1, common.Hash{})
		if oiHhcQbh9XAf {
			faCFaIZW := &aIOlc0qZvA{
				pgz9u076x_Q: rmHPuz8jT,
				dE4DeenWH:   evzmhL1,
			}
			faCFaIZW.ccteAaL6V = ! /*line e1NJ5P3SgaQw.go:1*/ faCFaIZW.pgz9u076x_Q.Next()
			faCFaIZW.h9AkahsaXXlY = true
			return faCFaIZW
		}

		zfco8dsA, _ := /*line TfXaI8L2p.go:1*/ hbsnnIi.Parent().StorageIterator(evzmhL1, common.Hash{})
		faCFaIZW := &aIOlc0qZvA{
			pgz9u076x_Q: rmHPuz8jT,
			aFrRKX:      zfco8dsA,
			dE4DeenWH:   evzmhL1,
		}
		faCFaIZW.ccteAaL6V = ! /*line Ce0LGCc.go:1*/ faCFaIZW.pgz9u076x_Q.Next()
		faCFaIZW.h9AkahsaXXlY = ! /*line Fir_a2gj.go:1*/ faCFaIZW.aFrRKX.Next()
		return faCFaIZW
	}

	rmHPuz8jT, oiHhcQbh9XAf := /*line f3M9QytDJ.go:1*/ hbsnnIi.StorageIterator(evzmhL1, common.Hash{})
	if oiHhcQbh9XAf {
		faCFaIZW := &aIOlc0qZvA{
			pgz9u076x_Q: rmHPuz8jT,
			dE4DeenWH:   evzmhL1,
		}
		faCFaIZW.ccteAaL6V = ! /*line VYXVvPb.go:1*/ faCFaIZW.pgz9u076x_Q.Next()
		faCFaIZW.h9AkahsaXXlY = true
		return faCFaIZW
	}
	faCFaIZW := &aIOlc0qZvA{
		pgz9u076x_Q: rmHPuz8jT,
		aFrRKX:/*line PV6b9mO9.go:1*/ f3dWTs7.z4rQNTf6c(evzmhL1),
		dE4DeenWH: evzmhL1,
	}
	faCFaIZW.ccteAaL6V = ! /*line ZNbJYa9E3.go:1*/ faCFaIZW.pgz9u076x_Q.Next()
	faCFaIZW.h9AkahsaXXlY = ! /*line gIgP960NZ.go:1*/ faCFaIZW.aFrRKX.Next()
	return faCFaIZW
}

func (cu8ZKpmK *aIOlc0qZvA) Next() bool {
	if cu8ZKpmK.ccteAaL6V && cu8ZKpmK.h9AkahsaXXlY {
		return false
	}
first:
	if cu8ZKpmK.ccteAaL6V {
		cu8ZKpmK.mZ_uN01xOx = /*line XxbjAem.go:1*/ cu8ZKpmK.aFrRKX.Hash()
		cu8ZKpmK.h9AkahsaXXlY = ! /*line UHnHbmLRv2tP.go:1*/ cu8ZKpmK.aFrRKX.Next()
		return true
	}
	if cu8ZKpmK.h9AkahsaXXlY {
		cu8ZKpmK.mZ_uN01xOx = /*line W7QJD8y5W.go:1*/ cu8ZKpmK.pgz9u076x_Q.Hash()
		cu8ZKpmK.ccteAaL6V = ! /*line AktSdwov.go:1*/ cu8ZKpmK.pgz9u076x_Q.Next()
		return true
	}
	kodNc4v, rs_rY6 := /*line waWuVoi2zu.go:1*/ cu8ZKpmK.pgz9u076x_Q.Hash() /*line eqr4zGDWdwdd.go:1*/, cu8ZKpmK.aFrRKX.Hash()
	if wMCMWH9Tfvzo := /*line KIwiyS.go:1*/ bytes.Compare(kodNc4v[:], rs_rY6[:]); wMCMWH9Tfvzo < 0 {
		cu8ZKpmK.ccteAaL6V = ! /*line BNvPziAyAUJ.go:1*/ cu8ZKpmK.pgz9u076x_Q.Next()
		cu8ZKpmK.mZ_uN01xOx = kodNc4v
		return true
	} else if wMCMWH9Tfvzo == 0 {

		cu8ZKpmK.ccteAaL6V = ! /*line YHBf9bfACY.go:1*/ cu8ZKpmK.pgz9u076x_Q.Next()
		goto first
	}
	cu8ZKpmK.h9AkahsaXXlY = ! /*line mWmzmvP0CE6.go:1*/ cu8ZKpmK.aFrRKX.Next()
	cu8ZKpmK.mZ_uN01xOx = rs_rY6
	return true
}

func (cu8ZKpmK *aIOlc0qZvA) Error() error {
	return cu8ZKpmK.ieH_2wJvG
}

func (cu8ZKpmK *aIOlc0qZvA) Hash() common.Hash {
	return cu8ZKpmK.mZ_uN01xOx
}

func (cu8ZKpmK *aIOlc0qZvA) Account() []byte {
	if !cu8ZKpmK.gwoHDzL7sp {
		return nil
	}

	b1jafJL, chyZ8Q := /*line o1GnMxlTQb.go:1*/ cu8ZKpmK.pgz9u076x_Q.(*diffAccountIterator).layer.AccountRLP(cu8ZKpmK.mZ_uN01xOx)
	if chyZ8Q != nil {
		cu8ZKpmK.ieH_2wJvG = chyZ8Q
		return nil
	}
	return b1jafJL
}

func (cu8ZKpmK *aIOlc0qZvA) Slot() []byte {
	if cu8ZKpmK.gwoHDzL7sp {
		return nil
	}
	b1jafJL, chyZ8Q := /*line Ftamic.go:1*/ cu8ZKpmK.pgz9u076x_Q.(*diffStorageIterator).layer.Storage(cu8ZKpmK.dE4DeenWH, cu8ZKpmK.mZ_uN01xOx)
	if chyZ8Q != nil {
		cu8ZKpmK.ieH_2wJvG = chyZ8Q
		return nil
	}
	return b1jafJL
}

func (cu8ZKpmK *aIOlc0qZvA) Release() {
	/*line J2yKkeSmMUiU.go:1*/ cu8ZKpmK.pgz9u076x_Q.Release()
	/*line A2npEWOddoc.go:1*/ cu8ZKpmK.aFrRKX.Release()
}

func (hbsnnIi *diffLayer) _YjFhl() Nq4YsH_ {
	lEtpKFfaKmLF := /*line vFVuwG.go:1*/ hbsnnIi.smV5bw0OO()
	return lEtpKFfaKmLF.(Nq4YsH_)
}

func (hbsnnIi *diffLayer) iiFIUyDqR6BJ(evzmhL1 common.Hash) PE_7UyJghxy {
	lEtpKFfaKmLF := /*line wklotaFZS.go:1*/ hbsnnIi.z4rQNTf6c(evzmhL1)
	return lEtpKFfaKmLF.(PE_7UyJghxy)
}

var _ bytes.Buffer
var _ = common.AbsolutePath
