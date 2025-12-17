// C# Approach:
//   - Uses abstract base class (DayBase) with virtual methods
//   - Relies on inheritance and polymorphism via interfaces (IDay)
//   - Uses reflection to discover day classes at runtime
//   - Pattern matching with regex to extract day numbers from class names
//
// Go Approach:
//   - Uses a simple interface (Solver) - Go favors composition over inheritance
//   - No base class needed - each day implements the interface directly
//   - Manual registration in a map (Registry) - explicit over implicit
//   - No reflection magic - simpler to understand and debug
//
// Both approaches allow adding new days without modifying existing code,
// but Go's is more explicit while C#'s is more automatic.
package solutions

// Solver defines the interface that each day's solution must implement.
// This is analogous to IDay in C#, but Go interfaces are implicitly
// satisfied - no "implements" keyword needed.
type Solver interface {
	Part1(lines []string) any
	Part2(lines []string) any
}

// Registry maps day numbers to their solvers.
// In C# this is done via reflection; in Go we prefer explicit registration.
// This makes the code more predictable and easier to debug.
var Registry = map[int]Solver{
	1: Day01{},
	2: Day02{},
	3: Day03{},
}
