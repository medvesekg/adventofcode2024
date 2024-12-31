package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	partOne()
}

func partOne() {
	input := utils.ReadFileLines("input")
	result := 0
	for _, line := range input {
		num := utils.StrToInt(line)
		for range 2000 {
			num = next(num)
		}
		result += num
	}
	fmt.Println(result)
}

func next(num int) int {
	num = prune(mix(num, num*64))
	num = prune(mix(num, num/32))
	num = prune((mix(num, num*2048)))

	return num
}

func mix(secret int, num int) int {
	return secret ^ num
}

func prune(num int) int {
	return num % 16777216
}
