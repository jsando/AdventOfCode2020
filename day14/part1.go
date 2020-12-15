package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	maskSet := 0
	maskClear := 0
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
			maskedValue := (value | maskSet) & maskClear
			fmt.Printf("mem[%d]=%d (%d)\n", addr, value, maskedValue)
			memory[addr] = maskedValue
		} else if cmd == "mas" {
			i := strings.Index(line, "=")
			s := line[i+2:]
			fmt.Printf("mask = %s\n", s)
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
			fmt.Printf("  set: %36b\n", maskSet)
			fmt.Printf("  clr: %36b\n", maskClear)
		}
	}
	sum := 0
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("part 1: %d\n", sum)
}
