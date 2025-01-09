package main

import (
	"fmt"
	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"time"
)

type coordinate struct {
	x, y int
}

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func (coor coordinate) ManhattanDistance(otherCoor coordinate) int {
	return absInt(coor.x-otherCoor.x) + absInt(coor.y-otherCoor.y)
}

type clawMachine struct {
	buttonA, buttonB, prize coordinate
}

func returnCoord(match []string) coordinate {
	x, err := strconv.Atoi(match[1])
	utils.DieOnError(err)
	y, err := strconv.Atoi(match[2])
	utils.DieOnError(err)
	return coordinate{x, y}
}

func parseClawMachines(input []string) (clawMachines []clawMachine) {
	var (
		matches                    [][]string
		buttonAs, buttonBs, prizes []coordinate
		rButtonA                   = regexp.MustCompile("Button A: X\\+(\\d+), Y\\+(\\d+)")
		rButtonB                   = regexp.MustCompile("Button B: X\\+(\\d+), Y\\+(\\d+)")
		rPrize                     = regexp.MustCompile("Prize: X\\=(\\d+), Y\\=(\\d+)")
	)

	for _, line := range input {
		if matches = rButtonA.FindAllStringSubmatch(line, -1); len(matches) > 0 {
			buttonAs = append(buttonAs, returnCoord(matches[0]))
		} else if matches = rButtonB.FindAllStringSubmatch(line, -1); len(matches) > 0 {
			buttonBs = append(buttonBs, returnCoord(matches[0]))
		} else if matches = rPrize.FindAllStringSubmatch(line, -1); len(matches) > 0 {
			coord := returnCoord(matches[0])
			prizes = append(prizes, coord)
		}
	}

	return lo.Map(buttonAs, func(coor coordinate, index int) clawMachine {
		return clawMachine{
			coor,
			buttonBs[index],
			prizes[index],
		}
	})
}

func starOne(input []string) {
	var (
		res   int
		start = time.Now()

		clawMachines             = parseClawMachines(input)
		buttonA, buttonB, prize  coordinate
		currentValX, currentValY int
		innerX, innerY           int
	)

clawMachineLoop:
	for _, clawMachine := range clawMachines {
		buttonA, buttonB, prize = clawMachine.buttonA, clawMachine.buttonB, clawMachine.prize
	innerLoop:
		for aMultiplier := 0; aMultiplier*buttonA.x < prize.x && aMultiplier*buttonA.y < prize.y; aMultiplier++ {
			currentValX, currentValY = aMultiplier*buttonA.x, aMultiplier*buttonA.y

			for bMultiplier := 0; true; bMultiplier++ {
				innerX = (currentValX + bMultiplier*buttonB.x)
				innerY = (currentValY + bMultiplier*buttonB.y)

				if (innerX == prize.x) && (innerY == prize.y) {
					res += (aMultiplier*3 + bMultiplier)
					continue clawMachineLoop
				}

				if (innerX > prize.x) || (innerY > prize.y) {
					continue innerLoop
				}
			}
		}
	}
	end := time.Now().Sub(start)

	fmt.Printf("Star One: %d in %fs\n", res, end.Seconds())
}

func starTwo(input []string) {
	var (
		start                   = time.Now()
		res                     int
		buttonA, buttonB, prize coordinate
	)
	clawMachines := parseClawMachines(input)
	for _, cm := range clawMachines {
		buttonA, buttonB, prize = cm.buttonA, cm.buttonB, cm.prize
		prize.x += 10000000000000
		prize.y += 10000000000000

		// First get the B
		nb := ((prize.y * buttonA.x) - (prize.x * buttonA.y))
		db := ((buttonA.x * buttonB.y) - (buttonB.x * buttonA.y))

		// Check if are divisible
		if nb%db == 0 {
			b := nb / db
			if na, da := (prize.x - (b * buttonB.x)), buttonA.x; na%da == 0 {
				a := (na / da)
				res += a*3 + b
			}
		}
	}

	end := time.Now().Sub(start).Seconds()
	fmt.Printf("Star Two: %d in %fs\n", res, end)

}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)
}
