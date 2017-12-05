package day1_advent2017

import (
	"strconv"
	"strings"
)

func Day1(num string) int {
	interval := 1
	return Day1_Totaler(num, interval)
}

func Day1_Part2(num string) int {
	interval := len(num) / 2
	return Day1_Totaler(num, interval)
}

func Day1_Totaler(num string, interval int) int {
	split := strings.Split(num, "")
	total := 0
	for i := 0; i < len(split); i++ {
		next := (i + interval) % len(split)
		if split[i] == split[next] {
			val, _ := strconv.Atoi(split[i])
			total += val
		}
	}

	return total
}
