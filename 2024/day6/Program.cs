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
    var mapArea = new List<List<char>>();
    var guardPosition = (X: -1, Y: -1);
    var currentVectorDirection = VectorDirection.Up;

    var visitedTiles = new HashSet<(int, int)>();

    var guardFound = false;
    int yCoord = 0;

    foreach (var line in input)
    {
        mapArea.Add([.. line.ToCharArray()]);
        if (!guardFound)
        {
            guardPosition.X = line.IndexOf('^');

            if (guardPosition.X > -1)
            {
                guardPosition.Y = yCoord;
                guardFound = true;
            }
        }

        yCoord++;
    }

    var width = mapArea[0].Count;
    var height = mapArea.Count;

    while (true)
    {
        var (aheadX, aheadY) = GetVector(currentVectorDirection);

        if (guardPosition.Y + aheadY < 0 || guardPosition.Y + aheadY >= height
        || guardPosition.X + aheadX < 0 || guardPosition.X + aheadX >= width
        )
        {
            visitedTiles.Add(guardPosition);
            break;
        }

        var tileAhead = mapArea[guardPosition.Y + aheadY][guardPosition.X + aheadX];

        if (tileAhead == '#')
        {
            currentVectorDirection = (VectorDirection)(((int)currentVectorDirection + 1) % 4);
        }
        else
        {
            visitedTiles.Add(guardPosition);
            guardPosition.X += aheadX;
            guardPosition.Y += aheadY;
        }
    }

    return visitedTiles.Count;
}

int Part2(string[] input)
{
    var result = 0;

    return result;
}

(int X, int Y) GetVector(VectorDirection vectDir)
{
    return vectDir switch
    {
        VectorDirection.Up => (0, -1),
        VectorDirection.Right => (1, 0),
        VectorDirection.Down => (0, 1),
        VectorDirection.Left => (-1, 0),
        _ => throw new Exception("unknown vector, that should not happen"),
    };
}

enum VectorDirection
{
    Up = 0,
    Right = 1,
    Down = 2,
    Left = 3,
}
