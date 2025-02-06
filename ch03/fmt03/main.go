package main

import "fmt"

func main() {
	var a = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n", a) // 00324.13
	fmt.Printf("%08.2g\n", a) // 03.2e+02
	fmt.Printf("%8.5g\n", a)  //   324.13
	fmt.Printf("%f\n", c)     // 3.140000
}
