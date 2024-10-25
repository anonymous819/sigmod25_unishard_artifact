//line :1
package pacemaker

import (
	blockchain "unishard/blockchain"
	crypto "unishard/crypto"
	quorum "unishard/quorum"
	types "unishard/types"
)

type TMO struct {
	types.Shard
	types.Epoch
	QmJZRap      types.View
	Ig3RQ9F      types.View
	BcEyAIp_     types.View
	Ln9AEaadIIY  types.NodeID
	UFs51prD6H1  types.BlockHeight
	K13ceE04nX   types.BlockHeight
	N9CdZ7bx2WlQ *blockchain.WorkerBlock
	GtrfyVMK9    *blockchain.WorkerBlock
	W52craOOoraa *quorum.HPOWa1PT0xzq
}

type TC struct {
	types.Shard
	types.Epoch
	Uc9aMd   types.NodeID
	NewView  types.View
	Z7uwaNa  types.View
	U6G7XuY_ types.BlockHeight
	Z4KEom01 *blockchain.WorkerBlock
	crypto.Pp__49cd
	crypto.MqlfmVE9d
}

func CDzFORl9BQA(f8bkQu_wY types.View, x4BUVZaNK map[types.NodeID]*TMO) (*TC, *blockchain.WorkerBlock) {

	var (
		h9neLDxya    types.Shard
		h83Fjm       types.View
		ud5BhBplY_aC types.BlockHeight
		aoOYYVnm1    *blockchain.WorkerBlock
		e2ncboMpKN   *blockchain.WorkerBlock
	)

	var (
		skr0nNw93k types.View
		eUOArMIB0Q types.BlockHeight
		lGc57XzkZ  types.BlockHeight
	)

	for _, eO94A1 := range x4BUVZaNK {
		if skr0nNw93k < eO94A1.BcEyAIp_ {
			skr0nNw93k = eO94A1.BcEyAIp_
		}

		if eUOArMIB0Q < eO94A1.UFs51prD6H1 {
			eUOArMIB0Q = eO94A1.UFs51prD6H1
			aoOYYVnm1 = eO94A1.N9CdZ7bx2WlQ
		}

		if lGc57XzkZ < eO94A1.K13ceE04nX {
			lGc57XzkZ = eO94A1.K13ceE04nX
			e2ncboMpKN = eO94A1.GtrfyVMK9
		}
	}

	h83Fjm = skr0nNw93k
	ud5BhBplY_aC = eUOArMIB0Q + 1

	if lGc57XzkZ < ud5BhBplY_aC {
		e2ncboMpKN = nil
	}
	return &TC{
		Shard:    h9neLDxya,
		NewView:  f8bkQu_wY,
		Z7uwaNa:  h83Fjm,
		U6G7XuY_: ud5BhBplY_aC,
		Z4KEom01: aoOYYVnm1,
	}, e2ncboMpKN
}

var _ blockchain.Accept
var _ crypto.Pp__49cd
var _ quorum.Commit

const _ = types.ABORT
