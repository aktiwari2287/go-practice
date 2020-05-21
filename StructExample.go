package main

import "fmt"

type Employee struct {
	 name string
	 address string
}
func main() {
	var emps = []Employee{
							{name:"Anand", address: "Ranchi"},
							{name: "Pooja", address: "Varanasi"},
							{name: "Bunty", address: "Bangalore"}}

	for i:=0;i<len(emps);i++ {
		fmt.Println(emps[i])
	}

	for _,e:= range emps {
		fmt.Println(e)
	}
}
