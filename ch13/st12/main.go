package main

import "fmt"

func main() {
	str1, str2, str3, str4 := "BBB", "aaaaAAA", "BBAD", "ZZZ"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4)
}

/**
BBB > aaaaAAA : false
BBB < BBAD : false
BBB <= ZZZ : true
*/
