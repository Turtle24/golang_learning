package main

import "fmt"

func main() {
	ans := -1

	switch ans {
	case 1, -1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Not a case")
	}
}
