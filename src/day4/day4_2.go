package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"reflect"
)

func GetCharCount(s string) (c map[rune]int) {
	c = make(map[rune]int)
	for _, runeValue := range s {
		if _, ok := c[runeValue]; ok {
			c[runeValue] += 1
		} else {
			c[runeValue] = 1
		}
	}
	return
}

func main() {
	fmt.Println("Advent of Code 2017 - Day 4 Part 2")
	fmt.Println("Enter puzzle input: ")

	count := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}

		valid := true
		encountered := map[string]map[rune]int{}

	Loop:
		for _, str := range strings.Fields(line) {
			chars := GetCharCount(str)
			for key, val := range encountered {
				if key == str || reflect.DeepEqual(val, chars) {
					encountered[str] = chars
					valid = false
					break Loop
				}
			}
			encountered[str] = chars
		}

		if valid {
			count++
		}
	}

	fmt.Println("Result: ", count)
}
