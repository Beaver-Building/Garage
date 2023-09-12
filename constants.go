package main

import (
	"fmt"
	"math"
)

// c string type constant
const s string = "thisIsAString"

// function create a serial of constants including int, float
func Constants() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
