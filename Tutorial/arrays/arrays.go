package main

import "fmt"

func main() {
	// var arr [5]int
	// arr[0] = 100
	// for x := 0; x < 5; x++ {
	// 	arr[x] = x
	// }
	// fmt.Println(arr)
	arr := [3]int{1, 2, 3}
	sum := 0

	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
