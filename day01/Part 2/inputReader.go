package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(fileName string) []string {
	fileContents := []string{}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}
	filePath := dir + "/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileContents = append(fileContents, line)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		return []string{}
	}
	return fileContents
}
