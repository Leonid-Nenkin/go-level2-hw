// Package main implements calculator
//
// The main runs calculator

package main

import (
	"fmt"
	"os"
)

// Start calculator
// Calculator supports operations like +, -, *, /

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	if b == 0 && op == "/" {
		fmt.Println("Деление на 0 запрещено")
		os.Exit(1)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
