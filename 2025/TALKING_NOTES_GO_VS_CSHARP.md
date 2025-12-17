# Go vs C# - Talking Notes for Advent of Code 2025

> **Context**: I'm primarily a C# engineer learning Go. These notes highlight the key differences I've discovered while implementing the same Advent of Code puzzles in both languages.

---

## 1. No Ternary Operator in Go

**The most jarring difference for C# developers.**

### C# (Day01.cs)
```csharp
var direction = l[0] == 'L' ? -1 : 1;
```

### Go (day01.go)
```go
direction := 1
if l[0] == 'L' {
    direction = -1
}
```

### Why Go Doesn't Have Ternaries
- Go prioritizes **readability over brevity**
- The Go team argues ternaries can become nested and unreadable: `a ? b ? c : d : e`
- Forces you to be explicit about conditional logic
- No "clever one-liners" — which Go considers a feature, not a bug

### Talking Point
> "Coming from C#, I initially found this verbose. But Go's philosophy is that code is read far more often than it's written, so clarity wins over conciseness."

---

## 2. Explicit Error Handling: `(value, error)` Pattern

**The signature Go idiom that C# developers must internalize.**

### C# (exceptions)
```csharp
var steps = int.Parse(l[1..]);  // Throws FormatException on failure
```

### Go (explicit errors)
```go
steps, err := strconv.Atoi(l[1:])
if err != nil {
    return 0, err  // or handle the error
}
```

### The `_` Discard
```go
steps, _ := strconv.Atoi(l[1:])  // Explicitly ignoring the error
```

### My Helper Pattern (helpers.go)
```go
func must[T any](val T, err error) T {
    if err != nil {
        panic(err)
    }
    return val
}

func mustAtoi(s string) int {
    return must(strconv.Atoi(s))
}
```

### Why Go Does This
- **Errors are values**, not exceptional control flow
- Forces you to think about error cases at every step
- No hidden exceptions bubbling up the stack
- Makes error paths explicit and visible in code review

### Talking Point
> "In C#, I'd wrap a whole method in try/catch. In Go, I'm forced to handle or acknowledge each potential failure point. It's more verbose but makes the failure modes explicit."

---

## 3. No Reflection-Based Discovery (Explicit Over Implicit)

**How we register and discover day solutions differs fundamentally.**

### C# (Program.cs + DayBase.cs) — Reflection Magic
```csharp
// Uses reflection to find all IDay implementations
var days = Assembly.GetExecutingAssembly()
    .GetTypes()
    .Where(t => typeof(IDay).IsAssignableFrom(t) && !t.IsAbstract)
    .Select(t => (IDay)Activator.CreateInstance(t)!)
    .ToList();

// Day number extracted via regex from class name
[GeneratedRegex(@"Day(\d+)")]
private static partial Regex DayNumberRegex();
```

### Go (solver.go) — Explicit Registry
```go
var Registry = map[int]Solver{
    1: Day01{},
    2: Day02{},
    3: Day03{},
}
```

### Key Differences
| Aspect | C# | Go |
|--------|----|----|
| Discovery | Automatic via reflection | Manual registration |
| Adding a day | Just create the class | Create struct + add to registry |
| Debugging | "Why isn't my class found?" | Explicit, obvious |
| Performance | Reflection has overhead | Zero overhead |

### Why Go Avoids Reflection
- Go has `reflect` package, but it's discouraged for everyday use
- Reflection defeats static type checking
- Makes code harder to understand and debug
- "If you need reflection, you probably have a design problem"

### Talking Point
> "In C#, I love the magic of 'just add a class and it works'. In Go, you pay for explicitness with a line in the registry, but you gain predictability and compile-time guarantees."

---

## 4. Generics in Go (Newer, More Limited)

**Go 1.18 (2022) added generics — but they're not C# generics.**

### My Example (helpers.go)
```go
func must[T any](val T, err error) T {
    if err != nil {
        panic(err)
    }
    return val
}
```

### Key Differences from C#

| Feature | C# | Go |
|---------|----|----|
| Introduced | 2005 (.NET 2.0) | 2022 (Go 1.18) |
| Constraints | `where T : class, IComparable` | `[T any]` or `[T comparable]` |
| Variance | Covariant/contravariant (`in`/`out`) | None |
| Type inference | Usually complete | Often requires explicit types |
| Generic methods on interfaces | Yes | Limited |

### What Go Generics Can Do
```go
// Constraint to comparable types (supports == and !=)
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}
```

### What's Missing vs C#
- No generic specialization
- No operator constraints (can't do `T + T`)
- No default values (`default(T)`)
- No generic attributes

### Talking Point
> "I haven't needed complex generics in my AoC solutions yet, but the `must[T any]` helper shows the basic syntax. Go generics are intentionally simpler — they solve the 'containers and algorithms' problem without the full complexity of C# generics."

---

## 5. LINQ vs Explicit Loops

**Perhaps the biggest stylistic adjustment for C# developers.**

### C# (Day02.cs) — LINQ Chains
```csharp
var productIds = lines[0].Split(',')
    .Select(r => {
        var parts = r.Split('-');
        return (Start: long.Parse(parts[0]), End: long.Parse(parts[1]));
    })
    .SelectMany(r => LongRange(r.Start, r.End - r.Start + 1))
    .ToHashSet()
    .Where(l => l.ToString().Length % 2 == 0)
    .ToList();

var count = productIds.Sum(id => /* ... */);
```

### Go (day02.go) — Explicit Loops
```go
result := make(map[int64]struct{})
ranges := strings.Split(input, ",")

for _, r := range ranges {
    parts := strings.Split(r, "-")
    start := mustParseInt64(parts[0])
    end := mustParseInt64(parts[1])
    
    for i := start; i <= end; i++ {
        result[i] = struct{}{}
    }
}
```

### Why Go Prefers Loops
- No method chaining on collections (no extension methods)
- Explicit loops are seen as clearer about what's happening
- Easier to debug (set breakpoints inside loops)
- No hidden allocations from intermediate collections

### Talking Point
> "I miss LINQ every single day. But Go's loops are explicit about allocations and control flow. There's no hidden `IEnumerable` machinery — what you see is what you get."

---

## 6. No Inline Tuple Syntax

### C#
```csharp
var result = (Start: long.Parse(parts[0]), End: long.Parse(parts[1]));
Console.WriteLine(result.Start);  // Named tuple access
```

### Go — Use Structs or Multiple Return Values
```go
// Option 1: Define a struct
type Range struct {
    Start, End int64
}

// Option 2: Multiple return values
func parseRange(s string) (start, end int64) {
    // ...
    return start, end
}
```

### Talking Point
> "C# tuples are incredibly convenient for ad-hoc groupings. In Go, I either define a small struct (more boilerplate) or use multiple return values. It pushes you toward explicit typing."

---

## 7. Interfaces: Implicit vs Explicit

### C# (explicit implementation)
```csharp
public interface IDay
{
    int DayNumber { get; }
    object Part1(string[] lines);
    object Part2(string[] lines);
}

public class Day01 : DayBase, IDay  // Explicit declaration
{
    // ...
}
```

### Go (implicit satisfaction)
```go
type Solver interface {
    Part1(lines []string) any
    Part2(lines []string) any
}

type Day01 struct{}  // No "implements" keyword

func (d Day01) Part1(lines []string) any { /* ... */ }
func (d Day01) Part2(lines []string) any { /* ... */ }
// Day01 now implicitly implements Solver!
```

### Why This Matters
- **Duck typing at compile time**: if it has the methods, it implements the interface
- No need to modify a type to implement a new interface
- Enables "interface segregation" naturally
- Can implement interfaces from other packages without modifying them

### Talking Point
> "In Go, I don't declare 'Day01 implements Solver'. If Day01 has the right methods, it *is* a Solver. This feels magical at first, but it's powerful for decoupling."

---

## 8. Other Syntax Differences

### String Slicing
```csharp
// C#: Range operator
var substring = line[1..];      // From index 1 to end
var middle = line[..mid];       // From start to mid
```
```go
// Go: Slice syntax
substring := line[1:]           // From index 1 to end
middle := line[:mid]            // From start to mid
```

### String from Characters
```csharp
// C#
var str = $"{line[i]}{line[j]}";           // Interpolation
var str2 = new string(charArray);          // From char[]
```
```go
// Go
str := string([]byte{line[i], line[j]})    // From byte slice
```

### Sets
```csharp
// C#
var set = new HashSet<long>();
set.Add(value);
```
```go
// Go (idiomatic empty struct as value)
set := make(map[int64]struct{})
set[value] = struct{}{}

// Check membership
if _, exists := set[value]; exists { }
```

---

## 9. Inheritance vs Composition

### C# — Inheritance Hierarchy
```csharp
public abstract class DayBase : IDay
{
    public int DayNumber { get; }
    protected DayBase() { /* extract from class name */ }
    public abstract object Part1(string[] lines);
    public abstract object Part2(string[] lines);
}

public class Day01 : DayBase { /* ... */ }
```

### Go — Composition (No Inheritance)
```go
type Day01 struct{}  // Just a struct, no base class

func (d Day01) Part1(lines []string) any { /* ... */ }
func (d Day01) Part2(lines []string) any { /* ... */ }
```

### Key Point
- Go has **no inheritance**
- Use **embedding** for composition: `type Extended struct { Base }`
- Prefer small interfaces over deep hierarchies

---

## Summary: The Go Philosophy

| Principle | Impact |
|-----------|--------|
| **Simplicity over features** | No ternary, limited generics |
| **Explicit over implicit** | Error handling, no reflection magic |
| **Composition over inheritance** | Interfaces, embedding |
| **Readability over brevity** | More lines, clearer intent |

### My Takeaway After 3 Days
> "Go forces me to write more lines of code, but each line is simpler and more obvious. There's no 'magic' — which is frustrating when you want it, but reassuring when debugging someone else's code."

---

## Questions to Expect from Go Developers

1. **"Why the `must*` helpers?"** — AoC inputs are trusted, so panicking on parse errors is acceptable. In production, you'd propagate errors.

2. **"Why not use a library for collections?"** — There are some (like `golang.org/x/exp/slices`), but I wanted to learn the core language first.

3. **"Your map-as-set pattern is correct!"** — Using `map[T]struct{}` is idiomatic; `struct{}` is zero-size.

