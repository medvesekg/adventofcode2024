package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"

	"seehuhn.de/go/ncurses"
)

func main() {
	partOne()
	partTwo()

}

func interactive() {
	win := ncurses.Init()
	grid, instructions := parseInput("input")
	grid = scaleUpGrid(grid)
	robotPosition := utils.FindInGrid(grid, "@")
	fmt.Println(robotPosition, instructions)
	for {
		win.Print(utils.SprintGrid(grid))
		ch := win.GetCh()
		if ch == ncurses.KeyUp {
			robotPosition = moveRobot(robotPosition, "^", grid)
		} else if ch == ncurses.KeyDown {
			robotPosition = moveRobot(robotPosition, "v", grid)
		} else if ch == ncurses.KeyLeft {
			robotPosition = moveRobot(robotPosition, "<", grid)
		} else if ch == ncurses.KeyRight {
			robotPosition = moveRobot(robotPosition, ">", grid)
		}
		win.Erase()
	}
}

func partOne() {
	grid, instructions := parseInput("input")
	robotPosition := utils.FindInGrid(grid, "@")

	for _, instruction := range instructions {
		robotPosition = move1(robotPosition, instruction, grid)
	}

	fmt.Println(computeSum(grid))
}

func partTwo() {
	grid, instructions := parseInput("input")
	grid = scaleUpGrid(grid)
	robotPosition := utils.FindInGrid(grid, "@")

	for _, instruction := range instructions {
		robotPosition = moveRobot(robotPosition, instruction, grid)
	}
	fmt.Println(computeSum2(grid))
}

func moveRobot(origin utils.Point, direction string, grid [][]string) utils.Point {
	newPosition := getNewPosition(origin, direction)
	newPositionValue := grid[newPosition.Y][newPosition.X]

	switch newPositionValue {
	case ".":
		grid[origin.Y][origin.X] = "."
		grid[newPosition.Y][newPosition.X] = "@"
		return newPosition
	case "#":
		return origin
	case "[", "]":
		boxCoordinates := getBoxCoordinates(newPosition, grid)
		if moveBox(boxCoordinates, direction, grid, true) != boxCoordinates {
			return moveRobot(origin, direction, grid)
		} else {
			return origin
		}
	default:
		panic("OH NO")
	}
}

func moveBox(boxCoordinates [2]utils.Point, direction string, grid [][]string, move bool) [2]utils.Point {
	newPositionL := getNewPosition(boxCoordinates[0], direction)
	newPositionR := getNewPosition(boxCoordinates[1], direction)
	newPositionLValue := grid[newPositionL.Y][newPositionL.X]
	newPositionRValue := grid[newPositionR.Y][newPositionR.X]

	adjacentBoxes := map[[2]utils.Point]bool{}
	if (newPositionLValue == "[" || newPositionLValue == "]") && newPositionL != boxCoordinates[1] {
		adjacentBoxes[getBoxCoordinates(newPositionL, grid)] = true
	}
	if (newPositionRValue == "[" || newPositionRValue == "]") && newPositionR != boxCoordinates[0] {
		adjacentBoxes[getBoxCoordinates(newPositionR, grid)] = true
	}

	if newPositionLValue == "#" || newPositionRValue == "#" {
		return boxCoordinates
	} else if len(adjacentBoxes) > 0 {
		for adjacentBoxCoordinates := range adjacentBoxes {
			newAdjacentBoxCoordinates := moveBox(adjacentBoxCoordinates, direction, grid, false)

			if newAdjacentBoxCoordinates == adjacentBoxCoordinates {
				return boxCoordinates
			}
		}
		if move {
			for adjacentBoxCoordinates := range adjacentBoxes {
				moveBox(adjacentBoxCoordinates, direction, grid, true)
			}
			return moveBox(boxCoordinates, direction, grid, true)
		}
		return [2]utils.Point{newPositionL, newPositionR}
	} else if (newPositionLValue == "." || newPositionL == boxCoordinates[1]) && (newPositionRValue == "." || newPositionR == boxCoordinates[0]) {
		if move {
			grid[boxCoordinates[0].Y][boxCoordinates[0].X] = "."
			grid[boxCoordinates[1].Y][boxCoordinates[1].X] = "."
			grid[newPositionL.Y][newPositionL.X] = "["
			grid[newPositionR.Y][newPositionR.X] = "]"
		}
		return [2]utils.Point{newPositionL, newPositionR}
	}

	panic("OH NO")

}

func getBoxCoordinates(origin utils.Point, grid [][]string) [2]utils.Point {
	originValue := grid[origin.Y][origin.X]
	switch originValue {
	case "[":
		return [2]utils.Point{origin, {X: origin.X + 1, Y: origin.Y}}
	case "]":
		return [2]utils.Point{{X: origin.X - 1, Y: origin.Y}, origin}
	default:
		panic("MALFORMED BOX!!!")
	}
}

func move1(origin utils.Point, direction string, grid [][]string) utils.Point {
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
		wasMoved := move1(newPosition, direction, grid) != newPosition
		if wasMoved {
			return move1(origin, direction, grid)
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

func scaleUpGrid(grid [][]string) [][]string {
	newGrid := [][]string{}
	for _, row := range grid {
		newRow := []string{}
		for _, val := range row {
			switch val {
			case "#":
				newRow = append(newRow, "#")
				newRow = append(newRow, "#")
			case "O":
				newRow = append(newRow, "[")
				newRow = append(newRow, "]")
			case ".":
				newRow = append(newRow, ".")
				newRow = append(newRow, ".")
			case "@":
				newRow = append(newRow, "@")
				newRow = append(newRow, ".")
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
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

func computeSum2(grid [][]string) int {
	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "[" {
				sum += ((100 * y) + x)
			}
		}
	}
	return sum
}
