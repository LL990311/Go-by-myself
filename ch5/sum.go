package main

import (
	"fmt"
	"math"
	"os"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//ex5.15

func max(vals ...int) int {
	if vals == nil {
		fmt.Println("at least one parameters")
		os.Exit(1)
	}
	max := math.MinInt32
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if vals == nil {
		fmt.Println("at least one parameters")
		os.Exit(1)
	}
	min := math.MaxInt32
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	//fmt.Println(sum())           //  0
	//fmt.Println(sum(3))          //  3
	//fmt.Println(sum(1, 2, 3, 4)) // 10
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(min(1, 2, 3, 4))
}
