package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

type coordinate struct {
	x, y int
}

type velocity struct {
	x, y int
}

type robot struct {
	coor coordinate
	vel  velocity
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func (r *robot) moveN(n, maxX, maxY int) coordinate {
	r.coor = coordinate{
		x: (r.coor.x + (r.vel.x * n)) % maxX,
		y: (r.coor.y + (r.vel.y * n)) % maxY,
	}
	if r.coor.x < 0 {
		r.coor.x = maxX - intAbs(r.coor.x)
	}
	if r.coor.y < 0 {
		r.coor.y = maxY - intAbs(r.coor.y)
	}

	return r.coor
}

func returnCoord(match []string) coordinate {
	x, err := strconv.Atoi(match[0])
	utils.DieOnError(err)
	y, err := strconv.Atoi(match[1])
	utils.DieOnError(err)
	return coordinate{x, y}
}

func returnVel(match []string) velocity {
	x, err := strconv.Atoi(match[0])
	utils.DieOnError(err)
	y, err := strconv.Atoi(match[1])
	utils.DieOnError(err)
	return velocity{x, y}
}

func parseRobots(input []string) (robots []robot) {
	var (
		matches [][]string
		rRobot  = regexp.MustCompile("p=(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")
	)

	for _, line := range input {
		if matches = rRobot.FindAllStringSubmatch(line, -1); len(matches) > 0 {
			targetMatch := matches[0][1:]
			robots = append(robots, robot{
				returnCoord(targetMatch[:2]),
				returnVel(targetMatch[2:]),
			},
			)
		}
	}

	return
}

func printRobots(robotPos map[coordinate]struct{}, maxX, maxY int, file *os.File) {
	var err error
	var fileContent string
	for y := range maxY {
		// Line
		for x := range maxX {
			tileVal := " "
			if _, ok := robotPos[coordinate{x, y}]; ok {
				tileVal = "\u25A0"
			}
			if file != nil {
				fileContent += tileVal
				_, err = file.Write([]byte(tileVal))
				utils.DieOnError(err)
			} else {
				fmt.Printf(tileVal)
			}
		}

		// Return
		if file != nil {
			_, err = file.Write([]byte("\n"))
		} else {
			fmt.Printf("\n")
		}
	}

	// Write to file
	_, err = file.Write([]byte("\n"))
}

func cycleTroughVisualizations(robots []robot, maxX, maxY int) (res int) {
	for true {
		robotPos := make(map[coordinate]struct{})
		newRobots := make([]robot, 0, len(robots))
		for _, r := range robots {
			r.coor = r.moveN(1, maxX, maxY)
			robotPos[r.coor] = struct{}{}
			newRobots = append(newRobots, r)
		}
		robots = newRobots
		printRobots(robotPos, maxX, maxY, nil)
		res += 1

		fmt.Printf("Current: %d", res)
		time.Sleep(500 * time.Millisecond)
		utils.CleanTerminal(maxY)
	}

	return res
}

func visualize(robots []robot, maxX, maxY int, seconds int) {
	robotPos := make(map[coordinate]struct{})
	newRobots := make([]robot, 0, len(robots))
	for _, r := range robots {
		r.coor = r.moveN(seconds, maxX, maxY)
		robotPos[r.coor] = struct{}{}
		newRobots = append(newRobots, r)
	}
	robots = newRobots
	printRobots(robotPos, maxX, maxY, nil)

}

func hugeDumbFileGenerator(robots []robot, filename string, maxX, maxY int) {
	var res int
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	utils.DieOnError(err)
	for range lo.Range(10000) {
		robotPos := make(map[coordinate]struct{})
		newRobots := make([]robot, 0, len(robots))
		for _, r := range robots {
			r.coor = r.moveN(1, maxX, maxY)
			robotPos[r.coor] = struct{}{}
			newRobots = append(newRobots, r)
		}
		robots = newRobots
		res += 1

		// Write info to file
		printRobots(robotPos, maxX, maxY, file)
		file.Write([]byte(fmt.Sprintf("Current: %d\n", res)))
	}
	utils.DieOnError(file.Close())
}

func findQuadrantCount(elements map[coordinate]int, halfX, halfY int, firstHalfX, firstHalfY bool) int {
	return lo.Sum(lo.FilterMap(lo.Keys(elements), func(coor coordinate, _ int) (int, bool) {
		res := true
		if (firstHalfX && coor.x >= halfX) ||
			(!firstHalfX && coor.x <= halfX) ||
			(firstHalfY && coor.y >= halfY) ||
			(!firstHalfY && coor.y <= halfY) {
			res = false
		}
		return elements[coor], res
	}))
}

func starOne(input []string, nSeconds, maxX, maxY int) {
	var (
		res    int = 1
		start      = time.Now()
		robots     = parseRobots(input)
	)

	finishGrid := make(map[coordinate]int)
	for _, bot := range robots {
		finishGrid[bot.moveN(nSeconds, maxX, maxY)] += 1
	}

	halfX, halfY := (maxX-1)/2, (maxY-1)/2
	for _, firstHalfX := range []bool{true, false} {
		for _, firstHalfY := range []bool{true, false} {
			res *= findQuadrantCount(finishGrid, halfX, halfY, firstHalfX, firstHalfY)
		}
	}

	end := time.Now().Sub(start)
	fmt.Printf("Star One: %d in %fs\n", res, end.Seconds())
}

func starTwo(input []string, maxX, maxY int) {
	// hugeDumbFileGenerator(parseRobots(input), "tree.txt", maxX, maxY)
	// fmt.Printf("Star One: %d", cycleTroughVisualizations(parseRobots(input), maxX, maxY))
	visualize(parseRobots(input), maxX, maxY, 7051)
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input, 100, 101, 103)
	starTwo(input, 101, 103)
}
