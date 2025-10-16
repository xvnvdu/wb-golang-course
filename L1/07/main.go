package main

import (
	"fmt"
	"sync"
	"time"
)

// Опишу только первого воркера, тк в отличие от второго
// он помимо записи в мапы еще и читает из них
func worker1(wg *sync.WaitGroup, mutex *sync.RWMutex, m *sync.Map, myMap map[int]int) {
	defer wg.Done()
	number := 1
	for {
		if number%2 == 0 {
			// Явно "обернем" потенциальную опасность в конструкцию лок-анлок,
			// таким образом записать данные может только одна горутина
			mutex.Lock()
			myMap[number] = number*2 - 5
			mutex.Unlock()

			// Также обернем потенциальную опасность, но лок уже будет только на запись,
			// то есть безопасно читать данные может сразу несколько горутин
			mutex.RLock()
			fmt.Println("Classic", myMap)
			mutex.RUnlock()
		} else {
			tempMap := make(map[any]any) // Создадим мапу для безопасного вывода данных из sync map
			m.Store(number, number+3)    // Записываем данные в sync мапу, это безопасная операция

			// Также можем безопасно прочитать данные, запишем их
			// в отдельную мапу, к которой у других горутин нет доступа,
			// чтобы в дальнейшем вывести их в похожем формате
			m.Range(func(key, value any) bool {
				tempMap[key] = value
				return true
			})
			fmt.Println("Sync", tempMap)
		}

		number++
		time.Sleep(time.Second)
	}
}

func worker2(wg *sync.WaitGroup, mutex *sync.RWMutex, m *sync.Map, myMap map[int]int) {
	defer wg.Done()
	number := -1
	for {
		if number%2 == 0 {
			mutex.Lock()
			myMap[number] = number*3 - 7
			mutex.Unlock()
		} else {
			m.Store(number, number-6)
		}

		number--
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup  // Как обычно создаем счетчик и делаем с ним все что нужно :)
	var mutex sync.RWMutex // Создаем rw mutex для явного взамидействия с обычной мапой
	var m sync.Map         // Создадим также sync мапу для наглядности

	myMap := make(map[int]int)

	wg.Add(2)
	// Запустим двух воркеров, которые будут записывать данные сразу в обе мапы:
	// получаем потенциальную угрозу гонки данных
	go worker1(&wg, &mutex, &m, myMap)
	go worker2(&wg, &mutex, &m, myMap)

	wg.Wait()
}
