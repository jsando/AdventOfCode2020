package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := loadGrid()
	grid.print()
	fmt.Printf("part 1: %d\n", part1(grid))
}

type grid interface {
	size() (x1, x2, y1, y2, z1, z2 int)
	set(x, y, z int, active bool)
	get(x, y, z int) bool
	neighbors(x, y, z int) int
	activeCount() int
	print()
}

func newGrid() grid {
	return &gridImpl{}
}

func part1(grid grid) int {
	for cycle := 0; cycle < 6; cycle++ {
		grid = bootCycle(grid)
		// fmt.Printf("\nAfter %d cycles:\n", (cycle + 1))
		// grid.print()
	}
	return grid.activeCount()
}

func bootCycle(input grid) grid {
	x1, x2, y1, y2, z1, z2 := input.size()
	output := newGrid()
	for z := z1 - 1; z <= z2+1; z++ {
		// fmt.Printf("Evaluating layer: %d\n", z)
		for y := y1 - 1; y <= y2+1; y++ {
			for x := x1 - 1; x <= x2+1; x++ {
				neighbors := input.neighbors(x, y, z)
				active := input.get(x, y, z)
				output.set(x, y, z, active)
				// if active {
				// 	fmt.Printf("%2d* ", neighbors)
				// } else {
				// 	fmt.Printf("%2d  ", neighbors)
				// }
				if active && neighbors != 2 && neighbors != 3 {
					output.set(x, y, z, false)
				}
				if !active && neighbors == 3 {
					output.set(x, y, z, true)
				}
			}
			fmt.Println()
		}
	}
	return output
}

func loadGrid() grid {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	grid := newGrid()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		for x, ch := range line {
			if ch == '#' {
				grid.set(x, y, 0, true)
			}
		}
		y++
	}
	return grid
}

type point struct {
	x, y, z int
}

type gridImpl struct {
	active     []point
	xmin, xmax int
	ymin, ymax int
	zmin, zmax int
}

func (g *gridImpl) size() (x1, x2, y1, y2, z1, z2 int) {
	return g.xmin, g.xmax, g.ymin, g.ymax, g.zmin, g.zmax
}

func (g *gridImpl) getPointIndex(x, y, z int) int {
	for i, p := range g.active {
		if p.x == x && p.y == y && p.z == z {
			return i
		}
	}
	return -1
}

func (g *gridImpl) set(x, y, z int, active bool) {
	i := g.getPointIndex(x, y, z)

	if active && i == -1 {
		g.active = append(g.active, point{x, y, z})
		g.xmin = min(g.xmin, x)
		g.xmax = max(g.xmax, x)

		g.ymin = min(g.ymin, y)
		g.ymax = max(g.ymax, y)

		g.zmin = min(g.zmin, z)
		g.zmax = max(g.zmax, z)
	}

	// to delete active[i], move the last item over it and truncate the length
	if !active && i != -1 {
		g.active[i] = g.active[len(g.active)-1]
		g.active = g.active[:len(g.active)-1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g *gridImpl) get(x, y, z int) bool {
	i := g.getPointIndex(x, y, z)
	if i == -1 {
		return false
	}
	return true
}

func (g *gridImpl) neighbors(x, y, z int) int {
	count := 0
	for _, p := range g.active {
		// ignore self
		if x == p.x && y == p.y && z == p.z {
			continue
		}
		// if adjacent then count it
		if adjacent(x, p.x) && adjacent(y, p.y) && adjacent(z, p.z) {
			count++
		}
	}
	return count
}

func adjacent(a, b int) bool {
	v := b - a
	if v < 0 {
		v = -v
	}
	if v <= 1 {
		return true
	}
	return false
}

func (g *gridImpl) activeCount() int {
	return len(g.active)
}

func (g *gridImpl) print() {
	for z := g.zmin; z <= g.zmax; z++ {
		fmt.Printf("\nz=%d\n", z)
		for y := g.ymin; y <= g.ymax; y++ {
			for x := g.xmin; x <= g.xmax; x++ {
				if g.get(x, y, z) {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	}
}
