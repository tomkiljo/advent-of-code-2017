package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2017 - Day 1 Part 2")
	fmt.Print("Enter puzzle input: ")
	text, _ := reader.ReadString('\n')

	runes := []rune(strings.TrimSpace(text))
	count := len(runes)

	sum := 0

	for i := 0; i <= count; i++ {
		cind := i
		if cind == count {
			cind = 0
		}

		nind := cind + (count / 2)
		if nind > (count - 1) {
			nind = nind - count
		}

		cur := int(runes[cind] - '0')
		nxt := int(runes[nind] - '0')

		if cur == nxt {
			sum += cur
		}
	}

	fmt.Println("Result: ", sum)
}
