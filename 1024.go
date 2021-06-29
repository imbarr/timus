package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	perm := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		_, err := fmt.Scanf("%d", &perm[i])
		if err != nil {
			panic(err)
		}
	}

	lcm := 1
	for current := 1; current <= n; current++ {
		i := perm[current]
		count := 1
		for current != i {
			i = perm[i]
			count++
		}
		lcm = count*lcm / GCD(count, lcm)
	}

	println(lcm)
}
