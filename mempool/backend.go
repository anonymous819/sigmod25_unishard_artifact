//line :1
package mempool

import (
	"container/list"
	"sync"

	message "unishard/message"

	"github.com/ethereum/go-ethereum/common"
)

type DUUdwSwZTCm struct {
	ckRwQZTSpMEy *list.List
	aDfpKT       int
	zJHJMtxDGNe  int64
	*EV4KClspBSq
	daeMr27MOLTF *sync.Mutex
}

func SAj4yfpZZaiv(aXWhnba8m int) *DUUdwSwZTCm {
	var cTdZXOVzI_ sync.Mutex
	return &DUUdwSwZTCm{
		ckRwQZTSpMEy:/*line Uj3jgdKUO.go:1*/ list.New(),
		EV4KClspBSq:/*line lrOUHIWla.go:1*/ Rq6zjAUwOnR6(),
		daeMr27MOLTF: &cTdZXOVzI_,
		aDfpKT:       aXWhnba8m,
	}
}

func (nr2GiUlY20Pz *DUUdwSwZTCm) ag6U3RXGgsa4(hDdlup4 *message.Transaction) {
	if hDdlup4 == nil {
		return
	}
	/*line rONdBGpJwq.go:1*/ nr2GiUlY20Pz.daeMr27MOLTF.Lock()
	defer /*line HakYtkl.go:1*/ nr2GiUlY20Pz.daeMr27MOLTF.Unlock()
	if /*line YiwWhtE.go:1*/ nr2GiUlY20Pz.eRan2aoK2() > nr2GiUlY20Pz.aDfpKT {
		return
	}
	nr2GiUlY20Pz.zJHJMtxDGNe++
	/*line MMFYTJ.go:1*/ nr2GiUlY20Pz.ckRwQZTSpMEy.PushBack(hDdlup4)
}

func (nr2GiUlY20Pz *DUUdwSwZTCm) b9rDah_A(hDdlup4 *message.Transaction) {
	if hDdlup4 == nil {
		return
	}
	/*line dTapnKGWS3.go:1*/ nr2GiUlY20Pz.daeMr27MOLTF.Lock()
	defer /*line w64Rh5h.go:1*/ nr2GiUlY20Pz.daeMr27MOLTF.Unlock()
	/*line lUBUuhcAq.go:1*/ nr2GiUlY20Pz.ckRwQZTSpMEy.PushFront(hDdlup4)
}

func (nr2GiUlY20Pz *DUUdwSwZTCm) eRan2aoK2() int {
	return /*line OyCNoByG37.go:1*/ nr2GiUlY20Pz.ckRwQZTSpMEy.Len()
}

func (nr2GiUlY20Pz *DUUdwSwZTCm) dcYEKvGn() *message.Transaction {
	if /*line f6iQsYKhw.go:1*/ nr2GiUlY20Pz.eRan2aoK2() == 0 {
		return nil
	}
	tDkzQYA6 := /*line IUssvP87DE.go:1*/ nr2GiUlY20Pz.ckRwQZTSpMEy.Front()
	fjp0LOqO, o2PUtete4 := tDkzQYA6.Value.(*message.Transaction)
	if !o2PUtete4 {
		return nil
	}
	/*line H5UganigKHC.go:1*/ nr2GiUlY20Pz.ckRwQZTSpMEy.Remove(tDkzQYA6)
	return fjp0LOqO
}

func (nr2GiUlY20Pz *DUUdwSwZTCm) lrQ0VJ(o8rcRtW int) []*message.Transaction {
	var ujBHTJ2u int
	/*line GxTnq28cC2.go:1*/ nr2GiUlY20Pz.daeMr27MOLTF.Lock()
	defer /*line E9rnaVGE06.go:1*/ nr2GiUlY20Pz.daeMr27MOLTF.Unlock()
	ujBHTJ2u = /*line Reh4ux.go:1*/ nr2GiUlY20Pz.eRan2aoK2()
	if ujBHTJ2u >= o8rcRtW {
		ujBHTJ2u = o8rcRtW
	}
	g5y6TBV_W1s := /*line KqRJbVR7IF.go:1*/ make([]*message.Transaction, 0, ujBHTJ2u)
	for iG7hNFS := 0; iG7hNFS < ujBHTJ2u; iG7hNFS++ {
		u9JlOd8Hs1 := /*line hDGSxorMqyhn.go:1*/ nr2GiUlY20Pz.dcYEKvGn()
		if u9JlOd8Hs1 == nil {
			break
		}
		g5y6TBV_W1s = /*line Xmcwk_xh.go:1*/ append(g5y6TBV_W1s, u9JlOd8Hs1)
	}
	return g5y6TBV_W1s
}

func (nr2GiUlY20Pz *DUUdwSwZTCm) vWUhfEYZr(pPKwnhwtDdHa common.Hash) {
	/*line jlaLatXavwXY.go:1*/ nr2GiUlY20Pz.Add(pPKwnhwtDdHa)
}

var _ list.Element
var _ sync.Cond
var _ message.ClientStart
var _ = common.AbsolutePath
