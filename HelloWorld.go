package main

import "fmt"

func calc(x, y int) (int, int, int, int) {
	sum := x+y
	mul := x*y
	div := x/7
	sub := x-y
	return sum, sub, mul, div
}
const (
	a, b int = 1 ,2
	c string = "Anand"
	d bool = true
)

func addArray(arr [5]int) int {
	var sum=0
	for i:=0;i<len(arr); i++{
		sum+=arr[i]
	}
	return sum
}
func main()  {
	mymap := map[string]int{
		"Anand":20,
	}
	fmt.Println(mymap)
}