//line :1
package pacemaker

import (
	"time"

	blockchain "unishard/blockchain"
	types "unishard/types"
)

type JjxBHN2smc interface {
	ProcessRemoteTmo(eO94A1 *TMO) (bool, *TC, *blockchain.O0GQK9U1)
	AdvanceView(hwXqUFsD types.View)
	UpdateView(hwXqUFsD types.View)
	ExecuteViewChange(f8bkQu_wY types.View)
	EnteringViewEvent() chan types.EpochView
	EnteringTmoEvent() chan types.EpochView
	UpdateAnchorView(h83Fjm types.View)
	UpdateNewView(f8bkQu_wY types.View)
	GetCurView() types.View
	GetNewView() types.View
	GetAnchorView() types.View
	GetCurRound() int
	GetCurEpoch() types.Epoch
	GetCurEpochView() (types.Epoch, types.View)
	GetTimerForView() time.Duration
	GetTimerForViewChange() time.Duration
	IsTimeToElect() bool
}

const _ = time.ANSIC

var _ blockchain.Accept

const _ = types.ABORT
