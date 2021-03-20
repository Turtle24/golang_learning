package main

import (
	"strconv"
)

func test(n int) []string {
	ss := []string{}
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 && i%5 == 0 {
			s = "FizzBuzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else {
			s = strconv.Itoa(i)
		}
		ss = append(ss, s)
	}
	return ss
}

func main() {
	n := 3
	test(n)
}
