package main

import (
	"flag"

	"github.com/jsando/aoc2020/day01"
	"github.com/jsando/aoc2020/day02"
	"github.com/jsando/aoc2020/day03"
	"github.com/jsando/aoc2020/day04"
	"github.com/jsando/aoc2020/day05"
	"github.com/jsando/aoc2020/day06"
	"github.com/jsando/aoc2020/day07"
	"github.com/jsando/aoc2020/day08"
	"github.com/jsando/aoc2020/day09"
	"github.com/jsando/aoc2020/day10"
	"github.com/jsando/aoc2020/day11"
	"github.com/jsando/aoc2020/day12"
	"github.com/jsando/aoc2020/day13"
	"github.com/jsando/aoc2020/day14"
	"github.com/jsando/aoc2020/day15"
	"github.com/jsando/aoc2020/day16"
	"github.com/jsando/aoc2020/day17"
	"github.com/jsando/aoc2020/day18"
	"github.com/jsando/aoc2020/day19"
	"github.com/jsando/aoc2020/day20"
	"github.com/jsando/aoc2020/day21"
	"github.com/jsando/aoc2020/day22"
	"github.com/jsando/aoc2020/day23"
	"github.com/jsando/aoc2020/day24"
	"github.com/jsando/aoc2020/day25"
)

var day = flag.Int("d", 0, "day number (1...25)")
var inputPath = flag.String("i", "", "optional input filename")

type runner func(inputPath string)

var runners []runner = []runner{
	day01.Run, day02.Run, day03.Run, day04.Run, day05.Run,
	day06.Run, day07.Run, day08.Run, day09.Run, day10.Run,
	day11.Run, day12.Run, day13.Run, day14.Run, day15.Run,
	day16.Run, day17.Run, day18.Run, day19.Run, day20.Run,
	day21.Run, day22.Run, day23.Run, day24.Run, day25.Run,
}

func main() {
	flag.Parse()
	runners[*day-1](*inputPath)
}
