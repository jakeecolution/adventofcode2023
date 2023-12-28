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
	star   rune = '*'
)

type Slice struct {
	first, last int
}

type SliceContext struct {
	slice Slice
	line  int
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
	if cur.first > -1 && cur.last > -1 {
		tmp := Slice{first: cur.first, last: cur.last}
		slices = append(slices, tmp)
	} else if cur.first > -1 {
		tmp := Slice{first: cur.first, last: cur.first}
		slices = append(slices, tmp)
	}
	return slices

}

func getNumber(line []rune, pos int) Slice {
	var lrange, rrange int
	for i := pos; i >= 0; i-- {
		if line[i] <= lnum && line[i] >= fnum {
			lrange = i
		} else {
			break
		}
	}
	for i := pos; i < len(line); i++ {
		if line[i] <= lnum && line[i] >= fnum {
			rrange = i
		} else {
			break
		}
	}
	return Slice{first: lrange, last: rrange}
}

func getNumberInt(line []rune, pos int) int {
	sliced := getNumber(line, pos)
	return stringToInt(string(line[sliced.first : sliced.last+1]))
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
		if (curline[lrange] != period) || (curline[rrange] != period) {
			if num.first == num.last {
				return stringToInt(string(curline[num.first]))
			}
			return stringToInt(string(curline[num.first : num.last+1]))
		}

	} else if num.first == lrange && num.last != rrange {
		if curline[rrange] != period {
			if num.first == num.last {
				return stringToInt(string(curline[num.first]))
			}
			return stringToInt(string(curline[num.first : num.last+1]))
		}
	} else if num.last == rrange && num.first != lrange {
		if curline[lrange] != period {
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
	for idx, line := range input {
		slices := getNumbers(line)
		for _, slice := range slices {

			if idx == 0 && len(input) > idx+1 {
				cur = isPartNumber(slice, line, nil, input[idx+1])
			} else if idx == len(input)-1 {
				cur = isPartNumber(slice, line, input[idx-1], nil)
			} else {
				cur = isPartNumber(slice, line, input[idx-1], input[idx+1])
			}
			if cur != 0 {
				sum += cur
			}
		}

	}
	fmt.Println("Sum: ", sum)
}

func IsDigit(char rune) bool {
	return char >= fnum && char <= lnum
}

func isGearStar(spos int, curline []rune, preline []rune, nexline []rune) int {
	var lrange, rrange int
	if spos >= 1 {
		lrange = spos - 1
	} else {
		lrange = 0
	}
	if spos < len(curline)-1 {
		rrange = spos + 1
	} else {
		rrange = spos
	}
	nums := make([]int, 0)
	if preline != nil {
		for i := lrange; i <= rrange; i++ {
			if IsDigit(preline[i]) {
				sliced := getNumber(preline, i)
				nums = append(nums, stringToInt(string(preline[sliced.first:sliced.last+1])))
				if sliced.first <= spos && sliced.last >= spos {
					break
				}
			}
		}
	}
	if nexline != nil {
		for i := lrange; i <= rrange; i++ {
			if IsDigit(nexline[i]) {
				sliced := getNumber(nexline, i)
				nums = append(nums, stringToInt(string(nexline[sliced.first:sliced.last+1])))
				if sliced.first <= spos && sliced.last >= spos {
					break
				}
			}
		}
	}
	if lrange != spos && rrange != spos {
		if IsDigit(curline[lrange]) {
			nums = append(nums, getNumberInt(curline, lrange))
		}
		if IsDigit(curline[rrange]) {
			nums = append(nums, getNumberInt(curline, rrange))
		}

	} else if spos == lrange && spos != rrange {
		if IsDigit(curline[rrange]) {
			nums = append(nums, getNumberInt(curline, rrange))
		}
	} else if spos == rrange && spos != lrange {
		if IsDigit(curline[lrange]) {
			nums = append(nums, getNumberInt(curline, lrange))
		}
	}
	if len(nums) == 2 {
		// fmt.Println("Gear: ", nums, " = ", nums[0]*nums[1])
		return nums[0] * nums[1]
	}
	return 0
}

func Exercise2(input [][]rune) {
	fmt.Println("Exercise 2")
	var sum, mygear int
	for idx, line := range input {
		for index, myrune := range line {
			if myrune != star {
				continue
			}
			var lrange, rrange int
			if index >= 1 {
				lrange = index - 1
			} else {
				lrange = 0
			}
			if index < len(line)-1 {
				rrange = index + 1
			} else {
				rrange = index
			}
			if idx == 0 && len(input) > idx+1 {
				fmt.Printf("Selected at index %d: \n\t[%d:%d]=%s\n\t%s\n", index, lrange, rrange+1, string(line[lrange:rrange+1]), string(input[idx+1][lrange:rrange+1]))
				// mygear = isGearStar(index, line, nil, input[idx+1])
			} else if idx == len(input)-1 {
				mygear = isGearStar(index, line, input[idx-1], nil)
				// fmt.Printf("Selected at index %d: \n\t%s\n\t%s\n", index, string(input[idx-1][lrange:rrange+1]), string(line[lrange:rrange+1]))
			} else {
				mygear = isGearStar(index, line, input[idx-1], input[idx+1])
				// fmt.Printf("Selected at index %d: \n\t%s\n\t[%d:%d]=%s\n\t%s\n", index, string(input[idx-1][lrange:rrange+1]), lrange, rrange+1, string(line[lrange:rrange+1]), string(input[idx+1][lrange:rrange+1]))
			}
			if mygear > 0 {
				sum += mygear
			}
		}
	}
	fmt.Println("Gear Sum: ", sum)

}

func main() {
	var inputfile string
	var exercise1 bool
	flag.StringVar(&inputfile, "input", "input.txt", "Input file")
	flag.BoolVar(&exercise1, "exercise1", false, "Exercise 1 is True, Exercise 2 is False")
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
		charlong := scanner.Text()
		if width != 0 && width < len(charlong) {
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
		Exercise2(myrune)
	}
}
