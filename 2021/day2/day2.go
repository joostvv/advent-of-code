package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

func byteArrayToString(data []byte) []string {
	lines := strings.Split(string(data), "\n")
	return lines
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func wayToGo(data []byte) (int, int) {
	x, y, aim := 0, 0, 0
	pointers := byteArrayToString(data)

	for i := 0; i < len(pointers); i++ {
		values := strings.Fields(pointers[i])
		m, _ := strconv.Atoi(values[1])
		switch values[0] {
		case "forward":
			x += m
			y += m * aim
		case "down":
			aim -= m
		case "up":
			aim += m
		}
	}
	return Abs(x * aim), Abs(x * y)
}

func main() {
	fmt.Println(wayToGo(input))
}
