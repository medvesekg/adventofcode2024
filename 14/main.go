package main

import (
	"adventofcode/utils"
	"fmt"
	"os/exec"
	"regexp"
)

type Robot struct {
	position utils.Point
	velocity utils.Point
}

func main() {
	//partOne()
	partTwo()
}

func partOne() {
	robots := parseInput("input")
	width := 101
	height := 103
	robots = runSimulation(robots, width, height, 100)
	quadrants := countQuadrants(robots, width, height)
	fmt.Println(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
}

func partTwo() {
	width := 101
	height := 103

	robots := parseInput("input")

	for i := 0; i < 10000; i++ {
		robots = tick(robots, width, height)
		//8178
		//if i == 41 || i == 144 || i == 247 {
		if (i-41)%103 == 0 {
			exec.Command("cls")
			fmt.Println()
			fmt.Println("ITERATION: ", i)
			draw(robots, width, height)
		}
	}
}

func runSimulation(robots []Robot, width int, height int, iterations int) []Robot {
	for i := 0; i < iterations; i++ {
		robots = tick(robots, width, height)
	}
	return robots
}

func tick(robots []Robot, width int, height int) []Robot {
	for j := range robots {
		robot := &robots[j]

		robot.position.X += robot.velocity.X
		if robot.position.X >= width {
			robot.position.X %= width
		} else if robot.position.X < 0 {
			robot.position.X += width
		}

		robot.position.Y += robot.velocity.Y
		if robot.position.Y >= height {
			robot.position.Y %= height
		} else if robot.position.Y < 0 {
			robot.position.Y += height
		}
	}
	return robots
}

func countQuadrants(robots []Robot, width int, height int) [4]int {
	quadrants := [4]int{0, 0, 0, 0}
	middleX := (width - 1) / 2
	middleY := (height - 1) / 2

	for _, robot := range robots {
		if robot.position.X == middleX || robot.position.Y == middleY {
			continue
		} else if robot.position.X < middleX && robot.position.Y < middleY {
			quadrants[0]++
		} else if robot.position.X < middleX && robot.position.Y > middleY {
			quadrants[1]++
		} else if robot.position.X > middleX && robot.position.Y < middleY {
			quadrants[2]++
		} else if robot.position.X > middleX && robot.position.Y > middleY {
			quadrants[3]++
		}
	}
	return quadrants
}

func parseInput(path string) []Robot {
	lines := utils.ReadFileLines(path)
	robots := []Robot{}
	r, _ := regexp.Compile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		startX := utils.StrToInt(matches[1])
		startY := utils.StrToInt(matches[2])
		velocityX := utils.StrToInt(matches[3])
		velocityY := utils.StrToInt(matches[4])
		robots = append(robots,
			Robot{
				position: utils.Point{X: startX, Y: startY},
				velocity: utils.Point{X: velocityX, Y: velocityY},
			})
	}

	return robots
}

func draw(robots []Robot, width int, height int) {
	counts := countRobots(robots)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			num, exists := counts[utils.Point{X: x, Y: y}]
			if exists {
				fmt.Print(num)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countRobots(robots []Robot) map[utils.Point]int {
	counts := map[utils.Point]int{}
	for _, robot := range robots {
		_, exists := counts[robot.position]
		if !exists {
			counts[robot.position] = 0
		}
		counts[robot.position]++
	}
	return counts
}
