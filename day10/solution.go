package day10

import (
	"fmt"
	"sort"

	"github.com/jsando/aoc2020/helpers"
)

// Run the thing.
func Run(inputPath string) {
	input := helpers.FileToIntSlice(inputPath)
	sort.Ints(input)
	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)
	fmt.Printf("Part 1: %d\n", part1(input)) // 2450
	fmt.Printf("Part 2: %d\n", part2(input)) // 32396521357312s
}

func part1(input []int) int {
	ones := 0
	threes := 0
	last := input[0]
	for i := 1; i < len(input); i++ {
		this := input[i]
		diff := this - last
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		} else {
			panic("more than 1 or 3 gap")
		}
		last = this
	}
	return ones * threes
}

// walk backwards through the list of joltage adapters and compute
// how many combinations from that point to the end will work.
func part2(input []int) int {
	paths := make([]int, len(input))
	paths[len(input)-1] = 1 // from the next to last one, there's only 1 path to the last one.
	for i := len(input) - 2; i >= 0; i-- {
		for j := i + 1; j < len(input); j++ {
			distance := input[j] - input[i]
			if distance > 3 {
				break
			}
			paths[i] += paths[j]
		}
	}
	return paths[0]
}

// previous attempt to brute force 2^104 subsets omitted to spare me the embarassment
