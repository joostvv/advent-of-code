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

func getData(data []byte) (map[string]int64, map[string]string, string) {
	pairs := make(map[string]int64)
	translations := make(map[string]string)
	list := strings.Split(string(data), "\n")
	final_letter := ""
	for i := 0; i < len(list); i++ {
		// Skip empty items.
		value := list[i]
		if value == "" {
			continue
		}

		// At the first line a pair items is expected.
		if i == 0 {
			length := len(value) - 1
			for x := 0; x < length; x++ {
				pair := value[x : x+2]
				addPair(&pairs, pair, 1)
			}
			final_letter = value[length : length+1]
			continue
		}

		// Finally create map of translations
		translation_pair := strings.Split(value, " -> ")
		if len(translation_pair) == 2 {
			translations[translation_pair[0]] = translation_pair[1]
		}
	}
	return pairs, translations, final_letter
}

func getPairs(pair string, letter string) (string, string) {
	return pair[0:1] + letter, letter + pair[1:2]
}

func addPair(pairs *map[string]int64, pair string, value int64) {
	(*pairs)[pair] += value
}

func addValue(letter_count *map[string]int64, letter string, value int64) {
	if val, ok := (*letter_count)[letter]; ok {
		(*letter_count)[letter] = val + value
	} else {
		(*letter_count)[letter] = value
	}
}

func doStep(pairs *map[string]int64, translations *map[string]string) map[string]int64 {
	new_pairs := make(map[string]int64)
	for key, value := range *pairs {
		letter := (*translations)[key]
		pair1, pair2 := getPairs(key, letter)
		addPair(&new_pairs, pair1, value)
		addPair(&new_pairs, pair2, value)
	}
	return new_pairs
}

func mostCommonElement(pairs *map[string]int64, final_letter string) (int64, int64) {
	letter_count := make(map[string]int64)
	least_count := int64(math.MaxInt64)
	max_count := int64(0)

	for key, value := range *pairs {
		letter1 := key[0:1]
		addValue(&letter_count, letter1, value)
	}

	// Add final_letter 1 time
	addValue(&letter_count, final_letter, 1)

	for _, value := range letter_count {
		if value > max_count {
			max_count = value
		}
		if value < least_count {
			least_count = value
		}
	}
	return max_count, least_count
}

func Polymerization(data []byte, steps int) int64 {
	pairs, translations, final_letter := getData((data))
	for step := 0; step < steps; step++ {
		pairs = doStep(&pairs, &translations)
	}
	most, least := mostCommonElement(&pairs, final_letter)
	return most - least
}

func main() {
	fmt.Println(Polymerization(input, 10))
	fmt.Println(Polymerization(input, 40))
}
