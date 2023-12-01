package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) IsMember(m string) bool {
	if _, ok := s[m]; ok {
		return true
	}

	return false
}

func (s Set) AddMember(m string) bool {
	if !s.IsMember(m) {
		s[m] = struct{}{}
		return true
	}

	return false
}

func (s Set) DelMember(m string) {
	if s.IsMember(m) {
		delete(s, m)
	}
}

func FindIntersection(s1, s2 Set) Set {
	intersection := NewSet()

	for k := range s1 {
		if s2.IsMember(k) {
			intersection.AddMember(k)
		}
	}

	return intersection
}

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var badges []string

	for lineNum := 0; scanner.Scan(); lineNum++ {
		elfStr1 := scanner.Text()
		scanner.Scan()
		elfStr2 := scanner.Text()
		scanner.Scan()
		elfStr3 := scanner.Text()

		elfStrings := [3]string{elfStr1, elfStr2, elfStr3}

		var elfInventories [3]Set

		// Populate elf inventories with unique items.
		for i := range elfInventories {
			elfInventories[i] = NewSet()

			for _, ch := range elfStrings[i] {
				elfInventories[i].AddMember(string(ch))
			}
		}

		common := FindIntersection(elfInventories[0], FindIntersection(elfInventories[1], elfInventories[2]))

		for k := range common {
			badges = append(badges, k)
		}
	}

	fmt.Println("Total points: ", calcPoints(badges))

}

func calcPoints(items []string) int {
	var total int

	for _, item := range items {
		var itemVal int

		asciiVal := []rune(item)[0]

		if unicode.IsLower(asciiVal) {
			// a = 97 in ASCII. Subtract 96 so it is worth 1 point.
			itemVal = int(asciiVal) - 96
		} else if unicode.IsUpper(asciiVal) {
			// A = 65 in ASCII. Subtract 38 so it is worth 27
			// points.
			itemVal = int(asciiVal) - 38
		}

		total += itemVal
	}

	return total
}
