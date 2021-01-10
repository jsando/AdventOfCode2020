package day17

import (
	"fmt"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run day 17.
func Run(inputPath string) {
	grid := loadGrid(inputPath)
	// grid.print()
	fmt.Printf("Part 1: %d\n", part1(grid)) // 276
	fmt.Printf("Part 2: %d\n", part2(grid)) // 2136
}

type grid interface {
	size() (x1, x2, y1, y2, z1, z2, w1, w2 int)
	set(x, y, z, w int, active bool)
	get(x, y, z, w int) bool
	neighbors(x, y, z, w int) int
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
	x1, x2, y1, y2, z1, z2, _, _ := input.size()
	output := newGrid()
	for z := z1 - 1; z <= z2+1; z++ {
		// fmt.Printf("Evaluating layer: %d\n", z)
		for y := y1 - 1; y <= y2+1; y++ {
			for x := x1 - 1; x <= x2+1; x++ {
				neighbors := input.neighbors(x, y, z, 0)
				active := input.get(x, y, z, 0)
				output.set(x, y, z, 0, active)
				// if active {
				// 	fmt.Printf("%2d* ", neighbors)
				// } else {
				// 	fmt.Printf("%2d  ", neighbors)
				// }
				if active && neighbors != 2 && neighbors != 3 {
					output.set(x, y, z, 0, false)
				}
				if !active && neighbors == 3 {
					output.set(x, y, z, 0, true)
				}
			}
			// fmt.Println()
		}
	}
	return output
}

func part2(grid grid) int {
	for cycle := 0; cycle < 6; cycle++ {
		grid = bootCycle2(grid)
		// fmt.Printf("\nAfter %d cycles:\n", (cycle + 1))
		// grid.print()
	}
	return grid.activeCount()
}

func bootCycle2(input grid) grid {
	x1, x2, y1, y2, z1, z2, w1, w2 := input.size()
	output := newGrid()
	for w := w1 - 1; w <= w2+1; w++ {
		for z := z1 - 1; z <= z2+1; z++ {
			// fmt.Printf("Evaluating layer: %d\n", z)
			for y := y1 - 1; y <= y2+1; y++ {
				for x := x1 - 1; x <= x2+1; x++ {
					neighbors := input.neighbors(x, y, z, w)
					active := input.get(x, y, z, w)
					output.set(x, y, z, w, active)
					// if active {
					// 	fmt.Printf("%2d* ", neighbors)
					// } else {
					// 	fmt.Printf("%2d  ", neighbors)
					// }
					if active && neighbors != 2 && neighbors != 3 {
						output.set(x, y, z, w, false)
					}
					if !active && neighbors == 3 {
						output.set(x, y, z, w, true)
					}
				}
				// fmt.Println()
			}
		}
	}
	return output
}

func loadGrid(inputPath string) grid {
	scanner := helpers.NewScanner(inputPath)
	y := 0
	grid := newGrid()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		for x, ch := range line {
			if ch == '#' {
				grid.set(x, y, 0, 0, true)
			}
		}
		y++
	}
	return grid
}

type point struct {
	x, y, z, w int
}

type gridImpl struct {
	active     []point
	xmin, xmax int
	ymin, ymax int
	zmin, zmax int
	wmin, wmax int
}

func (g *gridImpl) size() (x1, x2, y1, y2, z1, z2, w1, w2 int) {
	return g.xmin, g.xmax, g.ymin, g.ymax, g.zmin, g.zmax, g.wmin, g.wmax
}

func (g *gridImpl) getPointIndex(x, y, z, w int) int {
	for i, p := range g.active {
		if p.x == x && p.y == y && p.z == z && p.w == w {
			return i
		}
	}
	return -1
}

func (g *gridImpl) set(x, y, z, w int, active bool) {
	i := g.getPointIndex(x, y, z, w)

	if active && i == -1 {
		g.active = append(g.active, point{x, y, z, w})
		g.xmin = min(g.xmin, x)
		g.xmax = max(g.xmax, x)

		g.ymin = min(g.ymin, y)
		g.ymax = max(g.ymax, y)

		g.zmin = min(g.zmin, z)
		g.zmax = max(g.zmax, z)

		g.wmin = min(g.wmin, w)
		g.wmax = max(g.wmax, w)
	}

	// to delete active[i], move the last item over it and truncate the length
	if !active && i != -1 {
		g.active[i] = g.active[len(g.active)-1]
		g.active = g.active[:len(g.active)-1]
	}
}

func (g *gridImpl) get(x, y, z, w int) bool {
	i := g.getPointIndex(x, y, z, w)
	if i == -1 {
		return false
	}
	return true
}

func (g *gridImpl) neighbors(x, y, z, w int) int {
	count := 0
	for _, p := range g.active {
		// ignore self
		if x == p.x && y == p.y && z == p.z && w == p.w {
			continue
		}
		// if adjacent then count it
		if adjacent(x, p.x) && adjacent(y, p.y) && adjacent(z, p.z) && adjacent(w, p.w) {
			count++
		}
	}
	return count
}

func (g *gridImpl) print() {
	for w := g.wmin; w <= g.wmax; w++ {
		for z := g.zmin; z <= g.zmax; z++ {
			fmt.Printf("\nz=%d, w=%d\n", z, w)
			for y := g.ymin; y <= g.ymax; y++ {
				for x := g.xmin; x <= g.xmax; x++ {
					if g.get(x, y, z, w) {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Println()
			}
		}
	}
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
