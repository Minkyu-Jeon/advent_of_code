package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	re := regexp.MustCompile(`Card\s+(?P<cardNumber>\d+):(?P<winingNumbers>(\s[\d\s]{2})+)\s\|(?P<myNumbers>(\s[\d\s]{2})+)`)

	if part == 1 {
		ans := part1(input, re)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input, re)
		fmt.Println("Output:", ans)
	}
}

type game struct {
	cardNumber    int
	winingNumbers map[int]bool
	myNumbers     map[int]bool
}

func (g *game) getMatchedNumbers() int {
	intersection := []int{}

	for myNumber := range g.myNumbers {
		_, exists := g.winingNumbers[myNumber]

		if exists {
			intersection = append(intersection, myNumber)
		}
	}

	return len(intersection)
}

func numbersToMap(numbers string) map[int]bool {
	m := make(map[int]bool)

	for _, strNumber := range strings.Split(strings.ReplaceAll(numbers, "  ", " "), " ") {
		i, err := strconv.Atoi(strNumber)
		if err != nil {
			continue
		}
		m[i] = true
	}

	return m
}

func part1(input []string, re *regexp.Regexp) int {
	answer := 0

	for _, line := range input {
		match := re.FindStringSubmatch(line)
		params := make(map[string]string)
		for i, item := range re.SubexpNames() {
			if item == "" {
				continue
			}
			params[item] = match[i]
		}
		cardNumber, _ := strconv.Atoi(params["cardNumber"])
		winingNumbers := numbersToMap(params["winingNumbers"])
		myNumbers := numbersToMap(params["myNumbers"])

		g := game{
			cardNumber:    cardNumber,
			winingNumbers: winingNumbers,
			myNumbers:     myNumbers,
		}
		sum := int(math.Pow(float64(2), float64(g.getMatchedNumbers()-1)))
		answer += sum
	}

	return answer
}

func part2(input []string, re *regexp.Regexp) int {
	answer := 0
	accumulator := make(map[int]int)

	for _, line := range input {
		match := re.FindStringSubmatch(line)
		params := make(map[string]string)
		for i, item := range re.SubexpNames() {
			if item == "" {
				continue
			}
			params[item] = match[i]
		}
		cardNumber, _ := strconv.Atoi(params["cardNumber"])
		winingNumbers := numbersToMap(params["winingNumbers"])
		myNumbers := numbersToMap(params["myNumbers"])

		g := game{
			cardNumber:    cardNumber,
			winingNumbers: winingNumbers,
			myNumbers:     myNumbers,
		}

		_, exists := accumulator[g.cardNumber]
		matchedNumber := g.getMatchedNumbers()
		if !exists {
			accumulator[g.cardNumber] = 1
		}

		for i := 0; i < accumulator[g.cardNumber]; i++ {
			for j := 1; j <= matchedNumber; j++ {
				v, exists := accumulator[g.cardNumber+j]
				if exists {
					v += 1
				} else {
					v = 2
				}
				accumulator[g.cardNumber+j] = v
			}
		}
	}

	for cardNumber := 1; cardNumber <= len(input); cardNumber++ {
		answer += accumulator[cardNumber]
	}

	return answer
}
