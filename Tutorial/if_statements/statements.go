package main

import "fmt"

func main() {
	x := 13

	if x >= 18 {
		fmt.Println("Yes")
	} else if x < 18 && x > 12 {
		fmt.Println("You can ride with your parents")
	} else {
		fmt.Println("You cannot ride")
		fmt.Printf("Wait %d Years", 18-x)
	}
}
