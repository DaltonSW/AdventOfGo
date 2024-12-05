package main

import (
	"os"
	"regexp"
	"strings"

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
	var stepTotal int
	var strData string

	var dataRows []string
	var charMatrix [][]string

	strData = string(data)

	dataRows = strings.Split(strData, "\n")

	forwardRowRegex := regexp.MustCompile(`XMAS`)
	backwardRowRegex := regexp.MustCompile("SAMX")

	// This should check every horizontal possibility and build out the matrix
	for _, row := range dataRows {
		if row == "" {
			continue
		}

		forwardMatches := forwardRowRegex.FindAllString(row, -1)
		backwardMatches := backwardRowRegex.FindAllString(row, -1)
		// forwardMatches := regexp.
		// forwardMatches := regexp.FindAllString(row, -1)
		// matches := rowRegex.FindAllString(row, -1)
		// log.Debug("Row", "row", row, "forward matches", forwardMatches)
		// log.Debug("Row", "row", row, "backward matches", backwardMatches)
		log.Debug("Row", "row", row, "found", len(forwardMatches)+len(backwardMatches))
		stepTotal += len(forwardMatches) + len(backwardMatches) // This now gets XMAS and SMAX in every row
		charMatrix = append(charMatrix, strings.Split(row, ""))
	}

	log.Debug("Found", "horizontal", stepTotal)

	answer += stepTotal
	stepTotal = 0

	// Should check every vertical possibility
	for r := 0; r < len(charMatrix)-3; r++ {
		for c := 0; c < len(charMatrix[r]); c++ {
			checkStr := charMatrix[r][c] + charMatrix[r+1][c] + charMatrix[r+2][c] + charMatrix[r+3][c]
			if CheckStr(checkStr) {
				stepTotal++
			}
		}
	}

	log.Debug("Found", "vertical", stepTotal)

	answer += stepTotal
	stepTotal = 0

	for r := 0; r < len(charMatrix)-3; r++ {
		for c := 0; c < len(charMatrix[r])-3; c++ {
			checkStr := charMatrix[r][c] + charMatrix[r+1][c+1] + charMatrix[r+2][c+2] + charMatrix[r+3][c+3]
			if CheckStr(checkStr) {
				stepTotal++
			}
		}
	}

	log.Debug("Found", "forward diagonal", stepTotal)

	answer += stepTotal
	stepTotal = 0

	for r := 0; r < len(charMatrix)-3; r++ {
		for c := 3; c < len(charMatrix[r]); c++ {
			checkStr := charMatrix[r][c] + charMatrix[r+1][c-1] + charMatrix[r+2][c-2] + charMatrix[r+3][c-3]
			if CheckStr(checkStr) {
				stepTotal++
			}
		}
	}

	log.Debug("Found", "backwards diagonal", stepTotal)

	answer += stepTotal
	stepTotal = 0

	return answer
}

func PartTwo(data []byte) int {
	var answer int
	var strData string

	var dataRows []string
	var charMatrix [][]string

	strData = string(data)

	dataRows = strings.Split(strData, "\n")

	// This should check every horizontal possibility and build out the matrix
	for _, row := range dataRows {
		if row == "" {
			continue
		}

		charMatrix = append(charMatrix, strings.Split(row, ""))
	}

	for r := 1; r < len(charMatrix)-1; r++ {
		for c := 1; c < len(charMatrix[r])-1; c++ {
			wordOne := charMatrix[r-1][c-1] + charMatrix[r][c] + charMatrix[r+1][c+1]
			wordTwo := charMatrix[r-1][c+1] + charMatrix[r][c] + charMatrix[r+1][c-1]
			if (wordOne == "MAS" || wordOne == "SAM") && (wordTwo == "MAS" || wordTwo == "SAM") {
				answer++
			}
		}
	}

	return answer
}

func CheckStr(checkStr string) bool {
	log.Debug("Checking", "checkStr", checkStr)
	if checkStr == "XMAS" || checkStr == "SAMX" {
		return true
	}

	return false
}
