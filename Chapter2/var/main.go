package main

import "fmt"

// := で型推論
func typeInference() {
	x := 10
	x, y := 30, "Hello"
	fmt.Println(x, y)
}

func main() {
	// var x int = 10
	// var y = 10
	var (
		x    int
		y        = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)
	fmt.Println(x, y, z, d, e, f, g)
	typeInference()
}
