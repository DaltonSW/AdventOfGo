package main

import (
	// "os"
	// "strconv"
	// "strings"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ := os.ReadFile("input.txt")
	// exampleBytes, _ := os.ReadFile("example.txt")
	// exampleData := strings.Split(string(exampleBytes), "\n")

	data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	// aocgo.RunTest("Example One", PartOne, exampleData, 0)
	aocgo.RunSolve("Part One", PartOne, data)

	// aocgo.RunTest("Example Two", PartTwo, exampleData, 0)
	aocgo.RunSolve("Part Two", PartTwo, data)
}

func PartOne(data []string) int {
	var answer int

	return answer
}

func PartTwo(data []string) int {
	var answer int

	return answer
}
