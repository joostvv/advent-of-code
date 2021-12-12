package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input []byte

//go:embed input-sample.txt
var sample []byte

//go:embed input-sample2.txt
var sample2 []byte

//go:embed input-sample3.txt
var sample3 []byte

// small caves, caves
func getPaths(data []byte, small_cave_twice bool) (map[string][]string, []string) {

	paths := make(map[string][]string)
	var small_caves []string
	list := strings.Split(string(data), "\n")
	for _, c := range list {
		cave := strings.Split(c, "-")
		paths[cave[0]] = append(paths[cave[0]], cave[1])
		if cave[1] != "end" && cave[0] != "start" {
			paths[cave[1]] = append(paths[cave[1]], cave[0])
		}
	}
	for path := range paths {
		if IsLower(path) && path != "start" {
			small_caves = append(small_caves, path)
		}
	}
	if small_cave_twice {
		small_caves = append(small_caves, "joker")
	}

	return paths, small_caves
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remove(v string, s []string) []string {
	var output []string
	for _, value := range s {
		if value != v {
			output = append(output, value)
		}
	}
	return output
}

func pathFinder(taken_path string, paths map[string][]string, small_caves []string, total_path []string) int {
	ways := 0
	for _, path := range paths[taken_path] {
		if path == "end" {
			// total_path := append(total_path, path)
			// fmt.Println("Final path:", total_path)
			ways += 1
			continue
		}
		if IsLower(path) && stringInSlice(path, small_caves) {
			new_small_caves := remove(path, small_caves)
			new_total_path := append(total_path, path)
			ways += pathFinder(path, paths, new_small_caves, new_total_path)
		} else if !IsLower(path) {
			new_total_path := append(total_path, path)
			ways += pathFinder(path, paths, small_caves, new_total_path)
		} else if IsLower(path) && stringInSlice("joker", small_caves) && path != "start" {
			new_small_caves := remove("joker", small_caves)
			new_total_path := append(total_path, path)
			ways += pathFinder(path, paths, new_small_caves, new_total_path)
		}
	}
	return ways
}

func SecretTunnel(data []byte, small_caves_twice bool) int {
	ways := 0
	var total_path []string
	total_path = append(total_path, "start")
	paths, small_caves := getPaths(data, small_caves_twice)
	ways += pathFinder("start", paths, small_caves, total_path)
	return ways
}

func main() {
	fmt.Println(SecretTunnel(input, false))
	fmt.Println(SecretTunnel(input, true))
}
