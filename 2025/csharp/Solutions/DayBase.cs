using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions;

public abstract partial class DayBase : IDay
{
    public int DayNumber { get; }

    protected DayBase()
    {
        var match = DayNumberRegex().Match(GetType().Name);
        if (!match.Success)
        {
            throw new InvalidOperationException(
                $"Could not extract day number from class name: {GetType().Name}. " +
                "Class name should be in format 'DayXX' (e.g., Day01, Day25).");
        }
        DayNumber = int.Parse(match.Groups[1].Value);
    }

    public abstract object Part1(string[] lines);
    public abstract object Part2(string[] lines);

    [GeneratedRegex(@"Day(\d+)")]
    private static partial Regex DayNumberRegex();
}
