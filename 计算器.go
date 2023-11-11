package main

import "fmt"

type operation func(float64, float64) float64

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	return 0
}

func calculate(x, y float64, op operation) float64 {
	return op(x, y)
}

func main() {
	var x, y float64
	var operator string

	fmt.Println("第一个数字： ")
	fmt.Scan(&x)

	fmt.Println("第二个数字： ")
	fmt.Scan(&y)

	fmt.Println("运算符号 (+, -, *, /): ")
	fmt.Scan(&operator)

	var result float64

	switch operator {
	case "+":
		result = calculate(x, y, add)
	case "-":
		result = calculate(x, y, subtract)
	case "*":
		result = calculate(x, y, multiply)
	case "/":
		result = calculate(x, y, divide)
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("Result: %v\n", result)
}
