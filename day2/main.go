package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

/*
Part 1: Determine the number of safe reports

Solution: Read the input file and split each line into levels. Iterate through
each level and compare it to the previous level.

If the difference is greater than 3, less than -3 or 0 the report is not safe. When the
direction changes from up to down or down to up, the report is not safe.
*/
func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numberOfSafeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

		isSafe := isReportSafe(levels)

		if isSafe {
			numberOfSafeReports++
		}
	}

	fmt.Println("Part 1: ", numberOfSafeReports)
}

/*
Part 2: Determine the number of safe reports after removing one element ("the Problem Dampener")

Solution: Same as part 1, but we brute force all combinations of removing one element from the levels.

**Thoughts**
This solution is not efficient. It has a time complexity of O(n^2) where n is the number of levels in the report.
There must be a better way to solve this problem.
*/
func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numberOfSafeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

		if isReportSafe(levels) {
			numberOfSafeReports++
			continue
		}

		isSafe := false

		for i := 0; i < len(levels); i++ {
			levelsCopy := make([]string, len(levels))
			copy(levelsCopy, levels)

			if isReportSafe(append(levelsCopy[:i], levelsCopy[i+1:]...)) {
				isSafe = true
				break
			}
		}

		if isSafe {
			numberOfSafeReports++
		}
	}

	fmt.Println("Part 2: ", numberOfSafeReports)
}

func isReportSafe(levels []string) bool {
	direction := ""

	for i := 1; i < len(levels); i++ {
		previousLevel, err := strconv.Atoi(levels[i-1])
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return false
		}

		currentLevel, err := strconv.Atoi(levels[i])
		if err != nil {
			fmt.Println("Error converting to number:", err)
			return false
		}

		difference := currentLevel - previousLevel

		if (difference > 3 || difference < -3 || difference == 0) || (direction == "down" && difference > 0) || (direction == "up" && difference < 0) {
			return false
		}

		if difference < 0 {
			direction = "down"
		} else {
			direction = "up"
		}
	}

	return true
}
