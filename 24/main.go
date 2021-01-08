package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

func main() {
	tiles := map[image.Point]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		var point image.Point
		for i := 0; i < len(line); i++ {
			var dx, dy int
			if line[i] == 'w' {
				dx = -1
			} else if line[i] == 'e' {
				dx = 1
			} else {
				switch line[i : i+2] {
				case "nw":
					dy = 1
				case "ne":
					dx = 1
					dy = 1
				case "sw":
					dx = -1
					dy = -1
				case "se":
					dy = -1
				}
				i++
			}
			point.X += dx
			point.Y += dy
		}
		tiles[point] = !tiles[point]
		color := "white"
		if tiles[point] {
			color = "black"
		}
		fmt.Printf("%s is now %s\n", point, color)
	}
	count := countBlack(tiles)

	fmt.Printf("Part 1: %d\n", count)        // 230
	fmt.Printf("Part 2: %d\n", part2(tiles)) // 3565
}

func part2(tiles map[image.Point]bool) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	x1 := math.MaxInt64
	y1 := math.MaxInt64
	x2 := math.MinInt64
	y2 := math.MinInt64
	for v := range tiles {
		x1 = min(x1, v.X)
		y1 = min(y1, v.Y)
		x2 = max(x2, v.X)
		y2 = max(y2, v.Y)
	}
	for day := 0; day < 100; day++ {
		x1--
		y1--
		x2++
		y2++
		newMap := map[image.Point]bool{}
		pointsCovered := 0
		flipped := 0
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				pointsCovered++
				p := image.Point{x, y}
				blackAdjacent := countAdjacent(tiles, x, y)
				newMap[p] = tiles[p]
				if tiles[p] {
					if blackAdjacent == 0 || blackAdjacent > 2 {
						newMap[p] = false
						flipped++
					}
				} else {
					if blackAdjacent == 2 {
						newMap[p] = true
						flipped++
					}
				}
			}
		}
		fmt.Printf("Day %d: %d (%d flipped)\n", day+1, countBlack(newMap), flipped)
		tiles = newMap
	}
	return countBlack(tiles)
}

var adjacent []image.Point = []image.Point{
	{0, 1},   // nw
	{1, 1},   // ne
	{-1, 0},  // w
	{1, 0},   // e
	{-1, -1}, // sw
	{0, -1},  // se
}

func countAdjacent(tiles map[image.Point]bool, x, y int) int {
	bc := 0
	for _, p := range adjacent {
		loc := image.Point{p.X + x, p.Y + y}
		if tiles[loc] {
			bc++
		}
	}
	return bc
}

func countBlack(tiles map[image.Point]bool) int {
	count := 0
	for _, black := range tiles {
		if black {
			count++
		}
	}
	return count
}
