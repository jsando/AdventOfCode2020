package day02

import (
	"fmt"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run run away.
func Run(inputPath string) {
	validCount := 0
	valid2Count := 0
	for _, line := range helpers.FileToStringSlice(inputPath) {
		var min, max int
		var ch byte
		var pwd string
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &ch, &pwd)
		if err != nil {
			panic(fmt.Sprintf("error %s: input string '%s'", err.Error(), line))
		}
		if isValid(pwd, ch, min, max) {
			validCount++
		}
		if isValid2(pwd, ch, min, max) {
			valid2Count++
		}
	}
	fmt.Printf("Part 1: %d\n", validCount)  // 398
	fmt.Printf("Part 2: %d\n", valid2Count) // 562
}

// Count occurance of ch in pwd, if min <= count <= max return true
func isValid(pwd string, ch byte, min, max int) bool {
	count := strings.Count(pwd, string(ch))
	if count >= min && count <= max {
		return true
	}
	return false
}

// Return true if exactly one of the min or max char positions contains the given ch.
func isValid2(pwd string, ch byte, min, max int) bool {
	// todo not valid for UTF8!
	ch1 := pwd[min-1]
	ch2 := pwd[max-1]
	if (ch1 == ch && ch2 != ch) || (ch1 != ch && ch2 == ch) {
		return true
	}
	return false
}
