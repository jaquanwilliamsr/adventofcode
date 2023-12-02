package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type cubes struct {
	red   int
	green int
	blue  int
}

// min amount of cubes needed to play a game (max per round)
func getMax(max *cubes, roundCube *cubes) {
	if roundCube.red > max.red {
		max.red = roundCube.red
	}
	if roundCube.green > max.green {
		max.green = roundCube.green
	}
	if roundCube.blue > max.blue {
		max.blue = roundCube.blue
	}
}

func cubesValues(round string) *cubes {
	var roundCubes cubes

	cubesRegex := regexp.MustCompile(`(\d+)\s(blue|red|green)`)
	// Strip the number from the color later
	stripNum := regexp.MustCompile(`\d+`)

	// Complains if I don't initialize err here
	err := error(nil)

	matches := cubesRegex.FindAllString(round, -1)
	for _, match := range matches {
		color := strings.Trim(stripNum.ReplaceAllString(match[:], ""), " ")
		amount := strings.Trim(match[0:2], " ")
		switch color {
		case "red":
			roundCubes.red, err = strconv.Atoi(amount)
			if err != nil {
				panic(err)
			}
		case "green":
			roundCubes.green, err = strconv.Atoi(amount)
			if err != nil {
				panic(err)
			}
		case "blue":
			roundCubes.blue, err = strconv.Atoi(amount)
			if err != nil {
				panic(err)
			}
		}
	}
	return &roundCubes
}

func main() {
	var max cubes
	var total int

	games := ReadInput("input.txt")
	for _, rounds := range games {
		splitRounds := strings.Split(rounds, ";")
		for _, round := range splitRounds {
			roundValues := cubesValues(round)
			getMax(&max, roundValues)
		}
		total += max.red * max.green * max.blue
		// Reset max after each game
		max = cubes{}
	}
	fmt.Println(total)
}
