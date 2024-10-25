//line :1
package types

type NodeID string
type NW4p1ya []NodeID

type Shard uint
type View uint
type BlockHeight uint
type Epoch uint

type EpochView struct {
	Epoch
	View
}

type MaZWZHI byte

const (
	TRANSFER MaZWZHI = 0 + iota
	SMARTCONTRACT
	DEPLOY
	ABORT
)

type JH_6YK byte

const (
	NORMALTRANSFER JH_6YK = 0 + iota
	NORMALSMARTCONTRACT
	CROSSSHARDTRANSFER
	CROSSSHARDSMARTCONTRACT
)

type TbLE_w byte

const (
	PENDING TbLE_w = 0 + iota
	CONFIRMED
	REJECTED
)

type T30k6y_tK byte

const (
	READY T30k6y_tK = 0 + iota
	PREPARED
	VIEWCHANGING
	LOCKED
	SPECULATIVEEXECUTION
	FILLHOLE
	FINISHFILLHOLE
	COMMIT
	CONFIRMREQ
)

type FP9GjX byte

const (
	LEADER FP9GjX = 0 + iota
	CANDIDATE
	COMMITTEE
	VALIDATOR
)

type ZIxaqWCiw byte

const (
	QC ZIxaqWCiw = 0 + iota
	CQC
)

type FGQEDLInwsw struct {
	WNHy5_Jk    []HucrUMG7j4X `json:"externalRweSet"`
	N0k3j4Tt    string        `json:"function"`
	EMQpRmSxByt string        `json:"functionSelector"`
	EZoqawtB8r4 []string      `json:"readSet"`
	TLgnY5kyM   []string      `json:"writeSet"`
}

type J_zoDMw struct {
	ReadSet      []string
	WriteSet     []string
	BtvFRP3tVe8k []HucrUMG7j4X
}

type HucrUMG7j4X struct {
	OP7awm    string `json:"name"`
	EDwUuQQ38 string `json:"type"`
}
