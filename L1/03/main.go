package main

import (
	"fmt"
	"strconv"
	"time"
)

// Воркер работает фоном и принимает данные из канала
func worker(i int, ch chan string) {
	for {
		data := <-ch
		fmt.Printf("Alien hunter (worker) №%d reporting: %s\n", i+1, data)
	}
}

// Создает и запускает N воркеров, каждый из которых работает конкурентно
func makeWorkers(n int, ch chan string) {
	for i := range n {
		go worker(i, ch)
	}
}

func main() {
	workersAmount := 5
	ch := make(chan string) // (1) Создаем канал

	counter := 1
	go func() { // (2) Запускаем постоянную запись в канал фоном
		for {
			str := "Exobiological creature number " + strconv.Itoa(counter) + " neutralized."
			ch <- str

			counter++
			time.Sleep(2 * time.Second)
		}
	}()

	makeWorkers(workersAmount, ch) // (3) Создаем N воркеров
	select {}                      // (4) Блокируем главную горутину, пока продюсер и консьюмеры работают
}
