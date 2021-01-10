package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run it up the flagpole.
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 913
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 717
}

func part1(inputPath string) int {
	max := 0
	scanner := helpers.NewScanner(inputPath)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		binary := strings.Map(func(r rune) rune {
			switch r {
			case 'F':
				return '0'
			case 'B':
				return '1'
			case 'L':
				return '0'
			case 'R':
				return '1'
			}
			return r
		}, line)
		seatID, err := strconv.ParseInt(binary, 2, 64)
		if err != nil {
			panic(err)
		}
		if int(seatID) > max {
			max = int(seatID)
		}
		// fmt.Printf("%s -> %s (%d)\n", line, binary, seatID)
	}
	return max
}

func part2(inputPath string) int {
	scanner := helpers.NewScanner(inputPath)
	occupied := make([]bool, 1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		binary := strings.Map(func(r rune) rune {
			switch r {
			case 'F':
				return '0'
			case 'B':
				return '1'
			case 'L':
				return '0'
			case 'R':
				return '1'
			}
			return r
		}, line)
		seatID, err := strconv.ParseInt(binary, 2, 64)
		if err != nil {
			panic(err)
		}
		occupied[seatID] = true
		// fmt.Printf("%s -> %s (%d)\n", line, binary, seatID)
	}
	mySeatID := 0
	for i := 0; i < len(occupied)-3; i++ {
		if occupied[i] && !occupied[i+1] && occupied[i+2] {
			mySeatID = i + 1
			break
		}
	}
	return mySeatID
}
