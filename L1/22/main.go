package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Структура Calc хранит два больших числа
type Calc struct {
	x *big.Int
	y *big.Int
}

// Возвращает результат x*y между числами внутри Calc
func (c *Calc) Multiply() *big.Int {
	result := new(big.Int)
	result.Mul(c.x, c.y)

	return result
}

// Возвращает результат x/y между числами внутри Calc
func (c *Calc) Divide() *big.Int {
	result := new(big.Int)
	result.Div(c.x, c.y)

	return result
}

// Возвращает результат x+y между числами внутри Calc
func (c *Calc) Add() *big.Int {
	result := new(big.Int)
	result.Add(c.x, c.y)

	return result
}

// Возвращает результат x-y между числами внутри Calc
func (c *Calc) Subtract() *big.Int {
	result := new(big.Int)
	result.Sub(c.x, c.y)

	return result
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// Создаем два числа, которые не
	// поместились бы в 64 бита
	a.SetString(createBigNumber(346, 40), 10)
	b.SetString(createBigNumber(45, 20), 10)

	calc := &Calc{
		x: a,
		y: b,
	}

	fmt.Println("Multiplication result: ", calc.Multiply())
	fmt.Println("Division result:       ", calc.Divide())
	fmt.Println("Addition result:       ", calc.Add())
	fmt.Println("Subtraction result:    ", calc.Subtract())
}

// createBigNumber генерирует целое круглое число и возвращает
// его в строковом выражении.
//
// baseNum является "основой" генерируемого числа.
// addDigits отражает количество разрядов, добавляемых к основе.
//
// Пример:
//
// myNumber := createBigNumber(67, 10)
// fmt.Println(myNumber)    // 670000000000
func createBigNumber(baseNum, addDigits int) string {
	base := strconv.Itoa(baseNum)

	for range addDigits {
		base += "0"
	}
	return base
}
