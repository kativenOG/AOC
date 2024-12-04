package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Regexp for both stars
var mainOne = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
var mainTwo = regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don\\'t\\(\\)")
var numberR = regexp.MustCompile("\\d+")

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

func starOne(input []string) {
	var result int
	for _, line := range input {
		matches := mainOne.FindAllString(line, -1)
		for _, match := range matches {
			numbers := numberR.FindAllString(match, -1)
			if len(numbers) != 2 {
				dieOnError(fmt.Errorf("too many operands in match %s", match))
			}
			n1, err := strconv.Atoi(numbers[0])
			dieOnError(err)
			n2, err := strconv.Atoi(numbers[1])
			dieOnError(err)
			result += n1 * n2
		}
	}
	fmt.Printf("Star One: %d\n", result)
}

func starTwo(input []string) {
	var result int
	var valid bool = true
	for _, line := range input {
		matches := mainTwo.FindAllString(line, -1)
	matchLoop:
		for _, match := range matches {
			switch match {
			case "do()":
				valid = true
				continue matchLoop
			case "don't()":
				valid = false
				continue matchLoop
			}

			if !valid {
				continue matchLoop
			}

			numbers := numberR.FindAllString(match, -1)
			if len(numbers) != 2 {
				dieOnError(fmt.Errorf("too many operands in match %s", match))
			}
			n1, err := strconv.Atoi(numbers[0])
			dieOnError(err)
			n2, err := strconv.Atoi(numbers[1])
			dieOnError(err)
			result += n1 * n2
		}
	}
	fmt.Printf("Star Two: %d\n", result)
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.Parse()
	input := parseInputFIie(filename)

	starOne(input)
	starTwo(input)
}
