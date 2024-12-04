package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func parseInput(input string) [][]string {
	data := utils.ReadFileLines(input)
	rows := [][]string{}
	for _, line := range data {
		row := strings.Split(line, "")
		rows = append(rows, row)
	}
	return rows
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	pattern := []string{"X", "M", "A", "S"}
	directions := [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	xmasCount := 0
	rows := parseInput("input")

	for y := range rows {
		row := rows[y]
		for x := range row {
			for _, direction := range directions {
				check := checkDirection(rows, y, x, pattern, direction)
				if check {
					xmasCount++
				}

			}
		}
	}

	fmt.Println(xmasCount)
}

func partTwo() {
	xmasCount := 0
	rows := parseInput("input")
	for y := range rows {
		for x := range rows[y] {
			if checkXmas(rows, y, x) {
				xmasCount++
			}
		}
	}
	fmt.Println(xmasCount)

}

func checkXmas(rows [][]string, y int, x int) bool {
	if rows[y][x] != "A" {
		return false
	}

	xFound := 0
	for _, direction := range [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}} {
		nextY := y + direction[0]
		nextX := x + direction[1]

		if !checkBounds(nextY, nextX, rows) {
			continue
		}

		nextLetter := rows[nextY][nextX]

		if nextLetter == "M" {
			oppositeY := y - direction[0]
			oppositeX := x - direction[1]

			if !checkBounds(oppositeY, oppositeX, rows) {
				continue
			}

			oppositeLetter := rows[oppositeY][oppositeX]

			if oppositeLetter == "S" {
				xFound++
			}
		}
	}
	return xFound == 2
}

func checkDirection(rows [][]string, y int, x int, pattern []string, direction []int) bool {
	curY := y
	curX := x
	for _, patternLetter := range pattern {

		if !checkBounds(curY, curX, rows) {
			return false
		}

		if rows[curY][curX] != patternLetter {
			return false
		}

		curY += direction[0]
		curX += direction[1]

	}

	return true
}

func checkBounds[T any](y int, x int, grid [][]T) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y])

}
