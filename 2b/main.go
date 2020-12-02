package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go run main.go < ../2a/input.txt
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
		ch := fields[1][0]
		pwd := fields[2]
		valid := isValid(pwd, ch, min, max)
		fmt.Printf("%d-%d %c: %s (valid: %v)\n", min, max, ch, pwd, valid)
		if valid {
			validCount++
		}
	}
	fmt.Printf("Valid passwords: %d\n", validCount)
}

// Return true if exactly one of the min or max char positions contains the given ch.
func isValid(pwd string, ch byte, min, max int) bool {
	// todo not valid for UTF8!
	ch1 := pwd[min-1]
	ch2 := pwd[max-1]
	if (ch1 == ch && ch2 != ch) || (ch1 != ch && ch2 == ch) {
		return true
	}
	return false
}
