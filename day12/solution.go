package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run florest, run.
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 1152
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 58637
}

// dx,dy to move by based on current rotation
var vectors [][]int = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func part1(inputPath string) int {
	var x, y int // current ship position (starting at 0,0)
	rot := 0     // current rotation 0=0, 1=90, 2=180, 3=270
	scanner := helpers.NewScanner(inputPath)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		cmd := line[0]
		delta, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch cmd {
		case 'N':
			y -= delta
		case 'S':
			y += delta
		case 'E':
			x += delta
		case 'W':
			x -= delta
		case 'L':
			rot = rotate(rot, delta)
		case 'R':
			rot = rotate(rot, -delta)
		case 'F':
			vector := vectors[rot]
			x += (delta * vector[0])
			y += (delta * vector[1])
		}
		// fmt.Printf("Command %c%d, ship now at (%d, %d) with heading %d\n", cmd, delta, x, y, rot)
	}
	return helpers.AbsInt(x) + helpers.AbsInt(y)
}

func rotate(rot, degrees int) int {
	rot += (degrees / 90)
	for rot < 0 {
		rot += 4
	}
	for rot > 3 {
		rot -= 4
	}
	return rot
}

func part2(inputPath string) int {
	shipx := 0      // absolute x coordinate
	shipy := 0      // absoluite y coordinate
	waypointx := 10 // relative to ship x coordinate
	waypointy := 1  // relative to ship y coordinate
	scanner := helpers.NewScanner(inputPath)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		cmd := line[0]
		delta, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		degrees := 0
		switch cmd {
		case 'N':
			waypointy += delta
		case 'S':
			waypointy -= delta
		case 'E':
			waypointx += delta
		case 'W':
			waypointx -= delta
		case 'L':
			degrees = delta
		case 'R':
			degrees = 360 - delta
		case 'F':
			for i := 0; i < delta; i++ {
				shipx += waypointx
				shipy += waypointy
			}
		}
		for degrees > 0 {
			waypointx, waypointy = -waypointy, waypointx
			degrees -= 90
		}
		// fmt.Printf("Command %c%d, waypoint (%d, %d) ship (%d, %d)\n", cmd, delta, waypointx, waypointy, shipx, shipy)
	}
	return helpers.AbsInt(shipx) + helpers.AbsInt(shipy)
}
