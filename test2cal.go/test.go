package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Функция для проверки, является ли число римским
func isRomanNumeral(numeral string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, rn := range romanNumerals {
		if numeral == rn {
			return true
		}
	}
	return false
}

// Функция для преобразования римских чисел в арабские
func convertRomanToArabic(numeral string) int {
	romanToArabic := map[string]int{
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
	return romanToArabic[numeral]
}

// Функция для преобразования арабских чисел в римские
func convertArabicToRoman(number int) string {
	arabicToRoman := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	return arabicToRoman[number]
}

// Функция для выполнения операции сложения
func add(a, b int, isRoman bool) int {
	result := a + b
	if isRoman {
		return result
	} else {
		return result
	}
}

// Функция для выполнения операции вычитания
func subtract(a, b int, isRoman bool) int {
	result := a - b
	if isRoman {
		return result
	} else {
		return result
	}
}

// Функция для выполнения операции умножения
func multiply(a, b int, isRoman bool) int {
	result := a * b
	if isRoman {
		return result
	} else {
		return result
	}
}

// Функция для выполнения операции деления
func divide(a, b int, isRoman bool) int {
	result := a / b
	if isRoman {
		return result
	} else {
		return result
	}
}

func main() {
	for {
		var input string
		fmt.Print("Введите выражение: ")
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)
		// Разбиваем строку на операнды и оператор
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Ошибка: неправильный формат выражения")
			break
		}

		// Проверяем, являются ли операнды числами от 1 до 10
		a, isValid := convertRomanToArabic(parts[0]), false
		if a != 0 {
			isValid = true
		} else {
			a, isValid = convertRomanToArabic(parts[0]), false
			if a != 0 {
				isValid = true
			}
		}
		if !isValid {
			fmt.Println("Ошибка: операнды должны быть числами от 1 до 10")
			break
		}

		// Проверяем, является ли оператор допустимым
		operator := parts[1]
		if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			fmt.Println("Ошибка: неправильный оператор")
			break
		}

		// Выполняем операцию в зависимости от типа операнда и оператора
		b, _ := strconv.Atoi(parts[2])
		isRoman := isRomanNumeral(parts[0]) && isRomanNumeral(parts[2])
		switch operator {
		case "+":
			result := add(a, b, isRoman)
			if isRoman {
				fmt.Println(convertArabicToRoman(result))
			} else {
				fmt.Println(result)
			}
		case "-":
			result := subtract(a, b, isRoman)
			if isRoman {
				fmt.Println(convertArabicToRoman(result))
			} else {
				fmt.Println(result)
			}
		case "*":
			result := multiply(a, b, isRoman)
			if isRoman {
				fmt.Println(convertArabicToRoman(result))
			} else {
				fmt.Println(result)
			}
		case "/":
			if b == 0 {
				fmt.Println("Ошибка: деление на ноль")
				break
			}
			result := divide(a, b, isRoman)
			if isRoman {
				fmt.Println(convertArabicToRoman(result))
			} else {
				fmt.Println(result)
			}
		}
	}
}
