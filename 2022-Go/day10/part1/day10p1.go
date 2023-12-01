package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var errLog = log.New(os.Stderr, "Error: ", log.Lshortfile)

func main() {
	input, err := os.Open("../input/test")
	if err != nil {
		errLog.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// The CPU cycle number.
	cycle := 1
	// Current register value.
	var register int = 1

	for scanner.Scan() {
		inputLine := scanner.Text()
		tokens := strings.Split(inputLine, " ")
		nTokens := len(tokens)

		// CPU instruction and value.
		instr := tokens[0]
		var val int

		switch instr {
		case "noop":
			if nTokens != 1 {
				errLog.Fatalln("noop takes no values.")
			}

			cycle++
		case "addx":
			if nTokens != 2 {
				errLog.Fatalln("addx needs 1 value.")
			}

			val, err = strconv.Atoi(tokens[1])
			if err != nil {
				errLog.Fatalln(err)
			}

			cycle += 2
			register += val
		default:
			errLog.Fatalf("Unrecognised instruction `%v`", instr)
		}

		if (cycle-20)%40 == 0 {
			fmt.Println(cycle, register)
		}
	}
}
