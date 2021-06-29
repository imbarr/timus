package main

import "fmt"

func main() {
	var N int
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		panic(err)
	}

	start := N / 3

	for d := start; d >= 1; d-- {
		if N % d == 0 {
			L := N/d - 1
			println(L)
			return
		}
	}

	println(0)
}
