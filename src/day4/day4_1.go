package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 4 Part 1")
	fmt.Println("Enter puzzle input: ")

	count := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}

		valid := true
		encountered := map[string]bool{}
		for _, str := range strings.Fields(line) {
			if _, has := encountered[str]; has {
				valid = false
				break
			}
			encountered[str] = true
		}

		if valid {
			count++
		}
	}

	fmt.Println("Result: ", count)
}
