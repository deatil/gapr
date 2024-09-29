package ec

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once

var p256 = &p256Curve{newPoint: bign.NewP256Element}

func initP256() {
	p256.params = &elliptic.CurveParams{
    	Name:    "BIGN256V1",
    	BitSize: 256,
    	P:       bigFromHex("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff43"),
    	N:       bigFromHex("ffffffffffffffffffffffffffffffffd95c8ed60dfb4dfc7e5abf99263d6607"),
    	B:       bigFromHex("77ce6c1515f3a8edd2c13aabe4d8fbbe4cf55069978b9253b22e7d6bd69c03f1"),
    	Gx:      bigFromHex("0000000000000000000000000000000000000000000000000000000000000000"),
    	Gy:      bigFromHex("6bf7fc3cfb16d69f5ce4c9a351d6835d78913966c408f6521e29cf1804516a93"),
    }
}

func initAll() {
	initP256()
}

func P256() elliptic.Curve {
	initonce.Do(initAll)
	return p256
}

func bigFromHex(s string) *big.Int {
	b, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("bign/elliptic: internal error: invalid encoding")
	}
	return b
}
