using System.Diagnostics;
using System.Reflection;
using AdventOfCode.Solutions;
using Spectre.Console;

var days = DiscoverDays();

if (days.Count == 0)
{
    AnsiConsole.MarkupLine("[red]No solutions found![/] Create a class extending DayBase in the Solutions folder.");
    return;
}

while (true)
{
    AnsiConsole.Clear();
    AnsiConsole.Write(new FigletText("Advent of Code").Color(Color.Green));
    AnsiConsole.MarkupLine("[grey]2025 Edition[/]\n");

    var choices = days
        .OrderBy(d => d.DayNumber)
        .Select(d => $"Day {d.DayNumber:D2}")
        .Append("[red]Exit[/]")
        .ToList();

    var selection = AnsiConsole.Prompt(
        new SelectionPrompt<string>()
            .Title("Select a [green]day[/] to run:")
            .PageSize(10)
            .HighlightStyle(new Style(Color.Green, decoration: Decoration.Bold))
            .AddChoices(choices));

    if (selection.Contains("Exit"))
    {
        AnsiConsole.MarkupLine("[yellow]Goodbye![/]");
        break;
    }

    var dayNumber = int.Parse(selection.Replace("Day ", ""));
    var day = days.First(d => d.DayNumber == dayNumber);

    var inputType = AnsiConsole.Prompt(
        new SelectionPrompt<string>()
            .Title($"Select [green]input type[/] for Day {dayNumber:D2}:")
            .HighlightStyle(new Style(Color.Green, decoration: Decoration.Bold))
            .AddChoices("Puzzle Input", "Example Input", "[grey]Back[/]"));

    if (inputType.Contains("Back"))
    {
        continue;
    }

    var inputFile = inputType == "Example Input"
        ? $"Inputs/Day{dayNumber:D2}_example.txt"
        : $"Inputs/Day{dayNumber:D2}.txt";

    if (!File.Exists(inputFile))
    {
        AnsiConsole.MarkupLine($"\n[red]Input file not found:[/] {inputFile}");
        AnsiConsole.MarkupLine("[grey]Press any key to continue...[/]");
        Console.ReadKey(true);
        continue;
    }

    var lines = File.ReadAllLines(inputFile);

    AnsiConsole.WriteLine();
    AnsiConsole.Write(new Rule($"[green]Day {dayNumber:D2}[/]").RuleStyle("grey"));
    AnsiConsole.WriteLine();

    var table = new Table()
        .Border(TableBorder.Rounded)
        .AddColumn(new TableColumn("[bold]Part[/]").Centered())
        .AddColumn(new TableColumn("[bold]Answer[/]").Centered())
        .AddColumn(new TableColumn("[bold]Time[/]").Centered());

    var sw = new Stopwatch();

    sw.Restart();
    var result1 = day.Part1(lines);
    sw.Stop();
    var time1 = FormatTime(sw.Elapsed);
    table.AddRow("[cyan]Part 1[/]", $"[white]{result1}[/]", $"[grey]{time1}[/]");

    sw.Restart();
    var result2 = day.Part2(lines);
    sw.Stop();
    var time2 = FormatTime(sw.Elapsed);
    table.AddRow("[cyan]Part 2[/]", $"[white]{result2}[/]", $"[grey]{time2}[/]");

    AnsiConsole.Write(table);
    AnsiConsole.WriteLine();
    AnsiConsole.MarkupLine("[grey]Press any key to continue...[/]");
    Console.ReadKey(true);
}

static List<IDay> DiscoverDays()
{
    return Assembly.GetExecutingAssembly()
        .GetTypes()
        .Where(t => t.IsClass && !t.IsAbstract && typeof(IDay).IsAssignableFrom(t))
        .Select(t => (IDay)Activator.CreateInstance(t)!)
        .ToList();
}

static string FormatTime(TimeSpan elapsed)
{
    return elapsed.TotalMilliseconds switch
    {
        < 1 => $"{elapsed.TotalMicroseconds:F2} μs",
        < 1000 => $"{elapsed.TotalMilliseconds:F2} ms",
        _ => $"{elapsed.TotalSeconds:F2} s"
    };
}
