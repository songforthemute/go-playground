package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)

	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // 비교
	/**
	x.Cmp(y)
		-1: x가 작음
		1: x가 큼
		0: 두 값이 같음
	*/
}

/**
0.1 0.2 0.3 0.3
0
*/
