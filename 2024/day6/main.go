package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

type coordinate struct {
	x, y int
}

func (pos coordinate) coordinateSum(secondCoordinate coordinate) (newCoordinate coordinate) {
	return coordinate{
		pos.x + secondCoordinate.x,
		pos.y + secondCoordinate.y,
	}
}

func (pos coordinate) equal(otherPos coordinate) (res bool) {
	if (pos.x == otherPos.x) && (pos.y == otherPos.y) {
		res = true
	}
	return
}

// Obstacles
type tile int

const (
	FREE tile = iota
	OBSTACLE
)

type grid map[coordinate]tile

func parseInputGrid(input []string) (g grid) {
	g = make(map[coordinate]tile)
	for x, line := range input {
		runeLine := []rune(line)
		for y, val := range runeLine {
			var pos tile
			switch string(val) {
			case "^", ">", "<", "v", ".":
				pos = FREE
			case "#":
				pos = OBSTACLE
			}
			g[coordinate{x, y}] = pos
		}
	}
	return
}

// Obstacles
type direction int

const (
	SUD direction = iota
	NORD
	EST
	OVEST
)

func (dir direction) toTheRight() direction {
	switch dir {
	case SUD:
		return OVEST
	case OVEST:
		return NORD
	case EST:
		return SUD
	case NORD:
		return EST
	}
	panic(fmt.Sprint("Unsupported direction for turning %d", dir))
}

func (dir direction) moveOneStep() coordinate {
	switch dir {
	case NORD:
		return coordinate{-1, 0}
	case SUD:
		return coordinate{1, 0}
	case OVEST:
		return coordinate{0, -1}
	case EST:
		return coordinate{0, 1}
	}
	panic(fmt.Sprint("Unsupported direction for movemnt %d", dir))
}

type guard struct {
	position  coordinate
	direction direction
}

func parseOneGuard(input []string) *guard {
	rGuard := regexp.MustCompile("\\^|\\>|\\<|v")
	for x, line := range input {
		if res := rGuard.FindAllString(line, -1); len(res) == 1 && len(res[0]) > 0 {
			y := rGuard.FindAllStringIndex(line, -1)[0][0]
			var dir direction
			switch res[0] {
			case "^":
				dir = NORD
			case ">":
				dir = EST
			case "<":
				dir = OVEST
			case "v":
				dir = SUD
			default:
				panic(fmt.Sprintf("wtf the regex did not work lol %v", res))
			}
			return &guard{
				coordinate{x, y},
				dir,
			}
		}
	}
	panic("found no guard char in the input")
}

type museumEnv struct {
	museumGuard *guard
	museumGrid  grid
	maxX, maxY  int
	visited     map[coordinate]struct{} // Star One
}

func newMuseumEnv(input []string) *museumEnv {
	g := parseInputGrid(input)
	maxX := lo.Max(lo.Map(lo.Keys(g), func(coord coordinate, _ int) int {
		return coord.x
	}))
	maxY := lo.Max(lo.Map(lo.Keys(g), func(coord coordinate, _ int) int {
		return coord.y
	}))

	mEnv := &museumEnv{
		museumGuard: parseOneGuard(input),
		museumGrid:  g,
		maxX:        maxX,
		maxY:        maxY,
		visited:     map[coordinate]struct{}{},
	}
	mEnv.visited[mEnv.museumGuard.position] = struct{}{}

	return mEnv
}

func (mEnv *museumEnv) step() (done bool, hit) {
	// First check if the next position is an obstacle
	possibleNewPos := mEnv.museumGuard.position.coordinateSum(mEnv.museumGuard.direction.moveOneStep())
	if (possibleNewPos.x > mEnv.maxX) || (possibleNewPos.y > mEnv.maxY) ||
		(possibleNewPos.y < 0) || (possibleNewPos.y < 0) {
		done = true
		return
	}
	switch mEnv.museumGrid[possibleNewPos] {
	// ROTATE AND MOVE 1
	case OBSTACLE:
		mEnv.museumGuard.direction = mEnv.museumGuard.direction.toTheRight()
		newPos := mEnv.museumGuard.position.coordinateSum(mEnv.museumGuard.direction.moveOneStep())
		mEnv.museumGuard.position = newPos
		mEnv.visited[newPos] = struct{}{}

	// JUST UPDATE IF THE POSITION IS FREE
	case FREE:
		mEnv.museumGuard.position = possibleNewPos
		mEnv.visited[possibleNewPos] = struct{}{}
	}

	return
}

func (mEnv museumEnv) isInfiniteLoop() (res bool) {
	startPosition := mEnv.museumGuard.position
	var (
		done, hit bool
		hits  int
	)
	for true {
		done, hit = mEnv.step()
		hits += hit
		if hits>4 || done{
			break	
		} else if startPosition.equal(mEnv.museumGuard.position) {
			res = true
			break
		}
	}
	return
}

func starOne(input []string) {
	done := false
	mEnv := newMuseumEnv(input)
	start := time.Now()
	for !done {
		done = mEnv.step()
	}
	end := time.Now().Sub(start)

	fmt.Printf("Star One: %d in %#vs\n", len(lo.Keys(mEnv.visited)), end.Seconds())
}

// TODO: you have to run a isInfiniteLoop() for 2 grid configurations each step:
// - Obstacle to the left (of the direction)
// - Obstacle in front (of the direction)
// NB: 
// 	- You have to run the loop only once for guard struct unique value (use string as identifier).
//  - Only report once the loop for each coordinate (use mutex on map of coords and then count the values).
func starTwo(input []string) {
	fmt.Printf("Star Two: %d\n")
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	// starTwo(input)
}
