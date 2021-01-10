package day03

import (
	"fmt"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run run away.
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 207
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 2655892800
}

func part1(inputPath string) int {
	treeCount := 0
	scanner := helpers.NewScanner(inputPath)
	scanner.Scan() // skip first row
	x := 3
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		if x >= len(line) {
			x -= len(line)
		}
		if line[x] == '#' {
			treeCount++
		}
		x += 3
	}
	return treeCount
}

type path struct {
	dx, dy int
	cx     int
	trees  int
}

func part2(inputPath string) int {
	scanner := helpers.NewScanner(inputPath)
	paths := []*path{
		{dx: 1, dy: 1},
		{dx: 3, dy: 1},
		{dx: 5, dy: 1},
		{dx: 7, dy: 1},
		{dx: 1, dy: 2},
	}
	scanner.Scan() // skip row zero
	cy := 0
	for scanner.Scan() {
		cy++
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		max := len(line)
		for _, path := range paths {
			if cy%path.dy != 0 {
				continue
			}
			path.cx += path.dx
			if path.cx >= max {
				path.cx -= max
			}
			if line[path.cx] == '#' {
				path.trees++
			}
		}
	}
	multiple := 1
	for _, path := range paths {
		// fmt.Printf("dx: %d, dy: %d, count: %d\n", path.dx, path.dy, path.trees)
		multiple *= path.trees
	}
	return multiple
}
