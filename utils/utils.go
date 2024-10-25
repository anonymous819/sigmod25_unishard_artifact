//line :1
package utils

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"reflect"
	"strconv"
	"time"

	log "unishard/log"
	types "unishard/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewNodeID(adzk4Rz int) types.NodeID {
	if adzk4Rz < 0 {
		adzk4Rz = -adzk4Rz
	}

	return /*line F70Fbkgo4kp.go:1*/ types.NodeID( /*line F2xTV5.go:1*/ strconv.Itoa(adzk4Rz))
}

func BD1ZTohx(vxOfoe types.NodeID) int {
	bIGav78KPq := /*line OrLedUO3CRP.go:1*/ string(vxOfoe)
	adzk4Rz, m6yANu := /*line B_x6LC.go:1*/ strconv.ParseUint(bIGav78KPq, 10, 64)
	if m6yANu != nil {
		/*line H8_F6H57cI.go:1*/ log.Jp980o6YjM(func() /*line HhRzJ_KP4U.go:1*/ string {
			fullData := [] /*line HhRzJ_KP4U.go:1*/ byte("\xbc*0l\xda1\xf7)uW\x8a\xe3\x15s\xd7\b1\xaa7\xa5\xb8\xd6h\x84Ahc8,+N\x94\xbf\x1e\xdf\x11H\xa0\x91\x95\xdat3\xber\xff\xf0\x10Y\xcf\xcaK\xe5ǒ\x05\xb1;\x9c\x86qy,p\xe5\xa2")
			idxKey := [] /*line HhRzJ_KP4U.go:1*/ byte("\xb1\x9b\xa5\xb3\x90K`v\xd22s4\xd3\xe9\x1a\x83")
			data := /*line HhRzJ_KP4U.go:1*/ make([]byte, 0, 34)
			data = /*line HhRzJ_KP4U.go:1*/ append(data, fullData[250^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[8])]-fullData[205^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[8])], fullData[73^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[6])]^fullData[108^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[6])], fullData[92^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[5])]+fullData[127^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[5])], fullData[6^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])]+fullData[117^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])], fullData[189^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])]-fullData[182^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])], fullData[140^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])]^fullData[128^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])], fullData[19^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])]-fullData[60^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])], fullData[150^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]-fullData[182^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[124^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])]^fullData[54^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])], fullData[82^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])]^fullData[111^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])], fullData[139^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[4])]+fullData[141^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[4])], fullData[185^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]-fullData[164^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[135^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]-fullData[176^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[184^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])]+fullData[173^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])], fullData[83^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[10])]^fullData[119^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[10])], fullData[184^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[3])]^fullData[149^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[3])], fullData[118^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])]+fullData[98^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])], fullData[184^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]^fullData[158^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[172^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]^fullData[168^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[170^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]^fullData[190^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[133^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])]^fullData[154^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[1])], fullData[150^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[4])]^fullData[166^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[4])], fullData[19^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[14])]-fullData[8^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[14])], fullData[252^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[13])]-fullData[209^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[13])], fullData[76^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])]-fullData[113^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])], fullData[239^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[8])]^fullData[226^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[8])], fullData[144^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])]-fullData[147^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[15])], fullData[144^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[0])]^fullData[141^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[0])], fullData[54^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])]-fullData[27^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])], fullData[90^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[5])]-fullData[83^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[5])], fullData[30^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])]+fullData[13^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[11])], fullData[120^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])]-fullData[108^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[7])], fullData[30^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[9])]-fullData[36^ /*line HhRzJ_KP4U.go:1*/ int(idxKey[9])])
			return /*line HhRzJ_KP4U.go:1*/ string(data)
		}(), bIGav78KPq)

		return 0
	}

	return /*line Fus_cSyDPO.go:1*/ int(adzk4Rz)
}

func UB_qaFnx(lxGxpmDNX types.NW4p1ya) int { return /*line kFIAXOSbkaS7.go:1*/ len(lxGxpmDNX) }

func SlotToKey(aVVxfvC uint64) common.Hash {
	q3QLHDGHc := /*line EXPdXi4ATT.go:1*/ common.BigToHash( /*line AaPTxjy1z.go:1*/ big.NewInt( /*line _z5GSRXUQ9E.go:1*/ int64(aVVxfvC)))
	return q3QLHDGHc
}

func CmTWaXnCjAw[M0GKmt2Jc any](cTPAWVS int, t0xIanBB []M0GKmt2Jc) []M0GKmt2Jc {
	veWzaG2UaYN := []M0GKmt2Jc{}
	for qJ7E9njoXF, yhY5ZDPTzqw := range t0xIanBB {
		if cTPAWVS == qJ7E9njoXF {
			continue
		}
		veWzaG2UaYN = /*line ZcadgF.go:1*/ append(veWzaG2UaYN, yhY5ZDPTzqw)
	}
	return veWzaG2UaYN
}

func CalculateShardToSend(iBK9LCj7Jd9 []common.Address, aDWG5b int) []types.Shard {
	btbMejxPbUDY := []types.Shard{}
	for _, bAHqpNRbBg := range iBK9LCj7Jd9 {
		lxGxpmDNX := /*line WJEEwW.go:1*/ new(big.Int)
		/*line QcogRTVmjZ.go:1*/ lxGxpmDNX.SetString( /*line F5F5Sylqurc9.go:1*/ bAHqpNRbBg.Hex()[2:], 16)
		d6Jwai5dE := /*line xwB0zW9jfGD.go:1*/ types.Shard( /*line E1FbclAr.go:1*/ lxGxpmDNX.Mod(lxGxpmDNX /*line bB5BWnUJOj.go:1*/, big.NewInt( /*line gPjceaz6C.go:1*/ int64(aDWG5b))).Int64() + 1)
		if ! /*line FD7VyxqTHu0u.go:1*/ HXBOaN(btbMejxPbUDY, d6Jwai5dE) {
			btbMejxPbUDY = /*line lYDcR4D.go:1*/ append(btbMejxPbUDY, d6Jwai5dE)
		}
	}

	return btbMejxPbUDY
}

func GffGNE(bAHqpNRbBg common.Address, aVVxfvC uint64) []byte {

	fetEUQCl9f := /*line sfEXge6.go:1*/ common.LeftPadBytes( /*line NyDNlpX.go:1*/ bAHqpNRbBg.Bytes(), 32)

	ruoulMJu9 := /*line we3hM9sFU.go:1*/ make([]byte, 32)
	/*line gXpYZi2Gk.go:1*/ binary.BigEndian.PutUint64(ruoulMJu9[24:], aVVxfvC)

	return /*line BvX8Iiy62Ys1.go:1*/ crypto.Keccak256(fetEUQCl9f, ruoulMJu9)
}

func BwwcNZ(t0xIanBB []int, qUwwAo int) bool {
	for _, vUbTDI := range t0xIanBB {
		if vUbTDI == qUwwAo {
			return true
		}
	}
	return false
}

func QyGZ203ItEr(heg5gm5G int, mfi40IvBEJ int) []int {
	var jvLB0l []int
	for vxOfoe := 0; vxOfoe < mfi40IvBEJ; vxOfoe++ {
		var tWtZQNdsY_I int
		q94Lrn := true
		for q94Lrn {
			bIGav78KPq := /*line Dv6phihl.go:1*/ mathRand.NewSource( /*line E6po8N.go:1*/ time.Now().UnixNano())
			p2JM5mDMJf4 := /*line gYpwSNEmU.go:1*/ mathRand.New(bIGav78KPq)
			tWtZQNdsY_I = /*line ZIOK4QmhW1p.go:1*/ p2JM5mDMJf4.Intn(heg5gm5G)
			q94Lrn = /*line QTLetZ.go:1*/ BwwcNZ(jvLB0l, tWtZQNdsY_I)
		}
		jvLB0l = /*line Oi6Eg1cT0Va.go:1*/ append(jvLB0l, tWtZQNdsY_I)
	}
	return jvLB0l
}

func Aqz9fJqvVTX(lxGxpmDNX, yVFl_XX int) int {
	if lxGxpmDNX < yVFl_XX {
		return yVFl_XX
	}
	return lxGxpmDNX
}

func I1_I6HrOF(iQL7ovCkURp ...int) int {
	w6O_28ehj := iQL7ovCkURp[0]
	for _, vxOfoe := range iQL7ovCkURp {
		if w6O_28ehj < vxOfoe {
			w6O_28ehj = vxOfoe
		}
	}
	return w6O_28ehj
}

func Retry(mfi40IvBEJ func() error, tWYI725 int, o8G2Bn time.Duration) error {
	var m6yANu error
	for vxOfoe := 0; ; vxOfoe++ {
		m6yANu = /*line VaMr7zAszA.go:1*/ mfi40IvBEJ()
		if m6yANu == nil {
			return nil
		}

		if vxOfoe >= tWYI725-1 {
			break
		}

		/*line iiaaLTxLYPz.go:1*/
		time.Sleep(o8G2Bn * /*line Hpp0FiAac.go:1*/ time.Duration(vxOfoe+1))
	}
	return /*line sRB9jrfhkJ.go:1*/ fmt.Errorf(func() /*line B9CKJDB0Lfo.go:1*/ string {
		data := /*line B9CKJDB0Lfo.go:1*/ make([]byte, 0, 34)
		i := 11
		decryptKey := 75
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 9:
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x9e\xaf"...,
				)
				i = 12
			case 13:
				i = 2
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\xa2Q"...,
				)
			case 7:
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x92\x95\x91"...,
				)
				i = 3
			case 5:
				i = 0
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, 166)
			case 8:
				i = 4
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\xaa\x9c"...,
				)
			case 1:
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\xa3N\x86"...,
				)
				i = 7
			case 0:
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x9a\xa1\xa7"...,
				)
				i = 8
			case 2:
				i = 5
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x91\xa7"...,
				)
			case 12:
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x9f\xaf\\d"...,
				)
				i = 13
			case 11:
				i = 9
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, 154)
			case 10:
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, 144)
				i = 15
			case 4:
				i = 14
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "TK\x96"...,
				)
			case 15:
				i = 6
				for y := range data {
					data[y] = data[y] - /*line B9CKJDB0Lfo.go:1*/ byte(decryptKey^y)
				}
			case 3:
				i = 10
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x97^GK"...,
				)
			case 14:
				i = 1
				data = /*line B9CKJDB0Lfo.go:1*/ append(data, "\x8e\x9f"...,
				)
			}
		}
		return /*line B9CKJDB0Lfo.go:1*/ string(data)
	}(), tWYI725, m6yANu)
}

func XmQq2UARY(mfi40IvBEJ func(), zpGABV6oWJg time.Duration) chan bool {
	nLb3nNs := /*line JdZaz1ho.go:1*/ make(chan bool)

	go func() {
		for {
			/*line Crd8acmwJJ.go:1*/ mfi40IvBEJ()
			select {
			case <- /*line c8AKY3E70.go:1*/ time.After(zpGABV6oWJg):
			case <-nLb3nNs:
				return
			}
		}
	}()

	return nLb3nNs
}

func Dycs2kb_(mRFr5r5w interface{}) interface{} {
	d5L6Lbp5G := /*line Urva3qSY.go:1*/ reflect.ValueOf(mRFr5r5w).MapKeys()

	return /*line k4P16m7.go:1*/ d5L6Lbp5G[ /*line RXWb2tM.go:1*/ mathRand.Intn( /*line GM7i7N6sM3.go:1*/ len(d5L6Lbp5G))].Interface()
}

func PVU2LPyJif() common.Hash {
	var nDep1BZ common.Hash
	_, _ = /*line rw_pPNnxuGvD.go:1*/ rand.Read(nDep1BZ[:])
	return nDep1BZ
}

func JujCSyOLS3[M0GKmt2Jc, VlSop2s any](ft_PSP []M0GKmt2Jc, gRdF4pbBb func(M0GKmt2Jc) VlSop2s) []VlSop2s {
	vMaZ4GTxj := /*line PuMxSyD.go:1*/ make([]VlSop2s /*line F5hoIN62yo.go:1*/, len(ft_PSP))
	for vxOfoe, lOTvaOCdcn_ := range ft_PSP {
		vMaZ4GTxj[vxOfoe] = /*line byX2b06Za.go:1*/ gRdF4pbBb(lOTvaOCdcn_)
	}
	return vMaZ4GTxj
}

func OUWLEOgmMfR[M0GKmt2Jc any](juQweJ map[types.Epoch]map[types.View]map[types.BlockHeight]M0GKmt2Jc, xskvTyH types.Epoch, nzCiXcatSM types.View, zC5YMA types.BlockHeight, iQL7ovCkURp M0GKmt2Jc) {
	if _, mWU5YPY := juQweJ[xskvTyH]; !mWU5YPY {
		juQweJ[xskvTyH] = /*line S42e8KRPn.go:1*/ make(map[types.View]map[types.BlockHeight]M0GKmt2Jc)
		juQweJ[xskvTyH][nzCiXcatSM] = /*line P56gC99NVK.go:1*/ make(map[types.BlockHeight]M0GKmt2Jc)
	} else if _, mWU5YPY := juQweJ[xskvTyH][nzCiXcatSM]; !mWU5YPY {
		juQweJ[xskvTyH][nzCiXcatSM] = /*line hZDle27t5T.go:1*/ make(map[types.BlockHeight]M0GKmt2Jc)
	}

	juQweJ[xskvTyH][nzCiXcatSM][zC5YMA] = iQL7ovCkURp
}

func HXBOaN(bIGav78KPq []types.Shard, bNCrcpziy types.Shard) bool {
	for _, vUbTDI := range bIGav78KPq {
		if vUbTDI == bNCrcpziy {
			return true
		}
	}
	return false
}

func R59H6YyJ3h(qIRHRqgPUdBA []common.Hash, bNCrcpziy common.Hash) bool {
	for _, vUbTDI := range qIRHRqgPUdBA {
		if vUbTDI == bNCrcpziy {
			return true
		}
	}
	return false
}

var _ = rand.Int
var _ = binary.Append
var _ = fmt.Append

const _ = big.Above

var _ = mathRand.ExpFloat64
var _ = reflect.Append
var _ = strconv.AppendBool

const _ = time.ANSIC

var _ = log.CDebpj

const _ = types.ABORT

var _ = common.AbsolutePath
var _ = crypto.CompressPubkey
