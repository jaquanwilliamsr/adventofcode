package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func findNums(chars string) int {
	regex := regexp.MustCompile("[0-9]")
	matches := regex.FindAllString(chars, -1)
	numString := matches[0] + matches[len(matches)-1]

	num, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	sum := 0
	input := ReadInput("input.txt")

	for _, chars := range input {
		sum += findNums(chars)
	}
	fmt.Println(sum)
}
