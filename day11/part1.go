package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seats := make([][]byte, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		seats = append(seats, []byte(line))
	}

	// printSeats(seats)
	fmt.Printf("Available seats before seating: %d\n", countSeats(seats, byte('L')))
	fmt.Printf("Part 1 - occupied: %d\n", part1(seats)) // 2299
}

func printSeats(seats [][]byte) {
	for i := 0; i < len(seats); i++ {
		fmt.Printf("%s\n", string(seats[i]))
	}
	fmt.Println()
}

func part1(seats [][]byte) int {
	for {
		var changes int
		seats, changes = iterate(seats)
		if changes == 0 {
			break
		}
	}
	return countSeats(seats, '#')
}

func iterate(seats [][]byte) ([][]byte, int) {
	dup := make([][]byte, len(seats))
	for i := 0; i < len(seats); i++ {
		dup[i] = make([]byte, len(seats[i]))
		copy(dup[i], seats[i])
	}
	changes := 0
	for row := 0; row < len(seats); row++ {
		for col := 0; col < len(seats[row]); col++ {
			occupied := occupiedAround(seats, row, col)
			ch := seats[row][col]
			if ch == byte('#') {
				if occupied >= 4 {
					dup[row][col] = byte('L')
					changes++
				}
			} else if ch == byte('L') {
				if occupied == 0 {
					changes++
					dup[row][col] = byte('#')
				}
			}
		}
	}
	return dup, changes
}

func occupiedAround(seats [][]byte, row int, column int) int {
	count := 0
	if row > 0 {
		count += occupied(seats, row-1, column-1)
		count += occupied(seats, row-1, column)
		count += occupied(seats, row-1, column+1)
	}
	count += occupied(seats, row, column-1)
	count += occupied(seats, row, column+1)
	if row < len(seats) {
		count += occupied(seats, row+1, column-1)
		count += occupied(seats, row+1, column)
		count += occupied(seats, row+1, column+1)
	}
	return count
}

func occupied(seats [][]byte, row int, column int) int {
	if row < 0 || row >= len(seats) {
		return 0
	}
	if column < 0 || column >= len(seats[row]) {
		return 0
	}
	ch := seats[row][column]
	if ch == byte('#') {
		return 1
	}
	return 0
}

func countSeats(seats [][]byte, seatType byte) int {
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == seatType {
				count++
			}
		}
	}
	return count
}
