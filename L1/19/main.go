package main

import "fmt"

// Разворачивает строку "на месте"
func reverseString(str string) string {
	runes := []rune(str) // Создаем срез рун
	// Итерируемся с краев к центру, попарно меняя
	// местами элементы, попавшие под указатели
	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

func main() {
	// Это ужасно...
	str := "Skibidi rizzler cooked my gyatt after 6-7 " +
		"crash outs fr, goofy ahh sigma moment 😭💀🙏🏻"

	fmt.Println(reverseString(str))
}
