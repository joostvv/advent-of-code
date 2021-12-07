package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

func getCrab(data []byte, fish *[]int) {
	crab_list := strings.Split(string(data), ",")
	for _, s := range crab_list {
		conv, _ := strconv.Atoi(s)
		*fish = append(*fish, conv)
	}
}

func getHighestCrab(input *[]int) int {
	max := 0
	for _, s := range *input {
		if s > max {
			max = s
		}
	}
	return max
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sum(input *[]int) int {
	sum := 0
	for _, s := range *input {
		sum += s
	}
	return sum
}

func sum(x int) int {
	if x == 0 {
		return 0
	}
	return x + sum(x-1)
}

func findLowestDistance2(input *[]int) int {
	var list = make([]int, len(*input))
	min := math.MaxInt32
	max := getHighestCrab(input)
	for i := 0; i <= max; i++ {
		for j, value := range *input {
			list[j] = sum(Abs(value - i))
		}
		if dist := Sum(&list); dist < min {
			min = dist
		}
	}
	return min
}

func findLowestDistance(input *[]int) int {
	var list = make([]int, len(*input))
	min := math.MaxInt32
	max := getHighestCrab(input)
	for i := 0; i <= max; i++ {
		for j, value := range *input {
			list[j] = Abs(value - i)
		}
		if dist := Sum(&list); dist < min {
			min = dist
		}
	}
	return min
}

func CrabPeople(data []byte) int {
	var crabs []int
	// ocean := make([]int, 9)
	result := 0

	getCrab(data, &crabs)

	result = findLowestDistance(&crabs)

	return result
}

func CrabPeople2(data []byte) int {
	var crabs []int
	result := 0

	getCrab(data, &crabs)

	result = findLowestDistance2(&crabs)

	return result
}

func main() {

	fmt.Println(CrabPeople(input))
	fmt.Println(CrabPeople2(input))
}
