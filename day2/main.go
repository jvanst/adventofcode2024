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
	direction := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

		isCurrentReportSafe := true
		direction = ""

		for i := 1; i < len(levels); i++ {
			previousLevel, err := strconv.Atoi(levels[i-1])
			if err != nil {
				fmt.Println("Error converting to number:", err)
				return
			}

			currentLevel, err := strconv.Atoi(levels[i])
			if err != nil {
				fmt.Println("Error converting to number:", err)
				return
			}

			difference := currentLevel - previousLevel

			if difference > 3 || difference < -3 || difference == 0 {
				isCurrentReportSafe = false
				break
			}

			if (direction == "down" && difference > 0) || (direction == "up" && difference < 0) {
				isCurrentReportSafe = false
				break
			}

			if difference < 0 {
				direction = "down"
			} else {
				direction = "up"
			}
		}

		if isCurrentReportSafe {
			numberOfSafeReports++
		}
	}

	fmt.Println("Part 1: ", numberOfSafeReports)
}

/*
 */
func partTwo() {

}
