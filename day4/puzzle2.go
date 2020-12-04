package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	// fmt.Printf("%v\n", p)

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !isValidYear(p["byr"], 1920, 2002) {
		fmt.Printf("invalid byr: %s\n", p["byr"])
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !isValidYear(p["iyr"], 2010, 2020) {
		fmt.Printf("invalid iyr: %s\n", p["iyr"])
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !isValidYear(p["eyr"], 2020, 2030) {
		fmt.Printf("invalid eyr: %s\n", p["eyr"])
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if !isValidHeight(p["hgt"]) {
		fmt.Printf("invalid hgt: %s\n", p["hgt"])
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if !isValidHairColor(p["hcl"]) {
		fmt.Printf("invalid hcl: %s\n", p["hcl"])
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eclMatch, err := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, p["ecl"])
	if err != nil {
		panic(err)
	}
	if !eclMatch {
		fmt.Printf("invalid ecl: %s\n", p["ecl"])
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pidMatch, err := regexp.MatchString(`[0-9]{9}`, p["pid"])
	if err != nil {
		panic(err)
	}
	if !pidMatch {
		fmt.Printf("invalid pid: %s\n", p["pid"])
		return false
	}

	// cid (Country ID) - ignored, missing or not.
	return true
}

// Valid year ... 4 digits, between min and max inclusive.
func isValidYear(year string, min, max int) bool {
	if len(year) != 4 {
		return false
	}
	n, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return min <= n && n <= max
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func isValidHeight(h string) bool {
	if len(h) < 3 {
		return false
	}
	units := h[len(h)-2:]
	h = h[:len(h)-2]
	height, err := strconv.Atoi(h)
	if err != nil {
		return false
	}
	switch units {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	default:
		return false
	}
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func isValidHairColor(hcl string) bool {
	if len(hcl) != 7 {
		return false
	}
	matched, err := regexp.MatchString(`#[0-9a-f]{6}`, hcl)
	if err != nil {
		panic(err)
	}
	return matched
}
