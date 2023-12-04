package Day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Part01(path string) (string, error) {
	scanner := readFile(path)

	total := 0
	for scanner.Scan() {
		pattern := regexp.MustCompile(`\d`)
		matches := pattern.FindAllString(scanner.Text(), -1)

		digit, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		fmt.Println(digit)
		total = total + digit
	}

	result := strconv.Itoa(total)

	return result, nil
}

func Part02(path string) (string, error) {
	scanner := readFile(path)
	numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	var total int
	for scanner.Scan() {
		line := scanner.Text()
		digits := []rune{}
		match := []int{}

		re := regexp.MustCompile(`^one|^two|^three|^four|^five|^six|^seven|^eight|^nine`)
		for k, c := range line {

			lineLength := len(line)
			v := c - 48
			if v < 10 {
				digits = append(digits, v)
				continue
			}

			if lineLength-k >= 5 {
				match = re.FindStringIndex(line[k : k+5])
			} else {
				match = re.FindStringIndex(line[k:])
			}
			if match == nil {
				continue
			}

			digits = append(digits, rune(numbers[line[k+match[0]:k+match[1]]]))

		}
		total += int(digits[0]*10 + digits[len(digits)-1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
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
