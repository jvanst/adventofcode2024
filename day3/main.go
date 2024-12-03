package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

/*
Part 1: Determine the total of all multiplications in the input file

Solution:
A clear use case for regex, to parse out only valid mul() functions and extract the two numbers.

First, read the file content and then use a regex to find all matches of mul() functions. Iterate through each match
and extract the two numbers. Multiply the two numbers and add it to the total.
*/
func partOne() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0

	r := regexp.MustCompile(`mul\(\b\d{1,3}\b,\b\d{1,3}\b\)`)
	matches := r.FindAllString(string(fileContent), -1)

	for _, match := range matches {
		// Extract the two numbers from the match
		numbers := strings.Split(match[4:len(match)-1], ",")

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return
		}

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return
		}

		total += num1 * num2
	}

	fmt.Println("Part 1: ", total)
}
