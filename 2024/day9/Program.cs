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

    // Parse
    var diskRep = input[0].ToCharArray().Select(x => int.Parse(x.ToString())).ToArray();
    var diskSize = diskRep.Sum();
    var disk = new int[diskSize];

    ParseDiskRep(diskRep, disk);

    // Defrag
    var currentBlockIndex = 0;
    var detectedEnd = false;
    for (var i = diskSize - 1; i >= 0; i--)
    {
        if (disk[i] == -1)
        {
            continue;
        }

        while (disk[currentBlockIndex] != -1)
        {
            currentBlockIndex++;

            if (currentBlockIndex >= i)
            {
                detectedEnd = true;
                break;
            }
        }

        if (detectedEnd)
        {
            break;
        }

        disk[currentBlockIndex] = disk[i];
        disk[i] = -1;
    }

    // Checksum
    result = Checksum(disk);

    return result;
}

void ParseDiskRep(int[] diskRep, int[] inputDisk)
{
    var acc = 0;

    for (var blockIndex = 0; blockIndex < diskRep.Length; blockIndex++)
    {
        var number = diskRep[blockIndex];

        for (var blockScanIndex = 0; blockScanIndex < number; blockScanIndex++)
        {
            inputDisk[acc + blockScanIndex] = blockIndex % 2 == 0 ? blockIndex / 2 : -1;
        }

        acc += number;
    }
}

long Checksum(int[] inputDisk)
{
    long result = 0;

    for (var i = 0; i < inputDisk.Length; i++)
    {
        if (inputDisk[i] == -1)
        {
            continue;
        }

        result += inputDisk[i] * i;
    }

    return result;
}

long Part2(string[] input)
{
    long result = 0;

    // Parse
    var diskRep = input[0].ToCharArray().Select(x => int.Parse(x.ToString())).ToArray();
    var diskSize = diskRep.Sum();
    var disk = new int[diskSize];

    ParseDiskRep(diskRep, disk);

    // Defrag
    var currentBlockIndex = 0;
    var detectedEnd = false;
    for (var i = diskSize - 1; i >= 0; i--)
    {
        if (disk[i] == -1)
        {
            continue;
        }

        while (disk[currentBlockIndex] != -1)
        {
            currentBlockIndex++;

            if (currentBlockIndex > diskSize || currentBlockIndex >= i)
            {
                detectedEnd = true;
                break;
            }
        }

        if (detectedEnd)
        {
            break;
        }

        disk[currentBlockIndex] = disk[i];
        disk[i] = -1;
    }

    // Checksum
    result = Checksum(disk);

    return result;
}
