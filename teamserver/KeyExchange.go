package main

import (
	"math/big"
)

type Uint128 struct {
	LW uint64 // 低 8 个字节
	HG uint64 // 高 8 个字节
}

func NewUint128(low, high uint64) *Uint128 {
	return &Uint128{
		LW: low,
		HG: high,
	}
}

func (u *Uint128) ToBigInt() *big.Int {
	high := new(big.Int).Lsh(big.NewInt(0).SetUint64(u.HG), 64)
	low := new(big.Int).SetUint64(u.LW)
	return new(big.Int).Or(high, low)
}

func (u *Uint128) ModMulti(num *Uint128) *Uint128 {
	mod := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil), big.NewInt(1))
	result := new(big.Int).Mod(new(big.Int).Mul(u.ToBigInt(), num.ToBigInt()), mod)
	return NewUint128(result.Uint64(), new(big.Int).Rsh(result, 64).Uint64())
}

func (u *Uint128) ModPow(amount *Uint128) *Uint128 {
	mod := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil), big.NewInt(1))
	result := new(big.Int).Exp(u.ToBigInt(), amount.ToBigInt(), mod)
	return NewUint128(result.Uint64(), new(big.Int).Rsh(result, 64).Uint64())
}
