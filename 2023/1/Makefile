puzzle:
	go build -o puzzle puzzle.go

.PHONY: run1 run2 clean

run1: clean puzzle
	./puzzle -part=1 && rm -f puzzle

run2: clean puzzle
	./puzzle -part=2 && rm -f puzzle

clean:
	rm -f puzzle