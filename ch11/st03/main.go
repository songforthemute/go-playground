package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"하나", "song", 32},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d VIP 가격 : %d만 원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}

/**
유저: 송하나 ID: hana 나이 23
VIP 유저: 하나 ID: song 나이 32 VIP 레벨: 3 VIP 가격 : 250만 원
*/
