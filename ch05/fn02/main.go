package main

import "fmt"

func main() {
	math := 80
	eng := 74
	history := 95
	fmt.Println("1 평균 점수 : ", (math+eng+history)/3)

	math = 88
	eng = 92
	history = 53
	fmt.Println("2 평균 점수 : ", (math+eng+history)/3)

	math = 78
	eng = 93
	history = 78
	fmt.Println("3 평균 점수 : ", (math+eng+history)/3)
}

/**
1 평균 점수 :  83
2 평균 점수 :  77
3 평균 점수 :  83
*/
