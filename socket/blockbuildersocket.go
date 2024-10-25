//line :1
package socket

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	log "unishard/log"
	transport "unishard/transport"
	types "unishard/types"
	utils "unishard/utils"
)

type VN3GUe interface {
	Send(awwXMUiK6A types.NodeID, dRoNa19 interface{})

	MulticastQuorum(q7JtDot int, dRoNa19 interface{})

	Broadcast(dRoNa19 interface{})

	BroadcastToSome(nholaU []types.NodeID, dRoNa19 interface{})

	Recv() interface{}

	Close()

	Drop(zZn_LYj93f3D types.NodeID, dcqCa4GP7e1I int)
	Slow(zZn_LYj93f3D types.NodeID, aex_Fp_19Q int, dcqCa4GP7e1I int)
	Flaky(zZn_LYj93f3D types.NodeID, uAR3fm3wN float64, dcqCa4GP7e1I int)
	Crash(dcqCa4GP7e1I int)
}

type mxbaOtQRo struct {
	aRxCNe     string
	gJscERFw2l types.Shard
	lzxrOtg4E  transport.Transport
	aV7XBri    map[types.NodeID]string
	bUgt_faoa0 map[types.NodeID]transport.Transport

	biqRFZ8    bool
	f8zewAq87M map[types.NodeID]bool
	ddoQGJa    map[types.NodeID]int
	eOU3guNe2  map[types.NodeID]float64

	m6PLCmgKK sync.RWMutex
}

func P2LcT9Ia(ncqr73aTfs string, rWzgYte_vPZ types.Shard, n5zX44Wh map[types.Shard]map[types.NodeID]string) VN3GUe {
	ybTnuFJA4y1m := /*line HMVeHx.go:1*/ new(mxbaOtQRo)
	ybTnuFJA4y1m.aRxCNe = ncqr73aTfs
	ybTnuFJA4y1m.gJscERFw2l = rWzgYte_vPZ
	ybTnuFJA4y1m.lzxrOtg4E = /*line ysILvATxzvj2.go:1*/ transport.NewTransport(ncqr73aTfs)
	oAbspbLK7T := /*line BZBJfX_Ic.go:1*/ make(map[types.Shard]map[types.NodeID]string)
	oAbspbLK7T[rWzgYte_vPZ] = /*line Gt1XqeE_bQR.go:1*/ make(map[types.NodeID]string)
	for _, uEwAIUF14 := range n5zX44Wh {
		for zZn_LYj93f3D, aCyiLpB := range uEwAIUF14 {
			t9yLi5rbO := /*line JKvoksEM.go:1*/ strconv.Itoa(3999 + /*line _vulOtI6.go:1*/ int(rWzgYte_vPZ)*100 + /*line fN9vIvGmu.go:1*/ utils.BD1ZTohx(zZn_LYj93f3D))
			aCyiLpB += t9yLi5rbO
			oAbspbLK7T[rWzgYte_vPZ][zZn_LYj93f3D] = aCyiLpB
		}
	}
	ybTnuFJA4y1m.aV7XBri = oAbspbLK7T[rWzgYte_vPZ]
	ybTnuFJA4y1m.bUgt_faoa0 = /*line Kxdj9a8UGx.go:1*/ make(map[types.NodeID]transport.Transport)
	ybTnuFJA4y1m.biqRFZ8 = false
	ybTnuFJA4y1m.f8zewAq87M = /*line Foo3uYN0HhF.go:1*/ make(map[types.NodeID]bool)
	ybTnuFJA4y1m.ddoQGJa = /*line DQDGEO.go:1*/ make(map[types.NodeID]int)
	ybTnuFJA4y1m.eOU3guNe2 = /*line pUtpBtZtL.go:1*/ make(map[types.NodeID]float64)

	/*line qWt_3TNxMR.go:1*/
	ybTnuFJA4y1m.lzxrOtg4E.Listen()

	return ybTnuFJA4y1m
}

func (cT8AFuafk *mxbaOtQRo) Send(awwXMUiK6A types.NodeID, dRoNa19 interface{}) {

	if cT8AFuafk.biqRFZ8 {
		return
	}

	if cT8AFuafk.f8zewAq87M[awwXMUiK6A] {
		return
	}

	if uAR3fm3wN, tJTHt_i8o := cT8AFuafk.eOU3guNe2[awwXMUiK6A]; tJTHt_i8o && uAR3fm3wN > 0 {
		if /*line Za1rqjSNyqU.go:1*/ rand.Float64() < uAR3fm3wN {
			return
		}
	}

	/*line by2amparu4.go:1*/
	cT8AFuafk.m6PLCmgKK.RLock()
	dcqCa4GP7e1I, enay95JJyPS := cT8AFuafk.bUgt_faoa0[awwXMUiK6A]
	/*line TPq6Svpp.go:1*/ cT8AFuafk.m6PLCmgKK.RUnlock()
	if !enay95JJyPS {
		/*line H0xLfA.go:1*/ cT8AFuafk.m6PLCmgKK.RLock()
		uEwAIUF14, tJTHt_i8o := cT8AFuafk.aV7XBri[awwXMUiK6A]
		/*line D4aY2Kk.go:1*/ cT8AFuafk.m6PLCmgKK.RUnlock()
		if !tJTHt_i8o {
			/*line XWe7ih_acII.go:1*/ log.Jp980o6YjM(func() /*line nvAKuLENApf.go:1*/ string {
				data := /*line nvAKuLENApf.go:1*/ make([]byte, 0, 40)
				i := 11
				decryptKey := 95
				for counter := 0; i != 2; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 8:
						data = /*line nvAKuLENApf.go:1*/ append(data, "\xce\x17"...,
						)
						i = 12
					case 7:
						data = /*line nvAKuLENApf.go:1*/ append(data, 209)
						i = 4
					case 4:
						data = /*line nvAKuLENApf.go:1*/ append(data, "\x11\x17\x16"...,
						)
						i = 0
					case 0:
						data = /*line nvAKuLENApf.go:1*/ append(data, 47)
						i = 9
					case 15:
						i = 1
						data = /*line nvAKuLENApf.go:1*/ append(data, "\x11\x06\x14\xc3"...,
						)
					case 9:
						data = /*line nvAKuLENApf.go:1*/ append(data, "!21\xd9"...,
						)
						i = 5
					case 17:
						i = 13
						data = /*line nvAKuLENApf.go:1*/ append(data, "\xf3\xf3\xeb"...,
						)
					case 12:
						data = /*line nvAKuLENApf.go:1*/ append(data, "\x17\x1f\xca\x1d"...,
						)
						i = 6
					case 3:
						data = /*line nvAKuLENApf.go:1*/ append(data, "\x13\n"...,
						)
						i = 15
					case 1:
						i = 8
						data = /*line nvAKuLENApf.go:1*/ append(data, "\x06\x1c\x11\""...,
						)
					case 16:
						i = 2
						for y := range data {
							data[y] = data[y] - /*line nvAKuLENApf.go:1*/ byte(decryptKey^y)
						}
					case 14:
						i = 7
						data = /*line nvAKuLENApf.go:1*/ append(data, "-\x1b"...,
						)
					case 13:
						i = 10
						data = /*line nvAKuLENApf.go:1*/ append(data, "롥"...,
						)
					case 10:
						data = /*line nvAKuLENApf.go:1*/ append(data, 246)
						i = 16
					case 11:
						data = /*line nvAKuLENApf.go:1*/ append(data, 24)
						i = 3
					case 5:
						i = 17
						data = /*line nvAKuLENApf.go:1*/ append(data, "'!\xda"...,
						)
					case 6:
						data = /*line nvAKuLENApf.go:1*/ append(data, 21)
						i = 14
					}
				}
				return /*line nvAKuLENApf.go:1*/ string(data)
			}(), awwXMUiK6A)

			return
		}
		dcqCa4GP7e1I = /*line pGcoQI.go:1*/ transport.NewTransport(uEwAIUF14)
		fPGbxN := /*line r6m9SmT.go:1*/ utils.Retry(dcqCa4GP7e1I.Dial, 100 /*line j0iidxgeg.go:1*/, time.Duration(50)*time.Millisecond)
		if fPGbxN != nil {
			/*line SC66ej.go:1*/ panic(fPGbxN)
		}
		/*line YTvFiXa.go:1*/ cT8AFuafk.m6PLCmgKK.Lock()
		cT8AFuafk.bUgt_faoa0[awwXMUiK6A] = dcqCa4GP7e1I
		/*line cCqLrNa7gTjw.go:1*/ cT8AFuafk.m6PLCmgKK.Unlock()
	}

	if v4pla10, tJTHt_i8o := cT8AFuafk.ddoQGJa[awwXMUiK6A]; tJTHt_i8o && v4pla10 > 0 {
		gBxWvDSoa2k := /*line pz7kA11kLGK.go:1*/ time.NewTimer( /*line EOA_OaP86C.go:1*/ time.Duration(v4pla10) * time.Millisecond)
		go func() {
			<- /*line GOLQiWw7.go:1*/ gBxWvDSoa2k.C
			/*line Fcd0dkdihD_.go:1*/ dcqCa4GP7e1I.Send(dRoNa19)
		}()

		return
	}

	/*line AvJfWanXpHIq.go:1*/
	dcqCa4GP7e1I.Send(dRoNa19)

}

func (cT8AFuafk *mxbaOtQRo) Recv() interface{} {
	/*line s_EriOXZhLd.go:1*/ cT8AFuafk.m6PLCmgKK.RLock()
	dcqCa4GP7e1I := cT8AFuafk.lzxrOtg4E
	/*line XrHnMGJgflW.go:1*/ cT8AFuafk.m6PLCmgKK.RUnlock()
	for {
		dRoNa19 := /*line NDlKO08AaDU.go:1*/ dcqCa4GP7e1I.Recv()
		if !cT8AFuafk.biqRFZ8 {

			return dRoNa19
		}
	}
}

func (cT8AFuafk *mxbaOtQRo) MulticastQuorum(q7JtDot int, dRoNa19 interface{}) {

	ccY6QdReN9M := map[int]struct{}{}
	for dNgGJ90bR := 0; dNgGJ90bR < q7JtDot; dNgGJ90bR++ {
		kexI8k := /*line CvZqnXUx_.go:1*/ rand.Intn( /*line _PT6AW.go:1*/ len(cT8AFuafk.aV7XBri)) + 1
		_, enay95JJyPS := ccY6QdReN9M[kexI8k]
		if enay95JJyPS {
			continue
		}
		/*line a6QGa4b.go:1*/ cT8AFuafk.Send( /*line lJ47iwPw1CE1.go:1*/ utils.NewNodeID(kexI8k), dRoNa19)
		ccY6QdReN9M[kexI8k] = struct{}{}
	}
}

func (cT8AFuafk *mxbaOtQRo) Broadcast(dRoNa19 interface{}) {

	for zZn_LYj93f3D := range cT8AFuafk.aV7XBri {
		/*line DaERg5YjpX.go:1*/ cT8AFuafk.Send(zZn_LYj93f3D, dRoNa19)
	}

}

func (cT8AFuafk *mxbaOtQRo) BroadcastToSome(nholaU []types.NodeID, dRoNa19 interface{}) {

	for _, zZn_LYj93f3D := range nholaU {
		if _, yxSkNhArP := cT8AFuafk.aV7XBri[zZn_LYj93f3D]; !yxSkNhArP {
			continue
		}
		/*line uIsamr497.go:1*/ cT8AFuafk.Send(zZn_LYj93f3D, dRoNa19)
	}

}

func (cT8AFuafk *mxbaOtQRo) Close() {
	for _, dcqCa4GP7e1I := range cT8AFuafk.bUgt_faoa0 {
		/*line BVNty6FtzcM.go:1*/ dcqCa4GP7e1I.Close()
	}
}

func (cT8AFuafk *mxbaOtQRo) Drop(zZn_LYj93f3D types.NodeID, dcqCa4GP7e1I int) {
	cT8AFuafk.f8zewAq87M[zZn_LYj93f3D] = true
	gBxWvDSoa2k := /*line ZhSpN07d3n3A.go:1*/ time.NewTimer( /*line mFeDrxmeFdS.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
	go func() {
		<- /*line nPvub_.go:1*/ gBxWvDSoa2k.C
		cT8AFuafk.f8zewAq87M[zZn_LYj93f3D] = false
	}()
}

func (cT8AFuafk *mxbaOtQRo) Slow(zZn_LYj93f3D types.NodeID, v4pla10 int, dcqCa4GP7e1I int) {
	cT8AFuafk.ddoQGJa[zZn_LYj93f3D] = v4pla10
	gBxWvDSoa2k := /*line C7gjiUE.go:1*/ time.NewTimer( /*line aA_BoMWa_AS7.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
	go func() {
		<- /*line VyCJeVvUUDu.go:1*/ gBxWvDSoa2k.C
		cT8AFuafk.ddoQGJa[zZn_LYj93f3D] = 0
	}()
}

func (cT8AFuafk *mxbaOtQRo) Flaky(zZn_LYj93f3D types.NodeID, uAR3fm3wN float64, dcqCa4GP7e1I int) {
	cT8AFuafk.eOU3guNe2[zZn_LYj93f3D] = uAR3fm3wN
	gBxWvDSoa2k := /*line BmjnxGfUU.go:1*/ time.NewTimer( /*line Prlfbf.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
	go func() {
		<- /*line LHMqVo5.go:1*/ gBxWvDSoa2k.C
		cT8AFuafk.eOU3guNe2[zZn_LYj93f3D] = 0
	}()
}

func (cT8AFuafk *mxbaOtQRo) Crash(dcqCa4GP7e1I int) {
	cT8AFuafk.biqRFZ8 = true
	if dcqCa4GP7e1I > 0 {
		gBxWvDSoa2k := /*line ezaAfyfb.go:1*/ time.NewTimer( /*line HBlwIxVEU8.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
		go func() {
			<- /*line Jo_b5YeM.go:1*/ gBxWvDSoa2k.C
			cT8AFuafk.biqRFZ8 = false
		}()
	}
}

var _ = rand.ExpFloat64
var _ = strconv.AppendBool
var _ sync.Cond

const _ = time.ANSIC

var _ = log.CDebpj
var _ = transport.NewTransport

const _ = types.ABORT

var _ = utils.GffGNE
