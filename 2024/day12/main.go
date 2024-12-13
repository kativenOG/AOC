package main

import (
	"fmt"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

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

type grid map[coordinate]string

func (g grid) CoorValueString(coor coordinate) string {
	return fmt.Sprintf("Coordinate: %s, Value: %d", coor, g[coor])
}

func parseStringInputGrid(input []string) (g grid) {
	g = make(grid)
	for x, line := range input {
		runeLine := []rune(line)
		for y, val := range runeLine {
			g[coordinate{x, y}] = string(val)
		}
	}
	return
}

var (
	UP         = coordinate{1, 0}
	DOWN       = coordinate{-1, 0}
	LEFT       = coordinate{0, -1}
	RIGHT      = coordinate{0, 1}
	directions = []coordinate{UP, DOWN, LEFT, RIGHT}
)

func (g grid) validNeighbours(coord coordinate, targetVal string) (neighbours []coordinate) {
	var newPos coordinate
	for _, dir := range directions {
		newPos = coord.coordinateSum(dir)
		if coordVal, ok := g[newPos]; ok && coordVal == targetVal {
			neighbours = append(neighbours, newPos)
		}
	}
	return
}

func (g grid) findPlantation(startCoord coordinate, currentPlantation map[coordinate]struct{}) (plants []coordinate) {
	plantType := g[startCoord]
	neighbours := g.validNeighbours(startCoord, plantType)
	for _, coor := range neighbours {
		if _, ok := currentPlantation[coor]; !ok {
			currentPlantation[coor] = struct{}{}
			plants = append(plants, g.findPlantation(coor, currentPlantation)...)
		}
	}
	return append(plants, startCoord)
}

func (g grid) fenceCost(currentPlantation map[coordinate]struct{}, plantType string) (res int) {
	area := len(lo.Keys(currentPlantation))
	perimeter := lo.Sum(lo.Map(lo.Keys(currentPlantation), func(coor coordinate, _ int) int {
		return 4 - len(g.validNeighbours(coor, plantType))
	}))

	return area * perimeter
}

func starOne(input []string) {
	start := time.Now()

	var (
		res int
		g   = parseStringInputGrid(input)

		startCoor      coordinate
		visited        int
		haveToVisit    = len(lo.Values(g))
		alreadyVisited = make(map[coordinate]struct{})
	)

	for visited < haveToVisit {
		// Find a new Starting Point to look for a plantation
		for possibleStartCoor := range g {
			if _, ok := alreadyVisited[possibleStartCoor]; !ok {
				startCoor = possibleStartCoor
				break
			}
		}

		// Find all the adjacent plants of the same type
		plantType := g[startCoor]
		currentPlantation := lo.SliceToMap(
			g.findPlantation(startCoor, make(map[coordinate]struct{})),
			func(coor coordinate) (coordinate, struct{}) {
				return coor, struct{}{}
			},
		)

		// Update Visited Map
		lo.ForEach(lo.Keys(currentPlantation), func(coor coordinate, _ int) {
			alreadyVisited[coor] = struct{}{}
		})
		visited += len(lo.Keys(currentPlantation))

		// Calculate fence cost
		res += g.fenceCost(currentPlantation, plantType)

	}

	end := time.Now().Sub(start)
	fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
}

func starTwo(input []string) {
	start := time.Now()

	res := 0 // FIX
	end := time.Now().Sub(start)
	fmt.Printf("Star Two: %d in %#vs\n", res, end.Seconds())
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)
}
