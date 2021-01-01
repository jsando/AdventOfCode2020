package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	tiles := parseTiles(string(bytes))
	fmt.Printf("Read %d tiles.\n", len(tiles))

	s := newSolver(tiles)
	image := s.solve()
	//s.print()
	part1 := image[0][0].id * image[0][11].id * image[11][0].id * image[11][11].id
	fmt.Printf("Part 1: %d\n", part1) // 29584525501199
}

type tile struct {
	id     int
	pixels []int
	left   *tile
	top    *tile
	right  *tile
	bottom *tile
}

func parseTiles(input string) []*tile {
	var tiles []*tile
	for _, block := range strings.Split(strings.TrimSpace(input), "\n\n") {
		tiles = append(tiles, parseTile(block))
	}
	return tiles
}

func parseTile(input string) *tile {
	tile := &tile{}
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Tile ") {
			id, err := strconv.Atoi(strings.TrimSuffix(line[5:], ":"))
			if err != nil {
				panic(err)
			}
			tile.id = id
		} else {
			binary := strings.ReplaceAll(line, ".", "0")
			binary = strings.ReplaceAll(binary, "#", "1")
			t, err := strconv.ParseInt(binary, 2, 64)
			if err != nil {
				panic(err)
			}
			pixel := int(t)
			tile.pixels = append(tile.pixels, pixel)
		}
	}
	return tile
}

func (t *tile) String() string {
	pixels := ""
	for i := 0; i < len(t.pixels); i++ {
		pixels += fmt.Sprintf("%010b\n", t.pixels[i])
	}
	s := fmt.Sprintf("Tile %d:\n%s\n", t.id, pixels)
	if t.top != nil {
		s += fmt.Sprintf("Top: %d ", t.top.id)
	}
	if t.right != nil {
		s += fmt.Sprintf("Right: %d ", t.right.id)
	}
	if t.bottom != nil {
		s += fmt.Sprintf("Bottom: %d ", t.bottom.id)
	}
	if t.left != nil {
		s += fmt.Sprintf("Left: %d ", t.left.id)
	}
	s += "\n"
	return s
}

func (t *tile) rotate() {
	rotPixels := make([]int, len(t.pixels))
	bit := 0b_10_0000_0000
	for i := 0; i < len(rotPixels); i++ {
		val := 0
		for j := len(t.pixels) - 1; j >= 0; j-- {
			val <<= 1
			if t.pixels[j]&bit != 0 {
				val |= 1
			}
		}
		rotPixels[i] = val
		bit >>= 1
	}
	t.pixels = rotPixels
}

func (t *tile) vFlip() {
	flipPixels := make([]int, len(t.pixels))
	for i := 0; i < len(t.pixels); i++ {
		flipPixels[i] = t.pixels[len(t.pixels)-1-i]
	}
	t.pixels = flipPixels
}

func (t *tile) hFlip() {
	flipPixels := make([]int, len(t.pixels))
	for i := 0; i < len(t.pixels); i++ {
		bits := t.pixels[i]
		val := 0
		for j := 0; j < len(t.pixels); j++ {
			val <<= 1
			val |= bits & 0x01
			bits >>= 1
		}
		flipPixels[i] = val
	}
	t.pixels = flipPixels
}

type solver struct {
	tiles  []*tile
	toLink map[int]*tile
	unused []*tile
	placed []*tile
	image  [][]*tile
}

func newSolver(tiles []*tile) *solver {
	s := &solver{
		tiles:  tiles,
		toLink: make(map[int]*tile),
		unused: make([]*tile, len(tiles)-1),
		placed: []*tile{tiles[0]},
	}
	copy(s.unused, tiles[1:])
	return s
}

func (s *solver) print() {
	for _, row := range s.image {
		for _, tile := range row {
			fmt.Printf("Tile: %-5d ", tile.id)
		}
		fmt.Println()
		for i := 0; i < len(s.tiles[0].pixels); i++ {
			for _, tile := range row {
				fmt.Printf("%010b  ", tile.pixels[i])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (s *solver) solve() [][]*tile {
	current := s.tiles[0]
	for current != nil {
		if current.top == nil {
			s.uberMatch(current, "top")
		}
		if current.right == nil {
			s.uberMatch(current, "right")
		}
		if current.bottom == nil {
			s.uberMatch(current, "bottom")
		}
		if current.left == nil {
			s.uberMatch(current, "left")
		}
		delete(s.toLink, current.id)
		current = nil
		for _, v := range s.toLink {
			current = v
			break
		}
	}
	s.image = convertToArray(s.tiles[0])
	return s.image
}

func convertToArray(rando *tile) [][]*tile {
	// find the top left then walk the graph into an array
	var image [][]*tile
	var piece *tile
	for piece = rando; piece.left != nil; piece = piece.left {
	}
	for ; piece.top != nil; piece = piece.top {
	}
	for ; piece != nil; piece = piece.bottom {
		var row []*tile
		for current := piece; current != nil; current = current.right {
			row = append(row, current)
		}
		image = append(image, row)
	}
	return image
}

func (s *solver) uberMatch(current *tile, side string) {
	var compare compareFunc
	switch side {
	case "top":
		compare = compareTopToBottom
	case "right":
		compare = compareRightToLeft
	case "bottom":
		compare = compareBottomToTop
	case "left":
		compare = compareLeftToRight
	}
	var matchPiece *tile
	match := matchSide(s.unused, current, compare)
	if match != -1 {
		matchPiece = s.unused[match]
		s.placed = append(s.placed, matchPiece)
		s.unused = append(s.unused[:match], s.unused[match+1:]...)
		s.toLink[matchPiece.id] = matchPiece
	} else {
		// search the already placed pile, without rotating/flipping anything
		matchPiece = matchPlaced(s.placed, current, compare)
	}
	switch side {
	case "top":
		current.top = matchPiece
	case "right":
		current.right = matchPiece
	case "bottom":
		current.bottom = matchPiece
	case "left":
		current.left = matchPiece
	}
}

func matchSide(tiles []*tile, tile *tile, compare compareFunc) int {
	for i, candidate := range tiles {
		if candidate.id == tile.id {
			continue
		}
		for rotation := 0; rotation < 4; rotation++ {
			candidate.rotate()
			if compare(tile, candidate) {
				return i
			}
			candidate.hFlip()
			if compare(tile, candidate) {
				return i
			}
			candidate.hFlip()
			candidate.vFlip()
			if compare(tile, candidate) {
				return i
			}
			candidate.vFlip()
		}
	}
	return -1
}

func matchPlaced(placed []*tile, tile *tile, compare compareFunc) *tile {
	for _, candidate := range placed {
		if candidate.id == tile.id {
			continue
		}
		if compare(tile, candidate) {
			return candidate
		}
	}
	return nil
}

type compareFunc func(tile1 *tile, tile2 *tile) bool

func compareTopToBottom(tile1 *tile, tile2 *tile) bool {
	if tile1.pixels[0] == tile2.pixels[len(tile2.pixels)-1] {
		return true
	}
	return false
}

func compareBottomToTop(tile1 *tile, tile2 *tile) bool {
	if tile1.pixels[len(tile1.pixels)-1] == tile2.pixels[0] {
		return true
	}
	return false
}

func compareLeftToRight(tile1 *tile, tile2 *tile) bool {
	for i := 0; i < len(tile1.pixels); i++ {
		left := tile1.pixels[i]
		right := tile2.pixels[i]
		if (left&0b_10_0000_0000 == 0 && right&0b_01 != 0) ||
			(left&0b_10_0000_0000 != 0 && right&0b_01 == 0) {
			return false
		}
	}
	return true
}

func compareRightToLeft(tile1 *tile, tile2 *tile) bool {
	for i := 0; i < len(tile1.pixels); i++ {
		right := tile1.pixels[i]
		left := tile2.pixels[i]
		if (left&0b_10_0000_0000 == 0 && right&0b_01 != 0) ||
			(left&0b_10_0000_0000 != 0 && right&0b_01 == 0) {
			return false
		}
	}
	return true
}
