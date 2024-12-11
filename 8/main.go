package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	grid := utils.ParseFileGrid("input")
	antennaMap := createAntennaMap(grid)
	antinodes := map[[2]int]bool{}
	for _, positions := range antennaMap {
		for _, combination := range generateCombinations(positions) {

			antinode1 := [2]int{2*combination[0][0] - combination[1][0], 2*combination[0][1] - combination[1][1]}
			antinode2 := [2]int{2*combination[1][0] - combination[0][0], 2*combination[1][1] - combination[0][1]}

			if utils.CheckBounds(antinode1[0], antinode1[1], grid) {
				antinodes[antinode1] = true
			}
			if utils.CheckBounds(antinode2[0], antinode2[1], grid) {
				antinodes[antinode2] = true
			}

		}
	}
	fmt.Println(len(antinodes))
}

func partTwo() {
	grid := utils.ParseFileGrid("input")
	antennaMap := createAntennaMap(grid)
	antinodes := map[[2]int]bool{}
	for _, positions := range antennaMap {
		for _, combination := range generateCombinations(positions) {
			antinodes[combination[0]] = true
			antinodes[combination[1]] = true

			diffY := combination[0][0] - combination[1][0]
			diffX := combination[0][1] - combination[1][1]

			currentY := combination[0][0]
			currentX := combination[0][1]
			for {

				currentY += diffY
				currentX += diffX
				if utils.CheckBounds(currentY, currentX, grid) {
					antinodes[[2]int{currentY, currentX}] = true
				} else {
					break
				}
			}

			diffY = combination[1][0] - combination[0][0]
			diffX = combination[1][1] - combination[0][1]

			currentY = combination[1][0]
			currentX = combination[1][1]

			for {
				currentY += diffY
				currentX += diffX

				if utils.CheckBounds(currentY, currentX, grid) {
					antinodes[[2]int{currentY, currentX}] = true
				} else {
					break
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func createAntennaMap(grid [][]string) map[string][][2]int {
	antennaMap := map[string][][2]int{}
	for y, row := range grid {
		for x, cell := range row {
			if cell != "." {
				_, exists := antennaMap[cell]
				if !exists {
					antennaMap[cell] = [][2]int{}
				}
				antennaMap[cell] = append(antennaMap[cell], [2]int{y, x})
			}
		}
	}
	return antennaMap
}

func generateCombinations(antennaPositions [][2]int) [][2][2]int {
	combinations := [][2][2]int{}
	for i := range antennaPositions {
		for j := i + 1; j < len(antennaPositions); j++ {
			combinations = append(combinations, [2][2]int{antennaPositions[i], antennaPositions[j]})
		}
	}
	return combinations
}
