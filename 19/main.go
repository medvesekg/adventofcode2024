package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	towels, patterns := parseInput("input")

	countPossible := 0
	for _, pattern := range patterns {
		if isPossible(towels, pattern) {
			countPossible++
		}
	}
	fmt.Println(countPossible)
}

func parseInput(path string) ([]string, []string) {
	lines := utils.ReadFileLines(path)

	r, _ := regexp.Compile(`\w+`)
	towels := r.FindAllString(lines[0], -1)
	patterns := lines[1:]
	return towels, patterns
}

var cache = make(map[string]bool)

func isPossible(towels []string, pattern string) bool {

	cachedResult, cacheExists := cache[pattern]
	if cacheExists {
		return cachedResult
	}

	if pattern == "" {
		return true
	}

	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) && isPossible(towels, strings.TrimPrefix(pattern, towel)) {
			cache[pattern] = true
			return true
		}
	}
	cache[pattern] = false
	return false
}

var cache2 = make(map[string]int)

func countPossibilities(towels []string, pattern string) int {
	count := 0

	cachedResult, cacheExists := cache2[pattern]
	if cacheExists {
		return cachedResult
	}

	if pattern == "" {
		return 1
	}

	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			count += countPossibilities(towels, strings.TrimPrefix(pattern, towel))
		}
	}
	cache2[pattern] = count
	return count
}

func partTwo() {
	towels, patterns := parseInput("input")

	possibilities := 0
	for _, pattern := range patterns {
		possibilities += countPossibilities(towels, pattern)
	}
	fmt.Println(possibilities)
}
