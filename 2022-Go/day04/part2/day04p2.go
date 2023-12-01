package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// Number of fully overlapping ranges.
	numOverlaps := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
		var sectionRanges [4]int

		fmt.Sscanf(inputLine, "%d-%d,%d-%d", &sectionRanges[0],
			&sectionRanges[1], &sectionRanges[2], &sectionRanges[3])

		if isOverlap(sectionRanges) {
			numOverlaps++
		}
	}

	fmt.Printf("Number of overlaps: %v\n", numOverlaps)
}

// Checks if the section ranges overlap at all.
func isOverlap(sectionRanges [4]int) bool {
	firstStart, firstEnd, secondStart, secondEnd :=
		sectionRanges[0], sectionRanges[1], sectionRanges[2], sectionRanges[3]

	if secondStart >= firstStart && secondStart <= firstEnd {
		return true
	}

	if firstStart >= secondStart && firstStart <= secondEnd {
		return true
	}

	return false
}
