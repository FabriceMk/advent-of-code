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
    int result = 0;

    var listKeys = new List<(int,int,int,int,int)>();
    var listLocks = new List<(int,int,int,int,int)>();

    var parsingKey = false;
    var currentlyParsing = false;
    var currentTuple = (0,0,0,0,0);

    foreach (var line in input)
    {
        if (line == string.Empty)
        {
            currentlyParsing = false;

            if (parsingKey)
            {
                // Reduce the height of 1 to take into account the '#####' floor for keys.
                listKeys.Add((currentTuple.Item1-1, currentTuple.Item2-1, currentTuple.Item3-1, currentTuple.Item4-1, currentTuple.Item5-1));
            } else {
                listLocks.Add((currentTuple.Item1, currentTuple.Item2, currentTuple.Item3, currentTuple.Item4, currentTuple.Item5));
            }

            currentTuple = (0,0,0,0,0);

            continue;
        }

        if (!currentlyParsing)
        {
            currentlyParsing = true;
            parsingKey = line == ".....";
            continue;
        }

        currentTuple.Item1 += line[0] == '#' ? 1 : 0;
        currentTuple.Item2 += line[1] == '#' ? 1 : 0;
        currentTuple.Item3 += line[2] == '#' ? 1 : 0;
        currentTuple.Item4 += line[3] == '#' ? 1 : 0;
        currentTuple.Item5 += line[4] == '#' ? 1 : 0;
    }

    Console.WriteLine("locks");
    foreach(var month in listLocks)
    {
        Console.WriteLine(month);
    }

    Console.WriteLine("keys");
    foreach(var month in listKeys)
    {
        Console.WriteLine(month);
    }

    foreach (var key in listKeys)
    {
        foreach (var currentLock in listLocks)
        {
            if (currentLock.Item1 + key.Item1 <= 5 &&
            currentLock.Item2 + key.Item2 <= 5 &&
            currentLock.Item3 + key.Item3 <= 5 &&
            currentLock.Item4 + key.Item4 <= 5 &&
            currentLock.Item5 + key.Item5 <= 5) {
                result += 1;
            }
        }
    }

    return result;
}

int Part2(string[] input)
{
    var result = 0;

    return result;
}
