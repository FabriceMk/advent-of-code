using System.Diagnostics;
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
