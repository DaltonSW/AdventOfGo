package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ := os.ReadFile("input.txt")
	exampleBytes, _ := os.ReadFile("example.txt")
	exampleData := strings.Split(string(exampleBytes), "\n")

	data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	// log.Info("Example One", "answer", PartOne(exampleData))
	// log.Info("Part One", "answer", PartOne(data))

	log.Info("Example Two", "answer", PartTwo(exampleData))
	log.Info("Part Two", "answer", PartTwo(data))
}

func PartOne(data []string) int {
	var answer int

	for _, line := range data {
		if line == "" {
			continue
		}
		strNums := strings.Split(line, " ")
		target, _ := strconv.Atoi(strings.ReplaceAll(strNums[0], ":", ""))
		var nums []int
		for _, s := range strNums[1:] {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		if partOneRecur(nums, target, nums[0], 1) {
			log.Debug("Valid", "target", target, "nums", nums)
			answer += target
		} else {
			log.Debug("Invalid", "target", target, "nums", nums)
		}
	}

	return answer
}

func partOneRecur(numbers []int, target, result, index int) bool {
	if index == len(numbers) {
		return result == target
	}

	nextNum := numbers[index]

	if partOneRecur(numbers, target, result+nextNum, index+1) {
		return true
	} else if partOneRecur(numbers, target, result*nextNum, index+1) {
		return true
	}

	return false
}

func PartTwo(data []string) uint64 {
	var answer uint64

	for _, line := range data {
		if line == "" {
			continue
		}
		strNums := strings.Split(line, " ")
		target, _ := strconv.Atoi(strings.ReplaceAll(strNums[0], ":", ""))
		var nums []int
		for _, s := range strNums[1:] {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		if partTwoRecur(nums, target, nums[0], 1) {
			log.Debug("Valid", "target", target, "nums", nums)
			answer += uint64(target)
		} else {
			log.Debug("Invalid", "target", target, "nums", nums)
		}
	}

	return answer
}

func partTwoRecur(numbers []int, target, result, index int) bool {
	if index == len(numbers) {
		return result == target
	}

	nextNum := numbers[index]

	if partTwoRecur(numbers, target, result+nextNum, index+1) {
		return true
	} else if partTwoRecur(numbers, target, result*nextNum, index+1) {
		return true
	} else if partTwoRecur(numbers, target, concatInt(result, nextNum), index+1) {
		return true
	}

	return false
}

func concatInt(a, b int) int {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	outStr := strA + strB
	outInt, _ := strconv.Atoi(outStr)
	return outInt
}
