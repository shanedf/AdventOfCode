package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
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
		for idx, val := range nums {
			if (strings.Index(inputLn, val) < firstIdx) && (strings.Index(inputLn, val) > -1) {
				fmt.Println("Found", val, "at position", idx)
				firstIdx = strings.Index(inputLn, val)
				firstVal = val
			}
			if strings.LastIndex(inputLn, val) > lastIdx {
				fmt.Println("Found", val, "at position", idx)
				lastIdx = strings.LastIndex(inputLn, val)
				lastVal = val
			}
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
