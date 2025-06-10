package main

import "fmt"

func main() {
	var weight, height float64

	fmt.Print("Введите вес (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Введите рост (м): ")
	fmt.Scan(&height)

	bmi := weight / (height * height)
	fmt.Printf("Ваш ИМТ: %.2f\n", bmi)
}
