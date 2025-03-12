package main

import "fmt"

func main() {
	a, b := [5]int{1, 2, 3, 4, 5}, [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)

		if i == len(a) {
			fmt.Println()
		}
	}

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)

		if i == len(b) {
			fmt.Println()
		}
	}

	b = a

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

/**
a[0] = 1
a[1] = 2
a[2] = 3
a[3] = 4
a[4] = 5
b[0] = 500
b[1] = 400
b[2] = 300
b[3] = 200
b[4] = 100
b[0] = 1
b[1] = 2
b[2] = 3
b[3] = 4
b[4] = 5
*/
