//line :1
package types

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

func (ecigZUHIEiA *Log) EncodeRLP(fq5uSoxn io.Writer) error {
	yabbhOmaVr := /*line fEO29a.go:1*/ rlp.NewEncoderBuffer(fq5uSoxn)
	q8pI2Uc := /*line KMvTr3u3.go:1*/ yabbhOmaVr.List()
	/*line Qh1qNaB.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Address[:])
	falOt_ := /*line FzLliuoQ.go:1*/ yabbhOmaVr.List()
	for _, jvTIFU := range ecigZUHIEiA.Topics {
		/*line iRXTen.go:1*/ yabbhOmaVr.WriteBytes(jvTIFU[:])
	}
	/*line NXGTkzslo.go:1*/ yabbhOmaVr.ListEnd(falOt_)
	/*line EGeawf0Q.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Data)
	/*line epCiIJQ.go:1*/ yabbhOmaVr.ListEnd(q8pI2Uc)
	return /*line Jc8dEvjPcRP.go:1*/ yabbhOmaVr.Flush()
}

var _ = rlp.AppendUint64
var _ io.ByteReader
