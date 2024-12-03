package main

import (
	"os"
	"regexp"
	"strconv"
	// "strings"

	"github.com/charmbracelet/log"
	// "go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	// data, _ := os.ReadFile("example.txt")

	log.SetLevel(log.DebugLevel)

	ansOne := PartOne(data)
	log.Info("Part One", "answer", ansOne)
	ansTwo := PartTwo(data)
	log.Info("Part Two", "answer", ansTwo)
}

func PartOne(data []byte) int {
	// log.Debug(strData)

	var answer int

	regexPattern := "mul\\(\\d{1,3},\\d{1,3}\\)"

	re, err := regexp.Compile(regexPattern)
	if err != nil {
		log.Fatal(err)
	}

	matches := re.FindAll(data, -1)

	for _, match := range matches {
		numPattern := "\\d{1,3}"
		numRe, err := regexp.Compile(numPattern)
		if err != nil {
			log.Fatal(err)
		}

		nums := numRe.FindAll(match, 2)

		firstNum, _ := strconv.Atoi(string(nums[0]))
		secondNum, _ := strconv.Atoi(string(nums[1]))

		answer += firstNum * secondNum
	}

	return answer
}

func PartTwo(data []byte) int {
	var answer int

	regexPattern := "don't\\(\\).*do\\(\\)"

	re, err := regexp.Compile(regexPattern)
	if err != nil {
		log.Fatal(err)
	}

	newData := re.ReplaceAllString(string(data), "")

	regexPattern = "don't\\(\\).*$"

	re, err = regexp.Compile(regexPattern)
	if err != nil {
		log.Fatal(err)
	}

	newData = re.ReplaceAllString(string(newData), "")

	answer = PartOne([]byte(newData))

	return answer
}
