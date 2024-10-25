//line :1
package types

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

//go:generate go run ../../rlp/rlpgen -type StateAccount -out gen_account_rlp.go

type StateAccount struct {
	Nonce    uint64
	Balance  *uint256.Int
	Root     common.Hash
	CodeHash []byte
}

func D1o8Dbhahg() *StateAccount {
	return &StateAccount{
		Balance:/*line EMVnh9Pm.go:1*/ new(uint256.Int),
		Root: NrGuaNA21,
		CodeHash:/*line LKqYW2Is.go:1*/ JhrQnETMFm.Bytes(),
	}
}

func (pj_1pSBU *StateAccount) Copy() *StateAccount {
	var jDFdTHHKfq *uint256.Int
	if pj_1pSBU.Balance != nil {
		jDFdTHHKfq = /*line FXGsXzCi.go:1*/ new(uint256.Int).Set(pj_1pSBU.Balance)
	}
	return &StateAccount{
		Nonce:   pj_1pSBU.Nonce,
		Balance: jDFdTHHKfq,
		Root:    pj_1pSBU.Root,
		CodeHash:/*line BQnaHfH6UkU.go:1*/ common.CopyBytes(pj_1pSBU.CodeHash),
	}
}

type SlimAccount struct {
	Nonce    uint64
	Balance  *uint256.Int
	Root     []byte
	CodeHash []byte
}

func K2bJTLXfmaB(gaP1S8 StateAccount) []byte {
	lkvtPyzzX9n9 := SlimAccount{
		Nonce:   gaP1S8.Nonce,
		Balance: gaP1S8.Balance,
	}
	if gaP1S8.Root != NrGuaNA21 {
		lkvtPyzzX9n9.Root = gaP1S8.Root[:]
	}
	if ! /*line ocUU2iaD.go:1*/ bytes.Equal(gaP1S8.CodeHash, JhrQnETMFm[:]) {
		lkvtPyzzX9n9.CodeHash = gaP1S8.CodeHash
	}
	bgjAklkHvJWu, rfteMJju := /*line SBTH66KjB3VM.go:1*/ rlp.EncodeToBytes(lkvtPyzzX9n9)
	if rfteMJju != nil {
		/*line jSqKjJDMNxe5.go:1*/ panic(rfteMJju)
	}
	return bgjAklkHvJWu
}

func VGE1getZq8(bgjAklkHvJWu []byte) (*StateAccount, error) {
	var lkvtPyzzX9n9 SlimAccount
	if rfteMJju := /*line FR0mcURRoaRG.go:1*/ rlp.DecodeBytes(bgjAklkHvJWu, &lkvtPyzzX9n9); rfteMJju != nil {
		return nil, rfteMJju
	}
	var gaP1S8 StateAccount
	gaP1S8.Nonce, gaP1S8.Balance = lkvtPyzzX9n9.Nonce, lkvtPyzzX9n9.Balance

	if /*line lLn70r.go:1*/ len(lkvtPyzzX9n9.Root) == 0 {
		gaP1S8.Root = NrGuaNA21
	} else {
		gaP1S8.Root = /*line ACqG58Z.go:1*/ common.BytesToHash(lkvtPyzzX9n9.Root)
	}
	if /*line o2pD5W4_90P2.go:1*/ len(lkvtPyzzX9n9.CodeHash) == 0 {
		gaP1S8.CodeHash = JhrQnETMFm[:]
	} else {
		gaP1S8.CodeHash = lkvtPyzzX9n9.CodeHash
	}
	return &gaP1S8, nil
}

func R3iwnzGSe(bgjAklkHvJWu []byte) ([]byte, error) {
	gaP1S8, rfteMJju := /*line nW2NJq5X.go:1*/ VGE1getZq8(bgjAklkHvJWu)
	if rfteMJju != nil {
		return nil, rfteMJju
	}
	return /*line Dyceezex.go:1*/ rlp.EncodeToBytes(gaP1S8)
}

var _ bytes.Buffer
var _ = common.AbsolutePath
var _ = rlp.AppendUint64
var _ = uint256.ErrBadBufferLength
