package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
)

type Equation struct {
	test    int
	numbers []int
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println(computeResult(testEquation, "input"))
}

func partTwo() {
	fmt.Println(computeResult(testEquation2, "input"))
}

func computeResult(testFunction func(Equation) bool, path string) int {
	lines := parseInput(path)

	passed := utils.ArrayFilter(lines, testFunction)
	passedTestValues := utils.ArrayMap(passed, func(equation Equation) int {
		return equation.test
	})

	return utils.ArraySum(passedTestValues)
}

func parseInput(path string) []Equation {
	lines := utils.ReadFileLines(path)
	r, _ := regexp.Compile(`\d+`)

	return utils.ArrayMap(lines, func(line string) Equation {
		a := r.FindAllString(line, -1)
		b := utils.ArrayMap(a, utils.StrToInt)

		return Equation{b[0], b[1:]}
	})
}

func testEquation(equation Equation) bool {
	numbers := equation.numbers

	totals := []int{numbers[0]}

	for i := 1; i < len(numbers); i++ {
		newTotals := []int{}
		for _, total := range totals {
			newTotals = append(newTotals, total+numbers[i])
			newTotals = append(newTotals, total*numbers[i])
		}
		totals = newTotals
	}

	return utils.ArrayAny(totals, func(total int) bool {
		return total == equation.test
	})
}

func testEquation2(equation Equation) bool {
	numbers := equation.numbers

	totals := []int{numbers[0]}

	for i := 1; i < len(numbers); i++ {
		newTotals := []int{}
		for _, total := range totals {
			newTotals = append(newTotals, total+numbers[i])
			newTotals = append(newTotals, total*numbers[i])
			newTotals = append(newTotals, utils.StrToInt(utils.IntToStr(total)+utils.IntToStr(numbers[i])))
		}
		totals = newTotals
	}

	return utils.ArrayAny(totals, func(total int) bool {
		return total == equation.test
	})
}
