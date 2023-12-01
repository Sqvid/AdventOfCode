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
	scanner.Scan()
	dataStream := scanner.Text()

	windowSize := 13

	if len(dataStream) <= windowSize {
		fmt.Fprintln(os.Stderr, "Data-stream is shorter than 4 characters.")
		os.Exit(1)
	}

	for i := windowSize; i < len(dataStream); i++ {
		searchWindow := dataStream[i-windowSize : i]

		// Convert the search string into a map to easily check for
		// duplicates.
		searchMap := make(map[string]int)
		for i, ch := range searchWindow {
			searchMap[string(ch)] = i
		}

		currentSearch := string(dataStream[i])

		pos := posInString(searchMap, currentSearch)

		if pos == -1 {
			// There were no duplicates in the searchString
			if len(searchMap) == windowSize {
				// Print the position after the match.
				fmt.Println(i + 1)
				break
			}
		}

		// Advance the search windpw past the position where a match was
		// found.
		i += searchMap[currentSearch]
	}
}

// Lookup the index where search appears in the search window.
func posInString(window map[string]int, search string) int {
	if index, ok := window[search]; ok {
		return index
	}

	return -1
}
