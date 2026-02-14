package main

import "fmt"

//for -> only construct  in Go for looping
func main() {
	// while loop using for loop

	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//infinite loop

	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// classic for loop
	// for i := 1; i <= 4; i++ {
	// 	fmt.Println(i)
	// }

	//range
	for i := range 3 {
		fmt.Println(i)
	}
}
