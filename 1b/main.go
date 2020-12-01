package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go run main.go < ../1a/input
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
			for k := j + 1; k < len(entries); k++ {
				c := entries[k]
				if a+b+c == 2020 {
					fmt.Printf("%d + %d + %d = 2020, * = %d\n", a, b, c, a*b*c)
				}
			}
		}
	}
}
