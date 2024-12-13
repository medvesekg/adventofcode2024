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

func partOne() {
	grid := utils.ParseFileGridInt("input")
	trailheads := findTrailheads(grid)
	sum := 0
	for _, trailhead := range trailheads {
		sum += calcualteTrailheadScore(grid, trailhead)
	}
	fmt.Println(sum)
}

func partTwo() {
	grid := utils.ParseFileGridInt("input")
	trailheads := findTrailheads(grid)
	sum := 0
	for _, trailhead := range trailheads {
		sum += calcualteTrailheadRating(grid, trailhead)
	}
	fmt.Println(sum)
}

func findTrailheads(grid [][]int) [][2]int {
	trailheads := [][2]int{}
	for y := range grid {
		row := grid[y]
		for x := range row {
			val := row[x]
			if val == 0 {
				trailheads = append(trailheads, [2]int{y, x})
			}
		}
	}
	return trailheads
}

func calcualteTrailheadScore(grid [][]int, trailhead [2]int) int {
	score := 0
	frontier := [][2]int{trailhead}
	visited := map[[2]int]bool{}
	for len(frontier) > 0 {
		current := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		if visited[current] {
			continue
		}
		visited[current] = true

		if grid[current[0]][current[1]] == 9 {
			score++
			continue
		}

		successors := getSuccessors(grid, current)

		frontier = append(frontier, successors...)
	}

	return score
}

func getSuccessors(grid [][]int, origin [2]int) [][2]int {
	successors := [][2]int{}

	for _, direction := range utils.DIRECTIONS["CARDINAL"] {
		possibleSuccessor := [2]int{origin[0] + direction[0], origin[1] + direction[1]}

		if utils.CheckBounds(possibleSuccessor[0], possibleSuccessor[1], grid) {
			successorValue := grid[possibleSuccessor[0]][possibleSuccessor[1]]
			originValue := grid[origin[0]][origin[1]]
			if successorValue == originValue+1 {
				successors = append(successors, possibleSuccessor)
			}

		}
	}

	return successors
}

func calcualteTrailheadRating(grid [][]int, trailhead [2]int) int {
	rating := 0
	frontier := [][][2]int{{trailhead}}
	for len(frontier) > 0 {
		currentPath := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]
		last := currentPath[len(currentPath)-1]

		successors := getSuccessors(grid, last)

		for _, successor := range successors {
			successorValue := grid[successor[0]][successor[1]]
			if successorValue == 9 {
				rating++
				continue
			}
			if slices.Contains(currentPath, successor) {
				continue
			}
			currentPath = append(currentPath, successor)
			frontier = append(frontier, currentPath)
		}

	}

	return rating
}
