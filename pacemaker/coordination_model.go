//line :1
package pacemaker

import (
	blockchain "unishard/blockchain"
	crypto "unishard/crypto"
	quorum "unishard/quorum"
	types "unishard/types"
)

type H8NP1AOKTF struct {
	types.Shard
	types.Epoch
	W72t5Nhk     types.View
	C50bBOZxcYfV types.View
	RR2FRcA9     types.View
	RKtUCo69O    types.NodeID
	AWApIIs      types.BlockHeight
	A2Q6Z0L      types.BlockHeight
	QSw1K7B      *blockchain.CoordinationBlock
	Jth9oL0xv    *blockchain.CoordinationBlock
	DSNq8K7x     *quorum.HPOWa1PT0xzq
}

type RZ65hTAdXj struct {
	types.Shard
	types.Epoch
	H5NbPTQ     types.NodeID
	APadJA      types.View
	C8IzV0maUK  types.View
	BecWOQ4Lltj types.BlockHeight
	HXY_lyal    *blockchain.CoordinationBlock
	crypto.Pp__49cd
	crypto.MqlfmVE9d
}

func JLIUrx3(f8bkQu_wY types.View, x4BUVZaNK map[types.NodeID]*H8NP1AOKTF) (*RZ65hTAdXj, *blockchain.CoordinationBlock) {

	var (
		h9neLDxya    types.Shard
		h83Fjm       types.View
		ud5BhBplY_aC types.BlockHeight
		aoOYYVnm1    *blockchain.CoordinationBlock
		e2ncboMpKN   *blockchain.CoordinationBlock
	)

	var (
		skr0nNw93k types.View
		eUOArMIB0Q types.BlockHeight
		lGc57XzkZ  types.BlockHeight
	)

	for _, eO94A1 := range x4BUVZaNK {
		if skr0nNw93k < eO94A1.RR2FRcA9 {
			skr0nNw93k = eO94A1.RR2FRcA9
		}

		if eUOArMIB0Q < eO94A1.AWApIIs {
			eUOArMIB0Q = eO94A1.AWApIIs
			aoOYYVnm1 = eO94A1.QSw1K7B
		}

		if lGc57XzkZ < eO94A1.A2Q6Z0L {
			lGc57XzkZ = eO94A1.A2Q6Z0L
			e2ncboMpKN = eO94A1.Jth9oL0xv
		}
	}

	h83Fjm = skr0nNw93k
	ud5BhBplY_aC = eUOArMIB0Q + 1

	if lGc57XzkZ < ud5BhBplY_aC {
		e2ncboMpKN = nil
	}
	return &RZ65hTAdXj{
		Shard:       h9neLDxya,
		APadJA:      f8bkQu_wY,
		C8IzV0maUK:  h83Fjm,
		BecWOQ4Lltj: ud5BhBplY_aC,
		HXY_lyal:    aoOYYVnm1,
	}, e2ncboMpKN
}

var _ blockchain.Accept
var _ crypto.Pp__49cd
var _ quorum.Commit

const _ = types.ABORT
