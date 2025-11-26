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
    var result = 0;

    // Failed solution attempt with regexp
    // var lineLength = input[0].Length;

    // var inputSingleLine = string.Join("", input);

    // var patterns = new string[] {
    //     @"XMAS",
    //     @"SAMX",
    //     $@"(?<=X.{{{lineLength - 1}}}M.{{{lineLength - 1}}}A.{{{lineLength - 1}}}S)",
    //     $@"(?<=S.{{{lineLength - 1}}}A.{{{lineLength - 1}}}M.{{{lineLength - 1}}}X)",
    //     $@"(?<=X.{{{lineLength}}}M.{{{lineLength}}}A.{{{lineLength}}}S)",
    //     $@"(?<=S.{{{lineLength}}}A.{{{lineLength}}}M.{{{lineLength}}}X)",
    //     $@"(?<=X.{{{lineLength - 2}}}M.{{{lineLength - 2}}}A.{{{lineLength - 2}}}S)",
    //     $@"(?<=S.{{{lineLength - 2}}}A.{{{lineLength - 2}}}M.{{{lineLength - 2}}}X)"
    // };

    // foreach (var pattern in patterns)
    // {
    //     result += Regex.Matches(inputSingleLine, pattern).Count;
    // }

    var matrix = new List<List<char>>();

    foreach (var line in input)
    {
        matrix.Add([.. line.ToCharArray()]);
    }

    var width = matrix[0].Count;
    var height = matrix.Count;

    for (var i = 0; i < height; i++)
    {
        for (var j = 0; j < width; j++)
        {
            if (matrix[i][j] != 'X')
            {
                continue;
            }

            if (j >= 3) // Check back
            {
                if (matrix[i][j - 1] == 'M' && matrix[i][j - 2] == 'A' && matrix[i][j - 3] == 'S')
                {
                    result++;
                }

                if (i >= 3)
                {
                    if (matrix[i - 1][j - 1] == 'M' && matrix[i - 2][j - 2] == 'A' && matrix[i - 3][j - 3] == 'S')
                    {
                        result++;
                    }
                }

                if (i < height - 3)
                {
                    if (matrix[i + 1][j - 1] == 'M' && matrix[i + 2][j - 2] == 'A' && matrix[i + 3][j - 3] == 'S')
                    {
                        result++;
                    }
                }


            }

            if (i >= 3) // Check up
            {
                if (matrix[i - 1][j] == 'M' && matrix[i - 2][j] == 'A' && matrix[i - 3][j] == 'S')
                {
                    result++;
                }
            }

            if (i < height - 3) // Check down
            {
                if (matrix[i + 1][j] == 'M' && matrix[i + 2][j] == 'A' && matrix[i + 3][j] == 'S')
                {
                    result++;
                }
            }

            if (j < width - 3) // Check forward
            {
                if (matrix[i][j + 1] == 'M' && matrix[i][j + 2] == 'A' && matrix[i][j + 3] == 'S')
                {
                    result++;
                }

                if (i >= 3)
                {
                    if (matrix[i - 1][j + 1] == 'M' && matrix[i - 2][j + 2] == 'A' && matrix[i - 3][j + 3] == 'S')
                    {
                        result++;
                    }
                }

                if (i < height - 3)
                {
                    if (matrix[i + 1][j + 1] == 'M' && matrix[i + 2][j + 2] == 'A' && matrix[i + 3][j + 3] == 'S')
                    {
                        result++;
                    }
                }
            }


        }
    }

    return result;
}



int Part2(string[] input)
{
    var result = 0;

    var matrix = new List<List<char>>();

    foreach (var line in input)
    {
        matrix.Add([.. line.ToCharArray()]);
    }

    var width = matrix[0].Count;
    var height = matrix.Count;

    for (var i = 1; i < height - 1; i++)
    {
        for (var j = 1; j < width - 1; j++)
        {
            if (matrix[i][j] != 'A')
            {
                continue;
            }

            if (matrix[i - 1][j - 1] == 'M' && matrix[i + 1][j + 1] == 'S')
            {
                if ((matrix[i + 1][j - 1] == 'M' && matrix[i - 1][j + 1] == 'S') ||
                (matrix[i + 1][j - 1] == 'S' && matrix[i - 1][j + 1] == 'M'))
                {
                    result++;
                }
            }

            if (matrix[i - 1][j - 1] == 'S' && matrix[i + 1][j + 1] == 'M')
            {
                if ((matrix[i + 1][j - 1] == 'M' && matrix[i - 1][j + 1] == 'S') ||
                (matrix[i + 1][j - 1] == 'S' && matrix[i - 1][j + 1] == 'M'))
                {
                    result++;
                }
            }
        }

    }

    return result;
}
