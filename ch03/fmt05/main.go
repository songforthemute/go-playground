package main

import "fmt"

func main() {
	var a int
	var b int

	/**
	n: 성공적으로 입력한 값의 개수
	err: 입력 시 발생한 에러 반환
	*/
	n, err := fmt.Scan(&a, &b)

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
