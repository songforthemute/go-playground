package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날입니다.")
		fmt.Println("오늘은 팀 미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날입니다.")
		fmt.Println("오늘은 팀원 면접이 있습니다.")
	case 3:
		fmt.Println("셋째 날입니다.")
		fmt.Println("오늘은 설계안을 확정하는 날입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요")
	}
}

/**
셋째 날입니다.
오늘은 설계안을 확정하는 날입니다.
*/
