package main

import (
	"flag"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func dieOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// DIRECTIONS:
const (
	UNDEFINED = iota
	INCREASING
	DECREASING
)

func intAbs(a, b int) (res int) {
	res = a - b
	if res < 0 {
		res = -res
	}
	return
}

func parseDir(previous, current, currentDirection int) int {
	if (currentDirection == UNDEFINED) && (previous == current) {
		return UNDEFINED
	} else if previous > current {
		return DECREASING
	}
	return INCREASING
}

func parseReport(levels []int) bool {
	var direction int = UNDEFINED
	previous := levels[0]
	for _, level := range levels[1:] {
		parsedDirection := parseDir(previous, level, direction)
		if direction == 0 {
			direction = parsedDirection
		}
		if diff := intAbs(previous, level); (diff < 1 || diff > 3) ||
			(parsedDirection != direction) {
			return false
		}
		previous = level
	}
	return true
}

// parseReportDampened if the first check fails it checks iteritavely for a functioning levels set
func parseReportDampened(levels []int) (res bool) {
	res = parseReport(levels)
	if res {
		return
	}

	for index, _ := range levels {
		dampenedLevels := make([]int, len(levels))
		copy(dampenedLevels, levels)
		dampenedLevels = append(dampenedLevels[:index], dampenedLevels[index+1:]...)

		if res = parseReport(dampenedLevels); res {
			break
		}
	}

	return
}

func parseInputFIie(filename string) (inputList []string) {
	content, err := os.ReadFile(filename)
	dieOnError(err)
	inputList = strings.Split(string(content), "\n")
	if len(inputList[len(inputList)-1]) == 0 {
		inputList = inputList[:len(inputList)-1]
	}
	return
}

func starOne(inputs []string, dampened bool) {
	var validReports int
	for _, line := range inputs {
		stringLevels := strings.Split(line, " ")
		levels := lo.Map(stringLevels, func(val string, _ int) int {
			res, err := strconv.Atoi(strings.TrimSpace(val))
			dieOnError(err)
			return res
		})
		var ok bool
		if dampened {
			ok = parseReportDampened(levels)
		} else {
			ok = parseReport(levels)
		}
		if ok {
			validReports += 1
		}
	}

	if dampened {
		fmt.Printf("Star Two: %d\n", validReports)
		return
	}
	fmt.Printf("Star One: %d\n", validReports)
}

func starTwo(inputs []string) {
	starOne(inputs, true)
}

func main() {

	var filename string
	flag.StringVar(&filename, "filename", "inputs.txt", "the input file")
	flag.Parse()
	input := parseInputFIie(filename)

	starOne(input, false)
	starTwo(input)

}
