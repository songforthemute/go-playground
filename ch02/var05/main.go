package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a)

	fmt.Println(a) // 3456
	fmt.Println(c) // -128
}
