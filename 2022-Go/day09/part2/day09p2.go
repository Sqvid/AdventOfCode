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
	knots  []Position
	head   *Position
	tail   *Position
	nKnots int
}

func NewRope(nKnots int) *Rope {
	knots := make([]Position, nKnots)
	head := &knots[0]
	tail := &knots[nKnots-1]

	return &Rope{head: head, knots: knots, tail: tail, nKnots: nKnots}
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

// If one knot is too far from the other, and its position needs to be updated.
func isStretched(knot1, knot2 Position) bool {
	xDiff := int(math.Abs(float64(knot1[0] - knot2[0])))
	yDiff := int(math.Abs(float64(knot1[1] - knot2[1])))

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

func (r *Rope) Draw() {
	var minX, minY int
	maxX, maxY := 10, 10

	for _, knot := range r.knots {
		if knot[0] > maxX {
			maxX = knot[0]
		}

		if knot[0] < minX {
			minX = knot[0]
		}

		if knot[1] > maxY {
			maxY = knot[1]
		}

		if knot[1] < minY {
			minY = knot[1]
		}
	}

	graphWidth := maxX - minX + 1
	graphHeight := maxY - minY + 1
	graph := make([][]string, graphHeight)

	// Initialise graph
	for i := 0; i < graphHeight; i++ {
		for j := 0; j < graphWidth; j++ {
			graph[i] = append(graph[i], ".")
		}
	}

	// Draw knots
	for i, pos := range r.knots {
		graph[pos[1]-minY][pos[0]-minX] = fmt.Sprintf("%v", i)
	}

	// Draw origin
	graph[-minY][-minX] = "s"

	for i := len(graph) - 1; i >= 0; i-- {
		fmt.Printf("%3v", i+minY)
		fmt.Println(graph[i])
	}
	fmt.Printf("\n\n")
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
	rope := NewRope(10)

	// Visited coordinates
	visited := make(Set, 0)
	visited.AddMember(*rope.tail)

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
			//oldLeadKnotPos := *rope.head

			// Update head position.
			switch direction {
			case up:
				rope.head[1]++
			case down:
				rope.head[1]--
			case left:
				rope.head[0]--
			case right:
				rope.head[0]++
			}

			//rope.Draw()

			// Update knot positions if needed.
			for i := 0; i < rope.nKnots-1; i++ {
				if !isStretched(rope.knots[i], rope.knots[i+1]) {
					break
				}

				// When the lead knot moves it pulls the next
				// knot if the rope is stretched. When this
				// happens the pulled knot moves one step in the
				// direction of the lead knot's new position.
				leadKnot := &rope.knots[i]
				pulledKnot := &rope.knots[i+1]

				// Pull the knot in the appropriate directions
				// if needed.
				if leadKnot[0] < pulledKnot[0] {
					(*pulledKnot)[0]--
				} else if leadKnot[0] > pulledKnot[0] {
					(*pulledKnot)[0]++
				}

				if leadKnot[1] < pulledKnot[1] {
					(*pulledKnot)[1]--
				} else if leadKnot[1] > pulledKnot[1] {
					(*pulledKnot)[1]++
				}

				//rope.Draw()
			}

			visited.AddMember(*rope.tail)
		}
	}

	fmt.Println("Positions visited by tail: ", len(visited))
}
