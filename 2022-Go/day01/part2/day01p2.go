package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("../input/day01.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// Will store the top calorie scores.
	leaderboardSize := 3
	leaderboard := make([]int, leaderboardSize)

	// The total calories held by the current elf.
	currentElfSum := 0

	for scanner.Scan() {
		sumString := scanner.Text()

		if sumString == "" {
			placeInLeaderboard(currentElfSum, leaderboard)

			// Reset sum.
			currentElfSum = 0
			continue
		}

		numCalories, err := strconv.Atoi(sumString)
		if err != nil {
			log.Fatalln(err)
		}

		currentElfSum += numCalories
	}

	fmt.Println("Leaderboard: ", leaderboard)

	calorieSum := 0
	for i := 0; i < leaderboardSize; i++ {
		calorieSum += leaderboard[i]
	}

	fmt.Println("Sum: ", calorieSum)
}

func placeInLeaderboard(num int, leaderboard []int) {
	// num is not big enough to be on the leaderboard.
	if num <= leaderboard[0] {
		return
	}

	for i := range leaderboard {
		// No other options. num is at the top of the leaderboard.
		if i == len(leaderboard)-1 {
			leaderboard[i] = num
			return
		}

		next := leaderboard[i+1]

		if num < next {
			// num is less than the next value, put it here.
			leaderboard[i] = num
			return

		} else if num == next {
			// num already exists in the leaderboard.
			return
		}
	}
}
