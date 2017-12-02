package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Advent of Code 2017 - Day 2 Part 1")
	fmt.Println("Enter puzzle input: ")

	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}

		var values []int
		for _, str := range strings.Fields(line) {
			num, _ := strconv.Atoi(str)
			for _, val := range values {
				if MaxInt(num, val) % MinInt(num, val) == 0 {
					sum += MaxInt(num, val) / MinInt(num, val)
					break
				}
			}
			values = append(values, num)
		}
	}

	fmt.Println("Result: ", sum)
}
