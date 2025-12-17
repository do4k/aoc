namespace AdventOfCode.Solutions;

public class Day03 : DayBase
{
    public override object Part1(string[] lines)
    {
        long total = 0;
        foreach (var line in lines)
        {
            var max = 0;
            for (var i = 0; i < line.Length; i++)
            {
                for (var j = i + 1; j < line.Length; j++)
                {
                    var val = int.Parse($"{line[i]}{line[j]}");
                    if (val > max)
                    {
                        max = val;
                    }
                }
            }
            total += max;
        }
        return total;
    }

    public override object Part2(string[] lines)
    {
        long total = 0;
        foreach (var line in lines)
        {
            if (line.Length < 12)
                continue;
            var result = new char[12];
            var start = 0;
            for (var pos = 0; pos < 12; pos++)
            {
                var maxIdx = start;
                var end = line.Length - (12 - pos);
                for (var i = start; i <= end; i++)
                {
                    if (line[i] > line[maxIdx])
                        maxIdx = i;
                }
                result[pos] = line[maxIdx];
                start = maxIdx + 1;
            }
            total += long.Parse(new string(result));
        }
        return total;
    }
}
