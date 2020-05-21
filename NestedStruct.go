package main

import "fmt"

type Animal struct {
	Name,
	Origin string
}

type Bird struct {
	Animal
	CanFly bool
}
func main() {
	var dog = Animal{Name:"Dog", Origin: "India"}
	var crow = Bird{
		Animal: Animal { Name:"A", Origin: "B" },
		CanFly:true,
		}
	/*crow.Name = "crow"
	crow.Origin = "USA"
	crow.CanFly = true*/
	fmt.Println(dog, crow)
}
