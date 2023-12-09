package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func Task1() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	card := 0
	var scores = [11]int{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	var cards [198][]string
	var draws [198][]string
	var matches [198][]string
	total := 0
	for scanner.Scan() {
		inputLn := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")
		cards[card] = re.FindAllString(inputLn[1], -1)
		draws[card] = re.FindAllString(inputLn[0], -1)
		for _, v := range draws[card] {
			if slices.Contains(cards[card], v) {
				matches[card] = append(matches[card], v)
			}
		}
		matched := len(matches[card])
		score := scores[matched]
		total += int(score)
		fmt.Println("there were", matched, "matches on card", card+1, "scoring", score, "points")
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		card++
		fmt.Println("Total score is", total)
	}
}
