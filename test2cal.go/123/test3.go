package main

import (
	"fmt"
	"strings"
)

func main() {
	// Ввод данных с консоли
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	// Разделение строки по пробелам
	splitted := strings.Split(input, " ")
	if len(splitted) != 3 {
		fmt.Println("Ошибка: некорректное количество аргументов")
		return
	}

	// Парсинг чисел
	a, err := parseNumber(splitted[0])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	b, err := parseNumber(splitted[2])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Выполнение операции
	var result interface{}
	switch operator := splitted[1]; operator {
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
		fmt.Println("Ошибка: некорректная операция")
		return
	}

	// Вывод результата
	fmt.Println(result)
}

func parseNumber(s string) (int, error) {
	// Проверка является ли число арабским
	if isArabic(s) {
		return parseArabic(s)
	}

	// Проверка является ли число римским
	if isRoman(s) {
		return parseRoman(s)
	}

	return 0, fmt.Errorf("некорректное число")
}

func isArabic(s string) bool {
	_, err := fmt.Sscan(s, new(int))
	return err == nil
}

func parseArabic(s string) (int, error) {
	num, err := fmt.Sscan(s, new(int))
	if err != nil || num != 1 || !isInRange(s, 1, 10) {
		return 0, fmt.Errorf("некорректное арабское число")
	}
	return num, nil
}

func isRoman(s string) bool {
	for _, r := range s {
		if r != 'I' && r != 'V' && r != 'X' {
			return false
		}
	}
	return true
}

func parseRoman(s string) (int, error) {
	num, err := romanToInt(s)
	if err != nil || !isInRange(s, 1, 10) {
		return 0, fmt.Errorf("некорректное римское число")
	}
	return num, nil
}

func isInRange(s string, min, max int) bool {
	num, _ := romanToInt(s)
	return num >= min && num <= max
}

func romanToInt(s string) (int, error) {
	romanDict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	num := 0
	for i := 0; i < len(s); i++ {
		// Проверка валидности символов
		_, ok := romanDict[rune(s[i])]
		if !ok {
			return 0, fmt.Errorf("некорректное римское число")
		}

		// Проверка на правильность порядка символов
		if i > 0 && romanDict[rune(s[i])] > romanDict[rune(s[i-1])] {
			return 0, fmt.Errorf("некорректное римское число")
		}

		// Вычитание или сложение чисел в зависимости от порядка
		if i > 0 && romanDict[rune(s[i])] > romanDict[rune(s[i-1])] {
			num += romanDict[rune(s[i])] - 2*romanDict[rune(s[i-1])]
		} else {
			num += romanDict[rune(s[i])]
		}
	}

	return num, nil
}
