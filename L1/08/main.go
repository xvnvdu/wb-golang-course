package main

import "fmt"

func setBit(num, i int64, targetBit bool) int64 {
	// Создаем битовую маску из единицы:
	// двигаем первый бит на i-ю позицию
	// влево, то есть выполняем i - 1 движений.
	// Через нее сможем менять ТОЛЬКО i-й бит в num
	var bitMask int64 = 1 << (i - 1)

	if targetBit {
		// Если целевой бит равен единице,
		// выполняем поразрядное сложение:
		// 0101 <- num       = 5
		// 0010 <- i         = 2
		// 0111 -> результат = 7
		num |= bitMask
	} else {
		// Если нужен ноль, выполняем сброс бита:
		// единица в битовой маске на i-ю единицу
		// (и на ноль тоже) даст ноль, оставшиеся
		// в маске нули справа примут соответствующие
		// значения битов исходного числа:
		// 1100 <- num       = 12
		// 1000 <- i         = 8
		// 0100 <- результат = 4
		num &^= bitMask
	}
	return num
}

func main() {
	// Задаем число num и i-й бит, который хотим изменить в num
	var num, i int64 = 1000, 10
	var targetBit bool = false // Указываем целевое значение бита 0/1

	result := setBit(num, i, targetBit)

	fmt.Printf("Old value: %d (Binary form: %b)\n", num, num)
	fmt.Printf("New value: %d (Binary form: %b)\n", result, result)
}
