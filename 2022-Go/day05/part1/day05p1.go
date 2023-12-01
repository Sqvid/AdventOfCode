package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Stack struct {
	data []string
	size int
}

func NewStack() *Stack {
	return &Stack{data: make([]string, 0), size: 0}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Peek() (string, error) {
	if s.size == 0 {
		return "", errors.New("Peeking an empty stack.")
	}

	return s.data[s.size-1], nil
}

func (s *Stack) Pop() (string, error) {
	if s.size == 0 {
		return "", errors.New("Cannot pop from empty stack.")
	}

	poppedItem := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--

	return poppedItem, nil
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
	s.size++
}

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	var stackList []Stack
	var instructions []int

	scanner := bufio.NewScanner(input)
	stackInput := make([]byte, 0, 512)

	for scanner.Scan() {
		inputLine := scanner.Text()

		if len(inputLine) < 2 {
			continue
		}
		secondChar := rune(inputLine[1])

		// Everything above the row with numbers are stack items.
		if unicode.IsUpper(rune(inputLine[1])) || inputLine[1] == 32 {
			for i := 1; i < len(inputLine); i += 4 {
				ch := inputLine[i]
				stackInput = append(stackInput, ch)
			}

			// Not done reading in stack inputs. Loop back.
			continue
		} else if unicode.IsNumber(secondChar) {
			numStacks := (len(inputLine) + 1) / 4
			stackList = make([]Stack, numStacks, 10*numStacks)

			// Load the items into the stacks. We have to go in
			// reverse order to load the stacks from bottom to top.
			currentStackNum := numStacks - 1
			for i := len(stackInput) - 1; i >= 0; i-- {
				ch := stackInput[i]

				// Push the current byte onto the current stack.
				if string(ch) != " " {
					stackList[currentStackNum].Push(string(ch))
				}

				// Switch to the previous stack. Wrap if
				// necessary.
				if currentStackNum == 0 {
					currentStackNum = numStacks - 1
				} else {
					currentStackNum--
				}
			}
		} else if unicode.IsLower(secondChar) {
			var numItems, srcStack, destStack int

			fmt.Sscanf(inputLine, "move %d from %d to %d\n",
				&numItems, &srcStack, &destStack)

			instructions = append(instructions, numItems, srcStack,
				destStack)
		}
	}

	err = rearrangeStacks(stackList, instructions)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	topItems := ""

	for _, s := range stackList {
		topOfStack, _ := s.Peek()

		topItems = topItems + topOfStack
	}

	fmt.Println(topItems)
}

func printStacks(stackList []Stack) {
	// Print stacks
	for i, s := range stackList {
		fmt.Printf("Stack %v, size: %v\n", i+1, s.Size())

		for s.Size() > 0 {
			val, _ := s.Pop()

			fmt.Println(val)
		}
	}
}

func rearrangeStacks(stackList []Stack, instructions []int) error {
	if len(instructions)%3 != 0 {
		return errors.New("Instructions must come in sets of three.")
	}

	for i := 0; i < len(instructions); i += 3 {
		numItems := instructions[i]

		// Convert 1-indexing to 0-indexing.
		srcStack := instructions[i+1] - 1
		destStack := instructions[i+2] - 1

		// Error if trying to move more items than are available.
		if numItems > stackList[srcStack].Size() {
			return fmt.Errorf("Stack %v doesn't have %v values.",
				srcStack+1, numItems)
		}

		// Pop items off the source stack and push them onto the
		// destination stack.
		for i := 0; i < numItems; i++ {
			item, err := stackList[srcStack].Pop()

			if err != nil {
				return err
			}

			stackList[destStack].Push(item)
		}
	}

	return nil
}
