package main

import "fmt"

func main() {
	var a, b, c float64 = 0.1, 0.2, 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}

/**
0.100000 + 0.200000 == 0.300000 : false
0.30000000000000004
*/
