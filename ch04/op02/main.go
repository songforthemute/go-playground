package main

import "fmt"

func main() {
	var x1 int8 = 34
	var x2 int16 = 34
	var x3 int32 = 34
	var x4 int64 = 34

	fmt.Printf("^%d = %5d, \t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("^%d = %5d, \t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d, \t %016b\n", x4, ^x4, ^x4)
}

/**
^34 =   -35,     11011101
^34 =   -35,     1111111111011101
^34 =   -35,     -0100011
^34 =   -35,     -000000000100011
*/
