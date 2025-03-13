package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "서울시 강동구"
	house.Size = 20
	house.Price = 9.8
	house.Type = "Apartment"

	fmt.Println("주소: ", house.Address)
	fmt.Printf("크기: %d 평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입: ", house.Type)
}

/**
주소:  서울시 강동구
크기: 20 평
가격: 9.80억 원
타입:  Apartment
*/
