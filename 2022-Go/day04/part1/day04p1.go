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
	numFullOverlaps := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
		var sectionRanges [4]int

		fmt.Sscanf(inputLine, "%d-%d,%d-%d", &sectionRanges[0],
			&sectionRanges[1], &sectionRanges[2], &sectionRanges[3])

		fullOverlap := isFullOverlap(sectionRanges)

		switch fullOverlap {
		case 1, 2:
			fmt.Printf("%v: Range %v fully contains the other.\n", inputLine, fullOverlap)
			numFullOverlaps++
		default:
			fmt.Printf("%v: No complete overlap.\n", inputLine)
		}
	}

	fmt.Printf("\nNumber of full overlaps: %v\n", numFullOverlaps)
}

// Returns 1 if the first range fully contains the second.
// Returns 2 if the second range fully contains the first.
// Returns 0 otherwise.
func isFullOverlap(sectionRanges [4]int) int {
	// First range may contain second.
	if sectionRanges[0] <= sectionRanges[2] {
		if sectionRanges[1] >= sectionRanges[3] {
			return 1
		}
	}

	// Second range may contain first.
	if sectionRanges[2] <= sectionRanges[0] {
		if sectionRanges[1] <= sectionRanges[3] {
			return 2
		}
	}

	return 0
}
