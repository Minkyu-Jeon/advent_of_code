package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func readInput(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("There is no such file `%s`", path)
	}

	lines := strings.Split(string(data), "\n")

	return lines[:len(lines)-1]
}

func main() {
	var part int
	var input = readInput("./input.txt")

	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func getCalibaratedValue(digits [][]int) int {
	return digits[0][0]*10 + digits[len(digits)-1][0]
}

func transformPart1(line string) [][]int {
	answer := [][]int{}

	for index, c := range line {
		if unicode.IsDigit(c) {
			intVal, _ := strconv.Atoi(string(c))

			answer = append(answer, []int{intVal, index})
		}
	}

	return answer
}

func part1(input []string) int {
	answer := 0
	for _, item := range input {
		digits := transformPart1(item)
		answer += getCalibaratedValue(digits)
	}
	return answer
}

func transformPart2(line string, numbers []string) [][]int {
	answer := [][]int{}

	translateToNumberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, number := range numbers {
		re := regexp.MustCompile(number)
		indecies := re.FindAllIndex([]byte(line), -1)

		for _, item := range indecies {
			intVal, _ := strconv.Atoi(translateToNumberMap[number])
			answer = append(answer, []int{intVal, item[0]})
		}
	}

	answer = append(answer, transformPart1(line)...)

	sort.SliceStable(answer, func(a int, b int) bool {
		return answer[a][1] < answer[b][1]
	})

	return answer
}

func part2(input []string) int {
	answer := 0
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, item := range input {
		digits := transformPart2(item, numbers)
		answer += getCalibaratedValue(digits)
	}

	return answer
}
