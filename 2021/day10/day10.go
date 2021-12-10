package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

func getChunks(data []byte, chunk *[][]string) {
	list := strings.Split(string(data), "\n")
	for _, s := range list {
		temp := strings.Split(string(s), "")
		*chunk = append(*chunk, temp)
	}
}

func chunkHasFunc() int {
	result := 0
	var chunks [][]string
	getChunks(input, &chunks)
	for _, c := range chunks {
		result += findError(&c)
	}
	return result
}

func chunkHasFunc2() int {
	var results []int
	var chunks [][]string
	getChunks(input, &chunks)
	for _, c := range chunks {
		if value := isIncomplete(&c); value > 0 {
			results = append(results, value)
		}
	}
	sort.Ints(results)
	return results[len(results)/2]
}

func calcScore(value string) int {
	switch value {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	return 0
}

func remove(s *[]string, i int) {
	*s = append((*s)[:i], (*s)[i+1:]...)
}

// Returns if open or closed and the character needed for the close
func isOpenOrClosed(value string) string {
	switch value {
	case ")":
		return "("
	case "]":
		return "["
	case "}":
		return "{"
	case ">":
		return "<"
	default:
		return ""
	}
}

func findError(chunk *[]string) int {
	for i := 1; i < len(*chunk); {
		value := (*chunk)[i]
		if close := isOpenOrClosed(value); close != "" {
			if close != (*chunk)[i-1] {
				return calcScore(value)
			} else {
				// remove correctly closed parameters
				remove(chunk, i-1)
				remove(chunk, i-1)
				i = 1
			}
		} else {
			i++
		}
	}
	return 0
}

func isIncomplete(chunk *[]string) int {
	for i := 1; i < len(*chunk); {
		value := (*chunk)[i]
		if close := isOpenOrClosed(value); close != "" {
			if close != (*chunk)[i-1] {
				return 0
			} else {
				// remove correctly closed parameters
				remove(chunk, i-1)
				remove(chunk, i-1)
				i--
			}
		} else {
			i++
		}
	}
	score := 0
	for i := len(*chunk) - 1; i >= 0; i-- {
		score = score * 5
		score += calcScore((*chunk)[i])
	}
	return score
}

func main() {
	fmt.Println(chunkHasFunc())
	fmt.Println(chunkHasFunc2())
}
