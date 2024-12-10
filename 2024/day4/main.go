package main

import (
	"flag"
	"fmt"
	"github.com/AOC/2024/utils"
	"regexp"

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

	// Check for diagonal matches on main diagonal
	var (
		height       = len(input)
		width        = len(input[0])
		inputMatrix  = realMatrix(input)
		matrixTarget string
	)

	utils.DebugPrintf(shouldDebug, "MAIN DIAGONAL:\n")
	for i := 0; i < height; i++ {
		plainDiagonal := []string{}
		if i == 0 {
			plainDiagonal = append(plainDiagonal, inputMatrix[i][0])
		} else {
			// diagonal := []MatrixEntry{}
			for j := 0; j < i+1; j++ {
				matrixTarget = inputMatrix[i-j][j]
				plainDiagonal = append(plainDiagonal, matrixTarget)
			}
		}
		res += parseStringList(plainDiagonal, rs)
		utils.DebugPrintf(shouldDebug, "%v %d %s", plainDiagonal, parseStringList(plainDiagonal, rs), "\n")
	}
	utils.DebugPrintf(shouldDebug, "\n")

	// Check for diagonal matches on other diagonal
	utils.DebugPrintf(shouldDebug, "CROSS DIAGONAL:\n")
	for i := 0; i < height; i++ {
		plainDiagonal := []string{}
		if i == 0 {
			plainDiagonal = append(plainDiagonal, inputMatrix[i][width-1])
		} else {
			for k, j := 0, width-1; k < i+1; k, j = k+1, j-1 {
				matrixTarget = inputMatrix[i-k][j]
				plainDiagonal = append(plainDiagonal, matrixTarget)
			}
		}
		res += parseStringList(plainDiagonal, rs)

		utils.DebugPrintf(shouldDebug, "%v %d %s", plainDiagonal, parseStringList(plainDiagonal, rs), "\n")
	}
	utils.DebugPrintf(shouldDebug, "\n")

	return
}

func starOne(input []string, shouldDebug bool) {
	target := "XMAS"
	fmt.Printf("Star One: %d\n", crossWord(input, target, shouldDebug))
}

func starTwo(input []string, shouldDebug bool) {
}

func main() {
	filename, debug := utils.ParseFlags()
	input := utils.ParseInputFile(filename)
	starOne(input, debug)
	// starTwo(input, debug)
}
