package main

import (
	"flag"
	"fmt"
	"github.com/AOC/2024/utils"
	"regexp"
	"strconv"
)

// Regexp for both stars
var mainOne = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
var mainTwo = regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don\\'t\\(\\)")
var numberR = regexp.MustCompile("\\d+")

func starOne(input []string) {
	var result int
	for _, line := range input {
		matches := mainOne.FindAllString(line, -1)
		for _, match := range matches {
			numbers := numberR.FindAllString(match, -1)
			if len(numbers) != 2 {
				utils.DieOnError(fmt.Errorf("too many operands in match %s", match))
			}
			n1, err := strconv.Atoi(numbers[0])
			utils.DieOnError(err)
			n2, err := strconv.Atoi(numbers[1])
			utils.DieOnError(err)
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
				utils.DieOnError(fmt.Errorf("too many operands in match %s", match))
			}
			n1, err := strconv.Atoi(numbers[0])
			utils.DieOnError(err)
			n2, err := strconv.Atoi(numbers[1])
			utils.DieOnError(err)
			result += n1 * n2
		}
	}
	fmt.Printf("Star Two: %d\n", result)
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)
}
