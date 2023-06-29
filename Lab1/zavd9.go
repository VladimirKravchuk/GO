package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	//Задание.
	//1. Пояснить результаты операций
	//Додав пояснення у звіт
	// Логічне && (I), повертає true тільки якщо всі значення мають true, в інакшому випадку false
	// Логічне || (АБО), повертає true тільки якщо хоть одне значення має true, в інакшому випадку false
	// != (Не дорівнює), повертає true якщо значення не рівні міє собою, в інакшому випадку false
	// ! повертає альтернативне значення
	// > повертає true якщо значення більше, в інакшому випадку false
	// < повертає true якщо значення менше, в інакшому випадку false
	// == повертає true якщо значення рівні, в інакшому випадку false
}
