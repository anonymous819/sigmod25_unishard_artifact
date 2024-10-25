//line :1
package types

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

func (ecigZUHIEiA *StateAccount) EncodeRLP(fq5uSoxn io.Writer) error {
	yabbhOmaVr := /*line w0OSOg2.go:1*/ rlp.NewEncoderBuffer(fq5uSoxn)
	q8pI2Uc := /*line maln3z.go:1*/ yabbhOmaVr.List()
	/*line oqNM9TbqNLwp.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.Nonce)
	if ecigZUHIEiA.Balance == nil {
		/*line MHNrmqx.go:1*/ yabbhOmaVr.Write(rlp.EmptyString)
	} else {
		/*line J4B8ga9QcHi.go:1*/ yabbhOmaVr.WriteUint256(ecigZUHIEiA.Balance)
	}
	/*line H003i7aanbOs.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Root[:])
	/*line A6rls228Jf.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.CodeHash)
	/*line sBi6q0ARC.go:1*/ yabbhOmaVr.ListEnd(q8pI2Uc)
	return /*line ZSq_dDkvL.go:1*/ yabbhOmaVr.Flush()
}

var _ = rlp.AppendUint64
var _ io.ByteReader
