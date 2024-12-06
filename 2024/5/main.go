package main

import (
	// "os"

	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"go.dalton.dog/aocgo"
	// "go.dalton.dog/aocgo/aocutils"
)

func main() {
	var data []string
	// byteData, _ := os.ReadFile("input.txt")
	// byteData, _ := os.ReadFile("example.txt")
	// strData := string(byteData)
	// data = strings.Split(strData, "\n")

	data = aocgo.GetInputAsLineArray()

	log.SetLevel(log.DebugLevel)

	log.Debug(len(data))

	ansOne, rules, wrongUpdates := PartOne(data)
	log.Info("Part One", "answer", ansOne)
	ansTwo := PartTwo(wrongUpdates, rules)
	log.Info("Part Two", "answer", ansTwo)
}

func PartOne(data []string) (int, map[int][]int, [][]int) {
	var answer int
	var rules map[int][]int
	var updates []string
	var wrongUpdates [][]int

	rules = make(map[int][]int)

	for idx, rule := range data {
		nums := strings.Split(rule, "|")

		if len(nums) == 1 {
			updates = append(updates, data[idx+1:]...)
			break
		}

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		arr, ok := rules[x]
		if !ok {
			arr = make([]int, 0)
		}
		arr = append(arr, y)
		rules[x] = arr
	}

	for _, update := range updates {
		if update == "" {
			break
		}
		// log.Debug("Update!", "update", update)
		pages := strings.Split(update, ",")
		pageNums := make([]int, 0, len(pages))

		for _, page := range pages {
			p, _ := strconv.Atoi(page)
			pageNums = append(pageNums, p)
		}

		if checkPage(pageNums, rules) {
			log.Debug("", "true page", pageNums, "middle page", pageNums[len(pageNums)/2])
			answer += pageNums[len(pageNums)/2]
		} else {
			wrongUpdates = append(wrongUpdates, pageNums)
		}

	}

	return answer, rules, wrongUpdates
}

func checkPage(pageNums []int, rules map[int][]int) bool {
	if len(pageNums) < 2 {
		return false
	}
	for i := len(pageNums) - 1; i >= 0; i-- {
		pagesAfter, _ := rules[pageNums[i]]
		for j := 0; j < i; j++ {
			if slices.Contains(pagesAfter, pageNums[j]) {
				return false
			}
		}
	}
	return true
}

func PartTwo(updates [][]int, rules map[int][]int) int {
	var answer int

	for _, update := range updates {
		log.Debug("", "unsorted", update)
		slices.SortFunc(update, func(a, b int) int {
			rule, _ := rules[a]
			if slices.Contains(rule, b) {
				return -1
			}
			return 0
		})

		log.Debug("", "sorted", update)
		answer += update[len(update)/2]
	}

	return answer
}
