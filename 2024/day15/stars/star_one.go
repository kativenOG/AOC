package stars

import (
	"fmt"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

func (wh *warehouse) move(input string) {
	var dir coordinate
	switch input {
	case "v":
		dir = coordinate{1, 0}
	case "^":
		dir = coordinate{-1, 0}
	case ">":
		dir = coordinate{0, 1}
	case "<":
		dir = coordinate{0, -1}
	default:
		utils.DieOnError(fmt.Errorf("wtf, wrong direction input %s", input))
	}

	newRobotPos := wh.robot.coordinateSum(dir)
	switch tile := wh.g[newRobotPos]; tile {
	case ROBOT:
		panic("there can only be one warehouse robot")
	case WALL:
		return
	case EMPTY:
		wh.g[wh.robot] = EMPTY
		wh.g[newRobotPos] = ROBOT
		wh.robot = newRobotPos
	case BLOCK:
		var foundEmpty bool
		newCratePositions := []coordinate{}
		for newCratePos := newRobotPos.coordinateSum(dir); wh.g[newCratePos] == BLOCK || wh.g[newCratePos] == EMPTY; newCratePos = newCratePos.coordinateSum(dir) {
			newCratePositions = append(newCratePositions, newCratePos)
			if wh.g[newCratePos] == EMPTY {
				foundEmpty = true
				break
			}
		}

		if !foundEmpty {
			return
		}

		lo.ForEach(newCratePositions, func(pos coordinate, _ int) {
			wh.g[pos] = BLOCK
		})
		wh.g[wh.robot] = EMPTY
		wh.g[newRobotPos] = ROBOT
		wh.robot = newRobotPos
	}
}

func furtherParse(input []string) (wh warehouse, actions []string) {
	var foundEmptyLine bool
	wh.g = make(grid)

	for x, line := range input {
		if line == "" {
			foundEmptyLine = true
			continue
		}
		if !foundEmptyLine { // GRID
			runeline := []rune(line)
			for y, rune := range runeline {
				coor := coordinate{x, y}
				tileT, robot := parseTile(string(rune))
				if !robot {
					wh.g[coor] = tileT
				} else {
					wh.g[coor] = tileT
					wh.robot = coor
				}
			}
		} else { // ACTIONS
			runeline := []rune(line)
			for _, a := range runeline {
				actions = append(actions, string(a))
			}
		}
	}

	return
}

func StarOne(input []string, visualization bool) {
	var (
		start       = time.Now()
		wh, actions = furtherParse(input)
	)

	// Run all the Actions
	if visualization {
		wh.visualize("START")
	}
	for _, action := range actions {
		wh.move(action)
		if visualization {
			wh.visualize(action)
		}
	}

	end := time.Now().Sub(start)
	fmt.Printf("Star One: %d in %fs\n", wh.g.gpsScore(), end.Seconds())
}
