# 2025

## Language of the year

[Golang](https://go.dev/)

## Self-objectives

- Familiarize with TUI (Terminal UI) framework, especially [BubbleTea](https://github.com/charmbracelet/bubbletea) for Golang.
- Discover the [Charm](https://charm.land/) ecosystem (it looks cool).

## Prerequisites

- **Go** (>= 1.25): <https://go.dev/doc/install>
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

```sh
make run/day-1/part2 # Executes Day 1 Part 2
make run/day-2/part1 # Executes Day 2 Part 1
make run/day-3       # Executes all parts of Day 3

make test/day-2      # Executes all the tests for Day 2
```

## Self-imposed rules

- Nothing special this year as I won't have much time. But focus should be on using BubbleTea and have visualizations.

## Notes

I have taken the decision to separate `part1` and `part2` for each puzzle so I could have clear separation and have the possibility to have 2 different approaches for each part if needed. Maybe next year I will try to combine solutions.
