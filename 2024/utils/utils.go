package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func DieOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintAndDie(val interface{}) {
	fmt.Printf("%#v\n", val)
	DieOnError(fmt.Errorf("DEBUGGING"))
}

func DebugPrintf(shouldDebug bool, args ...interface{}) {
	if !shouldDebug {
		return
	}
	fmt.Printf(args[0].(string), args[1:]...)
}

func InvertString(target string) string {
	runes := []rune(target)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ParseInputFile(filename string) (inputList []string) {
	content, err := os.ReadFile(filename)
	DieOnError(err)
	inputList = strings.Split(string(content), "\n")
	if len(inputList[len(inputList)-1]) == 0 {
		inputList = inputList[:len(inputList)-1]
	}
	return
}

func ParseFlags() (filename string, debug bool) {
	flag.StringVar(&filename, "filename", "input.txt", "the input file")
	flag.BoolVar(&debug, "debug", false, "verbose output on lines and matches")
	flag.Parse()
	return
}

func CleanTerminal(maxY int) {
	for range maxY + 1 {
		fmt.Printf("\r\033[K\033[1A")
	}
}
