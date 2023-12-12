package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(operand1, operand2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("Неподдерживаемая операция: %s", operator)
	}
}

func isRoman(s string) bool {
	romanNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	_, err := strconv.Atoi(s)
	return err != nil && romanNumerals[s] > 0
}

func toArabic(s string) (int, error) {
	if isRoman(s) {
		romanNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
		return romanNumerals[s], nil
	}
	return strconv.Atoi(s)
}

func toRoman(num int) (string, error) {
	if num <= 0 {
		return "", fmt.Errorf("Результатом работы калькулятора с римскими числами могут быть только положительные числа.")
	}
	romanNumerals := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}

	result := ""
	for value, numeral := range romanNumerals {
		for num >= value {
			result += numeral
			num -= value
		}
	}
	return result, nil
}

func runCalculator(expression string) (string, error) {
	parts := strings.Fields(expression)

	if len(parts) != 3 {
		return "", fmt.Errorf("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	operand1Str, operator, operand2Str := parts[0], parts[1], parts[2]

	operand1, err := toArabic(operand1Str)
	if err != nil {
		return "", err
	}

	operand2, err := toArabic(operand2Str)
	if err != nil {
		return "", err
	}

	if isRoman(operand1Str) != isRoman(operand2Str) {
		return "", fmt.Errorf("Используются одновременно разные системы счисления.")
	}

	result, err := calculate(operand1, operand2, operator)
	if err != nil {
		return "", err
	}

	if isRoman(operand1Str) {
		romanResult, err := toRoman(result)
		if err != nil {
			return "", err
		}
		return romanResult, nil
	}

	return strconv.Itoa(result), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := runCalculator(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
