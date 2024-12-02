package main

import (
	"fmt"
	"math"
	"slices"

	"adventofcode/utils"
)

func parseInput(path string) ([]int, []int) {
	data := utils.ParseFile(path, utils.StrToInt)
	cols := utils.RowsToCols(data)

	return cols[0], cols[1]
}

func partOne() {
	leftCol, rightCol := parseInput("input")

	slices.Sort(leftCol)
	slices.Sort(rightCol)

	sum := 0

	for i := range leftCol {
		left := leftCol[i]
		right := rightCol[i]

		distance := math.Abs(float64(left) - float64(right))
		sum += int(distance)
	}

	fmt.Println(sum)
}

func partTwo() {
	leftCol, rightCol := parseInput("input")
	counts := utils.ArrayCountOccurances(rightCol)

	sum := 0
	for _, val := range leftCol {
		mul := counts[val]

		sum += val * mul
	}

	fmt.Println(sum)

}

func main() {
	partOne()
	partTwo()
}
