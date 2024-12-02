package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

/*
Part 1: Calculate the sum of the absolute differences between the two columns

Solution: Read the input file and store the left and right columns
in two separate arrays: left and right.

Sort the left and right arrays from smallest to largest, then iterate
through the arrays and calculate the absolute difference between the two columns.

-- Thoughts --
There is most certainly a more efficient approach that involves calculating the
total difference without sorting the arrays.
*/
func partOne() {
	var left []int
	var right []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		before, after, _ := strings.Cut(scanner.Text(), "   ")

		beforeNumber, err := strconv.Atoi(before)
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return
		}

		afterNumber, err := strconv.Atoi(after)
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return
		}

		left = append(left, beforeNumber)
		right = append(right, afterNumber)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(left)
	sort.Ints(right)

	var totalDifference int

	for i := 0; i < len(left); i++ {
		totalDifference += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println("Part 1: ", totalDifference)
}

/*
Part 2: Calculate the sum of the products of the left column numbers and the frequency of the right column numbers

Solution: Read the input file and store the left column numbers in an array called left.
Create a frequency map for the right column numbers.

Iterate through the left column numbers and calculate the product of the left column number
and the frequency of the right column number. Add the product to a total sum.
*/
func partTwo() {
	var left []int
	rightFrequencyMap := make(map[int]int)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		before, after, _ := strings.Cut(scanner.Text(), "   ")

		beforeNumber, err := strconv.Atoi(before)
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return
		}

		afterNumber, err := strconv.Atoi(after)
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return
		}

		left = append(left, beforeNumber)

		// Add right column numbers to a frequency map
		value, exists := rightFrequencyMap[afterNumber]
		if exists {
			rightFrequencyMap[afterNumber] = value + 1
		} else {
			rightFrequencyMap[afterNumber] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	var totalDifference int

	for i := 0; i < len(left); i++ {
		value, exists := rightFrequencyMap[left[i]]
		if exists {
			totalDifference += left[i] * value
		}
	}

	fmt.Println("Part 2: ", totalDifference)
}
