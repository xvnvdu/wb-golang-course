package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup // (1) Создаем счетчик задач
	// (2) Объявляем контекст и функцию остановки процесса
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel() // (3) Откладываем остановку процесса до завершения функции (Ctrl + C)

	wg.Add(1) // (4) Инкрементируем счетчик на количество процессов
	go sleepingCat(ctx, &wg)

	wg.Wait() // (7) Блокируем завершение до окончания всех задач
}

// Рендерим спящего котика
func sleepingCat(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // (5) Уменьшаем счетчик процессов по завершении сна кота
	text := "      (Zzz...)"

	for {
		select {
		case <-ctx.Done(): // (8) Когда задача выполнена и канал закрывается, срабатывает условие
			awakenedCat()
			return
		default: // (6) По умолчанию кот спит
			fmt.Print("\033[2J\n\033[H") // Нужно для очистки консоли
			fmt.Println("My cat is sleeping, don't bother him pls:")
			fmt.Println(text)
			fmt.Println("^___^ /\n≽ᵔᴥᵔ≼")
			text = text[:9] + "z" + text[9:]
			time.Sleep(time.Second)
		}
	}
}

// Будим котика
func awakenedCat() {
	fmt.Print("\033[2J\n\033[H") // Снова очищаем консоль
	fmt.Println("Noooo, why did you wake him up....")
	fmt.Println("      (Damn, what?)")
	fmt.Println("^___^ /\n≽*ᴥ*≼")
}
