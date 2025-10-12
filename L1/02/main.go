package main

import (
	"fmt"
	"sync"
)

func raiseToThePower(wg *sync.WaitGroup, num int) {
	defer wg.Done() // (3) Уменьшаем счетчик на единицу по завершении возведения
	result := num * num
	fmt.Println(result)
}

func startCounting(nums []int) {
	var wg sync.WaitGroup // (1) Создаем счетчик задач
	for _, num := range nums {
		wg.Add(1) // (2) Увеличиваем счетчик задач для каждого числа, вызывая возведения в степень конкурентно
		go raiseToThePower(&wg, num)
	}
	wg.Wait() // (4) Блокируем завершение функции, ожидая обнуления всех задач
	fmt.Println("That's it !")
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	startCounting(nums)
}
