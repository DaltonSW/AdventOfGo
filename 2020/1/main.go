package main

import (
	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
)

func main() {
	data := aocgo.GetInputAsString()

	log.SetLevel(log.DebugLevel)
	log.Debug(data)
}
