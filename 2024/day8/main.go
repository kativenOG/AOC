package main

import (
	"fmt"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

type identifier string

type coordinate struct {
	x, y int
}

func (coor coordinate) String() string {
	return fmt.Sprintf("(x:%d y:%d)", coor.x, coor.y)
}

func (coor coordinate) sum(other coordinate) coordinate {
	return coordinate{
		x: coor.x + other.x,
		y: coor.y + other.y,
	}
}
func (coor coordinate) outOfBounds(maxX, maxY int) bool {
	if (coor.x >= 0 && coor.x <= maxX) && (coor.y >= 0 && coor.y <= maxY) {
		return false
	}
	return true
}

func (coor coordinate) dist(other coordinate) coordinate {
	return coordinate{
		x: coor.x - other.x,
		y: coor.y - other.y,
	}
}

func (coor coordinate) multiply(scalar int) coordinate {
	return coordinate{
		x: coor.x * scalar,
		y: coor.y * scalar,
	}
}

type antenna struct {
	id   identifier
	coor coordinate
}

func (a antenna) String() string {
	return fmt.Sprintf("%s: %v", a.id, a.coor)
}

type grid struct {
	tiles    map[coordinate]identifier
	antennas map[identifier][]*antenna
}

func printAntinodes(antinodes map[coordinate]int, maxX, maxY int) {
	for y := range lo.Range(maxY + 1) {
		for x := range lo.Range(maxX + 1) {
			if _, ok := antinodes[coordinate{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("\n")
	}
	fmt.Println("\n")
}

func parseInput(input []string) (g grid, maxX, maxY int) {
	g.tiles = make(map[coordinate]identifier)
	g.antennas = make(map[identifier][]*antenna)
	maxX = len(input[0]) - 1
	maxY = len(input) - 1

	for y, line := range input {
		linee := []rune(line)

		for x, val := range linee {
			id := identifier(val)
			coor := coordinate{x, y}
			g.tiles[coor] = id
			if id != "." {
				newAntenna := &antenna{
					id:   id,
					coor: coor,
				}

				if prev, ok := g.antennas[id]; !ok {
					g.antennas[id] = []*antenna{newAntenna}
				} else {
					g.antennas[id] = append(prev, newAntenna)
				}
			}
		}
	}

	return
}

func StarOne(input []string) {
	start := time.Now()

	antinodes := make(map[coordinate]struct{})
	g, maxX, maxY := parseInput(input)

	var antiNodePos coordinate
	for _, antennas := range g.antennas {
		for index, first := range antennas {

			for _, other := range antennas[index+1:] {
				// Do it once for the first node of the pair
				antiNodePos = first.coor.sum(first.coor.dist(other.coor))
				if !antiNodePos.outOfBounds(maxX, maxY) {
					antinodes[antiNodePos] = struct{}{}
				}

				// Then do the mirror position for the second node
				antiNodePos = other.coor.sum(other.coor.dist(first.coor))
				if !antiNodePos.outOfBounds(maxX, maxY) {
					antinodes[antiNodePos] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("Star One: %d, in %fs\n", len(lo.Keys(antinodes)), time.Now().Sub(start).Seconds())
}

func StarTwo(input []string) {
	start := time.Now()

	antinodes := make(map[coordinate]struct{})
	g, maxX, maxY := parseInput(input)

	var diff coordinate
	for _, antennas := range g.antennas {
		for index, first := range antennas {
			for _, other := range antennas[index+1:] {
				// Do it once for the first node of the pair
				diff = first.coor.dist(other.coor)
				antiNodePosFirst := first.coor.sum(diff)
				for !antiNodePosFirst.outOfBounds(maxX, maxY) {
					antinodes[antiNodePosFirst] = struct{}{}
					antiNodePosFirst = antiNodePosFirst.sum(diff)
				}

				// Then do the mirror position for the second node
				diff = other.coor.dist(first.coor)
				antiNodePosSecond := other.coor.sum(diff)
				for !antiNodePosSecond.outOfBounds(maxX, maxY) {
					antinodes[antiNodePosSecond] = struct{}{}
					antiNodePosSecond = antiNodePosSecond.sum(diff)
				}
			}
		}
	}

	// Fill it with the normal antennas
	lo.ForEach(lo.Values(g.antennas), func(arr []*antenna, _ int) {
		for _, a := range arr {
			antinodes[a.coor] = struct{}{}
		}
	})

	fmt.Printf("Star Two: %d, in %fs\n", len(lo.Keys(antinodes)), time.Now().Sub(start).Seconds())
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	StarOne(input)
	StarTwo(input)
}
