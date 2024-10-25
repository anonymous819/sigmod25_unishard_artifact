//line :1
package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
)

func PrivSign(dwSdbZ []byte, tNRz8u K3fvz3wD) (MqlfmVE9d, error) {
	var tkzeb39, eJtyIP *big.Int
	var iYBOpw error
	if tNRz8u != nil {
		tkzeb39, eJtyIP, iYBOpw = /*line A0Ract.go:1*/ ecdsa.Sign(rand.Reader, jYDXng /*line nVVrMZH.go:1*/, tNRz8u.ComputeHash(dwSdbZ))
		if iYBOpw != nil {
			return nil, iYBOpw
		}
	} else {
		tkzeb39, eJtyIP, iYBOpw = /*line SAvnp4.go:1*/ ecdsa.Sign(rand.Reader, jYDXng, dwSdbZ)
		if iYBOpw != nil {
			return nil, iYBOpw
		}
	}
	llQDu_QS := /*line WSNHpZuht.go:1*/ make([][]byte, 2)
	llQDu_QS[0] = [] /*line aqAOOe.go:1*/ byte( /*line ZpKjAeJ.go:1*/ tkzeb39.String())
	llQDu_QS[1] = [] /*line aGSfRX.go:1*/ byte( /*line WeRnqDP.go:1*/ eJtyIP.String())
	return llQDu_QS, iYBOpw
}

func SkrCuscT(llQDu_QS MqlfmVE9d, pu2etb KDSldlLG06As) (bool, error) {
	omwxZNgdq := /*line jNWQJv.go:1*/ llQDu_QS.ToECDSA()
	bhcWZp2iy := /*line __CTb6a.go:1*/ ecdsa.Verify(nRHqoC0RfsK_, pu2etb, omwxZNgdq.igyJjLnMaZ, omwxZNgdq.wZmgNCG)
	return bhcWZp2iy, nil
}

var _ = ecdsa.GenerateKey
var _ = rand.Int

const _ = big.Above
