package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Input length: %d\n", len(input))
	missingSum := findMissingSum(input, 25) // 105950735
	fmt.Printf("Part 1: %d\n", missingSum)
	fmt.Printf("Part 2: %d\n", findXMASWeakness(input, missingSum))
}

func readInput() ([]int, error) {
	input := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		input = append(input, num)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return input, nil
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
