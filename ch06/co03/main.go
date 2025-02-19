package main

import "fmt"

const Pig, Cow, Chicken int = 0, 1, 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("Pig")
	} else if animal == Cow {
		fmt.Println("Cow")
	} else if animal == Chicken {
		fmt.Println("Chicken")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}

/**
Pig
Cow
Chicken
*/
