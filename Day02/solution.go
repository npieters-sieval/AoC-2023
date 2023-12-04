package Day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part01(path string) (string, error) {
	scanner := readFile(path)

	cubeMap := map[string]int{"red": 0, "green": 0, "blue": 0}
	total := 0
	for scanner.Scan() {
		game := strings.Split(strings.TrimSpace(scanner.Text()), ":")
		sets := strings.Split(strings.TrimSpace(game[1]), ";")

		currGame := strings.Split(game[0], " ")
		gameId, _ := strconv.Atoi(currGame[1])
		total += gameId

		for _, hands := range sets {
			cubeMap["red"] = 0
			cubeMap["green"] = 0
			cubeMap["blue"] = 0

			cubes := strings.Split(hands, ",")

			for _, cube := range cubes {
				x := strings.Split(strings.TrimSpace(cube), " ")
				amount, _ := strconv.Atoi(strings.TrimSpace(x[0]))
				cubeMap[strings.TrimSpace(x[1])] += amount
			}

			if cubeMap["red"] > 12 || cubeMap["green"] > 13 || cubeMap["blue"] > 14 {
				total = total - gameId
				break
			}
		}
	}

	result := strconv.Itoa(total)

	return result, nil
}

func Part02(path string) (string, error) {
	scanner := readFile(path)

	total := 0
	for scanner.Scan() {
		game := strings.Split(strings.TrimSpace(scanner.Text()), ":")
		sets := strings.Split(strings.TrimSpace(game[1]), ";")
		cubeMap := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, hands := range sets {
			cubes := strings.Split(hands, ",")

			for _, cube := range cubes {
				x := strings.Split(strings.TrimSpace(cube), " ")
				amount, _ := strconv.Atoi(strings.TrimSpace(x[0]))

				if amount > cubeMap[x[1]] {
					cubeMap[x[1]] = amount
				}
			}
		}

		multplied := cubeMap["red"] * cubeMap["green"] * cubeMap["blue"]
		total += multplied
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
