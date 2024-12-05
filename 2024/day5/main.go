package main

import (
	"flag"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis/passes/deepequalerrors"
)

func dieOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInputFile(filename string) (inputList []string) {
	content, err := os.ReadFile(filename)
	dieOnError(err)
	inputList = strings.Split(string(content), "\n")
	if len(inputList[len(inputList)-1]) == 0 {
		inputList = inputList[:len(inputList)-1]
	}
	return
}

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
		dieOnError(err)
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
		dieOnError(err)
		if _, ok := mappedRules[leftValue]; !ok {
			mappedRules[leftValue] = []int{}
		}
		rightValue, err := strconv.Atoi(strings.TrimSpace(values[1]))
		mappedRules[leftValue] = append(mappedRules[leftValue], rightValue)
	}
	// recursively deepenedMap for second star
	deepMappedRules := make(map[int][]int, len(mappedRules))
	if correctlyOrdered {
		// first deep copy the initial map
		for key, value := range mappedRules {
			var newValue []int
			copy(newValue, value)
			deepMappedRules[key] = value
		}
		// Then keep expanding it
		var changed bool
		newValues := map[int]map[int]struct{}{}
		for true {
			for key, value := range mappedRules {
				if _, ok := newValues[key]; !ok {

				}
				for _, val := range value {

				}
			}
		}
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
						continue inputLoop
					}
				}
			}
		}

		if correctlyOrdered {
			slices.SortFunc(inputList, func(a, b int) int {
				if slices.Contains(deepMappedRules[a], b) {
					return 1
				} else if slices.Contains(deepMappedRules[b], a) {
					return -1
				}
				return 0
			})
		}
		middleIndex := int(math.Floor(float64(len(inputList)) / 2.0))
		result += inputList[middleIndex]
	}
	return result
}

func starOne(input []string) {
	fmt.Printf("Star One: %d\n", findCorrectInputsMiddleValues(input, false))
}
func starTwo(input []string) {
	fmt.Printf("Star One: %d\n", findCorrectInputsMiddleValues(input, true))
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.Parse()
	input := parseInputFile(filename)

	starOne(input)
	starTwo(input)
}
