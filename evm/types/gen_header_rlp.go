//line :1
package types

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

func (ecigZUHIEiA *Header) EncodeRLP(fq5uSoxn io.Writer) error {
	yabbhOmaVr := /*line c_H1cGlpWQQ.go:1*/ rlp.NewEncoderBuffer(fq5uSoxn)
	q8pI2Uc := /*line ayyLoYJk.go:1*/ yabbhOmaVr.List()
	/*line AGw3ajWa.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.ParentHash[:])
	/*line htAa1OC9rs.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.UncleHash[:])
	/*line ZS6JGWlqX.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Coinbase[:])
	/*line ahdOg3xQea.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Root[:])
	/*line fQxKnS.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.TxHash[:])
	/*line G21ubydX.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.ReceiptHash[:])
	/*line tIsrVH.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Bloom[:])
	if ecigZUHIEiA.Difficulty == nil {
		/*line k32vNehaX.go:1*/ yabbhOmaVr.Write(rlp.EmptyString)
	} else {
		if /*line VQcSRqvzwo73.go:1*/ ecigZUHIEiA.Difficulty.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		/*line UjboIQz.go:1*/ yabbhOmaVr.WriteBigInt(ecigZUHIEiA.Difficulty)
	}
	if ecigZUHIEiA.Number == nil {
		/*line B0S65ngGdFNx.go:1*/ yabbhOmaVr.Write(rlp.EmptyString)
	} else {
		if /*line n4KPJ3X.go:1*/ ecigZUHIEiA.Number.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		/*line NJgD7tZNUFWX.go:1*/ yabbhOmaVr.WriteBigInt(ecigZUHIEiA.Number)
	}
	/*line qIAyML.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.GasLimit)
	/*line UR4cq81I1c.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.GasUsed)
	/*line DN985wQC6YQ.go:1*/ yabbhOmaVr.WriteUint64(ecigZUHIEiA.Time)
	/*line Y7Wcgo4BUV.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Extra)
	/*line FrneaQ.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.MixDigest[:])
	/*line s7eds4AfylM.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.Nonce[:])
	falOt_ := ecigZUHIEiA.BaseFee != nil
	jvTIFU := ecigZUHIEiA.WithdrawalsHash != nil
	smP2mXldo := ecigZUHIEiA.BlobGasUsed != nil
	rfy5xClnZlHf := ecigZUHIEiA.ExcessBlobGas != nil
	hXWBPoTmt := ecigZUHIEiA.ParentBeaconRoot != nil
	if falOt_ || jvTIFU || smP2mXldo || rfy5xClnZlHf || hXWBPoTmt {
		if ecigZUHIEiA.BaseFee == nil {
			/*line wyB5w5b6.go:1*/ yabbhOmaVr.Write(rlp.EmptyString)
		} else {
			if /*line OPytbj9j1s.go:1*/ ecigZUHIEiA.BaseFee.Sign() == -1 {
				return rlp.ErrNegativeBigInt
			}
			/*line DZQ4naH.go:1*/ yabbhOmaVr.WriteBigInt(ecigZUHIEiA.BaseFee)
		}
	}
	if jvTIFU || smP2mXldo || rfy5xClnZlHf || hXWBPoTmt {
		if ecigZUHIEiA.WithdrawalsHash == nil {
			/*line Aw4vMh.go:1*/ yabbhOmaVr.Write([]byte{0x80})
		} else {
			/*line _su54T.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.WithdrawalsHash[:])
		}
	}
	if smP2mXldo || rfy5xClnZlHf || hXWBPoTmt {
		if ecigZUHIEiA.BlobGasUsed == nil {
			/*line F72kaE9w.go:1*/ yabbhOmaVr.Write([]byte{0x80})
		} else {
			/*line XlrRYCP17B.go:1*/ yabbhOmaVr.WriteUint64((*ecigZUHIEiA.BlobGasUsed))
		}
	}
	if rfy5xClnZlHf || hXWBPoTmt {
		if ecigZUHIEiA.ExcessBlobGas == nil {
			/*line QkCboAumbJ.go:1*/ yabbhOmaVr.Write([]byte{0x80})
		} else {
			/*line qpeNCFNd.go:1*/ yabbhOmaVr.WriteUint64((*ecigZUHIEiA.ExcessBlobGas))
		}
	}
	if hXWBPoTmt {
		if ecigZUHIEiA.ParentBeaconRoot == nil {
			/*line sRNfRmZMq0.go:1*/ yabbhOmaVr.Write([]byte{0x80})
		} else {
			/*line G3QXgTAlT8VO.go:1*/ yabbhOmaVr.WriteBytes(ecigZUHIEiA.ParentBeaconRoot[:])
		}
	}
	/*line Da4OiSB.go:1*/ yabbhOmaVr.ListEnd(q8pI2Uc)
	return /*line TFpHlnF7.go:1*/ yabbhOmaVr.Flush()
}

var _ = rlp.AppendUint64
var _ io.ByteReader
