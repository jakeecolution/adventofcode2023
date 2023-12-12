package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exercise1() {
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

func exercise2() {

}

func main() {
	exercise1()
}
