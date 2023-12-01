package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var errLog = log.New(os.Stderr, "Error: ", log.Lshortfile)

func main() {
	input, err := os.Open("../input/test")
	if err != nil {
		errLog.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		inputLine := scanner.Text()
		fmt.Println(inputLine)
	}
}
