package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	groups := strings.Split(string(input), "\n\n")
	sum := 0
	for i, group := range groups {
		people := strings.Split(group, "\n")
		fmt.Printf("Group %d, people:\n%v\n", i, people)
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
		fmt.Printf("Group %d has %d common answers.\n\n", i, len(answers))
		sum += len(answers)
	}
	fmt.Printf("Total groups: %d, sum of answers: %d\n", len(groups), sum)
}
