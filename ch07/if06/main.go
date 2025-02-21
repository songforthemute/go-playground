// ch07/if06

package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You're young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You're cool", age)
	} else {
		fmt.Println("Error")
	}

	fmt.Println(age) // age는 소멸되었음
}
