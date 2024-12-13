package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	data := utils.ReadFile("input")
	stones := utils.ArrayMap(strings.Fields(data), utils.StrToInt)
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	fmt.Println(len(stones))
}

func partTwo() {
	data := utils.ReadFile("input")
	stones := utils.ArrayMap(strings.Fields(data), utils.StrToInt)
	count := 0
	for _, stone := range stones {
		count += computeCount(stone, 75)
	}

	fmt.Println(count)
}

var computeCountCache = map[[2]int]int{}

func computeCount(num int, iterations int) int {
	var res int
	cached, isCached := computeCountCache[[2]int{num, iterations}]

	if isCached {
		return cached
	} else if iterations == 0 {
		res = 1
	} else if num == 0 {
		res = computeCount(1, iterations-1)
	} else if hasEvenNumberOfDigits(num) {
		left, right := splitStone(num)
		res = computeCount(left, iterations-1) + computeCount(right, iterations-1)
	} else {
		res = computeCount(num*2024, iterations-1)
	}
	computeCountCache[[2]int{num, iterations}] = res
	return res
}

func blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		stone := stones[i]
		if stone == 0 {
			stones[i] = 1
		} else if hasEvenNumberOfDigits(stone) {
			left, right := splitStone(stone)
			stones[i] = left
			i++
			stones = slices.Insert(stones, i, right)
		} else {
			stones[i] = stone * 2024
		}
	}
	return stones
}

func intToDigits(number int) []int {
	current := number
	digits := []int{}
	for current > 0 {
		digit := current % 10
		digits = append(digits, digit)
		current = current / 10
	}
	slices.Reverse(digits)
	return digits
}

func digitsToInt(digits []int) int {
	number := 0
	for i, digit := range digits {
		factor := int(math.Pow(10, float64(len(digits)-i-1)))
		number += digit * factor
	}
	return number
}

func hasEvenNumberOfDigits(number int) bool {
	digits := intToDigits(number)
	return len(digits)%2 == 0
}

func splitStone(stone int) (int, int) {
	digits := intToDigits(stone)

	half := len(digits) / 2
	leftHalf := digits[:half]
	rightHalf := digits[half:]

	return digitsToInt(leftHalf), digitsToInt(rightHalf)
}
