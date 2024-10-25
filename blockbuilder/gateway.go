//line :1
package blockbuilder

import (
	"encoding/gob"
	"encoding/hex"
	"math"
	"os/exec"
	"sync"
	"time"
	blockchain "unishard/blockchain"
	config "unishard/config"
	httpclient "unishard/httpclient"
	log "unishard/log"
	message "unishard/message"
	node "unishard/node"
	transport "unishard/transport"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
)

type (
	U7hE62tSaXOO struct {
		node.YxdKRgsOCr
		httpclient.Client
		PnCpLXZlF0c  chan interface{}
		MlMWcTSzrzrb chan interface{}
		Jt_iPwHpmgAu chan interface{}
		cXjl8zw1iHNU transport.Transport
		fFwHeHft9B   map[types.Shard]string
		k0ccaa       map[types.Shard]transport.Transport
		zLUALENr     map[types.Shard][]string
		a_15l7BeIkK  []types.Shard
		jBFiCJzhj2g  map[common.Address]map[string]types.J_zoDMw
		hBYik1       map[types.Shard][]common.Hash
		RCNsNIbcO19  map[common.Hash]int64
		vgv5ZC9oy5J  map[types.Shard]bool
		wEFflhGA3Xdi map[types.Shard]message.R8SUVOg
		wnpf69wXu0   [][]int64
		xacoxmJX     [][]int64
		c_uXYSaILdm  map[types.Shard]message.Experiment
		uIr9qU       sync.Once
	}
)

func X9uPNUll(hnc1t9s_zr string, shBiykZ types.Shard, gJp1UIg5cl string) *U7hE62tSaXOO {
	auFiAvL := /*line QdtiSTVrVM.go:1*/ new(U7hE62tSaXOO)
	auFiAvL.Client = /*line sMey7U.go:1*/ httpclient.NewHTTPClient()
	auFiAvL.YxdKRgsOCr = /*line F41pgqEKiL.go:1*/ node.I2vD85cy(hnc1t9s_zr, shBiykZ)
	auFiAvL.PnCpLXZlF0c = /*line vEKYOW0.go:1*/ make(chan interface{}, 128)
	auFiAvL.MlMWcTSzrzrb = /*line RajDib.go:1*/ make(chan interface{}, 128)
	auFiAvL.Jt_iPwHpmgAu = /*line S0Rqn7wBBmV.go:1*/ make(chan interface{}, 128)
	auFiAvL.cXjl8zw1iHNU = /*line LAVB2JzZ8O3B.go:1*/ transport.NewTransport(gJp1UIg5cl)
	auFiAvL.fFwHeHft9B = /*line u3XqSxNPQe.go:1*/ make(map[types.Shard]string)
	auFiAvL.k0ccaa = /*line B0PcVJ6ftgUG.go:1*/ make(map[types.Shard]transport.Transport)
	auFiAvL.zLUALENr = /*line mZ7EiZNQqUn.go:1*/ make(map[types.Shard][]string)
	auFiAvL.a_15l7BeIkK = /*line H2GpMWnfn.go:1*/ make([]types.Shard, 0)
	auFiAvL.jBFiCJzhj2g = /*line HYHfM44qm.go:1*/ make(map[common.Address]map[string]types.J_zoDMw)
	auFiAvL.hBYik1 = /*line ZZ8J8UBfR.go:1*/ make(map[types.Shard][]common.Hash)
	auFiAvL.RCNsNIbcO19 = /*line C5pKzm4.go:1*/ make(map[common.Hash]int64)
	auFiAvL.vgv5ZC9oy5J = /*line ye6nBpR4vw.go:1*/ make(map[types.Shard]bool)
	auFiAvL.wEFflhGA3Xdi = /*line dhAs0aJ83EKj.go:1*/ make(map[types.Shard]message.R8SUVOg)
	auFiAvL.wnpf69wXu0 = /*line oeNZ87HH.go:1*/ make([][]int64, 0)
	auFiAvL.xacoxmJX = /*line nNawKN0Vu.go:1*/ make([][]int64, 0)
	auFiAvL.c_uXYSaILdm = /*line D_x2X2.go:1*/ make(map[types.Shard]message.Experiment)

	/*line Gsa1L0K0wR.go:1*/
	gob.Register(message.WorkerBuilderListRequest{})
	/*line t526QXIQIDh.go:1*/ gob.Register(message.WorkerSharedVariableRegisterResponse{})
	/*line JCks2Fyf.go:1*/ gob.Register(message.ClientStart{})
	/*line E32RRc_mc.go:1*/ gob.Register(message.TransactionForm{})
	/*line bKKa2vt9Q.go:1*/ gob.Register(message.Transaction{})
	/*line S3B2eI.go:1*/ gob.Register(blockchain.WorkerBlock{})
	/*line HtrtEI52.go:1*/ gob.Register(message.Experiment{})

	/*line sia1fw_YMn.go:1*/
	auFiAvL.Register(message.WorkerBuilderListResponse{}, auFiAvL.tXowSV)
	/*line eGzrUT02a.go:1*/ auFiAvL.Register(message.WorkerSharedVariableRegisterResponse{}, auFiAvL.l9QXWb)
	/*line dcGIUW11N1lo.go:1*/ auFiAvL.Register(message.ClientStart{}, auFiAvL.ahB5LLc)
	/*line Goqxrpn.go:1*/ auFiAvL.Register(message.TransactionForm{}, auFiAvL.uKkNnNM)
	/*line mbtiaY.go:1*/ auFiAvL.Register(blockchain.WorkerBlock{}, auFiAvL.uNmBCPm)
	/*line RZbBaO.go:1*/ auFiAvL.Register(message.Experiment{}, auFiAvL.xDUa91)

	return auFiAvL
}

func (auFiAvL *U7hE62tSaXOO) Start() {
	go /*line JiEPdNF.go:1*/ auFiAvL.YxdKRgsOCr.Run()

	aeeXFooet := /*line AjZVufOZWtt.go:1*/ utils.Retry(auFiAvL.cXjl8zw1iHNU.Dial, 100 /*line qz2kO8bg.go:1*/, time.Duration(50)*time.Millisecond)
	if aeeXFooet != nil {
		/*line BCCjhBn4B.go:1*/ panic(aeeXFooet)
	}

	hPNhTZcPo0S3, t3r6JZyHQ5r0, iKYcKbojvkG, yopKlt8 := /*line L7cwXkn.go:1*/ config.LuxobPjs()

	a3YCxZRO := /*line nNME9b.go:1*/ config.UVxlf_QrYL("train")
	t0ra44pxfG := /*line TQCTxxe.go:1*/ config.UVxlf_QrYL("hotel")
	yjdahIB := /*line mcZPs9X0flua.go:1*/ config.UVxlf_QrYL("travel")
	lf5D2Iia8 := /*line byQbhr6h.go:1*/ config.UVxlf_QrYL("payment")

	for _, aehfYPWIx := range t3r6JZyHQ5r0 {
		auFiAvL.jBFiCJzhj2g[aehfYPWIx[0]] = a3YCxZRO
	}
	for _, aehfYPWIx := range iKYcKbojvkG {
		auFiAvL.jBFiCJzhj2g[aehfYPWIx[0]] = t0ra44pxfG
	}
	for _, aehfYPWIx := range yopKlt8 {
		auFiAvL.jBFiCJzhj2g[aehfYPWIx[0]] = yjdahIB
	}

	for _, aehfYPWIx := range hPNhTZcPo0S3 {
		auFiAvL.jBFiCJzhj2g[aehfYPWIx] = lf5D2Iia8
	}

	if aeeXFooet = /*line Ja4TEsBleeUI.go:1*/ auFiAvL.Client.GetCoordinationBlockBuilder(func() /*line GakbZNY.go:1*/ string {
		data := [] /*line GakbZNY.go:1*/ byte("\xbbW\x85\x12k\xc7\x1bMu\xd1lde\x9c\xb9\xacV\xb2Oe\xb3\xfbɒ\v?")
		positions := [...]byte{9, 2, 7, 16, 13, 18, 7, 24, 6, 3, 22, 24, 9, 17, 14, 23, 13, 5, 21, 0, 17, 3, 16, 0, 3, 7, 6, 20, 15, 5, 16, 0}
		for i := 0; i < 32; i += 2 {
			localKey := /*line GakbZNY.go:1*/ byte(i) + /*line GakbZNY.go:1*/ byte(positions[i]^positions[i+1]) + 147
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return /*line GakbZNY.go:1*/ string(data)
	}(), func() /*line aGwJRq5.go:1*/ string {
		seed := /*line aGwJRq5.go:1*/ byte(46)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data = /*line aGwJRq5.go:1*/ append(data, x-seed); seed += x; return fnc }
		/*line aGwJRq5.go:1*/ fnc(117)(4)(27)(39)(96)(170)(108)(156)
		return /*line aGwJRq5.go:1*/ string(data)
	}()+ /*line gznj8DgloG.go:1*/ auFiAvL.GetIP()); aeeXFooet != nil {
		/*line kJwiSOPeH.go:1*/ panic(aeeXFooet)
	}

	go func() {
		for {
			/*line Tl8LaZO1vsR.go:1*/ oAwuXpW := 0
			jyRmXZsIEZ8, m2Ez8Zpq, pNJa85Mre, zO1lqLXb, mBBVfAk, e5rOvB7U, eNMhTxn5Jan, eaVDIXktM, lgNekXA7Ya := 0, 0, 0, 0, 0, 0, 0, 0, 0
			for {
				select {
				case qzJdt_ := <-auFiAvL.Jt_iPwHpmgAu:
					switch cp9xTLm4SBR8 := (qzJdt_).(type) {
					case message.WorkerBuilderListResponse:
						auFiAvL.fFwHeHft9B = cp9xTLm4SBR8.Builders

						go func( /*line dTv6who.go:1*/ zH4Jtr map[types.Shard]string) {
							for shBiykZ, q7A5DFQ_Oe := range zH4Jtr {
								if _, gvacP_ := auFiAvL.k0ccaa[shBiykZ]; !gvacP_ {
									tR01ENnYu := /*line i7P4RBh.go:1*/ transport.NewTransport(q7A5DFQ_Oe)

									if aeeXFooet := /*line _xgHXpP.go:1*/ tR01ENnYu.Dial(); aeeXFooet != nil {
										/*line hYnLvDsW8_k.go:1*/ panic(aeeXFooet)
									}
									auFiAvL.k0ccaa[shBiykZ] = tR01ENnYu
								}

								if aeeXFooet = /*line yNt14bYcOTYg.go:1*/ auFiAvL.Client.GetWorkerBlockBuilder(shBiykZ, func() /*line IFcn25PQZXnj.go:1*/ string {
									data := /*line IFcn25PQZXnj.go:1*/ make([]byte, 0, 38)
									i := 0
									decryptKey := 116
									for counter := 0; i != 12; counter++ {
										decryptKey ^= i * counter
										switch i {
										case 1:
											i = 9
											data = /*line IFcn25PQZXnj.go:1*/ append(data, 162)
										case 6:
											i = 8
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "nr{o"...,
											)
										case 4:
											i = 14
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "\x85~\x8e"...,
											)
										case 8:
											i = 13
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "[qr"...,
											)
										case 10:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "tSi"...,
											)
											i = 2
										case 14:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "n~"...,
											)
											i = 3
										case 5:
											i = 12
											for y := range data {
												data[y] = data[y] + /*line IFcn25PQZXnj.go:1*/ byte(decryptKey^y)
											}
										case 3:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "v\x8a|v"...,
											)
											i = 7
										case 2:
											i = 1
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "t\xb3"...,
											)
										case 7:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "gu\x85w"...,
											)
											i = 6
										case 13:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "ox|l"...,
											)
											i = 10
										case 11:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "\xb3y"...,
											)
											i = 5
										case 9:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, 179)
											i = 11
										case 0:
											data = /*line IFcn25PQZXnj.go:1*/ append(data, "Mt\x8f\x91"...,
											)
											i = 4
										}
									}
									return /*line IFcn25PQZXnj.go:1*/ string(data)
								}(), func() /*line D2xCSXjC.go:1*/ string {
									key := [] /*line D2xCSXjC.go:1*/ byte("E\xb9Inj$\xaf\x87")
									data := [] /*line D2xCSXjC.go:1*/ byte("\x02\xd8=\v\x1dEֺ")
									for i, b := range key {
										data[i] = data[i] ^ b
									}
									return /*line D2xCSXjC.go:1*/ string(data)
								}()+ /*line h4QlAtlpK.go:1*/ auFiAvL.GetIP()); aeeXFooet != nil {
									/*line CA_B_VjnTjj.go:1*/ panic(aeeXFooet)
								}

							}
						}(auFiAvL.fFwHeHft9B)

					case message.WorkerSharedVariableRegisterResponse:
						auFiAvL.zLUALENr[cp9xTLm4SBR8.PI77Zg94] = /*line LRIJKJpBo0_U.go:1*/ append(
							auFiAvL.zLUALENr[cp9xTLm4SBR8.PI77Zg94],
							cp9xTLm4SBR8.QxT5nG3X5Gn4,
						)

					case message.ClientStart:
						auFiAvL.a_15l7BeIkK = /*line wbBRUZFpSP.go:1*/ append(auFiAvL.a_15l7BeIkK, cp9xTLm4SBR8.Shard)

						if /*line DeqBiCzWBn.go:1*/ len(auFiAvL.a_15l7BeIkK) == /*line DomNshjaraF.go:1*/ config.GetConfig().ShardCount+1 {
							hQ8l_Dkg := /*line BkZxlvumK2.go:1*/ exec.Command("sh", "-c", func() /*line Rwe5MxcXu.go:1*/ string {
								seed := /*line Rwe5MxcXu.go:1*/ byte(209)
								var data []byte
								type decFunc func(byte) decFunc
								var fnc decFunc
								fnc = func(x byte) decFunc { data = /*line Rwe5MxcXu.go:1*/ append(data, x^seed); seed += x; return fnc }
								/*line Rwe5MxcXu.go:1*/ fnc(255)(255)(172)(23)(251)(232)(27)(228)(84)(238)
								return /*line Rwe5MxcXu.go:1*/ string(data)
							}())
							if aeeXFooet := /*line M6z7Juc7G.go:1*/ hQ8l_Dkg.Run(); aeeXFooet != nil {

							} else {

							}
						}
					case message.TransactionForm:
						switch {
						case /*line dxEXNir6B.go:1*/ len(cp9xTLm4SBR8.Data) > 0:
							if /*line GHggnaA.go:1*/ cp9xTLm4SBR8.To.String() == func() /*line tSgyfXiK.go:1*/ string {
								fullData := [] /*line tSgyfXiK.go:1*/ byte("5\x84M\xef\x8aG\x97\x1duP\x05\xe4r\xf6\x1f\x9a\xa65 \xacػ\xd1F-\xc3\xea\xcff\xb4\x96\xff\xd86[v\xd7F\xb5M8:\xe0\x9d\xaa-勅H\x97\bw\xe8C]\r\x14\x99\xb4\xf0\xad}\x8a\x9cX\x99$\xcf\x01\xa6\xfba\x10=\xc0m\xe8m\x13\xe3\xa6B\xa7")
								idxKey := [] /*line tSgyfXiK.go:1*/ byte("\x1d0\xd8cCR&\n\xcb\vzr\x18\xba\x13\x1c")
								data := /*line tSgyfXiK.go:1*/ make([]byte, 0, 43)
								data = /*line tSgyfXiK.go:1*/ append(data, fullData[20^ /*line tSgyfXiK.go:1*/ int(idxKey[1])]-fullData[99^ /*line tSgyfXiK.go:1*/ int(idxKey[1])], fullData[30^ /*line tSgyfXiK.go:1*/ int(idxKey[1])]-fullData[126^ /*line tSgyfXiK.go:1*/ int(idxKey[1])], fullData[244^ /*line tSgyfXiK.go:1*/ int(idxKey[2])]^fullData[215^ /*line tSgyfXiK.go:1*/ int(idxKey[2])], fullData[100^ /*line tSgyfXiK.go:1*/ int(idxKey[10])]^fullData[106^ /*line tSgyfXiK.go:1*/ int(idxKey[10])], fullData[30^ /*line tSgyfXiK.go:1*/ int(idxKey[7])]^fullData[63^ /*line tSgyfXiK.go:1*/ int(idxKey[7])], fullData[84^ /*line tSgyfXiK.go:1*/ int(idxKey[11])]-fullData[66^ /*line tSgyfXiK.go:1*/ int(idxKey[11])], fullData[26^ /*line tSgyfXiK.go:1*/ int(idxKey[14])]+fullData[57^ /*line tSgyfXiK.go:1*/ int(idxKey[14])], fullData[81^ /*line tSgyfXiK.go:1*/ int(idxKey[0])]+fullData[4^ /*line tSgyfXiK.go:1*/ int(idxKey[0])], fullData[51^ /*line tSgyfXiK.go:1*/ int(idxKey[9])]^fullData[65^ /*line tSgyfXiK.go:1*/ int(idxKey[9])], fullData[97^ /*line tSgyfXiK.go:1*/ int(idxKey[1])]+fullData[15^ /*line tSgyfXiK.go:1*/ int(idxKey[1])], fullData[21^ /*line tSgyfXiK.go:1*/ int(idxKey[0])]+fullData[8^ /*line tSgyfXiK.go:1*/ int(idxKey[0])], fullData[51^ /*line tSgyfXiK.go:1*/ int(idxKey[11])]+fullData[82^ /*line tSgyfXiK.go:1*/ int(idxKey[11])], fullData[99^ /*line tSgyfXiK.go:1*/ int(idxKey[6])]-fullData[48^ /*line tSgyfXiK.go:1*/ int(idxKey[6])], fullData[252^ /*line tSgyfXiK.go:1*/ int(idxKey[13])]+fullData[190^ /*line tSgyfXiK.go:1*/ int(idxKey[13])], fullData[115^ /*line tSgyfXiK.go:1*/ int(idxKey[1])]^fullData[9^ /*line tSgyfXiK.go:1*/ int(idxKey[1])], fullData[243^ /*line tSgyfXiK.go:1*/ int(idxKey[13])]^fullData[168^ /*line tSgyfXiK.go:1*/ int(idxKey[13])], fullData[22^ /*line tSgyfXiK.go:1*/ int(idxKey[12])]-fullData[27^ /*line tSgyfXiK.go:1*/ int(idxKey[12])], fullData[67^ /*line tSgyfXiK.go:1*/ int(idxKey[11])]+fullData[63^ /*line tSgyfXiK.go:1*/ int(idxKey[11])], fullData[2^ /*line tSgyfXiK.go:1*/ int(idxKey[14])]-fullData[25^ /*line tSgyfXiK.go:1*/ int(idxKey[14])], fullData[50^ /*line tSgyfXiK.go:1*/ int(idxKey[10])]+fullData[97^ /*line tSgyfXiK.go:1*/ int(idxKey[10])], fullData[100^ /*line tSgyfXiK.go:1*/ int(idxKey[5])]-fullData[29^ /*line tSgyfXiK.go:1*/ int(idxKey[5])], fullData[24^ /*line tSgyfXiK.go:1*/ int(idxKey[9])]^fullData[75^ /*line tSgyfXiK.go:1*/ int(idxKey[9])], fullData[246^ /*line tSgyfXiK.go:1*/ int(idxKey[8])]^fullData[224^ /*line tSgyfXiK.go:1*/ int(idxKey[8])], fullData[239^ /*line tSgyfXiK.go:1*/ int(idxKey[2])]-fullData[245^ /*line tSgyfXiK.go:1*/ int(idxKey[2])], fullData[109^ /*line tSgyfXiK.go:1*/ int(idxKey[6])]^fullData[26^ /*line tSgyfXiK.go:1*/ int(idxKey[6])], fullData[90^ /*line tSgyfXiK.go:1*/ int(idxKey[12])]+fullData[42^ /*line tSgyfXiK.go:1*/ int(idxKey[12])], fullData[20^ /*line tSgyfXiK.go:1*/ int(idxKey[12])]-fullData[74^ /*line tSgyfXiK.go:1*/ int(idxKey[12])], fullData[94^ /*line tSgyfXiK.go:1*/ int(idxKey[4])]-fullData[66^ /*line tSgyfXiK.go:1*/ int(idxKey[4])], fullData[88^ /*line tSgyfXiK.go:1*/ int(idxKey[15])]^fullData[3^ /*line tSgyfXiK.go:1*/ int(idxKey[15])], fullData[17^ /*line tSgyfXiK.go:1*/ int(idxKey[9])]+fullData[46^ /*line tSgyfXiK.go:1*/ int(idxKey[9])], fullData[15^ /*line tSgyfXiK.go:1*/ int(idxKey[14])]-fullData[50^ /*line tSgyfXiK.go:1*/ int(idxKey[14])], fullData[49^ /*line tSgyfXiK.go:1*/ int(idxKey[9])]+fullData[13^ /*line tSgyfXiK.go:1*/ int(idxKey[9])], fullData[93^ /*line tSgyfXiK.go:1*/ int(idxKey[11])]-fullData[80^ /*line tSgyfXiK.go:1*/ int(idxKey[11])], fullData[85^ /*line tSgyfXiK.go:1*/ int(idxKey[11])]^fullData[76^ /*line tSgyfXiK.go:1*/ int(idxKey[11])], fullData[140^ /*line tSgyfXiK.go:1*/ int(idxKey[8])]+fullData[203^ /*line tSgyfXiK.go:1*/ int(idxKey[8])], fullData[119^ /*line tSgyfXiK.go:1*/ int(idxKey[4])]^fullData[70^ /*line tSgyfXiK.go:1*/ int(idxKey[4])], fullData[121^ /*line tSgyfXiK.go:1*/ int(idxKey[11])]-fullData[73^ /*line tSgyfXiK.go:1*/ int(idxKey[11])], fullData[248^ /*line tSgyfXiK.go:1*/ int(idxKey[8])]^fullData[227^ /*line tSgyfXiK.go:1*/ int(idxKey[8])], fullData[5^ /*line tSgyfXiK.go:1*/ int(idxKey[6])]-fullData[49^ /*line tSgyfXiK.go:1*/ int(idxKey[6])], fullData[162^ /*line tSgyfXiK.go:1*/ int(idxKey[13])]^fullData[189^ /*line tSgyfXiK.go:1*/ int(idxKey[13])], fullData[50^ /*line tSgyfXiK.go:1*/ int(idxKey[1])]+fullData[96^ /*line tSgyfXiK.go:1*/ int(idxKey[1])], fullData[226^ /*line tSgyfXiK.go:1*/ int(idxKey[8])]+fullData[198^ /*line tSgyfXiK.go:1*/ int(idxKey[8])])
								return /*line tSgyfXiK.go:1*/ string(data)
							}() {

							} else {

								zkWXPvRf3 := /*line z1bvK5Dz5ar2.go:1*/ utils.CalculateShardToSend( /*line JUAlLt.go:1*/ append([]common.Address{cp9xTLm4SBR8.From, cp9xTLm4SBR8.To}, cp9xTLm4SBR8.ExternalAddressList...) /*line BkzMxv.go:1*/, config.GetConfig().ShardCount)
								zA3R5Feh := /*line KqBAkgAF1EY.go:1*/ len(zkWXPvRf3) > 1
								eqQU2o := /*line bnlETjN.go:1*/ hex.EncodeToString(cp9xTLm4SBR8.Data[:4])
								agLByaYfeQ := /*line mQ42EXvzC.go:1*/ make([]message.RwVariable, 0)
								fJaiXc := 0
								agLByaYfeQ = /*line BKISb2Z.go:1*/ append(agLByaYfeQ, message.RwVariable{Address: cp9xTLm4SBR8.To, Name: eqQU2o, J_zoDMw: auFiAvL.jBFiCJzhj2g[cp9xTLm4SBR8.To][eqQU2o]})

								for _, _k98zlR7LfP := range agLByaYfeQ[0].BtvFRP3tVe8k {
									reYEpyEoI, hEFdQf := /*line IHUW1mUi0M7U.go:1*/ auFiAvL.qxQi66Ua2T(fJaiXc, cp9xTLm4SBR8.ExternalAddressList, []types.HucrUMG7j4X{_k98zlR7LfP})
									fJaiXc = hEFdQf
									agLByaYfeQ = /*line kObONae8owkV.go:1*/ append(agLByaYfeQ, reYEpyEoI...)
								}
								for r4i6nnqEY := 0; r4i6nnqEY < /*line J3ZMPmY.go:1*/ len(agLByaYfeQ); r4i6nnqEY++ {
									switch agLByaYfeQ[r4i6nnqEY].Name {
									case func() /*line YIj7CFtQ.go:1*/ string {
										data := /*line YIj7CFtQ.go:1*/ make([]byte, 0, 9)
										i := 8
										decryptKey := 4
										for counter := 0; i != 5; counter++ {
											decryptKey ^= i * counter
											switch i {
											case 9:
												i = 1
												data = /*line YIj7CFtQ.go:1*/ append(data, 37)
											case 2:
												i = 3
												data = /*line YIj7CFtQ.go:1*/ append(data, 44)
											case 6:
												for y := range data {
													data[y] = data[y] + /*line YIj7CFtQ.go:1*/ byte(decryptKey^y)
												}
												i = 5
											case 3:
												data = /*line YIj7CFtQ.go:1*/ append(data, 43)
												i = 7
											case 8:
												data = /*line YIj7CFtQ.go:1*/ append(data, 89)
												i = 4
											case 7:
												data = /*line YIj7CFtQ.go:1*/ append(data, 88)
												i = 9
											case 1:
												i = 0
												data = /*line YIj7CFtQ.go:1*/ append(data, 85)
											case 0:
												i = 6
												data = /*line YIj7CFtQ.go:1*/ append(data, 84)
											case 4:
												data = /*line YIj7CFtQ.go:1*/ append(data, 44)
												i = 2
											}
										}
										return /*line YIj7CFtQ.go:1*/ string(data)
									}():

										gf0Wio_Zo_a, aAXR2A0RAXsf := /*line s43gKrD.go:1*/ auFiAvL.iKktjwV(agLByaYfeQ[r4i6nnqEY].J_zoDMw.ReadSet, agLByaYfeQ[r4i6nnqEY].J_zoDMw.WriteSet, cp9xTLm4SBR8.MappingIdx, "0")
										agLByaYfeQ[r4i6nnqEY].J_zoDMw.ReadSet = gf0Wio_Zo_a
										agLByaYfeQ[r4i6nnqEY].J_zoDMw.WriteSet = aAXR2A0RAXsf
									case func() /*line ph1iBCah5f.go:1*/ string {
										seed := /*line ph1iBCah5f.go:1*/ byte(183)
										var data []byte
										type decFunc func(byte) decFunc
										var fnc decFunc
										fnc = func(x byte) decFunc { data = /*line ph1iBCah5f.go:1*/ append(data, x^seed); seed += x; return fnc }
										/*line ph1iBCah5f.go:1*/ fnc(135)(13)(120)(241)(213)(186)(113)(128)
										return /*line ph1iBCah5f.go:1*/ string(data)
									}():

										gf0Wio_Zo_a, aAXR2A0RAXsf := /*line z2UB_G6.go:1*/ auFiAvL.iKktjwV(agLByaYfeQ[r4i6nnqEY].J_zoDMw.ReadSet, agLByaYfeQ[r4i6nnqEY].J_zoDMw.WriteSet, cp9xTLm4SBR8.MappingIdx, "0")
										agLByaYfeQ[r4i6nnqEY].J_zoDMw.ReadSet = gf0Wio_Zo_a
										agLByaYfeQ[r4i6nnqEY].J_zoDMw.WriteSet = aAXR2A0RAXsf
									case func() /*line ijVikL.go:1*/ string {
										data := [] /*line ijVikL.go:1*/ byte("52\x94\xf7j\x9aא")
										positions := [...]byte{4, 3, 4, 7, 7, 2, 3, 5, 4, 7, 6, 3}
										for i := 0; i < 12; i += 2 {
											localKey := /*line ijVikL.go:1*/ byte(i) + /*line ijVikL.go:1*/ byte(positions[i]^positions[i+1]) + 143
											data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
										}
										return /*line ijVikL.go:1*/ string(data)
									}():

										wauOukwDG6 := /*line adk5iIYkK46.go:1*/ utils.GffGNE(cp9xTLm4SBR8.ExternalAddressList[r4i6nnqEY-2], 0)
										gf0Wio_Zo_a, aAXR2A0RAXsf := /*line tUe7DaM_.go:1*/ auFiAvL.iKktjwV(agLByaYfeQ[r4i6nnqEY].J_zoDMw.ReadSet, agLByaYfeQ[r4i6nnqEY].J_zoDMw.WriteSet, cp9xTLm4SBR8.MappingIdx, "0")
										aAXR2A0RAXsf = /*line Ra2kQo.go:1*/ append(aAXR2A0RAXsf /*line yDN15pXX.go:1*/, hex.EncodeToString(wauOukwDG6))
										agLByaYfeQ[r4i6nnqEY].J_zoDMw.ReadSet = gf0Wio_Zo_a
										agLByaYfeQ[r4i6nnqEY].J_zoDMw.WriteSet = aAXR2A0RAXsf
									}
								}

								t1uYudwjr := /*line msXuNBg4u.go:1*/ message.AtmQ9HD(cp9xTLm4SBR8, agLByaYfeQ, zA3R5Feh /*line fEvtrv75.go:1*/, utils.NewNodeID(0))
								oAwuXpW++
								if /*line B4uCLlFbWv9.go:1*/ len(zkWXPvRf3) == 1 {
									switch zkWXPvRf3[0] {
									case 1:
										jyRmXZsIEZ8++
									case 2:
										pNJa85Mre++
									case 3:
										mBBVfAk++
									case 4:
										e5rOvB7U++
									}
								} else if /*line J1_86TItoH0.go:1*/ len(zkWXPvRf3) > 1 {
									lgNekXA7Ya++
									for _, shBiykZ := range zkWXPvRf3 {
										if shBiykZ == 1 {
											m2Ez8Zpq++
										}
										if shBiykZ == 2 {
											zO1lqLXb++
										}
										if shBiykZ == 3 {
											eNMhTxn5Jan++
										}
										if shBiykZ == 4 {
											eaVDIXktM++
										}
									}
								}
								for _, shBiykZ := range zkWXPvRf3 {
									if zA3R5Feh {

									} else {

									}
									auFiAvL.hBYik1[shBiykZ] = /*line q4dDhJG7i8s.go:1*/ append(auFiAvL.hBYik1[shBiykZ], t1uYudwjr.Hash)
									/*line cD65EeauSSwk.go:1*/ auFiAvL.k0ccaa[shBiykZ].Send(t1uYudwjr)
								}
							}
						default:

							zkWXPvRf3 := /*line WXQ1lnduAN.go:1*/ utils.CalculateShardToSend([]common.Address{cp9xTLm4SBR8.From, cp9xTLm4SBR8.To} /*line pUKlA45.go:1*/, config.GetConfig().ShardCount)
							zA3R5Feh := /*line zQT1GqcEyx.go:1*/ len(zkWXPvRf3) > 1
							t1uYudwjr := /*line WvlLGV.go:1*/ message.UjchOXDlIcT(cp9xTLm4SBR8, zA3R5Feh /*line G7xl61CRo.go:1*/, utils.NewNodeID(0))
							oAwuXpW++
							if /*line I6a8k3ESfRaF.go:1*/ len(zkWXPvRf3) == 1 {
								if zkWXPvRf3[0] == 1 {
									jyRmXZsIEZ8++
								} else if zkWXPvRf3[0] == 2 {
									pNJa85Mre++
								} else if zkWXPvRf3[0] == 3 {
									mBBVfAk++
								} else if zkWXPvRf3[0] == 4 {
									e5rOvB7U++
								}
							} else if /*line APzQoeAoWp.go:1*/ len(zkWXPvRf3) == 2 {
								lgNekXA7Ya++
								for _, shBiykZ := range zkWXPvRf3 {
									if shBiykZ == 1 {
										m2Ez8Zpq++
									} else if shBiykZ == 2 {
										zO1lqLXb++
									} else if shBiykZ == 3 {
										eNMhTxn5Jan++
									} else if shBiykZ == 4 {
										eaVDIXktM++
									}
								}
							}
							for _, shBiykZ := range zkWXPvRf3 {
								if zA3R5Feh {

								} else {

								}
								auFiAvL.hBYik1[shBiykZ] = /*line _zLhV8akQ4.go:1*/ append(auFiAvL.hBYik1[shBiykZ], t1uYudwjr.Hash)
								/*line LhYqebVG.go:1*/ auFiAvL.k0ccaa[shBiykZ].Send(t1uYudwjr)
							}
						}
						if oAwuXpW%1000 == 0 {
							/*line Se6VOnB2.go:1*/ log.Eo2LyO(func() /*line FUGb_TnDU.go:1*/ string {
								key := [] /*line FUGb_TnDU.go:1*/ byte("\xb3T\x9e\x82\x86\v\f\xb4\x8b%\xee)_-c\x00i\xb4\x83!")
								data := [] /*line FUGb_TnDU.go:1*/ byte("\x96\"\xbe\xf6\xf4jb\xc7\xeaF\x9a@0C\x10 \x1a\xd1\xedU")
								for i, b := range key {
									data[i] = data[i] ^ b
								}
								return /*line FUGb_TnDU.go:1*/ string(data)
							}(), oAwuXpW)
						}
					case blockchain.WorkerBlock:
						for _, i0N37a := range cp9xTLm4SBR8.BlockData.CJpmTwa4y {
							for cBVvrBnxea, qoIxUo := range auFiAvL.hBYik1[cp9xTLm4SBR8.BlockHeader.HK3UX13] {
								if qoIxUo == i0N37a.Hash {
									auFiAvL.hBYik1[cp9xTLm4SBR8.BlockHeader.HK3UX13] = /*line JiuorxYceIzI.go:1*/ append(auFiAvL.hBYik1[cp9xTLm4SBR8.BlockHeader.HK3UX13][:cBVvrBnxea], auFiAvL.hBYik1[cp9xTLm4SBR8.BlockHeader.HK3UX13][cBVvrBnxea+1:]...)
								}
							}
						}

					case message.Experiment:
						for _, uacsCGAyKh7 := range cp9xTLm4SBR8.Bkq028NruW {
							if _, gvacP_ := auFiAvL.RCNsNIbcO19[uacsCGAyKh7.W4puDs]; !gvacP_ {
								auFiAvL.RCNsNIbcO19[uacsCGAyKh7.W4puDs] = uacsCGAyKh7.DESc6onmvvK
							} else {
								if auFiAvL.RCNsNIbcO19[uacsCGAyKh7.W4puDs] > uacsCGAyKh7.DESc6onmvvK {
									auFiAvL.RCNsNIbcO19[uacsCGAyKh7.W4puDs] = uacsCGAyKh7.DESc6onmvvK
								}
							}
						}

						_, gvacP_ := auFiAvL.wEFflhGA3Xdi[cp9xTLm4SBR8.Shard]
						if cp9xTLm4SBR8.E2v01O_Ii && !gvacP_ {
							auFiAvL.wEFflhGA3Xdi[cp9xTLm4SBR8.Shard] = cp9xTLm4SBR8.Da8uC0Sys2u

						}
						auFiAvL.c_uXYSaILdm[cp9xTLm4SBR8.Shard] = cp9xTLm4SBR8
						for _, hY1KjfvuA := range cp9xTLm4SBR8.W2g99STnJRA1 {
							auFiAvL.xacoxmJX = /*line SK_CLag3.go:1*/ append(auFiAvL.xacoxmJX, hY1KjfvuA.UCK4axEV)
						}
						for _, j5eifAKKhUk := range cp9xTLm4SBR8.JUgObcOr {
							auFiAvL.wnpf69wXu0 = /*line ruBlf21LSK_.go:1*/ append(auFiAvL.wnpf69wXu0, j5eifAKKhUk.UCK4axEV)
						}

					default:

					}
					/*line wa0InBh9V.go:1*/ auFiAvL.uIr9qU.Do(func() {
						/*line syagbwpZYt_.go:1*/ log.Eal9ypd6VP(func() /*line AOdZuaG.go:1*/ string {
							data := /*line AOdZuaG.go:1*/ make([]byte, 0, 155)
							i := 14
							decryptKey := 20
							for counter := 0; i != 42; counter++ {
								decryptKey ^= i * counter
								switch i {
								case 44:
									data = /*line AOdZuaG.go:1*/ append(data, "\x93\xb2"...,
									)
									i = 10
								case 27:
									data = /*line AOdZuaG.go:1*/ append(data, "9U"...,
									)
									i = 64
								case 65:
									i = 41
									data = /*line AOdZuaG.go:1*/ append(data, "\xcc\xcf"...,
									)
								case 22:
									data = /*line AOdZuaG.go:1*/ append(data, 148)
									i = 1
								case 48:
									data = /*line AOdZuaG.go:1*/ append(data, "\xb4\xe2\xe6\xe9"...,
									)
									i = 47
								case 28:
									i = 43
									data = /*line AOdZuaG.go:1*/ append(data, "2/"...,
									)
								case 17:
									i = 23
									data = /*line AOdZuaG.go:1*/ append(data, 114)
								case 35:
									data = /*line AOdZuaG.go:1*/ append(data, 175)
									i = 32
								case 41:
									data = /*line AOdZuaG.go:1*/ append(data, "ֶ"...,
									)
									i = 24
								case 40:
									data = /*line AOdZuaG.go:1*/ append(data, "E9E"...,
									)
									i = 38
								case 62:
									data = /*line AOdZuaG.go:1*/ append(data, "e\x8d\x8f_"...,
									)
									i = 25
								case 6:
									i = 22
									data = /*line AOdZuaG.go:1*/ append(data, 133)
								case 38:
									i = 30
									data = /*line AOdZuaG.go:1*/ append(data, "=)=H"...,
									)
								case 36:
									i = 19
									data = /*line AOdZuaG.go:1*/ append(data, "\xb9\xa0"...,
									)
								case 5:
									i = 2
									data = /*line AOdZuaG.go:1*/ append(data, 199)
								case 23:
									data = /*line AOdZuaG.go:1*/ append(data, "pte"...,
									)
									i = 52
								case 9:
									i = 56
									data = /*line AOdZuaG.go:1*/ append(data, "\t1"...,
									)
								case 19:
									i = 18
									data = /*line AOdZuaG.go:1*/ append(data, "\xa3\xa5}"...,
									)
								case 31:
									i = 29
									data = /*line AOdZuaG.go:1*/ append(data, 211)
								case 58:
									data = /*line AOdZuaG.go:1*/ append(data, "wQy{"...,
									)
									i = 37
								case 20:
									i = 26
									data = /*line AOdZuaG.go:1*/ append(data, 153)
								case 26:
									i = 45
									data = /*line AOdZuaG.go:1*/ append(data, "\x98^}\x9f"...,
									)
								case 16:
									i = 66
									data = /*line AOdZuaG.go:1*/ append(data, 50)
								case 12:
									data = /*line AOdZuaG.go:1*/ append(data, "\x9ev\x9d"...,
									)
									i = 8
								case 11:
									i = 65
									data = /*line AOdZuaG.go:1*/ append(data, "\xac\x84\xa2\xd0"...,
									)
								case 4:
									i = 3
									data = /*line AOdZuaG.go:1*/ append(data, 168)
								case 53:
									i = 39
									data = /*line AOdZuaG.go:1*/ append(data, 134)
								case 50:
									data = /*line AOdZuaG.go:1*/ append(data, "\x89\x9bR|"...,
									)
									i = 6
								case 59:
									i = 12
									data = /*line AOdZuaG.go:1*/ append(data, "\x99\x94"...,
									)
								case 51:
									data = /*line AOdZuaG.go:1*/ append(data, "\xa1w\xaa"...,
									)
									i = 4
								case 15:
									i = 28
									data = /*line AOdZuaG.go:1*/ append(data, "\x157"...,
									)
								case 45:
									data = /*line AOdZuaG.go:1*/ append(data, "\x9a\x97"...,
									)
									i = 51
								case 32:
									data = /*line AOdZuaG.go:1*/ append(data, 170)
									i = 11
								case 14:
									i = 33
									data = /*line AOdZuaG.go:1*/ append(data, "\x96\xbb"...,
									)
								case 55:
									data = /*line AOdZuaG.go:1*/ append(data, "\x96\x82\x96"...,
									)
									i = 20
								case 10:
									data = /*line AOdZuaG.go:1*/ append(data, 212)
									i = 5
								case 52:
									i = 54
									data = /*line AOdZuaG.go:1*/ append(data, 117)
								case 33:
									data = /*line AOdZuaG.go:1*/ append(data, "\xb4\xb4\xac\xb3"...,
									)
									i = 59
								case 3:
									i = 61
									data = /*line AOdZuaG.go:1*/ append(data, "\xac\x9d\xad"...,
									)
								case 37:
									i = 27
									data = /*line AOdZuaG.go:1*/ append(data, "QbQ"...,
									)
								case 57:
									data = /*line AOdZuaG.go:1*/ append(data, 139)
									i = 62
								case 21:
									data = /*line AOdZuaG.go:1*/ append(data, 121)
									i = 0
								case 63:
									data = /*line AOdZuaG.go:1*/ append(data, "\xdc\xee\xde\xe6"...,
									)
									i = 60
								case 49:
									data = /*line AOdZuaG.go:1*/ append(data, 142)
									i = 57
								case 18:
									data = /*line AOdZuaG.go:1*/ append(data, "\x9cƹ"...,
									)
									i = 13
								case 54:
									data = /*line AOdZuaG.go:1*/ append(data, "yz"...,
									)
									i = 58
								case 34:
									data = /*line AOdZuaG.go:1*/ append(data, 227)
									i = 31
								case 25:
									i = 50
									data = /*line AOdZuaG.go:1*/ append(data, "\x92\x8f\x8e"...,
									)
								case 7:
									i = 21
									data = /*line AOdZuaG.go:1*/ append(data, 136)
								case 30:
									data = /*line AOdZuaG.go:1*/ append(data, 63)
									i = 46
								case 67:
									data = /*line AOdZuaG.go:1*/ append(data, 192)
									i = 63
								case 64:
									i = 53
									data = /*line AOdZuaG.go:1*/ append(data, "fU=S"...,
									)
								case 8:
									i = 36
									data = /*line AOdZuaG.go:1*/ append(data, "\xb7ï"...,
									)
								case 66:
									i = 15
									data = /*line AOdZuaG.go:1*/ append(data, "-?\xf6"...,
									)
								case 47:
									i = 67
									data = /*line AOdZuaG.go:1*/ append(data, 232)
								case 61:
									i = 9
									data = /*line AOdZuaG.go:1*/ append(data, "\xb1\xb2\xaf"...,
									)
								case 43:
									data = /*line AOdZuaG.go:1*/ append(data, "9#4;"...,
									)
									i = 40
								case 56:
									i = 16
									data = /*line AOdZuaG.go:1*/ append(data, "3\x0363"...,
									)
								case 29:
									data = /*line AOdZuaG.go:1*/ append(data, "\xdb\xcf\xec\x9e"...,
									)
									i = 48
								case 0:
									data = /*line AOdZuaG.go:1*/ append(data, "\x89\x8d"...,
									)
									i = 49
								case 2:
									data = /*line AOdZuaG.go:1*/ append(data, "\xccֵ\xc9"...,
									)
									i = 34
								case 46:
									i = 42
									for y := range data {
										data[y] = data[y] - /*line AOdZuaG.go:1*/ byte(decryptKey^y)
									}
								case 13:
									i = 35
									data = /*line AOdZuaG.go:1*/ append(data, "\xb6\xc0"...,
									)
								case 39:
									i = 7
									data = /*line AOdZuaG.go:1*/ append(data, 132)
								case 1:
									data = /*line AOdZuaG.go:1*/ append(data, "\x9e\x92\x96"...,
									)
									i = 55
								case 60:
									i = 17
									data = /*line AOdZuaG.go:1*/ append(data, "\xe2\xf7\xa9\xbf"...,
									)
								case 24:
									data = /*line AOdZuaG.go:1*/ append(data, "\xb1\xb3"...,
									)
									i = 44
								}
							}
							return /*line AOdZuaG.go:1*/ string(data)
						}())
					})

					if oAwuXpW == /*line wSqTyH.go:1*/ config.GetConfig().Benchmark.N {
						/*line EI70lzBHmNn.go:1*/ log.Eo2LyO(func() /*line VSJxq05vLCV.go:1*/ string {
							seed := /*line VSJxq05vLCV.go:1*/ byte(231)
							var data []byte
							type decFunc func(byte) decFunc
							var fnc decFunc
							fnc = func(x byte) decFunc { data = /*line VSJxq05vLCV.go:1*/ append(data, x^seed); seed += x; return fnc }
							/*line VSJxq05vLCV.go:1*/ fnc(129)(1)(7)(25)(250)(235)(11)(29)(182)(63)(238)(23)(244)(237)(31)(247)(167)(11)(79)(168)(68)(6)(27)(251)(227)(18)(230)(31)(227)(2)(1)(3)
							return /*line VSJxq05vLCV.go:1*/ string(data)
						}(), oAwuXpW)

						yUG3GGqt := /*line sE2Wl9tz.go:1*/ time.NewTicker(time.Millisecond * /*line _zuGNK0ORhY.go:1*/ time.Duration(5000))
						<-yUG3GGqt.C
						kEtdRbt := /*line b9_xMy.go:1*/ int64(0)
						for _, j_g49l1YinM9 := range auFiAvL.RCNsNIbcO19 {
							kEtdRbt += j_g49l1YinM9
						}
						hk9u5Ijvz := /*line uCtbIiHxAN.go:1*/ math.Round(( /*line oxCQfPHFu.go:1*/ float64(kEtdRbt)/1000)/( /*line rpkguoH60.go:1*/ float64( /*line AklDf2828hYJ.go:1*/ len(auFiAvL.RCNsNIbcO19)))*1000) / 1000

						oNpWKXqzzN6w := 0.0
						cmy1UX := 0.0
						et0Tfl := 0.0
						bNlgzfMd := 0.0
						iyAwBg := 0.0
						slURaOqOBR := 0.0
						j5eifAKKhUk := []int64{0, 0}
						hY1KjfvuA := []int64{0, 0, 0, 0, 0, 0, 0, 0}

						for _, iED2lmP := range auFiAvL.wnpf69wXu0 {
							j5eifAKKhUk[0] += iED2lmP[0]
							j5eifAKKhUk[1] += iED2lmP[1]
						}

						for _, iED2lmP := range auFiAvL.xacoxmJX {
							hY1KjfvuA[0] += iED2lmP[0]
							hY1KjfvuA[1] += iED2lmP[1]
							hY1KjfvuA[2] += iED2lmP[2]
							hY1KjfvuA[3] += iED2lmP[3]
							hY1KjfvuA[4] += iED2lmP[4]
							hY1KjfvuA[5] += iED2lmP[5]
							hY1KjfvuA[6] += iED2lmP[6]
							hY1KjfvuA[7] += iED2lmP[7]
						}

						baog5S := ( /*line yzN3x7.go:1*/ float64(hY1KjfvuA[0]) / 1000) / /*line XWXKH_oR.go:1*/ float64( /*line J8JM9b.go:1*/ len(auFiAvL.xacoxmJX))
						d5maSU := ( /*line TpRNnj5QQ.go:1*/ float64(hY1KjfvuA[1]) / 1000) / /*line FPT4LC.go:1*/ float64( /*line pFFaUG5fyb.go:1*/ len(auFiAvL.xacoxmJX))
						u6qEczQ := ( /*line iRi0Ux0xP.go:1*/ float64(hY1KjfvuA[2]) / 1000) / /*line TBi75s6.go:1*/ float64( /*line QP7g4NSovjs.go:1*/ len(auFiAvL.xacoxmJX))
						vg1Mw1Pk5i := ( /*line VMncxia.go:1*/ float64(hY1KjfvuA[3]) / 1000) / /*line abdjqMzUm4Wh.go:1*/ float64( /*line BmReSSwiA.go:1*/ len(auFiAvL.xacoxmJX))
						pYMgMKH := ( /*line FTrrexy6K.go:1*/ float64(hY1KjfvuA[4]) / 1000) / /*line IeNW3mn.go:1*/ float64( /*line D_rJd3.go:1*/ len(auFiAvL.xacoxmJX))
						yjax7trrrDA := ( /*line n3JrvH.go:1*/ float64(hY1KjfvuA[5]) / 1000) / /*line nabwVtNef.go:1*/ float64( /*line uuPAQo.go:1*/ len(auFiAvL.xacoxmJX))
						z4OpaBzjrd := ( /*line qwXmHwtp.go:1*/ float64(hY1KjfvuA[6]) / 1000) / /*line fR7VQAcEcHb.go:1*/ float64( /*line Rbd0cZ1I_.go:1*/ len(auFiAvL.xacoxmJX))
						cNYQR95c := ( /*line GZ6HPha.go:1*/ float64(hY1KjfvuA[7]) / 1000) / /*line Ei8VsLvKjfd.go:1*/ float64( /*line mhoR1qKtjAhf.go:1*/ len(auFiAvL.xacoxmJX))
						akEslUGRlaXc := ( /*line E5Z_TGGr2pr.go:1*/ float64(j5eifAKKhUk[0]) / 1000) / /*line GYgaJh8p_m0.go:1*/ float64( /*line CCV35sz.go:1*/ len(auFiAvL.wnpf69wXu0))
						oSnWLaqh6h := ( /*line CNq8UIdXE94.go:1*/ float64(j5eifAKKhUk[1]) / 1000) / /*line OkgFBDqyik.go:1*/ float64( /*line uGB2gptBZS3.go:1*/ len(auFiAvL.wnpf69wXu0))

						oqnaABvEuF := d5maSU + pYMgMKH
						SReAjLEC5E := u6qEczQ + yjax7trrrDA
						f4TAXiC := cNYQR95c
						bMdf2Pvo := baog5S + vg1Mw1Pk5i + z4OpaBzjrd

						for r4i6nnqEY := 1; r4i6nnqEY <= /*line jxctC31.go:1*/ config.GetConfig().ShardCount; r4i6nnqEY++ {
							hTfNQwFOnp := auFiAvL.wEFflhGA3Xdi[ /*line AHj81aGZc0.go:1*/ types.Shard(r4i6nnqEY)]
							l56np3Vi := auFiAvL.c_uXYSaILdm[ /*line R0ymKO8WNoNc.go:1*/ types.Shard(r4i6nnqEY)]
							oNpWKXqzzN6w += /*line A4KNaHV.go:1*/ math.Round(( /*line Im8vgd.go:1*/ float64(hTfNQwFOnp.RRTdrUMYO) / hTfNQwFOnp.XRLlpce9L))
							cmy1UX += /*line RRc6Cpm.go:1*/ math.Round(( /*line yAl_w2q.go:1*/ float64(hTfNQwFOnp.RRTdrUMYO) / hTfNQwFOnp.XRLlpce9L))
							et0Tfl += /*line oykcndl_.go:1*/ math.Round(( /*line PO0OeX_f.go:1*/ float64(hTfNQwFOnp.Zjb9_abe) / hTfNQwFOnp.XRLlpce9L))
							bNlgzfMd += /*line wCBRBiM.go:1*/ math.Round(( /*line slQtmpNf.go:1*/ float64(hTfNQwFOnp.GV9IpLO9e) / hTfNQwFOnp.XRLlpce9L))
							iyAwBg += l56np3Vi.RoQbttkP
							slURaOqOBR += hk9u5Ijvz
						}
						/*line T3ewAzAMWBi.go:1*/ log.Eal9ypd6VP(func() /*line hOr9PR.go:1*/ string {
							data := /*line hOr9PR.go:1*/ make([]byte, 0, 60)
							i := 13
							decryptKey := 67
							for counter := 0; i != 3; counter++ {
								decryptKey ^= i * counter
								switch i {
								case 21:
									data = /*line hOr9PR.go:1*/ append(data, "7|v"...,
									)
									i = 5
								case 6:
									i = 24
									data = /*line hOr9PR.go:1*/ append(data, "ZPJ\x1e"...,
									)
								case 11:
									i = 15
									data = /*line hOr9PR.go:1*/ append(data, "r&og"...,
									)
								case 0:
									i = 23
									data = /*line hOr9PR.go:1*/ append(data, "\rF@"...,
									)
								case 14:
									i = 11
									data = /*line hOr9PR.go:1*/ append(data, "bh"...,
									)
								case 4:
									data = /*line hOr9PR.go:1*/ append(data, "(e"...,
									)
									i = 19
								case 12:
									data = /*line hOr9PR.go:1*/ append(data, "KG["...,
									)
									i = 0
								case 7:
									i = 12
									data = /*line hOr9PR.go:1*/ append(data, "\nC"...,
									)
								case 5:
									data = /*line hOr9PR.go:1*/ append(data, "|^"...,
									)
									i = 7
								case 18:
									data = /*line hOr9PR.go:1*/ append(data, "ey#h"...,
									)
									i = 14
								case 20:
									data = /*line hOr9PR.go:1*/ append(data, "\x00M"...,
									)
									i = 17
								case 1:
									data = /*line hOr9PR.go:1*/ append(data, "b|"...,
									)
									i = 4
								case 10:
									for y := range data {
										data[y] = data[y] ^ /*line hOr9PR.go:1*/ byte(decryptKey^y)
									}
									i = 3
								case 13:
									i = 1
									data = /*line hOr9PR.go:1*/ append(data, 104)
								case 8:
									data = /*line hOr9PR.go:1*/ append(data, "yqye"...,
									)
									i = 21
								case 2:
									data = /*line hOr9PR.go:1*/ append(data, "MQ\x1bP"...,
									)
									i = 6
								case 16:
									i = 9
									data = /*line hOr9PR.go:1*/ append(data, "9r"...,
									)
								case 17:
									i = 2
									data = /*line hOr9PR.go:1*/ append(data, 69)
								case 15:
									i = 16
									data = /*line hOr9PR.go:1*/ append(data, "so"...,
									)
								case 23:
									data = /*line hOr9PR.go:1*/ append(data, "JT"...,
									)
									i = 20
								case 24:
									data = /*line hOr9PR.go:1*/ append(data, "W_[G"...,
									)
									i = 22
								case 9:
									i = 8
									data = /*line hOr9PR.go:1*/ append(data, "|vh<"...,
									)
								case 19:
									i = 18
									data = /*line hOr9PR.go:1*/ append(data, 109)
								case 22:
									i = 10
									data = /*line hOr9PR.go:1*/ append(data, 17)
								}
							}
							return /*line hOr9PR.go:1*/ string(data)
						}(),
							oNpWKXqzzN6w,
							cmy1UX/ /*line Lqtd7PK.go:1*/ float64( /*line IAAdp0BCF.go:1*/ config.GetConfig().ShardCount),
							et0Tfl/ /*line aAjBY2t.go:1*/ float64( /*line H2c7m3cYc.go:1*/ config.GetConfig().ShardCount),
							bNlgzfMd/ /*line VwIsHPc.go:1*/ float64( /*line qRdaN3z.go:1*/ config.GetConfig().ShardCount),
							iyAwBg/ /*line YujGb3SaF_1m.go:1*/ float64( /*line cDWH31Dx.go:1*/ config.GetConfig().ShardCount),
							slURaOqOBR/ /*line ypuzyP.go:1*/ float64( /*line QqE60Emf.go:1*/ config.GetConfig().ShardCount),
							oqnaABvEuF,
							SReAjLEC5E,
							f4TAXiC,
							bMdf2Pvo,
							oSnWLaqh6h,
							akEslUGRlaXc,
						)
					}

				case <-auFiAvL.PnCpLXZlF0c:
				case <-auFiAvL.MlMWcTSzrzrb:
					continue
				}
			}
		}
	}()
}

func (auFiAvL *U7hE62tSaXOO) tXowSV(qzJdt_ message.WorkerBuilderListResponse) {

	auFiAvL.Jt_iPwHpmgAu <- qzJdt_
}

func (auFiAvL *U7hE62tSaXOO) l9QXWb(qzJdt_ message.WorkerSharedVariableRegisterResponse) {

	auFiAvL.Jt_iPwHpmgAu <- qzJdt_
}

func (auFiAvL *U7hE62tSaXOO) uKkNnNM(qzJdt_ message.TransactionForm) {
	auFiAvL.Jt_iPwHpmgAu <- qzJdt_

}
func (auFiAvL *U7hE62tSaXOO) uNmBCPm(qtrCsrxrTmq blockchain.WorkerBlock) {
	auFiAvL.Jt_iPwHpmgAu <- qtrCsrxrTmq

}

func (auFiAvL *U7hE62tSaXOO) xDUa91(qzJdt_ message.Experiment) {
	auFiAvL.Jt_iPwHpmgAu <- qzJdt_
}

func (auFiAvL *U7hE62tSaXOO) ahB5LLc(qzJdt_ message.ClientStart) {
	auFiAvL.Jt_iPwHpmgAu <- qzJdt_
}

func (auFiAvL *U7hE62tSaXOO) qxQi66Ua2T(fJaiXc int, gXlkBAtMi []common.Address, _k98zlR7LfP []types.HucrUMG7j4X) ([]message.RwVariable, int) {
	agLByaYfeQ := /*line jIxnZQ5gpc.go:1*/ make([]message.RwVariable, 0)
	hEFdQf := fJaiXc
	for _, ehz9TWYbfU := range _k98zlR7LfP {
		if ehz9TWYbfU.EDwUuQQ38 == "read" {
			agLByaYfeQ = /*line qqxl4KeP3.go:1*/ append(agLByaYfeQ, message.RwVariable{
				Address: gXlkBAtMi[fJaiXc],
				Name:    ehz9TWYbfU.OP7awm,
				J_zoDMw: types.J_zoDMw{ReadSet: []string{ehz9TWYbfU.OP7awm}, WriteSet: []string{}},
			})
			hEFdQf++
		} else if ehz9TWYbfU.EDwUuQQ38 == "execute" {
			agLByaYfeQ = /*line SuK_kSR4d.go:1*/ append(agLByaYfeQ, message.RwVariable{
				Address: gXlkBAtMi[fJaiXc],
				Name:    ehz9TWYbfU.OP7awm,
				J_zoDMw: auFiAvL.jBFiCJzhj2g[gXlkBAtMi[fJaiXc]][ehz9TWYbfU.OP7awm],
			})
			hEFdQf++
			if /*line GuLqTSzQ.go:1*/ len(auFiAvL.jBFiCJzhj2g[gXlkBAtMi[fJaiXc]][ehz9TWYbfU.OP7awm].BtvFRP3tVe8k) > 0 {
				reYEpyEoI, aJEU5NSaAK := /*line ypozVXVys.go:1*/ auFiAvL.qxQi66Ua2T(hEFdQf, gXlkBAtMi, auFiAvL.jBFiCJzhj2g[gXlkBAtMi[fJaiXc]][ehz9TWYbfU.OP7awm].BtvFRP3tVe8k)
				agLByaYfeQ = /*line awLKQtj.go:1*/ append(agLByaYfeQ, reYEpyEoI...)
				hEFdQf = aJEU5NSaAK
			}
		}
	}
	return agLByaYfeQ, hEFdQf
}

func (auFiAvL *U7hE62tSaXOO) iKktjwV(gf0Wio_Zo_a []string, aAXR2A0RAXsf []string, fJaiXc []byte, uawV02a4cIL string) ([]string, []string) {
	fZvPslFv := []string{}
	woM7akqBf := []string{}

	for _, j5aV6sXBSAG9 := range gf0Wio_Zo_a {
		if j5aV6sXBSAG9 == uawV02a4cIL {
			fZvPslFv = /*line ABbyBWpM.go:1*/ append(fZvPslFv /*line LmnYm_um.go:1*/, hex.EncodeToString(fJaiXc))
		} else {
			fZvPslFv = /*line xWUAaX.go:1*/ append(fZvPslFv, j5aV6sXBSAG9)
		}
	}
	for _, j5aV6sXBSAG9 := range aAXR2A0RAXsf {
		if j5aV6sXBSAG9 == uawV02a4cIL {
			woM7akqBf = /*line BgOefm2amJWN.go:1*/ append(woM7akqBf /*line ycOWeYHJ.go:1*/, hex.EncodeToString(fJaiXc))
		} else {
			woM7akqBf = /*line bpB7fi7Espk.go:1*/ append(woM7akqBf, j5aV6sXBSAG9)
		}
	}
	return fZvPslFv, woM7akqBf
}

var _ gob.CommonType
var _ = hex.AppendDecode
var _ = math.Abs
var _ exec.Cmd
var _ sync.Cond

const _ = time.ANSIC

var _ blockchain.Accept
var _ config.Bconfig
var _ httpclient.Client
var _ = log.CDebpj
var _ message.ClientStart
var _ node.BlockBuilder
var _ = transport.NewTransport

const _ = types.ABORT

var _ = utils.GffGNE
var _ = common.AbsolutePath
