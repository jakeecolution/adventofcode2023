package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	fnum   rune = '0'
	lnum   rune = '9'
	period rune = '.'
)

type Slice struct {
	first, last int
}

func getNumbers(line []rune) []Slice {
	var slices []Slice
	cur := Slice{first: -1, last: -1}
	for i, char := range line {
		if char >= fnum && char <= lnum {
			if cur.first == -1 {
				cur.first = i
			} else {
				cur.last = i
			}
		} else {
			if cur.first > -1 && cur.last > -1 {
				tmp := Slice{first: cur.first, last: cur.last}
				slices = append(slices, tmp)
				cur.first = -1
				cur.last = -1
			} else if cur.first > -1 {
				tmp := Slice{first: cur.first, last: cur.first}
				slices = append(slices, tmp)
				cur.first = -1
				cur.last = -1
			}
		}
	}
	return slices

}

func stringToInt(slice string) int {
	res, err := strconv.Atoi(slice)
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return 0
	}
	return res
}

func isPartNumber(num Slice, curline []rune, preline []rune, nexline []rune) int {
	var lrange, rrange int
	if num.first >= 1 {
		lrange = num.first - 1
	} else {
		lrange = 0
	}
	if num.last < len(curline)-1 {
		rrange = num.last + 1
	} else {
		rrange = num.last
	}
	if preline != nil {
		for i := lrange; i <= rrange; i++ {
			if (preline[i] > lnum || preline[i] < fnum) && (preline[i] != period) {
				// symbol
				if num.first == num.last {
					return stringToInt(string(curline[num.first]))
				}
				return stringToInt(string(curline[num.first : num.last+1]))
			}
		}
	}
	if nexline != nil {
		for i := lrange; i <= rrange; i++ {
			if (nexline[i] > lnum || nexline[i] < fnum) && (nexline[i] != period) {
				// symbol
				if num.first == num.last {
					return stringToInt(string(curline[num.first]))
				}
				return stringToInt(string(curline[num.first : num.last+1]))
			}
		}
	}
	if lrange != num.first && rrange != num.last {
		if ((curline[lrange] > lnum || curline[num.first-lrange] < fnum) && (curline[lrange] != period)) || ((curline[rrange] > lnum || curline[rrange] < fnum) && (curline[rrange] != period)) {
			if num.first == num.last {
				return stringToInt(string(curline[num.first]))
			}
			return stringToInt(string(curline[num.first : num.last+1]))
		}

	} else if num.first == lrange && num.last != rrange {
		if (curline[rrange] > lnum || curline[rrange] < fnum) && (curline[rrange] != period) {
			if num.first == num.last {
				return stringToInt(string(curline[num.first]))
			}
			return stringToInt(string(curline[num.first : num.last+1]))
		}
	} else if num.last == rrange && num.first != lrange {
		if (curline[lrange] > lnum || curline[lrange] < fnum) && (curline[lrange] != period) {
			if num.first == num.last {
				return stringToInt(string(curline[num.first]))
			}
			return stringToInt(string(curline[num.first : num.last+1]))
		}
	}
	return 0
}

func Exercise1(input [][]rune) {
	fmt.Println("Exercise 1")
	var sum, cur int
	reader := bufio.NewReader(os.Stdin)
	for idx, line := range input {
		// fmt.Printf("Line %d:", idx)
		slices := getNumbers(line)
		fmt.Printf("Line %d: ", idx)
		for _, slice := range slices {
			// if slice.last != -1 {
			// 	fmt.Printf("%s, ", string(line[slice.first:slice.last+1]))
			// } else {
			// 	fmt.Printf("%s, ", string(line[slice.first]))
			// }
			if idx == 0 && len(input) > idx+1 {
				cur = isPartNumber(slice, line, nil, input[idx+1])
			} else if idx == len(input)-1 {
				cur = isPartNumber(slice, line, input[idx-1], nil)
			} else {
				cur = isPartNumber(slice, line, input[idx-1], input[idx+1])
			}
			fmt.Printf("%d, ", cur)
			sum += cur
		}

		if idx > 0 {
			fmt.Printf("\n%s", string(input[idx-1]))
		}
		fmt.Printf("\n%s", string(line))
		if idx < len(input)-1 {
			fmt.Printf("\n%s", string(input[idx+1]))
		}
		fmt.Println("\nCheck line above: ")
		_, _ = reader.ReadString('\n')
	}
	fmt.Println("Sum: ", sum)
}

func main() {
	var inputfile string
	var exercise1 bool
	flag.StringVar(&inputfile, "input", "input.txt", "Input file")
	flag.BoolVar(&exercise1, "exercise1", true, "Exercise 1 is True, Exercise 2 is False")
	flag.Parse()

	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	height, width := 0, 0
	for scanner.Scan() {
		height++
		if charlong := scanner.Text(); width != 0 && width < len(charlong) {
			width = len(charlong)
		} else if width == 0 {
			width = len(charlong)
		}
	}
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	myrune := make([][]rune, 0)
	var counter int
	// fmt.Println("My Runes:")
	for scanner.Scan() {
		myrune = append(myrune, make([]rune, 0))
		temp := scanner.Text()
		// fmt.Println("Temp: ", temp)
		for _, char := range temp {
			myrune[counter] = append(myrune[counter], char)
		}
		counter++
	}

	if exercise1 {
		Exercise1(myrune)
	} else {
		fmt.Println("Exercise 2")
	}
}
