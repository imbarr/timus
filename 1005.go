package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20

var ChanCount = 100

func pileDiff(weights []int, c chan int) {
	if len(weights) == 0 {
		c <- 0
		close(c)
		return
	}

	first := weights[0]
	subChan := make(chan int, ChanCount)

	go pileDiff(weights[1:], subChan)

	for diff := range subChan {
		c <- diff + first
		c <- diff - first
	}
	close(c)
}

func main() {
	var length int
	_, err := fmt.Scanf("%d\n", &length)

	if err != nil {
		panic(err)
	}

	weights := make([]int, length)

	for i := 0; i < length; i++ {
		_, err := fmt.Scanf("%d", &weights[i])

		if err != nil {
			panic(err)
		}
	}

	start := time.Now()

	c := make(chan int, ChanCount)

	go pileDiff(weights, c)

	minDiff := math.MaxInt64
	for diff := range c {
		absDiff := int(math.Abs(float64(diff)))
		if minDiff > absDiff {
			minDiff = absDiff
		}
	}

	elapsed := time.Since(start)

	println(minDiff)
	log.Printf("%s", elapsed)
}
