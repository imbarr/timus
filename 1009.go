package main

import (
	"fmt"
	"math"
)

var ChanCount9 = 100

func combinations(count int, c chan []bool, canStartWithZero bool) {
	if count == 0 {
		c <- make([]bool, 0)
		close(c)
		return
	}

	subChanZero := make(chan []bool, ChanCount9)

	go combinations(count - 1, subChanZero, true)

	for comb := range subChanZero {
		c <- append(comb, true)
	}

	if canStartWithZero {
		subChanNonZero := make(chan []bool, ChanCount9)
		go combinations(count - 1, subChanNonZero, false)

		for comb := range subChanNonZero {
			c <- append(comb, false)
		}
	}

	close(c)
}

func main() {
	var N int
	var K int

	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scanf("%d", &K)
	if err != nil {
		panic(err)
	}

	c := make(chan []bool, ChanCount9)

	go combinations(N, c, false)

	totalCount := 0
	for comb := range c {
		nonZeroCount := 0
		for _, isNonZero := range comb {
			if isNonZero {
				nonZeroCount++
			}
		}

		totalCount += int(math.Pow(float64(K - 1), float64(nonZeroCount)))
	}

	println(totalCount)
}
