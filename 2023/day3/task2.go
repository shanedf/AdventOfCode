package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+")
	as := regexp.MustCompile("\\*")
	//vals := "0123456789"
	scanner := bufio.NewScanner(file)
	var line int = 0
	var lines [140]string
	var astxs [140][][]int
	var nums [140][]string
	var idxs [140][][]int
	sumParts := 0
	for scanner.Scan() {
		inputLn := scanner.Text()
		lines[line] = inputLn
		astxs[line] = as.FindAllStringIndex(inputLn, -1)
		nums[line] = re.FindAllString(inputLn, -1)
		idxs[line] = re.FindAllStringIndex(inputLn, -1)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line++
	}
	for ln := 0; ln < 140; ln++ {
		fmt.Println("Line", ln)
		for idx, val := range astxs[ln] {
			fmt.Println("Asterisk", idx)
			// part := ""
			prev := 0
			curr := 0
			next := 0
			num1 := ""
			num2 := ""
			//var locs [2][]int
			// flag := false
			var min int
			if val[0] > 0 {
				min = val[0] - 1
			} else {
				min = 0
			}
			var max int
			if val[1] < 139 {
				max = val[1] + 1
			} else {
				max = 139
			}
			if ln > 0 {
				prev = len(re.FindAllString(lines[ln-1][min:max], -1))
				if prev > 0 {
					for i, v := range idxs[ln-1] {
						fmt.Println("Checking previous row if", min, max, "are present in", v, "-", nums[ln-1][i])
						if (min >= v[0] && min < v[1]) || (max > v[0] && max <= v[1]) {
							fmt.Println("Found", nums[ln-1][i], "in previous row")
							if num1 == "" {
								num1 = nums[ln-1][i]
							} else {
								if num2 == "" {
									num2 = nums[ln-1][i]
								} else {
									panic("More than two values attempted!")
								}
							}
						}
					}
				}
			}
			curr = len(re.FindAllString(lines[ln][min:max], -1))
			if curr > 0 {
				for i, v := range idxs[ln] {
					fmt.Println("Checking current row if", min, max, "are present in", v, "-", nums[ln][i])
					if (min >= v[0] && min < v[1]) || (max >= v[0] && max <= v[1]) || (v[0] >= min && v[1] <= max) {
						fmt.Println("Found", nums[ln][i], "on current line")
						if num1 == "" {
							num1 = nums[ln][i]
						} else {
							if num2 == "" {
								num2 = nums[ln][i]
							} else {
								panic("More than two values attempted!")
							}
						}
					}
				}
			}
			if ln < 139 {
				next = len(re.FindAllString(lines[ln+1][min:max], -1))
				if next > 0 {
					for i, v := range idxs[ln+1] {
						fmt.Println("Checking next row if", min, max, "are present in", v, "-", nums[ln+1][i])
						if (min >= v[0] && min < v[1]) || (max > v[0] && max < v[1]) || (v[0] >= min && v[1] <= max) {
							fmt.Println("Found", nums[ln+1][i], "in next row")
							if num1 == "" {
								num1 = nums[ln+1][i]
							} else {
								if num2 == "" {
									num2 = nums[ln+1][i]
								} else {
									panic("More than two values attempted!")
								}
							}
						}
					}
				}
			}
			if prev+curr+next == 2 {
				fmt.Println("Line", ln, ", asterisk", idx, "at", val, "has 2 adjacent numbers -", prev, "on the previous line,", curr, "on the same line and ", next, "on the following line")
				var nm1 int64
				var nm2 int64
				nm1, err = strconv.ParseInt(num1, 10, 16)
				if err != nil {
					panic(err)
				}
				nm2, err = strconv.ParseInt(num2, 10, 16)
				if err != nil {
					panic(err)
				}
				ratio := int(nm1) * int(nm2)
				//fmt.Println("Gear ratio found for line", ln, "asterisk", idx, "is", num1, num2, "result:", ratio)
				sumParts += ratio
			}
		}
	}
	fmt.Println("Total sun for parts:", sumParts)
}
