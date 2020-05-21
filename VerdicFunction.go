package main

import "fmt"
func addNumbers(nos ...int) int{
	sum:=0
	for _,i:=range nos{
		sum=sum+i
	}
	return sum
}
func main()  {
	fmt.Println(addNumbers(1,2,3,4,5))
}
