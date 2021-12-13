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

type Fold struct {
	axis  string
	value int
}

type Cord struct {
	x int
	y int
}

func getData(data []byte) ([][]string, []Fold) {
	var coordinates []Cord

	var folds []Fold
	list := strings.Split(string(data), "\n")
	split := false
	xmax, ymax := 0, 0
	for _, c := range list {
		if c == "" {
			split = true
			continue
		}
		if !split {
			// "111,222" to Cord(x:111,y:222)
			values := strings.Split(c, ",")
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(values[1])
			if x > xmax {
				xmax = x
			}
			if y > ymax {
				ymax = y
			}
			coordinates = append(coordinates, Cord{x: x, y: y})
		} else {
			// "fold along x=655" to Fold(axis:x, value:655)
			c = strings.Replace(c, "fold along ", "", -1)
			values := strings.Split(c, "=")
			value, _ := strconv.Atoi(values[1])
			folds = append(folds, Fold{axis: values[0], value: value})
		}
	}
	paper := make([][]string, ymax+1)
	for i := range paper {
		paper[i] = make([]string, xmax+1)
	}
	for _, coor := range coordinates {
		paper[coor.y][coor.x] = "#"
	}
	return paper, folds
}

func calcDots(paper [][]string, print bool) int {
	count := 0
	for _, y := range paper {
		for _, value := range y {
			if value != "#" {
				value = "."
			} else {
				count++
			}

			if print {
				fmt.Print(value)
			}
		}
		if print {
			fmt.Print("\n")
		}
	}
	return count
}

func foldPaper(paper [][]string, fold Fold) [][]string {
	// Check on which axis to fold
	var output [][]string
	if fold.axis == "x" {
		var first_half [][]string
		var second_half [][]string
		for _, half := range paper {
			first_half = append(first_half, half[:fold.value])
			second_half = append(second_half, half[fold.value+1:])
		}
		for y, y_val := range second_half {
			for x, x_val := range y_val {
				if x_val == "#" {
					first_half[y][fold.value-x-1] = x_val
				}
			}
		}
		output = first_half

	} else if fold.axis == "y" {
		first_half := paper[:fold.value][:]
		second_half := paper[fold.value+1:][:]
		for y, y_val := range second_half {
			for x, x_val := range y_val {
				if x_val == "#" {
					first_half[fold.value-y-1][x] = x_val
				}
			}
		}
		output = first_half
	}
	return output
}

func Origami(data []byte, fold_amount int, print bool) int {
	paper, folds := getData((data))
	for i := 0; i < fold_amount && i < len(folds); i++ {
		paper = foldPaper(paper, folds[i])
	}
	return calcDots(paper, print)
}

func main() {
	fmt.Println(Origami(input, 1, false))
	fmt.Println(Origami(input, 12, true))
}
