package main

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

type fileSystem []int

func (fs fileSystem) checksum() (checksumResult int) {
	return lo.Sum(lo.Map(fs, func(val, index int) int {
		// Skip free space
		if val == -1 {
			return 0
		}
		// Index based checksum
		return index * val
	},
	),
	)
}

type freeSpace struct {
	size, index int
	used        bool
}

func (loc freeSpace) String() string {
	return fmt.Sprintf("Location[size:%d, index:%d]", loc.size, loc.index)
}

type file struct {
	id, size, index int
}

func (f file) String() string {
	return fmt.Sprintf("File[id:%d, size:%d, index:%d]", f.id, f.size, f.index)
}

func parseFileSystem(input []string) (fs fileSystem, freeSpaces []freeSpace, files []file) {
	var id int

	// Should be only one line
	for _, line := range input {
		runeLine := []rune(line)
		var (
			isFile  bool = true
			err     error
			parsedN int
		)
		for _, r := range runeLine {
			parsedN, err = strconv.Atoi(string(r))
			utils.DieOnError(err)

			start := len(fs)
			fs = append(fs, lo.Map(lo.Range(parsedN), func(_, _ int) int {
				if isFile {
					return id
				}
				return -1
			})...)

			if isFile {
				files = append(files, file{
					id:    id,
					size:  parsedN,
					index: start,
				})
				id += 1
			} else {
				freeSpaces = append(
					freeSpaces,
					freeSpace{
						parsedN,
						start,
						false,
					},
				)
			}

			isFile = !isFile
		}

	}

	slices.SortFunc(files, func(a, b file) int {
		if a.id == b.id {
			return 0
		} else if a.id > b.id {
			return -1
		}
		return 1
	})

	return
}

func (fs fileSystem) reversedFileIndexes() (indexes []int) {
	for j := len(fs) - 1; j >= 0; j -= 1 {
		if fs[j] != -1 {
			indexes = append(indexes, j)
		}
	}

	return
}

func starOne(input []string) {
	start := time.Now()
	fs, _, _ := parseFileSystem(input)

	var currentReversedIndex int
	fileMemoryIndexes := fs.reversedFileIndexes()

	for i := 0; i < fileMemoryIndexes[currentReversedIndex]; i++ {
		if fs[i] == -1 {
			fs[fileMemoryIndexes[currentReversedIndex]], fs[i] = fs[i], fs[fileMemoryIndexes[currentReversedIndex]]
			currentReversedIndex += 1
		}
	}

	res := fs.checksum()

	end := time.Now().Sub(start)

	fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
}

func starTwo(input []string) {
	start := time.Now()
	fs, freeSpaces, files := parseFileSystem(input)

fileLoop:
	for _, file := range files {
		// We have to look for a suitable free position to switch
		for i, location := range freeSpaces {
			if location.index > file.index {
				continue fileLoop
			}

			if (location.size >= file.size) && (!location.used) {
				// Modify the fileSystem
				for s, e := location.index, file.index; s < (location.index + file.size); s, e = s+1, e+1 {
					fs[s], fs[e] = fs[e], fs[s]
				}

				// Update the Free Space Location
				if file.size == location.size {
					location.used = true
				} else {
					location.size = location.size - file.size
					location.index = location.index + file.size
				}

				freeSpaces[i] = location
				continue fileLoop
			}
		}
	}

	res := fs.checksum()

	end := time.Now().Sub(start)

	fmt.Printf("Star Two: %d in %#vs\n", res, end.Seconds())
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)
}
