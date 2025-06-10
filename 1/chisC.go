package main

import "fmt"

func main() {
	var x, y int
	var operation string

	fmt.Print("Введите число: ")
	fmt.Scan(&x)

	fmt.Print("Введите число: ")
	fmt.Scan(&y)

	fmt.Print("Введите оперцию (+, -, *, /, %, f): ")
	fmt.Scan(&operation)

	if y == 0 && (operation == "/" || operation == "%" || operation == "f") {
		fmt.Println("Ошибка: деление на ноль!")
		return
	}

	switch operation {
	case "+":
		sum := x + y
		fmt.Printf("Сумма: %d\n", sum)

	case "-":
		diff := x - y
		fmt.Printf("Разность: %d\n", diff)

	case "*":
		prod := x * y
		fmt.Printf("Произведение: %d\n", prod)

	case "/":
		qu := x / y
		fmt.Printf("Целочисленное деление: %d\n", qu)

	case "%":
		rm := x % y
		fmt.Printf("Остаток от деления: %d\n", rm)

	case "f":
		floatDiv := float64(x) / float64(y)
		fmt.Printf("Дробное деление: %.2f\n", floatDiv)

	default:
		fmt.Println("Некорректный ввод")
	}

	/*
		fmt.Printf("Сумма: %d\n", sum)
		fmt.Printf("Разность: %d\n", diff)
		fmt.Printf("Произведение: %d\n", prod)
		fmt.Printf("Целочисленное деление: %d\n", qu)
		fmt.Printf("Остаток от деления: %d\n", rm)
		fmt.Printf("Дробное деление: %.2f\n", floatDiv)
	*/
}
