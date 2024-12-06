package main

import (
	"os"

	"strings"

	"github.com/charmbracelet/log"
	// "go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ = os.ReadFile("input.txt")
	byteData, _ := os.ReadFile("example.txt")
	strData := string(byteData)
	data = strings.Split(strData, "\n")

	// data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(data)

	ansOne := PartOne(data)
	log.Info("Part One", "answer", ansOne)
	ansTwo := PartTwo(data)
	log.Info("Part Two", "answer", ansTwo)
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

func NewCoord(row, col int) Coord {
	return Coord{Row: row, Col: col}
}

func PartOne(data []string) int {
	var answer int

	var charMatrix [][]string
	for _, line := range data {
		charMatrix = append(charMatrix, strings.Split(line, ""))
	}

	obstacles, startCoord, startDir := processInput(charMatrix)

	maxBound := NewCoord(len(charMatrix), len(charMatrix[0]))

	answer = len(walkPath(startCoord, startDir, obstacles, maxBound))

	return answer
}

func walkPath(startCoord Coord, startFacingDir Direction, obstacles map[Coord]bool, maxBound Coord) map[Coord]bool {
	var walkedCoords = make(map[Coord]bool)

	facingDir := startFacingDir

	curCoord := startCoord

	log.Debug("Walking path")

	for {
		walkedCoords[curCoord] = true

		moveRow, moveCol := getDirAdjustment(facingDir)
		checkRow := curCoord.Row + moveRow
		checkCol := curCoord.Col + moveCol

		if checkRow < 0 || checkRow > maxBound.Row || checkCol < 0 || checkCol > maxBound.Col {
			return walkedCoords
		} else if checkRow == startCoord.Row && checkCol == startCoord.Col && facingDir == startFacingDir {
			return nil
		}

		val, ok := obstacles[NewCoord(checkRow, checkCol)]
		if ok && val {
			facingDir++
			if facingDir > LEFT {
				facingDir = UP
			}
		} else {
			curCoord.Row = checkRow
			curCoord.Col = checkCol
		}

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

func PartTwo(data []string) int {
	var answer int

	var charMatrix [][]string
	for _, line := range data {
		charMatrix = append(charMatrix, strings.Split(line, ""))
	}

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

	maxBound := NewCoord(len(charMatrix), len(charMatrix[0]))

	for k, _ := range walkPath(startCoord, facingDir, obstacles, maxBound) {
		newCoord := NewCoord(k.Row, k.Col)
		if newCoord.Row == startCoord.Row && newCoord.Col == startCoord.Col {
			continue
		}
		_, ok := obstacles[newCoord]
		if ok {
			continue
		}

		obstacles[newCoord] = true

		run := walkPath(startCoord, facingDir, obstacles, maxBound)
		if run == nil {
			answer++
		}

		delete(obstacles, newCoord)
	}

	for r := 0; r < len(charMatrix); r++ {
		for c := 0; c < len(charMatrix[r]); c++ {
			newCoord := NewCoord(r, c)
			if newCoord == startCoord {
				continue
			}

			val, ok := obstacles[newCoord]
			if ok && val {
				continue
			}

			obstacles[newCoord] = true

			run := walkPath(startCoord, facingDir, obstacles, maxBound)
			if run == nil {
				answer++
			}

			delete(obstacles, newCoord)
		}
	}

	return answer
}
