package main

import "fmt"

func factorial(a int) int {
	if a <= 1 {
		return a
	}
	return a * factorial(a-1)
}
func main() {
	fmt.Println(factorial(5))
}
