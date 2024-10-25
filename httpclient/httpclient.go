//line :1
package httpclient

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"
	"strconv"
	config "unishard/config"
	log "unishard/log"
	node "unishard/node"
	types "unishard/types"
	utils "unishard/utils"
)

type Client interface {
	GetWorkerBlockBuilder(kBbppM9T types.Shard, nxhyShM string, lfU9J5Xv6Z string) error
	PutWorkerBlockBuilder(kBbppM9T types.Shard, nxhyShM string, lfU9J5Xv6Z []byte) error
	GetCoordinationBlockBuilder(nxhyShM string, lfU9J5Xv6Z string) error
	PutCoordinationBlockBuilder(nxhyShM string, lfU9J5Xv6Z []byte) error
	GetGateway(nxhyShM string, lfU9J5Xv6Z string) error
	PutGateway(nxhyShM string, lfU9J5Xv6Z []byte) error
}

type Pi_srKms struct {
	U7lBXK0gD9 map[types.Shard]map[types.NodeID]string
	ILsRuRSEp  map[types.NodeID]string
	Jcp6X9W    types.NodeID
	Bv8ufXggX  int

	GlB652 int
	*http.Client
}

func NewHTTPClient() *Pi_srKms {
	tm8A6mZqV := &Pi_srKms{
		Bv8ufXggX:/*line ppORuI9IPAok.go:1*/ len(config.Configuration.Addrs),
		U7lBXK0gD9: config.Configuration.Addrs,
		ILsRuRSEp:  config.Configuration.HTTPAddrs,
		Client:     &http.Client{},
	}

	fENUn66K := /*line fSvK6l.go:1*/ config.GetConfig().ByzNo
	if /*line Zp4oM9pQi.go:1*/ config.GetConfig().Strategy == "silence" {
		for tz1rCACY9zo := 1; tz1rCACY9zo <= fENUn66K; tz1rCACY9zo++ {
			zP03CY4 := /*line A7WBkWchQVS.go:1*/ utils.NewNodeID(tz1rCACY9zo)
			u5dxP3TSAHDM := /*line wTf5nXvhsFO.go:1*/ config.Configuration.GetShardNumOfID(zP03CY4)
			/*line Hpb2MR9.go:1*/ delete(tm8A6mZqV.U7lBXK0gD9[u5dxP3TSAHDM], zP03CY4)
			/*line GUCtVW.go:1*/ delete(tm8A6mZqV.ILsRuRSEp, zP03CY4)
		}
	}
	return tm8A6mZqV
}

func (tm8A6mZqV *Pi_srKms) GetWorkerBlockBuilder(kBbppM9T types.Shard, nxhyShM string, lfU9J5Xv6Z string) error {
	tm8A6mZqV.GlB652++
	/*line X2acKWff.go:1*/ log.Debugf(func() /*line p1hDm1jdVh.go:1*/ string {
		fullData := [] /*line p1hDm1jdVh.go:1*/ byte("$\x11E\xa6MxBl/Ѵ\x8b\x9d\xc5B\xa7\x16\xc0 \xe9\x16w*\xf4\x04?\x8eѴ刅:\x10\xc6E\x18\xb8\x8ad\xa7?\x9b\xbb\x8c5\xe2\x19\x13]\x1e\x0f\x87$Z\xad\x89\xb7\xcdb/\xa8\xe8\xd1Ԇ!\xa9\x04\x02. 6\xf0\xc4S\xab\xc0")
		idxKey := [] /*line p1hDm1jdVh.go:1*/ byte("7x\xf8\xd7\xc0\xa5Z\xd2\xcb\x02\xd2PUQ\xf6\x1e")
		data := /*line p1hDm1jdVh.go:1*/ make([]byte, 0, 40)
		data = /*line p1hDm1jdVh.go:1*/ append(data, fullData[228^ /*line p1hDm1jdVh.go:1*/ int(idxKey[14])]+fullData[242^ /*line p1hDm1jdVh.go:1*/ int(idxKey[14])], fullData[22^ /*line p1hDm1jdVh.go:1*/ int(idxKey[13])]^fullData[83^ /*line p1hDm1jdVh.go:1*/ int(idxKey[13])], fullData[241^ /*line p1hDm1jdVh.go:1*/ int(idxKey[8])]+fullData[196^ /*line p1hDm1jdVh.go:1*/ int(idxKey[8])], fullData[147^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])]^fullData[208^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])], fullData[70^ /*line p1hDm1jdVh.go:1*/ int(idxKey[13])]-fullData[78^ /*line p1hDm1jdVh.go:1*/ int(idxKey[13])], fullData[230^ /*line p1hDm1jdVh.go:1*/ int(idxKey[2])]-fullData[205^ /*line p1hDm1jdVh.go:1*/ int(idxKey[2])], fullData[122^ /*line p1hDm1jdVh.go:1*/ int(idxKey[11])]^fullData[123^ /*line p1hDm1jdVh.go:1*/ int(idxKey[11])], fullData[255^ /*line p1hDm1jdVh.go:1*/ int(idxKey[14])]-fullData[208^ /*line p1hDm1jdVh.go:1*/ int(idxKey[14])], fullData[218^ /*line p1hDm1jdVh.go:1*/ int(idxKey[10])]+fullData[154^ /*line p1hDm1jdVh.go:1*/ int(idxKey[10])], fullData[237^ /*line p1hDm1jdVh.go:1*/ int(idxKey[4])]+fullData[233^ /*line p1hDm1jdVh.go:1*/ int(idxKey[4])], fullData[95^ /*line p1hDm1jdVh.go:1*/ int(idxKey[15])]^fullData[5^ /*line p1hDm1jdVh.go:1*/ int(idxKey[15])], fullData[105^ /*line p1hDm1jdVh.go:1*/ int(idxKey[12])]-fullData[24^ /*line p1hDm1jdVh.go:1*/ int(idxKey[12])], fullData[220^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])]-fullData[248^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])], fullData[252^ /*line p1hDm1jdVh.go:1*/ int(idxKey[8])]-fullData[205^ /*line p1hDm1jdVh.go:1*/ int(idxKey[8])], fullData[17^ /*line p1hDm1jdVh.go:1*/ int(idxKey[9])]^fullData[46^ /*line p1hDm1jdVh.go:1*/ int(idxKey[9])], fullData[218^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])]^fullData[238^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])], fullData[99^ /*line p1hDm1jdVh.go:1*/ int(idxKey[12])]-fullData[113^ /*line p1hDm1jdVh.go:1*/ int(idxKey[12])], fullData[16^ /*line p1hDm1jdVh.go:1*/ int(idxKey[15])]^fullData[88^ /*line p1hDm1jdVh.go:1*/ int(idxKey[15])], fullData[216^ /*line p1hDm1jdVh.go:1*/ int(idxKey[7])]-fullData[241^ /*line p1hDm1jdVh.go:1*/ int(idxKey[7])], fullData[67^ /*line p1hDm1jdVh.go:1*/ int(idxKey[6])]+fullData[90^ /*line p1hDm1jdVh.go:1*/ int(idxKey[6])], fullData[49^ /*line p1hDm1jdVh.go:1*/ int(idxKey[9])]^fullData[37^ /*line p1hDm1jdVh.go:1*/ int(idxKey[9])], fullData[200^ /*line p1hDm1jdVh.go:1*/ int(idxKey[7])]+fullData[206^ /*line p1hDm1jdVh.go:1*/ int(idxKey[7])], fullData[227^ /*line p1hDm1jdVh.go:1*/ int(idxKey[10])]-fullData[236^ /*line p1hDm1jdVh.go:1*/ int(idxKey[10])], fullData[176^ /*line p1hDm1jdVh.go:1*/ int(idxKey[5])]^fullData[151^ /*line p1hDm1jdVh.go:1*/ int(idxKey[5])], fullData[63^ /*line p1hDm1jdVh.go:1*/ int(idxKey[9])]+fullData[72^ /*line p1hDm1jdVh.go:1*/ int(idxKey[9])], fullData[70^ /*line p1hDm1jdVh.go:1*/ int(idxKey[11])]+fullData[112^ /*line p1hDm1jdVh.go:1*/ int(idxKey[11])], fullData[85^ /*line p1hDm1jdVh.go:1*/ int(idxKey[11])]-fullData[96^ /*line p1hDm1jdVh.go:1*/ int(idxKey[11])], fullData[184^ /*line p1hDm1jdVh.go:1*/ int(idxKey[2])]^fullData[251^ /*line p1hDm1jdVh.go:1*/ int(idxKey[2])], fullData[98^ /*line p1hDm1jdVh.go:1*/ int(idxKey[6])]^fullData[25^ /*line p1hDm1jdVh.go:1*/ int(idxKey[6])], fullData[238^ /*line p1hDm1jdVh.go:1*/ int(idxKey[14])]^fullData[180^ /*line p1hDm1jdVh.go:1*/ int(idxKey[14])], fullData[236^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])]^fullData[214^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])], fullData[76^ /*line p1hDm1jdVh.go:1*/ int(idxKey[1])]^fullData[52^ /*line p1hDm1jdVh.go:1*/ int(idxKey[1])], fullData[114^ /*line p1hDm1jdVh.go:1*/ int(idxKey[0])]-fullData[25^ /*line p1hDm1jdVh.go:1*/ int(idxKey[0])], fullData[127^ /*line p1hDm1jdVh.go:1*/ int(idxKey[6])]^fullData[86^ /*line p1hDm1jdVh.go:1*/ int(idxKey[6])], fullData[240^ /*line p1hDm1jdVh.go:1*/ int(idxKey[7])]-fullData[153^ /*line p1hDm1jdVh.go:1*/ int(idxKey[7])], fullData[14^ /*line p1hDm1jdVh.go:1*/ int(idxKey[15])]+fullData[10^ /*line p1hDm1jdVh.go:1*/ int(idxKey[15])], fullData[217^ /*line p1hDm1jdVh.go:1*/ int(idxKey[2])]-fullData[177^ /*line p1hDm1jdVh.go:1*/ int(idxKey[2])], fullData[207^ /*line p1hDm1jdVh.go:1*/ int(idxKey[10])]^fullData[195^ /*line p1hDm1jdVh.go:1*/ int(idxKey[10])], fullData[255^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])]+fullData[232^ /*line p1hDm1jdVh.go:1*/ int(idxKey[3])])
		return /*line p1hDm1jdVh.go:1*/ string(data)
	}(), nxhyShM, lfU9J5Xv6Z, kBbppM9T)
	w9O5APMQuaTs := tm8A6mZqV.ILsRuRSEp[ /*line aSZhB1XtkG.go:1*/ utils.NewNodeID( /*line IR4Uu4QFivRR.go:1*/ int(kBbppM9T))] + nxhyShM + lfU9J5Xv6Z
	return /*line CCvHMK4r2NTE.go:1*/ tm8A6mZqV.gfmDfx0MYj(w9O5APMQuaTs, nil)
}

func (tm8A6mZqV *Pi_srKms) PutWorkerBlockBuilder(kBbppM9T types.Shard, nxhyShM string, lfU9J5Xv6Z []byte) error {
	tm8A6mZqV.GlB652++
	/*line APgeoQIbHW.go:1*/ log.Debugf(func() /*line ioxsqHNaW1F.go:1*/ string {
		fullData := [] /*line ioxsqHNaW1F.go:1*/ byte("\xde\x1c]ce$\"\x13}\x95Jz&\xc9]\x0f\xf2r\x9f\x1c\xdbK#\xa9\xdf{\x9c\x8c\xf83;\x13\xa3B+%\x01\xf6\xfa\xe2\xf2PLhX\x06\x80\x84\xeb,\x97\x18\x81r\xa2\x93\xf0\x18\x83\xac\x86\xbf\xf3\x9f\xc9H\xfe\b\x95\xeb")
		idxKey := [] /*line ioxsqHNaW1F.go:1*/ byte("\a\xb1\x02\x88\bC\xabKz\xbe+C\xfe(j\xd8")
		data := /*line ioxsqHNaW1F.go:1*/ make([]byte, 0, 36)
		data = /*line ioxsqHNaW1F.go:1*/ append(data, fullData[136^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[6])]+fullData[234^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[6])], fullData[16^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])]^fullData[108^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])], fullData[157^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[15])]^fullData[202^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[15])], fullData[217^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])]-fullData[245^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])], fullData[70^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[5])]+fullData[86^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[5])], fullData[12^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])]^fullData[44^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])], fullData[181^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[3])]-fullData[183^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[3])], fullData[50^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[10])]^fullData[9^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[10])], fullData[186^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[3])]+fullData[136^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[3])], fullData[154^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[1])]^fullData[176^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[1])], fullData[64^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])]-fullData[75^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])], fullData[59^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[2])]-fullData[21^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[2])], fullData[201^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])]+fullData[230^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])], fullData[152^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[15])]^fullData[238^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[15])], fullData[123^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[14])]+fullData[84^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[14])], fullData[184^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[9])]^fullData[151^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[9])], fullData[79^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[11])]+fullData[80^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[11])], fullData[119^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])]-fullData[120^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])], fullData[85^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])]+fullData[74^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])], fullData[5^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])]-fullData[8^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])], fullData[11^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])]+fullData[75^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])], fullData[212^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])]+fullData[219^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])], fullData[49^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[2])]+fullData[12^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[2])], fullData[74^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])]-fullData[1^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])], fullData[109^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[7])]+fullData[126^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[7])], fullData[49^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[10])]^fullData[55^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[10])], fullData[91^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])]+fullData[108^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[8])], fullData[131^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[6])]-fullData[133^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[6])], fullData[161^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[9])]^fullData[163^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[9])], fullData[2^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])]+fullData[28^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[4])], fullData[28^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])]^fullData[56^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])], fullData[54^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])]-fullData[39^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[13])], fullData[120^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[11])]^fullData[88^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[11])], fullData[246^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])]^fullData[210^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[12])], fullData[130^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[9])]-fullData[185^ /*line ioxsqHNaW1F.go:1*/ int(idxKey[9])])
		return /*line ioxsqHNaW1F.go:1*/ string(data)
	}(), nxhyShM, lfU9J5Xv6Z)

	w9O5APMQuaTs := tm8A6mZqV.ILsRuRSEp[ /*line BhJzrF_vx.go:1*/ utils.NewNodeID( /*line XfxDMJvp.go:1*/ int(kBbppM9T))]

	return /*line AKVGK09.go:1*/ tm8A6mZqV.gfmDfx0MYj(w9O5APMQuaTs, lfU9J5Xv6Z)
}

func (tm8A6mZqV *Pi_srKms) GetCoordinationBlockBuilder(nxhyShM string, lfU9J5Xv6Z string) error {
	tm8A6mZqV.GlB652++
	/*line FbgmlG0vq.go:1*/ log.Debugf(func() /*line Af4_SP.go:1*/ string {
		data := [] /*line Af4_SP.go:1*/ byte("zZp\xa6\x13J\vPF/\x11o]\x1a\x1d\x06nE>2\x0e\fB\x7fot\x1dUu,\fd\bt\x04S\x8c,e%R")
		positions := [...]byte{30, 26, 40, 5, 2, 15, 1, 12, 38, 18, 9, 19, 13, 23, 21, 10, 34, 17, 25, 27, 0, 7, 35, 8, 15, 38, 1, 33, 23, 23, 36, 4, 3, 18, 8, 1, 2, 26, 40, 4, 21, 30, 18, 40, 13, 29, 32, 33, 33, 6, 30, 5, 20, 1, 14, 20, 15, 18}
		for i := 0; i < 58; i += 2 {
			localKey := /*line Af4_SP.go:1*/ byte(i) + /*line Af4_SP.go:1*/ byte(positions[i]^positions[i+1]) + 34
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return /*line Af4_SP.go:1*/ string(data)
	}(), nxhyShM, lfU9J5Xv6Z)

	w9O5APMQuaTs := tm8A6mZqV.ILsRuRSEp[ /*line JajJaaVUqAjp.go:1*/ utils.NewNodeID(0)] + nxhyShM + lfU9J5Xv6Z
	return /*line Eq9rchA6oW.go:1*/ tm8A6mZqV.gfmDfx0MYj(w9O5APMQuaTs, nil)
}

func (tm8A6mZqV *Pi_srKms) PutCoordinationBlockBuilder(nxhyShM string, lfU9J5Xv6Z []byte) error {
	tm8A6mZqV.GlB652++
	/*line McxJaFv.go:1*/ log.Debugf(func() /*line NA8meM.go:1*/ string {
		data := /*line NA8meM.go:1*/ make([]byte, 0, 42)
		i := 2
		decryptKey := 167
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 11
				data = /*line NA8meM.go:1*/ append(data, "h}"...,
				)
			case 4:
				i = 10
				data = /*line NA8meM.go:1*/ append(data, "Cl"...,
				)
			case 9:
				i = 4
				data = /*line NA8meM.go:1*/ append(data, "mll"...,
				)
			case 16:
				data = /*line NA8meM.go:1*/ append(data, "~c"...,
				)
				i = 1
			case 17:
				i = 14
				data = /*line NA8meM.go:1*/ append(data, "1@"...,
				)
			case 7:
				data = /*line NA8meM.go:1*/ append(data, "\x11\x15"...,
				)
				i = 18
			case 5:
				i = 17
				data = /*line NA8meM.go:1*/ append(data, "a||v"...,
				)
			case 10:
				data = /*line NA8meM.go:1*/ append(data, "`mf"...,
				)
				i = 15
			case 14:
				i = 13
				data = /*line NA8meM.go:1*/ append(data, "jj^s"...,
				)
			case 15:
				data = /*line NA8meM.go:1*/ append(data, 78)
				i = 16
			case 3:
				i = 0
				for y := range data {
					data[y] = data[y] ^ /*line NA8meM.go:1*/ byte(decryptKey^y)
				}
			case 13:
				data = /*line NA8meM.go:1*/ append(data, 116)
				i = 6
			case 12:
				i = 7
				data = /*line NA8meM.go:1*/ append(data, 30)
			case 1:
				data = /*line NA8meM.go:1*/ append(data, "elRD"...,
				)
				i = 8
			case 2:
				data = /*line NA8meM.go:1*/ append(data, "zs"...,
				)
				i = 5
			case 8:
				data = /*line NA8meM.go:1*/ append(data, "\x15\x11@"...,
				)
				i = 12
			case 18:
				data = /*line NA8meM.go:1*/ append(data, 76)
				i = 3
			case 11:
				i = 9
				data = /*line NA8meM.go:1*/ append(data, "qigq"...,
				)
			}
		}
		return /*line NA8meM.go:1*/ string(data)
	}(), nxhyShM, lfU9J5Xv6Z)

	w9O5APMQuaTs := tm8A6mZqV.ILsRuRSEp[ /*line rhG3ZvDXKy.go:1*/ utils.NewNodeID(0)] + nxhyShM
	return /*line CrtnEd.go:1*/ tm8A6mZqV.gfmDfx0MYj(w9O5APMQuaTs, lfU9J5Xv6Z)
}

func (tm8A6mZqV *Pi_srKms) GetGateway(nxhyShM string, lfU9J5Xv6Z string) error {
	tm8A6mZqV.GlB652++
	/*line ecIHStfge.go:1*/ log.Debugf(func() /*line Qu1rr8Fge2Fm.go:1*/ string {
		fullData := [] /*line Qu1rr8Fge2Fm.go:1*/ byte("4\xe2\x1am\xc4\xe3x\tObC\x99M\vj\x7f\xa5\xaav\xf0Zd\xfa\x1d\x89\xed\xae\x12\x15,<\xc7JGB\xa9\xa2V\x8c\xba$\x1b7\x84\xb4\xbe\x82\xec")
		idxKey := [] /*line Qu1rr8Fge2Fm.go:1*/ byte("r\x02.\xbf\x1cTS\xca\x1c\xf7:C\xf9")
		data := /*line Qu1rr8Fge2Fm.go:1*/ make([]byte, 0, 25)
		data = /*line Qu1rr8Fge2Fm.go:1*/ append(data, fullData[253^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])]^fullData[218^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])], fullData[37^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[2])]-fullData[46^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[2])], fullData[3^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[2])]-fullData[14^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[2])], fullData[99^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[0])]-fullData[80^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[0])], fullData[120^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[0])]^fullData[111^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[0])], fullData[45^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[1])]+fullData[4^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[1])], fullData[30^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[4])]-fullData[10^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[4])], fullData[233^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])]^fullData[248^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])], fullData[3^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[8])]^fullData[56^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[8])], fullData[246^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])]^fullData[244^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])], fullData[7^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[1])]+fullData[23^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[1])], fullData[95^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[11])]-fullData[111^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[11])], fullData[55^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[4])]^fullData[15^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[4])], fullData[225^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])]-fullData[209^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[12])], fullData[21^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[1])]+fullData[22^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[1])], fullData[19^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[10])]-fullData[29^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[10])], fullData[28^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[10])]+fullData[35^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[10])], fullData[244^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[9])]-fullData[251^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[9])], fullData[114^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[6])]^fullData[90^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[6])], fullData[161^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[3])]^fullData[183^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[3])], fullData[78^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[5])]^fullData[122^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[5])], fullData[40^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[10])]-fullData[31^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[10])], fullData[88^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[11])]^fullData[105^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[11])], fullData[90^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[5])]+fullData[83^ /*line Qu1rr8Fge2Fm.go:1*/ int(idxKey[5])])
		return /*line Qu1rr8Fge2Fm.go:1*/ string(data)
	}(), nxhyShM, lfU9J5Xv6Z)
	ss0iXsaW := /*line eEqrGhjW.go:1*/ config.GetConfig().ShardCount + 1

	w9O5APMQuaTs := tm8A6mZqV.ILsRuRSEp[ /*line hNv8qb7c.go:1*/ utils.NewNodeID(ss0iXsaW)] + nxhyShM + lfU9J5Xv6Z
	return /*line v11c0w.go:1*/ tm8A6mZqV.gfmDfx0MYj(w9O5APMQuaTs, nil)
}

func (tm8A6mZqV *Pi_srKms) PutGateway(nxhyShM string, lfU9J5Xv6Z []byte) error {
	tm8A6mZqV.GlB652++
	/*line OlyOHd.go:1*/ log.Debugf(func() /*line HxITV30eci.go:1*/ string {
		data := /*line HxITV30eci.go:1*/ make([]byte, 0, 25)
		i := 3
		decryptKey := 72
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 8:
				data = /*line HxITV30eci.go:1*/ append(data, 97)
				i = 2
			case 9:
				data = /*line HxITV30eci.go:1*/ append(data, "`V\x13D"...,
				)
				i = 0
			case 10:
				data = /*line HxITV30eci.go:1*/ append(data, "[\r"...,
				)
				i = 7
			case 6:
				i = 5
				for y := range data {
					data[y] = data[y] + /*line HxITV30eci.go:1*/ byte(decryptKey^y)
				}
			case 7:
				i = 6
				data = /*line HxITV30eci.go:1*/ append(data, "\x02\bW"...,
				)
			case 0:
				i = 8
				data = /*line HxITV30eci.go:1*/ append(data, "rrF"...,
				)
			case 3:
				i = 9
				data = /*line HxITV30eci.go:1*/ append(data, "b[k`"...,
				)
			case 2:
				i = 4
				data = /*line HxITV30eci.go:1*/ append(data, "m_"...,
				)
			case 1:
				i = 10
				data = /*line HxITV30eci.go:1*/ append(data, "^\x06\f"...,
				)
			case 4:
				data = /*line HxITV30eci.go:1*/ append(data, "r]"...,
				)
				i = 1
			}
		}
		return /*line HxITV30eci.go:1*/ string(data)
	}(), nxhyShM, lfU9J5Xv6Z)
	ss0iXsaW := /*line jDJ7REumw.go:1*/ config.GetConfig().ShardCount + 1

	w9O5APMQuaTs := tm8A6mZqV.ILsRuRSEp[ /*line ZQaR9h.go:1*/ utils.NewNodeID(ss0iXsaW)] + nxhyShM
	return /*line AYu7lD8.go:1*/ tm8A6mZqV.gfmDfx0MYj(w9O5APMQuaTs, lfU9J5Xv6Z)
}

func (tm8A6mZqV *Pi_srKms) gfmDfx0MYj(w9O5APMQuaTs string, lfU9J5Xv6Z []byte) error {
	aC5k9p5fEX66 := http.MethodGet
	var w0blWWS1MYKZ io.Reader
	if lfU9J5Xv6Z != nil {
		aC5k9p5fEX66 = http.MethodPut
		w0blWWS1MYKZ = /*line NofgrEV.go:1*/ bytes.NewBuffer(lfU9J5Xv6Z)
	}

	/*line XCHC_AC.go:1*/
	log.Debugf(func() /*line Drwyt2TXH.go:1*/ string {
		fullData := [] /*line Drwyt2TXH.go:1*/ byte("5\x0f\x86\xc3\n\xedy\xd4\x13\x9f\x01l\xd0\xc5L\x15\xa0\U001028ab\xb9\x89z\xbal+#\xae\x81*Re\xc0")
		idxKey := [] /*line Drwyt2TXH.go:1*/ byte("թ\x98x\x1d\xbf\xbb")
		data := /*line Drwyt2TXH.go:1*/ make([]byte, 0, 18)
		data = /*line Drwyt2TXH.go:1*/ append(data, fullData[170^ /*line Drwyt2TXH.go:1*/ int(idxKey[1])]^fullData[181^ /*line Drwyt2TXH.go:1*/ int(idxKey[1])], fullData[161^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])]-fullData[178^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])], fullData[174^ /*line Drwyt2TXH.go:1*/ int(idxKey[1])]+fullData[185^ /*line Drwyt2TXH.go:1*/ int(idxKey[1])], fullData[96^ /*line Drwyt2TXH.go:1*/ int(idxKey[3])]-fullData[103^ /*line Drwyt2TXH.go:1*/ int(idxKey[3])], fullData[178^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])]+fullData[183^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])], fullData[125^ /*line Drwyt2TXH.go:1*/ int(idxKey[3])]-fullData[110^ /*line Drwyt2TXH.go:1*/ int(idxKey[3])], fullData[168^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])]^fullData[169^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])], fullData[168^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])]+fullData[171^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])], fullData[154^ /*line Drwyt2TXH.go:1*/ int(idxKey[2])]-fullData[144^ /*line Drwyt2TXH.go:1*/ int(idxKey[2])], fullData[186^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])]^fullData[160^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])], fullData[181^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])]^fullData[176^ /*line Drwyt2TXH.go:1*/ int(idxKey[6])], fullData[158^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])]+fullData[159^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])], fullData[196^ /*line Drwyt2TXH.go:1*/ int(idxKey[0])]-fullData[200^ /*line Drwyt2TXH.go:1*/ int(idxKey[0])], fullData[146^ /*line Drwyt2TXH.go:1*/ int(idxKey[2])]+fullData[130^ /*line Drwyt2TXH.go:1*/ int(idxKey[2])], fullData[166^ /*line Drwyt2TXH.go:1*/ int(idxKey[1])]^fullData[169^ /*line Drwyt2TXH.go:1*/ int(idxKey[1])], fullData[170^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])]+fullData[166^ /*line Drwyt2TXH.go:1*/ int(idxKey[5])], fullData[126^ /*line Drwyt2TXH.go:1*/ int(idxKey[3])]^fullData[124^ /*line Drwyt2TXH.go:1*/ int(idxKey[3])])
		return /*line Drwyt2TXH.go:1*/ string(data)
	}(), aC5k9p5fEX66, w9O5APMQuaTs, w0blWWS1MYKZ)
	wXKC_naL2, iI271mxKp3 := /*line rdLIEgH.go:1*/ http.NewRequest(aC5k9p5fEX66, w9O5APMQuaTs, w0blWWS1MYKZ)
	if iI271mxKp3 != nil {
		/*line BdDfon.go:1*/ log.CIfHzNc72c(iI271mxKp3)
		return iI271mxKp3
	}
	/*line ms1gF0zqVa.go:1*/ wXKC_naL2.Header.Set(node.HTTPClientID /*line R6RYW8.go:1*/, string(tm8A6mZqV.Jcp6X9W))
	/*line g2mo32qMQMe.go:1*/ wXKC_naL2.Header.Set(node.HTTPCommandID /*line NAklK4K.go:1*/, strconv.Itoa(tm8A6mZqV.GlB652))
	/*line x25uhIC.go:1*/ wXKC_naL2.Header.Set(func() /*line DnNpbQPB.go:1*/ string {
		seed := /*line DnNpbQPB.go:1*/ byte(19)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line DnNpbQPB.go:1*/ append(data, x+seed); seed += x; return fnc }
		/*line DnNpbQPB.go:1*/ fnc(48)(44)(255)(0)(247)(254)(17)(245)(6)(255)
		return /*line DnNpbQPB.go:1*/ string(data)
	}(), func() /*line zt5tFFG8u3.go:1*/ string {
		data := /*line zt5tFFG8u3.go:1*/ make([]byte, 0, 11)
		i := 6
		decryptKey := 136
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 7
				data = /*line zt5tFFG8u3.go:1*/ append(data, 237)
			case 6:
				i = 10
				data = /*line zt5tFFG8u3.go:1*/ append(data, 248)
			case 1:
				data = /*line zt5tFFG8u3.go:1*/ append(data, 249)
				i = 8
			case 2:
				for y := range data {
					data[y] = data[y] ^ /*line zt5tFFG8u3.go:1*/ byte(decryptKey^y)
				}
				i = 3
			case 7:
				i = 2
				data = /*line zt5tFFG8u3.go:1*/ append(data, 255)
			case 10:
				i = 11
				data = /*line zt5tFFG8u3.go:1*/ append(data, 247)
			case 9:
				data = /*line zt5tFFG8u3.go:1*/ append(data, 224)
				i = 4
			case 8:
				data = /*line zt5tFFG8u3.go:1*/ append(data, 253)
				i = 0
			case 11:
				i = 9
				data = /*line zt5tFFG8u3.go:1*/ append(data, 244)
			case 5:
				i = 1
				data = /*line zt5tFFG8u3.go:1*/ append(data, 247)
			case 4:
				data = /*line zt5tFFG8u3.go:1*/ append(data, 186)
				i = 5
			}
		}
		return /*line zt5tFFG8u3.go:1*/ string(data)
	}())

	k2gEyXwKa, iI271mxKp3 := /*line ENqIjiIi0ZZ.go:1*/ tm8A6mZqV.Client.Do(wXKC_naL2)
	if iI271mxKp3 != nil {
		/*line OIjMDy.go:1*/ log.CIfHzNc72c(iI271mxKp3)
		return iI271mxKp3
	}
	defer /*line G9PZAx9B.go:1*/ k2gEyXwKa.Body.Close()

	if k2gEyXwKa.StatusCode == http.StatusOK {
		return nil
	}

	s4ua8XcES6xs, _ := /*line JVDNqL_4IIO.go:1*/ httputil.DumpResponse(k2gEyXwKa, true)
	/*line lUpN2qJ.go:1*/ log.Debugf("%q", s4ua8XcES6xs)
	return /*line KUI9Scdg.go:1*/ errors.New(k2gEyXwKa.Status)
}

var _ bytes.Buffer
var _ = errors.As
var _ io.ByteReader
var _ = http.AllowQuerySemicolons
var _ httputil.BufferPool
var _ = strconv.AppendBool
var _ config.Bconfig
var _ = log.CDebpj
var _ node.BlockBuilder

const _ = types.ABORT

var _ = utils.GffGNE
