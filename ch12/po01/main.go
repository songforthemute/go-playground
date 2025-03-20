package main

import "fmt"

func main() {
	var a int = 500
	var p *int = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100

	fmt.Printf("a의 값: %d\n", a)
}

/**
p의 값: 0x14000104020
p가 가리키는 메모리의 값: 500
a의 값: 100
*/
