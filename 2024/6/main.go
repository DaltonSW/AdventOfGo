package main

import (
	"os"

	"strings"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ = os.ReadFile("input.txt")
	exampleBytes, _ := os.ReadFile("example.txt")
	exampleData := strings.Split(string(exampleBytes), "\n")

	var exampleMatrix [][]string
	for _, line := range exampleData {
		if line == "" {
			break
		}
		exampleMatrix = append(exampleMatrix, strings.Split(line, ""))
	}
	data = aocgo.GetInputAsLineArray()

	var charMatrix [][]string
	for _, line := range data {
		if line == "" {
			break
		}
		charMatrix = append(charMatrix, strings.Split(line, ""))
	}
	// log.SetLevel(log.DebugLevel)

	// log.Debug(charMatrix)

	log.Info("Example One", "answer", PartOne(exampleMatrix))

	log.Info("Part One", "answer", PartOne(charMatrix))

	log.Info("Example Two", "answer", PartTwo(exampleMatrix))

	log.Info("Part Two", "answer", PartTwo(charMatrix))
}

type Coord struct {
	Row int
	Col int
}

type Direction int

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var DirMap = map[Direction]string{
	UP:    "Up",
	RIGHT: "Right",
	DOWN:  "Down",
	LEFT:  "Left",
}

func NewCoord(row, col int) Coord {
	return Coord{Row: row, Col: col}
}

func PartOne(charMatrix [][]string) int {
	// var answer int

	obstacles, startCoord, startDir := processInput(charMatrix)

	maxBound := NewCoord(len(charMatrix), len(charMatrix[0]))

	path, _ := walkPath(startCoord, startDir, obstacles, maxBound)

	// for v, _ := range path {
	// 	if v.Row < maxBound.Row && v.Col < maxBound.Col {
	// 		answer++
	// 		// charMatrix[v.Row][v.Col] = "X"
	// 	}
	// }

	// for _, row := range charMatrix {
	// 	log.Debug(row)
	// }

	return len(path)
}

func PartTwo(charMatrix [][]string) int {
	var answer int

	obstacles, startCoord, startDir := processInput(charMatrix)

	maxBound := NewCoord(len(charMatrix), len(charMatrix[0]))

	path, _ := walkPath(startCoord, startDir, obstacles, maxBound)

	log.Debug(len(path))

	for k, _ := range path {
		log.Debug("Loop", "val", k)
		if startCoord == k {
			continue
		}

		obstacles[k] = true

		_, ok := walkPath(startCoord, startDir, obstacles, maxBound)
		if !ok {
			answer++
		}

		delete(obstacles, k)
	}

	return answer
}

func walkPath(startCoord Coord, startFacingDir Direction, obstacles map[Coord]bool, maxBound Coord) (map[Coord]int, bool) {
	var walkedCoords = make(map[Coord]int)

	facingDir := startFacingDir

	curCoord := startCoord

	walkedCoords[startCoord]++

	// log.Debug("Walking path", "start coord", startCoord, "start dir", DirMap[startFacingDir], "unique spots", len(walkedCoords))

	for {
		moveRow, moveCol := getDirAdjustment(facingDir)
		checkRow := curCoord.Row + moveRow
		checkCol := curCoord.Col + moveCol

		if checkRow < 0 || checkRow >= maxBound.Row || checkCol < 0 || checkCol >= maxBound.Col {
			walkedCoords[curCoord]++
			return walkedCoords, true
		}

		walkedCoords[curCoord]++
		if walkedCoords[curCoord] > 10 {
			return walkedCoords, false
		}
		// log.Debug("New spot!")

		val, ok := obstacles[NewCoord(checkRow, checkCol)]
		if ok && val {
			facingDir++
			if facingDir > LEFT {
				facingDir = UP
			}
			continue
		} else {
			curCoord.Row = checkRow
			curCoord.Col = checkCol
		}

		log.Debug("Step", "new coord", curCoord, "new dir", DirMap[facingDir])
	}
}

func getDirAdjustment(dir Direction) (int, int) {
	switch dir {
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	case DOWN:
		return 1, 0
	case LEFT:
		return 0, -1
	default:
		return 0, 0
	}
}

func processInput(charMatrix [][]string) (map[Coord]bool, Coord, Direction) {
	var obstacles = make(map[Coord]bool)
	var startCoord Coord
	var facingDir Direction

	for row := 0; row < len(charMatrix); row++ {
		for col := 0; col < len(charMatrix[row]); col++ {
			char := charMatrix[row][col]
			coord := NewCoord(row, col)
			switch char {
			case "^":
				startCoord = coord
				facingDir = UP
			case ">":
				startCoord = coord
				facingDir = RIGHT
			case "v":
				startCoord = coord
				facingDir = DOWN
			case "<":
				startCoord = coord
				facingDir = LEFT
			case "#":
				obstacles[coord] = true
			}

		}
	}

	return obstacles, startCoord, facingDir
}
