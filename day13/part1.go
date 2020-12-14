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
	startTime, buses, err := readInput(os.Args[1])
	if err != nil {
		panic(err)
	}
	nextBus, nextTime := findNextBus(startTime, buses)
	waitTime := nextTime - startTime
	fmt.Printf("Next bus from start time %d is %d, wait time %d.  Part 1: %d\n", startTime, nextBus, waitTime, waitTime*nextBus)
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
