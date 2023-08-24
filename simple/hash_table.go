package simple_structs

import "fmt"

func HashTable() {
	// Инициализация хеш-таблицы
	m := make(map[string]string)

	// Добавление элементов
	m["name"] = "John"
	m["job"] = "Engineer"

	fmt.Printf("hash table: %#v\n", m)
}
