package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
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
	fmt.Printf("Pulic Key 1: %d, Public Key 2: %d\n", pk1, pk2)
	loop1 := getLoopSize(7, pk1)
	fmt.Printf("Loop size 1: %d\n", loop1)
	loop2 := getLoopSize(7, pk2)
	fmt.Printf("Loop size 2: %d\n", loop2)
	fmt.Printf("Key1: %d, Key2: %d\n", createPublicKey(pk1, loop2), createPublicKey(pk2, loop1))
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
