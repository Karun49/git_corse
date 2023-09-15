package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Калькулятор")

	for {
		fmt.Print("Введите выражение (два числа и операцию, разделенные пробелами): ")
		input, _ := reader.ReadString('\n')

		// Удаление символа новой строки в конце строки
		input = strings.TrimSpace(input)

		// Разделение строки на числа и операцию
		inputSlice := strings.Split(input, " ")
		if len(inputSlice) != 3 {
			fmt.Println("Ошибка! Пожалуйста, введите два числа и операцию, разделенные пробелами.")
			continue
		}

		// Парсинг первого числа
		num1, err := strconv.Atoi(inputSlice[0])
		if err != nil {
			fmt.Println("Ошибка! Пожалуйста, введите числа от 1 до 10 включительно.")
			continue
		}

		// Парсинг второго числа
		num2, err := strconv.Atoi(inputSlice[2])
		if err != nil {
			fmt.Println("Ошибка! Пожалуйста, введите числа от 1 до 10 включительно.")
			continue
		}

		// Проверка операции
		result := 0
		//	isRoman := false
		operator := inputSlice[1]
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		default:
			fmt.Println("Ошибка! Поддерживаемые операции: +, -, *, /")
			continue
		}

		// Вывод результата
		if strings.ContainsAny(inputSlice[0], "IVX") && strings.ContainsAny(inputSlice[2], "IVX") {
			//	isRoman = true
			romanResult := intToRoman(result)
			fmt.Printf("Результат: %s\n", romanResult)
		} else {
			fmt.Printf("Результат: %d\n", result)
		}

		// Завершение программы при вводе "exit"
		if strings.ToLower(input) == "exit" {
			break
		}
	}
}

func intToRoman(num int) string {
	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for _, numeral := range romanNumerals {
		for num >= numeral.value {
			result += numeral.symbol
			num -= numeral.value
		}
	}

	return result
}
