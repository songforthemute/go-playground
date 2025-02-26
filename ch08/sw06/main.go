package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 33")
	default:
		fmt.Println("My age is: ", age)
	}

	fmt.Println("age is: ", age) // age 변수는 사라짐
}
