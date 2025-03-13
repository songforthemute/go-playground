package main

import "fmt"

type Student struct {
	Age   int // 대문자로 시작하는 필드는 외부로 공개됨
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("Age: %d, No: %d, Score: %f", s.Age, s.No, s.Score)
}

func main() {
	student := Student{15, 23, 88.2}
	student2 := student

	PrintStudent(student2)
}

// Age: 15, No: 23, Score: 88.200000%
