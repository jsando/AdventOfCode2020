package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	shipx     int = 0  // absolute x coordinate
	shipy     int = 0  // absoluite y coordinate
	waypointx int = 10 // relative to ship x coordinate
	waypointy int = 1  // relative to ship y coordinate
)

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
			waypointy += delta
		case 'S':
			waypointy -= delta
		case 'E':
			waypointx += delta
		case 'W':
			waypointx -= delta
		case 'L':
			rotate(delta)
		case 'R':
			rotate(360 - delta)
		case 'F':
			for i := 0; i < delta; i++ {
				shipx += waypointx
				shipy += waypointy
			}
		}
		fmt.Printf("Command %c%d, waypoint (%d, %d) ship (%d, %d)\n", cmd, delta, waypointx, waypointy, shipx, shipy)
	}
	fmt.Printf("Part 2: %d\n", abs(shipx)+abs(shipy)) //
}

// rotate the waypoint (around the ship) by the given number of degrees
// to keep it simple it uses the fact that a 90 degree rotation of (x,y) = (-y,x)
func rotate(degrees int) {
	if degrees <= 0 || degrees > 360 {
		panic("degrees out of range")
	}
	for degrees > 0 {
		waypointx, waypointy = -waypointy, waypointx
		degrees -= 90
	}
}

// yes ... go has no built-in abs(int) function
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
