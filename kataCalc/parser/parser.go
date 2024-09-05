package parser

import (
	"strconv"
	"strings"
)

func Split(input string) []string {
	var op string
	arr := []string{}
	for {
		arr = strings.Split(input, " - ")
		if len(arr) == 2 {
			op = "-"
			break
		}
		arr = strings.Split(input, " + ")
		if len(arr) == 2 {
			op = "+"
			break
		}
		arr = strings.Split(input, " / ")
		if len(arr) == 2 {
			op = "/"
			break
		}
		arr = strings.Split(input, " * ")
		if len(arr) == 2 {
			op = "*"
			break
		}
		panic("Неправильный оператор или некорректное число входных значений")
	}

	_, err := strconv.Atoi(arr[0])
	if err == nil {
		panic("Числовое значение не может стоять в начале строки!")
	}

	g, err := strconv.Atoi(arr[1])
	if (op == "/" || op == "*") && err != nil {
		panic("Операции деления и умножения возможно проводить только с использованием целых чисел")
	}
	if (op == "/" || op == "*") && g > 10 {
		panic("Числовое значение не может быть больше 10")
	}
	if (op == "+" || op == "-") && ((len(strings.Trim(arr[0], `"`))+2 != len(arr[0])) || (len(strings.Trim(arr[1], `"`))+2 != len(arr[1]))) {
		panic("В операциях сложения и вычитания могут участвовать только строки окружённые кавычками!")
	}
	if (op == "+" || op == "-") && (len(arr[0]) > 12 || len(arr[1]) > 12) {
		panic("Длина одной строки не может быть больше 10 символов")
	}

	return []string{arr[0], op, arr[1]}
}