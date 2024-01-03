package adventofcode

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func ParseFlags() (string, bool) {
	var input string
	var exercise bool
	flag.StringVar(&input, "input", "testinput.txt", "Path to input file")
	flag.BoolVar(&exercise, "day", false, "Day to run")
	flag.Parse()
	return input, exercise
}

func ReadFile(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
