package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	qty   int
	color string
}

var rules map[string][]rule = make(map[string][]rule)

// go run main.go < input.txt
func main() {
	// word word bags contain ("no other bags" | bag-spec [, bag spec]+) "."
	// bag-spec := quantity word word bag[s]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sa := strings.Split(line, " bags contain ")
		bag := strings.TrimSpace(sa[0])
		rules[bag] = make([]rule, 0)
		if sa[1] == "no other bags." {
			continue
		}
		contents := strings.Split(sa[1], ", ")
		// fmt.Printf("%s -> %v\n", bag, contents)
		for _, bagspec := range contents {
			fields := strings.Split(bagspec, " ")
			qty, err := strconv.Atoi(fields[0])
			if err != nil {
				panic(err)
			}
			color := fields[1] + " " + fields[2]
			rules[bag] = append(rules[bag], rule{qty: qty, color: color})
			// fmt.Printf("  %d %s\n", qty, color)
		}
	}
	fmt.Printf("Rule count: %d\n", len(rules))
	count := 0
	for k := range rules {
		if canContain(k, "shiny gold") {
			count++
		}
	}
	fmt.Printf("Part 1: bags that can somehow contain 'shiny gold' bags: %d\n", count)
	fmt.Printf("Part 2: number of individual bags within a 'shiny gold' bag: %d\n", countBags("shiny gold"))
}

func canContain(key string, target string) bool {
	for _, v := range rules[key] {
		if v.color == target || canContain(v.color, target) {
			return true
		}
	}
	return false
}

func countBags(key string) int {
	count := 0
	for _, v := range rules[key] {
		count = count + v.qty*(countBags(v.color)+1)
	}
	return count
}
