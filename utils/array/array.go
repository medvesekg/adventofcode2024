package array

import (
	"strings"
)

func Map[T, V any](array []T, fn func(T) V) []V {
	mapped := []V{}
	for _, value := range array {
		mapped = append(mapped, fn(value))
	}
	return mapped
}

func Filter[T any](array []T, fn func(T) bool) []T {
	filtered := []T{}
	for _, value := range array {
		if fn(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func StringTrim(array []string) []string {
	trimmed := Map(array, func(item string) string {
		return strings.TrimSpace(item)
	})

	return Filter(trimmed, func(item string) bool {
		return item != ""
	})
}

func Split[T any](array []T, fn func(T) bool) [][]T {
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

func CountOccurances[T comparable](arr []T) map[T]int {
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

func Sum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func IndexValid[T any](array []T, i int) bool {
	return i >= 0 && i < len(array)
}
