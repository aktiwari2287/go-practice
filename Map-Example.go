package main

import "fmt"

func createMapUsingMake() {
	var m = make(map[string]string)
	m["name"] = "Anand"
	m["address"] = "Ranchi"
	for k, v := range m {
		fmt.Println(k + "\t" + v)
	}
}
func main() {
	var mymap = map[string]int{
		"Anand": 20,
		"Kumar": 33,
	}
	for k, v := range mymap {
		fmt.Println(k, "\t", v)
	}
	fmt.Println(mymap["Anand"])

	createMapUsingMake()
}
