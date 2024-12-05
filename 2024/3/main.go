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

	re := regexp.MustCompile(regexPattern)

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

	log.Debug(string(data))

	// regexPattern := "don't\\(\\)(?s).*do\\(\\)"
	regexPattern := `(mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\))`

	re := regexp.MustCompile(regexPattern)

	// shoutouts to u/BitE01 on from the megathread for showing this reference

	matches := re.FindAllStringSubmatch(string(data), -1)

	log.Debug(matches)

	multEnabled := true

	for _, match := range matches {
		switch match[1] {
		case "do()":
			multEnabled = true
		case "don't()":
			multEnabled = false
		default:
			if multEnabled {
				numPattern := "\\d{1,3}"
				re = regexp.MustCompile(numPattern)

				firstNum, _ := strconv.Atoi(match[2])
				secondNum, _ := strconv.Atoi(match[3])

				answer += firstNum * secondNum
			}

		}
	}
	// log.Debug(newData)
	//
	// answer = PartOne([]byte(newData))

	return answer
}
