package solutions

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func occurances[T comparable](input []T) map[T]int {
	counts := make(map[T]int)
	for _, item := range input {
		counts[item]++
	}
	return counts
}

func ordered_lists(input []string) ([]int, []int) {
	left := make([]int, len(input))
	right := make([]int, len(input))

	for i, line := range input {
		parts := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		right[i], _ = strconv.Atoi(parts[1])
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	return left, right
}

func Day01_Part1(input []string) string {
	left, right := ordered_lists(input)

	sum := 0.0
	for i := 0; i < len(left); i++ {
		sum += math.Abs(float64(left[i] - right[i]))
	}
	return strconv.Itoa(int(sum))
}

func Day01_Part2(input []string) string {
	left, right := ordered_lists(input)

	occuranceMap := occurances(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += left[i] * occuranceMap[left[i]]
	}

	return strconv.Itoa(sum)
}
