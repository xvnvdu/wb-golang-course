package main

import "fmt"

func binarySearch(sortedSlice []int, target int) int {
	l, r := 0, len(sortedSlice)-1 // Определим границы массива

	// Итерируемся по массиву пока границы не сомкнутся
	for l < r {
		// Определяем середину массива,
		// сохраняя исходные индексы
		mid := (l + r) / 2
		// Вернем "средний" индекс, если
		// по нему нашли искомое значение
		if target == sortedSlice[mid] {
			return mid
		} else if target > sortedSlice[mid] {
			// Если значение правее текущей середины,
			// рассматриваем только правую часть,
			// двигаем левую границу на следующий
			// после центрального элемент
			l = mid + 1
		} else {
			// Если искомое значение левее, текущая
			// середина становится правой границей,
			// повторим процесс пока границы не сомкнутся
			r = mid
		}
	}
	// Если границы сомкнулись, но функция еще не вернула значение,
	// значит целевого элемента нет в массиве
	return -1
}

func main() {
	var target int = 9
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	index := binarySearch(sortedSlice, target)
	if index == -1 {
		fmt.Println("Target value was not found")
	} else {
		fmt.Println("Target element found, index: ", index)
	}
}
