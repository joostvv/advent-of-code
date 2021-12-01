package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const SamplePath = "input-sample.txt"
const InputPath = "input.txt"

func ReadFileToStringArray(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Can't read file: ", path)
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func ReadFileToIntArray(path string) []int {
	var int_data []int
	data := ReadFileToStringArray(path)
	for _, s := range data {
		conv, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Can't convert value: ", s)
			panic(err)
		}
		int_data = append(int_data, conv)
	}
	return int_data
}

func thatIsDeepMan(data []int) int {
	count := 0
	for i, s := range data {
		if i == 0 {
			continue
		}
		if s > data[i-1] {
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
	data := ReadFileToIntArray(InputPath)
	fmt.Println(thatIsDeepMan(data))
	fmt.Println(thatIsDeepManPart2(data))
}
