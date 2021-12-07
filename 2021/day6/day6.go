package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type LanternFish int

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

func getFish(data []byte, fish *[]LanternFish) {
	fish_list := strings.Split(string(data), ",")
	for _, s := range fish_list {
		conv, _ := strconv.Atoi(s)
		*fish = append(*fish, LanternFish(conv))
	}
}

func GameOfLife(data []byte, days int) int {
	var fishes []LanternFish
	ocean := make([]int, 9)
	result := 0

	getFish(data, &fishes)

	for _, fish := range fishes {
		ocean[fish]++
	}

	for day := 0; day < days; day++ {
		next_day := make([]int, 9)
		duplicates := ocean[0]
		for i := range ocean {
			next_day[i] = ocean[((i + 1) % 9)]
		}
		next_day[6] += duplicates
		ocean = next_day
	}

	for i := range ocean {
		result += ocean[i]
	}

	return result
}

func main() {

	fmt.Println(GameOfLife(input, 18))
	fmt.Println(GameOfLife(input, 256))
}
