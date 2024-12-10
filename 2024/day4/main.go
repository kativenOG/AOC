package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func dieOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func debugPrintf(shouldDebug bool, args ...interface{}) {
	if !shouldDebug {
		return
	}
	fmt.Printf(args[0].(string), args[1:]...)
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

func invertString(target string) string {
	runes := []rune(target)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func naiveParse(input []string, rs []*regexp.Regexp, shouldDebug bool) (res int) {
	for _, line := range input {
		for _, r := range rs {
			matches := r.FindAllString(line, -1)
			debugPrintf(shouldDebug, "%s %d\n", line, len(matches))
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
	inverseR := regexp.MustCompile(invertString(target))
	rs := []*regexp.Regexp{
		directR,
		inverseR,
	}

	// Check for horizzontal matches
	debugPrintf(shouldDebug, "ORIZZONTAL:\n")
	res += naiveParse(input, rs, shouldDebug)
	debugPrintf(shouldDebug, "\n")

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
	debugPrintf(shouldDebug, "VERTICAL:\n")
	res += naiveParse(verticalInput, rs, shouldDebug)
	debugPrintf(shouldDebug, "\n")

	// Check for diagonal matches on main diagonal
	var (
		height       = len(input)
		width        = len(input[0])
		inputMatrix  = realMatrix(input)
		matrixTarget string
	)

	debugPrintf(shouldDebug, "MAIN DIAGONAL:\n")
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
		debugPrintf(shouldDebug, "%v %d %s", plainDiagonal, parseStringList(plainDiagonal, rs), "\n")
	}
	debugPrintf(shouldDebug, "\n")

	// Check for diagonal matches on other diagonal
	debugPrintf(shouldDebug, "CROSS DIAGONAL:\n")
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

		debugPrintf(shouldDebug, "%v %d %s", plainDiagonal, parseStringList(plainDiagonal, rs), "\n")
	}
	debugPrintf(shouldDebug, "\n")

	return
}

func starOne(input []string, shouldDebug bool) {
	target := "XMAS"
	fmt.Printf("Star One: %d\n", crossWord(input, target, shouldDebug))
}

func starTwo(input []string) {
}

func main() {
	var filename string
	var debug bool
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.BoolVar(&debug, "debug", false, "verbose output on lines and matches")
	flag.Parse()
	input := parseInputFIie(filename)
	starOne(input, debug)
}
