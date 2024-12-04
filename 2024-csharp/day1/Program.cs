﻿using System.Diagnostics;

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

if (part == "1")
{
    res = Part1(input);
}
else
{
    res = Part2(input);
}

watch.Stop();

Console.WriteLine("##########");
Console.WriteLine($"Execution time: {watch.ElapsedMilliseconds}ms");
Console.WriteLine("Output:");
Console.WriteLine(res);
Console.WriteLine("##########");

int Part1(string[] input)
{
    var leftList = new List<int>();
    var rightList = new List<int>();

    foreach (var line in input)
    {
        var numbers = line.Split("   ", StringSplitOptions.RemoveEmptyEntries).Select(int.Parse).ToList();

        leftList.Add(numbers[0]);
        rightList.Add(numbers[1]);
    }

    leftList.Sort();
    rightList.Sort();

    var result = leftList.Zip(rightList, (x, y) => Math.Abs(x - y)).Sum();

    return result;
}

int Part2(string[] input)
{
    var leftList = new List<int>();
    var countingTable = new Dictionary<int, int>();

    foreach (var line in input)
    {
        var numbers = line.Split("   ", StringSplitOptions.RemoveEmptyEntries).Select(int.Parse).ToList();

        leftList.Add(numbers[0]);
        countingTable[numbers[1]] = countingTable.TryGetValue(numbers[1], out int occurrences) ? occurrences + 1 : 1;
    }

    var result = leftList.Where(countingTable.ContainsKey).Aggregate(0, (a,b) => a + (countingTable.TryGetValue(b, out int occurrences) ? b * occurrences : 0));

    return result;
}

