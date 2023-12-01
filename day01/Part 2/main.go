package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	forwardMap = map[string]string{
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

	reversedMap = map[string]string{
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
)

func reverseString(chars string) string {
	for i, char := range chars {
		if i >= len(chars)/2 {
			break
		}
		chars = chars[:i] + string(chars[len(chars)-1-i]) + chars[i+1:]
		chars = chars[:len(chars)-1-i] + string(char) + chars[len(chars)-i:]
	}
	return chars
}

func createRegexMap(maps map[string]string) string {
	var pattern []string
	for key, value := range maps {
		pattern = append(pattern, key)
		pattern = append(pattern, value)
	}
	return strings.Join(pattern, "|")
}

func findFirstNum(chars string) string {
	pattern := createRegexMap(forwardMap)
	regex := regexp.MustCompile(pattern)
	match := regex.FindString(chars)
	if forwardMap[match] != "" {
		return forwardMap[match]
	}
	return match
}

func findLastNum(chars string) string {
	pattern := createRegexMap(reversedMap)
	regex := regexp.MustCompile(pattern)
	reversedInput := reverseString(chars)
	match := regex.FindString(reversedInput)
	if reversedMap[match] != "" {
		return reversedMap[match]
	}
	return match
}

func combineNums(num1 string, num2 string) int {
	numString := num1 + num2
	num, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	sum := 0
	input := ReadInput("input.txt")
	for _, char := range input {
		firstNum := findFirstNum(char)
		lastNum := findLastNum(char)
		num := combineNums(firstNum, lastNum)
		sum += num
	}
	fmt.Println(sum)
}
