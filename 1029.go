package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	var m int
	_, err = fmt.Scanf("%d\n", &m)
	if err != nil {
		panic(err)
	}

	weights := make([][]int, n)

	for i := 0; i < n; i++ {
		weights[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, err = fmt.Scanf("%d", &weights[i][j])
			if err != nil {
				panic(err)
			}
		}
		//_, err = fmt.Scanf("\n")
		//if err != nil {
		//	panic(err)
		//}
	}

	paths := make([][]int, n)
	paths[0] = make([]int, m)
	prev := make([][]int, n)
	prev[0] = make([]int, m)

	for j := 0; j < m; j++ {
		paths[0][j] = weights[0][j]
		prev[0][j] = -1
	}


	for i := 1; i < n; i++ {
		paths[i] = make([]int, m)
		prev[i] = make([]int, m)

		for j := 0; j < m; j++ {
			paths[i][j] = paths[i - 1][j] + weights[i][j]
			prev[i][j] = j
		}

		changed := true
		for changed {
			changed = false
			for j := 0; j < m; j++ {
				if j > 0 {
					newPath := paths[i][j - 1] + weights[i][j]
					if newPath < paths[i][j] {
						changed = true
						paths[i][j] = newPath
						prev[i][j] = j - 1
					}
				}

				if j < m - 1 {
					newPath := paths[i][j + 1] + weights[i][j]
					if newPath < paths[i][j] {
						changed = true
						paths[i][j] = newPath
						prev[i][j] = j + 1
					}
				}
			}
		}
	}

	minPath := math.MaxInt64
	minNode := -1

	for j := 0; j < m; j++ {
		if paths[n - 1][j] < minPath {
			minPath = paths[n - 1][j]
			minNode = j
		}
	}

	result := make([]int, 0)
	j := minNode
	i := n - 1

	for i != 0 {
		result = append(result, prev[i][j])
		if prev[i][j] == j {
			i--
		} else {
			j = prev[i][j]
		}
	}

	for k := len(result) - 1; k >= 0; k-- {
		print(result[k] + 1, " ")
	}
	print(result[0] + 1, "\n")
}
