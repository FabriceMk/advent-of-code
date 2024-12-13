using System.Data;
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

(Dictionary<int, List<int>>,  List<List<int>>) parseInput(string[] input) {
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

    return (dict, list);
}

int Part1(string[] input)
{
    var inputs = parseInput(input);
    var dict = inputs.Item1;
    var list = inputs.Item2;

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
    var inputs = parseInput(input);
    var dict = inputs.Item1;
    var list = inputs.Item2;

    var comparer = new CompareByCustomDict() { CustomDict = dict};

    var result = list.Where(x => !isValidUpdate(x, dict))
    .Select(x => {
        x.Sort(comparer.Compare);
        return x;
    })
    .Aggregate(0, (result, x) => result += x[x.Count / 2]);

    return result;
}

class CompareByCustomDict : IComparer<int>
{
    public required Dictionary<int, List<int>> CustomDict {get; init;}

    public int Compare(int x1, int x2)
    {
        if (CustomDict.TryGetValue(x1, out var successorList) && successorList.Contains(x2)) {
            return -1;
        }

        if (CustomDict.TryGetValue(x2, out successorList) && successorList.Contains(x1)) {
            return 1;
        }

        return 0;
    }
}