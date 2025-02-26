package main

import "fmt"

func main() {
	day := 3

	if day == 1 {
		fmt.Println("첫째 날입니다.")
		fmt.Println("오늘은 팀 미팅이 있습니다.")
	} else if day == 2 {
		fmt.Println("둘째 날입니다.")
		fmt.Println("오늘은 팀원 면접이 있습니다.")
	} else if day == 3 {
		fmt.Println("셋째 날입니다.")
		fmt.Println("오늘은 설계안을 확정하는 날입니다.")
	} else {
		fmt.Println("프로젝트를 진행하세요")
	}
}

/**
셋째 날입니다.
오늘은 설계안을 확정하는 날입니다.
*/
