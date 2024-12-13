using System.Diagnostics;
using System.Text.RegularExpressions;

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
        var pattern = @"mul\((\d+),(\d+)\)";
        var results = Regex.Matches(line, pattern);

        result += results.Select(x => int.Parse(x.Groups[1].Value) * int.Parse(x.Groups[2].Value)).Sum();
    }

    return result;
}



int Part2(string[] input)
{
    var result = 0;

    var pattern = @"mul\((\d+),(\d+)\)";
    var multEnabled = true;

    var singleLine = string.Join("", input);

    var dontSegments = singleLine.Split("don't()", StringSplitOptions.RemoveEmptyEntries);

    var processedSegment = Regex.Matches(dontSegments[0], pattern);
    result += processedSegment.Select(x => int.Parse(x.Groups[1].Value) * int.Parse(x.Groups[2].Value)).Sum();

    for (var i = 1; i < dontSegments.Length; i++)
    {
        multEnabled = false;
        var doSegments = dontSegments[i].Split("do()", StringSplitOptions.RemoveEmptyEntries);

        foreach (var doSegment in doSegments)
        {
            if (multEnabled)
            {
                processedSegment = Regex.Matches(doSegment, pattern);
                result += processedSegment.Select(x => int.Parse(x.Groups[1].Value) * int.Parse(x.Groups[2].Value)).Sum();
            }

            multEnabled = true;
        }
    }

    return result;
}
