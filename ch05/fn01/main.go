package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c) // 9
}
