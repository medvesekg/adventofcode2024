package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	data := utils.ParseFile("input", utils.StrToInt)
	results := utils.ArrayMap(data, isSafe)

	countSafe := utils.ArrayCountOccurances(results)[true]

	fmt.Println(countSafe)
}

func partTwo() {
	data := utils.ParseFile("input", utils.StrToInt)
	results := utils.ArrayMap(data, func(row []int) bool {
		for _, variation := range variations(row) {
			if isSafe(variation) {
				return true
			}
		}
		return false
	})

	countSafe := utils.ArrayCountOccurances(results)[true]

	fmt.Println(countSafe)
}

func isSafe(row []int) bool {
	diff := row[0] - row[1]
	prevVal := row[1]

	if diff < 0 && diff > -4 {
		for _, val := range row[2:] {
			diff := prevVal - val

			if diff > -1 || diff < -3 {
				return false
			}
			prevVal = val
		}

	} else if diff > 0 && diff < 4 {
		for _, val := range row[2:] {
			diff := prevVal - val

			if diff > 3 || diff < 1 {
				return false
			}
			prevVal = val
		}
	} else {
		return false
	}

	return true
}

func variations(array []int) [][]int {
	results := [][]int{array}
	for i := range array {
		leftSide := array[:i]
		rightSide := array[i+1:]
		result := []int{}
		result = append(result, leftSide...)
		result = append(result, rightSide...)
		results = append(results, result)
	}
	return results
}
