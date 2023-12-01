package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FsObjectType int

const (
	dirObj FsObjectType = iota
	fileObj
)

// A Filesystem object. Can be a plain file or a directory. Several FsObjects
// can link to form a tree.
type FsObject struct {
	name     string
	size     int
	objType  FsObjectType
	parent   *FsObject
	children map[string]*FsObject
}

// Return a pointer to a new directory object.
func NewDirObj(name string, parent *FsObject) *FsObject {
	children := make(map[string]*FsObject, 0)

	newObject := FsObject{
		name:     name,
		size:     0,
		objType:  dirObj,
		parent:   parent,
		children: children,
	}

	return &newObject
}

// Return a pointer to a new plain file object.
func NewFileObj(name string, size int, parent *FsObject) *FsObject {
	newObject := FsObject{
		name:     name,
		size:     size,
		objType:  fileObj,
		parent:   parent,
		children: nil,
	}

	return &newObject
}

// Register a directory as a child of another.
func AddChild(dir *FsObject, child *FsObject) {
	dir.children[child.name] = child
}

// Lookup children of a directory by name.
func FindChild(parent *FsObject, childName string) (*FsObject, error) {
	if parent.objType == fileObj {
		return nil, fmt.Errorf("Plain files (%v) have no children!", parent.name)
	}

	result := parent.children[childName]

	return result, nil
}

func PrintFsTree(node *FsObject, indentLevel int) {
	indent := strings.Repeat(" ", 4*indentLevel)

	fmt.Printf("%v (%v)\n", node.name, node.size)

	for _, object := range node.children {
		fmt.Printf(indent)

		if object.objType == dirObj {
			indentLevel++
			PrintFsTree(object, indentLevel)
			indentLevel--
		} else {
			fmt.Printf("%v (%v)\n", object.name, object.size)
		}
	}
}

func UpdateDirSizes(node *FsObject) {
	for _, object := range node.children {

		if object.objType == dirObj {
			UpdateDirSizes(object)
		}

		node.size += object.size
	}
}

func solveQuestion(node *FsObject, minDelete *int, reqSpace int) {
	potentialSpace := node.size

	if potentialSpace >= reqSpace && potentialSpace < *minDelete {
		*minDelete = node.size
	}

	for _, object := range node.children {
		if object.objType == dirObj {
			solveQuestion(object, minDelete, reqSpace)
		}
	}
}

func main() {
	input, err := os.Open("../input/input")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	errLog := log.New(os.Stderr, "Error: ", log.Lshortfile)

	scanner := bufio.NewScanner(input)

	rootNode := NewDirObj("/", nil)
	currentDir := rootNode

	scanner.Scan()

	for scanner.Scan() {
		inputLine := scanner.Text()
		tokens := strings.Split(inputLine, " ")
		nTokens := len(tokens)
		if nTokens < 2 || nTokens > 3 {
			errLog.Fatalln("Bad input!")
		}

		// The current line is a command.
		if tokens[0] == "$" {
			command := tokens[1]
			objName := ""

			if nTokens == 3 {
				objName = tokens[2]
			}

			if command == "ls" && nTokens == 2 {
				// Good input. Nothing to do. Continue to next
				// iteration.
				continue
			} else if command == "cd" && nTokens == 3 {
				// Handle special case `..`
				if objName == ".." {
					currentDir = currentDir.parent
					continue
				}

				// Locate the requested directory amongst
				// children.
				newDir, err := FindChild(currentDir, objName)
				if err != nil {
					errLog.Fatalln(err)
				}

				if newDir == nil {
					errLog.Fatalf("No child directory %v\n", objName)
				}

				// Change directory.
				currentDir = newDir
			} else {
				errLog.Fatalf("Unrecognised command `%v`.\n", inputLine)
			}

		} else {
			// We are reading the output of the last command.

			if len(tokens) != 2 {
				errLog.Fatalln("Bad input!")
			}

			var newObj *FsObject

			// Directory object
			if tokens[0] == "dir" {
				newObj = NewDirObj(tokens[1], currentDir)
			} else {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					errLog.Fatalln(err)
				}

				newObj = NewFileObj(tokens[1], size, currentDir)
			}

			AddChild(currentDir, newObj)
		}
	}

	UpdateDirSizes(rootNode)
	fmt.Println("Printing tree:")
	PrintFsTree(rootNode, 0)

	// Solve the question:
	unusedSpace := 70000000 - rootNode.size
	needSpace := 30000000 - unusedSpace
	minDelete := rootNode.size
	solveQuestion(rootNode, &minDelete, needSpace)
	fmt.Printf("\nSolution: %v\n", minDelete)
}
