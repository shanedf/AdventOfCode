package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	card := 0
	var cards [198][]string
	var draws [198][]string
	var matches [198][]string
	total := 0
	for scanner.Scan() {
		inputLn := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")
		cards[card] = re.FindAllString(inputLn[1], -1)
		draws[card] = re.FindAllString(inputLn[0], -1)
		// var intCard = []int{}
		// for _, c := range strCard {
		// 	v, err := strconv.Atoi(c)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	intCard = append(intCard, v)
		// }
		// cards[card] = intCard
		// var intDraw = []int{}
		// for _, d := range strDraw {
		// 	v, err := strconv.Atoi(d)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	intDraw = append(intDraw, v)
		// }
		// draws[card] = intDraw
		for _, v := range draws[card] {
			if slices.Contains(cards[card], v) {
				matches[card] = append(matches[card], v)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		card++
	}
}
