package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/MrDanOak/aoc/2024/golang/solutions"
)

func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func main() {
	day := flag.Int("d", 1, "Day to run")
	flag.Parse()

	dayStr := fmt.Sprintf("%02d", *day)
	inputFile := fmt.Sprintf("inputs/day%s.txt", dayStr)
	lines, err := readLines(inputFile)
	if err != nil {
		panic(err)
	}

	var part1 string
	var part2 string
	if *day == 1 {
		part1 = solutions.Day01_Part1(lines)
		part2 = solutions.Day01_Part2(lines)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
