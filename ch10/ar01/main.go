package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}

/**
24
25.9
27.8
26.9
26.2
*/
