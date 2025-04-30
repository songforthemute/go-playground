package main

import "fmt"

func main() {
	var str string = "Hello world"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%x\n", slice)
}

/**
Hello world
4865616c6f20776f726c64
*/
