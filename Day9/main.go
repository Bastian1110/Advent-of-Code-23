package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	var res []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cant open the file 😬")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func parseInput(lines []string) [][]int {
	var res [][]int
	for _, line := range lines {
		numbers := strings.Fields(strings.TrimSpace(line))
		var numbers_int []int
		for _, s := range numbers {
			sn, _ := strconv.Atoi(s)
			numbers_int = append(numbers_int, sn)
		}
		res = append(res, numbers_int)
	}
	return res
}

func isAllZero(numbers []int) bool {
	for _, value := range numbers {
		if value != 0 {
			return false
		}
	}
	return true
}

func calculateIntermediate(numbers []int) []int {
	var res []int
	for i := 0; i < len(numbers)-1; i++ {
		res = append(res, numbers[i+1]-numbers[i])
	}
	return res
}

func getLastValue(numbers []int) int {
	if isAllZero(numbers) {
		return 0
	}
	newRow := calculateIntermediate(numbers)
	return getLastValue(newRow) + numbers[len(numbers)-1]
}

func one(lines []string) int {
	res := 0
	input := parseInput(lines)
	for _, i := range input {
		res += getLastValue(i)
	}
	return res
}

func getFirstValue(numbers []int) int {
	if isAllZero(numbers) {
		return 0
	}
	newRow := calculateIntermediate(numbers)
	return numbers[0] - getFirstValue(newRow)
}

func two(lines []string) int {
	res := 0
	input := parseInput(lines)
	for _, i := range input {
		res += getFirstValue(i)
	}
	return res
}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(two(lines))
}
