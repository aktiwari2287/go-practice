package main

import "fmt"

func IntFilter(slice []int, predicate func(int) bool) []int {
	filtered := make([]int, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}
	return filtered
}
func matchingStrings(strings []string, queries []string) []int32 {
	i := 0
	arr := make([]int32, len(queries))
	for i < len(queries) {
		count := 0
		j := 0
		for j < len(strings) {
			if queries[i] == strings[j] {
				count += 1
			}
			j++
		}
		arr[i] = int32(count)
		i++
	}
	return arr
}
func main() {
	var arr = []string{"ab", "ba", "ab", "ba", "ab"}
	var query = []string{"ab", "ba"}
	r := matchingStrings(arr, query)
	fmt.Println(r)

}
