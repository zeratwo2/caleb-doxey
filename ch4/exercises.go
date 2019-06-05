//что выведет этот код?
package main

import "fmt"

func main() {
	i := 10

	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
	cisla()
	rofl()
}

//выводит слово ""small" потому что "i" не больше 10 и работает оператор else

//Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …).

func cisla() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {

			fmt.Println(i, "deleatsea na 3")
		}
	}
}

/*Напишите программу, которая выводит числа от 1 до 100.
Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz»,
 а для кратных как трём, так и пяти — «FizzBuzz». */
func rofl() {
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")

		}

	}
}
