package main

// 5457\d\d\d\d7553\d\d\d\d1657\d\d\d\d3753\d\d\d\d5057\d\d\d\d7153\d\d\d\d1257\d\d\d\d3353\d\d\d\d

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
)

type Computer struct {
	regA       int
	regB       int
	regC       int
	output     []int
	pointer    int
	program    []int
	terminated bool
}

func (c Computer) getOutput() string {
	return join(c.output)
}

func main() {
	testComputer()
	partOne()
	partTwo()
}

func partOne() {
	computer := parseInput("input")
	computer = run(computer)
	fmt.Println(computer.getOutput())
}

func partTwo() {
	initializer := parseInput2("input")
	base := 8

	program := initializer().program
	regA := int(math.Pow(float64(base), float64(len(program)-1)))

	for i := len(program) - 1; i >= 0; i-- {
		computer := initializer()
		n := int(math.Pow(float64(base), float64(i)))
		computer.regA = regA
		computer = run(computer)
		for computer.output[i] != computer.program[i] {
			computer = initializer()
			regA += n
			computer.regA = regA
			computer = run(computer)
		}

		for j := len(program) - 1; j >= i; j-- {
			if computer.output[j] != computer.program[j] {
				i = j + 1
			}
		}
	}

	fmt.Println(regA)
}

func testComputer() {
	var computer Computer

	computer = Computer{}
	computer.regC = 9
	computer.program = []int{2, 6}
	computer = tick(computer)
	if computer.regB != 1 {
		panic("Test failed")
	}

	computer = Computer{}
	computer.regA = 10
	computer.program = []int{5, 0, 5, 1, 5, 4}
	computer = run(computer)
	if computer.getOutput() != "0,1,2" {
		panic("Test failed")
	}

	computer = Computer{}
	computer.regA = 2024
	computer.program = []int{0, 1, 5, 4, 3, 0}
	computer = run(computer)
	if computer.getOutput() != "4,2,5,6,7,7,7,7,3,1,0" {
		panic("Test failed")
	}
	if computer.regA != 0 {
		panic("Test failed")
	}

	computer = Computer{}
	computer.regB = 29
	computer.program = []int{1, 7}
	computer = tick(computer)
	if computer.regB != 26 {
		panic("Test failed")
	}

	computer = Computer{}
	computer.regB = 2024
	computer.regC = 43690
	computer.program = []int{4, 0}
	computer = tick(computer)

	if computer.regB != 44354 {
		panic("Test failed")
	}

}

func parseInput(path string) Computer {
	lines := utils.ReadFileLines(path)
	r, _ := regexp.Compile(`\d+`)

	return Computer{
		regA:       utils.StrToInt(r.FindString(lines[0])),
		regB:       utils.StrToInt(r.FindString(lines[1])),
		regC:       utils.StrToInt(r.FindString(lines[2])),
		pointer:    0,
		terminated: false,
		output:     []int{},
		program:    utils.ArrayMap(r.FindAllString(lines[3], -1), utils.StrToInt),
	}
}

func parseInput2(path string) func() Computer {
	lines := utils.ReadFileLines(path)
	r, _ := regexp.Compile(`\d+`)
	regA := utils.StrToInt(r.FindString(lines[0]))
	regB := utils.StrToInt(r.FindString(lines[1]))
	regC := utils.StrToInt(r.FindString(lines[2]))
	program := utils.ArrayMap(r.FindAllString(lines[3], -1), utils.StrToInt)
	return func() Computer {
		return Computer{
			regA:       regA,
			regB:       regB,
			regC:       regC,
			pointer:    0,
			terminated: false,
			output:     []int{},
			program:    program,
		}
	}
}

func run(computer Computer) Computer {
	for !computer.terminated {
		computer = tick(computer)
	}
	return computer
}

func tick(computer Computer) Computer {
	if computer.pointer > len(computer.program)-2 {
		computer.terminated = true
		return computer
	}

	command := computer.program[computer.pointer]
	operand := computer.program[computer.pointer+1]

	return instructions[command](operand, computer)

}

func comboOperand(o int, c Computer) int {
	switch {
	case o < 4:
		return o
	case o == 4:
		return c.regA
	case o == 5:
		return c.regB
	case o == 6:
		return c.regC
	}
	panic("NO!!!!")
}

var instructions = []func(int, Computer) Computer{
	// 0 adv
	func(o int, c Computer) Computer {
		o = comboOperand(o, c)
		numerator := c.regA
		denominator := int(math.Pow(float64(2), float64(o)))
		result := numerator / denominator
		c.regA = result
		c.pointer += 2
		return c
	},
	// 1 bxl
	func(o int, c Computer) Computer {
		c.regB = c.regB ^ o
		c.pointer += 2
		return c
	},
	// 2 bst
	func(o int, c Computer) Computer {
		o = comboOperand(o, c)
		c.regB = o % 8
		c.pointer += 2
		return c
	},
	// 3 jnz
	func(o int, c Computer) Computer {
		if c.regA == 0 {
			c.pointer += 2
			return c
		}
		c.pointer = o
		return c
	},
	// 4 bxc
	func(o int, c Computer) Computer {
		c.regB = c.regB ^ c.regC
		c.pointer += 2
		return c
	},
	// 5 out
	func(o int, c Computer) Computer {
		o = comboOperand(o, c)
		c.output = append(c.output, o%8)
		c.pointer += 2
		return c
	},
	// 6 bdv
	func(o int, c Computer) Computer {
		o = comboOperand(o, c)
		numerator := c.regA
		denominator := int(math.Pow(float64(2), float64(o)))
		result := numerator / denominator
		c.regB = result
		c.pointer += 2
		return c
	},
	// 6 cdv
	func(o int, c Computer) Computer {
		o = comboOperand(o, c)
		numerator := c.regA
		denominator := int(math.Pow(float64(2), float64(o)))
		result := numerator / denominator
		c.regC = result
		c.pointer += 2
		return c
	},
}

func join(array []int) string {
	var joined string
	for i, o := range array {
		joined += utils.IntToStr(o)
		if i < len(array)-1 {
			joined += ","
		}
	}
	return joined
}
