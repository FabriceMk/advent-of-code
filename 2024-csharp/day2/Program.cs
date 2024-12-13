using System.Diagnostics;

// Extract argument to know which part to run
var part = args.FirstOrDefault();

Console.WriteLine("##########");
Console.WriteLine($"Running part {part}");

var filename = (args.Length > 1 && args[1] == "test") ? "input-test" : "input-full";
string[] input = File.ReadAllLines(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, filename));

var watch = Stopwatch.StartNew();

var res = (part == "1") ? Part1(input) : Part2(input);

watch.Stop();

Console.WriteLine("##########");
Console.WriteLine($"Execution time: {watch.ElapsedMilliseconds}ms");
Console.WriteLine("Output:");
Console.WriteLine(res);
Console.WriteLine("##########");

int Part1(string[] input)
{
    var result = 0;

    foreach (var line in input)
    {
        var report = line.Split(" ", StringSplitOptions.RemoveEmptyEntries).Select(int.Parse).ToList();

        if (isReportSafe(report))
        {
            result++;
        }
    }

    return result;
}

bool isReportSafe(List<int> report)
{
    var evolution = Evolution.None;

    var previous = report[0];

    foreach (var currentNumber in report.Skip(1))
    {
        if (currentNumber == previous || (evolution == Evolution.Increase && currentNumber < previous) ||
        (evolution == Evolution.Decrease && currentNumber > previous))
        {
            return false;
        }

        if (evolution == Evolution.None)
        {
            if (currentNumber > previous)
            {
                evolution = Evolution.Increase;
            }
            else
            {
                evolution = Evolution.Decrease;
            }
        }

        var gap = Math.Abs(currentNumber - previous);

        if (gap < 1 || gap > 3)
        {
            return false;
        }

        previous = currentNumber;
    }

    return true;
}

int Part2(string[] input)
{
    var result = 0;

    foreach (var line in input)
    {
        var report = line.Split(" ", StringSplitOptions.RemoveEmptyEntries).Select(int.Parse).ToList();

        if (isReportSafe(report))
        {
            result++;
            continue;
        }

        var strippedSafeFound = false;

        for (var i = 0; i < report.Count; i++)
        {
            var strippedReport = report.ToList();
            strippedReport.RemoveAt(i);

            if (isReportSafe(strippedReport))
            {
                strippedSafeFound = true;
                break;
            }
        }

        if (strippedSafeFound)
        {
            result++;
        }
    }

    return result;
}

enum Evolution
{
    None = 0,
    Increase = 1,
    Decrease = -1,
}

