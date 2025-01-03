package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Instruction struct {
	in1 string
	op  string
	in2 string
	out string
}

func main() {
	state, instructions := parseInput("input")
	state = run(state, instructions)
	output := readOutput(state)
	fmt.Println(output)
}

func run(initialState map[string]int, instructions []Instruction) map[string]int {
	state := initialState
	for {
		stop := true
		for _, insturction := range instructions {
			var executed bool
			state, executed = executeInstruction(insturction, state)
			if !executed {
				stop = false
			}
		}
		if stop {
			break
		}
	}
	return state
}

func parseInput(path string) (map[string]int, []Instruction) {
	input := utils.ReadFile(path)
	parts := strings.Split(input, "\n\n")
	r := regexp.MustCompile(`(?m)^(\w+):\s*(\d)`)
	matches := r.FindAllStringSubmatch(parts[0], -1)
	initialState := map[string]int{}
	for _, match := range matches {
		initialState[match[1]] = utils.StrToInt(match[2])
	}

	r = regexp.MustCompile(`(?m)^(\w+)\s*((?:XOR)|(?:OR)|(?:AND))\s*(\w+)\s*->\s*(\w+)`)
	matches = r.FindAllStringSubmatch(parts[1], -1)
	instructions := []Instruction{}
	for _, match := range matches {
		instructions = append(instructions, Instruction{
			in1: match[1],
			op:  match[2],
			in2: match[3],
			out: match[4],
		})
	}
	return initialState, instructions
}

func executeInstruction(instruction Instruction, state map[string]int) (map[string]int, bool) {
	arg1, in1Exists := state[instruction.in1]
	arg2, in2Exists := state[instruction.in2]

	if !in1Exists || !in2Exists {
		return state, false
	}

	state[instruction.out] = executeOp(instruction.op, arg1, arg2)

	return state, true
}

func executeOp(op string, arg1 int, arg2 int) int {
	switch op {
	case "AND":
		return and(arg1, arg2)
	case "OR":
		return or(arg1, arg2)
	case "XOR":
		return xor(arg1, arg2)
	default:
		panic("Unknown op")
	}
}

func and(val1 int, val2 int) int {
	if val1 == 1 && val2 == 1 {
		return 1
	}
	return 0
}

func or(val1 int, val2 int) int {
	if val1 == 1 || val2 == 1 {
		return 1
	}
	return 0
}

func xor(val1 int, val2 int) int {
	if val1 == val2 {
		return 0
	}
	return 1
}

func readOutput(state map[string]int) int {
	i := 0
	output := []int{}
	for {
		value, exists := state[z(i)]
		if !exists {
			break
		}
		output = append(output, value)
		i++
	}
	result := 0
	for i := range output {
		result += int(math.Pow(2, float64(i))) * output[i]
	}
	return result
}

func z(i int) string {
	return fmt.Sprintf("z%02d", i)
}
