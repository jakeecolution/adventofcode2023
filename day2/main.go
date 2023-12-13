package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	red   string = "red"
	green string = "green"
	blue  string = "blue"
)

var getGameNum = regexp.MustCompile(`(?m)^Game\s(\d+)`)
var getRoundNum = regexp.MustCompile(`(?m)(\d+)\s(blue|red|green)`)

type Round struct {
	red   int
	green int
	blue  int
}

func (r Round) String() string {
	return fmt.Sprintf("(Round: red=%d, green=%d, blue=%d)", r.red, r.green, r.blue)
}

type Game struct {
	id     int
	rounds []Round
}

func (g Game) String() string {
	return fmt.Sprintf("Game: id=%d, rounds=%v", g.id, g.rounds)
}

func (g *Game) create(input string) {
	ti := strings.Split(input, ":")
	games := strings.Split(strings.TrimSpace(ti[1]), ";")
	var err error
	g.id, err = strconv.Atoi(getGameNum.FindStringSubmatch(ti[0])[1])
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return
	}
	for _, j := range games {
		round := Round{}
		for _, match := range getRoundNum.FindAllStringSubmatch(j, -1) {
			switch strings.TrimSpace(match[2]) {
			case red:
				round.red, err = strconv.Atoi(match[1])
				if err != nil {
					fmt.Printf("Error converting string to int: %v\n", err)
					return
				}
			case green:
				round.green, err = strconv.Atoi(match[1])
				if err != nil {
					fmt.Printf("Error converting string to int: %v\n", err)
					return
				}
			case blue:
				round.blue, err = strconv.Atoi(match[1])
				if err != nil {
					fmt.Printf("Error converting string to int: %v\n", err)
					return
				}
			default:
				fmt.Printf("Error converting string to int: %v\n", err)
				return
			}
		}
		g.rounds = append(g.rounds, round)
	}
}

func (g Game) possible(total *Round) (bool, *Round) {
	for idx, i := range g.rounds {
		if i.red > total.red || i.blue > total.blue || i.green > total.green {
			fmt.Println(i)
			return false, &g.rounds[idx]
		}
	}
	return true, nil
}

func (g Game) LowestPossible() Round {
	var lowest Round
	for _, i := range g.rounds {
		if i.red > lowest.red {
			lowest.red = i.red
		}
		if i.blue > lowest.blue {
			lowest.blue = i.blue
		}
		if i.green > lowest.green {
			lowest.green = i.green
		}
	}
	return lowest
}

func (g Game) LowestCube() (int, Round) {
	lowest := g.LowestPossible()
	return lowest.red * lowest.blue * lowest.green, lowest
}

func main() {
	var sum int
	var redi, bluei, greeni int
	var exercise bool
	var filename string

	flag.IntVar(&redi, "red", 12, "Red")
	flag.IntVar(&bluei, "blue", 14, "Blue")
	flag.IntVar(&greeni, "green", 13, "Green")
	flag.StringVar(&filename, "file", "input.txt", "Input file")
	flag.BoolVar(&exercise, "exercise", false, "Exercise 1 is True, Exercise 2 is False")
	flag.Parse()

	total := Round{blue: bluei, green: greeni, red: redi}
	games := make([]Game, 0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		test := scanner.Text()
		cur := Game{}
		cur.create(test)
		games = append(games, cur)
	}

	fmt.Println("len(games) =", len(games))
	for _, i := range games {
		if exercise {
			if myb, rounding := i.possible(&total); myb {
				fmt.Println(i.id, ":Possible")
				sum += i.id
			} else {
				fmt.Println(i.id, ":Impossible")
				fmt.Println(rounding)
			}
		} else {
			tcube, myround := i.LowestCube()
			fmt.Println(i.id, ":LowestCube =", tcube, "; Round =", myround)
			sum += tcube
		}
	}
	fmt.Println("Sum:", sum)
}
