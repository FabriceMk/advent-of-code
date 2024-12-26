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
    var iterations = 2000;

    foreach (var line in input)
    {
        long.TryParse(line, out var secretNumber);

        for (var i = 0; i < iterations; i++)
        {
            secretNumber = nextSecret(secretNumber);
        }

       result += secretNumber;
    }

    return result;
}

long nextSecret(long secretNumber)
{
    var temp = secretNumber * 64;
    secretNumber ^= temp;
    secretNumber %= 16777216;
    temp = (long) Math.Floor((double) secretNumber/32);
    secretNumber ^= temp;
    secretNumber %= 16777216;
    temp = secretNumber * 2048;
    secretNumber ^= temp;
    secretNumber %= 16777216;

    return secretNumber;
}


int Part2(string[] input)
{
    var result = 0;

    var iterations = 2000;

    var currentSeller = 0;
    var sequences = new Dictionary<(int, int, int, int), Dictionary<int,int>>();

    foreach (var line in input)
    {
        currentSeller++;

        long.TryParse(line, out var secretNumber);

        var lastChanges = new Queue<int>();

        secretNumber = nextSecret(secretNumber);
        var previousSellingPrice = (int) secretNumber % 10;

        for (var i = 1; i < iterations; i++)
        {
            secretNumber = nextSecret(secretNumber);
            var currentSellingPrice = (int) secretNumber % 10;

            lastChanges.Enqueue(currentSellingPrice - previousSellingPrice);
            previousSellingPrice = currentSellingPrice;

            if (lastChanges.Count > 4) {
                lastChanges.Dequeue();
            }

            if (lastChanges.Count < 4) {
                continue;
            }

            var sequence = (lastChanges.ElementAt(0), lastChanges.ElementAt(1), lastChanges.ElementAt(2), lastChanges.ElementAt(3));

            var foundExistingSequence = sequences.TryGetValue(sequence, out var _);

            if (foundExistingSequence) {
                var foundPriceForSellerSequence = sequences[sequence].TryGetValue(currentSeller, out var _);
                if (!foundPriceForSellerSequence) {
                    sequences[sequence][currentSeller] = currentSellingPrice;
                }
            } else {
                sequences[sequence] = new Dictionary<int, int> {{ currentSeller, currentSellingPrice }};
            }
        }
    }

    foreach (var sequence in sequences)
    {
        var maxBananasForSequence = sequence.Value.Values.Sum();
        if (maxBananasForSequence > result)
        {
            result = maxBananasForSequence;
        }
    }

    return result;
}
