package main

import "fmt"

func mapDict() {
	// Инициализация карты
	m := make(map[string]int)

	// Добавление значений
	m["apple"] = 5
	m["banana"] = 7

	// Получение значения
	appleCount, exists := m["apple"]
	if exists {
		fmt.Println("map:", m, "Количество яблок:", appleCount)
	} else {
		fmt.Println("map:", m, "Яблоки не найдены")
	}

	// Удаление элемента
	delete(m, "apple")
	appleCount, exists = m["apple"]
	if exists {
		fmt.Println("map:", m, "Количество яблок:", appleCount)
	} else {
		fmt.Println("map:", m, "Яблоки не найдены")
	}
}
