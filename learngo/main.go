package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)


func main() {
	data := []float64{8.17, 8.63, 8.65, 8.92, 8.95, 9.22}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
