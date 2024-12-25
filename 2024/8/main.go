package main

import (
	"fmt"
	"math"
	"os"
	"time"

	// "strconv"
	"strings"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	"go.dalton.dog/aocgo/aocutils"
	// "go.dalton.dog/aocgo/aocutils"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    false,
	ReportTimestamp: false,
})

type Run func([][]string) int

func PrintRun(title string, run Run, data [][]string) {
	fmt.Printf("Run: %v\n", title)
	start := time.Now()
	fmt.Printf("Answer: %v\n", run(data))
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}

func main() {
	// byteData, _ := os.ReadFile("input.txt")
	exampleBytes, _ := os.ReadFile("example.txt")
	exampleData := strings.Split(string(exampleBytes), "\n")
	var exampleMatrix [][]string
	for _, line := range exampleData {
		exampleMatrix = append(exampleMatrix, strings.Split(line, ""))
	}

	data := aocgo.GetInputAsLineArray()
	var dataMatrix [][]string
	for _, line := range data {
		dataMatrix = append(dataMatrix, strings.Split(line, ""))
	}

	// log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	PrintRun("Example One", PartOne, exampleMatrix)
	PrintRun("Part One", PartOne, dataMatrix)
	PrintRun("Example Two", PartTwo, exampleMatrix)
	PrintRun("Part Two", PartTwo, dataMatrix)
}

type Coord struct {
	Row int
	Col int
}

func PartOne(data [][]string) int {
	var answer int

	var antennas map[string][]Coord
	antennas = make(map[string][]Coord)

	for x, r := range data {
		for y, c := range r {
			if c != "." {
				var arr []Coord
				arr, ok := antennas[c]
				if !ok {
					arr = make([]Coord, 0)
				}
				arr = append(arr, Coord{Row: x, Col: y})

				antennas[c] = arr
			}
		}
	}

	log.Debug("", "antennas", antennas)

	var antinodes []Coord
	antinodes = make([]Coord, 0)

	for x, r := range data {
		for y, _ := range r {
			coordToCheck := Coord{Row: x, Col: y}
			// log.Debug("", "coordToCheck", coordToCheck)
			if isAntinode(coordToCheck, antennas) {
				antinodes = append(antinodes, coordToCheck)
			}

		}
	}

	answer = len(antinodes)

	return answer
}

func isAntinode(coord Coord, antennas map[string][]Coord) bool {

	for _, coords := range antennas { // For every antenna type...
		for i := range len(coords) - 1 { // For every pairing of antennas...
			coordOne := coords[i]

			if coord == coordOne {
				continue
			}
			distOne := aocutils.Distance2D(coord.Row, coord.Col, coordOne.Row, coordOne.Col)
			slopeOne := calcSlope(coord.Row, coord.Col, coordOne.Row, coordOne.Col)
			for j := range len(coords) {
				coordTwo := coords[j]
				if coord == coordTwo {
					continue
				}
				distTwo := aocutils.Distance2D(coord.Row, coord.Col, coordTwo.Row, coordTwo.Col)
				slopeTwo := calcSlope(coord.Row, coord.Col, coordTwo.Row, coordTwo.Col)

				log.Debug("Pairing", "one", coordOne, "dist to one", distOne, "slope one", slopeOne)
				log.Debug("Pairing", "two", coordTwo, "dist to one", distTwo, "slope one", slopeTwo)

				if (distOne*2 == distTwo || distTwo*2 == distOne) && slopeOne == slopeTwo {
					log.Debug("Found", "coord", coord)
					return true
				}
			}
		}
	}

	return false

}

func calcSlope(x1, y1, x2, y2 int) float64 {
	if x1-x2 == 0 {
		return math.Inf(0)
	}
	return (float64(y2) - float64(y1)) / (float64(x2) - float64(x1))
}

func PartTwo(data [][]string) int {
	var answer int

	var antennas map[string][]Coord
	antennas = make(map[string][]Coord)

	for x, r := range data {
		for y, c := range r {
			if c != "." {
				var arr []Coord
				arr, ok := antennas[c]
				if !ok {
					arr = make([]Coord, 0)
				}
				arr = append(arr, Coord{Row: x, Col: y})

				antennas[c] = arr
			}
		}
	}

	log.Debug("", "antennas", antennas)

	var antinodes map[Coord]bool
	antinodes = make(map[Coord]bool)

	for x, r := range data {
		for y, _ := range r {
			coordToCheck := Coord{Row: x, Col: y}
			// log.Debug("", "coordToCheck", coordToCheck)
			if isAntinodeNoDist(coordToCheck, antennas) {
				antinodes[coordToCheck] = true
			}

		}
	}

	answer = len(antinodes)

	return answer
}

func isAntinodeNoDist(coord Coord, antennas map[string][]Coord) bool {
	for _, coords := range antennas { // For every antenna type...
		for i := range len(coords) - 1 { // For every pairing of antennas...
			coordOne := coords[i]

			slopeOne := calcSlope(coord.Row, coord.Col, coordOne.Row, coordOne.Col)
			for j := i + 1; j < len(coords); j++ {
				if coord == coordOne {
					return true
				}

				coordTwo := coords[j]
				slopeTwo := calcSlope(coord.Row, coord.Col, coordTwo.Row, coordTwo.Col)

				if coord == coordTwo {
					return true
				}

				// log.Debug("Pairing", "one", coordOne, "two", coordTwo, "slope one", slopeOne, "slope two", slopeTwo)

				if slopeOne == slopeTwo {
					log.Debug("Found", "coord", coord, "one", coordOne, "two", coordTwo)
					return true
				}
			}
		}
	}

	return false

}
