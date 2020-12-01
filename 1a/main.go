package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go run main.go < input
func main() {
	entries := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		entries = append(entries, entry)
	}
	for i := 0; i < len(entries); i++ {
		a := entries[i]
		for j := i + 1; j < len(entries); j++ {
			b := entries[j]
			if a+b == 2020 {
				fmt.Printf("%d + %d = 2020, %d * %d = %d\n", a, b, a, b, a*b)
			}
		}
	}
}
