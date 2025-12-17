package solutions

// Day03 solves the Day 3 puzzle.
//
// # C# vs Go: Key Differences for this puzzle
//
//  1. String formatting for parsing:
//     C#:  int.Parse($"{line[i]}{line[j]}") - string interpolation
//     Go:  strconv.Atoi(string([]byte{line[i], line[j]})) - explicit byte slice to string
//
//  2. Character comparison:
//     C#:  line[i] > line[maxIdx] - chars are comparable
//     Go:  line[i] > line[maxIdx] - bytes are comparable (same syntax!)
//
//  3. String from chars:
//     C#:  new string(result) - char array to string
//     Go:  string(result) - byte slice to string
//
//  4. Indexing strings:
//     C#:  line[i] returns char (16-bit Unicode)
//     Go:  line[i] returns byte (8-bit) - for ASCII this is fine
type Day03 struct{}

func (d Day03) Part1(lines []string) any {
	var total int64

	for _, line := range lines {
		max := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				// Create two-digit number from characters at positions i and j
				val := mustAtoi(string([]byte{line[i], line[j]}))
				if val > max {
					max = val
				}
			}
		}
		total += int64(max)
	}

	return total
}

func (d Day03) Part2(lines []string) any {
	var total int64

	for _, line := range lines {
		if len(line) < 12 {
			continue
		}

		result := make([]byte, 12)
		start := 0

		for pos := 0; pos < 12; pos++ {
			maxIdx := start
			end := len(line) - (12 - pos)

			for i := start; i <= end; i++ {
				if line[i] > line[maxIdx] {
					maxIdx = i
				}
			}

			result[pos] = line[maxIdx]
			start = maxIdx + 1
		}

		val := mustParseInt64(string(result))
		total += val
	}

	return total
}
