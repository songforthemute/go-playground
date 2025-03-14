package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a, b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}

/**
Hello 4
expected integer
3 4
2 3 4
*/
