package Day04

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part01(path string) (string, error) {
	scanner := readFile(path)

	total := 0
	cardMap := map[string]int{}

	for scanner.Scan() {
		re := regexp.MustCompile(`\d+`)

		card := strings.Split(strings.TrimSpace(scanner.Text()), ":")
		x := strings.Split(strings.TrimSpace(card[1]), "|")
		winningNumbers := re.FindAllString(x[0], -1)
		myNumbers := re.FindAllString(x[1], -1)

		for i := 0; i < len(winningNumbers); i++ {
			cardMap[winningNumbers[i]] += 1
		}

		for i := 0; i < len(myNumbers); i++ {
			cardMap[myNumbers[i]] += 1
		}

		score := 0
		for _, value := range cardMap {
			if value == 2 {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		total += score
		cardMap = map[string]int{}
	}

	result := strconv.Itoa(total)

	return result, nil
}

func Part02(path string) (string, error) {
	scanner := readFile(path)

	total := 0
	index := 0
	copies := map[int]int{}

	for scanner.Scan() {
		scoreMap := map[string]int{}
		re := regexp.MustCompile(`\d+`)

		card := strings.Split(strings.TrimSpace(scanner.Text()), ":")
		x := strings.Split(strings.TrimSpace(card[1]), "|")

		winningNumbers := re.FindAllString(x[0], -1)
		myNumbers := re.FindAllString(x[1], -1)

		for i := 0; i < len(winningNumbers); i++ {
			scoreMap[winningNumbers[i]] += 1
		}

		for i := 0; i < len(myNumbers); i++ {
			scoreMap[myNumbers[i]] += 1
		}

		matches := 0
		for _, value := range scoreMap {
			if value == 2 {
				matches++
			}
		}

		copies[index]++
		total++
		for j := 1; j <= matches; j++ {
			copies[index+j] += copies[index]
			total += copies[index]
		}
		index++
	}

	result := strconv.Itoa(total)

	return result, nil
}

func readFile(path string) *bufio.Scanner {
	dat, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(dat)

	return scanner
}
