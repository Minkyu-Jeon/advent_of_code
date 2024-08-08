package main

import (
	"flag"
	"fmt"
	"log"
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

	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input, cubeSet{red: 12, green: 13, blue: 14})
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

type cubeSet struct {
	red   int
	green int
	blue  int
}

type game struct {
	gameId int
	cubes  []cubeSet
}

// 가장 큰 수만을 담도록
func parseInput(line string, re *regexp.Regexp) game {
	match := re.FindStringSubmatch(line)
	gameId, _ := strconv.Atoi(match[1])
	rounds := match[2]

	cubes := []cubeSet{}

	for _, round := range strings.Split(rounds, "; ") {
		cube := cubeSet{red: 0, green: 0, blue: 0}

		for _, eachCubes := range strings.Split(round, ", ") {
			countAndColor := strings.Split(eachCubes, " ")
			cubeCount, _ := strconv.Atoi(countAndColor[0])
			cubeColor := countAndColor[1]

			switch cubeColor {
			case "red":
				cube.red = cubeCount
			case "green":
				cube.green = cubeCount
			case "blue":
				cube.blue = cubeCount
			}
		}
		cubes = append(cubes, cube)
	}

	return game{gameId: gameId, cubes: cubes}
}

func getMaxAmountofCubes(cubes []cubeSet) cubeSet {
	maxRed, maxGreen, maxBlue := 0, 0, 0

	for _, cube := range cubes {
		if cube.red > maxRed {
			maxRed = cube.red
		}

		if cube.green > maxGreen {
			maxGreen = cube.green
		}

		if cube.blue > maxBlue {
			maxBlue = cube.blue
		}
	}

	return cubeSet{red: maxRed, green: maxGreen, blue: maxBlue}
}

func compareGame(c1 cubeSet, c2 cubeSet) bool {
	return c1.red <= c2.red && c1.green <= c2.green && c1.blue <= c2.blue
}

func part1(input []string, totalCubes cubeSet) int {
	answer := 0
	re := regexp.MustCompile(`Game (?P<gameId>\d+): (.+)`)
	for _, item := range input {
		gm := parseInput(item, re)
		if compareGame(getMaxAmountofCubes(gm.cubes), totalCubes) {
			answer += gm.gameId
		}
	}
	return answer
}

func part2(input []string) int {
	answer := 0
	re := regexp.MustCompile(`Game (?P<gameId>\d+): (.+)`)
	for _, item := range input {
		gm := parseInput(item, re)
		cube := getMaxAmountofCubes(gm.cubes)
		answer += (cube.red * cube.green * cube.blue)
	}
	return answer
}
