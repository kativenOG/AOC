package main

import (
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/AOC/2024/utils"
)

func getRulesAndInputs(input []string) (rules []string, problemInput []string) {
	for i, line := range input {
		if len(line) == 0 {
			return input[:i], input[i+1:]
		}
	}
	panic("no empty line for separating rules and inputs")
}

func parseList(stringList string) (res []int) {
	values := strings.Split(stringList, ",")
	for _, value := range values {
		parsedValue, err := strconv.Atoi(strings.TrimSpace(value))
		utils.DieOnError(err)
		res = append(res, parsedValue)
	}
	return
}

func findCorrectInputsMiddleValues(input []string, correctlyOrdered bool) (result int) {
	// Map rules in a more convenient way
	rules, problemInputs := getRulesAndInputs(input)
	mappedRules := map[int][]int{}
	for _, rule := range rules {
		values := strings.Split(rule, "|")
		if len(values) != 2 {
			panic("more than 2 values in a rule")
		}
		leftValue, err := strconv.Atoi(strings.TrimSpace(values[0]))
		utils.DieOnError(err)
		if _, ok := mappedRules[leftValue]; !ok {
			mappedRules[leftValue] = []int{}
		}
		rightValue, err := strconv.Atoi(strings.TrimSpace(values[1]))
		mappedRules[leftValue] = append(mappedRules[leftValue], rightValue)
	}

inputLoop:
	for _, line := range problemInputs {
		inputList := parseList(line)
		// Map the input list to a direct map based on the index
		mappedDirectIndex := map[int]int{} // value -> index
		for i, value := range inputList {
			mappedDirectIndex[value] = i
		}
		// Check that the input lists respects all the relative rules
		for _, value := range inputList {
			// If there are rules relative to this number
			if relativeRules, ok := mappedRules[value]; ok {
				targetPos := mappedDirectIndex[value]
				for _, rule := range relativeRules {
					// If the rule is before the input the input automatically becomes invalid
					if rulePos, numberExists := mappedDirectIndex[rule]; numberExists && targetPos > rulePos {
						if correctlyOrdered {
							slices.SortFunc(inputList, func(a, b int) int {
								if rulesA, ok := mappedRules[a]; ok && slices.Contains(rulesA, b) {
									return -1
								} else if rulesB, ok := mappedRules[b]; ok && slices.Contains(rulesB, a) {
									return 1
								}
								return 0
							})

							middleIndex := int(math.Floor(float64(len(inputList)) / 2.0))
							result += inputList[middleIndex]
							continue inputLoop
						}
					}
				}
			}
		}

		if !correctlyOrdered {
			middleIndex := int(math.Floor(float64(len(inputList)) / 2.0))
			result += inputList[middleIndex]
		}
	}
	return result
}

func starOne(input []string) {
	fmt.Printf("Star One: %d\n", findCorrectInputsMiddleValues(input, false))
}
func starTwo(input []string) {
	fmt.Printf("Star Two: %d\n", findCorrectInputsMiddleValues(input, true))
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.Parse()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)
}
