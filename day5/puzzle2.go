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
	occupied := make([]bool, 1024)
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
		occupied[seatID] = true
		fmt.Printf("%s -> %s (%d)\n", line, binary, seatID)
	}
	mySeatID := 0
	for i := 0; i < len(occupied)-3; i++ {
		if occupied[i] && !occupied[i+1] && occupied[i+2] {
			mySeatID = i + 1
			break
		}
	}
	fmt.Printf("My Seat ID: %d\n", mySeatID)
}
