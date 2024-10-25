//line :1
package vm

import (
	"math/big"

	types "unishard/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type StateDB interface {
	CreateAccount(common.Address)

	SubBalance(common.Address, *uint256.Int)
	AddBalance(common.Address, *uint256.Int)
	GetBalance(common.Address) *uint256.Int

	GetNonce(common.Address) uint64
	SetNonce(common.Address, uint64)

	GetCodeHash(common.Address) common.Hash
	GetCode(common.Address) []byte
	GetDeploycode(common.Address) []byte
	SetCode(common.Address, []byte)
	SetDeployCode(common.Address, []byte)
	GetCodeSize(common.Address) int

	AddRefund(uint64)
	SubRefund(uint64)
	GetRefund() uint64

	GetCommittedState(common.Address, common.Hash) common.Hash
	GetState(common.Address, common.Hash) common.Hash
	SetState(common.Address, common.Hash, common.Hash)

	GetTransientState(ewLbwkms common.Address, m9RTgbd0bs6o common.Hash) common.Hash
	SetTransientState(ewLbwkms common.Address, m9RTgbd0bs6o, gVUwww common.Hash)

	SelfDestruct(common.Address)
	HasSelfDestructed(common.Address) bool

	Selfdestruct6780(common.Address)

	Exist(common.Address) bool

	Empty(common.Address) bool

	AddressInAccessList(ewLbwkms common.Address) bool
	SlotInAccessList(ewLbwkms common.Address, w8fL8G3a1Z common.Hash) (oyuZY3We bool, xmJRd8 bool)

	AddAddressToAccessList(ewLbwkms common.Address)

	AddSlotToAccessList(ewLbwkms common.Address, w8fL8G3a1Z common.Hash)
	Prepare(avYeWr0Mmle params.Rules, qXjMoBVlo7Kx, tpMisd common.Address, diQ3lfnTEsj *common.Address, m6aBsCMavS []common.Address, lixnsCat types.AccessList)

	RevertToSnapshot(int)
	Snapshot() int

	AddLog(*types.Log)
	AddPreimage(common.Hash, []byte)
}

type VgQsz5khgz interface {
	Call(mlSixlcgTwS *EVM, a4kcg21 Ck0iTH, ewLbwkms common.Address, eiZFOFAwLXjh []byte, a0SqxOtD, gVUwww *big.Int) ([]byte, error)

	CallCode(mlSixlcgTwS *EVM, a4kcg21 Ck0iTH, ewLbwkms common.Address, eiZFOFAwLXjh []byte, a0SqxOtD, gVUwww *big.Int) ([]byte, error)

	DelegateCall(mlSixlcgTwS *EVM, a4kcg21 Ck0iTH, ewLbwkms common.Address, eiZFOFAwLXjh []byte, a0SqxOtD *big.Int) ([]byte, error)

	Create(mlSixlcgTwS *EVM, a4kcg21 Ck0iTH, eiZFOFAwLXjh []byte, a0SqxOtD, gVUwww *big.Int) ([]byte, common.Address, error)
}

const _ = big.Above

var _ types.AccessList
var _ = common.AbsolutePath
var _ = params.AllCliqueProtocolChanges
var _ = uint256.ErrBadBufferLength
