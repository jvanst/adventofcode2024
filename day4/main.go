package main

import (
	"fmt"
	"os"
	"strings"
)

const INPUT = "input.txt"

func main() {
	partOne()
	partTwo()
}

/*
Part 1: Find all instances of XMAS and SAMX

Solution: I've opted for a brute force solution. I iterate through each character in the input file and
search for the strings "XMAS" and "SAMX" in all* directions (right, down, diagonal down right, diagonal down left).

* I say all, but I'm only searching in the 4 directions mentioned above. This is possible because
I'm searching from top left, in the right direction. Searching up, left, diagonal up left, diagonal up right
is not necessary because the strings are the same in reverse.

**Thoughts**
I'm pretty sure there is something better here
*/
func partOne() {
	fileContent, err := os.ReadFile(INPUT)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0

	lines := strings.Split(string(fileContent), "\n")

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if searchRight(lines, x, y) {
				total++
			}
			if searchDown(lines, x, y) {
				total++
			}
			if searchDiagonalDownRight(lines, x, y) {
				total++
			}
			if searchDiagonalDownLeft(lines, x, y) {
				total++
			}
		}
	}

	fmt.Println("Part 1: ", total)
}

func searchRight(lines []string, x int, y int) bool {
	if y >= len(lines) || x+4 > len(lines[y]) {
		return false
	}
	return isMatch(string(lines[y][x : x+4]))
}

func searchDown(lines []string, x int, y int) bool {
	if y+4 > len(lines) {
		return false
	}

	return isMatch(string(lines[y][x]) + string(lines[y+1][x]) + string(lines[y+2][x]) + string(lines[y+3][x]))
}

func searchDiagonalDownRight(lines []string, x int, y int) bool {
	if y+4 > len(lines) || x+4 > len(lines[y]) {
		return false
	}

	return isMatch(string(lines[y][x]) + string(lines[y+1][x+1]) + string(lines[y+2][x+2]) + string(lines[y+3][x+3]))
}

func searchDiagonalDownLeft(lines []string, x int, y int) bool {
	if y+4 > len(lines) || x-3 < 0 {
		return false
	}

	return isMatch(string(lines[y][x]) + string(lines[y+1][x-1]) + string(lines[y+2][x-2]) + string(lines[y+3][x-3]))
}

func isMatch(substring string) bool {
	if substring == "XMAS" || substring == "SAMX" {
		return true
	}
	return false
}

/*
Part 2: Find all cases of two MAS in the shape of an X:

Ex:
M.S
.A.
M.S

Solution: Iterate over each character, disregard all characters that are not "A".
Then checks if character to the diagonal up/down left and diagonal up/down right is exist. If they do
create a string with the characters and check if it's either "MAS" or "SAM". If it is, increment the total.
*/
func partTwo() {
	fileContent, err := os.ReadFile(INPUT)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0

	lines := strings.Split(string(fileContent), "\n")

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if string(lines[y][x]) != "A" {
				continue
			}

			if y-1 >= 0 && y+1 < len(lines) && x-1 >= 0 && x+1 < len(lines[y]) {
				descending := string(lines[y-1][x-1]) + string(lines[y][x]) + string(lines[y+1][x+1])
				accending := string(lines[y-1][x+1]) + string(lines[y][x]) + string(lines[y+1][x-1])

				if (accending == "MAS" || accending == "SAM") && (descending == "MAS" || descending == "SAM") {
					total += 1
				}
			}
		}
	}

	fmt.Println("Part 2: ", total)
}
