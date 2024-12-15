package utils

import (
	"cmp"
	"os"
	"strconv"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	CheckError(err)
	return string(data)
}

func SplitByLine(str string) []string {
	return strings.Split(str, "\n")
}

func SplitByWord(str string) []string {
	return strings.Fields(str)
}

func ArrayMap[T, V any](array []T, fn func(T) V) []V {
	mapped := []V{}
	for _, value := range array {
		mapped = append(mapped, fn(value))
	}
	return mapped
}

func ArrayFilter[T any](array []T, fn func(T) bool) []T {
	filtered := []T{}
	for _, value := range array {
		if fn(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func StringArrayTrim(array []string) []string {
	trimmed := ArrayMap(array, func(item string) string {
		return strings.TrimSpace(item)
	})

	return ArrayFilter(trimmed, func(item string) bool {
		return item != ""
	})
}

func ArraySplit[T any](array []T, fn func(T) bool) [][]T {
	chunks := [][]T{}
	chunk := []T{}
	for _, val := range array {
		if fn(val) {
			if len(chunk) > 0 {
				chunks = append(chunks, chunk)
			}
			chunk = []T{}
		} else {
			chunk = append(chunk, val)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	return chunks
}

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	CheckError(err)
	return num
}

func IntToStr(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func Identity[T any](item T) T {
	return item
}

func ReadFileLines(path string) []string {
	data := ReadFile(path)
	lines := SplitByLine(data)
	return StringArrayTrim(lines)
}

func ParseFile[T any](path string, fn func(val string) T) [][]T {
	rawLines := ReadFileLines(path)
	lines := ArrayMap(rawLines, func(line string) []T {
		return ArrayMap(strings.Fields(line), fn)
	})
	return lines
}

func RowsToCols[T any](grid [][]T) [][]T {
	inverted := [][]T{}

	for _, row := range grid {
		for i, val := range row {
			if len(inverted) <= i {
				inverted = append(inverted, []T{})
			}
			inverted[i] = append(inverted[i], val)
		}
	}

	return inverted
}

func ArrayCountOccurances[T comparable](arr []T) map[T]int {
	counts := map[T]int{}
	for _, val := range arr {
		_, exists := counts[val]
		if !exists {
			counts[val] = 0
		}
		counts[val]++
	}
	return counts
}

func ArraySum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func ArrayMax[T any, K cmp.Ordered](array []T, fn func(item T) K) K {
	maxSet := false
	var max K
	for _, item := range array {
		val := fn(item)
		if !maxSet {
			max = val
			maxSet = true
		}

		if maxSet && val > max {
			max = val
		}
	}
	return max
}

func ArrayMin[T any, K cmp.Ordered](array []T, fn func(item T) K) K {
	minSet := false
	var min K
	for _, item := range array {
		val := fn(item)
		if !minSet {
			min = val
			minSet = true
		}

		if minSet && val < min {
			min = val
		}
	}
	return min
}

func IndexValid[T any](array []T, i int) bool {
	return i >= 0 && i < len(array)
}

func CheckBounds[T any](y int, x int, grid [][]T) bool {
	return IndexValid(grid, y) && IndexValid(grid[y], x)
}

func PrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}

func ArrayUnique[T comparable](array []T) []T {
	seen := map[T]bool{}
	unique := []T{}
	for _, val := range array {
		if !seen[val] {
			unique = append(unique, val)
		}
		seen[val] = true
	}
	return unique
}

func ArrayAny[T any](array []T, fn func(T) bool) bool {
	for _, value := range array {
		if fn(value) {
			return true
		}
	}
	return false
}

func ParseFileGrid(path string) [][]string {
	lines := ReadFileLines(path)
	return ArrayMap(lines, func(line string) []string {
		return strings.Split(line, "")
	})
}

func ParseFileGridInt(path string) [][]int {
	lines := ReadFileLines(path)
	return ArrayMap(lines, func(line string) []int {
		return ArrayMap(strings.Split(line, ""), StrToInt)
	})
}

var DIRECTIONS = map[string][][2]int{
	"CARDINAL": {{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
}

var DIRECTIONS2 = map[string][]Point{
	"CARDINAL": {{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
}

type Point struct {
	Y int
	X int
}

func GridGetNeighbours[T any](grid [][]T, origin Point, directions []Point) []Point {
	neighbours := ArrayMap(directions, func(point Point) Point {
		return Point{X: origin.X + point.X, Y: origin.Y + point.Y}
	})
	return ArrayFilter(neighbours, func(point Point) bool {
		return CheckBounds(point.Y, point.X, grid)
	})
}
