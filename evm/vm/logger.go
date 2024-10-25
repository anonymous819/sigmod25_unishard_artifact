//line :1
package vm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EVMLogger interface {
	CaptureTxStart(dYbZ70 uint64)
	CaptureTxEnd(m0OQdQ0vv uint64)

	CaptureStart(mlSixlcgTwS *EVM, uB7ekFhZ common.Address, gEiLwA5i5pVB common.Address, wGdwdlv bool, sIVM3RLPSKSh []byte, a0SqxOtD uint64, gVUwww *big.Int)
	CaptureEnd(dhOXtmt2Uz9 []byte, yPtvVe5tw uint64, v1p7xkTykQgN error)

	CaptureEnter(oyspArz0 LxosJe8, uB7ekFhZ common.Address, gEiLwA5i5pVB common.Address, sIVM3RLPSKSh []byte, a0SqxOtD uint64, gVUwww *big.Int)
	CaptureExit(dhOXtmt2Uz9 []byte, yPtvVe5tw uint64, v1p7xkTykQgN error)

	CaptureState(eAorajeVy084 uint64, nCnJ0mc LxosJe8, a0SqxOtD, pIJBYyGRV uint64, oXFAbeW4qW *GUBKEu339CA, ia7bD5D []byte, agAOS88IBm int, v1p7xkTykQgN error)
	CaptureFault(eAorajeVy084 uint64, nCnJ0mc LxosJe8, a0SqxOtD, pIJBYyGRV uint64, oXFAbeW4qW *GUBKEu339CA, agAOS88IBm int, v1p7xkTykQgN error)
}

const _ = big.Above

var _ = common.AbsolutePath
