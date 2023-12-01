package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// An (x,y) coordinate.
type Position [2]int

// Simple set implementation.
type Set map[Position]struct{}

func (s Set) IsMember(m Position) bool {
	if _, ok := s[m]; ok {
		return true
	}

	return false
}

func (s Set) AddMember(m Position) bool {
	if !s.IsMember(m) {
		s[m] = struct{}{}
		return true
	}

	return false
}

func (s Set) DelMember(m Position) {
	if s.IsMember(m) {
		delete(s, m)
	}
}

// Type to hold the rope state. Needs to store the head and tail position of the
// rope.
type Rope struct {
	hPos Position
	tPos Position
}

func NewRope() *Rope {
	return &Rope{hPos: Position{}, tPos: Position{}}
}

// An instruction stores the change in position to be made.
// The first element strores the direction and the second stores the amount to
// move by.
type Instruction [2]int

const (
	up int = iota
	down
	left
	right
)

// If the head is too far from the tail, then the tail will be pulled along.
// Check if this needs to happen.
func (r *Rope) isStretched() bool {
	xDiff := int(math.Abs(float64(r.hPos[0] - r.tPos[0])))
	yDiff := int(math.Abs(float64(r.hPos[1] - r.tPos[1])))

	if xDiff > 1 || yDiff > 1 {
		return true
	}

	return false
}

// Parses the input file for movement instructions and returns an Instruction
// type.
func parseInstruction(move string) (Instruction, error) {
	var direction string
	var amount int

	_, err := fmt.Sscanf(move, "%s %d", &direction, &amount)

	if err != nil {
		return Instruction{}, err
	}

	switch direction {
	case "U":
		return Instruction{up, amount}, nil
	case "D":
		return Instruction{down, amount}, nil
	case "L":
		return Instruction{left, amount}, nil
	case "R":
		return Instruction{right, amount}, nil
	default:
		return Instruction{}, fmt.Errorf("Invalid instruction %v", move)
	}
}

var errLog = log.New(os.Stderr, "Error: ", log.Lshortfile)

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		errLog.Fatalln(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// Initialise a new rope at the origin.
	rope := NewRope()

	// Visited coordinates
	visited := make(Set, 0)
	visited.AddMember(rope.tPos)

	// Keep track of the coordinates the tail has been at least once. start
	// at (0,0). For each instruction update the tail coordinate and it to a
	// set map[coord]struct{}. Then the final answer is len(map).
	for scanner.Scan() {
		inputLine := scanner.Text()

		// Get next move.
		ins, err := parseInstruction(inputLine)
		if err != nil {
			errLog.Fatalln(err)
		}

		for i := 0; i < ins[1]; i++ {
			direction := ins[0]
			oldHeadPos := rope.hPos

			// Update head position.
			switch direction {
			case up:
				rope.hPos[1]++
			case down:
				rope.hPos[1]--
			case left:
				rope.hPos[0]--
			case right:
				rope.hPos[0]++
			}

			// Update tail position if needed.
			if rope.isStretched() {
				rope.tPos = oldHeadPos
				visited.AddMember(rope.tPos)
			}

		}

	}

	fmt.Println("Positions visited by tail: ", len(visited))
}
