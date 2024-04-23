package main

import (
	"crypto/rand"
	"math/big"
)

func Bytes_To_BigInt(data []byte) *big.Int {
	return new(big.Int).SetBytes(data)
}

func Random_Big_Int128() *big.Int {
	randomInt := new(big.Int)
	randomInt.SetBit(randomInt, 128, 1) // 设置为 1，确保最高位为 1

	for {
		randBits := make([]byte, 16)
		_, err := rand.Read(randBits)
		if err != nil {
			panic(err)
		}

		randomInt.SetBytes(randBits)

		if randomInt.BitLen() == 128 {
			return randomInt
		}
	}
}

func Mod_Multi(u, num *big.Int) *big.Int {
	mod := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil), big.NewInt(1))
	result := new(big.Int).Mul(u, num)
	result.Mod(result, mod)
	return result
}

func Mod_Pow(u, amount *big.Int) *big.Int {
	mod := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil), big.NewInt(1))
	result := new(big.Int).Exp(u, amount, mod)
	return result
}
