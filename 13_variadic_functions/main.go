package main

import "fmt"

// func sum(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }
// func allTypes(nums ...interface{}) {
// 	for _, num := range nums {
// 		fmt.Println(num)
// 	}
// }
func allTypes(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	// fmt.Println(1, 2, 3, 4, 5, "sdfghjk")
	// fmt.Println(sum(1, 2, 3, 4, 5))
	// allTypes(1, 2, 3, 4, 5, "sdfghjk")
	// for slice
	nums := []int{1, 2, 3, 4, 5}
	allTypes(nums...) // this will not work for interface{} type because it will be treated as a single argument of type slice, not as multiple arguments of type int
	fmt.Println(nums)

}
