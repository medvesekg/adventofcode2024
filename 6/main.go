package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	data := parseInput("input")
	path := runSimulation(data)
	visited := utils.ArrayUnique(path)
	fmt.Println(len(visited))
}

func partTwo() {
	data := parseInput("input")
	count := 0
	for i := range data {
		for j := range data[i] {
			prev := data[i][j]
			data[i][j] = "#"
			path := runSimulation2(data)
			if !path {
				count++
			}
			data[i][j] = prev
		}
	}
	fmt.Println(count)

}

func parseInput(path string) [][]string {
	lines := utils.ReadFileLines(path)
	data := [][]string{}
	for _, line := range lines {
		data = append(data, strings.Split(line, ""))
	}
	return data
}

func getGuardPosition(grid [][]string) [2]int {
	for y, row := range grid {
		for x, cell := range row {
			if cell == "^" {
				return [2]int{y, x}
			}
		}
	}
	return [2]int{0, 0}
}

func runSimulation(grid [][]string) [][2]int {

	guardPosition := getGuardPosition(grid)

	currentY := guardPosition[0]
	currentX := guardPosition[1]
	path := [][2]int{{currentY, currentX}}
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	currentDirectionIndex := 0

	for {
		currentDirection := directions[currentDirectionIndex]
		currentY = currentY + currentDirection[0]
		currentX = currentX + currentDirection[1]

		if !utils.CheckBounds(currentY, currentX, grid) {
			break
		}

		if grid[currentY][currentX] == "#" {
			currentY = currentY - currentDirection[0]
			currentX = currentX - currentDirection[1]
			currentDirectionIndex = (currentDirectionIndex + 1) % 4
			continue
		}
		path = append(path, [2]int{currentY, currentX})
	}

	return path
}

func runSimulation2(grid [][]string) bool {

	guardPosition := getGuardPosition(grid)

	currentY := guardPosition[0]
	currentX := guardPosition[1]
	path := [][2]int{{currentY, currentX}}

	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	currentDirectionIndex := 0

	for {
		currentDirection := directions[currentDirectionIndex]
		currentY = currentY + currentDirection[0]
		currentX = currentX + currentDirection[1]

		if !utils.CheckBounds(currentY, currentX, grid) {
			break
		}

		if len(path) >= 1000 {
			return false
		}

		if grid[currentY][currentX] == "#" {
			currentY = currentY - currentDirection[0]
			currentX = currentX - currentDirection[1]
			currentDirectionIndex = (currentDirectionIndex + 1) % 4

			path = append(path, [2]int{currentY, currentX})
			continue
		}
	}

	return true
}
