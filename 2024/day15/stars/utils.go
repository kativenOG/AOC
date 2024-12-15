package stars

import (
	"fmt"
	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

type coordinate struct {
	x, y int
}

func (coor coordinate) coordinateSum(otherCoordinate coordinate) coordinate {
	return coordinate{
		x: coor.x + otherCoordinate.x,
		y: coor.y + otherCoordinate.y,
	}
}

type grid map[coordinate]tileType

func (g grid) gpsScore() (res int) {
	for coor, tile := range g {
		if tile == BLOCK || tile == BIG_BLOCK_LEFT {
			res += 100*coor.x + coor.y
		}
	}
	return
}

type tileType int

const (
	WALL tileType = iota
	BLOCK
	EMPTY
	ROBOT
	BIG_BLOCK_LEFT
	BIG_BLOCK_RIGHT
)

func parseTile(input string) (tile tileType, r bool) {
	switch input {
	case "#":
		return WALL, false
	case "@":
		return ROBOT, true
	case "O":
		return BLOCK, false
	case "[":
		return BIG_BLOCK_LEFT, false
	case "]":
		return BIG_BLOCK_RIGHT, false
	case ".":
		return EMPTY, false
	case "\n":
	}
	panic(fmt.Sprintf("Wtf, this char shouldn't be here %s", input))
}

func reverseParseTile(tt tileType) string {
	switch tt {
	case WALL:
		return "#"
	case ROBOT:
		return "@"
	case BLOCK:
		return "O"
	case BIG_BLOCK_LEFT:
		return "["
	case BIG_BLOCK_RIGHT:
		return "]"
	case EMPTY:
		return "."
	}
	panic(fmt.Sprintf("Wtf, this char shouldn't be here %s", tt))
}

type warehouse struct {
	g     grid
	robot coordinate
}

func (wh warehouse) visualize(currentAction string) {
	maxX := lo.Max(lo.Map(lo.Keys(wh.g), func(coor coordinate, _ int) int {
		return coor.x
	}))
	maxY := lo.Max(lo.Map(lo.Keys(wh.g), func(coor coordinate, _ int) int {
		return coor.y
	}))
	utils.CleanTerminal(maxY + 1)
	for x := range lo.Range(maxX + 1) {
		for y := range lo.Range(maxY + 1) {
			fmt.Printf(reverseParseTile(wh.g[coordinate{x, y}]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Current action: %s", currentAction)
	var appo bool
	fmt.Scanln(&appo)

}
