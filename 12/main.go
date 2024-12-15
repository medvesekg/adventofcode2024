package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	grid := utils.ParseFileGrid("input")
	done := map[utils.Point]bool{}
	totalPrice := 0
	for y := range grid {
		for x := range grid[y] {
			point := utils.Point{X: x, Y: y}
			_, alreadyDone := done[point]
			if !alreadyDone {
				area, _ := mapArea(grid, utils.Point{X: x, Y: y})
				sides := countSides(area)
				totalPrice += len(area) * sides
				for _, pt := range area {
					done[pt] = true
				}
			}
		}
	}

	fmt.Println(totalPrice)
}

func partOne() {
	grid := utils.ParseFileGrid("input")
	done := map[utils.Point]bool{}
	totalPrice := 0
	for y := range grid {
		for x := range grid[y] {
			point := utils.Point{X: x, Y: y}
			_, alreadyDone := done[point]
			if !alreadyDone {
				area, fence := mapArea(grid, utils.Point{X: x, Y: y})
				totalPrice += len(area) * fence

				for _, pt := range area {
					done[pt] = true
				}

			}
		}
	}

	fmt.Println(totalPrice)
}

func mapArea(grid [][]string, origin utils.Point) ([]utils.Point, int) {
	area := []utils.Point{origin}
	checked := map[utils.Point]bool{}
	toCheck := []utils.Point{origin}
	val := grid[origin.Y][origin.X]
	fence := 0
	for len(toCheck) > 0 {
		current := toCheck[len(toCheck)-1]
		toCheck = toCheck[:len(toCheck)-1]
		checked[current] = true
		fence += 4

		for _, neighbour := range utils.GridGetNeighbours(grid, current, utils.DIRECTIONS2["CARDINAL"]) {
			neighbourVal := grid[neighbour.Y][neighbour.X]
			_, alreadyChecked := checked[neighbour]
			checked[neighbour] = true

			if neighbourVal == val {
				fence--
			}

			if neighbourVal == val && !alreadyChecked {
				area = append(area, neighbour)
				toCheck = append(toCheck, neighbour)
			}
		}
	}
	return area, fence
}

func countSides(area []utils.Point) int {
	maxY := utils.ArrayMax(area, func(p utils.Point) int {
		return p.Y
	})
	maxX := utils.ArrayMax(area, func(p utils.Point) int {
		return p.X
	})
	minY := utils.ArrayMin(area, func(p utils.Point) int {
		return p.Y
	})
	minX := utils.ArrayMin(area, func(p utils.Point) int {
		return p.X
	})

	lines := 0
	for y := minY; y <= maxY; y++ {
		lineFound := false
		for x := minX; x <= maxX; x++ {
			if (slices.Contains(area, utils.Point{X: x, Y: y}) && !slices.Contains(area, utils.Point{X: x, Y: y - 1})) {
				lineFound = true
			} else {
				if lineFound {
					lines++
					lineFound = false
				}
			}
		}

		if lineFound {
			lines++
			lineFound = false
		}

		lineFound = false
		for x := minX; x <= maxX; x++ {
			if (slices.Contains(area, utils.Point{X: x, Y: y}) && !slices.Contains(area, utils.Point{X: x, Y: y + 1})) {
				lineFound = true
			} else {
				if lineFound {
					lines++
					lineFound = false
				}
			}
		}

		if lineFound {
			lines++
			lineFound = false
		}
	}

	for x := minX; x <= maxX; x++ {
		lineFound := false
		for y := minY; y <= maxY; y++ {
			if (slices.Contains(area, utils.Point{X: x, Y: y}) && !slices.Contains(area, utils.Point{X: x + 1, Y: y})) {
				lineFound = true
			} else {
				if lineFound {
					lines++
					lineFound = false
				}
			}
		}

		if lineFound {
			lines++
			lineFound = false
		}

		lineFound = false
		for y := minY; y <= maxY; y++ {
			if (slices.Contains(area, utils.Point{X: x, Y: y}) && !slices.Contains(area, utils.Point{X: x - 1, Y: y})) {
				lineFound = true
			} else {
				if lineFound {
					lines++
					lineFound = false
				}
			}
		}

		if lineFound {
			lines++
			lineFound = false
		}
	}

	return lines
}
