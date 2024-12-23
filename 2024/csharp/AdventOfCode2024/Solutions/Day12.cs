namespace AdventOfCode2024.Solutions;

[RunSolution]
public class Day12 : BaseSolution, ISolution
{
    public override int Day() => 12;

    private record Node(int X, int Y, char Plot);

    private List<Node> RecursiveGroup(Node node, List<Node> graph, List<Node>? currentGroup = null)
    {
        currentGroup ??= [node];
        var neighbours = graph
            .Except(currentGroup)
            .Where(x =>
                x.X == node.X && x.Y == node.Y - 1 ||
                x.X == node.X && x.Y == node.Y + 1 ||
                x.X == node.X + 1 && x.Y == node.Y ||
                x.X == node.X - 1 && x.Y == node.Y)
            .ToList();
        
        currentGroup.AddRange(neighbours.SelectMany(neighbour => RecursiveGroup(neighbour, graph, currentGroup.Select(x => x).ToList())));

        return neighbours;
    }

    public object Part1()
    {
        var nodes = Lines.SelectMany((l, y) => l.Select((c, x) => new Node(x, y, c))).ToList();
        
        var groups = new List<List<Node>>();
        foreach (var currentNode in nodes)
        {
            if (groups.Any(group => group.Contains(currentNode)))
                continue;

            var potentialNodes = nodes
                .Where(x => x.Plot == currentNode.Plot && !groups.Any(g => g.Contains(x)))
                .ToList();
            var newGroup = RecursiveGroup(
                    currentNode, 
                    potentialNodes
                );
            
            groups.Add(newGroup);
        }

        return 0;
    }

    public object Part2()
    {
        return 0;
    }
}