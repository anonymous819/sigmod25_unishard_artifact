//line :1
package blockchain

import (
	"time"

	crypto "unishard/crypto"
	message "unishard/message"
	quorum "unishard/quorum"
	types "unishard/types"

	"github.com/ethereum/go-ethereum/common"
)

type Sequence []*message.Transaction
type LocalSnapshots []*LocalSnapshot

type HzSVQckh1 byte

const (
	Coordination HzSVQckh1 = 0 + iota
	Worker
)

type CoordinationBlockHeader struct {
	Epoch         types.Epoch
	View          types.View
	GC            types.BlockHeight
	StateRoot     common.Hash
	PrevBlockHash common.Hash
	BlockHeight   types.BlockHeight
	Timestamp     time.Time
}

type CoordinationBlockData struct {
	GlobalCoordinationSequence Sequence
	GlobalSnapshot             LocalSnapshots
	GlobalContractBundle       []*LocalContract
}

type CoordinationBlock struct {
	BlockHeader        *CoordinationBlockHeader
	BlockHash          common.Hash
	QC                 *quorum.HPOWa1PT0xzq
	CQC                *quorum.HPOWa1PT0xzq
	CommitteeSignature []crypto.MqlfmVE9d
	Proposer           types.NodeID
	BlockData          *CoordinationBlockData
}
type O0GQK9U1 interface {
	WorkerBlockHeader() *ZkVeRG
	CoordinationBlockHeader() *CoordinationBlockHeader
	QuorumCertificate() *quorum.HPOWa1PT0xzq
	CommitQuorumCertificate() *quorum.HPOWa1PT0xzq
	WorkerBlockData() *WCGubgPD
	CoordinationBlockData() *CoordinationBlockData
	GetBlockHash() common.Hash
}

func (tNbRuD7nT *WorkerBlock) QuorumCertificate() *quorum.HPOWa1PT0xzq {
	return tNbRuD7nT.O7aFnBRNh
}

func (dhXZOVnFcobf *CoordinationBlock) QuorumCertificate() *quorum.HPOWa1PT0xzq {
	return dhXZOVnFcobf.QC
}

func (tNbRuD7nT *WorkerBlock) WorkerBlockHeader() *ZkVeRG {
	return tNbRuD7nT.BlockHeader
}

func (dhXZOVnFcobf *CoordinationBlock) CoordinationBlockHeader() *CoordinationBlockHeader {
	return dhXZOVnFcobf.BlockHeader
}

func (tNbRuD7nT *WorkerBlock) WorkerBlockData() *WCGubgPD {
	return tNbRuD7nT.BlockData
}

func (dhXZOVnFcobf *CoordinationBlock) CoordinationBlockData() *CoordinationBlockData {
	return dhXZOVnFcobf.BlockData
}

func (tNbRuD7nT *WorkerBlock) CoordinationBlockHeader() *CoordinationBlockHeader {
	/*line DWZvjqp3IZx.go:1*/ panic(func() /*line i6Yo7N.go:1*/ string {
		data := /*line i6Yo7N.go:1*/ make([]byte, 0, 14)
		i := 2
		decryptKey := 14
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 3
				data = /*line i6Yo7N.go:1*/ append(data, "T^a"...,
				)
			case 3:
				i = 4
				data = /*line i6Yo7N.go:1*/ append(data, "SO"...,
				)
			case 2:
				data = /*line i6Yo7N.go:1*/ append(data, "\\VNS"...,
				)
				i = 6
			case 4:
				for y := range data {
					data[y] = data[y] + /*line i6Yo7N.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 5:
				i = 1
				data = /*line i6Yo7N.go:1*/ append(data, "FO"...,
				)
			case 6:
				i = 5
				data = /*line i6Yo7N.go:1*/ append(data, "SP"...,
				)
			}
		}
		return /*line i6Yo7N.go:1*/ string(data)
	}())
}

func (dhXZOVnFcobf *CoordinationBlock) WorkerBlockHeader() *ZkVeRG {
	/*line bvwKKFHo8qry.go:1*/ panic(func() /*line nFwXL3X1w.go:1*/ string {
		data := /*line nFwXL3X1w.go:1*/ make([]byte, 0, 14)
		i := 8
		decryptKey := 217
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data = /*line nFwXL3X1w.go:1*/ append(data, 115)
				i = 0
			case 5:
				data = /*line nFwXL3X1w.go:1*/ append(data, 106)
				i = 10
			case 0:
				i = 6
				data = /*line nFwXL3X1w.go:1*/ append(data, "pz}"...,
				)
			case 7:
				for y := range data {
					data[y] = data[y] + /*line nFwXL3X1w.go:1*/ byte(decryptKey^y)
				}
				i = 2
			case 8:
				data = /*line nFwXL3X1w.go:1*/ append(data, 120)
				i = 4
			case 1:
				data = /*line nFwXL3X1w.go:1*/ append(data, 116)
				i = 9
			case 6:
				data = /*line nFwXL3X1w.go:1*/ append(data, "os"...,
				)
				i = 7
			case 4:
				data = /*line nFwXL3X1w.go:1*/ append(data, 114)
				i = 5
			case 9:
				i = 3
				data = /*line nFwXL3X1w.go:1*/ append(data, 106)
			case 10:
				i = 1
				data = /*line nFwXL3X1w.go:1*/ append(data, "ow"...,
				)
			}
		}
		return /*line nFwXL3X1w.go:1*/ string(data)
	}())
}

func (tNbRuD7nT *WorkerBlock) CoordinationBlockData() *CoordinationBlockData {
	/*line JTfyZQSHvH.go:1*/ panic(func() /*line Y7dbjf.go:1*/ string {
		fullData := [] /*line Y7dbjf.go:1*/ byte("A\xa3\xcf\x135d#W\xa9\x88@\x8c\n\xceE\xe2\xbc$3\xc7,\x84:\xf0B^")
		idxKey := [] /*line Y7dbjf.go:1*/ byte("l\x84'\xe9\x9f")
		data := /*line Y7dbjf.go:1*/ make([]byte, 0, 14)
		data = /*line Y7dbjf.go:1*/ append(data, fullData[141^ /*line Y7dbjf.go:1*/ int(idxKey[1])]-fullData[135^ /*line Y7dbjf.go:1*/ int(idxKey[1])], fullData[34^ /*line Y7dbjf.go:1*/ int(idxKey[2])]^fullData[43^ /*line Y7dbjf.go:1*/ int(idxKey[2])], fullData[140^ /*line Y7dbjf.go:1*/ int(idxKey[4])]-fullData[134^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[49^ /*line Y7dbjf.go:1*/ int(idxKey[2])]^fullData[32^ /*line Y7dbjf.go:1*/ int(idxKey[2])], fullData[155^ /*line Y7dbjf.go:1*/ int(idxKey[4])]^fullData[145^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[149^ /*line Y7dbjf.go:1*/ int(idxKey[4])]+fullData[139^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[241^ /*line Y7dbjf.go:1*/ int(idxKey[3])]+fullData[239^ /*line Y7dbjf.go:1*/ int(idxKey[3])], fullData[158^ /*line Y7dbjf.go:1*/ int(idxKey[4])]^fullData[146^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[125^ /*line Y7dbjf.go:1*/ int(idxKey[0])]^fullData[108^ /*line Y7dbjf.go:1*/ int(idxKey[0])], fullData[144^ /*line Y7dbjf.go:1*/ int(idxKey[4])]^fullData[148^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[136^ /*line Y7dbjf.go:1*/ int(idxKey[4])]^fullData[138^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[143^ /*line Y7dbjf.go:1*/ int(idxKey[4])]+fullData[151^ /*line Y7dbjf.go:1*/ int(idxKey[4])], fullData[126^ /*line Y7dbjf.go:1*/ int(idxKey[0])]-fullData[110^ /*line Y7dbjf.go:1*/ int(idxKey[0])])
		return /*line Y7dbjf.go:1*/ string(data)
	}())
}

func (dhXZOVnFcobf *CoordinationBlock) WorkerBlockData() *WCGubgPD {
	/*line wjmL7PQyhXl.go:1*/ panic(func() /*line Nm9uBC3eFW.go:1*/ string {
		key := [] /*line Nm9uBC3eFW.go:1*/ byte("\x88ad\xf8\xed:\xa1%\xdcwa\xe4Y")
		data := [] /*line Nm9uBC3eFW.go:1*/ byte("\xfd\xcf\xcde]\xa6\x06\x92A\xe5\xd5I\xbd")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return /*line Nm9uBC3eFW.go:1*/ string(data)
	}())
}

func (tNbRuD7nT *WorkerBlock) GetBlockHash() common.Hash {
	return tNbRuD7nT.BlockHash
}

func (dhXZOVnFcobf *CoordinationBlock) GetBlockHash() common.Hash {
	return dhXZOVnFcobf.BlockHash
}
func (tNbRuD7nT *WorkerBlock) CommitQuorumCertificate() *quorum.HPOWa1PT0xzq {
	return tNbRuD7nT.Jy8nZ9h7dU
}

func (dhXZOVnFcobf *CoordinationBlock) CommitQuorumCertificate() *quorum.HPOWa1PT0xzq {
	return dhXZOVnFcobf.CQC
}

type CoordinationBlockWithoutHeader struct {
	CommitteeSignature    []crypto.MqlfmVE9d
	CoordinationBlockData *CoordinationBlockData
}

type ProposedWorkerBlock struct {
	WorkerBlock          *WorkerBlock
	GlobalSnapshot       []*LocalSnapshot
	GlobalContractBundle []*LocalContract
}

type WorkerBlock struct {
	BlockHeader        *ZkVeRG
	BlockHash          common.Hash
	O7aFnBRNh          *quorum.HPOWa1PT0xzq
	Jy8nZ9h7dU         *quorum.HPOWa1PT0xzq
	CommitteeSignature []crypto.MqlfmVE9d
	Proposer           types.NodeID
	BlockData          *WCGubgPD
}

type ZkVeRG struct {
	Epoch        types.Epoch
	View         types.View
	GC           types.BlockHeight
	XftVzp       common.Hash
	GvMaqmnWSRO  common.Hash
	BlockHeight  types.BlockHeight
	IGa3VRMdDduL time.Time
	HK3UX13      types.Shard
}

type Qhrpt6Ktz9 struct {
	ReceivedCrossTransaction   []*message.Transaction
	ReceivedLocalTransaction   []*message.Transaction
	GlobalCoordinationSequence []*message.Transaction
	GlobalSnapshot             []*LocalSnapshot
	GlobalContractBundle       []*LocalContract
	GC                         types.BlockHeight
}

type WorkerConsensusPayload struct {
	Data *Qhrpt6Ktz9
}

type WCGubgPD struct {
	CJpmTwa4y                []*message.Transaction
	ReceivedCrossTransaction []*message.Transaction
	LocalSnapshot            []*LocalSnapshot
	LocalContractBundle      []*LocalContract
}

type Y7M2DYD struct {
	Uy2puNVLTPk  *WCGubgPD
	Q75od61      []crypto.MqlfmVE9d
	ZhsYCn3FaA2r *ALPYXqBtG
}

type J_aDshw struct {
	EVa2ObNZf7 *Y7M2DYD
	AJvp2jR    []*LocalSnapshot
	UzYQ7u     []*LocalContract
}

type LocalSnapshot struct {
	Address common.Address
	Slot    common.Hash
	Value   string
	RTCS    []common.Hash
}

type LocalContract struct {
	Address common.Address
	Code    []byte
}

type ALPYXqBtG struct {
	GBebKHb []*LocalSnapshot
	Na_xTB  []*LocalContract
}

func BC1Hpz9(mA2QN8z []*message.Transaction, rjNEFY4KGh []*message.Transaction, i6zKSW6W []*LocalSnapshot, gdMLZKB []*LocalContract) *WCGubgPD {
	jx2s151m1 := /*line BMtwDYPvz.go:1*/ new(WCGubgPD)
	jx2s151m1.CJpmTwa4y = mA2QN8z
	jx2s151m1.ReceivedCrossTransaction = rjNEFY4KGh
	jx2s151m1.LocalSnapshot = i6zKSW6W
	jx2s151m1.LocalContractBundle = gdMLZKB

	return jx2s151m1
}

func BRx2UrRXo(w3CyWcT_Eq *Y7M2DYD, hSpZO_F types.Epoch, e_TNaqT types.View, m1XbIgZY9y4r types.BlockHeight, sWmUTBdaq common.Hash, xb5YFd common.Hash, kIpAgVaA types.BlockHeight, kFN7XktCcIp *quorum.HPOWa1PT0xzq, fddpyk_YAhM types.Shard) O0GQK9U1 {
	iXnEH0yOR := &ZkVeRG{
		Epoch:       hSpZO_F,
		View:        e_TNaqT,
		GC:          m1XbIgZY9y4r,
		XftVzp:      sWmUTBdaq,
		GvMaqmnWSRO: xb5YFd,
		BlockHeight: kIpAgVaA + 1,
		HK3UX13:     fddpyk_YAhM,
	}
	aq4zSRIRs := /*line oPWaPiwNg.go:1*/ new(WorkerBlock)
	aq4zSRIRs.BlockHeader = iXnEH0yOR
	aq4zSRIRs.BlockHash = /*line gdlel3WUM8.go:1*/ aq4zSRIRs.MakeHash(aq4zSRIRs.BlockHeader)
	aq4zSRIRs.BlockData = w3CyWcT_Eq.Uy2puNVLTPk
	aq4zSRIRs.CommitteeSignature = w3CyWcT_Eq.Q75od61
	aq4zSRIRs.O7aFnBRNh = kFN7XktCcIp

	return aq4zSRIRs
}

func AiKaUM7(w3CyWcT_Eq *CoordinationBlockWithoutHeader, hSpZO_F types.Epoch, e_TNaqT types.View, m1XbIgZY9y4r types.BlockHeight, dVvzH5G common.Hash, xb5YFd common.Hash, kIpAgVaA types.BlockHeight, kFN7XktCcIp *quorum.HPOWa1PT0xzq) O0GQK9U1 {
	iXnEH0yOR := &CoordinationBlockHeader{
		Epoch:         hSpZO_F,
		View:          e_TNaqT,
		GC:            m1XbIgZY9y4r,
		StateRoot:     dVvzH5G,
		PrevBlockHash: xb5YFd,
		BlockHeight:   kIpAgVaA + 1,
	}
	ad6V39xkYQ := /*line toxGB733jFr.go:1*/ new(CoordinationBlock)
	ad6V39xkYQ.BlockHeader = iXnEH0yOR
	ad6V39xkYQ.BlockHash = /*line JidIwDC.go:1*/ ad6V39xkYQ.MakeHash(ad6V39xkYQ.BlockHeader)
	ad6V39xkYQ.BlockData = w3CyWcT_Eq.CoordinationBlockData
	ad6V39xkYQ.CommitteeSignature = w3CyWcT_Eq.CommitteeSignature
	ad6V39xkYQ.QC = kFN7XktCcIp

	return ad6V39xkYQ
}

func (tNbRuD7nT *CoordinationBlockWithoutHeader) MakeHash(jx2s151m1 interface{}) common.Hash {
	return /*line CJxT8f6zaD.go:1*/ crypto.ZsUswPaGG4Z(jx2s151m1)
}

func (tNbRuD7nT *CoordinationBlockData) MakeHash(jx2s151m1 interface{}) common.Hash {
	return /*line ntUGoVuZD.go:1*/ crypto.ZsUswPaGG4Z(jx2s151m1)
}

func (tNbRuD7nT *CoordinationBlock) MakeHash(jx2s151m1 interface{}) common.Hash {
	return /*line BThc1k.go:1*/ crypto.ZsUswPaGG4Z(jx2s151m1)
}

func (tNbRuD7nT *WorkerConsensusPayload) MakeHash(jx2s151m1 interface{}) common.Hash {
	return /*line h5WiqQrO.go:1*/ crypto.ZsUswPaGG4Z(jx2s151m1)
}

func (tNbRuD7nT *WCGubgPD) MakeHash(jx2s151m1 interface{}) common.Hash {
	return /*line Ahm2Ww0H.go:1*/ crypto.ZsUswPaGG4Z(jx2s151m1)
}

func (tNbRuD7nT *WorkerBlock) MakeHash(jx2s151m1 interface{}) common.Hash {
	return /*line _EOU5E.go:1*/ crypto.ZsUswPaGG4Z(jx2s151m1)
}

type ZAD7w_1hMTmR struct {
	ZzEFD8XF O0GQK9U1
}

type AIpVC8nUWuMr struct {
	TSVaV2tB2Jr O0GQK9U1
	IzKjft      types.BlockHeight
}

type W8Aocd38zMBU struct {
	PXwHQ3K *O0GQK9U1
}

type qUzMIbMJ_W struct {
	types.View
	TehvTKmC468  *quorum.HPOWa1PT0xzq
	DfbtEPI      *quorum.HPOWa1PT0xzq
	K7NGMk       types.NodeID
	RgBnVVJ5D9ms []common.Hash
	TPGXBYmw     common.Hash
	AADQWVHlFO1  crypto.MqlfmVE9d
	U0IFJz3i1be  common.Hash
}

type Accept struct {
	CommittedBlock *WorkerBlock
	*quorum.HPOWa1PT0xzq
	Kg1tLS2Ii    []*LocalSnapshot
	K0CmW5aPRDdY []*LocalContract
	JTK381N2H    time.Time
}

type Qi_7sYpWjR8 struct {
	XjPd77f0 *CoordinationBlock
	*quorum.HPOWa1PT0xzq
	Vbha3pl time.Time
}

const _ = time.ANSIC

var _ crypto.Pp__49cd
var _ message.ClientStart
var _ quorum.Commit

const _ = types.ABORT

var _ = common.AbsolutePath
