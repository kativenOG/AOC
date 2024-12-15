package main

import (
	"github.com/AOC/2024/day15/stars"
	"github.com/AOC/2024/utils"
)

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	stars.StarOne(input, false)
	stars.StarTwo(input, true)
}
