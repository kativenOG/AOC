package main

import (
	"fmt"
	"slices"
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

func (g grid) nonValidNeighbours(coord coordinate, targetVal string) (neighbours []coordinate) {
	var newPos coordinate
	for _, dir := range directions {
		newPos = coord.coordinateSum(dir)
		if coordVal, ok := g[newPos]; ok && coordVal != targetVal {
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

func parseNonSequentialCoords(coorArr []coordinate) []coordinate {
	if len(coorArr) == 0 {
		return coorArr
	}

	// Get the axis
	var sameYaxis bool
	if coorArr[0].y == coorArr[0].y {
		sameYaxis = true
	}

	newCoorArr := make([]coordinate, 0, len(coorArr))
	slices.SortFunc(coorArr, func(a, b coordinate) int {
		valA, valB := a.y, b.y
		if sameYaxis {
			valA, valB = a.x, b.x
		}

		if valA == valB {
			return 0
		}
		if valA > valB {
			return 1
		}
		return -1
	})

	var previousVal int
	for i, coor := range coorArr {
		newVal := coor.y
		if sameYaxis {
			newVal = coor.x
		}

		if i == 0 {
			newCoorArr = append(newCoorArr, coor)
			previousVal = newVal
			continue
		}

		if (previousVal + 1) != newVal {
			break
		}

		newCoorArr = append(newCoorArr, coor)
		previousVal = newVal
	}

	return newCoorArr
}

func (g grid) discountedFenceCost(currentPlantation map[coordinate]struct{}, plantType string) (res int) {

	var (
		sides = 0
		area  = len(lo.Keys(currentPlantation))
	)

	visitedCounter := make(map[coordinate]int)
	lo.ForEach(lo.Keys(currentPlantation), func(coor coordinate, _ int) {
		nodes := g.nonValidNeighbours(coor, plantType)
		lo.ForEach(nodes, func(coor coordinate, _ int) {
			visitedCounter[coor] += 1
		})
	})
	outerNeighbours := lo.Keys(visitedCounter)

	var (
		targetSide coordinate
	)

	for lo.Sum(lo.Values(visitedCounter)) != 0 {
		var found bool
		for _, possibleTarget := range outerNeighbours {
			if val := visitedCounter[possibleTarget]; val > 0 {
				targetSide = possibleTarget
				found = true
				break
			}
		}
		if !found {
			utils.DieOnError(fmt.Errorf("wtf man, should've already exited"))
		}

		sameCoords := lo.Filter(outerNeighbours, func(coor coordinate, _ int) bool {
			if val := visitedCounter[coor]; val == 0 {
				return false
			}
			return targetSide.x == coor.x
		})

		if len(sameCoords) == 0 {
			sameCoords = lo.Filter(outerNeighbours, func(coor coordinate, _ int) bool {
				if val := visitedCounter[coor]; val == 0 {
					return false
				}
				return targetSide.y == coor.y
			})
		}

		if len(sameCoords) > 0 {
			sides += 1
			sameCoords = parseNonSequentialCoords(sameCoords)
			lo.ForEach(sameCoords, func(coord coordinate, _ int) {
				visitedCounter[coord] -= 1
			})
		}
	}

	fmt.Printf("%d - %d\n", area, sides)
	return area * sides
}

func starOne(input []string, discount bool) {
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
		if discount {
			res += g.discountedFenceCost(currentPlantation, plantType)
		} else {
			res += g.fenceCost(currentPlantation, plantType)
		}

	}

	end := time.Now().Sub(start)
	if !discount {
		fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
	} else {
		fmt.Printf("Star Two: %d in %#vs\n", res, end.Seconds())
	}
}

func starTwo(input []string) {
	starOne(input, true)
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input, false)
	starTwo(input)
}
