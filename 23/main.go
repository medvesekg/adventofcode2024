package main

import (
	"adventofcode/utils"
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
)

func main() {
	connections := parseInput("input")

	sets := map[[3]string]bool{}
	for computer, directConnections := range connections {
		if strings.HasPrefix(computer, "t") {
			for _, connectedComputer := range directConnections {
				for _, otherComputer := range connections[connectedComputer] {
					if utils.ArrayContains(directConnections, otherComputer) {
						set := []string{computer, connectedComputer, otherComputer}
						sort.Strings(set)
						sets[[3]string{set[0], set[1], set[2]}] = true
					}
				}
			}
		}

	}
	setsList := slices.Collect(maps.Keys(sets))
	fmt.Println(len(setsList))
}

func parseInput(path string) map[string][]string {
	lines := utils.ReadFileLines(path)
	input := utils.ArrayMap(lines, func(line string) []string {
		return strings.Split(line, "-")
	})
	connections := map[string][]string{}
	for _, pair := range input {
		connections[pair[0]] = append(connections[pair[0]], pair[1])
		connections[pair[1]] = append(connections[pair[1]], pair[0])
	}
	return connections
}
