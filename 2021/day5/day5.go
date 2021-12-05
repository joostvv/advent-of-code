package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Vector struct {
	x  int
	y  int
	dx int
	dy int
}

type Diagram []int

func (v *Vector) diagonal() bool {
	return (v.x != v.dx) && (v.y != v.dy)
}

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

// Parse a string to a vector
func convertStringToVector(input string) Vector {
	var values []int
	input = strings.Replace(input, " -> ", ",", -1)
	split := strings.Split(input, ",")
	for _, s := range split {
		conv, _ := strconv.Atoi(s)
		values = append(values, conv)
	}
	output := Vector{x: values[0], y: values[1], dx: values[2], dy: values[3]}
	return output
}

func getVentVectors(data []byte, vectors *[]Vector) {
	vector_list := strings.Split(string(data), "\n")
	for _, s := range vector_list {
		*vectors = append(*vectors, convertStringToVector(s))
	}
}

func DetermineStep(w, dw int) int {
	if w < dw {
		return 1
	}
	return -1
}

func makeRange(w, dw int) []int {
	count := Abs(dw-w) + 1
	step := DetermineStep(w, dw)

	s := make([]int, count)
	for i := range s {
		s[i] = w
		w += step
	}
	return s
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func drawInDiagram(vector Vector, diagram *Diagram, width int) {
	// Determine start and stop coordinates
	if vector.diagonal() {
		range_x := makeRange(vector.x, vector.dx)
		range_y := makeRange(vector.y, vector.dy)
		for d := 0; d < len(range_y); d++ {
			(*diagram)[range_x[d]+range_y[d]*width] += 1
		}
	} else {
		range_x := makeRange(vector.x, vector.dx)
		range_y := makeRange(vector.y, vector.dy)
		for _, y := range range_y {
			for _, x := range range_x {
				(*diagram)[x+y*width] += 1
			}
		}
	}
}

func PrintDiagram(diagram *Diagram, width int) {
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			if (*diagram)[x+y*width] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print((*diagram)[x+y*width])
			}

		}
		fmt.Println()
	}
}

func FindOverlaps(diagram *Diagram) int {
	count := 0
	for _, s := range *diagram {
		if s > 1 {
			count++
		}
	}
	return count
}

// Calculate the places where lines overlap
func IHaveToVent(data []byte, width int, diagonal bool) int {
	var vectors []Vector
	diagram := make(Diagram, (width * width))
	getVentVectors(data, &vectors)
	for _, v := range vectors {
		if v.diagonal() && !diagonal {
			continue
		}
		drawInDiagram(v, &diagram, width)
	}
	PrintDiagram(&diagram, width)
	return FindOverlaps(&diagram)
}

func main() {
	fmt.Println(IHaveToVent(input, 1000, false))
	fmt.Println(IHaveToVent(input, 1000, true))
}
