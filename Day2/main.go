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

func getColorMaps(input string) []map[string]int {
	groups := strings.Split(input, "; ")

	var results []map[string]int

	for _, group := range groups {
		colorCountPairs := strings.Split(group, ", ")
		colorCountMap := make(map[string]int)

		for _, pair := range colorCountPairs {
			parts := strings.Split(pair, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting count to integer:", err)
				continue
			}
			color := parts[1]
			colorCountMap[color] = count
		}

		results = append(results, colorCountMap)
	}
	return results
}

func validate(colorMaps map[string]int, config map[string]int) bool {
	res := true
	for color, count := range colorMaps {
		if count > config[color] {
			res = false
			break
		}
	}
	return res
}

func one(lines []string) int {
	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	res := 0
	for index, value := range lines {
		parts := strings.SplitN(value, ": ", 2)
		maps := getColorMaps(parts[1])
		valid := true
		for _, m := range maps {
			if !validate(m, config) {
				valid = false
			}
		}
		if valid {
			res += index + 1
		}
	}
	return res
}

func getMaxFromMaps(maps []map[string]int) int {
	red := 1
	green := 1
	blue := 1

	for _, m := range maps {
		r, ro := m["red"]
		if ro {
			if r > red {
				red = r
			}
		}
		g, eo := m["green"]
		if eo {
			if g > green {
				green = g
			}
		}
		b, bo := m["blue"]
		if bo {
			if b > blue {
				blue = b
			}
		}
	}
	return red * green * blue
}

func two(lines []string) int {
	res := 0
	for _, value := range lines {
		parts := strings.SplitN(value, ": ", 2)
		maps := getColorMaps(parts[1])
		res += getMaxFromMaps(maps)
	}
	return res
}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(two(lines))

}
