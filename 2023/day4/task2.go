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

func Task2() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	card := 0
	var cards [198][]string
	var draws [198][]string
	var matches [198]int
	var cardCounts [198]int
	for scanner.Scan() {
		inputLn := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")
		cards[card] = re.FindAllString(inputLn[1], -1)
		draws[card] = re.FindAllString(inputLn[0], -1)
		var matched []string
		for _, v := range draws[card] {
			if slices.Contains(cards[card], v) {
				matched = append(matched, v)
			}
		}
		matches[card] = len(matched)
		cardCounts[card] = 1
		card++
	}

	for i, v := range matches {
		for j := 0; j < cardCounts[i]; j++ {
			for k := 1; k <= v; k++ {
				cardCounts[i+k] += 1
			}
		}
	}

	total := 0
	for i := range cardCounts {
		total += cardCounts[i]
	}

	fmt.Println(total)
}
