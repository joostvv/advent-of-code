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

func byteArrayToUint(data []byte) []int64 {
	var int_data []int64
	lines := strings.Split(string(data), "\n")
	for _, s := range lines {
		conv, _ := strconv.ParseInt(s, 2, 64)
		int_data = append(int_data, conv)
	}
	return int_data
}

func checkBitsForPosition(data []int64, pos int) int {
	count := 0
	for _, value := range data {
		// Check if nth bit set
		if val := value & (1 << pos); val > 0 {
			count++
		}
	}
	return count
}

func powerConsumption(data []int64, bits_amount int) int {
	var gamma, epsilon int = 0, 0
	for pos := 0; pos < bits_amount; pos++ {
		count := checkBitsForPosition(data, pos)
		if count > (len(data) / 2) {
			gamma += (1 << pos)
		} else {
			epsilon += (1 << pos)
		}
	}
	return gamma * epsilon
}

func filterMCLCPosition(data []int64, pos int) ([]int64, []int64) {
	one_count, zero_count := 0, 0
	var one_values, zero_values []int64
	for _, value := range data {
		// Check if nth bit set and filter one values and zero values
		if val := value & (1 << pos); val > 0 {
			one_count++
			one_values = append(one_values, value)
		} else {
			zero_count++
			zero_values = append(zero_values, value)
		}
	}
	// When equal one_values have to be kept.
	if one_count >= zero_count {
		return one_values, zero_values
	}
	return zero_values, one_values
}

func lifeSupportRating(data []int64, bits_amount int) int64 {
	mc_values, lc_values := filterMCLCPosition(data, bits_amount-1)
	for pos := bits_amount - 2; pos >= 0; pos-- {
		if len(mc_values) == 1 && len(lc_values) == 1 {
			break
		}
		if len(mc_values) > 1 {
			mc_values, _ = filterMCLCPosition(mc_values, pos)
		}
		if len(lc_values) > 1 {
			_, lc_values = filterMCLCPosition(lc_values, pos)
		}
	}
	return mc_values[0] * lc_values[0]
}

func main() {
	fmt.Println(powerConsumption(byteArrayToUint(input), 12))
	fmt.Println(lifeSupportRating(byteArrayToUint(input), 12))
}
