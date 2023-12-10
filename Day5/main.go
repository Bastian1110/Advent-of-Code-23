package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

func parseInput(lines []string) ([]int, [][]int) {
	parts := strings.SplitN(lines[0], ": ", 2)
	seeds := strings.Fields(strings.TrimSpace(parts[1]))
	var seeds_int []int
	for _, s := range seeds {
		sn, _ := strconv.Atoi(s)
		seeds_int = append(seeds_int, sn)
	}

	lines = lines[1:]
	var maps [][]int
	var currentMap []int

	for _, line := range lines {
		if line == "" || strings.HasSuffix(line, "map:") {
			if len(currentMap) > 0 {
				maps = append(maps, currentMap)
				currentMap = nil
			}
			continue
		}
		parts := strings.Fields(line)
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			currentMap = append(currentMap, num)
		}
	}
	if len(currentMap) > 0 {
		maps = append(maps, currentMap)
	}
	return seeds_int, maps
}

func forwardMap(seed int, maap []int) int {
	for i := 0; i < len(maap); i += 3 {
		dest_range := maap[i+1] + maap[i+2]
		if maap[i+1] <= seed && seed < dest_range {
			n := maap[i] - maap[i+1]
			return seed + n
		}
	}
	return seed
}

func getLocation(seed int, maps [][]int) int {
	actual := seed
	for _, m := range maps {
		actual = forwardMap(actual, m)
	}
	return actual
}

func one(lines []string) int {
	seeds, maps := parseInput(lines)
	var locations []int
	for _, s := range seeds {
		locations = append(locations, getLocation(s, maps))
	}
	fmt.Println(locations)
	min := locations[0]
	for _, value := range locations {
		if value < min {
			min = value
		}
	}
	return min
}

func two(lines []string) int {
	seeds, maps := parseInput(lines)
	fmt.Println("seeds: ", len(seeds))
	var locations []int
	for _, s := range seeds {
		locations = append(locations, getLocation(s, maps))
	}
	min := locations[0]
	for _, value := range locations {
		if value < min {
			min = value
		}
	}
	return min
}

func streamRealSeedsToFile(seeds []int, file *os.File) {
	numCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numCores)

	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex for synchronizing file writes
	chunks := chunkSeeds(seeds, numCores)

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			processAndWriteChunk(chunk, file, &mu)
		}(chunk)
	}

	wg.Wait()
}

func processAndWriteChunk(chunk []int, file *os.File, mu *sync.Mutex) {
	for i := 0; i < len(chunk); i += 2 {
		for j := chunk[i]; j < chunk[i]+chunk[i+1]; j++ {
			mu.Lock()
			_, err := file.WriteString(fmt.Sprintf("%d\n", j))
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
			mu.Unlock()
		}
	}
}

func chunkSeeds(seeds []int, numChunks int) [][]int {
	var chunks [][]int
	chunkSize := (len(seeds) + numChunks - 1) / numChunks
	chunkSize += chunkSize % 2 // Ensure chunk size is even
	for i := 0; i < len(seeds); i += chunkSize {
		end := i + chunkSize
		if end > len(seeds) {
			end = len(seeds)
		}
		chunks = append(chunks, seeds[i:end])
	}
	return chunks
}

func writeSeedsFile(seeds []int) {
	file, err := os.Create("output_seeds.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	streamRealSeedsToFile(seeds, file)
}

func main() {
	lines := ReadFile("./input.txt")
	seeds, _ := parseInput(lines)
	writeSeedsFile(seeds)
}
