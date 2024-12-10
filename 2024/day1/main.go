package main

import (
	"flag"
	"fmt"
	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
	"slices"
	"strconv"
	"strings"
)

// parseIntoIntegerLists parses the input in two integer lists and sortes them DESC
func parseIntoIntegerLists(input []string) (listOne []int, listTwo []int) {
	for i, line := range input {
		contents := strings.Split(line, "   ")
		if contLen := len(contents); contLen != 2 {
			utils.DieOnError(fmt.Errorf("input in line %d has %d ids instead of 2", i, contLen))
		}

		numberOne, err := strconv.Atoi(contents[0])
		utils.DieOnError(err)
		listOne = append(listOne, numberOne)

		numberTwo, err := strconv.Atoi(strings.TrimSpace(contents[1]))
		utils.DieOnError(err)
		listTwo = append(listTwo, numberTwo)
	}
	slices.Sort(listOne)
	slices.Sort(listTwo)

	return
}

func starOne(input []string) {
	listOne, listTwo := parseIntoIntegerLists(input)
	result := lo.Map(listOne, func(val int, index int) int {
		val -= listTwo[index]
		if val < 0 {
			return -val
		}
		return val
	})
	fmt.Printf("Star 1 Results: %d\n", lo.Sum(result))
}

func starTwo(input []string) {
	listOne, listTwo := parseIntoIntegerLists(input)

	// Make Set DS from the first list
	setOne := []int{}
	visited := make(map[int]struct{})
	lo.ForEach(listOne, func(val int, _ int) {
		if _, ok := visited[val]; !ok {
			setOne = append(setOne, val)
			visited[val] = struct{}{}
		}
	})

	// Make Counter DS from second list
	listTwoCounter := make(map[int]int)
	lo.ForEach(listTwo, func(val int, _ int) {
		if _, ok := listTwoCounter[val]; !ok {
			listTwoCounter[val] = 0
		}
		listTwoCounter[val] += 1
	})

	// Calculate the similarity index
	result := lo.Sum(
		lo.Map(setOne, func(val int, _ int) int {
			return val * listTwoCounter[val]
		}),
	)

	fmt.Printf("Star 2 Results: %d\n", result)
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)

}
