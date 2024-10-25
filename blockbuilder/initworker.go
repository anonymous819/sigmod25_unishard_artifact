//line :1
package blockbuilder

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	blockchain "unishard/blockchain"
	config "unishard/config"
	log "unishard/log"
	message "unishard/message"
	transport "unishard/transport"
	types "unishard/types"
	utils "unishard/utils"
	"unishard/worker"
)

type (
	DzS5s1o struct {
		worker.WorkerBlockBuilderReplica

		dTCAFdn7LSD  string
		fdlgYve7     transport.Transport
		ef5TQ7       map[types.NodeID]transport.Transport
		f4DeZFIiwt7  []types.NodeID
		eVc4_4JaXJmF string
	}
)

func ZRaICa(hnc1t9s_zr string, shBiykZ types.Shard, gJp1UIg5cl string, i1Z6lPiLV string) *DzS5s1o {
	fz9M8fYEK := DzS5s1o{
		WorkerBlockBuilderReplica: * /*line ow3_jY.go:1*/ worker.NewWorkerBlockBuilder(hnc1t9s_zr, shBiykZ, i1Z6lPiLV),
		dTCAFdn7LSD:               gJp1UIg5cl,
		fdlgYve7:/*line gKfg6NMr.go:1*/ transport.NewTransport(gJp1UIg5cl),
		ef5TQ7:/*line AHhzyC.go:1*/ make(map[types.NodeID]transport.Transport),
		f4DeZFIiwt7:/*line NNcPsl.go:1*/ make([]types.NodeID, 0),
		eVc4_4JaXJmF: "Alive",
	}

	return &fz9M8fYEK
}

func (fz9M8fYEK *DzS5s1o) Start() {
	go /*line UJaSElc.go:1*/ fz9M8fYEK.WorkerBlockBuilderReplica.Start()

	aeeXFooet := /*line csNVQeEKUfL.go:1*/ utils.Retry(fz9M8fYEK.fdlgYve7.Dial, 100 /*line BG8WqNH.go:1*/, time.Duration(50)*time.Millisecond)
	if aeeXFooet != nil {
		/*line _aP1uCUoxa.go:1*/ panic(aeeXFooet)
	}

	zOZWg21Ku, _ := /*line SQKpl1ad.go:1*/ json.Marshal(message.WorkerBuilderRegister{
		SenderShard:/*line JLUUvaR.go:1*/ fz9M8fYEK.GetShard(),
		Address:/*line h3P30o4sfS.go:1*/ fz9M8fYEK.GetIP(),
	})
	if aeeXFooet = /*line FfB1jsFZ5m0.go:1*/ fz9M8fYEK.Client.PutCoordinationBlockBuilder(func() /*line noBVZW4iA5G.go:1*/ string {
		key := [] /*line noBVZW4iA5G.go:1*/ byte("!-\xb6\x9b\xb6\xd3\xdbњ4\xd1\x01\x1f\x18\xfa\x9es\x93\xccBY0")
		data := [] /*line noBVZW4iA5G.go:1*/ byte("\x0e*\xb9\u05f5\x92\x97q\xdb5\x9bcFZX\xc7\xf4֧2\fB")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line noBVZW4iA5G.go:1*/ string(data)
	}(), zOZWg21Ku); aeeXFooet != nil {
		/*line iJITv9.go:1*/ panic(aeeXFooet)
	}
	/*line bb4vU9XV.go:1*/ log.Debugf(func() /*line BFsxQQLO_2i7.go:1*/ string {
		fullData := [] /*line BFsxQQLO_2i7.go:1*/ byte("\xcfH]\x82m2\xf0\xc3+}\x19I>\x1c\xbf\x02MS\x1c\vW\xa5I\x1c.\x1f\xfc\x1cM\xc5\xda \xa3+\x0e\x90.A\xf2\xa1\xdbX\xa6\xe8\xe4YS\u05cb\x8b\x86f\x90\xae\x8b\xc6l\xd1\xe1+-ѕ\x95")
		idxKey := [] /*line BFsxQQLO_2i7.go:1*/ byte("6\x90s\xd8Yu\xc9\xf9~]\x9b\xbf\\Ax\xf6")
		data := /*line BFsxQQLO_2i7.go:1*/ make([]byte, 0, 33)
		data = /*line BFsxQQLO_2i7.go:1*/ append(data, fullData[148^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[1])]^fullData[137^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[1])], fullData[195^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[3])]+fullData[206^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[3])], fullData[80^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])]^fullData[113^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])], fullData[92^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])]-fullData[111^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])], fullData[219^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[7])]^fullData[240^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[7])], fullData[97^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])]-fullData[94^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])], fullData[91^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[14])]-fullData[112^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[14])], fullData[250^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[7])]+fullData[255^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[7])], fullData[74^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[5])]+fullData[68^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[5])], fullData[252^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])]-fullData[227^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])], fullData[118^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[2])]-fullData[116^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[2])], fullData[82^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[2])]-fullData[96^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[2])], fullData[64^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[9])]^fullData[119^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[9])], fullData[107^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])]-fullData[72^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])], fullData[77^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])]+fullData[81^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])], fullData[111^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[13])]-fullData[123^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[13])], fullData[78^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[4])]+fullData[88^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[4])], fullData[129^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[10])]^fullData[165^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[10])], fullData[234^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[3])]+fullData[243^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[3])], fullData[64^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[5])]-fullData[101^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[5])], fullData[217^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])]^fullData[214^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])], fullData[163^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[11])]+fullData[173^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[11])], fullData[108^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])]+fullData[112^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[12])], fullData[64^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[14])]^fullData[119^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[14])], fullData[208^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])]+fullData[238^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])], fullData[139^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[11])]-fullData[155^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[11])], fullData[101^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[4])]^fullData[112^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[4])], fullData[74^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[13])]^fullData[94^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[13])], fullData[220^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[7])]+fullData[194^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[7])], fullData[87^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[4])]^fullData[113^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[4])], fullData[215^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[6])]+fullData[255^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[6])], fullData[209^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])]+fullData[207^ /*line BFsxQQLO_2i7.go:1*/ int(idxKey[15])])
		return /*line BFsxQQLO_2i7.go:1*/ string(data)
	}())

	/*line syxzU_NlwXb.go:1*/
	log.Debugf(func() /*line saHgLRFj.go:1*/ string {
		key := [] /*line saHgLRFj.go:1*/ byte("N\x7f$\xfeň\xbf\xfc\x85\x86\xc1\xb1w\xe6X\xdf\x03\xceSR\xb7\xfd\t\xea\x8e\xdeԝưg")
		data := [] /*line saHgLRFj.go:1*/ byte("%\xf5=t\xaf\x98\xb8s\xed\xe5\xa4\xc1\xa9|\x1d\x8ai\x96\x12 ihm{\xe0\x96Lϩ\xbf\t")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return /*line saHgLRFj.go:1*/ string(data)
	}())
	go func() {
		for {
			select {
			case /*line wCj9NVDlhDBA.go:1*/ qzJdt_ := <-fz9M8fYEK.WorkerBlockBuilderReplica.WorkerBuilderTopic:
				if cp9xTLm4SBR8, pcXDH4KDf := qzJdt_.(message.WorkerSharedVariableRegisterRequest); pcXDH4KDf {

					if aeeXFooet = /*line NZ7Mu88V6m.go:1*/ fz9M8fYEK.Client.GetGateway(func() /*line oQDNTJPoPRVD.go:1*/ string {
						key := [] /*line oQDNTJPoPRVD.go:1*/ byte("\x97\x149\xa7a\x1e\xbe\xc41t\xd6\xd5\xd6x\x0f;\xf7؏>\xd0T\xf4\xcc\x18\xcfyv\x1f\x9bP\xe0\xdbJ#\x97\x84H")
						data := [] /*line oQDNTJPoPRVD.go:1*/ byte("\x98C6\xcb\nG\xb4\x8f7휐\x8e\xdeR7r\x89\xd3.\x95\xfeq\x9bQ\xa4\xfb\xefS\xb7\x15\x93\x95%K\xdc\xe1\xf7")
						for i, b := range key {
							data[i] = data[i] + b
						}
						return /*line oQDNTJPoPRVD.go:1*/ string(data)
					}(), func() /*line IIKEMmp2eb.go:1*/ string {
						seed := /*line IIKEMmp2eb.go:1*/ byte(27)
						var data []byte
						type decFunc func(byte) decFunc
						var fnc decFunc
						fnc = func(x byte) decFunc { data = /*line IIKEMmp2eb.go:1*/ append(data, x-seed); seed += x; return fnc }
						/*line IIKEMmp2eb.go:1*/ fnc(110)(238)(229)(192)(129)(15)(255)(19)(31)(79)(144)(249)
						return /*line IIKEMmp2eb.go:1*/ string(data)
					}()+ /*line rNy2R1hs.go:1*/ strconv.Itoa( /*line _8vwGVFL9.go:1*/ int( /*line CHmVTCDMFa.go:1*/ fz9M8fYEK.GetShard()))+func() /*line kC0dNf.go:1*/ string {
						data := [] /*line kC0dNf.go:1*/ byte("4V7q\xf9a\x1f\n\x1c(=")
						positions := [...]byte{4, 3, 7, 2, 3, 2, 9, 4, 0, 9, 8, 6}
						for i := 0; i < 12; i += 2 {
							localKey := /*line kC0dNf.go:1*/ byte(i) + /*line kC0dNf.go:1*/ byte(positions[i]^positions[i+1]) + 46
							data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
						}
						return /*line kC0dNf.go:1*/ string(data)
					}()+fz9M8fYEK.eVc4_4JaXJmF); aeeXFooet != nil {
						/*line cTxjnALjmwG7.go:1*/ panic(aeeXFooet)
					}
					/*line nDIaD0m.go:1*/ log.Debugf(func() /*line RmpDFJtdr.go:1*/ string {
						seed := /*line RmpDFJtdr.go:1*/ byte(204)
						var data []byte
						type decFunc func(byte) decFunc
						var fnc decFunc
						fnc = func(x byte) decFunc { data = /*line RmpDFJtdr.go:1*/ append(data, x+seed); seed += x; return fnc }
						/*line RmpDFJtdr.go:1*/ fnc(166)(243)(254)(2)(4)(13)(239)(255)(188)(82)(243)(12)(4)(240)(14)(1)(172)(79)(247)(186)(77)(12)(167)(83)(245)(249)(17)(243)(255)(188)(86)(235)(17)(247)(248)(1)(10)(249)(187)(70)(12)(253)(254)(179)(39)(26)(19)(241)(18)(234)(24)(167)(5)(78)
						return /*line RmpDFJtdr.go:1*/ string(data)
					}(), cp9xTLm4SBR8.UpcIp_)
				} else if cp9xTLm4SBR8, pcXDH4KDf := qzJdt_.(*blockchain.WorkerConsensusPayload); pcXDH4KDf {

					for _, tR01ENnYu := range fz9M8fYEK.ef5TQ7 {
						/*line qXNI3Pe3fySD.go:1*/ tR01ENnYu.Send(cp9xTLm4SBR8)
					}
				} else if cp9xTLm4SBR8, pcXDH4KDf := qzJdt_.(message.Experiment); pcXDH4KDf {
					/*line uez2BBEOXMzV.go:1*/ fz9M8fYEK.GatewaynodeTransport.Send(cp9xTLm4SBR8)
				} else {
					/*line NrBbN73Uf11U.go:1*/ log.Debugf(func() /*line CboR3w.go:1*/ string {
						data := /*line CboR3w.go:1*/ make([]byte, 0, 26)
						i := 0
						decryptKey := 80
						for counter := 0; i != 5; counter++ {
							decryptKey ^= i * counter
							switch i {
							case 6:
								i = 1
								data = /*line CboR3w.go:1*/ append(data, "s&v"...,
								)
							case 7:
								data = /*line CboR3w.go:1*/ append(data, "\x89\x83w"...,
								)
								i = 9
							case 10:
								i = 6
								data = /*line CboR3w.go:1*/ append(data, "le"...,
								)
							case 8:
								data = /*line CboR3w.go:1*/ append(data, ";\x8c"...,
								)
								i = 3
							case 2:
								i = 7
								data = /*line CboR3w.go:1*/ append(data, ".\x85"...,
								)
							case 1:
								i = 4
								data = /*line CboR3w.go:1*/ append(data, "m~}n"...,
								)
							case 3:
								i = 5
								for y := range data {
									data[y] = data[y] - /*line CboR3w.go:1*/ byte(decryptKey^y)
								}
							case 4:
								data = /*line CboR3w.go:1*/ append(data, "st"...,
								)
								i = 2
							case 0:
								data = /*line CboR3w.go:1*/ append(data, "jlog"...,
								)
								i = 10
							case 9:
								i = 8
								data = /*line CboR3w.go:1*/ append(data, "5N7"...,
								)
							}
						}
						return /*line CboR3w.go:1*/ string(data)
					}(), /*line JY62na4dU.go:1*/ reflect.TypeOf(cp9xTLm4SBR8).String())
				}

			case qzJdt_ := <-fz9M8fYEK.WorkerBlockBuilderReplica.ConsensusNodeTopic:
				if cp9xTLm4SBR8, pcXDH4KDf := qzJdt_.(message.ConsensusNodeRegister); pcXDH4KDf {
					tR01ENnYu := /*line tSQs4f9ll9.go:1*/ transport.NewTransport(cp9xTLm4SBR8.IP)
					aeeXFooet := /*line ItKI03ymnSsa.go:1*/ utils.Retry(tR01ENnYu.Dial, 100 /*line KfGCOOrrQZY.go:1*/, time.Duration(50)*time.Millisecond)
					if aeeXFooet != nil {
						/*line tUFYiR.go:1*/ panic(aeeXFooet)
					}
					/*line bgwreE2U1h.go:1*/ log.Debugf(func() /*line B4cVsaa4aU.go:1*/ string {
						fullData := [] /*line B4cVsaa4aU.go:1*/ byte("\xecIVd\xe5%\x8d\x81?C̳\xbd\xb9w>\xfb\x1e\b\x9cy\xad\xcb\x16٪\xfec\xa9KΥX!p \xc9 \x18\x10.[\"\xdb\x14v\x00\xbb\x15خ\x8a\xf5\x0f(\xcb\rg\xf6cރ\xc4$J\x14\xef\xd9yA\xa4\x14\xb6\xacxT=W|H\xf9\x8d҈tP\xa2`\x17\xcf")
						idxKey := [] /*line B4cVsaa4aU.go:1*/ byte("M{\x839\xfb\xb0 \x89\xd3\v\x8dB)0\x05\xc1")
						data := /*line B4cVsaa4aU.go:1*/ make([]byte, 0, 46)
						data = /*line B4cVsaa4aU.go:1*/ append(data, fullData[210^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])]^fullData[233^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])], fullData[115^ /*line B4cVsaa4aU.go:1*/ int(idxKey[6])]-fullData[5^ /*line B4cVsaa4aU.go:1*/ int(idxKey[6])], fullData[191^ /*line B4cVsaa4aU.go:1*/ int(idxKey[2])]+fullData[190^ /*line B4cVsaa4aU.go:1*/ int(idxKey[2])], fullData[10^ /*line B4cVsaa4aU.go:1*/ int(idxKey[12])]-fullData[27^ /*line B4cVsaa4aU.go:1*/ int(idxKey[12])], fullData[139^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])]-fullData[134^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])], fullData[108^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])]+fullData[68^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])], fullData[125^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])]^fullData[62^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])], fullData[151^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])]-fullData[149^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])], fullData[50^ /*line B4cVsaa4aU.go:1*/ int(idxKey[3])]-fullData[117^ /*line B4cVsaa4aU.go:1*/ int(idxKey[3])], fullData[172^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])]^fullData[181^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])], fullData[43^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])]^fullData[57^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])], fullData[232^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])]+fullData[238^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])], fullData[170^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])]-fullData[250^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])], fullData[96^ /*line B4cVsaa4aU.go:1*/ int(idxKey[3])]^fullData[13^ /*line B4cVsaa4aU.go:1*/ int(idxKey[3])], fullData[66^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])]+fullData[95^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])], fullData[136^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])]^fullData[163^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])], fullData[235^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])]^fullData[253^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])], fullData[155^ /*line B4cVsaa4aU.go:1*/ int(idxKey[8])]+fullData[254^ /*line B4cVsaa4aU.go:1*/ int(idxKey[8])], fullData[255^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])]-fullData[134^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])], fullData[228^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])]^fullData[251^ /*line B4cVsaa4aU.go:1*/ int(idxKey[4])], fullData[81^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])]-fullData[41^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])], fullData[162^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])]-fullData[138^ /*line B4cVsaa4aU.go:1*/ int(idxKey[10])], fullData[132^ /*line B4cVsaa4aU.go:1*/ int(idxKey[7])]+fullData[176^ /*line B4cVsaa4aU.go:1*/ int(idxKey[7])], fullData[78^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])]^fullData[8^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])], fullData[131^ /*line B4cVsaa4aU.go:1*/ int(idxKey[7])]-fullData[139^ /*line B4cVsaa4aU.go:1*/ int(idxKey[7])], fullData[33^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])]^fullData[63^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])], fullData[69^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])]^fullData[35^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])], fullData[45^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])]^fullData[24^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])], fullData[53^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])]-fullData[75^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])], fullData[50^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])]^fullData[83^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])], fullData[250^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])]+fullData[230^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])], fullData[13^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])]-fullData[19^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])], fullData[240^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])]^fullData[205^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])], fullData[244^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])]+fullData[224^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])], fullData[41^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])]-fullData[94^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])], fullData[64^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])]^fullData[83^ /*line B4cVsaa4aU.go:1*/ int(idxKey[9])], fullData[111^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])]+fullData[65^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])], fullData[20^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])]+fullData[118^ /*line B4cVsaa4aU.go:1*/ int(idxKey[13])], fullData[221^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])]^fullData[130^ /*line B4cVsaa4aU.go:1*/ int(idxKey[15])], fullData[163^ /*line B4cVsaa4aU.go:1*/ int(idxKey[2])]+fullData[175^ /*line B4cVsaa4aU.go:1*/ int(idxKey[2])], fullData[72^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])]+fullData[80^ /*line B4cVsaa4aU.go:1*/ int(idxKey[1])], fullData[231^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])]^fullData[241^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])], fullData[230^ /*line B4cVsaa4aU.go:1*/ int(idxKey[8])]-fullData[202^ /*line B4cVsaa4aU.go:1*/ int(idxKey[8])], fullData[242^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])]^fullData[174^ /*line B4cVsaa4aU.go:1*/ int(idxKey[5])], fullData[59^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])]^fullData[1^ /*line B4cVsaa4aU.go:1*/ int(idxKey[14])])
						return /*line B4cVsaa4aU.go:1*/ string(data)
					}(), /*line k6_2aYL.go:1*/ fz9M8fYEK.GetShard(), cp9xTLm4SBR8.ConsensusNodeID, cp9xTLm4SBR8.IP)
					fz9M8fYEK.ef5TQ7[cp9xTLm4SBR8.ConsensusNodeID] = tR01ENnYu
					if /*line rpu99NadLkI.go:1*/ len(fz9M8fYEK.ef5TQ7) == /*line hrEli6Zl.go:1*/ config.GetConfig().CommitteeNumber {
						fz9M8fYEK.GoFlag = true
						/*line fJkUU1M3u40C.go:1*/ log.Debugf(func() /*line A8WIh5c2ya.go:1*/ string {
							seed := /*line A8WIh5c2ya.go:1*/ byte(188)
							var data []byte
							type decFunc func(byte) decFunc
							var fnc decFunc
							fnc = func(x byte) decFunc { data = /*line A8WIh5c2ya.go:1*/ append(data, x-seed); seed += x; return fnc }
							/*line A8WIh5c2ya.go:1*/ fnc(15)(51)(95)(207)(144)(246)(210)(169)(163)(252)(236)(11)(55)(91)(199)(144)(204)(187)(159)(59)(114)(237)(224)
							return /*line A8WIh5c2ya.go:1*/ string(data)
						}(), /*line kLq6bkiEiy69.go:1*/ fz9M8fYEK.GetShard())
						/*line yGLIWYHam.go:1*/ fz9M8fYEK.GatewaynodeTransport.Send(message.ClientStart{
							Shard: /*line xl5uXU.go:1*/ fz9M8fYEK.GetShard(),
						})
					}
				}

			case qzJdt_ := <-fz9M8fYEK.WorkerBlockBuilderReplica.CoordinationBuilderTopic:
				if cp9xTLm4SBR8, pcXDH4KDf := qzJdt_.(blockchain.WorkerBlock); pcXDH4KDf {

					yrCtuXkXuEqk := /*line KxFlad1i13.go:1*/ config.GetBetweenShardTimer( /*line xiz0ZO.go:1*/ fz9M8fYEK.GetShard())
					<-yrCtuXkXuEqk.C
					for _, aM5zas_eDo := range cp9xTLm4SBR8.BlockData.ReceivedCrossTransaction {
						aM5zas_eDo.LatencyDissection.WorkerConsensusTime = /*line Cqw1sukb.go:1*/ time.Now().UnixMilli()
					}
					/*line b39hHz.go:1*/ fz9M8fYEK.fdlgYve7.Send(cp9xTLm4SBR8)
					/*line _UtzN1q.go:1*/ fz9M8fYEK.GatewaynodeTransport.Send(cp9xTLm4SBR8)
				}

			case <-fz9M8fYEK.GatewayTopic:
				continue
			}
		}
	}()
}

var _ = json.Compact
var _ = reflect.Append
var _ = strconv.AppendBool

const _ = time.ANSIC

var _ blockchain.Accept
var _ config.Bconfig
var _ = log.CDebpj
var _ message.ClientStart
var _ = transport.NewTransport

const _ = types.ABORT

var _ = utils.GffGNE
var _ = worker.NewReplica
