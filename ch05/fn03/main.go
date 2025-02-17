package main

import "fmt"

func PrintAvgScore(name string, math, eng, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, avg)
}

func main() {
	PrintAvgScore("1등", 80, 74, 95)
	PrintAvgScore("2등", 88, 92, 53)
	PrintAvgScore("3등", 78, 73, 78)
}

/**
1등 83
2등 77
3등 76
*/
