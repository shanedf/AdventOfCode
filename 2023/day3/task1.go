package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+")
	specials := "/*-+!£$@~#;:<>=%^_&"
	scanner := bufio.NewScanner(file)
	var line int = 0
	var lines [140]string
	var nums [140][]string
	var idxs [140][][]int
	sumParts := 0
	for scanner.Scan() {
		inputLn := scanner.Text()
		lines[line] = inputLn
		nums[line] = re.FindAllString(inputLn, -1)
		idxs[line] = re.FindAllStringIndex(inputLn, -1)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line++
	}
	for ln := 0; ln < 140; ln++ {
		fmt.Println("Line", ln, lines[ln], nums[ln], idxs[ln])
		for idx, val := range idxs[ln] {
			part := ""
			flag := false
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
				if strings.ContainsAny(lines[ln-1][min:max], specials) {
					part = nums[ln][idx]
					flag = true
					fmt.Println("Part", part)
				}
			}
			if strings.ContainsAny(lines[ln][min:max], specials) {
				part = nums[ln][idx]
				flag = true
				fmt.Println("Part", part)
			}
			if ln < 139 {
				if strings.ContainsAny(lines[ln+1][min:max], specials) {
					part = nums[ln][idx]
					flag = true
					fmt.Println("Part", part)
				}
			}
			if flag {
				var pt int64
				pt, err = strconv.ParseInt(part, 10, 16)
				if err != nil {
					panic(err)
				}
				sumParts += int(pt)
			}
		}
	}
	fmt.Println("Total sun for parts:", sumParts)
}
