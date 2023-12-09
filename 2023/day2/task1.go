package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Task1() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	gameSum := 0
	limitRed := 12
	limitGreen := 13
	limitBlue := 14
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		game := strings.Split(input, ":")[0]
		gameId := strings.Fields(game)[1]
		rounds := strings.Split(strings.Split(input, ":")[1], ";")
		trip := false
		for _, round := range rounds {
			draw := strings.Split(round, ",")
			for _, clr := range draw {
				var r int64
				res := strings.Fields(strings.TrimSpace(clr))
				switch res[1] {
				case "red":
					r, err = strconv.ParseInt(res[0], 10, 8)
					if err != nil {
						panic(err)
					}
					if int(r) > limitRed {
						trip = true
					}
				case "green":
					r, err = strconv.ParseInt(res[0], 10, 8)
					if err != nil {
						panic(err)
					}
					if int(r) > limitGreen {
						trip = true
					}
				case "blue":
					r, err = strconv.ParseInt(res[0], 10, 8)
					if err != nil {
						panic(err)
					}
					if int(r) > limitBlue {
						trip = true
					}
				}
			}

		}
		if trip == false {
			fmt.Println("Game", gameId, "is possible")
			var id int64
			id, err = strconv.ParseInt(gameId, 10, 8)
			if err != nil {
				panic(err)
			}
			gameSum = gameSum + int(id)
		}
	}
	fmt.Println("Total is:", gameSum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
