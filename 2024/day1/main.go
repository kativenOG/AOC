package main

import (
	"flag"
	"fmt"
	"github.com/samber/lo"
	"os"
	"slices"
	"strconv"
	"strings"
)

func dieOnErr(err error) {
	if err != nil {
		panic(fmt.Errorf("fatal error: %w", err))
	}
	return
}

// readinput returns the parsed file content in an array of string,
// in which each entry reppresents a line.
func readinput(filename string) []string {
	f, err := os.ReadFile(filename)
	dieOnErr(err)

	res := strings.Split(string(f), "\n")
	if len(res[len(res)-1]) == 0 {
		res = res[:len(res)-1]
	}

	return res
}

// parseIntoIntegerLists parses the input in two integer lists and sortes them DESC
func parseIntoIntegerLists(input []string) (listOne []int, listTwo []int) {
	for i, line := range input {
		contents := strings.Split(line, "   ")
		if contLen := len(contents); contLen != 2 {
			dieOnErr(fmt.Errorf("input in line %d has %d ids instead of 2", i, contLen))
		}

		numberOne, err := strconv.Atoi(contents[0])
		dieOnErr(err)
		listOne = append(listOne, numberOne)

		numberTwo, err := strconv.Atoi(strings.TrimSpace(contents[1]))
		dieOnErr(err)
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
	var filename string
	flag.StringVar(&filename, "filename", "input.txt", "the input file name")
	flag.Parse()
	input := readinput(filename)

	starOne(input)
	starTwo(input)

}
