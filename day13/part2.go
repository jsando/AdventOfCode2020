package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	buses, err := readInput(os.Args[1])
	if err != nil {
		panic(err)
	}
	time := findSequencedDeparture(buses)
	fmt.Printf("part 2: %d\n", time) // 725850285300475
}

func readInput(filename string) (buses []int, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	input := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	buses = make([]int, 0)
	busidStrings := strings.Split(input[1], ",")
	var busid int
	for _, s := range busidStrings {
		if s == "x" {
			busid = 0
		} else {
			busid, err = strconv.Atoi(s)
			if err != nil {
				return
			}
		}
		buses = append(buses, busid)
	}
	return
}

// There's no getting around having to test lots of possible solutions BUT we can
// significantly reduce the number that have to be tested by (a) starting with
// the largest id in the list and checking multiples of it for the others, and
// then (b) for each additional solution, multiply the increment times that id so
// we continue to only check multples of the existing known solution(s) for the
// remaining buses.
func findSequencedDeparture(buses []int) int {
	maxID, maxOffset := findMaxInt(buses)
	increment := maxID
	buses[maxOffset] = 0
	for t := maxID - maxOffset; t < math.MaxInt64; t += increment {
		for i, id := range buses {
			if id != 0 && (t+i)%id == 0 {
				increment *= id
				buses[i] = 0
			}
		}
		if allFound(buses) {
			return t
		}
	}
	return -1
}

func allFound(buses []int) bool {
	for _, id := range buses {
		if id != 0 {
			return false
		}
	}
	return true
}

func findMaxInt(va []int) (int, int) {
	maxv := 0
	maxp := 0
	for i, v := range va {
		if maxv < v {
			maxv = v
			maxp = i
		}
	}
	return maxv, maxp
}
