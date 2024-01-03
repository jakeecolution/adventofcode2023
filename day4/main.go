package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Run     int
	Winning []int
	Picks   []int
}

func (c *Card) NewCard() {
	c.Winning = make([]int, 0)
	c.Picks = make([]int, 0)
}

func (c *Card) String() string {
	return fmt.Sprintf("Run: %d\nWinning: %v\nPicks: %v\n", c.Run, c.Winning, c.Picks)
}

var re = regexp.MustCompile(`(?m)\d+`)

func (c *Card) getNumbers(line string) error {
	var err error
	run := strings.Split(line, ":")
	c.Run, err = strconv.Atoi(re.FindAllString(run[0], -1)[0])
	if err != nil {
		return err
	}
	cardstr := strings.Split(run[1], "|")
	for _, win := range re.FindAllString(cardstr[0], -1) {
		pickint, err := strconv.Atoi(win)
		if err != nil {
			return err
		}
		c.Winning = append(c.Winning, pickint)
	}
	for _, pick := range re.FindAllString(cardstr[1], -1) {
		pickint, err := strconv.Atoi(pick)
		if err != nil {
			return err
		}
		c.Picks = append(c.Picks, pickint)
	}
	return nil
}

func calculatePoints(winning []int, picks []int) int {
	var points int
	for _, pick := range picks {
		if slices.Contains(winning, pick) {
			points += 1
		}
	}

	points = int(math.Pow(2, float64(points-1)))
	return points
}

func Exercise1(cards []Card) int {
	var points int
	for _, card := range cards {
		points += calculatePoints(card.Winning, card.Picks)
	}
	return points
}

func main() {
	fmt.Println("Hello, World!")
	// flags for which text input to use and which exercise to run
	var textfile string
	var exercise bool
	flag.StringVar(&textfile, "textfile", "testinput.txt", "path to text file")
	flag.BoolVar(&exercise, "exercise", false, "run exercise 1")
	flag.Parse()
	// ingest text file
	file, err := os.Open(textfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	var cards []Card
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// parse text file into []Card
		card := Card{}
		card.NewCard()
		err := card.getNumbers(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(card.String())
		cards = append(cards, card)
	}
	if !exercise {
		fmt.Println("Exercise 1")
		fmt.Println(Exercise1(cards))
	} else {
		fmt.Println("Exercise 2 - not implemented")
	}
	// run Exercise1
}
