package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input-sample.txt
var sample string

func getInput(input string) [][]*bool {
	var cucumbers [][]*bool
	cucumber_rows := strings.Split(input, "\n")
	for _, cucumber_row := range cucumber_rows {
		var cucumber_row_values []*bool
		cucumber_values := strings.Split(cucumber_row, "")
		for _, cucumber_value := range cucumber_values {
			cucumber_row_values = append(cucumber_row_values, stringToCucumber(cucumber_value))
		}
		cucumbers = append(cucumbers, cucumber_row_values)
	}
	return cucumbers
}

func stringToCucumber(cucumber string) *bool {
	var value bool
	switch cucumber {
	case ".":
		return nil
	case ">":
		value = false
		return &value
	case "v":
		value = true
		return &value
	default:
		return nil
	}
}

func cucumberToString(cucumber *bool) string {
	if cucumber == nil {
		return "."
	}
	if *cucumber {
		return "v"
	} else {
		return ">"
	}
}

func printInput(cucumbers [][]*bool) {
	for _, cucumber_row := range cucumbers {
		for _, cucumber_value := range cucumber_row {
			fmt.Print(cucumberToString(cucumber_value))
		}
		fmt.Println("")
	}
}

// Do a round and check for stalemate
func doRound(cucumbers [][]*bool) (bool, [][]*bool) {
	stalemate := true
	ymax := len(cucumbers)
	xmax := len(cucumbers[0])
	order := []bool{false, true}
	for _, cucumber_dir_ := range order {
		new_round := duplicate(cucumbers)
		for y, cucumber_row := range cucumbers {
			for x, cucumber_value := range cucumber_row {
				if cucumber_value != nil && *cucumber_value == cucumber_dir_ {
					y_step := y
					x_step := x
					// Value = v, check downwards
					if *cucumber_value {
						if y == ymax-1 {
							y_step = 0
						} else {
							y_step++
						}
					} else {
						if x == xmax-1 {
							x_step = 0
						} else {
							x_step++
						}
					}
					if cucumbers[y_step][x_step] == nil {
						stalemate = false
						new_round[y_step][x_step] = cucumber_value
						new_round[y][x] = nil
					}
				}
			}
		}
		cucumbers = new_round
	}
	return stalemate, cucumbers
}

func duplicate(matrix [][]*bool) [][]*bool {
	duplicate := make([][]*bool, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]*bool, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func KevinCucumber(input string) int {
	rounds := 0
	cucumber := getInput(input)
	finished := false
	for !finished {
		finished, cucumber = doRound(cucumber)
		rounds++
	}
	// printInput(cucumber)

	return rounds
}

func main() {
	fmt.Println(KevinCucumber(input))
}
