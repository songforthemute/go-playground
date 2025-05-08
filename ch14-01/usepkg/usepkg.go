package main

import (
	"ch14-01/usepkg/usepkg/custompkg"
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo2/ch14/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

/**
Hello, this is a custom package!
This is Github expkg Sample
 10.00 ┤        ╭╮
  9.00 ┤   ╭╮   ││
  8.00 ┤   ││ ╭╮││
  7.00 ┤   │╰╮││││╭╮
  6.00 ┤  ╭╯ │││││││ ╭
  5.00 ┤ ╭╯  ╰╯╰╯│││╭╯
  4.00 ┤╭╯       ││││
  3.00 ┼╯        ││││
  2.00 ┤         ╰╯╰╯
*/
