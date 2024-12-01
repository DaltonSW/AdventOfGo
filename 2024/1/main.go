package main

import (
	"os"
	"sort"

	"github.com/charmbracelet/log"
	// "go.dalton.dog/aocgo"
	"go.dalton.dog/aocgo/aocutils"
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
	numbers := aocutils.ExtractIntsFromString(string(data))
	leftList := make([]int, 0, len(numbers)/2)
	rightList := make([]int, 0, len(numbers)/2)

	// log.Debug(numbers)

	idx := 0
	for idx < len(numbers) {
		if idx%2 == 0 {
			leftList = append(leftList, numbers[idx])
		} else {
			rightList = append(rightList, numbers[idx])
		}
		idx++
	}

	numbers = nil

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	// log.Debug("", "Left List", leftList)
	// log.Debug("", "Right List", rightList)

	var dist int
	idx = 0
	for idx < len(leftList) {
		tempDist := rightList[idx] - leftList[idx]
		if tempDist < 0 {
			tempDist *= -1
		}
		dist += tempDist
		idx++
	}

	// log.Info("", "dist", dist)

	return dist
}

func PartTwo(data []byte) int {
	numbers := aocutils.ExtractIntsFromString(string(data))
	leftList := make([]int, 0, len(numbers)/2)
	rightList := make(map[int]int)

	// log.Debug(numbers)

	idx := 0
	for idx < len(numbers) {
		num := numbers[idx]
		if idx%2 == 0 {
			leftList = append(leftList, num)
		} else {
			_, ok := rightList[num]
			if ok {
				rightList[num]++
			} else {
				rightList[num] = 1
			}
		}
		idx++
	}

	// log.Debug("", "Left List", leftList)
	// log.Debug("", "Right List", rightList)

	var similarity int
	idx = 0
	for idx < len(leftList) {
		similarity += leftList[idx] * rightList[leftList[idx]]
		idx++
	}

	// log.Info("", "similarity", similarity)

	return similarity
}
