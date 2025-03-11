package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}

	for i, v := range t {
		fmt.Println(i, v)
	}
}

/**
0 1.1
1 2.2
2 3.3
3 4.4
4 5.5
*/
