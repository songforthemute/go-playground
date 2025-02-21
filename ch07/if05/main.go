// ch07/if05

package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("감사합니다")
		} else {
			fmt.Println("더치페이!")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("감사합니다")
		} else {
			fmt.Println("더치페이?")
		}
	} else {
		fmt.Println("결제 완료")
	}
}

// 더치페이?
