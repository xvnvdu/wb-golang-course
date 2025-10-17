package main

import (
	"fmt"
)

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree", "cat"}
	checkMap := make(map[string]bool) // Так же используем мапу
	resultSlice := []string{}         // Пустой слайс для результатов

	// Можем решить эту задачу за единственную
	// итерацию по исходному слайсу - O(n)
	for _, word := range slice {
		// Проверим, существует ли в принципе в мапе такая строка,
		// если нет, значит мы ее еще не встречали - надо записать
		if _, ok := checkMap[word]; !ok {
			resultSlice = append(resultSlice, word)
			// Добавляем строку в мапу, чтобы больше ее не учитывать
			checkMap[word] = true
		}
	}
	// Таким образом сохраняем и порядок строк,
	// в котором они встречались первый раз
	fmt.Println(resultSlice)
}
