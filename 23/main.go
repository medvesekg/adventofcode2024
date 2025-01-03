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
	partOne()
	partTwo()
}

func partOne() {
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

func partTwo() {
	connections := parseInput("input")

	longestLen := 0
	longest := []string{}
	for computer := range connections {
		candidate := findMutualConnectionsRecursive([]string{computer}, connections)
		if len(candidate) > longestLen {
			longestLen = len(candidate)
			longest = candidate
		}
	}

	fmt.Println(strings.Join(longest, ","))
}

var cache = map[string][]string{}

func findMutualConnections(set []string, connections map[string][]string) []string {
	sort.Strings(set)
	key := strings.Join(set, ",")
	_, exists := cache[key]
	if exists {
		return cache[key]
	}
	mutual := connections[set[0]]
	for _, computer := range set {
		mutual = findMutual(mutual, connections[computer])
	}
	cache[key] = mutual
	return mutual
}

var cache2 = map[string][]string{}

func findMutualConnectionsRecursive(set []string, connections map[string][]string) []string {
	sort.Strings(set)
	key := strings.Join(set, ",")
	_, exists := cache2[key]
	if exists {
		return cache2[key]
	}
	mutual := findMutualConnections(set, connections)
	longestLen := len(set)
	longest := set
	for _, computer := range mutual {
		newSet := []string{}
		newSet = append(newSet, set...)
		newSet = append(newSet, computer)
		newMutual := findMutualConnectionsRecursive(newSet, connections)
		if len(newMutual) > longestLen {
			longestLen = len(newMutual)
			longest = newMutual
		}
	}
	cache2[key] = longest

	return longest
}

func findMutual(set1 []string, set2 []string) []string {
	mutual := []string{}
	for _, el := range set1 {
		if utils.ArrayContains(set2, el) {
			mutual = append(mutual, el)
		}
	}
	return mutual
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

func printConnections(connections map[string][]string) {
	for computer, directConnections := range connections {
		fmt.Println(computer, directConnections)
	}
}
