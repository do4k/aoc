package solutions

import (
	"strconv"
	"strings"
)

// Day02 solves the Day 2 puzzle.
//
// # C# vs Go: Key Differences for this puzzle
//
//  1. LINQ vs loops:
//     C#:  lines[0].Split(',').Select(...).Where(...).Sum(...)
//     Go:  Explicit for loops - Go has no built-in functional collection operations
//     This is perhaps the biggest stylistic difference. Go prefers explicit loops
//     over method chaining. It's more verbose but arguably clearer about what's happening.
//
//  2. Tuples:
//     C#:  (Start: long.Parse(parts[0]), End: long.Parse(parts[1]))
//     Go:  Use a struct or return multiple values - Go has no inline tuple syntax
//
//  3. HashSet:
//     C#:  .ToHashSet()
//     Go:  map[int64]struct{} - idiomatic empty struct as set value (zero memory)
//
//  4. String operations:
//     C#:  asString[..mid] (range), string.Concat(Enumerable.Repeat(...))
//     Go:  asString[:mid] (slice), strings.Repeat(...)
//
//  5. Number types:
//     C#:  long (64-bit), implicit conversions
//     Go:  int64 (explicit), no implicit numeric conversions
//
//  6. Generators/yield:
//     C#:  IEnumerable<long> LongRange(...) { yield return ... }
//     Go:  No yield; use slices, channels, or iterator functions
type Day02 struct{}

func (d Day02) Part1(lines []string) any {
	productIDs := parseAndExpandRanges(lines[0])

	var count int64
	for id := range productIDs {
		asString := strconv.FormatInt(id, 10)

		// Part 1: only process even-length numbers
		if len(asString)%2 != 0 {
			continue
		}

		mid := len(asString) / 2
		firstHalf := asString[:mid]
		secondHalf := asString[mid:]

		if firstHalf == secondHalf {
			count += id
		}
	}
	return count
}

func (d Day02) Part2(lines []string) any {
	productIDs := parseAndExpandRanges(lines[0])

	var count int64
	for id := range productIDs {
		asString := strconv.FormatInt(id, 10)

		// Check if the string is a repetition of any substring
		invalid := false
		for length := 1; length <= len(asString)/2; length++ {
			if len(asString)%length != 0 {
				continue
			}

			pattern := asString[:length]
			repeated := strings.Repeat(pattern, len(asString)/length)

			if repeated == asString {
				invalid = true
				break
			}
		}

		if invalid {
			count += id
		}
	}
	return count
}

// parseAndExpandRanges parses comma-separated ranges like "100-105,200-203"
// and returns a set of all numbers in those ranges.
//
// In C# this used LINQ with SelectMany and ToHashSet.
// In Go we build the set manually with explicit loops.
// Go uses map[T]struct{} as an idiomatic set (struct{} uses zero bytes).
func parseAndExpandRanges(input string) map[int64]struct{} {
	// map[int64]struct{} is Go's idiomatic way to represent a Set
	// struct{} takes zero bytes of memory
	result := make(map[int64]struct{})

	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start := mustParseInt64(parts[0])
		end := mustParseInt64(parts[1])

		// Expand range into set - equivalent to C#'s LongRange().ToHashSet()
		for i := start; i <= end; i++ {
			result[i] = struct{}{}
		}
	}

	return result
}
