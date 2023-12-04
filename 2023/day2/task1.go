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
	var gameSum int64
	limitRed := 12
	limitGreen := 13
	limitBlue := 14
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		var r int64
		sumRed := 0
		sumGreen := 0
		sumBlue := 0
		game := strings.Split(input, ":")[0]
		gameId := strings.Fields(game)[1]
		rounds := strings.Split(game, ";")
		for _, round := range rounds {
			draw := strings.Split(round, ",")
			for _, clr := range draw {
				res := strings.Fields(strings.TrimSpace(clr))
				switch res[1] {
				case "red":
					r, err = strconv.ParseInt(res[0], 10, 8)
				}
				if err != nil {
					panic(err)
				}
			}
		}
	}
	fmt.Println("Total is:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
