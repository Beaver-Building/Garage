package main

import "fmt"

// a function named variables and including output of all type of variables
func Variables() {
	// variable declaration
	var a = "initial"
	fmt.Println(a)

	// variable declaration with multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// variable declaration with multiple variables and type inference
	// for more about interface please check empty interface and type assertion
	var d = true
	fmt.Println(d)

	// variable declaration, zero-valued
	var e int
	fmt.Println(e)

	// variable declaration with short hand
	f := "short"
	fmt.Println(f)
}
