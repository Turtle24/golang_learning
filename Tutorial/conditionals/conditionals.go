package main

import "fmt"

func main() {
	//  ! Not eg. !(7 < 5), !true
	//  || or
	//  && And

	val := !(5 < 7)
	fmt.Printf("%t", val)
}
