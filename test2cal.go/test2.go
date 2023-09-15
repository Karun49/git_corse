package main

import (
	"fmt"
	"strings"
)

func main() {
	// Считываем пользовательский ввод
	var input string
	fmt.Print("Введите арифметическое выражение: ")
	fmt.Scanln(&input)

	// Разбиваем ввод на числа и операцию
	values := strings.Split(input, " ")
	if len(values) != 3 {
		fmt.Println("Ошибка: неверное выражение")
		return
	}

	// Парсим числа
	a, err := parseNumber(values[0])
	if err != nil {
		fmt.Println("Ошибка: неверное число")
		return
	}
	b, err := parseNumber(values[2])
	if err != nil {
		fmt.Println("Ошибка: неверное число")
		return
	}

	// Выполняем операцию
	var result int
	switch values[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неверная операция")
		return
	}

	// Выводим результат
	fmt.Printf("Результат: %s\n", formatNumber(result))
}

// Функция для парсинга числа
func parseNumber(s string) (int, error) {
	// Проверяем, является ли число римским
	if isRomanNumeral(s) {
		return romanToArabic(s)
	}

	// Парсим арабское число
	var result int
	_, err := fmt.Sscan(s, &result)
	return result, err
}

// Функция для форматирования числа
func formatNumber(n int) string {
	// Проверяем, является ли число отрицательным
	if n < 0 {
		return fmt.Sprintf("-%s", formatNumber(-n))
	}

	// Форматируем арабское число
	return fmt.Sprintf("%d", n)
}

// Функция для проверки, является ли строка римским числом
func isRomanNumeral(s string) bool {
	// Проверяем, что строка состоит только из символов I, V, X, L, C, D, M
	allowedChars := "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	return true
}

// Функция для преобразования римского числа в арабское
func romanToArabic(s string) (int, error) {
	// Создаем словарь с соответствием римских символов и чисел
	romanToArabicMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Преобразуем римское число в арабское, выполняя сложение или вычитание
	var result int
	prev := 0
	for _, char := range s {
		value := romanToArabicMap[char]
		if value == 0 {
			return 0, fmt.Errorf("неверное римское число: %s", s)
		}
		if prev < value {
			result -= prev
		} else {
			result += prev
		}
		prev = value
	}
	result += prev
	return result, nil
}
