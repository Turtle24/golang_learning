package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 4, 5}

	for i, elem := range a {
		for j := i + 1; j < len(a); j++ {
			elem2 := a[j]
			if elem2 == elem {
				fmt.Println(elem)
			}
		}
	}
}
