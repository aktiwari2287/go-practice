package main

import "fmt"

func main() {
	 var mymap = map[string]int{
		"Anand":20,
		"Kumar":33,
	}
	for k,v :=range mymap {
		fmt.Println(k, "\t" ,v)
	}
	fmt.Println(mymap["Anand"])
}
