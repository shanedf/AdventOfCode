package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
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
	var cards = [198][]int
	var draws = [198][]int
	var matches = [198]int
	total := 0
	for scanner.Scan() {
		inputLn := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")
		strCard := re.FindAllString(inputLn[1])
		strDraw := re.FindAllString(inputLn[0])

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		card++
	}
}
