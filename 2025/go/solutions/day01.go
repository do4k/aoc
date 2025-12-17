package solutions

// Day01 solves the Day 1 puzzle.
//
// # C# vs Go: Key Differences
//
//  1. Type declaration:
//     C#:  public class Day01 : DayBase { ... }
//     Go:  type Day01 struct{} - no inheritance, just a simple struct
//
//  2. Method receivers:
//     C#:  public override object Part1(string[] lines) { ... }
//     Go:  func (d Day01) Part1(lines []string) any { ... }
//     Go methods have explicit receivers; C# uses implicit 'this'
//
//  3. String slicing:
//     C#:  l[1..] (range operator)
//     Go:  l[1:] (slice syntax)
//
//  4. Parsing:
//     C#:  int.Parse(str) - throws on error
//     Go:  strconv.Atoi(str) - returns (value, error)
//     Go's explicit error handling vs C#'s exceptions
//
//  5. Ternary operator:
//     C#:  l[0] == 'L' ? -1 : 1
//     Go:  if/else (Go has no ternary operator)
//
//  6. Constants:
//     C#:  const int size = 100; (inside method)
//     Go:  const size = 100 (package or function level, type inferred)
type Day01 struct{}

func (d Day01) Part1(lines []string) any {
	return simulate(lines, false)
}

func (d Day01) Part2(lines []string) any {
	return simulate(lines, true)
}

func simulate(lines []string, countAllZeroes bool) int {
	const size = 100
	position := 50
	count := 0

	for _, l := range lines {
		// Go has no ternary operator, so we use if/else
		direction := 1
		if l[0] == 'L' {
			direction = -1
		}

		// strconv.Atoi returns (int, error) - Go's explicit error handling
		// We use mustAtoi to panic on parse errors
		steps := mustAtoi(l[1:])

		if countAllZeroes {
			for i := 0; i < steps; i++ {
				position = (position + direction + size) % size
				if position == 0 {
					count++
				}
			}
		} else {
			position = (position + direction*steps) % size
			if position < 0 {
				position += size
			}
			if position == 0 {
				count++
			}
		}
	}
	return count
}
