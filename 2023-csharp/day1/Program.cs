using System.Diagnostics;

// Extract argument to know which part to run
var part = args.FirstOrDefault();

string[] input = File.ReadAllLines(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "input-full"));

Console.WriteLine("##########");
Console.WriteLine($"Running part {part}");

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
Console.WriteLine($"Execution time: {watch.ElapsedMilliseconds}");
Console.WriteLine("Output:");
Console.WriteLine(res);
Console.WriteLine("##########");

int Part1(string[] input)
{
    // TODO
    return input.Length;
}

int Part2(string[] input)
{
    // TODO
    return input.Length;
}

