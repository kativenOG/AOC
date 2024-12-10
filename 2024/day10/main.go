package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func dieOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInputFile(filename string) (inputList []string) {
	content, err := os.ReadFile(filename)
	dieOnError(err)
	inputList = strings.Split(string(content), "\n")
	if len(inputList[len(inputList)-1]) == 0 {
		inputList = inputList[:len(inputList)-1]
	}
	return
}

type coordinate struct {
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("[%d-%d]", c.x, c.y)
}

func (pos coordinate) coordinateSum(secondCoordinate coordinate) (newCoordinate coordinate) {
	return coordinate{
		pos.x + secondCoordinate.x,
		pos.y + secondCoordinate.y,
	}
}

type grid map[coordinate]int

func (g grid) CoorValueString(coor coordinate) string {
	return fmt.Sprintf("Coordinate: %s, Value: %d", coor, g[coor])
}

func parseInputGrid(input []string) (g grid) {
	g = make(map[coordinate]int)
	for x, line := range input {
		runeLine := []rune(line)
		for y, val := range runeLine {
			val, err := strconv.Atoi(string(val))
			dieOnError(err)
			g[coordinate{x, y}] = val
		}
	}
	return
}

func (dir coordinate) allPossibleCoordinates(maxX, maxY int) []coordinate {
	res := []coordinate{}
	// NORD
	if (dir.x - 1) >= 0 {
		res = append(res, dir.coordinateSum(coordinate{-1, 0}))
	}
	// SUD
	if (dir.x + 1) <= maxX {
		res = append(res, dir.coordinateSum(coordinate{1, 0}))
	}
	// OVEST
	if (dir.y - 1) >= 0 {
		res = append(res, dir.coordinateSum(coordinate{0, -1}))
	}
	// EST
	if (dir.y + 1) <= maxY {
		res = append(res, dir.coordinateSum(coordinate{0, 1}))

	}
	return res
}

type trailHead struct {
	currentPos      coordinate
	currentPosValue int

	maxX, maxY int
}

func (head trailHead) String() string {
	return fmt.Sprintf("HEAD at coords=%s with value %d\n", head.currentPos, head.currentPosValue)
}

func findTrailHeads(g grid, maxX, maxY int) []trailHead {
	tHeads := []trailHead{}
	for key, value := range g {
		if value == 0 {
			newHead := trailHead{
				currentPos:      key,
				currentPosValue: value,
				maxX:            maxX,
				maxY:            maxY,
			}
			tHeads = append(tHeads, newHead)
		}
	}
	return tHeads
}

func recursiveTrailExploration(head trailHead, g grid, explored map[coordinate]struct{}, useExplored bool) (topCounter int) {
	if g[head.currentPos] == 9 {
		return 1
	}
	for _, pCoord := range head.currentPos.allPossibleCoordinates(head.maxX, head.maxY) {
		_, ok := explored[pCoord]
		pValue := g[pCoord]
		if ((pValue - head.currentPosValue) == 1) && (!ok || !useExplored) {
			newHead := trailHead{
				currentPos:      pCoord,
				currentPosValue: pValue,
				maxX:            head.maxX,
				maxY:            head.maxY,
			}
			explored[pCoord] = struct{}{}
			topCounter += recursiveTrailExploration(newHead, g, explored, useExplored)
		}
	}
	return
}

func starOne(input []string) {
	start := time.Now()

	maxX, maxY := len(input[0]), len(input)
	g := parseInputGrid(input)
	tHeads := findTrailHeads(g, maxX, maxY)

	res := 0
	for _, head := range tHeads {
		explored := map[coordinate]struct{}{
			head.currentPos: struct{}{},
		}
		tops := recursiveTrailExploration(head, g, explored, true)
		res += tops
	}
	end := time.Now().Sub(start)

	fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
}

func starTwo(input []string) {
	start := time.Now()

	maxX, maxY := len(input[0]), len(input)
	g := parseInputGrid(input)
	tHeads := findTrailHeads(g, maxX, maxY)

	res := 0
	for _, head := range tHeads {
		explored := map[coordinate]struct{}{
			head.currentPos: struct{}{},
		}
		tops := recursiveTrailExploration(head, g, explored, false)
		res += tops
	}
	end := time.Now().Sub(start)

	fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.Parse()
	input := parseInputFile(filename)

	starOne(input)
	starTwo(input)
}
