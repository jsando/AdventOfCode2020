package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go run puzzle1.go < input.txt
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	passport := make(map[string]string)
	valid := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if isValid(passport) {
				valid++
			}
			passport = make(map[string]string)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				kv := strings.Split(field, ":")
				passport[kv[0]] = kv[1]
			}
		}
	}
	if len(passport) > 0 {
		if isValid(passport) {
			valid++
		}
	}
	fmt.Printf("Valid passports: %d\n", valid)
}

func isValid(p map[string]string) bool {
	fmt.Printf("%v\n", p)
	valid := p["byr"] != "" &&
		p["iyr"] != "" &&
		p["eyr"] != "" &&
		p["hgt"] != "" &&
		p["hcl"] != "" &&
		p["ecl"] != "" &&
		p["pid"] != ""
	return valid
}
