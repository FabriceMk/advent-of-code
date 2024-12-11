using System.Data;
using System.Diagnostics;
using System.Runtime.CompilerServices;
using System.Text.RegularExpressions;

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
    var dict = new Dictionary<int, List<int>>();
    var list = new List<List<int>>();

    var parseDict = true;

    foreach (var line in input)
    {
        if (line == string.Empty)
        {
            parseDict = false;
            continue;
        }

        if (parseDict)
        {
            var splitted = line.Split("|").Select(int.Parse).ToArray();
            if (!dict.TryGetValue(splitted[0], out var successorList))
            {
                dict[splitted[0]] = [];
            }

            dict[splitted[0]].Add(splitted[1]);
        }
        else
        {
            list.Add([.. line.Split(",").Select(int.Parse)]);
        }
    }

    return list.Where(x => isValidUpdate(x, dict)).Aggregate(0, (result, x) => result += x[x.Count / 2]);
}

bool isValidUpdate(List<int> update, Dictionary<int, List<int>> dict)
{
    var validUpdate = true;
    var pivotIndex = 0;

    while (pivotIndex < update.Count - 1 && validUpdate)
    {
        foreach (var before in update[0..pivotIndex])
        {
            if (dict.TryGetValue(update[pivotIndex], out var listAfter) && listAfter.Contains(before))
            {
                validUpdate = false;
                break;
            }
        }

        if (!validUpdate)
        {
            break;
        }

        foreach (var after in update[(pivotIndex + 1)..])
        {
            if (dict.TryGetValue(after, out var listBefore) && listBefore.Contains(update[pivotIndex]))
            {
                validUpdate = false;
                break;
            }
        }

        pivotIndex++;
    }

    return validUpdate;
}


int Part2(string[] input)
{
    var dict = new Dictionary<int, List<int>>();
    var list = new List<List<int>>();

    var parseDict = true;

    foreach (var line in input)
    {
        if (line == string.Empty)
        {
            parseDict = false;
            continue;
        }

        if (parseDict)
        {
            var splitted = line.Split("|").Select(int.Parse).ToArray();
            if (!dict.TryGetValue(splitted[0], out var successorList))
            {
                dict[splitted[0]] = [];
            }

            dict[splitted[0]].Add(splitted[1]);
        }
        else
        {
            list.Add([.. line.Split(",").Select(int.Parse)]);
        }
    }

    var result = list.Where(x => isValidUpdate(x, dict))
    .Select(x => reorderUpdate(x, dict))
    .Aggregate(0, (result, x) => result += x[x.Count / 2]);

    return result;
}

List<int> reorderUpdate(List<int> update, Dictionary<int, List<int>> dict)
{
    //return update.Sort(() => {})
}
