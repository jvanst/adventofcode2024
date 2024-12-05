package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

func main() {
	partOne()
}

/*
Part 1: Build a map of page ordering rules, then validate printing updates against the rules

Solution: To start, we build the page ordering map by reading all lines that contain a pipe. The
map will contain the page number as the key and the pages it must be printed before as the value.

Next, we iterate through the lines that contain a comma. For each line, we iterate backwards through
the page numbers. For each page number, we add the pages it must be printed before to a list of invalid
pages. If the current page number is in the list of invalid pages, we break. If we reach the end of the
page numbers, we add the middle page number to the total.
*/
func partOne() {
	fileContent, err := os.ReadFile(INPUT)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0

	lines := strings.Split(string(fileContent), "\n")

	pageOrderingRules := make(map[string][]string)

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		// If the line contains a pipe, then it's a page ordering rule
		if strings.Contains(lines[lineIndex], "|") {
			pieces := strings.Split(lines[lineIndex], "|")
			pageOrderingRules[pieces[0]] = append(pageOrderingRules[pieces[0]], pieces[1])
			continue
		}

		// If the line contains a comma, then it's a printing update
		if strings.Contains(lines[lineIndex], ",") {
			pageNumbers := strings.Split(lines[lineIndex], ",")

			invalidPages := make([]string, 0)

			// Iterate backwards through the pieces
			for pageNumberIndex := len(pageNumbers) - 1; pageNumberIndex >= 0; pageNumberIndex-- {
				invalidPages = append(invalidPages, pageOrderingRules[pageNumbers[pageNumberIndex]]...)

				// If invalidPages contains the current page number, then we can break
				if contains(invalidPages, pageNumbers[pageNumberIndex]) {
					break
				}

				// If we've reached the end of the page numbers, add the middle page to the total
				if pageNumberIndex == 0 {
					middlePageNumber, err := strconv.Atoi(pageNumbers[len(pageNumbers)/2])
					if err != nil {
						fmt.Println("Error converting to number:", err)
						return
					}

					total += middlePageNumber
				}
			}
		}
	}

	fmt.Println("Part 1: ", total)
}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
