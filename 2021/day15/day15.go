package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

func getRiskMap(data []byte, five_times bool) {
	var output []int
	var output2 []int
	list := strings.Split(string(data), "\n")
	xmax := 0
	ymax := len(list)
	for _, s := range list {
		input := strings.Split(s, "")
		xmax = len(input)
		for _, value := range input {
			conv, _ := strconv.Atoi(value)
			output = append(output, conv)
		}
	}

	if five_times {
		for y := 0; y < 5; y++ {
			for w := 0; w < ymax; w++ {
				for x := 0; x < 5; x++ {
					for v := w * ymax; v < ((ymax * w) + ymax); v++ {
						value := output[v]
						value += x + y
						if value >= 10 {
							value -= 9
						}
						// fmt.Print(value)
						output2 = append(output2, value)
					}
				}
				// fmt.Print("\n")
			}
		}
		xmax *= 5
		ymax *= 5
		output = output2
	}

	graph := dijkstra.NewGraph()
	for i := range output {
		graph.AddVertex(i)
	}

	for i := range output {
		if i == 0 || (i%xmax) != xmax-1 {
			graph.AddArc(i, i+1, int64((output)[i+1]))
		}
		if i < len(output)-xmax {
			graph.AddArc(i, i+ymax, int64((output)[i+ymax]))
		}
		if (i % xmax) != 0 {
			graph.AddArc(i, i-1, int64((output)[i-1]))
		}
		if i >= ymax {
			graph.AddArc(i, i-ymax, int64((output)[i-ymax]))
		}
	}

	best, err := graph.Shortest(0, len(output)-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance)
}

func Chiton(data []byte, five_times bool) {
	getRiskMap(data, five_times)
}

func main() {
	Chiton(input, false)
	Chiton(input, true)
}
