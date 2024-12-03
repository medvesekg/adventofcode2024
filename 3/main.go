package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	inst := utils.ReadFile("input")
	r, err := regexp.Compile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	utils.CheckError(err)
	res := r.FindAllString(inst, -1)
	sum := exec(res)
	fmt.Println(sum)
}

func partOne() {
	inst := utils.ReadFile("input")

	regex, err := regexp.Compile(`mul\(\d+,\d+\)`)

	utils.CheckError(err)

	res := regex.FindAllString(inst, -1)

	muls := utils.ArrayMap(res, func(inst string) int {
		return execMul(inst)
	})

	sum := utils.ArraySum(muls)

	fmt.Println(sum)
}

func exec(cmds []string) int {

	enabled := true
	vals := []int{}

	for _, cmd := range cmds {
		switch cmd {
		case "do()":
			enabled = true

		case "don't()":
			enabled = false
		default:
			if enabled {
				vals = append(vals, execMul(cmd))
			}
		}
	}

	return utils.ArraySum(vals)
}

func execMul(cmd string) int {
	r, err := regexp.Compile(`\d+`)
	utils.CheckError(err)
	nums := r.FindAllString(cmd, -1)

	return utils.StrToInt(nums[0]) * utils.StrToInt(nums[1])
}
