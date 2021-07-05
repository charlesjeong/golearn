package main

import (
	"fmt"

	"github.com/charlesjeong/learngo/usepkg/custompkg"
	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 7, 5, 10, 2, 6, 7, 3, 4, 3, 4}
	graph := asciigraph.Plot(data)
	fmt.Println(data, graph)
}
