package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run for the hills.
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 254
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 184
}

func part1(inputPath string) int {
	valid := 0
	passport := make(map[string]string)
	scanner := helpers.NewScanner(inputPath)
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
	return valid
}

func isValid(p map[string]string) bool {
	// fmt.Printf("%v\n", p)
	valid := p["byr"] != "" &&
		p["iyr"] != "" &&
		p["eyr"] != "" &&
		p["hgt"] != "" &&
		p["hcl"] != "" &&
		p["ecl"] != "" &&
		p["pid"] != ""
	return valid
}

func part2(inputPath string) int {
	scanner := helpers.NewScanner(inputPath)
	passport := make(map[string]string)
	total := 0
	valid := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// fmt.Printf(">> %s\n", line)
		if len(line) == 0 {
			total++
			if isValid2(passport) {
				valid++
			}
			passport = make(map[string]string)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				kv := strings.Split(field, ":")
				_, present := passport[kv[0]]
				if present {
					panic("value already present!")
				}
				passport[kv[0]] = kv[1]
			}
		}
	}
	if len(passport) > 0 {
		total++
		if isValid2(passport) {
			valid++
		}
	}
	return valid
}

func isValid2(p map[string]string) bool {
	// fmt.Printf("byr: %s, iyr: %s, eyr: %s, hgt: %s, hcl: %s, ecl: %s, pid: %s\n",
	// p["byr"], p["iyr"], p["eyr"], p["hgt"], p["hcl"], p["ecl"], p["pid"])

	valid := true
	// fmt.Printf("pid: %s -- %v\n", p["pid"], isValidPassportID(p["pid"]))
	// return true
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !isValidYear(p["byr"], 1920, 2002) {
		// fmt.Printf("  invalid byr: %s\n", p["byr"])
		valid = false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !isValidYear(p["iyr"], 2010, 2020) {
		// fmt.Printf("  invalid iyr: %s\n", p["iyr"])
		valid = false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !isValidYear(p["eyr"], 2020, 2030) {
		// fmt.Printf("  invalid eyr: %s\n", p["eyr"])
		valid = false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if !isValidHeight(p["hgt"]) {
		// fmt.Printf("  invalid hgt: %s\n", p["hgt"])
		valid = false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if !isValidHairColor(p["hcl"]) {
		// fmt.Printf("  invalid hcl: %s\n", p["hcl"])
		valid = false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if !isValidEyeColor(p["ecl"]) {
		// fmt.Printf("  invalid ecl: %s\n", p["ecl"])
		valid = false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if !isValidPassportID(p["pid"]) {
		// fmt.Printf("  invalid pid: %s\n", p["pid"])
		valid = false
	}

	// cid (Country ID) - ignored, missing or not.
	return valid
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
	matched, err := regexp.MatchString(`^#[0-9a-f]{6}$`, hcl)
	if err != nil {
		panic(err)
	}
	return matched
}

func isValidEyeColor(ecl string) bool {
	eclMatch, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, ecl)
	if err != nil {
		panic(err)
	}
	return eclMatch
}

func isValidPassportID(id string) bool {
	pidMatch, err := regexp.MatchString(`^[0-9]{9}$`, id)
	if err != nil {
		panic(err)
	}
	return pidMatch
}
