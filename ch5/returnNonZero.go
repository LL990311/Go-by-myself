package main

import "fmt"

func main() {
	fmt.Println(returnNonZero(3))
}

func returnNonZero(n int) (res int) {
	defer func() {
		res = n
		recover()
	}()
	panic(1)
}
