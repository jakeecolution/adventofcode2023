package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	fnum rune = '0'
	lnum rune = '9'
)

func calibrate(input string, output chan int) {
	first, last := 'a', 'a'
	for _, runeval := range input {
		if runeval >= fnum && runeval <= lnum {
			if first == 'a' {
				first = runeval
			} else {
				last = runeval
			}
		}
	}
	var sum string
	if last == 'a' {
		sum = string(first) + string(first)
	} else {
		sum = string(first) + string(last)
	}

	num, err := strconv.Atoi(sum)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		output <- 0
		return
	}
	output <- num
}

func matchme(match string) (string, string) {
	switch match {
	case "one":
		return "1", "1e"
	case "two":
		return "2", "2o"
	case "three":
		return "3", "3e"
	case "four":
		return "4", "4r"
	case "five":
		return "5", "5e"
	case "six":
		return "6", "6x"
	case "seven":
		return "7", "7n"
	case "eight":
		return "8", "8t"
	case "nine":
		return "9", "9e"
	default:
		return match, match
	}
}

func calibrate_v2(input string) int {
	var re = regexp.MustCompile(`(?m)\d|one|two|three|four|five|six|seven|eight|nine`)
	var sum int
	var stotal string

	matches := re.FindAllStringSubmatchIndex(input, -1)
	if len(matches) > 0 {
		temp, replacement := matchme(input[matches[0][0]:matches[0][1]])
		input = input[:matches[0][0]] + replacement + input[matches[0][1]:]
		stotal += temp

	} else {
		return 0
	}
	matchesv2 := re.FindAllString(input, -1)
	count1 := len(matchesv2)
	if count1 == 0 {
		stotal = stotal + stotal
	} else {
		temp, _ := matchme(matchesv2[count1-1])
		stotal += temp
	}

	sum, err := strconv.Atoi(stotal)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return 0
	}
	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exercise1(mychan chan int, fs *bufio.Scanner, lc int) {

	// wg.Add(lc)
	for fs.Scan() {
		// wg.Add(1)
		temp := func(input string, output chan int) {
			// defer wg.Done()
			calibrate(input, output)
		}
		go temp(fs.Text(), mychan)
	}
	// wg.Wait()
	var result int = 0
	for {
		select {
		case r := <-mychan:
			result += r
		case <-time.After(1 * time.Second):
			fmt.Printf("Result: %v\n", result)
			return
		}
	}
}

func exercise2(mychan chan int, fs *bufio.Scanner, lc int) {
	var res int
	for fs.Scan() {
		// temp := func(input string, output chan int) {
		// 	calibrate_v2(input, output)
		// }
		// go temp(fs.Text(), mychan)
		res += calibrate_v2(fs.Text())
	}
	// var result int = 0
	// for {
	// 	select {
	// 	case r := <-mychan:
	// 		result += r
	// 	case <-time.After(2 * time.Second):
	fmt.Printf("Result: %v\n", res)
	// 	return
	// }
	// }

}

func main() {
	var inputfile string
	// var wg sync.WaitGroup
	flag.StringVar(&inputfile, "inputfile", "./analyzeme.txt", "A file with line(s) to calibrate given AOC day 1")
	flag.Parse()

	fmt.Println("Reading file: ", inputfile)
	f, err := os.Open(inputfile)
	check(err)
	defer f.Close()
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	lc := 0
	for fs.Scan() {
		lc++
	}
	fmt.Printf("Number of lines: %v\n", lc)
	mychan := make(chan int, lc+1)
	if _, err := f.Seek(0, 0); err != nil {
		fmt.Println("Error resetting file pointer: ", err)
		return
	}
	fs = bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	exercise1(mychan, fs, lc)
	// exercise2(mychan, fs, lc)
}
