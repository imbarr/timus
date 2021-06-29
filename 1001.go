package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	numbers := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		root := math.Sqrt(float64(numbers[i]))
		fmt.Printf("%.4f\n", root)
	}
}
