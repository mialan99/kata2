package main

import (
	"bufio"
	"fmt"
	"kata/operations"
	"kata/parser"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение: ")
	scanner.Scan()
	input := scanner.Text()
	arr := parser.Split(input)

	switch arr[1] {
	case "+":
		fmt.Println(operations.Sum(arr[0], arr[2]))
	case "-":
		fmt.Println(operations.Sub(arr[0], arr[2]))
	case "*":
		fmt.Println(operations.Multiply(arr[0], arr[2]))
	case "/":
		fmt.Println(operations.Div(arr[0], arr[2]))
	}
}
