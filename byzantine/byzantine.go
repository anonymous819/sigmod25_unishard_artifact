//line :1
package byzantine

import (
	"fmt"
	"time"

	types "unishard/types"
)

type ReportByzantine struct {
	types.Shard
	types.Epoch
	types.View
	DDj8mepzC  types.NodeID
	EaVb63J    bool
	EV7bXXBF0b time.Time
}
type ReplaceByzantine struct {
	types.Shard
	types.Epoch
	types.View
	NDKcLy      types.NodeID
	JEKGSzf     types.NodeID
	ALLbEpAL9Ak bool
	P3nbb_      time.Time
}

type BPJk2Q struct {
	*ReplaceByzantine
}

func MakeReportByzantine(we0FUSjoO types.NodeID, kf0HzAGeunI types.Shard, rKV8Jdp types.Epoch, t8qtpvs types.View) *ReportByzantine {
	yHM6zmHu4T9 := /*line YLqahmW6g.go:1*/ new(ReportByzantine)
	yHM6zmHu4T9.Shard = kf0HzAGeunI
	yHM6zmHu4T9.Epoch = rKV8Jdp
	yHM6zmHu4T9.View = t8qtpvs
	yHM6zmHu4T9.DDj8mepzC = we0FUSjoO
	yHM6zmHu4T9.EaVb63J = true
	yHM6zmHu4T9.EV7bXXBF0b = /*line goYqgQSLwBy.go:1*/ time.Now()

	return yHM6zmHu4T9
}
func JEsMGO(oT17o75 ReportByzantine, rWWl2fcUduOw types.NodeID) *ReplaceByzantine {
	yHM6zmHu4T9 := /*line vbiB84.go:1*/ new(ReplaceByzantine)
	yHM6zmHu4T9.Shard = oT17o75.Shard
	yHM6zmHu4T9.Epoch = oT17o75.Epoch
	yHM6zmHu4T9.View = oT17o75.View
	yHM6zmHu4T9.NDKcLy = oT17o75.DDj8mepzC
	yHM6zmHu4T9.JEKGSzf = rWWl2fcUduOw
	yHM6zmHu4T9.ALLbEpAL9Ak = oT17o75.EaVb63J
	yHM6zmHu4T9.P3nbb_ = /*line wPumwSP10a.go:1*/ time.Now()

	return yHM6zmHu4T9
}

func (g837mIVcfcWa ReportByzantine) String() string {

	return /*line MVo_N6.go:1*/ fmt.Sprintf(func() /*line WpwvdtHvnLoH.go:1*/ string {
		data := /*line WpwvdtHvnLoH.go:1*/ make([]byte, 0, 66)
		i := 14
		decryptKey := 11
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 12:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\x8a\x8f"...,
				)
				i = 2
			case 1:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "75\x87"...,
				)
				i = 23
			case 6:
				i = 13
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, 108)
			case 3:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, 173)
				i = 10
			case 18:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "ZN\x91\xa1"...,
				)
				i = 15
			case 4:
				i = 11
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, 242)
			case 9:
				i = 12
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "x\x9a\x88"...,
				)
			case 25:
				i = 4
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "w;\"("...,
				)
			case 11:
				for y := range data {
					data[y] = data[y] - /*line WpwvdtHvnLoH.go:1*/ byte(decryptKey^y)
				}
				i = 8
			case 14:
				i = 19
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\x8e\xa2"...,
				)
			case 20:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, 170)
				i = 5
			case 5:
				i = 3
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\xadZ}"...,
				)
			case 15:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, 163)
				i = 9
			case 19:
				i = 20
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\xae\xae"...,
				)
			case 16:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "B)"...,
				)
				i = 21
			case 21:
				i = 24
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "/\x81$"...,
				)
			case 24:
				i = 25
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "[ol"...,
				)
			case 22:
				i = 26
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "A}H\x92"...,
				)
			case 17:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\x9a\xa0\x98L"...,
				)
				i = 18
			case 0:
				i = 1
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "yP"...,
				)
			case 7:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "||qw"...,
				)
				i = 16
			case 13:
				i = 0
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\x82|\x86"...,
				)
			case 26:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "z>N8"...,
				)
				i = 6
			case 2:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\x8bZ"...,
				)
				i = 22
			case 10:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "\xaf\x97\xa5\xa4"...,
				)
				i = 17
			case 23:
				data = /*line WpwvdtHvnLoH.go:1*/ append(data, "2X"...,
				)
				i = 7
			}
		}
		return /*line WpwvdtHvnLoH.go:1*/ string(data)
	}(), g837mIVcfcWa.DDj8mepzC, g837mIVcfcWa.Shard, g837mIVcfcWa.Epoch, g837mIVcfcWa.View)
}

func (iSAVRj4k ReplaceByzantine) String() string {

	return /*line qhVELQ9oz.go:1*/ fmt.Sprintf(func() /*line e6FyRgKC.go:1*/ string {
		data := /*line e6FyRgKC.go:1*/ make([]byte, 0, 89)
		i := 0
		decryptKey := 123
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 16:
				i = 4
				data = /*line e6FyRgKC.go:1*/ append(data, "cu\xe0\xb7"...,
				)
			case 0:
				i = 10
				data = /*line e6FyRgKC.go:1*/ append(data, "\x8e\xa2"...,
				)
			case 3:
				i = 29
				data = /*line e6FyRgKC.go:1*/ append(data, "7\x89i-"...,
				)
			case 23:
				i = 2
				for y := range data {
					data[y] = data[y] - /*line e6FyRgKC.go:1*/ byte(decryptKey^y)
				}
			case 20:
				i = 16
				data = /*line e6FyRgKC.go:1*/ append(data, 105)
			case 32:
				data = /*line e6FyRgKC.go:1*/ append(data, "\x81=\x92\x8e"...,
				)
				i = 28
			case 5:
				i = 34
				data = /*line e6FyRgKC.go:1*/ append(data, "\x87\x7f"...,
				)
			case 19:
				i = 32
				data = /*line e6FyRgKC.go:1*/ append(data, 135)
			case 22:
				i = 13
				data = /*line e6FyRgKC.go:1*/ append(data, "\xbf\xeb\xe3\xd8"...,
				)
			case 27:
				i = 18
				data = /*line e6FyRgKC.go:1*/ append(data, 137)
			case 1:
				i = 27
				data = /*line e6FyRgKC.go:1*/ append(data, "\xd3\xe6\xa2"...,
				)
			case 26:
				i = 19
				data = /*line e6FyRgKC.go:1*/ append(data, "\xa0\x9ao\x91"...,
				)
			case 15:
				i = 1
				data = /*line e6FyRgKC.go:1*/ append(data, "\x93\xc2\xd6"...,
				)
			case 30:
				data = /*line e6FyRgKC.go:1*/ append(data, "\x90\x8a"...,
				)
				i = 35
			case 21:
				data = /*line e6FyRgKC.go:1*/ append(data, "\xee\x99"...,
				)
				i = 22
			case 12:
				i = 7
				data = /*line e6FyRgKC.go:1*/ append(data, "\xa5\xa4\x9a\xa0"...,
				)
			case 4:
				i = 21
				data = /*line e6FyRgKC.go:1*/ append(data, "\x9e\xa4"...,
				)
			case 33:
				data = /*line e6FyRgKC.go:1*/ append(data, "|0l"...,
				)
				i = 3
			case 14:
				i = 30
				data = /*line e6FyRgKC.go:1*/ append(data, "LZN\x92"...,
				)
			case 29:
				data = /*line e6FyRgKC.go:1*/ append(data, ";M("...,
				)
				i = 31
			case 18:
				i = 23
				data = /*line e6FyRgKC.go:1*/ append(data, "\x8f\xe1"...,
				)
			case 13:
				i = 17
				data = /*line e6FyRgKC.go:1*/ append(data, "ޱ\x90"...,
				)
			case 11:
				data = /*line e6FyRgKC.go:1*/ append(data, "\xae\xaa\xadZ"...,
				)
				i = 9
			case 24:
				i = 26
				data = /*line e6FyRgKC.go:1*/ append(data, "\x92\x89E\x88"...,
				)
			case 7:
				data = /*line e6FyRgKC.go:1*/ append(data, 152)
				i = 14
			case 25:
				i = 15
				data = /*line e6FyRgKC.go:1*/ append(data, 232)
			case 9:
				data = /*line e6FyRgKC.go:1*/ append(data, "}\xad\xaf\x97"...,
				)
				i = 12
			case 31:
				data = /*line e6FyRgKC.go:1*/ append(data, "d/\x81a"...,
				)
				i = 8
			case 28:
				i = 5
				data = /*line e6FyRgKC.go:1*/ append(data, 56)
			case 10:
				i = 11
				data = /*line e6FyRgKC.go:1*/ append(data, 174)
			case 34:
				i = 33
				data = /*line e6FyRgKC.go:1*/ append(data, "\x92b\x84z"...,
				)
			case 17:
				data = /*line e6FyRgKC.go:1*/ append(data, 150)
				i = 25
			case 8:
				data = /*line e6FyRgKC.go:1*/ append(data, 37)
				i = 6
			case 6:
				data = /*line e6FyRgKC.go:1*/ append(data, "5'S"...,
				)
				i = 20
			case 35:
				i = 24
				data = /*line e6FyRgKC.go:1*/ append(data, 152)
			}
		}
		return /*line e6FyRgKC.go:1*/ string(data)
	}(), iSAVRj4k.NDKcLy, iSAVRj4k.JEKGSzf, iSAVRj4k.Shard, iSAVRj4k.Epoch, iSAVRj4k.View)
}

var _ = fmt.Append

const _ = time.ANSIC
const _ = types.ABORT
