package main

import "fmt"

func main() {
	offset := 32768
	sum := 10000
	max := offset * 2 + sum + 1
	checkList := make([]bool, max)

	var N int
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		panic(err)
	}

	for i := 0; i < N; i++ {
		var number int
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			panic(err)
		}

		checkList[offset + number] = true
	}

	var K int
	_, err = fmt.Scanf("%d\n", &K)
	if err != nil {
		panic(err)
	}

	for i := 0; i < K; i++ {
		var number int
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			panic(err)
		}

		if checkList[offset + sum - number] {
			println("YES")
			return
		}
	}

	println("NO")
}
