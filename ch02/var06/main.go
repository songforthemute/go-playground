package main

import "fmt"

var g int = 10 // 전역 변수 선언

func main() {
	var m int = 20 // 지역 변수 선언

	{
		var s int = 50       // 지역 변수 선언
		fmt.Println(m, s, g) // 20 50 10
	}

	// m = s + 20 // Error
} // main 함수 종료
