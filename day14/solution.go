package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run in circles.
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 5875750429995
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 5272149590143
}

func part1(inputPath string) int {
	maskSet := 0
	maskClear := 0
	memory := map[int]int{}
	scanner := helpers.NewScanner(inputPath)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		cmd := line[0:3]
		if cmd == "mem" {
			i := strings.Index(line, "]")
			addr, err := strconv.Atoi(line[4:i])
			if err != nil {
				panic(err)
			}
			i = strings.Index(line, "=")
			value, err := strconv.Atoi(line[i+2:])
			maskedValue := (value | maskSet) & maskClear
			// fmt.Printf("mem[%d]=%d (%d)\n", addr, value, maskedValue)
			memory[addr] = maskedValue
		} else if cmd == "mas" {
			i := strings.Index(line, "=")
			s := line[i+2:]
			// fmt.Printf("mask = %s\n", s)
			v, err := strconv.ParseInt(strings.ReplaceAll(s, "X", "0"), 2, 64)
			if err != nil {
				panic(err)
			}
			maskSet = int(v)
			v, err = strconv.ParseInt(strings.ReplaceAll(s, "X", "1"), 2, 64)
			if err != nil {
				panic(err)
			}
			maskClear = int(v)
			// fmt.Printf("  set: %36b\n", maskSet)
			// fmt.Printf("  clr: %36b\n", maskClear)
		}
	}
	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

// not proud of this unmaintainable mess but it got the right answer
func part2(inputPath string) int {
	ormask := 0
	floatmask := 0
	floatbits := 0
	bitmap := []int{}
	memory := map[int]int{}
	scanner := helpers.NewScanner(inputPath)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		cmd := line[0:3]
		if cmd == "mem" {
			i := strings.Index(line, "]")
			addr, err := strconv.Atoi(line[4:i])
			if err != nil {
				panic(err)
			}
			i = strings.Index(line, "=")
			value, err := strconv.Atoi(line[i+2:])
			// fmt.Printf("mem[%d]=%d\n", addr, value)
			addr = addr | ormask
			addr = addr & (^floatmask)
			for j := 0; j < 1<<floatbits; j++ {
				actual := addr
				for bitnum, v := range bitmap {
					if j&(1<<bitnum) != 0 {
						actual |= v
					}
				}
				// fmt.Printf("  mem[%d]=%d\n", actual, value)
				memory[actual] = value
			}
		} else if cmd == "mas" {
			i := strings.Index(line, "=")
			s := line[i+2:]
			// fmt.Printf("mask = %s\n", s)
			v, err := strconv.ParseInt(strings.ReplaceAll(s, "X", "0"), 2, 64)
			if err != nil {
				panic(err)
			}
			ormask = int(v)
			s = strings.ReplaceAll(s, "1", "0")
			floatbits = strings.Count(s, "X")
			v, err = strconv.ParseInt(strings.ReplaceAll(s, "X", "1"), 2, 64)
			if err != nil {
				panic(err)
			}
			floatmask = int(v)
			bitmap = make([]int, floatbits)
			mask := 1
			for i = 0; i < floatbits; i++ {
				for floatmask&mask == 0 {
					mask <<= 1
				}
				bitmap[i] = mask
				// fmt.Printf("  mask[%d] = %36b\n", i, mask)
				mask <<= 1
			}
			// fmt.Printf("  set: %36b\n", ormask)
			// fmt.Printf("  flt: %36b\n", floatmask)
			// fmt.Printf("  bits: %d\n", floatbits)
		}
	}
	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}
