package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}

	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteByte('A' + byte(c-'a'))
		} else {
			builder.WriteByte(byte(c))
		}
	}

	return builder.String()
}

func main() {
	var str string = "Hello world"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}

/**
HELLO WORLD
HELLO WORLD
*/
