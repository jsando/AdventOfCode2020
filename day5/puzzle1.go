package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	max := int64(0)
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
		if seatID > max {
			max = seatID
		}
		fmt.Printf("%s -> %s (%d)\n", line, binary, seatID)
	}
	fmt.Printf("Max: %d\n", max)
}
