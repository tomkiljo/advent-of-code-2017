package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Last(n int) int {
	return (n*(n+1)/2)*8 + 1
}

func main() {
	fmt.Println("Advent of Code 2017 - Day 3 Part 1")
	fmt.Print("Enter puzzle input: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(text))

	n := 1
	for i := 1; val >= Last(i); i++ {
		n++
	}

	seq := Last(n) - val
	dis := seq - (seq/(n*2))*(n*2)
	fmt.Println("Result: ", dis)
}
