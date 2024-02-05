package wallet

import (
	"crypto/elliptic"
	"math/big"
)

type SerializableP256Curve struct {
    P *big.Int
    N *big.Int
    B *big.Int
    Gx *big.Int
    Gy *big.Int
    BitSize int
}

func NewSerializableCurve() *SerializableP256Curve {
    params := elliptic.P256().Params()
    return &SerializableP256Curve{
        P: params.P,
        N: params.N,
        B: params.B,
        Gx: params.Gx,
        Gy: params.Gy,
        BitSize: params.BitSize,
    }
}
