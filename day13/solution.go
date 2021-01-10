package day13

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Run dmc (deer motor corp).
func Run(inputPath string) {
	fmt.Printf("Part 1: %d\n", part1(inputPath)) // 4207
	fmt.Printf("Part 2: %d\n", part2(inputPath)) // 725850285300475
}

func part1(inputPath string) int {
	startTime, buses, err := readInput(inputPath)
	if err != nil {
		panic(err)
	}
	nextBus, nextTime := findNextBus(startTime, buses)
	waitTime := nextTime - startTime
	// fmt.Printf("Next bus from start time %d is %d, wait time %d.  Part 1: %d\n", startTime, nextBus, waitTime, waitTime*nextBus)
	return waitTime * nextBus
}

func readInput(filename string) (startTime int, buses []int, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	input := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	startTime, err = strconv.Atoi(input[0])
	if err != nil {
		return
	}
	buses = make([]int, 0)
	busidStrings := strings.Split(input[1], ",")
	for _, s := range busidStrings {
		if s == "x" {
			continue
		}
		var busid int
		busid, err = strconv.Atoi(s)
		if err != nil {
			return
		}
		buses = append(buses, busid)
	}
	return
}

func findNextBus(startTime int, buses []int) (bus int, time int) {
	for departTime := startTime; departTime < math.MaxInt64; departTime++ {
		for _, id := range buses {
			if departTime%id == 0 {
				return id, departTime
			}
		}
	}
	return -1, -1
}

func part2(inputPath string) int {
	buses, err := readInput2(inputPath)
	if err != nil {
		panic(err)
	}
	time := findSequencedDeparture(buses)
	return time
}

func readInput2(filename string) (buses []int, err error) {
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
