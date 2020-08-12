package main

func addNumbers(nos ...int) int {
	sum := 0
	for _, i := range nos {
		sum = sum + i
	}
	return sum
}
func sum() interface{} {
	var ar = [2]int{2, 4}
	return ar
}
func main() {
	ar := sum()
	for i := range ar {

	}
}
