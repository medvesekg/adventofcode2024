package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

var keypad = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"", "0", "A"},
}

var remote = [][]string{
	{"", "^", "A"},
	{"<", "v", ">"},
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	codes := parseInput("input")
	r, _ := regexp.Compile(`\d+`)

	result := 0
	for _, code := range codes {
		splitCode := strings.Split(code, "")
		codeValue := utils.StrToInt(r.FindString(code))
		result += computeInput(splitCode, 2) * codeValue
	}

	fmt.Println(result)
}

func partTwo() {
	codes := parseInput("input")
	r, _ := regexp.Compile(`\d+`)

	result := 0
	for _, code := range codes {
		splitCode := strings.Split(code, "")
		codeValue := utils.StrToInt(r.FindString(code))
		result += computeInput(splitCode, 25) * codeValue
	}

	fmt.Println(result)
}

func parseInput(path string) []string {
	return utils.ReadFileLines(path)
}

func combos(array []string, prefix []string) [][]string {
	if len(array) == 0 {
		return [][]string{prefix}
	}
	c := [][]string{}
	for _, item := range unique(array) {
		newPrefix := []string{}
		newPrefix = append(newPrefix, prefix...)
		newPrefix = append(newPrefix, item)
		index := findIndex(array, item)
		newArray := []string{}
		newArray = append(newArray, array[:index]...)
		newArray = append(newArray, array[index+1:]...)
		c = append(c, combos(newArray, newPrefix)...)
	}
	return c
}

func unique(array []string) []string {
	seen := map[string]bool{}
	unique := []string{}
	for _, item := range array {
		_, alreadySeen := seen[item]
		if !alreadySeen {
			unique = append(unique, item)
		}
		seen[item] = true
	}
	return unique
}

func findIndex(array []string, item string) int {
	for i, a := range array {
		if a == item {
			return i
		}
	}
	return -1
}

func checkPresses(presses []string, currentPosition utils.Point, forbidden utils.Point) bool {
	for _, press := range presses {
		switch press {
		case ">":
			currentPosition.X++
		case "<":
			currentPosition.X--
		case "v":
			currentPosition.Y++
		case "^":
			currentPosition.Y--
		}
		if currentPosition == forbidden {
			return false
		}
	}
	return true
}

func computeInput(code []string, depth int) int {
	finalCode := 0
	currentPosition := utils.FindInGrid(keypad, "A")
	for _, char := range code {
		target := utils.FindInGrid(keypad, char)
		shortest := math.MaxInt64
		for _, possiblity := range findWays(currentPosition, target, keypad) {
			length := computeInputRemote(possiblity, depth)
			if length < shortest {
				shortest = length
			}
		}
		finalCode += shortest
		currentPosition = target
	}
	return finalCode
}

var memo = map[string]int{}

func computeInputRemote(code []string, depth int) int {
	memoKey := fmt.Sprintf("%d-%s", depth, strings.Join(code, ""))

	result, exists := memo[memoKey]
	if exists {
		return result
	}

	if depth == 0 {
		return len(code)
	}
	finalCode := 0
	currentPosition := utils.FindInGrid(remote, "A")
	for _, char := range code {
		target := utils.FindInGrid(remote, char)
		shortest := math.MaxInt64
		for _, possiblity := range findWays(currentPosition, target, remote) {
			length := computeInputRemote(possiblity, depth-1)
			if length < shortest {
				shortest = length
			}
		}
		finalCode += shortest
		currentPosition = target
	}

	memo[memoKey] = finalCode

	return finalCode
}

func findWays(origin utils.Point, target utils.Point, grid [][]string) [][]string {
	diffX := target.X - origin.X
	diffY := target.Y - origin.Y

	var directionX string
	if diffX > 0 {
		directionX = ">"
	} else if diffX < 0 {
		directionX = "<"
	}

	var directionY string
	if diffY > 0 {
		directionY = "v"
	} else if diffY < 0 {
		directionY = "^"
	}

	buttons := []string{}

	for range int(math.Abs(float64(diffX))) {
		buttons = append(buttons, directionX)
	}
	for range int(math.Abs(float64(diffY))) {
		buttons = append(buttons, directionY)
	}

	combinations := combos(buttons, []string{})
	validCombinations := [][]string{}
	forbidden := utils.FindInGrid(grid, "")
	for _, combination := range combinations {
		if checkPresses(combination, origin, forbidden) {
			validCombinations = append(validCombinations, append(combination, "A"))
		}
	}
	return validCombinations
}
