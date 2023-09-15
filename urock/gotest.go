package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Введите выражение:")

	var input string
	fmt.Scanln(&input)

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculate(input string) (int, error) {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		return 0, fmt.Errorf("некорректное выражение")
	}

	num1, err := convertToNumber(parts[0])
	if err != nil {
		return 0, fmt.Errorf("некорректное число: %v", parts[0])
	}

	num2, err := convertToNumber(parts[2])
	if err != nil {
		return 0, fmt.Errorf("некорректное число: %v", parts[2])
	}

	operator := parts[1]
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("некорректный оператор: %v", operator)
	}
}

func convertToNumber(input string) (int, error) {
	romanToNumber := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if val, ok := romanToNumber[input]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("некорректное число: %v", input)
}
