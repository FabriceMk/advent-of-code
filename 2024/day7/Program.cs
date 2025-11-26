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

long Part1(string[] input)
{
    long result = 0;

    foreach (var line in input)
    {
        var splittedLine = line.Split(":").ToList();

        _ = long.TryParse(splittedLine[0], out var wantedRes);
        var operands = splittedLine[1].Trim().Split(" ", StringSplitOptions.TrimEntries).Select(int.Parse).ToList();

        result = isValid(wantedRes, 0, operands, false) ? result + wantedRes : result;
    }

    return result;
}

bool isValid(long wantedRes, long acc, List<int> operandsList, bool enableThirdOperator) {
    if (operandsList.Count == 0) {
        return wantedRes == acc;
    }

    if (operandsList.Count == 1) {
        return wantedRes == (acc + operandsList[0])
        || wantedRes == (acc * operandsList[0])
        || (enableThirdOperator && wantedRes == long.Parse(acc.ToString() + operandsList[0].ToString()));
    }

    return isValid(wantedRes, acc + operandsList[0], operandsList[1..], enableThirdOperator)
    || isValid(wantedRes, acc * operandsList[0], operandsList[1..], enableThirdOperator)
    || (enableThirdOperator && isValid(wantedRes, long.Parse(acc.ToString() + operandsList[0].ToString()), operandsList[1..], enableThirdOperator));
}

long Part2(string[] input)
{
    long result = 0;

    foreach (var line in input)
    {
        var splittedLine = line.Split(":").ToList();

        _ = long.TryParse(splittedLine[0], out var wantedRes);
        var operands = splittedLine[1].Trim().Split(" ", StringSplitOptions.TrimEntries).Select(int.Parse).ToList();

        result = isValid(wantedRes, 0, operands, true) ? result + wantedRes : result;
    }

    return result;
}
