package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	grid := utils.ParseFileGrid("input")
	origin := utils.FindInGrid(grid, "S")
	target := utils.FindInGrid(grid, "E")
	path := findPath(grid, origin, target)

	possibleCheatDirections := []utils.Point{
		{X: 2, Y: 0},
		{X: -2, Y: 0},
		{X: 0, Y: -2},
		{X: 0, Y: 2},
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 1},
		{X: -1, Y: 1},
	}

	cheats := map[[2]utils.Point]int{}
	cheatCount := map[int]int{}
	result := 0
	for i := 0; i < len(path)-3; i++ {
		point := path[i]
		for _, possibleDirection := range possibleCheatDirections {
			possiblePoint := utils.Point{X: point.X + possibleDirection.X, Y: point.Y + possibleDirection.Y}
			for j := i + 3; j < len(path); j++ {
				pointForwardOnPath := path[j]
				if possiblePoint == pointForwardOnPath {
					cheats[[2]utils.Point{point, pointForwardOnPath}] = 1
					cheatCount[j-i-2]++
					if j-i-2 >= 100 {
						result++
					}
				}
			}
		}
	}

	fmt.Println(result)

}

func findPath(grid [][]string, origin utils.Point, target utils.Point) []utils.Point {
	frontier := [][]utils.Point{{origin}}
	visited := map[utils.Point]bool{origin: true}
	for len(frontier) > 0 {
		currentPath := frontier[0]
		currentPoint := currentPath[len(currentPath)-1]
		frontier = frontier[1:]
		for _, neighbour := range utils.GridGetNeighbours(grid, currentPoint, utils.DIRECTIONS2["CARDINAL"]) {
			_, alreadyVisited := visited[neighbour]

			if neighbour == target {
				return append(currentPath, target)
			}

			if !alreadyVisited && grid[neighbour.Y][neighbour.X] != "#" {
				visited[neighbour] = true
				newPath := []utils.Point{}
				newPath = append(newPath, currentPath...)
				newPath = append(newPath, neighbour)
				frontier = append(frontier, newPath)
			}
		}
	}
	return []utils.Point{}
}

func visualizePath(grid [][]string, path []utils.Point) {
	for y := range grid {
		for x := range grid[y] {
			point := utils.Point{X: x, Y: y}
			if utils.ArrayContains(path, point) {
				fmt.Print("\x1b[36mX\x1b[0m")
			} else {
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println()
	}
}
