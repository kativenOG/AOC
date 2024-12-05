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
	for i, j := 0, len(runes); i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func naiveParse(input []string, rs []*regexp.Regexp) (res int) {
	for _, line := range input {
		for _, r := range rs {
			matches := r.FindAllString(line, -1)
			res += len(matches)
		}
	}
	return
}

func crossWord(input []string, target string) (res int) {
	directR := regexp.MustCompile(target)
	inverseR := regexp.MustCompile(invertString(target))
	rs := []*regexp.Regexp{
		directR,
		inverseR,
	}

	// Check for horizzontal matches
	res += naiveParse(input, rs)

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
	res += naiveParse(verticalInput, rs)

	// Check for diagonal matches
	var diagonalInput []string
	length, width := len(input), len(input)
	res += naiveParse(diagonalInput, rs)

	return
}

func starOne(input []string) {
	target := "XMAS"
	fmt.Println("Star One: %d", crossWord(input, target))
}

func starTwo(input []string) {
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.Parse()
	input := parseInputFIie(filename)
	fmt.Println(len(input))
}
