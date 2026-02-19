package main

// func processIt(fn func(a int) int) int {
// 	return fn(1)
// }
func processIt() func(a int) int {
	return func(a int) int {
		return 2
	}
}
func main() {
	// fn := func(a int) int {
	// 	return 2
	// }
	// fmt.Println(processIt(fn))
	fn := processIt()
	fn(5)

}
