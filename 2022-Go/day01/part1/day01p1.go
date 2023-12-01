package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var maxCalories int
	var currentElfSum int

	for scanner.Scan() {
		sumString := scanner.Text()

		if sumString == "" {
			if maxCalories < currentElfSum {
				maxCalories = currentElfSum
			}

			currentElfSum = 0
			continue
		}

		numCalories, err := strconv.Atoi(sumString)
		if err != nil {
			log.Fatalln(err)
		}

		currentElfSum += numCalories
	}

	fmt.Printf("Max Calories: %d\n", maxCalories)
}
