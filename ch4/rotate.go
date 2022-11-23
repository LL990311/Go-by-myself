// ex4.4
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 3)
	fmt.Println(s)
}

func rotate(arr []int, n int) {
	n %= len(arr)
	tmp := append(arr, arr[:n]...)
	fmt.Println(tmp)
	copy(arr, tmp[n:])
}
