package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	// mp := make(map[string]int)
	// mp["apple"] = 90
	// mp["kiwi"] = 10
	// delete(mp, "apple")

	val, ok := mp["apple"]

	fmt.Printf("%d: %t", val, ok)
}
