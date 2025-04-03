package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	var a = Actor{name, hp, speed}
	return &a
}

func Ex02() {
	var actor = NewActor("Bob", 100, 50)
	fmt.Println(actor)
}
