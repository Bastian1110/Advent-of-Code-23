package main

import (
	"bufio"
	"fmt"
	"os"
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

func getCoincidence(line string) int {
	parts := strings.SplitN(line, " | ", 2)
	listA := strings.Fields(strings.TrimSpace(parts[0]))
	listB := strings.Fields(strings.TrimSpace(parts[1]))

	itemCount := make(map[string]bool)
	count := 0

	// Add all items of list1 to the map
	for _, item := range listB {
		itemCount[item] = true
	}

	for _, item := range listA {
		if _, found := itemCount[item]; found {
			count++
			delete(itemCount, item)
		}
	}

	return count
}

func one(grid []string) int {
	res := 0
	for _, value := range grid {
		parts := strings.SplitN(value, ": ", 2)
		add := getCoincidence(parts[1])
		res += add
	}
	return res
}

func processCard(grid []string, index int, alreadyProcessed map[int]int) int {
	if count, exists := alreadyProcessed[index]; exists {
		return count
	}

	matches := getCoincidence(grid[index])
	count := 1

	for i := 1; i <= matches && index+i < len(grid); i++ {
		count += processCard(grid, index+i, alreadyProcessed)
	}

	alreadyProcessed[index] = count
	return count
}

func two(grid []string) int {
	totalCards := 0
	alreadyProcessed := make(map[int]int)

	for i := range grid {
		totalCards += processCard(grid, i, alreadyProcessed)
	}

	return totalCards
}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(two(lines))

}
