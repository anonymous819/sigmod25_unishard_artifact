//line :1
package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	log "unishard/log"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const path = "../common/"

var eALIfA6r04PU = /*line JbdLRlVKQtT.go:1*/ flag.String("config", func() /*line EdUSjMeW6a.go:1*/ string {
	key := [] /*line EdUSjMeW6a.go:1*/ byte("*&\x8e\xc87Cŝ\x8e\xb8f")
	data := [] /*line EdUSjMeW6a.go:1*/ byte("\x8d\x95\xfc.\xa0\xaa\xf3\a\x01'\xd4")
	for i, b := range key {
		data[i] = data[i] - b
	}
	return /*line EdUSjMeW6a.go:1*/ string(data)
}(), func() /*line GTaONZo.go:1*/ string {
	fullData := [] /*line GTaONZo.go:1*/ byte("\xe7ٰMe\xfb\xb4\xff\xf3U\x16ݪ\x90\xfe\x83|\x86D^Ҋ2\xbd\xdbr\xa6\x93v\xd5\x18\xd6刍\xef\x91s]\xe0\x88\xad\xe4\x9b\x1b[\x93\xa0;\x1b\x97C\xec:\xe1\xc6\x1e\xc1\x9c\x04\x85\x92K\x92N\xf3\x03uP\xfdr\xa8+\fA\xa6\xd0\x05V\xc47\xfchf\xcf\t\x04\x03\xc0\x8dF\xf1MpW]^5\xd9R\xe4G\xfd.\x11\xa6\xbf\xf6\x9a\x1dm\x81\xe4;\xf0\xc8*\x9e\xe7v\xcf\xd9dŉ\xf8")
	idxKey := [] /*line GTaONZo.go:1*/ byte("\xb0\xf6}\xf3\xf5q\x8e")
	data := /*line GTaONZo.go:1*/ make([]byte, 0, 64)
	data = /*line GTaONZo.go:1*/ append(data, fullData[232^ /*line GTaONZo.go:1*/ int(idxKey[1])]-fullData[235^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[183^ /*line GTaONZo.go:1*/ int(idxKey[3])]-fullData[197^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[105^ /*line GTaONZo.go:1*/ int(idxKey[5])]+fullData[95^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[228^ /*line GTaONZo.go:1*/ int(idxKey[3])]-fullData[173^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[159^ /*line GTaONZo.go:1*/ int(idxKey[1])]^fullData[162^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[10^ /*line GTaONZo.go:1*/ int(idxKey[2])]+fullData[38^ /*line GTaONZo.go:1*/ int(idxKey[2])], fullData[200^ /*line GTaONZo.go:1*/ int(idxKey[1])]-fullData[233^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[157^ /*line GTaONZo.go:1*/ int(idxKey[3])]+fullData[190^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[69^ /*line GTaONZo.go:1*/ int(idxKey[5])]+fullData[50^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[81^ /*line GTaONZo.go:1*/ int(idxKey[5])]^fullData[85^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[179^ /*line GTaONZo.go:1*/ int(idxKey[0])]-fullData[192^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[229^ /*line GTaONZo.go:1*/ int(idxKey[3])]^fullData[213^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[104^ /*line GTaONZo.go:1*/ int(idxKey[2])]^fullData[87^ /*line GTaONZo.go:1*/ int(idxKey[2])], fullData[55^ /*line GTaONZo.go:1*/ int(idxKey[5])]-fullData[18^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[191^ /*line GTaONZo.go:1*/ int(idxKey[1])]-fullData[189^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[130^ /*line GTaONZo.go:1*/ int(idxKey[1])]-fullData[207^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[244^ /*line GTaONZo.go:1*/ int(idxKey[3])]^fullData[232^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[38^ /*line GTaONZo.go:1*/ int(idxKey[5])]-fullData[4^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[210^ /*line GTaONZo.go:1*/ int(idxKey[4])]-fullData[173^ /*line GTaONZo.go:1*/ int(idxKey[4])], fullData[30^ /*line GTaONZo.go:1*/ int(idxKey[5])]-fullData[64^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[187^ /*line GTaONZo.go:1*/ int(idxKey[6])]+fullData[239^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[221^ /*line GTaONZo.go:1*/ int(idxKey[4])]-fullData[255^ /*line GTaONZo.go:1*/ int(idxKey[4])], fullData[236^ /*line GTaONZo.go:1*/ int(idxKey[1])]^fullData[231^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[157^ /*line GTaONZo.go:1*/ int(idxKey[4])]^fullData[208^ /*line GTaONZo.go:1*/ int(idxKey[4])], fullData[144^ /*line GTaONZo.go:1*/ int(idxKey[1])]+fullData[140^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[136^ /*line GTaONZo.go:1*/ int(idxKey[3])]^fullData[180^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[154^ /*line GTaONZo.go:1*/ int(idxKey[6])]+fullData[131^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[140^ /*line GTaONZo.go:1*/ int(idxKey[6])]-fullData[196^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[34^ /*line GTaONZo.go:1*/ int(idxKey[5])]+fullData[36^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[92^ /*line GTaONZo.go:1*/ int(idxKey[5])]-fullData[65^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[172^ /*line GTaONZo.go:1*/ int(idxKey[0])]^fullData[230^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[239^ /*line GTaONZo.go:1*/ int(idxKey[1])]+fullData[254^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[209^ /*line GTaONZo.go:1*/ int(idxKey[3])]^fullData[182^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[223^ /*line GTaONZo.go:1*/ int(idxKey[6])]+fullData[211^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[70^ /*line GTaONZo.go:1*/ int(idxKey[2])]+fullData[121^ /*line GTaONZo.go:1*/ int(idxKey[2])], fullData[29^ /*line GTaONZo.go:1*/ int(idxKey[5])]-fullData[33^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[194^ /*line GTaONZo.go:1*/ int(idxKey[4])]+fullData[222^ /*line GTaONZo.go:1*/ int(idxKey[4])], fullData[178^ /*line GTaONZo.go:1*/ int(idxKey[3])]^fullData[248^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[0^ /*line GTaONZo.go:1*/ int(idxKey[5])]^fullData[93^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[239^ /*line GTaONZo.go:1*/ int(idxKey[0])]+fullData[198^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[129^ /*line GTaONZo.go:1*/ int(idxKey[6])]-fullData[182^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[63^ /*line GTaONZo.go:1*/ int(idxKey[5])]-fullData[3^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[163^ /*line GTaONZo.go:1*/ int(idxKey[0])]+fullData[242^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[128^ /*line GTaONZo.go:1*/ int(idxKey[3])]+fullData[218^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[204^ /*line GTaONZo.go:1*/ int(idxKey[0])]-fullData[221^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[243^ /*line GTaONZo.go:1*/ int(idxKey[3])]+fullData[170^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[26^ /*line GTaONZo.go:1*/ int(idxKey[5])]^fullData[77^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[210^ /*line GTaONZo.go:1*/ int(idxKey[0])]+fullData[213^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[139^ /*line GTaONZo.go:1*/ int(idxKey[1])]+fullData[230^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[208^ /*line GTaONZo.go:1*/ int(idxKey[0])]-fullData[147^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[175^ /*line GTaONZo.go:1*/ int(idxKey[6])]-fullData[220^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[145^ /*line GTaONZo.go:1*/ int(idxKey[1])]^fullData[170^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[94^ /*line GTaONZo.go:1*/ int(idxKey[5])]^fullData[9^ /*line GTaONZo.go:1*/ int(idxKey[5])], fullData[113^ /*line GTaONZo.go:1*/ int(idxKey[2])]^fullData[50^ /*line GTaONZo.go:1*/ int(idxKey[2])], fullData[201^ /*line GTaONZo.go:1*/ int(idxKey[0])]^fullData[218^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[181^ /*line GTaONZo.go:1*/ int(idxKey[0])]-fullData[141^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[190^ /*line GTaONZo.go:1*/ int(idxKey[0])]-fullData[130^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[204^ /*line GTaONZo.go:1*/ int(idxKey[3])]+fullData[201^ /*line GTaONZo.go:1*/ int(idxKey[3])], fullData[206^ /*line GTaONZo.go:1*/ int(idxKey[6])]-fullData[234^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[131^ /*line GTaONZo.go:1*/ int(idxKey[0])]-fullData[252^ /*line GTaONZo.go:1*/ int(idxKey[0])], fullData[228^ /*line GTaONZo.go:1*/ int(idxKey[1])]^fullData[190^ /*line GTaONZo.go:1*/ int(idxKey[1])], fullData[136^ /*line GTaONZo.go:1*/ int(idxKey[6])]-fullData[212^ /*line GTaONZo.go:1*/ int(idxKey[6])], fullData[244^ /*line GTaONZo.go:1*/ int(idxKey[4])]+fullData[252^ /*line GTaONZo.go:1*/ int(idxKey[4])])
	return /*line GTaONZo.go:1*/ string(data)
}())
var piaXrZeka = /*line YQhnMqGe.go:1*/ flag.String("ips", "ips.txt", func() /*line Nwi9OC9t.go:1*/ string {
	data := /*line Nwi9OC9t.go:1*/ make([]byte, 0, 46)
	i := 11
	decryptKey := 174
	for counter := 0; i != 13; counter++ {
		decryptKey ^= i * counter
		switch i {
		case 2:
			i = 13
			for y := range data {
				data[y] = data[y] ^ /*line Nwi9OC9t.go:1*/ byte(decryptKey^y)
			}
		case 14:
			data = /*line Nwi9OC9t.go:1*/ append(data, 76)
			i = 18
		case 0:
			i = 5
			data = /*line Nwi9OC9t.go:1*/ append(data, "17 8"...,
			)
		case 12:
			i = 16
			data = /*line Nwi9OC9t.go:1*/ append(data, 49)
		case 16:
			i = 7
			data = /*line Nwi9OC9t.go:1*/ append(data, "+))y"...,
			)
		case 11:
			i = 4
			data = /*line Nwi9OC9t.go:1*/ append(data, 2)
		case 6:
			i = 10
			data = /*line Nwi9OC9t.go:1*/ append(data, 0)
		case 5:
			data = /*line Nwi9OC9t.go:1*/ append(data, "\x1f\x19I\x1c"...,
			)
			i = 6
		case 18:
			i = 2
			data = /*line Nwi9OC9t.go:1*/ append(data, "\x15\x18\x13"...,
			)
		case 15:
			i = 3
			data = /*line Nwi9OC9t.go:1*/ append(data, 63)
		case 3:
			i = 1
			data = /*line Nwi9OC9t.go:1*/ append(data, "l5#"...,
			)
		case 8:
			data = /*line Nwi9OC9t.go:1*/ append(data, 33)
			i = 15
		case 9:
			data = /*line Nwi9OC9t.go:1*/ append(data, 47)
			i = 17
		case 7:
			i = 9
			data = /*line Nwi9OC9t.go:1*/ append(data, "60:8"...,
			)
		case 4:
			data = /*line Nwi9OC9t.go:1*/ append(data, "::h)"...,
			)
			i = 8
		case 10:
			data = /*line Nwi9OC9t.go:1*/ append(data, "N\x04\x1c\x10"...,
			)
			i = 14
		case 1:
			i = 12
			data = /*line Nwi9OC9t.go:1*/ append(data, "-)#'"...,
			)
		case 17:
			data = /*line Nwi9OC9t.go:1*/ append(data, "}r\x155"...,
			)
			i = 0
		}
	}
	return /*line Nwi9OC9t.go:1*/ string(data)
}())
var vg42FpQQa5 = /*line NbWy18o_Hks2.go:1*/ flag.String("baseIPs", func() /*line mtGkd4.go:1*/ string {
	fullData := [] /*line mtGkd4.go:1*/ byte("\xd4?\xc0Սi\xa5\x7f \x82\xe7B\xe9\xbf\b\xa0\xafZ\xf1|L3\xf6R")
	idxKey := [] /*line mtGkd4.go:1*/ byte("\x93\x8a\x828r\xff9\x87\x16.")
	data := /*line mtGkd4.go:1*/ make([]byte, 0, 13)
	data = /*line mtGkd4.go:1*/ append(data, fullData[155^ /*line mtGkd4.go:1*/ int(idxKey[0])]+fullData[152^ /*line mtGkd4.go:1*/ int(idxKey[0])], fullData[136^ /*line mtGkd4.go:1*/ int(idxKey[7])]-fullData[134^ /*line mtGkd4.go:1*/ int(idxKey[7])], fullData[103^ /*line mtGkd4.go:1*/ int(idxKey[4])]-fullData[112^ /*line mtGkd4.go:1*/ int(idxKey[4])], fullData[143^ /*line mtGkd4.go:1*/ int(idxKey[2])]-fullData[147^ /*line mtGkd4.go:1*/ int(idxKey[2])], fullData[43^ /*line mtGkd4.go:1*/ int(idxKey[9])]+fullData[56^ /*line mtGkd4.go:1*/ int(idxKey[9])], fullData[1^ /*line mtGkd4.go:1*/ int(idxKey[8])]-fullData[26^ /*line mtGkd4.go:1*/ int(idxKey[8])], fullData[21^ /*line mtGkd4.go:1*/ int(idxKey[8])]^fullData[16^ /*line mtGkd4.go:1*/ int(idxKey[8])], fullData[49^ /*line mtGkd4.go:1*/ int(idxKey[3])]^fullData[42^ /*line mtGkd4.go:1*/ int(idxKey[3])], fullData[151^ /*line mtGkd4.go:1*/ int(idxKey[7])]+fullData[128^ /*line mtGkd4.go:1*/ int(idxKey[7])], fullData[245^ /*line mtGkd4.go:1*/ int(idxKey[5])]+fullData[251^ /*line mtGkd4.go:1*/ int(idxKey[5])], fullData[158^ /*line mtGkd4.go:1*/ int(idxKey[1])]-fullData[138^ /*line mtGkd4.go:1*/ int(idxKey[1])], fullData[241^ /*line mtGkd4.go:1*/ int(idxKey[5])]^fullData[236^ /*line mtGkd4.go:1*/ int(idxKey[5])])
	return /*line mtGkd4.go:1*/ string(data)
}(), func() /*line cutF48E.go:1*/ string {
	data := [] /*line cutF48E.go:1*/ byte("I\x7fR\x1f0\x01\xf4\x1eb s!UD\x9f\xa4\xb0\"#;\x00\xe5\r\t\f\xa8\xa41s&\xb3M~es\xb9\x98\x13lwF3-\x19s\x1b5\xf7 \n\xc5W\xeb\xe7x\x1a")
	positions := [...]byte{18, 46, 3, 38, 9, 42, 38, 24, 25, 11, 35, 1, 46, 46, 46, 25, 49, 11, 40, 7, 45, 2, 16, 18, 43, 17, 27, 41, 30, 30, 5, 55, 27, 24, 31, 29, 29, 51, 19, 26, 14, 22, 52, 15, 45, 32, 41, 36, 18, 24, 18, 12, 47, 1, 30, 39, 13, 43, 41, 30, 20, 20, 40, 4, 40, 37, 13, 50, 49, 4, 53, 22, 6, 19, 25, 23, 21, 37, 43, 24}
	for i := 0; i < 80; i += 2 {
		localKey := /*line cutF48E.go:1*/ byte(i) + /*line cutF48E.go:1*/ byte(positions[i]^positions[i+1]) + 37
		data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
	}
	return /*line cutF48E.go:1*/ string(data)
}())

type Config struct {
	Addrs     map[types.Shard]map[types.NodeID]string `json:"address"`
	HTTPAddrs map[types.NodeID]string                 `json:"http_address"`

	Policy    string  `json:"policy"`
	Threshold float64 `json:"threshold"`

	Thrifty           bool         `json:"thrifty"`
	BufferSize        int          `json:"buffer_size"`
	ChanBufferSize    int          `json:"chan_buffer_size"`
	MultiVersion      bool         `json:"multiversion"`
	Timeout           int          `json:"timeout"`
	ViewChangeTimeout int          `json:"viewchange_timeout"`
	ByzNo             int          `json:"byzNo"`
	BSize             int          `json:"bsize"`
	Fixed             bool         `json:"fixed"`
	Benchmark         Bconfig      `json:"benchmark"`
	Delta             int          `json:"delta"`
	Pprof             bool         `json:"pprof"`
	MaxRound          int          `json:"maxRound"`
	Strategy          string       `json:"strategy"`
	PayloadSize       int          `json:"payload_size"`
	Master            types.NodeID `json:"master"`
	Delay             int          `json:"delay"`
	DErr              int          `json:"derr"`
	MemSize           int          `json:"memsize"`
	Slow              int          `json:"slow"`
	Crash             int          `json:"crash"`

	ShardCount       int    `json:"shard_count"`
	DefaultBalance   int    `json:"default_balance"`
	ClientNumber     int    `json:"client_number"`
	ViewChangePeriod int    `json:"viewchange_period"`
	CommitteeNumber  int    `json:"committee_number"`
	RotatingElection string `json:"rotating_election"`
	InShardLatency   int    `json:"inshardlatency"`
	Shard01Latency   int    `json:"shard01latency"`
	Shard02Latency   int    `json:"shard02latency"`
	Shard03Latency   int    `json:"shard03latency"`
	Shard04Latency   int    `json:"shard04latency"`

	CoordinationShardBlockbuilderAddress string `json:"coordination_builder_address"`

	hasher string
	signer string

	n     int
	nBase int
}

type Bconfig struct {
	T            int
	N            int
	K            int
	Throttle     int
	Concurrency  int
	Distribution string

	Conflicts int
	Min       int

	Mu    float64
	Sigma float64
	Move  bool
	Speed int

	TXRatioPerType []int
	ZipfTheta      float64
}

type DwUGFadFbpap struct {
	F7zSaMJ common.Address
	T00fo63 []common.Address
}

var Configuration Config

func IXW32OGaz0J() []*common.Address {
	oM9zjG_eet, ypSeUZh := /*line R6luhXaNrx.go:1*/ os.ReadFile(func() /*line I912bSGY8B.go:1*/ string {
		data := [] /*line I912bSGY8B.go:1*/ byte("..\xaac7mG\xbf\x91\x8aad\xfer\x9a\xf8\xb3.\x06x\xf7")
		positions := [...]byte{2, 6, 9, 18, 4, 4, 4, 8, 6, 18, 2, 2, 15, 20, 9, 16, 12, 14, 4, 16, 12, 8, 7, 2}
		for i := 0; i < 24; i += 2 {
			localKey := /*line I912bSGY8B.go:1*/ byte(i) + /*line I912bSGY8B.go:1*/ byte(positions[i]^positions[i+1]) + 85
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line I912bSGY8B.go:1*/ string(data)
	}())
	if ypSeUZh != nil {
		/*line ahS6uNyY.go:1*/ panic(ypSeUZh)
	}
	ub3eM_aOI := /*line Sjw5VSZz.go:1*/ strings.Split( /*line ItXiYWf3SpL.go:1*/ string(oM9zjG_eet), "\n")
	guGUKEwm := /*line BZq4NNfb.go:1*/ make([]*common.Address, 0)
	for aas4qNe6 := 0; aas4qNe6 < /*line cwXYYrDpT8cM.go:1*/ len(ub3eM_aOI); aas4qNe6++ {
		eaUkMXV := /*line yFKBCfkmBKdM.go:1*/ common.HexToAddress(ub3eM_aOI[aas4qNe6])
		guGUKEwm = /*line bvqZma.go:1*/ append(guGUKEwm, &eaUkMXV)
	}
	return guGUKEwm
}

func GetConfig() Config {
	return Configuration
}

func b2twoSCi(yD1VVZBMMym error) {
	if yD1VVZBMMym != nil {
		/*line hyxDVzSGmCC.go:1*/ panic(yD1VVZBMMym)
	}
}

func LuxobPjs() (gESTB7YAm []common.Address, kjNXaQ, cJMbVlTiRG, uYd5Z0hKdfK [][]common.Address) {
	kjNXaQ = /*line MsEw6rD.go:1*/ make([][]common.Address, 0)
	cJMbVlTiRG = /*line fPhqsD1YGIN7.go:1*/ make([][]common.Address, 0)
	uYd5Z0hKdfK = /*line UWzbgw.go:1*/ make([][]common.Address, 0)
	gESTB7YAm = /*line OUUg8wGnc.go:1*/ make([]common.Address, 0)
	q12uG9VmOEC := []common.Address{}

	fIUszI := func() /*line tam0CD1DKLab.go:1*/ string {
		data := [] /*line tam0CD1DKLab.go:1*/ byte("\xd9.\u0082\xcd{\x80\xcd\xcf\xc54}/")
		positions := [...]byte{4, 7, 8, 3, 10, 9, 5, 9, 11, 8, 9, 11, 6, 2, 5, 8, 8, 0}
		for i := 0; i < 18; i += 2 {
			localKey := /*line tam0CD1DKLab.go:1*/ byte(i) + /*line tam0CD1DKLab.go:1*/ byte(positions[i]^positions[i+1]) + 159
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line tam0CD1DKLab.go:1*/ string(data)
	}() + /*line sLv1tLsn.go:1*/ strconv.Itoa(Configuration.ShardCount) + "/"

	for aas4qNe6 := 1; aas4qNe6 <= Configuration.ShardCount; aas4qNe6++ {
		dylCayz, ypSeUZh := /*line MVQmTL3a.go:1*/ os.ReadFile(fIUszI + "payment" + /*line WHE8xbCZ0.go:1*/ strconv.Itoa(aas4qNe6) + ".txt")
		/*line zICJcmLur.go:1*/ b2twoSCi(ypSeUZh)
		for fvBnQMMl_Ehh, _ZhVNkCZSi := range /*line ZUIBp9qr7c.go:1*/ strings.Split( /*line _QGVBeC.go:1*/ string(dylCayz), "\n") {
			if fvBnQMMl_Ehh == 0 {
				q12uG9VmOEC = /*line TIneniM.go:1*/ append(q12uG9VmOEC /*line PduRLRDFRJ.go:1*/, common.HexToAddress(_ZhVNkCZSi))
			}
			gESTB7YAm = /*line BTtsjcC.go:1*/ append(gESTB7YAm /*line yxbp5Bd.go:1*/, common.HexToAddress(_ZhVNkCZSi))
		}
	}

	for aas4qNe6 := 1; aas4qNe6 <= Configuration.ShardCount; aas4qNe6++ {
		for oReRWbQ8Z := 1; oReRWbQ8Z <= Configuration.ShardCount; oReRWbQ8Z++ {
			hD0zm5jtKxFl, ypSeUZh := /*line CetTpG4V.go:1*/ os.ReadFile(fIUszI + "payment" + /*line JbGfafP4YrS.go:1*/ strconv.Itoa(aas4qNe6) + "_shard" + /*line CmHgzlMqx.go:1*/ strconv.Itoa(oReRWbQ8Z) + func() /*line o2fZ3twB_hOW.go:1*/ string {
				fullData := [] /*line o2fZ3twB_hOW.go:1*/ byte("\x9c\x17$1d\xc5\x15\xd2=\xb7P\xd2[\xbb_\x98$\x10$\x06q]\x91\xf1")
				idxKey := [] /*line o2fZ3twB_hOW.go:1*/ byte("Q]")
				data := /*line o2fZ3twB_hOW.go:1*/ make([]byte, 0, 13)
				data = /*line o2fZ3twB_hOW.go:1*/ append(data, fullData[67^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]-fullData[84^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])], fullData[64^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]^fullData[85^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])], fullData[80^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])]+fullData[84^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])], fullData[78^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])]+fullData[81^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])], fullData[94^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]^fullData[70^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])], fullData[94^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])]+fullData[85^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])], fullData[91^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])]-fullData[86^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])], fullData[86^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]-fullData[71^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])], fullData[69^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]^fullData[95^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])], fullData[87^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])]^fullData[95^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[1])], fullData[81^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]-fullData[65^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])], fullData[68^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])]+fullData[80^ /*line o2fZ3twB_hOW.go:1*/ int(idxKey[0])])
				return /*line o2fZ3twB_hOW.go:1*/ string(data)
			}())
			/*line W9k6rT.go:1*/ b2twoSCi(ypSeUZh)
			for _, _ZhVNkCZSi := range /*line IrIFHoD.go:1*/ strings.Split( /*line JsuX8O2WuO_G.go:1*/ string(hD0zm5jtKxFl), "\n") {
				hD0zm5jtKxFl := /*line AzpQi73n.go:1*/ make([]common.Address, 0)
				hD0zm5jtKxFl = /*line DfL3EjcLfM.go:1*/ append(hD0zm5jtKxFl /*line RDzXnKT.go:1*/, common.HexToAddress(_ZhVNkCZSi))
				hD0zm5jtKxFl = /*line rpedvyhG.go:1*/ append(hD0zm5jtKxFl, q12uG9VmOEC[aas4qNe6-1])
				kjNXaQ = /*line H08vnMEw53RY.go:1*/ append(kjNXaQ, hD0zm5jtKxFl)
			}

			topfTn1M4t, ypSeUZh := /*line _FRwxoYC.go:1*/ os.ReadFile(fIUszI + "payment" + /*line boZZ8FzRd.go:1*/ strconv.Itoa(aas4qNe6) + "_shard" + /*line SFrvHDa.go:1*/ strconv.Itoa(oReRWbQ8Z) + func() /*line k5JRFP_q.go:1*/ string {
				data := /*line k5JRFP_q.go:1*/ make([]byte, 0, 13)
				i := 5
				decryptKey := 18
				for counter := 0; i != 6; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 0:
						i = 2
						data = /*line k5JRFP_q.go:1*/ append(data, 93)
					case 2:
						i = 3
						data = /*line k5JRFP_q.go:1*/ append(data, "gkW]"...,
						)
					case 5:
						data = /*line k5JRFP_q.go:1*/ append(data, 85)
						i = 0
					case 4:
						i = 1
						data = /*line k5JRFP_q.go:1*/ append(data, "xs"...,
						)
					case 3:
						data = /*line k5JRFP_q.go:1*/ append(data, "74,q"...,
						)
						i = 4
					case 1:
						for y := range data {
							data[y] = data[y] + /*line k5JRFP_q.go:1*/ byte(decryptKey^y)
						}
						i = 6
					}
				}
				return /*line k5JRFP_q.go:1*/ string(data)
			}())
			/*line weKtUGytSv4m.go:1*/ b2twoSCi(ypSeUZh)
			for _, _ZhVNkCZSi := range /*line CF2Jw1B.go:1*/ strings.Split( /*line yjxjNg.go:1*/ string(topfTn1M4t), "\n") {
				topfTn1M4t := /*line z7xXIIZ7KS.go:1*/ make([]common.Address, 0)
				topfTn1M4t = /*line GJBR40HEAsx.go:1*/ append(topfTn1M4t /*line _Ygda7Rw5.go:1*/, common.HexToAddress(_ZhVNkCZSi))
				topfTn1M4t = /*line KDMSL7nRV6V.go:1*/ append(topfTn1M4t, q12uG9VmOEC[aas4qNe6-1])
				cJMbVlTiRG = /*line FUhncEiq.go:1*/ append(cJMbVlTiRG, topfTn1M4t)
			}

			fZyRAe, ypSeUZh := /*line VCt0oy.go:1*/ os.ReadFile(fIUszI + "payment" + /*line fZsVcmNR.go:1*/ strconv.Itoa(aas4qNe6) + func() /*line S9QSRhGIWP.go:1*/ string {
				key := [] /*line S9QSRhGIWP.go:1*/ byte("@d\x83Hw1b\xfe\x91\xea\xff\xe4\xa5+\xd3lX̒")
				data := [] /*line S9QSRhGIWP.go:1*/ byte("\x1f\n\xec:\x1aP\x0e\xaa㋉\x81\xc9h\x92B,\xb4\xe6")
				for i, b := range key {
					data[i] = data[i] ^ b
				}
				return /*line S9QSRhGIWP.go:1*/ string(data)
			}())
			/*line PSzO1e.go:1*/ b2twoSCi(ypSeUZh)
			for _, _ZhVNkCZSi := range /*line aOPQHQWX.go:1*/ strings.Split( /*line lJ_z1hUhZfQ.go:1*/ string(fZyRAe), "\n") {
				dpP8rHtN := /*line DLBDrQR0wn.go:1*/ make([]common.Address, 0)
				for _, _ZhVNkCZSi := range /*line iDIlKfaUNt.go:1*/ strings.Split(_ZhVNkCZSi, ",") {
					dpP8rHtN = /*line EFE5VEu6CIM.go:1*/ append(dpP8rHtN /*line aIvUGVJ.go:1*/, common.HexToAddress(_ZhVNkCZSi))
				}
				uYd5Z0hKdfK = /*line DwDSLfh.go:1*/ append(uYd5Z0hKdfK, dpP8rHtN)
			}

			fZyRAe, ypSeUZh = /*line COucqw5Rl.go:1*/ os.ReadFile(fIUszI + "payment" + /*line SW5ayCMp4NNg.go:1*/ strconv.Itoa(aas4qNe6) + func() /*line SvCsQM3E_74.go:1*/ string {
				data := /*line SvCsQM3E_74.go:1*/ make([]byte, 0, 24)
				i := 7
				decryptKey := 248
				for counter := 0; i != 1; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 9:
						i = 10
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xce\xdc\xc9\xf8"...,
						)
					case 4:
						i = 1
						for y := range data {
							data[y] = data[y] ^ /*line SvCsQM3E_74.go:1*/ byte(decryptKey^y)
						}
					case 5:
						i = 8
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xf4\x9a\xc7"...,
						)
					case 3:
						data = /*line SvCsQM3E_74.go:1*/ append(data, 208)
						i = 2
					case 10:
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xd9\xcb"...,
						)
						i = 6
					case 6:
						i = 5
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xdf\xcd\xdb\xf5"...,
						)
					case 8:
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xca\xc5"...,
						)
						i = 4
					case 2:
						i = 11
						data = /*line SvCsQM3E_74.go:1*/ append(data, 209)
					case 0:
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xd7\xcb"...,
						)
						i = 3
					case 11:
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xf2\xc8"...,
						)
						i = 9
					case 7:
						i = 0
						data = /*line SvCsQM3E_74.go:1*/ append(data, "\xf8\xc5"...,
						)
					}
				}
				return /*line SvCsQM3E_74.go:1*/ string(data)
			}())
			/*line EDBR1FJ.go:1*/ b2twoSCi(ypSeUZh)
			for _, _ZhVNkCZSi := range /*line M9dym4p.go:1*/ strings.Split( /*line h3syCXgsM.go:1*/ string(fZyRAe), "\n") {
				dpP8rHtN := /*line FMJgrkMMhy.go:1*/ make([]common.Address, 0)
				for _, _ZhVNkCZSi := range /*line _50Tavma8F4.go:1*/ strings.Split(_ZhVNkCZSi, ",") {
					dpP8rHtN = /*line GMkfU8.go:1*/ append(dpP8rHtN /*line HgwZB1I.go:1*/, common.HexToAddress(_ZhVNkCZSi))
				}
				uYd5Z0hKdfK = /*line HGHMg5.go:1*/ append(uYd5Z0hKdfK, dpP8rHtN)
			}
		}
	}
	return gESTB7YAm, kjNXaQ, cJMbVlTiRG, uYd5Z0hKdfK
}

func UVxlf_QrYL(gWG7GbeU string) map[string]types.J_zoDMw {
	var qlxpfbft []types.FGQEDLInwsw
	ursrhE := /*line vPHBRK.go:1*/ make(map[string]types.J_zoDMw)
	qSUukID, ypSeUZh := /*line Iq12OPnxbSH.go:1*/ os.ReadFile(func() /*line yoXTOLL.go:1*/ string {
		seed := /*line yoXTOLL.go:1*/ byte(115)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line yoXTOLL.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line yoXTOLL.go:1*/ fnc(187)(0)(1)(52)(12)(254)(0)(2)(255)(193)(52)(12)(255)(6)(254)(239)(2)(17)(255)(188)
		return /*line yoXTOLL.go:1*/ string(data)
	}() + gWG7GbeU + func() /*line jMaDlAG.go:1*/ string {
		seed := /*line jMaDlAG.go:1*/ byte(176)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line jMaDlAG.go:1*/ append(data, x^seed); seed += x; return fnc }
		/*line jMaDlAG.go:1*/ fnc(159)(61)(251)(212)(62)(237)(168)(68)(1)(28)(225)
		return /*line jMaDlAG.go:1*/ string(data)
	}())
	if ypSeUZh != nil {
		/*line VytiUN7XtTSh.go:1*/ panic(ypSeUZh)
	}
	ypSeUZh = /*line Ldri_SKXjJ.go:1*/ json.Unmarshal(qSUukID, &qlxpfbft)
	/*line CLfDDkWck.go:1*/ b2twoSCi(ypSeUZh)

	for _, r60FsdJ9x := range qlxpfbft {
		ursrhE[r60FsdJ9x.EMQpRmSxByt] = types.J_zoDMw{
			ReadSet:      r60FsdJ9x.EZoqawtB8r4,
			WriteSet:     r60FsdJ9x.TLgnY5kyM,
			BtvFRP3tVe8k: r60FsdJ9x.WNHy5_Jk,
		}
	}
	return ursrhE
}

func EKXfbEt() abi.ABI {
	zqdfWn, _ := /*line eOHS3WPd.go:1*/ os.ReadFile(func() /*line isaowO.go:1*/ string {
		data := /*line isaowO.go:1*/ make([]byte, 0, 38)
		i := 15
		decryptKey := 3
		for counter := 0; i != 17; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 14:
				data = /*line isaowO.go:1*/ append(data, 6)
				i = 7
			case 8:
				i = 0
				data = /*line isaowO.go:1*/ append(data, 234)
			case 5:
				data = /*line isaowO.go:1*/ append(data, "\xb6\xfc"...,
				)
				i = 20
			case 4:
				i = 2
				data = /*line isaowO.go:1*/ append(data, "\v\x03\xc5\xfa"...,
				)
			case 20:
				data = /*line isaowO.go:1*/ append(data, "\xf3\xe3"...,
				)
				i = 11
			case 15:
				i = 19
				data = /*line isaowO.go:1*/ append(data, "\xcb\xcc\xce"...,
				)
			case 13:
				data = /*line isaowO.go:1*/ append(data, 252)
				i = 3
			case 19:
				i = 4
				data = /*line isaowO.go:1*/ append(data, "\x03\b\a\b"...,
				)
			case 9:
				i = 10
				data = /*line isaowO.go:1*/ append(data, "\xec \""...,
				)
			case 3:
				i = 8
				data = /*line isaowO.go:1*/ append(data, "\xec\x02"...,
				)
			case 2:
				i = 16
				data = /*line isaowO.go:1*/ append(data, 7)
			case 18:
				data = /*line isaowO.go:1*/ append(data, "\xf0\x02\x02"...,
				)
				i = 12
			case 16:
				i = 14
				data = /*line isaowO.go:1*/ append(data, 255)
			case 12:
				data = /*line isaowO.go:1*/ append(data, 191)
				i = 1
			case 0:
				data = /*line isaowO.go:1*/ append(data, 242)
				i = 5
			case 6:
				i = 17
				for y := range data {
					data[y] = data[y] + /*line isaowO.go:1*/ byte(decryptKey^y)
				}
			case 1:
				data = /*line isaowO.go:1*/ append(data, 253)
				i = 13
			case 7:
				i = 18
				data = /*line isaowO.go:1*/ append(data, "\x05\xf5"...,
				)
			case 10:
				data = /*line isaowO.go:1*/ append(data, 34)
				i = 6
			case 11:
				data = /*line isaowO.go:1*/ append(data, "\xf9\xe9)"...,
				)
				i = 9
			}
		}
		return /*line isaowO.go:1*/ string(data)
	}())
	m2ZQmpiaZ, _ := /*line Y2TaJMPnr.go:1*/ abi.JSON( /*line IVuA5aP.go:1*/ strings.NewReader( /*line FgzqUwb9MTpe.go:1*/ string(zqdfWn)))
	return m2ZQmpiaZ
}

func EGz0tn(wRwqCbbTs5U string, wAjEqkhzQS types.Shard) []DwUGFadFbpap {
	aaFFqt := /*line BOlKU6ZM.go:1*/ make([]DwUGFadFbpap, 0)
	l7XFJxH3k7 := Configuration.ShardCount
	gESTB7YAm := /*line UcNCc1c.go:1*/ make(map[types.Shard][]string)
	fIUszI := func() /*line DaZIVrYMo.go:1*/ string {
		data := [] /*line DaZIVrYMo.go:1*/ byte(".Ϛco\xc7m\\\xb4RcV\xda")
		positions := [...]byte{12, 7, 8, 2, 1, 5, 5, 7, 2, 11, 5, 1, 1, 9}
		for i := 0; i < 14; i += 2 {
			localKey := /*line DaZIVrYMo.go:1*/ byte(i) + /*line DaZIVrYMo.go:1*/ byte(positions[i]^positions[i+1]) + 200
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line DaZIVrYMo.go:1*/ string(data)
	}() + /*line dGIWSfK2.go:1*/ strconv.Itoa(l7XFJxH3k7) + "/"

	for aas4qNe6 := 1; aas4qNe6 <= l7XFJxH3k7; aas4qNe6++ {
		dylCayz, ypSeUZh := /*line GbOoCAfl.go:1*/ os.ReadFile(fIUszI + "payment" + /*line A8Y2YpilkpY.go:1*/ strconv.Itoa(aas4qNe6) + ".txt")
		/*line HBs_WJAxAo.go:1*/ b2twoSCi(ypSeUZh)
		gESTB7YAm[ /*line xs89QUIqKYK.go:1*/ types.Shard(aas4qNe6)] = /*line RcVu9pwae.go:1*/ strings.Split( /*line ferSTe0R.go:1*/ string(dylCayz), "\n")
	}
	dylCayz := /*line h9iHlS0.go:1*/ common.HexToAddress(gESTB7YAm[wAjEqkhzQS][0])

	uYd5Z0hKdfK, ypSeUZh := /*line ih9aTyFyv.go:1*/ os.ReadFile(fIUszI + "payment" + /*line HHt2rOuPbvV.go:1*/ strconv.Itoa( /*line XNzVNcO_W.go:1*/ int(wAjEqkhzQS)) + "_" + wRwqCbbTs5U + func() /*line YmcuqJcIf7oy.go:1*/ string {
		data := /*line YmcuqJcIf7oy.go:1*/ make([]byte, 0, 13)
		i := 1
		decryptKey := 154
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, 206)
				i = 5
			case 7:
				i = 6
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, 210)
			case 1:
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, "\xef\xc8\xd8"...,
				)
				i = 0
			case 6:
				i = 8
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, 254)
			case 5:
				i = 7
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, 218)
			case 8:
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, 253)
				i = 3
			case 3:
				data = /*line YmcuqJcIf7oy.go:1*/ append(data, "\x9d\xc6\xc9\xc4"...,
				)
				i = 4
			case 4:
				for y := range data {
					data[y] = data[y] ^ /*line YmcuqJcIf7oy.go:1*/ byte(decryptKey^y)
				}
				i = 2
			}
		}
		return /*line YmcuqJcIf7oy.go:1*/ string(data)
	}())
	/*line Dm180fyQ.go:1*/ b2twoSCi(ypSeUZh)
	f9rDctaLJ := /*line OaRKQi0.go:1*/ strings.Split( /*line OiGnKYZoHn.go:1*/ string(uYd5Z0hKdfK), "\n")

	for aas4qNe6 := 0; aas4qNe6 < /*line Oa4C34p05dpM.go:1*/ len(f9rDctaLJ); aas4qNe6++ {
		fZyRAe := /*line GJJkEBQt4.go:1*/ strings.Split(f9rDctaLJ[aas4qNe6], ",")
		aaFFqt = /*line _2OrUkEj4ksK.go:1*/ append(aaFFqt, DwUGFadFbpap{
			F7zSaMJ:/*line n58j1cPJiSPb.go:1*/ common.HexToAddress(fZyRAe[0]),
			T00fo63: []common.Address{ /*line wF4aAXyP.go:1*/ common.HexToAddress(fZyRAe[1]), dylCayz /*line cpYvLq4v.go:1*/, common.HexToAddress(fZyRAe[2]), dylCayz /*line pyhUPG.go:1*/, common.HexToAddress(fZyRAe[1]), dylCayz /*line NB670aAW8o5.go:1*/, common.HexToAddress(fZyRAe[2]), dylCayz},
		})
	}
	return aaFFqt
}

func (p7dFngSV Config) NBase() int {
	return p7dFngSV.nBase
}

func (p7dFngSV Config) N() int {
	return p7dFngSV.n
}

func (p7dFngSV Config) ClientN() int {
	return p7dFngSV.ClientNumber
}

func (p7dFngSV Config) NPerShard() int {
	return p7dFngSV.n / p7dFngSV.ShardCount
}

func (p7dFngSV *Config) Load() {
	gyEypFy, ypSeUZh := /*line hjcob9y.go:1*/ os.Open( /*line fncntiC.go:1*/ filepath.Join(func() /*line CBOw4Q7aT.go:1*/ string {
		data := [] /*line CBOw4Q7aT.go:1*/ byte("\x9e\xdd+v\x86mp\x8a-/")
		positions := [...]byte{7, 6, 0, 8, 3, 2, 2, 4, 2, 3, 7, 3, 0, 1}
		for i := 0; i < 14; i += 2 {
			localKey := /*line CBOw4Q7aT.go:1*/ byte(i) + /*line CBOw4Q7aT.go:1*/ byte(positions[i]^positions[i+1]) + 230
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line CBOw4Q7aT.go:1*/ string(data)
	}(), *eALIfA6r04PU))
	if ypSeUZh != nil {
		/*line A4HvfIfpP.go:1*/ log.FB5S6xdVix(ypSeUZh)
	}
	mFPgh6IE64 := /*line fFovUjrCAnV.go:1*/ json.NewDecoder(gyEypFy)
	ypSeUZh = /*line y0Kwi3VN9B.go:1*/ mFPgh6IE64.Decode(p7dFngSV)
	if ypSeUZh != nil {
		/*line TDOyowx.go:1*/ log.FB5S6xdVix(ypSeUZh)
	}

	m5nf8PzTB, ypSeUZh := /*line UavXoGETC.go:1*/ os.Open( /*line MNJX45.go:1*/ filepath.Join(func() /*line Ra5709gHVGZ.go:1*/ string {
		key := [] /*line Ra5709gHVGZ.go:1*/ byte("8R\x80RXCz\x1e\xb0\"")
		data := [] /*line Ra5709gHVGZ.go:1*/ byte("\xf6ܯ\x11\x17*\xf3Q\xbe\r")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line Ra5709gHVGZ.go:1*/ string(data)
	}(), *piaXrZeka))
	if ypSeUZh != nil {
		/*line Ab9n_ST7B4.go:1*/ log.FB5S6xdVix(ypSeUZh)
	}
	defer /*line aOayDar.go:1*/ m5nf8PzTB.Close()

	aS41VHGs91 := /*line f4SmH57.go:1*/ bufio.NewScanner(m5nf8PzTB)

	bzVQUcooi3Y0 := /*line VUW6rhpe.go:1*/ make(map[types.NodeID]string)
	aas4qNe6 := 1
	for /*line IiqJrG.go:1*/ aS41VHGs91.Scan() {
		qlX29FUy8 := /*line L6_WQxXqz.go:1*/ utils.NewNodeID(aas4qNe6)

		e96cCZo2Z := "tcp://" + /*line BrcHjsFPRNRp.go:1*/ aS41VHGs91.Text() + ":"
		hLE8VT := /*line XiPrnsu.go:1*/ strconv.Itoa(8069 + aas4qNe6)
		eRZE_Yx_Jeaz := "http://" + /*line C6zyWMPR9.go:1*/ aS41VHGs91.Text() + ":" + hLE8VT
		bzVQUcooi3Y0[qlX29FUy8] = e96cCZo2Z
		p7dFngSV.HTTPAddrs[qlX29FUy8] = eRZE_Yx_Jeaz
		aas4qNe6++
	}

	if ypSeUZh := /*line BxiMbeDHq6V.go:1*/ aS41VHGs91.Err(); ypSeUZh != nil {
		/*line FWmdyfS.go:1*/ fmt.Println(ypSeUZh)
	}

	p7dFngSV.n = /*line HzvPlhRHfiI.go:1*/ len(bzVQUcooi3Y0)

	for aas4qNe6 := 0; aas4qNe6 <= p7dFngSV.ShardCount; aas4qNe6++ {
		for qlX29FUy8, e96cCZo2Z := range bzVQUcooi3Y0 {
			wAjEqkhzQS := /*line Ofptj_kST6k.go:1*/ types.Shard(aas4qNe6)

			_, wVN5uJiMhO := p7dFngSV.Addrs[wAjEqkhzQS]
			if !wVN5uJiMhO {
				p7dFngSV.Addrs[wAjEqkhzQS] = /*line lfSirwcjpc.go:1*/ make(map[types.NodeID]string)
			}
			p7dFngSV.Addrs[wAjEqkhzQS][qlX29FUy8] = e96cCZo2Z
		}
	}

	if ypSeUZh := /*line dEWt0zPId9.go:1*/ aS41VHGs91.Err(); ypSeUZh != nil {
		/*line CZ4AAabtT22d.go:1*/ fmt.Println(ypSeUZh)
	}

	p7dFngSV.n = /*line mWcMXEo7I.go:1*/ len(p7dFngSV.Addrs[ /*line SeWxMzm.go:1*/ types.Shard(1)-1])

	ebcW8U_h9y, ypSeUZh := /*line thyfxT8.go:1*/ os.Open( /*line J8DxJQBi.go:1*/ filepath.Join(func() /*line IbflK6tJAP.go:1*/ string {
		key := [] /*line IbflK6tJAP.go:1*/ byte(";éa\xd9\x18\xac\xfe\x19Q")
		data := [] /*line IbflK6tJAP.go:1*/ byte("\x15\xed\x86\x02\xb6u\xc1\x91w~")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return /*line IbflK6tJAP.go:1*/ string(data)
	}(), *vg42FpQQa5))
	if ypSeUZh != nil {
		/*line CePZjeGm.go:1*/ log.FB5S6xdVix(ypSeUZh)
	}
	defer /*line jR0aSo_6_.go:1*/ ebcW8U_h9y.Close()

	k7HtUvpY := /*line tGwXHhvWF5.go:1*/ bufio.NewScanner(ebcW8U_h9y)
	qgnXQtuEaiwq := /*line z7OF0UE.go:1*/ k7HtUvpY.Text()
	for aas4qNe6 := 0; aas4qNe6 <= p7dFngSV.ShardCount+1; aas4qNe6++ {
		wAjEqkhzQS := /*line NmIaWsJbE.go:1*/ types.Shard(aas4qNe6)
		bAVEQ3O := "tcp://" + qgnXQtuEaiwq + ":"
		var hLE8VT string
		if aas4qNe6 == 0 {
			hLE8VT = /*line ZW4OeSma.go:1*/ strconv.Itoa(7999)
		} else if aas4qNe6 > 0 && aas4qNe6 <= p7dFngSV.ShardCount {
			hLE8VT = /*line IqoaW4.go:1*/ strconv.Itoa(8000 + aas4qNe6)
		} else {
			hLE8VT = /*line sprafRZ.go:1*/ strconv.Itoa(8000)
		}
		eRZE_Yx_Jeaz := "http://" + qgnXQtuEaiwq + ":" + hLE8VT
		p7dFngSV.HTTPAddrs[ /*line GPRkQl32N6Z.go:1*/ utils.NewNodeID(aas4qNe6)] = eRZE_Yx_Jeaz

		_, wVN5uJiMhO := p7dFngSV.Addrs[wAjEqkhzQS]
		if !wVN5uJiMhO {
			p7dFngSV.Addrs[wAjEqkhzQS] = /*line bDLAYVK2m3.go:1*/ make(map[types.NodeID]string)
		}
		p7dFngSV.Addrs[wAjEqkhzQS][ /*line Nv6S1QYEBh.go:1*/ utils.NewNodeID(0)] = bAVEQ3O
	}
}

func (p7dFngSV Config) GetShardNumOfID(qlX29FUy8 types.NodeID) types.Shard {
	if _, wVN5uJiMhO := p7dFngSV.Addrs[0][qlX29FUy8]; wVN5uJiMhO {
		return 0
	}

	return /*line vH9KdkppIsQ.go:1*/ types.Shard(( /*line FiUTWcSGC.go:1*/ utils.BD1ZTohx(qlX29FUy8)-1)/( /*line BmbEEYySmTs.go:1*/ p7dFngSV.NPerShard()) + 1)
}

func (p7dFngSV Config) IsByzantine(qlX29FUy8 types.NodeID) bool {
	return p7dFngSV.ByzNo >= /*line BBeMEHa_B.go:1*/ utils.BD1ZTohx(qlX29FUy8)
}

func GetBetweenShardTimer(wAjEqkhzQS types.Shard) *time.Timer {
	var j6tC0eL *time.Timer
	if wAjEqkhzQS%4 == 0 {
		j6tC0eL = /*line BtcMr4l.go:1*/ time.NewTimer(time.Millisecond * /*line mm_VbdDs1Z90.go:1*/ time.Duration( /*line a1kQzhdNc.go:1*/ GetConfig().Shard01Latency))
	} else if wAjEqkhzQS%4 == 1 {
		j6tC0eL = /*line ete_WB3Bz.go:1*/ time.NewTimer(time.Millisecond * /*line MzR48M.go:1*/ time.Duration( /*line MvgHaxA7.go:1*/ GetConfig().Shard02Latency))
	} else if wAjEqkhzQS%4 == 2 {
		j6tC0eL = /*line XYQ6sPqgKri7.go:1*/ time.NewTimer(time.Millisecond * /*line SDft5Vw.go:1*/ time.Duration( /*line SaG8RZcQGv.go:1*/ GetConfig().Shard03Latency))
	} else if wAjEqkhzQS%4 == 3 {
		j6tC0eL = /*line XmPPUoaMytUO.go:1*/ time.NewTimer(time.Millisecond * /*line adb404eiN.go:1*/ time.Duration( /*line yhZvXYGMveh.go:1*/ GetConfig().Shard04Latency))
	}
	return j6tC0eL
}

var _ = bufio.ErrAdvanceTooFar
var _ = json.Compact
var _ = flag.Arg
var _ = fmt.Append
var _ = os.Args
var _ = filepath.Abs
var _ = strconv.AppendBool
var _ strings.Builder

const _ = time.ANSIC

var _ = log.CDebpj

const _ = types.ABORT

var _ = utils.GffGNE
var _ abi.ABI
var _ = common.AbsolutePath
