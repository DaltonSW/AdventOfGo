package main

import (
	"dalton.dog/aocgo"
	"github.com/charmbracelet/log"
)

func main() {
	input := aocgo.GetInputAsString()
	// log.Info(input)
	PartA(input)
	PartB(input)
}

func PartA(input string) {
	floor := 0

	for i := range input {
		c := string(input[i])
		// log.Info(c)
		if c == "(" {
			floor++
		} else if c == ")" {
			floor--
		}
	}
	log.Infof("Part A: %v\n", floor)
	return

}

func PartB(input string) {
	floor := 0

	for i := range input {
		c := string(input[i])
		if c == "(" {
			floor++
		} else if c == ")" {
			floor--
			if floor < 0 {
				log.Infof("Part B: %v\n", i)
				return
			}
		}
	}
}
