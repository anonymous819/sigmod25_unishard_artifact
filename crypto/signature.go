//line :1
package crypto

import "math/big"

type MqlfmVE9d [][]byte
type Pp__49cd []MqlfmVE9d

type CkpjNpq struct {
	igyJjLnMaZ, wZmgNCG *big.Int
}

func (llQDu_QS *MqlfmVE9d) ToECDSA() CkpjNpq {
	var omwxZNgdq = /*line HxC458PFoEo.go:1*/ new(CkpjNpq)
	tlCX9t34s3 := *llQDu_QS
	var tkzeb39, eJtyIP big.Int
	/*line aGT3G209.go:1*/ tkzeb39.SetString( /*line ToLzI3K.go:1*/ string(tlCX9t34s3[0]), 10)
	/*line H0vqMaWjgd2H.go:1*/ eJtyIP.SetString( /*line ii6BL7mC1.go:1*/ string(tlCX9t34s3[1]), 10)
	omwxZNgdq.igyJjLnMaZ = &tkzeb39
	omwxZNgdq.wZmgNCG = &eJtyIP
	return *omwxZNgdq
}

const _ = big.Above
