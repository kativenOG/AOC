package stars

import (
	"fmt"
	"slices"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

func (g grid) moveBigBox(pos, dir coordinate) {
	switch tile := g[pos]; tile {
	case BIG_BLOCK_LEFT:
		otherPos := pos.coordinateSum(coordinate{0, 1})
		g[pos], g[otherPos] = EMPTY, EMPTY
		g[pos.coordinateSum(dir)] = BIG_BLOCK_LEFT
		g[otherPos.coordinateSum(dir)] = BIG_BLOCK_RIGHT
	case BIG_BLOCK_RIGHT:
		otherPos := pos.coordinateSum(coordinate{0, -1})
		g[pos], g[otherPos] = EMPTY, EMPTY
		g[pos.coordinateSum(dir)] = BIG_BLOCK_RIGHT
		g[otherPos.coordinateSum(dir)] = BIG_BLOCK_LEFT
	default:
		panic(fmt.Sprintf("in grid.moveBigBox trying to move a tile that is not a box: %s", reverseParseTile(tile)))
	}
}

func (g grid) getBigBox(pos coordinate) (res []coordinate) {
	var otherPos coordinate
	switch tile := g[pos]; tile {
	case BIG_BLOCK_LEFT:
		otherPos = pos.coordinateSum(coordinate{0, 1})
		res = append(res, pos, otherPos)
	case BIG_BLOCK_RIGHT:
		otherPos = pos.coordinateSum(coordinate{0, -1})
		res = append(res, otherPos, pos)
	default:
		panic("in grid.getBigBox trying to get a parse a tile that is not a box")
	}
	return
}

func (coor coordinate) isDirectionHorizzontal() (res bool) {
	if coor.y != 0 {
		res = true
	}
	return
}

func (wh warehouse) canMoveVertically(startPos, dir coordinate, visitedBoxesSet map[coordinate]struct{}) (res bool) {
	res = true
	if dir.y != 0 {
		panic(fmt.Sprintf("%s this should be a vertical movement", dir))
	}

	switch newPos := startPos.coordinateSum(dir); wh.g[newPos] {
	case BIG_BLOCK_LEFT, BIG_BLOCK_RIGHT:
		bigBox := wh.g.getBigBox(newPos)
		for _, boxDir := range bigBox {
			if _, ok := visitedBoxesSet[boxDir]; !ok {
				res = wh.canMoveVertically(boxDir, dir, visitedBoxesSet)
				if !res {
					break
				}
			}
		}
		lo.ForEach(bigBox, func(coor coordinate, _ int) {
			visitedBoxesSet[coor] = struct{}{}
		})
	case WALL:
		return false
	case EMPTY:
		return true
	}

	return
}

func (wh *warehouse) moveBigBoxes(input string) {
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
	case BIG_BLOCK_LEFT, BIG_BLOCK_RIGHT:
		var foundEmpty bool
		if dir.isDirectionHorizzontal() {
			cratePositions := []coordinate{newRobotPos}
			for newCratePos := newRobotPos.coordinateSum(coordinate{dir.x, dir.y * 2}); wh.g[newCratePos] == BLOCK || wh.g[newCratePos] == EMPTY; newCratePos = newCratePos.coordinateSum(coordinate{dir.x, dir.y * 2}) {
				if wh.g[newCratePos] == EMPTY {
					foundEmpty = true
					break
				}
				cratePositions = append(cratePositions, newCratePos)
			}

			if !foundEmpty {
				return
			}

			// Create unique box set:
			singleBoxArray := lo.SliceToMap(
				lo.Map(cratePositions,
					func(coor coordinate, _ int) coordinate {
						return wh.g.getBigBox(coor)[0]
					},
				),
				func(coor coordinate) (val coordinate, appo struct{}) {
					return coor, struct{}{}
				},
			)
			lo.ForEach(lo.Reverse(lo.Keys(singleBoxArray)),
				func(pos coordinate, _ int) {
					wh.g.moveBigBox(pos, dir)
				},
			)
		} else {
			haveVisitedBoxes := make(map[coordinate]struct{})
			if res := wh.canMoveVertically(wh.robot, dir, haveVisitedBoxes); res {
				fmt.Println(haveVisitedBoxes)
				targets := lo.Filter(lo.Keys(haveVisitedBoxes), func(coor coordinate, _ int) bool {
					return wh.g[coor] == BIG_BLOCK_LEFT
				})

				// Move every box only once and sort them so they
				// do not override each other while updating positions.
				slices.SortFunc(targets, func(a, b coordinate) int {
					if a.x == b.x {
						return 0
					} else if a.x > b.x {
						return -1 * dir.x
					}
					return 1 * dir.x
				})
				lo.ForEach(targets, func(coor coordinate, _ int) {
					fmt.Println(coor)
					wh.g.moveBigBox(coor, dir)
				})
			} else {
				return
			}

		}

		// Always
		wh.g[wh.robot] = EMPTY
		wh.g[newRobotPos] = ROBOT
		wh.robot = newRobotPos
	}
}

func furtherBigParse(input []string) (wh warehouse, actions []string) {
	var foundEmptyLine bool
	wh.g = make(grid)

	for x, line := range input {
		if line == "" {
			foundEmptyLine = true
			continue
		}
		if !foundEmptyLine { // GRID
			runeline := []rune(line)
			var y int
			for _, rune := range runeline {
				tileT, robot := parseTile(string(rune))
				coorA, coorB := coordinate{x, y}, coordinate{x, y + 1}
				if robot {
					wh.robot = coorA
					wh.g[coorA] = ROBOT
					wh.g[coorB] = EMPTY
				} else if tileT == BLOCK {
					wh.g[coorA] = BIG_BLOCK_LEFT
					wh.g[coorB] = BIG_BLOCK_RIGHT
				} else if tileT == WALL {
					wh.g[coorA] = WALL
					wh.g[coorB] = WALL
				} else {
					wh.g[coorA] = EMPTY
					wh.g[coorB] = EMPTY
				}
				y += 2
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

func StarTwo(input []string, visualization bool) {
	var (
		start       = time.Now()
		wh, actions = furtherBigParse(input)
	)

	// Run all the Actions
	if visualization {
		wh.visualize("START")
	}
	for _, action := range actions {
		wh.moveBigBoxes(action)
		if visualization {
			wh.visualize(action)
		}
	}

	end := time.Now().Sub(start)
	fmt.Printf("Star Two: %d in %fs\n", wh.g.gpsScore(), end.Seconds())
}
