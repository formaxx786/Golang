package main

import "fmt"

//iterating over data structures
func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// sum := 0
	// for i := 0; i < len(nums); i++ {
	// fmt.Println(nums[i])
	// 	sum += nums[i]
	// }
	// now using range
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)
	// for i, num := range nums {
	// 	fmt.Println("value:", num, "index:", i)
	// }

	// m := map[string]string{
	// 	"fname": "sunil",
	// 	"lname": "kumar",
	// }
	// for k, v := range m {
	// 	fmt.Println("key:", k, "value:", v)
	// }

	// c:gives Ascii value or unicode code point rune
	//i: giving starting byte of rune example: "Āolang"
	// for i, c := range "golang" {
	// 	fmt.Println(i, c)
	// }
	// for i, c := range "Āolang" {
	// 	fmt.Println(i, c)
	// }    // indexing will be like : 0 2 3 4 5 6

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
