package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hello world"
	var slice []byte = []byte(str)

	fmt.Printf("str: \t%p\n", unsafe.StringData(str))
	fmt.Printf("slice: \t%p\n", unsafe.SliceData(slice))
}

/**
str:    0x10477fd3e
slice:  0x14000102020
*/
