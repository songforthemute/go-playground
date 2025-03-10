package main

import "fmt"

func main() {
	a, b, found := 1, 1, false

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/**
5 * 9 = 45
*/
