package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
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

func (pos coordinate) coordinateSum(secondCoordinate coordinate) (newCoordinate coordinate) {
	return coordinate{
		pos.x + secondCoordinate.x,
		pos.y + secondCoordinate.y,
	}
}

type grid map[coordinate]int

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
	if (dir.x - 1) > 0 {
		res = append(res, dir.coordinateSum(coordinate{-1, 0}))
	}
	// SUD
	if (dir.x + 1) < maxX {
		res = append(res, dir.coordinateSum(coordinate{1, 0}))
	}

	// OVEST
	if (dir.y - 1) > 0 {
		res = append(res, dir.coordinateSum(coordinate{0, -1}))
	}
	// EST
	if (dir.y + 1) < maxY {
		res = append(res, dir.coordinateSum(coordinate{0, 1}))

	}
	return res
}

type trailHead struct {
	currentPos      coordinate
	currentPosValue int

	maxX, maxY int
}

func findTrailHeads(g grid, maxX, maxY int) []trailHead {
	tHeads := []trailHead{}
	for key, value := range g {
		if value == 0 {
			newHead := trailHead{
				currentPos:      key,
				currentPosValue: 0,
				maxX:            maxX,
				maxY:            maxY,
			}
			tHeads = append(tHeads, newHead)
		}
	}
	return tHeads
}

func recursiveTrailExploration(head trailHead, g grid) (topCounter int) {
	for _, pCoord := range head.currentPos.allPossibleCoordinates(head.maxX, head.maxY) {
		pValue := g[pCoord]
		if pValue == 9 {
			return 1
		} else if ((head.currentPosValue - pValue) == 1) || ((head.currentPosValue - pValue) == -1) {
			newHead := trailHead{
				currentPos:      pCoord,
				currentPosValue: pValue,
				maxX:            head.maxX,
				maxY:            head.maxY,
			}
			topCounter += recursiveTrailExploration(newHead, g)
		}
	}
	return
}

func starOne(input []string) {
	start := time.Now()

	maxX, maxY := len(input), len(input[0])
	g := parseInputGrid(input)
	tHeads := findTrailHeads(g, maxX, maxY)

	res := 0
	for _, head := range tHeads {
		res += recursiveTrailExploration(head, g)
	}
	end := time.Now().Sub(start)

	fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
}

func starTwo(input []string) {
	fmt.Printf("Star Two: %d\n")
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.Parse()
	input := parseInputFile(filename)

	starOne(input)
	// starTwo(input)
}
