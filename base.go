package main

import (
	// "os"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ := os.ReadFile("input.txt")
	// byteData, _ := os.ReadFile("example.txt")
	// strData := string(byteData)
	// data = strings.Split(strData, "\n")

	data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	ansOne := PartOne(data)
	log.Info("Part One", "answer", ansOne)
	ansTwo := PartTwo(data)
	log.Info("Part Two", "answer", ansTwo)
}

func PartOne(data []string) int {
	var answer int

	return answer
}

func PartTwo(data []string) int {
	var answer int

	return answer
}
