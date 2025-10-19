package main

import (
	"fmt"
	"sync"
	"time"
)

// Создадим структуру, которая будет хранить сам счетчик
// и мьютекс для безопасного управления счетчиком
type Counter struct {
	count int
	mutex *sync.RWMutex
}

// Метод для инкрементации счетчика, вызывается массово
func (c *Counter) IncrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Добавим некоторое условие для завершения счетчика
		if c.count == 10 {
			return
		}

		// Обернем инкрементацию в лок-анлок
		c.mutex.Lock()
		c.count++
		c.mutex.Unlock()

		time.Sleep(time.Second)
	}
}

// По завершении работы счетчика уведомляем об этом
func (c *Counter) PrintCount() {
	// Читаем данные через мьютекс, хотя это и так
	// безопасно, мы обращаемся к общей изменяемой
	// переменной, поэтому дабы соблюсти общий
	// контраст и быть уверенным, что мы получаем
	// последнее актуальное значение, используем
	// мьютекс на чтение
	c.mutex.RLock()
	fmt.Println("All goroutines are closed, count result: ", c.count)
	c.mutex.RUnlock()
}

// Пожалуй, было бы проще и компактнее реализовать
// через atomic, но мы пойдем более "явным" путем
func main() {
	var wg sync.WaitGroup

	// Создадим экземпляр структуры
	c := Counter{
		count: 0,
		mutex: &sync.RWMutex{},
	}

	fmt.Println("Start counting...")
	// Конкурентно запускаем 5 воркеров
	// через экземпляр структуры
	for range 5 {
		wg.Add(1)
		go c.IncrementCounter(&wg)
	}

	wg.Wait()
	// Уведомляем о завершении подсчета
	c.PrintCount()
}
