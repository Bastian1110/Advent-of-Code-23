package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Hand struct {
	cards []rune
	class int
	bid   int
}

func (h Hand) isGreaterThan(other Hand) bool {
	cardValues := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}
	if h.class != other.class {
		return (h.class >= other.class)
	}
	for i := 0; i < 5; i++ {
		if cardValues[h.cards[i]] > cardValues[other.cards[i]] {
			return true
		} else if cardValues[h.cards[i]] < cardValues[other.cards[i]] {
			return false
		}
	}
	return false
}

type HandSlice []Hand

// Implementing sort.Interface for MyObjectSlice
func (s HandSlice) Len() int           { return len(s) }
func (s HandSlice) Less(i, j int) bool { return s[j].isGreaterThan(s[i]) } // Use isGreaterThan here
func (s HandSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func containsExactCombination(m map[rune]int, combination []int) bool {
	mapCounts := make(map[int]int)
	for _, value := range m {
		mapCounts[value]++
	}
	comboCounts := make(map[int]int)
	for _, val := range combination {
		comboCounts[val]++
	}
	for val, count := range comboCounts {
		if mapCounts[val] != count {
			return false
		}
	}

	return true
}

func calculateClass(cards []rune) int {
	count := make(map[rune]int)
	jokers := 0

	for _, card := range cards {
		if card == 'J' {
			jokers++
		} else {
			count[card]++
		}
	}

	for j := 0; j < jokers; j++ {
		maxCard := rune(0)
		maxCount := 0
		for card, c := range count {
			if c > maxCount || (c == maxCount && card > maxCard) {
				maxCard = card
				maxCount = c
			}
		}
		count[maxCard]++
	}

	ty := 1
	switch len(count) {
	case 1:
		ty = 7
	case 2:
		if containsExactCombination(count, []int{4, 1}) {
			ty = 6
		} else if containsExactCombination(count, []int{3, 2}) {
			ty = 5
		}
	case 3:
		if containsExactCombination(count, []int{3, 1, 1}) {
			ty = 4
		} else if containsExactCombination(count, []int{2, 2, 1}) {
			ty = 3
		}
	case 4:
		ty = 2
	}
	return ty
}

func parseInput(lines []string) HandSlice {
	res := []Hand{}
	for _, value := range lines {
		parts := strings.SplitN(value, " ", 2)
		symbols := []rune(parts[0])
		number, _ := strconv.Atoi(parts[1])
		ty := calculateClass(symbols)
		hand := Hand{cards: symbols, class: ty, bid: number}
		res = append(res, hand)
	}
	return res
}

func one(lines []string) int {
	res := 0
	hands := parseInput(lines)
	sort.Sort(hands)
	for index, value := range hands {
		res += (index + 1) * value.bid
	}
	return res
}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(one(lines))
}
