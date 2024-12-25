package main

import (
	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	"go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	aocgo.RunSolve("Part One", PartOne, data)

	aocgo.RunSolve("Part Two", PartTwo, data)
}

func PartOne(data []string) int {
	var answer int

	for _, line := range data {
		nums := aocutils.ExtractIntsFromString(line)
		if len(nums) == 0 {
			break
		}
		l, w, h := nums[0], nums[1], nums[2]

		areaOne := l * w
		areaTwo := w * h
		areaThree := h * l

		total := areaOne*2 + areaTwo*2 + areaThree*2

		if areaOne <= areaTwo && areaOne <= areaThree {
			total += areaOne
		} else if areaTwo <= areaOne && areaTwo <= areaThree {
			total += areaTwo
		} else if areaThree <= areaOne && areaThree <= areaTwo {
			total += areaThree
		}

		answer += total
	}

	return answer
}

func PartTwo(data []string) int {
	var answer int

	for _, line := range data {
		nums := aocutils.ExtractIntsFromString(line)
		if len(nums) == 0 {
			break
		}
		l, w, h := nums[0], nums[1], nums[2]

		total := l * w * h

		if l >= w && l >= h {
			total += w + w + h + h
		} else if h >= l && h >= w {
			total += l + l + w + w
		} else if w >= l && w >= h {
			total += l + l + h + h
		}

		answer += total
	}

	return answer
}
