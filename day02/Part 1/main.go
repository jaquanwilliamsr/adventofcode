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

func isValidGame(max *cubes, round *cubes) bool {
	if max.red < round.red || max.green < round.green || max.blue < round.blue {
		return false
	}
	return true
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
	var valid bool
	var total int

	max := cubes{red: 12, green: 13, blue: 14}
	games := ReadInput("input.txt")

	for i, rounds := range games {
		splitRounds := strings.Split(rounds, ";")
		for _, round := range splitRounds {
			roundValues := cubesValues(round)
			if !isValidGame(&max, roundValues) {
				valid = false
				break
			}
			valid = true
		}
		if valid {
			total += i + 1
		}
	}
	fmt.Println(total)
}
