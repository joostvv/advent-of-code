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

func getCrab(data []byte, crab *[]int) {
	crab_list := strings.Split(string(data), ",")
	for _, s := range crab_list {
		conv, _ := strconv.Atoi(s)
		*crab = append(*crab, conv)
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

func returnLowest(input *[]int, min int) int {
	sum := 0
	for _, s := range *input {
		sum += s
	}
	if sum < min {
		return sum
	}
	return min
}

func findLowestDistance(input *[]int) (int, int) {
	var crab_noob = make([]int, len(*input))
	var crab_engineer = make([]int, len(*input))
	min_crab_noob := math.MaxInt32
	min_crab_engineer := math.MaxInt32
	max := getHighestCrab(input)
	for i := 0; i <= max; i++ {
		for j, value := range *input {
			crab_noob[j] = Abs(value - i)
			// sum[n] with n = 1,2,...,n-1, n == n * n+1 / 2
			crab_engineer[j] = (crab_noob[j] * (crab_noob[j] + 1)) / 2
		}
		min_crab_noob = returnLowest(&crab_noob, min_crab_noob)
		min_crab_engineer = returnLowest(&crab_engineer, min_crab_engineer)
	}
	return min_crab_noob, min_crab_engineer
}

func CrabPeople(data []byte) (int, int) {
	var crabs []int

	getCrab(data, &crabs)
	crab_noob, crab_engineer := findLowestDistance(&crabs)

	return crab_noob, crab_engineer
}

func main() {
	fmt.Println(CrabPeople(input))
}
