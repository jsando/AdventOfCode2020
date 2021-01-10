package day01

import (
	"fmt"

	"github.com/jsando/aoc2020/helpers"
)

// Run run reindeer.
func Run(inputPath string) {
	entries := helpers.FileToIntSlice(inputPath)
	part1 := 0
	part2 := 0
	for i := 0; i < len(entries); i++ {
		a := entries[i]
		for j := i + 1; j < len(entries); j++ {
			b := entries[j]
			if a+b == 2020 {
				part1 = a * b
			}
			for k := j + 1; k < len(entries); k++ {
				c := entries[k]
				if a+b+c == 2020 {
					part2 = a * b * c
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part1) // 542619
	fmt.Printf("Part 2: %d\n", part2) // 32858450
}
