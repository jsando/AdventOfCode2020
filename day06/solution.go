package day06

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Run back to me.
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 7283
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 3520
}

func part1(inputPath string) int {
	input, _ := ioutil.ReadFile(inputPath)
	groups := strings.Split(string(input), "\n\n")
	sum := 0
	for _, group := range groups {
		answers := map[rune]bool{}
		for _, c := range group {
			if c != '\n' {
				answers[c] = true
			}
		}
		// fmt.Printf("Group %d has %d answers.\n", i, len(answers))
		sum += len(answers)
	}
	return sum
}

func part2(inputPath string) int {
	input, _ := ioutil.ReadFile(inputPath)
	groups := strings.Split(string(input), "\n\n")
	sum := 0
	for _, group := range groups {
		people := strings.Split(group, "\n")
		// fmt.Printf("Group %d, people:\n%v\n", i, people)
		answers := map[rune]bool{}
		// For each answer in person 1's answer, see if the others in their group answered the same
		for _, c := range people[0] {
			allAnswered := true
			for i := 1; i < len(people); i++ {
				if len(people[i]) == 0 {
					continue // ignore blank lines included by srings.split
				}
				if !strings.ContainsRune(people[i], c) {
					allAnswered = false
					break
				}
			}
			if allAnswered {
				answers[c] = true
			}
		}
		// fmt.Printf("Group %d has %d common answers.\n\n", i, len(answers))
		sum += len(answers)
	}
	return sum
}
