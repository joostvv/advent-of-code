package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BingoCard = [25]byte

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

// Parse a string to a byte array
func convertStringToByteArray(input string, seperation string) []byte {
	var output []byte
	split := strings.Split(input, seperation)
	for _, s := range split {
		conv, _ := strconv.Atoi(s)
		output = append(output, byte(conv))
	}
	return output
}

// Parse a string to a bingo card
func convertStringToBingoCard(input string) BingoCard {
	var bingo_card BingoCard

	// Remove trailing whitespace
	space := regexp.MustCompile(`\s+`)
	input = space.ReplaceAllString(input, " ")

	// Remove leading whitespace
	card := convertStringToByteArray(input[1:], " ")

	copy(bingo_card[:], card)
	return bingo_card
}

// Parse the input values to a list of picked up values and bingo score lists.
func ParseBingoInput(data []byte) ([]byte, []BingoCard) {
	var bingo_cards []BingoCard
	lines := strings.Split(string(data), "\n")
	bingo_nrs := convertStringToByteArray(lines[0], ",")

	// Drop first two lines not needed for the bingo cards
	lines = lines[2:]
	result := ""
	count := 0
	for _, s := range lines {
		result = result + " " + s
		count++
		if count == 6 {
			bingo_cards = append(bingo_cards, convertStringToBingoCard(result))
			result = ""
			count = 0
			continue
		}
	}
	return bingo_nrs, bingo_cards
}

// Check if the bingo card has bingo.
func Bingo(bingo_card *BingoCard) bool {
	// Check bingo in rows.
	for y := 0; y < 5; y++ {
		count := 0
		for x := 0; x < 5; x++ {
			if bingo_card[x+y*5] != 100 {
				// Skip if the row does not have a called out nr.
				break
			}
			count++
		}
		// Bingo for the x rows
		if count == 5 {
			return true
		}
	}
	// Check bingo in columns.
	for x := 0; x < 5; x++ {
		count := 0
		for y := 0; y < 5; y++ {
			if bingo_card[x+y*5] != 100 {
				// Skip if the column does not have a called out nr.
				break
			}
			count++
		}
		// Bingo for the x rows
		if count == 5 {
			return true
		}
	}
	return false
}

// Calculate winning score.
func CalcScore(bingo_card *BingoCard, winning_nr byte) int {
	var sum = 0
	for _, value := range bingo_card {
		if value != 100 {
			sum += int(value)
		}
	}
	return sum * int(winning_nr)
}

// Check if score has value and set it if it has it.
func SetValueInCard(bingo_card *BingoCard, nr byte) {
	for i, value := range bingo_card {
		if value == nr {
			bingo_card[i] = 100
		}
	}
}

// Remove Bingo card on a position from the list
func removeBingoCard(cards []BingoCard, i int) []BingoCard {
	return append(cards[:i], cards[i+1:]...)
}

// Calculate the winning bingo card and return the winning score.
func Squidgames(data []byte) int {
	values, cards := ParseBingoInput(data)
	for _, value := range values {
		for i := 0; i < len(cards); i++ {
			SetValueInCard(&cards[i], value)
			if Bingo(&cards[i]) {
				return CalcScore(&cards[i], value)
			}
		}
	}
	return 0
}

// Calculate the winning bingo card and return the winning score.
func SquidgamesPart2(data []byte) int {
	values, cards := ParseBingoInput(data)
	for _, value := range values {
		for i := 0; i < len(cards); i++ {
			SetValueInCard(&cards[i], value)
			if Bingo(&cards[i]) {
				if len(cards) == 1 {
					return CalcScore(&cards[i], value)
				}
				cards = removeBingoCard(cards, i)
				i--
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(Squidgames(input))
	fmt.Println(SquidgamesPart2(input))
}
