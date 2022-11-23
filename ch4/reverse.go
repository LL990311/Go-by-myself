// ex4.3 reverse func with pointer param
package main

import "fmt"

func main() {
	s := [5]int{5, 6, 7, 8, 9}
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}

func reverse(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
