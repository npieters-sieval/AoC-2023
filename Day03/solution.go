package Day03

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Part01(path string) (string, error) {
	reNumbers := regexp.MustCompile(`\d+`)
	schematic := readFile(path)
	total := 0

	for i := 0; i < len(schematic); i++ {
		line := strings.Split(schematic[i], "")

		searchIndices := [][]int{}
		currIndex := []int{}

		for index, char := range line {
			matched := reNumbers.MatchString(char)
			if matched {
				currIndex = append(currIndex, index)
			}

			if !matched || index == len(line)-1 {
				if len(currIndex) > 0 {
					if currIndex[0]-1 >= 0 {
						currIndex = append([]int{currIndex[0] - 1}, currIndex...)
					}
					currIndex = append(currIndex, index)

					searchIndices = append(searchIndices, currIndex)
					currIndex = []int{}
				}

			}
		}

		min := []string{}
		plus := []string{}

		if i > 0 {
			min = strings.Split(schematic[i-1], "")
		}
		if i < len(schematic)-1 {
			plus = strings.Split(schematic[i+1], "")
		}

		for i := 0; i < len(searchIndices); i++ {
			addPart := false
			for _, val := range searchIndices[i] {
				if len(min) > 0 && checkIfCharacter(min[val]) {
					addPart = true
					break
				}

				if len(plus) > 0 && checkIfCharacter(plus[val]) {
					addPart = true
					break
				}

				if checkIfCharacter(line[val]) {
					addPart = true
					break
				}
			}

			if addPart {
				combinedString := strings.Join(line[searchIndices[i][0]:searchIndices[i][len(searchIndices[i])-1]+1], "")

				val, _ := strconv.Atoi(strings.Map(func(r rune) rune {
					if unicode.IsDigit(r) {
						return r
					}
					return -1
				}, combinedString))

				total = total + val
			}
		}
	}

	return strconv.Itoa(total), nil
}

func Part02(path string) (string, error) {
	schematic := readFile(path)
	total := 0

	for i := 0; i < len(schematic); i++ {
		line := strings.Split(schematic[i], "")

		gearIndices := [][]int{}
		currIndex := []int{}

		for index, char := range line {
			matched := char == "*"
			if matched {
				currIndex = append(currIndex, index)
			}

			if !matched || index == len(line)-1 {
				if len(currIndex) > 0 {
					gearIndices = append(gearIndices, currIndex)
					currIndex = []int{}
				}

			}
		}

		min := []string{}
		plus := []string{}

		if i > 0 {
			min = strings.Split(schematic[i-1], "")
		}
		if i < len(schematic)-1 {
			plus = strings.Split(schematic[i+1], "")
		}

		for i := 0; i < len(gearIndices); i++ {
			numbers := []int{}
			start := 0
			end := len(line)

			for _, val := range gearIndices[i] {

				starIndex := val
				n := extractNumbersAtIndex(strings.Join(line[start:end], ""), starIndex)
				numbers = append(n, numbers...)

				if len(min) > 0 {
					s := strings.Join(min, "")
					n = extractNumbersAtIndex(s, starIndex)
					numbers = append(n, numbers...)

				}

				if len(plus) > 0 {
					n = extractNumbersAtIndex(strings.Join(plus, ""), starIndex)
					numbers = append(n, numbers...)
				}
			}

			if len(numbers) == 2 {
				v1 := numbers[0]
				v2 := numbers[1]
				fmt.Println(v1, v2, total)
				total = total + (v1 * v2)

			}
		}

	}
	fmt.Println(total)
	return strconv.Itoa(total), nil
}

func readFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Cant find file")
	}
	return strings.Split(string(content), "\n")
}

func checkIfCharacter(input string) bool {
	return input != "." && input != "0" && input != "1" && input != "2" && input != "3" && input != "4" && input != "5" && input != "6" && input != "7" && input != "8" && input != "9"
}

func checkIfNumber(input string) bool {
	return input == "0" || input == "1" || input == "2" || input == "3" || input == "4" || input == "5" || input == "6" || input == "7" || input == "8" || input == "9"
}

func extractNumbersAtIndex(s string, index int) []int {
	var numbers []int
	numStr := ""
	startIndex := -1
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			if numStr == "" {
				startIndex = i
			}
			numStr += string(s[i])
		} else if numStr != "" {
			if (startIndex <= index+1 && startIndex >= index-1) || (i-1 <= index+1 && i-1 >= index-1) {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
			numStr = ""
		}
	}
	if numStr != "" && (startIndex <= index+1 && startIndex >= index-1) {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}
	return numbers
}

func findStarIndex(arr []string) int {
	for i, str := range arr {
		if str == "*" {
			return i
		}
	}
	return -1
}
