package main

import "fmt"

func main() {
	a, b := 1, 1

OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

// 5 * 9 = 45
