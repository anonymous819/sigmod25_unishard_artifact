//line :1
package election

import (
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Static struct {
	dfMzEjleeRU types.NW4p1ya
	frAhHE90    int
	zjFSfrMOC   types.NodeID
	k5poNJ32kfb types.Shard
}

func NewStatic(kBW2zD types.NodeID, nwfG8Xauo types.Shard, _SKyya_3 int) *Static {
	return &Static{
		zjFSfrMOC:   kBW2zD,
		k5poNJ32kfb: nwfG8Xauo,
		frAhHE90:    _SKyya_3,
		dfMzEjleeRU: []types.NodeID{},
	}
}

func (qyPmSkcq2 *Static) IsLeader(gL46kzPw7IM types.NodeID, gXBOXbK9mwD types.View, eW00eBxD5J types.Epoch) bool {
	return gL46kzPw7IM == qyPmSkcq2.zjFSfrMOC
}

func (qyPmSkcq2 *Static) FindLeaderFor(gXBOXbK9mwD types.View, eW00eBxD5J types.Epoch) types.NodeID {
	return qyPmSkcq2.zjFSfrMOC
}

func (qyPmSkcq2 *Static) IsCommittee(gL46kzPw7IM types.NodeID, eW00eBxD5J types.Epoch) bool {
	for _, A2SE5Xlf1o := range qyPmSkcq2.dfMzEjleeRU {
		if A2SE5Xlf1o == gL46kzPw7IM {
			return true
		}
	}
	return false
}

func (qyPmSkcq2 *Static) FindCommitteesFor(eW00eBxD5J types.Epoch) types.NW4p1ya {
	return qyPmSkcq2.dfMzEjleeRU
}

func (qyPmSkcq2 *Static) FindValidatorsFor(eW00eBxD5J types.Epoch) types.NW4p1ya {
	return nil
}

func (qyPmSkcq2 *Static) ElectLeader(gL46kzPw7IM types.NodeID, jTxHDS9yCnLX types.View) bool {
	return true
}

func (qyPmSkcq2 *Static) ElectCommittees(tau_Q7jMq common.Hash, pJEaQmNhFn types.Epoch) types.NW4p1ya {
	if /*line Io9TW8JA.go:1*/ len(qyPmSkcq2.dfMzEjleeRU) == 0 {
		for oHtoYLFocL := 1; oHtoYLFocL <= qyPmSkcq2.frAhHE90; oHtoYLFocL++ {
			qyPmSkcq2.dfMzEjleeRU = /*line xx497OLY.go:1*/ append(qyPmSkcq2.dfMzEjleeRU /*line X5OzLuJ.go:1*/, utils.NewNodeID(oHtoYLFocL))
		}
	}
	return qyPmSkcq2.dfMzEjleeRU
}

func (qyPmSkcq2 *Static) ReplaceCommittee(eW00eBxD5J types.Epoch, eZ4GM2DOpY types.NodeID, uv1RnabZ2 types.NodeID) types.NW4p1ya {
	return nil
}

const _ = types.ABORT

var _ = utils.GffGNE
var _ = common.AbsolutePath
