package main

import (
	"fmt"
	"sync"
	"time"
)

// Генерирует числа от 1 до arrLen и записывает их в первый канал
func generator(wg *sync.WaitGroup, arrLen int, ch1 chan int) {
	defer wg.Done() // (3) По завершении функции уменьшаем счетчик [1]
	for i := 1; i <= arrLen; i++ {
		ch1 <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch1) // Обязательно закрываем канал, данных больше не будет [1]
}

// Безопасно обрабатывает все значения из первого канала,
// умножая каждое на 2 и записывая результат во второй канал
func processor(wg *sync.WaitGroup, ch1, ch2 chan int) {
	defer wg.Done() // (3) По завершении функции уменьшаем счетчик [2]
	for x := range ch1 {
		ch2 <- x * 2
	}
	close(ch2) // Обязательно закрываем канал, данных больше не будет [2]
}

func main() {
	var wg sync.WaitGroup // (1) Создаем счетчик
	arrLen := 10          // Длина массива для генератора

	// Создаем два канала
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2) // (2) Увеличиаем счетчик на количество процессов
	go generator(&wg, arrLen, ch1)
	go processor(&wg, ch1, ch2)

	// Безопасно читаем из канала
	for result := range ch2 {
		fmt.Println(result)
	}

	wg.Wait() // (4) Блокируем завершение функции
}
