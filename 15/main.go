package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	grid, instructions := parseInput("input")
	robotPosition := utils.FindInGrid(grid, "@")

	for _, instruction := range instructions {
		robotPosition = move(robotPosition, instruction, grid)
	}

	fmt.Println(computeSum(grid))
}

func move(origin utils.Point, direction string, grid [][]string) utils.Point {
	originValue := grid[origin.Y][origin.X]
	newPosition := getNewPosition(origin, direction)
	newPositionValue := grid[newPosition.Y][newPosition.X]

	if newPositionValue == "." {
		grid[origin.Y][origin.X] = "."
		grid[newPosition.Y][newPosition.X] = originValue
		return newPosition
	} else if newPositionValue == "#" {
		return origin
	} else if newPositionValue == "O" {
		wasMoved := move(newPosition, direction, grid) != newPosition
		if wasMoved {
			return move(origin, direction, grid)
		}
		return origin
	}

	panic("SOMETHING IS WRONG!")
}

func parseInput(path string) ([][]string, []string) {
	data := utils.ReadFile(path)
	parts := strings.Split(data, "\n\n")
	grid := utils.ParseGrid(parts[0])
	instructions := strings.Split(parts[1], "")
	return grid, instructions
}

func getNewPosition(origin utils.Point, direction string) utils.Point {
	switch direction {
	case "<":
		return utils.Point{X: origin.X - 1, Y: origin.Y}
	case ">":
		return utils.Point{X: origin.X + 1, Y: origin.Y}
	case "^":
		return utils.Point{X: origin.X, Y: origin.Y - 1}
	case "v":
		return utils.Point{X: origin.X, Y: origin.Y + 1}
	}
	fmt.Println(direction)
	panic("Invalid direction")
}

func computeSum(grid [][]string) int {
	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" {
				sum += ((100 * y) + x)
			}
		}
	}
	return sum
}
