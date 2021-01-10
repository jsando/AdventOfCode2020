package day25

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Run day 25.
func Run(inputPath string) {
	bytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	s := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	pk1, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	pk2, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}
	loop1 := getLoopSize(7, pk1)
	fmt.Printf("Part 1: %d\n", createPublicKey(pk2, loop1)) // 17032383
}

// my first attempt called createPublicKey below repeatedly until it found the right
// key ... which required loopSize! re-calculations of the same value.  Duh.
func getLoopSize(subjectNumber, publicKey int) int {
	value := 1
	loop := 0
	for value != publicKey {
		value *= subjectNumber
		value %= 20201227
		loop++
	}
	return loop
}

func createPublicKey(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = (value * subjectNumber) % 20201227
	}
	return value
}
