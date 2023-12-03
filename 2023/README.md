# 2023

## Language of the year
[Golang](https://go.dev/)

## Self-objectives

- Familiarize more with Golang specificities as I only have 1 year of experience with it.
- And of course having fun with other friends.

## Prerequisites

- **Go** (>= 1.21): https://go.dev/doc/install
- `make` if you want to use the provided `Makefile` (should work natively on GNU/Linux or MacOS, through WSL or Git Bash on Windows)

## How to run

First, fetch dependencies with `go mod tidy` or `make dependencies` to get everything needed for this year.

Each day is self contained (except the `Makefile` and the `go.mod` at the root of the year folder).

Use the `Makefile` at the root to build and execute:

- `make run/day-{dayNumber}/part{1|2}` will execute the specified day and the specified part

- `make run/day-{dayNumber}` will execute the specified day and all the parts

These commands will clean any previous binary, build the required binary, execute the program and clean the binary.

To run tests you can just use `make test/day-{dayNumber}` to run all the tests for a specific day.

Examples:
```
make run/day-1/part2 # Executes Day 1 Part 2
make run/day-2/part1 # Executes Day 2 Part 1
make run/day-3       # Executes all parts of Day 3

make test/day-2      # Executes all the tests for Day 2
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

## Notes

I have taken the decision to separate `part1` and `part2` for each puzzle so I could have clear separation and have the possibility to have 2 different approaches for each part if needed. Maybe next year I will try to combine solutions.