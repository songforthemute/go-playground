package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Unknown"
	}
}

func getMyFavorColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is ", colorToString(getMyFavorColor()))
}

// My favorite color is  Blue
