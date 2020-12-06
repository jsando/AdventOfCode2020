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
		answers := map[rune]bool{}
		for _, c := range group {
			if c != '\n' {
				answers[c] = true
			}
		}
		fmt.Printf("Group %d has %d answers.\n", i, len(answers))
		sum += len(answers)
	}
	fmt.Printf("Total groups: %d, sum of answers: %d\n", len(groups), sum)
}
