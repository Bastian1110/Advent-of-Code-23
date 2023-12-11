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

	partsTime := strings.SplitN(lines[0], ": ", 2)
	numbersTime := strings.Fields(strings.TrimSpace(partsTime[1]))
	partsDistance := strings.SplitN(lines[1], ": ", 2)
	numbersDistance := strings.Fields(strings.TrimSpace(partsDistance[1]))

	var times []int
	for _, str := range numbersTime {
		num, _ := strconv.Atoi(str)
		times = append(times, num)
	}
	var distances []int
	for _, str := range numbersDistance {
		num, _ := strconv.Atoi(str)
		distances = append(distances, num)
	}
	for index := range numbersTime {
		res = append(res, []int{times[index], distances[index]})
	}
	return res
}

func getCombinations(velocity []int) int {
	res := 0
	for i := 1; i < velocity[0]; i++ {
		if (velocity[0]-i)*i > velocity[1] {
			res += 1
		}
	}
	return res
}

func one(lines []string) int {
	res := 1
	pairs := parseInput(lines)
	for _, pair := range pairs {
		res *= getCombinations(pair)
	}
	return res
}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(one(lines))
}
