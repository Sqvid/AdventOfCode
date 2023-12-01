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

func main() {
	input, err := os.Open("../input/day02.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	totalScore := 0

	for scanner.Scan() {
		playerChoices := scanner.Text()
		var iPlayed, theyPlayed string

		fmt.Sscanf(playerChoices, "%s %s", &theyPlayed, &iPlayed)
		totalScore += roundResult(theyPlayed, iPlayed)
	}

	fmt.Println(totalScore)
}

func roundResult(theyPlayed, iPlayed string) int {
	var result int

	loseCondition := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	drawCondition := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	shapeScore := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	if iPlayed == loseCondition[theyPlayed] {
		result = lose
	} else if iPlayed == drawCondition[theyPlayed] {
		result = draw
	} else {
		result = win
	}

	return shapeScore[iPlayed] + result
}
