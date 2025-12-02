package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_Part1(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"part1 should return correct result for test input": {
			input: test_input,
			want:  1227775554,
		},
	}

	for tn, tc := range testCases {
		assert := assert.New(t)

		t.Run(tn, func(t *testing.T) {
			// act
			res := part1(tc.input)

			// assert
			assert.Equal(tc.want, res)
		})
	}
}

func TestPuzzle_Part2(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"part2 should return correct result for test input 2": {
			input: test_input2,
			want:  4174379265,
		},
	}

	for tn, tc := range testCases {
		assert := assert.New(t)

		t.Run(tn, func(t *testing.T) {
			// act
			res := part2(tc.input)

			// assert
			assert.Equal(tc.want, res)
		})
	}
}
