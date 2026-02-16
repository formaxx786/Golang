package main

import (
	"fmt"
)

func main() {
	// simple switch statement
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	//multiple conditions switch statement

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a weekday.")
	// }
	// type switch statement

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an integer with value \n", t)

		case string:
			fmt.Println("I'm a string with value \n", t)

		case bool:
			fmt.Println("I'm a boolean with value \n", t)
		default:
			fmt.Printf("I don't know what type I am\n", t)

		}
	}
	whoAmI(43)

}
