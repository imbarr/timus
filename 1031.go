package main

import (
	"fmt"
	"math"
)

func main() {
	var L1, L2, L3, C1, C2, C3, N, A, B int

	_, err := fmt.Scanf("%d %d %d %d %d %d\n%d\n%d %d\n", &L1, &L2, &L3, &C1, &C2, &C3, &N, &A, &B)
	if err != nil {
		panic(err)
	}

	A = A - 1
	B = B - 1

	if A > B {
		A, B = B, A
	}

	dist := make([]int, N)
	dist[0] = 0

	for i := 1; i < N; i++ {
		_, err = fmt.Scanf("%d\n", &dist[i])
		if err != nil {
			panic(err)
		}
	}

	prices := make([]int, N)
	prices[A] = 0

	for i := A + 1; i <= B; i++ {
		minPrice := math.MaxInt64
		for j := i - 1; j >= A; j-- {
			d := dist[i] - dist[j]
			var price int
			if d <= L1 {
				price = C1
			} else if d <= L2 {
				price = C2
			} else if d <= L3 {
				price = C3
			} else {
				break
			}

			totalPrice := prices[j] + price
			if totalPrice < minPrice {
				minPrice = totalPrice
			}
		}
		prices[i] = minPrice
	}

	println(prices[B])
}
