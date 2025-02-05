package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a) // 1234.523
	fmt.Println(b) // 3456.123
	fmt.Println(c) // 4.266663e+06
	fmt.Println(d) // 1.2799989e+07
}
