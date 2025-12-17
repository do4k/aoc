namespace AdventOfCode.Solutions;

public interface IDay
{
    int DayNumber { get; }
    object Part1(string[] lines);
    object Part2(string[] lines);
}
