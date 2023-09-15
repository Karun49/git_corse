package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Read user input
	var input string
	fmt.Print("Enter an arithmetic expression: ")
	fmt.Scanln(&input)

	// Parse the input
	numbers, operator, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Perform the calculation and handle errors
	result, err := calculate(numbers, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the result
	fmt.Println("Result:", result)
}

// Function to parse the user input
func parseInput(input string) ([]int, string, error) {
	// Split the input by spaces
	items := strings.Split(input, " ")
	if len(items) != 3 {
		return nil, "", fmt.Errorf("Invalid input")
	}

	// Parse the first and third items as integers
	numbers := make([]int, 2)
	for i, item := range items {
		num, err := parseNumber(item)
		if err != nil {
			return nil, "", err
		}
		numbers[i] = num
	}

	// Parse the second item as the operator
	operator := items[1]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return nil, "", fmt.Errorf("Invalid operator")
	}

	return numbers, operator, nil
}

// Function to parse a number (either arabic or roman)
func parseNumber(item string) (int, error) {
	// Try parsing as arabic number
	num, err := parseArabicNumber(item)
	if err == nil {
		return num, nil
	}

	// Try parsing as roman number
	num, err = parseRomanNumber(item)
	if err == nil {
		return num, nil
	}

	return 0, fmt.Errorf("Invalid number")
}

// Function to parse an arabic number
func parseArabicNumber(item string) (int, error) {
	num, err := strconv.Atoi(item)
	if err != nil || num < 1 || num > 10 {
		return 0, fmt.Errorf("Invalid arabic number")
	}
	return num, nil
}

// Function to parse a roman number
func parseRomanNumber(item string) (int, error) {
	// Roman numeral mapping
	romanMapping := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Iterate over the characters in reverse order
	runes := []rune(item)
	total := 0
	prev := 0
	for i := len(runes) - 1; i >= 0; i-- {
		num := romanMapping[runes[i]]
		if num < prev {
			total -= num
		} else {
			total += num
		}
		prev = num
	}

	// Check for invalid roman numeral
	if convertToRomanNumeral(total) != item {
		return 0, fmt.Errorf("Invalid roman number")
	}

	return total, nil
}

// Function to convert an arabic number to roman numeral
func convertToRomanNumeral(num int) string {
	romanMapping := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	result := ""
	for num > 0 {
		for arabic, roman := range romanMapping {
			if num >= arabic {
				result += roman
				num -= arabic
				break
			}
		}
	}

	return result
}

// Function to calculate the result of the arithmetic expression
func calculate(numbers []int, operator string) (int, error) {
	a, b := numbers[0], numbers[1]

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Invalid operator")
	}
}
