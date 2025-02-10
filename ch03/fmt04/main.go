package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Print(str)
	/**
	Hello	Go		World
	"Go"is Awesome!
	*/

	fmt.Printf("%s", str)
	/**
	Hello	Go		World
	"Go"is Awesome!
	*/

	fmt.Printf("%q", str) // "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"
}
