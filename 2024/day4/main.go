package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"

	"github.com/AOC/2024/utils"

	"github.com/samber/lo"
)

func naiveParse(input []string, rs []*regexp.Regexp, shouldDebug bool) (res int) {
	for _, line := range input {
		for _, r := range rs {
			matches := r.FindAllString(line, -1)
			utils.DebugPrintf(shouldDebug, "%s %d\n", line, len(matches))
			res += len(matches)
		}
	}
	return
}

type coordinate struct {
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d-%d)", c.x, c.y)
}

func sortCoordinatesArray(coors *[]coordinate) {
	slices.SortFunc(*coors, func(a, b coordinate) int {
		if a.x == b.x {
			return 0
		} else if a.x > b.x {
			return 1
		}
		return -1
	})
}

type direction int

const (
	HORIZZONTAL = iota
	VERTICAL
	PRIMARY_DIAGONAL
	SECONDARY_DIAGONAL
)

func (d direction) String() (res string) {
	switch d {
	case HORIZZONTAL:
		res = "Horizzontal"
	case VERTICAL:
		res = "Vertical"
	case PRIMARY_DIAGONAL:
		res = "PrimaryDiag"
	case SECONDARY_DIAGONAL:
		res = "SecondaryDiag"
	default:
		panic("wtf, it's an enum golang !!!")
	}

	return
}

func (h direction) isAnX(other direction) (res bool) {
	if (h == HORIZZONTAL && other == VERTICAL) ||
		(h == VERTICAL && other == HORIZZONTAL) ||
		(h == PRIMARY_DIAGONAL && other == SECONDARY_DIAGONAL) ||
		(h == SECONDARY_DIAGONAL && other == PRIMARY_DIAGONAL) {
		res = true
	}

	return
}

func naiveParseIndeces(input []string, rs []*regexp.Regexp, shouldDebug bool) (res []coordinate) {
	var matches [][]int
	for i, line := range input {
		for _, r := range rs {
			matches = r.FindAllStringIndex(line, -1)
			utils.DebugPrintf(shouldDebug, "%s %d\n", line, len(matches))
			if len(matches) == 0 {
				continue
			}
			res = append(res, lo.Map(matches, func(match []int, _ int) coordinate {
				centerIndex := int(math.Ceil(float64((match[0] + match[1]) / 2)))
				return coordinate{i, centerIndex}
			})...)
		}
	}

	return
}

func parseCoordinates(matrix [][]string, rs []*regexp.Regexp, indeces [][]coordinate, shouldDebug bool) (res []coordinate) {
	var matches [][]int
	var input []string = lo.Map(indeces, func(coords []coordinate, _ int) string {
		val := ""
		for _, coor := range coords {
			val += matrix[coor.x][coor.y]
		}
		return val
	})
	for i, line := range input {
		for _, r := range rs {
			matches = r.FindAllStringIndex(line, -1)
			utils.DebugPrintf(shouldDebug, "%s %d\n", line, len(matches))
			if len(matches) == 0 {
				continue
			}
			res = append(res, lo.Map(matches, func(match []int, _ int) coordinate {
				centerIndex := int(math.Ceil(float64((match[0] + match[1]) / 2)))
				return indeces[i][centerIndex]
			})...)
		}
	}

	return
}

func realMatrix(input []string) (realMatrix [][]string) {
	realMatrix = make([][]string, 0, len(input))
	for _, line := range input {
		runeLine := []rune(line)
		parsedLine := lo.Map(runeLine, func(val rune, _ int) string {
			return string(val)
		})
		realMatrix = append(realMatrix, parsedLine)
	}

	return
}

func primaryDiagonals(matrix [][]string, h, w int) [][]string {
	var diff int
	diags := make(map[int][]string)
	for i := range lo.Range(h) {
		for j := range lo.Range(w) {
			diff = i - j
			if _, ok := diags[diff]; !ok {
				diags[diff] = []string{}
			}
			diags[diff] = append(diags[diff], matrix[i][j])
		}
	}

	return lo.Values(diags)
}

func indexPrimaryDiagonals(matrix [][]string, h, w int) [][]coordinate {
	var diff int
	diags := make(map[int][]coordinate)
	for i := range lo.Range(h) {
		for j := range lo.Range(w) {
			diff = i - j
			if _, ok := diags[diff]; !ok {
				diags[diff] = []coordinate{}
			}
			diags[diff] = append(diags[diff], coordinate{i, j})
		}
	}

	return lo.Values(diags)
}

func secondaryDiagonals(matrix [][]string, h, w int) [][]string {
	var summ int
	diags := make(map[int][]string)
	for i := range lo.Range(h) {
		for j := range lo.Range(w) {
			summ = i + j
			if _, ok := diags[summ]; !ok {
				diags[summ] = []string{}
			}
			diags[summ] = append(diags[summ], matrix[i][j])
		}
	}

	return lo.Values(diags)
}

func indexSecondaryDiagonals(matrix [][]string, h, w int) [][]coordinate {
	var summ int
	diags := make(map[int][]coordinate)
	for i := range lo.Range(h) {
		for j := range lo.Range(w) {
			summ = i + j
			if _, ok := diags[summ]; !ok {
				diags[summ] = []coordinate{}
			}
			diags[summ] = append(diags[summ], coordinate{i, j})
		}
	}

	return lo.Values(diags)
}

func parseStringList(stringList []string, rs []*regexp.Regexp) (res int) {
	var parsedList string
	for _, val := range stringList {
		parsedList += val
	}

	for _, r := range rs {
		res += len(r.FindAllString(parsedList, -1))
	}

	return
}

func crossWord(input []string, target string, shouldDebug bool) (res int) {
	directR := regexp.MustCompile(target)
	inverseR := regexp.MustCompile(utils.InvertString(target))
	rs := []*regexp.Regexp{
		directR,
		inverseR,
	}

	// Check for horizzontal matches
	utils.DebugPrintf(shouldDebug, "ORIZZONTAL:\n")
	res += naiveParse(input, rs, shouldDebug)
	utils.DebugPrintf(shouldDebug, "\n")

	// Create a vertical version
	n_colums := len(input[0])
	verticalAppo := make([][]rune, n_colums)
	for _, line := range input {
		runeLine := []rune(line)
		for i, val := range runeLine {
			verticalAppo[i] = append(verticalAppo[i], val)
		}
	}
	// Cast back the input to an array of Strings
	verticalInput := lo.Map(verticalAppo, func(column []rune, _ int) string {
		return string(column)
	})
	// Finally check for vertical matches
	utils.DebugPrintf(shouldDebug, "VERTICAL:\n")
	res += naiveParse(verticalInput, rs, shouldDebug)
	utils.DebugPrintf(shouldDebug, "\n")

	// Check for diagonal matches on main/cross diagonal
	var (
		height = len(input)
		width  = len(input[0])
		matrix = realMatrix(input)
	)
	diags := primaryDiagonals(matrix, height, width)
	diags = append(diags, secondaryDiagonals(matrix, height, width)...)
	res += naiveParse(lo.Map(diags, func(letters []string, _ int) string {
		res := ""
		for _, val := range letters {
			res += val
		}
		return res
	}), rs, shouldDebug)

	return
}

func xCrossWord(input []string, target string, shouldDebug bool) (res int) {
	directR := regexp.MustCompile(target)
	inverseR := regexp.MustCompile(utils.InvertString(target))
	rs := []*regexp.Regexp{
		directR,
		inverseR,
	}

	// Horizzontal
	resMap := make(map[coordinate][]direction)
	lo.ForEach(naiveParseIndeces(input, rs, shouldDebug), func(coor coordinate, _ int) {
		resMap[coor] = append(resMap[coor], HORIZZONTAL)
	})

	// Vertical
	// First create a vertical version
	n_colums := len(input[0])
	verticalAppo := make([][]rune, n_colums)
	for _, line := range input {
		runeLine := []rune(line)
		for i, val := range runeLine {
			verticalAppo[i] = append(verticalAppo[i], val)
		}
	}
	lo.ForEach(naiveParseIndeces(lo.Map(verticalAppo, func(column []rune, _ int) string {
		return string(column)
	}), rs, shouldDebug), func(coor coordinate, _ int) {
		resMap[coor] = append(resMap[coor], VERTICAL)
	})

	// Diagonals
	matrix := realMatrix(input)
	height, width := len(matrix), len(matrix[0])

	// Primary
	coords := indexPrimaryDiagonals(matrix, height, width)
	hits := parseCoordinates(matrix, rs, coords, shouldDebug)
	sortCoordinatesArray(&hits)
	lo.ForEach(hits, func(coor coordinate, _ int) {
		resMap[coor] = append(resMap[coor], PRIMARY_DIAGONAL)
	})
	utils.DebugPrintf(shouldDebug, "Primary Diagonal Matches Second Star %v\n", hits)

	// Secondary
	coords = indexSecondaryDiagonals(matrix, height, width)
	hits = parseCoordinates(matrix, rs, coords, shouldDebug)
	sortCoordinatesArray(&hits)
	lo.ForEach(hits, func(coor coordinate, _ int) {
		resMap[coor] = append(resMap[coor], SECONDARY_DIAGONAL)
	})

	utils.DebugPrintf(shouldDebug, "Secondary Diagonal Matches Second Star %v\n", hits)

	// Return sum of middle coordinates with more than one match
	utils.DebugPrintf(shouldDebug, "\nFinal Result Start Two:\n")

	for coor, count := range resMap {
		utils.DebugPrintf(shouldDebug, "%v %d\n", coor, count)
	}

	return lo.Sum(lo.FilterMap(lo.Values(resMap), func(dirs []direction, _ int) (val int, res bool) {
	mainLoop:
		for i, a := range dirs {
			for _, b := range dirs[i+1:] {
				if ok := a.isAnX(b); ok {
					val, res = 1, true
					break mainLoop
				}
			}
		}

		return
	}))

}

func starOne(input []string, shouldDebug bool) {
	target := "XMAS"
	fmt.Printf("Star One: %d\n", crossWord(input, target, shouldDebug))
}

func starTwo(input []string, shouldDebug bool) {
	target := "MAS"
	fmt.Printf("Star Two: %d\n", xCrossWord(input, target, shouldDebug))
}

func main() {
	filename, debug := utils.ParseFlags()
	input := utils.ParseInputFile(filename)
	starOne(input, debug)
	starTwo(input, debug)
}
