package main

import "fmt"

func main() {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // float64 -> int
	d := float64(a * c) // int -> float64

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3

	fmt.Println(g, h, f) // 10 9 53
}
