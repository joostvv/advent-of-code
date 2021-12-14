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

		// At the first line the input items is expected that can be conveted to pairs.
		// Example: KHS to map[KH] = 1, map[HS] = 1
		if i == 0 {
			length := len(value) - 1
			for x := 0; x < length; x++ {
				pair := value[x : x+2]
				pairs[pair] += 1
			}
			// Safe the last letter in the input for the end (needed to calculate all letters amount).
			final_letter = value[length : length+1]
			continue
		}

		// Finally create a map of translations of pairs to added value.
		// Example: FV -> H to map[FH]=H
		translation_pair := strings.Split(value, " -> ")
		translations[translation_pair[0]] = translation_pair[1]
	}
	return pairs, translations, final_letter
}

func doStep(pairs *map[string]int64, translations *map[string]string) map[string]int64 {
	new_pairs := make(map[string]int64)
	for key, value := range *pairs {
		letter := (*translations)[key]
		pair1, pair2 := key[0:1]+letter, letter+key[1:2]
		new_pairs[pair1] += value
		new_pairs[pair2] += value
	}
	return new_pairs
}

func mostCommonElement(pairs *map[string]int64, final_letter string) (int64, int64) {
	letter_count := make(map[string]int64)
	least_count := int64(math.MaxInt64)
	max_count := int64(0)

	// Add only left side of pairs, such that every side of a pair is evaluated once.
	for key, value := range *pairs {
		letter := key[0:1]
		letter_count[letter] += value
	}

	// Add final_letter 1 time to complete count of letters.
	letter_count[final_letter] += 1

	// Find max and min count of all letters.
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
