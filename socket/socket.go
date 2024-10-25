//line :1
package socket

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	config "unishard/config"
	log "unishard/log"
	transport "unishard/transport"
	types "unishard/types"
	utils "unishard/utils"
)

type EbznnrSS9z interface {
	GetBlockBuilder() transport.Transport

	GetAddresses() map[types.Shard]map[types.NodeID]string

	SendToBlockBuilder(dRoNa19 interface{})

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

type pQJ3C46N struct {
	eVa0etz9ixz  types.NodeID
	dz6rHsgvk    types.Shard
	q_hjniKP     transport.Transport
	wztTb4BxGqN7 map[types.Shard]map[types.NodeID]string
	jp6Kdhh9sl7  map[types.NodeID]transport.Transport

	_jVr32sd   bool
	eFTRp1DIiO map[types.NodeID]bool
	h0Bop6Z    map[types.NodeID]int
	lpqcRj9    map[types.NodeID]float64

	gSyYFh sync.RWMutex
}

func UxTMLURlR(zZn_LYj93f3D types.NodeID, n5zX44Wh map[types.Shard]map[types.NodeID]string, rWzgYte_vPZ types.Shard) EbznnrSS9z {
	oAbspbLK7T := /*line sNqobtQYc.go:1*/ make(map[types.Shard]map[types.NodeID]string)
	oAbspbLK7T[rWzgYte_vPZ] = /*line Ad4S50aYfF.go:1*/ make(map[types.NodeID]string)
	for hC3G29, uEwAIUF14 := range n5zX44Wh {
		if hC3G29 == rWzgYte_vPZ {
			for zZn_LYj93f3D, aCyiLpB := range uEwAIUF14 {
				t9yLi5rbO := /*line nKUC_j8oDX.go:1*/ strconv.Itoa(3999 + /*line a5ayXtB0WpR.go:1*/ int(rWzgYte_vPZ)*100 + /*line Rknvl2qwB.go:1*/ utils.BD1ZTohx(zZn_LYj93f3D))
				aCyiLpB = aCyiLpB + t9yLi5rbO
				oAbspbLK7T[rWzgYte_vPZ][zZn_LYj93f3D] = aCyiLpB
			}
		}
	}

	oAbspbLK7T[rWzgYte_vPZ][ /*line CBuYc6qCXhM2.go:1*/ utils.NewNodeID(0)] = n5zX44Wh[rWzgYte_vPZ][ /*line zXRQygmAE.go:1*/ utils.NewNodeID(0)] + /*line bpiTXav.go:1*/ strconv.Itoa(6000+ /*line u7NgVa.go:1*/ int(rWzgYte_vPZ))
	pQJ3C46N := &pQJ3C46N{
		eVa0etz9ixz: zZn_LYj93f3D,
		dz6rHsgvk:   rWzgYte_vPZ,
		q_hjniKP:/*line EMUJiYBU2uI.go:1*/ transport.NewTransport(oAbspbLK7T[rWzgYte_vPZ][ /*line uqqA893M0G.go:1*/ utils.NewNodeID(0)]),
		wztTb4BxGqN7: oAbspbLK7T,
		jp6Kdhh9sl7:/*line lZpp9x.go:1*/ make(map[types.NodeID]transport.Transport),
		_jVr32sd: false,
		eFTRp1DIiO:/*line ir6_S84n.go:1*/ make(map[types.NodeID]bool),
		h0Bop6Z:/*line qInSuid1X.go:1*/ make(map[types.NodeID]int),
		lpqcRj9:/*line rxiKXu8IoCT7.go:1*/ make(map[types.NodeID]float64),
	}

	fPGbxN := /*line YFSbGS.go:1*/ utils.Retry(pQJ3C46N.q_hjniKP.Dial, 100 /*line E1DDY_iV.go:1*/, time.Duration(50)*time.Millisecond)
	if fPGbxN != nil {
		/*line nfIbojDj3.go:1*/ panic(fPGbxN)
	}

	pQJ3C46N.jp6Kdhh9sl7[zZn_LYj93f3D] = /*line ayriOt9Nva.go:1*/ transport.NewTransport(pQJ3C46N.wztTb4BxGqN7[rWzgYte_vPZ][zZn_LYj93f3D])
	/*line fssPsr.go:1*/ pQJ3C46N.jp6Kdhh9sl7[zZn_LYj93f3D].Listen()

	return pQJ3C46N
}

func (cT8AFuafk *pQJ3C46N) GetBlockBuilder() transport.Transport {
	return cT8AFuafk.q_hjniKP
}

func (cT8AFuafk *pQJ3C46N) GetAddresses() map[types.Shard]map[types.NodeID]string {
	return cT8AFuafk.wztTb4BxGqN7
}

func (cT8AFuafk *pQJ3C46N) SendToBlockBuilder(dRoNa19 interface{}) {
	awwXMUiK6A := /*line xK9J_8r.go:1*/ utils.NewNodeID(0)
	if cT8AFuafk._jVr32sd {
		return
	}

	if cT8AFuafk.eFTRp1DIiO[awwXMUiK6A] {
		return
	}

	if uAR3fm3wN, tJTHt_i8o := cT8AFuafk.lpqcRj9[awwXMUiK6A]; tJTHt_i8o && uAR3fm3wN > 0 {
		if /*line uVFbGntA.go:1*/ rand.Float64() < uAR3fm3wN {
			return
		}
	}

	if v4pla10, tJTHt_i8o := cT8AFuafk.h0Bop6Z[awwXMUiK6A]; tJTHt_i8o && v4pla10 > 0 {
		gBxWvDSoa2k := /*line je6Hzb.go:1*/ time.NewTimer( /*line EnCdDIynsK.go:1*/ time.Duration(v4pla10) * time.Millisecond)
		go func() {
			<- /*line cjNrRX.go:1*/ gBxWvDSoa2k.C
			/*line rbu_e5.go:1*/ cT8AFuafk.q_hjniKP.Send(dRoNa19)
		}()

		return
	}

	/*line K7IkHReC8gHd.go:1*/
	cT8AFuafk.q_hjniKP.Send(dRoNa19)
}

func (cT8AFuafk *pQJ3C46N) Send(awwXMUiK6A types.NodeID, dRoNa19 interface{}) {

	if cT8AFuafk._jVr32sd {
		return
	}

	if cT8AFuafk.eFTRp1DIiO[awwXMUiK6A] {
		return
	}

	if uAR3fm3wN, tJTHt_i8o := cT8AFuafk.lpqcRj9[awwXMUiK6A]; tJTHt_i8o && uAR3fm3wN > 0 {
		if /*line uhQh__2a.go:1*/ rand.Float64() < uAR3fm3wN {
			return
		}
	}

	/*line BiWf5RL6.go:1*/
	cT8AFuafk.gSyYFh.RLock()
	dcqCa4GP7e1I, enay95JJyPS := cT8AFuafk.jp6Kdhh9sl7[awwXMUiK6A]
	/*line mRAUsS.go:1*/ cT8AFuafk.gSyYFh.RUnlock()
	if !enay95JJyPS {
		/*line JezDRy5lTOH.go:1*/ cT8AFuafk.gSyYFh.RLock()

		uEwAIUF14, tJTHt_i8o := cT8AFuafk.wztTb4BxGqN7[cT8AFuafk.dz6rHsgvk][awwXMUiK6A]
		/*line Cw3q94.go:1*/ cT8AFuafk.gSyYFh.RUnlock()
		if !tJTHt_i8o {
			/*line NMGKww.go:1*/ log.Jp980o6YjM(func() /*line p77PwOybzb.go:1*/ string {
				key := [] /*line p77PwOybzb.go:1*/ byte("\x1dt\"\xc0\\\xa3\x84\xa6\xd3\x16\xd8\xfa>;\x01\xfbc\xaeSL.Æ\xf5ܙ5\xde`\x90\xec\bI\xb1\xa19\xdd[\xec")
				data := [] /*line p77PwOybzb.go:1*/ byte("\x90\xe3\x85+\xc1\x17\xa4\nB{K\x1a\xac\xaau\x1b\xcb\x0fɱN$\xeaYN\xfe\xa8Q\x80\xffR(\xb7 \x05\x9e\xfd\x80_")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return /*line p77PwOybzb.go:1*/ string(data)
			}(), awwXMUiK6A)
			return
		}
		dcqCa4GP7e1I = /*line YlHfP3EB.go:1*/ transport.NewTransport(uEwAIUF14)
		fPGbxN := /*line IahB3D3Ou.go:1*/ utils.Retry(dcqCa4GP7e1I.Dial, 100 /*line JzXWcKW.go:1*/, time.Duration(50)*time.Millisecond)
		if fPGbxN != nil {
			/*line ILSXa5sggMC.go:1*/ panic(fPGbxN)
		}
		/*line _ORAybNc2Qk4.go:1*/ cT8AFuafk.gSyYFh.Lock()
		cT8AFuafk.jp6Kdhh9sl7[awwXMUiK6A] = dcqCa4GP7e1I
		/*line X4gTkGdj7C.go:1*/ cT8AFuafk.gSyYFh.Unlock()
	}

	if v4pla10, tJTHt_i8o := cT8AFuafk.h0Bop6Z[awwXMUiK6A]; tJTHt_i8o && v4pla10 > 0 {
		gBxWvDSoa2k := /*line TEgYTDZ.go:1*/ time.NewTimer( /*line oUybXOEmQ.go:1*/ time.Duration(v4pla10) * time.Millisecond)
		go func() {
			<- /*line WzGsl5lI.go:1*/ gBxWvDSoa2k.C
			/*line oKt9XnGMyG.go:1*/ dcqCa4GP7e1I.Send(dRoNa19)
		}()
		return
	}

	/*line EEVEbMl8SWv.go:1*/
	dcqCa4GP7e1I.Send(dRoNa19)

}

func (cT8AFuafk *pQJ3C46N) Recv() interface{} {
	/*line qAwBTU5l.go:1*/ cT8AFuafk.gSyYFh.RLock()
	dcqCa4GP7e1I := cT8AFuafk.jp6Kdhh9sl7[cT8AFuafk.eVa0etz9ixz]
	/*line _0pH8JkDeT.go:1*/ cT8AFuafk.gSyYFh.RUnlock()
	for {
		dRoNa19 := /*line DhOqRZ.go:1*/ dcqCa4GP7e1I.Recv()
		if !cT8AFuafk._jVr32sd {
			return dRoNa19
		}
	}
}

func (cT8AFuafk *pQJ3C46N) MulticastQuorum(q7JtDot int, dRoNa19 interface{}) {

	ccY6QdReN9M := map[int]struct{}{}
	for dNgGJ90bR := 0; dNgGJ90bR < q7JtDot; dNgGJ90bR++ {
		kexI8k := /*line XLKPpsvjIbyk.go:1*/ rand.Intn( /*line KdMyN5T6haF.go:1*/ len(cT8AFuafk.wztTb4BxGqN7)) + 1
		_, enay95JJyPS := ccY6QdReN9M[kexI8k]
		if enay95JJyPS {
			continue
		}
		/*line bQNzR8Jywxm.go:1*/ cT8AFuafk.Send( /*line Ax_y0Hegwb.go:1*/ utils.NewNodeID(kexI8k), dRoNa19)
		ccY6QdReN9M[kexI8k] = struct{}{}
	}
}

func (cT8AFuafk *pQJ3C46N) Broadcast(dRoNa19 interface{}) {
	ackrbDpOPe := /*line FIxoB1p.go:1*/ time.NewTimer(time.Millisecond * /*line Pzqw6x8nIV_W.go:1*/ time.Duration( /*line L9yX49_SXq.go:1*/ config.GetConfig().InShardLatency))

	<-ackrbDpOPe.C

	for zZn_LYj93f3D := range cT8AFuafk.wztTb4BxGqN7[cT8AFuafk.dz6rHsgvk] {
		if zZn_LYj93f3D == cT8AFuafk.eVa0etz9ixz {
			continue
		}
		/*line iYr6FM.go:1*/ cT8AFuafk.Send(zZn_LYj93f3D, dRoNa19)
	}

}

func (cT8AFuafk *pQJ3C46N) BroadcastToSome(nholaU []types.NodeID, dRoNa19 interface{}) {
	for _, zZn_LYj93f3D := range nholaU {
		if zZn_LYj93f3D == cT8AFuafk.eVa0etz9ixz {
			continue
		}

		if _, yxSkNhArP := cT8AFuafk.wztTb4BxGqN7[cT8AFuafk.dz6rHsgvk][zZn_LYj93f3D]; !yxSkNhArP {
			continue
		}
		/*line vmjMyMs2C.go:1*/ cT8AFuafk.Send(zZn_LYj93f3D, dRoNa19)
	}

}

func (cT8AFuafk *pQJ3C46N) Close() {
	for _, dcqCa4GP7e1I := range cT8AFuafk.jp6Kdhh9sl7 {
		/*line vLGMDXuM.go:1*/ dcqCa4GP7e1I.Close()
	}
}

func (cT8AFuafk *pQJ3C46N) Drop(zZn_LYj93f3D types.NodeID, dcqCa4GP7e1I int) {
	cT8AFuafk.eFTRp1DIiO[zZn_LYj93f3D] = true
	gBxWvDSoa2k := /*line YzxeEvfSV.go:1*/ time.NewTimer( /*line UkXweJmuZ.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
	go func() {
		<- /*line EaICVO.go:1*/ gBxWvDSoa2k.C
		cT8AFuafk.eFTRp1DIiO[zZn_LYj93f3D] = false
	}()
}

func (cT8AFuafk *pQJ3C46N) Slow(zZn_LYj93f3D types.NodeID, v4pla10 int, dcqCa4GP7e1I int) {
	cT8AFuafk.h0Bop6Z[zZn_LYj93f3D] = v4pla10
	gBxWvDSoa2k := /*line Ak_0jYOdZ.go:1*/ time.NewTimer( /*line Us1LrFcOg.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
	go func() {
		<- /*line OmMcuip.go:1*/ gBxWvDSoa2k.C
		cT8AFuafk.h0Bop6Z[zZn_LYj93f3D] = 0
	}()
}

func (cT8AFuafk *pQJ3C46N) Flaky(zZn_LYj93f3D types.NodeID, uAR3fm3wN float64, dcqCa4GP7e1I int) {
	cT8AFuafk.lpqcRj9[zZn_LYj93f3D] = uAR3fm3wN
	gBxWvDSoa2k := /*line FqCfar.go:1*/ time.NewTimer( /*line d7sAncllNOG.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
	go func() {
		<- /*line SIR3VVG.go:1*/ gBxWvDSoa2k.C
		cT8AFuafk.lpqcRj9[zZn_LYj93f3D] = 0
	}()
}

func (cT8AFuafk *pQJ3C46N) Crash(dcqCa4GP7e1I int) {
	cT8AFuafk._jVr32sd = true
	if dcqCa4GP7e1I > 0 {
		gBxWvDSoa2k := /*line E6NK6a0iH.go:1*/ time.NewTimer( /*line JTxt62.go:1*/ time.Duration(dcqCa4GP7e1I) * time.Second)
		go func() {
			<- /*line bJgdZikSIlLk.go:1*/ gBxWvDSoa2k.C
			cT8AFuafk._jVr32sd = false
		}()
	}
}

var _ = rand.ExpFloat64
var _ = strconv.AppendBool
var _ sync.Cond

const _ = time.ANSIC

var _ config.Bconfig
var _ = log.CDebpj
var _ = transport.NewTransport

const _ = types.ABORT

var _ = utils.GffGNE
