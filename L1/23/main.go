package main

import (
	"errors"
	"fmt"
)

// Функция pop использует дженерики для универсального взаимодейстия
// со слайсами. Принимает сам слайс, который необходимо изменить и
// индекс в формате целого числа для удаления в этом слайсе. Возвращает
// видоизмененный слайс и nil при успешном удалении элемента. При
// неудаче вернет исходный слайс и ошибку с соответствующим описанием
func pop[T any](slice []T, index int) ([]T, error) {
	// Проверяем наличие указанного индекса в слайсе
	if index > len(slice)-1 {
		return slice, errors.New("index out of range")
	}
	// Проверяем индекс на неотрицательность
	if index < 0 {
		return slice, errors.New("index must not be negative")
	}

	// "Сдвигаем" хвост слайса на место удаленного элемена
	copy(slice[index:], slice[index+1:])
	// Обрезаем дубликат последнего значения в слайсе
	return slice[:len(slice)-1], nil
}

func main() {
	myIntSlice := []int{1, 2, 3, 4, 5, 6, 7}
	myStrSlice := []string{"1", "2", "3", "4", "5", "6"}

	myIntSlice, err := pop(myIntSlice, 2)
	if err != nil {
		fmt.Println("Error removing slice element: ", err)
	}

	myStrSlice, err = pop(myStrSlice, 5)
	if err != nil {
		fmt.Println("Error removing slice element: ", err)
	}

	fmt.Printf("%v\n%v\n", myIntSlice, myStrSlice)
}
