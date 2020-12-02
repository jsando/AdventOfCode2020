package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go run main.go < input.txt
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	validCount := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line, " ")
		if len(fields) != 3 {
			panic(fmt.Sprintf("bad format '%s'", line))
		}
		minmax := strings.Split(fields[0], "-")
		if len(minmax) != 2 {
			panic(fmt.Sprintf("bad format '%s'", line))
		}
		// ya, no error handling ... bad elf
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		ch := fields[1][:1]
		pwd := fields[2]
		valid := isValid(pwd, ch, min, max)
		fmt.Printf("%d-%d %s: %s (valid: %v)\n", min, max, ch, pwd, valid)
		if valid {
			validCount++
		}
	}
	fmt.Printf("Valid passwords: %d\n", validCount)
}

// Count occurance of ch in pwd, if min <= count <= max return true
func isValid(pwd string, ch string, min, max int) bool {
	count := strings.Count(pwd, ch)
	if count >= min && count <= max {
		return true
	}
	return false
}
