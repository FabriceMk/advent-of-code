using System.Diagnostics;
using System.Security.AccessControl;

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
        string[] numbers = line.Split("   ");


        if (int.TryParse(numbers[0], out int leftNumber))
        {
            leftList.Add(leftNumber);
        }

        if (int.TryParse(numbers[1], out int rightNumber))
        {
            rightList.Add(rightNumber);
        }
    }

    leftList.Sort();
    rightList.Sort();

    var result = 0;

    for (var i = 0; i < leftList.Count; i++)
    {
        result += Math.Abs(rightList[i] - leftList[i]);
    }

    return result;
}

int Part2(string[] input)
{
    var leftList = new List<int>();
    var countingTable = new Dictionary<int, int>();

    foreach (var line in input)
    {
        string[] numbers = line.Split("   ");


        if (int.TryParse(numbers[0], out int leftNumber))
        {
            leftList.Add(leftNumber);
        }

        if (int.TryParse(numbers[1], out int rightNumber))
        {
            if (countingTable.ContainsKey(rightNumber))
            {
                countingTable[rightNumber]++;
            }
            else
            {
                countingTable[rightNumber] = 1;
            }
        }
    }

    var result = 0;

    foreach (var left in leftList)
    {
        if (countingTable.ContainsKey(left))
        {
            result += left * countingTable[left];
        }
    }

    return result;
}

