package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	//partOne()
	partTwo()
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

func lastDigit(num int) int {
	return num % 10
}

func partTwo() {
	sequences := parseSequences("input")

	analyzedSequences := map[[4]int]int{}
	for _, sequence := range sequences {
		subsequences := map[[4]int]int{}
		for i := 1; i+3 < len(sequence); i++ {
			s := [4]int{}
			for j := 0; j < 4; j++ {
				s[j] = sequence[i+j][1]
			}
			_, alreadyFound := subsequences[s]
			if alreadyFound {
				continue
			}
			value := sequence[i+3][0]

			subsequences[s] = value
		}

		for s, v := range subsequences {
			analyzedSequences[s] += v
		}
	}

	longest := 0
	//longestSequence := [4]int{}

	for _, v := range analyzedSequences {
		if v > longest {
			longest = v
			//longestSequence = s
		}
	}

	fmt.Println(longest)
}

func parseSequences(path string) map[int][][2]int {
	input := utils.ReadFileLines(path)
	sequences := map[int][][2]int{}
	for _, line := range input {
		startNum := utils.StrToInt(line)
		sequence := [][2]int{{lastDigit(startNum), 0}}
		num := startNum
		for range 2000 {
			num = next(num)
			previous := sequence[len(sequence)-1]
			current := lastDigit(num)
			sequence = append(sequence, [2]int{current, current - previous[0]})
		}
		sequences[startNum] = sequence
	}
	return sequences
}
