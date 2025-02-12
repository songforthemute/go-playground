package main

import "fmt"

func main() {
	var x, y int32 = 7, 3
	var s, t float32 = 3.14, 5

	fmt.Println("x + y = ", x+y) // x + y =  10
	fmt.Println("x - y = ", x-y) // x - y =  4
	fmt.Println("x * y = ", x*y) // x * y =  21
	fmt.Println("x / y = ", x/y) // x / y =  2
	fmt.Println("x % y = ", x%y) // x % y =  1

	fmt.Println("s * t = ", s*t) // s * t =  15.700001
	fmt.Println("s / t = ", s/t) // s / t =  0.628
}
