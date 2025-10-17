package main

import "fmt"

func main() {
	myMap := make(map[int][]float64) // Создаем мапу формата "диапазон: [колебания]"
	// Колебания температур закинем в слайс
	tempSwings := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Итерируемся по всему слайсу
	for _, temp := range tempSwings {
		// Ключ - целое усеченное (truncated) число
		key := int(temp/10) * 10
		// К соответствующему ключу в слайс кидаем температуру
		myMap[key] = append(myMap[key], temp)
	}
	fmt.Println(myMap)
}
