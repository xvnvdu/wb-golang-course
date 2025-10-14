package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// Горутина, завершающая свою работу при выполнении условия
func workerWithCondition(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		if i == 4 {
			fmt.Printf("workerWithCondition: retired on his pension after %d seconds of hard work\n", i)
			return
		}
		i++
		time.Sleep(time.Second)
	}
}

// Горутина, завершающая свою работу по уведомлению из канала
func workerWithChanel(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		value := <-ch
		if value >= 144 {
			fmt.Printf("workerWithChanel: is sick of reading fibonacci numbers, stopped at %d\n", value)
			return
		}
	}
}

// Горутина, завершающая свою работу через контекст (по таймеру)
func workerWithContext(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	value := 2
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("workerWithContext: was stopped by timeout while raising number 2 to a power, last value is %d\n", value)
			return
		default:
			value = value * 2
			time.Sleep(time.Second)
		}
	}
}

// Горутина, завершающая свою работу через Goexit()
func workerWithGoexit(wg *sync.WaitGroup) {
	defer wg.Done()
	now := time.Now()
	lifeTime := 5 * time.Second
	for {
		if time.Since(now) >= lifeTime {
			fmt.Printf("workerWithGoexit: lived only for %v seconds, was terminated by Goexit, rest in peace, buddy\n", lifeTime.Seconds())
			runtime.Goexit()
		}
	}
}

// Горутина с бесконечной работой, завершается только через Ctrl + C
func workerWithInfiniteCycle(wg *sync.WaitGroup, ctxSigint context.Context, ch chan int) {
	defer wg.Done()
	fibNumber := 0
	for {
		select {
		case <-ctxSigint.Done():
			fmt.Printf("\nworkerWithInfiniteCycle: received SIGINT signal, last fibonacci number was %d\n", fibNumber)
			return
		case fibNumber = <-ch:
			time.Sleep(time.Second)
		}
	}
}

// Небольшой продюсер, записывающий в канал последовательность Фибоначчи
func producer(ch chan int) {
	a, b := 0, 1
	for {
		fibNumber := a + b
		ch <- fibNumber

		a = b
		b = fibNumber

		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup // (1) Создаем счетчик
	ch := make(chan int)

	// (2) Объявляем два контекста: один с таймаутом, второй на отмену через SIGINT
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctxSigint, cancelSigint := signal.NotifyContext(context.Background(), syscall.SIGINT)

	// (3) Откладываем остановку процессов до завершения функции
	defer cancel()
	defer cancelSigint()

	// (4) Инкрементируем счетчик на количество воркеров и сразу их запускаем
	wg.Add(5)
	go workerWithCondition(&wg)
	go workerWithChanel(&wg, ch)
	go workerWithContext(&wg, ctx)
	go workerWithGoexit(&wg)
	go workerWithInfiniteCycle(&wg, ctxSigint, ch)

	go producer(ch) // (5) Запускаем продюсера

	fmt.Println("All goroutines were started")

	wg.Wait() // (6) Блокируем завершение до окончания задач
}
