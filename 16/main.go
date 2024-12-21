package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
)

func main() {
	partOne()
	partTwo()
}

type Path struct {
	steps []utils.Point
	score int
}

func partOne() {
	grid := parseInput("input")

	bestPaths := findBestPaths(grid)

	fmt.Println(bestPaths[0].score)
}

func findLowestScoreIndex(paths []Path) int {
	lowestSet := false
	lowestIndex := 0
	for i, path := range paths {
		if !lowestSet {
			lowestIndex = i
			lowestSet = true
		}

		if path.score < paths[lowestIndex].score {
			lowestIndex = i
		}
	}
	if !lowestSet {
		panic("OO")
	}
	return lowestIndex
}

func getDirection(steps []utils.Point) utils.Point {
	if len(steps) < 2 {
		return utils.Point{X: 1, Y: 0}
	}

	last := steps[len(steps)-1]
	secondToLast := steps[len(steps)-2]

	if last.X > secondToLast.X {
		return utils.Point{X: 1, Y: 0}
	} else if last.Y > secondToLast.Y {
		return utils.Point{X: 0, Y: 1}
	} else if last.X < secondToLast.X {
		return utils.Point{X: -1, Y: 0}
	} else if last.Y < secondToLast.Y {
		return utils.Point{X: 0, Y: -1}
	}
	panic("OH NO")
}

func getSuccessors(grid [][]string, p Path) []Path {
	currentPoint := p.steps[len(p.steps)-1]
	currentDirection := getDirection(p.steps)
	successors := []Path{}

	for _, direction := range utils.DIRECTIONS2["CARDINAL"] {
		neigbour := utils.Point{X: currentPoint.X + direction.X, Y: currentPoint.Y + direction.Y}

		if !utils.CheckBounds(neigbour.Y, neigbour.X, grid) {
			continue
		}

		neigbourValue := grid[neigbour.Y][neigbour.X]

		if neigbourValue == "#" {
			continue
		}
		score := 1
		if direction != currentDirection {
			score += 1000
			if math.Abs(float64(currentDirection.X)) == math.Abs(float64(direction.X)) && math.Abs(float64(currentDirection.Y)) == math.Abs(float64(direction.Y)) {
				score += 1000
			}
		}
		successor := Path{steps: []utils.Point{}, score: score + p.score}
		successor.steps = append(successor.steps, p.steps...)
		successor.steps = append(successor.steps, neigbour)
		successors = append(successors, successor)

	}

	return successors
}

func parseInput(path string) [][]string {
	return utils.ParseFileGrid(path)
}

func partTwo() {
	grid := parseInput("input")
	bestSpots := map[utils.Point]bool{}
	for _, path := range findBestPaths(grid) {
		for _, spot := range path.steps {
			bestSpots[spot] = true
		}
	}

	fmt.Println(len(bestSpots))
}

func ExistsInArray[T comparable](array []T, searchee T) bool {
	for _, item := range array {
		if item == searchee {
			return true
		}
	}
	return false
}

func findBestPaths(grid [][]string) []Path {

	origin := utils.FindInGrid(grid, "S")
	visited := map[[2]utils.Point]int{}
	frontier := []Path{{steps: []utils.Point{origin}, score: 0}}
	bestPaths := []Path{}

	for len(frontier) > 0 {
		lowestIndex := findLowestScoreIndex(frontier)
		currentPath := frontier[lowestIndex]
		frontier = append(frontier[:lowestIndex], frontier[lowestIndex+1:]...)
		currentPoint := currentPath.steps[len(currentPath.steps)-1]
		currentPointValue := grid[currentPoint.Y][currentPoint.X]
		if currentPointValue == "E" {
			if len(bestPaths) > 0 {
				if currentPath.score > bestPaths[len(bestPaths)-1].score {
					break
				}
			}
			bestPaths = append(bestPaths, currentPath)
			continue
		}
		currentDirection := getDirection(currentPath.steps)
		visited[[2]utils.Point{currentPoint, currentDirection}] = currentPath.score

		for _, successor := range getSuccessors(grid, currentPath) {
			successorDirection := getDirection(successor.steps)
			successorLastPoint := successor.steps[len(successor.steps)-1]
			alreadyVisited, hasAlreadyVisited := visited[[2]utils.Point{successorLastPoint, successorDirection}]
			if hasAlreadyVisited && alreadyVisited <= successor.score {
				continue
			}

			frontier = append(frontier, successor)
		}

	}

	return bestPaths
}
