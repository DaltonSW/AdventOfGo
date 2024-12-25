package main

import (
	"os"
	// "strconv"
	"strings"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ := os.ReadFile("input.txt")
	exampleBytes, _ := os.ReadFile("example.txt")
	exampleData := strings.Split(string(exampleBytes), "\n")

	data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	log.Info("Example One", "answer", PartOne(exampleData))
	log.Info("Part One", "answer", PartOne(data))

	log.Info("Example Two", "answer", PartTwo(exampleData))
	log.Info("Part Two", "answer", PartTwo(data))
}

func PartOne(data []string) int {
	var answer int

	return answer
}

func PartTwo(data []string) int {
	var answer int

	return answer
}
