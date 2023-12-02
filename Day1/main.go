package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

func firstAndLast(sentence string) int {
	first := ""
	last := ""
	numbers := 0
	for _, char := range sentence {
		if unicode.IsDigit(char) {
			if numbers == 0 {
				first = string(char)
			}
			last = string(char)
			numbers += 1
		}
	}
	res, _ := strconv.Atoi(first + last)
	return res
}

func one(lines []string) int {
	res := 0
	for _, value := range lines {
		res += firstAndLast(value)
	}
	return res
}
func replace(sentence string) string {
	numberWords := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	type occurrence struct {
		word  string
		index int
	}

	var occurrences []occurrence

	// Find all occurrences of number words
	for word := range numberWords {
		re := regexp.MustCompile(word)
		indexes := re.FindAllStringIndex(sentence, -1)
		for _, index := range indexes {
			occurrences = append(occurrences, occurrence{word, index[0]})
		}
	}

	// Sort occurrences by index
	sort.Slice(occurrences, func(i, j int) bool {
		return occurrences[i].index < occurrences[j].index
	})

	// Replace the first and last occurrences if they exist
	if len(occurrences) > 0 {
		first := occurrences[0]
		last := occurrences[len(occurrences)-1]
		sentence = regexp.MustCompile(first.word).ReplaceAllString(sentence, numberWords[first.word])
		if last.word != first.word || last.index != first.index {
			sentence = regexp.MustCompile(last.word).ReplaceAllString(sentence, numberWords[last.word])
		}
	}

	// Replace the remaining occurrences
	for word, digit := range numberWords {
		if len(occurrences) == 0 || (word != occurrences[0].word && (len(occurrences) == 1 || word != occurrences[len(occurrences)-1].word)) {
			re := regexp.MustCompile(word)
			sentence = re.ReplaceAllString(sentence, digit)
		}
	}

	return sentence
}
func two(lines []string) int {
	var newLines []string
	for _, value := range lines {
		newValue := replace(value)
		newLines = append(newLines, newValue)
	}
	return one(newLines)
}

func main() {
	lines := ReadFile("./test.txt")
	//fmt.Println(one(lines))
	fmt.Println(two(lines))

}
