package main

import "fmt"

func classicOperations(a, b int) {
	a += b
	b = a - b
	a -= b
	fmt.Printf("After classicOperations func a=%d, b=%d\n", a, b)
}

// A ^ B ^ B = A
func bitwiseOperations(c, d int) {
	c ^= d    // c = (c ^ d)
	d = c ^ d // Присваиваем d исходное значение c, используя (c ^ d) ^ d = c
	c ^= d    // Возвращаем саму переменную c в исходное значение (c ^ d) ^ d = c
	fmt.Printf("After bitwiseOperations func c=%d, d=%d\n", c, d)
}

func main() {
	a := 5
	b := 3
	c := 4
	d := 6
	fmt.Printf("Initially we had a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)

	classicOperations(a, b)
	bitwiseOperations(c, d)
}
