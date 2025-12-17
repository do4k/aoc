namespace AdventOfCode.Solutions;

public class Day01 : DayBase
{
    public override object Part1(string[] lines)
    {
        return Simulate(lines, false);
    }

    public override object Part2(string[] lines)
    {
        return Simulate(lines, true);
    }

    private static int Simulate(string[] lines, bool countAllZeroes)
    {
        const int size = 100;
        var position = 50;
        var count = 0;

        foreach (var l in lines)
        {
            var direction = l[0] == 'L' ? -1 : 1;
            var steps = int.Parse(l.ToString()[1..]);

            if (countAllZeroes)
            {
                for (var i = 0; i < steps; i++)
                {
                    position = (position + direction + size) % size;
                    if (position == 0)
                    {
                        count++;
                    }
                }
            }
            else
            {
                position = (position + direction * steps) % size;

                if (position < 0)
                {
                    position += size;
                }

                if (position == 0)
                {
                    count++;
                }
            }
        }
        return count;
    }

    // public override object Part1(string[] lines) =>
    //     lines
    //         .Select(l => (Direction: l[0] == 'L' ? -1 : 1, Steps: int.Parse(l[1..])))
    //         .Aggregate((Position: 50, Count: 0), (state, move) =>
    //         {
    //             var newPos = ((state.Position + move.Direction * move.Steps) % Size + Size) % Size;
    //             return (newPos, state.Count + (newPos == 0 ? 1 : 0));
    //         })
    //         .Count;

    // public override object Part2(string[] lines) =>
    //     lines
    //         .Select(l => (Direction: l[0] == 'L' ? -1 : 1, Steps: int.Parse(l[1..])))
    //         .SelectMany(move => Enumerable.Repeat(move.Direction, move.Steps))
    //         .Aggregate((Position: 50, Count: 0), (state, dir) =>
    //         {
    //             var newPos = ((state.Position + dir) % Size + Size) % Size;
    //             return (newPos, state.Count + (newPos == 0 ? 1 : 0));
    //         })
    //         .Count;
}
