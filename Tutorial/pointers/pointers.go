package main

import "fmt"

// func changeValue(str *string) {
// 	*str = "changed"
// }

// func changeValue2(str string) string {
// 	str = "changed!"
// 	return str
// }

// func main() {
// 	toChange := "hello"
// 	fmt.Println(toChange)
// 	changeValue(&toChange)
// 	fmt.Println(toChange)
// }

func main() {
	// toChange := "hello"
	// var pointer *string = &toChange
	// fmt.Println(pointer)
	x := 5
	y := x
	fmt.Println(&x, &y)
}
