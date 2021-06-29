package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	if n == 1 {
		println(1)
		return
	}

	if n == 0 {
		println(10)
		return
	}

	digitCount := new([10]int)

	for n != 1 {
		found := false
		for d := 9; d >= 2; d-- {
			if n % d == 0 {
				n = n / d
				digitCount[d]++
				found = true
				break
			}
		}
		if !found {
			println(-1)
			return
		}
	}

	for d := 2; d <= 9; d++ {
		print(strings.Repeat(strconv.Itoa(d), digitCount[d]))
	}
	println()
}
