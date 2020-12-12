package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// current ship position (starting at 0,0)
var x, y int

// current rotation 0=0, 1=90, 2=180, 3=270
var rot = 0

// dx,dy to move by based on current rotation
var vectors [][]int = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
			rotate(delta)
		case 'R':
			rotate(-delta)
		case 'F':
			vector := vectors[rot]
			x += (delta * vector[0])
			y += (delta * vector[1])
		}
		fmt.Printf("Command %c%d, ship now at (%d, %d) with heading %d\n", cmd, delta, x, y, rot)
	}
	fmt.Printf("Part 1: %d\n", abs(x)+abs(y)) // 1152
}

func rotate(degrees int) {
	rot += (degrees / 90)
	for rot < 0 {
		rot += 4
	}
	for rot > 3 {
		rot -= 4
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
