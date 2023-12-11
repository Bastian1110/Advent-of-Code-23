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

type DFA struct {
	actual string
	nodes  map[string][]string
}

func (m *DFA) Forward(input int) {
	m.actual = m.nodes[m.actual][input]
}

func (m *DFA) RunSequence(sequence []int, steps int) int {
	for i := 0; i < steps; i++ {
		index := i % len(sequence)
		symbol := sequence[index]
		m.Forward(symbol)
		fmt.Println("Input : ", symbol, "State :", m.actual)
		if m.actual == "ZZZ" {
			return i + 1
		}
	}
	return steps
}

func isLineEmpty(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return trimmedLine == ""
}

func parseInput(lines []string) ([]int, DFA) {
	var sequence []int
	for _, s := range lines[0] {
		if s == 'L' {
			sequence = append(sequence, 0)
		} else {
			sequence = append(sequence, 1)
		}
	}
	var nodes = make(map[string][]string)
	for i := 1; i < len(lines); i++ {
		if !isLineEmpty(lines[i]) {

			parts := strings.SplitN(lines[i], " = ", 2)
			trimmed := strings.Trim(parts[1], "()")
			partsDirections := strings.Split(trimmed, ",")
			for i, part := range partsDirections {
				partsDirections[i] = strings.TrimSpace(part)
			}
			nodes[parts[0]] = partsDirections
		}
	}
	myDFA := DFA{actual: "AAA", nodes: nodes}
	return sequence, myDFA
}

func one(lines []string) int {
	sequence, dfa := parseInput(lines)
	res := dfa.RunSequence(sequence, 100000)
	return res
}

type NFA struct {
	currentStates []string
	nodes         map[string][]string
}

func (n *NFA) Forward(input int) {
	nextStates := make([]string, 0)
	for _, state := range n.currentStates {
		nextState := n.nodes[state][input]
		nextStates = append(nextStates, nextState)
	}
	n.currentStates = nextStates
}

func allEndWithZ(states []string) bool {
	for _, state := range states {
		if !strings.HasSuffix(state, "Z") {
			return false
		}
	}
	return true
}

func two(lines []string) int {
	sequence, dfa := parseInput(lines)

	// Initialize NFA with all states ending with 'A'
	var startStates []string
	for state := range dfa.nodes {
		if strings.HasSuffix(state, "A") {
			startStates = append(startStates, state)
		}
	}

	nfa := NFA{currentStates: startStates, nodes: dfa.nodes}

	steps := 0
	for !allEndWithZ(nfa.currentStates) {
		index := steps % len(sequence)
		symbol := sequence[index]
		nfa.Forward(symbol)
		fmt.Println("Step:", steps, "States:", nfa.currentStates)
		steps++
	}

	return steps
}

func main() {
	lines := ReadFile("./input.txt")
	fmt.Println(two(lines))
}
