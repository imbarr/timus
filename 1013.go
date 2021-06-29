package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N int
	var K int
	var M int

	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scanf("%d\n", &K)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scanf("%d", &M)
	if err != nil {
		panic(err)
	}

	zeroComb := new(big.Int).SetInt64(1)
	nonZeroComb := new(big.Int).SetInt64(int64(K - 1))
	mBig := new(big.Int).SetInt64(int64(M))

	for i := 1; i < N; i++ {
		newZeroComb := new(big.Int).Set(nonZeroComb)

		nonZeroComb.Add(nonZeroComb, zeroComb)
		nonZeroComb.Mul(nonZeroComb, new(big.Int).SetInt64(int64(K - 1)))
		nonZeroComb.Mod(nonZeroComb, mBig)

		zeroComb = newZeroComb
	}

	println(nonZeroComb.String())
}
