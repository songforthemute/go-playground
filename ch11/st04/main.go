package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		User{"하나", "song", 32, 12},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d 유저 레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.User.Level,
	)
}

/**
유저: 송하나 ID: hana 나이 23
VIP 유저: 하나 ID: song 나이 32 VIP 레벨: 3 유저 레벨: 12
*/
