//line :1
package node

import (
	"net/http"
	"reflect"
	"sync"

	config "unishard/config"
	log "unishard/log"
	socket "unishard/socket"
	types "unishard/types"
)

type YxdKRgsOCr interface {
	socket.VN3GUe

	GetIP() string
	GetShard() types.Shard
	Run()
	Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{})
}

type gatewaynode struct {
	ip    string
	shard types.Shard

	socket.VN3GUe
	server *http.Server

	handles     map[string]reflect.Value
	MessageChan chan interface{}

	sync.RWMutex
}

func I2vD85cy(lq54woHa3C string, js85nByJG types.Shard) YxdKRgsOCr {
	mM6kT_d := /*line hIt5iU.go:1*/ new(gatewaynode)
	mM6kT_d.ip = lq54woHa3C
	mM6kT_d.shard = js85nByJG
	mM6kT_d.VN3GUe = /*line IuZut1.go:1*/ socket.P2LcT9Ia(lq54woHa3C, js85nByJG, config.Configuration.Addrs)
	mM6kT_d.handles = /*line tpdaeKJ.go:1*/ make(map[string]reflect.Value)
	mM6kT_d.MessageChan = /*line NOA450uS.go:1*/ make(chan interface{}, config.Configuration.ChanBufferSize)
	return mM6kT_d
}

func (mM6kT_d *gatewaynode) GetIP() string {
	return mM6kT_d.ip
}

func (mM6kT_d *gatewaynode) GetShard() types.Shard {
	return mM6kT_d.shard
}

func (mM6kT_d *gatewaynode) Register(cYu43eeVpCaK interface{}, cy0xXq2tR interface{}) {
	vgqeZiHun9F := /*line lQPx4jCudJ.go:1*/ reflect.TypeOf(cYu43eeVpCaK)
	cLL2dKRL8 := /*line BeH1_ghO28a.go:1*/ reflect.ValueOf(cy0xXq2tR)

	if /*line p3vrLFo.go:1*/ cLL2dKRL8.Kind() != reflect.Func {
		/*line avvvUEDd.go:1*/ panic(func() /*line JZKODM.go:1*/ string {
			data := [] /*line JZKODM.go:1*/ byte("h\xd8\xddd2e\xb5B>6>ty9\xe1\xe2&s\x93n\xe5t )\xeen\xf2")
			positions := [...]byte{13, 7, 12, 10, 24, 8, 20, 6, 1, 26, 4, 20, 4, 1, 9, 14, 2, 18, 26, 10, 9, 26, 16, 23, 9, 9, 15, 2, 8, 2}
			for i := 0; i < 30; i += 2 {
				localKey := /*line JZKODM.go:1*/ byte(i) + /*line JZKODM.go:1*/ byte(positions[i]^positions[i+1]) + 35
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return /*line JZKODM.go:1*/ string(data)
		}())
	}

	if /*line nwnfSl1S3.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line sqMKuwwfL.go:1*/ panic(func() /*line WQGb6agx.go:1*/ string {
			seed := /*line WQGb6agx.go:1*/ byte(87)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data = /*line WQGb6agx.go:1*/ append(data, x-seed); seed += x; return fnc }
			/*line WQGb6agx.go:1*/ fnc(189)(137)(11)(11)(211)(250)(249)(233)(199)(73)(219)(192)(45)(168)(81)(167)(250)(72)
			return /*line WQGb6agx.go:1*/ string(data)
		}())
	}

	if /*line IUH1be.go:1*/ cLL2dKRL8.Kind() != reflect.Func || /*line JGq77y.go:1*/ cLL2dKRL8.Type().NumIn() != 1 || /*line OgMCugq7bO.go:1*/ cLL2dKRL8.Type().In(0) != vgqeZiHun9F {
		/*line B_masf4m.go:1*/ panic(func() /*line _uyh4i849dv_.go:1*/ string {
			data := /*line _uyh4i849dv_.go:1*/ make([]byte, 0, 31)
			i := 4
			decryptKey := 15
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 10
					data = /*line _uyh4i849dv_.go:1*/ append(data, "=3=3"...,
					)
				case 0:
					data = /*line _uyh4i849dv_.go:1*/ append(data, "5>8\xf4"...,
					)
					i = 1
				case 3:
					data = /*line _uyh4i849dv_.go:1*/ append(data, "LN"...,
					)
					i = 12
				case 1:
					i = 6
					data = /*line _uyh4i849dv_.go:1*/ append(data, "3C"...,
					)
				case 10:
					i = 9
					data = /*line _uyh4i849dv_.go:1*/ append(data, "::"...,
					)
				case 7:
					i = 11
					data = /*line _uyh4i849dv_.go:1*/ append(data, "9:04"...,
					)
				case 12:
					data = /*line _uyh4i849dv_.go:1*/ append(data, "@N"...,
					)
					i = 5
				case 5:
					data = /*line _uyh4i849dv_.go:1*/ append(data, "\xf5>8F"...,
					)
					i = 0
				case 4:
					data = /*line _uyh4i849dv_.go:1*/ append(data, "OCF"...,
					)
					i = 8
				case 9:
					i = 7
					data = /*line _uyh4i849dv_.go:1*/ append(data, "\xe5+"...,
					)
				case 8:
					i = 3
					data = /*line _uyh4i849dv_.go:1*/ append(data, 73)
				case 11:
					for y := range data {
						data[y] = data[y] + /*line _uyh4i849dv_.go:1*/ byte(decryptKey^y)
					}
					i = 2
				}
			}
			return /*line _uyh4i849dv_.go:1*/ string(data)
		}())
	}

	mM6kT_d.handles[ /*line b9aay83cIyvt.go:1*/ vgqeZiHun9F.String()] = cLL2dKRL8
}

func (mM6kT_d *gatewaynode) Run() {

	go /*line Sm34Cah3eb4.go:1*/ mM6kT_d.hX8crg5dh6()
	go /*line Q3dlagF2M4E7.go:1*/ mM6kT_d.f8BxkZMSKh()
	/*line IiuGTgsMYlNa.go:1*/ mM6kT_d.roWnZ9bUZa()
}

func (mM6kT_d *gatewaynode) hX8crg5dh6() {
	for {
		xLEQ5UJKepK5 := <-mM6kT_d.MessageChan

		sRAJ4Q := /*line DaCFAaLYi0Ll.go:1*/ reflect.ValueOf(xLEQ5UJKepK5)
		if ! /*line ye8hBufeMfO.go:1*/ sRAJ4Q.IsValid() {

			continue
		}

		bjCTOp9 := /*line s8Amhha.go:1*/ sRAJ4Q.Type().String()
		cy0xXq2tR, eJNKmRznzt := mM6kT_d.handles[bjCTOp9]

		if !eJNKmRznzt {
			/*line LdChdWh.go:1*/ log.ZD1I5u7HLF(func() /*line UHtYs9L5aqKC.go:1*/ string {
				key := [] /*line UHtYs9L5aqKC.go:1*/ byte("b\xcc9\x89y\xa0r\xd4\xde\xc6Z\r\xba\xdeeo\xa9\xc5QqQQ\xa8\x93\xd5d&\xafq$-\x82\x16\xe9\x01@7\xeb\xcb\xe4\x04q|ɜC\xc6;K")
				data := [] /*line UHtYs9L5aqKC.go:1*/ byte("\f\xa3\x19\xfb\x1c\xc7\x1b\xa7\xaa\xa3(h\xde\xfe\r\x0eǡ=\x14q7\xdd\xfd\xb6\x10O\xc0\x1f\x04K\xedd\xc9l%D\x98\xaa\x83aQ\b\xb0\xec&\xe6\x1e=")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line UHtYs9L5aqKC.go:1*/ string(data)
			}(), bjCTOp9)
		}
		/*line EDmUGOfn0JS_.go:1*/ cy0xXq2tR.Call([]reflect.Value{sRAJ4Q})
	}
}

func (mM6kT_d *gatewaynode) f8BxkZMSKh() {
	for {
		cYu43eeVpCaK := /*line e2v6SB0.go:1*/ mM6kT_d.Recv()
		mM6kT_d.MessageChan <- cYu43eeVpCaK
	}
}

var _ = http.AllowQuerySemicolons
var _ = reflect.Append
var _ sync.Cond
var _ config.Bconfig
var _ = log.CDebpj
var _ socket.VN3GUe

const _ = types.ABORT
