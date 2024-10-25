//line :1
package types

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

func (ecigZUHIEiA *Withdrawal) EncodeRLP(fq5uSoxn io.Writer) error {
	yabbhOmaVr := /*line rZ5WcDzA02.go:1*/ rlp.NewEncoderBuffer(fq5uSoxn)
	q8pI2Uc := /*line MldTDRsVy.go:1*/ yabbhOmaVr.List()
	/*line uRv9zXdfM.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.Index)
	/*line PUxfKKwVQ.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.Validator)
	/*line lGsNRnqO.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Address[:])
	/*line VxQ8l2.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.Amount)
	/*line LAoB27W.go:1*/ yabbhOmaVr.ListEnd(q8pI2Uc)
	return /*line N0PYAS.go:1*/ yabbhOmaVr.Flush()
}

var _ = rlp.AppendUint64
var _ io.ByteReader
