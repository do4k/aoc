package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"aoc2025/solutions"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86")).
			MarginBottom(1)

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212"))

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func main() {
	day := flag.Int("day", 0, "Day number to run (1-25)")
	example := flag.Bool("example", false, "Use example input instead of puzzle input")
	flag.Parse()

	// If no day specified, run interactive mode
	if *day == 0 {
		runInteractive()
		return
	}

	runDay(*day, *example)
}

func runInteractive() {
	for {
		// Build day options from registry
		days := getSortedDays()
		dayOptions := make([]huh.Option[int], 0, len(days)+1)
		for _, d := range days {
			dayOptions = append(dayOptions, huh.NewOption(fmt.Sprintf("Day %02d", d), d))
		}
		dayOptions = append(dayOptions, huh.NewOption("Exit", 0))

		var selectedDay int
		err := huh.NewSelect[int]().
			Title("🎄 Advent of Code 2025").
			Description("Select a day to run").
			Options(dayOptions...).
			Value(&selectedDay).
			Run()

		if err != nil || selectedDay == 0 {
			fmt.Println(labelStyle.Render("Goodbye! 🎅"))
			return
		}

		// Select input type
		inputOptions := []huh.Option[string]{
			huh.NewOption("Puzzle Input", "puzzle"),
			huh.NewOption("Example Input", "example"),
			huh.NewOption("Back", "back"),
		}

		var inputType string
		err = huh.NewSelect[string]().
			Title(fmt.Sprintf("Day %02d - Input Type", selectedDay)).
			Options(inputOptions...).
			Value(&inputType).
			Run()

		if err != nil || inputType == "back" {
			continue
		}

		useExample := inputType == "example"
		runDay(selectedDay, useExample)

		// Pause before returning to menu
		fmt.Println()
		fmt.Println(labelStyle.Render("Press Enter to continue..."))
		fmt.Scanln()
	}
}

func runDay(day int, example bool) {
	solver, ok := solutions.Registry[day]
	if !ok {
		fmt.Printf("Day %d not implemented yet\n", day)
		return
	}

	lines, err := readInput(day, example)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	inputLabel := "puzzle"
	if example {
		inputLabel = "example"
	}

	fmt.Println()
	fmt.Println(titleStyle.Render(fmt.Sprintf("🎄 Day %02d (%s input)", day, inputLabel)))
	fmt.Println(labelStyle.Render("Part 1: ") + resultStyle.Render(fmt.Sprintf("%v", solver.Part1(lines))))
	fmt.Println(labelStyle.Render("Part 2: ") + resultStyle.Render(fmt.Sprintf("%v", solver.Part2(lines))))
}

func getSortedDays() []int {
	days := make([]int, 0, len(solutions.Registry))
	for d := range solutions.Registry {
		days = append(days, d)
	}
	sort.Ints(days)
	return days
}

func readInput(day int, example bool) ([]string, error) {
	suffix := ""
	if example {
		suffix = "_example"
	}
	filename := fmt.Sprintf("inputs/day%02d%s.txt", day, suffix)
	path := filepath.Join(".", filename)

	file, err := os.Open(path)
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
