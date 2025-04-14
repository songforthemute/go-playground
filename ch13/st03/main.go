package main

import "fmt"

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
}

/**
int32
54620
한
*/
