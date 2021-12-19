package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input-sample.txt
var sample string

type Vector struct {
	x int
	y int
}

type Area struct {
	x_start int
	x_end   int
	y_start int
	y_end   int
}

func getTargetArea(input string) Area {
	var target_area Area
	input = strings.Replace(input, "target area: ", "", -1)
	values := strings.Split(input, ", ")
	for _, i := range values {
		values_input := strings.Split(i, "=")
		if values_input[0] == "x" {
			values_input := strings.Split(values_input[1], "..")
			target_area.x_start, _ = strconv.Atoi(values_input[0])
			target_area.x_end, _ = strconv.Atoi(values_input[1])
		} else if values_input[0] == "y" {
			values_input := strings.Split(values_input[1], "..")
			target_area.y_end, _ = strconv.Atoi(values_input[0])
			target_area.y_start, _ = strconv.Atoi(values_input[1])
		}
	}
	return target_area
}

func calcVelocityChange(velocity *Vector) {
	if (*velocity).x > 0 {
		(*velocity).x--
	} else if (*velocity).x < 0 {
		(*velocity).x++
	}
	(*velocity).y--
}

func addVelocity(probe *Vector, velocity *Vector) {
	(*probe).x += velocity.x
	(*probe).y += velocity.y
}

func hasMissedTarget(probe *Vector, target *Area) bool {
	// If the probe is further than the target, stop
	if probe.x > target.x_end {
		return true
	} else if probe.y < target.y_end {
		return true
	}
	return false
}

func hasHitTarget(probe *Vector, target *Area) bool {
	// If the probe is further than the target, stop
	if probe.x >= target.x_start && probe.x <= target.x_end {
		if probe.y <= target.y_start && probe.y >= target.y_end {
			return true
		}
	}
	return false
}

func calcIfHit(velocity *Vector, target *Area) (bool, int) {
	probe := Vector{x: 0, y: 0}
	max_y := 0
	for !hasMissedTarget(&probe, target) {
		if max_y < probe.y {
			max_y = probe.y
		}
		if hasHitTarget(&probe, target) {
			return true, max_y

		}
		addVelocity(&probe, velocity)
		calcVelocityChange(velocity)
	}
	return false, 0
}

func ThreeSixtyNoScope(input string) (int, int) {
	var velocities []Vector
	target_area := getTargetArea(input)
	max_y := 0

	// Brute for
	for x := 0; x <= target_area.x_end; x++ {
		for y := target_area.y_end; y <= 500; y++ {
			vector := Vector{x: x, y: y}
			tmp_velocity := vector
			hit, max_y_probe := calcIfHit(&vector, &target_area)
			if hit {
				velocities = append(velocities, tmp_velocity)
				if max_y < max_y_probe {
					max_y = max_y_probe
				}
			}
		}
	}
	return max_y, len(velocities)
}

func main() {
	fmt.Println(ThreeSixtyNoScope(input))
}
