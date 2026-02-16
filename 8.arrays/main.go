package main

import "fmt"

func main() {
	// var nums [5]int

	// nums[0] = 10
	// nums[1] = 20
	// nums[2] = 30
	// nums[3] = 40
	// nums[4] = 50
	//array length
	// fmt.Println(len(nums))
	//array values
	// var bools [3]bool
	// bools[1] = true

	// var strings [9]string
	// fmt.Println(strings)
	// var nums [5]int = [5]int{10, 20, 30, 40, 50}
	// fmt.Println(nums)
	//2d array
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix)
	// fixed size that is predictable and efficient but less flexible than slices
	// memory optimization and performance benefits due to contiguous memory allocation
	// content time access
}
