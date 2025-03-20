package main

import "fmt"

func main() {
	var a, b int = 10, 20

	var p1, p2, p3 *int = &a, &a, &b

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
}

/**
p1 == p2 : true
p2 == p3 : false
*/
