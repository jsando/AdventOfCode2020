package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go run main.go < input.txt
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	treeCount := 0
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
	fmt.Printf("Tree count: %d\n", treeCount)
}
