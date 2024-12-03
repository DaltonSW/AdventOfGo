package main

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	// "go.dalton.dog/aocgo"
	"go.dalton.dog/aocgo/aocutils"
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
	strData := string(data)
	// log.Debug(strData)
	reports := strings.Split(strings.ReplaceAll(strData, "\r\n", "\n"), "\n")
	// log.Debug(len(reports))
	var safeReports int

	for _, report := range reports {
		if len(report) < 1 {
			continue
		}
		// log.Debug(report)
		nums := aocutils.ExtractIntsFromString(report)
		// log.Debug(nums)
		if checkDiff(nums) && checkProg(nums) {
			// log.Debug("Result", "safe", nums)
			safeReports++
		} else {
			// log.Debug("Result", "fail", nums)
		}
	}

	return safeReports
}

func PartTwo(data []byte) int {
	strData := string(data)
	reports := strings.Split(strings.ReplaceAll(strData, "\r\n", "\n"), "\n")
	var safeReports int

	for _, report := range reports {
		if len(report) < 1 {
			continue
		}
		// log.Debug(report)
		nums := aocutils.ExtractIntsFromString(report)
		// log.Debug(nums)
		if checkDiff(nums) && checkProg(nums) {
			// log.Debug("Result", "safe", nums)
			safeReports++
		} else if problemDampen(nums) {
			safeReports++
		} else {
			// log.Debug("Result", "fail", nums)
		}
	}

	return safeReports
}

func problemDampen(nums []int) bool {
	log.Debug("Trying to dampen", "nums", nums)
	for i := 0; i < len(nums); i++ {
		newNums := removeIndex(nums, i)
		// log.Debug(newNums)

		if checkDiff(newNums) && checkProg(newNums) {
			log.Debug("Result", "safe w/ dampen", newNums)
			return true
		}
	}
	return false

}

func removeIndex(nums []int, idx int) []int {
	outNum := make([]int, 0)
	outNum = append(outNum, nums[:idx]...)
	return append(outNum, nums[idx+1:]...)
}

func checkDiff(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 {
			return false
		} else if diff < 0 {
			diff *= -1
		}

		if diff > 3 {
			return false
		}
	}
	return true
}

func checkProg(nums []int) bool {
	var isIncreasing bool
	if nums[0] < nums[1] {
		isIncreasing = true
	} else if nums[1] < nums[0] {
		isIncreasing = false
	} else {
		return false
	}

	for i := 1; i < len(nums)-1; i++ {
		if (nums[i] <= nums[i+1]) && !isIncreasing {
			return false
		} else if (nums[i] >= nums[i+1]) && isIncreasing {
			return false
		}
	}
	return true
}
