# This Makefile should work on Linux, Mac and Windows (with MinGW)

# Detect OS for executable extension only
ifeq ($(OS),Windows_NT)
    EXE := .exe
else
    EXE :=
endif

MAKEFLAGS += --no-print-directory

.PHONY: dependencies build/day-% clean/day-%
dependencies:
	go mod tidy

build/day-%:
	go build -o ./day-$*/puzzle$(EXE) ./day-$*/puzzle.go

clean/day-%:
	-rm -f ./day-$*/puzzle$(EXE)

.PHONY: run/day-% run/day-%/part1 run/day-%/part2
run/day-%: clean/day-% build/day-%
	./day-$*/puzzle$(EXE) -part=1 && ./day-$*/puzzle$(EXE) -part=2 && rm -f ./day-$*/puzzle$(EXE)

run/day-%/part1:
	$(MAKE) clean/day-$*
	$(MAKE) build/day-$*
	./day-$*/puzzle$(EXE) -part=1 && rm -f ./day-$*/puzzle$(EXE)

run/day-%/part2:
	$(MAKE) clean/day-$*
	$(MAKE) build/day-$*
	./day-$*/puzzle$(EXE) -part=2 && rm -f ./day-$*/puzzle$(EXE)

.PHONY: test/day-%
test/day-%:
	go test ./day-$*/ -v
