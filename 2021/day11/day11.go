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

const xmax = 10
const ymax = 10

type Octi [ymax][xmax]byte

type point struct {
	x int
	y int
}

func getOcti(data []byte, octi *Octi) {
	list := strings.Split(string(data), "\n")
	for y, s := range list {
		x_values := strings.Split(s, "")
		for x, x_value := range x_values {
			conv, _ := strconv.Atoi(x_value)
			(*octi)[y][x] = byte(conv)
		}
	}
}

// Update all values in all directions and check if they flash because of the update.
func flash(p *point, octi *Octi) {
	if p.x < xmax-1 {
		// check back at the edge
		if (*octi)[p.y][p.x+1]++; (*octi)[p.y][p.x+1] == 10 {
			flash(&point{p.x + 1, p.y}, octi)
		}
		// check diagonals on the right side
		if p.y < ymax-1 {

			if (*octi)[p.y+1][p.x+1]++; (*octi)[p.y+1][p.x+1] == 10 {
				flash(&point{p.x + 1, p.y + 1}, octi)
			}
		}
		if p.y > 0 {
			if (*octi)[p.y-1][p.x+1]++; (*octi)[p.y-1][p.x+1] == 10 {
				flash(&point{p.x + 1, p.y - 1}, octi)
			}
		}
	}

	if p.x > 0 {
		// check back at the edge
		if (*octi)[p.y][p.x-1]++; (*octi)[p.y][p.x-1] == 10 {
			flash(&point{p.x - 1, p.y}, octi)
		}
		// heck diagonals on the left side
		if p.y < ymax-1 {

			if (*octi)[p.y+1][p.x-1]++; (*octi)[p.y+1][p.x-1] == 10 {
				flash(&point{p.x - 1, p.y + 1}, octi)
			}
		}
		if p.y > 0 {
			if (*octi)[p.y-1][p.x-1]++; (*octi)[p.y-1][p.x-1] == 10 {
				flash(&point{p.x - 1, p.y - 1}, octi)
			}
		}
	}

	// if not vertical bottom edge, check below
	if p.y < ymax-1 {
		if (*octi)[p.y+1][p.x]++; (*octi)[p.y+1][p.x] == 10 {
			flash(&point{p.x, p.y + 1}, octi)
		}
	}

	// if not vertical top edge, check above
	if p.y > 0 {
		if (*octi)[p.y-1][p.x]++; (*octi)[p.y-1][p.x] == 10 {
			flash(&point{p.x, p.y - 1}, octi)
		}
	}
	return
}

func blindingLights(octi *Octi) {
	for y, y_values := range *octi {
		for x, _ := range y_values {
			// Update every value with one.
			(*octi)[y][x] += 1
			// If because of this update the octopus flashes (and not other octi), flash.
			if (*octi)[y][x] == 10 {
				flash(&point{x, y}, octi)
			}
		}
	}
}

// Check values bigger than 10, count them and make them 0
func countFlashes(octi *Octi) int {
	flashes := 0
	for y, y_values := range *octi {
		for x, _ := range y_values {
			if (*octi)[y][x] >= 10 {
				(*octi)[y][x] = 0
				flashes++
			}
		}
	}
	return flashes
}

func printOcti(octi *Octi, step int, flashes int) {
	fmt.Println("Step:", step, "Flashes:", flashes)
	for _, y_values := range *octi {
		fmt.Println(y_values)
	}
	fmt.Println("")
}

func PinkElephantsOnParade(data []byte, steps int, print bool) (int, int) {
	var octi Octi
	var sync_flash []int
	total_flashes := 0
	first_sync_flash := 0

	getOcti(data, &octi)

	for step := 1; step <= steps; step++ {
		blindingLights(&octi)
		flashes := countFlashes(&octi)
		// Check if synchronized flash.
		if flashes == xmax*ymax {
			sync_flash = append(sync_flash, step)
		}
		total_flashes += flashes
		if print {
			printOcti(&octi, step, flashes)
		}
	}
	if sync_flash != nil {
		first_sync_flash = sync_flash[0]
	}

	return total_flashes, first_sync_flash
}

func main() {
	fmt.Println(PinkElephantsOnParade(input, 100, false))
	fmt.Println(PinkElephantsOnParade(input, 1000, false))
}
