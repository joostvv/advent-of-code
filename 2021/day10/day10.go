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

func chunkHasFunc() (int, int) {
	var incompletes []int
	var chunks [][]string
	corrupted_score := 0

	getChunks(input, &chunks)

	for _, c := range chunks {
		corrupted, incomplete := findError(&c)
		if incomplete > 0 {
			incompletes = append(incompletes, incomplete)
		}
		corrupted_score += corrupted
	}

	// Get middle incomplete score
	sort.Ints(incompletes)
	incomplete_score := incompletes[len(incompletes)/2]

	return corrupted_score, incomplete_score
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

// Returns the character corresponding with the closed character, otherwise ""
func getOpenCharacter(value string) string {
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

func findError(chunk *[]string) (int, int) {
	for i := 1; i < len(*chunk); {
		value := (*chunk)[i]
		// Check if a close character.
		if open := getOpenCharacter(value); open != "" {
			// Check if open character matches.
			if open != (*chunk)[i-1] {
				return calcScore(value), 0
			} else {
				// Remove correctly opened and closed parameters.
				remove(chunk, i-1)
				remove(chunk, i-1)
				i -= 2
			}
		} else {
			i++
		}
	}
	// Calculate score for the completion strings
	score := 0
	for i := len(*chunk) - 1; i >= 0; i-- {
		score = score * 5
		score += calcScore((*chunk)[i])
	}
	return 0, score
}

func main() {
	fmt.Println(chunkHasFunc())
}
