package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50}

	nums[2] = 300

	for _, num := range nums { // for i := 0; i < len(nums); i++ {
		fmt.Println(num) // fmt.Println(nums[i])
	}
}

/**
10
20
300
40
50
*/
