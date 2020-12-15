package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// not proud of this unmaintainable mess but it got the right answer
func main() {
	ormask := 0
	floatmask := 0
	floatbits := 0
	bitmap := []int{}
	memory := map[int]int{}
	scanner := bufio.NewScanner(os.Stdin)
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
			fmt.Printf("mem[%d]=%d\n", addr, value)
			addr = addr | ormask
			addr = addr & (^floatmask)
			for j := 0; j < 1<<floatbits; j++ {
				actual := addr
				for bitnum, v := range bitmap {
					if j&(1<<bitnum) != 0 {
						actual |= v
					}
				}
				fmt.Printf("  mem[%d]=%d\n", actual, value)
				memory[actual] = value
			}
		} else if cmd == "mas" {
			i := strings.Index(line, "=")
			s := line[i+2:]
			fmt.Printf("mask = %s\n", s)
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
				fmt.Printf("  mask[%d] = %36b\n", i, mask)
				mask <<= 1
			}
			fmt.Printf("  set: %36b\n", ormask)
			fmt.Printf("  flt: %36b\n", floatmask)
			fmt.Printf("  bits: %d\n", floatbits)
		}
	}
	sum := 0
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("part 2: %d\n", sum)
}
