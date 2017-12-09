package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 5 Part 2")
	fmt.Println("Enter puzzle input: ")

	var values []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		num, _ := strconv.Atoi(line)
		values = append(values, num)
	}

	i := 0
	var c int
	for c = 0; i >= 0 && i < len(values); c++ {
		n := values[i]
		if n < 3 {
			values[i] = n + 1
		} else {
			values[i] = n - 1
		}
		i += n
	}

	fmt.Println("Result: ", c)
}
