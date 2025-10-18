package main

import "fmt"

// Могу предположить, что в данном кейсе использовать
// v any на вход было бы более коректно с точки
// зрения синтаксиса и читаемости, тк это алиас
// пустого интерфейса, точно так же как chan any
func identifyType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Переменная %v: %T\n", v, v)
	case string:
		fmt.Printf("Переменная %v: %T\n", v, v)
	case bool:
		fmt.Printf("Переменная %v: %T\n", v, v)
	case chan any:
		// В этом случае chan any будет принимать
		// ТОЛЬКО каналы с переменными типа any, для
		// каналов с конкретными типами переменных
		// string, int, bool, map[string]int и тд
		// пришлось бы реализовывать каждый кейс
		// по отдельности
		fmt.Printf("Переменная %v: %T\n", v, v)
	}
}

func main() {
	var integer int = 100
	var str string = "14th task"
	var boolean bool = true
	var ch chan any = make(chan any)

	identifyType(integer)
	identifyType(str)
	identifyType(boolean)
	identifyType(ch)
}
