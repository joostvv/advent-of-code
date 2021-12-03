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

func wayToGo(data []string) int {
	x, y := 0, 0

	for i := 0; i < len(data); i++ {
		values := strings.Fields(data[i])
		m, _ := strconv.Atoi(values[1])
		switch values[0] {
		case "forward":
			x += m
		case "down":
			y -= m
		case "up":
			y += m

		}
	}
	fmt.Println("%d,%d", x, y)
	return x * y
}

func wayToGo2(data []string) int {
	x, y, aim := 0, 0, 0

	for i := 0; i < len(data); i++ {
		values := strings.Fields(data[i])
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
	fmt.Println("%d,%d", x, y)
	return x * y
}

func main() {
	fmt.Println(wayToGo2(byteArrayToString(input)))
}
