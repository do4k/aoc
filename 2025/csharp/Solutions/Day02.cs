namespace AdventOfCode.Solutions;

public class Day02 : DayBase
{
    public override object Part1(string[] lines)
    {
        var productIds = lines[0].Split(',').Select(r =>
        {
            var parts = r.Split('-');
            return (Start: long.Parse(parts[0]), End: long.Parse(parts[1]));
        }).Select(r => LongRange(r.Start, r.End - r.Start + 1).ToHashSet())
            .SelectMany(set => set)
            .Where(l => l.ToString().Length % 2 == 0)
            .ToList();

        var count = productIds.Sum(id =>
        {
            var asString = id.ToString();
            var mid = asString.Length / 2;
            var firstHalf = asString[..mid];
            var secondHalf = asString[mid..];
            var invalid = firstHalf.Equals(secondHalf);
            return invalid ? id : 0;
        });

        return count;
    }

    public override object Part2(string[] lines)
    {
        var productIds = lines[0].Split(',').Select(r =>
        {
            var parts = r.Split('-');
            return (Start: long.Parse(parts[0]), End: long.Parse(parts[1]));
        }).Select(r => LongRange(r.Start, r.End - r.Start + 1).ToHashSet())
            .SelectMany(set => set)
            .ToList();

        var count = productIds.Sum(id =>
        {
            var asString = id.ToString();
            var invalid = false;
            for (var len = 1; len <= asString.Length / 2; len++)
            {
                if (asString.Length % len != 0)
                {
                    continue;
                }

                var pattern = asString[..len];
                var repeated = string.Concat(Enumerable.Repeat(pattern, asString.Length / len));
                if (!repeated.Equals(asString))
                {
                    continue;
                }

                invalid = true;
                break;
            }
            return invalid ? id : 0;
        });

        return count;
    }

    private static IEnumerable<long> LongRange(long start, long count)
    {
        for (long i = 0; i < count; i++)
            yield return start + i;
    }
}