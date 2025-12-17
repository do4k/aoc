# Advent of Code 2025 - Go

A Go port of the C# solutions.

## Running

```bash
# Run a specific day with puzzle input
go run . -day=1

# Run with example input
go run . -day=1 -example

# List available days
go run .
```

## Project Structure

```
go/
├── main.go              # CLI entry point
├── solutions/
│   ├── solver.go        # Interface + registry
│   ├── day01.go         # Day 1 solution
│   └── day02.go         # Day 2 solution
└── inputs/
    ├── day01.txt        # Puzzle inputs
    └── day01_example.txt
```

## C# vs Go: Architectural Comparison

| Aspect | C# | Go |
|--------|----|----|
| **Inheritance** | `DayBase` abstract class with virtual methods | No inheritance; simple interface |
| **Discovery** | Reflection-based, automatic | Explicit registry map |
| **Interfaces** | Explicit `implements` | Implicit satisfaction |
| **Error Handling** | Exceptions | Multiple return values `(value, error)` |
| **Collections** | LINQ method chains | Explicit for loops |
| **Generics** | Long-established | Added in Go 1.18 |

## Adding a New Day

1. Create `solutions/dayNN.go`:

```go
package solutions

type DayNN struct{}

func (d DayNN) Part1(lines []string) any {
    // Your solution
    return 0
}

func (d DayNN) Part2(lines []string) any {
    // Your solution
    return 0
}
```

2. Register in `solutions/solver.go`:

```go
var Registry = map[int]Solver{
    // ...existing entries...
    NN: DayNN{},
}
```

3. Add input files to `inputs/dayNN.txt` and `inputs/dayNN_example.txt`
