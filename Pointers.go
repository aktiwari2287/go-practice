package main

import "fmt"

func main() {
	var p *int
	var a = 10
	b:=&a
	p=&a
	fmt.Println("Hello ",**&b, *&a, p)
}
