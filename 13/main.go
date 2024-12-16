package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Machine struct {
	a     utils.Point
	b     utils.Point
	prize utils.Point
}

func (m Machine) String() string {
	return fmt.Sprintf(
		" A: x%d y%d\n B: x%d y%d\n Prize: x:%d y:%d\n\n",
		m.a.X,
		m.a.Y,
		m.b.X,
		m.b.Y,
		m.prize.X,
		m.prize.Y,
	)
}

func main() {
	partOne()
	//partTwo()
}

func partOne() {
	machines := parseInput("input")
	total := 0
	for _, machine := range machines {
		solved, a, b := solveMachine(machine)
		if solved {
			total += 3*a + b
		}
	}
	fmt.Println(total)
}

func partTwo() {
	machines := parseInput2("input")
	total := 0
	for _, machine := range machines {
		solveMachine2(machine)
		break
	}
	fmt.Println(total)
}

func solveMachine(machine Machine) (bool, int, int) {
	db1 := machine.prize.X / machine.b.X
	db2 := machine.prize.Y / machine.b.Y

	db := int(math.Min(float64(db1), float64(db2)))
	i := 0

	for db > 0 {
		tryX := machine.b.X*db + machine.a.X*i
		tryY := machine.b.Y*db + machine.a.Y*i

		if tryX == machine.prize.X && tryY == machine.prize.Y {
			return true, i, db
		} else if tryX > machine.prize.X || tryY > machine.prize.Y {
			db--
		} else if tryX < machine.prize.X || tryY < machine.prize.Y {
			i++
		}

	}

	return false, 0, 0
}

func solveMachine2(machine Machine) {
	div := machine.prize.X / machine.b.X
	mul := machine.b.X * div
	diff := machine.prize.X - mul

	fmt.Println(diff)

}

func parseInput(path string) []Machine {
	data := utils.ReadFile(path)
	lines := strings.Split(data, "\n")

	machines := utils.ArraySplit(lines, func(s string) bool {
		return s == ""
	})

	r, _ := regexp.Compile(`X([+-]\d+), Y([+-]\d+)`)
	r2, _ := regexp.Compile(`X=(\d+), Y=(\d+)`)
	m := []Machine{}
	for _, machine := range machines {

		btnALine := r.FindStringSubmatch(machine[0])
		btnAPoint := utils.Point{X: toInt(btnALine[1]), Y: toInt(btnALine[2])}

		btnBLine := r.FindStringSubmatch(machine[1])
		btnBPoint := utils.Point{X: toInt(btnBLine[1]), Y: toInt(btnBLine[2])}

		prizeLine := r2.FindStringSubmatch(machine[2])
		prizePoint := utils.Point{X: utils.StrToInt(prizeLine[1]), Y: utils.StrToInt(prizeLine[2])}

		m = append(m, Machine{
			a:     btnAPoint,
			b:     btnBPoint,
			prize: prizePoint,
		})
	}

	return m
}

func parseInput2(path string) []Machine {
	data := utils.ReadFile(path)
	lines := strings.Split(data, "\n")

	machines := utils.ArraySplit(lines, func(s string) bool {
		return s == ""
	})

	r, _ := regexp.Compile(`X([+-]\d+), Y([+-]\d+)`)
	r2, _ := regexp.Compile(`X=(\d+), Y=(\d+)`)
	m := []Machine{}
	for _, machine := range machines {

		btnALine := r.FindStringSubmatch(machine[0])
		btnAPoint := utils.Point{X: toInt(btnALine[1]), Y: toInt(btnALine[2])}

		btnBLine := r.FindStringSubmatch(machine[1])
		btnBPoint := utils.Point{X: toInt(btnBLine[1]), Y: toInt(btnBLine[2])}

		prizeLine := r2.FindStringSubmatch(machine[2])
		prizePoint := utils.Point{X: utils.StrToInt(prizeLine[1]) + 10000000000000, Y: utils.StrToInt(prizeLine[2]) + 10000000000000}

		m = append(m, Machine{
			a:     btnAPoint,
			b:     btnBPoint,
			prize: prizePoint,
		})
	}

	return m
}

func toInt(s string) int {
	if string(s[0]) == "+" {
		return utils.StrToInt(s)
	}
	return -utils.StrToInt(s)
}
