package main

import "fmt"

func removeDup(arr []string) []string {
	for i := 0; i < len(arr)-1; {
		if arr[i] == arr[i+1] {
			copy(arr[i:], arr[i+1:])
			arr = arr[:len(arr)-1]
		} else {
			i++
		}
	}
	return arr
}

func main() {
	s := []string{"a", "b", "b", "c", "c", "a"}
	s = removeDup(s)
	fmt.Println(s)

}
