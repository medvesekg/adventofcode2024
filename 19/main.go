package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
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
