# 2024-csharp

## Language of the year
[C#](https://dotnet.microsoft.com/en-us/languages/csharp)

## Prerequisites

- **.NET 8**: https://dotnet.microsoft.com/en-us/download/dotnet/8.0

## How to run

You need choose which part to run and if you want to run against the test input or the full input.

By default the program runs against the full input file. By adding `test`, you can run it against the test input.


```
dotnet run --project /day{dayNumber}/day{dayNumber}.csproj {1|2} {test}
dotnet run --project /day4/day4.csproj 1   # Runs part 1 of the day 4 against input-full
dotnet run --project /day6/day6.csproj 2 test  # Runs part 2 of the day 6 against input-test
```

or directly from one of the day subfolders
```
dotnet run {1|2} {test}
dotnet run 1      # Runs part 1 of the day against the full input
dotnet run 2 test # Runs part 2 of the day against the test input
```

## Self-imposed rules
- No external dependency (except for the test files).
- Use less regexp as possible, this aspect is not the one that interests me this year.
- Try to avoid cryptic hacks when possible like bit flipping and other kind of Dark Magic.
- Depending on my mood, each puzzle can have different self-imposed rules like:
  - Trying to avoid too many nested loops
  - Solve the problem in one pass when parsing the input
  - Try to go for a fast solution even if it has more hardcoded stuff
  - Go for a simple and elegant solution that puts emphasis on future collaboration and is easy-to-read
  - etc...