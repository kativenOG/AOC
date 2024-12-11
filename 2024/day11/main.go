package main

import (
	"fmt"
	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"sync"
	"time"
)

type stone string
type stoneList []stone
type stoneCounter map[stone]int

func (sc stoneCounter) update(sl stoneList, multiplier int) {
	for _, s := range sl {
		sc[s] += multiplier
	}
}

func inputStones(input []string) (sc stoneCounter) {
	sc = make(map[stone]int)
	if len(input) != 1 {
		utils.DieOnError(fmt.Errorf("Input has more tham one string %d", len(input)))
	}
	stones := strings.Split(input[0], " ")
	for _, s := range stones {
		val, err := strconv.Atoi(s)
		utils.DieOnError(err)
		sc[stone(fmt.Sprintf("%d", val))] += 1
	}

	return
}

func (s stone) blink() (sl stoneList) {
	if stoneLen := len(s); (stoneLen%2 == 0) && (stoneLen != 0) {
		runeStone := []rune(s)

		// For the first half just need to split, no need to convert
		stoneA := stone(runeStone[:stoneLen/2])

		// Convert to remove zero padding
		stoneBInt, err := strconv.Atoi(string(runeStone[stoneLen/2:]))
		utils.DieOnError(err)
		stoneB := stone(fmt.Sprintf("%d", stoneBInt))

		sl = append(sl, stoneA, stoneB)

	} else if s == "0" {
		sl = append(sl, stone("1"))
	} else {
		newStoneInt, err := strconv.Atoi(string(s))
		utils.DieOnError(err)
		sl = append(sl, stone(fmt.Sprintf("%d", newStoneInt*2024)))
	}

	return
}

func (s stone) blinkSequence(nBlinks, initialMultiplier int) int {
	sc := make(stoneCounter)
	sc[s] = initialMultiplier
	for range lo.Range(nBlinks) {
		newCounter := make(stoneCounter)
		for s, multiplier := range sc {
			newCounter.update(s.blink(), multiplier)

		}
		sc = newCounter
	}

	return lo.Sum(lo.Values(sc))
}

func travel(input []string, nBlinks int) (res int) {
	sc := inputStones(input)

	var (
		wg sync.WaitGroup
		mx sync.Mutex
	)

	for s, multiplier := range sc {
		s, multiplier := s, multiplier
		wg.Add(1)
		go func() {
			defer wg.Done()

			sRes := s.blinkSequence(nBlinks, multiplier)

			mx.Lock()
			res += sRes
			mx.Unlock()
		}()
	}

	wg.Wait()

	return
}

func starOne(input []string) {
	start := time.Now()
	res := travel(input, 25)
	end := time.Now().Sub(start)
	fmt.Printf("Star One: %d in %#vs\n", res, end.Seconds())
}

func starTwo(input []string) {
	start := time.Now()
	res := travel(input, 75)
	end := time.Now().Sub(start)
	fmt.Printf("Star Two: %d in %#vs\n", res, end.Seconds())
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input)
	starTwo(input)
}
