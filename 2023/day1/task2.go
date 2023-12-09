package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Task2() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var firstIdx int
	var firstVal string
	var lastIdx int
	var lastVal string
	var sum int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLn := scanner.Text()
		firstIdx = len(inputLn)
		lastIdx = -1
		for _, val := range nums {
			if (strings.Index(inputLn, val) < firstIdx) && (strings.Index(inputLn, val) > -1) {
				// fmt.Println("Found", val, "at position", idx)
				firstIdx = strings.Index(inputLn, val)
				firstVal = val
			}
			if strings.LastIndex(inputLn, val) > lastIdx {
				// fmt.Println("Found", val, "at position", idx)
				lastIdx = strings.LastIndex(inputLn, val)
				lastVal = val
			}
		}
		if len(firstVal) > 1 {
			firstVal = getNumber(firstVal)
		}
		if len(lastVal) > 1 {
			lastVal = getNumber(lastVal)
		}
		if firstVal == "err" || lastVal == "err" {
			panic("one or more numbers failed to convert correctly")
		}
		tmp := firstVal + lastVal
		fmt.Println(scanner.Text(), len(inputLn), firstVal, lastVal, tmp)
		num, err := strconv.ParseInt(tmp, 10, 8)
		if err != nil {
			panic(err)
		}
		sum = sum + num
	}
	fmt.Println("Total is:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumber(num string) string {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	vals := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, n := range nums {
		if n == num {
			return vals[i]
		}
	}
	return "err"
}
