package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a: ", a, "b: ", b)                // a: 10b: 20
	fmt.Println("a: ", a, "b: ", b, "f: ", f)    // 0a:  10 b:  20 f:  3.27994387438297e+10
	fmt.Printf("a: %d, b: %d, f: %f\n", a, b, f) // a: 10, b: 20, f: 32799438743.829700
}
