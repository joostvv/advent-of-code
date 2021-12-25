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

func getInput(input string, rounds int) (map[int]bool, [][]bool) {
	enhancement_algorithm := make(map[int]bool)
	var image [][]bool
	values := strings.Split(input, "\n")
	enhancement_values := values[0]
	count := 0
	enhancements := strings.Split(enhancement_values, "")
	for _, onoff := range enhancements {
		enhancement_algorithm[count] = onoff == "#"
		count++
	}
	// Add one extra layer padding
	rounds++

	// Skip the newline in between
	values = values[2:]
	var padding []bool
	for i := 0; i < len(values[0])+(2*rounds); i++ {
		padding = append(padding, false)
	}
	for i := 0; i < rounds; i++ {
		image = append(image, padding)
	}
	for _, value := range values {

		var pixels []bool
		pixels_values := strings.Split(value, "")
		for i := 0; i < rounds; i++ {
			pixels = append(pixels, false)
		}
		for _, pixel := range pixels_values {

			pixels = append(pixels, pixel == "#")

		}
		for i := 0; i < rounds; i++ {
			pixels = append(pixels, false)
		}
		image = append(image, pixels)
	}
	for i := 0; i < rounds; i++ {

		image = append(image, padding)
	}
	return enhancement_algorithm, image
}

func printImage(image [][]bool) {
	for y := 0; y < len(image); y++ {
		for _, x := range image[y] {
			if x {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func calcEnhancement(image *[][]bool, x_start, y_start int) int {
	value := 0
	for y := y_start - 1; y <= y_start+1; y++ {
		for x := x_start - 1; x <= x_start+1; x++ {
			value = value << 1
			if (*image)[y][x] {
				value++
			}
		}
	}
	return value
}

func CopyImage(image *[][]bool) [][]bool {
	duplicate := make([][]bool, len(*image))
	for i := range *image {
		duplicate[i] = make([]bool, len((*image)[i]))
		copy(duplicate[i], (*image)[i])
	}
	return duplicate
}

func lightsOn(image *[][]bool) int {
	value := 0
	for y := 0; y < len(*image); y++ {
		for _, x := range (*image)[y] {
			if x {
				value++
			}
		}
	}
	return value
}

func Enhance(input string, rounds int) int {
	enhancement_algorithm, image := getInput(input, rounds)
	toggle := false
	if enhancement_algorithm[0] {
		toggle = !enhancement_algorithm[len(enhancement_algorithm)-1]
	}
	ymax := len(image)
	xmax := len(image[0])
	for round := 0; round < rounds; round++ {
		tmp := CopyImage(&image)
		for y := 0; y < ymax; y++ {
			for x := 0; x < xmax; x++ {
				// Do something special for the edges
				if (x == 0 || x == xmax-1) || (y == 0 || y == ymax-1) {
					if toggle {
						tmp[y][x] = round%2 == 0
					} else {
						tmp[y][x] = false
					}
				} else {
					value := calcEnhancement(&image, x, y)
					tmp[y][x] = enhancement_algorithm[value]
				}
			}
		}
		image = CopyImage(&tmp)
	}
	// fmt.Println("------------")
	// printImage(image)

	return lightsOn(&image)
}

func main() {
	fmt.Println(Enhance(input, 2))
	fmt.Println(Enhance(input, 50))
}
