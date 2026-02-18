package main

import "fmt"

// slice -> dynamic
// most used construct in go
// +usefull methods
func main() {
	//unintialized slice is nil and has length and capacity of 0
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))
	// var nums = make([]int, 3)     // not a nil slice but has length and capacity of 3
	// var nums2 = make([]int, 3, 5) // length of 3 and capacity of 5
	// var nums2 = make([]int, 0, 5)
	// nums2 = append(nums2, 10)
	// nums2 = append(nums2, 4)
	// nums2 = append(nums2, 7)
	// nums2 = append(nums2, 9)
	// nums2 = append(nums2, 9)
	// // fmt.Println(nums)
	// fmt.Println(nums2)
	// fmt.Println(cap(nums2))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// //copy function
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// var nums = []int{1, 2, 3, 4, 5, 7}

	// var nums1 = []int{2, 1, 2, 3, 4, 7}

	// fmt.Println(slices.Equal(nums, nums1))
	// fmt.Println(slices.Equal(nums[1:4], nums1[2:5]))

	// fmt.Println(nums == nums1) // will not work because slices are reference types and they point to different underlying arrays

	//slice operator
	// 	fmt.Println(nums[1:4])
	// 	fmt.Println(nums[:4])
	// 	fmt.Println(nums[1:])

	//2d slices
	var grid = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(grid)
}
