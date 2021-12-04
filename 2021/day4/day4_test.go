package main

import (
	_ "embed"
	"reflect"
	"testing"
)

type BingoCardTest struct {
	card   BingoCard
	result bool
}

type ScoreCardTest struct {
	card       BingoCard
	last_value byte
	result     int
}

type SetCardTest struct {
	card   BingoCard
	value  byte
	result BingoCard
}

func TestParseBingoInput(t *testing.T) {
	var expected_bingo_nrs = []byte{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	var expected_first_bingo_card = BingoCard{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}

	bingo_nrs, bingo_cards := ParseBingoInput(sample)

	if bingo_nrs == nil || bingo_cards == nil {
		t.Errorf("Return values are nil of TestParseBingoInput")
	} else if !reflect.DeepEqual(bingo_nrs, expected_bingo_nrs) {
		t.Errorf("Bingo numbers are not equal %d, %d", bingo_nrs, expected_bingo_nrs)
	} else if bingo_cards[0] != expected_first_bingo_card {
		t.Errorf("Bingo cards are not equal %d, %d", bingo_nrs, expected_bingo_nrs)
	}
}

// Test the check if the bingo card has bingo.
func TestBingo(t *testing.T) {
	cardtests := []BingoCardTest{
		BingoCardTest{card: BingoCard{100, 100, 100, 100, 100, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}, result: true},
		BingoCardTest{card: BingoCard{22, 13, 17, 11, 100, 8, 2, 23, 4, 100, 21, 9, 14, 16, 100, 6, 10, 3, 18, 100, 1, 12, 20, 15, 100}, result: true},
		BingoCardTest{card: BingoCard{100, 100, 100, 100, 1, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}, result: false},
		BingoCardTest{card: BingoCard{22, 13, 17, 11, 100, 8, 2, 23, 4, 100, 21, 9, 14, 16, 100, 6, 10, 3, 18, 100, 1, 12, 20, 15, 1}, result: false},
	}
	for _, cardtest := range cardtests {
		if result := Bingo(&cardtest.card); result != cardtest.result {
			t.Errorf("Bingo result not correct for card: %d\n Expected: %t, got %t", cardtest.card, cardtest.result, result)
		}
	}
}

// Test the calculation of the winning score.
func TestCalcScore(t *testing.T) {
	cardtests := []ScoreCardTest{
		ScoreCardTest{card: BingoCard{100, 100, 100, 100, 100, 10, 16, 15, 100, 19, 18, 8, 100, 26, 20, 22, 100, 13, 6, 100, 100, 100, 12, 3, 100}, last_value: 24, result: 4512},
	}

	for _, cardtest := range cardtests {
		if score := CalcScore(&cardtest.card, cardtest.last_value); score != cardtest.result {
			t.Errorf("End score not correct for card: %d and end value %d\n Expected: %d, got %d", cardtest.card, cardtest.last_value, cardtest.result, score)
		}
	}
}

// Test setting the value if the card has it.
func TestSetValueInCard(t *testing.T) {
	cardtests := []SetCardTest{
		SetCardTest{card: BingoCard{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19},
			value:  7,
			result: BingoCard{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 100, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}},
		SetCardTest{card: BingoCard{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7},
			value:  7,
			result: BingoCard{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 100}},
		SetCardTest{card: BingoCard{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19},
			value:  26,
			result: BingoCard{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}},
		SetCardTest{card: BingoCard{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7},
			value:  26,
			result: BingoCard{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 100, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7}},
	}

	for _, cardtest := range cardtests {
		card := cardtest.card
		if SetValueInCard(&card, cardtest.value); card != cardtest.result {
			t.Errorf("Set value not correct for card: %d and value %d\n Expected: %d, got %d", cardtest.card, cardtest.value, cardtest.result, card)
		}
	}
}
