package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Task2() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	gameSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		rounds := strings.Split(strings.Split(input, ":")[1], ";")
		minRed := 0
		minGreen := 0
		minBlue := 0
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
					if int(r) > minRed {
						minRed = int(r)
					}
				case "green":
					r, err = strconv.ParseInt(res[0], 10, 8)
					if err != nil {
						panic(err)
					}
					if int(r) > minGreen {
						minGreen = int(r)
					}
				case "blue":
					r, err = strconv.ParseInt(res[0], 10, 8)
					if err != nil {
						panic(err)
					}
					if int(r) > minBlue {
						minBlue = int(r)
					}
				}
			}
		}
		gameSum = gameSum + (minRed * minGreen * minBlue)
	}
	fmt.Println("Total is:", gameSum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
