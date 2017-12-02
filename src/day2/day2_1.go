package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 2 Part 1")
	fmt.Println("Enter puzzle input: ")

	checksum := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}

		max := math.MinInt16
		min := math.MaxInt16

		for _, str := range strings.Fields(line) {
			num, _ := strconv.Atoi(str)
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}

		checksum += max - min
	}

	fmt.Println("Result: ", checksum)
}
