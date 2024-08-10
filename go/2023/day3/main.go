package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var dx = []int{-1, 0, 1, -1, 1, -1, 0, 1}
var dy = []int{-1, -1, -1, 0, 0, 1, 1, 1}

func readInput(path string) [][]byte {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("There is no such file `%s`", path)
	}

	lines := strings.Split(string(data), "\n")

	result := make([][]byte, len(lines))
	for i, line := range lines[:len(lines)-1] {
		result[i] = []byte(line)
	}

	return result
}

func main() {
	var part int
	var input = readInput("./input.txt")
	visit := make([][]bool, len(input))
	for i, line := range input {
		visit[i] = make([]bool, len(line))
		for j := range line {
			visit[i][j] = false
		}
	}

	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input, visit)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input, visit)
		fmt.Println("Output:", ans)
	}
}

func part1(input [][]byte, visit [][]bool) int {
	answer := 0

	maxX := len(input) - 1
	maxY := len(input[0]) - 1

	for i, line := range input {
		for j, char := range line {
			if !(unicode.IsDigit(rune(char)) || char == '.') {
				for d := 0; d < 8; d++ {
					nextX := j + dx[d]
					nextY := i + dy[d]
					if !(nextY >= 0 && nextY <= maxY && nextX >= 0 && nextX <= maxX) {
						continue
					}

					if unicode.IsDigit(rune(input[nextY][nextX])) {
						digits := searchDigits(nextX, nextY, input, visit)
						answer += calc(digits)
					}
				}
			}
		}
	}

	return answer
}

func part2(input [][]byte, visit [][]bool) int {
	answer := 0

	maxX := len(input) - 1
	maxY := len(input[0]) - 1

	for i, line := range input {
		for j, char := range line {
			if char == '*' {
				part := []int{}
				for d := 0; d < 8; d++ {
					nextX := j + dx[d]
					nextY := i + dy[d]
					if !(nextY >= 0 && nextY <= maxY && nextX >= 0 && nextX <= maxX) {
						continue
					}

					if unicode.IsDigit(rune(input[nextY][nextX])) {
						digits := calc(searchDigits(nextX, nextY, input, visit))

						if digits != 0 {
							part = append(part, digits)
						}
					}
				}

				if len(part) == 2 {
					acc := 1
					for _, p := range part {
						acc *= p
					}
					answer += acc
				}
			}
		}
	}

	return answer
}

func searchDigits(x int, y int, input [][]byte, visit [][]bool) []int {
	if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) || visit[y][x] {
		return []int{}
	}

	visit[y][x] = true

	char := input[y][x]

	if !unicode.IsDigit(rune(char)) {
		return []int{}
	}

	s := string(char)
	i, err := strconv.Atoi(s)

	if err != nil {
		return []int{}
	}

	return append(searchDigits(x-1, y, input, visit), append([]int{i}, searchDigits(x+1, y, input, visit)...)...)
}

func calc(arr []int) int {
	answer := 0

	for i := len(arr) - 1; i >= 0; i-- {
		answer += int(math.Pow(float64(10), float64(i))) * arr[len(arr)-i-1]
	}

	return answer
}
