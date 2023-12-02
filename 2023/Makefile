.PHONY: dependencies build/day-% clean/day-%
dependencies:
	go mod tidy

build/day-%:
	go build -o ./day-$*/puzzle ./day-$*/puzzle.go

clean/day-%:
	rm -f ./day-$*/puzzle

.PHONY: run/day-% run/day-%/part1 run/day-%/part2
run/day-%: clean/day-% build/day-%
	./day-$*/puzzle -part=1 && ./day-$*/puzzle -part=2 && rm -f ./day-$*/puzzle

run/day-%/part1: clean/day-% build/day-%
	./day-$*/puzzle -part=1 && rm -f ./day-$*/puzzle

run/day-%/part2: clean/day-% build/day-%
	./day-$*/puzzle -part=2 && rm -f ./day-$*/puzzle

.PHONY: test/day-%
test/day-%:
	go test ./day-$*/ -v
