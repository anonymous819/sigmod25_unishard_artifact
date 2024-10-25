//line :1
package types

import (
	"bytes"
	"crypto/sha256"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

type BlobTx struct {
	ChainID    *uint256.Int
	Nonce      uint64
	GasTipCap  *uint256.Int
	GasFeeCap  *uint256.Int
	Gas        uint64
	To         common.Address
	Value      *uint256.Int
	Data       []byte
	AccessList AccessList
	BlobFeeCap *uint256.Int
	BlobHashes []common.Hash

	Sidecar *BlobTxSidecar `rlp:"-"`

	V *uint256.Int `json:"v" gencodec:"required"`
	R *uint256.Int `json:"r" gencodec:"required"`
	S *uint256.Int `json:"s" gencodec:"required"`
}

type BlobTxSidecar struct {
	Blobs       []kzg4844.Blob
	Commitments []kzg4844.Commitment
	Proofs      []kzg4844.Proof
}

func (dR4Hfi0 *BlobTxSidecar) BlobHashes() []common.Hash {
	qs4y7ialyr := /*line IOMnElB1lj.go:1*/ sha256.New()
	hP0dFanm4 := /*line cB3AzL00H.go:1*/ make([]common.Hash /*line u4z7qIPKsYh.go:1*/, len(dR4Hfi0.Commitments))
	for ibcOzM6f := range dR4Hfi0.Blobs {
		hP0dFanm4[ibcOzM6f] = /*line BLzH6nuDqO.go:1*/ kzg4844.CalcBlobHashV1(qs4y7ialyr, &dR4Hfi0.Commitments[ibcOzM6f])
	}
	return hP0dFanm4
}

func (dR4Hfi0 *BlobTxSidecar) qyTs7N2Qo54() uint64 {
	var lEWGRGoeJJXL, gp5EvJKhi91x, hbxhaWCy_ uint64
	for ibcOzM6f := range dR4Hfi0.Blobs {
		lEWGRGoeJJXL += /*line hOKpAbM.go:1*/ rlp.BytesSize(dR4Hfi0.Blobs[ibcOzM6f][:])
	}
	for ibcOzM6f := range dR4Hfi0.Commitments {
		gp5EvJKhi91x += /*line xkmuDRg0mB.go:1*/ rlp.BytesSize(dR4Hfi0.Commitments[ibcOzM6f][:])
	}
	for ibcOzM6f := range dR4Hfi0.Proofs {
		hbxhaWCy_ += /*line ECOS1Uur.go:1*/ rlp.BytesSize(dR4Hfi0.Proofs[ibcOzM6f][:])
	}
	return /*line uIEZLut.go:1*/ rlp.ListSize(lEWGRGoeJJXL) + /*line eaJiXI.go:1*/ rlp.ListSize(gp5EvJKhi91x) + /*line Gnryy8L.go:1*/ rlp.ListSize(hbxhaWCy_)
}

type blobTxWithBlobs struct {
	BlobTx      *BlobTx
	Blobs       []kzg4844.Blob
	Commitments []kzg4844.Commitment
	Proofs      []kzg4844.Proof
}

func (iPfjW8i91hC *BlobTx) acVy5yir() TxData {
	h5R0LDrm := &BlobTx{
		Nonce: iPfjW8i91hC.Nonce,
		To:    iPfjW8i91hC.To,
		Data:/*line YnD9iqsIbs2J.go:1*/ common.CopyBytes(iPfjW8i91hC.Data),
		Gas: iPfjW8i91hC.Gas,

		AccessList:/*line xicUFKVLxuW.go:1*/ make(AccessList /*line Cj9pEfTrl.go:1*/, len(iPfjW8i91hC.AccessList)),
		BlobHashes:/*line Q7zFKn5.go:1*/ make([]common.Hash /*line yqz2aY9PK.go:1*/, len(iPfjW8i91hC.BlobHashes)),
		Value:/*line UOZNBidOUd_k.go:1*/ new(uint256.Int),
		ChainID:/*line JsEGgBiN3t3.go:1*/ new(uint256.Int),
		GasTipCap:/*line NzP5Qi.go:1*/ new(uint256.Int),
		GasFeeCap:/*line DVZuHBHvURIK.go:1*/ new(uint256.Int),
		BlobFeeCap:/*line lL6TJMHH6.go:1*/ new(uint256.Int),
		V:/*line igcaTZ3.go:1*/ new(uint256.Int),
		R:/*line Hmidr6se.go:1*/ new(uint256.Int),
		S:/*line O5L7pDWX4AV.go:1*/ new(uint256.Int),
	}
	/*line tgJZlHajOGYA.go:1*/ copy(h5R0LDrm.AccessList, iPfjW8i91hC.AccessList)
	/*line IuoSzWeV.go:1*/ copy(h5R0LDrm.BlobHashes, iPfjW8i91hC.BlobHashes)

	if iPfjW8i91hC.Value != nil {
		/*line rBb92ccMawEe.go:1*/ h5R0LDrm.Value.Set(iPfjW8i91hC.Value)
	}
	if iPfjW8i91hC.ChainID != nil {
		/*line fooQ8ZrQZ.go:1*/ h5R0LDrm.ChainID.Set(iPfjW8i91hC.ChainID)
	}
	if iPfjW8i91hC.GasTipCap != nil {
		/*line rXv1UNv97vb.go:1*/ h5R0LDrm.GasTipCap.Set(iPfjW8i91hC.GasTipCap)
	}
	if iPfjW8i91hC.GasFeeCap != nil {
		/*line i8HzsFVoaUOA.go:1*/ h5R0LDrm.GasFeeCap.Set(iPfjW8i91hC.GasFeeCap)
	}
	if iPfjW8i91hC.BlobFeeCap != nil {
		/*line PlG6_2F.go:1*/ h5R0LDrm.BlobFeeCap.Set(iPfjW8i91hC.BlobFeeCap)
	}
	if iPfjW8i91hC.V != nil {
		/*line ewLtufaN_g.go:1*/ h5R0LDrm.V.Set(iPfjW8i91hC.V)
	}
	if iPfjW8i91hC.R != nil {
		/*line i3fAUn8.go:1*/ h5R0LDrm.R.Set(iPfjW8i91hC.R)
	}
	if iPfjW8i91hC.S != nil {
		/*line MeaiWoNjbY.go:1*/ h5R0LDrm.S.Set(iPfjW8i91hC.S)
	}
	if iPfjW8i91hC.Sidecar != nil {
		h5R0LDrm.Sidecar = &BlobTxSidecar{
			Blobs:/*line kXtFXR5RsZIj.go:1*/ append([] /*line wmU47gL666M.go:1*/ kzg4844.Blob(nil), iPfjW8i91hC.Sidecar.Blobs...),
			Commitments:/*line p_TEuIHV.go:1*/ append([] /*line jYAvGoJUV.go:1*/ kzg4844.Commitment(nil), iPfjW8i91hC.Sidecar.Commitments...),
			Proofs:/*line RcuBc2Swr3i.go:1*/ append([] /*line ozFby0E7W5.go:1*/ kzg4844.Proof(nil), iPfjW8i91hC.Sidecar.Proofs...),
		}
	}
	return h5R0LDrm
}

func (iPfjW8i91hC *BlobTx) b1pCNt1O() byte          { return BlobTxType }
func (iPfjW8i91hC *BlobTx) aVVDgwu() *big.Int       { return /*line SMF1jW.go:1*/ iPfjW8i91hC.ChainID.ToBig() }
func (iPfjW8i91hC *BlobTx) mJq92sCJCrD() AccessList { return iPfjW8i91hC.AccessList }
func (iPfjW8i91hC *BlobTx) bgjAklkHvJWu() []byte    { return iPfjW8i91hC.Data }
func (iPfjW8i91hC *BlobTx) go6R4MnIedw() uint64     { return iPfjW8i91hC.Gas }
func (iPfjW8i91hC *BlobTx) nwO6ag44() *big.Int {
	return /*line ggLMml2ex9A8.go:1*/ iPfjW8i91hC.GasFeeCap.ToBig()
}
func (iPfjW8i91hC *BlobTx) wIXPaTO() *big.Int {
	return /*line m4irgwP.go:1*/ iPfjW8i91hC.GasTipCap.ToBig()
}
func (iPfjW8i91hC *BlobTx) dtbQrx() *big.Int {
	return /*line aIEG15ZiB.go:1*/ iPfjW8i91hC.GasFeeCap.ToBig()
}
func (iPfjW8i91hC *BlobTx) fMqtswJ() *big.Int { return /*line rYOqrl.go:1*/ iPfjW8i91hC.Value.ToBig() }
func (iPfjW8i91hC *BlobTx) xiCJb__() uint64   { return iPfjW8i91hC.Nonce }
func (iPfjW8i91hC *BlobTx) ypKkxrFoZ() *common.Address {
	ckpdAKJvbnXP := iPfjW8i91hC.To
	return &ckpdAKJvbnXP
}
func (iPfjW8i91hC *BlobTx) zacAawX() uint64 {
	return params.BlobTxBlobGasPerBlob * /*line MES_Nu.go:1*/ uint64( /*line CnnLEcR8UC.go:1*/ len(iPfjW8i91hC.BlobHashes))
}

func (iPfjW8i91hC *BlobTx) iDoBrObf8(ySxt2pzafw *big.Int, oogLmTRZPjtz *big.Int) *big.Int {
	if oogLmTRZPjtz == nil {
		return /*line kif4XbD51oia.go:1*/ ySxt2pzafw.Set( /*line AeeaaoylXePa.go:1*/ iPfjW8i91hC.GasFeeCap.ToBig())
	}
	y64eru18K5 := /*line B9aiQR.go:1*/ ySxt2pzafw.Sub( /*line AcRZVnj.go:1*/ iPfjW8i91hC.GasFeeCap.ToBig(), oogLmTRZPjtz)
	if /*line tlCjCjVGX1.go:1*/ y64eru18K5.Cmp( /*line Okq_YLxHY.go:1*/ iPfjW8i91hC.GasTipCap.ToBig()) > 0 {
		/*line zWeA5z7Jt.go:1*/ y64eru18K5.Set( /*line fLcCzDftKfv.go:1*/ iPfjW8i91hC.GasTipCap.ToBig())
	}
	return /*line rbSlDQHBrYYi.go:1*/ y64eru18K5.Add(y64eru18K5, oogLmTRZPjtz)
}

func (iPfjW8i91hC *BlobTx) naFUnaJ8Fe() (pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	return /*line q5CRG9N5LNu.go:1*/ iPfjW8i91hC.V.ToBig() /*line m59_vgVM.go:1*/, iPfjW8i91hC.R.ToBig() /*line EbbrnC.go:1*/, iPfjW8i91hC.S.ToBig()
}

func (iPfjW8i91hC *BlobTx) ccFa3ozV(aVVDgwu, pW1UMRn1s, fWuxre8U, gDp0rg *big.Int) {
	/*line y9aSaWSaQ.go:1*/ iPfjW8i91hC.ChainID.SetFromBig(aVVDgwu)
	/*line PNTQ9kj.go:1*/ iPfjW8i91hC.V.SetFromBig(pW1UMRn1s)
	/*line W0X_X2S.go:1*/ iPfjW8i91hC.R.SetFromBig(fWuxre8U)
	/*line ldm1RxOR_rvx.go:1*/ iPfjW8i91hC.S.SetFromBig(gDp0rg)
}

func (iPfjW8i91hC *BlobTx) ini_YlFWY3() *BlobTx {
	h5R0LDrm := *iPfjW8i91hC
	h5R0LDrm.Sidecar = nil
	return &h5R0LDrm
}

func (iPfjW8i91hC *BlobTx) iJPSdycNFb(aYao5YS *bytes.Buffer) error {
	if iPfjW8i91hC.Sidecar == nil {
		return /*line A6T73pXccpC.go:1*/ rlp.Encode(aYao5YS, iPfjW8i91hC)
	}
	vv4w9F := &blobTxWithBlobs{
		BlobTx:      iPfjW8i91hC,
		Blobs:       iPfjW8i91hC.Sidecar.Blobs,
		Commitments: iPfjW8i91hC.Sidecar.Commitments,
		Proofs:      iPfjW8i91hC.Sidecar.Proofs,
	}
	return /*line IbaVQP.go:1*/ rlp.Encode(aYao5YS, vv4w9F)
}

func (iPfjW8i91hC *BlobTx) h5OeLZu4u(uzD2Up []byte) error {

	_QdL64ghyIP, _, rfteMJju := /*line j4ZbMFHKOn6.go:1*/ rlp.SplitList(uzD2Up)
	if rfteMJju != nil {
		return rfteMJju
	}
	c6m_rsAxNP, _, _, rfteMJju := /*line aGe4xIZJ2x.go:1*/ rlp.Split(_QdL64ghyIP)
	if rfteMJju != nil {
		return rfteMJju
	}

	if c6m_rsAxNP != rlp.List {
		return /*line JSFZsa8hY8m.go:1*/ rlp.DecodeBytes(uzD2Up, iPfjW8i91hC)
	}

	var vv4w9F blobTxWithBlobs
	if rfteMJju := /*line ZCKapw.go:1*/ rlp.DecodeBytes(uzD2Up, &vv4w9F); rfteMJju != nil {
		return rfteMJju
	}
	*iPfjW8i91hC = *vv4w9F.BlobTx
	iPfjW8i91hC.Sidecar = &BlobTxSidecar{
		Blobs:       vv4w9F.Blobs,
		Commitments: vv4w9F.Commitments,
		Proofs:      vv4w9F.Proofs,
	}
	return nil
}

var _ bytes.Buffer

const _ = sha256.BlockSize
const _ = big.Above

var _ = common.AbsolutePath
var _ kzg4844.Blob
var _ = params.AllCliqueProtocolChanges
var _ = rlp.AppendUint64
var _ = uint256.ErrBadBufferLength
