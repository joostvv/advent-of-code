package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTargetArea(t *testing.T) {
	test := map[Area]string{{x_start: 20, x_end: 30, y_start: -5, y_end: -10}: sample,
		{x_start: 265, x_end: 287, y_start: -58, y_end: -103}: input,
	}
	for target_area := range test {
		expected_area := getTargetArea(test[target_area])
		if !reflect.DeepEqual(expected_area, target_area) {
			t.Fail()
		}
	}
}

func TestCalcProbe(t *testing.T) {
	// Vector start maps to Vector end
	test := map[Vector]Vector{{x: 1, y: 30}: {x: 0, y: 29},
		{x: -1, y: 0}: {x: 0, y: -1},
		{x: 0, y: -1}: {x: 0, y: -2},
	}
	for start := range test {
		expected_result := test[start]
		// Calc one step with start_verlocity
		calcVelocityChange(&start)
		if !reflect.DeepEqual(expected_result, start) {
			fmt.Println(expected_result, start)
			t.Fail()
		}
	}
}

func TestCalcIfHit(t *testing.T) {
	// Vector start maps to Vector end
	test := map[Vector]bool{{x: 7, y: 2}: true,
		{x: 6, y: 3}:   true,
		{x: 9, y: 0}:   true,
		{x: 17, y: -4}: false,
	}
	target_area := getTargetArea(sample)
	for vector := range test {
		expected_result := test[vector]
		hit, _ := calcIfHit(&vector, &target_area)
		if hit != expected_result {
			fmt.Println(vector, hit, expected_result)
			t.Fail()
		}
	}
}
