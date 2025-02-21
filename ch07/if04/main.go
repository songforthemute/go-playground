// ch07/if04

package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 { // 함수가 호출되지 않음
		fmt.Println("1 increase")
	}

	if true || IncreaseAndReturn() < 5 { // 함수가 호출되지 않음
		fmt.Println("2 incerase")
	}

	fmt.Println("cnt: ", cnt)
}

// 2 incerase
// cnt:  0
