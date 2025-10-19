package main

import (
	"fmt"
)

// Разворачивает срез рун "на месте", принимает
// две границы, внутри которых включительно
// следует поменять местами элементы
func reverseString(runes []rune, start, finish int) {
	for l, r := start, finish; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
}

// Разворачивает уже сами слова
func reverseWords(words string) string {
	runes := []rune(words)

	// Сначала перевернем всю строку, теперь
	// слова стоят на своих местах, но задом
	// наперед, остается их только "собрать"
	// в приемлемый вид
	reverseString(runes, 0, len(runes)-1)
	// Затем, определив границы каждого
	// отдельного слова, будем использовать
	// функцию разворота строки снова, но
	// уже точечно на каждое слово
	for l, r := 0, 0; r <= len(runes); r++ {
		// Двигаем правую границу, пока не
		// найдем пробел или не достигнем
		// последнего элемента среза
		if r == len(runes) || runes[r] == ' ' {
			// Переворачиваем слово по его
			// границам без учета пробела
			reverseString(runes, l, r-1)
			l = r + 1
		}
	}
	return string(runes)
}

func main() {
	words := "Look guys I'm talking backwards"

	fmt.Println(reverseWords(words))
}
