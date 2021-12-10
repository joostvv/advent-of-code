package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

type point struct {
	x int
	y int
}

func getHeightMap(data []byte, output *[][]byte) {
	list := strings.Split(string(data), "\n")
	for _, s := range list {
		temp := make([]byte, 0)
		input := strings.Split(s, "")
		for _, value := range input {
			conv, _ := strconv.Atoi(value)
			temp = append(temp, byte(conv))
		}
		*output = append(*output, temp)
	}
}

// Find the lowest points
func atMyLowestPoints(heightmap *[][]byte) ([]byte, []point) {
	var lowest_points []byte
	var points []point
	xmax := len((*heightmap)[0]) - 1
	ymax := len((*heightmap)) - 1
	for y, xrow := range *heightmap {
		for x, height := range xrow {
			if x < xmax {
				// check further when no on the edge
				if height >= xrow[x+1] {
					continue
				}
			}
			if x > 0 {
				// check back at the edge
				if height >= xrow[x-1] {
					continue
				}
			}
			if y < ymax {
				// if not vertical bottom edge, check below
				if height >= (*heightmap)[y+1][x] {
					continue
				}
			}
			if y > 0 {
				// if not vertical top edge, check above
				if height >= (*heightmap)[y-1][x] {
					continue
				}
			}
			lowest_points = append(lowest_points, height)
			points = append(points, point{x, y})
		}
	}
	return lowest_points, points
}

func calculateScore(lowest_points *[]byte) int {
	result := 0
	for _, point := range *lowest_points {
		result += int(point + 1)
	}
	return result
}

func search(p *point, heightmap *[][]byte) int {
	count := 1
	xmax := len((*heightmap)[p.y]) - 1
	ymax := len((*heightmap)) - 1

	// Mark the cell
	(*heightmap)[p.y][p.x] = 9

	if p.x < xmax {
		// check further when no on the edge
		if (*heightmap)[p.y][p.x+1] != 9 {
			count += search(&point{p.x + 1, p.y}, heightmap)
		}
	}

	if p.x > 0 {
		// check back at the edge
		if (*heightmap)[p.y][p.x-1] != 9 {
			count += search(&point{p.x - 1, p.y}, heightmap)
		}
	}

	if p.y < ymax {
		// if not vertical bottom edge, check below
		if (*heightmap)[p.y+1][p.x] != 9 {
			count += search(&point{p.x, p.y + 1}, heightmap)
		}
	}

	if p.y > 0 {
		// if not vertical top edge, check above
		if (*heightmap)[p.y-1][p.x] != 9 {
			count += search(&point{p.x, p.y - 1}, heightmap)
		}
	}
	return count
}

func findBasins(points *[]point, heightmap *[][]byte) int {
	result := 0
	var results []int
	for _, p := range *points {
		results = append(results, search(&p, heightmap))
	}
	result = calculateScoreBasins(results)
	return result
}

func calculateScoreBasins(basins_count []int) int {
	sum := 1
	sort.Sort(sort.Reverse(sort.IntSlice(basins_count)))
	for i := 0; i < 3; i++ {
		sum *= basins_count[i]
	}
	return sum
}

// Teh teh teh, teh teh ta teh, teh teh teh, ta te tehhhhh
func SmokeUnderWater(data []byte) (int, int) {
	var heighthmap [][]byte

	getHeightMap(data, &heighthmap)
	lowest_points, coordiantes := atMyLowestPoints(&heighthmap)
	points := calculateScore(&lowest_points)
	basins := findBasins(&coordiantes, &heighthmap)
	return points, basins
}

func main() {
	fmt.Println(SmokeUnderWater(input))
}
