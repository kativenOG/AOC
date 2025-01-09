package main

import (
	"fmt"
	"github.com/AOC/2024/utils"
	"time"
)

//////////////////////
// GRID COORDINATES //
//////////////////////

type coordinate struct {
	x, y int
}

func (coor coordinate) isEqual(target coordinate) bool {
	return (coor.x == target.x) && (coor.y == target.y)
}

func (c coordinate) coordinateSum(other coordinate) coordinate {
	return coordinate{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

//////////
// GRID //
//////////

type tile int

const (
	EMPTY tile = iota + 1
	WALL
	DEST
)

type grid map[coordinate]tile

func parseProblem(input []string) (rs reinderState, g grid, target coordinate) {
	g = make(grid)

	var foundTarget, foundStart bool
	for x, line := range input {
		runeline := []rune(line)
		for y, r := range runeline {

			coor := coordinate{
				x: x,
				y: y,
			}

			switch val := string(r); val {
			case "#":
				g[coor] = WALL
			case ".":
				g[coor] = EMPTY
			case "S":
				if foundStart {
					panic(fmt.Sprintf("already found a start at %s", rs.pos))
				}
				rs = reinderState{
					dir:      RIGHT,
					pos:      coor,
					cost:     0,
					previous: nil,
				}
				foundStart = true
			case "E":
				if foundTarget {
					panic(fmt.Sprintf("already found a target at %s", target))
				}
				target = coor
				foundTarget = true
			}
		}
	}

	if !foundTarget {
		panic("found no end")
	} else if !foundStart {
		panic("found no start")
	}

	return
}

///////////////
// DIRECTION //
///////////////

type direction int

const (
	UP direction = iota + 1
	DOWN
	LEFT
	RIGHT
)

func (dir direction) clockwise() (res direction) {
	switch dir {
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	}
	panic(fmt.Sprintf("in clockWise: direction does not exist: %d", dir))
}

func (dir direction) counterclockwise() direction {
	switch dir {
	case UP:
		return LEFT
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	case RIGHT:
		return UP
	}
	panic(fmt.Sprintf("in counterClockWise: direction does not exist: %d", dir))

}

func (d direction) dirVector() coordinate {
	switch d {
	case UP:
		return coordinate{0, 1}
	case DOWN:
		return coordinate{0, -1}
	case LEFT:
		return coordinate{1, 0}
	case RIGHT:
		return coordinate{-1, 0}
	}

	panic(fmt.Sprintf("while creating a vector for a direction: the direction value is not supported: %d", d))
}

type reinderState struct {
	dir      direction
	pos      coordinate
	cost     int
	previous *reinderState
}

func (rs reinderState) String() string {
	return fmt.Sprintf("%d-%d", rs.pos.x, rs.pos.y)
}

func (rs reinderState) Identifier() string {
	return fmt.Sprintf("%d-%d-%d", rs.pos.x, rs.pos.y, rs.dir)
}

func (rs reinderState) validActions(g grid) (newStates []reinderState) {
	// Forward (only an option if there is no wall and not out of bound)
	newPos := rs.pos.coordinateSum(rs.dir.dirVector())
	if tileVal, ok := g[newPos]; ok && tileVal == EMPTY {
		newStates = append(newStates, reinderState{
			dir:  rs.dir,
			pos:  newPos,
			cost: rs.cost + 1,
		})
	}

	// Clockwise turn
	newStates = append(newStates, reinderState{
		dir:  rs.dir.clockwise(),
		pos:  rs.pos,
		cost: rs.cost + 9999,
	})

	// Counter Clockwise turn
	newStates = append(newStates, reinderState{
		dir:  rs.dir.counterclockwise(),
		pos:  rs.pos,
		cost: rs.cost + 9999,
	})

	// DEBUG assertion:
	for _, target := range newStates {
		if (target.dir == 0 || target.pos == coordinate{0, 0}) {
			panic(fmt.Sprintf("valid action assertion: %#v", newStates))
		}
	}

	return
}

type frontier []reinderState

func createFrontier(s reinderState) frontier {
	return []reinderState{s}
}

func (f frontier) pop() (reinderState, frontier) {
	return f[0], f[1:]
}

func StarOne(input []string) {
	start := time.Now()
	startState, g, endCoor := parseProblem(input)

	// Result:
	fmt.Println(fmt.Sprintf("Start: %s", startState.Identifier()))
	fmt.Println(fmt.Sprintf("Start: %s", startState.Identifier()))

	frontier := createFrontier(startState)
	bestCostsMap := make(map[coordinate]int)
	expanded := map[string]struct{}{}

	var state reinderState
	for len(frontier) > 0 {
		// Pop the first last node added to the frontier to explore
		state, frontier = frontier.pop()

		// Save the current cost if its a new position or if its less than the previous one
		if _, ok := bestCostsMap[state.pos]; ok {
			bestCostsMap[state.pos] = min(bestCostsMap[state.pos], state.cost)
		} else {
			bestCostsMap[state.pos] = state.cost
		}

		// Don't expand multiple times the same position
		id := state.Identifier()
		if _, ok := expanded[id]; !ok {
			actions := state.validActions(g)
			for _, action := range actions {
				if (state.previous != nil) && (state.previous.String() == action.String()) {
					continue
				}
				frontier = append(frontier, action)
			}
			expanded[id] = struct{}{}
		}
	}

	if _, ok := bestCostsMap[endCoor]; !ok {
		fmt.Println(expanded)
		panic("havent found the end state")
	}

	fmt.Printf("Star One: %d, in %fs\n", bestCostsMap[endCoor], time.Now().Sub(start).Seconds())
}

func StarTwo(input []string) {

}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	StarOne(input)
	// StarTwo(input)
}
