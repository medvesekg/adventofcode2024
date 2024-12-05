package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func parseOrdedringRules(raw string) map[int]map[int]bool {
	orderingRules := map[int]map[int]bool{}
	for _, line := range strings.Split(raw, "\n") {
		parts := strings.Split(line, "|")
		first := utils.StrToInt(parts[0])
		_, exists := orderingRules[first]
		if !exists {
			orderingRules[first] = map[int]bool{}
		}
		orderingRules[first][utils.StrToInt(parts[1])] = true
	}
	return orderingRules
}

func parseUpdates(raw string) [][]int {
	lines := strings.Split(raw, "\n")
	return utils.ArrayMap(lines, func(line string) []int {
		return utils.ArrayMap(strings.Split(line, ","), utils.StrToInt)
	})
}

func parseInput(path string) (map[int]map[int]bool, [][]int) {
	raw := utils.ReadFile(path)
	parts := strings.Split(raw, "\n\n")
	return parseOrdedringRules(parts[0]), parseUpdates(parts[1])
}

func orderIsCorrect(row []int, orderingRules map[int]map[int]bool) bool {
	numsBefore := map[int]bool{}
	for _, num := range row {
		mustBeBefore, rulesExist := orderingRules[num]
		if rulesExist {
			for mustmustBeBeforeNum := range mustBeBefore {
				if numsBefore[mustmustBeBeforeNum] {
					return false
				}
			}
		}
		numsBefore[num] = true
	}

	return true
}

func sumResult(rows [][]int) int {
	middleNumbers := utils.ArrayMap(rows, func(row []int) int {
		return row[len(row)/2]
	})
	return utils.ArraySum(middleNumbers)
}

func partOne() {
	orderingRules, updates := parseInput("input")
	correctUpdates := utils.ArrayFilter(updates, func(row []int) bool {
		return orderIsCorrect(row, orderingRules)
	})

	sum := sumResult(correctUpdates)

	fmt.Println(sum)

}

func partTwo() {
	orderingRules, updates := parseInput("input")
	incorrectUpdates := utils.ArrayFilter(updates, func(row []int) bool {
		return !orderIsCorrect(row, orderingRules)
	})

	for _, row := range incorrectUpdates {
		sort.SliceStable(row, func(i, j int) bool {
			a := row[i]
			b := row[j]

			aMustBeBefore := orderingRules[a][b]
			if aMustBeBefore {
				return true
			}
			bMustBeBefore := orderingRules[b][a]
			if bMustBeBefore {
				return false
			}
			return false
		})
	}

	sum := sumResult(incorrectUpdates)
	fmt.Println(sum)

}
