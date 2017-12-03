package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SumNeighbours(arr [][]int, x, y int) int {
	dim, sum := len(arr), 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < dim && j >= 0 && j < dim {
				sum += arr[i][j]
			}
		}
	}
	arr[x][y] = sum
	return sum
}

func ExpandSlice(arr [][]int) [][]int {
	dim := len(arr) + 2
	var res [][]int
	res = make([][]int, dim, dim)
	for i := 0; i < dim; i++ {
		res[i] = make([]int, dim)
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			res[i+1][j+1] = arr[i][j]
		}
	}
	return res
}

func main() {
	fmt.Println("Advent of Code 2017 - Day 3 Part 1")
	fmt.Print("Enter puzzle input: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(text))

	var arr [][]int
	arr = make([][]int, 1, 1)
	arr[0] = make([]int, 1)
	arr[0][0] = 1
	num := 0

Loop:
	for i := 1; ; i++ {
		arr = ExpandSlice(arr)
		x, y := 2*i, 2*i-1

		// right
		if num = SumNeighbours(arr, x, y); num > val {
			break Loop
		}
		// up
		for j := 0; j < i*2-1; j++ {
			y--
			if num = SumNeighbours(arr, x, y); num > val {
				break Loop
			}
		}
		// left
		for j := 0; j < i*2; j++ {
			x--
			if num = SumNeighbours(arr, x, y); num > val {
				break Loop
			}
		}
		// down
		for j := 0; j < i*2; j++ {
			y++
			if num = SumNeighbours(arr, x, y); num > val {
				break Loop
			}
		}
		// right
		for j := 0; j < i*2; j++ {
			x++
			if num = SumNeighbours(arr, x, y); num > val {
				break Loop
			}
		}
	}

	fmt.Println("Result: ", num)
}
