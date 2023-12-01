package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	lose int = 0
	draw     = 3
	win      = 6
)

const (
	rock int = iota + 1
	paper
	scissors
)

func main() {
	input, err := os.Open("../input/day02.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	totalScore := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		playerChoices := scanner.Text()
		var theyPlayed, needResult string

		fmt.Sscanf(playerChoices, "%s %s", &theyPlayed, &needResult)
		totalScore += roundResult(theyPlayed, needResult)
	}

	fmt.Println(totalScore)
}

func roundResult(theyPlayed, needResult string) int {
	resultNeeded := map[string]int{
		"X": lose,
		"Y": draw,
		"Z": win,
	}

	var iPlay int

	switch theyPlayed {
	case "A":
		switch needResult {
		case "X":
			iPlay = scissors
		case "Y":
			iPlay = rock
		case "Z":
			iPlay = paper
		}
	case "B":
		switch needResult {
		case "X":
			iPlay = rock
		case "Y":
			iPlay = paper
		case "Z":
			iPlay = scissors
		}
	case "C":
		switch needResult {
		case "X":
			iPlay = paper
		case "Y":
			iPlay = scissors
		case "Z":
			iPlay = rock
		}
	}

	return iPlay + resultNeeded[needResult]
}
