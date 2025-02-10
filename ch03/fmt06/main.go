package main

import "fmt"

func main() {
	var a, b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

/**
3 4
2 3 4
*/

/**
Hello 4
0 expected integer
*/

/**
4 Hello
1 expected integer
*/
