package day09

import (
	"fmt"

	"github.com/jsando/aoc2020/helpers"
)

// Run it all to heck.
func Run(inputPath string) {
	input := helpers.FileToIntSlice(inputPath)
	missingSum := findMissingSum(input, 25)
	fmt.Printf("Part 1: %d\n", missingSum)                          // 105950735
	fmt.Printf("Part 2: %d\n", findXMASWeakness(input, missingSum)) // 13826915
}

func findMissingSum(input []int, window int) int {
	for i := window; i < len(input); i++ {
		if !isSumWithinWindow(input[i-window:i], input[i]) {
			return input[i]
		}
	}
	panic("not found!")
}

func isSumWithinWindow(list []int, sum int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == sum {
				return true
			}
		}
	}
	return false
}

// Find a contiguous set of numbers that sum to the given sum and
// return the sum of the first and last in the set.
func findXMASWeakness(list []int, sum int) int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			// compute sum of numbers list[i ... j]
			test := 0
			min := 0
			max := 0
			for k := i; k < j; k++ {
				test += list[k]
				if list[k] < min || min == 0 {
					min = list[k]
				}
				if max < list[k] {
					max = list[k]
				}
			}
			if test == sum {
				//fmt.Printf("Sum to %d in range %d-%d, min = %d, max = %d\n", sum, i, j, min, max)
				return min + max
			}
		}
	}
	panic("didn't find it!")
}
