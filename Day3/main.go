package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func isValidSymbol(c rune) bool {
	return !unicode.IsDigit(c) && c != '.'
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func getNumber(grid []string, x, y int) string {
	number := ""
	for i := x; i >= 0 && isDigit(rune(grid[y][i])); i-- {
		number = string(grid[y][i]) + number
	}
	for i := x + 1; i < len(grid[y]) && isDigit(rune(grid[y][i])); i++ {
		number += string(grid[y][i])
	}
	return number
}

func getFullNumber(grid []string, x, y int) (string, int) {
	number := ""
	endX := x
	for ; endX < len(grid[y]) && isDigit(rune(grid[y][endX])); endX++ {
		number += string(grid[y][endX])
	}
	return number, endX - 1
}

func hasAdjacentSymbol(grid []string, startX, endX, y int) bool {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for d := 0; d < 8; d++ {
		for i := startX; i <= endX; i++ {
			nx, ny := i+dx[d], y+dy[d]
			if nx >= 0 && ny >= 0 && nx < len(grid[y]) && ny < len(grid) && isValidSymbol(rune(grid[ny][nx])) {
				return true
			}
		}
	}
	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getAdjacentNumbers(grid []string, x, y int) int {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	var numbers []string

	for d := 0; d < 8; d++ {
		nx, ny := x+dx[d], y+dy[d]
		if nx >= 0 && ny >= 0 && nx < len(grid[y]) && ny < len(grid) && unicode.IsDigit(rune(grid[ny][nx])) {
			number := getNumber(grid, nx, ny)
			if !stringInSlice(number, numbers) {
				numbers = append(numbers, number)
			}
		}
	}
	if len(numbers) == 2 {
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])
		return a * b
	}
	return 0
}

func one(grid []string) int {
	var numbers []string
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if isDigit(rune(grid[y][x])) && (x == 0 || !isDigit(rune(grid[y][x-1]))) {
				number, endX := getFullNumber(grid, x, y)
				if hasAdjacentSymbol(grid, x, endX, y) {
					fmt.Println(number)
					numbers = append(numbers, number)
				}
				x = endX
			}
		}
	}

	res := 0
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		res += n
	}
	return res
}

func two(grid []string) int {
	res := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if rune(grid[y][x]) == '*' {
				res += getAdjacentNumbers(grid, x, y)
			}
		}
	}
	return res

}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(two(lines))

}
