package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

var WIDTH = 5
var HEIGHT = 5

func main() {
	partOne()
}

func partOne() {
	locks, keys := parseInput("input")

	numCompatible := 0
	for _, key := range keys {
		for _, lock := range locks {
			if isCompatible(key, lock) {
				numCompatible++
			}
		}
	}
	fmt.Println(numCompatible)
}

func parseInput(path string) ([][5]int, [][5]int) {
	input := utils.ReadFile(path)
	splitInput := strings.Split(input, "\n\n")
	locks := [][5]int{}
	keys := [][5]int{}
	for _, part := range splitInput {
		lines := strings.Split(part, "\n")
		if isLock(lines) {
			locks = append(locks, parseLock(lines))
		} else {
			keys = append(keys, parseKey(lines))
		}
	}
	return locks, keys
}

func isLock(lines []string) bool {
	return string(lines[0][0]) == "#"
}

func parseLock(lines []string) [5]int {
	lock := [5]int{}
	for i := 0; i < WIDTH; i++ {
		num := 0
		for j := 1; j <= HEIGHT; j++ {
			if string(lines[j][i]) == "#" {
				num++
			}
		}
		lock[i] = num
	}
	return lock
}

func parseKey(lines []string) [5]int {
	keys := [5]int{}
	for i := 0; i < WIDTH; i++ {
		num := 0
		for j := HEIGHT; j >= 0; j-- {
			if string(lines[j][i]) == "#" {
				num++
			}
		}
		keys[i] = num
	}
	return keys
}

func isCompatible(key [5]int, lock [5]int) bool {
	for i := 0; i < WIDTH; i++ {
		if key[i]+lock[i] > HEIGHT {
			return false
		}
	}
	return true
}
