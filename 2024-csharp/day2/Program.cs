using System.Diagnostics;

// Extract argument to know which part to run
var part = args.FirstOrDefault();

Console.WriteLine("##########");
Console.WriteLine($"Running part {part}");

string inputArg;
var filename = "input-full";

if (args.Length > 1)
{
    inputArg = args[1];
    if (inputArg == "test")
    {
        filename = "input-test";
    }
}

string[] input = File.ReadAllLines(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, filename));


int res;

var watch = Stopwatch.StartNew();

if (part == "2")
{
    res = Part2(input);
}
else
{
    res = Part1(input);
}

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
        string[] report = line.Split(" ");

        if (isReportSafe(report))
        {
            result++;
        }
    }

    return result;
}

bool isReportSafe(string[] report)
{
    var evolution = Evolution.None;

    int.TryParse(report[0], out int previous);

    foreach (var numberString in report.Skip(1))
    {
        int.TryParse(numberString, out int currentNumber);

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
        string[] report = line.Split(" ");

        if (isReportSafe(report))
        {
            result++;
        }
        else
        {
            var strippedSafeFound = false;

            for (var i = 0; i < report.Length; i++)
            {
                var strippedReport = report.ToList();
                strippedReport.RemoveAt(i);

                if (isReportSafe([.. strippedReport]))
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
    }

    return result;
}

enum Evolution
{
    None = 0,
    Increase = 1,
    Decrease = -1,
}

