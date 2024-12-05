package main

import (
	"os"

	"github.com/charmbracelet/log"
	// "go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	// data, _ := os.ReadFile("example.txt")

	// log.SetLevel(log.DebugLevel)

	ansOne := PartOne(data)
	log.Info("Part One", "answer", ansOne)
	ansTwo := PartTwo(data)
	log.Info("Part Two", "answer", ansTwo)
}

func PartOne(data []byte) int {
	var answer int

	return answer
}

func PartTwo(data []byte) int {
	var answer int

	return answer
}
