package main

import "fmt"

func add(p1, p2, p3 *int) {
	*p3 = *p1 + *p2
}

func Ex01() {
	a, b, c := 3, 5, 0

	add(&a, &b, &c)
	fmt.Println(c) // 8
}
