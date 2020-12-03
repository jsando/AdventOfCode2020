package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type path struct {
	dx, dy int
	cx     int
	trees  int
}

// go run main.go < input.txt
func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
		fmt.Printf("dx: %d, dy: %d, count: %d\n", path.dx, path.dy, path.trees)
		multiple *= path.trees
	}
	fmt.Printf("Multiple: %d\n", multiple)
}
