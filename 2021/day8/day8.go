package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

func getData(data []byte, signal, output *[][]string) {
	list := strings.Split(string(data), "\n")
	for _, s := range list {
		s = strings.Replace(s, " | ", "|", -1)
		input := strings.Split(s, "|")
		*signal = append(*signal, strings.Split(input[0], " "))
		*output = append(*output, strings.Split(input[1], " "))
	}
}

func findUniqueSegments(output *[]string) int {
	count := 0
	for _, s := range *output {
		switch len(s) {
		case 2:
			count++
		case 3:
			count++
		case 4:
			count++
		case 7:
			count++
		default:
			continue
		}
	}
	return count
}

// find the letters representing 2 and 4
func find2And4(signal *[]string) {
	var twoandfour = make([]string, 2)
	for _, s := range *signal {
		switch len(s) {
		case 2:
			twoandfour[0] = s
		case 4:
			twoandfour[1] = s
		default:
			continue
		}
	}
	*signal = twoandfour
}

func Segments(data []byte) (int, int) {
	var signal, output [][]string
	digits_score := 0
	output_score := 0

	getData(data, &signal, &output)

	for i, out := range output {
		find2And4(&signal[i])
		output_score += decodeOutput(&signal[i], &out)
		digits_score += findUniqueSegments(&out)
	}
	return digits_score, output_score
}

// Decode values 0,6 and 9 which have 6 segment values
func decodeLenSix(twoandfour *[]string, output string) int {
	lettersTwo := strings.Split((*twoandfour)[0], "")
	lettersFour := strings.Split((*twoandfour)[1], "")

	count2 := 0
	for _, s := range lettersTwo {
		count2 += strings.Count(output, s)
	}

	// Six does not have the segment digits of 1
	if count2 != 2 {
		return 6
	}

	count4 := 0
	for _, s := range lettersFour {
		count4 += strings.Count(output, s)
	}

	// 9 does have all the digits of 4
	if count4 == 4 {
		return 9
	}

	// Only possible solution is then 0
	return 0
}

// Decode values 3, 2 and 5 which have 5 segment values
func decodeLenFive(twoandfour *[]string, output string) int {
	lettersTwo := strings.Split((*twoandfour)[0], "")
	lettersFour := strings.Split((*twoandfour)[1], "")

	count2 := 0
	for _, s := range lettersTwo {
		count2 += strings.Count(output, s)
	}

	// 3 does have the all the segment digits of 1 and the others don't
	if count2 == 2 {
		return 3
	}

	count4 := 0
	for _, s := range lettersFour {
		count4 += strings.Count(output, s)
	}

	// 5 does have 3 matches with the digits of 4, 2 has 2 matches
	if count4 == 3 {
		return 5
	}

	// Only possible solution is then 2
	return 2
}

func decodeOutput(twoandfour, output *[]string) int {
	sum := 0
	length := len(*output)
	for i := 0; i < length; i++ {
		value := 0
		letters := (*output)[i]
		switch len(letters) {
		case 2:
			value = 1
		case 3:
			value = 7
		case 4:
			value = 4
		case 5:
			value = decodeLenFive(twoandfour, letters)
		case 6:
			value = decodeLenSix(twoandfour, letters)
		case 7:
			value = 8
		}
		sum += int(math.Pow(10, float64(length-i-1))) * value
	}
	return sum
}

func main() {
	fmt.Println(Segments(input))
}
