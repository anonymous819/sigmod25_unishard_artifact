//line :1
package graph_dependency

import (
	blockchain "unishard/blockchain"
	types "unishard/types"

	"github.com/ethereum/go-ethereum/common"
)

type JF3zv0GskK struct {
	TransactionId common.Hash
	Write_set     []string
	Read_set      []string
}

func GenerateSet(xpAJp5 blockchain.Sequence) []JF3zv0GskK {
	mKEqdDLsg := /*line F_Xo7ADBL.go:1*/ make([]JF3zv0GskK, 0)
	for _, wxLuQo := range xpAJp5 {
		if wxLuQo.TXType == types.SMARTCONTRACT {
			for _, bssX3Kgd := range wxLuQo.RwSet {
				var uiIoi1x JF3zv0GskK
				uiIoi1x.TransactionId = wxLuQo.Hash
				for _, aJk1C00 := range bssX3Kgd.WriteSet {
					uiIoi1x.Write_set = /*line A8z2TDJucFC.go:1*/ append(uiIoi1x.Write_set /*line qXSujLaG4.go:1*/, bssX3Kgd.Address.Hex()+"@"+aJk1C00)
				}
				for _, wtAg5innW := range bssX3Kgd.ReadSet {
					uiIoi1x.Read_set = /*line yqBRXv.go:1*/ append(uiIoi1x.Read_set /*line nFc11vEE.go:1*/, bssX3Kgd.Address.Hex()+"@"+wtAg5innW)
				}
				uiIoi1x.Read_set = /*line pBJqT7TA.go:1*/ append(uiIoi1x.Read_set /*line N5KpEB.go:1*/, wxLuQo.From.Hex()+"@"+"0")
				uiIoi1x.Write_set = /*line FuY9HhlD.go:1*/ append(uiIoi1x.Write_set /*line RK0s5Sd.go:1*/, wxLuQo.From.Hex()+"@"+"0")
				mKEqdDLsg = /*line Sdxy5EmZeV9.go:1*/ append(mKEqdDLsg, uiIoi1x)
			}
		} else if wxLuQo.TXType == types.TRANSFER {
			var uiIoi1x JF3zv0GskK
			uiIoi1x.TransactionId = wxLuQo.Hash
			uiIoi1x.Read_set = /*line GaBA69qwDxl.go:1*/ append(uiIoi1x.Read_set, []string{ /*line eC5qAU1ty.go:1*/ wxLuQo.From.Hex() + "@" + "0" /*line EKBovxRaOX.go:1*/, wxLuQo.To.Hex() + "@" + "0"}...)
			uiIoi1x.Write_set = /*line pzZYoDar6XE.go:1*/ append(uiIoi1x.Write_set, []string{ /*line YWAHNRR_2KHI.go:1*/ wxLuQo.From.Hex() + "@" + "0" /*line a_mKR37.go:1*/, wxLuQo.To.Hex() + "@" + "0"}...)
			mKEqdDLsg = /*line EvGQBMQOPWDD.go:1*/ append(mKEqdDLsg, uiIoi1x)
		}
	}

	return mKEqdDLsg
}

var _ blockchain.Accept

const _ = types.ABORT

var _ = common.AbsolutePath
