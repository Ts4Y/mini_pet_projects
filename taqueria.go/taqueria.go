package taqueriago

import (
	"fmt"
)

func CountFood() {

	var total float32

	menu := map[string]float32{
		"Baja Taco":        4.00,
		"Burrito":          7.50,
		"Bowl":             8.50,
		"Nachos":           11.00,
		"Quesadilla":       8.50,
		"Super Burrito":    8.50,
		"Super Quesadilla": 9.50,
		"Taco":             3.00,
		"Tortilla Salad":   8.00,
	}

	fmt.Println("Добро пожаловать в Taqueria! Вы можете размещать заказ, вводя названия товаров. Для завершения заказа введите 'done'.")

	for {
		var EnterFood string
		fmt.Scan(&EnterFood)

		if EnterFood == "done" {
			break // Завершаем цикл, если введено "done"
		}

		if value, exist := menu[EnterFood]; exist {
			total += value
		} else {
			fmt.Println("Извините, этого блюда нет в меню.")
		}
	}

	fmt.Println("OBSHAYA SUMMA:", total)

}
