package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"reflect"
)

func Largest(arr []int) (int, int) {
	l, x := 0, 0
	for i, v := range arr {
		if v > l {
			x = i
			l = v
		}
	}
	return x, l
}

func main() {
	fmt.Println("Advent of Code 2017 - Day 6 Part 1")
	fmt.Print("Enter puzzle input: ")

	var values []int
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, str := range strings.Fields(strings.TrimSpace(line)) {
		num, _ := strconv.Atoi(str)
		values = append(values, num)
	}

	var res int
	seen := map[int][]int{}
Loop:
	for count := 0; ; count++ {
		i, v := Largest(values)
		values[i] = 0

		for ; v > 0; v-- {
			if i++; i >= len(values) {
				i = 0
			}
			values[i] = values[i] + 1
		}

		for key, val := range seen {
			if reflect.DeepEqual(values, val) {
				res = count - key
				break Loop
			}
		}
		seen[count] = append([]int(nil), values...)
	}

	fmt.Println("Result: ", res)
}
