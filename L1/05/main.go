package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const N = 5

func worker(ctx context.Context, ch chan any, wg *sync.WaitGroup) {
	defer wg.Done() // (8) Откладываем обнуление количества задач до закрытия канала и выхода воркера
	for {
		select {
		case <-ctx.Done(): // (9) Когда время выходит, канал закрывается и воркер прекращает слушать
			fmt.Println("Yup, that seems to be working")
			return
		case data := <-ch:
			fmt.Printf("Proof of 5th task completion in: %d\n", data)
		}
	}
}

func main() {
	var wg sync.WaitGroup // (1) Создаем счетчик

	ch := make(chan any)
	// (2) ОБъявляем контекст с таймаутом (вместо time.After, но там был бы свой канал, проверяли бы по нему)
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)

	defer cancel() // (3) Отложим закрытие канала до завершения функции

	go func() { // (4) Анонимный продюсер
		i := 5
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- i
				i--
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Add(1)               // (5) Инкрементируем количество процессов на один
	go worker(ctx, ch, &wg) // (6) Запускаем воркера

	wg.Wait() // (7) Блокируем завершение функции до окончания работы
}
