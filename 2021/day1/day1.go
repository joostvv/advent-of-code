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

func byteArrayToInt(data []byte) []int {
	var int_data []int
	lines := strings.Split(string(data), "\n")
	for _, s := range lines {
		conv, _ := strconv.Atoi(s)
		int_data = append(int_data, conv)
	}
	return int_data
}

func thatIsDeepMan(data []int) int {
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}
	return count
}

func thatIsDeepManPart2(data []int) int {
	var converted_data []int
	for i := 0; i < len(data)-2; i++ {
		sum := 0
		for _, s := range data[i : i+3] {
			sum += s
		}
		converted_data = append(converted_data, sum)
	}
	return thatIsDeepMan(converted_data)
}

func main() {
	fmt.Println(thatIsDeepMan(byteArrayToInt(input)))
	fmt.Println(thatIsDeepManPart2(byteArrayToInt(input)))
}
