package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 20
	C = 20
	fmt.Println(&C)
}

/**
./main.go:8:6: declared and not used: b
./main.go:9:2: cannot assign to C (neither addressable nor a map index expression)
./main.go:10:15: invalid operation: cannot take address of C (constant 10 of type int)
*/
