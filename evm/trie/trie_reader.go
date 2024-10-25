//line :1
package trie

import (
	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie/triestate"
	"github.com/ethereum/go-ethereum/triedb/database"
)

type trieReader struct {
	owner  common.Hash
	reader database.Reader
	banned map[string]struct{}
}

func fHP68lzj9vD1(iiQYh2c, zalM3_NR4XT common.Hash, kRC_1aK database.Database) (*trieReader, error) {
	if iiQYh2c == (common.Hash{}) || iiQYh2c == types.NrGuaNA21 {
		if iiQYh2c == (common.Hash{}) {
			/*line avR_fB6.go:1*/ log.Error(func() /*line ibuFQRnue.go:1*/ string {
				seed := /*line ibuFQRnue.go:1*/ byte(40)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data = /*line ibuFQRnue.go:1*/ append(data, x^seed); seed += x; return fnc }
				/*line ibuFQRnue.go:1*/ fnc(114)(255)(235)(235)(79)(205)(255)(235)(1)(19)(169)(64)(29)(224)(27)(170)(92)(241)(242)(27)(175)
				return /*line ibuFQRnue.go:1*/ string(data)
			}())
		}
		return &trieReader{owner: zalM3_NR4XT}, nil
	}
	dmVOxRYtQ, eZzE0KPS := /*line p792j6i.go:1*/ kRC_1aK.Reader(iiQYh2c)
	if eZzE0KPS != nil {
		return nil, &OZajCEax3{UiAlIYHWNd6: zalM3_NR4XT, QWet5V5UyH5m: iiQYh2c, q6c1HMK_: eZzE0KPS}
	}
	return &trieReader{owner: zalM3_NR4XT, reader: dmVOxRYtQ}, nil
}

func uAe_TJrsD() *trieReader {
	return &trieReader{}
}

func (iMFHMG *trieReader) uAjz8NxU_cL(qiRG6lpeaFl []byte, rNuN0sMPDJ common.Hash) ([]byte, error) {

	if iMFHMG.banned != nil {
		if _, yY_yPSd8vG := iMFHMG.banned[ /*line iscRTAv0u.go:1*/ string(qiRG6lpeaFl)]; yY_yPSd8vG {
			return nil, &OZajCEax3{UiAlIYHWNd6: iMFHMG.owner, QWet5V5UyH5m: rNuN0sMPDJ, Sm2UO87Ro4SN: qiRG6lpeaFl}
		}
	}
	if iMFHMG.reader == nil {
		return nil, &OZajCEax3{UiAlIYHWNd6: iMFHMG.owner, QWet5V5UyH5m: rNuN0sMPDJ, Sm2UO87Ro4SN: qiRG6lpeaFl}
	}
	aBHduJG, eZzE0KPS := /*line BLnanE.go:1*/ iMFHMG.reader.Node(iMFHMG.owner, qiRG6lpeaFl, rNuN0sMPDJ)
	if eZzE0KPS != nil || /*line PRNpy5MEH.go:1*/ len(aBHduJG) == 0 {
		return nil, &OZajCEax3{UiAlIYHWNd6: iMFHMG.owner, QWet5V5UyH5m: rNuN0sMPDJ, Sm2UO87Ro4SN: qiRG6lpeaFl, q6c1HMK_: eZzE0KPS}
	}
	return aBHduJG, nil
}

type SZ5vcrrMm struct {
	eL9bOz database.Database
}

func I5ajRKaUDN(kRC_1aK database.Database) *SZ5vcrrMm {
	return &SZ5vcrrMm{eL9bOz: kRC_1aK}
}

func (aHj7_4 *SZ5vcrrMm) OpenTrie(m7S8SE common.Hash) (triestate.Trie, error) {
	return /*line N292jPqYGXCw.go:1*/ RJaQ3esB( /*line A4bho8ArTm.go:1*/ K9sqwbQ(m7S8SE), aHj7_4.eL9bOz)
}

func (aHj7_4 *SZ5vcrrMm) OpenStorageTrie(iiQYh2c common.Hash, kzkViH5, m7S8SE common.Hash) (triestate.Trie, error) {
	return /*line gL4JTXm.go:1*/ RJaQ3esB( /*line W_FdAPTY.go:1*/ NYkaq0T(iiQYh2c, kzkViH5, m7S8SE), aHj7_4.eL9bOz)
}

var _ types.AccessList
var _ = common.AbsolutePath
var _ = log.Crit
var _ = triestate.Apply
var _ database.Database
