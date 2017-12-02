package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2017 - Day 1 Part 1")
	fmt.Print("Enter puzzle input: ")
	text, _ := reader.ReadString('\n')

	runes := []rune(strings.TrimSpace(text))
	count := len(runes)

	pre := -1
	sum := 0

	for i := 0; i <= count; i++ {
		ind := i
		if ind == count {
			ind = 0
		}

		cur := int(runes[ind] - '0')
		if pre == cur {
			sum += cur
		}
		pre = cur
	}

	fmt.Println("Result: ", sum)
}
