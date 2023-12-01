package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Set map[string]struct{}

func (s Set) IsMember(m string) bool {
	if _, ok := s[m]; ok {
		return true
	}

	return false
}

func (s Set) DelMember(m string) {
	if s.IsMember(m) {
		delete(s, m)
	}
}

func (s Set) AddMember(m string) bool {
	if !s.IsMember(m) {
		s[m] = struct{}{}
		return true
	}

	return false
}

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var commonItems []string

	for scanner.Scan() {
		inputLine := scanner.Text()
		fmt.Println(inputLine)
		rucksack1 := make(Set)
		rucksack2 := make(Set)

		for i, item := range inputLine {
			// Put the first half of the items into rucksack1 and
			// the remainder into rucksack2.
			if i < len(inputLine)/2 {
				rucksack1.AddMember(string(item))
			} else {
				rucksack2.AddMember(string(item))
			}
		}

		// Take members from rucksack1 and check if they are in
		// rucksack2. Since we are guaranteed only one match we can quit
		// when it is found.
		for k := range rucksack1 {
			if rucksack2.IsMember(k) {
				commonItems = append(commonItems, k)
				break
			}
		}
	}

	fmt.Println("Total points: ", calcPoints(commonItems))

}

func calcPoints(items []string) int {
	var total int

	for _, item := range items {
		var itemVal int

		asciiVal := []rune(item)[0]

		if unicode.IsLower(asciiVal) {
			// a = 97 in ASCII. Subtract 96 so it is worth 1 point.
			itemVal = int(asciiVal) - 96
			fmt.Printf("Lowercase: %v, ascii: %v\n", item, itemVal)
		} else if unicode.IsUpper(asciiVal) {
			// A = 65 in ASCII. Subtract 38 so it is worth 27
			// points.
			itemVal = int(asciiVal) - 38
			fmt.Printf("Uppercase: %v, ascii: %v\n", item, itemVal)
		}

		total += itemVal
	}

	return total
}
