package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	size := 71
	iterations := 1024
	input := "input"

	grid := createGrid(size)
	incoming := parseInput(input)
	grid = fall(grid, incoming, iterations)
	origin := utils.Point{X: 0, Y: 0}
	target := utils.Point{X: size - 1, Y: size - 1}
	shortest := shortestPath(grid, origin, target)
	fmt.Println(shortest)
}

func partTwo() {
	size := 71
	startIteration := 1024
	input := "input"

	origin := utils.Point{X: 0, Y: 0}
	target := utils.Point{X: size - 1, Y: size - 1}
	incoming := parseInput(input)
	grid := createGrid(size)
	grid = fall(grid, incoming, startIteration)

	for i := startIteration; i < len(incoming); i++ {
		pos := incoming[i]
		grid[pos[1]][pos[0]] = "#"
		exists := pathExists(grid, origin, target)
		if !exists {
			fmt.Printf("%d,%d\n", pos[0], pos[1])
			return
		}
	}
}

func parseInput(path string) [][]int {
	lines := utils.ReadFileLines(path)
	r, _ := regexp.Compile(`\d+`)
	return utils.ArrayMap(lines, func(line string) []int {
		matches := r.FindAllString(line, -1)
		return utils.ArrayMap(matches, utils.StrToInt)
	})
}

func createGrid(size int) [][]string {
	grid := make([][]string, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			grid[i][j] = "."
		}
	}
	return grid
}

func fall(grid [][]string, sequence [][]int, iterations int) [][]string {
	for i := 0; i < iterations; i++ {
		pos := sequence[i]
		grid[pos[1]][pos[0]] = "#"
	}

	return grid
}

func shortestPath(grid [][]string, origin, target utils.Point) int {

	frontier := [][]utils.Point{{origin}}
	visited := make(map[utils.Point]int, len(grid)*len(grid))
	visited[origin] = 0
	solutions := [][]utils.Point{}

	for len(frontier) > 0 {
		index := getClosestIndex(frontier, target)
		currentPath := frontier[index]
		frontier = append(frontier[:index], frontier[index+1:]...)
		currentPoint := currentPath[len(currentPath)-1]

		if currentPoint == target {
			solutions = append(solutions, currentPath)
		}

		neighbours := utils.GridGetNeighbours(grid, currentPoint, utils.DIRECTIONS2["CARDINAL"])
		for _, neigbour := range neighbours {
			vistiedLength, hasVisited := visited[neigbour]

			if grid[neigbour.Y][neigbour.X] == "#" || (hasVisited && len(currentPath)+1 > vistiedLength) {
				continue
			}
			visited[neigbour] = len(currentPath)
			newPath := []utils.Point{}
			newPath = append(newPath, currentPath...)
			frontier = append(frontier, append(newPath, neigbour))
		}
	}

	return utils.ArrayMin(solutions, func(path []utils.Point) int {
		return len(path) - 1
	})
}

func pathExists(grid [][]string, origin, target utils.Point) bool {

	frontier := [][]utils.Point{{origin}}
	visited := make(map[utils.Point]int, len(grid)*len(grid))
	visited[origin] = 0

	for len(frontier) > 0 {
		index := getClosestIndex(frontier, target)
		currentPath := frontier[index]
		frontier = append(frontier[:index], frontier[index+1:]...)
		currentPoint := currentPath[len(currentPath)-1]

		if currentPoint == target {
			return true
		}

		neighbours := utils.GridGetNeighbours(grid, currentPoint, utils.DIRECTIONS2["CARDINAL"])
		for _, neigbour := range neighbours {
			vistiedLength, hasVisited := visited[neigbour]

			if grid[neigbour.Y][neigbour.X] == "#" || (hasVisited && len(currentPath)+1 > vistiedLength) {
				continue
			}
			visited[neigbour] = len(currentPath)
			newPath := []utils.Point{}
			newPath = append(newPath, currentPath...)
			frontier = append(frontier, append(newPath, neigbour))
		}
	}

	return false
}

func visualizeSearch(grid [][]string, path []utils.Point) {
	newGrid := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	for _, point := range path {
		newGrid[point.Y][point.X] = "O"
	}
	utils.PrintGrid(newGrid)
	fmt.Println()
}

func calculateDistance(current utils.Point, target utils.Point) float64 {
	return math.Sqrt(math.Pow(float64(target.X)-float64(current.X), 2) + math.Pow(float64(target.Y)-float64(current.Y), 2))
}

func getClosestIndex(paths [][]utils.Point, target utils.Point) int {
	closestIndex := -1
	distance := math.Inf(1)

	for i, path := range paths {
		current := path[len(path)-1]
		currentDistance := calculateDistance(current, target)
		if currentDistance < distance {
			distance = currentDistance
			closestIndex = i
		}
	}
	return closestIndex
}
