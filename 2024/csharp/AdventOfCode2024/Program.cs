
using System.Diagnostics;
using System.Reflection;
using AdventOfCode2024;
using AdventOfCode2024.Solutions;

var solutions = Assembly 
    .GetExecutingAssembly() 
    .GetTypes() 
    .Where(x => typeof(ISolution).IsAssignableFrom(x) && 
                !x.IsAbstract && 
                x.GetConstructor(Type.EmptyTypes) != null &&
                x.IsDefined(typeof(RunSolutionAttribute), false)) 
    .Select(x => (ISolution)Activator.CreateInstance(x)!) 
    .ToList();

foreach (var s in solutions)
{
    var sw1 = Stopwatch.StartNew();
    var part1 = s.Part1();
    sw1.Stop();
    var sw2 = Stopwatch.StartNew();
    var part2 = s.Part2();
    sw2.Stop();
    Console.WriteLine($"Day {((BaseSolution)s!).Day()} Part 1: {part1} ({sw1.ElapsedMilliseconds}ms) Part 2: {part2} ({sw2.ElapsedMilliseconds}ms)");
}